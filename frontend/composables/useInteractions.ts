import { useMutation, useQuery } from '@vue/apollo-composable'
import {
  TOGGLE_BOOKMARK,
  REMOVE_BOOKMARK,
  TOGGLE_FOLLOW,
  REMOVE_FOLLOW,
  CHECK_USER_INTERACTIONS,
} from '~/graphql/interactions'

export const useEventInteractions = (eventId: string) => {
  const { user, isAuthenticated } = useAuth()

  // Check if user has bookmarked/followed
  const { result, refetch } = useQuery(
    CHECK_USER_INTERACTIONS,
    () => ({
      event_id: eventId,
      user_id: user.value?.id,
    }),
    () => ({ enabled: isAuthenticated.value })
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

    try {
      if (isBookmarked.value) {
        await deleteBookmark({ event_id: eventId })
      } else {
        await addBookmark({ event_id: eventId })
      }
      refetch()
    } catch (error) {
      console.error('Failed to toggle bookmark:', error)
    }
  }

  const toggleFollow = async () => {
    if (!isAuthenticated.value) {
      navigateTo('/auth/login')
      return
    }

    try {
      if (isFollowing.value) {
        await deleteFollow({ event_id: eventId })
      } else {
        await addFollow({ event_id: eventId })
      }
      refetch()
    } catch (error) {
      console.error('Failed to toggle follow:', error)
    }
  }

  return {
    isBookmarked,
    isFollowing,
    toggleBookmark,
    toggleFollow,
  }
}
