// +build !minimal

#define protected public
#define private public

#include "qml.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QDateTime>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFileSelector>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QJSEngine>
#include <QJSValue>
#include <QJSValueIterator>
#include <QLatin1String>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QNetworkAccessManager>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQmlAbstractUrlInterceptor>
#include <QQmlApplicationEngine>
#include <QQmlComponent>
#include <QQmlContext>
#include <QQmlEngine>
#include <QQmlEngineExtensionPlugin>
#include <QQmlError>
#include <QQmlExpression>
#include <QQmlFileSelector>
#include <QQmlImageProviderBase>
#include <QQmlIncubationController>
#include <QQmlIncubator>
#include <QQmlListReference>
#include <QQmlNetworkAccessManagerFactory>
#include <QQmlParserStatus>
#include <QQmlProperty>
#include <QQmlPropertyMap>
#include <QQmlPropertyValueSource>
#include <QQmlScriptString>
#include <QQuickItem>
#include <QRadioData>
#include <QRemoteObjectPendingCallWatcher>
#include <QString>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>

class MyQJSEngine: public QJSEngine
{
public:
	MyQJSEngine() : QJSEngine() {QJSEngine_QJSEngine_QRegisterMetaType();};
	MyQJSEngine(QObject *parent) : QJSEngine(parent) {QJSEngine_QJSEngine_QRegisterMetaType();};
	void Signal_UiLanguageChanged() { callbackQJSEngine_UiLanguageChanged(this); };
	 ~MyQJSEngine() { callbackQJSEngine_DestroyQJSEngine(this); };
	void childEvent(QChildEvent * event) { callbackQJSEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQJSEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQJSEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQJSEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQJSEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQJSEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQJSEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQJSEngine_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQJSEngine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQJSEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQJSEngine_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QJSEngine*)
Q_DECLARE_METATYPE(MyQJSEngine*)

int QJSEngine_QJSEngine_QRegisterMetaType(){qRegisterMetaType<QJSEngine*>(); return qRegisterMetaType<MyQJSEngine*>();}

void* QJSEngine_NewQJSEngine()
{
	return new MyQJSEngine();
}

void* QJSEngine_NewQJSEngine2(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QWindow*>(parent));
	} else {
		return new MyQJSEngine(static_cast<QObject*>(parent));
	}
}

void QJSEngine_CollectGarbage(void* ptr)
{
	static_cast<QJSEngine*>(ptr)->collectGarbage();
}

void* QJSEngine_Evaluate(void* ptr, struct QtQml_PackedString program, struct QtQml_PackedString fileName, int lineNumber)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->evaluate(QString::fromUtf8(program.data, program.len), QString::fromUtf8(fileName.data, fileName.len), lineNumber));
}

void* QJSEngine_FromScriptValue(void* ptr, void* value)
{
	return new QVariant(static_cast<QJSEngine*>(ptr)->fromScriptValue<QVariant>(*static_cast<QJSValue*>(value)));
}

void* QJSEngine_GlobalObject(void* ptr)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->globalObject());
}

void* QJSEngine_ImportModule(void* ptr, struct QtQml_PackedString fileName)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->importModule(QString::fromUtf8(fileName.data, fileName.len)));
}

void QJSEngine_InstallExtensions(void* ptr, long long extensions, void* object)
{
	static_cast<QJSEngine*>(ptr)->installExtensions(static_cast<QJSEngine::Extension>(extensions), *static_cast<QJSValue*>(object));
}

char QJSEngine_IsInterrupted(void* ptr)
{
	return static_cast<QJSEngine*>(ptr)->isInterrupted();
}

void* QJSEngine_NewArray(void* ptr, unsigned int length)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newArray(length));
}

void* QJSEngine_NewErrorObject(void* ptr, long long errorType, struct QtQml_PackedString message)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newErrorObject(static_cast<QJSValue::ErrorType>(errorType), QString::fromUtf8(message.data, message.len)));
}

void* QJSEngine_NewObject(void* ptr)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newObject());
}

void* QJSEngine_NewQMetaObject(void* ptr, void* metaObject)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newQMetaObject(static_cast<QMetaObject*>(metaObject)));
}

void* QJSEngine_NewQObject(void* ptr, void* object)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newQObject(static_cast<QObject*>(object)));
}

void QJSEngine_SetInterrupted(void* ptr, char interrupted)
{
	static_cast<QJSEngine*>(ptr)->setInterrupted(interrupted != 0);
}

void QJSEngine_SetUiLanguage(void* ptr, struct QtQml_PackedString language)
{
	static_cast<QJSEngine*>(ptr)->setUiLanguage(QString::fromUtf8(language.data, language.len));
}

void QJSEngine_ThrowError(void* ptr, struct QtQml_PackedString message)
{
	static_cast<QJSEngine*>(ptr)->throwError(QString::fromUtf8(message.data, message.len));
}

void QJSEngine_ThrowError2(void* ptr, long long errorType, struct QtQml_PackedString message)
{
	static_cast<QJSEngine*>(ptr)->throwError(static_cast<QJSValue::ErrorType>(errorType), QString::fromUtf8(message.data, message.len));
}

void* QJSEngine_ToScriptValue(void* ptr, void* value)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->toScriptValue(*static_cast<QVariant*>(value)));
}

struct QtQml_PackedString QJSEngine_UiLanguage(void* ptr)
{
	return ({ QByteArray* t326a6e = new QByteArray(static_cast<QJSEngine*>(ptr)->uiLanguage().toUtf8()); QtQml_PackedString { const_cast<char*>(t326a6e->prepend("WHITESPACE").constData()+10), t326a6e->size()-10, t326a6e }; });
}

void QJSEngine_ConnectUiLanguageChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QJSEngine*>(ptr), static_cast<void (QJSEngine::*)()>(&QJSEngine::uiLanguageChanged), static_cast<MyQJSEngine*>(ptr), static_cast<void (MyQJSEngine::*)()>(&MyQJSEngine::Signal_UiLanguageChanged), static_cast<Qt::ConnectionType>(t));
}

void QJSEngine_DisconnectUiLanguageChanged(void* ptr)
{
	QObject::disconnect(static_cast<QJSEngine*>(ptr), static_cast<void (QJSEngine::*)()>(&QJSEngine::uiLanguageChanged), static_cast<MyQJSEngine*>(ptr), static_cast<void (MyQJSEngine::*)()>(&MyQJSEngine::Signal_UiLanguageChanged));
}

void QJSEngine_UiLanguageChanged(void* ptr)
{
	static_cast<QJSEngine*>(ptr)->uiLanguageChanged();
}

void QJSEngine_DestroyQJSEngine(void* ptr)
{
	static_cast<QJSEngine*>(ptr)->~QJSEngine();
}

void QJSEngine_DestroyQJSEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QJSEngine_QJSEngine_qjsEngine(void* object)
{
	return qjsEngine(static_cast<QObject*>(object));
}

void* QJSEngine___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QJSEngine___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QJSEngine___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QJSEngine___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJSEngine___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QJSEngine___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QJSEngine___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QJSEngine___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QJSEngine___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QJSEngine___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QJSEngine___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QJSEngine___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QJSEngine_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQmlApplicationEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QJSEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QQmlApplicationEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QJSEngine_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQmlApplicationEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::customEvent(static_cast<QEvent*>(event));
	}
}

void QJSEngine_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QQmlApplicationEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::deleteLater();
	} else if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::deleteLater();
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::deleteLater();
	}
}

void QJSEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QQmlApplicationEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QJSEngine_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QQmlApplicationEngine*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQmlEngine*>(ptr)->QQmlEngine::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QJSEngine*>(ptr)->QJSEngine::event(static_cast<QEvent*>(e));
	}
}

char QJSEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQmlApplicationEngine*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQmlEngine*>(ptr)->QQmlEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QJSEngine*>(ptr)->QJSEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void* QJSEngine_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QQmlApplicationEngine*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::metaObject());
	} else if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QQmlEngine*>(ptr)->QQmlEngine::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QJSEngine*>(ptr)->QJSEngine::metaObject());
	}
}

void QJSEngine_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQmlApplicationEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

Q_DECLARE_METATYPE(QJSValue*)
void* QJSValue_NewQJSValue(long long value)
{
	return new QJSValue(static_cast<QJSValue::SpecialValue>(value));
}

void* QJSValue_NewQJSValue2(void* other)
{
	return new QJSValue(*static_cast<QJSValue*>(other));
}

void* QJSValue_NewQJSValue3(void* other)
{
	return new QJSValue(*static_cast<QJSValue*>(other));
}

void* QJSValue_NewQJSValue4(char value)
{
	return new QJSValue(value != 0);
}

void* QJSValue_NewQJSValue5(int value)
{
	return new QJSValue(value);
}

void* QJSValue_NewQJSValue6(unsigned int value)
{
	return new QJSValue(value);
}

void* QJSValue_NewQJSValue7(double value)
{
	return new QJSValue(value);
}

void* QJSValue_NewQJSValue8(struct QtQml_PackedString value)
{
	return new QJSValue(QString::fromUtf8(value.data, value.len));
}

void* QJSValue_NewQJSValue9(void* value)
{
	return new QJSValue(*static_cast<QLatin1String*>(value));
}

void* QJSValue_NewQJSValue10(char* value)
{
	return new QJSValue(const_cast<const char*>(value));
}

void* QJSValue_Call(void* ptr, void* args)
{
	return new QJSValue(static_cast<QJSValue*>(ptr)->call(({ QList<QJSValue>* tmpP = static_cast<QList<QJSValue>*>(args); QList<QJSValue> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; })));
}

void* QJSValue_CallAsConstructor(void* ptr, void* args)
{
	return new QJSValue(static_cast<QJSValue*>(ptr)->callAsConstructor(({ QList<QJSValue>* tmpP = static_cast<QList<QJSValue>*>(args); QList<QJSValue> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; })));
}

void* QJSValue_CallWithInstance(void* ptr, void* instance, void* args)
{
	return new QJSValue(static_cast<QJSValue*>(ptr)->callWithInstance(*static_cast<QJSValue*>(instance), ({ QList<QJSValue>* tmpP = static_cast<QList<QJSValue>*>(args); QList<QJSValue> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; })));
}

char QJSValue_DeleteProperty(void* ptr, struct QtQml_PackedString name)
{
	return static_cast<QJSValue*>(ptr)->deleteProperty(QString::fromUtf8(name.data, name.len));
}

char QJSValue_Equals(void* ptr, void* other)
{
	return static_cast<QJSValue*>(ptr)->equals(*static_cast<QJSValue*>(other));
}

long long QJSValue_ErrorType(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->errorType();
}

char QJSValue_HasOwnProperty(void* ptr, struct QtQml_PackedString name)
{
	return static_cast<QJSValue*>(ptr)->hasOwnProperty(QString::fromUtf8(name.data, name.len));
}

char QJSValue_HasProperty(void* ptr, struct QtQml_PackedString name)
{
	return static_cast<QJSValue*>(ptr)->hasProperty(QString::fromUtf8(name.data, name.len));
}

char QJSValue_IsArray(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isArray();
}

char QJSValue_IsBool(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isBool();
}

char QJSValue_IsCallable(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isCallable();
}

char QJSValue_IsDate(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isDate();
}

char QJSValue_IsError(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isError();
}

char QJSValue_IsNull(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isNull();
}

char QJSValue_IsNumber(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isNumber();
}

char QJSValue_IsObject(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isObject();
}

char QJSValue_IsQMetaObject(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isQMetaObject();
}

char QJSValue_IsQObject(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isQObject();
}

char QJSValue_IsRegExp(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isRegExp();
}

char QJSValue_IsString(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isString();
}

char QJSValue_IsUndefined(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isUndefined();
}

char QJSValue_IsVariant(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->isVariant();
}

void* QJSValue_Property(void* ptr, struct QtQml_PackedString name)
{
	return new QJSValue(static_cast<QJSValue*>(ptr)->property(QString::fromUtf8(name.data, name.len)));
}

void* QJSValue_Property2(void* ptr, unsigned int arrayIndex)
{
	return new QJSValue(static_cast<QJSValue*>(ptr)->property(arrayIndex));
}

void* QJSValue_Prototype(void* ptr)
{
	return new QJSValue(static_cast<QJSValue*>(ptr)->prototype());
}

void QJSValue_SetProperty(void* ptr, struct QtQml_PackedString name, void* value)
{
	static_cast<QJSValue*>(ptr)->setProperty(QString::fromUtf8(name.data, name.len), *static_cast<QJSValue*>(value));
}

void QJSValue_SetProperty2(void* ptr, unsigned int arrayIndex, void* value)
{
	static_cast<QJSValue*>(ptr)->setProperty(arrayIndex, *static_cast<QJSValue*>(value));
}

void QJSValue_SetPrototype(void* ptr, void* prototype)
{
	static_cast<QJSValue*>(ptr)->setPrototype(*static_cast<QJSValue*>(prototype));
}

