package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type GameController_ITF interface {
	std_core.QObject_ITF
	GameController_PTR() *GameController
}

func (ptr *GameController) GameController_PTR() *GameController {
	return ptr
}

func (ptr *GameController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *GameController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromGameController(ptr GameController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.GameController_PTR().Pointer()
	}
	return nil
}

func NewGameControllerFromPointer(ptr unsafe.Pointer) (n *GameController) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(GameController)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *GameController:
			n = deduced

		case *std_core.QObject:
			n = &GameController{QObject: *deduced}

		default:
			n = new(GameController)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackGameControllercbfc99_Constructor
func callbackGameControllercbfc99_Constructor(ptr unsafe.Pointer) {
	this := NewGameControllerFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackGameControllercbfc99_Login
func callbackGameControllercbfc99_Login(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "login"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *GameController) ConnectLogin(f func(data string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "login"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "login", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "login", f)
		}
	}
}

func (ptr *GameController) DisconnectLogin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "login")
	}
}

func (ptr *GameController) Login(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.GameControllercbfc99_Login(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackGameControllercbfc99_JoinRoom
func callbackGameControllercbfc99_JoinRoom(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "joinRoom"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *GameController) ConnectJoinRoom(f func(data string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "joinRoom"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "joinRoom", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "joinRoom", f)
		}
	}
}

func (ptr *GameController) DisconnectJoinRoom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "joinRoom")
	}
}

func (ptr *GameController) JoinRoom(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.GameControllercbfc99_JoinRoom(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackGameControllercbfc99_SendWord
func callbackGameControllercbfc99_SendWord(ptr unsafe.Pointer, data C.struct_Moc_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "sendWord"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *GameController) ConnectSendWord(f func(data string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sendWord"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sendWord", func(data string) bool {
				signal.(func(string) bool)(data)
				return f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendWord", f)
		}
	}
}

func (ptr *GameController) DisconnectSendWord() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sendWord")
	}
}

func (ptr *GameController) SendWord(data string) bool {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return C.GameControllercbfc99_SendWord(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))}) != 0
	}
	return false
}

//export callbackGameControllercbfc99_StartGame
func callbackGameControllercbfc99_StartGame(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "startGame"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *GameController) ConnectStartGame(f func(data string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startGame"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startGame", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startGame", f)
		}
	}
}

func (ptr *GameController) DisconnectStartGame() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startGame")
	}
}

func (ptr *GameController) StartGame(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.GameControllercbfc99_StartGame(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackGameControllercbfc99_CreateRoom
func callbackGameControllercbfc99_CreateRoom(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "createRoom"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *GameController) ConnectCreateRoom(f func(data string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createRoom"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createRoom", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createRoom", f)
		}
	}
}

func (ptr *GameController) DisconnectCreateRoom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createRoom")
	}
}

func (ptr *GameController) CreateRoom(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.GameControllercbfc99_CreateRoom(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackGameControllercbfc99_FinishGame
func callbackGameControllercbfc99_FinishGame(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finishGame"); signal != nil {
		signal.(func())()
	}

}

func (ptr *GameController) ConnectFinishGame(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "finishGame"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finishGame", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finishGame", f)
		}
	}
}

func (ptr *GameController) DisconnectFinishGame() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "finishGame")
	}
}

func (ptr *GameController) FinishGame() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_FinishGame(ptr.Pointer())
	}
}

//export callbackGameControllercbfc99_AppendUser
func callbackGameControllercbfc99_AppendUser(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "appendUser"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *GameController) ConnectAppendUser(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "appendUser") {
			C.GameControllercbfc99_ConnectAppendUser(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "appendUser"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "appendUser", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "appendUser", f)
		}
	}
}

func (ptr *GameController) DisconnectAppendUser() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectAppendUser(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "appendUser")
	}
}

func (ptr *GameController) AppendUser(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.GameControllercbfc99_AppendUser(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

//export callbackGameControllercbfc99_AppendRoom
func callbackGameControllercbfc99_AppendRoom(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "appendRoom"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *GameController) ConnectAppendRoom(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "appendRoom") {
			C.GameControllercbfc99_ConnectAppendRoom(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "appendRoom"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "appendRoom", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "appendRoom", f)
		}
	}
}

