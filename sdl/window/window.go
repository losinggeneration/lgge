package common

import (
	"github.com/losinggeneration/lgge/gl"
	"github.com/losinggeneration/lgge/window"

	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	window.RegisterWindow("sdl", &sdlWindow{})
}

type sdlWindow struct {
	window.WindowConfig

	window  *sdl.Window
	context sdl.GLContext
}

func (w *sdlWindow) Create(cfg window.WindowConfig) error {
	var flags uint32
	flags = sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE | sdl.WINDOW_OPENGL

	w.WindowConfig = cfg

	if w.Fullscreen {
		flags = flags | sdl.WINDOW_FULLSCREEN
	}

	window, err := sdl.CreateWindow(w.Title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, w.Width, w.Height, flags)
	if err != nil {
		return err
	}

	w.window = window

	context, err := sdl.GL_CreateContext(window)
	if err != nil {
		return err
	}

	if err := sdl.GL_SetSwapInterval(1); err != nil {
		return err
	}

	w.context = context
	w.Running = true

	gl.Init()
	w.Resize(w.Width, w.Height)

	return nil
}

func (w *sdlWindow) Destroy() error {
	w.window.Destroy()
	sdl.GL_DeleteContext(w.context)

	return nil
}

func (w sdlWindow) SwapBuffers() {
	sdl.GL_SwapWindow(w.window)
}

func (window *sdlWindow) Resize(w, h int) {
	if h <= 0 {
		h = 1
	}

	gl.Resize(w, h)
}
