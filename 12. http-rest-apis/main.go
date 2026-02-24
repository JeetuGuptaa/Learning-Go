package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// User struct for JSON examples
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Response struct for consistent API responses
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// In-memory user storage (simulating a database)
var users = []User{
	{ID: 1, Name: "Alice Johnson", Email: "alice@example.com", CreatedAt: time.Now()},
	{ID: 2, Name: "Bob Smith", Email: "bob@example.com", CreatedAt: time.Now()},
	{ID: 3, Name: "Charlie Brown", Email: "charlie@example.com", CreatedAt: time.Now()},
}
var nextID = 4

// --- Handlers ---

// Simple home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to Go REST API</h1><p>Try the following endpoints:</p>")
	fmt.Fprintf(w, "<ul>")
	fmt.Fprintf(w, "<li>GET /api/users - Get all users</li>")
	fmt.Fprintf(w, "<li>GET /api/users/{id} - Get user by ID</li>")
	fmt.Fprintf(w, "<li>POST /api/users/create - Create new user</li>")
	fmt.Fprintf(w, "<li>DELETE /api/users/{id} - Delete user</li>")
	fmt.Fprintf(w, "</ul>")
}

// Get all users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendJSONResponse(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	sendJSONResponse(w, http.StatusOK, Response{
		Success: true,
		Data:    users,
	})
}

// Get user by ID
func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendJSONResponse(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	// Extract ID from URL path (simple parsing)
	idStr := r.URL.Path[len("/api/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendJSONResponse(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	// Find user
	for _, user := range users {
		if user.ID == id {
			sendJSONResponse(w, http.StatusOK, Response{
				Success: true,
				Data:    user,
			})
			return
		}
	}

	sendJSONResponse(w, http.StatusNotFound, Response{
		Success: false,
		Message: "User not found",
	})
}

// Create new user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendJSONResponse(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Error reading request body",
		})
		return
	}
	defer r.Body.Close()

	// Parse JSON
	var newUser User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		sendJSONResponse(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid JSON format",
		})
		return
	}

	// Validate
	if newUser.Name == "" || newUser.Email == "" {
		sendJSONResponse(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Name and email are required",
		})
		return
	}

	// Create user
	newUser.ID = nextID
	nextID++
	newUser.CreatedAt = time.Now()
	users = append(users, newUser)

	sendJSONResponse(w, http.StatusCreated, Response{
		Success: true,
		Message: "User created successfully",
		Data:    newUser,
	})
}

// Delete user
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		sendJSONResponse(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	// Extract ID from URL path
	idStr := r.URL.Path[len("/api/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendJSONResponse(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	// Find and delete user
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			sendJSONResponse(w, http.StatusOK, Response{
				Success: true,
				Message: "User deleted successfully",
			})
			return
		}
	}

	sendJSONResponse(w, http.StatusNotFound, Response{
		Success: false,
		Message: "User not found",
	})
}

// Helper function to send JSON responses
func sendJSONResponse(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// --- Middleware ---

// Logging middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("Completed in %v", time.Since(start))
	}
}

// CORS middleware (for browser access)
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// Chain middleware
func withMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return corsMiddleware(loggingMiddleware(handler))
}

// --- HTTP Client Example ---

// Example function to make HTTP GET request
func fetchUserExample() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Printf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return
	}

	fmt.Println("\n--- HTTP Client Example ---")
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response: %s\n", string(body))
}

func main() {
	// Register routes
	http.HandleFunc("/", withMiddleware(homeHandler))
	http.HandleFunc("/api/users", withMiddleware(getUsersHandler))
	http.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		// Route to appropriate handler based on method and path
		if r.URL.Path == "/api/users/" || r.URL.Path == "/api/users" {
			withMiddleware(getUsersHandler)(w, r)
		} else if r.Method == http.MethodGet {
			withMiddleware(getUserByIDHandler)(w, r)
		} else if r.Method == http.MethodDelete {
			withMiddleware(deleteUserHandler)(w, r)
		} else {
			sendJSONResponse(w, http.StatusMethodNotAllowed, Response{
				Success: false,
				Message: "Method not allowed",
			})
		}
	})

	// Create endpoint for POST requests
	http.HandleFunc("/api/users/create", withMiddleware(createUserHandler))

	// Demonstrate HTTP client
	go func() {
		time.Sleep(1 * time.Second)
		fetchUserExample()
	}()

	// Start server
	port := ":8080"
	fmt.Printf("\n🚀 Server starting on http://localhost%s\n", port)
	fmt.Println("📝 API Endpoints:")
	fmt.Println("   GET    http://localhost:8080/api/users")
	fmt.Println("   GET    http://localhost:8080/api/users/1")
	fmt.Println("   POST   http://localhost:8080/api/users/create")
	fmt.Println("   DELETE http://localhost:8080/api/users/1")
	fmt.Println("\n💡 Try it with curl:")
	fmt.Println("   curl http://localhost:8080/api/users")
	fmt.Println(`   curl -X POST http://localhost:8080/api/users/create -H "Content-Type: application/json" -d '{"name":"Jane Doe","email":"jane@example.com"}'`)
	fmt.Println()

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}