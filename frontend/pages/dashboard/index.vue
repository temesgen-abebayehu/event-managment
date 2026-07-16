<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow">
      <div class="container mx-auto px-4 py-6">
        <div class="flex items-center justify-between">
          <NuxtLink to="/" class="text-2xl font-bold text-primary">
            EventHub Ethiopia
          </NuxtLink>
          <div class="flex items-center gap-4">
            <NuxtLink to="/dashboard/bookmarks" class="text-gray-600 hover:text-primary">
              Bookmarks
            </NuxtLink>
            <NuxtLink to="/dashboard/follows" class="text-gray-600 hover:text-primary">
              Following
            </NuxtLink>
            <NuxtLink to="/dashboard/tickets" class="text-gray-600 hover:text-primary">
              Tickets
            </NuxtLink>
            <span class="text-gray-600">{{ user?.full_name }}</span>
            <button @click="logout" class="text-gray-600 hover:text-primary">
              Logout
            </button>
          </div>
        </div>
      </div>
    </header>

    <main class="container mx-auto px-4 py-8">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold">My Events</h1>
        <NuxtLink
          to="/events/create"
          class="bg-primary text-white px-6 py-3 rounded-lg font-semibold hover:bg-primary-dark transition"
        >
          + Create Event
        </NuxtLink>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
      </div>

      <!-- Empty State -->
      <div v-else-if="myEvents.length === 0" class="text-center py-12">
        <svg class="w-24 h-24 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
        </svg>
        <h3 class="text-xl font-semibold text-gray-700 mb-2">No events yet</h3>
        <p class="text-gray-500 mb-6">Create your first event to get started</p>
        <NuxtLink
          to="/events/create"
          class="inline-block bg-primary text-white px-6 py-3 rounded-lg font-semibold hover:bg-primary-dark transition"
        >
          Create Event
        </NuxtLink>
      </div>

      <!-- Events List -->
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="event in myEvents"
          :key="event.id"
          class="bg-white rounded-lg shadow-md overflow-hidden"
        >
          <!-- Image -->
          <div class="relative h-48">
            <img
              v-if="event.event_images[0]?.url"
              :src="event.event_images[0].url"
              :alt="event.title"
              class="w-full h-full object-cover"
            />
            <div v-else class="w-full h-full bg-gray-200 flex items-center justify-center">
              <span class="text-gray-400">No image</span>
            </div>
          </div>

          <!-- Content -->
          <div class="p-4">
            <h3 class="text-lg font-semibold mb-2">{{ event.title }}</h3>
            <p class="text-sm text-gray-600 mb-2">
              {{ formatDate(event.event_date) }}
            </p>
            <p class="text-sm text-gray-600 mb-4">{{ event.venue }}</p>

            <!-- Actions -->
            <div class="flex gap-2">
              <NuxtLink
                :to="`/events/edit/${event.id}`"
                class="flex-1 bg-blue-500 text-white text-center py-2 rounded-lg hover:bg-blue-600 transition"
              >
                Edit
              </NuxtLink>
              <button
                @click="confirmDelete(event)"
                class="flex-1 bg-red-500 text-white py-2 rounded-lg hover:bg-red-600 transition"
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="showDeleteModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click="showDeleteModal = false"
    >
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4" @click.stop>
        <h3 class="text-xl font-bold mb-4">Delete Event?</h3>
        <p class="text-gray-600 mb-6">
          Are you sure you want to delete "{{ eventToDelete?.title }}"? This action cannot be undone.
        </p>
        <div class="flex gap-4">
          <button
            @click="handleDelete"
            :disabled="deleting"
            class="flex-1 bg-red-500 text-white py-2 rounded-lg hover:bg-red-600 transition disabled:opacity-50"
          >
            {{ deleting ? 'Deleting...' : 'Delete' }}
          </button>
          <button
            @click="showDeleteModal = false"
            class="flex-1 bg-gray-200 text-gray-700 py-2 rounded-lg hover:bg-gray-300 transition"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useMutation, useQuery } from '@vue/apollo-composable'
import { gql } from '@apollo/client/core'

definePageMeta({
  middleware: 'auth'
})

const { user, logout } = useAuth()

// GraphQL Queries
const MY_EVENTS = gql`
  query MyEvents($user_id: uuid!) {
    events(where: { user_id: { _eq: $user_id } }, order_by: { created_at: desc }) {
      id
      title
      description
      venue
      event_date
      price
      event_images(where: { is_featured: { _eq: true } }, limit: 1) {
        url
      }
    }
  }
`

const DELETE_EVENT = gql`
  mutation DeleteEvent($id: uuid!) {
    delete_events_by_pk(id: $id) {
      id
    }
  }
`

const { result, loading, refetch } = useQuery(MY_EVENTS, { user_id: user.value?.id })
const myEvents = computed(() => result.value?.events || [])

const { mutate: deleteEventMutation } = useMutation(DELETE_EVENT)

const showDeleteModal = ref(false)
const eventToDelete = ref(null)
const deleting = ref(false)

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('en-ET', {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

const confirmDelete = (event: any) => {
  eventToDelete.value = event
  showDeleteModal.value = true
}

const handleDelete = async () => {
  if (!eventToDelete.value) return

  deleting.value = true
  try {
    await deleteEventMutation({ id: eventToDelete.value.id })
    showDeleteModal.value = false
    eventToDelete.value = null
    refetch()
  } catch (error) {
    console.error('Failed to delete event:', error)
  } finally {
    deleting.value = false
  }
}
</script>
