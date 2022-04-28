//go:build !minimal
// +build !minimal

package sensors

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtSensors_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtSensors_PackedString) []byte {
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

type QAccelerometer struct {
	QSensor
}

type QAccelerometer_ITF interface {
	QSensor_ITF
	QAccelerometer_PTR() *QAccelerometer
}

func (ptr *QAccelerometer) QAccelerometer_PTR() *QAccelerometer {
	return ptr
}

func (ptr *QAccelerometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QAccelerometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQAccelerometer(ptr QAccelerometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometer_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerFromPointer(ptr unsafe.Pointer) (n *QAccelerometer) {
	n = new(QAccelerometer)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAccelerometer__AccelerationMode
//QAccelerometer::AccelerationMode
type QAccelerometer__AccelerationMode int64

const (
	QAccelerometer__Combined QAccelerometer__AccelerationMode = QAccelerometer__AccelerationMode(0)
	QAccelerometer__Gravity  QAccelerometer__AccelerationMode = QAccelerometer__AccelerationMode(1)
	QAccelerometer__User     QAccelerometer__AccelerationMode = QAccelerometer__AccelerationMode(2)
)

func NewQAccelerometer(parent core.QObject_ITF) *QAccelerometer {
	tmpValue := NewQAccelerometerFromPointer(C.QAccelerometer_NewQAccelerometer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAccelerometer) AccelerationMode() QAccelerometer__AccelerationMode {
	if ptr.Pointer() != nil {
		return QAccelerometer__AccelerationMode(C.QAccelerometer_AccelerationMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQAccelerometer_AccelerationModeChanged
func callbackQAccelerometer_AccelerationModeChanged(ptr unsafe.Pointer, accelerationMode C.longlong) {
	if signal := qt.GetSignal(ptr, "accelerationModeChanged"); signal != nil {
		(*(*func(QAccelerometer__AccelerationMode))(signal))(QAccelerometer__AccelerationMode(accelerationMode))
	}

}

func (ptr *QAccelerometer) ConnectAccelerationModeChanged(f func(accelerationMode QAccelerometer__AccelerationMode)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "accelerationModeChanged") {
			C.QAccelerometer_ConnectAccelerationModeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "accelerationModeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "accelerationModeChanged"); signal != nil {
			f := func(accelerationMode QAccelerometer__AccelerationMode) {
				(*(*func(QAccelerometer__AccelerationMode))(signal))(accelerationMode)
				f(accelerationMode)
			}
			qt.ConnectSignal(ptr.Pointer(), "accelerationModeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "accelerationModeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAccelerometer) DisconnectAccelerationModeChanged() {
	if ptr.Pointer() != nil {
		C.QAccelerometer_DisconnectAccelerationModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "accelerationModeChanged")
	}
}

func (ptr *QAccelerometer) AccelerationModeChanged(accelerationMode QAccelerometer__AccelerationMode) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_AccelerationModeChanged(ptr.Pointer(), C.longlong(accelerationMode))
	}
}

func (ptr *QAccelerometer) Reading() *QAccelerometerReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQAccelerometerReadingFromPointer(C.QAccelerometer_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAccelerometer) SetAccelerationMode(accelerationMode QAccelerometer__AccelerationMode) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_SetAccelerationMode(ptr.Pointer(), C.longlong(accelerationMode))
	}
}

//export callbackQAccelerometer_DestroyQAccelerometer
func callbackQAccelerometer_DestroyQAccelerometer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAccelerometer"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAccelerometerFromPointer(ptr).DestroyQAccelerometerDefault()
	}
}

func (ptr *QAccelerometer) ConnectDestroyQAccelerometer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAccelerometer"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAccelerometer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAccelerometer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAccelerometer) DisconnectDestroyQAccelerometer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAccelerometer")
	}
}

func (ptr *QAccelerometer) DestroyQAccelerometer() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAccelerometer_DestroyQAccelerometer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAccelerometer) DestroyQAccelerometerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAccelerometer_DestroyQAccelerometerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QAccelerometerFilter struct {
	QSensorFilter
}

type QAccelerometerFilter_ITF interface {
	QSensorFilter_ITF
	QAccelerometerFilter_PTR() *QAccelerometerFilter
}

func (ptr *QAccelerometerFilter) QAccelerometerFilter_PTR() *QAccelerometerFilter {
	return ptr
}

func (ptr *QAccelerometerFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QAccelerometerFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQAccelerometerFilter(ptr QAccelerometerFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerFilter_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerFilterFromPointer(ptr unsafe.Pointer) (n *QAccelerometerFilter) {
	n = new(QAccelerometerFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QAccelerometerFilter) DestroyQAccelerometerFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQAccelerometerFilter_Filter
func callbackQAccelerometerFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QAccelerometerReading) bool)(signal))(NewQAccelerometerReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAccelerometerFilter) ConnectFilter(f func(reading *QAccelerometerReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QAccelerometerReading) bool {
				(*(*func(*QAccelerometerReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAccelerometerFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QAccelerometerFilter) Filter(reading QAccelerometerReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAccelerometerFilter_Filter(ptr.Pointer(), PointerFromQAccelerometerReading(reading))) != 0
	}
	return false
}

type QAccelerometerReading struct {
	QSensorReading
}

type QAccelerometerReading_ITF interface {
	QSensorReading_ITF
	QAccelerometerReading_PTR() *QAccelerometerReading
}

func (ptr *QAccelerometerReading) QAccelerometerReading_PTR() *QAccelerometerReading {
	return ptr
}

func (ptr *QAccelerometerReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QAccelerometerReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQAccelerometerReading(ptr QAccelerometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerReading_PTR().Pointer()
	}
	return nil
}

func NewQAccelerometerReadingFromPointer(ptr unsafe.Pointer) (n *QAccelerometerReading) {
	n = new(QAccelerometerReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QAccelerometerReading) SetX(x float64) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QAccelerometerReading) SetY(y float64) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QAccelerometerReading) SetZ(z float64) {
	if ptr.Pointer() != nil {
		C.QAccelerometerReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

func (ptr *QAccelerometerReading) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccelerometerReading) Z() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAccelerometerReading_Z(ptr.Pointer()))
	}
	return 0
}

type QAltimeter struct {
	QSensor
}

type QAltimeter_ITF interface {
	QSensor_ITF
	QAltimeter_PTR() *QAltimeter
}

func (ptr *QAltimeter) QAltimeter_PTR() *QAltimeter {
	return ptr
}

