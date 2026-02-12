//go:build minimal
// +build minimal

package printsupport

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "printsupport-minimal.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtPrintSupport_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtPrintSupport_PackedString) []byte {
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

// QAbstractPrintDialog::PrintDialogOption
//
//go:generate stringer -type=QAbstractPrintDialog__PrintDialogOption
type QAbstractPrintDialog__PrintDialogOption int64

const (
	QAbstractPrintDialog__None               QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0000)
	QAbstractPrintDialog__PrintToFile        QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0001)
	QAbstractPrintDialog__PrintSelection     QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0002)
	QAbstractPrintDialog__PrintPageRange     QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0004)
	QAbstractPrintDialog__PrintShowPageSize  QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0008)
	QAbstractPrintDialog__PrintCollateCopies QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0010)
	QAbstractPrintDialog__DontUseSheet       QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0020)
	QAbstractPrintDialog__PrintCurrentPage   QAbstractPrintDialog__PrintDialogOption = QAbstractPrintDialog__PrintDialogOption(0x0040)
)

// QAbstractPrintDialog::PrintRange
//
//go:generate stringer -type=QAbstractPrintDialog__PrintRange
type QAbstractPrintDialog__PrintRange int64

const (
	QAbstractPrintDialog__AllPages    QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(0)
	QAbstractPrintDialog__Selection   QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(1)
	QAbstractPrintDialog__PageRange   QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(2)
	QAbstractPrintDialog__CurrentPage QAbstractPrintDialog__PrintRange = QAbstractPrintDialog__PrintRange(3)
)

// QPrintEngine::PrintEnginePropertyKey
//
//go:generate stringer -type=QPrintEngine__PrintEnginePropertyKey
type QPrintEngine__PrintEnginePropertyKey int64

const (
	QPrintEngine__PPK_CollateCopies          QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(0)
	QPrintEngine__PPK_ColorMode              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(1)
	QPrintEngine__PPK_Creator                QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(2)
	QPrintEngine__PPK_DocumentName           QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(3)
	QPrintEngine__PPK_FullPage               QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(4)
	QPrintEngine__PPK_NumberOfCopies         QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(5)
	QPrintEngine__PPK_Orientation            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(6)
	QPrintEngine__PPK_OutputFileName         QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(7)
	QPrintEngine__PPK_PageOrder              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(8)
	QPrintEngine__PPK_PageRect               QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(9)
	QPrintEngine__PPK_PageSize               QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(10)
	QPrintEngine__PPK_PaperRect              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(11)
	QPrintEngine__PPK_PaperSource            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(12)
	QPrintEngine__PPK_PrinterName            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(13)
	QPrintEngine__PPK_PrinterProgram         QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(14)
	QPrintEngine__PPK_Resolution             QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(15)
	QPrintEngine__PPK_SelectionOption        QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(16)
	QPrintEngine__PPK_SupportedResolutions   QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(17)
	QPrintEngine__PPK_WindowsPageSize        QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(18)
	QPrintEngine__PPK_FontEmbedding          QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(19)
	QPrintEngine__PPK_Duplex                 QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(20)
	QPrintEngine__PPK_PaperSources           QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(21)
	QPrintEngine__PPK_CustomPaperSize        QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(22)
	QPrintEngine__PPK_PageMargins            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(23)
	QPrintEngine__PPK_CopyCount              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(24)
	QPrintEngine__PPK_SupportsMultipleCopies QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(25)
	QPrintEngine__PPK_PaperName              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(26)
	QPrintEngine__PPK_QPageSize              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(27)
	QPrintEngine__PPK_QPageMargins           QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(28)
	QPrintEngine__PPK_QPageLayout            QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(29)
	QPrintEngine__PPK_PaperSize              QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(QPrintEngine__PPK_PageSize)
	QPrintEngine__PPK_CustomBase             QPrintEngine__PrintEnginePropertyKey = QPrintEngine__PrintEnginePropertyKey(0xff00)
)

