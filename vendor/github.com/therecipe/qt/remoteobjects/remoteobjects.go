//go:build !minimal
// +build !minimal

package remoteobjects

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "remoteobjects.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtRemoteObjects_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtRemoteObjects_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QRemoteObjectAbstractPersistedStore struct {
	core.QObject
}

type QRemoteObjectAbstractPersistedStore_ITF interface {
	core.QObject_ITF
	QRemoteObjectAbstractPersistedStore_PTR() *QRemoteObjectAbstractPersistedStore
}

func (ptr *QRemoteObjectAbstractPersistedStore) QRemoteObjectAbstractPersistedStore_PTR() *QRemoteObjectAbstractPersistedStore {
	return ptr
}

func (ptr *QRemoteObjectAbstractPersistedStore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectAbstractPersistedStore(ptr QRemoteObjectAbstractPersistedStore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectAbstractPersistedStore_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectAbstractPersistedStore) {
	n = new(QRemoteObjectAbstractPersistedStore)
	n.SetPointer(ptr)
	return
}
func NewQRemoteObjectAbstractPersistedStore(parent core.QObject_ITF) *QRemoteObjectAbstractPersistedStore {
	tmpValue := NewQRemoteObjectAbstractPersistedStoreFromPointer(C.QRemoteObjectAbstractPersistedStore_NewQRemoteObjectAbstractPersistedStore(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQRemoteObjectAbstractPersistedStore_RestoreProperties
func callbackQRemoteObjectAbstractPersistedStore_RestoreProperties(ptr unsafe.Pointer, repName C.struct_QtRemoteObjects_PackedString, repSig unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "restoreProperties"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQRemoteObjectAbstractPersistedStoreFromPointer(NewQRemoteObjectAbstractPersistedStoreFromPointer(nil).__restoreProperties_newList())
			for _, v := range (*(*func(string, *core.QByteArray) []*core.QVariant)(signal))(cGoUnpackString(repName), core.NewQByteArrayFromPointer(repSig)) {
				tmpList.__restoreProperties_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQRemoteObjectAbstractPersistedStoreFromPointer(NewQRemoteObjectAbstractPersistedStoreFromPointer(nil).__restoreProperties_newList())
		for _, v := range make([]*core.QVariant, 0) {
			tmpList.__restoreProperties_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QRemoteObjectAbstractPersistedStore) ConnectRestoreProperties(f func(repName string, repSig *core.QByteArray) []*core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "restoreProperties"); signal != nil {
			f := func(repName string, repSig *core.QByteArray) []*core.QVariant {
				(*(*func(string, *core.QByteArray) []*core.QVariant)(signal))(repName, repSig)
				return f(repName, repSig)
			}
			qt.ConnectSignal(ptr.Pointer(), "restoreProperties", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "restoreProperties", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DisconnectRestoreProperties() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "restoreProperties")
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) RestoreProperties(repName string, repSig core.QByteArray_ITF) []*core.QVariant {
	if ptr.Pointer() != nil {
		var repNameC *C.char
		if repName != "" {
			repNameC = C.CString(repName)
			defer C.free(unsafe.Pointer(repNameC))
		}
		return func(l C.struct_QtRemoteObjects_PackedList) []*core.QVariant {
			out := make([]*core.QVariant, int(l.len))
			tmpList := NewQRemoteObjectAbstractPersistedStoreFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__restoreProperties_atList(i)
			}
			return out
		}(C.QRemoteObjectAbstractPersistedStore_RestoreProperties(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: repNameC, len: C.longlong(len(repName))}, core.PointerFromQByteArray(repSig)))
	}
	return make([]*core.QVariant, 0)
}

//export callbackQRemoteObjectAbstractPersistedStore_SaveProperties
func callbackQRemoteObjectAbstractPersistedStore_SaveProperties(ptr unsafe.Pointer, repName C.struct_QtRemoteObjects_PackedString, repSig unsafe.Pointer, values C.struct_QtRemoteObjects_PackedList) {
	if signal := qt.GetSignal(ptr, "saveProperties"); signal != nil {
		(*(*func(string, *core.QByteArray, []*core.QVariant))(signal))(cGoUnpackString(repName), core.NewQByteArrayFromPointer(repSig), func(l C.struct_QtRemoteObjects_PackedList) []*core.QVariant {
			out := make([]*core.QVariant, int(l.len))
			tmpList := NewQRemoteObjectAbstractPersistedStoreFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__saveProperties_values_atList(i)
			}
			return out
		}(values))
	}

}

func (ptr *QRemoteObjectAbstractPersistedStore) ConnectSaveProperties(f func(repName string, repSig *core.QByteArray, values []*core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "saveProperties"); signal != nil {
			f := func(repName string, repSig *core.QByteArray, values []*core.QVariant) {
				(*(*func(string, *core.QByteArray, []*core.QVariant))(signal))(repName, repSig, values)
				f(repName, repSig, values)
			}
			qt.ConnectSignal(ptr.Pointer(), "saveProperties", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "saveProperties", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DisconnectSaveProperties() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "saveProperties")
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) SaveProperties(repName string, repSig core.QByteArray_ITF, values []*core.QVariant) {
	if ptr.Pointer() != nil {
		var repNameC *C.char
		if repName != "" {
			repNameC = C.CString(repName)
			defer C.free(unsafe.Pointer(repNameC))
		}
		C.QRemoteObjectAbstractPersistedStore_SaveProperties(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: repNameC, len: C.longlong(len(repName))}, core.PointerFromQByteArray(repSig), func() unsafe.Pointer {
			tmpList := NewQRemoteObjectAbstractPersistedStoreFromPointer(NewQRemoteObjectAbstractPersistedStoreFromPointer(nil).__saveProperties_values_newList())
			for _, v := range values {
				tmpList.__saveProperties_values_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __restoreProperties_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectAbstractPersistedStore___restoreProperties_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __restoreProperties_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___restoreProperties_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __restoreProperties_newList() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___restoreProperties_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __saveProperties_values_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectAbstractPersistedStore___saveProperties_values_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __saveProperties_values_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___saveProperties_values_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __saveProperties_values_newList() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___saveProperties_values_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectAbstractPersistedStore___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __children_newList() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___children_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectAbstractPersistedStore___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_newList() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___findChildren_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectAbstractPersistedStore___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_newList3() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___findChildren_newList3(ptr.Pointer())
}

//export callbackQRemoteObjectAbstractPersistedStore_ChildEvent
func callbackQRemoteObjectAbstractPersistedStore_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_ConnectNotify
func callbackQRemoteObjectAbstractPersistedStore_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_CustomEvent
func callbackQRemoteObjectAbstractPersistedStore_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_DeleteLater
func callbackQRemoteObjectAbstractPersistedStore_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRemoteObjectAbstractPersistedStore_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_Destroyed
func callbackQRemoteObjectAbstractPersistedStore_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQRemoteObjectAbstractPersistedStore_DisconnectNotify
func callbackQRemoteObjectAbstractPersistedStore_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_Event
func callbackQRemoteObjectAbstractPersistedStore_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QRemoteObjectAbstractPersistedStore) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectAbstractPersistedStore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQRemoteObjectAbstractPersistedStore_EventFilter
func callbackQRemoteObjectAbstractPersistedStore_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QRemoteObjectAbstractPersistedStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectAbstractPersistedStore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQRemoteObjectAbstractPersistedStore_MetaObject
func callbackQRemoteObjectAbstractPersistedStore_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QRemoteObjectAbstractPersistedStore) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRemoteObjectAbstractPersistedStore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQRemoteObjectAbstractPersistedStore_ObjectNameChanged
func callbackQRemoteObjectAbstractPersistedStore_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtRemoteObjects_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQRemoteObjectAbstractPersistedStore_TimerEvent
func callbackQRemoteObjectAbstractPersistedStore_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QRemoteObjectDynamicReplica struct {
	QRemoteObjectReplica
}

type QRemoteObjectDynamicReplica_ITF interface {
	QRemoteObjectReplica_ITF
	QRemoteObjectDynamicReplica_PTR() *QRemoteObjectDynamicReplica
}

func (ptr *QRemoteObjectDynamicReplica) QRemoteObjectDynamicReplica_PTR() *QRemoteObjectDynamicReplica {
	return ptr
}

func (ptr *QRemoteObjectDynamicReplica) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectReplica_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectDynamicReplica) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectReplica_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectDynamicReplica(ptr QRemoteObjectDynamicReplica_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectDynamicReplica_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectDynamicReplicaFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectDynamicReplica) {
	n = new(QRemoteObjectDynamicReplica)
	n.SetPointer(ptr)
	return
}

//export callbackQRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica
func callbackQRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QRemoteObjectDynamicReplica"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQRemoteObjectDynamicReplicaFromPointer(ptr).DestroyQRemoteObjectDynamicReplicaDefault()
	}
}

func (ptr *QRemoteObjectDynamicReplica) ConnectDestroyQRemoteObjectDynamicReplica(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QRemoteObjectDynamicReplica"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectDynamicReplica", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectDynamicReplica", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectDynamicReplica) DisconnectDestroyQRemoteObjectDynamicReplica() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QRemoteObjectDynamicReplica")
	}
}

