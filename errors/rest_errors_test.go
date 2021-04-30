package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewRestError(t *testing.T) {
	expectedMessage := "This is test error"
	e := "error test"
	err := NewRestError(expectedMessage, http.StatusForbidden, e, nil)

	if err == nil {
		t.Error("Error should not be nil")
	}
	if err != nil && err.Status() != 403 {
		t.Errorf("Error should be 403")
	}
	if err != nil && err.Message() != expectedMessage {
		t.Errorf("\nExpected message: %s\nReceived message: %s", expectedMessage, err.Message())
	}
	if err != nil && err.Error() != "error test" {
		t.Errorf("\nExpected error: %s\nReceived error: %s", "error test", err.Error())
	}
	if err != nil && err.Causes() != nil {
		t.Errorf("Causes should be nil")
	}

}

func TestNewInternalServerError(t *testing.T) {
	expectedMessage := "This is internal server error"
	e := errors.New("db may be failing test error")
	err := NewInternalServerError(expectedMessage, e)

	if err == nil {
		t.Error("Error should not be nil")
	}
	if err != nil && err.Status() != 500 {
		t.Errorf("Error should be 500")
	}
	if err != nil && err.Message() != expectedMessage {
		t.Errorf("\nExpected message: %s\nReceived message: %s", expectedMessage, err.Message())
	}
	if err != nil && err.Error() != "Internal server error" {
		t.Errorf("\nExpected error: %s\nReceived error: %s", "Internal server error", err.Error())
	}
	if err != nil && len(err.Causes()) != 1 {
		t.Errorf("Expected a non-empty slice")
	}
}

func TestNewBadRequestError(t *testing.T) {
	expectedMessage := "User Bad request test error"
	err := NewBadRequestError(expectedMessage)

	if err == nil {
		t.Error("Error should not be nil")
	}
	if err != nil && err.Status() != 400 {
		t.Error("Error should be 400")
	}
	if err != nil && err.Message() != expectedMessage {
		t.Errorf("\nExpected message: %s\nReceived message: %s", expectedMessage, err.Message())
	}
	if err != nil && err.Error() != "Bad Request" {
		t.Errorf("\nExpected error: %s\nReceived error: %s", "Bad Request", err.Error())
	}
}

func TestNewNotFoundError(t *testing.T) {
	expectedMessage := "User not found test error"
	err := NewNotFoundError(expectedMessage)

	if err == nil {
		t.Error("Error should not be nil")
	}
	if err != nil && err.Status() != 404 {
		t.Error("Error should be 404")
	}
	if err != nil && err.Message() != expectedMessage {
		t.Errorf("\nExpected message: %s\nReceived message: %s", expectedMessage, err.Message())
	}
	if err != nil && err.Error() != "Not found" {
		t.Errorf("\nExpected error: %s\nReceived error: %s", "Not found", err.Error())
	}
}

func TestNewUnauthorizedError(t *testing.T) {
	expectedMessage := "you have been unauthorized"
	err := NewUnauthorizedError(expectedMessage)

	if err == nil {
		t.Error("Error should not be nil")
	}

	if err != nil && err.Status() != 401 {
		t.Error("Error should be 401")
	}

	if err != nil && err.Message() != expectedMessage {
		t.Errorf("\nExpected message: %s\nReceived message: %s", expectedMessage, err.Message())
	}
	if err != nil && err.Error() != "Unauthorized" {
		t.Errorf("\nExpected error: %s\nReceived error: %s", "Not found", err.Error())
	}
}
