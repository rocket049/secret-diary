// +build !minimal

#define protected public
#define private public

#include "remoteobjects.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
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
#include <QItemSelectionModel>
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
#include <QRemoteObjectAbstractPersistedStore>
#include <QRemoteObjectDynamicReplica>
#include <QRemoteObjectHost>
#include <QRemoteObjectHostBase>
#include <QRemoteObjectNode>
#include <QRemoteObjectPendingCall>
#include <QRemoteObjectPendingCallWatcher>
#include <QRemoteObjectRegistry>
#include <QRemoteObjectRegistryHost>
#include <QRemoteObjectReplica>
#include <QString>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>

class MyQRemoteObjectAbstractPersistedStore: public QRemoteObjectAbstractPersistedStore
{
public:
	MyQRemoteObjectAbstractPersistedStore(QObject *parent = Q_NULLPTR) : QRemoteObjectAbstractPersistedStore(parent) {QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_QRegisterMetaType();};
	QList<QVariant> restoreProperties(const QString & repName, const QByteArray & repSig) { QByteArray* t7884ba = new QByteArray(repName.toUtf8()); QtRemoteObjects_PackedString repNamePacked = { const_cast<char*>(t7884ba->prepend("WHITESPACE").constData()+10), t7884ba->size()-10, t7884ba };return ({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(callbackQRemoteObjectAbstractPersistedStore_RestoreProperties(this, repNamePacked, const_cast<QByteArray*>(&repSig))); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void saveProperties(const QString & repName, const QByteArray & repSig, const QVariantList & values) { QByteArray* t7884ba = new QByteArray(repName.toUtf8()); QtRemoteObjects_PackedString repNamePacked = { const_cast<char*>(t7884ba->prepend("WHITESPACE").constData()+10), t7884ba->size()-10, t7884ba };callbackQRemoteObjectAbstractPersistedStore_SaveProperties(this, repNamePacked, const_cast<QByteArray*>(&repSig), ({ QList<QVariant>* tmpValue048b0c = new QList<QVariant>(values); QtRemoteObjects_PackedList { tmpValue048b0c, tmpValue048b0c->size() }; })); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectAbstractPersistedStore_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectAbstractPersistedStore_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectAbstractPersistedStore_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectAbstractPersistedStore_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectAbstractPersistedStore_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectAbstractPersistedStore_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectAbstractPersistedStore_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectAbstractPersistedStore_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectAbstractPersistedStore_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectAbstractPersistedStore_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectAbstractPersistedStore_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectAbstractPersistedStore*)
Q_DECLARE_METATYPE(MyQRemoteObjectAbstractPersistedStore*)

int QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectAbstractPersistedStore*>(); return qRegisterMetaType<MyQRemoteObjectAbstractPersistedStore*>();}

void* QRemoteObjectAbstractPersistedStore_NewQRemoteObjectAbstractPersistedStore(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QObject*>(parent));
	}
}

struct QtRemoteObjects_PackedList QRemoteObjectAbstractPersistedStore_RestoreProperties(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig)
{
	return ({ QList<QVariant>* tmpValue53915f = new QList<QVariant>(static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->restoreProperties(QString::fromUtf8(repName.data, repName.len), *static_cast<QByteArray*>(repSig))); QtRemoteObjects_PackedList { tmpValue53915f, tmpValue53915f->size() }; });
}

void QRemoteObjectAbstractPersistedStore_SaveProperties(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig, void* values)
{
	static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->saveProperties(QString::fromUtf8(repName.data, repName.len), *static_cast<QByteArray*>(repSig), ({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(values); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* QRemoteObjectAbstractPersistedStore___restoreProperties_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectAbstractPersistedStore___restoreProperties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectAbstractPersistedStore___restoreProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectAbstractPersistedStore___saveProperties_values_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectAbstractPersistedStore___saveProperties_values_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectAbstractPersistedStore___saveProperties_values_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectAbstractPersistedStore___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectAbstractPersistedStore___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectAbstractPersistedStore___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QRemoteObjectAbstractPersistedStore___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectAbstractPersistedStore___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectAbstractPersistedStore___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QRemoteObjectAbstractPersistedStore___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectAbstractPersistedStore___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectAbstractPersistedStore___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QRemoteObjectAbstractPersistedStore_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::childEvent(static_cast<QChildEvent*>(event));
}

void QRemoteObjectAbstractPersistedStore_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRemoteObjectAbstractPersistedStore_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::customEvent(static_cast<QEvent*>(event));
}

void QRemoteObjectAbstractPersistedStore_DeleteLaterDefault(void* ptr)
{
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::deleteLater();
}

void QRemoteObjectAbstractPersistedStore_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QRemoteObjectAbstractPersistedStore_EventDefault(void* ptr, void* e)
{
		return static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::event(static_cast<QEvent*>(e));
}

char QRemoteObjectAbstractPersistedStore_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void* QRemoteObjectAbstractPersistedStore_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::metaObject());
}

void QRemoteObjectAbstractPersistedStore_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQRemoteObjectDynamicReplica: public QRemoteObjectDynamicReplica
{
public:
	 ~MyQRemoteObjectDynamicReplica() { callbackQRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica(this); };
	void Signal_Initialized() { callbackQRemoteObjectReplica_Initialized(this); };
	void Signal_Notified() { callbackQRemoteObjectReplica_Notified(this); };
	void setNode(QRemoteObjectNode * node) { callbackQRemoteObjectReplica_SetNode(this, node); };
	void Signal_StateChanged(QRemoteObjectReplica::State state, QRemoteObjectReplica::State oldState) { callbackQRemoteObjectReplica_StateChanged(this, state, oldState); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectReplica_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectReplica_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectReplica_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectReplica_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectReplica_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectReplica_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectReplica_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectReplica_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectReplica_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectDynamicReplica*)
Q_DECLARE_METATYPE(MyQRemoteObjectDynamicReplica*)

int QRemoteObjectDynamicReplica_QRemoteObjectDynamicReplica_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectDynamicReplica*>(); return qRegisterMetaType<MyQRemoteObjectDynamicReplica*>();}

void QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica(void* ptr)
{
	static_cast<QRemoteObjectDynamicReplica*>(ptr)->~QRemoteObjectDynamicReplica();
}

void QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplicaDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQRemoteObjectHost: public QRemoteObjectHost
{
public:
	MyQRemoteObjectHost(QObject *parent = Q_NULLPTR) : QRemoteObjectHost(parent) {QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType();};
	MyQRemoteObjectHost(const QUrl &address, const QUrl &registryAddress = QUrl(), QRemoteObjectHostBase::AllowedSchemas allowedSchemas = BuiltInSchemasOnly, QObject *parent = Q_NULLPTR) : QRemoteObjectHost(address, registryAddress, allowedSchemas, parent) {QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType();};
	MyQRemoteObjectHost(const QUrl &address, QObject *parent) : QRemoteObjectHost(address, parent) {QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType();};
	QUrl hostUrl() const { return *static_cast<QUrl*>(callbackQRemoteObjectHost_HostUrl(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_HostUrlChanged() { callbackQRemoteObjectHost_HostUrlChanged(this); };
	bool setHostUrl(const QUrl & hostAddress, QRemoteObjectHostBase::AllowedSchemas allowedSchemas) { return callbackQRemoteObjectHost_SetHostUrl(this, const_cast<QUrl*>(&hostAddress), allowedSchemas) != 0; };
	void setName(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	void Signal_HeartbeatIntervalChanged(int heartbeatInterval) { callbackQRemoteObjectNode_HeartbeatIntervalChanged(this, heartbeatInterval); };
	bool setRegistryUrl(const QUrl & registryAddress) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryAddress)) != 0; };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QRemoteObjectHost*)
Q_DECLARE_METATYPE(MyQRemoteObjectHost*)

int QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectHost*>(); return qRegisterMetaType<MyQRemoteObjectHost*>();}

void* QRemoteObjectHost_NewQRemoteObjectHost(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectHost(static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectHost_NewQRemoteObjectHost2(void* address, void* registryAddress, long long allowedSchemas, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectHost_NewQRemoteObjectHost3(void* address, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectHost_HostUrl(void* ptr)
{
	return new QUrl(static_cast<QRemoteObjectHost*>(ptr)->hostUrl());
}

void* QRemoteObjectHost_HostUrlDefault(void* ptr)
{
		return new QUrl(static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::hostUrl());
}

void QRemoteObjectHost_ConnectHostUrlChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QRemoteObjectHost*>(ptr), static_cast<void (QRemoteObjectHost::*)()>(&QRemoteObjectHost::hostUrlChanged), static_cast<MyQRemoteObjectHost*>(ptr), static_cast<void (MyQRemoteObjectHost::*)()>(&MyQRemoteObjectHost::Signal_HostUrlChanged), static_cast<Qt::ConnectionType>(t));
}

void QRemoteObjectHost_DisconnectHostUrlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectHost*>(ptr), static_cast<void (QRemoteObjectHost::*)()>(&QRemoteObjectHost::hostUrlChanged), static_cast<MyQRemoteObjectHost*>(ptr), static_cast<void (MyQRemoteObjectHost::*)()>(&MyQRemoteObjectHost::Signal_HostUrlChanged));
}

void QRemoteObjectHost_HostUrlChanged(void* ptr)
{
	static_cast<QRemoteObjectHost*>(ptr)->hostUrlChanged();
}

char QRemoteObjectHost_SetHostUrl(void* ptr, void* hostAddress, long long allowedSchemas)
{
	return static_cast<QRemoteObjectHost*>(ptr)->setHostUrl(*static_cast<QUrl*>(hostAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas));
}

char QRemoteObjectHost_SetHostUrlDefault(void* ptr, void* hostAddress, long long allowedSchemas)
{
		return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::setHostUrl(*static_cast<QUrl*>(hostAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas));
}

class MyQRemoteObjectHostBase: public QRemoteObjectHostBase
{
public:
	void setName(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	void Signal_HeartbeatIntervalChanged(int heartbeatInterval) { callbackQRemoteObjectNode_HeartbeatIntervalChanged(this, heartbeatInterval); };
	bool setRegistryUrl(const QUrl & registryAddress) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryAddress)) != 0; };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QRemoteObjectHostBase*)
