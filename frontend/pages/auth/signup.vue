<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center px-4">
    <div class="max-w-md w-full">
      <!-- Logo/Header -->
      <div class="text-center mb-8">
        <NuxtLink to="/" class="text-3xl font-bold text-primary">
          EventHub Ethiopia
        </NuxtLink>
        <p class="text-gray-600 mt-2">Create your account</p>
      </div>

      <!-- Signup Form -->
      <div class="bg-white rounded-lg shadow-lg p-8">
        <!-- Back Button -->
        <button
          type="button"
          @click="$router.push('/')"
          class="flex items-center gap-2 text-gray-600 hover:text-purple-600 mb-6 transition-colors group"
        >
          <svg class="w-5 h-5 transform group-hover:-translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
          <span class="font-medium">Back to Home</span>
        </button>

        <form @submit="onSubmit">
          <!-- Full Name -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Full Name
            </label>
            <input
              v-model="fullName"
              type="text"
              class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
              :class="errors.fullName ? 'border-red-500' : 'border-gray-300'"
              placeholder="John Doe"
            />
            <p v-if="errors.fullName" class="mt-1 text-sm text-red-600">{{ errors.fullName }}</p>
          </div>

          <!-- Email -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Email
            </label>
            <input
              v-model="email"
              type="email"
              class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
              :class="errors.email ? 'border-red-500' : 'border-gray-300'"
              placeholder="your@email.com"
            />
            <p v-if="errors.email" class="mt-1 text-sm text-red-600">{{ errors.email }}</p>
          </div>

          <!-- Password -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Password
            </label>
            <input
              v-model="password"
              type="password"
              class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
              :class="errors.password ? 'border-red-500' : 'border-gray-300'"
              placeholder="••••••••"
            />
            <p v-if="errors.password" class="mt-1 text-sm text-red-600">{{ errors.password }}</p>
            <p class="text-xs text-gray-500 mt-1">Minimum 8 characters, with uppercase, lowercase, and number</p>
          </div>

          <!-- Confirm Password -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Confirm Password
            </label>
            <input
              v-model="confirmPassword"
              type="password"
              class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-primary focus:border-transparent"
              :class="errors.confirmPassword ? 'border-red-500' : 'border-gray-300'"
              placeholder="••••••••"
            />
            <p v-if="errors.confirmPassword" class="mt-1 text-sm text-red-600">{{ errors.confirmPassword }}</p>
          </div>

          <!-- Error Message -->
          <div v-if="serverError" class="mb-4 p-3 bg-red-50 text-red-700 rounded-lg text-sm">
            {{ serverError }}
          </div>

          <!-- Submit Button -->
          <button
            type="submit"
            :disabled="loading"
            class="w-full bg-primary text-white py-3 rounded-lg font-semibold hover:bg-primary-dark transition disabled:opacity-50"
          >
            {{ loading ? 'Creating account...' : 'Sign Up' }}
          </button>
        </form>

        <!-- Login Link -->
        <p class="text-center mt-6 text-gray-600">
          Already have an account?
          <NuxtLink to="/auth/login" class="text-primary hover:underline font-medium">
            Login
          </NuxtLink>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useForm, useField } from 'vee-validate'
import * as yup from 'yup'

const config = useRuntimeConfig()
const { setAuth } = useAuth()

definePageMeta({
  layout: 'blank'
})

useHead({ title: 'Sign Up' })

const schema = yup.object({
  fullName: yup.string().required('Full name is required').min(2, 'Name must be at least 2 characters'),
  email: yup.string().required('Email is required').email('Invalid email address'),
  password: yup.string()
    .required('Password is required')
    .min(8, 'Password must be at least 8 characters')
    .matches(/[A-Z]/, 'Password must contain at least one uppercase letter')
    .matches(/[a-z]/, 'Password must contain at least one lowercase letter')
    .matches(/[0-9]/, 'Password must contain at least one number'),
  confirmPassword: yup.string()
    .required('Please confirm your password')
    .oneOf([yup.ref('password')], 'Passwords must match'),
})

const { handleSubmit, errors } = useForm({
  validationSchema: schema,
})

const { value: fullName } = useField<string>('fullName')
const { value: email } = useField<string>('email')
const { value: password } = useField<string>('password')
const { value: confirmPassword } = useField<string>('confirmPassword')

const loading = ref(false)
const serverError = ref('')

const onSubmit = handleSubmit(async (values) => {
  loading.value = true
  serverError.value = ''

  try {
    const response = await $fetch(`${config.public.backendUrl}/actions/signup`, {
      method: 'POST',
      body: {
        input: {
          email: values.email,
          password: values.password,
          full_name: values.fullName,
        },
      },
    })

    if (response.access_token && response.user) {
      setAuth(response.access_token, response.user)
      navigateTo('/dashboard')
    }
  } catch (err: any) {
    serverError.value = err.data?.error || 'Failed to create account'
  } finally {
    loading.value = false
  }
})
</script>
