<template>
  <div class="min-h-screen bg-white">
    <!-- Modern Hero Section with gradient and search -->
    <section class="relative bg-gradient-to-br from-purple-600 via-pink-600 to-orange-500 overflow-hidden">
      <!-- Background Image -->
      <div class="absolute inset-0">
        <div class="absolute inset-0 bg-gradient-to-r from-purple-900/70 to-pink-900/70"></div>
        <img src="public/consert2.jpg" alt="Concert background" class="w-full h-full object-cover z-10" />
      </div>
      
      <div class="relative container mx-auto px-4 py-20 md:py-28">
        <div class="max-w-4xl mx-auto text-center text-white">
          <h1 class="text-5xl md:text-6xl lg:text-7xl font-bold mb-6 leading-tight">
            Find your next
            <span class="block bg-gradient-to-r from-yellow-300 to-pink-300 bg-clip-text text-transparent">
              experience
            </span>
          </h1>
          <p class="text-xl md:text-2xl mb-10 text-white/90 font-light">
            Discover amazing events happening in Ethiopia
          </p>
          
          <!-- Enhanced Search Bar -->
          <div class="bg-white rounded-2xl shadow-2xl p-2 max-w-3xl mx-auto">
            <div class="flex flex-col md:flex-row gap-2">
              <div class="flex-1 relative">
                <svg class="absolute left-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
                <input
                  v-model="searchInput"
                  type="text"
                  placeholder="Search events by name, artist, or keyword..."
                  class="w-full pl-12 pr-4 py-4 rounded-xl border-2 border-transparent focus:border-purple-500 focus:outline-none text-gray-800 text-lg"
                  @keyup.enter="handleSearch(searchInput)"
                />
              </div>
              <button
                @click="handleSearch(searchInput)"
                class="bg-gradient-to-r from-purple-600 to-pink-600 text-white px-8 py-4 rounded-xl font-semibold hover:shadow-lg transform hover:scale-105 transition-all duration-200"
              >
                Search
              </button>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Wave decoration -->
      <div class="absolute bottom-0 left-0 right-0">
        <svg viewBox="0 0 1440 120" class="w-full h-12 md:h-20" preserveAspectRatio="none">
          <path fill="#ffffff" d="M0,64L48,69.3C96,75,192,85,288,80C384,75,480,53,576,48C672,43,768,53,864,58.7C960,64,1056,64,1152,58.7C1248,53,1344,43,1392,37.3L1440,32L1440,120L1392,120C1344,120,1248,120,1152,120C1056,120,960,120,864,120C768,120,672,120,576,120C480,120,384,120,288,120C192,120,96,120,48,120L0,120Z"></path>
        </svg>
      </div>
    </section>

    <!-- Main Content -->
    <div class="container mx-auto px-4 -mt-8 relative z-10">
      <!-- Quick Category Pills -->
      <div class="mb-8 flex gap-3 overflow-x-auto pb-4 scrollbar-hide">
        <button
          v-for="cat in quickCategories"
          :key="cat"
          @click="handleQuickFilter(cat)"
          class="flex-shrink-0 px-6 py-3 rounded-full border-2 border-gray-200 hover:border-purple-500 hover:bg-purple-50 transition-all duration-200 font-medium text-gray-700 hover:text-purple-700"
        >
          {{ cat }}
        </button>
      </div>

      <!-- Advanced Filters & View Controls -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-100 p-4 mb-8">
        <div class="flex flex-wrap items-center justify-between gap-4">
          <div class="flex items-center gap-4">
            <button
              @click="showFilters = !showFilters"
              class="flex items-center gap-2 px-4 py-2 rounded-lg border border-gray-300 hover:border-purple-500 transition-colors"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z" />
              </svg>
              <span>Filters</span>
            </button>
            
            <span class="text-gray-600 font-medium">
              <span class="text-purple-600 font-bold">{{ totalCount }}</span> events
            </span>
          </div>

          <div class="flex items-center gap-2">
            <span class="text-sm text-gray-600 mr-2">View:</span>
            <button
              @click="viewMode = 'grid'"
              :class="viewMode === 'grid' ? 'bg-purple-100 text-purple-700 border-purple-300' : 'bg-white text-gray-600 border-gray-200'"
              class="p-2 rounded-lg border-2 transition-all"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
              </svg>
            </button>
            <button
              @click="viewMode = 'map'"
              :class="viewMode === 'map' ? 'bg-purple-100 text-purple-700 border-purple-300' : 'bg-white text-gray-600 border-gray-200'"
              class="p-2 rounded-lg border-2 transition-all"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7" />
              </svg>
            </button>
          </div>
        </div>

        <!-- Collapsible Filters -->
        <transition name="expand">
          <div v-if="showFilters" class="mt-4 pt-4 border-t border-gray-100">
            <SearchFilters 
              @update:filters="handleFiltersUpdate" 
              @search="handleSearch"
            />
          </div>
        </transition>
      </div>

      <!-- Content Area -->
      <transition name="fade" mode="out-in">
        <!-- Loading State -->
        <div v-if="isLoading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          <EventCardSkeleton v-for="n in 8" :key="n" />
        </div>

        <!-- Empty State -->
        <div v-else-if="displayEvents.length === 0" class="text-center py-20">
          <div class="inline-block p-8 rounded-full bg-gray-100 mb-4">
            <svg class="w-16 h-16 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h3 class="text-2xl font-bold text-gray-800 mb-2">No events found</h3>
          <p class="text-gray-600 mb-6">Try adjusting your search or filters</p>
          <button
            @click="resetFilters"
            class="px-6 py-3 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors"
          >
            Clear all filters
          </button>
        </div>

        <!-- Grid View -->
        <div v-else-if="viewMode === 'grid'">
          <transition-group
            name="stagger"
            tag="div"
            class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6"
          >
            <EventCard v-for="event in displayEvents" :key="event.id" :event="event" />
          </transition-group>
        </div>

        <!-- Map View -->
        <div v-else-if="viewMode === 'map'" class="rounded-xl overflow-hidden shadow-lg border border-gray-200">
          <ClientOnly>
            <EventMap 
              :events="displayEvents" 
              @marker-click="handleMarkerClick"
            />
          </ClientOnly>
        </div>
      </transition>

      <!-- Modern Pagination -->
      <div v-if="totalCount > pageSize && !isLoading" class="mt-12 mb-8">
        <div class="flex items-center justify-center gap-2">
          <button
            @click="currentPage > 1 && (currentPage--)"
            :disabled="currentPage === 1"
            class="px-4 py-2 rounded-lg border border-gray-300 hover:border-purple-500 disabled:opacity-40 disabled:cursor-not-allowed transition-colors"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          
          <div class="flex gap-2">
            <button
              v-for="page in paginationPages"
              :key="page"
              @click="typeof page === 'number' && (currentPage = page)"
              :class="currentPage === page ? 'bg-purple-600 text-white' : 'bg-white text-gray-700 border border-gray-300 hover:border-purple-500'"
              class="min-w-[40px] px-3 py-2 rounded-lg font-medium transition-all"
              :disabled="page === '...'"
            >
              {{ page }}
            </button>
          </div>
          
          <button
            @click="currentPage < totalPages && (currentPage++)"
            :disabled="currentPage === totalPages"
            class="px-4 py-2 rounded-lg border border-gray-300 hover:border-purple-500 disabled:opacity-40 disabled:cursor-not-allowed transition-colors"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
        </div>
        
        <p class="text-center text-gray-600 mt-4">
          Showing <span class="font-semibold">{{ ((currentPage - 1) * pageSize) + 1 }}</span> to 
          <span class="font-semibold">{{ Math.min(currentPage * pageSize, totalCount) }}</span> of 
          <span class="font-semibold">{{ totalCount }}</span> events
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
useHead({ title: 'Discover Events' })

