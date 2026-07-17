// Toast notification composable
export const useToast = () => {
  const toasts = useState<Array<{ id: string; message: string; type: 'success' | 'error' | 'info' }>>('toasts', () => [])

  const addToast = (message: string, type: 'success' | 'error' | 'info' = 'info', duration = 3000) => {
    const id = Math.random().toString(36).substring(7)
    toasts.value.push({ id, message, type })

    if (duration > 0) {
      setTimeout(() => {
        removeToast(id)
      }, duration)
    }
  }

  const removeToast = (id: string) => {
    toasts.value = toasts.value.filter(t => t.id !== id)
  }

  const success = (message: string, duration = 3000) => addToast(message, 'success', duration)
  const error = (message: string, duration = 5000) => addToast(message, 'error', duration)
  const info = (message: string, duration = 3000) => addToast(message, 'info', duration)

  return {
    toasts,
    addToast,
    removeToast,
    success,
    error,
    info,
  }
}
