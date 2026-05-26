package logging

import "testing"

func TestNewDevelopmentLogger(t *testing.T) {
	logger, err := NewDevelopmentLogger()
	if err != nil {
		t.Fatalf("expected logger without error, got %v", err)
	}

	if logger == nil {
		t.Fatal("expected logger instance, got nil")
	}
}

func TestNewNopLogger(t *testing.T) {
	logger := NewNopLogger()
	if logger == nil {
		t.Fatal("expected nop logger instance, got nil")
	}
}
