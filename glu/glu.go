package glu

var _Perspective func(fovy, aspect, near, far float64)
var _LookAt func(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64)

func RegisterPerspective(fn func(fovy, aspect, near, far float64)) {
	_Perspective = fn
}

func RegisterLookAt(fn func(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64)) {
	_LookAt = fn
}

func Perspective(fovy, aspect, near, far float64) {
	if _Perspective == nil {
		panic("Please register a Perspective with glu.RegisterPespective")
	}

	_Perspective(fovy, aspect, near, far)
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64) {
	if _LookAt == nil {
		panic("Please register a LookAt with glu.RegisterLookAt")
	}

	_LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ)
}
