package main

import (
	"fmt"
	"log"
)

// User represents a user entity with ID and First name.
type User struct {
	ID    int
	First string
}

// Datastore defines methods for user data storage.
type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(user User) error
}

// MockDatastore is a mock implementation of the Datastore interface for testing.
type MockDatastore struct {
	Users map[int]User
}

// GetUser retrieves a user by ID from the datastore.
func (md *MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}

// SaveUser saves a user to the datastore.
func (md *MockDatastore) SaveUser(user User) error {
	_, ok := md.Users[user.ID]
	if ok {
		return fmt.Errorf("User %v already exists", user.ID)
	}
	md.Users[user.ID] = user
	return nil
}

// Service provides methods to interact with the datastore.
type Service struct {
	ds Datastore
}

// GetUser retrieves a user by ID from the datastore.
func (s *Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

// SaveUser saves a user to the datastore.
func (s *Service) SaveUser(user User) error {
	return s.ds.SaveUser(user)
}

func main() {
	db := &MockDatastore{
		Users: make(map[int]User),
	}

	service := &Service{ds: db}

	user1 := User{
		ID:    1,
		First: "Alyapany",
	}

	// Try saving the user
	err := service.SaveUser(user1)
	if err != nil {
		fmt.Println("Error saving user:", err)
	}

	// Retrieve the same user
	user1Returned, err := service.GetUser(user1.ID)
	if err != nil {
		log.Fatalf("Error getting user: %v\n", err)
	}

	fmt.Println(user1)
	fmt.Println(user1Returned)
}