Q_DECLARE_METATYPE(MyQRemoteObjectHostBase*)

int QRemoteObjectHostBase_QRemoteObjectHostBase_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectHostBase*>(); return qRegisterMetaType<MyQRemoteObjectHostBase*>();}

void QRemoteObjectHostBase_AddHostSideConnection(void* ptr, void* ioDevice)
{
	static_cast<QRemoteObjectHostBase*>(ptr)->addHostSideConnection(static_cast<QIODevice*>(ioDevice));
}

char QRemoteObjectHostBase_DisableRemoting(void* ptr, void* remoteObject)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(remoteObject))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->disableRemoting(static_cast<QRemoteObjectPendingCallWatcher*>(remoteObject));
	} else {
		return static_cast<QRemoteObjectHostBase*>(ptr)->disableRemoting(static_cast<QObject*>(remoteObject));
	}
}

char QRemoteObjectHostBase_EnableRemoting2(void* ptr, void* object, struct QtRemoteObjects_PackedString name)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(object))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->enableRemoting(static_cast<QRemoteObjectPendingCallWatcher*>(object), QString::fromUtf8(name.data, name.len));
	} else {
		return static_cast<QRemoteObjectHostBase*>(ptr)->enableRemoting(static_cast<QObject*>(object), QString::fromUtf8(name.data, name.len));
	}
}

