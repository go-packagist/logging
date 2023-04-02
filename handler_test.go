package monolog

import (
	"github.com/go-packagist/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerable(t *testing.T) {
	h := NewHandlerable(
		WithLevel(logger.Notice),
	)

	assert.True(t, h.IsHandling(&Record{Level: logger.Emergency}))
	assert.True(t, h.IsHandling(&Record{Level: logger.Alert}))
	assert.True(t, h.IsHandling(&Record{Level: logger.Critical}))
	assert.True(t, h.IsHandling(&Record{Level: logger.Error}))
	assert.True(t, h.IsHandling(&Record{Level: logger.Warning}))
	assert.True(t, h.IsHandling(&Record{Level: logger.Notice}))
	assert.False(t, h.IsHandling(&Record{Level: logger.Info}))
	assert.False(t, h.IsHandling(&Record{Level: logger.Debug}))

	assert.False(t, h.Handle(&Record{}))
	assert.False(t, h.HandleBatch([]*Record{
		{
			Level: logger.Debug,
		},
	}))
}