char QJSValue_StrictlyEquals(void* ptr, void* other)
{
	return static_cast<QJSValue*>(ptr)->strictlyEquals(*static_cast<QJSValue*>(other));
}

char QJSValue_ToBool(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->toBool();
}

void* QJSValue_ToDateTime(void* ptr)
{
	return new QDateTime(static_cast<QJSValue*>(ptr)->toDateTime());
}

int QJSValue_ToInt(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->toInt();
}

double QJSValue_ToNumber(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->toNumber();
}

void* QJSValue_ToQMetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QJSValue*>(ptr)->toQMetaObject());
}

void* QJSValue_ToQObject(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->toQObject();
}

struct QtQml_PackedString QJSValue_ToString(void* ptr)
{
	return ({ QByteArray* t6adf54 = new QByteArray(static_cast<QJSValue*>(ptr)->toString().toUtf8()); QtQml_PackedString { const_cast<char*>(t6adf54->prepend("WHITESPACE").constData()+10), t6adf54->size()-10, t6adf54 }; });
}

unsigned int QJSValue_ToUInt(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->toUInt();
}

void* QJSValue_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QJSValue*>(ptr)->toVariant());
}

void QJSValue_DestroyQJSValue(void* ptr)
{
	static_cast<QJSValue*>(ptr)->~QJSValue();
}

void* QJSValue___call_args_atList(void* ptr, int i)
{
	return new QJSValue(({QJSValue tmp = static_cast<QList<QJSValue>*>(ptr)->at(i); if (i == static_cast<QList<QJSValue>*>(ptr)->size()-1) { static_cast<QList<QJSValue>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJSValue___call_args_setList(void* ptr, void* i)
{
	static_cast<QList<QJSValue>*>(ptr)->append(*static_cast<QJSValue*>(i));
}

void* QJSValue___call_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QJSValue>();
}

void* QJSValue___callAsConstructor_args_atList(void* ptr, int i)
{
	return new QJSValue(({QJSValue tmp = static_cast<QList<QJSValue>*>(ptr)->at(i); if (i == static_cast<QList<QJSValue>*>(ptr)->size()-1) { static_cast<QList<QJSValue>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJSValue___callAsConstructor_args_setList(void* ptr, void* i)
{
	static_cast<QList<QJSValue>*>(ptr)->append(*static_cast<QJSValue*>(i));
}

void* QJSValue___callAsConstructor_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QJSValue>();
}

void* QJSValue___callWithInstance_args_atList(void* ptr, int i)
{
	return new QJSValue(({QJSValue tmp = static_cast<QList<QJSValue>*>(ptr)->at(i); if (i == static_cast<QList<QJSValue>*>(ptr)->size()-1) { static_cast<QList<QJSValue>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJSValue___callWithInstance_args_setList(void* ptr, void* i)
{
	static_cast<QList<QJSValue>*>(ptr)->append(*static_cast<QJSValue*>(i));
}

void* QJSValue___callWithInstance_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QJSValue>();
}

Q_DECLARE_METATYPE(QJSValueIterator*)
void* QJSValueIterator_NewQJSValueIterator(void* object)
{
	return new QJSValueIterator(*static_cast<QJSValue*>(object));
}

char QJSValueIterator_HasNext(void* ptr)
{
	return static_cast<QJSValueIterator*>(ptr)->hasNext();
}

struct QtQml_PackedString QJSValueIterator_Name(void* ptr)
{
	return ({ QByteArray* t042ed1 = new QByteArray(static_cast<QJSValueIterator*>(ptr)->name().toUtf8()); QtQml_PackedString { const_cast<char*>(t042ed1->prepend("WHITESPACE").constData()+10), t042ed1->size()-10, t042ed1 }; });
}

char QJSValueIterator_Next(void* ptr)
{
	return static_cast<QJSValueIterator*>(ptr)->next();
}

void* QJSValueIterator_Value(void* ptr)
{
	return new QJSValue(static_cast<QJSValueIterator*>(ptr)->value());
}

void QJSValueIterator_DestroyQJSValueIterator(void* ptr)
{
	static_cast<QJSValueIterator*>(ptr)->~QJSValueIterator();
}

class MyQQmlAbstractUrlInterceptor: public QQmlAbstractUrlInterceptor
{
public:
	MyQQmlAbstractUrlInterceptor() : QQmlAbstractUrlInterceptor() {QQmlAbstractUrlInterceptor_QQmlAbstractUrlInterceptor_QRegisterMetaType();};
	QUrl intercept(const QUrl & url, QQmlAbstractUrlInterceptor::DataType ty) { return *static_cast<QUrl*>(callbackQQmlAbstractUrlInterceptor_Intercept(this, const_cast<QUrl*>(&url), ty)); };
	 ~MyQQmlAbstractUrlInterceptor() { callbackQQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(this); };
};

Q_DECLARE_METATYPE(QQmlAbstractUrlInterceptor*)
Q_DECLARE_METATYPE(MyQQmlAbstractUrlInterceptor*)

int QQmlAbstractUrlInterceptor_QQmlAbstractUrlInterceptor_QRegisterMetaType(){qRegisterMetaType<QQmlAbstractUrlInterceptor*>(); return qRegisterMetaType<MyQQmlAbstractUrlInterceptor*>();}

void* QQmlAbstractUrlInterceptor_NewQQmlAbstractUrlInterceptor()
{
	return new MyQQmlAbstractUrlInterceptor();
}

void* QQmlAbstractUrlInterceptor_Intercept(void* ptr, void* url, long long ty)
{
	return new QUrl(static_cast<QQmlAbstractUrlInterceptor*>(ptr)->intercept(*static_cast<QUrl*>(url), static_cast<QQmlAbstractUrlInterceptor::DataType>(ty)));
}

void QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(void* ptr)
{
	static_cast<QQmlAbstractUrlInterceptor*>(ptr)->~QQmlAbstractUrlInterceptor();
}

void QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQQmlApplicationEngine: public QQmlApplicationEngine
{
public:
	MyQQmlApplicationEngine(QObject *parent = Q_NULLPTR) : QQmlApplicationEngine(parent) {QQmlApplicationEngine_QQmlApplicationEngine_QRegisterMetaType();};
	MyQQmlApplicationEngine(const QUrl &url, QObject *parent = Q_NULLPTR) : QQmlApplicationEngine(url, parent) {QQmlApplicationEngine_QQmlApplicationEngine_QRegisterMetaType();};
	MyQQmlApplicationEngine(const QString &filePath, QObject *parent = Q_NULLPTR) : QQmlApplicationEngine(filePath, parent) {QQmlApplicationEngine_QQmlApplicationEngine_QRegisterMetaType();};
	void load(const QUrl & url) { callbackQQmlApplicationEngine_Load(this, const_cast<QUrl*>(&url)); };
	void load(const QString & filePath) { QByteArray* t7df503 = new QByteArray(filePath.toUtf8()); QtQml_PackedString filePathPacked = { const_cast<char*>(t7df503->prepend("WHITESPACE").constData()+10), t7df503->size()-10, t7df503 };callbackQQmlApplicationEngine_Load2(this, filePathPacked); };
	void loadData(const QByteArray & data, const QUrl & url) { callbackQQmlApplicationEngine_LoadData(this, const_cast<QByteArray*>(&data), const_cast<QUrl*>(&url)); };
	void Signal_ObjectCreated(QObject * object, const QUrl & url) { callbackQQmlApplicationEngine_ObjectCreated(this, object, const_cast<QUrl*>(&url)); };
	 ~MyQQmlApplicationEngine() { callbackQQmlApplicationEngine_DestroyQQmlApplicationEngine(this); };
	bool event(QEvent * e) { return callbackQJSEngine_Event(this, e) != 0; };
	void Signal_Exit(int retCode) { callbackQQmlEngine_Exit(this, retCode); };
	void Signal_Quit() { callbackQQmlEngine_Quit(this); };
	void retranslate() { callbackQQmlEngine_Retranslate(this); };
	void Signal_Warnings(const QList<QQmlError> & warnings) { callbackQQmlEngine_Warnings(this, ({ QList<QQmlError>* tmpValuebd207f = new QList<QQmlError>(warnings); QtQml_PackedList { tmpValuebd207f, tmpValuebd207f->size() }; })); };
	void Signal_UiLanguageChanged() { callbackQJSEngine_UiLanguageChanged(this); };
	void childEvent(QChildEvent * event) { callbackQJSEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQJSEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQJSEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQJSEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQJSEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQJSEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQJSEngine_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQJSEngine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQJSEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQJSEngine_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QQmlApplicationEngine*)
Q_DECLARE_METATYPE(MyQQmlApplicationEngine*)

int QQmlApplicationEngine_QQmlApplicationEngine_QRegisterMetaType(){qRegisterMetaType<QQmlApplicationEngine*>(); return qRegisterMetaType<MyQQmlApplicationEngine*>();}

void* QQmlApplicationEngine_NewQQmlApplicationEngine(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlApplicationEngine(static_cast<QObject*>(parent));
	}
}

void* QQmlApplicationEngine_NewQQmlApplicationEngine2(void* url, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlApplicationEngine(*static_cast<QUrl*>(url), static_cast<QObject*>(parent));
	}
}

void* QQmlApplicationEngine_NewQQmlApplicationEngine3(struct QtQml_PackedString filePath, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlApplicationEngine(QString::fromUtf8(filePath.data, filePath.len), static_cast<QObject*>(parent));
	}
}

void QQmlApplicationEngine_Load(void* ptr, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "load", Q_ARG(const QUrl, *static_cast<QUrl*>(url)));
}

void QQmlApplicationEngine_LoadDefault(void* ptr, void* url)
{
		static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::load(*static_cast<QUrl*>(url));
}

void QQmlApplicationEngine_Load2(void* ptr, struct QtQml_PackedString filePath)
{
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "load", Q_ARG(const QString, QString::fromUtf8(filePath.data, filePath.len)));
}

void QQmlApplicationEngine_Load2Default(void* ptr, struct QtQml_PackedString filePath)
{
		static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::load(QString::fromUtf8(filePath.data, filePath.len));
}

void QQmlApplicationEngine_LoadData(void* ptr, void* data, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQmlApplicationEngine*>(ptr), "loadData", Q_ARG(const QByteArray, *static_cast<QByteArray*>(data)), Q_ARG(const QUrl, *static_cast<QUrl*>(url)));
}

void QQmlApplicationEngine_LoadDataDefault(void* ptr, void* data, void* url)
{
		static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::loadData(*static_cast<QByteArray*>(data), *static_cast<QUrl*>(url));
}

void QQmlApplicationEngine_ConnectObjectCreated(void* ptr, long long t)
{
	QObject::connect(static_cast<QQmlApplicationEngine*>(ptr), static_cast<void (QQmlApplicationEngine::*)(QObject *, const QUrl &)>(&QQmlApplicationEngine::objectCreated), static_cast<MyQQmlApplicationEngine*>(ptr), static_cast<void (MyQQmlApplicationEngine::*)(QObject *, const QUrl &)>(&MyQQmlApplicationEngine::Signal_ObjectCreated), static_cast<Qt::ConnectionType>(t));
}

void QQmlApplicationEngine_DisconnectObjectCreated(void* ptr)
{
	QObject::disconnect(static_cast<QQmlApplicationEngine*>(ptr), static_cast<void (QQmlApplicationEngine::*)(QObject *, const QUrl &)>(&QQmlApplicationEngine::objectCreated), static_cast<MyQQmlApplicationEngine*>(ptr), static_cast<void (MyQQmlApplicationEngine::*)(QObject *, const QUrl &)>(&MyQQmlApplicationEngine::Signal_ObjectCreated));
}

void QQmlApplicationEngine_ObjectCreated(void* ptr, void* object, void* url)
{
	static_cast<QQmlApplicationEngine*>(ptr)->objectCreated(static_cast<QObject*>(object), *static_cast<QUrl*>(url));
}

struct QtQml_PackedList QQmlApplicationEngine_RootObjects(void* ptr)
{
	return ({ QList<QObject *>* tmpValue21875f = new QList<QObject *>(static_cast<QQmlApplicationEngine*>(ptr)->rootObjects()); QtQml_PackedList { tmpValue21875f, tmpValue21875f->size() }; });
}

void QQmlApplicationEngine_DestroyQQmlApplicationEngine(void* ptr)
{
	static_cast<QQmlApplicationEngine*>(ptr)->~QQmlApplicationEngine();
}

void QQmlApplicationEngine_DestroyQQmlApplicationEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQmlApplicationEngine___rootObjects_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlApplicationEngine___rootObjects_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlApplicationEngine___rootObjects_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QQmlApplicationEngine___setInitialProperties_initialProperties_atList(void* ptr, struct QtQml_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QQmlApplicationEngine___setInitialProperties_initialProperties_setList(void* ptr, struct QtQml_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QQmlApplicationEngine___setInitialProperties_initialProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtQml_PackedList QQmlApplicationEngine___setInitialProperties_initialProperties_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtQml_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtQml_PackedString QQmlApplicationEngine_____setInitialProperties_initialProperties_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtQml_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QQmlApplicationEngine_____setInitialProperties_initialProperties_keyList_setList(void* ptr, struct QtQml_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QQmlApplicationEngine_____setInitialProperties_initialProperties_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

class MyQQmlComponent: public QQmlComponent
{
public:
	MyQQmlComponent(QQmlEngine *engine, QObject *parent = Q_NULLPTR) : QQmlComponent(engine, parent) {QQmlComponent_QQmlComponent_QRegisterMetaType();};
	MyQQmlComponent(QQmlEngine *engine, const QString &fileName, QObject *parent = Q_NULLPTR) : QQmlComponent(engine, fileName, parent) {QQmlComponent_QQmlComponent_QRegisterMetaType();};
	MyQQmlComponent(QQmlEngine *engine, const QString &fileName, QQmlComponent::CompilationMode mode, QObject *parent = Q_NULLPTR) : QQmlComponent(engine, fileName, mode, parent) {QQmlComponent_QQmlComponent_QRegisterMetaType();};
	MyQQmlComponent(QQmlEngine *engine, const QUrl &url, QObject *parent = Q_NULLPTR) : QQmlComponent(engine, url, parent) {QQmlComponent_QQmlComponent_QRegisterMetaType();};
	MyQQmlComponent(QQmlEngine *engine, const QUrl &url, QQmlComponent::CompilationMode mode, QObject *parent = Q_NULLPTR) : QQmlComponent(engine, url, mode, parent) {QQmlComponent_QQmlComponent_QRegisterMetaType();};
	QObject * beginCreate(QQmlContext * publicContext) { return static_cast<QObject*>(callbackQQmlComponent_BeginCreate(this, publicContext)); };
	void completeCreate() { callbackQQmlComponent_CompleteCreate(this); };
	QObject * create(QQmlContext * context) { return static_cast<QObject*>(callbackQQmlComponent_Create(this, context)); };
	void loadUrl(const QUrl & url) { callbackQQmlComponent_LoadUrl(this, const_cast<QUrl*>(&url)); };
	void loadUrl(const QUrl & url, QQmlComponent::CompilationMode mode) { callbackQQmlComponent_LoadUrl2(this, const_cast<QUrl*>(&url), mode); };
	void Signal_ProgressChanged(qreal progress) { callbackQQmlComponent_ProgressChanged(this, progress); };
	void setData(const QByteArray & data, const QUrl & url) { callbackQQmlComponent_SetData(this, const_cast<QByteArray*>(&data), const_cast<QUrl*>(&url)); };
	void Signal_StatusChanged(QQmlComponent::Status status) { callbackQQmlComponent_StatusChanged(this, status); };
	 ~MyQQmlComponent() { callbackQQmlComponent_DestroyQQmlComponent(this); };
	void childEvent(QChildEvent * event) { callbackQQmlComponent_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlComponent_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQmlComponent_CustomEvent(this, event); };
	void deleteLater() { callbackQQmlComponent_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQmlComponent_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlComponent_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQQmlComponent_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlComponent_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlComponent_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQQmlComponent_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQmlComponent_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQQmlComponent*)

int QQmlComponent_QQmlComponent_QRegisterMetaType(){qRegisterMetaType<QQmlComponent*>(); return qRegisterMetaType<MyQQmlComponent*>();}

void* QQmlComponent_NewQQmlComponent2(void* engine, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
	}
}

void* QQmlComponent_NewQQmlComponent3(void* engine, struct QtQml_PackedString fileName, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QObject*>(parent));
	}
}

void* QQmlComponent_NewQQmlComponent4(void* engine, struct QtQml_PackedString fileName, long long mode, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), QString::fromUtf8(fileName.data, fileName.len), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QObject*>(parent));
	}
}

void* QQmlComponent_NewQQmlComponent5(void* engine, void* url, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QObject*>(parent));
	}
}

void* QQmlComponent_NewQQmlComponent6(void* engine, void* url, long long mode, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlComponent(static_cast<QQmlEngine*>(engine), *static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode), static_cast<QObject*>(parent));
	}
}

void* QQmlComponent_BeginCreate(void* ptr, void* publicContext)
{
	return static_cast<QQmlComponent*>(ptr)->beginCreate(static_cast<QQmlContext*>(publicContext));
}

void* QQmlComponent_BeginCreateDefault(void* ptr, void* publicContext)
{
		return static_cast<QQmlComponent*>(ptr)->QQmlComponent::beginCreate(static_cast<QQmlContext*>(publicContext));
}

void QQmlComponent_CompleteCreate(void* ptr)
{
	static_cast<QQmlComponent*>(ptr)->completeCreate();
}

void QQmlComponent_CompleteCreateDefault(void* ptr)
{
		static_cast<QQmlComponent*>(ptr)->QQmlComponent::completeCreate();
}

void* QQmlComponent_Create(void* ptr, void* context)
{
	return static_cast<QQmlComponent*>(ptr)->create(static_cast<QQmlContext*>(context));
}

void* QQmlComponent_CreateDefault(void* ptr, void* context)
{
		return static_cast<QQmlComponent*>(ptr)->QQmlComponent::create(static_cast<QQmlContext*>(context));
}

void QQmlComponent_Create2(void* ptr, void* incubator, void* context, void* forContext)
{
	static_cast<QQmlComponent*>(ptr)->create(*static_cast<QQmlIncubator*>(incubator), static_cast<QQmlContext*>(context), static_cast<QQmlContext*>(forContext));
}

void* QQmlComponent_CreateWithInitialProperties(void* ptr, void* initialProperties, void* context)
{
	return static_cast<QQmlComponent*>(ptr)->createWithInitialProperties(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(initialProperties); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QQmlContext*>(context));
}

void* QQmlComponent_CreationContext(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->creationContext();
}

void* QQmlComponent_Engine(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->engine();
}

struct QtQml_PackedList QQmlComponent_Errors(void* ptr)
{
	return ({ QList<QQmlError>* tmpValue422633 = new QList<QQmlError>(static_cast<QQmlComponent*>(ptr)->errors()); QtQml_PackedList { tmpValue422633, tmpValue422633->size() }; });
}

char QQmlComponent_IsError(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->isError();
}

char QQmlComponent_IsLoading(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->isLoading();
}

char QQmlComponent_IsNull(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->isNull();
}

char QQmlComponent_IsReady(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->isReady();
}

void QQmlComponent_LoadUrl(void* ptr, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "loadUrl", Q_ARG(const QUrl, *static_cast<QUrl*>(url)));
}

