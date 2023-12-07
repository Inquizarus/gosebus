package handler_test

import (
	"testing"

	"github.com/inquizarus/gosebus/pkg/event"
	"github.com/inquizarus/gosebus/pkg/handler"
	"github.com/stretchr/testify/assert"
)

func TestThatHandlerMatchesCorrectly(t *testing.T) {
	cases := []struct {
		eventHandler handler.Handler
		event        event.Event
		expected     bool
	}{
		{
			eventHandler: handler.New(func(e event.Event) {}, handler.WithPattern("correct_test_event")),
			event:        event.NewEvent("correct_test_event", nil),
			expected:     true,
		},
		{
			eventHandler: handler.New(func(e event.Event) {}, handler.WithPattern("*_test_event")),
			event:        event.NewEvent("correct_test_event", nil),
			expected:     true,
		},
		{
			eventHandler: handler.New(func(e event.Event) {}, handler.WithPattern("correct_test_*")),
			event:        event.NewEvent("correct_test_event", nil),
			expected:     true,
		},
		{
			eventHandler: handler.New(func(e event.Event) {}, handler.WithPattern("correct_*_event")),
			event:        event.NewEvent("correct_test_event", nil),
			expected:     true,
		},
		{
			eventHandler: handler.New(func(e event.Event) {}, handler.WithPattern("corr*ect_*_event")),
			event:        event.NewEvent("correct_test_event", nil),
			expected:     true,
		},
		{
			eventHandler: handler.New(func(e event.Event) {}, handler.WithPattern("corr*ect_*_eve*nt*")),
			event:        event.NewEvent("x", nil),
			expected:     false,
		},
		{
			eventHandler: handler.New(func(e event.Event) {}, handler.WithPattern("correct_test_event")),
			event:        event.NewEvent("wrong_test_event", nil),
			expected:     false,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, c.eventHandler.Match(c.event))
	}
}

func TestThatHandlerHandlesEventCorrectly(t *testing.T) {
	i := 0
	e := event.NewEvent("test_event", 10)
	eh := handler.New(func(e event.Event) {
		i += e.Data().(int)
	}, handler.WithPattern("test_event"))
	eh.Handle(e)
	assert.Equal(t, 10, i)
}

func TestRunOnceReturnsCorrectValue(t *testing.T) {
	// Test case 1: Assert that the ShouldRunOnce() method returns true
	assert.True(t, handler.New(
		func(e event.Event) {},
		handler.WithPattern("test_event"),
		handler.RunOnce(),
	).ShouldRunOnce())

	// Test case 2: Assert that the ShouldRunOnce() method returns false
	assert.False(t, handler.New(
		func(e event.Event) {},
		handler.WithPattern("test_event"),
	).ShouldRunOnce())
}

func TestHandlerEquals(t *testing.T) {
	handler1 := handler.New(
		func(e event.Event) {},
		handler.WithPattern("test_event"),
	)
	handler2 := handler.New(
		func(e event.Event) {},
		handler.WithPattern("test_event"),
	)
	handler3 := handler.New(
		func(e event.Event) {},
		handler.WithPattern("different_test_event"),
	)

	// Test case 1: Assert that Equals returns true for the same handler
	assert.True(t, handler1.Equals(handler1))

	// Test case 2: Assert that Equals returns false for different handlers with same pattern
	assert.False(t, handler1.Equals(handler2))

	// Test case 3: Assert that Equals returns false for handlers with different patterns
	assert.False(t, handler1.Equals(handler3))
}
