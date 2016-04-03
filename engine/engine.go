package engine

import (
	"runtime"
	"time"

	"github.com/losinggeneration/lgge/application"
	"github.com/losinggeneration/lgge/input"
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

	then := time.Now()
	for w.IsRunning() {
		events := processEvents()

		now := time.Now()
		since := now.Sub(then)
		then = now

		app.Prepare(events, float64(since)/float64(time.Second))
		app.Render()

		w.SwapBuffers()
	}

	return nil
}