char QRemoteObjectHostBase_EnableRemoting3(void* ptr, void* model, struct QtRemoteObjects_PackedString name, void* roles, void* selectionModel)
{
	return static_cast<QRemoteObjectHostBase*>(ptr)->enableRemoting(static_cast<QAbstractItemModel*>(model), QString::fromUtf8(name.data, name.len), *static_cast<QVector<int>*>(roles), static_cast<QItemSelectionModel*>(selectionModel));
}

int QRemoteObjectHostBase___enableRemoting_roles_atList3(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QRemoteObjectHostBase___enableRemoting_roles_setList3(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QRemoteObjectHostBase___enableRemoting_roles_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

class MyQRemoteObjectNode: public QRemoteObjectNode
{
public:
	MyQRemoteObjectNode(QObject *parent = Q_NULLPTR) : QRemoteObjectNode(parent) {QRemoteObjectNode_QRemoteObjectNode_QRegisterMetaType();};
	MyQRemoteObjectNode(const QUrl &registryAddress, QObject *parent = Q_NULLPTR) : QRemoteObjectNode(registryAddress, parent) {QRemoteObjectNode_QRemoteObjectNode_QRegisterMetaType();};
	void Signal_HeartbeatIntervalChanged(int heartbeatInterval) { callbackQRemoteObjectNode_HeartbeatIntervalChanged(this, heartbeatInterval); };
	void setName(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	bool setRegistryUrl(const QUrl & registryAddress) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryAddress)) != 0; };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QRemoteObjectNode*)
