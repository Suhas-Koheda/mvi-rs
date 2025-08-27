# Movie Booking API

A RESTful API for managing movie bookings, built with Go. This project provides endpoints for user authentication, movie management, hall and seat management, reservations, and showtimes. It follows a modular structure for scalability and maintainability.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Setup & Installation](#setup--installation)
- [Environment Variables](#environment-variables)
- [Database Migrations](#database-migrations)
- [API Endpoints](#api-endpoints)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- User registration, login, and authentication (OAuth support)
- CRUD operations for movies, halls, seats, and showtimes
- Reservation system for booking seats
- Environment variable management
- Modular repository and controller pattern
- Database migrations

---

## Project Structure

```
api/
  main.go                  # Entry point for the API server
internal/
  controllers/             # HTTP handlers and server initialization
    AuthController.go
    InitialiseServer.go
    MovieController.go
  hash/                    # Password hashing utilities
    hash.go
  migrations/              # Database migration scripts
    main.go
  models/                  # Database models
    Hall.go
    Movie.go
    Reservation.go
    Seat.go
    ShowTime.go
    User.go
  oauth/                   # OAuth authentication logic
    oauth.go
pkg/
  env/                     # Environment variable loader
    load_env.go
  initialisers/            # Database initialisation
    InitialiseDatabase.go
repository/
  MovieRepository.Go       # Movie data access logic
  UserRepository.go        # User data access logic
go.mod, go.sum             # Go module files
readme.md                  # Project documentation
```

---

## Setup & Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/Suhas-Koheda/mvi-rs.git
   cd mvi-rs
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Configure environment variables:**
   - Copy `.env.example` to `.env` and update values as needed.

4. **Run database migrations:**
   ```sh
   go run internal/migrations/main.go
   ```

5. **Start the API server:**
   ```sh
   go run api/main.go
   ```

---

## Environment Variables

Environment variables are loaded via `pkg/env/load_env.go`. Typical variables include:

- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` (Database config)
- `JWT_SECRET` (Authentication)
- `OAUTH_CLIENT_ID`, `OAUTH_CLIENT_SECRET` (OAuth)

---

## Database Migrations

Migration scripts are located in `internal/migrations/main.go`. Run migrations before starting the server to set up tables for movies, users, reservations, etc.

---

## API Endpoints

### Authentication

- `POST /register` — Register a new user
- `POST /login` — Login and receive JWT
- `GET /oauth/callback` — OAuth login

### Movies

- `GET /movies` — List all movies
- `POST /movies` — Add a new movie
- `GET /movies/{id}` — Get movie details
- `PUT /movies/{id}` — Update movie
- `DELETE /movies/{id}` — Delete movie

### Halls & Seats

- `GET /halls` — List halls
- `POST /halls` — Add hall
- `GET /seats` — List seats
- `POST /seats` — Add seat

### ShowTimes

- `GET /showtimes` — List showtimes
- `POST /showtimes` — Add showtime

### Reservations

- `POST /reservations` — Book a seat
- `GET /reservations` — List reservations

---

## Usage

- Use tools like Postman or curl to interact with the API.
- Authenticate using JWT for protected endpoints.
- Admin endpoints require elevated privileges.

---

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/foo`)
3. Commit your changes (`git commit -am 'Add feature'`)
4. Push to the branch (`git push origin feature/foo`)
5. Create a new Pull Request

---

## License

This project is licensed under the MIT License.

---
