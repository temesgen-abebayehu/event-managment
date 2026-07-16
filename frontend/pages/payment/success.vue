<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center">
    <div class="max-w-md w-full mx-4">
      <div class="bg-white rounded-lg shadow-lg p-8 text-center">
        <!-- Success Icon -->
        <div class="w-20 h-20 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-6">
          <svg class="w-10 h-10 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
        </div>

        <h1 class="text-2xl font-bold mb-4">Payment Successful!</h1>
        <p class="text-gray-600 mb-6">
          Your payment has been processed successfully. Your tickets will be available shortly.
        </p>

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
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: 'auth'
})
</script>