func (ptr *QRemoteObjectDynamicReplica) DestroyQRemoteObjectDynamicReplica() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRemoteObjectDynamicReplica) DestroyQRemoteObjectDynamicReplicaDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplicaDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QRemoteObjectHost struct {
	QRemoteObjectHostBase
}

type QRemoteObjectHost_ITF interface {
	QRemoteObjectHostBase_ITF
	QRemoteObjectHost_PTR() *QRemoteObjectHost
}

func (ptr *QRemoteObjectHost) QRemoteObjectHost_PTR() *QRemoteObjectHost {
	return ptr
}

func (ptr *QRemoteObjectHost) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHostBase_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectHost) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectHostBase_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectHost(ptr QRemoteObjectHost_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHost_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectHostFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectHost) {
	n = new(QRemoteObjectHost)
	n.SetPointer(ptr)
	return
}
func NewQRemoteObjectHost(parent core.QObject_ITF) *QRemoteObjectHost {
	tmpValue := NewQRemoteObjectHostFromPointer(C.QRemoteObjectHost_NewQRemoteObjectHost(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQRemoteObjectHost2(address core.QUrl_ITF, registryAddress core.QUrl_ITF, allowedSchemas QRemoteObjectHostBase__AllowedSchemas, parent core.QObject_ITF) *QRemoteObjectHost {
	tmpValue := NewQRemoteObjectHostFromPointer(C.QRemoteObjectHost_NewQRemoteObjectHost2(core.PointerFromQUrl(address), core.PointerFromQUrl(registryAddress), C.longlong(allowedSchemas), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQRemoteObjectHost3(address core.QUrl_ITF, parent core.QObject_ITF) *QRemoteObjectHost {
	tmpValue := NewQRemoteObjectHostFromPointer(C.QRemoteObjectHost_NewQRemoteObjectHost3(core.PointerFromQUrl(address), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQRemoteObjectHost_HostUrl
func callbackQRemoteObjectHost_HostUrl(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "hostUrl"); signal != nil {
		return core.PointerFromQUrl((*(*func() *core.QUrl)(signal))())
	}

	return core.PointerFromQUrl(NewQRemoteObjectHostFromPointer(ptr).HostUrlDefault())
}

func (ptr *QRemoteObjectHost) ConnectHostUrl(f func() *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hostUrl"); signal != nil {
			f := func() *core.QUrl {
				(*(*func() *core.QUrl)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hostUrl", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hostUrl", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectHost) DisconnectHostUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hostUrl")
	}
}

func (ptr *QRemoteObjectHost) HostUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QRemoteObjectHost_HostUrl(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectHost) HostUrlDefault() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QRemoteObjectHost_HostUrlDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQRemoteObjectHost_HostUrlChanged
func callbackQRemoteObjectHost_HostUrlChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hostUrlChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QRemoteObjectHost) ConnectHostUrlChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hostUrlChanged") {
			C.QRemoteObjectHost_ConnectHostUrlChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hostUrlChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hostUrlChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hostUrlChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hostUrlChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectHost) DisconnectHostUrlChanged() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectHost_DisconnectHostUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hostUrlChanged")
	}
}

func (ptr *QRemoteObjectHost) HostUrlChanged() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectHost_HostUrlChanged(ptr.Pointer())
	}
}

//export callbackQRemoteObjectHost_SetHostUrl
func callbackQRemoteObjectHost_SetHostUrl(ptr unsafe.Pointer, hostAddress unsafe.Pointer, allowedSchemas C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "setHostUrl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QUrl, QRemoteObjectHostBase__AllowedSchemas) bool)(signal))(core.NewQUrlFromPointer(hostAddress), QRemoteObjectHostBase__AllowedSchemas(allowedSchemas)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectHostFromPointer(ptr).SetHostUrlDefault(core.NewQUrlFromPointer(hostAddress), QRemoteObjectHostBase__AllowedSchemas(allowedSchemas)))))
}

func (ptr *QRemoteObjectHost) ConnectSetHostUrl(f func(hostAddress *core.QUrl, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHostUrl"); signal != nil {
			f := func(hostAddress *core.QUrl, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool {
				(*(*func(*core.QUrl, QRemoteObjectHostBase__AllowedSchemas) bool)(signal))(hostAddress, allowedSchemas)
				return f(hostAddress, allowedSchemas)
			}
			qt.ConnectSignal(ptr.Pointer(), "setHostUrl", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHostUrl", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectHost) DisconnectSetHostUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHostUrl")
	}
}

func (ptr *QRemoteObjectHost) SetHostUrl(hostAddress core.QUrl_ITF, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectHost_SetHostUrl(ptr.Pointer(), core.PointerFromQUrl(hostAddress), C.longlong(allowedSchemas))) != 0
	}
	return false
}

func (ptr *QRemoteObjectHost) SetHostUrlDefault(hostAddress core.QUrl_ITF, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectHost_SetHostUrlDefault(ptr.Pointer(), core.PointerFromQUrl(hostAddress), C.longlong(allowedSchemas))) != 0
	}
	return false
}

