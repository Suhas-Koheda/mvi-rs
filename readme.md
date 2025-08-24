# Movie Reservation System Backend

A robust backend API for a Movie Reservation System, built with Go and MySQL. This project supports user authentication, movie and showtime management, seat reservations, and admin reporting.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Setup & Installation](#setup--installation)
- [Database Schema](#database-schema)
- [API Endpoints](#api-endpoints)
- [Authentication & Authorization](#authentication--authorization)
- [Testing](#testing)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- User registration and login (JWT-based authentication)
- Admin and user roles
- CRUD operations for movies, halls, showtimes (admin only)
- Seat reservation with concurrency-safe logic
- Reporting endpoints for admins
- Input validation and error handling

## Tech Stack

- **Language:** Go
- **Web Framework:** Gin (or Echo)
- **Database:** MySQL
- **ORM/Driver:** go-sql-driver/mysql
- **Auth:** JWT (github.com/golang-jwt/jwt/v5)
- **Env Management:** godotenv
- **Password Hashing:** bcrypt

## Project Structure

```
/cmd
  /api
	 main.go          # Application entry point
/internal
  /handlers          # HTTP request handlers (Controllers)
  /models            # Struct definitions matching DB tables
  /repository        # Database interaction logic (Data Access Layer)
  /auth              # JWT token creation/validation logic
  /middleware        # Custom middleware (e.g., JWT auth, Admin check)
/pkg
  /database          # Database connection setup
  /utils             # Helper functions (e.g., password hashing)
.env
go.mod
go.sum
schema.sql
```

## Setup & Installation

1. **Clone the repository:**
	```sh
	git clone https://github.com/suhas-koheda/mvi-rs.git
	cd mvi-rs
	```

2. **Install dependencies:**
	```sh
	go mod tidy
	```

3. **Set up MySQL:**
	- Install MySQL and create a database:
	  ```sql
	  CREATE DATABASE movie_db;
	  ```
	- Update your `.env` file with DB credentials.

4. **Run migrations:**
	- Use the provided `schema.sql` to set up tables.

5. **Start the server:**
	```sh
	go run main.go
	```

## Database Schema

See [`schema.sql`](schema.sql) for full table definitions. Key tables:

- `users`: Stores user info and hashed passwords.
- `movies`: Movie details.
- `halls`: Theater halls and capacities.
- `showtimes`: Movie showtimes, linked to halls and movies.
- `reservations`: User reservations for showtimes.
- `seats`: (Optional) Tracks individual seat reservations.

## API Endpoints

### Auth

- `POST /api/signup` — Register a new user
- `POST /api/login` — Login and receive JWT

### Movies

- `GET /api/movies` — List all movies
- `POST /api/movies` — Add a movie (admin only)
- `PUT /api/movies/:id` — Update a movie (admin only)
- `DELETE /api/movies/:id` — Delete a movie (admin only)

### Showtimes

- `GET /api/movies/:id/showtimes?date=YYYY-MM-DD` — Get showtimes for a movie on a date
- `POST /api/showtimes` — Add showtime (admin only, with conflict check)

### Reservations

- `GET /api/showtimes/:id/seats` — Check seat availability
- `POST /api/reservations` — Make a reservation (auth required)
- `GET /api/user/reservations` — List user's reservations

### Admin Reports

- `GET /api/admin/reports/reservations?start_date=&end_date=` — Reservation reports (admin only)

## Authentication & Authorization

- JWT tokens are required for protected routes.
- Admin-only routes require the user to have the `admin` role.
- Use the `Authorization: Bearer <token>` header.

## Testing

- Use [Postman](https://www.postman.com/) or `curl` to test endpoints.
- Simulate concurrent reservations to ensure no overbooking.
- All errors are returned as JSON with appropriate HTTP status codes.

## Environment Variables

Create a `.env` file with:

```
DB_USER=your_db_user
DB_PASS=your_db_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=movie_db
JWT_SECRET=your_jwt_secret
ADMIN_EMAIL=admin@example.com
ADMIN_PASSWORD=adminpassword
```

## Contributing

Pull requests are welcome! Please open issues for suggestions or bugs.

## License

[MIT](LICENSE)

---

**Note:** For full endpoint details, request/response examples, and advanced usage, see the [API Documentation](#) or comments in the code.

---
