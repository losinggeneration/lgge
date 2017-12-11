package glu

import (
	"math"

	"github.com/go-gl/gl/v3.3-compatibility/gl"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/losinggeneration/lgge/glu"
)

func init() {
	glu.RegisterPerspective(_Perspective)
	glu.RegisterLookAt(_LookAt)
}

func _Perspective(fovy, aspect, near, far float64) {
	M := mgl64.Perspective((fovy*math.Pi)/180.0, aspect, near, far)
	gl.MultMatrixd(&M[0])
}

func _LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64) {
	M := mgl64.LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ)
	gl.MultMatrixd(&M[0])
}