// QPrintPreviewWidget::ViewMode
//
//go:generate stringer -type=QPrintPreviewWidget__ViewMode
type QPrintPreviewWidget__ViewMode int64

const (
	QPrintPreviewWidget__SinglePageView  QPrintPreviewWidget__ViewMode = QPrintPreviewWidget__ViewMode(0)
	QPrintPreviewWidget__FacingPagesView QPrintPreviewWidget__ViewMode = QPrintPreviewWidget__ViewMode(1)
	QPrintPreviewWidget__AllPagesView    QPrintPreviewWidget__ViewMode = QPrintPreviewWidget__ViewMode(2)
)

// QPrintPreviewWidget::ZoomMode
//
//go:generate stringer -type=QPrintPreviewWidget__ZoomMode
type QPrintPreviewWidget__ZoomMode int64

const (
	QPrintPreviewWidget__CustomZoom QPrintPreviewWidget__ZoomMode = QPrintPreviewWidget__ZoomMode(0)
	QPrintPreviewWidget__FitToWidth QPrintPreviewWidget__ZoomMode = QPrintPreviewWidget__ZoomMode(1)
	QPrintPreviewWidget__FitInView  QPrintPreviewWidget__ZoomMode = QPrintPreviewWidget__ZoomMode(2)
)

type QPrinter struct {
	gui.QPagedPaintDevice
}

type QPrinter_ITF interface {
	gui.QPagedPaintDevice_ITF
	QPrinter_PTR() *QPrinter
}

func (ptr *QPrinter) QPrinter_PTR() *QPrinter {
	return ptr
}

func (ptr *QPrinter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPagedPaintDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QPrinter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPagedPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQPrinter(ptr QPrinter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrinter_PTR().Pointer()
	}
	return nil
}

func NewQPrinterFromPointer(ptr unsafe.Pointer) (n *QPrinter) {
	n = new(QPrinter)
	n.SetPointer(ptr)
	return
}

// QPrinter::ColorMode
//
//go:generate stringer -type=QPrinter__ColorMode
type QPrinter__ColorMode int64

const (
	QPrinter__GrayScale QPrinter__ColorMode = QPrinter__ColorMode(0)
	QPrinter__Color     QPrinter__ColorMode = QPrinter__ColorMode(1)
)

// QPrinter::DuplexMode
//
//go:generate stringer -type=QPrinter__DuplexMode
type QPrinter__DuplexMode int64

const (
	QPrinter__DuplexNone      QPrinter__DuplexMode = QPrinter__DuplexMode(0)
	QPrinter__DuplexAuto      QPrinter__DuplexMode = QPrinter__DuplexMode(1)
	QPrinter__DuplexLongSide  QPrinter__DuplexMode = QPrinter__DuplexMode(2)
	QPrinter__DuplexShortSide QPrinter__DuplexMode = QPrinter__DuplexMode(3)
)

// QPrinter::Orientation
//
//go:generate stringer -type=QPrinter__Orientation
type QPrinter__Orientation int64

const (
	QPrinter__Portrait  QPrinter__Orientation = QPrinter__Orientation(0)
	QPrinter__Landscape QPrinter__Orientation = QPrinter__Orientation(1)
)

// QPrinter::OutputFormat
//
//go:generate stringer -type=QPrinter__OutputFormat
type QPrinter__OutputFormat int64

const (
	QPrinter__NativeFormat QPrinter__OutputFormat = QPrinter__OutputFormat(0)
	QPrinter__PdfFormat    QPrinter__OutputFormat = QPrinter__OutputFormat(1)
)

// QPrinter::PageOrder
//
//go:generate stringer -type=QPrinter__PageOrder
type QPrinter__PageOrder int64

const (
	QPrinter__FirstPageFirst QPrinter__PageOrder = QPrinter__PageOrder(0)
	QPrinter__LastPageFirst  QPrinter__PageOrder = QPrinter__PageOrder(1)
)

