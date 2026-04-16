package cafe

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func CallDivide() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

type LoginError struct {
	Username string
	Message  string
}

func (e *LoginError) Error() string {
	return fmt.Sprintf("Login error for user '%s': %s", e.Username, e.Message)
}

func login(username, password string) error {
	// Simulate checking username and password
	if username != "admin" || password != "password123" {
		return &LoginError{Username: username, Message: "invalid credentials"}
	}
	// Login successful
	return nil
}

func CallLogin() {
	// Attempt to login with wrong credentials
	err := login("user", "pass")
	if err != nil {
		switch e := err.(type) {
		case *LoginError:
			// Custom error handling
			fmt.Println("Custom error occurred:", e)
		default:
			// Other types of errors
			fmt.Println("Generic error occurred:", e)
		}
		return
	}

	// Continue with the rest of the program if login is successful
	fmt.Println("Login successful!")
}
