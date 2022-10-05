package win32

import (
	"syscall"
	"unsafe"
)

var (
	modKernel32         = syscall.NewLazyDLL("kernel32.dll")
	procGetModuleHandle = modKernel32.NewProc("GetModuleHandleW")

	moduser32                    = syscall.NewLazyDLL("user32.dll")
	procRegisterClassEx          = moduser32.NewProc("RegisterClassExW")
	procLoadIcon                 = moduser32.NewProc("LoadIconW")
	procLoadCursor               = moduser32.NewProc("LoadCursorW")
	procCreateWindowEx           = moduser32.NewProc("CreateWindowExW")
	procPostMessage              = moduser32.NewProc("PostMessageW")
	procGetCursorPos             = moduser32.NewProc("GetCursorPos")
	procSetForegroundWindow      = moduser32.NewProc("SetForegroundWindow")
	procCreatePopupMenu          = moduser32.NewProc("CreatePopupMenu")
	procTrackPopupMenu           = moduser32.NewProc("TrackPopupMenu")
	procDestroyMenu              = moduser32.NewProc("DestroyMenu")
	procAppendMenuW              = moduser32.NewProc("AppendMenuW")
	procCreateIconFromResourceEx = moduser32.NewProc("CreateIconFromResourceEx")
	procGetMessageW              = moduser32.NewProc("GetMessageW")
	procIsDialogMessage          = moduser32.NewProc("IsDialogMessageW")
	procTranslateMessage         = moduser32.NewProc("TranslateMessage")
	procDispatchMessage          = moduser32.NewProc("DispatchMessageW")

	modshell32          = syscall.NewLazyDLL("shell32.dll")
	procShellNotifyIcon = modshell32.NewProc("Shell_NotifyIconW")
)

type HANDLE uintptr
type HINSTANCE HANDLE
type HICON HANDLE
type HCURSOR HANDLE
type HBRUSH HANDLE
type HWND HANDLE
type HMENU HANDLE
type ATOM uint16

const (
	WM_LBUTTONUP     = 0x0202
	WM_LBUTTONDBLCLK = 0x0203
	WM_RBUTTONUP     = 0x0205
	WM_USER          = 0x0400
	WM_TRAYICON      = WM_USER + 69
	WM_SETTINGCHANGE = 0x001A

	WS_EX_APPWINDOW     = 0x00040000
	WS_OVERLAPPEDWINDOW = 0x00000000 | 0x00C00000 | 0x00080000 | 0x00040000 | 0x00020000 | 0x00010000
	CW_USEDEFAULT       = 0x80000000

	NIM_ADD        = 0x00000000
	NIM_MODIFY     = 0x00000001
	NIM_DELETE     = 0x00000002
	NIM_SETVERSION = 0x00000004

	NIF_MESSAGE = 0x00000001
	NIF_ICON    = 0x00000002
	NIF_TIP     = 0x00000004
	NIF_STATE   = 0x00000008
	NIF_INFO    = 0x00000010

	NIS_HIDDEN = 0x00000001

	NIIF_NONE               = 0x00000000
	NIIF_INFO               = 0x00000001
	NIIF_WARNING            = 0x00000002
	NIIF_ERROR              = 0x00000003
	NIIF_USER               = 0x00000004
	NIIF_NOSOUND            = 0x00000010
	NIIF_LARGE_ICON         = 0x00000020
	NIIF_RESPECT_QUIET_TIME = 0x00000080
	NIIF_ICON_MASK          = 0x0000000F

	IMAGE_BITMAP    = 0
	IMAGE_ICON      = 1
	LR_LOADFROMFILE = 0x00000010
	LR_DEFAULTSIZE  = 0x00000040

	IDC_ARROW     = 32512
	COLOR_WINDOW  = 5
	COLOR_BTNFACE = 15

	GWLP_USERDATA       = -21
	WS_CLIPSIBLINGS     = 0x04000000
	WS_EX_CONTROLPARENT = 0x00010000

	HWND_MESSAGE       = ^HWND(2)
	NOTIFYICON_VERSION = 4

	IDI_APPLICATION = 32512
	WM_APP          = 32768
	WM_COMMAND      = 273

	MenuItemMsgID       = WM_APP + 1024
	NotifyIconMessageId = WM_APP + iota

	MF_STRING       = 0x00000000
	MF_ENABLED      = 0x00000000
	MF_GRAYED       = 0x00000001
	MF_DISABLED     = 0x00000002
	MF_SEPARATOR    = 0x00000800
	MF_CHECKED      = 0x00000008
	MF_MENUBARBREAK = 0x00000020

	TPM_LEFTALIGN = 0x0000
	WM_NULL       = 0

	CS_VREDRAW = 0x0001
	CS_HREDRAW = 0x0002
)

type WindowProc func(hwnd HWND, msg uint32, wparam, lparam uintptr) uintptr

func GetModuleHandle(value uintptr) uintptr {
	result, _, _ := procGetModuleHandle.Call(value)
	return result
}

func GetMessage(msg *MSG) uintptr {
	rt, _, _ := procGetMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
	return rt
}

func PostMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procPostMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func ShellNotifyIcon(cmd uintptr, nid *NOTIFYICONDATA) uintptr {
	ret, _, _ := procShellNotifyIcon.Call(cmd, uintptr(unsafe.Pointer(nid)))
	return ret
}

func IsDialogMessage(hwnd HWND, msg *MSG) uintptr {
	ret, _, _ := procIsDialogMessage.Call(uintptr(hwnd), uintptr(unsafe.Pointer(msg)))
	return ret
}

func TranslateMessage(msg *MSG) uintptr {
	ret, _, _ := procTranslateMessage.Call(uintptr(unsafe.Pointer(msg)))
	return ret
}

func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := procDispatchMessage.Call(uintptr(unsafe.Pointer(msg)))
	return ret
}