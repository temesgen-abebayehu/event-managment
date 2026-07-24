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

EventHub Ethiopia is a comprehensive event management platform that allows users to discover, create, manage, and purchase tickets for events across Ethiopia. The platform features a modern UI, real-time search, location-based filtering, integrated payment processing, and cloud-based image management.

**Key Highlights:**
- 🔐 JWT-based authentication with secure password hashing
- 🗺️ Interactive map-based location selection and event browsing
- 🏷️ Tag-based event categorization and full-text search
- 💳 Integrated payment gateway for ticket purchases
- 📸 Multi-image upload with cloud storage
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
- **Database:** PostgreSQL 15+ (supports Neon serverless)
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

### Infrastructure
- **Database:** PostgreSQL 15+ / Neon (serverless)
- **GraphQL Engine:** Hasura (self-hosted or cloud)
- **Containerization:** Docker & Docker Compose (optional)

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

- **Docker & Docker Compose** (v20.10+) - For containerized setup
- **Node.js** (v18+) - For frontend development
- **Go** (v1.22+) - For backend development
- **PostgreSQL** (v15+) - For database (or use Neon serverless)
- **Git**

### External Services (Required)
- **Cloudinary Account** - For image storage ([Sign up](https://cloudinary.com/))
- **Chapa Account** - For payment processing ([Sign up](https://chapa.co/))

---

## 🌐 Demo

A live demo of the application is available for testing purposes.

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

Create `backend/.env`:

```env
# Database
DATABASE_URL=postgresql://user:password@host:port/dbname

# Cloudinary (Image Storage)
CLOUDINARY_CLOUD_NAME=your-cloud-name
CLOUDINARY_API_KEY=your-api-key
CLOUDINARY_API_SECRET=your-api-secret

# Chapa (Payment Gateway)
CHAPA_SECRET_KEY=your-chapa-secret-key
CHAPA_PUBLIC_KEY=your-chapa-public-key
CHAPA_CALLBACK_URL=http://localhost:3001/webhook/chapa
CHAPA_RETURN_URL=http://localhost:3000/payment/success

# JWT Authentication
JWT_SECRET=your-super-secret-jwt-key-min-32-characters-long

# Hasura
HASURA_ADMIN_SECRET=your-hasura-admin-secret
HASURA_ENDPOINT=http://localhost:8080/v1/graphql

# Server
PORT=3001
```

### Frontend Environment Variables

Create `frontend/.env` (local development):

```env
NUXT_PUBLIC_BACKEND_URL=http://localhost:3001
NUXT_PUBLIC_GRAPHQL_ENDPOINT=http://localhost:8080/v1/graphql
```

For production, create `frontend/.env.production`:

```env
NUXT_PUBLIC_BACKEND_URL=https://your-backend-domain.com
NUXT_PUBLIC_GRAPHQL_ENDPOINT=https://your-hasura-domain.com/v1/graphql
```

---

## 🚀 Quick Start (Local Development)

### Using Docker (Recommended for Local Dev)

```bash
# 1. Setup environment
cp .env.example .env
# Edit .env with your Cloudinary and Chapa credentials

# 2. Start everything
docker-compose up -d

# 3. Run migrations
# Connect to your PostgreSQL and run migrations 001-012
```

**Access:**
- Frontend: http://localhost:3000
- Backend: http://localhost:3001
- Hasura: http://localhost:8080/console
- PostgreSQL: localhost:5432

### Manual Setup (Without Docker)

**Backend:**
```bash
cd backend
cp .env.example .env
# Edit .env with your credentials
go mod download
go run cmd/main.go
```

**Frontend:**
```bash
cd frontend
cp .env.example .env
# Edit .env
npm install
npm run dev
```

**Database:**
- Install PostgreSQL locally or use Neon (serverless)
- Run all migrations (001-012)
- Seed categories if needed

---

## 🗄️ Database Migrations

### Applying Migrations

Run all migrations in order (001-012):

```bash
cd backend/migrations

# Connect to your PostgreSQL database
psql "your-connection-string" -f 001_create_users.sql
psql "your-connection-string" -f 002_create_categories.sql
psql "your-connection-string" -f 003_create_events.sql
psql "your-connection-string" -f 004_create_event_images.sql
psql "your-connection-string" -f 005_create_tags.sql
psql "your-connection-string" -f 006_create_bookmarks_follows.sql
psql "your-connection-string" -f 007_create_tickets.sql
psql "your-connection-string" -f 008_create_functions.sql
psql "your-connection-string" -f 009_create_triggers.sql
psql "your-connection-string" -f 010_add_computed_and_generated_fields.sql
psql "your-connection-string" -f 011_add_hasura_computed_fields.sql
psql "your-connection-string" -f 012_fix_search_vector_conflict.sql

# Or apply all at once
cat *.sql | psql "your-connection-string"
```

> **Important:** Migration 012 is required for tag-based search functionality.

### Database Schema Overview

**Core Tables:**
- `users` - User accounts with JWT authentication
- `categories` - Event categories (dynamic CRUD via Hasura)
- `events` - Event information with slug-based URLs
- `event_images` - Multiple images per event (up to 5, stored in Cloudinary)
- `tags` - Reusable tags for categorization
- `event_tags` - Many-to-many relationship between events and tags
- `bookmarks` - User bookmarks for events
- `follows` - User follows for events
- `tickets` - Ticket templates with pricing and availability
- `orders` - Ticket purchases via Chapa payment gateway

**PostgreSQL Features:**
- ✅ Full-text search with `search_vector` (includes tags via function)
- ✅ Slug generation for SEO-friendly URLs
- ✅ Triggers for auto-updating timestamps
- ✅ Functions for complex queries (search, nearby events)
- ✅ Proper indexes on FK and search columns
- ✅ Materialized view for optimized tag-based search

**Important Notes:**
- `search_vector` is a GENERATED column (cannot be manually updated)
- Tag search is handled via `get_event_tags_tsvector()` function
- Use `search_events_with_tags()` function for full search including tags

---

## 📚 API Documentation

### Hasura Actions (Backend Endpoints)

The backend provides these custom actions via Hasura:

**Authentication**
- `signup` - User registration with email/password
- `login` - User authentication with JWT token generation

**File Management**
- `upload` - Upload multiple images (max 5, 5MB each)
- `delete-files` - Delete images from cloud storage

**Payment**
- `initiate-payment` - Initialize payment gateway checkout
- `verify-payment` - Verify payment status and create order

### GraphQL API

All data operations use standard Hasura GraphQL queries and mutations:

**Events**
- Query events with filtering, sorting, pagination
- Create, update, delete events
- Search events by title, description, and tags
- Filter by category, date, location

**Tags**
- Create tags with automatic deduplication
- Link tags to events with conflict handling

**Orders & Tickets**
- Query user tickets and orders
- Track ticket availability
- View order history

Refer to Hasura GraphQL schema for complete API documentation.

---

## 🚢 Deployment

### Deployment Options

The application can be deployed using various platforms:

**Recommended Stack:**
- **Frontend:** Vercel, Netlify, or similar serverless platforms
- **Backend:** Render, Railway, Heroku, or any Go-compatible hosting
- **Database:** Neon (serverless PostgreSQL), AWS RDS, or self-hosted PostgreSQL
- **Hasura:** Hasura Cloud or self-hosted

### General Deployment Steps

1. **Database Setup**
   - Create PostgreSQL instance (cloud or self-hosted)
   - Run all migrations (001-012)
   - Seed initial data (categories)

2. **Hasura Setup**
   - Deploy Hasura instance (cloud or self-hosted)
   - Connect to database
   - Set `HASURA_GRAPHQL_UNAUTHORIZED_ROLE=anonymous`
   - Track all tables and relationships
   - Configure permissions for `anonymous` and `user` roles
   - Add computed fields: `tickets_remaining`, `is_sold_out`, `event_status`, `attendee_count`
   - Configure Hasura Actions pointing to backend endpoints

3. **Backend Deployment**
   - Build: `go build -o main cmd/main.go`
   - Configure all environment variables
   - Ensure CORS is properly configured
   - Deploy binary or use native Go runtime

4. **Frontend Deployment**
   - Build: `npm run build`
   - Set production environment variables
   - Deploy build output

### Hasura Actions Configuration

Configure these actions in Hasura Console pointing to your backend:

```yaml
- name: signup
  endpoint: https://your-backend/actions/signup
  
- name: login
  endpoint: https://your-backend/actions/login
  
- name: initiate_payment
  endpoint: https://your-backend/actions/initiate-payment
  
- name: verify_payment
  endpoint: https://your-backend/actions/verify-payment
```

### Production Checklist

- [ ] All 12 database migrations applied
- [ ] Environment variables configured for all services
- [ ] Hasura connected to database
- [ ] Hasura permissions configured
- [ ] Backend endpoints accessible
- [ ] Frontend environment variables set
- [ ] CORS properly configured
- [ ] SSL/HTTPS enabled
- [ ] Payment gateway keys (production mode)
- [ ] Image storage configured
- [ ] Database backups enabled

---

## 🧪 Testing

### Manual Testing Checklist

**Authentication:**
- [x] User signup with email validation
- [x] User login with JWT token
- [x] Protected routes redirect to login
- [x] Token persistence across sessions

**Event Management:**
- [x] Create event with multiple images
- [x] Create event with tags
- [x] Edit event and update images
- [x] Delete event images
- [x] Set featured image
- [x] Slug-based URLs work correctly
- [x] Events display on homepage
- [x] Event detail page shows all information

**Search & Filter:**
- [x] Search by title and description
- [x] Search includes tags (via migration 012)
- [x] Filter by category
- [x] Filter by date range
- [x] Filter by price
- [x] Map view displays events
- [x] Location-based filtering

**Payment Flow:**
- [x] Ticket purchase initiates Chapa payment
- [x] Redirect to Chapa checkout
- [x] Return to success page after payment
- [x] Payment verification works (not localhost)
- [x] Tickets appear in dashboard
- [x] Order history is accurate

**UI/UX:**
- [x] Mobile responsive design
- [x] Loading states for all async operations
- [x] Error messages are user-friendly
- [x] Image previews work correctly
- [x] Map picker functions properly

### Automated Testing (Future)

```bash
# Backend tests
cd backend
go test ./...

# Frontend tests
cd frontend
npm run test
```

---

## 📖 Additional Documentation

- **Requirements Analysis:** `REQUIREMENTS_ANALYSIS.md`
- **Compliance Summary:** `COMPLIANCE_SUMMARY.md`
- **Phase Completion:** `PROJECT_COMPLETE.md`
- **Gap Fixes:** `GAPS_FIXED_SUMMARY.txt`

---

## 🔧 Recent Updates & Fixes

### Key Improvements

**Database & Search**
- Fixed conflict between GENERATED columns and triggers (Migration 012)
- Tag-based search now works correctly
- Optimized search performance with materialized views

**Code Quality**
- Removed debug console.log statements
- Improved error handling with user-friendly messages
- Cleaned up unused code and handlers

**Features**
- Implemented SEO-friendly slug-based URLs for events
- Fixed payment verification to use environment variables
- Added graceful error handling for optional operations

### Architecture Decisions

**Why GENERATED Columns?**
The `search_vector` column uses PostgreSQL's GENERATED ALWAYS constraint to automatically index event titles and descriptions. Tags are indexed separately via a function to avoid conflicts.

**Why Separate Tag Search?**
Tags are dynamic and can change frequently. Using a separate function (`get_event_tags_tsvector()`) combined with a materialized view provides flexibility and performance.

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

**Database Connection Issues:**
```bash
# Verify database is accessible
psql "your-connection-string"

# Check connection string format
# Format: postgresql://user:password@host:port/database
```

**GraphQL/Hasura Errors:**
```bash
# Verify Hasura is accessible
curl https://your-hasura-url/healthz

# Check Hasura permissions
# Ensure 'anonymous' and 'user' roles are configured
```

**Image Upload Failing:**
- Verify Cloudinary credentials in environment variables
- Check file size limits (max 5MB per image, max 5 images per event)
- Ensure `event_id` is included in upload request
- Verify CORS is enabled on backend

**Payment Issues:**
- Verify payment gateway API keys
- Check callback URL is publicly accessible
- Ensure return URL points to correct frontend domain
- Verify environment variables use correct domains (not localhost in production)

**Tag Creation Errors:**
- Ensure migration 012 is applied
- Verify `on_conflict` clause in tag mutations
- Check Hasura permissions for `event_tags` table

**Slug-Based URLs Not Working:**
- Ensure migrations 010 and 011 are applied
- Verify slugs are auto-generated on event creation
- Check routing uses `slug` field instead of `id`

### Local Development Issues

**Backend Won't Start:**
```bash
# Verify Go version
go version  # Should be 1.22+

# Check dependencies
go mod tidy

# Verify environment variables
cat backend/.env
```

**Frontend Build Errors:**
```bash
# Clear cache
rm -rf node_modules .nuxt

# Reinstall dependencies
npm install

# Check Node version
node --version  # Should be 18+
```

**Docker Issues:**
```bash
# Check Docker services
docker-compose ps

# Restart services
docker-compose restart

# View logs
docker-compose logs -f
```

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

For issues, questions, or contributions:
- Open an issue in the repository
- Contact: https://t.me/Temuab21

---

**Made with ❤️ in Ethiopia**