// QPrinter::PaperSource
//
//go:generate stringer -type=QPrinter__PaperSource
type QPrinter__PaperSource int64

const (
	QPrinter__OnlyOne         QPrinter__PaperSource = QPrinter__PaperSource(0)
	QPrinter__Lower           QPrinter__PaperSource = QPrinter__PaperSource(1)
	QPrinter__Middle          QPrinter__PaperSource = QPrinter__PaperSource(2)
	QPrinter__Manual          QPrinter__PaperSource = QPrinter__PaperSource(3)
	QPrinter__Envelope        QPrinter__PaperSource = QPrinter__PaperSource(4)
	QPrinter__EnvelopeManual  QPrinter__PaperSource = QPrinter__PaperSource(5)
	QPrinter__Auto            QPrinter__PaperSource = QPrinter__PaperSource(6)
	QPrinter__Tractor         QPrinter__PaperSource = QPrinter__PaperSource(7)
	QPrinter__SmallFormat     QPrinter__PaperSource = QPrinter__PaperSource(8)
	QPrinter__LargeFormat     QPrinter__PaperSource = QPrinter__PaperSource(9)
	QPrinter__LargeCapacity   QPrinter__PaperSource = QPrinter__PaperSource(10)
	QPrinter__Cassette        QPrinter__PaperSource = QPrinter__PaperSource(11)
	QPrinter__FormSource      QPrinter__PaperSource = QPrinter__PaperSource(12)
	QPrinter__MaxPageSource   QPrinter__PaperSource = QPrinter__PaperSource(13)
	QPrinter__CustomSource    QPrinter__PaperSource = QPrinter__PaperSource(14)
	QPrinter__LastPaperSource QPrinter__PaperSource = QPrinter__PaperSource(QPrinter__CustomSource)
	QPrinter__Upper           QPrinter__PaperSource = QPrinter__PaperSource(QPrinter__OnlyOne)
)

// QPrinter::PrintRange
//
//go:generate stringer -type=QPrinter__PrintRange
type QPrinter__PrintRange int64

const (
	QPrinter__AllPages    QPrinter__PrintRange = QPrinter__PrintRange(0)
	QPrinter__Selection   QPrinter__PrintRange = QPrinter__PrintRange(1)
	QPrinter__PageRange   QPrinter__PrintRange = QPrinter__PrintRange(2)
	QPrinter__CurrentPage QPrinter__PrintRange = QPrinter__PrintRange(3)
)

// QPrinter::PrinterMode
//
//go:generate stringer -type=QPrinter__PrinterMode
type QPrinter__PrinterMode int64

const (
	QPrinter__ScreenResolution  QPrinter__PrinterMode = QPrinter__PrinterMode(0)
	QPrinter__PrinterResolution QPrinter__PrinterMode = QPrinter__PrinterMode(1)
	QPrinter__HighResolution    QPrinter__PrinterMode = QPrinter__PrinterMode(2)
)

// QPrinter::PrinterState
//
//go:generate stringer -type=QPrinter__PrinterState
type QPrinter__PrinterState int64

const (
	QPrinter__Idle    QPrinter__PrinterState = QPrinter__PrinterState(0)
	QPrinter__Active  QPrinter__PrinterState = QPrinter__PrinterState(1)
	QPrinter__Aborted QPrinter__PrinterState = QPrinter__PrinterState(2)
	QPrinter__Error   QPrinter__PrinterState = QPrinter__PrinterState(3)
)

// QPrinter::Unit
//
//go:generate stringer -type=QPrinter__Unit
type QPrinter__Unit int64

const (
	QPrinter__Millimeter  QPrinter__Unit = QPrinter__Unit(0)
	QPrinter__Point       QPrinter__Unit = QPrinter__Unit(1)
	QPrinter__Inch        QPrinter__Unit = QPrinter__Unit(2)
	QPrinter__Pica        QPrinter__Unit = QPrinter__Unit(3)
	QPrinter__Didot       QPrinter__Unit = QPrinter__Unit(4)
	QPrinter__Cicero      QPrinter__Unit = QPrinter__Unit(5)
	QPrinter__DevicePixel QPrinter__Unit = QPrinter__Unit(6)
)

