<template>
  <div>
    <div class="container mx-auto px-4 py-8">
      <!-- Back Button -->
      <button
        @click="$router.back()"
        class="flex items-center gap-2 text-gray-600 hover:text-purple-600 mb-6 transition-colors group"
      >
        <svg class="w-5 h-5 transform group-hover:-translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
        </svg>
        <span class="font-medium">Back</span>
      </button>

      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
      </div>

      <div v-else-if="event" class="max-w-4xl mx-auto">
        <!-- Image Carousel -->
        <div v-if="allImages.length > 0" class="relative h-96 rounded-lg overflow-hidden mb-8 group">
          <!-- Main Image -->
          <img 
            :src="currentImage" 
            :alt="event.title" 
            class="w-full h-full object-cover transition-opacity duration-300"
          />
          
          <!-- Navigation Arrows (show on hover if multiple images) -->
          <template v-if="allImages.length > 1">
            <button
              @click="prevImage"
              class="absolute left-4 top-1/2 -translate-y-1/2 bg-black/50 hover:bg-black/70 text-white p-3 rounded-full transition-all opacity-0 group-hover:opacity-100"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
              </svg>
            </button>
            <button
              @click="nextImage"
              class="absolute right-4 top-1/2 -translate-y-1/2 bg-black/50 hover:bg-black/70 text-white p-3 rounded-full transition-all opacity-0 group-hover:opacity-100"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </button>
          </template>
          
          <!-- Image Counter -->
          <div v-if="allImages.length > 1" class="absolute top-4 right-4 bg-black/70 text-white px-3 py-1 rounded-full text-sm">
            {{ currentImageIndex + 1 }} / {{ allImages.length }}
          </div>
          
          <!-- Thumbnail Navigation (show if multiple images) -->
          <div v-if="allImages.length > 1" class="absolute bottom-4 left-1/2 -translate-x-1/2 flex gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
            <button
              v-for="(img, index) in allImages"
              :key="index"
              @click="goToImage(index)"
              class="w-3 h-3 rounded-full transition-all"
              :class="index === currentImageIndex ? 'bg-white scale-125' : 'bg-white/50 hover:bg-white/80'"
            />
          </div>
        </div>
        
        <!-- Placeholder if no images -->
        <div v-else class="relative h-96 rounded-lg overflow-hidden mb-8 bg-gradient-to-br from-purple-100 via-pink-100 to-orange-100 flex items-center justify-center">
          <svg class="w-32 h-32 text-purple-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" 
                  d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
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

          <!-- Price and Ticket Availability -->
          <div class="mb-6">
            <p class="text-3xl font-bold text-primary mb-2">
              {{ event.price === 0 ? 'FREE' : `${formatPrice(event.price)} ETB` }}
            </p>
            <div v-if="ticket" class="flex items-center gap-2">
              <span 
                v-if="ticketsAvailable" 
                class="inline-flex items-center gap-1 text-sm font-medium text-green-700 bg-green-50 px-3 py-1 rounded-full"
              >
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" />
                </svg>
                {{ ticketsRemaining }} tickets remaining
              </span>
              <span 
                v-else 
                class="inline-flex items-center gap-1 text-sm font-medium text-red-700 bg-red-50 px-3 py-1 rounded-full"
              >
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
                </svg>
                SOLD OUT
              </span>
            </div>
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
            <div v-if="event.latitude && event.longitude">
              <ClientOnly>
                <EventMap 
                  :events="[event]" 
                  :center="[event.latitude, event.longitude]"
                />
              </ClientOnly>
            </div>
            <div v-else class="h-[300px] bg-gray-100 rounded-lg flex items-center justify-center text-gray-500">
              <div class="text-center">
                <svg class="w-16 h-16 mx-auto mb-2 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                        d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                        d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                <p>Location coordinates not available</p>
              </div>
            </div>
          </div>

          <!-- Action Buttons -->
          <div class="flex gap-4">
            <!-- Get Tickets Button - Only show if tickets available -->
            <template v-if="ticketsAvailable">
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
            </template>
            
            <!-- Sold Out Message -->
            <div
              v-else
              class="flex-1 bg-gray-200 text-gray-600 text-center py-3 rounded-lg font-semibold cursor-not-allowed"
            >
              Sold Out - No Tickets Available
            </div>
            
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
    </div>
  </div>
</template>

<script setup lang="ts">
import { useQuery } from '@vue/apollo-composable'
import { gql } from '@apollo/client/core'

const route = useRoute()
const eventSlug = route.params.id // Param name stays 'id' but contains slug

const { isAuthenticated } = useAuth()

const GET_EVENT = gql`
  query GetEvent($slug: String!) {
    events(where: { slug: { _eq: $slug } }, limit: 1) {
      id
      slug
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
      tickets {
        id
        quantity_total
        quantity_sold
      }
    }
  }
`

const { result, loading } = useQuery(
  GET_EVENT, 
  { slug: eventSlug },
  {
    context: {
      forceAnonymous: true // Use anonymous role for public event viewing
    }
  }
)

const event = computed(() => result.value?.events?.[0])
const eventId = computed(() => event.value?.id) // Extract ID for interactions
const { isBookmarked, isFollowing, toggleBookmark, toggleFollow } = useEventInteractions(eventId)

// Ticket availability computed properties
const ticket = computed(() => event.value?.tickets?.[0])
const ticketsRemaining = computed(() => {
  if (!ticket.value) return 0
  return ticket.value.quantity_total - ticket.value.quantity_sold
})
const ticketsAvailable = computed(() => ticketsRemaining.value > 0)
const soldOut = computed(() => ticketsRemaining.value === 0)

// All images for carousel
const allImages = computed(() => event.value?.event_images || [])
const currentImageIndex = ref(0)
const currentImage = computed(() => {
  if (allImages.value.length === 0) return null
  return allImages.value[currentImageIndex.value]?.url
})

const featuredImage = computed(() => {
  const images = event.value?.event_images || []
  const featured = images.find((img: any) => img.is_featured)
  return featured?.url || images[0]?.url || null
})

const nextImage = () => {
  if (currentImageIndex.value < allImages.value.length - 1) {
    currentImageIndex.value++
  } else {
    currentImageIndex.value = 0 // Loop to first
  }
}

const prevImage = () => {
  if (currentImageIndex.value > 0) {
    currentImageIndex.value--
  } else {
    currentImageIndex.value = allImages.value.length - 1 // Loop to last
  }
}

const goToImage = (index: number) => {
  currentImageIndex.value = index
}

// Keyboard navigation
onMounted(() => {
  const handleKeyPress = (e: KeyboardEvent) => {
    if (allImages.value.length <= 1) return
    if (e.key === 'ArrowLeft') prevImage()
    if (e.key === 'ArrowRight') nextImage()
  }
  window.addEventListener('keydown', handleKeyPress)
  
  onUnmounted(() => {
    window.removeEventListener('keydown', handleKeyPress)
  })
})

// Dynamic page title based on event name
useHead({
  title: computed(() => event.value?.title || 'Event Details'),
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
  navigateTo(`/payment/checkout/${event.value.slug}`)
}
</script>
