import { gql } from '@apollo/client/core'

export const TOGGLE_BOOKMARK = gql`
  mutation ToggleBookmark($event_id: uuid!) {
    insert_bookmarks_one(
      object: { event_id: $event_id }
      on_conflict: {
        constraint: bookmarks_pkey
        update_columns: []
      }
    ) {
      event_id
    }
  }
`

export const REMOVE_BOOKMARK = gql`
  mutation RemoveBookmark($event_id: uuid!) {
    delete_bookmarks(where: { event_id: { _eq: $event_id } }) {
      affected_rows
    }
  }
`

export const TOGGLE_FOLLOW = gql`
  mutation ToggleFollow($event_id: uuid!) {
    insert_follows_one(
      object: { event_id: $event_id }
      on_conflict: {
        constraint: follows_pkey
        update_columns: []
      }
    ) {
      event_id
    }
  }
`

export const REMOVE_FOLLOW = gql`
  mutation RemoveFollow($event_id: uuid!) {
    delete_follows(where: { event_id: { _eq: $event_id } }) {
      affected_rows
    }
  }
`

export const GET_USER_BOOKMARKS = gql`
  query GetUserBookmarks($user_id: uuid!) {
    bookmarks(where: { user_id: { _eq: $user_id } }) {
      event {
        id
        title
        description
        venue
        event_date
        price
        category {
          name
        }
        event_images(where: { is_featured: { _eq: true } }, limit: 1) {
          url
        }
      }
    }
  }
`

export const GET_USER_FOLLOWS = gql`
  query GetUserFollows($user_id: uuid!) {
    follows(where: { user_id: { _eq: $user_id } }) {
      event {
        id
        title
        description
        venue
        event_date
        price
        category {
          name
        }
        event_images(where: { is_featured: { _eq: true } }, limit: 1) {
          url
        }
      }
    }
  }
`

export const CHECK_USER_INTERACTIONS = gql`
  query CheckUserInteractions($event_id: uuid!, $user_id: uuid!) {
    bookmarks(where: { event_id: { _eq: $event_id }, user_id: { _eq: $user_id } }) {
      event_id
    }
    follows(where: { event_id: { _eq: $event_id }, user_id: { _eq: $user_id } }) {
      event_id
    }
  }
`
