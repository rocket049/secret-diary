// +build !minimal

#define protected public
#define private public

#include "serialport.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QIODevice>
#include <QLayout>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QRemoteObjectPendingCallWatcher>
#include <QSerialPort>
#include <QSerialPortInfo>
#include <QString>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

class MyQSerialPort: public QSerialPort
{
public:
	MyQSerialPort(QObject *parent = Q_NULLPTR) : QSerialPort(parent) {QSerialPort_QSerialPort_QRegisterMetaType();};
	MyQSerialPort(const QString &name, QObject *parent = Q_NULLPTR) : QSerialPort(name, parent) {QSerialPort_QSerialPort_QRegisterMetaType();};
	MyQSerialPort(const QSerialPortInfo &serialPortInfo, QObject *parent = Q_NULLPTR) : QSerialPort(serialPortInfo, parent) {QSerialPort_QSerialPort_QRegisterMetaType();};
	bool atEnd() const { return callbackQSerialPort_AtEnd(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	void Signal_BaudRateChanged(qint32 baudRate, QSerialPort::Directions directions) { callbackQSerialPort_BaudRateChanged(this, baudRate, directions); };
	void Signal_BreakEnabledChanged(bool set) { callbackQSerialPort_BreakEnabledChanged(this, set); };
	qint64 bytesAvailable() const { return callbackQSerialPort_BytesAvailable(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 bytesToWrite() const { return callbackQSerialPort_BytesToWrite(const_cast<void*>(static_cast<const void*>(this))); };
	bool canReadLine() const { return callbackQSerialPort_CanReadLine(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	void close() { callbackQSerialPort_Close(this); };
	void Signal_DataBitsChanged(QSerialPort::DataBits dataBits) { callbackQSerialPort_DataBitsChanged(this, dataBits); };
	void Signal_DataTerminalReadyChanged(bool set) { callbackQSerialPort_DataTerminalReadyChanged(this, set); };
	void Signal_ErrorOccurred(QSerialPort::SerialPortError error) { callbackQSerialPort_ErrorOccurred(this, error); };
	void Signal_FlowControlChanged(QSerialPort::FlowControl flow) { callbackQSerialPort_FlowControlChanged(this, flow); };
	bool isSequential() const { return callbackQSerialPort_IsSequential(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	bool open(QIODevice::OpenMode mode) { return callbackQSerialPort_Open(this, mode) != 0; };
	void Signal_ParityChanged(QSerialPort::Parity parity) { callbackQSerialPort_ParityChanged(this, parity); };
	qint64 readData(char * data, qint64 maxSize) { QtSerialPort_PackedString dataPacked = { data, maxSize, NULL };return callbackQSerialPort_ReadData(this, dataPacked, maxSize); };
	qint64 readLineData(char * data, qint64 maxSize) { QtSerialPort_PackedString dataPacked = { data, maxSize, NULL };return callbackQSerialPort_ReadLineData(this, dataPacked, maxSize); };
	void Signal_RequestToSendChanged(bool set) { callbackQSerialPort_RequestToSendChanged(this, set); };
	void Signal_StopBitsChanged(QSerialPort::StopBits stopBits) { callbackQSerialPort_StopBitsChanged(this, stopBits); };
	bool waitForBytesWritten(int msecs) { return callbackQSerialPort_WaitForBytesWritten(this, msecs) != 0; };
	bool waitForReadyRead(int msecs) { return callbackQSerialPort_WaitForReadyRead(this, msecs) != 0; };
	qint64 writeData(const char * data, qint64 maxSize) { QtSerialPort_PackedString dataPacked = { const_cast<char*>(data), maxSize, NULL };return callbackQSerialPort_WriteData(this, dataPacked, maxSize); };
	 ~MyQSerialPort() { callbackQSerialPort_DestroyQSerialPort(this); };
	void Signal_AboutToClose() { callbackQSerialPort_AboutToClose(this); };
	void Signal_BytesWritten(qint64 bytes) { callbackQSerialPort_BytesWritten(this, bytes); };
	void Signal_ChannelBytesWritten(int channel, qint64 bytes) { callbackQSerialPort_ChannelBytesWritten(this, channel, bytes); };
	void Signal_ChannelReadyRead(int channel) { callbackQSerialPort_ChannelReadyRead(this, channel); };
	qint64 pos() const { return callbackQSerialPort_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	void Signal_ReadChannelFinished() { callbackQSerialPort_ReadChannelFinished(this); };
	void Signal_ReadyRead() { callbackQSerialPort_ReadyRead(this); };
	bool reset() { return callbackQSerialPort_Reset(this) != 0; };
	bool seek(qint64 pos) { return callbackQSerialPort_Seek(this, pos) != 0; };
	qint64 size() const { return callbackQSerialPort_Size(const_cast<void*>(static_cast<const void*>(this))); };
	void childEvent(QChildEvent * event) { callbackQSerialPort_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSerialPort_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSerialPort_CustomEvent(this, event); };
	void deleteLater() { callbackQSerialPort_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSerialPort_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSerialPort_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSerialPort_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSerialPort_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSerialPort_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtSerialPort_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQSerialPort_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSerialPort_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QSerialPort*)
Q_DECLARE_METATYPE(MyQSerialPort*)

int QSerialPort_QSerialPort_QRegisterMetaType(){qRegisterMetaType<QSerialPort*>(); return qRegisterMetaType<MyQSerialPort*>();}

void* QSerialPort_NewQSerialPort(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(static_cast<QWindow*>(parent));
	} else {
		return new MyQSerialPort(static_cast<QObject*>(parent));
	}
}

void* QSerialPort_NewQSerialPort2(struct QtSerialPort_PackedString name, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQSerialPort(QString::fromUtf8(name.data, name.len), static_cast<QObject*>(parent));
	}
}

void* QSerialPort_NewQSerialPort3(void* serialPortInfo, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QWindow*>(parent));
	} else {
		return new MyQSerialPort(*static_cast<QSerialPortInfo*>(serialPortInfo), static_cast<QObject*>(parent));
	}
}

char QSerialPort_AtEndDefault(void* ptr)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::atEnd();
}

int QSerialPort_BaudRate(void* ptr, long long directions)
{
	return static_cast<QSerialPort*>(ptr)->baudRate(static_cast<QSerialPort::Direction>(directions));
}

void QSerialPort_ConnectBaudRateChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(qint32, QSerialPort::Directions)>(&QSerialPort::baudRateChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(qint32, QSerialPort::Directions)>(&MyQSerialPort::Signal_BaudRateChanged), static_cast<Qt::ConnectionType>(t));
}

void QSerialPort_DisconnectBaudRateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(qint32, QSerialPort::Directions)>(&QSerialPort::baudRateChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(qint32, QSerialPort::Directions)>(&MyQSerialPort::Signal_BaudRateChanged));
}

void QSerialPort_BaudRateChanged(void* ptr, int baudRate, long long directions)
{
	static_cast<QSerialPort*>(ptr)->baudRateChanged(baudRate, static_cast<QSerialPort::Direction>(directions));
}

void QSerialPort_ConnectBreakEnabledChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::breakEnabledChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_BreakEnabledChanged), static_cast<Qt::ConnectionType>(t));
}

