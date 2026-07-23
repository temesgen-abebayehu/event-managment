export default defineNuxtPlugin(() => {
  const { restoreUserFromToken } = useAuth()
  
  // Restore user state from token on app initialization
  restoreUserFromToken()
})
