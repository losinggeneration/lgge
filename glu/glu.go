package glu

import (
	"math"

	"github.com/go-gl/gl/v3.3-compatibility/gl"
	"github.com/go-gl/mathgl/mgl64"
)

func Perspective(fovy, aspect, near, far float64) {
	M := mgl64.Perspective((fovy*math.Pi)/180.0, aspect, near, far)
	gl.MultMatrixd(&M[0])
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64) {
	M := mgl64.LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ)
	gl.MultMatrixd(&M[0])
}
