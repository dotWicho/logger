package logger

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStandardLogger_Info(t *testing.T) {

	t.Run("buffer is empty if we have a Test Logger", func(t *testing.T) {

		// We define some vars
		buffer := &bytes.Buffer{}

		// Try to create Logger
		_logger := NewLogger(true)
		_logger.SetLevel(logrus.InfoLevel)

		// Logger is not nil
		assert.NotNil(t, _logger)

		// Some lines to _logger
		_logger.Info("This line appears on console? %v", true)
		_logger.Info("Another line with number %d", 22)

		// Simple validation
		assert.Empty(t, buffer.String())
	})

	t.Run("buffer have data if have a non Test Logger", func(t *testing.T) {

		// We define some vars
		buffer := &bytes.Buffer{}

		// Try to create Logger
		_logger := NewLogger(false)
		_logger.SetOutput(buffer)
		_logger.SetLevel(logrus.InfoLevel)

		// Logger is not nil
		assert.NotNil(t, _logger)

		// Some lines to _logger
		_logger.Info("This line will never show up on console, right? %v", true)
		_logger.Info("Another line with number %d", 44)

		// Simple validation
		assert.NotEmpty(t, buffer.String())
	})
}

func TestStandardLogger_Warn(t *testing.T) {

	t.Run("buffer is empty if we have a Test Logger", func(t *testing.T) {

		// We define some vars
		buffer := &bytes.Buffer{}

		// Try to create Logger
		_logger := NewLogger(true)
		_logger.SetLevel(logrus.WarnLevel)

		// Logger is not nil
		assert.NotNil(t, _logger)

		// Some lines to _logger
		_logger.Warn("This line appears on console? %v", true)
		_logger.Warn("Another line with number %d", 22)

		// Simple validation
		assert.Empty(t, buffer.String())
	})

	t.Run("buffer have data if have a non Test Logger", func(t *testing.T) {

		// We define some vars
		buffer := &bytes.Buffer{}

		// Try to create Logger
		_logger := NewLogger(false)
		_logger.SetOutput(buffer)
		_logger.SetLevel(logrus.WarnLevel)

		// Logger is not nil
		assert.NotNil(t, _logger)

		// Some lines to _logger
		_logger.Warn("This line will never show up on console, right? %v", true)
		_logger.Warn("Another line with number %d", 44)

		// Simple validation
		assert.NotEmpty(t, buffer.String())
	})
}

func TestStandardLogger_Error(t *testing.T) {

	t.Run("buffer is empty if we have a Test Logger", func(t *testing.T) {

		// We define some vars
		buffer := &bytes.Buffer{}

		// Try to create Logger
		_logger := NewLogger(true)
		_logger.SetLevel(logrus.ErrorLevel)

		// Logger is not nil
		assert.NotNil(t, _logger)

		// Some lines to _logger
		_logger.Error("This line appears on console? %v", true)
		_logger.Error("Another line with number %d", 22)

		// Simple validation
		assert.Empty(t, buffer.String())
	})

	t.Run("buffer have data if have a non Test Logger", func(t *testing.T) {

		// We define some vars
		buffer := &bytes.Buffer{}

		// Try to create Logger
		_logger := NewLogger(false)
		_logger.SetOutput(buffer)
		_logger.SetLevel(logrus.ErrorLevel)

		// Logger is not nil
		assert.NotNil(t, _logger)

		// Some lines to _logger
		_logger.Error("This line will never show up on console, right? %v", true)
		_logger.Error("Another line with number %d", 44)

		// Simple validation
		assert.NotEmpty(t, buffer.String())
	})
}

func TestStandardLogger_Debug(t *testing.T) {

	t.Run("buffer is empty if we have a Test Logger", func(t *testing.T) {

		// We define some vars
		buffer := &bytes.Buffer{}

		// Try to create Logger
		_logger := NewLogger(true)
		_logger.SetLevel(logrus.DebugLevel)

		// Logger is not nil
		assert.NotNil(t, _logger)

		// Some lines to _logger
		_logger.Debug("This line appears on console? %v", true)
		_logger.Debug("Another line with number %d", 22)

		// Simple validation
		assert.Empty(t, buffer.String())
	})

	t.Run("buffer have data if have a non Test Logger", func(t *testing.T) {

		// We define some vars
		buffer := &bytes.Buffer{}

		// Try to create Logger
		_logger := NewLogger(false)
		_logger.SetOutput(buffer)
		_logger.SetLevel(logrus.DebugLevel)

		// Logger is not nil
		assert.NotNil(t, _logger)

		// Some lines to _logger
		_logger.Debug("This line will never show up on console, right? %v", true)
		_logger.Debug("Another line with number %d", 44)

		// Simple validation
		assert.NotEmpty(t, buffer.String())
	})
}