const viewMode = ref<'grid' | 'map'>('grid')
const currentPage = ref(1)
const pageSize = 12
const searchQuery = ref('')
const searchInput = ref('')
const activeFilters = ref({})
const showFilters = ref(false)

const quickCategories = ['All Events', 'Music', 'Arts', 'Sports', 'Food & Drink', 'Business']

const filters = computed(() => ({
  ...activeFilters.value,
  limit: pageSize,
  offset: (currentPage.value - 1) * pageSize,
}))

const { events: filteredEvents, totalCount, loading: eventsLoading } = useEvents(filters)
const { events: searchResults, loading: searchLoading } = useSearchEvents(searchQuery)

const displayEvents = computed(() => {
  if (searchQuery.value.length > 0) {
    return searchResults.value || []
  }
  return filteredEvents.value || []
})

const isLoading = computed(() => {
  if (searchQuery.value.length > 0) return searchLoading.value
  return eventsLoading.value
})

const totalPages = computed(() => Math.ceil(totalCount.value / pageSize))

// Smart pagination (show ... for large page counts)
const paginationPages = computed(() => {
  const pages: (number | string)[] = []
  const total = totalPages.value
  const current = currentPage.value

  if (total <= 7) {
    for (let i = 1; i <= total; i++) pages.push(i)
  } else {
    if (current <= 3) {
      for (let i = 1; i <= 5; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    } else if (current >= total - 2) {
      pages.push(1)
      pages.push('...')
      for (let i = total - 4; i <= total; i++) pages.push(i)
    } else {
      pages.push(1)
      pages.push('...')
      for (let i = current - 1; i <= current + 1; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    }
  }
  return pages
})

const handleFiltersUpdate = (newFilters: any) => {
  activeFilters.value = newFilters
  currentPage.value = 1
}

const handleSearch = (query: string) => {
  searchQuery.value = query
  currentPage.value = 1
}

const handleQuickFilter = (category: string) => {
  if (category === 'All Events') {
    activeFilters.value = {}
  } else {
    activeFilters.value = { category }
  }
  currentPage.value = 1
}

const resetFilters = () => {
  activeFilters.value = {}
  searchQuery.value = ''
  searchInput.value = ''
  currentPage.value = 1
}

const handleMarkerClick = (event: any) => {
  navigateTo(`/events/${event.id}`)
}
</script>

<style scoped>
.scrollbar-hide::-webkit-scrollbar {
  display: none;
}

.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.stagger-move {
  transition: transform 0.3s ease;
}

.stagger-enter-active {
  transition: all 0.4s ease;
}

.stagger-enter-from {
  opacity: 0;
  transform: translateY(30px) scale(0.95);
}

.expand-enter-active,
.expand-leave-active {
  transition: all 0.3s ease;
  max-height: 500px;
  overflow: hidden;
}

.expand-enter-from,
.expand-leave-to {
  max-height: 0;
  opacity: 0;
}
</style>
