<template>
  <div class="bg-white rounded-lg shadow p-6 mb-6">
    <!-- Search -->
    <div class="mb-4">
      <input
        v-model="searchQuery"
        type="text"
        placeholder="Search events by title, description..."
        class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
        @input="onSearchChange"
      />
    </div>

    <!-- Filters Row -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <!-- Category Filter -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Category</label>
        <select
          v-model="filters.category_id"
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
          @change="onFilterChange"
        >
          <option value="">All Categories</option>
          <option v-for="category in categories" :key="category.id" :value="category.id">
            {{ category.name }}
          </option>
        </select>
      </div>

      <!-- Date Filter -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Date From</label>
        <input
          v-model="filters.date_from"
          type="date"
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
          @change="onFilterChange"
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Date To</label>
        <input
          v-model="filters.date_to"
          type="date"
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
          @change="onFilterChange"
        />
      </div>

      <!-- Price Filter -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Max Price (ETB)</label>
        <input
          v-model.number="filters.max_price"
          type="number"
          min="0"
          placeholder="Any price"
          class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
          @input="onFilterChange"
        />
      </div>
    </div>

    <!-- Quick Filters -->
    <div class="mt-4 flex flex-wrap gap-2">
      <button
        @click="setQuickFilter('free')"
        :class="quickFilter === 'free' ? 'bg-primary text-white' : 'bg-gray-100 text-gray-700'"
        class="px-4 py-2 rounded-full text-sm font-medium hover:bg-primary hover:text-white transition"
      >
        Free Events
      </button>
      <button
        @click="setQuickFilter('today')"
        :class="quickFilter === 'today' ? 'bg-primary text-white' : 'bg-gray-100 text-gray-700'"
        class="px-4 py-2 rounded-full text-sm font-medium hover:bg-primary hover:text-white transition"
      >
        Today
      </button>
      <button
        @click="setQuickFilter('this_week')"
        :class="quickFilter === 'this_week' ? 'bg-primary text-white' : 'bg-gray-100 text-gray-700'"
        class="px-4 py-2 rounded-full text-sm font-medium hover:bg-primary hover:text-white transition"
      >
        This Week
      </button>
      <button
        @click="clearFilters"
        class="px-4 py-2 rounded-full text-sm font-medium bg-gray-200 text-gray-700 hover:bg-gray-300 transition"
      >
        Clear All
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
const emit = defineEmits(['update:filters', 'search'])

const { categories } = useCategories()

const searchQuery = ref('')
const quickFilter = ref('')
const filters = ref({
  category_id: '',
  date_from: '',
  date_to: '',
  max_price: null,
  min_price: 0,
})

let searchTimeout: any = null

const onSearchChange = () => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    emit('search', searchQuery.value)
  }, 500)
}

const onFilterChange = () => {
  quickFilter.value = ''
  emit('update:filters', { ...filters.value })
}

const setQuickFilter = (type: string) => {
  quickFilter.value = type
  const now = new Date()
  
  if (type === 'free') {
    filters.value.max_price = 0
    filters.value.min_price = 0
  } else if (type === 'today') {
    const today = now.toISOString().split('T')[0]
    filters.value.date_from = today
    filters.value.date_to = today
  } else if (type === 'this_week') {
    const weekEnd = new Date(now)
    weekEnd.setDate(now.getDate() + 7)
    filters.value.date_from = now.toISOString().split('T')[0]
    filters.value.date_to = weekEnd.toISOString().split('T')[0]
  }
  
  onFilterChange()
}

const clearFilters = () => {
  searchQuery.value = ''
  quickFilter.value = ''
  filters.value = {
    category_id: '',
    date_from: '',
    date_to: '',
    max_price: null,
    min_price: 0,
  }
  emit('search', '')
  emit('update:filters', {})
}
</script>