func (ptr *QAltimeter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QAltimeter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQAltimeter(ptr QAltimeter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeter_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterFromPointer(ptr unsafe.Pointer) (n *QAltimeter) {
	n = new(QAltimeter)
	n.SetPointer(ptr)
	return
}
func NewQAltimeter(parent core.QObject_ITF) *QAltimeter {
	tmpValue := NewQAltimeterFromPointer(C.QAltimeter_NewQAltimeter(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAltimeter) Reading() *QAltimeterReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQAltimeterReadingFromPointer(C.QAltimeter_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAltimeter_DestroyQAltimeter
func callbackQAltimeter_DestroyQAltimeter(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAltimeter"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAltimeterFromPointer(ptr).DestroyQAltimeterDefault()
	}
}

func (ptr *QAltimeter) ConnectDestroyQAltimeter(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAltimeter"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAltimeter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAltimeter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAltimeter) DisconnectDestroyQAltimeter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAltimeter")
	}
}

func (ptr *QAltimeter) DestroyQAltimeter() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAltimeter_DestroyQAltimeter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAltimeter) DestroyQAltimeterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAltimeter_DestroyQAltimeterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QAltimeterFilter struct {
	QSensorFilter
}

type QAltimeterFilter_ITF interface {
	QSensorFilter_ITF
	QAltimeterFilter_PTR() *QAltimeterFilter
}

func (ptr *QAltimeterFilter) QAltimeterFilter_PTR() *QAltimeterFilter {
	return ptr
}

func (ptr *QAltimeterFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QAltimeterFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQAltimeterFilter(ptr QAltimeterFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterFilter_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterFilterFromPointer(ptr unsafe.Pointer) (n *QAltimeterFilter) {
	n = new(QAltimeterFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QAltimeterFilter) DestroyQAltimeterFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQAltimeterFilter_Filter
func callbackQAltimeterFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QAltimeterReading) bool)(signal))(NewQAltimeterReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAltimeterFilter) ConnectFilter(f func(reading *QAltimeterReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QAltimeterReading) bool {
				(*(*func(*QAltimeterReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAltimeterFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QAltimeterFilter) Filter(reading QAltimeterReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAltimeterFilter_Filter(ptr.Pointer(), PointerFromQAltimeterReading(reading))) != 0
	}
	return false
}

type QAltimeterReading struct {
	QSensorReading
}

type QAltimeterReading_ITF interface {
	QSensorReading_ITF
	QAltimeterReading_PTR() *QAltimeterReading
}

func (ptr *QAltimeterReading) QAltimeterReading_PTR() *QAltimeterReading {
	return ptr
}

func (ptr *QAltimeterReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QAltimeterReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQAltimeterReading(ptr QAltimeterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAltimeterReading_PTR().Pointer()
	}
	return nil
}

func NewQAltimeterReadingFromPointer(ptr unsafe.Pointer) (n *QAltimeterReading) {
	n = new(QAltimeterReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QAltimeterReading) Altitude() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAltimeterReading_Altitude(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAltimeterReading) SetAltitude(altitude float64) {
	if ptr.Pointer() != nil {
		C.QAltimeterReading_SetAltitude(ptr.Pointer(), C.double(altitude))
	}
}

type QAmbientLightFilter struct {
	QSensorFilter
}

type QAmbientLightFilter_ITF interface {
	QSensorFilter_ITF
	QAmbientLightFilter_PTR() *QAmbientLightFilter
}

func (ptr *QAmbientLightFilter) QAmbientLightFilter_PTR() *QAmbientLightFilter {
	return ptr
}

func (ptr *QAmbientLightFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientLightFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientLightFilter(ptr QAmbientLightFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightFilter_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightFilterFromPointer(ptr unsafe.Pointer) (n *QAmbientLightFilter) {
	n = new(QAmbientLightFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QAmbientLightFilter) DestroyQAmbientLightFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQAmbientLightFilter_Filter
func callbackQAmbientLightFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QAmbientLightReading) bool)(signal))(NewQAmbientLightReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAmbientLightFilter) ConnectFilter(f func(reading *QAmbientLightReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QAmbientLightReading) bool {
				(*(*func(*QAmbientLightReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAmbientLightFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QAmbientLightFilter) Filter(reading QAmbientLightReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAmbientLightFilter_Filter(ptr.Pointer(), PointerFromQAmbientLightReading(reading))) != 0
	}
	return false
}

type QAmbientLightReading struct {
	QSensorReading
}

type QAmbientLightReading_ITF interface {
	QSensorReading_ITF
	QAmbientLightReading_PTR() *QAmbientLightReading
}

func (ptr *QAmbientLightReading) QAmbientLightReading_PTR() *QAmbientLightReading {
	return ptr
}

func (ptr *QAmbientLightReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientLightReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientLightReading(ptr QAmbientLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightReading_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightReadingFromPointer(ptr unsafe.Pointer) (n *QAmbientLightReading) {
	n = new(QAmbientLightReading)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAmbientLightReading__LightLevel
//QAmbientLightReading::LightLevel
type QAmbientLightReading__LightLevel int64

const (
	QAmbientLightReading__Undefined QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(0)
	QAmbientLightReading__Dark      QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(1)
	QAmbientLightReading__Twilight  QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(2)
	QAmbientLightReading__Light     QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(3)
	QAmbientLightReading__Bright    QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(4)
	QAmbientLightReading__Sunny     QAmbientLightReading__LightLevel = QAmbientLightReading__LightLevel(5)
)

func (ptr *QAmbientLightReading) LightLevel() QAmbientLightReading__LightLevel {
	if ptr.Pointer() != nil {
		return QAmbientLightReading__LightLevel(C.QAmbientLightReading_LightLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAmbientLightReading) SetLightLevel(lightLevel QAmbientLightReading__LightLevel) {
	if ptr.Pointer() != nil {
		C.QAmbientLightReading_SetLightLevel(ptr.Pointer(), C.longlong(lightLevel))
	}
}

type QAmbientLightSensor struct {
	QSensor
}

type QAmbientLightSensor_ITF interface {
	QSensor_ITF
	QAmbientLightSensor_PTR() *QAmbientLightSensor
}

func (ptr *QAmbientLightSensor) QAmbientLightSensor_PTR() *QAmbientLightSensor {
	return ptr
}

func (ptr *QAmbientLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientLightSensor(ptr QAmbientLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightSensor_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightSensorFromPointer(ptr unsafe.Pointer) (n *QAmbientLightSensor) {
	n = new(QAmbientLightSensor)
	n.SetPointer(ptr)
	return
}
func NewQAmbientLightSensor(parent core.QObject_ITF) *QAmbientLightSensor {
	tmpValue := NewQAmbientLightSensorFromPointer(C.QAmbientLightSensor_NewQAmbientLightSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAmbientLightSensor) Reading() *QAmbientLightReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQAmbientLightReadingFromPointer(C.QAmbientLightSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAmbientLightSensor_DestroyQAmbientLightSensor
func callbackQAmbientLightSensor_DestroyQAmbientLightSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAmbientLightSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAmbientLightSensorFromPointer(ptr).DestroyQAmbientLightSensorDefault()
	}
}

func (ptr *QAmbientLightSensor) ConnectDestroyQAmbientLightSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAmbientLightSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAmbientLightSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAmbientLightSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAmbientLightSensor) DisconnectDestroyQAmbientLightSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAmbientLightSensor")
	}
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAmbientLightSensor_DestroyQAmbientLightSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientLightSensor) DestroyQAmbientLightSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAmbientLightSensor_DestroyQAmbientLightSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QAmbientTemperatureFilter struct {
	QSensorFilter
}

type QAmbientTemperatureFilter_ITF interface {
	QSensorFilter_ITF
	QAmbientTemperatureFilter_PTR() *QAmbientTemperatureFilter
}

func (ptr *QAmbientTemperatureFilter) QAmbientTemperatureFilter_PTR() *QAmbientTemperatureFilter {
	return ptr
}

func (ptr *QAmbientTemperatureFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientTemperatureFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientTemperatureFilter(ptr QAmbientTemperatureFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureFilter_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureFilterFromPointer(ptr unsafe.Pointer) (n *QAmbientTemperatureFilter) {
	n = new(QAmbientTemperatureFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QAmbientTemperatureFilter) DestroyQAmbientTemperatureFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQAmbientTemperatureFilter_Filter
func callbackQAmbientTemperatureFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QAmbientTemperatureReading) bool)(signal))(NewQAmbientTemperatureReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAmbientTemperatureFilter) ConnectFilter(f func(reading *QAmbientTemperatureReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QAmbientTemperatureReading) bool {
				(*(*func(*QAmbientTemperatureReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAmbientTemperatureFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QAmbientTemperatureFilter) Filter(reading QAmbientTemperatureReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAmbientTemperatureFilter_Filter(ptr.Pointer(), PointerFromQAmbientTemperatureReading(reading))) != 0
	}
	return false
}

type QAmbientTemperatureReading struct {
	QSensorReading
}

type QAmbientTemperatureReading_ITF interface {
	QSensorReading_ITF
	QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading
}

func (ptr *QAmbientTemperatureReading) QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading {
	return ptr
}

func (ptr *QAmbientTemperatureReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientTemperatureReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientTemperatureReading(ptr QAmbientTemperatureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureReading_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureReadingFromPointer(ptr unsafe.Pointer) (n *QAmbientTemperatureReading) {
	n = new(QAmbientTemperatureReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QAmbientTemperatureReading) SetTemperature(temperature float64) {
	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_SetTemperature(ptr.Pointer(), C.double(temperature))
	}
}

func (ptr *QAmbientTemperatureReading) Temperature() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QAmbientTemperatureReading_Temperature(ptr.Pointer()))
	}
	return 0
}

type QAmbientTemperatureSensor struct {
	QSensor
}

type QAmbientTemperatureSensor_ITF interface {
	QSensor_ITF
	QAmbientTemperatureSensor_PTR() *QAmbientTemperatureSensor
}

func (ptr *QAmbientTemperatureSensor) QAmbientTemperatureSensor_PTR() *QAmbientTemperatureSensor {
	return ptr
}

func (ptr *QAmbientTemperatureSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QAmbientTemperatureSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQAmbientTemperatureSensor(ptr QAmbientTemperatureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureSensor_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureSensorFromPointer(ptr unsafe.Pointer) (n *QAmbientTemperatureSensor) {
	n = new(QAmbientTemperatureSensor)
	n.SetPointer(ptr)
	return
}
func NewQAmbientTemperatureSensor(parent core.QObject_ITF) *QAmbientTemperatureSensor {
	tmpValue := NewQAmbientTemperatureSensorFromPointer(C.QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAmbientTemperatureSensor) Reading() *QAmbientTemperatureReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQAmbientTemperatureReadingFromPointer(C.QAmbientTemperatureSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor
func callbackQAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAmbientTemperatureSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAmbientTemperatureSensorFromPointer(ptr).DestroyQAmbientTemperatureSensorDefault()
	}
}

func (ptr *QAmbientTemperatureSensor) ConnectDestroyQAmbientTemperatureSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAmbientTemperatureSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAmbientTemperatureSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAmbientTemperatureSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAmbientTemperatureSensor) DisconnectDestroyQAmbientTemperatureSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAmbientTemperatureSensor")
	}
}

func (ptr *QAmbientTemperatureSensor) DestroyQAmbientTemperatureSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAmbientTemperatureSensor) DestroyQAmbientTemperatureSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QCompass struct {
	QSensor
}

type QCompass_ITF interface {
	QSensor_ITF
	QCompass_PTR() *QCompass
}

func (ptr *QCompass) QCompass_PTR() *QCompass {
	return ptr
}

func (ptr *QCompass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QCompass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQCompass(ptr QCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompass_PTR().Pointer()
	}
	return nil
}

func NewQCompassFromPointer(ptr unsafe.Pointer) (n *QCompass) {
	n = new(QCompass)
	n.SetPointer(ptr)
	return
}
func NewQCompass(parent core.QObject_ITF) *QCompass {
	tmpValue := NewQCompassFromPointer(C.QCompass_NewQCompass(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QCompass) Reading() *QCompassReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQCompassReadingFromPointer(C.QCompass_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQCompass_DestroyQCompass
func callbackQCompass_DestroyQCompass(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QCompass"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQCompassFromPointer(ptr).DestroyQCompassDefault()
	}
}

func (ptr *QCompass) ConnectDestroyQCompass(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QCompass"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QCompass", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QCompass", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCompass) DisconnectDestroyQCompass() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QCompass")
	}
}

func (ptr *QCompass) DestroyQCompass() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCompass_DestroyQCompass(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCompass) DestroyQCompassDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QCompass_DestroyQCompassDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QCompassFilter struct {
	QSensorFilter
}

type QCompassFilter_ITF interface {
	QSensorFilter_ITF
	QCompassFilter_PTR() *QCompassFilter
}

func (ptr *QCompassFilter) QCompassFilter_PTR() *QCompassFilter {
	return ptr
}

func (ptr *QCompassFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QCompassFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQCompassFilter(ptr QCompassFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassFilter_PTR().Pointer()
	}
	return nil
}

func NewQCompassFilterFromPointer(ptr unsafe.Pointer) (n *QCompassFilter) {
	n = new(QCompassFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QCompassFilter) DestroyQCompassFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQCompassFilter_Filter
func callbackQCompassFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QCompassReading) bool)(signal))(NewQCompassReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QCompassFilter) ConnectFilter(f func(reading *QCompassReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QCompassReading) bool {
				(*(*func(*QCompassReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QCompassFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QCompassFilter) Filter(reading QCompassReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QCompassFilter_Filter(ptr.Pointer(), PointerFromQCompassReading(reading))) != 0
	}
	return false
}

type QCompassReading struct {
	QSensorReading
}

type QCompassReading_ITF interface {
	QSensorReading_ITF
	QCompassReading_PTR() *QCompassReading
}

func (ptr *QCompassReading) QCompassReading_PTR() *QCompassReading {
	return ptr
}

func (ptr *QCompassReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QCompassReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQCompassReading(ptr QCompassReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassReading_PTR().Pointer()
	}
	return nil
}

func NewQCompassReadingFromPointer(ptr unsafe.Pointer) (n *QCompassReading) {
	n = new(QCompassReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QCompassReading) Azimuth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCompassReading_Azimuth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompassReading) CalibrationLevel() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QCompassReading_CalibrationLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompassReading) SetAzimuth(azimuth float64) {
	if ptr.Pointer() != nil {
		C.QCompassReading_SetAzimuth(ptr.Pointer(), C.double(azimuth))
	}
}

func (ptr *QCompassReading) SetCalibrationLevel(calibrationLevel float64) {
	if ptr.Pointer() != nil {
		C.QCompassReading_SetCalibrationLevel(ptr.Pointer(), C.double(calibrationLevel))
	}
}

type QDistanceFilter struct {
	QSensorFilter
}

type QDistanceFilter_ITF interface {
	QSensorFilter_ITF
	QDistanceFilter_PTR() *QDistanceFilter
}

func (ptr *QDistanceFilter) QDistanceFilter_PTR() *QDistanceFilter {
	return ptr
}

func (ptr *QDistanceFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QDistanceFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQDistanceFilter(ptr QDistanceFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceFilter_PTR().Pointer()
	}
	return nil
}

func NewQDistanceFilterFromPointer(ptr unsafe.Pointer) (n *QDistanceFilter) {
	n = new(QDistanceFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QDistanceFilter) DestroyQDistanceFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDistanceFilter_Filter
func callbackQDistanceFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QDistanceReading) bool)(signal))(NewQDistanceReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDistanceFilter) ConnectFilter(f func(reading *QDistanceReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QDistanceReading) bool {
				(*(*func(*QDistanceReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDistanceFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QDistanceFilter) Filter(reading QDistanceReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDistanceFilter_Filter(ptr.Pointer(), PointerFromQDistanceReading(reading))) != 0
	}
	return false
}

type QDistanceReading struct {
	QSensorReading
}

type QDistanceReading_ITF interface {
	QSensorReading_ITF
	QDistanceReading_PTR() *QDistanceReading
}

func (ptr *QDistanceReading) QDistanceReading_PTR() *QDistanceReading {
	return ptr
}

func (ptr *QDistanceReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QDistanceReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQDistanceReading(ptr QDistanceReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceReading_PTR().Pointer()
	}
	return nil
}

func NewQDistanceReadingFromPointer(ptr unsafe.Pointer) (n *QDistanceReading) {
	n = new(QDistanceReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QDistanceReading) Distance() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QDistanceReading_Distance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDistanceReading) SetDistance(distance float64) {
	if ptr.Pointer() != nil {
		C.QDistanceReading_SetDistance(ptr.Pointer(), C.double(distance))
	}
}

type QDistanceSensor struct {
	QSensor
}

type QDistanceSensor_ITF interface {
	QSensor_ITF
	QDistanceSensor_PTR() *QDistanceSensor
}

func (ptr *QDistanceSensor) QDistanceSensor_PTR() *QDistanceSensor {
	return ptr
}

func (ptr *QDistanceSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QDistanceSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQDistanceSensor(ptr QDistanceSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDistanceSensor_PTR().Pointer()
	}
	return nil
}

func NewQDistanceSensorFromPointer(ptr unsafe.Pointer) (n *QDistanceSensor) {
	n = new(QDistanceSensor)
	n.SetPointer(ptr)
	return
}
func NewQDistanceSensor(parent core.QObject_ITF) *QDistanceSensor {
	tmpValue := NewQDistanceSensorFromPointer(C.QDistanceSensor_NewQDistanceSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDistanceSensor) Reading() *QDistanceReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQDistanceReadingFromPointer(C.QDistanceSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDistanceSensor_DestroyQDistanceSensor
func callbackQDistanceSensor_DestroyQDistanceSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDistanceSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQDistanceSensorFromPointer(ptr).DestroyQDistanceSensorDefault()
	}
}

func (ptr *QDistanceSensor) ConnectDestroyQDistanceSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDistanceSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QDistanceSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDistanceSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QDistanceSensor) DisconnectDestroyQDistanceSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDistanceSensor")
	}
}

func (ptr *QDistanceSensor) DestroyQDistanceSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDistanceSensor_DestroyQDistanceSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDistanceSensor) DestroyQDistanceSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QDistanceSensor_DestroyQDistanceSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGyroscope struct {
	QSensor
}

type QGyroscope_ITF interface {
	QSensor_ITF
	QGyroscope_PTR() *QGyroscope
}

func (ptr *QGyroscope) QGyroscope_PTR() *QGyroscope {
	return ptr
}

func (ptr *QGyroscope) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QGyroscope) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQGyroscope(ptr QGyroscope_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscope_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeFromPointer(ptr unsafe.Pointer) (n *QGyroscope) {
	n = new(QGyroscope)
	n.SetPointer(ptr)
	return
}
func NewQGyroscope(parent core.QObject_ITF) *QGyroscope {
	tmpValue := NewQGyroscopeFromPointer(C.QGyroscope_NewQGyroscope(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGyroscope) Reading() *QGyroscopeReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQGyroscopeReadingFromPointer(C.QGyroscope_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGyroscope_DestroyQGyroscope
func callbackQGyroscope_DestroyQGyroscope(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGyroscope"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGyroscopeFromPointer(ptr).DestroyQGyroscopeDefault()
	}
}

func (ptr *QGyroscope) ConnectDestroyQGyroscope(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGyroscope"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGyroscope", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGyroscope", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGyroscope) DisconnectDestroyQGyroscope() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGyroscope")
	}
}

func (ptr *QGyroscope) DestroyQGyroscope() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGyroscope_DestroyQGyroscope(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGyroscope) DestroyQGyroscopeDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGyroscope_DestroyQGyroscopeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QGyroscopeFilter struct {
	QSensorFilter
}

type QGyroscopeFilter_ITF interface {
	QSensorFilter_ITF
	QGyroscopeFilter_PTR() *QGyroscopeFilter
}

func (ptr *QGyroscopeFilter) QGyroscopeFilter_PTR() *QGyroscopeFilter {
	return ptr
}

func (ptr *QGyroscopeFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QGyroscopeFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQGyroscopeFilter(ptr QGyroscopeFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeFilter_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeFilterFromPointer(ptr unsafe.Pointer) (n *QGyroscopeFilter) {
	n = new(QGyroscopeFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QGyroscopeFilter) DestroyQGyroscopeFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGyroscopeFilter_Filter
func callbackQGyroscopeFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QGyroscopeReading) bool)(signal))(NewQGyroscopeReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QGyroscopeFilter) ConnectFilter(f func(reading *QGyroscopeReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QGyroscopeReading) bool {
				(*(*func(*QGyroscopeReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGyroscopeFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QGyroscopeFilter) Filter(reading QGyroscopeReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QGyroscopeFilter_Filter(ptr.Pointer(), PointerFromQGyroscopeReading(reading))) != 0
	}
	return false
}

type QGyroscopeReading struct {
	QSensorReading
}

type QGyroscopeReading_ITF interface {
	QSensorReading_ITF
	QGyroscopeReading_PTR() *QGyroscopeReading
}

func (ptr *QGyroscopeReading) QGyroscopeReading_PTR() *QGyroscopeReading {
	return ptr
}

func (ptr *QGyroscopeReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QGyroscopeReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQGyroscopeReading(ptr QGyroscopeReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGyroscopeReading_PTR().Pointer()
	}
	return nil
}

func NewQGyroscopeReadingFromPointer(ptr unsafe.Pointer) (n *QGyroscopeReading) {
	n = new(QGyroscopeReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QGyroscopeReading) SetX(x float64) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QGyroscopeReading) SetY(y float64) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QGyroscopeReading) SetZ(z float64) {
	if ptr.Pointer() != nil {
		C.QGyroscopeReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

func (ptr *QGyroscopeReading) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGyroscopeReading) Z() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGyroscopeReading_Z(ptr.Pointer()))
	}
	return 0
}

type QHolsterFilter struct {
	QSensorFilter
}

type QHolsterFilter_ITF interface {
	QSensorFilter_ITF
	QHolsterFilter_PTR() *QHolsterFilter
}

func (ptr *QHolsterFilter) QHolsterFilter_PTR() *QHolsterFilter {
	return ptr
}

func (ptr *QHolsterFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QHolsterFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQHolsterFilter(ptr QHolsterFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterFilter_PTR().Pointer()
	}
	return nil
}

func NewQHolsterFilterFromPointer(ptr unsafe.Pointer) (n *QHolsterFilter) {
	n = new(QHolsterFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QHolsterFilter) DestroyQHolsterFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHolsterFilter_Filter
func callbackQHolsterFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QHolsterReading) bool)(signal))(NewQHolsterReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QHolsterFilter) ConnectFilter(f func(reading *QHolsterReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QHolsterReading) bool {
				(*(*func(*QHolsterReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHolsterFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QHolsterFilter) Filter(reading QHolsterReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHolsterFilter_Filter(ptr.Pointer(), PointerFromQHolsterReading(reading))) != 0
	}
	return false
}

type QHolsterReading struct {
	QSensorReading
}

type QHolsterReading_ITF interface {
	QSensorReading_ITF
	QHolsterReading_PTR() *QHolsterReading
}

func (ptr *QHolsterReading) QHolsterReading_PTR() *QHolsterReading {
	return ptr
}

func (ptr *QHolsterReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QHolsterReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQHolsterReading(ptr QHolsterReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterReading_PTR().Pointer()
	}
	return nil
}

func NewQHolsterReadingFromPointer(ptr unsafe.Pointer) (n *QHolsterReading) {
	n = new(QHolsterReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QHolsterReading) Holstered() bool {
	if ptr.Pointer() != nil {
		return int8(C.QHolsterReading_Holstered(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHolsterReading) SetHolstered(holstered bool) {
	if ptr.Pointer() != nil {
		C.QHolsterReading_SetHolstered(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(holstered))))
	}
}

type QHolsterSensor struct {
	QSensor
}

type QHolsterSensor_ITF interface {
	QSensor_ITF
	QHolsterSensor_PTR() *QHolsterSensor
}

func (ptr *QHolsterSensor) QHolsterSensor_PTR() *QHolsterSensor {
	return ptr
}

func (ptr *QHolsterSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QHolsterSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQHolsterSensor(ptr QHolsterSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHolsterSensor_PTR().Pointer()
	}
	return nil
}

func NewQHolsterSensorFromPointer(ptr unsafe.Pointer) (n *QHolsterSensor) {
	n = new(QHolsterSensor)
	n.SetPointer(ptr)
	return
}
func NewQHolsterSensor(parent core.QObject_ITF) *QHolsterSensor {
	tmpValue := NewQHolsterSensorFromPointer(C.QHolsterSensor_NewQHolsterSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHolsterSensor) Reading() *QHolsterReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQHolsterReadingFromPointer(C.QHolsterSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHolsterSensor_DestroyQHolsterSensor
func callbackQHolsterSensor_DestroyQHolsterSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHolsterSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHolsterSensorFromPointer(ptr).DestroyQHolsterSensorDefault()
	}
}

func (ptr *QHolsterSensor) ConnectDestroyQHolsterSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHolsterSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHolsterSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHolsterSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHolsterSensor) DisconnectDestroyQHolsterSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHolsterSensor")
	}
}

func (ptr *QHolsterSensor) DestroyQHolsterSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHolsterSensor_DestroyQHolsterSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHolsterSensor) DestroyQHolsterSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHolsterSensor_DestroyQHolsterSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QHumidityFilter struct {
	QSensorFilter
}

type QHumidityFilter_ITF interface {
	QSensorFilter_ITF
	QHumidityFilter_PTR() *QHumidityFilter
}

func (ptr *QHumidityFilter) QHumidityFilter_PTR() *QHumidityFilter {
	return ptr
}

func (ptr *QHumidityFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QHumidityFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQHumidityFilter(ptr QHumidityFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHumidityFilter_PTR().Pointer()
	}
	return nil
}

func NewQHumidityFilterFromPointer(ptr unsafe.Pointer) (n *QHumidityFilter) {
	n = new(QHumidityFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QHumidityFilter) DestroyQHumidityFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQHumidityFilter_Filter
func callbackQHumidityFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QHumidityReading) bool)(signal))(NewQHumidityReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QHumidityFilter) ConnectFilter(f func(reading *QHumidityReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QHumidityReading) bool {
				(*(*func(*QHumidityReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHumidityFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QHumidityFilter) Filter(reading QHumidityReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QHumidityFilter_Filter(ptr.Pointer(), PointerFromQHumidityReading(reading))) != 0
	}
	return false
}

type QHumidityReading struct {
	QSensorReading
}

type QHumidityReading_ITF interface {
	QSensorReading_ITF
	QHumidityReading_PTR() *QHumidityReading
}

func (ptr *QHumidityReading) QHumidityReading_PTR() *QHumidityReading {
	return ptr
}

func (ptr *QHumidityReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QHumidityReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQHumidityReading(ptr QHumidityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHumidityReading_PTR().Pointer()
	}
	return nil
}

func NewQHumidityReadingFromPointer(ptr unsafe.Pointer) (n *QHumidityReading) {
	n = new(QHumidityReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QHumidityReading) AbsoluteHumidity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QHumidityReading_AbsoluteHumidity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHumidityReading) RelativeHumidity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QHumidityReading_RelativeHumidity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHumidityReading) SetAbsoluteHumidity(value float64) {
	if ptr.Pointer() != nil {
		C.QHumidityReading_SetAbsoluteHumidity(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QHumidityReading) SetRelativeHumidity(humidity float64) {
	if ptr.Pointer() != nil {
		C.QHumidityReading_SetRelativeHumidity(ptr.Pointer(), C.double(humidity))
	}
}

type QHumiditySensor struct {
	QSensor
}

type QHumiditySensor_ITF interface {
	QSensor_ITF
	QHumiditySensor_PTR() *QHumiditySensor
}

func (ptr *QHumiditySensor) QHumiditySensor_PTR() *QHumiditySensor {
	return ptr
}

func (ptr *QHumiditySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QHumiditySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQHumiditySensor(ptr QHumiditySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHumiditySensor_PTR().Pointer()
	}
	return nil
}

func NewQHumiditySensorFromPointer(ptr unsafe.Pointer) (n *QHumiditySensor) {
	n = new(QHumiditySensor)
	n.SetPointer(ptr)
	return
}
func NewQHumiditySensor(parent core.QObject_ITF) *QHumiditySensor {
	tmpValue := NewQHumiditySensorFromPointer(C.QHumiditySensor_NewQHumiditySensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHumiditySensor) Reading() *QHumidityReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQHumidityReadingFromPointer(C.QHumiditySensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQHumiditySensor_DestroyQHumiditySensor
func callbackQHumiditySensor_DestroyQHumiditySensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QHumiditySensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQHumiditySensorFromPointer(ptr).DestroyQHumiditySensorDefault()
	}
}

func (ptr *QHumiditySensor) ConnectDestroyQHumiditySensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QHumiditySensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QHumiditySensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QHumiditySensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QHumiditySensor) DisconnectDestroyQHumiditySensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QHumiditySensor")
	}
}

func (ptr *QHumiditySensor) DestroyQHumiditySensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHumiditySensor_DestroyQHumiditySensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHumiditySensor) DestroyQHumiditySensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QHumiditySensor_DestroyQHumiditySensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QIRProximityFilter struct {
	QSensorFilter
}

type QIRProximityFilter_ITF interface {
	QSensorFilter_ITF
	QIRProximityFilter_PTR() *QIRProximityFilter
}

func (ptr *QIRProximityFilter) QIRProximityFilter_PTR() *QIRProximityFilter {
	return ptr
}

func (ptr *QIRProximityFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QIRProximityFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQIRProximityFilter(ptr QIRProximityFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityFilter_PTR().Pointer()
	}
	return nil
}

func NewQIRProximityFilterFromPointer(ptr unsafe.Pointer) (n *QIRProximityFilter) {
	n = new(QIRProximityFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QIRProximityFilter) DestroyQIRProximityFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQIRProximityFilter_Filter
func callbackQIRProximityFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QIRProximityReading) bool)(signal))(NewQIRProximityReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QIRProximityFilter) ConnectFilter(f func(reading *QIRProximityReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QIRProximityReading) bool {
				(*(*func(*QIRProximityReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIRProximityFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QIRProximityFilter) Filter(reading QIRProximityReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIRProximityFilter_Filter(ptr.Pointer(), PointerFromQIRProximityReading(reading))) != 0
	}
	return false
}

type QIRProximityReading struct {
	QSensorReading
}

type QIRProximityReading_ITF interface {
	QSensorReading_ITF
	QIRProximityReading_PTR() *QIRProximityReading
}

func (ptr *QIRProximityReading) QIRProximityReading_PTR() *QIRProximityReading {
	return ptr
}

func (ptr *QIRProximityReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QIRProximityReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQIRProximityReading(ptr QIRProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximityReading_PTR().Pointer()
	}
	return nil
}

func NewQIRProximityReadingFromPointer(ptr unsafe.Pointer) (n *QIRProximityReading) {
	n = new(QIRProximityReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QIRProximityReading) Reflectance() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QIRProximityReading_Reflectance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QIRProximityReading) SetReflectance(reflectance float64) {
	if ptr.Pointer() != nil {
		C.QIRProximityReading_SetReflectance(ptr.Pointer(), C.double(reflectance))
	}
}

type QIRProximitySensor struct {
	QSensor
}

type QIRProximitySensor_ITF interface {
	QSensor_ITF
	QIRProximitySensor_PTR() *QIRProximitySensor
}

func (ptr *QIRProximitySensor) QIRProximitySensor_PTR() *QIRProximitySensor {
	return ptr
}

func (ptr *QIRProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QIRProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQIRProximitySensor(ptr QIRProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIRProximitySensor_PTR().Pointer()
	}
	return nil
}

func NewQIRProximitySensorFromPointer(ptr unsafe.Pointer) (n *QIRProximitySensor) {
	n = new(QIRProximitySensor)
	n.SetPointer(ptr)
	return
}
func NewQIRProximitySensor(parent core.QObject_ITF) *QIRProximitySensor {
	tmpValue := NewQIRProximitySensorFromPointer(C.QIRProximitySensor_NewQIRProximitySensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QIRProximitySensor) Reading() *QIRProximityReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQIRProximityReadingFromPointer(C.QIRProximitySensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQIRProximitySensor_DestroyQIRProximitySensor
func callbackQIRProximitySensor_DestroyQIRProximitySensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QIRProximitySensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQIRProximitySensorFromPointer(ptr).DestroyQIRProximitySensorDefault()
	}
}

func (ptr *QIRProximitySensor) ConnectDestroyQIRProximitySensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QIRProximitySensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QIRProximitySensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QIRProximitySensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIRProximitySensor) DisconnectDestroyQIRProximitySensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QIRProximitySensor")
	}
}

