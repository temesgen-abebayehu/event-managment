<template>
  <div>
    <div class="container mx-auto px-4 py-8 max-w-3xl">
      <div class="flex items-center justify-between mb-8">
        <h1 class="text-3xl font-bold">Create New Event</h1>
        <button
          @click="$router.back()"
          class="flex items-center gap-2 text-gray-600 hover:text-purple-600 transition-colors group"
        >
          <svg class="w-5 h-5 transform group-hover:-translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
          <span class="font-medium">Back</span>
        </button>
      </div>

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
            Description (Optional)
          </label>
          <textarea
            v-model="description"
            rows="4"
            class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent border-gray-300"
            placeholder="Add more details about your event..."
          />
        </div>

        <!-- Images -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Event Images (Optional - Max 5 images, 5MB each)
          </label>
          <input
            type="file"
            multiple
            accept="image/*"
            @change="handleFileSelect"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg"
          />
          <p v-if="imageError" class="mt-2 text-sm text-red-600">{{ imageError }}</p>
          
          <!-- Image Previews -->
          <div v-if="imagePreviews.length > 0" class="mt-4 grid grid-cols-2 sm:grid-cols-3 md:grid-cols-5 gap-4">
            <div
              v-for="(preview, index) in imagePreviews"
              :key="index"
              class="relative group"
            >
              <img
                :src="preview.url"
                :alt="`Preview ${index + 1}`"
                class="w-full h-32 object-cover rounded-lg border-2"
                :class="preview.isFeatured ? 'border-purple-500' : 'border-gray-200'"
              />
              
              <!-- Featured Badge -->
              <div
                v-if="preview.isFeatured"
                class="absolute top-2 left-2 bg-purple-600 text-white text-xs px-2 py-1 rounded-full font-semibold"
              >
                Featured
              </div>
              
              <!-- Action Buttons -->
              <div class="absolute inset-0 bg-black bg-opacity-50 opacity-0 group-hover:opacity-100 transition-opacity rounded-lg flex items-center justify-center gap-2">
                <button
                  v-if="!preview.isFeatured"
                  @click="setFeaturedImage(index)"
                  type="button"
                  class="bg-white text-gray-800 px-3 py-1 rounded text-xs font-semibold hover:bg-purple-100"
                >
                  Set Featured
                </button>
                <button
                  @click="removeImage(index)"
                  type="button"
                  class="bg-red-500 text-white px-3 py-1 rounded text-xs font-semibold hover:bg-red-600"
                >
                  Remove
                </button>
              </div>
            </div>
          </div>
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
            Address *
          </label>
          <input
            v-model="address"
            type="text"
            class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
            :class="errors.address ? 'border-red-500' : 'border-gray-300'"
            placeholder="Event address"
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

        <!-- Ticket Quantity -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Total Tickets Available *
          </label>
          <input
            v-model.number="ticketQuantity"
            type="number"
            min="1"
            class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
            :class="errors.ticketQuantity ? 'border-red-500' : 'border-gray-300'"
            placeholder="e.g., 100"
          />
          <p v-if="errors.ticketQuantity" class="mt-1 text-sm text-red-600">{{ errors.ticketQuantity }}</p>
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
    </div>
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

useHead({ title: 'Create Event' })

const config = useRuntimeConfig()
const { user } = useAuth()
const { categories } = useCategories()

// Validation schema
const schema = yup.object({
  title: yup.string().required('Event title is required').min(3, 'Title must be at least 3 characters'),
  categoryId: yup.string().required('Please select a category'),
  description: yup.string(), // Optional - no validation
  venue: yup.string().required('Venue name is required'),
  address: yup.string().required('Address is required'),
  eventDate: yup.string().required('Event date is required'),
  price: yup.number().required('Price is required').min(0, 'Price cannot be negative'),
  ticketQuantity: yup.number().required('Total tickets is required').min(1, 'Must have at least 1 ticket'),
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
const { value: ticketQuantity } = useField<number>('ticketQuantity')

const selectedFiles = ref<File[]>([])
const imagePreviews = ref<{ file: File; url: string; isFeatured: boolean }[]>([])
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
    
    // Create previews
    imagePreviews.value = files.map((file, index) => ({
      file,
      url: URL.createObjectURL(file),
      isFeatured: index === 0 // First image is featured by default
    }))
  }
}