func (ptr *GameController) DisconnectAppendRoom() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectAppendRoom(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "appendRoom")
	}
}

func (ptr *GameController) AppendRoom(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.GameControllercbfc99_AppendRoom(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

//export callbackGameControllercbfc99_DeleteUser
func callbackGameControllercbfc99_DeleteUser(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "deleteUser"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *GameController) ConnectDeleteUser(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "deleteUser") {
			C.GameControllercbfc99_ConnectDeleteUser(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "deleteUser"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "deleteUser", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deleteUser", f)
		}
	}
}

func (ptr *GameController) DisconnectDeleteUser() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectDeleteUser(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "deleteUser")
	}
}

func (ptr *GameController) DeleteUser(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.GameControllercbfc99_DeleteUser(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

//export callbackGameControllercbfc99_UpdateRooms
func callbackGameControllercbfc99_UpdateRooms(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "updateRooms"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *GameController) ConnectUpdateRooms(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateRooms") {
			C.GameControllercbfc99_ConnectUpdateRooms(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateRooms"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateRooms", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateRooms", f)
		}
	}
}

func (ptr *GameController) DisconnectUpdateRooms() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectUpdateRooms(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateRooms")
	}
}

func (ptr *GameController) UpdateRooms(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.GameControllercbfc99_UpdateRooms(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

//export callbackGameControllercbfc99_SessionAuthenticated
func callbackGameControllercbfc99_SessionAuthenticated(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sessionAuthenticated"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *GameController) ConnectSessionAuthenticated(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sessionAuthenticated") {
			C.GameControllercbfc99_ConnectSessionAuthenticated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sessionAuthenticated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sessionAuthenticated", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sessionAuthenticated", f)
		}
	}
}

func (ptr *GameController) DisconnectSessionAuthenticated() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectSessionAuthenticated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sessionAuthenticated")
	}
}

func (ptr *GameController) SessionAuthenticated(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.GameControllercbfc99_SessionAuthenticated(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

//export callbackGameControllercbfc99_SessionAuthenticationError
func callbackGameControllercbfc99_SessionAuthenticationError(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sessionAuthenticationError"); signal != nil {
		signal.(func())()
	}

}

func (ptr *GameController) ConnectSessionAuthenticationError(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sessionAuthenticationError") {
			C.GameControllercbfc99_ConnectSessionAuthenticationError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sessionAuthenticationError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sessionAuthenticationError", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sessionAuthenticationError", f)
		}
	}
}

func (ptr *GameController) DisconnectSessionAuthenticationError() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectSessionAuthenticationError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sessionAuthenticationError")
	}
}

func (ptr *GameController) SessionAuthenticationError() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_SessionAuthenticationError(ptr.Pointer())
	}
}

//export callbackGameControllercbfc99_CreateRoomError
func callbackGameControllercbfc99_CreateRoomError(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "createRoomError"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *GameController) ConnectCreateRoomError(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "createRoomError") {
			C.GameControllercbfc99_ConnectCreateRoomError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "createRoomError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createRoomError", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createRoomError", f)
		}
	}
}

func (ptr *GameController) DisconnectCreateRoomError() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectCreateRoomError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "createRoomError")
	}
}

func (ptr *GameController) CreateRoomError(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.GameControllercbfc99_CreateRoomError(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

//export callbackGameControllercbfc99_JoinedRoom
func callbackGameControllercbfc99_JoinedRoom(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "joinedRoom"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *GameController) ConnectJoinedRoom(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "joinedRoom") {
			C.GameControllercbfc99_ConnectJoinedRoom(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "joinedRoom"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "joinedRoom", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "joinedRoom", f)
		}
	}
}

func (ptr *GameController) DisconnectJoinedRoom() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectJoinedRoom(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "joinedRoom")
	}
}

