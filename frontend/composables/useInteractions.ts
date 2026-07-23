import { useMutation, useQuery } from '@vue/apollo-composable'
import {
  TOGGLE_BOOKMARK,
  REMOVE_BOOKMARK,
  TOGGLE_FOLLOW,
  REMOVE_FOLLOW,
  CHECK_USER_INTERACTIONS,
} from '~/graphql/interactions'

export const useEventInteractions = (eventId: Ref<string | undefined>) => {
  const { user, isAuthenticated } = useAuth()
  const { error: showError } = useToast()

  // Check if user has bookmarked/followed
  const { result, refetch } = useQuery(
    CHECK_USER_INTERACTIONS,
    () => ({
      event_id: eventId.value,
      user_id: user.value?.id,
    }),
    () => ({ enabled: isAuthenticated.value && !!eventId.value })
  )

  const isBookmarked = computed(() => {
    return result.value?.bookmarks?.length > 0
  })

  const isFollowing = computed(() => {
    return result.value?.follows?.length > 0
  })

  // Mutations
  const { mutate: addBookmark } = useMutation(TOGGLE_BOOKMARK)
  const { mutate: deleteBookmark } = useMutation(REMOVE_BOOKMARK)
  const { mutate: addFollow } = useMutation(TOGGLE_FOLLOW)
  const { mutate: deleteFollow } = useMutation(REMOVE_FOLLOW)

  const toggleBookmark = async () => {
    if (!isAuthenticated.value) {
      navigateTo('/auth/login')
      return
    }

    if (!eventId.value) return

    try {
      if (isBookmarked.value) {
        await deleteBookmark({ event_id: eventId.value })
      } else {
        await addBookmark({ event_id: eventId.value })
      }
      refetch()
    } catch (error) {
      showError('Failed to update bookmark. Please try again.')
    }
  }

  const toggleFollow = async () => {
    if (!isAuthenticated.value) {
      navigateTo('/auth/login')
      return
    }

    if (!eventId.value) return

    try {
      if (isFollowing.value) {
        await deleteFollow({ event_id: eventId.value })
      } else {
        await addFollow({ event_id: eventId.value })
      }
      refetch()
    } catch (error) {
      showError('Failed to update follow. Please try again.')
    }
  }

  return {
    isBookmarked,
    isFollowing,
    toggleBookmark,
    toggleFollow,
  }
}
