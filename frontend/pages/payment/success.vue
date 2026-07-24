<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center">
    <div class="max-w-md w-full mx-4">
      <!-- Loading State -->
      <div v-if="verifying" class="bg-white rounded-lg shadow-lg p-8 text-center">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary mb-4"></div>
        <p class="text-gray-600">Verifying your payment...</p>
      </div>

      <!-- Success State -->
      <div v-else-if="verified" class="bg-white rounded-lg shadow-lg p-8 text-center">
        <!-- Success Icon -->
        <div class="w-20 h-20 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-6">
          <svg class="w-10 h-10 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
        </div>

        <h1 class="text-2xl font-bold mb-4">Payment Successful!</h1>
        <p class="text-gray-600 mb-6">
          Your payment has been processed successfully. Your tickets are now confirmed.
        </p>

        <!-- Order Details -->
        <div v-if="orderDetails" class="bg-gray-50 rounded-lg p-4 mb-6 text-left">
          <p class="text-sm text-gray-600 mb-2">Transaction Reference</p>
          <p class="font-mono text-sm mb-3">{{ orderDetails.tx_ref }}</p>
          <p class="text-sm text-gray-600 mb-2">Quantity</p>
          <p class="font-semibold mb-3">{{ orderDetails.quantity }} ticket(s)</p>
          <p class="text-sm text-gray-600 mb-2">Total Paid</p>
          <p class="font-bold text-lg text-green-600">{{ formatPrice(orderDetails.total_price) }} ETB</p>
        </div>

        <!-- Actions -->
        <div class="flex flex-col gap-3">
          <NuxtLink
            to="/dashboard/tickets"
            class="bg-primary text-white py-3 rounded-lg font-semibold hover:bg-primary-dark transition text-center"
          >
            View My Tickets
          </NuxtLink>
          <NuxtLink
            to="/"
            class="text-gray-600 hover:text-primary"
          >
            Browse More Events
          </NuxtLink>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-white rounded-lg shadow-lg p-8 text-center">
        <div class="w-20 h-20 bg-yellow-100 rounded-full flex items-center justify-center mx-auto mb-6">
          <svg class="w-10 h-10 text-yellow-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
        </div>

        <h1 class="text-2xl font-bold mb-4">Verification Issue</h1>
        <p class="text-gray-600 mb-6">{{ errorMessage }}</p>

        <div class="flex flex-col gap-3">
          <NuxtLink
            to="/dashboard/tickets"
            class="bg-primary text-white py-3 rounded-lg font-semibold hover:bg-primary-dark transition text-center"
          >
            Check My Tickets
          </NuxtLink>
          <NuxtLink
            to="/"
            class="text-gray-600 hover:text-primary"
          >
            Back to Home
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'blank',
  middleware: 'auth'
})

useHead({ title: 'Payment Verification' })

const route = useRoute()
const verifying = ref(true)
const verified = ref(false)
const error = ref(false)
const errorMessage = ref('')
const orderDetails = ref<any>(null)

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('en-ET').format(price)
}

// Verify payment on mount
onMounted(async () => {
  let txRef = route.query.tx_ref || route.query.trx_ref || route.query.txRef || route.query.transaction_id
  
  if (!txRef) {
    txRef = localStorage.getItem('pending_tx_ref')
  }
  
  if (!txRef) {
    const hash = window.location.hash
    const hashParams = new URLSearchParams(hash.substring(1))
    const hashTxRef = hashParams.get('tx_ref') || hashParams.get('trx_ref')
    
    if (hashTxRef) {
      await verifyPayment(hashTxRef as string)
      return
    }
    
    error.value = true
    errorMessage.value = 'Unable to verify payment. Please check "My Tickets" to confirm your order status, or contact support if you were charged.'
    verifying.value = false
    return
  }

  localStorage.removeItem('pending_tx_ref')
  localStorage.removeItem('pending_order_id')
  
  await verifyPayment(txRef as string)
})

async function verifyPayment(txRef: string) {
  try {
    const response = await fetch('http://localhost:3001/actions/verify-payment', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        input: {
          tx_ref: txRef
        }
      })
    })

    if (!response.ok) {
      throw new Error('Failed to verify payment')
    }

    const data = await response.json()
    
    if (data.status === 'completed') {
      verified.value = true
      orderDetails.value = data.order
    } else {
      error.value = true
      errorMessage.value = 'Payment was not successful. Please try again or contact support.'
    }
  } catch (err) {
    error.value = true
    errorMessage.value = 'Unable to verify payment. Please check "My Tickets" to confirm your order status.'
  } finally {
    verifying.value = false
  }
}
</script>
