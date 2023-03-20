package keyboard

import "syscall"

const (
	// evSync is used as markers to separate events. Events may be separated in time or in space, such as with the multitouch protocol.
	evSync ev = 0x00
	// evKeyChange is used to describe state changes of keyboards, buttons, or other key-like devices.
	evKeyChange ev = 0x01
)

// Possible events types
const (
	EventKeyPress   EventType = 0x1
	EventKeyRelease EventType = 0x0
)

const (
	// Normal characters
	KeyA Key = 0x1e
	KeyB Key = 0x30
	KeyC Key = 0x2e
	KeyD Key = 0x20
	KeyE Key = 0x12
	KeyF Key = 0x21
	KeyG Key = 0x22
	KeyH Key = 0x23
	KeyI Key = 0x17
	KeyJ Key = 0x24
	KeyK Key = 0x25
	KeyL Key = 0x26
	KeyM Key = 0x32
	KeyN Key = 0x31
	KeyO Key = 0x18
	KeyP Key = 0x19
	KeyQ Key = 0x10
	KeyR Key = 0x13
	KeyS Key = 0x1f
	KeyT Key = 0x14
	KeyU Key = 0x16
	KeyV Key = 0x2f
	KeyW Key = 0x11
	KeyX Key = 0x2d
	KeyY Key = 0x15
	KeyZ Key = 0x2c

	// Numeric characters
	Key0 Key = 0x0b
	Key1 Key = 0x02
	Key2 Key = 0x03
	Key3 Key = 0x04
	Key4 Key = 0x05
	Key5 Key = 0x06
	Key6 Key = 0x07
	Key7 Key = 0x08
	Key8 Key = 0x09
	Key9 Key = 0x0a

	// Function keys
	KeyF1  Key = 0x3b
	KeyF2  Key = 0x3c
	KeyF3  Key = 0x3d
	KeyF4  Key = 0x3e
	KeyF5  Key = 0x3f
	KeyF6  Key = 0x40
	KeyF7  Key = 0x41
	KeyF8  Key = 0x42
	KeyF9  Key = 0x43
	KeyF10 Key = 0x44
	KeyF11 Key = 0x57
	KeyF12 Key = 0x58

	// Special characters
	KeyEquals       Key = 0x0d
	KeyMinus        Key = 0x0c
	KeyComma        Key = 0x33
	KeyLeftBracket  Key = 0x1a
	KeyRightBracket Key = 0x1b
	KeyQuote        Key = 0x28
	KeySemicolon    Key = 0x27
	KeyBackSlash    Key = 0x2b
	KeySlash        Key = 0x35
	KeyPeriod       Key = 0x34
	KeyGrave        Key = 0x29
	KeyTab          Key = 0x0f
	KeySpace        Key = 0x39

	// Accesor keys
	KeyHome          Key = 0x66
	KeyEnd           Key = 0x6b
	KeyPageUp        Key = 0x68
	KeyPageDown      Key = 0x6d
	KeyArrowLeft     Key = 0x69
	KeyArrowRight    Key = 0x6a
	KeyArrowDown     Key = 0x6c
	KeyArrowUp       Key = 0x67
	KeyShift         Key = 0x2a
	KeyControl       Key = 0x1d
	KeyAlt           Key = 0x38
	KeyRightShift    Key = 0x36
	KeyRightControl  Key = 0x61
	KeyRightAlt      Key = 0x64
	KeyEnter         Key = 0x1c
	KeyEscape        Key = 0x01
	KeyDelete        Key = 0x6f
	KeyBackspace     Key = 0x0e
	KeyForwardDelete Key = 0x53
	KeyPause         Key = 0x77
	KeyInsert        Key = 0x6e
	KeyCapsLock      Key = 0x3a
	KeyNumlock       Key = 0x45
	KeyScrolllock    Key = 0x46

	// Keypad
	KeypadDecimal  Key = 0x34
	KeypadMultiply Key = 0x37
	KeypadDivide   Key = 0x62
	KeypadPlus     Key = 0x4e
	KeypadMinus    Key = 0x0c
	KeypadEquals   Key = 0x0d
	KeypadEnter    Key = 0x60
	Keypad0        Key = 0x0b
	Keypad1        Key = 0x02
	Keypad2        Key = 0x03
	Keypad3        Key = 0x04
	Keypad4        Key = 0x05
	Keypad5        Key = 0x06
	Keypad6        Key = 0x07
	Keypad7        Key = 0x08
	Keypad8        Key = 0x09
	Keypad9        Key = 0x0a
)

