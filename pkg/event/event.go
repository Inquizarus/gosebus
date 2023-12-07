package event

type Event interface {
	// Name returns the events name which is matched against listener patterns
	Name() string
	// Data returns whatever was sent along the event to be distributed
	Data() interface{}
	StringMap() map[string]string
	StringSlice() []string
	String() string
	Bytes() []byte
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

func (e *standardEvent) Bytes() []byte {
	return e.Data().([]byte)
}

func (e *standardEvent) String() string {
	return e.Data().(string)
}

func (e *standardEvent) StringMap() map[string]string {
	return e.Data().(map[string]string)
}

// StringSlice returns the string slice representation of the data.
//
// It does not modify the original event.
// It returns a slice of strings.
func (e *standardEvent) StringSlice() []string {
	return e.Data().([]string)
}

// NewEvent creates a new Event with the specified name and data.
//
// Parameters:
// - name: the name of the event.
// - data: the data associated with the event.
//
// Return:
// - Event: the newly created Event.
func NewEvent(name string, data interface{}) Event {
	return &standardEvent{
		name: name,
		data: data,
	}
}