func NewQPrinter(mode QPrinter__PrinterMode) *QPrinter {
	return NewQPrinterFromPointer(C.QPrinter_NewQPrinter(C.longlong(mode)))
}

func NewQPrinter2(printer QPrinterInfo_ITF, mode QPrinter__PrinterMode) *QPrinter {
	return NewQPrinterFromPointer(C.QPrinter_NewQPrinter2(PointerFromQPrinterInfo(printer), C.longlong(mode)))
}

func (ptr *QPrinter) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinter) OutputFileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinter_OutputFileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinter) OutputFormat() QPrinter__OutputFormat {
	if ptr.Pointer() != nil {
		return QPrinter__OutputFormat(C.QPrinter_OutputFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPrinter) Resolution() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_Resolution(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPrinter) SetOutputFileName(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QPrinter_SetOutputFileName(ptr.Pointer(), C.struct_QtPrintSupport_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

func (ptr *QPrinter) SetOutputFormat(format QPrinter__OutputFormat) {
	if ptr.Pointer() != nil {
		C.QPrinter_SetOutputFormat(ptr.Pointer(), C.longlong(format))
	}
}

//export callbackQPrinter_DestroyQPrinter
func callbackQPrinter_DestroyQPrinter(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPrinter"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPrinterFromPointer(ptr).DestroyQPrinterDefault()
	}
}

func (ptr *QPrinter) ConnectDestroyQPrinter(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPrinter"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPrinter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPrinter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPrinter) DisconnectDestroyQPrinter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPrinter")
	}
}

func (ptr *QPrinter) DestroyQPrinter() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPrinter_DestroyQPrinter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPrinter) DestroyQPrinterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPrinter_DestroyQPrinterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPrinter) __supportedResolutions_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter___supportedResolutions_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QPrinter) __supportedResolutions_setList(i int) {
	if ptr.Pointer() != nil {
		C.QPrinter___supportedResolutions_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QPrinter) __supportedResolutions_newList() unsafe.Pointer {
	return C.QPrinter___supportedResolutions_newList(ptr.Pointer())
}

//export callbackQPrinter_NewPage
func callbackQPrinter_NewPage(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "newPage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQPrinterFromPointer(ptr).NewPageDefault())))
}

func (ptr *QPrinter) NewPage() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_NewPage(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPrinter) NewPageDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPrinter_NewPageDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQPrinter_Metric
func callbackQPrinter_Metric(ptr unsafe.Pointer, metric C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(gui.QPaintDevice__PaintDeviceMetric) int)(signal))(gui.QPaintDevice__PaintDeviceMetric(metric))))
	}

	return C.int(int32(NewQPrinterFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(metric))))
}

func (ptr *QPrinter) MetricDefault(metric gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrinter_MetricDefault(ptr.Pointer(), C.longlong(metric))))
	}
	return 0
}

//export callbackQPrinter_PaintEngine
func callbackQPrinter_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine((*(*func() *gui.QPaintEngine)(signal))())
	}

	return gui.PointerFromQPaintEngine(NewQPrinterFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QPrinter) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QPrinter_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

type QPrinterInfo struct {
	ptr unsafe.Pointer
}

type QPrinterInfo_ITF interface {
	QPrinterInfo_PTR() *QPrinterInfo
}

func (ptr *QPrinterInfo) QPrinterInfo_PTR() *QPrinterInfo {
	return ptr
}

func (ptr *QPrinterInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPrinterInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPrinterInfo(ptr QPrinterInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPrinterInfo_PTR().Pointer()
	}
	return nil
}