func (ptr *GameController) JoinedRoom(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.GameControllercbfc99_JoinedRoom(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

//export callbackGameControllercbfc99_JoinRoomError
func callbackGameControllercbfc99_JoinRoomError(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "joinRoomError"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *GameController) ConnectJoinRoomError(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "joinRoomError") {
			C.GameControllercbfc99_ConnectJoinRoomError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "joinRoomError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "joinRoomError", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "joinRoomError", f)
		}
	}
}

func (ptr *GameController) DisconnectJoinRoomError() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectJoinRoomError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "joinRoomError")
	}
}

func (ptr *GameController) JoinRoomError(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.GameControllercbfc99_JoinRoomError(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

//export callbackGameControllercbfc99_PreparingGame
func callbackGameControllercbfc99_PreparingGame(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "preparingGame"); signal != nil {
		signal.(func())()
	}

}

func (ptr *GameController) ConnectPreparingGame(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "preparingGame") {
			C.GameControllercbfc99_ConnectPreparingGame(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "preparingGame"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "preparingGame", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preparingGame", f)
		}
	}
}

func (ptr *GameController) DisconnectPreparingGame() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectPreparingGame(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "preparingGame")
	}
}

func (ptr *GameController) PreparingGame() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_PreparingGame(ptr.Pointer())
	}
}

//export callbackGameControllercbfc99_ReadyGame
func callbackGameControllercbfc99_ReadyGame(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "readyGame"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *GameController) ConnectReadyGame(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "readyGame") {
			C.GameControllercbfc99_ConnectReadyGame(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "readyGame"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "readyGame", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readyGame", f)
		}
	}
}

func (ptr *GameController) DisconnectReadyGame() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectReadyGame(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "readyGame")
	}
}

func (ptr *GameController) ReadyGame(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.GameControllercbfc99_ReadyGame(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

//export callbackGameControllercbfc99_FinishedGame
func callbackGameControllercbfc99_FinishedGame(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "finishedGame"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *GameController) ConnectFinishedGame(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finishedGame") {
			C.GameControllercbfc99_ConnectFinishedGame(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finishedGame"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finishedGame", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finishedGame", f)
		}
	}
}

func (ptr *GameController) DisconnectFinishedGame() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectFinishedGame(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finishedGame")
	}
}

func (ptr *GameController) FinishedGame(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.GameControllercbfc99_FinishedGame(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

func GameController_QRegisterMetaType() int {
	return int(int32(C.GameControllercbfc99_GameControllercbfc99_QRegisterMetaType()))
}

func (ptr *GameController) QRegisterMetaType() int {
	return int(int32(C.GameControllercbfc99_GameControllercbfc99_QRegisterMetaType()))
}

func GameController_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.GameControllercbfc99_GameControllercbfc99_QRegisterMetaType2(typeNameC)))
}

func (ptr *GameController) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.GameControllercbfc99_GameControllercbfc99_QRegisterMetaType2(typeNameC)))
}

func GameController_QmlRegisterType() int {
	return int(int32(C.GameControllercbfc99_GameControllercbfc99_QmlRegisterType()))
}

func (ptr *GameController) QmlRegisterType() int {
	return int(int32(C.GameControllercbfc99_GameControllercbfc99_QmlRegisterType()))
}

func GameController_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.GameControllercbfc99_GameControllercbfc99_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *GameController) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.GameControllercbfc99_GameControllercbfc99_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *GameController) __children_atList(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.GameControllercbfc99___children_atList(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *GameController) __children_setList(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.GameControllercbfc99___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *GameController) __children_newList() unsafe.Pointer {
	return C.GameControllercbfc99___children_newList(ptr.Pointer())
}

func (ptr *GameController) __dynamicPropertyNames_atList(i int, p unsafe.Pointer) *std_core.QByteArray {
	tmpValue := std_core.NewQByteArrayFromPointer(C.GameControllercbfc99___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i)), p))
	runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *GameController) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF, p unsafe.Pointer) {
	C.GameControllercbfc99___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i), p)
}

