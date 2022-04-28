// +build !minimal

#define protected public
#define private public

#include "location.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QDateTime>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGeoCoordinate>
#include <QGeoManeuver>
#include <QGeoRectangle>
#include <QGeoRoute>
#include <QGeoRouteLeg>
#include <QGeoRouteReply>
#include <QGeoRouteRequest>
#include <QGeoRouteSegment>
#include <QGeoRoutingManager>
#include <QGeoRoutingManagerEngine>
#include <QGeoServiceProvider>
#include <QGeoServiceProviderFactory>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QLocale>
#include <QMap>
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
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QWidget>
#include <QWindow>

Q_DECLARE_METATYPE(QGeoManeuver)
Q_DECLARE_METATYPE(QGeoManeuver*)
void* QGeoManeuver_NewQGeoManeuver()
{
	return new QGeoManeuver();
}

void* QGeoManeuver_NewQGeoManeuver2(void* other)
{
	return new QGeoManeuver(*static_cast<QGeoManeuver*>(other));
}

long long QGeoManeuver_Direction(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->direction();
}

double QGeoManeuver_DistanceToNextInstruction(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->distanceToNextInstruction();
}

struct QtLocation_PackedList QGeoManeuver_ExtendedAttributes(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValuea18eb4 = new QMap<QString, QVariant>(static_cast<QGeoManeuver*>(ptr)->extendedAttributes()); QtLocation_PackedList { tmpValuea18eb4, tmpValuea18eb4->size() }; });
}

struct QtLocation_PackedString QGeoManeuver_InstructionText(void* ptr)
{
	return ({ QByteArray* t4cc4b3 = new QByteArray(static_cast<QGeoManeuver*>(ptr)->instructionText().toUtf8()); QtLocation_PackedString { const_cast<char*>(t4cc4b3->prepend("WHITESPACE").constData()+10), t4cc4b3->size()-10, t4cc4b3 }; });
}

char QGeoManeuver_IsValid(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->isValid();
}

void* QGeoManeuver_Position(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoManeuver*>(ptr)->position());
}

void QGeoManeuver_SetDirection(void* ptr, long long direction)
{
	static_cast<QGeoManeuver*>(ptr)->setDirection(static_cast<QGeoManeuver::InstructionDirection>(direction));
}

void QGeoManeuver_SetDistanceToNextInstruction(void* ptr, double distance)
{
	static_cast<QGeoManeuver*>(ptr)->setDistanceToNextInstruction(distance);
}

void QGeoManeuver_SetExtendedAttributes(void* ptr, void* extendedAttributes)
{
	static_cast<QGeoManeuver*>(ptr)->setExtendedAttributes(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(extendedAttributes); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void QGeoManeuver_SetInstructionText(void* ptr, struct QtLocation_PackedString instructionText)
{
	static_cast<QGeoManeuver*>(ptr)->setInstructionText(QString::fromUtf8(instructionText.data, instructionText.len));
}

void QGeoManeuver_SetPosition(void* ptr, void* position)
{
	static_cast<QGeoManeuver*>(ptr)->setPosition(*static_cast<QGeoCoordinate*>(position));
}

void QGeoManeuver_SetTimeToNextInstruction(void* ptr, int secs)
{
	static_cast<QGeoManeuver*>(ptr)->setTimeToNextInstruction(secs);
}

void QGeoManeuver_SetWaypoint(void* ptr, void* coordinate)
{
	static_cast<QGeoManeuver*>(ptr)->setWaypoint(*static_cast<QGeoCoordinate*>(coordinate));
}

int QGeoManeuver_TimeToNextInstruction(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->timeToNextInstruction();
}

void* QGeoManeuver_Waypoint(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoManeuver*>(ptr)->waypoint());
}

void QGeoManeuver_DestroyQGeoManeuver(void* ptr)
{
	static_cast<QGeoManeuver*>(ptr)->~QGeoManeuver();
}

void* QGeoManeuver___extendedAttributes_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoManeuver___extendedAttributes_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoManeuver___extendedAttributes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoManeuver___extendedAttributes_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoManeuver___setExtendedAttributes_extendedAttributes_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoManeuver___setExtendedAttributes_extendedAttributes_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoManeuver___setExtendedAttributes_extendedAttributes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoManeuver___setExtendedAttributes_extendedAttributes_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtLocation_PackedString QGeoManeuver_____extendedAttributes_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoManeuver_____extendedAttributes_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoManeuver_____extendedAttributes_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtLocation_PackedString QGeoManeuver_____setExtendedAttributes_extendedAttributes_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoManeuver_____setExtendedAttributes_extendedAttributes_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoManeuver_____setExtendedAttributes_extendedAttributes_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

Q_DECLARE_METATYPE(QGeoRoute)
Q_DECLARE_METATYPE(QGeoRoute*)
void* QGeoRoute_NewQGeoRoute()
{
	return new QGeoRoute();
}

void* QGeoRoute_NewQGeoRoute2(void* other)
{
	return new QGeoRoute(*static_cast<QGeoRoute*>(other));
}

void* QGeoRoute_Bounds(void* ptr)
{
	return new QGeoRectangle(static_cast<QGeoRoute*>(ptr)->bounds());
}

double QGeoRoute_Distance(void* ptr)
{
	return static_cast<QGeoRoute*>(ptr)->distance();
}

struct QtLocation_PackedList QGeoRoute_ExtendedAttributes(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue7e83bf = new QMap<QString, QVariant>(static_cast<QGeoRoute*>(ptr)->extendedAttributes()); QtLocation_PackedList { tmpValue7e83bf, tmpValue7e83bf->size() }; });
}

void* QGeoRoute_FirstRouteSegment(void* ptr)
{
	return new QGeoRouteSegment(static_cast<QGeoRoute*>(ptr)->firstRouteSegment());
}

struct QtLocation_PackedList QGeoRoute_Path(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValue38ece8 = new QList<QGeoCoordinate>(static_cast<QGeoRoute*>(ptr)->path()); QtLocation_PackedList { tmpValue38ece8, tmpValue38ece8->size() }; });
}

void* QGeoRoute_Request(void* ptr)
{
	return new QGeoRouteRequest(static_cast<QGeoRoute*>(ptr)->request());
}

struct QtLocation_PackedString QGeoRoute_RouteId(void* ptr)
{
	return ({ QByteArray* t7492be = new QByteArray(static_cast<QGeoRoute*>(ptr)->routeId().toUtf8()); QtLocation_PackedString { const_cast<char*>(t7492be->prepend("WHITESPACE").constData()+10), t7492be->size()-10, t7492be }; });
}

struct QtLocation_PackedList QGeoRoute_RouteLegs(void* ptr)
{
	return ({ QList<QGeoRouteLeg>* tmpValue7b4e55 = new QList<QGeoRouteLeg>(static_cast<QGeoRoute*>(ptr)->routeLegs()); QtLocation_PackedList { tmpValue7b4e55, tmpValue7b4e55->size() }; });
}

void QGeoRoute_SetBounds(void* ptr, void* bounds)
{
	static_cast<QGeoRoute*>(ptr)->setBounds(*static_cast<QGeoRectangle*>(bounds));
}

void QGeoRoute_SetDistance(void* ptr, double distance)
{
	static_cast<QGeoRoute*>(ptr)->setDistance(distance);
}

