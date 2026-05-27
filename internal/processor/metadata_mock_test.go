package processor

import (
	"testing"

	"lab3-detector/internal/storage"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockMetadataStore struct {
	mock.Mock
}

func (m *mockMetadataStore) Save(metadata storage.Metadata) error {
	args := m.Called(metadata)

	return args.Error(0)
}

func TestProcessAndStoreMetadata(t *testing.T) {
	store := new(mockMetadataStore)

	expected := storage.Metadata{
		ImageID:  "image-001",
		WorkerID: 7,
		Valid:    true,
	}

	store.On("Save", expected).Return(nil).Once()

	err := ProcessAndStoreMetadata("image-001", 7, store)

	require.NoError(t, err)
	store.AssertExpectations(t)
}

func TestProcessAndStoreMetadataSkipsInvalidImageID(t *testing.T) {
	store := new(mockMetadataStore)

	err := ProcessAndStoreMetadata("", 7, store)

	require.NoError(t, err)
	store.AssertNotCalled(t, "Save", mock.Anything)
}
