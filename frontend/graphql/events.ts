import { gql } from '@apollo/client/core'

export const GET_EVENTS = gql`
  query GetEvents(
    $where: events_bool_exp
    $order_by: [events_order_by!]
    $limit: Int
    $offset: Int
  ) {
    events(
      where: $where
      order_by: $order_by
      limit: $limit
      offset: $offset
    ) {
      id
      title
      description
      venue
      address
      latitude
      longitude
      price
      event_date
      created_at
      category {
        id
        name
        slug
      }
      event_images(where: { is_featured: { _eq: true } }, limit: 1) {
        url
      }
      event_tags {
        tag {
          id
          name
        }
      }
      user {
        full_name
      }
    }
    events_aggregate(where: $where) {
      aggregate {
        count
      }
    }
  }
`

export const GET_NEARBY_EVENTS = gql`
  query GetNearbyEvents($lat: numeric!, $lng: numeric!, $radius_km: numeric!) {
    get_nearby_events(args: { lat: $lat, lng: $lng, radius_km: $radius_km }) {
      event_id
      distance_km
    }
  }
`

export const GET_CATEGORIES = gql`
  query GetCategories {
    categories(order_by: { name: asc }) {
      id
      name
      slug
    }
  }
`