void QGeoRoute_SetExtendedAttributes(void* ptr, void* extendedAttributes)
{
	static_cast<QGeoRoute*>(ptr)->setExtendedAttributes(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(extendedAttributes); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void QGeoRoute_SetFirstRouteSegment(void* ptr, void* routeSegment)
{
	static_cast<QGeoRoute*>(ptr)->setFirstRouteSegment(*static_cast<QGeoRouteSegment*>(routeSegment));
}

void QGeoRoute_SetPath(void* ptr, void* path)
{
	static_cast<QGeoRoute*>(ptr)->setPath(*static_cast<QList<QGeoCoordinate>*>(path));
}

void QGeoRoute_SetRequest(void* ptr, void* request)
{
	static_cast<QGeoRoute*>(ptr)->setRequest(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoute_SetRouteId(void* ptr, struct QtLocation_PackedString id)
{
	static_cast<QGeoRoute*>(ptr)->setRouteId(QString::fromUtf8(id.data, id.len));
}

void QGeoRoute_SetRouteLegs(void* ptr, void* legs)
{
	static_cast<QGeoRoute*>(ptr)->setRouteLegs(*static_cast<QList<QGeoRouteLeg>*>(legs));
}

void QGeoRoute_SetTravelMode(void* ptr, long long mode)
{
	static_cast<QGeoRoute*>(ptr)->setTravelMode(static_cast<QGeoRouteRequest::TravelMode>(mode));
}

void QGeoRoute_SetTravelTime(void* ptr, int secs)
{
	static_cast<QGeoRoute*>(ptr)->setTravelTime(secs);
}

long long QGeoRoute_TravelMode(void* ptr)
{
	return static_cast<QGeoRoute*>(ptr)->travelMode();
}

int QGeoRoute_TravelTime(void* ptr)
{
	return static_cast<QGeoRoute*>(ptr)->travelTime();
}

void QGeoRoute_DestroyQGeoRoute(void* ptr)
{
	static_cast<QGeoRoute*>(ptr)->~QGeoRoute();
}

void* QGeoRoute___extendedAttributes_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoRoute___extendedAttributes_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoRoute___extendedAttributes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoRoute___extendedAttributes_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoRoute___path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRoute___path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRoute___path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoRoute___routeLegs_atList(void* ptr, int i)
{
	return new QGeoRouteLeg(({QGeoRouteLeg tmp = static_cast<QList<QGeoRouteLeg>*>(ptr)->at(i); if (i == static_cast<QList<QGeoRouteLeg>*>(ptr)->size()-1) { static_cast<QList<QGeoRouteLeg>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRoute___routeLegs_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRouteLeg>*>(ptr)->append(*static_cast<QGeoRouteLeg*>(i));
}

void* QGeoRoute___routeLegs_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRouteLeg>();
}

void* QGeoRoute___setExtendedAttributes_extendedAttributes_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoRoute___setExtendedAttributes_extendedAttributes_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoRoute___setExtendedAttributes_extendedAttributes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoRoute___setExtendedAttributes_extendedAttributes_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoRoute___setPath_path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRoute___setPath_path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRoute___setPath_path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoRoute___setRouteLegs_legs_atList(void* ptr, int i)
{
	return new QGeoRouteLeg(({QGeoRouteLeg tmp = static_cast<QList<QGeoRouteLeg>*>(ptr)->at(i); if (i == static_cast<QList<QGeoRouteLeg>*>(ptr)->size()-1) { static_cast<QList<QGeoRouteLeg>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRoute___setRouteLegs_legs_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRouteLeg>*>(ptr)->append(*static_cast<QGeoRouteLeg*>(i));
}

void* QGeoRoute___setRouteLegs_legs_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRouteLeg>();
}

struct QtLocation_PackedString QGeoRoute_____extendedAttributes_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoRoute_____extendedAttributes_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoRoute_____extendedAttributes_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtLocation_PackedString QGeoRoute_____setExtendedAttributes_extendedAttributes_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoRoute_____setExtendedAttributes_extendedAttributes_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoRoute_____setExtendedAttributes_extendedAttributes_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

Q_DECLARE_METATYPE(QGeoRouteLeg)
Q_DECLARE_METATYPE(QGeoRouteLeg*)
void* QGeoRouteLeg_NewQGeoRouteLeg()
{
	return new QGeoRouteLeg();
}

void* QGeoRouteLeg_NewQGeoRouteLeg2(void* other)
{
	return new QGeoRouteLeg(*static_cast<QGeoRouteLeg*>(other));
}

int QGeoRouteLeg_LegIndex(void* ptr)
{
	return static_cast<QGeoRouteLeg*>(ptr)->legIndex();
}

void* QGeoRouteLeg_OverallRoute(void* ptr)
{
	return new QGeoRoute(static_cast<QGeoRouteLeg*>(ptr)->overallRoute());
}

void QGeoRouteLeg_SetLegIndex(void* ptr, int idx)
{
	static_cast<QGeoRouteLeg*>(ptr)->setLegIndex(idx);
}

void QGeoRouteLeg_SetOverallRoute(void* ptr, void* route)
{
	static_cast<QGeoRouteLeg*>(ptr)->setOverallRoute(*static_cast<QGeoRoute*>(route));
}

void QGeoRouteLeg_DestroyQGeoRouteLeg(void* ptr)
{
	static_cast<QGeoRouteLeg*>(ptr)->~QGeoRouteLeg();
}

class MyQGeoRouteReply: public QGeoRouteReply
{
public:
	MyQGeoRouteReply(QGeoRouteReply::Error error, const QString &errorString, QObject *parent = Q_NULLPTR) : QGeoRouteReply(error, errorString, parent) {QGeoRouteReply_QGeoRouteReply_QRegisterMetaType();};
	MyQGeoRouteReply(const QGeoRouteRequest &request, QObject *parent = Q_NULLPTR) : QGeoRouteReply(request, parent) {QGeoRouteReply_QGeoRouteReply_QRegisterMetaType();};
	void abort() { callbackQGeoRouteReply_Abort(this); };
	void Signal_Aborted() { callbackQGeoRouteReply_Aborted(this); };
	void Signal_Error2(QGeoRouteReply::Error error, const QString & errorString) { QByteArray* tc8b6bd = new QByteArray(errorString.toUtf8()); QtLocation_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd->prepend("WHITESPACE").constData()+10), tc8b6bd->size()-10, tc8b6bd };callbackQGeoRouteReply_Error2(this, error, errorStringPacked); };
	void Signal_Finished() { callbackQGeoRouteReply_Finished(this); };
	 ~MyQGeoRouteReply() { callbackQGeoRouteReply_DestroyQGeoRouteReply(this); };
	void childEvent(QChildEvent * event) { callbackQGeoRouteReply_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoRouteReply_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoRouteReply_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoRouteReply_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoRouteReply_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoRouteReply_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGeoRouteReply_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoRouteReply_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoRouteReply_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtLocation_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGeoRouteReply_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoRouteReply_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QGeoRouteReply*)
Q_DECLARE_METATYPE(MyQGeoRouteReply*)

int QGeoRouteReply_QGeoRouteReply_QRegisterMetaType(){qRegisterMetaType<QGeoRouteReply*>(); return qRegisterMetaType<MyQGeoRouteReply*>();}

void* QGeoRouteReply_NewQGeoRouteReply(long long error, struct QtLocation_PackedString errorString, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QObject*>(parent));
	}
}

void* QGeoRouteReply_NewQGeoRouteReply2(void* request, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QWindow*>(parent));
	} else {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QObject*>(parent));
	}
}

void QGeoRouteReply_Abort(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->abort();
}

void QGeoRouteReply_AbortDefault(void* ptr)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::abort();
}

void QGeoRouteReply_ConnectAborted(void* ptr, long long t)
{
	QObject::connect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::aborted), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Aborted), static_cast<Qt::ConnectionType>(t));
}

void QGeoRouteReply_DisconnectAborted(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::aborted), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Aborted));
}

void QGeoRouteReply_Aborted(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->aborted();
}

void QGeoRouteReply_AddRoutes(void* ptr, void* routes)
{
	static_cast<QGeoRouteReply*>(ptr)->addRoutes(*static_cast<QList<QGeoRoute>*>(routes));
}

