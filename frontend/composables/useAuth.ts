export const useAuth = () => {
  const token = useCookie('auth_token')
  const user = useState('user', () => null)

  const isAuthenticated = computed(() => !!token.value)

  // Decode JWT token to extract user data
  const decodeToken = (jwtToken: string) => {
    try {
      const base64Url = jwtToken.split('.')[1]
      const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
      const jsonPayload = decodeURIComponent(
        atob(base64)
          .split('')
          .map((c) => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
          .join('')
      )
      return JSON.parse(jsonPayload)
    } catch (error) {
      console.error('Failed to decode token:', error)
      return null
    }
  }

  // Restore user state from token if available
  const restoreUserFromToken = () => {
    if (token.value && !user.value) {
      const decoded = decodeToken(token.value)
      if (decoded) {
        // Extract user info from JWT claims
        const hasuraClaims = decoded['https://hasura.io/jwt/claims']
        if (hasuraClaims && hasuraClaims['x-hasura-user-id']) {
          user.value = {
            id: hasuraClaims['x-hasura-user-id'],
            email: decoded.email || '',
            full_name: '', // We don't store full name in JWT, will be empty until proper fetch
          }
        }
      }
    }
  }

  // Initialize user state on composable creation
  restoreUserFromToken()

  const setAuth = (accessToken: string, userData: any) => {
    token.value = accessToken
    user.value = userData
  }

  const logout = () => {
    token.value = null
    user.value = null
    navigateTo('/auth/login')
  }

  return {
    token,
    user,
    isAuthenticated,
    setAuth,
    logout,
    restoreUserFromToken,
  }
}
