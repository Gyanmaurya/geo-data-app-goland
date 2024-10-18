package controllers

import (
    "encoding/json"
    "net/http"
    "golang.org/x/crypto/bcrypt"
)

var users = make(map[string]string) // In-memory storage for demo

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// Register a new user
func Register(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil || user.Username == "" || user.Password == "" {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Check if the user already exists
    if _, exists := users[user.Username]; exists {
        http.Error(w, "User already exists", http.StatusConflict)
        return
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Could not register user", http.StatusInternalServerError)
        return
    }

    users[user.Username] = string(hashedPassword) // Store the username and hashed password
    w.WriteHeader(http.StatusCreated)
}

import (
	"github.com/dgrijalva/jwt-go"
	"time"
 )
 
 // Login a user and generate a JWT token
 func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Password == "" {
	    http.Error(w, "Invalid input", http.StatusBadRequest)
	    return
	}
 
	// Check if user exists and verify password
	storedPassword, exists := users[user.Username]
	if !exists || bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)) != nil {
	    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	    return
	}
 
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	    "username": user.Username,
	    "exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
 
	tokenString, err := token.SignedString([]byte("your_secret_key")) // Use your own secret key
	if err != nil {
	    http.Error(w, "Could not generate token", http.StatusInternalServerError)
	    return
	}
 
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
 }
 