long long QGeoRouteReply_Error(void* ptr)
{
	return static_cast<QGeoRouteReply*>(ptr)->error();
}

void QGeoRouteReply_ConnectError2(void* ptr, long long t)
{
	QObject::connect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&QGeoRouteReply::error), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&MyQGeoRouteReply::Signal_Error2), static_cast<Qt::ConnectionType>(t));
}

void QGeoRouteReply_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&QGeoRouteReply::error), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&MyQGeoRouteReply::Signal_Error2));
}

void QGeoRouteReply_Error2(void* ptr, long long error, struct QtLocation_PackedString errorString)
{
	static_cast<QGeoRouteReply*>(ptr)->error(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len));
}

struct QtLocation_PackedString QGeoRouteReply_ErrorString(void* ptr)
{
	return ({ QByteArray* t834aee = new QByteArray(static_cast<QGeoRouteReply*>(ptr)->errorString().toUtf8()); QtLocation_PackedString { const_cast<char*>(t834aee->prepend("WHITESPACE").constData()+10), t834aee->size()-10, t834aee }; });
}

void QGeoRouteReply_ConnectFinished(void* ptr, long long t)
{
	QObject::connect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::finished), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Finished), static_cast<Qt::ConnectionType>(t));
}

void QGeoRouteReply_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::finished), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Finished));
}

void QGeoRouteReply_Finished(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->finished();
}

char QGeoRouteReply_IsFinished(void* ptr)
{
	return static_cast<QGeoRouteReply*>(ptr)->isFinished();
}

void* QGeoRouteReply_Request(void* ptr)
{
	return new QGeoRouteRequest(static_cast<QGeoRouteReply*>(ptr)->request());
}

struct QtLocation_PackedList QGeoRouteReply_Routes(void* ptr)
{
	return ({ QList<QGeoRoute>* tmpValuecee788 = new QList<QGeoRoute>(static_cast<QGeoRouteReply*>(ptr)->routes()); QtLocation_PackedList { tmpValuecee788, tmpValuecee788->size() }; });
}

void QGeoRouteReply_SetError(void* ptr, long long error, struct QtLocation_PackedString errorString)
{
	static_cast<QGeoRouteReply*>(ptr)->setError(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len));
}

void QGeoRouteReply_SetFinished(void* ptr, char finished)
{
	static_cast<QGeoRouteReply*>(ptr)->setFinished(finished != 0);
}

void QGeoRouteReply_SetRoutes(void* ptr, void* routes)
{
	static_cast<QGeoRouteReply*>(ptr)->setRoutes(*static_cast<QList<QGeoRoute>*>(routes));
}

void QGeoRouteReply_DestroyQGeoRouteReply(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->~QGeoRouteReply();
}

