package engine

import (
	"runtime"

	"github.com/losinggeneration/lgge/application"
	"github.com/losinggeneration/lgge/input"
	"github.com/losinggeneration/lgge/timer"
	"github.com/losinggeneration/lgge/window"
)

func Run(backend string, cfg window.WindowConfig, app application.Application) error {
	runtime.LockOSThread()

	w, err := window.GetWindow(backend)
	if err != nil {
		return err
	}

	processEvents, err := input.GetProcess(backend)
	if err != nil {
		return err
	}

	if err := w.Create(cfg); err != nil {
		return err
	}
	defer w.Destroy()

	if err := app.Init(); err != nil {
		return err
	}
	defer app.Destroy()

	app.SetWindow(w)
	t := timer.NewTimer()

	for app.IsRunning() {
		events := processEvents()

		t.Update()

		app.Prepare(events, t.Delta())
		app.Render()

		w.SwapBuffers()
	}

	return nil
}