void QSerialPort_DisconnectBreakEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::breakEnabledChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_BreakEnabledChanged));
}

void QSerialPort_BreakEnabledChanged(void* ptr, char set)
{
	static_cast<QSerialPort*>(ptr)->breakEnabledChanged(set != 0);
}

long long QSerialPort_BytesAvailableDefault(void* ptr)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::bytesAvailable();
}

long long QSerialPort_BytesToWriteDefault(void* ptr)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::bytesToWrite();
}

char QSerialPort_CanReadLineDefault(void* ptr)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::canReadLine();
}

char QSerialPort_Clear(void* ptr, long long directions)
{
	return static_cast<QSerialPort*>(ptr)->clear(static_cast<QSerialPort::Direction>(directions));
}

void QSerialPort_ClearError(void* ptr)
{
	static_cast<QSerialPort*>(ptr)->clearError();
}

void QSerialPort_CloseDefault(void* ptr)
{
		static_cast<QSerialPort*>(ptr)->QSerialPort::close();
}

long long QSerialPort_DataBits(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->dataBits();
}

void QSerialPort_ConnectDataBitsChanged(void* ptr, long long t)
{
	qRegisterMetaType<QSerialPort::DataBits>();
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::DataBits)>(&QSerialPort::dataBitsChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::DataBits)>(&MyQSerialPort::Signal_DataBitsChanged), static_cast<Qt::ConnectionType>(t));
}