void QGeoRouteReply_DestroyQGeoRouteReplyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGeoRouteReply___addRoutes_routes_atList(void* ptr, int i)
{
	return new QGeoRoute(({QGeoRoute tmp = static_cast<QList<QGeoRoute>*>(ptr)->at(i); if (i == static_cast<QList<QGeoRoute>*>(ptr)->size()-1) { static_cast<QList<QGeoRoute>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteReply___addRoutes_routes_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRoute>*>(ptr)->append(*static_cast<QGeoRoute*>(i));
}

void* QGeoRouteReply___addRoutes_routes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRoute>();
}

void* QGeoRouteReply___routes_atList(void* ptr, int i)
{
	return new QGeoRoute(({QGeoRoute tmp = static_cast<QList<QGeoRoute>*>(ptr)->at(i); if (i == static_cast<QList<QGeoRoute>*>(ptr)->size()-1) { static_cast<QList<QGeoRoute>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteReply___routes_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRoute>*>(ptr)->append(*static_cast<QGeoRoute*>(i));
}

void* QGeoRouteReply___routes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRoute>();
}

void* QGeoRouteReply___setRoutes_routes_atList(void* ptr, int i)
{
	return new QGeoRoute(({QGeoRoute tmp = static_cast<QList<QGeoRoute>*>(ptr)->at(i); if (i == static_cast<QList<QGeoRoute>*>(ptr)->size()-1) { static_cast<QList<QGeoRoute>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteReply___setRoutes_routes_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRoute>*>(ptr)->append(*static_cast<QGeoRoute*>(i));
}

void* QGeoRouteReply___setRoutes_routes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRoute>();
}

void* QGeoRouteReply___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoRouteReply___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRouteReply___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QGeoRouteReply___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteReply___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoRouteReply___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGeoRouteReply___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoRouteReply___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRouteReply___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoRouteReply___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoRouteReply___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRouteReply___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QGeoRouteReply_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRouteReply_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRouteReply_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::customEvent(static_cast<QEvent*>(event));
}

void QGeoRouteReply_DeleteLaterDefault(void* ptr)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::deleteLater();
}

void QGeoRouteReply_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoRouteReply_EventDefault(void* ptr, void* e)
{
		return static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::event(static_cast<QEvent*>(e));
}

char QGeoRouteReply_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoRouteReply_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::metaObject());
}

void QGeoRouteReply_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::timerEvent(static_cast<QTimerEvent*>(event));
}

Q_DECLARE_METATYPE(QGeoRouteRequest*)
void* QGeoRouteRequest_NewQGeoRouteRequest(void* waypoints)
{
	return new QGeoRouteRequest(*static_cast<QList<QGeoCoordinate>*>(waypoints));
}

void* QGeoRouteRequest_NewQGeoRouteRequest2(void* origin, void* destination)
{
	return new QGeoRouteRequest(*static_cast<QGeoCoordinate*>(origin), *static_cast<QGeoCoordinate*>(destination));
}

void* QGeoRouteRequest_NewQGeoRouteRequest3(void* other)
{
	return new QGeoRouteRequest(*static_cast<QGeoRouteRequest*>(other));
}

void* QGeoRouteRequest_DepartureTime(void* ptr)
{
	return new QDateTime(static_cast<QGeoRouteRequest*>(ptr)->departureTime());
}

struct QtLocation_PackedList QGeoRouteRequest_ExcludeAreas(void* ptr)
{
	return ({ QList<QGeoRectangle>* tmpValue327b2e = new QList<QGeoRectangle>(static_cast<QGeoRouteRequest*>(ptr)->excludeAreas()); QtLocation_PackedList { tmpValue327b2e, tmpValue327b2e->size() }; });
}

struct QtLocation_PackedList QGeoRouteRequest_ExtraParameters(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValued420fe = new QMap<QString, QVariant>(static_cast<QGeoRouteRequest*>(ptr)->extraParameters()); QtLocation_PackedList { tmpValued420fe, tmpValued420fe->size() }; });
}

long long QGeoRouteRequest_FeatureWeight(void* ptr, long long featureType)
{
	return static_cast<QGeoRouteRequest*>(ptr)->featureWeight(static_cast<QGeoRouteRequest::FeatureType>(featureType));
}

long long QGeoRouteRequest_ManeuverDetail(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->maneuverDetail();
}

int QGeoRouteRequest_NumberAlternativeRoutes(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->numberAlternativeRoutes();
}

long long QGeoRouteRequest_RouteOptimization(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->routeOptimization();
}

long long QGeoRouteRequest_SegmentDetail(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->segmentDetail();
}

void QGeoRouteRequest_SetDepartureTime(void* ptr, void* departureTime)
{
	static_cast<QGeoRouteRequest*>(ptr)->setDepartureTime(*static_cast<QDateTime*>(departureTime));
}

void QGeoRouteRequest_SetExcludeAreas(void* ptr, void* areas)
{
	static_cast<QGeoRouteRequest*>(ptr)->setExcludeAreas(*static_cast<QList<QGeoRectangle>*>(areas));
}

void QGeoRouteRequest_SetExtraParameters(void* ptr, void* extraParameters)
{
	static_cast<QGeoRouteRequest*>(ptr)->setExtraParameters(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(extraParameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void QGeoRouteRequest_SetFeatureWeight(void* ptr, long long featureType, long long featureWeight)
{
	static_cast<QGeoRouteRequest*>(ptr)->setFeatureWeight(static_cast<QGeoRouteRequest::FeatureType>(featureType), static_cast<QGeoRouteRequest::FeatureWeight>(featureWeight));
}

void QGeoRouteRequest_SetManeuverDetail(void* ptr, long long maneuverDetail)
{
	static_cast<QGeoRouteRequest*>(ptr)->setManeuverDetail(static_cast<QGeoRouteRequest::ManeuverDetail>(maneuverDetail));
}

void QGeoRouteRequest_SetNumberAlternativeRoutes(void* ptr, int alternatives)
{
	static_cast<QGeoRouteRequest*>(ptr)->setNumberAlternativeRoutes(alternatives);
}

void QGeoRouteRequest_SetRouteOptimization(void* ptr, long long optimization)
{
	static_cast<QGeoRouteRequest*>(ptr)->setRouteOptimization(static_cast<QGeoRouteRequest::RouteOptimization>(optimization));
}

void QGeoRouteRequest_SetSegmentDetail(void* ptr, long long segmentDetail)
{
	static_cast<QGeoRouteRequest*>(ptr)->setSegmentDetail(static_cast<QGeoRouteRequest::SegmentDetail>(segmentDetail));
}

void QGeoRouteRequest_SetTravelModes(void* ptr, long long travelModes)
{
	static_cast<QGeoRouteRequest*>(ptr)->setTravelModes(static_cast<QGeoRouteRequest::TravelMode>(travelModes));
}

void QGeoRouteRequest_SetWaypoints(void* ptr, void* waypoints)
{
	static_cast<QGeoRouteRequest*>(ptr)->setWaypoints(*static_cast<QList<QGeoCoordinate>*>(waypoints));
}

long long QGeoRouteRequest_TravelModes(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->travelModes();
}

struct QtLocation_PackedList QGeoRouteRequest_Waypoints(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValueccc416 = new QList<QGeoCoordinate>(static_cast<QGeoRouteRequest*>(ptr)->waypoints()); QtLocation_PackedList { tmpValueccc416, tmpValueccc416->size() }; });
}

void QGeoRouteRequest_DestroyQGeoRouteRequest(void* ptr)
{
	static_cast<QGeoRouteRequest*>(ptr)->~QGeoRouteRequest();
}

void* QGeoRouteRequest___QGeoRouteRequest_waypoints_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest___QGeoRouteRequest_waypoints_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRouteRequest___QGeoRouteRequest_waypoints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoRouteRequest___excludeAreas_atList(void* ptr, int i)
{
	return new QGeoRectangle(({QGeoRectangle tmp = static_cast<QList<QGeoRectangle>*>(ptr)->at(i); if (i == static_cast<QList<QGeoRectangle>*>(ptr)->size()-1) { static_cast<QList<QGeoRectangle>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest___excludeAreas_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRectangle>*>(ptr)->append(*static_cast<QGeoRectangle*>(i));
}

void* QGeoRouteRequest___excludeAreas_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRectangle>();
}

void* QGeoRouteRequest___extraParameters_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest___extraParameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoRouteRequest___extraParameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoRouteRequest___extraParameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoRouteRequest___setExcludeAreas_areas_atList(void* ptr, int i)
{
	return new QGeoRectangle(({QGeoRectangle tmp = static_cast<QList<QGeoRectangle>*>(ptr)->at(i); if (i == static_cast<QList<QGeoRectangle>*>(ptr)->size()-1) { static_cast<QList<QGeoRectangle>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest___setExcludeAreas_areas_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRectangle>*>(ptr)->append(*static_cast<QGeoRectangle*>(i));
}

void* QGeoRouteRequest___setExcludeAreas_areas_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRectangle>();
}

void* QGeoRouteRequest___setExtraParameters_extraParameters_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest___setExtraParameters_extraParameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoRouteRequest___setExtraParameters_extraParameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoRouteRequest___setExtraParameters_extraParameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoRouteRequest___setWaypoints_waypoints_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest___setWaypoints_waypoints_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRouteRequest___setWaypoints_waypoints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

struct QtLocation_PackedList QGeoRouteRequest___setWaypointsMetadata_waypointMetadata_atList(void* ptr, int i)
{
	return ({ QMap<QString, QVariant>* tmpValue5dc32e = new QMap<QString, QVariant>(({QMap<QString, QVariant> tmp = static_cast<QList<QMap<QString, QVariant>>*>(ptr)->at(i); if (i == static_cast<QList<QMap<QString, QVariant>>*>(ptr)->size()-1) { static_cast<QList<QMap<QString, QVariant>>*>(ptr)->~QList(); free(ptr); }; tmp; })); QtLocation_PackedList { tmpValue5dc32e, tmpValue5dc32e->size() }; });
}

void QGeoRouteRequest___setWaypointsMetadata_waypointMetadata_setList(void* ptr, void* i)
{
	static_cast<QList<QMap<QString, QVariant>>*>(ptr)->append(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(i); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void* QGeoRouteRequest___setWaypointsMetadata_waypointMetadata_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariantMap>();
}

void* QGeoRouteRequest___waypoints_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest___waypoints_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRouteRequest___waypoints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

struct QtLocation_PackedList QGeoRouteRequest___waypointsMetadata_atList(void* ptr, int i)
{
	return ({ QMap<QString, QVariant>* tmpValue5dc32e = new QMap<QString, QVariant>(({QMap<QString, QVariant> tmp = static_cast<QList<QMap<QString, QVariant>>*>(ptr)->at(i); if (i == static_cast<QList<QMap<QString, QVariant>>*>(ptr)->size()-1) { static_cast<QList<QMap<QString, QVariant>>*>(ptr)->~QList(); free(ptr); }; tmp; })); QtLocation_PackedList { tmpValue5dc32e, tmpValue5dc32e->size() }; });
}

void QGeoRouteRequest___waypointsMetadata_setList(void* ptr, void* i)
{
	static_cast<QList<QMap<QString, QVariant>>*>(ptr)->append(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(i); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void* QGeoRouteRequest___waypointsMetadata_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariantMap>();
}

struct QtLocation_PackedString QGeoRouteRequest_____extraParameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoRouteRequest_____extraParameters_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoRouteRequest_____extraParameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtLocation_PackedString QGeoRouteRequest_____setExtraParameters_extraParameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoRouteRequest_____setExtraParameters_extraParameters_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoRouteRequest_____setExtraParameters_extraParameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_atList_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_atList_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_atList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_atList_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_setList_i_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_setList_i_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_setList_i_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoRouteRequest_____setWaypointsMetadata_waypointMetadata_setList_i_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoRouteRequest_____waypointsMetadata_atList_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest_____waypointsMetadata_atList_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoRouteRequest_____waypointsMetadata_atList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoRouteRequest_____waypointsMetadata_atList_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoRouteRequest_____waypointsMetadata_setList_i_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoRouteRequest_____waypointsMetadata_setList_i_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoRouteRequest_____waypointsMetadata_setList_i_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoRouteRequest_____waypointsMetadata_setList_i_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtLocation_PackedString QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_atList_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_atList_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_atList_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtLocation_PackedString QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_setList_i_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_setList_i_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoRouteRequest_______setWaypointsMetadata_waypointMetadata_setList_i_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtLocation_PackedString QGeoRouteRequest_______waypointsMetadata_atList_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoRouteRequest_______waypointsMetadata_atList_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoRouteRequest_______waypointsMetadata_atList_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtLocation_PackedString QGeoRouteRequest_______waypointsMetadata_setList_i_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoRouteRequest_______waypointsMetadata_setList_i_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoRouteRequest_______waypointsMetadata_setList_i_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

Q_DECLARE_METATYPE(QGeoRouteSegment)
Q_DECLARE_METATYPE(QGeoRouteSegment*)
void* QGeoRouteSegment_NewQGeoRouteSegment()
{
	return new QGeoRouteSegment();
}

void* QGeoRouteSegment_NewQGeoRouteSegment2(void* other)
{
	return new QGeoRouteSegment(*static_cast<QGeoRouteSegment*>(other));
}

double QGeoRouteSegment_Distance(void* ptr)
{
	return static_cast<QGeoRouteSegment*>(ptr)->distance();
}

char QGeoRouteSegment_IsLegLastSegment(void* ptr)
{
	return static_cast<QGeoRouteSegment*>(ptr)->isLegLastSegment();
}

char QGeoRouteSegment_IsValid(void* ptr)
{
	return static_cast<QGeoRouteSegment*>(ptr)->isValid();
}

void* QGeoRouteSegment_Maneuver(void* ptr)
{
	return new QGeoManeuver(static_cast<QGeoRouteSegment*>(ptr)->maneuver());
}

void* QGeoRouteSegment_NextRouteSegment(void* ptr)
{
	return new QGeoRouteSegment(static_cast<QGeoRouteSegment*>(ptr)->nextRouteSegment());
}

struct QtLocation_PackedList QGeoRouteSegment_Path(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValue0158d6 = new QList<QGeoCoordinate>(static_cast<QGeoRouteSegment*>(ptr)->path()); QtLocation_PackedList { tmpValue0158d6, tmpValue0158d6->size() }; });
}

void QGeoRouteSegment_SetDistance(void* ptr, double distance)
{
	static_cast<QGeoRouteSegment*>(ptr)->setDistance(distance);
}

void QGeoRouteSegment_SetManeuver(void* ptr, void* maneuver)
{
	static_cast<QGeoRouteSegment*>(ptr)->setManeuver(*static_cast<QGeoManeuver*>(maneuver));
}

void QGeoRouteSegment_SetNextRouteSegment(void* ptr, void* routeSegment)
{
	static_cast<QGeoRouteSegment*>(ptr)->setNextRouteSegment(*static_cast<QGeoRouteSegment*>(routeSegment));
}

void QGeoRouteSegment_SetPath(void* ptr, void* path)
{
	static_cast<QGeoRouteSegment*>(ptr)->setPath(*static_cast<QList<QGeoCoordinate>*>(path));
}

void QGeoRouteSegment_SetTravelTime(void* ptr, int secs)
{
	static_cast<QGeoRouteSegment*>(ptr)->setTravelTime(secs);
}

int QGeoRouteSegment_TravelTime(void* ptr)
{
	return static_cast<QGeoRouteSegment*>(ptr)->travelTime();
}

void QGeoRouteSegment_DestroyQGeoRouteSegment(void* ptr)
{
	static_cast<QGeoRouteSegment*>(ptr)->~QGeoRouteSegment();
}

void* QGeoRouteSegment___path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteSegment___path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRouteSegment___path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoRouteSegment___setPath_path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRouteSegment___setPath_path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRouteSegment___setPath_path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

class MyQGeoRoutingManager: public QGeoRoutingManager
{
public:
	void Signal_Error(QGeoRouteReply * reply, QGeoRouteReply::Error error, QString errorString) { QByteArray* tc8b6bd = new QByteArray(errorString.toUtf8()); QtLocation_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd->prepend("WHITESPACE").constData()+10), tc8b6bd->size()-10, tc8b6bd };callbackQGeoRoutingManager_Error(this, reply, error, errorStringPacked); };
	void Signal_Finished(QGeoRouteReply * reply) { callbackQGeoRoutingManager_Finished(this, reply); };
	 ~MyQGeoRoutingManager() { callbackQGeoRoutingManager_DestroyQGeoRoutingManager(this); };
	void childEvent(QChildEvent * event) { callbackQGeoRoutingManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoRoutingManager_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoRoutingManager_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoRoutingManager_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGeoRoutingManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoRoutingManager_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoRoutingManager_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtLocation_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGeoRoutingManager_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoRoutingManager_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QGeoRoutingManager*)
Q_DECLARE_METATYPE(MyQGeoRoutingManager*)

int QGeoRoutingManager_QGeoRoutingManager_QRegisterMetaType(){qRegisterMetaType<QGeoRoutingManager*>(); return qRegisterMetaType<MyQGeoRoutingManager*>();}

void* QGeoRoutingManager_CalculateRoute(void* ptr, void* request)
{
	return static_cast<QGeoRoutingManager*>(ptr)->calculateRoute(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoutingManager_ConnectError(void* ptr, long long t)
{
	QObject::connect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManager::error), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManager::Signal_Error), static_cast<Qt::ConnectionType>(t));
}

void QGeoRoutingManager_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManager::error), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManager::Signal_Error));
}

void QGeoRoutingManager_Error(void* ptr, void* reply, long long error, struct QtLocation_PackedString errorString)
{
	static_cast<QGeoRoutingManager*>(ptr)->error(static_cast<QGeoRouteReply*>(reply), static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len));
}

void QGeoRoutingManager_ConnectFinished(void* ptr, long long t)
{
	QObject::connect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *)>(&QGeoRoutingManager::finished), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *)>(&MyQGeoRoutingManager::Signal_Finished), static_cast<Qt::ConnectionType>(t));
}

