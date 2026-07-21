<template>
  <article 
    class="group bg-white rounded-2xl shadow-sm hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-2 cursor-pointer overflow-hidden border border-gray-100"
    @click="navigateTo(`/events/${event.id}`)"
  >
    <!-- Image Container with Overlay -->
    <div class="relative h-56 overflow-hidden bg-gradient-to-br from-gray-100 to-gray-200">
      <LazyImage
        v-if="featuredImage"
        :src="featuredImage"
        :alt="event.title"
        className="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
        :showPlaceholder="true"
      />
      <div v-else class="w-full h-full bg-gradient-to-br from-purple-100 via-pink-100 to-orange-100 flex items-center justify-center">
        <svg class="w-20 h-20 text-purple-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" 
                d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
        </svg>
      </div>
      
      <!-- Gradient Overlay -->
      <div class="absolute inset-0 bg-gradient-to-t from-black/60 via-black/0 to-black/0 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
      
      <!-- Price Badge - Modern floating design -->
      <div class="absolute top-4 right-4 z-10">
        <div class="bg-white/95 backdrop-blur-sm px-4 py-2 rounded-full shadow-lg">
          <span class="font-bold text-purple-600">
            {{ event.price === 0 ? 'FREE' : `${formatPrice(event.price)} ETB` }}
          </span>
        </div>
      </div>

      <!-- Date Badge -->
      <div class="absolute top-4 left-4 z-10">
        <div class="bg-white rounded-xl p-2 shadow-lg text-center min-w-[60px]">
          <div class="text-xs font-semibold text-gray-500 uppercase">{{ formatMonth(event.event_date) }}</div>
          <div class="text-2xl font-bold text-gray-900">{{ formatDay(event.event_date) }}</div>
        </div>
      </div>
    </div>

    <!-- Content -->
    <div class="p-5">
      <!-- Category Badge -->
      <div class="mb-3">
        <span class="inline-flex items-center gap-1 bg-purple-50 text-purple-700 text-xs font-semibold px-3 py-1.5 rounded-full">
          <svg class="w-3.5 h-3.5" fill="currentColor" viewBox="0 0 20 20">
            <path d="M2 6a2 2 0 012-2h6a2 2 0 012 2v8a2 2 0 01-2 2H4a2 2 0 01-2-2V6zM14.553 7.106A1 1 0 0014 8v4a1 1 0 00.553.894l2 1A1 1 0 0018 13V7a1 1 0 00-1.447-.894l-2 1z" />
          </svg>
          {{ event.category.name }}
        </span>
      </div>

      <!-- Title with better typography -->
      <h3 class="text-xl font-bold mb-2 line-clamp-2 text-gray-900 group-hover:text-purple-600 transition-colors leading-tight">
        {{ event.title }}
      </h3>

      <!-- Description -->
      <p class="text-gray-600 text-sm mb-4 line-clamp-2 leading-relaxed">
        {{ event.description }}
      </p>

      <!-- Tags - Modern chip design -->
      <div v-if="event.event_tags && event.event_tags.length > 0" class="flex flex-wrap gap-1.5 mb-4">
        <span
          v-for="eventTag in event.event_tags.slice(0, 3)"
          :key="eventTag.tag.id"
          class="text-xs bg-gray-100 text-gray-700 px-2.5 py-1 rounded-full font-medium hover:bg-purple-100 hover:text-purple-700 transition-colors"
        >
          {{ eventTag.tag.name }}
        </span>
        <span v-if="event.event_tags.length > 3" class="text-xs text-purple-600 font-medium px-2 py-1">
          +{{ event.event_tags.length - 3 }}
        </span>
      </div>

      <!-- Footer Info -->
      <div class="pt-4 border-t border-gray-100 space-y-2">
        <!-- Location -->
        <div class="flex items-start gap-2 text-gray-600 text-sm">
          <svg class="w-5 h-5 flex-shrink-0 text-gray-400 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                  d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                  d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          <span class="line-clamp-1 font-medium">{{ event.venue }}</span>
        </div>

        <!-- Time -->
        <div class="flex items-center gap-2 text-gray-600 text-sm">
          <svg class="w-5 h-5 flex-shrink-0 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                  d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <span class="font-medium">{{ formatTime(event.event_date) }}</span>
        </div>
      </div>

      <!-- CTA Button appears on hover -->
      <div class="mt-4 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
        <div class="flex items-center justify-between text-purple-600 font-semibold">
          <span>View Details</span>
          <svg class="w-5 h-5 transform group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </div>
      </div>
    </div>
  </article>
</template>

<script setup lang="ts">
const props = defineProps<{
  event: any
}>()

const featuredImage = computed(() => {
  return props.event.event_images?.[0]?.url || null
})

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('en-ET').format(price)
}

const formatMonth = (date: string) => {
  return new Date(date).toLocaleDateString('en-US', { month: 'short' }).toUpperCase()
}

const formatDay = (date: string) => {
  return new Date(date).getDate()
}

const formatTime = (date: string) => {
  return new Date(date).toLocaleTimeString('en-US', { 
    hour: 'numeric', 
    minute: '2-digit',
    hour12: true 
  })
}
</script>

<style scoped>
/* Remove old styles as we're using Tailwind utilities */
</style>
