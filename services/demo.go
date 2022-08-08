package services

import "net/http"

// Register a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {}

// Login controller for user auths
func LoginUser(w http.ResponseWriter, r *http.Request) {}

// Get All user list
func UserList(w http.ResponseWriter, r *http.Request) {}

// Logout an user
func LogoutUser(w http.ResponseWriter, r *http.Request) {}
