package input

import "github.com/veandco/go-sdl2/sdl"

// Mirror sdl keys
const (
	K_UNKNOWN = sdl.K_UNKNOWN

	K_RETURN     = sdl.K_RETURN
	K_ESCAPE     = sdl.K_ESCAPE
	K_BACKSPACE  = sdl.K_BACKSPACE
	K_TAB        = sdl.K_TAB
	K_SPACE      = sdl.K_SPACE
	K_EXCLAIM    = sdl.K_EXCLAIM
	K_QUOTEDBL   = sdl.K_QUOTEDBL
	K_HASH       = sdl.K_HASH
	K_PERCENT    = sdl.K_PERCENT
	K_DOLLAR     = sdl.K_DOLLAR
	K_AMPERSAND  = sdl.K_AMPERSAND
	K_QUOTE      = sdl.K_QUOTE
	K_LEFTPAREN  = sdl.K_LEFTPAREN
	K_RIGHTPAREN = sdl.K_RIGHTPAREN
	K_ASTERISK   = sdl.K_ASTERISK
	K_PLUS       = sdl.K_PLUS
	K_COMMA      = sdl.K_COMMA
	K_MINUS      = sdl.K_MINUS
	K_PERIOD     = sdl.K_PERIOD
	K_SLASH      = sdl.K_SLASH
	K_0          = sdl.K_0
	K_1          = sdl.K_1
	K_2          = sdl.K_2
	K_3          = sdl.K_3
	K_4          = sdl.K_4
	K_5          = sdl.K_5
	K_6          = sdl.K_6
	K_7          = sdl.K_7
	K_8          = sdl.K_8
	K_9          = sdl.K_9
	K_COLON      = sdl.K_COLON
	K_SEMICOLON  = sdl.K_SEMICOLON
	K_LESS       = sdl.K_LESS
	K_EQUALS     = sdl.K_EQUALS
	K_GREATER    = sdl.K_GREATER
	K_QUESTION   = sdl.K_QUESTION
	K_AT         = sdl.K_AT

	K_LEFTBRACKET  = sdl.K_LEFTBRACKET
	K_BACKSLASH    = sdl.K_BACKSLASH
	K_RIGHTBRACKET = sdl.K_RIGHTBRACKET
	K_CARET        = sdl.K_CARET
	K_UNDERSCORE   = sdl.K_UNDERSCORE
	K_BACKQUOTE    = sdl.K_BACKQUOTE
	K_a            = sdl.K_a
	K_b            = sdl.K_b
	K_c            = sdl.K_c
	K_d            = sdl.K_d
	K_e            = sdl.K_e
	K_f            = sdl.K_f
	K_g            = sdl.K_g
	K_h            = sdl.K_h
	K_i            = sdl.K_i
	K_j            = sdl.K_j
	K_k            = sdl.K_k
	K_l            = sdl.K_l
	K_m            = sdl.K_m
	K_n            = sdl.K_n
	K_o            = sdl.K_o
	K_p            = sdl.K_p
	K_q            = sdl.K_q
	K_r            = sdl.K_r
	K_s            = sdl.K_s
	K_t            = sdl.K_t
	K_u            = sdl.K_u
	K_v            = sdl.K_v
	K_w            = sdl.K_w
	K_x            = sdl.K_x
	K_y            = sdl.K_y
	K_z            = sdl.K_z

	K_CAPSLOCK = sdl.K_CAPSLOCK

	K_F1  = sdl.K_F1
	K_F2  = sdl.K_F2
	K_F3  = sdl.K_F3
	K_F4  = sdl.K_F4
	K_F5  = sdl.K_F5
	K_F6  = sdl.K_F6
	K_F7  = sdl.K_F7
	K_F8  = sdl.K_F8
	K_F9  = sdl.K_F9
	K_F10 = sdl.K_F10
	K_F11 = sdl.K_F11
	K_F12 = sdl.K_F12

	K_PRINTSCREEN = sdl.K_PRINTSCREEN
	K_SCROLLLOCK  = sdl.K_SCROLLLOCK
	K_PAUSE       = sdl.K_PAUSE
	K_INSERT      = sdl.K_INSERT
	K_HOME        = sdl.K_HOME
	K_PAGEUP      = sdl.K_PAGEUP
	K_DELETE      = sdl.K_DELETE
	K_END         = sdl.K_END
	K_PAGEDOWN    = sdl.K_PAGEDOWN
	K_RIGHT       = sdl.K_RIGHT
	K_LEFT        = sdl.K_LEFT
	K_DOWN        = sdl.K_DOWN
	K_UP          = sdl.K_UP

	K_NUMLOCKCLEAR = sdl.K_NUMLOCKCLEAR
	K_KP_DIVIDE    = sdl.K_KP_DIVIDE
	K_KP_MULTIPLY  = sdl.K_KP_MULTIPLY
	K_KP_MINUS     = sdl.K_KP_MINUS
	K_KP_PLUS      = sdl.K_KP_PLUS
	K_KP_ENTER     = sdl.K_KP_ENTER
	K_KP_1         = sdl.K_KP_1
	K_KP_2         = sdl.K_KP_2
	K_KP_3         = sdl.K_KP_3
	K_KP_4         = sdl.K_KP_4
	K_KP_5         = sdl.K_KP_5
	K_KP_6         = sdl.K_KP_6
	K_KP_7         = sdl.K_KP_7
	K_KP_8         = sdl.K_KP_8
	K_KP_9         = sdl.K_KP_9
	K_KP_0         = sdl.K_KP_0
	K_KP_PERIOD    = sdl.K_KP_PERIOD

	K_APPLICATION    = sdl.K_APPLICATION
	K_POWER          = sdl.K_POWER
	K_KP_EQUALS      = sdl.K_KP_EQUALS
	K_F13            = sdl.K_F13
	K_F14            = sdl.K_F14
	K_F15            = sdl.K_F15
	K_F16            = sdl.K_F16
	K_F17            = sdl.K_F17
	K_F18            = sdl.K_F18
	K_F19            = sdl.K_F19
	K_F20            = sdl.K_F20
	K_F21            = sdl.K_F21
	K_F22            = sdl.K_F22
	K_F23            = sdl.K_F23
	K_F24            = sdl.K_F24
	K_EXECUTE        = sdl.K_EXECUTE
	K_HELP           = sdl.K_HELP
	K_MENU           = sdl.K_MENU
	K_SELECT         = sdl.K_SELECT
	K_STOP           = sdl.K_STOP
	K_AGAIN          = sdl.K_AGAIN
	K_UNDO           = sdl.K_UNDO
	K_CUT            = sdl.K_CUT
	K_COPY           = sdl.K_COPY
	K_PASTE          = sdl.K_PASTE
	K_FIND           = sdl.K_FIND
	K_MUTE           = sdl.K_MUTE
	K_VOLUMEUP       = sdl.K_VOLUMEUP
	K_VOLUMEDOWN     = sdl.K_VOLUMEDOWN
	K_KP_COMMA       = sdl.K_KP_COMMA
	K_KP_EQUALSAS400 = sdl.K_KP_EQUALSAS400

	K_ALTERASE   = sdl.K_ALTERASE
	K_SYSREQ     = sdl.K_SYSREQ
	K_CANCEL     = sdl.K_CANCEL
	K_CLEAR      = sdl.K_CLEAR
	K_PRIOR      = sdl.K_PRIOR
	K_RETURN2    = sdl.K_RETURN2
	K_SEPARATOR  = sdl.K_SEPARATOR
	K_OUT        = sdl.K_OUT
	K_OPER       = sdl.K_OPER
	K_CLEARAGAIN = sdl.K_CLEARAGAIN
	K_CRSEL      = sdl.K_CRSEL
	K_EXSEL      = sdl.K_EXSEL

	K_KP_00              = sdl.K_KP_00
	K_KP_000             = sdl.K_KP_000
	K_THOUSANDSSEPARATOR = sdl.K_THOUSANDSSEPARATOR
	K_DECIMALSEPARATOR   = sdl.K_DECIMALSEPARATOR
	K_CURRENCYUNIT       = sdl.K_CURRENCYUNIT
	K_CURRENCYSUBUNIT    = sdl.K_CURRENCYSUBUNIT
	K_KP_LEFTPAREN       = sdl.K_KP_LEFTPAREN
	K_KP_RIGHTPAREN      = sdl.K_KP_RIGHTPAREN
	K_KP_LEFTBRACE       = sdl.K_KP_LEFTBRACE
	K_KP_RIGHTBRACE      = sdl.K_KP_RIGHTBRACE
	K_KP_TAB             = sdl.K_KP_TAB
	K_KP_BACKSPACE       = sdl.K_KP_BACKSPACE
	K_KP_A               = sdl.K_KP_A
	K_KP_B               = sdl.K_KP_B
	K_KP_C               = sdl.K_KP_C
	K_KP_D               = sdl.K_KP_D
	K_KP_E               = sdl.K_KP_E
	K_KP_F               = sdl.K_KP_F
	K_KP_XOR             = sdl.K_KP_XOR
	K_KP_POWER           = sdl.K_KP_POWER
	K_KP_PERCENT         = sdl.K_KP_PERCENT
	K_KP_LESS            = sdl.K_KP_LESS
	K_KP_GREATER         = sdl.K_KP_GREATER
	K_KP_AMPERSAND       = sdl.K_KP_AMPERSAND
	K_KP_DBLAMPERSAND    = sdl.K_KP_DBLAMPERSAND
	K_KP_VERTICALBAR     = sdl.K_KP_VERTICALBAR
	K_KP_DBLVERTICALBAR  = sdl.K_KP_DBLVERTICALBAR
	K_KP_COLON           = sdl.K_KP_COLON
	K_KP_HASH            = sdl.K_KP_HASH
	K_KP_SPACE           = sdl.K_KP_SPACE
	K_KP_AT              = sdl.K_KP_AT
	K_KP_EXCLAM          = sdl.K_KP_EXCLAM
	K_KP_MEMSTORE        = sdl.K_KP_MEMSTORE
	K_KP_MEMRECALL       = sdl.K_KP_MEMRECALL
	K_KP_MEMCLEAR        = sdl.K_KP_MEMCLEAR
	K_KP_MEMADD          = sdl.K_KP_MEMADD
	K_KP_MEMSUBTRACT     = sdl.K_KP_MEMSUBTRACT
	K_KP_MEMMULTIPLY     = sdl.K_KP_MEMMULTIPLY
	K_KP_MEMDIVIDE       = sdl.K_KP_MEMDIVIDE
	K_KP_PLUSMINUS       = sdl.K_KP_PLUSMINUS
	K_KP_CLEAR           = sdl.K_KP_CLEAR
	K_KP_CLEARENTRY      = sdl.K_KP_CLEARENTRY
	K_KP_BINARY          = sdl.K_KP_BINARY
	K_KP_OCTAL           = sdl.K_KP_OCTAL
	K_KP_DECIMAL         = sdl.K_KP_DECIMAL
	K_KP_HEXADECIMAL     = sdl.K_KP_HEXADECIMAL

	K_LCTRL  = sdl.K_LCTRL
	K_LSHIFT = sdl.K_LSHIFT
	K_LALT   = sdl.K_LALT
	K_LGUI   = sdl.K_LGUI
	K_RCTRL  = sdl.K_RCTRL
	K_RSHIFT = sdl.K_RSHIFT
	K_RALT   = sdl.K_RALT
	K_RGUI   = sdl.K_RGUI

	K_MODE = sdl.K_MODE

	K_AUDIONEXT    = sdl.K_AUDIONEXT
	K_AUDIOPREV    = sdl.K_AUDIOPREV
	K_AUDIOSTOP    = sdl.K_AUDIOSTOP
	K_AUDIOPLAY    = sdl.K_AUDIOPLAY
	K_AUDIOMUTE    = sdl.K_AUDIOMUTE
	K_MEDIASELECT  = sdl.K_MEDIASELECT
	K_WWW          = sdl.K_WWW
	K_MAIL         = sdl.K_MAIL
	K_CALCULATOR   = sdl.K_CALCULATOR
	K_COMPUTER     = sdl.K_COMPUTER
	K_AC_SEARCH    = sdl.K_AC_SEARCH
	K_AC_HOME      = sdl.K_AC_HOME
	K_AC_BACK      = sdl.K_AC_BACK
	K_AC_FORWARD   = sdl.K_AC_FORWARD
	K_AC_STOP      = sdl.K_AC_STOP
	K_AC_REFRESH   = sdl.K_AC_REFRESH
	K_AC_BOOKMARKS = sdl.K_AC_BOOKMARKS

	K_BRIGHTNESSDOWN = sdl.K_BRIGHTNESSDOWN
	K_BRIGHTNESSUP   = sdl.K_BRIGHTNESSUP
	K_DISPLAYSWITCH  = sdl.K_DISPLAYSWITCH
	K_KBDILLUMTOGGLE = sdl.K_KBDILLUMTOGGLE
	K_KBDILLUMDOWN   = sdl.K_KBDILLUMDOWN
	K_KBDILLUMUP     = sdl.K_KBDILLUMUP
	K_EJECT          = sdl.K_EJECT
	K_SLEEP          = sdl.K_SLEEP
)