var charKeymap = map[byte][]Key{
	'0':  {Key0},
	'1':  {Key1},
	'2':  {Key2},
	'3':  {Key3},
	'4':  {Key4},
	'5':  {Key5},
	'6':  {Key6},
	'7':  {Key7},
	'8':  {Key8},
	'9':  {Key9},
	'a':  {KeyA},
	'b':  {KeyB},
	'c':  {KeyC},
	'd':  {KeyD},
	'e':  {KeyE},
	'f':  {KeyF},
	'g':  {KeyG},
	'h':  {KeyH},
	'i':  {KeyI},
	'j':  {KeyJ},
	'k':  {KeyK},
	'l':  {KeyL},
	'm':  {KeyM},
	'n':  {KeyN},
	'o':  {KeyO},
	'p':  {KeyP},
	'q':  {KeyQ},
	'r':  {KeyR},
	's':  {KeyS},
	't':  {KeyT},
	'u':  {KeyU},
	'v':  {KeyV},
	'w':  {KeyW},
	'x':  {KeyX},
	'y':  {KeyY},
	'z':  {KeyZ},
	'A':  {KeyShift, KeyA},
	'B':  {KeyShift, KeyB},
	'C':  {KeyShift, KeyC},
	'D':  {KeyShift, KeyD},
	'E':  {KeyShift, KeyE},
	'F':  {KeyShift, KeyF},
	'G':  {KeyShift, KeyG},
	'H':  {KeyShift, KeyH},
	'I':  {KeyShift, KeyI},
	'J':  {KeyShift, KeyJ},
	'K':  {KeyShift, KeyK},
	'L':  {KeyShift, KeyL},
	'M':  {KeyShift, KeyM},
	'N':  {KeyShift, KeyN},
	'O':  {KeyShift, KeyO},
	'P':  {KeyShift, KeyP},
	'Q':  {KeyShift, KeyQ},
	'R':  {KeyShift, KeyR},
	'S':  {KeyShift, KeyS},
	'T':  {KeyShift, KeyT},
	'U':  {KeyShift, KeyU},
	'V':  {KeyShift, KeyV},
	'W':  {KeyShift, KeyW},
	'X':  {KeyShift, KeyX},
	'Y':  {KeyShift, KeyY},
	'Z':  {KeyShift, KeyZ},
	'-':  {KeyMinus},
	'=':  {KeyEquals},
	'\b': {KeyTab},
	'\t': {KeyTab},
	'\n': {KeyEnter},
	'[':  {KeyLeftBracket},
	']':  {KeyRightBracket},
	';':  {KeySemicolon},
	'\\': {KeyBackSlash},
	'\'': {KeyQuote},
	'`':  {KeyGrave},
	',':  {KeyComma},
	'.':  {KeyPeriod},
	'/':  {KeySlash},
	' ':  {KeySpace},
	'{':  {KeyShift, KeyLeftBracket},
	'}':  {KeyShift, KeyRightBracket},
	':':  {KeyShift, KeySemicolon},
	'"':  {KeyShift, KeyQuote},
	'~':  {KeyShift, KeyGrave},
	'<':  {KeyShift, KeyComma},
	'>':  {KeyShift, KeyPeriod},
	'?':  {KeyShift, KeySlash},
	'_':  {KeyShift, KeyMinus},
	'+':  {KeyShift, KeyEquals},
	'|':  {KeyShift, KeyBackSlash},
}

type ev uint16

type inputEvent struct {
	Time syscall.Timeval
	EV   ev
	Key  Key
	Type EventType
}