func (ptr *GameController) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.GameControllercbfc99___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *GameController) __findChildren_atList2(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.GameControllercbfc99___findChildren_atList2(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *GameController) __findChildren_setList2(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.GameControllercbfc99___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *GameController) __findChildren_newList2() unsafe.Pointer {
	return C.GameControllercbfc99___findChildren_newList2(ptr.Pointer())
}

func (ptr *GameController) __findChildren_atList3(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.GameControllercbfc99___findChildren_atList3(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *GameController) __findChildren_setList3(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.GameControllercbfc99___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *GameController) __findChildren_newList3() unsafe.Pointer {
	return C.GameControllercbfc99___findChildren_newList3(ptr.Pointer())
}

func (ptr *GameController) __findChildren_atList(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.GameControllercbfc99___findChildren_atList(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *GameController) __findChildren_setList(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.GameControllercbfc99___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *GameController) __findChildren_newList() unsafe.Pointer {
	return C.GameControllercbfc99___findChildren_newList(ptr.Pointer())
}

func NewGameController(parent std_core.QObject_ITF) *GameController {
	tmpValue := NewGameControllerFromPointer(C.GameControllercbfc99_NewGameController(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackGameControllercbfc99_DestroyGameController
func callbackGameControllercbfc99_DestroyGameController(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~GameController"); signal != nil {
		signal.(func())()
	} else {
		NewGameControllerFromPointer(ptr).DestroyGameControllerDefault()
	}
}

func (ptr *GameController) ConnectDestroyGameController(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~GameController"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~GameController", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~GameController", f)
		}
	}
}

func (ptr *GameController) DisconnectDestroyGameController() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~GameController")
	}
}

func (ptr *GameController) DestroyGameController() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DestroyGameController(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *GameController) DestroyGameControllerDefault() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DestroyGameControllerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackGameControllercbfc99_TimerEvent
func callbackGameControllercbfc99_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewGameControllerFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *GameController) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

//export callbackGameControllercbfc99_ChildEvent
func callbackGameControllercbfc99_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewGameControllerFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *GameController) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackGameControllercbfc99_ConnectNotify
func callbackGameControllercbfc99_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewGameControllerFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *GameController) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackGameControllercbfc99_CustomEvent
func callbackGameControllercbfc99_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewGameControllerFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *GameController) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackGameControllercbfc99_DeleteLater
func callbackGameControllercbfc99_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewGameControllerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *GameController) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackGameControllercbfc99_Destroyed
func callbackGameControllercbfc99_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackGameControllercbfc99_DisconnectNotify
func callbackGameControllercbfc99_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewGameControllerFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *GameController) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.GameControllercbfc99_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackGameControllercbfc99_Event
func callbackGameControllercbfc99_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGameControllerFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *GameController) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.GameControllercbfc99_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackGameControllercbfc99_EventFilter
func callbackGameControllercbfc99_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGameControllerFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *GameController) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.GameControllercbfc99_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackGameControllercbfc99_ObjectNameChanged
func callbackGameControllercbfc99_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

type User_ITF interface {
	std_core.QObject_ITF
	User_PTR() *User
}

func (ptr *User) User_PTR() *User {
	return ptr
}

func (ptr *User) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *User) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromUser(ptr User_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.User_PTR().Pointer()
	}
	return nil
}

func NewUserFromPointer(ptr unsafe.Pointer) (n *User) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(User)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *User:
			n = deduced

		case *std_core.QObject:
			n = &User{QObject: *deduced}

		default:
			n = new(User)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackUsercbfc99_Constructor
func callbackUsercbfc99_Constructor(ptr unsafe.Pointer) {
	this := NewUserFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackUsercbfc99_State
func callbackUsercbfc99_State(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "state"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewUserFromPointer(ptr).StateDefault()))
}

func (ptr *User) ConnectState(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "state"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "state", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "state", f)
		}
	}
}

func (ptr *User) DisconnectState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "state")
	}
}

func (ptr *User) State() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Usercbfc99_State(ptr.Pointer())))
	}
	return 0
}

func (ptr *User) StateDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Usercbfc99_StateDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackUsercbfc99_SetState
func callbackUsercbfc99_SetState(ptr unsafe.Pointer, state C.int) {
	if signal := qt.GetSignal(ptr, "setState"); signal != nil {
		signal.(func(int))(int(int32(state)))
	} else {
		NewUserFromPointer(ptr).SetStateDefault(int(int32(state)))
	}
}

