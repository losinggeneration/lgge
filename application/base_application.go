package application

import (
	"github.com/losinggeneration/lgge/input"
	"github.com/losinggeneration/lgge/window"
)

type BaseApplication struct {
	Running bool
	Window  window.Window
}

func (b *BaseApplication) Init() error                   { b.Running = true; return nil }
func (b BaseApplication) IsRunning() bool                { return b.Running }
func (b BaseApplication) Destroy() error                 { return nil }
func (b BaseApplication) Prepare([]input.Event, float64) {}
func (b BaseApplication) Render()                        {}
func (b *BaseApplication) SetWindow(w window.Window)     { b.Window = w }