void QSerialPort_DisconnectDataBitsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::DataBits)>(&QSerialPort::dataBitsChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::DataBits)>(&MyQSerialPort::Signal_DataBitsChanged));
}

void QSerialPort_DataBitsChanged(void* ptr, long long dataBits)
{
	static_cast<QSerialPort*>(ptr)->dataBitsChanged(static_cast<QSerialPort::DataBits>(dataBits));
}

void QSerialPort_ConnectDataTerminalReadyChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::dataTerminalReadyChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_DataTerminalReadyChanged), static_cast<Qt::ConnectionType>(t));
}

void QSerialPort_DisconnectDataTerminalReadyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::dataTerminalReadyChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_DataTerminalReadyChanged));
}

void QSerialPort_DataTerminalReadyChanged(void* ptr, char set)
{
	static_cast<QSerialPort*>(ptr)->dataTerminalReadyChanged(set != 0);
}

long long QSerialPort_Error(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->error();
}

void QSerialPort_ConnectErrorOccurred(void* ptr, long long t)
{
	qRegisterMetaType<QSerialPort::SerialPortError>();
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::SerialPortError)>(&QSerialPort::errorOccurred), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::SerialPortError)>(&MyQSerialPort::Signal_ErrorOccurred), static_cast<Qt::ConnectionType>(t));
}

void QSerialPort_DisconnectErrorOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::SerialPortError)>(&QSerialPort::errorOccurred), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::SerialPortError)>(&MyQSerialPort::Signal_ErrorOccurred));
}

void QSerialPort_ErrorOccurred(void* ptr, long long error)
{
	static_cast<QSerialPort*>(ptr)->errorOccurred(static_cast<QSerialPort::SerialPortError>(error));
}

long long QSerialPort_FlowControl(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->flowControl();
}

void QSerialPort_ConnectFlowControlChanged(void* ptr, long long t)
{
	qRegisterMetaType<QSerialPort::FlowControl>();
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::FlowControl)>(&QSerialPort::flowControlChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::FlowControl)>(&MyQSerialPort::Signal_FlowControlChanged), static_cast<Qt::ConnectionType>(t));
}

void QSerialPort_DisconnectFlowControlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::FlowControl)>(&QSerialPort::flowControlChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::FlowControl)>(&MyQSerialPort::Signal_FlowControlChanged));
}

void QSerialPort_FlowControlChanged(void* ptr, long long flow)
{
	static_cast<QSerialPort*>(ptr)->flowControlChanged(static_cast<QSerialPort::FlowControl>(flow));
}

char QSerialPort_Flush(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->flush();
}

char QSerialPort_IsBreakEnabled(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->isBreakEnabled();
}

char QSerialPort_IsDataTerminalReady(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->isDataTerminalReady();
}

char QSerialPort_IsRequestToSend(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->isRequestToSend();
}

char QSerialPort_IsSequentialDefault(void* ptr)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::isSequential();
}

char QSerialPort_OpenDefault(void* ptr, long long mode)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::open(static_cast<QIODevice::OpenModeFlag>(mode));
}

long long QSerialPort_Parity(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->parity();
}

void QSerialPort_ConnectParityChanged(void* ptr, long long t)
{
	qRegisterMetaType<QSerialPort::Parity>();
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::Parity)>(&QSerialPort::parityChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::Parity)>(&MyQSerialPort::Signal_ParityChanged), static_cast<Qt::ConnectionType>(t));
}

void QSerialPort_DisconnectParityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::Parity)>(&QSerialPort::parityChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::Parity)>(&MyQSerialPort::Signal_ParityChanged));
}

void QSerialPort_ParityChanged(void* ptr, long long parity)
{
	static_cast<QSerialPort*>(ptr)->parityChanged(static_cast<QSerialPort::Parity>(parity));
}

long long QSerialPort_PinoutSignals(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->pinoutSignals();
}