func (ptr *QIRProximitySensor) DestroyQIRProximitySensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIRProximitySensor_DestroyQIRProximitySensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIRProximitySensor) DestroyQIRProximitySensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIRProximitySensor_DestroyQIRProximitySensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLidFilter struct {
	QSensorFilter
}

type QLidFilter_ITF interface {
	QSensorFilter_ITF
	QLidFilter_PTR() *QLidFilter
}

func (ptr *QLidFilter) QLidFilter_PTR() *QLidFilter {
	return ptr
}

func (ptr *QLidFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QLidFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQLidFilter(ptr QLidFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLidFilter_PTR().Pointer()
	}
	return nil
}

func NewQLidFilterFromPointer(ptr unsafe.Pointer) (n *QLidFilter) {
	n = new(QLidFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QLidFilter) DestroyQLidFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQLidFilter_Filter
func callbackQLidFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QLidReading) bool)(signal))(NewQLidReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QLidFilter) ConnectFilter(f func(reading *QLidReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QLidReading) bool {
				(*(*func(*QLidReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLidFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QLidFilter) Filter(reading QLidReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLidFilter_Filter(ptr.Pointer(), PointerFromQLidReading(reading))) != 0
	}
	return false
}

type QLidReading struct {
	QSensorReading
}

type QLidReading_ITF interface {
	QSensorReading_ITF
	QLidReading_PTR() *QLidReading
}

func (ptr *QLidReading) QLidReading_PTR() *QLidReading {
	return ptr
}

func (ptr *QLidReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QLidReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQLidReading(ptr QLidReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLidReading_PTR().Pointer()
	}
	return nil
}

func NewQLidReadingFromPointer(ptr unsafe.Pointer) (n *QLidReading) {
	n = new(QLidReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QLidReading) BackLidClosed() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLidReading_BackLidClosed(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLidReading) FrontLidClosed() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLidReading_FrontLidClosed(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLidReading) SetBackLidClosed(closed bool) {
	if ptr.Pointer() != nil {
		C.QLidReading_SetBackLidClosed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(closed))))
	}
}

func (ptr *QLidReading) SetFrontLidClosed(closed bool) {
	if ptr.Pointer() != nil {
		C.QLidReading_SetFrontLidClosed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(closed))))
	}
}

type QLidSensor struct {
	QSensor
}

type QLidSensor_ITF interface {
	QSensor_ITF
	QLidSensor_PTR() *QLidSensor
}

func (ptr *QLidSensor) QLidSensor_PTR() *QLidSensor {
	return ptr
}

func (ptr *QLidSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QLidSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQLidSensor(ptr QLidSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLidSensor_PTR().Pointer()
	}
	return nil
}

func NewQLidSensorFromPointer(ptr unsafe.Pointer) (n *QLidSensor) {
	n = new(QLidSensor)
	n.SetPointer(ptr)
	return
}
func NewQLidSensor(parent core.QObject_ITF) *QLidSensor {
	tmpValue := NewQLidSensorFromPointer(C.QLidSensor_NewQLidSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLidSensor) Reading() *QLidReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQLidReadingFromPointer(C.QLidSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQLidSensor_DestroyQLidSensor
func callbackQLidSensor_DestroyQLidSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLidSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLidSensorFromPointer(ptr).DestroyQLidSensorDefault()
	}
}

func (ptr *QLidSensor) ConnectDestroyQLidSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLidSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLidSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLidSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLidSensor) DisconnectDestroyQLidSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLidSensor")
	}
}

func (ptr *QLidSensor) DestroyQLidSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLidSensor_DestroyQLidSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLidSensor) DestroyQLidSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLidSensor_DestroyQLidSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLightFilter struct {
	QSensorFilter
}

type QLightFilter_ITF interface {
	QSensorFilter_ITF
	QLightFilter_PTR() *QLightFilter
}

func (ptr *QLightFilter) QLightFilter_PTR() *QLightFilter {
	return ptr
}

func (ptr *QLightFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QLightFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQLightFilter(ptr QLightFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightFilter_PTR().Pointer()
	}
	return nil
}

func NewQLightFilterFromPointer(ptr unsafe.Pointer) (n *QLightFilter) {
	n = new(QLightFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QLightFilter) DestroyQLightFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQLightFilter_Filter
func callbackQLightFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QLightReading) bool)(signal))(NewQLightReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QLightFilter) ConnectFilter(f func(reading *QLightReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QLightReading) bool {
				(*(*func(*QLightReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLightFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QLightFilter) Filter(reading QLightReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLightFilter_Filter(ptr.Pointer(), PointerFromQLightReading(reading))) != 0
	}
	return false
}

type QLightReading struct {
	QSensorReading
}

type QLightReading_ITF interface {
	QSensorReading_ITF
	QLightReading_PTR() *QLightReading
}

func (ptr *QLightReading) QLightReading_PTR() *QLightReading {
	return ptr
}

func (ptr *QLightReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QLightReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQLightReading(ptr QLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightReading_PTR().Pointer()
	}
	return nil
}

func NewQLightReadingFromPointer(ptr unsafe.Pointer) (n *QLightReading) {
	n = new(QLightReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QLightReading) Lux() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLightReading_Lux(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLightReading) SetLux(lux float64) {
	if ptr.Pointer() != nil {
		C.QLightReading_SetLux(ptr.Pointer(), C.double(lux))
	}
}

type QLightSensor struct {
	QSensor
}

type QLightSensor_ITF interface {
	QSensor_ITF
	QLightSensor_PTR() *QLightSensor
}

func (ptr *QLightSensor) QLightSensor_PTR() *QLightSensor {
	return ptr
}

func (ptr *QLightSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QLightSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQLightSensor(ptr QLightSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLightSensor_PTR().Pointer()
	}
	return nil
}

func NewQLightSensorFromPointer(ptr unsafe.Pointer) (n *QLightSensor) {
	n = new(QLightSensor)
	n.SetPointer(ptr)
	return
}
func NewQLightSensor(parent core.QObject_ITF) *QLightSensor {
	tmpValue := NewQLightSensorFromPointer(C.QLightSensor_NewQLightSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLightSensor) FieldOfView() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLightSensor_FieldOfView(ptr.Pointer()))
	}
	return 0
}

//export callbackQLightSensor_FieldOfViewChanged
func callbackQLightSensor_FieldOfViewChanged(ptr unsafe.Pointer, fieldOfView C.double) {
	if signal := qt.GetSignal(ptr, "fieldOfViewChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(fieldOfView))
	}

}

func (ptr *QLightSensor) ConnectFieldOfViewChanged(f func(fieldOfView float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fieldOfViewChanged") {
			C.QLightSensor_ConnectFieldOfViewChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "fieldOfViewChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fieldOfViewChanged"); signal != nil {
			f := func(fieldOfView float64) {
				(*(*func(float64))(signal))(fieldOfView)
				f(fieldOfView)
			}
			qt.ConnectSignal(ptr.Pointer(), "fieldOfViewChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fieldOfViewChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLightSensor) DisconnectFieldOfViewChanged() {
	if ptr.Pointer() != nil {
		C.QLightSensor_DisconnectFieldOfViewChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fieldOfViewChanged")
	}
}

func (ptr *QLightSensor) FieldOfViewChanged(fieldOfView float64) {
	if ptr.Pointer() != nil {
		C.QLightSensor_FieldOfViewChanged(ptr.Pointer(), C.double(fieldOfView))
	}
}

func (ptr *QLightSensor) Reading() *QLightReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQLightReadingFromPointer(C.QLightSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLightSensor) SetFieldOfView(fieldOfView float64) {
	if ptr.Pointer() != nil {
		C.QLightSensor_SetFieldOfView(ptr.Pointer(), C.double(fieldOfView))
	}
}

//export callbackQLightSensor_DestroyQLightSensor
func callbackQLightSensor_DestroyQLightSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLightSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLightSensorFromPointer(ptr).DestroyQLightSensorDefault()
	}
}

func (ptr *QLightSensor) ConnectDestroyQLightSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLightSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLightSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLightSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLightSensor) DisconnectDestroyQLightSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLightSensor")
	}
}

func (ptr *QLightSensor) DestroyQLightSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLightSensor_DestroyQLightSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLightSensor) DestroyQLightSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLightSensor_DestroyQLightSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QMagnetometer struct {
	QSensor
}

type QMagnetometer_ITF interface {
	QSensor_ITF
	QMagnetometer_PTR() *QMagnetometer
}

func (ptr *QMagnetometer) QMagnetometer_PTR() *QMagnetometer {
	return ptr
}

func (ptr *QMagnetometer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QMagnetometer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQMagnetometer(ptr QMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometer_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerFromPointer(ptr unsafe.Pointer) (n *QMagnetometer) {
	n = new(QMagnetometer)
	n.SetPointer(ptr)
	return
}
func NewQMagnetometer(parent core.QObject_ITF) *QMagnetometer {
	tmpValue := NewQMagnetometerFromPointer(C.QMagnetometer_NewQMagnetometer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QMagnetometer) Reading() *QMagnetometerReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQMagnetometerReadingFromPointer(C.QMagnetometer_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMagnetometer) ReturnGeoValues() bool {
	if ptr.Pointer() != nil {
		return int8(C.QMagnetometer_ReturnGeoValues(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQMagnetometer_ReturnGeoValuesChanged
func callbackQMagnetometer_ReturnGeoValuesChanged(ptr unsafe.Pointer, returnGeoValues C.char) {
	if signal := qt.GetSignal(ptr, "returnGeoValuesChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(returnGeoValues) != 0)
	}

}

func (ptr *QMagnetometer) ConnectReturnGeoValuesChanged(f func(returnGeoValues bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "returnGeoValuesChanged") {
			C.QMagnetometer_ConnectReturnGeoValuesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "returnGeoValuesChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "returnGeoValuesChanged"); signal != nil {
			f := func(returnGeoValues bool) {
				(*(*func(bool))(signal))(returnGeoValues)
				f(returnGeoValues)
			}
			qt.ConnectSignal(ptr.Pointer(), "returnGeoValuesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "returnGeoValuesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMagnetometer) DisconnectReturnGeoValuesChanged() {
	if ptr.Pointer() != nil {
		C.QMagnetometer_DisconnectReturnGeoValuesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "returnGeoValuesChanged")
	}
}

func (ptr *QMagnetometer) ReturnGeoValuesChanged(returnGeoValues bool) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_ReturnGeoValuesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(returnGeoValues))))
	}
}

func (ptr *QMagnetometer) SetReturnGeoValues(returnGeoValues bool) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_SetReturnGeoValues(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(returnGeoValues))))
	}
}

//export callbackQMagnetometer_DestroyQMagnetometer
func callbackQMagnetometer_DestroyQMagnetometer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QMagnetometer"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQMagnetometerFromPointer(ptr).DestroyQMagnetometerDefault()
	}
}

func (ptr *QMagnetometer) ConnectDestroyQMagnetometer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QMagnetometer"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QMagnetometer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QMagnetometer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMagnetometer) DisconnectDestroyQMagnetometer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QMagnetometer")
	}
}

func (ptr *QMagnetometer) DestroyQMagnetometer() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMagnetometer_DestroyQMagnetometer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMagnetometer) DestroyQMagnetometerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMagnetometer_DestroyQMagnetometerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QMagnetometerFilter struct {
	QSensorFilter
}

type QMagnetometerFilter_ITF interface {
	QSensorFilter_ITF
	QMagnetometerFilter_PTR() *QMagnetometerFilter
}

func (ptr *QMagnetometerFilter) QMagnetometerFilter_PTR() *QMagnetometerFilter {
	return ptr
}

func (ptr *QMagnetometerFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QMagnetometerFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQMagnetometerFilter(ptr QMagnetometerFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerFilter_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerFilterFromPointer(ptr unsafe.Pointer) (n *QMagnetometerFilter) {
	n = new(QMagnetometerFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QMagnetometerFilter) DestroyQMagnetometerFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQMagnetometerFilter_Filter
func callbackQMagnetometerFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QMagnetometerReading) bool)(signal))(NewQMagnetometerReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QMagnetometerFilter) ConnectFilter(f func(reading *QMagnetometerReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QMagnetometerReading) bool {
				(*(*func(*QMagnetometerReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMagnetometerFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QMagnetometerFilter) Filter(reading QMagnetometerReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QMagnetometerFilter_Filter(ptr.Pointer(), PointerFromQMagnetometerReading(reading))) != 0
	}
	return false
}

type QMagnetometerReading struct {
	QSensorReading
}

type QMagnetometerReading_ITF interface {
	QSensorReading_ITF
	QMagnetometerReading_PTR() *QMagnetometerReading
}

func (ptr *QMagnetometerReading) QMagnetometerReading_PTR() *QMagnetometerReading {
	return ptr
}

func (ptr *QMagnetometerReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QMagnetometerReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQMagnetometerReading(ptr QMagnetometerReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerReading_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerReadingFromPointer(ptr unsafe.Pointer) (n *QMagnetometerReading) {
	n = new(QMagnetometerReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QMagnetometerReading) CalibrationLevel() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_CalibrationLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) SetCalibrationLevel(calibrationLevel float64) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetCalibrationLevel(ptr.Pointer(), C.double(calibrationLevel))
	}
}

func (ptr *QMagnetometerReading) SetX(x float64) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QMagnetometerReading) SetY(y float64) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QMagnetometerReading) SetZ(z float64) {
	if ptr.Pointer() != nil {
		C.QMagnetometerReading_SetZ(ptr.Pointer(), C.double(z))
	}
}

func (ptr *QMagnetometerReading) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMagnetometerReading) Z() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMagnetometerReading_Z(ptr.Pointer()))
	}
	return 0
}

type QOrientationFilter struct {
	QSensorFilter
}

type QOrientationFilter_ITF interface {
	QSensorFilter_ITF
	QOrientationFilter_PTR() *QOrientationFilter
}

func (ptr *QOrientationFilter) QOrientationFilter_PTR() *QOrientationFilter {
	return ptr
}

func (ptr *QOrientationFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QOrientationFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQOrientationFilter(ptr QOrientationFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationFilter_PTR().Pointer()
	}
	return nil
}

func NewQOrientationFilterFromPointer(ptr unsafe.Pointer) (n *QOrientationFilter) {
	n = new(QOrientationFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QOrientationFilter) DestroyQOrientationFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQOrientationFilter_Filter
func callbackQOrientationFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QOrientationReading) bool)(signal))(NewQOrientationReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QOrientationFilter) ConnectFilter(f func(reading *QOrientationReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QOrientationReading) bool {
				(*(*func(*QOrientationReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QOrientationFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QOrientationFilter) Filter(reading QOrientationReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QOrientationFilter_Filter(ptr.Pointer(), PointerFromQOrientationReading(reading))) != 0
	}
	return false
}

type QOrientationReading struct {
	QSensorReading
}

type QOrientationReading_ITF interface {
	QSensorReading_ITF
	QOrientationReading_PTR() *QOrientationReading
}

func (ptr *QOrientationReading) QOrientationReading_PTR() *QOrientationReading {
	return ptr
}

func (ptr *QOrientationReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QOrientationReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQOrientationReading(ptr QOrientationReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationReading_PTR().Pointer()
	}
	return nil
}

func NewQOrientationReadingFromPointer(ptr unsafe.Pointer) (n *QOrientationReading) {
	n = new(QOrientationReading)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QOrientationReading__Orientation
//QOrientationReading::Orientation
type QOrientationReading__Orientation int64

const (
	QOrientationReading__Undefined QOrientationReading__Orientation = QOrientationReading__Orientation(0)
	QOrientationReading__TopUp     QOrientationReading__Orientation = QOrientationReading__Orientation(1)
	QOrientationReading__TopDown   QOrientationReading__Orientation = QOrientationReading__Orientation(2)
	QOrientationReading__LeftUp    QOrientationReading__Orientation = QOrientationReading__Orientation(3)
	QOrientationReading__RightUp   QOrientationReading__Orientation = QOrientationReading__Orientation(4)
	QOrientationReading__FaceUp    QOrientationReading__Orientation = QOrientationReading__Orientation(5)
	QOrientationReading__FaceDown  QOrientationReading__Orientation = QOrientationReading__Orientation(6)
)

func (ptr *QOrientationReading) Orientation() QOrientationReading__Orientation {
	if ptr.Pointer() != nil {
		return QOrientationReading__Orientation(C.QOrientationReading_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QOrientationReading) SetOrientation(orientation QOrientationReading__Orientation) {
	if ptr.Pointer() != nil {
		C.QOrientationReading_SetOrientation(ptr.Pointer(), C.longlong(orientation))
	}
}

type QOrientationSensor struct {
	QSensor
}

type QOrientationSensor_ITF interface {
	QSensor_ITF
	QOrientationSensor_PTR() *QOrientationSensor
}

func (ptr *QOrientationSensor) QOrientationSensor_PTR() *QOrientationSensor {
	return ptr
}

func (ptr *QOrientationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QOrientationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQOrientationSensor(ptr QOrientationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationSensor_PTR().Pointer()
	}
	return nil
}

func NewQOrientationSensorFromPointer(ptr unsafe.Pointer) (n *QOrientationSensor) {
	n = new(QOrientationSensor)
	n.SetPointer(ptr)
	return
}
func NewQOrientationSensor(parent core.QObject_ITF) *QOrientationSensor {
	tmpValue := NewQOrientationSensorFromPointer(C.QOrientationSensor_NewQOrientationSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QOrientationSensor) Reading() *QOrientationReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQOrientationReadingFromPointer(C.QOrientationSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQOrientationSensor_DestroyQOrientationSensor
func callbackQOrientationSensor_DestroyQOrientationSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QOrientationSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQOrientationSensorFromPointer(ptr).DestroyQOrientationSensorDefault()
	}
}

func (ptr *QOrientationSensor) ConnectDestroyQOrientationSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QOrientationSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QOrientationSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QOrientationSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QOrientationSensor) DisconnectDestroyQOrientationSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QOrientationSensor")
	}
}

func (ptr *QOrientationSensor) DestroyQOrientationSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QOrientationSensor_DestroyQOrientationSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QOrientationSensor) DestroyQOrientationSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QOrientationSensor_DestroyQOrientationSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QPressureFilter struct {
	QSensorFilter
}

type QPressureFilter_ITF interface {
	QSensorFilter_ITF
	QPressureFilter_PTR() *QPressureFilter
}

func (ptr *QPressureFilter) QPressureFilter_PTR() *QPressureFilter {
	return ptr
}

func (ptr *QPressureFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QPressureFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQPressureFilter(ptr QPressureFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureFilter_PTR().Pointer()
	}
	return nil
}

func NewQPressureFilterFromPointer(ptr unsafe.Pointer) (n *QPressureFilter) {
	n = new(QPressureFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QPressureFilter) DestroyQPressureFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQPressureFilter_Filter
func callbackQPressureFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QPressureReading) bool)(signal))(NewQPressureReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPressureFilter) ConnectFilter(f func(reading *QPressureReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QPressureReading) bool {
				(*(*func(*QPressureReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPressureFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QPressureFilter) Filter(reading QPressureReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPressureFilter_Filter(ptr.Pointer(), PointerFromQPressureReading(reading))) != 0
	}
	return false
}

type QPressureReading struct {
	QSensorReading
}

type QPressureReading_ITF interface {
	QSensorReading_ITF
	QPressureReading_PTR() *QPressureReading
}

func (ptr *QPressureReading) QPressureReading_PTR() *QPressureReading {
	return ptr
}

func (ptr *QPressureReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QPressureReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQPressureReading(ptr QPressureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureReading_PTR().Pointer()
	}
	return nil
}

func NewQPressureReadingFromPointer(ptr unsafe.Pointer) (n *QPressureReading) {
	n = new(QPressureReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QPressureReading) Pressure() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPressureReading_Pressure(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPressureReading) SetPressure(pressure float64) {
	if ptr.Pointer() != nil {
		C.QPressureReading_SetPressure(ptr.Pointer(), C.double(pressure))
	}
}

func (ptr *QPressureReading) SetTemperature(temperature float64) {
	if ptr.Pointer() != nil {
		C.QPressureReading_SetTemperature(ptr.Pointer(), C.double(temperature))
	}
}

func (ptr *QPressureReading) Temperature() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPressureReading_Temperature(ptr.Pointer()))
	}
	return 0
}

