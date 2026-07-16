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
            <button @click="logout" class="text-gray-600 hover:text-primary">
              Logout
            </button>
          </div>
        </div>
      </div>
    </header>

    <main class="container mx-auto px-4 py-8">
      <h1 class="text-3xl font-bold mb-6">My Tickets</h1>

      <!-- Loading -->
      <div v-if="loading" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
      </div>

      <!-- Empty State -->
      <div v-else-if="myOrders.length === 0" class="text-center py-12">
        <svg class="w-24 h-24 mx-auto text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                d="M15 5v2m0 4v2m0 4v2M5 5a2 2 0 00-2 2v3a2 2 0 110 4v3a2 2 0 002 2h14a2 2 0 002-2v-3a2 2 0 110-4V7a2 2 0 00-2-2H5z" />
        </svg>
        <h3 class="text-xl font-semibold text-gray-700 mb-2">No tickets yet</h3>
        <p class="text-gray-500 mb-6">Purchase tickets to events you want to attend</p>
        <NuxtLink to="/" class="text-primary hover:underline">
          Browse Events
        </NuxtLink>
      </div>

      <!-- Tickets List -->
      <div v-else class="space-y-4">
        <div
          v-for="order in myOrders"
          :key="order.id"
          class="bg-white rounded-lg shadow-md p-6"
        >
          <div class="flex justify-between items-start">
            <div class="flex-1">
              <h3 class="text-xl font-semibold mb-2">
                {{ order.ticket.event.title }}
              </h3>
              <p class="text-gray-600 mb-1">
                {{ formatDate(order.ticket.event.event_date) }}
              </p>
              <p class="text-gray-600 mb-1">
                {{ order.ticket.event.venue }}
              </p>
              <p class="text-sm text-gray-500 mb-4">
                Order ID: {{ order.id.slice(0, 8) }}
              </p>

              <div class="flex items-center gap-4">
                <span class="text-lg font-semibold">
                  {{ order.quantity }} × {{ formatPrice(order.ticket.price) }} ETB
                </span>
                <span
                  :class="{
                    'bg-green-100 text-green-800': order.status === 'completed',
                    'bg-yellow-100 text-yellow-800': order.status === 'pending',
                    'bg-red-100 text-red-800': order.status === 'failed',
                  }"
                  class="px-3 py-1 rounded-full text-sm font-medium"
                >
                  {{ order.status }}
                </span>
              </div>
            </div>

            <!-- QR Code Placeholder -->
            <div v-if="order.status === 'completed'" class="ml-4">
              <div class="w-32 h-32 bg-gray-100 rounded-lg flex items-center justify-center">
                <svg class="w-16 h-16 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                        d="M12 4v1m6 11h2m-6 0h-2v4m0-11v3m0 0h.01M12 12h4.01M16 20h4M4 12h4m12 0h.01M5 8h2a1 1 0 001-1V5a1 1 0 00-1-1H5a1 1 0 00-1 1v2a1 1 0 001 1zm12 0h2a1 1 0 001-1V5a1 1 0 00-1-1h-2a1 1 0 00-1 1v2a1 1 0 001 1zM5 20h2a1 1 0 001-1v-2a1 1 0 00-1-1H5a1 1 0 00-1 1v2a1 1 0 001 1z" />
                </svg>
              </div>
              <p class="text-xs text-gray-500 text-center mt-2">QR Code</p>
            </div>
          </div>

          <!-- Order Details -->
          <div class="mt-4 pt-4 border-t">
            <p class="text-sm text-gray-600">
              Purchased: {{ formatDate(order.created_at) }}
            </p>
            <p class="text-sm text-gray-600">
              Total: {{ formatPrice(order.total_price) }} ETB
            </p>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { useQuery } from '@vue/apollo-composable'
import { gql } from '@apollo/client/core'

definePageMeta({
  middleware: 'auth'
})

const { user, logout } = useAuth()

const MY_ORDERS = gql`
  query MyOrders($user_id: uuid!) {
    orders(where: { user_id: { _eq: $user_id } }, order_by: { created_at: desc }) {
      id
      quantity
      total_price
      status
      created_at
      ticket {
        price
        event {
          title
          venue
          event_date
        }
      }
    }
  }
`

const { result, loading } = useQuery(MY_ORDERS, { user_id: user.value?.id })
const myOrders = computed(() => result.value?.orders || [])

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('en-ET').format(price)
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('en-ET', {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}
</script>