void QGeoRoutingManager_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *)>(&QGeoRoutingManager::finished), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *)>(&MyQGeoRoutingManager::Signal_Finished));
}

void QGeoRoutingManager_Finished(void* ptr, void* reply)
{
	static_cast<QGeoRoutingManager*>(ptr)->finished(static_cast<QGeoRouteReply*>(reply));
}

void* QGeoRoutingManager_Locale(void* ptr)
{
	return new QLocale(static_cast<QGeoRoutingManager*>(ptr)->locale());
}

struct QtLocation_PackedString QGeoRoutingManager_ManagerName(void* ptr)
{
	return ({ QByteArray* t48ee82 = new QByteArray(static_cast<QGeoRoutingManager*>(ptr)->managerName().toUtf8()); QtLocation_PackedString { const_cast<char*>(t48ee82->prepend("WHITESPACE").constData()+10), t48ee82->size()-10, t48ee82 }; });
}

int QGeoRoutingManager_ManagerVersion(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->managerVersion();
}

long long QGeoRoutingManager_MeasurementSystem(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->measurementSystem();
}

void QGeoRoutingManager_SetLocale(void* ptr, void* locale)
{
	static_cast<QGeoRoutingManager*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoRoutingManager_SetMeasurementSystem(void* ptr, long long system)
{
	static_cast<QGeoRoutingManager*>(ptr)->setMeasurementSystem(static_cast<QLocale::MeasurementSystem>(system));
}

long long QGeoRoutingManager_SupportedFeatureTypes(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedFeatureTypes();
}

long long QGeoRoutingManager_SupportedFeatureWeights(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedFeatureWeights();
}

long long QGeoRoutingManager_SupportedManeuverDetails(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedManeuverDetails();
}

long long QGeoRoutingManager_SupportedRouteOptimizations(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedRouteOptimizations();
}

long long QGeoRoutingManager_SupportedSegmentDetails(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedSegmentDetails();
}

long long QGeoRoutingManager_SupportedTravelModes(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedTravelModes();
}

void* QGeoRoutingManager_UpdateRoute(void* ptr, void* route, void* position)
{
	return static_cast<QGeoRoutingManager*>(ptr)->updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void QGeoRoutingManager_DestroyQGeoRoutingManager(void* ptr)
{
	static_cast<QGeoRoutingManager*>(ptr)->~QGeoRoutingManager();
}

void QGeoRoutingManager_DestroyQGeoRoutingManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGeoRoutingManager___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoRoutingManager___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManager___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QGeoRoutingManager___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRoutingManager___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoRoutingManager___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGeoRoutingManager___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoRoutingManager___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManager___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoRoutingManager___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoRoutingManager___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManager___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QGeoRoutingManager_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRoutingManager_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManager_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::customEvent(static_cast<QEvent*>(event));
}

void QGeoRoutingManager_DeleteLaterDefault(void* ptr)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::deleteLater();
}

void QGeoRoutingManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoRoutingManager_EventDefault(void* ptr, void* e)
{
		return static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::event(static_cast<QEvent*>(e));
}

char QGeoRoutingManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoRoutingManager_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::metaObject());
}

void QGeoRoutingManager_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQGeoRoutingManagerEngine: public QGeoRoutingManagerEngine
{
public:
	MyQGeoRoutingManagerEngine(const QVariantMap &parameters, QObject *parent = Q_NULLPTR) : QGeoRoutingManagerEngine(parameters, parent) {QGeoRoutingManagerEngine_QGeoRoutingManagerEngine_QRegisterMetaType();};
	QGeoRouteReply * calculateRoute(const QGeoRouteRequest & request) { return static_cast<QGeoRouteReply*>(callbackQGeoRoutingManagerEngine_CalculateRoute(this, const_cast<QGeoRouteRequest*>(&request))); };
	void Signal_Error(QGeoRouteReply * reply, QGeoRouteReply::Error error, QString errorString) { QByteArray* tc8b6bd = new QByteArray(errorString.toUtf8()); QtLocation_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd->prepend("WHITESPACE").constData()+10), tc8b6bd->size()-10, tc8b6bd };callbackQGeoRoutingManagerEngine_Error(this, reply, error, errorStringPacked); };
	void Signal_Finished(QGeoRouteReply * reply) { callbackQGeoRoutingManagerEngine_Finished(this, reply); };
	QGeoRouteReply * updateRoute(const QGeoRoute & route, const QGeoCoordinate & position) { return static_cast<QGeoRouteReply*>(callbackQGeoRoutingManagerEngine_UpdateRoute(this, const_cast<QGeoRoute*>(&route), const_cast<QGeoCoordinate*>(&position))); };
	 ~MyQGeoRoutingManagerEngine() { callbackQGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(this); };
	void childEvent(QChildEvent * event) { callbackQGeoRoutingManagerEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManagerEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoRoutingManagerEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoRoutingManagerEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoRoutingManagerEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManagerEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGeoRoutingManagerEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoRoutingManagerEngine_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoRoutingManagerEngine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtLocation_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGeoRoutingManagerEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoRoutingManagerEngine_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QGeoRoutingManagerEngine*)
Q_DECLARE_METATYPE(MyQGeoRoutingManagerEngine*)

int QGeoRoutingManagerEngine_QGeoRoutingManagerEngine_QRegisterMetaType(){qRegisterMetaType<QGeoRoutingManagerEngine*>(); return qRegisterMetaType<MyQGeoRoutingManagerEngine*>();}

void* QGeoRoutingManagerEngine_NewQGeoRoutingManagerEngine(void* parameters, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QWindow*>(parent));
	} else {
		return new MyQGeoRoutingManagerEngine(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), static_cast<QObject*>(parent));
	}
}

void* QGeoRoutingManagerEngine_CalculateRoute(void* ptr, void* request)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->calculateRoute(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoutingManagerEngine_ConnectError(void* ptr, long long t)
{
	QObject::connect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManagerEngine::error), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManagerEngine::Signal_Error), static_cast<Qt::ConnectionType>(t));
}

void QGeoRoutingManagerEngine_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManagerEngine::error), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManagerEngine::Signal_Error));
}

void QGeoRoutingManagerEngine_Error(void* ptr, void* reply, long long error, struct QtLocation_PackedString errorString)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->error(static_cast<QGeoRouteReply*>(reply), static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len));
}

void QGeoRoutingManagerEngine_ConnectFinished(void* ptr, long long t)
{
	QObject::connect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&QGeoRoutingManagerEngine::finished), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&MyQGeoRoutingManagerEngine::Signal_Finished), static_cast<Qt::ConnectionType>(t));
}

void QGeoRoutingManagerEngine_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&QGeoRoutingManagerEngine::finished), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&MyQGeoRoutingManagerEngine::Signal_Finished));
}

void QGeoRoutingManagerEngine_Finished(void* ptr, void* reply)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->finished(static_cast<QGeoRouteReply*>(reply));
}

void* QGeoRoutingManagerEngine_Locale(void* ptr)
{
	return new QLocale(static_cast<QGeoRoutingManagerEngine*>(ptr)->locale());
}

struct QtLocation_PackedString QGeoRoutingManagerEngine_ManagerName(void* ptr)
{
	return ({ QByteArray* tce87b5 = new QByteArray(static_cast<QGeoRoutingManagerEngine*>(ptr)->managerName().toUtf8()); QtLocation_PackedString { const_cast<char*>(tce87b5->prepend("WHITESPACE").constData()+10), tce87b5->size()-10, tce87b5 }; });
}

int QGeoRoutingManagerEngine_ManagerVersion(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->managerVersion();
}

long long QGeoRoutingManagerEngine_MeasurementSystem(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->measurementSystem();
}

void QGeoRoutingManagerEngine_SetLocale(void* ptr, void* locale)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoRoutingManagerEngine_SetMeasurementSystem(void* ptr, long long system)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setMeasurementSystem(static_cast<QLocale::MeasurementSystem>(system));
}

void QGeoRoutingManagerEngine_SetSupportedFeatureTypes(void* ptr, long long featureTypes)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedFeatureTypes(static_cast<QGeoRouteRequest::FeatureType>(featureTypes));
}

void QGeoRoutingManagerEngine_SetSupportedFeatureWeights(void* ptr, long long featureWeights)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedFeatureWeights(static_cast<QGeoRouteRequest::FeatureWeight>(featureWeights));
}

void QGeoRoutingManagerEngine_SetSupportedManeuverDetails(void* ptr, long long maneuverDetails)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedManeuverDetails(static_cast<QGeoRouteRequest::ManeuverDetail>(maneuverDetails));
}

void QGeoRoutingManagerEngine_SetSupportedRouteOptimizations(void* ptr, long long optimizations)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedRouteOptimizations(static_cast<QGeoRouteRequest::RouteOptimization>(optimizations));
}

void QGeoRoutingManagerEngine_SetSupportedSegmentDetails(void* ptr, long long segmentDetails)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedSegmentDetails(static_cast<QGeoRouteRequest::SegmentDetail>(segmentDetails));
}

void QGeoRoutingManagerEngine_SetSupportedTravelModes(void* ptr, long long travelModes)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedTravelModes(static_cast<QGeoRouteRequest::TravelMode>(travelModes));
}

long long QGeoRoutingManagerEngine_SupportedFeatureTypes(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedFeatureTypes();
}

long long QGeoRoutingManagerEngine_SupportedFeatureWeights(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedFeatureWeights();
}

long long QGeoRoutingManagerEngine_SupportedManeuverDetails(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedManeuverDetails();
}

long long QGeoRoutingManagerEngine_SupportedRouteOptimizations(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedRouteOptimizations();
}

long long QGeoRoutingManagerEngine_SupportedSegmentDetails(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedSegmentDetails();
}

long long QGeoRoutingManagerEngine_SupportedTravelModes(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedTravelModes();
}