// mirror SDL mod keys
const (
	MOD_NONE   = sdl.KMOD_NONE
	MOD_LSHIFT = sdl.KMOD_LSHIFT
	MOD_RSHIFT = sdl.KMOD_RSHIFT
	MOD_LCTRL  = sdl.KMOD_LCTRL
	MOD_RCTRL  = sdl.KMOD_RCTRL
	MOD_LALT   = sdl.KMOD_LALT
	MOD_RALT   = sdl.KMOD_RALT
	MOD_LGUI   = sdl.KMOD_LGUI
	MOD_RGUI   = sdl.KMOD_RGUI
	MOD_NUM    = sdl.KMOD_NUM
	MOD_CAPS   = sdl.KMOD_CAPS
	MOD_MODE   = sdl.KMOD_MODE
	MOD_CTRL   = sdl.KMOD_CTRL
	MOD_SHIFT  = sdl.KMOD_SHIFT
	MOD_ALT    = sdl.KMOD_ALT
	MOD_GUI    = sdl.KMOD_GUI
)

// mirror SDL mouse buttons
const (
	BUTTON_LEFT   = sdl.BUTTON_LEFT
	BUTTON_MIDDLE = sdl.BUTTON_MIDDLE
	BUTTON_RIGHT  = sdl.BUTTON_RIGHT
	BUTTON_X1     = sdl.BUTTON_X1
	BUTTON_X2     = sdl.BUTTON_X2
)
