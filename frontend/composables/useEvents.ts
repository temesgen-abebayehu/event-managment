import { useQuery } from '@vue/apollo-composable'
import { GET_EVENTS, SEARCH_EVENTS, GET_CATEGORIES } from '~/graphql/events'

export const useEvents = (filters: any = {}) => {
  const buildWhere = () => {
    const where: any = {}

    // Category filter
    if (filters.value?.category_id) {
      where.category_id = { _eq: filters.value.category_id }
    }

    // Price filter
    if (filters.value?.min_price !== undefined || filters.value?.max_price !== undefined) {
      where.price = {}
      if (filters.value.min_price !== undefined) {
        where.price._gte = filters.value.min_price
      }
      if (filters.value.max_price !== undefined) {
        where.price._lte = filters.value.max_price
      }
    }

    // Date filter
    if (filters.value?.date_from || filters.value?.date_to) {
      where.event_date = {}
      if (filters.value.date_from) {
        where.event_date._gte = filters.value.date_from
      }
      if (filters.value.date_to) {
        where.event_date._lte = filters.value.date_to
      }
    }

    return where
  }

  const { result, loading, error, refetch } = useQuery(GET_EVENTS, {
    where: buildWhere(),
    order_by: filters.value?.order_by || [{ event_date: 'asc' }],
    limit: filters.value?.limit || 12,
    offset: filters.value?.offset || 0,
  })

  const events = computed(() => result.value?.events || [])
  const totalCount = computed(() => result.value?.events_aggregate?.aggregate?.count || 0)

  return {
    events,
    totalCount,
    loading,
    error,
    refetch,
  }
}

export const useSearchEvents = (searchTerm: Ref<string>) => {
  const { result, loading, error } = useQuery(
    SEARCH_EVENTS,
    () => ({ search_term: searchTerm.value }),
    () => ({ enabled: searchTerm.value.length > 0 })
  )

  const events = computed(() => result.value?.search_events || [])

  return {
    events,
    loading,
    error,
  }
}

export const useCategories = () => {
  const { result, loading, error } = useQuery(GET_CATEGORIES)

  const categories = computed(() => result.value?.categories || [])

  return {
    categories,
    loading,
    error,
  }
}
