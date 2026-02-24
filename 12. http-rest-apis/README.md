# HTTP/REST APIs in Go

This example demonstrates how to build RESTful APIs using Go's standard `net/http` package.

## Concepts Covered

### HTTP Server Basics
- Creating an HTTP server with `http.ListenAndServe()`
- Registering route handlers with `http.HandleFunc()`
- Writing responses with `http.ResponseWriter`
- Reading requests with `http.Request`

### RESTful Endpoints
- **GET** - Retrieve resources
- **POST** - Create new resources
- **DELETE** - Delete resources
- **PUT/PATCH** - Update resources (not shown but similar to POST)

### JSON Handling
- Marshaling Go structs to JSON with `json.Marshal()`
- Unmarshaling JSON to Go structs with `json.Unmarshal()`
- Using struct tags for JSON field mapping
- Streaming JSON with `json.Encoder` and `json.Decoder`

### Headers and Status Codes
- Setting Content-Type headers
- Returning appropriate HTTP status codes (200, 201, 400, 404, 500, etc.)
- CORS headers for browser access

### Middleware
- Logging requests
- CORS handling
- Chaining middleware functions
- Request/response processing

### HTTP Client
- Making GET requests with `http.Get()`
- Making POST requests with `http.Post()`
- Custom requests with `http.NewRequest()`
- Reading response bodies

## API Endpoints

### GET /api/users
Returns all users.

```bash
curl http://localhost:8080/api/users
```

### GET /api/users/{id}
Returns a specific user by ID.

```bash
curl http://localhost:8080/api/users/1
```

### POST /api/users/create
Creates a new user.

```bash
curl -X POST http://localhost:8080/api/users/create \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","email":"jane@example.com"}'
```

### DELETE /api/users/{id}
Deletes a user by ID.

```bash
curl -X DELETE http://localhost:8080/api/users/1
```

## Running the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Testing with curl

### Get all users
```bash
curl http://localhost:8080/api/users
```

### Get specific user
```bash
curl http://localhost:8080/api/users/2
```

### Create new user
```bash
curl -X POST http://localhost:8080/api/users/create \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Smith",
    "email": "john@example.com"
  }'
```

### Delete user
```bash
curl -X DELETE http://localhost:8080/api/users/3
```

## Testing with HTTPie (alternative to curl)

If you have HTTPie installed:

```bash
# Get all users
http GET localhost:8080/api/users

# Create user
http POST localhost:8080/api/users/create name="Jane" email="jane@test.com"

# Delete user
http DELETE localhost:8080/api/users/1
```

## Key Concepts

### Response Format
All API responses follow this structure:
```json
{
  "success": true,
  "message": "Optional message",
  "data": {} // Actual data
}
```

### HTTP Status Codes
- **200 OK** - Successful GET/DELETE
- **201 Created** - Successful POST
- **400 Bad Request** - Invalid input
- **404 Not Found** - Resource not found
- **405 Method Not Allowed** - Wrong HTTP method
- **500 Internal Server Error** - Server error

### Best Practices Shown
1. **Consistent response format** - All responses follow the same structure
2. **Proper error handling** - Returns appropriate status codes and messages
3. **Input validation** - Checks for required fields
4. **Clean separation** - Handlers, middleware, and helpers are separate
5. **Middleware pattern** - Reusable request/response processing
6. **CORS support** - Allows browser-based clients

## Next Steps

To improve this API, consider:
- Using a router like `gorilla/mux` or `chi` for better routing
- Adding authentication/authorization (JWT, OAuth)
- Connecting to a real database (PostgreSQL, MySQL)
- Adding input validation library
- Implementing pagination for list endpoints
- Adding request rate limiting
- Using environment variables for configuration
- Writing tests for handlers
- Adding OpenAPI/Swagger documentation
- Implementing graceful shutdown

## Popular Go Web Frameworks

While this example uses the standard library, you might also explore:
- **Gin** - Fast, minimalist framework
- **Echo** - High performance, extensible
- **Fiber** - Express-inspired, built on Fasthttp
- **Chi** - Lightweight, idiomatic router
- **Gorilla Mux** - Powerful URL router