void* QGeoRoutingManagerEngine_UpdateRoute(void* ptr, void* route, void* position)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void* QGeoRoutingManagerEngine_UpdateRouteDefault(void* ptr, void* route, void* position)
{
		return static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(void* ptr)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->~QGeoRoutingManagerEngine();
}

void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtLocation_PackedString QGeoRoutingManagerEngine_____QGeoRoutingManagerEngine_parameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoRoutingManagerEngine_____QGeoRoutingManagerEngine_parameters_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoRoutingManagerEngine_____QGeoRoutingManagerEngine_parameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* QGeoRoutingManagerEngine___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoRoutingManagerEngine___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManagerEngine___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QGeoRoutingManagerEngine___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRoutingManagerEngine___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoRoutingManagerEngine___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGeoRoutingManagerEngine___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoRoutingManagerEngine___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManagerEngine___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoRoutingManagerEngine___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoRoutingManagerEngine___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManagerEngine___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QGeoRoutingManagerEngine_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRoutingManagerEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManagerEngine_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::customEvent(static_cast<QEvent*>(event));
}

void QGeoRoutingManagerEngine_DeleteLaterDefault(void* ptr)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::deleteLater();
}

void QGeoRoutingManagerEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoRoutingManagerEngine_EventDefault(void* ptr, void* e)
{
		return static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::event(static_cast<QEvent*>(e));
}

char QGeoRoutingManagerEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoRoutingManagerEngine_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::metaObject());
}

void QGeoRoutingManagerEngine_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQGeoServiceProvider: public QGeoServiceProvider
{
public:
	MyQGeoServiceProvider(const QString &providerName, const QVariantMap &parameters = QVariantMap(), bool allowExperimental = false) : QGeoServiceProvider(providerName, parameters, allowExperimental) {QGeoServiceProvider_QGeoServiceProvider_QRegisterMetaType();};
	 ~MyQGeoServiceProvider() { callbackQGeoServiceProvider_DestroyQGeoServiceProvider(this); };
	void childEvent(QChildEvent * event) { callbackQGeoServiceProvider_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoServiceProvider_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoServiceProvider_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoServiceProvider_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoServiceProvider_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoServiceProvider_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGeoServiceProvider_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoServiceProvider_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoServiceProvider_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtLocation_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGeoServiceProvider_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoServiceProvider_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QGeoServiceProvider*)
Q_DECLARE_METATYPE(MyQGeoServiceProvider*)

int QGeoServiceProvider_QGeoServiceProvider_QRegisterMetaType(){qRegisterMetaType<QGeoServiceProvider*>(); return qRegisterMetaType<MyQGeoServiceProvider*>();}

int QGeoServiceProvider_OnlineGeocodingFeature_Type()
{
	return QGeoServiceProvider::OnlineGeocodingFeature;
}

int QGeoServiceProvider_OfflineGeocodingFeature_Type()
{
	return QGeoServiceProvider::OfflineGeocodingFeature;
}

int QGeoServiceProvider_ReverseGeocodingFeature_Type()
{
	return QGeoServiceProvider::ReverseGeocodingFeature;
}

int QGeoServiceProvider_LocalizedGeocodingFeature_Type()
{
	return QGeoServiceProvider::LocalizedGeocodingFeature;
}

int QGeoServiceProvider_AnyGeocodingFeatures_Type()
{
	return QGeoServiceProvider::AnyGeocodingFeatures;
}

int QGeoServiceProvider_OnlineMappingFeature_Type()
{
	return QGeoServiceProvider::OnlineMappingFeature;
}

int QGeoServiceProvider_OfflineMappingFeature_Type()
{
	return QGeoServiceProvider::OfflineMappingFeature;
}

int QGeoServiceProvider_LocalizedMappingFeature_Type()
{
	return QGeoServiceProvider::LocalizedMappingFeature;
}

int QGeoServiceProvider_AnyMappingFeatures_Type()
{
	return QGeoServiceProvider::AnyMappingFeatures;
}

int QGeoServiceProvider_OnlineNavigationFeature_Type()
{
	return QGeoServiceProvider::OnlineNavigationFeature;
}

int QGeoServiceProvider_OfflineNavigationFeature_Type()
{
	return QGeoServiceProvider::OfflineNavigationFeature;
}

int QGeoServiceProvider_AnyNavigationFeatures_Type()
{
	return QGeoServiceProvider::AnyNavigationFeatures;
}

int QGeoServiceProvider_OnlinePlacesFeature_Type()
{
	return QGeoServiceProvider::OnlinePlacesFeature;
}

int QGeoServiceProvider_OfflinePlacesFeature_Type()
{
	return QGeoServiceProvider::OfflinePlacesFeature;
}

int QGeoServiceProvider_SavePlaceFeature_Type()
{
	return QGeoServiceProvider::SavePlaceFeature;
}

int QGeoServiceProvider_RemovePlaceFeature_Type()
{
	return QGeoServiceProvider::RemovePlaceFeature;
}

int QGeoServiceProvider_SaveCategoryFeature_Type()
{
	return QGeoServiceProvider::SaveCategoryFeature;
}

int QGeoServiceProvider_RemoveCategoryFeature_Type()
{
	return QGeoServiceProvider::RemoveCategoryFeature;
}

int QGeoServiceProvider_PlaceRecommendationsFeature_Type()
{
	return QGeoServiceProvider::PlaceRecommendationsFeature;
}

int QGeoServiceProvider_SearchSuggestionsFeature_Type()
{
	return QGeoServiceProvider::SearchSuggestionsFeature;
}

int QGeoServiceProvider_LocalizedPlacesFeature_Type()
{
	return QGeoServiceProvider::LocalizedPlacesFeature;
}

int QGeoServiceProvider_NotificationsFeature_Type()
{
	return QGeoServiceProvider::NotificationsFeature;
}

int QGeoServiceProvider_PlaceMatchingFeature_Type()
{
	return QGeoServiceProvider::PlaceMatchingFeature;
}

int QGeoServiceProvider_AnyPlacesFeatures_Type()
{
	return QGeoServiceProvider::AnyPlacesFeatures;
}

int QGeoServiceProvider_OnlineRoutingFeature_Type()
{
	return QGeoServiceProvider::OnlineRoutingFeature;
}

int QGeoServiceProvider_OfflineRoutingFeature_Type()
{
	return QGeoServiceProvider::OfflineRoutingFeature;
}

int QGeoServiceProvider_LocalizedRoutingFeature_Type()
{
	return QGeoServiceProvider::LocalizedRoutingFeature;
}

int QGeoServiceProvider_RouteUpdatesFeature_Type()
{
	return QGeoServiceProvider::RouteUpdatesFeature;
}

int QGeoServiceProvider_AlternativeRoutesFeature_Type()
{
	return QGeoServiceProvider::AlternativeRoutesFeature;
}

int QGeoServiceProvider_ExcludeAreasRoutingFeature_Type()
{
	return QGeoServiceProvider::ExcludeAreasRoutingFeature;
}

int QGeoServiceProvider_AnyRoutingFeatures_Type()
{
	return QGeoServiceProvider::AnyRoutingFeatures;
}

void* QGeoServiceProvider_NewQGeoServiceProvider(struct QtLocation_PackedString providerName, void* parameters, char allowExperimental)
{
	return new MyQGeoServiceProvider(QString::fromUtf8(providerName.data, providerName.len), ({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }), allowExperimental != 0);
}

struct QtLocation_PackedString QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders()
{
	return ({ QByteArray* teec3ca = new QByteArray(QGeoServiceProvider::availableServiceProviders().join("¡¦!").toUtf8()); QtLocation_PackedString { const_cast<char*>(teec3ca->prepend("WHITESPACE").constData()+10), teec3ca->size()-10, teec3ca }; });
}

long long QGeoServiceProvider_Error(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->error();
}

struct QtLocation_PackedString QGeoServiceProvider_ErrorString(void* ptr)
{
	return ({ QByteArray* t90b712 = new QByteArray(static_cast<QGeoServiceProvider*>(ptr)->errorString().toUtf8()); QtLocation_PackedString { const_cast<char*>(t90b712->prepend("WHITESPACE").constData()+10), t90b712->size()-10, t90b712 }; });
}

long long QGeoServiceProvider_GeocodingError(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->geocodingError();
}

struct QtLocation_PackedString QGeoServiceProvider_GeocodingErrorString(void* ptr)
{
	return ({ QByteArray* t755a8d = new QByteArray(static_cast<QGeoServiceProvider*>(ptr)->geocodingErrorString().toUtf8()); QtLocation_PackedString { const_cast<char*>(t755a8d->prepend("WHITESPACE").constData()+10), t755a8d->size()-10, t755a8d }; });
}

long long QGeoServiceProvider_GeocodingFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->geocodingFeatures();
}

