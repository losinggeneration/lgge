package gl

import (
	"github.com/go-gl/gl/v3.3-compatibility/gl"

	_gl "github.com/losinggeneration/lgge/gl"
	"github.com/losinggeneration/lgge/glu"
	_ "github.com/losinggeneration/lgge/glu/v3.3"
)

func init() {
	_gl.RegisterInit(_Init)
	_gl.RegisterResize(_Resize)
}

func _Init() {
	gl.Init()
}

func _Resize(w, h int) {
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
