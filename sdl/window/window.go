package common

import (
	"github.com/losinggeneration/lgge/glu"
	"github.com/losinggeneration/lgge/window"

	"github.com/go-gl/gl/v3.3-compatibility/gl"
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

	w.context = context
	w.Running = true

	gl.Init()
	gl.Enable(gl.DEPTH_TEST)

	w.resize(w.Width, w.Height)

	return nil
}

func (w *sdlWindow) Destroy() error {
	w.window.Destroy()
	sdl.GL_DeleteContext(w.context)

	return nil
}

func (w sdlWindow) IsRunning() bool {
	return w.Running
}

func (w *sdlWindow) ProcessEvents() {
	event := sdl.PollEvent()
	for event != nil {
		switch e := event.(type) {
		case *sdl.KeyDownEvent:
			if e.Keysym.Sym == sdl.K_ESCAPE || e.Keysym.Sym == sdl.K_q {
				w.Running = false
			}

		case *sdl.WindowEvent:
			if e.Event == sdl.WINDOWEVENT_RESIZED {
				w.resize(int(e.Data1), int(e.Data2))
			}

		case *sdl.QuitEvent:
			w.Running = false
		}

		event = sdl.PollEvent()
	}
}

func (w sdlWindow) SwapBuffers() {
	sdl.GL_SwapWindow(w.window)
}

func (window *sdlWindow) resize(w, h int) {
	if h <= 0 {
		h = 1
	}

	gl.Viewport(0, 0, int32(w), int32(h))

	// Future OpenGL shader perspective matrix
	//projection := mgl.Perspective(60, float32(w)/float32(h), 1, 100)
	//view := mgl.LookAt(0, 0, 0, 0, 0, 0, 0, 0, 0)
	//model := mgl.Ident4()
	//mvp := projection.Mul4(view).Mul4(model)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(60, float64(w)/float64(h), 1, 100)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}
