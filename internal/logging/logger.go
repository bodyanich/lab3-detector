// Package logging provides application logger constructors.
package logging

import "go.uber.org/zap"

// NewDevelopmentLogger creates a structured logger for local development.
func NewDevelopmentLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

// NewNopLogger creates a logger for tests.
func NewNopLogger() *zap.Logger {
	return zap.NewNop()
}
