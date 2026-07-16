# EventHub Ethiopia - Frontend

Nuxt 4 frontend for Local Event Management Platform.

## Features

✅ **Public Event Browsing** - View all events without login  
✅ **Search** - Full-text search by title, description, tags  
✅ **Filters** - Category, date range, price  
✅ **Quick Filters** - Free events, Today, This week  
✅ **Grid View** - Responsive card layout  
✅ **Map View** - Interactive map with event markers  
✅ **Event Details** - Full event page with location map  
✅ **Pagination** - Navigate through events  

## Setup

```bash
cd frontend
npm install
cp .env.example .env
# Edit .env with your Hasura endpoint
npm run dev
```

Server will start on `http://localhost:3000`

## Tech Stack

- **Nuxt 4** - Vue 3 meta-framework with SSR
- **TailwindCSS** - Utility-first CSS
- **Apollo Client** - GraphQL client
- **Leaflet** - Interactive maps
- **VeeValidate** - Form validation (for auth/event forms)

## Project Structure

```
frontend/
├── app.vue                    # Root component
├── nuxt.config.ts             # Nuxt configuration
├── pages/
│   ├── index.vue             # Home page (event listing)
│   └── events/
│       └── [id].vue          # Event detail page
├── components/
│   ├── EventCard.vue         # Event card component
│   ├── SearchFilters.vue     # Search & filter UI
│   └── EventMap.vue          # Leaflet map component
├── composables/
│   ├── useAuth.ts            # Auth state management
│   └── useEvents.ts          # Event queries
├── graphql/
│   └── events.ts             # GraphQL queries
└── plugins/
    └── apollo.ts             # Apollo client setup
```

## Features Implemented (Phase 6)

### Search
- Full-text search using PostgreSQL `search_vector`
- Debounced search input (500ms)
- Real-time results

### Filters
- **Category** - Filter by event category
- **Date Range** - From/to date picker
- **Price** - Max price filter
- **Quick Filters** - Free, Today, This Week

### Views
- **Grid View** - Responsive 3-column layout
- **Map View** - Leaflet map with event markers
- Click markers to view event details

### Event Display
- Event cards with image, category, date, price
- Event detail page with full info + location map
- Responsive design (mobile-first)

## Next Steps

- Phase 7: Event Creation Form
- Phase 8: Auth Pages (Login/Signup)
- Phase 9: Event Edit/Delete
- Phase 10: Bookmarks & Follows
- Phase 11: Payment Integration
- Phase 12: UI/UX Polish

## Environment Variables

See `.env.example` - only needs Hasura GraphQL endpoint.
