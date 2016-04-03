package input

import (
	"github.com/losinggeneration/lgge/input"

	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	input.Register("sdl", processEvents)
}

func processEvents() []input.Event {
	events := make([]input.Event, 0)
	var windowEvent input.WindowEvent
	var hadWindowEvent bool
	mouseEvents := make(map[uint32]input.MouseEvent)

	event := sdl.PollEvent()
	for event != nil {
		switch e := event.(type) {
		case *sdl.KeyDownEvent:
			k := input.KeyDownEvent{
				input.KeyEvent{
					Key:    keys[e.Keysym.Sym],
					Mod:    mods[e.Keysym.Mod],
					Repeat: e.Repeat > 0,
				},
			}
			events = append(events, input.Event{Type: input.EventKeyDown, Event: k})

		case *sdl.KeyUpEvent:
			k := input.KeyUpEvent{
				input.KeyEvent{
					Key:    keys[e.Keysym.Sym],
					Mod:    mods[e.Keysym.Mod],
					Repeat: e.Repeat > 0,
				},
			}
			events = append(events, input.Event{Type: input.EventKeyUp, Event: k})

		case *sdl.MouseMotionEvent:
			m := mouseEvents[e.Which]
			m.X = int(e.X)
			m.Y = int(e.Y)
			mouseEvents[e.Which] = m

		case *sdl.MouseWheelEvent:
			m := mouseEvents[e.Which]
			m.Wheel = int(e.Y)
			mouseEvents[e.Which] = m

		case *sdl.MouseButtonEvent:
			m := mouseEvents[e.Which]
			m.Button = buttons[e.Button]
			m.State = e.State > 0
			mouseEvents[e.Which] = m

		case *sdl.WindowEvent:
			hadWindowEvent = true
			switch e.Event {
			case sdl.WINDOWEVENT_SHOWN:
				windowEvent.Shown = true
			case sdl.WINDOWEVENT_HIDDEN:
				windowEvent.Hidden = true
			case sdl.WINDOWEVENT_EXPOSED:
				windowEvent.Exposed = true
			case sdl.WINDOWEVENT_MINIMIZED:
				windowEvent.Minimized = true
			case sdl.WINDOWEVENT_MAXIMIZED:
				windowEvent.Maximized = true
			case sdl.WINDOWEVENT_RESTORED:
				windowEvent.Restored = true
			case sdl.WINDOWEVENT_ENTER:
				windowEvent.Enter = true
			case sdl.WINDOWEVENT_LEAVE:
				windowEvent.Leave = true
			case sdl.WINDOWEVENT_FOCUS_GAINED:
				windowEvent.FocusGained = true
			case sdl.WINDOWEVENT_FOCUS_LOST:
				windowEvent.FocusLost = true
			case sdl.WINDOWEVENT_CLOSE:
				windowEvent.Close = true

			case sdl.WINDOWEVENT_MOVED:
				windowEvent.Moved = input.WindowMovement{X: int(e.Data1), Y: int(e.Data2)}
			case sdl.WINDOWEVENT_RESIZED:
				windowEvent.Resized = input.WindowSize{W: int(e.Data1), H: int(e.Data2)}
			case sdl.WINDOWEVENT_SIZE_CHANGED:
				windowEvent.SizeChanged = input.WindowSize{W: int(e.Data1), H: int(e.Data2)}
			default:
				hadWindowEvent = false
			}

		case *sdl.QuitEvent:
			q := input.QuitEvent(true)
			events = append(events, input.Event{Type: input.EventQuit, Event: q})
		}

		event = sdl.PollEvent()
	}

	if hadWindowEvent {
		events = append(events, input.Event{Type: input.EventWindow, Event: windowEvent})
	}

	for w, m := range mouseEvents {
		m.Which = uint(w)
		events = append(events, input.Event{Type: input.EventMouse, Event: m})
	}

	return events
}