void QQmlComponent_LoadUrlDefault(void* ptr, void* url)
{
		static_cast<QQmlComponent*>(ptr)->QQmlComponent::loadUrl(*static_cast<QUrl*>(url));
}

void QQmlComponent_LoadUrl2(void* ptr, void* url, long long mode)
{
	qRegisterMetaType<QQmlComponent::CompilationMode>();
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "loadUrl", Q_ARG(const QUrl, *static_cast<QUrl*>(url)), Q_ARG(QQmlComponent::CompilationMode, static_cast<QQmlComponent::CompilationMode>(mode)));
}

void QQmlComponent_LoadUrl2Default(void* ptr, void* url, long long mode)
{
		static_cast<QQmlComponent*>(ptr)->QQmlComponent::loadUrl(*static_cast<QUrl*>(url), static_cast<QQmlComponent::CompilationMode>(mode));
}

double QQmlComponent_Progress(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->progress();
}

void QQmlComponent_ConnectProgressChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(qreal)>(&QQmlComponent::progressChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(qreal)>(&MyQQmlComponent::Signal_ProgressChanged), static_cast<Qt::ConnectionType>(t));
}

void QQmlComponent_DisconnectProgressChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(qreal)>(&QQmlComponent::progressChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(qreal)>(&MyQQmlComponent::Signal_ProgressChanged));
}

void QQmlComponent_ProgressChanged(void* ptr, double progress)
{
	static_cast<QQmlComponent*>(ptr)->progressChanged(progress);
}

void QQmlComponent_SetData(void* ptr, void* data, void* url)
{
	QMetaObject::invokeMethod(static_cast<QQmlComponent*>(ptr), "setData", Q_ARG(const QByteArray, *static_cast<QByteArray*>(data)), Q_ARG(const QUrl, *static_cast<QUrl*>(url)));
}

void QQmlComponent_SetDataDefault(void* ptr, void* data, void* url)
{
		static_cast<QQmlComponent*>(ptr)->QQmlComponent::setData(*static_cast<QByteArray*>(data), *static_cast<QUrl*>(url));
}

long long QQmlComponent_Status(void* ptr)
{
	return static_cast<QQmlComponent*>(ptr)->status();
}

void QQmlComponent_ConnectStatusChanged(void* ptr, long long t)
{
	qRegisterMetaType<QQmlComponent::Status>();
	QObject::connect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(QQmlComponent::Status)>(&QQmlComponent::statusChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(QQmlComponent::Status)>(&MyQQmlComponent::Signal_StatusChanged), static_cast<Qt::ConnectionType>(t));
}

void QQmlComponent_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlComponent*>(ptr), static_cast<void (QQmlComponent::*)(QQmlComponent::Status)>(&QQmlComponent::statusChanged), static_cast<MyQQmlComponent*>(ptr), static_cast<void (MyQQmlComponent::*)(QQmlComponent::Status)>(&MyQQmlComponent::Signal_StatusChanged));
}

void QQmlComponent_StatusChanged(void* ptr, long long status)
{
	static_cast<QQmlComponent*>(ptr)->statusChanged(static_cast<QQmlComponent::Status>(status));
}

void* QQmlComponent_Url(void* ptr)
{
	return new QUrl(static_cast<QQmlComponent*>(ptr)->url());
}

void QQmlComponent_DestroyQQmlComponent(void* ptr)
{
	static_cast<QQmlComponent*>(ptr)->~QQmlComponent();
}