Q_DECLARE_METATYPE(MyQRemoteObjectNode*)

int QRemoteObjectNode_QRemoteObjectNode_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectNode*>(); return qRegisterMetaType<MyQRemoteObjectNode*>();}

void* QRemoteObjectNode_NewQRemoteObjectNode(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectNode(static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectNode_NewQRemoteObjectNode2(void* registryAddress, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectNode_AcquireDynamic(void* ptr, struct QtRemoteObjects_PackedString name)
{
	return static_cast<QRemoteObjectNode*>(ptr)->acquireDynamic(QString::fromUtf8(name.data, name.len));
}

void QRemoteObjectNode_AddClientSideConnection(void* ptr, void* ioDevice)
{
	static_cast<QRemoteObjectNode*>(ptr)->addClientSideConnection(static_cast<QIODevice*>(ioDevice));
}

char QRemoteObjectNode_ConnectToNode(void* ptr, void* address)
{
	return static_cast<QRemoteObjectNode*>(ptr)->connectToNode(*static_cast<QUrl*>(address));
}

int QRemoteObjectNode_HeartbeatInterval(void* ptr)
{
	return static_cast<QRemoteObjectNode*>(ptr)->heartbeatInterval();
}

void QRemoteObjectNode_ConnectHeartbeatIntervalChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QRemoteObjectNode*>(ptr), static_cast<void (QRemoteObjectNode::*)(int)>(&QRemoteObjectNode::heartbeatIntervalChanged), static_cast<MyQRemoteObjectNode*>(ptr), static_cast<void (MyQRemoteObjectNode::*)(int)>(&MyQRemoteObjectNode::Signal_HeartbeatIntervalChanged), static_cast<Qt::ConnectionType>(t));
}

void QRemoteObjectNode_DisconnectHeartbeatIntervalChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectNode*>(ptr), static_cast<void (QRemoteObjectNode::*)(int)>(&QRemoteObjectNode::heartbeatIntervalChanged), static_cast<MyQRemoteObjectNode*>(ptr), static_cast<void (MyQRemoteObjectNode::*)(int)>(&MyQRemoteObjectNode::Signal_HeartbeatIntervalChanged));
}

void QRemoteObjectNode_HeartbeatIntervalChanged(void* ptr, int heartbeatInterval)
{
	static_cast<QRemoteObjectNode*>(ptr)->heartbeatIntervalChanged(heartbeatInterval);
}

struct QtRemoteObjects_PackedString QRemoteObjectNode_Instances2(void* ptr, struct QtRemoteObjects_PackedString typeName)
{
	return ({ QByteArray* t1bf5b2 = new QByteArray(static_cast<QRemoteObjectNode*>(ptr)->instances(QString::fromUtf8(typeName.data, typeName.len)).join("¡¦!").toUtf8()); QtRemoteObjects_PackedString { const_cast<char*>(t1bf5b2->prepend("WHITESPACE").constData()+10), t1bf5b2->size()-10, t1bf5b2 }; });
}

long long QRemoteObjectNode_LastError(void* ptr)
{
	return static_cast<QRemoteObjectNode*>(ptr)->lastError();
}