type QRemoteObjectHostBase struct {
	QRemoteObjectNode
}

type QRemoteObjectHostBase_ITF interface {
	QRemoteObjectNode_ITF
	QRemoteObjectHostBase_PTR() *QRemoteObjectHostBase
}

func (ptr *QRemoteObjectHostBase) QRemoteObjectHostBase_PTR() *QRemoteObjectHostBase {
	return ptr
}

func (ptr *QRemoteObjectHostBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectHostBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectNode_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectHostBase(ptr QRemoteObjectHostBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHostBase_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectHostBaseFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectHostBase) {
	n = new(QRemoteObjectHostBase)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QRemoteObjectHostBase__AllowedSchemas
//QRemoteObjectHostBase::AllowedSchemas
type QRemoteObjectHostBase__AllowedSchemas int64

const (
	QRemoteObjectHostBase__BuiltInSchemasOnly        QRemoteObjectHostBase__AllowedSchemas = QRemoteObjectHostBase__AllowedSchemas(0)
	QRemoteObjectHostBase__AllowExternalRegistration QRemoteObjectHostBase__AllowedSchemas = QRemoteObjectHostBase__AllowedSchemas(1)
)

func (ptr *QRemoteObjectHostBase) AddHostSideConnection(ioDevice core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectHostBase_AddHostSideConnection(ptr.Pointer(), core.PointerFromQIODevice(ioDevice))
	}
}

func (ptr *QRemoteObjectHostBase) DisableRemoting(remoteObject core.QObject_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectHostBase_DisableRemoting(ptr.Pointer(), core.PointerFromQObject(remoteObject))) != 0
	}
	return false
}

func (ptr *QRemoteObjectHostBase) EnableRemoting2(object core.QObject_ITF, name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QRemoteObjectHostBase_EnableRemoting2(ptr.Pointer(), core.PointerFromQObject(object), C.struct_QtRemoteObjects_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QRemoteObjectHostBase) EnableRemoting3(model core.QAbstractItemModel_ITF, name string, roles []int, selectionModel core.QItemSelectionModel_ITF) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QRemoteObjectHostBase_EnableRemoting3(ptr.Pointer(), core.PointerFromQAbstractItemModel(model), C.struct_QtRemoteObjects_PackedString{data: nameC, len: C.longlong(len(name))}, func() unsafe.Pointer {
			tmpList := NewQRemoteObjectHostBaseFromPointer(NewQRemoteObjectHostBaseFromPointer(nil).__enableRemoting_roles_newList3())
			for _, v := range roles {
				tmpList.__enableRemoting_roles_setList3(v)
			}
			return tmpList.Pointer()
		}(), core.PointerFromQItemSelectionModel(selectionModel))) != 0
	}
	return false
}

func (ptr *QRemoteObjectHostBase) __enableRemoting_roles_atList3(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRemoteObjectHostBase___enableRemoting_roles_atList3(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QRemoteObjectHostBase) __enableRemoting_roles_setList3(i int) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectHostBase___enableRemoting_roles_setList3(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QRemoteObjectHostBase) __enableRemoting_roles_newList3() unsafe.Pointer {
	return C.QRemoteObjectHostBase___enableRemoting_roles_newList3(ptr.Pointer())
}

type QRemoteObjectNode struct {
	core.QObject
}

type QRemoteObjectNode_ITF interface {
	core.QObject_ITF
	QRemoteObjectNode_PTR() *QRemoteObjectNode
}

func (ptr *QRemoteObjectNode) QRemoteObjectNode_PTR() *QRemoteObjectNode {
	return ptr
}

func (ptr *QRemoteObjectNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectNode(ptr QRemoteObjectNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectNode_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectNodeFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectNode) {
	n = new(QRemoteObjectNode)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QRemoteObjectNode__ErrorCode
//QRemoteObjectNode::ErrorCode
type QRemoteObjectNode__ErrorCode int64

const (
	QRemoteObjectNode__NoError                       QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(0)
	QRemoteObjectNode__RegistryNotAcquired           QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(1)
	QRemoteObjectNode__RegistryAlreadyHosted         QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(2)
	QRemoteObjectNode__NodeIsNoServer                QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(3)
	QRemoteObjectNode__ServerAlreadyCreated          QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(4)
	QRemoteObjectNode__UnintendedRegistryHosting     QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(5)
	QRemoteObjectNode__OperationNotValidOnClientNode QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(6)
	QRemoteObjectNode__SourceNotRegistered           QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(7)
	QRemoteObjectNode__MissingObjectName             QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(8)
	QRemoteObjectNode__HostUrlInvalid                QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(9)
	QRemoteObjectNode__ProtocolMismatch              QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(10)
	QRemoteObjectNode__ListenFailed                  QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(11)
)

func NewQRemoteObjectNode(parent core.QObject_ITF) *QRemoteObjectNode {
	tmpValue := NewQRemoteObjectNodeFromPointer(C.QRemoteObjectNode_NewQRemoteObjectNode(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQRemoteObjectNode2(registryAddress core.QUrl_ITF, parent core.QObject_ITF) *QRemoteObjectNode {
	tmpValue := NewQRemoteObjectNodeFromPointer(C.QRemoteObjectNode_NewQRemoteObjectNode2(core.PointerFromQUrl(registryAddress), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QRemoteObjectNode) AcquireDynamic(name string) *QRemoteObjectDynamicReplica {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQRemoteObjectDynamicReplicaFromPointer(C.QRemoteObjectNode_AcquireDynamic(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: nameC, len: C.longlong(len(name))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) AddClientSideConnection(ioDevice core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_AddClientSideConnection(ptr.Pointer(), core.PointerFromQIODevice(ioDevice))
	}
}

func (ptr *QRemoteObjectNode) ConnectToNode(address core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_ConnectToNode(ptr.Pointer(), core.PointerFromQUrl(address))) != 0
	}
	return false
}

func (ptr *QRemoteObjectNode) HeartbeatInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRemoteObjectNode_HeartbeatInterval(ptr.Pointer())))
	}
	return 0
}

//export callbackQRemoteObjectNode_HeartbeatIntervalChanged
func callbackQRemoteObjectNode_HeartbeatIntervalChanged(ptr unsafe.Pointer, heartbeatInterval C.int) {
	if signal := qt.GetSignal(ptr, "heartbeatIntervalChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(heartbeatInterval)))
	}

}

func (ptr *QRemoteObjectNode) ConnectHeartbeatIntervalChanged(f func(heartbeatInterval int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "heartbeatIntervalChanged") {
			C.QRemoteObjectNode_ConnectHeartbeatIntervalChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "heartbeatIntervalChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "heartbeatIntervalChanged"); signal != nil {
			f := func(heartbeatInterval int) {
				(*(*func(int))(signal))(heartbeatInterval)
				f(heartbeatInterval)
			}
			qt.ConnectSignal(ptr.Pointer(), "heartbeatIntervalChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "heartbeatIntervalChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectNode) DisconnectHeartbeatIntervalChanged() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_DisconnectHeartbeatIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "heartbeatIntervalChanged")
	}
}

func (ptr *QRemoteObjectNode) HeartbeatIntervalChanged(heartbeatInterval int) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_HeartbeatIntervalChanged(ptr.Pointer(), C.int(int32(heartbeatInterval)))
	}
}

