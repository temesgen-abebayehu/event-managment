<template>
  <div>
    <div class="container mx-auto px-4 py-8 max-w-3xl">
      <div v-if="loadingEvent" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
      </div>

      <div v-else-if="event">
        <div class="flex items-center justify-between mb-8">
          <h1 class="text-3xl font-bold">Edit Event</h1>
          <button @click="$router.back()" class="text-gray-600 hover:text-primary text-sm font-medium">← Back</button>
        </div>

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
            />
          </div>

          <!-- Current Images -->
          <div v-if="existingImages.length > 0" class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Current Images
            </label>
            <div class="grid grid-cols-3 gap-4">
              <div v-for="img in existingImages" :key="img.id" class="relative group">
                <img :src="img.url" class="w-full h-32 object-cover rounded-lg" />
                <span v-if="img.is_featured" class="absolute top-2 left-2 bg-purple-600 text-white text-xs px-2 py-1 rounded-full">
                  Featured
                </span>
                <!-- Delete Button -->
                <button
                  type="button"
                  @click="confirmDeleteImage(img)"
                  class="absolute top-2 right-2 bg-red-500 text-white p-1.5 rounded-full opacity-0 group-hover:opacity-100 transition-opacity"
                  title="Delete image"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
                <!-- Set Featured Button -->
                <button
                  v-if="!img.is_featured"
                  type="button"
                  @click="setFeaturedImage(img.id)"
                  class="absolute bottom-2 left-2 bg-purple-600 text-white text-xs px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity"
                >
                  Set Featured
                </button>
              </div>
            </div>
          </div>

          <!-- Upload New Images -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Add New Images (Optional)
            </label>
            <input
              type="file"
              accept="image/*"
              multiple
              @change="handleImageSelect"
              ref="fileInput"
              class="hidden"
            />
            <button
              type="button"
              @click="$refs.fileInput.click()"
              class="w-full px-4 py-3 border-2 border-dashed border-gray-300 rounded-lg hover:border-purple-500 transition-colors text-gray-600 hover:text-purple-600"
            >
              <svg class="w-8 h-8 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
              Click to upload images (Max 5)
            </button>

            <!-- New Images Preview -->
            <div v-if="newImagePreviews.length > 0" class="grid grid-cols-3 gap-4 mt-4">
              <div v-for="(preview, index) in newImagePreviews" :key="index" class="relative group">
                <img :src="preview.url" class="w-full h-32 object-cover rounded-lg" />
                <button
                  type="button"
                  @click="removeNewImage(index)"
                  class="absolute top-2 right-2 bg-red-500 text-white p-1.5 rounded-full opacity-0 group-hover:opacity-100 transition-opacity"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
                <span class="absolute bottom-2 left-2 bg-black bg-opacity-75 text-white text-xs px-2 py-1 rounded">
                  New
                </span>
              </div>
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
            />
          </div>

          <!-- Location -->
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
              />
              <input
                v-model.number="form.longitude"
                type="number"
                step="0.000001"
                required
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
              />
            </div>
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
              {{ loading ? 'Updating...' : 'Update Event' }}
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

    <!-- Delete Image Confirmation Modal -->
    <div
      v-if="showDeleteImageModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click="showDeleteImageModal = false"
    >
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4" @click.stop>
        <div class="flex items-center gap-3 mb-4">
          <div class="w-12 h-12 bg-red-100 rounded-full flex items-center justify-center flex-shrink-0">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
          </div>
          <div>
            <h3 class="text-lg font-bold text-gray-900">Delete Image?</h3>
            <p class="text-sm text-gray-500">This action cannot be undone</p>
          </div>
        </div>
        
        <p class="text-gray-600 mb-6">
          Are you sure you want to permanently delete this image? It will be removed from your event and from cloud storage.
        </p>
        
        <div class="flex gap-3">
          <button
            @click="handleDeleteImage"
            :disabled="deletingImage"
            class="flex-1 bg-red-600 text-white py-2.5 rounded-lg font-semibold hover:bg-red-700 transition disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ deletingImage ? 'Deleting...' : 'Delete Image' }}
          </button>
          <button
            @click="showDeleteImageModal = false"
            :disabled="deletingImage"
            class="flex-1 bg-gray-100 text-gray-700 py-2.5 rounded-lg font-semibold hover:bg-gray-200 transition disabled:opacity-50"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useQuery, useMutation } from '@vue/apollo-composable'
import { gql } from '@apollo/client/core'

definePageMeta({
  middleware: 'auth'
})

useHead({ title: 'Edit Event' })

const route = useRoute()
const router = useRouter()
const eventSlug = route.params.id
const { categories } = useCategories()
const config = useRuntimeConfig()

const GET_EVENT = gql`
  query GetEvent($slug: String!) {
    events(where: { slug: { _eq: $slug } }, limit: 1) {
      id
      slug
      title
      category_id
      description
      venue
      address
      latitude
      longitude
      event_date
      price
      event_images {
        id
        url
        public_id
        is_featured
      }
    }
  }
`

const UPDATE_EVENT = gql`
  mutation UpdateEvent($id: uuid!, $set: events_set_input!) {
    update_events_by_pk(pk_columns: { id: $id }, _set: $set) {
      id
    }
  }
`

const DELETE_IMAGE = gql`
  mutation DeleteImage($id: uuid!) {
    delete_event_images_by_pk(id: $id) {
      id
    }
  }
`