void* QRemoteObjectNode_PersistedStore(void* ptr)
{
	return static_cast<QRemoteObjectNode*>(ptr)->persistedStore();
}

void* QRemoteObjectNode_Registry(void* ptr)
{
	return const_cast<QRemoteObjectRegistry*>(static_cast<QRemoteObjectNode*>(ptr)->registry());
}

void* QRemoteObjectNode_RegistryUrl(void* ptr)
{
	return new QUrl(static_cast<QRemoteObjectNode*>(ptr)->registryUrl());
}

void QRemoteObjectNode_SetHeartbeatInterval(void* ptr, int interval)
{
	static_cast<QRemoteObjectNode*>(ptr)->setHeartbeatInterval(interval);
}

void QRemoteObjectNode_SetName(void* ptr, struct QtRemoteObjects_PackedString name)
{
	static_cast<QRemoteObjectNode*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QRemoteObjectNode_SetNameDefault(void* ptr, struct QtRemoteObjects_PackedString name)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::setName(QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::setName(QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::setName(QString::fromUtf8(name.data, name.len));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::setName(QString::fromUtf8(name.data, name.len));
	}
}

void QRemoteObjectNode_SetPersistedStore(void* ptr, void* persistedStore)
{
	static_cast<QRemoteObjectNode*>(ptr)->setPersistedStore(static_cast<QRemoteObjectAbstractPersistedStore*>(persistedStore));
}

char QRemoteObjectNode_SetRegistryUrl(void* ptr, void* registryAddress)
{
	return static_cast<QRemoteObjectNode*>(ptr)->setRegistryUrl(*static_cast<QUrl*>(registryAddress));
}

char QRemoteObjectNode_SetRegistryUrlDefault(void* ptr, void* registryAddress)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	} else {
		return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	}
}

void QRemoteObjectNode_TimerEventDefault(void* ptr, void* vqt)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::timerEvent(static_cast<QTimerEvent*>(vqt));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::timerEvent(static_cast<QTimerEvent*>(vqt));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::timerEvent(static_cast<QTimerEvent*>(vqt));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::timerEvent(static_cast<QTimerEvent*>(vqt));
	}
}

char QRemoteObjectNode_WaitForRegistry(void* ptr, int timeout)
{
	return static_cast<QRemoteObjectNode*>(ptr)->waitForRegistry(timeout);
}

int QRemoteObjectNode___acquireModel_rolesHint_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QRemoteObjectNode___acquireModel_rolesHint_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QRemoteObjectNode___acquireModel_rolesHint_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QRemoteObjectNode___persistProperties_props_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectNode___persistProperties_props_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectNode___persistProperties_props_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectNode___retrieveProperties_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectNode___retrieveProperties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectNode___retrieveProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectNode___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectNode___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectNode___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QRemoteObjectNode___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectNode___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRemoteObjectNode___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QRemoteObjectNode___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectNode___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectNode___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QRemoteObjectNode___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectNode___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectNode___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QRemoteObjectNode_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QRemoteObjectNode_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QRemoteObjectNode_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::customEvent(static_cast<QEvent*>(event));
	}
}

void QRemoteObjectNode_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::deleteLater();
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::deleteLater();
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::deleteLater();
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::deleteLater();
	}
}

void QRemoteObjectNode_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QRemoteObjectNode_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::event(static_cast<QEvent*>(e));
	}
}

char QRemoteObjectNode_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void* QRemoteObjectNode_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::metaObject());
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::metaObject());
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::metaObject());
	}
}

long long QRemoteObjectPendingCall_Error(void* ptr)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->error();
	} else {
		return static_cast<QRemoteObjectPendingCall*>(ptr)->error();
	}
}

char QRemoteObjectPendingCall_IsFinished(void* ptr)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->isFinished();
	} else {
		return static_cast<QRemoteObjectPendingCall*>(ptr)->isFinished();
	}
}

void* QRemoteObjectPendingCall_ReturnValue(void* ptr)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->returnValue());
	} else {
		return new QVariant(static_cast<QRemoteObjectPendingCall*>(ptr)->returnValue());
	}
}

