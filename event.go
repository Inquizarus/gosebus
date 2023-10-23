package gosebus

type Event interface {
	// Name returns the events name which is matched against listener patterns
	Name() string
	// Data returns whatever was sent along the event to be distributed
	Data() interface{}
}

type standardEvent struct {
	name string
	data interface{}
}

func (e *standardEvent) Name() string {
	return e.name
}

func (e *standardEvent) Data() interface{} {
	return e.data
}

func NewEvent(name string, data interface{}) Event {
	return &standardEvent{
		name,
		data,
	}
}