struct QtSerialPort_PackedString QSerialPort_PortName(void* ptr)
{
	return ({ QByteArray* t212a26 = new QByteArray(static_cast<QSerialPort*>(ptr)->portName().toUtf8()); QtSerialPort_PackedString { const_cast<char*>(t212a26->prepend("WHITESPACE").constData()+10), t212a26->size()-10, t212a26 }; });
}

long long QSerialPort_ReadBufferSize(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->readBufferSize();
}

long long QSerialPort_ReadData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QSerialPort*>(ptr)->readData(data, maxSize);
}

long long QSerialPort_ReadDataDefault(void* ptr, char* data, long long maxSize)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::readData(data, maxSize);
}

long long QSerialPort_ReadLineDataDefault(void* ptr, char* data, long long maxSize)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::readLineData(data, maxSize);
}

void QSerialPort_ConnectRequestToSendChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::requestToSendChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_RequestToSendChanged), static_cast<Qt::ConnectionType>(t));
}

void QSerialPort_DisconnectRequestToSendChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(bool)>(&QSerialPort::requestToSendChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(bool)>(&MyQSerialPort::Signal_RequestToSendChanged));
}

void QSerialPort_RequestToSendChanged(void* ptr, char set)
{
	static_cast<QSerialPort*>(ptr)->requestToSendChanged(set != 0);
}

char QSerialPort_SendBreak(void* ptr, int duration)
{
	return static_cast<QSerialPort*>(ptr)->sendBreak(duration);
}

char QSerialPort_SetBaudRate(void* ptr, int baudRate, long long directions)
{
	return static_cast<QSerialPort*>(ptr)->setBaudRate(baudRate, static_cast<QSerialPort::Direction>(directions));
}

char QSerialPort_SetBreakEnabled(void* ptr, char set)
{
	return static_cast<QSerialPort*>(ptr)->setBreakEnabled(set != 0);
}

char QSerialPort_SetDataBits(void* ptr, long long dataBits)
{
	return static_cast<QSerialPort*>(ptr)->setDataBits(static_cast<QSerialPort::DataBits>(dataBits));
}

char QSerialPort_SetDataTerminalReady(void* ptr, char set)
{
	return static_cast<QSerialPort*>(ptr)->setDataTerminalReady(set != 0);
}

char QSerialPort_SetFlowControl(void* ptr, long long flowControl)
{
	return static_cast<QSerialPort*>(ptr)->setFlowControl(static_cast<QSerialPort::FlowControl>(flowControl));
}

char QSerialPort_SetParity(void* ptr, long long parity)
{
	return static_cast<QSerialPort*>(ptr)->setParity(static_cast<QSerialPort::Parity>(parity));
}

void QSerialPort_SetPort(void* ptr, void* serialPortInfo)
{
	static_cast<QSerialPort*>(ptr)->setPort(*static_cast<QSerialPortInfo*>(serialPortInfo));
}

void QSerialPort_SetPortName(void* ptr, struct QtSerialPort_PackedString name)
{
	static_cast<QSerialPort*>(ptr)->setPortName(QString::fromUtf8(name.data, name.len));
}

void QSerialPort_SetReadBufferSize(void* ptr, long long size)
{
	static_cast<QSerialPort*>(ptr)->setReadBufferSize(size);
}

char QSerialPort_SetRequestToSend(void* ptr, char set)
{
	return static_cast<QSerialPort*>(ptr)->setRequestToSend(set != 0);
}

void QSerialPort_SetSettingsRestoredOnClose(void* ptr, char restore)
{
	static_cast<QSerialPort*>(ptr)->setSettingsRestoredOnClose(restore != 0);
}

char QSerialPort_SetStopBits(void* ptr, long long stopBits)
{
	return static_cast<QSerialPort*>(ptr)->setStopBits(static_cast<QSerialPort::StopBits>(stopBits));
}

char QSerialPort_SettingsRestoredOnClose(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->settingsRestoredOnClose();
}

long long QSerialPort_StopBits(void* ptr)
{
	return static_cast<QSerialPort*>(ptr)->stopBits();
}