func (ptr *QRemoteObjectNode) Instances2(typeName string) []string {
	if ptr.Pointer() != nil {
		var typeNameC *C.char
		if typeName != "" {
			typeNameC = C.CString(typeName)
			defer C.free(unsafe.Pointer(typeNameC))
		}
		return unpackStringList(cGoUnpackString(C.QRemoteObjectNode_Instances2(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: typeNameC, len: C.longlong(len(typeName))})))
	}
	return make([]string, 0)
}

func (ptr *QRemoteObjectNode) LastError() QRemoteObjectNode__ErrorCode {
	if ptr.Pointer() != nil {
		return QRemoteObjectNode__ErrorCode(C.QRemoteObjectNode_LastError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRemoteObjectNode) PersistedStore() *QRemoteObjectAbstractPersistedStore {
	if ptr.Pointer() != nil {
		tmpValue := NewQRemoteObjectAbstractPersistedStoreFromPointer(C.QRemoteObjectNode_PersistedStore(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) Registry() *QRemoteObjectRegistry {
	if ptr.Pointer() != nil {
		tmpValue := NewQRemoteObjectRegistryFromPointer(C.QRemoteObjectNode_Registry(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) RegistryUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QRemoteObjectNode_RegistryUrl(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) SetHeartbeatInterval(interval int) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_SetHeartbeatInterval(ptr.Pointer(), C.int(int32(interval)))
	}
}

//export callbackQRemoteObjectNode_SetName
func callbackQRemoteObjectNode_SetName(ptr unsafe.Pointer, name C.struct_QtRemoteObjects_PackedString) {
	if signal := qt.GetSignal(ptr, "setName"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(name))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).SetNameDefault(cGoUnpackString(name))
	}
}

func (ptr *QRemoteObjectNode) ConnectSetName(f func(name string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setName"); signal != nil {
			f := func(name string) {
				(*(*func(string))(signal))(name)
				f(name)
			}
			qt.ConnectSignal(ptr.Pointer(), "setName", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setName", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectNode) DisconnectSetName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setName")
	}
}

func (ptr *QRemoteObjectNode) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QRemoteObjectNode_SetName(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QRemoteObjectNode) SetNameDefault(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QRemoteObjectNode_SetNameDefault(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QRemoteObjectNode) SetPersistedStore(persistedStore QRemoteObjectAbstractPersistedStore_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_SetPersistedStore(ptr.Pointer(), PointerFromQRemoteObjectAbstractPersistedStore(persistedStore))
	}
}

//export callbackQRemoteObjectNode_SetRegistryUrl
func callbackQRemoteObjectNode_SetRegistryUrl(ptr unsafe.Pointer, registryAddress unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "setRegistryUrl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QUrl) bool)(signal))(core.NewQUrlFromPointer(registryAddress)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectNodeFromPointer(ptr).SetRegistryUrlDefault(core.NewQUrlFromPointer(registryAddress)))))
}

func (ptr *QRemoteObjectNode) ConnectSetRegistryUrl(f func(registryAddress *core.QUrl) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRegistryUrl"); signal != nil {
			f := func(registryAddress *core.QUrl) bool {
				(*(*func(*core.QUrl) bool)(signal))(registryAddress)
				return f(registryAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "setRegistryUrl", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRegistryUrl", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectNode) DisconnectSetRegistryUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRegistryUrl")
	}
}

func (ptr *QRemoteObjectNode) SetRegistryUrl(registryAddress core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_SetRegistryUrl(ptr.Pointer(), core.PointerFromQUrl(registryAddress))) != 0
	}
	return false
}

func (ptr *QRemoteObjectNode) SetRegistryUrlDefault(registryAddress core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_SetRegistryUrlDefault(ptr.Pointer(), core.PointerFromQUrl(registryAddress))) != 0
	}
	return false
}

//export callbackQRemoteObjectNode_TimerEvent
func callbackQRemoteObjectNode_TimerEvent(ptr unsafe.Pointer, vqt unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(vqt))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(vqt))
	}
}

func (ptr *QRemoteObjectNode) TimerEventDefault(vqt core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(vqt))
	}
}

func (ptr *QRemoteObjectNode) WaitForRegistry(timeout int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_WaitForRegistry(ptr.Pointer(), C.int(int32(timeout)))) != 0
	}
	return false
}

func (ptr *QRemoteObjectNode) __acquireModel_rolesHint_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRemoteObjectNode___acquireModel_rolesHint_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QRemoteObjectNode) __acquireModel_rolesHint_setList(i int) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___acquireModel_rolesHint_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QRemoteObjectNode) __acquireModel_rolesHint_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___acquireModel_rolesHint_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __persistProperties_props_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectNode___persistProperties_props_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __persistProperties_props_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___persistProperties_props_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectNode) __persistProperties_props_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___persistProperties_props_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __retrieveProperties_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectNode___retrieveProperties_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __retrieveProperties_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___retrieveProperties_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectNode) __retrieveProperties_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___retrieveProperties_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectNode___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectNode) __children_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___children_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QRemoteObjectNode___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QRemoteObjectNode) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectNode___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectNode) __findChildren_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___findChildren_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectNode___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectNode) __findChildren_newList3() unsafe.Pointer {
	return C.QRemoteObjectNode___findChildren_newList3(ptr.Pointer())
}

