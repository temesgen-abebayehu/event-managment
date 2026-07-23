import { useQuery } from '@vue/apollo-composable'
import { GET_EVENTS, GET_CATEGORIES } from '~/graphql/events'

export const useEvents = (filters: any = {}, searchTerm: Ref<string> = ref('')) => {
  const buildWhere = () => {
    const where: any = {}

    // Search functionality - multi-word search
    if (searchTerm.value && searchTerm.value.trim().length > 0) {
      const terms = searchTerm.value.trim().split(/\s+/)
      const conditions = terms.map(term => ({
        _or: [
          { title: { _ilike: `%${term}%` } },
          { description: { _ilike: `%${term}%` } },
          { venue: { _ilike: `%${term}%` } },
          { address: { _ilike: `%${term}%` } },
        ]
      }))
      where._and = conditions
    }

    // Category filter
    if (filters.value?.category_id) {
      where.category_id = { _eq: filters.value.category_id }
    }

    // Price filter
    if (filters.value?.min_price !== undefined && filters.value.min_price !== null && filters.value.min_price !== '') {
      if (!where.price) where.price = {}
      where.price._gte = Number(filters.value.min_price)
    }
    if (filters.value?.max_price !== undefined && filters.value.max_price !== null && filters.value.max_price !== '') {
      if (!where.price) where.price = {}
      where.price._lte = Number(filters.value.max_price)
    }

    // Date filter
    if (filters.value?.date_from) {
      if (!where.event_date) where.event_date = {}
      where.event_date._gte = filters.value.date_from + 'T00:00:00'
    }
    if (filters.value?.date_to) {
      if (!where.event_date) where.event_date = {}
      where.event_date._lte = filters.value.date_to + 'T23:59:59'
    }

    return where
  }

  // Reactive query - updates when filters or search changes
  const { result, loading, error, refetch } = useQuery(
    GET_EVENTS, 
    () => ({
      where: buildWhere(),
      order_by: filters.value?.order_by || [{ event_date: 'asc' }],
      limit: filters.value?.limit || 12,
      offset: filters.value?.offset || 0,
    }),
    {
      context: {
        forceAnonymous: true
      }
    }
  )

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

export const useCategories = () => {
  const { result, loading, error } = useQuery(GET_CATEGORIES, undefined, {
    context: { forceAnonymous: true }
  })

  const categories = computed(() => result.value?.categories || [])

  return {
    categories,
    loading,
    error,
  }
}