void QQmlComponent_DestroyQQmlComponentDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQmlComponent___createWithInitialProperties_initialProperties_atList(void* ptr, struct QtQml_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QQmlComponent___createWithInitialProperties_initialProperties_setList(void* ptr, struct QtQml_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QQmlComponent___createWithInitialProperties_initialProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtQml_PackedList QQmlComponent___createWithInitialProperties_initialProperties_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtQml_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QQmlComponent___errors_atList(void* ptr, int i)
{
	return new QQmlError(({QQmlError tmp = static_cast<QList<QQmlError>*>(ptr)->at(i); if (i == static_cast<QList<QQmlError>*>(ptr)->size()-1) { static_cast<QList<QQmlError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlComponent___errors_setList(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQmlComponent___errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>();
}

void* QQmlComponent___setInitialProperties_properties_atList(void* ptr, struct QtQml_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QQmlComponent___setInitialProperties_properties_setList(void* ptr, struct QtQml_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QQmlComponent___setInitialProperties_properties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtQml_PackedList QQmlComponent___setInitialProperties_properties_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtQml_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtQml_PackedString QQmlComponent_____createWithInitialProperties_initialProperties_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtQml_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QQmlComponent_____createWithInitialProperties_initialProperties_keyList_setList(void* ptr, struct QtQml_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QQmlComponent_____createWithInitialProperties_initialProperties_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtQml_PackedString QQmlComponent_____setInitialProperties_properties_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtQml_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QQmlComponent_____setInitialProperties_properties_keyList_setList(void* ptr, struct QtQml_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QQmlComponent_____setInitialProperties_properties_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* QQmlComponent___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlComponent___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlComponent___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QQmlComponent___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlComponent___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQmlComponent___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QQmlComponent___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlComponent___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlComponent___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQmlComponent___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlComponent___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlComponent___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QQmlComponent_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQmlComponent*>(ptr)->QQmlComponent::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlComponent_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlComponent*>(ptr)->QQmlComponent::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlComponent_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQmlComponent*>(ptr)->QQmlComponent::customEvent(static_cast<QEvent*>(event));
}

void QQmlComponent_DeleteLaterDefault(void* ptr)
{
		static_cast<QQmlComponent*>(ptr)->QQmlComponent::deleteLater();
}

void QQmlComponent_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlComponent*>(ptr)->QQmlComponent::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QQmlComponent_EventDefault(void* ptr, void* e)
{
		return static_cast<QQmlComponent*>(ptr)->QQmlComponent::event(static_cast<QEvent*>(e));
}

char QQmlComponent_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQmlComponent*>(ptr)->QQmlComponent::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlComponent_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQmlComponent*>(ptr)->QQmlComponent::metaObject());
}

void QQmlComponent_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQmlComponent*>(ptr)->QQmlComponent::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQQmlContext: public QQmlContext
{
public:
	MyQQmlContext(QQmlEngine *engine, QObject *parent = Q_NULLPTR) : QQmlContext(engine, parent) {QQmlContext_QQmlContext_QRegisterMetaType();};
	MyQQmlContext(QQmlContext *parentContext, QObject *parent = Q_NULLPTR) : QQmlContext(parentContext, parent) {QQmlContext_QQmlContext_QRegisterMetaType();};
	 ~MyQQmlContext() { callbackQQmlContext_DestroyQQmlContext(this); };
	void childEvent(QChildEvent * event) { callbackQQmlContext_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlContext_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQmlContext_CustomEvent(this, event); };
	void deleteLater() { callbackQQmlContext_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQmlContext_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlContext_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQQmlContext_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlContext_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlContext_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQQmlContext_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQmlContext_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QQmlContext*)
Q_DECLARE_METATYPE(MyQQmlContext*)

int QQmlContext_QQmlContext_QRegisterMetaType(){qRegisterMetaType<QQmlContext*>(); return qRegisterMetaType<MyQQmlContext*>();}

void* QQmlContext_NewQQmlContext(void* engine, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlContext(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
	}
}

void* QQmlContext_NewQQmlContext2(void* parentContext, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlContext(static_cast<QQmlContext*>(parentContext), static_cast<QObject*>(parent));
	}
}

void* QQmlContext_BaseUrl(void* ptr)
{
	return new QUrl(static_cast<QQmlContext*>(ptr)->baseUrl());
}

void* QQmlContext_ContextObject(void* ptr)
{
	return static_cast<QQmlContext*>(ptr)->contextObject();
}

void* QQmlContext_ContextProperty(void* ptr, struct QtQml_PackedString name)
{
	return new QVariant(static_cast<QQmlContext*>(ptr)->contextProperty(QString::fromUtf8(name.data, name.len)));
}

void* QQmlContext_Engine(void* ptr)
{
	return static_cast<QQmlContext*>(ptr)->engine();
}

char QQmlContext_IsValid(void* ptr)
{
	return static_cast<QQmlContext*>(ptr)->isValid();
}

struct QtQml_PackedString QQmlContext_NameForObject(void* ptr, void* object)
{
	return ({ QByteArray* t183ce4 = new QByteArray(static_cast<QQmlContext*>(ptr)->nameForObject(static_cast<QObject*>(object)).toUtf8()); QtQml_PackedString { const_cast<char*>(t183ce4->prepend("WHITESPACE").constData()+10), t183ce4->size()-10, t183ce4 }; });
}

void* QQmlContext_ParentContext(void* ptr)
{
	return static_cast<QQmlContext*>(ptr)->parentContext();
}

void* QQmlContext_ResolvedUrl(void* ptr, void* src)
{
	return new QUrl(static_cast<QQmlContext*>(ptr)->resolvedUrl(*static_cast<QUrl*>(src)));
}

void QQmlContext_SetBaseUrl(void* ptr, void* baseUrl)
{
	static_cast<QQmlContext*>(ptr)->setBaseUrl(*static_cast<QUrl*>(baseUrl));
}

void QQmlContext_SetContextObject(void* ptr, void* object)
{
	static_cast<QQmlContext*>(ptr)->setContextObject(static_cast<QObject*>(object));
}

void QQmlContext_SetContextProperty(void* ptr, struct QtQml_PackedString name, void* value)
{
	static_cast<QQmlContext*>(ptr)->setContextProperty(QString::fromUtf8(name.data, name.len), static_cast<QObject*>(value));
}

void QQmlContext_SetContextProperty2(void* ptr, struct QtQml_PackedString name, void* value)
{
	static_cast<QQmlContext*>(ptr)->setContextProperty(QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value));
}

void QQmlContext_DestroyQQmlContext(void* ptr)
{
	static_cast<QQmlContext*>(ptr)->~QQmlContext();
}

void QQmlContext_DestroyQQmlContextDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQmlContext___setContextProperties_properties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QQmlContext::PropertyPair>();
}

void* QQmlContext___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlContext___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlContext___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QQmlContext___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlContext___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQmlContext___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QQmlContext___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlContext___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlContext___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQmlContext___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlContext___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlContext___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QQmlContext_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQmlContext*>(ptr)->QQmlContext::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlContext_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlContext*>(ptr)->QQmlContext::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlContext_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQmlContext*>(ptr)->QQmlContext::customEvent(static_cast<QEvent*>(event));
}

void QQmlContext_DeleteLaterDefault(void* ptr)
{
		static_cast<QQmlContext*>(ptr)->QQmlContext::deleteLater();
}

void QQmlContext_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlContext*>(ptr)->QQmlContext::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QQmlContext_EventDefault(void* ptr, void* e)
{
		return static_cast<QQmlContext*>(ptr)->QQmlContext::event(static_cast<QEvent*>(e));
}

char QQmlContext_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQmlContext*>(ptr)->QQmlContext::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlContext_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQmlContext*>(ptr)->QQmlContext::metaObject());
}

void QQmlContext_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQmlContext*>(ptr)->QQmlContext::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQQmlEngine: public QQmlEngine
{
public:
	MyQQmlEngine(QObject *parent = Q_NULLPTR) : QQmlEngine(parent) {QQmlEngine_QQmlEngine_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQJSEngine_Event(this, e) != 0; };
	void Signal_Exit(int retCode) { callbackQQmlEngine_Exit(this, retCode); };
	void Signal_Quit() { callbackQQmlEngine_Quit(this); };
	void retranslate() { callbackQQmlEngine_Retranslate(this); };
	void Signal_Warnings(const QList<QQmlError> & warnings) { callbackQQmlEngine_Warnings(this, ({ QList<QQmlError>* tmpValuebd207f = new QList<QQmlError>(warnings); QtQml_PackedList { tmpValuebd207f, tmpValuebd207f->size() }; })); };
	 ~MyQQmlEngine() { callbackQQmlEngine_DestroyQQmlEngine(this); };
	void Signal_UiLanguageChanged() { callbackQJSEngine_UiLanguageChanged(this); };
	void childEvent(QChildEvent * event) { callbackQJSEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQJSEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQJSEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQJSEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQJSEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQJSEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQJSEngine_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQJSEngine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQJSEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQJSEngine_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QQmlEngine*)
Q_DECLARE_METATYPE(MyQQmlEngine*)

int QQmlEngine_QQmlEngine_QRegisterMetaType(){qRegisterMetaType<QQmlEngine*>(); return qRegisterMetaType<MyQQmlEngine*>();}

void* QQmlEngine_NewQQmlEngine(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlEngine(static_cast<QObject*>(parent));
	}
}

void QQmlEngine_AddImageProvider(void* ptr, struct QtQml_PackedString providerId, void* provider)
{
	static_cast<QQmlEngine*>(ptr)->addImageProvider(QString::fromUtf8(providerId.data, providerId.len), static_cast<QQmlImageProviderBase*>(provider));
}

void QQmlEngine_AddImportPath(void* ptr, struct QtQml_PackedString path)
{
	static_cast<QQmlEngine*>(ptr)->addImportPath(QString::fromUtf8(path.data, path.len));
}

void QQmlEngine_AddPluginPath(void* ptr, struct QtQml_PackedString path)
{
	static_cast<QQmlEngine*>(ptr)->addPluginPath(QString::fromUtf8(path.data, path.len));
}

void* QQmlEngine_BaseUrl(void* ptr)
{
	return new QUrl(static_cast<QQmlEngine*>(ptr)->baseUrl());
}

void QQmlEngine_ClearComponentCache(void* ptr)
{
	static_cast<QQmlEngine*>(ptr)->clearComponentCache();
}

void* QQmlEngine_QQmlEngine_ContextForObject(void* object)
{
	return QQmlEngine::contextForObject(static_cast<QObject*>(object));
}

void QQmlEngine_ConnectExit(void* ptr, long long t)
{
	QObject::connect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)(int)>(&QQmlEngine::exit), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)(int)>(&MyQQmlEngine::Signal_Exit), static_cast<Qt::ConnectionType>(t));
}

void QQmlEngine_DisconnectExit(void* ptr)
{
	QObject::disconnect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)(int)>(&QQmlEngine::exit), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)(int)>(&MyQQmlEngine::Signal_Exit));
}

void QQmlEngine_Exit(void* ptr, int retCode)
{
	static_cast<QQmlEngine*>(ptr)->exit(retCode);
}

void* QQmlEngine_ImageProvider(void* ptr, struct QtQml_PackedString providerId)
{
	return static_cast<QQmlEngine*>(ptr)->imageProvider(QString::fromUtf8(providerId.data, providerId.len));
}

struct QtQml_PackedString QQmlEngine_ImportPathList(void* ptr)
{
	return ({ QByteArray* t32e5da = new QByteArray(static_cast<QQmlEngine*>(ptr)->importPathList().join("¡¦!").toUtf8()); QtQml_PackedString { const_cast<char*>(t32e5da->prepend("WHITESPACE").constData()+10), t32e5da->size()-10, t32e5da }; });
}

char QQmlEngine_ImportPlugin(void* ptr, struct QtQml_PackedString filePath, struct QtQml_PackedString uri, void* errors)
{
	return static_cast<QQmlEngine*>(ptr)->importPlugin(QString::fromUtf8(filePath.data, filePath.len), QString::fromUtf8(uri.data, uri.len), static_cast<QList<QQmlError>*>(errors));
}

void* QQmlEngine_IncubationController(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->incubationController();
}

void* QQmlEngine_NetworkAccessManager(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->networkAccessManager();
}

void* QQmlEngine_NetworkAccessManagerFactory(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->networkAccessManagerFactory();
}

long long QQmlEngine_QQmlEngine_ObjectOwnership(void* object)
{
	return QQmlEngine::objectOwnership(static_cast<QObject*>(object));
}

struct QtQml_PackedString QQmlEngine_OfflineStorageDatabaseFilePath(void* ptr, struct QtQml_PackedString databaseName)
{
	return ({ QByteArray* t7b641b = new QByteArray(static_cast<QQmlEngine*>(ptr)->offlineStorageDatabaseFilePath(QString::fromUtf8(databaseName.data, databaseName.len)).toUtf8()); QtQml_PackedString { const_cast<char*>(t7b641b->prepend("WHITESPACE").constData()+10), t7b641b->size()-10, t7b641b }; });
}

struct QtQml_PackedString QQmlEngine_OfflineStoragePath(void* ptr)
{
	return ({ QByteArray* t6f33d6 = new QByteArray(static_cast<QQmlEngine*>(ptr)->offlineStoragePath().toUtf8()); QtQml_PackedString { const_cast<char*>(t6f33d6->prepend("WHITESPACE").constData()+10), t6f33d6->size()-10, t6f33d6 }; });
}

char QQmlEngine_OutputWarningsToStandardError(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->outputWarningsToStandardError();
}

struct QtQml_PackedString QQmlEngine_PluginPathList(void* ptr)
{
	return ({ QByteArray* t04b834 = new QByteArray(static_cast<QQmlEngine*>(ptr)->pluginPathList().join("¡¦!").toUtf8()); QtQml_PackedString { const_cast<char*>(t04b834->prepend("WHITESPACE").constData()+10), t04b834->size()-10, t04b834 }; });
}

void QQmlEngine_ConnectQuit(void* ptr, long long t)
{
	QObject::connect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)()>(&QQmlEngine::quit), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)()>(&MyQQmlEngine::Signal_Quit), static_cast<Qt::ConnectionType>(t));
}

void QQmlEngine_DisconnectQuit(void* ptr)
{
	QObject::disconnect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)()>(&QQmlEngine::quit), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)()>(&MyQQmlEngine::Signal_Quit));
}

void QQmlEngine_Quit(void* ptr)
{
	static_cast<QQmlEngine*>(ptr)->quit();
}

void QQmlEngine_RemoveImageProvider(void* ptr, struct QtQml_PackedString providerId)
{
	static_cast<QQmlEngine*>(ptr)->removeImageProvider(QString::fromUtf8(providerId.data, providerId.len));
}

void QQmlEngine_Retranslate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QQmlEngine*>(ptr), "retranslate");
}

void QQmlEngine_RetranslateDefault(void* ptr)
{
	if (dynamic_cast<QQmlApplicationEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlApplicationEngine*>(ptr)->QQmlApplicationEngine::retranslate();
	} else {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::retranslate();
	}
}

void* QQmlEngine_RootContext(void* ptr)
{
	return static_cast<QQmlEngine*>(ptr)->rootContext();
}

void QQmlEngine_SetBaseUrl(void* ptr, void* url)
{
	static_cast<QQmlEngine*>(ptr)->setBaseUrl(*static_cast<QUrl*>(url));
}

void QQmlEngine_QQmlEngine_SetContextForObject(void* object, void* context)
{
	QQmlEngine::setContextForObject(static_cast<QObject*>(object), static_cast<QQmlContext*>(context));
}

