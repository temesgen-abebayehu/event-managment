<template>
  <div class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-all duration-300 cursor-pointer"
       @click="navigateTo(`/events/${event.id}`)">
    <!-- Image -->
    <div class="relative h-48">
      <img 
        v-if="featuredImage" 
        :src="featuredImage" 
        :alt="event.title" 
        class="w-full h-full object-cover"
      />
      <div v-else class="w-full h-full bg-gray-200 flex items-center justify-center">
        <span class="text-gray-400">No image</span>
      </div>
      
      <!-- Price badge -->
      <div class="absolute bottom-2 left-2 bg-primary text-white px-3 py-1 rounded-full text-sm font-semibold">
        {{ event.price === 0 ? 'FREE' : `${formatPrice(event.price)} ETB` }}
      </div>
    </div>

    <!-- Content -->
    <div class="p-4">
      <!-- Category & Date -->
      <div class="flex items-center gap-2 mb-2">
        <span class="bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded">
          {{ event.category.name }}
        </span>
        <span class="text-gray-500 text-sm">
          {{ formatDate(event.event_date) }}
        </span>
      </div>

      <!-- Title -->
      <h3 class="text-lg font-semibold mb-2 line-clamp-2 hover:text-primary">
        {{ event.title }}
      </h3>

      <!-- Description -->
      <p class="text-gray-600 text-sm mb-3 line-clamp-2">
        {{ event.description }}
      </p>

      <!-- Location -->
      <div class="flex items-center text-gray-500 text-sm">
        <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
        </svg>
        <span class="line-clamp-1">{{ event.venue }}</span>
      </div>
    </div>
  </div>
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

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('en-ET', {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}
</script>
