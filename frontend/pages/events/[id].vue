<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow">
      <div class="container mx-auto px-4 py-6">
        <div class="flex items-center justify-between">
          <NuxtLink to="/" class="text-2xl font-bold text-primary">
            EventHub Ethiopia
          </NuxtLink>
          <button @click="$router.back()" class="text-gray-600 hover:text-primary">
            ← Back
          </button>
        </div>
      </div>
    </header>

    <main class="container mx-auto px-4 py-8">
      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
      </div>

      <div v-else-if="event" class="max-w-4xl mx-auto">
        <!-- Featured Image -->
        <div class="relative h-96 rounded-lg overflow-hidden mb-8">
          <img 
            v-if="featuredImage" 
            :src="featuredImage" 
            :alt="event.title" 
            class="w-full h-full object-cover"
          />
          <div v-else class="w-full h-full bg-gray-200 flex items-center justify-center">
            <span class="text-gray-400 text-xl">No image</span>
          </div>
        </div>

        <!-- Event Info -->
        <div class="bg-white rounded-lg shadow-lg p-8">
          <!-- Category -->
          <span class="inline-block bg-blue-100 text-blue-800 px-3 py-1 rounded text-sm font-medium mb-4">
            {{ event.category.name }}
          </span>

          <!-- Title -->
          <h1 class="text-4xl font-bold mb-4">{{ event.title }}</h1>

          <!-- Organizer -->
          <p class="text-gray-600 mb-6">
            Organized by <span class="font-semibold">{{ event.user.full_name }}</span>
          </p>

          <!-- Date & Time -->
          <div class="flex items-center gap-2 mb-4">
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                    d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
            <span class="text-gray-700">{{ formatDate(event.event_date) }}</span>
          </div>

          <!-- Location -->
          <div class="flex items-center gap-2 mb-6">
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                    d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                    d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            <div>
              <p class="font-semibold">{{ event.venue }}</p>
              <p class="text-sm text-gray-600">{{ event.address }}</p>
            </div>
          </div>

          <!-- Price -->
          <div class="mb-6">
            <p class="text-3xl font-bold text-primary">
              {{ event.price === 0 ? 'FREE' : `${formatPrice(event.price)} ETB` }}
            </p>
          </div>

          <!-- Tags -->
          <div v-if="event.event_tags && event.event_tags.length > 0" class="mb-6">
            <h3 class="text-lg font-semibold mb-2">Tags</h3>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="eventTag in event.event_tags"
                :key="eventTag.tag.id"
                class="inline-block bg-blue-100 text-blue-800 px-3 py-1 rounded-full text-sm"
              >
                #{{ eventTag.tag.name }}
              </span>
            </div>
          </div>

          <!-- Description -->
          <div class="mb-6">
            <h2 class="text-2xl font-semibold mb-3">About this event</h2>
            <p class="text-gray-700 whitespace-pre-line">{{ event.description }}</p>
          </div>

          <!-- Map -->
          <div class="mb-6">
            <h2 class="text-2xl font-semibold mb-3">Location</h2>
            <ClientOnly>
              <EventMap 
                :events="[event]" 
                :center="[event.latitude, event.longitude]"
              />
            </ClientOnly>
          </div>

          <!-- Action Buttons -->
          <div class="flex gap-4">
            <NuxtLink 
              v-if="!isAuthenticated"
              to="/auth/login"
              class="flex-1 bg-primary text-white text-center py-3 rounded-lg font-semibold hover:bg-primary-dark transition"
            >
              Get Tickets
            </NuxtLink>
            <button
              v-else
              @click="handleBuyTicket"
              class="flex-1 bg-primary text-white py-3 rounded-lg font-semibold hover:bg-primary-dark transition"
            >
              Get Tickets
            </button>
            
            <button 
              @click="toggleBookmark"
              class="px-6 py-3 border border-gray-300 rounded-lg hover:bg-gray-50 transition"
              :class="{ 'bg-yellow-50 border-yellow-500': isBookmarked }"
            >
              <svg class="w-6 h-6" :class="{ 'fill-yellow-500': isBookmarked }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                      d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
              </svg>
            </button>
            
            <button
              @click="toggleFollow"
              class="px-6 py-3 border border-gray-300 rounded-lg hover:bg-gray-50 transition"
              :class="{ 'bg-red-50 border-red-500': isFollowing }"
            >
              <svg class="w-6 h-6" :class="{ 'fill-red-500': isFollowing }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                      d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <div v-else class="text-center py-12">
        <p class="text-gray-600">Event not found</p>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { useQuery } from '@vue/apollo-composable'
import { gql } from '@apollo/client/core'

const route = useRoute()
const eventId = route.params.id

const { isAuthenticated } = useAuth()
const { isBookmarked, isFollowing, toggleBookmark, toggleFollow } = useEventInteractions(eventId as string)

const GET_EVENT = gql`
  query GetEvent($id: uuid!) {
    events_by_pk(id: $id) {
      id
      title
      description
      venue
      address
      latitude
      longitude
      price
      event_date
      created_at
      category {
        id
        name
        slug
      }
      event_images {
        url
        is_featured
      }
      event_tags {
        tag {
          id
          name
        }
      }
      user {
        full_name
      }
    }
  }
`

const { result, loading } = useQuery(GET_EVENT, { id: eventId })

const event = computed(() => result.value?.events_by_pk)
const featuredImage = computed(() => {
  const images = event.value?.event_images || []
  const featured = images.find((img: any) => img.is_featured)
  return featured?.url || images[0]?.url || null
})

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('en-ET').format(price)
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('en-ET', {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}

const handleBuyTicket = () => {
  navigateTo(`/payment/checkout/${eventId}`)
}
</script>
