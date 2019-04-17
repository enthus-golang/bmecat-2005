package bmecat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatetime_UnmarshalText(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		var d Datetime
		err := d.UnmarshalText([]byte(`2005-03-27T08:10:30+01:00`))
		assert.NoError(t, err)
		assert.Equal(t, "2005-03-27T08:10:30+01:00", d.Format(time.RFC3339))
	})
	t.Run("full_utc", func(t *testing.T) {
		var d Datetime
		err := d.UnmarshalText([]byte(`2005-03-27T08:10:30+01:00`))
		assert.NoError(t, err)
		assert.Equal(t, "2005-03-27T08:10:30+01:00", d.Format(time.RFC3339))
	})
	t.Run("y", func(t *testing.T) {
		var d Datetime
		err := d.UnmarshalText([]byte(`2005`))
		assert.NoError(t, err)
		assert.Equal(t, "2005-01-01T00:00:00Z", d.Format(time.RFC3339))
	})
	t.Run("ym", func(t *testing.T) {
		var d Datetime
		err := d.UnmarshalText([]byte(`2005-03`))
		assert.NoError(t, err)
		assert.Equal(t, "2005-03-01T00:00:00Z", d.Format(time.RFC3339))
	})
	t.Run("ymd", func(t *testing.T) {
		var d Datetime
		err := d.UnmarshalText([]byte(`2005-03-12`))
		assert.NoError(t, err)
		assert.Equal(t, "2005-03-12T00:00:00Z", d.Format(time.RFC3339))
	})
}
