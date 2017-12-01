package test

import (
	"testing"

	"github.com/gostores/assert"
	"github.com/gostores/logging"
)

func TestAllHooks(t *testing.T) {

	assert := assert.New(t)

	logger, hook := NewNullLogger()
	assert.Nil(hook.LastEntry())
	assert.Equal(0, len(hook.Entries))

	logger.Error("Hello error")
	assert.Equal(logging.ErrorLevel, hook.LastEntry().Level)
	assert.Equal("Hello error", hook.LastEntry().Message)
	assert.Equal(1, len(hook.Entries))

	logger.Warn("Hello warning")
	assert.Equal(logging.WarnLevel, hook.LastEntry().Level)
	assert.Equal("Hello warning", hook.LastEntry().Message)
	assert.Equal(2, len(hook.Entries))

	hook.Reset()
	assert.Nil(hook.LastEntry())
	assert.Equal(0, len(hook.Entries))

	hook = NewGlobal()

	logging.Error("Hello error")
	assert.Equal(logging.ErrorLevel, hook.LastEntry().Level)
	assert.Equal("Hello error", hook.LastEntry().Message)
	assert.Equal(1, len(hook.Entries))

}
