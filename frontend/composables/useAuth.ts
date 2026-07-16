export const useAuth = () => {
  const token = useCookie('auth_token')
  const user = useState('user', () => null)

  const isAuthenticated = computed(() => !!token.value)

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
  }
}
