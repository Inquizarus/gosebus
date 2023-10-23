package gosebus_test

import (
	"testing"
	"time"

	"github.com/inquizarus/gosebus"
	"github.com/stretchr/testify/assert"
)

func TestThatgosebusWorks(t *testing.T) {
	b := gosebus.New()
	count := 0
	b.On("test_event", func(e gosebus.Event) {
		count++
	})
	b.Publish(gosebus.NewEvent("test_event", nil))
	time.Sleep(time.Microsecond * 50) // Allow for some time as handling of events occurs in goroutines
	assert.Equal(t, 1, count)
}

func TestThatgosebusWildcardWorks(t *testing.T) {
	b := gosebus.New()
	count := 0
	b.On("*_test_*", func(e gosebus.Event) {
		count++
	})
	b.Publish(gosebus.NewEvent("trigger_test_event", nil))
	time.Sleep(time.Microsecond * 10) // Allow for some time as handling of events occurs in goroutines
	assert.Equal(t, 1, count)
}
