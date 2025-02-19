package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogstreamer(t *testing.T) {
	// Given
	testCases := []struct {
		stdType string
		name    string
	}{
		{stdType: StdOut, name: "test-stdout"},
		{stdType: StdErr, name: "test-stderr"},
	}

	for _, testCase := range testCases {
		// When
		streamer := NewLogstreamer(testCase.stdType, testCase.name)

		// Then
		assert.IsType(t, new(Logstreamer), streamer)

		assert.Equal(t, testCase.stdType, streamer.stdType)
		assert.Equal(t, testCase.name, streamer.name)
	}
}