//export callbackQRemoteObjectNode_ChildEvent
func callbackQRemoteObjectNode_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectNode) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQRemoteObjectNode_ConnectNotify
func callbackQRemoteObjectNode_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectNode) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectNode_CustomEvent
func callbackQRemoteObjectNode_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectNode) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQRemoteObjectNode_DeleteLater
func callbackQRemoteObjectNode_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRemoteObjectNode) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRemoteObjectNode_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQRemoteObjectNode_Destroyed
func callbackQRemoteObjectNode_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQRemoteObjectNode_DisconnectNotify
func callbackQRemoteObjectNode_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectNode) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectNode_Event
func callbackQRemoteObjectNode_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectNodeFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QRemoteObjectNode) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQRemoteObjectNode_EventFilter
func callbackQRemoteObjectNode_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectNodeFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QRemoteObjectNode) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQRemoteObjectNode_MetaObject
func callbackQRemoteObjectNode_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQRemoteObjectNodeFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QRemoteObjectNode) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRemoteObjectNode_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQRemoteObjectNode_ObjectNameChanged
func callbackQRemoteObjectNode_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtRemoteObjects_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

type QRemoteObjectPendingCall struct {
	ptr unsafe.Pointer
}

type QRemoteObjectPendingCall_ITF interface {
	QRemoteObjectPendingCall_PTR() *QRemoteObjectPendingCall
}

func (ptr *QRemoteObjectPendingCall) QRemoteObjectPendingCall_PTR() *QRemoteObjectPendingCall {
	return ptr
}

func (ptr *QRemoteObjectPendingCall) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRemoteObjectPendingCall) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRemoteObjectPendingCall(ptr QRemoteObjectPendingCall_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectPendingCall_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectPendingCallFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectPendingCall) {
	n = new(QRemoteObjectPendingCall)
	n.SetPointer(ptr)
	return
}
func (ptr *QRemoteObjectPendingCall) DestroyQRemoteObjectPendingCall() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QRemoteObjectPendingCall__Error
//QRemoteObjectPendingCall::Error
type QRemoteObjectPendingCall__Error int64

const (
	QRemoteObjectPendingCall__NoError        QRemoteObjectPendingCall__Error = QRemoteObjectPendingCall__Error(0)
	QRemoteObjectPendingCall__InvalidMessage QRemoteObjectPendingCall__Error = QRemoteObjectPendingCall__Error(1)
)

func (ptr *QRemoteObjectPendingCall) Error() QRemoteObjectPendingCall__Error {
	if ptr.Pointer() != nil {
		return QRemoteObjectPendingCall__Error(C.QRemoteObjectPendingCall_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRemoteObjectPendingCall) IsFinished() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectPendingCall_IsFinished(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRemoteObjectPendingCall) ReturnValue() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectPendingCall_ReturnValue(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

type QRemoteObjectPendingCallWatcher struct {
	core.QObject
	QRemoteObjectPendingCall
}

type QRemoteObjectPendingCallWatcher_ITF interface {
	core.QObject_ITF
	QRemoteObjectPendingCall_ITF
	QRemoteObjectPendingCallWatcher_PTR() *QRemoteObjectPendingCallWatcher
}

func (ptr *QRemoteObjectPendingCallWatcher) QRemoteObjectPendingCallWatcher_PTR() *QRemoteObjectPendingCallWatcher {
	return ptr
}

func (ptr *QRemoteObjectPendingCallWatcher) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectPendingCallWatcher) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QRemoteObjectPendingCall_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectPendingCallWatcher(ptr QRemoteObjectPendingCallWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectPendingCallWatcher_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectPendingCallWatcherFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectPendingCallWatcher) {
	n = new(QRemoteObjectPendingCallWatcher)
	n.SetPointer(ptr)
	return
}

//export callbackQRemoteObjectPendingCallWatcher_Finished
func callbackQRemoteObjectPendingCallWatcher_Finished(ptr unsafe.Pointer, se unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func(*QRemoteObjectPendingCallWatcher))(signal))(NewQRemoteObjectPendingCallWatcherFromPointer(se))
	}

}

func (ptr *QRemoteObjectPendingCallWatcher) ConnectFinished(f func(se *QRemoteObjectPendingCallWatcher)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QRemoteObjectPendingCallWatcher_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func(se *QRemoteObjectPendingCallWatcher) {
				(*(*func(*QRemoteObjectPendingCallWatcher))(signal))(se)
				f(se)
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) Finished(se QRemoteObjectPendingCallWatcher_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_Finished(ptr.Pointer(), PointerFromQRemoteObjectPendingCallWatcher(se))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) WaitForFinished() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_WaitForFinished(ptr.Pointer())
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectPendingCallWatcher___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectPendingCallWatcher) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) __children_newList() unsafe.Pointer {
	return C.QRemoteObjectPendingCallWatcher___children_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectPendingCallWatcher) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QRemoteObjectPendingCallWatcher___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectPendingCallWatcher) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QRemoteObjectPendingCallWatcher___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectPendingCallWatcher___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_newList() unsafe.Pointer {
	return C.QRemoteObjectPendingCallWatcher___findChildren_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectPendingCallWatcher___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) __findChildren_newList3() unsafe.Pointer {
	return C.QRemoteObjectPendingCallWatcher___findChildren_newList3(ptr.Pointer())
}