type QPressureSensor struct {
	QSensor
}

type QPressureSensor_ITF interface {
	QSensor_ITF
	QPressureSensor_PTR() *QPressureSensor
}

func (ptr *QPressureSensor) QPressureSensor_PTR() *QPressureSensor {
	return ptr
}

func (ptr *QPressureSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QPressureSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQPressureSensor(ptr QPressureSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPressureSensor_PTR().Pointer()
	}
	return nil
}

func NewQPressureSensorFromPointer(ptr unsafe.Pointer) (n *QPressureSensor) {
	n = new(QPressureSensor)
	n.SetPointer(ptr)
	return
}
func NewQPressureSensor(parent core.QObject_ITF) *QPressureSensor {
	tmpValue := NewQPressureSensorFromPointer(C.QPressureSensor_NewQPressureSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QPressureSensor) Reading() *QPressureReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQPressureReadingFromPointer(C.QPressureSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQPressureSensor_DestroyQPressureSensor
func callbackQPressureSensor_DestroyQPressureSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPressureSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPressureSensorFromPointer(ptr).DestroyQPressureSensorDefault()
	}
}

func (ptr *QPressureSensor) ConnectDestroyQPressureSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPressureSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPressureSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPressureSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPressureSensor) DisconnectDestroyQPressureSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPressureSensor")
	}
}

func (ptr *QPressureSensor) DestroyQPressureSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPressureSensor_DestroyQPressureSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPressureSensor) DestroyQPressureSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPressureSensor_DestroyQPressureSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QProximityFilter struct {
	QSensorFilter
}

type QProximityFilter_ITF interface {
	QSensorFilter_ITF
	QProximityFilter_PTR() *QProximityFilter
}

func (ptr *QProximityFilter) QProximityFilter_PTR() *QProximityFilter {
	return ptr
}

func (ptr *QProximityFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QProximityFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQProximityFilter(ptr QProximityFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityFilter_PTR().Pointer()
	}
	return nil
}

func NewQProximityFilterFromPointer(ptr unsafe.Pointer) (n *QProximityFilter) {
	n = new(QProximityFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QProximityFilter) DestroyQProximityFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQProximityFilter_Filter
func callbackQProximityFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QProximityReading) bool)(signal))(NewQProximityReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QProximityFilter) ConnectFilter(f func(reading *QProximityReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QProximityReading) bool {
				(*(*func(*QProximityReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QProximityFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QProximityFilter) Filter(reading QProximityReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QProximityFilter_Filter(ptr.Pointer(), PointerFromQProximityReading(reading))) != 0
	}
	return false
}

type QProximityReading struct {
	QSensorReading
}

type QProximityReading_ITF interface {
	QSensorReading_ITF
	QProximityReading_PTR() *QProximityReading
}

func (ptr *QProximityReading) QProximityReading_PTR() *QProximityReading {
	return ptr
}

func (ptr *QProximityReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QProximityReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQProximityReading(ptr QProximityReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximityReading_PTR().Pointer()
	}
	return nil
}

func NewQProximityReadingFromPointer(ptr unsafe.Pointer) (n *QProximityReading) {
	n = new(QProximityReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QProximityReading) Close() bool {
	if ptr.Pointer() != nil {
		return int8(C.QProximityReading_Close(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QProximityReading) SetClose(close bool) {
	if ptr.Pointer() != nil {
		C.QProximityReading_SetClose(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(close))))
	}
}

type QProximitySensor struct {
	QSensor
}

type QProximitySensor_ITF interface {
	QSensor_ITF
	QProximitySensor_PTR() *QProximitySensor
}

func (ptr *QProximitySensor) QProximitySensor_PTR() *QProximitySensor {
	return ptr
}

func (ptr *QProximitySensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QProximitySensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQProximitySensor(ptr QProximitySensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProximitySensor_PTR().Pointer()
	}
	return nil
}

func NewQProximitySensorFromPointer(ptr unsafe.Pointer) (n *QProximitySensor) {
	n = new(QProximitySensor)
	n.SetPointer(ptr)
	return
}
func NewQProximitySensor(parent core.QObject_ITF) *QProximitySensor {
	tmpValue := NewQProximitySensorFromPointer(C.QProximitySensor_NewQProximitySensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QProximitySensor) Reading() *QProximityReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQProximityReadingFromPointer(C.QProximitySensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQProximitySensor_DestroyQProximitySensor
func callbackQProximitySensor_DestroyQProximitySensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QProximitySensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQProximitySensorFromPointer(ptr).DestroyQProximitySensorDefault()
	}
}

func (ptr *QProximitySensor) ConnectDestroyQProximitySensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QProximitySensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QProximitySensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QProximitySensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QProximitySensor) DisconnectDestroyQProximitySensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QProximitySensor")
	}
}

func (ptr *QProximitySensor) DestroyQProximitySensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QProximitySensor_DestroyQProximitySensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QProximitySensor) DestroyQProximitySensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QProximitySensor_DestroyQProximitySensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QRotationFilter struct {
	QSensorFilter
}

type QRotationFilter_ITF interface {
	QSensorFilter_ITF
	QRotationFilter_PTR() *QRotationFilter
}

func (ptr *QRotationFilter) QRotationFilter_PTR() *QRotationFilter {
	return ptr
}

func (ptr *QRotationFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QRotationFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQRotationFilter(ptr QRotationFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationFilter_PTR().Pointer()
	}
	return nil
}

func NewQRotationFilterFromPointer(ptr unsafe.Pointer) (n *QRotationFilter) {
	n = new(QRotationFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QRotationFilter) DestroyQRotationFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQRotationFilter_Filter
func callbackQRotationFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QRotationReading) bool)(signal))(NewQRotationReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QRotationFilter) ConnectFilter(f func(reading *QRotationReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QRotationReading) bool {
				(*(*func(*QRotationReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRotationFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QRotationFilter) Filter(reading QRotationReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRotationFilter_Filter(ptr.Pointer(), PointerFromQRotationReading(reading))) != 0
	}
	return false
}

type QRotationReading struct {
	QSensorReading
}

type QRotationReading_ITF interface {
	QSensorReading_ITF
	QRotationReading_PTR() *QRotationReading
}

func (ptr *QRotationReading) QRotationReading_PTR() *QRotationReading {
	return ptr
}

func (ptr *QRotationReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QRotationReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQRotationReading(ptr QRotationReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationReading_PTR().Pointer()
	}
	return nil
}

func NewQRotationReadingFromPointer(ptr unsafe.Pointer) (n *QRotationReading) {
	n = new(QRotationReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QRotationReading) SetFromEuler(x float64, y float64, z float64) {
	if ptr.Pointer() != nil {
		C.QRotationReading_SetFromEuler(ptr.Pointer(), C.double(x), C.double(y), C.double(z))
	}
}

func (ptr *QRotationReading) X() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) Y() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRotationReading) Z() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRotationReading_Z(ptr.Pointer()))
	}
	return 0
}

type QRotationSensor struct {
	QSensor
}

type QRotationSensor_ITF interface {
	QSensor_ITF
	QRotationSensor_PTR() *QRotationSensor
}

func (ptr *QRotationSensor) QRotationSensor_PTR() *QRotationSensor {
	return ptr
}

func (ptr *QRotationSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QRotationSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQRotationSensor(ptr QRotationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationSensor_PTR().Pointer()
	}
	return nil
}

func NewQRotationSensorFromPointer(ptr unsafe.Pointer) (n *QRotationSensor) {
	n = new(QRotationSensor)
	n.SetPointer(ptr)
	return
}
func NewQRotationSensor(parent core.QObject_ITF) *QRotationSensor {
	tmpValue := NewQRotationSensorFromPointer(C.QRotationSensor_NewQRotationSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QRotationSensor) HasZ() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRotationSensor_HasZ(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQRotationSensor_HasZChanged
func callbackQRotationSensor_HasZChanged(ptr unsafe.Pointer, hasZ C.char) {
	if signal := qt.GetSignal(ptr, "hasZChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(hasZ) != 0)
	}

}

func (ptr *QRotationSensor) ConnectHasZChanged(f func(hasZ bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hasZChanged") {
			C.QRotationSensor_ConnectHasZChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hasZChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hasZChanged"); signal != nil {
			f := func(hasZ bool) {
				(*(*func(bool))(signal))(hasZ)
				f(hasZ)
			}
			qt.ConnectSignal(ptr.Pointer(), "hasZChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasZChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRotationSensor) DisconnectHasZChanged() {
	if ptr.Pointer() != nil {
		C.QRotationSensor_DisconnectHasZChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hasZChanged")
	}
}

func (ptr *QRotationSensor) HasZChanged(hasZ bool) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_HasZChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hasZ))))
	}
}

func (ptr *QRotationSensor) Reading() *QRotationReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQRotationReadingFromPointer(C.QRotationSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRotationSensor) SetHasZ(hasZ bool) {
	if ptr.Pointer() != nil {
		C.QRotationSensor_SetHasZ(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hasZ))))
	}
}

//export callbackQRotationSensor_DestroyQRotationSensor
func callbackQRotationSensor_DestroyQRotationSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QRotationSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQRotationSensorFromPointer(ptr).DestroyQRotationSensorDefault()
	}
}

func (ptr *QRotationSensor) ConnectDestroyQRotationSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QRotationSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QRotationSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QRotationSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QRotationSensor) DisconnectDestroyQRotationSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QRotationSensor")
	}
}

func (ptr *QRotationSensor) DestroyQRotationSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRotationSensor_DestroyQRotationSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRotationSensor) DestroyQRotationSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QRotationSensor_DestroyQRotationSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSensor struct {
	core.QObject
}

type QSensor_ITF interface {
	core.QObject_ITF
	QSensor_PTR() *QSensor
}

func (ptr *QSensor) QSensor_PTR() *QSensor {
	return ptr
}

func (ptr *QSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensor(ptr QSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func NewQSensorFromPointer(ptr unsafe.Pointer) (n *QSensor) {
	n = new(QSensor)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QSensor__AxesOrientationMode
//QSensor::AxesOrientationMode
type QSensor__AxesOrientationMode int64

const (
	QSensor__FixedOrientation     QSensor__AxesOrientationMode = QSensor__AxesOrientationMode(0)
	QSensor__AutomaticOrientation QSensor__AxesOrientationMode = QSensor__AxesOrientationMode(1)
	QSensor__UserOrientation      QSensor__AxesOrientationMode = QSensor__AxesOrientationMode(2)
)

//go:generate stringer -type=QSensor__Feature
//QSensor::Feature
type QSensor__Feature int64

const (
	QSensor__Buffering                 QSensor__Feature = QSensor__Feature(0)
	QSensor__AlwaysOn                  QSensor__Feature = QSensor__Feature(1)
	QSensor__GeoValues                 QSensor__Feature = QSensor__Feature(2)
	QSensor__FieldOfView               QSensor__Feature = QSensor__Feature(3)
	QSensor__AccelerationMode          QSensor__Feature = QSensor__Feature(4)
	QSensor__SkipDuplicates            QSensor__Feature = QSensor__Feature(5)
	QSensor__AxesOrientation           QSensor__Feature = QSensor__Feature(6)
	QSensor__PressureSensorTemperature QSensor__Feature = QSensor__Feature(7)
	QSensor__Reserved                  QSensor__Feature = QSensor__Feature(257)
)

func NewQSensor(ty core.QByteArray_ITF, parent core.QObject_ITF) *QSensor {
	tmpValue := NewQSensorFromPointer(C.QSensor_NewQSensor(core.PointerFromQByteArray(ty), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSensor_ActiveChanged
func callbackQSensor_ActiveChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activeChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensor) ConnectActiveChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activeChanged") {
			C.QSensor_ConnectActiveChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "activeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activeChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "activeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activeChanged")
	}
}

func (ptr *QSensor) ActiveChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_ActiveChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) AddFilter(filter QSensorFilter_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_AddFilter(ptr.Pointer(), PointerFromQSensorFilter(filter))
	}
}

//export callbackQSensor_AlwaysOnChanged
func callbackQSensor_AlwaysOnChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "alwaysOnChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensor) ConnectAlwaysOnChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "alwaysOnChanged") {
			C.QSensor_ConnectAlwaysOnChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "alwaysOnChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "alwaysOnChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "alwaysOnChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "alwaysOnChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectAlwaysOnChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAlwaysOnChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "alwaysOnChanged")
	}
}

func (ptr *QSensor) AlwaysOnChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_AlwaysOnChanged(ptr.Pointer())
	}
}

//export callbackQSensor_AvailableSensorsChanged
func callbackQSensor_AvailableSensorsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "availableSensorsChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensor) ConnectAvailableSensorsChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "availableSensorsChanged") {
			C.QSensor_ConnectAvailableSensorsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "availableSensorsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "availableSensorsChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "availableSensorsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "availableSensorsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectAvailableSensorsChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAvailableSensorsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "availableSensorsChanged")
	}
}

func (ptr *QSensor) AvailableSensorsChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_AvailableSensorsChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) AxesOrientationMode() QSensor__AxesOrientationMode {
	if ptr.Pointer() != nil {
		return QSensor__AxesOrientationMode(C.QSensor_AxesOrientationMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQSensor_AxesOrientationModeChanged
func callbackQSensor_AxesOrientationModeChanged(ptr unsafe.Pointer, axesOrientationMode C.longlong) {
	if signal := qt.GetSignal(ptr, "axesOrientationModeChanged"); signal != nil {
		(*(*func(QSensor__AxesOrientationMode))(signal))(QSensor__AxesOrientationMode(axesOrientationMode))
	}

}

func (ptr *QSensor) ConnectAxesOrientationModeChanged(f func(axesOrientationMode QSensor__AxesOrientationMode)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "axesOrientationModeChanged") {
			C.QSensor_ConnectAxesOrientationModeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "axesOrientationModeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "axesOrientationModeChanged"); signal != nil {
			f := func(axesOrientationMode QSensor__AxesOrientationMode) {
				(*(*func(QSensor__AxesOrientationMode))(signal))(axesOrientationMode)
				f(axesOrientationMode)
			}
			qt.ConnectSignal(ptr.Pointer(), "axesOrientationModeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "axesOrientationModeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectAxesOrientationModeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectAxesOrientationModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "axesOrientationModeChanged")
	}
}

func (ptr *QSensor) AxesOrientationModeChanged(axesOrientationMode QSensor__AxesOrientationMode) {
	if ptr.Pointer() != nil {
		C.QSensor_AxesOrientationModeChanged(ptr.Pointer(), C.longlong(axesOrientationMode))
	}
}

func (ptr *QSensor) BufferSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_BufferSize(ptr.Pointer())))
	}
	return 0
}

//export callbackQSensor_BufferSizeChanged
func callbackQSensor_BufferSizeChanged(ptr unsafe.Pointer, bufferSize C.int) {
	if signal := qt.GetSignal(ptr, "bufferSizeChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(bufferSize)))
	}

}

func (ptr *QSensor) ConnectBufferSizeChanged(f func(bufferSize int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "bufferSizeChanged") {
			C.QSensor_ConnectBufferSizeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "bufferSizeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "bufferSizeChanged"); signal != nil {
			f := func(bufferSize int) {
				(*(*func(int))(signal))(bufferSize)
				f(bufferSize)
			}
			qt.ConnectSignal(ptr.Pointer(), "bufferSizeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "bufferSizeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "bufferSizeChanged")
	}
}

func (ptr *QSensor) BufferSizeChanged(bufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_BufferSizeChanged(ptr.Pointer(), C.int(int32(bufferSize)))
	}
}

//export callbackQSensor_BusyChanged
func callbackQSensor_BusyChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "busyChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensor) ConnectBusyChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "busyChanged") {
			C.QSensor_ConnectBusyChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "busyChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "busyChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "busyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "busyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectBusyChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectBusyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "busyChanged")
	}
}

func (ptr *QSensor) BusyChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_BusyChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) ConnectToBackend() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_ConnectToBackend(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) CurrentOrientation() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_CurrentOrientation(ptr.Pointer())))
	}
	return 0
}

//export callbackQSensor_CurrentOrientationChanged
func callbackQSensor_CurrentOrientationChanged(ptr unsafe.Pointer, currentOrientation C.int) {
	if signal := qt.GetSignal(ptr, "currentOrientationChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(currentOrientation)))
	}

}

func (ptr *QSensor) ConnectCurrentOrientationChanged(f func(currentOrientation int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "currentOrientationChanged") {
			C.QSensor_ConnectCurrentOrientationChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "currentOrientationChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "currentOrientationChanged"); signal != nil {
			f := func(currentOrientation int) {
				(*(*func(int))(signal))(currentOrientation)
				f(currentOrientation)
			}
			qt.ConnectSignal(ptr.Pointer(), "currentOrientationChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "currentOrientationChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectCurrentOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectCurrentOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "currentOrientationChanged")
	}
}

func (ptr *QSensor) CurrentOrientationChanged(currentOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_CurrentOrientationChanged(ptr.Pointer(), C.int(int32(currentOrientation)))
	}
}

func (ptr *QSensor) DataRate() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_DataRate(ptr.Pointer())))
	}
	return 0
}

//export callbackQSensor_DataRateChanged
func callbackQSensor_DataRateChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dataRateChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensor) ConnectDataRateChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataRateChanged") {
			C.QSensor_ConnectDataRateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "dataRateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dataRateChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "dataRateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dataRateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectDataRateChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectDataRateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dataRateChanged")
	}
}

func (ptr *QSensor) DataRateChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DataRateChanged(ptr.Pointer())
	}
}

func QSensor_DefaultSensorForType(ty core.QByteArray_ITF) *core.QByteArray {
	tmpValue := core.NewQByteArrayFromPointer(C.QSensor_QSensor_DefaultSensorForType(core.PointerFromQByteArray(ty)))
	qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *QSensor) DefaultSensorForType(ty core.QByteArray_ITF) *core.QByteArray {
	tmpValue := core.NewQByteArrayFromPointer(C.QSensor_QSensor_DefaultSensorForType(core.PointerFromQByteArray(ty)))
	qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *QSensor) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSensor_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensor) EfficientBufferSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_EfficientBufferSize(ptr.Pointer())))
	}
	return 0
}

