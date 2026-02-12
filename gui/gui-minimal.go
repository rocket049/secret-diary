//go:build minimal
// +build minimal

package gui

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "gui-minimal.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtGui_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtGui_PackedString) []byte {
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

// QAccessible::Event
//
//go:generate stringer -type=QAccessible__Event
type QAccessible__Event int64

const (
	QAccessible__SoundPlayed                     QAccessible__Event = QAccessible__Event(0x0001)
	QAccessible__Alert                           QAccessible__Event = QAccessible__Event(0x0002)
	QAccessible__ForegroundChanged               QAccessible__Event = QAccessible__Event(0x0003)
	QAccessible__MenuStart                       QAccessible__Event = QAccessible__Event(0x0004)
	QAccessible__MenuEnd                         QAccessible__Event = QAccessible__Event(0x0005)
	QAccessible__PopupMenuStart                  QAccessible__Event = QAccessible__Event(0x0006)
	QAccessible__PopupMenuEnd                    QAccessible__Event = QAccessible__Event(0x0007)
	QAccessible__ContextHelpStart                QAccessible__Event = QAccessible__Event(0x000C)
	QAccessible__ContextHelpEnd                  QAccessible__Event = QAccessible__Event(0x000D)
	QAccessible__DragDropStart                   QAccessible__Event = QAccessible__Event(0x000E)
	QAccessible__DragDropEnd                     QAccessible__Event = QAccessible__Event(0x000F)
	QAccessible__DialogStart                     QAccessible__Event = QAccessible__Event(0x0010)
	QAccessible__DialogEnd                       QAccessible__Event = QAccessible__Event(0x0011)
	QAccessible__ScrollingStart                  QAccessible__Event = QAccessible__Event(0x0012)
	QAccessible__ScrollingEnd                    QAccessible__Event = QAccessible__Event(0x0013)
	QAccessible__MenuCommand                     QAccessible__Event = QAccessible__Event(0x0018)
	QAccessible__ActionChanged                   QAccessible__Event = QAccessible__Event(0x0101)
	QAccessible__ActiveDescendantChanged         QAccessible__Event = QAccessible__Event(0x0102)
	QAccessible__AttributeChanged                QAccessible__Event = QAccessible__Event(0x0103)
	QAccessible__DocumentContentChanged          QAccessible__Event = QAccessible__Event(0x0104)
	QAccessible__DocumentLoadComplete            QAccessible__Event = QAccessible__Event(0x0105)
	QAccessible__DocumentLoadStopped             QAccessible__Event = QAccessible__Event(0x0106)
	QAccessible__DocumentReload                  QAccessible__Event = QAccessible__Event(0x0107)
	QAccessible__HyperlinkEndIndexChanged        QAccessible__Event = QAccessible__Event(0x0108)
	QAccessible__HyperlinkNumberOfAnchorsChanged QAccessible__Event = QAccessible__Event(0x0109)
	QAccessible__HyperlinkSelectedLinkChanged    QAccessible__Event = QAccessible__Event(0x010A)
	QAccessible__HypertextLinkActivated          QAccessible__Event = QAccessible__Event(0x010B)
	QAccessible__HypertextLinkSelected           QAccessible__Event = QAccessible__Event(0x010C)
	QAccessible__HyperlinkStartIndexChanged      QAccessible__Event = QAccessible__Event(0x010D)
	QAccessible__HypertextChanged                QAccessible__Event = QAccessible__Event(0x010E)
	QAccessible__HypertextNLinksChanged          QAccessible__Event = QAccessible__Event(0x010F)
	QAccessible__ObjectAttributeChanged          QAccessible__Event = QAccessible__Event(0x0110)
	QAccessible__PageChanged                     QAccessible__Event = QAccessible__Event(0x0111)
	QAccessible__SectionChanged                  QAccessible__Event = QAccessible__Event(0x0112)
	QAccessible__TableCaptionChanged             QAccessible__Event = QAccessible__Event(0x0113)
	QAccessible__TableColumnDescriptionChanged   QAccessible__Event = QAccessible__Event(0x0114)
	QAccessible__TableColumnHeaderChanged        QAccessible__Event = QAccessible__Event(0x0115)
	QAccessible__TableModelChanged               QAccessible__Event = QAccessible__Event(0x0116)
	QAccessible__TableRowDescriptionChanged      QAccessible__Event = QAccessible__Event(0x0117)
	QAccessible__TableRowHeaderChanged           QAccessible__Event = QAccessible__Event(0x0118)
	QAccessible__TableSummaryChanged             QAccessible__Event = QAccessible__Event(0x0119)
	QAccessible__TextAttributeChanged            QAccessible__Event = QAccessible__Event(0x011A)
	QAccessible__TextCaretMoved                  QAccessible__Event = QAccessible__Event(0x011B)
	QAccessible__TextColumnChanged               QAccessible__Event = QAccessible__Event(0x011D)
	QAccessible__TextInserted                    QAccessible__Event = QAccessible__Event(0x011E)
	QAccessible__TextRemoved                     QAccessible__Event = QAccessible__Event(0x011F)
	QAccessible__TextUpdated                     QAccessible__Event = QAccessible__Event(0x0120)
	QAccessible__TextSelectionChanged            QAccessible__Event = QAccessible__Event(0x0121)
	QAccessible__VisibleDataChanged              QAccessible__Event = QAccessible__Event(0x0122)
	QAccessible__ObjectCreated                   QAccessible__Event = QAccessible__Event(0x8000)
	QAccessible__ObjectDestroyed                 QAccessible__Event = QAccessible__Event(0x8001)
	QAccessible__ObjectShow                      QAccessible__Event = QAccessible__Event(0x8002)
	QAccessible__ObjectHide                      QAccessible__Event = QAccessible__Event(0x8003)
	QAccessible__ObjectReorder                   QAccessible__Event = QAccessible__Event(0x8004)
	QAccessible__Focus                           QAccessible__Event = QAccessible__Event(0x8005)
	QAccessible__Selection                       QAccessible__Event = QAccessible__Event(0x8006)
	QAccessible__SelectionAdd                    QAccessible__Event = QAccessible__Event(0x8007)
	QAccessible__SelectionRemove                 QAccessible__Event = QAccessible__Event(0x8008)
	QAccessible__SelectionWithin                 QAccessible__Event = QAccessible__Event(0x8009)
	QAccessible__StateChanged                    QAccessible__Event = QAccessible__Event(0x800A)
	QAccessible__LocationChanged                 QAccessible__Event = QAccessible__Event(0x800B)
	QAccessible__NameChanged                     QAccessible__Event = QAccessible__Event(0x800C)
	QAccessible__DescriptionChanged              QAccessible__Event = QAccessible__Event(0x800D)
	QAccessible__ValueChanged                    QAccessible__Event = QAccessible__Event(0x800E)
	QAccessible__ParentChanged                   QAccessible__Event = QAccessible__Event(0x800F)
	QAccessible__HelpChanged                     QAccessible__Event = QAccessible__Event(0x80A0)
	QAccessible__DefaultActionChanged            QAccessible__Event = QAccessible__Event(0x80B0)
	QAccessible__AcceleratorChanged              QAccessible__Event = QAccessible__Event(0x80C0)
	QAccessible__InvalidEvent                    QAccessible__Event = QAccessible__Event(0x80c1)
)

// QAccessible::InterfaceType
//
//go:generate stringer -type=QAccessible__InterfaceType
type QAccessible__InterfaceType int64

const (
	QAccessible__TextInterface         QAccessible__InterfaceType = QAccessible__InterfaceType(0)
	QAccessible__EditableTextInterface QAccessible__InterfaceType = QAccessible__InterfaceType(1)
	QAccessible__ValueInterface        QAccessible__InterfaceType = QAccessible__InterfaceType(2)
	QAccessible__ActionInterface       QAccessible__InterfaceType = QAccessible__InterfaceType(3)
	QAccessible__ImageInterface        QAccessible__InterfaceType = QAccessible__InterfaceType(4)
	QAccessible__TableInterface        QAccessible__InterfaceType = QAccessible__InterfaceType(5)
	QAccessible__TableCellInterface    QAccessible__InterfaceType = QAccessible__InterfaceType(6)
)

// QAccessible::RelationFlag
//
//go:generate stringer -type=QAccessible__RelationFlag
type QAccessible__RelationFlag int64

const (
	QAccessible__Label        QAccessible__RelationFlag = QAccessible__RelationFlag(0x00000001)
	QAccessible__Labelled     QAccessible__RelationFlag = QAccessible__RelationFlag(0x00000002)
	QAccessible__Controller   QAccessible__RelationFlag = QAccessible__RelationFlag(0x00000004)
	QAccessible__Controlled   QAccessible__RelationFlag = QAccessible__RelationFlag(0x00000008)
	QAccessible__AllRelations QAccessible__RelationFlag = QAccessible__RelationFlag(0xffffffff)
)

// QAccessible::Role
//
//go:generate stringer -type=QAccessible__Role
type QAccessible__Role int64

const (
	QAccessible__NoRole               QAccessible__Role = QAccessible__Role(0x00000000)
	QAccessible__TitleBar             QAccessible__Role = QAccessible__Role(0x00000001)
	QAccessible__MenuBar              QAccessible__Role = QAccessible__Role(0x00000002)
	QAccessible__ScrollBar            QAccessible__Role = QAccessible__Role(0x00000003)
	QAccessible__Grip                 QAccessible__Role = QAccessible__Role(0x00000004)
	QAccessible__Sound                QAccessible__Role = QAccessible__Role(0x00000005)
	QAccessible__Cursor               QAccessible__Role = QAccessible__Role(0x00000006)
	QAccessible__Caret                QAccessible__Role = QAccessible__Role(0x00000007)
	QAccessible__AlertMessage         QAccessible__Role = QAccessible__Role(0x00000008)
	QAccessible__Window               QAccessible__Role = QAccessible__Role(0x00000009)
	QAccessible__Client               QAccessible__Role = QAccessible__Role(0x0000000A)
	QAccessible__PopupMenu            QAccessible__Role = QAccessible__Role(0x0000000B)
	QAccessible__MenuItem             QAccessible__Role = QAccessible__Role(0x0000000C)
	QAccessible__ToolTip              QAccessible__Role = QAccessible__Role(0x0000000D)
	QAccessible__Application          QAccessible__Role = QAccessible__Role(0x0000000E)
	QAccessible__Document             QAccessible__Role = QAccessible__Role(0x0000000F)
	QAccessible__Pane                 QAccessible__Role = QAccessible__Role(0x00000010)
	QAccessible__Chart                QAccessible__Role = QAccessible__Role(0x00000011)
	QAccessible__Dialog               QAccessible__Role = QAccessible__Role(0x00000012)
	QAccessible__Border               QAccessible__Role = QAccessible__Role(0x00000013)
	QAccessible__Grouping             QAccessible__Role = QAccessible__Role(0x00000014)
	QAccessible__Separator            QAccessible__Role = QAccessible__Role(0x00000015)
	QAccessible__ToolBar              QAccessible__Role = QAccessible__Role(0x00000016)
	QAccessible__StatusBar            QAccessible__Role = QAccessible__Role(0x00000017)
	QAccessible__Table                QAccessible__Role = QAccessible__Role(0x00000018)
	QAccessible__ColumnHeader         QAccessible__Role = QAccessible__Role(0x00000019)
	QAccessible__RowHeader            QAccessible__Role = QAccessible__Role(0x0000001A)
	QAccessible__Column               QAccessible__Role = QAccessible__Role(0x0000001B)
	QAccessible__Row                  QAccessible__Role = QAccessible__Role(0x0000001C)
	QAccessible__Cell                 QAccessible__Role = QAccessible__Role(0x0000001D)
	QAccessible__Link                 QAccessible__Role = QAccessible__Role(0x0000001E)
	QAccessible__HelpBalloon          QAccessible__Role = QAccessible__Role(0x0000001F)
	QAccessible__Assistant            QAccessible__Role = QAccessible__Role(0x00000020)
	QAccessible__List                 QAccessible__Role = QAccessible__Role(0x00000021)
	QAccessible__ListItem             QAccessible__Role = QAccessible__Role(0x00000022)
	QAccessible__Tree                 QAccessible__Role = QAccessible__Role(0x00000023)
	QAccessible__TreeItem             QAccessible__Role = QAccessible__Role(0x00000024)
	QAccessible__PageTab              QAccessible__Role = QAccessible__Role(0x00000025)
	QAccessible__PropertyPage         QAccessible__Role = QAccessible__Role(0x00000026)
	QAccessible__Indicator            QAccessible__Role = QAccessible__Role(0x00000027)
	QAccessible__Graphic              QAccessible__Role = QAccessible__Role(0x00000028)
	QAccessible__StaticText           QAccessible__Role = QAccessible__Role(0x00000029)
	QAccessible__EditableText         QAccessible__Role = QAccessible__Role(0x0000002A)
	QAccessible__Button               QAccessible__Role = QAccessible__Role(0x0000002B)
	QAccessible__CheckBox             QAccessible__Role = QAccessible__Role(0x0000002C)
	QAccessible__RadioButton          QAccessible__Role = QAccessible__Role(0x0000002D)
	QAccessible__ComboBox             QAccessible__Role = QAccessible__Role(0x0000002E)
	QAccessible__ProgressBar          QAccessible__Role = QAccessible__Role(0x00000030)
	QAccessible__Dial                 QAccessible__Role = QAccessible__Role(0x00000031)
	QAccessible__HotkeyField          QAccessible__Role = QAccessible__Role(0x00000032)
	QAccessible__Slider               QAccessible__Role = QAccessible__Role(0x00000033)
	QAccessible__SpinBox              QAccessible__Role = QAccessible__Role(0x00000034)
	QAccessible__Canvas               QAccessible__Role = QAccessible__Role(0x00000035)
	QAccessible__Animation            QAccessible__Role = QAccessible__Role(0x00000036)
	QAccessible__Equation             QAccessible__Role = QAccessible__Role(0x00000037)
	QAccessible__ButtonDropDown       QAccessible__Role = QAccessible__Role(0x00000038)
	QAccessible__ButtonMenu           QAccessible__Role = QAccessible__Role(0x00000039)
	QAccessible__ButtonDropGrid       QAccessible__Role = QAccessible__Role(0x0000003A)
	QAccessible__Whitespace           QAccessible__Role = QAccessible__Role(0x0000003B)
	QAccessible__PageTabList          QAccessible__Role = QAccessible__Role(0x0000003C)
	QAccessible__Clock                QAccessible__Role = QAccessible__Role(0x0000003D)
	QAccessible__Splitter             QAccessible__Role = QAccessible__Role(0x0000003E)
	QAccessible__LayeredPane          QAccessible__Role = QAccessible__Role(0x00000080)
	QAccessible__Terminal             QAccessible__Role = QAccessible__Role(0x00000081)
	QAccessible__Desktop              QAccessible__Role = QAccessible__Role(0x00000082)
	QAccessible__Paragraph            QAccessible__Role = QAccessible__Role(0x00000083)
	QAccessible__WebDocument          QAccessible__Role = QAccessible__Role(0x00000084)
	QAccessible__Section              QAccessible__Role = QAccessible__Role(0x00000085)
	QAccessible__Notification         QAccessible__Role = QAccessible__Role(0x00000086)
	QAccessible__ColorChooser         QAccessible__Role = QAccessible__Role(0x404)
	QAccessible__Footer               QAccessible__Role = QAccessible__Role(0x40E)
	QAccessible__Form                 QAccessible__Role = QAccessible__Role(0x410)
	QAccessible__Heading              QAccessible__Role = QAccessible__Role(0x414)
	QAccessible__Note                 QAccessible__Role = QAccessible__Role(0x41B)
	QAccessible__ComplementaryContent QAccessible__Role = QAccessible__Role(0x42C)
	QAccessible__UserRole             QAccessible__Role = QAccessible__Role(0x0000ffff)
)

// QAccessible::Text
//
//go:generate stringer -type=QAccessible__Text
type QAccessible__Text int64

const (
	QAccessible__Name             QAccessible__Text = QAccessible__Text(0)
	QAccessible__Description      QAccessible__Text = QAccessible__Text(1)
	QAccessible__Value            QAccessible__Text = QAccessible__Text(2)
	QAccessible__Help             QAccessible__Text = QAccessible__Text(3)
	QAccessible__Accelerator      QAccessible__Text = QAccessible__Text(4)
	QAccessible__DebugDescription QAccessible__Text = QAccessible__Text(5)
	QAccessible__UserText         QAccessible__Text = QAccessible__Text(0x0000ffff)
)

// QAccessible::TextBoundaryType
//
//go:generate stringer -type=QAccessible__TextBoundaryType
type QAccessible__TextBoundaryType int64

const (
	QAccessible__CharBoundary      QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(0)
	QAccessible__WordBoundary      QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(1)
	QAccessible__SentenceBoundary  QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(2)
	QAccessible__ParagraphBoundary QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(3)
	QAccessible__LineBoundary      QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(4)
	QAccessible__NoBoundary        QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(5)
)

// QAccessibleTableModelChangeEvent::ModelChangeType
//
//go:generate stringer -type=QAccessibleTableModelChangeEvent__ModelChangeType
type QAccessibleTableModelChangeEvent__ModelChangeType int64

const (
	QAccessibleTableModelChangeEvent__ModelReset      QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(0)
	QAccessibleTableModelChangeEvent__DataChanged     QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(1)
	QAccessibleTableModelChangeEvent__RowsInserted    QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(2)
	QAccessibleTableModelChangeEvent__ColumnsInserted QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(3)
	QAccessibleTableModelChangeEvent__RowsRemoved     QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(4)
	QAccessibleTableModelChangeEvent__ColumnsRemoved  QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(5)
)

type QBitmap struct {
	QPixmap
}

type QBitmap_ITF interface {
	QPixmap_ITF
	QBitmap_PTR() *QBitmap
}

func (ptr *QBitmap) QBitmap_PTR() *QBitmap {
	return ptr
}

func (ptr *QBitmap) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPixmap_PTR().Pointer()
	}
	return nil
}

func (ptr *QBitmap) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPixmap_PTR().SetPointer(p)
	}
}

func PointerFromQBitmap(ptr QBitmap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBitmap_PTR().Pointer()
	}
	return nil
}

func NewQBitmapFromPointer(ptr unsafe.Pointer) (n *QBitmap) {
	n = new(QBitmap)
	n.SetPointer(ptr)
	return
}
func NewQBitmap() *QBitmap {
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap())
}

func NewQBitmap2(pixmap QPixmap_ITF) *QBitmap {
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap2(PointerFromQPixmap(pixmap)))
}

func NewQBitmap3(width int, height int) *QBitmap {
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap3(C.int(int32(width)), C.int(int32(height))))
}

func NewQBitmap4(size core.QSize_ITF) *QBitmap {
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap4(core.PointerFromQSize(size)))
}

func NewQBitmap5(fileName string, format string) *QBitmap {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap5(C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, formatC))
}

func (ptr *QBitmap) Clear() {
	if ptr.Pointer() != nil {
		C.QBitmap_Clear(ptr.Pointer())
	}
}

func QBitmap_FromData(size core.QSize_ITF, bits string, monoFormat QImage__Format) *QBitmap {
	var bitsC *C.char
	if bits != "" {
		bitsC = C.CString(bits)
		defer C.free(unsafe.Pointer(bitsC))
	}
	tmpValue := NewQBitmapFromPointer(C.QBitmap_QBitmap_FromData(core.PointerFromQSize(size), bitsC, C.longlong(monoFormat)))
	qt.SetFinalizer(tmpValue, (*QBitmap).DestroyQBitmap)
	return tmpValue
}

func (ptr *QBitmap) FromData(size core.QSize_ITF, bits string, monoFormat QImage__Format) *QBitmap {
	var bitsC *C.char
	if bits != "" {
		bitsC = C.CString(bits)
		defer C.free(unsafe.Pointer(bitsC))
	}
	tmpValue := NewQBitmapFromPointer(C.QBitmap_QBitmap_FromData(core.PointerFromQSize(size), bitsC, C.longlong(monoFormat)))
	qt.SetFinalizer(tmpValue, (*QBitmap).DestroyQBitmap)
	return tmpValue
}

//export callbackQBitmap_DestroyQBitmap
func callbackQBitmap_DestroyQBitmap(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBitmap"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBitmapFromPointer(ptr).DestroyQBitmapDefault()
	}
}

func (ptr *QBitmap) ConnectDestroyQBitmap(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBitmap"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBitmap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBitmap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBitmap) DisconnectDestroyQBitmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBitmap")
	}
}

func (ptr *QBitmap) DestroyQBitmap() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBitmap_DestroyQBitmap(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBitmap) DestroyQBitmapDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBitmap_DestroyQBitmapDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QBrush struct {
	ptr unsafe.Pointer
}

type QBrush_ITF interface {
	QBrush_PTR() *QBrush
}

func (ptr *QBrush) QBrush_PTR() *QBrush {
	return ptr
}

func (ptr *QBrush) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBrush) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBrush(ptr QBrush_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBrush_PTR().Pointer()
	}
	return nil
}

func NewQBrushFromPointer(ptr unsafe.Pointer) (n *QBrush) {
	n = new(QBrush)
	n.SetPointer(ptr)
	return
}
func NewQBrush() *QBrush {
	tmpValue := NewQBrushFromPointer(C.QBrush_NewQBrush())
	qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
	return tmpValue
}

func NewQBrush2(style core.Qt__BrushStyle) *QBrush {
	tmpValue := NewQBrushFromPointer(C.QBrush_NewQBrush2(C.longlong(style)))
	qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
	return tmpValue
}

func NewQBrush3(color QColor_ITF, style core.Qt__BrushStyle) *QBrush {
	tmpValue := NewQBrushFromPointer(C.QBrush_NewQBrush3(PointerFromQColor(color), C.longlong(style)))
	qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
	return tmpValue
}

func NewQBrush4(color core.Qt__GlobalColor, style core.Qt__BrushStyle) *QBrush {
	tmpValue := NewQBrushFromPointer(C.QBrush_NewQBrush4(C.longlong(color), C.longlong(style)))
	qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
	return tmpValue
}

func NewQBrush5(color QColor_ITF, pixmap QPixmap_ITF) *QBrush {
	tmpValue := NewQBrushFromPointer(C.QBrush_NewQBrush5(PointerFromQColor(color), PointerFromQPixmap(pixmap)))
	qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
	return tmpValue
}

func NewQBrush6(color core.Qt__GlobalColor, pixmap QPixmap_ITF) *QBrush {
	tmpValue := NewQBrushFromPointer(C.QBrush_NewQBrush6(C.longlong(color), PointerFromQPixmap(pixmap)))
	qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
	return tmpValue
}

func NewQBrush7(pixmap QPixmap_ITF) *QBrush {
	tmpValue := NewQBrushFromPointer(C.QBrush_NewQBrush7(PointerFromQPixmap(pixmap)))
	qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
	return tmpValue
}

func NewQBrush8(image QImage_ITF) *QBrush {
	tmpValue := NewQBrushFromPointer(C.QBrush_NewQBrush8(PointerFromQImage(image)))
	qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
	return tmpValue
}

func NewQBrush9(other QBrush_ITF) *QBrush {
	tmpValue := NewQBrushFromPointer(C.QBrush_NewQBrush9(PointerFromQBrush(other)))
	qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
	return tmpValue
}

func NewQBrush10(gradient QGradient_ITF) *QBrush {
	tmpValue := NewQBrushFromPointer(C.QBrush_NewQBrush10(PointerFromQGradient(gradient)))
	qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
	return tmpValue
}

func (ptr *QBrush) Color() *QColor {
	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QBrush_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBrush) SetStyle(style core.Qt__BrushStyle) {
	if ptr.Pointer() != nil {
		C.QBrush_SetStyle(ptr.Pointer(), C.longlong(style))
	}
}

func (ptr *QBrush) Style() core.Qt__BrushStyle {
	if ptr.Pointer() != nil {
		return core.Qt__BrushStyle(C.QBrush_Style(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBrush) Transform() *QTransform {
	if ptr.Pointer() != nil {
		tmpValue := NewQTransformFromPointer(C.QBrush_Transform(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
		return tmpValue
	}
	return nil
}

func (ptr *QBrush) DestroyQBrush() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBrush_DestroyQBrush(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBrush) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QBrush_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

type QClipboard struct {
	core.QObject
}

type QClipboard_ITF interface {
	core.QObject_ITF
	QClipboard_PTR() *QClipboard
}

func (ptr *QClipboard) QClipboard_PTR() *QClipboard {
	return ptr
}

func (ptr *QClipboard) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QClipboard) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQClipboard(ptr QClipboard_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QClipboard_PTR().Pointer()
	}
	return nil
}

func NewQClipboardFromPointer(ptr unsafe.Pointer) (n *QClipboard) {
	n = new(QClipboard)
	n.SetPointer(ptr)
	return
}

// QClipboard::Mode
//
//go:generate stringer -type=QClipboard__Mode
type QClipboard__Mode int64

const (
	QClipboard__Clipboard  QClipboard__Mode = QClipboard__Mode(0)
	QClipboard__Selection  QClipboard__Mode = QClipboard__Mode(1)
	QClipboard__FindBuffer QClipboard__Mode = QClipboard__Mode(2)
	QClipboard__LastMode   QClipboard__Mode = QClipboard__Mode(QClipboard__FindBuffer)
)

//export callbackQClipboard_Changed
func callbackQClipboard_Changed(ptr unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(ptr, "changed"); signal != nil {
		(*(*func(QClipboard__Mode))(signal))(QClipboard__Mode(mode))
	}

}

func (ptr *QClipboard) ConnectChanged(f func(mode QClipboard__Mode)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "changed") {
			C.QClipboard_ConnectChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "changed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "changed"); signal != nil {
			f := func(mode QClipboard__Mode) {
				(*(*func(QClipboard__Mode))(signal))(mode)
				f(mode)
			}
			qt.ConnectSignal(ptr.Pointer(), "changed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "changed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QClipboard) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "changed")
	}
}

func (ptr *QClipboard) Changed(mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_Changed(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QClipboard) Clear(mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_Clear(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QClipboard) Image(mode QClipboard__Mode) *QImage {
	if ptr.Pointer() != nil {
		tmpValue := NewQImageFromPointer(C.QClipboard_Image(ptr.Pointer(), C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QClipboard) Pixmap(mode QClipboard__Mode) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QClipboard_Pixmap(ptr.Pointer(), C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

//export callbackQClipboard_SelectionChanged
func callbackQClipboard_SelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QClipboard) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionChanged") {
			C.QClipboard_ConnectSelectionChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "selectionChanged")))
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

func (ptr *QClipboard) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionChanged")
	}
}

func (ptr *QClipboard) SelectionChanged() {
	if ptr.Pointer() != nil {
		C.QClipboard_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QClipboard) SetPixmap(pixmap QPixmap_ITF, mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		C.QClipboard_SetPixmap(ptr.Pointer(), PointerFromQPixmap(pixmap), C.longlong(mode))
	}
}

func (ptr *QClipboard) SetText(text string, mode QClipboard__Mode) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QClipboard_SetText(ptr.Pointer(), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}, C.longlong(mode))
	}
}

func (ptr *QClipboard) Text(mode QClipboard__Mode) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QClipboard_Text(ptr.Pointer(), C.longlong(mode)))
	}
	return ""
}

func (ptr *QClipboard) Text2(subtype string, mode QClipboard__Mode) string {
	if ptr.Pointer() != nil {
		var subtypeC *C.char
		if subtype != "" {
			subtypeC = C.CString(subtype)
			defer C.free(unsafe.Pointer(subtypeC))
		}
		return cGoUnpackString(C.QClipboard_Text2(ptr.Pointer(), C.struct_QtGui_PackedString{data: subtypeC, len: C.longlong(len(subtype))}, C.longlong(mode)))
	}
	return ""
}

func (ptr *QClipboard) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QClipboard___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QClipboard) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QClipboard___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QClipboard) __children_newList() unsafe.Pointer {
	return C.QClipboard___children_newList(ptr.Pointer())
}

func (ptr *QClipboard) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QClipboard___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QClipboard) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QClipboard___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QClipboard) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QClipboard___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QClipboard) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QClipboard___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QClipboard) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QClipboard___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QClipboard) __findChildren_newList() unsafe.Pointer {
	return C.QClipboard___findChildren_newList(ptr.Pointer())
}

func (ptr *QClipboard) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QClipboard___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QClipboard) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QClipboard___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QClipboard) __findChildren_newList3() unsafe.Pointer {
	return C.QClipboard___findChildren_newList3(ptr.Pointer())
}

//export callbackQClipboard_ChildEvent
func callbackQClipboard_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQClipboardFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QClipboard) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QClipboard_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQClipboard_ConnectNotify
func callbackQClipboard_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQClipboardFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QClipboard) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QClipboard_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQClipboard_CustomEvent
func callbackQClipboard_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQClipboardFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QClipboard) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QClipboard_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQClipboard_DeleteLater
func callbackQClipboard_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQClipboardFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QClipboard) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QClipboard_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQClipboard_Destroyed
func callbackQClipboard_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQClipboard_DisconnectNotify
func callbackQClipboard_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQClipboardFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QClipboard) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QClipboard_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQClipboard_Event
func callbackQClipboard_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQClipboardFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QClipboard) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QClipboard_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQClipboard_EventFilter
func callbackQClipboard_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQClipboardFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QClipboard) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QClipboard_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQClipboard_ObjectNameChanged
func callbackQClipboard_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGui_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQClipboard_TimerEvent
func callbackQClipboard_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQClipboardFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QClipboard) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QClipboard_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QCloseEvent struct {
	core.QEvent
}

type QCloseEvent_ITF interface {
	core.QEvent_ITF
	QCloseEvent_PTR() *QCloseEvent
}

func (ptr *QCloseEvent) QCloseEvent_PTR() *QCloseEvent {
	return ptr
}

func (ptr *QCloseEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QCloseEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQCloseEvent(ptr QCloseEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCloseEvent_PTR().Pointer()
	}
	return nil
}

func NewQCloseEventFromPointer(ptr unsafe.Pointer) (n *QCloseEvent) {
	n = new(QCloseEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QCloseEvent) DestroyQCloseEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQCloseEvent() *QCloseEvent {
	tmpValue := NewQCloseEventFromPointer(C.QCloseEvent_NewQCloseEvent())
	qt.SetFinalizer(tmpValue, (*QCloseEvent).DestroyQCloseEvent)
	return tmpValue
}

type QColor struct {
	ptr unsafe.Pointer
}

type QColor_ITF interface {
	QColor_PTR() *QColor
}

func (ptr *QColor) QColor_PTR() *QColor {
	return ptr
}

func (ptr *QColor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QColor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQColor(ptr QColor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColor_PTR().Pointer()
	}
	return nil
}

func NewQColorFromPointer(ptr unsafe.Pointer) (n *QColor) {
	n = new(QColor)
	n.SetPointer(ptr)
	return
}
func (ptr *QColor) DestroyQColor() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QColor__NameFormat
//QColor::NameFormat
type QColor__NameFormat int64

const (
	QColor__HexRgb  QColor__NameFormat = QColor__NameFormat(0)
	QColor__HexArgb QColor__NameFormat = QColor__NameFormat(1)
)

// QColor::Spec
//
//go:generate stringer -type=QColor__Spec
type QColor__Spec int64

const (
	QColor__Invalid     QColor__Spec = QColor__Spec(0)
	QColor__Rgb         QColor__Spec = QColor__Spec(1)
	QColor__Hsv         QColor__Spec = QColor__Spec(2)
	QColor__Cmyk        QColor__Spec = QColor__Spec(3)
	QColor__Hsl         QColor__Spec = QColor__Spec(4)
	QColor__ExtendedRgb QColor__Spec = QColor__Spec(5)
)

func NewQColor() *QColor {
	tmpValue := NewQColorFromPointer(C.QColor_NewQColor())
	qt.SetFinalizer(tmpValue, (*QColor).DestroyQColor)
	return tmpValue
}

func NewQColor2(color core.Qt__GlobalColor) *QColor {
	tmpValue := NewQColorFromPointer(C.QColor_NewQColor2(C.longlong(color)))
	qt.SetFinalizer(tmpValue, (*QColor).DestroyQColor)
	return tmpValue
}

func NewQColor3(r int, g int, b int, a int) *QColor {
	tmpValue := NewQColorFromPointer(C.QColor_NewQColor3(C.int(int32(r)), C.int(int32(g)), C.int(int32(b)), C.int(int32(a))))
	qt.SetFinalizer(tmpValue, (*QColor).DestroyQColor)
	return tmpValue
}

func NewQColor4(color uint) *QColor {
	tmpValue := NewQColorFromPointer(C.QColor_NewQColor4(C.uint(uint32(color))))
	qt.SetFinalizer(tmpValue, (*QColor).DestroyQColor)
	return tmpValue
}

func NewQColor5(rgba64 QRgba64_ITF) *QColor {
	tmpValue := NewQColorFromPointer(C.QColor_NewQColor5(PointerFromQRgba64(rgba64)))
	qt.SetFinalizer(tmpValue, (*QColor).DestroyQColor)
	return tmpValue
}

func NewQColor6(name string) *QColor {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQColorFromPointer(C.QColor_NewQColor6(C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QColor).DestroyQColor)
	return tmpValue
}

func NewQColor8(name string) *QColor {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQColorFromPointer(C.QColor_NewQColor8(nameC))
	qt.SetFinalizer(tmpValue, (*QColor).DestroyQColor)
	return tmpValue
}

func NewQColor9(name core.QLatin1String_ITF) *QColor {
	tmpValue := NewQColorFromPointer(C.QColor_NewQColor9(core.PointerFromQLatin1String(name)))
	qt.SetFinalizer(tmpValue, (*QColor).DestroyQColor)
	return tmpValue
}

func (ptr *QColor) Alpha() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QColor_Alpha(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QColor_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QColor) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QColor_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QColor) Name2(format QColor__NameFormat) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QColor_Name2(ptr.Pointer(), C.longlong(format)))
	}
	return ""
}

func (ptr *QColor) Value() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QColor_Value(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColor) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QColor_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

// QColorSpace::NamedColorSpace
//
//go:generate stringer -type=QColorSpace__NamedColorSpace
type QColorSpace__NamedColorSpace int64

const (
	QColorSpace__NamedColorSpace_SRgb        QColorSpace__NamedColorSpace = QColorSpace__NamedColorSpace(1)
	QColorSpace__NamedColorSpace_SRgbLinear  QColorSpace__NamedColorSpace = QColorSpace__NamedColorSpace(2)
	QColorSpace__NamedColorSpace_AdobeRgb    QColorSpace__NamedColorSpace = QColorSpace__NamedColorSpace(3)
	QColorSpace__NamedColorSpace_DisplayP3   QColorSpace__NamedColorSpace = QColorSpace__NamedColorSpace(4)
	QColorSpace__NamedColorSpace_ProPhotoRgb QColorSpace__NamedColorSpace = QColorSpace__NamedColorSpace(5)
)

// QColorSpace::Primaries
//
//go:generate stringer -type=QColorSpace__Primaries
type QColorSpace__Primaries int64

const (
	QColorSpace__Primaries_Custom      QColorSpace__Primaries = QColorSpace__Primaries(0)
	QColorSpace__Primaries_SRgb        QColorSpace__Primaries = QColorSpace__Primaries(1)
	QColorSpace__Primaries_AdobeRgb    QColorSpace__Primaries = QColorSpace__Primaries(2)
	QColorSpace__Primaries_DciP3D65    QColorSpace__Primaries = QColorSpace__Primaries(3)
	QColorSpace__Primaries_ProPhotoRgb QColorSpace__Primaries = QColorSpace__Primaries(4)
)

// QColorSpace::TransferFunction
//
//go:generate stringer -type=QColorSpace__TransferFunction
type QColorSpace__TransferFunction int64

const (
	QColorSpace__TransferFunction_Custom      QColorSpace__TransferFunction = QColorSpace__TransferFunction(0)
	QColorSpace__TransferFunction_Linear      QColorSpace__TransferFunction = QColorSpace__TransferFunction(1)
	QColorSpace__TransferFunction_Gamma       QColorSpace__TransferFunction = QColorSpace__TransferFunction(2)
	QColorSpace__TransferFunction_SRgb        QColorSpace__TransferFunction = QColorSpace__TransferFunction(3)
	QColorSpace__TransferFunction_ProPhotoRgb QColorSpace__TransferFunction = QColorSpace__TransferFunction(4)
)

type QContextMenuEvent struct {
	QInputEvent
}

type QContextMenuEvent_ITF interface {
	QInputEvent_ITF
	QContextMenuEvent_PTR() *QContextMenuEvent
}

func (ptr *QContextMenuEvent) QContextMenuEvent_PTR() *QContextMenuEvent {
	return ptr
}

func (ptr *QContextMenuEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QContextMenuEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QInputEvent_PTR().SetPointer(p)
	}
}

func PointerFromQContextMenuEvent(ptr QContextMenuEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QContextMenuEvent_PTR().Pointer()
	}
	return nil
}

func NewQContextMenuEventFromPointer(ptr unsafe.Pointer) (n *QContextMenuEvent) {
	n = new(QContextMenuEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QContextMenuEvent) DestroyQContextMenuEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QContextMenuEvent__Reason
//QContextMenuEvent::Reason
type QContextMenuEvent__Reason int64

const (
	QContextMenuEvent__Mouse    QContextMenuEvent__Reason = QContextMenuEvent__Reason(0)
	QContextMenuEvent__Keyboard QContextMenuEvent__Reason = QContextMenuEvent__Reason(1)
	QContextMenuEvent__Other    QContextMenuEvent__Reason = QContextMenuEvent__Reason(2)
)

func NewQContextMenuEvent(reason QContextMenuEvent__Reason, pos core.QPoint_ITF, globalPos core.QPoint_ITF, modifiers core.Qt__KeyboardModifier) *QContextMenuEvent {
	tmpValue := NewQContextMenuEventFromPointer(C.QContextMenuEvent_NewQContextMenuEvent(C.longlong(reason), core.PointerFromQPoint(pos), core.PointerFromQPoint(globalPos), C.longlong(modifiers)))
	qt.SetFinalizer(tmpValue, (*QContextMenuEvent).DestroyQContextMenuEvent)
	return tmpValue
}

func NewQContextMenuEvent2(reason QContextMenuEvent__Reason, pos core.QPoint_ITF, globalPos core.QPoint_ITF) *QContextMenuEvent {
	tmpValue := NewQContextMenuEventFromPointer(C.QContextMenuEvent_NewQContextMenuEvent2(C.longlong(reason), core.PointerFromQPoint(pos), core.PointerFromQPoint(globalPos)))
	qt.SetFinalizer(tmpValue, (*QContextMenuEvent).DestroyQContextMenuEvent)
	return tmpValue
}

func NewQContextMenuEvent3(reason QContextMenuEvent__Reason, pos core.QPoint_ITF) *QContextMenuEvent {
	tmpValue := NewQContextMenuEventFromPointer(C.QContextMenuEvent_NewQContextMenuEvent3(C.longlong(reason), core.PointerFromQPoint(pos)))
	qt.SetFinalizer(tmpValue, (*QContextMenuEvent).DestroyQContextMenuEvent)
	return tmpValue
}

func (ptr *QContextMenuEvent) GlobalPos() *core.QPoint {
	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QContextMenuEvent_GlobalPos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QContextMenuEvent) Pos() *core.QPoint {
	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QContextMenuEvent_Pos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QContextMenuEvent) X() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QContextMenuEvent_X(ptr.Pointer())))
	}
	return 0
}

func (ptr *QContextMenuEvent) Y() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QContextMenuEvent_Y(ptr.Pointer())))
	}
	return 0
}

type QCursor struct {
	ptr unsafe.Pointer
}

type QCursor_ITF interface {
	QCursor_PTR() *QCursor
}

func (ptr *QCursor) QCursor_PTR() *QCursor {
	return ptr
}

func (ptr *QCursor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCursor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCursor(ptr QCursor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCursor_PTR().Pointer()
	}
	return nil
}

func NewQCursorFromPointer(ptr unsafe.Pointer) (n *QCursor) {
	n = new(QCursor)
	n.SetPointer(ptr)
	return
}
func NewQCursor() *QCursor {
	tmpValue := NewQCursorFromPointer(C.QCursor_NewQCursor())
	qt.SetFinalizer(tmpValue, (*QCursor).DestroyQCursor)
	return tmpValue
}

func NewQCursor2(shape core.Qt__CursorShape) *QCursor {
	tmpValue := NewQCursorFromPointer(C.QCursor_NewQCursor2(C.longlong(shape)))
	qt.SetFinalizer(tmpValue, (*QCursor).DestroyQCursor)
	return tmpValue
}

func NewQCursor3(bitmap QBitmap_ITF, mask QBitmap_ITF, hotX int, hotY int) *QCursor {
	tmpValue := NewQCursorFromPointer(C.QCursor_NewQCursor3(PointerFromQBitmap(bitmap), PointerFromQBitmap(mask), C.int(int32(hotX)), C.int(int32(hotY))))
	qt.SetFinalizer(tmpValue, (*QCursor).DestroyQCursor)
	return tmpValue
}

func NewQCursor4(pixmap QPixmap_ITF, hotX int, hotY int) *QCursor {
	tmpValue := NewQCursorFromPointer(C.QCursor_NewQCursor4(PointerFromQPixmap(pixmap), C.int(int32(hotX)), C.int(int32(hotY))))
	qt.SetFinalizer(tmpValue, (*QCursor).DestroyQCursor)
	return tmpValue
}

func NewQCursor5(c QCursor_ITF) *QCursor {
	tmpValue := NewQCursorFromPointer(C.QCursor_NewQCursor5(PointerFromQCursor(c)))
	qt.SetFinalizer(tmpValue, (*QCursor).DestroyQCursor)
	return tmpValue
}

func NewQCursor6(other QCursor_ITF) *QCursor {
	tmpValue := NewQCursorFromPointer(C.QCursor_NewQCursor6(PointerFromQCursor(other)))
	qt.SetFinalizer(tmpValue, (*QCursor).DestroyQCursor)
	return tmpValue
}

func (ptr *QCursor) Pixmap() *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QCursor_Pixmap(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func QCursor_Pos() *core.QPoint {
	tmpValue := core.NewQPointFromPointer(C.QCursor_QCursor_Pos())
	qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
	return tmpValue
}

func (ptr *QCursor) Pos() *core.QPoint {
	tmpValue := core.NewQPointFromPointer(C.QCursor_QCursor_Pos())
	qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
	return tmpValue
}

func QCursor_Pos2(screen QScreen_ITF) *core.QPoint {
	tmpValue := core.NewQPointFromPointer(C.QCursor_QCursor_Pos2(PointerFromQScreen(screen)))
	qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
	return tmpValue
}

func (ptr *QCursor) Pos2(screen QScreen_ITF) *core.QPoint {
	tmpValue := core.NewQPointFromPointer(C.QCursor_QCursor_Pos2(PointerFromQScreen(screen)))
	qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
	return tmpValue
}

func QCursor_SetPos(x int, y int) {
	C.QCursor_QCursor_SetPos(C.int(int32(x)), C.int(int32(y)))
}

func (ptr *QCursor) SetPos(x int, y int) {
	C.QCursor_QCursor_SetPos(C.int(int32(x)), C.int(int32(y)))
}

func QCursor_SetPos2(screen QScreen_ITF, x int, y int) {
	C.QCursor_QCursor_SetPos2(PointerFromQScreen(screen), C.int(int32(x)), C.int(int32(y)))
}

func (ptr *QCursor) SetPos2(screen QScreen_ITF, x int, y int) {
	C.QCursor_QCursor_SetPos2(PointerFromQScreen(screen), C.int(int32(x)), C.int(int32(y)))
}

func QCursor_SetPos3(p core.QPoint_ITF) {
	C.QCursor_QCursor_SetPos3(core.PointerFromQPoint(p))
}

func (ptr *QCursor) SetPos3(p core.QPoint_ITF) {
	C.QCursor_QCursor_SetPos3(core.PointerFromQPoint(p))
}

func QCursor_SetPos4(screen QScreen_ITF, p core.QPoint_ITF) {
	C.QCursor_QCursor_SetPos4(PointerFromQScreen(screen), core.PointerFromQPoint(p))
}

func (ptr *QCursor) SetPos4(screen QScreen_ITF, p core.QPoint_ITF) {
	C.QCursor_QCursor_SetPos4(PointerFromQScreen(screen), core.PointerFromQPoint(p))
}

func (ptr *QCursor) DestroyQCursor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCursor_DestroyQCursor(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QDoubleValidator::Notation
//
//go:generate stringer -type=QDoubleValidator__Notation
type QDoubleValidator__Notation int64

const (
	QDoubleValidator__StandardNotation   QDoubleValidator__Notation = QDoubleValidator__Notation(0)
	QDoubleValidator__ScientificNotation QDoubleValidator__Notation = QDoubleValidator__Notation(1)
)

type QFont struct {
	ptr unsafe.Pointer
}

type QFont_ITF interface {
	QFont_PTR() *QFont
}

func (ptr *QFont) QFont_PTR() *QFont {
	return ptr
}

func (ptr *QFont) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QFont) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQFont(ptr QFont_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFont_PTR().Pointer()
	}
	return nil
}

func NewQFontFromPointer(ptr unsafe.Pointer) (n *QFont) {
	n = new(QFont)
	n.SetPointer(ptr)
	return
}

// QFont::Capitalization
//
//go:generate stringer -type=QFont__Capitalization
type QFont__Capitalization int64

const (
	QFont__MixedCase    QFont__Capitalization = QFont__Capitalization(0)
	QFont__AllUppercase QFont__Capitalization = QFont__Capitalization(1)
	QFont__AllLowercase QFont__Capitalization = QFont__Capitalization(2)
	QFont__SmallCaps    QFont__Capitalization = QFont__Capitalization(3)
	QFont__Capitalize   QFont__Capitalization = QFont__Capitalization(4)
)

// QFont::HintingPreference
//
//go:generate stringer -type=QFont__HintingPreference
type QFont__HintingPreference int64

const (
	QFont__PreferDefaultHinting  QFont__HintingPreference = QFont__HintingPreference(0)
	QFont__PreferNoHinting       QFont__HintingPreference = QFont__HintingPreference(1)
	QFont__PreferVerticalHinting QFont__HintingPreference = QFont__HintingPreference(2)
	QFont__PreferFullHinting     QFont__HintingPreference = QFont__HintingPreference(3)
)

// QFont::SpacingType
//
//go:generate stringer -type=QFont__SpacingType
type QFont__SpacingType int64

const (
	QFont__PercentageSpacing QFont__SpacingType = QFont__SpacingType(0)
	QFont__AbsoluteSpacing   QFont__SpacingType = QFont__SpacingType(1)
)

// QFont::Stretch
//
//go:generate stringer -type=QFont__Stretch
type QFont__Stretch int64

const (
	QFont__AnyStretch     QFont__Stretch = QFont__Stretch(0)
	QFont__UltraCondensed QFont__Stretch = QFont__Stretch(50)
	QFont__ExtraCondensed QFont__Stretch = QFont__Stretch(62)
	QFont__Condensed      QFont__Stretch = QFont__Stretch(75)
	QFont__SemiCondensed  QFont__Stretch = QFont__Stretch(87)
	QFont__Unstretched    QFont__Stretch = QFont__Stretch(100)
	QFont__SemiExpanded   QFont__Stretch = QFont__Stretch(112)
	QFont__Expanded       QFont__Stretch = QFont__Stretch(125)
	QFont__ExtraExpanded  QFont__Stretch = QFont__Stretch(150)
	QFont__UltraExpanded  QFont__Stretch = QFont__Stretch(200)
)

// QFont::Style
//
//go:generate stringer -type=QFont__Style
type QFont__Style int64

var (
	QFont__StyleNormal  QFont__Style = QFont__Style(0)
	QFont__StyleItalic  QFont__Style = QFont__Style(1)
	QFont__StyleOblique QFont__Style = QFont__Style(2)
)

// QFont::StyleHint
//
//go:generate stringer -type=QFont__StyleHint
type QFont__StyleHint int64

var (
	QFont__Helvetica  QFont__StyleHint = QFont__StyleHint(0)
	QFont__SansSerif  QFont__StyleHint = QFont__StyleHint(QFont__Helvetica)
	QFont__Times      QFont__StyleHint = QFont__StyleHint(1)
	QFont__Serif      QFont__StyleHint = QFont__StyleHint(QFont__Times)
	QFont__Courier    QFont__StyleHint = QFont__StyleHint(2)
	QFont__TypeWriter QFont__StyleHint = QFont__StyleHint(QFont__Courier)
	QFont__OldEnglish QFont__StyleHint = QFont__StyleHint(3)
	QFont__Decorative QFont__StyleHint = QFont__StyleHint(QFont__OldEnglish)
	QFont__System     QFont__StyleHint = QFont__StyleHint(4)
	QFont__AnyStyle   QFont__StyleHint = QFont__StyleHint(5)
	QFont__Cursive    QFont__StyleHint = QFont__StyleHint(6)
	QFont__Monospace  QFont__StyleHint = QFont__StyleHint(7)
	QFont__Fantasy    QFont__StyleHint = QFont__StyleHint(8)
)

// QFont::StyleStrategy
//
//go:generate stringer -type=QFont__StyleStrategy
type QFont__StyleStrategy int64

var (
	QFont__PreferDefault       QFont__StyleStrategy = QFont__StyleStrategy(0x0001)
	QFont__PreferBitmap        QFont__StyleStrategy = QFont__StyleStrategy(0x0002)
	QFont__PreferDevice        QFont__StyleStrategy = QFont__StyleStrategy(0x0004)
	QFont__PreferOutline       QFont__StyleStrategy = QFont__StyleStrategy(0x0008)
	QFont__ForceOutline        QFont__StyleStrategy = QFont__StyleStrategy(0x0010)
	QFont__PreferMatch         QFont__StyleStrategy = QFont__StyleStrategy(0x0020)
	QFont__PreferQuality       QFont__StyleStrategy = QFont__StyleStrategy(0x0040)
	QFont__PreferAntialias     QFont__StyleStrategy = QFont__StyleStrategy(0x0080)
	QFont__NoAntialias         QFont__StyleStrategy = QFont__StyleStrategy(0x0100)
	QFont__OpenGLCompatible    QFont__StyleStrategy = QFont__StyleStrategy(0x0200)
	QFont__ForceIntegerMetrics QFont__StyleStrategy = QFont__StyleStrategy(0x0400)
	QFont__NoSubpixelAntialias QFont__StyleStrategy = QFont__StyleStrategy(0x0800)
	QFont__PreferNoShaping     QFont__StyleStrategy = QFont__StyleStrategy(0x1000)
	QFont__NoFontMerging       QFont__StyleStrategy = QFont__StyleStrategy(0x8000)
)

// QFont::Weight
//
//go:generate stringer -type=QFont__Weight
type QFont__Weight int64

const (
	QFont__Thin       QFont__Weight = QFont__Weight(0)
	QFont__ExtraLight QFont__Weight = QFont__Weight(12)
	QFont__Light      QFont__Weight = QFont__Weight(25)
	QFont__Normal     QFont__Weight = QFont__Weight(50)
	QFont__Medium     QFont__Weight = QFont__Weight(57)
	QFont__DemiBold   QFont__Weight = QFont__Weight(63)
	QFont__Bold       QFont__Weight = QFont__Weight(75)
	QFont__ExtraBold  QFont__Weight = QFont__Weight(81)
	QFont__Black      QFont__Weight = QFont__Weight(87)
)

func NewQFont() *QFont {
	tmpValue := NewQFontFromPointer(C.QFont_NewQFont())
	qt.SetFinalizer(tmpValue, (*QFont).DestroyQFont)
	return tmpValue
}

func NewQFont2(family string, pointSize int, weight int, italic bool) *QFont {
	var familyC *C.char
	if family != "" {
		familyC = C.CString(family)
		defer C.free(unsafe.Pointer(familyC))
	}
	tmpValue := NewQFontFromPointer(C.QFont_NewQFont2(C.struct_QtGui_PackedString{data: familyC, len: C.longlong(len(family))}, C.int(int32(pointSize)), C.int(int32(weight)), C.char(int8(qt.GoBoolToInt(italic)))))
	qt.SetFinalizer(tmpValue, (*QFont).DestroyQFont)
	return tmpValue
}

func NewQFont4(font QFont_ITF, pd QPaintDevice_ITF) *QFont {
	tmpValue := NewQFontFromPointer(C.QFont_NewQFont4(PointerFromQFont(font), PointerFromQPaintDevice(pd)))
	qt.SetFinalizer(tmpValue, (*QFont).DestroyQFont)
	return tmpValue
}

func NewQFont5(font QFont_ITF) *QFont {
	tmpValue := NewQFontFromPointer(C.QFont_NewQFont5(PointerFromQFont(font)))
	qt.SetFinalizer(tmpValue, (*QFont).DestroyQFont)
	return tmpValue
}

func (ptr *QFont) Bold() bool {
	if ptr.Pointer() != nil {
		return int8(C.QFont_Bold(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) Family() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QFont_Family(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFont) FromString(descrip string) bool {
	if ptr.Pointer() != nil {
		var descripC *C.char
		if descrip != "" {
			descripC = C.CString(descrip)
			defer C.free(unsafe.Pointer(descripC))
		}
		return int8(C.QFont_FromString(ptr.Pointer(), C.struct_QtGui_PackedString{data: descripC, len: C.longlong(len(descrip))})) != 0
	}
	return false
}

func (ptr *QFont) Italic() bool {
	if ptr.Pointer() != nil {
		return int8(C.QFont_Italic(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) Key() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QFont_Key(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFont) PointSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QFont_PointSize(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFont) SetFamily(family string) {
	if ptr.Pointer() != nil {
		var familyC *C.char
		if family != "" {
			familyC = C.CString(family)
			defer C.free(unsafe.Pointer(familyC))
		}
		C.QFont_SetFamily(ptr.Pointer(), C.struct_QtGui_PackedString{data: familyC, len: C.longlong(len(family))})
	}
}

func (ptr *QFont) SetPointSize(pointSize int) {
	if ptr.Pointer() != nil {
		C.QFont_SetPointSize(ptr.Pointer(), C.int(int32(pointSize)))
	}
}

func (ptr *QFont) SetStretch(factor int) {
	if ptr.Pointer() != nil {
		C.QFont_SetStretch(ptr.Pointer(), C.int(int32(factor)))
	}
}

func (ptr *QFont) SetStyle(style QFont__Style) {
	if ptr.Pointer() != nil {
		C.QFont_SetStyle(ptr.Pointer(), C.longlong(style))
	}
}

func (ptr *QFont) SetStyleName(styleName string) {
	if ptr.Pointer() != nil {
		var styleNameC *C.char
		if styleName != "" {
			styleNameC = C.CString(styleName)
			defer C.free(unsafe.Pointer(styleNameC))
		}
		C.QFont_SetStyleName(ptr.Pointer(), C.struct_QtGui_PackedString{data: styleNameC, len: C.longlong(len(styleName))})
	}
}

func (ptr *QFont) Stretch() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QFont_Stretch(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFont) StrikeOut() bool {
	if ptr.Pointer() != nil {
		return int8(C.QFont_StrikeOut(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) Style() QFont__Style {
	if ptr.Pointer() != nil {
		return QFont__Style(C.QFont_Style(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) StyleName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QFont_StyleName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFont) Underline() bool {
	if ptr.Pointer() != nil {
		return int8(C.QFont_Underline(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) Weight() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QFont_Weight(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFont) DestroyQFont() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFont_DestroyQFont(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFont) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QFont_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

// QFontDatabase::SystemFont
//
//go:generate stringer -type=QFontDatabase__SystemFont
type QFontDatabase__SystemFont int64

const (
	QFontDatabase__GeneralFont          QFontDatabase__SystemFont = QFontDatabase__SystemFont(0)
	QFontDatabase__FixedFont            QFontDatabase__SystemFont = QFontDatabase__SystemFont(1)
	QFontDatabase__TitleFont            QFontDatabase__SystemFont = QFontDatabase__SystemFont(2)
	QFontDatabase__SmallestReadableFont QFontDatabase__SystemFont = QFontDatabase__SystemFont(3)
)

// QFontDatabase::WritingSystem
//
//go:generate stringer -type=QFontDatabase__WritingSystem
type QFontDatabase__WritingSystem int64

const (
	QFontDatabase__Any                 QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(0)
	QFontDatabase__Latin               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(1)
	QFontDatabase__Greek               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(2)
	QFontDatabase__Cyrillic            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(3)
	QFontDatabase__Armenian            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(4)
	QFontDatabase__Hebrew              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(5)
	QFontDatabase__Arabic              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(6)
	QFontDatabase__Syriac              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(7)
	QFontDatabase__Thaana              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(8)
	QFontDatabase__Devanagari          QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(9)
	QFontDatabase__Bengali             QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(10)
	QFontDatabase__Gurmukhi            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(11)
	QFontDatabase__Gujarati            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(12)
	QFontDatabase__Oriya               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(13)
	QFontDatabase__Tamil               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(14)
	QFontDatabase__Telugu              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(15)
	QFontDatabase__Kannada             QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(16)
	QFontDatabase__Malayalam           QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(17)
	QFontDatabase__Sinhala             QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(18)
	QFontDatabase__Thai                QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(19)
	QFontDatabase__Lao                 QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(20)
	QFontDatabase__Tibetan             QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(21)
	QFontDatabase__Myanmar             QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(22)
	QFontDatabase__Georgian            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(23)
	QFontDatabase__Khmer               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(24)
	QFontDatabase__SimplifiedChinese   QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(25)
	QFontDatabase__TraditionalChinese  QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(26)
	QFontDatabase__Japanese            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(27)
	QFontDatabase__Korean              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(28)
	QFontDatabase__Vietnamese          QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(29)
	QFontDatabase__Symbol              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(30)
	QFontDatabase__Other               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(QFontDatabase__Symbol)
	QFontDatabase__Ogham               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(31)
	QFontDatabase__Runic               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(32)
	QFontDatabase__Nko                 QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(33)
	QFontDatabase__WritingSystemsCount QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(34)
)

type QFontMetrics struct {
	ptr unsafe.Pointer
}

type QFontMetrics_ITF interface {
	QFontMetrics_PTR() *QFontMetrics
}

func (ptr *QFontMetrics) QFontMetrics_PTR() *QFontMetrics {
	return ptr
}

func (ptr *QFontMetrics) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QFontMetrics) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQFontMetrics(ptr QFontMetrics_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontMetrics_PTR().Pointer()
	}
	return nil
}

func NewQFontMetricsFromPointer(ptr unsafe.Pointer) (n *QFontMetrics) {
	n = new(QFontMetrics)
	n.SetPointer(ptr)
	return
}
func NewQFontMetrics(font QFont_ITF) *QFontMetrics {
	tmpValue := NewQFontMetricsFromPointer(C.QFontMetrics_NewQFontMetrics(PointerFromQFont(font)))
	qt.SetFinalizer(tmpValue, (*QFontMetrics).DestroyQFontMetrics)
	return tmpValue
}

func NewQFontMetrics3(font QFont_ITF, paintdevice QPaintDevice_ITF) *QFontMetrics {
	tmpValue := NewQFontMetricsFromPointer(C.QFontMetrics_NewQFontMetrics3(PointerFromQFont(font), PointerFromQPaintDevice(paintdevice)))
	qt.SetFinalizer(tmpValue, (*QFontMetrics).DestroyQFontMetrics)
	return tmpValue
}

func NewQFontMetrics4(fm QFontMetrics_ITF) *QFontMetrics {
	tmpValue := NewQFontMetricsFromPointer(C.QFontMetrics_NewQFontMetrics4(PointerFromQFontMetrics(fm)))
	qt.SetFinalizer(tmpValue, (*QFontMetrics).DestroyQFontMetrics)
	return tmpValue
}

func (ptr *QFontMetrics) BoundingRect(ch core.QChar_ITF) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QFontMetrics_BoundingRect(ptr.Pointer(), core.PointerFromQChar(ch)))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QFontMetrics) BoundingRect2(text string) *core.QRect {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := core.NewQRectFromPointer(C.QFontMetrics_BoundingRect2(ptr.Pointer(), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QFontMetrics) BoundingRect3(rect core.QRect_ITF, flags int, text string, tabStops int, tabArray int) *core.QRect {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := core.NewQRectFromPointer(C.QFontMetrics_BoundingRect3(ptr.Pointer(), core.PointerFromQRect(rect), C.int(int32(flags)), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}, C.int(int32(tabStops)), C.int(int32(tabArray))))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QFontMetrics) BoundingRect4(x int, y int, width int, height int, flags int, text string, tabStops int, tabArray int) *core.QRect {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := core.NewQRectFromPointer(C.QFontMetrics_BoundingRect4(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(width)), C.int(int32(height)), C.int(int32(flags)), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}, C.int(int32(tabStops)), C.int(int32(tabArray))))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QFontMetrics) Height() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QFontMetrics_Height(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontMetrics) Size(flags int, text string, tabStops int, tabArray int) *core.QSize {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := core.NewQSizeFromPointer(C.QFontMetrics_Size(ptr.Pointer(), C.int(int32(flags)), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}, C.int(int32(tabStops)), C.int(int32(tabArray))))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QFontMetrics) DestroyQFontMetrics() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QFontMetrics_DestroyQFontMetrics(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGlyphRun struct {
	ptr unsafe.Pointer
}

type QGlyphRun_ITF interface {
	QGlyphRun_PTR() *QGlyphRun
}

func (ptr *QGlyphRun) QGlyphRun_PTR() *QGlyphRun {
	return ptr
}

func (ptr *QGlyphRun) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGlyphRun) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGlyphRun(ptr QGlyphRun_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGlyphRun_PTR().Pointer()
	}
	return nil
}

func NewQGlyphRunFromPointer(ptr unsafe.Pointer) (n *QGlyphRun) {
	n = new(QGlyphRun)
	n.SetPointer(ptr)
	return
}

// QGlyphRun::GlyphRunFlag
//
//go:generate stringer -type=QGlyphRun__GlyphRunFlag
type QGlyphRun__GlyphRunFlag int64

const (
	QGlyphRun__Overline      QGlyphRun__GlyphRunFlag = QGlyphRun__GlyphRunFlag(0x01)
	QGlyphRun__Underline     QGlyphRun__GlyphRunFlag = QGlyphRun__GlyphRunFlag(0x02)
	QGlyphRun__StrikeOut     QGlyphRun__GlyphRunFlag = QGlyphRun__GlyphRunFlag(0x04)
	QGlyphRun__RightToLeft   QGlyphRun__GlyphRunFlag = QGlyphRun__GlyphRunFlag(0x08)
	QGlyphRun__SplitLigature QGlyphRun__GlyphRunFlag = QGlyphRun__GlyphRunFlag(0x10)
)

func NewQGlyphRun() *QGlyphRun {
	tmpValue := NewQGlyphRunFromPointer(C.QGlyphRun_NewQGlyphRun())
	qt.SetFinalizer(tmpValue, (*QGlyphRun).DestroyQGlyphRun)
	return tmpValue
}

func NewQGlyphRun2(other QGlyphRun_ITF) *QGlyphRun {
	tmpValue := NewQGlyphRunFromPointer(C.QGlyphRun_NewQGlyphRun2(PointerFromQGlyphRun(other)))
	qt.SetFinalizer(tmpValue, (*QGlyphRun).DestroyQGlyphRun)
	return tmpValue
}

func (ptr *QGlyphRun) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QGlyphRun_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGlyphRun) Clear() {
	if ptr.Pointer() != nil {
		C.QGlyphRun_Clear(ptr.Pointer())
	}
}

func (ptr *QGlyphRun) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGlyphRun_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGlyphRun) StrikeOut() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGlyphRun_StrikeOut(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGlyphRun) Underline() bool {
	if ptr.Pointer() != nil {
		return int8(C.QGlyphRun_Underline(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGlyphRun) DestroyQGlyphRun() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGlyphRun_DestroyQGlyphRun(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGlyphRun) __glyphIndexes_atList(i int) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QGlyphRun___glyphIndexes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QGlyphRun) __glyphIndexes_setList(i uint) {
	if ptr.Pointer() != nil {
		C.QGlyphRun___glyphIndexes_setList(ptr.Pointer(), C.uint(uint32(i)))
	}
}

func (ptr *QGlyphRun) __glyphIndexes_newList() unsafe.Pointer {
	return C.QGlyphRun___glyphIndexes_newList(ptr.Pointer())
}

func (ptr *QGlyphRun) __positions_atList(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QGlyphRun___positions_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QGlyphRun) __positions_setList(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QGlyphRun___positions_setList(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QGlyphRun) __positions_newList() unsafe.Pointer {
	return C.QGlyphRun___positions_newList(ptr.Pointer())
}

func (ptr *QGlyphRun) __setGlyphIndexes_glyphIndexes_atList(i int) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QGlyphRun___setGlyphIndexes_glyphIndexes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QGlyphRun) __setGlyphIndexes_glyphIndexes_setList(i uint) {
	if ptr.Pointer() != nil {
		C.QGlyphRun___setGlyphIndexes_glyphIndexes_setList(ptr.Pointer(), C.uint(uint32(i)))
	}
}

func (ptr *QGlyphRun) __setGlyphIndexes_glyphIndexes_newList() unsafe.Pointer {
	return C.QGlyphRun___setGlyphIndexes_glyphIndexes_newList(ptr.Pointer())
}

func (ptr *QGlyphRun) __setPositions_positions_atList(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QGlyphRun___setPositions_positions_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QGlyphRun) __setPositions_positions_setList(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QGlyphRun___setPositions_positions_setList(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QGlyphRun) __setPositions_positions_newList() unsafe.Pointer {
	return C.QGlyphRun___setPositions_positions_newList(ptr.Pointer())
}

type QGradient struct {
	ptr unsafe.Pointer
}

type QGradient_ITF interface {
	QGradient_PTR() *QGradient
}

func (ptr *QGradient) QGradient_PTR() *QGradient {
	return ptr
}

func (ptr *QGradient) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGradient) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGradient(ptr QGradient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGradient_PTR().Pointer()
	}
	return nil
}

func NewQGradientFromPointer(ptr unsafe.Pointer) (n *QGradient) {
	n = new(QGradient)
	n.SetPointer(ptr)
	return
}
func (ptr *QGradient) DestroyQGradient() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QGradient__CoordinateMode
//QGradient::CoordinateMode
type QGradient__CoordinateMode int64

const (
	QGradient__LogicalMode         QGradient__CoordinateMode = QGradient__CoordinateMode(0)
	QGradient__StretchToDeviceMode QGradient__CoordinateMode = QGradient__CoordinateMode(1)
	QGradient__ObjectBoundingMode  QGradient__CoordinateMode = QGradient__CoordinateMode(2)
	QGradient__ObjectMode          QGradient__CoordinateMode = QGradient__CoordinateMode(3)
)

// QGradient::Preset
//
//go:generate stringer -type=QGradient__Preset
type QGradient__Preset int64

const (
	QGradient__WarmFlame        QGradient__Preset = QGradient__Preset(1)
	QGradient__NightFade        QGradient__Preset = QGradient__Preset(2)
	QGradient__SpringWarmth     QGradient__Preset = QGradient__Preset(3)
	QGradient__JuicyPeach       QGradient__Preset = QGradient__Preset(4)
	QGradient__YoungPassion     QGradient__Preset = QGradient__Preset(5)
	QGradient__LadyLips         QGradient__Preset = QGradient__Preset(6)
	QGradient__SunnyMorning     QGradient__Preset = QGradient__Preset(7)
	QGradient__RainyAshville    QGradient__Preset = QGradient__Preset(8)
	QGradient__FrozenDreams     QGradient__Preset = QGradient__Preset(9)
	QGradient__WinterNeva       QGradient__Preset = QGradient__Preset(10)
	QGradient__DustyGrass       QGradient__Preset = QGradient__Preset(11)
	QGradient__TemptingAzure    QGradient__Preset = QGradient__Preset(12)
	QGradient__HeavyRain        QGradient__Preset = QGradient__Preset(13)
	QGradient__AmyCrisp         QGradient__Preset = QGradient__Preset(14)
	QGradient__MeanFruit        QGradient__Preset = QGradient__Preset(15)
	QGradient__DeepBlue         QGradient__Preset = QGradient__Preset(16)
	QGradient__RipeMalinka      QGradient__Preset = QGradient__Preset(17)
	QGradient__CloudyKnoxville  QGradient__Preset = QGradient__Preset(18)
	QGradient__MalibuBeach      QGradient__Preset = QGradient__Preset(19)
	QGradient__NewLife          QGradient__Preset = QGradient__Preset(20)
	QGradient__TrueSunset       QGradient__Preset = QGradient__Preset(21)
	QGradient__MorpheusDen      QGradient__Preset = QGradient__Preset(22)
	QGradient__RareWind         QGradient__Preset = QGradient__Preset(23)
	QGradient__NearMoon         QGradient__Preset = QGradient__Preset(24)
	QGradient__WildApple        QGradient__Preset = QGradient__Preset(25)
	QGradient__SaintPetersburg  QGradient__Preset = QGradient__Preset(26)
	QGradient__PlumPlate        QGradient__Preset = QGradient__Preset(28)
	QGradient__EverlastingSky   QGradient__Preset = QGradient__Preset(29)
	QGradient__HappyFisher      QGradient__Preset = QGradient__Preset(30)
	QGradient__Blessing         QGradient__Preset = QGradient__Preset(31)
	QGradient__SharpeyeEagle    QGradient__Preset = QGradient__Preset(32)
	QGradient__LadogaBottom     QGradient__Preset = QGradient__Preset(33)
	QGradient__LemonGate        QGradient__Preset = QGradient__Preset(34)
	QGradient__ItmeoBranding    QGradient__Preset = QGradient__Preset(35)
	QGradient__ZeusMiracle      QGradient__Preset = QGradient__Preset(36)
	QGradient__OldHat           QGradient__Preset = QGradient__Preset(37)
	QGradient__StarWine         QGradient__Preset = QGradient__Preset(38)
	QGradient__HappyAcid        QGradient__Preset = QGradient__Preset(41)
	QGradient__AwesomePine      QGradient__Preset = QGradient__Preset(42)
	QGradient__NewYork          QGradient__Preset = QGradient__Preset(43)
	QGradient__ShyRainbow       QGradient__Preset = QGradient__Preset(44)
	QGradient__MixedHopes       QGradient__Preset = QGradient__Preset(46)
	QGradient__FlyHigh          QGradient__Preset = QGradient__Preset(47)
	QGradient__StrongBliss      QGradient__Preset = QGradient__Preset(48)
	QGradient__FreshMilk        QGradient__Preset = QGradient__Preset(49)
	QGradient__SnowAgain        QGradient__Preset = QGradient__Preset(50)
	QGradient__FebruaryInk      QGradient__Preset = QGradient__Preset(51)
	QGradient__KindSteel        QGradient__Preset = QGradient__Preset(52)
	QGradient__SoftGrass        QGradient__Preset = QGradient__Preset(53)
	QGradient__GrownEarly       QGradient__Preset = QGradient__Preset(54)
	QGradient__SharpBlues       QGradient__Preset = QGradient__Preset(55)
	QGradient__ShadyWater       QGradient__Preset = QGradient__Preset(56)
	QGradient__DirtyBeauty      QGradient__Preset = QGradient__Preset(57)
	QGradient__GreatWhale       QGradient__Preset = QGradient__Preset(58)
	QGradient__TeenNotebook     QGradient__Preset = QGradient__Preset(59)
	QGradient__PoliteRumors     QGradient__Preset = QGradient__Preset(60)
	QGradient__SweetPeriod      QGradient__Preset = QGradient__Preset(61)
	QGradient__WideMatrix       QGradient__Preset = QGradient__Preset(62)
	QGradient__SoftCherish      QGradient__Preset = QGradient__Preset(63)
	QGradient__RedSalvation     QGradient__Preset = QGradient__Preset(64)
	QGradient__BurningSpring    QGradient__Preset = QGradient__Preset(65)
	QGradient__NightParty       QGradient__Preset = QGradient__Preset(66)
	QGradient__SkyGlider        QGradient__Preset = QGradient__Preset(67)
	QGradient__HeavenPeach      QGradient__Preset = QGradient__Preset(68)
	QGradient__PurpleDivision   QGradient__Preset = QGradient__Preset(69)
	QGradient__AquaSplash       QGradient__Preset = QGradient__Preset(70)
	QGradient__SpikyNaga        QGradient__Preset = QGradient__Preset(72)
	QGradient__LoveKiss         QGradient__Preset = QGradient__Preset(73)
	QGradient__CleanMirror      QGradient__Preset = QGradient__Preset(75)
	QGradient__PremiumDark      QGradient__Preset = QGradient__Preset(76)
	QGradient__ColdEvening      QGradient__Preset = QGradient__Preset(77)
	QGradient__CochitiLake      QGradient__Preset = QGradient__Preset(78)
	QGradient__SummerGames      QGradient__Preset = QGradient__Preset(79)
	QGradient__PassionateBed    QGradient__Preset = QGradient__Preset(80)
	QGradient__MountainRock     QGradient__Preset = QGradient__Preset(81)
	QGradient__DesertHump       QGradient__Preset = QGradient__Preset(82)
	QGradient__JungleDay        QGradient__Preset = QGradient__Preset(83)
	QGradient__PhoenixStart     QGradient__Preset = QGradient__Preset(84)
	QGradient__OctoberSilence   QGradient__Preset = QGradient__Preset(85)
	QGradient__FarawayRiver     QGradient__Preset = QGradient__Preset(86)
	QGradient__AlchemistLab     QGradient__Preset = QGradient__Preset(87)
	QGradient__OverSun          QGradient__Preset = QGradient__Preset(88)
	QGradient__PremiumWhite     QGradient__Preset = QGradient__Preset(89)
	QGradient__MarsParty        QGradient__Preset = QGradient__Preset(90)
	QGradient__EternalConstance QGradient__Preset = QGradient__Preset(91)
	QGradient__JapanBlush       QGradient__Preset = QGradient__Preset(92)
	QGradient__SmilingRain      QGradient__Preset = QGradient__Preset(93)
	QGradient__CloudyApple      QGradient__Preset = QGradient__Preset(94)
	QGradient__BigMango         QGradient__Preset = QGradient__Preset(95)
	QGradient__HealthyWater     QGradient__Preset = QGradient__Preset(96)
	QGradient__AmourAmour       QGradient__Preset = QGradient__Preset(97)
	QGradient__RiskyConcrete    QGradient__Preset = QGradient__Preset(98)
	QGradient__StrongStick      QGradient__Preset = QGradient__Preset(99)
	QGradient__ViciousStance    QGradient__Preset = QGradient__Preset(100)
	QGradient__PaloAlto         QGradient__Preset = QGradient__Preset(101)
	QGradient__HappyMemories    QGradient__Preset = QGradient__Preset(102)
	QGradient__MidnightBloom    QGradient__Preset = QGradient__Preset(103)
	QGradient__Crystalline      QGradient__Preset = QGradient__Preset(104)
	QGradient__PartyBliss       QGradient__Preset = QGradient__Preset(106)
	QGradient__ConfidentCloud   QGradient__Preset = QGradient__Preset(107)
	QGradient__LeCocktail       QGradient__Preset = QGradient__Preset(108)
	QGradient__RiverCity        QGradient__Preset = QGradient__Preset(109)
	QGradient__FrozenBerry      QGradient__Preset = QGradient__Preset(110)
	QGradient__ChildCare        QGradient__Preset = QGradient__Preset(112)
	QGradient__FlyingLemon      QGradient__Preset = QGradient__Preset(113)
	QGradient__NewRetrowave     QGradient__Preset = QGradient__Preset(114)
	QGradient__HiddenJaguar     QGradient__Preset = QGradient__Preset(115)
	QGradient__AboveTheSky      QGradient__Preset = QGradient__Preset(116)
	QGradient__Nega             QGradient__Preset = QGradient__Preset(117)
	QGradient__DenseWater       QGradient__Preset = QGradient__Preset(118)
	QGradient__Seashore         QGradient__Preset = QGradient__Preset(120)
	QGradient__MarbleWall       QGradient__Preset = QGradient__Preset(121)
	QGradient__CheerfulCaramel  QGradient__Preset = QGradient__Preset(122)
	QGradient__NightSky         QGradient__Preset = QGradient__Preset(123)
	QGradient__MagicLake        QGradient__Preset = QGradient__Preset(124)
	QGradient__YoungGrass       QGradient__Preset = QGradient__Preset(125)
	QGradient__ColorfulPeach    QGradient__Preset = QGradient__Preset(126)
	QGradient__GentleCare       QGradient__Preset = QGradient__Preset(127)
	QGradient__PlumBath         QGradient__Preset = QGradient__Preset(128)
	QGradient__HappyUnicorn     QGradient__Preset = QGradient__Preset(129)
	QGradient__AfricanField     QGradient__Preset = QGradient__Preset(131)
	QGradient__SolidStone       QGradient__Preset = QGradient__Preset(132)
	QGradient__OrangeJuice      QGradient__Preset = QGradient__Preset(133)
	QGradient__GlassWater       QGradient__Preset = QGradient__Preset(134)
	QGradient__NorthMiracle     QGradient__Preset = QGradient__Preset(136)
	QGradient__FruitBlend       QGradient__Preset = QGradient__Preset(137)
	QGradient__MillenniumPine   QGradient__Preset = QGradient__Preset(138)
	QGradient__HighFlight       QGradient__Preset = QGradient__Preset(139)
	QGradient__MoleHall         QGradient__Preset = QGradient__Preset(140)
	QGradient__SpaceShift       QGradient__Preset = QGradient__Preset(142)
	QGradient__ForestInei       QGradient__Preset = QGradient__Preset(143)
	QGradient__RoyalGarden      QGradient__Preset = QGradient__Preset(144)
	QGradient__RichMetal        QGradient__Preset = QGradient__Preset(145)
	QGradient__JuicyCake        QGradient__Preset = QGradient__Preset(146)
	QGradient__SmartIndigo      QGradient__Preset = QGradient__Preset(147)
	QGradient__SandStrike       QGradient__Preset = QGradient__Preset(148)
	QGradient__NorseBeauty      QGradient__Preset = QGradient__Preset(149)
	QGradient__AquaGuidance     QGradient__Preset = QGradient__Preset(150)
	QGradient__SunVeggie        QGradient__Preset = QGradient__Preset(151)
	QGradient__SeaLord          QGradient__Preset = QGradient__Preset(152)
	QGradient__BlackSea         QGradient__Preset = QGradient__Preset(153)
	QGradient__GrassShampoo     QGradient__Preset = QGradient__Preset(154)
	QGradient__LandingAircraft  QGradient__Preset = QGradient__Preset(155)
	QGradient__WitchDance       QGradient__Preset = QGradient__Preset(156)
	QGradient__SleeplessNight   QGradient__Preset = QGradient__Preset(157)
	QGradient__AngelCare        QGradient__Preset = QGradient__Preset(158)
	QGradient__CrystalRiver     QGradient__Preset = QGradient__Preset(159)
	QGradient__SoftLipstick     QGradient__Preset = QGradient__Preset(160)
	QGradient__SaltMountain     QGradient__Preset = QGradient__Preset(161)
	QGradient__PerfectWhite     QGradient__Preset = QGradient__Preset(162)
	QGradient__FreshOasis       QGradient__Preset = QGradient__Preset(163)
	QGradient__StrictNovember   QGradient__Preset = QGradient__Preset(164)
	QGradient__MorningSalad     QGradient__Preset = QGradient__Preset(165)
	QGradient__DeepRelief       QGradient__Preset = QGradient__Preset(166)
	QGradient__SeaStrike        QGradient__Preset = QGradient__Preset(167)
	QGradient__NightCall        QGradient__Preset = QGradient__Preset(168)
	QGradient__SupremeSky       QGradient__Preset = QGradient__Preset(169)
	QGradient__LightBlue        QGradient__Preset = QGradient__Preset(170)
	QGradient__MindCrawl        QGradient__Preset = QGradient__Preset(171)
	QGradient__LilyMeadow       QGradient__Preset = QGradient__Preset(172)
	QGradient__SugarLollipop    QGradient__Preset = QGradient__Preset(173)
	QGradient__SweetDessert     QGradient__Preset = QGradient__Preset(174)
	QGradient__MagicRay         QGradient__Preset = QGradient__Preset(175)
	QGradient__TeenParty        QGradient__Preset = QGradient__Preset(176)
	QGradient__FrozenHeat       QGradient__Preset = QGradient__Preset(177)
	QGradient__GagarinView      QGradient__Preset = QGradient__Preset(178)
	QGradient__FabledSunset     QGradient__Preset = QGradient__Preset(179)
	QGradient__PerfectBlue      QGradient__Preset = QGradient__Preset(180)
	QGradient__NumPresets       QGradient__Preset = QGradient__Preset(181)
)

// QGradient::Spread
//
//go:generate stringer -type=QGradient__Spread
type QGradient__Spread int64

const (
	QGradient__PadSpread     QGradient__Spread = QGradient__Spread(0)
	QGradient__ReflectSpread QGradient__Spread = QGradient__Spread(1)
	QGradient__RepeatSpread  QGradient__Spread = QGradient__Spread(2)
)

// QGradient::Type
//
//go:generate stringer -type=QGradient__Type
type QGradient__Type int64

const (
	QGradient__LinearGradient  QGradient__Type = QGradient__Type(0)
	QGradient__RadialGradient  QGradient__Type = QGradient__Type(1)
	QGradient__ConicalGradient QGradient__Type = QGradient__Type(2)
	QGradient__NoGradient      QGradient__Type = QGradient__Type(3)
)

func NewQGradient2(preset QGradient__Preset) *QGradient {
	tmpValue := NewQGradientFromPointer(C.QGradient_NewQGradient2(C.longlong(preset)))
	qt.SetFinalizer(tmpValue, (*QGradient).DestroyQGradient)
	return tmpValue
}

func (ptr *QGradient) Type() QGradient__Type {
	if ptr.Pointer() != nil {
		return QGradient__Type(C.QGradient_Type(ptr.Pointer()))
	}
	return 0
}

type QGuiApplication struct {
	core.QCoreApplication
}

type QGuiApplication_ITF interface {
	core.QCoreApplication_ITF
	QGuiApplication_PTR() *QGuiApplication
}

func (ptr *QGuiApplication) QGuiApplication_PTR() *QGuiApplication {
	return ptr
}

func (ptr *QGuiApplication) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QCoreApplication_PTR().Pointer()
	}
	return nil
}

func (ptr *QGuiApplication) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QCoreApplication_PTR().SetPointer(p)
	}
}

func PointerFromQGuiApplication(ptr QGuiApplication_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGuiApplication_PTR().Pointer()
	}
	return nil
}

func NewQGuiApplicationFromPointer(ptr unsafe.Pointer) (n *QGuiApplication) {
	n = new(QGuiApplication)
	n.SetPointer(ptr)
	return
}
func NewQGuiApplication(argc int, argv []string) *QGuiApplication {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	tmpValue := NewQGuiApplicationFromPointer(C.QGuiApplication_NewQGuiApplication(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGuiApplication_ApplicationDisplayName() string {
	return cGoUnpackString(C.QGuiApplication_QGuiApplication_ApplicationDisplayName())
}

func (ptr *QGuiApplication) ApplicationDisplayName() string {
	return cGoUnpackString(C.QGuiApplication_QGuiApplication_ApplicationDisplayName())
}

func QGuiApplication_Clipboard() *QClipboard {
	tmpValue := NewQClipboardFromPointer(C.QGuiApplication_QGuiApplication_Clipboard())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGuiApplication) Clipboard() *QClipboard {
	tmpValue := NewQClipboardFromPointer(C.QGuiApplication_QGuiApplication_Clipboard())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGuiApplication_Event
func callbackQGuiApplication_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGuiApplicationFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QGuiApplication) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGuiApplication_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

func QGuiApplication_Exec() int {
	return int(int32(C.QGuiApplication_QGuiApplication_Exec()))
}

func (ptr *QGuiApplication) Exec() int {
	return int(int32(C.QGuiApplication_QGuiApplication_Exec()))
}

func QGuiApplication_Font() *QFont {
	tmpValue := NewQFontFromPointer(C.QGuiApplication_QGuiApplication_Font())
	qt.SetFinalizer(tmpValue, (*QFont).DestroyQFont)
	return tmpValue
}

func (ptr *QGuiApplication) Font() *QFont {
	tmpValue := NewQFontFromPointer(C.QGuiApplication_QGuiApplication_Font())
	qt.SetFinalizer(tmpValue, (*QFont).DestroyQFont)
	return tmpValue
}

func QGuiApplication_SetApplicationDisplayName(name string) {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	C.QGuiApplication_QGuiApplication_SetApplicationDisplayName(C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))})
}

func (ptr *QGuiApplication) SetApplicationDisplayName(name string) {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	C.QGuiApplication_QGuiApplication_SetApplicationDisplayName(C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))})
}

func QGuiApplication_SetFont(font QFont_ITF) {
	C.QGuiApplication_QGuiApplication_SetFont(PointerFromQFont(font))
}

func (ptr *QGuiApplication) SetFont(font QFont_ITF) {
	C.QGuiApplication_QGuiApplication_SetFont(PointerFromQFont(font))
}

func QGuiApplication_SetWindowIcon(icon QIcon_ITF) {
	C.QGuiApplication_QGuiApplication_SetWindowIcon(PointerFromQIcon(icon))
}

func (ptr *QGuiApplication) SetWindowIcon(icon QIcon_ITF) {
	C.QGuiApplication_QGuiApplication_SetWindowIcon(PointerFromQIcon(icon))
}

func QGuiApplication_WindowIcon() *QIcon {
	tmpValue := NewQIconFromPointer(C.QGuiApplication_QGuiApplication_WindowIcon())
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func (ptr *QGuiApplication) WindowIcon() *QIcon {
	tmpValue := NewQIconFromPointer(C.QGuiApplication_QGuiApplication_WindowIcon())
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

//export callbackQGuiApplication_DestroyQGuiApplication
func callbackQGuiApplication_DestroyQGuiApplication(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGuiApplication"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGuiApplicationFromPointer(ptr).DestroyQGuiApplicationDefault()
	}
}

func (ptr *QGuiApplication) ConnectDestroyQGuiApplication(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGuiApplication"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGuiApplication", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGuiApplication", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGuiApplication) DisconnectDestroyQGuiApplication() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGuiApplication")
	}
}

func (ptr *QGuiApplication) DestroyQGuiApplication() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGuiApplication_DestroyQGuiApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGuiApplication) DestroyQGuiApplicationDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGuiApplication_DestroyQGuiApplicationDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGuiApplication) __screens_atList(i int) *QScreen {
	if ptr.Pointer() != nil {
		tmpValue := NewQScreenFromPointer(C.QGuiApplication___screens_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __screens_setList(i QScreen_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___screens_setList(ptr.Pointer(), PointerFromQScreen(i))
	}
}

func (ptr *QGuiApplication) __screens_newList() unsafe.Pointer {
	return C.QGuiApplication___screens_newList(ptr.Pointer())
}

func (ptr *QGuiApplication) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGuiApplication___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGuiApplication) __children_newList() unsafe.Pointer {
	return C.QGuiApplication___children_newList(ptr.Pointer())
}

func (ptr *QGuiApplication) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGuiApplication___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGuiApplication) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGuiApplication___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGuiApplication) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGuiApplication___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGuiApplication) __findChildren_newList() unsafe.Pointer {
	return C.QGuiApplication___findChildren_newList(ptr.Pointer())
}

func (ptr *QGuiApplication) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGuiApplication___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGuiApplication) __findChildren_newList3() unsafe.Pointer {
	return C.QGuiApplication___findChildren_newList3(ptr.Pointer())
}

//export callbackQGuiApplication_Quit
func callbackQGuiApplication_Quit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "quit"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGuiApplicationFromPointer(ptr).QuitDefault()
	}
}

func (ptr *QGuiApplication) QuitDefault() {
	if ptr.Pointer() != nil {
		C.QGuiApplication_QuitDefault(ptr.Pointer())
	}
}

//export callbackQGuiApplication_ChildEvent
func callbackQGuiApplication_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGuiApplicationFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGuiApplication) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGuiApplication_ConnectNotify
func callbackQGuiApplication_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGuiApplicationFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGuiApplication) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGuiApplication_CustomEvent
func callbackQGuiApplication_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQGuiApplicationFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGuiApplication) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGuiApplication_DeleteLater
func callbackQGuiApplication_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGuiApplicationFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGuiApplication) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGuiApplication_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQGuiApplication_Destroyed
func callbackQGuiApplication_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGuiApplication_DisconnectNotify
func callbackQGuiApplication_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGuiApplicationFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGuiApplication) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGuiApplication_EventFilter
func callbackQGuiApplication_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGuiApplicationFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGuiApplication) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGuiApplication_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQGuiApplication_ObjectNameChanged
func callbackQGuiApplication_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGui_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQGuiApplication_TimerEvent
func callbackQGuiApplication_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGuiApplicationFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGuiApplication) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QHideEvent struct {
	core.QEvent
}

type QHideEvent_ITF interface {
	core.QEvent_ITF
	QHideEvent_PTR() *QHideEvent
}

func (ptr *QHideEvent) QHideEvent_PTR() *QHideEvent {
	return ptr
}

func (ptr *QHideEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QHideEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQHideEvent(ptr QHideEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHideEvent_PTR().Pointer()
	}
	return nil
}

func NewQHideEventFromPointer(ptr unsafe.Pointer) (n *QHideEvent) {
	n = new(QHideEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QHideEvent) DestroyQHideEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQHideEvent() *QHideEvent {
	tmpValue := NewQHideEventFromPointer(C.QHideEvent_NewQHideEvent())
	qt.SetFinalizer(tmpValue, (*QHideEvent).DestroyQHideEvent)
	return tmpValue
}

type QIcon struct {
	ptr unsafe.Pointer
}

type QIcon_ITF interface {
	QIcon_PTR() *QIcon
}

func (ptr *QIcon) QIcon_PTR() *QIcon {
	return ptr
}

func (ptr *QIcon) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QIcon) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQIcon(ptr QIcon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIcon_PTR().Pointer()
	}
	return nil
}

func NewQIconFromPointer(ptr unsafe.Pointer) (n *QIcon) {
	n = new(QIcon)
	n.SetPointer(ptr)
	return
}

// QIcon::Mode
//
//go:generate stringer -type=QIcon__Mode
type QIcon__Mode int64

const (
	QIcon__Normal   QIcon__Mode = QIcon__Mode(0)
	QIcon__Disabled QIcon__Mode = QIcon__Mode(1)
	QIcon__Active   QIcon__Mode = QIcon__Mode(2)
	QIcon__Selected QIcon__Mode = QIcon__Mode(3)
)

// QIcon::State
//
//go:generate stringer -type=QIcon__State
type QIcon__State int64

const (
	QIcon__On  QIcon__State = QIcon__State(0)
	QIcon__Off QIcon__State = QIcon__State(1)
)

func NewQIcon() *QIcon {
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon())
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func NewQIcon2(pixmap QPixmap_ITF) *QIcon {
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon2(PointerFromQPixmap(pixmap)))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func NewQIcon3(other QIcon_ITF) *QIcon {
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon3(PointerFromQIcon(other)))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func NewQIcon4(other QIcon_ITF) *QIcon {
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon4(PointerFromQIcon(other)))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func NewQIcon5(fileName string) *QIcon {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon5(C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func NewQIcon6(engine QIconEngine_ITF) *QIcon {
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon6(PointerFromQIconEngine(engine)))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func QIcon_FromTheme(name string) *QIcon {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQIconFromPointer(C.QIcon_QIcon_FromTheme(C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func (ptr *QIcon) FromTheme(name string) *QIcon {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQIconFromPointer(C.QIcon_QIcon_FromTheme(C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))}))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func QIcon_FromTheme2(name string, fallback QIcon_ITF) *QIcon {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQIconFromPointer(C.QIcon_QIcon_FromTheme2(C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQIcon(fallback)))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func (ptr *QIcon) FromTheme2(name string, fallback QIcon_ITF) *QIcon {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQIconFromPointer(C.QIcon_QIcon_FromTheme2(C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQIcon(fallback)))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func (ptr *QIcon) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QIcon_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIcon) Pixmap(size core.QSize_ITF, mode QIcon__Mode, state QIcon__State) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QIcon_Pixmap(ptr.Pointer(), core.PointerFromQSize(size), C.longlong(mode), C.longlong(state)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QIcon) Pixmap2(w int, h int, mode QIcon__Mode, state QIcon__State) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QIcon_Pixmap2(ptr.Pointer(), C.int(int32(w)), C.int(int32(h)), C.longlong(mode), C.longlong(state)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QIcon) Pixmap3(extent int, mode QIcon__Mode, state QIcon__State) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QIcon_Pixmap3(ptr.Pointer(), C.int(int32(extent)), C.longlong(mode), C.longlong(state)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QIcon) Pixmap4(window QWindow_ITF, size core.QSize_ITF, mode QIcon__Mode, state QIcon__State) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QIcon_Pixmap4(ptr.Pointer(), PointerFromQWindow(window), core.PointerFromQSize(size), C.longlong(mode), C.longlong(state)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QIcon) DestroyQIcon() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIcon_DestroyQIcon(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIcon) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QIcon_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QIcon) __availableSizes_atList(i int) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QIcon___availableSizes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QIcon) __availableSizes_setList(i core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QIcon___availableSizes_setList(ptr.Pointer(), core.PointerFromQSize(i))
	}
}

func (ptr *QIcon) __availableSizes_newList() unsafe.Pointer {
	return C.QIcon___availableSizes_newList(ptr.Pointer())
}

type QIconEngine struct {
	ptr unsafe.Pointer
}

type QIconEngine_ITF interface {
	QIconEngine_PTR() *QIconEngine
}

func (ptr *QIconEngine) QIconEngine_PTR() *QIconEngine {
	return ptr
}

func (ptr *QIconEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QIconEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQIconEngine(ptr QIconEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIconEngine_PTR().Pointer()
	}
	return nil
}

func NewQIconEngineFromPointer(ptr unsafe.Pointer) (n *QIconEngine) {
	n = new(QIconEngine)
	n.SetPointer(ptr)
	return
}

// QIconEngine::IconEngineHook
//
//go:generate stringer -type=QIconEngine__IconEngineHook
type QIconEngine__IconEngineHook int64

const (
	QIconEngine__AvailableSizesHook QIconEngine__IconEngineHook = QIconEngine__IconEngineHook(1)
	QIconEngine__IconNameHook       QIconEngine__IconEngineHook = QIconEngine__IconEngineHook(2)
	QIconEngine__IsNullHook         QIconEngine__IconEngineHook = QIconEngine__IconEngineHook(3)
	QIconEngine__ScaledPixmapHook   QIconEngine__IconEngineHook = QIconEngine__IconEngineHook(4)
)

func NewQIconEngine() *QIconEngine {
	return NewQIconEngineFromPointer(C.QIconEngine_NewQIconEngine())
}

//export callbackQIconEngine_Clone
func callbackQIconEngine_Clone(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "clone"); signal != nil {
		return PointerFromQIconEngine((*(*func() *QIconEngine)(signal))())
	}

	return PointerFromQIconEngine(NewQIconEngine())
}

func (ptr *QIconEngine) ConnectClone(f func() *QIconEngine) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clone"); signal != nil {
			f := func() *QIconEngine {
				(*(*func() *QIconEngine)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clone", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clone", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectClone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clone")
	}
}

func (ptr *QIconEngine) Clone() *QIconEngine {
	if ptr.Pointer() != nil {
		return NewQIconEngineFromPointer(C.QIconEngine_Clone(ptr.Pointer()))
	}
	return nil
}

//export callbackQIconEngine_Key
func callbackQIconEngine_Key(ptr unsafe.Pointer) C.struct_QtGui_PackedString {
	if signal := qt.GetSignal(ptr, "key"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_QtGui_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQIconEngineFromPointer(ptr).KeyDefault()
	return C.struct_QtGui_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QIconEngine) ConnectKey(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "key"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "key", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "key", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "key")
	}
}

func (ptr *QIconEngine) Key() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QIconEngine_Key(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIconEngine) KeyDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QIconEngine_KeyDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQIconEngine_Paint
func callbackQIconEngine_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, rect unsafe.Pointer, mode C.longlong, state C.longlong) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		(*(*func(*QPainter, *core.QRect, QIcon__Mode, QIcon__State))(signal))(NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), QIcon__Mode(mode), QIcon__State(state))
	}

}

func (ptr *QIconEngine) ConnectPaint(f func(painter *QPainter, rect *core.QRect, mode QIcon__Mode, state QIcon__State)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paint"); signal != nil {
			f := func(painter *QPainter, rect *core.QRect, mode QIcon__Mode, state QIcon__State) {
				(*(*func(*QPainter, *core.QRect, QIcon__Mode, QIcon__State))(signal))(painter, rect, mode, state)
				f(painter, rect, mode, state)
			}
			qt.ConnectSignal(ptr.Pointer(), "paint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectPaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paint")
	}
}

func (ptr *QIconEngine) Paint(painter QPainter_ITF, rect core.QRect_ITF, mode QIcon__Mode, state QIcon__State) {
	if ptr.Pointer() != nil {
		C.QIconEngine_Paint(ptr.Pointer(), PointerFromQPainter(painter), core.PointerFromQRect(rect), C.longlong(mode), C.longlong(state))
	}
}

//export callbackQIconEngine_Pixmap
func callbackQIconEngine_Pixmap(ptr unsafe.Pointer, size unsafe.Pointer, mode C.longlong, state C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "pixmap"); signal != nil {
		return PointerFromQPixmap((*(*func(*core.QSize, QIcon__Mode, QIcon__State) *QPixmap)(signal))(core.NewQSizeFromPointer(size), QIcon__Mode(mode), QIcon__State(state)))
	}

	return PointerFromQPixmap(NewQIconEngineFromPointer(ptr).PixmapDefault(core.NewQSizeFromPointer(size), QIcon__Mode(mode), QIcon__State(state)))
}

func (ptr *QIconEngine) ConnectPixmap(f func(size *core.QSize, mode QIcon__Mode, state QIcon__State) *QPixmap) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "pixmap"); signal != nil {
			f := func(size *core.QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {
				(*(*func(*core.QSize, QIcon__Mode, QIcon__State) *QPixmap)(signal))(size, mode, state)
				return f(size, mode, state)
			}
			qt.ConnectSignal(ptr.Pointer(), "pixmap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pixmap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "pixmap")
	}
}

func (ptr *QIconEngine) Pixmap(size core.QSize_ITF, mode QIcon__Mode, state QIcon__State) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QIconEngine_Pixmap(ptr.Pointer(), core.PointerFromQSize(size), C.longlong(mode), C.longlong(state)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QIconEngine) PixmapDefault(size core.QSize_ITF, mode QIcon__Mode, state QIcon__State) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QIconEngine_PixmapDefault(ptr.Pointer(), core.PointerFromQSize(size), C.longlong(mode), C.longlong(state)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

//export callbackQIconEngine_Read
func callbackQIconEngine_Read(ptr unsafe.Pointer, in unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "read"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QDataStream) bool)(signal))(core.NewQDataStreamFromPointer(in)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQIconEngineFromPointer(ptr).ReadDefault(core.NewQDataStreamFromPointer(in)))))
}

func (ptr *QIconEngine) ConnectRead(f func(in *core.QDataStream) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "read"); signal != nil {
			f := func(in *core.QDataStream) bool {
				(*(*func(*core.QDataStream) bool)(signal))(in)
				return f(in)
			}
			qt.ConnectSignal(ptr.Pointer(), "read", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "read", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectRead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "read")
	}
}

func (ptr *QIconEngine) Read(in core.QDataStream_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIconEngine_Read(ptr.Pointer(), core.PointerFromQDataStream(in))) != 0
	}
	return false
}

func (ptr *QIconEngine) ReadDefault(in core.QDataStream_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIconEngine_ReadDefault(ptr.Pointer(), core.PointerFromQDataStream(in))) != 0
	}
	return false
}

//export callbackQIconEngine_Write
func callbackQIconEngine_Write(ptr unsafe.Pointer, out unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "write"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QDataStream) bool)(signal))(core.NewQDataStreamFromPointer(out)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQIconEngineFromPointer(ptr).WriteDefault(core.NewQDataStreamFromPointer(out)))))
}

func (ptr *QIconEngine) ConnectWrite(f func(out *core.QDataStream) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "write"); signal != nil {
			f := func(out *core.QDataStream) bool {
				(*(*func(*core.QDataStream) bool)(signal))(out)
				return f(out)
			}
			qt.ConnectSignal(ptr.Pointer(), "write", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "write", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectWrite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "write")
	}
}

func (ptr *QIconEngine) Write(out core.QDataStream_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIconEngine_Write(ptr.Pointer(), core.PointerFromQDataStream(out))) != 0
	}
	return false
}

func (ptr *QIconEngine) WriteDefault(out core.QDataStream_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIconEngine_WriteDefault(ptr.Pointer(), core.PointerFromQDataStream(out))) != 0
	}
	return false
}

//export callbackQIconEngine_DestroyQIconEngine
func callbackQIconEngine_DestroyQIconEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QIconEngine"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQIconEngineFromPointer(ptr).DestroyQIconEngineDefault()
	}
}

func (ptr *QIconEngine) ConnectDestroyQIconEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QIconEngine"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QIconEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QIconEngine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectDestroyQIconEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QIconEngine")
	}
}

func (ptr *QIconEngine) DestroyQIconEngine() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIconEngine_DestroyQIconEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIconEngine) DestroyQIconEngineDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIconEngine_DestroyQIconEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIconEngine) __availableSizes_atList(i int) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QIconEngine___availableSizes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QIconEngine) __availableSizes_setList(i core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QIconEngine___availableSizes_setList(ptr.Pointer(), core.PointerFromQSize(i))
	}
}

func (ptr *QIconEngine) __availableSizes_newList() unsafe.Pointer {
	return C.QIconEngine___availableSizes_newList(ptr.Pointer())
}

type QImage struct {
	QPaintDevice
}

type QImage_ITF interface {
	QPaintDevice_ITF
	QImage_PTR() *QImage
}

func (ptr *QImage) QImage_PTR() *QImage {
	return ptr
}

func (ptr *QImage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QImage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQImage(ptr QImage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImage_PTR().Pointer()
	}
	return nil
}

func NewQImageFromPointer(ptr unsafe.Pointer) (n *QImage) {
	n = new(QImage)
	n.SetPointer(ptr)
	return
}

// QImage::Format
//
//go:generate stringer -type=QImage__Format
type QImage__Format int64

const (
	QImage__Format_Invalid                QImage__Format = QImage__Format(0)
	QImage__Format_Mono                   QImage__Format = QImage__Format(1)
	QImage__Format_MonoLSB                QImage__Format = QImage__Format(2)
	QImage__Format_Indexed8               QImage__Format = QImage__Format(3)
	QImage__Format_RGB32                  QImage__Format = QImage__Format(4)
	QImage__Format_ARGB32                 QImage__Format = QImage__Format(5)
	QImage__Format_ARGB32_Premultiplied   QImage__Format = QImage__Format(6)
	QImage__Format_RGB16                  QImage__Format = QImage__Format(7)
	QImage__Format_ARGB8565_Premultiplied QImage__Format = QImage__Format(8)
	QImage__Format_RGB666                 QImage__Format = QImage__Format(9)
	QImage__Format_ARGB6666_Premultiplied QImage__Format = QImage__Format(10)
	QImage__Format_RGB555                 QImage__Format = QImage__Format(11)
	QImage__Format_ARGB8555_Premultiplied QImage__Format = QImage__Format(12)
	QImage__Format_RGB888                 QImage__Format = QImage__Format(13)
	QImage__Format_RGB444                 QImage__Format = QImage__Format(14)
	QImage__Format_ARGB4444_Premultiplied QImage__Format = QImage__Format(15)
	QImage__Format_RGBX8888               QImage__Format = QImage__Format(16)
	QImage__Format_RGBA8888               QImage__Format = QImage__Format(17)
	QImage__Format_RGBA8888_Premultiplied QImage__Format = QImage__Format(18)
	QImage__Format_BGR30                  QImage__Format = QImage__Format(19)
	QImage__Format_A2BGR30_Premultiplied  QImage__Format = QImage__Format(20)
	QImage__Format_RGB30                  QImage__Format = QImage__Format(21)
	QImage__Format_A2RGB30_Premultiplied  QImage__Format = QImage__Format(22)
	QImage__Format_Alpha8                 QImage__Format = QImage__Format(23)
	QImage__Format_Grayscale8             QImage__Format = QImage__Format(24)
	QImage__Format_RGBX64                 QImage__Format = QImage__Format(25)
	QImage__Format_RGBA64                 QImage__Format = QImage__Format(26)
	QImage__Format_RGBA64_Premultiplied   QImage__Format = QImage__Format(27)
	QImage__Format_Grayscale16            QImage__Format = QImage__Format(28)
	QImage__Format_BGR888                 QImage__Format = QImage__Format(29)
)

// QImage::InvertMode
//
//go:generate stringer -type=QImage__InvertMode
type QImage__InvertMode int64

const (
	QImage__InvertRgb  QImage__InvertMode = QImage__InvertMode(0)
	QImage__InvertRgba QImage__InvertMode = QImage__InvertMode(1)
)

func NewQImage() *QImage {
	return NewQImageFromPointer(C.QImage_NewQImage())
}

func NewQImage2(size core.QSize_ITF, format QImage__Format) *QImage {
	return NewQImageFromPointer(C.QImage_NewQImage2(core.PointerFromQSize(size), C.longlong(format)))
}

func NewQImage3(width int, height int, format QImage__Format) *QImage {
	return NewQImageFromPointer(C.QImage_NewQImage3(C.int(int32(width)), C.int(int32(height)), C.longlong(format)))
}

func NewQImage4(data string, width int, height int, format QImage__Format) *QImage {
	var dataC *C.char
	if data != "" {
		dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
	}
	return NewQImageFromPointer(C.QImage_NewQImage4(dataC, C.int(int32(width)), C.int(int32(height)), C.longlong(format)))
}

func NewQImage5(data string, width int, height int, format QImage__Format) *QImage {
	var dataC *C.char
	if data != "" {
		dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
	}
	return NewQImageFromPointer(C.QImage_NewQImage5(dataC, C.int(int32(width)), C.int(int32(height)), C.longlong(format)))
}

func NewQImage6(data string, width int, height int, bytesPerLine int, format QImage__Format) *QImage {
	var dataC *C.char
	if data != "" {
		dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
	}
	return NewQImageFromPointer(C.QImage_NewQImage6(dataC, C.int(int32(width)), C.int(int32(height)), C.int(int32(bytesPerLine)), C.longlong(format)))
}

func NewQImage7(data string, width int, height int, bytesPerLine int, format QImage__Format) *QImage {
	var dataC *C.char
	if data != "" {
		dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
	}
	return NewQImageFromPointer(C.QImage_NewQImage7(dataC, C.int(int32(width)), C.int(int32(height)), C.int(int32(bytesPerLine)), C.longlong(format)))
}

func NewQImage9(fileName string, format string) *QImage {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	return NewQImageFromPointer(C.QImage_NewQImage9(C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, formatC))
}

func NewQImage10(image QImage_ITF) *QImage {
	return NewQImageFromPointer(C.QImage_NewQImage10(PointerFromQImage(image)))
}

func NewQImage11(other QImage_ITF) *QImage {
	return NewQImageFromPointer(C.QImage_NewQImage11(PointerFromQImage(other)))
}

func (ptr *QImage) Color(i int) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QImage_Color(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QImage) Copy(rectangle core.QRect_ITF) *QImage {
	if ptr.Pointer() != nil {
		tmpValue := NewQImageFromPointer(C.QImage_Copy(ptr.Pointer(), core.PointerFromQRect(rectangle)))
		qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QImage) Copy2(x int, y int, width int, height int) *QImage {
	if ptr.Pointer() != nil {
		tmpValue := NewQImageFromPointer(C.QImage_Copy2(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(width)), C.int(int32(height))))
		qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QImage) Format() QImage__Format {
	if ptr.Pointer() != nil {
		return QImage__Format(C.QImage_Format(ptr.Pointer()))
	}
	return 0
}

func QImage_FromData(data []byte, size int, format string) *QImage {
	var dataC *C.char
	if len(data) != 0 {
		dataC = (*C.char)(unsafe.Pointer(&data[0]))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQImageFromPointer(C.QImage_QImage_FromData(dataC, C.int(int32(size)), formatC))
	qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
	return tmpValue
}

func (ptr *QImage) FromData(data []byte, size int, format string) *QImage {
	var dataC *C.char
	if len(data) != 0 {
		dataC = (*C.char)(unsafe.Pointer(&data[0]))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQImageFromPointer(C.QImage_QImage_FromData(dataC, C.int(int32(size)), formatC))
	qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
	return tmpValue
}

func QImage_FromData2(data core.QByteArray_ITF, format string) *QImage {
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQImageFromPointer(C.QImage_QImage_FromData2(core.PointerFromQByteArray(data), formatC))
	qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
	return tmpValue
}

func (ptr *QImage) FromData2(data core.QByteArray_ITF, format string) *QImage {
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	tmpValue := NewQImageFromPointer(C.QImage_QImage_FromData2(core.PointerFromQByteArray(data), formatC))
	qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
	return tmpValue
}

func (ptr *QImage) Load(fileName string, format string) bool {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QImage_Load(ptr.Pointer(), C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, formatC)) != 0
	}
	return false
}

func (ptr *QImage) Load2(device core.QIODevice_ITF, format string) bool {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QImage_Load2(ptr.Pointer(), core.PointerFromQIODevice(device), formatC)) != 0
	}
	return false
}

func (ptr *QImage) LoadFromData(data []byte, l int, format string) bool {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QImage_LoadFromData(ptr.Pointer(), dataC, C.int(int32(l)), formatC)) != 0
	}
	return false
}

func (ptr *QImage) LoadFromData2(data core.QByteArray_ITF, format string) bool {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QImage_LoadFromData2(ptr.Pointer(), core.PointerFromQByteArray(data), formatC)) != 0
	}
	return false
}

func (ptr *QImage) Rect() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QImage_Rect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QImage) Save(fileName string, format string, quality int) bool {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QImage_Save(ptr.Pointer(), C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, formatC, C.int(int32(quality)))) != 0
	}
	return false
}

func (ptr *QImage) Save2(device core.QIODevice_ITF, format string, quality int) bool {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QImage_Save2(ptr.Pointer(), core.PointerFromQIODevice(device), formatC, C.int(int32(quality)))) != 0
	}
	return false
}

func (ptr *QImage) Scaled(size core.QSize_ITF, aspectRatioMode core.Qt__AspectRatioMode, transformMode core.Qt__TransformationMode) *QImage {
	if ptr.Pointer() != nil {
		tmpValue := NewQImageFromPointer(C.QImage_Scaled(ptr.Pointer(), core.PointerFromQSize(size), C.longlong(aspectRatioMode), C.longlong(transformMode)))
		qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QImage) Scaled2(width int, height int, aspectRatioMode core.Qt__AspectRatioMode, transformMode core.Qt__TransformationMode) *QImage {
	if ptr.Pointer() != nil {
		tmpValue := NewQImageFromPointer(C.QImage_Scaled2(ptr.Pointer(), C.int(int32(width)), C.int(int32(height)), C.longlong(aspectRatioMode), C.longlong(transformMode)))
		qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QImage) ScaledToHeight(height int, mode core.Qt__TransformationMode) *QImage {
	if ptr.Pointer() != nil {
		tmpValue := NewQImageFromPointer(C.QImage_ScaledToHeight(ptr.Pointer(), C.int(int32(height)), C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

func (ptr *QImage) SetText(key string, text string) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QImage_SetText(ptr.Pointer(), C.struct_QtGui_PackedString{data: keyC, len: C.longlong(len(key))}, C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QImage) Size() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QImage_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QImage) Text(key string) string {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		return cGoUnpackString(C.QImage_Text(ptr.Pointer(), C.struct_QtGui_PackedString{data: keyC, len: C.longlong(len(key))}))
	}
	return ""
}

func (ptr *QImage) Valid(pos core.QPoint_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QImage_Valid(ptr.Pointer(), core.PointerFromQPoint(pos))) != 0
	}
	return false
}

func (ptr *QImage) Valid2(x int, y int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QImage_Valid2(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)))) != 0
	}
	return false
}

//export callbackQImage_DestroyQImage
func callbackQImage_DestroyQImage(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QImage"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQImageFromPointer(ptr).DestroyQImageDefault()
	}
}

func (ptr *QImage) ConnectDestroyQImage(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QImage"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QImage", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QImage", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QImage) DisconnectDestroyQImage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QImage")
	}
}

func (ptr *QImage) DestroyQImage() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QImage_DestroyQImage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QImage) DestroyQImageDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QImage_DestroyQImageDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QImage) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QImage_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QImage) __colorTable_atList(i int) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QImage___colorTable_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QImage) __colorTable_setList(i uint) {
	if ptr.Pointer() != nil {
		C.QImage___colorTable_setList(ptr.Pointer(), C.uint(uint32(i)))
	}
}

func (ptr *QImage) __colorTable_newList() unsafe.Pointer {
	return C.QImage___colorTable_newList(ptr.Pointer())
}

func (ptr *QImage) __convertToFormat_colorTable_atList3(i int) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QImage___convertToFormat_colorTable_atList3(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QImage) __convertToFormat_colorTable_setList3(i uint) {
	if ptr.Pointer() != nil {
		C.QImage___convertToFormat_colorTable_setList3(ptr.Pointer(), C.uint(uint32(i)))
	}
}

func (ptr *QImage) __convertToFormat_colorTable_newList3() unsafe.Pointer {
	return C.QImage___convertToFormat_colorTable_newList3(ptr.Pointer())
}

func (ptr *QImage) __setColorTable_colors_atList(i int) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QImage___setColorTable_colors_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QImage) __setColorTable_colors_setList(i uint) {
	if ptr.Pointer() != nil {
		C.QImage___setColorTable_colors_setList(ptr.Pointer(), C.uint(uint32(i)))
	}
}

func (ptr *QImage) __setColorTable_colors_newList() unsafe.Pointer {
	return C.QImage___setColorTable_colors_newList(ptr.Pointer())
}

//export callbackQImage_PaintEngine
func callbackQImage_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return PointerFromQPaintEngine((*(*func() *QPaintEngine)(signal))())
	}

	return PointerFromQPaintEngine(NewQImageFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QImage) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QImage_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QImage) PaintEngineDefault() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QImage_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

// QImageIOHandler::ImageOption
//
//go:generate stringer -type=QImageIOHandler__ImageOption
type QImageIOHandler__ImageOption int64

const (
	QImageIOHandler__Size                 QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(0)
	QImageIOHandler__ClipRect             QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(1)
	QImageIOHandler__Description          QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(2)
	QImageIOHandler__ScaledClipRect       QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(3)
	QImageIOHandler__ScaledSize           QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(4)
	QImageIOHandler__CompressionRatio     QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(5)
	QImageIOHandler__Gamma                QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(6)
	QImageIOHandler__Quality              QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(7)
	QImageIOHandler__Name                 QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(8)
	QImageIOHandler__SubType              QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(9)
	QImageIOHandler__IncrementalReading   QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(10)
	QImageIOHandler__Endianness           QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(11)
	QImageIOHandler__Animation            QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(12)
	QImageIOHandler__BackgroundColor      QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(13)
	QImageIOHandler__ImageFormat          QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(14)
	QImageIOHandler__SupportedSubTypes    QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(15)
	QImageIOHandler__OptimizedWrite       QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(16)
	QImageIOHandler__ProgressiveScanWrite QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(17)
	QImageIOHandler__ImageTransformation  QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(18)
	QImageIOHandler__TransformedByDefault QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(19)
)

// QImageIOHandler::Transformation
//
//go:generate stringer -type=QImageIOHandler__Transformation
type QImageIOHandler__Transformation int64

const (
	QImageIOHandler__TransformationNone              QImageIOHandler__Transformation = QImageIOHandler__Transformation(0)
	QImageIOHandler__TransformationMirror            QImageIOHandler__Transformation = QImageIOHandler__Transformation(1)
	QImageIOHandler__TransformationFlip              QImageIOHandler__Transformation = QImageIOHandler__Transformation(2)
	QImageIOHandler__TransformationRotate180         QImageIOHandler__Transformation = QImageIOHandler__Transformation(QImageIOHandler__TransformationMirror | QImageIOHandler__TransformationFlip)
	QImageIOHandler__TransformationRotate90          QImageIOHandler__Transformation = QImageIOHandler__Transformation(4)
	QImageIOHandler__TransformationMirrorAndRotate90 QImageIOHandler__Transformation = QImageIOHandler__Transformation(QImageIOHandler__TransformationMirror | QImageIOHandler__TransformationRotate90)
	QImageIOHandler__TransformationFlipAndRotate90   QImageIOHandler__Transformation = QImageIOHandler__Transformation(QImageIOHandler__TransformationFlip | QImageIOHandler__TransformationRotate90)
	QImageIOHandler__TransformationRotate270         QImageIOHandler__Transformation = QImageIOHandler__Transformation(QImageIOHandler__TransformationRotate180 | QImageIOHandler__TransformationRotate90)
)

// QImageIOPlugin::Capability
//
//go:generate stringer -type=QImageIOPlugin__Capability
type QImageIOPlugin__Capability int64

const (
	QImageIOPlugin__CanRead            QImageIOPlugin__Capability = QImageIOPlugin__Capability(0x1)
	QImageIOPlugin__CanWrite           QImageIOPlugin__Capability = QImageIOPlugin__Capability(0x2)
	QImageIOPlugin__CanReadIncremental QImageIOPlugin__Capability = QImageIOPlugin__Capability(0x4)
)

// QImageReader::ImageReaderError
//
//go:generate stringer -type=QImageReader__ImageReaderError
type QImageReader__ImageReaderError int64

const (
	QImageReader__UnknownError           QImageReader__ImageReaderError = QImageReader__ImageReaderError(0)
	QImageReader__FileNotFoundError      QImageReader__ImageReaderError = QImageReader__ImageReaderError(1)
	QImageReader__DeviceError            QImageReader__ImageReaderError = QImageReader__ImageReaderError(2)
	QImageReader__UnsupportedFormatError QImageReader__ImageReaderError = QImageReader__ImageReaderError(3)
	QImageReader__InvalidDataError       QImageReader__ImageReaderError = QImageReader__ImageReaderError(4)
)

type QImageTextKeyLang struct {
	ptr unsafe.Pointer
}

type QImageTextKeyLang_ITF interface {
	QImageTextKeyLang_PTR() *QImageTextKeyLang
}

func (ptr *QImageTextKeyLang) QImageTextKeyLang_PTR() *QImageTextKeyLang {
	return ptr
}

func (ptr *QImageTextKeyLang) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QImageTextKeyLang) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQImageTextKeyLang(ptr QImageTextKeyLang_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageTextKeyLang_PTR().Pointer()
	}
	return nil
}

func NewQImageTextKeyLangFromPointer(ptr unsafe.Pointer) (n *QImageTextKeyLang) {
	n = new(QImageTextKeyLang)
	n.SetPointer(ptr)
	return
}
func (ptr *QImageTextKeyLang) DestroyQImageTextKeyLang() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QImageWriter__ImageWriterError
//QImageWriter::ImageWriterError
type QImageWriter__ImageWriterError int64

const (
	QImageWriter__UnknownError           QImageWriter__ImageWriterError = QImageWriter__ImageWriterError(0)
	QImageWriter__DeviceError            QImageWriter__ImageWriterError = QImageWriter__ImageWriterError(1)
	QImageWriter__UnsupportedFormatError QImageWriter__ImageWriterError = QImageWriter__ImageWriterError(2)
	QImageWriter__InvalidImageError      QImageWriter__ImageWriterError = QImageWriter__ImageWriterError(3)
)

type QInputEvent struct {
	core.QEvent
}

type QInputEvent_ITF interface {
	core.QEvent_ITF
	QInputEvent_PTR() *QInputEvent
}

func (ptr *QInputEvent) QInputEvent_PTR() *QInputEvent {
	return ptr
}

func (ptr *QInputEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QInputEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQInputEvent(ptr QInputEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputEvent_PTR().Pointer()
	}
	return nil
}

func NewQInputEventFromPointer(ptr unsafe.Pointer) (n *QInputEvent) {
	n = new(QInputEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QInputEvent) DestroyQInputEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QInputMethod__Action
//QInputMethod::Action
type QInputMethod__Action int64

const (
	QInputMethod__Click       QInputMethod__Action = QInputMethod__Action(0)
	QInputMethod__ContextMenu QInputMethod__Action = QInputMethod__Action(1)
)

// QInputMethodEvent::AttributeType
//
//go:generate stringer -type=QInputMethodEvent__AttributeType
type QInputMethodEvent__AttributeType int64

const (
	QInputMethodEvent__TextFormat QInputMethodEvent__AttributeType = QInputMethodEvent__AttributeType(0)
	QInputMethodEvent__Cursor     QInputMethodEvent__AttributeType = QInputMethodEvent__AttributeType(1)
	QInputMethodEvent__Language   QInputMethodEvent__AttributeType = QInputMethodEvent__AttributeType(2)
	QInputMethodEvent__Ruby       QInputMethodEvent__AttributeType = QInputMethodEvent__AttributeType(3)
	QInputMethodEvent__Selection  QInputMethodEvent__AttributeType = QInputMethodEvent__AttributeType(4)
)

type QIntValidator struct {
	QValidator
}

type QIntValidator_ITF interface {
	QValidator_ITF
	QIntValidator_PTR() *QIntValidator
}

func (ptr *QIntValidator) QIntValidator_PTR() *QIntValidator {
	return ptr
}

func (ptr *QIntValidator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QValidator_PTR().Pointer()
	}
	return nil
}

func (ptr *QIntValidator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QValidator_PTR().SetPointer(p)
	}
}

func PointerFromQIntValidator(ptr QIntValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIntValidator_PTR().Pointer()
	}
	return nil
}

func NewQIntValidatorFromPointer(ptr unsafe.Pointer) (n *QIntValidator) {
	n = new(QIntValidator)
	n.SetPointer(ptr)
	return
}
func NewQIntValidator(parent core.QObject_ITF) *QIntValidator {
	tmpValue := NewQIntValidatorFromPointer(C.QIntValidator_NewQIntValidator(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQIntValidator2(minimum int, maximum int, parent core.QObject_ITF) *QIntValidator {
	tmpValue := NewQIntValidatorFromPointer(C.QIntValidator_NewQIntValidator2(C.int(int32(minimum)), C.int(int32(maximum)), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQIntValidator_SetRange
func callbackQIntValidator_SetRange(ptr unsafe.Pointer, bottom C.int, top C.int) {
	if signal := qt.GetSignal(ptr, "setRange"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(bottom)), int(int32(top)))
	} else {
		NewQIntValidatorFromPointer(ptr).SetRangeDefault(int(int32(bottom)), int(int32(top)))
	}
}

func (ptr *QIntValidator) ConnectSetRange(f func(bottom int, top int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRange"); signal != nil {
			f := func(bottom int, top int) {
				(*(*func(int, int))(signal))(bottom, top)
				f(bottom, top)
			}
			qt.ConnectSignal(ptr.Pointer(), "setRange", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRange", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIntValidator) DisconnectSetRange() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRange")
	}
}

func (ptr *QIntValidator) SetRange(bottom int, top int) {
	if ptr.Pointer() != nil {
		C.QIntValidator_SetRange(ptr.Pointer(), C.int(int32(bottom)), C.int(int32(top)))
	}
}

func (ptr *QIntValidator) SetRangeDefault(bottom int, top int) {
	if ptr.Pointer() != nil {
		C.QIntValidator_SetRangeDefault(ptr.Pointer(), C.int(int32(bottom)), C.int(int32(top)))
	}
}

func (ptr *QIntValidator) SetTop(vin int) {
	if ptr.Pointer() != nil {
		C.QIntValidator_SetTop(ptr.Pointer(), C.int(int32(vin)))
	}
}

func (ptr *QIntValidator) Top() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QIntValidator_Top(ptr.Pointer())))
	}
	return 0
}

//export callbackQIntValidator_DestroyQIntValidator
func callbackQIntValidator_DestroyQIntValidator(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QIntValidator"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQIntValidatorFromPointer(ptr).DestroyQIntValidatorDefault()
	}
}

func (ptr *QIntValidator) ConnectDestroyQIntValidator(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QIntValidator"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QIntValidator", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QIntValidator", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIntValidator) DisconnectDestroyQIntValidator() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QIntValidator")
	}
}

func (ptr *QIntValidator) DestroyQIntValidator() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIntValidator_DestroyQIntValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIntValidator) DestroyQIntValidatorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIntValidator_DestroyQIntValidatorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQIntValidator_Validate
func callbackQIntValidator_Validate(ptr unsafe.Pointer, input C.struct_QtGui_PackedString, pos C.int) C.longlong {
	if signal := qt.GetSignal(ptr, "validate"); signal != nil {
		return C.longlong((*(*func(string, int) QValidator__State)(signal))(cGoUnpackString(input), int(int32(pos))))
	}

	return C.longlong(NewQIntValidatorFromPointer(ptr).ValidateDefault(cGoUnpackString(input), int(int32(pos))))
}

func (ptr *QIntValidator) Validate(input string, pos int) QValidator__State {
	if ptr.Pointer() != nil {
		var inputC *C.char
		if input != "" {
			inputC = C.CString(input)
			defer C.free(unsafe.Pointer(inputC))
		}
		return QValidator__State(C.QIntValidator_Validate(ptr.Pointer(), C.struct_QtGui_PackedString{data: inputC, len: C.longlong(len(input))}, C.int(int32(pos))))
	}
	return 0
}

func (ptr *QIntValidator) ValidateDefault(input string, pos int) QValidator__State {
	if ptr.Pointer() != nil {
		var inputC *C.char
		if input != "" {
			inputC = C.CString(input)
			defer C.free(unsafe.Pointer(inputC))
		}
		return QValidator__State(C.QIntValidator_ValidateDefault(ptr.Pointer(), C.struct_QtGui_PackedString{data: inputC, len: C.longlong(len(input))}, C.int(int32(pos))))
	}
	return 0
}

type QKeyEvent struct {
	QInputEvent
}

type QKeyEvent_ITF interface {
	QInputEvent_ITF
	QKeyEvent_PTR() *QKeyEvent
}

func (ptr *QKeyEvent) QKeyEvent_PTR() *QKeyEvent {
	return ptr
}

func (ptr *QKeyEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QKeyEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QInputEvent_PTR().SetPointer(p)
	}
}

func PointerFromQKeyEvent(ptr QKeyEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeyEvent_PTR().Pointer()
	}
	return nil
}

func NewQKeyEventFromPointer(ptr unsafe.Pointer) (n *QKeyEvent) {
	n = new(QKeyEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QKeyEvent) DestroyQKeyEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQKeyEvent(ty core.QEvent__Type, key int, modifiers core.Qt__KeyboardModifier, text string, autorep bool, count uint16) *QKeyEvent {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQKeyEventFromPointer(C.QKeyEvent_NewQKeyEvent(C.longlong(ty), C.int(int32(key)), C.longlong(modifiers), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}, C.char(int8(qt.GoBoolToInt(autorep))), C.ushort(count)))
	qt.SetFinalizer(tmpValue, (*QKeyEvent).DestroyQKeyEvent)
	return tmpValue
}

func NewQKeyEvent2(ty core.QEvent__Type, key int, modifiers core.Qt__KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool, count uint16) *QKeyEvent {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQKeyEventFromPointer(C.QKeyEvent_NewQKeyEvent2(C.longlong(ty), C.int(int32(key)), C.longlong(modifiers), C.uint(uint32(nativeScanCode)), C.uint(uint32(nativeVirtualKey)), C.uint(uint32(nativeModifiers)), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}, C.char(int8(qt.GoBoolToInt(autorep))), C.ushort(count)))
	qt.SetFinalizer(tmpValue, (*QKeyEvent).DestroyQKeyEvent)
	return tmpValue
}

func (ptr *QKeyEvent) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QKeyEvent_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QKeyEvent) Key() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QKeyEvent_Key(ptr.Pointer())))
	}
	return 0
}

func (ptr *QKeyEvent) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QKeyEvent_Text(ptr.Pointer()))
	}
	return ""
}

type QKeySequence struct {
	ptr unsafe.Pointer
}

type QKeySequence_ITF interface {
	QKeySequence_PTR() *QKeySequence
}

func (ptr *QKeySequence) QKeySequence_PTR() *QKeySequence {
	return ptr
}

func (ptr *QKeySequence) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QKeySequence) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQKeySequence(ptr QKeySequence_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeySequence_PTR().Pointer()
	}
	return nil
}

func NewQKeySequenceFromPointer(ptr unsafe.Pointer) (n *QKeySequence) {
	n = new(QKeySequence)
	n.SetPointer(ptr)
	return
}

// QKeySequence::SequenceFormat
//
//go:generate stringer -type=QKeySequence__SequenceFormat
type QKeySequence__SequenceFormat int64

const (
	QKeySequence__NativeText   QKeySequence__SequenceFormat = QKeySequence__SequenceFormat(0)
	QKeySequence__PortableText QKeySequence__SequenceFormat = QKeySequence__SequenceFormat(1)
)

// QKeySequence::SequenceMatch
//
//go:generate stringer -type=QKeySequence__SequenceMatch
type QKeySequence__SequenceMatch int64

const (
	QKeySequence__NoMatch      QKeySequence__SequenceMatch = QKeySequence__SequenceMatch(0)
	QKeySequence__PartialMatch QKeySequence__SequenceMatch = QKeySequence__SequenceMatch(1)
	QKeySequence__ExactMatch   QKeySequence__SequenceMatch = QKeySequence__SequenceMatch(2)
)

// QKeySequence::StandardKey
//
//go:generate stringer -type=QKeySequence__StandardKey
type QKeySequence__StandardKey int64

const (
	QKeySequence__UnknownKey               QKeySequence__StandardKey = QKeySequence__StandardKey(0)
	QKeySequence__HelpContents             QKeySequence__StandardKey = QKeySequence__StandardKey(1)
	QKeySequence__WhatsThis                QKeySequence__StandardKey = QKeySequence__StandardKey(2)
	QKeySequence__Open                     QKeySequence__StandardKey = QKeySequence__StandardKey(3)
	QKeySequence__Close                    QKeySequence__StandardKey = QKeySequence__StandardKey(4)
	QKeySequence__Save                     QKeySequence__StandardKey = QKeySequence__StandardKey(5)
	QKeySequence__New                      QKeySequence__StandardKey = QKeySequence__StandardKey(6)
	QKeySequence__Delete                   QKeySequence__StandardKey = QKeySequence__StandardKey(7)
	QKeySequence__Cut                      QKeySequence__StandardKey = QKeySequence__StandardKey(8)
	QKeySequence__Copy                     QKeySequence__StandardKey = QKeySequence__StandardKey(9)
	QKeySequence__Paste                    QKeySequence__StandardKey = QKeySequence__StandardKey(10)
	QKeySequence__Undo                     QKeySequence__StandardKey = QKeySequence__StandardKey(11)
	QKeySequence__Redo                     QKeySequence__StandardKey = QKeySequence__StandardKey(12)
	QKeySequence__Back                     QKeySequence__StandardKey = QKeySequence__StandardKey(13)
	QKeySequence__Forward                  QKeySequence__StandardKey = QKeySequence__StandardKey(14)
	QKeySequence__Refresh                  QKeySequence__StandardKey = QKeySequence__StandardKey(15)
	QKeySequence__ZoomIn                   QKeySequence__StandardKey = QKeySequence__StandardKey(16)
	QKeySequence__ZoomOut                  QKeySequence__StandardKey = QKeySequence__StandardKey(17)
	QKeySequence__Print                    QKeySequence__StandardKey = QKeySequence__StandardKey(18)
	QKeySequence__AddTab                   QKeySequence__StandardKey = QKeySequence__StandardKey(19)
	QKeySequence__NextChild                QKeySequence__StandardKey = QKeySequence__StandardKey(20)
	QKeySequence__PreviousChild            QKeySequence__StandardKey = QKeySequence__StandardKey(21)
	QKeySequence__Find                     QKeySequence__StandardKey = QKeySequence__StandardKey(22)
	QKeySequence__FindNext                 QKeySequence__StandardKey = QKeySequence__StandardKey(23)
	QKeySequence__FindPrevious             QKeySequence__StandardKey = QKeySequence__StandardKey(24)
	QKeySequence__Replace                  QKeySequence__StandardKey = QKeySequence__StandardKey(25)
	QKeySequence__SelectAll                QKeySequence__StandardKey = QKeySequence__StandardKey(26)
	QKeySequence__Bold                     QKeySequence__StandardKey = QKeySequence__StandardKey(27)
	QKeySequence__Italic                   QKeySequence__StandardKey = QKeySequence__StandardKey(28)
	QKeySequence__Underline                QKeySequence__StandardKey = QKeySequence__StandardKey(29)
	QKeySequence__MoveToNextChar           QKeySequence__StandardKey = QKeySequence__StandardKey(30)
	QKeySequence__MoveToPreviousChar       QKeySequence__StandardKey = QKeySequence__StandardKey(31)
	QKeySequence__MoveToNextWord           QKeySequence__StandardKey = QKeySequence__StandardKey(32)
	QKeySequence__MoveToPreviousWord       QKeySequence__StandardKey = QKeySequence__StandardKey(33)
	QKeySequence__MoveToNextLine           QKeySequence__StandardKey = QKeySequence__StandardKey(34)
	QKeySequence__MoveToPreviousLine       QKeySequence__StandardKey = QKeySequence__StandardKey(35)
	QKeySequence__MoveToNextPage           QKeySequence__StandardKey = QKeySequence__StandardKey(36)
	QKeySequence__MoveToPreviousPage       QKeySequence__StandardKey = QKeySequence__StandardKey(37)
	QKeySequence__MoveToStartOfLine        QKeySequence__StandardKey = QKeySequence__StandardKey(38)
	QKeySequence__MoveToEndOfLine          QKeySequence__StandardKey = QKeySequence__StandardKey(39)
	QKeySequence__MoveToStartOfBlock       QKeySequence__StandardKey = QKeySequence__StandardKey(40)
	QKeySequence__MoveToEndOfBlock         QKeySequence__StandardKey = QKeySequence__StandardKey(41)
	QKeySequence__MoveToStartOfDocument    QKeySequence__StandardKey = QKeySequence__StandardKey(42)
	QKeySequence__MoveToEndOfDocument      QKeySequence__StandardKey = QKeySequence__StandardKey(43)
	QKeySequence__SelectNextChar           QKeySequence__StandardKey = QKeySequence__StandardKey(44)
	QKeySequence__SelectPreviousChar       QKeySequence__StandardKey = QKeySequence__StandardKey(45)
	QKeySequence__SelectNextWord           QKeySequence__StandardKey = QKeySequence__StandardKey(46)
	QKeySequence__SelectPreviousWord       QKeySequence__StandardKey = QKeySequence__StandardKey(47)
	QKeySequence__SelectNextLine           QKeySequence__StandardKey = QKeySequence__StandardKey(48)
	QKeySequence__SelectPreviousLine       QKeySequence__StandardKey = QKeySequence__StandardKey(49)
	QKeySequence__SelectNextPage           QKeySequence__StandardKey = QKeySequence__StandardKey(50)
	QKeySequence__SelectPreviousPage       QKeySequence__StandardKey = QKeySequence__StandardKey(51)
	QKeySequence__SelectStartOfLine        QKeySequence__StandardKey = QKeySequence__StandardKey(52)
	QKeySequence__SelectEndOfLine          QKeySequence__StandardKey = QKeySequence__StandardKey(53)
	QKeySequence__SelectStartOfBlock       QKeySequence__StandardKey = QKeySequence__StandardKey(54)
	QKeySequence__SelectEndOfBlock         QKeySequence__StandardKey = QKeySequence__StandardKey(55)
	QKeySequence__SelectStartOfDocument    QKeySequence__StandardKey = QKeySequence__StandardKey(56)
	QKeySequence__SelectEndOfDocument      QKeySequence__StandardKey = QKeySequence__StandardKey(57)
	QKeySequence__DeleteStartOfWord        QKeySequence__StandardKey = QKeySequence__StandardKey(58)
	QKeySequence__DeleteEndOfWord          QKeySequence__StandardKey = QKeySequence__StandardKey(59)
	QKeySequence__DeleteEndOfLine          QKeySequence__StandardKey = QKeySequence__StandardKey(60)
	QKeySequence__InsertParagraphSeparator QKeySequence__StandardKey = QKeySequence__StandardKey(61)
	QKeySequence__InsertLineSeparator      QKeySequence__StandardKey = QKeySequence__StandardKey(62)
	QKeySequence__SaveAs                   QKeySequence__StandardKey = QKeySequence__StandardKey(63)
	QKeySequence__Preferences              QKeySequence__StandardKey = QKeySequence__StandardKey(64)
	QKeySequence__Quit                     QKeySequence__StandardKey = QKeySequence__StandardKey(65)
	QKeySequence__FullScreen               QKeySequence__StandardKey = QKeySequence__StandardKey(66)
	QKeySequence__Deselect                 QKeySequence__StandardKey = QKeySequence__StandardKey(67)
	QKeySequence__DeleteCompleteLine       QKeySequence__StandardKey = QKeySequence__StandardKey(68)
	QKeySequence__Backspace                QKeySequence__StandardKey = QKeySequence__StandardKey(69)
	QKeySequence__Cancel                   QKeySequence__StandardKey = QKeySequence__StandardKey(70)
)

func NewQKeySequence() *QKeySequence {
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence())
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func NewQKeySequence2(key string, format QKeySequence__SequenceFormat) *QKeySequence {
	var keyC *C.char
	if key != "" {
		keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
	}
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence2(C.struct_QtGui_PackedString{data: keyC, len: C.longlong(len(key))}, C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func NewQKeySequence3(k1 int, k2 int, k3 int, k4 int) *QKeySequence {
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence3(C.int(int32(k1)), C.int(int32(k2)), C.int(int32(k3)), C.int(int32(k4))))
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func NewQKeySequence4(keysequence QKeySequence_ITF) *QKeySequence {
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence4(PointerFromQKeySequence(keysequence)))
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func NewQKeySequence5(key QKeySequence__StandardKey) *QKeySequence {
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence5(C.longlong(key)))
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func (ptr *QKeySequence) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QKeySequence_Count(ptr.Pointer())))
	}
	return 0
}

func QKeySequence_FromString(str string, format QKeySequence__SequenceFormat) *QKeySequence {
	var strC *C.char
	if str != "" {
		strC = C.CString(str)
		defer C.free(unsafe.Pointer(strC))
	}
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_QKeySequence_FromString(C.struct_QtGui_PackedString{data: strC, len: C.longlong(len(str))}, C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func (ptr *QKeySequence) FromString(str string, format QKeySequence__SequenceFormat) *QKeySequence {
	var strC *C.char
	if str != "" {
		strC = C.CString(str)
		defer C.free(unsafe.Pointer(strC))
	}
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_QKeySequence_FromString(C.struct_QtGui_PackedString{data: strC, len: C.longlong(len(str))}, C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func (ptr *QKeySequence) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QKeySequence_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QKeySequence) DestroyQKeySequence() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QKeySequence_DestroyQKeySequence(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QKeySequence) __keyBindings_atList(i int) *QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := NewQKeySequenceFromPointer(C.QKeySequence___keyBindings_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QKeySequence) __keyBindings_setList(i QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QKeySequence___keyBindings_setList(ptr.Pointer(), PointerFromQKeySequence(i))
	}
}

func (ptr *QKeySequence) __keyBindings_newList() unsafe.Pointer {
	return C.QKeySequence___keyBindings_newList(ptr.Pointer())
}

func (ptr *QKeySequence) __listFromString_atList(i int) *QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := NewQKeySequenceFromPointer(C.QKeySequence___listFromString_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QKeySequence) __listFromString_setList(i QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QKeySequence___listFromString_setList(ptr.Pointer(), PointerFromQKeySequence(i))
	}
}

func (ptr *QKeySequence) __listFromString_newList() unsafe.Pointer {
	return C.QKeySequence___listFromString_newList(ptr.Pointer())
}

func (ptr *QKeySequence) __listToString_list_atList(i int) *QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := NewQKeySequenceFromPointer(C.QKeySequence___listToString_list_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QKeySequence) __listToString_list_setList(i QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QKeySequence___listToString_list_setList(ptr.Pointer(), PointerFromQKeySequence(i))
	}
}

func (ptr *QKeySequence) __listToString_list_newList() unsafe.Pointer {
	return C.QKeySequence___listToString_list_newList(ptr.Pointer())
}

type QMatrix4x4 struct {
	ptr unsafe.Pointer
}

type QMatrix4x4_ITF interface {
	QMatrix4x4_PTR() *QMatrix4x4
}

func (ptr *QMatrix4x4) QMatrix4x4_PTR() *QMatrix4x4 {
	return ptr
}

func (ptr *QMatrix4x4) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QMatrix4x4) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQMatrix4x4(ptr QMatrix4x4_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMatrix4x4_PTR().Pointer()
	}
	return nil
}

func NewQMatrix4x4FromPointer(ptr unsafe.Pointer) (n *QMatrix4x4) {
	n = new(QMatrix4x4)
	n.SetPointer(ptr)
	return
}
func (ptr *QMatrix4x4) DestroyQMatrix4x4() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQMatrix4x4() *QMatrix4x4 {
	tmpValue := NewQMatrix4x4FromPointer(C.QMatrix4x4_NewQMatrix4x4())
	qt.SetFinalizer(tmpValue, (*QMatrix4x4).DestroyQMatrix4x4)
	return tmpValue
}

func NewQMatrix4x43(values float32) *QMatrix4x4 {
	tmpValue := NewQMatrix4x4FromPointer(C.QMatrix4x4_NewQMatrix4x43(C.float(values)))
	qt.SetFinalizer(tmpValue, (*QMatrix4x4).DestroyQMatrix4x4)
	return tmpValue
}

func NewQMatrix4x44(m11 float32, m12 float32, m13 float32, m14 float32, m21 float32, m22 float32, m23 float32, m24 float32, m31 float32, m32 float32, m33 float32, m34 float32, m41 float32, m42 float32, m43 float32, m44 float32) *QMatrix4x4 {
	tmpValue := NewQMatrix4x4FromPointer(C.QMatrix4x4_NewQMatrix4x44(C.float(m11), C.float(m12), C.float(m13), C.float(m14), C.float(m21), C.float(m22), C.float(m23), C.float(m24), C.float(m31), C.float(m32), C.float(m33), C.float(m34), C.float(m41), C.float(m42), C.float(m43), C.float(m44)))
	qt.SetFinalizer(tmpValue, (*QMatrix4x4).DestroyQMatrix4x4)
	return tmpValue
}

func (ptr *QMatrix4x4) Column(index int) *QVector4D {
	if ptr.Pointer() != nil {
		tmpValue := NewQVector4DFromPointer(C.QMatrix4x4_Column(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix4x4) Data() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QMatrix4x4_Data(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMatrix4x4) Data2() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QMatrix4x4_Data2(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMatrix4x4) Map(point core.QPoint_ITF) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QMatrix4x4_Map(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix4x4) Map2(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QMatrix4x4_Map2(ptr.Pointer(), core.PointerFromQPointF(point)))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix4x4) Map3(point QVector3D_ITF) *QVector3D {
	if ptr.Pointer() != nil {
		tmpValue := NewQVector3DFromPointer(C.QMatrix4x4_Map3(ptr.Pointer(), PointerFromQVector3D(point)))
		qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix4x4) Map4(point QVector4D_ITF) *QVector4D {
	if ptr.Pointer() != nil {
		tmpValue := NewQVector4DFromPointer(C.QMatrix4x4_Map4(ptr.Pointer(), PointerFromQVector4D(point)))
		qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix4x4) Row(index int) *QVector4D {
	if ptr.Pointer() != nil {
		tmpValue := NewQVector4DFromPointer(C.QMatrix4x4_Row(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix4x4) Scale(vector QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Scale(ptr.Pointer(), PointerFromQVector3D(vector))
	}
}

func (ptr *QMatrix4x4) Scale2(x float32, y float32) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Scale2(ptr.Pointer(), C.float(x), C.float(y))
	}
}

func (ptr *QMatrix4x4) Scale3(x float32, y float32, z float32) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Scale3(ptr.Pointer(), C.float(x), C.float(y), C.float(z))
	}
}

func (ptr *QMatrix4x4) Scale4(factor float32) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Scale4(ptr.Pointer(), C.float(factor))
	}
}

func (ptr *QMatrix4x4) SetColumn(index int, value QVector4D_ITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_SetColumn(ptr.Pointer(), C.int(int32(index)), PointerFromQVector4D(value))
	}
}

type QMouseEvent struct {
	QInputEvent
}

type QMouseEvent_ITF interface {
	QInputEvent_ITF
	QMouseEvent_PTR() *QMouseEvent
}

func (ptr *QMouseEvent) QMouseEvent_PTR() *QMouseEvent {
	return ptr
}

func (ptr *QMouseEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QMouseEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QInputEvent_PTR().SetPointer(p)
	}
}

func PointerFromQMouseEvent(ptr QMouseEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMouseEvent_PTR().Pointer()
	}
	return nil
}

func NewQMouseEventFromPointer(ptr unsafe.Pointer) (n *QMouseEvent) {
	n = new(QMouseEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QMouseEvent) DestroyQMouseEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQMouseEvent(ty core.QEvent__Type, localPos core.QPointF_ITF, button core.Qt__MouseButton, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QMouseEvent {
	tmpValue := NewQMouseEventFromPointer(C.QMouseEvent_NewQMouseEvent(C.longlong(ty), core.PointerFromQPointF(localPos), C.longlong(button), C.longlong(buttons), C.longlong(modifiers)))
	qt.SetFinalizer(tmpValue, (*QMouseEvent).DestroyQMouseEvent)
	return tmpValue
}

func NewQMouseEvent2(ty core.QEvent__Type, localPos core.QPointF_ITF, screenPos core.QPointF_ITF, button core.Qt__MouseButton, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QMouseEvent {
	tmpValue := NewQMouseEventFromPointer(C.QMouseEvent_NewQMouseEvent2(C.longlong(ty), core.PointerFromQPointF(localPos), core.PointerFromQPointF(screenPos), C.longlong(button), C.longlong(buttons), C.longlong(modifiers)))
	qt.SetFinalizer(tmpValue, (*QMouseEvent).DestroyQMouseEvent)
	return tmpValue
}

func NewQMouseEvent3(ty core.QEvent__Type, localPos core.QPointF_ITF, windowPos core.QPointF_ITF, screenPos core.QPointF_ITF, button core.Qt__MouseButton, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QMouseEvent {
	tmpValue := NewQMouseEventFromPointer(C.QMouseEvent_NewQMouseEvent3(C.longlong(ty), core.PointerFromQPointF(localPos), core.PointerFromQPointF(windowPos), core.PointerFromQPointF(screenPos), C.longlong(button), C.longlong(buttons), C.longlong(modifiers)))
	qt.SetFinalizer(tmpValue, (*QMouseEvent).DestroyQMouseEvent)
	return tmpValue
}

func NewQMouseEvent4(ty core.QEvent__Type, localPos core.QPointF_ITF, windowPos core.QPointF_ITF, screenPos core.QPointF_ITF, button core.Qt__MouseButton, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, source core.Qt__MouseEventSource) *QMouseEvent {
	tmpValue := NewQMouseEventFromPointer(C.QMouseEvent_NewQMouseEvent4(C.longlong(ty), core.PointerFromQPointF(localPos), core.PointerFromQPointF(windowPos), core.PointerFromQPointF(screenPos), C.longlong(button), C.longlong(buttons), C.longlong(modifiers), C.longlong(source)))
	qt.SetFinalizer(tmpValue, (*QMouseEvent).DestroyQMouseEvent)
	return tmpValue
}

func (ptr *QMouseEvent) Button() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QMouseEvent_Button(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEvent) Buttons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QMouseEvent_Buttons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEvent) GlobalPos() *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QMouseEvent_GlobalPos(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QMouseEvent) Pos() *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QMouseEvent_Pos(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QMouseEvent) X() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QMouseEvent_X(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMouseEvent) Y() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QMouseEvent_Y(ptr.Pointer())))
	}
	return 0
}

// QMovie::CacheMode
//
//go:generate stringer -type=QMovie__CacheMode
type QMovie__CacheMode int64

const (
	QMovie__CacheNone QMovie__CacheMode = QMovie__CacheMode(0)
	QMovie__CacheAll  QMovie__CacheMode = QMovie__CacheMode(1)
)

// QMovie::MovieState
//
//go:generate stringer -type=QMovie__MovieState
type QMovie__MovieState int64

const (
	QMovie__NotRunning QMovie__MovieState = QMovie__MovieState(0)
	QMovie__Paused     QMovie__MovieState = QMovie__MovieState(1)
	QMovie__Running    QMovie__MovieState = QMovie__MovieState(2)
)

// QOpenGLBuffer::Access
//
//go:generate stringer -type=QOpenGLBuffer__Access
type QOpenGLBuffer__Access int64

const (
	QOpenGLBuffer__ReadOnly  QOpenGLBuffer__Access = QOpenGLBuffer__Access(0x88B8)
	QOpenGLBuffer__WriteOnly QOpenGLBuffer__Access = QOpenGLBuffer__Access(0x88B9)
	QOpenGLBuffer__ReadWrite QOpenGLBuffer__Access = QOpenGLBuffer__Access(0x88BA)
)

// QOpenGLBuffer::RangeAccessFlag
//
//go:generate stringer -type=QOpenGLBuffer__RangeAccessFlag
type QOpenGLBuffer__RangeAccessFlag int64

const (
	QOpenGLBuffer__RangeRead             QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0001)
	QOpenGLBuffer__RangeWrite            QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0002)
	QOpenGLBuffer__RangeInvalidate       QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0004)
	QOpenGLBuffer__RangeInvalidateBuffer QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0008)
	QOpenGLBuffer__RangeFlushExplicit    QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0010)
	QOpenGLBuffer__RangeUnsynchronized   QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0020)
)

// QOpenGLBuffer::Type
//
//go:generate stringer -type=QOpenGLBuffer__Type
type QOpenGLBuffer__Type int64

const (
	QOpenGLBuffer__VertexBuffer      QOpenGLBuffer__Type = QOpenGLBuffer__Type(0x8892)
	QOpenGLBuffer__IndexBuffer       QOpenGLBuffer__Type = QOpenGLBuffer__Type(0x8893)
	QOpenGLBuffer__PixelPackBuffer   QOpenGLBuffer__Type = QOpenGLBuffer__Type(0x88EB)
	QOpenGLBuffer__PixelUnpackBuffer QOpenGLBuffer__Type = QOpenGLBuffer__Type(0x88EC)
)

// QOpenGLBuffer::UsagePattern
//
//go:generate stringer -type=QOpenGLBuffer__UsagePattern
type QOpenGLBuffer__UsagePattern int64

const (
	QOpenGLBuffer__StreamDraw  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E0)
	QOpenGLBuffer__StreamRead  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E1)
	QOpenGLBuffer__StreamCopy  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E2)
	QOpenGLBuffer__StaticDraw  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E4)
	QOpenGLBuffer__StaticRead  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E5)
	QOpenGLBuffer__StaticCopy  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E6)
	QOpenGLBuffer__DynamicDraw QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E8)
	QOpenGLBuffer__DynamicRead QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E9)
	QOpenGLBuffer__DynamicCopy QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88EA)
)

// QOpenGLContext::OpenGLModuleType
//
//go:generate stringer -type=QOpenGLContext__OpenGLModuleType
type QOpenGLContext__OpenGLModuleType int64

const (
	QOpenGLContext__LibGL   QOpenGLContext__OpenGLModuleType = QOpenGLContext__OpenGLModuleType(0)
	QOpenGLContext__LibGLES QOpenGLContext__OpenGLModuleType = QOpenGLContext__OpenGLModuleType(1)
)

// QOpenGLDebugLogger::LoggingMode
//
//go:generate stringer -type=QOpenGLDebugLogger__LoggingMode
type QOpenGLDebugLogger__LoggingMode int64

const (
	QOpenGLDebugLogger__AsynchronousLogging QOpenGLDebugLogger__LoggingMode = QOpenGLDebugLogger__LoggingMode(0)
	QOpenGLDebugLogger__SynchronousLogging  QOpenGLDebugLogger__LoggingMode = QOpenGLDebugLogger__LoggingMode(1)
)

// QOpenGLDebugMessage::Severity
//
//go:generate stringer -type=QOpenGLDebugMessage__Severity
type QOpenGLDebugMessage__Severity int64

const (
	QOpenGLDebugMessage__InvalidSeverity      QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0x00000000)
	QOpenGLDebugMessage__HighSeverity         QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0x00000001)
	QOpenGLDebugMessage__MediumSeverity       QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0x00000002)
	QOpenGLDebugMessage__LowSeverity          QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0x00000004)
	QOpenGLDebugMessage__NotificationSeverity QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0x00000008)
	QOpenGLDebugMessage__LastSeverity         QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(QOpenGLDebugMessage__NotificationSeverity)
	QOpenGLDebugMessage__AnySeverity          QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0xffffffff)
)

// QOpenGLDebugMessage::Source
//
//go:generate stringer -type=QOpenGLDebugMessage__Source
type QOpenGLDebugMessage__Source int64

const (
	QOpenGLDebugMessage__InvalidSource        QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000000)
	QOpenGLDebugMessage__APISource            QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000001)
	QOpenGLDebugMessage__WindowSystemSource   QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000002)
	QOpenGLDebugMessage__ShaderCompilerSource QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000004)
	QOpenGLDebugMessage__ThirdPartySource     QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000008)
	QOpenGLDebugMessage__ApplicationSource    QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000010)
	QOpenGLDebugMessage__OtherSource          QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000020)
	QOpenGLDebugMessage__LastSource           QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(QOpenGLDebugMessage__OtherSource)
	QOpenGLDebugMessage__AnySource            QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0xffffffff)
)

// QOpenGLDebugMessage::Type
//
//go:generate stringer -type=QOpenGLDebugMessage__Type
type QOpenGLDebugMessage__Type int64

const (
	QOpenGLDebugMessage__InvalidType            QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000000)
	QOpenGLDebugMessage__ErrorType              QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000001)
	QOpenGLDebugMessage__DeprecatedBehaviorType QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000002)
	QOpenGLDebugMessage__UndefinedBehaviorType  QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000004)
	QOpenGLDebugMessage__PortabilityType        QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000008)
	QOpenGLDebugMessage__PerformanceType        QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000010)
	QOpenGLDebugMessage__OtherType              QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000020)
	QOpenGLDebugMessage__MarkerType             QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000040)
	QOpenGLDebugMessage__GroupPushType          QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000080)
	QOpenGLDebugMessage__GroupPopType           QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000100)
	QOpenGLDebugMessage__LastType               QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(QOpenGLDebugMessage__GroupPopType)
	QOpenGLDebugMessage__AnyType                QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0xffffffff)
)

// QOpenGLFramebufferObject::Attachment
//
//go:generate stringer -type=QOpenGLFramebufferObject__Attachment
type QOpenGLFramebufferObject__Attachment int64

const (
	QOpenGLFramebufferObject__NoAttachment         QOpenGLFramebufferObject__Attachment = QOpenGLFramebufferObject__Attachment(0)
	QOpenGLFramebufferObject__CombinedDepthStencil QOpenGLFramebufferObject__Attachment = QOpenGLFramebufferObject__Attachment(1)
	QOpenGLFramebufferObject__Depth                QOpenGLFramebufferObject__Attachment = QOpenGLFramebufferObject__Attachment(2)
)

// QOpenGLFramebufferObject::FramebufferRestorePolicy
//
//go:generate stringer -type=QOpenGLFramebufferObject__FramebufferRestorePolicy
type QOpenGLFramebufferObject__FramebufferRestorePolicy int64

const (
	QOpenGLFramebufferObject__DontRestoreFramebufferBinding      QOpenGLFramebufferObject__FramebufferRestorePolicy = QOpenGLFramebufferObject__FramebufferRestorePolicy(0)
	QOpenGLFramebufferObject__RestoreFramebufferBindingToDefault QOpenGLFramebufferObject__FramebufferRestorePolicy = QOpenGLFramebufferObject__FramebufferRestorePolicy(1)
	QOpenGLFramebufferObject__RestoreFrameBufferBinding          QOpenGLFramebufferObject__FramebufferRestorePolicy = QOpenGLFramebufferObject__FramebufferRestorePolicy(2)
)

// QOpenGLFunctions::OpenGLFeature
//
//go:generate stringer -type=QOpenGLFunctions__OpenGLFeature
type QOpenGLFunctions__OpenGLFeature int64

const (
	QOpenGLFunctions__Multitexture          QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0001)
	QOpenGLFunctions__Shaders               QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0002)
	QOpenGLFunctions__Buffers               QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0004)
	QOpenGLFunctions__Framebuffers          QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0008)
	QOpenGLFunctions__BlendColor            QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0010)
	QOpenGLFunctions__BlendEquation         QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0020)
	QOpenGLFunctions__BlendEquationSeparate QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0040)
	QOpenGLFunctions__BlendFuncSeparate     QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0080)
	QOpenGLFunctions__BlendSubtract         QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0100)
	QOpenGLFunctions__CompressedTextures    QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0200)
	QOpenGLFunctions__Multisample           QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0400)
	QOpenGLFunctions__StencilSeparate       QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0800)
	QOpenGLFunctions__NPOTTextures          QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x1000)
	QOpenGLFunctions__NPOTTextureRepeat     QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x2000)
	QOpenGLFunctions__FixedFunctionPipeline QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x4000)
	QOpenGLFunctions__TextureRGFormats      QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x8000)
	QOpenGLFunctions__MultipleRenderTargets QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x10000)
	QOpenGLFunctions__BlendEquationAdvanced QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x20000)
)

// QOpenGLShader::ShaderTypeBit
//
//go:generate stringer -type=QOpenGLShader__ShaderTypeBit
type QOpenGLShader__ShaderTypeBit int64

const (
	QOpenGLShader__Vertex                 QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0001)
	QOpenGLShader__Fragment               QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0002)
	QOpenGLShader__Geometry               QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0004)
	QOpenGLShader__TessellationControl    QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0008)
	QOpenGLShader__TessellationEvaluation QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0010)
	QOpenGLShader__Compute                QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0020)
)

// QOpenGLTexture::BindingTarget
//
//go:generate stringer -type=QOpenGLTexture__BindingTarget
type QOpenGLTexture__BindingTarget int64

const (
	QOpenGLTexture__BindingTarget1D                 QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8068)
	QOpenGLTexture__BindingTarget1DArray            QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8C1C)
	QOpenGLTexture__BindingTarget2D                 QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8069)
	QOpenGLTexture__BindingTarget2DArray            QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8C1D)
	QOpenGLTexture__BindingTarget3D                 QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x806A)
	QOpenGLTexture__BindingTargetCubeMap            QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8514)
	QOpenGLTexture__BindingTargetCubeMapArray       QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x900A)
	QOpenGLTexture__BindingTarget2DMultisample      QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x9104)
	QOpenGLTexture__BindingTarget2DMultisampleArray QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x9105)
	QOpenGLTexture__BindingTargetRectangle          QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x84F6)
	QOpenGLTexture__BindingTargetBuffer             QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8C2C)
)

// QOpenGLTexture::ComparisonFunction
//
//go:generate stringer -type=QOpenGLTexture__ComparisonFunction
type QOpenGLTexture__ComparisonFunction int64

const (
	QOpenGLTexture__CompareLessEqual    QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0203)
	QOpenGLTexture__CompareGreaterEqual QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0206)
	QOpenGLTexture__CompareLess         QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0201)
	QOpenGLTexture__CompareGreater      QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0204)
	QOpenGLTexture__CompareEqual        QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0202)
	QOpenGLTexture__CommpareNotEqual    QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0205)
	QOpenGLTexture__CompareAlways       QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0207)
	QOpenGLTexture__CompareNever        QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0200)
)

// QOpenGLTexture::ComparisonMode
//
//go:generate stringer -type=QOpenGLTexture__ComparisonMode
type QOpenGLTexture__ComparisonMode int64

const (
	QOpenGLTexture__CompareRefToTexture QOpenGLTexture__ComparisonMode = QOpenGLTexture__ComparisonMode(0x884E)
	QOpenGLTexture__CompareNone         QOpenGLTexture__ComparisonMode = QOpenGLTexture__ComparisonMode(0x0000)
)

// QOpenGLTexture::CoordinateDirection
//
//go:generate stringer -type=QOpenGLTexture__CoordinateDirection
type QOpenGLTexture__CoordinateDirection int64

const (
	QOpenGLTexture__DirectionS QOpenGLTexture__CoordinateDirection = QOpenGLTexture__CoordinateDirection(0x2802)
	QOpenGLTexture__DirectionT QOpenGLTexture__CoordinateDirection = QOpenGLTexture__CoordinateDirection(0x2803)
	QOpenGLTexture__DirectionR QOpenGLTexture__CoordinateDirection = QOpenGLTexture__CoordinateDirection(0x8072)
)

// QOpenGLTexture::CubeMapFace
//
//go:generate stringer -type=QOpenGLTexture__CubeMapFace
type QOpenGLTexture__CubeMapFace int64

const (
	QOpenGLTexture__CubeMapPositiveX QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x8515)
	QOpenGLTexture__CubeMapNegativeX QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x8516)
	QOpenGLTexture__CubeMapPositiveY QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x8517)
	QOpenGLTexture__CubeMapNegativeY QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x8518)
	QOpenGLTexture__CubeMapPositiveZ QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x8519)
	QOpenGLTexture__CubeMapNegativeZ QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x851A)
)

// QOpenGLTexture::DepthStencilMode
//
//go:generate stringer -type=QOpenGLTexture__DepthStencilMode
type QOpenGLTexture__DepthStencilMode int64

const (
	QOpenGLTexture__DepthMode   QOpenGLTexture__DepthStencilMode = QOpenGLTexture__DepthStencilMode(0x1902)
	QOpenGLTexture__StencilMode QOpenGLTexture__DepthStencilMode = QOpenGLTexture__DepthStencilMode(0x1901)
)

// QOpenGLTexture::Feature
//
//go:generate stringer -type=QOpenGLTexture__Feature
type QOpenGLTexture__Feature int64

const (
	QOpenGLTexture__ImmutableStorage            QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000001)
	QOpenGLTexture__ImmutableMultisampleStorage QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000002)
	QOpenGLTexture__TextureRectangle            QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000004)
	QOpenGLTexture__TextureArrays               QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000008)
	QOpenGLTexture__Texture3D                   QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000010)
	QOpenGLTexture__TextureMultisample          QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000020)
	QOpenGLTexture__TextureBuffer               QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000040)
	QOpenGLTexture__TextureCubeMapArrays        QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000080)
	QOpenGLTexture__Swizzle                     QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000100)
	QOpenGLTexture__StencilTexturing            QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000200)
	QOpenGLTexture__AnisotropicFiltering        QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000400)
	QOpenGLTexture__NPOTTextures                QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000800)
	QOpenGLTexture__NPOTTextureRepeat           QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00001000)
	QOpenGLTexture__Texture1D                   QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00002000)
	QOpenGLTexture__TextureComparisonOperators  QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00004000)
	QOpenGLTexture__TextureMipMapLevel          QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00008000)
)

// QOpenGLTexture::Filter
//
//go:generate stringer -type=QOpenGLTexture__Filter
type QOpenGLTexture__Filter int64

const (
	QOpenGLTexture__Nearest              QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2600)
	QOpenGLTexture__Linear               QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2601)
	QOpenGLTexture__NearestMipMapNearest QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2700)
	QOpenGLTexture__NearestMipMapLinear  QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2702)
	QOpenGLTexture__LinearMipMapNearest  QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2701)
	QOpenGLTexture__LinearMipMapLinear   QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2703)
)

// QOpenGLTexture::MipMapGeneration
//
//go:generate stringer -type=QOpenGLTexture__MipMapGeneration
type QOpenGLTexture__MipMapGeneration int64

const (
	QOpenGLTexture__GenerateMipMaps     QOpenGLTexture__MipMapGeneration = QOpenGLTexture__MipMapGeneration(0)
	QOpenGLTexture__DontGenerateMipMaps QOpenGLTexture__MipMapGeneration = QOpenGLTexture__MipMapGeneration(1)
)

// QOpenGLTexture::PixelFormat
//
//go:generate stringer -type=QOpenGLTexture__PixelFormat
type QOpenGLTexture__PixelFormat int64

const (
	QOpenGLTexture__NoSourceFormat QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0)
	QOpenGLTexture__Red            QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1903)
	QOpenGLTexture__RG             QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8227)
	QOpenGLTexture__RGB            QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1907)
	QOpenGLTexture__BGR            QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x80E0)
	QOpenGLTexture__RGBA           QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1908)
	QOpenGLTexture__BGRA           QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x80E1)
	QOpenGLTexture__Red_Integer    QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8D94)
	QOpenGLTexture__RG_Integer     QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8228)
	QOpenGLTexture__RGB_Integer    QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8D98)
	QOpenGLTexture__BGR_Integer    QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8D9A)
	QOpenGLTexture__RGBA_Integer   QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8D99)
	QOpenGLTexture__BGRA_Integer   QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8D9B)
	QOpenGLTexture__Stencil        QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1901)
	QOpenGLTexture__Depth          QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1902)
	QOpenGLTexture__DepthStencil   QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x84F9)
	QOpenGLTexture__Alpha          QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1906)
	QOpenGLTexture__Luminance      QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1909)
	QOpenGLTexture__LuminanceAlpha QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x190A)
)

// QOpenGLTexture::PixelType
//
//go:generate stringer -type=QOpenGLTexture__PixelType
type QOpenGLTexture__PixelType int64

const (
	QOpenGLTexture__NoPixelType               QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0)
	QOpenGLTexture__Int8                      QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1400)
	QOpenGLTexture__UInt8                     QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1401)
	QOpenGLTexture__Int16                     QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1402)
	QOpenGLTexture__UInt16                    QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1403)
	QOpenGLTexture__Int32                     QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1404)
	QOpenGLTexture__UInt32                    QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1405)
	QOpenGLTexture__Float16                   QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x140B)
	QOpenGLTexture__Float16OES                QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8D61)
	QOpenGLTexture__Float32                   QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1406)
	QOpenGLTexture__UInt32_RGB9_E5            QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8C3E)
	QOpenGLTexture__UInt32_RG11B10F           QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8C3B)
	QOpenGLTexture__UInt8_RG3B2               QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8032)
	QOpenGLTexture__UInt8_RG3B2_Rev           QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8362)
	QOpenGLTexture__UInt16_RGB5A1             QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8034)
	QOpenGLTexture__UInt16_RGB5A1_Rev         QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8366)
	QOpenGLTexture__UInt16_R5G6B5             QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8363)
	QOpenGLTexture__UInt16_R5G6B5_Rev         QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8364)
	QOpenGLTexture__UInt16_RGBA4              QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8033)
	QOpenGLTexture__UInt16_RGBA4_Rev          QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8365)
	QOpenGLTexture__UInt32_RGBA8              QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8035)
	QOpenGLTexture__UInt32_RGBA8_Rev          QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8367)
	QOpenGLTexture__UInt32_RGB10A2            QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8036)
	QOpenGLTexture__UInt32_RGB10A2_Rev        QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8368)
	QOpenGLTexture__UInt32_D24S8              QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x84FA)
	QOpenGLTexture__Float32_D32_UInt32_S8_X24 QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8DAD)
)

// QOpenGLTexture::SwizzleComponent
//
//go:generate stringer -type=QOpenGLTexture__SwizzleComponent
type QOpenGLTexture__SwizzleComponent int64

const (
	QOpenGLTexture__SwizzleRed   QOpenGLTexture__SwizzleComponent = QOpenGLTexture__SwizzleComponent(0x8E42)
	QOpenGLTexture__SwizzleGreen QOpenGLTexture__SwizzleComponent = QOpenGLTexture__SwizzleComponent(0x8E43)
	QOpenGLTexture__SwizzleBlue  QOpenGLTexture__SwizzleComponent = QOpenGLTexture__SwizzleComponent(0x8E44)
	QOpenGLTexture__SwizzleAlpha QOpenGLTexture__SwizzleComponent = QOpenGLTexture__SwizzleComponent(0x8E45)
)

// QOpenGLTexture::SwizzleValue
//
//go:generate stringer -type=QOpenGLTexture__SwizzleValue
type QOpenGLTexture__SwizzleValue int64

const (
	QOpenGLTexture__RedValue   QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(0x1903)
	QOpenGLTexture__GreenValue QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(0x1904)
	QOpenGLTexture__BlueValue  QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(0x1905)
	QOpenGLTexture__AlphaValue QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(0x1906)
	QOpenGLTexture__ZeroValue  QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(0)
	QOpenGLTexture__OneValue   QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(1)
)

// QOpenGLTexture::Target
//
//go:generate stringer -type=QOpenGLTexture__Target
type QOpenGLTexture__Target int64

const (
	QOpenGLTexture__Target1D                 QOpenGLTexture__Target = QOpenGLTexture__Target(0x0DE0)
	QOpenGLTexture__Target1DArray            QOpenGLTexture__Target = QOpenGLTexture__Target(0x8C18)
	QOpenGLTexture__Target2D                 QOpenGLTexture__Target = QOpenGLTexture__Target(0x0DE1)
	QOpenGLTexture__Target2DArray            QOpenGLTexture__Target = QOpenGLTexture__Target(0x8C1A)
	QOpenGLTexture__Target3D                 QOpenGLTexture__Target = QOpenGLTexture__Target(0x806F)
	QOpenGLTexture__TargetCubeMap            QOpenGLTexture__Target = QOpenGLTexture__Target(0x8513)
	QOpenGLTexture__TargetCubeMapArray       QOpenGLTexture__Target = QOpenGLTexture__Target(0x9009)
	QOpenGLTexture__Target2DMultisample      QOpenGLTexture__Target = QOpenGLTexture__Target(0x9100)
	QOpenGLTexture__Target2DMultisampleArray QOpenGLTexture__Target = QOpenGLTexture__Target(0x9102)
	QOpenGLTexture__TargetRectangle          QOpenGLTexture__Target = QOpenGLTexture__Target(0x84F5)
	QOpenGLTexture__TargetBuffer             QOpenGLTexture__Target = QOpenGLTexture__Target(0x8C2A)
)

// QOpenGLTexture::TextureFormat
//
//go:generate stringer -type=QOpenGLTexture__TextureFormat
type QOpenGLTexture__TextureFormat int64

const (
	QOpenGLTexture__NoFormat                       QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0)
	QOpenGLTexture__R8_UNorm                       QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8229)
	QOpenGLTexture__RG8_UNorm                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822B)
	QOpenGLTexture__RGB8_UNorm                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8051)
	QOpenGLTexture__RGBA8_UNorm                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8058)
	QOpenGLTexture__R16_UNorm                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822A)
	QOpenGLTexture__RG16_UNorm                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822C)
	QOpenGLTexture__RGB16_UNorm                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8054)
	QOpenGLTexture__RGBA16_UNorm                   QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x805B)
	QOpenGLTexture__R8_SNorm                       QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F94)
	QOpenGLTexture__RG8_SNorm                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F95)
	QOpenGLTexture__RGB8_SNorm                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F96)
	QOpenGLTexture__RGBA8_SNorm                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F97)
	QOpenGLTexture__R16_SNorm                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F98)
	QOpenGLTexture__RG16_SNorm                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F99)
	QOpenGLTexture__RGB16_SNorm                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F9A)
	QOpenGLTexture__RGBA16_SNorm                   QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F9B)
	QOpenGLTexture__R8U                            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8232)
	QOpenGLTexture__RG8U                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8238)
	QOpenGLTexture__RGB8U                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D7D)
	QOpenGLTexture__RGBA8U                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D7C)
	QOpenGLTexture__R16U                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8234)
	QOpenGLTexture__RG16U                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x823A)
	QOpenGLTexture__RGB16U                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D77)
	QOpenGLTexture__RGBA16U                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D76)
	QOpenGLTexture__R32U                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8236)
	QOpenGLTexture__RG32U                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x823C)
	QOpenGLTexture__RGB32U                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D71)
	QOpenGLTexture__RGBA32U                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D70)
	QOpenGLTexture__R8I                            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8231)
	QOpenGLTexture__RG8I                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8237)
	QOpenGLTexture__RGB8I                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D8F)
	QOpenGLTexture__RGBA8I                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D8E)
	QOpenGLTexture__R16I                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8233)
	QOpenGLTexture__RG16I                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8239)
	QOpenGLTexture__RGB16I                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D89)
	QOpenGLTexture__RGBA16I                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D88)
	QOpenGLTexture__R32I                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8235)
	QOpenGLTexture__RG32I                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x823B)
	QOpenGLTexture__RGB32I                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D83)
	QOpenGLTexture__RGBA32I                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D82)
	QOpenGLTexture__R16F                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822D)
	QOpenGLTexture__RG16F                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822F)
	QOpenGLTexture__RGB16F                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x881B)
	QOpenGLTexture__RGBA16F                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x881A)
	QOpenGLTexture__R32F                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822E)
	QOpenGLTexture__RG32F                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8230)
	QOpenGLTexture__RGB32F                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8815)
	QOpenGLTexture__RGBA32F                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8814)
	QOpenGLTexture__RGB9E5                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C3D)
	QOpenGLTexture__RG11B10F                       QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C3A)
	QOpenGLTexture__RG3B2                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x2A10)
	QOpenGLTexture__R5G6B5                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D62)
	QOpenGLTexture__RGB5A1                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8057)
	QOpenGLTexture__RGBA4                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8056)
	QOpenGLTexture__RGB10A2                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x906F)
	QOpenGLTexture__D16                            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x81A5)
	QOpenGLTexture__D24                            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x81A6)
	QOpenGLTexture__D24S8                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x88F0)
	QOpenGLTexture__D32                            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x81A7)
	QOpenGLTexture__D32F                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8CAC)
	QOpenGLTexture__D32FS8X24                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8CAD)
	QOpenGLTexture__S8                             QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D48)
	QOpenGLTexture__RGB_DXT1                       QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x83F0)
	QOpenGLTexture__RGBA_DXT1                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x83F1)
	QOpenGLTexture__RGBA_DXT3                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x83F2)
	QOpenGLTexture__RGBA_DXT5                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x83F3)
	QOpenGLTexture__R_ATI1N_UNorm                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8DBB)
	QOpenGLTexture__R_ATI1N_SNorm                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8DBC)
	QOpenGLTexture__RG_ATI2N_UNorm                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8DBD)
	QOpenGLTexture__RG_ATI2N_SNorm                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8DBE)
	QOpenGLTexture__RGB_BP_UNSIGNED_FLOAT          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8E8F)
	QOpenGLTexture__RGB_BP_SIGNED_FLOAT            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8E8E)
	QOpenGLTexture__RGB_BP_UNorm                   QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8E8C)
	QOpenGLTexture__R11_EAC_UNorm                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9270)
	QOpenGLTexture__R11_EAC_SNorm                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9271)
	QOpenGLTexture__RG11_EAC_UNorm                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9272)
	QOpenGLTexture__RG11_EAC_SNorm                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9273)
	QOpenGLTexture__RGB8_ETC2                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9274)
	QOpenGLTexture__SRGB8_ETC2                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9275)
	QOpenGLTexture__RGB8_PunchThrough_Alpha1_ETC2  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9276)
	QOpenGLTexture__SRGB8_PunchThrough_Alpha1_ETC2 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9277)
	QOpenGLTexture__RGBA8_ETC2_EAC                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9278)
	QOpenGLTexture__SRGB8_Alpha8_ETC2_EAC          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9279)
	QOpenGLTexture__RGB8_ETC1                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D64)
	QOpenGLTexture__RGBA_ASTC_4x4                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B0)
	QOpenGLTexture__RGBA_ASTC_5x4                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B1)
	QOpenGLTexture__RGBA_ASTC_5x5                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B2)
	QOpenGLTexture__RGBA_ASTC_6x5                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B3)
	QOpenGLTexture__RGBA_ASTC_6x6                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B4)
	QOpenGLTexture__RGBA_ASTC_8x5                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B5)
	QOpenGLTexture__RGBA_ASTC_8x6                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B6)
	QOpenGLTexture__RGBA_ASTC_8x8                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B7)
	QOpenGLTexture__RGBA_ASTC_10x5                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B8)
	QOpenGLTexture__RGBA_ASTC_10x6                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B9)
	QOpenGLTexture__RGBA_ASTC_10x8                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93BA)
	QOpenGLTexture__RGBA_ASTC_10x10                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93BB)
	QOpenGLTexture__RGBA_ASTC_12x10                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93BC)
	QOpenGLTexture__RGBA_ASTC_12x12                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93BD)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_4x4          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D0)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_5x4          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D1)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_5x5          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D2)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_6x5          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D3)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_6x6          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D4)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_8x5          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D5)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_8x6          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D6)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_8x8          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D7)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_10x5         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D8)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_10x6         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D9)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_10x8         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93DA)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_10x10        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93DB)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_12x10        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93DC)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_12x12        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93DD)
	QOpenGLTexture__SRGB8                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C41)
	QOpenGLTexture__SRGB8_Alpha8                   QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C43)
	QOpenGLTexture__SRGB_DXT1                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C4C)
	QOpenGLTexture__SRGB_Alpha_DXT1                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C4D)
	QOpenGLTexture__SRGB_Alpha_DXT3                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C4E)
	QOpenGLTexture__SRGB_Alpha_DXT5                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C4F)
	QOpenGLTexture__SRGB_BP_UNorm                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8E8D)
	QOpenGLTexture__DepthFormat                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x1902)
	QOpenGLTexture__AlphaFormat                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x1906)
	QOpenGLTexture__RGBFormat                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x1907)
	QOpenGLTexture__RGBAFormat                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x1908)
	QOpenGLTexture__LuminanceFormat                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x1909)
	QOpenGLTexture__LuminanceAlphaFormat           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x190A)
)

// QOpenGLTexture::TextureUnitReset
//
//go:generate stringer -type=QOpenGLTexture__TextureUnitReset
type QOpenGLTexture__TextureUnitReset int64

const (
	QOpenGLTexture__ResetTextureUnit     QOpenGLTexture__TextureUnitReset = QOpenGLTexture__TextureUnitReset(0)
	QOpenGLTexture__DontResetTextureUnit QOpenGLTexture__TextureUnitReset = QOpenGLTexture__TextureUnitReset(1)
)

// QOpenGLTexture::WrapMode
//
//go:generate stringer -type=QOpenGLTexture__WrapMode
type QOpenGLTexture__WrapMode int64

const (
	QOpenGLTexture__Repeat         QOpenGLTexture__WrapMode = QOpenGLTexture__WrapMode(0x2901)
	QOpenGLTexture__MirroredRepeat QOpenGLTexture__WrapMode = QOpenGLTexture__WrapMode(0x8370)
	QOpenGLTexture__ClampToEdge    QOpenGLTexture__WrapMode = QOpenGLTexture__WrapMode(0x812F)
	QOpenGLTexture__ClampToBorder  QOpenGLTexture__WrapMode = QOpenGLTexture__WrapMode(0x812D)
)

// QOpenGLTextureBlitter::Origin
//
//go:generate stringer -type=QOpenGLTextureBlitter__Origin
type QOpenGLTextureBlitter__Origin int64

const (
	QOpenGLTextureBlitter__OriginBottomLeft QOpenGLTextureBlitter__Origin = QOpenGLTextureBlitter__Origin(0)
	QOpenGLTextureBlitter__OriginTopLeft    QOpenGLTextureBlitter__Origin = QOpenGLTextureBlitter__Origin(1)
)

// QOpenGLWindow::UpdateBehavior
//
//go:generate stringer -type=QOpenGLWindow__UpdateBehavior
type QOpenGLWindow__UpdateBehavior int64

const (
	QOpenGLWindow__NoPartialUpdate    QOpenGLWindow__UpdateBehavior = QOpenGLWindow__UpdateBehavior(0)
	QOpenGLWindow__PartialUpdateBlit  QOpenGLWindow__UpdateBehavior = QOpenGLWindow__UpdateBehavior(1)
	QOpenGLWindow__PartialUpdateBlend QOpenGLWindow__UpdateBehavior = QOpenGLWindow__UpdateBehavior(2)
)

// QPageLayout::Mode
//
//go:generate stringer -type=QPageLayout__Mode
type QPageLayout__Mode int64

const (
	QPageLayout__StandardMode QPageLayout__Mode = QPageLayout__Mode(0)
	QPageLayout__FullPageMode QPageLayout__Mode = QPageLayout__Mode(1)
)

// QPageLayout::Orientation
//
//go:generate stringer -type=QPageLayout__Orientation
type QPageLayout__Orientation int64

const (
	QPageLayout__Portrait  QPageLayout__Orientation = QPageLayout__Orientation(0)
	QPageLayout__Landscape QPageLayout__Orientation = QPageLayout__Orientation(1)
)

// QPageLayout::Unit
//
//go:generate stringer -type=QPageLayout__Unit
type QPageLayout__Unit int64

const (
	QPageLayout__Millimeter QPageLayout__Unit = QPageLayout__Unit(0)
	QPageLayout__Point      QPageLayout__Unit = QPageLayout__Unit(1)
	QPageLayout__Inch       QPageLayout__Unit = QPageLayout__Unit(2)
	QPageLayout__Pica       QPageLayout__Unit = QPageLayout__Unit(3)
	QPageLayout__Didot      QPageLayout__Unit = QPageLayout__Unit(4)
	QPageLayout__Cicero     QPageLayout__Unit = QPageLayout__Unit(5)
)

type QPageSize struct {
	ptr unsafe.Pointer
}

type QPageSize_ITF interface {
	QPageSize_PTR() *QPageSize
}

func (ptr *QPageSize) QPageSize_PTR() *QPageSize {
	return ptr
}

func (ptr *QPageSize) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPageSize) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPageSize(ptr QPageSize_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPageSize_PTR().Pointer()
	}
	return nil
}

func NewQPageSizeFromPointer(ptr unsafe.Pointer) (n *QPageSize) {
	n = new(QPageSize)
	n.SetPointer(ptr)
	return
}

// QPageSize::PageSizeId
//
//go:generate stringer -type=QPageSize__PageSizeId
type QPageSize__PageSizeId int64

const (
	QPageSize__A4                 QPageSize__PageSizeId = QPageSize__PageSizeId(0)
	QPageSize__B5                 QPageSize__PageSizeId = QPageSize__PageSizeId(1)
	QPageSize__Letter             QPageSize__PageSizeId = QPageSize__PageSizeId(2)
	QPageSize__Legal              QPageSize__PageSizeId = QPageSize__PageSizeId(3)
	QPageSize__Executive          QPageSize__PageSizeId = QPageSize__PageSizeId(4)
	QPageSize__A0                 QPageSize__PageSizeId = QPageSize__PageSizeId(5)
	QPageSize__A1                 QPageSize__PageSizeId = QPageSize__PageSizeId(6)
	QPageSize__A2                 QPageSize__PageSizeId = QPageSize__PageSizeId(7)
	QPageSize__A3                 QPageSize__PageSizeId = QPageSize__PageSizeId(8)
	QPageSize__A5                 QPageSize__PageSizeId = QPageSize__PageSizeId(9)
	QPageSize__A6                 QPageSize__PageSizeId = QPageSize__PageSizeId(10)
	QPageSize__A7                 QPageSize__PageSizeId = QPageSize__PageSizeId(11)
	QPageSize__A8                 QPageSize__PageSizeId = QPageSize__PageSizeId(12)
	QPageSize__A9                 QPageSize__PageSizeId = QPageSize__PageSizeId(13)
	QPageSize__B0                 QPageSize__PageSizeId = QPageSize__PageSizeId(14)
	QPageSize__B1                 QPageSize__PageSizeId = QPageSize__PageSizeId(15)
	QPageSize__B10                QPageSize__PageSizeId = QPageSize__PageSizeId(16)
	QPageSize__B2                 QPageSize__PageSizeId = QPageSize__PageSizeId(17)
	QPageSize__B3                 QPageSize__PageSizeId = QPageSize__PageSizeId(18)
	QPageSize__B4                 QPageSize__PageSizeId = QPageSize__PageSizeId(19)
	QPageSize__B6                 QPageSize__PageSizeId = QPageSize__PageSizeId(20)
	QPageSize__B7                 QPageSize__PageSizeId = QPageSize__PageSizeId(21)
	QPageSize__B8                 QPageSize__PageSizeId = QPageSize__PageSizeId(22)
	QPageSize__B9                 QPageSize__PageSizeId = QPageSize__PageSizeId(23)
	QPageSize__C5E                QPageSize__PageSizeId = QPageSize__PageSizeId(24)
	QPageSize__Comm10E            QPageSize__PageSizeId = QPageSize__PageSizeId(25)
	QPageSize__DLE                QPageSize__PageSizeId = QPageSize__PageSizeId(26)
	QPageSize__Folio              QPageSize__PageSizeId = QPageSize__PageSizeId(27)
	QPageSize__Ledger             QPageSize__PageSizeId = QPageSize__PageSizeId(28)
	QPageSize__Tabloid            QPageSize__PageSizeId = QPageSize__PageSizeId(29)
	QPageSize__Custom             QPageSize__PageSizeId = QPageSize__PageSizeId(30)
	QPageSize__A10                QPageSize__PageSizeId = QPageSize__PageSizeId(31)
	QPageSize__A3Extra            QPageSize__PageSizeId = QPageSize__PageSizeId(32)
	QPageSize__A4Extra            QPageSize__PageSizeId = QPageSize__PageSizeId(33)
	QPageSize__A4Plus             QPageSize__PageSizeId = QPageSize__PageSizeId(34)
	QPageSize__A4Small            QPageSize__PageSizeId = QPageSize__PageSizeId(35)
	QPageSize__A5Extra            QPageSize__PageSizeId = QPageSize__PageSizeId(36)
	QPageSize__B5Extra            QPageSize__PageSizeId = QPageSize__PageSizeId(37)
	QPageSize__JisB0              QPageSize__PageSizeId = QPageSize__PageSizeId(38)
	QPageSize__JisB1              QPageSize__PageSizeId = QPageSize__PageSizeId(39)
	QPageSize__JisB2              QPageSize__PageSizeId = QPageSize__PageSizeId(40)
	QPageSize__JisB3              QPageSize__PageSizeId = QPageSize__PageSizeId(41)
	QPageSize__JisB4              QPageSize__PageSizeId = QPageSize__PageSizeId(42)
	QPageSize__JisB5              QPageSize__PageSizeId = QPageSize__PageSizeId(43)
	QPageSize__JisB6              QPageSize__PageSizeId = QPageSize__PageSizeId(44)
	QPageSize__JisB7              QPageSize__PageSizeId = QPageSize__PageSizeId(45)
	QPageSize__JisB8              QPageSize__PageSizeId = QPageSize__PageSizeId(46)
	QPageSize__JisB9              QPageSize__PageSizeId = QPageSize__PageSizeId(47)
	QPageSize__JisB10             QPageSize__PageSizeId = QPageSize__PageSizeId(48)
	QPageSize__AnsiC              QPageSize__PageSizeId = QPageSize__PageSizeId(49)
	QPageSize__AnsiD              QPageSize__PageSizeId = QPageSize__PageSizeId(50)
	QPageSize__AnsiE              QPageSize__PageSizeId = QPageSize__PageSizeId(51)
	QPageSize__LegalExtra         QPageSize__PageSizeId = QPageSize__PageSizeId(52)
	QPageSize__LetterExtra        QPageSize__PageSizeId = QPageSize__PageSizeId(53)
	QPageSize__LetterPlus         QPageSize__PageSizeId = QPageSize__PageSizeId(54)
	QPageSize__LetterSmall        QPageSize__PageSizeId = QPageSize__PageSizeId(55)
	QPageSize__TabloidExtra       QPageSize__PageSizeId = QPageSize__PageSizeId(56)
	QPageSize__ArchA              QPageSize__PageSizeId = QPageSize__PageSizeId(57)
	QPageSize__ArchB              QPageSize__PageSizeId = QPageSize__PageSizeId(58)
	QPageSize__ArchC              QPageSize__PageSizeId = QPageSize__PageSizeId(59)
	QPageSize__ArchD              QPageSize__PageSizeId = QPageSize__PageSizeId(60)
	QPageSize__ArchE              QPageSize__PageSizeId = QPageSize__PageSizeId(61)
	QPageSize__Imperial7x9        QPageSize__PageSizeId = QPageSize__PageSizeId(62)
	QPageSize__Imperial8x10       QPageSize__PageSizeId = QPageSize__PageSizeId(63)
	QPageSize__Imperial9x11       QPageSize__PageSizeId = QPageSize__PageSizeId(64)
	QPageSize__Imperial9x12       QPageSize__PageSizeId = QPageSize__PageSizeId(65)
	QPageSize__Imperial10x11      QPageSize__PageSizeId = QPageSize__PageSizeId(66)
	QPageSize__Imperial10x13      QPageSize__PageSizeId = QPageSize__PageSizeId(67)
	QPageSize__Imperial10x14      QPageSize__PageSizeId = QPageSize__PageSizeId(68)
	QPageSize__Imperial12x11      QPageSize__PageSizeId = QPageSize__PageSizeId(69)
	QPageSize__Imperial15x11      QPageSize__PageSizeId = QPageSize__PageSizeId(70)
	QPageSize__ExecutiveStandard  QPageSize__PageSizeId = QPageSize__PageSizeId(71)
	QPageSize__Note               QPageSize__PageSizeId = QPageSize__PageSizeId(72)
	QPageSize__Quarto             QPageSize__PageSizeId = QPageSize__PageSizeId(73)
	QPageSize__Statement          QPageSize__PageSizeId = QPageSize__PageSizeId(74)
	QPageSize__SuperA             QPageSize__PageSizeId = QPageSize__PageSizeId(75)
	QPageSize__SuperB             QPageSize__PageSizeId = QPageSize__PageSizeId(76)
	QPageSize__Postcard           QPageSize__PageSizeId = QPageSize__PageSizeId(77)
	QPageSize__DoublePostcard     QPageSize__PageSizeId = QPageSize__PageSizeId(78)
	QPageSize__Prc16K             QPageSize__PageSizeId = QPageSize__PageSizeId(79)
	QPageSize__Prc32K             QPageSize__PageSizeId = QPageSize__PageSizeId(80)
	QPageSize__Prc32KBig          QPageSize__PageSizeId = QPageSize__PageSizeId(81)
	QPageSize__FanFoldUS          QPageSize__PageSizeId = QPageSize__PageSizeId(82)
	QPageSize__FanFoldGerman      QPageSize__PageSizeId = QPageSize__PageSizeId(83)
	QPageSize__FanFoldGermanLegal QPageSize__PageSizeId = QPageSize__PageSizeId(84)
	QPageSize__EnvelopeB4         QPageSize__PageSizeId = QPageSize__PageSizeId(85)
	QPageSize__EnvelopeB5         QPageSize__PageSizeId = QPageSize__PageSizeId(86)
	QPageSize__EnvelopeB6         QPageSize__PageSizeId = QPageSize__PageSizeId(87)
	QPageSize__EnvelopeC0         QPageSize__PageSizeId = QPageSize__PageSizeId(88)
	QPageSize__EnvelopeC1         QPageSize__PageSizeId = QPageSize__PageSizeId(89)
	QPageSize__EnvelopeC2         QPageSize__PageSizeId = QPageSize__PageSizeId(90)
	QPageSize__EnvelopeC3         QPageSize__PageSizeId = QPageSize__PageSizeId(91)
	QPageSize__EnvelopeC4         QPageSize__PageSizeId = QPageSize__PageSizeId(92)
	QPageSize__EnvelopeC6         QPageSize__PageSizeId = QPageSize__PageSizeId(93)
	QPageSize__EnvelopeC65        QPageSize__PageSizeId = QPageSize__PageSizeId(94)
	QPageSize__EnvelopeC7         QPageSize__PageSizeId = QPageSize__PageSizeId(95)
	QPageSize__Envelope9          QPageSize__PageSizeId = QPageSize__PageSizeId(96)
	QPageSize__Envelope11         QPageSize__PageSizeId = QPageSize__PageSizeId(97)
	QPageSize__Envelope12         QPageSize__PageSizeId = QPageSize__PageSizeId(98)
	QPageSize__Envelope14         QPageSize__PageSizeId = QPageSize__PageSizeId(99)
	QPageSize__EnvelopeMonarch    QPageSize__PageSizeId = QPageSize__PageSizeId(100)
	QPageSize__EnvelopePersonal   QPageSize__PageSizeId = QPageSize__PageSizeId(101)
	QPageSize__EnvelopeChou3      QPageSize__PageSizeId = QPageSize__PageSizeId(102)
	QPageSize__EnvelopeChou4      QPageSize__PageSizeId = QPageSize__PageSizeId(103)
	QPageSize__EnvelopeInvite     QPageSize__PageSizeId = QPageSize__PageSizeId(104)
	QPageSize__EnvelopeItalian    QPageSize__PageSizeId = QPageSize__PageSizeId(105)
	QPageSize__EnvelopeKaku2      QPageSize__PageSizeId = QPageSize__PageSizeId(106)
	QPageSize__EnvelopeKaku3      QPageSize__PageSizeId = QPageSize__PageSizeId(107)
	QPageSize__EnvelopePrc1       QPageSize__PageSizeId = QPageSize__PageSizeId(108)
	QPageSize__EnvelopePrc2       QPageSize__PageSizeId = QPageSize__PageSizeId(109)
	QPageSize__EnvelopePrc3       QPageSize__PageSizeId = QPageSize__PageSizeId(110)
	QPageSize__EnvelopePrc4       QPageSize__PageSizeId = QPageSize__PageSizeId(111)
	QPageSize__EnvelopePrc5       QPageSize__PageSizeId = QPageSize__PageSizeId(112)
	QPageSize__EnvelopePrc6       QPageSize__PageSizeId = QPageSize__PageSizeId(113)
	QPageSize__EnvelopePrc7       QPageSize__PageSizeId = QPageSize__PageSizeId(114)
	QPageSize__EnvelopePrc8       QPageSize__PageSizeId = QPageSize__PageSizeId(115)
	QPageSize__EnvelopePrc9       QPageSize__PageSizeId = QPageSize__PageSizeId(116)
	QPageSize__EnvelopePrc10      QPageSize__PageSizeId = QPageSize__PageSizeId(117)
	QPageSize__EnvelopeYou4       QPageSize__PageSizeId = QPageSize__PageSizeId(118)
	QPageSize__LastPageSize       QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__EnvelopeYou4)
	QPageSize__NPageSize          QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__LastPageSize)
	QPageSize__NPaperSize         QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__LastPageSize)
	QPageSize__AnsiA              QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__Letter)
	QPageSize__AnsiB              QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__Ledger)
	QPageSize__EnvelopeC5         QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__C5E)
	QPageSize__EnvelopeDL         QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__DLE)
	QPageSize__Envelope10         QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__Comm10E)
)

// QPageSize::SizeMatchPolicy
//
//go:generate stringer -type=QPageSize__SizeMatchPolicy
type QPageSize__SizeMatchPolicy int64

const (
	QPageSize__FuzzyMatch            QPageSize__SizeMatchPolicy = QPageSize__SizeMatchPolicy(0)
	QPageSize__FuzzyOrientationMatch QPageSize__SizeMatchPolicy = QPageSize__SizeMatchPolicy(1)
	QPageSize__ExactMatch            QPageSize__SizeMatchPolicy = QPageSize__SizeMatchPolicy(2)
)

// QPageSize::Unit
//
//go:generate stringer -type=QPageSize__Unit
type QPageSize__Unit int64

const (
	QPageSize__Millimeter QPageSize__Unit = QPageSize__Unit(0)
	QPageSize__Point      QPageSize__Unit = QPageSize__Unit(1)
	QPageSize__Inch       QPageSize__Unit = QPageSize__Unit(2)
	QPageSize__Pica       QPageSize__Unit = QPageSize__Unit(3)
	QPageSize__Didot      QPageSize__Unit = QPageSize__Unit(4)
	QPageSize__Cicero     QPageSize__Unit = QPageSize__Unit(5)
)

func NewQPageSize() *QPageSize {
	tmpValue := NewQPageSizeFromPointer(C.QPageSize_NewQPageSize())
	qt.SetFinalizer(tmpValue, (*QPageSize).DestroyQPageSize)
	return tmpValue
}

func NewQPageSize2(pageSize QPageSize__PageSizeId) *QPageSize {
	tmpValue := NewQPageSizeFromPointer(C.QPageSize_NewQPageSize2(C.longlong(pageSize)))
	qt.SetFinalizer(tmpValue, (*QPageSize).DestroyQPageSize)
	return tmpValue
}

func NewQPageSize3(pointSize core.QSize_ITF, name string, matchPolicy QPageSize__SizeMatchPolicy) *QPageSize {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQPageSizeFromPointer(C.QPageSize_NewQPageSize3(core.PointerFromQSize(pointSize), C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(matchPolicy)))
	qt.SetFinalizer(tmpValue, (*QPageSize).DestroyQPageSize)
	return tmpValue
}

func NewQPageSize4(size core.QSizeF_ITF, units QPageSize__Unit, name string, matchPolicy QPageSize__SizeMatchPolicy) *QPageSize {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQPageSizeFromPointer(C.QPageSize_NewQPageSize4(core.PointerFromQSizeF(size), C.longlong(units), C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(matchPolicy)))
	qt.SetFinalizer(tmpValue, (*QPageSize).DestroyQPageSize)
	return tmpValue
}

func NewQPageSize5(other QPageSize_ITF) *QPageSize {
	tmpValue := NewQPageSizeFromPointer(C.QPageSize_NewQPageSize5(PointerFromQPageSize(other)))
	qt.SetFinalizer(tmpValue, (*QPageSize).DestroyQPageSize)
	return tmpValue
}

func (ptr *QPageSize) Id() QPageSize__PageSizeId {
	if ptr.Pointer() != nil {
		return QPageSize__PageSizeId(C.QPageSize_Id(ptr.Pointer()))
	}
	return 0
}

func QPageSize_Id2(pointSize core.QSize_ITF, matchPolicy QPageSize__SizeMatchPolicy) QPageSize__PageSizeId {
	return QPageSize__PageSizeId(C.QPageSize_QPageSize_Id2(core.PointerFromQSize(pointSize), C.longlong(matchPolicy)))
}

func (ptr *QPageSize) Id2(pointSize core.QSize_ITF, matchPolicy QPageSize__SizeMatchPolicy) QPageSize__PageSizeId {
	return QPageSize__PageSizeId(C.QPageSize_QPageSize_Id2(core.PointerFromQSize(pointSize), C.longlong(matchPolicy)))
}

func QPageSize_Id3(size core.QSizeF_ITF, units QPageSize__Unit, matchPolicy QPageSize__SizeMatchPolicy) QPageSize__PageSizeId {
	return QPageSize__PageSizeId(C.QPageSize_QPageSize_Id3(core.PointerFromQSizeF(size), C.longlong(units), C.longlong(matchPolicy)))
}

func (ptr *QPageSize) Id3(size core.QSizeF_ITF, units QPageSize__Unit, matchPolicy QPageSize__SizeMatchPolicy) QPageSize__PageSizeId {
	return QPageSize__PageSizeId(C.QPageSize_QPageSize_Id3(core.PointerFromQSizeF(size), C.longlong(units), C.longlong(matchPolicy)))
}

func QPageSize_Id4(windowsId int) QPageSize__PageSizeId {
	return QPageSize__PageSizeId(C.QPageSize_QPageSize_Id4(C.int(int32(windowsId))))
}

func (ptr *QPageSize) Id4(windowsId int) QPageSize__PageSizeId {
	return QPageSize__PageSizeId(C.QPageSize_QPageSize_Id4(C.int(int32(windowsId))))
}

func (ptr *QPageSize) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPageSize_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPageSize) Key() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPageSize_Key(ptr.Pointer()))
	}
	return ""
}

func QPageSize_Key2(pageSizeId QPageSize__PageSizeId) string {
	return cGoUnpackString(C.QPageSize_QPageSize_Key2(C.longlong(pageSizeId)))
}

func (ptr *QPageSize) Key2(pageSizeId QPageSize__PageSizeId) string {
	return cGoUnpackString(C.QPageSize_QPageSize_Key2(C.longlong(pageSizeId)))
}

func (ptr *QPageSize) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPageSize_Name(ptr.Pointer()))
	}
	return ""
}

func QPageSize_Name2(pageSizeId QPageSize__PageSizeId) string {
	return cGoUnpackString(C.QPageSize_QPageSize_Name2(C.longlong(pageSizeId)))
}

func (ptr *QPageSize) Name2(pageSizeId QPageSize__PageSizeId) string {
	return cGoUnpackString(C.QPageSize_QPageSize_Name2(C.longlong(pageSizeId)))
}

func (ptr *QPageSize) Rect(units QPageSize__Unit) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QPageSize_Rect(ptr.Pointer(), C.longlong(units)))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPageSize) Size(units QPageSize__Unit) *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QPageSize_Size(ptr.Pointer(), C.longlong(units)))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func QPageSize_Size2(pageSizeId QPageSize__PageSizeId, units QPageSize__Unit) *core.QSizeF {
	tmpValue := core.NewQSizeFFromPointer(C.QPageSize_QPageSize_Size2(C.longlong(pageSizeId), C.longlong(units)))
	qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
	return tmpValue
}

func (ptr *QPageSize) Size2(pageSizeId QPageSize__PageSizeId, units QPageSize__Unit) *core.QSizeF {
	tmpValue := core.NewQSizeFFromPointer(C.QPageSize_QPageSize_Size2(C.longlong(pageSizeId), C.longlong(units)))
	qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
	return tmpValue
}

func (ptr *QPageSize) DestroyQPageSize() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPageSize_DestroyQPageSize(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QPagedPaintDevice struct {
	QPaintDevice
}

type QPagedPaintDevice_ITF interface {
	QPaintDevice_ITF
	QPagedPaintDevice_PTR() *QPagedPaintDevice
}

func (ptr *QPagedPaintDevice) QPagedPaintDevice_PTR() *QPagedPaintDevice {
	return ptr
}

func (ptr *QPagedPaintDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QPagedPaintDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQPagedPaintDevice(ptr QPagedPaintDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPagedPaintDevice_PTR().Pointer()
	}
	return nil
}

func NewQPagedPaintDeviceFromPointer(ptr unsafe.Pointer) (n *QPagedPaintDevice) {
	n = new(QPagedPaintDevice)
	n.SetPointer(ptr)
	return
}

// QPagedPaintDevice::PageSize
//
//go:generate stringer -type=QPagedPaintDevice__PageSize
type QPagedPaintDevice__PageSize int64

const (
	QPagedPaintDevice__A4                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(0)
	QPagedPaintDevice__B5                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(1)
	QPagedPaintDevice__Letter             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(2)
	QPagedPaintDevice__Legal              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(3)
	QPagedPaintDevice__Executive          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(4)
	QPagedPaintDevice__A0                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(5)
	QPagedPaintDevice__A1                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(6)
	QPagedPaintDevice__A2                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(7)
	QPagedPaintDevice__A3                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(8)
	QPagedPaintDevice__A5                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(9)
	QPagedPaintDevice__A6                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(10)
	QPagedPaintDevice__A7                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(11)
	QPagedPaintDevice__A8                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(12)
	QPagedPaintDevice__A9                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(13)
	QPagedPaintDevice__B0                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(14)
	QPagedPaintDevice__B1                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(15)
	QPagedPaintDevice__B10                QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(16)
	QPagedPaintDevice__B2                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(17)
	QPagedPaintDevice__B3                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(18)
	QPagedPaintDevice__B4                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(19)
	QPagedPaintDevice__B6                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(20)
	QPagedPaintDevice__B7                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(21)
	QPagedPaintDevice__B8                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(22)
	QPagedPaintDevice__B9                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(23)
	QPagedPaintDevice__C5E                QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(24)
	QPagedPaintDevice__Comm10E            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(25)
	QPagedPaintDevice__DLE                QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(26)
	QPagedPaintDevice__Folio              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(27)
	QPagedPaintDevice__Ledger             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(28)
	QPagedPaintDevice__Tabloid            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(29)
	QPagedPaintDevice__Custom             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(30)
	QPagedPaintDevice__A10                QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(31)
	QPagedPaintDevice__A3Extra            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(32)
	QPagedPaintDevice__A4Extra            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(33)
	QPagedPaintDevice__A4Plus             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(34)
	QPagedPaintDevice__A4Small            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(35)
	QPagedPaintDevice__A5Extra            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(36)
	QPagedPaintDevice__B5Extra            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(37)
	QPagedPaintDevice__JisB0              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(38)
	QPagedPaintDevice__JisB1              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(39)
	QPagedPaintDevice__JisB2              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(40)
	QPagedPaintDevice__JisB3              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(41)
	QPagedPaintDevice__JisB4              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(42)
	QPagedPaintDevice__JisB5              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(43)
	QPagedPaintDevice__JisB6              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(44)
	QPagedPaintDevice__JisB7              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(45)
	QPagedPaintDevice__JisB8              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(46)
	QPagedPaintDevice__JisB9              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(47)
	QPagedPaintDevice__JisB10             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(48)
	QPagedPaintDevice__AnsiC              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(49)
	QPagedPaintDevice__AnsiD              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(50)
	QPagedPaintDevice__AnsiE              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(51)
	QPagedPaintDevice__LegalExtra         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(52)
	QPagedPaintDevice__LetterExtra        QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(53)
	QPagedPaintDevice__LetterPlus         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(54)
	QPagedPaintDevice__LetterSmall        QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(55)
	QPagedPaintDevice__TabloidExtra       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(56)
	QPagedPaintDevice__ArchA              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(57)
	QPagedPaintDevice__ArchB              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(58)
	QPagedPaintDevice__ArchC              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(59)
	QPagedPaintDevice__ArchD              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(60)
	QPagedPaintDevice__ArchE              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(61)
	QPagedPaintDevice__Imperial7x9        QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(62)
	QPagedPaintDevice__Imperial8x10       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(63)
	QPagedPaintDevice__Imperial9x11       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(64)
	QPagedPaintDevice__Imperial9x12       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(65)
	QPagedPaintDevice__Imperial10x11      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(66)
	QPagedPaintDevice__Imperial10x13      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(67)
	QPagedPaintDevice__Imperial10x14      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(68)
	QPagedPaintDevice__Imperial12x11      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(69)
	QPagedPaintDevice__Imperial15x11      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(70)
	QPagedPaintDevice__ExecutiveStandard  QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(71)
	QPagedPaintDevice__Note               QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(72)
	QPagedPaintDevice__Quarto             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(73)
	QPagedPaintDevice__Statement          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(74)
	QPagedPaintDevice__SuperA             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(75)
	QPagedPaintDevice__SuperB             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(76)
	QPagedPaintDevice__Postcard           QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(77)
	QPagedPaintDevice__DoublePostcard     QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(78)
	QPagedPaintDevice__Prc16K             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(79)
	QPagedPaintDevice__Prc32K             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(80)
	QPagedPaintDevice__Prc32KBig          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(81)
	QPagedPaintDevice__FanFoldUS          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(82)
	QPagedPaintDevice__FanFoldGerman      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(83)
	QPagedPaintDevice__FanFoldGermanLegal QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(84)
	QPagedPaintDevice__EnvelopeB4         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(85)
	QPagedPaintDevice__EnvelopeB5         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(86)
	QPagedPaintDevice__EnvelopeB6         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(87)
	QPagedPaintDevice__EnvelopeC0         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(88)
	QPagedPaintDevice__EnvelopeC1         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(89)
	QPagedPaintDevice__EnvelopeC2         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(90)
	QPagedPaintDevice__EnvelopeC3         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(91)
	QPagedPaintDevice__EnvelopeC4         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(92)
	QPagedPaintDevice__EnvelopeC6         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(93)
	QPagedPaintDevice__EnvelopeC65        QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(94)
	QPagedPaintDevice__EnvelopeC7         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(95)
	QPagedPaintDevice__Envelope9          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(96)
	QPagedPaintDevice__Envelope11         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(97)
	QPagedPaintDevice__Envelope12         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(98)
	QPagedPaintDevice__Envelope14         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(99)
	QPagedPaintDevice__EnvelopeMonarch    QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(100)
	QPagedPaintDevice__EnvelopePersonal   QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(101)
	QPagedPaintDevice__EnvelopeChou3      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(102)
	QPagedPaintDevice__EnvelopeChou4      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(103)
	QPagedPaintDevice__EnvelopeInvite     QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(104)
	QPagedPaintDevice__EnvelopeItalian    QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(105)
	QPagedPaintDevice__EnvelopeKaku2      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(106)
	QPagedPaintDevice__EnvelopeKaku3      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(107)
	QPagedPaintDevice__EnvelopePrc1       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(108)
	QPagedPaintDevice__EnvelopePrc2       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(109)
	QPagedPaintDevice__EnvelopePrc3       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(110)
	QPagedPaintDevice__EnvelopePrc4       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(111)
	QPagedPaintDevice__EnvelopePrc5       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(112)
	QPagedPaintDevice__EnvelopePrc6       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(113)
	QPagedPaintDevice__EnvelopePrc7       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(114)
	QPagedPaintDevice__EnvelopePrc8       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(115)
	QPagedPaintDevice__EnvelopePrc9       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(116)
	QPagedPaintDevice__EnvelopePrc10      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(117)
	QPagedPaintDevice__EnvelopeYou4       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(118)
	QPagedPaintDevice__LastPageSize       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__EnvelopeYou4)
	QPagedPaintDevice__NPageSize          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__LastPageSize)
	QPagedPaintDevice__NPaperSize         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__LastPageSize)
	QPagedPaintDevice__AnsiA              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__Letter)
	QPagedPaintDevice__AnsiB              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__Ledger)
	QPagedPaintDevice__EnvelopeC5         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__C5E)
	QPagedPaintDevice__EnvelopeDL         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__DLE)
	QPagedPaintDevice__Envelope10         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__Comm10E)
)

// QPagedPaintDevice::PdfVersion
//
//go:generate stringer -type=QPagedPaintDevice__PdfVersion
type QPagedPaintDevice__PdfVersion int64

const (
	QPagedPaintDevice__PdfVersion_1_4 QPagedPaintDevice__PdfVersion = QPagedPaintDevice__PdfVersion(0)
	QPagedPaintDevice__PdfVersion_A1b QPagedPaintDevice__PdfVersion = QPagedPaintDevice__PdfVersion(1)
	QPagedPaintDevice__PdfVersion_1_6 QPagedPaintDevice__PdfVersion = QPagedPaintDevice__PdfVersion(2)
)

//export callbackQPagedPaintDevice_NewPage
func callbackQPagedPaintDevice_NewPage(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "newPage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPagedPaintDevice) ConnectNewPage(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "newPage"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "newPage", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newPage", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPagedPaintDevice) DisconnectNewPage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "newPage")
	}
}

func (ptr *QPagedPaintDevice) NewPage() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPagedPaintDevice_NewPage(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPagedPaintDevice_DestroyQPagedPaintDevice
func callbackQPagedPaintDevice_DestroyQPagedPaintDevice(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPagedPaintDevice"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPagedPaintDeviceFromPointer(ptr).DestroyQPagedPaintDeviceDefault()
	}
}

func (ptr *QPagedPaintDevice) ConnectDestroyQPagedPaintDevice(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPagedPaintDevice"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPagedPaintDevice", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPagedPaintDevice", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPagedPaintDevice) DisconnectDestroyQPagedPaintDevice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPagedPaintDevice")
	}
}

func (ptr *QPagedPaintDevice) DestroyQPagedPaintDevice() {
	if ptr.Pointer() != nil {
		C.QPagedPaintDevice_DestroyQPagedPaintDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPagedPaintDevice) DestroyQPagedPaintDeviceDefault() {
	if ptr.Pointer() != nil {
		C.QPagedPaintDevice_DestroyQPagedPaintDeviceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQPagedPaintDevice_PaintEngine
func callbackQPagedPaintDevice_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return PointerFromQPaintEngine((*(*func() *QPaintEngine)(signal))())
	}

	return PointerFromQPaintEngine(NewQPagedPaintDeviceFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QPagedPaintDevice) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPagedPaintDevice_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPagedPaintDevice) PaintEngineDefault() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPagedPaintDevice_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

type QPaintDevice struct {
	ptr unsafe.Pointer
}

type QPaintDevice_ITF interface {
	QPaintDevice_PTR() *QPaintDevice
}

func (ptr *QPaintDevice) QPaintDevice_PTR() *QPaintDevice {
	return ptr
}

func (ptr *QPaintDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPaintDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPaintDevice(ptr QPaintDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func NewQPaintDeviceFromPointer(ptr unsafe.Pointer) (n *QPaintDevice) {
	n = new(QPaintDevice)
	n.SetPointer(ptr)
	return
}

// QPaintDevice::PaintDeviceMetric
//
//go:generate stringer -type=QPaintDevice__PaintDeviceMetric
type QPaintDevice__PaintDeviceMetric int64

const (
	QPaintDevice__PdmWidth                  QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(1)
	QPaintDevice__PdmHeight                 QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(2)
	QPaintDevice__PdmWidthMM                QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(3)
	QPaintDevice__PdmHeightMM               QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(4)
	QPaintDevice__PdmNumColors              QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(5)
	QPaintDevice__PdmDepth                  QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(6)
	QPaintDevice__PdmDpiX                   QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(7)
	QPaintDevice__PdmDpiY                   QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(8)
	QPaintDevice__PdmPhysicalDpiX           QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(9)
	QPaintDevice__PdmPhysicalDpiY           QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(10)
	QPaintDevice__PdmDevicePixelRatio       QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(11)
	QPaintDevice__PdmDevicePixelRatioScaled QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(12)
)

func NewQPaintDevice() *QPaintDevice {
	return NewQPaintDeviceFromPointer(C.QPaintDevice_NewQPaintDevice())
}

func (ptr *QPaintDevice) Height() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPaintDevice_Height(ptr.Pointer())))
	}
	return 0
}

//export callbackQPaintDevice_Metric
func callbackQPaintDevice_Metric(ptr unsafe.Pointer, metric C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(QPaintDevice__PaintDeviceMetric) int)(signal))(QPaintDevice__PaintDeviceMetric(metric))))
	}

	return C.int(int32(NewQPaintDeviceFromPointer(ptr).MetricDefault(QPaintDevice__PaintDeviceMetric(metric))))
}

func (ptr *QPaintDevice) ConnectMetric(f func(metric QPaintDevice__PaintDeviceMetric) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "metric"); signal != nil {
			f := func(metric QPaintDevice__PaintDeviceMetric) int {
				(*(*func(QPaintDevice__PaintDeviceMetric) int)(signal))(metric)
				return f(metric)
			}
			qt.ConnectSignal(ptr.Pointer(), "metric", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "metric", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintDevice) DisconnectMetric() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "metric")
	}
}

func (ptr *QPaintDevice) Metric(metric QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPaintDevice_Metric(ptr.Pointer(), C.longlong(metric))))
	}
	return 0
}

func (ptr *QPaintDevice) MetricDefault(metric QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPaintDevice_MetricDefault(ptr.Pointer(), C.longlong(metric))))
	}
	return 0
}

//export callbackQPaintDevice_PaintEngine
func callbackQPaintDevice_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return PointerFromQPaintEngine((*(*func() *QPaintEngine)(signal))())
	}

	return PointerFromQPaintEngine(NewQPaintEngine(0))
}

func (ptr *QPaintDevice) ConnectPaintEngine(f func() *QPaintEngine) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paintEngine"); signal != nil {
			f := func() *QPaintEngine {
				(*(*func() *QPaintEngine)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "paintEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paintEngine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintDevice) DisconnectPaintEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paintEngine")
	}
}

func (ptr *QPaintDevice) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPaintDevice_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintDevice) Width() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPaintDevice_Width(ptr.Pointer())))
	}
	return 0
}

//export callbackQPaintDevice_DestroyQPaintDevice
func callbackQPaintDevice_DestroyQPaintDevice(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPaintDevice"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPaintDeviceFromPointer(ptr).DestroyQPaintDeviceDefault()
	}
}

func (ptr *QPaintDevice) ConnectDestroyQPaintDevice(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPaintDevice"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPaintDevice", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPaintDevice", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintDevice) DisconnectDestroyQPaintDevice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPaintDevice")
	}
}

func (ptr *QPaintDevice) DestroyQPaintDevice() {
	if ptr.Pointer() != nil {
		C.QPaintDevice_DestroyQPaintDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPaintDevice) DestroyQPaintDeviceDefault() {
	if ptr.Pointer() != nil {
		C.QPaintDevice_DestroyQPaintDeviceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QPaintEngine struct {
	ptr unsafe.Pointer
}

type QPaintEngine_ITF interface {
	QPaintEngine_PTR() *QPaintEngine
}

func (ptr *QPaintEngine) QPaintEngine_PTR() *QPaintEngine {
	return ptr
}

func (ptr *QPaintEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPaintEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPaintEngine(ptr QPaintEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEngine_PTR().Pointer()
	}
	return nil
}

func NewQPaintEngineFromPointer(ptr unsafe.Pointer) (n *QPaintEngine) {
	n = new(QPaintEngine)
	n.SetPointer(ptr)
	return
}

// QPaintEngine::DirtyFlag
//
//go:generate stringer -type=QPaintEngine__DirtyFlag
type QPaintEngine__DirtyFlag int64

const (
	QPaintEngine__DirtyPen             QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0001)
	QPaintEngine__DirtyBrush           QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0002)
	QPaintEngine__DirtyBrushOrigin     QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0004)
	QPaintEngine__DirtyFont            QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0008)
	QPaintEngine__DirtyBackground      QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0010)
	QPaintEngine__DirtyBackgroundMode  QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0020)
	QPaintEngine__DirtyTransform       QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0040)
	QPaintEngine__DirtyClipRegion      QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0080)
	QPaintEngine__DirtyClipPath        QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0100)
	QPaintEngine__DirtyHints           QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0200)
	QPaintEngine__DirtyCompositionMode QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0400)
	QPaintEngine__DirtyClipEnabled     QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0800)
	QPaintEngine__DirtyOpacity         QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x1000)
	QPaintEngine__AllDirty             QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0xffff)
)

// QPaintEngine::PaintEngineFeature
//
//go:generate stringer -type=QPaintEngine__PaintEngineFeature
type QPaintEngine__PaintEngineFeature int64

const (
	QPaintEngine__PrimitiveTransform          QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000001)
	QPaintEngine__PatternTransform            QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000002)
	QPaintEngine__PixmapTransform             QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000004)
	QPaintEngine__PatternBrush                QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000008)
	QPaintEngine__LinearGradientFill          QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000010)
	QPaintEngine__RadialGradientFill          QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000020)
	QPaintEngine__ConicalGradientFill         QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000040)
	QPaintEngine__AlphaBlend                  QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000080)
	QPaintEngine__PorterDuff                  QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000100)
	QPaintEngine__PainterPaths                QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000200)
	QPaintEngine__Antialiasing                QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000400)
	QPaintEngine__BrushStroke                 QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000800)
	QPaintEngine__ConstantOpacity             QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00001000)
	QPaintEngine__MaskedBrush                 QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00002000)
	QPaintEngine__PerspectiveTransform        QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00004000)
	QPaintEngine__BlendModes                  QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00008000)
	QPaintEngine__ObjectBoundingModeGradients QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00010000)
	QPaintEngine__RasterOpModes               QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00020000)
	QPaintEngine__PaintOutsidePaintEvent      QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x20000000)
	QPaintEngine__AllFeatures                 QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0xffffffff)
)

// QPaintEngine::PolygonDrawMode
//
//go:generate stringer -type=QPaintEngine__PolygonDrawMode
type QPaintEngine__PolygonDrawMode int64

const (
	QPaintEngine__OddEvenMode  QPaintEngine__PolygonDrawMode = QPaintEngine__PolygonDrawMode(0)
	QPaintEngine__WindingMode  QPaintEngine__PolygonDrawMode = QPaintEngine__PolygonDrawMode(1)
	QPaintEngine__ConvexMode   QPaintEngine__PolygonDrawMode = QPaintEngine__PolygonDrawMode(2)
	QPaintEngine__PolylineMode QPaintEngine__PolygonDrawMode = QPaintEngine__PolygonDrawMode(3)
)

// QPaintEngine::Type
//
//go:generate stringer -type=QPaintEngine__Type
type QPaintEngine__Type int64

const (
	QPaintEngine__X11           QPaintEngine__Type = QPaintEngine__Type(0)
	QPaintEngine__Windows       QPaintEngine__Type = QPaintEngine__Type(1)
	QPaintEngine__QuickDraw     QPaintEngine__Type = QPaintEngine__Type(2)
	QPaintEngine__CoreGraphics  QPaintEngine__Type = QPaintEngine__Type(3)
	QPaintEngine__MacPrinter    QPaintEngine__Type = QPaintEngine__Type(4)
	QPaintEngine__QWindowSystem QPaintEngine__Type = QPaintEngine__Type(5)
	QPaintEngine__PostScript    QPaintEngine__Type = QPaintEngine__Type(6)
	QPaintEngine__OpenGL        QPaintEngine__Type = QPaintEngine__Type(7)
	QPaintEngine__Picture       QPaintEngine__Type = QPaintEngine__Type(8)
	QPaintEngine__SVG           QPaintEngine__Type = QPaintEngine__Type(9)
	QPaintEngine__Raster        QPaintEngine__Type = QPaintEngine__Type(10)
	QPaintEngine__Direct3D      QPaintEngine__Type = QPaintEngine__Type(11)
	QPaintEngine__Pdf           QPaintEngine__Type = QPaintEngine__Type(12)
	QPaintEngine__OpenVG        QPaintEngine__Type = QPaintEngine__Type(13)
	QPaintEngine__OpenGL2       QPaintEngine__Type = QPaintEngine__Type(14)
	QPaintEngine__PaintBuffer   QPaintEngine__Type = QPaintEngine__Type(15)
	QPaintEngine__Blitter       QPaintEngine__Type = QPaintEngine__Type(16)
	QPaintEngine__Direct2D      QPaintEngine__Type = QPaintEngine__Type(17)
	QPaintEngine__User          QPaintEngine__Type = QPaintEngine__Type(50)
	QPaintEngine__MaxUser       QPaintEngine__Type = QPaintEngine__Type(100)
)

func NewQPaintEngine(caps QPaintEngine__PaintEngineFeature) *QPaintEngine {
	return NewQPaintEngineFromPointer(C.QPaintEngine_NewQPaintEngine(C.longlong(caps)))
}

//export callbackQPaintEngine_Begin
func callbackQPaintEngine_Begin(ptr unsafe.Pointer, pdev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "begin"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QPaintDevice) bool)(signal))(NewQPaintDeviceFromPointer(pdev)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPaintEngine) ConnectBegin(f func(pdev *QPaintDevice) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "begin"); signal != nil {
			f := func(pdev *QPaintDevice) bool {
				(*(*func(*QPaintDevice) bool)(signal))(pdev)
				return f(pdev)
			}
			qt.ConnectSignal(ptr.Pointer(), "begin", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "begin", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectBegin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "begin")
	}
}

func (ptr *QPaintEngine) Begin(pdev QPaintDevice_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPaintEngine_Begin(ptr.Pointer(), PointerFromQPaintDevice(pdev))) != 0
	}
	return false
}

//export callbackQPaintEngine_DrawPixmap
func callbackQPaintEngine_DrawPixmap(ptr unsafe.Pointer, r unsafe.Pointer, pm unsafe.Pointer, sr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawPixmap"); signal != nil {
		(*(*func(*core.QRectF, *QPixmap, *core.QRectF))(signal))(core.NewQRectFFromPointer(r), NewQPixmapFromPointer(pm), core.NewQRectFFromPointer(sr))
	}

}

func (ptr *QPaintEngine) ConnectDrawPixmap(f func(r *core.QRectF, pm *QPixmap, sr *core.QRectF)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "drawPixmap"); signal != nil {
			f := func(r *core.QRectF, pm *QPixmap, sr *core.QRectF) {
				(*(*func(*core.QRectF, *QPixmap, *core.QRectF))(signal))(r, pm, sr)
				f(r, pm, sr)
			}
			qt.ConnectSignal(ptr.Pointer(), "drawPixmap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "drawPixmap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectDrawPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "drawPixmap")
	}
}

func (ptr *QPaintEngine) DrawPixmap(r core.QRectF_ITF, pm QPixmap_ITF, sr core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPixmap(ptr.Pointer(), core.PointerFromQRectF(r), PointerFromQPixmap(pm), core.PointerFromQRectF(sr))
	}
}

//export callbackQPaintEngine_End
func callbackQPaintEngine_End(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "end"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPaintEngine) ConnectEnd(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "end"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "end", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "end", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectEnd() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "end")
	}
}

func (ptr *QPaintEngine) End() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPaintEngine_End(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPaintEngine) SetActive(state bool) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_SetActive(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(state))))
	}
}

//export callbackQPaintEngine_Type
func callbackQPaintEngine_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QPaintEngine__Type)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QPaintEngine) ConnectType(f func() QPaintEngine__Type) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QPaintEngine__Type {
				(*(*func() QPaintEngine__Type)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QPaintEngine) Type() QPaintEngine__Type {
	if ptr.Pointer() != nil {
		return QPaintEngine__Type(C.QPaintEngine_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQPaintEngine_UpdateState
func callbackQPaintEngine_UpdateState(ptr unsafe.Pointer, state unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateState"); signal != nil {
		(*(*func(*QPaintEngineState))(signal))(NewQPaintEngineStateFromPointer(state))
	}

}

func (ptr *QPaintEngine) ConnectUpdateState(f func(state *QPaintEngineState)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateState"); signal != nil {
			f := func(state *QPaintEngineState) {
				(*(*func(*QPaintEngineState))(signal))(state)
				f(state)
			}
			qt.ConnectSignal(ptr.Pointer(), "updateState", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateState", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectUpdateState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateState")
	}
}

func (ptr *QPaintEngine) UpdateState(state QPaintEngineState_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_UpdateState(ptr.Pointer(), PointerFromQPaintEngineState(state))
	}
}

//export callbackQPaintEngine_DestroyQPaintEngine
func callbackQPaintEngine_DestroyQPaintEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPaintEngine"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPaintEngineFromPointer(ptr).DestroyQPaintEngineDefault()
	}
}

func (ptr *QPaintEngine) ConnectDestroyQPaintEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPaintEngine"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPaintEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPaintEngine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectDestroyQPaintEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPaintEngine")
	}
}

func (ptr *QPaintEngine) DestroyQPaintEngine() {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DestroyQPaintEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPaintEngine) DestroyQPaintEngineDefault() {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DestroyQPaintEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QPaintEngineState struct {
	ptr unsafe.Pointer
}

type QPaintEngineState_ITF interface {
	QPaintEngineState_PTR() *QPaintEngineState
}

func (ptr *QPaintEngineState) QPaintEngineState_PTR() *QPaintEngineState {
	return ptr
}

func (ptr *QPaintEngineState) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPaintEngineState) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPaintEngineState(ptr QPaintEngineState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEngineState_PTR().Pointer()
	}
	return nil
}

func NewQPaintEngineStateFromPointer(ptr unsafe.Pointer) (n *QPaintEngineState) {
	n = new(QPaintEngineState)
	n.SetPointer(ptr)
	return
}
func (ptr *QPaintEngineState) DestroyQPaintEngineState() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func (ptr *QPaintEngineState) Brush() *QBrush {
	if ptr.Pointer() != nil {
		tmpValue := NewQBrushFromPointer(C.QPaintEngineState_Brush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QPaintEngineState) Font() *QFont {
	if ptr.Pointer() != nil {
		tmpValue := NewQFontFromPointer(C.QPaintEngineState_Font(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QPaintEngineState) Transform() *QTransform {
	if ptr.Pointer() != nil {
		tmpValue := NewQTransformFromPointer(C.QPaintEngineState_Transform(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
		return tmpValue
	}
	return nil
}

type QPaintEvent struct {
	core.QEvent
}

type QPaintEvent_ITF interface {
	core.QEvent_ITF
	QPaintEvent_PTR() *QPaintEvent
}

func (ptr *QPaintEvent) QPaintEvent_PTR() *QPaintEvent {
	return ptr
}

func (ptr *QPaintEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QPaintEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQPaintEvent(ptr QPaintEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEvent_PTR().Pointer()
	}
	return nil
}

func NewQPaintEventFromPointer(ptr unsafe.Pointer) (n *QPaintEvent) {
	n = new(QPaintEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QPaintEvent) DestroyQPaintEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQPaintEvent(paintRegion QRegion_ITF) *QPaintEvent {
	tmpValue := NewQPaintEventFromPointer(C.QPaintEvent_NewQPaintEvent(PointerFromQRegion(paintRegion)))
	qt.SetFinalizer(tmpValue, (*QPaintEvent).DestroyQPaintEvent)
	return tmpValue
}

func NewQPaintEvent2(paintRect core.QRect_ITF) *QPaintEvent {
	tmpValue := NewQPaintEventFromPointer(C.QPaintEvent_NewQPaintEvent2(core.PointerFromQRect(paintRect)))
	qt.SetFinalizer(tmpValue, (*QPaintEvent).DestroyQPaintEvent)
	return tmpValue
}

func (ptr *QPaintEvent) Rect() *core.QRect {
	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QPaintEvent_Rect(ptr.Pointer()))
	}
	return nil
}

type QPainter struct {
	ptr unsafe.Pointer
}

type QPainter_ITF interface {
	QPainter_PTR() *QPainter
}

func (ptr *QPainter) QPainter_PTR() *QPainter {
	return ptr
}

func (ptr *QPainter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPainter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPainter(ptr QPainter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPainter_PTR().Pointer()
	}
	return nil
}

func NewQPainterFromPointer(ptr unsafe.Pointer) (n *QPainter) {
	n = new(QPainter)
	n.SetPointer(ptr)
	return
}

// QPainter::CompositionMode
//
//go:generate stringer -type=QPainter__CompositionMode
type QPainter__CompositionMode int64

const (
	QPainter__CompositionMode_SourceOver          QPainter__CompositionMode = QPainter__CompositionMode(0)
	QPainter__CompositionMode_DestinationOver     QPainter__CompositionMode = QPainter__CompositionMode(1)
	QPainter__CompositionMode_Clear               QPainter__CompositionMode = QPainter__CompositionMode(2)
	QPainter__CompositionMode_Source              QPainter__CompositionMode = QPainter__CompositionMode(3)
	QPainter__CompositionMode_Destination         QPainter__CompositionMode = QPainter__CompositionMode(4)
	QPainter__CompositionMode_SourceIn            QPainter__CompositionMode = QPainter__CompositionMode(5)
	QPainter__CompositionMode_DestinationIn       QPainter__CompositionMode = QPainter__CompositionMode(6)
	QPainter__CompositionMode_SourceOut           QPainter__CompositionMode = QPainter__CompositionMode(7)
	QPainter__CompositionMode_DestinationOut      QPainter__CompositionMode = QPainter__CompositionMode(8)
	QPainter__CompositionMode_SourceAtop          QPainter__CompositionMode = QPainter__CompositionMode(9)
	QPainter__CompositionMode_DestinationAtop     QPainter__CompositionMode = QPainter__CompositionMode(10)
	QPainter__CompositionMode_Xor                 QPainter__CompositionMode = QPainter__CompositionMode(11)
	QPainter__CompositionMode_Plus                QPainter__CompositionMode = QPainter__CompositionMode(12)
	QPainter__CompositionMode_Multiply            QPainter__CompositionMode = QPainter__CompositionMode(13)
	QPainter__CompositionMode_Screen              QPainter__CompositionMode = QPainter__CompositionMode(14)
	QPainter__CompositionMode_Overlay             QPainter__CompositionMode = QPainter__CompositionMode(15)
	QPainter__CompositionMode_Darken              QPainter__CompositionMode = QPainter__CompositionMode(16)
	QPainter__CompositionMode_Lighten             QPainter__CompositionMode = QPainter__CompositionMode(17)
	QPainter__CompositionMode_ColorDodge          QPainter__CompositionMode = QPainter__CompositionMode(18)
	QPainter__CompositionMode_ColorBurn           QPainter__CompositionMode = QPainter__CompositionMode(19)
	QPainter__CompositionMode_HardLight           QPainter__CompositionMode = QPainter__CompositionMode(20)
	QPainter__CompositionMode_SoftLight           QPainter__CompositionMode = QPainter__CompositionMode(21)
	QPainter__CompositionMode_Difference          QPainter__CompositionMode = QPainter__CompositionMode(22)
	QPainter__CompositionMode_Exclusion           QPainter__CompositionMode = QPainter__CompositionMode(23)
	QPainter__RasterOp_SourceOrDestination        QPainter__CompositionMode = QPainter__CompositionMode(24)
	QPainter__RasterOp_SourceAndDestination       QPainter__CompositionMode = QPainter__CompositionMode(25)
	QPainter__RasterOp_SourceXorDestination       QPainter__CompositionMode = QPainter__CompositionMode(26)
	QPainter__RasterOp_NotSourceAndNotDestination QPainter__CompositionMode = QPainter__CompositionMode(27)
	QPainter__RasterOp_NotSourceOrNotDestination  QPainter__CompositionMode = QPainter__CompositionMode(28)
	QPainter__RasterOp_NotSourceXorDestination    QPainter__CompositionMode = QPainter__CompositionMode(29)
	QPainter__RasterOp_NotSource                  QPainter__CompositionMode = QPainter__CompositionMode(30)
	QPainter__RasterOp_NotSourceAndDestination    QPainter__CompositionMode = QPainter__CompositionMode(31)
	QPainter__RasterOp_SourceAndNotDestination    QPainter__CompositionMode = QPainter__CompositionMode(32)
	QPainter__RasterOp_NotSourceOrDestination     QPainter__CompositionMode = QPainter__CompositionMode(33)
	QPainter__RasterOp_SourceOrNotDestination     QPainter__CompositionMode = QPainter__CompositionMode(34)
	QPainter__RasterOp_ClearDestination           QPainter__CompositionMode = QPainter__CompositionMode(35)
	QPainter__RasterOp_SetDestination             QPainter__CompositionMode = QPainter__CompositionMode(36)
	QPainter__RasterOp_NotDestination             QPainter__CompositionMode = QPainter__CompositionMode(37)
)

// QPainter::PixmapFragmentHint
//
//go:generate stringer -type=QPainter__PixmapFragmentHint
type QPainter__PixmapFragmentHint int64

const (
	QPainter__OpaqueHint QPainter__PixmapFragmentHint = QPainter__PixmapFragmentHint(0x01)
)

// QPainter::RenderHint
//
//go:generate stringer -type=QPainter__RenderHint
type QPainter__RenderHint int64

const (
	QPainter__Antialiasing            QPainter__RenderHint = QPainter__RenderHint(0x01)
	QPainter__TextAntialiasing        QPainter__RenderHint = QPainter__RenderHint(0x02)
	QPainter__SmoothPixmapTransform   QPainter__RenderHint = QPainter__RenderHint(0x04)
	QPainter__HighQualityAntialiasing QPainter__RenderHint = QPainter__RenderHint(0x08)
	QPainter__NonCosmeticDefaultPen   QPainter__RenderHint = QPainter__RenderHint(0x10)
	QPainter__Qt4CompatiblePainting   QPainter__RenderHint = QPainter__RenderHint(0x20)
	QPainter__LosslessImageRendering  QPainter__RenderHint = QPainter__RenderHint(0x40)
)

func NewQPainter() *QPainter {
	tmpValue := NewQPainterFromPointer(C.QPainter_NewQPainter())
	qt.SetFinalizer(tmpValue, (*QPainter).DestroyQPainter)
	return tmpValue
}

func NewQPainter2(device QPaintDevice_ITF) *QPainter {
	tmpValue := NewQPainterFromPointer(C.QPainter_NewQPainter2(PointerFromQPaintDevice(device)))
	qt.SetFinalizer(tmpValue, (*QPainter).DestroyQPainter)
	return tmpValue
}

func (ptr *QPainter) Background() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPainter_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPainter) Begin(device QPaintDevice_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPainter_Begin(ptr.Pointer(), PointerFromQPaintDevice(device))) != 0
	}
	return false
}

func (ptr *QPainter) BoundingRect(rectangle core.QRectF_ITF, flags int, text string) *core.QRectF {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := core.NewQRectFFromPointer(C.QPainter_BoundingRect(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.int(int32(flags)), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) BoundingRect2(rectangle core.QRect_ITF, flags int, text string) *core.QRect {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := core.NewQRectFromPointer(C.QPainter_BoundingRect2(ptr.Pointer(), core.PointerFromQRect(rectangle), C.int(int32(flags)), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) BoundingRect3(x int, y int, w int, h int, flags int, text string) *core.QRect {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := core.NewQRectFromPointer(C.QPainter_BoundingRect3(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(w)), C.int(int32(h)), C.int(int32(flags)), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) BoundingRect4(rectangle core.QRectF_ITF, text string, option QTextOption_ITF) *core.QRectF {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := core.NewQRectFFromPointer(C.QPainter_BoundingRect4(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQTextOption(option)))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) Brush() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPainter_Brush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPainter) Device() *QPaintDevice {
	if ptr.Pointer() != nil {
		return NewQPaintDeviceFromPointer(C.QPainter_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPainter) End() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPainter_End(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPainter) Font() *QFont {
	if ptr.Pointer() != nil {
		return NewQFontFromPointer(C.QPainter_Font(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPainter) FontMetrics() *QFontMetrics {
	if ptr.Pointer() != nil {
		tmpValue := NewQFontMetricsFromPointer(C.QPainter_FontMetrics(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QFontMetrics).DestroyQFontMetrics)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) Save() {
	if ptr.Pointer() != nil {
		C.QPainter_Save(ptr.Pointer())
	}
}

func (ptr *QPainter) Scale(sx float64, sy float64) {
	if ptr.Pointer() != nil {
		C.QPainter_Scale(ptr.Pointer(), C.double(sx), C.double(sy))
	}
}

func (ptr *QPainter) SetBackground(brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetBackground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QPainter) SetFont(font QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetFont(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QPainter) SetWindow(rectangle core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetWindow(ptr.Pointer(), core.PointerFromQRect(rectangle))
	}
}

func (ptr *QPainter) SetWindow2(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QPainter_SetWindow2(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(width)), C.int(int32(height)))
	}
}

func (ptr *QPainter) Transform() *QTransform {
	if ptr.Pointer() != nil {
		return NewQTransformFromPointer(C.QPainter_Transform(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPainter) Window() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QPainter_Window(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) DestroyQPainter() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPainter_DestroyQPainter(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPainter) __drawLines_lines_atList2(i int) *core.QLineF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLineFFromPointer(C.QPainter___drawLines_lines_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QLineF).DestroyQLineF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawLines_lines_setList2(i core.QLineF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawLines_lines_setList2(ptr.Pointer(), core.PointerFromQLineF(i))
	}
}

func (ptr *QPainter) __drawLines_lines_newList2() unsafe.Pointer {
	return C.QPainter___drawLines_lines_newList2(ptr.Pointer())
}

func (ptr *QPainter) __drawLines_pointPairs_atList4(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QPainter___drawLines_pointPairs_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawLines_pointPairs_setList4(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawLines_pointPairs_setList4(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QPainter) __drawLines_pointPairs_newList4() unsafe.Pointer {
	return C.QPainter___drawLines_pointPairs_newList4(ptr.Pointer())
}

func (ptr *QPainter) __drawLines_lines_atList6(i int) *core.QLine {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLineFromPointer(C.QPainter___drawLines_lines_atList6(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QLine).DestroyQLine)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawLines_lines_setList6(i core.QLine_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawLines_lines_setList6(ptr.Pointer(), core.PointerFromQLine(i))
	}
}

func (ptr *QPainter) __drawLines_lines_newList6() unsafe.Pointer {
	return C.QPainter___drawLines_lines_newList6(ptr.Pointer())
}

func (ptr *QPainter) __drawLines_pointPairs_atList8(i int) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QPainter___drawLines_pointPairs_atList8(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawLines_pointPairs_setList8(i core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawLines_pointPairs_setList8(ptr.Pointer(), core.PointerFromQPoint(i))
	}
}

func (ptr *QPainter) __drawLines_pointPairs_newList8() unsafe.Pointer {
	return C.QPainter___drawLines_pointPairs_newList8(ptr.Pointer())
}

func (ptr *QPainter) __drawRects_rectangles_atList2(i int) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QPainter___drawRects_rectangles_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawRects_rectangles_setList2(i core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawRects_rectangles_setList2(ptr.Pointer(), core.PointerFromQRectF(i))
	}
}

func (ptr *QPainter) __drawRects_rectangles_newList2() unsafe.Pointer {
	return C.QPainter___drawRects_rectangles_newList2(ptr.Pointer())
}

func (ptr *QPainter) __drawRects_rectangles_atList4(i int) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QPainter___drawRects_rectangles_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawRects_rectangles_setList4(i core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawRects_rectangles_setList4(ptr.Pointer(), core.PointerFromQRect(i))
	}
}

func (ptr *QPainter) __drawRects_rectangles_newList4() unsafe.Pointer {
	return C.QPainter___drawRects_rectangles_newList4(ptr.Pointer())
}

type QPainterPath struct {
	ptr unsafe.Pointer
}

type QPainterPath_ITF interface {
	QPainterPath_PTR() *QPainterPath
}

func (ptr *QPainterPath) QPainterPath_PTR() *QPainterPath {
	return ptr
}

func (ptr *QPainterPath) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPainterPath) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPainterPath(ptr QPainterPath_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPainterPath_PTR().Pointer()
	}
	return nil
}

func NewQPainterPathFromPointer(ptr unsafe.Pointer) (n *QPainterPath) {
	n = new(QPainterPath)
	n.SetPointer(ptr)
	return
}

// QPainterPath::ElementType
//
//go:generate stringer -type=QPainterPath__ElementType
type QPainterPath__ElementType int64

const (
	QPainterPath__MoveToElement      QPainterPath__ElementType = QPainterPath__ElementType(0)
	QPainterPath__LineToElement      QPainterPath__ElementType = QPainterPath__ElementType(1)
	QPainterPath__CurveToElement     QPainterPath__ElementType = QPainterPath__ElementType(2)
	QPainterPath__CurveToDataElement QPainterPath__ElementType = QPainterPath__ElementType(3)
)

func NewQPainterPath() *QPainterPath {
	tmpValue := NewQPainterPathFromPointer(C.QPainterPath_NewQPainterPath())
	qt.SetFinalizer(tmpValue, (*QPainterPath).DestroyQPainterPath)
	return tmpValue
}

func NewQPainterPath2(startPoint core.QPointF_ITF) *QPainterPath {
	tmpValue := NewQPainterPathFromPointer(C.QPainterPath_NewQPainterPath2(core.PointerFromQPointF(startPoint)))
	qt.SetFinalizer(tmpValue, (*QPainterPath).DestroyQPainterPath)
	return tmpValue
}

func NewQPainterPath3(path QPainterPath_ITF) *QPainterPath {
	tmpValue := NewQPainterPathFromPointer(C.QPainterPath_NewQPainterPath3(PointerFromQPainterPath(path)))
	qt.SetFinalizer(tmpValue, (*QPainterPath).DestroyQPainterPath)
	return tmpValue
}

func (ptr *QPainterPath) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QPainterPath_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainterPath) Clear() {
	if ptr.Pointer() != nil {
		C.QPainterPath_Clear(ptr.Pointer())
	}
}

func (ptr *QPainterPath) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPainterPath_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPainterPath) Length() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPainterPath_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPath) DestroyQPainterPath() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPainterPath_DestroyQPainterPath(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPainterPath) __toFillPolygons_atList(i int) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QPainterPath___toFillPolygons_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainterPath) __toFillPolygons_setList(i QPolygonF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath___toFillPolygons_setList(ptr.Pointer(), PointerFromQPolygonF(i))
	}
}

func (ptr *QPainterPath) __toFillPolygons_newList() unsafe.Pointer {
	return C.QPainterPath___toFillPolygons_newList(ptr.Pointer())
}

func (ptr *QPainterPath) __toFillPolygons_atList2(i int) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QPainterPath___toFillPolygons_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainterPath) __toFillPolygons_setList2(i QPolygonF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath___toFillPolygons_setList2(ptr.Pointer(), PointerFromQPolygonF(i))
	}
}

func (ptr *QPainterPath) __toFillPolygons_newList2() unsafe.Pointer {
	return C.QPainterPath___toFillPolygons_newList2(ptr.Pointer())
}

func (ptr *QPainterPath) __toSubpathPolygons_atList(i int) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QPainterPath___toSubpathPolygons_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainterPath) __toSubpathPolygons_setList(i QPolygonF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath___toSubpathPolygons_setList(ptr.Pointer(), PointerFromQPolygonF(i))
	}
}

func (ptr *QPainterPath) __toSubpathPolygons_newList() unsafe.Pointer {
	return C.QPainterPath___toSubpathPolygons_newList(ptr.Pointer())
}

func (ptr *QPainterPath) __toSubpathPolygons_atList2(i int) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QPainterPath___toSubpathPolygons_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainterPath) __toSubpathPolygons_setList2(i QPolygonF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath___toSubpathPolygons_setList2(ptr.Pointer(), PointerFromQPolygonF(i))
	}
}

func (ptr *QPainterPath) __toSubpathPolygons_newList2() unsafe.Pointer {
	return C.QPainterPath___toSubpathPolygons_newList2(ptr.Pointer())
}

// QPalette::ColorGroup
//
//go:generate stringer -type=QPalette__ColorGroup
type QPalette__ColorGroup int64

const (
	QPalette__Active       QPalette__ColorGroup = QPalette__ColorGroup(0)
	QPalette__Disabled     QPalette__ColorGroup = QPalette__ColorGroup(1)
	QPalette__Inactive     QPalette__ColorGroup = QPalette__ColorGroup(2)
	QPalette__NColorGroups QPalette__ColorGroup = QPalette__ColorGroup(3)
	QPalette__Current      QPalette__ColorGroup = QPalette__ColorGroup(4)
	QPalette__All          QPalette__ColorGroup = QPalette__ColorGroup(5)
	QPalette__Normal       QPalette__ColorGroup = QPalette__ColorGroup(QPalette__Active)
)

// QPalette::ColorRole
//
//go:generate stringer -type=QPalette__ColorRole
type QPalette__ColorRole int64

var (
	QPalette__WindowText      QPalette__ColorRole = QPalette__ColorRole(0)
	QPalette__Button          QPalette__ColorRole = QPalette__ColorRole(1)
	QPalette__Light           QPalette__ColorRole = QPalette__ColorRole(2)
	QPalette__Midlight        QPalette__ColorRole = QPalette__ColorRole(3)
	QPalette__Dark            QPalette__ColorRole = QPalette__ColorRole(4)
	QPalette__Mid             QPalette__ColorRole = QPalette__ColorRole(5)
	QPalette__Text            QPalette__ColorRole = QPalette__ColorRole(6)
	QPalette__BrightText      QPalette__ColorRole = QPalette__ColorRole(7)
	QPalette__ButtonText      QPalette__ColorRole = QPalette__ColorRole(8)
	QPalette__Base            QPalette__ColorRole = QPalette__ColorRole(9)
	QPalette__Window          QPalette__ColorRole = QPalette__ColorRole(10)
	QPalette__Shadow          QPalette__ColorRole = QPalette__ColorRole(11)
	QPalette__Highlight       QPalette__ColorRole = QPalette__ColorRole(12)
	QPalette__HighlightedText QPalette__ColorRole = QPalette__ColorRole(13)
	QPalette__Link            QPalette__ColorRole = QPalette__ColorRole(14)
	QPalette__LinkVisited     QPalette__ColorRole = QPalette__ColorRole(15)
	QPalette__AlternateBase   QPalette__ColorRole = QPalette__ColorRole(16)
	QPalette__NoRole          QPalette__ColorRole = QPalette__ColorRole(17)
	QPalette__ToolTipBase     QPalette__ColorRole = QPalette__ColorRole(18)
	QPalette__ToolTipText     QPalette__ColorRole = QPalette__ColorRole(19)
	QPalette__PlaceholderText QPalette__ColorRole = QPalette__ColorRole(20)
	QPalette__NColorRoles     QPalette__ColorRole = QPalette__ColorRole(C.QPalette_NColorRoles_Type())
	QPalette__Foreground      QPalette__ColorRole = QPalette__ColorRole(QPalette__WindowText)
	QPalette__Background      QPalette__ColorRole = QPalette__ColorRole(QPalette__Window)
)

// QPixelFormat::AlphaPosition
//
//go:generate stringer -type=QPixelFormat__AlphaPosition
type QPixelFormat__AlphaPosition int64

const (
	QPixelFormat__AtBeginning QPixelFormat__AlphaPosition = QPixelFormat__AlphaPosition(0)
	QPixelFormat__AtEnd       QPixelFormat__AlphaPosition = QPixelFormat__AlphaPosition(1)
)

// QPixelFormat::AlphaPremultiplied
//
//go:generate stringer -type=QPixelFormat__AlphaPremultiplied
type QPixelFormat__AlphaPremultiplied int64

const (
	QPixelFormat__NotPremultiplied QPixelFormat__AlphaPremultiplied = QPixelFormat__AlphaPremultiplied(0)
	QPixelFormat__Premultiplied    QPixelFormat__AlphaPremultiplied = QPixelFormat__AlphaPremultiplied(1)
)

// QPixelFormat::AlphaUsage
//
//go:generate stringer -type=QPixelFormat__AlphaUsage
type QPixelFormat__AlphaUsage int64

const (
	QPixelFormat__UsesAlpha    QPixelFormat__AlphaUsage = QPixelFormat__AlphaUsage(0)
	QPixelFormat__IgnoresAlpha QPixelFormat__AlphaUsage = QPixelFormat__AlphaUsage(1)
)

// QPixelFormat::ByteOrder
//
//go:generate stringer -type=QPixelFormat__ByteOrder
type QPixelFormat__ByteOrder int64

const (
	QPixelFormat__LittleEndian        QPixelFormat__ByteOrder = QPixelFormat__ByteOrder(0)
	QPixelFormat__BigEndian           QPixelFormat__ByteOrder = QPixelFormat__ByteOrder(1)
	QPixelFormat__CurrentSystemEndian QPixelFormat__ByteOrder = QPixelFormat__ByteOrder(2)
)

// QPixelFormat::ColorModel
//
//go:generate stringer -type=QPixelFormat__ColorModel
type QPixelFormat__ColorModel int64

const (
	QPixelFormat__RGB       QPixelFormat__ColorModel = QPixelFormat__ColorModel(0)
	QPixelFormat__BGR       QPixelFormat__ColorModel = QPixelFormat__ColorModel(1)
	QPixelFormat__Indexed   QPixelFormat__ColorModel = QPixelFormat__ColorModel(2)
	QPixelFormat__Grayscale QPixelFormat__ColorModel = QPixelFormat__ColorModel(3)
	QPixelFormat__CMYK      QPixelFormat__ColorModel = QPixelFormat__ColorModel(4)
	QPixelFormat__HSL       QPixelFormat__ColorModel = QPixelFormat__ColorModel(5)
	QPixelFormat__HSV       QPixelFormat__ColorModel = QPixelFormat__ColorModel(6)
	QPixelFormat__YUV       QPixelFormat__ColorModel = QPixelFormat__ColorModel(7)
	QPixelFormat__Alpha     QPixelFormat__ColorModel = QPixelFormat__ColorModel(8)
)

// QPixelFormat::TypeInterpretation
//
//go:generate stringer -type=QPixelFormat__TypeInterpretation
type QPixelFormat__TypeInterpretation int64

const (
	QPixelFormat__UnsignedInteger QPixelFormat__TypeInterpretation = QPixelFormat__TypeInterpretation(0)
	QPixelFormat__UnsignedShort   QPixelFormat__TypeInterpretation = QPixelFormat__TypeInterpretation(1)
	QPixelFormat__UnsignedByte    QPixelFormat__TypeInterpretation = QPixelFormat__TypeInterpretation(2)
	QPixelFormat__FloatingPoint   QPixelFormat__TypeInterpretation = QPixelFormat__TypeInterpretation(3)
)

// QPixelFormat::YUVLayout
//
//go:generate stringer -type=QPixelFormat__YUVLayout
type QPixelFormat__YUVLayout int64

const (
	QPixelFormat__YUV444   QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(0)
	QPixelFormat__YUV422   QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(1)
	QPixelFormat__YUV411   QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(2)
	QPixelFormat__YUV420P  QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(3)
	QPixelFormat__YUV420SP QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(4)
	QPixelFormat__YV12     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(5)
	QPixelFormat__UYVY     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(6)
	QPixelFormat__YUYV     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(7)
	QPixelFormat__NV12     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(8)
	QPixelFormat__NV21     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(9)
	QPixelFormat__IMC1     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(10)
	QPixelFormat__IMC2     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(11)
	QPixelFormat__IMC3     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(12)
	QPixelFormat__IMC4     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(13)
	QPixelFormat__Y8       QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(14)
	QPixelFormat__Y16      QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(15)
)

type QPixmap struct {
	QPaintDevice
}

type QPixmap_ITF interface {
	QPaintDevice_ITF
	QPixmap_PTR() *QPixmap
}

func (ptr *QPixmap) QPixmap_PTR() *QPixmap {
	return ptr
}

func (ptr *QPixmap) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QPixmap) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQPixmap(ptr QPixmap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPixmap_PTR().Pointer()
	}
	return nil
}

func NewQPixmapFromPointer(ptr unsafe.Pointer) (n *QPixmap) {
	n = new(QPixmap)
	n.SetPointer(ptr)
	return
}
func NewQPixmap() *QPixmap {
	return NewQPixmapFromPointer(C.QPixmap_NewQPixmap())
}

func NewQPixmap2(size core.QSize_ITF) *QPixmap {
	return NewQPixmapFromPointer(C.QPixmap_NewQPixmap2(core.PointerFromQSize(size)))
}

func NewQPixmap3(fileName string, format string, flags core.Qt__ImageConversionFlag) *QPixmap {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	return NewQPixmapFromPointer(C.QPixmap_NewQPixmap3(C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, formatC, C.longlong(flags)))
}

func NewQPixmap5(pixmap QPixmap_ITF) *QPixmap {
	return NewQPixmapFromPointer(C.QPixmap_NewQPixmap5(PointerFromQPixmap(pixmap)))
}

func (ptr *QPixmap) Copy(rectangle core.QRect_ITF) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QPixmap_Copy(ptr.Pointer(), core.PointerFromQRect(rectangle)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QPixmap) Copy2(x int, y int, width int, height int) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QPixmap_Copy2(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(width)), C.int(int32(height))))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QPixmap) Load(fileName string, format string, flags core.Qt__ImageConversionFlag) bool {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QPixmap_Load(ptr.Pointer(), C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, formatC, C.longlong(flags))) != 0
	}
	return false
}

func (ptr *QPixmap) LoadFromData(data []byte, l uint, format string, flags core.Qt__ImageConversionFlag) bool {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QPixmap_LoadFromData(ptr.Pointer(), dataC, C.uint(uint32(l)), formatC, C.longlong(flags))) != 0
	}
	return false
}

func (ptr *QPixmap) LoadFromData2(data core.QByteArray_ITF, format string, flags core.Qt__ImageConversionFlag) bool {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QPixmap_LoadFromData2(ptr.Pointer(), core.PointerFromQByteArray(data), formatC, C.longlong(flags))) != 0
	}
	return false
}

func (ptr *QPixmap) Rect() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QPixmap_Rect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QPixmap) Save(fileName string, format string, quality int) bool {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QPixmap_Save(ptr.Pointer(), C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, formatC, C.int(int32(quality)))) != 0
	}
	return false
}

func (ptr *QPixmap) Save2(device core.QIODevice_ITF, format string, quality int) bool {
	if ptr.Pointer() != nil {
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QPixmap_Save2(ptr.Pointer(), core.PointerFromQIODevice(device), formatC, C.int(int32(quality)))) != 0
	}
	return false
}

func (ptr *QPixmap) Scaled(size core.QSize_ITF, aspectRatioMode core.Qt__AspectRatioMode, transformMode core.Qt__TransformationMode) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QPixmap_Scaled(ptr.Pointer(), core.PointerFromQSize(size), C.longlong(aspectRatioMode), C.longlong(transformMode)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QPixmap) Scaled2(width int, height int, aspectRatioMode core.Qt__AspectRatioMode, transformMode core.Qt__TransformationMode) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QPixmap_Scaled2(ptr.Pointer(), C.int(int32(width)), C.int(int32(height)), C.longlong(aspectRatioMode), C.longlong(transformMode)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QPixmap) ScaledToHeight(height int, mode core.Qt__TransformationMode) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QPixmap_ScaledToHeight(ptr.Pointer(), C.int(int32(height)), C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QPixmap) Scroll(dx int, dy int, x int, y int, width int, height int, exposed QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_Scroll(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)), C.int(int32(x)), C.int(int32(y)), C.int(int32(width)), C.int(int32(height)), PointerFromQRegion(exposed))
	}
}

func (ptr *QPixmap) Scroll2(dx int, dy int, rect core.QRect_ITF, exposed QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_Scroll2(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)), core.PointerFromQRect(rect), PointerFromQRegion(exposed))
	}
}

func (ptr *QPixmap) Size() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QPixmap_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPixmap) ToImage() *QImage {
	if ptr.Pointer() != nil {
		tmpValue := NewQImageFromPointer(C.QPixmap_ToImage(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QImage).DestroyQImage)
		return tmpValue
	}
	return nil
}

//export callbackQPixmap_DestroyQPixmap
func callbackQPixmap_DestroyQPixmap(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPixmap"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPixmapFromPointer(ptr).DestroyQPixmapDefault()
	}
}

func (ptr *QPixmap) ConnectDestroyQPixmap(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPixmap"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPixmap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPixmap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPixmap) DisconnectDestroyQPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPixmap")
	}
}

func (ptr *QPixmap) DestroyQPixmap() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPixmap_DestroyQPixmap(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPixmap) DestroyQPixmapDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPixmap_DestroyQPixmapDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPixmap) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QPixmap_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPixmap_PaintEngine
func callbackQPixmap_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return PointerFromQPaintEngine((*(*func() *QPaintEngine)(signal))())
	}

	return PointerFromQPaintEngine(NewQPixmapFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QPixmap) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPixmap_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPixmap) PaintEngineDefault() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPixmap_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

// QPlatformSurfaceEvent::SurfaceEventType
//
//go:generate stringer -type=QPlatformSurfaceEvent__SurfaceEventType
type QPlatformSurfaceEvent__SurfaceEventType int64

const (
	QPlatformSurfaceEvent__SurfaceCreated            QPlatformSurfaceEvent__SurfaceEventType = QPlatformSurfaceEvent__SurfaceEventType(0)
	QPlatformSurfaceEvent__SurfaceAboutToBeDestroyed QPlatformSurfaceEvent__SurfaceEventType = QPlatformSurfaceEvent__SurfaceEventType(1)
)

type QPolygon struct {
	core.QVector
}

type QPolygon_ITF interface {
	core.QVector_ITF
	QPolygon_PTR() *QPolygon
}

func (ptr *QPolygon) QPolygon_PTR() *QPolygon {
	return ptr
}

func (ptr *QPolygon) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector_PTR().Pointer()
	}
	return nil
}

func (ptr *QPolygon) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QVector_PTR().SetPointer(p)
	}
}

func PointerFromQPolygon(ptr QPolygon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPolygon_PTR().Pointer()
	}
	return nil
}

func NewQPolygonFromPointer(ptr unsafe.Pointer) (n *QPolygon) {
	n = new(QPolygon)
	n.SetPointer(ptr)
	return
}
func NewQPolygon() *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon())
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func NewQPolygon2(size int) *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon2(C.int(int32(size))))
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func NewQPolygon3(points []*core.QPoint) *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon3(func() unsafe.Pointer {
		tmpList := NewQPolygonFromPointer(NewQPolygonFromPointer(nil).__QPolygon_points_newList3())
		for _, v := range points {
			tmpList.__QPolygon_points_setList3(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func NewQPolygon5(rectangle core.QRect_ITF, closed bool) *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon5(core.PointerFromQRect(rectangle), C.char(int8(qt.GoBoolToInt(closed)))))
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func (ptr *QPolygon) BoundingRect() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QPolygon_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) Point(index int, x int, y int) {
	if ptr.Pointer() != nil {
		C.QPolygon_Point(ptr.Pointer(), C.int(int32(index)), C.int(int32(x)), C.int(int32(y)))
	}
}

func (ptr *QPolygon) Point2(index int) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QPolygon_Point2(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) SetPoint(index int, x int, y int) {
	if ptr.Pointer() != nil {
		C.QPolygon_SetPoint(ptr.Pointer(), C.int(int32(index)), C.int(int32(x)), C.int(int32(y)))
	}
}

func (ptr *QPolygon) SetPoint2(index int, point core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon_SetPoint2(ptr.Pointer(), C.int(int32(index)), core.PointerFromQPoint(point))
	}
}

func (ptr *QPolygon) DestroyQPolygon() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPolygon_DestroyQPolygon(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPolygon) __QPolygon_points_atList3(i int) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QPolygon___QPolygon_points_atList3(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __QPolygon_points_setList3(i core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___QPolygon_points_setList3(ptr.Pointer(), core.PointerFromQPoint(i))
	}
}

func (ptr *QPolygon) __QPolygon_points_newList3() unsafe.Pointer {
	return C.QPolygon___QPolygon_points_newList3(ptr.Pointer())
}

func (ptr *QPolygon) __QPolygon_v_atList4(i int) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QPolygon___QPolygon_v_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __QPolygon_v_setList4(i core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___QPolygon_v_setList4(ptr.Pointer(), core.PointerFromQPoint(i))
	}
}

func (ptr *QPolygon) __QPolygon_v_newList4() unsafe.Pointer {
	return C.QPolygon___QPolygon_v_newList4(ptr.Pointer())
}

func (ptr *QPolygon) __QVector_other_atList4(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___QVector_other_atList4(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __QVector_other_setList4(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___QVector_other_setList4(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __QVector_other_newList4() unsafe.Pointer {
	return C.QPolygon___QVector_other_newList4(ptr.Pointer())
}

func (ptr *QPolygon) __QVector_other_atList5(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___QVector_other_atList5(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __QVector_other_setList5(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___QVector_other_setList5(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __QVector_other_newList5() unsafe.Pointer {
	return C.QPolygon___QVector_other_newList5(ptr.Pointer())
}

func (ptr *QPolygon) __append_value_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___append_value_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __append_value_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___append_value_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __append_value_newList3() unsafe.Pointer {
	return C.QPolygon___append_value_newList3(ptr.Pointer())
}

func (ptr *QPolygon) __fill_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___fill_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __fill_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___fill_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __fill_newList() unsafe.Pointer {
	return C.QPolygon___fill_newList(ptr.Pointer())
}

func (ptr *QPolygon) __fromList_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___fromList_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __fromList_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___fromList_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __fromList_newList() unsafe.Pointer {
	return C.QPolygon___fromList_newList(ptr.Pointer())
}

func (ptr *QPolygon) __fromList_list_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___fromList_list_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __fromList_list_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___fromList_list_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __fromList_list_newList() unsafe.Pointer {
	return C.QPolygon___fromList_list_newList(ptr.Pointer())
}

func (ptr *QPolygon) __fromStdVector_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___fromStdVector_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __fromStdVector_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___fromStdVector_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __fromStdVector_newList() unsafe.Pointer {
	return C.QPolygon___fromStdVector_newList(ptr.Pointer())
}

func (ptr *QPolygon) __mid_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___mid_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __mid_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___mid_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __mid_newList() unsafe.Pointer {
	return C.QPolygon___mid_newList(ptr.Pointer())
}

func (ptr *QPolygon) __swap_other_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___swap_other_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __swap_other_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___swap_other_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __swap_other_newList() unsafe.Pointer {
	return C.QPolygon___swap_other_newList(ptr.Pointer())
}

func (ptr *QPolygon) __toList_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___toList_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __toList_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___toList_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __toList_newList() unsafe.Pointer {
	return C.QPolygon___toList_newList(ptr.Pointer())
}

type QPolygonF struct {
	core.QVector
}

type QPolygonF_ITF interface {
	core.QVector_ITF
	QPolygonF_PTR() *QPolygonF
}

func (ptr *QPolygonF) QPolygonF_PTR() *QPolygonF {
	return ptr
}

func (ptr *QPolygonF) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector_PTR().Pointer()
	}
	return nil
}

func (ptr *QPolygonF) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QVector_PTR().SetPointer(p)
	}
}

func PointerFromQPolygonF(ptr QPolygonF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPolygonF_PTR().Pointer()
	}
	return nil
}

func NewQPolygonFFromPointer(ptr unsafe.Pointer) (n *QPolygonF) {
	n = new(QPolygonF)
	n.SetPointer(ptr)
	return
}
func NewQPolygonF() *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF())
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF2(size int) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF2(C.int(int32(size))))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF3(points []*core.QPointF) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF3(func() unsafe.Pointer {
		tmpList := NewQPolygonFFromPointer(NewQPolygonFFromPointer(nil).__QPolygonF_points_newList3())
		for _, v := range points {
			tmpList.__QPolygonF_points_setList3(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF5(rectangle core.QRectF_ITF) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF5(core.PointerFromQRectF(rectangle)))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF6(polygon QPolygon_ITF) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF6(PointerFromQPolygon(polygon)))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF7(polygon QPolygonF_ITF) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF7(PointerFromQPolygonF(polygon)))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func (ptr *QPolygonF) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QPolygonF_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) DestroyQPolygonF() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPolygonF_DestroyQPolygonF(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPolygonF) __QPolygonF_points_atList3(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QPolygonF___QPolygonF_points_atList3(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __QPolygonF_points_setList3(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___QPolygonF_points_setList3(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QPolygonF) __QPolygonF_points_newList3() unsafe.Pointer {
	return C.QPolygonF___QPolygonF_points_newList3(ptr.Pointer())
}

func (ptr *QPolygonF) __QPolygonF_v_atList4(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QPolygonF___QPolygonF_v_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __QPolygonF_v_setList4(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___QPolygonF_v_setList4(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QPolygonF) __QPolygonF_v_newList4() unsafe.Pointer {
	return C.QPolygonF___QPolygonF_v_newList4(ptr.Pointer())
}

func (ptr *QPolygonF) __QVector_other_atList4(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___QVector_other_atList4(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __QVector_other_setList4(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___QVector_other_setList4(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __QVector_other_newList4() unsafe.Pointer {
	return C.QPolygonF___QVector_other_newList4(ptr.Pointer())
}

func (ptr *QPolygonF) __QVector_other_atList5(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___QVector_other_atList5(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __QVector_other_setList5(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___QVector_other_setList5(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __QVector_other_newList5() unsafe.Pointer {
	return C.QPolygonF___QVector_other_newList5(ptr.Pointer())
}

func (ptr *QPolygonF) __append_value_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___append_value_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __append_value_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___append_value_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __append_value_newList3() unsafe.Pointer {
	return C.QPolygonF___append_value_newList3(ptr.Pointer())
}

func (ptr *QPolygonF) __fill_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___fill_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __fill_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___fill_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __fill_newList() unsafe.Pointer {
	return C.QPolygonF___fill_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __fromList_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___fromList_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __fromList_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___fromList_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __fromList_newList() unsafe.Pointer {
	return C.QPolygonF___fromList_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __fromList_list_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___fromList_list_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __fromList_list_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___fromList_list_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __fromList_list_newList() unsafe.Pointer {
	return C.QPolygonF___fromList_list_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __fromStdVector_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___fromStdVector_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __fromStdVector_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___fromStdVector_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __fromStdVector_newList() unsafe.Pointer {
	return C.QPolygonF___fromStdVector_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __mid_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___mid_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __mid_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___mid_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __mid_newList() unsafe.Pointer {
	return C.QPolygonF___mid_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __swap_other_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___swap_other_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __swap_other_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___swap_other_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __swap_other_newList() unsafe.Pointer {
	return C.QPolygonF___swap_other_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __toList_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___toList_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __toList_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___toList_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __toList_newList() unsafe.Pointer {
	return C.QPolygonF___toList_newList(ptr.Pointer())
}

// QRawFont::AntialiasingType
//
//go:generate stringer -type=QRawFont__AntialiasingType
type QRawFont__AntialiasingType int64

const (
	QRawFont__PixelAntialiasing    QRawFont__AntialiasingType = QRawFont__AntialiasingType(0)
	QRawFont__SubPixelAntialiasing QRawFont__AntialiasingType = QRawFont__AntialiasingType(1)
)

// QRawFont::LayoutFlag
//
//go:generate stringer -type=QRawFont__LayoutFlag
type QRawFont__LayoutFlag int64

const (
	QRawFont__SeparateAdvances QRawFont__LayoutFlag = QRawFont__LayoutFlag(0)
	QRawFont__KernedAdvances   QRawFont__LayoutFlag = QRawFont__LayoutFlag(1)
	QRawFont__UseDesignMetrics QRawFont__LayoutFlag = QRawFont__LayoutFlag(2)
)

type QRegion struct {
	ptr unsafe.Pointer
}

type QRegion_ITF interface {
	QRegion_PTR() *QRegion
}

func (ptr *QRegion) QRegion_PTR() *QRegion {
	return ptr
}

func (ptr *QRegion) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRegion) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRegion(ptr QRegion_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegion_PTR().Pointer()
	}
	return nil
}

func NewQRegionFromPointer(ptr unsafe.Pointer) (n *QRegion) {
	n = new(QRegion)
	n.SetPointer(ptr)
	return
}
func (ptr *QRegion) DestroyQRegion() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QRegion__RegionType
//QRegion::RegionType
type QRegion__RegionType int64

const (
	QRegion__Rectangle QRegion__RegionType = QRegion__RegionType(0)
	QRegion__Ellipse   QRegion__RegionType = QRegion__RegionType(1)
)

func NewQRegion() *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion())
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion2(x int, y int, w int, h int, t QRegion__RegionType) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion2(C.int(int32(x)), C.int(int32(y)), C.int(int32(w)), C.int(int32(h)), C.longlong(t)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion3(r core.QRect_ITF, t QRegion__RegionType) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion3(core.PointerFromQRect(r), C.longlong(t)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion4(a QPolygon_ITF, fillRule core.Qt__FillRule) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion4(PointerFromQPolygon(a), C.longlong(fillRule)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion5(r QRegion_ITF) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion5(PointerFromQRegion(r)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion6(other QRegion_ITF) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion6(PointerFromQRegion(other)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion7(bm QBitmap_ITF) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion7(PointerFromQBitmap(bm)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func (ptr *QRegion) BoundingRect() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QRegion_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QRegion) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRegion_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRegion) __rects_atList(i int) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QRegion___rects_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QRegion) __rects_setList(i core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QRegion___rects_setList(ptr.Pointer(), core.PointerFromQRect(i))
	}
}

func (ptr *QRegion) __rects_newList() unsafe.Pointer {
	return C.QRegion___rects_newList(ptr.Pointer())
}

type QResizeEvent struct {
	core.QEvent
}

type QResizeEvent_ITF interface {
	core.QEvent_ITF
	QResizeEvent_PTR() *QResizeEvent
}

func (ptr *QResizeEvent) QResizeEvent_PTR() *QResizeEvent {
	return ptr
}

func (ptr *QResizeEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QResizeEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQResizeEvent(ptr QResizeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QResizeEvent_PTR().Pointer()
	}
	return nil
}

func NewQResizeEventFromPointer(ptr unsafe.Pointer) (n *QResizeEvent) {
	n = new(QResizeEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QResizeEvent) DestroyQResizeEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQResizeEvent(size core.QSize_ITF, oldSize core.QSize_ITF) *QResizeEvent {
	tmpValue := NewQResizeEventFromPointer(C.QResizeEvent_NewQResizeEvent(core.PointerFromQSize(size), core.PointerFromQSize(oldSize)))
	qt.SetFinalizer(tmpValue, (*QResizeEvent).DestroyQResizeEvent)
	return tmpValue
}

func (ptr *QResizeEvent) Size() *core.QSize {
	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QResizeEvent_Size(ptr.Pointer()))
	}
	return nil
}

type QRgba64 struct {
	ptr unsafe.Pointer
}

type QRgba64_ITF interface {
	QRgba64_PTR() *QRgba64
}

func (ptr *QRgba64) QRgba64_PTR() *QRgba64 {
	return ptr
}

func (ptr *QRgba64) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRgba64) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRgba64(ptr QRgba64_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRgba64_PTR().Pointer()
	}
	return nil
}

func NewQRgba64FromPointer(ptr unsafe.Pointer) (n *QRgba64) {
	n = new(QRgba64)
	n.SetPointer(ptr)
	return
}
func (ptr *QRgba64) DestroyQRgba64() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func (ptr *QRgba64) Alpha() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QRgba64_Alpha(ptr.Pointer()))
	}
	return 0
}

type QScreen struct {
	core.QObject
}

type QScreen_ITF interface {
	core.QObject_ITF
	QScreen_PTR() *QScreen
}

func (ptr *QScreen) QScreen_PTR() *QScreen {
	return ptr
}

func (ptr *QScreen) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScreen) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScreen(ptr QScreen_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScreen_PTR().Pointer()
	}
	return nil
}

func NewQScreenFromPointer(ptr unsafe.Pointer) (n *QScreen) {
	n = new(QScreen)
	n.SetPointer(ptr)
	return
}
func (ptr *QScreen) Geometry() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QScreen_Geometry(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) Model() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScreen_Model(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScreen) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScreen_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScreen) Orientation() core.Qt__ScreenOrientation {
	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) Size() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QScreen_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQScreen_DestroyQScreen
func callbackQScreen_DestroyQScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScreenFromPointer(ptr).DestroyQScreenDefault()
	}
}

func (ptr *QScreen) ConnectDestroyQScreen(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScreen"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QScreen", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScreen", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScreen) DisconnectDestroyQScreen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QScreen")
	}
}

func (ptr *QScreen) DestroyQScreen() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScreen_DestroyQScreen(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScreen) DestroyQScreenDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScreen_DestroyQScreenDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScreen) __virtualSiblings_atList(i int) *QScreen {
	if ptr.Pointer() != nil {
		tmpValue := NewQScreenFromPointer(C.QScreen___virtualSiblings_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __virtualSiblings_setList(i QScreen_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___virtualSiblings_setList(ptr.Pointer(), PointerFromQScreen(i))
	}
}

func (ptr *QScreen) __virtualSiblings_newList() unsafe.Pointer {
	return C.QScreen___virtualSiblings_newList(ptr.Pointer())
}

func (ptr *QScreen) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScreen___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScreen) __children_newList() unsafe.Pointer {
	return C.QScreen___children_newList(ptr.Pointer())
}

func (ptr *QScreen) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QScreen___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScreen) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QScreen___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QScreen) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScreen___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScreen) __findChildren_newList() unsafe.Pointer {
	return C.QScreen___findChildren_newList(ptr.Pointer())
}

func (ptr *QScreen) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScreen___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScreen) __findChildren_newList3() unsafe.Pointer {
	return C.QScreen___findChildren_newList3(ptr.Pointer())
}

//export callbackQScreen_ChildEvent
func callbackQScreen_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScreenFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScreen) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScreen_ConnectNotify
func callbackQScreen_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScreenFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScreen) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScreen_CustomEvent
func callbackQScreen_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQScreenFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScreen) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScreen_DeleteLater
func callbackQScreen_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScreenFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScreen) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScreen_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQScreen_Destroyed
func callbackQScreen_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScreen_DisconnectNotify
func callbackQScreen_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScreenFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScreen) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScreen_Event
func callbackQScreen_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScreenFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScreen) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScreen_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQScreen_EventFilter
func callbackQScreen_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScreenFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScreen) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScreen_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQScreen_ObjectNameChanged
func callbackQScreen_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGui_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQScreen_TimerEvent
func callbackQScreen_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScreenFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScreen) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

// QScrollEvent::ScrollState
//
//go:generate stringer -type=QScrollEvent__ScrollState
type QScrollEvent__ScrollState int64

const (
	QScrollEvent__ScrollStarted  QScrollEvent__ScrollState = QScrollEvent__ScrollState(0)
	QScrollEvent__ScrollUpdated  QScrollEvent__ScrollState = QScrollEvent__ScrollState(1)
	QScrollEvent__ScrollFinished QScrollEvent__ScrollState = QScrollEvent__ScrollState(2)
)

// QSessionManager::RestartHint
//
//go:generate stringer -type=QSessionManager__RestartHint
type QSessionManager__RestartHint int64

const (
	QSessionManager__RestartIfRunning   QSessionManager__RestartHint = QSessionManager__RestartHint(0)
	QSessionManager__RestartAnyway      QSessionManager__RestartHint = QSessionManager__RestartHint(1)
	QSessionManager__RestartImmediately QSessionManager__RestartHint = QSessionManager__RestartHint(2)
	QSessionManager__RestartNever       QSessionManager__RestartHint = QSessionManager__RestartHint(3)
)

type QShowEvent struct {
	core.QEvent
}

type QShowEvent_ITF interface {
	core.QEvent_ITF
	QShowEvent_PTR() *QShowEvent
}

func (ptr *QShowEvent) QShowEvent_PTR() *QShowEvent {
	return ptr
}

func (ptr *QShowEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QShowEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQShowEvent(ptr QShowEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QShowEvent_PTR().Pointer()
	}
	return nil
}

func NewQShowEventFromPointer(ptr unsafe.Pointer) (n *QShowEvent) {
	n = new(QShowEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QShowEvent) DestroyQShowEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQShowEvent() *QShowEvent {
	tmpValue := NewQShowEventFromPointer(C.QShowEvent_NewQShowEvent())
	qt.SetFinalizer(tmpValue, (*QShowEvent).DestroyQShowEvent)
	return tmpValue
}

type QStandardItem struct {
	ptr unsafe.Pointer
}

type QStandardItem_ITF interface {
	QStandardItem_PTR() *QStandardItem
}

func (ptr *QStandardItem) QStandardItem_PTR() *QStandardItem {
	return ptr
}

func (ptr *QStandardItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QStandardItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQStandardItem(ptr QStandardItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStandardItem_PTR().Pointer()
	}
	return nil
}

func NewQStandardItemFromPointer(ptr unsafe.Pointer) (n *QStandardItem) {
	n = new(QStandardItem)
	n.SetPointer(ptr)
	return
}

// QStandardItem::ItemType
//
//go:generate stringer -type=QStandardItem__ItemType
type QStandardItem__ItemType int64

const (
	QStandardItem__Type     QStandardItem__ItemType = QStandardItem__ItemType(0)
	QStandardItem__UserType QStandardItem__ItemType = QStandardItem__ItemType(1000)
)

func NewQStandardItem() *QStandardItem {
	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem())
}

func NewQStandardItem2(text string) *QStandardItem {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem2(C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}))
}

func NewQStandardItem3(icon QIcon_ITF, text string) *QStandardItem {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem3(PointerFromQIcon(icon), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}))
}

func NewQStandardItem4(rows int, columns int) *QStandardItem {
	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem4(C.int(int32(rows)), C.int(int32(columns))))
}

func NewQStandardItem5(other QStandardItem_ITF) *QStandardItem {
	return NewQStandardItemFromPointer(C.QStandardItem_NewQStandardItem5(PointerFromQStandardItem(other)))
}

func (ptr *QStandardItem) AccessibleDescription() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QStandardItem_AccessibleDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) AccessibleText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QStandardItem_AccessibleText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) AppendColumn(items []*QStandardItem) {
	if ptr.Pointer() != nil {
		C.QStandardItem_AppendColumn(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQStandardItemFromPointer(NewQStandardItemFromPointer(nil).__appendColumn_items_newList())
			for _, v := range items {
				tmpList.__appendColumn_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QStandardItem) AppendRow(items []*QStandardItem) {
	if ptr.Pointer() != nil {
		C.QStandardItem_AppendRow(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQStandardItemFromPointer(NewQStandardItemFromPointer(nil).__appendRow_items_newList())
			for _, v := range items {
				tmpList.__appendRow_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QStandardItem) AppendRow2(item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_AppendRow2(ptr.Pointer(), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) AppendRows(items []*QStandardItem) {
	if ptr.Pointer() != nil {
		C.QStandardItem_AppendRows(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQStandardItemFromPointer(NewQStandardItemFromPointer(nil).__appendRows_items_newList())
			for _, v := range items {
				tmpList.__appendRows_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QStandardItem) Background() *QBrush {
	if ptr.Pointer() != nil {
		tmpValue := NewQBrushFromPointer(C.QStandardItem_Background(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItem) Child(row int, column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Child(ptr.Pointer(), C.int(int32(row)), C.int(int32(column))))
	}
	return nil
}

func (ptr *QStandardItem) Column() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItem_Column(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItem_ColumnCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQStandardItem_Data
func callbackQStandardItem_Data(ptr unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQVariant((*(*func(int) *core.QVariant)(signal))(int(int32(role))))
	}

	return core.PointerFromQVariant(NewQStandardItemFromPointer(ptr).DataDefault(int(int32(role))))
}

func (ptr *QStandardItem) ConnectData(f func(role int) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			f := func(role int) *core.QVariant {
				(*(*func(int) *core.QVariant)(signal))(role)
				return f(role)
			}
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItem) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *QStandardItem) Data(role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QStandardItem_Data(ptr.Pointer(), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItem) DataDefault(role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QStandardItem_DataDefault(ptr.Pointer(), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItem) Font() *QFont {
	if ptr.Pointer() != nil {
		tmpValue := NewQFontFromPointer(C.QStandardItem_Font(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItem) Foreground() *QBrush {
	if ptr.Pointer() != nil {
		tmpValue := NewQBrushFromPointer(C.QStandardItem_Foreground(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItem) HasChildren() bool {
	if ptr.Pointer() != nil {
		return int8(C.QStandardItem_HasChildren(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStandardItem) Icon() *QIcon {
	if ptr.Pointer() != nil {
		tmpValue := NewQIconFromPointer(C.QStandardItem_Icon(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItem) Index() *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QStandardItem_Index(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItem) InsertColumn(column int, items []*QStandardItem) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertColumn(ptr.Pointer(), C.int(int32(column)), func() unsafe.Pointer {
			tmpList := NewQStandardItemFromPointer(NewQStandardItemFromPointer(nil).__insertColumn_items_newList())
			for _, v := range items {
				tmpList.__insertColumn_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QStandardItem) InsertColumns(column int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)))
	}
}

func (ptr *QStandardItem) InsertRow(row int, items []*QStandardItem) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRow(ptr.Pointer(), C.int(int32(row)), func() unsafe.Pointer {
			tmpList := NewQStandardItemFromPointer(NewQStandardItemFromPointer(nil).__insertRow_items_newList())
			for _, v := range items {
				tmpList.__insertRow_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QStandardItem) InsertRow2(row int, item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRow2(ptr.Pointer(), C.int(int32(row)), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) InsertRows(row int, items []*QStandardItem) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRows(ptr.Pointer(), C.int(int32(row)), func() unsafe.Pointer {
			tmpList := NewQStandardItemFromPointer(NewQStandardItemFromPointer(nil).__insertRows_items_newList())
			for _, v := range items {
				tmpList.__insertRows_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QStandardItem) InsertRows2(row int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_InsertRows2(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)))
	}
}

func (ptr *QStandardItem) Model() *QStandardItemModel {
	if ptr.Pointer() != nil {
		tmpValue := NewQStandardItemModelFromPointer(C.QStandardItem_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItem) Parent() *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_Parent(ptr.Pointer()))
	}
	return nil
}

//export callbackQStandardItem_Read
func callbackQStandardItem_Read(ptr unsafe.Pointer, in unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "read"); signal != nil {
		(*(*func(*core.QDataStream))(signal))(core.NewQDataStreamFromPointer(in))
	} else {
		NewQStandardItemFromPointer(ptr).ReadDefault(core.NewQDataStreamFromPointer(in))
	}
}

func (ptr *QStandardItem) ConnectRead(f func(in *core.QDataStream)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "read"); signal != nil {
			f := func(in *core.QDataStream) {
				(*(*func(*core.QDataStream))(signal))(in)
				f(in)
			}
			qt.ConnectSignal(ptr.Pointer(), "read", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "read", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItem) DisconnectRead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "read")
	}
}

func (ptr *QStandardItem) Read(in core.QDataStream_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_Read(ptr.Pointer(), core.PointerFromQDataStream(in))
	}
}

func (ptr *QStandardItem) ReadDefault(in core.QDataStream_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_ReadDefault(ptr.Pointer(), core.PointerFromQDataStream(in))
	}
}

func (ptr *QStandardItem) RemoveColumn(column int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveColumn(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QStandardItem) RemoveColumns(column int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)))
	}
}

func (ptr *QStandardItem) RemoveRow(row int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveRow(ptr.Pointer(), C.int(int32(row)))
	}
}

func (ptr *QStandardItem) RemoveRows(row int, count int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_RemoveRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)))
	}
}

func (ptr *QStandardItem) Row() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItem_Row(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) RowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItem_RowCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) SetAccessibleDescription(accessibleDescription string) {
	if ptr.Pointer() != nil {
		var accessibleDescriptionC *C.char
		if accessibleDescription != "" {
			accessibleDescriptionC = C.CString(accessibleDescription)
			defer C.free(unsafe.Pointer(accessibleDescriptionC))
		}
		C.QStandardItem_SetAccessibleDescription(ptr.Pointer(), C.struct_QtGui_PackedString{data: accessibleDescriptionC, len: C.longlong(len(accessibleDescription))})
	}
}

func (ptr *QStandardItem) SetAccessibleText(accessibleText string) {
	if ptr.Pointer() != nil {
		var accessibleTextC *C.char
		if accessibleText != "" {
			accessibleTextC = C.CString(accessibleText)
			defer C.free(unsafe.Pointer(accessibleTextC))
		}
		C.QStandardItem_SetAccessibleText(ptr.Pointer(), C.struct_QtGui_PackedString{data: accessibleTextC, len: C.longlong(len(accessibleText))})
	}
}

func (ptr *QStandardItem) SetBackground(brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetBackground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QStandardItem) SetCheckable(checkable bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetCheckable(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(checkable))))
	}
}

func (ptr *QStandardItem) SetChild(row int, column int, item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetChild(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) SetChild2(row int, item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetChild2(ptr.Pointer(), C.int(int32(row)), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItem) SetColumnCount(columns int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetColumnCount(ptr.Pointer(), C.int(int32(columns)))
	}
}

//export callbackQStandardItem_SetData
func callbackQStandardItem_SetData(ptr unsafe.Pointer, value unsafe.Pointer, role C.int) {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		(*(*func(*core.QVariant, int))(signal))(core.NewQVariantFromPointer(value), int(int32(role)))
	} else {
		NewQStandardItemFromPointer(ptr).SetDataDefault(core.NewQVariantFromPointer(value), int(int32(role)))
	}
}

func (ptr *QStandardItem) ConnectSetData(f func(value *core.QVariant, role int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setData"); signal != nil {
			f := func(value *core.QVariant, role int) {
				(*(*func(*core.QVariant, int))(signal))(value, role)
				f(value, role)
			}
			qt.ConnectSignal(ptr.Pointer(), "setData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItem) DisconnectSetData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setData")
	}
}

func (ptr *QStandardItem) SetData(value core.QVariant_ITF, role int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetData(ptr.Pointer(), core.PointerFromQVariant(value), C.int(int32(role)))
	}
}

func (ptr *QStandardItem) SetDataDefault(value core.QVariant_ITF, role int) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetDataDefault(ptr.Pointer(), core.PointerFromQVariant(value), C.int(int32(role)))
	}
}

func (ptr *QStandardItem) SetEditable(editable bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetEditable(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(editable))))
	}
}

func (ptr *QStandardItem) SetEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QStandardItem) SetFont(font QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetFont(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QStandardItem) SetForeground(brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetForeground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QStandardItem) SetIcon(icon QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_SetIcon(ptr.Pointer(), PointerFromQIcon(icon))
	}
}

func (ptr *QStandardItem) SetText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QStandardItem_SetText(ptr.Pointer(), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QStandardItem) SetToolTip(toolTip string) {
	if ptr.Pointer() != nil {
		var toolTipC *C.char
		if toolTip != "" {
			toolTipC = C.CString(toolTip)
			defer C.free(unsafe.Pointer(toolTipC))
		}
		C.QStandardItem_SetToolTip(ptr.Pointer(), C.struct_QtGui_PackedString{data: toolTipC, len: C.longlong(len(toolTip))})
	}
}

func (ptr *QStandardItem) SizeHint() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QStandardItem_SizeHint(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItem) TakeChild(row int, column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem_TakeChild(ptr.Pointer(), C.int(int32(row)), C.int(int32(column))))
	}
	return nil
}

func (ptr *QStandardItem) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QStandardItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStandardItem) ToolTip() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QStandardItem_ToolTip(ptr.Pointer()))
	}
	return ""
}

//export callbackQStandardItem_Type
func callbackQStandardItem_Type(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQStandardItemFromPointer(ptr).TypeDefault()))
}

func (ptr *QStandardItem) ConnectType(f func() int) {
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

func (ptr *QStandardItem) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QStandardItem) Type() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItem_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStandardItem) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItem_TypeDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQStandardItem_Write
func callbackQStandardItem_Write(ptr unsafe.Pointer, out unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "write"); signal != nil {
		(*(*func(*core.QDataStream))(signal))(core.NewQDataStreamFromPointer(out))
	} else {
		NewQStandardItemFromPointer(ptr).WriteDefault(core.NewQDataStreamFromPointer(out))
	}
}

func (ptr *QStandardItem) ConnectWrite(f func(out *core.QDataStream)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "write"); signal != nil {
			f := func(out *core.QDataStream) {
				(*(*func(*core.QDataStream))(signal))(out)
				f(out)
			}
			qt.ConnectSignal(ptr.Pointer(), "write", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "write", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItem) DisconnectWrite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "write")
	}
}

func (ptr *QStandardItem) Write(out core.QDataStream_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_Write(ptr.Pointer(), core.PointerFromQDataStream(out))
	}
}

func (ptr *QStandardItem) WriteDefault(out core.QDataStream_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem_WriteDefault(ptr.Pointer(), core.PointerFromQDataStream(out))
	}
}

//export callbackQStandardItem_DestroyQStandardItem
func callbackQStandardItem_DestroyQStandardItem(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QStandardItem"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQStandardItemFromPointer(ptr).DestroyQStandardItemDefault()
	}
}

func (ptr *QStandardItem) ConnectDestroyQStandardItem(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QStandardItem"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QStandardItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QStandardItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItem) DisconnectDestroyQStandardItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QStandardItem")
	}
}

func (ptr *QStandardItem) DestroyQStandardItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStandardItem_DestroyQStandardItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStandardItem) DestroyQStandardItemDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStandardItem_DestroyQStandardItemDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStandardItem) __appendColumn_items_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem___appendColumn_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItem) __appendColumn_items_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem___appendColumn_items_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItem) __appendColumn_items_newList() unsafe.Pointer {
	return C.QStandardItem___appendColumn_items_newList(ptr.Pointer())
}

func (ptr *QStandardItem) __appendRow_items_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem___appendRow_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItem) __appendRow_items_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem___appendRow_items_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItem) __appendRow_items_newList() unsafe.Pointer {
	return C.QStandardItem___appendRow_items_newList(ptr.Pointer())
}

func (ptr *QStandardItem) __appendRows_items_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem___appendRows_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItem) __appendRows_items_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem___appendRows_items_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItem) __appendRows_items_newList() unsafe.Pointer {
	return C.QStandardItem___appendRows_items_newList(ptr.Pointer())
}

func (ptr *QStandardItem) __insertColumn_items_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem___insertColumn_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItem) __insertColumn_items_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem___insertColumn_items_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItem) __insertColumn_items_newList() unsafe.Pointer {
	return C.QStandardItem___insertColumn_items_newList(ptr.Pointer())
}

func (ptr *QStandardItem) __insertRow_items_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem___insertRow_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItem) __insertRow_items_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem___insertRow_items_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItem) __insertRow_items_newList() unsafe.Pointer {
	return C.QStandardItem___insertRow_items_newList(ptr.Pointer())
}

func (ptr *QStandardItem) __insertRows_items_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem___insertRows_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItem) __insertRows_items_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem___insertRows_items_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItem) __insertRows_items_newList() unsafe.Pointer {
	return C.QStandardItem___insertRows_items_newList(ptr.Pointer())
}

func (ptr *QStandardItem) __takeColumn_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem___takeColumn_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItem) __takeColumn_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem___takeColumn_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItem) __takeColumn_newList() unsafe.Pointer {
	return C.QStandardItem___takeColumn_newList(ptr.Pointer())
}

func (ptr *QStandardItem) __takeRow_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItem___takeRow_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItem) __takeRow_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItem___takeRow_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItem) __takeRow_newList() unsafe.Pointer {
	return C.QStandardItem___takeRow_newList(ptr.Pointer())
}

type QStandardItemModel struct {
	core.QAbstractItemModel
}

type QStandardItemModel_ITF interface {
	core.QAbstractItemModel_ITF
	QStandardItemModel_PTR() *QStandardItemModel
}

func (ptr *QStandardItemModel) QStandardItemModel_PTR() *QStandardItemModel {
	return ptr
}

func (ptr *QStandardItemModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QStandardItemModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractItemModel_PTR().SetPointer(p)
	}
}

func PointerFromQStandardItemModel(ptr QStandardItemModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStandardItemModel_PTR().Pointer()
	}
	return nil
}

func NewQStandardItemModelFromPointer(ptr unsafe.Pointer) (n *QStandardItemModel) {
	n = new(QStandardItemModel)
	n.SetPointer(ptr)
	return
}
func NewQStandardItemModel(parent core.QObject_ITF) *QStandardItemModel {
	tmpValue := NewQStandardItemModelFromPointer(C.QStandardItemModel_NewQStandardItemModel(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQStandardItemModel2(rows int, columns int, parent core.QObject_ITF) *QStandardItemModel {
	tmpValue := NewQStandardItemModelFromPointer(C.QStandardItemModel_NewQStandardItemModel2(C.int(int32(rows)), C.int(int32(columns)), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QStandardItemModel) AppendColumn(items []*QStandardItem) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_AppendColumn(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQStandardItemModelFromPointer(NewQStandardItemModelFromPointer(nil).__appendColumn_items_newList())
			for _, v := range items {
				tmpList.__appendColumn_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QStandardItemModel) AppendRow(items []*QStandardItem) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_AppendRow(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQStandardItemModelFromPointer(NewQStandardItemModelFromPointer(nil).__appendRow_items_newList())
			for _, v := range items {
				tmpList.__appendRow_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QStandardItemModel) AppendRow2(item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_AppendRow2(ptr.Pointer(), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) Clear() {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_Clear(ptr.Pointer())
	}
}

//export callbackQStandardItemModel_ColumnCount
func callbackQStandardItemModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*core.QModelIndex) int)(signal))(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQStandardItemModelFromPointer(ptr).ColumnCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QStandardItemModel) ConnectColumnCount(f func(parent *core.QModelIndex) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "columnCount"); signal != nil {
			f := func(parent *core.QModelIndex) int {
				(*(*func(*core.QModelIndex) int)(signal))(parent)
				return f(parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "columnCount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "columnCount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItemModel) DisconnectColumnCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "columnCount")
	}
}

func (ptr *QStandardItemModel) ColumnCount(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QStandardItemModel) ColumnCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel_ColumnCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQStandardItemModel_Data
func callbackQStandardItemModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return core.PointerFromQVariant((*(*func(*core.QModelIndex, int) *core.QVariant)(signal))(core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return core.PointerFromQVariant(NewQStandardItemModelFromPointer(ptr).DataDefault(core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *QStandardItemModel) ConnectData(f func(index *core.QModelIndex, role int) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			f := func(index *core.QModelIndex, role int) *core.QVariant {
				(*(*func(*core.QModelIndex, int) *core.QVariant)(signal))(index, role)
				return f(index, role)
			}
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItemModel) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *QStandardItemModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QStandardItemModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) DataDefault(index core.QModelIndex_ITF, role int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QStandardItemModel_DataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQStandardItemModel_HasChildren
func callbackQStandardItemModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex) bool)(signal))(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQStandardItemModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QStandardItemModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QStandardItemModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQStandardItemModel_Index
func callbackQStandardItemModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(int, int, *core.QModelIndex) *core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQStandardItemModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QStandardItemModel) ConnectIndex(f func(row int, column int, parent *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "index"); signal != nil {
			f := func(row int, column int, parent *core.QModelIndex) *core.QModelIndex {
				(*(*func(int, int, *core.QModelIndex) *core.QModelIndex)(signal))(row, column, parent)
				return f(row, column, parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "index", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "index", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItemModel) DisconnectIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "index")
	}
}

func (ptr *QStandardItemModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QStandardItemModel_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QStandardItemModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) InsertColumn(column int, items []*QStandardItem) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_InsertColumn(ptr.Pointer(), C.int(int32(column)), func() unsafe.Pointer {
			tmpList := NewQStandardItemModelFromPointer(NewQStandardItemModelFromPointer(nil).__insertColumn_items_newList())
			for _, v := range items {
				tmpList.__insertColumn_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQStandardItemModel_InsertColumns
func callbackQStandardItemModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQStandardItemModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QStandardItemModel) InsertColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QStandardItemModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) InsertRow(row int, items []*QStandardItem) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_InsertRow(ptr.Pointer(), C.int(int32(row)), func() unsafe.Pointer {
			tmpList := NewQStandardItemModelFromPointer(NewQStandardItemModelFromPointer(nil).__insertRow_items_newList())
			for _, v := range items {
				tmpList.__insertRow_items_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QStandardItemModel) InsertRow2(row int, item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_InsertRow2(ptr.Pointer(), C.int(int32(row)), PointerFromQStandardItem(item))
	}
}

//export callbackQStandardItemModel_InsertRows
func callbackQStandardItemModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQStandardItemModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QStandardItemModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QStandardItemModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) Item(row int, column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_Item(ptr.Pointer(), C.int(int32(row)), C.int(int32(column))))
	}
	return nil
}

func (ptr *QStandardItemModel) ItemFromIndex(index core.QModelIndex_ITF) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_ItemFromIndex(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

//export callbackQStandardItemModel_Parent
func callbackQStandardItemModel_Parent(ptr unsafe.Pointer, child unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return core.PointerFromQModelIndex((*(*func(*core.QModelIndex) *core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(child)))
	}

	return core.PointerFromQModelIndex(NewQStandardItemModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(child)))
}

func (ptr *QStandardItemModel) ConnectParent(f func(child *core.QModelIndex) *core.QModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parent"); signal != nil {
			f := func(child *core.QModelIndex) *core.QModelIndex {
				(*(*func(*core.QModelIndex) *core.QModelIndex)(signal))(child)
				return f(child)
			}
			qt.ConnectSignal(ptr.Pointer(), "parent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItemModel) DisconnectParent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parent")
	}
}

func (ptr *QStandardItemModel) Parent(child core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QStandardItemModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(child)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) ParentDefault(child core.QModelIndex_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QStandardItemModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(child)))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQStandardItemModel_RemoveColumns
func callbackQStandardItemModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQStandardItemModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QStandardItemModel) RemoveColumnsDefault(column int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QStandardItemModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQStandardItemModel_RemoveRows
func callbackQStandardItemModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQStandardItemModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QStandardItemModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QStandardItemModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackQStandardItemModel_RowCount
func callbackQStandardItemModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*core.QModelIndex) int)(signal))(core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewQStandardItemModelFromPointer(ptr).RowCountDefault(core.NewQModelIndexFromPointer(parent))))
}

func (ptr *QStandardItemModel) ConnectRowCount(f func(parent *core.QModelIndex) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rowCount"); signal != nil {
			f := func(parent *core.QModelIndex) int {
				(*(*func(*core.QModelIndex) int)(signal))(parent)
				return f(parent)
			}
			qt.ConnectSignal(ptr.Pointer(), "rowCount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rowCount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItemModel) DisconnectRowCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rowCount")
	}
}

func (ptr *QStandardItemModel) RowCount(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QStandardItemModel) RowCountDefault(parent core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel_RowCountDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QStandardItemModel) SetColumnCount(columns int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetColumnCount(ptr.Pointer(), C.int(int32(columns)))
	}
}

//export callbackQStandardItemModel_SetData
func callbackQStandardItemModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QModelIndex, *core.QVariant, int) bool)(signal))(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQStandardItemModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QStandardItemModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QStandardItemModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

func (ptr *QStandardItemModel) SetHorizontalHeaderLabels(labels []string) {
	if ptr.Pointer() != nil {
		labelsC := C.CString(strings.Join(labels, "¡¦!"))
		defer C.free(unsafe.Pointer(labelsC))
		C.QStandardItemModel_SetHorizontalHeaderLabels(ptr.Pointer(), C.struct_QtGui_PackedString{data: labelsC, len: C.longlong(len(strings.Join(labels, "¡¦!")))})
	}
}

func (ptr *QStandardItemModel) SetItem(row int, column int, item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItem(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) SetItem2(row int, item QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_SetItem2(ptr.Pointer(), C.int(int32(row)), PointerFromQStandardItem(item))
	}
}

func (ptr *QStandardItemModel) TakeItem(row int, column int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel_TakeItem(ptr.Pointer(), C.int(int32(row)), C.int(int32(column))))
	}
	return nil
}

//export callbackQStandardItemModel_DestroyQStandardItemModel
func callbackQStandardItemModel_DestroyQStandardItemModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QStandardItemModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQStandardItemModelFromPointer(ptr).DestroyQStandardItemModelDefault()
	}
}

func (ptr *QStandardItemModel) ConnectDestroyQStandardItemModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QStandardItemModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QStandardItemModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QStandardItemModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QStandardItemModel) DisconnectDestroyQStandardItemModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QStandardItemModel")
	}
}

func (ptr *QStandardItemModel) DestroyQStandardItemModel() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStandardItemModel_DestroyQStandardItemModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStandardItemModel) DestroyQStandardItemModelDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStandardItemModel_DestroyQStandardItemModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QStandardItemModel) __appendColumn_items_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel___appendColumn_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItemModel) __appendColumn_items_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___appendColumn_items_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItemModel) __appendColumn_items_newList() unsafe.Pointer {
	return C.QStandardItemModel___appendColumn_items_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __appendRow_items_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel___appendRow_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItemModel) __appendRow_items_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___appendRow_items_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItemModel) __appendRow_items_newList() unsafe.Pointer {
	return C.QStandardItemModel___appendRow_items_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __findItems_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel___findItems_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItemModel) __findItems_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___findItems_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItemModel) __findItems_newList() unsafe.Pointer {
	return C.QStandardItemModel___findItems_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __insertColumn_items_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel___insertColumn_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItemModel) __insertColumn_items_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___insertColumn_items_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItemModel) __insertColumn_items_newList() unsafe.Pointer {
	return C.QStandardItemModel___insertColumn_items_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __insertRow_items_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel___insertRow_items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItemModel) __insertRow_items_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___insertRow_items_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItemModel) __insertRow_items_newList() unsafe.Pointer {
	return C.QStandardItemModel___insertRow_items_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __itemData_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QStandardItemModel___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __itemData_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___itemData_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QStandardItemModel) __itemData_newList() unsafe.Pointer {
	return C.QStandardItemModel___itemData_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtGui_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQStandardItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.QStandardItemModel___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QStandardItemModel) __mimeData_indexes_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QStandardItemModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __mimeData_indexes_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___mimeData_indexes_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QStandardItemModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.QStandardItemModel___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __setItemData_roles_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QStandardItemModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __setItemData_roles_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QStandardItemModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.QStandardItemModel___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtGui_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQStandardItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.QStandardItemModel___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QStandardItemModel) __setItemRoleNames_roleNames_atList(v int, i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QStandardItemModel___setItemRoleNames_roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __setItemRoleNames_roleNames_setList(key int, i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___setItemRoleNames_roleNames_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQByteArray(i))
	}
}

func (ptr *QStandardItemModel) __setItemRoleNames_roleNames_newList() unsafe.Pointer {
	return C.QStandardItemModel___setItemRoleNames_roleNames_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __setItemRoleNames_roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtGui_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQStandardItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemRoleNames_roleNames_keyList_atList(i)
			}
			return out
		}(C.QStandardItemModel___setItemRoleNames_roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QStandardItemModel) __takeColumn_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel___takeColumn_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItemModel) __takeColumn_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___takeColumn_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItemModel) __takeColumn_newList() unsafe.Pointer {
	return C.QStandardItemModel___takeColumn_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __takeRow_atList(i int) *QStandardItem {
	if ptr.Pointer() != nil {
		return NewQStandardItemFromPointer(C.QStandardItemModel___takeRow_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QStandardItemModel) __takeRow_setList(i QStandardItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___takeRow_setList(ptr.Pointer(), PointerFromQStandardItem(i))
	}
}

func (ptr *QStandardItemModel) __takeRow_newList() unsafe.Pointer {
	return C.QStandardItemModel___takeRow_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QStandardItemModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QStandardItemModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.QStandardItemModel_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QStandardItemModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QStandardItemModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.QStandardItemModel_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) ____setItemRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel_____setItemRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QStandardItemModel) ____setItemRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_____setItemRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QStandardItemModel) ____setItemRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QStandardItemModel_____setItemRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __changePersistentIndexList_from_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QStandardItemModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __changePersistentIndexList_from_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___changePersistentIndexList_from_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QStandardItemModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.QStandardItemModel___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __changePersistentIndexList_to_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QStandardItemModel___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __changePersistentIndexList_to_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___changePersistentIndexList_to_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QStandardItemModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.QStandardItemModel___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QStandardItemModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QStandardItemModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.QStandardItemModel___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __layoutAboutToBeChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QStandardItemModel___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __layoutAboutToBeChanged_parents_setList(i core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QStandardItemModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.QStandardItemModel___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __layoutChanged_parents_atList(i int) *core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPersistentModelIndexFromPointer(C.QStandardItemModel___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __layoutChanged_parents_setList(i core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___layoutChanged_parents_setList(ptr.Pointer(), core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *QStandardItemModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.QStandardItemModel___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __match_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QStandardItemModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __match_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___match_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QStandardItemModel) __match_newList() unsafe.Pointer {
	return C.QStandardItemModel___match_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __persistentIndexList_atList(i int) *core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQModelIndexFromPointer(C.QStandardItemModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __persistentIndexList_setList(i core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___persistentIndexList_setList(ptr.Pointer(), core.PointerFromQModelIndex(i))
	}
}

func (ptr *QStandardItemModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.QStandardItemModel___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __roleNames_atList(v int, i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QStandardItemModel___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __roleNames_setList(key int, i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___roleNames_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQByteArray(i))
	}
}

func (ptr *QStandardItemModel) __roleNames_newList() unsafe.Pointer {
	return C.QStandardItemModel___roleNames_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtGui_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQStandardItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.QStandardItemModel___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QStandardItemModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QStandardItemModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QStandardItemModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QStandardItemModel_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QStandardItemModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QStandardItemModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.QStandardItemModel_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QStandardItemModel_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QStandardItemModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QStandardItemModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.QStandardItemModel_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QStandardItemModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QStandardItemModel) __children_newList() unsafe.Pointer {
	return C.QStandardItemModel___children_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QStandardItemModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QStandardItemModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QStandardItemModel___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QStandardItemModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QStandardItemModel) __findChildren_newList() unsafe.Pointer {
	return C.QStandardItemModel___findChildren_newList(ptr.Pointer())
}

func (ptr *QStandardItemModel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QStandardItemModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QStandardItemModel) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QStandardItemModel) __findChildren_newList3() unsafe.Pointer {
	return C.QStandardItemModel___findChildren_newList3(ptr.Pointer())
}

//export callbackQStandardItemModel_Match
func callbackQStandardItemModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQStandardItemModelFromPointer(NewQStandardItemModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*core.QModelIndex, int, *core.QVariant, int, core.Qt__MatchFlag) []*core.QModelIndex)(signal))(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQStandardItemModelFromPointer(NewQStandardItemModelFromPointer(nil).__match_newList())
		for _, v := range NewQStandardItemModelFromPointer(ptr).MatchDefault(core.NewQModelIndexFromPointer(start), int(int32(role)), core.NewQVariantFromPointer(value), int(int32(hits)), core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QStandardItemModel) MatchDefault(start core.QModelIndex_ITF, role int, value core.QVariant_ITF, hits int, flags core.Qt__MatchFlag) []*core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtGui_PackedList) []*core.QModelIndex {
			out := make([]*core.QModelIndex, int(l.len))
			tmpList := NewQStandardItemModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.QStandardItemModel_MatchDefault(ptr.Pointer(), core.PointerFromQModelIndex(start), C.int(int32(role)), core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*core.QModelIndex, 0)
}

//export callbackQStandardItemModel_ChildEvent
func callbackQStandardItemModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQStandardItemModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QStandardItemModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQStandardItemModel_ConnectNotify
func callbackQStandardItemModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQStandardItemModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QStandardItemModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQStandardItemModel_CustomEvent
func callbackQStandardItemModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQStandardItemModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QStandardItemModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQStandardItemModel_DeleteLater
func callbackQStandardItemModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQStandardItemModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QStandardItemModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QStandardItemModel_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQStandardItemModel_Destroyed
func callbackQStandardItemModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQStandardItemModel_DisconnectNotify
func callbackQStandardItemModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQStandardItemModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QStandardItemModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQStandardItemModel_Event
func callbackQStandardItemModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQStandardItemModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QStandardItemModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QStandardItemModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQStandardItemModel_EventFilter
func callbackQStandardItemModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQStandardItemModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QStandardItemModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QStandardItemModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQStandardItemModel_ObjectNameChanged
func callbackQStandardItemModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGui_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQStandardItemModel_TimerEvent
func callbackQStandardItemModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQStandardItemModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QStandardItemModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QStandardItemModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

// QStaticText::PerformanceHint
//
//go:generate stringer -type=QStaticText__PerformanceHint
type QStaticText__PerformanceHint int64

const (
	QStaticText__ModerateCaching   QStaticText__PerformanceHint = QStaticText__PerformanceHint(0)
	QStaticText__AggressiveCaching QStaticText__PerformanceHint = QStaticText__PerformanceHint(1)
)

type QSurface struct {
	ptr unsafe.Pointer
}

type QSurface_ITF interface {
	QSurface_PTR() *QSurface
}

func (ptr *QSurface) QSurface_PTR() *QSurface {
	return ptr
}

func (ptr *QSurface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSurface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSurface(ptr QSurface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurface_PTR().Pointer()
	}
	return nil
}

func NewQSurfaceFromPointer(ptr unsafe.Pointer) (n *QSurface) {
	n = new(QSurface)
	n.SetPointer(ptr)
	return
}

// QSurface::SurfaceClass
//
//go:generate stringer -type=QSurface__SurfaceClass
type QSurface__SurfaceClass int64

const (
	QSurface__Window    QSurface__SurfaceClass = QSurface__SurfaceClass(0)
	QSurface__Offscreen QSurface__SurfaceClass = QSurface__SurfaceClass(1)
)

// QSurface::SurfaceType
//
//go:generate stringer -type=QSurface__SurfaceType
type QSurface__SurfaceType int64

const (
	QSurface__RasterSurface   QSurface__SurfaceType = QSurface__SurfaceType(0)
	QSurface__OpenGLSurface   QSurface__SurfaceType = QSurface__SurfaceType(1)
	QSurface__RasterGLSurface QSurface__SurfaceType = QSurface__SurfaceType(2)
	QSurface__OpenVGSurface   QSurface__SurfaceType = QSurface__SurfaceType(3)
	QSurface__VulkanSurface   QSurface__SurfaceType = QSurface__SurfaceType(4)
	QSurface__MetalSurface    QSurface__SurfaceType = QSurface__SurfaceType(5)
)

//export callbackQSurface_Format
func callbackQSurface_Format(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "format"); signal != nil {
		return PointerFromQSurfaceFormat((*(*func() *QSurfaceFormat)(signal))())
	}

	return PointerFromQSurfaceFormat(NewQSurfaceFormat())
}

func (ptr *QSurface) ConnectFormat(f func() *QSurfaceFormat) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "format"); signal != nil {
			f := func() *QSurfaceFormat {
				(*(*func() *QSurfaceFormat)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "format", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "format", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSurface) DisconnectFormat() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "format")
	}
}

func (ptr *QSurface) Format() *QSurfaceFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQSurfaceFormatFromPointer(C.QSurface_Format(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

//export callbackQSurface_Size
func callbackQSurface_Size(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(core.NewQSize())
}

func (ptr *QSurface) ConnectSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "size"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "size", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "size", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSurface) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "size")
	}
}

func (ptr *QSurface) Size() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QSurface_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSurface_SurfaceType
func callbackQSurface_SurfaceType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "surfaceType"); signal != nil {
		return C.longlong((*(*func() QSurface__SurfaceType)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QSurface) ConnectSurfaceType(f func() QSurface__SurfaceType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "surfaceType"); signal != nil {
			f := func() QSurface__SurfaceType {
				(*(*func() QSurface__SurfaceType)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "surfaceType", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "surfaceType", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSurface) DisconnectSurfaceType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "surfaceType")
	}
}

func (ptr *QSurface) SurfaceType() QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QSurface_SurfaceType(ptr.Pointer()))
	}
	return 0
}

//export callbackQSurface_DestroyQSurface
func callbackQSurface_DestroyQSurface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSurface"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSurfaceFromPointer(ptr).DestroyQSurfaceDefault()
	}
}

func (ptr *QSurface) ConnectDestroyQSurface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSurface"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSurface", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSurface", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSurface) DisconnectDestroyQSurface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSurface")
	}
}

func (ptr *QSurface) DestroyQSurface() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSurface_DestroyQSurface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSurface) DestroyQSurfaceDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSurface_DestroyQSurfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSurfaceFormat struct {
	ptr unsafe.Pointer
}

type QSurfaceFormat_ITF interface {
	QSurfaceFormat_PTR() *QSurfaceFormat
}

func (ptr *QSurfaceFormat) QSurfaceFormat_PTR() *QSurfaceFormat {
	return ptr
}

func (ptr *QSurfaceFormat) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSurfaceFormat) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSurfaceFormat(ptr QSurfaceFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurfaceFormat_PTR().Pointer()
	}
	return nil
}

func NewQSurfaceFormatFromPointer(ptr unsafe.Pointer) (n *QSurfaceFormat) {
	n = new(QSurfaceFormat)
	n.SetPointer(ptr)
	return
}

// QSurfaceFormat::ColorSpace
//
//go:generate stringer -type=QSurfaceFormat__ColorSpace
type QSurfaceFormat__ColorSpace int64

const (
	QSurfaceFormat__DefaultColorSpace QSurfaceFormat__ColorSpace = QSurfaceFormat__ColorSpace(0)
	QSurfaceFormat__sRGBColorSpace    QSurfaceFormat__ColorSpace = QSurfaceFormat__ColorSpace(1)
)

// QSurfaceFormat::FormatOption
//
//go:generate stringer -type=QSurfaceFormat__FormatOption
type QSurfaceFormat__FormatOption int64

const (
	QSurfaceFormat__StereoBuffers       QSurfaceFormat__FormatOption = QSurfaceFormat__FormatOption(0x0001)
	QSurfaceFormat__DebugContext        QSurfaceFormat__FormatOption = QSurfaceFormat__FormatOption(0x0002)
	QSurfaceFormat__DeprecatedFunctions QSurfaceFormat__FormatOption = QSurfaceFormat__FormatOption(0x0004)
	QSurfaceFormat__ResetNotification   QSurfaceFormat__FormatOption = QSurfaceFormat__FormatOption(0x0008)
)

// QSurfaceFormat::OpenGLContextProfile
//
//go:generate stringer -type=QSurfaceFormat__OpenGLContextProfile
type QSurfaceFormat__OpenGLContextProfile int64

const (
	QSurfaceFormat__NoProfile            QSurfaceFormat__OpenGLContextProfile = QSurfaceFormat__OpenGLContextProfile(0)
	QSurfaceFormat__CoreProfile          QSurfaceFormat__OpenGLContextProfile = QSurfaceFormat__OpenGLContextProfile(1)
	QSurfaceFormat__CompatibilityProfile QSurfaceFormat__OpenGLContextProfile = QSurfaceFormat__OpenGLContextProfile(2)
)

// QSurfaceFormat::RenderableType
//
//go:generate stringer -type=QSurfaceFormat__RenderableType
type QSurfaceFormat__RenderableType int64

const (
	QSurfaceFormat__DefaultRenderableType QSurfaceFormat__RenderableType = QSurfaceFormat__RenderableType(0x0)
	QSurfaceFormat__OpenGL                QSurfaceFormat__RenderableType = QSurfaceFormat__RenderableType(0x1)
	QSurfaceFormat__OpenGLES              QSurfaceFormat__RenderableType = QSurfaceFormat__RenderableType(0x2)
	QSurfaceFormat__OpenVG                QSurfaceFormat__RenderableType = QSurfaceFormat__RenderableType(0x4)
)

// QSurfaceFormat::SwapBehavior
//
//go:generate stringer -type=QSurfaceFormat__SwapBehavior
type QSurfaceFormat__SwapBehavior int64

const (
	QSurfaceFormat__DefaultSwapBehavior QSurfaceFormat__SwapBehavior = QSurfaceFormat__SwapBehavior(0)
	QSurfaceFormat__SingleBuffer        QSurfaceFormat__SwapBehavior = QSurfaceFormat__SwapBehavior(1)
	QSurfaceFormat__DoubleBuffer        QSurfaceFormat__SwapBehavior = QSurfaceFormat__SwapBehavior(2)
	QSurfaceFormat__TripleBuffer        QSurfaceFormat__SwapBehavior = QSurfaceFormat__SwapBehavior(3)
)

func NewQSurfaceFormat() *QSurfaceFormat {
	tmpValue := NewQSurfaceFormatFromPointer(C.QSurfaceFormat_NewQSurfaceFormat())
	qt.SetFinalizer(tmpValue, (*QSurfaceFormat).DestroyQSurfaceFormat)
	return tmpValue
}

func NewQSurfaceFormat2(options QSurfaceFormat__FormatOption) *QSurfaceFormat {
	tmpValue := NewQSurfaceFormatFromPointer(C.QSurfaceFormat_NewQSurfaceFormat2(C.longlong(options)))
	qt.SetFinalizer(tmpValue, (*QSurfaceFormat).DestroyQSurfaceFormat)
	return tmpValue
}

func NewQSurfaceFormat3(other QSurfaceFormat_ITF) *QSurfaceFormat {
	tmpValue := NewQSurfaceFormatFromPointer(C.QSurfaceFormat_NewQSurfaceFormat3(PointerFromQSurfaceFormat(other)))
	qt.SetFinalizer(tmpValue, (*QSurfaceFormat).DestroyQSurfaceFormat)
	return tmpValue
}

func (ptr *QSurfaceFormat) DestroyQSurfaceFormat() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSurfaceFormat_DestroyQSurfaceFormat(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

// QTabletEvent::PointerType
//
//go:generate stringer -type=QTabletEvent__PointerType
type QTabletEvent__PointerType int64

const (
	QTabletEvent__UnknownPointer QTabletEvent__PointerType = QTabletEvent__PointerType(0)
	QTabletEvent__Pen            QTabletEvent__PointerType = QTabletEvent__PointerType(1)
	QTabletEvent__Cursor         QTabletEvent__PointerType = QTabletEvent__PointerType(2)
	QTabletEvent__Eraser         QTabletEvent__PointerType = QTabletEvent__PointerType(3)
)

// QTabletEvent::TabletDevice
//
//go:generate stringer -type=QTabletEvent__TabletDevice
type QTabletEvent__TabletDevice int64

const (
	QTabletEvent__NoDevice       QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(0)
	QTabletEvent__Puck           QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(1)
	QTabletEvent__Stylus         QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(2)
	QTabletEvent__Airbrush       QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(3)
	QTabletEvent__FourDMouse     QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(4)
	QTabletEvent__XFreeEraser    QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(5)
	QTabletEvent__RotationStylus QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(6)
)

type QTextBlock struct {
	ptr unsafe.Pointer
}

type QTextBlock_ITF interface {
	QTextBlock_PTR() *QTextBlock
}

func (ptr *QTextBlock) QTextBlock_PTR() *QTextBlock {
	return ptr
}

func (ptr *QTextBlock) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextBlock) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextBlock(ptr QTextBlock_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlock_PTR().Pointer()
	}
	return nil
}

func NewQTextBlockFromPointer(ptr unsafe.Pointer) (n *QTextBlock) {
	n = new(QTextBlock)
	n.SetPointer(ptr)
	return
}
func (ptr *QTextBlock) DestroyQTextBlock() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQTextBlock3(other QTextBlock_ITF) *QTextBlock {
	tmpValue := NewQTextBlockFromPointer(C.QTextBlock_NewQTextBlock3(PointerFromQTextBlock(other)))
	qt.SetFinalizer(tmpValue, (*QTextBlock).DestroyQTextBlock)
	return tmpValue
}

func (ptr *QTextBlock) BlockFormat() *QTextBlockFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextBlockFormatFromPointer(C.QTextBlock_BlockFormat(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextBlockFormat).DestroyQTextBlockFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QTextBlock) CharFormat() *QTextCharFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextCharFormatFromPointer(C.QTextBlock_CharFormat(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextCharFormat).DestroyQTextCharFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QTextBlock) Document() *QTextDocument {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextDocumentFromPointer(C.QTextBlock_Document(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextBlock) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextBlock_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextBlock) Layout() *QTextLayout {
	if ptr.Pointer() != nil {
		return NewQTextLayoutFromPointer(C.QTextBlock_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextBlock) Length() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextBlock_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) Next() *QTextBlock {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextBlockFromPointer(C.QTextBlock_Next(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextBlock).DestroyQTextBlock)
		return tmpValue
	}
	return nil
}

func (ptr *QTextBlock) Position() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextBlock_Position(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlock) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextBlock_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextBlock) __textFormats_newList() unsafe.Pointer {
	return C.QTextBlock___textFormats_newList(ptr.Pointer())
}

type QTextBlockFormat struct {
	QTextFormat
}

type QTextBlockFormat_ITF interface {
	QTextFormat_ITF
	QTextBlockFormat_PTR() *QTextBlockFormat
}

func (ptr *QTextBlockFormat) QTextBlockFormat_PTR() *QTextBlockFormat {
	return ptr
}

func (ptr *QTextBlockFormat) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFormat_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextBlockFormat) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextFormat_PTR().SetPointer(p)
	}
}

func PointerFromQTextBlockFormat(ptr QTextBlockFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextBlockFormatFromPointer(ptr unsafe.Pointer) (n *QTextBlockFormat) {
	n = new(QTextBlockFormat)
	n.SetPointer(ptr)
	return
}
func (ptr *QTextBlockFormat) DestroyQTextBlockFormat() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QTextBlockFormat__LineHeightTypes
//QTextBlockFormat::LineHeightTypes
type QTextBlockFormat__LineHeightTypes int64

const (
	QTextBlockFormat__SingleHeight       QTextBlockFormat__LineHeightTypes = QTextBlockFormat__LineHeightTypes(0)
	QTextBlockFormat__ProportionalHeight QTextBlockFormat__LineHeightTypes = QTextBlockFormat__LineHeightTypes(1)
	QTextBlockFormat__FixedHeight        QTextBlockFormat__LineHeightTypes = QTextBlockFormat__LineHeightTypes(2)
	QTextBlockFormat__MinimumHeight      QTextBlockFormat__LineHeightTypes = QTextBlockFormat__LineHeightTypes(3)
	QTextBlockFormat__LineDistanceHeight QTextBlockFormat__LineHeightTypes = QTextBlockFormat__LineHeightTypes(4)
)

// QTextBlockFormat::MarkerType
//
//go:generate stringer -type=QTextBlockFormat__MarkerType
type QTextBlockFormat__MarkerType int64

const (
	QTextBlockFormat__NoMarker  QTextBlockFormat__MarkerType = QTextBlockFormat__MarkerType(0)
	QTextBlockFormat__Unchecked QTextBlockFormat__MarkerType = QTextBlockFormat__MarkerType(1)
	QTextBlockFormat__Checked   QTextBlockFormat__MarkerType = QTextBlockFormat__MarkerType(2)
)

func NewQTextBlockFormat() *QTextBlockFormat {
	tmpValue := NewQTextBlockFormatFromPointer(C.QTextBlockFormat_NewQTextBlockFormat())
	qt.SetFinalizer(tmpValue, (*QTextBlockFormat).DestroyQTextBlockFormat)
	return tmpValue
}

func (ptr *QTextBlockFormat) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextBlockFormat_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) Indent() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextBlockFormat_Indent(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBlockFormat) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetAlignment(ptr.Pointer(), C.longlong(alignment))
	}
}

func (ptr *QTextBlockFormat) SetIndent(indentation int) {
	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetIndent(ptr.Pointer(), C.int(int32(indentation)))
	}
}

func (ptr *QTextBlockFormat) SetTopMargin(margin float64) {
	if ptr.Pointer() != nil {
		C.QTextBlockFormat_SetTopMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextBlockFormat) TopMargin() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextBlockFormat_TopMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBlockFormat) __setTabPositions_tabs_newList() unsafe.Pointer {
	return C.QTextBlockFormat___setTabPositions_tabs_newList(ptr.Pointer())
}

func (ptr *QTextBlockFormat) __tabPositions_newList() unsafe.Pointer {
	return C.QTextBlockFormat___tabPositions_newList(ptr.Pointer())
}

type QTextBlockGroup struct {
	QTextObject
}

type QTextBlockGroup_ITF interface {
	QTextObject_ITF
	QTextBlockGroup_PTR() *QTextBlockGroup
}

func (ptr *QTextBlockGroup) QTextBlockGroup_PTR() *QTextBlockGroup {
	return ptr
}

func (ptr *QTextBlockGroup) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextBlockGroup) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextObject_PTR().SetPointer(p)
	}
}

func PointerFromQTextBlockGroup(ptr QTextBlockGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockGroup_PTR().Pointer()
	}
	return nil
}

func NewQTextBlockGroupFromPointer(ptr unsafe.Pointer) (n *QTextBlockGroup) {
	n = new(QTextBlockGroup)
	n.SetPointer(ptr)
	return
}
func NewQTextBlockGroup(document QTextDocument_ITF) *QTextBlockGroup {
	tmpValue := NewQTextBlockGroupFromPointer(C.QTextBlockGroup_NewQTextBlockGroup(PointerFromQTextDocument(document)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTextBlockGroup_DestroyQTextBlockGroup
func callbackQTextBlockGroup_DestroyQTextBlockGroup(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTextBlockGroup"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextBlockGroupFromPointer(ptr).DestroyQTextBlockGroupDefault()
	}
}

func (ptr *QTextBlockGroup) ConnectDestroyQTextBlockGroup(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTextBlockGroup"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTextBlockGroup", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTextBlockGroup", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextBlockGroup) DisconnectDestroyQTextBlockGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTextBlockGroup")
	}
}

func (ptr *QTextBlockGroup) DestroyQTextBlockGroup() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextBlockGroup_DestroyQTextBlockGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextBlockGroup) DestroyQTextBlockGroupDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextBlockGroup_DestroyQTextBlockGroupDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextBlockGroup) __blockList_atList(i int) *QTextBlock {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextBlockFromPointer(C.QTextBlockGroup___blockList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QTextBlock).DestroyQTextBlock)
		return tmpValue
	}
	return nil
}

func (ptr *QTextBlockGroup) __blockList_setList(i QTextBlock_ITF) {
	if ptr.Pointer() != nil {
		C.QTextBlockGroup___blockList_setList(ptr.Pointer(), PointerFromQTextBlock(i))
	}
}

func (ptr *QTextBlockGroup) __blockList_newList() unsafe.Pointer {
	return C.QTextBlockGroup___blockList_newList(ptr.Pointer())
}

type QTextCharFormat struct {
	QTextFormat
}

type QTextCharFormat_ITF interface {
	QTextFormat_ITF
	QTextCharFormat_PTR() *QTextCharFormat
}

func (ptr *QTextCharFormat) QTextCharFormat_PTR() *QTextCharFormat {
	return ptr
}

func (ptr *QTextCharFormat) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFormat_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextCharFormat) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextFormat_PTR().SetPointer(p)
	}
}

func PointerFromQTextCharFormat(ptr QTextCharFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextCharFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextCharFormatFromPointer(ptr unsafe.Pointer) (n *QTextCharFormat) {
	n = new(QTextCharFormat)
	n.SetPointer(ptr)
	return
}
func (ptr *QTextCharFormat) DestroyQTextCharFormat() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QTextCharFormat__FontPropertiesInheritanceBehavior
//QTextCharFormat::FontPropertiesInheritanceBehavior
type QTextCharFormat__FontPropertiesInheritanceBehavior int64

const (
	QTextCharFormat__FontPropertiesSpecifiedOnly QTextCharFormat__FontPropertiesInheritanceBehavior = QTextCharFormat__FontPropertiesInheritanceBehavior(0)
	QTextCharFormat__FontPropertiesAll           QTextCharFormat__FontPropertiesInheritanceBehavior = QTextCharFormat__FontPropertiesInheritanceBehavior(1)
)

// QTextCharFormat::UnderlineStyle
//
//go:generate stringer -type=QTextCharFormat__UnderlineStyle
type QTextCharFormat__UnderlineStyle int64

var (
	QTextCharFormat__NoUnderline         QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(0)
	QTextCharFormat__SingleUnderline     QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(1)
	QTextCharFormat__DashUnderline       QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(2)
	QTextCharFormat__DotLine             QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(3)
	QTextCharFormat__DashDotLine         QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(4)
	QTextCharFormat__DashDotDotLine      QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(5)
	QTextCharFormat__WaveUnderline       QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(6)
	QTextCharFormat__SpellCheckUnderline QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(7)
)

// QTextCharFormat::VerticalAlignment
//
//go:generate stringer -type=QTextCharFormat__VerticalAlignment
type QTextCharFormat__VerticalAlignment int64

const (
	QTextCharFormat__AlignNormal      QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(0)
	QTextCharFormat__AlignSuperScript QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(1)
	QTextCharFormat__AlignSubScript   QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(2)
	QTextCharFormat__AlignMiddle      QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(3)
	QTextCharFormat__AlignTop         QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(4)
	QTextCharFormat__AlignBottom      QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(5)
	QTextCharFormat__AlignBaseline    QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(6)
)

func NewQTextCharFormat() *QTextCharFormat {
	tmpValue := NewQTextCharFormatFromPointer(C.QTextCharFormat_NewQTextCharFormat())
	qt.SetFinalizer(tmpValue, (*QTextCharFormat).DestroyQTextCharFormat)
	return tmpValue
}

func (ptr *QTextCharFormat) Font() *QFont {
	if ptr.Pointer() != nil {
		tmpValue := NewQFontFromPointer(C.QTextCharFormat_Font(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QTextCharFormat) FontItalic() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextCharFormat_FontItalic(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontPointSize() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextCharFormat_FontPointSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) FontStrikeOut() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextCharFormat_FontStrikeOut(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontUnderline() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextCharFormat_FontUnderline(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontWeight() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextCharFormat_FontWeight(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCharFormat) SetFont(font QFont_ITF, behavior QTextCharFormat__FontPropertiesInheritanceBehavior) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFont(ptr.Pointer(), PointerFromQFont(font), C.longlong(behavior))
	}
}

func (ptr *QTextCharFormat) SetFont2(font QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFont2(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QTextCharFormat) SetFontItalic(italic bool) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontItalic(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(italic))))
	}
}

func (ptr *QTextCharFormat) SetFontPointSize(size float64) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontPointSize(ptr.Pointer(), C.double(size))
	}
}

func (ptr *QTextCharFormat) SetFontStrikeOut(strikeOut bool) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontStrikeOut(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(strikeOut))))
	}
}

func (ptr *QTextCharFormat) SetFontUnderline(underline bool) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontUnderline(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(underline))))
	}
}

func (ptr *QTextCharFormat) SetFontWeight(weight int) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontWeight(ptr.Pointer(), C.int(int32(weight)))
	}
}

func (ptr *QTextCharFormat) SetToolTip(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextCharFormat_SetToolTip(ptr.Pointer(), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextCharFormat) ToolTip() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextCharFormat_ToolTip(ptr.Pointer()))
	}
	return ""
}

type QTextCursor struct {
	ptr unsafe.Pointer
}

type QTextCursor_ITF interface {
	QTextCursor_PTR() *QTextCursor
}

func (ptr *QTextCursor) QTextCursor_PTR() *QTextCursor {
	return ptr
}

func (ptr *QTextCursor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextCursor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextCursor(ptr QTextCursor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextCursor_PTR().Pointer()
	}
	return nil
}

func NewQTextCursorFromPointer(ptr unsafe.Pointer) (n *QTextCursor) {
	n = new(QTextCursor)
	n.SetPointer(ptr)
	return
}

// QTextCursor::MoveMode
//
//go:generate stringer -type=QTextCursor__MoveMode
type QTextCursor__MoveMode int64

const (
	QTextCursor__MoveAnchor QTextCursor__MoveMode = QTextCursor__MoveMode(0)
	QTextCursor__KeepAnchor QTextCursor__MoveMode = QTextCursor__MoveMode(1)
)

// QTextCursor::MoveOperation
//
//go:generate stringer -type=QTextCursor__MoveOperation
type QTextCursor__MoveOperation int64

const (
	QTextCursor__NoMove            QTextCursor__MoveOperation = QTextCursor__MoveOperation(0)
	QTextCursor__Start             QTextCursor__MoveOperation = QTextCursor__MoveOperation(1)
	QTextCursor__Up                QTextCursor__MoveOperation = QTextCursor__MoveOperation(2)
	QTextCursor__StartOfLine       QTextCursor__MoveOperation = QTextCursor__MoveOperation(3)
	QTextCursor__StartOfBlock      QTextCursor__MoveOperation = QTextCursor__MoveOperation(4)
	QTextCursor__StartOfWord       QTextCursor__MoveOperation = QTextCursor__MoveOperation(5)
	QTextCursor__PreviousBlock     QTextCursor__MoveOperation = QTextCursor__MoveOperation(6)
	QTextCursor__PreviousCharacter QTextCursor__MoveOperation = QTextCursor__MoveOperation(7)
	QTextCursor__PreviousWord      QTextCursor__MoveOperation = QTextCursor__MoveOperation(8)
	QTextCursor__Left              QTextCursor__MoveOperation = QTextCursor__MoveOperation(9)
	QTextCursor__WordLeft          QTextCursor__MoveOperation = QTextCursor__MoveOperation(10)
	QTextCursor__End               QTextCursor__MoveOperation = QTextCursor__MoveOperation(11)
	QTextCursor__Down              QTextCursor__MoveOperation = QTextCursor__MoveOperation(12)
	QTextCursor__EndOfLine         QTextCursor__MoveOperation = QTextCursor__MoveOperation(13)
	QTextCursor__EndOfWord         QTextCursor__MoveOperation = QTextCursor__MoveOperation(14)
	QTextCursor__EndOfBlock        QTextCursor__MoveOperation = QTextCursor__MoveOperation(15)
	QTextCursor__NextBlock         QTextCursor__MoveOperation = QTextCursor__MoveOperation(16)
	QTextCursor__NextCharacter     QTextCursor__MoveOperation = QTextCursor__MoveOperation(17)
	QTextCursor__NextWord          QTextCursor__MoveOperation = QTextCursor__MoveOperation(18)
	QTextCursor__Right             QTextCursor__MoveOperation = QTextCursor__MoveOperation(19)
	QTextCursor__WordRight         QTextCursor__MoveOperation = QTextCursor__MoveOperation(20)
	QTextCursor__NextCell          QTextCursor__MoveOperation = QTextCursor__MoveOperation(21)
	QTextCursor__PreviousCell      QTextCursor__MoveOperation = QTextCursor__MoveOperation(22)
	QTextCursor__NextRow           QTextCursor__MoveOperation = QTextCursor__MoveOperation(23)
	QTextCursor__PreviousRow       QTextCursor__MoveOperation = QTextCursor__MoveOperation(24)
)

// QTextCursor::SelectionType
//
//go:generate stringer -type=QTextCursor__SelectionType
type QTextCursor__SelectionType int64

const (
	QTextCursor__WordUnderCursor  QTextCursor__SelectionType = QTextCursor__SelectionType(0)
	QTextCursor__LineUnderCursor  QTextCursor__SelectionType = QTextCursor__SelectionType(1)
	QTextCursor__BlockUnderCursor QTextCursor__SelectionType = QTextCursor__SelectionType(2)
	QTextCursor__Document         QTextCursor__SelectionType = QTextCursor__SelectionType(3)
)

func NewQTextCursor() *QTextCursor {
	tmpValue := NewQTextCursorFromPointer(C.QTextCursor_NewQTextCursor())
	qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
	return tmpValue
}

func NewQTextCursor2(document QTextDocument_ITF) *QTextCursor {
	tmpValue := NewQTextCursorFromPointer(C.QTextCursor_NewQTextCursor2(PointerFromQTextDocument(document)))
	qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
	return tmpValue
}

func NewQTextCursor3(frame QTextFrame_ITF) *QTextCursor {
	tmpValue := NewQTextCursorFromPointer(C.QTextCursor_NewQTextCursor3(PointerFromQTextFrame(frame)))
	qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
	return tmpValue
}

func NewQTextCursor4(block QTextBlock_ITF) *QTextCursor {
	tmpValue := NewQTextCursorFromPointer(C.QTextCursor_NewQTextCursor4(PointerFromQTextBlock(block)))
	qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
	return tmpValue
}

func NewQTextCursor5(cursor QTextCursor_ITF) *QTextCursor {
	tmpValue := NewQTextCursorFromPointer(C.QTextCursor_NewQTextCursor5(PointerFromQTextCursor(cursor)))
	qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
	return tmpValue
}

func (ptr *QTextCursor) Anchor() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextCursor_Anchor(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) BeginEditBlock() {
	if ptr.Pointer() != nil {
		C.QTextCursor_BeginEditBlock(ptr.Pointer())
	}
}

func (ptr *QTextCursor) Block() *QTextBlock {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextBlockFromPointer(C.QTextCursor_Block(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextBlock).DestroyQTextBlock)
		return tmpValue
	}
	return nil
}

func (ptr *QTextCursor) BlockFormat() *QTextBlockFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextBlockFormatFromPointer(C.QTextCursor_BlockFormat(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextBlockFormat).DestroyQTextBlockFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QTextCursor) CharFormat() *QTextCharFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextCharFormatFromPointer(C.QTextCursor_CharFormat(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextCharFormat).DestroyQTextCharFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QTextCursor) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QTextCursor_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QTextCursor) CreateList(format QTextListFormat_ITF) *QTextList {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextListFromPointer(C.QTextCursor_CreateList(ptr.Pointer(), PointerFromQTextListFormat(format)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextCursor) CreateList2(style QTextListFormat__Style) *QTextList {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextListFromPointer(C.QTextCursor_CreateList2(ptr.Pointer(), C.longlong(style)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextCursor) CurrentList() *QTextList {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextListFromPointer(C.QTextCursor_CurrentList(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextCursor) Document() *QTextDocument {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextDocumentFromPointer(C.QTextCursor_Document(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextCursor) EndEditBlock() {
	if ptr.Pointer() != nil {
		C.QTextCursor_EndEditBlock(ptr.Pointer())
	}
}

func (ptr *QTextCursor) HasSelection() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextCursor_HasSelection(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCursor) InsertImage(format QTextImageFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertImage(ptr.Pointer(), PointerFromQTextImageFormat(format))
	}
}

func (ptr *QTextCursor) InsertImage2(format QTextImageFormat_ITF, alignment QTextFrameFormat__Position) {
	if ptr.Pointer() != nil {
		C.QTextCursor_InsertImage2(ptr.Pointer(), PointerFromQTextImageFormat(format), C.longlong(alignment))
	}
}

func (ptr *QTextCursor) InsertImage3(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QTextCursor_InsertImage3(ptr.Pointer(), C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QTextCursor) InsertImage4(image QImage_ITF, name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QTextCursor_InsertImage4(ptr.Pointer(), PointerFromQImage(image), C.struct_QtGui_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QTextCursor) InsertTable(rows int, columns int, format QTextTableFormat_ITF) *QTextTable {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextTableFromPointer(C.QTextCursor_InsertTable(ptr.Pointer(), C.int(int32(rows)), C.int(int32(columns)), PointerFromQTextTableFormat(format)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextCursor) InsertTable2(rows int, columns int) *QTextTable {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextTableFromPointer(C.QTextCursor_InsertTable2(ptr.Pointer(), C.int(int32(rows)), C.int(int32(columns))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextCursor) InsertText(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextCursor_InsertText(ptr.Pointer(), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextCursor) InsertText2(text string, format QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextCursor_InsertText2(ptr.Pointer(), C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQTextCharFormat(format))
	}
}

func (ptr *QTextCursor) MergeCharFormat(modifier QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_MergeCharFormat(ptr.Pointer(), PointerFromQTextCharFormat(modifier))
	}
}

func (ptr *QTextCursor) Position() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextCursor_Position(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) RemoveSelectedText() {
	if ptr.Pointer() != nil {
		C.QTextCursor_RemoveSelectedText(ptr.Pointer())
	}
}

func (ptr *QTextCursor) Select(selection QTextCursor__SelectionType) {
	if ptr.Pointer() != nil {
		C.QTextCursor_Select(ptr.Pointer(), C.longlong(selection))
	}
}

func (ptr *QTextCursor) SelectedText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextCursor_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextCursor) Selection() *QTextDocumentFragment {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextDocumentFragmentFromPointer(C.QTextCursor_Selection(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextDocumentFragment).DestroyQTextDocumentFragment)
		return tmpValue
	}
	return nil
}

func (ptr *QTextCursor) SelectionEnd() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextCursor_SelectionEnd(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) SelectionStart() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextCursor_SelectionStart(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCursor) SetBlockFormat(format QTextBlockFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SetBlockFormat(ptr.Pointer(), PointerFromQTextBlockFormat(format))
	}
}

func (ptr *QTextCursor) SetCharFormat(format QTextCharFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SetCharFormat(ptr.Pointer(), PointerFromQTextCharFormat(format))
	}
}

func (ptr *QTextCursor) SetPosition(pos int, m QTextCursor__MoveMode) {
	if ptr.Pointer() != nil {
		C.QTextCursor_SetPosition(ptr.Pointer(), C.int(int32(pos)), C.longlong(m))
	}
}

func (ptr *QTextCursor) DestroyQTextCursor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextCursor_DestroyQTextCursor(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTextDocument struct {
	core.QObject
}

type QTextDocument_ITF interface {
	core.QObject_ITF
	QTextDocument_PTR() *QTextDocument
}

func (ptr *QTextDocument) QTextDocument_PTR() *QTextDocument {
	return ptr
}

func (ptr *QTextDocument) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextDocument) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQTextDocument(ptr QTextDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextDocument_PTR().Pointer()
	}
	return nil
}

func NewQTextDocumentFromPointer(ptr unsafe.Pointer) (n *QTextDocument) {
	n = new(QTextDocument)
	n.SetPointer(ptr)
	return
}

// QTextDocument::FindFlag
//
//go:generate stringer -type=QTextDocument__FindFlag
type QTextDocument__FindFlag int64

const (
	QTextDocument__FindBackward        QTextDocument__FindFlag = QTextDocument__FindFlag(0x00001)
	QTextDocument__FindCaseSensitively QTextDocument__FindFlag = QTextDocument__FindFlag(0x00002)
	QTextDocument__FindWholeWords      QTextDocument__FindFlag = QTextDocument__FindFlag(0x00004)
)

// QTextDocument::MetaInformation
//
//go:generate stringer -type=QTextDocument__MetaInformation
type QTextDocument__MetaInformation int64

const (
	QTextDocument__DocumentTitle QTextDocument__MetaInformation = QTextDocument__MetaInformation(0)
	QTextDocument__DocumentUrl   QTextDocument__MetaInformation = QTextDocument__MetaInformation(1)
)

// QTextDocument::ResourceType
//
//go:generate stringer -type=QTextDocument__ResourceType
type QTextDocument__ResourceType int64

const (
	QTextDocument__UnknownResource    QTextDocument__ResourceType = QTextDocument__ResourceType(0)
	QTextDocument__HtmlResource       QTextDocument__ResourceType = QTextDocument__ResourceType(1)
	QTextDocument__ImageResource      QTextDocument__ResourceType = QTextDocument__ResourceType(2)
	QTextDocument__StyleSheetResource QTextDocument__ResourceType = QTextDocument__ResourceType(3)
	QTextDocument__MarkdownResource   QTextDocument__ResourceType = QTextDocument__ResourceType(4)
	QTextDocument__UserResource       QTextDocument__ResourceType = QTextDocument__ResourceType(100)
)

// QTextDocument::Stacks
//
//go:generate stringer -type=QTextDocument__Stacks
type QTextDocument__Stacks int64

const (
	QTextDocument__UndoStack         QTextDocument__Stacks = QTextDocument__Stacks(0x01)
	QTextDocument__RedoStack         QTextDocument__Stacks = QTextDocument__Stacks(0x02)
	QTextDocument__UndoAndRedoStacks QTextDocument__Stacks = QTextDocument__Stacks(QTextDocument__UndoStack | QTextDocument__RedoStack)
)

func NewQTextDocument(parent core.QObject_ITF) *QTextDocument {
	tmpValue := NewQTextDocumentFromPointer(C.QTextDocument_NewQTextDocument(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQTextDocument2(text string, parent core.QObject_ITF) *QTextDocument {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQTextDocumentFromPointer(C.QTextDocument_NewQTextDocument2(C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QTextDocument) AddResource(ty int, name core.QUrl_ITF, resource core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_AddResource(ptr.Pointer(), C.int(int32(ty)), core.PointerFromQUrl(name), core.PointerFromQVariant(resource))
	}
}

func (ptr *QTextDocument) Begin() *QTextBlock {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextBlockFromPointer(C.QTextDocument_Begin(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextBlock).DestroyQTextBlock)
		return tmpValue
	}
	return nil
}

//export callbackQTextDocument_Clear
func callbackQTextDocument_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextDocumentFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QTextDocument) ConnectClear(f func()) {
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

func (ptr *QTextDocument) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *QTextDocument) Clear() {
	if ptr.Pointer() != nil {
		C.QTextDocument_Clear(ptr.Pointer())
	}
}

func (ptr *QTextDocument) ClearDefault() {
	if ptr.Pointer() != nil {
		C.QTextDocument_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QTextDocument) End() *QTextBlock {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextBlockFromPointer(C.QTextDocument_End(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextBlock).DestroyQTextBlock)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) Find(subString string, cursor QTextCursor_ITF, options QTextDocument__FindFlag) *QTextCursor {
	if ptr.Pointer() != nil {
		var subStringC *C.char
		if subString != "" {
			subStringC = C.CString(subString)
			defer C.free(unsafe.Pointer(subStringC))
		}
		tmpValue := NewQTextCursorFromPointer(C.QTextDocument_Find(ptr.Pointer(), C.struct_QtGui_PackedString{data: subStringC, len: C.longlong(len(subString))}, PointerFromQTextCursor(cursor), C.longlong(options)))
		qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) Find2(subString string, position int, options QTextDocument__FindFlag) *QTextCursor {
	if ptr.Pointer() != nil {
		var subStringC *C.char
		if subString != "" {
			subStringC = C.CString(subString)
			defer C.free(unsafe.Pointer(subStringC))
		}
		tmpValue := NewQTextCursorFromPointer(C.QTextDocument_Find2(ptr.Pointer(), C.struct_QtGui_PackedString{data: subStringC, len: C.longlong(len(subString))}, C.int(int32(position)), C.longlong(options)))
		qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) Find3(expr core.QRegExp_ITF, from int, options QTextDocument__FindFlag) *QTextCursor {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextCursorFromPointer(C.QTextDocument_Find3(ptr.Pointer(), core.PointerFromQRegExp(expr), C.int(int32(from)), C.longlong(options)))
		qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) Find4(expr core.QRegExp_ITF, cursor QTextCursor_ITF, options QTextDocument__FindFlag) *QTextCursor {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextCursorFromPointer(C.QTextDocument_Find4(ptr.Pointer(), core.PointerFromQRegExp(expr), PointerFromQTextCursor(cursor), C.longlong(options)))
		qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) Find5(expr core.QRegularExpression_ITF, from int, options QTextDocument__FindFlag) *QTextCursor {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextCursorFromPointer(C.QTextDocument_Find5(ptr.Pointer(), core.PointerFromQRegularExpression(expr), C.int(int32(from)), C.longlong(options)))
		qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) Find6(expr core.QRegularExpression_ITF, cursor QTextCursor_ITF, options QTextDocument__FindFlag) *QTextCursor {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextCursorFromPointer(C.QTextDocument_Find6(ptr.Pointer(), core.PointerFromQRegularExpression(expr), PointerFromQTextCursor(cursor), C.longlong(options)))
		qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) FindBlock(pos int) *QTextBlock {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextBlockFromPointer(C.QTextDocument_FindBlock(ptr.Pointer(), C.int(int32(pos))))
		qt.SetFinalizer(tmpValue, (*QTextBlock).DestroyQTextBlock)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) FirstBlock() *QTextBlock {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextBlockFromPointer(C.QTextDocument_FirstBlock(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextBlock).DestroyQTextBlock)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextDocument_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextDocument) IsModified() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextDocument_IsModified(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextDocument) Object(objectIndex int) *QTextObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextObjectFromPointer(C.QTextDocument_Object(ptr.Pointer(), C.int(int32(objectIndex))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) Print(printer QPagedPaintDevice_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_Print(ptr.Pointer(), PointerFromQPagedPaintDevice(printer))
	}
}

func (ptr *QTextDocument) Resource(ty int, name core.QUrl_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QTextDocument_Resource(ptr.Pointer(), C.int(int32(ty)), core.PointerFromQUrl(name)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) RootFrame() *QTextFrame {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextFrameFromPointer(C.QTextDocument_RootFrame(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) SetHtml(html string) {
	if ptr.Pointer() != nil {
		var htmlC *C.char
		if html != "" {
			htmlC = C.CString(html)
			defer C.free(unsafe.Pointer(htmlC))
		}
		C.QTextDocument_SetHtml(ptr.Pointer(), C.struct_QtGui_PackedString{data: htmlC, len: C.longlong(len(html))})
	}
}

//export callbackQTextDocument_SetModified
func callbackQTextDocument_SetModified(ptr unsafe.Pointer, m C.char) {
	if signal := qt.GetSignal(ptr, "setModified"); signal != nil {
		(*(*func(bool))(signal))(int8(m) != 0)
	} else {
		NewQTextDocumentFromPointer(ptr).SetModifiedDefault(int8(m) != 0)
	}
}

func (ptr *QTextDocument) ConnectSetModified(f func(m bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setModified"); signal != nil {
			f := func(m bool) {
				(*(*func(bool))(signal))(m)
				f(m)
			}
			qt.ConnectSignal(ptr.Pointer(), "setModified", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setModified", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextDocument) DisconnectSetModified() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setModified")
	}
}

func (ptr *QTextDocument) SetModified(m bool) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(m))))
	}
}

func (ptr *QTextDocument) SetModifiedDefault(m bool) {
	if ptr.Pointer() != nil {
		C.QTextDocument_SetModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(m))))
	}
}

func (ptr *QTextDocument) Size() *core.QSizeF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFFromPointer(C.QTextDocument_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSizeF).DestroyQSizeF)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) ToHtml(encoding core.QByteArray_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextDocument_ToHtml(ptr.Pointer(), core.PointerFromQByteArray(encoding)))
	}
	return ""
}

func (ptr *QTextDocument) ToPlainText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextDocument_ToPlainText(ptr.Pointer()))
	}
	return ""
}

//export callbackQTextDocument_DestroyQTextDocument
func callbackQTextDocument_DestroyQTextDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTextDocument"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextDocumentFromPointer(ptr).DestroyQTextDocumentDefault()
	}
}

func (ptr *QTextDocument) ConnectDestroyQTextDocument(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTextDocument"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTextDocument", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTextDocument", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextDocument) DisconnectDestroyQTextDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTextDocument")
	}
}

func (ptr *QTextDocument) DestroyQTextDocument() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextDocument_DestroyQTextDocument(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextDocument) DestroyQTextDocumentDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextDocument_DestroyQTextDocumentDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextDocument) __allFormats_atList(i int) *QTextFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextFormatFromPointer(C.QTextDocument___allFormats_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QTextFormat).DestroyQTextFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) __allFormats_setList(i QTextFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument___allFormats_setList(ptr.Pointer(), PointerFromQTextFormat(i))
	}
}

func (ptr *QTextDocument) __allFormats_newList() unsafe.Pointer {
	return C.QTextDocument___allFormats_newList(ptr.Pointer())
}

func (ptr *QTextDocument) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextDocument___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextDocument) __children_newList() unsafe.Pointer {
	return C.QTextDocument___children_newList(ptr.Pointer())
}

func (ptr *QTextDocument) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QTextDocument___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QTextDocument) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QTextDocument___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QTextDocument) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextDocument___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextDocument) __findChildren_newList() unsafe.Pointer {
	return C.QTextDocument___findChildren_newList(ptr.Pointer())
}

func (ptr *QTextDocument) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextDocument___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocument) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextDocument) __findChildren_newList3() unsafe.Pointer {
	return C.QTextDocument___findChildren_newList3(ptr.Pointer())
}

//export callbackQTextDocument_ChildEvent
func callbackQTextDocument_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTextDocumentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTextDocument) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTextDocument_ConnectNotify
func callbackQTextDocument_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTextDocumentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTextDocument) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTextDocument_CustomEvent
func callbackQTextDocument_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQTextDocumentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextDocument) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTextDocument_DeleteLater
func callbackQTextDocument_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextDocumentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTextDocument) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextDocument_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQTextDocument_Destroyed
func callbackQTextDocument_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQTextDocument_DisconnectNotify
func callbackQTextDocument_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTextDocumentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTextDocument) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTextDocument_Event
func callbackQTextDocument_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTextDocumentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTextDocument) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextDocument_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQTextDocument_EventFilter
func callbackQTextDocument_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTextDocumentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTextDocument) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextDocument_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQTextDocument_ObjectNameChanged
func callbackQTextDocument_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGui_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQTextDocument_TimerEvent
func callbackQTextDocument_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTextDocumentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTextDocument) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocument_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QTextDocumentFragment struct {
	ptr unsafe.Pointer
}

type QTextDocumentFragment_ITF interface {
	QTextDocumentFragment_PTR() *QTextDocumentFragment
}

func (ptr *QTextDocumentFragment) QTextDocumentFragment_PTR() *QTextDocumentFragment {
	return ptr
}

func (ptr *QTextDocumentFragment) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextDocumentFragment) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextDocumentFragment(ptr QTextDocumentFragment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextDocumentFragment_PTR().Pointer()
	}
	return nil
}

func NewQTextDocumentFragmentFromPointer(ptr unsafe.Pointer) (n *QTextDocumentFragment) {
	n = new(QTextDocumentFragment)
	n.SetPointer(ptr)
	return
}
func NewQTextDocumentFragment() *QTextDocumentFragment {
	tmpValue := NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment())
	qt.SetFinalizer(tmpValue, (*QTextDocumentFragment).DestroyQTextDocumentFragment)
	return tmpValue
}

func NewQTextDocumentFragment2(document QTextDocument_ITF) *QTextDocumentFragment {
	tmpValue := NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment2(PointerFromQTextDocument(document)))
	qt.SetFinalizer(tmpValue, (*QTextDocumentFragment).DestroyQTextDocumentFragment)
	return tmpValue
}

func NewQTextDocumentFragment3(cursor QTextCursor_ITF) *QTextDocumentFragment {
	tmpValue := NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment3(PointerFromQTextCursor(cursor)))
	qt.SetFinalizer(tmpValue, (*QTextDocumentFragment).DestroyQTextDocumentFragment)
	return tmpValue
}

func NewQTextDocumentFragment4(other QTextDocumentFragment_ITF) *QTextDocumentFragment {
	tmpValue := NewQTextDocumentFragmentFromPointer(C.QTextDocumentFragment_NewQTextDocumentFragment4(PointerFromQTextDocumentFragment(other)))
	qt.SetFinalizer(tmpValue, (*QTextDocumentFragment).DestroyQTextDocumentFragment)
	return tmpValue
}

func (ptr *QTextDocumentFragment) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextDocumentFragment_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextDocumentFragment) ToHtml(encoding core.QByteArray_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextDocumentFragment_ToHtml(ptr.Pointer(), core.PointerFromQByteArray(encoding)))
	}
	return ""
}

func (ptr *QTextDocumentFragment) ToPlainText() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextDocumentFragment_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextDocumentFragment) DestroyQTextDocumentFragment() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextDocumentFragment_DestroyQTextDocumentFragment(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTextDocumentWriter struct {
	ptr unsafe.Pointer
}

type QTextDocumentWriter_ITF interface {
	QTextDocumentWriter_PTR() *QTextDocumentWriter
}

func (ptr *QTextDocumentWriter) QTextDocumentWriter_PTR() *QTextDocumentWriter {
	return ptr
}

func (ptr *QTextDocumentWriter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextDocumentWriter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextDocumentWriter(ptr QTextDocumentWriter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextDocumentWriter_PTR().Pointer()
	}
	return nil
}

func NewQTextDocumentWriterFromPointer(ptr unsafe.Pointer) (n *QTextDocumentWriter) {
	n = new(QTextDocumentWriter)
	n.SetPointer(ptr)
	return
}
func NewQTextDocumentWriter() *QTextDocumentWriter {
	tmpValue := NewQTextDocumentWriterFromPointer(C.QTextDocumentWriter_NewQTextDocumentWriter())
	qt.SetFinalizer(tmpValue, (*QTextDocumentWriter).DestroyQTextDocumentWriter)
	return tmpValue
}

func NewQTextDocumentWriter2(device core.QIODevice_ITF, format core.QByteArray_ITF) *QTextDocumentWriter {
	tmpValue := NewQTextDocumentWriterFromPointer(C.QTextDocumentWriter_NewQTextDocumentWriter2(core.PointerFromQIODevice(device), core.PointerFromQByteArray(format)))
	qt.SetFinalizer(tmpValue, (*QTextDocumentWriter).DestroyQTextDocumentWriter)
	return tmpValue
}

func NewQTextDocumentWriter3(fileName string, format core.QByteArray_ITF) *QTextDocumentWriter {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	tmpValue := NewQTextDocumentWriterFromPointer(C.QTextDocumentWriter_NewQTextDocumentWriter3(C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, core.PointerFromQByteArray(format)))
	qt.SetFinalizer(tmpValue, (*QTextDocumentWriter).DestroyQTextDocumentWriter)
	return tmpValue
}

func (ptr *QTextDocumentWriter) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQIODeviceFromPointer(C.QTextDocumentWriter_Device(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocumentWriter) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextDocumentWriter_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextDocumentWriter) Format() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QTextDocumentWriter_Format(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocumentWriter) Write(document QTextDocument_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextDocumentWriter_Write(ptr.Pointer(), PointerFromQTextDocument(document))) != 0
	}
	return false
}

func (ptr *QTextDocumentWriter) Write2(fragment QTextDocumentFragment_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextDocumentWriter_Write2(ptr.Pointer(), PointerFromQTextDocumentFragment(fragment))) != 0
	}
	return false
}

func (ptr *QTextDocumentWriter) DestroyQTextDocumentWriter() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextDocumentWriter_DestroyQTextDocumentWriter(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextDocumentWriter) __supportedDocumentFormats_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QTextDocumentWriter___supportedDocumentFormats_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTextDocumentWriter) __supportedDocumentFormats_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTextDocumentWriter___supportedDocumentFormats_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QTextDocumentWriter) __supportedDocumentFormats_newList() unsafe.Pointer {
	return C.QTextDocumentWriter___supportedDocumentFormats_newList(ptr.Pointer())
}

type QTextFormat struct {
	ptr unsafe.Pointer
}

type QTextFormat_ITF interface {
	QTextFormat_PTR() *QTextFormat
}

func (ptr *QTextFormat) QTextFormat_PTR() *QTextFormat {
	return ptr
}

func (ptr *QTextFormat) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextFormat) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextFormat(ptr QTextFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextFormatFromPointer(ptr unsafe.Pointer) (n *QTextFormat) {
	n = new(QTextFormat)
	n.SetPointer(ptr)
	return
}

// QTextFormat::FormatType
//
//go:generate stringer -type=QTextFormat__FormatType
type QTextFormat__FormatType int64

const (
	QTextFormat__InvalidFormat QTextFormat__FormatType = QTextFormat__FormatType(-1)
	QTextFormat__BlockFormat   QTextFormat__FormatType = QTextFormat__FormatType(1)
	QTextFormat__CharFormat    QTextFormat__FormatType = QTextFormat__FormatType(2)
	QTextFormat__ListFormat    QTextFormat__FormatType = QTextFormat__FormatType(3)
	QTextFormat__TableFormat   QTextFormat__FormatType = QTextFormat__FormatType(4)
	QTextFormat__FrameFormat   QTextFormat__FormatType = QTextFormat__FormatType(5)
	QTextFormat__UserFormat    QTextFormat__FormatType = QTextFormat__FormatType(100)
)

// QTextFormat::ObjectTypes
//
//go:generate stringer -type=QTextFormat__ObjectTypes
type QTextFormat__ObjectTypes int64

const (
	QTextFormat__NoObject        QTextFormat__ObjectTypes = QTextFormat__ObjectTypes(0)
	QTextFormat__ImageObject     QTextFormat__ObjectTypes = QTextFormat__ObjectTypes(1)
	QTextFormat__TableObject     QTextFormat__ObjectTypes = QTextFormat__ObjectTypes(2)
	QTextFormat__TableCellObject QTextFormat__ObjectTypes = QTextFormat__ObjectTypes(3)
	QTextFormat__UserObject      QTextFormat__ObjectTypes = QTextFormat__ObjectTypes(0x1000)
)

// QTextFormat::PageBreakFlag
//
//go:generate stringer -type=QTextFormat__PageBreakFlag
type QTextFormat__PageBreakFlag int64

const (
	QTextFormat__PageBreak_Auto         QTextFormat__PageBreakFlag = QTextFormat__PageBreakFlag(0)
	QTextFormat__PageBreak_AlwaysBefore QTextFormat__PageBreakFlag = QTextFormat__PageBreakFlag(0x001)
	QTextFormat__PageBreak_AlwaysAfter  QTextFormat__PageBreakFlag = QTextFormat__PageBreakFlag(0x010)
)

// QTextFormat::Property
//
//go:generate stringer -type=QTextFormat__Property
type QTextFormat__Property int64

const (
	QTextFormat__ObjectIndex                       QTextFormat__Property = QTextFormat__Property(0x0)
	QTextFormat__CssFloat                          QTextFormat__Property = QTextFormat__Property(0x0800)
	QTextFormat__LayoutDirection                   QTextFormat__Property = QTextFormat__Property(0x0801)
	QTextFormat__OutlinePen                        QTextFormat__Property = QTextFormat__Property(0x810)
	QTextFormat__BackgroundBrush                   QTextFormat__Property = QTextFormat__Property(0x820)
	QTextFormat__ForegroundBrush                   QTextFormat__Property = QTextFormat__Property(0x821)
	QTextFormat__BackgroundImageUrl                QTextFormat__Property = QTextFormat__Property(0x823)
	QTextFormat__BlockAlignment                    QTextFormat__Property = QTextFormat__Property(0x1010)
	QTextFormat__BlockTopMargin                    QTextFormat__Property = QTextFormat__Property(0x1030)
	QTextFormat__BlockBottomMargin                 QTextFormat__Property = QTextFormat__Property(0x1031)
	QTextFormat__BlockLeftMargin                   QTextFormat__Property = QTextFormat__Property(0x1032)
	QTextFormat__BlockRightMargin                  QTextFormat__Property = QTextFormat__Property(0x1033)
	QTextFormat__TextIndent                        QTextFormat__Property = QTextFormat__Property(0x1034)
	QTextFormat__TabPositions                      QTextFormat__Property = QTextFormat__Property(0x1035)
	QTextFormat__BlockIndent                       QTextFormat__Property = QTextFormat__Property(0x1040)
	QTextFormat__LineHeight                        QTextFormat__Property = QTextFormat__Property(0x1048)
	QTextFormat__LineHeightType                    QTextFormat__Property = QTextFormat__Property(0x1049)
	QTextFormat__BlockNonBreakableLines            QTextFormat__Property = QTextFormat__Property(0x1050)
	QTextFormat__BlockTrailingHorizontalRulerWidth QTextFormat__Property = QTextFormat__Property(0x1060)
	QTextFormat__HeadingLevel                      QTextFormat__Property = QTextFormat__Property(0x1070)
	QTextFormat__BlockQuoteLevel                   QTextFormat__Property = QTextFormat__Property(0x1080)
	QTextFormat__BlockCodeLanguage                 QTextFormat__Property = QTextFormat__Property(0x1090)
	QTextFormat__BlockCodeFence                    QTextFormat__Property = QTextFormat__Property(0x1091)
	QTextFormat__BlockMarker                       QTextFormat__Property = QTextFormat__Property(0x10A0)
	QTextFormat__FirstFontProperty                 QTextFormat__Property = QTextFormat__Property(0x1FE0)
	QTextFormat__FontCapitalization                QTextFormat__Property = QTextFormat__Property(QTextFormat__FirstFontProperty)
	QTextFormat__FontLetterSpacingType             QTextFormat__Property = QTextFormat__Property(0x2033)
	QTextFormat__FontLetterSpacing                 QTextFormat__Property = QTextFormat__Property(0x1FE1)
	QTextFormat__FontWordSpacing                   QTextFormat__Property = QTextFormat__Property(0x1FE2)
	QTextFormat__FontStretch                       QTextFormat__Property = QTextFormat__Property(0x2034)
	QTextFormat__FontStyleHint                     QTextFormat__Property = QTextFormat__Property(0x1FE3)
	QTextFormat__FontStyleStrategy                 QTextFormat__Property = QTextFormat__Property(0x1FE4)
	QTextFormat__FontKerning                       QTextFormat__Property = QTextFormat__Property(0x1FE5)
	QTextFormat__FontHintingPreference             QTextFormat__Property = QTextFormat__Property(0x1FE6)
	QTextFormat__FontFamilies                      QTextFormat__Property = QTextFormat__Property(0x1FE7)
	QTextFormat__FontStyleName                     QTextFormat__Property = QTextFormat__Property(0x1FE8)
	QTextFormat__FontFamily                        QTextFormat__Property = QTextFormat__Property(0x2000)
	QTextFormat__FontPointSize                     QTextFormat__Property = QTextFormat__Property(0x2001)
	QTextFormat__FontSizeAdjustment                QTextFormat__Property = QTextFormat__Property(0x2002)
	QTextFormat__FontSizeIncrement                 QTextFormat__Property = QTextFormat__Property(QTextFormat__FontSizeAdjustment)
	QTextFormat__FontWeight                        QTextFormat__Property = QTextFormat__Property(0x2003)
	QTextFormat__FontItalic                        QTextFormat__Property = QTextFormat__Property(0x2004)
	QTextFormat__FontUnderline                     QTextFormat__Property = QTextFormat__Property(0x2005)
	QTextFormat__FontOverline                      QTextFormat__Property = QTextFormat__Property(0x2006)
	QTextFormat__FontStrikeOut                     QTextFormat__Property = QTextFormat__Property(0x2007)
	QTextFormat__FontFixedPitch                    QTextFormat__Property = QTextFormat__Property(0x2008)
	QTextFormat__FontPixelSize                     QTextFormat__Property = QTextFormat__Property(0x2009)
	QTextFormat__LastFontProperty                  QTextFormat__Property = QTextFormat__Property(QTextFormat__FontPixelSize)
	QTextFormat__TextUnderlineColor                QTextFormat__Property = QTextFormat__Property(0x2010)
	QTextFormat__TextVerticalAlignment             QTextFormat__Property = QTextFormat__Property(0x2021)
	QTextFormat__TextOutline                       QTextFormat__Property = QTextFormat__Property(0x2022)
	QTextFormat__TextUnderlineStyle                QTextFormat__Property = QTextFormat__Property(0x2023)
	QTextFormat__TextToolTip                       QTextFormat__Property = QTextFormat__Property(0x2024)
	QTextFormat__IsAnchor                          QTextFormat__Property = QTextFormat__Property(0x2030)
	QTextFormat__AnchorHref                        QTextFormat__Property = QTextFormat__Property(0x2031)
	QTextFormat__AnchorName                        QTextFormat__Property = QTextFormat__Property(0x2032)
	QTextFormat__ObjectType                        QTextFormat__Property = QTextFormat__Property(0x2f00)
	QTextFormat__ListStyle                         QTextFormat__Property = QTextFormat__Property(0x3000)
	QTextFormat__ListIndent                        QTextFormat__Property = QTextFormat__Property(0x3001)
	QTextFormat__ListNumberPrefix                  QTextFormat__Property = QTextFormat__Property(0x3002)
	QTextFormat__ListNumberSuffix                  QTextFormat__Property = QTextFormat__Property(0x3003)
	QTextFormat__FrameBorder                       QTextFormat__Property = QTextFormat__Property(0x4000)
	QTextFormat__FrameMargin                       QTextFormat__Property = QTextFormat__Property(0x4001)
	QTextFormat__FramePadding                      QTextFormat__Property = QTextFormat__Property(0x4002)
	QTextFormat__FrameWidth                        QTextFormat__Property = QTextFormat__Property(0x4003)
	QTextFormat__FrameHeight                       QTextFormat__Property = QTextFormat__Property(0x4004)
	QTextFormat__FrameTopMargin                    QTextFormat__Property = QTextFormat__Property(0x4005)
	QTextFormat__FrameBottomMargin                 QTextFormat__Property = QTextFormat__Property(0x4006)
	QTextFormat__FrameLeftMargin                   QTextFormat__Property = QTextFormat__Property(0x4007)
	QTextFormat__FrameRightMargin                  QTextFormat__Property = QTextFormat__Property(0x4008)
	QTextFormat__FrameBorderBrush                  QTextFormat__Property = QTextFormat__Property(0x4009)
	QTextFormat__FrameBorderStyle                  QTextFormat__Property = QTextFormat__Property(0x4010)
	QTextFormat__TableColumns                      QTextFormat__Property = QTextFormat__Property(0x4100)
	QTextFormat__TableColumnWidthConstraints       QTextFormat__Property = QTextFormat__Property(0x4101)
	QTextFormat__TableCellSpacing                  QTextFormat__Property = QTextFormat__Property(0x4102)
	QTextFormat__TableCellPadding                  QTextFormat__Property = QTextFormat__Property(0x4103)
	QTextFormat__TableHeaderRowCount               QTextFormat__Property = QTextFormat__Property(0x4104)
	QTextFormat__TableBorderCollapse               QTextFormat__Property = QTextFormat__Property(0x4105)
	QTextFormat__TableCellRowSpan                  QTextFormat__Property = QTextFormat__Property(0x4810)
	QTextFormat__TableCellColumnSpan               QTextFormat__Property = QTextFormat__Property(0x4811)
	QTextFormat__TableCellTopPadding               QTextFormat__Property = QTextFormat__Property(0x4812)
	QTextFormat__TableCellBottomPadding            QTextFormat__Property = QTextFormat__Property(0x4813)
	QTextFormat__TableCellLeftPadding              QTextFormat__Property = QTextFormat__Property(0x4814)
	QTextFormat__TableCellRightPadding             QTextFormat__Property = QTextFormat__Property(0x4815)
	QTextFormat__TableCellTopBorder                QTextFormat__Property = QTextFormat__Property(0x4816)
	QTextFormat__TableCellBottomBorder             QTextFormat__Property = QTextFormat__Property(0x4817)
	QTextFormat__TableCellLeftBorder               QTextFormat__Property = QTextFormat__Property(0x4818)
	QTextFormat__TableCellRightBorder              QTextFormat__Property = QTextFormat__Property(0x4819)
	QTextFormat__TableCellTopBorderStyle           QTextFormat__Property = QTextFormat__Property(0x481a)
	QTextFormat__TableCellBottomBorderStyle        QTextFormat__Property = QTextFormat__Property(0x481b)
	QTextFormat__TableCellLeftBorderStyle          QTextFormat__Property = QTextFormat__Property(0x481c)
	QTextFormat__TableCellRightBorderStyle         QTextFormat__Property = QTextFormat__Property(0x481d)
	QTextFormat__TableCellTopBorderBrush           QTextFormat__Property = QTextFormat__Property(0x481e)
	QTextFormat__TableCellBottomBorderBrush        QTextFormat__Property = QTextFormat__Property(0x481f)
	QTextFormat__TableCellLeftBorderBrush          QTextFormat__Property = QTextFormat__Property(0x4820)
	QTextFormat__TableCellRightBorderBrush         QTextFormat__Property = QTextFormat__Property(0x4821)
	QTextFormat__ImageName                         QTextFormat__Property = QTextFormat__Property(0x5000)
	QTextFormat__ImageTitle                        QTextFormat__Property = QTextFormat__Property(0x5001)
	QTextFormat__ImageAltText                      QTextFormat__Property = QTextFormat__Property(0x5002)
	QTextFormat__ImageWidth                        QTextFormat__Property = QTextFormat__Property(0x5010)
	QTextFormat__ImageHeight                       QTextFormat__Property = QTextFormat__Property(0x5011)
	QTextFormat__ImageQuality                      QTextFormat__Property = QTextFormat__Property(0x5014)
	QTextFormat__FullWidthSelection                QTextFormat__Property = QTextFormat__Property(0x06000)
	QTextFormat__PageBreakPolicy                   QTextFormat__Property = QTextFormat__Property(0x7000)
	QTextFormat__UserProperty                      QTextFormat__Property = QTextFormat__Property(0x100000)
)

func NewQTextFormat() *QTextFormat {
	tmpValue := NewQTextFormatFromPointer(C.QTextFormat_NewQTextFormat())
	qt.SetFinalizer(tmpValue, (*QTextFormat).DestroyQTextFormat)
	return tmpValue
}

func NewQTextFormat2(ty int) *QTextFormat {
	tmpValue := NewQTextFormatFromPointer(C.QTextFormat_NewQTextFormat2(C.int(int32(ty))))
	qt.SetFinalizer(tmpValue, (*QTextFormat).DestroyQTextFormat)
	return tmpValue
}

func NewQTextFormat3(other QTextFormat_ITF) *QTextFormat {
	tmpValue := NewQTextFormatFromPointer(C.QTextFormat_NewQTextFormat3(PointerFromQTextFormat(other)))
	qt.SetFinalizer(tmpValue, (*QTextFormat).DestroyQTextFormat)
	return tmpValue
}

func (ptr *QTextFormat) Background() *QBrush {
	if ptr.Pointer() != nil {
		tmpValue := NewQBrushFromPointer(C.QTextFormat_Background(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QTextFormat) ClearBackground() {
	if ptr.Pointer() != nil {
		C.QTextFormat_ClearBackground(ptr.Pointer())
	}
}

func (ptr *QTextFormat) ClearForeground() {
	if ptr.Pointer() != nil {
		C.QTextFormat_ClearForeground(ptr.Pointer())
	}
}

func (ptr *QTextFormat) Foreground() *QBrush {
	if ptr.Pointer() != nil {
		tmpValue := NewQBrushFromPointer(C.QTextFormat_Foreground(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QTextFormat) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextFormat_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextFormat) IsTableCellFormat() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextFormat_IsTableCellFormat(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextFormat) IsTableFormat() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextFormat_IsTableFormat(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextFormat_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextFormat) Merge(other QTextFormat_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFormat_Merge(ptr.Pointer(), PointerFromQTextFormat(other))
	}
}

func (ptr *QTextFormat) ObjectIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextFormat_ObjectIndex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextFormat) Property(propertyId int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QTextFormat_Property(ptr.Pointer(), C.int(int32(propertyId))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QTextFormat) SetBackground(brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFormat_SetBackground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QTextFormat) SetForeground(brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFormat_SetForeground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QTextFormat) SetObjectIndex(index int) {
	if ptr.Pointer() != nil {
		C.QTextFormat_SetObjectIndex(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QTextFormat) Type() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextFormat_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextFormat) DestroyQTextFormat() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextFormat_DestroyQTextFormat(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextFormat) __lengthVectorProperty_atList(i int) *QTextLength {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextLengthFromPointer(C.QTextFormat___lengthVectorProperty_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QTextLength).DestroyQTextLength)
		return tmpValue
	}
	return nil
}

func (ptr *QTextFormat) __lengthVectorProperty_setList(i QTextLength_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFormat___lengthVectorProperty_setList(ptr.Pointer(), PointerFromQTextLength(i))
	}
}

func (ptr *QTextFormat) __lengthVectorProperty_newList() unsafe.Pointer {
	return C.QTextFormat___lengthVectorProperty_newList(ptr.Pointer())
}

func (ptr *QTextFormat) __properties_atList(v int, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QTextFormat___properties_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QTextFormat) __properties_setList(key int, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFormat___properties_setList(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(i))
	}
}

func (ptr *QTextFormat) __properties_newList() unsafe.Pointer {
	return C.QTextFormat___properties_newList(ptr.Pointer())
}

func (ptr *QTextFormat) __properties_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtGui_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewQTextFormatFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____properties_keyList_atList(i)
			}
			return out
		}(C.QTextFormat___properties_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QTextFormat) __setProperty_value_atList2(i int) *QTextLength {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextLengthFromPointer(C.QTextFormat___setProperty_value_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QTextLength).DestroyQTextLength)
		return tmpValue
	}
	return nil
}

func (ptr *QTextFormat) __setProperty_value_setList2(i QTextLength_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFormat___setProperty_value_setList2(ptr.Pointer(), PointerFromQTextLength(i))
	}
}

func (ptr *QTextFormat) __setProperty_value_newList2() unsafe.Pointer {
	return C.QTextFormat___setProperty_value_newList2(ptr.Pointer())
}

func (ptr *QTextFormat) ____properties_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextFormat_____properties_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QTextFormat) ____properties_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.QTextFormat_____properties_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QTextFormat) ____properties_keyList_newList() unsafe.Pointer {
	return C.QTextFormat_____properties_keyList_newList(ptr.Pointer())
}

type QTextFrame struct {
	QTextObject
}

type QTextFrame_ITF interface {
	QTextObject_ITF
	QTextFrame_PTR() *QTextFrame
}

func (ptr *QTextFrame) QTextFrame_PTR() *QTextFrame {
	return ptr
}

func (ptr *QTextFrame) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextFrame) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextObject_PTR().SetPointer(p)
	}
}

func PointerFromQTextFrame(ptr QTextFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFrame_PTR().Pointer()
	}
	return nil
}

func NewQTextFrameFromPointer(ptr unsafe.Pointer) (n *QTextFrame) {
	n = new(QTextFrame)
	n.SetPointer(ptr)
	return
}
func NewQTextFrame(document QTextDocument_ITF) *QTextFrame {
	tmpValue := NewQTextFrameFromPointer(C.QTextFrame_NewQTextFrame(PointerFromQTextDocument(document)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QTextFrame) ChildFrames() []*QTextFrame {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtGui_PackedList) []*QTextFrame {
			out := make([]*QTextFrame, int(l.len))
			tmpList := NewQTextFrameFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__childFrames_atList(i)
			}
			return out
		}(C.QTextFrame_ChildFrames(ptr.Pointer()))
	}
	return make([]*QTextFrame, 0)
}

func (ptr *QTextFrame) FirstCursorPosition() *QTextCursor {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextCursorFromPointer(C.QTextFrame_FirstCursorPosition(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QTextFrame) FrameFormat() *QTextFrameFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextFrameFormatFromPointer(C.QTextFrame_FrameFormat(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextFrameFormat).DestroyQTextFrameFormat)
		return tmpValue
	}
	return nil
}

//export callbackQTextFrame_DestroyQTextFrame
func callbackQTextFrame_DestroyQTextFrame(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTextFrame"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextFrameFromPointer(ptr).DestroyQTextFrameDefault()
	}
}

func (ptr *QTextFrame) ConnectDestroyQTextFrame(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTextFrame"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTextFrame", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTextFrame", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextFrame) DisconnectDestroyQTextFrame() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTextFrame")
	}
}

func (ptr *QTextFrame) DestroyQTextFrame() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextFrame_DestroyQTextFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextFrame) DestroyQTextFrameDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextFrame_DestroyQTextFrameDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextFrame) __childFrames_atList(i int) *QTextFrame {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextFrameFromPointer(C.QTextFrame___childFrames_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextFrame) __childFrames_setList(i QTextFrame_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFrame___childFrames_setList(ptr.Pointer(), PointerFromQTextFrame(i))
	}
}

func (ptr *QTextFrame) __childFrames_newList() unsafe.Pointer {
	return C.QTextFrame___childFrames_newList(ptr.Pointer())
}

type QTextFrameFormat struct {
	QTextFormat
}

type QTextFrameFormat_ITF interface {
	QTextFormat_ITF
	QTextFrameFormat_PTR() *QTextFrameFormat
}

func (ptr *QTextFrameFormat) QTextFrameFormat_PTR() *QTextFrameFormat {
	return ptr
}

func (ptr *QTextFrameFormat) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFormat_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextFrameFormat) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextFormat_PTR().SetPointer(p)
	}
}

func PointerFromQTextFrameFormat(ptr QTextFrameFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFrameFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextFrameFormatFromPointer(ptr unsafe.Pointer) (n *QTextFrameFormat) {
	n = new(QTextFrameFormat)
	n.SetPointer(ptr)
	return
}
func (ptr *QTextFrameFormat) DestroyQTextFrameFormat() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QTextFrameFormat__BorderStyle
//QTextFrameFormat::BorderStyle
type QTextFrameFormat__BorderStyle int64

var (
	QTextFrameFormat__BorderStyle_None       QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(0)
	QTextFrameFormat__BorderStyle_Dotted     QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(1)
	QTextFrameFormat__BorderStyle_Dashed     QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(2)
	QTextFrameFormat__BorderStyle_Solid      QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(3)
	QTextFrameFormat__BorderStyle_Double     QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(4)
	QTextFrameFormat__BorderStyle_DotDash    QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(5)
	QTextFrameFormat__BorderStyle_DotDotDash QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(6)
	QTextFrameFormat__BorderStyle_Groove     QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(7)
	QTextFrameFormat__BorderStyle_Ridge      QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(8)
	QTextFrameFormat__BorderStyle_Inset      QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(9)
	QTextFrameFormat__BorderStyle_Outset     QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(10)
)

// QTextFrameFormat::Position
//
//go:generate stringer -type=QTextFrameFormat__Position
type QTextFrameFormat__Position int64

const (
	QTextFrameFormat__InFlow     QTextFrameFormat__Position = QTextFrameFormat__Position(0)
	QTextFrameFormat__FloatLeft  QTextFrameFormat__Position = QTextFrameFormat__Position(1)
	QTextFrameFormat__FloatRight QTextFrameFormat__Position = QTextFrameFormat__Position(2)
)

func NewQTextFrameFormat() *QTextFrameFormat {
	tmpValue := NewQTextFrameFormatFromPointer(C.QTextFrameFormat_NewQTextFrameFormat())
	qt.SetFinalizer(tmpValue, (*QTextFrameFormat).DestroyQTextFrameFormat)
	return tmpValue
}

func (ptr *QTextFrameFormat) Border() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextFrameFormat_Border(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) BorderBrush() *QBrush {
	if ptr.Pointer() != nil {
		tmpValue := NewQBrushFromPointer(C.QTextFrameFormat_BorderBrush(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBrush).DestroyQBrush)
		return tmpValue
	}
	return nil
}

func (ptr *QTextFrameFormat) Height() *QTextLength {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextLengthFromPointer(C.QTextFrameFormat_Height(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextLength).DestroyQTextLength)
		return tmpValue
	}
	return nil
}

func (ptr *QTextFrameFormat) Margin() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextFrameFormat_Margin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) Position() QTextFrameFormat__Position {
	if ptr.Pointer() != nil {
		return QTextFrameFormat__Position(C.QTextFrameFormat_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) SetBorder(width float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetBorder(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QTextFrameFormat) SetBorderBrush(brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetBorderBrush(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QTextFrameFormat) SetPosition(policy QTextFrameFormat__Position) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetPosition(ptr.Pointer(), C.longlong(policy))
	}
}

func (ptr *QTextFrameFormat) SetTopMargin(margin float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetTopMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextFrameFormat) TopMargin() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextFrameFormat_TopMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) Width() *QTextLength {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextLengthFromPointer(C.QTextFrameFormat_Width(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextLength).DestroyQTextLength)
		return tmpValue
	}
	return nil
}

type QTextImageFormat struct {
	QTextCharFormat
}

type QTextImageFormat_ITF interface {
	QTextCharFormat_ITF
	QTextImageFormat_PTR() *QTextImageFormat
}

func (ptr *QTextImageFormat) QTextImageFormat_PTR() *QTextImageFormat {
	return ptr
}

func (ptr *QTextImageFormat) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextCharFormat_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextImageFormat) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextCharFormat_PTR().SetPointer(p)
	}
}

func PointerFromQTextImageFormat(ptr QTextImageFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextImageFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextImageFormatFromPointer(ptr unsafe.Pointer) (n *QTextImageFormat) {
	n = new(QTextImageFormat)
	n.SetPointer(ptr)
	return
}
func (ptr *QTextImageFormat) DestroyQTextImageFormat() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQTextImageFormat() *QTextImageFormat {
	tmpValue := NewQTextImageFormatFromPointer(C.QTextImageFormat_NewQTextImageFormat())
	qt.SetFinalizer(tmpValue, (*QTextImageFormat).DestroyQTextImageFormat)
	return tmpValue
}

func (ptr *QTextImageFormat) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextImageFormat_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextImageFormat) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextImageFormat_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextImageFormat) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextImageFormat_Width(ptr.Pointer()))
	}
	return 0
}

// QTextItem::RenderFlag
//
//go:generate stringer -type=QTextItem__RenderFlag
type QTextItem__RenderFlag int64

const (
	QTextItem__RightToLeft QTextItem__RenderFlag = QTextItem__RenderFlag(0x1)
	QTextItem__Overline    QTextItem__RenderFlag = QTextItem__RenderFlag(0x10)
	QTextItem__Underline   QTextItem__RenderFlag = QTextItem__RenderFlag(0x20)
	QTextItem__StrikeOut   QTextItem__RenderFlag = QTextItem__RenderFlag(0x40)
	QTextItem__Dummy       QTextItem__RenderFlag = QTextItem__RenderFlag(0xffffffff)
)

type QTextLayout struct {
	ptr unsafe.Pointer
}

type QTextLayout_ITF interface {
	QTextLayout_PTR() *QTextLayout
}

func (ptr *QTextLayout) QTextLayout_PTR() *QTextLayout {
	return ptr
}

func (ptr *QTextLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextLayout(ptr QTextLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextLayout_PTR().Pointer()
	}
	return nil
}

func NewQTextLayoutFromPointer(ptr unsafe.Pointer) (n *QTextLayout) {
	n = new(QTextLayout)
	n.SetPointer(ptr)
	return
}

// QTextLayout::CursorMode
//
//go:generate stringer -type=QTextLayout__CursorMode
type QTextLayout__CursorMode int64

const (
	QTextLayout__SkipCharacters QTextLayout__CursorMode = QTextLayout__CursorMode(0)
	QTextLayout__SkipWords      QTextLayout__CursorMode = QTextLayout__CursorMode(1)
)

func NewQTextLayout() *QTextLayout {
	tmpValue := NewQTextLayoutFromPointer(C.QTextLayout_NewQTextLayout())
	qt.SetFinalizer(tmpValue, (*QTextLayout).DestroyQTextLayout)
	return tmpValue
}

func NewQTextLayout2(text string) *QTextLayout {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQTextLayoutFromPointer(C.QTextLayout_NewQTextLayout2(C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}))
	qt.SetFinalizer(tmpValue, (*QTextLayout).DestroyQTextLayout)
	return tmpValue
}

func NewQTextLayout4(text string, font QFont_ITF, paintdevice QPaintDevice_ITF) *QTextLayout {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	tmpValue := NewQTextLayoutFromPointer(C.QTextLayout_NewQTextLayout4(C.struct_QtGui_PackedString{data: textC, len: C.longlong(len(text))}, PointerFromQFont(font), PointerFromQPaintDevice(paintdevice)))
	qt.SetFinalizer(tmpValue, (*QTextLayout).DestroyQTextLayout)
	return tmpValue
}

func (ptr *QTextLayout) BoundingRect() *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QTextLayout_BoundingRect(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QTextLayout) Font() *QFont {
	if ptr.Pointer() != nil {
		tmpValue := NewQFontFromPointer(C.QTextLayout_Font(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QFont).DestroyQFont)
		return tmpValue
	}
	return nil
}

func (ptr *QTextLayout) MaximumWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextLayout_MaximumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLayout) MinimumWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextLayout_MinimumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLayout) Position() *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QTextLayout_Position(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QTextLayout) SetFont(font QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QTextLayout_SetFont(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QTextLayout) SetPosition(p core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QTextLayout_SetPosition(ptr.Pointer(), core.PointerFromQPointF(p))
	}
}

func (ptr *QTextLayout) SetText(stri string) {
	if ptr.Pointer() != nil {
		var striC *C.char
		if stri != "" {
			striC = C.CString(stri)
			defer C.free(unsafe.Pointer(striC))
		}
		C.QTextLayout_SetText(ptr.Pointer(), C.struct_QtGui_PackedString{data: striC, len: C.longlong(len(stri))})
	}
}

func (ptr *QTextLayout) Text() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextLayout_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextLayout) DestroyQTextLayout() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextLayout_DestroyQTextLayout(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextLayout) __formats_newList() unsafe.Pointer {
	return C.QTextLayout___formats_newList(ptr.Pointer())
}

func (ptr *QTextLayout) __glyphRuns_atList(i int) *QGlyphRun {
	if ptr.Pointer() != nil {
		tmpValue := NewQGlyphRunFromPointer(C.QTextLayout___glyphRuns_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QGlyphRun).DestroyQGlyphRun)
		return tmpValue
	}
	return nil
}

func (ptr *QTextLayout) __glyphRuns_setList(i QGlyphRun_ITF) {
	if ptr.Pointer() != nil {
		C.QTextLayout___glyphRuns_setList(ptr.Pointer(), PointerFromQGlyphRun(i))
	}
}

func (ptr *QTextLayout) __glyphRuns_newList() unsafe.Pointer {
	return C.QTextLayout___glyphRuns_newList(ptr.Pointer())
}

type QTextLength struct {
	ptr unsafe.Pointer
}

type QTextLength_ITF interface {
	QTextLength_PTR() *QTextLength
}

func (ptr *QTextLength) QTextLength_PTR() *QTextLength {
	return ptr
}

func (ptr *QTextLength) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextLength) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextLength(ptr QTextLength_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextLength_PTR().Pointer()
	}
	return nil
}

func NewQTextLengthFromPointer(ptr unsafe.Pointer) (n *QTextLength) {
	n = new(QTextLength)
	n.SetPointer(ptr)
	return
}
func (ptr *QTextLength) DestroyQTextLength() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QTextLength__Type
//QTextLength::Type
type QTextLength__Type int64

const (
	QTextLength__VariableLength   QTextLength__Type = QTextLength__Type(0)
	QTextLength__FixedLength      QTextLength__Type = QTextLength__Type(1)
	QTextLength__PercentageLength QTextLength__Type = QTextLength__Type(2)
)

func NewQTextLength() *QTextLength {
	tmpValue := NewQTextLengthFromPointer(C.QTextLength_NewQTextLength())
	qt.SetFinalizer(tmpValue, (*QTextLength).DestroyQTextLength)
	return tmpValue
}

func NewQTextLength2(ty QTextLength__Type, value float64) *QTextLength {
	tmpValue := NewQTextLengthFromPointer(C.QTextLength_NewQTextLength2(C.longlong(ty), C.double(value)))
	qt.SetFinalizer(tmpValue, (*QTextLength).DestroyQTextLength)
	return tmpValue
}

func (ptr *QTextLength) Type() QTextLength__Type {
	if ptr.Pointer() != nil {
		return QTextLength__Type(C.QTextLength_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLength) Value(maximumLength float64) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextLength_Value(ptr.Pointer(), C.double(maximumLength)))
	}
	return 0
}

// QTextLine::CursorPosition
//
//go:generate stringer -type=QTextLine__CursorPosition
type QTextLine__CursorPosition int64

const (
	QTextLine__CursorBetweenCharacters QTextLine__CursorPosition = QTextLine__CursorPosition(0)
	QTextLine__CursorOnCharacter       QTextLine__CursorPosition = QTextLine__CursorPosition(1)
)

// QTextLine::Edge
//
//go:generate stringer -type=QTextLine__Edge
type QTextLine__Edge int64

const (
	QTextLine__Leading  QTextLine__Edge = QTextLine__Edge(0)
	QTextLine__Trailing QTextLine__Edge = QTextLine__Edge(1)
)

type QTextList struct {
	QTextBlockGroup
}

type QTextList_ITF interface {
	QTextBlockGroup_ITF
	QTextList_PTR() *QTextList
}

func (ptr *QTextList) QTextList_PTR() *QTextList {
	return ptr
}

func (ptr *QTextList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBlockGroup_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextBlockGroup_PTR().SetPointer(p)
	}
}

func PointerFromQTextList(ptr QTextList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextList_PTR().Pointer()
	}
	return nil
}

func NewQTextListFromPointer(ptr unsafe.Pointer) (n *QTextList) {
	n = new(QTextList)
	n.SetPointer(ptr)
	return
}
func (ptr *QTextList) Add(block QTextBlock_ITF) {
	if ptr.Pointer() != nil {
		C.QTextList_Add(ptr.Pointer(), PointerFromQTextBlock(block))
	}
}

func (ptr *QTextList) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextList_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextList) Format() *QTextListFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextListFormatFromPointer(C.QTextList_Format(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextListFormat).DestroyQTextListFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QTextList) Item(i int) *QTextBlock {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextBlockFromPointer(C.QTextList_Item(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QTextBlock).DestroyQTextBlock)
		return tmpValue
	}
	return nil
}

func (ptr *QTextList) Remove(block QTextBlock_ITF) {
	if ptr.Pointer() != nil {
		C.QTextList_Remove(ptr.Pointer(), PointerFromQTextBlock(block))
	}
}

type QTextListFormat struct {
	QTextFormat
}

type QTextListFormat_ITF interface {
	QTextFormat_ITF
	QTextListFormat_PTR() *QTextListFormat
}

func (ptr *QTextListFormat) QTextListFormat_PTR() *QTextListFormat {
	return ptr
}

func (ptr *QTextListFormat) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFormat_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextListFormat) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextFormat_PTR().SetPointer(p)
	}
}

func PointerFromQTextListFormat(ptr QTextListFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextListFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextListFormatFromPointer(ptr unsafe.Pointer) (n *QTextListFormat) {
	n = new(QTextListFormat)
	n.SetPointer(ptr)
	return
}
func (ptr *QTextListFormat) DestroyQTextListFormat() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QTextListFormat__Style
//QTextListFormat::Style
type QTextListFormat__Style int64

var (
	QTextListFormat__ListDisc           QTextListFormat__Style = QTextListFormat__Style(-1)
	QTextListFormat__ListCircle         QTextListFormat__Style = QTextListFormat__Style(-2)
	QTextListFormat__ListSquare         QTextListFormat__Style = QTextListFormat__Style(-3)
	QTextListFormat__ListDecimal        QTextListFormat__Style = QTextListFormat__Style(-4)
	QTextListFormat__ListLowerAlpha     QTextListFormat__Style = QTextListFormat__Style(-5)
	QTextListFormat__ListUpperAlpha     QTextListFormat__Style = QTextListFormat__Style(-6)
	QTextListFormat__ListLowerRoman     QTextListFormat__Style = QTextListFormat__Style(-7)
	QTextListFormat__ListUpperRoman     QTextListFormat__Style = QTextListFormat__Style(-8)
	QTextListFormat__ListStyleUndefined QTextListFormat__Style = QTextListFormat__Style(0)
)

func NewQTextListFormat() *QTextListFormat {
	tmpValue := NewQTextListFormatFromPointer(C.QTextListFormat_NewQTextListFormat())
	qt.SetFinalizer(tmpValue, (*QTextListFormat).DestroyQTextListFormat)
	return tmpValue
}

func (ptr *QTextListFormat) Indent() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextListFormat_Indent(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextListFormat) SetIndent(indentation int) {
	if ptr.Pointer() != nil {
		C.QTextListFormat_SetIndent(ptr.Pointer(), C.int(int32(indentation)))
	}
}

func (ptr *QTextListFormat) SetStyle(style QTextListFormat__Style) {
	if ptr.Pointer() != nil {
		C.QTextListFormat_SetStyle(ptr.Pointer(), C.longlong(style))
	}
}

func (ptr *QTextListFormat) Style() QTextListFormat__Style {
	if ptr.Pointer() != nil {
		return QTextListFormat__Style(C.QTextListFormat_Style(ptr.Pointer()))
	}
	return 0
}

type QTextObject struct {
	core.QObject
}

type QTextObject_ITF interface {
	core.QObject_ITF
	QTextObject_PTR() *QTextObject
}

func (ptr *QTextObject) QTextObject_PTR() *QTextObject {
	return ptr
}

func (ptr *QTextObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQTextObject(ptr QTextObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextObject_PTR().Pointer()
	}
	return nil
}

func NewQTextObjectFromPointer(ptr unsafe.Pointer) (n *QTextObject) {
	n = new(QTextObject)
	n.SetPointer(ptr)
	return
}
func NewQTextObject(document QTextDocument_ITF) *QTextObject {
	tmpValue := NewQTextObjectFromPointer(C.QTextObject_NewQTextObject(PointerFromQTextDocument(document)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QTextObject) Document() *QTextDocument {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextDocumentFromPointer(C.QTextObject_Document(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextObject) Format() *QTextFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextFormatFromPointer(C.QTextObject_Format(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextFormat).DestroyQTextFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QTextObject) ObjectIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextObject_ObjectIndex(ptr.Pointer())))
	}
	return 0
}

//export callbackQTextObject_DestroyQTextObject
func callbackQTextObject_DestroyQTextObject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTextObject"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextObjectFromPointer(ptr).DestroyQTextObjectDefault()
	}
}

func (ptr *QTextObject) ConnectDestroyQTextObject(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTextObject"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTextObject", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTextObject", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTextObject) DisconnectDestroyQTextObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTextObject")
	}
}

func (ptr *QTextObject) DestroyQTextObject() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextObject_DestroyQTextObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextObject) DestroyQTextObjectDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextObject_DestroyQTextObjectDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextObject) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextObject___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextObject) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextObject___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextObject) __children_newList() unsafe.Pointer {
	return C.QTextObject___children_newList(ptr.Pointer())
}

func (ptr *QTextObject) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QTextObject___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTextObject) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTextObject___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QTextObject) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QTextObject___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QTextObject) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextObject___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextObject) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextObject___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextObject) __findChildren_newList() unsafe.Pointer {
	return C.QTextObject___findChildren_newList(ptr.Pointer())
}

func (ptr *QTextObject) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextObject___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextObject) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextObject___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextObject) __findChildren_newList3() unsafe.Pointer {
	return C.QTextObject___findChildren_newList3(ptr.Pointer())
}

//export callbackQTextObject_ChildEvent
func callbackQTextObject_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTextObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTextObject) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTextObject_ConnectNotify
func callbackQTextObject_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTextObjectFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTextObject) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTextObject_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTextObject_CustomEvent
func callbackQTextObject_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQTextObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextObject) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTextObject_DeleteLater
func callbackQTextObject_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTextObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTextObject) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextObject_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQTextObject_Destroyed
func callbackQTextObject_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQTextObject_DisconnectNotify
func callbackQTextObject_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTextObjectFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTextObject) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTextObject_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTextObject_Event
func callbackQTextObject_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTextObjectFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTextObject) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextObject_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQTextObject_EventFilter
func callbackQTextObject_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTextObjectFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTextObject) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextObject_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQTextObject_ObjectNameChanged
func callbackQTextObject_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGui_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQTextObject_TimerEvent
func callbackQTextObject_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTextObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTextObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QTextOption struct {
	ptr unsafe.Pointer
}

type QTextOption_ITF interface {
	QTextOption_PTR() *QTextOption
}

func (ptr *QTextOption) QTextOption_PTR() *QTextOption {
	return ptr
}

func (ptr *QTextOption) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextOption) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextOption(ptr QTextOption_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextOption_PTR().Pointer()
	}
	return nil
}

func NewQTextOptionFromPointer(ptr unsafe.Pointer) (n *QTextOption) {
	n = new(QTextOption)
	n.SetPointer(ptr)
	return
}

// QTextOption::Flag
//
//go:generate stringer -type=QTextOption__Flag
type QTextOption__Flag int64

const (
	QTextOption__ShowTabsAndSpaces                     QTextOption__Flag = QTextOption__Flag(0x1)
	QTextOption__ShowLineAndParagraphSeparators        QTextOption__Flag = QTextOption__Flag(0x2)
	QTextOption__AddSpaceForLineAndParagraphSeparators QTextOption__Flag = QTextOption__Flag(0x4)
	QTextOption__SuppressColors                        QTextOption__Flag = QTextOption__Flag(0x8)
	QTextOption__ShowDocumentTerminator                QTextOption__Flag = QTextOption__Flag(0x10)
	QTextOption__IncludeTrailingSpaces                 QTextOption__Flag = QTextOption__Flag(0x80000000)
)

// QTextOption::TabType
//
//go:generate stringer -type=QTextOption__TabType
type QTextOption__TabType int64

const (
	QTextOption__LeftTab      QTextOption__TabType = QTextOption__TabType(0)
	QTextOption__RightTab     QTextOption__TabType = QTextOption__TabType(1)
	QTextOption__CenterTab    QTextOption__TabType = QTextOption__TabType(2)
	QTextOption__DelimiterTab QTextOption__TabType = QTextOption__TabType(3)
)

// QTextOption::WrapMode
//
//go:generate stringer -type=QTextOption__WrapMode
type QTextOption__WrapMode int64

const (
	QTextOption__NoWrap                       QTextOption__WrapMode = QTextOption__WrapMode(0)
	QTextOption__WordWrap                     QTextOption__WrapMode = QTextOption__WrapMode(1)
	QTextOption__ManualWrap                   QTextOption__WrapMode = QTextOption__WrapMode(2)
	QTextOption__WrapAnywhere                 QTextOption__WrapMode = QTextOption__WrapMode(3)
	QTextOption__WrapAtWordBoundaryOrAnywhere QTextOption__WrapMode = QTextOption__WrapMode(4)
)

func NewQTextOption() *QTextOption {
	tmpValue := NewQTextOptionFromPointer(C.QTextOption_NewQTextOption())
	qt.SetFinalizer(tmpValue, (*QTextOption).DestroyQTextOption)
	return tmpValue
}

func NewQTextOption2(alignment core.Qt__AlignmentFlag) *QTextOption {
	tmpValue := NewQTextOptionFromPointer(C.QTextOption_NewQTextOption2(C.longlong(alignment)))
	qt.SetFinalizer(tmpValue, (*QTextOption).DestroyQTextOption)
	return tmpValue
}

func NewQTextOption3(other QTextOption_ITF) *QTextOption {
	tmpValue := NewQTextOptionFromPointer(C.QTextOption_NewQTextOption3(PointerFromQTextOption(other)))
	qt.SetFinalizer(tmpValue, (*QTextOption).DestroyQTextOption)
	return tmpValue
}

func (ptr *QTextOption) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextOption_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextOption) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QTextOption_SetAlignment(ptr.Pointer(), C.longlong(alignment))
	}
}

func (ptr *QTextOption) DestroyQTextOption() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextOption_DestroyQTextOption(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextOption) __setTabArray_tabStops_atList(i int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextOption___setTabArray_tabStops_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QTextOption) __setTabArray_tabStops_setList(i float64) {
	if ptr.Pointer() != nil {
		C.QTextOption___setTabArray_tabStops_setList(ptr.Pointer(), C.double(i))
	}
}

func (ptr *QTextOption) __setTabArray_tabStops_newList() unsafe.Pointer {
	return C.QTextOption___setTabArray_tabStops_newList(ptr.Pointer())
}

func (ptr *QTextOption) __tabArray_atList(i int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextOption___tabArray_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QTextOption) __tabArray_setList(i float64) {
	if ptr.Pointer() != nil {
		C.QTextOption___tabArray_setList(ptr.Pointer(), C.double(i))
	}
}

func (ptr *QTextOption) __tabArray_newList() unsafe.Pointer {
	return C.QTextOption___tabArray_newList(ptr.Pointer())
}

type QTextTable struct {
	QTextFrame
}

type QTextTable_ITF interface {
	QTextFrame_ITF
	QTextTable_PTR() *QTextTable
}

func (ptr *QTextTable) QTextTable_PTR() *QTextTable {
	return ptr
}

func (ptr *QTextTable) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFrame_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextTable) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextFrame_PTR().SetPointer(p)
	}
}

func PointerFromQTextTable(ptr QTextTable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTable_PTR().Pointer()
	}
	return nil
}

func NewQTextTableFromPointer(ptr unsafe.Pointer) (n *QTextTable) {
	n = new(QTextTable)
	n.SetPointer(ptr)
	return
}
func (ptr *QTextTable) AppendColumns(count int) {
	if ptr.Pointer() != nil {
		C.QTextTable_AppendColumns(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QTextTable) AppendRows(count int) {
	if ptr.Pointer() != nil {
		C.QTextTable_AppendRows(ptr.Pointer(), C.int(int32(count)))
	}
}

func (ptr *QTextTable) CellAt(row int, column int) *QTextTableCell {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextTableCellFromPointer(C.QTextTable_CellAt(ptr.Pointer(), C.int(int32(row)), C.int(int32(column))))
		qt.SetFinalizer(tmpValue, (*QTextTableCell).DestroyQTextTableCell)
		return tmpValue
	}
	return nil
}

func (ptr *QTextTable) CellAt2(position int) *QTextTableCell {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextTableCellFromPointer(C.QTextTable_CellAt2(ptr.Pointer(), C.int(int32(position))))
		qt.SetFinalizer(tmpValue, (*QTextTableCell).DestroyQTextTableCell)
		return tmpValue
	}
	return nil
}

func (ptr *QTextTable) CellAt3(cursor QTextCursor_ITF) *QTextTableCell {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextTableCellFromPointer(C.QTextTable_CellAt3(ptr.Pointer(), PointerFromQTextCursor(cursor)))
		qt.SetFinalizer(tmpValue, (*QTextTableCell).DestroyQTextTableCell)
		return tmpValue
	}
	return nil
}

func (ptr *QTextTable) Columns() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextTable_Columns(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTable) Format() *QTextTableFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextTableFormatFromPointer(C.QTextTable_Format(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextTableFormat).DestroyQTextTableFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QTextTable) InsertColumns(index int, columns int) {
	if ptr.Pointer() != nil {
		C.QTextTable_InsertColumns(ptr.Pointer(), C.int(int32(index)), C.int(int32(columns)))
	}
}

func (ptr *QTextTable) InsertRows(index int, rows int) {
	if ptr.Pointer() != nil {
		C.QTextTable_InsertRows(ptr.Pointer(), C.int(int32(index)), C.int(int32(rows)))
	}
}

func (ptr *QTextTable) RemoveColumns(index int, columns int) {
	if ptr.Pointer() != nil {
		C.QTextTable_RemoveColumns(ptr.Pointer(), C.int(int32(index)), C.int(int32(columns)))
	}
}

func (ptr *QTextTable) RemoveRows(index int, rows int) {
	if ptr.Pointer() != nil {
		C.QTextTable_RemoveRows(ptr.Pointer(), C.int(int32(index)), C.int(int32(rows)))
	}
}

func (ptr *QTextTable) Resize(rows int, columns int) {
	if ptr.Pointer() != nil {
		C.QTextTable_Resize(ptr.Pointer(), C.int(int32(rows)), C.int(int32(columns)))
	}
}

func (ptr *QTextTable) Rows() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextTable_Rows(ptr.Pointer())))
	}
	return 0
}

type QTextTableCell struct {
	ptr unsafe.Pointer
}

type QTextTableCell_ITF interface {
	QTextTableCell_PTR() *QTextTableCell
}

func (ptr *QTextTableCell) QTextTableCell_PTR() *QTextTableCell {
	return ptr
}

func (ptr *QTextTableCell) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextTableCell) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextTableCell(ptr QTextTableCell_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTableCell_PTR().Pointer()
	}
	return nil
}

func NewQTextTableCellFromPointer(ptr unsafe.Pointer) (n *QTextTableCell) {
	n = new(QTextTableCell)
	n.SetPointer(ptr)
	return
}
func NewQTextTableCell() *QTextTableCell {
	tmpValue := NewQTextTableCellFromPointer(C.QTextTableCell_NewQTextTableCell())
	qt.SetFinalizer(tmpValue, (*QTextTableCell).DestroyQTextTableCell)
	return tmpValue
}

func NewQTextTableCell2(other QTextTableCell_ITF) *QTextTableCell {
	tmpValue := NewQTextTableCellFromPointer(C.QTextTableCell_NewQTextTableCell2(PointerFromQTextTableCell(other)))
	qt.SetFinalizer(tmpValue, (*QTextTableCell).DestroyQTextTableCell)
	return tmpValue
}

func (ptr *QTextTableCell) Column() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextTableCell_Column(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableCell) FirstCursorPosition() *QTextCursor {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextCursorFromPointer(C.QTextTableCell_FirstCursorPosition(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextCursor).DestroyQTextCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QTextTableCell) Format() *QTextCharFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextCharFormatFromPointer(C.QTextTableCell_Format(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QTextCharFormat).DestroyQTextCharFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QTextTableCell) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextTableCell_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextTableCell) Row() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextTableCell_Row(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableCell) DestroyQTextTableCell() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTextTableCell_DestroyQTextTableCell(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTextTableFormat struct {
	QTextFrameFormat
}

type QTextTableFormat_ITF interface {
	QTextFrameFormat_ITF
	QTextTableFormat_PTR() *QTextTableFormat
}

func (ptr *QTextTableFormat) QTextTableFormat_PTR() *QTextTableFormat {
	return ptr
}

func (ptr *QTextTableFormat) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFrameFormat_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextTableFormat) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextFrameFormat_PTR().SetPointer(p)
	}
}

func PointerFromQTextTableFormat(ptr QTextTableFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextTableFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextTableFormatFromPointer(ptr unsafe.Pointer) (n *QTextTableFormat) {
	n = new(QTextTableFormat)
	n.SetPointer(ptr)
	return
}
func (ptr *QTextTableFormat) DestroyQTextTableFormat() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQTextTableFormat() *QTextTableFormat {
	tmpValue := NewQTextTableFormatFromPointer(C.QTextTableFormat_NewQTextTableFormat())
	qt.SetFinalizer(tmpValue, (*QTextTableFormat).DestroyQTextTableFormat)
	return tmpValue
}

func (ptr *QTextTableFormat) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextTableFormat_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextTableFormat) Columns() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QTextTableFormat_Columns(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextTableFormat) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QTextTableFormat_SetAlignment(ptr.Pointer(), C.longlong(alignment))
	}
}

func (ptr *QTextTableFormat) __columnWidthConstraints_atList(i int) *QTextLength {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextLengthFromPointer(C.QTextTableFormat___columnWidthConstraints_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QTextLength).DestroyQTextLength)
		return tmpValue
	}
	return nil
}

func (ptr *QTextTableFormat) __columnWidthConstraints_setList(i QTextLength_ITF) {
	if ptr.Pointer() != nil {
		C.QTextTableFormat___columnWidthConstraints_setList(ptr.Pointer(), PointerFromQTextLength(i))
	}
}

func (ptr *QTextTableFormat) __columnWidthConstraints_newList() unsafe.Pointer {
	return C.QTextTableFormat___columnWidthConstraints_newList(ptr.Pointer())
}

func (ptr *QTextTableFormat) __setColumnWidthConstraints_constraints_atList(i int) *QTextLength {
	if ptr.Pointer() != nil {
		tmpValue := NewQTextLengthFromPointer(C.QTextTableFormat___setColumnWidthConstraints_constraints_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QTextLength).DestroyQTextLength)
		return tmpValue
	}
	return nil
}

func (ptr *QTextTableFormat) __setColumnWidthConstraints_constraints_setList(i QTextLength_ITF) {
	if ptr.Pointer() != nil {
		C.QTextTableFormat___setColumnWidthConstraints_constraints_setList(ptr.Pointer(), PointerFromQTextLength(i))
	}
}

func (ptr *QTextTableFormat) __setColumnWidthConstraints_constraints_newList() unsafe.Pointer {
	return C.QTextTableFormat___setColumnWidthConstraints_constraints_newList(ptr.Pointer())
}

// QTouchDevice::CapabilityFlag
//
//go:generate stringer -type=QTouchDevice__CapabilityFlag
type QTouchDevice__CapabilityFlag int64

const (
	QTouchDevice__Position           QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0001)
	QTouchDevice__Area               QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0002)
	QTouchDevice__Pressure           QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0004)
	QTouchDevice__Velocity           QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0008)
	QTouchDevice__RawPositions       QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0010)
	QTouchDevice__NormalizedPosition QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0020)
	QTouchDevice__MouseEmulation     QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0040)
)

// QTouchDevice::DeviceType
//
//go:generate stringer -type=QTouchDevice__DeviceType
type QTouchDevice__DeviceType int64

const (
	QTouchDevice__TouchScreen QTouchDevice__DeviceType = QTouchDevice__DeviceType(0)
	QTouchDevice__TouchPad    QTouchDevice__DeviceType = QTouchDevice__DeviceType(1)
)

type QTransform struct {
	ptr unsafe.Pointer
}

type QTransform_ITF interface {
	QTransform_PTR() *QTransform
}

func (ptr *QTransform) QTransform_PTR() *QTransform {
	return ptr
}

func (ptr *QTransform) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTransform) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTransform(ptr QTransform_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTransform_PTR().Pointer()
	}
	return nil
}

func NewQTransformFromPointer(ptr unsafe.Pointer) (n *QTransform) {
	n = new(QTransform)
	n.SetPointer(ptr)
	return
}
func (ptr *QTransform) DestroyQTransform() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QTransform__TransformationType
//QTransform::TransformationType
type QTransform__TransformationType int64

const (
	QTransform__TxNone      QTransform__TransformationType = QTransform__TransformationType(0x00)
	QTransform__TxTranslate QTransform__TransformationType = QTransform__TransformationType(0x01)
	QTransform__TxScale     QTransform__TransformationType = QTransform__TransformationType(0x02)
	QTransform__TxRotate    QTransform__TransformationType = QTransform__TransformationType(0x04)
	QTransform__TxShear     QTransform__TransformationType = QTransform__TransformationType(0x08)
	QTransform__TxProject   QTransform__TransformationType = QTransform__TransformationType(0x10)
)

func NewQTransform2() *QTransform {
	tmpValue := NewQTransformFromPointer(C.QTransform_NewQTransform2())
	qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
	return tmpValue
}

func NewQTransform3(m11 float64, m12 float64, m13 float64, m21 float64, m22 float64, m23 float64, m31 float64, m32 float64, m33 float64) *QTransform {
	tmpValue := NewQTransformFromPointer(C.QTransform_NewQTransform3(C.double(m11), C.double(m12), C.double(m13), C.double(m21), C.double(m22), C.double(m23), C.double(m31), C.double(m32), C.double(m33)))
	qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
	return tmpValue
}

func NewQTransform4(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) *QTransform {
	tmpValue := NewQTransformFromPointer(C.QTransform_NewQTransform4(C.double(m11), C.double(m12), C.double(m21), C.double(m22), C.double(dx), C.double(dy)))
	qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
	return tmpValue
}

func (ptr *QTransform) Map(x float64, y float64, tx float64, ty float64) {
	if ptr.Pointer() != nil {
		C.QTransform_Map(ptr.Pointer(), C.double(x), C.double(y), C.double(tx), C.double(ty))
	}
}

func (ptr *QTransform) Map2(point core.QPoint_ITF) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QTransform_Map2(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map3(p core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QTransform_Map3(ptr.Pointer(), core.PointerFromQPointF(p)))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map4(l core.QLine_ITF) *core.QLine {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLineFromPointer(C.QTransform_Map4(ptr.Pointer(), core.PointerFromQLine(l)))
		qt.SetFinalizer(tmpValue, (*core.QLine).DestroyQLine)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map5(line core.QLineF_ITF) *core.QLineF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLineFFromPointer(C.QTransform_Map5(ptr.Pointer(), core.PointerFromQLineF(line)))
		qt.SetFinalizer(tmpValue, (*core.QLineF).DestroyQLineF)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map6(polygon QPolygonF_ITF) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QTransform_Map6(ptr.Pointer(), PointerFromQPolygonF(polygon)))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map7(polygon QPolygon_ITF) *QPolygon {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFromPointer(C.QTransform_Map7(ptr.Pointer(), PointerFromQPolygon(polygon)))
		qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map8(region QRegion_ITF) *QRegion {
	if ptr.Pointer() != nil {
		tmpValue := NewQRegionFromPointer(C.QTransform_Map8(ptr.Pointer(), PointerFromQRegion(region)))
		qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map9(path QPainterPath_ITF) *QPainterPath {
	if ptr.Pointer() != nil {
		tmpValue := NewQPainterPathFromPointer(C.QTransform_Map9(ptr.Pointer(), PointerFromQPainterPath(path)))
		qt.SetFinalizer(tmpValue, (*QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map10(x int, y int, tx int, ty int) {
	if ptr.Pointer() != nil {
		C.QTransform_Map10(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(tx)), C.int(int32(ty)))
	}
}

func (ptr *QTransform) Scale(sx float64, sy float64) *QTransform {
	if ptr.Pointer() != nil {
		tmpValue := NewQTransformFromPointer(C.QTransform_Scale(ptr.Pointer(), C.double(sx), C.double(sy)))
		qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Type() QTransform__TransformationType {
	if ptr.Pointer() != nil {
		return QTransform__TransformationType(C.QTransform_Type(ptr.Pointer()))
	}
	return 0
}

type QValidator struct {
	core.QObject
}

type QValidator_ITF interface {
	core.QObject_ITF
	QValidator_PTR() *QValidator
}

func (ptr *QValidator) QValidator_PTR() *QValidator {
	return ptr
}

func (ptr *QValidator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QValidator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQValidator(ptr QValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QValidator_PTR().Pointer()
	}
	return nil
}

func NewQValidatorFromPointer(ptr unsafe.Pointer) (n *QValidator) {
	n = new(QValidator)
	n.SetPointer(ptr)
	return
}

// QValidator::State
//
//go:generate stringer -type=QValidator__State
type QValidator__State int64

const (
	QValidator__Invalid      QValidator__State = QValidator__State(0)
	QValidator__Intermediate QValidator__State = QValidator__State(1)
	QValidator__Acceptable   QValidator__State = QValidator__State(2)
)

func NewQValidator(parent core.QObject_ITF) *QValidator {
	tmpValue := NewQValidatorFromPointer(C.QValidator_NewQValidator(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQValidator_Changed
func callbackQValidator_Changed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changed"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QValidator) ConnectChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "changed") {
			C.QValidator_ConnectChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "changed")))
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

func (ptr *QValidator) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.QValidator_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "changed")
	}
}

func (ptr *QValidator) Changed() {
	if ptr.Pointer() != nil {
		C.QValidator_Changed(ptr.Pointer())
	}
}

//export callbackQValidator_Validate
func callbackQValidator_Validate(ptr unsafe.Pointer, input C.struct_QtGui_PackedString, pos C.int) C.longlong {
	if signal := qt.GetSignal(ptr, "validate"); signal != nil {
		return C.longlong((*(*func(string, int) QValidator__State)(signal))(cGoUnpackString(input), int(int32(pos))))
	}

	return C.longlong(0)
}

func (ptr *QValidator) ConnectValidate(f func(input string, pos int) QValidator__State) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "validate"); signal != nil {
			f := func(input string, pos int) QValidator__State {
				(*(*func(string, int) QValidator__State)(signal))(input, pos)
				return f(input, pos)
			}
			qt.ConnectSignal(ptr.Pointer(), "validate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "validate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValidator) DisconnectValidate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "validate")
	}
}

func (ptr *QValidator) Validate(input string, pos int) QValidator__State {
	if ptr.Pointer() != nil {
		var inputC *C.char
		if input != "" {
			inputC = C.CString(input)
			defer C.free(unsafe.Pointer(inputC))
		}
		return QValidator__State(C.QValidator_Validate(ptr.Pointer(), C.struct_QtGui_PackedString{data: inputC, len: C.longlong(len(input))}, C.int(int32(pos))))
	}
	return 0
}

//export callbackQValidator_DestroyQValidator
func callbackQValidator_DestroyQValidator(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QValidator"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQValidatorFromPointer(ptr).DestroyQValidatorDefault()
	}
}

func (ptr *QValidator) ConnectDestroyQValidator(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QValidator"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QValidator", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QValidator", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QValidator) DisconnectDestroyQValidator() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QValidator")
	}
}

func (ptr *QValidator) DestroyQValidator() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QValidator_DestroyQValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QValidator) DestroyQValidatorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QValidator_DestroyQValidatorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QValidator) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QValidator___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QValidator) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QValidator___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QValidator) __children_newList() unsafe.Pointer {
	return C.QValidator___children_newList(ptr.Pointer())
}

func (ptr *QValidator) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QValidator___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QValidator) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QValidator___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QValidator) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QValidator___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QValidator) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QValidator___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QValidator) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QValidator___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QValidator) __findChildren_newList() unsafe.Pointer {
	return C.QValidator___findChildren_newList(ptr.Pointer())
}

func (ptr *QValidator) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QValidator___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QValidator) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QValidator___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QValidator) __findChildren_newList3() unsafe.Pointer {
	return C.QValidator___findChildren_newList3(ptr.Pointer())
}

//export callbackQValidator_ChildEvent
func callbackQValidator_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQValidatorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QValidator) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QValidator_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQValidator_ConnectNotify
func callbackQValidator_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQValidatorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QValidator) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QValidator_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQValidator_CustomEvent
func callbackQValidator_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQValidatorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QValidator) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QValidator_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQValidator_DeleteLater
func callbackQValidator_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQValidatorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QValidator) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QValidator_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQValidator_Destroyed
func callbackQValidator_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQValidator_DisconnectNotify
func callbackQValidator_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQValidatorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QValidator) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QValidator_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQValidator_Event
func callbackQValidator_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQValidatorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QValidator) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QValidator_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQValidator_EventFilter
func callbackQValidator_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQValidatorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QValidator) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QValidator_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQValidator_ObjectNameChanged
func callbackQValidator_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGui_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQValidator_TimerEvent
func callbackQValidator_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQValidatorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QValidator) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QValidator_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QVector2D struct {
	ptr unsafe.Pointer
}

type QVector2D_ITF interface {
	QVector2D_PTR() *QVector2D
}

func (ptr *QVector2D) QVector2D_PTR() *QVector2D {
	return ptr
}

func (ptr *QVector2D) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVector2D) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVector2D(ptr QVector2D_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector2D_PTR().Pointer()
	}
	return nil
}

func NewQVector2DFromPointer(ptr unsafe.Pointer) (n *QVector2D) {
	n = new(QVector2D)
	n.SetPointer(ptr)
	return
}
func (ptr *QVector2D) DestroyQVector2D() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQVector2D() *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D())
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func NewQVector2D3(xpos float32, ypos float32) *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D3(C.float(xpos), C.float(ypos)))
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func NewQVector2D4(point core.QPoint_ITF) *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D4(core.PointerFromQPoint(point)))
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func NewQVector2D5(point core.QPointF_ITF) *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D5(core.PointerFromQPointF(point)))
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func NewQVector2D6(vector QVector3D_ITF) *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D6(PointerFromQVector3D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func NewQVector2D7(vector QVector4D_ITF) *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D7(PointerFromQVector4D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func (ptr *QVector2D) Length() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector2D_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector2D) X() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector2D_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector2D) Y() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector2D_Y(ptr.Pointer()))
	}
	return 0
}

type QVector3D struct {
	ptr unsafe.Pointer
}

type QVector3D_ITF interface {
	QVector3D_PTR() *QVector3D
}

func (ptr *QVector3D) QVector3D_PTR() *QVector3D {
	return ptr
}

func (ptr *QVector3D) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVector3D) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVector3D(ptr QVector3D_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector3D_PTR().Pointer()
	}
	return nil
}

func NewQVector3DFromPointer(ptr unsafe.Pointer) (n *QVector3D) {
	n = new(QVector3D)
	n.SetPointer(ptr)
	return
}
func (ptr *QVector3D) DestroyQVector3D() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQVector3D() *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D())
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D3(xpos float32, ypos float32, zpos float32) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D3(C.float(xpos), C.float(ypos), C.float(zpos)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D4(point core.QPoint_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D4(core.PointerFromQPoint(point)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D5(point core.QPointF_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D5(core.PointerFromQPointF(point)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D6(vector QVector2D_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D6(PointerFromQVector2D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D7(vector QVector2D_ITF, zpos float32) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D7(PointerFromQVector2D(vector), C.float(zpos)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D8(vector QVector4D_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D8(PointerFromQVector4D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func (ptr *QVector3D) Length() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector3D_Length(ptr.Pointer()))
	}
	return 0
}

func QVector3D_Normal(v1 QVector3D_ITF, v2 QVector3D_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_QVector3D_Normal(PointerFromQVector3D(v1), PointerFromQVector3D(v2)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func (ptr *QVector3D) Normal(v1 QVector3D_ITF, v2 QVector3D_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_QVector3D_Normal(PointerFromQVector3D(v1), PointerFromQVector3D(v2)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func QVector3D_Normal2(v1 QVector3D_ITF, v2 QVector3D_ITF, v3 QVector3D_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_QVector3D_Normal2(PointerFromQVector3D(v1), PointerFromQVector3D(v2), PointerFromQVector3D(v3)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func (ptr *QVector3D) Normal2(v1 QVector3D_ITF, v2 QVector3D_ITF, v3 QVector3D_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_QVector3D_Normal2(PointerFromQVector3D(v1), PointerFromQVector3D(v2), PointerFromQVector3D(v3)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func (ptr *QVector3D) X() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector3D_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector3D) Y() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector3D_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector3D) Z() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector3D_Z(ptr.Pointer()))
	}
	return 0
}

type QVector4D struct {
	ptr unsafe.Pointer
}

type QVector4D_ITF interface {
	QVector4D_PTR() *QVector4D
}

func (ptr *QVector4D) QVector4D_PTR() *QVector4D {
	return ptr
}

func (ptr *QVector4D) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVector4D) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVector4D(ptr QVector4D_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector4D_PTR().Pointer()
	}
	return nil
}

func NewQVector4DFromPointer(ptr unsafe.Pointer) (n *QVector4D) {
	n = new(QVector4D)
	n.SetPointer(ptr)
	return
}
func (ptr *QVector4D) DestroyQVector4D() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQVector4D() *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D())
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D3(xpos float32, ypos float32, zpos float32, wpos float32) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D3(C.float(xpos), C.float(ypos), C.float(zpos), C.float(wpos)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D4(point core.QPoint_ITF) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D4(core.PointerFromQPoint(point)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D5(point core.QPointF_ITF) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D5(core.PointerFromQPointF(point)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D6(vector QVector2D_ITF) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D6(PointerFromQVector2D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D7(vector QVector2D_ITF, zpos float32, wpos float32) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D7(PointerFromQVector2D(vector), C.float(zpos), C.float(wpos)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D8(vector QVector3D_ITF) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D8(PointerFromQVector3D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D9(vector QVector3D_ITF, wpos float32) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D9(PointerFromQVector3D(vector), C.float(wpos)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func (ptr *QVector4D) Length() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector4D_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector4D) SetW(w float32) {
	if ptr.Pointer() != nil {
		C.QVector4D_SetW(ptr.Pointer(), C.float(w))
	}
}

func (ptr *QVector4D) W() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector4D_W(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector4D) X() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector4D_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector4D) Y() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector4D_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector4D) Z() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector4D_Z(ptr.Pointer()))
	}
	return 0
}

// QVulkanInstance::Flag
//
//go:generate stringer -type=QVulkanInstance__Flag
type QVulkanInstance__Flag int64

const (
	QVulkanInstance__NoDebugOutputRedirect QVulkanInstance__Flag = QVulkanInstance__Flag(0x01)
)

// QVulkanWindow::Flag
//
//go:generate stringer -type=QVulkanWindow__Flag
type QVulkanWindow__Flag int64

const (
	QVulkanWindow__PersistentResources QVulkanWindow__Flag = QVulkanWindow__Flag(0x01)
)

type QWindow struct {
	core.QObject
	QSurface
}

type QWindow_ITF interface {
	core.QObject_ITF
	QSurface_ITF
	QWindow_PTR() *QWindow
}

func (ptr *QWindow) QWindow_PTR() *QWindow {
	return ptr
}

func (ptr *QWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QSurface_PTR().SetPointer(p)
	}
}

func PointerFromQWindow(ptr QWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWindow_PTR().Pointer()
	}
	return nil
}

func NewQWindowFromPointer(ptr unsafe.Pointer) (n *QWindow) {
	n = new(QWindow)
	n.SetPointer(ptr)
	return
}

// QWindow::AncestorMode
//
//go:generate stringer -type=QWindow__AncestorMode
type QWindow__AncestorMode int64

const (
	QWindow__ExcludeTransients QWindow__AncestorMode = QWindow__AncestorMode(0)
	QWindow__IncludeTransients QWindow__AncestorMode = QWindow__AncestorMode(1)
)

// QWindow::Visibility
//
//go:generate stringer -type=QWindow__Visibility
type QWindow__Visibility int64

const (
	QWindow__Hidden              QWindow__Visibility = QWindow__Visibility(0)
	QWindow__AutomaticVisibility QWindow__Visibility = QWindow__Visibility(1)
	QWindow__Windowed            QWindow__Visibility = QWindow__Visibility(2)
	QWindow__Minimized           QWindow__Visibility = QWindow__Visibility(3)
	QWindow__Maximized           QWindow__Visibility = QWindow__Visibility(4)
	QWindow__FullScreen          QWindow__Visibility = QWindow__Visibility(5)
)

func NewQWindow(targetScreen QScreen_ITF) *QWindow {
	tmpValue := NewQWindowFromPointer(C.QWindow_NewQWindow(PointerFromQScreen(targetScreen)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQWindow2(parent QWindow_ITF) *QWindow {
	tmpValue := NewQWindowFromPointer(C.QWindow_NewQWindow2(PointerFromQWindow(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWindow_Close
func callbackQWindow_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWindowFromPointer(ptr).CloseDefault())))
}

func (ptr *QWindow) ConnectClose(f func() bool) {
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

func (ptr *QWindow) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QWindow) Close() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWindow_Close(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWindow) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWindow_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWindow) Create() {
	if ptr.Pointer() != nil {
		C.QWindow_Create(ptr.Pointer())
	}
}

func (ptr *QWindow) Cursor() *QCursor {
	if ptr.Pointer() != nil {
		tmpValue := NewQCursorFromPointer(C.QWindow_Cursor(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QCursor).DestroyQCursor)
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) Destroy() {
	if ptr.Pointer() != nil {
		C.QWindow_Destroy(ptr.Pointer())
	}
}

//export callbackQWindow_Event
func callbackQWindow_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWindowFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev)))))
}

func (ptr *QWindow) ConnectEvent(f func(ev *core.QEvent) bool) {
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

func (ptr *QWindow) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "event")
	}
}

func (ptr *QWindow) Event(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWindow_Event(ptr.Pointer(), core.PointerFromQEvent(ev))) != 0
	}
	return false
}

func (ptr *QWindow) EventDefault(ev core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWindow_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))) != 0
	}
	return false
}

//export callbackQWindow_Format
func callbackQWindow_Format(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "format"); signal != nil {
		return PointerFromQSurfaceFormat((*(*func() *QSurfaceFormat)(signal))())
	}

	return PointerFromQSurfaceFormat(NewQWindowFromPointer(ptr).FormatDefault())
}

func (ptr *QWindow) ConnectFormat(f func() *QSurfaceFormat) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "format"); signal != nil {
			f := func() *QSurfaceFormat {
				(*(*func() *QSurfaceFormat)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "format", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "format", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectFormat() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "format")
	}
}

func (ptr *QWindow) Format() *QSurfaceFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQSurfaceFormatFromPointer(C.QWindow_Format(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) FormatDefault() *QSurfaceFormat {
	if ptr.Pointer() != nil {
		tmpValue := NewQSurfaceFormatFromPointer(C.QWindow_FormatDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QSurfaceFormat).DestroyQSurfaceFormat)
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) Geometry() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QWindow_Geometry(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) Height() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWindow_Height(ptr.Pointer())))
	}
	return 0
}

//export callbackQWindow_Hide
func callbackQWindow_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWindowFromPointer(ptr).HideDefault()
	}
}

func (ptr *QWindow) ConnectHide(f func()) {
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

func (ptr *QWindow) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hide")
	}
}

func (ptr *QWindow) Hide() {
	if ptr.Pointer() != nil {
		C.QWindow_Hide(ptr.Pointer())
	}
}

func (ptr *QWindow) HideDefault() {
	if ptr.Pointer() != nil {
		C.QWindow_HideDefault(ptr.Pointer())
	}
}

//export callbackQWindow_HideEvent
func callbackQWindow_HideEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*QHideEvent))(signal))(NewQHideEventFromPointer(ev))
	} else {
		NewQWindowFromPointer(ptr).HideEventDefault(NewQHideEventFromPointer(ev))
	}
}

func (ptr *QWindow) ConnectHideEvent(f func(ev *QHideEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hideEvent"); signal != nil {
			f := func(ev *QHideEvent) {
				(*(*func(*QHideEvent))(signal))(ev)
				f(ev)
			}
			qt.ConnectSignal(ptr.Pointer(), "hideEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hideEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectHideEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hideEvent")
	}
}

func (ptr *QWindow) HideEvent(ev QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_HideEvent(ptr.Pointer(), PointerFromQHideEvent(ev))
	}
}

func (ptr *QWindow) HideEventDefault(ev QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_HideEventDefault(ptr.Pointer(), PointerFromQHideEvent(ev))
	}
}

func (ptr *QWindow) Icon() *QIcon {
	if ptr.Pointer() != nil {
		tmpValue := NewQIconFromPointer(C.QWindow_Icon(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

//export callbackQWindow_KeyPressEvent
func callbackQWindow_KeyPressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*QKeyEvent))(signal))(NewQKeyEventFromPointer(ev))
	} else {
		NewQWindowFromPointer(ptr).KeyPressEventDefault(NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QWindow) ConnectKeyPressEvent(f func(ev *QKeyEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keyPressEvent"); signal != nil {
			f := func(ev *QKeyEvent) {
				(*(*func(*QKeyEvent))(signal))(ev)
				f(ev)
			}
			qt.ConnectSignal(ptr.Pointer(), "keyPressEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keyPressEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectKeyPressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "keyPressEvent")
	}
}

func (ptr *QWindow) KeyPressEvent(ev QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_KeyPressEvent(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
}

func (ptr *QWindow) KeyPressEventDefault(ev QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_KeyPressEventDefault(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
}

//export callbackQWindow_KeyReleaseEvent
func callbackQWindow_KeyReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*QKeyEvent))(signal))(NewQKeyEventFromPointer(ev))
	} else {
		NewQWindowFromPointer(ptr).KeyReleaseEventDefault(NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QWindow) ConnectKeyReleaseEvent(f func(ev *QKeyEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keyReleaseEvent"); signal != nil {
			f := func(ev *QKeyEvent) {
				(*(*func(*QKeyEvent))(signal))(ev)
				f(ev)
			}
			qt.ConnectSignal(ptr.Pointer(), "keyReleaseEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keyReleaseEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectKeyReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "keyReleaseEvent")
	}
}

func (ptr *QWindow) KeyReleaseEvent(ev QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_KeyReleaseEvent(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
}

func (ptr *QWindow) KeyReleaseEventDefault(ev QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_KeyReleaseEventDefault(ptr.Pointer(), PointerFromQKeyEvent(ev))
	}
}

//export callbackQWindow_Lower
func callbackQWindow_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWindowFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QWindow) ConnectLower(f func()) {
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

func (ptr *QWindow) DisconnectLower() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lower")
	}
}

func (ptr *QWindow) Lower() {
	if ptr.Pointer() != nil {
		C.QWindow_Lower(ptr.Pointer())
	}
}

func (ptr *QWindow) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QWindow_LowerDefault(ptr.Pointer())
	}
}

func (ptr *QWindow) MaximumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWindow_MaximumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) MaximumWidth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWindow_MaximumWidth(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) MinimumHeight() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWindow_MinimumHeight(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) MinimumSize() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWindow_MinimumSize(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) MinimumWidth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWindow_MinimumWidth(ptr.Pointer())))
	}
	return 0
}

//export callbackQWindow_MouseDoubleClickEvent
func callbackQWindow_MouseDoubleClickEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*QMouseEvent))(signal))(NewQMouseEventFromPointer(ev))
	} else {
		NewQWindowFromPointer(ptr).MouseDoubleClickEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWindow) ConnectMouseDoubleClickEvent(f func(ev *QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mouseDoubleClickEvent"); signal != nil {
			f := func(ev *QMouseEvent) {
				(*(*func(*QMouseEvent))(signal))(ev)
				f(ev)
			}
			qt.ConnectSignal(ptr.Pointer(), "mouseDoubleClickEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mouseDoubleClickEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectMouseDoubleClickEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mouseDoubleClickEvent")
	}
}

func (ptr *QWindow) MouseDoubleClickEvent(ev QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_MouseDoubleClickEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QWindow) MouseDoubleClickEventDefault(ev QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_MouseDoubleClickEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

//export callbackQWindow_MousePressEvent
func callbackQWindow_MousePressEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*QMouseEvent))(signal))(NewQMouseEventFromPointer(ev))
	} else {
		NewQWindowFromPointer(ptr).MousePressEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWindow) ConnectMousePressEvent(f func(ev *QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mousePressEvent"); signal != nil {
			f := func(ev *QMouseEvent) {
				(*(*func(*QMouseEvent))(signal))(ev)
				f(ev)
			}
			qt.ConnectSignal(ptr.Pointer(), "mousePressEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mousePressEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectMousePressEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mousePressEvent")
	}
}

func (ptr *QWindow) MousePressEvent(ev QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_MousePressEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QWindow) MousePressEventDefault(ev QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_MousePressEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

//export callbackQWindow_MouseReleaseEvent
func callbackQWindow_MouseReleaseEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*QMouseEvent))(signal))(NewQMouseEventFromPointer(ev))
	} else {
		NewQWindowFromPointer(ptr).MouseReleaseEventDefault(NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWindow) ConnectMouseReleaseEvent(f func(ev *QMouseEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "mouseReleaseEvent"); signal != nil {
			f := func(ev *QMouseEvent) {
				(*(*func(*QMouseEvent))(signal))(ev)
				f(ev)
			}
			qt.ConnectSignal(ptr.Pointer(), "mouseReleaseEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mouseReleaseEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectMouseReleaseEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "mouseReleaseEvent")
	}
}

func (ptr *QWindow) MouseReleaseEvent(ev QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_MouseReleaseEvent(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QWindow) MouseReleaseEventDefault(ev QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_MouseReleaseEventDefault(ptr.Pointer(), PointerFromQMouseEvent(ev))
	}
}

func (ptr *QWindow) Parent(mode QWindow__AncestorMode) *QWindow {
	if ptr.Pointer() != nil {
		tmpValue := NewQWindowFromPointer(C.QWindow_Parent(ptr.Pointer(), C.longlong(mode)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) Parent2() *QWindow {
	if ptr.Pointer() != nil {
		tmpValue := NewQWindowFromPointer(C.QWindow_Parent2(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) Position() *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QWindow_Position(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) Resize(newSize core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_Resize(ptr.Pointer(), core.PointerFromQSize(newSize))
	}
}

func (ptr *QWindow) Resize2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWindow_Resize2(ptr.Pointer(), C.int(int32(w)), C.int(int32(h)))
	}
}

//export callbackQWindow_ResizeEvent
func callbackQWindow_ResizeEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*QResizeEvent))(signal))(NewQResizeEventFromPointer(ev))
	} else {
		NewQWindowFromPointer(ptr).ResizeEventDefault(NewQResizeEventFromPointer(ev))
	}
}

func (ptr *QWindow) ConnectResizeEvent(f func(ev *QResizeEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resizeEvent"); signal != nil {
			f := func(ev *QResizeEvent) {
				(*(*func(*QResizeEvent))(signal))(ev)
				f(ev)
			}
			qt.ConnectSignal(ptr.Pointer(), "resizeEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resizeEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectResizeEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resizeEvent")
	}
}

func (ptr *QWindow) ResizeEvent(ev QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_ResizeEvent(ptr.Pointer(), PointerFromQResizeEvent(ev))
	}
}

func (ptr *QWindow) ResizeEventDefault(ev QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_ResizeEventDefault(ptr.Pointer(), PointerFromQResizeEvent(ev))
	}
}

//export callbackQWindow_SetGeometry
func callbackQWindow_SetGeometry(ptr unsafe.Pointer, posx C.int, posy C.int, w C.int, h C.int) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		(*(*func(int, int, int, int))(signal))(int(int32(posx)), int(int32(posy)), int(int32(w)), int(int32(h)))
	} else {
		NewQWindowFromPointer(ptr).SetGeometryDefault(int(int32(posx)), int(int32(posy)), int(int32(w)), int(int32(h)))
	}
}

func (ptr *QWindow) ConnectSetGeometry(f func(posx int, posy int, w int, h int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGeometry"); signal != nil {
			f := func(posx int, posy int, w int, h int) {
				(*(*func(int, int, int, int))(signal))(posx, posy, w, h)
				f(posx, posy, w, h)
			}
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGeometry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectSetGeometry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGeometry")
	}
}

func (ptr *QWindow) SetGeometry(posx int, posy int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetGeometry(ptr.Pointer(), C.int(int32(posx)), C.int(int32(posy)), C.int(int32(w)), C.int(int32(h)))
	}
}

func (ptr *QWindow) SetGeometryDefault(posx int, posy int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetGeometryDefault(ptr.Pointer(), C.int(int32(posx)), C.int(int32(posy)), C.int(int32(w)), C.int(int32(h)))
	}
}

//export callbackQWindow_SetGeometry2
func callbackQWindow_SetGeometry2(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry2"); signal != nil {
		(*(*func(*core.QRect))(signal))(core.NewQRectFromPointer(rect))
	} else {
		NewQWindowFromPointer(ptr).SetGeometry2Default(core.NewQRectFromPointer(rect))
	}
}

func (ptr *QWindow) ConnectSetGeometry2(f func(rect *core.QRect)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGeometry2"); signal != nil {
			f := func(rect *core.QRect) {
				(*(*func(*core.QRect))(signal))(rect)
				f(rect)
			}
			qt.ConnectSignal(ptr.Pointer(), "setGeometry2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGeometry2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectSetGeometry2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGeometry2")
	}
}

func (ptr *QWindow) SetGeometry2(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetGeometry2(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWindow) SetGeometry2Default(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetGeometry2Default(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWindow) SetIcon(icon QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetIcon(ptr.Pointer(), PointerFromQIcon(icon))
	}
}

//export callbackQWindow_SetMaximumWidth
func callbackQWindow_SetMaximumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(ptr, "setMaximumWidth"); signal != nil {
		(*(*func(int))(signal))(int(int32(w)))
	} else {
		NewQWindowFromPointer(ptr).SetMaximumWidthDefault(int(int32(w)))
	}
}

func (ptr *QWindow) ConnectSetMaximumWidth(f func(w int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMaximumWidth"); signal != nil {
			f := func(w int) {
				(*(*func(int))(signal))(w)
				f(w)
			}
			qt.ConnectSignal(ptr.Pointer(), "setMaximumWidth", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMaximumWidth", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectSetMaximumWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMaximumWidth")
	}
}

func (ptr *QWindow) SetMaximumWidth(w int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumWidth(ptr.Pointer(), C.int(int32(w)))
	}
}

func (ptr *QWindow) SetMaximumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMaximumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

func (ptr *QWindow) SetMinimumSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

//export callbackQWindow_SetMinimumWidth
func callbackQWindow_SetMinimumWidth(ptr unsafe.Pointer, w C.int) {
	if signal := qt.GetSignal(ptr, "setMinimumWidth"); signal != nil {
		(*(*func(int))(signal))(int(int32(w)))
	} else {
		NewQWindowFromPointer(ptr).SetMinimumWidthDefault(int(int32(w)))
	}
}

func (ptr *QWindow) ConnectSetMinimumWidth(f func(w int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMinimumWidth"); signal != nil {
			f := func(w int) {
				(*(*func(int))(signal))(w)
				f(w)
			}
			qt.ConnectSignal(ptr.Pointer(), "setMinimumWidth", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMinimumWidth", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectSetMinimumWidth() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMinimumWidth")
	}
}

func (ptr *QWindow) SetMinimumWidth(w int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumWidth(ptr.Pointer(), C.int(int32(w)))
	}
}

func (ptr *QWindow) SetMinimumWidthDefault(w int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetMinimumWidthDefault(ptr.Pointer(), C.int(int32(w)))
	}
}

func (ptr *QWindow) SetPosition(pt core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_SetPosition(ptr.Pointer(), core.PointerFromQPoint(pt))
	}
}

func (ptr *QWindow) SetPosition2(posx int, posy int) {
	if ptr.Pointer() != nil {
		C.QWindow_SetPosition2(ptr.Pointer(), C.int(int32(posx)), C.int(int32(posy)))
	}
}

//export callbackQWindow_Show
func callbackQWindow_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWindowFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWindow) ConnectShow(f func()) {
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

func (ptr *QWindow) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "show")
	}
}

func (ptr *QWindow) Show() {
	if ptr.Pointer() != nil {
		C.QWindow_Show(ptr.Pointer())
	}
}

func (ptr *QWindow) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowDefault(ptr.Pointer())
	}
}

//export callbackQWindow_ShowEvent
func callbackQWindow_ShowEvent(ptr unsafe.Pointer, ev unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*QShowEvent))(signal))(NewQShowEventFromPointer(ev))
	} else {
		NewQWindowFromPointer(ptr).ShowEventDefault(NewQShowEventFromPointer(ev))
	}
}

func (ptr *QWindow) ConnectShowEvent(f func(ev *QShowEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "showEvent"); signal != nil {
			f := func(ev *QShowEvent) {
				(*(*func(*QShowEvent))(signal))(ev)
				f(ev)
			}
			qt.ConnectSignal(ptr.Pointer(), "showEvent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "showEvent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectShowEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "showEvent")
	}
}

func (ptr *QWindow) ShowEvent(ev QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_ShowEvent(ptr.Pointer(), PointerFromQShowEvent(ev))
	}
}

func (ptr *QWindow) ShowEventDefault(ev QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_ShowEventDefault(ptr.Pointer(), PointerFromQShowEvent(ev))
	}
}

//export callbackQWindow_ShowMaximized
func callbackQWindow_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWindowFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QWindow) ConnectShowMaximized(f func()) {
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

func (ptr *QWindow) DisconnectShowMaximized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "showMaximized")
	}
}

func (ptr *QWindow) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QWindow) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QWindow_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQWindow_Size
func callbackQWindow_Size(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return core.PointerFromQSize((*(*func() *core.QSize)(signal))())
	}

	return core.PointerFromQSize(NewQWindowFromPointer(ptr).SizeDefault())
}

func (ptr *QWindow) ConnectSize(f func() *core.QSize) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "size"); signal != nil {
			f := func() *core.QSize {
				(*(*func() *core.QSize)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "size", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "size", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectSize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "size")
	}
}

func (ptr *QWindow) Size() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWindow_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) SizeDefault() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QWindow_SizeDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWindow_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWindow) Type() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWindow_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) Width() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWindow_Width(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) X() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWindow_X(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWindow) Y() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWindow_Y(ptr.Pointer())))
	}
	return 0
}

//export callbackQWindow_DestroyQWindow
func callbackQWindow_DestroyQWindow(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWindow"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWindowFromPointer(ptr).DestroyQWindowDefault()
	}
}

func (ptr *QWindow) ConnectDestroyQWindow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWindow"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWindow", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWindow", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWindow) DisconnectDestroyQWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWindow")
	}
}

func (ptr *QWindow) DestroyQWindow() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWindow_DestroyQWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWindow) DestroyQWindowDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWindow_DestroyQWindowDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWindow) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWindow___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWindow) __children_newList() unsafe.Pointer {
	return C.QWindow___children_newList(ptr.Pointer())
}

func (ptr *QWindow) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWindow___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWindow) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWindow___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWindow) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWindow___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWindow) __findChildren_newList() unsafe.Pointer {
	return C.QWindow___findChildren_newList(ptr.Pointer())
}

func (ptr *QWindow) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWindow___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWindow) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWindow) __findChildren_newList3() unsafe.Pointer {
	return C.QWindow___findChildren_newList3(ptr.Pointer())
}

//export callbackQWindow_ChildEvent
func callbackQWindow_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWindowFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWindow) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWindow) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWindow_ConnectNotify
func callbackQWindow_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWindowFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWindow) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWindow) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWindow_CustomEvent
func callbackQWindow_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWindowFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWindow) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWindow) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWindow_DeleteLater
func callbackQWindow_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWindowFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWindow) DeleteLater() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWindow_DeleteLater(ptr.Pointer())
	}
}

func (ptr *QWindow) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWindow_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQWindow_Destroyed
func callbackQWindow_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWindow_DisconnectNotify
func callbackQWindow_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWindowFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWindow) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWindow) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWindow_EventFilter
func callbackQWindow_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWindowFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWindow) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWindow_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QWindow) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWindow_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWindow_ObjectNameChanged
func callbackQWindow_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGui_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWindow_TimerEvent
func callbackQWindow_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWindowFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWindow) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWindow) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWindow_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWindow_SurfaceType
func callbackQWindow_SurfaceType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "surfaceType"); signal != nil {
		return C.longlong((*(*func() QSurface__SurfaceType)(signal))())
	}

	return C.longlong(NewQWindowFromPointer(ptr).SurfaceTypeDefault())
}

func (ptr *QWindow) SurfaceType() QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QWindow_SurfaceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWindow) SurfaceTypeDefault() QSurface__SurfaceType {
	if ptr.Pointer() != nil {
		return QSurface__SurfaceType(C.QWindow_SurfaceTypeDefault(ptr.Pointer()))
	}
	return 0
}

func init() {
}