func (ptr *User) ConnectSetState(f func(state int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setState"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setState", func(state int) {
				signal.(func(int))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setState", f)
		}
	}
}

func (ptr *User) DisconnectSetState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setState")
	}
}

func (ptr *User) SetState(state int) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_SetState(ptr.Pointer(), C.int(int32(state)))
	}
}

func (ptr *User) SetStateDefault(state int) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_SetStateDefault(ptr.Pointer(), C.int(int32(state)))
	}
}

//export callbackUsercbfc99_StateChanged
func callbackUsercbfc99_StateChanged(ptr unsafe.Pointer, state C.int) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(int))(int(int32(state)))
	}

}

func (ptr *User) ConnectStateChanged(f func(state int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.Usercbfc99_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(state int) {
				signal.(func(int))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
		}
	}
}

func (ptr *User) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.Usercbfc99_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *User) StateChanged(state int) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_StateChanged(ptr.Pointer(), C.int(int32(state)))
	}
}

//export callbackUsercbfc99_Data
func callbackUsercbfc99_Data(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQJsonObject(signal.(func() *std_core.QJsonObject)())
	}

	return std_core.PointerFromQJsonObject(NewUserFromPointer(ptr).DataDefault())
}

func (ptr *User) ConnectData(f func() *std_core.QJsonObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "data", func() *std_core.QJsonObject {
				signal.(func() *std_core.QJsonObject)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", f)
		}
	}
}

func (ptr *User) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *User) Data() *std_core.QJsonObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQJsonObjectFromPointer(C.Usercbfc99_Data(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QJsonObject).DestroyQJsonObject)
		return tmpValue
	}
	return nil
}

func (ptr *User) DataDefault() *std_core.QJsonObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQJsonObjectFromPointer(C.Usercbfc99_DataDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QJsonObject).DestroyQJsonObject)
		return tmpValue
	}
	return nil
}

//export callbackUsercbfc99_SetData
func callbackUsercbfc99_SetData(ptr unsafe.Pointer, data unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		signal.(func(*std_core.QJsonObject))(std_core.NewQJsonObjectFromPointer(data))
	} else {
		NewUserFromPointer(ptr).SetDataDefault(std_core.NewQJsonObjectFromPointer(data))
	}
}

func (ptr *User) ConnectSetData(f func(data *std_core.QJsonObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setData", func(data *std_core.QJsonObject) {
				signal.(func(*std_core.QJsonObject))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setData", f)
		}
	}
}

func (ptr *User) DisconnectSetData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setData")
	}
}

func (ptr *User) SetData(data std_core.QJsonObject_ITF) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_SetData(ptr.Pointer(), std_core.PointerFromQJsonObject(data))
	}
}

func (ptr *User) SetDataDefault(data std_core.QJsonObject_ITF) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_SetDataDefault(ptr.Pointer(), std_core.PointerFromQJsonObject(data))
	}
}

//export callbackUsercbfc99_DataChanged
func callbackUsercbfc99_DataChanged(ptr unsafe.Pointer, data unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QJsonObject))(std_core.NewQJsonObjectFromPointer(data))
	}

}

func (ptr *User) ConnectDataChanged(f func(data *std_core.QJsonObject)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataChanged") {
			C.Usercbfc99_ConnectDataChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dataChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dataChanged", func(data *std_core.QJsonObject) {
				signal.(func(*std_core.QJsonObject))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dataChanged", f)
		}
	}
}

func (ptr *User) DisconnectDataChanged() {
	if ptr.Pointer() != nil {
		C.Usercbfc99_DisconnectDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dataChanged")
	}
}

func (ptr *User) DataChanged(data std_core.QJsonObject_ITF) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_DataChanged(ptr.Pointer(), std_core.PointerFromQJsonObject(data))
	}
}

func User_QRegisterMetaType() int {
	return int(int32(C.Usercbfc99_Usercbfc99_QRegisterMetaType()))
}