void QQmlEngine_SetImportPathList(void* ptr, struct QtQml_PackedString paths)
{
	static_cast<QQmlEngine*>(ptr)->setImportPathList(QString::fromUtf8(paths.data, paths.len).split("¡¦!", QString::SkipEmptyParts));
}

void QQmlEngine_SetIncubationController(void* ptr, void* controller)
{
	static_cast<QQmlEngine*>(ptr)->setIncubationController(static_cast<QQmlIncubationController*>(controller));
}

void QQmlEngine_SetNetworkAccessManagerFactory(void* ptr, void* factory)
{
	static_cast<QQmlEngine*>(ptr)->setNetworkAccessManagerFactory(static_cast<QQmlNetworkAccessManagerFactory*>(factory));
}

void QQmlEngine_QQmlEngine_SetObjectOwnership(void* object, long long ownership)
{
	QQmlEngine::setObjectOwnership(static_cast<QObject*>(object), static_cast<QQmlEngine::ObjectOwnership>(ownership));
}

void QQmlEngine_SetOfflineStoragePath(void* ptr, struct QtQml_PackedString dir)
{
	static_cast<QQmlEngine*>(ptr)->setOfflineStoragePath(QString::fromUtf8(dir.data, dir.len));
}

void QQmlEngine_SetOutputWarningsToStandardError(void* ptr, char enabled)
{
	static_cast<QQmlEngine*>(ptr)->setOutputWarningsToStandardError(enabled != 0);
}

void QQmlEngine_SetPluginPathList(void* ptr, struct QtQml_PackedString paths)
{
	static_cast<QQmlEngine*>(ptr)->setPluginPathList(QString::fromUtf8(paths.data, paths.len).split("¡¦!", QString::SkipEmptyParts));
}

void QQmlEngine_TrimComponentCache(void* ptr)
{
	static_cast<QQmlEngine*>(ptr)->trimComponentCache();
}

void QQmlEngine_ConnectWarnings(void* ptr, long long t)
{
	QObject::connect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)(const QList<QQmlError> &)>(&QQmlEngine::warnings), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)(const QList<QQmlError> &)>(&MyQQmlEngine::Signal_Warnings), static_cast<Qt::ConnectionType>(t));
}

void QQmlEngine_DisconnectWarnings(void* ptr)
{
	QObject::disconnect(static_cast<QQmlEngine*>(ptr), static_cast<void (QQmlEngine::*)(const QList<QQmlError> &)>(&QQmlEngine::warnings), static_cast<MyQQmlEngine*>(ptr), static_cast<void (MyQQmlEngine::*)(const QList<QQmlError> &)>(&MyQQmlEngine::Signal_Warnings));
}

void QQmlEngine_Warnings(void* ptr, void* warnings)
{
	static_cast<QQmlEngine*>(ptr)->warnings(*static_cast<QList<QQmlError>*>(warnings));
}

void QQmlEngine_DestroyQQmlEngine(void* ptr)
{
	static_cast<QQmlEngine*>(ptr)->~QQmlEngine();
}

void QQmlEngine_DestroyQQmlEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

int QQmlEngine_QQmlEngine_QmlRegisterSingletonType(void* url, char* uri, int versionMajor, int versionMinor, char* qmlName)
{
	return qmlRegisterSingletonType(*static_cast<QUrl*>(url), const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
}

int QQmlEngine_QQmlEngine_QmlRegisterType(void* url, char* uri, int versionMajor, int versionMinor, char* qmlName)
{
	return qmlRegisterType(*static_cast<QUrl*>(url), const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
}

int QQmlEngine_QQmlEngine_QmlRegisterSingletonInstance(char* uri, int versionMajor, int versionMinor, char* typeName, void* cppObject)
{
	return qmlRegisterSingletonInstance(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(typeName), static_cast<QObject*>(cppObject));
}

void* QQmlEngine___importPlugin_errors_atList(void* ptr, int i)
{
	return new QQmlError(({QQmlError tmp = static_cast<QList<QQmlError>*>(ptr)->at(i); if (i == static_cast<QList<QQmlError>*>(ptr)->size()-1) { static_cast<QList<QQmlError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlEngine___importPlugin_errors_setList(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQmlEngine___importPlugin_errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>();
}

void* QQmlEngine___warnings_warnings_atList(void* ptr, int i)
{
	return new QQmlError(({QQmlError tmp = static_cast<QList<QQmlError>*>(ptr)->at(i); if (i == static_cast<QList<QQmlError>*>(ptr)->size()-1) { static_cast<QList<QQmlError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlEngine___warnings_warnings_setList(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQmlEngine___warnings_warnings_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>();
}

class MyQQmlEngineExtensionPlugin: public QQmlEngineExtensionPlugin
{
public:
	MyQQmlEngineExtensionPlugin(QObject *parent = Q_NULLPTR) : QQmlEngineExtensionPlugin(parent) {QQmlEngineExtensionPlugin_QQmlEngineExtensionPlugin_QRegisterMetaType();};
	void initializeEngine(QQmlEngine * engine, const char * uri) { QtQml_PackedString uriPacked = { const_cast<char*>(uri), -1, NULL };callbackQQmlEngineExtensionPlugin_InitializeEngine(this, engine, uriPacked); };
	void childEvent(QChildEvent * event) { callbackQQmlEngineExtensionPlugin_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlEngineExtensionPlugin_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQmlEngineExtensionPlugin_CustomEvent(this, event); };
	void deleteLater() { callbackQQmlEngineExtensionPlugin_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQmlEngineExtensionPlugin_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlEngineExtensionPlugin_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQQmlEngineExtensionPlugin_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlEngineExtensionPlugin_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlEngineExtensionPlugin_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQQmlEngineExtensionPlugin_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQmlEngineExtensionPlugin_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QQmlEngineExtensionPlugin*)
Q_DECLARE_METATYPE(MyQQmlEngineExtensionPlugin*)

int QQmlEngineExtensionPlugin_QQmlEngineExtensionPlugin_QRegisterMetaType(){qRegisterMetaType<QQmlEngineExtensionPlugin*>(); return qRegisterMetaType<MyQQmlEngineExtensionPlugin*>();}

void* QQmlEngineExtensionPlugin_NewQQmlEngineExtensionPlugin3(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngineExtensionPlugin(static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlEngineExtensionPlugin(static_cast<QObject*>(parent));
	}
}

void QQmlEngineExtensionPlugin_InitializeEngine(void* ptr, void* engine, char* uri)
{
	static_cast<QQmlEngineExtensionPlugin*>(ptr)->initializeEngine(static_cast<QQmlEngine*>(engine), const_cast<const char*>(uri));
}

void QQmlEngineExtensionPlugin_InitializeEngineDefault(void* ptr, void* engine, char* uri)
{
		static_cast<QQmlEngineExtensionPlugin*>(ptr)->QQmlEngineExtensionPlugin::initializeEngine(static_cast<QQmlEngine*>(engine), const_cast<const char*>(uri));
}

void* QQmlEngineExtensionPlugin___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlEngineExtensionPlugin___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlEngineExtensionPlugin___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QQmlEngineExtensionPlugin___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlEngineExtensionPlugin___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQmlEngineExtensionPlugin___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QQmlEngineExtensionPlugin___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlEngineExtensionPlugin___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlEngineExtensionPlugin___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQmlEngineExtensionPlugin___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlEngineExtensionPlugin___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlEngineExtensionPlugin___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QQmlEngineExtensionPlugin_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQmlEngineExtensionPlugin*>(ptr)->QQmlEngineExtensionPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlEngineExtensionPlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlEngineExtensionPlugin*>(ptr)->QQmlEngineExtensionPlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlEngineExtensionPlugin_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQmlEngineExtensionPlugin*>(ptr)->QQmlEngineExtensionPlugin::customEvent(static_cast<QEvent*>(event));
}

void QQmlEngineExtensionPlugin_DeleteLaterDefault(void* ptr)
{
		static_cast<QQmlEngineExtensionPlugin*>(ptr)->QQmlEngineExtensionPlugin::deleteLater();
}

void QQmlEngineExtensionPlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlEngineExtensionPlugin*>(ptr)->QQmlEngineExtensionPlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QQmlEngineExtensionPlugin_EventDefault(void* ptr, void* e)
{
		return static_cast<QQmlEngineExtensionPlugin*>(ptr)->QQmlEngineExtensionPlugin::event(static_cast<QEvent*>(e));
}

char QQmlEngineExtensionPlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQmlEngineExtensionPlugin*>(ptr)->QQmlEngineExtensionPlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlEngineExtensionPlugin_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQmlEngineExtensionPlugin*>(ptr)->QQmlEngineExtensionPlugin::metaObject());
}

void QQmlEngineExtensionPlugin_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQmlEngineExtensionPlugin*>(ptr)->QQmlEngineExtensionPlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

Q_DECLARE_METATYPE(QQmlError)
Q_DECLARE_METATYPE(QQmlError*)
void* QQmlError_NewQQmlError()
{
	return new QQmlError();
}

void* QQmlError_NewQQmlError2(void* other)
{
	return new QQmlError(*static_cast<QQmlError*>(other));
}

int QQmlError_Column(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->column();
}

struct QtQml_PackedString QQmlError_Description(void* ptr)
{
	return ({ QByteArray* tff8552 = new QByteArray(static_cast<QQmlError*>(ptr)->description().toUtf8()); QtQml_PackedString { const_cast<char*>(tff8552->prepend("WHITESPACE").constData()+10), tff8552->size()-10, tff8552 }; });
}

char QQmlError_IsValid(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->isValid();
}

int QQmlError_Line(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->line();
}

void* QQmlError_Object(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->object();
}

void QQmlError_SetColumn(void* ptr, int column)
{
	static_cast<QQmlError*>(ptr)->setColumn(column);
}

void QQmlError_SetDescription(void* ptr, struct QtQml_PackedString description)
{
	static_cast<QQmlError*>(ptr)->setDescription(QString::fromUtf8(description.data, description.len));
}

void QQmlError_SetLine(void* ptr, int line)
{
	static_cast<QQmlError*>(ptr)->setLine(line);
}

void QQmlError_SetObject(void* ptr, void* object)
{
	static_cast<QQmlError*>(ptr)->setObject(static_cast<QObject*>(object));
}

void QQmlError_SetUrl(void* ptr, void* url)
{
	static_cast<QQmlError*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

struct QtQml_PackedString QQmlError_ToString(void* ptr)
{
	return ({ QByteArray* t77ab7b = new QByteArray(static_cast<QQmlError*>(ptr)->toString().toUtf8()); QtQml_PackedString { const_cast<char*>(t77ab7b->prepend("WHITESPACE").constData()+10), t77ab7b->size()-10, t77ab7b }; });
}

void* QQmlError_Url(void* ptr)
{
	return new QUrl(static_cast<QQmlError*>(ptr)->url());
}

class MyQQmlExpression: public QQmlExpression
{
public:
	MyQQmlExpression() : QQmlExpression() {QQmlExpression_QQmlExpression_QRegisterMetaType();};
	MyQQmlExpression(QQmlContext *ctxt, QObject *scope, const QString &expression, QObject *parent = Q_NULLPTR) : QQmlExpression(ctxt, scope, expression, parent) {QQmlExpression_QQmlExpression_QRegisterMetaType();};
	MyQQmlExpression(const QQmlScriptString &scri, QQmlContext *ctxt = Q_NULLPTR, QObject *scope = Q_NULLPTR, QObject *parent = Q_NULLPTR) : QQmlExpression(scri, ctxt, scope, parent) {QQmlExpression_QQmlExpression_QRegisterMetaType();};
	void Signal_ValueChanged() { callbackQQmlExpression_ValueChanged(this); };
	 ~MyQQmlExpression() { callbackQQmlExpression_DestroyQQmlExpression(this); };
	void childEvent(QChildEvent * event) { callbackQQmlExpression_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlExpression_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQmlExpression_CustomEvent(this, event); };
	void deleteLater() { callbackQQmlExpression_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQmlExpression_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlExpression_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQQmlExpression_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlExpression_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlExpression_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQQmlExpression_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQmlExpression_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QQmlExpression*)
Q_DECLARE_METATYPE(MyQQmlExpression*)

int QQmlExpression_QQmlExpression_QRegisterMetaType(){qRegisterMetaType<QQmlExpression*>(); return qRegisterMetaType<MyQQmlExpression*>();}

void* QQmlExpression_NewQQmlExpression()
{
	return new MyQQmlExpression();
}

void* QQmlExpression_NewQQmlExpression2(void* ctxt, void* scope, struct QtQml_PackedString expression, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QCameraImageCapture*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QDBusPendingCallWatcher*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QExtensionFactory*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QExtensionManager*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QGraphicsObject*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QGraphicsWidget*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QLayout*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QMediaPlaylist*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QMediaRecorder*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QOffscreenSurface*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QPaintDeviceWindow*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QPdfWriter*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QQuickItem*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QRadioData*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QRemoteObjectPendingCallWatcher*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QWidget*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QWindow*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	} else {
		return new MyQQmlExpression(static_cast<QQmlContext*>(ctxt), static_cast<QObject*>(scope), QString::fromUtf8(expression.data, expression.len), static_cast<QObject*>(parent));
	}
}

void* QQmlExpression_NewQQmlExpression3(void* scri, void* ctxt, void* scope, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QCameraImageCapture*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QDBusPendingCallWatcher*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QExtensionFactory*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QExtensionManager*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QGraphicsObject*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QGraphicsWidget*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QLayout*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QMediaPlaylist*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QMediaRecorder*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QOffscreenSurface*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QPaintDeviceWindow*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QPdfWriter*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QQuickItem*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QRadioData*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QRemoteObjectPendingCallWatcher*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QWidget*>(scope), static_cast<QObject*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(scope))) {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QWindow*>(scope), static_cast<QObject*>(parent));
	} else {
		return new MyQQmlExpression(*static_cast<QQmlScriptString*>(scri), static_cast<QQmlContext*>(ctxt), static_cast<QObject*>(scope), static_cast<QObject*>(parent));
	}
}

void QQmlExpression_ClearError(void* ptr)
{
	static_cast<QQmlExpression*>(ptr)->clearError();
}

int QQmlExpression_ColumnNumber(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->columnNumber();
}

void* QQmlExpression_Context(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->context();
}

void* QQmlExpression_Engine(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->engine();
}

void* QQmlExpression_Error(void* ptr)
{
	return new QQmlError(static_cast<QQmlExpression*>(ptr)->error());
}

void* QQmlExpression_Evaluate(void* ptr, char* valueIsUndefined)
{
	return new QVariant(static_cast<QQmlExpression*>(ptr)->evaluate(reinterpret_cast<bool*>(valueIsUndefined)));
}

struct QtQml_PackedString QQmlExpression_Expression(void* ptr)
{
	return ({ QByteArray* t1c3c3e = new QByteArray(static_cast<QQmlExpression*>(ptr)->expression().toUtf8()); QtQml_PackedString { const_cast<char*>(t1c3c3e->prepend("WHITESPACE").constData()+10), t1c3c3e->size()-10, t1c3c3e }; });
}

char QQmlExpression_HasError(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->hasError();
}

int QQmlExpression_LineNumber(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->lineNumber();
}

char QQmlExpression_NotifyOnValueChanged(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->notifyOnValueChanged();
}

void* QQmlExpression_ScopeObject(void* ptr)
{
	return static_cast<QQmlExpression*>(ptr)->scopeObject();
}

void QQmlExpression_SetExpression(void* ptr, struct QtQml_PackedString expression)
{
	static_cast<QQmlExpression*>(ptr)->setExpression(QString::fromUtf8(expression.data, expression.len));
}

void QQmlExpression_SetNotifyOnValueChanged(void* ptr, char notifyOnChange)
{
	static_cast<QQmlExpression*>(ptr)->setNotifyOnValueChanged(notifyOnChange != 0);
}

void QQmlExpression_SetSourceLocation(void* ptr, struct QtQml_PackedString url, int line, int column)
{
	static_cast<QQmlExpression*>(ptr)->setSourceLocation(QString::fromUtf8(url.data, url.len), line, column);
}

struct QtQml_PackedString QQmlExpression_SourceFile(void* ptr)
{
	return ({ QByteArray* t4b188b = new QByteArray(static_cast<QQmlExpression*>(ptr)->sourceFile().toUtf8()); QtQml_PackedString { const_cast<char*>(t4b188b->prepend("WHITESPACE").constData()+10), t4b188b->size()-10, t4b188b }; });
}

void QQmlExpression_ConnectValueChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QQmlExpression*>(ptr), static_cast<void (QQmlExpression::*)()>(&QQmlExpression::valueChanged), static_cast<MyQQmlExpression*>(ptr), static_cast<void (MyQQmlExpression::*)()>(&MyQQmlExpression::Signal_ValueChanged), static_cast<Qt::ConnectionType>(t));
}

void QQmlExpression_DisconnectValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlExpression*>(ptr), static_cast<void (QQmlExpression::*)()>(&QQmlExpression::valueChanged), static_cast<MyQQmlExpression*>(ptr), static_cast<void (MyQQmlExpression::*)()>(&MyQQmlExpression::Signal_ValueChanged));
}

void QQmlExpression_ValueChanged(void* ptr)
{
	static_cast<QQmlExpression*>(ptr)->valueChanged();
}

void QQmlExpression_DestroyQQmlExpression(void* ptr)
{
	static_cast<QQmlExpression*>(ptr)->~QQmlExpression();
}

void QQmlExpression_DestroyQQmlExpressionDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQmlExpression___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlExpression___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlExpression___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QQmlExpression___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlExpression___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQmlExpression___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QQmlExpression___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlExpression___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlExpression___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQmlExpression___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlExpression___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlExpression___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QQmlExpression_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQmlExpression*>(ptr)->QQmlExpression::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlExpression_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlExpression*>(ptr)->QQmlExpression::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlExpression_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQmlExpression*>(ptr)->QQmlExpression::customEvent(static_cast<QEvent*>(event));
}

void QQmlExpression_DeleteLaterDefault(void* ptr)
{
		static_cast<QQmlExpression*>(ptr)->QQmlExpression::deleteLater();
}

void QQmlExpression_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlExpression*>(ptr)->QQmlExpression::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QQmlExpression_EventDefault(void* ptr, void* e)
{
		return static_cast<QQmlExpression*>(ptr)->QQmlExpression::event(static_cast<QEvent*>(e));
}

char QQmlExpression_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQmlExpression*>(ptr)->QQmlExpression::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlExpression_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQmlExpression*>(ptr)->QQmlExpression::metaObject());
}

void QQmlExpression_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQmlExpression*>(ptr)->QQmlExpression::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQQmlFileSelector: public QQmlFileSelector
{
public:
	MyQQmlFileSelector(QQmlEngine *engine, QObject *parent = Q_NULLPTR) : QQmlFileSelector(engine, parent) {QQmlFileSelector_QQmlFileSelector_QRegisterMetaType();};
	 ~MyQQmlFileSelector() { callbackQQmlFileSelector_DestroyQQmlFileSelector(this); };
	void childEvent(QChildEvent * event) { callbackQQmlFileSelector_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlFileSelector_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQmlFileSelector_CustomEvent(this, event); };
	void deleteLater() { callbackQQmlFileSelector_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQmlFileSelector_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlFileSelector_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQQmlFileSelector_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlFileSelector_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlFileSelector_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQQmlFileSelector_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQmlFileSelector_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QQmlFileSelector*)
Q_DECLARE_METATYPE(MyQQmlFileSelector*)

int QQmlFileSelector_QQmlFileSelector_QRegisterMetaType(){qRegisterMetaType<QQmlFileSelector*>(); return qRegisterMetaType<MyQQmlFileSelector*>();}

void* QQmlFileSelector_NewQQmlFileSelector(void* engine, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
	}
}

void* QQmlFileSelector_QQmlFileSelector_Get(void* engine)
{
	return QQmlFileSelector::get(static_cast<QQmlEngine*>(engine));
}

void* QQmlFileSelector_Selector(void* ptr)
{
	return static_cast<QQmlFileSelector*>(ptr)->selector();
}

void QQmlFileSelector_SetExtraSelectors(void* ptr, struct QtQml_PackedString strin)
{
	static_cast<QQmlFileSelector*>(ptr)->setExtraSelectors(*(new QStringList(QString::fromUtf8(strin.data, strin.len).split("¡¦!", QString::SkipEmptyParts))));
}

void QQmlFileSelector_SetExtraSelectors2(void* ptr, struct QtQml_PackedString strin)
{
	static_cast<QQmlFileSelector*>(ptr)->setExtraSelectors(QString::fromUtf8(strin.data, strin.len).split("¡¦!", QString::SkipEmptyParts));
}

void QQmlFileSelector_SetSelector(void* ptr, void* selector)
{
	static_cast<QQmlFileSelector*>(ptr)->setSelector(static_cast<QFileSelector*>(selector));
}

void QQmlFileSelector_DestroyQQmlFileSelector(void* ptr)
{
	static_cast<QQmlFileSelector*>(ptr)->~QQmlFileSelector();
}

void QQmlFileSelector_DestroyQQmlFileSelectorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQmlFileSelector___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlFileSelector___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlFileSelector___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QQmlFileSelector___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlFileSelector___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQmlFileSelector___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QQmlFileSelector___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlFileSelector___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlFileSelector___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQmlFileSelector___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlFileSelector___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlFileSelector___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QQmlFileSelector_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlFileSelector_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlFileSelector_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::customEvent(static_cast<QEvent*>(event));
}