func NewQPrinterInfoFromPointer(ptr unsafe.Pointer) (n *QPrinterInfo) {
	n = new(QPrinterInfo)
	n.SetPointer(ptr)
	return
}
func NewQPrinterInfo() *QPrinterInfo {
	tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo_NewQPrinterInfo())
	qt.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func NewQPrinterInfo2(other QPrinterInfo_ITF) *QPrinterInfo {
	tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo_NewQPrinterInfo2(PointerFromQPrinterInfo(other)))
	qt.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func NewQPrinterInfo3(printer QPrinter_ITF) *QPrinterInfo {
	tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo_NewQPrinterInfo3(PointerFromQPrinter(printer)))
	qt.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
	return tmpValue
}

func (ptr *QPrinterInfo) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QPrinterInfo_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPrinterInfo) DestroyQPrinterInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPrinterInfo_DestroyQPrinterInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPrinterInfo) __availablePrinters_atList(i int) *QPrinterInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQPrinterInfoFromPointer(C.QPrinterInfo___availablePrinters_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPrinterInfo).DestroyQPrinterInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinterInfo) __availablePrinters_setList(i QPrinterInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QPrinterInfo___availablePrinters_setList(ptr.Pointer(), PointerFromQPrinterInfo(i))
	}
}

func (ptr *QPrinterInfo) __availablePrinters_newList() unsafe.Pointer {
	return C.QPrinterInfo___availablePrinters_newList(ptr.Pointer())
}

func (ptr *QPrinterInfo) __supportedColorModes_atList(i int) QPrinter__ColorMode {
	if ptr.Pointer() != nil {
		return QPrinter__ColorMode(C.QPrinterInfo___supportedColorModes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QPrinterInfo) __supportedColorModes_setList(i QPrinter__ColorMode) {
	if ptr.Pointer() != nil {
		C.QPrinterInfo___supportedColorModes_setList(ptr.Pointer(), C.longlong(i))
	}
}

func (ptr *QPrinterInfo) __supportedColorModes_newList() unsafe.Pointer {
	return C.QPrinterInfo___supportedColorModes_newList(ptr.Pointer())
}

func (ptr *QPrinterInfo) __supportedDuplexModes_atList(i int) QPrinter__DuplexMode {
	if ptr.Pointer() != nil {
		return QPrinter__DuplexMode(C.QPrinterInfo___supportedDuplexModes_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QPrinterInfo) __supportedDuplexModes_setList(i QPrinter__DuplexMode) {
	if ptr.Pointer() != nil {
		C.QPrinterInfo___supportedDuplexModes_setList(ptr.Pointer(), C.longlong(i))
	}
}

func (ptr *QPrinterInfo) __supportedDuplexModes_newList() unsafe.Pointer {
	return C.QPrinterInfo___supportedDuplexModes_newList(ptr.Pointer())
}

func (ptr *QPrinterInfo) __supportedPageSizes_atList(i int) *gui.QPageSize {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPageSizeFromPointer(C.QPrinterInfo___supportedPageSizes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*gui.QPageSize).DestroyQPageSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPrinterInfo) __supportedPageSizes_setList(i gui.QPageSize_ITF) {
	if ptr.Pointer() != nil {
		C.QPrinterInfo___supportedPageSizes_setList(ptr.Pointer(), gui.PointerFromQPageSize(i))
	}
}

func (ptr *QPrinterInfo) __supportedPageSizes_newList() unsafe.Pointer {
	return C.QPrinterInfo___supportedPageSizes_newList(ptr.Pointer())
}

func (ptr *QPrinterInfo) __supportedPaperSizes_newList() unsafe.Pointer {
	return C.QPrinterInfo___supportedPaperSizes_newList(ptr.Pointer())
}

func (ptr *QPrinterInfo) __supportedResolutions_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPrinterInfo___supportedResolutions_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QPrinterInfo) __supportedResolutions_setList(i int) {
	if ptr.Pointer() != nil {
		C.QPrinterInfo___supportedResolutions_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QPrinterInfo) __supportedResolutions_newList() unsafe.Pointer {
	return C.QPrinterInfo___supportedResolutions_newList(ptr.Pointer())
}

func init() {
}