class MyQRemoteObjectPendingCallWatcher: public QRemoteObjectPendingCallWatcher
{
public:
	void Signal_Finished(QRemoteObjectPendingCallWatcher * se) { callbackQRemoteObjectPendingCallWatcher_Finished(this, se); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectPendingCallWatcher_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectPendingCallWatcher_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectPendingCallWatcher_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectPendingCallWatcher_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectPendingCallWatcher_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectPendingCallWatcher_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectPendingCallWatcher_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectPendingCallWatcher_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectPendingCallWatcher_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectPendingCallWatcher_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectPendingCallWatcher_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectPendingCallWatcher*)
Q_DECLARE_METATYPE(MyQRemoteObjectPendingCallWatcher*)

int QRemoteObjectPendingCallWatcher_QRemoteObjectPendingCallWatcher_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectPendingCallWatcher*>(); return qRegisterMetaType<MyQRemoteObjectPendingCallWatcher*>();}

void QRemoteObjectPendingCallWatcher_ConnectFinished(void* ptr, long long t)
{
	QObject::connect(static_cast<QRemoteObjectPendingCallWatcher*>(ptr), static_cast<void (QRemoteObjectPendingCallWatcher::*)(QRemoteObjectPendingCallWatcher *)>(&QRemoteObjectPendingCallWatcher::finished), static_cast<MyQRemoteObjectPendingCallWatcher*>(ptr), static_cast<void (MyQRemoteObjectPendingCallWatcher::*)(QRemoteObjectPendingCallWatcher *)>(&MyQRemoteObjectPendingCallWatcher::Signal_Finished), static_cast<Qt::ConnectionType>(t));
}

void QRemoteObjectPendingCallWatcher_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectPendingCallWatcher*>(ptr), static_cast<void (QRemoteObjectPendingCallWatcher::*)(QRemoteObjectPendingCallWatcher *)>(&QRemoteObjectPendingCallWatcher::finished), static_cast<MyQRemoteObjectPendingCallWatcher*>(ptr), static_cast<void (MyQRemoteObjectPendingCallWatcher::*)(QRemoteObjectPendingCallWatcher *)>(&MyQRemoteObjectPendingCallWatcher::Signal_Finished));
}

void QRemoteObjectPendingCallWatcher_Finished(void* ptr, void* se)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->finished(static_cast<QRemoteObjectPendingCallWatcher*>(se));
}

void QRemoteObjectPendingCallWatcher_WaitForFinished(void* ptr)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->waitForFinished();
}

void* QRemoteObjectPendingCallWatcher___children_atList(void* ptr, int i)
{
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectPendingCallWatcher___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectPendingCallWatcher___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>();
}

void* QRemoteObjectPendingCallWatcher___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectPendingCallWatcher___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRemoteObjectPendingCallWatcher___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>();
}

void* QRemoteObjectPendingCallWatcher___findChildren_atList(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectPendingCallWatcher___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectPendingCallWatcher___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QRemoteObjectPendingCallWatcher___findChildren_atList3(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectPendingCallWatcher___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectPendingCallWatcher___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void QRemoteObjectPendingCallWatcher_ChildEvent(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRemoteObjectPendingCallWatcher_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::childEvent(static_cast<QChildEvent*>(event));
}

void QRemoteObjectPendingCallWatcher_ConnectNotify(void* ptr, void* sign)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRemoteObjectPendingCallWatcher_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRemoteObjectPendingCallWatcher_CustomEvent(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRemoteObjectPendingCallWatcher_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::customEvent(static_cast<QEvent*>(event));
}

void QRemoteObjectPendingCallWatcher_DeleteLater(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QRemoteObjectPendingCallWatcher*>(ptr), "deleteLater");
}

void QRemoteObjectPendingCallWatcher_DeleteLaterDefault(void* ptr)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::deleteLater();
}

void QRemoteObjectPendingCallWatcher_DisconnectNotify(void* ptr, void* sign)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRemoteObjectPendingCallWatcher_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QRemoteObjectPendingCallWatcher_Event(void* ptr, void* e)
{
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->event(static_cast<QEvent*>(e));
}

