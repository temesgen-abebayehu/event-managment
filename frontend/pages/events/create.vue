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

      <form @submit="onSubmit" class="bg-white rounded-lg shadow-lg p-8">
        <!-- Title -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Event Title *
          </label>
          <input
            v-model="title"
            type="text"
            class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
            :class="errors.title ? 'border-red-500' : 'border-gray-300'"
            placeholder="Enter event title"
          />
          <p v-if="errors.title" class="mt-1 text-sm text-red-600">{{ errors.title }}</p>
        </div>

        <!-- Category -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Category *
          </label>
          <select
            v-model="categoryId"
            class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
            :class="errors.categoryId ? 'border-red-500' : 'border-gray-300'"
          >
            <option value="">Select category</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.id">
              {{ cat.name }}
            </option>
          </select>
          <p v-if="errors.categoryId" class="mt-1 text-sm text-red-600">{{ errors.categoryId }}</p>
        </div>

        <!-- Description -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Description *
          </label>
          <textarea
            v-model="description"
            rows="6"
            class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
            :class="errors.description ? 'border-red-500' : 'border-gray-300'"
            placeholder="Describe your event..."
          />
          <p v-if="errors.description" class="mt-1 text-sm text-red-600">{{ errors.description }}</p>
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
          <p v-if="imageError" class="mt-1 text-sm text-red-600">{{ imageError }}</p>
        </div>

        <!-- Venue -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Venue Name *
          </label>
          <input
            v-model="venue"
            type="text"
            class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
            :class="errors.venue ? 'border-red-500' : 'border-gray-300'"
            placeholder="e.g., Millennium Hall"
          />
          <p v-if="errors.venue" class="mt-1 text-sm text-red-600">{{ errors.venue }}</p>
        </div>

        <!-- Address -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Full Address *
          </label>
          <input
            v-model="address"
            type="text"
            class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
            :class="errors.address ? 'border-red-500' : 'border-gray-300'"
            placeholder="Full address with landmarks"
          />
          <p v-if="errors.address" class="mt-1 text-sm text-red-600">{{ errors.address }}</p>
        </div>

        <!-- Location Map Picker -->
        <div class="mb-6">
          <ClientOnly>
            <LocationPicker
              :initial-lat="latitude"
              :initial-lng="longitude"
              @update="handleLocationUpdate"
            />
          </ClientOnly>
        </div>

        <!-- Date & Time -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Event Date & Time *
          </label>
          <input
            v-model="eventDate"
            type="datetime-local"
            class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
            :class="errors.eventDate ? 'border-red-500' : 'border-gray-300'"
          />
          <p v-if="errors.eventDate" class="mt-1 text-sm text-red-600">{{ errors.eventDate }}</p>
        </div>

        <!-- Price -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Ticket Price (ETB) *
          </label>
          <input
            v-model.number="price"
            type="number"
            step="0.01"
            min="0"
            class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
            :class="errors.price ? 'border-red-500' : 'border-gray-300'"
            placeholder="0 for free events"
          />
          <p v-if="errors.price" class="mt-1 text-sm text-red-600">{{ errors.price }}</p>
        </div>

        <!-- Tags -->
        <div class="mb-6">
          <TagInput v-model="tags" />
        </div>

        <!-- Error Message -->
        <div v-if="serverError" class="mb-6 p-3 bg-red-50 text-red-700 rounded-lg text-sm">
          {{ serverError }}
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
import { useForm, useField } from 'vee-validate'
import * as yup from 'yup'

definePageMeta({
  middleware: 'auth'
})

const config = useRuntimeConfig()
const { user } = useAuth()
const { categories } = useCategories()

// Validation schema
const schema = yup.object({
  title: yup.string().required('Event title is required').min(5, 'Title must be at least 5 characters'),
  categoryId: yup.string().required('Please select a category'),
  description: yup.string().required('Description is required').min(20, 'Description must be at least 20 characters'),
  venue: yup.string().required('Venue name is required'),
  address: yup.string().required('Full address is required'),
  eventDate: yup.string().required('Event date is required'),
  price: yup.number().required('Price is required').min(0, 'Price cannot be negative'),
})

const { handleSubmit, errors } = useForm({
  validationSchema: schema,
})

const { value: title } = useField<string>('title')
const { value: categoryId } = useField<string>('categoryId')
const { value: description } = useField<string>('description')
const { value: venue } = useField<string>('venue')
const { value: address } = useField<string>('address')
const { value: eventDate } = useField<string>('eventDate')
const { value: price } = useField<number>('price')

const selectedFiles = ref<File[]>([])
const latitude = ref(9.005401) // Addis Ababa default
const longitude = ref(38.763611)
const tags = ref<string[]>([])
const loading = ref(false)
const serverError = ref('')
const imageError = ref('')

const handleFileSelect = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files) {
    const files = Array.from(target.files).slice(0, 5)
    
    // Validate file sizes
    const invalidFiles = files.filter(f => f.size > 5 * 1024 * 1024)
    if (invalidFiles.length > 0) {
      imageError.value = 'Some files exceed 5MB limit'
      return
    }
    
    selectedFiles.value = files
    imageError.value = ''
  }
}

const handleLocationUpdate = ({ lat, lng }: { lat: number; lng: number }) => {
  latitude.value = lat
  longitude.value = lng
}

const CREATE_EVENT = gql`
  mutation CreateEvent($object: events_insert_input!) {
    insert_events_one(object: $object) {
      id
    }
  }
`

const { mutate: createEventMutation } = useMutation(CREATE_EVENT)

const onSubmit = handleSubmit(async (values) => {
  // Validate images
  if (selectedFiles.value.length === 0) {
    imageError.value = 'Please select at least one image'
    return
  }

  loading.value = true
  serverError.value = ''

  try {
    // 1. Upload images
    const formData = new FormData()
    selectedFiles.value.forEach(file => {
      formData.append('files', file)
    })
    formData.append('user_id', user.value.id)
    formData.append('event_id', 'temp')

    const uploadResponse = await $fetch(`${config.public.backendUrl}/actions/upload`, {
      method: 'POST',
      body: formData,
    })

    const uploadedFiles = uploadResponse.files

    // 2. Prepare tags data
    const tagsData = tags.value.map(tagName => ({
      tag: {
        data: { name: tagName },
        on_conflict: {
          constraint: 'tags_name_key',
          update_columns: ['name']
        }
      }
    }))

    // 3. Create event
    const eventData = {
      title: values.title,
      description: values.description,
      venue: values.venue,
      address: values.address,
      latitude: latitude.value,
      longitude: longitude.value,
      price: values.price,
      event_date: values.eventDate,
      category_id: values.categoryId,
      user_id: user.value.id,
      event_images: {
        data: uploadedFiles.map((file: any, index: number) => ({
          url: file.url,
          public_id: file.public_id,
          is_featured: index === 0,
        })),
      },
      event_tags: {
        data: tagsData,
      },
    }

    const result = await createEventMutation({ object: eventData })

    if (result?.data?.insert_events_one) {
      navigateTo('/dashboard')
    }
  } catch (err: any) {
    serverError.value = err.data?.error || 'Failed to create event'
  } finally {
    loading.value = false
  }
})
</script>
