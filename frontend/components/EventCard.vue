<template>
  <div class="group bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1 cursor-pointer"
       @click="navigateTo(`/events/${event.id}`)">
    <!-- Image -->
    <div class="relative h-48 overflow-hidden">
      <LazyImage
        v-if="featuredImage"
        :src="featuredImage"
        :alt="event.title"
        className="w-full h-full object-cover transition-transform duration-300 group-hover:scale-110"
        :showPlaceholder="true"
      />
      <div v-else class="w-full h-full bg-gradient-to-br from-gray-200 to-gray-300 flex items-center justify-center">
        <svg class="w-16 h-16 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
        </svg>
      </div>
      
      <!-- Price badge with animation -->
      <div class="absolute bottom-2 left-2 bg-primary text-white px-3 py-1 rounded-full text-sm font-semibold shadow-lg animate-fade-in">
        {{ event.price === 0 ? 'FREE' : `${formatPrice(event.price)} ETB` }}
      </div>
    </div>

    <!-- Content -->
    <div class="p-4">
      <!-- Category & Date -->
      <div class="flex items-center gap-2 mb-2 flex-wrap">
        <span class="bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded font-medium">
          {{ event.category.name }}
        </span>
        <span class="text-gray-500 text-xs flex items-center gap-1">
          <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                  d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          {{ formatDate(event.event_date) }}
        </span>
      </div>

      <!-- Title -->
      <h3 class="text-lg font-semibold mb-2 line-clamp-2 hover:text-primary transition-colors">
        {{ event.title }}
      </h3>

      <!-- Description -->
      <p class="text-gray-600 text-sm mb-3 line-clamp-2">
        {{ event.description }}
      </p>

      <!-- Location & Stats -->
      <div class="flex items-center justify-between pt-3 border-t border-gray-100">
        <div class="flex items-center gap-1 text-gray-500 text-sm">
          <svg class="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                  d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                  d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          <span class="truncate">{{ event.venue }}</span>
        </div>
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
  })
}
</script>

<style scoped>
@keyframes fade-in {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in {
  animation: fade-in 0.3s ease-out;
}
</style>
