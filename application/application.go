package application

import "github.com/losinggeneration/lgge/input"

type Application interface {
	Init() error
	Destroy() error
	Prepare(events []input.Event, dt float64)
	Render()
}