func (ptr *User) QRegisterMetaType() int {
	return int(int32(C.Usercbfc99_Usercbfc99_QRegisterMetaType()))
}

func User_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Usercbfc99_Usercbfc99_QRegisterMetaType2(typeNameC)))
}

func (ptr *User) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Usercbfc99_Usercbfc99_QRegisterMetaType2(typeNameC)))
}

func User_QmlRegisterType() int {
	return int(int32(C.Usercbfc99_Usercbfc99_QmlRegisterType()))
}

func (ptr *User) QmlRegisterType() int {
	return int(int32(C.Usercbfc99_Usercbfc99_QmlRegisterType()))
}

func User_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Usercbfc99_Usercbfc99_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *User) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Usercbfc99_Usercbfc99_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *User) __children_atList(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Usercbfc99___children_atList(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *User) __children_setList(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Usercbfc99___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *User) __children_newList() unsafe.Pointer {
	return C.Usercbfc99___children_newList(ptr.Pointer())
}

func (ptr *User) __dynamicPropertyNames_atList(i int, p unsafe.Pointer) *std_core.QByteArray {
	tmpValue := std_core.NewQByteArrayFromPointer(C.Usercbfc99___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i)), p))
	runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *User) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF, p unsafe.Pointer) {
	C.Usercbfc99___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i), p)
}

func (ptr *User) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Usercbfc99___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *User) __findChildren_atList2(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Usercbfc99___findChildren_atList2(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *User) __findChildren_setList2(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Usercbfc99___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *User) __findChildren_newList2() unsafe.Pointer {
	return C.Usercbfc99___findChildren_newList2(ptr.Pointer())
}

func (ptr *User) __findChildren_atList3(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Usercbfc99___findChildren_atList3(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *User) __findChildren_setList3(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Usercbfc99___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *User) __findChildren_newList3() unsafe.Pointer {
	return C.Usercbfc99___findChildren_newList3(ptr.Pointer())
}

func (ptr *User) __findChildren_atList(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Usercbfc99___findChildren_atList(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *User) __findChildren_setList(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Usercbfc99___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *User) __findChildren_newList() unsafe.Pointer {
	return C.Usercbfc99___findChildren_newList(ptr.Pointer())
}

func NewUser(parent std_core.QObject_ITF) *User {
	tmpValue := NewUserFromPointer(C.Usercbfc99_NewUser(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackUsercbfc99_DestroyUser
func callbackUsercbfc99_DestroyUser(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~User"); signal != nil {
		signal.(func())()
	} else {
		NewUserFromPointer(ptr).DestroyUserDefault()
	}
}

func (ptr *User) ConnectDestroyUser(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~User"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~User", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~User", f)
		}
	}
}

func (ptr *User) DisconnectDestroyUser() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~User")
	}
}

func (ptr *User) DestroyUser() {
	if ptr.Pointer() != nil {
		C.Usercbfc99_DestroyUser(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *User) DestroyUserDefault() {
	if ptr.Pointer() != nil {
		C.Usercbfc99_DestroyUserDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackUsercbfc99_TimerEvent
func callbackUsercbfc99_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewUserFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *User) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

//export callbackUsercbfc99_ChildEvent
func callbackUsercbfc99_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewUserFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *User) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackUsercbfc99_ConnectNotify
func callbackUsercbfc99_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewUserFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *User) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackUsercbfc99_CustomEvent
func callbackUsercbfc99_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewUserFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *User) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackUsercbfc99_DeleteLater
func callbackUsercbfc99_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewUserFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *User) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Usercbfc99_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackUsercbfc99_Destroyed
func callbackUsercbfc99_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackUsercbfc99_DisconnectNotify
func callbackUsercbfc99_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewUserFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *User) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Usercbfc99_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackUsercbfc99_Event
func callbackUsercbfc99_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewUserFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *User) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Usercbfc99_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackUsercbfc99_EventFilter
func callbackUsercbfc99_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewUserFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *User) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Usercbfc99_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackUsercbfc99_ObjectNameChanged
func callbackUsercbfc99_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}
