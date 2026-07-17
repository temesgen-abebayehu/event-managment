<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow">
      <div class="container mx-auto px-4 py-6">
        <div class="flex items-center justify-between">
          <NuxtLink to="/" class="text-2xl font-bold text-primary">
            EventHub Ethiopia - Admin
          </NuxtLink>
          <div class="flex items-center gap-4">
            <NuxtLink to="/dashboard" class="text-gray-600 hover:text-primary">
              Dashboard
            </NuxtLink>
            <button @click="logout" class="text-gray-600 hover:text-primary">
              Logout
            </button>
          </div>
        </div>
      </div>
    </header>

    <main class="container mx-auto px-4 py-8">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold">Category Management</h1>
        <button
          @click="openCreateModal"
          class="bg-primary text-white px-6 py-3 rounded-lg font-semibold hover:bg-primary-dark transition"
        >
          + Add Category
        </button>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
      </div>

      <!-- Categories List -->
      <div v-else class="bg-white rounded-lg shadow overflow-hidden">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Name
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Slug
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Events
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="category in categories" :key="category.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ category.name }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-500">{{ category.slug }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-500">{{ category.events_aggregate?.aggregate?.count || 0 }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <button
                  @click="openEditModal(category)"
                  class="text-blue-600 hover:text-blue-900 mr-4"
                >
                  Edit
                </button>
                <button
                  @click="confirmDelete(category)"
                  class="text-red-600 hover:text-red-900"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </main>

    <!-- Create/Edit Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click="showModal = false"
    >
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4" @click.stop>
        <h3 class="text-xl font-bold mb-4">
          {{ editingCategory ? 'Edit Category' : 'Add Category' }}
        </h3>

        <form @submit.prevent="handleSubmit">
          <!-- Name -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Category Name *
            </label>
            <input
              v-model="form.name"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
              placeholder="e.g., Music"
            />
          </div>

          <!-- Slug -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Slug *
            </label>
            <input
              v-model="form.slug"
              type="text"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
              placeholder="e.g., music"
            />
            <p class="text-xs text-gray-500 mt-1">URL-friendly identifier (lowercase, no spaces)</p>
          </div>

          <!-- Error Message -->
          <div v-if="error" class="mb-4 p-3 bg-red-50 text-red-700 rounded-lg text-sm">
            {{ error }}
          </div>

          <!-- Buttons -->
          <div class="flex gap-4">
            <button
              type="submit"
              :disabled="submitting"
              class="flex-1 bg-primary text-white py-2 rounded-lg hover:bg-primary-dark transition disabled:opacity-50"
            >
              {{ submitting ? 'Saving...' : 'Save' }}
            </button>
            <button
              type="button"
              @click="showModal = false"
              class="flex-1 bg-gray-200 text-gray-700 py-2 rounded-lg hover:bg-gray-300 transition"
            >
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="showDeleteModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click="showDeleteModal = false"
    >
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4" @click.stop>
        <h3 class="text-xl font-bold mb-4">Delete Category?</h3>
        <p class="text-gray-600 mb-6">
          Are you sure you want to delete "{{ categoryToDelete?.name }}"? 
          This will affect {{ categoryToDelete?.events_aggregate?.aggregate?.count || 0 }} events.
        </p>
        <div class="flex gap-4">
          <button
            @click="handleDelete"
            :disabled="deleting"
            class="flex-1 bg-red-500 text-white py-2 rounded-lg hover:bg-red-600 transition disabled:opacity-50"
          >
            {{ deleting ? 'Deleting...' : 'Delete' }}
          </button>
          <button
            @click="showDeleteModal = false"
            class="flex-1 bg-gray-200 text-gray-700 py-2 rounded-lg hover:bg-gray-300 transition"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useMutation, useQuery } from '@vue/apollo-composable'
import { gql } from '@apollo/client/core'

definePageMeta({
  middleware: 'auth'
})

const { logout } = useAuth()

const GET_CATEGORIES = gql`
  query GetAllCategories {
    categories(order_by: { name: asc }) {
      id
      name
      slug
      events_aggregate {
        aggregate {
          count
        }
      }
    }
  }
`

const CREATE_CATEGORY = gql`
  mutation CreateCategory($name: String!, $slug: String!) {
    insert_categories_one(object: { name: $name, slug: $slug }) {
      id
    }
  }
`

const UPDATE_CATEGORY = gql`
  mutation UpdateCategory($id: uuid!, $name: String!, $slug: String!) {
    update_categories_by_pk(pk_columns: { id: $id }, _set: { name: $name, slug: $slug }) {
      id
    }
  }
`

const DELETE_CATEGORY = gql`
  mutation DeleteCategory($id: uuid!) {
    delete_categories_by_pk(id: $id) {
      id
    }
  }
`

const { result, loading, refetch } = useQuery(GET_CATEGORIES)
const categories = computed(() => result.value?.categories || [])

const { mutate: createCategory } = useMutation(CREATE_CATEGORY)
const { mutate: updateCategory } = useMutation(UPDATE_CATEGORY)
const { mutate: deleteCategory } = useMutation(DELETE_CATEGORY)

const showModal = ref(false)
const showDeleteModal = ref(false)
const editingCategory = ref(null)
const categoryToDelete = ref(null)
const submitting = ref(false)
const deleting = ref(false)
const error = ref('')

const form = ref({
  name: '',
  slug: '',
})

const openCreateModal = () => {
  editingCategory.value = null
  form.value = { name: '', slug: '' }
  error.value = ''
  showModal.value = true
}

const openEditModal = (category: any) => {
  editingCategory.value = category
  form.value = { name: category.name, slug: category.slug }
  error.value = ''
  showModal.value = true
}

const handleSubmit = async () => {
  submitting.value = true
  error.value = ''

  try {
    if (editingCategory.value) {
      await updateCategory({
        id: editingCategory.value.id,
        name: form.value.name,
        slug: form.value.slug,
      })
    } else {
      await createCategory({
        name: form.value.name,
        slug: form.value.slug,
      })
    }
    showModal.value = false
    refetch()
  } catch (err: any) {
    error.value = err.message || 'Failed to save category'
  } finally {
    submitting.value = false
  }
}

const confirmDelete = (category: any) => {
  categoryToDelete.value = category
  showDeleteModal.value = true
}

const handleDelete = async () => {
  if (!categoryToDelete.value) return

  deleting.value = true
  try {
    await deleteCategory({ id: categoryToDelete.value.id })
    showDeleteModal.value = false
    categoryToDelete.value = null
    refetch()
  } catch (err) {
    console.error('Failed to delete category:', err)
  } finally {
    deleting.value = false
  }
}

// Auto-generate slug from name
watch(() => form.value.name, (newName) => {
  if (!editingCategory.value) {
    form.value.slug = newName.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/^-|-$/g, '')
  }
})
</script>
