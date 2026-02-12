//go:build minimal
// +build minimal

package widgets

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "widgets-minimal.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtWidgets_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtWidgets_PackedString) []byte {
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

type QAbstractButton struct {
	QWidget
}

type QAbstractButton_ITF interface {
	QWidget_ITF
	QAbstractButton_PTR() *QAbstractButton
}

func (ptr *QAbstractButton) QAbstractButton_PTR() *QAbstractButton {
	return ptr
}

func (ptr *QAbstractButton) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractButton) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractButton(ptr QAbstractButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractButton_PTR().Pointer()
	}
	return nil
}

func NewQAbstractButtonFromPointer(ptr unsafe.Pointer) (n *QAbstractButton) {
	n = new(QAbstractButton)
	n.SetPointer(ptr)
	return
}
func NewQAbstractButton(parent QWidget_ITF) *QAbstractButton {
	tmpValue := NewQAbstractButtonFromPointer(C.QAbstractButton_NewQAbstractButton(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAbstractButton_Click
func callbackQAbstractButton_Click(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "click"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractButtonFromPointer(ptr).ClickDefault()
	}
}

func (ptr *QAbstractButton) ConnectClick(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "click"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "click", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "click", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractButton) DisconnectClick() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "click")
	}
}

func (ptr *QAbstractButton) Click() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_Click(ptr.Pointer())
	}
}

func (ptr *QAbstractButton) ClickDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_ClickDefault(ptr.Pointer())
	}
}

//export callbackQAbstractButton_Clicked
func callbackQAbstractButton_Clicked(ptr unsafe.Pointer, checked C.char) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(bool))(signal))(int8(checked) != 0)
	}

}

func (ptr *QAbstractButton) ConnectClicked(f func(checked bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QAbstractButton_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(checked bool) {
				(*(*func(bool))(signal))(checked)
				f(checked)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractButton) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QAbstractButton) Clicked(checked bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_Clicked(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

func (ptr *QAbstractButton) Group() *QButtonGroup {
	if ptr.Pointer() != nil {
		tmpValue := NewQButtonGroupFromPointer(C.QAbstractButton_Group(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractButton) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QAbstractButton_Icon(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractButton) IsChecked() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractButton_IsChecked(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQAbstractButton_PaintEvent
func callbackQAbstractButton_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*gui.QPaintEvent))(signal))(gui.NewQPaintEventFromPointer(e))
	}

}

func (ptr *QAbstractButton) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paintEvent"); signal != nil {
			f := func(e *gui.QPaintEvent) {
				(*(*func(*gui.QPaintEvent))(signal))(e)
				f(e)
			}
			qt.ConnectSignal(ptr.Pointer(), "paintEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paintEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractButton) DisconnectPaintEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paintEvent")
	}
}

func (ptr *QAbstractButton) PaintEvent(e gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QAbstractButton) SetCheckable(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetCheckable(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQAbstractButton_SetChecked
func callbackQAbstractButton_SetChecked(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setChecked"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewQAbstractButtonFromPointer(ptr).SetCheckedDefault(int8(vbo) != 0)
	}
}

func (ptr *QAbstractButton) ConnectSetChecked(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setChecked"); signal != nil {
			f := func(vbo bool) {
				(*(*func(bool))(signal))(vbo)
				f(vbo)
			}
			qt.ConnectSignal(ptr.Pointer(), "setChecked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setChecked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractButton) DisconnectSetChecked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setChecked")
	}
}

func (ptr *QAbstractButton) SetChecked(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetChecked(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QAbstractButton) SetCheckedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetCheckedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QAbstractButton) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QAbstractButton) SetShortcut(key gui.QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractButton_SetShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(key))
	}
}

func (ptr *QAbstractButton) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QAbstractButton_SetText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QAbstractButton) Shortcut() *gui.QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQKeySequenceFromPointer(C.QAbstractButton_Shortcut(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractButton) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractButton_Text(ptr.Pointer()))
	}
	return ""
}

//export callbackQAbstractButton_DestroyQAbstractButton
func callbackQAbstractButton_DestroyQAbstractButton(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractButton"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractButtonFromPointer(ptr).DestroyQAbstractButtonDefault()
	}
}

func (ptr *QAbstractButton) ConnectDestroyQAbstractButton(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractButton"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractButton", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractButton", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractButton) DisconnectDestroyQAbstractButton() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractButton")
	}
}

func (ptr *QAbstractButton) DestroyQAbstractButton() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractButton_DestroyQAbstractButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractButton) DestroyQAbstractButtonDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractButton_DestroyQAbstractButtonDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QAbstractItemDelegate::EndEditHint
//
//go:generate stringer -type=QAbstractItemDelegate__EndEditHint
type QAbstractItemDelegate__EndEditHint int64

const (
	QAbstractItemDelegate__NoHint           QAbstractItemDelegate__EndEditHint = QAbstractItemDelegate__EndEditHint(0)
	QAbstractItemDelegate__EditNextItem     QAbstractItemDelegate__EndEditHint = QAbstractItemDelegate__EndEditHint(1)
	QAbstractItemDelegate__EditPreviousItem QAbstractItemDelegate__EndEditHint = QAbstractItemDelegate__EndEditHint(2)
	QAbstractItemDelegate__SubmitModelCache QAbstractItemDelegate__EndEditHint = QAbstractItemDelegate__EndEditHint(3)
	QAbstractItemDelegate__RevertModelCache QAbstractItemDelegate__EndEditHint = QAbstractItemDelegate__EndEditHint(4)
)

type QAbstractItemView struct {
	QAbstractScrollArea
}

type QAbstractItemView_ITF interface {
	QAbstractScrollArea_ITF
	QAbstractItemView_PTR() *QAbstractItemView
}

func (ptr *QAbstractItemView) QAbstractItemView_PTR() *QAbstractItemView {
	return ptr
}

func (ptr *QAbstractItemView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractItemView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractScrollArea_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractItemView(ptr QAbstractItemView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemView_PTR().Pointer()
	}
	return nil
}

func NewQAbstractItemViewFromPointer(ptr unsafe.Pointer) (n *QAbstractItemView) {
	n = new(QAbstractItemView)
	n.SetPointer(ptr)
	return
}

// QAbstractItemView::CursorAction
//
//go:generate stringer -type=QAbstractItemView__CursorAction
type QAbstractItemView__CursorAction int64

const (
	QAbstractItemView__MoveUp       QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(0)
	QAbstractItemView__MoveDown     QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(1)
	QAbstractItemView__MoveLeft     QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(2)
	QAbstractItemView__MoveRight    QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(3)
	QAbstractItemView__MoveHome     QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(4)
	QAbstractItemView__MoveEnd      QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(5)
	QAbstractItemView__MovePageUp   QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(6)
	QAbstractItemView__MovePageDown QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(7)
	QAbstractItemView__MoveNext     QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(8)
	QAbstractItemView__MovePrevious QAbstractItemView__CursorAction = QAbstractItemView__CursorAction(9)
)

// QAbstractItemView::DragDropMode
//
//go:generate stringer -type=QAbstractItemView__DragDropMode
type QAbstractItemView__DragDropMode int64

const (
	QAbstractItemView__NoDragDrop   QAbstractItemView__DragDropMode = QAbstractItemView__DragDropMode(0)
	QAbstractItemView__DragOnly     QAbstractItemView__DragDropMode = QAbstractItemView__DragDropMode(1)
	QAbstractItemView__DropOnly     QAbstractItemView__DragDropMode = QAbstractItemView__DragDropMode(2)
	QAbstractItemView__DragDrop     QAbstractItemView__DragDropMode = QAbstractItemView__DragDropMode(3)
	QAbstractItemView__InternalMove QAbstractItemView__DragDropMode = QAbstractItemView__DragDropMode(4)
)

// QAbstractItemView::DropIndicatorPosition
//
//go:generate stringer -type=QAbstractItemView__DropIndicatorPosition
type QAbstractItemView__DropIndicatorPosition int64

const (
	QAbstractItemView__OnItem     QAbstractItemView__DropIndicatorPosition = QAbstractItemView__DropIndicatorPosition(0)
	QAbstractItemView__AboveItem  QAbstractItemView__DropIndicatorPosition = QAbstractItemView__DropIndicatorPosition(1)
	QAbstractItemView__BelowItem  QAbstractItemView__DropIndicatorPosition = QAbstractItemView__DropIndicatorPosition(2)
	QAbstractItemView__OnViewport QAbstractItemView__DropIndicatorPosition = QAbstractItemView__DropIndicatorPosition(3)
)

// QAbstractItemView::EditTrigger
//
//go:generate stringer -type=QAbstractItemView__EditTrigger
type QAbstractItemView__EditTrigger int64

const (
	QAbstractItemView__NoEditTriggers  QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(0)
	QAbstractItemView__CurrentChanged  QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(1)
	QAbstractItemView__DoubleClicked   QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(2)
	QAbstractItemView__SelectedClicked QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(4)
	QAbstractItemView__EditKeyPressed  QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(8)
	QAbstractItemView__AnyKeyPressed   QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(16)
	QAbstractItemView__AllEditTriggers QAbstractItemView__EditTrigger = QAbstractItemView__EditTrigger(31)
)

// QAbstractItemView::ScrollHint
//
//go:generate stringer -type=QAbstractItemView__ScrollHint
type QAbstractItemView__ScrollHint int64

const (
	QAbstractItemView__EnsureVisible    QAbstractItemView__ScrollHint = QAbstractItemView__ScrollHint(0)
	QAbstractItemView__PositionAtTop    QAbstractItemView__ScrollHint = QAbstractItemView__ScrollHint(1)
	QAbstractItemView__PositionAtBottom QAbstractItemView__ScrollHint = QAbstractItemView__ScrollHint(2)
	QAbstractItemView__PositionAtCenter QAbstractItemView__ScrollHint = QAbstractItemView__ScrollHint(3)
)

// QAbstractItemView::ScrollMode
//
//go:generate stringer -type=QAbstractItemView__ScrollMode
type QAbstractItemView__ScrollMode int64

const (
	QAbstractItemView__ScrollPerItem  QAbstractItemView__ScrollMode = QAbstractItemView__ScrollMode(0)
	QAbstractItemView__ScrollPerPixel QAbstractItemView__ScrollMode = QAbstractItemView__ScrollMode(1)
)

// QAbstractItemView::SelectionBehavior
//
//go:generate stringer -type=QAbstractItemView__SelectionBehavior
type QAbstractItemView__SelectionBehavior int64

const (
	QAbstractItemView__SelectItems   QAbstractItemView__SelectionBehavior = QAbstractItemView__SelectionBehavior(0)
	QAbstractItemView__SelectRows    QAbstractItemView__SelectionBehavior = QAbstractItemView__SelectionBehavior(1)
	QAbstractItemView__SelectColumns QAbstractItemView__SelectionBehavior = QAbstractItemView__SelectionBehavior(2)
)

// QAbstractItemView::SelectionMode
//
//go:generate stringer -type=QAbstractItemView__SelectionMode
type QAbstractItemView__SelectionMode int64

const (
	QAbstractItemView__NoSelection         QAbstractItemView__SelectionMode = QAbstractItemView__SelectionMode(0)
	QAbstractItemView__SingleSelection     QAbstractItemView__SelectionMode = QAbstractItemView__SelectionMode(1)
	QAbstractItemView__MultiSelection      QAbstractItemView__SelectionMode = QAbstractItemView__SelectionMode(2)
	QAbstractItemView__ExtendedSelection   QAbstractItemView__SelectionMode = QAbstractItemView__SelectionMode(3)
	QAbstractItemView__ContiguousSelection QAbstractItemView__SelectionMode = QAbstractItemView__SelectionMode(4)
)

// QAbstractItemView::State
//
//go:generate stringer -type=QAbstractItemView__State
type QAbstractItemView__State int64

const (
	QAbstractItemView__NoState            QAbstractItemView__State = QAbstractItemView__State(0)
	QAbstractItemView__DraggingState      QAbstractItemView__State = QAbstractItemView__State(1)
	QAbstractItemView__DragSelectingState QAbstractItemView__State = QAbstractItemView__State(2)
	QAbstractItemView__EditingState       QAbstractItemView__State = QAbstractItemView__State(3)
	QAbstractItemView__ExpandingState     QAbstractItemView__State = QAbstractItemView__State(4)
	QAbstractItemView__CollapsingState    QAbstractItemView__State = QAbstractItemView__State(5)
	QAbstractItemView__AnimatingState     QAbstractItemView__State = QAbstractItemView__State(6)
)

func NewQAbstractItemView(parent QWidget_ITF) *QAbstractItemView {
	tmpValue := NewQAbstractItemViewFromPointer(C.QAbstractItemView_NewQAbstractItemView(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAbstractItemView_Activated
func callbackQAbstractItemView_Activated(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QAbstractItemView) ConnectActivated(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activated") {
			C.QAbstractItemView_ConnectActivated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "activated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activated"); signal != nil {
			f := func(index *core.QModelIndex) {
				(*(*func(*core.QModelIndex))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "activated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activated", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activated")
	}
}

func (ptr *QAbstractItemView) Activated(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_Activated(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQAbstractItemView_ClearSelection
func callbackQAbstractItemView_ClearSelection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearSelection"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractItemViewFromPointer(ptr).ClearSelectionDefault()
	}
}

func (ptr *QAbstractItemView) ConnectClearSelection(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clearSelection"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clearSelection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clearSelection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectClearSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clearSelection")
	}
}

func (ptr *QAbstractItemView) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QAbstractItemView) ClearSelectionDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ClearSelectionDefault(ptr.Pointer())
	}
}

//export callbackQAbstractItemView_Clicked
func callbackQAbstractItemView_Clicked(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QAbstractItemView) ConnectClicked(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QAbstractItemView_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func(index *core.QModelIndex) {
				(*(*func(*core.QModelIndex))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QAbstractItemView) Clicked(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_Clicked(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) CurrentIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QAbstractItemView_CurrentIndex(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractItemView_Edit
func callbackQAbstractItemView_Edit(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "edit"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQAbstractItemViewFromPointer(ptr).EditDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QAbstractItemView) ConnectEdit(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "edit"); signal != nil {
			f := func(index *core.QModelIndex) {
				(*(*func(*core.QModelIndex))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "edit", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "edit", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectEdit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "edit")
	}
}

func (ptr *QAbstractItemView) Edit(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_Edit(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) EditDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_EditDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQAbstractItemView_Edit2
func callbackQAbstractItemView_Edit2(ptr unsafe.Pointer, index unsafe.Pointer, trigger C.longlong, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "edit2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, QAbstractItemView__EditTrigger, *core.QEvent) bool)(signal))(core.NewQModelIndexFromPointer(index), QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractItemViewFromPointer(ptr).Edit2Default(core.NewQModelIndexFromPointer(index), QAbstractItemView__EditTrigger(trigger), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractItemView) ConnectEdit2(f func(index *core.QModelIndex, trigger QAbstractItemView__EditTrigger, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "edit2"); signal != nil {
			f := func(index *core.QModelIndex, trigger QAbstractItemView__EditTrigger, event *core.QEvent) bool {
				(*(*func(*core.QModelIndex, QAbstractItemView__EditTrigger, *core.QEvent) bool)(signal))(index, trigger, event)
				return f(index, trigger, event)
			}
			qt.ConnectSignal(ptr.Pointer(), "edit2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "edit2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectEdit2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "edit2")
	}
}

func (ptr *QAbstractItemView) Edit2(index core.QModelIndex_ITF, trigger QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemView_Edit2(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QAbstractItemView) Edit2Default(index core.QModelIndex_ITF, trigger QAbstractItemView__EditTrigger, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemView_Edit2Default(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(trigger), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAbstractItemView_HorizontalOffset
func callbackQAbstractItemView_HorizontalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "horizontalOffset"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(0))
}

func (ptr *QAbstractItemView) ConnectHorizontalOffset(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "horizontalOffset"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "horizontalOffset", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "horizontalOffset", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectHorizontalOffset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "horizontalOffset")
	}
}

func (ptr *QAbstractItemView) HorizontalOffset() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemView_HorizontalOffset(ptr.Pointer())))
	}
	return 0
}

//export callbackQAbstractItemView_IndexAt
func callbackQAbstractItemView_IndexAt(ptr unsafe.Pointer, point unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "indexAt"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QPoint) *core.QModelIndex)(signal))(core.NewQPointFromPointer(point)))
	}

	return core.PointerFromQModelIndex(core.NewQModelIndex())
}

func (ptr *QAbstractItemView) ConnectIndexAt(f func(point *core.QPoint) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "indexAt"); signal != nil {
			f := func(point *core.QPoint) *core.QModelIndex {
				(*(*func(*core.QPoint) *core.QModelIndex)(signal))(point)
				return f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "indexAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectIndexAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "indexAt")
	}
}

func (ptr *QAbstractItemView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QAbstractItemView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractItemView_IsIndexHidden
func callbackQAbstractItemView_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAbstractItemView) ConnectIsIndexHidden(f func(index *core.QModelIndex) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isIndexHidden"); signal != nil {
			f := func(index *core.QModelIndex) bool {
				(*(*func(*core.QModelIndex) bool)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "isIndexHidden", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isIndexHidden", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectIsIndexHidden() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isIndexHidden")
	}
}

func (ptr *QAbstractItemView) IsIndexHidden(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractItemView_IsIndexHidden(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QAbstractItemView) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QAbstractItemView_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractItemView_MoveCursor
func callbackQAbstractItemView_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "moveCursor"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(signal))(QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(core.NewQModelIndex())
}

func (ptr *QAbstractItemView) ConnectMoveCursor(f func(cursorAction QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "moveCursor"); signal != nil {
			f := func(cursorAction QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
				(*(*func(QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(signal))(cursorAction, modifiers)
				return f(cursorAction, modifiers)
			}
			qt.ConnectSignal(ptr.Pointer(), "moveCursor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "moveCursor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectMoveCursor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "moveCursor")
	}
}

func (ptr *QAbstractItemView) MoveCursor(cursorAction QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QAbstractItemView_MoveCursor(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemView) RootIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QAbstractItemView_RootIndex(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractItemView_ScrollTo
func callbackQAbstractItemView_ScrollTo(ptr unsafe.Pointer, index unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "scrollTo"); signal != nil {
		(*(*func(*core.QModelIndex, QAbstractItemView__ScrollHint))(signal))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	}

}

func (ptr *QAbstractItemView) ConnectScrollTo(f func(index *core.QModelIndex, hint QAbstractItemView__ScrollHint)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "scrollTo"); signal != nil {
			f := func(index *core.QModelIndex, hint QAbstractItemView__ScrollHint) {
				(*(*func(*core.QModelIndex, QAbstractItemView__ScrollHint))(signal))(index, hint)
				f(index, hint)
			}
			qt.ConnectSignal(ptr.Pointer(), "scrollTo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scrollTo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectScrollTo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "scrollTo")
	}
}

func (ptr *QAbstractItemView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackQAbstractItemView_SelectionChanged
func callbackQAbstractItemView_SelectionChanged(ptr unsafe.Pointer, selected unsafe.Pointer, deselected unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		(*(*func(*core.QItemSelection, *core.QItemSelection))(signal))(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	} else {
		NewQAbstractItemViewFromPointer(ptr).SelectionChangedDefault(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	}
}

func (ptr *QAbstractItemView) ConnectSelectionChanged(f func(selected *core.QItemSelection, deselected *core.QItemSelection)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectionChanged"); signal != nil {
			f := func(selected *core.QItemSelection, deselected *core.QItemSelection) {
				(*(*func(*core.QItemSelection, *core.QItemSelection))(signal))(selected, deselected)
				f(selected, deselected)
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectionChanged")
	}
}

func (ptr *QAbstractItemView) SelectionChanged(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SelectionChanged(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

func (ptr *QAbstractItemView) SelectionChangedDefault(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SelectionChangedDefault(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

func (ptr *QAbstractItemView) SelectionMode() QAbstractItemView__SelectionMode {
	if ptr.Pointer() != nil {
		return QAbstractItemView__SelectionMode(C.QAbstractItemView_SelectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) SetAutoScroll(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetAutoScroll(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQAbstractItemView_SetCurrentIndex
func callbackQAbstractItemView_SetCurrentIndex(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setCurrentIndex"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQAbstractItemViewFromPointer(ptr).SetCurrentIndexDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QAbstractItemView) ConnectSetCurrentIndex(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCurrentIndex"); signal != nil {
			f := func(index *core.QModelIndex) {
				(*(*func(*core.QModelIndex))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "setCurrentIndex", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCurrentIndex", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectSetCurrentIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCurrentIndex")
	}
}

func (ptr *QAbstractItemView) SetCurrentIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetCurrentIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) SetCurrentIndexDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetCurrentIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQAbstractItemView_SetModel
func callbackQAbstractItemView_SetModel(ptr unsafe.Pointer, model unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setModel"); signal != nil {
		(*(*func(*core.QAbstractItemModel))(signal))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQAbstractItemViewFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QAbstractItemView) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setModel"); signal != nil {
			f := func(model *core.QAbstractItemModel) {
				(*(*func(*core.QAbstractItemModel))(signal))(model)
				f(model)
			}
			qt.ConnectSignal(ptr.Pointer(), "setModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectSetModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setModel")
	}
}

func (ptr *QAbstractItemView) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QAbstractItemView) SetModelDefault(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

//export callbackQAbstractItemView_SetSelection
func callbackQAbstractItemView_SetSelection(ptr unsafe.Pointer, rect unsafe.Pointer, flags C.longlong) {
	if signal := qt.GetSignal(ptr, "setSelection"); signal != nil {
		(*(*func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(signal))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(flags))
	}

}

func (ptr *QAbstractItemView) ConnectSetSelection(f func(rect *core.QRect, flags core.QItemSelectionModel__SelectionFlag)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSelection"); signal != nil {
			f := func(rect *core.QRect, flags core.QItemSelectionModel__SelectionFlag) {
				(*(*func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(signal))(rect, flags)
				f(rect, flags)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSelection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSelection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectSetSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSelection")
	}
}

func (ptr *QAbstractItemView) SetSelection(rect core.QRect_ITF, flags core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(flags))
	}
}

func (ptr *QAbstractItemView) SetSelectionMode(mode QAbstractItemView__SelectionMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetSelectionMode(ptr.Pointer(), C.longlong(mode))
	}
}

//export callbackQAbstractItemView_Update
func callbackQAbstractItemView_Update(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQAbstractItemViewFromPointer(ptr).UpdateDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QAbstractItemView) ConnectUpdate(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "update"); signal != nil {
			f := func(index *core.QModelIndex) {
				(*(*func(*core.QModelIndex))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "update", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "update", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "update")
	}
}

func (ptr *QAbstractItemView) Update(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_Update(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) UpdateDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_UpdateDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQAbstractItemView_VerticalOffset
func callbackQAbstractItemView_VerticalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "verticalOffset"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(0))
}

func (ptr *QAbstractItemView) ConnectVerticalOffset(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "verticalOffset"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "verticalOffset", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "verticalOffset", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectVerticalOffset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "verticalOffset")
	}
}

func (ptr *QAbstractItemView) VerticalOffset() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemView_VerticalOffset(ptr.Pointer())))
	}
	return 0
}

//export callbackQAbstractItemView_VisualRect
func callbackQAbstractItemView_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRect"); signal != nil {
		return core.PointerFromQRect((*(*func(*core.QModelIndex) *core.QRect)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(core.NewQRect())
}

func (ptr *QAbstractItemView) ConnectVisualRect(f func(index *core.QModelIndex) *core.QRect) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "visualRect"); signal != nil {
			f := func(index *core.QModelIndex) *core.QRect {
				(*(*func(*core.QModelIndex) *core.QRect)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "visualRect", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "visualRect", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectVisualRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "visualRect")
	}
}

func (ptr *QAbstractItemView) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QAbstractItemView_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractItemView_VisualRegionForSelection
func callbackQAbstractItemView_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion((*(*func(*core.QItemSelection) *gui.QRegion)(signal))(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(gui.NewQRegion())
}

func (ptr *QAbstractItemView) ConnectVisualRegionForSelection(f func(selection *core.QItemSelection) *gui.QRegion) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "visualRegionForSelection"); signal != nil {
			f := func(selection *core.QItemSelection) *gui.QRegion {
				(*(*func(*core.QItemSelection) *gui.QRegion)(signal))(selection)
				return f(selection)
			}
			qt.ConnectSignal(ptr.Pointer(), "visualRegionForSelection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "visualRegionForSelection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectVisualRegionForSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "visualRegionForSelection")
	}
}

func (ptr *QAbstractItemView) VisualRegionForSelection(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQRegionFromPointer(C.QAbstractItemView_VisualRegionForSelection(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		qt.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractItemView_DestroyQAbstractItemView
func callbackQAbstractItemView_DestroyQAbstractItemView(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractItemView"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractItemViewFromPointer(ptr).DestroyQAbstractItemViewDefault()
	}
}

func (ptr *QAbstractItemView) ConnectDestroyQAbstractItemView(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractItemView"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractItemView", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractItemView", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractItemView) DisconnectDestroyQAbstractItemView() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractItemView")
	}
}

func (ptr *QAbstractItemView) DestroyQAbstractItemView() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractItemView_DestroyQAbstractItemView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractItemView) DestroyQAbstractItemViewDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractItemView_DestroyQAbstractItemViewDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractItemView) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractItemView___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QAbstractItemView) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QAbstractItemView) __dataChanged_roles_newList() unsafe.Pointer {
	return C.QAbstractItemView___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *QAbstractItemView) __selectedIndexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QAbstractItemView___selectedIndexes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractItemView) __selectedIndexes_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView___selectedIndexes_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QAbstractItemView) __selectedIndexes_newList() unsafe.Pointer {
	return C.QAbstractItemView___selectedIndexes_newList(ptr.Pointer())
}

type QAbstractScrollArea struct {
	QFrame
}

type QAbstractScrollArea_ITF interface {
	QFrame_ITF
	QAbstractScrollArea_PTR() *QAbstractScrollArea
}

func (ptr *QAbstractScrollArea) QAbstractScrollArea_PTR() *QAbstractScrollArea {
	return ptr
}

func (ptr *QAbstractScrollArea) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractScrollArea) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QFrame_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractScrollArea(ptr QAbstractScrollArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func NewQAbstractScrollAreaFromPointer(ptr unsafe.Pointer) (n *QAbstractScrollArea) {
	n = new(QAbstractScrollArea)
	n.SetPointer(ptr)
	return
}

// QAbstractScrollArea::SizeAdjustPolicy
//
//go:generate stringer -type=QAbstractScrollArea__SizeAdjustPolicy
type QAbstractScrollArea__SizeAdjustPolicy int64

const (
	QAbstractScrollArea__AdjustIgnored               QAbstractScrollArea__SizeAdjustPolicy = QAbstractScrollArea__SizeAdjustPolicy(0)
	QAbstractScrollArea__AdjustToContentsOnFirstShow QAbstractScrollArea__SizeAdjustPolicy = QAbstractScrollArea__SizeAdjustPolicy(1)
	QAbstractScrollArea__AdjustToContents            QAbstractScrollArea__SizeAdjustPolicy = QAbstractScrollArea__SizeAdjustPolicy(2)
)

func NewQAbstractScrollArea(parent QWidget_ITF) *QAbstractScrollArea {
	tmpValue := NewQAbstractScrollAreaFromPointer(C.QAbstractScrollArea_NewQAbstractScrollArea(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	if ptr.Pointer() != nil {
		tmpValue := NewQScrollBarFromPointer(C.QAbstractScrollArea_HorizontalScrollBar(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractScrollArea) HorizontalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_HorizontalScrollBarPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBar(scrollBar QScrollBar_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBar(ptr.Pointer(), PointerFromQScrollBar(scrollBar))
	}
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBarPolicy(vqt core.Qt__ScrollBarPolicy) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBarPolicy(ptr.Pointer(), C.longlong(vqt))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBar(scrollBar QScrollBar_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBar(ptr.Pointer(), PointerFromQScrollBar(scrollBar))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBarPolicy(vqt core.Qt__ScrollBarPolicy) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBarPolicy(ptr.Pointer(), C.longlong(vqt))
	}
}

func (ptr *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	if ptr.Pointer() != nil {
		tmpValue := NewQScrollBarFromPointer(C.QAbstractScrollArea_VerticalScrollBar(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractScrollArea) VerticalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_VerticalScrollBarPolicy(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractScrollArea_DestroyQAbstractScrollArea
func callbackQAbstractScrollArea_DestroyQAbstractScrollArea(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractScrollArea"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractScrollAreaFromPointer(ptr).DestroyQAbstractScrollAreaDefault()
	}
}

func (ptr *QAbstractScrollArea) ConnectDestroyQAbstractScrollArea(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractScrollArea"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractScrollArea", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractScrollArea", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractScrollArea) DisconnectDestroyQAbstractScrollArea() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractScrollArea")
	}
}

func (ptr *QAbstractScrollArea) DestroyQAbstractScrollArea() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractScrollArea_DestroyQAbstractScrollArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractScrollArea) DestroyQAbstractScrollAreaDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractScrollArea_DestroyQAbstractScrollAreaDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractScrollArea) __scrollBarWidgets_atList(i int) *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QAbstractScrollArea___scrollBarWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractScrollArea) __scrollBarWidgets_setList(i QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea___scrollBarWidgets_setList(ptr.Pointer(), PointerFromQWidget(i))
	}
}

func (ptr *QAbstractScrollArea) __scrollBarWidgets_newList() unsafe.Pointer {
	return C.QAbstractScrollArea___scrollBarWidgets_newList(ptr.Pointer())
}

type QAbstractSlider struct {
	QWidget
}

type QAbstractSlider_ITF interface {
	QWidget_ITF
	QAbstractSlider_PTR() *QAbstractSlider
}

func (ptr *QAbstractSlider) QAbstractSlider_PTR() *QAbstractSlider {
	return ptr
}

func (ptr *QAbstractSlider) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractSlider) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractSlider(ptr QAbstractSlider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSlider_PTR().Pointer()
	}
	return nil
}

func NewQAbstractSliderFromPointer(ptr unsafe.Pointer) (n *QAbstractSlider) {
	n = new(QAbstractSlider)
	n.SetPointer(ptr)
	return
}

// QAbstractSlider::SliderAction
//
//go:generate stringer -type=QAbstractSlider__SliderAction
type QAbstractSlider__SliderAction int64

const (
	QAbstractSlider__SliderNoAction      QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(0)
	QAbstractSlider__SliderSingleStepAdd QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(1)
	QAbstractSlider__SliderSingleStepSub QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(2)
	QAbstractSlider__SliderPageStepAdd   QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(3)
	QAbstractSlider__SliderPageStepSub   QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(4)
	QAbstractSlider__SliderToMinimum     QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(5)
	QAbstractSlider__SliderToMaximum     QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(6)
	QAbstractSlider__SliderMove          QAbstractSlider__SliderAction = QAbstractSlider__SliderAction(7)
)

// QAbstractSlider::SliderChange
//
//go:generate stringer -type=QAbstractSlider__SliderChange
type QAbstractSlider__SliderChange int64

const (
	QAbstractSlider__SliderRangeChange       QAbstractSlider__SliderChange = QAbstractSlider__SliderChange(0)
	QAbstractSlider__SliderOrientationChange QAbstractSlider__SliderChange = QAbstractSlider__SliderChange(1)
	QAbstractSlider__SliderStepsChange       QAbstractSlider__SliderChange = QAbstractSlider__SliderChange(2)
	QAbstractSlider__SliderValueChange       QAbstractSlider__SliderChange = QAbstractSlider__SliderChange(3)
)

func NewQAbstractSlider(parent QWidget_ITF) *QAbstractSlider {
	tmpValue := NewQAbstractSliderFromPointer(C.QAbstractSlider_NewQAbstractSlider(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAbstractSlider) Maximum() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractSlider_Maximum(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSlider) Minimum() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractSlider_Minimum(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractSlider) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QAbstractSlider_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) SetMaximum(vin int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetMaximum(ptr.Pointer(), C.int(int32(vin)))
	}
}

func (ptr *QAbstractSlider) SetMinimum(vin int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetMinimum(ptr.Pointer(), C.int(int32(vin)))
	}
}

//export callbackQAbstractSlider_SetOrientation
func callbackQAbstractSlider_SetOrientation(ptr unsafe.Pointer, vqt C.longlong) {
	if signal := qt.GetSignal(ptr, "setOrientation"); signal != nil {
		(*(*func(core.Qt__Orientation))(signal))(core.Qt__Orientation(vqt))
	} else {
		NewQAbstractSliderFromPointer(ptr).SetOrientationDefault(core.Qt__Orientation(vqt))
	}
}

func (ptr *QAbstractSlider) ConnectSetOrientation(f func(vqt core.Qt__Orientation)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setOrientation"); signal != nil {
			f := func(vqt core.Qt__Orientation) {
				(*(*func(core.Qt__Orientation))(signal))(vqt)
				f(vqt)
			}
			qt.ConnectSignal(ptr.Pointer(), "setOrientation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setOrientation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSlider) DisconnectSetOrientation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setOrientation")
	}
}

func (ptr *QAbstractSlider) SetOrientation(vqt core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetOrientation(ptr.Pointer(), C.longlong(vqt))
	}
}

func (ptr *QAbstractSlider) SetOrientationDefault(vqt core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetOrientationDefault(ptr.Pointer(), C.longlong(vqt))
	}
}

//export callbackQAbstractSlider_SetRange
func callbackQAbstractSlider_SetRange(ptr unsafe.Pointer, min C.int, max C.int) {
	if signal := qt.GetSignal(ptr, "setRange"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(min)), int(int32(max)))
	} else {
		NewQAbstractSliderFromPointer(ptr).SetRangeDefault(int(int32(min)), int(int32(max)))
	}
}

func (ptr *QAbstractSlider) ConnectSetRange(f func(min int, max int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRange"); signal != nil {
			f := func(min int, max int) {
				(*(*func(int, int))(signal))(min, max)
				f(min, max)
			}
			qt.ConnectSignal(ptr.Pointer(), "setRange", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRange", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSlider) DisconnectSetRange() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRange")
	}
}

func (ptr *QAbstractSlider) SetRange(min int, max int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetRange(ptr.Pointer(), C.int(int32(min)), C.int(int32(max)))
	}
}

func (ptr *QAbstractSlider) SetRangeDefault(min int, max int) {
	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetRangeDefault(ptr.Pointer(), C.int(int32(min)), C.int(int32(max)))
	}
}

func (ptr *QAbstractSlider) Value() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAbstractSlider_Value(ptr.Pointer())))
	}
	return 0
}

//export callbackQAbstractSlider_DestroyQAbstractSlider
func callbackQAbstractSlider_DestroyQAbstractSlider(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractSlider"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractSliderFromPointer(ptr).DestroyQAbstractSliderDefault()
	}
}

func (ptr *QAbstractSlider) ConnectDestroyQAbstractSlider(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractSlider"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractSlider", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractSlider", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractSlider) DisconnectDestroyQAbstractSlider() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractSlider")
	}
}

func (ptr *QAbstractSlider) DestroyQAbstractSlider() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractSlider_DestroyQAbstractSlider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSlider) DestroyQAbstractSliderDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractSlider_DestroyQAbstractSliderDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QAbstractSpinBox::ButtonSymbols
//
//go:generate stringer -type=QAbstractSpinBox__ButtonSymbols
type QAbstractSpinBox__ButtonSymbols int64

const (
	QAbstractSpinBox__UpDownArrows QAbstractSpinBox__ButtonSymbols = QAbstractSpinBox__ButtonSymbols(0)
	QAbstractSpinBox__PlusMinus    QAbstractSpinBox__ButtonSymbols = QAbstractSpinBox__ButtonSymbols(1)
	QAbstractSpinBox__NoButtons    QAbstractSpinBox__ButtonSymbols = QAbstractSpinBox__ButtonSymbols(2)
)

// QAbstractSpinBox::CorrectionMode
//
//go:generate stringer -type=QAbstractSpinBox__CorrectionMode
type QAbstractSpinBox__CorrectionMode int64

const (
	QAbstractSpinBox__CorrectToPreviousValue QAbstractSpinBox__CorrectionMode = QAbstractSpinBox__CorrectionMode(0)
	QAbstractSpinBox__CorrectToNearestValue  QAbstractSpinBox__CorrectionMode = QAbstractSpinBox__CorrectionMode(1)
)

// QAbstractSpinBox::StepEnabledFlag
//
//go:generate stringer -type=QAbstractSpinBox__StepEnabledFlag
type QAbstractSpinBox__StepEnabledFlag int64

const (
	QAbstractSpinBox__StepNone        QAbstractSpinBox__StepEnabledFlag = QAbstractSpinBox__StepEnabledFlag(0x00)
	QAbstractSpinBox__StepUpEnabled   QAbstractSpinBox__StepEnabledFlag = QAbstractSpinBox__StepEnabledFlag(0x01)
	QAbstractSpinBox__StepDownEnabled QAbstractSpinBox__StepEnabledFlag = QAbstractSpinBox__StepEnabledFlag(0x02)
)

// QAbstractSpinBox::StepType
//
//go:generate stringer -type=QAbstractSpinBox__StepType
type QAbstractSpinBox__StepType int64

const (
	QAbstractSpinBox__DefaultStepType         QAbstractSpinBox__StepType = QAbstractSpinBox__StepType(0)
	QAbstractSpinBox__AdaptiveDecimalStepType QAbstractSpinBox__StepType = QAbstractSpinBox__StepType(1)
)

type QAction struct {
	core.QObject
}

type QAction_ITF interface {
	core.QObject_ITF
	QAction_PTR() *QAction
}

func (ptr *QAction) QAction_PTR() *QAction {
	return ptr
}

func (ptr *QAction) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAction) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAction(ptr QAction_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAction_PTR().Pointer()
	}
	return nil
}

func NewQActionFromPointer(ptr unsafe.Pointer) (n *QAction) {
	n = new(QAction)
	n.SetPointer(ptr)
	return
}

// QAction::ActionEvent
//
//go:generate stringer -type=QAction__ActionEvent
type QAction__ActionEvent int64

const (
	QAction__Trigger QAction__ActionEvent = QAction__ActionEvent(0)
	QAction__Hover   QAction__ActionEvent = QAction__ActionEvent(1)
)

// QAction::MenuRole
//
//go:generate stringer -type=QAction__MenuRole
type QAction__MenuRole int64

const (
	QAction__NoRole                  QAction__MenuRole = QAction__MenuRole(0)
	QAction__TextHeuristicRole       QAction__MenuRole = QAction__MenuRole(1)
	QAction__ApplicationSpecificRole QAction__MenuRole = QAction__MenuRole(2)
	QAction__AboutQtRole             QAction__MenuRole = QAction__MenuRole(3)
	QAction__AboutRole               QAction__MenuRole = QAction__MenuRole(4)
	QAction__PreferencesRole         QAction__MenuRole = QAction__MenuRole(5)
	QAction__QuitRole                QAction__MenuRole = QAction__MenuRole(6)
)

// QAction::Priority
//
//go:generate stringer -type=QAction__Priority
type QAction__Priority int64

const (
	QAction__LowPriority    QAction__Priority = QAction__Priority(0)
	QAction__NormalPriority QAction__Priority = QAction__Priority(128)
	QAction__HighPriority   QAction__Priority = QAction__Priority(256)
)

func NewQAction(parent core.QObject_ITF) *QAction {
	tmpValue := NewQActionFromPointer(C.QAction_NewQAction(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQAction2(text string, parent core.QObject_ITF) *QAction {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQActionFromPointer(C.QAction_NewQAction2(C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQAction3(icon gui.QIcon_ITF, text string, parent core.QObject_ITF) *QAction {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQActionFromPointer(C.QAction_NewQAction3(gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAction) Activate(event QAction__ActionEvent) {
	if ptr.Pointer() != nil {
		C.QAction_Activate(ptr.Pointer(), C.longlong(event))
	}
}

//export callbackQAction_Changed
func callbackQAction_Changed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changed"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAction) ConnectChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "changed") {
			C.QAction_ConnectChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "changed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "changed"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "changed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "changed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAction) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "changed")
	}
}

func (ptr *QAction) Changed() {
	if ptr.Pointer() != nil {
		C.QAction_Changed(ptr.Pointer())
	}
}

func (ptr *QAction) Data() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QAction_Data(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQAction_Event
func callbackQAction_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQActionFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAction) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAction_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

func (ptr *QAction) Font() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QAction_Font(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QAction) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QAction_Icon(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QAction) IsChecked() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAction_IsChecked(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAction) Menu() *QMenu {
	if ptr.Pointer() != nil {
		tmpValue := NewQMenuFromPointer(C.QAction_Menu(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) Priority() QAction__Priority {
	if ptr.Pointer() != nil {
		return QAction__Priority(C.QAction_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAction) SetCheckable(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetCheckable(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQAction_SetChecked
func callbackQAction_SetChecked(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setChecked"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewQActionFromPointer(ptr).SetCheckedDefault(int8(vbo) != 0)
	}
}

func (ptr *QAction) ConnectSetChecked(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setChecked"); signal != nil {
			f := func(vbo bool) {
				(*(*func(bool))(signal))(vbo)
				f(vbo)
			}
			qt.ConnectSignal(ptr.Pointer(), "setChecked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setChecked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAction) DisconnectSetChecked() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setChecked")
	}
}

func (ptr *QAction) SetChecked(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetChecked(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QAction) SetCheckedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetCheckedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QAction) SetData(userData core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetData(ptr.Pointer(), core.PointerFromQVariant(userData))
	}
}

//export callbackQAction_SetDisabled
func callbackQAction_SetDisabled(ptr unsafe.Pointer, b C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(b) != 0)
	} else {
		NewQActionFromPointer(ptr).SetDisabledDefault(int8(b) != 0)
	}
}

func (ptr *QAction) ConnectSetDisabled(f func(b bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDisabled"); signal != nil {
			f := func(b bool) {
				(*(*func(bool))(signal))(b)
				f(b)
			}
			qt.ConnectSignal(ptr.Pointer(), "setDisabled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDisabled", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAction) DisconnectSetDisabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDisabled")
	}
}

func (ptr *QAction) SetDisabled(b bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(b))))
	}
}

func (ptr *QAction) SetDisabledDefault(b bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(b))))
	}
}

//export callbackQAction_SetEnabled
func callbackQAction_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewQActionFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QAction) ConnectSetEnabled(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEnabled"); signal != nil {
			f := func(vbo bool) {
				(*(*func(bool))(signal))(vbo)
				f(vbo)
			}
			qt.ConnectSignal(ptr.Pointer(), "setEnabled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEnabled", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAction) DisconnectSetEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEnabled")
	}
}

func (ptr *QAction) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QAction) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QAction) SetFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QAction) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QAction) SetPriority(priority QAction__Priority) {
	if ptr.Pointer() != nil {
		C.QAction_SetPriority(ptr.Pointer(), C.longlong(priority))
	}
}

func (ptr *QAction) SetShortcut(shortcut gui.QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(shortcut))
	}
}

func (ptr *QAction) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QAction_SetText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QAction) SetToolTip(tip string) {
	if ptr.Pointer() != nil {
		var tipC *C.char
		if tip != "" {
			tipC = C.CString(tip)
			defer C.free(unsafe.Pointer(tipC))
		}
		C.QAction_SetToolTip(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: tipC, len: C.longlong(len(tip))})
	}
}

func (ptr *QAction) Shortcut() *gui.QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQKeySequenceFromPointer(C.QAction_Shortcut(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QAction) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAction_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAction) ToolTip() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAction_ToolTip(ptr.Pointer()))
	}
	return ""
}

//export callbackQAction_Trigger
func callbackQAction_Trigger(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "trigger"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQActionFromPointer(ptr).TriggerDefault()
	}
}

func (ptr *QAction) ConnectTrigger(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "trigger"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "trigger", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "trigger", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAction) DisconnectTrigger() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "trigger")
	}
}

func (ptr *QAction) Trigger() {
	if ptr.Pointer() != nil {
		C.QAction_Trigger(ptr.Pointer())
	}
}

func (ptr *QAction) TriggerDefault() {
	if ptr.Pointer() != nil {
		C.QAction_TriggerDefault(ptr.Pointer())
	}
}

//export callbackQAction_Triggered
func callbackQAction_Triggered(ptr unsafe.Pointer, checked C.char) {
	if signal := qt.GetSignal(ptr, "triggered"); signal != nil {
		(*(*func(bool))(signal))(int8(checked) != 0)
	}

}

func (ptr *QAction) ConnectTriggered(f func(checked bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "triggered") {
			C.QAction_ConnectTriggered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "triggered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "triggered"); signal != nil {
			f := func(checked bool) {
				(*(*func(bool))(signal))(checked)
				f(checked)
			}
			qt.ConnectSignal(ptr.Pointer(), "triggered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "triggered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAction) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "triggered")
	}
}

func (ptr *QAction) Triggered(checked bool) {
	if ptr.Pointer() != nil {
		C.QAction_Triggered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(checked))))
	}
}

//export callbackQAction_DestroyQAction
func callbackQAction_DestroyQAction(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAction"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQActionFromPointer(ptr).DestroyQActionDefault()
	}
}

func (ptr *QAction) ConnectDestroyQAction(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAction"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAction", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAction", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAction) DisconnectDestroyQAction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAction")
	}
}

func (ptr *QAction) DestroyQAction() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAction_DestroyQAction(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAction) DestroyQActionDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAction_DestroyQActionDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAction) __associatedGraphicsWidgets_atList(i int) *QGraphicsWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQGraphicsWidgetFromPointer(C.QAction___associatedGraphicsWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __associatedGraphicsWidgets_setList(i QGraphicsWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___associatedGraphicsWidgets_setList(ptr.Pointer(), PointerFromQGraphicsWidget(i))
	}
}

func (ptr *QAction) __associatedGraphicsWidgets_newList() unsafe.Pointer {
	return C.QAction___associatedGraphicsWidgets_newList(ptr.Pointer())
}

func (ptr *QAction) __associatedWidgets_atList(i int) *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QAction___associatedWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __associatedWidgets_setList(i QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___associatedWidgets_setList(ptr.Pointer(), PointerFromQWidget(i))
	}
}

func (ptr *QAction) __associatedWidgets_newList() unsafe.Pointer {
	return C.QAction___associatedWidgets_newList(ptr.Pointer())
}

func (ptr *QAction) __setShortcuts_shortcuts_atList(i int) *gui.QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQKeySequenceFromPointer(C.QAction___setShortcuts_shortcuts_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*gui.QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __setShortcuts_shortcuts_setList(i gui.QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___setShortcuts_shortcuts_setList(ptr.Pointer(), gui.PointerFromQKeySequence(i))
	}
}

func (ptr *QAction) __setShortcuts_shortcuts_newList() unsafe.Pointer {
	return C.QAction___setShortcuts_shortcuts_newList(ptr.Pointer())
}

func (ptr *QAction) __shortcuts_atList(i int) *gui.QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQKeySequenceFromPointer(C.QAction___shortcuts_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*gui.QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __shortcuts_setList(i gui.QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___shortcuts_setList(ptr.Pointer(), gui.PointerFromQKeySequence(i))
	}
}

func (ptr *QAction) __shortcuts_newList() unsafe.Pointer {
	return C.QAction___shortcuts_newList(ptr.Pointer())
}

func (ptr *QAction) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAction___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAction) __children_newList() unsafe.Pointer {
	return C.QAction___children_newList(ptr.Pointer())
}

func (ptr *QAction) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAction___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QAction) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QAction___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QAction) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAction___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAction) __findChildren_newList() unsafe.Pointer {
	return C.QAction___findChildren_newList(ptr.Pointer())
}

func (ptr *QAction) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAction___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAction) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAction___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAction) __findChildren_newList3() unsafe.Pointer {
	return C.QAction___findChildren_newList3(ptr.Pointer())
}

//export callbackQAction_ChildEvent
func callbackQAction_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQActionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAction) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAction_ConnectNotify
func callbackQAction_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQActionFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAction) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAction_CustomEvent
func callbackQAction_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQActionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAction) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAction_DeleteLater
func callbackQAction_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQActionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAction) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAction_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQAction_Destroyed
func callbackQAction_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQAction_DisconnectNotify
func callbackQAction_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQActionFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAction) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAction_EventFilter
func callbackQAction_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQActionFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAction) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAction_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAction_ObjectNameChanged
func callbackQAction_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQAction_TimerEvent
func callbackQAction_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQActionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAction) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

// QActionGroup::ExclusionPolicy
//
//go:generate stringer -type=QActionGroup__ExclusionPolicy
type QActionGroup__ExclusionPolicy int64

const (
	QActionGroup__None              QActionGroup__ExclusionPolicy = QActionGroup__ExclusionPolicy(0)
	QActionGroup__Exclusive         QActionGroup__ExclusionPolicy = QActionGroup__ExclusionPolicy(1)
	QActionGroup__ExclusiveOptional QActionGroup__ExclusionPolicy = QActionGroup__ExclusionPolicy(2)
)

type QApplication struct {
	gui.QGuiApplication
}

type QApplication_ITF interface {
	gui.QGuiApplication_ITF
	QApplication_PTR() *QApplication
}

func (ptr *QApplication) QApplication_PTR() *QApplication {
	return ptr
}

func (ptr *QApplication) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGuiApplication_PTR().Pointer()
	}
	return nil
}

func (ptr *QApplication) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGuiApplication_PTR().SetPointer(p)
	}
}

func PointerFromQApplication(ptr QApplication_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QApplication_PTR().Pointer()
	}
	return nil
}

func NewQApplicationFromPointer(ptr unsafe.Pointer) (n *QApplication) {
	n = new(QApplication)
	n.SetPointer(ptr)
	return
}
func NewQApplication(argc int, argv []string) *QApplication {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	tmpValue := NewQApplicationFromPointer(C.QApplication_NewQApplication(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QApplication_ActiveWindow() *QWidget {
	tmpValue := NewQWidgetFromPointer(C.QApplication_QApplication_ActiveWindow())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QApplication) ActiveWindow() *QWidget {
	tmpValue := NewQWidgetFromPointer(C.QApplication_QApplication_ActiveWindow())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQApplication_Event
func callbackQApplication_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQApplicationFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QApplication) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QApplication_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

func QApplication_Exec() int {
	return int(int32(C.QApplication_QApplication_Exec()))
}

func (ptr *QApplication) Exec() int {
	return int(int32(C.QApplication_QApplication_Exec()))
}

func QApplication_Font() *gui.QFont {
	tmpValue := gui.NewQFontFromPointer(C.QApplication_QApplication_Font())
	qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
	return tmpValue
}

func (ptr *QApplication) Font() *gui.QFont {
	tmpValue := gui.NewQFontFromPointer(C.QApplication_QApplication_Font())
	qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
	return tmpValue
}

func QApplication_Font2(widget QWidget_ITF) *gui.QFont {
	tmpValue := gui.NewQFontFromPointer(C.QApplication_QApplication_Font2(PointerFromQWidget(widget)))
	qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
	return tmpValue
}

func (ptr *QApplication) Font2(widget QWidget_ITF) *gui.QFont {
	tmpValue := gui.NewQFontFromPointer(C.QApplication_QApplication_Font2(PointerFromQWidget(widget)))
	qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
	return tmpValue
}

func QApplication_Font3(className string) *gui.QFont {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	tmpValue := gui.NewQFontFromPointer(C.QApplication_QApplication_Font3(classNameC))
	qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
	return tmpValue
}

func (ptr *QApplication) Font3(className string) *gui.QFont {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	tmpValue := gui.NewQFontFromPointer(C.QApplication_QApplication_Font3(classNameC))
	qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
	return tmpValue
}

func QApplication_FontMetrics() *gui.QFontMetrics {
	tmpValue := gui.NewQFontMetricsFromPointer(C.QApplication_QApplication_FontMetrics())
	qt.SetFinalizer(tmpValue, (*gui.QFontMetrics).DestroyQFontMetrics)
	return tmpValue
}

func (ptr *QApplication) FontMetrics() *gui.QFontMetrics {
	tmpValue := gui.NewQFontMetricsFromPointer(C.QApplication_QApplication_FontMetrics())
	qt.SetFinalizer(tmpValue, (*gui.QFontMetrics).DestroyQFontMetrics)
	return tmpValue
}

func QApplication_SetActiveWindow(active QWidget_ITF) {
	C.QApplication_QApplication_SetActiveWindow(PointerFromQWidget(active))
}

func (ptr *QApplication) SetActiveWindow(active QWidget_ITF) {
	C.QApplication_QApplication_SetActiveWindow(PointerFromQWidget(active))
}

func QApplication_SetFont(font gui.QFont_ITF, className string) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	C.QApplication_QApplication_SetFont(gui.PointerFromQFont(font), classNameC)
}

func (ptr *QApplication) SetFont(font gui.QFont_ITF, className string) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	C.QApplication_QApplication_SetFont(gui.PointerFromQFont(font), classNameC)
}

func QApplication_SetStyle(style QStyle_ITF) {
	C.QApplication_QApplication_SetStyle(PointerFromQStyle(style))
}

func (ptr *QApplication) SetStyle(style QStyle_ITF) {
	C.QApplication_QApplication_SetStyle(PointerFromQStyle(style))
}

func QApplication_SetStyle2(style string) *QStyle {
	var styleC *C.char
	if style != "" {
		styleC = C.CString(style)
		defer C.free(unsafe.Pointer(styleC))
	}
	tmpValue := NewQStyleFromPointer(C.QApplication_QApplication_SetStyle2(C.struct_QtWidgets_PackedString{data: styleC, len: C.longlong(len(style))}))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QApplication) SetStyle2(style string) *QStyle {
	var styleC *C.char
	if style != "" {
		styleC = C.CString(style)
		defer C.free(unsafe.Pointer(styleC))
	}
	tmpValue := NewQStyleFromPointer(C.QApplication_QApplication_SetStyle2(C.struct_QtWidgets_PackedString{data: styleC, len: C.longlong(len(style))}))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQApplication_SetStyleSheet
func callbackQApplication_SetStyleSheet(ptr unsafe.Pointer, sheet C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(sheet))
	} else {
		NewQApplicationFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(sheet))
	}
}

func (ptr *QApplication) ConnectSetStyleSheet(f func(sheet string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setStyleSheet"); signal != nil {
			f := func(sheet string) {
				(*(*func(string))(signal))(sheet)
				f(sheet)
			}
			qt.ConnectSignal(ptr.Pointer(), "setStyleSheet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setStyleSheet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QApplication) DisconnectSetStyleSheet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setStyleSheet")
	}
}

func (ptr *QApplication) SetStyleSheet(sheet string) {
	if ptr.Pointer() != nil {
		var sheetC *C.char
		if sheet != "" {
			sheetC = C.CString(sheet)
			defer C.free(unsafe.Pointer(sheetC))
		}
		C.QApplication_SetStyleSheet(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: sheetC, len: C.longlong(len(sheet))})
	}
}

func (ptr *QApplication) SetStyleSheetDefault(sheet string) {
	if ptr.Pointer() != nil {
		var sheetC *C.char
		if sheet != "" {
			sheetC = C.CString(sheet)
			defer C.free(unsafe.Pointer(sheetC))
		}
		C.QApplication_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: sheetC, len: C.longlong(len(sheet))})
	}
}

func QApplication_SetWindowIcon(icon gui.QIcon_ITF) {
	C.QApplication_QApplication_SetWindowIcon(gui.PointerFromQIcon(icon))
}

func (ptr *QApplication) SetWindowIcon(icon gui.QIcon_ITF) {
	C.QApplication_QApplication_SetWindowIcon(gui.PointerFromQIcon(icon))
}

func QApplication_Style() *QStyle {
	tmpValue := NewQStyleFromPointer(C.QApplication_QApplication_Style())
	return tmpValue
}

func (ptr *QApplication) Style() *QStyle {
	tmpValue := NewQStyleFromPointer(C.QApplication_QApplication_Style())
	return tmpValue
}

func (ptr *QApplication) StyleSheet() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QApplication_StyleSheet(ptr.Pointer()))
	}
	return ""
}

func QApplication_WindowIcon() *gui.QIcon {
	tmpValue := gui.NewQIconFromPointer(C.QApplication_QApplication_WindowIcon())
	qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
	return tmpValue
}

func (ptr *QApplication) WindowIcon() *gui.QIcon {
	tmpValue := gui.NewQIconFromPointer(C.QApplication_QApplication_WindowIcon())
	qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
	return tmpValue
}

//export callbackQApplication_DestroyQApplication
func callbackQApplication_DestroyQApplication(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QApplication"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQApplicationFromPointer(ptr).DestroyQApplicationDefault()
	}
}

func (ptr *QApplication) ConnectDestroyQApplication(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QApplication"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QApplication", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QApplication", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QApplication) DisconnectDestroyQApplication() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QApplication")
	}
}

func (ptr *QApplication) DestroyQApplication() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QApplication_DestroyQApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QApplication) DestroyQApplicationDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QApplication_DestroyQApplicationDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QApplication) __allWidgets_atList(i int) *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QApplication___allWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __allWidgets_setList(i QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___allWidgets_setList(ptr.Pointer(), PointerFromQWidget(i))
	}
}

func (ptr *QApplication) __allWidgets_newList() unsafe.Pointer {
	return C.QApplication___allWidgets_newList(ptr.Pointer())
}

func (ptr *QApplication) __topLevelWidgets_atList(i int) *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QApplication___topLevelWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __topLevelWidgets_setList(i QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___topLevelWidgets_setList(ptr.Pointer(), PointerFromQWidget(i))
	}
}

func (ptr *QApplication) __topLevelWidgets_newList() unsafe.Pointer {
	return C.QApplication___topLevelWidgets_newList(ptr.Pointer())
}

func (ptr *QApplication) __screens_atList(i int) *gui.QScreen {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQScreenFromPointer(C.QApplication___screens_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __screens_setList(i gui.QScreen_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___screens_setList(ptr.Pointer(), gui.PointerFromQScreen(i))
	}
}

func (ptr *QApplication) __screens_newList() unsafe.Pointer {
	return C.QApplication___screens_newList(ptr.Pointer())
}

func (ptr *QApplication) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QApplication___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QApplication) __children_newList() unsafe.Pointer {
	return C.QApplication___children_newList(ptr.Pointer())
}

func (ptr *QApplication) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QApplication___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QApplication) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QApplication___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QApplication) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QApplication___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QApplication) __findChildren_newList() unsafe.Pointer {
	return C.QApplication___findChildren_newList(ptr.Pointer())
}

func (ptr *QApplication) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QApplication___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QApplication) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QApplication) __findChildren_newList3() unsafe.Pointer {
	return C.QApplication___findChildren_newList3(ptr.Pointer())
}

//export callbackQApplication_Quit
func callbackQApplication_Quit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "quit"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQApplicationFromPointer(ptr).QuitDefault()
	}
}

func (ptr *QApplication) QuitDefault() {
	if ptr.Pointer() != nil {
		C.QApplication_QuitDefault(ptr.Pointer())
	}
}

//export callbackQApplication_ChildEvent
func callbackQApplication_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQApplicationFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QApplication) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQApplication_ConnectNotify
func callbackQApplication_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQApplicationFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QApplication) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQApplication_CustomEvent
func callbackQApplication_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQApplicationFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QApplication) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQApplication_DeleteLater
func callbackQApplication_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQApplicationFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QApplication) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QApplication_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQApplication_Destroyed
func callbackQApplication_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQApplication_DisconnectNotify
func callbackQApplication_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQApplicationFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QApplication) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQApplication_EventFilter
func callbackQApplication_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQApplicationFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QApplication) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QApplication_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQApplication_ObjectNameChanged
func callbackQApplication_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQApplication_TimerEvent
func callbackQApplication_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQApplicationFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QApplication) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QApplication_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QBoxLayout struct {
	QLayout
}

type QBoxLayout_ITF interface {
	QLayout_ITF
	QBoxLayout_PTR() *QBoxLayout
}

func (ptr *QBoxLayout) QBoxLayout_PTR() *QBoxLayout {
	return ptr
}

func (ptr *QBoxLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QBoxLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLayout_PTR().SetPointer(p)
	}
}

func PointerFromQBoxLayout(ptr QBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQBoxLayoutFromPointer(ptr unsafe.Pointer) (n *QBoxLayout) {
	n = new(QBoxLayout)
	n.SetPointer(ptr)
	return
}

// QBoxLayout::Direction
//
//go:generate stringer -type=QBoxLayout__Direction
type QBoxLayout__Direction int64

const (
	QBoxLayout__LeftToRight QBoxLayout__Direction = QBoxLayout__Direction(0)
	QBoxLayout__RightToLeft QBoxLayout__Direction = QBoxLayout__Direction(1)
	QBoxLayout__TopToBottom QBoxLayout__Direction = QBoxLayout__Direction(2)
	QBoxLayout__BottomToTop QBoxLayout__Direction = QBoxLayout__Direction(3)
	QBoxLayout__Down        QBoxLayout__Direction = QBoxLayout__Direction(QBoxLayout__TopToBottom)
	QBoxLayout__Up          QBoxLayout__Direction = QBoxLayout__Direction(QBoxLayout__BottomToTop)
)

func NewQBoxLayout(dir QBoxLayout__Direction, parent QWidget_ITF) *QBoxLayout {
	tmpValue := NewQBoxLayoutFromPointer(C.QBoxLayout_NewQBoxLayout(C.longlong(dir), PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBoxLayout_AddItem
func callbackQBoxLayout_AddItem(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addItem"); signal != nil {
		(*(*func(*QLayoutItem))(signal))(NewQLayoutItemFromPointer(item))
	} else {
		NewQBoxLayoutFromPointer(ptr).AddItemDefault(NewQLayoutItemFromPointer(item))
	}
}

func (ptr *QBoxLayout) ConnectAddItem(f func(item *QLayoutItem)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addItem"); signal != nil {
			f := func(item *QLayoutItem) {
				(*(*func(*QLayoutItem))(signal))(item)
				f(item)
			}
			qt.ConnectSignal(ptr.Pointer(), "addItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxLayout) DisconnectAddItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addItem")
	}
}

func (ptr *QBoxLayout) AddItem(item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QBoxLayout) AddItemDefault(item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddItemDefault(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QBoxLayout) AddLayout(layout QLayout_ITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddLayout(ptr.Pointer(), PointerFromQLayout(layout), C.int(int32(stretch)))
	}
}

func (ptr *QBoxLayout) AddWidget(widget QWidget_ITF, stretch int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(int32(stretch)), C.longlong(alignment))
	}
}

//export callbackQBoxLayout_Count
func callbackQBoxLayout_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQBoxLayoutFromPointer(ptr).CountDefault()))
}

func (ptr *QBoxLayout) ConnectCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "count"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxLayout) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "count")
	}
}

func (ptr *QBoxLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBoxLayout_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBoxLayout) CountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBoxLayout_CountDefault(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBoxLayout) Direction() QBoxLayout__Direction {
	if ptr.Pointer() != nil {
		return QBoxLayout__Direction(C.QBoxLayout_Direction(ptr.Pointer()))
	}
	return 0
}

//export callbackQBoxLayout_ItemAt
func callbackQBoxLayout_ItemAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return PointerFromQLayoutItem((*(*func(int) *QLayoutItem)(signal))(int(int32(index))))
	}

	return PointerFromQLayoutItem(NewQBoxLayoutFromPointer(ptr).ItemAtDefault(int(int32(index))))
}

func (ptr *QBoxLayout) ConnectItemAt(f func(index int) *QLayoutItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "itemAt"); signal != nil {
			f := func(index int) *QLayoutItem {
				(*(*func(int) *QLayoutItem)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxLayout) DisconnectItemAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "itemAt")
	}
}

func (ptr *QBoxLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_ItemAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QBoxLayout) ItemAtDefault(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_ItemAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QBoxLayout) SetStretch(index int, stretch int) {
	if ptr.Pointer() != nil {
		C.QBoxLayout_SetStretch(ptr.Pointer(), C.int(int32(index)), C.int(int32(stretch)))
	}
}

func (ptr *QBoxLayout) SetStretchFactor(widget QWidget_ITF, stretch int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBoxLayout_SetStretchFactor(ptr.Pointer(), PointerFromQWidget(widget), C.int(int32(stretch)))) != 0
	}
	return false
}

func (ptr *QBoxLayout) SetStretchFactor2(layout QLayout_ITF, stretch int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBoxLayout_SetStretchFactor2(ptr.Pointer(), PointerFromQLayout(layout), C.int(int32(stretch)))) != 0
	}
	return false
}

//export callbackQBoxLayout_SizeHint
func callbackQBoxLayout_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQBoxLayoutFromPointer(ptr).SizeHintDefault())
}

func (ptr *QBoxLayout) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sizeHint"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxLayout) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sizeHint")
	}
}

func (ptr *QBoxLayout) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QBoxLayout_SizeHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QBoxLayout) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QBoxLayout_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QBoxLayout) Stretch(index int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBoxLayout_Stretch(ptr.Pointer(), C.int(int32(index)))))
	}
	return 0
}

//export callbackQBoxLayout_TakeAt
func callbackQBoxLayout_TakeAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "takeAt"); signal != nil {
		return PointerFromQLayoutItem((*(*func(int) *QLayoutItem)(signal))(int(int32(index))))
	}

	return PointerFromQLayoutItem(NewQBoxLayoutFromPointer(ptr).TakeAtDefault(int(int32(index))))
}

func (ptr *QBoxLayout) ConnectTakeAt(f func(index int) *QLayoutItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "takeAt"); signal != nil {
			f := func(index int) *QLayoutItem {
				(*(*func(int) *QLayoutItem)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "takeAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "takeAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxLayout) DisconnectTakeAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "takeAt")
	}
}

func (ptr *QBoxLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_TakeAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QBoxLayout) TakeAtDefault(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QBoxLayout_TakeAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQBoxLayout_DestroyQBoxLayout
func callbackQBoxLayout_DestroyQBoxLayout(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBoxLayout"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBoxLayoutFromPointer(ptr).DestroyQBoxLayoutDefault()
	}
}

func (ptr *QBoxLayout) ConnectDestroyQBoxLayout(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBoxLayout"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBoxLayout", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBoxLayout", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBoxLayout) DisconnectDestroyQBoxLayout() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBoxLayout")
	}
}

func (ptr *QBoxLayout) DestroyQBoxLayout() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBoxLayout_DestroyQBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBoxLayout) DestroyQBoxLayoutDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBoxLayout_DestroyQBoxLayoutDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QButtonGroup struct {
	core.QObject
}

type QButtonGroup_ITF interface {
	core.QObject_ITF
	QButtonGroup_PTR() *QButtonGroup
}

func (ptr *QButtonGroup) QButtonGroup_PTR() *QButtonGroup {
	return ptr
}

func (ptr *QButtonGroup) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QButtonGroup) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQButtonGroup(ptr QButtonGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QButtonGroup_PTR().Pointer()
	}
	return nil
}

func NewQButtonGroupFromPointer(ptr unsafe.Pointer) (n *QButtonGroup) {
	n = new(QButtonGroup)
	n.SetPointer(ptr)
	return
}
func NewQButtonGroup(parent core.QObject_ITF) *QButtonGroup {
	tmpValue := NewQButtonGroupFromPointer(C.QButtonGroup_NewQButtonGroup(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QButtonGroup) Button(id int) *QAbstractButton {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractButtonFromPointer(C.QButtonGroup_Button(ptr.Pointer(), C.int(int32(id))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) Buttons() []*QAbstractButton {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWidgets_PackedList) []*QAbstractButton {
			out := make([]*QAbstractButton, int(l.len))
			tmpList := NewQButtonGroupFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__buttons_atList(i)
			}
			return out
		}(C.QButtonGroup_Buttons(ptr.Pointer()))
	}
	return make([]*QAbstractButton, 0)
}

func (ptr *QButtonGroup) Id(button QAbstractButton_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QButtonGroup_Id(ptr.Pointer(), PointerFromQAbstractButton(button))))
	}
	return 0
}

//export callbackQButtonGroup_DestroyQButtonGroup
func callbackQButtonGroup_DestroyQButtonGroup(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QButtonGroup"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQButtonGroupFromPointer(ptr).DestroyQButtonGroupDefault()
	}
}

func (ptr *QButtonGroup) ConnectDestroyQButtonGroup(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QButtonGroup"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QButtonGroup", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QButtonGroup", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QButtonGroup) DisconnectDestroyQButtonGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QButtonGroup")
	}
}

func (ptr *QButtonGroup) DestroyQButtonGroup() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QButtonGroup_DestroyQButtonGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QButtonGroup) DestroyQButtonGroupDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QButtonGroup_DestroyQButtonGroupDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QButtonGroup) __buttons_atList(i int) *QAbstractButton {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractButtonFromPointer(C.QButtonGroup___buttons_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __buttons_setList(i QAbstractButton_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___buttons_setList(ptr.Pointer(), PointerFromQAbstractButton(i))
	}
}

func (ptr *QButtonGroup) __buttons_newList() unsafe.Pointer {
	return C.QButtonGroup___buttons_newList(ptr.Pointer())
}

func (ptr *QButtonGroup) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QButtonGroup___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QButtonGroup) __children_newList() unsafe.Pointer {
	return C.QButtonGroup___children_newList(ptr.Pointer())
}

func (ptr *QButtonGroup) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QButtonGroup___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QButtonGroup) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QButtonGroup___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QButtonGroup) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QButtonGroup___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QButtonGroup) __findChildren_newList() unsafe.Pointer {
	return C.QButtonGroup___findChildren_newList(ptr.Pointer())
}

func (ptr *QButtonGroup) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QButtonGroup___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QButtonGroup) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QButtonGroup) __findChildren_newList3() unsafe.Pointer {
	return C.QButtonGroup___findChildren_newList3(ptr.Pointer())
}

//export callbackQButtonGroup_ChildEvent
func callbackQButtonGroup_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQButtonGroupFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QButtonGroup) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQButtonGroup_ConnectNotify
func callbackQButtonGroup_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQButtonGroupFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QButtonGroup) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQButtonGroup_CustomEvent
func callbackQButtonGroup_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQButtonGroupFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QButtonGroup) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQButtonGroup_DeleteLater
func callbackQButtonGroup_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQButtonGroupFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QButtonGroup) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QButtonGroup_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQButtonGroup_Destroyed
func callbackQButtonGroup_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQButtonGroup_DisconnectNotify
func callbackQButtonGroup_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQButtonGroupFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QButtonGroup) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQButtonGroup_Event
func callbackQButtonGroup_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQButtonGroupFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QButtonGroup) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QButtonGroup_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQButtonGroup_EventFilter
func callbackQButtonGroup_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQButtonGroupFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QButtonGroup) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QButtonGroup_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQButtonGroup_ObjectNameChanged
func callbackQButtonGroup_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQButtonGroup_TimerEvent
func callbackQButtonGroup_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQButtonGroupFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QButtonGroup) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QButtonGroup_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

// QCalendarWidget::HorizontalHeaderFormat
//
//go:generate stringer -type=QCalendarWidget__HorizontalHeaderFormat
type QCalendarWidget__HorizontalHeaderFormat int64

const (
	QCalendarWidget__NoHorizontalHeader   QCalendarWidget__HorizontalHeaderFormat = QCalendarWidget__HorizontalHeaderFormat(0)
	QCalendarWidget__SingleLetterDayNames QCalendarWidget__HorizontalHeaderFormat = QCalendarWidget__HorizontalHeaderFormat(1)
	QCalendarWidget__ShortDayNames        QCalendarWidget__HorizontalHeaderFormat = QCalendarWidget__HorizontalHeaderFormat(2)
	QCalendarWidget__LongDayNames         QCalendarWidget__HorizontalHeaderFormat = QCalendarWidget__HorizontalHeaderFormat(3)
)

// QCalendarWidget::SelectionMode
//
//go:generate stringer -type=QCalendarWidget__SelectionMode
type QCalendarWidget__SelectionMode int64

const (
	QCalendarWidget__NoSelection     QCalendarWidget__SelectionMode = QCalendarWidget__SelectionMode(0)
	QCalendarWidget__SingleSelection QCalendarWidget__SelectionMode = QCalendarWidget__SelectionMode(1)
)

// QCalendarWidget::VerticalHeaderFormat
//
//go:generate stringer -type=QCalendarWidget__VerticalHeaderFormat
type QCalendarWidget__VerticalHeaderFormat int64

const (
	QCalendarWidget__NoVerticalHeader QCalendarWidget__VerticalHeaderFormat = QCalendarWidget__VerticalHeaderFormat(0)
	QCalendarWidget__ISOWeekNumbers   QCalendarWidget__VerticalHeaderFormat = QCalendarWidget__VerticalHeaderFormat(1)
)

type QColorDialog struct {
	QDialog
}

type QColorDialog_ITF interface {
	QDialog_ITF
	QColorDialog_PTR() *QColorDialog
}

func (ptr *QColorDialog) QColorDialog_PTR() *QColorDialog {
	return ptr
}

func (ptr *QColorDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QColorDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQColorDialog(ptr QColorDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColorDialog_PTR().Pointer()
	}
	return nil
}

func NewQColorDialogFromPointer(ptr unsafe.Pointer) (n *QColorDialog) {
	n = new(QColorDialog)
	n.SetPointer(ptr)
	return
}

// QColorDialog::ColorDialogOption
//
//go:generate stringer -type=QColorDialog__ColorDialogOption
type QColorDialog__ColorDialogOption int64

const (
	QColorDialog__ShowAlphaChannel    QColorDialog__ColorDialogOption = QColorDialog__ColorDialogOption(0x00000001)
	QColorDialog__NoButtons           QColorDialog__ColorDialogOption = QColorDialog__ColorDialogOption(0x00000002)
	QColorDialog__DontUseNativeDialog QColorDialog__ColorDialogOption = QColorDialog__ColorDialogOption(0x00000004)
)

func NewQColorDialog(parent QWidget_ITF) *QColorDialog {
	tmpValue := NewQColorDialogFromPointer(C.QColorDialog_NewQColorDialog(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQColorDialog2(initial gui.QColor_ITF, parent QWidget_ITF) *QColorDialog {
	tmpValue := NewQColorDialogFromPointer(C.QColorDialog_NewQColorDialog2(gui.PointerFromQColor(initial), PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQColorDialog_Done
func callbackQColorDialog_Done(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		(*(*func(int))(signal))(int(int32(result)))
	} else {
		NewQColorDialogFromPointer(ptr).DoneDefault(int(int32(result)))
	}
}

func (ptr *QColorDialog) ConnectDone(f func(result int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "done"); signal != nil {
			f := func(result int) {
				(*(*func(int))(signal))(result)
				f(result)
			}
			qt.ConnectSignal(ptr.Pointer(), "done", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "done", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QColorDialog) DisconnectDone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "done")
	}
}

func (ptr *QColorDialog) Done(result int) {
	if ptr.Pointer() != nil {
		C.QColorDialog_Done(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QColorDialog) DoneDefault(result int) {
	if ptr.Pointer() != nil {
		C.QColorDialog_DoneDefault(ptr.Pointer(), C.int(int32(result)))
	}
}

func QColorDialog_GetColor(initial gui.QColor_ITF, parent QWidget_ITF, title string, options QColorDialog__ColorDialogOption) *gui.QColor {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	tmpValue := gui.NewQColorFromPointer(C.QColorDialog_QColorDialog_GetColor(gui.PointerFromQColor(initial), PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.longlong(options)))
	qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
	return tmpValue
}

func (ptr *QColorDialog) GetColor(initial gui.QColor_ITF, parent QWidget_ITF, title string, options QColorDialog__ColorDialogOption) *gui.QColor {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	tmpValue := gui.NewQColorFromPointer(C.QColorDialog_QColorDialog_GetColor(gui.PointerFromQColor(initial), PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.longlong(options)))
	qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
	return tmpValue
}

func (ptr *QColorDialog) Open(receiver core.QObject_ITF, member string) {
	if ptr.Pointer() != nil {
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		C.QColorDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), memberC)
	}
}

//export callbackQColorDialog_DestroyQColorDialog
func callbackQColorDialog_DestroyQColorDialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QColorDialog"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQColorDialogFromPointer(ptr).DestroyQColorDialogDefault()
	}
}

func (ptr *QColorDialog) ConnectDestroyQColorDialog(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QColorDialog"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QColorDialog", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QColorDialog", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QColorDialog) DisconnectDestroyQColorDialog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QColorDialog")
	}
}

func (ptr *QColorDialog) DestroyQColorDialog() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QColorDialog_DestroyQColorDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QColorDialog) DestroyQColorDialogDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QColorDialog_DestroyQColorDialogDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QColormap::Mode
//
//go:generate stringer -type=QColormap__Mode
type QColormap__Mode int64

const (
	QColormap__Direct  QColormap__Mode = QColormap__Mode(0)
	QColormap__Indexed QColormap__Mode = QColormap__Mode(1)
	QColormap__Gray    QColormap__Mode = QColormap__Mode(2)
)

type QComboBox struct {
	QWidget
}

type QComboBox_ITF interface {
	QWidget_ITF
	QComboBox_PTR() *QComboBox
}

func (ptr *QComboBox) QComboBox_PTR() *QComboBox {
	return ptr
}

func (ptr *QComboBox) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QComboBox) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQComboBox(ptr QComboBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QComboBox_PTR().Pointer()
	}
	return nil
}

func NewQComboBoxFromPointer(ptr unsafe.Pointer) (n *QComboBox) {
	n = new(QComboBox)
	n.SetPointer(ptr)
	return
}

// QComboBox::InsertPolicy
//
//go:generate stringer -type=QComboBox__InsertPolicy
type QComboBox__InsertPolicy int64

const (
	QComboBox__NoInsert             QComboBox__InsertPolicy = QComboBox__InsertPolicy(0)
	QComboBox__InsertAtTop          QComboBox__InsertPolicy = QComboBox__InsertPolicy(1)
	QComboBox__InsertAtCurrent      QComboBox__InsertPolicy = QComboBox__InsertPolicy(2)
	QComboBox__InsertAtBottom       QComboBox__InsertPolicy = QComboBox__InsertPolicy(3)
	QComboBox__InsertAfterCurrent   QComboBox__InsertPolicy = QComboBox__InsertPolicy(4)
	QComboBox__InsertBeforeCurrent  QComboBox__InsertPolicy = QComboBox__InsertPolicy(5)
	QComboBox__InsertAlphabetically QComboBox__InsertPolicy = QComboBox__InsertPolicy(6)
)

// QComboBox::SizeAdjustPolicy
//
//go:generate stringer -type=QComboBox__SizeAdjustPolicy
type QComboBox__SizeAdjustPolicy int64

var (
	QComboBox__AdjustToContents                      QComboBox__SizeAdjustPolicy = QComboBox__SizeAdjustPolicy(0)
	QComboBox__AdjustToContentsOnFirstShow           QComboBox__SizeAdjustPolicy = QComboBox__SizeAdjustPolicy(1)
	QComboBox__AdjustToMinimumContentsLength         QComboBox__SizeAdjustPolicy = QComboBox__SizeAdjustPolicy(2)
	QComboBox__AdjustToMinimumContentsLengthWithIcon QComboBox__SizeAdjustPolicy = QComboBox__SizeAdjustPolicy(C.QComboBox_AdjustToMinimumContentsLengthWithIcon_Type())
)

func NewQComboBox(parent QWidget_ITF) *QComboBox {
	tmpValue := NewQComboBoxFromPointer(C.QComboBox_NewQComboBox(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQComboBox_Activated
func callbackQComboBox_Activated(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QComboBox) ConnectActivated(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activated") {
			C.QComboBox_ConnectActivated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "activated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activated"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "activated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activated", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QComboBox) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QComboBox_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activated")
	}
}

func (ptr *QComboBox) Activated(index int) {
	if ptr.Pointer() != nil {
		C.QComboBox_Activated(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QComboBox) AddItem(text string, userData core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QComboBox_AddItem(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) AddItem2(icon gui.QIcon_ITF, text string, userData core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QComboBox_AddItem2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQVariant(userData))
	}
}

func (ptr *QComboBox) AddItems(texts []string) {
	if ptr.Pointer() != nil {
		textsC := C.CString(strings.Join(texts, "¡¦!"))
		defer C.free(unsafe.Pointer(textsC))
		C.QComboBox_AddItems(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textsC, len: C.longlong(len(strings.Join(texts, "¡¦!")))})
	}
}

//export callbackQComboBox_Clear
func callbackQComboBox_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQComboBoxFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QComboBox) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QComboBox) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *QComboBox) Clear() {
	if ptr.Pointer() != nil {
		C.QComboBox_Clear(ptr.Pointer())
	}
}

func (ptr *QComboBox) ClearDefault() {
	if ptr.Pointer() != nil {
		C.QComboBox_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QComboBox) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QComboBox_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QComboBox) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QComboBox_CurrentIndex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QComboBox) CurrentText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QComboBox_CurrentText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QComboBox) LineEdit() *QLineEdit {
	if ptr.Pointer() != nil {
		tmpValue := NewQLineEditFromPointer(C.QComboBox_LineEdit(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QComboBox) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QComboBox_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QComboBox) PlaceholderText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QComboBox_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

//export callbackQComboBox_SetCurrentIndex
func callbackQComboBox_SetCurrentIndex(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "setCurrentIndex"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	} else {
		NewQComboBoxFromPointer(ptr).SetCurrentIndexDefault(int(int32(index)))
	}
}

func (ptr *QComboBox) ConnectSetCurrentIndex(f func(index int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCurrentIndex"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "setCurrentIndex", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCurrentIndex", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QComboBox) DisconnectSetCurrentIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCurrentIndex")
	}
}

func (ptr *QComboBox) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetCurrentIndex(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QComboBox) SetCurrentIndexDefault(index int) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetCurrentIndexDefault(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQComboBox_SetCurrentText
func callbackQComboBox_SetCurrentText(ptr unsafe.Pointer, text C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setCurrentText"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	} else {
		NewQComboBoxFromPointer(ptr).SetCurrentTextDefault(cGoUnpackString(text))
	}
}

func (ptr *QComboBox) ConnectSetCurrentText(f func(text string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCurrentText"); signal != nil {
			f := func(text string) {
				(*(*func(string))(signal))(text)
				f(text)
			}
			qt.ConnectSignal(ptr.Pointer(), "setCurrentText", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCurrentText", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QComboBox) DisconnectSetCurrentText() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCurrentText")
	}
}

func (ptr *QComboBox) SetCurrentText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QComboBox_SetCurrentText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QComboBox) SetCurrentTextDefault(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QComboBox_SetCurrentTextDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QComboBox) SetEditable(editable bool) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetEditable(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(editable))))
	}
}

func (ptr *QComboBox) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QComboBox) SetPlaceholderText(placeholderText string) {
	if ptr.Pointer() != nil {
		var placeholderTextC *C.char
		if placeholderText != "" {
			placeholderTextC = C.CString(placeholderText)
			defer C.free(unsafe.Pointer(placeholderTextC))
		}
		C.QComboBox_SetPlaceholderText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: placeholderTextC, len: C.longlong(len(placeholderText))})
	}
}

func (ptr *QComboBox) SetValidator(validator gui.QValidator_ITF) {
	if ptr.Pointer() != nil {
		C.QComboBox_SetValidator(ptr.Pointer(), gui.PointerFromQValidator(validator))
	}
}

func (ptr *QComboBox) Validator() *gui.QValidator {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQValidatorFromPointer(C.QComboBox_Validator(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QComboBox) View() *QAbstractItemView {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractItemViewFromPointer(C.QComboBox_View(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQComboBox_DestroyQComboBox
func callbackQComboBox_DestroyQComboBox(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QComboBox"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQComboBoxFromPointer(ptr).DestroyQComboBoxDefault()
	}
}

func (ptr *QComboBox) ConnectDestroyQComboBox(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QComboBox"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QComboBox", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QComboBox", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QComboBox) DisconnectDestroyQComboBox() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QComboBox")
	}
}

func (ptr *QComboBox) DestroyQComboBox() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QComboBox_DestroyQComboBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QComboBox) DestroyQComboBoxDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QComboBox_DestroyQComboBoxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QCompleter::CompletionMode
//
//go:generate stringer -type=QCompleter__CompletionMode
type QCompleter__CompletionMode int64

const (
	QCompleter__PopupCompletion           QCompleter__CompletionMode = QCompleter__CompletionMode(0)
	QCompleter__UnfilteredPopupCompletion QCompleter__CompletionMode = QCompleter__CompletionMode(1)
	QCompleter__InlineCompletion          QCompleter__CompletionMode = QCompleter__CompletionMode(2)
)

// QCompleter::ModelSorting
//
//go:generate stringer -type=QCompleter__ModelSorting
type QCompleter__ModelSorting int64

const (
	QCompleter__UnsortedModel                QCompleter__ModelSorting = QCompleter__ModelSorting(0)
	QCompleter__CaseSensitivelySortedModel   QCompleter__ModelSorting = QCompleter__ModelSorting(1)
	QCompleter__CaseInsensitivelySortedModel QCompleter__ModelSorting = QCompleter__ModelSorting(2)
)

// QDataWidgetMapper::SubmitPolicy
//
//go:generate stringer -type=QDataWidgetMapper__SubmitPolicy
type QDataWidgetMapper__SubmitPolicy int64

const (
	QDataWidgetMapper__AutoSubmit   QDataWidgetMapper__SubmitPolicy = QDataWidgetMapper__SubmitPolicy(0)
	QDataWidgetMapper__ManualSubmit QDataWidgetMapper__SubmitPolicy = QDataWidgetMapper__SubmitPolicy(1)
)

// QDateTimeEdit::Section
//
//go:generate stringer -type=QDateTimeEdit__Section
type QDateTimeEdit__Section int64

const (
	QDateTimeEdit__NoSection     QDateTimeEdit__Section = QDateTimeEdit__Section(0x0000)
	QDateTimeEdit__AmPmSection   QDateTimeEdit__Section = QDateTimeEdit__Section(0x0001)
	QDateTimeEdit__MSecSection   QDateTimeEdit__Section = QDateTimeEdit__Section(0x0002)
	QDateTimeEdit__SecondSection QDateTimeEdit__Section = QDateTimeEdit__Section(0x0004)
	QDateTimeEdit__MinuteSection QDateTimeEdit__Section = QDateTimeEdit__Section(0x0008)
	QDateTimeEdit__HourSection   QDateTimeEdit__Section = QDateTimeEdit__Section(0x0010)
	QDateTimeEdit__DaySection    QDateTimeEdit__Section = QDateTimeEdit__Section(0x0100)
	QDateTimeEdit__MonthSection  QDateTimeEdit__Section = QDateTimeEdit__Section(0x0200)
	QDateTimeEdit__YearSection   QDateTimeEdit__Section = QDateTimeEdit__Section(0x0400)
)

type QDial struct {
	QAbstractSlider
}

type QDial_ITF interface {
	QAbstractSlider_ITF
	QDial_PTR() *QDial
}

func (ptr *QDial) QDial_PTR() *QDial {
	return ptr
}

func (ptr *QDial) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSlider_PTR().Pointer()
	}
	return nil
}

func (ptr *QDial) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSlider_PTR().SetPointer(p)
	}
}

func PointerFromQDial(ptr QDial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDial_PTR().Pointer()
	}
	return nil
}

func NewQDialFromPointer(ptr unsafe.Pointer) (n *QDial) {
	n = new(QDial)
	n.SetPointer(ptr)
	return
}
func NewQDial(parent QWidget_ITF) *QDial {
	tmpValue := NewQDialFromPointer(C.QDial_NewQDial(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDial_DestroyQDial
func callbackQDial_DestroyQDial(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDial"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDialFromPointer(ptr).DestroyQDialDefault()
	}
}

func (ptr *QDial) ConnectDestroyQDial(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDial"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QDial", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDial", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDial) DisconnectDestroyQDial() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDial")
	}
}

func (ptr *QDial) DestroyQDial() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDial_DestroyQDial(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDial) DestroyQDialDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDial_DestroyQDialDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDialog struct {
	QWidget
}

type QDialog_ITF interface {
	QWidget_ITF
	QDialog_PTR() *QDialog
}

func (ptr *QDialog) QDialog_PTR() *QDialog {
	return ptr
}

func (ptr *QDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQDialog(ptr QDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func NewQDialogFromPointer(ptr unsafe.Pointer) (n *QDialog) {
	n = new(QDialog)
	n.SetPointer(ptr)
	return
}

// QDialog::DialogCode
//
//go:generate stringer -type=QDialog__DialogCode
type QDialog__DialogCode int64

const (
	QDialog__Rejected QDialog__DialogCode = QDialog__DialogCode(0)
	QDialog__Accepted QDialog__DialogCode = QDialog__DialogCode(1)
)

func NewQDialog(parent QWidget_ITF, ff core.Qt__WindowType) *QDialog {
	tmpValue := NewQDialogFromPointer(C.QDialog_NewQDialog(PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDialog_Accept
func callbackQDialog_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accept"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QDialog) ConnectAccept(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "accept"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "accept", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "accept", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDialog) DisconnectAccept() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "accept")
	}
}

func (ptr *QDialog) Accept() {
	if ptr.Pointer() != nil {
		C.QDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QDialog) AcceptDefault() {
	if ptr.Pointer() != nil {
		C.QDialog_AcceptDefault(ptr.Pointer())
	}
}

//export callbackQDialog_Done
func callbackQDialog_Done(ptr unsafe.Pointer, r C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		(*(*func(int))(signal))(int(int32(r)))
	} else {
		NewQDialogFromPointer(ptr).DoneDefault(int(int32(r)))
	}
}

func (ptr *QDialog) ConnectDone(f func(r int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "done"); signal != nil {
			f := func(r int) {
				(*(*func(int))(signal))(r)
				f(r)
			}
			qt.ConnectSignal(ptr.Pointer(), "done", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "done", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDialog) DisconnectDone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "done")
	}
}

func (ptr *QDialog) Done(r int) {
	if ptr.Pointer() != nil {
		C.QDialog_Done(ptr.Pointer(), C.int(int32(r)))
	}
}

func (ptr *QDialog) DoneDefault(r int) {
	if ptr.Pointer() != nil {
		C.QDialog_DoneDefault(ptr.Pointer(), C.int(int32(r)))
	}
}

//export callbackQDialog_Exec
func callbackQDialog_Exec(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "exec"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQDialogFromPointer(ptr).ExecDefault()))
}

func (ptr *QDialog) ConnectExec(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exec"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "exec", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exec", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDialog) DisconnectExec() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "exec")
	}
}

func (ptr *QDialog) Exec() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDialog_Exec(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDialog) ExecDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDialog_ExecDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQDialog_Finished
func callbackQDialog_Finished(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func(int))(signal))(int(int32(result)))
	}

}

func (ptr *QDialog) ConnectFinished(f func(result int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QDialog_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func(result int) {
				(*(*func(int))(signal))(result)
				f(result)
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDialog) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QDialog_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QDialog) Finished(result int) {
	if ptr.Pointer() != nil {
		C.QDialog_Finished(ptr.Pointer(), C.int(int32(result)))
	}
}

//export callbackQDialog_Open
func callbackQDialog_Open(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDialogFromPointer(ptr).OpenDefault()
	}
}

func (ptr *QDialog) ConnectOpen(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "open", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDialog) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
	}
}

func (ptr *QDialog) Open() {
	if ptr.Pointer() != nil {
		C.QDialog_Open(ptr.Pointer())
	}
}

func (ptr *QDialog) OpenDefault() {
	if ptr.Pointer() != nil {
		C.QDialog_OpenDefault(ptr.Pointer())
	}
}

func (ptr *QDialog) Result() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDialog_Result(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDialog) SetModal(modal bool) {
	if ptr.Pointer() != nil {
		C.QDialog_SetModal(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(modal))))
	}
}

//export callbackQDialog_DestroyQDialog
func callbackQDialog_DestroyQDialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDialog"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDialogFromPointer(ptr).DestroyQDialogDefault()
	}
}

func (ptr *QDialog) ConnectDestroyQDialog(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDialog"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QDialog", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDialog", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDialog) DisconnectDestroyQDialog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDialog")
	}
}

func (ptr *QDialog) DestroyQDialog() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDialog_DestroyQDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDialog) DestroyQDialogDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDialog_DestroyQDialogDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QDialogButtonBox::ButtonLayout
//
//go:generate stringer -type=QDialogButtonBox__ButtonLayout
type QDialogButtonBox__ButtonLayout int64

var (
	QDialogButtonBox__WinLayout     QDialogButtonBox__ButtonLayout = QDialogButtonBox__ButtonLayout(0)
	QDialogButtonBox__MacLayout     QDialogButtonBox__ButtonLayout = QDialogButtonBox__ButtonLayout(1)
	QDialogButtonBox__KdeLayout     QDialogButtonBox__ButtonLayout = QDialogButtonBox__ButtonLayout(2)
	QDialogButtonBox__GnomeLayout   QDialogButtonBox__ButtonLayout = QDialogButtonBox__ButtonLayout(3)
	QDialogButtonBox__AndroidLayout QDialogButtonBox__ButtonLayout = QDialogButtonBox__ButtonLayout(C.QDialogButtonBox_AndroidLayout_Type())
)

// QDialogButtonBox::ButtonRole
//
//go:generate stringer -type=QDialogButtonBox__ButtonRole
type QDialogButtonBox__ButtonRole int64

const (
	QDialogButtonBox__InvalidRole     QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(-1)
	QDialogButtonBox__AcceptRole      QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(0)
	QDialogButtonBox__RejectRole      QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(1)
	QDialogButtonBox__DestructiveRole QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(2)
	QDialogButtonBox__ActionRole      QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(3)
	QDialogButtonBox__HelpRole        QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(4)
	QDialogButtonBox__YesRole         QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(5)
	QDialogButtonBox__NoRole          QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(6)
	QDialogButtonBox__ResetRole       QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(7)
	QDialogButtonBox__ApplyRole       QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(8)
	QDialogButtonBox__NRoles          QDialogButtonBox__ButtonRole = QDialogButtonBox__ButtonRole(9)
)

// QDialogButtonBox::StandardButton
//
//go:generate stringer -type=QDialogButtonBox__StandardButton
type QDialogButtonBox__StandardButton int64

const (
	QDialogButtonBox__NoButton        QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00000000)
	QDialogButtonBox__Ok              QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00000400)
	QDialogButtonBox__Save            QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00000800)
	QDialogButtonBox__SaveAll         QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00001000)
	QDialogButtonBox__Open            QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00002000)
	QDialogButtonBox__Yes             QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00004000)
	QDialogButtonBox__YesToAll        QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00008000)
	QDialogButtonBox__No              QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00010000)
	QDialogButtonBox__NoToAll         QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00020000)
	QDialogButtonBox__Abort           QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00040000)
	QDialogButtonBox__Retry           QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00080000)
	QDialogButtonBox__Ignore          QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00100000)
	QDialogButtonBox__Close           QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00200000)
	QDialogButtonBox__Cancel          QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00400000)
	QDialogButtonBox__Discard         QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x00800000)
	QDialogButtonBox__Help            QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x01000000)
	QDialogButtonBox__Apply           QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x02000000)
	QDialogButtonBox__Reset           QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x04000000)
	QDialogButtonBox__RestoreDefaults QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(0x08000000)
	QDialogButtonBox__FirstButton     QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(QDialogButtonBox__Ok)
	QDialogButtonBox__LastButton      QDialogButtonBox__StandardButton = QDialogButtonBox__StandardButton(QDialogButtonBox__RestoreDefaults)
)

// QDirModel::Roles
//
//go:generate stringer -type=QDirModel__Roles
type QDirModel__Roles int64

var (
	QDirModel__FileIconRole QDirModel__Roles = QDirModel__Roles(core.Qt__DecorationRole)
	QDirModel__FilePathRole QDirModel__Roles = QDirModel__Roles(C.QDirModel_FilePathRole_Type())
	QDirModel__FileNameRole QDirModel__Roles = QDirModel__Roles(258)
)

type QDockWidget struct {
	QWidget
}

type QDockWidget_ITF interface {
	QWidget_ITF
	QDockWidget_PTR() *QDockWidget
}

func (ptr *QDockWidget) QDockWidget_PTR() *QDockWidget {
	return ptr
}

func (ptr *QDockWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QDockWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQDockWidget(ptr QDockWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDockWidget_PTR().Pointer()
	}
	return nil
}

func NewQDockWidgetFromPointer(ptr unsafe.Pointer) (n *QDockWidget) {
	n = new(QDockWidget)
	n.SetPointer(ptr)
	return
}

// QDockWidget::DockWidgetFeature
//
//go:generate stringer -type=QDockWidget__DockWidgetFeature
type QDockWidget__DockWidgetFeature int64

const (
	QDockWidget__DockWidgetClosable         QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x01)
	QDockWidget__DockWidgetMovable          QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x02)
	QDockWidget__DockWidgetFloatable        QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x04)
	QDockWidget__DockWidgetVerticalTitleBar QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x08)
	QDockWidget__DockWidgetFeatureMask      QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x0f)
	QDockWidget__NoDockWidgetFeatures       QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0x00)
	QDockWidget__Reserved                   QDockWidget__DockWidgetFeature = QDockWidget__DockWidgetFeature(0xff)
)

func NewQDockWidget(title string, parent QWidget_ITF, flags core.Qt__WindowType) *QDockWidget {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	tmpValue := NewQDockWidgetFromPointer(C.QDockWidget_NewQDockWidget(C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQDockWidget2(parent QWidget_ITF, flags core.Qt__WindowType) *QDockWidget {
	tmpValue := NewQDockWidgetFromPointer(C.QDockWidget_NewQDockWidget2(PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDockWidget) SetWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDockWidget_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QDockWidget) Widget() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QDockWidget_Widget(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDockWidget_DestroyQDockWidget
func callbackQDockWidget_DestroyQDockWidget(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDockWidget"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDockWidgetFromPointer(ptr).DestroyQDockWidgetDefault()
	}
}

func (ptr *QDockWidget) ConnectDestroyQDockWidget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDockWidget"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QDockWidget", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDockWidget", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDockWidget) DisconnectDestroyQDockWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDockWidget")
	}
}

func (ptr *QDockWidget) DestroyQDockWidget() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDockWidget_DestroyQDockWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDockWidget) DestroyQDockWidgetDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDockWidget_DestroyQDockWidgetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QFileDialog struct {
	QDialog
}

type QFileDialog_ITF interface {
	QDialog_ITF
	QFileDialog_PTR() *QFileDialog
}

func (ptr *QFileDialog) QFileDialog_PTR() *QFileDialog {
	return ptr
}

func (ptr *QFileDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QFileDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQFileDialog(ptr QFileDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileDialog_PTR().Pointer()
	}
	return nil
}

func NewQFileDialogFromPointer(ptr unsafe.Pointer) (n *QFileDialog) {
	n = new(QFileDialog)
	n.SetPointer(ptr)
	return
}

// QFileDialog::AcceptMode
//
//go:generate stringer -type=QFileDialog__AcceptMode
type QFileDialog__AcceptMode int64

const (
	QFileDialog__AcceptOpen QFileDialog__AcceptMode = QFileDialog__AcceptMode(0)
	QFileDialog__AcceptSave QFileDialog__AcceptMode = QFileDialog__AcceptMode(1)
)

// QFileDialog::DialogLabel
//
//go:generate stringer -type=QFileDialog__DialogLabel
type QFileDialog__DialogLabel int64

const (
	QFileDialog__LookIn   QFileDialog__DialogLabel = QFileDialog__DialogLabel(0)
	QFileDialog__FileName QFileDialog__DialogLabel = QFileDialog__DialogLabel(1)
	QFileDialog__FileType QFileDialog__DialogLabel = QFileDialog__DialogLabel(2)
	QFileDialog__Accept   QFileDialog__DialogLabel = QFileDialog__DialogLabel(3)
	QFileDialog__Reject   QFileDialog__DialogLabel = QFileDialog__DialogLabel(4)
)

// QFileDialog::FileMode
//
//go:generate stringer -type=QFileDialog__FileMode
type QFileDialog__FileMode int64

const (
	QFileDialog__AnyFile       QFileDialog__FileMode = QFileDialog__FileMode(0)
	QFileDialog__ExistingFile  QFileDialog__FileMode = QFileDialog__FileMode(1)
	QFileDialog__Directory     QFileDialog__FileMode = QFileDialog__FileMode(2)
	QFileDialog__ExistingFiles QFileDialog__FileMode = QFileDialog__FileMode(3)
	QFileDialog__DirectoryOnly QFileDialog__FileMode = QFileDialog__FileMode(4)
)

// QFileDialog::Option
//
//go:generate stringer -type=QFileDialog__Option
type QFileDialog__Option int64

const (
	QFileDialog__ShowDirsOnly                QFileDialog__Option = QFileDialog__Option(0x00000001)
	QFileDialog__DontResolveSymlinks         QFileDialog__Option = QFileDialog__Option(0x00000002)
	QFileDialog__DontConfirmOverwrite        QFileDialog__Option = QFileDialog__Option(0x00000004)
	QFileDialog__DontUseSheet                QFileDialog__Option = QFileDialog__Option(0x00000008)
	QFileDialog__DontUseNativeDialog         QFileDialog__Option = QFileDialog__Option(0x00000010)
	QFileDialog__ReadOnly                    QFileDialog__Option = QFileDialog__Option(0x00000020)
	QFileDialog__HideNameFilterDetails       QFileDialog__Option = QFileDialog__Option(0x00000040)
	QFileDialog__DontUseCustomDirectoryIcons QFileDialog__Option = QFileDialog__Option(0x00000080)
)

// QFileDialog::ViewMode
//
//go:generate stringer -type=QFileDialog__ViewMode
type QFileDialog__ViewMode int64

const (
	QFileDialog__Detail QFileDialog__ViewMode = QFileDialog__ViewMode(0)
	QFileDialog__List   QFileDialog__ViewMode = QFileDialog__ViewMode(1)
)

func NewQFileDialog(parent QWidget_ITF, flags core.Qt__WindowType) *QFileDialog {
	tmpValue := NewQFileDialogFromPointer(C.QFileDialog_NewQFileDialog(PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQFileDialog2(parent QWidget_ITF, caption string, directory string, filter string) *QFileDialog {
	var captionC *C.char
	if caption != "" {
		captionC = C.CString(caption)
		defer C.free(unsafe.Pointer(captionC))
	}
	var directoryC *C.char
	if directory != "" {
		directoryC = C.CString(directory)
		defer C.free(unsafe.Pointer(directoryC))
	}
	var filterC *C.char
	if filter != "" {
		filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
	}
	tmpValue := NewQFileDialogFromPointer(C.QFileDialog_NewQFileDialog2(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: captionC, len: C.longlong(len(caption))}, C.struct_QtWidgets_PackedString{data: directoryC, len: C.longlong(len(directory))}, C.struct_QtWidgets_PackedString{data: filterC, len: C.longlong(len(filter))}))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQFileDialog_Accept
func callbackQFileDialog_Accept(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "accept"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQFileDialogFromPointer(ptr).AcceptDefault()
	}
}

func (ptr *QFileDialog) ConnectAccept(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "accept"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "accept", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "accept", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QFileDialog) DisconnectAccept() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "accept")
	}
}

func (ptr *QFileDialog) Accept() {
	if ptr.Pointer() != nil {
		C.QFileDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QFileDialog) AcceptDefault() {
	if ptr.Pointer() != nil {
		C.QFileDialog_AcceptDefault(ptr.Pointer())
	}
}

func (ptr *QFileDialog) Directory() *core.QDir {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDirFromPointer(C.QFileDialog_Directory(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDir).DestroyQDir)
		return tmpValue
	}
	return nil
}

//export callbackQFileDialog_Done
func callbackQFileDialog_Done(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		(*(*func(int))(signal))(int(int32(result)))
	} else {
		NewQFileDialogFromPointer(ptr).DoneDefault(int(int32(result)))
	}
}

func (ptr *QFileDialog) ConnectDone(f func(result int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "done"); signal != nil {
			f := func(result int) {
				(*(*func(int))(signal))(result)
				f(result)
			}
			qt.ConnectSignal(ptr.Pointer(), "done", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "done", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QFileDialog) DisconnectDone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "done")
	}
}

func (ptr *QFileDialog) Done(result int) {
	if ptr.Pointer() != nil {
		C.QFileDialog_Done(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QFileDialog) DoneDefault(result int) {
	if ptr.Pointer() != nil {
		C.QFileDialog_DoneDefault(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QFileDialog) Filter() core.QDir__Filter {
	if ptr.Pointer() != nil {
		return core.QDir__Filter(C.QFileDialog_Filter(ptr.Pointer()))
	}
	return 0
}

func QFileDialog_GetExistingDirectory(parent QWidget_ITF, caption string, dir string, options QFileDialog__Option) string {
	var captionC *C.char
	if caption != "" {
		captionC = C.CString(caption)
		defer C.free(unsafe.Pointer(captionC))
	}
	var dirC *C.char
	if dir != "" {
		dirC = C.CString(dir)
		defer C.free(unsafe.Pointer(dirC))
	}
	return cGoUnpackString(C.QFileDialog_QFileDialog_GetExistingDirectory(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: captionC, len: C.longlong(len(caption))}, C.struct_QtWidgets_PackedString{data: dirC, len: C.longlong(len(dir))}, C.longlong(options)))
}

func (ptr *QFileDialog) GetExistingDirectory(parent QWidget_ITF, caption string, dir string, options QFileDialog__Option) string {
	var captionC *C.char
	if caption != "" {
		captionC = C.CString(caption)
		defer C.free(unsafe.Pointer(captionC))
	}
	var dirC *C.char
	if dir != "" {
		dirC = C.CString(dir)
		defer C.free(unsafe.Pointer(dirC))
	}
	return cGoUnpackString(C.QFileDialog_QFileDialog_GetExistingDirectory(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: captionC, len: C.longlong(len(caption))}, C.struct_QtWidgets_PackedString{data: dirC, len: C.longlong(len(dir))}, C.longlong(options)))
}

func QFileDialog_GetOpenFileName(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	var captionC *C.char
	if caption != "" {
		captionC = C.CString(caption)
		defer C.free(unsafe.Pointer(captionC))
	}
	var dirC *C.char
	if dir != "" {
		dirC = C.CString(dir)
		defer C.free(unsafe.Pointer(dirC))
	}
	var filterC *C.char
	if filter != "" {
		filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
	}
	var selectedFilterC *C.char
	if selectedFilter != "" {
		selectedFilterC = C.CString(selectedFilter)
		defer C.free(unsafe.Pointer(selectedFilterC))
	}
	return cGoUnpackString(C.QFileDialog_QFileDialog_GetOpenFileName(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: captionC, len: C.longlong(len(caption))}, C.struct_QtWidgets_PackedString{data: dirC, len: C.longlong(len(dir))}, C.struct_QtWidgets_PackedString{data: filterC, len: C.longlong(len(filter))}, C.struct_QtWidgets_PackedString{data: selectedFilterC, len: C.longlong(len(selectedFilter))}, C.longlong(options)))
}

func (ptr *QFileDialog) GetOpenFileName(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	var captionC *C.char
	if caption != "" {
		captionC = C.CString(caption)
		defer C.free(unsafe.Pointer(captionC))
	}
	var dirC *C.char
	if dir != "" {
		dirC = C.CString(dir)
		defer C.free(unsafe.Pointer(dirC))
	}
	var filterC *C.char
	if filter != "" {
		filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
	}
	var selectedFilterC *C.char
	if selectedFilter != "" {
		selectedFilterC = C.CString(selectedFilter)
		defer C.free(unsafe.Pointer(selectedFilterC))
	}
	return cGoUnpackString(C.QFileDialog_QFileDialog_GetOpenFileName(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: captionC, len: C.longlong(len(caption))}, C.struct_QtWidgets_PackedString{data: dirC, len: C.longlong(len(dir))}, C.struct_QtWidgets_PackedString{data: filterC, len: C.longlong(len(filter))}, C.struct_QtWidgets_PackedString{data: selectedFilterC, len: C.longlong(len(selectedFilter))}, C.longlong(options)))
}

func QFileDialog_GetSaveFileName(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	var captionC *C.char
	if caption != "" {
		captionC = C.CString(caption)
		defer C.free(unsafe.Pointer(captionC))
	}
	var dirC *C.char
	if dir != "" {
		dirC = C.CString(dir)
		defer C.free(unsafe.Pointer(dirC))
	}
	var filterC *C.char
	if filter != "" {
		filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
	}
	var selectedFilterC *C.char
	if selectedFilter != "" {
		selectedFilterC = C.CString(selectedFilter)
		defer C.free(unsafe.Pointer(selectedFilterC))
	}
	return cGoUnpackString(C.QFileDialog_QFileDialog_GetSaveFileName(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: captionC, len: C.longlong(len(caption))}, C.struct_QtWidgets_PackedString{data: dirC, len: C.longlong(len(dir))}, C.struct_QtWidgets_PackedString{data: filterC, len: C.longlong(len(filter))}, C.struct_QtWidgets_PackedString{data: selectedFilterC, len: C.longlong(len(selectedFilter))}, C.longlong(options)))
}

func (ptr *QFileDialog) GetSaveFileName(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	var captionC *C.char
	if caption != "" {
		captionC = C.CString(caption)
		defer C.free(unsafe.Pointer(captionC))
	}
	var dirC *C.char
	if dir != "" {
		dirC = C.CString(dir)
		defer C.free(unsafe.Pointer(dirC))
	}
	var filterC *C.char
	if filter != "" {
		filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
	}
	var selectedFilterC *C.char
	if selectedFilter != "" {
		selectedFilterC = C.CString(selectedFilter)
		defer C.free(unsafe.Pointer(selectedFilterC))
	}
	return cGoUnpackString(C.QFileDialog_QFileDialog_GetSaveFileName(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: captionC, len: C.longlong(len(caption))}, C.struct_QtWidgets_PackedString{data: dirC, len: C.longlong(len(dir))}, C.struct_QtWidgets_PackedString{data: filterC, len: C.longlong(len(filter))}, C.struct_QtWidgets_PackedString{data: selectedFilterC, len: C.longlong(len(selectedFilter))}, C.longlong(options)))
}

func (ptr *QFileDialog) Open(receiver core.QObject_ITF, member string) {
	if ptr.Pointer() != nil {
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		C.QFileDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), memberC)
	}
}

func (ptr *QFileDialog) SetReadOnly(enabled bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetReadOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQFileDialog_DestroyQFileDialog
func callbackQFileDialog_DestroyQFileDialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QFileDialog"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQFileDialogFromPointer(ptr).DestroyQFileDialogDefault()
	}
}

func (ptr *QFileDialog) ConnectDestroyQFileDialog(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QFileDialog"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QFileDialog", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QFileDialog", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QFileDialog) DisconnectDestroyQFileDialog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QFileDialog")
	}
}

func (ptr *QFileDialog) DestroyQFileDialog() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFileDialog_DestroyQFileDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFileDialog) DestroyQFileDialogDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFileDialog_DestroyQFileDialogDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFileDialog) __getOpenFileUrls_atList(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QFileDialog___getOpenFileUrls_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QFileDialog) __getOpenFileUrls_setList(i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog___getOpenFileUrls_setList(ptr.Pointer(), core.PointerFromQUrl(i))
	}
}

func (ptr *QFileDialog) __getOpenFileUrls_newList() unsafe.Pointer {
	return C.QFileDialog___getOpenFileUrls_newList(ptr.Pointer())
}

func (ptr *QFileDialog) __selectedUrls_atList(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QFileDialog___selectedUrls_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QFileDialog) __selectedUrls_setList(i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog___selectedUrls_setList(ptr.Pointer(), core.PointerFromQUrl(i))
	}
}

func (ptr *QFileDialog) __selectedUrls_newList() unsafe.Pointer {
	return C.QFileDialog___selectedUrls_newList(ptr.Pointer())
}

func (ptr *QFileDialog) __setSidebarUrls_urls_atList(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QFileDialog___setSidebarUrls_urls_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QFileDialog) __setSidebarUrls_urls_setList(i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog___setSidebarUrls_urls_setList(ptr.Pointer(), core.PointerFromQUrl(i))
	}
}

func (ptr *QFileDialog) __setSidebarUrls_urls_newList() unsafe.Pointer {
	return C.QFileDialog___setSidebarUrls_urls_newList(ptr.Pointer())
}

func (ptr *QFileDialog) __sidebarUrls_atList(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QFileDialog___sidebarUrls_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QFileDialog) __sidebarUrls_setList(i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog___sidebarUrls_setList(ptr.Pointer(), core.PointerFromQUrl(i))
	}
}

func (ptr *QFileDialog) __sidebarUrls_newList() unsafe.Pointer {
	return C.QFileDialog___sidebarUrls_newList(ptr.Pointer())
}

func (ptr *QFileDialog) __urlsSelected_urls_atList(i int) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QFileDialog___urlsSelected_urls_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QFileDialog) __urlsSelected_urls_setList(i core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog___urlsSelected_urls_setList(ptr.Pointer(), core.PointerFromQUrl(i))
	}
}

func (ptr *QFileDialog) __urlsSelected_urls_newList() unsafe.Pointer {
	return C.QFileDialog___urlsSelected_urls_newList(ptr.Pointer())
}

// QFileIconProvider::IconType
//
//go:generate stringer -type=QFileIconProvider__IconType
type QFileIconProvider__IconType int64

const (
	QFileIconProvider__Computer QFileIconProvider__IconType = QFileIconProvider__IconType(0)
	QFileIconProvider__Desktop  QFileIconProvider__IconType = QFileIconProvider__IconType(1)
	QFileIconProvider__Trashcan QFileIconProvider__IconType = QFileIconProvider__IconType(2)
	QFileIconProvider__Network  QFileIconProvider__IconType = QFileIconProvider__IconType(3)
	QFileIconProvider__Drive    QFileIconProvider__IconType = QFileIconProvider__IconType(4)
	QFileIconProvider__Folder   QFileIconProvider__IconType = QFileIconProvider__IconType(5)
	QFileIconProvider__File     QFileIconProvider__IconType = QFileIconProvider__IconType(6)
)

// QFileIconProvider::Option
//
//go:generate stringer -type=QFileIconProvider__Option
type QFileIconProvider__Option int64

const (
	QFileIconProvider__DontUseCustomDirectoryIcons QFileIconProvider__Option = QFileIconProvider__Option(0x00000001)
)

// QFileSystemModel::Option
//
//go:generate stringer -type=QFileSystemModel__Option
type QFileSystemModel__Option int64

const (
	QFileSystemModel__DontWatchForChanges         QFileSystemModel__Option = QFileSystemModel__Option(0x00000001)
	QFileSystemModel__DontResolveSymlinks         QFileSystemModel__Option = QFileSystemModel__Option(0x00000002)
	QFileSystemModel__DontUseCustomDirectoryIcons QFileSystemModel__Option = QFileSystemModel__Option(0x00000004)
)

// QFileSystemModel::Roles
//
//go:generate stringer -type=QFileSystemModel__Roles
type QFileSystemModel__Roles int64

var (
	QFileSystemModel__FileIconRole    QFileSystemModel__Roles = QFileSystemModel__Roles(core.Qt__DecorationRole)
	QFileSystemModel__FilePathRole    QFileSystemModel__Roles = QFileSystemModel__Roles(C.QFileSystemModel_FilePathRole_Type())
	QFileSystemModel__FileNameRole    QFileSystemModel__Roles = QFileSystemModel__Roles(C.QFileSystemModel_FileNameRole_Type())
	QFileSystemModel__FilePermissions QFileSystemModel__Roles = QFileSystemModel__Roles(C.QFileSystemModel_FilePermissions_Type())
)

// QFontComboBox::FontFilter
//
//go:generate stringer -type=QFontComboBox__FontFilter
type QFontComboBox__FontFilter int64

const (
	QFontComboBox__AllFonts          QFontComboBox__FontFilter = QFontComboBox__FontFilter(0)
	QFontComboBox__ScalableFonts     QFontComboBox__FontFilter = QFontComboBox__FontFilter(0x1)
	QFontComboBox__NonScalableFonts  QFontComboBox__FontFilter = QFontComboBox__FontFilter(0x2)
	QFontComboBox__MonospacedFonts   QFontComboBox__FontFilter = QFontComboBox__FontFilter(0x4)
	QFontComboBox__ProportionalFonts QFontComboBox__FontFilter = QFontComboBox__FontFilter(0x8)
)

type QFontDialog struct {
	QDialog
}

type QFontDialog_ITF interface {
	QDialog_ITF
	QFontDialog_PTR() *QFontDialog
}

func (ptr *QFontDialog) QFontDialog_PTR() *QFontDialog {
	return ptr
}

func (ptr *QFontDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QFontDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQFontDialog(ptr QFontDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontDialog_PTR().Pointer()
	}
	return nil
}

func NewQFontDialogFromPointer(ptr unsafe.Pointer) (n *QFontDialog) {
	n = new(QFontDialog)
	n.SetPointer(ptr)
	return
}

// QFontDialog::FontDialogOption
//
//go:generate stringer -type=QFontDialog__FontDialogOption
type QFontDialog__FontDialogOption int64

const (
	QFontDialog__NoButtons           QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000001)
	QFontDialog__DontUseNativeDialog QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000002)
	QFontDialog__ScalableFonts       QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000004)
	QFontDialog__NonScalableFonts    QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000008)
	QFontDialog__MonospacedFonts     QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000010)
	QFontDialog__ProportionalFonts   QFontDialog__FontDialogOption = QFontDialog__FontDialogOption(0x00000020)
)

func NewQFontDialog(parent QWidget_ITF) *QFontDialog {
	tmpValue := NewQFontDialogFromPointer(C.QFontDialog_NewQFontDialog(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQFontDialog2(initial gui.QFont_ITF, parent QWidget_ITF) *QFontDialog {
	tmpValue := NewQFontDialogFromPointer(C.QFontDialog_NewQFontDialog2(gui.PointerFromQFont(initial), PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QFontDialog) CurrentFont() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QFontDialog_CurrentFont(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

//export callbackQFontDialog_Done
func callbackQFontDialog_Done(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		(*(*func(int))(signal))(int(int32(result)))
	} else {
		NewQFontDialogFromPointer(ptr).DoneDefault(int(int32(result)))
	}
}

func (ptr *QFontDialog) ConnectDone(f func(result int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "done"); signal != nil {
			f := func(result int) {
				(*(*func(int))(signal))(result)
				f(result)
			}
			qt.ConnectSignal(ptr.Pointer(), "done", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "done", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QFontDialog) DisconnectDone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "done")
	}
}

func (ptr *QFontDialog) Done(result int) {
	if ptr.Pointer() != nil {
		C.QFontDialog_Done(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QFontDialog) DoneDefault(result int) {
	if ptr.Pointer() != nil {
		C.QFontDialog_DoneDefault(ptr.Pointer(), C.int(int32(result)))
	}
}

func QFontDialog_GetFont(ok *bool, initial gui.QFont_ITF, parent QWidget_ITF, title string, options QFontDialog__FontDialogOption) *gui.QFont {
	var okC C.char
	if ok != nil {
		okC = C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
	}
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	tmpValue := gui.NewQFontFromPointer(C.QFontDialog_QFontDialog_GetFont(&okC, gui.PointerFromQFont(initial), PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.longlong(options)))
	qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
	return tmpValue
}

func (ptr *QFontDialog) GetFont(ok *bool, initial gui.QFont_ITF, parent QWidget_ITF, title string, options QFontDialog__FontDialogOption) *gui.QFont {
	var okC C.char
	if ok != nil {
		okC = C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
	}
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	tmpValue := gui.NewQFontFromPointer(C.QFontDialog_QFontDialog_GetFont(&okC, gui.PointerFromQFont(initial), PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.longlong(options)))
	qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
	return tmpValue
}

func QFontDialog_GetFont2(ok *bool, parent QWidget_ITF) *gui.QFont {
	var okC C.char
	if ok != nil {
		okC = C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
	}
	tmpValue := gui.NewQFontFromPointer(C.QFontDialog_QFontDialog_GetFont2(&okC, PointerFromQWidget(parent)))
	qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
	return tmpValue
}

func (ptr *QFontDialog) GetFont2(ok *bool, parent QWidget_ITF) *gui.QFont {
	var okC C.char
	if ok != nil {
		okC = C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
	}
	tmpValue := gui.NewQFontFromPointer(C.QFontDialog_QFontDialog_GetFont2(&okC, PointerFromQWidget(parent)))
	qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
	return tmpValue
}

func (ptr *QFontDialog) Open(receiver core.QObject_ITF, member string) {
	if ptr.Pointer() != nil {
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		C.QFontDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), memberC)
	}
}

// QFormLayout::FieldGrowthPolicy
//
//go:generate stringer -type=QFormLayout__FieldGrowthPolicy
type QFormLayout__FieldGrowthPolicy int64

const (
	QFormLayout__FieldsStayAtSizeHint  QFormLayout__FieldGrowthPolicy = QFormLayout__FieldGrowthPolicy(0)
	QFormLayout__ExpandingFieldsGrow   QFormLayout__FieldGrowthPolicy = QFormLayout__FieldGrowthPolicy(1)
	QFormLayout__AllNonFixedFieldsGrow QFormLayout__FieldGrowthPolicy = QFormLayout__FieldGrowthPolicy(2)
)

// QFormLayout::ItemRole
//
//go:generate stringer -type=QFormLayout__ItemRole
type QFormLayout__ItemRole int64

const (
	QFormLayout__LabelRole    QFormLayout__ItemRole = QFormLayout__ItemRole(0)
	QFormLayout__FieldRole    QFormLayout__ItemRole = QFormLayout__ItemRole(1)
	QFormLayout__SpanningRole QFormLayout__ItemRole = QFormLayout__ItemRole(2)
)

// QFormLayout::RowWrapPolicy
//
//go:generate stringer -type=QFormLayout__RowWrapPolicy
type QFormLayout__RowWrapPolicy int64

const (
	QFormLayout__DontWrapRows QFormLayout__RowWrapPolicy = QFormLayout__RowWrapPolicy(0)
	QFormLayout__WrapLongRows QFormLayout__RowWrapPolicy = QFormLayout__RowWrapPolicy(1)
	QFormLayout__WrapAllRows  QFormLayout__RowWrapPolicy = QFormLayout__RowWrapPolicy(2)
)

type QFrame struct {
	QWidget
}

type QFrame_ITF interface {
	QWidget_ITF
	QFrame_PTR() *QFrame
}

func (ptr *QFrame) QFrame_PTR() *QFrame {
	return ptr
}

func (ptr *QFrame) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QFrame) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQFrame(ptr QFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func NewQFrameFromPointer(ptr unsafe.Pointer) (n *QFrame) {
	n = new(QFrame)
	n.SetPointer(ptr)
	return
}

// QFrame::Shadow
//
//go:generate stringer -type=QFrame__Shadow
type QFrame__Shadow int64

const (
	QFrame__Plain  QFrame__Shadow = QFrame__Shadow(0x0010)
	QFrame__Raised QFrame__Shadow = QFrame__Shadow(0x0020)
	QFrame__Sunken QFrame__Shadow = QFrame__Shadow(0x0030)
)

// QFrame::Shape
//
//go:generate stringer -type=QFrame__Shape
type QFrame__Shape int64

const (
	QFrame__NoFrame     QFrame__Shape = QFrame__Shape(0)
	QFrame__Box         QFrame__Shape = QFrame__Shape(0x0001)
	QFrame__Panel       QFrame__Shape = QFrame__Shape(0x0002)
	QFrame__WinPanel    QFrame__Shape = QFrame__Shape(0x0003)
	QFrame__HLine       QFrame__Shape = QFrame__Shape(0x0004)
	QFrame__VLine       QFrame__Shape = QFrame__Shape(0x0005)
	QFrame__StyledPanel QFrame__Shape = QFrame__Shape(0x0006)
)

// QFrame::StyleMask
//
//go:generate stringer -type=QFrame__StyleMask
type QFrame__StyleMask int64

var (
	QFrame__Shadow_Mask QFrame__StyleMask = QFrame__StyleMask(0x00f0)
	QFrame__Shape_Mask  QFrame__StyleMask = QFrame__StyleMask(0x000f)
)

func NewQFrame(parent QWidget_ITF, ff core.Qt__WindowType) *QFrame {
	tmpValue := NewQFrameFromPointer(C.QFrame_NewQFrame(PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQFrame_DestroyQFrame
func callbackQFrame_DestroyQFrame(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QFrame"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQFrameFromPointer(ptr).DestroyQFrameDefault()
	}
}

func (ptr *QFrame) ConnectDestroyQFrame(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QFrame"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QFrame", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QFrame", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QFrame) DisconnectDestroyQFrame() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QFrame")
	}
}

func (ptr *QFrame) DestroyQFrame() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFrame_DestroyQFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFrame) DestroyQFrameDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFrame_DestroyQFrameDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QGesture::GestureCancelPolicy
//
//go:generate stringer -type=QGesture__GestureCancelPolicy
type QGesture__GestureCancelPolicy int64

const (
	QGesture__CancelNone         QGesture__GestureCancelPolicy = QGesture__GestureCancelPolicy(0)
	QGesture__CancelAllInContext QGesture__GestureCancelPolicy = QGesture__GestureCancelPolicy(1)
)

// QGestureRecognizer::ResultFlag
//
//go:generate stringer -type=QGestureRecognizer__ResultFlag
type QGestureRecognizer__ResultFlag int64

const (
	QGestureRecognizer__Ignore           QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0001)
	QGestureRecognizer__MayBeGesture     QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0002)
	QGestureRecognizer__TriggerGesture   QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0004)
	QGestureRecognizer__FinishGesture    QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0008)
	QGestureRecognizer__CancelGesture    QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0010)
	QGestureRecognizer__ResultState_Mask QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x00ff)
	QGestureRecognizer__ConsumeEventHint QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0x0100)
	QGestureRecognizer__ResultHint_Mask  QGestureRecognizer__ResultFlag = QGestureRecognizer__ResultFlag(0xff00)
)

// QGraphicsBlurEffect::BlurHint
//
//go:generate stringer -type=QGraphicsBlurEffect__BlurHint
type QGraphicsBlurEffect__BlurHint int64

const (
	QGraphicsBlurEffect__PerformanceHint QGraphicsBlurEffect__BlurHint = QGraphicsBlurEffect__BlurHint(0x00)
	QGraphicsBlurEffect__QualityHint     QGraphicsBlurEffect__BlurHint = QGraphicsBlurEffect__BlurHint(0x01)
	QGraphicsBlurEffect__AnimationHint   QGraphicsBlurEffect__BlurHint = QGraphicsBlurEffect__BlurHint(0x02)
)

// QGraphicsEffect::ChangeFlag
//
//go:generate stringer -type=QGraphicsEffect__ChangeFlag
type QGraphicsEffect__ChangeFlag int64

const (
	QGraphicsEffect__SourceAttached            QGraphicsEffect__ChangeFlag = QGraphicsEffect__ChangeFlag(0x1)
	QGraphicsEffect__SourceDetached            QGraphicsEffect__ChangeFlag = QGraphicsEffect__ChangeFlag(0x2)
	QGraphicsEffect__SourceBoundingRectChanged QGraphicsEffect__ChangeFlag = QGraphicsEffect__ChangeFlag(0x4)
	QGraphicsEffect__SourceInvalidated         QGraphicsEffect__ChangeFlag = QGraphicsEffect__ChangeFlag(0x8)
)

// QGraphicsEffect::PixmapPadMode
//
//go:generate stringer -type=QGraphicsEffect__PixmapPadMode
type QGraphicsEffect__PixmapPadMode int64

const (
	QGraphicsEffect__NoPad                      QGraphicsEffect__PixmapPadMode = QGraphicsEffect__PixmapPadMode(0)
	QGraphicsEffect__PadToTransparentBorder     QGraphicsEffect__PixmapPadMode = QGraphicsEffect__PixmapPadMode(1)
	QGraphicsEffect__PadToEffectiveBoundingRect QGraphicsEffect__PixmapPadMode = QGraphicsEffect__PixmapPadMode(2)
)

// QGraphicsEllipseItem::anonymous
//
//go:generate stringer -type=QGraphicsEllipseItem__anonymous
type QGraphicsEllipseItem__anonymous int64

const (
	QGraphicsEllipseItem__Type QGraphicsEllipseItem__anonymous = QGraphicsEllipseItem__anonymous(4)
)

type QGraphicsItem struct {
	ptr unsafe.Pointer
}

type QGraphicsItem_ITF interface {
	QGraphicsItem_PTR() *QGraphicsItem
}

func (ptr *QGraphicsItem) QGraphicsItem_PTR() *QGraphicsItem {
	return ptr
}

func (ptr *QGraphicsItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGraphicsItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGraphicsItem(ptr QGraphicsItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsItemFromPointer(ptr unsafe.Pointer) (n *QGraphicsItem) {
	n = new(QGraphicsItem)
	n.SetPointer(ptr)
	return
}

// QGraphicsItem::CacheMode
//
//go:generate stringer -type=QGraphicsItem__CacheMode
type QGraphicsItem__CacheMode int64

const (
	QGraphicsItem__NoCache               QGraphicsItem__CacheMode = QGraphicsItem__CacheMode(0)
	QGraphicsItem__ItemCoordinateCache   QGraphicsItem__CacheMode = QGraphicsItem__CacheMode(1)
	QGraphicsItem__DeviceCoordinateCache QGraphicsItem__CacheMode = QGraphicsItem__CacheMode(2)
)

// QGraphicsItem::GraphicsItemChange
//
//go:generate stringer -type=QGraphicsItem__GraphicsItemChange
type QGraphicsItem__GraphicsItemChange int64

const (
	QGraphicsItem__ItemPositionChange                 QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(0)
	QGraphicsItem__ItemMatrixChange                   QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(1)
	QGraphicsItem__ItemVisibleChange                  QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(2)
	QGraphicsItem__ItemEnabledChange                  QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(3)
	QGraphicsItem__ItemSelectedChange                 QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(4)
	QGraphicsItem__ItemParentChange                   QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(5)
	QGraphicsItem__ItemChildAddedChange               QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(6)
	QGraphicsItem__ItemChildRemovedChange             QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(7)
	QGraphicsItem__ItemTransformChange                QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(8)
	QGraphicsItem__ItemPositionHasChanged             QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(9)
	QGraphicsItem__ItemTransformHasChanged            QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(10)
	QGraphicsItem__ItemSceneChange                    QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(11)
	QGraphicsItem__ItemVisibleHasChanged              QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(12)
	QGraphicsItem__ItemEnabledHasChanged              QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(13)
	QGraphicsItem__ItemSelectedHasChanged             QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(14)
	QGraphicsItem__ItemParentHasChanged               QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(15)
	QGraphicsItem__ItemSceneHasChanged                QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(16)
	QGraphicsItem__ItemCursorChange                   QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(17)
	QGraphicsItem__ItemCursorHasChanged               QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(18)
	QGraphicsItem__ItemToolTipChange                  QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(19)
	QGraphicsItem__ItemToolTipHasChanged              QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(20)
	QGraphicsItem__ItemFlagsChange                    QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(21)
	QGraphicsItem__ItemFlagsHaveChanged               QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(22)
	QGraphicsItem__ItemZValueChange                   QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(23)
	QGraphicsItem__ItemZValueHasChanged               QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(24)
	QGraphicsItem__ItemOpacityChange                  QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(25)
	QGraphicsItem__ItemOpacityHasChanged              QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(26)
	QGraphicsItem__ItemScenePositionHasChanged        QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(27)
	QGraphicsItem__ItemRotationChange                 QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(28)
	QGraphicsItem__ItemRotationHasChanged             QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(29)
	QGraphicsItem__ItemScaleChange                    QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(30)
	QGraphicsItem__ItemScaleHasChanged                QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(31)
	QGraphicsItem__ItemTransformOriginPointChange     QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(32)
	QGraphicsItem__ItemTransformOriginPointHasChanged QGraphicsItem__GraphicsItemChange = QGraphicsItem__GraphicsItemChange(33)
)

// QGraphicsItem::GraphicsItemFlag
//
//go:generate stringer -type=QGraphicsItem__GraphicsItemFlag
type QGraphicsItem__GraphicsItemFlag int64

const (
	QGraphicsItem__ItemIsMovable                        QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x1)
	QGraphicsItem__ItemIsSelectable                     QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x2)
	QGraphicsItem__ItemIsFocusable                      QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x4)
	QGraphicsItem__ItemClipsToShape                     QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x8)
	QGraphicsItem__ItemClipsChildrenToShape             QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x10)
	QGraphicsItem__ItemIgnoresTransformations           QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x20)
	QGraphicsItem__ItemIgnoresParentOpacity             QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x40)
	QGraphicsItem__ItemDoesntPropagateOpacityToChildren QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x80)
	QGraphicsItem__ItemStacksBehindParent               QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x100)
	QGraphicsItem__ItemUsesExtendedStyleOption          QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x200)
	QGraphicsItem__ItemHasNoContents                    QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x400)
	QGraphicsItem__ItemSendsGeometryChanges             QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x800)
	QGraphicsItem__ItemAcceptsInputMethod               QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x1000)
	QGraphicsItem__ItemNegativeZStacksBehindParent      QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x2000)
	QGraphicsItem__ItemIsPanel                          QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x4000)
	QGraphicsItem__ItemIsFocusScope                     QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x8000)
	QGraphicsItem__ItemSendsScenePositionChanges        QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x10000)
	QGraphicsItem__ItemStopsClickFocusPropagation       QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x20000)
	QGraphicsItem__ItemStopsFocusHandling               QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x40000)
	QGraphicsItem__ItemContainsChildrenInShape          QGraphicsItem__GraphicsItemFlag = QGraphicsItem__GraphicsItemFlag(0x80000)
)

// QGraphicsItem::PanelModality
//
//go:generate stringer -type=QGraphicsItem__PanelModality
type QGraphicsItem__PanelModality int64

const (
	QGraphicsItem__NonModal   QGraphicsItem__PanelModality = QGraphicsItem__PanelModality(0)
	QGraphicsItem__PanelModal QGraphicsItem__PanelModality = QGraphicsItem__PanelModality(1)
	QGraphicsItem__SceneModal QGraphicsItem__PanelModality = QGraphicsItem__PanelModality(2)
)

// QGraphicsItem::anonymous
//
//go:generate stringer -type=QGraphicsItem__anonymous
type QGraphicsItem__anonymous int64

const (
	QGraphicsItem__Type     QGraphicsItem__anonymous = QGraphicsItem__anonymous(1)
	QGraphicsItem__UserType QGraphicsItem__anonymous = QGraphicsItem__anonymous(65536)
)

func NewQGraphicsItem(parent QGraphicsItem_ITF) *QGraphicsItem {
	return NewQGraphicsItemFromPointer(C.QGraphicsItem_NewQGraphicsItem(PointerFromQGraphicsItem(parent)))
}

//export callbackQGraphicsItem_BoundingRect
func callbackQGraphicsItem_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF((*(*func() *core.QRectF)(signal))())
	}

	return core.PointerFromQRectF(core.NewQRectF())
}

func (ptr *QGraphicsItem) ConnectBoundingRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "boundingRect"); signal != nil {
			f := func() *core.QRectF {
				(*(*func() *core.QRectF)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "boundingRect", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "boundingRect", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectBoundingRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "boundingRect")
	}
}

func (ptr *QGraphicsItem) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsItem_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsItem_ContextMenuEvent
func callbackQGraphicsItem_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*QGraphicsSceneContextMenuEvent))(signal))(NewQGraphicsSceneContextMenuEventFromPointer(event))
	} else {
		NewQGraphicsItemFromPointer(ptr).ContextMenuEventDefault(NewQGraphicsSceneContextMenuEventFromPointer(event))
	}
}

func (ptr *QGraphicsItem) ConnectContextMenuEvent(f func(event *QGraphicsSceneContextMenuEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contextMenuEvent"); signal != nil {
			f := func(event *QGraphicsSceneContextMenuEvent) {
				(*(*func(*QGraphicsSceneContextMenuEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "contextMenuEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contextMenuEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectContextMenuEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "contextMenuEvent")
	}
}

func (ptr *QGraphicsItem) ContextMenuEvent(event QGraphicsSceneContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_ContextMenuEvent(ptr.Pointer(), PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

func (ptr *QGraphicsItem) ContextMenuEventDefault(event QGraphicsSceneContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_ContextMenuEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

func (ptr *QGraphicsItem) Cursor() *gui.QCursor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQCursorFromPointer(C.QGraphicsItem_Cursor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QCursor).DestroyQCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItem) Data(key int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QGraphicsItem_Data(ptr.Pointer(), C.int(int32(key))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItem) Group() *QGraphicsItemGroup {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemGroupFromPointer(C.QGraphicsItem_Group(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsItem) Hide() {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_Hide(ptr.Pointer())
	}
}

//export callbackQGraphicsItem_KeyPressEvent
func callbackQGraphicsItem_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keyPressEvent"); signal != nil {
			f := func(event *gui.QKeyEvent) {
				(*(*func(*gui.QKeyEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "keyPressEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keyPressEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "keyPressEvent")
	}
}

func (ptr *QGraphicsItem) KeyPressEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQGraphicsItem_KeyReleaseEvent
func callbackQGraphicsItem_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keyReleaseEvent"); signal != nil {
			f := func(event *gui.QKeyEvent) {
				(*(*func(*gui.QKeyEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "keyReleaseEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keyReleaseEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "keyReleaseEvent")
	}
}

func (ptr *QGraphicsItem) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQGraphicsItem_MouseDoubleClickEvent
func callbackQGraphicsItem_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*QGraphicsSceneMouseEvent))(signal))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsItemFromPointer(ptr).MouseDoubleClickEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsItem) ConnectMouseDoubleClickEvent(f func(event *QGraphicsSceneMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mouseDoubleClickEvent"); signal != nil {
			f := func(event *QGraphicsSceneMouseEvent) {
				(*(*func(*QGraphicsSceneMouseEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "mouseDoubleClickEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mouseDoubleClickEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mouseDoubleClickEvent")
	}
}

func (ptr *QGraphicsItem) MouseDoubleClickEvent(event QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_MouseDoubleClickEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsItem) MouseDoubleClickEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_MouseDoubleClickEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsItem_MousePressEvent
func callbackQGraphicsItem_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*QGraphicsSceneMouseEvent))(signal))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsItemFromPointer(ptr).MousePressEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsItem) ConnectMousePressEvent(f func(event *QGraphicsSceneMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mousePressEvent"); signal != nil {
			f := func(event *QGraphicsSceneMouseEvent) {
				(*(*func(*QGraphicsSceneMouseEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "mousePressEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mousePressEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mousePressEvent")
	}
}

func (ptr *QGraphicsItem) MousePressEvent(event QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_MousePressEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsItem) MousePressEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_MousePressEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsItem_MouseReleaseEvent
func callbackQGraphicsItem_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*QGraphicsSceneMouseEvent))(signal))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsItemFromPointer(ptr).MouseReleaseEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsItem) ConnectMouseReleaseEvent(f func(event *QGraphicsSceneMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mouseReleaseEvent"); signal != nil {
			f := func(event *QGraphicsSceneMouseEvent) {
				(*(*func(*QGraphicsSceneMouseEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "mouseReleaseEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mouseReleaseEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mouseReleaseEvent")
	}
}

func (ptr *QGraphicsItem) MouseReleaseEvent(event QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_MouseReleaseEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsItem) MouseReleaseEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_MouseReleaseEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsItem_Paint
func callbackQGraphicsItem_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		(*(*func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(signal))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QGraphicsItem) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paint"); signal != nil {
			f := func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
				(*(*func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(signal))(painter, option, widget)
				f(painter, option, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "paint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectPaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paint")
	}
}

func (ptr *QGraphicsItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsItem) Pos() *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QGraphicsItem_Pos(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItem) Scroll(dx float64, dy float64, rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_Scroll(ptr.Pointer(), C.double(dx), C.double(dy), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsItem) SetActive(active bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_SetActive(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(active))))
	}
}

func (ptr *QGraphicsItem) SetData(key int, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_SetData(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(value))
	}
}

func (ptr *QGraphicsItem) SetEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QGraphicsItem) SetFocus(focusReason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_SetFocus(ptr.Pointer(), C.longlong(focusReason))
	}
}

func (ptr *QGraphicsItem) SetPos(pos core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_SetPos(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func (ptr *QGraphicsItem) SetPos2(x float64, y float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_SetPos2(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QGraphicsItem) SetToolTip(toolTip string) {
	if ptr.Pointer() != nil {
		var toolTipC *C.char
		if toolTip != "" {
			toolTipC = C.CString(toolTip)
			defer C.free(unsafe.Pointer(toolTipC))
		}
		C.QGraphicsItem_SetToolTip(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: toolTipC, len: C.longlong(len(toolTip))})
	}
}

func (ptr *QGraphicsItem) Show() {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_Show(ptr.Pointer())
	}
}

func (ptr *QGraphicsItem) ToolTip() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGraphicsItem_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsItem) Transform() *gui.QTransform {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQTransformFromPointer(C.QGraphicsItem_Transform(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QTransform).DestroyQTransform)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsItem_Type
func callbackQGraphicsItem_Type(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQGraphicsItemFromPointer(ptr).TypeDefault()))
}

func (ptr *QGraphicsItem) ConnectType(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QGraphicsItem) Type() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGraphicsItem_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsItem) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGraphicsItem_TypeDefault(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsItem) Update(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_Update(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsItem) Update2(x float64, y float64, width float64, height float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem_Update2(ptr.Pointer(), C.double(x), C.double(y), C.double(width), C.double(height))
	}
}

func (ptr *QGraphicsItem) Window() *QGraphicsWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQGraphicsWidgetFromPointer(C.QGraphicsItem_Window(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItem) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsItem_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsItem) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsItem_Y(ptr.Pointer()))
	}
	return 0
}

//export callbackQGraphicsItem_DestroyQGraphicsItem
func callbackQGraphicsItem_DestroyQGraphicsItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsItem"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsItemFromPointer(ptr).DestroyQGraphicsItemDefault()
	}
}

func (ptr *QGraphicsItem) ConnectDestroyQGraphicsItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsItem"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItem) DisconnectDestroyQGraphicsItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsItem")
	}
}

func (ptr *QGraphicsItem) DestroyQGraphicsItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsItem_DestroyQGraphicsItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsItem) DestroyQGraphicsItemDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsItem_DestroyQGraphicsItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsItem) __childItems_atList(i int) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsItem___childItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QGraphicsItem) __childItems_setList(i QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem___childItems_setList(ptr.Pointer(), PointerFromQGraphicsItem(i))
	}
}

func (ptr *QGraphicsItem) __childItems_newList() unsafe.Pointer {
	return C.QGraphicsItem___childItems_newList(ptr.Pointer())
}

func (ptr *QGraphicsItem) __children_atList(i int) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsItem___children_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QGraphicsItem) __children_setList(i QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem___children_setList(ptr.Pointer(), PointerFromQGraphicsItem(i))
	}
}

func (ptr *QGraphicsItem) __children_newList() unsafe.Pointer {
	return C.QGraphicsItem___children_newList(ptr.Pointer())
}

func (ptr *QGraphicsItem) __collidingItems_atList(i int) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsItem___collidingItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QGraphicsItem) __collidingItems_setList(i QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem___collidingItems_setList(ptr.Pointer(), PointerFromQGraphicsItem(i))
	}
}

func (ptr *QGraphicsItem) __collidingItems_newList() unsafe.Pointer {
	return C.QGraphicsItem___collidingItems_newList(ptr.Pointer())
}

func (ptr *QGraphicsItem) __setTransformations_transformations_atList(i int) *QGraphicsTransform {
	if ptr.Pointer() != nil {
		tmpValue := NewQGraphicsTransformFromPointer(C.QGraphicsItem___setTransformations_transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItem) __setTransformations_transformations_setList(i QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem___setTransformations_transformations_setList(ptr.Pointer(), PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QGraphicsItem) __setTransformations_transformations_newList() unsafe.Pointer {
	return C.QGraphicsItem___setTransformations_transformations_newList(ptr.Pointer())
}

func (ptr *QGraphicsItem) __transformations_atList(i int) *QGraphicsTransform {
	if ptr.Pointer() != nil {
		tmpValue := NewQGraphicsTransformFromPointer(C.QGraphicsItem___transformations_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItem) __transformations_setList(i QGraphicsTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItem___transformations_setList(ptr.Pointer(), PointerFromQGraphicsTransform(i))
	}
}

func (ptr *QGraphicsItem) __transformations_newList() unsafe.Pointer {
	return C.QGraphicsItem___transformations_newList(ptr.Pointer())
}

type QGraphicsItemGroup struct {
	QGraphicsItem
}

type QGraphicsItemGroup_ITF interface {
	QGraphicsItem_ITF
	QGraphicsItemGroup_PTR() *QGraphicsItemGroup
}

func (ptr *QGraphicsItemGroup) QGraphicsItemGroup_PTR() *QGraphicsItemGroup {
	return ptr
}

func (ptr *QGraphicsItemGroup) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItem_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsItemGroup) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsItem_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsItemGroup(ptr QGraphicsItemGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItemGroup_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsItemGroupFromPointer(ptr unsafe.Pointer) (n *QGraphicsItemGroup) {
	n = new(QGraphicsItemGroup)
	n.SetPointer(ptr)
	return
}

// QGraphicsItemGroup::anonymous
//
//go:generate stringer -type=QGraphicsItemGroup__anonymous
type QGraphicsItemGroup__anonymous int64

const (
	QGraphicsItemGroup__Type QGraphicsItemGroup__anonymous = QGraphicsItemGroup__anonymous(10)
)

func NewQGraphicsItemGroup(parent QGraphicsItem_ITF) *QGraphicsItemGroup {
	return NewQGraphicsItemGroupFromPointer(C.QGraphicsItemGroup_NewQGraphicsItemGroup(PointerFromQGraphicsItem(parent)))
}

//export callbackQGraphicsItemGroup_BoundingRect
func callbackQGraphicsItemGroup_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF((*(*func() *core.QRectF)(signal))())
	}

	return core.PointerFromQRectF(NewQGraphicsItemGroupFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QGraphicsItemGroup) ConnectBoundingRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "boundingRect"); signal != nil {
			f := func() *core.QRectF {
				(*(*func() *core.QRectF)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "boundingRect", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "boundingRect", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItemGroup) DisconnectBoundingRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "boundingRect")
	}
}

func (ptr *QGraphicsItemGroup) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsItemGroup_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsItemGroup) BoundingRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsItemGroup_BoundingRectDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsItemGroup_DestroyQGraphicsItemGroup
func callbackQGraphicsItemGroup_DestroyQGraphicsItemGroup(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsItemGroup"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsItemGroupFromPointer(ptr).DestroyQGraphicsItemGroupDefault()
	}
}

func (ptr *QGraphicsItemGroup) ConnectDestroyQGraphicsItemGroup(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsItemGroup"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsItemGroup", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsItemGroup", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsItemGroup) DisconnectDestroyQGraphicsItemGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsItemGroup")
	}
}

func (ptr *QGraphicsItemGroup) DestroyQGraphicsItemGroup() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsItemGroup_DestroyQGraphicsItemGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsItemGroup) DestroyQGraphicsItemGroupDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsItemGroup_DestroyQGraphicsItemGroupDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGraphicsItemGroup_Paint
func callbackQGraphicsItemGroup_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		(*(*func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(signal))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsItemGroupFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsItemGroup) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsItemGroup) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

type QGraphicsLayout struct {
	QGraphicsLayoutItem
}

type QGraphicsLayout_ITF interface {
	QGraphicsLayoutItem_ITF
	QGraphicsLayout_PTR() *QGraphicsLayout
}

func (ptr *QGraphicsLayout) QGraphicsLayout_PTR() *QGraphicsLayout {
	return ptr
}

func (ptr *QGraphicsLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLayoutItem_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsLayoutItem_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsLayout(ptr QGraphicsLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLayout_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsLayoutFromPointer(ptr unsafe.Pointer) (n *QGraphicsLayout) {
	n = new(QGraphicsLayout)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsLayout(parent QGraphicsLayoutItem_ITF) *QGraphicsLayout {
	return NewQGraphicsLayoutFromPointer(C.QGraphicsLayout_NewQGraphicsLayout(PointerFromQGraphicsLayoutItem(parent)))
}

func (ptr *QGraphicsLayout) Activate() {
	if ptr.Pointer() != nil {
		C.QGraphicsLayout_Activate(ptr.Pointer())
	}
}

//export callbackQGraphicsLayout_Count
func callbackQGraphicsLayout_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(0))
}

func (ptr *QGraphicsLayout) ConnectCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "count"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayout) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "count")
	}
}

func (ptr *QGraphicsLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGraphicsLayout_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQGraphicsLayout_Invalidate
func callbackQGraphicsLayout_Invalidate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "invalidate"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsLayoutFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QGraphicsLayout) ConnectInvalidate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "invalidate"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "invalidate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "invalidate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayout) DisconnectInvalidate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "invalidate")
	}
}

func (ptr *QGraphicsLayout) Invalidate() {
	if ptr.Pointer() != nil {
		C.QGraphicsLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QGraphicsLayout) InvalidateDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsLayout_InvalidateDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsLayout_ItemAt
func callbackQGraphicsLayout_ItemAt(ptr unsafe.Pointer, i C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return PointerFromQGraphicsLayoutItem((*(*func(int) *QGraphicsLayoutItem)(signal))(int(int32(i))))
	}

	return PointerFromQGraphicsLayoutItem(nil)
}

func (ptr *QGraphicsLayout) ConnectItemAt(f func(i int) *QGraphicsLayoutItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "itemAt"); signal != nil {
			f := func(i int) *QGraphicsLayoutItem {
				(*(*func(int) *QGraphicsLayoutItem)(signal))(i)
				return f(i)
			}
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayout) DisconnectItemAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "itemAt")
	}
}

func (ptr *QGraphicsLayout) ItemAt(i int) *QGraphicsLayoutItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutItemFromPointer(C.QGraphicsLayout_ItemAt(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

//export callbackQGraphicsLayout_RemoveAt
func callbackQGraphicsLayout_RemoveAt(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "removeAt"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QGraphicsLayout) ConnectRemoveAt(f func(index int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeAt"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "removeAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayout) DisconnectRemoveAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeAt")
	}
}

func (ptr *QGraphicsLayout) RemoveAt(index int) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayout_RemoveAt(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQGraphicsLayout_DestroyQGraphicsLayout
func callbackQGraphicsLayout_DestroyQGraphicsLayout(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsLayout"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsLayoutFromPointer(ptr).DestroyQGraphicsLayoutDefault()
	}
}

func (ptr *QGraphicsLayout) ConnectDestroyQGraphicsLayout(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsLayout"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsLayout", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsLayout", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayout) DisconnectDestroyQGraphicsLayout() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsLayout")
	}
}

func (ptr *QGraphicsLayout) DestroyQGraphicsLayout() {
	if ptr.Pointer() != nil {
		C.QGraphicsLayout_DestroyQGraphicsLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsLayout) DestroyQGraphicsLayoutDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsLayout_DestroyQGraphicsLayoutDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGraphicsLayout_SizeHint
func callbackQGraphicsLayout_SizeHint(ptr unsafe.Pointer, which C.longlong, constraint unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSizeF((*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
	}

	return core.PointerFromQSizeF(NewQGraphicsLayoutFromPointer(ptr).SizeHintDefault(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
}

func (ptr *QGraphicsLayout) SizeHint(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsLayout_SizeHint(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsLayout) SizeHintDefault(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsLayout_SizeHintDefault(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

type QGraphicsLayoutItem struct {
	ptr unsafe.Pointer
}

type QGraphicsLayoutItem_ITF interface {
	QGraphicsLayoutItem_PTR() *QGraphicsLayoutItem
}

func (ptr *QGraphicsLayoutItem) QGraphicsLayoutItem_PTR() *QGraphicsLayoutItem {
	return ptr
}

func (ptr *QGraphicsLayoutItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGraphicsLayoutItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGraphicsLayoutItem(ptr QGraphicsLayoutItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsLayoutItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsLayoutItemFromPointer(ptr unsafe.Pointer) (n *QGraphicsLayoutItem) {
	n = new(QGraphicsLayoutItem)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsLayoutItem(parent QGraphicsLayoutItem_ITF, isLayout bool) *QGraphicsLayoutItem {
	return NewQGraphicsLayoutItemFromPointer(C.QGraphicsLayoutItem_NewQGraphicsLayoutItem(PointerFromQGraphicsLayoutItem(parent), C.char(int8(qt.GoBoolToInt(isLayout)))))
}

func (ptr *QGraphicsLayoutItem) Geometry() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsLayoutItem_Geometry(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsLayoutItem) MaximumSize() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsLayoutItem_MaximumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsLayoutItem) MaximumWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsLayoutItem_MaximumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsLayoutItem) MinimumHeight() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsLayoutItem_MinimumHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsLayoutItem) MinimumSize() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsLayoutItem_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsLayoutItem) MinimumWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsLayoutItem_MinimumWidth(ptr.Pointer()))
	}
	return 0
}

//export callbackQGraphicsLayoutItem_SetGeometry
func callbackQGraphicsLayoutItem_SetGeometry(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		(*(*func(*core.QRectF))(signal))(core.NewQRectFFromPointer(rect))
	} else {
		NewQGraphicsLayoutItemFromPointer(ptr).SetGeometryDefault(core.NewQRectFFromPointer(rect))
	}
}

func (ptr *QGraphicsLayoutItem) ConnectSetGeometry(f func(rect *core.QRectF)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGeometry"); signal != nil {
			f := func(rect *core.QRectF) {
				(*(*func(*core.QRectF))(signal))(rect)
				f(rect)
			}
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayoutItem) DisconnectSetGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGeometry")
	}
}

func (ptr *QGraphicsLayoutItem) SetGeometry(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayoutItem_SetGeometry(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsLayoutItem) SetGeometryDefault(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayoutItem_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsLayoutItem) SetMaximumWidth(width float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayoutItem_SetMaximumWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QGraphicsLayoutItem) SetMinimumSize(size core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayoutItem_SetMinimumSize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QGraphicsLayoutItem) SetMinimumSize2(w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayoutItem_SetMinimumSize2(ptr.Pointer(), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsLayoutItem) SetMinimumWidth(width float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayoutItem_SetMinimumWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QGraphicsLayoutItem) SetSizePolicy(policy QSizePolicy_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayoutItem_SetSizePolicy(ptr.Pointer(), PointerFromQSizePolicy(policy))
	}
}

func (ptr *QGraphicsLayoutItem) SetSizePolicy2(hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy, controlType QSizePolicy__ControlType) {
	if ptr.Pointer() != nil {
		C.QGraphicsLayoutItem_SetSizePolicy2(ptr.Pointer(), C.longlong(hPolicy), C.longlong(vPolicy), C.longlong(controlType))
	}
}

//export callbackQGraphicsLayoutItem_SizeHint
func callbackQGraphicsLayoutItem_SizeHint(ptr unsafe.Pointer, which C.longlong, constraint unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSizeF((*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
	}

	return core.PointerFromQSizeF(core.NewQSizeF())
}

func (ptr *QGraphicsLayoutItem) ConnectSizeHint(f func(which core.Qt__SizeHint, constraint *core.QSizeF) *core.QSizeF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sizeHint"); signal != nil {
			f := func(which core.Qt__SizeHint, constraint *core.QSizeF) *core.QSizeF {
				(*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(which, constraint)
				return f(which, constraint)
			}
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayoutItem) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sizeHint")
	}
}

func (ptr *QGraphicsLayoutItem) SizeHint(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsLayoutItem_SizeHint(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsLayoutItem) SizePolicy() *QSizePolicy {
	if ptr.Pointer() != nil {
		tmpValue := NewQSizePolicyFromPointer(C.QGraphicsLayoutItem_SizePolicy(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSizePolicy).DestroyQSizePolicy)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsLayoutItem_DestroyQGraphicsLayoutItem
func callbackQGraphicsLayoutItem_DestroyQGraphicsLayoutItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsLayoutItem"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsLayoutItemFromPointer(ptr).DestroyQGraphicsLayoutItemDefault()
	}
}

func (ptr *QGraphicsLayoutItem) ConnectDestroyQGraphicsLayoutItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsLayoutItem"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsLayoutItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsLayoutItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsLayoutItem) DisconnectDestroyQGraphicsLayoutItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsLayoutItem")
	}
}

func (ptr *QGraphicsLayoutItem) DestroyQGraphicsLayoutItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsLayoutItem_DestroyQGraphicsLayoutItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsLayoutItem) DestroyQGraphicsLayoutItemDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsLayoutItem_DestroyQGraphicsLayoutItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QGraphicsLineItem::anonymous
//
//go:generate stringer -type=QGraphicsLineItem__anonymous
type QGraphicsLineItem__anonymous int64

const (
	QGraphicsLineItem__Type QGraphicsLineItem__anonymous = QGraphicsLineItem__anonymous(6)
)

type QGraphicsObject struct {
	core.QObject
	QGraphicsItem
}

type QGraphicsObject_ITF interface {
	core.QObject_ITF
	QGraphicsItem_ITF
	QGraphicsObject_PTR() *QGraphicsObject
}

func (ptr *QGraphicsObject) QGraphicsObject_PTR() *QGraphicsObject {
	return ptr
}

func (ptr *QGraphicsObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QGraphicsItem_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsObject(ptr QGraphicsObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsObject_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsObjectFromPointer(ptr unsafe.Pointer) (n *QGraphicsObject) {
	n = new(QGraphicsObject)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsObject(parent QGraphicsItem_ITF) *QGraphicsObject {
	tmpValue := NewQGraphicsObjectFromPointer(C.QGraphicsObject_NewQGraphicsObject(PointerFromQGraphicsItem(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGraphicsObject_Event
func callbackQGraphicsObject_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsObjectFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev)))))
}

func (ptr *QGraphicsObject) ConnectEvent(f func(ev *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "event"); signal != nil {
			f := func(ev *core.QEvent) bool {
				(*(*func(*core.QEvent) bool)(signal))(ev)
				return f(ev)
			}
			qt.ConnectSignal(ptr.Pointer(), "event", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "event", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsObject) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "event")
	}
}

func (ptr *QGraphicsObject) Event(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsObject_Event(ptr.Pointer(), core.PointerFromQEvent(ev))) != 0
	}
	return false
}

func (ptr *QGraphicsObject) EventDefault(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsObject_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))) != 0
	}
	return false
}

//export callbackQGraphicsObject_DestroyQGraphicsObject
func callbackQGraphicsObject_DestroyQGraphicsObject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsObject"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsObjectFromPointer(ptr).DestroyQGraphicsObjectDefault()
	}
}

func (ptr *QGraphicsObject) ConnectDestroyQGraphicsObject(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsObject"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsObject", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsObject", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsObject) DisconnectDestroyQGraphicsObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsObject")
	}
}

func (ptr *QGraphicsObject) DestroyQGraphicsObject() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsObject_DestroyQGraphicsObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsObject) DestroyQGraphicsObjectDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsObject_DestroyQGraphicsObjectDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsObject) SetEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QGraphicsObject) Parent() *QGraphicsObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQGraphicsObjectFromPointer(C.QGraphicsObject_Parent(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) Pos() *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QGraphicsObject_Pos(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) SetPos(pos core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_SetPos(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func (ptr *QGraphicsObject) Scale() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsObject_Scale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsObject) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsObject_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsObject) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsObject_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsObject) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsObject___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsObject) __children_newList() unsafe.Pointer {
	return C.QGraphicsObject___children_newList(ptr.Pointer())
}

func (ptr *QGraphicsObject) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGraphicsObject___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGraphicsObject) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGraphicsObject___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGraphicsObject) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsObject___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsObject) __findChildren_newList() unsafe.Pointer {
	return C.QGraphicsObject___findChildren_newList(ptr.Pointer())
}

func (ptr *QGraphicsObject) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsObject___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsObject) __findChildren_newList3() unsafe.Pointer {
	return C.QGraphicsObject___findChildren_newList3(ptr.Pointer())
}

//export callbackQGraphicsObject_ChildEvent
func callbackQGraphicsObject_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsObject) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsObject) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGraphicsObject_ConnectNotify
func callbackQGraphicsObject_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsObjectFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsObject) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGraphicsObject) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsObject_CustomEvent
func callbackQGraphicsObject_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsObject) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsObject) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsObject_DeleteLater
func callbackQGraphicsObject_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGraphicsObject) DeleteLater() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsObject_DeleteLater(ptr.Pointer())
	}
}

func (ptr *QGraphicsObject) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsObject_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsObject_Destroyed
func callbackQGraphicsObject_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGraphicsObject_DisconnectNotify
func callbackQGraphicsObject_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsObjectFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsObject) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGraphicsObject) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsObject_EventFilter
func callbackQGraphicsObject_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsObjectFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsObject) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsObject_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QGraphicsObject) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsObject_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGraphicsObject_ObjectNameChanged
func callbackQGraphicsObject_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGraphicsObject_TimerEvent
func callbackQGraphicsObject_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsObject) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGraphicsObject_BoundingRect
func callbackQGraphicsObject_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF((*(*func() *core.QRectF)(signal))())
	}

	return core.PointerFromQRectF(NewQGraphicsObjectFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QGraphicsObject) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsObject_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsObject) BoundingRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsObject_BoundingRectDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsObject_Paint
func callbackQGraphicsObject_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		(*(*func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(signal))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsObjectFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsObject) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsObject) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsObject_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

// QGraphicsPathItem::anonymous
//
//go:generate stringer -type=QGraphicsPathItem__anonymous
type QGraphicsPathItem__anonymous int64

const (
	QGraphicsPathItem__Type QGraphicsPathItem__anonymous = QGraphicsPathItem__anonymous(2)
)

// QGraphicsPixmapItem::ShapeMode
//
//go:generate stringer -type=QGraphicsPixmapItem__ShapeMode
type QGraphicsPixmapItem__ShapeMode int64

const (
	QGraphicsPixmapItem__MaskShape          QGraphicsPixmapItem__ShapeMode = QGraphicsPixmapItem__ShapeMode(0)
	QGraphicsPixmapItem__BoundingRectShape  QGraphicsPixmapItem__ShapeMode = QGraphicsPixmapItem__ShapeMode(1)
	QGraphicsPixmapItem__HeuristicMaskShape QGraphicsPixmapItem__ShapeMode = QGraphicsPixmapItem__ShapeMode(2)
)

// QGraphicsPixmapItem::anonymous
//
//go:generate stringer -type=QGraphicsPixmapItem__anonymous
type QGraphicsPixmapItem__anonymous int64

const (
	QGraphicsPixmapItem__Type QGraphicsPixmapItem__anonymous = QGraphicsPixmapItem__anonymous(7)
)

// QGraphicsPolygonItem::anonymous
//
//go:generate stringer -type=QGraphicsPolygonItem__anonymous
type QGraphicsPolygonItem__anonymous int64

const (
	QGraphicsPolygonItem__Type QGraphicsPolygonItem__anonymous = QGraphicsPolygonItem__anonymous(5)
)

// QGraphicsProxyWidget::anonymous
//
//go:generate stringer -type=QGraphicsProxyWidget__anonymous
type QGraphicsProxyWidget__anonymous int64

const (
	QGraphicsProxyWidget__Type QGraphicsProxyWidget__anonymous = QGraphicsProxyWidget__anonymous(12)
)

// QGraphicsRectItem::anonymous
//
//go:generate stringer -type=QGraphicsRectItem__anonymous
type QGraphicsRectItem__anonymous int64

const (
	QGraphicsRectItem__Type QGraphicsRectItem__anonymous = QGraphicsRectItem__anonymous(3)
)

// QGraphicsScene::ItemIndexMethod
//
//go:generate stringer -type=QGraphicsScene__ItemIndexMethod
type QGraphicsScene__ItemIndexMethod int64

const (
	QGraphicsScene__BspTreeIndex QGraphicsScene__ItemIndexMethod = QGraphicsScene__ItemIndexMethod(0)
	QGraphicsScene__NoIndex      QGraphicsScene__ItemIndexMethod = QGraphicsScene__ItemIndexMethod(-1)
)

// QGraphicsScene::SceneLayer
//
//go:generate stringer -type=QGraphicsScene__SceneLayer
type QGraphicsScene__SceneLayer int64

const (
	QGraphicsScene__ItemLayer       QGraphicsScene__SceneLayer = QGraphicsScene__SceneLayer(0x1)
	QGraphicsScene__BackgroundLayer QGraphicsScene__SceneLayer = QGraphicsScene__SceneLayer(0x2)
	QGraphicsScene__ForegroundLayer QGraphicsScene__SceneLayer = QGraphicsScene__SceneLayer(0x4)
	QGraphicsScene__AllLayers       QGraphicsScene__SceneLayer = QGraphicsScene__SceneLayer(0xffff)
)

type QGraphicsSceneContextMenuEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneContextMenuEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneContextMenuEvent_PTR() *QGraphicsSceneContextMenuEvent
}

func (ptr *QGraphicsSceneContextMenuEvent) QGraphicsSceneContextMenuEvent_PTR() *QGraphicsSceneContextMenuEvent {
	return ptr
}

func (ptr *QGraphicsSceneContextMenuEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsSceneContextMenuEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsSceneEvent_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsSceneContextMenuEvent(ptr QGraphicsSceneContextMenuEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneContextMenuEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneContextMenuEventFromPointer(ptr unsafe.Pointer) (n *QGraphicsSceneContextMenuEvent) {
	n = new(QGraphicsSceneContextMenuEvent)
	n.SetPointer(ptr)
	return
}

// QGraphicsSceneContextMenuEvent::Reason
//
//go:generate stringer -type=QGraphicsSceneContextMenuEvent__Reason
type QGraphicsSceneContextMenuEvent__Reason int64

const (
	QGraphicsSceneContextMenuEvent__Mouse    QGraphicsSceneContextMenuEvent__Reason = QGraphicsSceneContextMenuEvent__Reason(0)
	QGraphicsSceneContextMenuEvent__Keyboard QGraphicsSceneContextMenuEvent__Reason = QGraphicsSceneContextMenuEvent__Reason(1)
	QGraphicsSceneContextMenuEvent__Other    QGraphicsSceneContextMenuEvent__Reason = QGraphicsSceneContextMenuEvent__Reason(2)
)

func (ptr *QGraphicsSceneContextMenuEvent) Pos() *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QGraphicsSceneContextMenuEvent_Pos(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEvent
func callbackQGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsSceneContextMenuEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsSceneContextMenuEventFromPointer(ptr).DestroyQGraphicsSceneContextMenuEventDefault()
	}
}

func (ptr *QGraphicsSceneContextMenuEvent) ConnectDestroyQGraphicsSceneContextMenuEvent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsSceneContextMenuEvent"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsSceneContextMenuEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsSceneContextMenuEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsSceneContextMenuEvent) DisconnectDestroyQGraphicsSceneContextMenuEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsSceneContextMenuEvent")
	}
}

func (ptr *QGraphicsSceneContextMenuEvent) DestroyQGraphicsSceneContextMenuEvent() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEvent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsSceneContextMenuEvent) DestroyQGraphicsSceneContextMenuEventDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEventDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGraphicsSceneEvent struct {
	core.QEvent
}

type QGraphicsSceneEvent_ITF interface {
	core.QEvent_ITF
	QGraphicsSceneEvent_PTR() *QGraphicsSceneEvent
}

func (ptr *QGraphicsSceneEvent) QGraphicsSceneEvent_PTR() *QGraphicsSceneEvent {
	return ptr
}

func (ptr *QGraphicsSceneEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsSceneEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsSceneEvent(ptr QGraphicsSceneEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneEventFromPointer(ptr unsafe.Pointer) (n *QGraphicsSceneEvent) {
	n = new(QGraphicsSceneEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QGraphicsSceneEvent) Widget() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QGraphicsSceneEvent_Widget(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSceneEvent_DestroyQGraphicsSceneEvent
func callbackQGraphicsSceneEvent_DestroyQGraphicsSceneEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsSceneEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsSceneEventFromPointer(ptr).DestroyQGraphicsSceneEventDefault()
	}
}

func (ptr *QGraphicsSceneEvent) ConnectDestroyQGraphicsSceneEvent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsSceneEvent"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsSceneEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsSceneEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsSceneEvent) DisconnectDestroyQGraphicsSceneEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsSceneEvent")
	}
}

func (ptr *QGraphicsSceneEvent) DestroyQGraphicsSceneEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneEvent_DestroyQGraphicsSceneEvent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsSceneEvent) DestroyQGraphicsSceneEventDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneEvent_DestroyQGraphicsSceneEventDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGraphicsSceneMouseEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneMouseEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneMouseEvent_PTR() *QGraphicsSceneMouseEvent
}

func (ptr *QGraphicsSceneMouseEvent) QGraphicsSceneMouseEvent_PTR() *QGraphicsSceneMouseEvent {
	return ptr
}

func (ptr *QGraphicsSceneMouseEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsSceneMouseEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsSceneEvent_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsSceneMouseEvent(ptr QGraphicsSceneMouseEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneMouseEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneMouseEventFromPointer(ptr unsafe.Pointer) (n *QGraphicsSceneMouseEvent) {
	n = new(QGraphicsSceneMouseEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QGraphicsSceneMouseEvent) Button() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneMouseEvent_Button(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Buttons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneMouseEvent_Buttons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Pos() *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QGraphicsSceneMouseEvent_Pos(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent
func callbackQGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsSceneMouseEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsSceneMouseEventFromPointer(ptr).DestroyQGraphicsSceneMouseEventDefault()
	}
}

func (ptr *QGraphicsSceneMouseEvent) ConnectDestroyQGraphicsSceneMouseEvent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsSceneMouseEvent"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsSceneMouseEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsSceneMouseEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsSceneMouseEvent) DisconnectDestroyQGraphicsSceneMouseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsSceneMouseEvent")
	}
}

func (ptr *QGraphicsSceneMouseEvent) DestroyQGraphicsSceneMouseEvent() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsSceneMouseEvent) DestroyQGraphicsSceneMouseEventDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEventDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGraphicsSceneResizeEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneResizeEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneResizeEvent_PTR() *QGraphicsSceneResizeEvent
}

func (ptr *QGraphicsSceneResizeEvent) QGraphicsSceneResizeEvent_PTR() *QGraphicsSceneResizeEvent {
	return ptr
}

func (ptr *QGraphicsSceneResizeEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsSceneResizeEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsSceneEvent_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsSceneResizeEvent(ptr QGraphicsSceneResizeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneResizeEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneResizeEventFromPointer(ptr unsafe.Pointer) (n *QGraphicsSceneResizeEvent) {
	n = new(QGraphicsSceneResizeEvent)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsSceneResizeEvent2() *QGraphicsSceneResizeEvent {
	tmpValue := NewQGraphicsSceneResizeEventFromPointer(C.QGraphicsSceneResizeEvent_NewQGraphicsSceneResizeEvent2())
	qt.SetFinalizer(tmpValue, (*QGraphicsSceneResizeEvent).DestroyQGraphicsSceneResizeEvent)
	return tmpValue
}

//export callbackQGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEvent
func callbackQGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsSceneResizeEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsSceneResizeEventFromPointer(ptr).DestroyQGraphicsSceneResizeEventDefault()
	}
}

func (ptr *QGraphicsSceneResizeEvent) ConnectDestroyQGraphicsSceneResizeEvent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsSceneResizeEvent"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsSceneResizeEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsSceneResizeEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsSceneResizeEvent) DisconnectDestroyQGraphicsSceneResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsSceneResizeEvent")
	}
}

func (ptr *QGraphicsSceneResizeEvent) DestroyQGraphicsSceneResizeEvent() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEvent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsSceneResizeEvent) DestroyQGraphicsSceneResizeEventDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEventDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QGraphicsSimpleTextItem::anonymous
//
//go:generate stringer -type=QGraphicsSimpleTextItem__anonymous
type QGraphicsSimpleTextItem__anonymous int64

const (
	QGraphicsSimpleTextItem__Type QGraphicsSimpleTextItem__anonymous = QGraphicsSimpleTextItem__anonymous(9)
)

// QGraphicsTextItem::anonymous
//
//go:generate stringer -type=QGraphicsTextItem__anonymous
type QGraphicsTextItem__anonymous int64

const (
	QGraphicsTextItem__Type QGraphicsTextItem__anonymous = QGraphicsTextItem__anonymous(8)
)

type QGraphicsTransform struct {
	core.QObject
}

type QGraphicsTransform_ITF interface {
	core.QObject_ITF
	QGraphicsTransform_PTR() *QGraphicsTransform
}

func (ptr *QGraphicsTransform) QGraphicsTransform_PTR() *QGraphicsTransform {
	return ptr
}

func (ptr *QGraphicsTransform) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsTransform) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsTransform(ptr QGraphicsTransform_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsTransform_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsTransformFromPointer(ptr unsafe.Pointer) (n *QGraphicsTransform) {
	n = new(QGraphicsTransform)
	n.SetPointer(ptr)
	return
}
func NewQGraphicsTransform(parent core.QObject_ITF) *QGraphicsTransform {
	tmpValue := NewQGraphicsTransformFromPointer(C.QGraphicsTransform_NewQGraphicsTransform(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGraphicsTransform_ApplyTo
func callbackQGraphicsTransform_ApplyTo(ptr unsafe.Pointer, matrix unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "applyTo"); signal != nil {
		(*(*func(*gui.QMatrix4x4))(signal))(gui.NewQMatrix4x4FromPointer(matrix))
	}

}

func (ptr *QGraphicsTransform) ConnectApplyTo(f func(matrix *gui.QMatrix4x4)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "applyTo"); signal != nil {
			f := func(matrix *gui.QMatrix4x4) {
				(*(*func(*gui.QMatrix4x4))(signal))(matrix)
				f(matrix)
			}
			qt.ConnectSignal(ptr.Pointer(), "applyTo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "applyTo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsTransform) DisconnectApplyTo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "applyTo")
	}
}

func (ptr *QGraphicsTransform) ApplyTo(matrix gui.QMatrix4x4_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform_ApplyTo(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

//export callbackQGraphicsTransform_Update
func callbackQGraphicsTransform_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsTransformFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QGraphicsTransform) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "update"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "update", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "update", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsTransform) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "update")
	}
}

func (ptr *QGraphicsTransform) Update() {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform_Update(ptr.Pointer())
	}
}

func (ptr *QGraphicsTransform) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsTransform_DestroyQGraphicsTransform
func callbackQGraphicsTransform_DestroyQGraphicsTransform(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsTransform"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsTransformFromPointer(ptr).DestroyQGraphicsTransformDefault()
	}
}

func (ptr *QGraphicsTransform) ConnectDestroyQGraphicsTransform(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsTransform"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsTransform", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsTransform", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsTransform) DisconnectDestroyQGraphicsTransform() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsTransform")
	}
}

func (ptr *QGraphicsTransform) DestroyQGraphicsTransform() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsTransform_DestroyQGraphicsTransform(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsTransform) DestroyQGraphicsTransformDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsTransform_DestroyQGraphicsTransformDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsTransform) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsTransform___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsTransform) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsTransform) __children_newList() unsafe.Pointer {
	return C.QGraphicsTransform___children_newList(ptr.Pointer())
}

func (ptr *QGraphicsTransform) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGraphicsTransform___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsTransform) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGraphicsTransform) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGraphicsTransform___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGraphicsTransform) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsTransform___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsTransform) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsTransform) __findChildren_newList() unsafe.Pointer {
	return C.QGraphicsTransform___findChildren_newList(ptr.Pointer())
}

func (ptr *QGraphicsTransform) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGraphicsTransform___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsTransform) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGraphicsTransform) __findChildren_newList3() unsafe.Pointer {
	return C.QGraphicsTransform___findChildren_newList3(ptr.Pointer())
}

//export callbackQGraphicsTransform_ChildEvent
func callbackQGraphicsTransform_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsTransformFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsTransform) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGraphicsTransform_ConnectNotify
func callbackQGraphicsTransform_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsTransformFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsTransform) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsTransform_CustomEvent
func callbackQGraphicsTransform_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsTransformFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsTransform) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsTransform_DeleteLater
func callbackQGraphicsTransform_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsTransformFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGraphicsTransform) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsTransform_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsTransform_Destroyed
func callbackQGraphicsTransform_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGraphicsTransform_DisconnectNotify
func callbackQGraphicsTransform_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsTransformFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsTransform) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsTransform_Event
func callbackQGraphicsTransform_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsTransformFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGraphicsTransform) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsTransform_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQGraphicsTransform_EventFilter
func callbackQGraphicsTransform_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsTransformFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsTransform) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsTransform_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGraphicsTransform_ObjectNameChanged
func callbackQGraphicsTransform_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGraphicsTransform_TimerEvent
func callbackQGraphicsTransform_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsTransformFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsTransform) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTransform_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

// QGraphicsView::CacheModeFlag
//
//go:generate stringer -type=QGraphicsView__CacheModeFlag
type QGraphicsView__CacheModeFlag int64

const (
	QGraphicsView__CacheNone       QGraphicsView__CacheModeFlag = QGraphicsView__CacheModeFlag(0x0)
	QGraphicsView__CacheBackground QGraphicsView__CacheModeFlag = QGraphicsView__CacheModeFlag(0x1)
)

// QGraphicsView::DragMode
//
//go:generate stringer -type=QGraphicsView__DragMode
type QGraphicsView__DragMode int64

const (
	QGraphicsView__NoDrag         QGraphicsView__DragMode = QGraphicsView__DragMode(0)
	QGraphicsView__ScrollHandDrag QGraphicsView__DragMode = QGraphicsView__DragMode(1)
	QGraphicsView__RubberBandDrag QGraphicsView__DragMode = QGraphicsView__DragMode(2)
)

// QGraphicsView::OptimizationFlag
//
//go:generate stringer -type=QGraphicsView__OptimizationFlag
type QGraphicsView__OptimizationFlag int64

const (
	QGraphicsView__DontClipPainter           QGraphicsView__OptimizationFlag = QGraphicsView__OptimizationFlag(0x1)
	QGraphicsView__DontSavePainterState      QGraphicsView__OptimizationFlag = QGraphicsView__OptimizationFlag(0x2)
	QGraphicsView__DontAdjustForAntialiasing QGraphicsView__OptimizationFlag = QGraphicsView__OptimizationFlag(0x4)
	QGraphicsView__IndirectPainting          QGraphicsView__OptimizationFlag = QGraphicsView__OptimizationFlag(0x8)
)

// QGraphicsView::ViewportAnchor
//
//go:generate stringer -type=QGraphicsView__ViewportAnchor
type QGraphicsView__ViewportAnchor int64

const (
	QGraphicsView__NoAnchor         QGraphicsView__ViewportAnchor = QGraphicsView__ViewportAnchor(0)
	QGraphicsView__AnchorViewCenter QGraphicsView__ViewportAnchor = QGraphicsView__ViewportAnchor(1)
	QGraphicsView__AnchorUnderMouse QGraphicsView__ViewportAnchor = QGraphicsView__ViewportAnchor(2)
)

// QGraphicsView::ViewportUpdateMode
//
//go:generate stringer -type=QGraphicsView__ViewportUpdateMode
type QGraphicsView__ViewportUpdateMode int64

const (
	QGraphicsView__FullViewportUpdate         QGraphicsView__ViewportUpdateMode = QGraphicsView__ViewportUpdateMode(0)
	QGraphicsView__MinimalViewportUpdate      QGraphicsView__ViewportUpdateMode = QGraphicsView__ViewportUpdateMode(1)
	QGraphicsView__SmartViewportUpdate        QGraphicsView__ViewportUpdateMode = QGraphicsView__ViewportUpdateMode(2)
	QGraphicsView__NoViewportUpdate           QGraphicsView__ViewportUpdateMode = QGraphicsView__ViewportUpdateMode(3)
	QGraphicsView__BoundingRectViewportUpdate QGraphicsView__ViewportUpdateMode = QGraphicsView__ViewportUpdateMode(4)
)

type QGraphicsWidget struct {
	QGraphicsObject
	QGraphicsLayoutItem
}

type QGraphicsWidget_ITF interface {
	QGraphicsObject_ITF
	QGraphicsLayoutItem_ITF
	QGraphicsWidget_PTR() *QGraphicsWidget
}

func (ptr *QGraphicsWidget) QGraphicsWidget_PTR() *QGraphicsWidget {
	return ptr
}

func (ptr *QGraphicsWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QGraphicsWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGraphicsObject_PTR().SetPointer(p)
		ptr.QGraphicsLayoutItem_PTR().SetPointer(p)
	}
}

func PointerFromQGraphicsWidget(ptr QGraphicsWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsWidget_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsWidgetFromPointer(ptr unsafe.Pointer) (n *QGraphicsWidget) {
	n = new(QGraphicsWidget)
	n.SetPointer(ptr)
	return
}

// QGraphicsWidget::anonymous
//
//go:generate stringer -type=QGraphicsWidget__anonymous
type QGraphicsWidget__anonymous int64

const (
	QGraphicsWidget__Type QGraphicsWidget__anonymous = QGraphicsWidget__anonymous(11)
)

func NewQGraphicsWidget(parent QGraphicsItem_ITF, wFlags core.Qt__WindowType) *QGraphicsWidget {
	tmpValue := NewQGraphicsWidgetFromPointer(C.QGraphicsWidget_NewQGraphicsWidget(PointerFromQGraphicsItem(parent), C.longlong(wFlags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGraphicsWidget) Actions() []*QAction {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWidgets_PackedList) []*QAction {
			out := make([]*QAction, int(l.len))
			tmpList := NewQGraphicsWidgetFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__actions_atList(i)
			}
			return out
		}(C.QGraphicsWidget_Actions(ptr.Pointer()))
	}
	return make([]*QAction, 0)
}

func (ptr *QGraphicsWidget) AddAction(action QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_AddAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

//export callbackQGraphicsWidget_BoundingRect
func callbackQGraphicsWidget_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "boundingRect"); signal != nil {
		return core.PointerFromQRectF((*(*func() *core.QRectF)(signal))())
	}

	return core.PointerFromQRectF(NewQGraphicsWidgetFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QGraphicsWidget) ConnectBoundingRect(f func() *core.QRectF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "boundingRect"); signal != nil {
			f := func() *core.QRectF {
				(*(*func() *core.QRectF)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "boundingRect", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "boundingRect", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsWidget) DisconnectBoundingRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "boundingRect")
	}
}

func (ptr *QGraphicsWidget) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsWidget_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) BoundingRectDefault() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsWidget_BoundingRectDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsWidget_Close
func callbackQGraphicsWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QGraphicsWidget) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsWidget) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QGraphicsWidget) Close() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsWidget_Close(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGraphicsWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQGraphicsWidget_CloseEvent
func callbackQGraphicsWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*gui.QCloseEvent))(signal))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "closeEvent"); signal != nil {
			f := func(event *gui.QCloseEvent) {
				(*(*func(*gui.QCloseEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "closeEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "closeEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsWidget) DisconnectCloseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "closeEvent")
	}
}

func (ptr *QGraphicsWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGraphicsWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGraphicsWidget) Font() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QGraphicsWidget_Font(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsWidget_HideEvent
func callbackQGraphicsWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hideEvent"); signal != nil {
			f := func(event *gui.QHideEvent) {
				(*(*func(*gui.QHideEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "hideEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hideEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsWidget) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hideEvent")
	}
}

func (ptr *QGraphicsWidget) HideEvent(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QGraphicsWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QGraphicsWidget) Layout() *QGraphicsLayout {
	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutFromPointer(C.QGraphicsWidget_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWidget) Rect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsWidget_Rect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) RemoveAction(action QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_RemoveAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QGraphicsWidget) Resize(size core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_Resize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QGraphicsWidget) Resize2(w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_Resize2(ptr.Pointer(), C.double(w), C.double(h))
	}
}

//export callbackQGraphicsWidget_ResizeEvent
func callbackQGraphicsWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*QGraphicsSceneResizeEvent))(signal))(NewQGraphicsSceneResizeEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).ResizeEventDefault(NewQGraphicsSceneResizeEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) ConnectResizeEvent(f func(event *QGraphicsSceneResizeEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resizeEvent"); signal != nil {
			f := func(event *QGraphicsSceneResizeEvent) {
				(*(*func(*QGraphicsSceneResizeEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "resizeEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resizeEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsWidget) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resizeEvent")
	}
}

func (ptr *QGraphicsWidget) ResizeEvent(event QGraphicsSceneResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ResizeEvent(ptr.Pointer(), PointerFromQGraphicsSceneResizeEvent(event))
	}
}

func (ptr *QGraphicsWidget) ResizeEventDefault(event QGraphicsSceneResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ResizeEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneResizeEvent(event))
	}
}

func (ptr *QGraphicsWidget) SetFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QGraphicsWidget) SetGeometry2(x float64, y float64, w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetGeometry2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsWidget) SetLayout(layout QGraphicsLayout_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetLayout(ptr.Pointer(), PointerFromQGraphicsLayout(layout))
	}
}

func (ptr *QGraphicsWidget) SetStyle(style QStyle_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetStyle(ptr.Pointer(), PointerFromQStyle(style))
	}
}

func (ptr *QGraphicsWidget) SetWindowTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QGraphicsWidget_SetWindowTitle(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

//export callbackQGraphicsWidget_ShowEvent
func callbackQGraphicsWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "showEvent"); signal != nil {
			f := func(event *gui.QShowEvent) {
				(*(*func(*gui.QShowEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "showEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "showEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsWidget) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "showEvent")
	}
}

func (ptr *QGraphicsWidget) ShowEvent(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QGraphicsWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QGraphicsWidget) Size() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsWidget_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsWidget_SizeHint
func callbackQGraphicsWidget_SizeHint(ptr unsafe.Pointer, which C.longlong, constraint unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSizeF((*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
	}

	return core.PointerFromQSizeF(NewQGraphicsWidgetFromPointer(ptr).SizeHintDefault(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
}

func (ptr *QGraphicsWidget) ConnectSizeHint(f func(which core.Qt__SizeHint, constraint *core.QSizeF) *core.QSizeF) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sizeHint"); signal != nil {
			f := func(which core.Qt__SizeHint, constraint *core.QSizeF) *core.QSizeF {
				(*(*func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(signal))(which, constraint)
				return f(which, constraint)
			}
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsWidget) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sizeHint")
	}
}

func (ptr *QGraphicsWidget) SizeHint(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsWidget_SizeHint(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) SizeHintDefault(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsWidget_SizeHintDefault(ptr.Pointer(), C.longlong(which), core.PointerFromQSizeF(constraint)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) Style() *QStyle {
	if ptr.Pointer() != nil {
		tmpValue := NewQStyleFromPointer(C.QGraphicsWidget_Style(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) WindowTitle() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QGraphicsWidget_WindowTitle(ptr.Pointer()))
	}
	return ""
}

//export callbackQGraphicsWidget_DestroyQGraphicsWidget
func callbackQGraphicsWidget_DestroyQGraphicsWidget(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGraphicsWidget"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGraphicsWidgetFromPointer(ptr).DestroyQGraphicsWidgetDefault()
	}
}

func (ptr *QGraphicsWidget) ConnectDestroyQGraphicsWidget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGraphicsWidget"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsWidget", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGraphicsWidget", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGraphicsWidget) DisconnectDestroyQGraphicsWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGraphicsWidget")
	}
}

func (ptr *QGraphicsWidget) DestroyQGraphicsWidget() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsWidget_DestroyQGraphicsWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsWidget) DestroyQGraphicsWidgetDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGraphicsWidget_DestroyQGraphicsWidgetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsWidget) Geometry() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGraphicsWidget_Geometry(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) MaximumSize() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsWidget_MaximumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) MinimumSize() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QGraphicsWidget_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) SetMinimumSize(minimumSize core.QSizeF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetMinimumSize(ptr.Pointer(), core.PointerFromQSizeF(minimumSize))
	}
}

func (ptr *QGraphicsWidget) SizePolicy() *QSizePolicy {
	if ptr.Pointer() != nil {
		tmpValue := NewQSizePolicyFromPointer(C.QGraphicsWidget_SizePolicy(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSizePolicy).DestroyQSizePolicy)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) SetSizePolicy(sizePolicy QSizePolicy_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetSizePolicy(ptr.Pointer(), PointerFromQSizePolicy(sizePolicy))
	}
}

func (ptr *QGraphicsWidget) __actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QGraphicsWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) __actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget___actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QGraphicsWidget) __actions_newList() unsafe.Pointer {
	return C.QGraphicsWidget___actions_newList(ptr.Pointer())
}

func (ptr *QGraphicsWidget) __addActions_actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QGraphicsWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) __addActions_actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget___addActions_actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QGraphicsWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QGraphicsWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QGraphicsWidget) __insertActions_actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QGraphicsWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsWidget) __insertActions_actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsWidget___insertActions_actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QGraphicsWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QGraphicsWidget___insertActions_actions_newList(ptr.Pointer())
}

type QGridLayout struct {
	QLayout
}

type QGridLayout_ITF interface {
	QLayout_ITF
	QGridLayout_PTR() *QGridLayout
}

func (ptr *QGridLayout) QGridLayout_PTR() *QGridLayout {
	return ptr
}

func (ptr *QGridLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QGridLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLayout_PTR().SetPointer(p)
	}
}

func PointerFromQGridLayout(ptr QGridLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGridLayout_PTR().Pointer()
	}
	return nil
}

func NewQGridLayoutFromPointer(ptr unsafe.Pointer) (n *QGridLayout) {
	n = new(QGridLayout)
	n.SetPointer(ptr)
	return
}
func NewQGridLayout(parent QWidget_ITF) *QGridLayout {
	tmpValue := NewQGridLayoutFromPointer(C.QGridLayout_NewQGridLayout(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQGridLayout2() *QGridLayout {
	tmpValue := NewQGridLayoutFromPointer(C.QGridLayout_NewQGridLayout2())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGridLayout) AddItem(item QLayoutItem_ITF, row int, column int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item), C.int(int32(row)), C.int(int32(column)), C.int(int32(rowSpan)), C.int(int32(columnSpan)), C.longlong(alignment))
	}
}

//export callbackQGridLayout_AddItem2
func callbackQGridLayout_AddItem2(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addItem2"); signal != nil {
		(*(*func(*QLayoutItem))(signal))(NewQLayoutItemFromPointer(item))
	} else {
		NewQGridLayoutFromPointer(ptr).AddItem2Default(NewQLayoutItemFromPointer(item))
	}
}

func (ptr *QGridLayout) ConnectAddItem2(f func(item *QLayoutItem)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addItem2"); signal != nil {
			f := func(item *QLayoutItem) {
				(*(*func(*QLayoutItem))(signal))(item)
				f(item)
			}
			qt.ConnectSignal(ptr.Pointer(), "addItem2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addItem2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGridLayout) DisconnectAddItem2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addItem2")
	}
}

func (ptr *QGridLayout) AddItem2(item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddItem2(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QGridLayout) AddItem2Default(item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddItem2Default(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QGridLayout) AddLayout(layout QLayout_ITF, row int, column int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddLayout(ptr.Pointer(), PointerFromQLayout(layout), C.int(int32(row)), C.int(int32(column)), C.longlong(alignment))
	}
}

func (ptr *QGridLayout) AddLayout2(layout QLayout_ITF, row int, column int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddLayout2(ptr.Pointer(), PointerFromQLayout(layout), C.int(int32(row)), C.int(int32(column)), C.int(int32(rowSpan)), C.int(int32(columnSpan)), C.longlong(alignment))
	}
}

func (ptr *QGridLayout) AddWidget2(widget QWidget_ITF, row int, column int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddWidget2(ptr.Pointer(), PointerFromQWidget(widget), C.int(int32(row)), C.int(int32(column)), C.longlong(alignment))
	}
}

func (ptr *QGridLayout) AddWidget3(widget QWidget_ITF, fromRow int, fromColumn int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGridLayout_AddWidget3(ptr.Pointer(), PointerFromQWidget(widget), C.int(int32(fromRow)), C.int(int32(fromColumn)), C.int(int32(rowSpan)), C.int(int32(columnSpan)), C.longlong(alignment))
	}
}

func (ptr *QGridLayout) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGridLayout_ColumnCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQGridLayout_Count
func callbackQGridLayout_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQGridLayoutFromPointer(ptr).CountDefault()))
}

func (ptr *QGridLayout) ConnectCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "count"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGridLayout) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "count")
	}
}

func (ptr *QGridLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGridLayout_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGridLayout) CountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGridLayout_CountDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQGridLayout_ItemAt
func callbackQGridLayout_ItemAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return PointerFromQLayoutItem((*(*func(int) *QLayoutItem)(signal))(int(int32(index))))
	}

	return PointerFromQLayoutItem(NewQGridLayoutFromPointer(ptr).ItemAtDefault(int(int32(index))))
}

func (ptr *QGridLayout) ConnectItemAt(f func(index int) *QLayoutItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "itemAt"); signal != nil {
			f := func(index int) *QLayoutItem {
				(*(*func(int) *QLayoutItem)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGridLayout) DisconnectItemAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "itemAt")
	}
}

func (ptr *QGridLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QGridLayout_ItemAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QGridLayout) ItemAtDefault(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QGridLayout_ItemAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QGridLayout) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGridLayout_RowCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQGridLayout_SizeHint
func callbackQGridLayout_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQGridLayoutFromPointer(ptr).SizeHintDefault())
}

func (ptr *QGridLayout) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sizeHint"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGridLayout) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sizeHint")
	}
}

func (ptr *QGridLayout) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QGridLayout_SizeHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QGridLayout) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QGridLayout_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQGridLayout_TakeAt
func callbackQGridLayout_TakeAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "takeAt"); signal != nil {
		return PointerFromQLayoutItem((*(*func(int) *QLayoutItem)(signal))(int(int32(index))))
	}

	return PointerFromQLayoutItem(NewQGridLayoutFromPointer(ptr).TakeAtDefault(int(int32(index))))
}

func (ptr *QGridLayout) ConnectTakeAt(f func(index int) *QLayoutItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "takeAt"); signal != nil {
			f := func(index int) *QLayoutItem {
				(*(*func(int) *QLayoutItem)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "takeAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "takeAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGridLayout) DisconnectTakeAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "takeAt")
	}
}

func (ptr *QGridLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QGridLayout_TakeAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QGridLayout) TakeAtDefault(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QGridLayout_TakeAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQGridLayout_DestroyQGridLayout
func callbackQGridLayout_DestroyQGridLayout(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGridLayout"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGridLayoutFromPointer(ptr).DestroyQGridLayoutDefault()
	}
}

func (ptr *QGridLayout) ConnectDestroyQGridLayout(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGridLayout"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGridLayout", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGridLayout", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGridLayout) DisconnectDestroyQGridLayout() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGridLayout")
	}
}

func (ptr *QGridLayout) DestroyQGridLayout() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGridLayout_DestroyQGridLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGridLayout) DestroyQGridLayoutDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGridLayout_DestroyQGridLayoutDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QHBoxLayout struct {
	QBoxLayout
}

type QHBoxLayout_ITF interface {
	QBoxLayout_ITF
	QHBoxLayout_PTR() *QHBoxLayout
}

func (ptr *QHBoxLayout) QHBoxLayout_PTR() *QHBoxLayout {
	return ptr
}

func (ptr *QHBoxLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QHBoxLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QBoxLayout_PTR().SetPointer(p)
	}
}

func PointerFromQHBoxLayout(ptr QHBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQHBoxLayoutFromPointer(ptr unsafe.Pointer) (n *QHBoxLayout) {
	n = new(QHBoxLayout)
	n.SetPointer(ptr)
	return
}
func NewQHBoxLayout() *QHBoxLayout {
	tmpValue := NewQHBoxLayoutFromPointer(C.QHBoxLayout_NewQHBoxLayout())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQHBoxLayout2(parent QWidget_ITF) *QHBoxLayout {
	tmpValue := NewQHBoxLayoutFromPointer(C.QHBoxLayout_NewQHBoxLayout2(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQHBoxLayout_DestroyQHBoxLayout
func callbackQHBoxLayout_DestroyQHBoxLayout(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHBoxLayout"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHBoxLayoutFromPointer(ptr).DestroyQHBoxLayoutDefault()
	}
}

func (ptr *QHBoxLayout) ConnectDestroyQHBoxLayout(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHBoxLayout"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHBoxLayout", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHBoxLayout", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHBoxLayout) DisconnectDestroyQHBoxLayout() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHBoxLayout")
	}
}

func (ptr *QHBoxLayout) DestroyQHBoxLayout() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHBoxLayout_DestroyQHBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHBoxLayout) DestroyQHBoxLayoutDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHBoxLayout_DestroyQHBoxLayoutDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QHeaderView struct {
	QAbstractItemView
}

type QHeaderView_ITF interface {
	QAbstractItemView_ITF
	QHeaderView_PTR() *QHeaderView
}

func (ptr *QHeaderView) QHeaderView_PTR() *QHeaderView {
	return ptr
}

func (ptr *QHeaderView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemView_PTR().Pointer()
	}
	return nil
}

func (ptr *QHeaderView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractItemView_PTR().SetPointer(p)
	}
}

func PointerFromQHeaderView(ptr QHeaderView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHeaderView_PTR().Pointer()
	}
	return nil
}

func NewQHeaderViewFromPointer(ptr unsafe.Pointer) (n *QHeaderView) {
	n = new(QHeaderView)
	n.SetPointer(ptr)
	return
}

// QHeaderView::ResizeMode
//
//go:generate stringer -type=QHeaderView__ResizeMode
type QHeaderView__ResizeMode int64

const (
	QHeaderView__Interactive      QHeaderView__ResizeMode = QHeaderView__ResizeMode(0)
	QHeaderView__Stretch          QHeaderView__ResizeMode = QHeaderView__ResizeMode(1)
	QHeaderView__Fixed            QHeaderView__ResizeMode = QHeaderView__ResizeMode(2)
	QHeaderView__ResizeToContents QHeaderView__ResizeMode = QHeaderView__ResizeMode(3)
	QHeaderView__Custom           QHeaderView__ResizeMode = QHeaderView__ResizeMode(QHeaderView__Fixed)
)

func NewQHeaderView(orientation core.Qt__Orientation, parent QWidget_ITF) *QHeaderView {
	tmpValue := NewQHeaderViewFromPointer(C.QHeaderView_NewQHeaderView(C.longlong(orientation), PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHeaderView) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHeaderView_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHeaderView_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QHeaderView_Orientation(ptr.Pointer()))
	}
	return 0
}

//export callbackQHeaderView_SetSelection
func callbackQHeaderView_SetSelection(ptr unsafe.Pointer, rect unsafe.Pointer, flags C.longlong) {
	if signal := qt.GetSignal(ptr, "setSelection"); signal != nil {
		(*(*func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(signal))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(flags))
	} else {
		NewQHeaderViewFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(flags))
	}
}

func (ptr *QHeaderView) ConnectSetSelection(f func(rect *core.QRect, flags core.QItemSelectionModel__SelectionFlag)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSelection"); signal != nil {
			f := func(rect *core.QRect, flags core.QItemSelectionModel__SelectionFlag) {
				(*(*func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(signal))(rect, flags)
				f(rect, flags)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSelection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSelection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHeaderView) DisconnectSetSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSelection")
	}
}

func (ptr *QHeaderView) SetSelection(rect core.QRect_ITF, flags core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(flags))
	}
}

func (ptr *QHeaderView) SetSelectionDefault(rect core.QRect_ITF, flags core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(flags))
	}
}

//export callbackQHeaderView_DestroyQHeaderView
func callbackQHeaderView_DestroyQHeaderView(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHeaderView"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHeaderViewFromPointer(ptr).DestroyQHeaderViewDefault()
	}
}

func (ptr *QHeaderView) ConnectDestroyQHeaderView(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHeaderView"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHeaderView", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHeaderView", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHeaderView) DisconnectDestroyQHeaderView() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHeaderView")
	}
}

func (ptr *QHeaderView) DestroyQHeaderView() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHeaderView_DestroyQHeaderView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHeaderView) DestroyQHeaderViewDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHeaderView_DestroyQHeaderViewDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHeaderView_HorizontalOffset
func callbackQHeaderView_HorizontalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "horizontalOffset"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQHeaderViewFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *QHeaderView) HorizontalOffset() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHeaderView_HorizontalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) HorizontalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHeaderView_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHeaderView_IndexAt
func callbackQHeaderView_IndexAt(ptr unsafe.Pointer, point unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "indexAt"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QPoint) *core.QModelIndex)(signal))(core.NewQPointFromPointer(point)))
	}

	return core.PointerFromQModelIndex(NewQHeaderViewFromPointer(ptr).IndexAtDefault(core.NewQPointFromPointer(point)))
}

func (ptr *QHeaderView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHeaderView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHeaderView) IndexAtDefault(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHeaderView_IndexAtDefault(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHeaderView_IsIndexHidden
func callbackQHeaderView_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHeaderViewFromPointer(ptr).IsIndexHiddenDefault(core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QHeaderView) IsIndexHidden(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHeaderView_IsIndexHidden(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QHeaderView) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHeaderView_IsIndexHiddenDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

//export callbackQHeaderView_MoveCursor
func callbackQHeaderView_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "moveCursor"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(signal))(QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(NewQHeaderViewFromPointer(ptr).MoveCursorDefault(QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *QHeaderView) MoveCursor(cursorAction QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHeaderView_MoveCursor(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QHeaderView) MoveCursorDefault(cursorAction QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QHeaderView_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQHeaderView_ScrollTo
func callbackQHeaderView_ScrollTo(ptr unsafe.Pointer, index unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "scrollTo"); signal != nil {
		(*(*func(*core.QModelIndex, QAbstractItemView__ScrollHint))(signal))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	} else {
		NewQHeaderViewFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QHeaderView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

func (ptr *QHeaderView) ScrollToDefault(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackQHeaderView_VerticalOffset
func callbackQHeaderView_VerticalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "verticalOffset"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQHeaderViewFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *QHeaderView) VerticalOffset() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHeaderView_VerticalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) VerticalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QHeaderView_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQHeaderView_VisualRect
func callbackQHeaderView_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRect"); signal != nil {
		return core.PointerFromQRect((*(*func(*core.QModelIndex) *core.QRect)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(NewQHeaderViewFromPointer(ptr).VisualRectDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QHeaderView) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QHeaderView_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QHeaderView) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QHeaderView_VisualRectDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQHeaderView_VisualRegionForSelection
func callbackQHeaderView_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion((*(*func(*core.QItemSelection) *gui.QRegion)(signal))(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(NewQHeaderViewFromPointer(ptr).VisualRegionForSelectionDefault(core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *QHeaderView) VisualRegionForSelection(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQRegionFromPointer(C.QHeaderView_VisualRegionForSelection(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		qt.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

func (ptr *QHeaderView) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQRegionFromPointer(C.QHeaderView_VisualRegionForSelectionDefault(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		qt.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

type QInputDialog struct {
	QDialog
}

type QInputDialog_ITF interface {
	QDialog_ITF
	QInputDialog_PTR() *QInputDialog
}

func (ptr *QInputDialog) QInputDialog_PTR() *QInputDialog {
	return ptr
}

func (ptr *QInputDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QInputDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQInputDialog(ptr QInputDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputDialog_PTR().Pointer()
	}
	return nil
}

func NewQInputDialogFromPointer(ptr unsafe.Pointer) (n *QInputDialog) {
	n = new(QInputDialog)
	n.SetPointer(ptr)
	return
}

// QInputDialog::InputDialogOption
//
//go:generate stringer -type=QInputDialog__InputDialogOption
type QInputDialog__InputDialogOption int64

const (
	QInputDialog__NoButtons                    QInputDialog__InputDialogOption = QInputDialog__InputDialogOption(0x00000001)
	QInputDialog__UseListViewForComboBoxItems  QInputDialog__InputDialogOption = QInputDialog__InputDialogOption(0x00000002)
	QInputDialog__UsePlainTextEditForTextInput QInputDialog__InputDialogOption = QInputDialog__InputDialogOption(0x00000004)
)

// QInputDialog::InputMode
//
//go:generate stringer -type=QInputDialog__InputMode
type QInputDialog__InputMode int64

const (
	QInputDialog__TextInput   QInputDialog__InputMode = QInputDialog__InputMode(0)
	QInputDialog__IntInput    QInputDialog__InputMode = QInputDialog__InputMode(1)
	QInputDialog__DoubleInput QInputDialog__InputMode = QInputDialog__InputMode(2)
)

func NewQInputDialog(parent QWidget_ITF, flags core.Qt__WindowType) *QInputDialog {
	tmpValue := NewQInputDialogFromPointer(C.QInputDialog_NewQInputDialog(PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQInputDialog_Done
func callbackQInputDialog_Done(ptr unsafe.Pointer, result C.int) {
	if signal := qt.GetSignal(ptr, "done"); signal != nil {
		(*(*func(int))(signal))(int(int32(result)))
	} else {
		NewQInputDialogFromPointer(ptr).DoneDefault(int(int32(result)))
	}
}

func (ptr *QInputDialog) ConnectDone(f func(result int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "done"); signal != nil {
			f := func(result int) {
				(*(*func(int))(signal))(result)
				f(result)
			}
			qt.ConnectSignal(ptr.Pointer(), "done", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "done", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QInputDialog) DisconnectDone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "done")
	}
}

func (ptr *QInputDialog) Done(result int) {
	if ptr.Pointer() != nil {
		C.QInputDialog_Done(ptr.Pointer(), C.int(int32(result)))
	}
}

func (ptr *QInputDialog) DoneDefault(result int) {
	if ptr.Pointer() != nil {
		C.QInputDialog_DoneDefault(ptr.Pointer(), C.int(int32(result)))
	}
}

func QInputDialog_GetItem(parent QWidget_ITF, title string, label string, items []string, current int, editable bool, ok *bool, flags core.Qt__WindowType, inputMethodHints core.Qt__InputMethodHint) string {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var labelC *C.char
	if label != "" {
		labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
	}
	itemsC := C.CString(strings.Join(items, "¡¦!"))
	defer C.free(unsafe.Pointer(itemsC))
	var okC C.char
	if ok != nil {
		okC = C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
	}
	return cGoUnpackString(C.QInputDialog_QInputDialog_GetItem(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: labelC, len: C.longlong(len(label))}, C.struct_QtWidgets_PackedString{data: itemsC, len: C.longlong(len(strings.Join(items, "¡¦!")))}, C.int(int32(current)), C.char(int8(qt.GoBoolToInt(editable))), &okC, C.longlong(flags), C.longlong(inputMethodHints)))
}

func (ptr *QInputDialog) GetItem(parent QWidget_ITF, title string, label string, items []string, current int, editable bool, ok *bool, flags core.Qt__WindowType, inputMethodHints core.Qt__InputMethodHint) string {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var labelC *C.char
	if label != "" {
		labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
	}
	itemsC := C.CString(strings.Join(items, "¡¦!"))
	defer C.free(unsafe.Pointer(itemsC))
	var okC C.char
	if ok != nil {
		okC = C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
	}
	return cGoUnpackString(C.QInputDialog_QInputDialog_GetItem(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: labelC, len: C.longlong(len(label))}, C.struct_QtWidgets_PackedString{data: itemsC, len: C.longlong(len(strings.Join(items, "¡¦!")))}, C.int(int32(current)), C.char(int8(qt.GoBoolToInt(editable))), &okC, C.longlong(flags), C.longlong(inputMethodHints)))
}

func QInputDialog_GetText(parent QWidget_ITF, title string, label string, mode QLineEdit__EchoMode, text string, ok *bool, flags core.Qt__WindowType, inputMethodHints core.Qt__InputMethodHint) string {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var labelC *C.char
	if label != "" {
		labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	var okC C.char
	if ok != nil {
		okC = C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
	}
	return cGoUnpackString(C.QInputDialog_QInputDialog_GetText(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: labelC, len: C.longlong(len(label))}, C.longlong(mode), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, &okC, C.longlong(flags), C.longlong(inputMethodHints)))
}

func (ptr *QInputDialog) GetText(parent QWidget_ITF, title string, label string, mode QLineEdit__EchoMode, text string, ok *bool, flags core.Qt__WindowType, inputMethodHints core.Qt__InputMethodHint) string {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var labelC *C.char
	if label != "" {
		labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	var okC C.char
	if ok != nil {
		okC = C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
	}
	return cGoUnpackString(C.QInputDialog_QInputDialog_GetText(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: labelC, len: C.longlong(len(label))}, C.longlong(mode), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, &okC, C.longlong(flags), C.longlong(inputMethodHints)))
}

func (ptr *QInputDialog) Open(receiver core.QObject_ITF, member string) {
	if ptr.Pointer() != nil {
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		C.QInputDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), memberC)
	}
}

//export callbackQInputDialog_DestroyQInputDialog
func callbackQInputDialog_DestroyQInputDialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QInputDialog"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQInputDialogFromPointer(ptr).DestroyQInputDialogDefault()
	}
}

func (ptr *QInputDialog) ConnectDestroyQInputDialog(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QInputDialog"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QInputDialog", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QInputDialog", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QInputDialog) DisconnectDestroyQInputDialog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QInputDialog")
	}
}

func (ptr *QInputDialog) DestroyQInputDialog() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QInputDialog_DestroyQInputDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QInputDialog) DestroyQInputDialogDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QInputDialog_DestroyQInputDialogDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QLCDNumber::Mode
//
//go:generate stringer -type=QLCDNumber__Mode
type QLCDNumber__Mode int64

const (
	QLCDNumber__Hex QLCDNumber__Mode = QLCDNumber__Mode(0)
	QLCDNumber__Dec QLCDNumber__Mode = QLCDNumber__Mode(1)
	QLCDNumber__Oct QLCDNumber__Mode = QLCDNumber__Mode(2)
	QLCDNumber__Bin QLCDNumber__Mode = QLCDNumber__Mode(3)
)

// QLCDNumber::SegmentStyle
//
//go:generate stringer -type=QLCDNumber__SegmentStyle
type QLCDNumber__SegmentStyle int64

var (
	QLCDNumber__Outline QLCDNumber__SegmentStyle = QLCDNumber__SegmentStyle(0)
	QLCDNumber__Filled  QLCDNumber__SegmentStyle = QLCDNumber__SegmentStyle(1)
	QLCDNumber__Flat    QLCDNumber__SegmentStyle = QLCDNumber__SegmentStyle(2)
)

type QLabel struct {
	QFrame
}

type QLabel_ITF interface {
	QFrame_ITF
	QLabel_PTR() *QLabel
}

func (ptr *QLabel) QLabel_PTR() *QLabel {
	return ptr
}

func (ptr *QLabel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func (ptr *QLabel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QFrame_PTR().SetPointer(p)
	}
}

func PointerFromQLabel(ptr QLabel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLabel_PTR().Pointer()
	}
	return nil
}

func NewQLabelFromPointer(ptr unsafe.Pointer) (n *QLabel) {
	n = new(QLabel)
	n.SetPointer(ptr)
	return
}
func NewQLabel(parent QWidget_ITF, ff core.Qt__WindowType) *QLabel {
	tmpValue := NewQLabelFromPointer(C.QLabel_NewQLabel(PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQLabel2(text string, parent QWidget_ITF, ff core.Qt__WindowType) *QLabel {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQLabelFromPointer(C.QLabel_NewQLabel2(C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLabel) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLabel_Alignment(ptr.Pointer()))
	}
	return 0
}

//export callbackQLabel_Clear
func callbackQLabel_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLabelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QLabel) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLabel) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *QLabel) Clear() {
	if ptr.Pointer() != nil {
		C.QLabel_Clear(ptr.Pointer())
	}
}

func (ptr *QLabel) ClearDefault() {
	if ptr.Pointer() != nil {
		C.QLabel_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QLabel) Indent() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLabel_Indent(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLabel) Margin() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLabel_Margin(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLabel) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLabel_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLabel) SelectionStart() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLabel_SelectionStart(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLabel) SetAlignment(vqt core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QLabel_SetAlignment(ptr.Pointer(), C.longlong(vqt))
	}
}

func (ptr *QLabel) SetIndent(vin int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetIndent(ptr.Pointer(), C.int(int32(vin)))
	}
}

//export callbackQLabel_SetPixmap
func callbackQLabel_SetPixmap(ptr unsafe.Pointer, vqp unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setPixmap"); signal != nil {
		(*(*func(*gui.QPixmap))(signal))(gui.NewQPixmapFromPointer(vqp))
	} else {
		NewQLabelFromPointer(ptr).SetPixmapDefault(gui.NewQPixmapFromPointer(vqp))
	}
}

func (ptr *QLabel) ConnectSetPixmap(f func(vqp *gui.QPixmap)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPixmap"); signal != nil {
			f := func(vqp *gui.QPixmap) {
				(*(*func(*gui.QPixmap))(signal))(vqp)
				f(vqp)
			}
			qt.ConnectSignal(ptr.Pointer(), "setPixmap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPixmap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLabel) DisconnectSetPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPixmap")
	}
}

func (ptr *QLabel) SetPixmap(vqp gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QLabel_SetPixmap(ptr.Pointer(), gui.PointerFromQPixmap(vqp))
	}
}

func (ptr *QLabel) SetPixmapDefault(vqp gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QLabel_SetPixmapDefault(ptr.Pointer(), gui.PointerFromQPixmap(vqp))
	}
}

func (ptr *QLabel) SetSelection(start int, length int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetSelection(ptr.Pointer(), C.int(int32(start)), C.int(int32(length)))
	}
}

//export callbackQLabel_SetText
func callbackQLabel_SetText(ptr unsafe.Pointer, vqs C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setText"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewQLabelFromPointer(ptr).SetTextDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QLabel) ConnectSetText(f func(vqs string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setText"); signal != nil {
			f := func(vqs string) {
				(*(*func(string))(signal))(vqs)
				f(vqs)
			}
			qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLabel) DisconnectSetText() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setText")
	}
}

func (ptr *QLabel) SetText(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QLabel_SetText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QLabel) SetTextDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QLabel_SetTextDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QLabel) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLabel_Text(ptr.Pointer()))
	}
	return ""
}

//export callbackQLabel_DestroyQLabel
func callbackQLabel_DestroyQLabel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLabel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLabelFromPointer(ptr).DestroyQLabelDefault()
	}
}

func (ptr *QLabel) ConnectDestroyQLabel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLabel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLabel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLabel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLabel) DisconnectDestroyQLabel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLabel")
	}
}

func (ptr *QLabel) DestroyQLabel() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLabel_DestroyQLabel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLabel) DestroyQLabelDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLabel_DestroyQLabelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLayout struct {
	core.QObject
	QLayoutItem
}

type QLayout_ITF interface {
	core.QObject_ITF
	QLayoutItem_ITF
	QLayout_PTR() *QLayout
}

func (ptr *QLayout) QLayout_PTR() *QLayout {
	return ptr
}

func (ptr *QLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QLayoutItem_PTR().SetPointer(p)
	}
}

func PointerFromQLayout(ptr QLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayout_PTR().Pointer()
	}
	return nil
}

func NewQLayoutFromPointer(ptr unsafe.Pointer) (n *QLayout) {
	n = new(QLayout)
	n.SetPointer(ptr)
	return
}

// QLayout::SizeConstraint
//
//go:generate stringer -type=QLayout__SizeConstraint
type QLayout__SizeConstraint int64

const (
	QLayout__SetDefaultConstraint QLayout__SizeConstraint = QLayout__SizeConstraint(0)
	QLayout__SetNoConstraint      QLayout__SizeConstraint = QLayout__SizeConstraint(1)
	QLayout__SetMinimumSize       QLayout__SizeConstraint = QLayout__SizeConstraint(2)
	QLayout__SetFixedSize         QLayout__SizeConstraint = QLayout__SizeConstraint(3)
	QLayout__SetMaximumSize       QLayout__SizeConstraint = QLayout__SizeConstraint(4)
	QLayout__SetMinAndMaxSize     QLayout__SizeConstraint = QLayout__SizeConstraint(5)
)

func NewQLayout(parent QWidget_ITF) *QLayout {
	tmpValue := NewQLayoutFromPointer(C.QLayout_NewQLayout(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQLayout2() *QLayout {
	tmpValue := NewQLayoutFromPointer(C.QLayout_NewQLayout2())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLayout) Activate() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_Activate(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLayout_AddItem
func callbackQLayout_AddItem(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addItem"); signal != nil {
		(*(*func(*QLayoutItem))(signal))(NewQLayoutItemFromPointer(item))
	}

}

func (ptr *QLayout) ConnectAddItem(f func(item *QLayoutItem)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addItem"); signal != nil {
			f := func(item *QLayoutItem) {
				(*(*func(*QLayoutItem))(signal))(item)
				f(item)
			}
			qt.ConnectSignal(ptr.Pointer(), "addItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectAddItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addItem")
	}
}

func (ptr *QLayout) AddItem(item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QLayout) AddWidget(w QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(w))
	}
}

//export callbackQLayout_ChildEvent
func callbackQLayout_ChildEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(e))
	} else {
		NewQLayoutFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(e))
	}
}

func (ptr *QLayout) ConnectChildEvent(f func(e *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "childEvent"); signal != nil {
			f := func(e *core.QChildEvent) {
				(*(*func(*core.QChildEvent))(signal))(e)
				f(e)
			}
			qt.ConnectSignal(ptr.Pointer(), "childEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "childEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "childEvent")
	}
}

func (ptr *QLayout) ChildEvent(e core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(e))
	}
}

func (ptr *QLayout) ChildEventDefault(e core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(e))
	}
}

//export callbackQLayout_Count
func callbackQLayout_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(0))
}

func (ptr *QLayout) ConnectCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "count"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "count", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "count")
	}
}

func (ptr *QLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLayout_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQLayout_ExpandingDirections
func callbackQLayout_ExpandingDirections(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expandingDirections"); signal != nil {
		return C.longlong((*(*func() core.Qt__Orientation)(signal))())
	}

	return C.longlong(NewQLayoutFromPointer(ptr).ExpandingDirectionsDefault())
}

func (ptr *QLayout) ConnectExpandingDirections(f func() core.Qt__Orientation) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "expandingDirections"); signal != nil {
			f := func() core.Qt__Orientation {
				(*(*func() core.Qt__Orientation)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "expandingDirections", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "expandingDirections", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectExpandingDirections() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "expandingDirections")
	}
}

func (ptr *QLayout) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayout_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) ExpandingDirectionsDefault() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayout_ExpandingDirectionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLayout_Geometry
func callbackQLayout_Geometry(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "geometry"); signal != nil {
		return core.PointerFromQRect((*(*func() *core.QRect)(signal))())
	}

	return core.PointerFromQRect(NewQLayoutFromPointer(ptr).GeometryDefault())
}

func (ptr *QLayout) ConnectGeometry(f func() *core.QRect) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "geometry"); signal != nil {
			f := func() *core.QRect {
				(*(*func() *core.QRect)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "geometry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "geometry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "geometry")
	}
}

func (ptr *QLayout) Geometry() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QLayout_Geometry(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) GeometryDefault() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QLayout_GeometryDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQLayout_IndexOf
func callbackQLayout_IndexOf(ptr unsafe.Pointer, widget unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "indexOf"); signal != nil {
		return C.int(int32((*(*func(*QWidget) int)(signal))(NewQWidgetFromPointer(widget))))
	}

	return C.int(int32(NewQLayoutFromPointer(ptr).IndexOfDefault(NewQWidgetFromPointer(widget))))
}

func (ptr *QLayout) ConnectIndexOf(f func(widget *QWidget) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "indexOf"); signal != nil {
			f := func(widget *QWidget) int {
				(*(*func(*QWidget) int)(signal))(widget)
				return f(widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "indexOf", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexOf", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectIndexOf() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "indexOf")
	}
}

func (ptr *QLayout) IndexOf(widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLayout_IndexOf(ptr.Pointer(), PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QLayout) IndexOfDefault(widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLayout_IndexOfDefault(ptr.Pointer(), PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QLayout) IndexOf2(layoutItem QLayoutItem_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLayout_IndexOf2(ptr.Pointer(), PointerFromQLayoutItem(layoutItem))))
	}
	return 0
}

//export callbackQLayout_IsEmpty
func callbackQLayout_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLayoutFromPointer(ptr).IsEmptyDefault())))
}

func (ptr *QLayout) ConnectIsEmpty(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isEmpty"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "isEmpty", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isEmpty", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectIsEmpty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isEmpty")
	}
}

func (ptr *QLayout) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLayout) IsEmptyDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_IsEmptyDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLayout_ItemAt
func callbackQLayout_ItemAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return PointerFromQLayoutItem((*(*func(int) *QLayoutItem)(signal))(int(int32(index))))
	}

	return PointerFromQLayoutItem(NewQLayoutItem(0))
}

func (ptr *QLayout) ConnectItemAt(f func(index int) *QLayoutItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "itemAt"); signal != nil {
			f := func(index int) *QLayoutItem {
				(*(*func(int) *QLayoutItem)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "itemAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectItemAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "itemAt")
	}
}

func (ptr *QLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_ItemAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQLayout_MaximumSize
func callbackQLayout_MaximumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "maximumSize"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQLayoutFromPointer(ptr).MaximumSizeDefault())
}

func (ptr *QLayout) ConnectMaximumSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "maximumSize"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "maximumSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maximumSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectMaximumSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "maximumSize")
	}
}

func (ptr *QLayout) MaximumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_MaximumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) MaximumSizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_MaximumSizeDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) MenuBar() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QLayout_MenuBar(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQLayout_MinimumSize
func callbackQLayout_MinimumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSize"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQLayoutFromPointer(ptr).MinimumSizeDefault())
}

func (ptr *QLayout) ConnectMinimumSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumSize"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectMinimumSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "minimumSize")
	}
}

func (ptr *QLayout) MinimumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) MinimumSizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_MinimumSizeDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) SetAlignment(w QWidget_ITF, alignment core.Qt__AlignmentFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_SetAlignment(ptr.Pointer(), PointerFromQWidget(w), C.longlong(alignment))) != 0
	}
	return false
}

func (ptr *QLayout) SetAlignment2(l QLayout_ITF, alignment core.Qt__AlignmentFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_SetAlignment2(ptr.Pointer(), PointerFromQLayout(l), C.longlong(alignment))) != 0
	}
	return false
}

func (ptr *QLayout) SetEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QLayout_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQLayout_SetGeometry
func callbackQLayout_SetGeometry(ptr unsafe.Pointer, r unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		(*(*func(*core.QRect))(signal))(core.NewQRectFromPointer(r))
	} else {
		NewQLayoutFromPointer(ptr).SetGeometryDefault(core.NewQRectFromPointer(r))
	}
}

func (ptr *QLayout) ConnectSetGeometry(f func(r *core.QRect)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGeometry"); signal != nil {
			f := func(r *core.QRect) {
				(*(*func(*core.QRect))(signal))(r)
				f(r)
			}
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectSetGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGeometry")
	}
}

func (ptr *QLayout) SetGeometry(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QLayout) SetGeometryDefault(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

//export callbackQLayout_TakeAt
func callbackQLayout_TakeAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "takeAt"); signal != nil {
		return PointerFromQLayoutItem((*(*func(int) *QLayoutItem)(signal))(int(int32(index))))
	}

	return PointerFromQLayoutItem(NewQLayoutItem(0))
}

func (ptr *QLayout) ConnectTakeAt(f func(index int) *QLayoutItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "takeAt"); signal != nil {
			f := func(index int) *QLayoutItem {
				(*(*func(int) *QLayoutItem)(signal))(index)
				return f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "takeAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "takeAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayout) DisconnectTakeAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "takeAt")
	}
}

func (ptr *QLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_TakeAt(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

func (ptr *QLayout) Update() {
	if ptr.Pointer() != nil {
		C.QLayout_Update(ptr.Pointer())
	}
}

func (ptr *QLayout) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLayout___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLayout) __children_newList() unsafe.Pointer {
	return C.QLayout___children_newList(ptr.Pointer())
}

func (ptr *QLayout) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLayout___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QLayout) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QLayout___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QLayout) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLayout___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLayout) __findChildren_newList() unsafe.Pointer {
	return C.QLayout___findChildren_newList(ptr.Pointer())
}

func (ptr *QLayout) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLayout___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLayout) __findChildren_newList3() unsafe.Pointer {
	return C.QLayout___findChildren_newList3(ptr.Pointer())
}

//export callbackQLayout_ConnectNotify
func callbackQLayout_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLayoutFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLayout) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLayout) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLayout_CustomEvent
func callbackQLayout_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQLayoutFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLayout) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLayout) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLayout_DeleteLater
func callbackQLayout_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLayoutFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLayout) DeleteLater() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLayout_DeleteLater(ptr.Pointer())
	}
}

func (ptr *QLayout) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLayout_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQLayout_Destroyed
func callbackQLayout_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQLayout_DisconnectNotify
func callbackQLayout_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLayoutFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLayout) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLayout) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLayout_Event
func callbackQLayout_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLayoutFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLayout) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_Event(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

func (ptr *QLayout) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQLayout_EventFilter
func callbackQLayout_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLayoutFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLayout) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QLayout) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayout_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLayout_ObjectNameChanged
func callbackQLayout_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQLayout_TimerEvent
func callbackQLayout_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLayoutFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLayout) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLayout) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLayout_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLayout_SizeHint
func callbackQLayout_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQLayoutFromPointer(ptr).SizeHintDefault())
}

func (ptr *QLayout) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_SizeHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QLayout) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayout_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

type QLayoutItem struct {
	ptr unsafe.Pointer
}

type QLayoutItem_ITF interface {
	QLayoutItem_PTR() *QLayoutItem
}

func (ptr *QLayoutItem) QLayoutItem_PTR() *QLayoutItem {
	return ptr
}

func (ptr *QLayoutItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLayoutItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQLayoutItem(ptr QLayoutItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayoutItem_PTR().Pointer()
	}
	return nil
}

func NewQLayoutItemFromPointer(ptr unsafe.Pointer) (n *QLayoutItem) {
	n = new(QLayoutItem)
	n.SetPointer(ptr)
	return
}
func NewQLayoutItem(alignment core.Qt__AlignmentFlag) *QLayoutItem {
	return NewQLayoutItemFromPointer(C.QLayoutItem_NewQLayoutItem(C.longlong(alignment)))
}

func (ptr *QLayoutItem) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLayoutItem_Alignment(ptr.Pointer()))
	}
	return 0
}

//export callbackQLayoutItem_ControlTypes
func callbackQLayoutItem_ControlTypes(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "controlTypes"); signal != nil {
		return C.longlong((*(*func() QSizePolicy__ControlType)(signal))())
	}

	return C.longlong(NewQLayoutItemFromPointer(ptr).ControlTypesDefault())
}

func (ptr *QLayoutItem) ConnectControlTypes(f func() QSizePolicy__ControlType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "controlTypes"); signal != nil {
			f := func() QSizePolicy__ControlType {
				(*(*func() QSizePolicy__ControlType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "controlTypes", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "controlTypes", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectControlTypes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "controlTypes")
	}
}

func (ptr *QLayoutItem) ControlTypes() QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QLayoutItem_ControlTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayoutItem) ControlTypesDefault() QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QLayoutItem_ControlTypesDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLayoutItem_ExpandingDirections
func callbackQLayoutItem_ExpandingDirections(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expandingDirections"); signal != nil {
		return C.longlong((*(*func() core.Qt__Orientation)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QLayoutItem) ConnectExpandingDirections(f func() core.Qt__Orientation) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "expandingDirections"); signal != nil {
			f := func() core.Qt__Orientation {
				(*(*func() core.Qt__Orientation)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "expandingDirections", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "expandingDirections", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectExpandingDirections() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "expandingDirections")
	}
}

func (ptr *QLayoutItem) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayoutItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

//export callbackQLayoutItem_Geometry
func callbackQLayoutItem_Geometry(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "geometry"); signal != nil {
		return core.PointerFromQRect((*(*func() *core.QRect)(signal))())
	}

	return core.PointerFromQRect(core.NewQRect())
}

func (ptr *QLayoutItem) ConnectGeometry(f func() *core.QRect) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "geometry"); signal != nil {
			f := func() *core.QRect {
				(*(*func() *core.QRect)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "geometry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "geometry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "geometry")
	}
}

func (ptr *QLayoutItem) Geometry() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QLayoutItem_Geometry(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQLayoutItem_HasHeightForWidth
func callbackQLayoutItem_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLayoutItemFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QLayoutItem) ConnectHasHeightForWidth(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasHeightForWidth"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hasHeightForWidth", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasHeightForWidth", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectHasHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasHeightForWidth")
	}
}

func (ptr *QLayoutItem) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayoutItem_HasHeightForWidth(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLayoutItem) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayoutItem_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLayoutItem_HeightForWidth
func callbackQLayoutItem_HeightForWidth(ptr unsafe.Pointer, vin C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(vin)))))
	}

	return C.int(int32(NewQLayoutItemFromPointer(ptr).HeightForWidthDefault(int(int32(vin)))))
}

func (ptr *QLayoutItem) ConnectHeightForWidth(f func(vin int) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "heightForWidth"); signal != nil {
			f := func(vin int) int {
				(*(*func(int) int)(signal))(vin)
				return f(vin)
			}
			qt.ConnectSignal(ptr.Pointer(), "heightForWidth", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "heightForWidth", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "heightForWidth")
	}
}

func (ptr *QLayoutItem) HeightForWidth(vin int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLayoutItem_HeightForWidth(ptr.Pointer(), C.int(int32(vin)))))
	}
	return 0
}

func (ptr *QLayoutItem) HeightForWidthDefault(vin int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLayoutItem_HeightForWidthDefault(ptr.Pointer(), C.int(int32(vin)))))
	}
	return 0
}

//export callbackQLayoutItem_Invalidate
func callbackQLayoutItem_Invalidate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "invalidate"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLayoutItemFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QLayoutItem) ConnectInvalidate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "invalidate"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "invalidate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "invalidate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectInvalidate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "invalidate")
	}
}

func (ptr *QLayoutItem) Invalidate() {
	if ptr.Pointer() != nil {
		C.QLayoutItem_Invalidate(ptr.Pointer())
	}
}

func (ptr *QLayoutItem) InvalidateDefault() {
	if ptr.Pointer() != nil {
		C.QLayoutItem_InvalidateDefault(ptr.Pointer())
	}
}

//export callbackQLayoutItem_IsEmpty
func callbackQLayoutItem_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QLayoutItem) ConnectIsEmpty(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isEmpty"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "isEmpty", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isEmpty", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectIsEmpty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isEmpty")
	}
}

func (ptr *QLayoutItem) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLayoutItem_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLayoutItem_Layout
func callbackQLayoutItem_Layout(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "layout"); signal != nil {
		return PointerFromQLayout((*(*func() *QLayout)(signal))())
	}

	return PointerFromQLayout(NewQLayoutItemFromPointer(ptr).LayoutDefault())
}

func (ptr *QLayoutItem) ConnectLayout(f func() *QLayout) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "layout"); signal != nil {
			f := func() *QLayout {
				(*(*func() *QLayout)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "layout", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "layout", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectLayout() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "layout")
	}
}

func (ptr *QLayoutItem) Layout() *QLayout {
	if ptr.Pointer() != nil {
		tmpValue := NewQLayoutFromPointer(C.QLayoutItem_Layout(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayoutItem) LayoutDefault() *QLayout {
	if ptr.Pointer() != nil {
		tmpValue := NewQLayoutFromPointer(C.QLayoutItem_LayoutDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQLayoutItem_MaximumSize
func callbackQLayoutItem_MaximumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "maximumSize"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QLayoutItem) ConnectMaximumSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "maximumSize"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "maximumSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maximumSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectMaximumSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "maximumSize")
	}
}

func (ptr *QLayoutItem) MaximumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayoutItem_MaximumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQLayoutItem_MinimumHeightForWidth
func callbackQLayoutItem_MinimumHeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "minimumHeightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewQLayoutItemFromPointer(ptr).MinimumHeightForWidthDefault(int(int32(w)))))
}

func (ptr *QLayoutItem) ConnectMinimumHeightForWidth(f func(w int) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumHeightForWidth"); signal != nil {
			f := func(w int) int {
				(*(*func(int) int)(signal))(w)
				return f(w)
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumHeightForWidth", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumHeightForWidth", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectMinimumHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "minimumHeightForWidth")
	}
}

func (ptr *QLayoutItem) MinimumHeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLayoutItem_MinimumHeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QLayoutItem) MinimumHeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLayoutItem_MinimumHeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQLayoutItem_MinimumSize
func callbackQLayoutItem_MinimumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSize"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QLayoutItem) ConnectMinimumSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumSize"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectMinimumSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "minimumSize")
	}
}

func (ptr *QLayoutItem) MinimumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayoutItem_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QLayoutItem) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QLayoutItem_SetAlignment(ptr.Pointer(), C.longlong(alignment))
	}
}

//export callbackQLayoutItem_SetGeometry
func callbackQLayoutItem_SetGeometry(ptr unsafe.Pointer, r unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		(*(*func(*core.QRect))(signal))(core.NewQRectFromPointer(r))
	}

}

func (ptr *QLayoutItem) ConnectSetGeometry(f func(r *core.QRect)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGeometry"); signal != nil {
			f := func(r *core.QRect) {
				(*(*func(*core.QRect))(signal))(r)
				f(r)
			}
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectSetGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGeometry")
	}
}

func (ptr *QLayoutItem) SetGeometry(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QLayoutItem_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

//export callbackQLayoutItem_SizeHint
func callbackQLayoutItem_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QLayoutItem) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sizeHint"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sizeHint")
	}
}

func (ptr *QLayoutItem) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QLayoutItem_SizeHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQLayoutItem_SpacerItem
func callbackQLayoutItem_SpacerItem(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "spacerItem"); signal != nil {
		return PointerFromQSpacerItem((*(*func() *QSpacerItem)(signal))())
	}

	return PointerFromQSpacerItem(NewQLayoutItemFromPointer(ptr).SpacerItemDefault())
}

func (ptr *QLayoutItem) ConnectSpacerItem(f func() *QSpacerItem) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "spacerItem"); signal != nil {
			f := func() *QSpacerItem {
				(*(*func() *QSpacerItem)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "spacerItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "spacerItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectSpacerItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "spacerItem")
	}
}

func (ptr *QLayoutItem) SpacerItem() *QSpacerItem {
	if ptr.Pointer() != nil {
		return NewQSpacerItemFromPointer(C.QLayoutItem_SpacerItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayoutItem) SpacerItemDefault() *QSpacerItem {
	if ptr.Pointer() != nil {
		return NewQSpacerItemFromPointer(C.QLayoutItem_SpacerItemDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQLayoutItem_Widget
func callbackQLayoutItem_Widget(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "widget"); signal != nil {
		return PointerFromQWidget((*(*func() *QWidget)(signal))())
	}

	return PointerFromQWidget(NewQLayoutItemFromPointer(ptr).WidgetDefault())
}

func (ptr *QLayoutItem) ConnectWidget(f func() *QWidget) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "widget"); signal != nil {
			f := func() *QWidget {
				(*(*func() *QWidget)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "widget", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "widget", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "widget")
	}
}

func (ptr *QLayoutItem) Widget() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QLayoutItem_Widget(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLayoutItem) WidgetDefault() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QLayoutItem_WidgetDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQLayoutItem_DestroyQLayoutItem
func callbackQLayoutItem_DestroyQLayoutItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLayoutItem"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLayoutItemFromPointer(ptr).DestroyQLayoutItemDefault()
	}
}

func (ptr *QLayoutItem) ConnectDestroyQLayoutItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLayoutItem"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLayoutItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLayoutItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLayoutItem) DisconnectDestroyQLayoutItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLayoutItem")
	}
}

func (ptr *QLayoutItem) DestroyQLayoutItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLayoutItem_DestroyQLayoutItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLayoutItem) DestroyQLayoutItemDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLayoutItem_DestroyQLayoutItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLineEdit struct {
	QWidget
}

type QLineEdit_ITF interface {
	QWidget_ITF
	QLineEdit_PTR() *QLineEdit
}

func (ptr *QLineEdit) QLineEdit_PTR() *QLineEdit {
	return ptr
}

func (ptr *QLineEdit) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QLineEdit) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQLineEdit(ptr QLineEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLineEdit_PTR().Pointer()
	}
	return nil
}

func NewQLineEditFromPointer(ptr unsafe.Pointer) (n *QLineEdit) {
	n = new(QLineEdit)
	n.SetPointer(ptr)
	return
}

// QLineEdit::ActionPosition
//
//go:generate stringer -type=QLineEdit__ActionPosition
type QLineEdit__ActionPosition int64

const (
	QLineEdit__LeadingPosition  QLineEdit__ActionPosition = QLineEdit__ActionPosition(0)
	QLineEdit__TrailingPosition QLineEdit__ActionPosition = QLineEdit__ActionPosition(1)
)

// QLineEdit::EchoMode
//
//go:generate stringer -type=QLineEdit__EchoMode
type QLineEdit__EchoMode int64

const (
	QLineEdit__Normal             QLineEdit__EchoMode = QLineEdit__EchoMode(0)
	QLineEdit__NoEcho             QLineEdit__EchoMode = QLineEdit__EchoMode(1)
	QLineEdit__Password           QLineEdit__EchoMode = QLineEdit__EchoMode(2)
	QLineEdit__PasswordEchoOnEdit QLineEdit__EchoMode = QLineEdit__EchoMode(3)
)

func NewQLineEdit(parent QWidget_ITF) *QLineEdit {
	tmpValue := NewQLineEditFromPointer(C.QLineEdit_NewQLineEdit(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQLineEdit2(contents string, parent QWidget_ITF) *QLineEdit {
	var contentsC *C.char
	if contents != "" {
		contentsC = C.CString(contents)
		defer C.free(unsafe.Pointer(contentsC))
	}
	tmpValue := NewQLineEditFromPointer(C.QLineEdit_NewQLineEdit2(C.struct_QtWidgets_PackedString{data: contentsC, len: C.longlong(len(contents))}, PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLineEdit) AddAction(action QAction_ITF, position QLineEdit__ActionPosition) {
	if ptr.Pointer() != nil {
		C.QLineEdit_AddAction(ptr.Pointer(), PointerFromQAction(action), C.longlong(position))
	}
}

func (ptr *QLineEdit) AddAction2(icon gui.QIcon_ITF, position QLineEdit__ActionPosition) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QLineEdit_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.longlong(position)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLineEdit) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLineEdit_Alignment(ptr.Pointer()))
	}
	return 0
}

//export callbackQLineEdit_Clear
func callbackQLineEdit_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLineEditFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QLineEdit) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLineEdit) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *QLineEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QLineEdit) ClearDefault() {
	if ptr.Pointer() != nil {
		C.QLineEdit_ClearDefault(ptr.Pointer())
	}
}

//export callbackQLineEdit_Copy
func callbackQLineEdit_Copy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "copy"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLineEditFromPointer(ptr).CopyDefault()
	}
}

func (ptr *QLineEdit) ConnectCopy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "copy"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "copy", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "copy", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLineEdit) DisconnectCopy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "copy")
	}
}

func (ptr *QLineEdit) Copy() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QLineEdit) CopyDefault() {
	if ptr.Pointer() != nil {
		C.QLineEdit_CopyDefault(ptr.Pointer())
	}
}

func (ptr *QLineEdit) CreateStandardContextMenu() *QMenu {
	if ptr.Pointer() != nil {
		tmpValue := NewQMenuFromPointer(C.QLineEdit_CreateStandardContextMenu(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLineEdit) CursorPosition() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLineEdit_CursorPosition(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLineEdit) Del() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Del(ptr.Pointer())
	}
}

func (ptr *QLineEdit) EchoMode() QLineEdit__EchoMode {
	if ptr.Pointer() != nil {
		return QLineEdit__EchoMode(C.QLineEdit_EchoMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQLineEdit_EditingFinished
func callbackQLineEdit_EditingFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "editingFinished"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLineEdit) ConnectEditingFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "editingFinished") {
			C.QLineEdit_ConnectEditingFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "editingFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "editingFinished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "editingFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "editingFinished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLineEdit) DisconnectEditingFinished() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectEditingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "editingFinished")
	}
}

func (ptr *QLineEdit) EditingFinished() {
	if ptr.Pointer() != nil {
		C.QLineEdit_EditingFinished(ptr.Pointer())
	}
}

func (ptr *QLineEdit) End(mark bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_End(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(mark))))
	}
}

func (ptr *QLineEdit) Home(mark bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_Home(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(mark))))
	}
}

func (ptr *QLineEdit) Insert(newText string) {
	if ptr.Pointer() != nil {
		var newTextC *C.char
		if newText != "" {
			newTextC = C.CString(newText)
			defer C.free(unsafe.Pointer(newTextC))
		}
		C.QLineEdit_Insert(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: newTextC, len: C.longlong(len(newText))})
	}
}

func (ptr *QLineEdit) IsModified() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLineEdit_IsModified(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQLineEdit_Paste
func callbackQLineEdit_Paste(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paste"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLineEditFromPointer(ptr).PasteDefault()
	}
}

func (ptr *QLineEdit) ConnectPaste(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paste"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "paste", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paste", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLineEdit) DisconnectPaste() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paste")
	}
}

func (ptr *QLineEdit) Paste() {
	if ptr.Pointer() != nil {
		C.QLineEdit_Paste(ptr.Pointer())
	}
}

func (ptr *QLineEdit) PasteDefault() {
	if ptr.Pointer() != nil {
		C.QLineEdit_PasteDefault(ptr.Pointer())
	}
}

func (ptr *QLineEdit) PlaceholderText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLineEdit_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLineEdit) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLineEdit_SelectedText(ptr.Pointer()))
	}
	return ""
}

//export callbackQLineEdit_SelectionChanged
func callbackQLineEdit_SelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLineEdit) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionChanged") {
			C.QLineEdit_ConnectSelectionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectionChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLineEdit) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionChanged")
	}
}

func (ptr *QLineEdit) SelectionChanged() {
	if ptr.Pointer() != nil {
		C.QLineEdit_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QLineEdit) SelectionEnd() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLineEdit_SelectionEnd(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLineEdit) SelectionStart() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLineEdit_SelectionStart(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLineEdit) SetAlignment(flag core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetAlignment(ptr.Pointer(), C.longlong(flag))
	}
}

func (ptr *QLineEdit) SetEchoMode(vql QLineEdit__EchoMode) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetEchoMode(ptr.Pointer(), C.longlong(vql))
	}
}

func (ptr *QLineEdit) SetModified(vbo bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QLineEdit) SetPlaceholderText(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QLineEdit_SetPlaceholderText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QLineEdit) SetReadOnly(vbo bool) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetReadOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QLineEdit) SetSelection(start int, length int) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetSelection(ptr.Pointer(), C.int(int32(start)), C.int(int32(length)))
	}
}

//export callbackQLineEdit_SetText
func callbackQLineEdit_SetText(ptr unsafe.Pointer, vqs C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setText"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewQLineEditFromPointer(ptr).SetTextDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QLineEdit) ConnectSetText(f func(vqs string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setText"); signal != nil {
			f := func(vqs string) {
				(*(*func(string))(signal))(vqs)
				f(vqs)
			}
			qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLineEdit) DisconnectSetText() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setText")
	}
}

func (ptr *QLineEdit) SetText(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QLineEdit_SetText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QLineEdit) SetTextDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QLineEdit_SetTextDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QLineEdit) SetValidator(v gui.QValidator_ITF) {
	if ptr.Pointer() != nil {
		C.QLineEdit_SetValidator(ptr.Pointer(), gui.PointerFromQValidator(v))
	}
}

func (ptr *QLineEdit) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLineEdit_Text(ptr.Pointer()))
	}
	return ""
}

//export callbackQLineEdit_TextChanged
func callbackQLineEdit_TextChanged(ptr unsafe.Pointer, text C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "textChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	}

}

func (ptr *QLineEdit) ConnectTextChanged(f func(text string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "textChanged") {
			C.QLineEdit_ConnectTextChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "textChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "textChanged"); signal != nil {
			f := func(text string) {
				(*(*func(string))(signal))(text)
				f(text)
			}
			qt.ConnectSignal(ptr.Pointer(), "textChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLineEdit) DisconnectTextChanged() {
	if ptr.Pointer() != nil {
		C.QLineEdit_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "textChanged")
	}
}

func (ptr *QLineEdit) TextChanged(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QLineEdit_TextChanged(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QLineEdit) Validator() *gui.QValidator {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQValidatorFromPointer(C.QLineEdit_Validator(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQLineEdit_DestroyQLineEdit
func callbackQLineEdit_DestroyQLineEdit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLineEdit"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLineEditFromPointer(ptr).DestroyQLineEditDefault()
	}
}

func (ptr *QLineEdit) ConnectDestroyQLineEdit(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLineEdit"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLineEdit", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLineEdit", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLineEdit) DisconnectDestroyQLineEdit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLineEdit")
	}
}

func (ptr *QLineEdit) DestroyQLineEdit() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLineEdit_DestroyQLineEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLineEdit) DestroyQLineEditDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLineEdit_DestroyQLineEditDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QListView::Flow
//
//go:generate stringer -type=QListView__Flow
type QListView__Flow int64

const (
	QListView__LeftToRight QListView__Flow = QListView__Flow(0)
	QListView__TopToBottom QListView__Flow = QListView__Flow(1)
)

// QListView::LayoutMode
//
//go:generate stringer -type=QListView__LayoutMode
type QListView__LayoutMode int64

const (
	QListView__SinglePass QListView__LayoutMode = QListView__LayoutMode(0)
	QListView__Batched    QListView__LayoutMode = QListView__LayoutMode(1)
)

// QListView::Movement
//
//go:generate stringer -type=QListView__Movement
type QListView__Movement int64

const (
	QListView__Static QListView__Movement = QListView__Movement(0)
	QListView__Free   QListView__Movement = QListView__Movement(1)
	QListView__Snap   QListView__Movement = QListView__Movement(2)
)

// QListView::ResizeMode
//
//go:generate stringer -type=QListView__ResizeMode
type QListView__ResizeMode int64

const (
	QListView__Fixed  QListView__ResizeMode = QListView__ResizeMode(0)
	QListView__Adjust QListView__ResizeMode = QListView__ResizeMode(1)
)

// QListView::ViewMode
//
//go:generate stringer -type=QListView__ViewMode
type QListView__ViewMode int64

const (
	QListView__ListMode QListView__ViewMode = QListView__ViewMode(0)
	QListView__IconMode QListView__ViewMode = QListView__ViewMode(1)
)

// QListWidgetItem::ItemType
//
//go:generate stringer -type=QListWidgetItem__ItemType
type QListWidgetItem__ItemType int64

const (
	QListWidgetItem__Type     QListWidgetItem__ItemType = QListWidgetItem__ItemType(0)
	QListWidgetItem__UserType QListWidgetItem__ItemType = QListWidgetItem__ItemType(1000)
)

type QMainWindow struct {
	QWidget
}

type QMainWindow_ITF interface {
	QWidget_ITF
	QMainWindow_PTR() *QMainWindow
}

func (ptr *QMainWindow) QMainWindow_PTR() *QMainWindow {
	return ptr
}

func (ptr *QMainWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QMainWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQMainWindow(ptr QMainWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindow_PTR().Pointer()
	}
	return nil
}

func NewQMainWindowFromPointer(ptr unsafe.Pointer) (n *QMainWindow) {
	n = new(QMainWindow)
	n.SetPointer(ptr)
	return
}

// QMainWindow::DockOption
//
//go:generate stringer -type=QMainWindow__DockOption
type QMainWindow__DockOption int64

const (
	QMainWindow__AnimatedDocks    QMainWindow__DockOption = QMainWindow__DockOption(0x01)
	QMainWindow__AllowNestedDocks QMainWindow__DockOption = QMainWindow__DockOption(0x02)
	QMainWindow__AllowTabbedDocks QMainWindow__DockOption = QMainWindow__DockOption(0x04)
	QMainWindow__ForceTabbedDocks QMainWindow__DockOption = QMainWindow__DockOption(0x08)
	QMainWindow__VerticalTabs     QMainWindow__DockOption = QMainWindow__DockOption(0x10)
	QMainWindow__GroupedDragging  QMainWindow__DockOption = QMainWindow__DockOption(0x20)
)

func NewQMainWindow(parent QWidget_ITF, flags core.Qt__WindowType) *QMainWindow {
	tmpValue := NewQMainWindowFromPointer(C.QMainWindow_NewQMainWindow(PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QMainWindow) AddToolBar(area core.Qt__ToolBarArea, toolbar QToolBar_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBar(ptr.Pointer(), C.longlong(area), PointerFromQToolBar(toolbar))
	}
}

func (ptr *QMainWindow) AddToolBar2(toolbar QToolBar_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBar2(ptr.Pointer(), PointerFromQToolBar(toolbar))
	}
}

func (ptr *QMainWindow) AddToolBar3(title string) *QToolBar {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		tmpValue := NewQToolBarFromPointer(C.QMainWindow_AddToolBar3(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMainWindow) CentralWidget() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QMainWindow_CentralWidget(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMainWindow) MenuBar() *QMenuBar {
	if ptr.Pointer() != nil {
		tmpValue := NewQMenuBarFromPointer(C.QMainWindow_MenuBar(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMainWindow) SetCentralWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetCentralWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QMainWindow) StatusBar() *QStatusBar {
	if ptr.Pointer() != nil {
		tmpValue := NewQStatusBarFromPointer(C.QMainWindow_StatusBar(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQMainWindow_DestroyQMainWindow
func callbackQMainWindow_DestroyQMainWindow(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QMainWindow"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQMainWindowFromPointer(ptr).DestroyQMainWindowDefault()
	}
}

func (ptr *QMainWindow) ConnectDestroyQMainWindow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QMainWindow"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QMainWindow", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QMainWindow", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMainWindow) DisconnectDestroyQMainWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QMainWindow")
	}
}

func (ptr *QMainWindow) DestroyQMainWindow() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMainWindow_DestroyQMainWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMainWindow) DestroyQMainWindowDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMainWindow_DestroyQMainWindowDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMainWindow) __resizeDocks_docks_atList(i int) *QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQDockWidgetFromPointer(C.QMainWindow___resizeDocks_docks_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMainWindow) __resizeDocks_docks_setList(i QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow___resizeDocks_docks_setList(ptr.Pointer(), PointerFromQDockWidget(i))
	}
}

func (ptr *QMainWindow) __resizeDocks_docks_newList() unsafe.Pointer {
	return C.QMainWindow___resizeDocks_docks_newList(ptr.Pointer())
}

func (ptr *QMainWindow) __resizeDocks_sizes_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QMainWindow___resizeDocks_sizes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QMainWindow) __resizeDocks_sizes_setList(i int) {
	if ptr.Pointer() != nil {
		C.QMainWindow___resizeDocks_sizes_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QMainWindow) __resizeDocks_sizes_newList() unsafe.Pointer {
	return C.QMainWindow___resizeDocks_sizes_newList(ptr.Pointer())
}

func (ptr *QMainWindow) __tabifiedDockWidgets_atList(i int) *QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQDockWidgetFromPointer(C.QMainWindow___tabifiedDockWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMainWindow) __tabifiedDockWidgets_setList(i QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow___tabifiedDockWidgets_setList(ptr.Pointer(), PointerFromQDockWidget(i))
	}
}

func (ptr *QMainWindow) __tabifiedDockWidgets_newList() unsafe.Pointer {
	return C.QMainWindow___tabifiedDockWidgets_newList(ptr.Pointer())
}

// QMdiArea::AreaOption
//
//go:generate stringer -type=QMdiArea__AreaOption
type QMdiArea__AreaOption int64

const (
	QMdiArea__DontMaximizeSubWindowOnActivation QMdiArea__AreaOption = QMdiArea__AreaOption(0x1)
)

// QMdiArea::ViewMode
//
//go:generate stringer -type=QMdiArea__ViewMode
type QMdiArea__ViewMode int64

const (
	QMdiArea__SubWindowView QMdiArea__ViewMode = QMdiArea__ViewMode(0)
	QMdiArea__TabbedView    QMdiArea__ViewMode = QMdiArea__ViewMode(1)
)

// QMdiArea::WindowOrder
//
//go:generate stringer -type=QMdiArea__WindowOrder
type QMdiArea__WindowOrder int64

const (
	QMdiArea__CreationOrder          QMdiArea__WindowOrder = QMdiArea__WindowOrder(0)
	QMdiArea__StackingOrder          QMdiArea__WindowOrder = QMdiArea__WindowOrder(1)
	QMdiArea__ActivationHistoryOrder QMdiArea__WindowOrder = QMdiArea__WindowOrder(2)
)

// QMdiSubWindow::SubWindowOption
//
//go:generate stringer -type=QMdiSubWindow__SubWindowOption
type QMdiSubWindow__SubWindowOption int64

const (
	QMdiSubWindow__AllowOutsideAreaHorizontally QMdiSubWindow__SubWindowOption = QMdiSubWindow__SubWindowOption(0x1)
	QMdiSubWindow__AllowOutsideAreaVertically   QMdiSubWindow__SubWindowOption = QMdiSubWindow__SubWindowOption(0x2)
	QMdiSubWindow__RubberBandResize             QMdiSubWindow__SubWindowOption = QMdiSubWindow__SubWindowOption(0x4)
	QMdiSubWindow__RubberBandMove               QMdiSubWindow__SubWindowOption = QMdiSubWindow__SubWindowOption(0x8)
)

type QMenu struct {
	QWidget
}

type QMenu_ITF interface {
	QWidget_ITF
	QMenu_PTR() *QMenu
}

func (ptr *QMenu) QMenu_PTR() *QMenu {
	return ptr
}

func (ptr *QMenu) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QMenu) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQMenu(ptr QMenu_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMenu_PTR().Pointer()
	}
	return nil
}

func NewQMenuFromPointer(ptr unsafe.Pointer) (n *QMenu) {
	n = new(QMenu)
	n.SetPointer(ptr)
	return
}
func NewQMenu(parent QWidget_ITF) *QMenu {
	tmpValue := NewQMenuFromPointer(C.QMenu_NewQMenu(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQMenu2(title string, parent QWidget_ITF) *QMenu {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	tmpValue := NewQMenuFromPointer(C.QMenu_NewQMenu2(C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QMenu) AddAction(text string) *QAction {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := NewQActionFromPointer(C.QMenu_AddAction(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) AddAction2(icon gui.QIcon_ITF, text string) *QAction {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := NewQActionFromPointer(C.QMenu_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) AddAction3(text string, receiver core.QObject_ITF, member string, shortcut gui.QKeySequence_ITF) *QAction {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		tmpValue := NewQActionFromPointer(C.QMenu_AddAction3(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQObject(receiver), memberC, gui.PointerFromQKeySequence(shortcut)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) AddAction4(icon gui.QIcon_ITF, text string, receiver core.QObject_ITF, member string, shortcut gui.QKeySequence_ITF) *QAction {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		tmpValue := NewQActionFromPointer(C.QMenu_AddAction4(ptr.Pointer(), gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQObject(receiver), memberC, gui.PointerFromQKeySequence(shortcut)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) AddMenu(menu QMenu_ITF) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QMenu_AddMenu(ptr.Pointer(), PointerFromQMenu(menu)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) AddMenu2(title string) *QMenu {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		tmpValue := NewQMenuFromPointer(C.QMenu_AddMenu2(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) AddMenu3(icon gui.QIcon_ITF, title string) *QMenu {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		tmpValue := NewQMenuFromPointer(C.QMenu_AddMenu3(ptr.Pointer(), gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) AddSeparator() *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QMenu_AddSeparator(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) Clear() {
	if ptr.Pointer() != nil {
		C.QMenu_Clear(ptr.Pointer())
	}
}

func (ptr *QMenu) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QMenu_ColumnCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMenu) Exec() *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QMenu_Exec(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) Exec2(p core.QPoint_ITF, action QAction_ITF) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QMenu_Exec2(ptr.Pointer(), core.PointerFromQPoint(p), PointerFromQAction(action)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QMenu_Exec3(actions []*QAction, pos core.QPoint_ITF, at QAction_ITF, parent QWidget_ITF) *QAction {
	tmpValue := NewQActionFromPointer(C.QMenu_QMenu_Exec3(func() unsafe.Pointer {
		tmpList := NewQMenuFromPointer(NewQMenuFromPointer(nil).__exec_actions_newList3())
		for _, v := range actions {
			tmpList.__exec_actions_setList3(v)
		}
		return tmpList.Pointer()
	}(), core.PointerFromQPoint(pos), PointerFromQAction(at), PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QMenu) Exec3(actions []*QAction, pos core.QPoint_ITF, at QAction_ITF, parent QWidget_ITF) *QAction {
	tmpValue := NewQActionFromPointer(C.QMenu_QMenu_Exec3(func() unsafe.Pointer {
		tmpList := NewQMenuFromPointer(NewQMenuFromPointer(nil).__exec_actions_newList3())
		for _, v := range actions {
			tmpList.__exec_actions_setList3(v)
		}
		return tmpList.Pointer()
	}(), core.PointerFromQPoint(pos), PointerFromQAction(at), PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QMenu) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QMenu_Icon(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QMenu_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMenu) Popup(p core.QPoint_ITF, atAction QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QMenu_Popup(ptr.Pointer(), core.PointerFromQPoint(p), PointerFromQAction(atAction))
	}
}

func (ptr *QMenu) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QMenu_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QMenu) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QMenu_Title(ptr.Pointer()))
	}
	return ""
}

//export callbackQMenu_Triggered
func callbackQMenu_Triggered(ptr unsafe.Pointer, action unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "triggered"); signal != nil {
		(*(*func(*QAction))(signal))(NewQActionFromPointer(action))
	}

}

func (ptr *QMenu) ConnectTriggered(f func(action *QAction)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "triggered") {
			C.QMenu_ConnectTriggered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "triggered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "triggered"); signal != nil {
			f := func(action *QAction) {
				(*(*func(*QAction))(signal))(action)
				f(action)
			}
			qt.ConnectSignal(ptr.Pointer(), "triggered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "triggered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMenu) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QMenu_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "triggered")
	}
}

func (ptr *QMenu) Triggered(action QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QMenu_Triggered(ptr.Pointer(), PointerFromQAction(action))
	}
}

//export callbackQMenu_DestroyQMenu
func callbackQMenu_DestroyQMenu(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QMenu"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQMenuFromPointer(ptr).DestroyQMenuDefault()
	}
}

func (ptr *QMenu) ConnectDestroyQMenu(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QMenu"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QMenu", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QMenu", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMenu) DisconnectDestroyQMenu() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QMenu")
	}
}

func (ptr *QMenu) DestroyQMenu() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMenu_DestroyQMenu(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMenu) DestroyQMenuDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMenu_DestroyQMenuDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMenu) __exec_actions_atList3(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QMenu___exec_actions_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenu) __exec_actions_setList3(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QMenu___exec_actions_setList3(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QMenu) __exec_actions_newList3() unsafe.Pointer {
	return C.QMenu___exec_actions_newList3(ptr.Pointer())
}

type QMenuBar struct {
	QWidget
}

type QMenuBar_ITF interface {
	QWidget_ITF
	QMenuBar_PTR() *QMenuBar
}

func (ptr *QMenuBar) QMenuBar_PTR() *QMenuBar {
	return ptr
}

func (ptr *QMenuBar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QMenuBar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQMenuBar(ptr QMenuBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMenuBar_PTR().Pointer()
	}
	return nil
}

func NewQMenuBarFromPointer(ptr unsafe.Pointer) (n *QMenuBar) {
	n = new(QMenuBar)
	n.SetPointer(ptr)
	return
}
func NewQMenuBar(parent QWidget_ITF) *QMenuBar {
	tmpValue := NewQMenuBarFromPointer(C.QMenuBar_NewQMenuBar(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QMenuBar) AddAction(text string) *QAction {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := NewQActionFromPointer(C.QMenuBar_AddAction(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenuBar) AddAction2(text string, receiver core.QObject_ITF, member string) *QAction {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		tmpValue := NewQActionFromPointer(C.QMenuBar_AddAction2(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQObject(receiver), memberC))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenuBar) AddMenu(menu QMenu_ITF) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QMenuBar_AddMenu(ptr.Pointer(), PointerFromQMenu(menu)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenuBar) AddMenu2(title string) *QMenu {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		tmpValue := NewQMenuFromPointer(C.QMenuBar_AddMenu2(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenuBar) AddMenu3(icon gui.QIcon_ITF, title string) *QMenu {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		tmpValue := NewQMenuFromPointer(C.QMenuBar_AddMenu3(ptr.Pointer(), gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenuBar) AddSeparator() *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QMenuBar_AddSeparator(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMenuBar) Clear() {
	if ptr.Pointer() != nil {
		C.QMenuBar_Clear(ptr.Pointer())
	}
}

//export callbackQMenuBar_Triggered
func callbackQMenuBar_Triggered(ptr unsafe.Pointer, action unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "triggered"); signal != nil {
		(*(*func(*QAction))(signal))(NewQActionFromPointer(action))
	}

}

func (ptr *QMenuBar) ConnectTriggered(f func(action *QAction)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "triggered") {
			C.QMenuBar_ConnectTriggered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "triggered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "triggered"); signal != nil {
			f := func(action *QAction) {
				(*(*func(*QAction))(signal))(action)
				f(action)
			}
			qt.ConnectSignal(ptr.Pointer(), "triggered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "triggered", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMenuBar) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QMenuBar_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "triggered")
	}
}

func (ptr *QMenuBar) Triggered(action QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QMenuBar_Triggered(ptr.Pointer(), PointerFromQAction(action))
	}
}

//export callbackQMenuBar_DestroyQMenuBar
func callbackQMenuBar_DestroyQMenuBar(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QMenuBar"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQMenuBarFromPointer(ptr).DestroyQMenuBarDefault()
	}
}

func (ptr *QMenuBar) ConnectDestroyQMenuBar(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QMenuBar"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QMenuBar", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QMenuBar", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMenuBar) DisconnectDestroyQMenuBar() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QMenuBar")
	}
}

func (ptr *QMenuBar) DestroyQMenuBar() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMenuBar_DestroyQMenuBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMenuBar) DestroyQMenuBarDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMenuBar_DestroyQMenuBarDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QMessageBox struct {
	QDialog
}

type QMessageBox_ITF interface {
	QDialog_ITF
	QMessageBox_PTR() *QMessageBox
}

func (ptr *QMessageBox) QMessageBox_PTR() *QMessageBox {
	return ptr
}

func (ptr *QMessageBox) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *QMessageBox) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromQMessageBox(ptr QMessageBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMessageBox_PTR().Pointer()
	}
	return nil
}

func NewQMessageBoxFromPointer(ptr unsafe.Pointer) (n *QMessageBox) {
	n = new(QMessageBox)
	n.SetPointer(ptr)
	return
}

// QMessageBox::ButtonRole
//
//go:generate stringer -type=QMessageBox__ButtonRole
type QMessageBox__ButtonRole int64

const (
	QMessageBox__InvalidRole     QMessageBox__ButtonRole = QMessageBox__ButtonRole(-1)
	QMessageBox__AcceptRole      QMessageBox__ButtonRole = QMessageBox__ButtonRole(0)
	QMessageBox__RejectRole      QMessageBox__ButtonRole = QMessageBox__ButtonRole(1)
	QMessageBox__DestructiveRole QMessageBox__ButtonRole = QMessageBox__ButtonRole(2)
	QMessageBox__ActionRole      QMessageBox__ButtonRole = QMessageBox__ButtonRole(3)
	QMessageBox__HelpRole        QMessageBox__ButtonRole = QMessageBox__ButtonRole(4)
	QMessageBox__YesRole         QMessageBox__ButtonRole = QMessageBox__ButtonRole(5)
	QMessageBox__NoRole          QMessageBox__ButtonRole = QMessageBox__ButtonRole(6)
	QMessageBox__ResetRole       QMessageBox__ButtonRole = QMessageBox__ButtonRole(7)
	QMessageBox__ApplyRole       QMessageBox__ButtonRole = QMessageBox__ButtonRole(8)
	QMessageBox__NRoles          QMessageBox__ButtonRole = QMessageBox__ButtonRole(9)
)

// QMessageBox::Icon
//
//go:generate stringer -type=QMessageBox__Icon
type QMessageBox__Icon int64

const (
	QMessageBox__NoIcon      QMessageBox__Icon = QMessageBox__Icon(0)
	QMessageBox__Information QMessageBox__Icon = QMessageBox__Icon(1)
	QMessageBox__Warning     QMessageBox__Icon = QMessageBox__Icon(2)
	QMessageBox__Critical    QMessageBox__Icon = QMessageBox__Icon(3)
	QMessageBox__Question    QMessageBox__Icon = QMessageBox__Icon(4)
)

// QMessageBox::StandardButton
//
//go:generate stringer -type=QMessageBox__StandardButton
type QMessageBox__StandardButton int64

var (
	QMessageBox__NoButton        QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000000)
	QMessageBox__Ok              QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000400)
	QMessageBox__Save            QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000800)
	QMessageBox__SaveAll         QMessageBox__StandardButton = QMessageBox__StandardButton(0x00001000)
	QMessageBox__Open            QMessageBox__StandardButton = QMessageBox__StandardButton(0x00002000)
	QMessageBox__Yes             QMessageBox__StandardButton = QMessageBox__StandardButton(0x00004000)
	QMessageBox__YesToAll        QMessageBox__StandardButton = QMessageBox__StandardButton(0x00008000)
	QMessageBox__No              QMessageBox__StandardButton = QMessageBox__StandardButton(0x00010000)
	QMessageBox__NoToAll         QMessageBox__StandardButton = QMessageBox__StandardButton(0x00020000)
	QMessageBox__Abort           QMessageBox__StandardButton = QMessageBox__StandardButton(0x00040000)
	QMessageBox__Retry           QMessageBox__StandardButton = QMessageBox__StandardButton(0x00080000)
	QMessageBox__Ignore          QMessageBox__StandardButton = QMessageBox__StandardButton(0x00100000)
	QMessageBox__Close           QMessageBox__StandardButton = QMessageBox__StandardButton(0x00200000)
	QMessageBox__Cancel          QMessageBox__StandardButton = QMessageBox__StandardButton(0x00400000)
	QMessageBox__Discard         QMessageBox__StandardButton = QMessageBox__StandardButton(0x00800000)
	QMessageBox__Help            QMessageBox__StandardButton = QMessageBox__StandardButton(0x01000000)
	QMessageBox__Apply           QMessageBox__StandardButton = QMessageBox__StandardButton(0x02000000)
	QMessageBox__Reset           QMessageBox__StandardButton = QMessageBox__StandardButton(0x04000000)
	QMessageBox__RestoreDefaults QMessageBox__StandardButton = QMessageBox__StandardButton(0x08000000)
	QMessageBox__FirstButton     QMessageBox__StandardButton = QMessageBox__StandardButton(QMessageBox__Ok)
	QMessageBox__LastButton      QMessageBox__StandardButton = QMessageBox__StandardButton(QMessageBox__RestoreDefaults)
	QMessageBox__YesAll          QMessageBox__StandardButton = QMessageBox__StandardButton(QMessageBox__YesToAll)
	QMessageBox__NoAll           QMessageBox__StandardButton = QMessageBox__StandardButton(QMessageBox__NoToAll)
	QMessageBox__Default         QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000100)
	QMessageBox__Escape          QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000200)
	QMessageBox__FlagMask        QMessageBox__StandardButton = QMessageBox__StandardButton(0x00000300)
	QMessageBox__ButtonMask      QMessageBox__StandardButton = QMessageBox__StandardButton(C.QMessageBox_ButtonMask_Type())
)

func NewQMessageBox(parent QWidget_ITF) *QMessageBox {
	tmpValue := NewQMessageBoxFromPointer(C.QMessageBox_NewQMessageBox(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQMessageBox2(icon QMessageBox__Icon, title string, text string, buttons QMessageBox__StandardButton, parent QWidget_ITF, ff core.Qt__WindowType) *QMessageBox {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQMessageBoxFromPointer(C.QMessageBox_NewQMessageBox2(C.longlong(icon), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(buttons), PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QMessageBox_About(parent QWidget_ITF, title string, text string) {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	C.QMessageBox_QMessageBox_About(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
}

func (ptr *QMessageBox) About(parent QWidget_ITF, title string, text string) {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	C.QMessageBox_QMessageBox_About(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
}

func (ptr *QMessageBox) Button(which QMessageBox__StandardButton) *QAbstractButton {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractButtonFromPointer(C.QMessageBox_Button(ptr.Pointer(), C.longlong(which)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMessageBox) Buttons() []*QAbstractButton {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWidgets_PackedList) []*QAbstractButton {
			out := make([]*QAbstractButton, int(l.len))
			tmpList := NewQMessageBoxFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__buttons_atList(i)
			}
			return out
		}(C.QMessageBox_Buttons(ptr.Pointer()))
	}
	return make([]*QAbstractButton, 0)
}

func (ptr *QMessageBox) Icon() QMessageBox__Icon {
	if ptr.Pointer() != nil {
		return QMessageBox__Icon(C.QMessageBox_Icon(ptr.Pointer()))
	}
	return 0
}

func QMessageBox_Information(parent QWidget_ITF, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	return QMessageBox__StandardButton(C.QMessageBox_QMessageBox_Information(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(buttons), C.longlong(defaultButton)))
}

func (ptr *QMessageBox) Information(parent QWidget_ITF, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	return QMessageBox__StandardButton(C.QMessageBox_QMessageBox_Information(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(buttons), C.longlong(defaultButton)))
}

func (ptr *QMessageBox) Open(receiver core.QObject_ITF, member string) {
	if ptr.Pointer() != nil {
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		C.QMessageBox_Open(ptr.Pointer(), core.PointerFromQObject(receiver), memberC)
	}
}

func QMessageBox_Question(parent QWidget_ITF, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	return QMessageBox__StandardButton(C.QMessageBox_QMessageBox_Question(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(buttons), C.longlong(defaultButton)))
}

func (ptr *QMessageBox) Question(parent QWidget_ITF, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	return QMessageBox__StandardButton(C.QMessageBox_QMessageBox_Question(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(buttons), C.longlong(defaultButton)))
}

func (ptr *QMessageBox) SetIcon(vqm QMessageBox__Icon) {
	if ptr.Pointer() != nil {
		C.QMessageBox_SetIcon(ptr.Pointer(), C.longlong(vqm))
	}
}

func (ptr *QMessageBox) SetStandardButtons(buttons QMessageBox__StandardButton) {
	if ptr.Pointer() != nil {
		C.QMessageBox_SetStandardButtons(ptr.Pointer(), C.longlong(buttons))
	}
}

func (ptr *QMessageBox) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QMessageBox_SetText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QMessageBox) SetWindowTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QMessageBox_SetWindowTitle(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *QMessageBox) StandardButton(button QAbstractButton_ITF) QMessageBox__StandardButton {
	if ptr.Pointer() != nil {
		return QMessageBox__StandardButton(C.QMessageBox_StandardButton(ptr.Pointer(), PointerFromQAbstractButton(button)))
	}
	return 0
}

func (ptr *QMessageBox) StandardButtons() QMessageBox__StandardButton {
	if ptr.Pointer() != nil {
		return QMessageBox__StandardButton(C.QMessageBox_StandardButtons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMessageBox) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QMessageBox_Text(ptr.Pointer()))
	}
	return ""
}

func QMessageBox_Warning(parent QWidget_ITF, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	return QMessageBox__StandardButton(C.QMessageBox_QMessageBox_Warning(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(buttons), C.longlong(defaultButton)))
}

func (ptr *QMessageBox) Warning(parent QWidget_ITF, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	return QMessageBox__StandardButton(C.QMessageBox_QMessageBox_Warning(PointerFromQWidget(parent), C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(buttons), C.longlong(defaultButton)))
}

//export callbackQMessageBox_DestroyQMessageBox
func callbackQMessageBox_DestroyQMessageBox(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QMessageBox"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQMessageBoxFromPointer(ptr).DestroyQMessageBoxDefault()
	}
}

func (ptr *QMessageBox) ConnectDestroyQMessageBox(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QMessageBox"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QMessageBox", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QMessageBox", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMessageBox) DisconnectDestroyQMessageBox() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QMessageBox")
	}
}

func (ptr *QMessageBox) DestroyQMessageBox() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMessageBox_DestroyQMessageBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMessageBox) DestroyQMessageBoxDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMessageBox_DestroyQMessageBoxDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMessageBox) __buttons_atList(i int) *QAbstractButton {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractButtonFromPointer(C.QMessageBox___buttons_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMessageBox) __buttons_setList(i QAbstractButton_ITF) {
	if ptr.Pointer() != nil {
		C.QMessageBox___buttons_setList(ptr.Pointer(), PointerFromQAbstractButton(i))
	}
}

func (ptr *QMessageBox) __buttons_newList() unsafe.Pointer {
	return C.QMessageBox___buttons_newList(ptr.Pointer())
}

// QOpenGLWidget::UpdateBehavior
//
//go:generate stringer -type=QOpenGLWidget__UpdateBehavior
type QOpenGLWidget__UpdateBehavior int64

const (
	QOpenGLWidget__NoPartialUpdate QOpenGLWidget__UpdateBehavior = QOpenGLWidget__UpdateBehavior(0)
	QOpenGLWidget__PartialUpdate   QOpenGLWidget__UpdateBehavior = QOpenGLWidget__UpdateBehavior(1)
)

// QPinchGesture::ChangeFlag
//
//go:generate stringer -type=QPinchGesture__ChangeFlag
type QPinchGesture__ChangeFlag int64

const (
	QPinchGesture__ScaleFactorChanged   QPinchGesture__ChangeFlag = QPinchGesture__ChangeFlag(0x1)
	QPinchGesture__RotationAngleChanged QPinchGesture__ChangeFlag = QPinchGesture__ChangeFlag(0x2)
	QPinchGesture__CenterPointChanged   QPinchGesture__ChangeFlag = QPinchGesture__ChangeFlag(0x4)
)

// QPlainTextEdit::LineWrapMode
//
//go:generate stringer -type=QPlainTextEdit__LineWrapMode
type QPlainTextEdit__LineWrapMode int64

const (
	QPlainTextEdit__NoWrap      QPlainTextEdit__LineWrapMode = QPlainTextEdit__LineWrapMode(0)
	QPlainTextEdit__WidgetWidth QPlainTextEdit__LineWrapMode = QPlainTextEdit__LineWrapMode(1)
)

// QProgressBar::Direction
//
//go:generate stringer -type=QProgressBar__Direction
type QProgressBar__Direction int64

const (
	QProgressBar__TopToBottom QProgressBar__Direction = QProgressBar__Direction(0)
	QProgressBar__BottomToTop QProgressBar__Direction = QProgressBar__Direction(1)
)

type QPushButton struct {
	QAbstractButton
}

type QPushButton_ITF interface {
	QAbstractButton_ITF
	QPushButton_PTR() *QPushButton
}

func (ptr *QPushButton) QPushButton_PTR() *QPushButton {
	return ptr
}

func (ptr *QPushButton) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractButton_PTR().Pointer()
	}
	return nil
}

func (ptr *QPushButton) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractButton_PTR().SetPointer(p)
	}
}

func PointerFromQPushButton(ptr QPushButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPushButton_PTR().Pointer()
	}
	return nil
}

func NewQPushButtonFromPointer(ptr unsafe.Pointer) (n *QPushButton) {
	n = new(QPushButton)
	n.SetPointer(ptr)
	return
}
func NewQPushButton(parent QWidget_ITF) *QPushButton {
	tmpValue := NewQPushButtonFromPointer(C.QPushButton_NewQPushButton(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQPushButton2(text string, parent QWidget_ITF) *QPushButton {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQPushButtonFromPointer(C.QPushButton_NewQPushButton2(C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQPushButton3(icon gui.QIcon_ITF, text string, parent QWidget_ITF) *QPushButton {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQPushButtonFromPointer(C.QPushButton_NewQPushButton3(gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QPushButton) Menu() *QMenu {
	if ptr.Pointer() != nil {
		tmpValue := NewQMenuFromPointer(C.QPushButton_Menu(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQPushButton_DestroyQPushButton
func callbackQPushButton_DestroyQPushButton(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPushButton"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPushButtonFromPointer(ptr).DestroyQPushButtonDefault()
	}
}

func (ptr *QPushButton) ConnectDestroyQPushButton(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPushButton"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPushButton", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPushButton", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPushButton) DisconnectDestroyQPushButton() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPushButton")
	}
}

func (ptr *QPushButton) DestroyQPushButton() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPushButton_DestroyQPushButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPushButton) DestroyQPushButtonDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPushButton_DestroyQPushButtonDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQPushButton_PaintEvent
func callbackQPushButton_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*gui.QPaintEvent))(signal))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQPushButtonFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QPushButton) PaintEvent(e gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPushButton_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QPushButton) PaintEventDefault(e gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QPushButton_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

// QRubberBand::Shape
//
//go:generate stringer -type=QRubberBand__Shape
type QRubberBand__Shape int64

const (
	QRubberBand__Line      QRubberBand__Shape = QRubberBand__Shape(0)
	QRubberBand__Rectangle QRubberBand__Shape = QRubberBand__Shape(1)
)

type QScrollArea struct {
	QAbstractScrollArea
}

type QScrollArea_ITF interface {
	QAbstractScrollArea_ITF
	QScrollArea_PTR() *QScrollArea
}

func (ptr *QScrollArea) QScrollArea_PTR() *QScrollArea {
	return ptr
}

func (ptr *QScrollArea) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func (ptr *QScrollArea) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractScrollArea_PTR().SetPointer(p)
	}
}

func PointerFromQScrollArea(ptr QScrollArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollArea_PTR().Pointer()
	}
	return nil
}

func NewQScrollAreaFromPointer(ptr unsafe.Pointer) (n *QScrollArea) {
	n = new(QScrollArea)
	n.SetPointer(ptr)
	return
}
func NewQScrollArea(parent QWidget_ITF) *QScrollArea {
	tmpValue := NewQScrollAreaFromPointer(C.QScrollArea_NewQScrollArea(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QScrollArea) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QScrollArea_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScrollArea) SetAlignment(vqt core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QScrollArea_SetAlignment(ptr.Pointer(), C.longlong(vqt))
	}
}

func (ptr *QScrollArea) SetWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QScrollArea_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QScrollArea) Widget() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QScrollArea_Widget(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScrollArea_DestroyQScrollArea
func callbackQScrollArea_DestroyQScrollArea(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScrollArea"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScrollAreaFromPointer(ptr).DestroyQScrollAreaDefault()
	}
}

func (ptr *QScrollArea) ConnectDestroyQScrollArea(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScrollArea"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QScrollArea", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScrollArea", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScrollArea) DisconnectDestroyQScrollArea() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QScrollArea")
	}
}

func (ptr *QScrollArea) DestroyQScrollArea() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScrollArea_DestroyQScrollArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScrollArea) DestroyQScrollAreaDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScrollArea_DestroyQScrollAreaDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScrollBar struct {
	QAbstractSlider
}

type QScrollBar_ITF interface {
	QAbstractSlider_ITF
	QScrollBar_PTR() *QScrollBar
}

func (ptr *QScrollBar) QScrollBar_PTR() *QScrollBar {
	return ptr
}

func (ptr *QScrollBar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSlider_PTR().Pointer()
	}
	return nil
}

func (ptr *QScrollBar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSlider_PTR().SetPointer(p)
	}
}

func PointerFromQScrollBar(ptr QScrollBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollBar_PTR().Pointer()
	}
	return nil
}

func NewQScrollBarFromPointer(ptr unsafe.Pointer) (n *QScrollBar) {
	n = new(QScrollBar)
	n.SetPointer(ptr)
	return
}
func NewQScrollBar(parent QWidget_ITF) *QScrollBar {
	tmpValue := NewQScrollBarFromPointer(C.QScrollBar_NewQScrollBar(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQScrollBar2(orientation core.Qt__Orientation, parent QWidget_ITF) *QScrollBar {
	tmpValue := NewQScrollBarFromPointer(C.QScrollBar_NewQScrollBar2(C.longlong(orientation), PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQScrollBar_DestroyQScrollBar
func callbackQScrollBar_DestroyQScrollBar(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScrollBar"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScrollBarFromPointer(ptr).DestroyQScrollBarDefault()
	}
}

func (ptr *QScrollBar) ConnectDestroyQScrollBar(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScrollBar"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QScrollBar", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScrollBar", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScrollBar) DisconnectDestroyQScrollBar() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QScrollBar")
	}
}

func (ptr *QScrollBar) DestroyQScrollBar() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScrollBar_DestroyQScrollBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScrollBar) DestroyQScrollBarDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScrollBar_DestroyQScrollBarDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QScroller::Input
//
//go:generate stringer -type=QScroller__Input
type QScroller__Input int64

const (
	QScroller__InputPress   QScroller__Input = QScroller__Input(1)
	QScroller__InputMove    QScroller__Input = QScroller__Input(2)
	QScroller__InputRelease QScroller__Input = QScroller__Input(3)
)

// QScroller::ScrollerGestureType
//
//go:generate stringer -type=QScroller__ScrollerGestureType
type QScroller__ScrollerGestureType int64

const (
	QScroller__TouchGesture             QScroller__ScrollerGestureType = QScroller__ScrollerGestureType(0)
	QScroller__LeftMouseButtonGesture   QScroller__ScrollerGestureType = QScroller__ScrollerGestureType(1)
	QScroller__RightMouseButtonGesture  QScroller__ScrollerGestureType = QScroller__ScrollerGestureType(2)
	QScroller__MiddleMouseButtonGesture QScroller__ScrollerGestureType = QScroller__ScrollerGestureType(3)
)

// QScroller::State
//
//go:generate stringer -type=QScroller__State
type QScroller__State int64

const (
	QScroller__Inactive  QScroller__State = QScroller__State(0)
	QScroller__Pressed   QScroller__State = QScroller__State(1)
	QScroller__Dragging  QScroller__State = QScroller__State(2)
	QScroller__Scrolling QScroller__State = QScroller__State(3)
)

// QScrollerProperties::FrameRates
//
//go:generate stringer -type=QScrollerProperties__FrameRates
type QScrollerProperties__FrameRates int64

const (
	QScrollerProperties__Standard QScrollerProperties__FrameRates = QScrollerProperties__FrameRates(0)
	QScrollerProperties__Fps60    QScrollerProperties__FrameRates = QScrollerProperties__FrameRates(1)
	QScrollerProperties__Fps30    QScrollerProperties__FrameRates = QScrollerProperties__FrameRates(2)
	QScrollerProperties__Fps20    QScrollerProperties__FrameRates = QScrollerProperties__FrameRates(3)
)

// QScrollerProperties::OvershootPolicy
//
//go:generate stringer -type=QScrollerProperties__OvershootPolicy
type QScrollerProperties__OvershootPolicy int64

const (
	QScrollerProperties__OvershootWhenScrollable QScrollerProperties__OvershootPolicy = QScrollerProperties__OvershootPolicy(0)
	QScrollerProperties__OvershootAlwaysOff      QScrollerProperties__OvershootPolicy = QScrollerProperties__OvershootPolicy(1)
	QScrollerProperties__OvershootAlwaysOn       QScrollerProperties__OvershootPolicy = QScrollerProperties__OvershootPolicy(2)
)

// QScrollerProperties::ScrollMetric
//
//go:generate stringer -type=QScrollerProperties__ScrollMetric
type QScrollerProperties__ScrollMetric int64

const (
	QScrollerProperties__MousePressEventDelay           QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(0)
	QScrollerProperties__DragStartDistance              QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(1)
	QScrollerProperties__DragVelocitySmoothingFactor    QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(2)
	QScrollerProperties__AxisLockThreshold              QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(3)
	QScrollerProperties__ScrollingCurve                 QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(4)
	QScrollerProperties__DecelerationFactor             QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(5)
	QScrollerProperties__MinimumVelocity                QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(6)
	QScrollerProperties__MaximumVelocity                QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(7)
	QScrollerProperties__MaximumClickThroughVelocity    QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(8)
	QScrollerProperties__AcceleratingFlickMaximumTime   QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(9)
	QScrollerProperties__AcceleratingFlickSpeedupFactor QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(10)
	QScrollerProperties__SnapPositionRatio              QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(11)
	QScrollerProperties__SnapTime                       QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(12)
	QScrollerProperties__OvershootDragResistanceFactor  QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(13)
	QScrollerProperties__OvershootDragDistanceFactor    QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(14)
	QScrollerProperties__OvershootScrollDistanceFactor  QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(15)
	QScrollerProperties__OvershootScrollTime            QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(16)
	QScrollerProperties__HorizontalOvershootPolicy      QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(17)
	QScrollerProperties__VerticalOvershootPolicy        QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(18)
	QScrollerProperties__FrameRate                      QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(19)
	QScrollerProperties__ScrollMetricCount              QScrollerProperties__ScrollMetric = QScrollerProperties__ScrollMetric(20)
)

type QSizePolicy struct {
	ptr unsafe.Pointer
}

type QSizePolicy_ITF interface {
	QSizePolicy_PTR() *QSizePolicy
}

func (ptr *QSizePolicy) QSizePolicy_PTR() *QSizePolicy {
	return ptr
}

func (ptr *QSizePolicy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSizePolicy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSizePolicy(ptr QSizePolicy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSizePolicy_PTR().Pointer()
	}
	return nil
}

func NewQSizePolicyFromPointer(ptr unsafe.Pointer) (n *QSizePolicy) {
	n = new(QSizePolicy)
	n.SetPointer(ptr)
	return
}
func (ptr *QSizePolicy) DestroyQSizePolicy() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QSizePolicy__ControlType
//QSizePolicy::ControlType
type QSizePolicy__ControlType int64

const (
	QSizePolicy__DefaultType QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000001)
	QSizePolicy__ButtonBox   QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000002)
	QSizePolicy__CheckBox    QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000004)
	QSizePolicy__ComboBox    QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000008)
	QSizePolicy__Frame       QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000010)
	QSizePolicy__GroupBox    QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000020)
	QSizePolicy__Label       QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000040)
	QSizePolicy__Line        QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000080)
	QSizePolicy__LineEdit    QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000100)
	QSizePolicy__PushButton  QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000200)
	QSizePolicy__RadioButton QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000400)
	QSizePolicy__Slider      QSizePolicy__ControlType = QSizePolicy__ControlType(0x00000800)
	QSizePolicy__SpinBox     QSizePolicy__ControlType = QSizePolicy__ControlType(0x00001000)
	QSizePolicy__TabWidget   QSizePolicy__ControlType = QSizePolicy__ControlType(0x00002000)
	QSizePolicy__ToolButton  QSizePolicy__ControlType = QSizePolicy__ControlType(0x00004000)
)

// QSizePolicy::Policy
//
//go:generate stringer -type=QSizePolicy__Policy
type QSizePolicy__Policy int64

const (
	QSizePolicy__Fixed            QSizePolicy__Policy = QSizePolicy__Policy(0)
	QSizePolicy__Minimum          QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__GrowFlag)
	QSizePolicy__Maximum          QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__ShrinkFlag)
	QSizePolicy__Preferred        QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__GrowFlag | QSizePolicy__ShrinkFlag)
	QSizePolicy__MinimumExpanding QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__GrowFlag | QSizePolicy__ExpandFlag)
	QSizePolicy__Expanding        QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__GrowFlag | QSizePolicy__ShrinkFlag | QSizePolicy__ExpandFlag)
	QSizePolicy__Ignored          QSizePolicy__Policy = QSizePolicy__Policy(QSizePolicy__ShrinkFlag | QSizePolicy__GrowFlag | QSizePolicy__IgnoreFlag)
)

// QSizePolicy::PolicyFlag
//
//go:generate stringer -type=QSizePolicy__PolicyFlag
type QSizePolicy__PolicyFlag int64

const (
	QSizePolicy__GrowFlag   QSizePolicy__PolicyFlag = QSizePolicy__PolicyFlag(1)
	QSizePolicy__ExpandFlag QSizePolicy__PolicyFlag = QSizePolicy__PolicyFlag(2)
	QSizePolicy__ShrinkFlag QSizePolicy__PolicyFlag = QSizePolicy__PolicyFlag(4)
	QSizePolicy__IgnoreFlag QSizePolicy__PolicyFlag = QSizePolicy__PolicyFlag(8)
)

func NewQSizePolicy() *QSizePolicy {
	tmpValue := NewQSizePolicyFromPointer(C.QSizePolicy_NewQSizePolicy())
	qt.SetFinalizer(tmpValue, (*QSizePolicy).DestroyQSizePolicy)
	return tmpValue
}

func NewQSizePolicy2(horizontal QSizePolicy__Policy, vertical QSizePolicy__Policy, ty QSizePolicy__ControlType) *QSizePolicy {
	tmpValue := NewQSizePolicyFromPointer(C.QSizePolicy_NewQSizePolicy2(C.longlong(horizontal), C.longlong(vertical), C.longlong(ty)))
	qt.SetFinalizer(tmpValue, (*QSizePolicy).DestroyQSizePolicy)
	return tmpValue
}

func (ptr *QSizePolicy) ControlType() QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QSizePolicy_ControlType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizePolicy) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSizePolicy_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizePolicy) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSizePolicy_HasHeightForWidth(ptr.Pointer())) != 0
	}
	return false
}

// QSlider::TickPosition
//
//go:generate stringer -type=QSlider__TickPosition
type QSlider__TickPosition int64

const (
	QSlider__NoTicks        QSlider__TickPosition = QSlider__TickPosition(0)
	QSlider__TicksAbove     QSlider__TickPosition = QSlider__TickPosition(1)
	QSlider__TicksLeft      QSlider__TickPosition = QSlider__TickPosition(QSlider__TicksAbove)
	QSlider__TicksBelow     QSlider__TickPosition = QSlider__TickPosition(2)
	QSlider__TicksRight     QSlider__TickPosition = QSlider__TickPosition(QSlider__TicksBelow)
	QSlider__TicksBothSides QSlider__TickPosition = QSlider__TickPosition(3)
)

type QSpacerItem struct {
	QLayoutItem
}

type QSpacerItem_ITF interface {
	QLayoutItem_ITF
	QSpacerItem_PTR() *QSpacerItem
}

func (ptr *QSpacerItem) QSpacerItem_PTR() *QSpacerItem {
	return ptr
}

func (ptr *QSpacerItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayoutItem_PTR().Pointer()
	}
	return nil
}

func (ptr *QSpacerItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QLayoutItem_PTR().SetPointer(p)
	}
}

func PointerFromQSpacerItem(ptr QSpacerItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSpacerItem_PTR().Pointer()
	}
	return nil
}

func NewQSpacerItemFromPointer(ptr unsafe.Pointer) (n *QSpacerItem) {
	n = new(QSpacerItem)
	n.SetPointer(ptr)
	return
}
func NewQSpacerItem(w int, h int, hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) *QSpacerItem {
	return NewQSpacerItemFromPointer(C.QSpacerItem_NewQSpacerItem(C.int(int32(w)), C.int(int32(h)), C.longlong(hPolicy), C.longlong(vPolicy)))
}

//export callbackQSpacerItem_ExpandingDirections
func callbackQSpacerItem_ExpandingDirections(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expandingDirections"); signal != nil {
		return C.longlong((*(*func() core.Qt__Orientation)(signal))())
	}

	return C.longlong(NewQSpacerItemFromPointer(ptr).ExpandingDirectionsDefault())
}

func (ptr *QSpacerItem) ConnectExpandingDirections(f func() core.Qt__Orientation) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "expandingDirections"); signal != nil {
			f := func() core.Qt__Orientation {
				(*(*func() core.Qt__Orientation)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "expandingDirections", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "expandingDirections", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSpacerItem) DisconnectExpandingDirections() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "expandingDirections")
	}
}

func (ptr *QSpacerItem) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSpacerItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSpacerItem) ExpandingDirectionsDefault() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSpacerItem_ExpandingDirectionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSpacerItem_Geometry
func callbackQSpacerItem_Geometry(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "geometry"); signal != nil {
		return core.PointerFromQRect((*(*func() *core.QRect)(signal))())
	}

	return core.PointerFromQRect(NewQSpacerItemFromPointer(ptr).GeometryDefault())
}

func (ptr *QSpacerItem) ConnectGeometry(f func() *core.QRect) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "geometry"); signal != nil {
			f := func() *core.QRect {
				(*(*func() *core.QRect)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "geometry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "geometry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSpacerItem) DisconnectGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "geometry")
	}
}

func (ptr *QSpacerItem) Geometry() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QSpacerItem_Geometry(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QSpacerItem) GeometryDefault() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QSpacerItem_GeometryDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQSpacerItem_IsEmpty
func callbackQSpacerItem_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSpacerItemFromPointer(ptr).IsEmptyDefault())))
}

func (ptr *QSpacerItem) ConnectIsEmpty(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isEmpty"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "isEmpty", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isEmpty", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSpacerItem) DisconnectIsEmpty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isEmpty")
	}
}

func (ptr *QSpacerItem) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSpacerItem_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSpacerItem) IsEmptyDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSpacerItem_IsEmptyDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSpacerItem_MaximumSize
func callbackQSpacerItem_MaximumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "maximumSize"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQSpacerItemFromPointer(ptr).MaximumSizeDefault())
}

func (ptr *QSpacerItem) ConnectMaximumSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "maximumSize"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "maximumSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maximumSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSpacerItem) DisconnectMaximumSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "maximumSize")
	}
}

func (ptr *QSpacerItem) MaximumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSpacerItem_MaximumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSpacerItem) MaximumSizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSpacerItem_MaximumSizeDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSpacerItem_MinimumSize
func callbackQSpacerItem_MinimumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSize"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQSpacerItemFromPointer(ptr).MinimumSizeDefault())
}

func (ptr *QSpacerItem) ConnectMinimumSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "minimumSize"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSpacerItem) DisconnectMinimumSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "minimumSize")
	}
}

func (ptr *QSpacerItem) MinimumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSpacerItem_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSpacerItem) MinimumSizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSpacerItem_MinimumSizeDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSpacerItem_SetGeometry
func callbackQSpacerItem_SetGeometry(ptr unsafe.Pointer, r unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		(*(*func(*core.QRect))(signal))(core.NewQRectFromPointer(r))
	} else {
		NewQSpacerItemFromPointer(ptr).SetGeometryDefault(core.NewQRectFromPointer(r))
	}
}

func (ptr *QSpacerItem) ConnectSetGeometry(f func(r *core.QRect)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGeometry"); signal != nil {
			f := func(r *core.QRect) {
				(*(*func(*core.QRect))(signal))(r)
				f(r)
			}
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSpacerItem) DisconnectSetGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGeometry")
	}
}

func (ptr *QSpacerItem) SetGeometry(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QSpacerItem_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QSpacerItem) SetGeometryDefault(r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QSpacerItem_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

//export callbackQSpacerItem_SizeHint
func callbackQSpacerItem_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQSpacerItemFromPointer(ptr).SizeHintDefault())
}

func (ptr *QSpacerItem) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sizeHint"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSpacerItem) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sizeHint")
	}
}

func (ptr *QSpacerItem) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSpacerItem_SizeHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSpacerItem) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSpacerItem_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSpacerItem) SizePolicy() *QSizePolicy {
	if ptr.Pointer() != nil {
		tmpValue := NewQSizePolicyFromPointer(C.QSpacerItem_SizePolicy(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSizePolicy).DestroyQSizePolicy)
		return tmpValue
	}
	return nil
}

//export callbackQSpacerItem_DestroyQSpacerItem
func callbackQSpacerItem_DestroyQSpacerItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSpacerItem"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSpacerItemFromPointer(ptr).DestroyQSpacerItemDefault()
	}
}

func (ptr *QSpacerItem) ConnectDestroyQSpacerItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSpacerItem"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSpacerItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSpacerItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSpacerItem) DisconnectDestroyQSpacerItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSpacerItem")
	}
}

func (ptr *QSpacerItem) DestroyQSpacerItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSpacerItem_DestroyQSpacerItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSpacerItem) DestroyQSpacerItemDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSpacerItem_DestroyQSpacerItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSplitter struct {
	QFrame
}

type QSplitter_ITF interface {
	QFrame_ITF
	QSplitter_PTR() *QSplitter
}

func (ptr *QSplitter) QSplitter_PTR() *QSplitter {
	return ptr
}

func (ptr *QSplitter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func (ptr *QSplitter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QFrame_PTR().SetPointer(p)
	}
}

func PointerFromQSplitter(ptr QSplitter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplitter_PTR().Pointer()
	}
	return nil
}

func NewQSplitterFromPointer(ptr unsafe.Pointer) (n *QSplitter) {
	n = new(QSplitter)
	n.SetPointer(ptr)
	return
}
func NewQSplitter(parent QWidget_ITF) *QSplitter {
	tmpValue := NewQSplitterFromPointer(C.QSplitter_NewQSplitter(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSplitter2(orientation core.Qt__Orientation, parent QWidget_ITF) *QSplitter {
	tmpValue := NewQSplitterFromPointer(C.QSplitter_NewQSplitter2(C.longlong(orientation), PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSplitter) AddWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QSplitter_AddWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QSplitter) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSplitter_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSplitter) IndexOf(widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSplitter_IndexOf(ptr.Pointer(), PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QSplitter) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSplitter_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitter) Refresh() {
	if ptr.Pointer() != nil {
		C.QSplitter_Refresh(ptr.Pointer())
	}
}

func (ptr *QSplitter) SetOrientation(vqt core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QSplitter_SetOrientation(ptr.Pointer(), C.longlong(vqt))
	}
}

func (ptr *QSplitter) SetStretchFactor(index int, stretch int) {
	if ptr.Pointer() != nil {
		C.QSplitter_SetStretchFactor(ptr.Pointer(), C.int(int32(index)), C.int(int32(stretch)))
	}
}

func (ptr *QSplitter) Widget(index int) *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QSplitter_Widget(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSplitter_DestroyQSplitter
func callbackQSplitter_DestroyQSplitter(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSplitter"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSplitterFromPointer(ptr).DestroyQSplitterDefault()
	}
}

func (ptr *QSplitter) ConnectDestroyQSplitter(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSplitter"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSplitter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSplitter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSplitter) DisconnectDestroyQSplitter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSplitter")
	}
}

func (ptr *QSplitter) DestroyQSplitter() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSplitter_DestroyQSplitter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSplitter) DestroyQSplitterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSplitter_DestroyQSplitterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSplitter) __setSizes_list_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSplitter___setSizes_list_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSplitter) __setSizes_list_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSplitter___setSizes_list_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSplitter) __setSizes_list_newList() unsafe.Pointer {
	return C.QSplitter___setSizes_list_newList(ptr.Pointer())
}

func (ptr *QSplitter) __sizes_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSplitter___sizes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSplitter) __sizes_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSplitter___sizes_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSplitter) __sizes_newList() unsafe.Pointer {
	return C.QSplitter___sizes_newList(ptr.Pointer())
}

// QStackedLayout::StackingMode
//
//go:generate stringer -type=QStackedLayout__StackingMode
type QStackedLayout__StackingMode int64

const (
	QStackedLayout__StackOne QStackedLayout__StackingMode = QStackedLayout__StackingMode(0)
	QStackedLayout__StackAll QStackedLayout__StackingMode = QStackedLayout__StackingMode(1)
)

type QStatusBar struct {
	QWidget
}

type QStatusBar_ITF interface {
	QWidget_ITF
	QStatusBar_PTR() *QStatusBar
}

func (ptr *QStatusBar) QStatusBar_PTR() *QStatusBar {
	return ptr
}

func (ptr *QStatusBar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QStatusBar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQStatusBar(ptr QStatusBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStatusBar_PTR().Pointer()
	}
	return nil
}

func NewQStatusBarFromPointer(ptr unsafe.Pointer) (n *QStatusBar) {
	n = new(QStatusBar)
	n.SetPointer(ptr)
	return
}
func NewQStatusBar(parent QWidget_ITF) *QStatusBar {
	tmpValue := NewQStatusBarFromPointer(C.QStatusBar_NewQStatusBar(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QStatusBar) AddWidget(widget QWidget_ITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QStatusBar_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(int32(stretch)))
	}
}

//export callbackQStatusBar_ShowMessage
func callbackQStatusBar_ShowMessage(ptr unsafe.Pointer, message C.struct_QtWidgets_PackedString, timeout C.int) {
	if signal := qt.GetSignal(ptr, "showMessage"); signal != nil {
		(*(*func(string, int))(signal))(cGoUnpackString(message), int(int32(timeout)))
	} else {
		NewQStatusBarFromPointer(ptr).ShowMessageDefault(cGoUnpackString(message), int(int32(timeout)))
	}
}

func (ptr *QStatusBar) ConnectShowMessage(f func(message string, timeout int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "showMessage"); signal != nil {
			f := func(message string, timeout int) {
				(*(*func(string, int))(signal))(message, timeout)
				f(message, timeout)
			}
			qt.ConnectSignal(ptr.Pointer(), "showMessage", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "showMessage", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStatusBar) DisconnectShowMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "showMessage")
	}
}

func (ptr *QStatusBar) ShowMessage(message string, timeout int) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.QStatusBar_ShowMessage(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: messageC, len: C.longlong(len(message))}, C.int(int32(timeout)))
	}
}

func (ptr *QStatusBar) ShowMessageDefault(message string, timeout int) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.QStatusBar_ShowMessageDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: messageC, len: C.longlong(len(message))}, C.int(int32(timeout)))
	}
}

//export callbackQStatusBar_DestroyQStatusBar
func callbackQStatusBar_DestroyQStatusBar(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QStatusBar"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQStatusBarFromPointer(ptr).DestroyQStatusBarDefault()
	}
}

func (ptr *QStatusBar) ConnectDestroyQStatusBar(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QStatusBar"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QStatusBar", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QStatusBar", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStatusBar) DisconnectDestroyQStatusBar() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QStatusBar")
	}
}

func (ptr *QStatusBar) DestroyQStatusBar() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStatusBar_DestroyQStatusBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStatusBar) DestroyQStatusBarDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStatusBar_DestroyQStatusBarDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QStyle struct {
	core.QObject
}

type QStyle_ITF interface {
	core.QObject_ITF
	QStyle_PTR() *QStyle
}

func (ptr *QStyle) QStyle_PTR() *QStyle {
	return ptr
}

func (ptr *QStyle) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QStyle) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQStyle(ptr QStyle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyle_PTR().Pointer()
	}
	return nil
}

func NewQStyleFromPointer(ptr unsafe.Pointer) (n *QStyle) {
	n = new(QStyle)
	n.SetPointer(ptr)
	return
}

// QStyle::ComplexControl
//
//go:generate stringer -type=QStyle__ComplexControl
type QStyle__ComplexControl int64

const (
	QStyle__CC_SpinBox     QStyle__ComplexControl = QStyle__ComplexControl(0)
	QStyle__CC_ComboBox    QStyle__ComplexControl = QStyle__ComplexControl(1)
	QStyle__CC_ScrollBar   QStyle__ComplexControl = QStyle__ComplexControl(2)
	QStyle__CC_Slider      QStyle__ComplexControl = QStyle__ComplexControl(3)
	QStyle__CC_ToolButton  QStyle__ComplexControl = QStyle__ComplexControl(4)
	QStyle__CC_TitleBar    QStyle__ComplexControl = QStyle__ComplexControl(5)
	QStyle__CC_Dial        QStyle__ComplexControl = QStyle__ComplexControl(6)
	QStyle__CC_GroupBox    QStyle__ComplexControl = QStyle__ComplexControl(7)
	QStyle__CC_MdiControls QStyle__ComplexControl = QStyle__ComplexControl(8)
	QStyle__CC_CustomBase  QStyle__ComplexControl = QStyle__ComplexControl(0xf0000000)
)

// QStyle::ContentsType
//
//go:generate stringer -type=QStyle__ContentsType
type QStyle__ContentsType int64

const (
	QStyle__CT_PushButton    QStyle__ContentsType = QStyle__ContentsType(0)
	QStyle__CT_CheckBox      QStyle__ContentsType = QStyle__ContentsType(1)
	QStyle__CT_RadioButton   QStyle__ContentsType = QStyle__ContentsType(2)
	QStyle__CT_ToolButton    QStyle__ContentsType = QStyle__ContentsType(3)
	QStyle__CT_ComboBox      QStyle__ContentsType = QStyle__ContentsType(4)
	QStyle__CT_Splitter      QStyle__ContentsType = QStyle__ContentsType(5)
	QStyle__CT_ProgressBar   QStyle__ContentsType = QStyle__ContentsType(6)
	QStyle__CT_MenuItem      QStyle__ContentsType = QStyle__ContentsType(7)
	QStyle__CT_MenuBarItem   QStyle__ContentsType = QStyle__ContentsType(8)
	QStyle__CT_MenuBar       QStyle__ContentsType = QStyle__ContentsType(9)
	QStyle__CT_Menu          QStyle__ContentsType = QStyle__ContentsType(10)
	QStyle__CT_TabBarTab     QStyle__ContentsType = QStyle__ContentsType(11)
	QStyle__CT_Slider        QStyle__ContentsType = QStyle__ContentsType(12)
	QStyle__CT_ScrollBar     QStyle__ContentsType = QStyle__ContentsType(13)
	QStyle__CT_LineEdit      QStyle__ContentsType = QStyle__ContentsType(14)
	QStyle__CT_SpinBox       QStyle__ContentsType = QStyle__ContentsType(15)
	QStyle__CT_SizeGrip      QStyle__ContentsType = QStyle__ContentsType(16)
	QStyle__CT_TabWidget     QStyle__ContentsType = QStyle__ContentsType(17)
	QStyle__CT_DialogButtons QStyle__ContentsType = QStyle__ContentsType(18)
	QStyle__CT_HeaderSection QStyle__ContentsType = QStyle__ContentsType(19)
	QStyle__CT_GroupBox      QStyle__ContentsType = QStyle__ContentsType(20)
	QStyle__CT_MdiControls   QStyle__ContentsType = QStyle__ContentsType(21)
	QStyle__CT_ItemViewItem  QStyle__ContentsType = QStyle__ContentsType(22)
	QStyle__CT_CustomBase    QStyle__ContentsType = QStyle__ContentsType(0xf0000000)
)

// QStyle::ControlElement
//
//go:generate stringer -type=QStyle__ControlElement
type QStyle__ControlElement int64

const (
	QStyle__CE_PushButton          QStyle__ControlElement = QStyle__ControlElement(0)
	QStyle__CE_PushButtonBevel     QStyle__ControlElement = QStyle__ControlElement(1)
	QStyle__CE_PushButtonLabel     QStyle__ControlElement = QStyle__ControlElement(2)
	QStyle__CE_CheckBox            QStyle__ControlElement = QStyle__ControlElement(3)
	QStyle__CE_CheckBoxLabel       QStyle__ControlElement = QStyle__ControlElement(4)
	QStyle__CE_RadioButton         QStyle__ControlElement = QStyle__ControlElement(5)
	QStyle__CE_RadioButtonLabel    QStyle__ControlElement = QStyle__ControlElement(6)
	QStyle__CE_TabBarTab           QStyle__ControlElement = QStyle__ControlElement(7)
	QStyle__CE_TabBarTabShape      QStyle__ControlElement = QStyle__ControlElement(8)
	QStyle__CE_TabBarTabLabel      QStyle__ControlElement = QStyle__ControlElement(9)
	QStyle__CE_ProgressBar         QStyle__ControlElement = QStyle__ControlElement(10)
	QStyle__CE_ProgressBarGroove   QStyle__ControlElement = QStyle__ControlElement(11)
	QStyle__CE_ProgressBarContents QStyle__ControlElement = QStyle__ControlElement(12)
	QStyle__CE_ProgressBarLabel    QStyle__ControlElement = QStyle__ControlElement(13)
	QStyle__CE_MenuItem            QStyle__ControlElement = QStyle__ControlElement(14)
	QStyle__CE_MenuScroller        QStyle__ControlElement = QStyle__ControlElement(15)
	QStyle__CE_MenuVMargin         QStyle__ControlElement = QStyle__ControlElement(16)
	QStyle__CE_MenuHMargin         QStyle__ControlElement = QStyle__ControlElement(17)
	QStyle__CE_MenuTearoff         QStyle__ControlElement = QStyle__ControlElement(18)
	QStyle__CE_MenuEmptyArea       QStyle__ControlElement = QStyle__ControlElement(19)
	QStyle__CE_MenuBarItem         QStyle__ControlElement = QStyle__ControlElement(20)
	QStyle__CE_MenuBarEmptyArea    QStyle__ControlElement = QStyle__ControlElement(21)
	QStyle__CE_ToolButtonLabel     QStyle__ControlElement = QStyle__ControlElement(22)
	QStyle__CE_Header              QStyle__ControlElement = QStyle__ControlElement(23)
	QStyle__CE_HeaderSection       QStyle__ControlElement = QStyle__ControlElement(24)
	QStyle__CE_HeaderLabel         QStyle__ControlElement = QStyle__ControlElement(25)
	QStyle__CE_ToolBoxTab          QStyle__ControlElement = QStyle__ControlElement(26)
	QStyle__CE_SizeGrip            QStyle__ControlElement = QStyle__ControlElement(27)
	QStyle__CE_Splitter            QStyle__ControlElement = QStyle__ControlElement(28)
	QStyle__CE_RubberBand          QStyle__ControlElement = QStyle__ControlElement(29)
	QStyle__CE_DockWidgetTitle     QStyle__ControlElement = QStyle__ControlElement(30)
	QStyle__CE_ScrollBarAddLine    QStyle__ControlElement = QStyle__ControlElement(31)
	QStyle__CE_ScrollBarSubLine    QStyle__ControlElement = QStyle__ControlElement(32)
	QStyle__CE_ScrollBarAddPage    QStyle__ControlElement = QStyle__ControlElement(33)
	QStyle__CE_ScrollBarSubPage    QStyle__ControlElement = QStyle__ControlElement(34)
	QStyle__CE_ScrollBarSlider     QStyle__ControlElement = QStyle__ControlElement(35)
	QStyle__CE_ScrollBarFirst      QStyle__ControlElement = QStyle__ControlElement(36)
	QStyle__CE_ScrollBarLast       QStyle__ControlElement = QStyle__ControlElement(37)
	QStyle__CE_FocusFrame          QStyle__ControlElement = QStyle__ControlElement(38)
	QStyle__CE_ComboBoxLabel       QStyle__ControlElement = QStyle__ControlElement(39)
	QStyle__CE_ToolBar             QStyle__ControlElement = QStyle__ControlElement(40)
	QStyle__CE_ToolBoxTabShape     QStyle__ControlElement = QStyle__ControlElement(41)
	QStyle__CE_ToolBoxTabLabel     QStyle__ControlElement = QStyle__ControlElement(42)
	QStyle__CE_HeaderEmptyArea     QStyle__ControlElement = QStyle__ControlElement(43)
	QStyle__CE_ColumnViewGrip      QStyle__ControlElement = QStyle__ControlElement(44)
	QStyle__CE_ItemViewItem        QStyle__ControlElement = QStyle__ControlElement(45)
	QStyle__CE_ShapedFrame         QStyle__ControlElement = QStyle__ControlElement(46)
	QStyle__CE_CustomBase          QStyle__ControlElement = QStyle__ControlElement(0xf0000000)
)

// QStyle::PixelMetric
//
//go:generate stringer -type=QStyle__PixelMetric
type QStyle__PixelMetric int64

var (
	QStyle__PM_ButtonMargin                       QStyle__PixelMetric = QStyle__PixelMetric(0)
	QStyle__PM_ButtonDefaultIndicator             QStyle__PixelMetric = QStyle__PixelMetric(1)
	QStyle__PM_MenuButtonIndicator                QStyle__PixelMetric = QStyle__PixelMetric(2)
	QStyle__PM_ButtonShiftHorizontal              QStyle__PixelMetric = QStyle__PixelMetric(3)
	QStyle__PM_ButtonShiftVertical                QStyle__PixelMetric = QStyle__PixelMetric(4)
	QStyle__PM_DefaultFrameWidth                  QStyle__PixelMetric = QStyle__PixelMetric(5)
	QStyle__PM_SpinBoxFrameWidth                  QStyle__PixelMetric = QStyle__PixelMetric(6)
	QStyle__PM_ComboBoxFrameWidth                 QStyle__PixelMetric = QStyle__PixelMetric(7)
	QStyle__PM_MaximumDragDistance                QStyle__PixelMetric = QStyle__PixelMetric(8)
	QStyle__PM_ScrollBarExtent                    QStyle__PixelMetric = QStyle__PixelMetric(9)
	QStyle__PM_ScrollBarSliderMin                 QStyle__PixelMetric = QStyle__PixelMetric(10)
	QStyle__PM_SliderThickness                    QStyle__PixelMetric = QStyle__PixelMetric(11)
	QStyle__PM_SliderControlThickness             QStyle__PixelMetric = QStyle__PixelMetric(12)
	QStyle__PM_SliderLength                       QStyle__PixelMetric = QStyle__PixelMetric(13)
	QStyle__PM_SliderTickmarkOffset               QStyle__PixelMetric = QStyle__PixelMetric(14)
	QStyle__PM_SliderSpaceAvailable               QStyle__PixelMetric = QStyle__PixelMetric(15)
	QStyle__PM_DockWidgetSeparatorExtent          QStyle__PixelMetric = QStyle__PixelMetric(16)
	QStyle__PM_DockWidgetHandleExtent             QStyle__PixelMetric = QStyle__PixelMetric(17)
	QStyle__PM_DockWidgetFrameWidth               QStyle__PixelMetric = QStyle__PixelMetric(18)
	QStyle__PM_TabBarTabOverlap                   QStyle__PixelMetric = QStyle__PixelMetric(19)
	QStyle__PM_TabBarTabHSpace                    QStyle__PixelMetric = QStyle__PixelMetric(20)
	QStyle__PM_TabBarTabVSpace                    QStyle__PixelMetric = QStyle__PixelMetric(21)
	QStyle__PM_TabBarBaseHeight                   QStyle__PixelMetric = QStyle__PixelMetric(22)
	QStyle__PM_TabBarBaseOverlap                  QStyle__PixelMetric = QStyle__PixelMetric(23)
	QStyle__PM_ProgressBarChunkWidth              QStyle__PixelMetric = QStyle__PixelMetric(24)
	QStyle__PM_SplitterWidth                      QStyle__PixelMetric = QStyle__PixelMetric(25)
	QStyle__PM_TitleBarHeight                     QStyle__PixelMetric = QStyle__PixelMetric(26)
	QStyle__PM_MenuScrollerHeight                 QStyle__PixelMetric = QStyle__PixelMetric(27)
	QStyle__PM_MenuHMargin                        QStyle__PixelMetric = QStyle__PixelMetric(28)
	QStyle__PM_MenuVMargin                        QStyle__PixelMetric = QStyle__PixelMetric(29)
	QStyle__PM_MenuPanelWidth                     QStyle__PixelMetric = QStyle__PixelMetric(30)
	QStyle__PM_MenuTearoffHeight                  QStyle__PixelMetric = QStyle__PixelMetric(31)
	QStyle__PM_MenuDesktopFrameWidth              QStyle__PixelMetric = QStyle__PixelMetric(32)
	QStyle__PM_MenuBarPanelWidth                  QStyle__PixelMetric = QStyle__PixelMetric(33)
	QStyle__PM_MenuBarItemSpacing                 QStyle__PixelMetric = QStyle__PixelMetric(34)
	QStyle__PM_MenuBarVMargin                     QStyle__PixelMetric = QStyle__PixelMetric(35)
	QStyle__PM_MenuBarHMargin                     QStyle__PixelMetric = QStyle__PixelMetric(36)
	QStyle__PM_IndicatorWidth                     QStyle__PixelMetric = QStyle__PixelMetric(37)
	QStyle__PM_IndicatorHeight                    QStyle__PixelMetric = QStyle__PixelMetric(38)
	QStyle__PM_ExclusiveIndicatorWidth            QStyle__PixelMetric = QStyle__PixelMetric(39)
	QStyle__PM_ExclusiveIndicatorHeight           QStyle__PixelMetric = QStyle__PixelMetric(40)
	QStyle__PM_DialogButtonsSeparator             QStyle__PixelMetric = QStyle__PixelMetric(41)
	QStyle__PM_DialogButtonsButtonWidth           QStyle__PixelMetric = QStyle__PixelMetric(42)
	QStyle__PM_DialogButtonsButtonHeight          QStyle__PixelMetric = QStyle__PixelMetric(43)
	QStyle__PM_MdiSubWindowFrameWidth             QStyle__PixelMetric = QStyle__PixelMetric(44)
	QStyle__PM_MdiSubWindowMinimizedWidth         QStyle__PixelMetric = QStyle__PixelMetric(45)
	QStyle__PM_MDIFrameWidth                      QStyle__PixelMetric = QStyle__PixelMetric(QStyle__PM_MdiSubWindowFrameWidth)
	QStyle__PM_MDIMinimizedWidth                  QStyle__PixelMetric = QStyle__PixelMetric(QStyle__PM_MdiSubWindowMinimizedWidth)
	QStyle__PM_HeaderMargin                       QStyle__PixelMetric = QStyle__PixelMetric(46)
	QStyle__PM_HeaderMarkSize                     QStyle__PixelMetric = QStyle__PixelMetric(47)
	QStyle__PM_HeaderGripMargin                   QStyle__PixelMetric = QStyle__PixelMetric(48)
	QStyle__PM_TabBarTabShiftHorizontal           QStyle__PixelMetric = QStyle__PixelMetric(49)
	QStyle__PM_TabBarTabShiftVertical             QStyle__PixelMetric = QStyle__PixelMetric(50)
	QStyle__PM_TabBarScrollButtonWidth            QStyle__PixelMetric = QStyle__PixelMetric(51)
	QStyle__PM_ToolBarFrameWidth                  QStyle__PixelMetric = QStyle__PixelMetric(52)
	QStyle__PM_ToolBarHandleExtent                QStyle__PixelMetric = QStyle__PixelMetric(53)
	QStyle__PM_ToolBarItemSpacing                 QStyle__PixelMetric = QStyle__PixelMetric(54)
	QStyle__PM_ToolBarItemMargin                  QStyle__PixelMetric = QStyle__PixelMetric(55)
	QStyle__PM_ToolBarSeparatorExtent             QStyle__PixelMetric = QStyle__PixelMetric(56)
	QStyle__PM_ToolBarExtensionExtent             QStyle__PixelMetric = QStyle__PixelMetric(57)
	QStyle__PM_SpinBoxSliderHeight                QStyle__PixelMetric = QStyle__PixelMetric(58)
	QStyle__PM_DefaultTopLevelMargin              QStyle__PixelMetric = QStyle__PixelMetric(59)
	QStyle__PM_DefaultChildMargin                 QStyle__PixelMetric = QStyle__PixelMetric(60)
	QStyle__PM_DefaultLayoutSpacing               QStyle__PixelMetric = QStyle__PixelMetric(61)
	QStyle__PM_ToolBarIconSize                    QStyle__PixelMetric = QStyle__PixelMetric(C.QStyle_PM_ToolBarIconSize_Type())
	QStyle__PM_ListViewIconSize                   QStyle__PixelMetric = QStyle__PixelMetric(63)
	QStyle__PM_IconViewIconSize                   QStyle__PixelMetric = QStyle__PixelMetric(64)
	QStyle__PM_SmallIconSize                      QStyle__PixelMetric = QStyle__PixelMetric(65)
	QStyle__PM_LargeIconSize                      QStyle__PixelMetric = QStyle__PixelMetric(66)
	QStyle__PM_FocusFrameVMargin                  QStyle__PixelMetric = QStyle__PixelMetric(67)
	QStyle__PM_FocusFrameHMargin                  QStyle__PixelMetric = QStyle__PixelMetric(68)
	QStyle__PM_ToolTipLabelFrameWidth             QStyle__PixelMetric = QStyle__PixelMetric(69)
	QStyle__PM_CheckBoxLabelSpacing               QStyle__PixelMetric = QStyle__PixelMetric(70)
	QStyle__PM_TabBarIconSize                     QStyle__PixelMetric = QStyle__PixelMetric(71)
	QStyle__PM_SizeGripSize                       QStyle__PixelMetric = QStyle__PixelMetric(72)
	QStyle__PM_DockWidgetTitleMargin              QStyle__PixelMetric = QStyle__PixelMetric(73)
	QStyle__PM_MessageBoxIconSize                 QStyle__PixelMetric = QStyle__PixelMetric(74)
	QStyle__PM_ButtonIconSize                     QStyle__PixelMetric = QStyle__PixelMetric(75)
	QStyle__PM_DockWidgetTitleBarButtonMargin     QStyle__PixelMetric = QStyle__PixelMetric(76)
	QStyle__PM_RadioButtonLabelSpacing            QStyle__PixelMetric = QStyle__PixelMetric(77)
	QStyle__PM_LayoutLeftMargin                   QStyle__PixelMetric = QStyle__PixelMetric(78)
	QStyle__PM_LayoutTopMargin                    QStyle__PixelMetric = QStyle__PixelMetric(79)
	QStyle__PM_LayoutRightMargin                  QStyle__PixelMetric = QStyle__PixelMetric(80)
	QStyle__PM_LayoutBottomMargin                 QStyle__PixelMetric = QStyle__PixelMetric(81)
	QStyle__PM_LayoutHorizontalSpacing            QStyle__PixelMetric = QStyle__PixelMetric(82)
	QStyle__PM_LayoutVerticalSpacing              QStyle__PixelMetric = QStyle__PixelMetric(83)
	QStyle__PM_TabBar_ScrollButtonOverlap         QStyle__PixelMetric = QStyle__PixelMetric(84)
	QStyle__PM_TextCursorWidth                    QStyle__PixelMetric = QStyle__PixelMetric(85)
	QStyle__PM_TabCloseIndicatorWidth             QStyle__PixelMetric = QStyle__PixelMetric(86)
	QStyle__PM_TabCloseIndicatorHeight            QStyle__PixelMetric = QStyle__PixelMetric(87)
	QStyle__PM_ScrollView_ScrollBarSpacing        QStyle__PixelMetric = QStyle__PixelMetric(88)
	QStyle__PM_ScrollView_ScrollBarOverlap        QStyle__PixelMetric = QStyle__PixelMetric(89)
	QStyle__PM_SubMenuOverlap                     QStyle__PixelMetric = QStyle__PixelMetric(90)
	QStyle__PM_TreeViewIndentation                QStyle__PixelMetric = QStyle__PixelMetric(91)
	QStyle__PM_HeaderDefaultSectionSizeHorizontal QStyle__PixelMetric = QStyle__PixelMetric(92)
	QStyle__PM_HeaderDefaultSectionSizeVertical   QStyle__PixelMetric = QStyle__PixelMetric(93)
	QStyle__PM_TitleBarButtonIconSize             QStyle__PixelMetric = QStyle__PixelMetric(94)
	QStyle__PM_TitleBarButtonSize                 QStyle__PixelMetric = QStyle__PixelMetric(95)
	QStyle__PM_CustomBase                         QStyle__PixelMetric = QStyle__PixelMetric(0xf0000000)
)

// QStyle::PrimitiveElement
//
//go:generate stringer -type=QStyle__PrimitiveElement
type QStyle__PrimitiveElement int64

const (
	QStyle__PE_Frame                           QStyle__PrimitiveElement = QStyle__PrimitiveElement(0)
	QStyle__PE_FrameDefaultButton              QStyle__PrimitiveElement = QStyle__PrimitiveElement(1)
	QStyle__PE_FrameDockWidget                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(2)
	QStyle__PE_FrameFocusRect                  QStyle__PrimitiveElement = QStyle__PrimitiveElement(3)
	QStyle__PE_FrameGroupBox                   QStyle__PrimitiveElement = QStyle__PrimitiveElement(4)
	QStyle__PE_FrameLineEdit                   QStyle__PrimitiveElement = QStyle__PrimitiveElement(5)
	QStyle__PE_FrameMenu                       QStyle__PrimitiveElement = QStyle__PrimitiveElement(6)
	QStyle__PE_FrameStatusBarItem              QStyle__PrimitiveElement = QStyle__PrimitiveElement(7)
	QStyle__PE_FrameStatusBar                  QStyle__PrimitiveElement = QStyle__PrimitiveElement(QStyle__PE_FrameStatusBarItem)
	QStyle__PE_FrameTabWidget                  QStyle__PrimitiveElement = QStyle__PrimitiveElement(8)
	QStyle__PE_FrameWindow                     QStyle__PrimitiveElement = QStyle__PrimitiveElement(9)
	QStyle__PE_FrameButtonBevel                QStyle__PrimitiveElement = QStyle__PrimitiveElement(10)
	QStyle__PE_FrameButtonTool                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(11)
	QStyle__PE_FrameTabBarBase                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(12)
	QStyle__PE_PanelButtonCommand              QStyle__PrimitiveElement = QStyle__PrimitiveElement(13)
	QStyle__PE_PanelButtonBevel                QStyle__PrimitiveElement = QStyle__PrimitiveElement(14)
	QStyle__PE_PanelButtonTool                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(15)
	QStyle__PE_PanelMenuBar                    QStyle__PrimitiveElement = QStyle__PrimitiveElement(16)
	QStyle__PE_PanelToolBar                    QStyle__PrimitiveElement = QStyle__PrimitiveElement(17)
	QStyle__PE_PanelLineEdit                   QStyle__PrimitiveElement = QStyle__PrimitiveElement(18)
	QStyle__PE_IndicatorArrowDown              QStyle__PrimitiveElement = QStyle__PrimitiveElement(19)
	QStyle__PE_IndicatorArrowLeft              QStyle__PrimitiveElement = QStyle__PrimitiveElement(20)
	QStyle__PE_IndicatorArrowRight             QStyle__PrimitiveElement = QStyle__PrimitiveElement(21)
	QStyle__PE_IndicatorArrowUp                QStyle__PrimitiveElement = QStyle__PrimitiveElement(22)
	QStyle__PE_IndicatorBranch                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(23)
	QStyle__PE_IndicatorButtonDropDown         QStyle__PrimitiveElement = QStyle__PrimitiveElement(24)
	QStyle__PE_IndicatorItemViewItemCheck      QStyle__PrimitiveElement = QStyle__PrimitiveElement(25)
	QStyle__PE_IndicatorViewItemCheck          QStyle__PrimitiveElement = QStyle__PrimitiveElement(QStyle__PE_IndicatorItemViewItemCheck)
	QStyle__PE_IndicatorCheckBox               QStyle__PrimitiveElement = QStyle__PrimitiveElement(26)
	QStyle__PE_IndicatorDockWidgetResizeHandle QStyle__PrimitiveElement = QStyle__PrimitiveElement(27)
	QStyle__PE_IndicatorHeaderArrow            QStyle__PrimitiveElement = QStyle__PrimitiveElement(28)
	QStyle__PE_IndicatorMenuCheckMark          QStyle__PrimitiveElement = QStyle__PrimitiveElement(29)
	QStyle__PE_IndicatorProgressChunk          QStyle__PrimitiveElement = QStyle__PrimitiveElement(30)
	QStyle__PE_IndicatorRadioButton            QStyle__PrimitiveElement = QStyle__PrimitiveElement(31)
	QStyle__PE_IndicatorSpinDown               QStyle__PrimitiveElement = QStyle__PrimitiveElement(32)
	QStyle__PE_IndicatorSpinMinus              QStyle__PrimitiveElement = QStyle__PrimitiveElement(33)
	QStyle__PE_IndicatorSpinPlus               QStyle__PrimitiveElement = QStyle__PrimitiveElement(34)
	QStyle__PE_IndicatorSpinUp                 QStyle__PrimitiveElement = QStyle__PrimitiveElement(35)
	QStyle__PE_IndicatorToolBarHandle          QStyle__PrimitiveElement = QStyle__PrimitiveElement(36)
	QStyle__PE_IndicatorToolBarSeparator       QStyle__PrimitiveElement = QStyle__PrimitiveElement(37)
	QStyle__PE_PanelTipLabel                   QStyle__PrimitiveElement = QStyle__PrimitiveElement(38)
	QStyle__PE_IndicatorTabTear                QStyle__PrimitiveElement = QStyle__PrimitiveElement(39)
	QStyle__PE_IndicatorTabTearLeft            QStyle__PrimitiveElement = QStyle__PrimitiveElement(QStyle__PE_IndicatorTabTear)
	QStyle__PE_PanelScrollAreaCorner           QStyle__PrimitiveElement = QStyle__PrimitiveElement(40)
	QStyle__PE_Widget                          QStyle__PrimitiveElement = QStyle__PrimitiveElement(41)
	QStyle__PE_IndicatorColumnViewArrow        QStyle__PrimitiveElement = QStyle__PrimitiveElement(42)
	QStyle__PE_IndicatorItemViewItemDrop       QStyle__PrimitiveElement = QStyle__PrimitiveElement(43)
	QStyle__PE_PanelItemViewItem               QStyle__PrimitiveElement = QStyle__PrimitiveElement(44)
	QStyle__PE_PanelItemViewRow                QStyle__PrimitiveElement = QStyle__PrimitiveElement(45)
	QStyle__PE_PanelStatusBar                  QStyle__PrimitiveElement = QStyle__PrimitiveElement(46)
	QStyle__PE_IndicatorTabClose               QStyle__PrimitiveElement = QStyle__PrimitiveElement(47)
	QStyle__PE_PanelMenu                       QStyle__PrimitiveElement = QStyle__PrimitiveElement(48)
	QStyle__PE_IndicatorTabTearRight           QStyle__PrimitiveElement = QStyle__PrimitiveElement(49)
	QStyle__PE_CustomBase                      QStyle__PrimitiveElement = QStyle__PrimitiveElement(0xf000000)
)

// QStyle::RequestSoftwareInputPanel
//
//go:generate stringer -type=QStyle__RequestSoftwareInputPanel
type QStyle__RequestSoftwareInputPanel int64

const (
	QStyle__RSIP_OnMouseClickAndAlreadyFocused QStyle__RequestSoftwareInputPanel = QStyle__RequestSoftwareInputPanel(0)
	QStyle__RSIP_OnMouseClick                  QStyle__RequestSoftwareInputPanel = QStyle__RequestSoftwareInputPanel(1)
)

// QStyle::StandardPixmap
//
//go:generate stringer -type=QStyle__StandardPixmap
type QStyle__StandardPixmap int64

const (
	QStyle__SP_TitleBarMenuButton               QStyle__StandardPixmap = QStyle__StandardPixmap(0)
	QStyle__SP_TitleBarMinButton                QStyle__StandardPixmap = QStyle__StandardPixmap(1)
	QStyle__SP_TitleBarMaxButton                QStyle__StandardPixmap = QStyle__StandardPixmap(2)
	QStyle__SP_TitleBarCloseButton              QStyle__StandardPixmap = QStyle__StandardPixmap(3)
	QStyle__SP_TitleBarNormalButton             QStyle__StandardPixmap = QStyle__StandardPixmap(4)
	QStyle__SP_TitleBarShadeButton              QStyle__StandardPixmap = QStyle__StandardPixmap(5)
	QStyle__SP_TitleBarUnshadeButton            QStyle__StandardPixmap = QStyle__StandardPixmap(6)
	QStyle__SP_TitleBarContextHelpButton        QStyle__StandardPixmap = QStyle__StandardPixmap(7)
	QStyle__SP_DockWidgetCloseButton            QStyle__StandardPixmap = QStyle__StandardPixmap(8)
	QStyle__SP_MessageBoxInformation            QStyle__StandardPixmap = QStyle__StandardPixmap(9)
	QStyle__SP_MessageBoxWarning                QStyle__StandardPixmap = QStyle__StandardPixmap(10)
	QStyle__SP_MessageBoxCritical               QStyle__StandardPixmap = QStyle__StandardPixmap(11)
	QStyle__SP_MessageBoxQuestion               QStyle__StandardPixmap = QStyle__StandardPixmap(12)
	QStyle__SP_DesktopIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(13)
	QStyle__SP_TrashIcon                        QStyle__StandardPixmap = QStyle__StandardPixmap(14)
	QStyle__SP_ComputerIcon                     QStyle__StandardPixmap = QStyle__StandardPixmap(15)
	QStyle__SP_DriveFDIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(16)
	QStyle__SP_DriveHDIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(17)
	QStyle__SP_DriveCDIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(18)
	QStyle__SP_DriveDVDIcon                     QStyle__StandardPixmap = QStyle__StandardPixmap(19)
	QStyle__SP_DriveNetIcon                     QStyle__StandardPixmap = QStyle__StandardPixmap(20)
	QStyle__SP_DirOpenIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(21)
	QStyle__SP_DirClosedIcon                    QStyle__StandardPixmap = QStyle__StandardPixmap(22)
	QStyle__SP_DirLinkIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(23)
	QStyle__SP_DirLinkOpenIcon                  QStyle__StandardPixmap = QStyle__StandardPixmap(24)
	QStyle__SP_FileIcon                         QStyle__StandardPixmap = QStyle__StandardPixmap(25)
	QStyle__SP_FileLinkIcon                     QStyle__StandardPixmap = QStyle__StandardPixmap(26)
	QStyle__SP_ToolBarHorizontalExtensionButton QStyle__StandardPixmap = QStyle__StandardPixmap(27)
	QStyle__SP_ToolBarVerticalExtensionButton   QStyle__StandardPixmap = QStyle__StandardPixmap(28)
	QStyle__SP_FileDialogStart                  QStyle__StandardPixmap = QStyle__StandardPixmap(29)
	QStyle__SP_FileDialogEnd                    QStyle__StandardPixmap = QStyle__StandardPixmap(30)
	QStyle__SP_FileDialogToParent               QStyle__StandardPixmap = QStyle__StandardPixmap(31)
	QStyle__SP_FileDialogNewFolder              QStyle__StandardPixmap = QStyle__StandardPixmap(32)
	QStyle__SP_FileDialogDetailedView           QStyle__StandardPixmap = QStyle__StandardPixmap(33)
	QStyle__SP_FileDialogInfoView               QStyle__StandardPixmap = QStyle__StandardPixmap(34)
	QStyle__SP_FileDialogContentsView           QStyle__StandardPixmap = QStyle__StandardPixmap(35)
	QStyle__SP_FileDialogListView               QStyle__StandardPixmap = QStyle__StandardPixmap(36)
	QStyle__SP_FileDialogBack                   QStyle__StandardPixmap = QStyle__StandardPixmap(37)
	QStyle__SP_DirIcon                          QStyle__StandardPixmap = QStyle__StandardPixmap(38)
	QStyle__SP_DialogOkButton                   QStyle__StandardPixmap = QStyle__StandardPixmap(39)
	QStyle__SP_DialogCancelButton               QStyle__StandardPixmap = QStyle__StandardPixmap(40)
	QStyle__SP_DialogHelpButton                 QStyle__StandardPixmap = QStyle__StandardPixmap(41)
	QStyle__SP_DialogOpenButton                 QStyle__StandardPixmap = QStyle__StandardPixmap(42)
	QStyle__SP_DialogSaveButton                 QStyle__StandardPixmap = QStyle__StandardPixmap(43)
	QStyle__SP_DialogCloseButton                QStyle__StandardPixmap = QStyle__StandardPixmap(44)
	QStyle__SP_DialogApplyButton                QStyle__StandardPixmap = QStyle__StandardPixmap(45)
	QStyle__SP_DialogResetButton                QStyle__StandardPixmap = QStyle__StandardPixmap(46)
	QStyle__SP_DialogDiscardButton              QStyle__StandardPixmap = QStyle__StandardPixmap(47)
	QStyle__SP_DialogYesButton                  QStyle__StandardPixmap = QStyle__StandardPixmap(48)
	QStyle__SP_DialogNoButton                   QStyle__StandardPixmap = QStyle__StandardPixmap(49)
	QStyle__SP_ArrowUp                          QStyle__StandardPixmap = QStyle__StandardPixmap(50)
	QStyle__SP_ArrowDown                        QStyle__StandardPixmap = QStyle__StandardPixmap(51)
	QStyle__SP_ArrowLeft                        QStyle__StandardPixmap = QStyle__StandardPixmap(52)
	QStyle__SP_ArrowRight                       QStyle__StandardPixmap = QStyle__StandardPixmap(53)
	QStyle__SP_ArrowBack                        QStyle__StandardPixmap = QStyle__StandardPixmap(54)
	QStyle__SP_ArrowForward                     QStyle__StandardPixmap = QStyle__StandardPixmap(55)
	QStyle__SP_DirHomeIcon                      QStyle__StandardPixmap = QStyle__StandardPixmap(56)
	QStyle__SP_CommandLink                      QStyle__StandardPixmap = QStyle__StandardPixmap(57)
	QStyle__SP_VistaShield                      QStyle__StandardPixmap = QStyle__StandardPixmap(58)
	QStyle__SP_BrowserReload                    QStyle__StandardPixmap = QStyle__StandardPixmap(59)
	QStyle__SP_BrowserStop                      QStyle__StandardPixmap = QStyle__StandardPixmap(60)
	QStyle__SP_MediaPlay                        QStyle__StandardPixmap = QStyle__StandardPixmap(61)
	QStyle__SP_MediaStop                        QStyle__StandardPixmap = QStyle__StandardPixmap(62)
	QStyle__SP_MediaPause                       QStyle__StandardPixmap = QStyle__StandardPixmap(63)
	QStyle__SP_MediaSkipForward                 QStyle__StandardPixmap = QStyle__StandardPixmap(64)
	QStyle__SP_MediaSkipBackward                QStyle__StandardPixmap = QStyle__StandardPixmap(65)
	QStyle__SP_MediaSeekForward                 QStyle__StandardPixmap = QStyle__StandardPixmap(66)
	QStyle__SP_MediaSeekBackward                QStyle__StandardPixmap = QStyle__StandardPixmap(67)
	QStyle__SP_MediaVolume                      QStyle__StandardPixmap = QStyle__StandardPixmap(68)
	QStyle__SP_MediaVolumeMuted                 QStyle__StandardPixmap = QStyle__StandardPixmap(69)
	QStyle__SP_LineEditClearButton              QStyle__StandardPixmap = QStyle__StandardPixmap(70)
	QStyle__SP_DialogYesToAllButton             QStyle__StandardPixmap = QStyle__StandardPixmap(71)
	QStyle__SP_DialogNoToAllButton              QStyle__StandardPixmap = QStyle__StandardPixmap(72)
	QStyle__SP_DialogSaveAllButton              QStyle__StandardPixmap = QStyle__StandardPixmap(73)
	QStyle__SP_DialogAbortButton                QStyle__StandardPixmap = QStyle__StandardPixmap(74)
	QStyle__SP_DialogRetryButton                QStyle__StandardPixmap = QStyle__StandardPixmap(75)
	QStyle__SP_DialogIgnoreButton               QStyle__StandardPixmap = QStyle__StandardPixmap(76)
	QStyle__SP_RestoreDefaultsButton            QStyle__StandardPixmap = QStyle__StandardPixmap(77)
	QStyle__SP_CustomBase                       QStyle__StandardPixmap = QStyle__StandardPixmap(0xf0000000)
)

// QStyle::StateFlag
//
//go:generate stringer -type=QStyle__StateFlag
type QStyle__StateFlag int64

const (
	QStyle__State_None                QStyle__StateFlag = QStyle__StateFlag(0x00000000)
	QStyle__State_Enabled             QStyle__StateFlag = QStyle__StateFlag(0x00000001)
	QStyle__State_Raised              QStyle__StateFlag = QStyle__StateFlag(0x00000002)
	QStyle__State_Sunken              QStyle__StateFlag = QStyle__StateFlag(0x00000004)
	QStyle__State_Off                 QStyle__StateFlag = QStyle__StateFlag(0x00000008)
	QStyle__State_NoChange            QStyle__StateFlag = QStyle__StateFlag(0x00000010)
	QStyle__State_On                  QStyle__StateFlag = QStyle__StateFlag(0x00000020)
	QStyle__State_DownArrow           QStyle__StateFlag = QStyle__StateFlag(0x00000040)
	QStyle__State_Horizontal          QStyle__StateFlag = QStyle__StateFlag(0x00000080)
	QStyle__State_HasFocus            QStyle__StateFlag = QStyle__StateFlag(0x00000100)
	QStyle__State_Top                 QStyle__StateFlag = QStyle__StateFlag(0x00000200)
	QStyle__State_Bottom              QStyle__StateFlag = QStyle__StateFlag(0x00000400)
	QStyle__State_FocusAtBorder       QStyle__StateFlag = QStyle__StateFlag(0x00000800)
	QStyle__State_AutoRaise           QStyle__StateFlag = QStyle__StateFlag(0x00001000)
	QStyle__State_MouseOver           QStyle__StateFlag = QStyle__StateFlag(0x00002000)
	QStyle__State_UpArrow             QStyle__StateFlag = QStyle__StateFlag(0x00004000)
	QStyle__State_Selected            QStyle__StateFlag = QStyle__StateFlag(0x00008000)
	QStyle__State_Active              QStyle__StateFlag = QStyle__StateFlag(0x00010000)
	QStyle__State_Window              QStyle__StateFlag = QStyle__StateFlag(0x00020000)
	QStyle__State_Open                QStyle__StateFlag = QStyle__StateFlag(0x00040000)
	QStyle__State_Children            QStyle__StateFlag = QStyle__StateFlag(0x00080000)
	QStyle__State_Item                QStyle__StateFlag = QStyle__StateFlag(0x00100000)
	QStyle__State_Sibling             QStyle__StateFlag = QStyle__StateFlag(0x00200000)
	QStyle__State_Editing             QStyle__StateFlag = QStyle__StateFlag(0x00400000)
	QStyle__State_KeyboardFocusChange QStyle__StateFlag = QStyle__StateFlag(0x00800000)
	QStyle__State_HasEditFocus        QStyle__StateFlag = QStyle__StateFlag(0x01000000)
	QStyle__State_ReadOnly            QStyle__StateFlag = QStyle__StateFlag(0x02000000)
	QStyle__State_Small               QStyle__StateFlag = QStyle__StateFlag(0x04000000)
	QStyle__State_Mini                QStyle__StateFlag = QStyle__StateFlag(0x08000000)
)

// QStyle::StyleHint
//
//go:generate stringer -type=QStyle__StyleHint
type QStyle__StyleHint int64

var (
	QStyle__SH_EtchDisabledText                               QStyle__StyleHint = QStyle__StyleHint(0)
	QStyle__SH_DitherDisabledText                             QStyle__StyleHint = QStyle__StyleHint(1)
	QStyle__SH_ScrollBar_MiddleClickAbsolutePosition          QStyle__StyleHint = QStyle__StyleHint(2)
	QStyle__SH_ScrollBar_ScrollWhenPointerLeavesControl       QStyle__StyleHint = QStyle__StyleHint(3)
	QStyle__SH_TabBar_SelectMouseType                         QStyle__StyleHint = QStyle__StyleHint(4)
	QStyle__SH_TabBar_Alignment                               QStyle__StyleHint = QStyle__StyleHint(5)
	QStyle__SH_Header_ArrowAlignment                          QStyle__StyleHint = QStyle__StyleHint(6)
	QStyle__SH_Slider_SnapToValue                             QStyle__StyleHint = QStyle__StyleHint(7)
	QStyle__SH_Slider_SloppyKeyEvents                         QStyle__StyleHint = QStyle__StyleHint(8)
	QStyle__SH_ProgressDialog_CenterCancelButton              QStyle__StyleHint = QStyle__StyleHint(9)
	QStyle__SH_ProgressDialog_TextLabelAlignment              QStyle__StyleHint = QStyle__StyleHint(10)
	QStyle__SH_PrintDialog_RightAlignButtons                  QStyle__StyleHint = QStyle__StyleHint(11)
	QStyle__SH_MainWindow_SpaceBelowMenuBar                   QStyle__StyleHint = QStyle__StyleHint(12)
	QStyle__SH_FontDialog_SelectAssociatedText                QStyle__StyleHint = QStyle__StyleHint(13)
	QStyle__SH_Menu_AllowActiveAndDisabled                    QStyle__StyleHint = QStyle__StyleHint(14)
	QStyle__SH_Menu_SpaceActivatesItem                        QStyle__StyleHint = QStyle__StyleHint(15)
	QStyle__SH_Menu_SubMenuPopupDelay                         QStyle__StyleHint = QStyle__StyleHint(16)
	QStyle__SH_ScrollView_FrameOnlyAroundContents             QStyle__StyleHint = QStyle__StyleHint(17)
	QStyle__SH_MenuBar_AltKeyNavigation                       QStyle__StyleHint = QStyle__StyleHint(18)
	QStyle__SH_ComboBox_ListMouseTracking                     QStyle__StyleHint = QStyle__StyleHint(19)
	QStyle__SH_Menu_MouseTracking                             QStyle__StyleHint = QStyle__StyleHint(20)
	QStyle__SH_MenuBar_MouseTracking                          QStyle__StyleHint = QStyle__StyleHint(21)
	QStyle__SH_ItemView_ChangeHighlightOnFocus                QStyle__StyleHint = QStyle__StyleHint(22)
	QStyle__SH_Widget_ShareActivation                         QStyle__StyleHint = QStyle__StyleHint(23)
	QStyle__SH_Workspace_FillSpaceOnMaximize                  QStyle__StyleHint = QStyle__StyleHint(24)
	QStyle__SH_ComboBox_Popup                                 QStyle__StyleHint = QStyle__StyleHint(25)
	QStyle__SH_TitleBar_NoBorder                              QStyle__StyleHint = QStyle__StyleHint(26)
	QStyle__SH_Slider_StopMouseOverSlider                     QStyle__StyleHint = QStyle__StyleHint(27)
	QStyle__SH_ScrollBar_StopMouseOverSlider                  QStyle__StyleHint = QStyle__StyleHint(QStyle__SH_Slider_StopMouseOverSlider)
	QStyle__SH_BlinkCursorWhenTextSelected                    QStyle__StyleHint = QStyle__StyleHint(28)
	QStyle__SH_RichText_FullWidthSelection                    QStyle__StyleHint = QStyle__StyleHint(29)
	QStyle__SH_Menu_Scrollable                                QStyle__StyleHint = QStyle__StyleHint(30)
	QStyle__SH_GroupBox_TextLabelVerticalAlignment            QStyle__StyleHint = QStyle__StyleHint(31)
	QStyle__SH_GroupBox_TextLabelColor                        QStyle__StyleHint = QStyle__StyleHint(32)
	QStyle__SH_Menu_SloppySubMenus                            QStyle__StyleHint = QStyle__StyleHint(33)
	QStyle__SH_Table_GridLineColor                            QStyle__StyleHint = QStyle__StyleHint(34)
	QStyle__SH_LineEdit_PasswordCharacter                     QStyle__StyleHint = QStyle__StyleHint(35)
	QStyle__SH_DialogButtons_DefaultButton                    QStyle__StyleHint = QStyle__StyleHint(36)
	QStyle__SH_ToolBox_SelectedPageTitleBold                  QStyle__StyleHint = QStyle__StyleHint(37)
	QStyle__SH_TabBar_PreferNoArrows                          QStyle__StyleHint = QStyle__StyleHint(38)
	QStyle__SH_ScrollBar_LeftClickAbsolutePosition            QStyle__StyleHint = QStyle__StyleHint(39)
	QStyle__SH_ListViewExpand_SelectMouseType                 QStyle__StyleHint = QStyle__StyleHint(40)
	QStyle__SH_UnderlineShortcut                              QStyle__StyleHint = QStyle__StyleHint(41)
	QStyle__SH_SpinBox_AnimateButton                          QStyle__StyleHint = QStyle__StyleHint(42)
	QStyle__SH_SpinBox_KeyPressAutoRepeatRate                 QStyle__StyleHint = QStyle__StyleHint(43)
	QStyle__SH_SpinBox_ClickAutoRepeatRate                    QStyle__StyleHint = QStyle__StyleHint(44)
	QStyle__SH_Menu_FillScreenWithScroll                      QStyle__StyleHint = QStyle__StyleHint(45)
	QStyle__SH_ToolTipLabel_Opacity                           QStyle__StyleHint = QStyle__StyleHint(46)
	QStyle__SH_DrawMenuBarSeparator                           QStyle__StyleHint = QStyle__StyleHint(47)
	QStyle__SH_TitleBar_ModifyNotification                    QStyle__StyleHint = QStyle__StyleHint(48)
	QStyle__SH_Button_FocusPolicy                             QStyle__StyleHint = QStyle__StyleHint(49)
	QStyle__SH_MessageBox_UseBorderForButtonSpacing           QStyle__StyleHint = QStyle__StyleHint(50)
	QStyle__SH_TitleBar_AutoRaise                             QStyle__StyleHint = QStyle__StyleHint(51)
	QStyle__SH_ToolButton_PopupDelay                          QStyle__StyleHint = QStyle__StyleHint(52)
	QStyle__SH_FocusFrame_Mask                                QStyle__StyleHint = QStyle__StyleHint(53)
	QStyle__SH_RubberBand_Mask                                QStyle__StyleHint = QStyle__StyleHint(54)
	QStyle__SH_WindowFrame_Mask                               QStyle__StyleHint = QStyle__StyleHint(55)
	QStyle__SH_SpinControls_DisableOnBounds                   QStyle__StyleHint = QStyle__StyleHint(56)
	QStyle__SH_Dial_BackgroundRole                            QStyle__StyleHint = QStyle__StyleHint(57)
	QStyle__SH_ComboBox_LayoutDirection                       QStyle__StyleHint = QStyle__StyleHint(58)
	QStyle__SH_ItemView_EllipsisLocation                      QStyle__StyleHint = QStyle__StyleHint(59)
	QStyle__SH_ItemView_ShowDecorationSelected                QStyle__StyleHint = QStyle__StyleHint(60)
	QStyle__SH_ItemView_ActivateItemOnSingleClick             QStyle__StyleHint = QStyle__StyleHint(61)
	QStyle__SH_ScrollBar_ContextMenu                          QStyle__StyleHint = QStyle__StyleHint(62)
	QStyle__SH_ScrollBar_RollBetweenButtons                   QStyle__StyleHint = QStyle__StyleHint(63)
	QStyle__SH_Slider_AbsoluteSetButtons                      QStyle__StyleHint = QStyle__StyleHint(64)
	QStyle__SH_Slider_PageSetButtons                          QStyle__StyleHint = QStyle__StyleHint(65)
	QStyle__SH_Menu_KeyboardSearch                            QStyle__StyleHint = QStyle__StyleHint(66)
	QStyle__SH_TabBar_ElideMode                               QStyle__StyleHint = QStyle__StyleHint(67)
	QStyle__SH_DialogButtonLayout                             QStyle__StyleHint = QStyle__StyleHint(68)
	QStyle__SH_ComboBox_PopupFrameStyle                       QStyle__StyleHint = QStyle__StyleHint(69)
	QStyle__SH_MessageBox_TextInteractionFlags                QStyle__StyleHint = QStyle__StyleHint(70)
	QStyle__SH_DialogButtonBox_ButtonsHaveIcons               QStyle__StyleHint = QStyle__StyleHint(71)
	QStyle__SH_SpellCheckUnderlineStyle                       QStyle__StyleHint = QStyle__StyleHint(72)
	QStyle__SH_MessageBox_CenterButtons                       QStyle__StyleHint = QStyle__StyleHint(73)
	QStyle__SH_Menu_SelectionWrap                             QStyle__StyleHint = QStyle__StyleHint(74)
	QStyle__SH_ItemView_MovementWithoutUpdatingSelection      QStyle__StyleHint = QStyle__StyleHint(75)
	QStyle__SH_ToolTip_Mask                                   QStyle__StyleHint = QStyle__StyleHint(76)
	QStyle__SH_FocusFrame_AboveWidget                         QStyle__StyleHint = QStyle__StyleHint(77)
	QStyle__SH_TextControl_FocusIndicatorTextCharFormat       QStyle__StyleHint = QStyle__StyleHint(78)
	QStyle__SH_WizardStyle                                    QStyle__StyleHint = QStyle__StyleHint(79)
	QStyle__SH_ItemView_ArrowKeysNavigateIntoChildren         QStyle__StyleHint = QStyle__StyleHint(80)
	QStyle__SH_Menu_Mask                                      QStyle__StyleHint = QStyle__StyleHint(81)
	QStyle__SH_Menu_FlashTriggeredItem                        QStyle__StyleHint = QStyle__StyleHint(82)
	QStyle__SH_Menu_FadeOutOnHide                             QStyle__StyleHint = QStyle__StyleHint(83)
	QStyle__SH_SpinBox_ClickAutoRepeatThreshold               QStyle__StyleHint = QStyle__StyleHint(84)
	QStyle__SH_ItemView_PaintAlternatingRowColorsForEmptyArea QStyle__StyleHint = QStyle__StyleHint(85)
	QStyle__SH_FormLayoutWrapPolicy                           QStyle__StyleHint = QStyle__StyleHint(86)
	QStyle__SH_TabWidget_DefaultTabPosition                   QStyle__StyleHint = QStyle__StyleHint(87)
	QStyle__SH_ToolBar_Movable                                QStyle__StyleHint = QStyle__StyleHint(88)
	QStyle__SH_FormLayoutFieldGrowthPolicy                    QStyle__StyleHint = QStyle__StyleHint(89)
	QStyle__SH_FormLayoutFormAlignment                        QStyle__StyleHint = QStyle__StyleHint(90)
	QStyle__SH_FormLayoutLabelAlignment                       QStyle__StyleHint = QStyle__StyleHint(91)
	QStyle__SH_ItemView_DrawDelegateFrame                     QStyle__StyleHint = QStyle__StyleHint(92)
	QStyle__SH_TabBar_CloseButtonPosition                     QStyle__StyleHint = QStyle__StyleHint(93)
	QStyle__SH_DockWidget_ButtonsHaveFrame                    QStyle__StyleHint = QStyle__StyleHint(94)
	QStyle__SH_ToolButtonStyle                                QStyle__StyleHint = QStyle__StyleHint(95)
	QStyle__SH_RequestSoftwareInputPanel                      QStyle__StyleHint = QStyle__StyleHint(96)
	QStyle__SH_ScrollBar_Transient                            QStyle__StyleHint = QStyle__StyleHint(97)
	QStyle__SH_Menu_SupportsSections                          QStyle__StyleHint = QStyle__StyleHint(98)
	QStyle__SH_ToolTip_WakeUpDelay                            QStyle__StyleHint = QStyle__StyleHint(99)
	QStyle__SH_ToolTip_FallAsleepDelay                        QStyle__StyleHint = QStyle__StyleHint(100)
	QStyle__SH_Widget_Animate                                 QStyle__StyleHint = QStyle__StyleHint(101)
	QStyle__SH_Splitter_OpaqueResize                          QStyle__StyleHint = QStyle__StyleHint(102)
	QStyle__SH_ComboBox_UseNativePopup                        QStyle__StyleHint = QStyle__StyleHint(103)
	QStyle__SH_LineEdit_PasswordMaskDelay                     QStyle__StyleHint = QStyle__StyleHint(104)
	QStyle__SH_TabBar_ChangeCurrentDelay                      QStyle__StyleHint = QStyle__StyleHint(105)
	QStyle__SH_Menu_SubMenuUniDirection                       QStyle__StyleHint = QStyle__StyleHint(106)
	QStyle__SH_Menu_SubMenuUniDirectionFailCount              QStyle__StyleHint = QStyle__StyleHint(107)
	QStyle__SH_Menu_SubMenuSloppySelectOtherActions           QStyle__StyleHint = QStyle__StyleHint(108)
	QStyle__SH_Menu_SubMenuSloppyCloseTimeout                 QStyle__StyleHint = QStyle__StyleHint(109)
	QStyle__SH_Menu_SubMenuResetWhenReenteringParent          QStyle__StyleHint = QStyle__StyleHint(110)
	QStyle__SH_Menu_SubMenuDontStartSloppyOnLeave             QStyle__StyleHint = QStyle__StyleHint(111)
	QStyle__SH_ItemView_ScrollMode                            QStyle__StyleHint = QStyle__StyleHint(112)
	QStyle__SH_TitleBar_ShowToolTipsOnButtons                 QStyle__StyleHint = QStyle__StyleHint(113)
	QStyle__SH_Widget_Animation_Duration                      QStyle__StyleHint = QStyle__StyleHint(114)
	QStyle__SH_ComboBox_AllowWheelScrolling                   QStyle__StyleHint = QStyle__StyleHint(115)
	QStyle__SH_SpinBox_ButtonsInsideFrame                     QStyle__StyleHint = QStyle__StyleHint(116)
	QStyle__SH_SpinBox_StepModifier                           QStyle__StyleHint = QStyle__StyleHint(117)
	QStyle__SH_CustomBase                                     QStyle__StyleHint = QStyle__StyleHint(0xf0000000)
)

// QStyle::SubControl
//
//go:generate stringer -type=QStyle__SubControl
type QStyle__SubControl int64

const (
	QStyle__SC_None                      QStyle__SubControl = QStyle__SubControl(0x00000000)
	QStyle__SC_ScrollBarAddLine          QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_ScrollBarSubLine          QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_ScrollBarAddPage          QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_ScrollBarSubPage          QStyle__SubControl = QStyle__SubControl(0x00000008)
	QStyle__SC_ScrollBarFirst            QStyle__SubControl = QStyle__SubControl(0x00000010)
	QStyle__SC_ScrollBarLast             QStyle__SubControl = QStyle__SubControl(0x00000020)
	QStyle__SC_ScrollBarSlider           QStyle__SubControl = QStyle__SubControl(0x00000040)
	QStyle__SC_ScrollBarGroove           QStyle__SubControl = QStyle__SubControl(0x00000080)
	QStyle__SC_SpinBoxUp                 QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_SpinBoxDown               QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_SpinBoxFrame              QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_SpinBoxEditField          QStyle__SubControl = QStyle__SubControl(0x00000008)
	QStyle__SC_ComboBoxFrame             QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_ComboBoxEditField         QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_ComboBoxArrow             QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_ComboBoxListBoxPopup      QStyle__SubControl = QStyle__SubControl(0x00000008)
	QStyle__SC_SliderGroove              QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_SliderHandle              QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_SliderTickmarks           QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_ToolButton                QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_ToolButtonMenu            QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_TitleBarSysMenu           QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_TitleBarMinButton         QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_TitleBarMaxButton         QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_TitleBarCloseButton       QStyle__SubControl = QStyle__SubControl(0x00000008)
	QStyle__SC_TitleBarNormalButton      QStyle__SubControl = QStyle__SubControl(0x00000010)
	QStyle__SC_TitleBarShadeButton       QStyle__SubControl = QStyle__SubControl(0x00000020)
	QStyle__SC_TitleBarUnshadeButton     QStyle__SubControl = QStyle__SubControl(0x00000040)
	QStyle__SC_TitleBarContextHelpButton QStyle__SubControl = QStyle__SubControl(0x00000080)
	QStyle__SC_TitleBarLabel             QStyle__SubControl = QStyle__SubControl(0x00000100)
	QStyle__SC_DialGroove                QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_DialHandle                QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_DialTickmarks             QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_GroupBoxCheckBox          QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_GroupBoxLabel             QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_GroupBoxContents          QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_GroupBoxFrame             QStyle__SubControl = QStyle__SubControl(0x00000008)
	QStyle__SC_MdiMinButton              QStyle__SubControl = QStyle__SubControl(0x00000001)
	QStyle__SC_MdiNormalButton           QStyle__SubControl = QStyle__SubControl(0x00000002)
	QStyle__SC_MdiCloseButton            QStyle__SubControl = QStyle__SubControl(0x00000004)
	QStyle__SC_CustomBase                QStyle__SubControl = QStyle__SubControl(0xf0000000)
	QStyle__SC_All                       QStyle__SubControl = QStyle__SubControl(0xffffffff)
)

// QStyle::SubElement
//
//go:generate stringer -type=QStyle__SubElement
type QStyle__SubElement int64

var (
	QStyle__SE_PushButtonContents         QStyle__SubElement = QStyle__SubElement(0)
	QStyle__SE_PushButtonFocusRect        QStyle__SubElement = QStyle__SubElement(1)
	QStyle__SE_CheckBoxIndicator          QStyle__SubElement = QStyle__SubElement(2)
	QStyle__SE_CheckBoxContents           QStyle__SubElement = QStyle__SubElement(3)
	QStyle__SE_CheckBoxFocusRect          QStyle__SubElement = QStyle__SubElement(4)
	QStyle__SE_CheckBoxClickRect          QStyle__SubElement = QStyle__SubElement(5)
	QStyle__SE_RadioButtonIndicator       QStyle__SubElement = QStyle__SubElement(6)
	QStyle__SE_RadioButtonContents        QStyle__SubElement = QStyle__SubElement(7)
	QStyle__SE_RadioButtonFocusRect       QStyle__SubElement = QStyle__SubElement(8)
	QStyle__SE_RadioButtonClickRect       QStyle__SubElement = QStyle__SubElement(9)
	QStyle__SE_ComboBoxFocusRect          QStyle__SubElement = QStyle__SubElement(10)
	QStyle__SE_SliderFocusRect            QStyle__SubElement = QStyle__SubElement(11)
	QStyle__SE_ProgressBarGroove          QStyle__SubElement = QStyle__SubElement(12)
	QStyle__SE_ProgressBarContents        QStyle__SubElement = QStyle__SubElement(13)
	QStyle__SE_ProgressBarLabel           QStyle__SubElement = QStyle__SubElement(14)
	QStyle__SE_ToolBoxTabContents         QStyle__SubElement = QStyle__SubElement(15)
	QStyle__SE_HeaderLabel                QStyle__SubElement = QStyle__SubElement(16)
	QStyle__SE_HeaderArrow                QStyle__SubElement = QStyle__SubElement(17)
	QStyle__SE_TabWidgetTabBar            QStyle__SubElement = QStyle__SubElement(18)
	QStyle__SE_TabWidgetTabPane           QStyle__SubElement = QStyle__SubElement(19)
	QStyle__SE_TabWidgetTabContents       QStyle__SubElement = QStyle__SubElement(20)
	QStyle__SE_TabWidgetLeftCorner        QStyle__SubElement = QStyle__SubElement(21)
	QStyle__SE_TabWidgetRightCorner       QStyle__SubElement = QStyle__SubElement(22)
	QStyle__SE_ItemViewItemCheckIndicator QStyle__SubElement = QStyle__SubElement(23)
	QStyle__SE_ViewItemCheckIndicator     QStyle__SubElement = QStyle__SubElement(QStyle__SE_ItemViewItemCheckIndicator)
	QStyle__SE_TabBarTearIndicator        QStyle__SubElement = QStyle__SubElement(24)
	QStyle__SE_TabBarTearIndicatorLeft    QStyle__SubElement = QStyle__SubElement(QStyle__SE_TabBarTearIndicator)
	QStyle__SE_TreeViewDisclosureItem     QStyle__SubElement = QStyle__SubElement(25)
	QStyle__SE_LineEditContents           QStyle__SubElement = QStyle__SubElement(26)
	QStyle__SE_FrameContents              QStyle__SubElement = QStyle__SubElement(27)
	QStyle__SE_DockWidgetCloseButton      QStyle__SubElement = QStyle__SubElement(28)
	QStyle__SE_DockWidgetFloatButton      QStyle__SubElement = QStyle__SubElement(29)
	QStyle__SE_DockWidgetTitleBarText     QStyle__SubElement = QStyle__SubElement(30)
	QStyle__SE_DockWidgetIcon             QStyle__SubElement = QStyle__SubElement(31)
	QStyle__SE_CheckBoxLayoutItem         QStyle__SubElement = QStyle__SubElement(32)
	QStyle__SE_ComboBoxLayoutItem         QStyle__SubElement = QStyle__SubElement(33)
	QStyle__SE_DateTimeEditLayoutItem     QStyle__SubElement = QStyle__SubElement(34)
	QStyle__SE_DialogButtonBoxLayoutItem  QStyle__SubElement = QStyle__SubElement(35)
	QStyle__SE_LabelLayoutItem            QStyle__SubElement = QStyle__SubElement(C.QStyle_SE_LabelLayoutItem_Type())
	QStyle__SE_ProgressBarLayoutItem      QStyle__SubElement = QStyle__SubElement(37)
	QStyle__SE_PushButtonLayoutItem       QStyle__SubElement = QStyle__SubElement(38)
	QStyle__SE_RadioButtonLayoutItem      QStyle__SubElement = QStyle__SubElement(39)
	QStyle__SE_SliderLayoutItem           QStyle__SubElement = QStyle__SubElement(40)
	QStyle__SE_SpinBoxLayoutItem          QStyle__SubElement = QStyle__SubElement(41)
	QStyle__SE_ToolButtonLayoutItem       QStyle__SubElement = QStyle__SubElement(42)
	QStyle__SE_FrameLayoutItem            QStyle__SubElement = QStyle__SubElement(43)
	QStyle__SE_GroupBoxLayoutItem         QStyle__SubElement = QStyle__SubElement(44)
	QStyle__SE_TabWidgetLayoutItem        QStyle__SubElement = QStyle__SubElement(45)
	QStyle__SE_ItemViewItemDecoration     QStyle__SubElement = QStyle__SubElement(46)
	QStyle__SE_ItemViewItemText           QStyle__SubElement = QStyle__SubElement(47)
	QStyle__SE_ItemViewItemFocusRect      QStyle__SubElement = QStyle__SubElement(48)
	QStyle__SE_TabBarTabLeftButton        QStyle__SubElement = QStyle__SubElement(49)
	QStyle__SE_TabBarTabRightButton       QStyle__SubElement = QStyle__SubElement(50)
	QStyle__SE_TabBarTabText              QStyle__SubElement = QStyle__SubElement(51)
	QStyle__SE_ShapedFrameContents        QStyle__SubElement = QStyle__SubElement(52)
	QStyle__SE_ToolBarHandle              QStyle__SubElement = QStyle__SubElement(53)
	QStyle__SE_TabBarScrollLeftButton     QStyle__SubElement = QStyle__SubElement(54)
	QStyle__SE_TabBarScrollRightButton    QStyle__SubElement = QStyle__SubElement(55)
	QStyle__SE_TabBarTearIndicatorRight   QStyle__SubElement = QStyle__SubElement(56)
	QStyle__SE_PushButtonBevel            QStyle__SubElement = QStyle__SubElement(57)
	QStyle__SE_CustomBase                 QStyle__SubElement = QStyle__SubElement(0xf0000000)
)

func NewQStyle2() *QStyle {
	tmpValue := NewQStyleFromPointer(C.QStyle_NewQStyle2())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQStyle_DrawComplexControl
func callbackQStyle_DrawComplexControl(ptr unsafe.Pointer, control C.longlong, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawComplexControl"); signal != nil {
		(*(*func(QStyle__ComplexControl, *QStyleOptionComplex, *gui.QPainter, *QWidget))(signal))(QStyle__ComplexControl(control), NewQStyleOptionComplexFromPointer(option), gui.NewQPainterFromPointer(painter), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QStyle) ConnectDrawComplexControl(f func(control QStyle__ComplexControl, option *QStyleOptionComplex, painter *gui.QPainter, widget *QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "drawComplexControl"); signal != nil {
			f := func(control QStyle__ComplexControl, option *QStyleOptionComplex, painter *gui.QPainter, widget *QWidget) {
				(*(*func(QStyle__ComplexControl, *QStyleOptionComplex, *gui.QPainter, *QWidget))(signal))(control, option, painter, widget)
				f(control, option, painter, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "drawComplexControl", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "drawComplexControl", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectDrawComplexControl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "drawComplexControl")
	}
}

func (ptr *QStyle) DrawComplexControl(control QStyle__ComplexControl, option QStyleOptionComplex_ITF, painter gui.QPainter_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle_DrawComplexControl(ptr.Pointer(), C.longlong(control), PointerFromQStyleOptionComplex(option), gui.PointerFromQPainter(painter), PointerFromQWidget(widget))
	}
}

//export callbackQStyle_DrawControl
func callbackQStyle_DrawControl(ptr unsafe.Pointer, element C.longlong, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawControl"); signal != nil {
		(*(*func(QStyle__ControlElement, *QStyleOption, *gui.QPainter, *QWidget))(signal))(QStyle__ControlElement(element), NewQStyleOptionFromPointer(option), gui.NewQPainterFromPointer(painter), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QStyle) ConnectDrawControl(f func(element QStyle__ControlElement, option *QStyleOption, painter *gui.QPainter, widget *QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "drawControl"); signal != nil {
			f := func(element QStyle__ControlElement, option *QStyleOption, painter *gui.QPainter, widget *QWidget) {
				(*(*func(QStyle__ControlElement, *QStyleOption, *gui.QPainter, *QWidget))(signal))(element, option, painter, widget)
				f(element, option, painter, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "drawControl", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "drawControl", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectDrawControl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "drawControl")
	}
}

func (ptr *QStyle) DrawControl(element QStyle__ControlElement, option QStyleOption_ITF, painter gui.QPainter_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle_DrawControl(ptr.Pointer(), C.longlong(element), PointerFromQStyleOption(option), gui.PointerFromQPainter(painter), PointerFromQWidget(widget))
	}
}

//export callbackQStyle_DrawPrimitive
func callbackQStyle_DrawPrimitive(ptr unsafe.Pointer, element C.longlong, option unsafe.Pointer, painter unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawPrimitive"); signal != nil {
		(*(*func(QStyle__PrimitiveElement, *QStyleOption, *gui.QPainter, *QWidget))(signal))(QStyle__PrimitiveElement(element), NewQStyleOptionFromPointer(option), gui.NewQPainterFromPointer(painter), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QStyle) ConnectDrawPrimitive(f func(element QStyle__PrimitiveElement, option *QStyleOption, painter *gui.QPainter, widget *QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "drawPrimitive"); signal != nil {
			f := func(element QStyle__PrimitiveElement, option *QStyleOption, painter *gui.QPainter, widget *QWidget) {
				(*(*func(QStyle__PrimitiveElement, *QStyleOption, *gui.QPainter, *QWidget))(signal))(element, option, painter, widget)
				f(element, option, painter, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "drawPrimitive", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "drawPrimitive", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectDrawPrimitive() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "drawPrimitive")
	}
}

func (ptr *QStyle) DrawPrimitive(element QStyle__PrimitiveElement, option QStyleOption_ITF, painter gui.QPainter_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle_DrawPrimitive(ptr.Pointer(), C.longlong(element), PointerFromQStyleOption(option), gui.PointerFromQPainter(painter), PointerFromQWidget(widget))
	}
}

//export callbackQStyle_GeneratedIconPixmap
func callbackQStyle_GeneratedIconPixmap(ptr unsafe.Pointer, iconMode C.longlong, pixmap unsafe.Pointer, option unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "generatedIconPixmap"); signal != nil {
		return gui.PointerFromQPixmap((*(*func(gui.QIcon__Mode, *gui.QPixmap, *QStyleOption) *gui.QPixmap)(signal))(gui.QIcon__Mode(iconMode), gui.NewQPixmapFromPointer(pixmap), NewQStyleOptionFromPointer(option)))
	}

	return gui.PointerFromQPixmap(gui.NewQPixmap())
}

func (ptr *QStyle) ConnectGeneratedIconPixmap(f func(iconMode gui.QIcon__Mode, pixmap *gui.QPixmap, option *QStyleOption) *gui.QPixmap) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "generatedIconPixmap"); signal != nil {
			f := func(iconMode gui.QIcon__Mode, pixmap *gui.QPixmap, option *QStyleOption) *gui.QPixmap {
				(*(*func(gui.QIcon__Mode, *gui.QPixmap, *QStyleOption) *gui.QPixmap)(signal))(iconMode, pixmap, option)
				return f(iconMode, pixmap, option)
			}
			qt.ConnectSignal(ptr.Pointer(), "generatedIconPixmap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "generatedIconPixmap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectGeneratedIconPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "generatedIconPixmap")
	}
}

func (ptr *QStyle) GeneratedIconPixmap(iconMode gui.QIcon__Mode, pixmap gui.QPixmap_ITF, option QStyleOption_ITF) *gui.QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPixmapFromPointer(C.QStyle_GeneratedIconPixmap(ptr.Pointer(), C.longlong(iconMode), gui.PointerFromQPixmap(pixmap), PointerFromQStyleOption(option)))
		qt.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

//export callbackQStyle_HitTestComplexControl
func callbackQStyle_HitTestComplexControl(ptr unsafe.Pointer, control C.longlong, option unsafe.Pointer, position unsafe.Pointer, widget unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "hitTestComplexControl"); signal != nil {
		return C.longlong((*(*func(QStyle__ComplexControl, *QStyleOptionComplex, *core.QPoint, *QWidget) QStyle__SubControl)(signal))(QStyle__ComplexControl(control), NewQStyleOptionComplexFromPointer(option), core.NewQPointFromPointer(position), NewQWidgetFromPointer(widget)))
	}

	return C.longlong(0)
}

func (ptr *QStyle) ConnectHitTestComplexControl(f func(control QStyle__ComplexControl, option *QStyleOptionComplex, position *core.QPoint, widget *QWidget) QStyle__SubControl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hitTestComplexControl"); signal != nil {
			f := func(control QStyle__ComplexControl, option *QStyleOptionComplex, position *core.QPoint, widget *QWidget) QStyle__SubControl {
				(*(*func(QStyle__ComplexControl, *QStyleOptionComplex, *core.QPoint, *QWidget) QStyle__SubControl)(signal))(control, option, position, widget)
				return f(control, option, position, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "hitTestComplexControl", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hitTestComplexControl", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectHitTestComplexControl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hitTestComplexControl")
	}
}

func (ptr *QStyle) HitTestComplexControl(control QStyle__ComplexControl, option QStyleOptionComplex_ITF, position core.QPoint_ITF, widget QWidget_ITF) QStyle__SubControl {
	if ptr.Pointer() != nil {
		return QStyle__SubControl(C.QStyle_HitTestComplexControl(ptr.Pointer(), C.longlong(control), PointerFromQStyleOptionComplex(option), core.PointerFromQPoint(position), PointerFromQWidget(widget)))
	}
	return 0
}

//export callbackQStyle_LayoutSpacing
func callbackQStyle_LayoutSpacing(ptr unsafe.Pointer, control1 C.longlong, control2 C.longlong, orientation C.longlong, option unsafe.Pointer, widget unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "layoutSpacing"); signal != nil {
		return C.int(int32((*(*func(QSizePolicy__ControlType, QSizePolicy__ControlType, core.Qt__Orientation, *QStyleOption, *QWidget) int)(signal))(QSizePolicy__ControlType(control1), QSizePolicy__ControlType(control2), core.Qt__Orientation(orientation), NewQStyleOptionFromPointer(option), NewQWidgetFromPointer(widget))))
	}

	return C.int(int32(0))
}

func (ptr *QStyle) ConnectLayoutSpacing(f func(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option *QStyleOption, widget *QWidget) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "layoutSpacing"); signal != nil {
			f := func(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option *QStyleOption, widget *QWidget) int {
				(*(*func(QSizePolicy__ControlType, QSizePolicy__ControlType, core.Qt__Orientation, *QStyleOption, *QWidget) int)(signal))(control1, control2, orientation, option, widget)
				return f(control1, control2, orientation, option, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "layoutSpacing", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "layoutSpacing", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectLayoutSpacing() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "layoutSpacing")
	}
}

func (ptr *QStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option QStyleOption_ITF, widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStyle_LayoutSpacing(ptr.Pointer(), C.longlong(control1), C.longlong(control2), C.longlong(orientation), PointerFromQStyleOption(option), PointerFromQWidget(widget))))
	}
	return 0
}

//export callbackQStyle_PixelMetric
func callbackQStyle_PixelMetric(ptr unsafe.Pointer, metric C.longlong, option unsafe.Pointer, widget unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "pixelMetric"); signal != nil {
		return C.int(int32((*(*func(QStyle__PixelMetric, *QStyleOption, *QWidget) int)(signal))(QStyle__PixelMetric(metric), NewQStyleOptionFromPointer(option), NewQWidgetFromPointer(widget))))
	}

	return C.int(int32(0))
}

func (ptr *QStyle) ConnectPixelMetric(f func(metric QStyle__PixelMetric, option *QStyleOption, widget *QWidget) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "pixelMetric"); signal != nil {
			f := func(metric QStyle__PixelMetric, option *QStyleOption, widget *QWidget) int {
				(*(*func(QStyle__PixelMetric, *QStyleOption, *QWidget) int)(signal))(metric, option, widget)
				return f(metric, option, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "pixelMetric", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pixelMetric", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectPixelMetric() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "pixelMetric")
	}
}

func (ptr *QStyle) PixelMetric(metric QStyle__PixelMetric, option QStyleOption_ITF, widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStyle_PixelMetric(ptr.Pointer(), C.longlong(metric), PointerFromQStyleOption(option), PointerFromQWidget(widget))))
	}
	return 0
}

//export callbackQStyle_SizeFromContents
func callbackQStyle_SizeFromContents(ptr unsafe.Pointer, ty C.longlong, option unsafe.Pointer, contentsSize unsafe.Pointer, widget unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeFromContents"); signal != nil {
		return core.PointerFromQSize((*(*func(QStyle__ContentsType, *QStyleOption, *core.QSize, *QWidget) *core.QSize)(signal))(QStyle__ContentsType(ty), NewQStyleOptionFromPointer(option), core.NewQSizeFromPointer(contentsSize), NewQWidgetFromPointer(widget)))
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QStyle) ConnectSizeFromContents(f func(ty QStyle__ContentsType, option *QStyleOption, contentsSize *core.QSize, widget *QWidget) *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sizeFromContents"); signal != nil {
			f := func(ty QStyle__ContentsType, option *QStyleOption, contentsSize *core.QSize, widget *QWidget) *core.QSize {
				(*(*func(QStyle__ContentsType, *QStyleOption, *core.QSize, *QWidget) *core.QSize)(signal))(ty, option, contentsSize, widget)
				return f(ty, option, contentsSize, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "sizeFromContents", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sizeFromContents", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectSizeFromContents() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sizeFromContents")
	}
}

func (ptr *QStyle) SizeFromContents(ty QStyle__ContentsType, option QStyleOption_ITF, contentsSize core.QSize_ITF, widget QWidget_ITF) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QStyle_SizeFromContents(ptr.Pointer(), C.longlong(ty), PointerFromQStyleOption(option), core.PointerFromQSize(contentsSize), PointerFromQWidget(widget)))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQStyle_StandardIcon
func callbackQStyle_StandardIcon(ptr unsafe.Pointer, standardIcon C.longlong, option unsafe.Pointer, widget unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "standardIcon"); signal != nil {
		return gui.PointerFromQIcon((*(*func(QStyle__StandardPixmap, *QStyleOption, *QWidget) *gui.QIcon)(signal))(QStyle__StandardPixmap(standardIcon), NewQStyleOptionFromPointer(option), NewQWidgetFromPointer(widget)))
	}

	return gui.PointerFromQIcon(gui.NewQIcon())
}

func (ptr *QStyle) ConnectStandardIcon(f func(standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *gui.QIcon) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "standardIcon"); signal != nil {
			f := func(standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *gui.QIcon {
				(*(*func(QStyle__StandardPixmap, *QStyleOption, *QWidget) *gui.QIcon)(signal))(standardIcon, option, widget)
				return f(standardIcon, option, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "standardIcon", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "standardIcon", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectStandardIcon() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "standardIcon")
	}
}

func (ptr *QStyle) StandardIcon(standardIcon QStyle__StandardPixmap, option QStyleOption_ITF, widget QWidget_ITF) *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QStyle_StandardIcon(ptr.Pointer(), C.longlong(standardIcon), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
		qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

//export callbackQStyle_StyleHint
func callbackQStyle_StyleHint(ptr unsafe.Pointer, hint C.longlong, option unsafe.Pointer, widget unsafe.Pointer, returnData unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "styleHint"); signal != nil {
		return C.int(int32((*(*func(QStyle__StyleHint, *QStyleOption, *QWidget, *QStyleHintReturn) int)(signal))(QStyle__StyleHint(hint), NewQStyleOptionFromPointer(option), NewQWidgetFromPointer(widget), NewQStyleHintReturnFromPointer(returnData))))
	}

	return C.int(int32(0))
}

func (ptr *QStyle) ConnectStyleHint(f func(hint QStyle__StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "styleHint"); signal != nil {
			f := func(hint QStyle__StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int {
				(*(*func(QStyle__StyleHint, *QStyleOption, *QWidget, *QStyleHintReturn) int)(signal))(hint, option, widget, returnData)
				return f(hint, option, widget, returnData)
			}
			qt.ConnectSignal(ptr.Pointer(), "styleHint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "styleHint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectStyleHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "styleHint")
	}
}

func (ptr *QStyle) StyleHint(hint QStyle__StyleHint, option QStyleOption_ITF, widget QWidget_ITF, returnData QStyleHintReturn_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStyle_StyleHint(ptr.Pointer(), C.longlong(hint), PointerFromQStyleOption(option), PointerFromQWidget(widget), PointerFromQStyleHintReturn(returnData))))
	}
	return 0
}

//export callbackQStyle_SubControlRect
func callbackQStyle_SubControlRect(ptr unsafe.Pointer, control C.longlong, option unsafe.Pointer, subControl C.longlong, widget unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "subControlRect"); signal != nil {
		return core.PointerFromQRect((*(*func(QStyle__ComplexControl, *QStyleOptionComplex, QStyle__SubControl, *QWidget) *core.QRect)(signal))(QStyle__ComplexControl(control), NewQStyleOptionComplexFromPointer(option), QStyle__SubControl(subControl), NewQWidgetFromPointer(widget)))
	}

	return core.PointerFromQRect(core.NewQRect())
}

func (ptr *QStyle) ConnectSubControlRect(f func(control QStyle__ComplexControl, option *QStyleOptionComplex, subControl QStyle__SubControl, widget *QWidget) *core.QRect) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "subControlRect"); signal != nil {
			f := func(control QStyle__ComplexControl, option *QStyleOptionComplex, subControl QStyle__SubControl, widget *QWidget) *core.QRect {
				(*(*func(QStyle__ComplexControl, *QStyleOptionComplex, QStyle__SubControl, *QWidget) *core.QRect)(signal))(control, option, subControl, widget)
				return f(control, option, subControl, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "subControlRect", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "subControlRect", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectSubControlRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "subControlRect")
	}
}

func (ptr *QStyle) SubControlRect(control QStyle__ComplexControl, option QStyleOptionComplex_ITF, subControl QStyle__SubControl, widget QWidget_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QStyle_SubControlRect(ptr.Pointer(), C.longlong(control), PointerFromQStyleOptionComplex(option), C.longlong(subControl), PointerFromQWidget(widget)))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQStyle_SubElementRect
func callbackQStyle_SubElementRect(ptr unsafe.Pointer, element C.longlong, option unsafe.Pointer, widget unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "subElementRect"); signal != nil {
		return core.PointerFromQRect((*(*func(QStyle__SubElement, *QStyleOption, *QWidget) *core.QRect)(signal))(QStyle__SubElement(element), NewQStyleOptionFromPointer(option), NewQWidgetFromPointer(widget)))
	}

	return core.PointerFromQRect(core.NewQRect())
}

func (ptr *QStyle) ConnectSubElementRect(f func(element QStyle__SubElement, option *QStyleOption, widget *QWidget) *core.QRect) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "subElementRect"); signal != nil {
			f := func(element QStyle__SubElement, option *QStyleOption, widget *QWidget) *core.QRect {
				(*(*func(QStyle__SubElement, *QStyleOption, *QWidget) *core.QRect)(signal))(element, option, widget)
				return f(element, option, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "subElementRect", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "subElementRect", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectSubElementRect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "subElementRect")
	}
}

func (ptr *QStyle) SubElementRect(element QStyle__SubElement, option QStyleOption_ITF, widget QWidget_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QStyle_SubElementRect(ptr.Pointer(), C.longlong(element), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQStyle_DestroyQStyle
func callbackQStyle_DestroyQStyle(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QStyle"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQStyleFromPointer(ptr).DestroyQStyleDefault()
	}
}

func (ptr *QStyle) ConnectDestroyQStyle(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QStyle"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QStyle", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QStyle", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectDestroyQStyle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QStyle")
	}
}

func (ptr *QStyle) DestroyQStyle() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStyle_DestroyQStyle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStyle) DestroyQStyleDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStyle_DestroyQStyleDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQStyle_StandardPixmap
func callbackQStyle_StandardPixmap(ptr unsafe.Pointer, standardIcon C.longlong, option unsafe.Pointer, widget unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "standardPixmap"); signal != nil {
		return gui.PointerFromQPixmap((*(*func(QStyle__StandardPixmap, *QStyleOption, *QWidget) *gui.QPixmap)(signal))(QStyle__StandardPixmap(standardIcon), NewQStyleOptionFromPointer(option), NewQWidgetFromPointer(widget)))
	}

	return gui.PointerFromQPixmap(gui.NewQPixmap())
}

func (ptr *QStyle) ConnectStandardPixmap(f func(standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *gui.QPixmap) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "standardPixmap"); signal != nil {
			f := func(standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *gui.QPixmap {
				(*(*func(QStyle__StandardPixmap, *QStyleOption, *QWidget) *gui.QPixmap)(signal))(standardIcon, option, widget)
				return f(standardIcon, option, widget)
			}
			qt.ConnectSignal(ptr.Pointer(), "standardPixmap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "standardPixmap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStyle) DisconnectStandardPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "standardPixmap")
	}
}

func (ptr *QStyle) StandardPixmap(standardIcon QStyle__StandardPixmap, option QStyleOption_ITF, widget QWidget_ITF) *gui.QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPixmapFromPointer(C.QStyle_StandardPixmap(ptr.Pointer(), C.longlong(standardIcon), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
		qt.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QStyle) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QStyle___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QStyle) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QStyle) __children_newList() unsafe.Pointer {
	return C.QStyle___children_newList(ptr.Pointer())
}

func (ptr *QStyle) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QStyle___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QStyle) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QStyle) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QStyle___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QStyle) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QStyle___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QStyle) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QStyle) __findChildren_newList() unsafe.Pointer {
	return C.QStyle___findChildren_newList(ptr.Pointer())
}

func (ptr *QStyle) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QStyle___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QStyle) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QStyle) __findChildren_newList3() unsafe.Pointer {
	return C.QStyle___findChildren_newList3(ptr.Pointer())
}

//export callbackQStyle_ChildEvent
func callbackQStyle_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQStyleFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QStyle) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQStyle_ConnectNotify
func callbackQStyle_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQStyleFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QStyle) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQStyle_CustomEvent
func callbackQStyle_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQStyleFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStyle) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQStyle_DeleteLater
func callbackQStyle_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQStyleFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QStyle) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStyle_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQStyle_Destroyed
func callbackQStyle_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQStyle_DisconnectNotify
func callbackQStyle_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQStyleFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QStyle) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQStyle_Event
func callbackQStyle_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQStyleFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QStyle) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QStyle_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQStyle_EventFilter
func callbackQStyle_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQStyleFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QStyle) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QStyle_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQStyle_ObjectNameChanged
func callbackQStyle_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQStyle_TimerEvent
func callbackQStyle_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQStyleFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QStyle) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QStyle_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QStyleHintReturn struct {
	ptr unsafe.Pointer
}

type QStyleHintReturn_ITF interface {
	QStyleHintReturn_PTR() *QStyleHintReturn
}

func (ptr *QStyleHintReturn) QStyleHintReturn_PTR() *QStyleHintReturn {
	return ptr
}

func (ptr *QStyleHintReturn) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QStyleHintReturn) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQStyleHintReturn(ptr QStyleHintReturn_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleHintReturn_PTR().Pointer()
	}
	return nil
}

func NewQStyleHintReturnFromPointer(ptr unsafe.Pointer) (n *QStyleHintReturn) {
	n = new(QStyleHintReturn)
	n.SetPointer(ptr)
	return
}
func (ptr *QStyleHintReturn) DestroyQStyleHintReturn() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QStyleHintReturn__HintReturnType
//QStyleHintReturn::HintReturnType
type QStyleHintReturn__HintReturnType int64

const (
	QStyleHintReturn__SH_Default QStyleHintReturn__HintReturnType = QStyleHintReturn__HintReturnType(0xf000)
	QStyleHintReturn__SH_Mask    QStyleHintReturn__HintReturnType = QStyleHintReturn__HintReturnType(0xf001)
	QStyleHintReturn__SH_Variant QStyleHintReturn__HintReturnType = QStyleHintReturn__HintReturnType(0xf002)
)

// QStyleHintReturn::StyleOptionType
//
//go:generate stringer -type=QStyleHintReturn__StyleOptionType
type QStyleHintReturn__StyleOptionType int64

var (
	QStyleHintReturn__Type QStyleHintReturn__StyleOptionType = QStyleHintReturn__StyleOptionType(QStyleHintReturn__SH_Default)
)

// QStyleHintReturn::StyleOptionVersion
//
//go:generate stringer -type=QStyleHintReturn__StyleOptionVersion
type QStyleHintReturn__StyleOptionVersion int64

var (
	QStyleHintReturn__Version QStyleHintReturn__StyleOptionVersion = QStyleHintReturn__StyleOptionVersion(1)
)

func NewQStyleHintReturn(version int, ty int) *QStyleHintReturn {
	tmpValue := NewQStyleHintReturnFromPointer(C.QStyleHintReturn_NewQStyleHintReturn(C.int(int32(version)), C.int(int32(ty))))
	qt.SetFinalizer(tmpValue, (*QStyleHintReturn).DestroyQStyleHintReturn)
	return tmpValue
}

func (ptr *QStyleHintReturn) Type() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStyleHintReturn_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleHintReturn) Version() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStyleHintReturn_Version(ptr.Pointer())))
	}
	return 0
}

// QStyleHintReturnMask::StyleOptionType
//
//go:generate stringer -type=QStyleHintReturnMask__StyleOptionType
type QStyleHintReturnMask__StyleOptionType int64

var (
	QStyleHintReturnMask__Type QStyleHintReturnMask__StyleOptionType = QStyleHintReturnMask__StyleOptionType(QStyleHintReturn__SH_Mask)
)

// QStyleHintReturnMask::StyleOptionVersion
//
//go:generate stringer -type=QStyleHintReturnMask__StyleOptionVersion
type QStyleHintReturnMask__StyleOptionVersion int64

var (
	QStyleHintReturnMask__Version QStyleHintReturnMask__StyleOptionVersion = QStyleHintReturnMask__StyleOptionVersion(1)
)

// QStyleHintReturnVariant::StyleOptionType
//
//go:generate stringer -type=QStyleHintReturnVariant__StyleOptionType
type QStyleHintReturnVariant__StyleOptionType int64

var (
	QStyleHintReturnVariant__Type QStyleHintReturnVariant__StyleOptionType = QStyleHintReturnVariant__StyleOptionType(QStyleHintReturn__SH_Variant)
)

// QStyleHintReturnVariant::StyleOptionVersion
//
//go:generate stringer -type=QStyleHintReturnVariant__StyleOptionVersion
type QStyleHintReturnVariant__StyleOptionVersion int64

var (
	QStyleHintReturnVariant__Version QStyleHintReturnVariant__StyleOptionVersion = QStyleHintReturnVariant__StyleOptionVersion(1)
)

type QStyleOption struct {
	ptr unsafe.Pointer
}

type QStyleOption_ITF interface {
	QStyleOption_PTR() *QStyleOption
}

func (ptr *QStyleOption) QStyleOption_PTR() *QStyleOption {
	return ptr
}

func (ptr *QStyleOption) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QStyleOption) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQStyleOption(ptr QStyleOption_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOption_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionFromPointer(ptr unsafe.Pointer) (n *QStyleOption) {
	n = new(QStyleOption)
	n.SetPointer(ptr)
	return
}

// QStyleOption::OptionType
//
//go:generate stringer -type=QStyleOption__OptionType
type QStyleOption__OptionType int64

const (
	QStyleOption__SO_Default           QStyleOption__OptionType = QStyleOption__OptionType(0)
	QStyleOption__SO_FocusRect         QStyleOption__OptionType = QStyleOption__OptionType(1)
	QStyleOption__SO_Button            QStyleOption__OptionType = QStyleOption__OptionType(2)
	QStyleOption__SO_Tab               QStyleOption__OptionType = QStyleOption__OptionType(3)
	QStyleOption__SO_MenuItem          QStyleOption__OptionType = QStyleOption__OptionType(4)
	QStyleOption__SO_Frame             QStyleOption__OptionType = QStyleOption__OptionType(5)
	QStyleOption__SO_ProgressBar       QStyleOption__OptionType = QStyleOption__OptionType(6)
	QStyleOption__SO_ToolBox           QStyleOption__OptionType = QStyleOption__OptionType(7)
	QStyleOption__SO_Header            QStyleOption__OptionType = QStyleOption__OptionType(8)
	QStyleOption__SO_DockWidget        QStyleOption__OptionType = QStyleOption__OptionType(9)
	QStyleOption__SO_ViewItem          QStyleOption__OptionType = QStyleOption__OptionType(10)
	QStyleOption__SO_TabWidgetFrame    QStyleOption__OptionType = QStyleOption__OptionType(11)
	QStyleOption__SO_TabBarBase        QStyleOption__OptionType = QStyleOption__OptionType(12)
	QStyleOption__SO_RubberBand        QStyleOption__OptionType = QStyleOption__OptionType(13)
	QStyleOption__SO_ToolBar           QStyleOption__OptionType = QStyleOption__OptionType(14)
	QStyleOption__SO_GraphicsItem      QStyleOption__OptionType = QStyleOption__OptionType(15)
	QStyleOption__SO_Complex           QStyleOption__OptionType = QStyleOption__OptionType(0xf0000)
	QStyleOption__SO_Slider            QStyleOption__OptionType = QStyleOption__OptionType(0xf0001)
	QStyleOption__SO_SpinBox           QStyleOption__OptionType = QStyleOption__OptionType(0xf0002)
	QStyleOption__SO_ToolButton        QStyleOption__OptionType = QStyleOption__OptionType(0xf0003)
	QStyleOption__SO_ComboBox          QStyleOption__OptionType = QStyleOption__OptionType(0xf0004)
	QStyleOption__SO_TitleBar          QStyleOption__OptionType = QStyleOption__OptionType(0xf0005)
	QStyleOption__SO_GroupBox          QStyleOption__OptionType = QStyleOption__OptionType(0xf0006)
	QStyleOption__SO_SizeGrip          QStyleOption__OptionType = QStyleOption__OptionType(0xf0007)
	QStyleOption__SO_CustomBase        QStyleOption__OptionType = QStyleOption__OptionType(0xf00)
	QStyleOption__SO_ComplexCustomBase QStyleOption__OptionType = QStyleOption__OptionType(0xf000000)
)

// QStyleOption::StyleOptionType
//
//go:generate stringer -type=QStyleOption__StyleOptionType
type QStyleOption__StyleOptionType int64

var (
	QStyleOption__Type QStyleOption__StyleOptionType = QStyleOption__StyleOptionType(QStyleOption__SO_Default)
)

// QStyleOption::StyleOptionVersion
//
//go:generate stringer -type=QStyleOption__StyleOptionVersion
type QStyleOption__StyleOptionVersion int64

var (
	QStyleOption__Version QStyleOption__StyleOptionVersion = QStyleOption__StyleOptionVersion(1)
)

func NewQStyleOption(version int, ty int) *QStyleOption {
	tmpValue := NewQStyleOptionFromPointer(C.QStyleOption_NewQStyleOption(C.int(int32(version)), C.int(int32(ty))))
	qt.SetFinalizer(tmpValue, (*QStyleOption).DestroyQStyleOption)
	return tmpValue
}

func NewQStyleOption2(other QStyleOption_ITF) *QStyleOption {
	tmpValue := NewQStyleOptionFromPointer(C.QStyleOption_NewQStyleOption2(PointerFromQStyleOption(other)))
	qt.SetFinalizer(tmpValue, (*QStyleOption).DestroyQStyleOption)
	return tmpValue
}

func (ptr *QStyleOption) DestroyQStyleOption() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStyleOption_DestroyQStyleOption(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStyleOption) Direction() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QStyleOption_Direction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStyleOption) FontMetrics() *gui.QFontMetrics {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontMetricsFromPointer(C.QStyleOption_FontMetrics(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFontMetrics).DestroyQFontMetrics)
		return tmpValue
	}
	return nil
}

func (ptr *QStyleOption) Rect() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QStyleOption_Rect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QStyleOption) Type() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStyleOption_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStyleOption) Version() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStyleOption_Version(ptr.Pointer())))
	}
	return 0
}

// QStyleOptionButton::ButtonFeature
//
//go:generate stringer -type=QStyleOptionButton__ButtonFeature
type QStyleOptionButton__ButtonFeature int64

const (
	QStyleOptionButton__None              QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x00)
	QStyleOptionButton__Flat              QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x01)
	QStyleOptionButton__HasMenu           QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x02)
	QStyleOptionButton__DefaultButton     QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x04)
	QStyleOptionButton__AutoDefaultButton QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x08)
	QStyleOptionButton__CommandLinkButton QStyleOptionButton__ButtonFeature = QStyleOptionButton__ButtonFeature(0x10)
)

// QStyleOptionButton::StyleOptionType
//
//go:generate stringer -type=QStyleOptionButton__StyleOptionType
type QStyleOptionButton__StyleOptionType int64

var (
	QStyleOptionButton__Type QStyleOptionButton__StyleOptionType = QStyleOptionButton__StyleOptionType(QStyleOption__SO_Button)
)

// QStyleOptionButton::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionButton__StyleOptionVersion
type QStyleOptionButton__StyleOptionVersion int64

var (
	QStyleOptionButton__Version QStyleOptionButton__StyleOptionVersion = QStyleOptionButton__StyleOptionVersion(1)
)

// QStyleOptionComboBox::StyleOptionType
//
//go:generate stringer -type=QStyleOptionComboBox__StyleOptionType
type QStyleOptionComboBox__StyleOptionType int64

var (
	QStyleOptionComboBox__Type QStyleOptionComboBox__StyleOptionType = QStyleOptionComboBox__StyleOptionType(QStyleOption__SO_ComboBox)
)

// QStyleOptionComboBox::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionComboBox__StyleOptionVersion
type QStyleOptionComboBox__StyleOptionVersion int64

var (
	QStyleOptionComboBox__Version QStyleOptionComboBox__StyleOptionVersion = QStyleOptionComboBox__StyleOptionVersion(1)
)

type QStyleOptionComplex struct {
	QStyleOption
}

type QStyleOptionComplex_ITF interface {
	QStyleOption_ITF
	QStyleOptionComplex_PTR() *QStyleOptionComplex
}

func (ptr *QStyleOptionComplex) QStyleOptionComplex_PTR() *QStyleOptionComplex {
	return ptr
}

func (ptr *QStyleOptionComplex) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOption_PTR().Pointer()
	}
	return nil
}

func (ptr *QStyleOptionComplex) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QStyleOption_PTR().SetPointer(p)
	}
}

func PointerFromQStyleOptionComplex(ptr QStyleOptionComplex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionComplex_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionComplexFromPointer(ptr unsafe.Pointer) (n *QStyleOptionComplex) {
	n = new(QStyleOptionComplex)
	n.SetPointer(ptr)
	return
}
func (ptr *QStyleOptionComplex) DestroyQStyleOptionComplex() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QStyleOptionComplex__StyleOptionType
//QStyleOptionComplex::StyleOptionType
type QStyleOptionComplex__StyleOptionType int64

var (
	QStyleOptionComplex__Type QStyleOptionComplex__StyleOptionType = QStyleOptionComplex__StyleOptionType(QStyleOption__SO_Complex)
)

// QStyleOptionComplex::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionComplex__StyleOptionVersion
type QStyleOptionComplex__StyleOptionVersion int64

var (
	QStyleOptionComplex__Version QStyleOptionComplex__StyleOptionVersion = QStyleOptionComplex__StyleOptionVersion(1)
)

func NewQStyleOptionComplex(version int, ty int) *QStyleOptionComplex {
	tmpValue := NewQStyleOptionComplexFromPointer(C.QStyleOptionComplex_NewQStyleOptionComplex(C.int(int32(version)), C.int(int32(ty))))
	qt.SetFinalizer(tmpValue, (*QStyleOptionComplex).DestroyQStyleOptionComplex)
	return tmpValue
}

func NewQStyleOptionComplex2(other QStyleOptionComplex_ITF) *QStyleOptionComplex {
	tmpValue := NewQStyleOptionComplexFromPointer(C.QStyleOptionComplex_NewQStyleOptionComplex2(PointerFromQStyleOptionComplex(other)))
	qt.SetFinalizer(tmpValue, (*QStyleOptionComplex).DestroyQStyleOptionComplex)
	return tmpValue
}

// QStyleOptionDockWidget::StyleOptionType
//
//go:generate stringer -type=QStyleOptionDockWidget__StyleOptionType
type QStyleOptionDockWidget__StyleOptionType int64

var (
	QStyleOptionDockWidget__Type QStyleOptionDockWidget__StyleOptionType = QStyleOptionDockWidget__StyleOptionType(QStyleOption__SO_DockWidget)
)

// QStyleOptionDockWidget::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionDockWidget__StyleOptionVersion
type QStyleOptionDockWidget__StyleOptionVersion int64

var (
	QStyleOptionDockWidget__Version QStyleOptionDockWidget__StyleOptionVersion = QStyleOptionDockWidget__StyleOptionVersion(2)
)

// QStyleOptionFocusRect::StyleOptionType
//
//go:generate stringer -type=QStyleOptionFocusRect__StyleOptionType
type QStyleOptionFocusRect__StyleOptionType int64

var (
	QStyleOptionFocusRect__Type QStyleOptionFocusRect__StyleOptionType = QStyleOptionFocusRect__StyleOptionType(QStyleOption__SO_FocusRect)
)

// QStyleOptionFocusRect::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionFocusRect__StyleOptionVersion
type QStyleOptionFocusRect__StyleOptionVersion int64

var (
	QStyleOptionFocusRect__Version QStyleOptionFocusRect__StyleOptionVersion = QStyleOptionFocusRect__StyleOptionVersion(1)
)

// QStyleOptionFrame::FrameFeature
//
//go:generate stringer -type=QStyleOptionFrame__FrameFeature
type QStyleOptionFrame__FrameFeature int64

const (
	QStyleOptionFrame__None    QStyleOptionFrame__FrameFeature = QStyleOptionFrame__FrameFeature(0x00)
	QStyleOptionFrame__Flat    QStyleOptionFrame__FrameFeature = QStyleOptionFrame__FrameFeature(0x01)
	QStyleOptionFrame__Rounded QStyleOptionFrame__FrameFeature = QStyleOptionFrame__FrameFeature(0x02)
)

// QStyleOptionFrame::StyleOptionType
//
//go:generate stringer -type=QStyleOptionFrame__StyleOptionType
type QStyleOptionFrame__StyleOptionType int64

var (
	QStyleOptionFrame__Type QStyleOptionFrame__StyleOptionType = QStyleOptionFrame__StyleOptionType(QStyleOption__SO_Frame)
)

// QStyleOptionFrame::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionFrame__StyleOptionVersion
type QStyleOptionFrame__StyleOptionVersion int64

var (
	QStyleOptionFrame__Version QStyleOptionFrame__StyleOptionVersion = QStyleOptionFrame__StyleOptionVersion(3)
)

type QStyleOptionGraphicsItem struct {
	QStyleOption
}

type QStyleOptionGraphicsItem_ITF interface {
	QStyleOption_ITF
	QStyleOptionGraphicsItem_PTR() *QStyleOptionGraphicsItem
}

func (ptr *QStyleOptionGraphicsItem) QStyleOptionGraphicsItem_PTR() *QStyleOptionGraphicsItem {
	return ptr
}

func (ptr *QStyleOptionGraphicsItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOption_PTR().Pointer()
	}
	return nil
}

func (ptr *QStyleOptionGraphicsItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QStyleOption_PTR().SetPointer(p)
	}
}

func PointerFromQStyleOptionGraphicsItem(ptr QStyleOptionGraphicsItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionGraphicsItem_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionGraphicsItemFromPointer(ptr unsafe.Pointer) (n *QStyleOptionGraphicsItem) {
	n = new(QStyleOptionGraphicsItem)
	n.SetPointer(ptr)
	return
}
func (ptr *QStyleOptionGraphicsItem) DestroyQStyleOptionGraphicsItem() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QStyleOptionGraphicsItem__StyleOptionType
//QStyleOptionGraphicsItem::StyleOptionType
type QStyleOptionGraphicsItem__StyleOptionType int64

var (
	QStyleOptionGraphicsItem__Type QStyleOptionGraphicsItem__StyleOptionType = QStyleOptionGraphicsItem__StyleOptionType(QStyleOption__SO_GraphicsItem)
)

// QStyleOptionGraphicsItem::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionGraphicsItem__StyleOptionVersion
type QStyleOptionGraphicsItem__StyleOptionVersion int64

var (
	QStyleOptionGraphicsItem__Version QStyleOptionGraphicsItem__StyleOptionVersion = QStyleOptionGraphicsItem__StyleOptionVersion(1)
)

func NewQStyleOptionGraphicsItem() *QStyleOptionGraphicsItem {
	tmpValue := NewQStyleOptionGraphicsItemFromPointer(C.QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem())
	qt.SetFinalizer(tmpValue, (*QStyleOptionGraphicsItem).DestroyQStyleOptionGraphicsItem)
	return tmpValue
}

func NewQStyleOptionGraphicsItem2(other QStyleOptionGraphicsItem_ITF) *QStyleOptionGraphicsItem {
	tmpValue := NewQStyleOptionGraphicsItemFromPointer(C.QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem2(PointerFromQStyleOptionGraphicsItem(other)))
	qt.SetFinalizer(tmpValue, (*QStyleOptionGraphicsItem).DestroyQStyleOptionGraphicsItem)
	return tmpValue
}

// QStyleOptionGroupBox::StyleOptionType
//
//go:generate stringer -type=QStyleOptionGroupBox__StyleOptionType
type QStyleOptionGroupBox__StyleOptionType int64

var (
	QStyleOptionGroupBox__Type QStyleOptionGroupBox__StyleOptionType = QStyleOptionGroupBox__StyleOptionType(QStyleOption__SO_GroupBox)
)

// QStyleOptionGroupBox::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionGroupBox__StyleOptionVersion
type QStyleOptionGroupBox__StyleOptionVersion int64

var (
	QStyleOptionGroupBox__Version QStyleOptionGroupBox__StyleOptionVersion = QStyleOptionGroupBox__StyleOptionVersion(1)
)

// QStyleOptionHeader::SectionPosition
//
//go:generate stringer -type=QStyleOptionHeader__SectionPosition
type QStyleOptionHeader__SectionPosition int64

const (
	QStyleOptionHeader__Beginning      QStyleOptionHeader__SectionPosition = QStyleOptionHeader__SectionPosition(0)
	QStyleOptionHeader__Middle         QStyleOptionHeader__SectionPosition = QStyleOptionHeader__SectionPosition(1)
	QStyleOptionHeader__End            QStyleOptionHeader__SectionPosition = QStyleOptionHeader__SectionPosition(2)
	QStyleOptionHeader__OnlyOneSection QStyleOptionHeader__SectionPosition = QStyleOptionHeader__SectionPosition(3)
)

// QStyleOptionHeader::SelectedPosition
//
//go:generate stringer -type=QStyleOptionHeader__SelectedPosition
type QStyleOptionHeader__SelectedPosition int64

const (
	QStyleOptionHeader__NotAdjacent                QStyleOptionHeader__SelectedPosition = QStyleOptionHeader__SelectedPosition(0)
	QStyleOptionHeader__NextIsSelected             QStyleOptionHeader__SelectedPosition = QStyleOptionHeader__SelectedPosition(1)
	QStyleOptionHeader__PreviousIsSelected         QStyleOptionHeader__SelectedPosition = QStyleOptionHeader__SelectedPosition(2)
	QStyleOptionHeader__NextAndPreviousAreSelected QStyleOptionHeader__SelectedPosition = QStyleOptionHeader__SelectedPosition(3)
)

// QStyleOptionHeader::SortIndicator
//
//go:generate stringer -type=QStyleOptionHeader__SortIndicator
type QStyleOptionHeader__SortIndicator int64

const (
	QStyleOptionHeader__None     QStyleOptionHeader__SortIndicator = QStyleOptionHeader__SortIndicator(0)
	QStyleOptionHeader__SortUp   QStyleOptionHeader__SortIndicator = QStyleOptionHeader__SortIndicator(1)
	QStyleOptionHeader__SortDown QStyleOptionHeader__SortIndicator = QStyleOptionHeader__SortIndicator(2)
)

// QStyleOptionHeader::StyleOptionType
//
//go:generate stringer -type=QStyleOptionHeader__StyleOptionType
type QStyleOptionHeader__StyleOptionType int64

var (
	QStyleOptionHeader__Type QStyleOptionHeader__StyleOptionType = QStyleOptionHeader__StyleOptionType(QStyleOption__SO_Header)
)

// QStyleOptionHeader::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionHeader__StyleOptionVersion
type QStyleOptionHeader__StyleOptionVersion int64

var (
	QStyleOptionHeader__Version QStyleOptionHeader__StyleOptionVersion = QStyleOptionHeader__StyleOptionVersion(1)
)

// QStyleOptionMenuItem::CheckType
//
//go:generate stringer -type=QStyleOptionMenuItem__CheckType
type QStyleOptionMenuItem__CheckType int64

const (
	QStyleOptionMenuItem__NotCheckable QStyleOptionMenuItem__CheckType = QStyleOptionMenuItem__CheckType(0)
	QStyleOptionMenuItem__Exclusive    QStyleOptionMenuItem__CheckType = QStyleOptionMenuItem__CheckType(1)
	QStyleOptionMenuItem__NonExclusive QStyleOptionMenuItem__CheckType = QStyleOptionMenuItem__CheckType(2)
)

// QStyleOptionMenuItem::MenuItemType
//
//go:generate stringer -type=QStyleOptionMenuItem__MenuItemType
type QStyleOptionMenuItem__MenuItemType int64

const (
	QStyleOptionMenuItem__Normal      QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(0)
	QStyleOptionMenuItem__DefaultItem QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(1)
	QStyleOptionMenuItem__Separator   QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(2)
	QStyleOptionMenuItem__SubMenu     QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(3)
	QStyleOptionMenuItem__Scroller    QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(4)
	QStyleOptionMenuItem__TearOff     QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(5)
	QStyleOptionMenuItem__Margin      QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(6)
	QStyleOptionMenuItem__EmptyArea   QStyleOptionMenuItem__MenuItemType = QStyleOptionMenuItem__MenuItemType(7)
)

// QStyleOptionMenuItem::StyleOptionType
//
//go:generate stringer -type=QStyleOptionMenuItem__StyleOptionType
type QStyleOptionMenuItem__StyleOptionType int64

var (
	QStyleOptionMenuItem__Type QStyleOptionMenuItem__StyleOptionType = QStyleOptionMenuItem__StyleOptionType(QStyleOption__SO_MenuItem)
)

// QStyleOptionMenuItem::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionMenuItem__StyleOptionVersion
type QStyleOptionMenuItem__StyleOptionVersion int64

var (
	QStyleOptionMenuItem__Version QStyleOptionMenuItem__StyleOptionVersion = QStyleOptionMenuItem__StyleOptionVersion(1)
)

// QStyleOptionProgressBar::StyleOptionType
//
//go:generate stringer -type=QStyleOptionProgressBar__StyleOptionType
type QStyleOptionProgressBar__StyleOptionType int64

var (
	QStyleOptionProgressBar__Type QStyleOptionProgressBar__StyleOptionType = QStyleOptionProgressBar__StyleOptionType(QStyleOption__SO_ProgressBar)
)

// QStyleOptionProgressBar::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionProgressBar__StyleOptionVersion
type QStyleOptionProgressBar__StyleOptionVersion int64

var (
	QStyleOptionProgressBar__Version QStyleOptionProgressBar__StyleOptionVersion = QStyleOptionProgressBar__StyleOptionVersion(2)
)

// QStyleOptionRubberBand::StyleOptionType
//
//go:generate stringer -type=QStyleOptionRubberBand__StyleOptionType
type QStyleOptionRubberBand__StyleOptionType int64

var (
	QStyleOptionRubberBand__Type QStyleOptionRubberBand__StyleOptionType = QStyleOptionRubberBand__StyleOptionType(QStyleOption__SO_RubberBand)
)

// QStyleOptionRubberBand::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionRubberBand__StyleOptionVersion
type QStyleOptionRubberBand__StyleOptionVersion int64

var (
	QStyleOptionRubberBand__Version QStyleOptionRubberBand__StyleOptionVersion = QStyleOptionRubberBand__StyleOptionVersion(1)
)

// QStyleOptionSizeGrip::StyleOptionType
//
//go:generate stringer -type=QStyleOptionSizeGrip__StyleOptionType
type QStyleOptionSizeGrip__StyleOptionType int64

var (
	QStyleOptionSizeGrip__Type QStyleOptionSizeGrip__StyleOptionType = QStyleOptionSizeGrip__StyleOptionType(QStyleOption__SO_SizeGrip)
)

// QStyleOptionSizeGrip::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionSizeGrip__StyleOptionVersion
type QStyleOptionSizeGrip__StyleOptionVersion int64

var (
	QStyleOptionSizeGrip__Version QStyleOptionSizeGrip__StyleOptionVersion = QStyleOptionSizeGrip__StyleOptionVersion(1)
)

// QStyleOptionSlider::StyleOptionType
//
//go:generate stringer -type=QStyleOptionSlider__StyleOptionType
type QStyleOptionSlider__StyleOptionType int64

var (
	QStyleOptionSlider__Type QStyleOptionSlider__StyleOptionType = QStyleOptionSlider__StyleOptionType(QStyleOption__SO_Slider)
)

// QStyleOptionSlider::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionSlider__StyleOptionVersion
type QStyleOptionSlider__StyleOptionVersion int64

var (
	QStyleOptionSlider__Version QStyleOptionSlider__StyleOptionVersion = QStyleOptionSlider__StyleOptionVersion(1)
)

// QStyleOptionSpinBox::StyleOptionType
//
//go:generate stringer -type=QStyleOptionSpinBox__StyleOptionType
type QStyleOptionSpinBox__StyleOptionType int64

var (
	QStyleOptionSpinBox__Type QStyleOptionSpinBox__StyleOptionType = QStyleOptionSpinBox__StyleOptionType(QStyleOption__SO_SpinBox)
)

// QStyleOptionSpinBox::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionSpinBox__StyleOptionVersion
type QStyleOptionSpinBox__StyleOptionVersion int64

var (
	QStyleOptionSpinBox__Version QStyleOptionSpinBox__StyleOptionVersion = QStyleOptionSpinBox__StyleOptionVersion(1)
)

// QStyleOptionTab::CornerWidget
//
//go:generate stringer -type=QStyleOptionTab__CornerWidget
type QStyleOptionTab__CornerWidget int64

const (
	QStyleOptionTab__NoCornerWidgets   QStyleOptionTab__CornerWidget = QStyleOptionTab__CornerWidget(0x00)
	QStyleOptionTab__LeftCornerWidget  QStyleOptionTab__CornerWidget = QStyleOptionTab__CornerWidget(0x01)
	QStyleOptionTab__RightCornerWidget QStyleOptionTab__CornerWidget = QStyleOptionTab__CornerWidget(0x02)
)

// QStyleOptionTab::SelectedPosition
//
//go:generate stringer -type=QStyleOptionTab__SelectedPosition
type QStyleOptionTab__SelectedPosition int64

const (
	QStyleOptionTab__NotAdjacent        QStyleOptionTab__SelectedPosition = QStyleOptionTab__SelectedPosition(0)
	QStyleOptionTab__NextIsSelected     QStyleOptionTab__SelectedPosition = QStyleOptionTab__SelectedPosition(1)
	QStyleOptionTab__PreviousIsSelected QStyleOptionTab__SelectedPosition = QStyleOptionTab__SelectedPosition(2)
)

// QStyleOptionTab::StyleOptionType
//
//go:generate stringer -type=QStyleOptionTab__StyleOptionType
type QStyleOptionTab__StyleOptionType int64

var (
	QStyleOptionTab__Type QStyleOptionTab__StyleOptionType = QStyleOptionTab__StyleOptionType(QStyleOption__SO_Tab)
)

// QStyleOptionTab::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionTab__StyleOptionVersion
type QStyleOptionTab__StyleOptionVersion int64

var (
	QStyleOptionTab__Version QStyleOptionTab__StyleOptionVersion = QStyleOptionTab__StyleOptionVersion(3)
)

// QStyleOptionTab::TabFeature
//
//go:generate stringer -type=QStyleOptionTab__TabFeature
type QStyleOptionTab__TabFeature int64

const (
	QStyleOptionTab__None     QStyleOptionTab__TabFeature = QStyleOptionTab__TabFeature(0x00)
	QStyleOptionTab__HasFrame QStyleOptionTab__TabFeature = QStyleOptionTab__TabFeature(0x01)
)

// QStyleOptionTab::TabPosition
//
//go:generate stringer -type=QStyleOptionTab__TabPosition
type QStyleOptionTab__TabPosition int64

const (
	QStyleOptionTab__Beginning  QStyleOptionTab__TabPosition = QStyleOptionTab__TabPosition(0)
	QStyleOptionTab__Middle     QStyleOptionTab__TabPosition = QStyleOptionTab__TabPosition(1)
	QStyleOptionTab__End        QStyleOptionTab__TabPosition = QStyleOptionTab__TabPosition(2)
	QStyleOptionTab__OnlyOneTab QStyleOptionTab__TabPosition = QStyleOptionTab__TabPosition(3)
)

// QStyleOptionTabBarBase::StyleOptionType
//
//go:generate stringer -type=QStyleOptionTabBarBase__StyleOptionType
type QStyleOptionTabBarBase__StyleOptionType int64

var (
	QStyleOptionTabBarBase__Type QStyleOptionTabBarBase__StyleOptionType = QStyleOptionTabBarBase__StyleOptionType(QStyleOption__SO_TabBarBase)
)

// QStyleOptionTabBarBase::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionTabBarBase__StyleOptionVersion
type QStyleOptionTabBarBase__StyleOptionVersion int64

var (
	QStyleOptionTabBarBase__Version QStyleOptionTabBarBase__StyleOptionVersion = QStyleOptionTabBarBase__StyleOptionVersion(2)
)

// QStyleOptionTabWidgetFrame::StyleOptionType
//
//go:generate stringer -type=QStyleOptionTabWidgetFrame__StyleOptionType
type QStyleOptionTabWidgetFrame__StyleOptionType int64

var (
	QStyleOptionTabWidgetFrame__Type QStyleOptionTabWidgetFrame__StyleOptionType = QStyleOptionTabWidgetFrame__StyleOptionType(QStyleOption__SO_TabWidgetFrame)
)

// QStyleOptionTabWidgetFrame::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionTabWidgetFrame__StyleOptionVersion
type QStyleOptionTabWidgetFrame__StyleOptionVersion int64

var (
	QStyleOptionTabWidgetFrame__Version QStyleOptionTabWidgetFrame__StyleOptionVersion = QStyleOptionTabWidgetFrame__StyleOptionVersion(2)
)

// QStyleOptionTitleBar::StyleOptionType
//
//go:generate stringer -type=QStyleOptionTitleBar__StyleOptionType
type QStyleOptionTitleBar__StyleOptionType int64

var (
	QStyleOptionTitleBar__Type QStyleOptionTitleBar__StyleOptionType = QStyleOptionTitleBar__StyleOptionType(QStyleOption__SO_TitleBar)
)

// QStyleOptionTitleBar::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionTitleBar__StyleOptionVersion
type QStyleOptionTitleBar__StyleOptionVersion int64

var (
	QStyleOptionTitleBar__Version QStyleOptionTitleBar__StyleOptionVersion = QStyleOptionTitleBar__StyleOptionVersion(1)
)

// QStyleOptionToolBar::StyleOptionType
//
//go:generate stringer -type=QStyleOptionToolBar__StyleOptionType
type QStyleOptionToolBar__StyleOptionType int64

var (
	QStyleOptionToolBar__Type QStyleOptionToolBar__StyleOptionType = QStyleOptionToolBar__StyleOptionType(QStyleOption__SO_ToolBar)
)

// QStyleOptionToolBar::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionToolBar__StyleOptionVersion
type QStyleOptionToolBar__StyleOptionVersion int64

var (
	QStyleOptionToolBar__Version QStyleOptionToolBar__StyleOptionVersion = QStyleOptionToolBar__StyleOptionVersion(1)
)

// QStyleOptionToolBar::ToolBarFeature
//
//go:generate stringer -type=QStyleOptionToolBar__ToolBarFeature
type QStyleOptionToolBar__ToolBarFeature int64

const (
	QStyleOptionToolBar__None    QStyleOptionToolBar__ToolBarFeature = QStyleOptionToolBar__ToolBarFeature(0x0)
	QStyleOptionToolBar__Movable QStyleOptionToolBar__ToolBarFeature = QStyleOptionToolBar__ToolBarFeature(0x1)
)

// QStyleOptionToolBar::ToolBarPosition
//
//go:generate stringer -type=QStyleOptionToolBar__ToolBarPosition
type QStyleOptionToolBar__ToolBarPosition int64

const (
	QStyleOptionToolBar__Beginning QStyleOptionToolBar__ToolBarPosition = QStyleOptionToolBar__ToolBarPosition(0)
	QStyleOptionToolBar__Middle    QStyleOptionToolBar__ToolBarPosition = QStyleOptionToolBar__ToolBarPosition(1)
	QStyleOptionToolBar__End       QStyleOptionToolBar__ToolBarPosition = QStyleOptionToolBar__ToolBarPosition(2)
	QStyleOptionToolBar__OnlyOne   QStyleOptionToolBar__ToolBarPosition = QStyleOptionToolBar__ToolBarPosition(3)
)

// QStyleOptionToolBox::SelectedPosition
//
//go:generate stringer -type=QStyleOptionToolBox__SelectedPosition
type QStyleOptionToolBox__SelectedPosition int64

const (
	QStyleOptionToolBox__NotAdjacent        QStyleOptionToolBox__SelectedPosition = QStyleOptionToolBox__SelectedPosition(0)
	QStyleOptionToolBox__NextIsSelected     QStyleOptionToolBox__SelectedPosition = QStyleOptionToolBox__SelectedPosition(1)
	QStyleOptionToolBox__PreviousIsSelected QStyleOptionToolBox__SelectedPosition = QStyleOptionToolBox__SelectedPosition(2)
)

// QStyleOptionToolBox::StyleOptionType
//
//go:generate stringer -type=QStyleOptionToolBox__StyleOptionType
type QStyleOptionToolBox__StyleOptionType int64

var (
	QStyleOptionToolBox__Type QStyleOptionToolBox__StyleOptionType = QStyleOptionToolBox__StyleOptionType(QStyleOption__SO_ToolBox)
)

// QStyleOptionToolBox::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionToolBox__StyleOptionVersion
type QStyleOptionToolBox__StyleOptionVersion int64

var (
	QStyleOptionToolBox__Version QStyleOptionToolBox__StyleOptionVersion = QStyleOptionToolBox__StyleOptionVersion(2)
)

// QStyleOptionToolBox::TabPosition
//
//go:generate stringer -type=QStyleOptionToolBox__TabPosition
type QStyleOptionToolBox__TabPosition int64

const (
	QStyleOptionToolBox__Beginning  QStyleOptionToolBox__TabPosition = QStyleOptionToolBox__TabPosition(0)
	QStyleOptionToolBox__Middle     QStyleOptionToolBox__TabPosition = QStyleOptionToolBox__TabPosition(1)
	QStyleOptionToolBox__End        QStyleOptionToolBox__TabPosition = QStyleOptionToolBox__TabPosition(2)
	QStyleOptionToolBox__OnlyOneTab QStyleOptionToolBox__TabPosition = QStyleOptionToolBox__TabPosition(3)
)

// QStyleOptionToolButton::StyleOptionType
//
//go:generate stringer -type=QStyleOptionToolButton__StyleOptionType
type QStyleOptionToolButton__StyleOptionType int64

var (
	QStyleOptionToolButton__Type QStyleOptionToolButton__StyleOptionType = QStyleOptionToolButton__StyleOptionType(QStyleOption__SO_ToolButton)
)

// QStyleOptionToolButton::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionToolButton__StyleOptionVersion
type QStyleOptionToolButton__StyleOptionVersion int64

var (
	QStyleOptionToolButton__Version QStyleOptionToolButton__StyleOptionVersion = QStyleOptionToolButton__StyleOptionVersion(1)
)

// QStyleOptionToolButton::ToolButtonFeature
//
//go:generate stringer -type=QStyleOptionToolButton__ToolButtonFeature
type QStyleOptionToolButton__ToolButtonFeature int64

const (
	QStyleOptionToolButton__None            QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(0x00)
	QStyleOptionToolButton__Arrow           QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(0x01)
	QStyleOptionToolButton__Menu            QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(0x04)
	QStyleOptionToolButton__MenuButtonPopup QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(QStyleOptionToolButton__Menu)
	QStyleOptionToolButton__PopupDelay      QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(0x08)
	QStyleOptionToolButton__HasMenu         QStyleOptionToolButton__ToolButtonFeature = QStyleOptionToolButton__ToolButtonFeature(0x10)
)

// QStyleOptionViewItem::Position
//
//go:generate stringer -type=QStyleOptionViewItem__Position
type QStyleOptionViewItem__Position int64

const (
	QStyleOptionViewItem__Left   QStyleOptionViewItem__Position = QStyleOptionViewItem__Position(0)
	QStyleOptionViewItem__Right  QStyleOptionViewItem__Position = QStyleOptionViewItem__Position(1)
	QStyleOptionViewItem__Top    QStyleOptionViewItem__Position = QStyleOptionViewItem__Position(2)
	QStyleOptionViewItem__Bottom QStyleOptionViewItem__Position = QStyleOptionViewItem__Position(3)
)

// QStyleOptionViewItem::StyleOptionType
//
//go:generate stringer -type=QStyleOptionViewItem__StyleOptionType
type QStyleOptionViewItem__StyleOptionType int64

var (
	QStyleOptionViewItem__Type QStyleOptionViewItem__StyleOptionType = QStyleOptionViewItem__StyleOptionType(QStyleOption__SO_ViewItem)
)

// QStyleOptionViewItem::StyleOptionVersion
//
//go:generate stringer -type=QStyleOptionViewItem__StyleOptionVersion
type QStyleOptionViewItem__StyleOptionVersion int64

var (
	QStyleOptionViewItem__Version QStyleOptionViewItem__StyleOptionVersion = QStyleOptionViewItem__StyleOptionVersion(4)
)

// QStyleOptionViewItem::ViewItemFeature
//
//go:generate stringer -type=QStyleOptionViewItem__ViewItemFeature
type QStyleOptionViewItem__ViewItemFeature int64

const (
	QStyleOptionViewItem__None              QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x00)
	QStyleOptionViewItem__WrapText          QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x01)
	QStyleOptionViewItem__Alternate         QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x02)
	QStyleOptionViewItem__HasCheckIndicator QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x04)
	QStyleOptionViewItem__HasDisplay        QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x08)
	QStyleOptionViewItem__HasDecoration     QStyleOptionViewItem__ViewItemFeature = QStyleOptionViewItem__ViewItemFeature(0x10)
)

// QStyleOptionViewItem::ViewItemPosition
//
//go:generate stringer -type=QStyleOptionViewItem__ViewItemPosition
type QStyleOptionViewItem__ViewItemPosition int64

const (
	QStyleOptionViewItem__Invalid   QStyleOptionViewItem__ViewItemPosition = QStyleOptionViewItem__ViewItemPosition(0)
	QStyleOptionViewItem__Beginning QStyleOptionViewItem__ViewItemPosition = QStyleOptionViewItem__ViewItemPosition(1)
	QStyleOptionViewItem__Middle    QStyleOptionViewItem__ViewItemPosition = QStyleOptionViewItem__ViewItemPosition(2)
	QStyleOptionViewItem__End       QStyleOptionViewItem__ViewItemPosition = QStyleOptionViewItem__ViewItemPosition(3)
	QStyleOptionViewItem__OnlyOne   QStyleOptionViewItem__ViewItemPosition = QStyleOptionViewItem__ViewItemPosition(4)
)

// QSwipeGesture::SwipeDirection
//
//go:generate stringer -type=QSwipeGesture__SwipeDirection
type QSwipeGesture__SwipeDirection int64

const (
	QSwipeGesture__NoDirection QSwipeGesture__SwipeDirection = QSwipeGesture__SwipeDirection(0)
	QSwipeGesture__Left        QSwipeGesture__SwipeDirection = QSwipeGesture__SwipeDirection(1)
	QSwipeGesture__Right       QSwipeGesture__SwipeDirection = QSwipeGesture__SwipeDirection(2)
	QSwipeGesture__Up          QSwipeGesture__SwipeDirection = QSwipeGesture__SwipeDirection(3)
	QSwipeGesture__Down        QSwipeGesture__SwipeDirection = QSwipeGesture__SwipeDirection(4)
)

// QSystemTrayIcon::ActivationReason
//
//go:generate stringer -type=QSystemTrayIcon__ActivationReason
type QSystemTrayIcon__ActivationReason int64

const (
	QSystemTrayIcon__Unknown     QSystemTrayIcon__ActivationReason = QSystemTrayIcon__ActivationReason(0)
	QSystemTrayIcon__Context     QSystemTrayIcon__ActivationReason = QSystemTrayIcon__ActivationReason(1)
	QSystemTrayIcon__DoubleClick QSystemTrayIcon__ActivationReason = QSystemTrayIcon__ActivationReason(2)
	QSystemTrayIcon__Trigger     QSystemTrayIcon__ActivationReason = QSystemTrayIcon__ActivationReason(3)
	QSystemTrayIcon__MiddleClick QSystemTrayIcon__ActivationReason = QSystemTrayIcon__ActivationReason(4)
)

// QSystemTrayIcon::MessageIcon
//
//go:generate stringer -type=QSystemTrayIcon__MessageIcon
type QSystemTrayIcon__MessageIcon int64

const (
	QSystemTrayIcon__NoIcon      QSystemTrayIcon__MessageIcon = QSystemTrayIcon__MessageIcon(0)
	QSystemTrayIcon__Information QSystemTrayIcon__MessageIcon = QSystemTrayIcon__MessageIcon(1)
	QSystemTrayIcon__Warning     QSystemTrayIcon__MessageIcon = QSystemTrayIcon__MessageIcon(2)
	QSystemTrayIcon__Critical    QSystemTrayIcon__MessageIcon = QSystemTrayIcon__MessageIcon(3)
)

// QTabBar::ButtonPosition
//
//go:generate stringer -type=QTabBar__ButtonPosition
type QTabBar__ButtonPosition int64

const (
	QTabBar__LeftSide  QTabBar__ButtonPosition = QTabBar__ButtonPosition(0)
	QTabBar__RightSide QTabBar__ButtonPosition = QTabBar__ButtonPosition(1)
)

// QTabBar::SelectionBehavior
//
//go:generate stringer -type=QTabBar__SelectionBehavior
type QTabBar__SelectionBehavior int64

const (
	QTabBar__SelectLeftTab     QTabBar__SelectionBehavior = QTabBar__SelectionBehavior(0)
	QTabBar__SelectRightTab    QTabBar__SelectionBehavior = QTabBar__SelectionBehavior(1)
	QTabBar__SelectPreviousTab QTabBar__SelectionBehavior = QTabBar__SelectionBehavior(2)
)

// QTabBar::Shape
//
//go:generate stringer -type=QTabBar__Shape
type QTabBar__Shape int64

const (
	QTabBar__RoundedNorth    QTabBar__Shape = QTabBar__Shape(0)
	QTabBar__RoundedSouth    QTabBar__Shape = QTabBar__Shape(1)
	QTabBar__RoundedWest     QTabBar__Shape = QTabBar__Shape(2)
	QTabBar__RoundedEast     QTabBar__Shape = QTabBar__Shape(3)
	QTabBar__TriangularNorth QTabBar__Shape = QTabBar__Shape(4)
	QTabBar__TriangularSouth QTabBar__Shape = QTabBar__Shape(5)
	QTabBar__TriangularWest  QTabBar__Shape = QTabBar__Shape(6)
	QTabBar__TriangularEast  QTabBar__Shape = QTabBar__Shape(7)
)

// QTabWidget::TabPosition
//
//go:generate stringer -type=QTabWidget__TabPosition
type QTabWidget__TabPosition int64

const (
	QTabWidget__North QTabWidget__TabPosition = QTabWidget__TabPosition(0)
	QTabWidget__South QTabWidget__TabPosition = QTabWidget__TabPosition(1)
	QTabWidget__West  QTabWidget__TabPosition = QTabWidget__TabPosition(2)
	QTabWidget__East  QTabWidget__TabPosition = QTabWidget__TabPosition(3)
)

// QTabWidget::TabShape
//
//go:generate stringer -type=QTabWidget__TabShape
type QTabWidget__TabShape int64

const (
	QTabWidget__Rounded    QTabWidget__TabShape = QTabWidget__TabShape(0)
	QTabWidget__Triangular QTabWidget__TabShape = QTabWidget__TabShape(1)
)

// QTableWidgetItem::ItemType
//
//go:generate stringer -type=QTableWidgetItem__ItemType
type QTableWidgetItem__ItemType int64

const (
	QTableWidgetItem__Type     QTableWidgetItem__ItemType = QTableWidgetItem__ItemType(0)
	QTableWidgetItem__UserType QTableWidgetItem__ItemType = QTableWidgetItem__ItemType(1000)
)

type QTextEdit struct {
	QAbstractScrollArea
}

type QTextEdit_ITF interface {
	QAbstractScrollArea_ITF
	QTextEdit_PTR() *QTextEdit
}

func (ptr *QTextEdit) QTextEdit_PTR() *QTextEdit {
	return ptr
}

func (ptr *QTextEdit) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextEdit) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractScrollArea_PTR().SetPointer(p)
	}
}

func PointerFromQTextEdit(ptr QTextEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextEdit_PTR().Pointer()
	}
	return nil
}

func NewQTextEditFromPointer(ptr unsafe.Pointer) (n *QTextEdit) {
	n = new(QTextEdit)
	n.SetPointer(ptr)
	return
}

// QTextEdit::AutoFormattingFlag
//
//go:generate stringer -type=QTextEdit__AutoFormattingFlag
type QTextEdit__AutoFormattingFlag int64

const (
	QTextEdit__AutoNone       QTextEdit__AutoFormattingFlag = QTextEdit__AutoFormattingFlag(0)
	QTextEdit__AutoBulletList QTextEdit__AutoFormattingFlag = QTextEdit__AutoFormattingFlag(0x00000001)
	QTextEdit__AutoAll        QTextEdit__AutoFormattingFlag = QTextEdit__AutoFormattingFlag(0xffffffff)
)

// QTextEdit::LineWrapMode
//
//go:generate stringer -type=QTextEdit__LineWrapMode
type QTextEdit__LineWrapMode int64

const (
	QTextEdit__NoWrap           QTextEdit__LineWrapMode = QTextEdit__LineWrapMode(0)
	QTextEdit__WidgetWidth      QTextEdit__LineWrapMode = QTextEdit__LineWrapMode(1)
	QTextEdit__FixedPixelWidth  QTextEdit__LineWrapMode = QTextEdit__LineWrapMode(2)
	QTextEdit__FixedColumnWidth QTextEdit__LineWrapMode = QTextEdit__LineWrapMode(3)
)

func NewQTextEdit(parent QWidget_ITF) *QTextEdit {
	tmpValue := NewQTextEditFromPointer(C.QTextEdit_NewQTextEdit(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQTextEdit2(text string, parent QWidget_ITF) *QTextEdit {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQTextEditFromPointer(C.QTextEdit_NewQTextEdit2(C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QTextEdit) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextEdit_Alignment(ptr.Pointer()))
	}
	return 0
}

//export callbackQTextEdit_Append
func callbackQTextEdit_Append(ptr unsafe.Pointer, text C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "append"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	} else {
		NewQTextEditFromPointer(ptr).AppendDefault(cGoUnpackString(text))
	}
}

func (ptr *QTextEdit) ConnectAppend(f func(text string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "append"); signal != nil {
			f := func(text string) {
				(*(*func(string))(signal))(text)
				f(text)
			}
			qt.ConnectSignal(ptr.Pointer(), "append", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "append", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectAppend() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "append")
	}
}

func (ptr *QTextEdit) Append(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextEdit_Append(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextEdit) AppendDefault(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextEdit_AppendDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQTextEdit_Clear
func callbackQTextEdit_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextEditFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QTextEdit) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *QTextEdit) Clear() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QTextEdit) ClearDefault() {
	if ptr.Pointer() != nil {
		C.QTextEdit_ClearDefault(ptr.Pointer())
	}
}

//export callbackQTextEdit_Copy
func callbackQTextEdit_Copy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "copy"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextEditFromPointer(ptr).CopyDefault()
	}
}

func (ptr *QTextEdit) ConnectCopy(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "copy"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "copy", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "copy", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectCopy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "copy")
	}
}

func (ptr *QTextEdit) Copy() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Copy(ptr.Pointer())
	}
}

func (ptr *QTextEdit) CopyDefault() {
	if ptr.Pointer() != nil {
		C.QTextEdit_CopyDefault(ptr.Pointer())
	}
}

func (ptr *QTextEdit) CreateStandardContextMenu() *QMenu {
	if ptr.Pointer() != nil {
		tmpValue := NewQMenuFromPointer(C.QTextEdit_CreateStandardContextMenu(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextEdit) CreateStandardContextMenu2(position core.QPoint_ITF) *QMenu {
	if ptr.Pointer() != nil {
		tmpValue := NewQMenuFromPointer(C.QTextEdit_CreateStandardContextMenu2(ptr.Pointer(), core.PointerFromQPoint(position)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextEdit) CurrentCharFormat() *gui.QTextCharFormat {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQTextCharFormatFromPointer(C.QTextEdit_CurrentCharFormat(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QTextCharFormat).DestroyQTextCharFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QTextEdit) CurrentFont() *gui.QFont {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontFromPointer(C.QTextEdit_CurrentFont(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QTextEdit) Document() *gui.QTextDocument {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQTextDocumentFromPointer(C.QTextEdit_Document(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextEdit) Find(exp string, options gui.QTextDocument__FindFlag) bool {
	if ptr.Pointer() != nil {
		var expC *C.char
		if exp != "" {
			expC = C.CString(exp)
			defer C.free(unsafe.Pointer(expC))
		}
		return int8(C.QTextEdit_Find(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: expC, len: C.longlong(len(exp))}, C.longlong(options))) != 0
	}
	return false
}

func (ptr *QTextEdit) Find2(exp core.QRegExp_ITF, options gui.QTextDocument__FindFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextEdit_Find2(ptr.Pointer(), core.PointerFromQRegExp(exp), C.longlong(options))) != 0
	}
	return false
}

func (ptr *QTextEdit) Find3(exp core.QRegularExpression_ITF, options gui.QTextDocument__FindFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextEdit_Find3(ptr.Pointer(), core.PointerFromQRegularExpression(exp), C.longlong(options))) != 0
	}
	return false
}

func (ptr *QTextEdit) FontItalic() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextEdit_FontItalic(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) FontPointSize() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextEdit_FontPointSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextEdit) FontUnderline() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextEdit_FontUnderline(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) FontWeight() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextEdit_FontWeight(ptr.Pointer())))
	}
	return 0
}

//export callbackQTextEdit_InsertPlainText
func callbackQTextEdit_InsertPlainText(ptr unsafe.Pointer, text C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "insertPlainText"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	} else {
		NewQTextEditFromPointer(ptr).InsertPlainTextDefault(cGoUnpackString(text))
	}
}

func (ptr *QTextEdit) ConnectInsertPlainText(f func(text string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "insertPlainText"); signal != nil {
			f := func(text string) {
				(*(*func(string))(signal))(text)
				f(text)
			}
			qt.ConnectSignal(ptr.Pointer(), "insertPlainText", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "insertPlainText", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectInsertPlainText() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "insertPlainText")
	}
}

func (ptr *QTextEdit) InsertPlainText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextEdit_InsertPlainText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextEdit) InsertPlainTextDefault(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextEdit_InsertPlainTextDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextEdit) MergeCurrentCharFormat(modifier gui.QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_MergeCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(modifier))
	}
}

//export callbackQTextEdit_Paste
func callbackQTextEdit_Paste(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paste"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextEditFromPointer(ptr).PasteDefault()
	}
}

func (ptr *QTextEdit) ConnectPaste(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paste"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "paste", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paste", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectPaste() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paste")
	}
}

func (ptr *QTextEdit) Paste() {
	if ptr.Pointer() != nil {
		C.QTextEdit_Paste(ptr.Pointer())
	}
}

func (ptr *QTextEdit) PasteDefault() {
	if ptr.Pointer() != nil {
		C.QTextEdit_PasteDefault(ptr.Pointer())
	}
}

func (ptr *QTextEdit) PlaceholderText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextEdit_PlaceholderText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) Print(printer gui.QPagedPaintDevice_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_Print(ptr.Pointer(), gui.PointerFromQPagedPaintDevice(printer))
	}
}

//export callbackQTextEdit_SelectionChanged
func callbackQTextEdit_SelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QTextEdit) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionChanged") {
			C.QTextEdit_ConnectSelectionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectionChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionChanged")
	}
}

func (ptr *QTextEdit) SelectionChanged() {
	if ptr.Pointer() != nil {
		C.QTextEdit_SelectionChanged(ptr.Pointer())
	}
}

//export callbackQTextEdit_SetAlignment
func callbackQTextEdit_SetAlignment(ptr unsafe.Pointer, a C.longlong) {
	if signal := qt.GetSignal(ptr, "setAlignment"); signal != nil {
		(*(*func(core.Qt__AlignmentFlag))(signal))(core.Qt__AlignmentFlag(a))
	} else {
		NewQTextEditFromPointer(ptr).SetAlignmentDefault(core.Qt__AlignmentFlag(a))
	}
}

func (ptr *QTextEdit) ConnectSetAlignment(f func(a core.Qt__AlignmentFlag)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAlignment"); signal != nil {
			f := func(a core.Qt__AlignmentFlag) {
				(*(*func(core.Qt__AlignmentFlag))(signal))(a)
				f(a)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAlignment", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAlignment", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectSetAlignment() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAlignment")
	}
}

func (ptr *QTextEdit) SetAlignment(a core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetAlignment(ptr.Pointer(), C.longlong(a))
	}
}

func (ptr *QTextEdit) SetAlignmentDefault(a core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetAlignmentDefault(ptr.Pointer(), C.longlong(a))
	}
}

func (ptr *QTextEdit) SetCurrentCharFormat(format gui.QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetCurrentCharFormat(ptr.Pointer(), gui.PointerFromQTextCharFormat(format))
	}
}

//export callbackQTextEdit_SetFontItalic
func callbackQTextEdit_SetFontItalic(ptr unsafe.Pointer, italic C.char) {
	if signal := qt.GetSignal(ptr, "setFontItalic"); signal != nil {
		(*(*func(bool))(signal))(int8(italic) != 0)
	} else {
		NewQTextEditFromPointer(ptr).SetFontItalicDefault(int8(italic) != 0)
	}
}

func (ptr *QTextEdit) ConnectSetFontItalic(f func(italic bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFontItalic"); signal != nil {
			f := func(italic bool) {
				(*(*func(bool))(signal))(italic)
				f(italic)
			}
			qt.ConnectSignal(ptr.Pointer(), "setFontItalic", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFontItalic", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectSetFontItalic() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFontItalic")
	}
}

func (ptr *QTextEdit) SetFontItalic(italic bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontItalic(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(italic))))
	}
}

func (ptr *QTextEdit) SetFontItalicDefault(italic bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontItalicDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(italic))))
	}
}

//export callbackQTextEdit_SetFontPointSize
func callbackQTextEdit_SetFontPointSize(ptr unsafe.Pointer, s C.double) {
	if signal := qt.GetSignal(ptr, "setFontPointSize"); signal != nil {
		(*(*func(float64))(signal))(float64(s))
	} else {
		NewQTextEditFromPointer(ptr).SetFontPointSizeDefault(float64(s))
	}
}

func (ptr *QTextEdit) ConnectSetFontPointSize(f func(s float64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFontPointSize"); signal != nil {
			f := func(s float64) {
				(*(*func(float64))(signal))(s)
				f(s)
			}
			qt.ConnectSignal(ptr.Pointer(), "setFontPointSize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFontPointSize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectSetFontPointSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFontPointSize")
	}
}

func (ptr *QTextEdit) SetFontPointSize(s float64) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontPointSize(ptr.Pointer(), C.double(s))
	}
}

func (ptr *QTextEdit) SetFontPointSizeDefault(s float64) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontPointSizeDefault(ptr.Pointer(), C.double(s))
	}
}

//export callbackQTextEdit_SetFontUnderline
func callbackQTextEdit_SetFontUnderline(ptr unsafe.Pointer, underline C.char) {
	if signal := qt.GetSignal(ptr, "setFontUnderline"); signal != nil {
		(*(*func(bool))(signal))(int8(underline) != 0)
	} else {
		NewQTextEditFromPointer(ptr).SetFontUnderlineDefault(int8(underline) != 0)
	}
}

func (ptr *QTextEdit) ConnectSetFontUnderline(f func(underline bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFontUnderline"); signal != nil {
			f := func(underline bool) {
				(*(*func(bool))(signal))(underline)
				f(underline)
			}
			qt.ConnectSignal(ptr.Pointer(), "setFontUnderline", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFontUnderline", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectSetFontUnderline() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFontUnderline")
	}
}

func (ptr *QTextEdit) SetFontUnderline(underline bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontUnderline(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(underline))))
	}
}

func (ptr *QTextEdit) SetFontUnderlineDefault(underline bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontUnderlineDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(underline))))
	}
}

//export callbackQTextEdit_SetFontWeight
func callbackQTextEdit_SetFontWeight(ptr unsafe.Pointer, weight C.int) {
	if signal := qt.GetSignal(ptr, "setFontWeight"); signal != nil {
		(*(*func(int))(signal))(int(int32(weight)))
	} else {
		NewQTextEditFromPointer(ptr).SetFontWeightDefault(int(int32(weight)))
	}
}

func (ptr *QTextEdit) ConnectSetFontWeight(f func(weight int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFontWeight"); signal != nil {
			f := func(weight int) {
				(*(*func(int))(signal))(weight)
				f(weight)
			}
			qt.ConnectSignal(ptr.Pointer(), "setFontWeight", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFontWeight", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectSetFontWeight() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFontWeight")
	}
}

func (ptr *QTextEdit) SetFontWeight(weight int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontWeight(ptr.Pointer(), C.int(int32(weight)))
	}
}

func (ptr *QTextEdit) SetFontWeightDefault(weight int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetFontWeightDefault(ptr.Pointer(), C.int(int32(weight)))
	}
}

//export callbackQTextEdit_SetHtml
func callbackQTextEdit_SetHtml(ptr unsafe.Pointer, text C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setHtml"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	} else {
		NewQTextEditFromPointer(ptr).SetHtmlDefault(cGoUnpackString(text))
	}
}

func (ptr *QTextEdit) ConnectSetHtml(f func(text string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHtml"); signal != nil {
			f := func(text string) {
				(*(*func(string))(signal))(text)
				f(text)
			}
			qt.ConnectSignal(ptr.Pointer(), "setHtml", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHtml", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectSetHtml() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHtml")
	}
}

func (ptr *QTextEdit) SetHtml(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextEdit_SetHtml(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextEdit) SetHtmlDefault(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextEdit_SetHtmlDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextEdit) SetPlaceholderText(placeholderText string) {
	if ptr.Pointer() != nil {
		var placeholderTextC *C.char
		if placeholderText != "" {
			placeholderTextC = C.CString(placeholderText)
			defer C.free(unsafe.Pointer(placeholderTextC))
		}
		C.QTextEdit_SetPlaceholderText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: placeholderTextC, len: C.longlong(len(placeholderText))})
	}
}

func (ptr *QTextEdit) SetReadOnly(ro bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetReadOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ro))))
	}
}

func (ptr *QTextEdit) SetTabChangesFocus(b bool) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTabChangesFocus(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(b))))
	}
}

func (ptr *QTextEdit) SetTabStopWidth(width int) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTabStopWidth(ptr.Pointer(), C.int(int32(width)))
	}
}

//export callbackQTextEdit_SetText
func callbackQTextEdit_SetText(ptr unsafe.Pointer, text C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setText"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	} else {
		NewQTextEditFromPointer(ptr).SetTextDefault(cGoUnpackString(text))
	}
}

func (ptr *QTextEdit) ConnectSetText(f func(text string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setText"); signal != nil {
			f := func(text string) {
				(*(*func(string))(signal))(text)
				f(text)
			}
			qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setText", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectSetText() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setText")
	}
}

func (ptr *QTextEdit) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextEdit_SetText(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextEdit) SetTextDefault(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextEdit_SetTextDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextEdit) SetTextCursor(cursor gui.QTextCursor_ITF) {
	if ptr.Pointer() != nil {
		C.QTextEdit_SetTextCursor(ptr.Pointer(), gui.PointerFromQTextCursor(cursor))
	}
}

func (ptr *QTextEdit) TabChangesFocus() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextEdit_TabChangesFocus(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextEdit) TabStopWidth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextEdit_TabStopWidth(ptr.Pointer())))
	}
	return 0
}

//export callbackQTextEdit_TextChanged
func callbackQTextEdit_TextChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "textChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QTextEdit) ConnectTextChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "textChanged") {
			C.QTextEdit_ConnectTextChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "textChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "textChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "textChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectTextChanged() {
	if ptr.Pointer() != nil {
		C.QTextEdit_DisconnectTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "textChanged")
	}
}

func (ptr *QTextEdit) TextChanged() {
	if ptr.Pointer() != nil {
		C.QTextEdit_TextChanged(ptr.Pointer())
	}
}

func (ptr *QTextEdit) TextColor() *gui.QColor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQColorFromPointer(C.QTextEdit_TextColor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
		return tmpValue
	}
	return nil
}

func (ptr *QTextEdit) TextCursor() *gui.QTextCursor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQTextCursorFromPointer(C.QTextEdit_TextCursor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QTextCursor).DestroyQTextCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QTextEdit) ToHtml() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextEdit_ToHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextEdit) ToPlainText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextEdit_ToPlainText(ptr.Pointer()))
	}
	return ""
}

//export callbackQTextEdit_DestroyQTextEdit
func callbackQTextEdit_DestroyQTextEdit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTextEdit"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextEditFromPointer(ptr).DestroyQTextEditDefault()
	}
}

func (ptr *QTextEdit) ConnectDestroyQTextEdit(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTextEdit"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTextEdit", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTextEdit", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextEdit) DisconnectDestroyQTextEdit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTextEdit")
	}
}

func (ptr *QTextEdit) DestroyQTextEdit() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextEdit_DestroyQTextEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextEdit) DestroyQTextEditDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextEdit_DestroyQTextEditDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QToolBar struct {
	QWidget
}

type QToolBar_ITF interface {
	QWidget_ITF
	QToolBar_PTR() *QToolBar
}

func (ptr *QToolBar) QToolBar_PTR() *QToolBar {
	return ptr
}

func (ptr *QToolBar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QToolBar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQToolBar(ptr QToolBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolBar_PTR().Pointer()
	}
	return nil
}

func NewQToolBarFromPointer(ptr unsafe.Pointer) (n *QToolBar) {
	n = new(QToolBar)
	n.SetPointer(ptr)
	return
}
func NewQToolBar(title string, parent QWidget_ITF) *QToolBar {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	tmpValue := NewQToolBarFromPointer(C.QToolBar_NewQToolBar(C.struct_QtWidgets_PackedString{data: titleC, len: C.longlong(len(title))}, PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQToolBar2(parent QWidget_ITF) *QToolBar {
	tmpValue := NewQToolBarFromPointer(C.QToolBar_NewQToolBar2(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QToolBar) AddAction(text string) *QAction {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := NewQActionFromPointer(C.QToolBar_AddAction(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QToolBar) AddAction2(icon gui.QIcon_ITF, text string) *QAction {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := NewQActionFromPointer(C.QToolBar_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QToolBar) AddAction3(text string, receiver core.QObject_ITF, member string) *QAction {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		tmpValue := NewQActionFromPointer(C.QToolBar_AddAction3(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQObject(receiver), memberC))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QToolBar) AddAction4(icon gui.QIcon_ITF, text string, receiver core.QObject_ITF, member string) *QAction {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		var memberC *C.char
		if member != "" {
			memberC = C.CString(member)
			defer C.free(unsafe.Pointer(memberC))
		}
		tmpValue := NewQActionFromPointer(C.QToolBar_AddAction4(ptr.Pointer(), gui.PointerFromQIcon(icon), C.struct_QtWidgets_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQObject(receiver), memberC))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QToolBar) AddSeparator() *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QToolBar_AddSeparator(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QToolBar) AddWidget(widget QWidget_ITF) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QToolBar_AddWidget(ptr.Pointer(), PointerFromQWidget(widget)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QToolBar) Clear() {
	if ptr.Pointer() != nil {
		C.QToolBar_Clear(ptr.Pointer())
	}
}

func (ptr *QToolBar) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QToolBar_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolBar) SetOrientation(orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QToolBar_SetOrientation(ptr.Pointer(), C.longlong(orientation))
	}
}

//export callbackQToolBar_DestroyQToolBar
func callbackQToolBar_DestroyQToolBar(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QToolBar"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQToolBarFromPointer(ptr).DestroyQToolBarDefault()
	}
}

func (ptr *QToolBar) ConnectDestroyQToolBar(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QToolBar"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QToolBar", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QToolBar", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QToolBar) DisconnectDestroyQToolBar() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QToolBar")
	}
}

func (ptr *QToolBar) DestroyQToolBar() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QToolBar_DestroyQToolBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QToolBar) DestroyQToolBarDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QToolBar_DestroyQToolBarDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QToolButton::ToolButtonPopupMode
//
//go:generate stringer -type=QToolButton__ToolButtonPopupMode
type QToolButton__ToolButtonPopupMode int64

const (
	QToolButton__DelayedPopup    QToolButton__ToolButtonPopupMode = QToolButton__ToolButtonPopupMode(0)
	QToolButton__MenuButtonPopup QToolButton__ToolButtonPopupMode = QToolButton__ToolButtonPopupMode(1)
	QToolButton__InstantPopup    QToolButton__ToolButtonPopupMode = QToolButton__ToolButtonPopupMode(2)
)

type QTreeView struct {
	QAbstractItemView
}

type QTreeView_ITF interface {
	QAbstractItemView_ITF
	QTreeView_PTR() *QTreeView
}

func (ptr *QTreeView) QTreeView_PTR() *QTreeView {
	return ptr
}

func (ptr *QTreeView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemView_PTR().Pointer()
	}
	return nil
}

func (ptr *QTreeView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractItemView_PTR().SetPointer(p)
	}
}

func PointerFromQTreeView(ptr QTreeView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTreeView_PTR().Pointer()
	}
	return nil
}

func NewQTreeViewFromPointer(ptr unsafe.Pointer) (n *QTreeView) {
	n = new(QTreeView)
	n.SetPointer(ptr)
	return
}
func NewQTreeView(parent QWidget_ITF) *QTreeView {
	tmpValue := NewQTreeViewFromPointer(C.QTreeView_NewQTreeView(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTreeView_Collapse
func callbackQTreeView_Collapse(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "collapse"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQTreeViewFromPointer(ptr).CollapseDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QTreeView) ConnectCollapse(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "collapse"); signal != nil {
			f := func(index *core.QModelIndex) {
				(*(*func(*core.QModelIndex))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "collapse", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "collapse", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTreeView) DisconnectCollapse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "collapse")
	}
}

func (ptr *QTreeView) Collapse(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_Collapse(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) CollapseDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_CollapseDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQTreeView_Collapsed
func callbackQTreeView_Collapsed(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "collapsed"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QTreeView) ConnectCollapsed(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "collapsed") {
			C.QTreeView_ConnectCollapsed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "collapsed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "collapsed"); signal != nil {
			f := func(index *core.QModelIndex) {
				(*(*func(*core.QModelIndex))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "collapsed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "collapsed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTreeView) DisconnectCollapsed() {
	if ptr.Pointer() != nil {
		C.QTreeView_DisconnectCollapsed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "collapsed")
	}
}

func (ptr *QTreeView) Collapsed(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_Collapsed(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQTreeView_Expand
func callbackQTreeView_Expand(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "expand"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	} else {
		NewQTreeViewFromPointer(ptr).ExpandDefault(core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QTreeView) ConnectExpand(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "expand"); signal != nil {
			f := func(index *core.QModelIndex) {
				(*(*func(*core.QModelIndex))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "expand", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "expand", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTreeView) DisconnectExpand() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "expand")
	}
}

func (ptr *QTreeView) Expand(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_Expand(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) ExpandDefault(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_ExpandDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

//export callbackQTreeView_Expanded
func callbackQTreeView_Expanded(ptr unsafe.Pointer, index unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "expanded"); signal != nil {
		(*(*func(*core.QModelIndex))(signal))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QTreeView) ConnectExpanded(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "expanded") {
			C.QTreeView_ConnectExpanded(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "expanded")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "expanded"); signal != nil {
			f := func(index *core.QModelIndex) {
				(*(*func(*core.QModelIndex))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "expanded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "expanded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTreeView) DisconnectExpanded() {
	if ptr.Pointer() != nil {
		C.QTreeView_DisconnectExpanded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "expanded")
	}
}

func (ptr *QTreeView) Expanded(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_Expanded(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QTreeView) Header() *QHeaderView {
	if ptr.Pointer() != nil {
		tmpValue := NewQHeaderViewFromPointer(C.QTreeView_Header(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQTreeView_IndexAt
func callbackQTreeView_IndexAt(ptr unsafe.Pointer, point unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "indexAt"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QPoint) *core.QModelIndex)(signal))(core.NewQPointFromPointer(point)))
	}

	return core.PointerFromQModelIndex(NewQTreeViewFromPointer(ptr).IndexAtDefault(core.NewQPointFromPointer(point)))
}

func (ptr *QTreeView) ConnectIndexAt(f func(point *core.QPoint) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "indexAt"); signal != nil {
			f := func(point *core.QPoint) *core.QModelIndex {
				(*(*func(*core.QPoint) *core.QModelIndex)(signal))(point)
				return f(point)
			}
			qt.ConnectSignal(ptr.Pointer(), "indexAt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexAt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTreeView) DisconnectIndexAt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "indexAt")
	}
}

func (ptr *QTreeView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QTreeView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QTreeView) IndexAtDefault(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QTreeView_IndexAtDefault(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QTreeView) IsExpanded(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTreeView_IsExpanded(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

//export callbackQTreeView_ResizeColumnToContents
func callbackQTreeView_ResizeColumnToContents(ptr unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "resizeColumnToContents"); signal != nil {
		(*(*func(int))(signal))(int(int32(column)))
	} else {
		NewQTreeViewFromPointer(ptr).ResizeColumnToContentsDefault(int(int32(column)))
	}
}

func (ptr *QTreeView) ConnectResizeColumnToContents(f func(column int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resizeColumnToContents"); signal != nil {
			f := func(column int) {
				(*(*func(int))(signal))(column)
				f(column)
			}
			qt.ConnectSignal(ptr.Pointer(), "resizeColumnToContents", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resizeColumnToContents", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTreeView) DisconnectResizeColumnToContents() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resizeColumnToContents")
	}
}

func (ptr *QTreeView) ResizeColumnToContents(column int) {
	if ptr.Pointer() != nil {
		C.QTreeView_ResizeColumnToContents(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QTreeView) ResizeColumnToContentsDefault(column int) {
	if ptr.Pointer() != nil {
		C.QTreeView_ResizeColumnToContentsDefault(ptr.Pointer(), C.int(int32(column)))
	}
}

//export callbackQTreeView_SelectionChanged
func callbackQTreeView_SelectionChanged(ptr unsafe.Pointer, selected unsafe.Pointer, deselected unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		(*(*func(*core.QItemSelection, *core.QItemSelection))(signal))(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	} else {
		NewQTreeViewFromPointer(ptr).SelectionChangedDefault(core.NewQItemSelectionFromPointer(selected), core.NewQItemSelectionFromPointer(deselected))
	}
}

func (ptr *QTreeView) ConnectSelectionChanged(f func(selected *core.QItemSelection, deselected *core.QItemSelection)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectionChanged"); signal != nil {
			f := func(selected *core.QItemSelection, deselected *core.QItemSelection) {
				(*(*func(*core.QItemSelection, *core.QItemSelection))(signal))(selected, deselected)
				f(selected, deselected)
			}
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTreeView) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectionChanged")
	}
}

func (ptr *QTreeView) SelectionChanged(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_SelectionChanged(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

func (ptr *QTreeView) SelectionChangedDefault(selected core.QItemSelection_ITF, deselected core.QItemSelection_ITF) {
	if ptr.Pointer() != nil {
		C.QTreeView_SelectionChangedDefault(ptr.Pointer(), core.PointerFromQItemSelection(selected), core.PointerFromQItemSelection(deselected))
	}
}

//export callbackQTreeView_SetSelection
func callbackQTreeView_SetSelection(ptr unsafe.Pointer, rect unsafe.Pointer, command C.longlong) {
	if signal := qt.GetSignal(ptr, "setSelection"); signal != nil {
		(*(*func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(signal))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQTreeViewFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QTreeView) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSelection"); signal != nil {
			f := func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag) {
				(*(*func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(signal))(rect, command)
				f(rect, command)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSelection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSelection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTreeView) DisconnectSetSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSelection")
	}
}

func (ptr *QTreeView) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

func (ptr *QTreeView) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	if ptr.Pointer() != nil {
		C.QTreeView_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.longlong(command))
	}
}

//export callbackQTreeView_DestroyQTreeView
func callbackQTreeView_DestroyQTreeView(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTreeView"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTreeViewFromPointer(ptr).DestroyQTreeViewDefault()
	}
}

func (ptr *QTreeView) ConnectDestroyQTreeView(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTreeView"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTreeView", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTreeView", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTreeView) DisconnectDestroyQTreeView() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTreeView")
	}
}

func (ptr *QTreeView) DestroyQTreeView() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTreeView_DestroyQTreeView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTreeView) DestroyQTreeViewDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTreeView_DestroyQTreeViewDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQTreeView_HorizontalOffset
func callbackQTreeView_HorizontalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "horizontalOffset"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQTreeViewFromPointer(ptr).HorizontalOffsetDefault()))
}

func (ptr *QTreeView) HorizontalOffset() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTreeView_HorizontalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeView) HorizontalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTreeView_HorizontalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQTreeView_IsIndexHidden
func callbackQTreeView_IsIndexHidden(ptr unsafe.Pointer, index unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isIndexHidden"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTreeViewFromPointer(ptr).IsIndexHiddenDefault(core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QTreeView) IsIndexHidden(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTreeView_IsIndexHidden(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QTreeView) IsIndexHiddenDefault(index core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTreeView_IsIndexHiddenDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

//export callbackQTreeView_MoveCursor
func callbackQTreeView_MoveCursor(ptr unsafe.Pointer, cursorAction C.longlong, modifiers C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "moveCursor"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(QAbstractItemView__CursorAction, core.Qt__KeyboardModifier) *core.QModelIndex)(signal))(QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
	}

	return core.PointerFromQModelIndex(NewQTreeViewFromPointer(ptr).MoveCursorDefault(QAbstractItemView__CursorAction(cursorAction), core.Qt__KeyboardModifier(modifiers)))
}

func (ptr *QTreeView) MoveCursor(cursorAction QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QTreeView_MoveCursor(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QTreeView) MoveCursorDefault(cursorAction QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QTreeView_MoveCursorDefault(ptr.Pointer(), C.longlong(cursorAction), C.longlong(modifiers)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQTreeView_ScrollTo
func callbackQTreeView_ScrollTo(ptr unsafe.Pointer, index unsafe.Pointer, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "scrollTo"); signal != nil {
		(*(*func(*core.QModelIndex, QAbstractItemView__ScrollHint))(signal))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	} else {
		NewQTreeViewFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QTreeView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QTreeView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

func (ptr *QTreeView) ScrollToDefault(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QTreeView_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.longlong(hint))
	}
}

//export callbackQTreeView_VerticalOffset
func callbackQTreeView_VerticalOffset(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "verticalOffset"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQTreeViewFromPointer(ptr).VerticalOffsetDefault()))
}

func (ptr *QTreeView) VerticalOffset() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTreeView_VerticalOffset(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTreeView) VerticalOffsetDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTreeView_VerticalOffsetDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQTreeView_VisualRect
func callbackQTreeView_VisualRect(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRect"); signal != nil {
		return core.PointerFromQRect((*(*func(*core.QModelIndex) *core.QRect)(signal))(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQRect(NewQTreeViewFromPointer(ptr).VisualRectDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QTreeView) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QTreeView_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QTreeView) VisualRectDefault(index core.QModelIndex_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QTreeView_VisualRectDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQTreeView_VisualRegionForSelection
func callbackQTreeView_VisualRegionForSelection(ptr unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "visualRegionForSelection"); signal != nil {
		return gui.PointerFromQRegion((*(*func(*core.QItemSelection) *gui.QRegion)(signal))(core.NewQItemSelectionFromPointer(selection)))
	}

	return gui.PointerFromQRegion(NewQTreeViewFromPointer(ptr).VisualRegionForSelectionDefault(core.NewQItemSelectionFromPointer(selection)))
}

func (ptr *QTreeView) VisualRegionForSelection(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQRegionFromPointer(C.QTreeView_VisualRegionForSelection(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		qt.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

func (ptr *QTreeView) VisualRegionForSelectionDefault(selection core.QItemSelection_ITF) *gui.QRegion {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQRegionFromPointer(C.QTreeView_VisualRegionForSelectionDefault(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
		qt.SetFinalizer(tmpValue, (*gui.QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

// QTreeWidgetItem::ChildIndicatorPolicy
//
//go:generate stringer -type=QTreeWidgetItem__ChildIndicatorPolicy
type QTreeWidgetItem__ChildIndicatorPolicy int64

const (
	QTreeWidgetItem__ShowIndicator                  QTreeWidgetItem__ChildIndicatorPolicy = QTreeWidgetItem__ChildIndicatorPolicy(0)
	QTreeWidgetItem__DontShowIndicator              QTreeWidgetItem__ChildIndicatorPolicy = QTreeWidgetItem__ChildIndicatorPolicy(1)
	QTreeWidgetItem__DontShowIndicatorWhenChildless QTreeWidgetItem__ChildIndicatorPolicy = QTreeWidgetItem__ChildIndicatorPolicy(2)
)

// QTreeWidgetItem::ItemType
//
//go:generate stringer -type=QTreeWidgetItem__ItemType
type QTreeWidgetItem__ItemType int64

const (
	QTreeWidgetItem__Type     QTreeWidgetItem__ItemType = QTreeWidgetItem__ItemType(0)
	QTreeWidgetItem__UserType QTreeWidgetItem__ItemType = QTreeWidgetItem__ItemType(1000)
)

// QTreeWidgetItemIterator::IteratorFlag
//
//go:generate stringer -type=QTreeWidgetItemIterator__IteratorFlag
type QTreeWidgetItemIterator__IteratorFlag int64

const (
	QTreeWidgetItemIterator__All           QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000000)
	QTreeWidgetItemIterator__Hidden        QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000001)
	QTreeWidgetItemIterator__NotHidden     QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000002)
	QTreeWidgetItemIterator__Selected      QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000004)
	QTreeWidgetItemIterator__Unselected    QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000008)
	QTreeWidgetItemIterator__Selectable    QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000010)
	QTreeWidgetItemIterator__NotSelectable QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000020)
	QTreeWidgetItemIterator__DragEnabled   QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000040)
	QTreeWidgetItemIterator__DragDisabled  QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000080)
	QTreeWidgetItemIterator__DropEnabled   QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000100)
	QTreeWidgetItemIterator__DropDisabled  QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000200)
	QTreeWidgetItemIterator__HasChildren   QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000400)
	QTreeWidgetItemIterator__NoChildren    QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00000800)
	QTreeWidgetItemIterator__Checked       QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00001000)
	QTreeWidgetItemIterator__NotChecked    QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00002000)
	QTreeWidgetItemIterator__Enabled       QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00004000)
	QTreeWidgetItemIterator__Disabled      QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00008000)
	QTreeWidgetItemIterator__Editable      QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00010000)
	QTreeWidgetItemIterator__NotEditable   QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x00020000)
	QTreeWidgetItemIterator__UserFlag      QTreeWidgetItemIterator__IteratorFlag = QTreeWidgetItemIterator__IteratorFlag(0x01000000)
)

type QVBoxLayout struct {
	QBoxLayout
}

type QVBoxLayout_ITF interface {
	QBoxLayout_ITF
	QVBoxLayout_PTR() *QVBoxLayout
}

func (ptr *QVBoxLayout) QVBoxLayout_PTR() *QVBoxLayout {
	return ptr
}

func (ptr *QVBoxLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QBoxLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QVBoxLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QBoxLayout_PTR().SetPointer(p)
	}
}

func PointerFromQVBoxLayout(ptr QVBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQVBoxLayoutFromPointer(ptr unsafe.Pointer) (n *QVBoxLayout) {
	n = new(QVBoxLayout)
	n.SetPointer(ptr)
	return
}
func NewQVBoxLayout() *QVBoxLayout {
	tmpValue := NewQVBoxLayoutFromPointer(C.QVBoxLayout_NewQVBoxLayout())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQVBoxLayout2(parent QWidget_ITF) *QVBoxLayout {
	tmpValue := NewQVBoxLayoutFromPointer(C.QVBoxLayout_NewQVBoxLayout2(PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQVBoxLayout_DestroyQVBoxLayout
func callbackQVBoxLayout_DestroyQVBoxLayout(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QVBoxLayout"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQVBoxLayoutFromPointer(ptr).DestroyQVBoxLayoutDefault()
	}
}

func (ptr *QVBoxLayout) ConnectDestroyQVBoxLayout(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QVBoxLayout"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QVBoxLayout", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QVBoxLayout", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QVBoxLayout) DisconnectDestroyQVBoxLayout() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QVBoxLayout")
	}
}

func (ptr *QVBoxLayout) DestroyQVBoxLayout() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVBoxLayout_DestroyQVBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVBoxLayout) DestroyQVBoxLayoutDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QVBoxLayout_DestroyQVBoxLayoutDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWidget struct {
	core.QObject
	gui.QPaintDevice
}

type QWidget_ITF interface {
	core.QObject_ITF
	gui.QPaintDevice_ITF
	QWidget_PTR() *QWidget
}

func (ptr *QWidget) QWidget_PTR() *QWidget {
	return ptr
}

func (ptr *QWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQWidget(ptr QWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func NewQWidgetFromPointer(ptr unsafe.Pointer) (n *QWidget) {
	n = new(QWidget)
	n.SetPointer(ptr)
	return
}

// QWidget::RenderFlag
//
//go:generate stringer -type=QWidget__RenderFlag
type QWidget__RenderFlag int64

const (
	QWidget__DrawWindowBackground QWidget__RenderFlag = QWidget__RenderFlag(0x1)
	QWidget__DrawChildren         QWidget__RenderFlag = QWidget__RenderFlag(0x2)
	QWidget__IgnoreMask           QWidget__RenderFlag = QWidget__RenderFlag(0x4)
)

func NewQWidget(parent QWidget_ITF, ff core.Qt__WindowType) *QWidget {
	tmpValue := NewQWidgetFromPointer(C.QWidget_NewQWidget(PointerFromQWidget(parent), C.longlong(ff)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWidget) AccessibleDescription() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWidget_AccessibleDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) Actions() []*QAction {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWidgets_PackedList) []*QAction {
			out := make([]*QAction, int(l.len))
			tmpList := NewQWidgetFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__actions_atList(i)
			}
			return out
		}(C.QWidget_Actions(ptr.Pointer()))
	}
	return make([]*QAction, 0)
}

func (ptr *QWidget) ActivateWindow() {
	if ptr.Pointer() != nil {
		C.QWidget_ActivateWindow(ptr.Pointer())
	}
}

func (ptr *QWidget) AddAction(action QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_AddAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

//export callbackQWidget_Close
func callbackQWidget_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QWidget) ConnectClose(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QWidget) Close() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWidget_Close(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWidget_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQWidget_CloseEvent
func callbackQWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*gui.QCloseEvent))(signal))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "closeEvent"); signal != nil {
			f := func(event *gui.QCloseEvent) {
				(*(*func(*gui.QCloseEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "closeEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "closeEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectCloseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "closeEvent")
	}
}

func (ptr *QWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQWidget_ContextMenuEvent
func callbackQWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*gui.QContextMenuEvent))(signal))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contextMenuEvent"); signal != nil {
			f := func(event *gui.QContextMenuEvent) {
				(*(*func(*gui.QContextMenuEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "contextMenuEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contextMenuEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectContextMenuEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "contextMenuEvent")
	}
}

func (ptr *QWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QWidget) Create(window uintptr, initializeWindow bool, destroyOldWindow bool) {
	if ptr.Pointer() != nil {
		C.QWidget_Create(ptr.Pointer(), C.uintptr_t(window), C.char(int8(qt.GoBoolToInt(initializeWindow))), C.char(int8(qt.GoBoolToInt(destroyOldWindow))))
	}
}

func (ptr *QWidget) Cursor() *gui.QCursor {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQCursorFromPointer(C.QWidget_Cursor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QCursor).DestroyQCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) Destroy(destroyWindow bool, destroySubWindows bool) {
	if ptr.Pointer() != nil {
		C.QWidget_Destroy(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(destroyWindow))), C.char(int8(qt.GoBoolToInt(destroySubWindows))))
	}
}

//export callbackQWidget_Event
func callbackQWidget_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWidgetFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QWidget) ConnectEvent(f func(event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "event"); signal != nil {
			f := func(event *core.QEvent) bool {
				(*(*func(*core.QEvent) bool)(signal))(event)
				return f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "event", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "event", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "event")
	}
}

func (ptr *QWidget) Event(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWidget_Event(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QWidget) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWidget_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func QWidget_Find(id uintptr) *QWidget {
	tmpValue := NewQWidgetFromPointer(C.QWidget_QWidget_Find(C.uintptr_t(id)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWidget) Find(id uintptr) *QWidget {
	tmpValue := NewQWidgetFromPointer(C.QWidget_QWidget_Find(C.uintptr_t(id)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWidget) Font() *gui.QFont {
	if ptr.Pointer() != nil {
		return gui.NewQFontFromPointer(C.QWidget_Font(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) FontMetrics() *gui.QFontMetrics {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQFontMetricsFromPointer(C.QWidget_FontMetrics(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QFontMetrics).DestroyQFontMetrics)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) Geometry() *core.QRect {
	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWidget_Geometry(ptr.Pointer()))
	}
	return nil
}

//export callbackQWidget_HasHeightForWidth
func callbackQWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QWidget) ConnectHasHeightForWidth(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasHeightForWidth"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hasHeightForWidth", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasHeightForWidth", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectHasHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasHeightForWidth")
	}
}

func (ptr *QWidget) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWidget_HasHeightForWidth(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWidget_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) Height() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_Height(ptr.Pointer())))
	}
	return 0
}

//export callbackQWidget_HeightForWidth
func callbackQWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewQWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QWidget) ConnectHeightForWidth(f func(w int) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "heightForWidth"); signal != nil {
			f := func(w int) int {
				(*(*func(int) int)(signal))(w)
				return f(w)
			}
			qt.ConnectSignal(ptr.Pointer(), "heightForWidth", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "heightForWidth", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectHeightForWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "heightForWidth")
	}
}

func (ptr *QWidget) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QWidget) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQWidget_Hide
func callbackQWidget_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QWidget) ConnectHide(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hide"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hide", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hide", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hide")
	}
}

func (ptr *QWidget) Hide() {
	if ptr.Pointer() != nil {
		C.QWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QWidget) HideDefault() {
	if ptr.Pointer() != nil {
		C.QWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQWidget_HideEvent
func callbackQWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*gui.QHideEvent))(signal))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hideEvent"); signal != nil {
			f := func(event *gui.QHideEvent) {
				(*(*func(*gui.QHideEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "hideEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hideEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hideEvent")
	}
}

func (ptr *QWidget) HideEvent(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQWidget_KeyPressEvent
func callbackQWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keyPressEvent"); signal != nil {
			f := func(event *gui.QKeyEvent) {
				(*(*func(*gui.QKeyEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "keyPressEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keyPressEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "keyPressEvent")
	}
}

func (ptr *QWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWidget_KeyReleaseEvent
func callbackQWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*gui.QKeyEvent))(signal))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keyReleaseEvent"); signal != nil {
			f := func(event *gui.QKeyEvent) {
				(*(*func(*gui.QKeyEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "keyReleaseEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keyReleaseEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "keyReleaseEvent")
	}
}

func (ptr *QWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWidget) Layout() *QLayout {
	if ptr.Pointer() != nil {
		tmpValue := NewQLayoutFromPointer(C.QWidget_Layout(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWidget_Lower
func callbackQWidget_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QWidget) ConnectLower(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lower"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lower", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lower", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lower")
	}
}

func (ptr *QWidget) Lower() {
	if ptr.Pointer() != nil {
		C.QWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QWidget) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QWidget_LowerDefault(ptr.Pointer())
	}
}

func (ptr *QWidget) MaximumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWidget_MaximumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) MaximumWidth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_MaximumWidth(ptr.Pointer())))
	}
	return 0
}

//export callbackQWidget_Metric
func callbackQWidget_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(gui.QPaintDevice__PaintDeviceMetric) int)(signal))(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQWidgetFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QWidget) ConnectMetric(f func(m gui.QPaintDevice__PaintDeviceMetric) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "metric"); signal != nil {
			f := func(m gui.QPaintDevice__PaintDeviceMetric) int {
				(*(*func(gui.QPaintDevice__PaintDeviceMetric) int)(signal))(m)
				return f(m)
			}
			qt.ConnectSignal(ptr.Pointer(), "metric", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "metric", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectMetric() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "metric")
	}
}

func (ptr *QWidget) Metric(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_Metric(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

func (ptr *QWidget) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

func (ptr *QWidget) MinimumHeight() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_MinimumHeight(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) MinimumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWidget_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) MinimumWidth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_MinimumWidth(ptr.Pointer())))
	}
	return 0
}

//export callbackQWidget_MouseDoubleClickEvent
func callbackQWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mouseDoubleClickEvent"); signal != nil {
			f := func(event *gui.QMouseEvent) {
				(*(*func(*gui.QMouseEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "mouseDoubleClickEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mouseDoubleClickEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mouseDoubleClickEvent")
	}
}

func (ptr *QWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWidget_MousePressEvent
func callbackQWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mousePressEvent"); signal != nil {
			f := func(event *gui.QMouseEvent) {
				(*(*func(*gui.QMouseEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "mousePressEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mousePressEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mousePressEvent")
	}
}

func (ptr *QWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWidget_MouseReleaseEvent
func callbackQWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*gui.QMouseEvent))(signal))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mouseReleaseEvent"); signal != nil {
			f := func(event *gui.QMouseEvent) {
				(*(*func(*gui.QMouseEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "mouseReleaseEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mouseReleaseEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mouseReleaseEvent")
	}
}

func (ptr *QWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWidget) Move(vqp core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Move(ptr.Pointer(), core.PointerFromQPoint(vqp))
	}
}

func (ptr *QWidget) Move2(x int, y int) {
	if ptr.Pointer() != nil {
		C.QWidget_Move2(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)))
	}
}

func (ptr *QWidget) Pos() *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QWidget_Pos(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) Rect() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QWidget_Rect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) RemoveAction(action QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_RemoveAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QWidget) Resize(vqs core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Resize(ptr.Pointer(), core.PointerFromQSize(vqs))
	}
}

func (ptr *QWidget) Resize2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_Resize2(ptr.Pointer(), C.int(int32(w)), C.int(int32(h)))
	}
}

//export callbackQWidget_ResizeEvent
func callbackQWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*gui.QResizeEvent))(signal))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resizeEvent"); signal != nil {
			f := func(event *gui.QResizeEvent) {
				(*(*func(*gui.QResizeEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "resizeEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resizeEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resizeEvent")
	}
}

func (ptr *QWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWidget) Scroll(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QWidget_Scroll(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

func (ptr *QWidget) Scroll2(dx int, dy int, r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Scroll2(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)), core.PointerFromQRect(r))
	}
}

func (ptr *QWidget) SetAccessibleDescription(description string) {
	if ptr.Pointer() != nil {
		var descriptionC *C.char
		if description != "" {
			descriptionC = C.CString(description)
			defer C.free(unsafe.Pointer(descriptionC))
		}
		C.QWidget_SetAccessibleDescription(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: descriptionC, len: C.longlong(len(description))})
	}
}

//export callbackQWidget_SetDisabled
func callbackQWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewQWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QWidget) ConnectSetDisabled(f func(disable bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDisabled"); signal != nil {
			f := func(disable bool) {
				(*(*func(bool))(signal))(disable)
				f(disable)
			}
			qt.ConnectSignal(ptr.Pointer(), "setDisabled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDisabled", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectSetDisabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDisabled")
	}
}

func (ptr *QWidget) SetDisabled(disable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QWidget) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQWidget_SetEnabled
func callbackQWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewQWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QWidget) ConnectSetEnabled(f func(vbo bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEnabled"); signal != nil {
			f := func(vbo bool) {
				(*(*func(bool))(signal))(vbo)
				f(vbo)
			}
			qt.ConnectSignal(ptr.Pointer(), "setEnabled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEnabled", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectSetEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEnabled")
	}
}

func (ptr *QWidget) SetEnabled(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QWidget) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QWidget) SetFixedSize(s core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedSize(ptr.Pointer(), core.PointerFromQSize(s))
	}
}

func (ptr *QWidget) SetFixedSize2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedSize2(ptr.Pointer(), C.int(int32(w)), C.int(int32(h)))
	}
}

func (ptr *QWidget) SetFixedWidth(w int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedWidth(ptr.Pointer(), C.int(int32(w)))
	}
}

func (ptr *QWidget) SetFocus(reason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocus(ptr.Pointer(), C.longlong(reason))
	}
}

//export callbackQWidget_SetFocus2
func callbackQWidget_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QWidget) ConnectSetFocus2(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFocus2"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "setFocus2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFocus2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectSetFocus2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFocus2")
	}
}

func (ptr *QWidget) SetFocus2() {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QWidget) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocus2Default(ptr.Pointer())
	}
}

func (ptr *QWidget) SetFont(vqf gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFont(ptr.Pointer(), gui.PointerFromQFont(vqf))
	}
}

func (ptr *QWidget) SetGeometry(vqr core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetGeometry(ptr.Pointer(), core.PointerFromQRect(vqr))
	}
}

func (ptr *QWidget) SetGeometry2(x int, y int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetGeometry2(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(w)), C.int(int32(h)))
	}
}

func (ptr *QWidget) SetLayout(layout QLayout_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetLayout(ptr.Pointer(), PointerFromQLayout(layout))
	}
}

func (ptr *QWidget) SetMaximumWidth(maxw int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumWidth(ptr.Pointer(), C.int(int32(maxw)))
	}
}

func (ptr *QWidget) SetMinimumSize(vqs core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumSize(ptr.Pointer(), core.PointerFromQSize(vqs))
	}
}

func (ptr *QWidget) SetMinimumSize2(minw int, minh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumSize2(ptr.Pointer(), C.int(int32(minw)), C.int(int32(minh)))
	}
}

func (ptr *QWidget) SetMinimumWidth(minw int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumWidth(ptr.Pointer(), C.int(int32(minw)))
	}
}

func (ptr *QWidget) SetSizePolicy(vqs QSizePolicy_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetSizePolicy(ptr.Pointer(), PointerFromQSizePolicy(vqs))
	}
}

func (ptr *QWidget) SetSizePolicy2(horizontal QSizePolicy__Policy, vertical QSizePolicy__Policy) {
	if ptr.Pointer() != nil {
		C.QWidget_SetSizePolicy2(ptr.Pointer(), C.longlong(horizontal), C.longlong(vertical))
	}
}

func (ptr *QWidget) SetStyle(style QStyle_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetStyle(ptr.Pointer(), PointerFromQStyle(style))
	}
}

//export callbackQWidget_SetStyleSheet
func callbackQWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewQWidgetFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setStyleSheet"); signal != nil {
			f := func(styleSheet string) {
				(*(*func(string))(signal))(styleSheet)
				f(styleSheet)
			}
			qt.ConnectSignal(ptr.Pointer(), "setStyleSheet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setStyleSheet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectSetStyleSheet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setStyleSheet")
	}
}

func (ptr *QWidget) SetStyleSheet(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QWidget_SetStyleSheet(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

func (ptr *QWidget) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QWidget_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

func (ptr *QWidget) SetToolTip(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QWidget_SetToolTip(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QWidget) SetWindowIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

//export callbackQWidget_SetWindowTitle
func callbackQWidget_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewQWidgetFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QWidget) ConnectSetWindowTitle(f func(vqs string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setWindowTitle"); signal != nil {
			f := func(vqs string) {
				(*(*func(string))(signal))(vqs)
				f(vqs)
			}
			qt.ConnectSignal(ptr.Pointer(), "setWindowTitle", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setWindowTitle", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectSetWindowTitle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setWindowTitle")
	}
}

func (ptr *QWidget) SetWindowTitle(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QWidget_SetWindowTitle(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QWidget) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QWidget_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtWidgets_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQWidget_Show
func callbackQWidget_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWidget) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "show"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "show", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "show", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "show")
	}
}

func (ptr *QWidget) Show() {
	if ptr.Pointer() != nil {
		C.QWidget_Show(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQWidget_ShowEvent
func callbackQWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*gui.QShowEvent))(signal))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "showEvent"); signal != nil {
			f := func(event *gui.QShowEvent) {
				(*(*func(*gui.QShowEvent))(signal))(event)
				f(event)
			}
			qt.ConnectSignal(ptr.Pointer(), "showEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "showEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "showEvent")
	}
}

func (ptr *QWidget) ShowEvent(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQWidget_ShowMaximized
func callbackQWidget_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QWidget) ConnectShowMaximized(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "showMaximized"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "showMaximized", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "showMaximized", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectShowMaximized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "showMaximized")
	}
}

func (ptr *QWidget) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

func (ptr *QWidget) Size() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWidget_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQWidget_SizeHint
func callbackQWidget_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QWidget) ConnectSizeHint(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sizeHint"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sizeHint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectSizeHint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sizeHint")
	}
}

func (ptr *QWidget) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWidget_SizeHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWidget_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) SizePolicy() *QSizePolicy {
	if ptr.Pointer() != nil {
		tmpValue := NewQSizePolicyFromPointer(C.QWidget_SizePolicy(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSizePolicy).DestroyQSizePolicy)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) Style() *QStyle {
	if ptr.Pointer() != nil {
		tmpValue := NewQStyleFromPointer(C.QWidget_Style(ptr.Pointer()))
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) StyleSheet() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWidget_StyleSheet(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) ToolTip() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWidget_ToolTip(ptr.Pointer()))
	}
	return ""
}

//export callbackQWidget_Update
func callbackQWidget_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QWidget) ConnectUpdate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "update"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "update", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "update", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectUpdate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "update")
	}
}

func (ptr *QWidget) Update() {
	if ptr.Pointer() != nil {
		C.QWidget_Update(ptr.Pointer())
	}
}

func (ptr *QWidget) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QWidget_UpdateDefault(ptr.Pointer())
	}
}

func (ptr *QWidget) Update2(x int, y int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_Update2(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(w)), C.int(int32(h)))
	}
}

func (ptr *QWidget) Update3(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Update3(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWidget) Update4(rgn gui.QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Update4(ptr.Pointer(), gui.PointerFromQRegion(rgn))
	}
}

func (ptr *QWidget) Width() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_Width(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) Window() *QWidget {
	if ptr.Pointer() != nil {
		tmpValue := NewQWidgetFromPointer(C.QWidget_Window(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) WindowIcon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QWidget_WindowIcon(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) WindowTitle() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWidget_WindowTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) X() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_X(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) Y() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWidget_Y(ptr.Pointer())))
	}
	return 0
}

//export callbackQWidget_DestroyQWidget
func callbackQWidget_DestroyQWidget(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWidget"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWidgetFromPointer(ptr).DestroyQWidgetDefault()
	}
}

func (ptr *QWidget) ConnectDestroyQWidget(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWidget"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWidget", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWidget", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWidget) DisconnectDestroyQWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWidget")
	}
}

func (ptr *QWidget) DestroyQWidget() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWidget_DestroyQWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWidget) DestroyQWidgetDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWidget_DestroyQWidgetDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWidget) __actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QWidget___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QWidget) __actions_newList() unsafe.Pointer {
	return C.QWidget___actions_newList(ptr.Pointer())
}

func (ptr *QWidget) __addActions_actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QWidget___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __addActions_actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___addActions_actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QWidget) __addActions_actions_newList() unsafe.Pointer {
	return C.QWidget___addActions_actions_newList(ptr.Pointer())
}

func (ptr *QWidget) __insertActions_actions_atList(i int) *QAction {
	if ptr.Pointer() != nil {
		tmpValue := NewQActionFromPointer(C.QWidget___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __insertActions_actions_setList(i QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___insertActions_actions_setList(ptr.Pointer(), PointerFromQAction(i))
	}
}

func (ptr *QWidget) __insertActions_actions_newList() unsafe.Pointer {
	return C.QWidget___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *QWidget) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWidget___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWidget) __children_newList() unsafe.Pointer {
	return C.QWidget___children_newList(ptr.Pointer())
}

func (ptr *QWidget) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWidget___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWidget) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWidget___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWidget) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWidget___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWidget) __findChildren_newList() unsafe.Pointer {
	return C.QWidget___findChildren_newList(ptr.Pointer())
}

func (ptr *QWidget) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWidget___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWidget) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWidget) __findChildren_newList3() unsafe.Pointer {
	return C.QWidget___findChildren_newList3(ptr.Pointer())
}

//export callbackQWidget_ChildEvent
func callbackQWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWidget) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWidget_ConnectNotify
func callbackQWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWidget_CustomEvent
func callbackQWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWidget) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWidget) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWidget_DeleteLater
func callbackQWidget_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWidget) DeleteLater() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWidget_DeleteLater(ptr.Pointer())
	}
}

func (ptr *QWidget) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWidget_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQWidget_Destroyed
func callbackQWidget_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWidget_DisconnectNotify
func callbackQWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWidget_EventFilter
func callbackQWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWidget_ObjectNameChanged
func callbackQWidget_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWidgets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWidget_TimerEvent
func callbackQWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWidget) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWidget_PaintEngine
func callbackQWidget_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine((*(*func() *gui.QPaintEngine)(signal))())
	}

	return gui.PointerFromQPaintEngine(NewQWidgetFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QWidget) PaintEngine() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWidget_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWidget_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

// QWizard::WizardButton
//
//go:generate stringer -type=QWizard__WizardButton
type QWizard__WizardButton int64

const (
	QWizard__BackButton       QWizard__WizardButton = QWizard__WizardButton(0)
	QWizard__NextButton       QWizard__WizardButton = QWizard__WizardButton(1)
	QWizard__CommitButton     QWizard__WizardButton = QWizard__WizardButton(2)
	QWizard__FinishButton     QWizard__WizardButton = QWizard__WizardButton(3)
	QWizard__CancelButton     QWizard__WizardButton = QWizard__WizardButton(4)
	QWizard__HelpButton       QWizard__WizardButton = QWizard__WizardButton(5)
	QWizard__CustomButton1    QWizard__WizardButton = QWizard__WizardButton(6)
	QWizard__CustomButton2    QWizard__WizardButton = QWizard__WizardButton(7)
	QWizard__CustomButton3    QWizard__WizardButton = QWizard__WizardButton(8)
	QWizard__Stretch          QWizard__WizardButton = QWizard__WizardButton(9)
	QWizard__NoButton         QWizard__WizardButton = QWizard__WizardButton(-1)
	QWizard__NStandardButtons QWizard__WizardButton = QWizard__WizardButton(6)
	QWizard__NButtons         QWizard__WizardButton = QWizard__WizardButton(9)
)

// QWizard::WizardOption
//
//go:generate stringer -type=QWizard__WizardOption
type QWizard__WizardOption int64

const (
	QWizard__IndependentPages             QWizard__WizardOption = QWizard__WizardOption(0x00000001)
	QWizard__IgnoreSubTitles              QWizard__WizardOption = QWizard__WizardOption(0x00000002)
	QWizard__ExtendedWatermarkPixmap      QWizard__WizardOption = QWizard__WizardOption(0x00000004)
	QWizard__NoDefaultButton              QWizard__WizardOption = QWizard__WizardOption(0x00000008)
	QWizard__NoBackButtonOnStartPage      QWizard__WizardOption = QWizard__WizardOption(0x00000010)
	QWizard__NoBackButtonOnLastPage       QWizard__WizardOption = QWizard__WizardOption(0x00000020)
	QWizard__DisabledBackButtonOnLastPage QWizard__WizardOption = QWizard__WizardOption(0x00000040)
	QWizard__HaveNextButtonOnLastPage     QWizard__WizardOption = QWizard__WizardOption(0x00000080)
	QWizard__HaveFinishButtonOnEarlyPages QWizard__WizardOption = QWizard__WizardOption(0x00000100)
	QWizard__NoCancelButton               QWizard__WizardOption = QWizard__WizardOption(0x00000200)
	QWizard__CancelButtonOnLeft           QWizard__WizardOption = QWizard__WizardOption(0x00000400)
	QWizard__HaveHelpButton               QWizard__WizardOption = QWizard__WizardOption(0x00000800)
	QWizard__HelpButtonOnRight            QWizard__WizardOption = QWizard__WizardOption(0x00001000)
	QWizard__HaveCustomButton1            QWizard__WizardOption = QWizard__WizardOption(0x00002000)
	QWizard__HaveCustomButton2            QWizard__WizardOption = QWizard__WizardOption(0x00004000)
	QWizard__HaveCustomButton3            QWizard__WizardOption = QWizard__WizardOption(0x00008000)
	QWizard__NoCancelButtonOnLastPage     QWizard__WizardOption = QWizard__WizardOption(0x00010000)
)

// QWizard::WizardPixmap
//
//go:generate stringer -type=QWizard__WizardPixmap
type QWizard__WizardPixmap int64

const (
	QWizard__WatermarkPixmap  QWizard__WizardPixmap = QWizard__WizardPixmap(0)
	QWizard__LogoPixmap       QWizard__WizardPixmap = QWizard__WizardPixmap(1)
	QWizard__BannerPixmap     QWizard__WizardPixmap = QWizard__WizardPixmap(2)
	QWizard__BackgroundPixmap QWizard__WizardPixmap = QWizard__WizardPixmap(3)
	QWizard__NPixmaps         QWizard__WizardPixmap = QWizard__WizardPixmap(4)
)

// QWizard::WizardStyle
//
//go:generate stringer -type=QWizard__WizardStyle
type QWizard__WizardStyle int64

var (
	QWizard__ClassicStyle QWizard__WizardStyle = QWizard__WizardStyle(0)
	QWizard__ModernStyle  QWizard__WizardStyle = QWizard__WizardStyle(1)
	QWizard__MacStyle     QWizard__WizardStyle = QWizard__WizardStyle(2)
	QWizard__AeroStyle    QWizard__WizardStyle = QWizard__WizardStyle(3)
	QWizard__NStyles      QWizard__WizardStyle = QWizard__WizardStyle(4)
)

func init() {
}
