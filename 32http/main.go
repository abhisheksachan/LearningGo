package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// ============================================================
// SERVER SIDE
// ============================================================

// Handler for GET /hello — returns plain text
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // /hello?name=Alice
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

// Handler for GET /users — returns JSON
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler for POST /echo — reads JSON body, echoes it back
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// ============================================================
// CLIENT SIDE
// ============================================================

// doRequest makes a request, reads the body, and closes the connection
// immediately — before the next request starts. This returns the connection
// to the pool right away instead of waiting until the caller function exits.
func doRequest(client *http.Client, method, url, payload string) (int, []byte, error) {
	var bodyReader io.Reader
	if payload != "" {
		bodyReader = strings.NewReader(payload)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return 0, nil, err
	}
	if payload != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close() // scoped here — connection freed when doRequest returns

	body, err := io.ReadAll(resp.Body)
	return resp.StatusCode, body, err
}

func runClient(base string) {
	client := &http.Client{Timeout: 5 * time.Second}

	// GET plain text
	status, body, err := doRequest(client, "GET", base+"/hello?name=Alice", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GET /hello → %d\n%s\n", status, body)

	// GET JSON
	status, body, err = doRequest(client, "GET", base+"/users", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GET /users → %d\n%s\n", status, body)

	// POST with JSON body
	status, body, err = doRequest(client, "POST", base+"/echo", `{"message":"hello"}`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("POST /echo → %d\n%s\n", status, body)
}

// ============================================================
// MAIN
// ============================================================

func main() {
	const addr = "localhost:8080"

	// Register routes
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/echo", echoHandler)

	// Start server in background, run client, then exit
	go func() {
		log.Printf("Server listening on %s", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(50 * time.Millisecond) // wait for server to start
	runClient("http://" + addr)
}