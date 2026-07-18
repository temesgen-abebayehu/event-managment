// Domain models matching the backend and GraphQL schema

export interface User {
  id: string
  email: string
  full_name: string
  phone?: string
  avatar_url?: string
  role?: string
  created_at?: string
  updated_at?: string
}

export interface Category {
  id: string
  name: string
  slug: string
  events_aggregate?: {
    aggregate: {
      count: number
    }
  }
}

export interface Tag {
  id: string
  name: string
}

export interface EventImage {
  url: string
  public_id?: string
  is_featured?: boolean
}

export interface EventTag {
  tag: Tag
}

export interface Event {
  id: string
  title: string
  description: string
  venue: string
  address: string
  latitude: number
  longitude: number
  price: number
  event_date: string
  created_at?: string
  category: Category
  event_images: EventImage[]
  event_tags: EventTag[]
  user: {
    full_name: string
  }
}

export interface Ticket {
  id: string
  event_id: string
  price: number
  quantity_total: number
  quantity_sold: number
  event?: Event
}

export interface Order {
  id: string
  user_id: string
  ticket_id: string
  quantity: number
  total_price: number
  status: 'pending' | 'completed' | 'failed' | 'refunded'
  chapa_tx_ref: string
  created_at: string
  ticket: {
    price: number
    event: {
      title: string
      venue: string
      event_date: string
    }
  }
}

export interface Bookmark {
  event: Event
}

export interface Follow {
  event: Event
}

export interface Toast {
  id: string
  message: string
  type: 'success' | 'error' | 'info'
}

export interface UploadedFile {
  url: string
  public_id: string
}

export interface UploadResponse {
  files: UploadedFile[]
}