const INSERT_IMAGES = gql`
  mutation InsertImages($objects: [event_images_insert_input!]!) {
    insert_event_images(objects: $objects) {
      affected_rows
    }
  }
`

const SET_FEATURED_IMAGE = gql`
  mutation SetFeaturedImage($event_id: uuid!, $image_id: uuid!) {
    update_event_images(
      where: { event_id: { _eq: $event_id } }
      _set: { is_featured: false }
    ) {
      affected_rows
    }
    update_event_images_by_pk(
      pk_columns: { id: $image_id }
      _set: { is_featured: true }
    ) {
      id
    }
  }
`

const { result: eventResult, loading: loadingEvent, refetch } = useQuery(GET_EVENT, { slug: eventSlug })
const event = computed(() => eventResult.value?.events?.[0])
const eventId = computed(() => event.value?.id) // Keep ID for mutations
const existingImages = computed(() => event.value?.event_images || [])

const { mutate: updateEventMutation } = useMutation(UPDATE_EVENT)
const { mutate: deleteImageMutation } = useMutation(DELETE_IMAGE)
const { mutate: insertImagesMutation } = useMutation(INSERT_IMAGES)
const { mutate: setFeaturedMutation } = useMutation(SET_FEATURED_IMAGE)

const form = ref({
  title: '',
  category_id: '',
  description: '',
  venue: '',
  address: '',
  latitude: 9.005401,
  longitude: 38.763611,
  event_date: '',
  price: 0,
})

const loading = ref(false)
const error = ref('')
const newImageFiles = ref<File[]>([])
const newImagePreviews = ref<{url: string, file: File}[]>([])
const fileInput = ref(null)

// Delete image confirmation modal
const showDeleteImageModal = ref(false)
const imageToDelete = ref(null)
const deletingImage = ref(false)

// Load event data into form
watch(event, (newEvent) => {
  if (newEvent) {
    form.value = {
      title: newEvent.title,
      category_id: newEvent.category_id,
      description: newEvent.description,
      venue: newEvent.venue,
      address: newEvent.address,
      latitude: newEvent.latitude,
      longitude: newEvent.longitude,
      event_date: new Date(newEvent.event_date).toISOString().slice(0, 16),
      price: newEvent.price,
    }
  }
})

const handleImageSelect = (e: Event) => {
  const files = Array.from((e.target as HTMLInputElement).files || [])
  
  // Limit total images to 5
  const totalImages = existingImages.value.length + newImagePreviews.value.length + files.length
  if (totalImages > 5) {
    error.value = `You can only have up to 5 images total. Currently you have ${existingImages.value.length} existing and ${newImagePreviews.value.length} new images.`
    return
  }

  files.forEach(file => {
    const reader = new FileReader()
    reader.onload = (e) => {
      newImagePreviews.value.push({
        url: e.target?.result as string,
        file: file
      })
    }
    reader.readAsDataURL(file)
  })
}

const removeNewImage = (index: number) => {
  newImagePreviews.value.splice(index, 1)
}

const confirmDeleteImage = (img: any) => {
  imageToDelete.value = img
  showDeleteImageModal.value = true
}

const handleDeleteImage = async () => {
  if (!imageToDelete.value) return

  deletingImage.value = true
  try {
    // Delete from Cloudinary first
    await $fetch(`${config.public.backendUrl}/actions/delete-files`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${useCookie('auth_token').value}`,
      },
      body: {
        input: {
          public_ids: [imageToDelete.value.public_id]
        }
      }
    })

    // Delete from database
    await deleteImageMutation({ id: imageToDelete.value.id })
    
    showDeleteImageModal.value = false
    imageToDelete.value = null
    
    // Refetch event to update UI
    await refetch()
  } catch (err) {
    error.value = 'Failed to delete image'
  } finally {
    deletingImage.value = false
  }
}

const setFeaturedImage = async (imageId: string) => {
  try {
    await setFeaturedMutation({
      event_id: eventId.value,
      image_id: imageId
    })
    await refetch()
  } catch (err) {
    error.value = 'Failed to set featured image'
  }
}

const handleSubmit = async () => {
  loading.value = true
  error.value = ''

  try {
    // Update event details
    await updateEventMutation({
      id: eventId.value,
      set: form.value,
    })

    // Upload new images if any
    if (newImagePreviews.value.length > 0) {
      const formData = new FormData()
      formData.append('event_id', eventId.value) // Add event_id
      newImagePreviews.value.forEach(preview => {
        formData.append('files', preview.file)
      })

      const uploadResponse = await $fetch(`${config.public.backendUrl}/actions/upload`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${useCookie('auth_token').value}`,
        },
        body: formData,
      })

      // Insert image records - uploadResponse.files is the array
      const imageObjects = uploadResponse.files.map((img: any, index: number) => ({
        event_id: eventId.value,
        url: img.url,
        public_id: img.public_id,
        is_featured: existingImages.value.length === 0 && index === 0, // First image is featured if no existing images
      }))

      await insertImagesMutation({ objects: imageObjects })
    }

    router.push('/dashboard')
  } catch (err: any) {
    error.value = 'Failed to update event'
    console.error('Update error:', err)
  } finally {
    loading.value = false
  }
}

// Cleanup preview URLs on unmount
onUnmounted(() => {
  newImagePreviews.value.forEach(preview => {
    URL.revokeObjectURL(preview.url)
  })
})
</script>
