package errors

import (
	"encoding/json"
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

func TestNewRestErrorFromBytes(t *testing.T) {
	t.Run("Shouldn't have error", func(t *testing.T) {
		message := "This is test err"
		e := "new error"
		errTest := restErr{
			M: message,
			S: 401,
			E: e,
			C: nil,
		}
		errBytes, _ := json.Marshal(errTest)
		rErr, err := NewRestErrorFromBytes(errBytes)

		if err != nil {
			t.Error("Error should be nil")
		}

		if rErr == nil {
			t.Error("RestError should not be nil")
		} else {
			if rErr.Message() != message {
				t.Errorf("\nExpected message: %s\nReceived message: %s", message, rErr.Message())
			}
			if rErr.Status() != 401 {
				t.Errorf("\nExpected status: %d\nReceived status: %d", 401, rErr.Status())
			}
			if rErr.Causes() != nil {
				t.Errorf("\nExpected causes: %v\nReceived causes: %v", nil, rErr.Causes())
			}
			if rErr.Error() != e {
				t.Errorf("\nExpected error: %s\nReceived error: %s", e, rErr.Error())
			}
		}
	})
	t.Run("Should have error", func(t *testing.T) {
		errTest := "{This is not a valid json}"
		errBytes := []byte(errTest)

		rErr, err := NewRestErrorFromBytes(errBytes)
		invalidRes := "invalid json"

		if rErr != nil {
			t.Error("RestErr should be nil")
		}
		if err == nil {
			t.Error("Error should not be nil")
		} else {
			if err.Error() != invalidRes {
				t.Errorf("\nExpected error: %s\nReceived error: %s", invalidRes, err.Error())
			}
		}
	})

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