//export callbackQRemoteObjectPendingCallWatcher_ChildEvent
func callbackQRemoteObjectPendingCallWatcher_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRemoteObjectPendingCallWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQRemoteObjectPendingCallWatcher_ConnectNotify
func callbackQRemoteObjectPendingCallWatcher_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectPendingCallWatcherFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectPendingCallWatcher_CustomEvent
func callbackQRemoteObjectPendingCallWatcher_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQRemoteObjectPendingCallWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQRemoteObjectPendingCallWatcher_DeleteLater
func callbackQRemoteObjectPendingCallWatcher_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQRemoteObjectPendingCallWatcherFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) DeleteLater() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRemoteObjectPendingCallWatcher_DeleteLater(ptr.Pointer())
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRemoteObjectPendingCallWatcher_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQRemoteObjectPendingCallWatcher_Destroyed
func callbackQRemoteObjectPendingCallWatcher_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQRemoteObjectPendingCallWatcher_DisconnectNotify
func callbackQRemoteObjectPendingCallWatcher_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectPendingCallWatcherFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectPendingCallWatcher_Event
func callbackQRemoteObjectPendingCallWatcher_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectPendingCallWatcherFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QRemoteObjectPendingCallWatcher) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectPendingCallWatcher_Event(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

func (ptr *QRemoteObjectPendingCallWatcher) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectPendingCallWatcher_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQRemoteObjectPendingCallWatcher_EventFilter
func callbackQRemoteObjectPendingCallWatcher_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectPendingCallWatcherFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QRemoteObjectPendingCallWatcher) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectPendingCallWatcher_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QRemoteObjectPendingCallWatcher) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectPendingCallWatcher_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQRemoteObjectPendingCallWatcher_MetaObject
func callbackQRemoteObjectPendingCallWatcher_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQRemoteObjectPendingCallWatcherFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QRemoteObjectPendingCallWatcher) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRemoteObjectPendingCallWatcher_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRemoteObjectPendingCallWatcher) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRemoteObjectPendingCallWatcher_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQRemoteObjectPendingCallWatcher_ObjectNameChanged
func callbackQRemoteObjectPendingCallWatcher_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtRemoteObjects_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQRemoteObjectPendingCallWatcher_TimerEvent
func callbackQRemoteObjectPendingCallWatcher_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRemoteObjectPendingCallWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QRemoteObjectPendingCallWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectPendingCallWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QRemoteObjectPendingReply struct {
	QRemoteObjectPendingCall
}

type QRemoteObjectPendingReply_ITF interface {
	QRemoteObjectPendingCall_ITF
	QRemoteObjectPendingReply_PTR() *QRemoteObjectPendingReply
}

func (ptr *QRemoteObjectPendingReply) QRemoteObjectPendingReply_PTR() *QRemoteObjectPendingReply {
	return ptr
}

func (ptr *QRemoteObjectPendingReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectPendingCall_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectPendingReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectPendingCall_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectPendingReply(ptr QRemoteObjectPendingReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectPendingReply_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectPendingReplyFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectPendingReply) {
	n = new(QRemoteObjectPendingReply)
	n.SetPointer(ptr)
	return
}
func (ptr *QRemoteObjectPendingReply) DestroyQRemoteObjectPendingReply() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QRemoteObjectRegistry struct {
	QRemoteObjectReplica
}

type QRemoteObjectRegistry_ITF interface {
	QRemoteObjectReplica_ITF
	QRemoteObjectRegistry_PTR() *QRemoteObjectRegistry
}

func (ptr *QRemoteObjectRegistry) QRemoteObjectRegistry_PTR() *QRemoteObjectRegistry {
	return ptr
}

func (ptr *QRemoteObjectRegistry) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectReplica_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectRegistry) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectReplica_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectRegistry(ptr QRemoteObjectRegistry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectRegistry_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectRegistryFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectRegistry) {
	n = new(QRemoteObjectRegistry)
	n.SetPointer(ptr)
	return
}

//export callbackQRemoteObjectRegistry_DestroyQRemoteObjectRegistry
func callbackQRemoteObjectRegistry_DestroyQRemoteObjectRegistry(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QRemoteObjectRegistry"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQRemoteObjectRegistryFromPointer(ptr).DestroyQRemoteObjectRegistryDefault()
	}
}

func (ptr *QRemoteObjectRegistry) ConnectDestroyQRemoteObjectRegistry(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QRemoteObjectRegistry"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectRegistry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectRegistry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectRegistry) DisconnectDestroyQRemoteObjectRegistry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QRemoteObjectRegistry")
	}
}

func (ptr *QRemoteObjectRegistry) DestroyQRemoteObjectRegistry() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRemoteObjectRegistry_DestroyQRemoteObjectRegistry(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRemoteObjectRegistry) DestroyQRemoteObjectRegistryDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRemoteObjectRegistry_DestroyQRemoteObjectRegistryDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QRemoteObjectRegistryHost struct {
	QRemoteObjectHostBase
}

type QRemoteObjectRegistryHost_ITF interface {
	QRemoteObjectHostBase_ITF
	QRemoteObjectRegistryHost_PTR() *QRemoteObjectRegistryHost
}

func (ptr *QRemoteObjectRegistryHost) QRemoteObjectRegistryHost_PTR() *QRemoteObjectRegistryHost {
	return ptr
}

func (ptr *QRemoteObjectRegistryHost) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHostBase_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectRegistryHost) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectHostBase_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectRegistryHost(ptr QRemoteObjectRegistryHost_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectRegistryHost_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectRegistryHostFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectRegistryHost) {
	n = new(QRemoteObjectRegistryHost)
	n.SetPointer(ptr)
	return
}
func NewQRemoteObjectRegistryHost(registryAddress core.QUrl_ITF, parent core.QObject_ITF) *QRemoteObjectRegistryHost {
	tmpValue := NewQRemoteObjectRegistryHostFromPointer(C.QRemoteObjectRegistryHost_NewQRemoteObjectRegistryHost(core.PointerFromQUrl(registryAddress), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

type QRemoteObjectReplica struct {
	core.QObject
}

type QRemoteObjectReplica_ITF interface {
	core.QObject_ITF
	QRemoteObjectReplica_PTR() *QRemoteObjectReplica
}

func (ptr *QRemoteObjectReplica) QRemoteObjectReplica_PTR() *QRemoteObjectReplica {
	return ptr
}

func (ptr *QRemoteObjectReplica) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectReplica) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectReplica(ptr QRemoteObjectReplica_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectReplica_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectReplicaFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectReplica) {
	n = new(QRemoteObjectReplica)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QRemoteObjectReplica__State
//QRemoteObjectReplica::State
type QRemoteObjectReplica__State int64

const (
	QRemoteObjectReplica__Uninitialized     QRemoteObjectReplica__State = QRemoteObjectReplica__State(0)
	QRemoteObjectReplica__Default           QRemoteObjectReplica__State = QRemoteObjectReplica__State(1)
	QRemoteObjectReplica__Valid             QRemoteObjectReplica__State = QRemoteObjectReplica__State(2)
	QRemoteObjectReplica__Suspect           QRemoteObjectReplica__State = QRemoteObjectReplica__State(3)
	QRemoteObjectReplica__SignatureMismatch QRemoteObjectReplica__State = QRemoteObjectReplica__State(4)
)

//export callbackQRemoteObjectReplica_Initialized
func callbackQRemoteObjectReplica_Initialized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initialized"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QRemoteObjectReplica) ConnectInitialized(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "initialized") {
			C.QRemoteObjectReplica_ConnectInitialized(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "initialized")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "initialized"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "initialized", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initialized", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectReplica) DisconnectInitialized() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_DisconnectInitialized(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "initialized")
	}
}

func (ptr *QRemoteObjectReplica) Initialized() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_Initialized(ptr.Pointer())
	}
}