char QRemoteObjectPendingCallWatcher_EventDefault(void* ptr, void* e)
{
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::event(static_cast<QEvent*>(e));
}

char QRemoteObjectPendingCallWatcher_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QRemoteObjectPendingCallWatcher_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void* QRemoteObjectPendingCallWatcher_MetaObject(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->metaObject());
}

void* QRemoteObjectPendingCallWatcher_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::metaObject());
}

void QRemoteObjectPendingCallWatcher_TimerEvent(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRemoteObjectPendingCallWatcher_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQRemoteObjectRegistry: public QRemoteObjectRegistry
{
public:
	 ~MyQRemoteObjectRegistry() { callbackQRemoteObjectRegistry_DestroyQRemoteObjectRegistry(this); };
	void Signal_Initialized() { callbackQRemoteObjectReplica_Initialized(this); };
	void Signal_Notified() { callbackQRemoteObjectReplica_Notified(this); };
	void setNode(QRemoteObjectNode * node) { callbackQRemoteObjectReplica_SetNode(this, node); };
	void Signal_StateChanged(QRemoteObjectReplica::State state, QRemoteObjectReplica::State oldState) { callbackQRemoteObjectReplica_StateChanged(this, state, oldState); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectReplica_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectReplica_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectReplica_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectReplica_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectReplica_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectReplica_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectReplica_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectReplica_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectReplica_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectRegistry*)
Q_DECLARE_METATYPE(MyQRemoteObjectRegistry*)

int QRemoteObjectRegistry_QRemoteObjectRegistry_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectRegistry*>(); return qRegisterMetaType<MyQRemoteObjectRegistry*>();}

void QRemoteObjectRegistry_DestroyQRemoteObjectRegistry(void* ptr)
{
	static_cast<QRemoteObjectRegistry*>(ptr)->~QRemoteObjectRegistry();
}

void QRemoteObjectRegistry_DestroyQRemoteObjectRegistryDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQRemoteObjectRegistryHost: public QRemoteObjectRegistryHost
{
public:
	MyQRemoteObjectRegistryHost(const QUrl &registryAddress = QUrl(), QObject *parent = Q_NULLPTR) : QRemoteObjectRegistryHost(registryAddress, parent) {QRemoteObjectRegistryHost_QRemoteObjectRegistryHost_QRegisterMetaType();};
	bool setRegistryUrl(const QUrl & registryUrl) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryUrl)) != 0; };
	void setName(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	void Signal_HeartbeatIntervalChanged(int heartbeatInterval) { callbackQRemoteObjectNode_HeartbeatIntervalChanged(this, heartbeatInterval); };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QRemoteObjectRegistryHost*)
Q_DECLARE_METATYPE(MyQRemoteObjectRegistryHost*)

int QRemoteObjectRegistryHost_QRemoteObjectRegistryHost_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectRegistryHost*>(); return qRegisterMetaType<MyQRemoteObjectRegistryHost*>();}

void* QRemoteObjectRegistryHost_NewQRemoteObjectRegistryHost(void* registryAddress, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QObject*>(parent));
	}
}

class MyQRemoteObjectReplica: public QRemoteObjectReplica
{
public:
	void Signal_Initialized() { callbackQRemoteObjectReplica_Initialized(this); };
	void Signal_Notified() { callbackQRemoteObjectReplica_Notified(this); };
	void setNode(QRemoteObjectNode * node) { callbackQRemoteObjectReplica_SetNode(this, node); };
	void Signal_StateChanged(QRemoteObjectReplica::State state, QRemoteObjectReplica::State oldState) { callbackQRemoteObjectReplica_StateChanged(this, state, oldState); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectReplica_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectReplica_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectReplica_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectReplica_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectReplica_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectReplica_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectReplica_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectReplica_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectReplica_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectReplica*)
Q_DECLARE_METATYPE(MyQRemoteObjectReplica*)

int QRemoteObjectReplica_QRemoteObjectReplica_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectReplica*>(); return qRegisterMetaType<MyQRemoteObjectReplica*>();}

