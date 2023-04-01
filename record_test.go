package monolog

import (
	"github.com/go-packagist/logger"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRecord_Marshal(t *testing.T) {
	ti, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	r := &Record{
		Channel: "test",
		Message: "test",
		Level:   logger.Debug,
		Time:    ti,
		Extra:   nil,
	}

	marshal, err := r.Marshal()

	assert.Nil(t, err)
	assert.Equal(t, `{"channel":"test","message":"test","level":"debug","time":"2020-01-01T00:00:00Z"}`, string(marshal))
}

func TestRecord_Unmarshal(t *testing.T) {
	r := `{"channel":"test","message":"test","level":"debug","time":"2020-01-01T00:00:00Z"}`

	var record Record
	err := record.Unmarshal([]byte(r))

	assert.Nil(t, err)
	assert.Equal(t, "test", record.Channel)
	assert.Equal(t, "2020-01-01T00:00:00Z", record.Time.Format(time.RFC3339))
	assert.Equal(t, "test", record.Message)
	assert.Equal(t, logger.Debug, record.Level)
	assert.Nil(t, record.Extra)
}
