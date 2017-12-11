package gl

var (
	_Init   func()
	_Resize func(int, int)
)

func RegisterInit(fn func()) {
	_Init = fn
}

func RegisterResize(fn func(int, int)) {
	_Resize = fn
}

func Init() {
	if _Init == nil {
		panic("register an Init function with RegisterInit")
	}

	_Init()
}

func Resize(w, h int) {
	if _Resize == nil {
		panic("register an Resize function with RegisterResize")
	}

	_Resize(w, h)
}