const setFeaturedImage = (index: number) => {
  imagePreviews.value.forEach((preview, i) => {
    preview.isFeatured = i === index
  })
}

const removeImage = (index: number) => {
  // Revoke object URL to free memory
  URL.revokeObjectURL(imagePreviews.value[index].url)
  
  imagePreviews.value.splice(index, 1)
  selectedFiles.value.splice(index, 1)
  
  // If removed image was featured, make first image featured
  if (imagePreviews.value.length > 0 && !imagePreviews.value.some(p => p.isFeatured)) {
    imagePreviews.value[0].isFeatured = true
  }
}

// Cleanup object URLs when component unmounts
onUnmounted(() => {
  imagePreviews.value.forEach(preview => {
    URL.revokeObjectURL(preview.url)
  })
})

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

const CREATE_TICKET = gql`
  mutation CreateTicket($event_id: uuid!, $price: numeric!, $quantity_total: Int!) {
    insert_tickets_one(object: {
      event_id: $event_id
      price: $price
      quantity_total: $quantity_total
    }) {
      id
    }
  }
`

const { mutate: createEventMutation } = useMutation(CREATE_EVENT)
const { mutate: createTicketMutation } = useMutation(CREATE_TICKET)

const onSubmit = handleSubmit(async (values) => {
  loading.value = true
  serverError.value = ''

  let eventData: any = {}

  try {
    let uploadedFiles = []

    // 1. Upload images (optional)
    if (selectedFiles.value.length > 0) {
      const formData = new FormData()
      selectedFiles.value.forEach(file => {
        formData.append('files', file)
      })
      formData.append('event_id', 'temp')

      const uploadResponse = await $fetch(`${config.public.backendUrl}/actions/upload`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${useCookie('auth_token').value}`,
        },
        body: formData,
      })

      uploadedFiles = uploadResponse.files
    }

    // 2. Create event
    eventData = {
      title: values.title,
      description: values.description || '',
      venue: values.venue,
      address: values.address,
      latitude: latitude.value,
      longitude: longitude.value,
      price: values.price,
      event_date: values.eventDate,
      category_id: values.categoryId,
    }

    // Add images only if uploaded
    if (uploadedFiles.length > 0) {
      // Find which preview was marked as featured
      const featuredIndex = imagePreviews.value.findIndex(p => p.isFeatured)
      
      eventData.event_images = {
        data: uploadedFiles.map((file: any, index: number) => ({
          url: file.url,
          public_id: file.public_id,
          is_featured: index === featuredIndex,
        })),
      }
    }

    if (tags.value.length > 0) {
      eventData.event_tags = {
        data: tags.value.map(tagName => ({
          tag: {
            data: { name: tagName },
            on_conflict: {
              constraint: 'tags_name_key',
              update_columns: ['name'],
            },
          },
        })),
      }
    }

    console.log('Event data being sent:', JSON.stringify(eventData, null, 2))
    
    const result = await createEventMutation({ object: eventData })

    if (result?.data?.insert_events_one) {
      const eventId = result.data.insert_events_one.id
      
      // 3. Create ticket for the event
      await createTicketMutation({
        event_id: eventId,
        price: values.price,
        quantity_total: values.ticketQuantity
      })
      
      navigateTo('/dashboard')
    }
  } catch (err: any) {
    console.error('Create event error:', err)
    console.error('Event data:', eventData)
    serverError.value = err.message || err.data?.error || 'Failed to create event'
  } finally {
    loading.value = false
  }
})
</script>
