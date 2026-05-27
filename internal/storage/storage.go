// Package storage contains abstractions for saving processed image metadata.
package storage

// Metadata describes extracted image metadata.
type Metadata struct {
	ImageID  string
	WorkerID int
	Valid    bool
}

// MetadataStore defines storage behavior for processed metadata.
type MetadataStore interface {
	Save(metadata Metadata) error
}
