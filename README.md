# REST-API (Go)

This project is a simple RESTful API built with **Go** and the **Gin** framework. It demonstrates how to use standard HTTP methods to manage event data and user authentication.

## Features

- Built with **Go** and **Gin Web Framework**
- **JWT Authentication**: Secure login and signup functions
- **SQLite Database**: Persistent data storage for users and events
- **CRUD Operations**: Create, Read, Update, and Delete resources
- Middleware for protected routes

## Main Functions

- **User Management**: Signup and login with hashed passwords
- **Event Management**: Manage event details like name, location, and time
- **Token Verification**: Only authorized users can modify data

## Installation

1. **Clone the project:**

   ```bash
   git clone [https://github.com/tanset/REST-API.git](https://github.com/tanset/REST-API.git)
   cd REST-API
   ```

2. **Download dependencies:**

   ```bash
   go mod tidy
   ```

3. **Start the server:**

   ```bash
   go run .
   ```
   *The server will run at `http://localhost:8080`.*

## API Endpoints

| Method | URL            | Description           | Auth Required |
| ------ | -------------- | --------------------- | ------------- |
| POST   | `/signup`      | Register a new user   | No            |
| POST   | `/login`       | Login and get token   | No            |
| GET    | `/events`      | List all events       | No            |
| GET    | `/events/:id`  | Get a single event    | No            |
| POST   | `/events`      | Create a new event    | **Yes**       |
| PUT    | `/events/:id`  | Update an event       | **Yes**       |
| DELETE | `/events/:id`  | Delete an event       | **Yes**       |


## License

MIT License
