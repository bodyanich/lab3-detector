package processor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseImageMetadata(t *testing.T) {
	result, err := ParseImageMetadata("image_data_7_timestamp_123456")

	require.NoError(t, err)
	require.Equal(t, "7", result.WorkerID)
	require.Equal(t, "123456", result.Timestamp)
}

func TestParseImageMetadataInvalidInput(t *testing.T) {
	_, err := ParseImageMetadata("invalid-input")

	require.Error(t, err)
}
