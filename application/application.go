package application

import (
	"github.com/losinggeneration/lgge/input"
	"github.com/losinggeneration/lgge/timer"
	"github.com/losinggeneration/lgge/window"
)

type Application interface {
	Init() error
	Destroy() error
	IsRunning() bool
	Prepare(events []input.Event, t *timer.Timer)
	Render()
	SetWindow(w window.Window)
}