void QSerialPort_ConnectStopBitsChanged(void* ptr, long long t)
{
	qRegisterMetaType<QSerialPort::StopBits>();
	QObject::connect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::StopBits)>(&QSerialPort::stopBitsChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::StopBits)>(&MyQSerialPort::Signal_StopBitsChanged), static_cast<Qt::ConnectionType>(t));
}

void QSerialPort_DisconnectStopBitsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSerialPort*>(ptr), static_cast<void (QSerialPort::*)(QSerialPort::StopBits)>(&QSerialPort::stopBitsChanged), static_cast<MyQSerialPort*>(ptr), static_cast<void (MyQSerialPort::*)(QSerialPort::StopBits)>(&MyQSerialPort::Signal_StopBitsChanged));
}

void QSerialPort_StopBitsChanged(void* ptr, long long stopBits)
{
	static_cast<QSerialPort*>(ptr)->stopBitsChanged(static_cast<QSerialPort::StopBits>(stopBits));
}

char QSerialPort_WaitForBytesWrittenDefault(void* ptr, int msecs)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::waitForBytesWritten(msecs);
}

char QSerialPort_WaitForReadyReadDefault(void* ptr, int msecs)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::waitForReadyRead(msecs);
}

long long QSerialPort_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QSerialPort*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

long long QSerialPort_WriteDataDefault(void* ptr, char* data, long long maxSize)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::writeData(const_cast<const char*>(data), maxSize);
}

void QSerialPort_DestroyQSerialPort(void* ptr)
{
	static_cast<QSerialPort*>(ptr)->~QSerialPort();
}

void QSerialPort_DestroyQSerialPortDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSerialPort___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSerialPort___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSerialPort___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QSerialPort___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSerialPort___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSerialPort___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QSerialPort___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSerialPort___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSerialPort___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QSerialPort___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSerialPort___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSerialPort___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

long long QSerialPort_PosDefault(void* ptr)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::pos();
}

char QSerialPort_ResetDefault(void* ptr)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::reset();
}

char QSerialPort_SeekDefault(void* ptr, long long pos)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::seek(pos);
}

long long QSerialPort_SizeDefault(void* ptr)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::size();
}

void QSerialPort_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSerialPort*>(ptr)->QSerialPort::childEvent(static_cast<QChildEvent*>(event));
}

void QSerialPort_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSerialPort*>(ptr)->QSerialPort::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSerialPort_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSerialPort*>(ptr)->QSerialPort::customEvent(static_cast<QEvent*>(event));
}

void QSerialPort_DeleteLaterDefault(void* ptr)
{
		static_cast<QSerialPort*>(ptr)->QSerialPort::deleteLater();
}

void QSerialPort_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSerialPort*>(ptr)->QSerialPort::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSerialPort_EventDefault(void* ptr, void* e)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::event(static_cast<QEvent*>(e));
}

char QSerialPort_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QSerialPort*>(ptr)->QSerialPort::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSerialPort_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSerialPort*>(ptr)->QSerialPort::metaObject());
}

void QSerialPort_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSerialPort*>(ptr)->QSerialPort::timerEvent(static_cast<QTimerEvent*>(event));
}

Q_DECLARE_METATYPE(QSerialPortInfo)
Q_DECLARE_METATYPE(QSerialPortInfo*)
void* QSerialPortInfo_NewQSerialPortInfo()
{
	return new QSerialPortInfo();
}

void* QSerialPortInfo_NewQSerialPortInfo2(void* port)
{
	return new QSerialPortInfo(*static_cast<QSerialPort*>(port));
}

void* QSerialPortInfo_NewQSerialPortInfo3(struct QtSerialPort_PackedString name)
{
	return new QSerialPortInfo(QString::fromUtf8(name.data, name.len));
}

void* QSerialPortInfo_NewQSerialPortInfo4(void* other)
{
	return new QSerialPortInfo(*static_cast<QSerialPortInfo*>(other));
}

struct QtSerialPort_PackedList QSerialPortInfo_QSerialPortInfo_AvailablePorts()
{
	return ({ QList<QSerialPortInfo>* tmpValue9efc0e = new QList<QSerialPortInfo>(QSerialPortInfo::availablePorts()); QtSerialPort_PackedList { tmpValue9efc0e, tmpValue9efc0e->size() }; });
}