void QQmlFileSelector_DeleteLaterDefault(void* ptr)
{
		static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::deleteLater();
}

void QQmlFileSelector_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QQmlFileSelector_EventDefault(void* ptr, void* e)
{
		return static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::event(static_cast<QEvent*>(e));
}

char QQmlFileSelector_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlFileSelector_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::metaObject());
}

void QQmlFileSelector_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQmlFileSelector*>(ptr)->QQmlFileSelector::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQQmlImageProviderBase: public QQmlImageProviderBase
{
public:
	QQmlImageProviderBase::Flags flags() const { return static_cast<QQmlImageProviderBase::Flag>(callbackQQmlImageProviderBase_Flags(const_cast<void*>(static_cast<const void*>(this)))); };
	QQmlImageProviderBase::ImageType imageType() const { return static_cast<QQmlImageProviderBase::ImageType>(callbackQQmlImageProviderBase_ImageType(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QQmlImageProviderBase*)
Q_DECLARE_METATYPE(MyQQmlImageProviderBase*)

int QQmlImageProviderBase_QQmlImageProviderBase_QRegisterMetaType(){qRegisterMetaType<QQmlImageProviderBase*>(); return qRegisterMetaType<MyQQmlImageProviderBase*>();}

long long QQmlImageProviderBase_Flags(void* ptr)
{
	return static_cast<QQmlImageProviderBase*>(ptr)->flags();
}

long long QQmlImageProviderBase_ImageType(void* ptr)
{
	return static_cast<QQmlImageProviderBase*>(ptr)->imageType();
}

class MyQQmlIncubationController: public QQmlIncubationController
{
public:
	MyQQmlIncubationController() : QQmlIncubationController() {QQmlIncubationController_QQmlIncubationController_QRegisterMetaType();};
	void incubatingObjectCountChanged(int incubatingObjectCount) { callbackQQmlIncubationController_IncubatingObjectCountChanged(this, incubatingObjectCount); };
};

Q_DECLARE_METATYPE(QQmlIncubationController*)
Q_DECLARE_METATYPE(MyQQmlIncubationController*)

int QQmlIncubationController_QQmlIncubationController_QRegisterMetaType(){qRegisterMetaType<QQmlIncubationController*>(); return qRegisterMetaType<MyQQmlIncubationController*>();}

void* QQmlIncubationController_NewQQmlIncubationController2()
{
	return new MyQQmlIncubationController();
}

void* QQmlIncubationController_Engine(void* ptr)
{
	return static_cast<QQmlIncubationController*>(ptr)->engine();
}

void QQmlIncubationController_IncubateFor(void* ptr, int msecs)
{
	static_cast<QQmlIncubationController*>(ptr)->incubateFor(msecs);
}

int QQmlIncubationController_IncubatingObjectCount(void* ptr)
{
	return static_cast<QQmlIncubationController*>(ptr)->incubatingObjectCount();
}

void QQmlIncubationController_IncubatingObjectCountChanged(void* ptr, int incubatingObjectCount)
{
	static_cast<QQmlIncubationController*>(ptr)->incubatingObjectCountChanged(incubatingObjectCount);
}

void QQmlIncubationController_IncubatingObjectCountChangedDefault(void* ptr, int incubatingObjectCount)
{
		static_cast<QQmlIncubationController*>(ptr)->QQmlIncubationController::incubatingObjectCountChanged(incubatingObjectCount);
}

class MyQQmlIncubator: public QQmlIncubator
{
public:
	MyQQmlIncubator(QQmlIncubator::IncubationMode mode = Asynchronous) : QQmlIncubator(mode) {QQmlIncubator_QQmlIncubator_QRegisterMetaType();};
	void setInitialState(QObject * object) { callbackQQmlIncubator_SetInitialState(this, object); };
	void statusChanged(QQmlIncubator::Status status) { callbackQQmlIncubator_StatusChanged(this, status); };
};

Q_DECLARE_METATYPE(QQmlIncubator*)
Q_DECLARE_METATYPE(MyQQmlIncubator*)

int QQmlIncubator_QQmlIncubator_QRegisterMetaType(){qRegisterMetaType<QQmlIncubator*>(); return qRegisterMetaType<MyQQmlIncubator*>();}

void* QQmlIncubator_NewQQmlIncubator2(long long mode)
{
	return new MyQQmlIncubator(static_cast<QQmlIncubator::IncubationMode>(mode));
}

void QQmlIncubator_Clear(void* ptr)
{
	static_cast<QQmlIncubator*>(ptr)->clear();
}

struct QtQml_PackedList QQmlIncubator_Errors(void* ptr)
{
	return ({ QList<QQmlError>* tmpValuea88c54 = new QList<QQmlError>(static_cast<QQmlIncubator*>(ptr)->errors()); QtQml_PackedList { tmpValuea88c54, tmpValuea88c54->size() }; });
}

void QQmlIncubator_ForceCompletion(void* ptr)
{
	static_cast<QQmlIncubator*>(ptr)->forceCompletion();
}

long long QQmlIncubator_IncubationMode(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->incubationMode();
}

char QQmlIncubator_IsError(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->isError();
}

char QQmlIncubator_IsLoading(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->isLoading();
}

char QQmlIncubator_IsNull(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->isNull();
}

char QQmlIncubator_IsReady(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->isReady();
}

void* QQmlIncubator_Object(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->object();
}

void QQmlIncubator_SetInitialState(void* ptr, void* object)
{
	static_cast<QQmlIncubator*>(ptr)->setInitialState(static_cast<QObject*>(object));
}

void QQmlIncubator_SetInitialStateDefault(void* ptr, void* object)
{
		static_cast<QQmlIncubator*>(ptr)->QQmlIncubator::setInitialState(static_cast<QObject*>(object));
}

long long QQmlIncubator_Status(void* ptr)
{
	return static_cast<QQmlIncubator*>(ptr)->status();
}

void QQmlIncubator_StatusChanged(void* ptr, long long status)
{
	static_cast<QQmlIncubator*>(ptr)->statusChanged(static_cast<QQmlIncubator::Status>(status));
}

void QQmlIncubator_StatusChangedDefault(void* ptr, long long status)
{
		static_cast<QQmlIncubator*>(ptr)->QQmlIncubator::statusChanged(static_cast<QQmlIncubator::Status>(status));
}

void* QQmlIncubator___errors_atList(void* ptr, int i)
{
	return new QQmlError(({QQmlError tmp = static_cast<QList<QQmlError>*>(ptr)->at(i); if (i == static_cast<QList<QQmlError>*>(ptr)->size()-1) { static_cast<QList<QQmlError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlIncubator___errors_setList(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQmlIncubator___errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>();
}

void* QQmlIncubator___setInitialProperties_initialProperties_atList(void* ptr, struct QtQml_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QQmlIncubator___setInitialProperties_initialProperties_setList(void* ptr, struct QtQml_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QQmlIncubator___setInitialProperties_initialProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtQml_PackedList QQmlIncubator___setInitialProperties_initialProperties_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtQml_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtQml_PackedString QQmlIncubator_____setInitialProperties_initialProperties_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtQml_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QQmlIncubator_____setInitialProperties_initialProperties_keyList_setList(void* ptr, struct QtQml_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QQmlIncubator_____setInitialProperties_initialProperties_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

Q_DECLARE_METATYPE(QQmlListReference*)
void* QQmlListReference_NewQQmlListReference()
{
	return new QQmlListReference();
}

void* QQmlListReference_NewQQmlListReference2(void* object, char* property, void* engine)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QCameraImageCapture*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QDBusPendingCallWatcher*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QExtensionFactory*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QExtensionManager*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QGraphicsObject*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QGraphicsWidget*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QLayout*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QMediaPlaylist*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QMediaRecorder*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QOffscreenSurface*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QPaintDeviceWindow*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QPdfWriter*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QQuickItem*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QRadioData*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QRemoteObjectPendingCallWatcher*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QWidget*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(object))) {
		return new QQmlListReference(static_cast<QWindow*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	} else {
		return new QQmlListReference(static_cast<QObject*>(object), const_cast<const char*>(property), static_cast<QQmlEngine*>(engine));
	}
}

char QQmlListReference_Append(void* ptr, void* object)
{
	return static_cast<QQmlListReference*>(ptr)->append(static_cast<QObject*>(object));
}

void* QQmlListReference_At(void* ptr, int index)
{
	return static_cast<QQmlListReference*>(ptr)->at(index);
}

char QQmlListReference_CanAppend(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->canAppend();
}

char QQmlListReference_CanAt(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->canAt();
}

char QQmlListReference_CanClear(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->canClear();
}

char QQmlListReference_CanCount(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->canCount();
}

char QQmlListReference_CanRemoveLast(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->canRemoveLast();
}

char QQmlListReference_CanReplace(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->canReplace();
}

char QQmlListReference_Clear(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->clear();
}

int QQmlListReference_Count(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->count();
}

char QQmlListReference_IsManipulable(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->isManipulable();
}

char QQmlListReference_IsReadable(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->isReadable();
}

char QQmlListReference_IsValid(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->isValid();
}

void* QQmlListReference_ListElementType(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QQmlListReference*>(ptr)->listElementType());
}

void* QQmlListReference_Object(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->object();
}

char QQmlListReference_RemoveLast(void* ptr)
{
	return static_cast<QQmlListReference*>(ptr)->removeLast();
}

char QQmlListReference_Replace(void* ptr, int index, void* object)
{
	return static_cast<QQmlListReference*>(ptr)->replace(index, static_cast<QObject*>(object));
}

class MyQQmlNetworkAccessManagerFactory: public QQmlNetworkAccessManagerFactory
{
public:
	MyQQmlNetworkAccessManagerFactory() : QQmlNetworkAccessManagerFactory() {QQmlNetworkAccessManagerFactory_QQmlNetworkAccessManagerFactory_QRegisterMetaType();};
	QNetworkAccessManager * create(QObject * parent) { return static_cast<QNetworkAccessManager*>(callbackQQmlNetworkAccessManagerFactory_Create(this, parent)); };
	 ~MyQQmlNetworkAccessManagerFactory() { callbackQQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(this); };
};

Q_DECLARE_METATYPE(QQmlNetworkAccessManagerFactory*)
Q_DECLARE_METATYPE(MyQQmlNetworkAccessManagerFactory*)

int QQmlNetworkAccessManagerFactory_QQmlNetworkAccessManagerFactory_QRegisterMetaType(){qRegisterMetaType<QQmlNetworkAccessManagerFactory*>(); return qRegisterMetaType<MyQQmlNetworkAccessManagerFactory*>();}

void* QQmlNetworkAccessManagerFactory_Create(void* ptr, void* parent)
{
	return static_cast<QQmlNetworkAccessManagerFactory*>(ptr)->create(static_cast<QObject*>(parent));
}

void QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(void* ptr)
{
	static_cast<QQmlNetworkAccessManagerFactory*>(ptr)->~QQmlNetworkAccessManagerFactory();
}

void QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactoryDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQmlNetworkAccessManagerFactory_NewQQmlNetworkAccessManagerFactory()
{
	return new MyQQmlNetworkAccessManagerFactory();
}

class MyQQmlParserStatus: public QQmlParserStatus
{
public:
	void classBegin() { callbackQQmlParserStatus_ClassBegin(this); };
	void componentComplete() { callbackQQmlParserStatus_ComponentComplete(this); };
};

Q_DECLARE_METATYPE(QQmlParserStatus*)
Q_DECLARE_METATYPE(MyQQmlParserStatus*)

int QQmlParserStatus_QQmlParserStatus_QRegisterMetaType(){qRegisterMetaType<QQmlParserStatus*>(); return qRegisterMetaType<MyQQmlParserStatus*>();}

void QQmlParserStatus_ClassBegin(void* ptr)
{
	static_cast<QQmlParserStatus*>(ptr)->classBegin();
}

void QQmlParserStatus_ComponentComplete(void* ptr)
{
	static_cast<QQmlParserStatus*>(ptr)->componentComplete();
}

Q_DECLARE_METATYPE(QQmlProperty)
Q_DECLARE_METATYPE(QQmlProperty*)
void* QQmlProperty_NewQQmlProperty()
{
	return new QQmlProperty();
}

void* QQmlProperty_NewQQmlProperty2(void* obj)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QCameraImageCapture*>(obj));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QDBusPendingCallWatcher*>(obj));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionFactory*>(obj));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionManager*>(obj));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsObject*>(obj));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsWidget*>(obj));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QLayout*>(obj));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaPlaylist*>(obj));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaRecorder*>(obj));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QOffscreenSurface*>(obj));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPaintDeviceWindow*>(obj));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPdfWriter*>(obj));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QQuickItem*>(obj));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRadioData*>(obj));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRemoteObjectPendingCallWatcher*>(obj));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWidget*>(obj));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWindow*>(obj));
	} else {
		return new QQmlProperty(static_cast<QObject*>(obj));
	}
}

