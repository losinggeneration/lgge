package application

import "github.com/losinggeneration/lgge/input"

type BaseApplication struct{}

func (e BaseApplication) Init() error                    { return nil }
func (e BaseApplication) Destroy() error                 { return nil }
func (e BaseApplication) Prepare([]input.Event, float64) {}
func (e BaseApplication) Render()                        {}
