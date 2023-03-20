package keyboard

// Possible events types
const (
	EventKeyPress    EventType = 10
	EventKeyRelease  EventType = 11
	eventKeyModifier EventType = 12
)

const (
	// Normal characters
	KeyA Key = 0x00
	KeyB Key = 0x0B
	KeyC Key = 0x08
	KeyD Key = 0x02
	KeyE Key = 0x0E
	KeyF Key = 0x03
	KeyG Key = 0x05
	KeyH Key = 0x04
	KeyI Key = 0x22
	KeyJ Key = 0x26
	KeyK Key = 0x28
	KeyL Key = 0x25
	KeyM Key = 0x2E
	KeyN Key = 0x2D
	KeyO Key = 0x1F
	KeyP Key = 0x23
	KeyQ Key = 0x0C
	KeyR Key = 0x0F
	KeyS Key = 0x01
	KeyT Key = 0x11
	KeyU Key = 0x20
	KeyV Key = 0x09
	KeyW Key = 0x0D
	KeyX Key = 0x07
	KeyY Key = 0x10
	KeyZ Key = 0x06

	// Numeric characters
	Key0 Key = 0x1D
	Key1 Key = 0x12
	Key2 Key = 0x13
	Key3 Key = 0x14
	Key4 Key = 0x15
	Key5 Key = 0x17
	Key6 Key = 0x16
	Key7 Key = 0x1A
	Key8 Key = 0x1C
	Key9 Key = 0x19

	// Function keys
	KeyF1  Key = 0x7A
	KeyF2  Key = 0x78
	KeyF3  Key = 0x63
	KeyF4  Key = 0x76
	KeyF5  Key = 0x60
	KeyF6  Key = 0x61
	KeyF7  Key = 0x62
	KeyF8  Key = 0x64
	KeyF9  Key = 0x65
	KeyF10 Key = 0x6D
	KeyF11 Key = 0x67
	KeyF12 Key = 0x6F
	KeyF13 Key = 0x69
	KeyF14 Key = 0x6B
	KeyF15 Key = 0x71
	KeyF16 Key = 0x6A
	KeyF17 Key = 0x40
	KeyF18 Key = 0x4F
	KeyF19 Key = 0x50
	KeyF20 Key = 0x5A

	// Special characters
	KeyEquals       Key = 0x18
	KeyMinus        Key = 0x1B
	KeyComma        Key = 0x2B
	KeyLeftBracket  Key = 0x21
	KeyRightBracket Key = 0x1E
	KeyQuote        Key = 0x27
	KeySemicolon    Key = 0x29
	KeyBackSlash    Key = 0x2A
	KeySlash        Key = 0x2C
	KeyPeriod       Key = 0x2F
	KeyGrave        Key = 0x32
	KeyTab          Key = 0x30
	KeySpace        Key = 0x31

	// Accesor keys
	KeyHome          Key = 0x73
	KeyEnd           Key = 0x77
	KeyPageUp        Key = 0x74
	KeyPageDown      Key = 0x79
	KeyArrowLeft     Key = 0x7B
	KeyArrowRight    Key = 0x7C
	KeyArrowDown     Key = 0x7D
	KeyArrowUp       Key = 0x7E
	KeyCommand       Key = 0x37
	KeyShift         Key = 0x38
	KeyOption        Key = 0x3A
	KeyControl       Key = 0x3B
	KeyRightCommand  Key = 0x36
	KeyRightShift    Key = 0x3C
	KeyRightOption   Key = 0x3D
	KeyRightControl  Key = 0x3E
	KeyCapsLock      Key = 0x39
	KeyFunction      Key = 0x3F
	KeyHelp          Key = 0x72
	KeyEnter         Key = 0x24
	KeyEscape        Key = 0x35
	KeyDelete        Key = 0x33
	KeyForwardDelete Key = 0x75

	// Volumnes
	KeyMute       Key = 0x4A
	KeyVolumeUp   Key = 0x48
	KeyVolumeDown Key = 0x49

	// Keypad
	KeypadDecimal  Key = 0x41
	KeypadMultiply Key = 0x43
	KeypadDivide   Key = 0x4B
	KeypadPlus     Key = 0x45
	KeypadMinus    Key = 0x4E
	KeypadEquals   Key = 0x51
	KeypadEnter    Key = 0x4C
	KeypadClear    Key = 0x47
	Keypad0        Key = 0x52
	Keypad1        Key = 0x53
	Keypad2        Key = 0x54
	Keypad3        Key = 0x55
	Keypad4        Key = 0x56
	Keypad5        Key = 0x57
	Keypad6        Key = 0x58
	Keypad7        Key = 0x59
	Keypad8        Key = 0x5B
	Keypad9        Key = 0x5C
)