struct QtSerialPort_PackedString QSerialPortInfo_Description(void* ptr)
{
	return ({ QByteArray* t2edb91 = new QByteArray(static_cast<QSerialPortInfo*>(ptr)->description().toUtf8()); QtSerialPort_PackedString { const_cast<char*>(t2edb91->prepend("WHITESPACE").constData()+10), t2edb91->size()-10, t2edb91 }; });
}

char QSerialPortInfo_HasProductIdentifier(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->hasProductIdentifier();
}

char QSerialPortInfo_HasVendorIdentifier(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->hasVendorIdentifier();
}

char QSerialPortInfo_IsNull(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->isNull();
}

struct QtSerialPort_PackedString QSerialPortInfo_Manufacturer(void* ptr)
{
	return ({ QByteArray* ta0f36d = new QByteArray(static_cast<QSerialPortInfo*>(ptr)->manufacturer().toUtf8()); QtSerialPort_PackedString { const_cast<char*>(ta0f36d->prepend("WHITESPACE").constData()+10), ta0f36d->size()-10, ta0f36d }; });
}

struct QtSerialPort_PackedString QSerialPortInfo_PortName(void* ptr)
{
	return ({ QByteArray* tea4d96 = new QByteArray(static_cast<QSerialPortInfo*>(ptr)->portName().toUtf8()); QtSerialPort_PackedString { const_cast<char*>(tea4d96->prepend("WHITESPACE").constData()+10), tea4d96->size()-10, tea4d96 }; });
}

unsigned short QSerialPortInfo_ProductIdentifier(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->productIdentifier();
}

struct QtSerialPort_PackedString QSerialPortInfo_SerialNumber(void* ptr)
{
	return ({ QByteArray* tb01b4d = new QByteArray(static_cast<QSerialPortInfo*>(ptr)->serialNumber().toUtf8()); QtSerialPort_PackedString { const_cast<char*>(tb01b4d->prepend("WHITESPACE").constData()+10), tb01b4d->size()-10, tb01b4d }; });
}

struct QtSerialPort_PackedList QSerialPortInfo_QSerialPortInfo_StandardBaudRates()
{
	return ({ QList<qint32>* tmpValue059d90 = new QList<qint32>(QSerialPortInfo::standardBaudRates()); QtSerialPort_PackedList { tmpValue059d90, tmpValue059d90->size() }; });
}

void QSerialPortInfo_Swap(void* ptr, void* other)
{
	static_cast<QSerialPortInfo*>(ptr)->swap(*static_cast<QSerialPortInfo*>(other));
}

struct QtSerialPort_PackedString QSerialPortInfo_SystemLocation(void* ptr)
{
	return ({ QByteArray* te47b34 = new QByteArray(static_cast<QSerialPortInfo*>(ptr)->systemLocation().toUtf8()); QtSerialPort_PackedString { const_cast<char*>(te47b34->prepend("WHITESPACE").constData()+10), te47b34->size()-10, te47b34 }; });
}

unsigned short QSerialPortInfo_VendorIdentifier(void* ptr)
{
	return static_cast<QSerialPortInfo*>(ptr)->vendorIdentifier();
}

void QSerialPortInfo_DestroyQSerialPortInfo(void* ptr)
{
	static_cast<QSerialPortInfo*>(ptr)->~QSerialPortInfo();
}

void* QSerialPortInfo___availablePorts_atList(void* ptr, int i)
{
	return new QSerialPortInfo(({QSerialPortInfo tmp = static_cast<QList<QSerialPortInfo>*>(ptr)->at(i); if (i == static_cast<QList<QSerialPortInfo>*>(ptr)->size()-1) { static_cast<QList<QSerialPortInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QSerialPortInfo___availablePorts_setList(void* ptr, void* i)
{
	static_cast<QList<QSerialPortInfo>*>(ptr)->append(*static_cast<QSerialPortInfo*>(i));
}

void* QSerialPortInfo___availablePorts_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSerialPortInfo>();
}

int QSerialPortInfo___standardBaudRates_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSerialPortInfo___standardBaudRates_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* QSerialPortInfo___standardBaudRates_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

