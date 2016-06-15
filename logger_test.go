package flamingo_test

import (
	"testing"

	"github.com/hasanozgan/flamingo"
	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestLoggerFactoryWhenEmptyFields(t *testing.T) {
	loggerFactory := flamingo.NewLoggerFactory()
	assert.Equal(t, len(loggerFactory.GetFields()), 0)
}

func TestLoggerFactoryForNewLoggerWithName(t *testing.T) {
	loggerFactory := flamingo.NewLoggerFactory()
	logger := loggerFactory.NewLogger("Test Logger")
	assert.Equal(t, logger.Name, "Test Logger")
}