//export callbackQSensor_EfficientBufferSizeChanged
func callbackQSensor_EfficientBufferSizeChanged(ptr unsafe.Pointer, efficientBufferSize C.int) {
	if signal := qt.GetSignal(ptr, "efficientBufferSizeChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(efficientBufferSize)))
	}

}

func (ptr *QSensor) ConnectEfficientBufferSizeChanged(f func(efficientBufferSize int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "efficientBufferSizeChanged") {
			C.QSensor_ConnectEfficientBufferSizeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "efficientBufferSizeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "efficientBufferSizeChanged"); signal != nil {
			f := func(efficientBufferSize int) {
				(*(*func(int))(signal))(efficientBufferSize)
				f(efficientBufferSize)
			}
			qt.ConnectSignal(ptr.Pointer(), "efficientBufferSizeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "efficientBufferSizeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectEfficientBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectEfficientBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "efficientBufferSizeChanged")
	}
}

func (ptr *QSensor) EfficientBufferSizeChanged(efficientBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_EfficientBufferSizeChanged(ptr.Pointer(), C.int(int32(efficientBufferSize)))
	}
}

func (ptr *QSensor) Error() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_Error(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) Filters() []*QSensorFilter {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSensors_PackedList) []*QSensorFilter {
			out := make([]*QSensorFilter, int(l.len))
			tmpList := NewQSensorFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__filters_atList(i)
			}
			return out
		}(C.QSensor_Filters(ptr.Pointer()))
	}
	return make([]*QSensorFilter, 0)
}

func (ptr *QSensor) Identifier() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSensor_Identifier(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) IsActive() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_IsActive(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) IsAlwaysOn() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_IsAlwaysOn(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) IsBusy() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_IsBusy(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) IsConnectedToBackend() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_IsConnectedToBackend(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) IsFeatureSupported(feature QSensor__Feature) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_IsFeatureSupported(ptr.Pointer(), C.longlong(feature))) != 0
	}
	return false
}

func (ptr *QSensor) MaxBufferSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_MaxBufferSize(ptr.Pointer())))
	}
	return 0
}

//export callbackQSensor_MaxBufferSizeChanged
func callbackQSensor_MaxBufferSizeChanged(ptr unsafe.Pointer, maxBufferSize C.int) {
	if signal := qt.GetSignal(ptr, "maxBufferSizeChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(maxBufferSize)))
	}

}

func (ptr *QSensor) ConnectMaxBufferSizeChanged(f func(maxBufferSize int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "maxBufferSizeChanged") {
			C.QSensor_ConnectMaxBufferSizeChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "maxBufferSizeChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "maxBufferSizeChanged"); signal != nil {
			f := func(maxBufferSize int) {
				(*(*func(int))(signal))(maxBufferSize)
				f(maxBufferSize)
			}
			qt.ConnectSignal(ptr.Pointer(), "maxBufferSizeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maxBufferSizeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectMaxBufferSizeChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectMaxBufferSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "maxBufferSizeChanged")
	}
}

func (ptr *QSensor) MaxBufferSizeChanged(maxBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_MaxBufferSizeChanged(ptr.Pointer(), C.int(int32(maxBufferSize)))
	}
}

func (ptr *QSensor) OutputRange() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_OutputRange(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensor) Reading() *QSensorReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQSensorReadingFromPointer(C.QSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSensor_ReadingChanged
func callbackQSensor_ReadingChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readingChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensor) ConnectReadingChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "readingChanged") {
			C.QSensor_ConnectReadingChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "readingChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "readingChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "readingChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readingChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectReadingChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectReadingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "readingChanged")
	}
}

func (ptr *QSensor) ReadingChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_ReadingChanged(ptr.Pointer())
	}
}

func (ptr *QSensor) RemoveFilter(filter QSensorFilter_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_RemoveFilter(ptr.Pointer(), PointerFromQSensorFilter(filter))
	}
}

//export callbackQSensor_SensorError
func callbackQSensor_SensorError(ptr unsafe.Pointer, error C.int) {
	if signal := qt.GetSignal(ptr, "sensorError"); signal != nil {
		(*(*func(int))(signal))(int(int32(error)))
	}

}

func (ptr *QSensor) ConnectSensorError(f func(error int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sensorError") {
			C.QSensor_ConnectSensorError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sensorError")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sensorError"); signal != nil {
			f := func(error int) {
				(*(*func(int))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "sensorError", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sensorError", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectSensorError() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSensorError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sensorError")
	}
}

func (ptr *QSensor) SensorError(error int) {
	if ptr.Pointer() != nil {
		C.QSensor_SensorError(ptr.Pointer(), C.int(int32(error)))
	}
}

func QSensor_SensorTypes() []*core.QByteArray {
	return func(l C.struct_QtSensors_PackedList) []*core.QByteArray {
		out := make([]*core.QByteArray, int(l.len))
		tmpList := NewQSensorFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__sensorTypes_atList(i)
		}
		return out
	}(C.QSensor_QSensor_SensorTypes())
}

func (ptr *QSensor) SensorTypes() []*core.QByteArray {
	return func(l C.struct_QtSensors_PackedList) []*core.QByteArray {
		out := make([]*core.QByteArray, int(l.len))
		tmpList := NewQSensorFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__sensorTypes_atList(i)
		}
		return out
	}(C.QSensor_QSensor_SensorTypes())
}

func QSensor_SensorsForType(ty core.QByteArray_ITF) []*core.QByteArray {
	return func(l C.struct_QtSensors_PackedList) []*core.QByteArray {
		out := make([]*core.QByteArray, int(l.len))
		tmpList := NewQSensorFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__sensorsForType_atList(i)
		}
		return out
	}(C.QSensor_QSensor_SensorsForType(core.PointerFromQByteArray(ty)))
}

func (ptr *QSensor) SensorsForType(ty core.QByteArray_ITF) []*core.QByteArray {
	return func(l C.struct_QtSensors_PackedList) []*core.QByteArray {
		out := make([]*core.QByteArray, int(l.len))
		tmpList := NewQSensorFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__sensorsForType_atList(i)
		}
		return out
	}(C.QSensor_QSensor_SensorsForType(core.PointerFromQByteArray(ty)))
}

func (ptr *QSensor) SetActive(active bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetActive(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(active))))
	}
}

func (ptr *QSensor) SetAlwaysOn(alwaysOn bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetAlwaysOn(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(alwaysOn))))
	}
}

func (ptr *QSensor) SetAxesOrientationMode(axesOrientationMode QSensor__AxesOrientationMode) {
	if ptr.Pointer() != nil {
		C.QSensor_SetAxesOrientationMode(ptr.Pointer(), C.longlong(axesOrientationMode))
	}
}

func (ptr *QSensor) SetBufferSize(bufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetBufferSize(ptr.Pointer(), C.int(int32(bufferSize)))
	}
}

func (ptr *QSensor) SetCurrentOrientation(currentOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetCurrentOrientation(ptr.Pointer(), C.int(int32(currentOrientation)))
	}
}

func (ptr *QSensor) SetDataRate(rate int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetDataRate(ptr.Pointer(), C.int(int32(rate)))
	}
}

func (ptr *QSensor) SetEfficientBufferSize(efficientBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetEfficientBufferSize(ptr.Pointer(), C.int(int32(efficientBufferSize)))
	}
}

func (ptr *QSensor) SetIdentifier(identifier core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_SetIdentifier(ptr.Pointer(), core.PointerFromQByteArray(identifier))
	}
}

func (ptr *QSensor) SetMaxBufferSize(maxBufferSize int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetMaxBufferSize(ptr.Pointer(), C.int(int32(maxBufferSize)))
	}
}

func (ptr *QSensor) SetOutputRange(index int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetOutputRange(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QSensor) SetSkipDuplicates(skipDuplicates bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SetSkipDuplicates(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(skipDuplicates))))
	}
}

func (ptr *QSensor) SetUserOrientation(userOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_SetUserOrientation(ptr.Pointer(), C.int(int32(userOrientation)))
	}
}

func (ptr *QSensor) SkipDuplicates() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_SkipDuplicates(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSensor_SkipDuplicatesChanged
func callbackQSensor_SkipDuplicatesChanged(ptr unsafe.Pointer, skipDuplicates C.char) {
	if signal := qt.GetSignal(ptr, "skipDuplicatesChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(skipDuplicates) != 0)
	}

}

func (ptr *QSensor) ConnectSkipDuplicatesChanged(f func(skipDuplicates bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "skipDuplicatesChanged") {
			C.QSensor_ConnectSkipDuplicatesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "skipDuplicatesChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "skipDuplicatesChanged"); signal != nil {
			f := func(skipDuplicates bool) {
				(*(*func(bool))(signal))(skipDuplicates)
				f(skipDuplicates)
			}
			qt.ConnectSignal(ptr.Pointer(), "skipDuplicatesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "skipDuplicatesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectSkipDuplicatesChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectSkipDuplicatesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "skipDuplicatesChanged")
	}
}

func (ptr *QSensor) SkipDuplicatesChanged(skipDuplicates bool) {
	if ptr.Pointer() != nil {
		C.QSensor_SkipDuplicatesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(skipDuplicates))))
	}
}

//export callbackQSensor_Start
func callbackQSensor_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorFromPointer(ptr).StartDefault())))
}

func (ptr *QSensor) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "start")
	}
}

func (ptr *QSensor) Start() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_Start(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensor) StartDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_StartDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSensor_Stop
func callbackQSensor_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorFromPointer(ptr).StopDefault()
	}
}

func (ptr *QSensor) ConnectStop(f func()) {
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

func (ptr *QSensor) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QSensor) Stop() {
	if ptr.Pointer() != nil {
		C.QSensor_Stop(ptr.Pointer())
	}
}

func (ptr *QSensor) StopDefault() {
	if ptr.Pointer() != nil {
		C.QSensor_StopDefault(ptr.Pointer())
	}
}

func (ptr *QSensor) Type() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSensor_Type(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) UserOrientation() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensor_UserOrientation(ptr.Pointer())))
	}
	return 0
}

//export callbackQSensor_UserOrientationChanged
func callbackQSensor_UserOrientationChanged(ptr unsafe.Pointer, userOrientation C.int) {
	if signal := qt.GetSignal(ptr, "userOrientationChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(userOrientation)))
	}

}

func (ptr *QSensor) ConnectUserOrientationChanged(f func(userOrientation int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "userOrientationChanged") {
			C.QSensor_ConnectUserOrientationChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "userOrientationChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "userOrientationChanged"); signal != nil {
			f := func(userOrientation int) {
				(*(*func(int))(signal))(userOrientation)
				f(userOrientation)
			}
			qt.ConnectSignal(ptr.Pointer(), "userOrientationChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "userOrientationChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectUserOrientationChanged() {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectUserOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "userOrientationChanged")
	}
}

func (ptr *QSensor) UserOrientationChanged(userOrientation int) {
	if ptr.Pointer() != nil {
		C.QSensor_UserOrientationChanged(ptr.Pointer(), C.int(int32(userOrientation)))
	}
}

//export callbackQSensor_DestroyQSensor
func callbackQSensor_DestroyQSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorFromPointer(ptr).DestroyQSensorDefault()
	}
}

func (ptr *QSensor) ConnectDestroyQSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensor) DisconnectDestroyQSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSensor")
	}
}

func (ptr *QSensor) DestroyQSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensor_DestroyQSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensor) DestroyQSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensor_DestroyQSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensor) __filters_atList(i int) *QSensorFilter {
	if ptr.Pointer() != nil {
		return NewQSensorFilterFromPointer(C.QSensor___filters_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QSensor) __filters_setList(i QSensorFilter_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor___filters_setList(ptr.Pointer(), PointerFromQSensorFilter(i))
	}
}

func (ptr *QSensor) __filters_newList() unsafe.Pointer {
	return C.QSensor___filters_newList(ptr.Pointer())
}

func (ptr *QSensor) __sensorTypes_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSensor___sensorTypes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) __sensorTypes_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor___sensorTypes_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSensor) __sensorTypes_newList() unsafe.Pointer {
	return C.QSensor___sensorTypes_newList(ptr.Pointer())
}

func (ptr *QSensor) __sensorsForType_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSensor___sensorsForType_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) __sensorsForType_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor___sensorsForType_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSensor) __sensorsForType_newList() unsafe.Pointer {
	return C.QSensor___sensorsForType_newList(ptr.Pointer())
}

func (ptr *QSensor) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensor___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensor) __children_newList() unsafe.Pointer {
	return C.QSensor___children_newList(ptr.Pointer())
}

func (ptr *QSensor) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSensor___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSensor) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSensor___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSensor) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensor___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensor) __findChildren_newList() unsafe.Pointer {
	return C.QSensor___findChildren_newList(ptr.Pointer())
}

func (ptr *QSensor) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensor___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensor) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensor) __findChildren_newList3() unsafe.Pointer {
	return C.QSensor___findChildren_newList3(ptr.Pointer())
}

//export callbackQSensor_ChildEvent
func callbackQSensor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensor_ConnectNotify
func callbackQSensor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensor_CustomEvent
func callbackQSensor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensor_DeleteLater
func callbackQSensor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensor_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQSensor_Destroyed
func callbackQSensor_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSensor_DisconnectNotify
func callbackQSensor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensor_Event
func callbackQSensor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSensor_EventFilter
func callbackQSensor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSensor_MetaObject
func callbackQSensor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQSensorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSensor_ObjectNameChanged
func callbackQSensor_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSensors_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQSensor_TimerEvent
func callbackQSensor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSensorBackend struct {
	core.QObject
}

type QSensorBackend_ITF interface {
	core.QObject_ITF
	QSensorBackend_PTR() *QSensorBackend
}

func (ptr *QSensorBackend) QSensorBackend_PTR() *QSensorBackend {
	return ptr
}

func (ptr *QSensorBackend) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensorBackend) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensorBackend(ptr QSensorBackend_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func NewQSensorBackendFromPointer(ptr unsafe.Pointer) (n *QSensorBackend) {
	n = new(QSensorBackend)
	n.SetPointer(ptr)
	return
}
func (ptr *QSensorBackend) AddDataRate(min float64, max float64) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_AddDataRate(ptr.Pointer(), C.double(min), C.double(max))
	}
}

func (ptr *QSensorBackend) AddOutputRange(min float64, max float64, accuracy float64) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_AddOutputRange(ptr.Pointer(), C.double(min), C.double(max), C.double(accuracy))
	}
}

//export callbackQSensorBackend_IsFeatureSupported
func callbackQSensorBackend_IsFeatureSupported(ptr unsafe.Pointer, feature C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "isFeatureSupported"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(QSensor__Feature) bool)(signal))(QSensor__Feature(feature)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorBackendFromPointer(ptr).IsFeatureSupportedDefault(QSensor__Feature(feature)))))
}

func (ptr *QSensorBackend) ConnectIsFeatureSupported(f func(feature QSensor__Feature) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isFeatureSupported"); signal != nil {
			f := func(feature QSensor__Feature) bool {
				(*(*func(QSensor__Feature) bool)(signal))(feature)
				return f(feature)
			}
			qt.ConnectSignal(ptr.Pointer(), "isFeatureSupported", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isFeatureSupported", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorBackend) DisconnectIsFeatureSupported() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isFeatureSupported")
	}
}

func (ptr *QSensorBackend) IsFeatureSupported(feature QSensor__Feature) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorBackend_IsFeatureSupported(ptr.Pointer(), C.longlong(feature))) != 0
	}
	return false
}

func (ptr *QSensorBackend) IsFeatureSupportedDefault(feature QSensor__Feature) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorBackend_IsFeatureSupportedDefault(ptr.Pointer(), C.longlong(feature))) != 0
	}
	return false
}

func (ptr *QSensorBackend) NewReadingAvailable() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_NewReadingAvailable(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) Reading() *QSensorReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQSensorReadingFromPointer(C.QSensorBackend_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorBackend) Sensor() *QSensor {
	if ptr.Pointer() != nil {
		tmpValue := NewQSensorFromPointer(C.QSensorBackend_Sensor(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorBackend) SensorBusy() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorBusy(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) SensorError(error int) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorError(ptr.Pointer(), C.int(int32(error)))
	}
}

func (ptr *QSensorBackend) SensorStopped() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorStopped(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) SetDataRates(otherSensor QSensor_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_SetDataRates(ptr.Pointer(), PointerFromQSensor(otherSensor))
	}
}

func (ptr *QSensorBackend) SetDescription(description string) {
	if ptr.Pointer() != nil {
		var descriptionC *C.char
		if description != "" {
			descriptionC = C.CString(description)
			defer C.free(unsafe.Pointer(descriptionC))
		}
		C.QSensorBackend_SetDescription(ptr.Pointer(), C.struct_QtSensors_PackedString{data: descriptionC, len: C.longlong(len(description))})
	}
}

//export callbackQSensorBackend_Start
func callbackQSensorBackend_Start(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "start"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensorBackend) ConnectStart(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorBackend) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "start")
	}
}

func (ptr *QSensorBackend) Start() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_Start(ptr.Pointer())
	}
}

//export callbackQSensorBackend_Stop
func callbackQSensorBackend_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensorBackend) ConnectStop(f func()) {
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

func (ptr *QSensorBackend) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QSensorBackend) Stop() {
	if ptr.Pointer() != nil {
		C.QSensorBackend_Stop(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorBackend___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorBackend) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorBackend) __children_newList() unsafe.Pointer {
	return C.QSensorBackend___children_newList(ptr.Pointer())
}

func (ptr *QSensorBackend) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSensorBackend___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensorBackend) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSensorBackend) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSensorBackend___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSensorBackend) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorBackend___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorBackend) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorBackend) __findChildren_newList() unsafe.Pointer {
	return C.QSensorBackend___findChildren_newList(ptr.Pointer())
}

func (ptr *QSensorBackend) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorBackend___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorBackend) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorBackend) __findChildren_newList3() unsafe.Pointer {
	return C.QSensorBackend___findChildren_newList3(ptr.Pointer())
}

//export callbackQSensorBackend_ChildEvent
func callbackQSensorBackend_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorBackendFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorBackend) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensorBackend_ConnectNotify
func callbackQSensorBackend_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorBackendFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorBackend) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorBackend_CustomEvent
func callbackQSensorBackend_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorBackendFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorBackend) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensorBackend_DeleteLater
func callbackQSensorBackend_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorBackendFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensorBackend) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorBackend_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQSensorBackend_Destroyed
func callbackQSensorBackend_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSensorBackend_DisconnectNotify
func callbackQSensorBackend_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorBackendFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorBackend) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorBackend_Event
func callbackQSensorBackend_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorBackendFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensorBackend) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorBackend_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSensorBackend_EventFilter
func callbackQSensorBackend_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorBackendFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensorBackend) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorBackend_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSensorBackend_MetaObject
func callbackQSensorBackend_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQSensorBackendFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensorBackend) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorBackend_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSensorBackend_ObjectNameChanged
func callbackQSensorBackend_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSensors_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQSensorBackend_TimerEvent
func callbackQSensorBackend_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorBackendFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorBackend) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorBackend_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSensorBackendFactory struct {
	ptr unsafe.Pointer
}

type QSensorBackendFactory_ITF interface {
	QSensorBackendFactory_PTR() *QSensorBackendFactory
}

func (ptr *QSensorBackendFactory) QSensorBackendFactory_PTR() *QSensorBackendFactory {
	return ptr
}

