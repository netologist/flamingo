package vodka_test

import (
	"testing"

	"github.com/hasanozgan/vodka"
	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestLoggerFactoryWhenEmptyFields(t *testing.T) {
	loggerFactory := vodka.NewLoggerFactory()
	assert.Equal(t, len(loggerFactory.GetFields()), 0)
}

func TestLoggerFactoryForNewLoggerWithName(t *testing.T) {
	loggerFactory := vodka.NewLoggerFactory()
	logger := loggerFactory.NewLogger("Test Logger")
	assert.Equal(t, logger.Name, "Test Logger")
}