void* QQmlProperty_NewQQmlProperty3(void* obj, void* ctxt)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QCameraImageCapture*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QDBusPendingCallWatcher*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionFactory*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionManager*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsObject*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsWidget*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QLayout*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaPlaylist*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaRecorder*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QOffscreenSurface*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPaintDeviceWindow*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPdfWriter*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QQuickItem*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRadioData*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRemoteObjectPendingCallWatcher*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWidget*>(obj), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWindow*>(obj), static_cast<QQmlContext*>(ctxt));
	} else {
		return new QQmlProperty(static_cast<QObject*>(obj), static_cast<QQmlContext*>(ctxt));
	}
}

void* QQmlProperty_NewQQmlProperty4(void* obj, void* engine)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QCameraImageCapture*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QDBusPendingCallWatcher*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionFactory*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionManager*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsObject*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsWidget*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QLayout*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaPlaylist*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaRecorder*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QOffscreenSurface*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPaintDeviceWindow*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPdfWriter*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QQuickItem*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRadioData*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRemoteObjectPendingCallWatcher*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWidget*>(obj), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWindow*>(obj), static_cast<QQmlEngine*>(engine));
	} else {
		return new QQmlProperty(static_cast<QObject*>(obj), static_cast<QQmlEngine*>(engine));
	}
}

void* QQmlProperty_NewQQmlProperty5(void* obj, struct QtQml_PackedString name)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QCameraImageCapture*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QDBusPendingCallWatcher*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionFactory*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionManager*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsObject*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsWidget*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QLayout*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaPlaylist*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaRecorder*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QOffscreenSurface*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPaintDeviceWindow*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPdfWriter*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QQuickItem*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRadioData*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRemoteObjectPendingCallWatcher*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWidget*>(obj), QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWindow*>(obj), QString::fromUtf8(name.data, name.len));
	} else {
		return new QQmlProperty(static_cast<QObject*>(obj), QString::fromUtf8(name.data, name.len));
	}
}

void* QQmlProperty_NewQQmlProperty6(void* obj, struct QtQml_PackedString name, void* ctxt)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QCameraImageCapture*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QDBusPendingCallWatcher*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionFactory*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionManager*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsObject*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsWidget*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QLayout*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaPlaylist*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaRecorder*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QOffscreenSurface*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPaintDeviceWindow*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPdfWriter*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QQuickItem*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRadioData*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRemoteObjectPendingCallWatcher*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWidget*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWindow*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	} else {
		return new QQmlProperty(static_cast<QObject*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt));
	}
}

void* QQmlProperty_NewQQmlProperty7(void* obj, struct QtQml_PackedString name, void* engine)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QCameraImageCapture*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QDBusPendingCallWatcher*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionFactory*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QExtensionManager*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsObject*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QGraphicsWidget*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QLayout*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaPlaylist*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QMediaRecorder*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QOffscreenSurface*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPaintDeviceWindow*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QPdfWriter*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QQuickItem*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRadioData*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QRemoteObjectPendingCallWatcher*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWidget*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(obj))) {
		return new QQmlProperty(static_cast<QWindow*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	} else {
		return new QQmlProperty(static_cast<QObject*>(obj), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine));
	}
}

void* QQmlProperty_NewQQmlProperty8(void* other)
{
	return new QQmlProperty(*static_cast<QQmlProperty*>(other));
}

char QQmlProperty_ConnectNotifySignal(void* ptr, void* dest, char* slot)
{
	return static_cast<QQmlProperty*>(ptr)->connectNotifySignal(static_cast<QObject*>(dest), const_cast<const char*>(slot));
}

char QQmlProperty_ConnectNotifySignal2(void* ptr, void* dest, int method)
{
	return static_cast<QQmlProperty*>(ptr)->connectNotifySignal(static_cast<QObject*>(dest), method);
}

char QQmlProperty_HasNotifySignal(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->hasNotifySignal();
}

int QQmlProperty_Index(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->index();
}

char QQmlProperty_IsDesignable(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isDesignable();
}

