package event_test

import (
	"testing"

	"github.com/inquizarus/gosebus/pkg/event"
	"github.com/stretchr/testify/assert"
)

func TestEventTypes(t *testing.T) {
	t.Run("TestStringEvent", func(t *testing.T) {
		event := event.NewEvent("test_event", "test_data")
		assert.Equal(t, "test_event", event.Name())
		assert.Equal(t, "test_data", event.Data())
		assert.Equal(t, "test_data", event.String())
	})

	t.Run("TestBytesEvent", func(t *testing.T) {
		event := event.NewEvent("test_event", []byte("test_data"))
		assert.Equal(t, "test_event", event.Name())
		assert.Equal(t, []byte("test_data"), event.Bytes())
	})

	t.Run("TestStringSliceEvent", func(t *testing.T) {
		event := event.NewEvent("test_event", []string{"test_data1", "test_data2"})
		assert.Equal(t, "test_event", event.Name())
		assert.Equal(t, []string{"test_data1", "test_data2"}, event.StringSlice())
	})

	t.Run("TestStringMapEvent", func(t *testing.T) {
		event := event.NewEvent("test_event", map[string]string{"key": "value"})
		assert.Equal(t, "test_event", event.Name())
		assert.Equal(t, map[string]string{"key": "value"}, event.StringMap())
	})
}