func (ptr *QRemoteObjectReplica) IsInitialized() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectReplica_IsInitialized(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRemoteObjectReplica) IsReplicaValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectReplica_IsReplicaValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRemoteObjectReplica) Node() *QRemoteObjectNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQRemoteObjectNodeFromPointer(C.QRemoteObjectReplica_Node(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQRemoteObjectReplica_Notified
func callbackQRemoteObjectReplica_Notified(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "notified"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QRemoteObjectReplica) ConnectNotified(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "notified") {
			C.QRemoteObjectReplica_ConnectNotified(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "notified")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "notified"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "notified", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "notified", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectReplica) DisconnectNotified() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_DisconnectNotified(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "notified")
	}
}

func (ptr *QRemoteObjectReplica) Notified() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_Notified(ptr.Pointer())
	}
}

//export callbackQRemoteObjectReplica_SetNode
func callbackQRemoteObjectReplica_SetNode(ptr unsafe.Pointer, node unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setNode"); signal != nil {
		(*(*func(*QRemoteObjectNode))(signal))(NewQRemoteObjectNodeFromPointer(node))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).SetNodeDefault(NewQRemoteObjectNodeFromPointer(node))
	}
}

func (ptr *QRemoteObjectReplica) ConnectSetNode(f func(node *QRemoteObjectNode)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setNode"); signal != nil {
			f := func(node *QRemoteObjectNode) {
				(*(*func(*QRemoteObjectNode))(signal))(node)
				f(node)
			}
			qt.ConnectSignal(ptr.Pointer(), "setNode", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setNode", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectReplica) DisconnectSetNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setNode")
	}
}

func (ptr *QRemoteObjectReplica) SetNode(node QRemoteObjectNode_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_SetNode(ptr.Pointer(), PointerFromQRemoteObjectNode(node))
	}
}

func (ptr *QRemoteObjectReplica) SetNodeDefault(node QRemoteObjectNode_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_SetNodeDefault(ptr.Pointer(), PointerFromQRemoteObjectNode(node))
	}
}

func (ptr *QRemoteObjectReplica) State() QRemoteObjectReplica__State {
	if ptr.Pointer() != nil {
		return QRemoteObjectReplica__State(C.QRemoteObjectReplica_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQRemoteObjectReplica_StateChanged
func callbackQRemoteObjectReplica_StateChanged(ptr unsafe.Pointer, state C.longlong, oldState C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(QRemoteObjectReplica__State, QRemoteObjectReplica__State))(signal))(QRemoteObjectReplica__State(state), QRemoteObjectReplica__State(oldState))
	}

}

func (ptr *QRemoteObjectReplica) ConnectStateChanged(f func(state QRemoteObjectReplica__State, oldState QRemoteObjectReplica__State)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QRemoteObjectReplica_ConnectStateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "stateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			f := func(state QRemoteObjectReplica__State, oldState QRemoteObjectReplica__State) {
				(*(*func(QRemoteObjectReplica__State, QRemoteObjectReplica__State))(signal))(state, oldState)
				f(state, oldState)
			}
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRemoteObjectReplica) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QRemoteObjectReplica) StateChanged(state QRemoteObjectReplica__State, oldState QRemoteObjectReplica__State) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_StateChanged(ptr.Pointer(), C.longlong(state), C.longlong(oldState))
	}
}

func (ptr *QRemoteObjectReplica) WaitForSource(timeout int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectReplica_WaitForSource(ptr.Pointer(), C.int(int32(timeout)))) != 0
	}
	return false
}

func (ptr *QRemoteObjectReplica) __persistProperties_props_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectReplica___persistProperties_props_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __persistProperties_props_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___persistProperties_props_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectReplica) __persistProperties_props_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___persistProperties_props_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __retrieveProperties_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectReplica___retrieveProperties_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __retrieveProperties_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___retrieveProperties_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectReplica) __retrieveProperties_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___retrieveProperties_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __send_args_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectReplica___send_args_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __send_args_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___send_args_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectReplica) __send_args_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___send_args_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __sendWithReply_args_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectReplica___sendWithReply_args_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __sendWithReply_args_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___sendWithReply_args_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectReplica) __sendWithReply_args_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___sendWithReply_args_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __setProperties_properties_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectReplica___setProperties_properties_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __setProperties_properties_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___setProperties_properties_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectReplica) __setProperties_properties_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___setProperties_properties_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectReplica___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectReplica) __children_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___children_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QRemoteObjectReplica___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QRemoteObjectReplica) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectReplica___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectReplica) __findChildren_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___findChildren_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectReplica___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectReplica) __findChildren_newList3() unsafe.Pointer {
	return C.QRemoteObjectReplica___findChildren_newList3(ptr.Pointer())
}

//export callbackQRemoteObjectReplica_ChildEvent
func callbackQRemoteObjectReplica_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectReplica) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQRemoteObjectReplica_ConnectNotify
func callbackQRemoteObjectReplica_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectReplica) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectReplica_CustomEvent
func callbackQRemoteObjectReplica_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectReplica) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQRemoteObjectReplica_DeleteLater
func callbackQRemoteObjectReplica_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRemoteObjectReplica) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRemoteObjectReplica_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQRemoteObjectReplica_Destroyed
func callbackQRemoteObjectReplica_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQRemoteObjectReplica_DisconnectNotify
func callbackQRemoteObjectReplica_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectReplica) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectReplica_Event
func callbackQRemoteObjectReplica_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectReplicaFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QRemoteObjectReplica) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectReplica_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQRemoteObjectReplica_EventFilter
func callbackQRemoteObjectReplica_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectReplicaFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QRemoteObjectReplica) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectReplica_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQRemoteObjectReplica_MetaObject
func callbackQRemoteObjectReplica_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQRemoteObjectReplicaFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QRemoteObjectReplica) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRemoteObjectReplica_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQRemoteObjectReplica_ObjectNameChanged
func callbackQRemoteObjectReplica_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtRemoteObjects_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQRemoteObjectReplica_TimerEvent
func callbackQRemoteObjectReplica_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectReplica) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QtROClientFactory struct {
	ptr unsafe.Pointer
}