char QQmlProperty_IsProperty(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isProperty();
}

char QQmlProperty_IsResettable(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isResettable();
}

char QQmlProperty_IsSignalProperty(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isSignalProperty();
}

char QQmlProperty_IsValid(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isValid();
}

char QQmlProperty_IsWritable(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->isWritable();
}

void* QQmlProperty_Method(void* ptr)
{
	return new QMetaMethod(static_cast<QQmlProperty*>(ptr)->method());
}

struct QtQml_PackedString QQmlProperty_Name(void* ptr)
{
	return ({ QByteArray* t576900 = new QByteArray(static_cast<QQmlProperty*>(ptr)->name().toUtf8()); QtQml_PackedString { const_cast<char*>(t576900->prepend("WHITESPACE").constData()+10), t576900->size()-10, t576900 }; });
}

char QQmlProperty_NeedsNotifySignal(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->needsNotifySignal();
}

void* QQmlProperty_Object(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->object();
}

int QQmlProperty_PropertyType(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->propertyType();
}

long long QQmlProperty_PropertyTypeCategory(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->propertyTypeCategory();
}

struct QtQml_PackedString QQmlProperty_PropertyTypeName(void* ptr)
{
	return QtQml_PackedString { const_cast<char*>(static_cast<QQmlProperty*>(ptr)->propertyTypeName()), -1, NULL };
}

void* QQmlProperty_Read(void* ptr)
{
	return new QVariant(static_cast<QQmlProperty*>(ptr)->read());
}

void* QQmlProperty_QQmlProperty_Read2(void* object, struct QtQml_PackedString name)
{
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString::fromUtf8(name.data, name.len)));
}

void* QQmlProperty_QQmlProperty_Read3(void* object, struct QtQml_PackedString name, void* ctxt)
{
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString::fromUtf8(name.data, name.len), static_cast<QQmlContext*>(ctxt)));
}

void* QQmlProperty_QQmlProperty_Read4(void* object, struct QtQml_PackedString name, void* engine)
{
	return new QVariant(QQmlProperty::read(static_cast<QObject*>(object), QString::fromUtf8(name.data, name.len), static_cast<QQmlEngine*>(engine)));
}

char QQmlProperty_Reset(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->reset();
}

long long QQmlProperty_Type(void* ptr)
{
	return static_cast<QQmlProperty*>(ptr)->type();
}

char QQmlProperty_Write(void* ptr, void* value)
{
	return static_cast<QQmlProperty*>(ptr)->write(*static_cast<QVariant*>(value));
}

char QQmlProperty_QQmlProperty_Write2(void* object, struct QtQml_PackedString name, void* value)
{
	return QQmlProperty::write(static_cast<QObject*>(object), QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value));
}

char QQmlProperty_QQmlProperty_Write3(void* object, struct QtQml_PackedString name, void* value, void* ctxt)
{
	return QQmlProperty::write(static_cast<QObject*>(object), QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value), static_cast<QQmlContext*>(ctxt));
}

char QQmlProperty_QQmlProperty_Write4(void* object, struct QtQml_PackedString name, void* value, void* engine)
{
	return QQmlProperty::write(static_cast<QObject*>(object), QString::fromUtf8(name.data, name.len), *static_cast<QVariant*>(value), static_cast<QQmlEngine*>(engine));
}

class MyQQmlPropertyMap: public QQmlPropertyMap
{
public:
	MyQQmlPropertyMap(QObject *parent = Q_NULLPTR) : QQmlPropertyMap(parent) {QQmlPropertyMap_QQmlPropertyMap_QRegisterMetaType();};
	QVariant updateValue(const QString & key, const QVariant & input) { QByteArray* ta62f22 = new QByteArray(key.toUtf8()); QtQml_PackedString keyPacked = { const_cast<char*>(ta62f22->prepend("WHITESPACE").constData()+10), ta62f22->size()-10, ta62f22 };return *static_cast<QVariant*>(callbackQQmlPropertyMap_UpdateValue(this, keyPacked, const_cast<QVariant*>(&input))); };
	void Signal_ValueChanged(const QString & key, const QVariant & value) { QByteArray* ta62f22 = new QByteArray(key.toUtf8()); QtQml_PackedString keyPacked = { const_cast<char*>(ta62f22->prepend("WHITESPACE").constData()+10), ta62f22->size()-10, ta62f22 };callbackQQmlPropertyMap_ValueChanged(this, keyPacked, const_cast<QVariant*>(&value)); };
	 ~MyQQmlPropertyMap() { callbackQQmlPropertyMap_DestroyQQmlPropertyMap(this); };
	void childEvent(QChildEvent * event) { callbackQQmlPropertyMap_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQQmlPropertyMap_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQQmlPropertyMap_CustomEvent(this, event); };
	void deleteLater() { callbackQQmlPropertyMap_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQQmlPropertyMap_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQQmlPropertyMap_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQQmlPropertyMap_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQQmlPropertyMap_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQQmlPropertyMap_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQQmlPropertyMap_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQQmlPropertyMap_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QQmlPropertyMap*)
Q_DECLARE_METATYPE(MyQQmlPropertyMap*)

int QQmlPropertyMap_QQmlPropertyMap_QRegisterMetaType(){qRegisterMetaType<QQmlPropertyMap*>(); return qRegisterMetaType<MyQQmlPropertyMap*>();}

void* QQmlPropertyMap_NewQQmlPropertyMap(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlPropertyMap(static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlPropertyMap(static_cast<QObject*>(parent));
	}
}

void QQmlPropertyMap_Clear(void* ptr, struct QtQml_PackedString key)
{
	static_cast<QQmlPropertyMap*>(ptr)->clear(QString::fromUtf8(key.data, key.len));
}

char QQmlPropertyMap_Contains(void* ptr, struct QtQml_PackedString key)
{
	return static_cast<QQmlPropertyMap*>(ptr)->contains(QString::fromUtf8(key.data, key.len));
}

int QQmlPropertyMap_Count(void* ptr)
{
	return static_cast<QQmlPropertyMap*>(ptr)->count();
}

void QQmlPropertyMap_Insert(void* ptr, struct QtQml_PackedString key, void* value)
{
	static_cast<QQmlPropertyMap*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(value));
}

char QQmlPropertyMap_IsEmpty(void* ptr)
{
	return static_cast<QQmlPropertyMap*>(ptr)->isEmpty();
}

struct QtQml_PackedString QQmlPropertyMap_Keys(void* ptr)
{
	return ({ QByteArray* t4c814d = new QByteArray(static_cast<QQmlPropertyMap*>(ptr)->keys().join("¡¦!").toUtf8()); QtQml_PackedString { const_cast<char*>(t4c814d->prepend("WHITESPACE").constData()+10), t4c814d->size()-10, t4c814d }; });
}

int QQmlPropertyMap_Size(void* ptr)
{
	return static_cast<QQmlPropertyMap*>(ptr)->size();
}

void* QQmlPropertyMap_UpdateValue(void* ptr, struct QtQml_PackedString key, void* input)
{
	return new QVariant(static_cast<QQmlPropertyMap*>(ptr)->updateValue(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(input)));
}

void* QQmlPropertyMap_UpdateValueDefault(void* ptr, struct QtQml_PackedString key, void* input)
{
		return new QVariant(static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::updateValue(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(input)));
}

void* QQmlPropertyMap_Value(void* ptr, struct QtQml_PackedString key)
{
	return new QVariant(static_cast<QQmlPropertyMap*>(ptr)->value(QString::fromUtf8(key.data, key.len)));
}

void QQmlPropertyMap_ConnectValueChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QQmlPropertyMap*>(ptr), static_cast<void (QQmlPropertyMap::*)(const QString &, const QVariant &)>(&QQmlPropertyMap::valueChanged), static_cast<MyQQmlPropertyMap*>(ptr), static_cast<void (MyQQmlPropertyMap::*)(const QString &, const QVariant &)>(&MyQQmlPropertyMap::Signal_ValueChanged), static_cast<Qt::ConnectionType>(t));
}

void QQmlPropertyMap_DisconnectValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QQmlPropertyMap*>(ptr), static_cast<void (QQmlPropertyMap::*)(const QString &, const QVariant &)>(&QQmlPropertyMap::valueChanged), static_cast<MyQQmlPropertyMap*>(ptr), static_cast<void (MyQQmlPropertyMap::*)(const QString &, const QVariant &)>(&MyQQmlPropertyMap::Signal_ValueChanged));
}

void QQmlPropertyMap_ValueChanged(void* ptr, struct QtQml_PackedString key, void* value)
{
	static_cast<QQmlPropertyMap*>(ptr)->valueChanged(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(value));
}

void QQmlPropertyMap_DestroyQQmlPropertyMap(void* ptr)
{
	static_cast<QQmlPropertyMap*>(ptr)->~QQmlPropertyMap();
}

void QQmlPropertyMap_DestroyQQmlPropertyMapDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QQmlPropertyMap___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlPropertyMap___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlPropertyMap___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QQmlPropertyMap___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlPropertyMap___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QQmlPropertyMap___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QQmlPropertyMap___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlPropertyMap___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlPropertyMap___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QQmlPropertyMap___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QQmlPropertyMap___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QQmlPropertyMap___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QQmlPropertyMap_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlPropertyMap_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QQmlPropertyMap_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::customEvent(static_cast<QEvent*>(event));
}

void QQmlPropertyMap_DeleteLaterDefault(void* ptr)
{
		static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::deleteLater();
}

void QQmlPropertyMap_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QQmlPropertyMap_EventDefault(void* ptr, void* e)
{
		return static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::event(static_cast<QEvent*>(e));
}

char QQmlPropertyMap_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QQmlPropertyMap_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::metaObject());
}

void QQmlPropertyMap_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QQmlPropertyMap*>(ptr)->QQmlPropertyMap::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQQmlPropertyValueSource: public QQmlPropertyValueSource
{
public:
	MyQQmlPropertyValueSource() : QQmlPropertyValueSource() {QQmlPropertyValueSource_QQmlPropertyValueSource_QRegisterMetaType();};
	void setTarget(const QQmlProperty & property) { callbackQQmlPropertyValueSource_SetTarget(this, const_cast<QQmlProperty*>(&property)); };
	 ~MyQQmlPropertyValueSource() { callbackQQmlPropertyValueSource_DestroyQQmlPropertyValueSource(this); };
};

Q_DECLARE_METATYPE(QQmlPropertyValueSource*)
Q_DECLARE_METATYPE(MyQQmlPropertyValueSource*)

int QQmlPropertyValueSource_QQmlPropertyValueSource_QRegisterMetaType(){qRegisterMetaType<QQmlPropertyValueSource*>(); return qRegisterMetaType<MyQQmlPropertyValueSource*>();}

void* QQmlPropertyValueSource_NewQQmlPropertyValueSource()
{
	return new MyQQmlPropertyValueSource();
}

void QQmlPropertyValueSource_SetTarget(void* ptr, void* property)
{
	static_cast<QQmlPropertyValueSource*>(ptr)->setTarget(*static_cast<QQmlProperty*>(property));
}

void QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(void* ptr)
{
	static_cast<QQmlPropertyValueSource*>(ptr)->~QQmlPropertyValueSource();
}

void QQmlPropertyValueSource_DestroyQQmlPropertyValueSourceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

Q_DECLARE_METATYPE(QQmlScriptString*)
void* QQmlScriptString_NewQQmlScriptString()
{
	return new QQmlScriptString();
}

void* QQmlScriptString_NewQQmlScriptString2(void* other)
{
	return new QQmlScriptString(*static_cast<QQmlScriptString*>(other));
}

char QQmlScriptString_BooleanLiteral(void* ptr, char* ok)
{
	return static_cast<QQmlScriptString*>(ptr)->booleanLiteral(reinterpret_cast<bool*>(ok));
}

char QQmlScriptString_IsEmpty(void* ptr)
{
	return static_cast<QQmlScriptString*>(ptr)->isEmpty();
}

char QQmlScriptString_IsNullLiteral(void* ptr)
{
	return static_cast<QQmlScriptString*>(ptr)->isNullLiteral();
}

char QQmlScriptString_IsUndefinedLiteral(void* ptr)
{
	return static_cast<QQmlScriptString*>(ptr)->isUndefinedLiteral();
}

double QQmlScriptString_NumberLiteral(void* ptr, char* ok)
{
	return static_cast<QQmlScriptString*>(ptr)->numberLiteral(reinterpret_cast<bool*>(ok));
}

struct QtQml_PackedString QQmlScriptString_StringLiteral(void* ptr)
{
	return ({ QByteArray* t77944a = new QByteArray(static_cast<QQmlScriptString*>(ptr)->stringLiteral().toUtf8()); QtQml_PackedString { const_cast<char*>(t77944a->prepend("WHITESPACE").constData()+10), t77944a->size()-10, t77944a }; });
}

