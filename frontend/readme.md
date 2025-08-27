# Movie Booking Frontend

This is a simple React-based frontend for the Movie Booking API. It provides user authentication, movie browsing, reservations, and more, interacting with the backend API.

## Prerequisites

- Node.js (v18 or above recommended)
- npm or yarn

## Setup & Installation

1. **Navigate to the frontend directory:**
   ```sh
   cd frontend
   ```

2. **Install dependencies:**
   ```sh
   npm install
   # or
   yarn install
   ```

3. **Start the development server:**
   ```sh
   npm start
   # or
   yarn start
   ```
   The app will open at `http://localhost:8080` by default.

## Project Structure

```
frontend/
  components/        # React components (pages, navbar, etc.)
  public/            # Static files (index.html)
  App.jsx            # Main app component
  index.js           # Entry point
  package.json       # Project config
  webpack.config.js  # Webpack config
  .babelrc           # Babel config
```

## API Proxy

API requests are proxied to the backend (default: `http://localhost:8080`). Ensure your backend server is running before using the frontend.

## Features

- User login and registration
- Movie listing and details
- View reservations
- Simple navigation

## Build for Production

```sh
npm run build
# or
yarn build
```
The production build will be in the `dist/` folder.

## Notes
- This frontend assumes the backend API endpoints as described in the main project README.
- For custom backend URLs, update the proxy settings in `webpack.config.js`.

---