type QtROClientFactory_ITF interface {
	QtROClientFactory_PTR() *QtROClientFactory
}

func (ptr *QtROClientFactory) QtROClientFactory_PTR() *QtROClientFactory {
	return ptr
}

func (ptr *QtROClientFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QtROClientFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQtROClientFactory(ptr QtROClientFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtROClientFactory_PTR().Pointer()
	}
	return nil
}

func NewQtROClientFactoryFromPointer(ptr unsafe.Pointer) (n *QtROClientFactory) {
	n = new(QtROClientFactory)
	n.SetPointer(ptr)
	return
}
func (ptr *QtROClientFactory) DestroyQtROClientFactory() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QtROServerFactory struct {
	ptr unsafe.Pointer
}

type QtROServerFactory_ITF interface {
	QtROServerFactory_PTR() *QtROServerFactory
}

func (ptr *QtROServerFactory) QtROServerFactory_PTR() *QtROServerFactory {
	return ptr
}

func (ptr *QtROServerFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QtROServerFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQtROServerFactory(ptr QtROServerFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtROServerFactory_PTR().Pointer()
	}
	return nil
}

func NewQtROServerFactoryFromPointer(ptr unsafe.Pointer) (n *QtROServerFactory) {
	n = new(QtROServerFactory)
	n.SetPointer(ptr)
	return
}
func (ptr *QtROServerFactory) DestroyQtROServerFactory() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func init() {
	qt.ItfMap["remoteobjects.QRemoteObjectAbstractPersistedStore_ITF"] = QRemoteObjectAbstractPersistedStore{}
	qt.FuncMap["remoteobjects.NewQRemoteObjectAbstractPersistedStore"] = NewQRemoteObjectAbstractPersistedStore
	qt.ItfMap["remoteobjects.QRemoteObjectDynamicReplica_ITF"] = QRemoteObjectDynamicReplica{}
	qt.ItfMap["remoteobjects.QRemoteObjectHost_ITF"] = QRemoteObjectHost{}
	qt.FuncMap["remoteobjects.NewQRemoteObjectHost"] = NewQRemoteObjectHost
	qt.FuncMap["remoteobjects.NewQRemoteObjectHost2"] = NewQRemoteObjectHost2
	qt.FuncMap["remoteobjects.NewQRemoteObjectHost3"] = NewQRemoteObjectHost3
	qt.ItfMap["remoteobjects.QRemoteObjectHostBase_ITF"] = QRemoteObjectHostBase{}
	qt.EnumMap["remoteobjects.QRemoteObjectHostBase__BuiltInSchemasOnly"] = int64(QRemoteObjectHostBase__BuiltInSchemasOnly)
	qt.EnumMap["remoteobjects.QRemoteObjectHostBase__AllowExternalRegistration"] = int64(QRemoteObjectHostBase__AllowExternalRegistration)
	qt.ItfMap["remoteobjects.QRemoteObjectNode_ITF"] = QRemoteObjectNode{}
	qt.FuncMap["remoteobjects.NewQRemoteObjectNode"] = NewQRemoteObjectNode
	qt.FuncMap["remoteobjects.NewQRemoteObjectNode2"] = NewQRemoteObjectNode2
	qt.EnumMap["remoteobjects.QRemoteObjectNode__NoError"] = int64(QRemoteObjectNode__NoError)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__RegistryNotAcquired"] = int64(QRemoteObjectNode__RegistryNotAcquired)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__RegistryAlreadyHosted"] = int64(QRemoteObjectNode__RegistryAlreadyHosted)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__NodeIsNoServer"] = int64(QRemoteObjectNode__NodeIsNoServer)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__ServerAlreadyCreated"] = int64(QRemoteObjectNode__ServerAlreadyCreated)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__UnintendedRegistryHosting"] = int64(QRemoteObjectNode__UnintendedRegistryHosting)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__OperationNotValidOnClientNode"] = int64(QRemoteObjectNode__OperationNotValidOnClientNode)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__SourceNotRegistered"] = int64(QRemoteObjectNode__SourceNotRegistered)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__MissingObjectName"] = int64(QRemoteObjectNode__MissingObjectName)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__HostUrlInvalid"] = int64(QRemoteObjectNode__HostUrlInvalid)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__ProtocolMismatch"] = int64(QRemoteObjectNode__ProtocolMismatch)
	qt.EnumMap["remoteobjects.QRemoteObjectNode__ListenFailed"] = int64(QRemoteObjectNode__ListenFailed)
	qt.ItfMap["remoteobjects.QRemoteObjectPendingCall_ITF"] = QRemoteObjectPendingCall{}
	qt.EnumMap["remoteobjects.QRemoteObjectPendingCall__NoError"] = int64(QRemoteObjectPendingCall__NoError)
	qt.EnumMap["remoteobjects.QRemoteObjectPendingCall__InvalidMessage"] = int64(QRemoteObjectPendingCall__InvalidMessage)
	qt.ItfMap["remoteobjects.QRemoteObjectPendingCallWatcher_ITF"] = QRemoteObjectPendingCallWatcher{}
	qt.ItfMap["remoteobjects.QRemoteObjectPendingReply_ITF"] = QRemoteObjectPendingReply{}
	qt.ItfMap["remoteobjects.QRemoteObjectRegistry_ITF"] = QRemoteObjectRegistry{}
	qt.ItfMap["remoteobjects.QRemoteObjectRegistryHost_ITF"] = QRemoteObjectRegistryHost{}
	qt.FuncMap["remoteobjects.NewQRemoteObjectRegistryHost"] = NewQRemoteObjectRegistryHost
	qt.ItfMap["remoteobjects.QRemoteObjectReplica_ITF"] = QRemoteObjectReplica{}
	qt.EnumMap["remoteobjects.QRemoteObjectReplica__Uninitialized"] = int64(QRemoteObjectReplica__Uninitialized)
	qt.EnumMap["remoteobjects.QRemoteObjectReplica__Default"] = int64(QRemoteObjectReplica__Default)
	qt.EnumMap["remoteobjects.QRemoteObjectReplica__Valid"] = int64(QRemoteObjectReplica__Valid)
	qt.EnumMap["remoteobjects.QRemoteObjectReplica__Suspect"] = int64(QRemoteObjectReplica__Suspect)
	qt.EnumMap["remoteobjects.QRemoteObjectReplica__SignatureMismatch"] = int64(QRemoteObjectReplica__SignatureMismatch)
	qt.ItfMap["remoteobjects.QtROClientFactory_ITF"] = QtROClientFactory{}
	qt.ItfMap["remoteobjects.QtROServerFactory_ITF"] = QtROServerFactory{}
}
