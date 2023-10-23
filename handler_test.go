package gosebus_test

import (
	"testing"

	"github.com/inquizarus/gosebus"
	"github.com/stretchr/testify/assert"
)

func TestThatHandlerMatchesCorrectly(t *testing.T) {
	cases := []struct {
		eventHandler gosebus.Handler
		event        gosebus.Event
		expected     bool
	}{
		{
			eventHandler: gosebus.NewStandardEventHandler("correct_test_event", func(e gosebus.Event) {}),
			event:        gosebus.NewEvent("correct_test_event", nil),
			expected:     true,
		},
		{
			eventHandler: gosebus.NewStandardEventHandler("*_test_event", func(e gosebus.Event) {}),
			event:        gosebus.NewEvent("correct_test_event", nil),
			expected:     true,
		},
		{
			eventHandler: gosebus.NewStandardEventHandler("correct_test_*", func(e gosebus.Event) {}),
			event:        gosebus.NewEvent("correct_test_event", nil),
			expected:     true,
		},
		{
			eventHandler: gosebus.NewStandardEventHandler("correct_*_event", func(e gosebus.Event) {}),
			event:        gosebus.NewEvent("correct_test_event", nil),
			expected:     true,
		},
		{
			eventHandler: gosebus.NewStandardEventHandler("corr*ect_*_event", func(e gosebus.Event) {}),
			event:        gosebus.NewEvent("correct_test_event", nil),
			expected:     true,
		},
		{
			eventHandler: gosebus.NewStandardEventHandler("corr*ect_*_eve*nt*", func(e gosebus.Event) {}),
			event:        gosebus.NewEvent("x", nil),
			expected:     false,
		},
		{
			eventHandler: gosebus.NewStandardEventHandler("correct_test_event", func(e gosebus.Event) {}),
			event:        gosebus.NewEvent("wrong_test_event", nil),
			expected:     false,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, c.eventHandler.Match(c.event))
	}
}

func TestThatHandlerHandlesEventCorrectly(t *testing.T) {
	i := 0
	e := gosebus.NewEvent("test_event", 10)
	eh := gosebus.NewStandardEventHandler("test_event", func(e gosebus.Event) {
		i += e.Data().(int)
	})
	eh.Handle(e)
	assert.Equal(t, 10, i)
}
