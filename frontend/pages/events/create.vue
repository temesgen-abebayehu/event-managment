<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow">
      <div class="container mx-auto px-4 py-6">
        <div class="flex items-center justify-between">
          <NuxtLink to="/dashboard" class="text-2xl font-bold text-primary">
            EventHub Ethiopia
          </NuxtLink>
          <button @click="$router.back()" class="text-gray-600 hover:text-primary">
            ← Back
          </button>
        </div>
      </div>
    </header>

    <main class="container mx-auto px-4 py-8 max-w-3xl">
      <h1 class="text-3xl font-bold mb-8">Create New Event</h1>

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
            placeholder="Enter event title"
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
            placeholder="Describe your event..."
          />
        </div>

        <!-- Images -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Event Images * (Max 5 images, 5MB each)
          </label>
          <input
            type="file"
            multiple
            accept="image/*"
            @change="handleFileSelect"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg"
          />
          <div v-if="selectedFiles.length > 0" class="mt-2 text-sm text-gray-600">
            {{ selectedFiles.length }} file(s) selected
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
            placeholder="e.g., Millennium Hall"
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
            placeholder="Full address with landmarks"
          />
        </div>

        <!-- Location (Map - simplified for now) -->
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
              placeholder="Latitude (e.g., 9.005401)"
            />
            <input
              v-model.number="form.longitude"
              type="number"
              step="0.000001"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
              placeholder="Longitude (e.g., 38.763611)"
            />
          </div>
          <p class="text-xs text-gray-500 mt-1">Default: Addis Ababa center</p>
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
            placeholder="0 for free events"
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
            {{ loading ? 'Creating...' : 'Create Event' }}
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
    </main>
  </div>
</template>

<script setup lang="ts">
import { useMutation } from '@vue/apollo-composable'
import { gql } from '@apollo/client/core'

definePageMeta({
  middleware: 'auth'
})

const config = useRuntimeConfig()
const { user } = useAuth()
const { categories } = useCategories()

const form = ref({
  title: '',
  category_id: '',
  description: '',
  venue: '',
  address: '',
  latitude: 9.005401, // Addis Ababa default
  longitude: 38.763611,
  event_date: '',
  price: 0,
})

const selectedFiles = ref<File[]>([])
const loading = ref(false)
const error = ref('')

const handleFileSelect = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files) {
    selectedFiles.value = Array.from(target.files).slice(0, 5)
  }
}

const CREATE_EVENT = gql`
  mutation CreateEvent($object: events_insert_input!) {
    insert_events_one(object: $object) {
      id
    }
  }
`

const { mutate: createEventMutation } = useMutation(CREATE_EVENT)

const handleSubmit = async () => {
  loading.value = true
  error.value = ''

  try {
    // 1. Upload images
    if (selectedFiles.value.length === 0) {
      error.value = 'Please select at least one image'
      loading.value = false
      return
    }

    const formData = new FormData()
    selectedFiles.value.forEach(file => {
      formData.append('files', file)
    })
    formData.append('user_id', user.value.id)
    formData.append('event_id', 'temp') // Will be replaced after event creation

    const uploadResponse = await $fetch(`${config.public.backendUrl}/actions/upload`, {
      method: 'POST',
      body: formData,
    })

    const uploadedFiles = uploadResponse.files

    // 2. Create event
    const eventData = {
      ...form.value,
      user_id: user.value.id,
      event_images: {
        data: uploadedFiles.map((file: any, index: number) => ({
          url: file.url,
          public_id: file.public_id,
          is_featured: index === 0, // First image is featured
        })),
      },
    }

    const result = await createEventMutation({ object: eventData })

    if (result?.data?.insert_events_one) {
      navigateTo('/dashboard')
    }
  } catch (err: any) {
    error.value = err.data?.error || 'Failed to create event'
  } finally {
    loading.value = false
  }
}
</script>
