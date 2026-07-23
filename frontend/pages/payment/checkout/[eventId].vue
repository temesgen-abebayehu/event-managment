<template>
  <div>
    <div class="container mx-auto px-4 py-8 max-w-2xl">
      <div v-if="loadingEvent" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
      </div>

      <div v-else-if="event" class="bg-white rounded-lg shadow-lg p-8">
        <!-- Back Button -->
        <button
          @click="$router.back()"
          class="flex items-center gap-2 text-gray-600 hover:text-purple-600 mb-6 transition-colors group"
        >
          <svg class="w-5 h-5 transform group-hover:-translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
          <span class="font-medium">Back</span>
        </button>

        <h1 class="text-3xl font-bold mb-6">Checkout</h1>

        <!-- Event Summary -->
        <div class="border-b pb-6 mb-6">
          <h2 class="text-xl font-semibold mb-4">{{ event.title }}</h2>
          <p class="text-gray-600 mb-2">{{ formatDate(event.event_date) }}</p>
          <p class="text-gray-600">{{ event.venue }}</p>
        </div>

        <!-- Quantity Selection -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Number of Tickets
          </label>
          <input
            v-model.number="quantity"
            type="number"
            min="1"
            :max="ticketsRemaining"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
          />
          <p class="mt-1 text-sm text-gray-600">
            {{ ticketsRemaining }} tickets available
          </p>
        </div>

        <!-- Price Summary -->
        <div class="bg-gray-50 rounded-lg p-4 mb-6">
          <div class="flex justify-between mb-2">
            <span>Price per ticket</span>
            <span class="font-semibold">{{ formatPrice(event.price) }} ETB</span>
          </div>
          <div class="flex justify-between mb-2">
            <span>Quantity</span>
            <span class="font-semibold">{{ quantity }}</span>
          </div>
          <div class="border-t pt-2 flex justify-between">
            <span class="font-bold">Total</span>
            <span class="font-bold text-xl text-primary">
              {{ formatPrice(totalPrice) }} ETB
            </span>
          </div>
        </div>

        <!-- Error Message -->
        <div v-if="error" class="mb-6 p-3 bg-red-50 text-red-700 rounded-lg text-sm">
          {{ error }}
        </div>

        <!-- Action Buttons -->
        <div class="flex gap-4">
          <button
            @click="handlePayment"
            :disabled="loading"
            class="flex-1 bg-primary text-white py-3 rounded-lg font-semibold hover:bg-primary-dark transition disabled:opacity-50"
          >
            {{ loading ? 'Processing...' : 'Pay with Chapa' }}
          </button>
          <button
            @click="$router.back()"
            class="px-6 py-3 border border-gray-300 rounded-lg hover:bg-gray-50 transition"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useQuery } from '@vue/apollo-composable'
import { gql } from '@apollo/client/core'

definePageMeta({
  middleware: 'auth'
})

useHead({ title: 'Checkout' })

const route = useRoute()
const config = useRuntimeConfig()
const eventId = route.params.eventId

const GET_EVENT = gql`
  query GetEvent($id: uuid!) {
    events_by_pk(id: $id) {
      id
      title
      venue
      event_date
      price
      tickets {
        id
        quantity_total
        quantity_sold
      }
    }
  }
`

const { result, loading: loadingEvent } = useQuery(GET_EVENT, { id: eventId })
const event = computed(() => result.value?.events_by_pk)

const ticket = computed(() => event.value?.tickets?.[0])
const ticketsRemaining = computed(() => {
  if (!ticket.value) return 0
  return ticket.value.quantity_total - ticket.value.quantity_sold
})

const quantity = ref(1)
const loading = ref(false)
const error = ref('')

// Validate quantity doesn't exceed available tickets
watch([quantity, ticketsRemaining], () => {
  if (quantity.value > ticketsRemaining.value) {
    quantity.value = ticketsRemaining.value
  }
  if (quantity.value < 1) {
    quantity.value = 1
  }
})

const totalPrice = computed(() => {
  return (event.value?.price || 0) * quantity.value
})

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('en-ET').format(price)
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('en-ET', {
    weekday: 'long',
    month: 'long',
    day: 'numeric',
    year: 'numeric',
  })
}

const handlePayment = async () => {
  // Validate tickets available
  if (ticketsRemaining.value === 0) {
    error.value = 'Sorry, this event is sold out'
    return
  }

  if (quantity.value > ticketsRemaining.value) {
    error.value = `Only ${ticketsRemaining.value} tickets available`
    return
  }

  loading.value = true
  error.value = ''

  try {
    const response = await $fetch(`${config.public.backendUrl}/actions/initiate-payment`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${useCookie('auth_token').value}`,
      },
      body: {
        input: {
          event_id: eventId,
          quantity: quantity.value,
        },
        session_variables: {
          'x-hasura-user-id': useAuth().user.value?.id,
        },
      },
    })

    if (response.checkout_url) {
      // Redirect to Chapa payment page
      window.location.href = response.checkout_url
    }
  } catch (err: any) {
    error.value = err.data?.error || 'Failed to initialize payment'
  } finally {
    loading.value = false
  }
}
</script>
