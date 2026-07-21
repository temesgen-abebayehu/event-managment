<template>
  <div class="min-h-screen bg-gray-50 flex flex-col">
    <!-- Header -->
    <header class="bg-white shadow-sm sticky top-0 z-40">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <NuxtLink to="/" class="text-2xl font-bold text-primary hover:text-primary-dark transition-colors">
            EventHub
          </NuxtLink>

          <!-- Navigation -->
          <nav class="flex items-center gap-2 md:gap-4">
            <template v-if="isAuthenticated">
              <NuxtLink
                to="/"
                class="hidden md:inline-block px-3 py-2 text-sm font-medium rounded-lg transition-colors"
                :class="isActiveRoute('/') ? 'text-primary bg-primary/10' : 'text-gray-600 hover:text-primary hover:bg-gray-100'"
              >
                Browse
              </NuxtLink>
              <NuxtLink
                to="/dashboard"
                class="hidden md:inline-block px-3 py-2 text-sm font-medium rounded-lg transition-colors"
                :class="isActiveRoute('/dashboard') ? 'text-primary bg-primary/10' : 'text-gray-600 hover:text-primary hover:bg-gray-100'"
              >
                My Events
              </NuxtLink>
              <NuxtLink
                to="/dashboard/bookmarks"
                class="hidden md:inline-block px-3 py-2 text-sm font-medium rounded-lg transition-colors"
                :class="isActiveRoute('/dashboard/bookmarks') ? 'text-primary bg-primary/10' : 'text-gray-600 hover:text-primary hover:bg-gray-100'"
              >
                Bookmarks
              </NuxtLink>
              <NuxtLink
                to="/dashboard/tickets"
                class="hidden md:inline-block px-3 py-2 text-sm font-medium rounded-lg transition-colors"
                :class="isActiveRoute('/dashboard/tickets') ? 'text-primary bg-primary/10' : 'text-gray-600 hover:text-primary hover:bg-gray-100'"
              >
                Tickets
              </NuxtLink>

              <!-- User Menu -->
              <div class="flex items-center gap-3 ml-2 pl-2 md:pl-4 border-l border-gray-200">
                <span class="hidden md:inline text-sm font-medium text-gray-700">
                  {{ user?.full_name }}
                </span>
                <button
                  @click="logout"
                  class="px-3 py-2 text-sm font-medium text-gray-600 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
                >
                  Logout
                </button>
              </div>
            </template>

            <template v-else>
              <NuxtLink
                to="/auth/login"
                class="px-3 py-2 text-sm font-medium text-gray-600 hover:text-primary rounded-lg transition-colors"
              >
                Login
              </NuxtLink>
              <NuxtLink
                to="/auth/signup"
                class="px-4 py-2 text-sm font-semibold bg-primary text-white rounded-lg hover:bg-primary-dark transition-colors"
              >
                Sign Up
              </NuxtLink>
            </template>
          </nav>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="flex-1">
      <slot />
    </main>

    <!-- Footer -->
    <footer class="bg-gray-800 text-white py-8 mt-16">
      <div class="container mx-auto px-4 text-center">
        <p class="text-gray-400">&copy; {{ new Date().getFullYear() }} EventHub Ethiopia. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
const route = useRoute()
const { isAuthenticated, user, logout } = useAuth()

const isActiveRoute = (path: string) => {
  if (path === '/') return route.path === '/'
  return route.path.startsWith(path)
}
</script>
