<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow">
      <div class="container mx-auto px-4 py-6">
        <div class="flex items-center justify-between">
          <NuxtLink to="/" class="text-2xl font-bold text-primary">
            EventHub Ethiopia
          </NuxtLink>
          <div class="flex items-center gap-4">
            <NuxtLink to="/dashboard" class="text-gray-600 hover:text-primary">
              My Events
            </NuxtLink>
            <NuxtLink to="/dashboard/follows" class="text-gray-600 hover:text-primary">
              Following
            </NuxtLink>
            <button @click="logout" class="text-gray-600 hover:text-primary">
              Logout
            </button>
          </div>
        </div>
      </div>
    </header>

    <main class="container mx-auto px-4 py-8">
      <h1 class="text-3xl font-bold mb-6">Bookmarked Events</h1>

      <!-- Loading -->
      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
      </div>

      <!-- Empty State -->
      <div v-else-if="bookmarkedEvents.length === 0" class="text-center py-12">
        <svg class="w-24 h-24 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
        </svg>
        <h3 class="text-xl font-semibold text-gray-700 mb-2">No bookmarks yet</h3>
        <p class="text-gray-500 mb-6">Start bookmarking events you're interested in</p>
        <NuxtLink to="/" class="text-primary hover:underline">
          Browse Events
        </NuxtLink>
      </div>

      <!-- Events Grid -->
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <EventCard v-for="bookmark in bookmarkedEvents" :key="bookmark.event.id" :event="bookmark.event" />
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { useQuery } from '@vue/apollo-composable'
import { GET_USER_BOOKMARKS } from '~/graphql/interactions'

definePageMeta({
  middleware: 'auth'
})

const { user, logout } = useAuth()

const { result, loading } = useQuery(GET_USER_BOOKMARKS, { user_id: user.value?.id })
const bookmarkedEvents = computed(() => result.value?.bookmarks || [])
</script>
