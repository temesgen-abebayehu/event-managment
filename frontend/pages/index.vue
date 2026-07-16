<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow">
      <div class="container mx-auto px-4 py-6">
        <div class="flex items-center justify-between">
          <h1 class="text-3xl font-bold text-primary">EventHub Ethiopia</h1>
          <nav class="flex gap-4">
            <NuxtLink to="/auth/login" class="text-gray-600 hover:text-primary">
              Login
            </NuxtLink>
            <NuxtLink to="/auth/signup" class="bg-primary text-white px-4 py-2 rounded-lg hover:bg-primary-dark">
              Sign Up
            </NuxtLink>
          </nav>
        </div>
      </div>
    </header>

    <!-- Hero Section -->
    <section class="bg-gradient-to-r from-primary to-primary-dark text-white py-16">
      <div class="container mx-auto px-4 text-center">
        <h2 class="text-4xl md:text-5xl font-bold mb-4">
          Discover Amazing Events in Ethiopia
        </h2>
        <p class="text-xl mb-8">
          Find concerts, festivals, workshops, and more happening near you
        </p>
      </div>
    </section>

    <!-- Main Content -->
    <main class="container mx-auto px-4 py-8">
      <!-- Search & Filters -->
      <SearchFilters 
        @update:filters="handleFiltersUpdate" 
        @search="handleSearch"
      />

      <!-- View Toggle -->
      <div class="flex justify-between items-center mb-6">
        <p class="text-gray-600">
          {{ totalCount }} events found
        </p>
        <div class="flex gap-2">
          <button
            @click="viewMode = 'grid'"
            :class="viewMode === 'grid' ? 'bg-primary text-white' : 'bg-gray-200 text-gray-700'"
            class="px-4 py-2 rounded-lg"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                    d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
            </svg>
          </button>
          <button
            @click="viewMode = 'map'"
            :class="viewMode === 'map' ? 'bg-primary text-white' : 'bg-gray-200 text-gray-700'"
            class="px-4 py-2 rounded-lg"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                    d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7" />
            </svg>
          </button>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
        <p class="mt-4 text-gray-600">Loading events...</p>
      </div>

      <!-- Grid View -->
      <div v-else-if="viewMode === 'grid'">
        <div v-if="displayEvents.length === 0" class="text-center py-12">
          <p class="text-gray-600">No events found matching your criteria</p>
        </div>
        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <EventCard v-for="event in displayEvents" :key="event.id" :event="event" />
        </div>
      </div>

      <!-- Map View -->
      <div v-else-if="viewMode === 'map'">
        <ClientOnly>
          <EventMap 
            :events="displayEvents" 
            @marker-click="handleMarkerClick"
          />
        </ClientOnly>
      </div>

      <!-- Pagination -->
      <div v-if="totalCount > pageSize" class="mt-8 flex justify-center gap-2">
        <button
          v-for="page in totalPages"
          :key="page"
          @click="currentPage = page"
          :class="currentPage === page ? 'bg-primary text-white' : 'bg-gray-200 text-gray-700'"
          class="px-4 py-2 rounded-lg hover:bg-primary hover:text-white transition"
        >
          {{ page }}
        </button>
      </div>
    </main>

    <!-- Footer -->
    <footer class="bg-gray-800 text-white py-8 mt-16">
      <div class="container mx-auto px-4 text-center">
        <p>&copy; 2024 EventHub Ethiopia. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
const viewMode = ref<'grid' | 'map'>('grid')
const currentPage = ref(1)
const pageSize = 12
const searchQuery = ref('')
const activeFilters = ref({})

const filters = computed(() => ({
  ...activeFilters.value,
  limit: pageSize,
  offset: (currentPage.value - 1) * pageSize,
}))

// Use events or search based on search query
const { events, totalCount, loading } = searchQuery.value 
  ? useSearchEvents(searchQuery)
  : useEvents(filters)

const displayEvents = computed(() => events.value || [])
const totalPages = computed(() => Math.ceil(totalCount.value / pageSize))

const handleFiltersUpdate = (newFilters: any) => {
  activeFilters.value = newFilters
  currentPage.value = 1
}

const handleSearch = (query: string) => {
  searchQuery.value = query
  currentPage.value = 1
}

const handleMarkerClick = (event: any) => {
  navigateTo(`/events/${event.id}`)
}
</script>
