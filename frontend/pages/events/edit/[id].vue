<template>
  <div>
    <div class="container mx-auto px-4 py-8 max-w-3xl">
      <div v-if="loadingEvent" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
      </div>

      <div v-else-if="event">
        <div class="flex items-center justify-between mb-8">
          <h1 class="text-3xl font-bold">Edit Event</h1>
          <button @click="$router.back()" class="text-gray-600 hover:text-primary text-sm font-medium">← Back</button>
        </div>

        <form @submit.prevent="handleSubmit" class="bg-white rounded-lg shadow-lg p-8">
          <!-- Title -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Event Title *
            </label>
            <input
              v-model="form.title"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
            />
          </div>

          <!-- Category -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Category *
            </label>
            <select
              v-model="form.category_id"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
            >
              <option value="">Select category</option>
              <option v-for="cat in categories" :key="cat.id" :value="cat.id">
                {{ cat.name }}
              </option>
            </select>
          </div>

          <!-- Description -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Description *
            </label>
            <textarea
              v-model="form.description"
              required
              rows="6"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
            />
          </div>

          <!-- Current Images -->
          <div v-if="existingImages.length > 0" class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Current Images
            </label>
            <div class="grid grid-cols-3 gap-4">
              <div v-for="img in existingImages" :key="img.id" class="relative">
                <img :src="img.url" class="w-full h-32 object-cover rounded-lg" />
                <span v-if="img.is_featured" class="absolute top-1 left-1 bg-primary text-white text-xs px-2 py-1 rounded">
                  Featured
                </span>
              </div>
            </div>
          </div>

          <!-- Venue -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Venue Name *
            </label>
            <input
              v-model="form.venue"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
            />
          </div>

          <!-- Address -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Full Address *
            </label>
            <input
              v-model="form.address"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
            />
          </div>

          <!-- Location -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Location Coordinates *
            </label>
            <div class="grid grid-cols-2 gap-4">
              <input
                v-model.number="form.latitude"
                type="number"
                step="0.000001"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
              />
              <input
                v-model.number="form.longitude"
                type="number"
                step="0.000001"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
              />
            </div>
          </div>

          <!-- Date & Time -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Event Date & Time *
            </label>
            <input
              v-model="form.event_date"
              type="datetime-local"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
            />
          </div>

          <!-- Price -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Ticket Price (ETB) *
            </label>
            <input
              v-model.number="form.price"
              type="number"
              step="0.01"
              min="0"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
            />
          </div>

          <!-- Error Message -->
          <div v-if="error" class="mb-6 p-3 bg-red-50 text-red-700 rounded-lg text-sm">
            {{ error }}
          </div>

          <!-- Submit -->
          <div class="flex gap-4">
            <button
              type="submit"
              :disabled="loading"
              class="flex-1 bg-primary text-white py-3 rounded-lg font-semibold hover:bg-primary-dark transition disabled:opacity-50"
            >
              {{ loading ? 'Updating...' : 'Update Event' }}
            </button>
            <button
              type="button"
              @click="$router.back()"
              class="px-6 py-3 border border-gray-300 rounded-lg hover:bg-gray-50 transition"
            >
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useQuery, useMutation } from '@vue/apollo-composable'
import { gql } from '@apollo/client/core'

definePageMeta({
  middleware: 'auth'
})

useHead({ title: 'Edit Event' })

const route = useRoute()
const eventId = route.params.id
const { categories } = useCategories()

const GET_EVENT = gql`
  query GetEvent($id: uuid!) {
    events_by_pk(id: $id) {
      id
      title
      category_id
      description
      venue
      address
      latitude
      longitude
      event_date
      price
      event_images {
        id
        url
        public_id
        is_featured
      }
    }
  }
`

const UPDATE_EVENT = gql`
  mutation UpdateEvent($id: uuid!, $set: events_set_input!) {
    update_events_by_pk(pk_columns: { id: $id }, _set: $set) {
      id
    }
  }
`

const { result: eventResult, loading: loadingEvent } = useQuery(GET_EVENT, { id: eventId })
const event = computed(() => eventResult.value?.events_by_pk)
const existingImages = computed(() => event.value?.event_images || [])

const { mutate: updateEventMutation } = useMutation(UPDATE_EVENT)

const form = ref({
  title: '',
  category_id: '',
  description: '',
  venue: '',
  address: '',
  latitude: 9.005401,
  longitude: 38.763611,
  event_date: '',
  price: 0,
})

const loading = ref(false)
const error = ref('')

// Load event data into form
watch(event, (newEvent) => {
  if (newEvent) {
    form.value = {
      title: newEvent.title,
      category_id: newEvent.category_id,
      description: newEvent.description,
      venue: newEvent.venue,
      address: newEvent.address,
      latitude: newEvent.latitude,
      longitude: newEvent.longitude,
      event_date: new Date(newEvent.event_date).toISOString().slice(0, 16),
      price: newEvent.price,
    }
  }
})

const handleSubmit = async () => {
  loading.value = true
  error.value = ''

  try {
    await updateEventMutation({
      id: eventId,
      set: form.value,
    })

    navigateTo('/dashboard')
  } catch (err: any) {
    error.value = 'Failed to update event'
  } finally {
    loading.value = false
  }
}
</script>
