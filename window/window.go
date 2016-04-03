package window

import (
	"errors"
	"reflect"
)

var ErrUnregisteredWindow = errors.New("not a registered window type")

type WindowConfig struct {
	Title      string
	Width      int
	Height     int
	Fullscreen bool

	Running bool
}

type Window interface {
	Create(WindowConfig) error
	Destroy() error

	IsRunning() bool
	ProcessEvents()
	SwapBuffers()
}

var registeredWindows = make(map[string]Window)

func RegisterWindow(name string, implementation Window) {
	registeredWindows[name] = implementation
}

func GetWindow(name string) (Window, error) {
	w := registeredWindows[name]
	if w == nil {
		return nil, ErrUnregisteredWindow
	}

	// Get the value of w,  which is a pointer, then get the underlying element's type
	t := reflect.ValueOf(w).Elem().Type()

	// Create a new instance based of w's type
	return reflect.New(t).Interface().(Window), nil
}
