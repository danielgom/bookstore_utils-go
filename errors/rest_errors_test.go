package errors

import "testing"

func TestNewInternalServerError(t *testing.T) {
	expectedMessage := "This is internal server error"
	err := NewInternalServerError(expectedMessage)

	if err == nil {
		t.Error("Error should not be nil")
	}
	if err != nil && err.Status != 500 {
		t.Errorf("Error should be 500")
	}
	if err != nil && err.Message != expectedMessage {
		t.Errorf("\nExpected message: %s\nReceived message: %s", expectedMessage, err.Message)
	}
	if err != nil && err.Error != "Internal server error" {
		t.Errorf("\nExpected error: %s\nReceived error: %s", "Internal server error", err.Error)
	}
}

func TestNewBadRequestError(t *testing.T) {
	expectedMessage := "User Bad request test error"
	err := NewBadRequestError(expectedMessage)

	if err == nil {
		t.Error("Error should not be nil")
	}
	if err != nil && err.Status != 400 {
		t.Error("Error should be 400")
	}
	if err != nil && err.Message != expectedMessage {
		t.Errorf("\nExpected message: %s\nReceived message: %s", expectedMessage, err.Message)
	}
	if err != nil && err.Error != "Bad Request" {
		t.Errorf("\nExpected error: %s\nReceived error: %s", "Bad Request", err.Error)
	}
}

func TestNewNotFoundError(t *testing.T) {
	expectedMessage := "User not found test error"
	err := NewNotFoundError(expectedMessage)

	if err == nil {
		t.Error("Error should not be nil")
	}
	if err != nil && err.Status != 404 {
		t.Error("Error should be 404")
	}
	if err != nil && err.Message != expectedMessage {
		t.Errorf("\nExpected message: %s\nReceived message: %s", expectedMessage, err.Message)
	}
	if err != nil && err.Error != "Not found" {
		t.Errorf("\nExpected error: %s\nReceived error: %s", "Not found", err.Error)
	}
}