func (ptr *QSensorBackendFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSensorBackendFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSensorBackendFactory(ptr QSensorBackendFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackendFactory_PTR().Pointer()
	}
	return nil
}

func NewQSensorBackendFactoryFromPointer(ptr unsafe.Pointer) (n *QSensorBackendFactory) {
	n = new(QSensorBackendFactory)
	n.SetPointer(ptr)
	return
}
func (ptr *QSensorBackendFactory) DestroyQSensorBackendFactory() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorBackendFactory_CreateBackend
func callbackQSensorBackendFactory_CreateBackend(ptr unsafe.Pointer, sensor unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createBackend"); signal != nil {
		return PointerFromQSensorBackend((*(*func(*QSensor) *QSensorBackend)(signal))(NewQSensorFromPointer(sensor)))
	}

	return PointerFromQSensorBackend(nil)
}

func (ptr *QSensorBackendFactory) ConnectCreateBackend(f func(sensor *QSensor) *QSensorBackend) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createBackend"); signal != nil {
			f := func(sensor *QSensor) *QSensorBackend {
				(*(*func(*QSensor) *QSensorBackend)(signal))(sensor)
				return f(sensor)
			}
			qt.ConnectSignal(ptr.Pointer(), "createBackend", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createBackend", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorBackendFactory) DisconnectCreateBackend() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createBackend")
	}
}

func (ptr *QSensorBackendFactory) CreateBackend(sensor QSensor_ITF) *QSensorBackend {
	if ptr.Pointer() != nil {
		tmpValue := NewQSensorBackendFromPointer(C.QSensorBackendFactory_CreateBackend(ptr.Pointer(), PointerFromQSensor(sensor)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QSensorChangesInterface struct {
	ptr unsafe.Pointer
}

type QSensorChangesInterface_ITF interface {
	QSensorChangesInterface_PTR() *QSensorChangesInterface
}

func (ptr *QSensorChangesInterface) QSensorChangesInterface_PTR() *QSensorChangesInterface {
	return ptr
}

func (ptr *QSensorChangesInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSensorChangesInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSensorChangesInterface(ptr QSensorChangesInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorChangesInterface_PTR().Pointer()
	}
	return nil
}

func NewQSensorChangesInterfaceFromPointer(ptr unsafe.Pointer) (n *QSensorChangesInterface) {
	n = new(QSensorChangesInterface)
	n.SetPointer(ptr)
	return
}
func (ptr *QSensorChangesInterface) DestroyQSensorChangesInterface() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorChangesInterface_SensorsChanged
func callbackQSensorChangesInterface_SensorsChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sensorsChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensorChangesInterface) ConnectSensorsChanged(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sensorsChanged"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sensorsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sensorsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorChangesInterface) DisconnectSensorsChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sensorsChanged")
	}
}

func (ptr *QSensorChangesInterface) SensorsChanged() {
	if ptr.Pointer() != nil {
		C.QSensorChangesInterface_SensorsChanged(ptr.Pointer())
	}
}

type QSensorFilter struct {
	ptr unsafe.Pointer
}

type QSensorFilter_ITF interface {
	QSensorFilter_PTR() *QSensorFilter
}

func (ptr *QSensorFilter) QSensorFilter_PTR() *QSensorFilter {
	return ptr
}

func (ptr *QSensorFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSensorFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSensorFilter(ptr QSensorFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func NewQSensorFilterFromPointer(ptr unsafe.Pointer) (n *QSensorFilter) {
	n = new(QSensorFilter)
	n.SetPointer(ptr)
	return
}

//export callbackQSensorFilter_Filter
func callbackQSensorFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QSensorReading) bool)(signal))(NewQSensorReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSensorFilter) ConnectFilter(f func(reading *QSensorReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QSensorReading) bool {
				(*(*func(*QSensorReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QSensorFilter) Filter(reading QSensorReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorFilter_Filter(ptr.Pointer(), PointerFromQSensorReading(reading))) != 0
	}
	return false
}

//export callbackQSensorFilter_DestroyQSensorFilter
func callbackQSensorFilter_DestroyQSensorFilter(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSensorFilter"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorFilterFromPointer(ptr).DestroyQSensorFilterDefault()
	}
}

func (ptr *QSensorFilter) ConnectDestroyQSensorFilter(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSensorFilter"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSensorFilter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSensorFilter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorFilter) DisconnectDestroyQSensorFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSensorFilter")
	}
}

func (ptr *QSensorFilter) DestroyQSensorFilter() {
	if ptr.Pointer() != nil {
		C.QSensorFilter_DestroyQSensorFilter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorFilter) DestroyQSensorFilterDefault() {
	if ptr.Pointer() != nil {
		C.QSensorFilter_DestroyQSensorFilterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSensorGesture struct {
	core.QObject
}

type QSensorGesture_ITF interface {
	core.QObject_ITF
	QSensorGesture_PTR() *QSensorGesture
}

func (ptr *QSensorGesture) QSensorGesture_PTR() *QSensorGesture {
	return ptr
}

func (ptr *QSensorGesture) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensorGesture) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensorGesture(ptr QSensorGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesture_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureFromPointer(ptr unsafe.Pointer) (n *QSensorGesture) {
	n = new(QSensorGesture)
	n.SetPointer(ptr)
	return
}
func NewQSensorGesture(ids []string, parent core.QObject_ITF) *QSensorGesture {
	idsC := C.CString(strings.Join(ids, "¡¦!"))
	defer C.free(unsafe.Pointer(idsC))
	tmpValue := NewQSensorGestureFromPointer(C.QSensorGesture_NewQSensorGesture(C.struct_QtSensors_PackedString{data: idsC, len: C.longlong(len(strings.Join(ids, "¡¦!")))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSensorGesture_Detected
func callbackQSensorGesture_Detected(ptr unsafe.Pointer, vqs C.struct_QtSensors_PackedString) {
	if signal := qt.GetSignal(ptr, "detected"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	}

}

func (ptr *QSensorGesture) ConnectDetected(f func(vqs string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "detected") {
			C.QSensorGesture_ConnectDetected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "detected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "detected"); signal != nil {
			f := func(vqs string) {
				(*(*func(string))(signal))(vqs)
				f(vqs)
			}
			qt.ConnectSignal(ptr.Pointer(), "detected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "detected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGesture) DisconnectDetected() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_DisconnectDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "detected")
	}
}

func (ptr *QSensorGesture) Detected(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QSensorGesture_Detected(ptr.Pointer(), C.struct_QtSensors_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QSensorGesture) GestureSignals() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QSensorGesture_GestureSignals(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) InvalidIds() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QSensorGesture_InvalidIds(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) IsActive() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGesture_IsActive(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensorGesture) StartDetection() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_StartDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) StopDetection() {
	if ptr.Pointer() != nil {
		C.QSensorGesture_StopDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) ValidIds() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QSensorGesture_ValidIds(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQSensorGesture_DestroyQSensorGesture
func callbackQSensorGesture_DestroyQSensorGesture(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSensorGesture"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorGestureFromPointer(ptr).DestroyQSensorGestureDefault()
	}
}

func (ptr *QSensorGesture) ConnectDestroyQSensorGesture(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSensorGesture"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSensorGesture", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSensorGesture", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGesture) DisconnectDestroyQSensorGesture() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSensorGesture")
	}
}

func (ptr *QSensorGesture) DestroyQSensorGesture() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorGesture_DestroyQSensorGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGesture) DestroyQSensorGestureDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorGesture_DestroyQSensorGestureDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGesture) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorGesture___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGesture) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorGesture) __children_newList() unsafe.Pointer {
	return C.QSensorGesture___children_newList(ptr.Pointer())
}

func (ptr *QSensorGesture) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSensorGesture___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGesture) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSensorGesture) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSensorGesture___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSensorGesture) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorGesture___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGesture) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorGesture) __findChildren_newList() unsafe.Pointer {
	return C.QSensorGesture___findChildren_newList(ptr.Pointer())
}

func (ptr *QSensorGesture) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorGesture___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGesture) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorGesture) __findChildren_newList3() unsafe.Pointer {
	return C.QSensorGesture___findChildren_newList3(ptr.Pointer())
}

//export callbackQSensorGesture_ChildEvent
func callbackQSensorGesture_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensorGesture_ConnectNotify
func callbackQSensorGesture_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGesture) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGesture_CustomEvent
func callbackQSensorGesture_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensorGesture_DeleteLater
func callbackQSensorGesture_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorGestureFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensorGesture) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorGesture_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQSensorGesture_Destroyed
func callbackQSensorGesture_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSensorGesture_DisconnectNotify
func callbackQSensorGesture_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGesture) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGesture_Event
func callbackQSensorGesture_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensorGesture) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGesture_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSensorGesture_EventFilter
func callbackQSensorGesture_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensorGesture) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGesture_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSensorGesture_MetaObject
func callbackQSensorGesture_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQSensorGestureFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensorGesture) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorGesture_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSensorGesture_ObjectNameChanged
func callbackQSensorGesture_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSensors_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQSensorGesture_TimerEvent
func callbackQSensorGesture_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSensorGestureManager struct {
	core.QObject
}

type QSensorGestureManager_ITF interface {
	core.QObject_ITF
	QSensorGestureManager_PTR() *QSensorGestureManager
}

func (ptr *QSensorGestureManager) QSensorGestureManager_PTR() *QSensorGestureManager {
	return ptr
}

func (ptr *QSensorGestureManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensorGestureManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensorGestureManager(ptr QSensorGestureManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureManager_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureManagerFromPointer(ptr unsafe.Pointer) (n *QSensorGestureManager) {
	n = new(QSensorGestureManager)
	n.SetPointer(ptr)
	return
}
func NewQSensorGestureManager(parent core.QObject_ITF) *QSensorGestureManager {
	tmpValue := NewQSensorGestureManagerFromPointer(C.QSensorGestureManager_NewQSensorGestureManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSensorGestureManager) GestureIds() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QSensorGestureManager_GestureIds(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQSensorGestureManager_NewSensorGestureAvailable
func callbackQSensorGestureManager_NewSensorGestureAvailable(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "newSensorGestureAvailable"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensorGestureManager) ConnectNewSensorGestureAvailable(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "newSensorGestureAvailable") {
			C.QSensorGestureManager_ConnectNewSensorGestureAvailable(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "newSensorGestureAvailable")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "newSensorGestureAvailable"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "newSensorGestureAvailable", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newSensorGestureAvailable", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGestureManager) DisconnectNewSensorGestureAvailable() {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DisconnectNewSensorGestureAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "newSensorGestureAvailable")
	}
}

func (ptr *QSensorGestureManager) NewSensorGestureAvailable() {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_NewSensorGestureAvailable(ptr.Pointer())
	}
}

func (ptr *QSensorGestureManager) RecognizerSignals(gestureId string) []string {
	if ptr.Pointer() != nil {
		var gestureIdC *C.char
		if gestureId != "" {
			gestureIdC = C.CString(gestureId)
			defer C.free(unsafe.Pointer(gestureIdC))
		}
		return unpackStringList(cGoUnpackString(C.QSensorGestureManager_RecognizerSignals(ptr.Pointer(), C.struct_QtSensors_PackedString{data: gestureIdC, len: C.longlong(len(gestureId))})))
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureManager) RegisterSensorGestureRecognizer(recognizer QSensorGestureRecognizer_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGestureManager_RegisterSensorGestureRecognizer(ptr.Pointer(), PointerFromQSensorGestureRecognizer(recognizer))) != 0
	}
	return false
}

func QSensorGestureManager_SensorGestureRecognizer(id string) *QSensorGestureRecognizer {
	var idC *C.char
	if id != "" {
		idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
	}
	tmpValue := NewQSensorGestureRecognizerFromPointer(C.QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(C.struct_QtSensors_PackedString{data: idC, len: C.longlong(len(id))}))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSensorGestureManager) SensorGestureRecognizer(id string) *QSensorGestureRecognizer {
	var idC *C.char
	if id != "" {
		idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
	}
	tmpValue := NewQSensorGestureRecognizerFromPointer(C.QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(C.struct_QtSensors_PackedString{data: idC, len: C.longlong(len(id))}))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSensorGestureManager_DestroyQSensorGestureManager
func callbackQSensorGestureManager_DestroyQSensorGestureManager(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSensorGestureManager"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorGestureManagerFromPointer(ptr).DestroyQSensorGestureManagerDefault()
	}
}

func (ptr *QSensorGestureManager) ConnectDestroyQSensorGestureManager(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSensorGestureManager"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSensorGestureManager", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSensorGestureManager", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGestureManager) DisconnectDestroyQSensorGestureManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSensorGestureManager")
	}
}

func (ptr *QSensorGestureManager) DestroyQSensorGestureManager() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorGestureManager_DestroyQSensorGestureManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGestureManager) DestroyQSensorGestureManagerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorGestureManager_DestroyQSensorGestureManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGestureManager) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorGestureManager___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGestureManager) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorGestureManager) __children_newList() unsafe.Pointer {
	return C.QSensorGestureManager___children_newList(ptr.Pointer())
}

func (ptr *QSensorGestureManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSensorGestureManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGestureManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSensorGestureManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSensorGestureManager___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSensorGestureManager) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorGestureManager___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGestureManager) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorGestureManager) __findChildren_newList() unsafe.Pointer {
	return C.QSensorGestureManager___findChildren_newList(ptr.Pointer())
}

func (ptr *QSensorGestureManager) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorGestureManager___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGestureManager) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorGestureManager) __findChildren_newList3() unsafe.Pointer {
	return C.QSensorGestureManager___findChildren_newList3(ptr.Pointer())
}

//export callbackQSensorGestureManager_ChildEvent
func callbackQSensorGestureManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGestureManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensorGestureManager_ConnectNotify
func callbackQSensorGestureManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGestureManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGestureManager_CustomEvent
func callbackQSensorGestureManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGestureManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensorGestureManager_DeleteLater
func callbackQSensorGestureManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorGestureManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensorGestureManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorGestureManager_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQSensorGestureManager_Destroyed
func callbackQSensorGestureManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSensorGestureManager_DisconnectNotify
func callbackQSensorGestureManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGestureManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGestureManager_Event
func callbackQSensorGestureManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensorGestureManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGestureManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSensorGestureManager_EventFilter
func callbackQSensorGestureManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensorGestureManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGestureManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSensorGestureManager_MetaObject
func callbackQSensorGestureManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQSensorGestureManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensorGestureManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorGestureManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSensorGestureManager_ObjectNameChanged
func callbackQSensorGestureManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSensors_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQSensorGestureManager_TimerEvent
func callbackQSensorGestureManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGestureManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSensorGesturePluginInterface struct {
	ptr unsafe.Pointer
}

type QSensorGesturePluginInterface_ITF interface {
	QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface
}

func (ptr *QSensorGesturePluginInterface) QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface {
	return ptr
}

func (ptr *QSensorGesturePluginInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSensorGesturePluginInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSensorGesturePluginInterface(ptr QSensorGesturePluginInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesturePluginInterface_PTR().Pointer()
	}
	return nil
}

func NewQSensorGesturePluginInterfaceFromPointer(ptr unsafe.Pointer) (n *QSensorGesturePluginInterface) {
	n = new(QSensorGesturePluginInterface)
	n.SetPointer(ptr)
	return
}
func NewQSensorGesturePluginInterface() *QSensorGesturePluginInterface {
	return NewQSensorGesturePluginInterfaceFromPointer(C.QSensorGesturePluginInterface_NewQSensorGesturePluginInterface())
}

//export callbackQSensorGesturePluginInterface_CreateRecognizers
func callbackQSensorGesturePluginInterface_CreateRecognizers(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createRecognizers"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQSensorGesturePluginInterfaceFromPointer(NewQSensorGesturePluginInterfaceFromPointer(nil).__createRecognizers_newList())
			for _, v := range (*(*func() []*QSensorGestureRecognizer)(signal))() {
				tmpList.__createRecognizers_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQSensorGesturePluginInterfaceFromPointer(NewQSensorGesturePluginInterfaceFromPointer(nil).__createRecognizers_newList())
		for _, v := range make([]*QSensorGestureRecognizer, 0) {
			tmpList.__createRecognizers_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QSensorGesturePluginInterface) ConnectCreateRecognizers(f func() []*QSensorGestureRecognizer) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createRecognizers"); signal != nil {
			f := func() []*QSensorGestureRecognizer {
				(*(*func() []*QSensorGestureRecognizer)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "createRecognizers", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createRecognizers", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGesturePluginInterface) DisconnectCreateRecognizers() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createRecognizers")
	}
}

func (ptr *QSensorGesturePluginInterface) CreateRecognizers() []*QSensorGestureRecognizer {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSensors_PackedList) []*QSensorGestureRecognizer {
			out := make([]*QSensorGestureRecognizer, int(l.len))
			tmpList := NewQSensorGesturePluginInterfaceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__createRecognizers_atList(i)
			}
			return out
		}(C.QSensorGesturePluginInterface_CreateRecognizers(ptr.Pointer()))
	}
	return make([]*QSensorGestureRecognizer, 0)
}

//export callbackQSensorGesturePluginInterface_Name
func callbackQSensorGesturePluginInterface_Name(ptr unsafe.Pointer) C.struct_QtSensors_PackedString {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_QtSensors_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtSensors_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QSensorGesturePluginInterface) ConnectName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "name"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "name", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "name", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGesturePluginInterface) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "name")
	}
}

func (ptr *QSensorGesturePluginInterface) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSensorGesturePluginInterface_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQSensorGesturePluginInterface_SupportedIds
func callbackQSensorGesturePluginInterface_SupportedIds(ptr unsafe.Pointer) C.struct_QtSensors_PackedString {
	if signal := qt.GetSignal(ptr, "supportedIds"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_QtSensors_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := make([]string, 0)
	return C.struct_QtSensors_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *QSensorGesturePluginInterface) ConnectSupportedIds(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportedIds"); signal != nil {
			f := func() []string {
				(*(*func() []string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "supportedIds", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportedIds", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGesturePluginInterface) DisconnectSupportedIds() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "supportedIds")
	}
}

func (ptr *QSensorGesturePluginInterface) SupportedIds() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QSensorGesturePluginInterface_SupportedIds(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface
func callbackQSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSensorGesturePluginInterface"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorGesturePluginInterfaceFromPointer(ptr).DestroyQSensorGesturePluginInterfaceDefault()
	}
}

func (ptr *QSensorGesturePluginInterface) ConnectDestroyQSensorGesturePluginInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSensorGesturePluginInterface"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSensorGesturePluginInterface", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSensorGesturePluginInterface", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGesturePluginInterface) DisconnectDestroyQSensorGesturePluginInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSensorGesturePluginInterface")
	}
}

