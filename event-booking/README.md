# Event Booking Project

## Overview
The Event Booking project is a Go application designed to manage events and user registrations. It demonstrates the use of Go for building RESTful APIs, middleware, and utility functions.

## Project Structure
```
event-booking/
├── go.mod                # Go module file
├── main.go               # Entry point of the application
├── db/
│   └── db.go            # Database connection and operations
├── middlewares/
│   └── auth.go          # Authentication middleware
├── models/
│   ├── event.go         # Event data model
│   └── user.go          # User data model
├── routes/
│   ├── events.go        # Routes for event operations
│   ├── registrations.go # Routes for user registrations
│   ├── routes.go        # Main route configuration
│   └── users.go         # Routes for user operations
├── utils/
    ├── hash.go          # Utility for hashing
    └── jwt.go           # Utility for JWT handling
```

## Features
- **Event Management**: Create, update, delete, and retrieve events.
- **User Management**: Register, authenticate, and manage users.
- **Middleware**: Authentication middleware for securing routes.
- **Utilities**: Hashing and JWT utilities for secure operations.

## Getting Started

### Prerequisites
- Go 1.20 or later

### Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```bash
   cd event-booking
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Application
To start the application, run:
```bash
go run main.go
```

### Project Configuration
Database and other configurations are managed in `db/db.go`. Update this file to configure database connections.

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments
- Inspired by Go's best practices for building RESTful APIs.
- Special thanks to the Go community for their excellent resources and support.