void* QGeoServiceProvider_GeocodingManager(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->geocodingManager();
}

long long QGeoServiceProvider_MappingError(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->mappingError();
}

struct QtLocation_PackedString QGeoServiceProvider_MappingErrorString(void* ptr)
{
	return ({ QByteArray* ta0acf4 = new QByteArray(static_cast<QGeoServiceProvider*>(ptr)->mappingErrorString().toUtf8()); QtLocation_PackedString { const_cast<char*>(ta0acf4->prepend("WHITESPACE").constData()+10), ta0acf4->size()-10, ta0acf4 }; });
}

long long QGeoServiceProvider_MappingFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->mappingFeatures();
}

long long QGeoServiceProvider_NavigationError(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->navigationError();
}

struct QtLocation_PackedString QGeoServiceProvider_NavigationErrorString(void* ptr)
{
	return ({ QByteArray* t14312f = new QByteArray(static_cast<QGeoServiceProvider*>(ptr)->navigationErrorString().toUtf8()); QtLocation_PackedString { const_cast<char*>(t14312f->prepend("WHITESPACE").constData()+10), t14312f->size()-10, t14312f }; });
}

long long QGeoServiceProvider_NavigationFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->navigationFeatures();
}

void* QGeoServiceProvider_PlaceManager(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->placeManager();
}

long long QGeoServiceProvider_PlacesError(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->placesError();
}

struct QtLocation_PackedString QGeoServiceProvider_PlacesErrorString(void* ptr)
{
	return ({ QByteArray* t14e0cd = new QByteArray(static_cast<QGeoServiceProvider*>(ptr)->placesErrorString().toUtf8()); QtLocation_PackedString { const_cast<char*>(t14e0cd->prepend("WHITESPACE").constData()+10), t14e0cd->size()-10, t14e0cd }; });
}

long long QGeoServiceProvider_PlacesFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->placesFeatures();
}

long long QGeoServiceProvider_RoutingError(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->routingError();
}

struct QtLocation_PackedString QGeoServiceProvider_RoutingErrorString(void* ptr)
{
	return ({ QByteArray* tb7c210 = new QByteArray(static_cast<QGeoServiceProvider*>(ptr)->routingErrorString().toUtf8()); QtLocation_PackedString { const_cast<char*>(tb7c210->prepend("WHITESPACE").constData()+10), tb7c210->size()-10, tb7c210 }; });
}

long long QGeoServiceProvider_RoutingFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->routingFeatures();
}

void* QGeoServiceProvider_RoutingManager(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->routingManager();
}

void QGeoServiceProvider_SetAllowExperimental(void* ptr, char allow)
{
	static_cast<QGeoServiceProvider*>(ptr)->setAllowExperimental(allow != 0);
}

void QGeoServiceProvider_SetLocale(void* ptr, void* locale)
{
	static_cast<QGeoServiceProvider*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoServiceProvider_SetParameters(void* ptr, void* parameters)
{
	static_cast<QGeoServiceProvider*>(ptr)->setParameters(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void QGeoServiceProvider_DestroyQGeoServiceProvider(void* ptr)
{
	static_cast<QGeoServiceProvider*>(ptr)->~QGeoServiceProvider();
}

void QGeoServiceProvider_DestroyQGeoServiceProviderDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGeoServiceProvider___QGeoServiceProvider_parameters_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoServiceProvider___QGeoServiceProvider_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProvider___QGeoServiceProvider_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoServiceProvider___QGeoServiceProvider_parameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoServiceProvider___setParameters_parameters_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoServiceProvider___setParameters_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProvider___setParameters_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoServiceProvider___setParameters_parameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtLocation_PackedString QGeoServiceProvider_____QGeoServiceProvider_parameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoServiceProvider_____QGeoServiceProvider_parameters_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProvider_____QGeoServiceProvider_parameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtLocation_PackedString QGeoServiceProvider_____setParameters_parameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoServiceProvider_____setParameters_parameters_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProvider_____setParameters_parameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* QGeoServiceProvider___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoServiceProvider___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoServiceProvider___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QGeoServiceProvider___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoServiceProvider___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoServiceProvider___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGeoServiceProvider___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoServiceProvider___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoServiceProvider___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoServiceProvider___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoServiceProvider___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoServiceProvider___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QGeoServiceProvider_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoServiceProvider_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoServiceProvider_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::customEvent(static_cast<QEvent*>(event));
}

void QGeoServiceProvider_DeleteLaterDefault(void* ptr)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::deleteLater();
}

void QGeoServiceProvider_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoServiceProvider_EventDefault(void* ptr, void* e)
{
		return static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::event(static_cast<QEvent*>(e));
}

char QGeoServiceProvider_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoServiceProvider_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::metaObject());
}

void QGeoServiceProvider_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQGeoServiceProviderFactory: public QGeoServiceProviderFactory
{
public:
	 ~MyQGeoServiceProviderFactory() { callbackQGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(this); };
};

Q_DECLARE_METATYPE(QGeoServiceProviderFactory*)
Q_DECLARE_METATYPE(MyQGeoServiceProviderFactory*)

int QGeoServiceProviderFactory_QGeoServiceProviderFactory_QRegisterMetaType(){qRegisterMetaType<QGeoServiceProviderFactory*>(); return qRegisterMetaType<MyQGeoServiceProviderFactory*>();}

void QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(void* ptr)
{
	static_cast<QGeoServiceProviderFactory*>(ptr)->~QGeoServiceProviderFactory();
}

void QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactoryDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoServiceProviderFactory___createMappingManagerEngine_parameters_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoServiceProviderFactory___createMappingManagerEngine_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProviderFactory___createMappingManagerEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoServiceProviderFactory___createMappingManagerEngine_parameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_atList(void* ptr, struct QtLocation_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtLocation_PackedList QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtLocation_PackedString QGeoServiceProviderFactory_____createGeocodingManagerEngine_parameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoServiceProviderFactory_____createGeocodingManagerEngine_parameters_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProviderFactory_____createGeocodingManagerEngine_parameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtLocation_PackedString QGeoServiceProviderFactory_____createMappingManagerEngine_parameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoServiceProviderFactory_____createMappingManagerEngine_parameters_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProviderFactory_____createMappingManagerEngine_parameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtLocation_PackedString QGeoServiceProviderFactory_____createPlaceManagerEngine_parameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoServiceProviderFactory_____createPlaceManagerEngine_parameters_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProviderFactory_____createPlaceManagerEngine_parameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtLocation_PackedString QGeoServiceProviderFactory_____createRoutingManagerEngine_parameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtLocation_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QGeoServiceProviderFactory_____createRoutingManagerEngine_parameters_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProviderFactory_____createRoutingManagerEngine_parameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

