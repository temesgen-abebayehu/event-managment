<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center px-4">
    <div class="max-w-md w-full">
      <!-- Logo/Header -->
      <div class="text-center mb-8">
        <NuxtLink to="/" class="text-3xl font-bold text-primary">
          EventHub Ethiopia
        </NuxtLink>
        <p class="text-gray-600 mt-2">Login to your account</p>
      </div>

      <!-- Login Form -->
      <div class="bg-white rounded-lg shadow-lg p-8">
        <form @submit.prevent="handleLogin">
          <!-- Email -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Email
            </label>
            <input
              v-model="form.email"
              type="email"
              required
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
              placeholder="your@email.com"
            />
          </div>

          <!-- Password -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Password
            </label>
            <input
              v-model="form.password"
              type="password"
              required
              minlength="8"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
              placeholder="••••••••"
            />
          </div>

          <!-- Error Message -->
          <div v-if="error" class="mb-4 p-3 bg-red-50 text-red-700 rounded-lg text-sm">
            {{ error }}
          </div>

          <!-- Submit Button -->
          <button
            type="submit"
            :disabled="loading"
            class="w-full bg-primary text-white py-3 rounded-lg font-semibold hover:bg-primary-dark transition disabled:opacity-50"
          >
            {{ loading ? 'Logging in...' : 'Login' }}
          </button>
        </form>

        <!-- Signup Link -->
        <p class="text-center mt-6 text-gray-600">
          Don't have an account?
          <NuxtLink to="/auth/signup" class="text-primary hover:underline font-medium">
            Sign up
          </NuxtLink>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const config = useRuntimeConfig()
const { setAuth } = useAuth()

const form = ref({
  email: '',
  password: '',
})

const loading = ref(false)
const error = ref('')

const handleLogin = async () => {
  loading.value = true
  error.value = ''

  try {
    const response = await $fetch(`${config.public.backendUrl}/actions/login`, {
      method: 'POST',
      body: {
        input: {
          email: form.value.email,
          password: form.value.password,
        },
      },
    })

    if (response.access_token && response.user) {
      setAuth(response.access_token, response.user)
      navigateTo('/dashboard')
    }
  } catch (err: any) {
    error.value = err.data?.error || 'Invalid email or password'
  } finally {
    loading.value = false
  }
}
</script>
