# EventHub Ethiopia 🎉

> A modern, full-featured event management platform for Ethiopia built with Golang, Hasura, PostgreSQL, Nuxt 4, and Vue 3.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)](https://go.dev/)
[![Nuxt](https://img.shields.io/badge/Nuxt-3.9-00DC82?logo=nuxt.js)](https://nuxt.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-316192?logo=postgresql)](https://www.postgresql.org/)

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Database Migrations](#database-migrations)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

---

## 🎯 Overview

EventHub Ethiopia is a comprehensive event management platform that allows users to discover, create, manage, and purchase tickets for events across Ethiopia. The platform features a modern UI, real-time search, location-based filtering, payment integration via Chapa, and image management through Cloudinary.

**Key Highlights:**
- 🔐 JWT-based authentication with secure password hashing
- 🗺️ Interactive map-based location selection and event browsing
- 🏷️ Tag-based event categorization and search
- 💳 Integrated Chapa payment gateway for ticket purchases
- 📸 Multi-image upload with Cloudinary integration
- 🔍 Full-text search with PostgreSQL
- 📱 Fully responsive mobile-first design
- ⚡ Real-time GraphQL API via Hasura

---

## ✨ Features

### For All Users (Public)
- 🔍 Browse and search events by title, description, and tags
- 🗂️ Filter events by category, date, price, and location
- 🗺️ View events on an interactive map
- 📅 View event details with images, description, and venue information

### For Authenticated Users
- 👤 Secure signup and login
- 🎫 Create, edit, and delete events
- 📸 Upload multiple images (up to 5 per event, 5MB each)
- 🗺️ Set event location using interactive map picker
- 🏷️ Add multiple tags to events
- ❤️ Bookmark favorite events
- 👁️ Follow events for updates
- 💰 Purchase tickets with Chapa payment integration
- 🎟️ View purchased tickets in dashboard

### For Admins
- 📂 Manage categories (CRUD operations)
- 📊 View event statistics

---

## 🛠️ Tech Stack

### Backend
- **Language:** Golang 1.22+
- **GraphQL Engine:** Hasura
- **Database:** PostgreSQL 15+
- **Authentication:** JWT with bcrypt
- **File Storage:** Cloudinary
- **Payment Gateway:** Chapa API
- **Architecture:** Clean Architecture (Domain → Usecase → Repository → Delivery)

### Frontend
- **Framework:** Nuxt 4 (Vue 3)
- **UI Library:** TailwindCSS
- **GraphQL Client:** Vue Apollo
- **Form Validation:** VeeValidate + Yup
- **Maps:** Leaflet
- **State Management:** Vue Composition API

### DevOps
- **Containerization:** Docker & Docker Compose
- **Version Control:** Git

---

## 📁 Project Structure

```
event-management/
├── backend/                    # Golang backend
│   ├── cmd/
│   │   └── main.go            # Application entry point
│   ├── config/                # Configuration management
│   ├── delivery/              # HTTP handlers (Hasura actions)
│   ├── domain/                # Business entities
│   ├── infrastructure/        # External services (Cloudinary, Chapa, Hasura)
│   ├── repository/            # Data access layer
│   ├── usecase/              # Business logic
│   ├── utils/                # Utilities (JWT, password)
│   ├── migrations/           # SQL migrations
│   ├── go.mod
│   └── .env.example
│
├── frontend/                  # Nuxt 4 frontend
│   ├── pages/                # File-based routing
│   ├── components/           # Reusable Vue components
│   ├── composables/          # Vue composables
│   ├── graphql/              # GraphQL queries/mutations
│   ├── middleware/           # Route middleware
│   ├── plugins/              # Nuxt plugins
│   ├── assets/               # Static assets
│   ├── nuxt.config.ts
│   ├── tailwind.config.js
│   └── package.json
│
├── docker-compose.yml        # Docker services configuration
├── .gitignore
└── README.md
```

---

## 📦 Prerequisites

Before you begin, ensure you have the following installed:

- **Docker & Docker Compose** (v20.10+)
- **Node.js** (v18+ for frontend development)
- **Go** (v1.22+ for backend development)
- **Git**

### External Services (Required)
- **Cloudinary Account** - [Sign up here](https://cloudinary.com/)
- **Chapa Account** - [Sign up here](https://chapa.co/)

---

## 🚀 Installation

### 1. Clone the Repository

```bash
git clone <repository-url>
cd event-management
```

### 2. Backend Setup

```bash
cd backend

# Copy environment file
cp .env.example .env

# Edit .env with your credentials
# nano .env  # or use your preferred editor

# Install Go dependencies
go mod download
```

### 3. Frontend Setup

```bash
cd frontend

# Copy environment file
cp .env.example .env

# Edit .env with your credentials
# nano .env

# Install npm dependencies
npm install
```

---

## ⚙️ Configuration

### Backend Environment Variables

Edit `backend/.env`:

```env
# Database
DATABASE_URL=postgresql://postgres:postgres@localhost:5432/events_db

# Cloudinary
CLOUDINARY_CLOUD_NAME=your-cloud-name
CLOUDINARY_API_KEY=your-api-key
CLOUDINARY_API_SECRET=your-api-secret

# Chapa
CHAPA_SECRET_KEY=your-chapa-secret-key
CHAPA_PUBLIC_KEY=your-chapa-public-key
CHAPA_CALLBACK_URL=http://localhost:3001/api/payment/callback
CHAPA_RETURN_URL=http://localhost:3000/payment/success

# JWT
JWT_SECRET=your-super-secret-jwt-key-min-32-characters-long

# Hasura
HASURA_ADMIN_SECRET=your-hasura-admin-secret
HASURA_ENDPOINT=http://localhost:8080/v1/graphql

# Server
PORT=3001
```

### Frontend Environment Variables

Edit `frontend/.env`:

```env
BACKEND_URL=http://localhost:3001
HASURA_GRAPHQL_URL=http://localhost:8080/v1/graphql
```

---

## 🚀 Quick Start

```bash
# 1. Setup environment
cp .env.example .env
# Edit .env with your Cloudinary and Chapa credentials

# 2. Start everything
make build
make up

# 3. Run migrations
make migrate
```

**Done! Access:**
- Frontend: http://localhost:3000
- Backend: http://localhost:3001
- Hasura: http://localhost:8080/console (password in .env)

**Commands:**
```bash
make logs    # View logs
make down    # Stop
make clean   # Remove everything
```

---

## 🗄️ Database Migrations

### Apply Migrations

```bash
cd backend/migrations

# Using psql
psql -U postgres -d events_db -f 001_create_users.sql
psql -U postgres -d events_db -f 002_create_categories.sql
psql -U postgres -d events_db -f 003_create_events.sql
psql -U postgres -d events_db -f 004_create_event_images.sql
psql -U postgres -d events_db -f 005_create_tags.sql
psql -U postgres -d events_db -f 006_create_bookmarks_follows.sql
psql -U postgres -d events_db -f 007_create_tickets.sql
psql -U postgres -d events_db -f 008_create_functions.sql
psql -U postgres -d events_db -f 009_create_triggers.sql

# Or apply all at once
cat *.sql | psql -U postgres -d events_db
```

### Database Schema Overview

**Core Tables:**
- `users` - User accounts
- `categories` - Event categories (dynamic CRUD)
- `events` - Event information
- `event_images` - Multiple images per event
- `tags` - Reusable tags
- `event_tags` - Many-to-many relationship
- `bookmarks` - User bookmarks
- `follows` - User follows
- `tickets` - Ticket templates
- `orders` - Ticket purchases

**PostgreSQL Features:**
- ✅ Full-text search with generated `search_vector`
- ✅ Triggers for auto-updating timestamps
- ✅ Functions for complex queries (search, nearby events)
- ✅ Proper indexes on FK and search columns

---

## 📚 API Documentation

### Hasura Actions (Golang Endpoints)

#### Authentication

**Signup**
```graphql
mutation Signup($input: SignupInput!) {
  signup(input: $input) {
    access_token
    user {
      id
      email
      full_name
    }
  }
}
```

**Login**
```graphql
mutation Login($input: LoginInput!) {
  login(input: $input) {
    access_token
    user {
      id
      email
      full_name
    }
  }
}
```

#### File Upload

**Upload Images**
```bash
POST /api/upload
Content-Type: multipart/form-data
Authorization: Bearer <token>

files: [File, File, ...]
user_id: uuid
event_id: uuid
```

#### Payment

**Initialize Payment**
```bash
POST /api/payment/initialize
Authorization: Bearer <token>

{
  "event_id": "uuid",
  "quantity": 1,
  "user_email": "user@example.com",
  "user_name": "John Doe"
}
```

**Payment Callback**
```bash
POST /api/payment/callback

{
  "tx_ref": "string",
  "status": "success"
}
```

### GraphQL Queries

All other operations use standard Hasura GraphQL queries:

```graphql
# Get Events
query GetEvents {
  events(order_by: { event_date: asc }) {
    id
    title
    description
    venue
    price
    event_date
    category { name }
    event_images { url is_featured }
    event_tags { tag { name } }
  }
}

# Search Events
query SearchEvents($search_term: String!) {
  search_events(args: { search_term: $search_term }) {
    id
    title
    description
  }
}

# Get Nearby Events
query GetNearbyEvents($lat: numeric!, $lng: numeric!, $radius_km: numeric!) {
  get_nearby_events(args: { lat: $lat, lng: $lng, radius_km: $radius_km }) {
    event_id
    distance_km
  }
}
```

---

## 🚢 Deployment

### Production Checklist

- [ ] Update environment variables for production URLs
- [ ] Set strong JWT secret (32+ characters)
- [ ] Configure HTTPS/SSL
- [ ] Set up Cloudinary production account
- [ ] Set up Chapa production keys
- [ ] Configure Hasura permissions
- [ ] Set up database backups
- [ ] Configure monitoring (logs, errors)
- [ ] Set up CDN for static assets
- [ ] Enable rate limiting
- [ ] Review security headers

### Docker Production Build

```bash
# Build images
docker-compose -f docker-compose.prod.yml build

# Start services
docker-compose -f docker-compose.prod.yml up -d

# View logs
docker-compose -f docker-compose.prod.yml logs -f
```

### Frontend Production Build

```bash
cd frontend

# Build for production
npm run build

# Preview production build
npm run preview

# Or deploy to Vercel/Netlify
# Follow their deployment guides
```

---

## 🧪 Testing

```bash
# Backend tests
cd backend
go test ./...

# Frontend tests (if configured)
cd frontend
npm run test
```

### Manual Testing Checklist

- [ ] User signup and login
- [ ] Create event with images
- [ ] Edit and delete events
- [ ] Search and filter events
- [ ] Map view and location picker
- [ ] Bookmark and follow events
- [ ] Purchase tickets
- [ ] View dashboard
- [ ] Admin category management
- [ ] Mobile responsiveness

---

## 📖 Additional Documentation

- **Requirements Analysis:** `REQUIREMENTS_ANALYSIS.md`
- **Compliance Summary:** `COMPLIANCE_SUMMARY.md`
- **Phase Completion:** `PROJECT_COMPLETE.md`
- **Gap Fixes:** `GAPS_FIXED_SUMMARY.txt`

---

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code Style

- **Backend:** Follow Go conventions, use `gofmt`
- **Frontend:** Follow Vue 3 Composition API patterns
- **Commits:** Use conventional commits format

---

## 🐛 Troubleshooting

### Common Issues

**Hasura not connecting to database:**
```bash
# Check Docker services
docker-compose ps

# Restart services
docker-compose restart postgres hasura
```

**Frontend GraphQL errors:**
```bash
# Verify Hasura is running
curl http://localhost:8080/healthz

# Check environment variables
cat frontend/.env
```

**Image upload failing:**
- Verify Cloudinary credentials in `backend/.env`
- Check file size limits (max 5MB per image)

**Payment not working:**
- Verify Chapa API keys
- Check callback URL is accessible
- Review Chapa dashboard for errors

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Authors

- Temesgen A. - Initial work

---

## 🙏 Acknowledgments

- Hasura for the GraphQL engine
- Cloudinary for image hosting
- Chapa for payment processing
- OpenStreetMap for map tiles
- The Vue.js and Go communities

---

## 📞 Support

For support, email https://t.me/Temuab21 or open an issue in the repository.

---

**Made with ❤️ in Ethiopia**
