package processor

import "testing"

func FuzzParseImageMetadata(f *testing.F) {
	f.Add("image_data_1_timestamp_123456789")
	f.Add("image_data_0_timestamp_0")
	f.Add("invalid-input")
	f.Add("")
	f.Add("image_data_x_timestamp_y")

	f.Fuzz(func(_ *testing.T, input string) {
		_, _ = ParseImageMetadata(input)
	})
}