func (ptr *QSensorGesturePluginInterface) DestroyQSensorGesturePluginInterface() {
	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGesturePluginInterface) DestroyQSensorGesturePluginInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGesturePluginInterface) __createRecognizers_atList(i int) *QSensorGestureRecognizer {
	if ptr.Pointer() != nil {
		tmpValue := NewQSensorGestureRecognizerFromPointer(C.QSensorGesturePluginInterface___createRecognizers_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGesturePluginInterface) __createRecognizers_setList(i QSensorGestureRecognizer_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface___createRecognizers_setList(ptr.Pointer(), PointerFromQSensorGestureRecognizer(i))
	}
}

func (ptr *QSensorGesturePluginInterface) __createRecognizers_newList() unsafe.Pointer {
	return C.QSensorGesturePluginInterface___createRecognizers_newList(ptr.Pointer())
}

type QSensorGestureRecognizer struct {
	core.QObject
}

type QSensorGestureRecognizer_ITF interface {
	core.QObject_ITF
	QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer
}

func (ptr *QSensorGestureRecognizer) QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer {
	return ptr
}

func (ptr *QSensorGestureRecognizer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensorGestureRecognizer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensorGestureRecognizer(ptr QSensorGestureRecognizer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureRecognizer_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureRecognizerFromPointer(ptr unsafe.Pointer) (n *QSensorGestureRecognizer) {
	n = new(QSensorGestureRecognizer)
	n.SetPointer(ptr)
	return
}
func NewQSensorGestureRecognizer(parent core.QObject_ITF) *QSensorGestureRecognizer {
	tmpValue := NewQSensorGestureRecognizerFromPointer(C.QSensorGestureRecognizer_NewQSensorGestureRecognizer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSensorGestureRecognizer_Create
func callbackQSensorGestureRecognizer_Create(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "create"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensorGestureRecognizer) ConnectCreate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "create"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "create", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "create", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "create")
	}
}

func (ptr *QSensorGestureRecognizer) Create() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_Create(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) CreateBackend() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CreateBackend(ptr.Pointer())
	}
}

//export callbackQSensorGestureRecognizer_Detected
func callbackQSensorGestureRecognizer_Detected(ptr unsafe.Pointer, vqs C.struct_QtSensors_PackedString) {
	if signal := qt.GetSignal(ptr, "detected"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	}

}

func (ptr *QSensorGestureRecognizer) ConnectDetected(f func(vqs string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "detected") {
			C.QSensorGestureRecognizer_ConnectDetected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "detected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "detected"); signal != nil {
			f := func(vqs string) {
				(*(*func(string))(signal))(vqs)
				f(vqs)
			}
			qt.ConnectSignal(ptr.Pointer(), "detected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "detected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectDetected() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DisconnectDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "detected")
	}
}

func (ptr *QSensorGestureRecognizer) Detected(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QSensorGestureRecognizer_Detected(ptr.Pointer(), C.struct_QtSensors_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

func (ptr *QSensorGestureRecognizer) GestureSignals() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QSensorGestureRecognizer_GestureSignals(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackQSensorGestureRecognizer_Id
func callbackQSensorGestureRecognizer_Id(ptr unsafe.Pointer) C.struct_QtSensors_PackedString {
	if signal := qt.GetSignal(ptr, "id"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_QtSensors_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtSensors_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QSensorGestureRecognizer) ConnectId(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "id"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "id", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "id", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "id")
	}
}

func (ptr *QSensorGestureRecognizer) Id() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSensorGestureRecognizer_Id(ptr.Pointer()))
	}
	return ""
}

//export callbackQSensorGestureRecognizer_IsActive
func callbackQSensorGestureRecognizer_IsActive(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isActive"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSensorGestureRecognizer) ConnectIsActive(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isActive"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "isActive", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isActive", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectIsActive() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isActive")
	}
}

func (ptr *QSensorGestureRecognizer) IsActive() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGestureRecognizer_IsActive(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQSensorGestureRecognizer_Start
func callbackQSensorGestureRecognizer_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSensorGestureRecognizer) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "start")
	}
}

func (ptr *QSensorGestureRecognizer) Start() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGestureRecognizer_Start(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensorGestureRecognizer) StartBackend() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StartBackend(ptr.Pointer())
	}
}

//export callbackQSensorGestureRecognizer_Stop
func callbackQSensorGestureRecognizer_Stop(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSensorGestureRecognizer) ConnectStop(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QSensorGestureRecognizer) Stop() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGestureRecognizer_Stop(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSensorGestureRecognizer) StopBackend() {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StopBackend(ptr.Pointer())
	}
}

//export callbackQSensorGestureRecognizer_DestroyQSensorGestureRecognizer
func callbackQSensorGestureRecognizer_DestroyQSensorGestureRecognizer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSensorGestureRecognizer"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).DestroyQSensorGestureRecognizerDefault()
	}
}

func (ptr *QSensorGestureRecognizer) ConnectDestroyQSensorGestureRecognizer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSensorGestureRecognizer"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSensorGestureRecognizer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSensorGestureRecognizer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectDestroyQSensorGestureRecognizer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSensorGestureRecognizer")
	}
}

func (ptr *QSensorGestureRecognizer) DestroyQSensorGestureRecognizer() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGestureRecognizer) DestroyQSensorGestureRecognizerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorGestureRecognizer_DestroyQSensorGestureRecognizerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGestureRecognizer) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorGestureRecognizer___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGestureRecognizer) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorGestureRecognizer) __children_newList() unsafe.Pointer {
	return C.QSensorGestureRecognizer___children_newList(ptr.Pointer())
}

func (ptr *QSensorGestureRecognizer) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSensorGestureRecognizer___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGestureRecognizer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSensorGestureRecognizer) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSensorGestureRecognizer___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSensorGestureRecognizer) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorGestureRecognizer___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGestureRecognizer) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorGestureRecognizer) __findChildren_newList() unsafe.Pointer {
	return C.QSensorGestureRecognizer___findChildren_newList(ptr.Pointer())
}

func (ptr *QSensorGestureRecognizer) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorGestureRecognizer___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorGestureRecognizer) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorGestureRecognizer) __findChildren_newList3() unsafe.Pointer {
	return C.QSensorGestureRecognizer___findChildren_newList3(ptr.Pointer())
}

//export callbackQSensorGestureRecognizer_ChildEvent
func callbackQSensorGestureRecognizer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensorGestureRecognizer_ConnectNotify
func callbackQSensorGestureRecognizer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGestureRecognizer_CustomEvent
func callbackQSensorGestureRecognizer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensorGestureRecognizer_DeleteLater
func callbackQSensorGestureRecognizer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensorGestureRecognizer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorGestureRecognizer_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQSensorGestureRecognizer_Destroyed
func callbackQSensorGestureRecognizer_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSensorGestureRecognizer_DisconnectNotify
func callbackQSensorGestureRecognizer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorGestureRecognizer_Event
func callbackQSensorGestureRecognizer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureRecognizerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensorGestureRecognizer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGestureRecognizer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSensorGestureRecognizer_EventFilter
func callbackQSensorGestureRecognizer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorGestureRecognizerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensorGestureRecognizer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorGestureRecognizer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSensorGestureRecognizer_MetaObject
func callbackQSensorGestureRecognizer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQSensorGestureRecognizerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensorGestureRecognizer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorGestureRecognizer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSensorGestureRecognizer_ObjectNameChanged
func callbackQSensorGestureRecognizer_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSensors_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQSensorGestureRecognizer_TimerEvent
func callbackQSensorGestureRecognizer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QSensorManager struct {
	ptr unsafe.Pointer
}

type QSensorManager_ITF interface {
	QSensorManager_PTR() *QSensorManager
}

func (ptr *QSensorManager) QSensorManager_PTR() *QSensorManager {
	return ptr
}

func (ptr *QSensorManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSensorManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSensorManager(ptr QSensorManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorManager_PTR().Pointer()
	}
	return nil
}

func NewQSensorManagerFromPointer(ptr unsafe.Pointer) (n *QSensorManager) {
	n = new(QSensorManager)
	n.SetPointer(ptr)
	return
}
func (ptr *QSensorManager) DestroyQSensorManager() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func QSensorManager_CreateBackend(sensor QSensor_ITF) *QSensorBackend {
	tmpValue := NewQSensorBackendFromPointer(C.QSensorManager_QSensorManager_CreateBackend(PointerFromQSensor(sensor)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSensorManager) CreateBackend(sensor QSensor_ITF) *QSensorBackend {
	tmpValue := NewQSensorBackendFromPointer(C.QSensorManager_QSensorManager_CreateBackend(PointerFromQSensor(sensor)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QSensorManager_IsBackendRegistered(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) bool {
	return int8(C.QSensorManager_QSensorManager_IsBackendRegistered(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))) != 0
}

func (ptr *QSensorManager) IsBackendRegistered(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) bool {
	return int8(C.QSensorManager_QSensorManager_IsBackendRegistered(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))) != 0
}

func QSensorManager_RegisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF, factory QSensorBackendFactory_ITF) {
	C.QSensorManager_QSensorManager_RegisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier), PointerFromQSensorBackendFactory(factory))
}

func (ptr *QSensorManager) RegisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF, factory QSensorBackendFactory_ITF) {
	C.QSensorManager_QSensorManager_RegisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier), PointerFromQSensorBackendFactory(factory))
}

func QSensorManager_SetDefaultBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	C.QSensorManager_QSensorManager_SetDefaultBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

func (ptr *QSensorManager) SetDefaultBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	C.QSensorManager_QSensorManager_SetDefaultBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

func QSensorManager_UnregisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	C.QSensorManager_QSensorManager_UnregisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

func (ptr *QSensorManager) UnregisterBackend(ty core.QByteArray_ITF, identifier core.QByteArray_ITF) {
	C.QSensorManager_QSensorManager_UnregisterBackend(core.PointerFromQByteArray(ty), core.PointerFromQByteArray(identifier))
}

type QSensorPluginInterface struct {
	ptr unsafe.Pointer
}

type QSensorPluginInterface_ITF interface {
	QSensorPluginInterface_PTR() *QSensorPluginInterface
}

func (ptr *QSensorPluginInterface) QSensorPluginInterface_PTR() *QSensorPluginInterface {
	return ptr
}

func (ptr *QSensorPluginInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSensorPluginInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSensorPluginInterface(ptr QSensorPluginInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorPluginInterface_PTR().Pointer()
	}
	return nil
}

func NewQSensorPluginInterfaceFromPointer(ptr unsafe.Pointer) (n *QSensorPluginInterface) {
	n = new(QSensorPluginInterface)
	n.SetPointer(ptr)
	return
}
func (ptr *QSensorPluginInterface) DestroyQSensorPluginInterface() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSensorPluginInterface_RegisterSensors
func callbackQSensorPluginInterface_RegisterSensors(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "registerSensors"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QSensorPluginInterface) ConnectRegisterSensors(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "registerSensors"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "registerSensors", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "registerSensors", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSensorPluginInterface) DisconnectRegisterSensors() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "registerSensors")
	}
}

func (ptr *QSensorPluginInterface) RegisterSensors() {
	if ptr.Pointer() != nil {
		C.QSensorPluginInterface_RegisterSensors(ptr.Pointer())
	}
}

type QSensorReading struct {
	core.QObject
}

type QSensorReading_ITF interface {
	core.QObject_ITF
	QSensorReading_PTR() *QSensorReading
}

func (ptr *QSensorReading) QSensorReading_PTR() *QSensorReading {
	return ptr
}

func (ptr *QSensorReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSensorReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSensorReading(ptr QSensorReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func NewQSensorReadingFromPointer(ptr unsafe.Pointer) (n *QSensorReading) {
	n = new(QSensorReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QSensorReading) SetTimestamp(timestamp uint64) {
	if ptr.Pointer() != nil {
		C.QSensorReading_SetTimestamp(ptr.Pointer(), C.ulonglong(timestamp))
	}
}

func (ptr *QSensorReading) Timestamp() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QSensorReading_Timestamp(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSensorReading) Value(index int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSensorReading_Value(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSensorReading) ValueCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSensorReading_ValueCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSensorReading) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorReading___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorReading) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorReading) __children_newList() unsafe.Pointer {
	return C.QSensorReading___children_newList(ptr.Pointer())
}

func (ptr *QSensorReading) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSensorReading___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSensorReading) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSensorReading) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSensorReading___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSensorReading) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorReading___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorReading) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorReading) __findChildren_newList() unsafe.Pointer {
	return C.QSensorReading___findChildren_newList(ptr.Pointer())
}

func (ptr *QSensorReading) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSensorReading___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSensorReading) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSensorReading) __findChildren_newList3() unsafe.Pointer {
	return C.QSensorReading___findChildren_newList3(ptr.Pointer())
}

//export callbackQSensorReading_ChildEvent
func callbackQSensorReading_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorReading) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSensorReading_ConnectNotify
func callbackQSensorReading_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorReadingFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorReading) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorReading_CustomEvent
func callbackQSensorReading_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorReading) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSensorReading_DeleteLater
func callbackQSensorReading_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSensorReadingFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSensorReading) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSensorReading_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQSensorReading_Destroyed
func callbackQSensorReading_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSensorReading_DisconnectNotify
func callbackQSensorReading_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSensorReadingFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSensorReading) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSensorReading_Event
func callbackQSensorReading_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorReadingFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSensorReading) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorReading_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSensorReading_EventFilter
func callbackQSensorReading_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSensorReadingFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSensorReading) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSensorReading_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSensorReading_MetaObject
func callbackQSensorReading_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQSensorReadingFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSensorReading) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSensorReading_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSensorReading_ObjectNameChanged
func callbackQSensorReading_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSensors_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQSensorReading_TimerEvent
func callbackQSensorReading_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSensorReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QTapFilter struct {
	QSensorFilter
}

type QTapFilter_ITF interface {
	QSensorFilter_ITF
	QTapFilter_PTR() *QTapFilter
}

func (ptr *QTapFilter) QTapFilter_PTR() *QTapFilter {
	return ptr
}

func (ptr *QTapFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QTapFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQTapFilter(ptr QTapFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapFilter_PTR().Pointer()
	}
	return nil
}

func NewQTapFilterFromPointer(ptr unsafe.Pointer) (n *QTapFilter) {
	n = new(QTapFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QTapFilter) DestroyQTapFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQTapFilter_Filter
func callbackQTapFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QTapReading) bool)(signal))(NewQTapReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QTapFilter) ConnectFilter(f func(reading *QTapReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QTapReading) bool {
				(*(*func(*QTapReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTapFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QTapFilter) Filter(reading QTapReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTapFilter_Filter(ptr.Pointer(), PointerFromQTapReading(reading))) != 0
	}
	return false
}

type QTapReading struct {
	QSensorReading
}

type QTapReading_ITF interface {
	QSensorReading_ITF
	QTapReading_PTR() *QTapReading
}

func (ptr *QTapReading) QTapReading_PTR() *QTapReading {
	return ptr
}

func (ptr *QTapReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QTapReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQTapReading(ptr QTapReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapReading_PTR().Pointer()
	}
	return nil
}

func NewQTapReadingFromPointer(ptr unsafe.Pointer) (n *QTapReading) {
	n = new(QTapReading)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QTapReading__TapDirection
//QTapReading::TapDirection
type QTapReading__TapDirection int64

const (
	QTapReading__Undefined QTapReading__TapDirection = QTapReading__TapDirection(0)
	QTapReading__X         QTapReading__TapDirection = QTapReading__TapDirection(0x0001)
	QTapReading__Y         QTapReading__TapDirection = QTapReading__TapDirection(0x0002)
	QTapReading__Z         QTapReading__TapDirection = QTapReading__TapDirection(0x0004)
	QTapReading__X_Pos     QTapReading__TapDirection = QTapReading__TapDirection(0x0011)
	QTapReading__Y_Pos     QTapReading__TapDirection = QTapReading__TapDirection(0x0022)
	QTapReading__Z_Pos     QTapReading__TapDirection = QTapReading__TapDirection(0x0044)
	QTapReading__X_Neg     QTapReading__TapDirection = QTapReading__TapDirection(0x0101)
	QTapReading__Y_Neg     QTapReading__TapDirection = QTapReading__TapDirection(0x0202)
	QTapReading__Z_Neg     QTapReading__TapDirection = QTapReading__TapDirection(0x0404)
	QTapReading__X_Both    QTapReading__TapDirection = QTapReading__TapDirection(0x0111)
	QTapReading__Y_Both    QTapReading__TapDirection = QTapReading__TapDirection(0x0222)
	QTapReading__Z_Both    QTapReading__TapDirection = QTapReading__TapDirection(0x0444)
)

func (ptr *QTapReading) IsDoubleTap() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTapReading_IsDoubleTap(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTapReading) SetDoubleTap(doubleTap bool) {
	if ptr.Pointer() != nil {
		C.QTapReading_SetDoubleTap(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(doubleTap))))
	}
}

func (ptr *QTapReading) SetTapDirection(tapDirection QTapReading__TapDirection) {
	if ptr.Pointer() != nil {
		C.QTapReading_SetTapDirection(ptr.Pointer(), C.longlong(tapDirection))
	}
}

func (ptr *QTapReading) TapDirection() QTapReading__TapDirection {
	if ptr.Pointer() != nil {
		return QTapReading__TapDirection(C.QTapReading_TapDirection(ptr.Pointer()))
	}
	return 0
}

type QTapSensor struct {
	QSensor
}

type QTapSensor_ITF interface {
	QSensor_ITF
	QTapSensor_PTR() *QTapSensor
}

func (ptr *QTapSensor) QTapSensor_PTR() *QTapSensor {
	return ptr
}

func (ptr *QTapSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QTapSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQTapSensor(ptr QTapSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapSensor_PTR().Pointer()
	}
	return nil
}

func NewQTapSensorFromPointer(ptr unsafe.Pointer) (n *QTapSensor) {
	n = new(QTapSensor)
	n.SetPointer(ptr)
	return
}
func NewQTapSensor(parent core.QObject_ITF) *QTapSensor {
	tmpValue := NewQTapSensorFromPointer(C.QTapSensor_NewQTapSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QTapSensor) Reading() *QTapReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQTapReadingFromPointer(C.QTapSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTapSensor) ReturnDoubleTapEvents() bool {
	if ptr.Pointer() != nil {
		return int8(C.QTapSensor_ReturnDoubleTapEvents(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQTapSensor_ReturnDoubleTapEventsChanged
func callbackQTapSensor_ReturnDoubleTapEventsChanged(ptr unsafe.Pointer, returnDoubleTapEvents C.char) {
	if signal := qt.GetSignal(ptr, "returnDoubleTapEventsChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(returnDoubleTapEvents) != 0)
	}

}

func (ptr *QTapSensor) ConnectReturnDoubleTapEventsChanged(f func(returnDoubleTapEvents bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "returnDoubleTapEventsChanged") {
			C.QTapSensor_ConnectReturnDoubleTapEventsChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "returnDoubleTapEventsChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "returnDoubleTapEventsChanged"); signal != nil {
			f := func(returnDoubleTapEvents bool) {
				(*(*func(bool))(signal))(returnDoubleTapEvents)
				f(returnDoubleTapEvents)
			}
			qt.ConnectSignal(ptr.Pointer(), "returnDoubleTapEventsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "returnDoubleTapEventsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTapSensor) DisconnectReturnDoubleTapEventsChanged() {
	if ptr.Pointer() != nil {
		C.QTapSensor_DisconnectReturnDoubleTapEventsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "returnDoubleTapEventsChanged")
	}
}

func (ptr *QTapSensor) ReturnDoubleTapEventsChanged(returnDoubleTapEvents bool) {
	if ptr.Pointer() != nil {
		C.QTapSensor_ReturnDoubleTapEventsChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(returnDoubleTapEvents))))
	}
}

func (ptr *QTapSensor) SetReturnDoubleTapEvents(returnDoubleTapEvents bool) {
	if ptr.Pointer() != nil {
		C.QTapSensor_SetReturnDoubleTapEvents(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(returnDoubleTapEvents))))
	}
}

//export callbackQTapSensor_DestroyQTapSensor
func callbackQTapSensor_DestroyQTapSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTapSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTapSensorFromPointer(ptr).DestroyQTapSensorDefault()
	}
}

