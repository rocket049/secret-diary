//go:build minimal
// +build minimal

package core

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "core-minimal.h"
import "C"
import (
	"github.com/therecipe/qt"
	"math"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtCore_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtCore_PackedString) []byte {
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

// QAbstractAnimation::DeletionPolicy
//
//go:generate stringer -type=QAbstractAnimation__DeletionPolicy
type QAbstractAnimation__DeletionPolicy int64

const (
	QAbstractAnimation__KeepWhenStopped   QAbstractAnimation__DeletionPolicy = QAbstractAnimation__DeletionPolicy(0)
	QAbstractAnimation__DeleteWhenStopped QAbstractAnimation__DeletionPolicy = QAbstractAnimation__DeletionPolicy(1)
)

// QAbstractAnimation::Direction
//
//go:generate stringer -type=QAbstractAnimation__Direction
type QAbstractAnimation__Direction int64

const (
	QAbstractAnimation__Forward  QAbstractAnimation__Direction = QAbstractAnimation__Direction(0)
	QAbstractAnimation__Backward QAbstractAnimation__Direction = QAbstractAnimation__Direction(1)
)

// QAbstractAnimation::State
//
//go:generate stringer -type=QAbstractAnimation__State
type QAbstractAnimation__State int64

const (
	QAbstractAnimation__Stopped QAbstractAnimation__State = QAbstractAnimation__State(0)
	QAbstractAnimation__Paused  QAbstractAnimation__State = QAbstractAnimation__State(1)
	QAbstractAnimation__Running QAbstractAnimation__State = QAbstractAnimation__State(2)
)

type QAbstractItemModel struct {
	QObject
}

type QAbstractItemModel_ITF interface {
	QObject_ITF
	QAbstractItemModel_PTR() *QAbstractItemModel
}

func (ptr *QAbstractItemModel) QAbstractItemModel_PTR() *QAbstractItemModel {
	return ptr
}

func (ptr *QAbstractItemModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractItemModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractItemModel(ptr QAbstractItemModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractItemModelFromPointer(ptr unsafe.Pointer) (n *QAbstractItemModel) {
	n = new(QAbstractItemModel)
	n.SetPointer(ptr)
	return
}

// QAbstractItemModel::CheckIndexOption
//
//go:generate stringer -type=QAbstractItemModel__CheckIndexOption
type QAbstractItemModel__CheckIndexOption int64

const (
	QAbstractItemModel__NoOption        QAbstractItemModel__CheckIndexOption = QAbstractItemModel__CheckIndexOption(0x0000)
	QAbstractItemModel__IndexIsValid    QAbstractItemModel__CheckIndexOption = QAbstractItemModel__CheckIndexOption(0x0001)
	QAbstractItemModel__DoNotUseParent  QAbstractItemModel__CheckIndexOption = QAbstractItemModel__CheckIndexOption(0x0002)
	QAbstractItemModel__ParentIsInvalid QAbstractItemModel__CheckIndexOption = QAbstractItemModel__CheckIndexOption(0x0004)
)

// QAbstractItemModel::LayoutChangeHint
//
//go:generate stringer -type=QAbstractItemModel__LayoutChangeHint
type QAbstractItemModel__LayoutChangeHint int64

const (
	QAbstractItemModel__NoLayoutChangeHint QAbstractItemModel__LayoutChangeHint = QAbstractItemModel__LayoutChangeHint(0)
	QAbstractItemModel__VerticalSortHint   QAbstractItemModel__LayoutChangeHint = QAbstractItemModel__LayoutChangeHint(1)
	QAbstractItemModel__HorizontalSortHint QAbstractItemModel__LayoutChangeHint = QAbstractItemModel__LayoutChangeHint(2)
)

func NewQAbstractItemModel(parent QObject_ITF) *QAbstractItemModel {
	tmpValue := NewQAbstractItemModelFromPointer(C.QAbstractItemModel_NewQAbstractItemModel(PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAbstractItemModel_ColumnCount
func callbackQAbstractItemModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*QModelIndex) int)(signal))(NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(0))
}

func (ptr *QAbstractItemModel) ConnectColumnCount(f func(parent *QModelIndex) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "columnCount"); signal != nil {
			f := func(parent *QModelIndex) int {
				(*(*func(*QModelIndex) int)(signal))(parent)
				return f(parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "columnCount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "columnCount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectColumnCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "columnCount")
	}
}

func (ptr *QAbstractItemModel) ColumnCount(parent QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemModel_ColumnCount(ptr.Pointer(), PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQAbstractItemModel_Data
func callbackQAbstractItemModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return PointerFromQVariant((*(*func(*QModelIndex, int) *QVariant)(signal))(NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return PointerFromQVariant(NewQVariant())
}

func (ptr *QAbstractItemModel) ConnectData(f func(index *QModelIndex, role int) *QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			f := func(index *QModelIndex, role int) *QVariant {
				(*(*func(*QModelIndex, int) *QVariant)(signal))(index, role)
				return f(index, role)
			}
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *QAbstractItemModel) Data(index QModelIndex_ITF, role int) *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QAbstractItemModel_Data(ptr.Pointer(), PointerFromQModelIndex(index), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractItemModel_HasChildren
func callbackQAbstractItemModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QModelIndex) bool)(signal))(NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractItemModelFromPointer(ptr).HasChildrenDefault(NewQModelIndexFromPointer(parent)))))
}

func (ptr *QAbstractItemModel) ConnectHasChildren(f func(parent *QModelIndex) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasChildren"); signal != nil {
			f := func(parent *QModelIndex) bool {
				(*(*func(*QModelIndex) bool)(signal))(parent)
				return f(parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "hasChildren", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasChildren", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectHasChildren() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasChildren")
	}
}

func (ptr *QAbstractItemModel) HasChildren(parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_HasChildren(ptr.Pointer(), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) HasChildrenDefault(parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_HasChildrenDefault(ptr.Pointer(), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQAbstractItemModel_Index
func callbackQAbstractItemModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return PointerFromQModelIndex((*(*func(int, int, *QModelIndex) *QModelIndex)(signal))(int(int32(row)), int(int32(column)), NewQModelIndexFromPointer(parent)))
	}

	return PointerFromQModelIndex(NewQModelIndex())
}

func (ptr *QAbstractItemModel) ConnectIndex(f func(row int, column int, parent *QModelIndex) *QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "index"); signal != nil {
			f := func(row int, column int, parent *QModelIndex) *QModelIndex {
				(*(*func(int, int, *QModelIndex) *QModelIndex)(signal))(row, column, parent)
				return f(row, column, parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "index", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "index", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "index")
	}
}

func (ptr *QAbstractItemModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QAbstractItemModel_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), PointerFromQModelIndex(parent)))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) InsertColumn(column int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_InsertColumn(ptr.Pointer(), C.int(int32(column)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQAbstractItemModel_InsertColumns
func callbackQAbstractItemModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractItemModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), NewQModelIndexFromPointer(parent)))))
}

func (ptr *QAbstractItemModel) ConnectInsertColumns(f func(column int, count int, parent *QModelIndex) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "insertColumns"); signal != nil {
			f := func(column int, count int, parent *QModelIndex) bool {
				(*(*func(int, int, *QModelIndex) bool)(signal))(column, count, parent)
				return f(column, count, parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "insertColumns", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "insertColumns", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectInsertColumns() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "insertColumns")
	}
}

func (ptr *QAbstractItemModel) InsertColumns(column int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_InsertColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) InsertColumnsDefault(column int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) InsertRow(row int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_InsertRow(ptr.Pointer(), C.int(int32(row)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQAbstractItemModel_InsertRows
func callbackQAbstractItemModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractItemModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), NewQModelIndexFromPointer(parent)))))
}

func (ptr *QAbstractItemModel) ConnectInsertRows(f func(row int, count int, parent *QModelIndex) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "insertRows"); signal != nil {
			f := func(row int, count int, parent *QModelIndex) bool {
				(*(*func(int, int, *QModelIndex) bool)(signal))(row, count, parent)
				return f(row, count, parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "insertRows", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "insertRows", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectInsertRows() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "insertRows")
	}
}

func (ptr *QAbstractItemModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_InsertRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) InsertRowsDefault(row int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQAbstractItemModel_Match
func callbackQAbstractItemModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQAbstractItemModelFromPointer(NewQAbstractItemModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*QModelIndex, int, *QVariant, int, Qt__MatchFlag) []*QModelIndex)(signal))(NewQModelIndexFromPointer(start), int(int32(role)), NewQVariantFromPointer(value), int(int32(hits)), Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQAbstractItemModelFromPointer(NewQAbstractItemModelFromPointer(nil).__match_newList())
		for _, v := range NewQAbstractItemModelFromPointer(ptr).MatchDefault(NewQModelIndexFromPointer(start), int(int32(role)), NewQVariantFromPointer(value), int(int32(hits)), Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QAbstractItemModel) ConnectMatch(f func(start *QModelIndex, role int, value *QVariant, hits int, flags Qt__MatchFlag) []*QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "match"); signal != nil {
			f := func(start *QModelIndex, role int, value *QVariant, hits int, flags Qt__MatchFlag) []*QModelIndex {
				(*(*func(*QModelIndex, int, *QVariant, int, Qt__MatchFlag) []*QModelIndex)(signal))(start, role, value, hits, flags)
				return f(start, role, value, hits, flags)
			}
			qt.ConnectSignal(ptr.Pointer(), "match", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "match", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectMatch() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "match")
	}
}

func (ptr *QAbstractItemModel) Match(start QModelIndex_ITF, role int, value QVariant_ITF, hits int, flags Qt__MatchFlag) []*QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []*QModelIndex {
			out := make([]*QModelIndex, int(l.len))
			tmpList := NewQAbstractItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.QAbstractItemModel_Match(ptr.Pointer(), PointerFromQModelIndex(start), C.int(int32(role)), PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*QModelIndex, 0)
}

func (ptr *QAbstractItemModel) MatchDefault(start QModelIndex_ITF, role int, value QVariant_ITF, hits int, flags Qt__MatchFlag) []*QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []*QModelIndex {
			out := make([]*QModelIndex, int(l.len))
			tmpList := NewQAbstractItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.QAbstractItemModel_MatchDefault(ptr.Pointer(), PointerFromQModelIndex(start), C.int(int32(role)), PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*QModelIndex, 0)
}

//export callbackQAbstractItemModel_Parent
func callbackQAbstractItemModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return PointerFromQModelIndex((*(*func(*QModelIndex) *QModelIndex)(signal))(NewQModelIndexFromPointer(index)))
	}

	return PointerFromQModelIndex(NewQModelIndex())
}

func (ptr *QAbstractItemModel) ConnectParent(f func(index *QModelIndex) *QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parent"); signal != nil {
			f := func(index *QModelIndex) *QModelIndex {
				(*(*func(*QModelIndex) *QModelIndex)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "parent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectParent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parent")
	}
}

func (ptr *QAbstractItemModel) Parent(index QModelIndex_ITF) *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QAbstractItemModel_Parent(ptr.Pointer(), PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) RemoveColumn(column int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_RemoveColumn(ptr.Pointer(), C.int(int32(column)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQAbstractItemModel_RemoveColumns
func callbackQAbstractItemModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractItemModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), NewQModelIndexFromPointer(parent)))))
}

func (ptr *QAbstractItemModel) ConnectRemoveColumns(f func(column int, count int, parent *QModelIndex) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeColumns"); signal != nil {
			f := func(column int, count int, parent *QModelIndex) bool {
				(*(*func(int, int, *QModelIndex) bool)(signal))(column, count, parent)
				return f(column, count, parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "removeColumns", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeColumns", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectRemoveColumns() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeColumns")
	}
}

func (ptr *QAbstractItemModel) RemoveColumns(column int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_RemoveColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) RemoveColumnsDefault(column int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) RemoveRow(row int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_RemoveRow(ptr.Pointer(), C.int(int32(row)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQAbstractItemModel_RemoveRows
func callbackQAbstractItemModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractItemModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), NewQModelIndexFromPointer(parent)))))
}

func (ptr *QAbstractItemModel) ConnectRemoveRows(f func(row int, count int, parent *QModelIndex) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeRows"); signal != nil {
			f := func(row int, count int, parent *QModelIndex) bool {
				(*(*func(int, int, *QModelIndex) bool)(signal))(row, count, parent)
				return f(row, count, parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "removeRows", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeRows", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectRemoveRows() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeRows")
	}
}

func (ptr *QAbstractItemModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_RemoveRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) RemoveRowsDefault(row int, count int, parent QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQAbstractItemModel_RowCount
func callbackQAbstractItemModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*QModelIndex) int)(signal))(NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(0))
}

func (ptr *QAbstractItemModel) ConnectRowCount(f func(parent *QModelIndex) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rowCount"); signal != nil {
			f := func(parent *QModelIndex) int {
				(*(*func(*QModelIndex) int)(signal))(parent)
				return f(parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "rowCount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rowCount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectRowCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rowCount")
	}
}

func (ptr *QAbstractItemModel) RowCount(parent QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemModel_RowCount(ptr.Pointer(), PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQAbstractItemModel_SetData
func callbackQAbstractItemModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QModelIndex, *QVariant, int) bool)(signal))(NewQModelIndexFromPointer(index), NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractItemModelFromPointer(ptr).SetDataDefault(NewQModelIndexFromPointer(index), NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QAbstractItemModel) ConnectSetData(f func(index *QModelIndex, value *QVariant, role int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setData"); signal != nil {
			f := func(index *QModelIndex, value *QVariant, role int) bool {
				(*(*func(*QModelIndex, *QVariant, int) bool)(signal))(index, value, role)
				return f(index, value, role)
			}
			qt.ConnectSignal(ptr.Pointer(), "setData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectSetData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setData")
	}
}

func (ptr *QAbstractItemModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_SetData(ptr.Pointer(), PointerFromQModelIndex(index), PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

func (ptr *QAbstractItemModel) SetDataDefault(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemModel_SetDataDefault(ptr.Pointer(), PointerFromQModelIndex(index), PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackQAbstractItemModel_DestroyQAbstractItemModel
func callbackQAbstractItemModel_DestroyQAbstractItemModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractItemModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractItemModelFromPointer(ptr).DestroyQAbstractItemModelDefault()
	}
}

func (ptr *QAbstractItemModel) ConnectDestroyQAbstractItemModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractItemModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractItemModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractItemModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemModel) DisconnectDestroyQAbstractItemModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractItemModel")
	}
}

func (ptr *QAbstractItemModel) DestroyQAbstractItemModel() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractItemModel_DestroyQAbstractItemModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractItemModel) DestroyQAbstractItemModelDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractItemModel_DestroyQAbstractItemModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractItemModel) __changePersistentIndexList_from_atList(i int) *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QAbstractItemModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __changePersistentIndexList_from_setList(i QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___changePersistentIndexList_from_setList(ptr.Pointer(), PointerFromQModelIndex(i))
	}
}

func (ptr *QAbstractItemModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.QAbstractItemModel___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __changePersistentIndexList_to_atList(i int) *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QAbstractItemModel___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __changePersistentIndexList_to_setList(i QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___changePersistentIndexList_to_setList(ptr.Pointer(), PointerFromQModelIndex(i))
	}
}

func (ptr *QAbstractItemModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.QAbstractItemModel___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemModel___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QAbstractItemModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QAbstractItemModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.QAbstractItemModel___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __doSetRoleNames_roleNames_atList(v int, i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QAbstractItemModel___doSetRoleNames_roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __doSetRoleNames_roleNames_setList(key int, i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___doSetRoleNames_roleNames_setList(ptr.Pointer(), C.int(int32(key)), PointerFromQByteArray(i))
	}
}

func (ptr *QAbstractItemModel) __doSetRoleNames_roleNames_newList() unsafe.Pointer {
	return C.QAbstractItemModel___doSetRoleNames_roleNames_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __doSetRoleNames_roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQAbstractItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____doSetRoleNames_roleNames_keyList_atList(i)
			}
			return out
		}(C.QAbstractItemModel___doSetRoleNames_roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QAbstractItemModel) __encodeData_indexes_atList(i int) *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QAbstractItemModel___encodeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __encodeData_indexes_setList(i QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___encodeData_indexes_setList(ptr.Pointer(), PointerFromQModelIndex(i))
	}
}

func (ptr *QAbstractItemModel) __encodeData_indexes_newList() unsafe.Pointer {
	return C.QAbstractItemModel___encodeData_indexes_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __itemData_atList(v int, i int) *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QAbstractItemModel___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __itemData_setList(key int, i QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___itemData_setList(ptr.Pointer(), C.int(int32(key)), PointerFromQVariant(i))
	}
}

func (ptr *QAbstractItemModel) __itemData_newList() unsafe.Pointer {
	return C.QAbstractItemModel___itemData_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQAbstractItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.QAbstractItemModel___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QAbstractItemModel) __layoutAboutToBeChanged_parents_atList(i int) *QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQPersistentModelIndexFromPointer(C.QAbstractItemModel___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __layoutAboutToBeChanged_parents_setList(i QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QAbstractItemModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.QAbstractItemModel___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __layoutChanged_parents_atList(i int) *QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQPersistentModelIndexFromPointer(C.QAbstractItemModel___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __layoutChanged_parents_setList(i QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___layoutChanged_parents_setList(ptr.Pointer(), PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QAbstractItemModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.QAbstractItemModel___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __match_atList(i int) *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QAbstractItemModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __match_setList(i QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___match_setList(ptr.Pointer(), PointerFromQModelIndex(i))
	}
}

func (ptr *QAbstractItemModel) __match_newList() unsafe.Pointer {
	return C.QAbstractItemModel___match_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __mimeData_indexes_atList(i int) *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QAbstractItemModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __mimeData_indexes_setList(i QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___mimeData_indexes_setList(ptr.Pointer(), PointerFromQModelIndex(i))
	}
}

func (ptr *QAbstractItemModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.QAbstractItemModel___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __persistentIndexList_atList(i int) *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QAbstractItemModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __persistentIndexList_setList(i QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___persistentIndexList_setList(ptr.Pointer(), PointerFromQModelIndex(i))
	}
}

func (ptr *QAbstractItemModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.QAbstractItemModel___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __roleNames_atList(v int, i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QAbstractItemModel___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __roleNames_setList(key int, i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___roleNames_setList(ptr.Pointer(), C.int(int32(key)), PointerFromQByteArray(i))
	}
}

func (ptr *QAbstractItemModel) __roleNames_newList() unsafe.Pointer {
	return C.QAbstractItemModel___roleNames_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQAbstractItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.QAbstractItemModel___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QAbstractItemModel) __setItemData_roles_atList(v int, i int) *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QAbstractItemModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __setItemData_roles_setList(key int, i QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), PointerFromQVariant(i))
	}
}

func (ptr *QAbstractItemModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.QAbstractItemModel___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQAbstractItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.QAbstractItemModel___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QAbstractItemModel) __setRoleNames_roleNames_atList(v int, i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QAbstractItemModel___setRoleNames_roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemModel) __setRoleNames_roleNames_setList(key int, i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel___setRoleNames_roleNames_setList(ptr.Pointer(), C.int(int32(key)), PointerFromQByteArray(i))
	}
}

func (ptr *QAbstractItemModel) __setRoleNames_roleNames_newList() unsafe.Pointer {
	return C.QAbstractItemModel___setRoleNames_roleNames_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) __setRoleNames_roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQAbstractItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoleNames_roleNames_keyList_atList(i)
			}
			return out
		}(C.QAbstractItemModel___setRoleNames_roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QAbstractItemModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemModel_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QAbstractItemModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QAbstractItemModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QAbstractItemModel_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemModel_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QAbstractItemModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QAbstractItemModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.QAbstractItemModel_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemModel_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QAbstractItemModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QAbstractItemModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.QAbstractItemModel_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemModel_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QAbstractItemModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QAbstractItemModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.QAbstractItemModel_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *QAbstractItemModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemModel_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QAbstractItemModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QAbstractItemModel_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QAbstractItemModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QAbstractItemModel_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

// QAbstractTransition::TransitionType
//
//go:generate stringer -type=QAbstractTransition__TransitionType
type QAbstractTransition__TransitionType int64

const (
	QAbstractTransition__ExternalTransition QAbstractTransition__TransitionType = QAbstractTransition__TransitionType(0)
	QAbstractTransition__InternalTransition QAbstractTransition__TransitionType = QAbstractTransition__TransitionType(1)
)

type QBitArray struct {
	ptr unsafe.Pointer
}

type QBitArray_ITF interface {
	QBitArray_PTR() *QBitArray
}

func (ptr *QBitArray) QBitArray_PTR() *QBitArray {
	return ptr
}

func (ptr *QBitArray) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBitArray) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBitArray(ptr QBitArray_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBitArray_PTR().Pointer()
	}
	return nil
}

func NewQBitArrayFromPointer(ptr unsafe.Pointer) (n *QBitArray) {
	n = new(QBitArray)
	n.SetPointer(ptr)
	return
}
func (ptr *QBitArray) DestroyQBitArray() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQBitArray() *QBitArray {
	tmpValue := NewQBitArrayFromPointer(C.QBitArray_NewQBitArray())
	qt.SetFinalizer(tmpValue, (*QBitArray).DestroyQBitArray)
	return tmpValue
}

func NewQBitArray2(size int, value bool) *QBitArray {
	tmpValue := NewQBitArrayFromPointer(C.QBitArray_NewQBitArray2(C.int(int32(size)), C.char(int8(qt.GoBoolToInt(value)))))
	qt.SetFinalizer(tmpValue, (*QBitArray).DestroyQBitArray)
	return tmpValue
}

func NewQBitArray3(other QBitArray_ITF) *QBitArray {
	tmpValue := NewQBitArrayFromPointer(C.QBitArray_NewQBitArray3(PointerFromQBitArray(other)))
	qt.SetFinalizer(tmpValue, (*QBitArray).DestroyQBitArray)
	return tmpValue
}

func NewQBitArray4(other QBitArray_ITF) *QBitArray {
	tmpValue := NewQBitArrayFromPointer(C.QBitArray_NewQBitArray4(PointerFromQBitArray(other)))
	qt.SetFinalizer(tmpValue, (*QBitArray).DestroyQBitArray)
	return tmpValue
}

func (ptr *QBitArray) At(i int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBitArray_At(ptr.Pointer(), C.int(int32(i)))) != 0
	}
	return false
}

func (ptr *QBitArray) Clear() {
	if ptr.Pointer() != nil {
		C.QBitArray_Clear(ptr.Pointer())
	}
}

func (ptr *QBitArray) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBitArray_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBitArray) Count2(on bool) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBitArray_Count2(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(on))))))
	}
	return 0
}

func (ptr *QBitArray) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBitArray_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBitArray) Resize(size int) {
	if ptr.Pointer() != nil {
		C.QBitArray_Resize(ptr.Pointer(), C.int(int32(size)))
	}
}

func (ptr *QBitArray) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBitArray_Size(ptr.Pointer())))
	}
	return 0
}

type QBuffer struct {
	QIODevice
}

type QBuffer_ITF interface {
	QIODevice_ITF
	QBuffer_PTR() *QBuffer
}

func (ptr *QBuffer) QBuffer_PTR() *QBuffer {
	return ptr
}

func (ptr *QBuffer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QBuffer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQBuffer(ptr QBuffer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBuffer_PTR().Pointer()
	}
	return nil
}

func NewQBufferFromPointer(ptr unsafe.Pointer) (n *QBuffer) {
	n = new(QBuffer)
	n.SetPointer(ptr)
	return
}
func NewQBuffer(parent QObject_ITF) *QBuffer {
	tmpValue := NewQBufferFromPointer(C.QBuffer_NewQBuffer(PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBuffer2(byteArray QByteArray_ITF, parent QObject_ITF) *QBuffer {
	tmpValue := NewQBufferFromPointer(C.QBuffer_NewQBuffer2(PointerFromQByteArray(byteArray), PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBuffer) Buffer() *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QBuffer_Buffer(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QBuffer) Buffer2() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Buffer2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) Data() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QBuffer_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBuffer) SetData(data QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBuffer_SetData(ptr.Pointer(), PointerFromQByteArray(data))
	}
}

func (ptr *QBuffer) SetData2(data []byte, size int) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		C.QBuffer_SetData2(ptr.Pointer(), dataC, C.int(int32(size)))
	}
}

//export callbackQBuffer_DestroyQBuffer
func callbackQBuffer_DestroyQBuffer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBuffer"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBufferFromPointer(ptr).DestroyQBufferDefault()
	}
}

func (ptr *QBuffer) ConnectDestroyQBuffer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBuffer"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBuffer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBuffer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBuffer) DisconnectDestroyQBuffer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBuffer")
	}
}

func (ptr *QBuffer) DestroyQBuffer() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBuffer_DestroyQBuffer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBuffer) DestroyQBufferDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBuffer_DestroyQBufferDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBuffer_ReadData
func callbackQBuffer_ReadData(ptr unsafe.Pointer, data C.struct_QtCore_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		retS := cGoUnpackString(data)
		ret := C.longlong((*(*func(*string, int64) int64)(signal))(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
		}
		return ret
	}
	retS := cGoUnpackString(data)
	ret := C.longlong(NewQBufferFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
	}
	return ret
}

func (ptr *QBuffer) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QBuffer_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QBuffer) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QBuffer_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQBuffer_WriteData
func callbackQBuffer_WriteData(ptr unsafe.Pointer, data C.struct_QtCore_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxSize)))
	}

	return C.longlong(NewQBufferFromPointer(ptr).WriteDataDefault(cGoUnpackBytes(data), int64(maxSize)))
}

func (ptr *QBuffer) WriteData(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QBuffer_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBuffer) WriteDataDefault(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QBuffer_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

type QByteArray struct {
	ptr unsafe.Pointer
}

type QByteArray_ITF interface {
	QByteArray_PTR() *QByteArray
}

func (ptr *QByteArray) QByteArray_PTR() *QByteArray {
	return ptr
}

func (ptr *QByteArray) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QByteArray) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQByteArray(ptr QByteArray_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QByteArray_PTR().Pointer()
	}
	return nil
}

func NewQByteArrayFromPointer(ptr unsafe.Pointer) (n *QByteArray) {
	n = new(QByteArray)
	n.SetPointer(ptr)
	return
}

// QByteArray::Base64Option
//
//go:generate stringer -type=QByteArray__Base64Option
type QByteArray__Base64Option int64

const (
	QByteArray__Base64Encoding              QByteArray__Base64Option = QByteArray__Base64Option(0)
	QByteArray__Base64UrlEncoding           QByteArray__Base64Option = QByteArray__Base64Option(1)
	QByteArray__KeepTrailingEquals          QByteArray__Base64Option = QByteArray__Base64Option(0)
	QByteArray__OmitTrailingEquals          QByteArray__Base64Option = QByteArray__Base64Option(2)
	QByteArray__IgnoreBase64DecodingErrors  QByteArray__Base64Option = QByteArray__Base64Option(0)
	QByteArray__AbortOnBase64DecodingErrors QByteArray__Base64Option = QByteArray__Base64Option(4)
)

func NewQByteArray() *QByteArray {
	tmpValue := NewQByteArrayFromPointer(C.QByteArray_NewQByteArray())
	qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
	return tmpValue
}

func NewQByteArray2(data string, size int) *QByteArray {
	var dataC *C.char
	if data != "" {
		dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
	}
	tmpValue := NewQByteArrayFromPointer(C.QByteArray_NewQByteArray2(dataC, C.int(int32(size))))
	qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
	return tmpValue
}

func NewQByteArray3(size int, ch string) *QByteArray {
	var chC *C.char
	if ch != "" {
		chC = C.CString(ch)
		defer C.free(unsafe.Pointer(chC))
	}
	tmpValue := NewQByteArrayFromPointer(C.QByteArray_NewQByteArray3(C.int(int32(size)), chC))
	qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
	return tmpValue
}

func NewQByteArray4(other QByteArray_ITF) *QByteArray {
	tmpValue := NewQByteArrayFromPointer(C.QByteArray_NewQByteArray4(PointerFromQByteArray(other)))
	qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *QByteArray) Append(ba QByteArray_ITF) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Append(ptr.Pointer(), PointerFromQByteArray(ba)))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Append2(ch string) *QByteArray {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Append2(ptr.Pointer(), chC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Append3(count int, ch string) *QByteArray {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Append3(ptr.Pointer(), C.int(int32(count)), chC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Append4(str string) *QByteArray {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Append4(ptr.Pointer(), strC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Append5(str string, l int) *QByteArray {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Append5(ptr.Pointer(), strC, C.int(int32(l))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) At(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QByteArray_At(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QByteArray) Back() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QByteArray_Back(ptr.Pointer()))
	}
	return ""
}

func (ptr *QByteArray) Clear() {
	if ptr.Pointer() != nil {
		C.QByteArray_Clear(ptr.Pointer())
	}
}

func (ptr *QByteArray) Count(ba QByteArray_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QByteArray_Count(ptr.Pointer(), PointerFromQByteArray(ba))))
	}
	return 0
}

func (ptr *QByteArray) Count2(ch string) int {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int(int32(C.QByteArray_Count2(ptr.Pointer(), chC)))
	}
	return 0
}

func (ptr *QByteArray) Count3(str string) int {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		return int(int32(C.QByteArray_Count3(ptr.Pointer(), strC)))
	}
	return 0
}

func (ptr *QByteArray) Count4() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QByteArray_Count4(ptr.Pointer())))
	}
	return 0
}

func (ptr *QByteArray) Data() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QByteArray_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QByteArray) Data2() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QByteArray_Data2(ptr.Pointer()))
	}
	return ""
}

func (ptr *QByteArray) IndexOf(ba QByteArray_ITF, from int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QByteArray_IndexOf(ptr.Pointer(), PointerFromQByteArray(ba), C.int(int32(from)))))
	}
	return 0
}

func (ptr *QByteArray) IndexOf2(ch string, from int) int {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		return int(int32(C.QByteArray_IndexOf2(ptr.Pointer(), chC, C.int(int32(from)))))
	}
	return 0
}

func (ptr *QByteArray) IndexOf3(str string, from int) int {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		return int(int32(C.QByteArray_IndexOf3(ptr.Pointer(), strC, C.int(int32(from)))))
	}
	return 0
}

func (ptr *QByteArray) Insert(i int, ba QByteArray_ITF) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Insert(ptr.Pointer(), C.int(int32(i)), PointerFromQByteArray(ba)))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Insert2(i int, ch string) *QByteArray {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Insert2(ptr.Pointer(), C.int(int32(i)), chC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Insert3(i int, count int, ch string) *QByteArray {
	if ptr.Pointer() != nil {
		var chC *C.char
		if ch != "" {
			chC = C.CString(ch)
			defer C.free(unsafe.Pointer(chC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Insert3(ptr.Pointer(), C.int(int32(i)), C.int(int32(count)), chC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Insert4(i int, str string) *QByteArray {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Insert4(ptr.Pointer(), C.int(int32(i)), strC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Insert5(i int, str string, l int) *QByteArray {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Insert5(ptr.Pointer(), C.int(int32(i)), strC, C.int(int32(l))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QByteArray_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QByteArray) Left(l int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Left(ptr.Pointer(), C.int(int32(l))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QByteArray_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QByteArray) Remove(pos int, l int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Remove(ptr.Pointer(), C.int(int32(pos)), C.int(int32(l))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace(pos int, l int, after QByteArray_ITF) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace(ptr.Pointer(), C.int(int32(pos)), C.int(int32(l)), PointerFromQByteArray(after)))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace2(pos int, l int, after string) *QByteArray {
	if ptr.Pointer() != nil {
		var afterC *C.char
		if after != "" {
			afterC = C.CString(after)
			defer C.free(unsafe.Pointer(afterC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace2(ptr.Pointer(), C.int(int32(pos)), C.int(int32(l)), afterC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace3(pos int, l int, after string, alen int) *QByteArray {
	if ptr.Pointer() != nil {
		var afterC *C.char
		if after != "" {
			afterC = C.CString(after)
			defer C.free(unsafe.Pointer(afterC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace3(ptr.Pointer(), C.int(int32(pos)), C.int(int32(l)), afterC, C.int(int32(alen))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace4(before string, after string) *QByteArray {
	if ptr.Pointer() != nil {
		var beforeC *C.char
		if before != "" {
			beforeC = C.CString(before)
			defer C.free(unsafe.Pointer(beforeC))
		}
		var afterC *C.char
		if after != "" {
			afterC = C.CString(after)
			defer C.free(unsafe.Pointer(afterC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace4(ptr.Pointer(), beforeC, afterC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace5(before string, after QByteArray_ITF) *QByteArray {
	if ptr.Pointer() != nil {
		var beforeC *C.char
		if before != "" {
			beforeC = C.CString(before)
			defer C.free(unsafe.Pointer(beforeC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace5(ptr.Pointer(), beforeC, PointerFromQByteArray(after)))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace6(before string, after string) *QByteArray {
	if ptr.Pointer() != nil {
		var beforeC *C.char
		if before != "" {
			beforeC = C.CString(before)
			defer C.free(unsafe.Pointer(beforeC))
		}
		var afterC *C.char
		if after != "" {
			afterC = C.CString(after)
			defer C.free(unsafe.Pointer(afterC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace6(ptr.Pointer(), beforeC, afterC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace7(before string, bsize int, after string, asize int) *QByteArray {
	if ptr.Pointer() != nil {
		var beforeC *C.char
		if before != "" {
			beforeC = C.CString(before)
			defer C.free(unsafe.Pointer(beforeC))
		}
		var afterC *C.char
		if after != "" {
			afterC = C.CString(after)
			defer C.free(unsafe.Pointer(afterC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace7(ptr.Pointer(), beforeC, C.int(int32(bsize)), afterC, C.int(int32(asize))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace8(before QByteArray_ITF, after QByteArray_ITF) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace8(ptr.Pointer(), PointerFromQByteArray(before), PointerFromQByteArray(after)))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace9(before QByteArray_ITF, after string) *QByteArray {
	if ptr.Pointer() != nil {
		var afterC *C.char
		if after != "" {
			afterC = C.CString(after)
			defer C.free(unsafe.Pointer(afterC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace9(ptr.Pointer(), PointerFromQByteArray(before), afterC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace10(before string, after QByteArray_ITF) *QByteArray {
	if ptr.Pointer() != nil {
		var beforeC *C.char
		if before != "" {
			beforeC = C.CString(before)
			defer C.free(unsafe.Pointer(beforeC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace10(ptr.Pointer(), beforeC, PointerFromQByteArray(after)))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Replace11(before string, after string) *QByteArray {
	if ptr.Pointer() != nil {
		var beforeC *C.char
		if before != "" {
			beforeC = C.CString(before)
			defer C.free(unsafe.Pointer(beforeC))
		}
		var afterC *C.char
		if after != "" {
			afterC = C.CString(after)
			defer C.free(unsafe.Pointer(afterC))
		}
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Replace11(ptr.Pointer(), beforeC, afterC))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Resize(size int) {
	if ptr.Pointer() != nil {
		C.QByteArray_Resize(ptr.Pointer(), C.int(int32(size)))
	}
}

func (ptr *QByteArray) Right(l int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_Right(ptr.Pointer(), C.int(int32(l))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QByteArray_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QByteArray) Split(sep string) []*QByteArray {
	if ptr.Pointer() != nil {
		var sepC *C.char
		if sep != "" {
			sepC = C.CString(sep)
			defer C.free(unsafe.Pointer(sepC))
		}
		return func(l C.struct_QtCore_PackedList) []*QByteArray {
			out := make([]*QByteArray, int(l.len))
			tmpList := NewQByteArrayFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__split_atList(i)
			}
			return out
		}(C.QByteArray_Split(ptr.Pointer(), sepC))
	}
	return make([]*QByteArray, 0)
}

func (ptr *QByteArray) ToInt(ok *bool, base int) int {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return int(int32(C.QByteArray_ToInt(ptr.Pointer(), &okC, C.int(int32(base)))))
	}
	return 0
}

func (ptr *QByteArray) ToLower() *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QByteArray_ToLower(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) DestroyQByteArray() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QByteArray_DestroyQByteArray(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QByteArray) __split_atList(i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QByteArray___split_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QByteArray) __split_setList(i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QByteArray___split_setList(ptr.Pointer(), PointerFromQByteArray(i))
	}
}

func (ptr *QByteArray) __split_newList() unsafe.Pointer {
	return C.QByteArray___split_newList(ptr.Pointer())
}

type QByteRef struct {
	ptr unsafe.Pointer
}

type QByteRef_ITF interface {
	QByteRef_PTR() *QByteRef
}

func (ptr *QByteRef) QByteRef_PTR() *QByteRef {
	return ptr
}

func (ptr *QByteRef) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QByteRef) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQByteRef(ptr QByteRef_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QByteRef_PTR().Pointer()
	}
	return nil
}

func NewQByteRefFromPointer(ptr unsafe.Pointer) (n *QByteRef) {
	n = new(QByteRef)
	n.SetPointer(ptr)
	return
}
func (ptr *QByteRef) DestroyQByteRef() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QCalendar struct {
	ptr unsafe.Pointer
}

type QCalendar_ITF interface {
	QCalendar_PTR() *QCalendar
}

func (ptr *QCalendar) QCalendar_PTR() *QCalendar {
	return ptr
}

func (ptr *QCalendar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCalendar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCalendar(ptr QCalendar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCalendar_PTR().Pointer()
	}
	return nil
}

func NewQCalendarFromPointer(ptr unsafe.Pointer) (n *QCalendar) {
	n = new(QCalendar)
	n.SetPointer(ptr)
	return
}
func (ptr *QCalendar) DestroyQCalendar() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QCalendar__System
//QCalendar::System
type QCalendar__System int64

const (
	QCalendar__Gregorian    QCalendar__System = QCalendar__System(0)
	QCalendar__Julian       QCalendar__System = QCalendar__System(8)
	QCalendar__Milankovic   QCalendar__System = QCalendar__System(9)
	QCalendar__Jalali       QCalendar__System = QCalendar__System(10)
	QCalendar__IslamicCivil QCalendar__System = QCalendar__System(11)
	QCalendar__Last         QCalendar__System = QCalendar__System(11)
	QCalendar__User         QCalendar__System = QCalendar__System(-1)
)

func NewQCalendar() *QCalendar {
	tmpValue := NewQCalendarFromPointer(C.QCalendar_NewQCalendar())
	qt.SetFinalizer(tmpValue, (*QCalendar).DestroyQCalendar)
	return tmpValue
}

func NewQCalendar2(system QCalendar__System) *QCalendar {
	tmpValue := NewQCalendarFromPointer(C.QCalendar_NewQCalendar2(C.longlong(system)))
	qt.SetFinalizer(tmpValue, (*QCalendar).DestroyQCalendar)
	return tmpValue
}

func NewQCalendar3(name QLatin1String_ITF) *QCalendar {
	tmpValue := NewQCalendarFromPointer(C.QCalendar_NewQCalendar3(PointerFromQLatin1String(name)))
	qt.SetFinalizer(tmpValue, (*QCalendar).DestroyQCalendar)
	return tmpValue
}

func NewQCalendar4(name QStringView_ITF) *QCalendar {
	tmpValue := NewQCalendarFromPointer(C.QCalendar_NewQCalendar4(PointerFromQStringView(name)))
	qt.SetFinalizer(tmpValue, (*QCalendar).DestroyQCalendar)
	return tmpValue
}

func (ptr *QCalendar) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QCalendar_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCalendar) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCalendar_Name(ptr.Pointer()))
	}
	return ""
}

// QCborError::Code
//
//go:generate stringer -type=QCborError__Code
type QCborError__Code int64

const (
	QCborError__UnknownError      QCborError__Code = QCborError__Code(1)
	QCborError__AdvancePastEnd    QCborError__Code = QCborError__Code(3)
	QCborError__InputOutputError  QCborError__Code = QCborError__Code(4)
	QCborError__GarbageAtEnd      QCborError__Code = QCborError__Code(256)
	QCborError__EndOfFile         QCborError__Code = QCborError__Code(257)
	QCborError__UnexpectedBreak   QCborError__Code = QCborError__Code(258)
	QCborError__UnknownType       QCborError__Code = QCborError__Code(259)
	QCborError__IllegalType       QCborError__Code = QCborError__Code(260)
	QCborError__IllegalNumber     QCborError__Code = QCborError__Code(261)
	QCborError__IllegalSimpleType QCborError__Code = QCborError__Code(262)
	QCborError__InvalidUtf8String QCborError__Code = QCborError__Code(516)
	QCborError__DataTooLarge      QCborError__Code = QCborError__Code(1024)
	QCborError__NestingTooDeep    QCborError__Code = QCborError__Code(1025)
	QCborError__UnsupportedType   QCborError__Code = QCborError__Code(1026)
	QCborError__NoError           QCborError__Code = QCborError__Code(0)
)

// QCborStreamReader::StringResultCode
//
//go:generate stringer -type=QCborStreamReader__StringResultCode
type QCborStreamReader__StringResultCode int64

const (
	QCborStreamReader__EndOfString QCborStreamReader__StringResultCode = QCborStreamReader__StringResultCode(0)
	QCborStreamReader__Ok          QCborStreamReader__StringResultCode = QCborStreamReader__StringResultCode(1)
	QCborStreamReader__Error       QCborStreamReader__StringResultCode = QCborStreamReader__StringResultCode(-1)
)

// QCborStreamReader::Type
//
//go:generate stringer -type=QCborStreamReader__Type
type QCborStreamReader__Type int64

const (
	QCborStreamReader__UnsignedInteger QCborStreamReader__Type = QCborStreamReader__Type(0x00)
	QCborStreamReader__NegativeInteger QCborStreamReader__Type = QCborStreamReader__Type(0x20)
	QCborStreamReader__ByteString      QCborStreamReader__Type = QCborStreamReader__Type(0x40)
	QCborStreamReader__ByteArray       QCborStreamReader__Type = QCborStreamReader__Type(QCborStreamReader__ByteString)
	QCborStreamReader__TextString      QCborStreamReader__Type = QCborStreamReader__Type(0x60)
	QCborStreamReader__String          QCborStreamReader__Type = QCborStreamReader__Type(QCborStreamReader__TextString)
	QCborStreamReader__Array           QCborStreamReader__Type = QCborStreamReader__Type(0x80)
	QCborStreamReader__Map             QCborStreamReader__Type = QCborStreamReader__Type(0xa0)
	QCborStreamReader__Tag             QCborStreamReader__Type = QCborStreamReader__Type(0xc0)
	QCborStreamReader__SimpleType      QCborStreamReader__Type = QCborStreamReader__Type(0xe0)
	QCborStreamReader__HalfFloat       QCborStreamReader__Type = QCborStreamReader__Type(0xf9)
	QCborStreamReader__Float16         QCborStreamReader__Type = QCborStreamReader__Type(QCborStreamReader__HalfFloat)
	QCborStreamReader__Float           QCborStreamReader__Type = QCborStreamReader__Type(0xfa)
	QCborStreamReader__Double          QCborStreamReader__Type = QCborStreamReader__Type(0xfb)
	QCborStreamReader__Invalid         QCborStreamReader__Type = QCborStreamReader__Type(0xff)
)

// QCborValue::DiagnosticNotationOption
//
//go:generate stringer -type=QCborValue__DiagnosticNotationOption
type QCborValue__DiagnosticNotationOption int64

const (
	QCborValue__Compact        QCborValue__DiagnosticNotationOption = QCborValue__DiagnosticNotationOption(0x00)
	QCborValue__LineWrapped    QCborValue__DiagnosticNotationOption = QCborValue__DiagnosticNotationOption(0x01)
	QCborValue__ExtendedFormat QCborValue__DiagnosticNotationOption = QCborValue__DiagnosticNotationOption(0x02)
)

// QCborValue::EncodingOption
//
//go:generate stringer -type=QCborValue__EncodingOption
type QCborValue__EncodingOption int64

const (
	QCborValue__SortKeysInMaps   QCborValue__EncodingOption = QCborValue__EncodingOption(0x01)
	QCborValue__UseFloat         QCborValue__EncodingOption = QCborValue__EncodingOption(0x02)
	QCborValue__UseFloat16       QCborValue__EncodingOption = QCborValue__EncodingOption(QCborValue__UseFloat | 0x04)
	QCborValue__UseIntegers      QCborValue__EncodingOption = QCborValue__EncodingOption(0x08)
	QCborValue__NoTransformation QCborValue__EncodingOption = QCborValue__EncodingOption(0)
)

// QCborValue::Type
//
//go:generate stringer -type=QCborValue__Type
type QCborValue__Type int64

var (
	QCborValue__Integer           QCborValue__Type = QCborValue__Type(0x00)
	QCborValue__ByteArray         QCborValue__Type = QCborValue__Type(0x40)
	QCborValue__String            QCborValue__Type = QCborValue__Type(0x60)
	QCborValue__Array             QCborValue__Type = QCborValue__Type(0x80)
	QCborValue__Map               QCborValue__Type = QCborValue__Type(0xa0)
	QCborValue__Tag               QCborValue__Type = QCborValue__Type(0xc0)
	QCborValue__SimpleType        QCborValue__Type = QCborValue__Type(0x100)
	QCborValue__False             QCborValue__Type = QCborValue__Type(C.QCborValue_False_Type())
	QCborValue__True              QCborValue__Type = QCborValue__Type(C.QCborValue_True_Type())
	QCborValue__Null              QCborValue__Type = QCborValue__Type(C.QCborValue_Null_Type())
	QCborValue__Undefined         QCborValue__Type = QCborValue__Type(C.QCborValue_Undefined_Type())
	QCborValue__Double            QCborValue__Type = QCborValue__Type(0x202)
	QCborValue__DateTime          QCborValue__Type = QCborValue__Type(0x10000)
	QCborValue__Url               QCborValue__Type = QCborValue__Type(0x10020)
	QCborValue__RegularExpression QCborValue__Type = QCborValue__Type(0x10023)
	QCborValue__Uuid              QCborValue__Type = QCborValue__Type(0x10025)
	QCborValue__Invalid           QCborValue__Type = QCborValue__Type(-1)
)

type QChar struct {
	ptr unsafe.Pointer
}

type QChar_ITF interface {
	QChar_PTR() *QChar
}

func (ptr *QChar) QChar_PTR() *QChar {
	return ptr
}

func (ptr *QChar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QChar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQChar(ptr QChar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QChar_PTR().Pointer()
	}
	return nil
}

func NewQCharFromPointer(ptr unsafe.Pointer) (n *QChar) {
	n = new(QChar)
	n.SetPointer(ptr)
	return
}
func (ptr *QChar) DestroyQChar() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QChar__Category
//QChar::Category
type QChar__Category int64

const (
	QChar__Mark_NonSpacing          QChar__Category = QChar__Category(0)
	QChar__Mark_SpacingCombining    QChar__Category = QChar__Category(1)
	QChar__Mark_Enclosing           QChar__Category = QChar__Category(2)
	QChar__Number_DecimalDigit      QChar__Category = QChar__Category(3)
	QChar__Number_Letter            QChar__Category = QChar__Category(4)
	QChar__Number_Other             QChar__Category = QChar__Category(5)
	QChar__Separator_Space          QChar__Category = QChar__Category(6)
	QChar__Separator_Line           QChar__Category = QChar__Category(7)
	QChar__Separator_Paragraph      QChar__Category = QChar__Category(8)
	QChar__Other_Control            QChar__Category = QChar__Category(9)
	QChar__Other_Format             QChar__Category = QChar__Category(10)
	QChar__Other_Surrogate          QChar__Category = QChar__Category(11)
	QChar__Other_PrivateUse         QChar__Category = QChar__Category(12)
	QChar__Other_NotAssigned        QChar__Category = QChar__Category(13)
	QChar__Letter_Uppercase         QChar__Category = QChar__Category(14)
	QChar__Letter_Lowercase         QChar__Category = QChar__Category(15)
	QChar__Letter_Titlecase         QChar__Category = QChar__Category(16)
	QChar__Letter_Modifier          QChar__Category = QChar__Category(17)
	QChar__Letter_Other             QChar__Category = QChar__Category(18)
	QChar__Punctuation_Connector    QChar__Category = QChar__Category(19)
	QChar__Punctuation_Dash         QChar__Category = QChar__Category(20)
	QChar__Punctuation_Open         QChar__Category = QChar__Category(21)
	QChar__Punctuation_Close        QChar__Category = QChar__Category(22)
	QChar__Punctuation_InitialQuote QChar__Category = QChar__Category(23)
	QChar__Punctuation_FinalQuote   QChar__Category = QChar__Category(24)
	QChar__Punctuation_Other        QChar__Category = QChar__Category(25)
	QChar__Symbol_Math              QChar__Category = QChar__Category(26)
	QChar__Symbol_Currency          QChar__Category = QChar__Category(27)
	QChar__Symbol_Modifier          QChar__Category = QChar__Category(28)
	QChar__Symbol_Other             QChar__Category = QChar__Category(29)
)

// QChar::Decomposition
//
//go:generate stringer -type=QChar__Decomposition
type QChar__Decomposition int64

const (
	QChar__NoDecomposition QChar__Decomposition = QChar__Decomposition(0)
	QChar__Canonical       QChar__Decomposition = QChar__Decomposition(1)
	QChar__Font            QChar__Decomposition = QChar__Decomposition(2)
	QChar__NoBreak         QChar__Decomposition = QChar__Decomposition(3)
	QChar__Initial         QChar__Decomposition = QChar__Decomposition(4)
	QChar__Medial          QChar__Decomposition = QChar__Decomposition(5)
	QChar__Final           QChar__Decomposition = QChar__Decomposition(6)
	QChar__Isolated        QChar__Decomposition = QChar__Decomposition(7)
	QChar__Circle          QChar__Decomposition = QChar__Decomposition(8)
	QChar__Super           QChar__Decomposition = QChar__Decomposition(9)
	QChar__Sub             QChar__Decomposition = QChar__Decomposition(10)
	QChar__Vertical        QChar__Decomposition = QChar__Decomposition(11)
	QChar__Wide            QChar__Decomposition = QChar__Decomposition(12)
	QChar__Narrow          QChar__Decomposition = QChar__Decomposition(13)
	QChar__Small           QChar__Decomposition = QChar__Decomposition(14)
	QChar__Square          QChar__Decomposition = QChar__Decomposition(15)
	QChar__Compat          QChar__Decomposition = QChar__Decomposition(16)
	QChar__Fraction        QChar__Decomposition = QChar__Decomposition(17)
)

// QChar::Direction
//
//go:generate stringer -type=QChar__Direction
type QChar__Direction int64

const (
	QChar__DirL   QChar__Direction = QChar__Direction(0)
	QChar__DirR   QChar__Direction = QChar__Direction(1)
	QChar__DirEN  QChar__Direction = QChar__Direction(2)
	QChar__DirES  QChar__Direction = QChar__Direction(3)
	QChar__DirET  QChar__Direction = QChar__Direction(4)
	QChar__DirAN  QChar__Direction = QChar__Direction(5)
	QChar__DirCS  QChar__Direction = QChar__Direction(6)
	QChar__DirB   QChar__Direction = QChar__Direction(7)
	QChar__DirS   QChar__Direction = QChar__Direction(8)
	QChar__DirWS  QChar__Direction = QChar__Direction(9)
	QChar__DirON  QChar__Direction = QChar__Direction(10)
	QChar__DirLRE QChar__Direction = QChar__Direction(11)
	QChar__DirLRO QChar__Direction = QChar__Direction(12)
	QChar__DirAL  QChar__Direction = QChar__Direction(13)
	QChar__DirRLE QChar__Direction = QChar__Direction(14)
	QChar__DirRLO QChar__Direction = QChar__Direction(15)
	QChar__DirPDF QChar__Direction = QChar__Direction(16)
	QChar__DirNSM QChar__Direction = QChar__Direction(17)
	QChar__DirBN  QChar__Direction = QChar__Direction(18)
	QChar__DirLRI QChar__Direction = QChar__Direction(19)
	QChar__DirRLI QChar__Direction = QChar__Direction(20)
	QChar__DirFSI QChar__Direction = QChar__Direction(21)
	QChar__DirPDI QChar__Direction = QChar__Direction(22)
)

// QChar::JoiningType
//
//go:generate stringer -type=QChar__JoiningType
type QChar__JoiningType int64

const (
	QChar__Joining_None        QChar__JoiningType = QChar__JoiningType(0)
	QChar__Joining_Causing     QChar__JoiningType = QChar__JoiningType(1)
	QChar__Joining_Dual        QChar__JoiningType = QChar__JoiningType(2)
	QChar__Joining_Right       QChar__JoiningType = QChar__JoiningType(3)
	QChar__Joining_Left        QChar__JoiningType = QChar__JoiningType(4)
	QChar__Joining_Transparent QChar__JoiningType = QChar__JoiningType(5)
)

// QChar::Script
//
//go:generate stringer -type=QChar__Script
type QChar__Script int64

const (
	QChar__Script_Unknown               QChar__Script = QChar__Script(0)
	QChar__Script_Inherited             QChar__Script = QChar__Script(1)
	QChar__Script_Common                QChar__Script = QChar__Script(2)
	QChar__Script_Latin                 QChar__Script = QChar__Script(3)
	QChar__Script_Greek                 QChar__Script = QChar__Script(4)
	QChar__Script_Cyrillic              QChar__Script = QChar__Script(5)
	QChar__Script_Armenian              QChar__Script = QChar__Script(6)
	QChar__Script_Hebrew                QChar__Script = QChar__Script(7)
	QChar__Script_Arabic                QChar__Script = QChar__Script(8)
	QChar__Script_Syriac                QChar__Script = QChar__Script(9)
	QChar__Script_Thaana                QChar__Script = QChar__Script(10)
	QChar__Script_Devanagari            QChar__Script = QChar__Script(11)
	QChar__Script_Bengali               QChar__Script = QChar__Script(12)
	QChar__Script_Gurmukhi              QChar__Script = QChar__Script(13)
	QChar__Script_Gujarati              QChar__Script = QChar__Script(14)
	QChar__Script_Oriya                 QChar__Script = QChar__Script(15)
	QChar__Script_Tamil                 QChar__Script = QChar__Script(16)
	QChar__Script_Telugu                QChar__Script = QChar__Script(17)
	QChar__Script_Kannada               QChar__Script = QChar__Script(18)
	QChar__Script_Malayalam             QChar__Script = QChar__Script(19)
	QChar__Script_Sinhala               QChar__Script = QChar__Script(20)
	QChar__Script_Thai                  QChar__Script = QChar__Script(21)
	QChar__Script_Lao                   QChar__Script = QChar__Script(22)
	QChar__Script_Tibetan               QChar__Script = QChar__Script(23)
	QChar__Script_Myanmar               QChar__Script = QChar__Script(24)
	QChar__Script_Georgian              QChar__Script = QChar__Script(25)
	QChar__Script_Hangul                QChar__Script = QChar__Script(26)
	QChar__Script_Ethiopic              QChar__Script = QChar__Script(27)
	QChar__Script_Cherokee              QChar__Script = QChar__Script(28)
	QChar__Script_CanadianAboriginal    QChar__Script = QChar__Script(29)
	QChar__Script_Ogham                 QChar__Script = QChar__Script(30)
	QChar__Script_Runic                 QChar__Script = QChar__Script(31)
	QChar__Script_Khmer                 QChar__Script = QChar__Script(32)
	QChar__Script_Mongolian             QChar__Script = QChar__Script(33)
	QChar__Script_Hiragana              QChar__Script = QChar__Script(34)
	QChar__Script_Katakana              QChar__Script = QChar__Script(35)
	QChar__Script_Bopomofo              QChar__Script = QChar__Script(36)
	QChar__Script_Han                   QChar__Script = QChar__Script(37)
	QChar__Script_Yi                    QChar__Script = QChar__Script(38)
	QChar__Script_OldItalic             QChar__Script = QChar__Script(39)
	QChar__Script_Gothic                QChar__Script = QChar__Script(40)
	QChar__Script_Deseret               QChar__Script = QChar__Script(41)
	QChar__Script_Tagalog               QChar__Script = QChar__Script(42)
	QChar__Script_Hanunoo               QChar__Script = QChar__Script(43)
	QChar__Script_Buhid                 QChar__Script = QChar__Script(44)
	QChar__Script_Tagbanwa              QChar__Script = QChar__Script(45)
	QChar__Script_Coptic                QChar__Script = QChar__Script(46)
	QChar__Script_Limbu                 QChar__Script = QChar__Script(47)
	QChar__Script_TaiLe                 QChar__Script = QChar__Script(48)
	QChar__Script_LinearB               QChar__Script = QChar__Script(49)
	QChar__Script_Ugaritic              QChar__Script = QChar__Script(50)
	QChar__Script_Shavian               QChar__Script = QChar__Script(51)
	QChar__Script_Osmanya               QChar__Script = QChar__Script(52)
	QChar__Script_Cypriot               QChar__Script = QChar__Script(53)
	QChar__Script_Braille               QChar__Script = QChar__Script(54)
	QChar__Script_Buginese              QChar__Script = QChar__Script(55)
	QChar__Script_NewTaiLue             QChar__Script = QChar__Script(56)
	QChar__Script_Glagolitic            QChar__Script = QChar__Script(57)
	QChar__Script_Tifinagh              QChar__Script = QChar__Script(58)
	QChar__Script_SylotiNagri           QChar__Script = QChar__Script(59)
	QChar__Script_OldPersian            QChar__Script = QChar__Script(60)
	QChar__Script_Kharoshthi            QChar__Script = QChar__Script(61)
	QChar__Script_Balinese              QChar__Script = QChar__Script(62)
	QChar__Script_Cuneiform             QChar__Script = QChar__Script(63)
	QChar__Script_Phoenician            QChar__Script = QChar__Script(64)
	QChar__Script_PhagsPa               QChar__Script = QChar__Script(65)
	QChar__Script_Nko                   QChar__Script = QChar__Script(66)
	QChar__Script_Sundanese             QChar__Script = QChar__Script(67)
	QChar__Script_Lepcha                QChar__Script = QChar__Script(68)
	QChar__Script_OlChiki               QChar__Script = QChar__Script(69)
	QChar__Script_Vai                   QChar__Script = QChar__Script(70)
	QChar__Script_Saurashtra            QChar__Script = QChar__Script(71)
	QChar__Script_KayahLi               QChar__Script = QChar__Script(72)
	QChar__Script_Rejang                QChar__Script = QChar__Script(73)
	QChar__Script_Lycian                QChar__Script = QChar__Script(74)
	QChar__Script_Carian                QChar__Script = QChar__Script(75)
	QChar__Script_Lydian                QChar__Script = QChar__Script(76)
	QChar__Script_Cham                  QChar__Script = QChar__Script(77)
	QChar__Script_TaiTham               QChar__Script = QChar__Script(78)
	QChar__Script_TaiViet               QChar__Script = QChar__Script(79)
	QChar__Script_Avestan               QChar__Script = QChar__Script(80)
	QChar__Script_EgyptianHieroglyphs   QChar__Script = QChar__Script(81)
	QChar__Script_Samaritan             QChar__Script = QChar__Script(82)
	QChar__Script_Lisu                  QChar__Script = QChar__Script(83)
	QChar__Script_Bamum                 QChar__Script = QChar__Script(84)
	QChar__Script_Javanese              QChar__Script = QChar__Script(85)
	QChar__Script_MeeteiMayek           QChar__Script = QChar__Script(86)
	QChar__Script_ImperialAramaic       QChar__Script = QChar__Script(87)
	QChar__Script_OldSouthArabian       QChar__Script = QChar__Script(88)
	QChar__Script_InscriptionalParthian QChar__Script = QChar__Script(89)
	QChar__Script_InscriptionalPahlavi  QChar__Script = QChar__Script(90)
	QChar__Script_OldTurkic             QChar__Script = QChar__Script(91)
	QChar__Script_Kaithi                QChar__Script = QChar__Script(92)
	QChar__Script_Batak                 QChar__Script = QChar__Script(93)
	QChar__Script_Brahmi                QChar__Script = QChar__Script(94)
	QChar__Script_Mandaic               QChar__Script = QChar__Script(95)
	QChar__Script_Chakma                QChar__Script = QChar__Script(96)
	QChar__Script_MeroiticCursive       QChar__Script = QChar__Script(97)
	QChar__Script_MeroiticHieroglyphs   QChar__Script = QChar__Script(98)
	QChar__Script_Miao                  QChar__Script = QChar__Script(99)
	QChar__Script_Sharada               QChar__Script = QChar__Script(100)
	QChar__Script_SoraSompeng           QChar__Script = QChar__Script(101)
	QChar__Script_Takri                 QChar__Script = QChar__Script(102)
	QChar__Script_CaucasianAlbanian     QChar__Script = QChar__Script(103)
	QChar__Script_BassaVah              QChar__Script = QChar__Script(104)
	QChar__Script_Duployan              QChar__Script = QChar__Script(105)
	QChar__Script_Elbasan               QChar__Script = QChar__Script(106)
	QChar__Script_Grantha               QChar__Script = QChar__Script(107)
	QChar__Script_PahawhHmong           QChar__Script = QChar__Script(108)
	QChar__Script_Khojki                QChar__Script = QChar__Script(109)
	QChar__Script_LinearA               QChar__Script = QChar__Script(110)
	QChar__Script_Mahajani              QChar__Script = QChar__Script(111)
	QChar__Script_Manichaean            QChar__Script = QChar__Script(112)
	QChar__Script_MendeKikakui          QChar__Script = QChar__Script(113)
	QChar__Script_Modi                  QChar__Script = QChar__Script(114)
	QChar__Script_Mro                   QChar__Script = QChar__Script(115)
	QChar__Script_OldNorthArabian       QChar__Script = QChar__Script(116)
	QChar__Script_Nabataean             QChar__Script = QChar__Script(117)
	QChar__Script_Palmyrene             QChar__Script = QChar__Script(118)
	QChar__Script_PauCinHau             QChar__Script = QChar__Script(119)
	QChar__Script_OldPermic             QChar__Script = QChar__Script(120)
	QChar__Script_PsalterPahlavi        QChar__Script = QChar__Script(121)
	QChar__Script_Siddham               QChar__Script = QChar__Script(122)
	QChar__Script_Khudawadi             QChar__Script = QChar__Script(123)
	QChar__Script_Tirhuta               QChar__Script = QChar__Script(124)
	QChar__Script_WarangCiti            QChar__Script = QChar__Script(125)
	QChar__Script_Ahom                  QChar__Script = QChar__Script(126)
	QChar__Script_AnatolianHieroglyphs  QChar__Script = QChar__Script(127)
	QChar__Script_Hatran                QChar__Script = QChar__Script(128)
	QChar__Script_Multani               QChar__Script = QChar__Script(129)
	QChar__Script_OldHungarian          QChar__Script = QChar__Script(130)
	QChar__Script_SignWriting           QChar__Script = QChar__Script(131)
	QChar__Script_Adlam                 QChar__Script = QChar__Script(132)
	QChar__Script_Bhaiksuki             QChar__Script = QChar__Script(133)
	QChar__Script_Marchen               QChar__Script = QChar__Script(134)
	QChar__Script_Newa                  QChar__Script = QChar__Script(135)
	QChar__Script_Osage                 QChar__Script = QChar__Script(136)
	QChar__Script_Tangut                QChar__Script = QChar__Script(137)
	QChar__Script_MasaramGondi          QChar__Script = QChar__Script(138)
	QChar__Script_Nushu                 QChar__Script = QChar__Script(139)
	QChar__Script_Soyombo               QChar__Script = QChar__Script(140)
	QChar__Script_ZanabazarSquare       QChar__Script = QChar__Script(141)
	QChar__Script_Dogra                 QChar__Script = QChar__Script(142)
	QChar__Script_GunjalaGondi          QChar__Script = QChar__Script(143)
	QChar__Script_HanifiRohingya        QChar__Script = QChar__Script(144)
	QChar__Script_Makasar               QChar__Script = QChar__Script(145)
	QChar__Script_Medefaidrin           QChar__Script = QChar__Script(146)
	QChar__Script_OldSogdian            QChar__Script = QChar__Script(147)
	QChar__Script_Sogdian               QChar__Script = QChar__Script(148)
	QChar__Script_Elymaic               QChar__Script = QChar__Script(149)
	QChar__Script_Nandinagari           QChar__Script = QChar__Script(150)
	QChar__Script_NyiakengPuachueHmong  QChar__Script = QChar__Script(151)
	QChar__Script_Wancho                QChar__Script = QChar__Script(152)
	QChar__Script_Chorasmian            QChar__Script = QChar__Script(153)
	QChar__Script_DivesAkuru            QChar__Script = QChar__Script(154)
	QChar__Script_KhitanSmallScript     QChar__Script = QChar__Script(155)
	QChar__Script_Yezidi                QChar__Script = QChar__Script(156)
	QChar__ScriptCount                  QChar__Script = QChar__Script(157)
)

// QChar::SpecialCharacter
//
//go:generate stringer -type=QChar__SpecialCharacter
type QChar__SpecialCharacter int64

const (
	QChar__Null                       QChar__SpecialCharacter = QChar__SpecialCharacter(0x0000)
	QChar__Tabulation                 QChar__SpecialCharacter = QChar__SpecialCharacter(0x0009)
	QChar__LineFeed                   QChar__SpecialCharacter = QChar__SpecialCharacter(0x000a)
	QChar__FormFeed                   QChar__SpecialCharacter = QChar__SpecialCharacter(0x000c)
	QChar__CarriageReturn             QChar__SpecialCharacter = QChar__SpecialCharacter(0x000d)
	QChar__Space                      QChar__SpecialCharacter = QChar__SpecialCharacter(0x0020)
	QChar__Nbsp                       QChar__SpecialCharacter = QChar__SpecialCharacter(0x00a0)
	QChar__SoftHyphen                 QChar__SpecialCharacter = QChar__SpecialCharacter(0x00ad)
	QChar__ReplacementCharacter       QChar__SpecialCharacter = QChar__SpecialCharacter(0xfffd)
	QChar__ObjectReplacementCharacter QChar__SpecialCharacter = QChar__SpecialCharacter(0xfffc)
	QChar__ByteOrderMark              QChar__SpecialCharacter = QChar__SpecialCharacter(0xfeff)
	QChar__ByteOrderSwapped           QChar__SpecialCharacter = QChar__SpecialCharacter(0xfffe)
	QChar__ParagraphSeparator         QChar__SpecialCharacter = QChar__SpecialCharacter(0x2029)
	QChar__LineSeparator              QChar__SpecialCharacter = QChar__SpecialCharacter(0x2028)
	QChar__LastValidCodePoint         QChar__SpecialCharacter = QChar__SpecialCharacter(0x10ffff)
)

// QChar::UnicodeVersion
//
//go:generate stringer -type=QChar__UnicodeVersion
type QChar__UnicodeVersion int64

const (
	QChar__Unicode_Unassigned QChar__UnicodeVersion = QChar__UnicodeVersion(0)
	QChar__Unicode_1_1        QChar__UnicodeVersion = QChar__UnicodeVersion(1)
	QChar__Unicode_2_0        QChar__UnicodeVersion = QChar__UnicodeVersion(2)
	QChar__Unicode_2_1_2      QChar__UnicodeVersion = QChar__UnicodeVersion(3)
	QChar__Unicode_3_0        QChar__UnicodeVersion = QChar__UnicodeVersion(4)
	QChar__Unicode_3_1        QChar__UnicodeVersion = QChar__UnicodeVersion(5)
	QChar__Unicode_3_2        QChar__UnicodeVersion = QChar__UnicodeVersion(6)
	QChar__Unicode_4_0        QChar__UnicodeVersion = QChar__UnicodeVersion(7)
	QChar__Unicode_4_1        QChar__UnicodeVersion = QChar__UnicodeVersion(8)
	QChar__Unicode_5_0        QChar__UnicodeVersion = QChar__UnicodeVersion(9)
	QChar__Unicode_5_1        QChar__UnicodeVersion = QChar__UnicodeVersion(10)
	QChar__Unicode_5_2        QChar__UnicodeVersion = QChar__UnicodeVersion(11)
	QChar__Unicode_6_0        QChar__UnicodeVersion = QChar__UnicodeVersion(12)
	QChar__Unicode_6_1        QChar__UnicodeVersion = QChar__UnicodeVersion(13)
	QChar__Unicode_6_2        QChar__UnicodeVersion = QChar__UnicodeVersion(14)
	QChar__Unicode_6_3        QChar__UnicodeVersion = QChar__UnicodeVersion(15)
	QChar__Unicode_7_0        QChar__UnicodeVersion = QChar__UnicodeVersion(16)
	QChar__Unicode_8_0        QChar__UnicodeVersion = QChar__UnicodeVersion(17)
	QChar__Unicode_9_0        QChar__UnicodeVersion = QChar__UnicodeVersion(18)
	QChar__Unicode_10_0       QChar__UnicodeVersion = QChar__UnicodeVersion(19)
	QChar__Unicode_11_0       QChar__UnicodeVersion = QChar__UnicodeVersion(20)
	QChar__Unicode_12_0       QChar__UnicodeVersion = QChar__UnicodeVersion(21)
	QChar__Unicode_12_1       QChar__UnicodeVersion = QChar__UnicodeVersion(22)
	QChar__Unicode_13_0       QChar__UnicodeVersion = QChar__UnicodeVersion(23)
)

func NewQChar() *QChar {
	tmpValue := NewQCharFromPointer(C.QChar_NewQChar())
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func NewQChar2(code uint16) *QChar {
	tmpValue := NewQCharFromPointer(C.QChar_NewQChar2(C.ushort(code)))
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func NewQChar3(cell string, row string) *QChar {
	var cellC *C.char
	if cell != "" {
		cellC = C.CString(cell)
		defer C.free(unsafe.Pointer(cellC))
	}
	var rowC *C.char
	if row != "" {
		rowC = C.CString(row)
		defer C.free(unsafe.Pointer(rowC))
	}
	tmpValue := NewQCharFromPointer(C.QChar_NewQChar3(cellC, rowC))
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func NewQChar4(code int16) *QChar {
	tmpValue := NewQCharFromPointer(C.QChar_NewQChar4(C.short(code)))
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func NewQChar5(code uint) *QChar {
	tmpValue := NewQCharFromPointer(C.QChar_NewQChar5(C.uint(uint32(code))))
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func NewQChar6(code int) *QChar {
	tmpValue := NewQCharFromPointer(C.QChar_NewQChar6(C.int(int32(code))))
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func NewQChar7(ch QChar__SpecialCharacter) *QChar {
	tmpValue := NewQCharFromPointer(C.QChar_NewQChar7(C.longlong(ch)))
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func NewQChar8(ch QLatin1Char_ITF) *QChar {
	tmpValue := NewQCharFromPointer(C.QChar_NewQChar8(PointerFromQLatin1Char(ch)))
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func NewQChar10(ch string) *QChar {
	var chC *C.char
	if ch != "" {
		chC = C.CString(ch)
		defer C.free(unsafe.Pointer(chC))
	}
	tmpValue := NewQCharFromPointer(C.QChar_NewQChar10(chC))
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func NewQChar11(ch string) *QChar {
	var chC *C.char
	if ch != "" {
		chC = C.CString(ch)
		defer C.free(unsafe.Pointer(chC))
	}
	tmpValue := NewQCharFromPointer(C.QChar_NewQChar11(chC))
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func (ptr *QChar) Category() QChar__Category {
	if ptr.Pointer() != nil {
		return QChar__Category(C.QChar_Category(ptr.Pointer()))
	}
	return 0
}

func QChar_Category2(ucs4 uint) QChar__Category {
	return QChar__Category(C.QChar_QChar_Category2(C.uint(uint32(ucs4))))
}

func (ptr *QChar) Category2(ucs4 uint) QChar__Category {
	return QChar__Category(C.QChar_QChar_Category2(C.uint(uint32(ucs4))))
}

func (ptr *QChar) Cell() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QChar_Cell(ptr.Pointer()))
	}
	return ""
}

func (ptr *QChar) Direction() QChar__Direction {
	if ptr.Pointer() != nil {
		return QChar__Direction(C.QChar_Direction(ptr.Pointer()))
	}
	return 0
}

func QChar_Direction2(ucs4 uint) QChar__Direction {
	return QChar__Direction(C.QChar_QChar_Direction2(C.uint(uint32(ucs4))))
}

func (ptr *QChar) Direction2(ucs4 uint) QChar__Direction {
	return QChar__Direction(C.QChar_QChar_Direction2(C.uint(uint32(ucs4))))
}

func (ptr *QChar) Row() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QChar_Row(ptr.Pointer()))
	}
	return ""
}

func (ptr *QChar) ToLower() *QChar {
	if ptr.Pointer() != nil {
		tmpValue := NewQCharFromPointer(C.QChar_ToLower(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
		return tmpValue
	}
	return nil
}

func QChar_ToLower2(ucs4 uint) uint {
	return uint(uint32(C.QChar_QChar_ToLower2(C.uint(uint32(ucs4)))))
}

func (ptr *QChar) ToLower2(ucs4 uint) uint {
	return uint(uint32(C.QChar_QChar_ToLower2(C.uint(uint32(ucs4)))))
}

type QChildEvent struct {
	QEvent
}

type QChildEvent_ITF interface {
	QEvent_ITF
	QChildEvent_PTR() *QChildEvent
}

func (ptr *QChildEvent) QChildEvent_PTR() *QChildEvent {
	return ptr
}

func (ptr *QChildEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QChildEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQChildEvent(ptr QChildEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QChildEvent_PTR().Pointer()
	}
	return nil
}

func NewQChildEventFromPointer(ptr unsafe.Pointer) (n *QChildEvent) {
	n = new(QChildEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QChildEvent) DestroyQChildEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQChildEvent(ty QEvent__Type, child QObject_ITF) *QChildEvent {
	tmpValue := NewQChildEventFromPointer(C.QChildEvent_NewQChildEvent(C.longlong(ty), PointerFromQObject(child)))
	qt.SetFinalizer(tmpValue, (*QChildEvent).DestroyQChildEvent)
	return tmpValue
}

func (ptr *QChildEvent) Child() *QObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQObjectFromPointer(C.QChildEvent_Child(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

// QCommandLineOption::Flag
//
//go:generate stringer -type=QCommandLineOption__Flag
type QCommandLineOption__Flag int64

const (
	QCommandLineOption__HiddenFromHelp   QCommandLineOption__Flag = QCommandLineOption__Flag(0x1)
	QCommandLineOption__ShortOptionStyle QCommandLineOption__Flag = QCommandLineOption__Flag(0x2)
)

// QCommandLineParser::OptionsAfterPositionalArgumentsMode
//
//go:generate stringer -type=QCommandLineParser__OptionsAfterPositionalArgumentsMode
type QCommandLineParser__OptionsAfterPositionalArgumentsMode int64

const (
	QCommandLineParser__ParseAsOptions             QCommandLineParser__OptionsAfterPositionalArgumentsMode = QCommandLineParser__OptionsAfterPositionalArgumentsMode(0)
	QCommandLineParser__ParseAsPositionalArguments QCommandLineParser__OptionsAfterPositionalArgumentsMode = QCommandLineParser__OptionsAfterPositionalArgumentsMode(1)
)

// QCommandLineParser::SingleDashWordOptionMode
//
//go:generate stringer -type=QCommandLineParser__SingleDashWordOptionMode
type QCommandLineParser__SingleDashWordOptionMode int64

const (
	QCommandLineParser__ParseAsCompactedShortOptions QCommandLineParser__SingleDashWordOptionMode = QCommandLineParser__SingleDashWordOptionMode(0)
	QCommandLineParser__ParseAsLongOptions           QCommandLineParser__SingleDashWordOptionMode = QCommandLineParser__SingleDashWordOptionMode(1)
)

type QCoreApplication struct {
	QObject
}

type QCoreApplication_ITF interface {
	QObject_ITF
	QCoreApplication_PTR() *QCoreApplication
}

func (ptr *QCoreApplication) QCoreApplication_PTR() *QCoreApplication {
	return ptr
}

func (ptr *QCoreApplication) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QCoreApplication) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQCoreApplication(ptr QCoreApplication_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCoreApplication_PTR().Pointer()
	}
	return nil
}

func NewQCoreApplicationFromPointer(ptr unsafe.Pointer) (n *QCoreApplication) {
	n = new(QCoreApplication)
	n.SetPointer(ptr)
	return
}
func NewQCoreApplication(argc int, argv []string) *QCoreApplication {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	tmpValue := NewQCoreApplicationFromPointer(C.QCoreApplication_NewQCoreApplication(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QCoreApplication_ApplicationName() string {
	return cGoUnpackString(C.QCoreApplication_QCoreApplication_ApplicationName())
}

func (ptr *QCoreApplication) ApplicationName() string {
	return cGoUnpackString(C.QCoreApplication_QCoreApplication_ApplicationName())
}

func QCoreApplication_ApplicationVersion() string {
	return cGoUnpackString(C.QCoreApplication_QCoreApplication_ApplicationVersion())
}

func (ptr *QCoreApplication) ApplicationVersion() string {
	return cGoUnpackString(C.QCoreApplication_QCoreApplication_ApplicationVersion())
}

func QCoreApplication_Exec() int {
	return int(int32(C.QCoreApplication_QCoreApplication_Exec()))
}

func (ptr *QCoreApplication) Exec() int {
	return int(int32(C.QCoreApplication_QCoreApplication_Exec()))
}

//export callbackQCoreApplication_Quit
func callbackQCoreApplication_Quit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "quit"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQCoreApplicationFromPointer(ptr).QuitDefault()
	}
}

func (ptr *QCoreApplication) ConnectQuit(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "quit"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "quit", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "quit", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCoreApplication) DisconnectQuit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "quit")
	}
}

func (ptr *QCoreApplication) Quit() {
	if ptr.Pointer() != nil {
		C.QCoreApplication_Quit(ptr.Pointer())
	}
}

func (ptr *QCoreApplication) QuitDefault() {
	if ptr.Pointer() != nil {
		C.QCoreApplication_QuitDefault(ptr.Pointer())
	}
}

func QCoreApplication_SetApplicationName(application string) {
	var applicationC *C.char
	if application != "" {
		applicationC = C.CString(application)
		defer C.free(unsafe.Pointer(applicationC))
	}
	C.QCoreApplication_QCoreApplication_SetApplicationName(C.struct_QtCore_PackedString{data: applicationC, len: C.longlong(len(application))})
}

func (ptr *QCoreApplication) SetApplicationName(application string) {
	var applicationC *C.char
	if application != "" {
		applicationC = C.CString(application)
		defer C.free(unsafe.Pointer(applicationC))
	}
	C.QCoreApplication_QCoreApplication_SetApplicationName(C.struct_QtCore_PackedString{data: applicationC, len: C.longlong(len(application))})
}

func QCoreApplication_SetApplicationVersion(version string) {
	var versionC *C.char
	if version != "" {
		versionC = C.CString(version)
		defer C.free(unsafe.Pointer(versionC))
	}
	C.QCoreApplication_QCoreApplication_SetApplicationVersion(C.struct_QtCore_PackedString{data: versionC, len: C.longlong(len(version))})
}

func (ptr *QCoreApplication) SetApplicationVersion(version string) {
	var versionC *C.char
	if version != "" {
		versionC = C.CString(version)
		defer C.free(unsafe.Pointer(versionC))
	}
	C.QCoreApplication_QCoreApplication_SetApplicationVersion(C.struct_QtCore_PackedString{data: versionC, len: C.longlong(len(version))})
}

//export callbackQCoreApplication_DestroyQCoreApplication
func callbackQCoreApplication_DestroyQCoreApplication(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QCoreApplication"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQCoreApplicationFromPointer(ptr).DestroyQCoreApplicationDefault()
	}
}

func (ptr *QCoreApplication) ConnectDestroyQCoreApplication(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QCoreApplication"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QCoreApplication", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QCoreApplication", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCoreApplication) DisconnectDestroyQCoreApplication() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QCoreApplication")
	}
}

func (ptr *QCoreApplication) DestroyQCoreApplication() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCoreApplication_DestroyQCoreApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCoreApplication) DestroyQCoreApplicationDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCoreApplication_DestroyQCoreApplicationDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QCryptographicHash::Algorithm
//
//go:generate stringer -type=QCryptographicHash__Algorithm
type QCryptographicHash__Algorithm int64

const (
	QCryptographicHash__Md4          QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(0)
	QCryptographicHash__Md5          QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(1)
	QCryptographicHash__Sha1         QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(2)
	QCryptographicHash__Sha224       QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(3)
	QCryptographicHash__Sha256       QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(4)
	QCryptographicHash__Sha384       QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(5)
	QCryptographicHash__Sha512       QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(6)
	QCryptographicHash__Keccak_224   QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(7)
	QCryptographicHash__Keccak_256   QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(8)
	QCryptographicHash__Keccak_384   QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(9)
	QCryptographicHash__Keccak_512   QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(10)
	QCryptographicHash__RealSha3_224 QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(11)
	QCryptographicHash__RealSha3_256 QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(12)
	QCryptographicHash__RealSha3_384 QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(13)
	QCryptographicHash__RealSha3_512 QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(14)
	QCryptographicHash__Sha3_224     QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(QCryptographicHash__RealSha3_224)
	QCryptographicHash__Sha3_256     QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(QCryptographicHash__RealSha3_256)
	QCryptographicHash__Sha3_384     QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(QCryptographicHash__RealSha3_384)
	QCryptographicHash__Sha3_512     QCryptographicHash__Algorithm = QCryptographicHash__Algorithm(QCryptographicHash__RealSha3_512)
)

type QDataStream struct {
	ptr unsafe.Pointer
}

type QDataStream_ITF interface {
	QDataStream_PTR() *QDataStream
}

func (ptr *QDataStream) QDataStream_PTR() *QDataStream {
	return ptr
}

func (ptr *QDataStream) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDataStream) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDataStream(ptr QDataStream_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDataStream_PTR().Pointer()
	}
	return nil
}

func NewQDataStreamFromPointer(ptr unsafe.Pointer) (n *QDataStream) {
	n = new(QDataStream)
	n.SetPointer(ptr)
	return
}

// QDataStream::ByteOrder
//
//go:generate stringer -type=QDataStream__ByteOrder
type QDataStream__ByteOrder int64

const (
	QDataStream__BigEndian    QDataStream__ByteOrder = QDataStream__ByteOrder(QSysInfo__BigEndian)
	QDataStream__LittleEndian QDataStream__ByteOrder = QDataStream__ByteOrder(QSysInfo__LittleEndian)
)

// QDataStream::FloatingPointPrecision
//
//go:generate stringer -type=QDataStream__FloatingPointPrecision
type QDataStream__FloatingPointPrecision int64

const (
	QDataStream__SinglePrecision QDataStream__FloatingPointPrecision = QDataStream__FloatingPointPrecision(0)
	QDataStream__DoublePrecision QDataStream__FloatingPointPrecision = QDataStream__FloatingPointPrecision(1)
)

// QDataStream::Status
//
//go:generate stringer -type=QDataStream__Status
type QDataStream__Status int64

const (
	QDataStream__Ok              QDataStream__Status = QDataStream__Status(0)
	QDataStream__ReadPastEnd     QDataStream__Status = QDataStream__Status(1)
	QDataStream__ReadCorruptData QDataStream__Status = QDataStream__Status(2)
	QDataStream__WriteFailed     QDataStream__Status = QDataStream__Status(3)
)

// QDataStream::Version
//
//go:generate stringer -type=QDataStream__Version
type QDataStream__Version int64

const (
	QDataStream__Qt_1_0                    QDataStream__Version = QDataStream__Version(1)
	QDataStream__Qt_2_0                    QDataStream__Version = QDataStream__Version(2)
	QDataStream__Qt_2_1                    QDataStream__Version = QDataStream__Version(3)
	QDataStream__Qt_3_0                    QDataStream__Version = QDataStream__Version(4)
	QDataStream__Qt_3_1                    QDataStream__Version = QDataStream__Version(5)
	QDataStream__Qt_3_3                    QDataStream__Version = QDataStream__Version(6)
	QDataStream__Qt_4_0                    QDataStream__Version = QDataStream__Version(7)
	QDataStream__Qt_4_1                    QDataStream__Version = QDataStream__Version(QDataStream__Qt_4_0)
	QDataStream__Qt_4_2                    QDataStream__Version = QDataStream__Version(8)
	QDataStream__Qt_4_3                    QDataStream__Version = QDataStream__Version(9)
	QDataStream__Qt_4_4                    QDataStream__Version = QDataStream__Version(10)
	QDataStream__Qt_4_5                    QDataStream__Version = QDataStream__Version(11)
	QDataStream__Qt_4_6                    QDataStream__Version = QDataStream__Version(12)
	QDataStream__Qt_4_7                    QDataStream__Version = QDataStream__Version(QDataStream__Qt_4_6)
	QDataStream__Qt_4_8                    QDataStream__Version = QDataStream__Version(QDataStream__Qt_4_7)
	QDataStream__Qt_4_9                    QDataStream__Version = QDataStream__Version(QDataStream__Qt_4_8)
	QDataStream__Qt_5_0                    QDataStream__Version = QDataStream__Version(13)
	QDataStream__Qt_5_1                    QDataStream__Version = QDataStream__Version(14)
	QDataStream__Qt_5_2                    QDataStream__Version = QDataStream__Version(15)
	QDataStream__Qt_5_3                    QDataStream__Version = QDataStream__Version(QDataStream__Qt_5_2)
	QDataStream__Qt_5_4                    QDataStream__Version = QDataStream__Version(16)
	QDataStream__Qt_5_5                    QDataStream__Version = QDataStream__Version(QDataStream__Qt_5_4)
	QDataStream__Qt_5_6                    QDataStream__Version = QDataStream__Version(17)
	QDataStream__Qt_5_7                    QDataStream__Version = QDataStream__Version(QDataStream__Qt_5_6)
	QDataStream__Qt_5_8                    QDataStream__Version = QDataStream__Version(QDataStream__Qt_5_7)
	QDataStream__Qt_5_9                    QDataStream__Version = QDataStream__Version(QDataStream__Qt_5_8)
	QDataStream__Qt_5_10                   QDataStream__Version = QDataStream__Version(QDataStream__Qt_5_9)
	QDataStream__Qt_5_11                   QDataStream__Version = QDataStream__Version(QDataStream__Qt_5_10)
	QDataStream__Qt_5_12                   QDataStream__Version = QDataStream__Version(18)
	QDataStream__Qt_5_13                   QDataStream__Version = QDataStream__Version(19)
	QDataStream__Qt_5_14                   QDataStream__Version = QDataStream__Version(QDataStream__Qt_5_13)
	QDataStream__Qt_5_15                   QDataStream__Version = QDataStream__Version(QDataStream__Qt_5_14)
	QDataStream__Qt_DefaultCompiledVersion QDataStream__Version = QDataStream__Version(QDataStream__Qt_5_15)
)

func NewQDataStream() *QDataStream {
	tmpValue := NewQDataStreamFromPointer(C.QDataStream_NewQDataStream())
	qt.SetFinalizer(tmpValue, (*QDataStream).DestroyQDataStream)
	return tmpValue
}

func NewQDataStream2(d QIODevice_ITF) *QDataStream {
	tmpValue := NewQDataStreamFromPointer(C.QDataStream_NewQDataStream2(PointerFromQIODevice(d)))
	qt.SetFinalizer(tmpValue, (*QDataStream).DestroyQDataStream)
	return tmpValue
}

func NewQDataStream3(a QByteArray_ITF, mode QIODevice__OpenModeFlag) *QDataStream {
	tmpValue := NewQDataStreamFromPointer(C.QDataStream_NewQDataStream3(PointerFromQByteArray(a), C.longlong(mode)))
	qt.SetFinalizer(tmpValue, (*QDataStream).DestroyQDataStream)
	return tmpValue
}

func NewQDataStream4(a QByteArray_ITF) *QDataStream {
	tmpValue := NewQDataStreamFromPointer(C.QDataStream_NewQDataStream4(PointerFromQByteArray(a)))
	qt.SetFinalizer(tmpValue, (*QDataStream).DestroyQDataStream)
	return tmpValue
}

func (ptr *QDataStream) Device() *QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := NewQIODeviceFromPointer(C.QDataStream_Device(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDataStream) Status() QDataStream__Status {
	if ptr.Pointer() != nil {
		return QDataStream__Status(C.QDataStream_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDataStream) Version() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDataStream_Version(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDataStream) DestroyQDataStream() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDataStream_DestroyQDataStream(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDate struct {
	ptr unsafe.Pointer
}

type QDate_ITF interface {
	QDate_PTR() *QDate
}

func (ptr *QDate) QDate_PTR() *QDate {
	return ptr
}

func (ptr *QDate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDate(ptr QDate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDate_PTR().Pointer()
	}
	return nil
}

func NewQDateFromPointer(ptr unsafe.Pointer) (n *QDate) {
	n = new(QDate)
	n.SetPointer(ptr)
	return
}
func (ptr *QDate) DestroyQDate() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QDate__MonthNameType
//QDate::MonthNameType
type QDate__MonthNameType int64

const (
	QDate__DateFormat       QDate__MonthNameType = QDate__MonthNameType(0)
	QDate__StandaloneFormat QDate__MonthNameType = QDate__MonthNameType(1)
)

func NewQDate2() *QDate {
	tmpValue := NewQDateFromPointer(C.QDate_NewQDate2())
	qt.SetFinalizer(tmpValue, (*QDate).DestroyQDate)
	return tmpValue
}

func NewQDate3(y int, m int, d int) *QDate {
	tmpValue := NewQDateFromPointer(C.QDate_NewQDate3(C.int(int32(y)), C.int(int32(m)), C.int(int32(d))))
	qt.SetFinalizer(tmpValue, (*QDate).DestroyQDate)
	return tmpValue
}

func (ptr *QDate) Day(cal QCalendar_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDate_Day(ptr.Pointer(), PointerFromQCalendar(cal))))
	}
	return 0
}

func (ptr *QDate) Day2() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDate_Day2(ptr.Pointer())))
	}
	return 0
}

func QDate_FromString(stri string, format Qt__DateFormat) *QDate {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQDateFromPointer(C.QDate_QDate_FromString(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QDate).DestroyQDate)
	return tmpValue
}

func (ptr *QDate) FromString(stri string, format Qt__DateFormat) *QDate {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQDateFromPointer(C.QDate_QDate_FromString(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QDate).DestroyQDate)
	return tmpValue
}

func QDate_FromString2(stri string, format string) *QDate {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQDateFromPointer(C.QDate_QDate_FromString2(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.struct_QtCore_PackedString{data: formatC, len: C.longlong(len(format))}))
	qt.SetFinalizer(tmpValue, (*QDate).DestroyQDate)
	return tmpValue
}

func (ptr *QDate) FromString2(stri string, format string) *QDate {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQDateFromPointer(C.QDate_QDate_FromString2(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.struct_QtCore_PackedString{data: formatC, len: C.longlong(len(format))}))
	qt.SetFinalizer(tmpValue, (*QDate).DestroyQDate)
	return tmpValue
}

func QDate_FromString3(stri string, format string, cal QCalendar_ITF) *QDate {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQDateFromPointer(C.QDate_QDate_FromString3(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.struct_QtCore_PackedString{data: formatC, len: C.longlong(len(format))}, PointerFromQCalendar(cal)))
	qt.SetFinalizer(tmpValue, (*QDate).DestroyQDate)
	return tmpValue
}

func (ptr *QDate) FromString3(stri string, format string, cal QCalendar_ITF) *QDate {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQDateFromPointer(C.QDate_QDate_FromString3(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.struct_QtCore_PackedString{data: formatC, len: C.longlong(len(format))}, PointerFromQCalendar(cal)))
	qt.SetFinalizer(tmpValue, (*QDate).DestroyQDate)
	return tmpValue
}

func (ptr *QDate) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDate_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func QDate_IsValid2(year int, month int, day int) bool {
	return int8(C.QDate_QDate_IsValid2(C.int(int32(year)), C.int(int32(month)), C.int(int32(day)))) != 0
}

func (ptr *QDate) IsValid2(year int, month int, day int) bool {
	return int8(C.QDate_QDate_IsValid2(C.int(int32(year)), C.int(int32(month)), C.int(int32(day)))) != 0
}

func (ptr *QDate) Month(cal QCalendar_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDate_Month(ptr.Pointer(), PointerFromQCalendar(cal))))
	}
	return 0
}

func (ptr *QDate) Month2() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDate_Month2(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDate) Year(cal QCalendar_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDate_Year(ptr.Pointer(), PointerFromQCalendar(cal))))
	}
	return 0
}

func (ptr *QDate) Year2() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDate_Year2(ptr.Pointer())))
	}
	return 0
}

type QDateTime struct {
	ptr unsafe.Pointer
}

type QDateTime_ITF interface {
	QDateTime_PTR() *QDateTime
}

func (ptr *QDateTime) QDateTime_PTR() *QDateTime {
	return ptr
}

func (ptr *QDateTime) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDateTime) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDateTime(ptr QDateTime_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateTime_PTR().Pointer()
	}
	return nil
}

func NewQDateTimeFromPointer(ptr unsafe.Pointer) (n *QDateTime) {
	n = new(QDateTime)
	n.SetPointer(ptr)
	return
}

// QDateTime::YearRange
//
//go:generate stringer -type=QDateTime__YearRange
type QDateTime__YearRange int64

const (
	QDateTime__First QDateTime__YearRange = QDateTime__YearRange(-292275056)
)

func NewQDateTime() *QDateTime {
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_NewQDateTime())
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func NewQDateTime2(date QDate_ITF) *QDateTime {
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_NewQDateTime2(PointerFromQDate(date)))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func NewQDateTime3(date QDate_ITF, ti QTime_ITF, spec Qt__TimeSpec) *QDateTime {
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_NewQDateTime3(PointerFromQDate(date), PointerFromQTime(ti), C.longlong(spec)))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func NewQDateTime4(date QDate_ITF, ti QTime_ITF, spec Qt__TimeSpec, offsetSeconds int) *QDateTime {
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_NewQDateTime4(PointerFromQDate(date), PointerFromQTime(ti), C.longlong(spec), C.int(int32(offsetSeconds))))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func NewQDateTime5(date QDate_ITF, ti QTime_ITF, timeZone QTimeZone_ITF) *QDateTime {
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_NewQDateTime5(PointerFromQDate(date), PointerFromQTime(ti), PointerFromQTimeZone(timeZone)))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func NewQDateTime6(other QDateTime_ITF) *QDateTime {
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_NewQDateTime6(PointerFromQDateTime(other)))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func NewQDateTime7(other QDateTime_ITF) *QDateTime {
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_NewQDateTime7(PointerFromQDateTime(other)))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func QDateTime_FromString(stri string, format Qt__DateFormat) *QDateTime {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_QDateTime_FromString(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func (ptr *QDateTime) FromString(stri string, format Qt__DateFormat) *QDateTime {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_QDateTime_FromString(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func QDateTime_FromString2(stri string, format string) *QDateTime {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_QDateTime_FromString2(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.struct_QtCore_PackedString{data: formatC, len: C.longlong(len(format))}))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func (ptr *QDateTime) FromString2(stri string, format string) *QDateTime {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_QDateTime_FromString2(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.struct_QtCore_PackedString{data: formatC, len: C.longlong(len(format))}))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func QDateTime_FromString3(stri string, format string, cal QCalendar_ITF) *QDateTime {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_QDateTime_FromString3(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.struct_QtCore_PackedString{data: formatC, len: C.longlong(len(format))}, PointerFromQCalendar(cal)))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func (ptr *QDateTime) FromString3(stri string, format string, cal QCalendar_ITF) *QDateTime {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQDateTimeFromPointer(C.QDateTime_QDateTime_FromString3(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.struct_QtCore_PackedString{data: formatC, len: C.longlong(len(format))}, PointerFromQCalendar(cal)))
	qt.SetFinalizer(tmpValue, (*QDateTime).DestroyQDateTime)
	return tmpValue
}

func (ptr *QDateTime) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDateTime_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDateTime) Time() *QTime {
	if ptr.Pointer() != nil {
		tmpValue := NewQTimeFromPointer(C.QDateTime_Time(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTime).DestroyQTime)
		return tmpValue
	}
	return nil
}

func (ptr *QDateTime) DestroyQDateTime() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDateTime_DestroyQDateTime(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QDeadlineTimer::ForeverConstant
//
//go:generate stringer -type=QDeadlineTimer__ForeverConstant
type QDeadlineTimer__ForeverConstant int64

const (
	QDeadlineTimer__Forever QDeadlineTimer__ForeverConstant = QDeadlineTimer__ForeverConstant(0)
)

// QDebug::VerbosityLevel
//
//go:generate stringer -type=QDebug__VerbosityLevel
type QDebug__VerbosityLevel int64

const (
	QDebug__MinimumVerbosity QDebug__VerbosityLevel = QDebug__VerbosityLevel(0)
	QDebug__DefaultVerbosity QDebug__VerbosityLevel = QDebug__VerbosityLevel(2)
	QDebug__MaximumVerbosity QDebug__VerbosityLevel = QDebug__VerbosityLevel(7)
)

type QDir struct {
	ptr unsafe.Pointer
}

type QDir_ITF interface {
	QDir_PTR() *QDir
}

func (ptr *QDir) QDir_PTR() *QDir {
	return ptr
}

func (ptr *QDir) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDir) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDir(ptr QDir_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDir_PTR().Pointer()
	}
	return nil
}

func NewQDirFromPointer(ptr unsafe.Pointer) (n *QDir) {
	n = new(QDir)
	n.SetPointer(ptr)
	return
}

// QDir::Filter
//
//go:generate stringer -type=QDir__Filter
type QDir__Filter int64

const (
	QDir__Dirs           QDir__Filter = QDir__Filter(0x001)
	QDir__Files          QDir__Filter = QDir__Filter(0x002)
	QDir__Drives         QDir__Filter = QDir__Filter(0x004)
	QDir__NoSymLinks     QDir__Filter = QDir__Filter(0x008)
	QDir__AllEntries     QDir__Filter = QDir__Filter(QDir__Dirs | QDir__Files | QDir__Drives)
	QDir__TypeMask       QDir__Filter = QDir__Filter(0x00f)
	QDir__Readable       QDir__Filter = QDir__Filter(0x010)
	QDir__Writable       QDir__Filter = QDir__Filter(0x020)
	QDir__Executable     QDir__Filter = QDir__Filter(0x040)
	QDir__PermissionMask QDir__Filter = QDir__Filter(0x070)
	QDir__Modified       QDir__Filter = QDir__Filter(0x080)
	QDir__Hidden         QDir__Filter = QDir__Filter(0x100)
	QDir__System         QDir__Filter = QDir__Filter(0x200)
	QDir__AccessMask     QDir__Filter = QDir__Filter(0x3F0)
	QDir__AllDirs        QDir__Filter = QDir__Filter(0x400)
	QDir__CaseSensitive  QDir__Filter = QDir__Filter(0x800)
	QDir__NoDot          QDir__Filter = QDir__Filter(0x2000)
	QDir__NoDotDot       QDir__Filter = QDir__Filter(0x4000)
	QDir__NoDotAndDotDot QDir__Filter = QDir__Filter(QDir__NoDot | QDir__NoDotDot)
	QDir__NoFilter       QDir__Filter = QDir__Filter(-1)
)

// QDir::SortFlag
//
//go:generate stringer -type=QDir__SortFlag
type QDir__SortFlag int64

const (
	QDir__Name        QDir__SortFlag = QDir__SortFlag(0x00)
	QDir__Time        QDir__SortFlag = QDir__SortFlag(0x01)
	QDir__Size        QDir__SortFlag = QDir__SortFlag(0x02)
	QDir__Unsorted    QDir__SortFlag = QDir__SortFlag(0x03)
	QDir__SortByMask  QDir__SortFlag = QDir__SortFlag(0x03)
	QDir__DirsFirst   QDir__SortFlag = QDir__SortFlag(0x04)
	QDir__Reversed    QDir__SortFlag = QDir__SortFlag(0x08)
	QDir__IgnoreCase  QDir__SortFlag = QDir__SortFlag(0x10)
	QDir__DirsLast    QDir__SortFlag = QDir__SortFlag(0x20)
	QDir__LocaleAware QDir__SortFlag = QDir__SortFlag(0x40)
	QDir__Type        QDir__SortFlag = QDir__SortFlag(0x80)
	QDir__NoSort      QDir__SortFlag = QDir__SortFlag(-1)
)

func NewQDir(dir QDir_ITF) *QDir {
	tmpValue := NewQDirFromPointer(C.QDir_NewQDir(PointerFromQDir(dir)))
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func NewQDir2(path string) *QDir {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	tmpValue := NewQDirFromPointer(C.QDir_NewQDir2(C.struct_QtCore_PackedString{data: pathC, len: C.longlong(len(path))}))
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func NewQDir3(path string, nameFilter string, sort QDir__SortFlag, filters QDir__Filter) *QDir {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	var nameFilterC *C.char
	if nameFilter != "" {
		nameFilterC = C.CString(nameFilter)
		defer C.free(unsafe.Pointer(nameFilterC))
	}
	tmpValue := NewQDirFromPointer(C.QDir_NewQDir3(C.struct_QtCore_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtCore_PackedString{data: nameFilterC, len: C.longlong(len(nameFilter))}, C.longlong(sort), C.longlong(filters)))
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func (ptr *QDir) Cd(dirName string) bool {
	if ptr.Pointer() != nil {
		var dirNameC *C.char
		if dirName != "" {
			dirNameC = C.CString(dirName)
			defer C.free(unsafe.Pointer(dirNameC))
		}
		return int8(C.QDir_Cd(ptr.Pointer(), C.struct_QtCore_PackedString{data: dirNameC, len: C.longlong(len(dirName))})) != 0
	}
	return false
}

func (ptr *QDir) Count() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QDir_Count(ptr.Pointer())))
	}
	return 0
}

func QDir_Current() *QDir {
	tmpValue := NewQDirFromPointer(C.QDir_QDir_Current())
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func (ptr *QDir) Current() *QDir {
	tmpValue := NewQDirFromPointer(C.QDir_QDir_Current())
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func (ptr *QDir) Exists(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QDir_Exists(ptr.Pointer(), C.struct_QtCore_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QDir) Exists2() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDir_Exists2(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDir) Filter() QDir__Filter {
	if ptr.Pointer() != nil {
		return QDir__Filter(C.QDir_Filter(ptr.Pointer()))
	}
	return 0
}

func QDir_Home() *QDir {
	tmpValue := NewQDirFromPointer(C.QDir_QDir_Home())
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func (ptr *QDir) Home() *QDir {
	tmpValue := NewQDirFromPointer(C.QDir_QDir_Home())
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func (ptr *QDir) IsEmpty(filters QDir__Filter) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDir_IsEmpty(ptr.Pointer(), C.longlong(filters))) != 0
	}
	return false
}

func QDir_Match(filter string, fileName string) bool {
	var filterC *C.char
	if filter != "" {
		filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
	}
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	return int8(C.QDir_QDir_Match(C.struct_QtCore_PackedString{data: filterC, len: C.longlong(len(filter))}, C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
}

func (ptr *QDir) Match(filter string, fileName string) bool {
	var filterC *C.char
	if filter != "" {
		filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
	}
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	return int8(C.QDir_QDir_Match(C.struct_QtCore_PackedString{data: filterC, len: C.longlong(len(filter))}, C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
}

func QDir_Match2(filters []string, fileName string) bool {
	filtersC := C.CString(strings.Join(filters, "¡¦!"))
	defer C.free(unsafe.Pointer(filtersC))
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	return int8(C.QDir_QDir_Match2(C.struct_QtCore_PackedString{data: filtersC, len: C.longlong(len(strings.Join(filters, "¡¦!")))}, C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
}

func (ptr *QDir) Match2(filters []string, fileName string) bool {
	filtersC := C.CString(strings.Join(filters, "¡¦!"))
	defer C.free(unsafe.Pointer(filtersC))
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	return int8(C.QDir_QDir_Match2(C.struct_QtCore_PackedString{data: filtersC, len: C.longlong(len(strings.Join(filters, "¡¦!")))}, C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
}

func (ptr *QDir) Mkdir(dirName string) bool {
	if ptr.Pointer() != nil {
		var dirNameC *C.char
		if dirName != "" {
			dirNameC = C.CString(dirName)
			defer C.free(unsafe.Pointer(dirNameC))
		}
		return int8(C.QDir_Mkdir(ptr.Pointer(), C.struct_QtCore_PackedString{data: dirNameC, len: C.longlong(len(dirName))})) != 0
	}
	return false
}

func (ptr *QDir) Path() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDir_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDir) Refresh() {
	if ptr.Pointer() != nil {
		C.QDir_Refresh(ptr.Pointer())
	}
}

func (ptr *QDir) Remove(fileName string) bool {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		return int8(C.QDir_Remove(ptr.Pointer(), C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
	}
	return false
}

func (ptr *QDir) Rename(oldName string, newName string) bool {
	if ptr.Pointer() != nil {
		var oldNameC *C.char
		if oldName != "" {
			oldNameC = C.CString(oldName)
			defer C.free(unsafe.Pointer(oldNameC))
		}
		var newNameC *C.char
		if newName != "" {
			newNameC = C.CString(newName)
			defer C.free(unsafe.Pointer(newNameC))
		}
		return int8(C.QDir_Rename(ptr.Pointer(), C.struct_QtCore_PackedString{data: oldNameC, len: C.longlong(len(oldName))}, C.struct_QtCore_PackedString{data: newNameC, len: C.longlong(len(newName))})) != 0
	}
	return false
}

func QDir_Root() *QDir {
	tmpValue := NewQDirFromPointer(C.QDir_QDir_Root())
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func (ptr *QDir) Root() *QDir {
	tmpValue := NewQDirFromPointer(C.QDir_QDir_Root())
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func QDir_Separator() *QChar {
	tmpValue := NewQCharFromPointer(C.QDir_QDir_Separator())
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func (ptr *QDir) Separator() *QChar {
	tmpValue := NewQCharFromPointer(C.QDir_QDir_Separator())
	qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
	return tmpValue
}

func QDir_SetCurrent(path string) bool {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	return int8(C.QDir_QDir_SetCurrent(C.struct_QtCore_PackedString{data: pathC, len: C.longlong(len(path))})) != 0
}

func (ptr *QDir) SetCurrent(path string) bool {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	return int8(C.QDir_QDir_SetCurrent(C.struct_QtCore_PackedString{data: pathC, len: C.longlong(len(path))})) != 0
}

func QDir_Temp() *QDir {
	tmpValue := NewQDirFromPointer(C.QDir_QDir_Temp())
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func (ptr *QDir) Temp() *QDir {
	tmpValue := NewQDirFromPointer(C.QDir_QDir_Temp())
	qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
	return tmpValue
}

func QDir_ToNativeSeparators(pathName string) string {
	var pathNameC *C.char
	if pathName != "" {
		pathNameC = C.CString(pathName)
		defer C.free(unsafe.Pointer(pathNameC))
	}
	return cGoUnpackString(C.QDir_QDir_ToNativeSeparators(C.struct_QtCore_PackedString{data: pathNameC, len: C.longlong(len(pathName))}))
}

func (ptr *QDir) ToNativeSeparators(pathName string) string {
	var pathNameC *C.char
	if pathName != "" {
		pathNameC = C.CString(pathName)
		defer C.free(unsafe.Pointer(pathNameC))
	}
	return cGoUnpackString(C.QDir_QDir_ToNativeSeparators(C.struct_QtCore_PackedString{data: pathNameC, len: C.longlong(len(pathName))}))
}

func (ptr *QDir) DestroyQDir() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDir_DestroyQDir(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDir) __drives_atList(i int) *QFileInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQFileInfoFromPointer(C.QDir___drives_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QFileInfo).DestroyQFileInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QDir) __drives_setList(i QFileInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QDir___drives_setList(ptr.Pointer(), PointerFromQFileInfo(i))
	}
}

func (ptr *QDir) __drives_newList() unsafe.Pointer {
	return C.QDir___drives_newList(ptr.Pointer())
}

func (ptr *QDir) __entryInfoList_atList(i int) *QFileInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQFileInfoFromPointer(C.QDir___entryInfoList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QFileInfo).DestroyQFileInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QDir) __entryInfoList_setList(i QFileInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QDir___entryInfoList_setList(ptr.Pointer(), PointerFromQFileInfo(i))
	}
}

func (ptr *QDir) __entryInfoList_newList() unsafe.Pointer {
	return C.QDir___entryInfoList_newList(ptr.Pointer())
}

func (ptr *QDir) __entryInfoList_atList2(i int) *QFileInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQFileInfoFromPointer(C.QDir___entryInfoList_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QFileInfo).DestroyQFileInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QDir) __entryInfoList_setList2(i QFileInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QDir___entryInfoList_setList2(ptr.Pointer(), PointerFromQFileInfo(i))
	}
}

func (ptr *QDir) __entryInfoList_newList2() unsafe.Pointer {
	return C.QDir___entryInfoList_newList2(ptr.Pointer())
}

// QDirIterator::IteratorFlag
//
//go:generate stringer -type=QDirIterator__IteratorFlag
type QDirIterator__IteratorFlag int64

const (
	QDirIterator__NoIteratorFlags QDirIterator__IteratorFlag = QDirIterator__IteratorFlag(0x0)
	QDirIterator__FollowSymlinks  QDirIterator__IteratorFlag = QDirIterator__IteratorFlag(0x1)
	QDirIterator__Subdirectories  QDirIterator__IteratorFlag = QDirIterator__IteratorFlag(0x2)
)

type QEasingCurve struct {
	ptr unsafe.Pointer
}

type QEasingCurve_ITF interface {
	QEasingCurve_PTR() *QEasingCurve
}

func (ptr *QEasingCurve) QEasingCurve_PTR() *QEasingCurve {
	return ptr
}

func (ptr *QEasingCurve) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QEasingCurve) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQEasingCurve(ptr QEasingCurve_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEasingCurve_PTR().Pointer()
	}
	return nil
}

func NewQEasingCurveFromPointer(ptr unsafe.Pointer) (n *QEasingCurve) {
	n = new(QEasingCurve)
	n.SetPointer(ptr)
	return
}

// QEasingCurve::Type
//
//go:generate stringer -type=QEasingCurve__Type
type QEasingCurve__Type int64

const (
	QEasingCurve__Linear       QEasingCurve__Type = QEasingCurve__Type(0)
	QEasingCurve__InQuad       QEasingCurve__Type = QEasingCurve__Type(1)
	QEasingCurve__OutQuad      QEasingCurve__Type = QEasingCurve__Type(2)
	QEasingCurve__InOutQuad    QEasingCurve__Type = QEasingCurve__Type(3)
	QEasingCurve__OutInQuad    QEasingCurve__Type = QEasingCurve__Type(4)
	QEasingCurve__InCubic      QEasingCurve__Type = QEasingCurve__Type(5)
	QEasingCurve__OutCubic     QEasingCurve__Type = QEasingCurve__Type(6)
	QEasingCurve__InOutCubic   QEasingCurve__Type = QEasingCurve__Type(7)
	QEasingCurve__OutInCubic   QEasingCurve__Type = QEasingCurve__Type(8)
	QEasingCurve__InQuart      QEasingCurve__Type = QEasingCurve__Type(9)
	QEasingCurve__OutQuart     QEasingCurve__Type = QEasingCurve__Type(10)
	QEasingCurve__InOutQuart   QEasingCurve__Type = QEasingCurve__Type(11)
	QEasingCurve__OutInQuart   QEasingCurve__Type = QEasingCurve__Type(12)
	QEasingCurve__InQuint      QEasingCurve__Type = QEasingCurve__Type(13)
	QEasingCurve__OutQuint     QEasingCurve__Type = QEasingCurve__Type(14)
	QEasingCurve__InOutQuint   QEasingCurve__Type = QEasingCurve__Type(15)
	QEasingCurve__OutInQuint   QEasingCurve__Type = QEasingCurve__Type(16)
	QEasingCurve__InSine       QEasingCurve__Type = QEasingCurve__Type(17)
	QEasingCurve__OutSine      QEasingCurve__Type = QEasingCurve__Type(18)
	QEasingCurve__InOutSine    QEasingCurve__Type = QEasingCurve__Type(19)
	QEasingCurve__OutInSine    QEasingCurve__Type = QEasingCurve__Type(20)
	QEasingCurve__InExpo       QEasingCurve__Type = QEasingCurve__Type(21)
	QEasingCurve__OutExpo      QEasingCurve__Type = QEasingCurve__Type(22)
	QEasingCurve__InOutExpo    QEasingCurve__Type = QEasingCurve__Type(23)
	QEasingCurve__OutInExpo    QEasingCurve__Type = QEasingCurve__Type(24)
	QEasingCurve__InCirc       QEasingCurve__Type = QEasingCurve__Type(25)
	QEasingCurve__OutCirc      QEasingCurve__Type = QEasingCurve__Type(26)
	QEasingCurve__InOutCirc    QEasingCurve__Type = QEasingCurve__Type(27)
	QEasingCurve__OutInCirc    QEasingCurve__Type = QEasingCurve__Type(28)
	QEasingCurve__InElastic    QEasingCurve__Type = QEasingCurve__Type(29)
	QEasingCurve__OutElastic   QEasingCurve__Type = QEasingCurve__Type(30)
	QEasingCurve__InOutElastic QEasingCurve__Type = QEasingCurve__Type(31)
	QEasingCurve__OutInElastic QEasingCurve__Type = QEasingCurve__Type(32)
	QEasingCurve__InBack       QEasingCurve__Type = QEasingCurve__Type(33)
	QEasingCurve__OutBack      QEasingCurve__Type = QEasingCurve__Type(34)
	QEasingCurve__InOutBack    QEasingCurve__Type = QEasingCurve__Type(35)
	QEasingCurve__OutInBack    QEasingCurve__Type = QEasingCurve__Type(36)
	QEasingCurve__InBounce     QEasingCurve__Type = QEasingCurve__Type(37)
	QEasingCurve__OutBounce    QEasingCurve__Type = QEasingCurve__Type(38)
	QEasingCurve__InOutBounce  QEasingCurve__Type = QEasingCurve__Type(39)
	QEasingCurve__OutInBounce  QEasingCurve__Type = QEasingCurve__Type(40)
	QEasingCurve__InCurve      QEasingCurve__Type = QEasingCurve__Type(41)
	QEasingCurve__OutCurve     QEasingCurve__Type = QEasingCurve__Type(42)
	QEasingCurve__SineCurve    QEasingCurve__Type = QEasingCurve__Type(43)
	QEasingCurve__CosineCurve  QEasingCurve__Type = QEasingCurve__Type(44)
	QEasingCurve__BezierSpline QEasingCurve__Type = QEasingCurve__Type(45)
	QEasingCurve__TCBSpline    QEasingCurve__Type = QEasingCurve__Type(46)
	QEasingCurve__Custom       QEasingCurve__Type = QEasingCurve__Type(47)
	QEasingCurve__NCurveTypes  QEasingCurve__Type = QEasingCurve__Type(48)
)

func NewQEasingCurve(ty QEasingCurve__Type) *QEasingCurve {
	tmpValue := NewQEasingCurveFromPointer(C.QEasingCurve_NewQEasingCurve(C.longlong(ty)))
	qt.SetFinalizer(tmpValue, (*QEasingCurve).DestroyQEasingCurve)
	return tmpValue
}

func NewQEasingCurve2(other QEasingCurve_ITF) *QEasingCurve {
	tmpValue := NewQEasingCurveFromPointer(C.QEasingCurve_NewQEasingCurve2(PointerFromQEasingCurve(other)))
	qt.SetFinalizer(tmpValue, (*QEasingCurve).DestroyQEasingCurve)
	return tmpValue
}

func NewQEasingCurve3(other QEasingCurve_ITF) *QEasingCurve {
	tmpValue := NewQEasingCurveFromPointer(C.QEasingCurve_NewQEasingCurve3(PointerFromQEasingCurve(other)))
	qt.SetFinalizer(tmpValue, (*QEasingCurve).DestroyQEasingCurve)
	return tmpValue
}

func (ptr *QEasingCurve) Type() QEasingCurve__Type {
	if ptr.Pointer() != nil {
		return QEasingCurve__Type(C.QEasingCurve_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEasingCurve) DestroyQEasingCurve() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QEasingCurve_DestroyQEasingCurve(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QEasingCurve) __cubicBezierSpline_atList(i int) *QPointF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPointFFromPointer(C.QEasingCurve___cubicBezierSpline_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QEasingCurve) __cubicBezierSpline_setList(i QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QEasingCurve___cubicBezierSpline_setList(ptr.Pointer(), PointerFromQPointF(i))
	}
}

func (ptr *QEasingCurve) __cubicBezierSpline_newList() unsafe.Pointer {
	return C.QEasingCurve___cubicBezierSpline_newList(ptr.Pointer())
}

func (ptr *QEasingCurve) __toCubicSpline_atList(i int) *QPointF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPointFFromPointer(C.QEasingCurve___toCubicSpline_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QEasingCurve) __toCubicSpline_setList(i QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QEasingCurve___toCubicSpline_setList(ptr.Pointer(), PointerFromQPointF(i))
	}
}

func (ptr *QEasingCurve) __toCubicSpline_newList() unsafe.Pointer {
	return C.QEasingCurve___toCubicSpline_newList(ptr.Pointer())
}

// QElapsedTimer::ClockType
//
//go:generate stringer -type=QElapsedTimer__ClockType
type QElapsedTimer__ClockType int64

const (
	QElapsedTimer__SystemTime         QElapsedTimer__ClockType = QElapsedTimer__ClockType(0)
	QElapsedTimer__MonotonicClock     QElapsedTimer__ClockType = QElapsedTimer__ClockType(1)
	QElapsedTimer__TickCounter        QElapsedTimer__ClockType = QElapsedTimer__ClockType(2)
	QElapsedTimer__MachAbsoluteTime   QElapsedTimer__ClockType = QElapsedTimer__ClockType(3)
	QElapsedTimer__PerformanceCounter QElapsedTimer__ClockType = QElapsedTimer__ClockType(4)
)

type QEvent struct {
	ptr unsafe.Pointer
}

type QEvent_ITF interface {
	QEvent_PTR() *QEvent
}

func (ptr *QEvent) QEvent_PTR() *QEvent {
	return ptr
}

func (ptr *QEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQEvent(ptr QEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func NewQEventFromPointer(ptr unsafe.Pointer) (n *QEvent) {
	n = new(QEvent)
	n.SetPointer(ptr)
	return
}

// QEvent::Type
//
//go:generate stringer -type=QEvent__Type
type QEvent__Type int64

const (
	QEvent__None                             QEvent__Type = QEvent__Type(0)
	QEvent__Timer                            QEvent__Type = QEvent__Type(1)
	QEvent__MouseButtonPress                 QEvent__Type = QEvent__Type(2)
	QEvent__MouseButtonRelease               QEvent__Type = QEvent__Type(3)
	QEvent__MouseButtonDblClick              QEvent__Type = QEvent__Type(4)
	QEvent__MouseMove                        QEvent__Type = QEvent__Type(5)
	QEvent__KeyPress                         QEvent__Type = QEvent__Type(6)
	QEvent__KeyRelease                       QEvent__Type = QEvent__Type(7)
	QEvent__FocusIn                          QEvent__Type = QEvent__Type(8)
	QEvent__FocusOut                         QEvent__Type = QEvent__Type(9)
	QEvent__FocusAboutToChange               QEvent__Type = QEvent__Type(23)
	QEvent__Enter                            QEvent__Type = QEvent__Type(10)
	QEvent__Leave                            QEvent__Type = QEvent__Type(11)
	QEvent__Paint                            QEvent__Type = QEvent__Type(12)
	QEvent__Move                             QEvent__Type = QEvent__Type(13)
	QEvent__Resize                           QEvent__Type = QEvent__Type(14)
	QEvent__Create                           QEvent__Type = QEvent__Type(15)
	QEvent__Destroy                          QEvent__Type = QEvent__Type(16)
	QEvent__Show                             QEvent__Type = QEvent__Type(17)
	QEvent__Hide                             QEvent__Type = QEvent__Type(18)
	QEvent__Close                            QEvent__Type = QEvent__Type(19)
	QEvent__Quit                             QEvent__Type = QEvent__Type(20)
	QEvent__ParentChange                     QEvent__Type = QEvent__Type(21)
	QEvent__ParentAboutToChange              QEvent__Type = QEvent__Type(131)
	QEvent__ThreadChange                     QEvent__Type = QEvent__Type(22)
	QEvent__WindowActivate                   QEvent__Type = QEvent__Type(24)
	QEvent__WindowDeactivate                 QEvent__Type = QEvent__Type(25)
	QEvent__ShowToParent                     QEvent__Type = QEvent__Type(26)
	QEvent__HideToParent                     QEvent__Type = QEvent__Type(27)
	QEvent__Wheel                            QEvent__Type = QEvent__Type(31)
	QEvent__WindowTitleChange                QEvent__Type = QEvent__Type(33)
	QEvent__WindowIconChange                 QEvent__Type = QEvent__Type(34)
	QEvent__ApplicationWindowIconChange      QEvent__Type = QEvent__Type(35)
	QEvent__ApplicationFontChange            QEvent__Type = QEvent__Type(36)
	QEvent__ApplicationLayoutDirectionChange QEvent__Type = QEvent__Type(37)
	QEvent__ApplicationPaletteChange         QEvent__Type = QEvent__Type(38)
	QEvent__PaletteChange                    QEvent__Type = QEvent__Type(39)
	QEvent__Clipboard                        QEvent__Type = QEvent__Type(40)
	QEvent__Speech                           QEvent__Type = QEvent__Type(42)
	QEvent__MetaCall                         QEvent__Type = QEvent__Type(43)
	QEvent__SockAct                          QEvent__Type = QEvent__Type(50)
	QEvent__WinEventAct                      QEvent__Type = QEvent__Type(132)
	QEvent__DeferredDelete                   QEvent__Type = QEvent__Type(52)
	QEvent__DragEnter                        QEvent__Type = QEvent__Type(60)
	QEvent__DragMove                         QEvent__Type = QEvent__Type(61)
	QEvent__DragLeave                        QEvent__Type = QEvent__Type(62)
	QEvent__Drop                             QEvent__Type = QEvent__Type(63)
	QEvent__DragResponse                     QEvent__Type = QEvent__Type(64)
	QEvent__ChildAdded                       QEvent__Type = QEvent__Type(68)
	QEvent__ChildPolished                    QEvent__Type = QEvent__Type(69)
	QEvent__ChildRemoved                     QEvent__Type = QEvent__Type(71)
	QEvent__ShowWindowRequest                QEvent__Type = QEvent__Type(73)
	QEvent__PolishRequest                    QEvent__Type = QEvent__Type(74)
	QEvent__Polish                           QEvent__Type = QEvent__Type(75)
	QEvent__LayoutRequest                    QEvent__Type = QEvent__Type(76)
	QEvent__UpdateRequest                    QEvent__Type = QEvent__Type(77)
	QEvent__UpdateLater                      QEvent__Type = QEvent__Type(78)
	QEvent__EmbeddingControl                 QEvent__Type = QEvent__Type(79)
	QEvent__ActivateControl                  QEvent__Type = QEvent__Type(80)
	QEvent__DeactivateControl                QEvent__Type = QEvent__Type(81)
	QEvent__ContextMenu                      QEvent__Type = QEvent__Type(82)
	QEvent__InputMethod                      QEvent__Type = QEvent__Type(83)
	QEvent__TabletMove                       QEvent__Type = QEvent__Type(87)
	QEvent__LocaleChange                     QEvent__Type = QEvent__Type(88)
	QEvent__LanguageChange                   QEvent__Type = QEvent__Type(89)
	QEvent__LayoutDirectionChange            QEvent__Type = QEvent__Type(90)
	QEvent__Style                            QEvent__Type = QEvent__Type(91)
	QEvent__TabletPress                      QEvent__Type = QEvent__Type(92)
	QEvent__TabletRelease                    QEvent__Type = QEvent__Type(93)
	QEvent__OkRequest                        QEvent__Type = QEvent__Type(94)
	QEvent__HelpRequest                      QEvent__Type = QEvent__Type(95)
	QEvent__IconDrag                         QEvent__Type = QEvent__Type(96)
	QEvent__FontChange                       QEvent__Type = QEvent__Type(97)
	QEvent__EnabledChange                    QEvent__Type = QEvent__Type(98)
	QEvent__ActivationChange                 QEvent__Type = QEvent__Type(99)
	QEvent__StyleChange                      QEvent__Type = QEvent__Type(100)
	QEvent__IconTextChange                   QEvent__Type = QEvent__Type(101)
	QEvent__ModifiedChange                   QEvent__Type = QEvent__Type(102)
	QEvent__MouseTrackingChange              QEvent__Type = QEvent__Type(109)
	QEvent__WindowBlocked                    QEvent__Type = QEvent__Type(103)
	QEvent__WindowUnblocked                  QEvent__Type = QEvent__Type(104)
	QEvent__WindowStateChange                QEvent__Type = QEvent__Type(105)
	QEvent__ReadOnlyChange                   QEvent__Type = QEvent__Type(106)
	QEvent__ToolTip                          QEvent__Type = QEvent__Type(110)
	QEvent__WhatsThis                        QEvent__Type = QEvent__Type(111)
	QEvent__StatusTip                        QEvent__Type = QEvent__Type(112)
	QEvent__ActionChanged                    QEvent__Type = QEvent__Type(113)
	QEvent__ActionAdded                      QEvent__Type = QEvent__Type(114)
	QEvent__ActionRemoved                    QEvent__Type = QEvent__Type(115)
	QEvent__FileOpen                         QEvent__Type = QEvent__Type(116)
	QEvent__Shortcut                         QEvent__Type = QEvent__Type(117)
	QEvent__ShortcutOverride                 QEvent__Type = QEvent__Type(51)
	QEvent__WhatsThisClicked                 QEvent__Type = QEvent__Type(118)
	QEvent__ToolBarChange                    QEvent__Type = QEvent__Type(120)
	QEvent__ApplicationActivate              QEvent__Type = QEvent__Type(121)
	QEvent__ApplicationActivated             QEvent__Type = QEvent__Type(QEvent__ApplicationActivate)
	QEvent__ApplicationDeactivate            QEvent__Type = QEvent__Type(122)
	QEvent__ApplicationDeactivated           QEvent__Type = QEvent__Type(QEvent__ApplicationDeactivate)
	QEvent__QueryWhatsThis                   QEvent__Type = QEvent__Type(123)
	QEvent__EnterWhatsThisMode               QEvent__Type = QEvent__Type(124)
	QEvent__LeaveWhatsThisMode               QEvent__Type = QEvent__Type(125)
	QEvent__ZOrderChange                     QEvent__Type = QEvent__Type(126)
	QEvent__HoverEnter                       QEvent__Type = QEvent__Type(127)
	QEvent__HoverLeave                       QEvent__Type = QEvent__Type(128)
	QEvent__HoverMove                        QEvent__Type = QEvent__Type(129)
	QEvent__EnterEditFocus                   QEvent__Type = QEvent__Type(150)
	QEvent__LeaveEditFocus                   QEvent__Type = QEvent__Type(151)
	QEvent__AcceptDropsChange                QEvent__Type = QEvent__Type(152)
	QEvent__ZeroTimerEvent                   QEvent__Type = QEvent__Type(154)
	QEvent__GraphicsSceneMouseMove           QEvent__Type = QEvent__Type(155)
	QEvent__GraphicsSceneMousePress          QEvent__Type = QEvent__Type(156)
	QEvent__GraphicsSceneMouseRelease        QEvent__Type = QEvent__Type(157)
	QEvent__GraphicsSceneMouseDoubleClick    QEvent__Type = QEvent__Type(158)
	QEvent__GraphicsSceneContextMenu         QEvent__Type = QEvent__Type(159)
	QEvent__GraphicsSceneHoverEnter          QEvent__Type = QEvent__Type(160)
	QEvent__GraphicsSceneHoverMove           QEvent__Type = QEvent__Type(161)
	QEvent__GraphicsSceneHoverLeave          QEvent__Type = QEvent__Type(162)
	QEvent__GraphicsSceneHelp                QEvent__Type = QEvent__Type(163)
	QEvent__GraphicsSceneDragEnter           QEvent__Type = QEvent__Type(164)
	QEvent__GraphicsSceneDragMove            QEvent__Type = QEvent__Type(165)
	QEvent__GraphicsSceneDragLeave           QEvent__Type = QEvent__Type(166)
	QEvent__GraphicsSceneDrop                QEvent__Type = QEvent__Type(167)
	QEvent__GraphicsSceneWheel               QEvent__Type = QEvent__Type(168)
	QEvent__KeyboardLayoutChange             QEvent__Type = QEvent__Type(169)
	QEvent__DynamicPropertyChange            QEvent__Type = QEvent__Type(170)
	QEvent__TabletEnterProximity             QEvent__Type = QEvent__Type(171)
	QEvent__TabletLeaveProximity             QEvent__Type = QEvent__Type(172)
	QEvent__NonClientAreaMouseMove           QEvent__Type = QEvent__Type(173)
	QEvent__NonClientAreaMouseButtonPress    QEvent__Type = QEvent__Type(174)
	QEvent__NonClientAreaMouseButtonRelease  QEvent__Type = QEvent__Type(175)
	QEvent__NonClientAreaMouseButtonDblClick QEvent__Type = QEvent__Type(176)
	QEvent__MacSizeChange                    QEvent__Type = QEvent__Type(177)
	QEvent__ContentsRectChange               QEvent__Type = QEvent__Type(178)
	QEvent__MacGLWindowChange                QEvent__Type = QEvent__Type(179)
	QEvent__FutureCallOut                    QEvent__Type = QEvent__Type(180)
	QEvent__GraphicsSceneResize              QEvent__Type = QEvent__Type(181)
	QEvent__GraphicsSceneMove                QEvent__Type = QEvent__Type(182)
	QEvent__CursorChange                     QEvent__Type = QEvent__Type(183)
	QEvent__ToolTipChange                    QEvent__Type = QEvent__Type(184)
	QEvent__NetworkReplyUpdated              QEvent__Type = QEvent__Type(185)
	QEvent__GrabMouse                        QEvent__Type = QEvent__Type(186)
	QEvent__UngrabMouse                      QEvent__Type = QEvent__Type(187)
	QEvent__GrabKeyboard                     QEvent__Type = QEvent__Type(188)
	QEvent__UngrabKeyboard                   QEvent__Type = QEvent__Type(189)
	QEvent__MacGLClearDrawable               QEvent__Type = QEvent__Type(191)
	QEvent__StateMachineSignal               QEvent__Type = QEvent__Type(192)
	QEvent__StateMachineWrapped              QEvent__Type = QEvent__Type(193)
	QEvent__TouchBegin                       QEvent__Type = QEvent__Type(194)
	QEvent__TouchUpdate                      QEvent__Type = QEvent__Type(195)
	QEvent__TouchEnd                         QEvent__Type = QEvent__Type(196)
	QEvent__NativeGesture                    QEvent__Type = QEvent__Type(197)
	QEvent__RequestSoftwareInputPanel        QEvent__Type = QEvent__Type(199)
	QEvent__CloseSoftwareInputPanel          QEvent__Type = QEvent__Type(200)
	QEvent__WinIdChange                      QEvent__Type = QEvent__Type(203)
	QEvent__Gesture                          QEvent__Type = QEvent__Type(198)
	QEvent__GestureOverride                  QEvent__Type = QEvent__Type(202)
	QEvent__ScrollPrepare                    QEvent__Type = QEvent__Type(204)
	QEvent__Scroll                           QEvent__Type = QEvent__Type(205)
	QEvent__Expose                           QEvent__Type = QEvent__Type(206)
	QEvent__InputMethodQuery                 QEvent__Type = QEvent__Type(207)
	QEvent__OrientationChange                QEvent__Type = QEvent__Type(208)
	QEvent__TouchCancel                      QEvent__Type = QEvent__Type(209)
	QEvent__ThemeChange                      QEvent__Type = QEvent__Type(210)
	QEvent__SockClose                        QEvent__Type = QEvent__Type(211)
	QEvent__PlatformPanel                    QEvent__Type = QEvent__Type(212)
	QEvent__StyleAnimationUpdate             QEvent__Type = QEvent__Type(213)
	QEvent__ApplicationStateChange           QEvent__Type = QEvent__Type(214)
	QEvent__WindowChangeInternal             QEvent__Type = QEvent__Type(215)
	QEvent__ScreenChangeInternal             QEvent__Type = QEvent__Type(216)
	QEvent__PlatformSurface                  QEvent__Type = QEvent__Type(217)
	QEvent__Pointer                          QEvent__Type = QEvent__Type(218)
	QEvent__TabletTrackingChange             QEvent__Type = QEvent__Type(219)
	QEvent__User                             QEvent__Type = QEvent__Type(1000)
	QEvent__MaxUser                          QEvent__Type = QEvent__Type(65535)
)

func NewQEvent(ty QEvent__Type) *QEvent {
	tmpValue := NewQEventFromPointer(C.QEvent_NewQEvent(C.longlong(ty)))
	qt.SetFinalizer(tmpValue, (*QEvent).DestroyQEvent)
	return tmpValue
}

func (ptr *QEvent) Accept() {
	if ptr.Pointer() != nil {
		C.QEvent_Accept(ptr.Pointer())
	}
}

func (ptr *QEvent) Type() QEvent__Type {
	if ptr.Pointer() != nil {
		return QEvent__Type(C.QEvent_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQEvent_DestroyQEvent
func callbackQEvent_DestroyQEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQEventFromPointer(ptr).DestroyQEventDefault()
	}
}

func (ptr *QEvent) ConnectDestroyQEvent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QEvent"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QEvent) DisconnectDestroyQEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QEvent")
	}
}

func (ptr *QEvent) DestroyQEvent() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QEvent_DestroyQEvent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QEvent) DestroyQEventDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QEvent_DestroyQEventDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QEventLoop::ProcessEventsFlag
//
//go:generate stringer -type=QEventLoop__ProcessEventsFlag
type QEventLoop__ProcessEventsFlag int64

const (
	QEventLoop__AllEvents              QEventLoop__ProcessEventsFlag = QEventLoop__ProcessEventsFlag(0x00)
	QEventLoop__ExcludeUserInputEvents QEventLoop__ProcessEventsFlag = QEventLoop__ProcessEventsFlag(0x01)
	QEventLoop__ExcludeSocketNotifiers QEventLoop__ProcessEventsFlag = QEventLoop__ProcessEventsFlag(0x02)
	QEventLoop__WaitForMoreEvents      QEventLoop__ProcessEventsFlag = QEventLoop__ProcessEventsFlag(0x04)
	QEventLoop__X11ExcludeTimers       QEventLoop__ProcessEventsFlag = QEventLoop__ProcessEventsFlag(0x08)
	QEventLoop__EventLoopExec          QEventLoop__ProcessEventsFlag = QEventLoop__ProcessEventsFlag(0x20)
	QEventLoop__DialogExec             QEventLoop__ProcessEventsFlag = QEventLoop__ProcessEventsFlag(0x40)
)

type QFile struct {
	QFileDevice
}

type QFile_ITF interface {
	QFileDevice_ITF
	QFile_PTR() *QFile
}

func (ptr *QFile) QFile_PTR() *QFile {
	return ptr
}

func (ptr *QFile) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QFile) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QFileDevice_PTR().SetPointer(p)
	}
}

func PointerFromQFile(ptr QFile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFile_PTR().Pointer()
	}
	return nil
}

func NewQFileFromPointer(ptr unsafe.Pointer) (n *QFile) {
	n = new(QFile)
	n.SetPointer(ptr)
	return
}
func NewQFile() *QFile {
	tmpValue := NewQFileFromPointer(C.QFile_NewQFile())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQFile2(name string) *QFile {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQFileFromPointer(C.QFile_NewQFile2(C.struct_QtCore_PackedString{data: nameC, len: C.longlong(len(name))}))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQFile3(parent QObject_ITF) *QFile {
	tmpValue := NewQFileFromPointer(C.QFile_NewQFile3(PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQFile4(name string, parent QObject_ITF) *QFile {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQFileFromPointer(C.QFile_NewQFile4(C.struct_QtCore_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QFile) Copy(newName string) bool {
	if ptr.Pointer() != nil {
		var newNameC *C.char
		if newName != "" {
			newNameC = C.CString(newName)
			defer C.free(unsafe.Pointer(newNameC))
		}
		return int8(C.QFile_Copy(ptr.Pointer(), C.struct_QtCore_PackedString{data: newNameC, len: C.longlong(len(newName))})) != 0
	}
	return false
}

func QFile_Copy2(fileName string, newName string) bool {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var newNameC *C.char
	if newName != "" {
		newNameC = C.CString(newName)
		defer C.free(unsafe.Pointer(newNameC))
	}
	return int8(C.QFile_QFile_Copy2(C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.struct_QtCore_PackedString{data: newNameC, len: C.longlong(len(newName))})) != 0
}

func (ptr *QFile) Copy2(fileName string, newName string) bool {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var newNameC *C.char
	if newName != "" {
		newNameC = C.CString(newName)
		defer C.free(unsafe.Pointer(newNameC))
	}
	return int8(C.QFile_QFile_Copy2(C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.struct_QtCore_PackedString{data: newNameC, len: C.longlong(len(newName))})) != 0
}

func QFile_Exists(fileName string) bool {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	return int8(C.QFile_QFile_Exists(C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
}

func (ptr *QFile) Exists(fileName string) bool {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	return int8(C.QFile_QFile_Exists(C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
}

func (ptr *QFile) Exists2() bool {
	if ptr.Pointer() != nil {
		return int8(C.QFile_Exists2(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFile) Open3(fd int, mode QIODevice__OpenModeFlag, handleFlags QFileDevice__FileHandleFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QFile_Open3(ptr.Pointer(), C.int(int32(fd)), C.longlong(mode), C.longlong(handleFlags))) != 0
	}
	return false
}

func (ptr *QFile) Remove() bool {
	if ptr.Pointer() != nil {
		return int8(C.QFile_Remove(ptr.Pointer())) != 0
	}
	return false
}

func QFile_Remove2(fileName string) bool {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	return int8(C.QFile_QFile_Remove2(C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
}

func (ptr *QFile) Remove2(fileName string) bool {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	return int8(C.QFile_QFile_Remove2(C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))})) != 0
}

func (ptr *QFile) Rename(newName string) bool {
	if ptr.Pointer() != nil {
		var newNameC *C.char
		if newName != "" {
			newNameC = C.CString(newName)
			defer C.free(unsafe.Pointer(newNameC))
		}
		return int8(C.QFile_Rename(ptr.Pointer(), C.struct_QtCore_PackedString{data: newNameC, len: C.longlong(len(newName))})) != 0
	}
	return false
}

func QFile_Rename2(oldName string, newName string) bool {
	var oldNameC *C.char
	if oldName != "" {
		oldNameC = C.CString(oldName)
		defer C.free(unsafe.Pointer(oldNameC))
	}
	var newNameC *C.char
	if newName != "" {
		newNameC = C.CString(newName)
		defer C.free(unsafe.Pointer(newNameC))
	}
	return int8(C.QFile_QFile_Rename2(C.struct_QtCore_PackedString{data: oldNameC, len: C.longlong(len(oldName))}, C.struct_QtCore_PackedString{data: newNameC, len: C.longlong(len(newName))})) != 0
}

func (ptr *QFile) Rename2(oldName string, newName string) bool {
	var oldNameC *C.char
	if oldName != "" {
		oldNameC = C.CString(oldName)
		defer C.free(unsafe.Pointer(oldNameC))
	}
	var newNameC *C.char
	if newName != "" {
		newNameC = C.CString(newName)
		defer C.free(unsafe.Pointer(newNameC))
	}
	return int8(C.QFile_QFile_Rename2(C.struct_QtCore_PackedString{data: oldNameC, len: C.longlong(len(oldName))}, C.struct_QtCore_PackedString{data: newNameC, len: C.longlong(len(newName))})) != 0
}

func QFile_Resize2(fileName string, sz int64) bool {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	return int8(C.QFile_QFile_Resize2(C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.longlong(sz))) != 0
}

func (ptr *QFile) Resize2(fileName string, sz int64) bool {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	return int8(C.QFile_QFile_Resize2(C.struct_QtCore_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.longlong(sz))) != 0
}

//export callbackQFile_DestroyQFile
func callbackQFile_DestroyQFile(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QFile"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQFileFromPointer(ptr).DestroyQFileDefault()
	}
}

func (ptr *QFile) ConnectDestroyQFile(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QFile"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QFile", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QFile", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QFile) DisconnectDestroyQFile() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QFile")
	}
}

func (ptr *QFile) DestroyQFile() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFile_DestroyQFile(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFile) DestroyQFileDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFile_DestroyQFileDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QFileDevice struct {
	QIODevice
}

type QFileDevice_ITF interface {
	QIODevice_ITF
	QFileDevice_PTR() *QFileDevice
}

func (ptr *QFileDevice) QFileDevice_PTR() *QFileDevice {
	return ptr
}

func (ptr *QFileDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QFileDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQFileDevice(ptr QFileDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileDevice_PTR().Pointer()
	}
	return nil
}

func NewQFileDeviceFromPointer(ptr unsafe.Pointer) (n *QFileDevice) {
	n = new(QFileDevice)
	n.SetPointer(ptr)
	return
}

// QFileDevice::FileError
//
//go:generate stringer -type=QFileDevice__FileError
type QFileDevice__FileError int64

const (
	QFileDevice__NoError          QFileDevice__FileError = QFileDevice__FileError(0)
	QFileDevice__ReadError        QFileDevice__FileError = QFileDevice__FileError(1)
	QFileDevice__WriteError       QFileDevice__FileError = QFileDevice__FileError(2)
	QFileDevice__FatalError       QFileDevice__FileError = QFileDevice__FileError(3)
	QFileDevice__ResourceError    QFileDevice__FileError = QFileDevice__FileError(4)
	QFileDevice__OpenError        QFileDevice__FileError = QFileDevice__FileError(5)
	QFileDevice__AbortError       QFileDevice__FileError = QFileDevice__FileError(6)
	QFileDevice__TimeOutError     QFileDevice__FileError = QFileDevice__FileError(7)
	QFileDevice__UnspecifiedError QFileDevice__FileError = QFileDevice__FileError(8)
	QFileDevice__RemoveError      QFileDevice__FileError = QFileDevice__FileError(9)
	QFileDevice__RenameError      QFileDevice__FileError = QFileDevice__FileError(10)
	QFileDevice__PositionError    QFileDevice__FileError = QFileDevice__FileError(11)
	QFileDevice__ResizeError      QFileDevice__FileError = QFileDevice__FileError(12)
	QFileDevice__PermissionsError QFileDevice__FileError = QFileDevice__FileError(13)
	QFileDevice__CopyError        QFileDevice__FileError = QFileDevice__FileError(14)
)

// QFileDevice::FileHandleFlag
//
//go:generate stringer -type=QFileDevice__FileHandleFlag
type QFileDevice__FileHandleFlag int64

const (
	QFileDevice__AutoCloseHandle QFileDevice__FileHandleFlag = QFileDevice__FileHandleFlag(0x0001)
	QFileDevice__DontCloseHandle QFileDevice__FileHandleFlag = QFileDevice__FileHandleFlag(0)
)

// QFileDevice::FileTime
//
//go:generate stringer -type=QFileDevice__FileTime
type QFileDevice__FileTime int64

const (
	QFileDevice__FileAccessTime         QFileDevice__FileTime = QFileDevice__FileTime(0)
	QFileDevice__FileBirthTime          QFileDevice__FileTime = QFileDevice__FileTime(1)
	QFileDevice__FileMetadataChangeTime QFileDevice__FileTime = QFileDevice__FileTime(2)
	QFileDevice__FileModificationTime   QFileDevice__FileTime = QFileDevice__FileTime(3)
)

// QFileDevice::MemoryMapFlags
//
//go:generate stringer -type=QFileDevice__MemoryMapFlags
type QFileDevice__MemoryMapFlags int64

const (
	QFileDevice__NoOptions        QFileDevice__MemoryMapFlags = QFileDevice__MemoryMapFlags(0)
	QFileDevice__MapPrivateOption QFileDevice__MemoryMapFlags = QFileDevice__MemoryMapFlags(0x0001)
)

// QFileDevice::Permission
//
//go:generate stringer -type=QFileDevice__Permission
type QFileDevice__Permission int64

const (
	QFileDevice__ReadOwner  QFileDevice__Permission = QFileDevice__Permission(0x4000)
	QFileDevice__WriteOwner QFileDevice__Permission = QFileDevice__Permission(0x2000)
	QFileDevice__ExeOwner   QFileDevice__Permission = QFileDevice__Permission(0x1000)
	QFileDevice__ReadUser   QFileDevice__Permission = QFileDevice__Permission(0x0400)
	QFileDevice__WriteUser  QFileDevice__Permission = QFileDevice__Permission(0x0200)
	QFileDevice__ExeUser    QFileDevice__Permission = QFileDevice__Permission(0x0100)
	QFileDevice__ReadGroup  QFileDevice__Permission = QFileDevice__Permission(0x0040)
	QFileDevice__WriteGroup QFileDevice__Permission = QFileDevice__Permission(0x0020)
	QFileDevice__ExeGroup   QFileDevice__Permission = QFileDevice__Permission(0x0010)
	QFileDevice__ReadOther  QFileDevice__Permission = QFileDevice__Permission(0x0004)
	QFileDevice__WriteOther QFileDevice__Permission = QFileDevice__Permission(0x0002)
	QFileDevice__ExeOther   QFileDevice__Permission = QFileDevice__Permission(0x0001)
)

func (ptr *QFileDevice) Error() QFileDevice__FileError {
	if ptr.Pointer() != nil {
		return QFileDevice__FileError(C.QFileDevice_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQFileDevice_FileName
func callbackQFileDevice_FileName(ptr unsafe.Pointer) C.struct_QtCore_PackedString {
	if signal := qt.GetSignal(ptr, "fileName"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_QtCore_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQFileDeviceFromPointer(ptr).FileNameDefault()
	return C.struct_QtCore_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QFileDevice) ConnectFileName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fileName"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "fileName", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fileName", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QFileDevice) DisconnectFileName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fileName")
	}
}

func (ptr *QFileDevice) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QFileDevice_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileDevice) FileNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QFileDevice_FileNameDefault(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileDevice) Map(offset int64, size int64, flags QFileDevice__MemoryMapFlags) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QFileDevice_Map(ptr.Pointer(), C.longlong(offset), C.longlong(size), C.longlong(flags)))
	}
	return ""
}

//export callbackQFileDevice_Resize
func callbackQFileDevice_Resize(ptr unsafe.Pointer, sz C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "resize"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int64) bool)(signal))(int64(sz)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQFileDeviceFromPointer(ptr).ResizeDefault(int64(sz)))))
}

func (ptr *QFileDevice) ConnectResize(f func(sz int64) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resize"); signal != nil {
			f := func(sz int64) bool {
				(*(*func(int64) bool)(signal))(sz)
				return f(sz)
			}
			qt.ConnectSignal(ptr.Pointer(), "resize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QFileDevice) DisconnectResize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resize")
	}
}

func (ptr *QFileDevice) Resize(sz int64) bool {
	if ptr.Pointer() != nil {
		return int8(C.QFileDevice_Resize(ptr.Pointer(), C.longlong(sz))) != 0
	}
	return false
}

func (ptr *QFileDevice) ResizeDefault(sz int64) bool {
	if ptr.Pointer() != nil {
		return int8(C.QFileDevice_ResizeDefault(ptr.Pointer(), C.longlong(sz))) != 0
	}
	return false
}

//export callbackQFileDevice_DestroyQFileDevice
func callbackQFileDevice_DestroyQFileDevice(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QFileDevice"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQFileDeviceFromPointer(ptr).DestroyQFileDeviceDefault()
	}
}

func (ptr *QFileDevice) ConnectDestroyQFileDevice(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QFileDevice"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QFileDevice", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QFileDevice", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QFileDevice) DisconnectDestroyQFileDevice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QFileDevice")
	}
}

func (ptr *QFileDevice) DestroyQFileDevice() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFileDevice_DestroyQFileDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFileDevice) DestroyQFileDeviceDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFileDevice_DestroyQFileDeviceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQFileDevice_ReadData
func callbackQFileDevice_ReadData(ptr unsafe.Pointer, data C.struct_QtCore_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		retS := cGoUnpackString(data)
		ret := C.longlong((*(*func(*string, int64) int64)(signal))(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
		}
		return ret
	}
	retS := cGoUnpackString(data)
	ret := C.longlong(NewQFileDeviceFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
	}
	return ret
}

func (ptr *QFileDevice) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QFileDevice_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QFileDevice) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QFileDevice_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQFileDevice_WriteData
func callbackQFileDevice_WriteData(ptr unsafe.Pointer, data C.struct_QtCore_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxSize)))
	}

	return C.longlong(NewQFileDeviceFromPointer(ptr).WriteDataDefault(cGoUnpackBytes(data), int64(maxSize)))
}

func (ptr *QFileDevice) WriteData(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QFileDevice_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QFileDevice) WriteDataDefault(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QFileDevice_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

type QFileInfo struct {
	ptr unsafe.Pointer
}

type QFileInfo_ITF interface {
	QFileInfo_PTR() *QFileInfo
}

func (ptr *QFileInfo) QFileInfo_PTR() *QFileInfo {
	return ptr
}

func (ptr *QFileInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QFileInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQFileInfo(ptr QFileInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileInfo_PTR().Pointer()
	}
	return nil
}

func NewQFileInfoFromPointer(ptr unsafe.Pointer) (n *QFileInfo) {
	n = new(QFileInfo)
	n.SetPointer(ptr)
	return
}
func NewQFileInfo2() *QFileInfo {
	tmpValue := NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo2())
	qt.SetFinalizer(tmpValue, (*QFileInfo).DestroyQFileInfo)
	return tmpValue
}

func NewQFileInfo3(file string) *QFileInfo {
	var fileC *C.char
	if file != "" {
		fileC = C.CString(file)
		defer C.free(unsafe.Pointer(fileC))
	}
	tmpValue := NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo3(C.struct_QtCore_PackedString{data: fileC, len: C.longlong(len(file))}))
	qt.SetFinalizer(tmpValue, (*QFileInfo).DestroyQFileInfo)
	return tmpValue
}

func NewQFileInfo4(file QFile_ITF) *QFileInfo {
	tmpValue := NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo4(PointerFromQFile(file)))
	qt.SetFinalizer(tmpValue, (*QFileInfo).DestroyQFileInfo)
	return tmpValue
}

func NewQFileInfo5(dir QDir_ITF, file string) *QFileInfo {
	var fileC *C.char
	if file != "" {
		fileC = C.CString(file)
		defer C.free(unsafe.Pointer(fileC))
	}
	tmpValue := NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo5(PointerFromQDir(dir), C.struct_QtCore_PackedString{data: fileC, len: C.longlong(len(file))}))
	qt.SetFinalizer(tmpValue, (*QFileInfo).DestroyQFileInfo)
	return tmpValue
}

func NewQFileInfo6(fileinfo QFileInfo_ITF) *QFileInfo {
	tmpValue := NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo6(PointerFromQFileInfo(fileinfo)))
	qt.SetFinalizer(tmpValue, (*QFileInfo).DestroyQFileInfo)
	return tmpValue
}

func (ptr *QFileInfo) Dir() *QDir {
	if ptr.Pointer() != nil {
		tmpValue := NewQDirFromPointer(C.QFileInfo_Dir(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QDir).DestroyQDir)
		return tmpValue
	}
	return nil
}

func (ptr *QFileInfo) Exists() bool {
	if ptr.Pointer() != nil {
		return int8(C.QFileInfo_Exists(ptr.Pointer())) != 0
	}
	return false
}

func QFileInfo_Exists2(file string) bool {
	var fileC *C.char
	if file != "" {
		fileC = C.CString(file)
		defer C.free(unsafe.Pointer(fileC))
	}
	return int8(C.QFileInfo_QFileInfo_Exists2(C.struct_QtCore_PackedString{data: fileC, len: C.longlong(len(file))})) != 0
}

func (ptr *QFileInfo) Exists2(file string) bool {
	var fileC *C.char
	if file != "" {
		fileC = C.CString(file)
		defer C.free(unsafe.Pointer(fileC))
	}
	return int8(C.QFileInfo_QFileInfo_Exists2(C.struct_QtCore_PackedString{data: fileC, len: C.longlong(len(file))})) != 0
}

func (ptr *QFileInfo) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QFileInfo_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Group() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QFileInfo_Group(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) IsDir() bool {
	if ptr.Pointer() != nil {
		return int8(C.QFileInfo_IsDir(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) Path() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QFileInfo_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Refresh() {
	if ptr.Pointer() != nil {
		C.QFileInfo_Refresh(ptr.Pointer())
	}
}

func (ptr *QFileInfo) Size() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QFileInfo_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileInfo) Suffix() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QFileInfo_Suffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) DestroyQFileInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFileInfo_DestroyQFileInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QHistoryState::HistoryType
//
//go:generate stringer -type=QHistoryState__HistoryType
type QHistoryState__HistoryType int64

const (
	QHistoryState__ShallowHistory QHistoryState__HistoryType = QHistoryState__HistoryType(0)
	QHistoryState__DeepHistory    QHistoryState__HistoryType = QHistoryState__HistoryType(1)
)

type QIODevice struct {
	QObject
}

type QIODevice_ITF interface {
	QObject_ITF
	QIODevice_PTR() *QIODevice
}

func (ptr *QIODevice) QIODevice_PTR() *QIODevice {
	return ptr
}

func (ptr *QIODevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QIODevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQIODevice(ptr QIODevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func NewQIODeviceFromPointer(ptr unsafe.Pointer) (n *QIODevice) {
	n = new(QIODevice)
	n.SetPointer(ptr)
	return
}

// QIODevice::OpenModeFlag
//
//go:generate stringer -type=QIODevice__OpenModeFlag
type QIODevice__OpenModeFlag int64

const (
	QIODevice__NotOpen      QIODevice__OpenModeFlag = QIODevice__OpenModeFlag(0x0000)
	QIODevice__ReadOnly     QIODevice__OpenModeFlag = QIODevice__OpenModeFlag(0x0001)
	QIODevice__WriteOnly    QIODevice__OpenModeFlag = QIODevice__OpenModeFlag(0x0002)
	QIODevice__ReadWrite    QIODevice__OpenModeFlag = QIODevice__OpenModeFlag(QIODevice__ReadOnly | QIODevice__WriteOnly)
	QIODevice__Append       QIODevice__OpenModeFlag = QIODevice__OpenModeFlag(0x0004)
	QIODevice__Truncate     QIODevice__OpenModeFlag = QIODevice__OpenModeFlag(0x0008)
	QIODevice__Text         QIODevice__OpenModeFlag = QIODevice__OpenModeFlag(0x0010)
	QIODevice__Unbuffered   QIODevice__OpenModeFlag = QIODevice__OpenModeFlag(0x0020)
	QIODevice__NewOnly      QIODevice__OpenModeFlag = QIODevice__OpenModeFlag(0x0040)
	QIODevice__ExistingOnly QIODevice__OpenModeFlag = QIODevice__OpenModeFlag(0x0080)
)

func NewQIODevice() *QIODevice {
	tmpValue := NewQIODeviceFromPointer(C.QIODevice_NewQIODevice())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQIODevice2(parent QObject_ITF) *QIODevice {
	tmpValue := NewQIODeviceFromPointer(C.QIODevice_NewQIODevice2(PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQIODevice_Close
func callbackQIODevice_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQIODeviceFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QIODevice) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIODevice) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QIODevice) Close() {
	if ptr.Pointer() != nil {
		C.QIODevice_Close(ptr.Pointer())
	}
}

func (ptr *QIODevice) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QIODevice_CloseDefault(ptr.Pointer())
	}
}

//export callbackQIODevice_Open
func callbackQIODevice_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(QIODevice__OpenModeFlag) bool)(signal))(QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQIODeviceFromPointer(ptr).OpenDefault(QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QIODevice) ConnectOpen(f func(mode QIODevice__OpenModeFlag) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			f := func(mode QIODevice__OpenModeFlag) bool {
				(*(*func(QIODevice__OpenModeFlag) bool)(signal))(mode)
				return f(mode)
			}
			qt.ConnectSignal(ptr.Pointer(), "open", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIODevice) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
	}
}

func (ptr *QIODevice) Open(mode QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIODevice_Open(ptr.Pointer(), C.longlong(mode))) != 0
	}
	return false
}

func (ptr *QIODevice) OpenDefault(mode QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIODevice_OpenDefault(ptr.Pointer(), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQIODevice_Pos
func callbackQIODevice_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "pos"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQIODeviceFromPointer(ptr).PosDefault())
}

func (ptr *QIODevice) ConnectPos(f func() int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "pos"); signal != nil {
			f := func() int64 {
				(*(*func() int64)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "pos", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pos", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIODevice) DisconnectPos() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "pos")
	}
}

func (ptr *QIODevice) Pos() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QIODevice_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QIODevice_PosDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) Read(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QIODevice_Read(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QIODevice) Read2(maxSize int64) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QIODevice_Read2(ptr.Pointer(), C.longlong(maxSize)))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

//export callbackQIODevice_ReadData
func callbackQIODevice_ReadData(ptr unsafe.Pointer, data C.struct_QtCore_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		retS := cGoUnpackString(data)
		ret := C.longlong((*(*func(*string, int64) int64)(signal))(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
		}
		return ret
	}

	return C.longlong(0)
}

func (ptr *QIODevice) ConnectReadData(f func(data *string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "readData"); signal != nil {
			f := func(data *string, maxSize int64) int64 {
				(*(*func(*string, int64) int64)(signal))(data, maxSize)
				return f(data, maxSize)
			}
			qt.ConnectSignal(ptr.Pointer(), "readData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIODevice) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "readData")
	}
}

func (ptr *QIODevice) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QIODevice_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QIODevice) ReadLine(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QIODevice_ReadLine(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QIODevice) ReadLine2(maxSize int64) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QIODevice_ReadLine2(ptr.Pointer(), C.longlong(maxSize)))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

//export callbackQIODevice_Size
func callbackQIODevice_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQIODeviceFromPointer(ptr).SizeDefault())
}

func (ptr *QIODevice) ConnectSize(f func() int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "size"); signal != nil {
			f := func() int64 {
				(*(*func() int64)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "size", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "size", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIODevice) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "size")
	}
}

func (ptr *QIODevice) Size() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QIODevice_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QIODevice_SizeDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIODevice) Write(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QIODevice_Write(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QIODevice) Write2(data string) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return int64(C.QIODevice_Write2(ptr.Pointer(), dataC))
	}
	return 0
}

func (ptr *QIODevice) Write3(byteArray QByteArray_ITF) int64 {
	if ptr.Pointer() != nil {
		return int64(C.QIODevice_Write3(ptr.Pointer(), PointerFromQByteArray(byteArray)))
	}
	return 0
}

//export callbackQIODevice_WriteData
func callbackQIODevice_WriteData(ptr unsafe.Pointer, data C.struct_QtCore_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxSize)))
	}

	return C.longlong(0)
}

func (ptr *QIODevice) ConnectWriteData(f func(data []byte, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "writeData"); signal != nil {
			f := func(data []byte, maxSize int64) int64 {
				(*(*func([]byte, int64) int64)(signal))(data, maxSize)
				return f(data, maxSize)
			}
			qt.ConnectSignal(ptr.Pointer(), "writeData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "writeData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIODevice) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "writeData")
	}
}

func (ptr *QIODevice) WriteData(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QIODevice_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQIODevice_DestroyQIODevice
func callbackQIODevice_DestroyQIODevice(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QIODevice"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQIODeviceFromPointer(ptr).DestroyQIODeviceDefault()
	}
}

func (ptr *QIODevice) ConnectDestroyQIODevice(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QIODevice"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QIODevice", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QIODevice", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIODevice) DisconnectDestroyQIODevice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QIODevice")
	}
}

func (ptr *QIODevice) DestroyQIODevice() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIODevice_DestroyQIODevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIODevice) DestroyQIODeviceDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIODevice_DestroyQIODeviceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QItemSelection struct {
	ptr unsafe.Pointer
}

type QItemSelection_ITF interface {
	QItemSelection_PTR() *QItemSelection
}

func (ptr *QItemSelection) QItemSelection_PTR() *QItemSelection {
	return ptr
}

func (ptr *QItemSelection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QItemSelection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQItemSelection(ptr QItemSelection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemSelection_PTR().Pointer()
	}
	return nil
}

func NewQItemSelectionFromPointer(ptr unsafe.Pointer) (n *QItemSelection) {
	n = new(QItemSelection)
	n.SetPointer(ptr)
	return
}
func (ptr *QItemSelection) DestroyQItemSelection() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQItemSelection() *QItemSelection {
	tmpValue := NewQItemSelectionFromPointer(C.QItemSelection_NewQItemSelection())
	qt.SetFinalizer(tmpValue, (*QItemSelection).DestroyQItemSelection)
	return tmpValue
}

func NewQItemSelection2(topLeft QModelIndex_ITF, bottomRight QModelIndex_ITF) *QItemSelection {
	tmpValue := NewQItemSelectionFromPointer(C.QItemSelection_NewQItemSelection2(PointerFromQModelIndex(topLeft), PointerFromQModelIndex(bottomRight)))
	qt.SetFinalizer(tmpValue, (*QItemSelection).DestroyQItemSelection)
	return tmpValue
}

func (ptr *QItemSelection) Indexes() []*QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []*QModelIndex {
			out := make([]*QModelIndex, int(l.len))
			tmpList := NewQItemSelectionFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__indexes_atList(i)
			}
			return out
		}(C.QItemSelection_Indexes(ptr.Pointer()))
	}
	return make([]*QModelIndex, 0)
}

func (ptr *QItemSelection) Merge(other QItemSelection_ITF, command QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QItemSelection_Merge(ptr.Pointer(), PointerFromQItemSelection(other), C.longlong(command))
	}
}

func (ptr *QItemSelection) Select(topLeft QModelIndex_ITF, bottomRight QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QItemSelection_Select(ptr.Pointer(), PointerFromQModelIndex(topLeft), PointerFromQModelIndex(bottomRight))
	}
}

func QItemSelection_Split(ran QItemSelectionRange_ITF, other QItemSelectionRange_ITF, result QItemSelection_ITF) {
	C.QItemSelection_QItemSelection_Split(PointerFromQItemSelectionRange(ran), PointerFromQItemSelectionRange(other), PointerFromQItemSelection(result))
}

func (ptr *QItemSelection) Split(ran QItemSelectionRange_ITF, other QItemSelectionRange_ITF, result QItemSelection_ITF) {
	C.QItemSelection_QItemSelection_Split(PointerFromQItemSelectionRange(ran), PointerFromQItemSelectionRange(other), PointerFromQItemSelection(result))
}

func (ptr *QItemSelection) __indexes_atList(i int) *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QItemSelection___indexes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QItemSelection) __indexes_setList(i QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QItemSelection___indexes_setList(ptr.Pointer(), PointerFromQModelIndex(i))
	}
}

func (ptr *QItemSelection) __indexes_newList() unsafe.Pointer {
	return C.QItemSelection___indexes_newList(ptr.Pointer())
}

// QItemSelectionModel::SelectionFlag
//
//go:generate stringer -type=QItemSelectionModel__SelectionFlag
type QItemSelectionModel__SelectionFlag int64

const (
	QItemSelectionModel__NoUpdate       QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(0x0000)
	QItemSelectionModel__Clear          QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(0x0001)
	QItemSelectionModel__Select         QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(0x0002)
	QItemSelectionModel__Deselect       QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(0x0004)
	QItemSelectionModel__Toggle         QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(0x0008)
	QItemSelectionModel__Current        QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(0x0010)
	QItemSelectionModel__Rows           QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(0x0020)
	QItemSelectionModel__Columns        QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(0x0040)
	QItemSelectionModel__SelectCurrent  QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(QItemSelectionModel__Select | QItemSelectionModel__Current)
	QItemSelectionModel__ToggleCurrent  QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(QItemSelectionModel__Toggle | QItemSelectionModel__Current)
	QItemSelectionModel__ClearAndSelect QItemSelectionModel__SelectionFlag = QItemSelectionModel__SelectionFlag(QItemSelectionModel__Clear | QItemSelectionModel__Select)
)

type QItemSelectionRange struct {
	ptr unsafe.Pointer
}

type QItemSelectionRange_ITF interface {
	QItemSelectionRange_PTR() *QItemSelectionRange
}

func (ptr *QItemSelectionRange) QItemSelectionRange_PTR() *QItemSelectionRange {
	return ptr
}

func (ptr *QItemSelectionRange) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QItemSelectionRange) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQItemSelectionRange(ptr QItemSelectionRange_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemSelectionRange_PTR().Pointer()
	}
	return nil
}

func NewQItemSelectionRangeFromPointer(ptr unsafe.Pointer) (n *QItemSelectionRange) {
	n = new(QItemSelectionRange)
	n.SetPointer(ptr)
	return
}
func (ptr *QItemSelectionRange) DestroyQItemSelectionRange() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQItemSelectionRange() *QItemSelectionRange {
	tmpValue := NewQItemSelectionRangeFromPointer(C.QItemSelectionRange_NewQItemSelectionRange())
	qt.SetFinalizer(tmpValue, (*QItemSelectionRange).DestroyQItemSelectionRange)
	return tmpValue
}

func NewQItemSelectionRange2(other QItemSelectionRange_ITF) *QItemSelectionRange {
	tmpValue := NewQItemSelectionRangeFromPointer(C.QItemSelectionRange_NewQItemSelectionRange2(PointerFromQItemSelectionRange(other)))
	qt.SetFinalizer(tmpValue, (*QItemSelectionRange).DestroyQItemSelectionRange)
	return tmpValue
}

func NewQItemSelectionRange4(topLeft QModelIndex_ITF, bottomRight QModelIndex_ITF) *QItemSelectionRange {
	tmpValue := NewQItemSelectionRangeFromPointer(C.QItemSelectionRange_NewQItemSelectionRange4(PointerFromQModelIndex(topLeft), PointerFromQModelIndex(bottomRight)))
	qt.SetFinalizer(tmpValue, (*QItemSelectionRange).DestroyQItemSelectionRange)
	return tmpValue
}

func NewQItemSelectionRange5(index QModelIndex_ITF) *QItemSelectionRange {
	tmpValue := NewQItemSelectionRangeFromPointer(C.QItemSelectionRange_NewQItemSelectionRange5(PointerFromQModelIndex(index)))
	qt.SetFinalizer(tmpValue, (*QItemSelectionRange).DestroyQItemSelectionRange)
	return tmpValue
}

func (ptr *QItemSelectionRange) Height() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QItemSelectionRange_Height(ptr.Pointer())))
	}
	return 0
}

func (ptr *QItemSelectionRange) Indexes() []*QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []*QModelIndex {
			out := make([]*QModelIndex, int(l.len))
			tmpList := NewQItemSelectionRangeFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__indexes_atList(i)
			}
			return out
		}(C.QItemSelectionRange_Indexes(ptr.Pointer()))
	}
	return make([]*QModelIndex, 0)
}

func (ptr *QItemSelectionRange) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QItemSelectionRange_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QItemSelectionRange) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QItemSelectionRange_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QItemSelectionRange) Left() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QItemSelectionRange_Left(ptr.Pointer())))
	}
	return 0
}

func (ptr *QItemSelectionRange) Model() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractItemModelFromPointer(C.QItemSelectionRange_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QItemSelectionRange) Parent() *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QItemSelectionRange_Parent(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QItemSelectionRange) Right() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QItemSelectionRange_Right(ptr.Pointer())))
	}
	return 0
}

func (ptr *QItemSelectionRange) Top() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QItemSelectionRange_Top(ptr.Pointer())))
	}
	return 0
}

func (ptr *QItemSelectionRange) Width() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QItemSelectionRange_Width(ptr.Pointer())))
	}
	return 0
}

func (ptr *QItemSelectionRange) __indexes_atList(i int) *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QItemSelectionRange___indexes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QItemSelectionRange) __indexes_setList(i QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QItemSelectionRange___indexes_setList(ptr.Pointer(), PointerFromQModelIndex(i))
	}
}

func (ptr *QItemSelectionRange) __indexes_newList() unsafe.Pointer {
	return C.QItemSelectionRange___indexes_newList(ptr.Pointer())
}

type QJsonArray struct {
	ptr unsafe.Pointer
}

type QJsonArray_ITF interface {
	QJsonArray_PTR() *QJsonArray
}

func (ptr *QJsonArray) QJsonArray_PTR() *QJsonArray {
	return ptr
}

func (ptr *QJsonArray) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QJsonArray) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQJsonArray(ptr QJsonArray_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonArray_PTR().Pointer()
	}
	return nil
}

func NewQJsonArrayFromPointer(ptr unsafe.Pointer) (n *QJsonArray) {
	n = new(QJsonArray)
	n.SetPointer(ptr)
	return
}
func NewQJsonArray() *QJsonArray {
	tmpValue := NewQJsonArrayFromPointer(C.QJsonArray_NewQJsonArray())
	qt.SetFinalizer(tmpValue, (*QJsonArray).DestroyQJsonArray)
	return tmpValue
}

func NewQJsonArray3(other QJsonArray_ITF) *QJsonArray {
	tmpValue := NewQJsonArrayFromPointer(C.QJsonArray_NewQJsonArray3(PointerFromQJsonArray(other)))
	qt.SetFinalizer(tmpValue, (*QJsonArray).DestroyQJsonArray)
	return tmpValue
}

func NewQJsonArray4(other QJsonArray_ITF) *QJsonArray {
	tmpValue := NewQJsonArrayFromPointer(C.QJsonArray_NewQJsonArray4(PointerFromQJsonArray(other)))
	qt.SetFinalizer(tmpValue, (*QJsonArray).DestroyQJsonArray)
	return tmpValue
}

func (ptr *QJsonArray) Append(value QJsonValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Append(ptr.Pointer(), PointerFromQJsonValue(value))
	}
}

func (ptr *QJsonArray) At(i int) *QJsonValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJsonValueFromPointer(C.QJsonArray_At(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonArray) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QJsonArray_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonArray) Empty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJsonArray_Empty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonArray) First() *QJsonValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJsonValueFromPointer(C.QJsonArray_First(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonArray) Insert(i int, value QJsonValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Insert(ptr.Pointer(), C.int(int32(i)), PointerFromQJsonValue(value))
	}
}

func (ptr *QJsonArray) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJsonArray_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonArray) Last() *QJsonValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJsonValueFromPointer(C.QJsonArray_Last(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonArray) Replace(i int, value QJsonValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray_Replace(ptr.Pointer(), C.int(int32(i)), PointerFromQJsonValue(value))
	}
}

func (ptr *QJsonArray) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QJsonArray_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonArray) TakeAt(i int) *QJsonValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJsonValueFromPointer(C.QJsonArray_TakeAt(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonArray) DestroyQJsonArray() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QJsonArray_DestroyQJsonArray(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QJsonArray) __fromVariantList_list_atList(i int) *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QJsonArray___fromVariantList_list_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonArray) __fromVariantList_list_setList(i QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray___fromVariantList_list_setList(ptr.Pointer(), PointerFromQVariant(i))
	}
}

func (ptr *QJsonArray) __fromVariantList_list_newList() unsafe.Pointer {
	return C.QJsonArray___fromVariantList_list_newList(ptr.Pointer())
}

func (ptr *QJsonArray) __toVariantList_atList(i int) *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QJsonArray___toVariantList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonArray) __toVariantList_setList(i QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonArray___toVariantList_setList(ptr.Pointer(), PointerFromQVariant(i))
	}
}

func (ptr *QJsonArray) __toVariantList_newList() unsafe.Pointer {
	return C.QJsonArray___toVariantList_newList(ptr.Pointer())
}

type QJsonDocument struct {
	ptr unsafe.Pointer
}

type QJsonDocument_ITF interface {
	QJsonDocument_PTR() *QJsonDocument
}

func (ptr *QJsonDocument) QJsonDocument_PTR() *QJsonDocument {
	return ptr
}

func (ptr *QJsonDocument) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QJsonDocument) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQJsonDocument(ptr QJsonDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonDocument_PTR().Pointer()
	}
	return nil
}

func NewQJsonDocumentFromPointer(ptr unsafe.Pointer) (n *QJsonDocument) {
	n = new(QJsonDocument)
	n.SetPointer(ptr)
	return
}

// QJsonDocument::DataValidation
//
//go:generate stringer -type=QJsonDocument__DataValidation
type QJsonDocument__DataValidation int64

const (
	QJsonDocument__Validate         QJsonDocument__DataValidation = QJsonDocument__DataValidation(0)
	QJsonDocument__BypassValidation QJsonDocument__DataValidation = QJsonDocument__DataValidation(1)
)

// QJsonDocument::JsonFormat
//
//go:generate stringer -type=QJsonDocument__JsonFormat
type QJsonDocument__JsonFormat int64

const (
	QJsonDocument__Indented QJsonDocument__JsonFormat = QJsonDocument__JsonFormat(0)
	QJsonDocument__Compact  QJsonDocument__JsonFormat = QJsonDocument__JsonFormat(1)
)

func NewQJsonDocument() *QJsonDocument {
	tmpValue := NewQJsonDocumentFromPointer(C.QJsonDocument_NewQJsonDocument())
	qt.SetFinalizer(tmpValue, (*QJsonDocument).DestroyQJsonDocument)
	return tmpValue
}

func NewQJsonDocument2(object QJsonObject_ITF) *QJsonDocument {
	tmpValue := NewQJsonDocumentFromPointer(C.QJsonDocument_NewQJsonDocument2(PointerFromQJsonObject(object)))
	qt.SetFinalizer(tmpValue, (*QJsonDocument).DestroyQJsonDocument)
	return tmpValue
}

func NewQJsonDocument3(array QJsonArray_ITF) *QJsonDocument {
	tmpValue := NewQJsonDocumentFromPointer(C.QJsonDocument_NewQJsonDocument3(PointerFromQJsonArray(array)))
	qt.SetFinalizer(tmpValue, (*QJsonDocument).DestroyQJsonDocument)
	return tmpValue
}

func NewQJsonDocument4(other QJsonDocument_ITF) *QJsonDocument {
	tmpValue := NewQJsonDocumentFromPointer(C.QJsonDocument_NewQJsonDocument4(PointerFromQJsonDocument(other)))
	qt.SetFinalizer(tmpValue, (*QJsonDocument).DestroyQJsonDocument)
	return tmpValue
}

func NewQJsonDocument5(other QJsonDocument_ITF) *QJsonDocument {
	tmpValue := NewQJsonDocumentFromPointer(C.QJsonDocument_NewQJsonDocument5(PointerFromQJsonDocument(other)))
	qt.SetFinalizer(tmpValue, (*QJsonDocument).DestroyQJsonDocument)
	return tmpValue
}

func (ptr *QJsonDocument) Array() *QJsonArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQJsonArrayFromPointer(C.QJsonDocument_Array(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QJsonArray).DestroyQJsonArray)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonDocument) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJsonDocument_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonDocument) Object() *QJsonObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQJsonObjectFromPointer(C.QJsonDocument_Object(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QJsonObject).DestroyQJsonObject)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonDocument) SetObject(object QJsonObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonDocument_SetObject(ptr.Pointer(), PointerFromQJsonObject(object))
	}
}

func (ptr *QJsonDocument) ToVariant() *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QJsonDocument_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonDocument) DestroyQJsonDocument() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QJsonDocument_DestroyQJsonDocument(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QJsonObject struct {
	ptr unsafe.Pointer
}

type QJsonObject_ITF interface {
	QJsonObject_PTR() *QJsonObject
}

func (ptr *QJsonObject) QJsonObject_PTR() *QJsonObject {
	return ptr
}

func (ptr *QJsonObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QJsonObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQJsonObject(ptr QJsonObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonObject_PTR().Pointer()
	}
	return nil
}

func NewQJsonObjectFromPointer(ptr unsafe.Pointer) (n *QJsonObject) {
	n = new(QJsonObject)
	n.SetPointer(ptr)
	return
}
func NewQJsonObject() *QJsonObject {
	tmpValue := NewQJsonObjectFromPointer(C.QJsonObject_NewQJsonObject())
	qt.SetFinalizer(tmpValue, (*QJsonObject).DestroyQJsonObject)
	return tmpValue
}

func NewQJsonObject3(other QJsonObject_ITF) *QJsonObject {
	tmpValue := NewQJsonObjectFromPointer(C.QJsonObject_NewQJsonObject3(PointerFromQJsonObject(other)))
	qt.SetFinalizer(tmpValue, (*QJsonObject).DestroyQJsonObject)
	return tmpValue
}

func NewQJsonObject4(other QJsonObject_ITF) *QJsonObject {
	tmpValue := NewQJsonObjectFromPointer(C.QJsonObject_NewQJsonObject4(PointerFromQJsonObject(other)))
	qt.SetFinalizer(tmpValue, (*QJsonObject).DestroyQJsonObject)
	return tmpValue
}

func (ptr *QJsonObject) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QJsonObject_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonObject) Empty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJsonObject_Empty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonObject) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJsonObject_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonObject) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QJsonObject_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonObject) Remove(key string) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QJsonObject_Remove(ptr.Pointer(), C.struct_QtCore_PackedString{data: keyC, len: C.longlong(len(key))})
	}
}

func (ptr *QJsonObject) Remove2(key QStringView_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonObject_Remove2(ptr.Pointer(), PointerFromQStringView(key))
	}
}

func (ptr *QJsonObject) Remove3(key QLatin1String_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonObject_Remove3(ptr.Pointer(), PointerFromQLatin1String(key))
	}
}

func (ptr *QJsonObject) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QJsonObject_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonObject) Take(key string) *QJsonValue {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := NewQJsonValueFromPointer(C.QJsonObject_Take(ptr.Pointer(), C.struct_QtCore_PackedString{data: keyC, len: C.longlong(len(key))}))
		qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonObject) Take2(key QStringView_ITF) *QJsonValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJsonValueFromPointer(C.QJsonObject_Take2(ptr.Pointer(), PointerFromQStringView(key)))
		qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonObject) Take3(key QLatin1String_ITF) *QJsonValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJsonValueFromPointer(C.QJsonObject_Take3(ptr.Pointer(), PointerFromQLatin1String(key)))
		qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonObject) Value(key string) *QJsonValue {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := NewQJsonValueFromPointer(C.QJsonObject_Value(ptr.Pointer(), C.struct_QtCore_PackedString{data: keyC, len: C.longlong(len(key))}))
		qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonObject) Value2(key QStringView_ITF) *QJsonValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJsonValueFromPointer(C.QJsonObject_Value2(ptr.Pointer(), PointerFromQStringView(key)))
		qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonObject) Value3(key QLatin1String_ITF) *QJsonValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJsonValueFromPointer(C.QJsonObject_Value3(ptr.Pointer(), PointerFromQLatin1String(key)))
		qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonObject) DestroyQJsonObject() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QJsonObject_DestroyQJsonObject(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QJsonObject) __fromVariantHash_hash_atList(v string, i int) *QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := NewQVariantFromPointer(C.QJsonObject___fromVariantHash_hash_atList(ptr.Pointer(), C.struct_QtCore_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonObject) __fromVariantHash_hash_setList(key string, i QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QJsonObject___fromVariantHash_hash_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: keyC, len: C.longlong(len(key))}, PointerFromQVariant(i))
	}
}

func (ptr *QJsonObject) __fromVariantHash_hash_newList() unsafe.Pointer {
	return C.QJsonObject___fromVariantHash_hash_newList(ptr.Pointer())
}

func (ptr *QJsonObject) __fromVariantHash_hash_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQJsonObjectFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____fromVariantHash_hash_keyList_atList(i)
			}
			return out
		}(C.QJsonObject___fromVariantHash_hash_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QJsonObject) __toVariantHash_atList(v string, i int) *QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := NewQVariantFromPointer(C.QJsonObject___toVariantHash_atList(ptr.Pointer(), C.struct_QtCore_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonObject) __toVariantHash_setList(key string, i QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QJsonObject___toVariantHash_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: keyC, len: C.longlong(len(key))}, PointerFromQVariant(i))
	}
}

func (ptr *QJsonObject) __toVariantHash_newList() unsafe.Pointer {
	return C.QJsonObject___toVariantHash_newList(ptr.Pointer())
}

func (ptr *QJsonObject) __toVariantHash_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQJsonObjectFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____toVariantHash_keyList_atList(i)
			}
			return out
		}(C.QJsonObject___toVariantHash_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QJsonObject) __toVariantMap_atList(v string, i int) *QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := NewQVariantFromPointer(C.QJsonObject___toVariantMap_atList(ptr.Pointer(), C.struct_QtCore_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonObject) __toVariantMap_setList(key string, i QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QJsonObject___toVariantMap_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: keyC, len: C.longlong(len(key))}, PointerFromQVariant(i))
	}
}

func (ptr *QJsonObject) __toVariantMap_newList() unsafe.Pointer {
	return C.QJsonObject___toVariantMap_newList(ptr.Pointer())
}

func (ptr *QJsonObject) __toVariantMap_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQJsonObjectFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____toVariantMap_keyList_atList(i)
			}
			return out
		}(C.QJsonObject___toVariantMap_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QJsonObject) ____fromVariantHash_hash_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QJsonObject_____fromVariantHash_hash_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QJsonObject) ____fromVariantHash_hash_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QJsonObject_____fromVariantHash_hash_keyList_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QJsonObject) ____fromVariantHash_hash_keyList_newList() unsafe.Pointer {
	return C.QJsonObject_____fromVariantHash_hash_keyList_newList(ptr.Pointer())
}

func (ptr *QJsonObject) ____fromVariantMap_map_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QJsonObject_____fromVariantMap_map_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QJsonObject) ____fromVariantMap_map_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QJsonObject_____fromVariantMap_map_keyList_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QJsonObject) ____fromVariantMap_map_keyList_newList() unsafe.Pointer {
	return C.QJsonObject_____fromVariantMap_map_keyList_newList(ptr.Pointer())
}

func (ptr *QJsonObject) ____toVariantHash_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QJsonObject_____toVariantHash_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QJsonObject) ____toVariantHash_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QJsonObject_____toVariantHash_keyList_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QJsonObject) ____toVariantHash_keyList_newList() unsafe.Pointer {
	return C.QJsonObject_____toVariantHash_keyList_newList(ptr.Pointer())
}

func (ptr *QJsonObject) ____toVariantMap_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QJsonObject_____toVariantMap_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QJsonObject) ____toVariantMap_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QJsonObject_____toVariantMap_keyList_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QJsonObject) ____toVariantMap_keyList_newList() unsafe.Pointer {
	return C.QJsonObject_____toVariantMap_keyList_newList(ptr.Pointer())
}

// QJsonParseError::ParseError
//
//go:generate stringer -type=QJsonParseError__ParseError
type QJsonParseError__ParseError int64

const (
	QJsonParseError__NoError               QJsonParseError__ParseError = QJsonParseError__ParseError(0)
	QJsonParseError__UnterminatedObject    QJsonParseError__ParseError = QJsonParseError__ParseError(1)
	QJsonParseError__MissingNameSeparator  QJsonParseError__ParseError = QJsonParseError__ParseError(2)
	QJsonParseError__UnterminatedArray     QJsonParseError__ParseError = QJsonParseError__ParseError(3)
	QJsonParseError__MissingValueSeparator QJsonParseError__ParseError = QJsonParseError__ParseError(4)
	QJsonParseError__IllegalValue          QJsonParseError__ParseError = QJsonParseError__ParseError(5)
	QJsonParseError__TerminationByNumber   QJsonParseError__ParseError = QJsonParseError__ParseError(6)
	QJsonParseError__IllegalNumber         QJsonParseError__ParseError = QJsonParseError__ParseError(7)
	QJsonParseError__IllegalEscapeSequence QJsonParseError__ParseError = QJsonParseError__ParseError(8)
	QJsonParseError__IllegalUTF8String     QJsonParseError__ParseError = QJsonParseError__ParseError(9)
	QJsonParseError__UnterminatedString    QJsonParseError__ParseError = QJsonParseError__ParseError(10)
	QJsonParseError__MissingObject         QJsonParseError__ParseError = QJsonParseError__ParseError(11)
	QJsonParseError__DeepNesting           QJsonParseError__ParseError = QJsonParseError__ParseError(12)
	QJsonParseError__DocumentTooLarge      QJsonParseError__ParseError = QJsonParseError__ParseError(13)
	QJsonParseError__GarbageAtEnd          QJsonParseError__ParseError = QJsonParseError__ParseError(14)
)

type QJsonValue struct {
	ptr unsafe.Pointer
}

type QJsonValue_ITF interface {
	QJsonValue_PTR() *QJsonValue
}

func (ptr *QJsonValue) QJsonValue_PTR() *QJsonValue {
	return ptr
}

func (ptr *QJsonValue) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QJsonValue) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQJsonValue(ptr QJsonValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonValue_PTR().Pointer()
	}
	return nil
}

func NewQJsonValueFromPointer(ptr unsafe.Pointer) (n *QJsonValue) {
	n = new(QJsonValue)
	n.SetPointer(ptr)
	return
}

// QJsonValue::Type
//
//go:generate stringer -type=QJsonValue__Type
type QJsonValue__Type int64

const (
	QJsonValue__Null      QJsonValue__Type = QJsonValue__Type(0x0)
	QJsonValue__Bool      QJsonValue__Type = QJsonValue__Type(0x1)
	QJsonValue__Double    QJsonValue__Type = QJsonValue__Type(0x2)
	QJsonValue__String    QJsonValue__Type = QJsonValue__Type(0x3)
	QJsonValue__Array     QJsonValue__Type = QJsonValue__Type(0x4)
	QJsonValue__Object    QJsonValue__Type = QJsonValue__Type(0x5)
	QJsonValue__Undefined QJsonValue__Type = QJsonValue__Type(0x80)
)

func NewQJsonValue(ty QJsonValue__Type) *QJsonValue {
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue(C.longlong(ty)))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue2(b bool) *QJsonValue {
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue2(C.char(int8(qt.GoBoolToInt(b)))))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue3(v float64) *QJsonValue {
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue3(C.double(v)))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue4(v int) *QJsonValue {
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue4(C.int(int32(v))))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue5(v int64) *QJsonValue {
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue5(C.longlong(v)))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue6(s string) *QJsonValue {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue6(C.struct_QtCore_PackedString{data: sC, len: C.longlong(len(s))}))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue7(s QLatin1String_ITF) *QJsonValue {
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue7(PointerFromQLatin1String(s)))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue8(s string) *QJsonValue {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue8(sC))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue9(a QJsonArray_ITF) *QJsonValue {
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue9(PointerFromQJsonArray(a)))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue10(o QJsonObject_ITF) *QJsonValue {
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue10(PointerFromQJsonObject(o)))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue11(other QJsonValue_ITF) *QJsonValue {
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue11(PointerFromQJsonValue(other)))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func NewQJsonValue12(other QJsonValue_ITF) *QJsonValue {
	tmpValue := NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue12(PointerFromQJsonValue(other)))
	qt.SetFinalizer(tmpValue, (*QJsonValue).DestroyQJsonValue)
	return tmpValue
}

func (ptr *QJsonValue) ToInt(defaultValue int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QJsonValue_ToInt(ptr.Pointer(), C.int(int32(defaultValue)))))
	}
	return 0
}

func (ptr *QJsonValue) ToVariant() *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QJsonValue_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJsonValue) Type() QJsonValue__Type {
	if ptr.Pointer() != nil {
		return QJsonValue__Type(C.QJsonValue_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJsonValue) DestroyQJsonValue() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QJsonValue_DestroyQJsonValue(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLatin1Char struct {
	ptr unsafe.Pointer
}

type QLatin1Char_ITF interface {
	QLatin1Char_PTR() *QLatin1Char
}

func (ptr *QLatin1Char) QLatin1Char_PTR() *QLatin1Char {
	return ptr
}

func (ptr *QLatin1Char) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLatin1Char) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLatin1Char(ptr QLatin1Char_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLatin1Char_PTR().Pointer()
	}
	return nil
}

func NewQLatin1CharFromPointer(ptr unsafe.Pointer) (n *QLatin1Char) {
	n = new(QLatin1Char)
	n.SetPointer(ptr)
	return
}
func (ptr *QLatin1Char) DestroyQLatin1Char() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQLatin1Char(c string) *QLatin1Char {
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	tmpValue := NewQLatin1CharFromPointer(C.QLatin1Char_NewQLatin1Char(cC))
	qt.SetFinalizer(tmpValue, (*QLatin1Char).DestroyQLatin1Char)
	return tmpValue
}

type QLatin1String struct {
	ptr unsafe.Pointer
}

type QLatin1String_ITF interface {
	QLatin1String_PTR() *QLatin1String
}

func (ptr *QLatin1String) QLatin1String_PTR() *QLatin1String {
	return ptr
}

func (ptr *QLatin1String) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLatin1String) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLatin1String(ptr QLatin1String_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLatin1String_PTR().Pointer()
	}
	return nil
}

func NewQLatin1StringFromPointer(ptr unsafe.Pointer) (n *QLatin1String) {
	n = new(QLatin1String)
	n.SetPointer(ptr)
	return
}
func (ptr *QLatin1String) DestroyQLatin1String() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQLatin1String() *QLatin1String {
	tmpValue := NewQLatin1StringFromPointer(C.QLatin1String_NewQLatin1String())
	qt.SetFinalizer(tmpValue, (*QLatin1String).DestroyQLatin1String)
	return tmpValue
}

func NewQLatin1String2(str string) *QLatin1String {
	var strC *C.char
	if str != "" {
		strC = C.CString(str)
		defer C.free(unsafe.Pointer(strC))
	}
	tmpValue := NewQLatin1StringFromPointer(C.QLatin1String_NewQLatin1String2(strC))
	qt.SetFinalizer(tmpValue, (*QLatin1String).DestroyQLatin1String)
	return tmpValue
}

func NewQLatin1String3(first string, last string) *QLatin1String {
	var firstC *C.char
	if first != "" {
		firstC = C.CString(first)
		defer C.free(unsafe.Pointer(firstC))
	}
	var lastC *C.char
	if last != "" {
		lastC = C.CString(last)
		defer C.free(unsafe.Pointer(lastC))
	}
	tmpValue := NewQLatin1StringFromPointer(C.QLatin1String_NewQLatin1String3(firstC, lastC))
	qt.SetFinalizer(tmpValue, (*QLatin1String).DestroyQLatin1String)
	return tmpValue
}

func NewQLatin1String4(str string, size int) *QLatin1String {
	var strC *C.char
	if str != "" {
		strC = C.CString(str)
		defer C.free(unsafe.Pointer(strC))
	}
	tmpValue := NewQLatin1StringFromPointer(C.QLatin1String_NewQLatin1String4(strC, C.int(int32(size))))
	qt.SetFinalizer(tmpValue, (*QLatin1String).DestroyQLatin1String)
	return tmpValue
}

func NewQLatin1String5(str QByteArray_ITF) *QLatin1String {
	tmpValue := NewQLatin1StringFromPointer(C.QLatin1String_NewQLatin1String5(PointerFromQByteArray(str)))
	qt.SetFinalizer(tmpValue, (*QLatin1String).DestroyQLatin1String)
	return tmpValue
}

func (ptr *QLatin1String) Data() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLatin1String_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLatin1String) IndexOf(str QStringView_ITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLatin1String_IndexOf(ptr.Pointer(), PointerFromQStringView(str), C.int(int32(from)), C.longlong(cs))))
	}
	return 0
}

func (ptr *QLatin1String) IndexOf2(l1 QLatin1String_ITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLatin1String_IndexOf2(ptr.Pointer(), PointerFromQLatin1String(l1), C.int(int32(from)), C.longlong(cs))))
	}
	return 0
}

func (ptr *QLatin1String) IndexOf3(c QChar_ITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLatin1String_IndexOf3(ptr.Pointer(), PointerFromQChar(c), C.int(int32(from)), C.longlong(cs))))
	}
	return 0
}

func (ptr *QLatin1String) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLatin1String_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLatin1String) Left(length int) *QLatin1String {
	if ptr.Pointer() != nil {
		tmpValue := NewQLatin1StringFromPointer(C.QLatin1String_Left(ptr.Pointer(), C.int(int32(length))))
		qt.SetFinalizer(tmpValue, (*QLatin1String).DestroyQLatin1String)
		return tmpValue
	}
	return nil
}

func (ptr *QLatin1String) Right(length int) *QLatin1String {
	if ptr.Pointer() != nil {
		tmpValue := NewQLatin1StringFromPointer(C.QLatin1String_Right(ptr.Pointer(), C.int(int32(length))))
		qt.SetFinalizer(tmpValue, (*QLatin1String).DestroyQLatin1String)
		return tmpValue
	}
	return nil
}

func (ptr *QLatin1String) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLatin1String_Size(ptr.Pointer())))
	}
	return 0
}

// QLibrary::LoadHint
//
//go:generate stringer -type=QLibrary__LoadHint
type QLibrary__LoadHint int64

const (
	QLibrary__ResolveAllSymbolsHint     QLibrary__LoadHint = QLibrary__LoadHint(0x01)
	QLibrary__ExportExternalSymbolsHint QLibrary__LoadHint = QLibrary__LoadHint(0x02)
	QLibrary__LoadArchiveMemberHint     QLibrary__LoadHint = QLibrary__LoadHint(0x04)
	QLibrary__PreventUnloadHint         QLibrary__LoadHint = QLibrary__LoadHint(0x08)
	QLibrary__DeepBindHint              QLibrary__LoadHint = QLibrary__LoadHint(0x10)
)

// QLibraryInfo::LibraryLocation
//
//go:generate stringer -type=QLibraryInfo__LibraryLocation
type QLibraryInfo__LibraryLocation int64

const (
	QLibraryInfo__PrefixPath             QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(0)
	QLibraryInfo__DocumentationPath      QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(1)
	QLibraryInfo__HeadersPath            QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(2)
	QLibraryInfo__LibrariesPath          QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(3)
	QLibraryInfo__LibraryExecutablesPath QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(4)
	QLibraryInfo__BinariesPath           QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(5)
	QLibraryInfo__PluginsPath            QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(6)
	QLibraryInfo__ImportsPath            QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(7)
	QLibraryInfo__Qml2ImportsPath        QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(8)
	QLibraryInfo__ArchDataPath           QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(9)
	QLibraryInfo__DataPath               QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(10)
	QLibraryInfo__TranslationsPath       QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(11)
	QLibraryInfo__ExamplesPath           QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(12)
	QLibraryInfo__TestsPath              QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(13)
	QLibraryInfo__SettingsPath           QLibraryInfo__LibraryLocation = QLibraryInfo__LibraryLocation(100)
)

type QLine struct {
	ptr unsafe.Pointer
}

type QLine_ITF interface {
	QLine_PTR() *QLine
}

func (ptr *QLine) QLine_PTR() *QLine {
	return ptr
}

func (ptr *QLine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLine(ptr QLine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLine_PTR().Pointer()
	}
	return nil
}

func NewQLineFromPointer(ptr unsafe.Pointer) (n *QLine) {
	n = new(QLine)
	n.SetPointer(ptr)
	return
}
func (ptr *QLine) DestroyQLine() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQLine() *QLine {
	tmpValue := NewQLineFromPointer(C.QLine_NewQLine())
	qt.SetFinalizer(tmpValue, (*QLine).DestroyQLine)
	return tmpValue
}

func NewQLine2(p1 QPoint_ITF, p2 QPoint_ITF) *QLine {
	tmpValue := NewQLineFromPointer(C.QLine_NewQLine2(PointerFromQPoint(p1), PointerFromQPoint(p2)))
	qt.SetFinalizer(tmpValue, (*QLine).DestroyQLine)
	return tmpValue
}

func NewQLine3(x1 int, y1 int, x2 int, y2 int) *QLine {
	tmpValue := NewQLineFromPointer(C.QLine_NewQLine3(C.int(int32(x1)), C.int(int32(y1)), C.int(int32(x2)), C.int(int32(y2))))
	qt.SetFinalizer(tmpValue, (*QLine).DestroyQLine)
	return tmpValue
}

func (ptr *QLine) Center() *QPoint {
	if ptr.Pointer() != nil {
		tmpValue := NewQPointFromPointer(C.QLine_Center(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

type QLineF struct {
	ptr unsafe.Pointer
}

type QLineF_ITF interface {
	QLineF_PTR() *QLineF
}

func (ptr *QLineF) QLineF_PTR() *QLineF {
	return ptr
}

func (ptr *QLineF) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLineF) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLineF(ptr QLineF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLineF_PTR().Pointer()
	}
	return nil
}

func NewQLineFFromPointer(ptr unsafe.Pointer) (n *QLineF) {
	n = new(QLineF)
	n.SetPointer(ptr)
	return
}
func (ptr *QLineF) DestroyQLineF() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQLineF() *QLineF {
	tmpValue := NewQLineFFromPointer(C.QLineF_NewQLineF())
	qt.SetFinalizer(tmpValue, (*QLineF).DestroyQLineF)
	return tmpValue
}

func NewQLineF2(p1 QPointF_ITF, p2 QPointF_ITF) *QLineF {
	tmpValue := NewQLineFFromPointer(C.QLineF_NewQLineF2(PointerFromQPointF(p1), PointerFromQPointF(p2)))
	qt.SetFinalizer(tmpValue, (*QLineF).DestroyQLineF)
	return tmpValue
}

func NewQLineF3(x1 float64, y1 float64, x2 float64, y2 float64) *QLineF {
	tmpValue := NewQLineFFromPointer(C.QLineF_NewQLineF3(C.double(x1), C.double(y1), C.double(x2), C.double(y2)))
	qt.SetFinalizer(tmpValue, (*QLineF).DestroyQLineF)
	return tmpValue
}

func NewQLineF4(line QLine_ITF) *QLineF {
	tmpValue := NewQLineFFromPointer(C.QLineF_NewQLineF4(PointerFromQLine(line)))
	qt.SetFinalizer(tmpValue, (*QLineF).DestroyQLineF)
	return tmpValue
}

func (ptr *QLineF) Center() *QPointF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPointFFromPointer(C.QLineF_Center(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QLineF) Length() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLineF_Length(ptr.Pointer()))
	}
	return 0
}

type QLocale struct {
	ptr unsafe.Pointer
}

type QLocale_ITF interface {
	QLocale_PTR() *QLocale
}

func (ptr *QLocale) QLocale_PTR() *QLocale {
	return ptr
}

func (ptr *QLocale) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLocale) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLocale(ptr QLocale_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocale_PTR().Pointer()
	}
	return nil
}

func NewQLocaleFromPointer(ptr unsafe.Pointer) (n *QLocale) {
	n = new(QLocale)
	n.SetPointer(ptr)
	return
}

// QLocale::Country
//
//go:generate stringer -type=QLocale__Country
type QLocale__Country int64

const (
	QLocale__AnyCountry                             QLocale__Country = QLocale__Country(0)
	QLocale__Afghanistan                            QLocale__Country = QLocale__Country(1)
	QLocale__Albania                                QLocale__Country = QLocale__Country(2)
	QLocale__Algeria                                QLocale__Country = QLocale__Country(3)
	QLocale__AmericanSamoa                          QLocale__Country = QLocale__Country(4)
	QLocale__Andorra                                QLocale__Country = QLocale__Country(5)
	QLocale__Angola                                 QLocale__Country = QLocale__Country(6)
	QLocale__Anguilla                               QLocale__Country = QLocale__Country(7)
	QLocale__Antarctica                             QLocale__Country = QLocale__Country(8)
	QLocale__AntiguaAndBarbuda                      QLocale__Country = QLocale__Country(9)
	QLocale__Argentina                              QLocale__Country = QLocale__Country(10)
	QLocale__Armenia                                QLocale__Country = QLocale__Country(11)
	QLocale__Aruba                                  QLocale__Country = QLocale__Country(12)
	QLocale__Australia                              QLocale__Country = QLocale__Country(13)
	QLocale__Austria                                QLocale__Country = QLocale__Country(14)
	QLocale__Azerbaijan                             QLocale__Country = QLocale__Country(15)
	QLocale__Bahamas                                QLocale__Country = QLocale__Country(16)
	QLocale__Bahrain                                QLocale__Country = QLocale__Country(17)
	QLocale__Bangladesh                             QLocale__Country = QLocale__Country(18)
	QLocale__Barbados                               QLocale__Country = QLocale__Country(19)
	QLocale__Belarus                                QLocale__Country = QLocale__Country(20)
	QLocale__Belgium                                QLocale__Country = QLocale__Country(21)
	QLocale__Belize                                 QLocale__Country = QLocale__Country(22)
	QLocale__Benin                                  QLocale__Country = QLocale__Country(23)
	QLocale__Bermuda                                QLocale__Country = QLocale__Country(24)
	QLocale__Bhutan                                 QLocale__Country = QLocale__Country(25)
	QLocale__Bolivia                                QLocale__Country = QLocale__Country(26)
	QLocale__BosniaAndHerzegowina                   QLocale__Country = QLocale__Country(27)
	QLocale__Botswana                               QLocale__Country = QLocale__Country(28)
	QLocale__BouvetIsland                           QLocale__Country = QLocale__Country(29)
	QLocale__Brazil                                 QLocale__Country = QLocale__Country(30)
	QLocale__BritishIndianOceanTerritory            QLocale__Country = QLocale__Country(31)
	QLocale__Brunei                                 QLocale__Country = QLocale__Country(32)
	QLocale__Bulgaria                               QLocale__Country = QLocale__Country(33)
	QLocale__BurkinaFaso                            QLocale__Country = QLocale__Country(34)
	QLocale__Burundi                                QLocale__Country = QLocale__Country(35)
	QLocale__Cambodia                               QLocale__Country = QLocale__Country(36)
	QLocale__Cameroon                               QLocale__Country = QLocale__Country(37)
	QLocale__Canada                                 QLocale__Country = QLocale__Country(38)
	QLocale__CapeVerde                              QLocale__Country = QLocale__Country(39)
	QLocale__CaymanIslands                          QLocale__Country = QLocale__Country(40)
	QLocale__CentralAfricanRepublic                 QLocale__Country = QLocale__Country(41)
	QLocale__Chad                                   QLocale__Country = QLocale__Country(42)
	QLocale__Chile                                  QLocale__Country = QLocale__Country(43)
	QLocale__China                                  QLocale__Country = QLocale__Country(44)
	QLocale__ChristmasIsland                        QLocale__Country = QLocale__Country(45)
	QLocale__CocosIslands                           QLocale__Country = QLocale__Country(46)
	QLocale__Colombia                               QLocale__Country = QLocale__Country(47)
	QLocale__Comoros                                QLocale__Country = QLocale__Country(48)
	QLocale__CongoKinshasa                          QLocale__Country = QLocale__Country(49)
	QLocale__CongoBrazzaville                       QLocale__Country = QLocale__Country(50)
	QLocale__CookIslands                            QLocale__Country = QLocale__Country(51)
	QLocale__CostaRica                              QLocale__Country = QLocale__Country(52)
	QLocale__IvoryCoast                             QLocale__Country = QLocale__Country(53)
	QLocale__Croatia                                QLocale__Country = QLocale__Country(54)
	QLocale__Cuba                                   QLocale__Country = QLocale__Country(55)
	QLocale__Cyprus                                 QLocale__Country = QLocale__Country(56)
	QLocale__CzechRepublic                          QLocale__Country = QLocale__Country(57)
	QLocale__Denmark                                QLocale__Country = QLocale__Country(58)
	QLocale__Djibouti                               QLocale__Country = QLocale__Country(59)
	QLocale__Dominica                               QLocale__Country = QLocale__Country(60)
	QLocale__DominicanRepublic                      QLocale__Country = QLocale__Country(61)
	QLocale__EastTimor                              QLocale__Country = QLocale__Country(62)
	QLocale__Ecuador                                QLocale__Country = QLocale__Country(63)
	QLocale__Egypt                                  QLocale__Country = QLocale__Country(64)
	QLocale__ElSalvador                             QLocale__Country = QLocale__Country(65)
	QLocale__EquatorialGuinea                       QLocale__Country = QLocale__Country(66)
	QLocale__Eritrea                                QLocale__Country = QLocale__Country(67)
	QLocale__Estonia                                QLocale__Country = QLocale__Country(68)
	QLocale__Ethiopia                               QLocale__Country = QLocale__Country(69)
	QLocale__FalklandIslands                        QLocale__Country = QLocale__Country(70)
	QLocale__FaroeIslands                           QLocale__Country = QLocale__Country(71)
	QLocale__Fiji                                   QLocale__Country = QLocale__Country(72)
	QLocale__Finland                                QLocale__Country = QLocale__Country(73)
	QLocale__France                                 QLocale__Country = QLocale__Country(74)
	QLocale__Guernsey                               QLocale__Country = QLocale__Country(75)
	QLocale__FrenchGuiana                           QLocale__Country = QLocale__Country(76)
	QLocale__FrenchPolynesia                        QLocale__Country = QLocale__Country(77)
	QLocale__FrenchSouthernTerritories              QLocale__Country = QLocale__Country(78)
	QLocale__Gabon                                  QLocale__Country = QLocale__Country(79)
	QLocale__Gambia                                 QLocale__Country = QLocale__Country(80)
	QLocale__Georgia                                QLocale__Country = QLocale__Country(81)
	QLocale__Germany                                QLocale__Country = QLocale__Country(82)
	QLocale__Ghana                                  QLocale__Country = QLocale__Country(83)
	QLocale__Gibraltar                              QLocale__Country = QLocale__Country(84)
	QLocale__Greece                                 QLocale__Country = QLocale__Country(85)
	QLocale__Greenland                              QLocale__Country = QLocale__Country(86)
	QLocale__Grenada                                QLocale__Country = QLocale__Country(87)
	QLocale__Guadeloupe                             QLocale__Country = QLocale__Country(88)
	QLocale__Guam                                   QLocale__Country = QLocale__Country(89)
	QLocale__Guatemala                              QLocale__Country = QLocale__Country(90)
	QLocale__Guinea                                 QLocale__Country = QLocale__Country(91)
	QLocale__GuineaBissau                           QLocale__Country = QLocale__Country(92)
	QLocale__Guyana                                 QLocale__Country = QLocale__Country(93)
	QLocale__Haiti                                  QLocale__Country = QLocale__Country(94)
	QLocale__HeardAndMcDonaldIslands                QLocale__Country = QLocale__Country(95)
	QLocale__Honduras                               QLocale__Country = QLocale__Country(96)
	QLocale__HongKong                               QLocale__Country = QLocale__Country(97)
	QLocale__Hungary                                QLocale__Country = QLocale__Country(98)
	QLocale__Iceland                                QLocale__Country = QLocale__Country(99)
	QLocale__India                                  QLocale__Country = QLocale__Country(100)
	QLocale__Indonesia                              QLocale__Country = QLocale__Country(101)
	QLocale__Iran                                   QLocale__Country = QLocale__Country(102)
	QLocale__Iraq                                   QLocale__Country = QLocale__Country(103)
	QLocale__Ireland                                QLocale__Country = QLocale__Country(104)
	QLocale__Israel                                 QLocale__Country = QLocale__Country(105)
	QLocale__Italy                                  QLocale__Country = QLocale__Country(106)
	QLocale__Jamaica                                QLocale__Country = QLocale__Country(107)
	QLocale__Japan                                  QLocale__Country = QLocale__Country(108)
	QLocale__Jordan                                 QLocale__Country = QLocale__Country(109)
	QLocale__Kazakhstan                             QLocale__Country = QLocale__Country(110)
	QLocale__Kenya                                  QLocale__Country = QLocale__Country(111)
	QLocale__Kiribati                               QLocale__Country = QLocale__Country(112)
	QLocale__NorthKorea                             QLocale__Country = QLocale__Country(113)
	QLocale__SouthKorea                             QLocale__Country = QLocale__Country(114)
	QLocale__Kuwait                                 QLocale__Country = QLocale__Country(115)
	QLocale__Kyrgyzstan                             QLocale__Country = QLocale__Country(116)
	QLocale__Laos                                   QLocale__Country = QLocale__Country(117)
	QLocale__Latvia                                 QLocale__Country = QLocale__Country(118)
	QLocale__Lebanon                                QLocale__Country = QLocale__Country(119)
	QLocale__Lesotho                                QLocale__Country = QLocale__Country(120)
	QLocale__Liberia                                QLocale__Country = QLocale__Country(121)
	QLocale__Libya                                  QLocale__Country = QLocale__Country(122)
	QLocale__Liechtenstein                          QLocale__Country = QLocale__Country(123)
	QLocale__Lithuania                              QLocale__Country = QLocale__Country(124)
	QLocale__Luxembourg                             QLocale__Country = QLocale__Country(125)
	QLocale__Macau                                  QLocale__Country = QLocale__Country(126)
	QLocale__Macedonia                              QLocale__Country = QLocale__Country(127)
	QLocale__Madagascar                             QLocale__Country = QLocale__Country(128)
	QLocale__Malawi                                 QLocale__Country = QLocale__Country(129)
	QLocale__Malaysia                               QLocale__Country = QLocale__Country(130)
	QLocale__Maldives                               QLocale__Country = QLocale__Country(131)
	QLocale__Mali                                   QLocale__Country = QLocale__Country(132)
	QLocale__Malta                                  QLocale__Country = QLocale__Country(133)
	QLocale__MarshallIslands                        QLocale__Country = QLocale__Country(134)
	QLocale__Martinique                             QLocale__Country = QLocale__Country(135)
	QLocale__Mauritania                             QLocale__Country = QLocale__Country(136)
	QLocale__Mauritius                              QLocale__Country = QLocale__Country(137)
	QLocale__Mayotte                                QLocale__Country = QLocale__Country(138)
	QLocale__Mexico                                 QLocale__Country = QLocale__Country(139)
	QLocale__Micronesia                             QLocale__Country = QLocale__Country(140)
	QLocale__Moldova                                QLocale__Country = QLocale__Country(141)
	QLocale__Monaco                                 QLocale__Country = QLocale__Country(142)
	QLocale__Mongolia                               QLocale__Country = QLocale__Country(143)
	QLocale__Montserrat                             QLocale__Country = QLocale__Country(144)
	QLocale__Morocco                                QLocale__Country = QLocale__Country(145)
	QLocale__Mozambique                             QLocale__Country = QLocale__Country(146)
	QLocale__Myanmar                                QLocale__Country = QLocale__Country(147)
	QLocale__Namibia                                QLocale__Country = QLocale__Country(148)
	QLocale__NauruCountry                           QLocale__Country = QLocale__Country(149)
	QLocale__Nepal                                  QLocale__Country = QLocale__Country(150)
	QLocale__Netherlands                            QLocale__Country = QLocale__Country(151)
	QLocale__CuraSao                                QLocale__Country = QLocale__Country(152)
	QLocale__NewCaledonia                           QLocale__Country = QLocale__Country(153)
	QLocale__NewZealand                             QLocale__Country = QLocale__Country(154)
	QLocale__Nicaragua                              QLocale__Country = QLocale__Country(155)
	QLocale__Niger                                  QLocale__Country = QLocale__Country(156)
	QLocale__Nigeria                                QLocale__Country = QLocale__Country(157)
	QLocale__Niue                                   QLocale__Country = QLocale__Country(158)
	QLocale__NorfolkIsland                          QLocale__Country = QLocale__Country(159)
	QLocale__NorthernMarianaIslands                 QLocale__Country = QLocale__Country(160)
	QLocale__Norway                                 QLocale__Country = QLocale__Country(161)
	QLocale__Oman                                   QLocale__Country = QLocale__Country(162)
	QLocale__Pakistan                               QLocale__Country = QLocale__Country(163)
	QLocale__Palau                                  QLocale__Country = QLocale__Country(164)
	QLocale__PalestinianTerritories                 QLocale__Country = QLocale__Country(165)
	QLocale__Panama                                 QLocale__Country = QLocale__Country(166)
	QLocale__PapuaNewGuinea                         QLocale__Country = QLocale__Country(167)
	QLocale__Paraguay                               QLocale__Country = QLocale__Country(168)
	QLocale__Peru                                   QLocale__Country = QLocale__Country(169)
	QLocale__Philippines                            QLocale__Country = QLocale__Country(170)
	QLocale__Pitcairn                               QLocale__Country = QLocale__Country(171)
	QLocale__Poland                                 QLocale__Country = QLocale__Country(172)
	QLocale__Portugal                               QLocale__Country = QLocale__Country(173)
	QLocale__PuertoRico                             QLocale__Country = QLocale__Country(174)
	QLocale__Qatar                                  QLocale__Country = QLocale__Country(175)
	QLocale__Reunion                                QLocale__Country = QLocale__Country(176)
	QLocale__Romania                                QLocale__Country = QLocale__Country(177)
	QLocale__Russia                                 QLocale__Country = QLocale__Country(178)
	QLocale__Rwanda                                 QLocale__Country = QLocale__Country(179)
	QLocale__SaintKittsAndNevis                     QLocale__Country = QLocale__Country(180)
	QLocale__SaintLucia                             QLocale__Country = QLocale__Country(181)
	QLocale__SaintVincentAndTheGrenadines           QLocale__Country = QLocale__Country(182)
	QLocale__Samoa                                  QLocale__Country = QLocale__Country(183)
	QLocale__SanMarino                              QLocale__Country = QLocale__Country(184)
	QLocale__SaoTomeAndPrincipe                     QLocale__Country = QLocale__Country(185)
	QLocale__SaudiArabia                            QLocale__Country = QLocale__Country(186)
	QLocale__Senegal                                QLocale__Country = QLocale__Country(187)
	QLocale__Seychelles                             QLocale__Country = QLocale__Country(188)
	QLocale__SierraLeone                            QLocale__Country = QLocale__Country(189)
	QLocale__Singapore                              QLocale__Country = QLocale__Country(190)
	QLocale__Slovakia                               QLocale__Country = QLocale__Country(191)
	QLocale__Slovenia                               QLocale__Country = QLocale__Country(192)
	QLocale__SolomonIslands                         QLocale__Country = QLocale__Country(193)
	QLocale__Somalia                                QLocale__Country = QLocale__Country(194)
	QLocale__SouthAfrica                            QLocale__Country = QLocale__Country(195)
	QLocale__SouthGeorgiaAndTheSouthSandwichIslands QLocale__Country = QLocale__Country(196)
	QLocale__Spain                                  QLocale__Country = QLocale__Country(197)
	QLocale__SriLanka                               QLocale__Country = QLocale__Country(198)
	QLocale__SaintHelena                            QLocale__Country = QLocale__Country(199)
	QLocale__SaintPierreAndMiquelon                 QLocale__Country = QLocale__Country(200)
	QLocale__Sudan                                  QLocale__Country = QLocale__Country(201)
	QLocale__Suriname                               QLocale__Country = QLocale__Country(202)
	QLocale__SvalbardAndJanMayenIslands             QLocale__Country = QLocale__Country(203)
	QLocale__Swaziland                              QLocale__Country = QLocale__Country(204)
	QLocale__Sweden                                 QLocale__Country = QLocale__Country(205)
	QLocale__Switzerland                            QLocale__Country = QLocale__Country(206)
	QLocale__Syria                                  QLocale__Country = QLocale__Country(207)
	QLocale__Taiwan                                 QLocale__Country = QLocale__Country(208)
	QLocale__Tajikistan                             QLocale__Country = QLocale__Country(209)
	QLocale__Tanzania                               QLocale__Country = QLocale__Country(210)
	QLocale__Thailand                               QLocale__Country = QLocale__Country(211)
	QLocale__Togo                                   QLocale__Country = QLocale__Country(212)
	QLocale__TokelauCountry                         QLocale__Country = QLocale__Country(213)
	QLocale__Tonga                                  QLocale__Country = QLocale__Country(214)
	QLocale__TrinidadAndTobago                      QLocale__Country = QLocale__Country(215)
	QLocale__Tunisia                                QLocale__Country = QLocale__Country(216)
	QLocale__Turkey                                 QLocale__Country = QLocale__Country(217)
	QLocale__Turkmenistan                           QLocale__Country = QLocale__Country(218)
	QLocale__TurksAndCaicosIslands                  QLocale__Country = QLocale__Country(219)
	QLocale__TuvaluCountry                          QLocale__Country = QLocale__Country(220)
	QLocale__Uganda                                 QLocale__Country = QLocale__Country(221)
	QLocale__Ukraine                                QLocale__Country = QLocale__Country(222)
	QLocale__UnitedArabEmirates                     QLocale__Country = QLocale__Country(223)
	QLocale__UnitedKingdom                          QLocale__Country = QLocale__Country(224)
	QLocale__UnitedStates                           QLocale__Country = QLocale__Country(225)
	QLocale__UnitedStatesMinorOutlyingIslands       QLocale__Country = QLocale__Country(226)
	QLocale__Uruguay                                QLocale__Country = QLocale__Country(227)
	QLocale__Uzbekistan                             QLocale__Country = QLocale__Country(228)
	QLocale__Vanuatu                                QLocale__Country = QLocale__Country(229)
	QLocale__VaticanCityState                       QLocale__Country = QLocale__Country(230)
	QLocale__Venezuela                              QLocale__Country = QLocale__Country(231)
	QLocale__Vietnam                                QLocale__Country = QLocale__Country(232)
	QLocale__BritishVirginIslands                   QLocale__Country = QLocale__Country(233)
	QLocale__UnitedStatesVirginIslands              QLocale__Country = QLocale__Country(234)
	QLocale__WallisAndFutunaIslands                 QLocale__Country = QLocale__Country(235)
	QLocale__WesternSahara                          QLocale__Country = QLocale__Country(236)
	QLocale__Yemen                                  QLocale__Country = QLocale__Country(237)
	QLocale__CanaryIslands                          QLocale__Country = QLocale__Country(238)
	QLocale__Zambia                                 QLocale__Country = QLocale__Country(239)
	QLocale__Zimbabwe                               QLocale__Country = QLocale__Country(240)
	QLocale__ClippertonIsland                       QLocale__Country = QLocale__Country(241)
	QLocale__Montenegro                             QLocale__Country = QLocale__Country(242)
	QLocale__Serbia                                 QLocale__Country = QLocale__Country(243)
	QLocale__SaintBarthelemy                        QLocale__Country = QLocale__Country(244)
	QLocale__SaintMartin                            QLocale__Country = QLocale__Country(245)
	QLocale__LatinAmerica                           QLocale__Country = QLocale__Country(246)
	QLocale__AscensionIsland                        QLocale__Country = QLocale__Country(247)
	QLocale__AlandIslands                           QLocale__Country = QLocale__Country(248)
	QLocale__DiegoGarcia                            QLocale__Country = QLocale__Country(249)
	QLocale__CeutaAndMelilla                        QLocale__Country = QLocale__Country(250)
	QLocale__IsleOfMan                              QLocale__Country = QLocale__Country(251)
	QLocale__Jersey                                 QLocale__Country = QLocale__Country(252)
	QLocale__TristanDaCunha                         QLocale__Country = QLocale__Country(253)
	QLocale__SouthSudan                             QLocale__Country = QLocale__Country(254)
	QLocale__Bonaire                                QLocale__Country = QLocale__Country(255)
	QLocale__SintMaarten                            QLocale__Country = QLocale__Country(256)
	QLocale__Kosovo                                 QLocale__Country = QLocale__Country(257)
	QLocale__EuropeanUnion                          QLocale__Country = QLocale__Country(258)
	QLocale__OutlyingOceania                        QLocale__Country = QLocale__Country(259)
	QLocale__World                                  QLocale__Country = QLocale__Country(260)
	QLocale__Europe                                 QLocale__Country = QLocale__Country(261)
	QLocale__DemocraticRepublicOfCongo              QLocale__Country = QLocale__Country(QLocale__CongoKinshasa)
	QLocale__DemocraticRepublicOfKorea              QLocale__Country = QLocale__Country(QLocale__NorthKorea)
	QLocale__LatinAmericaAndTheCaribbean            QLocale__Country = QLocale__Country(QLocale__LatinAmerica)
	QLocale__PeoplesRepublicOfCongo                 QLocale__Country = QLocale__Country(QLocale__CongoBrazzaville)
	QLocale__RepublicOfKorea                        QLocale__Country = QLocale__Country(QLocale__SouthKorea)
	QLocale__RussianFederation                      QLocale__Country = QLocale__Country(QLocale__Russia)
	QLocale__SyrianArabRepublic                     QLocale__Country = QLocale__Country(QLocale__Syria)
	QLocale__Tokelau                                QLocale__Country = QLocale__Country(QLocale__TokelauCountry)
	QLocale__Tuvalu                                 QLocale__Country = QLocale__Country(QLocale__TuvaluCountry)
	QLocale__LastCountry                            QLocale__Country = QLocale__Country(QLocale__Europe)
)

// QLocale::CurrencySymbolFormat
//
//go:generate stringer -type=QLocale__CurrencySymbolFormat
type QLocale__CurrencySymbolFormat int64

const (
	QLocale__CurrencyIsoCode     QLocale__CurrencySymbolFormat = QLocale__CurrencySymbolFormat(0)
	QLocale__CurrencySymbol      QLocale__CurrencySymbolFormat = QLocale__CurrencySymbolFormat(1)
	QLocale__CurrencyDisplayName QLocale__CurrencySymbolFormat = QLocale__CurrencySymbolFormat(2)
)

// QLocale::DataSizeFormat
//
//go:generate stringer -type=QLocale__DataSizeFormat
type QLocale__DataSizeFormat int64

const (
	QLocale__DataSizeBase1000          QLocale__DataSizeFormat = QLocale__DataSizeFormat(1)
	QLocale__DataSizeSIQuantifiers     QLocale__DataSizeFormat = QLocale__DataSizeFormat(2)
	QLocale__DataSizeIecFormat         QLocale__DataSizeFormat = QLocale__DataSizeFormat(0)
	QLocale__DataSizeTraditionalFormat QLocale__DataSizeFormat = QLocale__DataSizeFormat(QLocale__DataSizeSIQuantifiers)
	QLocale__DataSizeSIFormat          QLocale__DataSizeFormat = QLocale__DataSizeFormat(QLocale__DataSizeBase1000 | QLocale__DataSizeSIQuantifiers)
)

// QLocale::FloatingPointPrecisionOption
//
//go:generate stringer -type=QLocale__FloatingPointPrecisionOption
type QLocale__FloatingPointPrecisionOption int64

const (
	QLocale__FloatingPointShortest QLocale__FloatingPointPrecisionOption = QLocale__FloatingPointPrecisionOption(-128)
)

// QLocale::FormatType
//
//go:generate stringer -type=QLocale__FormatType
type QLocale__FormatType int64

const (
	QLocale__LongFormat   QLocale__FormatType = QLocale__FormatType(0)
	QLocale__ShortFormat  QLocale__FormatType = QLocale__FormatType(1)
	QLocale__NarrowFormat QLocale__FormatType = QLocale__FormatType(2)
)

// QLocale::Language
//
//go:generate stringer -type=QLocale__Language
type QLocale__Language int64

const (
	QLocale__AnyLanguage               QLocale__Language = QLocale__Language(0)
	QLocale__C                         QLocale__Language = QLocale__Language(1)
	QLocale__Abkhazian                 QLocale__Language = QLocale__Language(2)
	QLocale__Oromo                     QLocale__Language = QLocale__Language(3)
	QLocale__Afar                      QLocale__Language = QLocale__Language(4)
	QLocale__Afrikaans                 QLocale__Language = QLocale__Language(5)
	QLocale__Albanian                  QLocale__Language = QLocale__Language(6)
	QLocale__Amharic                   QLocale__Language = QLocale__Language(7)
	QLocale__Arabic                    QLocale__Language = QLocale__Language(8)
	QLocale__Armenian                  QLocale__Language = QLocale__Language(9)
	QLocale__Assamese                  QLocale__Language = QLocale__Language(10)
	QLocale__Aymara                    QLocale__Language = QLocale__Language(11)
	QLocale__Azerbaijani               QLocale__Language = QLocale__Language(12)
	QLocale__Bashkir                   QLocale__Language = QLocale__Language(13)
	QLocale__Basque                    QLocale__Language = QLocale__Language(14)
	QLocale__Bengali                   QLocale__Language = QLocale__Language(15)
	QLocale__Dzongkha                  QLocale__Language = QLocale__Language(16)
	QLocale__Bihari                    QLocale__Language = QLocale__Language(17)
	QLocale__Bislama                   QLocale__Language = QLocale__Language(18)
	QLocale__Breton                    QLocale__Language = QLocale__Language(19)
	QLocale__Bulgarian                 QLocale__Language = QLocale__Language(20)
	QLocale__Burmese                   QLocale__Language = QLocale__Language(21)
	QLocale__Belarusian                QLocale__Language = QLocale__Language(22)
	QLocale__Khmer                     QLocale__Language = QLocale__Language(23)
	QLocale__Catalan                   QLocale__Language = QLocale__Language(24)
	QLocale__Chinese                   QLocale__Language = QLocale__Language(25)
	QLocale__Corsican                  QLocale__Language = QLocale__Language(26)
	QLocale__Croatian                  QLocale__Language = QLocale__Language(27)
	QLocale__Czech                     QLocale__Language = QLocale__Language(28)
	QLocale__Danish                    QLocale__Language = QLocale__Language(29)
	QLocale__Dutch                     QLocale__Language = QLocale__Language(30)
	QLocale__English                   QLocale__Language = QLocale__Language(31)
	QLocale__Esperanto                 QLocale__Language = QLocale__Language(32)
	QLocale__Estonian                  QLocale__Language = QLocale__Language(33)
	QLocale__Faroese                   QLocale__Language = QLocale__Language(34)
	QLocale__Fijian                    QLocale__Language = QLocale__Language(35)
	QLocale__Finnish                   QLocale__Language = QLocale__Language(36)
	QLocale__French                    QLocale__Language = QLocale__Language(37)
	QLocale__WesternFrisian            QLocale__Language = QLocale__Language(38)
	QLocale__Gaelic                    QLocale__Language = QLocale__Language(39)
	QLocale__Galician                  QLocale__Language = QLocale__Language(40)
	QLocale__Georgian                  QLocale__Language = QLocale__Language(41)
	QLocale__German                    QLocale__Language = QLocale__Language(42)
	QLocale__Greek                     QLocale__Language = QLocale__Language(43)
	QLocale__Greenlandic               QLocale__Language = QLocale__Language(44)
	QLocale__Guarani                   QLocale__Language = QLocale__Language(45)
	QLocale__Gujarati                  QLocale__Language = QLocale__Language(46)
	QLocale__Hausa                     QLocale__Language = QLocale__Language(47)
	QLocale__Hebrew                    QLocale__Language = QLocale__Language(48)
	QLocale__Hindi                     QLocale__Language = QLocale__Language(49)
	QLocale__Hungarian                 QLocale__Language = QLocale__Language(50)
	QLocale__Icelandic                 QLocale__Language = QLocale__Language(51)
	QLocale__Indonesian                QLocale__Language = QLocale__Language(52)
	QLocale__Interlingua               QLocale__Language = QLocale__Language(53)
	QLocale__Interlingue               QLocale__Language = QLocale__Language(54)
	QLocale__Inuktitut                 QLocale__Language = QLocale__Language(55)
	QLocale__Inupiak                   QLocale__Language = QLocale__Language(56)
	QLocale__Irish                     QLocale__Language = QLocale__Language(57)
	QLocale__Italian                   QLocale__Language = QLocale__Language(58)
	QLocale__Japanese                  QLocale__Language = QLocale__Language(59)
	QLocale__Javanese                  QLocale__Language = QLocale__Language(60)
	QLocale__Kannada                   QLocale__Language = QLocale__Language(61)
	QLocale__Kashmiri                  QLocale__Language = QLocale__Language(62)
	QLocale__Kazakh                    QLocale__Language = QLocale__Language(63)
	QLocale__Kinyarwanda               QLocale__Language = QLocale__Language(64)
	QLocale__Kirghiz                   QLocale__Language = QLocale__Language(65)
	QLocale__Korean                    QLocale__Language = QLocale__Language(66)
	QLocale__Kurdish                   QLocale__Language = QLocale__Language(67)
	QLocale__Rundi                     QLocale__Language = QLocale__Language(68)
	QLocale__Lao                       QLocale__Language = QLocale__Language(69)
	QLocale__Latin                     QLocale__Language = QLocale__Language(70)
	QLocale__Latvian                   QLocale__Language = QLocale__Language(71)
	QLocale__Lingala                   QLocale__Language = QLocale__Language(72)
	QLocale__Lithuanian                QLocale__Language = QLocale__Language(73)
	QLocale__Macedonian                QLocale__Language = QLocale__Language(74)
	QLocale__Malagasy                  QLocale__Language = QLocale__Language(75)
	QLocale__Malay                     QLocale__Language = QLocale__Language(76)
	QLocale__Malayalam                 QLocale__Language = QLocale__Language(77)
	QLocale__Maltese                   QLocale__Language = QLocale__Language(78)
	QLocale__Maori                     QLocale__Language = QLocale__Language(79)
	QLocale__Marathi                   QLocale__Language = QLocale__Language(80)
	QLocale__Marshallese               QLocale__Language = QLocale__Language(81)
	QLocale__Mongolian                 QLocale__Language = QLocale__Language(82)
	QLocale__NauruLanguage             QLocale__Language = QLocale__Language(83)
	QLocale__Nepali                    QLocale__Language = QLocale__Language(84)
	QLocale__NorwegianBokmal           QLocale__Language = QLocale__Language(85)
	QLocale__Occitan                   QLocale__Language = QLocale__Language(86)
	QLocale__Oriya                     QLocale__Language = QLocale__Language(87)
	QLocale__Pashto                    QLocale__Language = QLocale__Language(88)
	QLocale__Persian                   QLocale__Language = QLocale__Language(89)
	QLocale__Polish                    QLocale__Language = QLocale__Language(90)
	QLocale__Portuguese                QLocale__Language = QLocale__Language(91)
	QLocale__Punjabi                   QLocale__Language = QLocale__Language(92)
	QLocale__Quechua                   QLocale__Language = QLocale__Language(93)
	QLocale__Romansh                   QLocale__Language = QLocale__Language(94)
	QLocale__Romanian                  QLocale__Language = QLocale__Language(95)
	QLocale__Russian                   QLocale__Language = QLocale__Language(96)
	QLocale__Samoan                    QLocale__Language = QLocale__Language(97)
	QLocale__Sango                     QLocale__Language = QLocale__Language(98)
	QLocale__Sanskrit                  QLocale__Language = QLocale__Language(99)
	QLocale__Serbian                   QLocale__Language = QLocale__Language(100)
	QLocale__Ossetic                   QLocale__Language = QLocale__Language(101)
	QLocale__SouthernSotho             QLocale__Language = QLocale__Language(102)
	QLocale__Tswana                    QLocale__Language = QLocale__Language(103)
	QLocale__Shona                     QLocale__Language = QLocale__Language(104)
	QLocale__Sindhi                    QLocale__Language = QLocale__Language(105)
	QLocale__Sinhala                   QLocale__Language = QLocale__Language(106)
	QLocale__Swati                     QLocale__Language = QLocale__Language(107)
	QLocale__Slovak                    QLocale__Language = QLocale__Language(108)
	QLocale__Slovenian                 QLocale__Language = QLocale__Language(109)
	QLocale__Somali                    QLocale__Language = QLocale__Language(110)
	QLocale__Spanish                   QLocale__Language = QLocale__Language(111)
	QLocale__Sundanese                 QLocale__Language = QLocale__Language(112)
	QLocale__Swahili                   QLocale__Language = QLocale__Language(113)
	QLocale__Swedish                   QLocale__Language = QLocale__Language(114)
	QLocale__Sardinian                 QLocale__Language = QLocale__Language(115)
	QLocale__Tajik                     QLocale__Language = QLocale__Language(116)
	QLocale__Tamil                     QLocale__Language = QLocale__Language(117)
	QLocale__Tatar                     QLocale__Language = QLocale__Language(118)
	QLocale__Telugu                    QLocale__Language = QLocale__Language(119)
	QLocale__Thai                      QLocale__Language = QLocale__Language(120)
	QLocale__Tibetan                   QLocale__Language = QLocale__Language(121)
	QLocale__Tigrinya                  QLocale__Language = QLocale__Language(122)
	QLocale__Tongan                    QLocale__Language = QLocale__Language(123)
	QLocale__Tsonga                    QLocale__Language = QLocale__Language(124)
	QLocale__Turkish                   QLocale__Language = QLocale__Language(125)
	QLocale__Turkmen                   QLocale__Language = QLocale__Language(126)
	QLocale__Tahitian                  QLocale__Language = QLocale__Language(127)
	QLocale__Uighur                    QLocale__Language = QLocale__Language(128)
	QLocale__Ukrainian                 QLocale__Language = QLocale__Language(129)
	QLocale__Urdu                      QLocale__Language = QLocale__Language(130)
	QLocale__Uzbek                     QLocale__Language = QLocale__Language(131)
	QLocale__Vietnamese                QLocale__Language = QLocale__Language(132)
	QLocale__Volapuk                   QLocale__Language = QLocale__Language(133)
	QLocale__Welsh                     QLocale__Language = QLocale__Language(134)
	QLocale__Wolof                     QLocale__Language = QLocale__Language(135)
	QLocale__Xhosa                     QLocale__Language = QLocale__Language(136)
	QLocale__Yiddish                   QLocale__Language = QLocale__Language(137)
	QLocale__Yoruba                    QLocale__Language = QLocale__Language(138)
	QLocale__Zhuang                    QLocale__Language = QLocale__Language(139)
	QLocale__Zulu                      QLocale__Language = QLocale__Language(140)
	QLocale__NorwegianNynorsk          QLocale__Language = QLocale__Language(141)
	QLocale__Bosnian                   QLocale__Language = QLocale__Language(142)
	QLocale__Divehi                    QLocale__Language = QLocale__Language(143)
	QLocale__Manx                      QLocale__Language = QLocale__Language(144)
	QLocale__Cornish                   QLocale__Language = QLocale__Language(145)
	QLocale__Akan                      QLocale__Language = QLocale__Language(146)
	QLocale__Konkani                   QLocale__Language = QLocale__Language(147)
	QLocale__Ga                        QLocale__Language = QLocale__Language(148)
	QLocale__Igbo                      QLocale__Language = QLocale__Language(149)
	QLocale__Kamba                     QLocale__Language = QLocale__Language(150)
	QLocale__Syriac                    QLocale__Language = QLocale__Language(151)
	QLocale__Blin                      QLocale__Language = QLocale__Language(152)
	QLocale__Geez                      QLocale__Language = QLocale__Language(153)
	QLocale__Koro                      QLocale__Language = QLocale__Language(154)
	QLocale__Sidamo                    QLocale__Language = QLocale__Language(155)
	QLocale__Atsam                     QLocale__Language = QLocale__Language(156)
	QLocale__Tigre                     QLocale__Language = QLocale__Language(157)
	QLocale__Jju                       QLocale__Language = QLocale__Language(158)
	QLocale__Friulian                  QLocale__Language = QLocale__Language(159)
	QLocale__Venda                     QLocale__Language = QLocale__Language(160)
	QLocale__Ewe                       QLocale__Language = QLocale__Language(161)
	QLocale__Walamo                    QLocale__Language = QLocale__Language(162)
	QLocale__Hawaiian                  QLocale__Language = QLocale__Language(163)
	QLocale__Tyap                      QLocale__Language = QLocale__Language(164)
	QLocale__Nyanja                    QLocale__Language = QLocale__Language(165)
	QLocale__Filipino                  QLocale__Language = QLocale__Language(166)
	QLocale__SwissGerman               QLocale__Language = QLocale__Language(167)
	QLocale__SichuanYi                 QLocale__Language = QLocale__Language(168)
	QLocale__Kpelle                    QLocale__Language = QLocale__Language(169)
	QLocale__LowGerman                 QLocale__Language = QLocale__Language(170)
	QLocale__SouthNdebele              QLocale__Language = QLocale__Language(171)
	QLocale__NorthernSotho             QLocale__Language = QLocale__Language(172)
	QLocale__NorthernSami              QLocale__Language = QLocale__Language(173)
	QLocale__Taroko                    QLocale__Language = QLocale__Language(174)
	QLocale__Gusii                     QLocale__Language = QLocale__Language(175)
	QLocale__Taita                     QLocale__Language = QLocale__Language(176)
	QLocale__Fulah                     QLocale__Language = QLocale__Language(177)
	QLocale__Kikuyu                    QLocale__Language = QLocale__Language(178)
	QLocale__Samburu                   QLocale__Language = QLocale__Language(179)
	QLocale__Sena                      QLocale__Language = QLocale__Language(180)
	QLocale__NorthNdebele              QLocale__Language = QLocale__Language(181)
	QLocale__Rombo                     QLocale__Language = QLocale__Language(182)
	QLocale__Tachelhit                 QLocale__Language = QLocale__Language(183)
	QLocale__Kabyle                    QLocale__Language = QLocale__Language(184)
	QLocale__Nyankole                  QLocale__Language = QLocale__Language(185)
	QLocale__Bena                      QLocale__Language = QLocale__Language(186)
	QLocale__Vunjo                     QLocale__Language = QLocale__Language(187)
	QLocale__Bambara                   QLocale__Language = QLocale__Language(188)
	QLocale__Embu                      QLocale__Language = QLocale__Language(189)
	QLocale__Cherokee                  QLocale__Language = QLocale__Language(190)
	QLocale__Morisyen                  QLocale__Language = QLocale__Language(191)
	QLocale__Makonde                   QLocale__Language = QLocale__Language(192)
	QLocale__Langi                     QLocale__Language = QLocale__Language(193)
	QLocale__Ganda                     QLocale__Language = QLocale__Language(194)
	QLocale__Bemba                     QLocale__Language = QLocale__Language(195)
	QLocale__Kabuverdianu              QLocale__Language = QLocale__Language(196)
	QLocale__Meru                      QLocale__Language = QLocale__Language(197)
	QLocale__Kalenjin                  QLocale__Language = QLocale__Language(198)
	QLocale__Nama                      QLocale__Language = QLocale__Language(199)
	QLocale__Machame                   QLocale__Language = QLocale__Language(200)
	QLocale__Colognian                 QLocale__Language = QLocale__Language(201)
	QLocale__Masai                     QLocale__Language = QLocale__Language(202)
	QLocale__Soga                      QLocale__Language = QLocale__Language(203)
	QLocale__Luyia                     QLocale__Language = QLocale__Language(204)
	QLocale__Asu                       QLocale__Language = QLocale__Language(205)
	QLocale__Teso                      QLocale__Language = QLocale__Language(206)
	QLocale__Saho                      QLocale__Language = QLocale__Language(207)
	QLocale__KoyraChiini               QLocale__Language = QLocale__Language(208)
	QLocale__Rwa                       QLocale__Language = QLocale__Language(209)
	QLocale__Luo                       QLocale__Language = QLocale__Language(210)
	QLocale__Chiga                     QLocale__Language = QLocale__Language(211)
	QLocale__CentralMoroccoTamazight   QLocale__Language = QLocale__Language(212)
	QLocale__KoyraboroSenni            QLocale__Language = QLocale__Language(213)
	QLocale__Shambala                  QLocale__Language = QLocale__Language(214)
	QLocale__Bodo                      QLocale__Language = QLocale__Language(215)
	QLocale__Avaric                    QLocale__Language = QLocale__Language(216)
	QLocale__Chamorro                  QLocale__Language = QLocale__Language(217)
	QLocale__Chechen                   QLocale__Language = QLocale__Language(218)
	QLocale__Church                    QLocale__Language = QLocale__Language(219)
	QLocale__Chuvash                   QLocale__Language = QLocale__Language(220)
	QLocale__Cree                      QLocale__Language = QLocale__Language(221)
	QLocale__Haitian                   QLocale__Language = QLocale__Language(222)
	QLocale__Herero                    QLocale__Language = QLocale__Language(223)
	QLocale__HiriMotu                  QLocale__Language = QLocale__Language(224)
	QLocale__Kanuri                    QLocale__Language = QLocale__Language(225)
	QLocale__Komi                      QLocale__Language = QLocale__Language(226)
	QLocale__Kongo                     QLocale__Language = QLocale__Language(227)
	QLocale__Kwanyama                  QLocale__Language = QLocale__Language(228)
	QLocale__Limburgish                QLocale__Language = QLocale__Language(229)
	QLocale__LubaKatanga               QLocale__Language = QLocale__Language(230)
	QLocale__Luxembourgish             QLocale__Language = QLocale__Language(231)
	QLocale__Navaho                    QLocale__Language = QLocale__Language(232)
	QLocale__Ndonga                    QLocale__Language = QLocale__Language(233)
	QLocale__Ojibwa                    QLocale__Language = QLocale__Language(234)
	QLocale__Pali                      QLocale__Language = QLocale__Language(235)
	QLocale__Walloon                   QLocale__Language = QLocale__Language(236)
	QLocale__Aghem                     QLocale__Language = QLocale__Language(237)
	QLocale__Basaa                     QLocale__Language = QLocale__Language(238)
	QLocale__Zarma                     QLocale__Language = QLocale__Language(239)
	QLocale__Duala                     QLocale__Language = QLocale__Language(240)
	QLocale__JolaFonyi                 QLocale__Language = QLocale__Language(241)
	QLocale__Ewondo                    QLocale__Language = QLocale__Language(242)
	QLocale__Bafia                     QLocale__Language = QLocale__Language(243)
	QLocale__MakhuwaMeetto             QLocale__Language = QLocale__Language(244)
	QLocale__Mundang                   QLocale__Language = QLocale__Language(245)
	QLocale__Kwasio                    QLocale__Language = QLocale__Language(246)
	QLocale__Nuer                      QLocale__Language = QLocale__Language(247)
	QLocale__Sakha                     QLocale__Language = QLocale__Language(248)
	QLocale__Sangu                     QLocale__Language = QLocale__Language(249)
	QLocale__CongoSwahili              QLocale__Language = QLocale__Language(250)
	QLocale__Tasawaq                   QLocale__Language = QLocale__Language(251)
	QLocale__Vai                       QLocale__Language = QLocale__Language(252)
	QLocale__Walser                    QLocale__Language = QLocale__Language(253)
	QLocale__Yangben                   QLocale__Language = QLocale__Language(254)
	QLocale__Avestan                   QLocale__Language = QLocale__Language(255)
	QLocale__Asturian                  QLocale__Language = QLocale__Language(256)
	QLocale__Ngomba                    QLocale__Language = QLocale__Language(257)
	QLocale__Kako                      QLocale__Language = QLocale__Language(258)
	QLocale__Meta                      QLocale__Language = QLocale__Language(259)
	QLocale__Ngiemboon                 QLocale__Language = QLocale__Language(260)
	QLocale__Aragonese                 QLocale__Language = QLocale__Language(261)
	QLocale__Akkadian                  QLocale__Language = QLocale__Language(262)
	QLocale__AncientEgyptian           QLocale__Language = QLocale__Language(263)
	QLocale__AncientGreek              QLocale__Language = QLocale__Language(264)
	QLocale__Aramaic                   QLocale__Language = QLocale__Language(265)
	QLocale__Balinese                  QLocale__Language = QLocale__Language(266)
	QLocale__Bamun                     QLocale__Language = QLocale__Language(267)
	QLocale__BatakToba                 QLocale__Language = QLocale__Language(268)
	QLocale__Buginese                  QLocale__Language = QLocale__Language(269)
	QLocale__Buhid                     QLocale__Language = QLocale__Language(270)
	QLocale__Carian                    QLocale__Language = QLocale__Language(271)
	QLocale__Chakma                    QLocale__Language = QLocale__Language(272)
	QLocale__ClassicalMandaic          QLocale__Language = QLocale__Language(273)
	QLocale__Coptic                    QLocale__Language = QLocale__Language(274)
	QLocale__Dogri                     QLocale__Language = QLocale__Language(275)
	QLocale__EasternCham               QLocale__Language = QLocale__Language(276)
	QLocale__EasternKayah              QLocale__Language = QLocale__Language(277)
	QLocale__Etruscan                  QLocale__Language = QLocale__Language(278)
	QLocale__Gothic                    QLocale__Language = QLocale__Language(279)
	QLocale__Hanunoo                   QLocale__Language = QLocale__Language(280)
	QLocale__Ingush                    QLocale__Language = QLocale__Language(281)
	QLocale__LargeFloweryMiao          QLocale__Language = QLocale__Language(282)
	QLocale__Lepcha                    QLocale__Language = QLocale__Language(283)
	QLocale__Limbu                     QLocale__Language = QLocale__Language(284)
	QLocale__Lisu                      QLocale__Language = QLocale__Language(285)
	QLocale__Lu                        QLocale__Language = QLocale__Language(286)
	QLocale__Lycian                    QLocale__Language = QLocale__Language(287)
	QLocale__Lydian                    QLocale__Language = QLocale__Language(288)
	QLocale__Mandingo                  QLocale__Language = QLocale__Language(289)
	QLocale__Manipuri                  QLocale__Language = QLocale__Language(290)
	QLocale__Meroitic                  QLocale__Language = QLocale__Language(291)
	QLocale__NorthernThai              QLocale__Language = QLocale__Language(292)
	QLocale__OldIrish                  QLocale__Language = QLocale__Language(293)
	QLocale__OldNorse                  QLocale__Language = QLocale__Language(294)
	QLocale__OldPersian                QLocale__Language = QLocale__Language(295)
	QLocale__OldTurkish                QLocale__Language = QLocale__Language(296)
	QLocale__Pahlavi                   QLocale__Language = QLocale__Language(297)
	QLocale__Parthian                  QLocale__Language = QLocale__Language(298)
	QLocale__Phoenician                QLocale__Language = QLocale__Language(299)
	QLocale__PrakritLanguage           QLocale__Language = QLocale__Language(300)
	QLocale__Rejang                    QLocale__Language = QLocale__Language(301)
	QLocale__Sabaean                   QLocale__Language = QLocale__Language(302)
	QLocale__Samaritan                 QLocale__Language = QLocale__Language(303)
	QLocale__Santali                   QLocale__Language = QLocale__Language(304)
	QLocale__Saurashtra                QLocale__Language = QLocale__Language(305)
	QLocale__Sora                      QLocale__Language = QLocale__Language(306)
	QLocale__Sylheti                   QLocale__Language = QLocale__Language(307)
	QLocale__Tagbanwa                  QLocale__Language = QLocale__Language(308)
	QLocale__TaiDam                    QLocale__Language = QLocale__Language(309)
	QLocale__TaiNua                    QLocale__Language = QLocale__Language(310)
	QLocale__Ugaritic                  QLocale__Language = QLocale__Language(311)
	QLocale__Akoose                    QLocale__Language = QLocale__Language(312)
	QLocale__Lakota                    QLocale__Language = QLocale__Language(313)
	QLocale__StandardMoroccanTamazight QLocale__Language = QLocale__Language(314)
	QLocale__Mapuche                   QLocale__Language = QLocale__Language(315)
	QLocale__CentralKurdish            QLocale__Language = QLocale__Language(316)
	QLocale__LowerSorbian              QLocale__Language = QLocale__Language(317)
	QLocale__UpperSorbian              QLocale__Language = QLocale__Language(318)
	QLocale__Kenyang                   QLocale__Language = QLocale__Language(319)
	QLocale__Mohawk                    QLocale__Language = QLocale__Language(320)
	QLocale__Nko                       QLocale__Language = QLocale__Language(321)
	QLocale__Prussian                  QLocale__Language = QLocale__Language(322)
	QLocale__Kiche                     QLocale__Language = QLocale__Language(323)
	QLocale__SouthernSami              QLocale__Language = QLocale__Language(324)
	QLocale__LuleSami                  QLocale__Language = QLocale__Language(325)
	QLocale__InariSami                 QLocale__Language = QLocale__Language(326)
	QLocale__SkoltSami                 QLocale__Language = QLocale__Language(327)
	QLocale__Warlpiri                  QLocale__Language = QLocale__Language(328)
	QLocale__ManichaeanMiddlePersian   QLocale__Language = QLocale__Language(329)
	QLocale__Mende                     QLocale__Language = QLocale__Language(330)
	QLocale__AncientNorthArabian       QLocale__Language = QLocale__Language(331)
	QLocale__LinearA                   QLocale__Language = QLocale__Language(332)
	QLocale__HmongNjua                 QLocale__Language = QLocale__Language(333)
	QLocale__Ho                        QLocale__Language = QLocale__Language(334)
	QLocale__Lezghian                  QLocale__Language = QLocale__Language(335)
	QLocale__Bassa                     QLocale__Language = QLocale__Language(336)
	QLocale__Mono                      QLocale__Language = QLocale__Language(337)
	QLocale__TedimChin                 QLocale__Language = QLocale__Language(338)
	QLocale__Maithili                  QLocale__Language = QLocale__Language(339)
	QLocale__Ahom                      QLocale__Language = QLocale__Language(340)
	QLocale__AmericanSignLanguage      QLocale__Language = QLocale__Language(341)
	QLocale__ArdhamagadhiPrakrit       QLocale__Language = QLocale__Language(342)
	QLocale__Bhojpuri                  QLocale__Language = QLocale__Language(343)
	QLocale__HieroglyphicLuwian        QLocale__Language = QLocale__Language(344)
	QLocale__LiteraryChinese           QLocale__Language = QLocale__Language(345)
	QLocale__Mazanderani               QLocale__Language = QLocale__Language(346)
	QLocale__Mru                       QLocale__Language = QLocale__Language(347)
	QLocale__Newari                    QLocale__Language = QLocale__Language(348)
	QLocale__NorthernLuri              QLocale__Language = QLocale__Language(349)
	QLocale__Palauan                   QLocale__Language = QLocale__Language(350)
	QLocale__Papiamento                QLocale__Language = QLocale__Language(351)
	QLocale__Saraiki                   QLocale__Language = QLocale__Language(352)
	QLocale__TokelauLanguage           QLocale__Language = QLocale__Language(353)
	QLocale__TokPisin                  QLocale__Language = QLocale__Language(354)
	QLocale__TuvaluLanguage            QLocale__Language = QLocale__Language(355)
	QLocale__UncodedLanguages          QLocale__Language = QLocale__Language(356)
	QLocale__Cantonese                 QLocale__Language = QLocale__Language(357)
	QLocale__Osage                     QLocale__Language = QLocale__Language(358)
	QLocale__Tangut                    QLocale__Language = QLocale__Language(359)
	QLocale__Ido                       QLocale__Language = QLocale__Language(360)
	QLocale__Lojban                    QLocale__Language = QLocale__Language(361)
	QLocale__Sicilian                  QLocale__Language = QLocale__Language(362)
	QLocale__SouthernKurdish           QLocale__Language = QLocale__Language(363)
	QLocale__WesternBalochi            QLocale__Language = QLocale__Language(364)
	QLocale__Cebuano                   QLocale__Language = QLocale__Language(365)
	QLocale__Erzya                     QLocale__Language = QLocale__Language(366)
	QLocale__Chickasaw                 QLocale__Language = QLocale__Language(367)
	QLocale__Muscogee                  QLocale__Language = QLocale__Language(368)
	QLocale__Silesian                  QLocale__Language = QLocale__Language(369)
	QLocale__NigerianPidgin            QLocale__Language = QLocale__Language(370)
	QLocale__Afan                      QLocale__Language = QLocale__Language(QLocale__Oromo)
	QLocale__Bhutani                   QLocale__Language = QLocale__Language(QLocale__Dzongkha)
	QLocale__Byelorussian              QLocale__Language = QLocale__Language(QLocale__Belarusian)
	QLocale__Cambodian                 QLocale__Language = QLocale__Language(QLocale__Khmer)
	QLocale__Chewa                     QLocale__Language = QLocale__Language(QLocale__Nyanja)
	QLocale__Frisian                   QLocale__Language = QLocale__Language(QLocale__WesternFrisian)
	QLocale__Kurundi                   QLocale__Language = QLocale__Language(QLocale__Rundi)
	QLocale__Moldavian                 QLocale__Language = QLocale__Language(QLocale__Romanian)
	QLocale__Norwegian                 QLocale__Language = QLocale__Language(QLocale__NorwegianBokmal)
	QLocale__RhaetoRomance             QLocale__Language = QLocale__Language(QLocale__Romansh)
	QLocale__SerboCroatian             QLocale__Language = QLocale__Language(QLocale__Serbian)
	QLocale__Tagalog                   QLocale__Language = QLocale__Language(QLocale__Filipino)
	QLocale__Twi                       QLocale__Language = QLocale__Language(QLocale__Akan)
	QLocale__Uigur                     QLocale__Language = QLocale__Language(QLocale__Uighur)
	QLocale__LastLanguage              QLocale__Language = QLocale__Language(QLocale__NigerianPidgin)
)

// QLocale::MeasurementSystem
//
//go:generate stringer -type=QLocale__MeasurementSystem
type QLocale__MeasurementSystem int64

const (
	QLocale__MetricSystem     QLocale__MeasurementSystem = QLocale__MeasurementSystem(0)
	QLocale__ImperialUSSystem QLocale__MeasurementSystem = QLocale__MeasurementSystem(1)
	QLocale__ImperialUKSystem QLocale__MeasurementSystem = QLocale__MeasurementSystem(2)
	QLocale__ImperialSystem   QLocale__MeasurementSystem = QLocale__MeasurementSystem(QLocale__ImperialUSSystem)
)

// QLocale::NumberOption
//
//go:generate stringer -type=QLocale__NumberOption
type QLocale__NumberOption int64

const (
	QLocale__DefaultNumberOptions          QLocale__NumberOption = QLocale__NumberOption(0x0)
	QLocale__OmitGroupSeparator            QLocale__NumberOption = QLocale__NumberOption(0x01)
	QLocale__RejectGroupSeparator          QLocale__NumberOption = QLocale__NumberOption(0x02)
	QLocale__OmitLeadingZeroInExponent     QLocale__NumberOption = QLocale__NumberOption(0x04)
	QLocale__RejectLeadingZeroInExponent   QLocale__NumberOption = QLocale__NumberOption(0x08)
	QLocale__IncludeTrailingZeroesAfterDot QLocale__NumberOption = QLocale__NumberOption(0x10)
	QLocale__RejectTrailingZeroesAfterDot  QLocale__NumberOption = QLocale__NumberOption(0x20)
)

// QLocale::QuotationStyle
//
//go:generate stringer -type=QLocale__QuotationStyle
type QLocale__QuotationStyle int64

var (
	QLocale__StandardQuotation  QLocale__QuotationStyle = QLocale__QuotationStyle(0)
	QLocale__AlternateQuotation QLocale__QuotationStyle = QLocale__QuotationStyle(1)
)

// QLocale::Script
//
//go:generate stringer -type=QLocale__Script
type QLocale__Script int64

const (
	QLocale__AnyScript                   QLocale__Script = QLocale__Script(0)
	QLocale__ArabicScript                QLocale__Script = QLocale__Script(1)
	QLocale__CyrillicScript              QLocale__Script = QLocale__Script(2)
	QLocale__DeseretScript               QLocale__Script = QLocale__Script(3)
	QLocale__GurmukhiScript              QLocale__Script = QLocale__Script(4)
	QLocale__SimplifiedHanScript         QLocale__Script = QLocale__Script(5)
	QLocale__TraditionalHanScript        QLocale__Script = QLocale__Script(6)
	QLocale__LatinScript                 QLocale__Script = QLocale__Script(7)
	QLocale__MongolianScript             QLocale__Script = QLocale__Script(8)
	QLocale__TifinaghScript              QLocale__Script = QLocale__Script(9)
	QLocale__ArmenianScript              QLocale__Script = QLocale__Script(10)
	QLocale__BengaliScript               QLocale__Script = QLocale__Script(11)
	QLocale__CherokeeScript              QLocale__Script = QLocale__Script(12)
	QLocale__DevanagariScript            QLocale__Script = QLocale__Script(13)
	QLocale__EthiopicScript              QLocale__Script = QLocale__Script(14)
	QLocale__GeorgianScript              QLocale__Script = QLocale__Script(15)
	QLocale__GreekScript                 QLocale__Script = QLocale__Script(16)
	QLocale__GujaratiScript              QLocale__Script = QLocale__Script(17)
	QLocale__HebrewScript                QLocale__Script = QLocale__Script(18)
	QLocale__JapaneseScript              QLocale__Script = QLocale__Script(19)
	QLocale__KhmerScript                 QLocale__Script = QLocale__Script(20)
	QLocale__KannadaScript               QLocale__Script = QLocale__Script(21)
	QLocale__KoreanScript                QLocale__Script = QLocale__Script(22)
	QLocale__LaoScript                   QLocale__Script = QLocale__Script(23)
	QLocale__MalayalamScript             QLocale__Script = QLocale__Script(24)
	QLocale__MyanmarScript               QLocale__Script = QLocale__Script(25)
	QLocale__OriyaScript                 QLocale__Script = QLocale__Script(26)
	QLocale__TamilScript                 QLocale__Script = QLocale__Script(27)
	QLocale__TeluguScript                QLocale__Script = QLocale__Script(28)
	QLocale__ThaanaScript                QLocale__Script = QLocale__Script(29)
	QLocale__ThaiScript                  QLocale__Script = QLocale__Script(30)
	QLocale__TibetanScript               QLocale__Script = QLocale__Script(31)
	QLocale__SinhalaScript               QLocale__Script = QLocale__Script(32)
	QLocale__SyriacScript                QLocale__Script = QLocale__Script(33)
	QLocale__YiScript                    QLocale__Script = QLocale__Script(34)
	QLocale__VaiScript                   QLocale__Script = QLocale__Script(35)
	QLocale__AvestanScript               QLocale__Script = QLocale__Script(36)
	QLocale__BalineseScript              QLocale__Script = QLocale__Script(37)
	QLocale__BamumScript                 QLocale__Script = QLocale__Script(38)
	QLocale__BatakScript                 QLocale__Script = QLocale__Script(39)
	QLocale__BopomofoScript              QLocale__Script = QLocale__Script(40)
	QLocale__BrahmiScript                QLocale__Script = QLocale__Script(41)
	QLocale__BugineseScript              QLocale__Script = QLocale__Script(42)
	QLocale__BuhidScript                 QLocale__Script = QLocale__Script(43)
	QLocale__CanadianAboriginalScript    QLocale__Script = QLocale__Script(44)
	QLocale__CarianScript                QLocale__Script = QLocale__Script(45)
	QLocale__ChakmaScript                QLocale__Script = QLocale__Script(46)
	QLocale__ChamScript                  QLocale__Script = QLocale__Script(47)
	QLocale__CopticScript                QLocale__Script = QLocale__Script(48)
	QLocale__CypriotScript               QLocale__Script = QLocale__Script(49)
	QLocale__EgyptianHieroglyphsScript   QLocale__Script = QLocale__Script(50)
	QLocale__FraserScript                QLocale__Script = QLocale__Script(51)
	QLocale__GlagoliticScript            QLocale__Script = QLocale__Script(52)
	QLocale__GothicScript                QLocale__Script = QLocale__Script(53)
	QLocale__HanScript                   QLocale__Script = QLocale__Script(54)
	QLocale__HangulScript                QLocale__Script = QLocale__Script(55)
	QLocale__HanunooScript               QLocale__Script = QLocale__Script(56)
	QLocale__ImperialAramaicScript       QLocale__Script = QLocale__Script(57)
	QLocale__InscriptionalPahlaviScript  QLocale__Script = QLocale__Script(58)
	QLocale__InscriptionalParthianScript QLocale__Script = QLocale__Script(59)
	QLocale__JavaneseScript              QLocale__Script = QLocale__Script(60)
	QLocale__KaithiScript                QLocale__Script = QLocale__Script(61)
	QLocale__KatakanaScript              QLocale__Script = QLocale__Script(62)
	QLocale__KayahLiScript               QLocale__Script = QLocale__Script(63)
	QLocale__KharoshthiScript            QLocale__Script = QLocale__Script(64)
	QLocale__LannaScript                 QLocale__Script = QLocale__Script(65)
	QLocale__LepchaScript                QLocale__Script = QLocale__Script(66)
	QLocale__LimbuScript                 QLocale__Script = QLocale__Script(67)
	QLocale__LinearBScript               QLocale__Script = QLocale__Script(68)
	QLocale__LycianScript                QLocale__Script = QLocale__Script(69)
	QLocale__LydianScript                QLocale__Script = QLocale__Script(70)
	QLocale__MandaeanScript              QLocale__Script = QLocale__Script(71)
	QLocale__MeiteiMayekScript           QLocale__Script = QLocale__Script(72)
	QLocale__MeroiticScript              QLocale__Script = QLocale__Script(73)
	QLocale__MeroiticCursiveScript       QLocale__Script = QLocale__Script(74)
	QLocale__NkoScript                   QLocale__Script = QLocale__Script(75)
	QLocale__NewTaiLueScript             QLocale__Script = QLocale__Script(76)
	QLocale__OghamScript                 QLocale__Script = QLocale__Script(77)
	QLocale__OlChikiScript               QLocale__Script = QLocale__Script(78)
	QLocale__OldItalicScript             QLocale__Script = QLocale__Script(79)
	QLocale__OldPersianScript            QLocale__Script = QLocale__Script(80)
	QLocale__OldSouthArabianScript       QLocale__Script = QLocale__Script(81)
	QLocale__OrkhonScript                QLocale__Script = QLocale__Script(82)
	QLocale__OsmanyaScript               QLocale__Script = QLocale__Script(83)
	QLocale__PhagsPaScript               QLocale__Script = QLocale__Script(84)
	QLocale__PhoenicianScript            QLocale__Script = QLocale__Script(85)
	QLocale__PollardPhoneticScript       QLocale__Script = QLocale__Script(86)
	QLocale__RejangScript                QLocale__Script = QLocale__Script(87)
	QLocale__RunicScript                 QLocale__Script = QLocale__Script(88)
	QLocale__SamaritanScript             QLocale__Script = QLocale__Script(89)
	QLocale__SaurashtraScript            QLocale__Script = QLocale__Script(90)
	QLocale__SharadaScript               QLocale__Script = QLocale__Script(91)
	QLocale__ShavianScript               QLocale__Script = QLocale__Script(92)
	QLocale__SoraSompengScript           QLocale__Script = QLocale__Script(93)
	QLocale__CuneiformScript             QLocale__Script = QLocale__Script(94)
	QLocale__SundaneseScript             QLocale__Script = QLocale__Script(95)
	QLocale__SylotiNagriScript           QLocale__Script = QLocale__Script(96)
	QLocale__TagalogScript               QLocale__Script = QLocale__Script(97)
	QLocale__TagbanwaScript              QLocale__Script = QLocale__Script(98)
	QLocale__TaiLeScript                 QLocale__Script = QLocale__Script(99)
	QLocale__TaiVietScript               QLocale__Script = QLocale__Script(100)
	QLocale__TakriScript                 QLocale__Script = QLocale__Script(101)
	QLocale__UgariticScript              QLocale__Script = QLocale__Script(102)
	QLocale__BrailleScript               QLocale__Script = QLocale__Script(103)
	QLocale__HiraganaScript              QLocale__Script = QLocale__Script(104)
	QLocale__CaucasianAlbanianScript     QLocale__Script = QLocale__Script(105)
	QLocale__BassaVahScript              QLocale__Script = QLocale__Script(106)
	QLocale__DuployanScript              QLocale__Script = QLocale__Script(107)
	QLocale__ElbasanScript               QLocale__Script = QLocale__Script(108)
	QLocale__GranthaScript               QLocale__Script = QLocale__Script(109)
	QLocale__PahawhHmongScript           QLocale__Script = QLocale__Script(110)
	QLocale__KhojkiScript                QLocale__Script = QLocale__Script(111)
	QLocale__LinearAScript               QLocale__Script = QLocale__Script(112)
	QLocale__MahajaniScript              QLocale__Script = QLocale__Script(113)
	QLocale__ManichaeanScript            QLocale__Script = QLocale__Script(114)
	QLocale__MendeKikakuiScript          QLocale__Script = QLocale__Script(115)
	QLocale__ModiScript                  QLocale__Script = QLocale__Script(116)
	QLocale__MroScript                   QLocale__Script = QLocale__Script(117)
	QLocale__OldNorthArabianScript       QLocale__Script = QLocale__Script(118)
	QLocale__NabataeanScript             QLocale__Script = QLocale__Script(119)
	QLocale__PalmyreneScript             QLocale__Script = QLocale__Script(120)
	QLocale__PauCinHauScript             QLocale__Script = QLocale__Script(121)
	QLocale__OldPermicScript             QLocale__Script = QLocale__Script(122)
	QLocale__PsalterPahlaviScript        QLocale__Script = QLocale__Script(123)
	QLocale__SiddhamScript               QLocale__Script = QLocale__Script(124)
	QLocale__KhudawadiScript             QLocale__Script = QLocale__Script(125)
	QLocale__TirhutaScript               QLocale__Script = QLocale__Script(126)
	QLocale__VarangKshitiScript          QLocale__Script = QLocale__Script(127)
	QLocale__AhomScript                  QLocale__Script = QLocale__Script(128)
	QLocale__AnatolianHieroglyphsScript  QLocale__Script = QLocale__Script(129)
	QLocale__HatranScript                QLocale__Script = QLocale__Script(130)
	QLocale__MultaniScript               QLocale__Script = QLocale__Script(131)
	QLocale__OldHungarianScript          QLocale__Script = QLocale__Script(132)
	QLocale__SignWritingScript           QLocale__Script = QLocale__Script(133)
	QLocale__AdlamScript                 QLocale__Script = QLocale__Script(134)
	QLocale__BhaiksukiScript             QLocale__Script = QLocale__Script(135)
	QLocale__MarchenScript               QLocale__Script = QLocale__Script(136)
	QLocale__NewaScript                  QLocale__Script = QLocale__Script(137)
	QLocale__OsageScript                 QLocale__Script = QLocale__Script(138)
	QLocale__TangutScript                QLocale__Script = QLocale__Script(139)
	QLocale__HanWithBopomofoScript       QLocale__Script = QLocale__Script(140)
	QLocale__JamoScript                  QLocale__Script = QLocale__Script(141)
	QLocale__SimplifiedChineseScript     QLocale__Script = QLocale__Script(QLocale__SimplifiedHanScript)
	QLocale__TraditionalChineseScript    QLocale__Script = QLocale__Script(QLocale__TraditionalHanScript)
	QLocale__LastScript                  QLocale__Script = QLocale__Script(QLocale__JamoScript)
)

func NewQLocale() *QLocale {
	tmpValue := NewQLocaleFromPointer(C.QLocale_NewQLocale())
	qt.SetFinalizer(tmpValue, (*QLocale).DestroyQLocale)
	return tmpValue
}

func NewQLocale2(name string) *QLocale {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQLocaleFromPointer(C.QLocale_NewQLocale2(C.struct_QtCore_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QLocale).DestroyQLocale)
	return tmpValue
}

func NewQLocale3(language QLocale__Language, country QLocale__Country) *QLocale {
	tmpValue := NewQLocaleFromPointer(C.QLocale_NewQLocale3(C.longlong(language), C.longlong(country)))
	qt.SetFinalizer(tmpValue, (*QLocale).DestroyQLocale)
	return tmpValue
}

func NewQLocale4(language QLocale__Language, scri QLocale__Script, country QLocale__Country) *QLocale {
	tmpValue := NewQLocaleFromPointer(C.QLocale_NewQLocale4(C.longlong(language), C.longlong(scri), C.longlong(country)))
	qt.SetFinalizer(tmpValue, (*QLocale).DestroyQLocale)
	return tmpValue
}

func NewQLocale5(other QLocale_ITF) *QLocale {
	tmpValue := NewQLocaleFromPointer(C.QLocale_NewQLocale5(PointerFromQLocale(other)))
	qt.SetFinalizer(tmpValue, (*QLocale).DestroyQLocale)
	return tmpValue
}

func QLocale_C() *QLocale {
	tmpValue := NewQLocaleFromPointer(C.QLocale_QLocale_C())
	qt.SetFinalizer(tmpValue, (*QLocale).DestroyQLocale)
	return tmpValue
}

func (ptr *QLocale) C() *QLocale {
	tmpValue := NewQLocaleFromPointer(C.QLocale_QLocale_C())
	qt.SetFinalizer(tmpValue, (*QLocale).DestroyQLocale)
	return tmpValue
}

func (ptr *QLocale) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLocale_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocale) ToInt(s string, ok *bool) int {
	if ptr.Pointer() != nil {
		var sC *C.char
		if s != "" {
			sC = C.CString(s)
			defer C.free(unsafe.Pointer(sC))
		}
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return int(int32(C.QLocale_ToInt(ptr.Pointer(), C.struct_QtCore_PackedString{data: sC, len: C.longlong(len(s))}, &okC)))
	}
	return 0
}

func (ptr *QLocale) ToInt2(s QStringRef_ITF, ok *bool) int {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return int(int32(C.QLocale_ToInt2(ptr.Pointer(), PointerFromQStringRef(s), &okC)))
	}
	return 0
}

func (ptr *QLocale) ToInt3(s QStringView_ITF, ok *bool) int {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return int(int32(C.QLocale_ToInt3(ptr.Pointer(), PointerFromQStringView(s), &okC)))
	}
	return 0
}

func (ptr *QLocale) ToLower(str string) string {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		return cGoUnpackString(C.QLocale_ToLower(ptr.Pointer(), C.struct_QtCore_PackedString{data: strC, len: C.longlong(len(str))}))
	}
	return ""
}

func (ptr *QLocale) DestroyQLocale() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLocale_DestroyQLocale(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocale) __matchingLocales_atList(i int) *QLocale {
	if ptr.Pointer() != nil {
		tmpValue := NewQLocaleFromPointer(C.QLocale___matchingLocales_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

func (ptr *QLocale) __matchingLocales_setList(i QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QLocale___matchingLocales_setList(ptr.Pointer(), PointerFromQLocale(i))
	}
}

func (ptr *QLocale) __matchingLocales_newList() unsafe.Pointer {
	return C.QLocale___matchingLocales_newList(ptr.Pointer())
}

func (ptr *QLocale) __weekdays_atList(i int) Qt__DayOfWeek {
	if ptr.Pointer() != nil {
		return Qt__DayOfWeek(C.QLocale___weekdays_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QLocale) __weekdays_setList(i Qt__DayOfWeek) {
	if ptr.Pointer() != nil {
		C.QLocale___weekdays_setList(ptr.Pointer(), C.longlong(i))
	}
}

func (ptr *QLocale) __weekdays_newList() unsafe.Pointer {
	return C.QLocale___weekdays_newList(ptr.Pointer())
}

// QLockFile::LockError
//
//go:generate stringer -type=QLockFile__LockError
type QLockFile__LockError int64

const (
	QLockFile__NoError         QLockFile__LockError = QLockFile__LockError(0)
	QLockFile__LockFailedError QLockFile__LockError = QLockFile__LockError(1)
	QLockFile__PermissionError QLockFile__LockError = QLockFile__LockError(2)
	QLockFile__UnknownError    QLockFile__LockError = QLockFile__LockError(3)
)

type QMetaMethod struct {
	ptr unsafe.Pointer
}

type QMetaMethod_ITF interface {
	QMetaMethod_PTR() *QMetaMethod
}

func (ptr *QMetaMethod) QMetaMethod_PTR() *QMetaMethod {
	return ptr
}

func (ptr *QMetaMethod) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QMetaMethod) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQMetaMethod(ptr QMetaMethod_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaMethod_PTR().Pointer()
	}
	return nil
}

func NewQMetaMethodFromPointer(ptr unsafe.Pointer) (n *QMetaMethod) {
	n = new(QMetaMethod)
	n.SetPointer(ptr)
	return
}
func (ptr *QMetaMethod) DestroyQMetaMethod() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QMetaMethod__Access
//QMetaMethod::Access
type QMetaMethod__Access int64

const (
	QMetaMethod__Private   QMetaMethod__Access = QMetaMethod__Access(0)
	QMetaMethod__Protected QMetaMethod__Access = QMetaMethod__Access(1)
	QMetaMethod__Public    QMetaMethod__Access = QMetaMethod__Access(2)
)

// QMetaMethod::MethodType
//
//go:generate stringer -type=QMetaMethod__MethodType
type QMetaMethod__MethodType int64

const (
	QMetaMethod__Method      QMetaMethod__MethodType = QMetaMethod__MethodType(0)
	QMetaMethod__Signal      QMetaMethod__MethodType = QMetaMethod__MethodType(1)
	QMetaMethod__Slot        QMetaMethod__MethodType = QMetaMethod__MethodType(2)
	QMetaMethod__Constructor QMetaMethod__MethodType = QMetaMethod__MethodType(3)
)

func (ptr *QMetaMethod) Access() QMetaMethod__Access {
	if ptr.Pointer() != nil {
		return QMetaMethod__Access(C.QMetaMethod_Access(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaMethod) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QMetaMethod_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaMethod) Name() *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QMetaMethod_Name(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QMetaMethod) __parameterNames_atList(i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QMetaMethod___parameterNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QMetaMethod) __parameterNames_setList(i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMetaMethod___parameterNames_setList(ptr.Pointer(), PointerFromQByteArray(i))
	}
}

func (ptr *QMetaMethod) __parameterNames_newList() unsafe.Pointer {
	return C.QMetaMethod___parameterNames_newList(ptr.Pointer())
}

func (ptr *QMetaMethod) __parameterTypes_atList(i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QMetaMethod___parameterTypes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QMetaMethod) __parameterTypes_setList(i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMetaMethod___parameterTypes_setList(ptr.Pointer(), PointerFromQByteArray(i))
	}
}

func (ptr *QMetaMethod) __parameterTypes_newList() unsafe.Pointer {
	return C.QMetaMethod___parameterTypes_newList(ptr.Pointer())
}

type QMetaObject struct {
	ptr unsafe.Pointer
}

type QMetaObject_ITF interface {
	QMetaObject_PTR() *QMetaObject
}

func (ptr *QMetaObject) QMetaObject_PTR() *QMetaObject {
	return ptr
}

func (ptr *QMetaObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QMetaObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQMetaObject(ptr QMetaObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaObject_PTR().Pointer()
	}
	return nil
}

func NewQMetaObjectFromPointer(ptr unsafe.Pointer) (n *QMetaObject) {
	n = new(QMetaObject)
	n.SetPointer(ptr)
	return
}
func (ptr *QMetaObject) DestroyQMetaObject() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func (ptr *QMetaObject) Constructor(index int) *QMetaMethod {
	if ptr.Pointer() != nil {
		tmpValue := NewQMetaMethodFromPointer(C.QMetaObject_Constructor(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*QMetaMethod).DestroyQMetaMethod)
		return tmpValue
	}
	return nil
}

func (ptr *QMetaObject) Method(index int) *QMetaMethod {
	if ptr.Pointer() != nil {
		tmpValue := NewQMetaMethodFromPointer(C.QMetaObject_Method(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*QMetaMethod).DestroyQMetaMethod)
		return tmpValue
	}
	return nil
}

type QMetaProperty struct {
	ptr unsafe.Pointer
}

type QMetaProperty_ITF interface {
	QMetaProperty_PTR() *QMetaProperty
}

func (ptr *QMetaProperty) QMetaProperty_PTR() *QMetaProperty {
	return ptr
}

func (ptr *QMetaProperty) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QMetaProperty) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQMetaProperty(ptr QMetaProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaProperty_PTR().Pointer()
	}
	return nil
}

func NewQMetaPropertyFromPointer(ptr unsafe.Pointer) (n *QMetaProperty) {
	n = new(QMetaProperty)
	n.SetPointer(ptr)
	return
}
func (ptr *QMetaProperty) DestroyQMetaProperty() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func (ptr *QMetaProperty) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QMetaProperty_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaProperty) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QMetaProperty_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMetaProperty) Read(object QObject_ITF) *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QMetaProperty_Read(ptr.Pointer(), PointerFromQObject(object)))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QMetaProperty) Type() QVariant__Type {
	if ptr.Pointer() != nil {
		return QVariant__Type(C.QMetaProperty_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaProperty) Write(object QObject_ITF, value QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QMetaProperty_Write(ptr.Pointer(), PointerFromQObject(object), PointerFromQVariant(value))) != 0
	}
	return false
}

// QMetaType::Type
//
//go:generate stringer -type=QMetaType__Type
type QMetaType__Type int64

const (
	QMetaType__UnknownType           QMetaType__Type = QMetaType__Type(0)
	QMetaType__Bool                  QMetaType__Type = QMetaType__Type(1)
	QMetaType__Int                   QMetaType__Type = QMetaType__Type(2)
	QMetaType__UInt                  QMetaType__Type = QMetaType__Type(3)
	QMetaType__LongLong              QMetaType__Type = QMetaType__Type(4)
	QMetaType__ULongLong             QMetaType__Type = QMetaType__Type(5)
	QMetaType__Double                QMetaType__Type = QMetaType__Type(6)
	QMetaType__Long                  QMetaType__Type = QMetaType__Type(32)
	QMetaType__Short                 QMetaType__Type = QMetaType__Type(33)
	QMetaType__Char                  QMetaType__Type = QMetaType__Type(34)
	QMetaType__ULong                 QMetaType__Type = QMetaType__Type(35)
	QMetaType__UShort                QMetaType__Type = QMetaType__Type(36)
	QMetaType__UChar                 QMetaType__Type = QMetaType__Type(37)
	QMetaType__Float                 QMetaType__Type = QMetaType__Type(38)
	QMetaType__VoidStar              QMetaType__Type = QMetaType__Type(31)
	QMetaType__QChar                 QMetaType__Type = QMetaType__Type(7)
	QMetaType__QString               QMetaType__Type = QMetaType__Type(10)
	QMetaType__QStringList           QMetaType__Type = QMetaType__Type(11)
	QMetaType__QByteArray            QMetaType__Type = QMetaType__Type(12)
	QMetaType__QBitArray             QMetaType__Type = QMetaType__Type(13)
	QMetaType__QDate                 QMetaType__Type = QMetaType__Type(14)
	QMetaType__QTime                 QMetaType__Type = QMetaType__Type(15)
	QMetaType__QDateTime             QMetaType__Type = QMetaType__Type(16)
	QMetaType__QUrl                  QMetaType__Type = QMetaType__Type(17)
	QMetaType__QLocale               QMetaType__Type = QMetaType__Type(18)
	QMetaType__QRect                 QMetaType__Type = QMetaType__Type(19)
	QMetaType__QRectF                QMetaType__Type = QMetaType__Type(20)
	QMetaType__QSize                 QMetaType__Type = QMetaType__Type(21)
	QMetaType__QSizeF                QMetaType__Type = QMetaType__Type(22)
	QMetaType__QLine                 QMetaType__Type = QMetaType__Type(23)
	QMetaType__QLineF                QMetaType__Type = QMetaType__Type(24)
	QMetaType__QPoint                QMetaType__Type = QMetaType__Type(25)
	QMetaType__QPointF               QMetaType__Type = QMetaType__Type(26)
	QMetaType__QRegExp               QMetaType__Type = QMetaType__Type(27)
	QMetaType__QEasingCurve          QMetaType__Type = QMetaType__Type(29)
	QMetaType__QUuid                 QMetaType__Type = QMetaType__Type(30)
	QMetaType__QVariant              QMetaType__Type = QMetaType__Type(41)
	QMetaType__QModelIndex           QMetaType__Type = QMetaType__Type(42)
	QMetaType__QPersistentModelIndex QMetaType__Type = QMetaType__Type(50)
	QMetaType__QRegularExpression    QMetaType__Type = QMetaType__Type(44)
	QMetaType__QJsonValue            QMetaType__Type = QMetaType__Type(45)
	QMetaType__QJsonObject           QMetaType__Type = QMetaType__Type(46)
	QMetaType__QJsonArray            QMetaType__Type = QMetaType__Type(47)
	QMetaType__QJsonDocument         QMetaType__Type = QMetaType__Type(48)
	QMetaType__QByteArrayList        QMetaType__Type = QMetaType__Type(49)
	QMetaType__QObjectStar           QMetaType__Type = QMetaType__Type(39)
	QMetaType__SChar                 QMetaType__Type = QMetaType__Type(40)
	QMetaType__Void                  QMetaType__Type = QMetaType__Type(43)
	QMetaType__Nullptr               QMetaType__Type = QMetaType__Type(51)
	QMetaType__QVariantMap           QMetaType__Type = QMetaType__Type(8)
	QMetaType__QVariantList          QMetaType__Type = QMetaType__Type(9)
	QMetaType__QVariantHash          QMetaType__Type = QMetaType__Type(28)
	QMetaType__QCborSimpleType       QMetaType__Type = QMetaType__Type(52)
	QMetaType__QCborValue            QMetaType__Type = QMetaType__Type(53)
	QMetaType__QCborArray            QMetaType__Type = QMetaType__Type(54)
	QMetaType__QCborMap              QMetaType__Type = QMetaType__Type(55)
	QMetaType__QFont                 QMetaType__Type = QMetaType__Type(64)
	QMetaType__QPixmap               QMetaType__Type = QMetaType__Type(65)
	QMetaType__QBrush                QMetaType__Type = QMetaType__Type(66)
	QMetaType__QColor                QMetaType__Type = QMetaType__Type(67)
	QMetaType__QPalette              QMetaType__Type = QMetaType__Type(68)
	QMetaType__QIcon                 QMetaType__Type = QMetaType__Type(69)
	QMetaType__QImage                QMetaType__Type = QMetaType__Type(70)
	QMetaType__QPolygon              QMetaType__Type = QMetaType__Type(71)
	QMetaType__QRegion               QMetaType__Type = QMetaType__Type(72)
	QMetaType__QBitmap               QMetaType__Type = QMetaType__Type(73)
	QMetaType__QCursor               QMetaType__Type = QMetaType__Type(74)
	QMetaType__QKeySequence          QMetaType__Type = QMetaType__Type(75)
	QMetaType__QPen                  QMetaType__Type = QMetaType__Type(76)
	QMetaType__QTextLength           QMetaType__Type = QMetaType__Type(77)
	QMetaType__QTextFormat           QMetaType__Type = QMetaType__Type(78)
	QMetaType__QMatrix               QMetaType__Type = QMetaType__Type(79)
	QMetaType__QTransform            QMetaType__Type = QMetaType__Type(80)
	QMetaType__QMatrix4x4            QMetaType__Type = QMetaType__Type(81)
	QMetaType__QVector2D             QMetaType__Type = QMetaType__Type(82)
	QMetaType__QVector3D             QMetaType__Type = QMetaType__Type(83)
	QMetaType__QVector4D             QMetaType__Type = QMetaType__Type(84)
	QMetaType__QQuaternion           QMetaType__Type = QMetaType__Type(85)
	QMetaType__QPolygonF             QMetaType__Type = QMetaType__Type(86)
	QMetaType__QColorSpace           QMetaType__Type = QMetaType__Type(87)
	QMetaType__QSizePolicy           QMetaType__Type = QMetaType__Type(121)
	QMetaType__LastCoreType          QMetaType__Type = QMetaType__Type(QMetaType__QCborMap)
	QMetaType__LastGuiType           QMetaType__Type = QMetaType__Type(QMetaType__QColorSpace)
	QMetaType__User                  QMetaType__Type = QMetaType__Type(1024)
)

// QMetaType::TypeFlag
//
//go:generate stringer -type=QMetaType__TypeFlag
type QMetaType__TypeFlag int64

const (
	QMetaType__NeedsConstruction        QMetaType__TypeFlag = QMetaType__TypeFlag(0x1)
	QMetaType__NeedsDestruction         QMetaType__TypeFlag = QMetaType__TypeFlag(0x2)
	QMetaType__MovableType              QMetaType__TypeFlag = QMetaType__TypeFlag(0x4)
	QMetaType__PointerToQObject         QMetaType__TypeFlag = QMetaType__TypeFlag(0x8)
	QMetaType__IsEnumeration            QMetaType__TypeFlag = QMetaType__TypeFlag(0x10)
	QMetaType__SharedPointerToQObject   QMetaType__TypeFlag = QMetaType__TypeFlag(0x20)
	QMetaType__WeakPointerToQObject     QMetaType__TypeFlag = QMetaType__TypeFlag(0x40)
	QMetaType__TrackingPointerToQObject QMetaType__TypeFlag = QMetaType__TypeFlag(0x80)
	QMetaType__WasDeclaredAsMetaType    QMetaType__TypeFlag = QMetaType__TypeFlag(0x100)
	QMetaType__IsGadget                 QMetaType__TypeFlag = QMetaType__TypeFlag(0x200)
	QMetaType__PointerToGadget          QMetaType__TypeFlag = QMetaType__TypeFlag(0x400)
)

// QMimeDatabase::MatchMode
//
//go:generate stringer -type=QMimeDatabase__MatchMode
type QMimeDatabase__MatchMode int64

const (
	QMimeDatabase__MatchDefault   QMimeDatabase__MatchMode = QMimeDatabase__MatchMode(0x0)
	QMimeDatabase__MatchExtension QMimeDatabase__MatchMode = QMimeDatabase__MatchMode(0x1)
	QMimeDatabase__MatchContent   QMimeDatabase__MatchMode = QMimeDatabase__MatchMode(0x2)
)

type QModelIndex struct {
	ptr unsafe.Pointer
}

type QModelIndex_ITF interface {
	QModelIndex_PTR() *QModelIndex
}

func (ptr *QModelIndex) QModelIndex_PTR() *QModelIndex {
	return ptr
}

func (ptr *QModelIndex) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QModelIndex) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQModelIndex(ptr QModelIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModelIndex_PTR().Pointer()
	}
	return nil
}

func NewQModelIndexFromPointer(ptr unsafe.Pointer) (n *QModelIndex) {
	n = new(QModelIndex)
	n.SetPointer(ptr)
	return
}
func (ptr *QModelIndex) DestroyQModelIndex() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQModelIndex() *QModelIndex {
	tmpValue := NewQModelIndexFromPointer(C.QModelIndex_NewQModelIndex())
	qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
	return tmpValue
}

func (ptr *QModelIndex) Column() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModelIndex_Column(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModelIndex) Data(role int) *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QModelIndex_Data(ptr.Pointer(), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QModelIndex) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QModelIndex_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QModelIndex) Model() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractItemModelFromPointer(C.QModelIndex_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModelIndex) Parent() *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QModelIndex_Parent(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QModelIndex) Row() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModelIndex_Row(ptr.Pointer())))
	}
	return 0
}

// QMutex::RecursionMode
//
//go:generate stringer -type=QMutex__RecursionMode
type QMutex__RecursionMode int64

const (
	QMutex__NonRecursive QMutex__RecursionMode = QMutex__RecursionMode(0)
	QMutex__Recursive    QMutex__RecursionMode = QMutex__RecursionMode(1)
)

type QObject struct {
	ptr unsafe.Pointer
}

type QObject_ITF interface {
	QObject_PTR() *QObject
}

func (ptr *QObject) QObject_PTR() *QObject {
	return ptr
}

func (ptr *QObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQObject(ptr QObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func NewQObjectFromPointer(ptr unsafe.Pointer) (n *QObject) {
	n = new(QObject)
	n.SetPointer(ptr)
	return
}

func (ptr *QObject) ConnectSignal(f interface{}, a interface{}, t Qt__ConnectionType) {
	fn := strings.TrimSuffix(strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".Connect")[1], "-fm")
	qt.RegisterConnectionType(ptr.Pointer(), strings.ToLower(fn[:1])+fn[1:], int64(t))
	reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(a)})
}

func NewQObject(parent QObject_ITF) *QObject {
	tmpValue := NewQObjectFromPointer(C.QObject_NewQObject(PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQObject_ChildEvent
func callbackQObject_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*QChildEvent))(signal))(NewQChildEventFromPointer(event))
	} else {
		NewQObjectFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QObject) ConnectChildEvent(f func(event *QChildEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "childEvent"); signal != nil {
			f := func(event *QChildEvent) {
				(*(*func(*QChildEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "childEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "childEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "childEvent")
	}
}

func (ptr *QObject) ChildEvent(event QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QObject) ChildEventDefault(event QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QObject) Children() []*QObject {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []*QObject {
			out := make([]*QObject, int(l.len))
			tmpList := NewQObjectFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__children_atList(i)
			}
			return out
		}(C.QObject_Children(ptr.Pointer()))
	}
	return make([]*QObject, 0)
}

//export callbackQObject_ConnectNotify
func callbackQObject_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*QMetaMethod))(signal))(NewQMetaMethodFromPointer(sign))
	} else {
		NewQObjectFromPointer(ptr).ConnectNotifyDefault(NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QObject) ConnectConnectNotify(f func(sign *QMetaMethod)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "connectNotify"); signal != nil {
			f := func(sign *QMetaMethod) {
				(*(*func(*QMetaMethod))(signal))(sign)
				f(sign)
			}
			qt.ConnectSignal(ptr.Pointer(), "connectNotify", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectNotify", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "connectNotify")
	}
}

func (ptr *QObject) ConnectNotify(sign QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_ConnectNotify(ptr.Pointer(), PointerFromQMetaMethod(sign))
	}
}

func (ptr *QObject) ConnectNotifyDefault(sign QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_ConnectNotifyDefault(ptr.Pointer(), PointerFromQMetaMethod(sign))
	}
}

//export callbackQObject_CustomEvent
func callbackQObject_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*QEvent))(signal))(NewQEventFromPointer(event))
	} else {
		NewQObjectFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QObject) ConnectCustomEvent(f func(event *QEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "customEvent"); signal != nil {
			f := func(event *QEvent) {
				(*(*func(*QEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "customEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "customEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "customEvent")
	}
}

func (ptr *QObject) CustomEvent(event QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QObject) CustomEventDefault(event QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}

//export callbackQObject_DeleteLater
func callbackQObject_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QObject) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "deleteLater"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "deleteLater", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deleteLater", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "deleteLater")
	}
}

func (ptr *QObject) DeleteLater() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QObject_DeleteLater(ptr.Pointer())
	}
}

func (ptr *QObject) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QObject_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQObject_Destroyed
func callbackQObject_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*QObject))(signal))(NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

func (ptr *QObject) ConnectDestroyed(f func(obj *QObject)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "destroyed") {
			C.QObject_ConnectDestroyed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "destroyed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "destroyed"); signal != nil {
			f := func(obj *QObject) {
				(*(*func(*QObject))(signal))(obj)
				f(obj)
			}
			qt.ConnectSignal(ptr.Pointer(), "destroyed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "destroyed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectDestroyed() {
	if ptr.Pointer() != nil {
		C.QObject_DisconnectDestroyed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "destroyed")
	}
}

func (ptr *QObject) Destroyed(obj QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_Destroyed(ptr.Pointer(), PointerFromQObject(obj))
	}
}

func QObject_Disconnect(sender QObject_ITF, sign string, receiver QObject_ITF, method string) bool {
	var signC *C.char
	if sign != "" {
		signC = C.CString(sign)
		defer C.free(unsafe.Pointer(signC))
	}
	var methodC *C.char
	if method != "" {
		methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
	}
	return int8(C.QObject_QObject_Disconnect(PointerFromQObject(sender), signC, PointerFromQObject(receiver), methodC)) != 0
}

func (ptr *QObject) Disconnect(sender QObject_ITF, sign string, receiver QObject_ITF, method string) bool {
	var signC *C.char
	if sign != "" {
		signC = C.CString(sign)
		defer C.free(unsafe.Pointer(signC))
	}
	var methodC *C.char
	if method != "" {
		methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
	}
	return int8(C.QObject_QObject_Disconnect(PointerFromQObject(sender), signC, PointerFromQObject(receiver), methodC)) != 0
}

func QObject_Disconnect2(sender QObject_ITF, sign QMetaMethod_ITF, receiver QObject_ITF, method QMetaMethod_ITF) bool {
	return int8(C.QObject_QObject_Disconnect2(PointerFromQObject(sender), PointerFromQMetaMethod(sign), PointerFromQObject(receiver), PointerFromQMetaMethod(method))) != 0
}

func (ptr *QObject) Disconnect2(sender QObject_ITF, sign QMetaMethod_ITF, receiver QObject_ITF, method QMetaMethod_ITF) bool {
	return int8(C.QObject_QObject_Disconnect2(PointerFromQObject(sender), PointerFromQMetaMethod(sign), PointerFromQObject(receiver), PointerFromQMetaMethod(method))) != 0
}

func (ptr *QObject) Disconnect3(sign string, receiver QObject_ITF, method string) bool {
	if ptr.Pointer() != nil {
		var signC *C.char
		if sign != "" {
			signC = C.CString(sign)
			defer C.free(unsafe.Pointer(signC))
		}
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		return int8(C.QObject_Disconnect3(ptr.Pointer(), signC, PointerFromQObject(receiver), methodC)) != 0
	}
	return false
}

func (ptr *QObject) Disconnect4(receiver QObject_ITF, method string) bool {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		return int8(C.QObject_Disconnect4(ptr.Pointer(), PointerFromQObject(receiver), methodC)) != 0
	}
	return false
}

//export callbackQObject_DisconnectNotify
func callbackQObject_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*QMetaMethod))(signal))(NewQMetaMethodFromPointer(sign))
	} else {
		NewQObjectFromPointer(ptr).DisconnectNotifyDefault(NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QObject) ConnectDisconnectNotify(f func(sign *QMetaMethod)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "disconnectNotify"); signal != nil {
			f := func(sign *QMetaMethod) {
				(*(*func(*QMetaMethod))(signal))(sign)
				f(sign)
			}
			qt.ConnectSignal(ptr.Pointer(), "disconnectNotify", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnectNotify", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "disconnectNotify")
	}
}

func (ptr *QObject) DisconnectNotify(sign QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_DisconnectNotify(ptr.Pointer(), PointerFromQMetaMethod(sign))
	}
}

func (ptr *QObject) DisconnectNotifyDefault(sign QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_DisconnectNotifyDefault(ptr.Pointer(), PointerFromQMetaMethod(sign))
	}
}

//export callbackQObject_Event
func callbackQObject_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QEvent) bool)(signal))(NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQObjectFromPointer(ptr).EventDefault(NewQEventFromPointer(e)))))
}

func (ptr *QObject) ConnectEvent(f func(e *QEvent) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "event"); signal != nil {
			f := func(e *QEvent) bool {
				(*(*func(*QEvent) bool)(signal))(e)
				return f(e)
			}
			qt.ConnectSignal(ptr.Pointer(), "event", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "event", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "event")
	}
}

func (ptr *QObject) Event(e QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QObject_Event(ptr.Pointer(), PointerFromQEvent(e))) != 0
	}
	return false
}

func (ptr *QObject) EventDefault(e QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QObject_EventDefault(ptr.Pointer(), PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQObject_EventFilter
func callbackQObject_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QObject, *QEvent) bool)(signal))(NewQObjectFromPointer(watched), NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQObjectFromPointer(ptr).EventFilterDefault(NewQObjectFromPointer(watched), NewQEventFromPointer(event)))))
}

func (ptr *QObject) ConnectEventFilter(f func(watched *QObject, event *QEvent) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "eventFilter"); signal != nil {
			f := func(watched *QObject, event *QEvent) bool {
				(*(*func(*QObject, *QEvent) bool)(signal))(watched, event)
				return f(watched, event)
			}
			qt.ConnectSignal(ptr.Pointer(), "eventFilter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "eventFilter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "eventFilter")
	}
}

func (ptr *QObject) EventFilter(watched QObject_ITF, event QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QObject_EventFilter(ptr.Pointer(), PointerFromQObject(watched), PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QObject) EventFilterDefault(watched QObject_ITF, event QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QObject_EventFilterDefault(ptr.Pointer(), PointerFromQObject(watched), PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QObject) ObjectName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QObject_ObjectName(ptr.Pointer()))
	}
	return ""
}

//export callbackQObject_ObjectNameChanged
func callbackQObject_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtCore_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

func (ptr *QObject) ConnectObjectNameChanged(f func(objectName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "objectNameChanged") {
			C.QObject_ConnectObjectNameChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "objectNameChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "objectNameChanged"); signal != nil {
			f := func(objectName string) {
				(*(*func(string))(signal))(objectName)
				f(objectName)
			}
			qt.ConnectSignal(ptr.Pointer(), "objectNameChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "objectNameChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectObjectNameChanged() {
	if ptr.Pointer() != nil {
		C.QObject_DisconnectObjectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "objectNameChanged")
	}
}

func (ptr *QObject) Parent() *QObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQObjectFromPointer(C.QObject_Parent(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QObject) Property(name string) *QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQVariantFromPointer(C.QObject_Property(ptr.Pointer(), nameC))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QObject) SetObjectName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QObject_SetObjectName(ptr.Pointer(), C.struct_QtCore_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

//export callbackQObject_TimerEvent
func callbackQObject_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*QTimerEvent))(signal))(NewQTimerEventFromPointer(event))
	} else {
		NewQObjectFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QObject) ConnectTimerEvent(f func(event *QTimerEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "timerEvent"); signal != nil {
			f := func(event *QTimerEvent) {
				(*(*func(*QTimerEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "timerEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "timerEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "timerEvent")
	}
}

func (ptr *QObject) TimerEvent(event QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QObject) TimerEventDefault(event QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QObject_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func QObject_Tr(sourceText string, disambiguation string, n int) string {
	var sourceTextC *C.char
	if sourceText != "" {
		sourceTextC = C.CString(sourceText)
		defer C.free(unsafe.Pointer(sourceTextC))
	}
	var disambiguationC *C.char
	if disambiguation != "" {
		disambiguationC = C.CString(disambiguation)
		defer C.free(unsafe.Pointer(disambiguationC))
	}
	return cGoUnpackString(C.QObject_QObject_Tr(sourceTextC, disambiguationC, C.int(int32(n))))
}

func (ptr *QObject) Tr(sourceText string, disambiguation string, n int) string {
	var sourceTextC *C.char
	if sourceText != "" {
		sourceTextC = C.CString(sourceText)
		defer C.free(unsafe.Pointer(sourceTextC))
	}
	var disambiguationC *C.char
	if disambiguation != "" {
		disambiguationC = C.CString(disambiguation)
		defer C.free(unsafe.Pointer(disambiguationC))
	}
	return cGoUnpackString(C.QObject_QObject_Tr(sourceTextC, disambiguationC, C.int(int32(n))))
}

//export callbackQObject_DestroyQObject
func callbackQObject_DestroyQObject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QObject"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQObjectFromPointer(ptr).DestroyQObjectDefault()
	}
}

func (ptr *QObject) ConnectDestroyQObject(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QObject"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QObject", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QObject", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QObject) DisconnectDestroyQObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QObject")
	}
}

func (ptr *QObject) DestroyQObject() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QObject_DestroyQObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QObject) DestroyQObjectDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QObject_DestroyQObjectDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QObject) ToVariant() *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QObject_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QObject) __children_atList(i int) *QObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQObjectFromPointer(C.QObject___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QObject) __children_setList(i QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QObject___children_setList(ptr.Pointer(), PointerFromQObject(i))
	}
}

func (ptr *QObject) __children_newList() unsafe.Pointer {
	return C.QObject___children_newList(ptr.Pointer())
}

func (ptr *QObject) __dynamicPropertyNames_atList(i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QObject___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QObject) __dynamicPropertyNames_setList(i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QObject___dynamicPropertyNames_setList(ptr.Pointer(), PointerFromQByteArray(i))
	}
}

func (ptr *QObject) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QObject___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QObject) __findChildren_atList(i int) *QObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQObjectFromPointer(C.QObject___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QObject) __findChildren_setList(i QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QObject___findChildren_setList(ptr.Pointer(), PointerFromQObject(i))
	}
}

func (ptr *QObject) __findChildren_newList() unsafe.Pointer {
	return C.QObject___findChildren_newList(ptr.Pointer())
}

func (ptr *QObject) __findChildren_atList2(i int) *QObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQObjectFromPointer(C.QObject___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QObject) __findChildren_setList2(i QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QObject___findChildren_setList2(ptr.Pointer(), PointerFromQObject(i))
	}
}

func (ptr *QObject) __findChildren_newList2() unsafe.Pointer {
	return C.QObject___findChildren_newList2(ptr.Pointer())
}

func (ptr *QObject) __findChildren_atList3(i int) *QObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQObjectFromPointer(C.QObject___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QObject) __findChildren_setList3(i QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QObject___findChildren_setList3(ptr.Pointer(), PointerFromQObject(i))
	}
}

func (ptr *QObject) __findChildren_newList3() unsafe.Pointer {
	return C.QObject___findChildren_newList3(ptr.Pointer())
}

// QOperatingSystemVersion::OSType
//
//go:generate stringer -type=QOperatingSystemVersion__OSType
type QOperatingSystemVersion__OSType int64

const (
	QOperatingSystemVersion__Unknown QOperatingSystemVersion__OSType = QOperatingSystemVersion__OSType(0)
	QOperatingSystemVersion__Windows QOperatingSystemVersion__OSType = QOperatingSystemVersion__OSType(1)
	QOperatingSystemVersion__MacOS   QOperatingSystemVersion__OSType = QOperatingSystemVersion__OSType(2)
	QOperatingSystemVersion__IOS     QOperatingSystemVersion__OSType = QOperatingSystemVersion__OSType(3)
	QOperatingSystemVersion__TvOS    QOperatingSystemVersion__OSType = QOperatingSystemVersion__OSType(4)
	QOperatingSystemVersion__WatchOS QOperatingSystemVersion__OSType = QOperatingSystemVersion__OSType(5)
	QOperatingSystemVersion__Android QOperatingSystemVersion__OSType = QOperatingSystemVersion__OSType(6)
)

type QPersistentModelIndex struct {
	ptr unsafe.Pointer
}

type QPersistentModelIndex_ITF interface {
	QPersistentModelIndex_PTR() *QPersistentModelIndex
}

func (ptr *QPersistentModelIndex) QPersistentModelIndex_PTR() *QPersistentModelIndex {
	return ptr
}

func (ptr *QPersistentModelIndex) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPersistentModelIndex) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPersistentModelIndex(ptr QPersistentModelIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPersistentModelIndex_PTR().Pointer()
	}
	return nil
}

func NewQPersistentModelIndexFromPointer(ptr unsafe.Pointer) (n *QPersistentModelIndex) {
	n = new(QPersistentModelIndex)
	n.SetPointer(ptr)
	return
}
func (ptr *QPersistentModelIndex) DestroyQPersistentModelIndex() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQPersistentModelIndex2(index QModelIndex_ITF) *QPersistentModelIndex {
	tmpValue := NewQPersistentModelIndexFromPointer(C.QPersistentModelIndex_NewQPersistentModelIndex2(PointerFromQModelIndex(index)))
	qt.SetFinalizer(tmpValue, (*QPersistentModelIndex).DestroyQPersistentModelIndex)
	return tmpValue
}

func NewQPersistentModelIndex3(other QPersistentModelIndex_ITF) *QPersistentModelIndex {
	tmpValue := NewQPersistentModelIndexFromPointer(C.QPersistentModelIndex_NewQPersistentModelIndex3(PointerFromQPersistentModelIndex(other)))
	qt.SetFinalizer(tmpValue, (*QPersistentModelIndex).DestroyQPersistentModelIndex)
	return tmpValue
}

func NewQPersistentModelIndex4(other QPersistentModelIndex_ITF) *QPersistentModelIndex {
	tmpValue := NewQPersistentModelIndexFromPointer(C.QPersistentModelIndex_NewQPersistentModelIndex4(PointerFromQPersistentModelIndex(other)))
	qt.SetFinalizer(tmpValue, (*QPersistentModelIndex).DestroyQPersistentModelIndex)
	return tmpValue
}

func (ptr *QPersistentModelIndex) Column() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPersistentModelIndex_Column(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPersistentModelIndex) Data(role int) *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QPersistentModelIndex_Data(ptr.Pointer(), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QPersistentModelIndex) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPersistentModelIndex_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPersistentModelIndex) Model() *QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractItemModelFromPointer(C.QPersistentModelIndex_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPersistentModelIndex) Parent() *QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQModelIndexFromPointer(C.QPersistentModelIndex_Parent(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QPersistentModelIndex) Row() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPersistentModelIndex_Row(ptr.Pointer())))
	}
	return 0
}

type QPoint struct {
	ptr unsafe.Pointer
}

type QPoint_ITF interface {
	QPoint_PTR() *QPoint
}

func (ptr *QPoint) QPoint_PTR() *QPoint {
	return ptr
}

func (ptr *QPoint) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPoint) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPoint(ptr QPoint_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPoint_PTR().Pointer()
	}
	return nil
}

func NewQPointFromPointer(ptr unsafe.Pointer) (n *QPoint) {
	n = new(QPoint)
	n.SetPointer(ptr)
	return
}
func (ptr *QPoint) DestroyQPoint() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQPoint() *QPoint {
	tmpValue := NewQPointFromPointer(C.QPoint_NewQPoint())
	qt.SetFinalizer(tmpValue, (*QPoint).DestroyQPoint)
	return tmpValue
}

func NewQPoint2(xpos int, ypos int) *QPoint {
	tmpValue := NewQPointFromPointer(C.QPoint_NewQPoint2(C.int(int32(xpos)), C.int(int32(ypos))))
	qt.SetFinalizer(tmpValue, (*QPoint).DestroyQPoint)
	return tmpValue
}

func (ptr *QPoint) X() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPoint_X(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPoint) Y() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPoint_Y(ptr.Pointer())))
	}
	return 0
}

type QPointF struct {
	ptr unsafe.Pointer
}

type QPointF_ITF interface {
	QPointF_PTR() *QPointF
}

func (ptr *QPointF) QPointF_PTR() *QPointF {
	return ptr
}

func (ptr *QPointF) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPointF) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPointF(ptr QPointF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPointF_PTR().Pointer()
	}
	return nil
}

func NewQPointFFromPointer(ptr unsafe.Pointer) (n *QPointF) {
	n = new(QPointF)
	n.SetPointer(ptr)
	return
}
func (ptr *QPointF) DestroyQPointF() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQPointF() *QPointF {
	tmpValue := NewQPointFFromPointer(C.QPointF_NewQPointF())
	qt.SetFinalizer(tmpValue, (*QPointF).DestroyQPointF)
	return tmpValue
}

func NewQPointF2(point QPoint_ITF) *QPointF {
	tmpValue := NewQPointFFromPointer(C.QPointF_NewQPointF2(PointerFromQPoint(point)))
	qt.SetFinalizer(tmpValue, (*QPointF).DestroyQPointF)
	return tmpValue
}

func NewQPointF3(xpos float64, ypos float64) *QPointF {
	tmpValue := NewQPointFFromPointer(C.QPointF_NewQPointF3(C.double(xpos), C.double(ypos)))
	qt.SetFinalizer(tmpValue, (*QPointF).DestroyQPointF)
	return tmpValue
}

func (ptr *QPointF) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPointF_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPointF) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPointF_Y(ptr.Pointer()))
	}
	return 0
}

// QProcess::ExitStatus
//
//go:generate stringer -type=QProcess__ExitStatus
type QProcess__ExitStatus int64

const (
	QProcess__NormalExit QProcess__ExitStatus = QProcess__ExitStatus(0)
	QProcess__CrashExit  QProcess__ExitStatus = QProcess__ExitStatus(1)
)

// QProcess::InputChannelMode
//
//go:generate stringer -type=QProcess__InputChannelMode
type QProcess__InputChannelMode int64

const (
	QProcess__ManagedInputChannel   QProcess__InputChannelMode = QProcess__InputChannelMode(0)
	QProcess__ForwardedInputChannel QProcess__InputChannelMode = QProcess__InputChannelMode(1)
)

// QProcess::ProcessChannel
//
//go:generate stringer -type=QProcess__ProcessChannel
type QProcess__ProcessChannel int64

const (
	QProcess__StandardOutput QProcess__ProcessChannel = QProcess__ProcessChannel(0)
	QProcess__StandardError  QProcess__ProcessChannel = QProcess__ProcessChannel(1)
)

// QProcess::ProcessChannelMode
//
//go:generate stringer -type=QProcess__ProcessChannelMode
type QProcess__ProcessChannelMode int64

const (
	QProcess__SeparateChannels       QProcess__ProcessChannelMode = QProcess__ProcessChannelMode(0)
	QProcess__MergedChannels         QProcess__ProcessChannelMode = QProcess__ProcessChannelMode(1)
	QProcess__ForwardedChannels      QProcess__ProcessChannelMode = QProcess__ProcessChannelMode(2)
	QProcess__ForwardedOutputChannel QProcess__ProcessChannelMode = QProcess__ProcessChannelMode(3)
	QProcess__ForwardedErrorChannel  QProcess__ProcessChannelMode = QProcess__ProcessChannelMode(4)
)

// QProcess::ProcessError
//
//go:generate stringer -type=QProcess__ProcessError
type QProcess__ProcessError int64

const (
	QProcess__FailedToStart QProcess__ProcessError = QProcess__ProcessError(0)
	QProcess__Crashed       QProcess__ProcessError = QProcess__ProcessError(1)
	QProcess__Timedout      QProcess__ProcessError = QProcess__ProcessError(2)
	QProcess__ReadError     QProcess__ProcessError = QProcess__ProcessError(3)
	QProcess__WriteError    QProcess__ProcessError = QProcess__ProcessError(4)
	QProcess__UnknownError  QProcess__ProcessError = QProcess__ProcessError(5)
)

// QProcess::ProcessState
//
//go:generate stringer -type=QProcess__ProcessState
type QProcess__ProcessState int64

const (
	QProcess__NotRunning QProcess__ProcessState = QProcess__ProcessState(0)
	QProcess__Starting   QProcess__ProcessState = QProcess__ProcessState(1)
	QProcess__Running    QProcess__ProcessState = QProcess__ProcessState(2)
)

// QReadWriteLock::RecursionMode
//
//go:generate stringer -type=QReadWriteLock__RecursionMode
type QReadWriteLock__RecursionMode int64

const (
	QReadWriteLock__NonRecursive QReadWriteLock__RecursionMode = QReadWriteLock__RecursionMode(0)
	QReadWriteLock__Recursive    QReadWriteLock__RecursionMode = QReadWriteLock__RecursionMode(1)
)

type QRect struct {
	ptr unsafe.Pointer
}

type QRect_ITF interface {
	QRect_PTR() *QRect
}

func (ptr *QRect) QRect_PTR() *QRect {
	return ptr
}

func (ptr *QRect) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRect) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRect(ptr QRect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRect_PTR().Pointer()
	}
	return nil
}

func NewQRectFromPointer(ptr unsafe.Pointer) (n *QRect) {
	n = new(QRect)
	n.SetPointer(ptr)
	return
}
func (ptr *QRect) DestroyQRect() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQRect() *QRect {
	tmpValue := NewQRectFromPointer(C.QRect_NewQRect())
	qt.SetFinalizer(tmpValue, (*QRect).DestroyQRect)
	return tmpValue
}

func NewQRect2(topLeft QPoint_ITF, bottomRight QPoint_ITF) *QRect {
	tmpValue := NewQRectFromPointer(C.QRect_NewQRect2(PointerFromQPoint(topLeft), PointerFromQPoint(bottomRight)))
	qt.SetFinalizer(tmpValue, (*QRect).DestroyQRect)
	return tmpValue
}

func NewQRect3(topLeft QPoint_ITF, size QSize_ITF) *QRect {
	tmpValue := NewQRectFromPointer(C.QRect_NewQRect3(PointerFromQPoint(topLeft), PointerFromQSize(size)))
	qt.SetFinalizer(tmpValue, (*QRect).DestroyQRect)
	return tmpValue
}

func NewQRect4(x int, y int, width int, height int) *QRect {
	tmpValue := NewQRectFromPointer(C.QRect_NewQRect4(C.int(int32(x)), C.int(int32(y)), C.int(int32(width)), C.int(int32(height))))
	qt.SetFinalizer(tmpValue, (*QRect).DestroyQRect)
	return tmpValue
}

func (ptr *QRect) Center() *QPoint {
	if ptr.Pointer() != nil {
		tmpValue := NewQPointFromPointer(C.QRect_Center(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QRect) Height() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRect_Height(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRect_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRect) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRect_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRect) Left() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRect_Left(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) Right() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRect_Right(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) SetSize(size QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QRect_SetSize(ptr.Pointer(), PointerFromQSize(size))
	}
}

func (ptr *QRect) SetTop(y int) {
	if ptr.Pointer() != nil {
		C.QRect_SetTop(ptr.Pointer(), C.int(int32(y)))
	}
}

func (ptr *QRect) Size() *QSize {
	if ptr.Pointer() != nil {
		tmpValue := NewQSizeFromPointer(C.QRect_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QRect) Top() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRect_Top(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) Width() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRect_Width(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) X() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRect_X(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) Y() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRect_Y(ptr.Pointer())))
	}
	return 0
}

type QRectF struct {
	ptr unsafe.Pointer
}

type QRectF_ITF interface {
	QRectF_PTR() *QRectF
}

func (ptr *QRectF) QRectF_PTR() *QRectF {
	return ptr
}

func (ptr *QRectF) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRectF) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRectF(ptr QRectF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRectF_PTR().Pointer()
	}
	return nil
}

func NewQRectFFromPointer(ptr unsafe.Pointer) (n *QRectF) {
	n = new(QRectF)
	n.SetPointer(ptr)
	return
}
func (ptr *QRectF) DestroyQRectF() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQRectF() *QRectF {
	tmpValue := NewQRectFFromPointer(C.QRectF_NewQRectF())
	qt.SetFinalizer(tmpValue, (*QRectF).DestroyQRectF)
	return tmpValue
}

func NewQRectF2(topLeft QPointF_ITF, size QSizeF_ITF) *QRectF {
	tmpValue := NewQRectFFromPointer(C.QRectF_NewQRectF2(PointerFromQPointF(topLeft), PointerFromQSizeF(size)))
	qt.SetFinalizer(tmpValue, (*QRectF).DestroyQRectF)
	return tmpValue
}

func NewQRectF3(topLeft QPointF_ITF, bottomRight QPointF_ITF) *QRectF {
	tmpValue := NewQRectFFromPointer(C.QRectF_NewQRectF3(PointerFromQPointF(topLeft), PointerFromQPointF(bottomRight)))
	qt.SetFinalizer(tmpValue, (*QRectF).DestroyQRectF)
	return tmpValue
}

func NewQRectF4(x float64, y float64, width float64, height float64) *QRectF {
	tmpValue := NewQRectFFromPointer(C.QRectF_NewQRectF4(C.double(x), C.double(y), C.double(width), C.double(height)))
	qt.SetFinalizer(tmpValue, (*QRectF).DestroyQRectF)
	return tmpValue
}

func NewQRectF5(rectangle QRect_ITF) *QRectF {
	tmpValue := NewQRectFFromPointer(C.QRectF_NewQRectF5(PointerFromQRect(rectangle)))
	qt.SetFinalizer(tmpValue, (*QRectF).DestroyQRectF)
	return tmpValue
}

func (ptr *QRectF) Center() *QPointF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPointFFromPointer(C.QRectF_Center(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QRectF) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRectF_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRectF_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRectF) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRectF_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRectF) Left() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRectF_Left(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) Right() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRectF_Right(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) SetSize(size QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.QRectF_SetSize(ptr.Pointer(), PointerFromQSizeF(size))
	}
}

func (ptr *QRectF) SetTop(y float64) {
	if ptr.Pointer() != nil {
		C.QRectF_SetTop(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QRectF) Size() *QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := NewQSizeFFromPointer(C.QRectF_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QRectF) Top() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRectF_Top(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRectF_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRectF_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRectF_Y(ptr.Pointer()))
	}
	return 0
}

type QRegExp struct {
	ptr unsafe.Pointer
}

type QRegExp_ITF interface {
	QRegExp_PTR() *QRegExp
}

func (ptr *QRegExp) QRegExp_PTR() *QRegExp {
	return ptr
}

func (ptr *QRegExp) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRegExp) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRegExp(ptr QRegExp_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegExp_PTR().Pointer()
	}
	return nil
}

func NewQRegExpFromPointer(ptr unsafe.Pointer) (n *QRegExp) {
	n = new(QRegExp)
	n.SetPointer(ptr)
	return
}

// QRegExp::CaretMode
//
//go:generate stringer -type=QRegExp__CaretMode
type QRegExp__CaretMode int64

const (
	QRegExp__CaretAtZero    QRegExp__CaretMode = QRegExp__CaretMode(0)
	QRegExp__CaretAtOffset  QRegExp__CaretMode = QRegExp__CaretMode(1)
	QRegExp__CaretWontMatch QRegExp__CaretMode = QRegExp__CaretMode(2)
)

// QRegExp::PatternSyntax
//
//go:generate stringer -type=QRegExp__PatternSyntax
type QRegExp__PatternSyntax int64

const (
	QRegExp__RegExp         QRegExp__PatternSyntax = QRegExp__PatternSyntax(0)
	QRegExp__Wildcard       QRegExp__PatternSyntax = QRegExp__PatternSyntax(1)
	QRegExp__FixedString    QRegExp__PatternSyntax = QRegExp__PatternSyntax(2)
	QRegExp__RegExp2        QRegExp__PatternSyntax = QRegExp__PatternSyntax(3)
	QRegExp__WildcardUnix   QRegExp__PatternSyntax = QRegExp__PatternSyntax(4)
	QRegExp__W3CXmlSchema11 QRegExp__PatternSyntax = QRegExp__PatternSyntax(5)
)

func NewQRegExp() *QRegExp {
	tmpValue := NewQRegExpFromPointer(C.QRegExp_NewQRegExp())
	qt.SetFinalizer(tmpValue, (*QRegExp).DestroyQRegExp)
	return tmpValue
}

func NewQRegExp2(pattern string, cs Qt__CaseSensitivity, syntax QRegExp__PatternSyntax) *QRegExp {
	var patternC *C.char
	if pattern != "" {
		patternC = C.CString(pattern)
		defer C.free(unsafe.Pointer(patternC))
	}
	tmpValue := NewQRegExpFromPointer(C.QRegExp_NewQRegExp2(C.struct_QtCore_PackedString{data: patternC, len: C.longlong(len(pattern))}, C.longlong(cs), C.longlong(syntax)))
	qt.SetFinalizer(tmpValue, (*QRegExp).DestroyQRegExp)
	return tmpValue
}

func NewQRegExp3(rx QRegExp_ITF) *QRegExp {
	tmpValue := NewQRegExpFromPointer(C.QRegExp_NewQRegExp3(PointerFromQRegExp(rx)))
	qt.SetFinalizer(tmpValue, (*QRegExp).DestroyQRegExp)
	return tmpValue
}

func (ptr *QRegExp) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRegExp_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRegExp) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRegExp_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRegExp) Pattern() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QRegExp_Pattern(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRegExp) Pos(nth int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRegExp_Pos(ptr.Pointer(), C.int(int32(nth)))))
	}
	return 0
}

func (ptr *QRegExp) DestroyQRegExp() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRegExp_DestroyQRegExp(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QRegularExpression struct {
	ptr unsafe.Pointer
}

type QRegularExpression_ITF interface {
	QRegularExpression_PTR() *QRegularExpression
}

func (ptr *QRegularExpression) QRegularExpression_PTR() *QRegularExpression {
	return ptr
}

func (ptr *QRegularExpression) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRegularExpression) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRegularExpression(ptr QRegularExpression_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegularExpression_PTR().Pointer()
	}
	return nil
}

func NewQRegularExpressionFromPointer(ptr unsafe.Pointer) (n *QRegularExpression) {
	n = new(QRegularExpression)
	n.SetPointer(ptr)
	return
}

// QRegularExpression::MatchOption
//
//go:generate stringer -type=QRegularExpression__MatchOption
type QRegularExpression__MatchOption int64

const (
	QRegularExpression__NoMatchOption                     QRegularExpression__MatchOption = QRegularExpression__MatchOption(0x0000)
	QRegularExpression__AnchoredMatchOption               QRegularExpression__MatchOption = QRegularExpression__MatchOption(0x0001)
	QRegularExpression__DontCheckSubjectStringMatchOption QRegularExpression__MatchOption = QRegularExpression__MatchOption(0x0002)
)

// QRegularExpression::MatchType
//
//go:generate stringer -type=QRegularExpression__MatchType
type QRegularExpression__MatchType int64

const (
	QRegularExpression__NormalMatch                QRegularExpression__MatchType = QRegularExpression__MatchType(0)
	QRegularExpression__PartialPreferCompleteMatch QRegularExpression__MatchType = QRegularExpression__MatchType(1)
	QRegularExpression__PartialPreferFirstMatch    QRegularExpression__MatchType = QRegularExpression__MatchType(2)
	QRegularExpression__NoMatch                    QRegularExpression__MatchType = QRegularExpression__MatchType(3)
)

// QRegularExpression::PatternOption
//
//go:generate stringer -type=QRegularExpression__PatternOption
type QRegularExpression__PatternOption int64

const (
	QRegularExpression__NoPatternOption                 QRegularExpression__PatternOption = QRegularExpression__PatternOption(0x0000)
	QRegularExpression__CaseInsensitiveOption           QRegularExpression__PatternOption = QRegularExpression__PatternOption(0x0001)
	QRegularExpression__DotMatchesEverythingOption      QRegularExpression__PatternOption = QRegularExpression__PatternOption(0x0002)
	QRegularExpression__MultilineOption                 QRegularExpression__PatternOption = QRegularExpression__PatternOption(0x0004)
	QRegularExpression__ExtendedPatternSyntaxOption     QRegularExpression__PatternOption = QRegularExpression__PatternOption(0x0008)
	QRegularExpression__InvertedGreedinessOption        QRegularExpression__PatternOption = QRegularExpression__PatternOption(0x0010)
	QRegularExpression__DontCaptureOption               QRegularExpression__PatternOption = QRegularExpression__PatternOption(0x0020)
	QRegularExpression__UseUnicodePropertiesOption      QRegularExpression__PatternOption = QRegularExpression__PatternOption(0x0040)
	QRegularExpression__OptimizeOnFirstUsageOption      QRegularExpression__PatternOption = QRegularExpression__PatternOption(0x0080)
	QRegularExpression__DontAutomaticallyOptimizeOption QRegularExpression__PatternOption = QRegularExpression__PatternOption(0x0100)
)

func NewQRegularExpression() *QRegularExpression {
	tmpValue := NewQRegularExpressionFromPointer(C.QRegularExpression_NewQRegularExpression())
	qt.SetFinalizer(tmpValue, (*QRegularExpression).DestroyQRegularExpression)
	return tmpValue
}

func NewQRegularExpression2(pattern string, options QRegularExpression__PatternOption) *QRegularExpression {
	var patternC *C.char
	if pattern != "" {
		patternC = C.CString(pattern)
		defer C.free(unsafe.Pointer(patternC))
	}
	tmpValue := NewQRegularExpressionFromPointer(C.QRegularExpression_NewQRegularExpression2(C.struct_QtCore_PackedString{data: patternC, len: C.longlong(len(pattern))}, C.longlong(options)))
	qt.SetFinalizer(tmpValue, (*QRegularExpression).DestroyQRegularExpression)
	return tmpValue
}

func NewQRegularExpression3(re QRegularExpression_ITF) *QRegularExpression {
	tmpValue := NewQRegularExpressionFromPointer(C.QRegularExpression_NewQRegularExpression3(PointerFromQRegularExpression(re)))
	qt.SetFinalizer(tmpValue, (*QRegularExpression).DestroyQRegularExpression)
	return tmpValue
}

func (ptr *QRegularExpression) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRegularExpression_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRegularExpression) Match(subject string, offset int, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatch {
	if ptr.Pointer() != nil {
		var subjectC *C.char
		if subject != "" {
			subjectC = C.CString(subject)
			defer C.free(unsafe.Pointer(subjectC))
		}
		tmpValue := NewQRegularExpressionMatchFromPointer(C.QRegularExpression_Match(ptr.Pointer(), C.struct_QtCore_PackedString{data: subjectC, len: C.longlong(len(subject))}, C.int(int32(offset)), C.longlong(matchType), C.longlong(matchOptions)))
		qt.SetFinalizer(tmpValue, (*QRegularExpressionMatch).DestroyQRegularExpressionMatch)
		return tmpValue
	}
	return nil
}

func (ptr *QRegularExpression) Match2(subjectRef QStringRef_ITF, offset int, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatch {
	if ptr.Pointer() != nil {
		tmpValue := NewQRegularExpressionMatchFromPointer(C.QRegularExpression_Match2(ptr.Pointer(), PointerFromQStringRef(subjectRef), C.int(int32(offset)), C.longlong(matchType), C.longlong(matchOptions)))
		qt.SetFinalizer(tmpValue, (*QRegularExpressionMatch).DestroyQRegularExpressionMatch)
		return tmpValue
	}
	return nil
}

func (ptr *QRegularExpression) Pattern() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QRegularExpression_Pattern(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRegularExpression) DestroyQRegularExpression() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRegularExpression_DestroyQRegularExpression(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QRegularExpressionMatch struct {
	ptr unsafe.Pointer
}

type QRegularExpressionMatch_ITF interface {
	QRegularExpressionMatch_PTR() *QRegularExpressionMatch
}

func (ptr *QRegularExpressionMatch) QRegularExpressionMatch_PTR() *QRegularExpressionMatch {
	return ptr
}

func (ptr *QRegularExpressionMatch) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRegularExpressionMatch) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRegularExpressionMatch(ptr QRegularExpressionMatch_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegularExpressionMatch_PTR().Pointer()
	}
	return nil
}

func NewQRegularExpressionMatchFromPointer(ptr unsafe.Pointer) (n *QRegularExpressionMatch) {
	n = new(QRegularExpressionMatch)
	n.SetPointer(ptr)
	return
}
func NewQRegularExpressionMatch() *QRegularExpressionMatch {
	tmpValue := NewQRegularExpressionMatchFromPointer(C.QRegularExpressionMatch_NewQRegularExpressionMatch())
	qt.SetFinalizer(tmpValue, (*QRegularExpressionMatch).DestroyQRegularExpressionMatch)
	return tmpValue
}

func NewQRegularExpressionMatch2(match QRegularExpressionMatch_ITF) *QRegularExpressionMatch {
	tmpValue := NewQRegularExpressionMatchFromPointer(C.QRegularExpressionMatch_NewQRegularExpressionMatch2(PointerFromQRegularExpressionMatch(match)))
	qt.SetFinalizer(tmpValue, (*QRegularExpressionMatch).DestroyQRegularExpressionMatch)
	return tmpValue
}

func (ptr *QRegularExpressionMatch) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRegularExpressionMatch_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRegularExpressionMatch) DestroyQRegularExpressionMatch() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRegularExpressionMatch_DestroyQRegularExpressionMatch(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QResource::Compression
//
//go:generate stringer -type=QResource__Compression
type QResource__Compression int64

const (
	QResource__NoCompression   QResource__Compression = QResource__Compression(0)
	QResource__ZlibCompression QResource__Compression = QResource__Compression(1)
	QResource__ZstdCompression QResource__Compression = QResource__Compression(2)
)

// QSettings::Format
//
//go:generate stringer -type=QSettings__Format
type QSettings__Format int64

const (
	QSettings__NativeFormat     QSettings__Format = QSettings__Format(0)
	QSettings__IniFormat        QSettings__Format = QSettings__Format(1)
	QSettings__Registry32Format QSettings__Format = QSettings__Format(2)
	QSettings__Registry64Format QSettings__Format = QSettings__Format(3)
	QSettings__InvalidFormat    QSettings__Format = QSettings__Format(16)
	QSettings__CustomFormat1    QSettings__Format = QSettings__Format(17)
	QSettings__CustomFormat2    QSettings__Format = QSettings__Format(18)
	QSettings__CustomFormat3    QSettings__Format = QSettings__Format(19)
	QSettings__CustomFormat4    QSettings__Format = QSettings__Format(20)
	QSettings__CustomFormat5    QSettings__Format = QSettings__Format(21)
	QSettings__CustomFormat6    QSettings__Format = QSettings__Format(22)
	QSettings__CustomFormat7    QSettings__Format = QSettings__Format(23)
	QSettings__CustomFormat8    QSettings__Format = QSettings__Format(24)
	QSettings__CustomFormat9    QSettings__Format = QSettings__Format(25)
	QSettings__CustomFormat10   QSettings__Format = QSettings__Format(26)
	QSettings__CustomFormat11   QSettings__Format = QSettings__Format(27)
	QSettings__CustomFormat12   QSettings__Format = QSettings__Format(28)
	QSettings__CustomFormat13   QSettings__Format = QSettings__Format(29)
	QSettings__CustomFormat14   QSettings__Format = QSettings__Format(30)
	QSettings__CustomFormat15   QSettings__Format = QSettings__Format(31)
	QSettings__CustomFormat16   QSettings__Format = QSettings__Format(32)
)

// QSettings::Scope
//
//go:generate stringer -type=QSettings__Scope
type QSettings__Scope int64

const (
	QSettings__UserScope   QSettings__Scope = QSettings__Scope(0)
	QSettings__SystemScope QSettings__Scope = QSettings__Scope(1)
)

// QSettings::Status
//
//go:generate stringer -type=QSettings__Status
type QSettings__Status int64

const (
	QSettings__NoError     QSettings__Status = QSettings__Status(0)
	QSettings__AccessError QSettings__Status = QSettings__Status(1)
	QSettings__FormatError QSettings__Status = QSettings__Status(2)
)

// QSharedMemory::AccessMode
//
//go:generate stringer -type=QSharedMemory__AccessMode
type QSharedMemory__AccessMode int64

const (
	QSharedMemory__ReadOnly  QSharedMemory__AccessMode = QSharedMemory__AccessMode(0)
	QSharedMemory__ReadWrite QSharedMemory__AccessMode = QSharedMemory__AccessMode(1)
)

// QSharedMemory::SharedMemoryError
//
//go:generate stringer -type=QSharedMemory__SharedMemoryError
type QSharedMemory__SharedMemoryError int64

const (
	QSharedMemory__NoError          QSharedMemory__SharedMemoryError = QSharedMemory__SharedMemoryError(0)
	QSharedMemory__PermissionDenied QSharedMemory__SharedMemoryError = QSharedMemory__SharedMemoryError(1)
	QSharedMemory__InvalidSize      QSharedMemory__SharedMemoryError = QSharedMemory__SharedMemoryError(2)
	QSharedMemory__KeyError         QSharedMemory__SharedMemoryError = QSharedMemory__SharedMemoryError(3)
	QSharedMemory__AlreadyExists    QSharedMemory__SharedMemoryError = QSharedMemory__SharedMemoryError(4)
	QSharedMemory__NotFound         QSharedMemory__SharedMemoryError = QSharedMemory__SharedMemoryError(5)
	QSharedMemory__LockError        QSharedMemory__SharedMemoryError = QSharedMemory__SharedMemoryError(6)
	QSharedMemory__OutOfResources   QSharedMemory__SharedMemoryError = QSharedMemory__SharedMemoryError(7)
	QSharedMemory__UnknownError     QSharedMemory__SharedMemoryError = QSharedMemory__SharedMemoryError(8)
)

type QSize struct {
	ptr unsafe.Pointer
}

type QSize_ITF interface {
	QSize_PTR() *QSize
}

func (ptr *QSize) QSize_PTR() *QSize {
	return ptr
}

func (ptr *QSize) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSize) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSize(ptr QSize_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSize_PTR().Pointer()
	}
	return nil
}

func NewQSizeFromPointer(ptr unsafe.Pointer) (n *QSize) {
	n = new(QSize)
	n.SetPointer(ptr)
	return
}
func (ptr *QSize) DestroyQSize() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQSize() *QSize {
	tmpValue := NewQSizeFromPointer(C.QSize_NewQSize())
	qt.SetFinalizer(tmpValue, (*QSize).DestroyQSize)
	return tmpValue
}

func NewQSize2(width int, height int) *QSize {
	tmpValue := NewQSizeFromPointer(C.QSize_NewQSize2(C.int(int32(width)), C.int(int32(height))))
	qt.SetFinalizer(tmpValue, (*QSize).DestroyQSize)
	return tmpValue
}

func (ptr *QSize) Height() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSize_Height(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSize) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSize_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSize) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSize_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSize) Scale(width int, height int, mode Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QSize_Scale(ptr.Pointer(), C.int(int32(width)), C.int(int32(height)), C.longlong(mode))
	}
}

func (ptr *QSize) Scale2(size QSize_ITF, mode Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QSize_Scale2(ptr.Pointer(), PointerFromQSize(size), C.longlong(mode))
	}
}

func (ptr *QSize) Scaled(width int, height int, mode Qt__AspectRatioMode) *QSize {
	if ptr.Pointer() != nil {
		tmpValue := NewQSizeFromPointer(C.QSize_Scaled(ptr.Pointer(), C.int(int32(width)), C.int(int32(height)), C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSize) Scaled2(s QSize_ITF, mode Qt__AspectRatioMode) *QSize {
	if ptr.Pointer() != nil {
		tmpValue := NewQSizeFromPointer(C.QSize_Scaled2(ptr.Pointer(), PointerFromQSize(s), C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSize) Width() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSize_Width(ptr.Pointer())))
	}
	return 0
}

type QSizeF struct {
	ptr unsafe.Pointer
}

type QSizeF_ITF interface {
	QSizeF_PTR() *QSizeF
}

func (ptr *QSizeF) QSizeF_PTR() *QSizeF {
	return ptr
}

func (ptr *QSizeF) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSizeF) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSizeF(ptr QSizeF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSizeF_PTR().Pointer()
	}
	return nil
}

func NewQSizeFFromPointer(ptr unsafe.Pointer) (n *QSizeF) {
	n = new(QSizeF)
	n.SetPointer(ptr)
	return
}
func (ptr *QSizeF) DestroyQSizeF() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQSizeF() *QSizeF {
	tmpValue := NewQSizeFFromPointer(C.QSizeF_NewQSizeF())
	qt.SetFinalizer(tmpValue, (*QSizeF).DestroyQSizeF)
	return tmpValue
}

func NewQSizeF2(size QSize_ITF) *QSizeF {
	tmpValue := NewQSizeFFromPointer(C.QSizeF_NewQSizeF2(PointerFromQSize(size)))
	qt.SetFinalizer(tmpValue, (*QSizeF).DestroyQSizeF)
	return tmpValue
}

func NewQSizeF3(width float64, height float64) *QSizeF {
	tmpValue := NewQSizeFFromPointer(C.QSizeF_NewQSizeF3(C.double(width), C.double(height)))
	qt.SetFinalizer(tmpValue, (*QSizeF).DestroyQSizeF)
	return tmpValue
}

func (ptr *QSizeF) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QSizeF_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizeF) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSizeF_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSizeF) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSizeF_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSizeF) Scale(width float64, height float64, mode Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QSizeF_Scale(ptr.Pointer(), C.double(width), C.double(height), C.longlong(mode))
	}
}

func (ptr *QSizeF) Scale2(size QSizeF_ITF, mode Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QSizeF_Scale2(ptr.Pointer(), PointerFromQSizeF(size), C.longlong(mode))
	}
}

func (ptr *QSizeF) Scaled(width float64, height float64, mode Qt__AspectRatioMode) *QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := NewQSizeFFromPointer(C.QSizeF_Scaled(ptr.Pointer(), C.double(width), C.double(height), C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QSizeF) Scaled2(s QSizeF_ITF, mode Qt__AspectRatioMode) *QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := NewQSizeFFromPointer(C.QSizeF_Scaled2(ptr.Pointer(), PointerFromQSizeF(s), C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QSizeF) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QSizeF_Width(ptr.Pointer()))
	}
	return 0
}

// QSocketNotifier::Type
//
//go:generate stringer -type=QSocketNotifier__Type
type QSocketNotifier__Type int64

const (
	QSocketNotifier__Read      QSocketNotifier__Type = QSocketNotifier__Type(0)
	QSocketNotifier__Write     QSocketNotifier__Type = QSocketNotifier__Type(1)
	QSocketNotifier__Exception QSocketNotifier__Type = QSocketNotifier__Type(2)
)

// QStandardPaths::LocateOption
//
//go:generate stringer -type=QStandardPaths__LocateOption
type QStandardPaths__LocateOption int64

const (
	QStandardPaths__LocateFile      QStandardPaths__LocateOption = QStandardPaths__LocateOption(0x0)
	QStandardPaths__LocateDirectory QStandardPaths__LocateOption = QStandardPaths__LocateOption(0x1)
)

// QStandardPaths::StandardLocation
//
//go:generate stringer -type=QStandardPaths__StandardLocation
type QStandardPaths__StandardLocation int64

const (
	QStandardPaths__DesktopLocation       QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(0)
	QStandardPaths__DocumentsLocation     QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(1)
	QStandardPaths__FontsLocation         QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(2)
	QStandardPaths__ApplicationsLocation  QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(3)
	QStandardPaths__MusicLocation         QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(4)
	QStandardPaths__MoviesLocation        QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(5)
	QStandardPaths__PicturesLocation      QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(6)
	QStandardPaths__TempLocation          QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(7)
	QStandardPaths__HomeLocation          QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(8)
	QStandardPaths__DataLocation          QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(9)
	QStandardPaths__CacheLocation         QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(10)
	QStandardPaths__GenericDataLocation   QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(11)
	QStandardPaths__RuntimeLocation       QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(12)
	QStandardPaths__ConfigLocation        QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(13)
	QStandardPaths__DownloadLocation      QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(14)
	QStandardPaths__GenericCacheLocation  QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(15)
	QStandardPaths__GenericConfigLocation QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(16)
	QStandardPaths__AppDataLocation       QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(17)
	QStandardPaths__AppConfigLocation     QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(18)
	QStandardPaths__AppLocalDataLocation  QStandardPaths__StandardLocation = QStandardPaths__StandardLocation(QStandardPaths__DataLocation)
)

// QState::ChildMode
//
//go:generate stringer -type=QState__ChildMode
type QState__ChildMode int64

const (
	QState__ExclusiveStates QState__ChildMode = QState__ChildMode(0)
	QState__ParallelStates  QState__ChildMode = QState__ChildMode(1)
)

// QState::RestorePolicy
//
//go:generate stringer -type=QState__RestorePolicy
type QState__RestorePolicy int64

const (
	QState__DontRestoreProperties QState__RestorePolicy = QState__RestorePolicy(0)
	QState__RestoreProperties     QState__RestorePolicy = QState__RestorePolicy(1)
)

// QStateMachine::Error
//
//go:generate stringer -type=QStateMachine__Error
type QStateMachine__Error int64

const (
	QStateMachine__NoError                                 QStateMachine__Error = QStateMachine__Error(0)
	QStateMachine__NoInitialStateError                     QStateMachine__Error = QStateMachine__Error(1)
	QStateMachine__NoDefaultStateInHistoryStateError       QStateMachine__Error = QStateMachine__Error(2)
	QStateMachine__NoCommonAncestorForTransitionError      QStateMachine__Error = QStateMachine__Error(3)
	QStateMachine__StateMachineChildModeSetToParallelError QStateMachine__Error = QStateMachine__Error(4)
)

// QStateMachine::EventPriority
//
//go:generate stringer -type=QStateMachine__EventPriority
type QStateMachine__EventPriority int64

const (
	QStateMachine__NormalPriority QStateMachine__EventPriority = QStateMachine__EventPriority(0)
	QStateMachine__HighPriority   QStateMachine__EventPriority = QStateMachine__EventPriority(1)
)

type QString struct {
	ptr unsafe.Pointer
}

type QString_ITF interface {
	QString_PTR() *QString
}

func (ptr *QString) QString_PTR() *QString {
	return ptr
}

func (ptr *QString) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QString) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQString(ptr QString_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QString_PTR().Pointer()
	}
	return nil
}

func NewQStringFromPointer(ptr unsafe.Pointer) (n *QString) {
	n = new(QString)
	n.SetPointer(ptr)
	return
}

// QString::NormalizationForm
//
//go:generate stringer -type=QString__NormalizationForm
type QString__NormalizationForm int64

const (
	QString__NormalizationForm_D  QString__NormalizationForm = QString__NormalizationForm(0)
	QString__NormalizationForm_C  QString__NormalizationForm = QString__NormalizationForm(1)
	QString__NormalizationForm_KD QString__NormalizationForm = QString__NormalizationForm(2)
	QString__NormalizationForm_KC QString__NormalizationForm = QString__NormalizationForm(3)
)

// QString::SectionFlag
//
//go:generate stringer -type=QString__SectionFlag
type QString__SectionFlag int64

const (
	QString__SectionDefault             QString__SectionFlag = QString__SectionFlag(0x00)
	QString__SectionSkipEmpty           QString__SectionFlag = QString__SectionFlag(0x01)
	QString__SectionIncludeLeadingSep   QString__SectionFlag = QString__SectionFlag(0x02)
	QString__SectionIncludeTrailingSep  QString__SectionFlag = QString__SectionFlag(0x04)
	QString__SectionCaseInsensitiveSeps QString__SectionFlag = QString__SectionFlag(0x08)
)

type QStringList struct {
	ptr unsafe.Pointer
}

type QStringList_ITF interface {
	QStringList_PTR() *QStringList
}

func (ptr *QStringList) QStringList_PTR() *QStringList {
	return ptr
}

func (ptr *QStringList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QStringList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQStringList(ptr QStringList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringList_PTR().Pointer()
	}
	return nil
}

func NewQStringListFromPointer(ptr unsafe.Pointer) (n *QStringList) {
	n = new(QStringList)
	n.SetPointer(ptr)
	return
}
func (ptr *QStringList) DestroyQStringList() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QStringRef struct {
	ptr unsafe.Pointer
}

type QStringRef_ITF interface {
	QStringRef_PTR() *QStringRef
}

func (ptr *QStringRef) QStringRef_PTR() *QStringRef {
	return ptr
}

func (ptr *QStringRef) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QStringRef) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQStringRef(ptr QStringRef_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringRef_PTR().Pointer()
	}
	return nil
}

func NewQStringRefFromPointer(ptr unsafe.Pointer) (n *QStringRef) {
	n = new(QStringRef)
	n.SetPointer(ptr)
	return
}
func NewQStringRef() *QStringRef {
	tmpValue := NewQStringRefFromPointer(C.QStringRef_NewQStringRef())
	qt.SetFinalizer(tmpValue, (*QStringRef).DestroyQStringRef)
	return tmpValue
}

func NewQStringRef2(stri string, position int, length int) *QStringRef {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQStringRefFromPointer(C.QStringRef_NewQStringRef2(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.int(int32(position)), C.int(int32(length))))
	qt.SetFinalizer(tmpValue, (*QStringRef).DestroyQStringRef)
	return tmpValue
}

func NewQStringRef3(stri string) *QStringRef {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQStringRefFromPointer(C.QStringRef_NewQStringRef3(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}))
	qt.SetFinalizer(tmpValue, (*QStringRef).DestroyQStringRef)
	return tmpValue
}

func NewQStringRef4(other QStringRef_ITF) *QStringRef {
	tmpValue := NewQStringRefFromPointer(C.QStringRef_NewQStringRef4(PointerFromQStringRef(other)))
	qt.SetFinalizer(tmpValue, (*QStringRef).DestroyQStringRef)
	return tmpValue
}

func (ptr *QStringRef) At(position int) *QChar {
	if ptr.Pointer() != nil {
		tmpValue := NewQCharFromPointer(C.QStringRef_At(ptr.Pointer(), C.int(int32(position))))
		qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
		return tmpValue
	}
	return nil
}

func (ptr *QStringRef) Back() *QChar {
	if ptr.Pointer() != nil {
		tmpValue := NewQCharFromPointer(C.QStringRef_Back(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
		return tmpValue
	}
	return nil
}

func (ptr *QStringRef) Clear() {
	if ptr.Pointer() != nil {
		C.QStringRef_Clear(ptr.Pointer())
	}
}

func (ptr *QStringRef) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringRef_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStringRef) Count2(str string, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		return int(int32(C.QStringRef_Count2(ptr.Pointer(), C.struct_QtCore_PackedString{data: strC, len: C.longlong(len(str))}, C.longlong(cs))))
	}
	return 0
}

func (ptr *QStringRef) Count3(ch QChar_ITF, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringRef_Count3(ptr.Pointer(), PointerFromQChar(ch), C.longlong(cs))))
	}
	return 0
}

func (ptr *QStringRef) Count4(str QStringRef_ITF, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringRef_Count4(ptr.Pointer(), PointerFromQStringRef(str), C.longlong(cs))))
	}
	return 0
}

func (ptr *QStringRef) Data() *QChar {
	if ptr.Pointer() != nil {
		return NewQCharFromPointer(C.QStringRef_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStringRef) IndexOf(str string, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		return int(int32(C.QStringRef_IndexOf(ptr.Pointer(), C.struct_QtCore_PackedString{data: strC, len: C.longlong(len(str))}, C.int(int32(from)), C.longlong(cs))))
	}
	return 0
}

func (ptr *QStringRef) IndexOf2(str QStringRef_ITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringRef_IndexOf2(ptr.Pointer(), PointerFromQStringRef(str), C.int(int32(from)), C.longlong(cs))))
	}
	return 0
}

func (ptr *QStringRef) IndexOf3(str QStringView_ITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringRef_IndexOf3(ptr.Pointer(), PointerFromQStringView(str), C.int(int32(from)), C.longlong(cs))))
	}
	return 0
}

func (ptr *QStringRef) IndexOf4(ch QChar_ITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringRef_IndexOf4(ptr.Pointer(), PointerFromQChar(ch), C.int(int32(from)), C.longlong(cs))))
	}
	return 0
}

func (ptr *QStringRef) IndexOf5(str QLatin1String_ITF, from int, cs Qt__CaseSensitivity) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringRef_IndexOf5(ptr.Pointer(), PointerFromQLatin1String(str), C.int(int32(from)), C.longlong(cs))))
	}
	return 0
}

func (ptr *QStringRef) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QStringRef_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStringRef) Left(n int) *QStringRef {
	if ptr.Pointer() != nil {
		tmpValue := NewQStringRefFromPointer(C.QStringRef_Left(ptr.Pointer(), C.int(int32(n))))
		qt.SetFinalizer(tmpValue, (*QStringRef).DestroyQStringRef)
		return tmpValue
	}
	return nil
}

func (ptr *QStringRef) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringRef_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStringRef) Position() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringRef_Position(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStringRef) Right(n int) *QStringRef {
	if ptr.Pointer() != nil {
		tmpValue := NewQStringRefFromPointer(C.QStringRef_Right(ptr.Pointer(), C.int(int32(n))))
		qt.SetFinalizer(tmpValue, (*QStringRef).DestroyQStringRef)
		return tmpValue
	}
	return nil
}

func (ptr *QStringRef) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringRef_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStringRef) Split(sep string, behavior Qt__SplitBehaviorFlags, cs Qt__CaseSensitivity) []*QStringRef {
	if ptr.Pointer() != nil {
		var sepC *C.char
		if sep != "" {
			sepC = C.CString(sep)
			defer C.free(unsafe.Pointer(sepC))
		}
		return func(l C.struct_QtCore_PackedList) []*QStringRef {
			out := make([]*QStringRef, int(l.len))
			tmpList := NewQStringRefFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__split_atList(i)
			}
			return out
		}(C.QStringRef_Split(ptr.Pointer(), C.struct_QtCore_PackedString{data: sepC, len: C.longlong(len(sep))}, C.longlong(behavior), C.longlong(cs)))
	}
	return make([]*QStringRef, 0)
}

func (ptr *QStringRef) Split4(sep QChar_ITF, behavior Qt__SplitBehaviorFlags, cs Qt__CaseSensitivity) []*QStringRef {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []*QStringRef {
			out := make([]*QStringRef, int(l.len))
			tmpList := NewQStringRefFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__split_atList4(i)
			}
			return out
		}(C.QStringRef_Split4(ptr.Pointer(), PointerFromQChar(sep), C.longlong(behavior), C.longlong(cs)))
	}
	return make([]*QStringRef, 0)
}

func (ptr *QStringRef) String() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QStringRef_String(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStringRef) ToInt(ok *bool, base int) int {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return int(int32(C.QStringRef_ToInt(ptr.Pointer(), &okC, C.int(int32(base)))))
	}
	return 0
}

func (ptr *QStringRef) DestroyQStringRef() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStringRef_DestroyQStringRef(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStringRef) __split_atList(i int) *QStringRef {
	if ptr.Pointer() != nil {
		tmpValue := NewQStringRefFromPointer(C.QStringRef___split_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QStringRef).DestroyQStringRef)
		return tmpValue
	}
	return nil
}

func (ptr *QStringRef) __split_setList(i QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QStringRef___split_setList(ptr.Pointer(), PointerFromQStringRef(i))
	}
}

func (ptr *QStringRef) __split_newList() unsafe.Pointer {
	return C.QStringRef___split_newList(ptr.Pointer())
}

func (ptr *QStringRef) __split_atList2(i int) *QStringRef {
	if ptr.Pointer() != nil {
		tmpValue := NewQStringRefFromPointer(C.QStringRef___split_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QStringRef).DestroyQStringRef)
		return tmpValue
	}
	return nil
}

func (ptr *QStringRef) __split_setList2(i QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QStringRef___split_setList2(ptr.Pointer(), PointerFromQStringRef(i))
	}
}

func (ptr *QStringRef) __split_newList2() unsafe.Pointer {
	return C.QStringRef___split_newList2(ptr.Pointer())
}

func (ptr *QStringRef) __split_atList4(i int) *QStringRef {
	if ptr.Pointer() != nil {
		tmpValue := NewQStringRefFromPointer(C.QStringRef___split_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QStringRef).DestroyQStringRef)
		return tmpValue
	}
	return nil
}

func (ptr *QStringRef) __split_setList4(i QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QStringRef___split_setList4(ptr.Pointer(), PointerFromQStringRef(i))
	}
}

func (ptr *QStringRef) __split_newList4() unsafe.Pointer {
	return C.QStringRef___split_newList4(ptr.Pointer())
}

func (ptr *QStringRef) __toUcs4_atList(i int) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QStringRef___toUcs4_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QStringRef) __toUcs4_setList(i uint) {
	if ptr.Pointer() != nil {
		C.QStringRef___toUcs4_setList(ptr.Pointer(), C.uint(uint32(i)))
	}
}

func (ptr *QStringRef) __toUcs4_newList() unsafe.Pointer {
	return C.QStringRef___toUcs4_newList(ptr.Pointer())
}

type QStringView struct {
	ptr unsafe.Pointer
}

type QStringView_ITF interface {
	QStringView_PTR() *QStringView
}

func (ptr *QStringView) QStringView_PTR() *QStringView {
	return ptr
}

func (ptr *QStringView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QStringView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQStringView(ptr QStringView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringView_PTR().Pointer()
	}
	return nil
}

func NewQStringViewFromPointer(ptr unsafe.Pointer) (n *QStringView) {
	n = new(QStringView)
	n.SetPointer(ptr)
	return
}
func (ptr *QStringView) DestroyQStringView() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQStringView() *QStringView {
	tmpValue := NewQStringViewFromPointer(C.QStringView_NewQStringView())
	qt.SetFinalizer(tmpValue, (*QStringView).DestroyQStringView)
	return tmpValue
}

func NewQStringView7(str string) *QStringView {
	var strC *C.char
	if str != "" {
		strC = C.CString(str)
		defer C.free(unsafe.Pointer(strC))
	}
	tmpValue := NewQStringViewFromPointer(C.QStringView_NewQStringView7(C.struct_QtCore_PackedString{data: strC, len: C.longlong(len(str))}))
	qt.SetFinalizer(tmpValue, (*QStringView).DestroyQStringView)
	return tmpValue
}

func NewQStringView8(str QStringRef_ITF) *QStringView {
	tmpValue := NewQStringViewFromPointer(C.QStringView_NewQStringView8(PointerFromQStringRef(str)))
	qt.SetFinalizer(tmpValue, (*QStringView).DestroyQStringView)
	return tmpValue
}

func (ptr *QStringView) Back() *QChar {
	if ptr.Pointer() != nil {
		tmpValue := NewQCharFromPointer(C.QStringView_Back(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
		return tmpValue
	}
	return nil
}

func (ptr *QStringView) Empty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QStringView_Empty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStringView) First() *QChar {
	if ptr.Pointer() != nil {
		tmpValue := NewQCharFromPointer(C.QStringView_First(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
		return tmpValue
	}
	return nil
}

func (ptr *QStringView) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QStringView_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStringView) Last() *QChar {
	if ptr.Pointer() != nil {
		tmpValue := NewQCharFromPointer(C.QStringView_Last(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QChar).DestroyQChar)
		return tmpValue
	}
	return nil
}

func (ptr *QStringView) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStringView_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStringView) Split(sep QStringView_ITF, behavior Qt__SplitBehaviorFlags, cs Qt__CaseSensitivity) []*QStringView {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []*QStringView {
			out := make([]*QStringView, int(l.len))
			tmpList := NewQStringViewFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__split_atList(i)
			}
			return out
		}(C.QStringView_Split(ptr.Pointer(), PointerFromQStringView(sep), C.longlong(behavior), C.longlong(cs)))
	}
	return make([]*QStringView, 0)
}

func (ptr *QStringView) Split2(sep QChar_ITF, behavior Qt__SplitBehaviorFlags, cs Qt__CaseSensitivity) []*QStringView {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []*QStringView {
			out := make([]*QStringView, int(l.len))
			tmpList := NewQStringViewFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__split_atList2(i)
			}
			return out
		}(C.QStringView_Split2(ptr.Pointer(), PointerFromQChar(sep), C.longlong(behavior), C.longlong(cs)))
	}
	return make([]*QStringView, 0)
}

func (ptr *QStringView) Split3(sep QRegularExpression_ITF, behavior Qt__SplitBehaviorFlags) []*QStringView {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []*QStringView {
			out := make([]*QStringView, int(l.len))
			tmpList := NewQStringViewFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__split_atList3(i)
			}
			return out
		}(C.QStringView_Split3(ptr.Pointer(), PointerFromQRegularExpression(sep), C.longlong(behavior)))
	}
	return make([]*QStringView, 0)
}

func (ptr *QStringView) __convertToUcs4_atList(i int) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QStringView___convertToUcs4_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QStringView) __convertToUcs4_setList(i uint) {
	if ptr.Pointer() != nil {
		C.QStringView___convertToUcs4_setList(ptr.Pointer(), C.uint(uint32(i)))
	}
}

func (ptr *QStringView) __convertToUcs4_newList() unsafe.Pointer {
	return C.QStringView___convertToUcs4_newList(ptr.Pointer())
}

func (ptr *QStringView) __split_atList(i int) *QStringView {
	if ptr.Pointer() != nil {
		tmpValue := NewQStringViewFromPointer(C.QStringView___split_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QStringView).DestroyQStringView)
		return tmpValue
	}
	return nil
}

func (ptr *QStringView) __split_setList(i QStringView_ITF) {
	if ptr.Pointer() != nil {
		C.QStringView___split_setList(ptr.Pointer(), PointerFromQStringView(i))
	}
}

func (ptr *QStringView) __split_newList() unsafe.Pointer {
	return C.QStringView___split_newList(ptr.Pointer())
}

func (ptr *QStringView) __split_atList2(i int) *QStringView {
	if ptr.Pointer() != nil {
		tmpValue := NewQStringViewFromPointer(C.QStringView___split_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QStringView).DestroyQStringView)
		return tmpValue
	}
	return nil
}

func (ptr *QStringView) __split_setList2(i QStringView_ITF) {
	if ptr.Pointer() != nil {
		C.QStringView___split_setList2(ptr.Pointer(), PointerFromQStringView(i))
	}
}

func (ptr *QStringView) __split_newList2() unsafe.Pointer {
	return C.QStringView___split_newList2(ptr.Pointer())
}

func (ptr *QStringView) __split_atList3(i int) *QStringView {
	if ptr.Pointer() != nil {
		tmpValue := NewQStringViewFromPointer(C.QStringView___split_atList3(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QStringView).DestroyQStringView)
		return tmpValue
	}
	return nil
}

func (ptr *QStringView) __split_setList3(i QStringView_ITF) {
	if ptr.Pointer() != nil {
		C.QStringView___split_setList3(ptr.Pointer(), PointerFromQStringView(i))
	}
}

func (ptr *QStringView) __split_newList3() unsafe.Pointer {
	return C.QStringView___split_newList3(ptr.Pointer())
}

func (ptr *QStringView) __toUcs4_atList(i int) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QStringView___toUcs4_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QStringView) __toUcs4_setList(i uint) {
	if ptr.Pointer() != nil {
		C.QStringView___toUcs4_setList(ptr.Pointer(), C.uint(uint32(i)))
	}
}

func (ptr *QStringView) __toUcs4_newList() unsafe.Pointer {
	return C.QStringView___toUcs4_newList(ptr.Pointer())
}

// QSysInfo::Endian
//
//go:generate stringer -type=QSysInfo__Endian
type QSysInfo__Endian int64

const (
	QSysInfo__BigEndian    QSysInfo__Endian = QSysInfo__Endian(0)
	QSysInfo__LittleEndian QSysInfo__Endian = QSysInfo__Endian(1)
)

// QSysInfo::Sizes
//
//go:generate stringer -type=QSysInfo__Sizes
type QSysInfo__Sizes int64

var (
	QSysInfo__WordSize QSysInfo__Sizes = QSysInfo__Sizes(C.QSysInfo_WordSize_Type())
)

// QSystemSemaphore::AccessMode
//
//go:generate stringer -type=QSystemSemaphore__AccessMode
type QSystemSemaphore__AccessMode int64

const (
	QSystemSemaphore__Open   QSystemSemaphore__AccessMode = QSystemSemaphore__AccessMode(0)
	QSystemSemaphore__Create QSystemSemaphore__AccessMode = QSystemSemaphore__AccessMode(1)
)

// QSystemSemaphore::SystemSemaphoreError
//
//go:generate stringer -type=QSystemSemaphore__SystemSemaphoreError
type QSystemSemaphore__SystemSemaphoreError int64

const (
	QSystemSemaphore__NoError          QSystemSemaphore__SystemSemaphoreError = QSystemSemaphore__SystemSemaphoreError(0)
	QSystemSemaphore__PermissionDenied QSystemSemaphore__SystemSemaphoreError = QSystemSemaphore__SystemSemaphoreError(1)
	QSystemSemaphore__KeyError         QSystemSemaphore__SystemSemaphoreError = QSystemSemaphore__SystemSemaphoreError(2)
	QSystemSemaphore__AlreadyExists    QSystemSemaphore__SystemSemaphoreError = QSystemSemaphore__SystemSemaphoreError(3)
	QSystemSemaphore__NotFound         QSystemSemaphore__SystemSemaphoreError = QSystemSemaphore__SystemSemaphoreError(4)
	QSystemSemaphore__OutOfResources   QSystemSemaphore__SystemSemaphoreError = QSystemSemaphore__SystemSemaphoreError(5)
	QSystemSemaphore__UnknownError     QSystemSemaphore__SystemSemaphoreError = QSystemSemaphore__SystemSemaphoreError(6)
)

// QTextBoundaryFinder::BoundaryReason
//
//go:generate stringer -type=QTextBoundaryFinder__BoundaryReason
type QTextBoundaryFinder__BoundaryReason int64

const (
	QTextBoundaryFinder__NotAtBoundary    QTextBoundaryFinder__BoundaryReason = QTextBoundaryFinder__BoundaryReason(0)
	QTextBoundaryFinder__BreakOpportunity QTextBoundaryFinder__BoundaryReason = QTextBoundaryFinder__BoundaryReason(0x1f)
	QTextBoundaryFinder__StartOfItem      QTextBoundaryFinder__BoundaryReason = QTextBoundaryFinder__BoundaryReason(0x20)
	QTextBoundaryFinder__EndOfItem        QTextBoundaryFinder__BoundaryReason = QTextBoundaryFinder__BoundaryReason(0x40)
	QTextBoundaryFinder__MandatoryBreak   QTextBoundaryFinder__BoundaryReason = QTextBoundaryFinder__BoundaryReason(0x80)
	QTextBoundaryFinder__SoftHyphen       QTextBoundaryFinder__BoundaryReason = QTextBoundaryFinder__BoundaryReason(0x100)
)

// QTextBoundaryFinder::BoundaryType
//
//go:generate stringer -type=QTextBoundaryFinder__BoundaryType
type QTextBoundaryFinder__BoundaryType int64

const (
	QTextBoundaryFinder__Grapheme QTextBoundaryFinder__BoundaryType = QTextBoundaryFinder__BoundaryType(0)
	QTextBoundaryFinder__Word     QTextBoundaryFinder__BoundaryType = QTextBoundaryFinder__BoundaryType(1)
	QTextBoundaryFinder__Sentence QTextBoundaryFinder__BoundaryType = QTextBoundaryFinder__BoundaryType(2)
	QTextBoundaryFinder__Line     QTextBoundaryFinder__BoundaryType = QTextBoundaryFinder__BoundaryType(3)
)

// QTextCodec::ConversionFlag
//
//go:generate stringer -type=QTextCodec__ConversionFlag
type QTextCodec__ConversionFlag int64

const (
	QTextCodec__DefaultConversion    QTextCodec__ConversionFlag = QTextCodec__ConversionFlag(0)
	QTextCodec__ConvertInvalidToNull QTextCodec__ConversionFlag = QTextCodec__ConversionFlag(0x80000000)
	QTextCodec__IgnoreHeader         QTextCodec__ConversionFlag = QTextCodec__ConversionFlag(0x1)
	QTextCodec__FreeFunction         QTextCodec__ConversionFlag = QTextCodec__ConversionFlag(0x2)
)

type QTextStream struct {
	ptr unsafe.Pointer
}

type QTextStream_ITF interface {
	QTextStream_PTR() *QTextStream
}

func (ptr *QTextStream) QTextStream_PTR() *QTextStream {
	return ptr
}

func (ptr *QTextStream) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextStream) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextStream(ptr QTextStream_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextStream_PTR().Pointer()
	}
	return nil
}

func NewQTextStreamFromPointer(ptr unsafe.Pointer) (n *QTextStream) {
	n = new(QTextStream)
	n.SetPointer(ptr)
	return
}

// QTextStream::FieldAlignment
//
//go:generate stringer -type=QTextStream__FieldAlignment
type QTextStream__FieldAlignment int64

const (
	QTextStream__AlignLeft            QTextStream__FieldAlignment = QTextStream__FieldAlignment(0)
	QTextStream__AlignRight           QTextStream__FieldAlignment = QTextStream__FieldAlignment(1)
	QTextStream__AlignCenter          QTextStream__FieldAlignment = QTextStream__FieldAlignment(2)
	QTextStream__AlignAccountingStyle QTextStream__FieldAlignment = QTextStream__FieldAlignment(3)
)

// QTextStream::NumberFlag
//
//go:generate stringer -type=QTextStream__NumberFlag
type QTextStream__NumberFlag int64

const (
	QTextStream__ShowBase        QTextStream__NumberFlag = QTextStream__NumberFlag(0x1)
	QTextStream__ForcePoint      QTextStream__NumberFlag = QTextStream__NumberFlag(0x2)
	QTextStream__ForceSign       QTextStream__NumberFlag = QTextStream__NumberFlag(0x4)
	QTextStream__UppercaseBase   QTextStream__NumberFlag = QTextStream__NumberFlag(0x8)
	QTextStream__UppercaseDigits QTextStream__NumberFlag = QTextStream__NumberFlag(0x10)
)

// QTextStream::RealNumberNotation
//
//go:generate stringer -type=QTextStream__RealNumberNotation
type QTextStream__RealNumberNotation int64

const (
	QTextStream__SmartNotation      QTextStream__RealNumberNotation = QTextStream__RealNumberNotation(0)
	QTextStream__FixedNotation      QTextStream__RealNumberNotation = QTextStream__RealNumberNotation(1)
	QTextStream__ScientificNotation QTextStream__RealNumberNotation = QTextStream__RealNumberNotation(2)
)

// QTextStream::Status
//
//go:generate stringer -type=QTextStream__Status
type QTextStream__Status int64

const (
	QTextStream__Ok              QTextStream__Status = QTextStream__Status(0)
	QTextStream__ReadPastEnd     QTextStream__Status = QTextStream__Status(1)
	QTextStream__ReadCorruptData QTextStream__Status = QTextStream__Status(2)
	QTextStream__WriteFailed     QTextStream__Status = QTextStream__Status(3)
)

func NewQTextStream() *QTextStream {
	tmpValue := NewQTextStreamFromPointer(C.QTextStream_NewQTextStream())
	qt.SetFinalizer(tmpValue, (*QTextStream).DestroyQTextStream)
	return tmpValue
}

func NewQTextStream2(device QIODevice_ITF) *QTextStream {
	tmpValue := NewQTextStreamFromPointer(C.QTextStream_NewQTextStream2(PointerFromQIODevice(device)))
	qt.SetFinalizer(tmpValue, (*QTextStream).DestroyQTextStream)
	return tmpValue
}

func NewQTextStream4(stri string, openMode QIODevice__OpenModeFlag) *QTextStream {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQTextStreamFromPointer(C.QTextStream_NewQTextStream4(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.longlong(openMode)))
	qt.SetFinalizer(tmpValue, (*QTextStream).DestroyQTextStream)
	return tmpValue
}

func NewQTextStream5(array QByteArray_ITF, openMode QIODevice__OpenModeFlag) *QTextStream {
	tmpValue := NewQTextStreamFromPointer(C.QTextStream_NewQTextStream5(PointerFromQByteArray(array), C.longlong(openMode)))
	qt.SetFinalizer(tmpValue, (*QTextStream).DestroyQTextStream)
	return tmpValue
}

func NewQTextStream6(array QByteArray_ITF, openMode QIODevice__OpenModeFlag) *QTextStream {
	tmpValue := NewQTextStreamFromPointer(C.QTextStream_NewQTextStream6(PointerFromQByteArray(array), C.longlong(openMode)))
	qt.SetFinalizer(tmpValue, (*QTextStream).DestroyQTextStream)
	return tmpValue
}

func (ptr *QTextStream) Device() *QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := NewQIODeviceFromPointer(C.QTextStream_Device(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextStream) Pos() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QTextStream_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) Read(maxlen int64) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextStream_Read(ptr.Pointer(), C.longlong(maxlen)))
	}
	return ""
}

func (ptr *QTextStream) ReadLine(maxlen int64) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextStream_ReadLine(ptr.Pointer(), C.longlong(maxlen)))
	}
	return ""
}

func (ptr *QTextStream) Status() QTextStream__Status {
	if ptr.Pointer() != nil {
		return QTextStream__Status(C.QTextStream_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextStream) String() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextStream_String(ptr.Pointer()))
	}
	return ""
}

//export callbackQTextStream_DestroyQTextStream
func callbackQTextStream_DestroyQTextStream(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTextStream"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextStreamFromPointer(ptr).DestroyQTextStreamDefault()
	}
}

func (ptr *QTextStream) ConnectDestroyQTextStream(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTextStream"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTextStream", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTextStream", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextStream) DisconnectDestroyQTextStream() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTextStream")
	}
}

func (ptr *QTextStream) DestroyQTextStream() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextStream_DestroyQTextStream(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextStream) DestroyQTextStreamDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextStream_DestroyQTextStreamDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QThread::Priority
//
//go:generate stringer -type=QThread__Priority
type QThread__Priority int64

const (
	QThread__IdlePriority         QThread__Priority = QThread__Priority(0)
	QThread__LowestPriority       QThread__Priority = QThread__Priority(1)
	QThread__LowPriority          QThread__Priority = QThread__Priority(2)
	QThread__NormalPriority       QThread__Priority = QThread__Priority(3)
	QThread__HighPriority         QThread__Priority = QThread__Priority(4)
	QThread__HighestPriority      QThread__Priority = QThread__Priority(5)
	QThread__TimeCriticalPriority QThread__Priority = QThread__Priority(6)
	QThread__InheritPriority      QThread__Priority = QThread__Priority(7)
)

type QTime struct {
	ptr unsafe.Pointer
}

type QTime_ITF interface {
	QTime_PTR() *QTime
}

func (ptr *QTime) QTime_PTR() *QTime {
	return ptr
}

func (ptr *QTime) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTime) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTime(ptr QTime_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTime_PTR().Pointer()
	}
	return nil
}

func NewQTimeFromPointer(ptr unsafe.Pointer) (n *QTime) {
	n = new(QTime)
	n.SetPointer(ptr)
	return
}
func (ptr *QTime) DestroyQTime() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQTime2() *QTime {
	tmpValue := NewQTimeFromPointer(C.QTime_NewQTime2())
	qt.SetFinalizer(tmpValue, (*QTime).DestroyQTime)
	return tmpValue
}

func NewQTime3(h int, m int, s int, ms int) *QTime {
	tmpValue := NewQTimeFromPointer(C.QTime_NewQTime3(C.int(int32(h)), C.int(int32(m)), C.int(int32(s)), C.int(int32(ms))))
	qt.SetFinalizer(tmpValue, (*QTime).DestroyQTime)
	return tmpValue
}

func QTime_FromString(stri string, format Qt__DateFormat) *QTime {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQTimeFromPointer(C.QTime_QTime_FromString(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QTime).DestroyQTime)
	return tmpValue
}

func (ptr *QTime) FromString(stri string, format Qt__DateFormat) *QTime {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQTimeFromPointer(C.QTime_QTime_FromString(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QTime).DestroyQTime)
	return tmpValue
}

func QTime_FromString2(stri string, format string) *QTime {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQTimeFromPointer(C.QTime_QTime_FromString2(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.struct_QtCore_PackedString{data: formatC, len: C.longlong(len(format))}))
	qt.SetFinalizer(tmpValue, (*QTime).DestroyQTime)
	return tmpValue
}

func (ptr *QTime) FromString2(stri string, format string) *QTime {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQTimeFromPointer(C.QTime_QTime_FromString2(C.struct_QtCore_PackedString{data: striC, len: C.longlong(len(stri))}, C.struct_QtCore_PackedString{data: formatC, len: C.longlong(len(format))}))
	qt.SetFinalizer(tmpValue, (*QTime).DestroyQTime)
	return tmpValue
}

func (ptr *QTime) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTime_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func QTime_IsValid2(h int, m int, s int, ms int) bool {
	return int8(C.QTime_QTime_IsValid2(C.int(int32(h)), C.int(int32(m)), C.int(int32(s)), C.int(int32(ms)))) != 0
}

func (ptr *QTime) IsValid2(h int, m int, s int, ms int) bool {
	return int8(C.QTime_QTime_IsValid2(C.int(int32(h)), C.int(int32(m)), C.int(int32(s)), C.int(int32(ms)))) != 0
}

// QTimeLine::Direction
//
//go:generate stringer -type=QTimeLine__Direction
type QTimeLine__Direction int64

const (
	QTimeLine__Forward  QTimeLine__Direction = QTimeLine__Direction(0)
	QTimeLine__Backward QTimeLine__Direction = QTimeLine__Direction(1)
)

// QTimeLine::State
//
//go:generate stringer -type=QTimeLine__State
type QTimeLine__State int64

const (
	QTimeLine__NotRunning QTimeLine__State = QTimeLine__State(0)
	QTimeLine__Paused     QTimeLine__State = QTimeLine__State(1)
	QTimeLine__Running    QTimeLine__State = QTimeLine__State(2)
)

type QTimeZone struct {
	ptr unsafe.Pointer
}

type QTimeZone_ITF interface {
	QTimeZone_PTR() *QTimeZone
}

func (ptr *QTimeZone) QTimeZone_PTR() *QTimeZone {
	return ptr
}

func (ptr *QTimeZone) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTimeZone) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTimeZone(ptr QTimeZone_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimeZone_PTR().Pointer()
	}
	return nil
}

func NewQTimeZoneFromPointer(ptr unsafe.Pointer) (n *QTimeZone) {
	n = new(QTimeZone)
	n.SetPointer(ptr)
	return
}

// QTimeZone::NameType
//
//go:generate stringer -type=QTimeZone__NameType
type QTimeZone__NameType int64

const (
	QTimeZone__DefaultName QTimeZone__NameType = QTimeZone__NameType(0)
	QTimeZone__LongName    QTimeZone__NameType = QTimeZone__NameType(1)
	QTimeZone__ShortName   QTimeZone__NameType = QTimeZone__NameType(2)
	QTimeZone__OffsetName  QTimeZone__NameType = QTimeZone__NameType(3)
)

// QTimeZone::TimeType
//
//go:generate stringer -type=QTimeZone__TimeType
type QTimeZone__TimeType int64

const (
	QTimeZone__StandardTime QTimeZone__TimeType = QTimeZone__TimeType(0)
	QTimeZone__DaylightTime QTimeZone__TimeType = QTimeZone__TimeType(1)
	QTimeZone__GenericTime  QTimeZone__TimeType = QTimeZone__TimeType(2)
)

func NewQTimeZone() *QTimeZone {
	tmpValue := NewQTimeZoneFromPointer(C.QTimeZone_NewQTimeZone())
	qt.SetFinalizer(tmpValue, (*QTimeZone).DestroyQTimeZone)
	return tmpValue
}

func NewQTimeZone2(ianaId QByteArray_ITF) *QTimeZone {
	tmpValue := NewQTimeZoneFromPointer(C.QTimeZone_NewQTimeZone2(PointerFromQByteArray(ianaId)))
	qt.SetFinalizer(tmpValue, (*QTimeZone).DestroyQTimeZone)
	return tmpValue
}

func NewQTimeZone3(offsetSeconds int) *QTimeZone {
	tmpValue := NewQTimeZoneFromPointer(C.QTimeZone_NewQTimeZone3(C.int(int32(offsetSeconds))))
	qt.SetFinalizer(tmpValue, (*QTimeZone).DestroyQTimeZone)
	return tmpValue
}

func NewQTimeZone4(ianaId QByteArray_ITF, offsetSeconds int, name string, abbreviation string, country QLocale__Country, comment string) *QTimeZone {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	var abbreviationC *C.char
	if abbreviation != "" {
		abbreviationC = C.CString(abbreviation)
		defer C.free(unsafe.Pointer(abbreviationC))
	}
	var commentC *C.char
	if comment != "" {
		commentC = C.CString(comment)
		defer C.free(unsafe.Pointer(commentC))
	}
	tmpValue := NewQTimeZoneFromPointer(C.QTimeZone_NewQTimeZone4(PointerFromQByteArray(ianaId), C.int(int32(offsetSeconds)), C.struct_QtCore_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtCore_PackedString{data: abbreviationC, len: C.longlong(len(abbreviation))}, C.longlong(country), C.struct_QtCore_PackedString{data: commentC, len: C.longlong(len(comment))}))
	qt.SetFinalizer(tmpValue, (*QTimeZone).DestroyQTimeZone)
	return tmpValue
}

func NewQTimeZone5(other QTimeZone_ITF) *QTimeZone {
	tmpValue := NewQTimeZoneFromPointer(C.QTimeZone_NewQTimeZone5(PointerFromQTimeZone(other)))
	qt.SetFinalizer(tmpValue, (*QTimeZone).DestroyQTimeZone)
	return tmpValue
}

func (ptr *QTimeZone) Comment() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTimeZone_Comment(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTimeZone) DisplayName(atDateTime QDateTime_ITF, nameType QTimeZone__NameType, locale QLocale_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTimeZone_DisplayName(ptr.Pointer(), PointerFromQDateTime(atDateTime), C.longlong(nameType), PointerFromQLocale(locale)))
	}
	return ""
}

func (ptr *QTimeZone) DisplayName2(timeType QTimeZone__TimeType, nameType QTimeZone__NameType, locale QLocale_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTimeZone_DisplayName2(ptr.Pointer(), C.longlong(timeType), C.longlong(nameType), PointerFromQLocale(locale)))
	}
	return ""
}

func (ptr *QTimeZone) Id() *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QTimeZone_Id(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTimeZone) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTimeZone_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTimeZone) DestroyQTimeZone() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTimeZone_DestroyQTimeZone(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTimeZone) __availableTimeZoneIds_atList(i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QTimeZone___availableTimeZoneIds_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTimeZone) __availableTimeZoneIds_setList(i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTimeZone___availableTimeZoneIds_setList(ptr.Pointer(), PointerFromQByteArray(i))
	}
}

func (ptr *QTimeZone) __availableTimeZoneIds_newList() unsafe.Pointer {
	return C.QTimeZone___availableTimeZoneIds_newList(ptr.Pointer())
}

func (ptr *QTimeZone) __availableTimeZoneIds_atList2(i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QTimeZone___availableTimeZoneIds_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTimeZone) __availableTimeZoneIds_setList2(i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTimeZone___availableTimeZoneIds_setList2(ptr.Pointer(), PointerFromQByteArray(i))
	}
}

func (ptr *QTimeZone) __availableTimeZoneIds_newList2() unsafe.Pointer {
	return C.QTimeZone___availableTimeZoneIds_newList2(ptr.Pointer())
}

func (ptr *QTimeZone) __availableTimeZoneIds_atList3(i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QTimeZone___availableTimeZoneIds_atList3(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTimeZone) __availableTimeZoneIds_setList3(i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTimeZone___availableTimeZoneIds_setList3(ptr.Pointer(), PointerFromQByteArray(i))
	}
}

func (ptr *QTimeZone) __availableTimeZoneIds_newList3() unsafe.Pointer {
	return C.QTimeZone___availableTimeZoneIds_newList3(ptr.Pointer())
}

func (ptr *QTimeZone) __windowsIdToIanaIds_atList(i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QTimeZone___windowsIdToIanaIds_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTimeZone) __windowsIdToIanaIds_setList(i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTimeZone___windowsIdToIanaIds_setList(ptr.Pointer(), PointerFromQByteArray(i))
	}
}

func (ptr *QTimeZone) __windowsIdToIanaIds_newList() unsafe.Pointer {
	return C.QTimeZone___windowsIdToIanaIds_newList(ptr.Pointer())
}

func (ptr *QTimeZone) __windowsIdToIanaIds_atList2(i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QTimeZone___windowsIdToIanaIds_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTimeZone) __windowsIdToIanaIds_setList2(i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTimeZone___windowsIdToIanaIds_setList2(ptr.Pointer(), PointerFromQByteArray(i))
	}
}

func (ptr *QTimeZone) __windowsIdToIanaIds_newList2() unsafe.Pointer {
	return C.QTimeZone___windowsIdToIanaIds_newList2(ptr.Pointer())
}

type QTimer struct {
	QObject
}

type QTimer_ITF interface {
	QObject_ITF
	QTimer_PTR() *QTimer
}

func (ptr *QTimer) QTimer_PTR() *QTimer {
	return ptr
}

func (ptr *QTimer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTimer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQTimer(ptr QTimer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimer_PTR().Pointer()
	}
	return nil
}

func NewQTimerFromPointer(ptr unsafe.Pointer) (n *QTimer) {
	n = new(QTimer)
	n.SetPointer(ptr)
	return
}
func NewQTimer(parent QObject_ITF) *QTimer {
	tmpValue := NewQTimerFromPointer(C.QTimer_NewQTimer(PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTimer_Start
func callbackQTimer_Start(ptr unsafe.Pointer, msec C.int) {
	if signal := qt.GetSignal(ptr, "start"); signal != nil {
		(*(*func(int))(signal))(int(int32(msec)))
	} else {
		NewQTimerFromPointer(ptr).StartDefault(int(int32(msec)))
	}
}

func (ptr *QTimer) ConnectStart(f func(msec int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			f := func(msec int) {
				(*(*func(int))(signal))(msec)
				f(msec)
			}
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTimer) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "start")
	}
}

func (ptr *QTimer) Start(msec int) {
	if ptr.Pointer() != nil {
		C.QTimer_Start(ptr.Pointer(), C.int(int32(msec)))
	}
}

func (ptr *QTimer) StartDefault(msec int) {
	if ptr.Pointer() != nil {
		C.QTimer_StartDefault(ptr.Pointer(), C.int(int32(msec)))
	}
}

//export callbackQTimer_Start2
func callbackQTimer_Start2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "start2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTimerFromPointer(ptr).Start2Default()
	}
}

func (ptr *QTimer) ConnectStart2(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start2"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "start2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTimer) DisconnectStart2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "start2")
	}
}

func (ptr *QTimer) Start2() {
	if ptr.Pointer() != nil {
		C.QTimer_Start2(ptr.Pointer())
	}
}

func (ptr *QTimer) Start2Default() {
	if ptr.Pointer() != nil {
		C.QTimer_Start2Default(ptr.Pointer())
	}
}

//export callbackQTimer_Stop
func callbackQTimer_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTimerFromPointer(ptr).StopDefault()
	}
}

func (ptr *QTimer) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTimer) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QTimer) Stop() {
	if ptr.Pointer() != nil {
		C.QTimer_Stop(ptr.Pointer())
	}
}

func (ptr *QTimer) StopDefault() {
	if ptr.Pointer() != nil {
		C.QTimer_StopDefault(ptr.Pointer())
	}
}

//export callbackQTimer_DestroyQTimer
func callbackQTimer_DestroyQTimer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTimer"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTimerFromPointer(ptr).DestroyQTimerDefault()
	}
}

func (ptr *QTimer) ConnectDestroyQTimer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTimer"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTimer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTimer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTimer) DisconnectDestroyQTimer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTimer")
	}
}

func (ptr *QTimer) DestroyQTimer() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTimer_DestroyQTimer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTimer) DestroyQTimerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTimer_DestroyQTimerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTimerEvent struct {
	QEvent
}

type QTimerEvent_ITF interface {
	QEvent_ITF
	QTimerEvent_PTR() *QTimerEvent
}

func (ptr *QTimerEvent) QTimerEvent_PTR() *QTimerEvent {
	return ptr
}

func (ptr *QTimerEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QTimerEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQTimerEvent(ptr QTimerEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimerEvent_PTR().Pointer()
	}
	return nil
}

func NewQTimerEventFromPointer(ptr unsafe.Pointer) (n *QTimerEvent) {
	n = new(QTimerEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QTimerEvent) DestroyQTimerEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQTimerEvent(timerId int) *QTimerEvent {
	tmpValue := NewQTimerEventFromPointer(C.QTimerEvent_NewQTimerEvent(C.int(int32(timerId))))
	qt.SetFinalizer(tmpValue, (*QTimerEvent).DestroyQTimerEvent)
	return tmpValue
}

type QUrl struct {
	ptr unsafe.Pointer
}

type QUrl_ITF interface {
	QUrl_PTR() *QUrl
}

func (ptr *QUrl) QUrl_PTR() *QUrl {
	return ptr
}

func (ptr *QUrl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QUrl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQUrl(ptr QUrl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUrl_PTR().Pointer()
	}
	return nil
}

func NewQUrlFromPointer(ptr unsafe.Pointer) (n *QUrl) {
	n = new(QUrl)
	n.SetPointer(ptr)
	return
}

// QUrl::ComponentFormattingOption
//
//go:generate stringer -type=QUrl__ComponentFormattingOption
type QUrl__ComponentFormattingOption int64

const (
	QUrl__PrettyDecoded    QUrl__ComponentFormattingOption = QUrl__ComponentFormattingOption(0x000000)
	QUrl__EncodeSpaces     QUrl__ComponentFormattingOption = QUrl__ComponentFormattingOption(0x100000)
	QUrl__EncodeUnicode    QUrl__ComponentFormattingOption = QUrl__ComponentFormattingOption(0x200000)
	QUrl__EncodeDelimiters QUrl__ComponentFormattingOption = QUrl__ComponentFormattingOption(0x400000 | 0x800000)
	QUrl__EncodeReserved   QUrl__ComponentFormattingOption = QUrl__ComponentFormattingOption(0x1000000)
	QUrl__DecodeReserved   QUrl__ComponentFormattingOption = QUrl__ComponentFormattingOption(0x2000000)
	QUrl__FullyEncoded     QUrl__ComponentFormattingOption = QUrl__ComponentFormattingOption(QUrl__EncodeSpaces | QUrl__EncodeUnicode | QUrl__EncodeDelimiters | QUrl__EncodeReserved)
	QUrl__FullyDecoded     QUrl__ComponentFormattingOption = QUrl__ComponentFormattingOption(QUrl__FullyEncoded | QUrl__DecodeReserved | 0x4000000)
)

// QUrl::ParsingMode
//
//go:generate stringer -type=QUrl__ParsingMode
type QUrl__ParsingMode int64

const (
	QUrl__TolerantMode QUrl__ParsingMode = QUrl__ParsingMode(0)
	QUrl__StrictMode   QUrl__ParsingMode = QUrl__ParsingMode(1)
	QUrl__DecodedMode  QUrl__ParsingMode = QUrl__ParsingMode(2)
)

// QUrl::UrlFormattingOption
//
//go:generate stringer -type=QUrl__UrlFormattingOption
type QUrl__UrlFormattingOption int64

const (
	QUrl__None                  QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x0)
	QUrl__RemoveScheme          QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x1)
	QUrl__RemovePassword        QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x2)
	QUrl__RemoveUserInfo        QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(QUrl__RemovePassword | 0x4)
	QUrl__RemovePort            QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x8)
	QUrl__RemoveAuthority       QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(QUrl__RemoveUserInfo | QUrl__RemovePort | 0x10)
	QUrl__RemovePath            QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x20)
	QUrl__RemoveQuery           QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x40)
	QUrl__RemoveFragment        QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x80)
	QUrl__PreferLocalFile       QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x200)
	QUrl__StripTrailingSlash    QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x400)
	QUrl__RemoveFilename        QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x800)
	QUrl__NormalizePathSegments QUrl__UrlFormattingOption = QUrl__UrlFormattingOption(0x1000)
)

// QUrl::UserInputResolutionOption
//
//go:generate stringer -type=QUrl__UserInputResolutionOption
type QUrl__UserInputResolutionOption int64

const (
	QUrl__DefaultResolution QUrl__UserInputResolutionOption = QUrl__UserInputResolutionOption(0)
	QUrl__AssumeLocalFile   QUrl__UserInputResolutionOption = QUrl__UserInputResolutionOption(1)
)

func NewQUrl() *QUrl {
	tmpValue := NewQUrlFromPointer(C.QUrl_NewQUrl())
	qt.SetFinalizer(tmpValue, (*QUrl).DestroyQUrl)
	return tmpValue
}

func NewQUrl2(other QUrl_ITF) *QUrl {
	tmpValue := NewQUrlFromPointer(C.QUrl_NewQUrl2(PointerFromQUrl(other)))
	qt.SetFinalizer(tmpValue, (*QUrl).DestroyQUrl)
	return tmpValue
}

func NewQUrl3(url string, parsingMode QUrl__ParsingMode) *QUrl {
	var urlC *C.char
	if url != "" {
		urlC = C.CString(url)
		defer C.free(unsafe.Pointer(urlC))
	}
	tmpValue := NewQUrlFromPointer(C.QUrl_NewQUrl3(C.struct_QtCore_PackedString{data: urlC, len: C.longlong(len(url))}, C.longlong(parsingMode)))
	qt.SetFinalizer(tmpValue, (*QUrl).DestroyQUrl)
	return tmpValue
}

func NewQUrl4(other QUrl_ITF) *QUrl {
	tmpValue := NewQUrlFromPointer(C.QUrl_NewQUrl4(PointerFromQUrl(other)))
	qt.SetFinalizer(tmpValue, (*QUrl).DestroyQUrl)
	return tmpValue
}

func (ptr *QUrl) Clear() {
	if ptr.Pointer() != nil {
		C.QUrl_Clear(ptr.Pointer())
	}
}

func (ptr *QUrl) FileName(options QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QUrl_FileName(ptr.Pointer(), C.longlong(options)))
	}
	return ""
}

func (ptr *QUrl) Fragment(options QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QUrl_Fragment(ptr.Pointer(), C.longlong(options)))
	}
	return ""
}

func (ptr *QUrl) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QUrl_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUrl) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QUrl_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUrl) Password(options QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QUrl_Password(ptr.Pointer(), C.longlong(options)))
	}
	return ""
}

func (ptr *QUrl) Path(options QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QUrl_Path(ptr.Pointer(), C.longlong(options)))
	}
	return ""
}

func (ptr *QUrl) Query(options QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QUrl_Query(ptr.Pointer(), C.longlong(options)))
	}
	return ""
}

func (ptr *QUrl) Url(options QUrl__UrlFormattingOption) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QUrl_Url(ptr.Pointer(), C.longlong(options)))
	}
	return ""
}

func (ptr *QUrl) UserName(options QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QUrl_UserName(ptr.Pointer(), C.longlong(options)))
	}
	return ""
}

func (ptr *QUrl) DestroyQUrl() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QUrl_DestroyQUrl(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QUrl) __allEncodedQueryItemValues_atList(i int) *QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := NewQByteArrayFromPointer(C.QUrl___allEncodedQueryItemValues_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QUrl) __allEncodedQueryItemValues_setList(i QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QUrl___allEncodedQueryItemValues_setList(ptr.Pointer(), PointerFromQByteArray(i))
	}
}

func (ptr *QUrl) __allEncodedQueryItemValues_newList() unsafe.Pointer {
	return C.QUrl___allEncodedQueryItemValues_newList(ptr.Pointer())
}

func (ptr *QUrl) __fromStringList_atList(i int) *QUrl {
	if ptr.Pointer() != nil {
		tmpValue := NewQUrlFromPointer(C.QUrl___fromStringList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QUrl) __fromStringList_setList(i QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QUrl___fromStringList_setList(ptr.Pointer(), PointerFromQUrl(i))
	}
}

func (ptr *QUrl) __fromStringList_newList() unsafe.Pointer {
	return C.QUrl___fromStringList_newList(ptr.Pointer())
}

func (ptr *QUrl) __toStringList_urls_atList(i int) *QUrl {
	if ptr.Pointer() != nil {
		tmpValue := NewQUrlFromPointer(C.QUrl___toStringList_urls_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QUrl) __toStringList_urls_setList(i QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QUrl___toStringList_urls_setList(ptr.Pointer(), PointerFromQUrl(i))
	}
}

func (ptr *QUrl) __toStringList_urls_newList() unsafe.Pointer {
	return C.QUrl___toStringList_urls_newList(ptr.Pointer())
}

type QUuid struct {
	ptr unsafe.Pointer
}

type QUuid_ITF interface {
	QUuid_PTR() *QUuid
}

func (ptr *QUuid) QUuid_PTR() *QUuid {
	return ptr
}

func (ptr *QUuid) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QUuid) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQUuid(ptr QUuid_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUuid_PTR().Pointer()
	}
	return nil
}

func NewQUuidFromPointer(ptr unsafe.Pointer) (n *QUuid) {
	n = new(QUuid)
	n.SetPointer(ptr)
	return
}
func (ptr *QUuid) DestroyQUuid() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QUuid__StringFormat
//QUuid::StringFormat
type QUuid__StringFormat int64

const (
	QUuid__WithBraces    QUuid__StringFormat = QUuid__StringFormat(0)
	QUuid__WithoutBraces QUuid__StringFormat = QUuid__StringFormat(1)
	QUuid__Id128         QUuid__StringFormat = QUuid__StringFormat(3)
)

// QUuid::Variant
//
//go:generate stringer -type=QUuid__Variant
type QUuid__Variant int64

const (
	QUuid__VarUnknown QUuid__Variant = QUuid__Variant(-1)
	QUuid__NCS        QUuid__Variant = QUuid__Variant(0)
	QUuid__DCE        QUuid__Variant = QUuid__Variant(2)
	QUuid__Microsoft  QUuid__Variant = QUuid__Variant(6)
	QUuid__Reserved   QUuid__Variant = QUuid__Variant(7)
)

// QUuid::Version
//
//go:generate stringer -type=QUuid__Version
type QUuid__Version int64

const (
	QUuid__VerUnknown    QUuid__Version = QUuid__Version(-1)
	QUuid__Time          QUuid__Version = QUuid__Version(1)
	QUuid__EmbeddedPOSIX QUuid__Version = QUuid__Version(2)
	QUuid__Md5           QUuid__Version = QUuid__Version(3)
	QUuid__Name          QUuid__Version = QUuid__Version(QUuid__Md5)
	QUuid__Random        QUuid__Version = QUuid__Version(4)
	QUuid__Sha1          QUuid__Version = QUuid__Version(5)
)

func NewQUuid2() *QUuid {
	tmpValue := NewQUuidFromPointer(C.QUuid_NewQUuid2())
	qt.SetFinalizer(tmpValue, (*QUuid).DestroyQUuid)
	return tmpValue
}

func NewQUuid3(l uint, w1 uint16, w2 uint16, b1 string, b2 string, b3 string, b4 string, b5 string, b6 string, b7 string, b8 string) *QUuid {
	var b1C *C.char
	if b1 != "" {
		b1C = C.CString(b1)
		defer C.free(unsafe.Pointer(b1C))
	}
	var b2C *C.char
	if b2 != "" {
		b2C = C.CString(b2)
		defer C.free(unsafe.Pointer(b2C))
	}
	var b3C *C.char
	if b3 != "" {
		b3C = C.CString(b3)
		defer C.free(unsafe.Pointer(b3C))
	}
	var b4C *C.char
	if b4 != "" {
		b4C = C.CString(b4)
		defer C.free(unsafe.Pointer(b4C))
	}
	var b5C *C.char
	if b5 != "" {
		b5C = C.CString(b5)
		defer C.free(unsafe.Pointer(b5C))
	}
	var b6C *C.char
	if b6 != "" {
		b6C = C.CString(b6)
		defer C.free(unsafe.Pointer(b6C))
	}
	var b7C *C.char
	if b7 != "" {
		b7C = C.CString(b7)
		defer C.free(unsafe.Pointer(b7C))
	}
	var b8C *C.char
	if b8 != "" {
		b8C = C.CString(b8)
		defer C.free(unsafe.Pointer(b8C))
	}
	tmpValue := NewQUuidFromPointer(C.QUuid_NewQUuid3(C.uint(uint32(l)), C.ushort(w1), C.ushort(w2), b1C, b2C, b3C, b4C, b5C, b6C, b7C, b8C))
	qt.SetFinalizer(tmpValue, (*QUuid).DestroyQUuid)
	return tmpValue
}

func NewQUuid4(text string) *QUuid {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQUuidFromPointer(C.QUuid_NewQUuid4(C.struct_QtCore_PackedString{data: textC, len: C.longlong(len(text))}))
	qt.SetFinalizer(tmpValue, (*QUuid).DestroyQUuid)
	return tmpValue
}

func NewQUuid(text QByteArray_ITF) *QUuid {
	tmpValue := NewQUuidFromPointer(C.QUuid_NewQUuid(PointerFromQByteArray(text)))
	qt.SetFinalizer(tmpValue, (*QUuid).DestroyQUuid)
	return tmpValue
}

func QUuid_FromString(text QStringView_ITF) *QUuid {
	tmpValue := NewQUuidFromPointer(C.QUuid_QUuid_FromString(PointerFromQStringView(text)))
	qt.SetFinalizer(tmpValue, (*QUuid).DestroyQUuid)
	return tmpValue
}

func (ptr *QUuid) FromString(text QStringView_ITF) *QUuid {
	tmpValue := NewQUuidFromPointer(C.QUuid_QUuid_FromString(PointerFromQStringView(text)))
	qt.SetFinalizer(tmpValue, (*QUuid).DestroyQUuid)
	return tmpValue
}

func QUuid_FromString2(text QLatin1String_ITF) *QUuid {
	tmpValue := NewQUuidFromPointer(C.QUuid_QUuid_FromString2(PointerFromQLatin1String(text)))
	qt.SetFinalizer(tmpValue, (*QUuid).DestroyQUuid)
	return tmpValue
}

func (ptr *QUuid) FromString2(text QLatin1String_ITF) *QUuid {
	tmpValue := NewQUuidFromPointer(C.QUuid_QUuid_FromString2(PointerFromQLatin1String(text)))
	qt.SetFinalizer(tmpValue, (*QUuid).DestroyQUuid)
	return tmpValue
}

func (ptr *QUuid) Variant() QUuid__Variant {
	if ptr.Pointer() != nil {
		return QUuid__Variant(C.QUuid_Variant(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUuid) Version() QUuid__Version {
	if ptr.Pointer() != nil {
		return QUuid__Version(C.QUuid_Version(ptr.Pointer()))
	}
	return 0
}

type QVariant struct {
	ptr unsafe.Pointer
}

type QVariant_ITF interface {
	QVariant_PTR() *QVariant
}

func (ptr *QVariant) QVariant_PTR() *QVariant {
	return ptr
}

func (ptr *QVariant) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVariant) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVariant(ptr QVariant_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVariant_PTR().Pointer()
	}
	return nil
}

func NewQVariantFromPointer(ptr unsafe.Pointer) (n *QVariant) {
	n = new(QVariant)
	n.SetPointer(ptr)
	return
}

type qVariant_ITF interface{ ToVariant() *QVariant }

func NewQVariant1(i interface{}) *QVariant {
	switch d := i.(type) {
	case *QVariant:
		return d
	case QVariant__Type:
		return NewQVariant2(d)
	case *QDataStream:
		return NewQVariant4(d)
	case int:
		return NewQVariant5(d)
	case uint:
		return NewQVariant6(d)
	case int64:
		return NewQVariant7(d)
	case uint64:
		return NewQVariant8(d)
	case bool:
		return NewQVariant9(d)
	case float64:
		return NewQVariant10(d)
	case float32:
		return NewQVariant11(d)
	case string:
		return NewQVariant12(d)
	case *QByteArray:
		return NewQVariant13(d)
	case *QBitArray:
		return NewQVariant14(d)
	case *QLatin1String:
		return NewQVariant16(d)
	case []string:
		return NewQVariant17(d)
	case *QChar:
		return NewQVariant18(d)
	case *QDate:
		return NewQVariant19(d)
	case *QTime:
		return NewQVariant20(d)
	case *QDateTime:
		return NewQVariant21(d)
	case []*QVariant:
		return NewQVariant22(d)
	case map[string]*QVariant:
		return NewQVariant23(d)
	case *QSize:
		return NewQVariant25(d)
	case *QSizeF:
		return NewQVariant26(d)
	case *QPoint:
		return NewQVariant27(d)
	case *QPointF:
		return NewQVariant28(d)
	case *QLine:
		return NewQVariant29(d)
	case *QLineF:
		return NewQVariant30(d)
	case *QRect:
		return NewQVariant31(d)
	case *QRectF:
		return NewQVariant32(d)
	case *QLocale:
		return NewQVariant33(d)
	case *QRegExp:
		return NewQVariant34(d)
	case *QRegularExpression:
		return NewQVariant35(d)
	case *QEasingCurve:
		return NewQVariant36(d)
	case *QUuid:
		return NewQVariant37(d)
	case *QUrl:
		return NewQVariant38(d)
	case *QJsonValue:
		return NewQVariant39(d)
	case *QJsonObject:
		return NewQVariant40(d)
	case *QJsonArray:
		return NewQVariant41(d)
	case *QJsonDocument:
		return NewQVariant42(d)
	case *QModelIndex:
		return NewQVariant43(d)
	case *QPersistentModelIndex:
		return NewQVariant44(d)
	case qVariant_ITF:
		return d.ToVariant()
	default:
		s := reflect.ValueOf(i)
		if s.Kind() == reflect.Ptr {
			s = s.Elem()
		}
		switch s.Kind() {
		case reflect.Struct:
			tmp := make(map[string]*QVariant, s.NumField())
			for id := 0; id < s.NumField(); id++ {
				field := s.Field(id)
				if !field.CanInterface() {
					continue
				}
				if tag, ok := s.Type().Field(id).Tag.Lookup("json"); ok {
					switch {
					case (strings.HasSuffix(tag, ",omitempty") && isZero(field)) || tag == "-":
					case strings.Count(tag, ",") == 0:
						tmp[tag] = NewQVariant1(field.Interface())
					default:
						tags := strings.Split(tag, ",")
						name := s.Type().Field(id).Name
						if len(tags[0]) != 0 {
							name = tags[0]
						}
						v := NewQVariant1(field.Interface())
						if tags[1] == "string" {
							v = NewQVariant1(v.ToString()) //TODO: pure go type conversion
						}
						tmp[name] = v
					}
				} else {
					tmp[s.Type().Field(id).Name] = NewQVariant1(field.Interface())
				}
			}
			return NewQVariant1(tmp)

		case reflect.Slice, reflect.Array:
			tmp := make([]*QVariant, s.Len())
			for id := 0; id < s.Len(); id++ {
				if field := s.Index(id); field.CanInterface() {
					tmp[id] = NewQVariant1(field.Interface())
				}
			}
			return NewQVariant1(tmp)

		case reflect.Map:
			tmp := make(map[string]*QVariant, s.Len())
			for _, id := range s.MapKeys() {
				if field := s.MapIndex(id); field.CanInterface() {
					if key := NewQVariant1(id.Interface()); key.CanConvert(int(QMetaType__QString)) {
						tmp[key.ToString()] = NewQVariant1(field.Interface())
					}
				}
			}
			return NewQVariant1(tmp)

		case reflect.UnsafePointer:
			return NewQVariant1(uint64(uintptr(s.Interface().(unsafe.Pointer))))

		case reflect.Uintptr:
			return NewQVariant1(uint64(s.Interface().(uintptr)))
		}

		if i != nil && s.Type().ConvertibleTo(reflect.TypeOf(int8(0))) {
			if s.Kind() == reflect.Int64 {
				return NewQVariant1(s.Convert(reflect.TypeOf(int64(0))).Interface())
			}
			return NewQVariant1(s.Interface())
		}

		return NewQVariant()
	}
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	default:
		// This should never happens, but will act as a safeguard for
		// later, as a default value doesn't makes sense here.
		panic(&reflect.ValueError{"reflect.Value.IsZero", v.Kind()})
	}
}
func (ptr *QVariant) ToInterface() interface{} {
	switch ptr.Type() {
	case QVariant__Bool:
		return ptr.ToBool()
	case QVariant__Int:
		return ptr.ToInt(nil)
	case QVariant__UInt:
		return ptr.ToUInt(nil)
	case QVariant__LongLong:
		return ptr.ToLongLong(nil)
	case QVariant__ULongLong:
		return ptr.ToULongLong(nil)
	case QVariant__Double:
		return ptr.ToDouble(nil)
	case QVariant__Map:
		return ptr.ToMap()
	case QVariant__List:
		return ptr.ToList()
	case QVariant__String:
		return ptr.ToString()
	case QVariant__StringList:
		return ptr.ToStringList()
	case QVariant__Hash:
		return ptr.ToHash()
	case QVariant__Image:
		return ptr.ToImage()

	}
	return ptr
}

func (ptr *QVariant) ToGoType(dst interface{}) {
	v := reflect.ValueOf(dst)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch ptr.Type() {
	case QVariant__List:
		d := ptr.ToList()

		switch v.Kind() {
		case reflect.Slice:
			v.Set(reflect.MakeSlice(v.Type(), len(d), len(d)))

		case reflect.Array:
			v.Set(reflect.Indirect(reflect.New(v.Type())))
		}

		for i := 0; i < len(d); i++ {
			switch v.Type().Elem().Kind() {
			case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
				s := reflect.New(v.Type().Elem())
				d[i].ToGoType(s.Interface())
				v.Index(i).Set(reflect.Indirect(s))

			default:
				v.Index(i).Set(reflect.ValueOf(d[i].ToInterface()).Convert(v.Type().Elem()))
			}
		}

	case QVariant__Map:
		d := ptr.ToMap()

		if v.Kind() == reflect.Struct {
			for k, val := range d {
				realName := k
				toInt := false
				for id := 0; id < v.NumField(); id++ {
					if tag, ok := v.Type().Field(id).Tag.Lookup("json"); ok {
						switch {
						case strings.Count(tag, ",") == 0:
							if k == tag {
								realName = v.Type().Field(id).Name
							}
						default:
							tags := strings.Split(tag, ",")
							if k == tags[0] || (len(tags[0]) == 0 && k == v.Type().Field(id).Name) {
								realName = v.Type().Field(id).Name
								if tags[1] == "string" {
									toInt = true
								}
							}
						}
					}
				}
				field := v.FieldByName(realName)
				if !field.IsValid() {
					continue
				}
				if toInt {
					field.Set(reflect.ValueOf(val.ToLongLong(nil)).Convert(field.Type())) //TODO: pure go type conversion
				} else {
					val.ToGoType(field.Addr().Interface())
				}
			}
		} else {
			v.Set(reflect.MakeMapWithSize(v.Type(), len(d)))
			for k, val := range d {
				switch v.Type().Elem().Kind() {
				case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
					s := reflect.New(v.Type().Elem())
					val.ToGoType(s.Interface())
					v.SetMapIndex(reflect.ValueOf(k).Convert(v.Type().Key()), reflect.Indirect(s))

				default:
					v.SetMapIndex(reflect.ValueOf(k).Convert(v.Type().Key()), reflect.ValueOf(val.ToInterface()).Convert(v.Type().Elem()))
				}
			}
		}

	default:
		if v.Kind() == reflect.String {
			v.Set(reflect.ValueOf(ptr.ToString()))
		} else {
			v.Set(reflect.ValueOf(ptr.ToInterface()).Convert(v.Type()))
		}
	}
}

// QVariant::Type
//
//go:generate stringer -type=QVariant__Type
type QVariant__Type int64

const (
	QVariant__Invalid              QVariant__Type = QVariant__Type(QMetaType__UnknownType)
	QVariant__Bool                 QVariant__Type = QVariant__Type(QMetaType__Bool)
	QVariant__Int                  QVariant__Type = QVariant__Type(QMetaType__Int)
	QVariant__UInt                 QVariant__Type = QVariant__Type(QMetaType__UInt)
	QVariant__LongLong             QVariant__Type = QVariant__Type(QMetaType__LongLong)
	QVariant__ULongLong            QVariant__Type = QVariant__Type(QMetaType__ULongLong)
	QVariant__Double               QVariant__Type = QVariant__Type(QMetaType__Double)
	QVariant__Char                 QVariant__Type = QVariant__Type(QMetaType__QChar)
	QVariant__Map                  QVariant__Type = QVariant__Type(QMetaType__QVariantMap)
	QVariant__List                 QVariant__Type = QVariant__Type(QMetaType__QVariantList)
	QVariant__String               QVariant__Type = QVariant__Type(QMetaType__QString)
	QVariant__StringList           QVariant__Type = QVariant__Type(QMetaType__QStringList)
	QVariant__ByteArray            QVariant__Type = QVariant__Type(QMetaType__QByteArray)
	QVariant__BitArray             QVariant__Type = QVariant__Type(QMetaType__QBitArray)
	QVariant__Date                 QVariant__Type = QVariant__Type(QMetaType__QDate)
	QVariant__Time                 QVariant__Type = QVariant__Type(QMetaType__QTime)
	QVariant__DateTime             QVariant__Type = QVariant__Type(QMetaType__QDateTime)
	QVariant__Url                  QVariant__Type = QVariant__Type(QMetaType__QUrl)
	QVariant__Locale               QVariant__Type = QVariant__Type(QMetaType__QLocale)
	QVariant__Rect                 QVariant__Type = QVariant__Type(QMetaType__QRect)
	QVariant__RectF                QVariant__Type = QVariant__Type(QMetaType__QRectF)
	QVariant__Size                 QVariant__Type = QVariant__Type(QMetaType__QSize)
	QVariant__SizeF                QVariant__Type = QVariant__Type(QMetaType__QSizeF)
	QVariant__Line                 QVariant__Type = QVariant__Type(QMetaType__QLine)
	QVariant__LineF                QVariant__Type = QVariant__Type(QMetaType__QLineF)
	QVariant__Point                QVariant__Type = QVariant__Type(QMetaType__QPoint)
	QVariant__PointF               QVariant__Type = QVariant__Type(QMetaType__QPointF)
	QVariant__RegExp               QVariant__Type = QVariant__Type(QMetaType__QRegExp)
	QVariant__RegularExpression    QVariant__Type = QVariant__Type(QMetaType__QRegularExpression)
	QVariant__Hash                 QVariant__Type = QVariant__Type(QMetaType__QVariantHash)
	QVariant__EasingCurve          QVariant__Type = QVariant__Type(QMetaType__QEasingCurve)
	QVariant__Uuid                 QVariant__Type = QVariant__Type(QMetaType__QUuid)
	QVariant__ModelIndex           QVariant__Type = QVariant__Type(QMetaType__QModelIndex)
	QVariant__PersistentModelIndex QVariant__Type = QVariant__Type(QMetaType__QPersistentModelIndex)
	QVariant__Font                 QVariant__Type = QVariant__Type(QMetaType__QFont)
	QVariant__Pixmap               QVariant__Type = QVariant__Type(QMetaType__QPixmap)
	QVariant__Brush                QVariant__Type = QVariant__Type(QMetaType__QBrush)
	QVariant__Color                QVariant__Type = QVariant__Type(QMetaType__QColor)
	QVariant__Palette              QVariant__Type = QVariant__Type(QMetaType__QPalette)
	QVariant__Image                QVariant__Type = QVariant__Type(QMetaType__QImage)
	QVariant__Polygon              QVariant__Type = QVariant__Type(QMetaType__QPolygon)
	QVariant__Region               QVariant__Type = QVariant__Type(QMetaType__QRegion)
	QVariant__Bitmap               QVariant__Type = QVariant__Type(QMetaType__QBitmap)
	QVariant__Cursor               QVariant__Type = QVariant__Type(QMetaType__QCursor)
	QVariant__KeySequence          QVariant__Type = QVariant__Type(QMetaType__QKeySequence)
	QVariant__Pen                  QVariant__Type = QVariant__Type(QMetaType__QPen)
	QVariant__TextLength           QVariant__Type = QVariant__Type(QMetaType__QTextLength)
	QVariant__TextFormat           QVariant__Type = QVariant__Type(QMetaType__QTextFormat)
	QVariant__Matrix               QVariant__Type = QVariant__Type(QMetaType__QMatrix)
	QVariant__Transform            QVariant__Type = QVariant__Type(QMetaType__QTransform)
	QVariant__Matrix4x4            QVariant__Type = QVariant__Type(QMetaType__QMatrix4x4)
	QVariant__Vector2D             QVariant__Type = QVariant__Type(QMetaType__QVector2D)
	QVariant__Vector3D             QVariant__Type = QVariant__Type(QMetaType__QVector3D)
	QVariant__Vector4D             QVariant__Type = QVariant__Type(QMetaType__QVector4D)
	QVariant__Quaternion           QVariant__Type = QVariant__Type(QMetaType__QQuaternion)
	QVariant__PolygonF             QVariant__Type = QVariant__Type(QMetaType__QPolygonF)
	QVariant__Icon                 QVariant__Type = QVariant__Type(QMetaType__QIcon)
	QVariant__SizePolicy           QVariant__Type = QVariant__Type(QMetaType__QSizePolicy)
	QVariant__UserType             QVariant__Type = QVariant__Type(QMetaType__User)
	QVariant__LastType             QVariant__Type = QVariant__Type(0xffffffff)
)

func NewQVariant() *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant())
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant2(ty QVariant__Type) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant2(C.longlong(ty)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant3(typeId int, copy unsafe.Pointer) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant3(C.int(int32(typeId)), copy))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant4(s QDataStream_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant4(PointerFromQDataStream(s)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant5(val int) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant5(C.int(int32(val))))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant6(val uint) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant6(C.uint(uint32(val))))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant7(val int64) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant7(C.longlong(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant8(val uint64) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant8(C.ulonglong(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant9(val bool) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant9(C.char(int8(qt.GoBoolToInt(val)))))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant10(val float64) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant10(C.double(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant11(val float32) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant11(C.float(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant12(val string) *QVariant {
	var valC *C.char
	if val != "" {
		valC = C.CString(val)
		defer C.free(unsafe.Pointer(valC))
	}
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant12(valC))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant13(val QByteArray_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant13(PointerFromQByteArray(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant14(val QBitArray_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant14(PointerFromQBitArray(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant15(val string) *QVariant {
	var valC *C.char
	if val != "" {
		valC = C.CString(val)
		defer C.free(unsafe.Pointer(valC))
	}
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant15(C.struct_QtCore_PackedString{data: valC, len: C.longlong(len(val))}))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant16(val QLatin1String_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant16(PointerFromQLatin1String(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant17(val []string) *QVariant {
	valC := C.CString(strings.Join(val, "¡¦!"))
	defer C.free(unsafe.Pointer(valC))
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant17(C.struct_QtCore_PackedString{data: valC, len: C.longlong(len(strings.Join(val, "¡¦!")))}))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant18(c QChar_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant18(PointerFromQChar(c)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant19(val QDate_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant19(PointerFromQDate(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant20(val QTime_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant20(PointerFromQTime(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant21(val QDateTime_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant21(PointerFromQDateTime(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant22(val []*QVariant) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant22(func() unsafe.Pointer {
		tmpList := NewQVariantFromPointer(NewQVariantFromPointer(nil).__QVariant_val_newList22())
		for _, v := range val {
			tmpList.__QVariant_val_setList22(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant23(val map[string]*QVariant) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant23(func() unsafe.Pointer {
		tmpList := NewQVariantFromPointer(NewQVariantFromPointer(nil).__QVariant_val_newList23())
		for k, v := range val {
			tmpList.__QVariant_val_setList23(k, v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant24(val map[string]*QVariant) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant24(func() unsafe.Pointer {
		tmpList := NewQVariantFromPointer(NewQVariantFromPointer(nil).__QVariant_val_newList24())
		for k, v := range val {
			tmpList.__QVariant_val_setList24(k, v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant25(val QSize_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant25(PointerFromQSize(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant26(val QSizeF_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant26(PointerFromQSizeF(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant27(val QPoint_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant27(PointerFromQPoint(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant28(val QPointF_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant28(PointerFromQPointF(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant29(val QLine_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant29(PointerFromQLine(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant30(val QLineF_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant30(PointerFromQLineF(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant31(val QRect_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant31(PointerFromQRect(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant32(val QRectF_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant32(PointerFromQRectF(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant33(l QLocale_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant33(PointerFromQLocale(l)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant34(regExp QRegExp_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant34(PointerFromQRegExp(regExp)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant35(re QRegularExpression_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant35(PointerFromQRegularExpression(re)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant36(val QEasingCurve_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant36(PointerFromQEasingCurve(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant37(val QUuid_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant37(PointerFromQUuid(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant38(val QUrl_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant38(PointerFromQUrl(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant39(val QJsonValue_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant39(PointerFromQJsonValue(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant40(val QJsonObject_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant40(PointerFromQJsonObject(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant41(val QJsonArray_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant41(PointerFromQJsonArray(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant42(val QJsonDocument_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant42(PointerFromQJsonDocument(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant43(val QModelIndex_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant43(PointerFromQModelIndex(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant44(val QPersistentModelIndex_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant44(PointerFromQPersistentModelIndex(val)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func NewQVariant45(other QVariant_ITF) *QVariant {
	tmpValue := NewQVariantFromPointer(C.QVariant_NewQVariant45(PointerFromQVariant(other)))
	qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QVariant) CanConvert(targetTypeId int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QVariant_CanConvert(ptr.Pointer(), C.int(int32(targetTypeId)))) != 0
	}
	return false
}

func (ptr *QVariant) Clear() {
	if ptr.Pointer() != nil {
		C.QVariant_Clear(ptr.Pointer())
	}
}

func (ptr *QVariant) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVariant_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVariant) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVariant_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVariant) ToBool() bool {
	if ptr.Pointer() != nil {
		return int8(C.QVariant_ToBool(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QVariant) ToDouble(ok *bool) float64 {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return float64(C.QVariant_ToDouble(ptr.Pointer(), &okC))
	}
	return 0
}

func (ptr *QVariant) ToHash() map[string]*QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) map[string]*QVariant {
			out := make(map[string]*QVariant, int(l.len))
			tmpList := NewQVariantFromPointer(l.data)
			for i, v := range tmpList.__toHash_keyList() {
				out[v] = tmpList.__toHash_atList(v, i)
			}
			return out
		}(C.QVariant_ToHash(ptr.Pointer()))
	}
	return make(map[string]*QVariant, 0)
}

func (ptr *QVariant) ToInt(ok *bool) int {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return int(int32(C.QVariant_ToInt(ptr.Pointer(), &okC)))
	}
	return 0
}

func (ptr *QVariant) ToList() []*QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []*QVariant {
			out := make([]*QVariant, int(l.len))
			tmpList := NewQVariantFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__toList_atList(i)
			}
			return out
		}(C.QVariant_ToList(ptr.Pointer()))
	}
	return make([]*QVariant, 0)
}

func (ptr *QVariant) ToLongLong(ok *bool) int64 {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return int64(C.QVariant_ToLongLong(ptr.Pointer(), &okC))
	}
	return 0
}

func (ptr *QVariant) ToMap() map[string]*QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) map[string]*QVariant {
			out := make(map[string]*QVariant, int(l.len))
			tmpList := NewQVariantFromPointer(l.data)
			for i, v := range tmpList.__toMap_keyList() {
				out[v] = tmpList.__toMap_atList(v, i)
			}
			return out
		}(C.QVariant_ToMap(ptr.Pointer()))
	}
	return make(map[string]*QVariant, 0)
}

func (ptr *QVariant) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVariant_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QVariant) ToStringList() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QVariant_ToStringList(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QVariant) ToUInt(ok *bool) uint {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return uint(uint32(C.QVariant_ToUInt(ptr.Pointer(), &okC)))
	}
	return 0
}

func (ptr *QVariant) ToULongLong(ok *bool) uint64 {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return uint64(C.QVariant_ToULongLong(ptr.Pointer(), &okC))
	}
	return 0
}

func (ptr *QVariant) Type() QVariant__Type {
	if ptr.Pointer() != nil {
		return QVariant__Type(C.QVariant_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVariant) DestroyQVariant() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVariant_DestroyQVariant(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVariant) ToImage() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QVariant_ToImage(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) __QVariant_val_atList22(i int) *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QVariant___QVariant_val_atList22(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVariant) __QVariant_val_setList22(i QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QVariant___QVariant_val_setList22(ptr.Pointer(), PointerFromQVariant(i))
	}
}

func (ptr *QVariant) __QVariant_val_newList22() unsafe.Pointer {
	return C.QVariant___QVariant_val_newList22(ptr.Pointer())
}

func (ptr *QVariant) __QVariant_val_atList23(v string, i int) *QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := NewQVariantFromPointer(C.QVariant___QVariant_val_atList23(ptr.Pointer(), C.struct_QtCore_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVariant) __QVariant_val_setList23(key string, i QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QVariant___QVariant_val_setList23(ptr.Pointer(), C.struct_QtCore_PackedString{data: keyC, len: C.longlong(len(key))}, PointerFromQVariant(i))
	}
}

func (ptr *QVariant) __QVariant_val_newList23() unsafe.Pointer {
	return C.QVariant___QVariant_val_newList23(ptr.Pointer())
}

func (ptr *QVariant) __QVariant_val_keyList23() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQVariantFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____QVariant_val_keyList_atList23(i)
			}
			return out
		}(C.QVariant___QVariant_val_keyList23(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QVariant) __QVariant_val_atList24(v string, i int) *QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := NewQVariantFromPointer(C.QVariant___QVariant_val_atList24(ptr.Pointer(), C.struct_QtCore_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVariant) __QVariant_val_setList24(key string, i QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QVariant___QVariant_val_setList24(ptr.Pointer(), C.struct_QtCore_PackedString{data: keyC, len: C.longlong(len(key))}, PointerFromQVariant(i))
	}
}

func (ptr *QVariant) __QVariant_val_newList24() unsafe.Pointer {
	return C.QVariant___QVariant_val_newList24(ptr.Pointer())
}

func (ptr *QVariant) __QVariant_val_keyList24() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQVariantFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____QVariant_val_keyList_atList24(i)
			}
			return out
		}(C.QVariant___QVariant_val_keyList24(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QVariant) __toHash_atList(v string, i int) *QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := NewQVariantFromPointer(C.QVariant___toHash_atList(ptr.Pointer(), C.struct_QtCore_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVariant) __toHash_setList(key string, i QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QVariant___toHash_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: keyC, len: C.longlong(len(key))}, PointerFromQVariant(i))
	}
}

func (ptr *QVariant) __toHash_newList() unsafe.Pointer {
	return C.QVariant___toHash_newList(ptr.Pointer())
}

func (ptr *QVariant) __toHash_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQVariantFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____toHash_keyList_atList(i)
			}
			return out
		}(C.QVariant___toHash_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QVariant) __toList_atList(i int) *QVariant {
	if ptr.Pointer() != nil {
		tmpValue := NewQVariantFromPointer(C.QVariant___toList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVariant) __toList_setList(i QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QVariant___toList_setList(ptr.Pointer(), PointerFromQVariant(i))
	}
}

func (ptr *QVariant) __toList_newList() unsafe.Pointer {
	return C.QVariant___toList_newList(ptr.Pointer())
}

func (ptr *QVariant) __toMap_atList(v string, i int) *QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := NewQVariantFromPointer(C.QVariant___toMap_atList(ptr.Pointer(), C.struct_QtCore_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QVariant) __toMap_setList(key string, i QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QVariant___toMap_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: keyC, len: C.longlong(len(key))}, PointerFromQVariant(i))
	}
}

func (ptr *QVariant) __toMap_newList() unsafe.Pointer {
	return C.QVariant___toMap_newList(ptr.Pointer())
}

func (ptr *QVariant) __toMap_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtCore_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQVariantFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____toMap_keyList_atList(i)
			}
			return out
		}(C.QVariant___toMap_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QVariant) ____QVariant_val_keyList_atList23(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVariant_____QVariant_val_keyList_atList23(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QVariant) ____QVariant_val_keyList_setList23(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QVariant_____QVariant_val_keyList_setList23(ptr.Pointer(), C.struct_QtCore_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QVariant) ____QVariant_val_keyList_newList23() unsafe.Pointer {
	return C.QVariant_____QVariant_val_keyList_newList23(ptr.Pointer())
}

func (ptr *QVariant) ____QVariant_val_keyList_atList24(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVariant_____QVariant_val_keyList_atList24(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QVariant) ____QVariant_val_keyList_setList24(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QVariant_____QVariant_val_keyList_setList24(ptr.Pointer(), C.struct_QtCore_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QVariant) ____QVariant_val_keyList_newList24() unsafe.Pointer {
	return C.QVariant_____QVariant_val_keyList_newList24(ptr.Pointer())
}

func (ptr *QVariant) ____toHash_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVariant_____toHash_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QVariant) ____toHash_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QVariant_____toHash_keyList_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QVariant) ____toHash_keyList_newList() unsafe.Pointer {
	return C.QVariant_____toHash_keyList_newList(ptr.Pointer())
}

func (ptr *QVariant) ____toMap_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVariant_____toMap_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QVariant) ____toMap_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QVariant_____toMap_keyList_setList(ptr.Pointer(), C.struct_QtCore_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QVariant) ____toMap_keyList_newList() unsafe.Pointer {
	return C.QVariant_____toMap_keyList_newList(ptr.Pointer())
}

type QVector struct {
	ptr unsafe.Pointer
}

type QVector_ITF interface {
	QVector_PTR() *QVector
}

func (ptr *QVector) QVector_PTR() *QVector {
	return ptr
}

func (ptr *QVector) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVector) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVector(ptr QVector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector_PTR().Pointer()
	}
	return nil
}

func NewQVectorFromPointer(ptr unsafe.Pointer) (n *QVector) {
	n = new(QVector)
	n.SetPointer(ptr)
	return
}

// QXmlStreamReader::Error
//
//go:generate stringer -type=QXmlStreamReader__Error
type QXmlStreamReader__Error int64

const (
	QXmlStreamReader__NoError                     QXmlStreamReader__Error = QXmlStreamReader__Error(0)
	QXmlStreamReader__UnexpectedElementError      QXmlStreamReader__Error = QXmlStreamReader__Error(1)
	QXmlStreamReader__CustomError                 QXmlStreamReader__Error = QXmlStreamReader__Error(2)
	QXmlStreamReader__NotWellFormedError          QXmlStreamReader__Error = QXmlStreamReader__Error(3)
	QXmlStreamReader__PrematureEndOfDocumentError QXmlStreamReader__Error = QXmlStreamReader__Error(4)
)

// QXmlStreamReader::ReadElementTextBehaviour
//
//go:generate stringer -type=QXmlStreamReader__ReadElementTextBehaviour
type QXmlStreamReader__ReadElementTextBehaviour int64

const (
	QXmlStreamReader__ErrorOnUnexpectedElement QXmlStreamReader__ReadElementTextBehaviour = QXmlStreamReader__ReadElementTextBehaviour(0)
	QXmlStreamReader__IncludeChildElements     QXmlStreamReader__ReadElementTextBehaviour = QXmlStreamReader__ReadElementTextBehaviour(1)
	QXmlStreamReader__SkipChildElements        QXmlStreamReader__ReadElementTextBehaviour = QXmlStreamReader__ReadElementTextBehaviour(2)
)

// QXmlStreamReader::TokenType
//
//go:generate stringer -type=QXmlStreamReader__TokenType
type QXmlStreamReader__TokenType int64

const (
	QXmlStreamReader__NoToken               QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(0)
	QXmlStreamReader__Invalid               QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(1)
	QXmlStreamReader__StartDocument         QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(2)
	QXmlStreamReader__EndDocument           QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(3)
	QXmlStreamReader__StartElement          QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(4)
	QXmlStreamReader__EndElement            QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(5)
	QXmlStreamReader__Characters            QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(6)
	QXmlStreamReader__Comment               QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(7)
	QXmlStreamReader__DTD                   QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(8)
	QXmlStreamReader__EntityReference       QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(9)
	QXmlStreamReader__ProcessingInstruction QXmlStreamReader__TokenType = QXmlStreamReader__TokenType(10)
)

type Qt struct {
	ptr unsafe.Pointer
}

type Qt_ITF interface {
	Qt_PTR() *Qt
}

func (ptr *Qt) Qt_PTR() *Qt {
	return ptr
}

func (ptr *Qt) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Qt) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQt(ptr Qt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Qt_PTR().Pointer()
	}
	return nil
}

func NewQtFromPointer(ptr unsafe.Pointer) (n *Qt) {
	n = new(Qt)
	n.SetPointer(ptr)
	return
}
func (ptr *Qt) DestroyQt() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=Qt__AlignmentFlag
//Qt::AlignmentFlag
type Qt__AlignmentFlag int64

const (
	Qt__AlignLeft            Qt__AlignmentFlag = Qt__AlignmentFlag(0x0001)
	Qt__AlignLeading         Qt__AlignmentFlag = Qt__AlignmentFlag(Qt__AlignLeft)
	Qt__AlignRight           Qt__AlignmentFlag = Qt__AlignmentFlag(0x0002)
	Qt__AlignTrailing        Qt__AlignmentFlag = Qt__AlignmentFlag(Qt__AlignRight)
	Qt__AlignHCenter         Qt__AlignmentFlag = Qt__AlignmentFlag(0x0004)
	Qt__AlignJustify         Qt__AlignmentFlag = Qt__AlignmentFlag(0x0008)
	Qt__AlignAbsolute        Qt__AlignmentFlag = Qt__AlignmentFlag(0x0010)
	Qt__AlignHorizontal_Mask Qt__AlignmentFlag = Qt__AlignmentFlag(Qt__AlignLeft | Qt__AlignRight | Qt__AlignHCenter | Qt__AlignJustify | Qt__AlignAbsolute)
	Qt__AlignTop             Qt__AlignmentFlag = Qt__AlignmentFlag(0x0020)
	Qt__AlignBottom          Qt__AlignmentFlag = Qt__AlignmentFlag(0x0040)
	Qt__AlignVCenter         Qt__AlignmentFlag = Qt__AlignmentFlag(0x0080)
	Qt__AlignBaseline        Qt__AlignmentFlag = Qt__AlignmentFlag(0x0100)
	Qt__AlignVertical_Mask   Qt__AlignmentFlag = Qt__AlignmentFlag(Qt__AlignTop | Qt__AlignBottom | Qt__AlignVCenter | Qt__AlignBaseline)
	Qt__AlignCenter          Qt__AlignmentFlag = Qt__AlignmentFlag(Qt__AlignVCenter | Qt__AlignHCenter)
)

// Qt::AnchorPoint
//
//go:generate stringer -type=Qt__AnchorPoint
type Qt__AnchorPoint int64

const (
	Qt__AnchorLeft             Qt__AnchorPoint = Qt__AnchorPoint(0)
	Qt__AnchorHorizontalCenter Qt__AnchorPoint = Qt__AnchorPoint(1)
	Qt__AnchorRight            Qt__AnchorPoint = Qt__AnchorPoint(2)
	Qt__AnchorTop              Qt__AnchorPoint = Qt__AnchorPoint(3)
	Qt__AnchorVerticalCenter   Qt__AnchorPoint = Qt__AnchorPoint(4)
	Qt__AnchorBottom           Qt__AnchorPoint = Qt__AnchorPoint(5)
)

// Qt::ApplicationAttribute
//
//go:generate stringer -type=Qt__ApplicationAttribute
type Qt__ApplicationAttribute int64

const (
	Qt__AA_ImmediateWidgetCreation                 Qt__ApplicationAttribute = Qt__ApplicationAttribute(0)
	Qt__AA_MSWindowsUseDirect3DByDefault           Qt__ApplicationAttribute = Qt__ApplicationAttribute(1)
	Qt__AA_DontShowIconsInMenus                    Qt__ApplicationAttribute = Qt__ApplicationAttribute(2)
	Qt__AA_NativeWindows                           Qt__ApplicationAttribute = Qt__ApplicationAttribute(3)
	Qt__AA_DontCreateNativeWidgetSiblings          Qt__ApplicationAttribute = Qt__ApplicationAttribute(4)
	Qt__AA_PluginApplication                       Qt__ApplicationAttribute = Qt__ApplicationAttribute(5)
	Qt__AA_MacPluginApplication                    Qt__ApplicationAttribute = Qt__ApplicationAttribute(Qt__AA_PluginApplication)
	Qt__AA_DontUseNativeMenuBar                    Qt__ApplicationAttribute = Qt__ApplicationAttribute(6)
	Qt__AA_MacDontSwapCtrlAndMeta                  Qt__ApplicationAttribute = Qt__ApplicationAttribute(7)
	Qt__AA_Use96Dpi                                Qt__ApplicationAttribute = Qt__ApplicationAttribute(8)
	Qt__AA_DisableNativeVirtualKeyboard            Qt__ApplicationAttribute = Qt__ApplicationAttribute(9)
	Qt__AA_X11InitThreads                          Qt__ApplicationAttribute = Qt__ApplicationAttribute(10)
	Qt__AA_SynthesizeTouchForUnhandledMouseEvents  Qt__ApplicationAttribute = Qt__ApplicationAttribute(11)
	Qt__AA_SynthesizeMouseForUnhandledTouchEvents  Qt__ApplicationAttribute = Qt__ApplicationAttribute(12)
	Qt__AA_UseHighDpiPixmaps                       Qt__ApplicationAttribute = Qt__ApplicationAttribute(13)
	Qt__AA_ForceRasterWidgets                      Qt__ApplicationAttribute = Qt__ApplicationAttribute(14)
	Qt__AA_UseDesktopOpenGL                        Qt__ApplicationAttribute = Qt__ApplicationAttribute(15)
	Qt__AA_UseOpenGLES                             Qt__ApplicationAttribute = Qt__ApplicationAttribute(16)
	Qt__AA_UseSoftwareOpenGL                       Qt__ApplicationAttribute = Qt__ApplicationAttribute(17)
	Qt__AA_ShareOpenGLContexts                     Qt__ApplicationAttribute = Qt__ApplicationAttribute(18)
	Qt__AA_SetPalette                              Qt__ApplicationAttribute = Qt__ApplicationAttribute(19)
	Qt__AA_EnableHighDpiScaling                    Qt__ApplicationAttribute = Qt__ApplicationAttribute(20)
	Qt__AA_DisableHighDpiScaling                   Qt__ApplicationAttribute = Qt__ApplicationAttribute(21)
	Qt__AA_UseStyleSheetPropagationInWidgetStyles  Qt__ApplicationAttribute = Qt__ApplicationAttribute(22)
	Qt__AA_DontUseNativeDialogs                    Qt__ApplicationAttribute = Qt__ApplicationAttribute(23)
	Qt__AA_SynthesizeMouseForUnhandledTabletEvents Qt__ApplicationAttribute = Qt__ApplicationAttribute(24)
	Qt__AA_CompressHighFrequencyEvents             Qt__ApplicationAttribute = Qt__ApplicationAttribute(25)
	Qt__AA_DontCheckOpenGLContextThreadAffinity    Qt__ApplicationAttribute = Qt__ApplicationAttribute(26)
	Qt__AA_DisableShaderDiskCache                  Qt__ApplicationAttribute = Qt__ApplicationAttribute(27)
	Qt__AA_DontShowShortcutsInContextMenus         Qt__ApplicationAttribute = Qt__ApplicationAttribute(28)
	Qt__AA_CompressTabletEvents                    Qt__ApplicationAttribute = Qt__ApplicationAttribute(29)
	Qt__AA_DisableWindowContextHelpButton          Qt__ApplicationAttribute = Qt__ApplicationAttribute(30)
	Qt__AA_DisableSessionManager                   Qt__ApplicationAttribute = Qt__ApplicationAttribute(31)
	Qt__AA_AttributeCount                          Qt__ApplicationAttribute = Qt__ApplicationAttribute(32)
)

// Qt::ApplicationState
//
//go:generate stringer -type=Qt__ApplicationState
type Qt__ApplicationState int64

const (
	Qt__ApplicationSuspended Qt__ApplicationState = Qt__ApplicationState(0x00000000)
	Qt__ApplicationHidden    Qt__ApplicationState = Qt__ApplicationState(0x00000001)
	Qt__ApplicationInactive  Qt__ApplicationState = Qt__ApplicationState(0x00000002)
	Qt__ApplicationActive    Qt__ApplicationState = Qt__ApplicationState(0x00000004)
)

// Qt::ArrowType
//
//go:generate stringer -type=Qt__ArrowType
type Qt__ArrowType int64

const (
	Qt__NoArrow    Qt__ArrowType = Qt__ArrowType(0)
	Qt__UpArrow    Qt__ArrowType = Qt__ArrowType(1)
	Qt__DownArrow  Qt__ArrowType = Qt__ArrowType(2)
	Qt__LeftArrow  Qt__ArrowType = Qt__ArrowType(3)
	Qt__RightArrow Qt__ArrowType = Qt__ArrowType(4)
)

// Qt::AspectRatioMode
//
//go:generate stringer -type=Qt__AspectRatioMode
type Qt__AspectRatioMode int64

const (
	Qt__IgnoreAspectRatio          Qt__AspectRatioMode = Qt__AspectRatioMode(0)
	Qt__KeepAspectRatio            Qt__AspectRatioMode = Qt__AspectRatioMode(1)
	Qt__KeepAspectRatioByExpanding Qt__AspectRatioMode = Qt__AspectRatioMode(2)
)

// Qt::Axis
//
//go:generate stringer -type=Qt__Axis
type Qt__Axis int64

const (
	Qt__XAxis Qt__Axis = Qt__Axis(0)
	Qt__YAxis Qt__Axis = Qt__Axis(1)
	Qt__ZAxis Qt__Axis = Qt__Axis(2)
)

// Qt::BGMode
//
//go:generate stringer -type=Qt__BGMode
type Qt__BGMode int64

const (
	Qt__TransparentMode Qt__BGMode = Qt__BGMode(0)
	Qt__OpaqueMode      Qt__BGMode = Qt__BGMode(1)
)

// Qt::BrushStyle
//
//go:generate stringer -type=Qt__BrushStyle
type Qt__BrushStyle int64

var (
	Qt__NoBrush                Qt__BrushStyle = Qt__BrushStyle(0)
	Qt__SolidPattern           Qt__BrushStyle = Qt__BrushStyle(1)
	Qt__Dense1Pattern          Qt__BrushStyle = Qt__BrushStyle(2)
	Qt__Dense2Pattern          Qt__BrushStyle = Qt__BrushStyle(3)
	Qt__Dense3Pattern          Qt__BrushStyle = Qt__BrushStyle(4)
	Qt__Dense4Pattern          Qt__BrushStyle = Qt__BrushStyle(5)
	Qt__Dense5Pattern          Qt__BrushStyle = Qt__BrushStyle(6)
	Qt__Dense6Pattern          Qt__BrushStyle = Qt__BrushStyle(7)
	Qt__Dense7Pattern          Qt__BrushStyle = Qt__BrushStyle(8)
	Qt__HorPattern             Qt__BrushStyle = Qt__BrushStyle(9)
	Qt__VerPattern             Qt__BrushStyle = Qt__BrushStyle(10)
	Qt__CrossPattern           Qt__BrushStyle = Qt__BrushStyle(11)
	Qt__BDiagPattern           Qt__BrushStyle = Qt__BrushStyle(12)
	Qt__FDiagPattern           Qt__BrushStyle = Qt__BrushStyle(13)
	Qt__DiagCrossPattern       Qt__BrushStyle = Qt__BrushStyle(14)
	Qt__LinearGradientPattern  Qt__BrushStyle = Qt__BrushStyle(15)
	Qt__RadialGradientPattern  Qt__BrushStyle = Qt__BrushStyle(16)
	Qt__ConicalGradientPattern Qt__BrushStyle = Qt__BrushStyle(17)
	Qt__TexturePattern         Qt__BrushStyle = Qt__BrushStyle(24)
)

// Qt::CaseSensitivity
//
//go:generate stringer -type=Qt__CaseSensitivity
type Qt__CaseSensitivity int64

const (
	Qt__CaseInsensitive Qt__CaseSensitivity = Qt__CaseSensitivity(0)
	Qt__CaseSensitive   Qt__CaseSensitivity = Qt__CaseSensitivity(1)
)

// Qt::CheckState
//
//go:generate stringer -type=Qt__CheckState
type Qt__CheckState int64

const (
	Qt__Unchecked        Qt__CheckState = Qt__CheckState(0)
	Qt__PartiallyChecked Qt__CheckState = Qt__CheckState(1)
	Qt__Checked          Qt__CheckState = Qt__CheckState(2)
)

// Qt::ChecksumType
//
//go:generate stringer -type=Qt__ChecksumType
type Qt__ChecksumType int64

const (
	Qt__ChecksumIso3309 Qt__ChecksumType = Qt__ChecksumType(0)
	Qt__ChecksumItuV41  Qt__ChecksumType = Qt__ChecksumType(1)
)

// Qt::ClipOperation
//
//go:generate stringer -type=Qt__ClipOperation
type Qt__ClipOperation int64

const (
	Qt__NoClip        Qt__ClipOperation = Qt__ClipOperation(0)
	Qt__ReplaceClip   Qt__ClipOperation = Qt__ClipOperation(1)
	Qt__IntersectClip Qt__ClipOperation = Qt__ClipOperation(2)
)

// Qt::ConnectionType
//
//go:generate stringer -type=Qt__ConnectionType
type Qt__ConnectionType int64

const (
	Qt__AutoConnection           Qt__ConnectionType = Qt__ConnectionType(0)
	Qt__DirectConnection         Qt__ConnectionType = Qt__ConnectionType(1)
	Qt__QueuedConnection         Qt__ConnectionType = Qt__ConnectionType(2)
	Qt__BlockingQueuedConnection Qt__ConnectionType = Qt__ConnectionType(3)
	Qt__UniqueConnection         Qt__ConnectionType = Qt__ConnectionType(0x80)
)

// Qt::ContextMenuPolicy
//
//go:generate stringer -type=Qt__ContextMenuPolicy
type Qt__ContextMenuPolicy int64

const (
	Qt__NoContextMenu      Qt__ContextMenuPolicy = Qt__ContextMenuPolicy(0)
	Qt__DefaultContextMenu Qt__ContextMenuPolicy = Qt__ContextMenuPolicy(1)
	Qt__ActionsContextMenu Qt__ContextMenuPolicy = Qt__ContextMenuPolicy(2)
	Qt__CustomContextMenu  Qt__ContextMenuPolicy = Qt__ContextMenuPolicy(3)
	Qt__PreventContextMenu Qt__ContextMenuPolicy = Qt__ContextMenuPolicy(4)
)

// Qt::CoordinateSystem
//
//go:generate stringer -type=Qt__CoordinateSystem
type Qt__CoordinateSystem int64

const (
	Qt__DeviceCoordinates  Qt__CoordinateSystem = Qt__CoordinateSystem(0)
	Qt__LogicalCoordinates Qt__CoordinateSystem = Qt__CoordinateSystem(1)
)

// Qt::Corner
//
//go:generate stringer -type=Qt__Corner
type Qt__Corner int64

const (
	Qt__TopLeftCorner     Qt__Corner = Qt__Corner(0x00000)
	Qt__TopRightCorner    Qt__Corner = Qt__Corner(0x00001)
	Qt__BottomLeftCorner  Qt__Corner = Qt__Corner(0x00002)
	Qt__BottomRightCorner Qt__Corner = Qt__Corner(0x00003)
)

// Qt::CursorMoveStyle
//
//go:generate stringer -type=Qt__CursorMoveStyle
type Qt__CursorMoveStyle int64

var (
	Qt__LogicalMoveStyle Qt__CursorMoveStyle = Qt__CursorMoveStyle(0)
	Qt__VisualMoveStyle  Qt__CursorMoveStyle = Qt__CursorMoveStyle(1)
)

// Qt::CursorShape
//
//go:generate stringer -type=Qt__CursorShape
type Qt__CursorShape int64

const (
	Qt__ArrowCursor        Qt__CursorShape = Qt__CursorShape(0)
	Qt__UpArrowCursor      Qt__CursorShape = Qt__CursorShape(1)
	Qt__CrossCursor        Qt__CursorShape = Qt__CursorShape(2)
	Qt__WaitCursor         Qt__CursorShape = Qt__CursorShape(3)
	Qt__IBeamCursor        Qt__CursorShape = Qt__CursorShape(4)
	Qt__SizeVerCursor      Qt__CursorShape = Qt__CursorShape(5)
	Qt__SizeHorCursor      Qt__CursorShape = Qt__CursorShape(6)
	Qt__SizeBDiagCursor    Qt__CursorShape = Qt__CursorShape(7)
	Qt__SizeFDiagCursor    Qt__CursorShape = Qt__CursorShape(8)
	Qt__SizeAllCursor      Qt__CursorShape = Qt__CursorShape(9)
	Qt__BlankCursor        Qt__CursorShape = Qt__CursorShape(10)
	Qt__SplitVCursor       Qt__CursorShape = Qt__CursorShape(11)
	Qt__SplitHCursor       Qt__CursorShape = Qt__CursorShape(12)
	Qt__PointingHandCursor Qt__CursorShape = Qt__CursorShape(13)
	Qt__ForbiddenCursor    Qt__CursorShape = Qt__CursorShape(14)
	Qt__WhatsThisCursor    Qt__CursorShape = Qt__CursorShape(15)
	Qt__BusyCursor         Qt__CursorShape = Qt__CursorShape(16)
	Qt__OpenHandCursor     Qt__CursorShape = Qt__CursorShape(17)
	Qt__ClosedHandCursor   Qt__CursorShape = Qt__CursorShape(18)
	Qt__DragCopyCursor     Qt__CursorShape = Qt__CursorShape(19)
	Qt__DragMoveCursor     Qt__CursorShape = Qt__CursorShape(20)
	Qt__DragLinkCursor     Qt__CursorShape = Qt__CursorShape(21)
	Qt__LastCursor         Qt__CursorShape = Qt__CursorShape(Qt__DragLinkCursor)
	Qt__BitmapCursor       Qt__CursorShape = Qt__CursorShape(24)
	Qt__CustomCursor       Qt__CursorShape = Qt__CursorShape(25)
)

// Qt::DateFormat
//
//go:generate stringer -type=Qt__DateFormat
type Qt__DateFormat int64

const (
	Qt__TextDate               Qt__DateFormat = Qt__DateFormat(0)
	Qt__ISODate                Qt__DateFormat = Qt__DateFormat(1)
	Qt__SystemLocaleDate       Qt__DateFormat = Qt__DateFormat(2)
	Qt__LocalDate              Qt__DateFormat = Qt__DateFormat(2)
	Qt__LocaleDate             Qt__DateFormat = Qt__DateFormat(3)
	Qt__SystemLocaleShortDate  Qt__DateFormat = Qt__DateFormat(4)
	Qt__SystemLocaleLongDate   Qt__DateFormat = Qt__DateFormat(5)
	Qt__DefaultLocaleShortDate Qt__DateFormat = Qt__DateFormat(6)
	Qt__DefaultLocaleLongDate  Qt__DateFormat = Qt__DateFormat(7)
	Qt__RFC2822Date            Qt__DateFormat = Qt__DateFormat(8)
	Qt__ISODateWithMs          Qt__DateFormat = Qt__DateFormat(9)
)

// Qt::DayOfWeek
//
//go:generate stringer -type=Qt__DayOfWeek
type Qt__DayOfWeek int64

const (
	Qt__Monday    Qt__DayOfWeek = Qt__DayOfWeek(1)
	Qt__Tuesday   Qt__DayOfWeek = Qt__DayOfWeek(2)
	Qt__Wednesday Qt__DayOfWeek = Qt__DayOfWeek(3)
	Qt__Thursday  Qt__DayOfWeek = Qt__DayOfWeek(4)
	Qt__Friday    Qt__DayOfWeek = Qt__DayOfWeek(5)
	Qt__Saturday  Qt__DayOfWeek = Qt__DayOfWeek(6)
	Qt__Sunday    Qt__DayOfWeek = Qt__DayOfWeek(7)
)

// Qt::DockWidgetArea
//
//go:generate stringer -type=Qt__DockWidgetArea
type Qt__DockWidgetArea int64

const (
	Qt__LeftDockWidgetArea   Qt__DockWidgetArea = Qt__DockWidgetArea(0x1)
	Qt__RightDockWidgetArea  Qt__DockWidgetArea = Qt__DockWidgetArea(0x2)
	Qt__TopDockWidgetArea    Qt__DockWidgetArea = Qt__DockWidgetArea(0x4)
	Qt__BottomDockWidgetArea Qt__DockWidgetArea = Qt__DockWidgetArea(0x8)
	Qt__DockWidgetArea_Mask  Qt__DockWidgetArea = Qt__DockWidgetArea(0xf)
	Qt__AllDockWidgetAreas   Qt__DockWidgetArea = Qt__DockWidgetArea(Qt__DockWidgetArea_Mask)
	Qt__NoDockWidgetArea     Qt__DockWidgetArea = Qt__DockWidgetArea(0)
)

// Qt::DropAction
//
//go:generate stringer -type=Qt__DropAction
type Qt__DropAction int64

const (
	Qt__CopyAction       Qt__DropAction = Qt__DropAction(0x1)
	Qt__MoveAction       Qt__DropAction = Qt__DropAction(0x2)
	Qt__LinkAction       Qt__DropAction = Qt__DropAction(0x4)
	Qt__ActionMask       Qt__DropAction = Qt__DropAction(0xff)
	Qt__TargetMoveAction Qt__DropAction = Qt__DropAction(0x8002)
	Qt__IgnoreAction     Qt__DropAction = Qt__DropAction(0x0)
)

// Qt::Edge
//
//go:generate stringer -type=Qt__Edge
type Qt__Edge int64

const (
	Qt__TopEdge    Qt__Edge = Qt__Edge(0x00001)
	Qt__LeftEdge   Qt__Edge = Qt__Edge(0x00002)
	Qt__RightEdge  Qt__Edge = Qt__Edge(0x00004)
	Qt__BottomEdge Qt__Edge = Qt__Edge(0x00008)
)

// Qt::EnterKeyType
//
//go:generate stringer -type=Qt__EnterKeyType
type Qt__EnterKeyType int64

const (
	Qt__EnterKeyDefault  Qt__EnterKeyType = Qt__EnterKeyType(0)
	Qt__EnterKeyReturn   Qt__EnterKeyType = Qt__EnterKeyType(1)
	Qt__EnterKeyDone     Qt__EnterKeyType = Qt__EnterKeyType(2)
	Qt__EnterKeyGo       Qt__EnterKeyType = Qt__EnterKeyType(3)
	Qt__EnterKeySend     Qt__EnterKeyType = Qt__EnterKeyType(4)
	Qt__EnterKeySearch   Qt__EnterKeyType = Qt__EnterKeyType(5)
	Qt__EnterKeyNext     Qt__EnterKeyType = Qt__EnterKeyType(6)
	Qt__EnterKeyPrevious Qt__EnterKeyType = Qt__EnterKeyType(7)
)

// Qt::EventPriority
//
//go:generate stringer -type=Qt__EventPriority
type Qt__EventPriority int64

const (
	Qt__HighEventPriority   Qt__EventPriority = Qt__EventPriority(1)
	Qt__NormalEventPriority Qt__EventPriority = Qt__EventPriority(0)
	Qt__LowEventPriority    Qt__EventPriority = Qt__EventPriority(-1)
)

// Qt::FillRule
//
//go:generate stringer -type=Qt__FillRule
type Qt__FillRule int64

const (
	Qt__OddEvenFill Qt__FillRule = Qt__FillRule(0)
	Qt__WindingFill Qt__FillRule = Qt__FillRule(1)
)

// Qt::FindChildOption
//
//go:generate stringer -type=Qt__FindChildOption
type Qt__FindChildOption int64

const (
	Qt__FindDirectChildrenOnly  Qt__FindChildOption = Qt__FindChildOption(0x0)
	Qt__FindChildrenRecursively Qt__FindChildOption = Qt__FindChildOption(0x1)
)

// Qt::FocusPolicy
//
//go:generate stringer -type=Qt__FocusPolicy
type Qt__FocusPolicy int64

const (
	Qt__NoFocus     Qt__FocusPolicy = Qt__FocusPolicy(0)
	Qt__TabFocus    Qt__FocusPolicy = Qt__FocusPolicy(0x1)
	Qt__ClickFocus  Qt__FocusPolicy = Qt__FocusPolicy(0x2)
	Qt__StrongFocus Qt__FocusPolicy = Qt__FocusPolicy(Qt__TabFocus | Qt__ClickFocus | 0x8)
	Qt__WheelFocus  Qt__FocusPolicy = Qt__FocusPolicy(Qt__StrongFocus | 0x4)
)

// Qt::FocusReason
//
//go:generate stringer -type=Qt__FocusReason
type Qt__FocusReason int64

const (
	Qt__MouseFocusReason        Qt__FocusReason = Qt__FocusReason(0)
	Qt__TabFocusReason          Qt__FocusReason = Qt__FocusReason(1)
	Qt__BacktabFocusReason      Qt__FocusReason = Qt__FocusReason(2)
	Qt__ActiveWindowFocusReason Qt__FocusReason = Qt__FocusReason(3)
	Qt__PopupFocusReason        Qt__FocusReason = Qt__FocusReason(4)
	Qt__ShortcutFocusReason     Qt__FocusReason = Qt__FocusReason(5)
	Qt__MenuBarFocusReason      Qt__FocusReason = Qt__FocusReason(6)
	Qt__OtherFocusReason        Qt__FocusReason = Qt__FocusReason(7)
	Qt__NoFocusReason           Qt__FocusReason = Qt__FocusReason(8)
)

// Qt::GestureFlag
//
//go:generate stringer -type=Qt__GestureFlag
type Qt__GestureFlag int64

const (
	Qt__DontStartGestureOnChildren       Qt__GestureFlag = Qt__GestureFlag(0x01)
	Qt__ReceivePartialGestures           Qt__GestureFlag = Qt__GestureFlag(0x02)
	Qt__IgnoredGesturesPropagateToParent Qt__GestureFlag = Qt__GestureFlag(0x04)
)

// Qt::GestureState
//
//go:generate stringer -type=Qt__GestureState
type Qt__GestureState int64

const (
	Qt__NoGesture       Qt__GestureState = Qt__GestureState(0)
	Qt__GestureStarted  Qt__GestureState = Qt__GestureState(1)
	Qt__GestureUpdated  Qt__GestureState = Qt__GestureState(2)
	Qt__GestureFinished Qt__GestureState = Qt__GestureState(3)
	Qt__GestureCanceled Qt__GestureState = Qt__GestureState(4)
)

// Qt::GestureType
//
//go:generate stringer -type=Qt__GestureType
type Qt__GestureType int64

var (
	Qt__TapGesture        Qt__GestureType = Qt__GestureType(1)
	Qt__TapAndHoldGesture Qt__GestureType = Qt__GestureType(2)
	Qt__PanGesture        Qt__GestureType = Qt__GestureType(3)
	Qt__PinchGesture      Qt__GestureType = Qt__GestureType(4)
	Qt__SwipeGesture      Qt__GestureType = Qt__GestureType(5)
	Qt__CustomGesture     Qt__GestureType = Qt__GestureType(0x0100)
	Qt__LastGestureType   Qt__GestureType = Qt__GestureType(C.Qt_LastGestureType_Type())
)

// Qt::GlobalColor
//
//go:generate stringer -type=Qt__GlobalColor
type Qt__GlobalColor int64

const (
	Qt__color0      Qt__GlobalColor = Qt__GlobalColor(0)
	Qt__color1      Qt__GlobalColor = Qt__GlobalColor(1)
	Qt__black       Qt__GlobalColor = Qt__GlobalColor(2)
	Qt__white       Qt__GlobalColor = Qt__GlobalColor(3)
	Qt__darkGray    Qt__GlobalColor = Qt__GlobalColor(4)
	Qt__gray        Qt__GlobalColor = Qt__GlobalColor(5)
	Qt__lightGray   Qt__GlobalColor = Qt__GlobalColor(6)
	Qt__red         Qt__GlobalColor = Qt__GlobalColor(7)
	Qt__green       Qt__GlobalColor = Qt__GlobalColor(8)
	Qt__blue        Qt__GlobalColor = Qt__GlobalColor(9)
	Qt__cyan        Qt__GlobalColor = Qt__GlobalColor(10)
	Qt__magenta     Qt__GlobalColor = Qt__GlobalColor(11)
	Qt__yellow      Qt__GlobalColor = Qt__GlobalColor(12)
	Qt__darkRed     Qt__GlobalColor = Qt__GlobalColor(13)
	Qt__darkGreen   Qt__GlobalColor = Qt__GlobalColor(14)
	Qt__darkBlue    Qt__GlobalColor = Qt__GlobalColor(15)
	Qt__darkCyan    Qt__GlobalColor = Qt__GlobalColor(16)
	Qt__darkMagenta Qt__GlobalColor = Qt__GlobalColor(17)
	Qt__darkYellow  Qt__GlobalColor = Qt__GlobalColor(18)
	Qt__transparent Qt__GlobalColor = Qt__GlobalColor(19)
)

// Qt::HighDpiScaleFactorRoundingPolicy
//
//go:generate stringer -type=Qt__HighDpiScaleFactorRoundingPolicy
type Qt__HighDpiScaleFactorRoundingPolicy int64

const (
	Qt__Unset            Qt__HighDpiScaleFactorRoundingPolicy = Qt__HighDpiScaleFactorRoundingPolicy(0)
	Qt__Round            Qt__HighDpiScaleFactorRoundingPolicy = Qt__HighDpiScaleFactorRoundingPolicy(1)
	Qt__Ceil             Qt__HighDpiScaleFactorRoundingPolicy = Qt__HighDpiScaleFactorRoundingPolicy(2)
	Qt__Floor            Qt__HighDpiScaleFactorRoundingPolicy = Qt__HighDpiScaleFactorRoundingPolicy(3)
	Qt__RoundPreferFloor Qt__HighDpiScaleFactorRoundingPolicy = Qt__HighDpiScaleFactorRoundingPolicy(4)
	Qt__PassThrough      Qt__HighDpiScaleFactorRoundingPolicy = Qt__HighDpiScaleFactorRoundingPolicy(5)
)

// Qt::HitTestAccuracy
//
//go:generate stringer -type=Qt__HitTestAccuracy
type Qt__HitTestAccuracy int64

const (
	Qt__ExactHit Qt__HitTestAccuracy = Qt__HitTestAccuracy(0)
	Qt__FuzzyHit Qt__HitTestAccuracy = Qt__HitTestAccuracy(1)
)

// Qt::ImageConversionFlag
//
//go:generate stringer -type=Qt__ImageConversionFlag
type Qt__ImageConversionFlag int64

const (
	Qt__ColorMode_Mask       Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000003)
	Qt__AutoColor            Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000000)
	Qt__ColorOnly            Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000003)
	Qt__MonoOnly             Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000002)
	Qt__AlphaDither_Mask     Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x0000000c)
	Qt__ThresholdAlphaDither Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000000)
	Qt__OrderedAlphaDither   Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000004)
	Qt__DiffuseAlphaDither   Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000008)
	Qt__NoAlpha              Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x0000000c)
	Qt__Dither_Mask          Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000030)
	Qt__DiffuseDither        Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000000)
	Qt__OrderedDither        Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000010)
	Qt__ThresholdDither      Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000020)
	Qt__DitherMode_Mask      Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x000000c0)
	Qt__AutoDither           Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000000)
	Qt__PreferDither         Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000040)
	Qt__AvoidDither          Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000080)
	Qt__NoOpaqueDetection    Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000100)
	Qt__NoFormatConversion   Qt__ImageConversionFlag = Qt__ImageConversionFlag(0x00000200)
)

// Qt::InputMethodHint
//
//go:generate stringer -type=Qt__InputMethodHint
type Qt__InputMethodHint int64

const (
	Qt__ImhNone                   Qt__InputMethodHint = Qt__InputMethodHint(0x0)
	Qt__ImhHiddenText             Qt__InputMethodHint = Qt__InputMethodHint(0x1)
	Qt__ImhSensitiveData          Qt__InputMethodHint = Qt__InputMethodHint(0x2)
	Qt__ImhNoAutoUppercase        Qt__InputMethodHint = Qt__InputMethodHint(0x4)
	Qt__ImhPreferNumbers          Qt__InputMethodHint = Qt__InputMethodHint(0x8)
	Qt__ImhPreferUppercase        Qt__InputMethodHint = Qt__InputMethodHint(0x10)
	Qt__ImhPreferLowercase        Qt__InputMethodHint = Qt__InputMethodHint(0x20)
	Qt__ImhNoPredictiveText       Qt__InputMethodHint = Qt__InputMethodHint(0x40)
	Qt__ImhDate                   Qt__InputMethodHint = Qt__InputMethodHint(0x80)
	Qt__ImhTime                   Qt__InputMethodHint = Qt__InputMethodHint(0x100)
	Qt__ImhPreferLatin            Qt__InputMethodHint = Qt__InputMethodHint(0x200)
	Qt__ImhMultiLine              Qt__InputMethodHint = Qt__InputMethodHint(0x400)
	Qt__ImhNoEditMenu             Qt__InputMethodHint = Qt__InputMethodHint(0x800)
	Qt__ImhNoTextHandles          Qt__InputMethodHint = Qt__InputMethodHint(0x1000)
	Qt__ImhDigitsOnly             Qt__InputMethodHint = Qt__InputMethodHint(0x10000)
	Qt__ImhFormattedNumbersOnly   Qt__InputMethodHint = Qt__InputMethodHint(0x20000)
	Qt__ImhUppercaseOnly          Qt__InputMethodHint = Qt__InputMethodHint(0x40000)
	Qt__ImhLowercaseOnly          Qt__InputMethodHint = Qt__InputMethodHint(0x80000)
	Qt__ImhDialableCharactersOnly Qt__InputMethodHint = Qt__InputMethodHint(0x100000)
	Qt__ImhEmailCharactersOnly    Qt__InputMethodHint = Qt__InputMethodHint(0x200000)
	Qt__ImhUrlCharactersOnly      Qt__InputMethodHint = Qt__InputMethodHint(0x400000)
	Qt__ImhLatinOnly              Qt__InputMethodHint = Qt__InputMethodHint(0x800000)
	Qt__ImhExclusiveInputMask     Qt__InputMethodHint = Qt__InputMethodHint(0xffff0000)
)

// Qt::InputMethodQuery
//
//go:generate stringer -type=Qt__InputMethodQuery
type Qt__InputMethodQuery int64

const (
	Qt__ImEnabled                Qt__InputMethodQuery = Qt__InputMethodQuery(0x1)
	Qt__ImCursorRectangle        Qt__InputMethodQuery = Qt__InputMethodQuery(0x2)
	Qt__ImMicroFocus             Qt__InputMethodQuery = Qt__InputMethodQuery(0x2)
	Qt__ImFont                   Qt__InputMethodQuery = Qt__InputMethodQuery(0x4)
	Qt__ImCursorPosition         Qt__InputMethodQuery = Qt__InputMethodQuery(0x8)
	Qt__ImSurroundingText        Qt__InputMethodQuery = Qt__InputMethodQuery(0x10)
	Qt__ImCurrentSelection       Qt__InputMethodQuery = Qt__InputMethodQuery(0x20)
	Qt__ImMaximumTextLength      Qt__InputMethodQuery = Qt__InputMethodQuery(0x40)
	Qt__ImAnchorPosition         Qt__InputMethodQuery = Qt__InputMethodQuery(0x80)
	Qt__ImHints                  Qt__InputMethodQuery = Qt__InputMethodQuery(0x100)
	Qt__ImPreferredLanguage      Qt__InputMethodQuery = Qt__InputMethodQuery(0x200)
	Qt__ImAbsolutePosition       Qt__InputMethodQuery = Qt__InputMethodQuery(0x400)
	Qt__ImTextBeforeCursor       Qt__InputMethodQuery = Qt__InputMethodQuery(0x800)
	Qt__ImTextAfterCursor        Qt__InputMethodQuery = Qt__InputMethodQuery(0x1000)
	Qt__ImEnterKeyType           Qt__InputMethodQuery = Qt__InputMethodQuery(0x2000)
	Qt__ImAnchorRectangle        Qt__InputMethodQuery = Qt__InputMethodQuery(0x4000)
	Qt__ImInputItemClipRectangle Qt__InputMethodQuery = Qt__InputMethodQuery(0x8000)
	Qt__ImPlatformData           Qt__InputMethodQuery = Qt__InputMethodQuery(0x80000000)
	Qt__ImQueryAll               Qt__InputMethodQuery = Qt__InputMethodQuery(0xffffffff)
)

// Qt::ItemDataRole
//
//go:generate stringer -type=Qt__ItemDataRole
type Qt__ItemDataRole int64

const (
	Qt__DisplayRole               Qt__ItemDataRole = Qt__ItemDataRole(0)
	Qt__DecorationRole            Qt__ItemDataRole = Qt__ItemDataRole(1)
	Qt__EditRole                  Qt__ItemDataRole = Qt__ItemDataRole(2)
	Qt__ToolTipRole               Qt__ItemDataRole = Qt__ItemDataRole(3)
	Qt__StatusTipRole             Qt__ItemDataRole = Qt__ItemDataRole(4)
	Qt__WhatsThisRole             Qt__ItemDataRole = Qt__ItemDataRole(5)
	Qt__FontRole                  Qt__ItemDataRole = Qt__ItemDataRole(6)
	Qt__TextAlignmentRole         Qt__ItemDataRole = Qt__ItemDataRole(7)
	Qt__BackgroundRole            Qt__ItemDataRole = Qt__ItemDataRole(8)
	Qt__ForegroundRole            Qt__ItemDataRole = Qt__ItemDataRole(9)
	Qt__BackgroundColorRole       Qt__ItemDataRole = Qt__ItemDataRole(Qt__BackgroundRole)
	Qt__TextColorRole             Qt__ItemDataRole = Qt__ItemDataRole(Qt__ForegroundRole)
	Qt__CheckStateRole            Qt__ItemDataRole = Qt__ItemDataRole(10)
	Qt__AccessibleTextRole        Qt__ItemDataRole = Qt__ItemDataRole(11)
	Qt__AccessibleDescriptionRole Qt__ItemDataRole = Qt__ItemDataRole(12)
	Qt__SizeHintRole              Qt__ItemDataRole = Qt__ItemDataRole(13)
	Qt__InitialSortOrderRole      Qt__ItemDataRole = Qt__ItemDataRole(14)
	Qt__DisplayPropertyRole       Qt__ItemDataRole = Qt__ItemDataRole(27)
	Qt__DecorationPropertyRole    Qt__ItemDataRole = Qt__ItemDataRole(28)
	Qt__ToolTipPropertyRole       Qt__ItemDataRole = Qt__ItemDataRole(29)
	Qt__StatusTipPropertyRole     Qt__ItemDataRole = Qt__ItemDataRole(30)
	Qt__WhatsThisPropertyRole     Qt__ItemDataRole = Qt__ItemDataRole(31)
	Qt__UserRole                  Qt__ItemDataRole = Qt__ItemDataRole(0x0100)
)

// Qt::ItemFlag
//
//go:generate stringer -type=Qt__ItemFlag
type Qt__ItemFlag int64

const (
	Qt__NoItemFlags          Qt__ItemFlag = Qt__ItemFlag(0)
	Qt__ItemIsSelectable     Qt__ItemFlag = Qt__ItemFlag(1)
	Qt__ItemIsEditable       Qt__ItemFlag = Qt__ItemFlag(2)
	Qt__ItemIsDragEnabled    Qt__ItemFlag = Qt__ItemFlag(4)
	Qt__ItemIsDropEnabled    Qt__ItemFlag = Qt__ItemFlag(8)
	Qt__ItemIsUserCheckable  Qt__ItemFlag = Qt__ItemFlag(16)
	Qt__ItemIsEnabled        Qt__ItemFlag = Qt__ItemFlag(32)
	Qt__ItemIsAutoTristate   Qt__ItemFlag = Qt__ItemFlag(64)
	Qt__ItemIsTristate       Qt__ItemFlag = Qt__ItemFlag(Qt__ItemIsAutoTristate)
	Qt__ItemNeverHasChildren Qt__ItemFlag = Qt__ItemFlag(128)
	Qt__ItemIsUserTristate   Qt__ItemFlag = Qt__ItemFlag(256)
)

// Qt::ItemSelectionMode
//
//go:generate stringer -type=Qt__ItemSelectionMode
type Qt__ItemSelectionMode int64

const (
	Qt__ContainsItemShape          Qt__ItemSelectionMode = Qt__ItemSelectionMode(0x0)
	Qt__IntersectsItemShape        Qt__ItemSelectionMode = Qt__ItemSelectionMode(0x1)
	Qt__ContainsItemBoundingRect   Qt__ItemSelectionMode = Qt__ItemSelectionMode(0x2)
	Qt__IntersectsItemBoundingRect Qt__ItemSelectionMode = Qt__ItemSelectionMode(0x3)
)

// Qt::ItemSelectionOperation
//
//go:generate stringer -type=Qt__ItemSelectionOperation
type Qt__ItemSelectionOperation int64

const (
	Qt__ReplaceSelection Qt__ItemSelectionOperation = Qt__ItemSelectionOperation(0)
	Qt__AddToSelection   Qt__ItemSelectionOperation = Qt__ItemSelectionOperation(1)
)

// Qt::Key
//
//go:generate stringer -type=Qt__Key
type Qt__Key int64

const (
	Qt__Key_Escape                  Qt__Key = Qt__Key(0x01000000)
	Qt__Key_Tab                     Qt__Key = Qt__Key(0x01000001)
	Qt__Key_Backtab                 Qt__Key = Qt__Key(0x01000002)
	Qt__Key_Backspace               Qt__Key = Qt__Key(0x01000003)
	Qt__Key_Return                  Qt__Key = Qt__Key(0x01000004)
	Qt__Key_Enter                   Qt__Key = Qt__Key(0x01000005)
	Qt__Key_Insert                  Qt__Key = Qt__Key(0x01000006)
	Qt__Key_Delete                  Qt__Key = Qt__Key(0x01000007)
	Qt__Key_Pause                   Qt__Key = Qt__Key(0x01000008)
	Qt__Key_Print                   Qt__Key = Qt__Key(0x01000009)
	Qt__Key_SysReq                  Qt__Key = Qt__Key(0x0100000a)
	Qt__Key_Clear                   Qt__Key = Qt__Key(0x0100000b)
	Qt__Key_Home                    Qt__Key = Qt__Key(0x01000010)
	Qt__Key_End                     Qt__Key = Qt__Key(0x01000011)
	Qt__Key_Left                    Qt__Key = Qt__Key(0x01000012)
	Qt__Key_Up                      Qt__Key = Qt__Key(0x01000013)
	Qt__Key_Right                   Qt__Key = Qt__Key(0x01000014)
	Qt__Key_Down                    Qt__Key = Qt__Key(0x01000015)
	Qt__Key_PageUp                  Qt__Key = Qt__Key(0x01000016)
	Qt__Key_PageDown                Qt__Key = Qt__Key(0x01000017)
	Qt__Key_Shift                   Qt__Key = Qt__Key(0x01000020)
	Qt__Key_Control                 Qt__Key = Qt__Key(0x01000021)
	Qt__Key_Meta                    Qt__Key = Qt__Key(0x01000022)
	Qt__Key_Alt                     Qt__Key = Qt__Key(0x01000023)
	Qt__Key_CapsLock                Qt__Key = Qt__Key(0x01000024)
	Qt__Key_NumLock                 Qt__Key = Qt__Key(0x01000025)
	Qt__Key_ScrollLock              Qt__Key = Qt__Key(0x01000026)
	Qt__Key_F1                      Qt__Key = Qt__Key(0x01000030)
	Qt__Key_F2                      Qt__Key = Qt__Key(0x01000031)
	Qt__Key_F3                      Qt__Key = Qt__Key(0x01000032)
	Qt__Key_F4                      Qt__Key = Qt__Key(0x01000033)
	Qt__Key_F5                      Qt__Key = Qt__Key(0x01000034)
	Qt__Key_F6                      Qt__Key = Qt__Key(0x01000035)
	Qt__Key_F7                      Qt__Key = Qt__Key(0x01000036)
	Qt__Key_F8                      Qt__Key = Qt__Key(0x01000037)
	Qt__Key_F9                      Qt__Key = Qt__Key(0x01000038)
	Qt__Key_F10                     Qt__Key = Qt__Key(0x01000039)
	Qt__Key_F11                     Qt__Key = Qt__Key(0x0100003a)
	Qt__Key_F12                     Qt__Key = Qt__Key(0x0100003b)
	Qt__Key_F13                     Qt__Key = Qt__Key(0x0100003c)
	Qt__Key_F14                     Qt__Key = Qt__Key(0x0100003d)
	Qt__Key_F15                     Qt__Key = Qt__Key(0x0100003e)
	Qt__Key_F16                     Qt__Key = Qt__Key(0x0100003f)
	Qt__Key_F17                     Qt__Key = Qt__Key(0x01000040)
	Qt__Key_F18                     Qt__Key = Qt__Key(0x01000041)
	Qt__Key_F19                     Qt__Key = Qt__Key(0x01000042)
	Qt__Key_F20                     Qt__Key = Qt__Key(0x01000043)
	Qt__Key_F21                     Qt__Key = Qt__Key(0x01000044)
	Qt__Key_F22                     Qt__Key = Qt__Key(0x01000045)
	Qt__Key_F23                     Qt__Key = Qt__Key(0x01000046)
	Qt__Key_F24                     Qt__Key = Qt__Key(0x01000047)
	Qt__Key_F25                     Qt__Key = Qt__Key(0x01000048)
	Qt__Key_F26                     Qt__Key = Qt__Key(0x01000049)
	Qt__Key_F27                     Qt__Key = Qt__Key(0x0100004a)
	Qt__Key_F28                     Qt__Key = Qt__Key(0x0100004b)
	Qt__Key_F29                     Qt__Key = Qt__Key(0x0100004c)
	Qt__Key_F30                     Qt__Key = Qt__Key(0x0100004d)
	Qt__Key_F31                     Qt__Key = Qt__Key(0x0100004e)
	Qt__Key_F32                     Qt__Key = Qt__Key(0x0100004f)
	Qt__Key_F33                     Qt__Key = Qt__Key(0x01000050)
	Qt__Key_F34                     Qt__Key = Qt__Key(0x01000051)
	Qt__Key_F35                     Qt__Key = Qt__Key(0x01000052)
	Qt__Key_Super_L                 Qt__Key = Qt__Key(0x01000053)
	Qt__Key_Super_R                 Qt__Key = Qt__Key(0x01000054)
	Qt__Key_Menu                    Qt__Key = Qt__Key(0x01000055)
	Qt__Key_Hyper_L                 Qt__Key = Qt__Key(0x01000056)
	Qt__Key_Hyper_R                 Qt__Key = Qt__Key(0x01000057)
	Qt__Key_Help                    Qt__Key = Qt__Key(0x01000058)
	Qt__Key_Direction_L             Qt__Key = Qt__Key(0x01000059)
	Qt__Key_Direction_R             Qt__Key = Qt__Key(0x01000060)
	Qt__Key_Space                   Qt__Key = Qt__Key(0x20)
	Qt__Key_Any                     Qt__Key = Qt__Key(Qt__Key_Space)
	Qt__Key_Exclam                  Qt__Key = Qt__Key(0x21)
	Qt__Key_QuoteDbl                Qt__Key = Qt__Key(0x22)
	Qt__Key_NumberSign              Qt__Key = Qt__Key(0x23)
	Qt__Key_Dollar                  Qt__Key = Qt__Key(0x24)
	Qt__Key_Percent                 Qt__Key = Qt__Key(0x25)
	Qt__Key_Ampersand               Qt__Key = Qt__Key(0x26)
	Qt__Key_Apostrophe              Qt__Key = Qt__Key(0x27)
	Qt__Key_ParenLeft               Qt__Key = Qt__Key(0x28)
	Qt__Key_ParenRight              Qt__Key = Qt__Key(0x29)
	Qt__Key_Asterisk                Qt__Key = Qt__Key(0x2a)
	Qt__Key_Plus                    Qt__Key = Qt__Key(0x2b)
	Qt__Key_Comma                   Qt__Key = Qt__Key(0x2c)
	Qt__Key_Minus                   Qt__Key = Qt__Key(0x2d)
	Qt__Key_Period                  Qt__Key = Qt__Key(0x2e)
	Qt__Key_Slash                   Qt__Key = Qt__Key(0x2f)
	Qt__Key_0                       Qt__Key = Qt__Key(0x30)
	Qt__Key_1                       Qt__Key = Qt__Key(0x31)
	Qt__Key_2                       Qt__Key = Qt__Key(0x32)
	Qt__Key_3                       Qt__Key = Qt__Key(0x33)
	Qt__Key_4                       Qt__Key = Qt__Key(0x34)
	Qt__Key_5                       Qt__Key = Qt__Key(0x35)
	Qt__Key_6                       Qt__Key = Qt__Key(0x36)
	Qt__Key_7                       Qt__Key = Qt__Key(0x37)
	Qt__Key_8                       Qt__Key = Qt__Key(0x38)
	Qt__Key_9                       Qt__Key = Qt__Key(0x39)
	Qt__Key_Colon                   Qt__Key = Qt__Key(0x3a)
	Qt__Key_Semicolon               Qt__Key = Qt__Key(0x3b)
	Qt__Key_Less                    Qt__Key = Qt__Key(0x3c)
	Qt__Key_Equal                   Qt__Key = Qt__Key(0x3d)
	Qt__Key_Greater                 Qt__Key = Qt__Key(0x3e)
	Qt__Key_Question                Qt__Key = Qt__Key(0x3f)
	Qt__Key_At                      Qt__Key = Qt__Key(0x40)
	Qt__Key_A                       Qt__Key = Qt__Key(0x41)
	Qt__Key_B                       Qt__Key = Qt__Key(0x42)
	Qt__Key_C                       Qt__Key = Qt__Key(0x43)
	Qt__Key_D                       Qt__Key = Qt__Key(0x44)
	Qt__Key_E                       Qt__Key = Qt__Key(0x45)
	Qt__Key_F                       Qt__Key = Qt__Key(0x46)
	Qt__Key_G                       Qt__Key = Qt__Key(0x47)
	Qt__Key_H                       Qt__Key = Qt__Key(0x48)
	Qt__Key_I                       Qt__Key = Qt__Key(0x49)
	Qt__Key_J                       Qt__Key = Qt__Key(0x4a)
	Qt__Key_K                       Qt__Key = Qt__Key(0x4b)
	Qt__Key_L                       Qt__Key = Qt__Key(0x4c)
	Qt__Key_M                       Qt__Key = Qt__Key(0x4d)
	Qt__Key_N                       Qt__Key = Qt__Key(0x4e)
	Qt__Key_O                       Qt__Key = Qt__Key(0x4f)
	Qt__Key_P                       Qt__Key = Qt__Key(0x50)
	Qt__Key_Q                       Qt__Key = Qt__Key(0x51)
	Qt__Key_R                       Qt__Key = Qt__Key(0x52)
	Qt__Key_S                       Qt__Key = Qt__Key(0x53)
	Qt__Key_T                       Qt__Key = Qt__Key(0x54)
	Qt__Key_U                       Qt__Key = Qt__Key(0x55)
	Qt__Key_V                       Qt__Key = Qt__Key(0x56)
	Qt__Key_W                       Qt__Key = Qt__Key(0x57)
	Qt__Key_X                       Qt__Key = Qt__Key(0x58)
	Qt__Key_Y                       Qt__Key = Qt__Key(0x59)
	Qt__Key_Z                       Qt__Key = Qt__Key(0x5a)
	Qt__Key_BracketLeft             Qt__Key = Qt__Key(0x5b)
	Qt__Key_Backslash               Qt__Key = Qt__Key(0x5c)
	Qt__Key_BracketRight            Qt__Key = Qt__Key(0x5d)
	Qt__Key_AsciiCircum             Qt__Key = Qt__Key(0x5e)
	Qt__Key_Underscore              Qt__Key = Qt__Key(0x5f)
	Qt__Key_QuoteLeft               Qt__Key = Qt__Key(0x60)
	Qt__Key_BraceLeft               Qt__Key = Qt__Key(0x7b)
	Qt__Key_Bar                     Qt__Key = Qt__Key(0x7c)
	Qt__Key_BraceRight              Qt__Key = Qt__Key(0x7d)
	Qt__Key_AsciiTilde              Qt__Key = Qt__Key(0x7e)
	Qt__Key_nobreakspace            Qt__Key = Qt__Key(0x0a0)
	Qt__Key_exclamdown              Qt__Key = Qt__Key(0x0a1)
	Qt__Key_cent                    Qt__Key = Qt__Key(0x0a2)
	Qt__Key_sterling                Qt__Key = Qt__Key(0x0a3)
	Qt__Key_currency                Qt__Key = Qt__Key(0x0a4)
	Qt__Key_yen                     Qt__Key = Qt__Key(0x0a5)
	Qt__Key_brokenbar               Qt__Key = Qt__Key(0x0a6)
	Qt__Key_section                 Qt__Key = Qt__Key(0x0a7)
	Qt__Key_diaeresis               Qt__Key = Qt__Key(0x0a8)
	Qt__Key_copyright               Qt__Key = Qt__Key(0x0a9)
	Qt__Key_ordfeminine             Qt__Key = Qt__Key(0x0aa)
	Qt__Key_guillemotleft           Qt__Key = Qt__Key(0x0ab)
	Qt__Key_notsign                 Qt__Key = Qt__Key(0x0ac)
	Qt__Key_hyphen                  Qt__Key = Qt__Key(0x0ad)
	Qt__Key_registered              Qt__Key = Qt__Key(0x0ae)
	Qt__Key_macron                  Qt__Key = Qt__Key(0x0af)
	Qt__Key_degree                  Qt__Key = Qt__Key(0x0b0)
	Qt__Key_plusminus               Qt__Key = Qt__Key(0x0b1)
	Qt__Key_twosuperior             Qt__Key = Qt__Key(0x0b2)
	Qt__Key_threesuperior           Qt__Key = Qt__Key(0x0b3)
	Qt__Key_acute                   Qt__Key = Qt__Key(0x0b4)
	Qt__Key_mu                      Qt__Key = Qt__Key(0x0b5)
	Qt__Key_paragraph               Qt__Key = Qt__Key(0x0b6)
	Qt__Key_periodcentered          Qt__Key = Qt__Key(0x0b7)
	Qt__Key_cedilla                 Qt__Key = Qt__Key(0x0b8)
	Qt__Key_onesuperior             Qt__Key = Qt__Key(0x0b9)
	Qt__Key_masculine               Qt__Key = Qt__Key(0x0ba)
	Qt__Key_guillemotright          Qt__Key = Qt__Key(0x0bb)
	Qt__Key_onequarter              Qt__Key = Qt__Key(0x0bc)
	Qt__Key_onehalf                 Qt__Key = Qt__Key(0x0bd)
	Qt__Key_threequarters           Qt__Key = Qt__Key(0x0be)
	Qt__Key_questiondown            Qt__Key = Qt__Key(0x0bf)
	Qt__Key_Agrave                  Qt__Key = Qt__Key(0x0c0)
	Qt__Key_Aacute                  Qt__Key = Qt__Key(0x0c1)
	Qt__Key_Acircumflex             Qt__Key = Qt__Key(0x0c2)
	Qt__Key_Atilde                  Qt__Key = Qt__Key(0x0c3)
	Qt__Key_Adiaeresis              Qt__Key = Qt__Key(0x0c4)
	Qt__Key_Aring                   Qt__Key = Qt__Key(0x0c5)
	Qt__Key_AE                      Qt__Key = Qt__Key(0x0c6)
	Qt__Key_Ccedilla                Qt__Key = Qt__Key(0x0c7)
	Qt__Key_Egrave                  Qt__Key = Qt__Key(0x0c8)
	Qt__Key_Eacute                  Qt__Key = Qt__Key(0x0c9)
	Qt__Key_Ecircumflex             Qt__Key = Qt__Key(0x0ca)
	Qt__Key_Ediaeresis              Qt__Key = Qt__Key(0x0cb)
	Qt__Key_Igrave                  Qt__Key = Qt__Key(0x0cc)
	Qt__Key_Iacute                  Qt__Key = Qt__Key(0x0cd)
	Qt__Key_Icircumflex             Qt__Key = Qt__Key(0x0ce)
	Qt__Key_Idiaeresis              Qt__Key = Qt__Key(0x0cf)
	Qt__Key_ETH                     Qt__Key = Qt__Key(0x0d0)
	Qt__Key_Ntilde                  Qt__Key = Qt__Key(0x0d1)
	Qt__Key_Ograve                  Qt__Key = Qt__Key(0x0d2)
	Qt__Key_Oacute                  Qt__Key = Qt__Key(0x0d3)
	Qt__Key_Ocircumflex             Qt__Key = Qt__Key(0x0d4)
	Qt__Key_Otilde                  Qt__Key = Qt__Key(0x0d5)
	Qt__Key_Odiaeresis              Qt__Key = Qt__Key(0x0d6)
	Qt__Key_multiply                Qt__Key = Qt__Key(0x0d7)
	Qt__Key_Ooblique                Qt__Key = Qt__Key(0x0d8)
	Qt__Key_Ugrave                  Qt__Key = Qt__Key(0x0d9)
	Qt__Key_Uacute                  Qt__Key = Qt__Key(0x0da)
	Qt__Key_Ucircumflex             Qt__Key = Qt__Key(0x0db)
	Qt__Key_Udiaeresis              Qt__Key = Qt__Key(0x0dc)
	Qt__Key_Yacute                  Qt__Key = Qt__Key(0x0dd)
	Qt__Key_THORN                   Qt__Key = Qt__Key(0x0de)
	Qt__Key_ssharp                  Qt__Key = Qt__Key(0x0df)
	Qt__Key_division                Qt__Key = Qt__Key(0x0f7)
	Qt__Key_ydiaeresis              Qt__Key = Qt__Key(0x0ff)
	Qt__Key_AltGr                   Qt__Key = Qt__Key(0x01001103)
	Qt__Key_Multi_key               Qt__Key = Qt__Key(0x01001120)
	Qt__Key_Codeinput               Qt__Key = Qt__Key(0x01001137)
	Qt__Key_SingleCandidate         Qt__Key = Qt__Key(0x0100113c)
	Qt__Key_MultipleCandidate       Qt__Key = Qt__Key(0x0100113d)
	Qt__Key_PreviousCandidate       Qt__Key = Qt__Key(0x0100113e)
	Qt__Key_Mode_switch             Qt__Key = Qt__Key(0x0100117e)
	Qt__Key_Kanji                   Qt__Key = Qt__Key(0x01001121)
	Qt__Key_Muhenkan                Qt__Key = Qt__Key(0x01001122)
	Qt__Key_Henkan                  Qt__Key = Qt__Key(0x01001123)
	Qt__Key_Romaji                  Qt__Key = Qt__Key(0x01001124)
	Qt__Key_Hiragana                Qt__Key = Qt__Key(0x01001125)
	Qt__Key_Katakana                Qt__Key = Qt__Key(0x01001126)
	Qt__Key_Hiragana_Katakana       Qt__Key = Qt__Key(0x01001127)
	Qt__Key_Zenkaku                 Qt__Key = Qt__Key(0x01001128)
	Qt__Key_Hankaku                 Qt__Key = Qt__Key(0x01001129)
	Qt__Key_Zenkaku_Hankaku         Qt__Key = Qt__Key(0x0100112a)
	Qt__Key_Touroku                 Qt__Key = Qt__Key(0x0100112b)
	Qt__Key_Massyo                  Qt__Key = Qt__Key(0x0100112c)
	Qt__Key_Kana_Lock               Qt__Key = Qt__Key(0x0100112d)
	Qt__Key_Kana_Shift              Qt__Key = Qt__Key(0x0100112e)
	Qt__Key_Eisu_Shift              Qt__Key = Qt__Key(0x0100112f)
	Qt__Key_Eisu_toggle             Qt__Key = Qt__Key(0x01001130)
	Qt__Key_Hangul                  Qt__Key = Qt__Key(0x01001131)
	Qt__Key_Hangul_Start            Qt__Key = Qt__Key(0x01001132)
	Qt__Key_Hangul_End              Qt__Key = Qt__Key(0x01001133)
	Qt__Key_Hangul_Hanja            Qt__Key = Qt__Key(0x01001134)
	Qt__Key_Hangul_Jamo             Qt__Key = Qt__Key(0x01001135)
	Qt__Key_Hangul_Romaja           Qt__Key = Qt__Key(0x01001136)
	Qt__Key_Hangul_Jeonja           Qt__Key = Qt__Key(0x01001138)
	Qt__Key_Hangul_Banja            Qt__Key = Qt__Key(0x01001139)
	Qt__Key_Hangul_PreHanja         Qt__Key = Qt__Key(0x0100113a)
	Qt__Key_Hangul_PostHanja        Qt__Key = Qt__Key(0x0100113b)
	Qt__Key_Hangul_Special          Qt__Key = Qt__Key(0x0100113f)
	Qt__Key_Dead_Grave              Qt__Key = Qt__Key(0x01001250)
	Qt__Key_Dead_Acute              Qt__Key = Qt__Key(0x01001251)
	Qt__Key_Dead_Circumflex         Qt__Key = Qt__Key(0x01001252)
	Qt__Key_Dead_Tilde              Qt__Key = Qt__Key(0x01001253)
	Qt__Key_Dead_Macron             Qt__Key = Qt__Key(0x01001254)
	Qt__Key_Dead_Breve              Qt__Key = Qt__Key(0x01001255)
	Qt__Key_Dead_Abovedot           Qt__Key = Qt__Key(0x01001256)
	Qt__Key_Dead_Diaeresis          Qt__Key = Qt__Key(0x01001257)
	Qt__Key_Dead_Abovering          Qt__Key = Qt__Key(0x01001258)
	Qt__Key_Dead_Doubleacute        Qt__Key = Qt__Key(0x01001259)
	Qt__Key_Dead_Caron              Qt__Key = Qt__Key(0x0100125a)
	Qt__Key_Dead_Cedilla            Qt__Key = Qt__Key(0x0100125b)
	Qt__Key_Dead_Ogonek             Qt__Key = Qt__Key(0x0100125c)
	Qt__Key_Dead_Iota               Qt__Key = Qt__Key(0x0100125d)
	Qt__Key_Dead_Voiced_Sound       Qt__Key = Qt__Key(0x0100125e)
	Qt__Key_Dead_Semivoiced_Sound   Qt__Key = Qt__Key(0x0100125f)
	Qt__Key_Dead_Belowdot           Qt__Key = Qt__Key(0x01001260)
	Qt__Key_Dead_Hook               Qt__Key = Qt__Key(0x01001261)
	Qt__Key_Dead_Horn               Qt__Key = Qt__Key(0x01001262)
	Qt__Key_Dead_Stroke             Qt__Key = Qt__Key(0x01001263)
	Qt__Key_Dead_Abovecomma         Qt__Key = Qt__Key(0x01001264)
	Qt__Key_Dead_Abovereversedcomma Qt__Key = Qt__Key(0x01001265)
	Qt__Key_Dead_Doublegrave        Qt__Key = Qt__Key(0x01001266)
	Qt__Key_Dead_Belowring          Qt__Key = Qt__Key(0x01001267)
	Qt__Key_Dead_Belowmacron        Qt__Key = Qt__Key(0x01001268)
	Qt__Key_Dead_Belowcircumflex    Qt__Key = Qt__Key(0x01001269)
	Qt__Key_Dead_Belowtilde         Qt__Key = Qt__Key(0x0100126a)
	Qt__Key_Dead_Belowbreve         Qt__Key = Qt__Key(0x0100126b)
	Qt__Key_Dead_Belowdiaeresis     Qt__Key = Qt__Key(0x0100126c)
	Qt__Key_Dead_Invertedbreve      Qt__Key = Qt__Key(0x0100126d)
	Qt__Key_Dead_Belowcomma         Qt__Key = Qt__Key(0x0100126e)
	Qt__Key_Dead_Currency           Qt__Key = Qt__Key(0x0100126f)
	Qt__Key_Dead_a                  Qt__Key = Qt__Key(0x01001280)
	Qt__Key_Dead_A                  Qt__Key = Qt__Key(0x01001281)
	Qt__Key_Dead_e                  Qt__Key = Qt__Key(0x01001282)
	Qt__Key_Dead_E                  Qt__Key = Qt__Key(0x01001283)
	Qt__Key_Dead_i                  Qt__Key = Qt__Key(0x01001284)
	Qt__Key_Dead_I                  Qt__Key = Qt__Key(0x01001285)
	Qt__Key_Dead_o                  Qt__Key = Qt__Key(0x01001286)
	Qt__Key_Dead_O                  Qt__Key = Qt__Key(0x01001287)
	Qt__Key_Dead_u                  Qt__Key = Qt__Key(0x01001288)
	Qt__Key_Dead_U                  Qt__Key = Qt__Key(0x01001289)
	Qt__Key_Dead_Small_Schwa        Qt__Key = Qt__Key(0x0100128a)
	Qt__Key_Dead_Capital_Schwa      Qt__Key = Qt__Key(0x0100128b)
	Qt__Key_Dead_Greek              Qt__Key = Qt__Key(0x0100128c)
	Qt__Key_Dead_Lowline            Qt__Key = Qt__Key(0x01001290)
	Qt__Key_Dead_Aboveverticalline  Qt__Key = Qt__Key(0x01001291)
	Qt__Key_Dead_Belowverticalline  Qt__Key = Qt__Key(0x01001292)
	Qt__Key_Dead_Longsolidusoverlay Qt__Key = Qt__Key(0x01001293)
	Qt__Key_Back                    Qt__Key = Qt__Key(0x01000061)
	Qt__Key_Forward                 Qt__Key = Qt__Key(0x01000062)
	Qt__Key_Stop                    Qt__Key = Qt__Key(0x01000063)
	Qt__Key_Refresh                 Qt__Key = Qt__Key(0x01000064)
	Qt__Key_VolumeDown              Qt__Key = Qt__Key(0x01000070)
	Qt__Key_VolumeMute              Qt__Key = Qt__Key(0x01000071)
	Qt__Key_VolumeUp                Qt__Key = Qt__Key(0x01000072)
	Qt__Key_BassBoost               Qt__Key = Qt__Key(0x01000073)
	Qt__Key_BassUp                  Qt__Key = Qt__Key(0x01000074)
	Qt__Key_BassDown                Qt__Key = Qt__Key(0x01000075)
	Qt__Key_TrebleUp                Qt__Key = Qt__Key(0x01000076)
	Qt__Key_TrebleDown              Qt__Key = Qt__Key(0x01000077)
	Qt__Key_MediaPlay               Qt__Key = Qt__Key(0x01000080)
	Qt__Key_MediaStop               Qt__Key = Qt__Key(0x01000081)
	Qt__Key_MediaPrevious           Qt__Key = Qt__Key(0x01000082)
	Qt__Key_MediaNext               Qt__Key = Qt__Key(0x01000083)
	Qt__Key_MediaRecord             Qt__Key = Qt__Key(0x01000084)
	Qt__Key_MediaPause              Qt__Key = Qt__Key(0x1000085)
	Qt__Key_MediaTogglePlayPause    Qt__Key = Qt__Key(0x1000086)
	Qt__Key_HomePage                Qt__Key = Qt__Key(0x01000090)
	Qt__Key_Favorites               Qt__Key = Qt__Key(0x01000091)
	Qt__Key_Search                  Qt__Key = Qt__Key(0x01000092)
	Qt__Key_Standby                 Qt__Key = Qt__Key(0x01000093)
	Qt__Key_OpenUrl                 Qt__Key = Qt__Key(0x01000094)
	Qt__Key_LaunchMail              Qt__Key = Qt__Key(0x010000a0)
	Qt__Key_LaunchMedia             Qt__Key = Qt__Key(0x010000a1)
	Qt__Key_Launch0                 Qt__Key = Qt__Key(0x010000a2)
	Qt__Key_Launch1                 Qt__Key = Qt__Key(0x010000a3)
	Qt__Key_Launch2                 Qt__Key = Qt__Key(0x010000a4)
	Qt__Key_Launch3                 Qt__Key = Qt__Key(0x010000a5)
	Qt__Key_Launch4                 Qt__Key = Qt__Key(0x010000a6)
	Qt__Key_Launch5                 Qt__Key = Qt__Key(0x010000a7)
	Qt__Key_Launch6                 Qt__Key = Qt__Key(0x010000a8)
	Qt__Key_Launch7                 Qt__Key = Qt__Key(0x010000a9)
	Qt__Key_Launch8                 Qt__Key = Qt__Key(0x010000aa)
	Qt__Key_Launch9                 Qt__Key = Qt__Key(0x010000ab)
	Qt__Key_LaunchA                 Qt__Key = Qt__Key(0x010000ac)
	Qt__Key_LaunchB                 Qt__Key = Qt__Key(0x010000ad)
	Qt__Key_LaunchC                 Qt__Key = Qt__Key(0x010000ae)
	Qt__Key_LaunchD                 Qt__Key = Qt__Key(0x010000af)
	Qt__Key_LaunchE                 Qt__Key = Qt__Key(0x010000b0)
	Qt__Key_LaunchF                 Qt__Key = Qt__Key(0x010000b1)
	Qt__Key_MonBrightnessUp         Qt__Key = Qt__Key(0x010000b2)
	Qt__Key_MonBrightnessDown       Qt__Key = Qt__Key(0x010000b3)
	Qt__Key_KeyboardLightOnOff      Qt__Key = Qt__Key(0x010000b4)
	Qt__Key_KeyboardBrightnessUp    Qt__Key = Qt__Key(0x010000b5)
	Qt__Key_KeyboardBrightnessDown  Qt__Key = Qt__Key(0x010000b6)
	Qt__Key_PowerOff                Qt__Key = Qt__Key(0x010000b7)
	Qt__Key_WakeUp                  Qt__Key = Qt__Key(0x010000b8)
	Qt__Key_Eject                   Qt__Key = Qt__Key(0x010000b9)
	Qt__Key_ScreenSaver             Qt__Key = Qt__Key(0x010000ba)
	Qt__Key_WWW                     Qt__Key = Qt__Key(0x010000bb)
	Qt__Key_Memo                    Qt__Key = Qt__Key(0x010000bc)
	Qt__Key_LightBulb               Qt__Key = Qt__Key(0x010000bd)
	Qt__Key_Shop                    Qt__Key = Qt__Key(0x010000be)
	Qt__Key_History                 Qt__Key = Qt__Key(0x010000bf)
	Qt__Key_AddFavorite             Qt__Key = Qt__Key(0x010000c0)
	Qt__Key_HotLinks                Qt__Key = Qt__Key(0x010000c1)
	Qt__Key_BrightnessAdjust        Qt__Key = Qt__Key(0x010000c2)
	Qt__Key_Finance                 Qt__Key = Qt__Key(0x010000c3)
	Qt__Key_Community               Qt__Key = Qt__Key(0x010000c4)
	Qt__Key_AudioRewind             Qt__Key = Qt__Key(0x010000c5)
	Qt__Key_BackForward             Qt__Key = Qt__Key(0x010000c6)
	Qt__Key_ApplicationLeft         Qt__Key = Qt__Key(0x010000c7)
	Qt__Key_ApplicationRight        Qt__Key = Qt__Key(0x010000c8)
	Qt__Key_Book                    Qt__Key = Qt__Key(0x010000c9)
	Qt__Key_CD                      Qt__Key = Qt__Key(0x010000ca)
	Qt__Key_Calculator              Qt__Key = Qt__Key(0x010000cb)
	Qt__Key_ToDoList                Qt__Key = Qt__Key(0x010000cc)
	Qt__Key_ClearGrab               Qt__Key = Qt__Key(0x010000cd)
	Qt__Key_Close                   Qt__Key = Qt__Key(0x010000ce)
	Qt__Key_Copy                    Qt__Key = Qt__Key(0x010000cf)
	Qt__Key_Cut                     Qt__Key = Qt__Key(0x010000d0)
	Qt__Key_Display                 Qt__Key = Qt__Key(0x010000d1)
	Qt__Key_DOS                     Qt__Key = Qt__Key(0x010000d2)
	Qt__Key_Documents               Qt__Key = Qt__Key(0x010000d3)
	Qt__Key_Excel                   Qt__Key = Qt__Key(0x010000d4)
	Qt__Key_Explorer                Qt__Key = Qt__Key(0x010000d5)
	Qt__Key_Game                    Qt__Key = Qt__Key(0x010000d6)
	Qt__Key_Go                      Qt__Key = Qt__Key(0x010000d7)
	Qt__Key_iTouch                  Qt__Key = Qt__Key(0x010000d8)
	Qt__Key_LogOff                  Qt__Key = Qt__Key(0x010000d9)
	Qt__Key_Market                  Qt__Key = Qt__Key(0x010000da)
	Qt__Key_Meeting                 Qt__Key = Qt__Key(0x010000db)
	Qt__Key_MenuKB                  Qt__Key = Qt__Key(0x010000dc)
	Qt__Key_MenuPB                  Qt__Key = Qt__Key(0x010000dd)
	Qt__Key_MySites                 Qt__Key = Qt__Key(0x010000de)
	Qt__Key_News                    Qt__Key = Qt__Key(0x010000df)
	Qt__Key_OfficeHome              Qt__Key = Qt__Key(0x010000e0)
	Qt__Key_Option                  Qt__Key = Qt__Key(0x010000e1)
	Qt__Key_Paste                   Qt__Key = Qt__Key(0x010000e2)
	Qt__Key_Phone                   Qt__Key = Qt__Key(0x010000e3)
	Qt__Key_Calendar                Qt__Key = Qt__Key(0x010000e4)
	Qt__Key_Reply                   Qt__Key = Qt__Key(0x010000e5)
	Qt__Key_Reload                  Qt__Key = Qt__Key(0x010000e6)
	Qt__Key_RotateWindows           Qt__Key = Qt__Key(0x010000e7)
	Qt__Key_RotationPB              Qt__Key = Qt__Key(0x010000e8)
	Qt__Key_RotationKB              Qt__Key = Qt__Key(0x010000e9)
	Qt__Key_Save                    Qt__Key = Qt__Key(0x010000ea)
	Qt__Key_Send                    Qt__Key = Qt__Key(0x010000eb)
	Qt__Key_Spell                   Qt__Key = Qt__Key(0x010000ec)
	Qt__Key_SplitScreen             Qt__Key = Qt__Key(0x010000ed)
	Qt__Key_Support                 Qt__Key = Qt__Key(0x010000ee)
	Qt__Key_TaskPane                Qt__Key = Qt__Key(0x010000ef)
	Qt__Key_Terminal                Qt__Key = Qt__Key(0x010000f0)
	Qt__Key_Tools                   Qt__Key = Qt__Key(0x010000f1)
	Qt__Key_Travel                  Qt__Key = Qt__Key(0x010000f2)
	Qt__Key_Video                   Qt__Key = Qt__Key(0x010000f3)
	Qt__Key_Word                    Qt__Key = Qt__Key(0x010000f4)
	Qt__Key_Xfer                    Qt__Key = Qt__Key(0x010000f5)
	Qt__Key_ZoomIn                  Qt__Key = Qt__Key(0x010000f6)
	Qt__Key_ZoomOut                 Qt__Key = Qt__Key(0x010000f7)
	Qt__Key_Away                    Qt__Key = Qt__Key(0x010000f8)
	Qt__Key_Messenger               Qt__Key = Qt__Key(0x010000f9)
	Qt__Key_WebCam                  Qt__Key = Qt__Key(0x010000fa)
	Qt__Key_MailForward             Qt__Key = Qt__Key(0x010000fb)
	Qt__Key_Pictures                Qt__Key = Qt__Key(0x010000fc)
	Qt__Key_Music                   Qt__Key = Qt__Key(0x010000fd)
	Qt__Key_Battery                 Qt__Key = Qt__Key(0x010000fe)
	Qt__Key_Bluetooth               Qt__Key = Qt__Key(0x010000ff)
	Qt__Key_WLAN                    Qt__Key = Qt__Key(0x01000100)
	Qt__Key_UWB                     Qt__Key = Qt__Key(0x01000101)
	Qt__Key_AudioForward            Qt__Key = Qt__Key(0x01000102)
	Qt__Key_AudioRepeat             Qt__Key = Qt__Key(0x01000103)
	Qt__Key_AudioRandomPlay         Qt__Key = Qt__Key(0x01000104)
	Qt__Key_Subtitle                Qt__Key = Qt__Key(0x01000105)
	Qt__Key_AudioCycleTrack         Qt__Key = Qt__Key(0x01000106)
	Qt__Key_Time                    Qt__Key = Qt__Key(0x01000107)
	Qt__Key_Hibernate               Qt__Key = Qt__Key(0x01000108)
	Qt__Key_View                    Qt__Key = Qt__Key(0x01000109)
	Qt__Key_TopMenu                 Qt__Key = Qt__Key(0x0100010a)
	Qt__Key_PowerDown               Qt__Key = Qt__Key(0x0100010b)
	Qt__Key_Suspend                 Qt__Key = Qt__Key(0x0100010c)
	Qt__Key_ContrastAdjust          Qt__Key = Qt__Key(0x0100010d)
	Qt__Key_LaunchG                 Qt__Key = Qt__Key(0x0100010e)
	Qt__Key_LaunchH                 Qt__Key = Qt__Key(0x0100010f)
	Qt__Key_TouchpadToggle          Qt__Key = Qt__Key(0x01000110)
	Qt__Key_TouchpadOn              Qt__Key = Qt__Key(0x01000111)
	Qt__Key_TouchpadOff             Qt__Key = Qt__Key(0x01000112)
	Qt__Key_MicMute                 Qt__Key = Qt__Key(0x01000113)
	Qt__Key_Red                     Qt__Key = Qt__Key(0x01000114)
	Qt__Key_Green                   Qt__Key = Qt__Key(0x01000115)
	Qt__Key_Yellow                  Qt__Key = Qt__Key(0x01000116)
	Qt__Key_Blue                    Qt__Key = Qt__Key(0x01000117)
	Qt__Key_ChannelUp               Qt__Key = Qt__Key(0x01000118)
	Qt__Key_ChannelDown             Qt__Key = Qt__Key(0x01000119)
	Qt__Key_Guide                   Qt__Key = Qt__Key(0x0100011a)
	Qt__Key_Info                    Qt__Key = Qt__Key(0x0100011b)
	Qt__Key_Settings                Qt__Key = Qt__Key(0x0100011c)
	Qt__Key_MicVolumeUp             Qt__Key = Qt__Key(0x0100011d)
	Qt__Key_MicVolumeDown           Qt__Key = Qt__Key(0x0100011e)
	Qt__Key_New                     Qt__Key = Qt__Key(0x01000120)
	Qt__Key_Open                    Qt__Key = Qt__Key(0x01000121)
	Qt__Key_Find                    Qt__Key = Qt__Key(0x01000122)
	Qt__Key_Undo                    Qt__Key = Qt__Key(0x01000123)
	Qt__Key_Redo                    Qt__Key = Qt__Key(0x01000124)
	Qt__Key_MediaLast               Qt__Key = Qt__Key(0x0100ffff)
	Qt__Key_Select                  Qt__Key = Qt__Key(0x01010000)
	Qt__Key_Yes                     Qt__Key = Qt__Key(0x01010001)
	Qt__Key_No                      Qt__Key = Qt__Key(0x01010002)
	Qt__Key_Cancel                  Qt__Key = Qt__Key(0x01020001)
	Qt__Key_Printer                 Qt__Key = Qt__Key(0x01020002)
	Qt__Key_Execute                 Qt__Key = Qt__Key(0x01020003)
	Qt__Key_Sleep                   Qt__Key = Qt__Key(0x01020004)
	Qt__Key_Play                    Qt__Key = Qt__Key(0x01020005)
	Qt__Key_Zoom                    Qt__Key = Qt__Key(0x01020006)
	Qt__Key_Exit                    Qt__Key = Qt__Key(0x0102000a)
	Qt__Key_Context1                Qt__Key = Qt__Key(0x01100000)
	Qt__Key_Context2                Qt__Key = Qt__Key(0x01100001)
	Qt__Key_Context3                Qt__Key = Qt__Key(0x01100002)
	Qt__Key_Context4                Qt__Key = Qt__Key(0x01100003)
	Qt__Key_Call                    Qt__Key = Qt__Key(0x01100004)
	Qt__Key_Hangup                  Qt__Key = Qt__Key(0x01100005)
	Qt__Key_Flip                    Qt__Key = Qt__Key(0x01100006)
	Qt__Key_ToggleCallHangup        Qt__Key = Qt__Key(0x01100007)
	Qt__Key_VoiceDial               Qt__Key = Qt__Key(0x01100008)
	Qt__Key_LastNumberRedial        Qt__Key = Qt__Key(0x01100009)
	Qt__Key_Camera                  Qt__Key = Qt__Key(0x01100020)
	Qt__Key_CameraFocus             Qt__Key = Qt__Key(0x01100021)
	Qt__Key_unknown                 Qt__Key = Qt__Key(0x01ffffff)
)

// Qt::KeyboardModifier
//
//go:generate stringer -type=Qt__KeyboardModifier
type Qt__KeyboardModifier int64

const (
	Qt__NoModifier           Qt__KeyboardModifier = Qt__KeyboardModifier(0x00000000)
	Qt__ShiftModifier        Qt__KeyboardModifier = Qt__KeyboardModifier(0x02000000)
	Qt__ControlModifier      Qt__KeyboardModifier = Qt__KeyboardModifier(0x04000000)
	Qt__AltModifier          Qt__KeyboardModifier = Qt__KeyboardModifier(0x08000000)
	Qt__MetaModifier         Qt__KeyboardModifier = Qt__KeyboardModifier(0x10000000)
	Qt__KeypadModifier       Qt__KeyboardModifier = Qt__KeyboardModifier(0x20000000)
	Qt__GroupSwitchModifier  Qt__KeyboardModifier = Qt__KeyboardModifier(0x40000000)
	Qt__KeyboardModifierMask Qt__KeyboardModifier = Qt__KeyboardModifier(0xfe000000)
)

// Qt::LayoutDirection
//
//go:generate stringer -type=Qt__LayoutDirection
type Qt__LayoutDirection int64

const (
	Qt__LeftToRight         Qt__LayoutDirection = Qt__LayoutDirection(0)
	Qt__RightToLeft         Qt__LayoutDirection = Qt__LayoutDirection(1)
	Qt__LayoutDirectionAuto Qt__LayoutDirection = Qt__LayoutDirection(2)
)

// Qt::MaskMode
//
//go:generate stringer -type=Qt__MaskMode
type Qt__MaskMode int64

const (
	Qt__MaskInColor  Qt__MaskMode = Qt__MaskMode(0)
	Qt__MaskOutColor Qt__MaskMode = Qt__MaskMode(1)
)

// Qt::MatchFlag
//
//go:generate stringer -type=Qt__MatchFlag
type Qt__MatchFlag int64

const (
	Qt__MatchExactly           Qt__MatchFlag = Qt__MatchFlag(0)
	Qt__MatchContains          Qt__MatchFlag = Qt__MatchFlag(1)
	Qt__MatchStartsWith        Qt__MatchFlag = Qt__MatchFlag(2)
	Qt__MatchEndsWith          Qt__MatchFlag = Qt__MatchFlag(3)
	Qt__MatchRegExp            Qt__MatchFlag = Qt__MatchFlag(4)
	Qt__MatchWildcard          Qt__MatchFlag = Qt__MatchFlag(5)
	Qt__MatchFixedString       Qt__MatchFlag = Qt__MatchFlag(8)
	Qt__MatchRegularExpression Qt__MatchFlag = Qt__MatchFlag(9)
	Qt__MatchCaseSensitive     Qt__MatchFlag = Qt__MatchFlag(16)
	Qt__MatchWrap              Qt__MatchFlag = Qt__MatchFlag(32)
	Qt__MatchRecursive         Qt__MatchFlag = Qt__MatchFlag(64)
)

// Qt::Modifier
//
//go:generate stringer -type=Qt__Modifier
type Qt__Modifier int64

const (
	Qt__META          Qt__Modifier = Qt__Modifier(Qt__MetaModifier)
	Qt__SHIFT         Qt__Modifier = Qt__Modifier(Qt__ShiftModifier)
	Qt__CTRL          Qt__Modifier = Qt__Modifier(Qt__ControlModifier)
	Qt__ALT           Qt__Modifier = Qt__Modifier(Qt__AltModifier)
	Qt__MODIFIER_MASK Qt__Modifier = Qt__Modifier(Qt__KeyboardModifierMask)
	Qt__UNICODE_ACCEL Qt__Modifier = Qt__Modifier(0x00000000)
)

// Qt::MouseButton
//
//go:generate stringer -type=Qt__MouseButton
type Qt__MouseButton int64

const (
	Qt__NoButton        Qt__MouseButton = Qt__MouseButton(0x00000000)
	Qt__LeftButton      Qt__MouseButton = Qt__MouseButton(0x00000001)
	Qt__RightButton     Qt__MouseButton = Qt__MouseButton(0x00000002)
	Qt__MiddleButton    Qt__MouseButton = Qt__MouseButton(0x00000004)
	Qt__MidButton       Qt__MouseButton = Qt__MouseButton(Qt__MiddleButton)
	Qt__BackButton      Qt__MouseButton = Qt__MouseButton(0x00000008)
	Qt__XButton1        Qt__MouseButton = Qt__MouseButton(Qt__BackButton)
	Qt__ExtraButton1    Qt__MouseButton = Qt__MouseButton(Qt__XButton1)
	Qt__ForwardButton   Qt__MouseButton = Qt__MouseButton(0x00000010)
	Qt__XButton2        Qt__MouseButton = Qt__MouseButton(Qt__ForwardButton)
	Qt__ExtraButton2    Qt__MouseButton = Qt__MouseButton(Qt__ForwardButton)
	Qt__TaskButton      Qt__MouseButton = Qt__MouseButton(0x00000020)
	Qt__ExtraButton3    Qt__MouseButton = Qt__MouseButton(Qt__TaskButton)
	Qt__ExtraButton4    Qt__MouseButton = Qt__MouseButton(0x00000040)
	Qt__ExtraButton5    Qt__MouseButton = Qt__MouseButton(0x00000080)
	Qt__ExtraButton6    Qt__MouseButton = Qt__MouseButton(0x00000100)
	Qt__ExtraButton7    Qt__MouseButton = Qt__MouseButton(0x00000200)
	Qt__ExtraButton8    Qt__MouseButton = Qt__MouseButton(0x00000400)
	Qt__ExtraButton9    Qt__MouseButton = Qt__MouseButton(0x00000800)
	Qt__ExtraButton10   Qt__MouseButton = Qt__MouseButton(0x00001000)
	Qt__ExtraButton11   Qt__MouseButton = Qt__MouseButton(0x00002000)
	Qt__ExtraButton12   Qt__MouseButton = Qt__MouseButton(0x00004000)
	Qt__ExtraButton13   Qt__MouseButton = Qt__MouseButton(0x00008000)
	Qt__ExtraButton14   Qt__MouseButton = Qt__MouseButton(0x00010000)
	Qt__ExtraButton15   Qt__MouseButton = Qt__MouseButton(0x00020000)
	Qt__ExtraButton16   Qt__MouseButton = Qt__MouseButton(0x00040000)
	Qt__ExtraButton17   Qt__MouseButton = Qt__MouseButton(0x00080000)
	Qt__ExtraButton18   Qt__MouseButton = Qt__MouseButton(0x00100000)
	Qt__ExtraButton19   Qt__MouseButton = Qt__MouseButton(0x00200000)
	Qt__ExtraButton20   Qt__MouseButton = Qt__MouseButton(0x00400000)
	Qt__ExtraButton21   Qt__MouseButton = Qt__MouseButton(0x00800000)
	Qt__ExtraButton22   Qt__MouseButton = Qt__MouseButton(0x01000000)
	Qt__ExtraButton23   Qt__MouseButton = Qt__MouseButton(0x02000000)
	Qt__ExtraButton24   Qt__MouseButton = Qt__MouseButton(0x04000000)
	Qt__AllButtons      Qt__MouseButton = Qt__MouseButton(0x07ffffff)
	Qt__MaxMouseButton  Qt__MouseButton = Qt__MouseButton(Qt__ExtraButton24)
	Qt__MouseButtonMask Qt__MouseButton = Qt__MouseButton(0xffffffff)
)

// Qt::MouseEventFlag
//
//go:generate stringer -type=Qt__MouseEventFlag
type Qt__MouseEventFlag int64

const (
	Qt__MouseEventCreatedDoubleClick Qt__MouseEventFlag = Qt__MouseEventFlag(0x01)
	Qt__MouseEventFlagMask           Qt__MouseEventFlag = Qt__MouseEventFlag(0xFF)
)

// Qt::MouseEventSource
//
//go:generate stringer -type=Qt__MouseEventSource
type Qt__MouseEventSource int64

const (
	Qt__MouseEventNotSynthesized           Qt__MouseEventSource = Qt__MouseEventSource(0)
	Qt__MouseEventSynthesizedBySystem      Qt__MouseEventSource = Qt__MouseEventSource(1)
	Qt__MouseEventSynthesizedByQt          Qt__MouseEventSource = Qt__MouseEventSource(2)
	Qt__MouseEventSynthesizedByApplication Qt__MouseEventSource = Qt__MouseEventSource(3)
)

// Qt::NativeGestureType
//
//go:generate stringer -type=Qt__NativeGestureType
type Qt__NativeGestureType int64

const (
	Qt__BeginNativeGesture     Qt__NativeGestureType = Qt__NativeGestureType(0)
	Qt__EndNativeGesture       Qt__NativeGestureType = Qt__NativeGestureType(1)
	Qt__PanNativeGesture       Qt__NativeGestureType = Qt__NativeGestureType(2)
	Qt__ZoomNativeGesture      Qt__NativeGestureType = Qt__NativeGestureType(3)
	Qt__SmartZoomNativeGesture Qt__NativeGestureType = Qt__NativeGestureType(4)
	Qt__RotateNativeGesture    Qt__NativeGestureType = Qt__NativeGestureType(5)
	Qt__SwipeNativeGesture     Qt__NativeGestureType = Qt__NativeGestureType(6)
)

// Qt::NavigationMode
//
//go:generate stringer -type=Qt__NavigationMode
type Qt__NavigationMode int64

const (
	Qt__NavigationModeNone               Qt__NavigationMode = Qt__NavigationMode(0)
	Qt__NavigationModeKeypadTabOrder     Qt__NavigationMode = Qt__NavigationMode(1)
	Qt__NavigationModeKeypadDirectional  Qt__NavigationMode = Qt__NavigationMode(2)
	Qt__NavigationModeCursorAuto         Qt__NavigationMode = Qt__NavigationMode(3)
	Qt__NavigationModeCursorForceVisible Qt__NavigationMode = Qt__NavigationMode(4)
)

// Qt::Orientation
//
//go:generate stringer -type=Qt__Orientation
type Qt__Orientation int64

const (
	Qt__Horizontal Qt__Orientation = Qt__Orientation(0x1)
	Qt__Vertical   Qt__Orientation = Qt__Orientation(0x2)
)

// Qt::PenCapStyle
//
//go:generate stringer -type=Qt__PenCapStyle
type Qt__PenCapStyle int64

var (
	Qt__FlatCap      Qt__PenCapStyle = Qt__PenCapStyle(0x00)
	Qt__SquareCap    Qt__PenCapStyle = Qt__PenCapStyle(0x10)
	Qt__RoundCap     Qt__PenCapStyle = Qt__PenCapStyle(0x20)
	Qt__MPenCapStyle Qt__PenCapStyle = Qt__PenCapStyle(0x30)
)

// Qt::PenJoinStyle
//
//go:generate stringer -type=Qt__PenJoinStyle
type Qt__PenJoinStyle int64

var (
	Qt__MiterJoin     Qt__PenJoinStyle = Qt__PenJoinStyle(0x00)
	Qt__BevelJoin     Qt__PenJoinStyle = Qt__PenJoinStyle(0x40)
	Qt__RoundJoin     Qt__PenJoinStyle = Qt__PenJoinStyle(0x80)
	Qt__SvgMiterJoin  Qt__PenJoinStyle = Qt__PenJoinStyle(0x100)
	Qt__MPenJoinStyle Qt__PenJoinStyle = Qt__PenJoinStyle(0x1c0)
)

// Qt::PenStyle
//
//go:generate stringer -type=Qt__PenStyle
type Qt__PenStyle int64

var (
	Qt__NoPen          Qt__PenStyle = Qt__PenStyle(0)
	Qt__SolidLine      Qt__PenStyle = Qt__PenStyle(1)
	Qt__DashLine       Qt__PenStyle = Qt__PenStyle(2)
	Qt__DotLine        Qt__PenStyle = Qt__PenStyle(3)
	Qt__DashDotLine    Qt__PenStyle = Qt__PenStyle(4)
	Qt__DashDotDotLine Qt__PenStyle = Qt__PenStyle(5)
	Qt__CustomDashLine Qt__PenStyle = Qt__PenStyle(6)
	Qt__MPenStyle      Qt__PenStyle = Qt__PenStyle(0x0f)
)

// Qt::ReturnByValueConstant
//
//go:generate stringer -type=Qt__ReturnByValueConstant
type Qt__ReturnByValueConstant int64

const (
	Qt__ReturnByValue Qt__ReturnByValueConstant = Qt__ReturnByValueConstant(0)
)

// Qt::ScreenOrientation
//
//go:generate stringer -type=Qt__ScreenOrientation
type Qt__ScreenOrientation int64

const (
	Qt__PrimaryOrientation           Qt__ScreenOrientation = Qt__ScreenOrientation(0x00000000)
	Qt__PortraitOrientation          Qt__ScreenOrientation = Qt__ScreenOrientation(0x00000001)
	Qt__LandscapeOrientation         Qt__ScreenOrientation = Qt__ScreenOrientation(0x00000002)
	Qt__InvertedPortraitOrientation  Qt__ScreenOrientation = Qt__ScreenOrientation(0x00000004)
	Qt__InvertedLandscapeOrientation Qt__ScreenOrientation = Qt__ScreenOrientation(0x00000008)
)

// Qt::ScrollBarPolicy
//
//go:generate stringer -type=Qt__ScrollBarPolicy
type Qt__ScrollBarPolicy int64

const (
	Qt__ScrollBarAsNeeded  Qt__ScrollBarPolicy = Qt__ScrollBarPolicy(0)
	Qt__ScrollBarAlwaysOff Qt__ScrollBarPolicy = Qt__ScrollBarPolicy(1)
	Qt__ScrollBarAlwaysOn  Qt__ScrollBarPolicy = Qt__ScrollBarPolicy(2)
)

// Qt::ScrollPhase
//
//go:generate stringer -type=Qt__ScrollPhase
type Qt__ScrollPhase int64

const (
	Qt__NoScrollPhase  Qt__ScrollPhase = Qt__ScrollPhase(0)
	Qt__ScrollBegin    Qt__ScrollPhase = Qt__ScrollPhase(1)
	Qt__ScrollUpdate   Qt__ScrollPhase = Qt__ScrollPhase(2)
	Qt__ScrollEnd      Qt__ScrollPhase = Qt__ScrollPhase(3)
	Qt__ScrollMomentum Qt__ScrollPhase = Qt__ScrollPhase(4)
)

// Qt::ShortcutContext
//
//go:generate stringer -type=Qt__ShortcutContext
type Qt__ShortcutContext int64

const (
	Qt__WidgetShortcut             Qt__ShortcutContext = Qt__ShortcutContext(0)
	Qt__WindowShortcut             Qt__ShortcutContext = Qt__ShortcutContext(1)
	Qt__ApplicationShortcut        Qt__ShortcutContext = Qt__ShortcutContext(2)
	Qt__WidgetWithChildrenShortcut Qt__ShortcutContext = Qt__ShortcutContext(3)
)

// Qt::SizeHint
//
//go:generate stringer -type=Qt__SizeHint
type Qt__SizeHint int64

const (
	Qt__MinimumSize    Qt__SizeHint = Qt__SizeHint(0)
	Qt__PreferredSize  Qt__SizeHint = Qt__SizeHint(1)
	Qt__MaximumSize    Qt__SizeHint = Qt__SizeHint(2)
	Qt__MinimumDescent Qt__SizeHint = Qt__SizeHint(3)
	Qt__NSizeHints     Qt__SizeHint = Qt__SizeHint(4)
)

// Qt::SizeMode
//
//go:generate stringer -type=Qt__SizeMode
type Qt__SizeMode int64

const (
	Qt__AbsoluteSize Qt__SizeMode = Qt__SizeMode(0)
	Qt__RelativeSize Qt__SizeMode = Qt__SizeMode(1)
)

// Qt::SortOrder
//
//go:generate stringer -type=Qt__SortOrder
type Qt__SortOrder int64

const (
	Qt__AscendingOrder  Qt__SortOrder = Qt__SortOrder(0)
	Qt__DescendingOrder Qt__SortOrder = Qt__SortOrder(1)
)

// Qt::SplitBehaviorFlags
//
//go:generate stringer -type=Qt__SplitBehaviorFlags
type Qt__SplitBehaviorFlags int64

const (
	Qt__KeepEmptyParts Qt__SplitBehaviorFlags = Qt__SplitBehaviorFlags(0)
	Qt__SkipEmptyParts Qt__SplitBehaviorFlags = Qt__SplitBehaviorFlags(0x1)
)

// Qt::TabFocusBehavior
//
//go:generate stringer -type=Qt__TabFocusBehavior
type Qt__TabFocusBehavior int64

const (
	Qt__NoTabFocus           Qt__TabFocusBehavior = Qt__TabFocusBehavior(0x00)
	Qt__TabFocusTextControls Qt__TabFocusBehavior = Qt__TabFocusBehavior(0x01)
	Qt__TabFocusListControls Qt__TabFocusBehavior = Qt__TabFocusBehavior(0x02)
	Qt__TabFocusAllControls  Qt__TabFocusBehavior = Qt__TabFocusBehavior(0xff)
)

// Qt::TextElideMode
//
//go:generate stringer -type=Qt__TextElideMode
type Qt__TextElideMode int64

const (
	Qt__ElideLeft   Qt__TextElideMode = Qt__TextElideMode(0)
	Qt__ElideRight  Qt__TextElideMode = Qt__TextElideMode(1)
	Qt__ElideMiddle Qt__TextElideMode = Qt__TextElideMode(2)
	Qt__ElideNone   Qt__TextElideMode = Qt__TextElideMode(3)
)

// Qt::TextFlag
//
//go:generate stringer -type=Qt__TextFlag
type Qt__TextFlag int64

const (
	Qt__TextSingleLine            Qt__TextFlag = Qt__TextFlag(0x0100)
	Qt__TextDontClip              Qt__TextFlag = Qt__TextFlag(0x0200)
	Qt__TextExpandTabs            Qt__TextFlag = Qt__TextFlag(0x0400)
	Qt__TextShowMnemonic          Qt__TextFlag = Qt__TextFlag(0x0800)
	Qt__TextWordWrap              Qt__TextFlag = Qt__TextFlag(0x1000)
	Qt__TextWrapAnywhere          Qt__TextFlag = Qt__TextFlag(0x2000)
	Qt__TextDontPrint             Qt__TextFlag = Qt__TextFlag(0x4000)
	Qt__TextIncludeTrailingSpaces Qt__TextFlag = Qt__TextFlag(0x08000000)
	Qt__TextHideMnemonic          Qt__TextFlag = Qt__TextFlag(0x8000)
	Qt__TextJustificationForced   Qt__TextFlag = Qt__TextFlag(0x10000)
	Qt__TextForceLeftToRight      Qt__TextFlag = Qt__TextFlag(0x20000)
	Qt__TextForceRightToLeft      Qt__TextFlag = Qt__TextFlag(0x40000)
	Qt__TextLongestVariant        Qt__TextFlag = Qt__TextFlag(0x80000)
	Qt__TextBypassShaping         Qt__TextFlag = Qt__TextFlag(0x100000)
)

// Qt::TextFormat
//
//go:generate stringer -type=Qt__TextFormat
type Qt__TextFormat int64

const (
	Qt__PlainText    Qt__TextFormat = Qt__TextFormat(0)
	Qt__RichText     Qt__TextFormat = Qt__TextFormat(1)
	Qt__AutoText     Qt__TextFormat = Qt__TextFormat(2)
	Qt__MarkdownText Qt__TextFormat = Qt__TextFormat(3)
)

// Qt::TextInteractionFlag
//
//go:generate stringer -type=Qt__TextInteractionFlag
type Qt__TextInteractionFlag int64

const (
	Qt__NoTextInteraction         Qt__TextInteractionFlag = Qt__TextInteractionFlag(0)
	Qt__TextSelectableByMouse     Qt__TextInteractionFlag = Qt__TextInteractionFlag(1)
	Qt__TextSelectableByKeyboard  Qt__TextInteractionFlag = Qt__TextInteractionFlag(2)
	Qt__LinksAccessibleByMouse    Qt__TextInteractionFlag = Qt__TextInteractionFlag(4)
	Qt__LinksAccessibleByKeyboard Qt__TextInteractionFlag = Qt__TextInteractionFlag(8)
	Qt__TextEditable              Qt__TextInteractionFlag = Qt__TextInteractionFlag(16)
	Qt__TextEditorInteraction     Qt__TextInteractionFlag = Qt__TextInteractionFlag(Qt__TextSelectableByMouse | Qt__TextSelectableByKeyboard | Qt__TextEditable)
	Qt__TextBrowserInteraction    Qt__TextInteractionFlag = Qt__TextInteractionFlag(Qt__TextSelectableByMouse | Qt__LinksAccessibleByMouse | Qt__LinksAccessibleByKeyboard)
)

// Qt::TileRule
//
//go:generate stringer -type=Qt__TileRule
type Qt__TileRule int64

const (
	Qt__StretchTile Qt__TileRule = Qt__TileRule(0)
	Qt__RepeatTile  Qt__TileRule = Qt__TileRule(1)
	Qt__RoundTile   Qt__TileRule = Qt__TileRule(2)
)

// Qt::TimeSpec
//
//go:generate stringer -type=Qt__TimeSpec
type Qt__TimeSpec int64

const (
	Qt__LocalTime     Qt__TimeSpec = Qt__TimeSpec(0)
	Qt__UTC           Qt__TimeSpec = Qt__TimeSpec(1)
	Qt__OffsetFromUTC Qt__TimeSpec = Qt__TimeSpec(2)
	Qt__TimeZone      Qt__TimeSpec = Qt__TimeSpec(3)
)

// Qt::TimerType
//
//go:generate stringer -type=Qt__TimerType
type Qt__TimerType int64

const (
	Qt__PreciseTimer    Qt__TimerType = Qt__TimerType(0)
	Qt__CoarseTimer     Qt__TimerType = Qt__TimerType(1)
	Qt__VeryCoarseTimer Qt__TimerType = Qt__TimerType(2)
)

// Qt::ToolBarArea
//
//go:generate stringer -type=Qt__ToolBarArea
type Qt__ToolBarArea int64

const (
	Qt__LeftToolBarArea   Qt__ToolBarArea = Qt__ToolBarArea(0x1)
	Qt__RightToolBarArea  Qt__ToolBarArea = Qt__ToolBarArea(0x2)
	Qt__TopToolBarArea    Qt__ToolBarArea = Qt__ToolBarArea(0x4)
	Qt__BottomToolBarArea Qt__ToolBarArea = Qt__ToolBarArea(0x8)
	Qt__ToolBarArea_Mask  Qt__ToolBarArea = Qt__ToolBarArea(0xf)
	Qt__AllToolBarAreas   Qt__ToolBarArea = Qt__ToolBarArea(Qt__ToolBarArea_Mask)
	Qt__NoToolBarArea     Qt__ToolBarArea = Qt__ToolBarArea(0)
)

// Qt::ToolButtonStyle
//
//go:generate stringer -type=Qt__ToolButtonStyle
type Qt__ToolButtonStyle int64

var (
	Qt__ToolButtonIconOnly       Qt__ToolButtonStyle = Qt__ToolButtonStyle(0)
	Qt__ToolButtonTextOnly       Qt__ToolButtonStyle = Qt__ToolButtonStyle(1)
	Qt__ToolButtonTextBesideIcon Qt__ToolButtonStyle = Qt__ToolButtonStyle(2)
	Qt__ToolButtonTextUnderIcon  Qt__ToolButtonStyle = Qt__ToolButtonStyle(3)
	Qt__ToolButtonFollowStyle    Qt__ToolButtonStyle = Qt__ToolButtonStyle(4)
)

// Qt::TouchPointState
//
//go:generate stringer -type=Qt__TouchPointState
type Qt__TouchPointState int64

const (
	Qt__TouchPointPressed    Qt__TouchPointState = Qt__TouchPointState(0x01)
	Qt__TouchPointMoved      Qt__TouchPointState = Qt__TouchPointState(0x02)
	Qt__TouchPointStationary Qt__TouchPointState = Qt__TouchPointState(0x04)
	Qt__TouchPointReleased   Qt__TouchPointState = Qt__TouchPointState(0x08)
)

// Qt::TransformationMode
//
//go:generate stringer -type=Qt__TransformationMode
type Qt__TransformationMode int64

const (
	Qt__FastTransformation   Qt__TransformationMode = Qt__TransformationMode(0)
	Qt__SmoothTransformation Qt__TransformationMode = Qt__TransformationMode(1)
)

// Qt::UIEffect
//
//go:generate stringer -type=Qt__UIEffect
type Qt__UIEffect int64

const (
	Qt__UI_General        Qt__UIEffect = Qt__UIEffect(0)
	Qt__UI_AnimateMenu    Qt__UIEffect = Qt__UIEffect(1)
	Qt__UI_FadeMenu       Qt__UIEffect = Qt__UIEffect(2)
	Qt__UI_AnimateCombo   Qt__UIEffect = Qt__UIEffect(3)
	Qt__UI_AnimateTooltip Qt__UIEffect = Qt__UIEffect(4)
	Qt__UI_FadeTooltip    Qt__UIEffect = Qt__UIEffect(5)
	Qt__UI_AnimateToolBox Qt__UIEffect = Qt__UIEffect(6)
)

// Qt::WhiteSpaceMode
//
//go:generate stringer -type=Qt__WhiteSpaceMode
type Qt__WhiteSpaceMode int64

const (
	Qt__WhiteSpaceNormal        Qt__WhiteSpaceMode = Qt__WhiteSpaceMode(0)
	Qt__WhiteSpacePre           Qt__WhiteSpaceMode = Qt__WhiteSpaceMode(1)
	Qt__WhiteSpaceNoWrap        Qt__WhiteSpaceMode = Qt__WhiteSpaceMode(2)
	Qt__WhiteSpaceModeUndefined Qt__WhiteSpaceMode = Qt__WhiteSpaceMode(-1)
)

// Qt::WidgetAttribute
//
//go:generate stringer -type=Qt__WidgetAttribute
type Qt__WidgetAttribute int64

const (
	Qt__WA_Disabled                        Qt__WidgetAttribute = Qt__WidgetAttribute(0)
	Qt__WA_UnderMouse                      Qt__WidgetAttribute = Qt__WidgetAttribute(1)
	Qt__WA_MouseTracking                   Qt__WidgetAttribute = Qt__WidgetAttribute(2)
	Qt__WA_ContentsPropagated              Qt__WidgetAttribute = Qt__WidgetAttribute(3)
	Qt__WA_OpaquePaintEvent                Qt__WidgetAttribute = Qt__WidgetAttribute(4)
	Qt__WA_NoBackground                    Qt__WidgetAttribute = Qt__WidgetAttribute(Qt__WA_OpaquePaintEvent)
	Qt__WA_StaticContents                  Qt__WidgetAttribute = Qt__WidgetAttribute(5)
	Qt__WA_LaidOut                         Qt__WidgetAttribute = Qt__WidgetAttribute(7)
	Qt__WA_PaintOnScreen                   Qt__WidgetAttribute = Qt__WidgetAttribute(8)
	Qt__WA_NoSystemBackground              Qt__WidgetAttribute = Qt__WidgetAttribute(9)
	Qt__WA_UpdatesDisabled                 Qt__WidgetAttribute = Qt__WidgetAttribute(10)
	Qt__WA_Mapped                          Qt__WidgetAttribute = Qt__WidgetAttribute(11)
	Qt__WA_MacNoClickThrough               Qt__WidgetAttribute = Qt__WidgetAttribute(12)
	Qt__WA_InputMethodEnabled              Qt__WidgetAttribute = Qt__WidgetAttribute(14)
	Qt__WA_WState_Visible                  Qt__WidgetAttribute = Qt__WidgetAttribute(15)
	Qt__WA_WState_Hidden                   Qt__WidgetAttribute = Qt__WidgetAttribute(16)
	Qt__WA_ForceDisabled                   Qt__WidgetAttribute = Qt__WidgetAttribute(32)
	Qt__WA_KeyCompression                  Qt__WidgetAttribute = Qt__WidgetAttribute(33)
	Qt__WA_PendingMoveEvent                Qt__WidgetAttribute = Qt__WidgetAttribute(34)
	Qt__WA_PendingResizeEvent              Qt__WidgetAttribute = Qt__WidgetAttribute(35)
	Qt__WA_SetPalette                      Qt__WidgetAttribute = Qt__WidgetAttribute(36)
	Qt__WA_SetFont                         Qt__WidgetAttribute = Qt__WidgetAttribute(37)
	Qt__WA_SetCursor                       Qt__WidgetAttribute = Qt__WidgetAttribute(38)
	Qt__WA_NoChildEventsFromChildren       Qt__WidgetAttribute = Qt__WidgetAttribute(39)
	Qt__WA_WindowModified                  Qt__WidgetAttribute = Qt__WidgetAttribute(41)
	Qt__WA_Resized                         Qt__WidgetAttribute = Qt__WidgetAttribute(42)
	Qt__WA_Moved                           Qt__WidgetAttribute = Qt__WidgetAttribute(43)
	Qt__WA_PendingUpdate                   Qt__WidgetAttribute = Qt__WidgetAttribute(44)
	Qt__WA_InvalidSize                     Qt__WidgetAttribute = Qt__WidgetAttribute(45)
	Qt__WA_MacBrushedMetal                 Qt__WidgetAttribute = Qt__WidgetAttribute(46)
	Qt__WA_MacMetalStyle                   Qt__WidgetAttribute = Qt__WidgetAttribute(46)
	Qt__WA_CustomWhatsThis                 Qt__WidgetAttribute = Qt__WidgetAttribute(47)
	Qt__WA_LayoutOnEntireRect              Qt__WidgetAttribute = Qt__WidgetAttribute(48)
	Qt__WA_OutsideWSRange                  Qt__WidgetAttribute = Qt__WidgetAttribute(49)
	Qt__WA_GrabbedShortcut                 Qt__WidgetAttribute = Qt__WidgetAttribute(50)
	Qt__WA_TransparentForMouseEvents       Qt__WidgetAttribute = Qt__WidgetAttribute(51)
	Qt__WA_PaintUnclipped                  Qt__WidgetAttribute = Qt__WidgetAttribute(52)
	Qt__WA_SetWindowIcon                   Qt__WidgetAttribute = Qt__WidgetAttribute(53)
	Qt__WA_NoMouseReplay                   Qt__WidgetAttribute = Qt__WidgetAttribute(54)
	Qt__WA_DeleteOnClose                   Qt__WidgetAttribute = Qt__WidgetAttribute(55)
	Qt__WA_RightToLeft                     Qt__WidgetAttribute = Qt__WidgetAttribute(56)
	Qt__WA_SetLayoutDirection              Qt__WidgetAttribute = Qt__WidgetAttribute(57)
	Qt__WA_NoChildEventsForParent          Qt__WidgetAttribute = Qt__WidgetAttribute(58)
	Qt__WA_ForceUpdatesDisabled            Qt__WidgetAttribute = Qt__WidgetAttribute(59)
	Qt__WA_WState_Created                  Qt__WidgetAttribute = Qt__WidgetAttribute(60)
	Qt__WA_WState_CompressKeys             Qt__WidgetAttribute = Qt__WidgetAttribute(61)
	Qt__WA_WState_InPaintEvent             Qt__WidgetAttribute = Qt__WidgetAttribute(62)
	Qt__WA_WState_Reparented               Qt__WidgetAttribute = Qt__WidgetAttribute(63)
	Qt__WA_WState_ConfigPending            Qt__WidgetAttribute = Qt__WidgetAttribute(64)
	Qt__WA_WState_Polished                 Qt__WidgetAttribute = Qt__WidgetAttribute(66)
	Qt__WA_WState_DND                      Qt__WidgetAttribute = Qt__WidgetAttribute(67)
	Qt__WA_WState_OwnSizePolicy            Qt__WidgetAttribute = Qt__WidgetAttribute(68)
	Qt__WA_WState_ExplicitShowHide         Qt__WidgetAttribute = Qt__WidgetAttribute(69)
	Qt__WA_ShowModal                       Qt__WidgetAttribute = Qt__WidgetAttribute(70)
	Qt__WA_MouseNoMask                     Qt__WidgetAttribute = Qt__WidgetAttribute(71)
	Qt__WA_GroupLeader                     Qt__WidgetAttribute = Qt__WidgetAttribute(72)
	Qt__WA_NoMousePropagation              Qt__WidgetAttribute = Qt__WidgetAttribute(73)
	Qt__WA_Hover                           Qt__WidgetAttribute = Qt__WidgetAttribute(74)
	Qt__WA_InputMethodTransparent          Qt__WidgetAttribute = Qt__WidgetAttribute(75)
	Qt__WA_QuitOnClose                     Qt__WidgetAttribute = Qt__WidgetAttribute(76)
	Qt__WA_KeyboardFocusChange             Qt__WidgetAttribute = Qt__WidgetAttribute(77)
	Qt__WA_AcceptDrops                     Qt__WidgetAttribute = Qt__WidgetAttribute(78)
	Qt__WA_DropSiteRegistered              Qt__WidgetAttribute = Qt__WidgetAttribute(79)
	Qt__WA_ForceAcceptDrops                Qt__WidgetAttribute = Qt__WidgetAttribute(Qt__WA_DropSiteRegistered)
	Qt__WA_WindowPropagation               Qt__WidgetAttribute = Qt__WidgetAttribute(80)
	Qt__WA_NoX11EventCompression           Qt__WidgetAttribute = Qt__WidgetAttribute(81)
	Qt__WA_TintedBackground                Qt__WidgetAttribute = Qt__WidgetAttribute(82)
	Qt__WA_X11OpenGLOverlay                Qt__WidgetAttribute = Qt__WidgetAttribute(83)
	Qt__WA_AlwaysShowToolTips              Qt__WidgetAttribute = Qt__WidgetAttribute(84)
	Qt__WA_MacOpaqueSizeGrip               Qt__WidgetAttribute = Qt__WidgetAttribute(85)
	Qt__WA_SetStyle                        Qt__WidgetAttribute = Qt__WidgetAttribute(86)
	Qt__WA_SetLocale                       Qt__WidgetAttribute = Qt__WidgetAttribute(87)
	Qt__WA_MacShowFocusRect                Qt__WidgetAttribute = Qt__WidgetAttribute(88)
	Qt__WA_MacNormalSize                   Qt__WidgetAttribute = Qt__WidgetAttribute(89)
	Qt__WA_MacSmallSize                    Qt__WidgetAttribute = Qt__WidgetAttribute(90)
	Qt__WA_MacMiniSize                     Qt__WidgetAttribute = Qt__WidgetAttribute(91)
	Qt__WA_LayoutUsesWidgetRect            Qt__WidgetAttribute = Qt__WidgetAttribute(92)
	Qt__WA_StyledBackground                Qt__WidgetAttribute = Qt__WidgetAttribute(93)
	Qt__WA_MSWindowsUseDirect3D            Qt__WidgetAttribute = Qt__WidgetAttribute(94)
	Qt__WA_CanHostQMdiSubWindowTitleBar    Qt__WidgetAttribute = Qt__WidgetAttribute(95)
	Qt__WA_MacAlwaysShowToolWindow         Qt__WidgetAttribute = Qt__WidgetAttribute(96)
	Qt__WA_StyleSheet                      Qt__WidgetAttribute = Qt__WidgetAttribute(97)
	Qt__WA_ShowWithoutActivating           Qt__WidgetAttribute = Qt__WidgetAttribute(98)
	Qt__WA_X11BypassTransientForHint       Qt__WidgetAttribute = Qt__WidgetAttribute(99)
	Qt__WA_NativeWindow                    Qt__WidgetAttribute = Qt__WidgetAttribute(100)
	Qt__WA_DontCreateNativeAncestors       Qt__WidgetAttribute = Qt__WidgetAttribute(101)
	Qt__WA_MacVariableSize                 Qt__WidgetAttribute = Qt__WidgetAttribute(102)
	Qt__WA_DontShowOnScreen                Qt__WidgetAttribute = Qt__WidgetAttribute(103)
	Qt__WA_X11NetWmWindowTypeDesktop       Qt__WidgetAttribute = Qt__WidgetAttribute(104)
	Qt__WA_X11NetWmWindowTypeDock          Qt__WidgetAttribute = Qt__WidgetAttribute(105)
	Qt__WA_X11NetWmWindowTypeToolBar       Qt__WidgetAttribute = Qt__WidgetAttribute(106)
	Qt__WA_X11NetWmWindowTypeMenu          Qt__WidgetAttribute = Qt__WidgetAttribute(107)
	Qt__WA_X11NetWmWindowTypeUtility       Qt__WidgetAttribute = Qt__WidgetAttribute(108)
	Qt__WA_X11NetWmWindowTypeSplash        Qt__WidgetAttribute = Qt__WidgetAttribute(109)
	Qt__WA_X11NetWmWindowTypeDialog        Qt__WidgetAttribute = Qt__WidgetAttribute(110)
	Qt__WA_X11NetWmWindowTypeDropDownMenu  Qt__WidgetAttribute = Qt__WidgetAttribute(111)
	Qt__WA_X11NetWmWindowTypePopupMenu     Qt__WidgetAttribute = Qt__WidgetAttribute(112)
	Qt__WA_X11NetWmWindowTypeToolTip       Qt__WidgetAttribute = Qt__WidgetAttribute(113)
	Qt__WA_X11NetWmWindowTypeNotification  Qt__WidgetAttribute = Qt__WidgetAttribute(114)
	Qt__WA_X11NetWmWindowTypeCombo         Qt__WidgetAttribute = Qt__WidgetAttribute(115)
	Qt__WA_X11NetWmWindowTypeDND           Qt__WidgetAttribute = Qt__WidgetAttribute(116)
	Qt__WA_MacFrameworkScaled              Qt__WidgetAttribute = Qt__WidgetAttribute(117)
	Qt__WA_SetWindowModality               Qt__WidgetAttribute = Qt__WidgetAttribute(118)
	Qt__WA_WState_WindowOpacitySet         Qt__WidgetAttribute = Qt__WidgetAttribute(119)
	Qt__WA_TranslucentBackground           Qt__WidgetAttribute = Qt__WidgetAttribute(120)
	Qt__WA_AcceptTouchEvents               Qt__WidgetAttribute = Qt__WidgetAttribute(121)
	Qt__WA_WState_AcceptedTouchBeginEvent  Qt__WidgetAttribute = Qt__WidgetAttribute(122)
	Qt__WA_TouchPadAcceptSingleTouchEvents Qt__WidgetAttribute = Qt__WidgetAttribute(123)
	Qt__WA_X11DoNotAcceptFocus             Qt__WidgetAttribute = Qt__WidgetAttribute(126)
	Qt__WA_MacNoShadow                     Qt__WidgetAttribute = Qt__WidgetAttribute(127)
	Qt__WA_AlwaysStackOnTop                Qt__WidgetAttribute = Qt__WidgetAttribute(128)
	Qt__WA_TabletTracking                  Qt__WidgetAttribute = Qt__WidgetAttribute(129)
	Qt__WA_ContentsMarginsRespectsSafeArea Qt__WidgetAttribute = Qt__WidgetAttribute(130)
	Qt__WA_StyleSheetTarget                Qt__WidgetAttribute = Qt__WidgetAttribute(131)
	Qt__WA_AttributeCount                  Qt__WidgetAttribute = Qt__WidgetAttribute(132)
)

// Qt::WindowFrameSection
//
//go:generate stringer -type=Qt__WindowFrameSection
type Qt__WindowFrameSection int64

const (
	Qt__NoSection          Qt__WindowFrameSection = Qt__WindowFrameSection(0)
	Qt__LeftSection        Qt__WindowFrameSection = Qt__WindowFrameSection(1)
	Qt__TopLeftSection     Qt__WindowFrameSection = Qt__WindowFrameSection(2)
	Qt__TopSection         Qt__WindowFrameSection = Qt__WindowFrameSection(3)
	Qt__TopRightSection    Qt__WindowFrameSection = Qt__WindowFrameSection(4)
	Qt__RightSection       Qt__WindowFrameSection = Qt__WindowFrameSection(5)
	Qt__BottomRightSection Qt__WindowFrameSection = Qt__WindowFrameSection(6)
	Qt__BottomSection      Qt__WindowFrameSection = Qt__WindowFrameSection(7)
	Qt__BottomLeftSection  Qt__WindowFrameSection = Qt__WindowFrameSection(8)
	Qt__TitleBarArea       Qt__WindowFrameSection = Qt__WindowFrameSection(9)
)

// Qt::WindowModality
//
//go:generate stringer -type=Qt__WindowModality
type Qt__WindowModality int64

const (
	Qt__NonModal         Qt__WindowModality = Qt__WindowModality(0)
	Qt__WindowModal      Qt__WindowModality = Qt__WindowModality(1)
	Qt__ApplicationModal Qt__WindowModality = Qt__WindowModality(2)
)

// Qt::WindowState
//
//go:generate stringer -type=Qt__WindowState
type Qt__WindowState int64

const (
	Qt__WindowNoState    Qt__WindowState = Qt__WindowState(0x00000000)
	Qt__WindowMinimized  Qt__WindowState = Qt__WindowState(0x00000001)
	Qt__WindowMaximized  Qt__WindowState = Qt__WindowState(0x00000002)
	Qt__WindowFullScreen Qt__WindowState = Qt__WindowState(0x00000004)
	Qt__WindowActive     Qt__WindowState = Qt__WindowState(0x00000008)
)

// Qt::WindowType
//
//go:generate stringer -type=Qt__WindowType
type Qt__WindowType int64

const (
	Qt__Widget                              Qt__WindowType = Qt__WindowType(0x00000000)
	Qt__Window                              Qt__WindowType = Qt__WindowType(0x00000001)
	Qt__Dialog                              Qt__WindowType = Qt__WindowType(0x00000002 | Qt__Window)
	Qt__Sheet                               Qt__WindowType = Qt__WindowType(0x00000004 | Qt__Window)
	Qt__Drawer                              Qt__WindowType = Qt__WindowType(Qt__Sheet | Qt__Dialog)
	Qt__Popup                               Qt__WindowType = Qt__WindowType(0x00000008 | Qt__Window)
	Qt__Tool                                Qt__WindowType = Qt__WindowType(Qt__Popup | Qt__Dialog)
	Qt__ToolTip                             Qt__WindowType = Qt__WindowType(Qt__Popup | Qt__Sheet)
	Qt__SplashScreen                        Qt__WindowType = Qt__WindowType(Qt__ToolTip | Qt__Dialog)
	Qt__Desktop                             Qt__WindowType = Qt__WindowType(0x00000010 | Qt__Window)
	Qt__SubWindow                           Qt__WindowType = Qt__WindowType(0x00000012)
	Qt__ForeignWindow                       Qt__WindowType = Qt__WindowType(0x00000020 | Qt__Window)
	Qt__CoverWindow                         Qt__WindowType = Qt__WindowType(0x00000040 | Qt__Window)
	Qt__WindowType_Mask                     Qt__WindowType = Qt__WindowType(0x000000ff)
	Qt__MSWindowsFixedSizeDialogHint        Qt__WindowType = Qt__WindowType(0x00000100)
	Qt__MSWindowsOwnDC                      Qt__WindowType = Qt__WindowType(0x00000200)
	Qt__BypassWindowManagerHint             Qt__WindowType = Qt__WindowType(0x00000400)
	Qt__X11BypassWindowManagerHint          Qt__WindowType = Qt__WindowType(Qt__BypassWindowManagerHint)
	Qt__FramelessWindowHint                 Qt__WindowType = Qt__WindowType(0x00000800)
	Qt__WindowTitleHint                     Qt__WindowType = Qt__WindowType(0x00001000)
	Qt__WindowSystemMenuHint                Qt__WindowType = Qt__WindowType(0x00002000)
	Qt__WindowMinimizeButtonHint            Qt__WindowType = Qt__WindowType(0x00004000)
	Qt__WindowMaximizeButtonHint            Qt__WindowType = Qt__WindowType(0x00008000)
	Qt__WindowMinMaxButtonsHint             Qt__WindowType = Qt__WindowType(Qt__WindowMinimizeButtonHint | Qt__WindowMaximizeButtonHint)
	Qt__WindowContextHelpButtonHint         Qt__WindowType = Qt__WindowType(0x00010000)
	Qt__WindowShadeButtonHint               Qt__WindowType = Qt__WindowType(0x00020000)
	Qt__WindowStaysOnTopHint                Qt__WindowType = Qt__WindowType(0x00040000)
	Qt__WindowTransparentForInput           Qt__WindowType = Qt__WindowType(0x00080000)
	Qt__WindowOverridesSystemGestures       Qt__WindowType = Qt__WindowType(0x00100000)
	Qt__WindowDoesNotAcceptFocus            Qt__WindowType = Qt__WindowType(0x00200000)
	Qt__MaximizeUsingFullscreenGeometryHint Qt__WindowType = Qt__WindowType(0x00400000)
	Qt__CustomizeWindowHint                 Qt__WindowType = Qt__WindowType(0x02000000)
	Qt__WindowStaysOnBottomHint             Qt__WindowType = Qt__WindowType(0x04000000)
	Qt__WindowCloseButtonHint               Qt__WindowType = Qt__WindowType(0x08000000)
	Qt__MacWindowToolBarButtonHint          Qt__WindowType = Qt__WindowType(0x10000000)
	Qt__BypassGraphicsProxyWidget           Qt__WindowType = Qt__WindowType(0x20000000)
	Qt__NoDropShadowWindowHint              Qt__WindowType = Qt__WindowType(0x40000000)
	Qt__WindowFullscreenButtonHint          Qt__WindowType = Qt__WindowType(0x80000000)
)

func init() {
}