void QRemoteObjectReplica_ConnectInitialized(void* ptr, long long t)
{
	QObject::connect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)()>(&QRemoteObjectReplica::initialized), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)()>(&MyQRemoteObjectReplica::Signal_Initialized), static_cast<Qt::ConnectionType>(t));
}

void QRemoteObjectReplica_DisconnectInitialized(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)()>(&QRemoteObjectReplica::initialized), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)()>(&MyQRemoteObjectReplica::Signal_Initialized));
}

void QRemoteObjectReplica_Initialized(void* ptr)
{
	static_cast<QRemoteObjectReplica*>(ptr)->initialized();
}

char QRemoteObjectReplica_IsInitialized(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->isInitialized();
}

char QRemoteObjectReplica_IsReplicaValid(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->isReplicaValid();
}

void* QRemoteObjectReplica_Node(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->node();
}

void QRemoteObjectReplica_ConnectNotified(void* ptr, long long t)
{
	QObject::connect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)()>(&QRemoteObjectReplica::notified), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)()>(&MyQRemoteObjectReplica::Signal_Notified), static_cast<Qt::ConnectionType>(t));
}

void QRemoteObjectReplica_DisconnectNotified(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)()>(&QRemoteObjectReplica::notified), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)()>(&MyQRemoteObjectReplica::Signal_Notified));
}

void QRemoteObjectReplica_Notified(void* ptr)
{
	static_cast<QRemoteObjectReplica*>(ptr)->notified();
}

void QRemoteObjectReplica_SetNode(void* ptr, void* node)
{
	static_cast<QRemoteObjectReplica*>(ptr)->setNode(static_cast<QRemoteObjectNode*>(node));
}

void QRemoteObjectReplica_SetNodeDefault(void* ptr, void* node)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::setNode(static_cast<QRemoteObjectNode*>(node));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::setNode(static_cast<QRemoteObjectNode*>(node));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::setNode(static_cast<QRemoteObjectNode*>(node));
	}
}

long long QRemoteObjectReplica_State(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->state();
}

void QRemoteObjectReplica_ConnectStateChanged(void* ptr, long long t)
{
	qRegisterMetaType<QRemoteObjectReplica::State>();
	qRegisterMetaType<QRemoteObjectReplica::State>();
	QObject::connect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&QRemoteObjectReplica::stateChanged), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&MyQRemoteObjectReplica::Signal_StateChanged), static_cast<Qt::ConnectionType>(t));
}

void QRemoteObjectReplica_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&QRemoteObjectReplica::stateChanged), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&MyQRemoteObjectReplica::Signal_StateChanged));
}

void QRemoteObjectReplica_StateChanged(void* ptr, long long state, long long oldState)
{
	static_cast<QRemoteObjectReplica*>(ptr)->stateChanged(static_cast<QRemoteObjectReplica::State>(state), static_cast<QRemoteObjectReplica::State>(oldState));
}

char QRemoteObjectReplica_WaitForSource(void* ptr, int timeout)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->waitForSource(timeout);
}

void* QRemoteObjectReplica___persistProperties_props_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___persistProperties_props_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___persistProperties_props_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectReplica___retrieveProperties_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___retrieveProperties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___retrieveProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectReplica___send_args_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___send_args_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___send_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectReplica___sendWithReply_args_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___sendWithReply_args_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___sendWithReply_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectReplica___setProperties_properties_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___setProperties_properties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___setProperties_properties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectReplica___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectReplica___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectReplica___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QRemoteObjectReplica___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRemoteObjectReplica___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QRemoteObjectReplica___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectReplica___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectReplica___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QRemoteObjectReplica___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectReplica___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectReplica___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QRemoteObjectReplica_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QRemoteObjectReplica_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QRemoteObjectReplica_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::customEvent(static_cast<QEvent*>(event));
	}
}

void QRemoteObjectReplica_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::deleteLater();
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::deleteLater();
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::deleteLater();
	}
}

void QRemoteObjectReplica_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QRemoteObjectReplica_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::event(static_cast<QEvent*>(e));
	}
}

char QRemoteObjectReplica_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void* QRemoteObjectReplica_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::metaObject());
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::metaObject());
	}
}

void QRemoteObjectReplica_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

