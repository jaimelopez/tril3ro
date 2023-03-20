package keyboard

// Possible events types
const (
	EventKeyPress   EventType = 256
	EventKeyRelease EventType = 257
)

const (
	// Normal characters
	KeyA Key = 0x41
	KeyB Key = 0x42
	KeyC Key = 0x43
	KeyD Key = 0x44
	KeyE Key = 0x45
	KeyF Key = 0x46
	KeyG Key = 0x47
	KeyH Key = 0x48
	KeyI Key = 0x49
	KeyJ Key = 0x4A
	KeyK Key = 0x4B
	KeyL Key = 0x4C
	KeyM Key = 0x4D
	KeyN Key = 0x4E
	KeyO Key = 0x4F
	KeyP Key = 0x50
	KeyQ Key = 0x51
	KeyR Key = 0x52
	KeyS Key = 0x53
	KeyT Key = 0x54
	KeyU Key = 0x55
	KeyV Key = 0x56
	KeyW Key = 0x57
	KeyX Key = 0x58
	KeyY Key = 0x59
	KeyZ Key = 0x5A

	// Numeric characters
	Key0 Key = 0x30
	Key1 Key = 0x31
	Key2 Key = 0x32
	Key3 Key = 0x33
	Key4 Key = 0x34
	Key5 Key = 0x35
	Key6 Key = 0x36
	Key7 Key = 0x37
	Key8 Key = 0x38
	Key9 Key = 0x39

	// Function keys
	KeyF1  Key = 0x70
	KeyF2  Key = 0x71
	KeyF3  Key = 0x72
	KeyF4  Key = 0x73
	KeyF5  Key = 0x74
	KeyF6  Key = 0x75
	KeyF7  Key = 0x76
	KeyF8  Key = 0x77
	KeyF9  Key = 0x78
	KeyF10 Key = 0x79
	KeyF11 Key = 0x7A
	KeyF12 Key = 0x7B
	KeyF13 Key = 0x7C
	KeyF14 Key = 0x7D
	KeyF15 Key = 0x7E
	KeyF16 Key = 0x7F
	KeyF17 Key = 0x80
	KeyF18 Key = 0x81
	KeyF19 Key = 0x82
	KeyF20 Key = 0x83
	KeyF21 Key = 0x84
	KeyF22 Key = 0x85
	KeyF23 Key = 0x86
	KeyF24 Key = 0x87

	// Special characters
	KeyDecimal  Key = 0x6E
	KeyMultiply Key = 0x6A
	KeyDivide   Key = 0x6F
	KeyPlus     Key = 0x6B
	KeyMinus    Key = 0x6D
	KeyTab      Key = 0x09
	KeySpace    Key = 0x20

	// Accesor keys
	KeyHome         Key = 0x24
	KeyEnd          Key = 0x23
	KeyPageUp       Key = 0x21
	KeyPageDown     Key = 0x22
	KeyArrowLeft    Key = 0x25
	KeyArrowRight   Key = 0x27
	KeyArrowDown    Key = 0x28
	KeyArrowUp      Key = 0x26
	KeyShift        Key = 0x10
	KeyControl      Key = 0x11
	KeyAlt          Key = 0x12
	KeyLeftShift    Key = 0xA0
	KeyLeftControl  Key = 0xA2
	KeyLeftMenu     Key = 0xA4
	KeyLeftWindows  Key = 0x5B
	KeyRightShift   Key = 0xA1
	KeyRightControl Key = 0xA3
	KeyRightMenu    Key = 0xA5
	KeyRightWindows Key = 0x5C
	KeyCapsLock     Key = 0x14
	KeyHelp         Key = 0x2F
	KeyEnter        Key = 0x0D
	KeyEscape       Key = 0x1B
	KeyDelete       Key = 0x2E
	KeyBackspace    Key = 0x08
	KeyPause        Key = 0x13
	KeyExecute      Key = 0x2B
	KeySnapshot     Key = 0x2C
	KeyInsert       Key = 0x2D
	KeyApps         Key = 0x5D
	KeySleep        Key = 0x5F
	KeySeparator    Key = 0x6C
	KeyNumlock      Key = 0x90
	KeyScroll       Key = 0x91
	KeyClear        Key = 0x0C

	// Volumnes
	KeyMute       Key = 0xAD
	KeyVolumeUp   Key = 0xAF
	KeyVolumeDown Key = 0xAE

	// Keypad
	Keypad0 Key = 0x60
	Keypad1 Key = 0x61
	Keypad2 Key = 0x62
	Keypad3 Key = 0x63
	Keypad4 Key = 0x64
	Keypad5 Key = 0x65
	Keypad6 Key = 0x66
	Keypad7 Key = 0x67
	Keypad8 Key = 0x68
	Keypad9 Key = 0x69
)

const (
	// Type of hook when calling SetWindowsHookExW
	// https://learn.microsoft.com/en-gb/windows/win32/api/winuser/nf-winuser-setwindowshookexa
	hookKeyboardLL uintptr = 13

	// Represents the shift modifier when comes from VkKeyScanA
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-vkkeyscana#return-value
	scanShiftModifier uintptr = 0x0001
)
