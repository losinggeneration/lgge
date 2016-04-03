package application

import (
	"github.com/losinggeneration/lgge/input"
	"github.com/losinggeneration/lgge/window"
)

type Application interface {
	Init() error
	Destroy() error
	IsRunning() bool
	Prepare(events []input.Event, dt float64)
	Render()
	SetWindow(w window.Window)
}
