export default defineNuxtRouteMiddleware((to, from) => {
  const { isAuthenticated, user } = useAuth()

  // First check authentication
  if (!isAuthenticated.value) {
    return navigateTo('/auth/login')
  }

  // Then check admin role
  if (user.value?.role !== 'admin') {
    return navigateTo('/dashboard')
  }
})