func (ptr *QTapSensor) ConnectDestroyQTapSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTapSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTapSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTapSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTapSensor) DisconnectDestroyQTapSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTapSensor")
	}
}

func (ptr *QTapSensor) DestroyQTapSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTapSensor_DestroyQTapSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTapSensor) DestroyQTapSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTapSensor_DestroyQTapSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QTiltFilter struct {
	QSensorFilter
}

type QTiltFilter_ITF interface {
	QSensorFilter_ITF
	QTiltFilter_PTR() *QTiltFilter
}

func (ptr *QTiltFilter) QTiltFilter_PTR() *QTiltFilter {
	return ptr
}

func (ptr *QTiltFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorFilter_PTR().Pointer()
	}
	return nil
}

func (ptr *QTiltFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorFilter_PTR().SetPointer(p)
	}
}

func PointerFromQTiltFilter(ptr QTiltFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltFilter_PTR().Pointer()
	}
	return nil
}

func NewQTiltFilterFromPointer(ptr unsafe.Pointer) (n *QTiltFilter) {
	n = new(QTiltFilter)
	n.SetPointer(ptr)
	return
}
func (ptr *QTiltFilter) DestroyQTiltFilter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQTiltFilter_Filter
func callbackQTiltFilter_Filter(ptr unsafe.Pointer, reading unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "filter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QTiltReading) bool)(signal))(NewQTiltReadingFromPointer(reading)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QTiltFilter) ConnectFilter(f func(reading *QTiltReading) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "filter"); signal != nil {
			f := func(reading *QTiltReading) bool {
				(*(*func(*QTiltReading) bool)(signal))(reading)
				return f(reading)
			}
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "filter", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTiltFilter) DisconnectFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "filter")
	}
}

func (ptr *QTiltFilter) Filter(reading QTiltReading_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTiltFilter_Filter(ptr.Pointer(), PointerFromQTiltReading(reading))) != 0
	}
	return false
}

type QTiltReading struct {
	QSensorReading
}

type QTiltReading_ITF interface {
	QSensorReading_ITF
	QTiltReading_PTR() *QTiltReading
}

func (ptr *QTiltReading) QTiltReading_PTR() *QTiltReading {
	return ptr
}

func (ptr *QTiltReading) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorReading_PTR().Pointer()
	}
	return nil
}

func (ptr *QTiltReading) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensorReading_PTR().SetPointer(p)
	}
}

func PointerFromQTiltReading(ptr QTiltReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltReading_PTR().Pointer()
	}
	return nil
}

func NewQTiltReadingFromPointer(ptr unsafe.Pointer) (n *QTiltReading) {
	n = new(QTiltReading)
	n.SetPointer(ptr)
	return
}
func (ptr *QTiltReading) SetXRotation(x float64) {
	if ptr.Pointer() != nil {
		C.QTiltReading_SetXRotation(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QTiltReading) SetYRotation(y float64) {
	if ptr.Pointer() != nil {
		C.QTiltReading_SetYRotation(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QTiltReading) XRotation() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTiltReading_XRotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTiltReading) YRotation() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTiltReading_YRotation(ptr.Pointer()))
	}
	return 0
}

type QTiltSensor struct {
	QSensor
}

type QTiltSensor_ITF interface {
	QSensor_ITF
	QTiltSensor_PTR() *QTiltSensor
}

func (ptr *QTiltSensor) QTiltSensor_PTR() *QTiltSensor {
	return ptr
}

func (ptr *QTiltSensor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensor_PTR().Pointer()
	}
	return nil
}

func (ptr *QTiltSensor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSensor_PTR().SetPointer(p)
	}
}

func PointerFromQTiltSensor(ptr QTiltSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTiltSensor_PTR().Pointer()
	}
	return nil
}

func NewQTiltSensorFromPointer(ptr unsafe.Pointer) (n *QTiltSensor) {
	n = new(QTiltSensor)
	n.SetPointer(ptr)
	return
}
func NewQTiltSensor(parent core.QObject_ITF) *QTiltSensor {
	tmpValue := NewQTiltSensorFromPointer(C.QTiltSensor_NewQTiltSensor(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QTiltSensor) Calibrate() {
	if ptr.Pointer() != nil {
		C.QTiltSensor_Calibrate(ptr.Pointer())
	}
}

func (ptr *QTiltSensor) Reading() *QTiltReading {
	if ptr.Pointer() != nil {
		tmpValue := NewQTiltReadingFromPointer(C.QTiltSensor_Reading(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQTiltSensor_DestroyQTiltSensor
func callbackQTiltSensor_DestroyQTiltSensor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTiltSensor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQTiltSensorFromPointer(ptr).DestroyQTiltSensorDefault()
	}
}

func (ptr *QTiltSensor) ConnectDestroyQTiltSensor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTiltSensor"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QTiltSensor", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTiltSensor", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QTiltSensor) DisconnectDestroyQTiltSensor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTiltSensor")
	}
}

func (ptr *QTiltSensor) DestroyQTiltSensor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTiltSensor_DestroyQTiltSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTiltSensor) DestroyQTiltSensorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QTiltSensor_DestroyQTiltSensorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func init() {
	qt.ItfMap["sensors.QAccelerometer_ITF"] = QAccelerometer{}
	qt.FuncMap["sensors.NewQAccelerometer"] = NewQAccelerometer
	qt.EnumMap["sensors.QAccelerometer__Combined"] = int64(QAccelerometer__Combined)
	qt.EnumMap["sensors.QAccelerometer__Gravity"] = int64(QAccelerometer__Gravity)
	qt.EnumMap["sensors.QAccelerometer__User"] = int64(QAccelerometer__User)
	qt.ItfMap["sensors.QAccelerometerFilter_ITF"] = QAccelerometerFilter{}
	qt.ItfMap["sensors.QAccelerometerReading_ITF"] = QAccelerometerReading{}
	qt.ItfMap["sensors.QAltimeter_ITF"] = QAltimeter{}
	qt.FuncMap["sensors.NewQAltimeter"] = NewQAltimeter
	qt.ItfMap["sensors.QAltimeterFilter_ITF"] = QAltimeterFilter{}
	qt.ItfMap["sensors.QAltimeterReading_ITF"] = QAltimeterReading{}
	qt.ItfMap["sensors.QAmbientLightFilter_ITF"] = QAmbientLightFilter{}
	qt.ItfMap["sensors.QAmbientLightReading_ITF"] = QAmbientLightReading{}
	qt.EnumMap["sensors.QAmbientLightReading__Undefined"] = int64(QAmbientLightReading__Undefined)
	qt.EnumMap["sensors.QAmbientLightReading__Dark"] = int64(QAmbientLightReading__Dark)
	qt.EnumMap["sensors.QAmbientLightReading__Twilight"] = int64(QAmbientLightReading__Twilight)
	qt.EnumMap["sensors.QAmbientLightReading__Light"] = int64(QAmbientLightReading__Light)
	qt.EnumMap["sensors.QAmbientLightReading__Bright"] = int64(QAmbientLightReading__Bright)
	qt.EnumMap["sensors.QAmbientLightReading__Sunny"] = int64(QAmbientLightReading__Sunny)
	qt.ItfMap["sensors.QAmbientLightSensor_ITF"] = QAmbientLightSensor{}
	qt.FuncMap["sensors.NewQAmbientLightSensor"] = NewQAmbientLightSensor
	qt.ItfMap["sensors.QAmbientTemperatureFilter_ITF"] = QAmbientTemperatureFilter{}
	qt.ItfMap["sensors.QAmbientTemperatureReading_ITF"] = QAmbientTemperatureReading{}
	qt.ItfMap["sensors.QAmbientTemperatureSensor_ITF"] = QAmbientTemperatureSensor{}
	qt.FuncMap["sensors.NewQAmbientTemperatureSensor"] = NewQAmbientTemperatureSensor
	qt.ItfMap["sensors.QCompass_ITF"] = QCompass{}
	qt.FuncMap["sensors.NewQCompass"] = NewQCompass
	qt.ItfMap["sensors.QCompassFilter_ITF"] = QCompassFilter{}
	qt.ItfMap["sensors.QCompassReading_ITF"] = QCompassReading{}
	qt.ItfMap["sensors.QDistanceFilter_ITF"] = QDistanceFilter{}
	qt.ItfMap["sensors.QDistanceReading_ITF"] = QDistanceReading{}
	qt.ItfMap["sensors.QDistanceSensor_ITF"] = QDistanceSensor{}
	qt.FuncMap["sensors.NewQDistanceSensor"] = NewQDistanceSensor
	qt.ItfMap["sensors.QGyroscope_ITF"] = QGyroscope{}
	qt.FuncMap["sensors.NewQGyroscope"] = NewQGyroscope
	qt.ItfMap["sensors.QGyroscopeFilter_ITF"] = QGyroscopeFilter{}
	qt.ItfMap["sensors.QGyroscopeReading_ITF"] = QGyroscopeReading{}
	qt.ItfMap["sensors.QHolsterFilter_ITF"] = QHolsterFilter{}
	qt.ItfMap["sensors.QHolsterReading_ITF"] = QHolsterReading{}
	qt.ItfMap["sensors.QHolsterSensor_ITF"] = QHolsterSensor{}
	qt.FuncMap["sensors.NewQHolsterSensor"] = NewQHolsterSensor
	qt.ItfMap["sensors.QHumidityFilter_ITF"] = QHumidityFilter{}
	qt.ItfMap["sensors.QHumidityReading_ITF"] = QHumidityReading{}
	qt.ItfMap["sensors.QHumiditySensor_ITF"] = QHumiditySensor{}
	qt.FuncMap["sensors.NewQHumiditySensor"] = NewQHumiditySensor
	qt.ItfMap["sensors.QIRProximityFilter_ITF"] = QIRProximityFilter{}
	qt.ItfMap["sensors.QIRProximityReading_ITF"] = QIRProximityReading{}
	qt.ItfMap["sensors.QIRProximitySensor_ITF"] = QIRProximitySensor{}
	qt.FuncMap["sensors.NewQIRProximitySensor"] = NewQIRProximitySensor
	qt.ItfMap["sensors.QLidFilter_ITF"] = QLidFilter{}
	qt.ItfMap["sensors.QLidReading_ITF"] = QLidReading{}
	qt.ItfMap["sensors.QLidSensor_ITF"] = QLidSensor{}
	qt.FuncMap["sensors.NewQLidSensor"] = NewQLidSensor
	qt.ItfMap["sensors.QLightFilter_ITF"] = QLightFilter{}
	qt.ItfMap["sensors.QLightReading_ITF"] = QLightReading{}
	qt.ItfMap["sensors.QLightSensor_ITF"] = QLightSensor{}
	qt.FuncMap["sensors.NewQLightSensor"] = NewQLightSensor
	qt.ItfMap["sensors.QMagnetometer_ITF"] = QMagnetometer{}
	qt.FuncMap["sensors.NewQMagnetometer"] = NewQMagnetometer
	qt.ItfMap["sensors.QMagnetometerFilter_ITF"] = QMagnetometerFilter{}
	qt.ItfMap["sensors.QMagnetometerReading_ITF"] = QMagnetometerReading{}
	qt.ItfMap["sensors.QOrientationFilter_ITF"] = QOrientationFilter{}
	qt.ItfMap["sensors.QOrientationReading_ITF"] = QOrientationReading{}
	qt.EnumMap["sensors.QOrientationReading__Undefined"] = int64(QOrientationReading__Undefined)
	qt.EnumMap["sensors.QOrientationReading__TopUp"] = int64(QOrientationReading__TopUp)
	qt.EnumMap["sensors.QOrientationReading__TopDown"] = int64(QOrientationReading__TopDown)
	qt.EnumMap["sensors.QOrientationReading__LeftUp"] = int64(QOrientationReading__LeftUp)
	qt.EnumMap["sensors.QOrientationReading__RightUp"] = int64(QOrientationReading__RightUp)
	qt.EnumMap["sensors.QOrientationReading__FaceUp"] = int64(QOrientationReading__FaceUp)
	qt.EnumMap["sensors.QOrientationReading__FaceDown"] = int64(QOrientationReading__FaceDown)
	qt.ItfMap["sensors.QOrientationSensor_ITF"] = QOrientationSensor{}
	qt.FuncMap["sensors.NewQOrientationSensor"] = NewQOrientationSensor
	qt.ItfMap["sensors.QPressureFilter_ITF"] = QPressureFilter{}
	qt.ItfMap["sensors.QPressureReading_ITF"] = QPressureReading{}
	qt.ItfMap["sensors.QPressureSensor_ITF"] = QPressureSensor{}
	qt.FuncMap["sensors.NewQPressureSensor"] = NewQPressureSensor
	qt.ItfMap["sensors.QProximityFilter_ITF"] = QProximityFilter{}
	qt.ItfMap["sensors.QProximityReading_ITF"] = QProximityReading{}
	qt.ItfMap["sensors.QProximitySensor_ITF"] = QProximitySensor{}
	qt.FuncMap["sensors.NewQProximitySensor"] = NewQProximitySensor
	qt.ItfMap["sensors.QRotationFilter_ITF"] = QRotationFilter{}
	qt.ItfMap["sensors.QRotationReading_ITF"] = QRotationReading{}
	qt.ItfMap["sensors.QRotationSensor_ITF"] = QRotationSensor{}
	qt.FuncMap["sensors.NewQRotationSensor"] = NewQRotationSensor
	qt.ItfMap["sensors.QSensor_ITF"] = QSensor{}
	qt.FuncMap["sensors.NewQSensor"] = NewQSensor
	qt.FuncMap["sensors.QSensor_DefaultSensorForType"] = QSensor_DefaultSensorForType
	qt.FuncMap["sensors.QSensor_SensorTypes"] = QSensor_SensorTypes
	qt.FuncMap["sensors.QSensor_SensorsForType"] = QSensor_SensorsForType
	qt.EnumMap["sensors.QSensor__FixedOrientation"] = int64(QSensor__FixedOrientation)
	qt.EnumMap["sensors.QSensor__AutomaticOrientation"] = int64(QSensor__AutomaticOrientation)
	qt.EnumMap["sensors.QSensor__UserOrientation"] = int64(QSensor__UserOrientation)
	qt.EnumMap["sensors.QSensor__Buffering"] = int64(QSensor__Buffering)
	qt.EnumMap["sensors.QSensor__AlwaysOn"] = int64(QSensor__AlwaysOn)
	qt.EnumMap["sensors.QSensor__GeoValues"] = int64(QSensor__GeoValues)
	qt.EnumMap["sensors.QSensor__FieldOfView"] = int64(QSensor__FieldOfView)
	qt.EnumMap["sensors.QSensor__AccelerationMode"] = int64(QSensor__AccelerationMode)
	qt.EnumMap["sensors.QSensor__SkipDuplicates"] = int64(QSensor__SkipDuplicates)
	qt.EnumMap["sensors.QSensor__AxesOrientation"] = int64(QSensor__AxesOrientation)
	qt.EnumMap["sensors.QSensor__PressureSensorTemperature"] = int64(QSensor__PressureSensorTemperature)
	qt.EnumMap["sensors.QSensor__Reserved"] = int64(QSensor__Reserved)
	qt.ItfMap["sensors.QSensorBackend_ITF"] = QSensorBackend{}
	qt.ItfMap["sensors.QSensorBackendFactory_ITF"] = QSensorBackendFactory{}
	qt.ItfMap["sensors.QSensorChangesInterface_ITF"] = QSensorChangesInterface{}
	qt.ItfMap["sensors.QSensorFilter_ITF"] = QSensorFilter{}
	qt.ItfMap["sensors.QSensorGesture_ITF"] = QSensorGesture{}
	qt.FuncMap["sensors.NewQSensorGesture"] = NewQSensorGesture
	qt.ItfMap["sensors.QSensorGestureManager_ITF"] = QSensorGestureManager{}
	qt.FuncMap["sensors.NewQSensorGestureManager"] = NewQSensorGestureManager
	qt.FuncMap["sensors.QSensorGestureManager_SensorGestureRecognizer"] = QSensorGestureManager_SensorGestureRecognizer
	qt.ItfMap["sensors.QSensorGesturePluginInterface_ITF"] = QSensorGesturePluginInterface{}
	qt.FuncMap["sensors.NewQSensorGesturePluginInterface"] = NewQSensorGesturePluginInterface
	qt.ItfMap["sensors.QSensorGestureRecognizer_ITF"] = QSensorGestureRecognizer{}
	qt.FuncMap["sensors.NewQSensorGestureRecognizer"] = NewQSensorGestureRecognizer
	qt.ItfMap["sensors.QSensorManager_ITF"] = QSensorManager{}
	qt.FuncMap["sensors.QSensorManager_CreateBackend"] = QSensorManager_CreateBackend
	qt.FuncMap["sensors.QSensorManager_IsBackendRegistered"] = QSensorManager_IsBackendRegistered
	qt.FuncMap["sensors.QSensorManager_RegisterBackend"] = QSensorManager_RegisterBackend
	qt.FuncMap["sensors.QSensorManager_SetDefaultBackend"] = QSensorManager_SetDefaultBackend
	qt.FuncMap["sensors.QSensorManager_UnregisterBackend"] = QSensorManager_UnregisterBackend
	qt.ItfMap["sensors.QSensorPluginInterface_ITF"] = QSensorPluginInterface{}
	qt.ItfMap["sensors.QSensorReading_ITF"] = QSensorReading{}
	qt.ItfMap["sensors.QTapFilter_ITF"] = QTapFilter{}
	qt.ItfMap["sensors.QTapReading_ITF"] = QTapReading{}
	qt.EnumMap["sensors.QTapReading__Undefined"] = int64(QTapReading__Undefined)
	qt.EnumMap["sensors.QTapReading__X"] = int64(QTapReading__X)
	qt.EnumMap["sensors.QTapReading__Y"] = int64(QTapReading__Y)
	qt.EnumMap["sensors.QTapReading__Z"] = int64(QTapReading__Z)
	qt.EnumMap["sensors.QTapReading__X_Pos"] = int64(QTapReading__X_Pos)
	qt.EnumMap["sensors.QTapReading__Y_Pos"] = int64(QTapReading__Y_Pos)
	qt.EnumMap["sensors.QTapReading__Z_Pos"] = int64(QTapReading__Z_Pos)
	qt.EnumMap["sensors.QTapReading__X_Neg"] = int64(QTapReading__X_Neg)
	qt.EnumMap["sensors.QTapReading__Y_Neg"] = int64(QTapReading__Y_Neg)
	qt.EnumMap["sensors.QTapReading__Z_Neg"] = int64(QTapReading__Z_Neg)
	qt.EnumMap["sensors.QTapReading__X_Both"] = int64(QTapReading__X_Both)
	qt.EnumMap["sensors.QTapReading__Y_Both"] = int64(QTapReading__Y_Both)
	qt.EnumMap["sensors.QTapReading__Z_Both"] = int64(QTapReading__Z_Both)
	qt.ItfMap["sensors.QTapSensor_ITF"] = QTapSensor{}
	qt.FuncMap["sensors.NewQTapSensor"] = NewQTapSensor
	qt.ItfMap["sensors.QTiltFilter_ITF"] = QTiltFilter{}
	qt.ItfMap["sensors.QTiltReading_ITF"] = QTiltReading{}
	qt.ItfMap["sensors.QTiltSensor_ITF"] = QTiltSensor{}
	qt.FuncMap["sensors.NewQTiltSensor"] = NewQTiltSensor
}
