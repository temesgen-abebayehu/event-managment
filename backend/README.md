# Event Management Backend

Backend for Hasura Actions - handles auth, file upload, and payment.

## Structure
```
backend/
├── cmd/main.go                         # Server with dependency injection
├── config/config.go                    # Environment configuration
├── delivery/                           # HTTP handlers
│   ├── auth_handler.go                # Signup, Login
│   ├── upload_handler.go              # File upload
│   └── payment_handler.go             # Payment initiation & verification
├── usecase/                            # Business logic
│   ├── auth_usecase.go                # Password hashing, JWT generation
│   ├── upload_usecase.go              # File validation, Cloudinary upload
│   └── payment_usecase.go             # Payment flow, Chapa integration
├── repository/                         # Data access (only for actions)
│   ├── user_repository.go             # Users (auth)
│   ├── ticket_repository.go           # Tickets (payment)
│   └── order_repository.go            # Orders (payment)
├── infrastructure/                     # External services
│   ├── hasura/client.go               # Hasura GraphQL client
│   ├── cloudinary/client.go           # Cloudinary SDK
│   └── chapa/client.go                # Chapa payment API
├── domain/                             # Business entities
│   ├── user.go
│   ├── ticket.go
│   ├── upload.go
│   └── payment.go
└── utils/
    ├── jwt.go
    └── password.go
```

## Setup

```bash
cd backend
cp .env.example .env
# Edit .env with your credentials
go mod download
go run cmd/main.go
```

## Endpoints

### Auth Actions
- `POST /actions/signup` - Create user account
- `POST /actions/login` - Login and get JWT

### Upload Actions
- `POST /actions/upload` - Upload files to Cloudinary
- `POST /actions/delete-files` - Delete files from Cloudinary

### Payment Actions
- `POST /actions/initiate-payment` - Start payment flow with Chapa
- `POST /actions/verify-payment` - Verify payment status

### Webhooks
- `POST /webhook/chapa` - Chapa payment callback

## File Upload

**Naming:** `events/{user_id}/{event_id}/{timestamp}_{uuid}.jpg`

**Example:**
```
events/abc-123/evt-456/1704067200_a1b2c3d4.jpg
```

**Max:** 10 files, 5MB each

## Payment Flow

1. Frontend calls `initiate-payment` action
2. Backend creates pending order, calls Chapa API
3. Backend returns Chapa checkout URL
4. User completes payment on Chapa
5. Chapa sends webhook to `/webhook/chapa`
6. Backend updates order status to "completed"
7. Trigger automatically updates `quantity_sold` in tickets table

## Environment Variables

All required env vars are validated on startup. See `.env.example`.

## What Hasura Does Directly

Events, categories, bookmarks, follows, tags - all accessed by frontend via Hasura GraphQL. Backend only handles what requires custom logic (auth, upload, payment).
