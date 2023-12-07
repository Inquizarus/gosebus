package gosebus_test

import (
	"testing"
	"time"

	"github.com/inquizarus/gosebus"
	"github.com/inquizarus/gosebus/pkg/event"
	"github.com/inquizarus/gosebus/pkg/handling"
)

func TestThatgosebusWorks(t *testing.T) {
	b := gosebus.New()
	done := make(chan bool)

	b.On("test_event", func(_ event.Event) {
		done <- true
	})

	b.Publish(event.NewEvent("test_event", nil))

	select {
	case <-done:
		// Test passed
	case <-time.After(time.Millisecond * 50):
		t.Errorf("timed out waiting for event handler")
	}
}

func TestThatgosebusWildcardWorks(t *testing.T) {
	b := gosebus.New()
	done := make(chan bool)
	b.On("*_test_*", func(e event.Event) {
		done <- true
	})
	b.Publish(event.NewEvent("trigger_test_event", nil))

	select {
	case <-done:
	// Test passed
	case <-time.After(time.Millisecond * 50):
		t.Errorf("timed out waiting for event handler")
	}
}

func TestHandlerAppliedOnlyOnce(t *testing.T) {
	b := gosebus.New()
	counter := 0
	handler := handling.NewStandardEventHandler(func(e event.Event) {
		counter++
	}, handling.HandlerOptionWithPattern("test_event"), handling.HandlerOptionRunOnce())

	b.Handle(handler)
	b.Publish(event.NewEvent("test_event", nil))
	b.Publish(event.NewEvent("test_event", nil))

	if counter != 1 {
		t.Errorf("handler ran more than once")
	}
}
