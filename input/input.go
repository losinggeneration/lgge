package input

import "errors"

type Key int
type Mod uint16
type Button int
type Type int

const (
	EventQuit = Type(iota)
	EventKeyDown
	EventKeyUp
	EventMouse
	EventWindow
)

type Event struct {
	Type  Type
	Event interface{}
}

type QuitEvent bool

type KeyEvent struct {
	Key    Key
	Mod    Mod
	Repeat bool
}

type KeyDownEvent struct {
	KeyEvent
}

type KeyUpEvent struct {
	KeyEvent
}

type MouseEvent struct {
	Which  uint   // Which mouse
	Button Button // The mouse button pressed
	State  bool   // true for pressed, false for released
	X, Y   int    // X & Y possition
	Wheel  int    // Wheel delta
}

type WindowMovement struct {
	X, Y int
}

type WindowSize struct {
	W, H int
}

type WindowEvent struct {
	Shown       bool
	Hidden      bool
	Exposed     bool
	Minimized   bool
	Maximized   bool
	Restored    bool
	Enter       bool
	Leave       bool
	FocusGained bool
	FocusLost   bool
	Close       bool

	Moved       WindowMovement
	Resized     WindowSize
	SizeChanged WindowSize
}

type Process func() []Event

var registered = make(map[string]Process)

var ErrUnregisteredInput = errors.New("not a registered input type")

func Register(name string, process Process) {
	registered[name] = process
}

func GetProcess(name string) (Process, error) {
	if registered[name] == nil {
		return nil, ErrUnregisteredInput
	}

	return registered[name], nil
}
