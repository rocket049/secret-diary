// +build minimal

#define protected public
#define private public

#include "gui-minimal.h"
#include "_cgo_export.h"

#include <QBitmap>
#include <QBrush>
#include <QByteArray>
#include <QChar>
#include <QChildEvent>
#include <QClipboard>
#include <QCloseEvent>
#include <QColor>
#include <QContextMenuEvent>
#include <QCursor>
#include <QDataStream>
#include <QEvent>
#include <QFont>
#include <QFontMetrics>
#include <QGlyphRun>
#include <QGradient>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QGuiApplication>
#include <QHash>
#include <QHideEvent>
#include <QIODevice>
#include <QIcon>
#include <QIconEngine>
#include <QImage>
#include <QInputEvent>
#include <QIntValidator>
#include <QKeyEvent>
#include <QKeySequence>
#include <QLatin1String>
#include <QLayout>
#include <QLine>
#include <QLineF>
#include <QMap>
#include <QMatrix4x4>
#include <QMetaMethod>
#include <QMetaObject>
#include <QModelIndex>
#include <QMouseEvent>
#include <QObject>
#include <QPageSize>
#include <QPagedPaintDevice>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEngineState>
#include <QPaintEvent>
#include <QPainter>
#include <QPainterPath>
#include <QPalette>
#include <QPersistentModelIndex>
#include <QPixmap>
#include <QPoint>
#include <QPointF>
#include <QPolygon>
#include <QPolygonF>
#include <QRect>
#include <QRectF>
#include <QRegExp>
#include <QRegion>
#include <QRegularExpression>
#include <QResizeEvent>
#include <QRgba64>
#include <QScreen>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QStandardItem>
#include <QStandardItemModel>
#include <QString>
#include <QSurface>
#include <QSurfaceFormat>
#include <QTextBlock>
#include <QTextBlockFormat>
#include <QTextBlockGroup>
#include <QTextCharFormat>
#include <QTextCursor>
#include <QTextDocument>
#include <QTextDocumentFragment>
#include <QTextDocumentWriter>
#include <QTextFormat>
#include <QTextFrame>
#include <QTextFrameFormat>
#include <QTextImageFormat>
#include <QTextLayout>
#include <QTextLength>
#include <QTextList>
#include <QTextListFormat>
#include <QTextObject>
#include <QTextOption>
#include <QTextTable>
#include <QTextTableCell>
#include <QTextTableFormat>
#include <QTimerEvent>
#include <QTransform>
#include <QUrl>
#include <QValidator>
#include <QVariant>
#include <QVector>
#include <QVector2D>
#include <QVector3D>
#include <QVector4D>
#include <QWidget>
#include <QWindow>
#include <QStringList>

class MyQBitmap: public QBitmap
{
public:
	MyQBitmap() : QBitmap() {QBitmap_QBitmap_QRegisterMetaType();};
	MyQBitmap(const QPixmap &pixmap) : QBitmap(pixmap) {QBitmap_QBitmap_QRegisterMetaType();};
	MyQBitmap(int width, int height) : QBitmap(width, height) {QBitmap_QBitmap_QRegisterMetaType();};
	MyQBitmap(const QSize &size) : QBitmap(size) {QBitmap_QBitmap_QRegisterMetaType();};
	MyQBitmap(const QString &fileName, const char *format = Q_NULLPTR) : QBitmap(fileName, format) {QBitmap_QBitmap_QRegisterMetaType();};
	 ~MyQBitmap() { callbackQBitmap_DestroyQBitmap(this); };
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPaintDevice_Metric(const_cast<void*>(static_cast<const void*>(this)), metric); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPixmap_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QBitmap*)
Q_DECLARE_METATYPE(MyQBitmap*)

int QBitmap_QBitmap_QRegisterMetaType(){qRegisterMetaType<QBitmap*>(); return qRegisterMetaType<MyQBitmap*>();}

void* QBitmap_NewQBitmap()
{
	return new MyQBitmap();
}

void* QBitmap_NewQBitmap2(void* pixmap)
{
	return new MyQBitmap(*static_cast<QPixmap*>(pixmap));
}

void* QBitmap_NewQBitmap3(int width, int height)
{
	return new MyQBitmap(width, height);
}

void* QBitmap_NewQBitmap4(void* size)
{
	return new MyQBitmap(*static_cast<QSize*>(size));
}

void* QBitmap_NewQBitmap5(struct QtGui_PackedString fileName, char* format)
{
	return new MyQBitmap(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format));
}

void QBitmap_Clear(void* ptr)
{
	static_cast<QBitmap*>(ptr)->clear();
}

void* QBitmap_QBitmap_FromData(void* size, char* bits, long long monoFormat)
{
	return new QBitmap(QBitmap::fromData(*static_cast<QSize*>(size), const_cast<const uchar*>(static_cast<uchar*>(static_cast<void*>(bits))), static_cast<QImage::Format>(monoFormat)));
}

void QBitmap_DestroyQBitmap(void* ptr)
{
	static_cast<QBitmap*>(ptr)->~QBitmap();
}

void QBitmap_DestroyQBitmapDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

Q_DECLARE_METATYPE(QBrush)
Q_DECLARE_METATYPE(QBrush*)
void* QBrush_NewQBrush()
{
	return new QBrush();
}

void* QBrush_NewQBrush2(long long style)
{
	return new QBrush(static_cast<Qt::BrushStyle>(style));
}

void* QBrush_NewQBrush3(void* color, long long style)
{
	return new QBrush(*static_cast<QColor*>(color), static_cast<Qt::BrushStyle>(style));
}

void* QBrush_NewQBrush4(long long color, long long style)
{
	return new QBrush(static_cast<Qt::GlobalColor>(color), static_cast<Qt::BrushStyle>(style));
}

void* QBrush_NewQBrush5(void* color, void* pixmap)
{
	return new QBrush(*static_cast<QColor*>(color), *static_cast<QPixmap*>(pixmap));
}

void* QBrush_NewQBrush6(long long color, void* pixmap)
{
	return new QBrush(static_cast<Qt::GlobalColor>(color), *static_cast<QPixmap*>(pixmap));
}

void* QBrush_NewQBrush7(void* pixmap)
{
	return new QBrush(*static_cast<QPixmap*>(pixmap));
}

void* QBrush_NewQBrush8(void* image)
{
	return new QBrush(*static_cast<QImage*>(image));
}

void* QBrush_NewQBrush9(void* other)
{
	return new QBrush(*static_cast<QBrush*>(other));
}

void* QBrush_NewQBrush10(void* gradient)
{
	return new QBrush(*static_cast<QGradient*>(gradient));
}

void* QBrush_Color(void* ptr)
{
	return const_cast<QColor*>(&static_cast<QBrush*>(ptr)->color());
}

void QBrush_SetStyle(void* ptr, long long style)
{
	static_cast<QBrush*>(ptr)->setStyle(static_cast<Qt::BrushStyle>(style));
}

long long QBrush_Style(void* ptr)
{
	return static_cast<QBrush*>(ptr)->style();
}

void* QBrush_Transform(void* ptr)
{
	return new QTransform(static_cast<QBrush*>(ptr)->transform());
}

void QBrush_DestroyQBrush(void* ptr)
{
	static_cast<QBrush*>(ptr)->~QBrush();
}

void* QBrush_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QBrush*>(ptr));
}

class MyQClipboard: public QClipboard
{
public:
	void Signal_Changed(QClipboard::Mode mode) { callbackQClipboard_Changed(this, mode); };
	void Signal_SelectionChanged() { callbackQClipboard_SelectionChanged(this); };
	void childEvent(QChildEvent * event) { callbackQClipboard_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQClipboard_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQClipboard_CustomEvent(this, event); };
	void deleteLater() { callbackQClipboard_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQClipboard_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQClipboard_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQClipboard_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQClipboard_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQClipboard_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQClipboard_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QClipboard*)
Q_DECLARE_METATYPE(MyQClipboard*)

int QClipboard_QClipboard_QRegisterMetaType(){qRegisterMetaType<QClipboard*>(); return qRegisterMetaType<MyQClipboard*>();}

void QClipboard_ConnectChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)(QClipboard::Mode)>(&QClipboard::changed), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)(QClipboard::Mode)>(&MyQClipboard::Signal_Changed), static_cast<Qt::ConnectionType>(t));
}

void QClipboard_DisconnectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)(QClipboard::Mode)>(&QClipboard::changed), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)(QClipboard::Mode)>(&MyQClipboard::Signal_Changed));
}

void QClipboard_Changed(void* ptr, long long mode)
{
	static_cast<QClipboard*>(ptr)->changed(static_cast<QClipboard::Mode>(mode));
}

void QClipboard_Clear(void* ptr, long long mode)
{
	static_cast<QClipboard*>(ptr)->clear(static_cast<QClipboard::Mode>(mode));
}

void* QClipboard_Image(void* ptr, long long mode)
{
	return new QImage(static_cast<QClipboard*>(ptr)->image(static_cast<QClipboard::Mode>(mode)));
}

void* QClipboard_Pixmap(void* ptr, long long mode)
{
	return new QPixmap(static_cast<QClipboard*>(ptr)->pixmap(static_cast<QClipboard::Mode>(mode)));
}

void QClipboard_ConnectSelectionChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::selectionChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_SelectionChanged), static_cast<Qt::ConnectionType>(t));
}

void QClipboard_DisconnectSelectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::selectionChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_SelectionChanged));
}

void QClipboard_SelectionChanged(void* ptr)
{
	static_cast<QClipboard*>(ptr)->selectionChanged();
}

void QClipboard_SetPixmap(void* ptr, void* pixmap, long long mode)
{
	static_cast<QClipboard*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QClipboard::Mode>(mode));
}

void QClipboard_SetText(void* ptr, struct QtGui_PackedString text, long long mode)
{
	static_cast<QClipboard*>(ptr)->setText(QString::fromUtf8(text.data, text.len), static_cast<QClipboard::Mode>(mode));
}

struct QtGui_PackedString QClipboard_Text(void* ptr, long long mode)
{
	return ({ QByteArray* tc36a83 = new QByteArray(static_cast<QClipboard*>(ptr)->text(static_cast<QClipboard::Mode>(mode)).toUtf8()); QtGui_PackedString { const_cast<char*>(tc36a83->prepend("WHITESPACE").constData()+10), tc36a83->size()-10, tc36a83 }; });
}

struct QtGui_PackedString QClipboard_Text2(void* ptr, struct QtGui_PackedString subtype, long long mode)
{
	return ({ QByteArray* tde8270 = new QByteArray(static_cast<QClipboard*>(ptr)->text(*(new QString(QString::fromUtf8(subtype.data, subtype.len))), static_cast<QClipboard::Mode>(mode)).toUtf8()); QtGui_PackedString { const_cast<char*>(tde8270->prepend("WHITESPACE").constData()+10), tde8270->size()-10, tde8270 }; });
}

void* QClipboard___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QClipboard___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QClipboard___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QClipboard___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QClipboard___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QClipboard___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QClipboard___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QClipboard___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QClipboard___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QClipboard___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QClipboard___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QClipboard___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QClipboard_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QClipboard*>(ptr)->QClipboard::childEvent(static_cast<QChildEvent*>(event));
}

void QClipboard_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QClipboard*>(ptr)->QClipboard::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QClipboard_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QClipboard*>(ptr)->QClipboard::customEvent(static_cast<QEvent*>(event));
}

void QClipboard_DeleteLaterDefault(void* ptr)
{
		static_cast<QClipboard*>(ptr)->QClipboard::deleteLater();
}

void QClipboard_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QClipboard*>(ptr)->QClipboard::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QClipboard_EventDefault(void* ptr, void* e)
{
		return static_cast<QClipboard*>(ptr)->QClipboard::event(static_cast<QEvent*>(e));
}

char QClipboard_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QClipboard*>(ptr)->QClipboard::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QClipboard*>(ptr)->QClipboard::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QClipboard_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QClipboard*>(ptr)->QClipboard::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQCloseEvent: public QCloseEvent
{
public:
	MyQCloseEvent() : QCloseEvent() {QCloseEvent_QCloseEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QCloseEvent*)
Q_DECLARE_METATYPE(MyQCloseEvent*)

int QCloseEvent_QCloseEvent_QRegisterMetaType(){qRegisterMetaType<QCloseEvent*>(); return qRegisterMetaType<MyQCloseEvent*>();}

void* QCloseEvent_NewQCloseEvent()
{
	return new MyQCloseEvent();
}

Q_DECLARE_METATYPE(QColor)
Q_DECLARE_METATYPE(QColor*)
void* QColor_NewQColor()
{
	return new QColor();
}

void* QColor_NewQColor2(long long color)
{
	return new QColor(static_cast<Qt::GlobalColor>(color));
}

void* QColor_NewQColor3(int r, int g, int b, int a)
{
	return new QColor(r, g, b, a);
}

void* QColor_NewQColor4(unsigned int color)
{
	return new QColor(color);
}

void* QColor_NewQColor5(void* rgba64)
{
	return new QColor(*static_cast<QRgba64*>(rgba64));
}

void* QColor_NewQColor6(struct QtGui_PackedString name)
{
	return new QColor(QString::fromUtf8(name.data, name.len));
}

void* QColor_NewQColor8(char* name)
{
	return new QColor(const_cast<const char*>(name));
}

void* QColor_NewQColor9(void* name)
{
	return new QColor(*static_cast<QLatin1String*>(name));
}

int QColor_Alpha(void* ptr)
{
	return static_cast<QColor*>(ptr)->alpha();
}

char QColor_IsValid(void* ptr)
{
	return static_cast<QColor*>(ptr)->isValid();
}

struct QtGui_PackedString QColor_Name(void* ptr)
{
	return ({ QByteArray* t9b3be4 = new QByteArray(static_cast<QColor*>(ptr)->name().toUtf8()); QtGui_PackedString { const_cast<char*>(t9b3be4->prepend("WHITESPACE").constData()+10), t9b3be4->size()-10, t9b3be4 }; });
}

struct QtGui_PackedString QColor_Name2(void* ptr, long long format)
{
	return ({ QByteArray* t4331f3 = new QByteArray(static_cast<QColor*>(ptr)->name(static_cast<QColor::NameFormat>(format)).toUtf8()); QtGui_PackedString { const_cast<char*>(t4331f3->prepend("WHITESPACE").constData()+10), t4331f3->size()-10, t4331f3 }; });
}

int QColor_Value(void* ptr)
{
	return static_cast<QColor*>(ptr)->value();
}

void* QColor_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QColor*>(ptr));
}

class MyQContextMenuEvent: public QContextMenuEvent
{
public:
	MyQContextMenuEvent(QContextMenuEvent::Reason reason, const QPoint &pos, const QPoint &globalPos, Qt::KeyboardModifiers modifiers) : QContextMenuEvent(reason, pos, globalPos, modifiers) {QContextMenuEvent_QContextMenuEvent_QRegisterMetaType();};
	MyQContextMenuEvent(QContextMenuEvent::Reason reason, const QPoint &pos, const QPoint &globalPos) : QContextMenuEvent(reason, pos, globalPos) {QContextMenuEvent_QContextMenuEvent_QRegisterMetaType();};
	MyQContextMenuEvent(QContextMenuEvent::Reason reason, const QPoint &pos) : QContextMenuEvent(reason, pos) {QContextMenuEvent_QContextMenuEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QContextMenuEvent*)
Q_DECLARE_METATYPE(MyQContextMenuEvent*)

int QContextMenuEvent_QContextMenuEvent_QRegisterMetaType(){qRegisterMetaType<QContextMenuEvent*>(); return qRegisterMetaType<MyQContextMenuEvent*>();}

void* QContextMenuEvent_NewQContextMenuEvent(long long reason, void* pos, void* globalPos, long long modifiers)
{
	return new MyQContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QContextMenuEvent_NewQContextMenuEvent2(long long reason, void* pos, void* globalPos)
{
	return new MyQContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos));
}

void* QContextMenuEvent_NewQContextMenuEvent3(long long reason, void* pos)
{
	return new MyQContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos));
}

void* QContextMenuEvent_GlobalPos(void* ptr)
{
	return const_cast<QPoint*>(&static_cast<QContextMenuEvent*>(ptr)->globalPos());
}

void* QContextMenuEvent_Pos(void* ptr)
{
	return const_cast<QPoint*>(&static_cast<QContextMenuEvent*>(ptr)->pos());
}

int QContextMenuEvent_X(void* ptr)
{
	return static_cast<QContextMenuEvent*>(ptr)->x();
}

int QContextMenuEvent_Y(void* ptr)
{
	return static_cast<QContextMenuEvent*>(ptr)->y();
}

Q_DECLARE_METATYPE(QCursor)
Q_DECLARE_METATYPE(QCursor*)
void* QCursor_NewQCursor()
{
	return new QCursor();
}

void* QCursor_NewQCursor2(long long shape)
{
	return new QCursor(static_cast<Qt::CursorShape>(shape));
}

void* QCursor_NewQCursor3(void* bitmap, void* mask, int hotX, int hotY)
{
	return new QCursor(*static_cast<QBitmap*>(bitmap), *static_cast<QBitmap*>(mask), hotX, hotY);
}

void* QCursor_NewQCursor4(void* pixmap, int hotX, int hotY)
{
	return new QCursor(*static_cast<QPixmap*>(pixmap), hotX, hotY);
}

void* QCursor_NewQCursor5(void* c)
{
	return new QCursor(*static_cast<QCursor*>(c));
}

void* QCursor_NewQCursor6(void* other)
{
	return new QCursor(*static_cast<QCursor*>(other));
}

void* QCursor_Pixmap(void* ptr)
{
	return new QPixmap(static_cast<QCursor*>(ptr)->pixmap());
}

void* QCursor_QCursor_Pos()
{
	return ({ QPoint tmpValue = QCursor::pos(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QCursor_QCursor_Pos2(void* screen)
{
	return ({ QPoint tmpValue = QCursor::pos(static_cast<QScreen*>(screen)); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QCursor_QCursor_SetPos(int x, int y)
{
	QCursor::setPos(x, y);
}

void QCursor_QCursor_SetPos2(void* screen, int x, int y)
{
	QCursor::setPos(static_cast<QScreen*>(screen), x, y);
}

void QCursor_QCursor_SetPos3(void* p)
{
	QCursor::setPos(*static_cast<QPoint*>(p));
}

void QCursor_QCursor_SetPos4(void* screen, void* p)
{
	QCursor::setPos(static_cast<QScreen*>(screen), *static_cast<QPoint*>(p));
}

void QCursor_DestroyQCursor(void* ptr)
{
	static_cast<QCursor*>(ptr)->~QCursor();
}

Q_DECLARE_METATYPE(QFont)
Q_DECLARE_METATYPE(QFont*)
void* QFont_NewQFont()
{
	return new QFont();
}

void* QFont_NewQFont2(struct QtGui_PackedString family, int pointSize, int weight, char italic)
{
	return new QFont(QString::fromUtf8(family.data, family.len), pointSize, weight, italic != 0);
}

void* QFont_NewQFont4(void* font, void* pd)
{
	if (dynamic_cast<QWidget*>(static_cast<QObject*>(pd))) {
		return new QFont(*static_cast<QFont*>(font), static_cast<QWidget*>(pd));
	} else {
		return new QFont(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(pd));
	}
}

void* QFont_NewQFont5(void* font)
{
	return new QFont(*static_cast<QFont*>(font));
}

char QFont_Bold(void* ptr)
{
	return static_cast<QFont*>(ptr)->bold();
}

struct QtGui_PackedString QFont_Family(void* ptr)
{
	return ({ QByteArray* t52247e = new QByteArray(static_cast<QFont*>(ptr)->family().toUtf8()); QtGui_PackedString { const_cast<char*>(t52247e->prepend("WHITESPACE").constData()+10), t52247e->size()-10, t52247e }; });
}

char QFont_FromString(void* ptr, struct QtGui_PackedString descrip)
{
	return static_cast<QFont*>(ptr)->fromString(QString::fromUtf8(descrip.data, descrip.len));
}

char QFont_Italic(void* ptr)
{
	return static_cast<QFont*>(ptr)->italic();
}

struct QtGui_PackedString QFont_Key(void* ptr)
{
	return ({ QByteArray* t9324a4 = new QByteArray(static_cast<QFont*>(ptr)->key().toUtf8()); QtGui_PackedString { const_cast<char*>(t9324a4->prepend("WHITESPACE").constData()+10), t9324a4->size()-10, t9324a4 }; });
}

int QFont_PointSize(void* ptr)
{
	return static_cast<QFont*>(ptr)->pointSize();
}

void QFont_SetFamily(void* ptr, struct QtGui_PackedString family)
{
	static_cast<QFont*>(ptr)->setFamily(QString::fromUtf8(family.data, family.len));
}

void QFont_SetPointSize(void* ptr, int pointSize)
{
	static_cast<QFont*>(ptr)->setPointSize(pointSize);
}

void QFont_SetStretch(void* ptr, int factor)
{
	static_cast<QFont*>(ptr)->setStretch(factor);
}

void QFont_SetStyle(void* ptr, long long style)
{
	static_cast<QFont*>(ptr)->setStyle(static_cast<QFont::Style>(style));
}

void QFont_SetStyleName(void* ptr, struct QtGui_PackedString styleName)
{
	static_cast<QFont*>(ptr)->setStyleName(QString::fromUtf8(styleName.data, styleName.len));
}

int QFont_Stretch(void* ptr)
{
	return static_cast<QFont*>(ptr)->stretch();
}

char QFont_StrikeOut(void* ptr)
{
	return static_cast<QFont*>(ptr)->strikeOut();
}

long long QFont_Style(void* ptr)
{
	return static_cast<QFont*>(ptr)->style();
}

struct QtGui_PackedString QFont_StyleName(void* ptr)
{
	return ({ QByteArray* t8d3474 = new QByteArray(static_cast<QFont*>(ptr)->styleName().toUtf8()); QtGui_PackedString { const_cast<char*>(t8d3474->prepend("WHITESPACE").constData()+10), t8d3474->size()-10, t8d3474 }; });
}

char QFont_Underline(void* ptr)
{
	return static_cast<QFont*>(ptr)->underline();
}

int QFont_Weight(void* ptr)
{
	return static_cast<QFont*>(ptr)->weight();
}

void QFont_DestroyQFont(void* ptr)
{
	static_cast<QFont*>(ptr)->~QFont();
}

void* QFont_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QFont*>(ptr));
}

Q_DECLARE_METATYPE(QFontMetrics*)
void* QFontMetrics_NewQFontMetrics(void* font)
{
	return new QFontMetrics(*static_cast<QFont*>(font));
}

void* QFontMetrics_NewQFontMetrics3(void* font, void* paintdevice)
{
	if (dynamic_cast<QWidget*>(static_cast<QObject*>(paintdevice))) {
		return new QFontMetrics(*static_cast<QFont*>(font), static_cast<QWidget*>(paintdevice));
	} else {
		return new QFontMetrics(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
	}
}

void* QFontMetrics_NewQFontMetrics4(void* fm)
{
	return new QFontMetrics(*static_cast<QFontMetrics*>(fm));
}

void* QFontMetrics_BoundingRect(void* ptr, void* ch)
{
	return ({ QRect tmpValue = static_cast<QFontMetrics*>(ptr)->boundingRect(*static_cast<QChar*>(ch)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QFontMetrics_BoundingRect2(void* ptr, struct QtGui_PackedString text)
{
	return ({ QRect tmpValue = static_cast<QFontMetrics*>(ptr)->boundingRect(QString::fromUtf8(text.data, text.len)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QFontMetrics_BoundingRect3(void* ptr, void* rect, int flags, struct QtGui_PackedString text, int tabStops, int tabArray)
{
	return ({ QRect tmpValue = static_cast<QFontMetrics*>(ptr)->boundingRect(*static_cast<QRect*>(rect), flags, QString::fromUtf8(text.data, text.len), tabStops, &tabArray); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QFontMetrics_BoundingRect4(void* ptr, int x, int y, int width, int height, int flags, struct QtGui_PackedString text, int tabStops, int tabArray)
{
	return ({ QRect tmpValue = static_cast<QFontMetrics*>(ptr)->boundingRect(x, y, width, height, flags, QString::fromUtf8(text.data, text.len), tabStops, &tabArray); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

int QFontMetrics_Height(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->height();
}

void* QFontMetrics_Size(void* ptr, int flags, struct QtGui_PackedString text, int tabStops, int tabArray)
{
	return ({ QSize tmpValue = static_cast<QFontMetrics*>(ptr)->size(flags, QString::fromUtf8(text.data, text.len), tabStops, &tabArray); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QFontMetrics_DestroyQFontMetrics(void* ptr)
{
	static_cast<QFontMetrics*>(ptr)->~QFontMetrics();
}

Q_DECLARE_METATYPE(QGlyphRun)
Q_DECLARE_METATYPE(QGlyphRun*)
void* QGlyphRun_NewQGlyphRun()
{
	return new QGlyphRun();
}

void* QGlyphRun_NewQGlyphRun2(void* other)
{
	return new QGlyphRun(*static_cast<QGlyphRun*>(other));
}

void* QGlyphRun_BoundingRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QGlyphRun*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QGlyphRun_Clear(void* ptr)
{
	static_cast<QGlyphRun*>(ptr)->clear();
}

char QGlyphRun_IsEmpty(void* ptr)
{
	return static_cast<QGlyphRun*>(ptr)->isEmpty();
}

char QGlyphRun_StrikeOut(void* ptr)
{
	return static_cast<QGlyphRun*>(ptr)->strikeOut();
}

char QGlyphRun_Underline(void* ptr)
{
	return static_cast<QGlyphRun*>(ptr)->underline();
}

void QGlyphRun_DestroyQGlyphRun(void* ptr)
{
	static_cast<QGlyphRun*>(ptr)->~QGlyphRun();
}

unsigned int QGlyphRun___glyphIndexes_atList(void* ptr, int i)
{
	return ({quint32 tmp = static_cast<QVector<quint32>*>(ptr)->at(i); if (i == static_cast<QVector<quint32>*>(ptr)->size()-1) { static_cast<QVector<quint32>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QGlyphRun___glyphIndexes_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<quint32>*>(ptr)->append(i);
}

void* QGlyphRun___glyphIndexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<quint32>();
}

void* QGlyphRun___positions_atList(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QGlyphRun___positions_setList(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QGlyphRun___positions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

unsigned int QGlyphRun___setGlyphIndexes_glyphIndexes_atList(void* ptr, int i)
{
	return ({quint32 tmp = static_cast<QVector<quint32>*>(ptr)->at(i); if (i == static_cast<QVector<quint32>*>(ptr)->size()-1) { static_cast<QVector<quint32>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QGlyphRun___setGlyphIndexes_glyphIndexes_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<quint32>*>(ptr)->append(i);
}

void* QGlyphRun___setGlyphIndexes_glyphIndexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<quint32>();
}

void* QGlyphRun___setPositions_positions_atList(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QGlyphRun___setPositions_positions_setList(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QGlyphRun___setPositions_positions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

Q_DECLARE_METATYPE(QGradient*)
void* QGradient_NewQGradient2(long long preset)
{
	return new QGradient(static_cast<QGradient::Preset>(preset));
}

long long QGradient_Type(void* ptr)
{
	return static_cast<QGradient*>(ptr)->type();
}

class MyQGuiApplication: public QGuiApplication
{
public:
	MyQGuiApplication(int &argc, char **argv) : QGuiApplication(argc, argv) {QGuiApplication_QGuiApplication_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQGuiApplication_Event(this, e) != 0; };
	 ~MyQGuiApplication() { callbackQGuiApplication_DestroyQGuiApplication(this); };
	void quit() { callbackQGuiApplication_Quit(this); };
	void childEvent(QChildEvent * event) { callbackQGuiApplication_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGuiApplication_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGuiApplication_CustomEvent(this, event); };
	void deleteLater() { callbackQGuiApplication_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGuiApplication_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGuiApplication_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGuiApplication_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGuiApplication_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGuiApplication_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QGuiApplication*)
Q_DECLARE_METATYPE(MyQGuiApplication*)

int QGuiApplication_QGuiApplication_QRegisterMetaType(){qRegisterMetaType<QGuiApplication*>(); return qRegisterMetaType<MyQGuiApplication*>();}

void* QGuiApplication_NewQGuiApplication(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new MyQGuiApplication(argcs, argvs);
}

struct QtGui_PackedString QGuiApplication_QGuiApplication_ApplicationDisplayName()
{
	return ({ QByteArray* tb2b492 = new QByteArray(QGuiApplication::applicationDisplayName().toUtf8()); QtGui_PackedString { const_cast<char*>(tb2b492->prepend("WHITESPACE").constData()+10), tb2b492->size()-10, tb2b492 }; });
}

void* QGuiApplication_QGuiApplication_Clipboard()
{
	return QGuiApplication::clipboard();
}

char QGuiApplication_EventDefault(void* ptr, void* e)
{
		return static_cast<QGuiApplication*>(ptr)->QGuiApplication::event(static_cast<QEvent*>(e));
}

int QGuiApplication_QGuiApplication_Exec()
{
	return QGuiApplication::exec();
}

void* QGuiApplication_QGuiApplication_Font()
{
	return new QFont(QGuiApplication::font());
}

void QGuiApplication_QGuiApplication_SetApplicationDisplayName(struct QtGui_PackedString name)
{
	QGuiApplication::setApplicationDisplayName(QString::fromUtf8(name.data, name.len));
}

void QGuiApplication_QGuiApplication_SetFont(void* font)
{
	QGuiApplication::setFont(*static_cast<QFont*>(font));
}

void QGuiApplication_QGuiApplication_SetWindowIcon(void* icon)
{
	QGuiApplication::setWindowIcon(*static_cast<QIcon*>(icon));
}

void* QGuiApplication_QGuiApplication_WindowIcon()
{
	return new QIcon(QGuiApplication::windowIcon());
}

void QGuiApplication_DestroyQGuiApplication(void* ptr)
{
	static_cast<QGuiApplication*>(ptr)->~QGuiApplication();
}

void QGuiApplication_DestroyQGuiApplicationDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGuiApplication___screens_atList(void* ptr, int i)
{
	return ({QScreen * tmp = static_cast<QList<QScreen *>*>(ptr)->at(i); if (i == static_cast<QList<QScreen *>*>(ptr)->size()-1) { static_cast<QList<QScreen *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___screens_setList(void* ptr, void* i)
{
	static_cast<QList<QScreen *>*>(ptr)->append(static_cast<QScreen*>(i));
}

void* QGuiApplication___screens_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QScreen *>();
}

void* QGuiApplication___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGuiApplication___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QGuiApplication___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGuiApplication___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGuiApplication___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGuiApplication___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGuiApplication___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGuiApplication___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGuiApplication___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QGuiApplication_QuitDefault(void* ptr)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::quit();
}

void QGuiApplication_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::childEvent(static_cast<QChildEvent*>(event));
}

void QGuiApplication_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGuiApplication_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::customEvent(static_cast<QEvent*>(event));
}

void QGuiApplication_DeleteLaterDefault(void* ptr)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::deleteLater();
}

void QGuiApplication_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGuiApplication_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QGuiApplication*>(ptr)->QGuiApplication::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QGuiApplication*>(ptr)->QGuiApplication::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QGuiApplication_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQHideEvent: public QHideEvent
{
public:
	MyQHideEvent() : QHideEvent() {QHideEvent_QHideEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QHideEvent*)
Q_DECLARE_METATYPE(MyQHideEvent*)

int QHideEvent_QHideEvent_QRegisterMetaType(){qRegisterMetaType<QHideEvent*>(); return qRegisterMetaType<MyQHideEvent*>();}

void* QHideEvent_NewQHideEvent()
{
	return new MyQHideEvent();
}

Q_DECLARE_METATYPE(QIcon)
Q_DECLARE_METATYPE(QIcon*)
void* QIcon_NewQIcon()
{
	return new QIcon();
}

void* QIcon_NewQIcon2(void* pixmap)
{
	return new QIcon(*static_cast<QPixmap*>(pixmap));
}

void* QIcon_NewQIcon3(void* other)
{
	return new QIcon(*static_cast<QIcon*>(other));
}

void* QIcon_NewQIcon4(void* other)
{
	return new QIcon(*static_cast<QIcon*>(other));
}

void* QIcon_NewQIcon5(struct QtGui_PackedString fileName)
{
	return new QIcon(QString::fromUtf8(fileName.data, fileName.len));
}

void* QIcon_NewQIcon6(void* engine)
{
	return new QIcon(static_cast<QIconEngine*>(engine));
}

void* QIcon_QIcon_FromTheme(struct QtGui_PackedString name)
{
	return new QIcon(QIcon::fromTheme(QString::fromUtf8(name.data, name.len)));
}

void* QIcon_QIcon_FromTheme2(struct QtGui_PackedString name, void* fallback)
{
	return new QIcon(QIcon::fromTheme(QString::fromUtf8(name.data, name.len), *static_cast<QIcon*>(fallback)));
}

struct QtGui_PackedString QIcon_Name(void* ptr)
{
	return ({ QByteArray* t03700a = new QByteArray(static_cast<QIcon*>(ptr)->name().toUtf8()); QtGui_PackedString { const_cast<char*>(t03700a->prepend("WHITESPACE").constData()+10), t03700a->size()-10, t03700a }; });
}

void* QIcon_Pixmap(void* ptr, void* size, long long mode, long long state)
{
	return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIcon_Pixmap2(void* ptr, int w, int h, long long mode, long long state)
{
	return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(w, h, static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIcon_Pixmap3(void* ptr, int extent, long long mode, long long state)
{
	return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(extent, static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIcon_Pixmap4(void* ptr, void* window, void* size, long long mode, long long state)
{
		return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(static_cast<QWindow*>(window), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void QIcon_DestroyQIcon(void* ptr)
{
	static_cast<QIcon*>(ptr)->~QIcon();
}

void* QIcon_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QIcon*>(ptr));
}

void* QIcon___availableSizes_atList(void* ptr, int i)
{
	return ({ QSize tmpValue = ({QSize tmp = static_cast<QList<QSize>*>(ptr)->at(i); if (i == static_cast<QList<QSize>*>(ptr)->size()-1) { static_cast<QList<QSize>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QIcon___availableSizes_setList(void* ptr, void* i)
{
	static_cast<QList<QSize>*>(ptr)->append(*static_cast<QSize*>(i));
}

void* QIcon___availableSizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSize>();
}

class MyQIconEngine: public QIconEngine
{
public:
	MyQIconEngine() : QIconEngine() {QIconEngine_QIconEngine_QRegisterMetaType();};
	QIconEngine * clone() const { return static_cast<QIconEngine*>(callbackQIconEngine_Clone(const_cast<void*>(static_cast<const void*>(this)))); };
	QString key() const { return ({ QtGui_PackedString tempVal = callbackQIconEngine_Key(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void paint(QPainter * painter, const QRect & rect, QIcon::Mode mode, QIcon::State state) { callbackQIconEngine_Paint(this, painter, const_cast<QRect*>(&rect), mode, state); };
	QPixmap pixmap(const QSize & size, QIcon::Mode mode, QIcon::State state) { return *static_cast<QPixmap*>(callbackQIconEngine_Pixmap(this, const_cast<QSize*>(&size), mode, state)); };
	bool read(QDataStream & in) { return callbackQIconEngine_Read(this, static_cast<QDataStream*>(&in)) != 0; };
	bool write(QDataStream & out) const { return callbackQIconEngine_Write(const_cast<void*>(static_cast<const void*>(this)), static_cast<QDataStream*>(&out)) != 0; };
	 ~MyQIconEngine() { callbackQIconEngine_DestroyQIconEngine(this); };
};

Q_DECLARE_METATYPE(QIconEngine*)
Q_DECLARE_METATYPE(MyQIconEngine*)

int QIconEngine_QIconEngine_QRegisterMetaType(){qRegisterMetaType<QIconEngine*>(); return qRegisterMetaType<MyQIconEngine*>();}

void* QIconEngine_NewQIconEngine()
{
	return new MyQIconEngine();
}

void* QIconEngine_Clone(void* ptr)
{
	return static_cast<QIconEngine*>(ptr)->clone();
}

struct QtGui_PackedString QIconEngine_Key(void* ptr)
{
	return ({ QByteArray* tfa2543 = new QByteArray(static_cast<QIconEngine*>(ptr)->key().toUtf8()); QtGui_PackedString { const_cast<char*>(tfa2543->prepend("WHITESPACE").constData()+10), tfa2543->size()-10, tfa2543 }; });
}

struct QtGui_PackedString QIconEngine_KeyDefault(void* ptr)
{
		return ({ QByteArray* t9979b6 = new QByteArray(static_cast<QIconEngine*>(ptr)->QIconEngine::key().toUtf8()); QtGui_PackedString { const_cast<char*>(t9979b6->prepend("WHITESPACE").constData()+10), t9979b6->size()-10, t9979b6 }; });
}

void QIconEngine_Paint(void* ptr, void* painter, void* rect, long long mode, long long state)
{
	static_cast<QIconEngine*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void* QIconEngine_Pixmap(void* ptr, void* size, long long mode, long long state)
{
	return new QPixmap(static_cast<QIconEngine*>(ptr)->pixmap(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIconEngine_PixmapDefault(void* ptr, void* size, long long mode, long long state)
{
		return new QPixmap(static_cast<QIconEngine*>(ptr)->QIconEngine::pixmap(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

char QIconEngine_Read(void* ptr, void* in)
{
	return static_cast<QIconEngine*>(ptr)->read(*static_cast<QDataStream*>(in));
}

char QIconEngine_ReadDefault(void* ptr, void* in)
{
		return static_cast<QIconEngine*>(ptr)->QIconEngine::read(*static_cast<QDataStream*>(in));
}

char QIconEngine_Write(void* ptr, void* out)
{
	return static_cast<QIconEngine*>(ptr)->write(*static_cast<QDataStream*>(out));
}

char QIconEngine_WriteDefault(void* ptr, void* out)
{
		return static_cast<QIconEngine*>(ptr)->QIconEngine::write(*static_cast<QDataStream*>(out));
}

void QIconEngine_DestroyQIconEngine(void* ptr)
{
	static_cast<QIconEngine*>(ptr)->~QIconEngine();
}

void QIconEngine_DestroyQIconEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QIconEngine___availableSizes_atList(void* ptr, int i)
{
	return ({ QSize tmpValue = ({QSize tmp = static_cast<QList<QSize>*>(ptr)->at(i); if (i == static_cast<QList<QSize>*>(ptr)->size()-1) { static_cast<QList<QSize>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QIconEngine___availableSizes_setList(void* ptr, void* i)
{
	static_cast<QList<QSize>*>(ptr)->append(*static_cast<QSize*>(i));
}

void* QIconEngine___availableSizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSize>();
}

class MyQImage: public QImage
{
public:
	MyQImage() : QImage() {QImage_QImage_QRegisterMetaType();};
	MyQImage(const QSize &size, QImage::Format format) : QImage(size, format) {QImage_QImage_QRegisterMetaType();};
	MyQImage(int width, int height, QImage::Format format) : QImage(width, height, format) {QImage_QImage_QRegisterMetaType();};
	MyQImage(uchar *data, int width, int height, QImage::Format format) : QImage(data, width, height, format) {QImage_QImage_QRegisterMetaType();};
	MyQImage(const uchar *data, int width, int height, QImage::Format format) : QImage(data, width, height, format) {QImage_QImage_QRegisterMetaType();};
	MyQImage(uchar *data, int width, int height, int bytesPerLine, QImage::Format format) : QImage(data, width, height, bytesPerLine, format) {QImage_QImage_QRegisterMetaType();};
	MyQImage(const uchar *data, int width, int height, int bytesPerLine, QImage::Format format) : QImage(data, width, height, bytesPerLine, format) {QImage_QImage_QRegisterMetaType();};
	MyQImage(const QString &fileName, const char *format = Q_NULLPTR) : QImage(fileName, format) {QImage_QImage_QRegisterMetaType();};
	MyQImage(const QImage &image) : QImage(image) {QImage_QImage_QRegisterMetaType();};
	MyQImage(QImage &&other) : QImage(other) {QImage_QImage_QRegisterMetaType();};
	 ~MyQImage() { callbackQImage_DestroyQImage(this); };
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPaintDevice_Metric(const_cast<void*>(static_cast<const void*>(this)), metric); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQImage_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QImage*)
Q_DECLARE_METATYPE(MyQImage*)

int QImage_QImage_QRegisterMetaType(){qRegisterMetaType<QImage*>(); return qRegisterMetaType<MyQImage*>();}

void* QImage_NewQImage()
{
	return new MyQImage();
}

void* QImage_NewQImage2(void* size, long long format)
{
	return new MyQImage(*static_cast<QSize*>(size), static_cast<QImage::Format>(format));
}

void* QImage_NewQImage3(int width, int height, long long format)
{
	return new MyQImage(width, height, static_cast<QImage::Format>(format));
}

void* QImage_NewQImage4(char* data, int width, int height, long long format)
{
	return new MyQImage(static_cast<uchar*>(static_cast<void*>(data)), width, height, static_cast<QImage::Format>(format));
}

void* QImage_NewQImage5(char* data, int width, int height, long long format)
{
	return new MyQImage(const_cast<const uchar*>(static_cast<uchar*>(static_cast<void*>(data))), width, height, static_cast<QImage::Format>(format));
}

void* QImage_NewQImage6(char* data, int width, int height, int bytesPerLine, long long format)
{
	return new MyQImage(static_cast<uchar*>(static_cast<void*>(data)), width, height, bytesPerLine, static_cast<QImage::Format>(format));
}

void* QImage_NewQImage7(char* data, int width, int height, int bytesPerLine, long long format)
{
	return new MyQImage(const_cast<const uchar*>(static_cast<uchar*>(static_cast<void*>(data))), width, height, bytesPerLine, static_cast<QImage::Format>(format));
}

void* QImage_NewQImage9(struct QtGui_PackedString fileName, char* format)
{
	return new MyQImage(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format));
}

void* QImage_NewQImage10(void* image)
{
	return new MyQImage(*static_cast<QImage*>(image));
}

void* QImage_NewQImage11(void* other)
{
	return new MyQImage(*static_cast<QImage*>(other));
}

unsigned int QImage_Color(void* ptr, int i)
{
	return static_cast<QImage*>(ptr)->color(i);
}

void* QImage_Copy(void* ptr, void* rectangle)
{
	return new QImage(static_cast<QImage*>(ptr)->copy(*static_cast<QRect*>(rectangle)));
}

void* QImage_Copy2(void* ptr, int x, int y, int width, int height)
{
	return new QImage(static_cast<QImage*>(ptr)->copy(x, y, width, height));
}

long long QImage_Format(void* ptr)
{
	return static_cast<QImage*>(ptr)->format();
}

void* QImage_QImage_FromData(char* data, int size, char* format)
{
	return new QImage(QImage::fromData(const_cast<const uchar*>(static_cast<uchar*>(static_cast<void*>(data))), size, const_cast<const char*>(format)));
}

void* QImage_QImage_FromData2(void* data, char* format)
{
	return new QImage(QImage::fromData(*static_cast<QByteArray*>(data), const_cast<const char*>(format)));
}

char QImage_Load(void* ptr, struct QtGui_PackedString fileName, char* format)
{
	return static_cast<QImage*>(ptr)->load(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format));
}

char QImage_Load2(void* ptr, void* device, char* format)
{
	return static_cast<QImage*>(ptr)->load(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

char QImage_LoadFromData(void* ptr, char* data, int l, char* format)
{
	return static_cast<QImage*>(ptr)->loadFromData(const_cast<const uchar*>(static_cast<uchar*>(static_cast<void*>(data))), l, const_cast<const char*>(format));
}

char QImage_LoadFromData2(void* ptr, void* data, char* format)
{
	return static_cast<QImage*>(ptr)->loadFromData(*static_cast<QByteArray*>(data), const_cast<const char*>(format));
}

void* QImage_Rect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QImage*>(ptr)->rect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

char QImage_Save(void* ptr, struct QtGui_PackedString fileName, char* format, int quality)
{
	return static_cast<QImage*>(ptr)->save(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format), quality);
}

char QImage_Save2(void* ptr, void* device, char* format, int quality)
{
	return static_cast<QImage*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format), quality);
}

void* QImage_Scaled(void* ptr, void* size, long long aspectRatioMode, long long transformMode)
{
	return new QImage(static_cast<QImage*>(ptr)->scaled(*static_cast<QSize*>(size), static_cast<Qt::AspectRatioMode>(aspectRatioMode), static_cast<Qt::TransformationMode>(transformMode)));
}

void* QImage_Scaled2(void* ptr, int width, int height, long long aspectRatioMode, long long transformMode)
{
	return new QImage(static_cast<QImage*>(ptr)->scaled(width, height, static_cast<Qt::AspectRatioMode>(aspectRatioMode), static_cast<Qt::TransformationMode>(transformMode)));
}

void* QImage_ScaledToHeight(void* ptr, int height, long long mode)
{
	return new QImage(static_cast<QImage*>(ptr)->scaledToHeight(height, static_cast<Qt::TransformationMode>(mode)));
}

void QImage_SetText(void* ptr, struct QtGui_PackedString key, struct QtGui_PackedString text)
{
	static_cast<QImage*>(ptr)->setText(QString::fromUtf8(key.data, key.len), QString::fromUtf8(text.data, text.len));
}

void* QImage_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QImage*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtGui_PackedString QImage_Text(void* ptr, struct QtGui_PackedString key)
{
	return ({ QByteArray* t3cc0e5 = new QByteArray(static_cast<QImage*>(ptr)->text(QString::fromUtf8(key.data, key.len)).toUtf8()); QtGui_PackedString { const_cast<char*>(t3cc0e5->prepend("WHITESPACE").constData()+10), t3cc0e5->size()-10, t3cc0e5 }; });
}

char QImage_Valid(void* ptr, void* pos)
{
	return static_cast<QImage*>(ptr)->valid(*static_cast<QPoint*>(pos));
}

char QImage_Valid2(void* ptr, int x, int y)
{
	return static_cast<QImage*>(ptr)->valid(x, y);
}

void QImage_DestroyQImage(void* ptr)
{
	static_cast<QImage*>(ptr)->~QImage();
}

void QImage_DestroyQImageDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QImage_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QImage*>(ptr));
}

unsigned int QImage___colorTable_atList(void* ptr, int i)
{
	return ({QRgb tmp = static_cast<QVector<QRgb>*>(ptr)->at(i); if (i == static_cast<QVector<QRgb>*>(ptr)->size()-1) { static_cast<QVector<QRgb>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QImage___colorTable_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<QRgb>*>(ptr)->append(i);
}

void* QImage___colorTable_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRgb>();
}

unsigned int QImage___convertToFormat_colorTable_atList3(void* ptr, int i)
{
	return ({QRgb tmp = static_cast<QVector<QRgb>*>(ptr)->at(i); if (i == static_cast<QVector<QRgb>*>(ptr)->size()-1) { static_cast<QVector<QRgb>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QImage___convertToFormat_colorTable_setList3(void* ptr, unsigned int i)
{
	static_cast<QVector<QRgb>*>(ptr)->append(i);
}

void* QImage___convertToFormat_colorTable_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRgb>();
}

unsigned int QImage___setColorTable_colors_atList(void* ptr, int i)
{
	return ({QRgb tmp = static_cast<QVector<QRgb>*>(ptr)->at(i); if (i == static_cast<QVector<QRgb>*>(ptr)->size()-1) { static_cast<QVector<QRgb>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QImage___setColorTable_colors_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<QRgb>*>(ptr)->append(i);
}

void* QImage___setColorTable_colors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRgb>();
}

void* QImage_PaintEngine(void* ptr)
{
	return static_cast<QImage*>(ptr)->paintEngine();
}

void* QImage_PaintEngineDefault(void* ptr)
{
		return static_cast<QImage*>(ptr)->QImage::paintEngine();
}

class MyQInputEvent: public QInputEvent
{
public:
};

Q_DECLARE_METATYPE(QInputEvent*)
Q_DECLARE_METATYPE(MyQInputEvent*)

int QInputEvent_QInputEvent_QRegisterMetaType(){qRegisterMetaType<QInputEvent*>(); return qRegisterMetaType<MyQInputEvent*>();}

class MyQIntValidator: public QIntValidator
{
public:
	MyQIntValidator(QObject *parent = Q_NULLPTR) : QIntValidator(parent) {QIntValidator_QIntValidator_QRegisterMetaType();};
	MyQIntValidator(int minimum, int maximum, QObject *parent = Q_NULLPTR) : QIntValidator(minimum, maximum, parent) {QIntValidator_QIntValidator_QRegisterMetaType();};
	void setRange(int bottom, int top) { callbackQIntValidator_SetRange(this, bottom, top); };
	 ~MyQIntValidator() { callbackQIntValidator_DestroyQIntValidator(this); };
	void Signal_Changed() { callbackQValidator_Changed(this); };
	QValidator::State validate(QString & input, int & pos) const { QByteArray* t140f86 = new QByteArray(input.toUtf8()); QtGui_PackedString inputPacked = { const_cast<char*>(t140f86->prepend("WHITESPACE").constData()+10), t140f86->size()-10, t140f86 };return static_cast<QValidator::State>(callbackQIntValidator_Validate(const_cast<void*>(static_cast<const void*>(this)), inputPacked, pos)); };
	void childEvent(QChildEvent * event) { callbackQValidator_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQValidator_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQValidator_CustomEvent(this, event); };
	void deleteLater() { callbackQValidator_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQValidator_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQValidator_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQValidator_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQValidator_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQValidator_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQValidator_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QIntValidator*)
Q_DECLARE_METATYPE(MyQIntValidator*)

int QIntValidator_QIntValidator_QRegisterMetaType(){qRegisterMetaType<QIntValidator*>(); return qRegisterMetaType<MyQIntValidator*>();}

void* QIntValidator_NewQIntValidator(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQIntValidator(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIntValidator(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQIntValidator(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIntValidator(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQIntValidator(static_cast<QWindow*>(parent));
	} else {
		return new MyQIntValidator(static_cast<QObject*>(parent));
	}
}

void* QIntValidator_NewQIntValidator2(int minimum, int maximum, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQIntValidator(minimum, maximum, static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIntValidator(minimum, maximum, static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQIntValidator(minimum, maximum, static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIntValidator(minimum, maximum, static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQIntValidator(minimum, maximum, static_cast<QWindow*>(parent));
	} else {
		return new MyQIntValidator(minimum, maximum, static_cast<QObject*>(parent));
	}
}

void QIntValidator_SetRange(void* ptr, int bottom, int top)
{
	static_cast<QIntValidator*>(ptr)->setRange(bottom, top);
}

void QIntValidator_SetRangeDefault(void* ptr, int bottom, int top)
{
		static_cast<QIntValidator*>(ptr)->QIntValidator::setRange(bottom, top);
}

void QIntValidator_SetTop(void* ptr, int vin)
{
	static_cast<QIntValidator*>(ptr)->setTop(vin);
}

int QIntValidator_Top(void* ptr)
{
	return static_cast<QIntValidator*>(ptr)->top();
}

void QIntValidator_DestroyQIntValidator(void* ptr)
{
	static_cast<QIntValidator*>(ptr)->~QIntValidator();
}

void QIntValidator_DestroyQIntValidatorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QIntValidator_Validate(void* ptr, struct QtGui_PackedString input, int pos)
{
	return static_cast<QIntValidator*>(ptr)->validate(*(new QString(QString::fromUtf8(input.data, input.len))), pos);
}

long long QIntValidator_ValidateDefault(void* ptr, struct QtGui_PackedString input, int pos)
{
		return static_cast<QIntValidator*>(ptr)->QIntValidator::validate(*(new QString(QString::fromUtf8(input.data, input.len))), pos);
}

class MyQKeyEvent: public QKeyEvent
{
public:
	MyQKeyEvent(QEvent::Type ty, int key, Qt::KeyboardModifiers modifiers, const QString &text = QString(), bool autorep = false, ushort count = 1) : QKeyEvent(ty, key, modifiers, text, autorep, count) {QKeyEvent_QKeyEvent_QRegisterMetaType();};
	MyQKeyEvent(QEvent::Type ty, int key, Qt::KeyboardModifiers modifiers, quint32 nativeScanCode, quint32 nativeVirtualKey, quint32 nativeModifiers, const QString &text = QString(), bool autorep = false, ushort count = 1) : QKeyEvent(ty, key, modifiers, nativeScanCode, nativeVirtualKey, nativeModifiers, text, autorep, count) {QKeyEvent_QKeyEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QKeyEvent*)
Q_DECLARE_METATYPE(MyQKeyEvent*)

int QKeyEvent_QKeyEvent_QRegisterMetaType(){qRegisterMetaType<QKeyEvent*>(); return qRegisterMetaType<MyQKeyEvent*>();}

void* QKeyEvent_NewQKeyEvent(long long ty, int key, long long modifiers, struct QtGui_PackedString text, char autorep, unsigned short count)
{
	return new MyQKeyEvent(static_cast<QEvent::Type>(ty), key, static_cast<Qt::KeyboardModifier>(modifiers), QString::fromUtf8(text.data, text.len), autorep != 0, count);
}

void* QKeyEvent_NewQKeyEvent2(long long ty, int key, long long modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, struct QtGui_PackedString text, char autorep, unsigned short count)
{
	return new MyQKeyEvent(static_cast<QEvent::Type>(ty), key, static_cast<Qt::KeyboardModifier>(modifiers), nativeScanCode, nativeVirtualKey, nativeModifiers, QString::fromUtf8(text.data, text.len), autorep != 0, count);
}

int QKeyEvent_Count(void* ptr)
{
	return static_cast<QKeyEvent*>(ptr)->count();
}

int QKeyEvent_Key(void* ptr)
{
	return static_cast<QKeyEvent*>(ptr)->key();
}

struct QtGui_PackedString QKeyEvent_Text(void* ptr)
{
	return ({ QByteArray* tac962f = new QByteArray(static_cast<QKeyEvent*>(ptr)->text().toUtf8()); QtGui_PackedString { const_cast<char*>(tac962f->prepend("WHITESPACE").constData()+10), tac962f->size()-10, tac962f }; });
}

Q_DECLARE_METATYPE(QKeySequence)
Q_DECLARE_METATYPE(QKeySequence*)
void* QKeySequence_NewQKeySequence()
{
	return new QKeySequence();
}

void* QKeySequence_NewQKeySequence2(struct QtGui_PackedString key, long long format)
{
	return new QKeySequence(QString::fromUtf8(key.data, key.len), static_cast<QKeySequence::SequenceFormat>(format));
}

void* QKeySequence_NewQKeySequence3(int k1, int k2, int k3, int k4)
{
	return new QKeySequence(k1, k2, k3, k4);
}

void* QKeySequence_NewQKeySequence4(void* keysequence)
{
	return new QKeySequence(*static_cast<QKeySequence*>(keysequence));
}

void* QKeySequence_NewQKeySequence5(long long key)
{
	return new QKeySequence(static_cast<QKeySequence::StandardKey>(key));
}

int QKeySequence_Count(void* ptr)
{
	return static_cast<QKeySequence*>(ptr)->count();
}

void* QKeySequence_QKeySequence_FromString(struct QtGui_PackedString str, long long format)
{
	return new QKeySequence(QKeySequence::fromString(QString::fromUtf8(str.data, str.len), static_cast<QKeySequence::SequenceFormat>(format)));
}

char QKeySequence_IsEmpty(void* ptr)
{
	return static_cast<QKeySequence*>(ptr)->isEmpty();
}

void QKeySequence_DestroyQKeySequence(void* ptr)
{
	static_cast<QKeySequence*>(ptr)->~QKeySequence();
}

void* QKeySequence___keyBindings_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QKeySequence___keyBindings_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QKeySequence___keyBindings_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QKeySequence___listFromString_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QKeySequence___listFromString_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QKeySequence___listFromString_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QKeySequence___listToString_list_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QKeySequence___listToString_list_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QKeySequence___listToString_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

Q_DECLARE_METATYPE(QMatrix4x4)
Q_DECLARE_METATYPE(QMatrix4x4*)
void* QMatrix4x4_NewQMatrix4x4()
{
	return new QMatrix4x4();
}

void* QMatrix4x4_NewQMatrix4x43(float values)
{
	return new QMatrix4x4(const_cast<const float*>(&values));
}

void* QMatrix4x4_NewQMatrix4x44(float m11, float m12, float m13, float m14, float m21, float m22, float m23, float m24, float m31, float m32, float m33, float m34, float m41, float m42, float m43, float m44)
{
	return new QMatrix4x4(m11, m12, m13, m14, m21, m22, m23, m24, m31, m32, m33, m34, m41, m42, m43, m44);
}

void* QMatrix4x4_Column(void* ptr, int index)
{
	return new QVector4D(static_cast<QMatrix4x4*>(ptr)->column(index));
}

float QMatrix4x4_Data(void* ptr)
{
	return *static_cast<QMatrix4x4*>(ptr)->data();
}

float QMatrix4x4_Data2(void* ptr)
{
	return *static_cast<QMatrix4x4*>(ptr)->data();
}

void* QMatrix4x4_Map(void* ptr, void* point)
{
	return ({ QPoint tmpValue = static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QPoint*>(point)); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QMatrix4x4_Map2(void* ptr, void* point)
{
	return ({ QPointF tmpValue = static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QPointF*>(point)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QMatrix4x4_Map3(void* ptr, void* point)
{
	return new QVector3D(static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QVector3D*>(point)));
}

void* QMatrix4x4_Map4(void* ptr, void* point)
{
	return new QVector4D(static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QVector4D*>(point)));
}

void* QMatrix4x4_Row(void* ptr, int index)
{
	return new QVector4D(static_cast<QMatrix4x4*>(ptr)->row(index));
}

void QMatrix4x4_Scale(void* ptr, void* vector)
{
	static_cast<QMatrix4x4*>(ptr)->scale(*static_cast<QVector3D*>(vector));
}

void QMatrix4x4_Scale2(void* ptr, float x, float y)
{
	static_cast<QMatrix4x4*>(ptr)->scale(x, y);
}

void QMatrix4x4_Scale3(void* ptr, float x, float y, float z)
{
	static_cast<QMatrix4x4*>(ptr)->scale(x, y, z);
}

void QMatrix4x4_Scale4(void* ptr, float factor)
{
	static_cast<QMatrix4x4*>(ptr)->scale(factor);
}

void QMatrix4x4_SetColumn(void* ptr, int index, void* value)
{
	static_cast<QMatrix4x4*>(ptr)->setColumn(index, *static_cast<QVector4D*>(value));
}

class MyQMouseEvent: public QMouseEvent
{
public:
	MyQMouseEvent(QEvent::Type ty, const QPointF &localPos, Qt::MouseButton button, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers) : QMouseEvent(ty, localPos, button, buttons, modifiers) {QMouseEvent_QMouseEvent_QRegisterMetaType();};
	MyQMouseEvent(QEvent::Type ty, const QPointF &localPos, const QPointF &screenPos, Qt::MouseButton button, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers) : QMouseEvent(ty, localPos, screenPos, button, buttons, modifiers) {QMouseEvent_QMouseEvent_QRegisterMetaType();};
	MyQMouseEvent(QEvent::Type ty, const QPointF &localPos, const QPointF &windowPos, const QPointF &screenPos, Qt::MouseButton button, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers) : QMouseEvent(ty, localPos, windowPos, screenPos, button, buttons, modifiers) {QMouseEvent_QMouseEvent_QRegisterMetaType();};
	MyQMouseEvent(QEvent::Type ty, const QPointF &localPos, const QPointF &windowPos, const QPointF &screenPos, Qt::MouseButton button, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers, Qt::MouseEventSource source) : QMouseEvent(ty, localPos, windowPos, screenPos, button, buttons, modifiers, source) {QMouseEvent_QMouseEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QMouseEvent*)
Q_DECLARE_METATYPE(MyQMouseEvent*)

int QMouseEvent_QMouseEvent_QRegisterMetaType(){qRegisterMetaType<QMouseEvent*>(); return qRegisterMetaType<MyQMouseEvent*>();}

void* QMouseEvent_NewQMouseEvent(long long ty, void* localPos, long long button, long long buttons, long long modifiers)
{
	return new MyQMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QMouseEvent_NewQMouseEvent2(long long ty, void* localPos, void* screenPos, long long button, long long buttons, long long modifiers)
{
	return new MyQMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QMouseEvent_NewQMouseEvent3(long long ty, void* localPos, void* windowPos, void* screenPos, long long button, long long buttons, long long modifiers)
{
	return new MyQMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(windowPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QMouseEvent_NewQMouseEvent4(long long ty, void* localPos, void* windowPos, void* screenPos, long long button, long long buttons, long long modifiers, long long source)
{
	return new MyQMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(windowPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::MouseEventSource>(source));
}

long long QMouseEvent_Button(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->button();
}

long long QMouseEvent_Buttons(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->buttons();
}

void* QMouseEvent_GlobalPos(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QMouseEvent*>(ptr)->globalPos(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QMouseEvent_Pos(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QMouseEvent*>(ptr)->pos(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

int QMouseEvent_X(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->x();
}

int QMouseEvent_Y(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->y();
}

Q_DECLARE_METATYPE(QPageSize*)
void* QPageSize_NewQPageSize()
{
	return new QPageSize();
}

void* QPageSize_NewQPageSize2(long long pageSize)
{
	return new QPageSize(static_cast<QPageSize::PageSizeId>(pageSize));
}

void* QPageSize_NewQPageSize3(void* pointSize, struct QtGui_PackedString name, long long matchPolicy)
{
	return new QPageSize(*static_cast<QSize*>(pointSize), QString::fromUtf8(name.data, name.len), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

void* QPageSize_NewQPageSize4(void* size, long long units, struct QtGui_PackedString name, long long matchPolicy)
{
	return new QPageSize(*static_cast<QSizeF*>(size), static_cast<QPageSize::Unit>(units), QString::fromUtf8(name.data, name.len), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

void* QPageSize_NewQPageSize5(void* other)
{
	return new QPageSize(*static_cast<QPageSize*>(other));
}

long long QPageSize_Id(void* ptr)
{
	return static_cast<QPageSize*>(ptr)->id();
}

long long QPageSize_QPageSize_Id2(void* pointSize, long long matchPolicy)
{
	return QPageSize::id(*static_cast<QSize*>(pointSize), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

long long QPageSize_QPageSize_Id3(void* size, long long units, long long matchPolicy)
{
	return QPageSize::id(*static_cast<QSizeF*>(size), static_cast<QPageSize::Unit>(units), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

long long QPageSize_QPageSize_Id4(int windowsId)
{
	return QPageSize::id(windowsId);
}

char QPageSize_IsValid(void* ptr)
{
	return static_cast<QPageSize*>(ptr)->isValid();
}

struct QtGui_PackedString QPageSize_Key(void* ptr)
{
	return ({ QByteArray* t01da2c = new QByteArray(static_cast<QPageSize*>(ptr)->key().toUtf8()); QtGui_PackedString { const_cast<char*>(t01da2c->prepend("WHITESPACE").constData()+10), t01da2c->size()-10, t01da2c }; });
}

struct QtGui_PackedString QPageSize_QPageSize_Key2(long long pageSizeId)
{
	return ({ QByteArray* tba9d4f = new QByteArray(QPageSize::key(static_cast<QPageSize::PageSizeId>(pageSizeId)).toUtf8()); QtGui_PackedString { const_cast<char*>(tba9d4f->prepend("WHITESPACE").constData()+10), tba9d4f->size()-10, tba9d4f }; });
}

struct QtGui_PackedString QPageSize_Name(void* ptr)
{
	return ({ QByteArray* t2620ab = new QByteArray(static_cast<QPageSize*>(ptr)->name().toUtf8()); QtGui_PackedString { const_cast<char*>(t2620ab->prepend("WHITESPACE").constData()+10), t2620ab->size()-10, t2620ab }; });
}

struct QtGui_PackedString QPageSize_QPageSize_Name2(long long pageSizeId)
{
	return ({ QByteArray* t8d3c9f = new QByteArray(QPageSize::name(static_cast<QPageSize::PageSizeId>(pageSizeId)).toUtf8()); QtGui_PackedString { const_cast<char*>(t8d3c9f->prepend("WHITESPACE").constData()+10), t8d3c9f->size()-10, t8d3c9f }; });
}

void* QPageSize_Rect(void* ptr, long long units)
{
	return ({ QRectF tmpValue = static_cast<QPageSize*>(ptr)->rect(static_cast<QPageSize::Unit>(units)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QPageSize_Size(void* ptr, long long units)
{
	return ({ QSizeF tmpValue = static_cast<QPageSize*>(ptr)->size(static_cast<QPageSize::Unit>(units)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* QPageSize_QPageSize_Size2(long long pageSizeId, long long units)
{
	return ({ QSizeF tmpValue = QPageSize::size(static_cast<QPageSize::PageSizeId>(pageSizeId), static_cast<QPageSize::Unit>(units)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void QPageSize_DestroyQPageSize(void* ptr)
{
	static_cast<QPageSize*>(ptr)->~QPageSize();
}

class MyQPagedPaintDevice: public QPagedPaintDevice
{
public:
	bool newPage() { return callbackQPagedPaintDevice_NewPage(this) != 0; };
	 ~MyQPagedPaintDevice() { callbackQPagedPaintDevice_DestroyQPagedPaintDevice(this); };
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPaintDevice_Metric(const_cast<void*>(static_cast<const void*>(this)), metric); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPagedPaintDevice_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QPagedPaintDevice*)
Q_DECLARE_METATYPE(MyQPagedPaintDevice*)

int QPagedPaintDevice_QPagedPaintDevice_QRegisterMetaType(){qRegisterMetaType<QPagedPaintDevice*>(); return qRegisterMetaType<MyQPagedPaintDevice*>();}

char QPagedPaintDevice_NewPage(void* ptr)
{
	return static_cast<QPagedPaintDevice*>(ptr)->newPage();
}

void QPagedPaintDevice_DestroyQPagedPaintDevice(void* ptr)
{
	static_cast<QPagedPaintDevice*>(ptr)->~QPagedPaintDevice();
}

void QPagedPaintDevice_DestroyQPagedPaintDeviceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QPagedPaintDevice_PaintEngine(void* ptr)
{
	return static_cast<QPagedPaintDevice*>(ptr)->paintEngine();
}

void* QPagedPaintDevice_PaintEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);
	
}

class MyQPaintDevice: public QPaintDevice
{
public:
	MyQPaintDevice() : QPaintDevice() {QPaintDevice_QPaintDevice_QRegisterMetaType();};
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPaintDevice_Metric(const_cast<void*>(static_cast<const void*>(this)), metric); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPaintDevice_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQPaintDevice() { callbackQPaintDevice_DestroyQPaintDevice(this); };
};

Q_DECLARE_METATYPE(QPaintDevice*)
Q_DECLARE_METATYPE(MyQPaintDevice*)

int QPaintDevice_QPaintDevice_QRegisterMetaType(){qRegisterMetaType<QPaintDevice*>(); return qRegisterMetaType<MyQPaintDevice*>();}

void* QPaintDevice_NewQPaintDevice()
{
	return new MyQPaintDevice();
}

int QPaintDevice_Height(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->height();
}

int QPaintDevice_Metric(void* ptr, long long metric)
{
	return static_cast<QPaintDevice*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QPaintDevice_MetricDefault(void* ptr, long long metric)
{
	if (dynamic_cast<QBitmap*>(static_cast<QPaintDevice*>(ptr))) {
		return static_cast<QBitmap*>(ptr)->QBitmap::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
	} else if (dynamic_cast<QPixmap*>(static_cast<QPaintDevice*>(ptr))) {
		return static_cast<QPixmap*>(ptr)->QPixmap::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
	} else if (dynamic_cast<QPagedPaintDevice*>(static_cast<QPaintDevice*>(ptr))) {
		return static_cast<QPagedPaintDevice*>(ptr)->QPagedPaintDevice::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
	} else if (dynamic_cast<QImage*>(static_cast<QPaintDevice*>(ptr))) {
		return static_cast<QImage*>(ptr)->QImage::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
	} else {
		return static_cast<QPaintDevice*>(ptr)->QPaintDevice::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
	}
}

void* QPaintDevice_PaintEngine(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->paintEngine();
}

int QPaintDevice_Width(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->width();
}

void QPaintDevice_DestroyQPaintDevice(void* ptr)
{
	static_cast<QPaintDevice*>(ptr)->~QPaintDevice();
}

void QPaintDevice_DestroyQPaintDeviceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQPaintEngine: public QPaintEngine
{
public:
	MyQPaintEngine(QPaintEngine::PaintEngineFeatures caps = PaintEngineFeatures()) : QPaintEngine(caps) {QPaintEngine_QPaintEngine_QRegisterMetaType();};
	bool begin(QPaintDevice * pdev) { return callbackQPaintEngine_Begin(this, pdev) != 0; };
	void drawPixmap(const QRectF & r, const QPixmap & pm, const QRectF & sr) { callbackQPaintEngine_DrawPixmap(this, const_cast<QRectF*>(&r), const_cast<QPixmap*>(&pm), const_cast<QRectF*>(&sr)); };
	bool end() { return callbackQPaintEngine_End(this) != 0; };
	QPaintEngine::Type type() const { return static_cast<QPaintEngine::Type>(callbackQPaintEngine_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	void updateState(const QPaintEngineState & state) { callbackQPaintEngine_UpdateState(this, const_cast<QPaintEngineState*>(&state)); };
	 ~MyQPaintEngine() { callbackQPaintEngine_DestroyQPaintEngine(this); };
};

Q_DECLARE_METATYPE(QPaintEngine*)
Q_DECLARE_METATYPE(MyQPaintEngine*)

int QPaintEngine_QPaintEngine_QRegisterMetaType(){qRegisterMetaType<QPaintEngine*>(); return qRegisterMetaType<MyQPaintEngine*>();}

void* QPaintEngine_NewQPaintEngine(long long caps)
{
	return new MyQPaintEngine(static_cast<QPaintEngine::PaintEngineFeature>(caps));
}

char QPaintEngine_Begin(void* ptr, void* pdev)
{
	return static_cast<QPaintEngine*>(ptr)->begin(static_cast<QPaintDevice*>(pdev));
}

void QPaintEngine_DrawPixmap(void* ptr, void* r, void* pm, void* sr)
{
	static_cast<QPaintEngine*>(ptr)->drawPixmap(*static_cast<QRectF*>(r), *static_cast<QPixmap*>(pm), *static_cast<QRectF*>(sr));
}

char QPaintEngine_End(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->end();
}

void QPaintEngine_SetActive(void* ptr, char state)
{
	static_cast<QPaintEngine*>(ptr)->setActive(state != 0);
}

long long QPaintEngine_Type(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->type();
}

void QPaintEngine_UpdateState(void* ptr, void* state)
{
	static_cast<QPaintEngine*>(ptr)->updateState(*static_cast<QPaintEngineState*>(state));
}

void QPaintEngine_DestroyQPaintEngine(void* ptr)
{
	static_cast<QPaintEngine*>(ptr)->~QPaintEngine();
}

void QPaintEngine_DestroyQPaintEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QPaintEngineState_Brush(void* ptr)
{
	return new QBrush(static_cast<QPaintEngineState*>(ptr)->brush());
}

void* QPaintEngineState_Font(void* ptr)
{
	return new QFont(static_cast<QPaintEngineState*>(ptr)->font());
}

void* QPaintEngineState_Transform(void* ptr)
{
	return new QTransform(static_cast<QPaintEngineState*>(ptr)->transform());
}

class MyQPaintEvent: public QPaintEvent
{
public:
	MyQPaintEvent(const QRegion &paintRegion) : QPaintEvent(paintRegion) {QPaintEvent_QPaintEvent_QRegisterMetaType();};
	MyQPaintEvent(const QRect &paintRect) : QPaintEvent(paintRect) {QPaintEvent_QPaintEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QPaintEvent*)
Q_DECLARE_METATYPE(MyQPaintEvent*)

int QPaintEvent_QPaintEvent_QRegisterMetaType(){qRegisterMetaType<QPaintEvent*>(); return qRegisterMetaType<MyQPaintEvent*>();}

void* QPaintEvent_NewQPaintEvent(void* paintRegion)
{
	return new MyQPaintEvent(*static_cast<QRegion*>(paintRegion));
}

void* QPaintEvent_NewQPaintEvent2(void* paintRect)
{
	return new MyQPaintEvent(*static_cast<QRect*>(paintRect));
}

void* QPaintEvent_Rect(void* ptr)
{
	return const_cast<QRect*>(&static_cast<QPaintEvent*>(ptr)->rect());
}

Q_DECLARE_METATYPE(QPainter*)
void* QPainter_NewQPainter()
{
	return new QPainter();
}

void* QPainter_NewQPainter2(void* device)
{
	if (dynamic_cast<QWidget*>(static_cast<QObject*>(device))) {
		return new QPainter(static_cast<QWidget*>(device));
	} else {
		return new QPainter(static_cast<QPaintDevice*>(device));
	}
}

void* QPainter_Background(void* ptr)
{
	return const_cast<QBrush*>(&static_cast<QPainter*>(ptr)->background());
}

char QPainter_Begin(void* ptr, void* device)
{
	return static_cast<QPainter*>(ptr)->begin(static_cast<QPaintDevice*>(device));
}

void* QPainter_BoundingRect(void* ptr, void* rectangle, int flags, struct QtGui_PackedString text)
{
	return ({ QRectF tmpValue = static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRectF*>(rectangle), flags, QString::fromUtf8(text.data, text.len)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QPainter_BoundingRect2(void* ptr, void* rectangle, int flags, struct QtGui_PackedString text)
{
	return ({ QRect tmpValue = static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRect*>(rectangle), flags, QString::fromUtf8(text.data, text.len)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QPainter_BoundingRect3(void* ptr, int x, int y, int w, int h, int flags, struct QtGui_PackedString text)
{
	return ({ QRect tmpValue = static_cast<QPainter*>(ptr)->boundingRect(x, y, w, h, flags, QString::fromUtf8(text.data, text.len)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QPainter_BoundingRect4(void* ptr, void* rectangle, struct QtGui_PackedString text, void* option)
{
	return ({ QRectF tmpValue = static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRectF*>(rectangle), QString::fromUtf8(text.data, text.len), *static_cast<QTextOption*>(option)); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QPainter_Brush(void* ptr)
{
	return const_cast<QBrush*>(&static_cast<QPainter*>(ptr)->brush());
}

void* QPainter_Device(void* ptr)
{
	return static_cast<QPainter*>(ptr)->device();
}

char QPainter_End(void* ptr)
{
	return static_cast<QPainter*>(ptr)->end();
}

void* QPainter_Font(void* ptr)
{
	return const_cast<QFont*>(&static_cast<QPainter*>(ptr)->font());
}

void* QPainter_FontMetrics(void* ptr)
{
	return new QFontMetrics(static_cast<QPainter*>(ptr)->fontMetrics());
}

void QPainter_Save(void* ptr)
{
	static_cast<QPainter*>(ptr)->save();
}

void QPainter_Scale(void* ptr, double sx, double sy)
{
	static_cast<QPainter*>(ptr)->scale(sx, sy);
}

void QPainter_SetBackground(void* ptr, void* brush)
{
	static_cast<QPainter*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QPainter_SetFont(void* ptr, void* font)
{
	static_cast<QPainter*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QPainter_SetWindow(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->setWindow(*static_cast<QRect*>(rectangle));
}

void QPainter_SetWindow2(void* ptr, int x, int y, int width, int height)
{
	static_cast<QPainter*>(ptr)->setWindow(x, y, width, height);
}

void* QPainter_Transform(void* ptr)
{
	return const_cast<QTransform*>(&static_cast<QPainter*>(ptr)->transform());
}

void* QPainter_Window(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QPainter*>(ptr)->window(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPainter_DestroyQPainter(void* ptr)
{
	static_cast<QPainter*>(ptr)->~QPainter();
}

void* QPainter___drawLines_lines_atList2(void* ptr, int i)
{
	return ({ QLineF tmpValue = ({QLineF tmp = static_cast<QVector<QLineF>*>(ptr)->at(i); if (i == static_cast<QVector<QLineF>*>(ptr)->size()-1) { static_cast<QVector<QLineF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QLineF(tmpValue.p1(), tmpValue.p2()); });
}

void QPainter___drawLines_lines_setList2(void* ptr, void* i)
{
	static_cast<QVector<QLineF>*>(ptr)->append(*static_cast<QLineF*>(i));
}

void* QPainter___drawLines_lines_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QLineF>();
}

void* QPainter___drawLines_pointPairs_atList4(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QPainter___drawLines_pointPairs_setList4(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QPainter___drawLines_pointPairs_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

void* QPainter___drawLines_lines_atList6(void* ptr, int i)
{
	return ({ QLine tmpValue = ({QLine tmp = static_cast<QVector<QLine>*>(ptr)->at(i); if (i == static_cast<QVector<QLine>*>(ptr)->size()-1) { static_cast<QVector<QLine>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QLine(tmpValue.p1(), tmpValue.p2()); });
}

void QPainter___drawLines_lines_setList6(void* ptr, void* i)
{
	static_cast<QVector<QLine>*>(ptr)->append(*static_cast<QLine*>(i));
}

void* QPainter___drawLines_lines_newList6(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QLine>();
}

void* QPainter___drawLines_pointPairs_atList8(void* ptr, int i)
{
	return ({ QPoint tmpValue = ({QPoint tmp = static_cast<QVector<QPoint>*>(ptr)->at(i); if (i == static_cast<QVector<QPoint>*>(ptr)->size()-1) { static_cast<QVector<QPoint>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPainter___drawLines_pointPairs_setList8(void* ptr, void* i)
{
	static_cast<QVector<QPoint>*>(ptr)->append(*static_cast<QPoint*>(i));
}

void* QPainter___drawLines_pointPairs_newList8(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPoint>();
}

void* QPainter___drawRects_rectangles_atList2(void* ptr, int i)
{
	return ({ QRectF tmpValue = ({QRectF tmp = static_cast<QVector<QRectF>*>(ptr)->at(i); if (i == static_cast<QVector<QRectF>*>(ptr)->size()-1) { static_cast<QVector<QRectF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPainter___drawRects_rectangles_setList2(void* ptr, void* i)
{
	static_cast<QVector<QRectF>*>(ptr)->append(*static_cast<QRectF*>(i));
}

void* QPainter___drawRects_rectangles_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRectF>();
}

void* QPainter___drawRects_rectangles_atList4(void* ptr, int i)
{
	return ({ QRect tmpValue = ({QRect tmp = static_cast<QVector<QRect>*>(ptr)->at(i); if (i == static_cast<QVector<QRect>*>(ptr)->size()-1) { static_cast<QVector<QRect>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPainter___drawRects_rectangles_setList4(void* ptr, void* i)
{
	static_cast<QVector<QRect>*>(ptr)->append(*static_cast<QRect*>(i));
}

void* QPainter___drawRects_rectangles_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRect>();
}

Q_DECLARE_METATYPE(QPainterPath)
Q_DECLARE_METATYPE(QPainterPath*)
void* QPainterPath_NewQPainterPath()
{
	return new QPainterPath();
}

void* QPainterPath_NewQPainterPath2(void* startPoint)
{
	return new QPainterPath(*static_cast<QPointF*>(startPoint));
}

void* QPainterPath_NewQPainterPath3(void* path)
{
	return new QPainterPath(*static_cast<QPainterPath*>(path));
}

void* QPainterPath_BoundingRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QPainterPath*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPainterPath_Clear(void* ptr)
{
	static_cast<QPainterPath*>(ptr)->clear();
}

char QPainterPath_IsEmpty(void* ptr)
{
	return static_cast<QPainterPath*>(ptr)->isEmpty();
}

double QPainterPath_Length(void* ptr)
{
	return static_cast<QPainterPath*>(ptr)->length();
}

void QPainterPath_DestroyQPainterPath(void* ptr)
{
	static_cast<QPainterPath*>(ptr)->~QPainterPath();
}

void* QPainterPath___toFillPolygons_atList(void* ptr, int i)
{
	return new QPolygonF(({QPolygonF tmp = static_cast<QList<QPolygonF>*>(ptr)->at(i); if (i == static_cast<QList<QPolygonF>*>(ptr)->size()-1) { static_cast<QList<QPolygonF>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPainterPath___toFillPolygons_setList(void* ptr, void* i)
{
	static_cast<QList<QPolygonF>*>(ptr)->append(*static_cast<QPolygonF*>(i));
}

void* QPainterPath___toFillPolygons_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPolygonF>();
}

void* QPainterPath___toFillPolygons_atList2(void* ptr, int i)
{
	return new QPolygonF(({QPolygonF tmp = static_cast<QList<QPolygonF>*>(ptr)->at(i); if (i == static_cast<QList<QPolygonF>*>(ptr)->size()-1) { static_cast<QList<QPolygonF>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPainterPath___toFillPolygons_setList2(void* ptr, void* i)
{
	static_cast<QList<QPolygonF>*>(ptr)->append(*static_cast<QPolygonF*>(i));
}

void* QPainterPath___toFillPolygons_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPolygonF>();
}

void* QPainterPath___toSubpathPolygons_atList(void* ptr, int i)
{
	return new QPolygonF(({QPolygonF tmp = static_cast<QList<QPolygonF>*>(ptr)->at(i); if (i == static_cast<QList<QPolygonF>*>(ptr)->size()-1) { static_cast<QList<QPolygonF>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPainterPath___toSubpathPolygons_setList(void* ptr, void* i)
{
	static_cast<QList<QPolygonF>*>(ptr)->append(*static_cast<QPolygonF*>(i));
}

void* QPainterPath___toSubpathPolygons_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPolygonF>();
}

void* QPainterPath___toSubpathPolygons_atList2(void* ptr, int i)
{
	return new QPolygonF(({QPolygonF tmp = static_cast<QList<QPolygonF>*>(ptr)->at(i); if (i == static_cast<QList<QPolygonF>*>(ptr)->size()-1) { static_cast<QList<QPolygonF>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPainterPath___toSubpathPolygons_setList2(void* ptr, void* i)
{
	static_cast<QList<QPolygonF>*>(ptr)->append(*static_cast<QPolygonF*>(i));
}

void* QPainterPath___toSubpathPolygons_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPolygonF>();
}

int QPalette_NColorRoles_Type()
{
	return QPalette::NColorRoles;
}

class MyQPixmap: public QPixmap
{
public:
	MyQPixmap() : QPixmap() {QPixmap_QPixmap_QRegisterMetaType();};
	MyQPixmap(int width, int height) : QPixmap(width, height) {QPixmap_QPixmap_QRegisterMetaType();};
	MyQPixmap(const QSize &size) : QPixmap(size) {QPixmap_QPixmap_QRegisterMetaType();};
	MyQPixmap(const QString &fileName, const char *format = Q_NULLPTR, Qt::ImageConversionFlags flags = Qt::AutoColor) : QPixmap(fileName, format, flags) {QPixmap_QPixmap_QRegisterMetaType();};
	MyQPixmap(const QPixmap &pixmap) : QPixmap(pixmap) {QPixmap_QPixmap_QRegisterMetaType();};
	 ~MyQPixmap() { callbackQPixmap_DestroyQPixmap(this); };
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPaintDevice_Metric(const_cast<void*>(static_cast<const void*>(this)), metric); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPixmap_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QPixmap*)
Q_DECLARE_METATYPE(MyQPixmap*)

int QPixmap_QPixmap_QRegisterMetaType(){qRegisterMetaType<QPixmap*>(); return qRegisterMetaType<MyQPixmap*>();}

void* QPixmap_NewQPixmap()
{
	return new MyQPixmap();
}

void* QPixmap_NewQPixmap2(void* size)
{
	return new MyQPixmap(*static_cast<QSize*>(size));
}

void* QPixmap_NewQPixmap3(struct QtGui_PackedString fileName, char* format, long long flags)
{
	return new MyQPixmap(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

void* QPixmap_NewQPixmap5(void* pixmap)
{
	return new MyQPixmap(*static_cast<QPixmap*>(pixmap));
}

void* QPixmap_Copy(void* ptr, void* rectangle)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->copy(*static_cast<QRect*>(rectangle)));
}

void* QPixmap_Copy2(void* ptr, int x, int y, int width, int height)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->copy(x, y, width, height));
}

char QPixmap_Load(void* ptr, struct QtGui_PackedString fileName, char* format, long long flags)
{
	return static_cast<QPixmap*>(ptr)->load(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

char QPixmap_LoadFromData(void* ptr, char* data, unsigned int l, char* format, long long flags)
{
	return static_cast<QPixmap*>(ptr)->loadFromData(const_cast<const uchar*>(static_cast<uchar*>(static_cast<void*>(data))), l, const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

char QPixmap_LoadFromData2(void* ptr, void* data, char* format, long long flags)
{
	return static_cast<QPixmap*>(ptr)->loadFromData(*static_cast<QByteArray*>(data), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

void* QPixmap_Rect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QPixmap*>(ptr)->rect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

char QPixmap_Save(void* ptr, struct QtGui_PackedString fileName, char* format, int quality)
{
	return static_cast<QPixmap*>(ptr)->save(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format), quality);
}

char QPixmap_Save2(void* ptr, void* device, char* format, int quality)
{
	return static_cast<QPixmap*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format), quality);
}

void* QPixmap_Scaled(void* ptr, void* size, long long aspectRatioMode, long long transformMode)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->scaled(*static_cast<QSize*>(size), static_cast<Qt::AspectRatioMode>(aspectRatioMode), static_cast<Qt::TransformationMode>(transformMode)));
}

void* QPixmap_Scaled2(void* ptr, int width, int height, long long aspectRatioMode, long long transformMode)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->scaled(width, height, static_cast<Qt::AspectRatioMode>(aspectRatioMode), static_cast<Qt::TransformationMode>(transformMode)));
}

void* QPixmap_ScaledToHeight(void* ptr, int height, long long mode)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->scaledToHeight(height, static_cast<Qt::TransformationMode>(mode)));
}

void QPixmap_Scroll(void* ptr, int dx, int dy, int x, int y, int width, int height, void* exposed)
{
	static_cast<QPixmap*>(ptr)->scroll(dx, dy, x, y, width, height, static_cast<QRegion*>(exposed));
}

void QPixmap_Scroll2(void* ptr, int dx, int dy, void* rect, void* exposed)
{
	static_cast<QPixmap*>(ptr)->scroll(dx, dy, *static_cast<QRect*>(rect), static_cast<QRegion*>(exposed));
}

void* QPixmap_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPixmap*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QPixmap_ToImage(void* ptr)
{
	return new QImage(static_cast<QPixmap*>(ptr)->toImage());
}

void QPixmap_DestroyQPixmap(void* ptr)
{
	static_cast<QPixmap*>(ptr)->~QPixmap();
}

void QPixmap_DestroyQPixmapDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QPixmap_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QPixmap*>(ptr));
}

void* QPixmap_PaintEngine(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->paintEngine();
}

void* QPixmap_PaintEngineDefault(void* ptr)
{
	if (dynamic_cast<QBitmap*>(static_cast<QPixmap*>(ptr))) {
		return static_cast<QBitmap*>(ptr)->QBitmap::paintEngine();
	} else {
		return static_cast<QPixmap*>(ptr)->QPixmap::paintEngine();
	}
}

Q_DECLARE_METATYPE(QPolygon)
Q_DECLARE_METATYPE(QPolygon*)
void* QPolygon_NewQPolygon()
{
	return new QPolygon();
}

void* QPolygon_NewQPolygon2(int size)
{
	return new QPolygon(size);
}

void* QPolygon_NewQPolygon3(void* points)
{
	return new QPolygon(*static_cast<QVector<QPoint>*>(points));
}

void* QPolygon_NewQPolygon5(void* rectangle, char closed)
{
	return new QPolygon(*static_cast<QRect*>(rectangle), closed != 0);
}

void* QPolygon_BoundingRect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QPolygon*>(ptr)->boundingRect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPolygon_Point(void* ptr, int index, int x, int y)
{
	static_cast<QPolygon*>(ptr)->point(index, &x, &y);
}

void* QPolygon_Point2(void* ptr, int index)
{
	return ({ QPoint tmpValue = static_cast<QPolygon*>(ptr)->point(index); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPolygon_SetPoint(void* ptr, int index, int x, int y)
{
	static_cast<QPolygon*>(ptr)->setPoint(index, x, y);
}

void QPolygon_SetPoint2(void* ptr, int index, void* point)
{
	static_cast<QPolygon*>(ptr)->setPoint(index, *static_cast<QPoint*>(point));
}

void QPolygon_DestroyQPolygon(void* ptr)
{
	static_cast<QPolygon*>(ptr)->~QPolygon();
}

void* QPolygon___QPolygon_points_atList3(void* ptr, int i)
{
	return ({ QPoint tmpValue = ({QPoint tmp = static_cast<QVector<QPoint>*>(ptr)->at(i); if (i == static_cast<QVector<QPoint>*>(ptr)->size()-1) { static_cast<QVector<QPoint>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPolygon___QPolygon_points_setList3(void* ptr, void* i)
{
	static_cast<QVector<QPoint>*>(ptr)->append(*static_cast<QPoint*>(i));
}

void* QPolygon___QPolygon_points_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPoint>();
}

void* QPolygon___QPolygon_v_atList4(void* ptr, int i)
{
	return ({ QPoint tmpValue = ({QPoint tmp = static_cast<QVector<QPoint>*>(ptr)->at(i); if (i == static_cast<QVector<QPoint>*>(ptr)->size()-1) { static_cast<QVector<QPoint>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPolygon___QPolygon_v_setList4(void* ptr, void* i)
{
	static_cast<QVector<QPoint>*>(ptr)->append(*static_cast<QPoint*>(i));
}

void* QPolygon___QPolygon_v_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPoint>();
}

void* QPolygon___QVector_other_atList4(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___QVector_other_setList4(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___QVector_other_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___QVector_other_atList5(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___QVector_other_setList5(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___QVector_other_newList5(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___append_value_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___append_value_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___append_value_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fill_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fill_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___fill_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fromList_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fromList_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___fromList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fromList_list_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fromList_list_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___fromList_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fromStdVector_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fromStdVector_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___fromStdVector_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___mid_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___mid_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___mid_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___swap_other_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___swap_other_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___swap_other_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___toList_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___toList_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___toList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

Q_DECLARE_METATYPE(QPolygonF)
Q_DECLARE_METATYPE(QPolygonF*)
void* QPolygonF_NewQPolygonF()
{
	return new QPolygonF();
}

void* QPolygonF_NewQPolygonF2(int size)
{
	return new QPolygonF(size);
}

void* QPolygonF_NewQPolygonF3(void* points)
{
	return new QPolygonF(*static_cast<QVector<QPointF>*>(points));
}

void* QPolygonF_NewQPolygonF5(void* rectangle)
{
	return new QPolygonF(*static_cast<QRectF*>(rectangle));
}

void* QPolygonF_NewQPolygonF6(void* polygon)
{
	return new QPolygonF(*static_cast<QPolygon*>(polygon));
}

void* QPolygonF_NewQPolygonF7(void* polygon)
{
	return new QPolygonF(*static_cast<QPolygonF*>(polygon));
}

void* QPolygonF_BoundingRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QPolygonF*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPolygonF_DestroyQPolygonF(void* ptr)
{
	static_cast<QPolygonF*>(ptr)->~QPolygonF();
}

void* QPolygonF___QPolygonF_points_atList3(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QPolygonF___QPolygonF_points_setList3(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QPolygonF___QPolygonF_points_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

void* QPolygonF___QPolygonF_v_atList4(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QPolygonF___QPolygonF_v_setList4(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QPolygonF___QPolygonF_v_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

void* QPolygonF___QVector_other_atList4(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___QVector_other_setList4(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygonF___QVector_other_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___QVector_other_atList5(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___QVector_other_setList5(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygonF___QVector_other_newList5(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___append_value_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___append_value_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygonF___append_value_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___fill_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___fill_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygonF___fill_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___fromList_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___fromList_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygonF___fromList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___fromList_list_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___fromList_list_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygonF___fromList_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___fromStdVector_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___fromStdVector_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygonF___fromStdVector_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___mid_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___mid_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygonF___mid_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___swap_other_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___swap_other_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygonF___swap_other_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygonF___toList_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygonF___toList_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygonF___toList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

Q_DECLARE_METATYPE(QRegion)
Q_DECLARE_METATYPE(QRegion*)
void* QRegion_NewQRegion()
{
	return new QRegion();
}

void* QRegion_NewQRegion2(int x, int y, int w, int h, long long t)
{
	return new QRegion(x, y, w, h, static_cast<QRegion::RegionType>(t));
}

void* QRegion_NewQRegion3(void* r, long long t)
{
	return new QRegion(*static_cast<QRect*>(r), static_cast<QRegion::RegionType>(t));
}

void* QRegion_NewQRegion4(void* a, long long fillRule)
{
	return new QRegion(*static_cast<QPolygon*>(a), static_cast<Qt::FillRule>(fillRule));
}

void* QRegion_NewQRegion5(void* r)
{
	return new QRegion(*static_cast<QRegion*>(r));
}

void* QRegion_NewQRegion6(void* other)
{
	return new QRegion(*static_cast<QRegion*>(other));
}

void* QRegion_NewQRegion7(void* bm)
{
	return new QRegion(*static_cast<QBitmap*>(bm));
}

void* QRegion_BoundingRect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QRegion*>(ptr)->boundingRect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

char QRegion_IsEmpty(void* ptr)
{
	return static_cast<QRegion*>(ptr)->isEmpty();
}

void* QRegion___rects_atList(void* ptr, int i)
{
	return ({ QRect tmpValue = ({QRect tmp = static_cast<QVector<QRect>*>(ptr)->at(i); if (i == static_cast<QVector<QRect>*>(ptr)->size()-1) { static_cast<QVector<QRect>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QRegion___rects_setList(void* ptr, void* i)
{
	static_cast<QVector<QRect>*>(ptr)->append(*static_cast<QRect*>(i));
}

void* QRegion___rects_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRect>();
}

class MyQResizeEvent: public QResizeEvent
{
public:
	MyQResizeEvent(const QSize &size, const QSize &oldSize) : QResizeEvent(size, oldSize) {QResizeEvent_QResizeEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QResizeEvent*)
Q_DECLARE_METATYPE(MyQResizeEvent*)

int QResizeEvent_QResizeEvent_QRegisterMetaType(){qRegisterMetaType<QResizeEvent*>(); return qRegisterMetaType<MyQResizeEvent*>();}

void* QResizeEvent_NewQResizeEvent(void* size, void* oldSize)
{
	return new MyQResizeEvent(*static_cast<QSize*>(size), *static_cast<QSize*>(oldSize));
}

void* QResizeEvent_Size(void* ptr)
{
	return const_cast<QSize*>(&static_cast<QResizeEvent*>(ptr)->size());
}

unsigned short QRgba64_Alpha(void* ptr)
{
	return static_cast<QRgba64*>(ptr)->alpha();
}

class MyQScreen: public QScreen
{
public:
	 ~MyQScreen() { callbackQScreen_DestroyQScreen(this); };
	void childEvent(QChildEvent * event) { callbackQScreen_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScreen_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScreen_CustomEvent(this, event); };
	void deleteLater() { callbackQScreen_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScreen_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScreen_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScreen_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScreen_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQScreen_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScreen_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QScreen*)
Q_DECLARE_METATYPE(MyQScreen*)

int QScreen_QScreen_QRegisterMetaType(){qRegisterMetaType<QScreen*>(); return qRegisterMetaType<MyQScreen*>();}

void* QScreen_Geometry(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QScreen*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

struct QtGui_PackedString QScreen_Model(void* ptr)
{
#ifndef Q_OS_WIN
	return ({ QByteArray* t131f94 = new QByteArray(static_cast<QScreen*>(ptr)->model().toUtf8()); QtGui_PackedString { const_cast<char*>(t131f94->prepend("WHITESPACE").constData()+10), t131f94->size()-10, t131f94 }; });
#endif
}

struct QtGui_PackedString QScreen_Name(void* ptr)
{
	return ({ QByteArray* tc60f02 = new QByteArray(static_cast<QScreen*>(ptr)->name().toUtf8()); QtGui_PackedString { const_cast<char*>(tc60f02->prepend("WHITESPACE").constData()+10), tc60f02->size()-10, tc60f02 }; });
}

long long QScreen_Orientation(void* ptr)
{
	return static_cast<QScreen*>(ptr)->orientation();
}

void* QScreen_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QScreen*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QScreen_DestroyQScreen(void* ptr)
{
	static_cast<QScreen*>(ptr)->~QScreen();
}

void QScreen_DestroyQScreenDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QScreen___virtualSiblings_atList(void* ptr, int i)
{
	return ({QScreen * tmp = static_cast<QList<QScreen *>*>(ptr)->at(i); if (i == static_cast<QList<QScreen *>*>(ptr)->size()-1) { static_cast<QList<QScreen *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___virtualSiblings_setList(void* ptr, void* i)
{
	static_cast<QList<QScreen *>*>(ptr)->append(static_cast<QScreen*>(i));
}

void* QScreen___virtualSiblings_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QScreen *>();
}

void* QScreen___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QScreen___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QScreen___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QScreen___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScreen___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QScreen___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QScreen___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QScreen___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QScreen___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QScreen_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QScreen*>(ptr)->QScreen::childEvent(static_cast<QChildEvent*>(event));
}

void QScreen_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScreen*>(ptr)->QScreen::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScreen_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QScreen*>(ptr)->QScreen::customEvent(static_cast<QEvent*>(event));
}

void QScreen_DeleteLaterDefault(void* ptr)
{
		static_cast<QScreen*>(ptr)->QScreen::deleteLater();
}

void QScreen_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScreen*>(ptr)->QScreen::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QScreen_EventDefault(void* ptr, void* e)
{
		return static_cast<QScreen*>(ptr)->QScreen::event(static_cast<QEvent*>(e));
}

char QScreen_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QScreen*>(ptr)->QScreen::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QScreen*>(ptr)->QScreen::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QScreen_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QScreen*>(ptr)->QScreen::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQShowEvent: public QShowEvent
{
public:
	MyQShowEvent() : QShowEvent() {QShowEvent_QShowEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QShowEvent*)
Q_DECLARE_METATYPE(MyQShowEvent*)

int QShowEvent_QShowEvent_QRegisterMetaType(){qRegisterMetaType<QShowEvent*>(); return qRegisterMetaType<MyQShowEvent*>();}

void* QShowEvent_NewQShowEvent()
{
	return new MyQShowEvent();
}

class MyQStandardItem: public QStandardItem
{
public:
	MyQStandardItem() : QStandardItem() {QStandardItem_QStandardItem_QRegisterMetaType();};
	MyQStandardItem(const QString &text) : QStandardItem(text) {QStandardItem_QStandardItem_QRegisterMetaType();};
	MyQStandardItem(const QIcon &icon, const QString &text) : QStandardItem(icon, text) {QStandardItem_QStandardItem_QRegisterMetaType();};
	MyQStandardItem(int rows, int columns = 1) : QStandardItem(rows, columns) {QStandardItem_QStandardItem_QRegisterMetaType();};
	MyQStandardItem(const QStandardItem &other) : QStandardItem(other) {QStandardItem_QStandardItem_QRegisterMetaType();};
	QVariant data(int role) const { return *static_cast<QVariant*>(callbackQStandardItem_Data(const_cast<void*>(static_cast<const void*>(this)), role)); };
	void read(QDataStream & in) { callbackQStandardItem_Read(this, static_cast<QDataStream*>(&in)); };
	void setData(const QVariant & value, int role) { callbackQStandardItem_SetData(this, const_cast<QVariant*>(&value), role); };
	int type() const { return callbackQStandardItem_Type(const_cast<void*>(static_cast<const void*>(this))); };
	void write(QDataStream & out) const { callbackQStandardItem_Write(const_cast<void*>(static_cast<const void*>(this)), static_cast<QDataStream*>(&out)); };
	 ~MyQStandardItem() { callbackQStandardItem_DestroyQStandardItem(this); };
};

Q_DECLARE_METATYPE(QStandardItem*)
Q_DECLARE_METATYPE(MyQStandardItem*)

int QStandardItem_QStandardItem_QRegisterMetaType(){qRegisterMetaType<QStandardItem*>(); return qRegisterMetaType<MyQStandardItem*>();}

void* QStandardItem_NewQStandardItem()
{
	return new MyQStandardItem();
}

void* QStandardItem_NewQStandardItem2(struct QtGui_PackedString text)
{
	return new MyQStandardItem(QString::fromUtf8(text.data, text.len));
}

void* QStandardItem_NewQStandardItem3(void* icon, struct QtGui_PackedString text)
{
	return new MyQStandardItem(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len));
}

void* QStandardItem_NewQStandardItem4(int rows, int columns)
{
	return new MyQStandardItem(rows, columns);
}

void* QStandardItem_NewQStandardItem5(void* other)
{
	return new MyQStandardItem(*static_cast<QStandardItem*>(other));
}

struct QtGui_PackedString QStandardItem_AccessibleDescription(void* ptr)
{
	return ({ QByteArray* t03d242 = new QByteArray(static_cast<QStandardItem*>(ptr)->accessibleDescription().toUtf8()); QtGui_PackedString { const_cast<char*>(t03d242->prepend("WHITESPACE").constData()+10), t03d242->size()-10, t03d242 }; });
}

struct QtGui_PackedString QStandardItem_AccessibleText(void* ptr)
{
	return ({ QByteArray* t3ec2e5 = new QByteArray(static_cast<QStandardItem*>(ptr)->accessibleText().toUtf8()); QtGui_PackedString { const_cast<char*>(t3ec2e5->prepend("WHITESPACE").constData()+10), t3ec2e5->size()-10, t3ec2e5 }; });
}

void QStandardItem_AppendColumn(void* ptr, void* items)
{
	static_cast<QStandardItem*>(ptr)->appendColumn(*static_cast<QList<QStandardItem *>*>(items));
}

void QStandardItem_AppendRow(void* ptr, void* items)
{
	static_cast<QStandardItem*>(ptr)->appendRow(*static_cast<QList<QStandardItem *>*>(items));
}

void QStandardItem_AppendRow2(void* ptr, void* item)
{
	static_cast<QStandardItem*>(ptr)->appendRow(static_cast<QStandardItem*>(item));
}

void QStandardItem_AppendRows(void* ptr, void* items)
{
	static_cast<QStandardItem*>(ptr)->appendRows(*static_cast<QList<QStandardItem *>*>(items));
}

void* QStandardItem_Background(void* ptr)
{
	return new QBrush(static_cast<QStandardItem*>(ptr)->background());
}

void* QStandardItem_Child(void* ptr, int row, int column)
{
	return static_cast<QStandardItem*>(ptr)->child(row, column);
}

int QStandardItem_Column(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->column();
}

int QStandardItem_ColumnCount(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->columnCount();
}

void* QStandardItem_Data(void* ptr, int role)
{
	return new QVariant(static_cast<QStandardItem*>(ptr)->data(role));
}

void* QStandardItem_DataDefault(void* ptr, int role)
{
		return new QVariant(static_cast<QStandardItem*>(ptr)->QStandardItem::data(role));
}

void* QStandardItem_Font(void* ptr)
{
	return new QFont(static_cast<QStandardItem*>(ptr)->font());
}

void* QStandardItem_Foreground(void* ptr)
{
	return new QBrush(static_cast<QStandardItem*>(ptr)->foreground());
}

char QStandardItem_HasChildren(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->hasChildren();
}

void* QStandardItem_Icon(void* ptr)
{
	return new QIcon(static_cast<QStandardItem*>(ptr)->icon());
}

void* QStandardItem_Index(void* ptr)
{
	return new QModelIndex(static_cast<QStandardItem*>(ptr)->index());
}

void QStandardItem_InsertColumn(void* ptr, int column, void* items)
{
	static_cast<QStandardItem*>(ptr)->insertColumn(column, *static_cast<QList<QStandardItem *>*>(items));
}

void QStandardItem_InsertColumns(void* ptr, int column, int count)
{
	static_cast<QStandardItem*>(ptr)->insertColumns(column, count);
}

void QStandardItem_InsertRow(void* ptr, int row, void* items)
{
	static_cast<QStandardItem*>(ptr)->insertRow(row, *static_cast<QList<QStandardItem *>*>(items));
}

void QStandardItem_InsertRow2(void* ptr, int row, void* item)
{
	static_cast<QStandardItem*>(ptr)->insertRow(row, static_cast<QStandardItem*>(item));
}

void QStandardItem_InsertRows(void* ptr, int row, void* items)
{
	static_cast<QStandardItem*>(ptr)->insertRows(row, *static_cast<QList<QStandardItem *>*>(items));
}

void QStandardItem_InsertRows2(void* ptr, int row, int count)
{
	static_cast<QStandardItem*>(ptr)->insertRows(row, count);
}

void* QStandardItem_Model(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->model();
}

void* QStandardItem_Parent(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->parent();
}

void QStandardItem_Read(void* ptr, void* in)
{
	static_cast<QStandardItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

void QStandardItem_ReadDefault(void* ptr, void* in)
{
		static_cast<QStandardItem*>(ptr)->QStandardItem::read(*static_cast<QDataStream*>(in));
}

void QStandardItem_RemoveColumn(void* ptr, int column)
{
	static_cast<QStandardItem*>(ptr)->removeColumn(column);
}

void QStandardItem_RemoveColumns(void* ptr, int column, int count)
{
	static_cast<QStandardItem*>(ptr)->removeColumns(column, count);
}

void QStandardItem_RemoveRow(void* ptr, int row)
{
	static_cast<QStandardItem*>(ptr)->removeRow(row);
}

void QStandardItem_RemoveRows(void* ptr, int row, int count)
{
	static_cast<QStandardItem*>(ptr)->removeRows(row, count);
}

int QStandardItem_Row(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->row();
}

int QStandardItem_RowCount(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->rowCount();
}

void QStandardItem_SetAccessibleDescription(void* ptr, struct QtGui_PackedString accessibleDescription)
{
	static_cast<QStandardItem*>(ptr)->setAccessibleDescription(QString::fromUtf8(accessibleDescription.data, accessibleDescription.len));
}

void QStandardItem_SetAccessibleText(void* ptr, struct QtGui_PackedString accessibleText)
{
	static_cast<QStandardItem*>(ptr)->setAccessibleText(QString::fromUtf8(accessibleText.data, accessibleText.len));
}

void QStandardItem_SetBackground(void* ptr, void* brush)
{
	static_cast<QStandardItem*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QStandardItem_SetCheckable(void* ptr, char checkable)
{
	static_cast<QStandardItem*>(ptr)->setCheckable(checkable != 0);
}

void QStandardItem_SetChild(void* ptr, int row, int column, void* item)
{
	static_cast<QStandardItem*>(ptr)->setChild(row, column, static_cast<QStandardItem*>(item));
}

void QStandardItem_SetChild2(void* ptr, int row, void* item)
{
	static_cast<QStandardItem*>(ptr)->setChild(row, static_cast<QStandardItem*>(item));
}

void QStandardItem_SetColumnCount(void* ptr, int columns)
{
	static_cast<QStandardItem*>(ptr)->setColumnCount(columns);
}

void QStandardItem_SetData(void* ptr, void* value, int role)
{
	static_cast<QStandardItem*>(ptr)->setData(*static_cast<QVariant*>(value), role);
}

void QStandardItem_SetDataDefault(void* ptr, void* value, int role)
{
		static_cast<QStandardItem*>(ptr)->QStandardItem::setData(*static_cast<QVariant*>(value), role);
}

void QStandardItem_SetEditable(void* ptr, char editable)
{
	static_cast<QStandardItem*>(ptr)->setEditable(editable != 0);
}

void QStandardItem_SetEnabled(void* ptr, char enabled)
{
	static_cast<QStandardItem*>(ptr)->setEnabled(enabled != 0);
}

void QStandardItem_SetFont(void* ptr, void* font)
{
	static_cast<QStandardItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QStandardItem_SetForeground(void* ptr, void* brush)
{
	static_cast<QStandardItem*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QStandardItem_SetIcon(void* ptr, void* icon)
{
	static_cast<QStandardItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QStandardItem_SetText(void* ptr, struct QtGui_PackedString text)
{
	static_cast<QStandardItem*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

void QStandardItem_SetToolTip(void* ptr, struct QtGui_PackedString toolTip)
{
	static_cast<QStandardItem*>(ptr)->setToolTip(QString::fromUtf8(toolTip.data, toolTip.len));
}

void* QStandardItem_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QStandardItem*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QStandardItem_TakeChild(void* ptr, int row, int column)
{
	return static_cast<QStandardItem*>(ptr)->takeChild(row, column);
}

struct QtGui_PackedString QStandardItem_Text(void* ptr)
{
	return ({ QByteArray* t847a7a = new QByteArray(static_cast<QStandardItem*>(ptr)->text().toUtf8()); QtGui_PackedString { const_cast<char*>(t847a7a->prepend("WHITESPACE").constData()+10), t847a7a->size()-10, t847a7a }; });
}

struct QtGui_PackedString QStandardItem_ToolTip(void* ptr)
{
	return ({ QByteArray* t938860 = new QByteArray(static_cast<QStandardItem*>(ptr)->toolTip().toUtf8()); QtGui_PackedString { const_cast<char*>(t938860->prepend("WHITESPACE").constData()+10), t938860->size()-10, t938860 }; });
}

int QStandardItem_Type(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->type();
}

int QStandardItem_TypeDefault(void* ptr)
{
		return static_cast<QStandardItem*>(ptr)->QStandardItem::type();
}

void QStandardItem_Write(void* ptr, void* out)
{
	static_cast<QStandardItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QStandardItem_WriteDefault(void* ptr, void* out)
{
		static_cast<QStandardItem*>(ptr)->QStandardItem::write(*static_cast<QDataStream*>(out));
}

void QStandardItem_DestroyQStandardItem(void* ptr)
{
	static_cast<QStandardItem*>(ptr)->~QStandardItem();
}

void QStandardItem_DestroyQStandardItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QStandardItem___appendColumn_items_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItem___appendColumn_items_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItem___appendColumn_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItem___appendRow_items_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItem___appendRow_items_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItem___appendRow_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItem___appendRows_items_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItem___appendRows_items_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItem___appendRows_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItem___insertColumn_items_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItem___insertColumn_items_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItem___insertColumn_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItem___insertRow_items_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItem___insertRow_items_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItem___insertRow_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItem___insertRows_items_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItem___insertRows_items_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItem___insertRows_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItem___takeColumn_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItem___takeColumn_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItem___takeColumn_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItem___takeRow_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItem___takeRow_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItem___takeRow_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

class MyQStandardItemModel: public QStandardItemModel
{
public:
	MyQStandardItemModel(QObject *parent = Q_NULLPTR) : QStandardItemModel(parent) {QStandardItemModel_QStandardItemModel_QRegisterMetaType();};
	MyQStandardItemModel(int rows, int columns, QObject *parent = Q_NULLPTR) : QStandardItemModel(rows, columns, parent) {QStandardItemModel_QStandardItemModel_QRegisterMetaType();};
	int columnCount(const QModelIndex & parent) const { return callbackQStandardItemModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQStandardItemModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQStandardItemModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQStandardItemModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQStandardItemModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQStandardItemModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex parent(const QModelIndex & child) const { return *static_cast<QModelIndex*>(callbackQStandardItemModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&child))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQStandardItemModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQStandardItemModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	int rowCount(const QModelIndex & parent) const { return callbackQStandardItemModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQStandardItemModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	 ~MyQStandardItemModel() { callbackQStandardItemModel_DestroyQStandardItemModel(this); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQStandardItemModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void childEvent(QChildEvent * event) { callbackQStandardItemModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQStandardItemModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQStandardItemModel_CustomEvent(this, event); };
	void deleteLater() { callbackQStandardItemModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQStandardItemModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQStandardItemModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQStandardItemModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQStandardItemModel_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQStandardItemModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQStandardItemModel_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QStandardItemModel*)
Q_DECLARE_METATYPE(MyQStandardItemModel*)

int QStandardItemModel_QStandardItemModel_QRegisterMetaType(){qRegisterMetaType<QStandardItemModel*>(); return qRegisterMetaType<MyQStandardItemModel*>();}

void* QStandardItemModel_NewQStandardItemModel(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQStandardItemModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQStandardItemModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQStandardItemModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQStandardItemModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQStandardItemModel(static_cast<QWindow*>(parent));
	} else {
		return new MyQStandardItemModel(static_cast<QObject*>(parent));
	}
}

void* QStandardItemModel_NewQStandardItemModel2(int rows, int columns, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQStandardItemModel(rows, columns, static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQStandardItemModel(rows, columns, static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQStandardItemModel(rows, columns, static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQStandardItemModel(rows, columns, static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQStandardItemModel(rows, columns, static_cast<QWindow*>(parent));
	} else {
		return new MyQStandardItemModel(rows, columns, static_cast<QObject*>(parent));
	}
}

void QStandardItemModel_AppendColumn(void* ptr, void* items)
{
	static_cast<QStandardItemModel*>(ptr)->appendColumn(*static_cast<QList<QStandardItem *>*>(items));
}

void QStandardItemModel_AppendRow(void* ptr, void* items)
{
	static_cast<QStandardItemModel*>(ptr)->appendRow(*static_cast<QList<QStandardItem *>*>(items));
}

void QStandardItemModel_AppendRow2(void* ptr, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->appendRow(static_cast<QStandardItem*>(item));
}

void QStandardItemModel_Clear(void* ptr)
{
	static_cast<QStandardItemModel*>(ptr)->clear();
}

int QStandardItemModel_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_ColumnCountDefault(void* ptr, void* parent)
{
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* QStandardItemModel_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QStandardItemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QStandardItemModel_DataDefault(void* ptr, void* index, int role)
{
		return new QVariant(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::data(*static_cast<QModelIndex*>(index), role));
}

char QStandardItemModel_HasChildrenDefault(void* ptr, void* parent)
{
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QStandardItemModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QStandardItemModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
		return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void QStandardItemModel_InsertColumn(void* ptr, int column, void* items)
{
	static_cast<QStandardItemModel*>(ptr)->insertColumn(column, *static_cast<QList<QStandardItem *>*>(items));
}

char QStandardItemModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void QStandardItemModel_InsertRow(void* ptr, int row, void* items)
{
	static_cast<QStandardItemModel*>(ptr)->insertRow(row, *static_cast<QList<QStandardItem *>*>(items));
}

void QStandardItemModel_InsertRow2(void* ptr, int row, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->insertRow(row, static_cast<QStandardItem*>(item));
}

char QStandardItemModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

void* QStandardItemModel_Item(void* ptr, int row, int column)
{
	return static_cast<QStandardItemModel*>(ptr)->item(row, column);
}

void* QStandardItemModel_ItemFromIndex(void* ptr, void* index)
{
	return static_cast<QStandardItemModel*>(ptr)->itemFromIndex(*static_cast<QModelIndex*>(index));
}

void* QStandardItemModel_Parent(void* ptr, void* child)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)));
}

void* QStandardItemModel_ParentDefault(void* ptr, void* child)
{
		return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::parent(*static_cast<QModelIndex*>(child)));
}

char QStandardItemModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QStandardItemModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RowCountDefault(void* ptr, void* parent)
{
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::rowCount(*static_cast<QModelIndex*>(parent));
}

void QStandardItemModel_SetColumnCount(void* ptr, int columns)
{
	static_cast<QStandardItemModel*>(ptr)->setColumnCount(columns);
}

char QStandardItemModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QStandardItemModel_SetHorizontalHeaderLabels(void* ptr, struct QtGui_PackedString labels)
{
	static_cast<QStandardItemModel*>(ptr)->setHorizontalHeaderLabels(QString::fromUtf8(labels.data, labels.len).split("¡¦!", QString::SkipEmptyParts));
}

void QStandardItemModel_SetItem(void* ptr, int row, int column, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->setItem(row, column, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetItem2(void* ptr, int row, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->setItem(row, static_cast<QStandardItem*>(item));
}

void* QStandardItemModel_TakeItem(void* ptr, int row, int column)
{
	return static_cast<QStandardItemModel*>(ptr)->takeItem(row, column);
}

void QStandardItemModel_DestroyQStandardItemModel(void* ptr)
{
	static_cast<QStandardItemModel*>(ptr)->~QStandardItemModel();
}

void QStandardItemModel_DestroyQStandardItemModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QStandardItemModel___appendColumn_items_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel___appendColumn_items_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItemModel___appendColumn_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItemModel___appendRow_items_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel___appendRow_items_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItemModel___appendRow_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItemModel___findItems_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel___findItems_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItemModel___findItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItemModel___insertColumn_items_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel___insertColumn_items_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItemModel___insertColumn_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItemModel___insertRow_items_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel___insertRow_items_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItemModel___insertRow_items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItemModel___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QStandardItemModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QStandardItemModel___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtGui_PackedList QStandardItemModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtGui_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* QStandardItemModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStandardItemModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QStandardItemModel___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QStandardItemModel___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QStandardItemModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QStandardItemModel___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtGui_PackedList QStandardItemModel___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtGui_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* QStandardItemModel___setItemRoleNames_roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QStandardItemModel___setItemRoleNames_roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QStandardItemModel___setItemRoleNames_roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtGui_PackedList QStandardItemModel___setItemRoleNames_roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtGui_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

void* QStandardItemModel___takeColumn_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel___takeColumn_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItemModel___takeColumn_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

void* QStandardItemModel___takeRow_atList(void* ptr, int i)
{
	return ({QStandardItem * tmp = static_cast<QList<QStandardItem *>*>(ptr)->at(i); if (i == static_cast<QList<QStandardItem *>*>(ptr)->size()-1) { static_cast<QList<QStandardItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel___takeRow_setList(void* ptr, void* i)
{
	static_cast<QList<QStandardItem *>*>(ptr)->append(static_cast<QStandardItem*>(i));
}

void* QStandardItemModel___takeRow_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStandardItem *>();
}

int QStandardItemModel_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QStandardItemModel_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QStandardItemModel_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QStandardItemModel_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QStandardItemModel_____setItemRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel_____setItemRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QStandardItemModel_____setItemRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QStandardItemModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStandardItemModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QStandardItemModel___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QStandardItemModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStandardItemModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QStandardItemModel___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int QStandardItemModel___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QStandardItemModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QStandardItemModel___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QStandardItemModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStandardItemModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QStandardItemModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QStandardItemModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStandardItemModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QStandardItemModel___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QStandardItemModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStandardItemModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QStandardItemModel___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QStandardItemModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStandardItemModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QStandardItemModel___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QStandardItemModel___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QStandardItemModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QStandardItemModel___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtGui_PackedList QStandardItemModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtGui_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

int QStandardItemModel_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QStandardItemModel_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QStandardItemModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QStandardItemModel_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QStandardItemModel_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QStandardItemModel_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QStandardItemModel___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QStandardItemModel___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QStandardItemModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStandardItemModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QStandardItemModel___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QStandardItemModel___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QStandardItemModel___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QStandardItemModel___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStandardItemModel___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QStandardItemModel___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

struct QtGui_PackedList QStandardItemModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
		return ({ QList<QModelIndex>* tmpValue0c2bda = new QList<QModelIndex>(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtGui_PackedList { tmpValue0c2bda, tmpValue0c2bda->size() }; });
}

void QStandardItemModel_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::childEvent(static_cast<QChildEvent*>(event));
}

void QStandardItemModel_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QStandardItemModel_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::customEvent(static_cast<QEvent*>(event));
}

void QStandardItemModel_DeleteLaterDefault(void* ptr)
{
		static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::deleteLater();
}

void QStandardItemModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QStandardItemModel_EventDefault(void* ptr, void* e)
{
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::event(static_cast<QEvent*>(e));
}

char QStandardItemModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QStandardItemModel_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQSurface: public QSurface
{
public:
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQSurface_Format(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize size() const { return *static_cast<QSize*>(callbackQSurface_Size(const_cast<void*>(static_cast<const void*>(this)))); };
	QSurface::SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQSurface_SurfaceType(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQSurface() { callbackQSurface_DestroyQSurface(this); };
};

Q_DECLARE_METATYPE(MyQSurface*)

int QSurface_QSurface_QRegisterMetaType(){qRegisterMetaType<QSurface*>(); return qRegisterMetaType<MyQSurface*>();}

void* QSurface_Format(void* ptr)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(ptr))) {
		return new QSurfaceFormat(static_cast<QWindow*>(ptr)->format());
	} else {
		return new QSurfaceFormat(static_cast<QSurface*>(ptr)->format());
	}
}

void* QSurface_Size(void* ptr)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QWindow*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QSurface*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

long long QSurface_SurfaceType(void* ptr)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(ptr))) {
		return static_cast<QWindow*>(ptr)->surfaceType();
	} else {
		return static_cast<QSurface*>(ptr)->surfaceType();
	}
}

void QSurface_DestroyQSurface(void* ptr)
{
	static_cast<QSurface*>(ptr)->~QSurface();
}

void QSurface_DestroyQSurfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

Q_DECLARE_METATYPE(QSurfaceFormat)
Q_DECLARE_METATYPE(QSurfaceFormat*)
void* QSurfaceFormat_NewQSurfaceFormat()
{
	return new QSurfaceFormat();
}

void* QSurfaceFormat_NewQSurfaceFormat2(long long options)
{
	return new QSurfaceFormat(static_cast<QSurfaceFormat::FormatOption>(options));
}

void* QSurfaceFormat_NewQSurfaceFormat3(void* other)
{
	return new QSurfaceFormat(*static_cast<QSurfaceFormat*>(other));
}

void QSurfaceFormat_DestroyQSurfaceFormat(void* ptr)
{
	static_cast<QSurfaceFormat*>(ptr)->~QSurfaceFormat();
}

Q_DECLARE_METATYPE(QTextBlock*)
void* QTextBlock_NewQTextBlock3(void* other)
{
	return new QTextBlock(*static_cast<QTextBlock*>(other));
}

void* QTextBlock_BlockFormat(void* ptr)
{
	return new QTextBlockFormat(static_cast<QTextBlock*>(ptr)->blockFormat());
}

void* QTextBlock_CharFormat(void* ptr)
{
	return new QTextCharFormat(static_cast<QTextBlock*>(ptr)->charFormat());
}

void* QTextBlock_Document(void* ptr)
{
	return const_cast<QTextDocument*>(static_cast<QTextBlock*>(ptr)->document());
}

char QTextBlock_IsValid(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->isValid();
}

void* QTextBlock_Layout(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->layout();
}

int QTextBlock_Length(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->length();
}

void* QTextBlock_Next(void* ptr)
{
	return new QTextBlock(static_cast<QTextBlock*>(ptr)->next());
}

int QTextBlock_Position(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->position();
}

struct QtGui_PackedString QTextBlock_Text(void* ptr)
{
	return ({ QByteArray* t4bcf6c = new QByteArray(static_cast<QTextBlock*>(ptr)->text().toUtf8()); QtGui_PackedString { const_cast<char*>(t4bcf6c->prepend("WHITESPACE").constData()+10), t4bcf6c->size()-10, t4bcf6c }; });
}

void* QTextBlock___textFormats_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QTextLayout::FormatRange>();
}

Q_DECLARE_METATYPE(QTextBlockFormat)
Q_DECLARE_METATYPE(QTextBlockFormat*)
void* QTextBlockFormat_NewQTextBlockFormat()
{
	return new QTextBlockFormat();
}

long long QTextBlockFormat_Alignment(void* ptr)
{
	return static_cast<QTextBlockFormat*>(ptr)->alignment();
}

int QTextBlockFormat_Indent(void* ptr)
{
	return static_cast<QTextBlockFormat*>(ptr)->indent();
}

void QTextBlockFormat_SetAlignment(void* ptr, long long alignment)
{
	static_cast<QTextBlockFormat*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextBlockFormat_SetIndent(void* ptr, int indentation)
{
	static_cast<QTextBlockFormat*>(ptr)->setIndent(indentation);
}

void QTextBlockFormat_SetTopMargin(void* ptr, double margin)
{
	static_cast<QTextBlockFormat*>(ptr)->setTopMargin(margin);
}

double QTextBlockFormat_TopMargin(void* ptr)
{
	return static_cast<QTextBlockFormat*>(ptr)->topMargin();
}

void* QTextBlockFormat___setTabPositions_tabs_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QTextOption::Tab>();
}

void* QTextBlockFormat___tabPositions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QTextOption::Tab>();
}

class MyQTextBlockGroup: public QTextBlockGroup
{
public:
	MyQTextBlockGroup(QTextDocument *document) : QTextBlockGroup(document) {QTextBlockGroup_QTextBlockGroup_QRegisterMetaType();};
	 ~MyQTextBlockGroup() { callbackQTextBlockGroup_DestroyQTextBlockGroup(this); };
	void childEvent(QChildEvent * event) { callbackQTextObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTextObject_CustomEvent(this, event); };
	void deleteLater() { callbackQTextObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQTextObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQTextObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQTextObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQTextObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QTextBlockGroup*)
Q_DECLARE_METATYPE(MyQTextBlockGroup*)

int QTextBlockGroup_QTextBlockGroup_QRegisterMetaType(){qRegisterMetaType<QTextBlockGroup*>(); return qRegisterMetaType<MyQTextBlockGroup*>();}

void* QTextBlockGroup_NewQTextBlockGroup(void* document)
{
	return new MyQTextBlockGroup(static_cast<QTextDocument*>(document));
}

void QTextBlockGroup_DestroyQTextBlockGroup(void* ptr)
{
	static_cast<QTextBlockGroup*>(ptr)->~QTextBlockGroup();
}

void QTextBlockGroup_DestroyQTextBlockGroupDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QTextBlockGroup___blockList_atList(void* ptr, int i)
{
	return new QTextBlock(({QTextBlock tmp = static_cast<QList<QTextBlock>*>(ptr)->at(i); if (i == static_cast<QList<QTextBlock>*>(ptr)->size()-1) { static_cast<QList<QTextBlock>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTextBlockGroup___blockList_setList(void* ptr, void* i)
{
	static_cast<QList<QTextBlock>*>(ptr)->append(*static_cast<QTextBlock*>(i));
}

void* QTextBlockGroup___blockList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QTextBlock>();
}

Q_DECLARE_METATYPE(QTextCharFormat)
Q_DECLARE_METATYPE(QTextCharFormat*)
void* QTextCharFormat_NewQTextCharFormat()
{
	return new QTextCharFormat();
}

void* QTextCharFormat_Font(void* ptr)
{
	return new QFont(static_cast<QTextCharFormat*>(ptr)->font());
}

char QTextCharFormat_FontItalic(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontItalic();
}

double QTextCharFormat_FontPointSize(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontPointSize();
}

char QTextCharFormat_FontStrikeOut(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontStrikeOut();
}

char QTextCharFormat_FontUnderline(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontUnderline();
}

int QTextCharFormat_FontWeight(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontWeight();
}

void QTextCharFormat_SetFont(void* ptr, void* font, long long behavior)
{
	static_cast<QTextCharFormat*>(ptr)->setFont(*static_cast<QFont*>(font), static_cast<QTextCharFormat::FontPropertiesInheritanceBehavior>(behavior));
}

void QTextCharFormat_SetFont2(void* ptr, void* font)
{
	static_cast<QTextCharFormat*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTextCharFormat_SetFontItalic(void* ptr, char italic)
{
	static_cast<QTextCharFormat*>(ptr)->setFontItalic(italic != 0);
}

void QTextCharFormat_SetFontPointSize(void* ptr, double size)
{
	static_cast<QTextCharFormat*>(ptr)->setFontPointSize(size);
}

void QTextCharFormat_SetFontStrikeOut(void* ptr, char strikeOut)
{
	static_cast<QTextCharFormat*>(ptr)->setFontStrikeOut(strikeOut != 0);
}

void QTextCharFormat_SetFontUnderline(void* ptr, char underline)
{
	static_cast<QTextCharFormat*>(ptr)->setFontUnderline(underline != 0);
}

void QTextCharFormat_SetFontWeight(void* ptr, int weight)
{
	static_cast<QTextCharFormat*>(ptr)->setFontWeight(weight);
}

void QTextCharFormat_SetToolTip(void* ptr, struct QtGui_PackedString text)
{
	static_cast<QTextCharFormat*>(ptr)->setToolTip(QString::fromUtf8(text.data, text.len));
}

struct QtGui_PackedString QTextCharFormat_ToolTip(void* ptr)
{
	return ({ QByteArray* t7e7f7f = new QByteArray(static_cast<QTextCharFormat*>(ptr)->toolTip().toUtf8()); QtGui_PackedString { const_cast<char*>(t7e7f7f->prepend("WHITESPACE").constData()+10), t7e7f7f->size()-10, t7e7f7f }; });
}

Q_DECLARE_METATYPE(QTextCursor)
Q_DECLARE_METATYPE(QTextCursor*)
void* QTextCursor_NewQTextCursor()
{
	return new QTextCursor();
}

void* QTextCursor_NewQTextCursor2(void* document)
{
	return new QTextCursor(static_cast<QTextDocument*>(document));
}

void* QTextCursor_NewQTextCursor3(void* frame)
{
	return new QTextCursor(static_cast<QTextFrame*>(frame));
}

void* QTextCursor_NewQTextCursor4(void* block)
{
	return new QTextCursor(*static_cast<QTextBlock*>(block));
}

void* QTextCursor_NewQTextCursor5(void* cursor)
{
	return new QTextCursor(*static_cast<QTextCursor*>(cursor));
}

int QTextCursor_Anchor(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->anchor();
}

void QTextCursor_BeginEditBlock(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->beginEditBlock();
}

void* QTextCursor_Block(void* ptr)
{
	return new QTextBlock(static_cast<QTextCursor*>(ptr)->block());
}

void* QTextCursor_BlockFormat(void* ptr)
{
	return new QTextBlockFormat(static_cast<QTextCursor*>(ptr)->blockFormat());
}

void* QTextCursor_CharFormat(void* ptr)
{
	return new QTextCharFormat(static_cast<QTextCursor*>(ptr)->charFormat());
}

void QTextCursor_ClearSelection(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->clearSelection();
}

void* QTextCursor_CreateList(void* ptr, void* format)
{
	return static_cast<QTextCursor*>(ptr)->createList(*static_cast<QTextListFormat*>(format));
}

void* QTextCursor_CreateList2(void* ptr, long long style)
{
	return static_cast<QTextCursor*>(ptr)->createList(static_cast<QTextListFormat::Style>(style));
}

void* QTextCursor_CurrentList(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->currentList();
}

void* QTextCursor_Document(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->document();
}

void QTextCursor_EndEditBlock(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->endEditBlock();
}

char QTextCursor_HasSelection(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->hasSelection();
}

void QTextCursor_InsertImage(void* ptr, void* format)
{
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QTextImageFormat*>(format));
}

void QTextCursor_InsertImage2(void* ptr, void* format, long long alignment)
{
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QTextImageFormat*>(format), static_cast<QTextFrameFormat::Position>(alignment));
}

void QTextCursor_InsertImage3(void* ptr, struct QtGui_PackedString name)
{
	static_cast<QTextCursor*>(ptr)->insertImage(QString::fromUtf8(name.data, name.len));
}

void QTextCursor_InsertImage4(void* ptr, void* image, struct QtGui_PackedString name)
{
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QImage*>(image), QString::fromUtf8(name.data, name.len));
}

void* QTextCursor_InsertTable(void* ptr, int rows, int columns, void* format)
{
	return static_cast<QTextCursor*>(ptr)->insertTable(rows, columns, *static_cast<QTextTableFormat*>(format));
}

void* QTextCursor_InsertTable2(void* ptr, int rows, int columns)
{
	return static_cast<QTextCursor*>(ptr)->insertTable(rows, columns);
}

void QTextCursor_InsertText(void* ptr, struct QtGui_PackedString text)
{
	static_cast<QTextCursor*>(ptr)->insertText(QString::fromUtf8(text.data, text.len));
}

void QTextCursor_InsertText2(void* ptr, struct QtGui_PackedString text, void* format)
{
	static_cast<QTextCursor*>(ptr)->insertText(QString::fromUtf8(text.data, text.len), *static_cast<QTextCharFormat*>(format));
}

void QTextCursor_MergeCharFormat(void* ptr, void* modifier)
{
	static_cast<QTextCursor*>(ptr)->mergeCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

int QTextCursor_Position(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->position();
}

void QTextCursor_RemoveSelectedText(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->removeSelectedText();
}

void QTextCursor_Select(void* ptr, long long selection)
{
	static_cast<QTextCursor*>(ptr)->select(static_cast<QTextCursor::SelectionType>(selection));
}

struct QtGui_PackedString QTextCursor_SelectedText(void* ptr)
{
	return ({ QByteArray* t495ef8 = new QByteArray(static_cast<QTextCursor*>(ptr)->selectedText().toUtf8()); QtGui_PackedString { const_cast<char*>(t495ef8->prepend("WHITESPACE").constData()+10), t495ef8->size()-10, t495ef8 }; });
}

void* QTextCursor_Selection(void* ptr)
{
	return new QTextDocumentFragment(static_cast<QTextCursor*>(ptr)->selection());
}

int QTextCursor_SelectionEnd(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->selectionEnd();
}

int QTextCursor_SelectionStart(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->selectionStart();
}

void QTextCursor_SetBlockFormat(void* ptr, void* format)
{
	static_cast<QTextCursor*>(ptr)->setBlockFormat(*static_cast<QTextBlockFormat*>(format));
}

void QTextCursor_SetCharFormat(void* ptr, void* format)
{
	static_cast<QTextCursor*>(ptr)->setCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextCursor_SetPosition(void* ptr, int pos, long long m)
{
	static_cast<QTextCursor*>(ptr)->setPosition(pos, static_cast<QTextCursor::MoveMode>(m));
}

void QTextCursor_DestroyQTextCursor(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->~QTextCursor();
}

class MyQTextDocument: public QTextDocument
{
public:
	MyQTextDocument(QObject *parent = Q_NULLPTR) : QTextDocument(parent) {QTextDocument_QTextDocument_QRegisterMetaType();};
	MyQTextDocument(const QString &text, QObject *parent = Q_NULLPTR) : QTextDocument(text, parent) {QTextDocument_QTextDocument_QRegisterMetaType();};
	void clear() { callbackQTextDocument_Clear(this); };
	void setModified(bool m) { callbackQTextDocument_SetModified(this, m); };
	 ~MyQTextDocument() { callbackQTextDocument_DestroyQTextDocument(this); };
	void childEvent(QChildEvent * event) { callbackQTextDocument_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextDocument_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTextDocument_CustomEvent(this, event); };
	void deleteLater() { callbackQTextDocument_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQTextDocument_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextDocument_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQTextDocument_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextDocument_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQTextDocument_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQTextDocument_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QTextDocument*)
Q_DECLARE_METATYPE(MyQTextDocument*)

int QTextDocument_QTextDocument_QRegisterMetaType(){qRegisterMetaType<QTextDocument*>(); return qRegisterMetaType<MyQTextDocument*>();}

void* QTextDocument_NewQTextDocument(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTextDocument(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextDocument(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTextDocument(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextDocument(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextDocument(static_cast<QWindow*>(parent));
	} else {
		return new MyQTextDocument(static_cast<QObject*>(parent));
	}
}

void* QTextDocument_NewQTextDocument2(struct QtGui_PackedString text, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTextDocument(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextDocument(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTextDocument(QString::fromUtf8(text.data, text.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTextDocument(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTextDocument(QString::fromUtf8(text.data, text.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQTextDocument(QString::fromUtf8(text.data, text.len), static_cast<QObject*>(parent));
	}
}

void QTextDocument_AddResource(void* ptr, int ty, void* name, void* resource)
{
	static_cast<QTextDocument*>(ptr)->addResource(ty, *static_cast<QUrl*>(name), *static_cast<QVariant*>(resource));
}

void* QTextDocument_Begin(void* ptr)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->begin());
}

void QTextDocument_Clear(void* ptr)
{
	static_cast<QTextDocument*>(ptr)->clear();
}

void QTextDocument_ClearDefault(void* ptr)
{
		static_cast<QTextDocument*>(ptr)->QTextDocument::clear();
}

void* QTextDocument_End(void* ptr)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->end());
}

void* QTextDocument_Find(void* ptr, struct QtGui_PackedString subString, void* cursor, long long options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(QString::fromUtf8(subString.data, subString.len), *static_cast<QTextCursor*>(cursor), static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_Find2(void* ptr, struct QtGui_PackedString subString, int position, long long options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(QString::fromUtf8(subString.data, subString.len), position, static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_Find3(void* ptr, void* expr, int from, long long options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(*static_cast<QRegExp*>(expr), from, static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_Find4(void* ptr, void* expr, void* cursor, long long options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(*static_cast<QRegExp*>(expr), *static_cast<QTextCursor*>(cursor), static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_Find5(void* ptr, void* expr, int from, long long options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(*static_cast<QRegularExpression*>(expr), from, static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_Find6(void* ptr, void* expr, void* cursor, long long options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(*static_cast<QRegularExpression*>(expr), *static_cast<QTextCursor*>(cursor), static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_FindBlock(void* ptr, int pos)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->findBlock(pos));
}

void* QTextDocument_FirstBlock(void* ptr)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->firstBlock());
}

char QTextDocument_IsEmpty(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->isEmpty();
}

char QTextDocument_IsModified(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->isModified();
}

void* QTextDocument_Object(void* ptr, int objectIndex)
{
	return static_cast<QTextDocument*>(ptr)->object(objectIndex);
}

void QTextDocument_Print(void* ptr, void* printer)
{
#ifndef Q_OS_IOS
	static_cast<QTextDocument*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
#endif
}

void* QTextDocument_Resource(void* ptr, int ty, void* name)
{
	return new QVariant(static_cast<QTextDocument*>(ptr)->resource(ty, *static_cast<QUrl*>(name)));
}

void* QTextDocument_RootFrame(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->rootFrame();
}

void QTextDocument_SetHtml(void* ptr, struct QtGui_PackedString html)
{
	static_cast<QTextDocument*>(ptr)->setHtml(QString::fromUtf8(html.data, html.len));
}

void QTextDocument_SetModified(void* ptr, char m)
{
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "setModified", Q_ARG(bool, m != 0));
}

void QTextDocument_SetModifiedDefault(void* ptr, char m)
{
		static_cast<QTextDocument*>(ptr)->QTextDocument::setModified(m != 0);
}

void* QTextDocument_Size(void* ptr)
{
	return ({ QSizeF tmpValue = static_cast<QTextDocument*>(ptr)->size(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

struct QtGui_PackedString QTextDocument_ToHtml(void* ptr, void* encoding)
{
	return ({ QByteArray* t48be10 = new QByteArray(static_cast<QTextDocument*>(ptr)->toHtml(*static_cast<QByteArray*>(encoding)).toUtf8()); QtGui_PackedString { const_cast<char*>(t48be10->prepend("WHITESPACE").constData()+10), t48be10->size()-10, t48be10 }; });
}

struct QtGui_PackedString QTextDocument_ToPlainText(void* ptr)
{
	return ({ QByteArray* t7d8da3 = new QByteArray(static_cast<QTextDocument*>(ptr)->toPlainText().toUtf8()); QtGui_PackedString { const_cast<char*>(t7d8da3->prepend("WHITESPACE").constData()+10), t7d8da3->size()-10, t7d8da3 }; });
}

void QTextDocument_DestroyQTextDocument(void* ptr)
{
	static_cast<QTextDocument*>(ptr)->~QTextDocument();
}

void QTextDocument_DestroyQTextDocumentDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QTextDocument___allFormats_atList(void* ptr, int i)
{
	return new QTextFormat(({QTextFormat tmp = static_cast<QVector<QTextFormat>*>(ptr)->at(i); if (i == static_cast<QVector<QTextFormat>*>(ptr)->size()-1) { static_cast<QVector<QTextFormat>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QTextDocument___allFormats_setList(void* ptr, void* i)
{
	static_cast<QVector<QTextFormat>*>(ptr)->append(*static_cast<QTextFormat*>(i));
}

void* QTextDocument___allFormats_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QTextFormat>();
}

void* QTextDocument___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextDocument___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTextDocument___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QTextDocument___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTextDocument___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTextDocument___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTextDocument___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextDocument___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTextDocument___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QTextDocument___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextDocument___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTextDocument___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QTextDocument_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QTextDocument*>(ptr)->QTextDocument::childEvent(static_cast<QChildEvent*>(event));
}

void QTextDocument_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QTextDocument*>(ptr)->QTextDocument::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextDocument_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QTextDocument*>(ptr)->QTextDocument::customEvent(static_cast<QEvent*>(event));
}

void QTextDocument_DeleteLaterDefault(void* ptr)
{
		static_cast<QTextDocument*>(ptr)->QTextDocument::deleteLater();
}

void QTextDocument_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QTextDocument*>(ptr)->QTextDocument::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QTextDocument_EventDefault(void* ptr, void* e)
{
		return static_cast<QTextDocument*>(ptr)->QTextDocument::event(static_cast<QEvent*>(e));
}

char QTextDocument_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QTextDocument*>(ptr)->QTextDocument::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QTextDocument*>(ptr)->QTextDocument::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QTextDocument_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QTextDocument*>(ptr)->QTextDocument::timerEvent(static_cast<QTimerEvent*>(event));
}

Q_DECLARE_METATYPE(QTextDocumentFragment)
Q_DECLARE_METATYPE(QTextDocumentFragment*)
void* QTextDocumentFragment_NewQTextDocumentFragment()
{
	return new QTextDocumentFragment();
}

void* QTextDocumentFragment_NewQTextDocumentFragment2(void* document)
{
	return new QTextDocumentFragment(static_cast<QTextDocument*>(document));
}

void* QTextDocumentFragment_NewQTextDocumentFragment3(void* cursor)
{
	return new QTextDocumentFragment(*static_cast<QTextCursor*>(cursor));
}

void* QTextDocumentFragment_NewQTextDocumentFragment4(void* other)
{
	return new QTextDocumentFragment(*static_cast<QTextDocumentFragment*>(other));
}

char QTextDocumentFragment_IsEmpty(void* ptr)
{
	return static_cast<QTextDocumentFragment*>(ptr)->isEmpty();
}

struct QtGui_PackedString QTextDocumentFragment_ToHtml(void* ptr, void* encoding)
{
	return ({ QByteArray* t6dea0f = new QByteArray(static_cast<QTextDocumentFragment*>(ptr)->toHtml(*static_cast<QByteArray*>(encoding)).toUtf8()); QtGui_PackedString { const_cast<char*>(t6dea0f->prepend("WHITESPACE").constData()+10), t6dea0f->size()-10, t6dea0f }; });
}

struct QtGui_PackedString QTextDocumentFragment_ToPlainText(void* ptr)
{
	return ({ QByteArray* td3666e = new QByteArray(static_cast<QTextDocumentFragment*>(ptr)->toPlainText().toUtf8()); QtGui_PackedString { const_cast<char*>(td3666e->prepend("WHITESPACE").constData()+10), td3666e->size()-10, td3666e }; });
}

void QTextDocumentFragment_DestroyQTextDocumentFragment(void* ptr)
{
	static_cast<QTextDocumentFragment*>(ptr)->~QTextDocumentFragment();
}

Q_DECLARE_METATYPE(QTextDocumentWriter*)
void* QTextDocumentWriter_NewQTextDocumentWriter()
{
	return new QTextDocumentWriter();
}

void* QTextDocumentWriter_NewQTextDocumentWriter2(void* device, void* format)
{
	return new QTextDocumentWriter(static_cast<QIODevice*>(device), *static_cast<QByteArray*>(format));
}

void* QTextDocumentWriter_NewQTextDocumentWriter3(struct QtGui_PackedString fileName, void* format)
{
	return new QTextDocumentWriter(QString::fromUtf8(fileName.data, fileName.len), *static_cast<QByteArray*>(format));
}

void* QTextDocumentWriter_Device(void* ptr)
{
	return static_cast<QTextDocumentWriter*>(ptr)->device();
}

struct QtGui_PackedString QTextDocumentWriter_FileName(void* ptr)
{
	return ({ QByteArray* tbc6b5c = new QByteArray(static_cast<QTextDocumentWriter*>(ptr)->fileName().toUtf8()); QtGui_PackedString { const_cast<char*>(tbc6b5c->prepend("WHITESPACE").constData()+10), tbc6b5c->size()-10, tbc6b5c }; });
}

void* QTextDocumentWriter_Format(void* ptr)
{
	return new QByteArray(static_cast<QTextDocumentWriter*>(ptr)->format());
}

char QTextDocumentWriter_Write(void* ptr, void* document)
{
	return static_cast<QTextDocumentWriter*>(ptr)->write(static_cast<QTextDocument*>(document));
}

char QTextDocumentWriter_Write2(void* ptr, void* fragment)
{
	return static_cast<QTextDocumentWriter*>(ptr)->write(*static_cast<QTextDocumentFragment*>(fragment));
}

void QTextDocumentWriter_DestroyQTextDocumentWriter(void* ptr)
{
	static_cast<QTextDocumentWriter*>(ptr)->~QTextDocumentWriter();
}

void* QTextDocumentWriter___supportedDocumentFormats_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTextDocumentWriter___supportedDocumentFormats_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTextDocumentWriter___supportedDocumentFormats_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

Q_DECLARE_METATYPE(QTextFormat)
Q_DECLARE_METATYPE(QTextFormat*)
void* QTextFormat_NewQTextFormat()
{
	return new QTextFormat();
}

void* QTextFormat_NewQTextFormat2(int ty)
{
	return new QTextFormat(ty);
}

void* QTextFormat_NewQTextFormat3(void* other)
{
	return new QTextFormat(*static_cast<QTextFormat*>(other));
}

void* QTextFormat_Background(void* ptr)
{
	return new QBrush(static_cast<QTextFormat*>(ptr)->background());
}

void QTextFormat_ClearBackground(void* ptr)
{
	static_cast<QTextFormat*>(ptr)->clearBackground();
}

void QTextFormat_ClearForeground(void* ptr)
{
	static_cast<QTextFormat*>(ptr)->clearForeground();
}

void* QTextFormat_Foreground(void* ptr)
{
	return new QBrush(static_cast<QTextFormat*>(ptr)->foreground());
}

char QTextFormat_IsEmpty(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isEmpty();
}

char QTextFormat_IsTableCellFormat(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isTableCellFormat();
}

char QTextFormat_IsTableFormat(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isTableFormat();
}

char QTextFormat_IsValid(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isValid();
}

void QTextFormat_Merge(void* ptr, void* other)
{
	static_cast<QTextFormat*>(ptr)->merge(*static_cast<QTextFormat*>(other));
}

int QTextFormat_ObjectIndex(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->objectIndex();
}

void* QTextFormat_Property(void* ptr, int propertyId)
{
	return new QVariant(static_cast<QTextFormat*>(ptr)->property(propertyId));
}

void QTextFormat_SetBackground(void* ptr, void* brush)
{
	static_cast<QTextFormat*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QTextFormat_SetForeground(void* ptr, void* brush)
{
	static_cast<QTextFormat*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QTextFormat_SetObjectIndex(void* ptr, int index)
{
	static_cast<QTextFormat*>(ptr)->setObjectIndex(index);
}

int QTextFormat_Type(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->type();
}

void QTextFormat_DestroyQTextFormat(void* ptr)
{
	static_cast<QTextFormat*>(ptr)->~QTextFormat();
}

void* QTextFormat___lengthVectorProperty_atList(void* ptr, int i)
{
	return new QTextLength(({QTextLength tmp = static_cast<QVector<QTextLength>*>(ptr)->at(i); if (i == static_cast<QVector<QTextLength>*>(ptr)->size()-1) { static_cast<QVector<QTextLength>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QTextFormat___lengthVectorProperty_setList(void* ptr, void* i)
{
	static_cast<QVector<QTextLength>*>(ptr)->append(*static_cast<QTextLength*>(i));
}

void* QTextFormat___lengthVectorProperty_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QTextLength>();
}

void* QTextFormat___properties_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QTextFormat___properties_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QTextFormat___properties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtGui_PackedList QTextFormat___properties_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtGui_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* QTextFormat___setProperty_value_atList2(void* ptr, int i)
{
	return new QTextLength(({QTextLength tmp = static_cast<QVector<QTextLength>*>(ptr)->at(i); if (i == static_cast<QVector<QTextLength>*>(ptr)->size()-1) { static_cast<QVector<QTextLength>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QTextFormat___setProperty_value_setList2(void* ptr, void* i)
{
	static_cast<QVector<QTextLength>*>(ptr)->append(*static_cast<QTextLength*>(i));
}

void* QTextFormat___setProperty_value_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QTextLength>();
}

int QTextFormat_____properties_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextFormat_____properties_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QTextFormat_____properties_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

class MyQTextFrame: public QTextFrame
{
public:
	MyQTextFrame(QTextDocument *document) : QTextFrame(document) {QTextFrame_QTextFrame_QRegisterMetaType();};
	 ~MyQTextFrame() { callbackQTextFrame_DestroyQTextFrame(this); };
	void childEvent(QChildEvent * event) { callbackQTextObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTextObject_CustomEvent(this, event); };
	void deleteLater() { callbackQTextObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQTextObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQTextObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQTextObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQTextObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QTextFrame*)
Q_DECLARE_METATYPE(MyQTextFrame*)

int QTextFrame_QTextFrame_QRegisterMetaType(){qRegisterMetaType<QTextFrame*>(); return qRegisterMetaType<MyQTextFrame*>();}

void* QTextFrame_NewQTextFrame(void* document)
{
	return new MyQTextFrame(static_cast<QTextDocument*>(document));
}

struct QtGui_PackedList QTextFrame_ChildFrames(void* ptr)
{
	return ({ QList<QTextFrame *>* tmpValue45166f = new QList<QTextFrame *>(static_cast<QTextFrame*>(ptr)->childFrames()); QtGui_PackedList { tmpValue45166f, tmpValue45166f->size() }; });
}

void* QTextFrame_FirstCursorPosition(void* ptr)
{
	return new QTextCursor(static_cast<QTextFrame*>(ptr)->firstCursorPosition());
}

void* QTextFrame_FrameFormat(void* ptr)
{
	return new QTextFrameFormat(static_cast<QTextFrame*>(ptr)->frameFormat());
}

void QTextFrame_DestroyQTextFrame(void* ptr)
{
	static_cast<QTextFrame*>(ptr)->~QTextFrame();
}

void QTextFrame_DestroyQTextFrameDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QTextFrame___childFrames_atList(void* ptr, int i)
{
	return ({QTextFrame * tmp = static_cast<QList<QTextFrame *>*>(ptr)->at(i); if (i == static_cast<QList<QTextFrame *>*>(ptr)->size()-1) { static_cast<QList<QTextFrame *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextFrame___childFrames_setList(void* ptr, void* i)
{
	static_cast<QList<QTextFrame *>*>(ptr)->append(static_cast<QTextFrame*>(i));
}

void* QTextFrame___childFrames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QTextFrame *>();
}

Q_DECLARE_METATYPE(QTextFrameFormat)
Q_DECLARE_METATYPE(QTextFrameFormat*)
void* QTextFrameFormat_NewQTextFrameFormat()
{
	return new QTextFrameFormat();
}

double QTextFrameFormat_Border(void* ptr)
{
	return static_cast<QTextFrameFormat*>(ptr)->border();
}

void* QTextFrameFormat_BorderBrush(void* ptr)
{
	return new QBrush(static_cast<QTextFrameFormat*>(ptr)->borderBrush());
}

void* QTextFrameFormat_Height(void* ptr)
{
	return new QTextLength(static_cast<QTextFrameFormat*>(ptr)->height());
}

double QTextFrameFormat_Margin(void* ptr)
{
	return static_cast<QTextFrameFormat*>(ptr)->margin();
}

long long QTextFrameFormat_Position(void* ptr)
{
	return static_cast<QTextFrameFormat*>(ptr)->position();
}

void QTextFrameFormat_SetBorder(void* ptr, double width)
{
	static_cast<QTextFrameFormat*>(ptr)->setBorder(width);
}

void QTextFrameFormat_SetBorderBrush(void* ptr, void* brush)
{
	static_cast<QTextFrameFormat*>(ptr)->setBorderBrush(*static_cast<QBrush*>(brush));
}

void QTextFrameFormat_SetPosition(void* ptr, long long policy)
{
	static_cast<QTextFrameFormat*>(ptr)->setPosition(static_cast<QTextFrameFormat::Position>(policy));
}

void QTextFrameFormat_SetTopMargin(void* ptr, double margin)
{
	static_cast<QTextFrameFormat*>(ptr)->setTopMargin(margin);
}

double QTextFrameFormat_TopMargin(void* ptr)
{
	return static_cast<QTextFrameFormat*>(ptr)->topMargin();
}

void* QTextFrameFormat_Width(void* ptr)
{
	return new QTextLength(static_cast<QTextFrameFormat*>(ptr)->width());
}

Q_DECLARE_METATYPE(QTextImageFormat)
Q_DECLARE_METATYPE(QTextImageFormat*)
void* QTextImageFormat_NewQTextImageFormat()
{
	return new QTextImageFormat();
}

double QTextImageFormat_Height(void* ptr)
{
	return static_cast<QTextImageFormat*>(ptr)->height();
}

struct QtGui_PackedString QTextImageFormat_Name(void* ptr)
{
	return ({ QByteArray* t290038 = new QByteArray(static_cast<QTextImageFormat*>(ptr)->name().toUtf8()); QtGui_PackedString { const_cast<char*>(t290038->prepend("WHITESPACE").constData()+10), t290038->size()-10, t290038 }; });
}

double QTextImageFormat_Width(void* ptr)
{
	return static_cast<QTextImageFormat*>(ptr)->width();
}

Q_DECLARE_METATYPE(QTextLayout*)
void* QTextLayout_NewQTextLayout()
{
	return new QTextLayout();
}

void* QTextLayout_NewQTextLayout2(struct QtGui_PackedString text)
{
	return new QTextLayout(QString::fromUtf8(text.data, text.len));
}

void* QTextLayout_NewQTextLayout4(struct QtGui_PackedString text, void* font, void* paintdevice)
{
	if (dynamic_cast<QWidget*>(static_cast<QObject*>(paintdevice))) {
		return new QTextLayout(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), static_cast<QWidget*>(paintdevice));
	} else {
		return new QTextLayout(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
	}
}

void* QTextLayout_BoundingRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QTextLayout*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QTextLayout_Font(void* ptr)
{
	return new QFont(static_cast<QTextLayout*>(ptr)->font());
}

double QTextLayout_MaximumWidth(void* ptr)
{
	return static_cast<QTextLayout*>(ptr)->maximumWidth();
}

double QTextLayout_MinimumWidth(void* ptr)
{
	return static_cast<QTextLayout*>(ptr)->minimumWidth();
}

void* QTextLayout_Position(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QTextLayout*>(ptr)->position(); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QTextLayout_SetFont(void* ptr, void* font)
{
	static_cast<QTextLayout*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTextLayout_SetPosition(void* ptr, void* p)
{
	static_cast<QTextLayout*>(ptr)->setPosition(*static_cast<QPointF*>(p));
}

void QTextLayout_SetText(void* ptr, struct QtGui_PackedString stri)
{
	static_cast<QTextLayout*>(ptr)->setText(QString::fromUtf8(stri.data, stri.len));
}

struct QtGui_PackedString QTextLayout_Text(void* ptr)
{
	return ({ QByteArray* t7e8d3c = new QByteArray(static_cast<QTextLayout*>(ptr)->text().toUtf8()); QtGui_PackedString { const_cast<char*>(t7e8d3c->prepend("WHITESPACE").constData()+10), t7e8d3c->size()-10, t7e8d3c }; });
}

void QTextLayout_DestroyQTextLayout(void* ptr)
{
	static_cast<QTextLayout*>(ptr)->~QTextLayout();
}

void* QTextLayout___formats_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QTextLayout::FormatRange>();
}

void* QTextLayout___glyphRuns_atList(void* ptr, int i)
{
	return new QGlyphRun(({QGlyphRun tmp = static_cast<QList<QGlyphRun>*>(ptr)->at(i); if (i == static_cast<QList<QGlyphRun>*>(ptr)->size()-1) { static_cast<QList<QGlyphRun>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTextLayout___glyphRuns_setList(void* ptr, void* i)
{
	static_cast<QList<QGlyphRun>*>(ptr)->append(*static_cast<QGlyphRun*>(i));
}

void* QTextLayout___glyphRuns_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGlyphRun>();
}

Q_DECLARE_METATYPE(QTextLength)
Q_DECLARE_METATYPE(QTextLength*)
void* QTextLength_NewQTextLength()
{
	return new QTextLength();
}

void* QTextLength_NewQTextLength2(long long ty, double value)
{
	return new QTextLength(static_cast<QTextLength::Type>(ty), value);
}

long long QTextLength_Type(void* ptr)
{
	return static_cast<QTextLength*>(ptr)->type();
}

double QTextLength_Value(void* ptr, double maximumLength)
{
	return static_cast<QTextLength*>(ptr)->value(maximumLength);
}

class MyQTextList: public QTextList
{
public:
	void childEvent(QChildEvent * event) { callbackQTextObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTextObject_CustomEvent(this, event); };
	void deleteLater() { callbackQTextObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQTextObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQTextObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQTextObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQTextObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QTextList*)
Q_DECLARE_METATYPE(MyQTextList*)

int QTextList_QTextList_QRegisterMetaType(){qRegisterMetaType<QTextList*>(); return qRegisterMetaType<MyQTextList*>();}

void QTextList_Add(void* ptr, void* block)
{
	static_cast<QTextList*>(ptr)->add(*static_cast<QTextBlock*>(block));
}

int QTextList_Count(void* ptr)
{
	return static_cast<QTextList*>(ptr)->count();
}

void* QTextList_Format(void* ptr)
{
	return new QTextListFormat(static_cast<QTextList*>(ptr)->format());
}

void* QTextList_Item(void* ptr, int i)
{
	return new QTextBlock(static_cast<QTextList*>(ptr)->item(i));
}

void QTextList_Remove(void* ptr, void* block)
{
	static_cast<QTextList*>(ptr)->remove(*static_cast<QTextBlock*>(block));
}

Q_DECLARE_METATYPE(QTextListFormat)
Q_DECLARE_METATYPE(QTextListFormat*)
void* QTextListFormat_NewQTextListFormat()
{
	return new QTextListFormat();
}

int QTextListFormat_Indent(void* ptr)
{
	return static_cast<QTextListFormat*>(ptr)->indent();
}

void QTextListFormat_SetIndent(void* ptr, int indentation)
{
	static_cast<QTextListFormat*>(ptr)->setIndent(indentation);
}

void QTextListFormat_SetStyle(void* ptr, long long style)
{
	static_cast<QTextListFormat*>(ptr)->setStyle(static_cast<QTextListFormat::Style>(style));
}

long long QTextListFormat_Style(void* ptr)
{
	return static_cast<QTextListFormat*>(ptr)->style();
}

class MyQTextObject: public QTextObject
{
public:
	MyQTextObject(QTextDocument *document) : QTextObject(document) {QTextObject_QTextObject_QRegisterMetaType();};
	 ~MyQTextObject() { callbackQTextObject_DestroyQTextObject(this); };
	void childEvent(QChildEvent * event) { callbackQTextObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTextObject_CustomEvent(this, event); };
	void deleteLater() { callbackQTextObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQTextObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQTextObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQTextObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQTextObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QTextObject*)
Q_DECLARE_METATYPE(MyQTextObject*)

int QTextObject_QTextObject_QRegisterMetaType(){qRegisterMetaType<QTextObject*>(); return qRegisterMetaType<MyQTextObject*>();}

void* QTextObject_NewQTextObject(void* document)
{
	return new MyQTextObject(static_cast<QTextDocument*>(document));
}

void* QTextObject_Document(void* ptr)
{
	return static_cast<QTextObject*>(ptr)->document();
}

void* QTextObject_Format(void* ptr)
{
	return new QTextFormat(static_cast<QTextObject*>(ptr)->format());
}

int QTextObject_ObjectIndex(void* ptr)
{
	return static_cast<QTextObject*>(ptr)->objectIndex();
}

void QTextObject_DestroyQTextObject(void* ptr)
{
	static_cast<QTextObject*>(ptr)->~QTextObject();
}

void QTextObject_DestroyQTextObjectDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QTextObject___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextObject___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTextObject___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QTextObject___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTextObject___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTextObject___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTextObject___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextObject___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTextObject___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QTextObject___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextObject___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTextObject___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QTextObject_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTextTable*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextTable*>(ptr)->QTextTable::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QTextFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextFrame*>(ptr)->QTextFrame::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QTextList*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextList*>(ptr)->QTextList::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QTextBlockGroup*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QTextObject*>(ptr)->QTextObject::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QTextObject_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QTextTable*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextTable*>(ptr)->QTextTable::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTextFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextFrame*>(ptr)->QTextFrame::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTextList*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextList*>(ptr)->QTextList::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTextBlockGroup*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QTextObject*>(ptr)->QTextObject::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QTextObject_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTextTable*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextTable*>(ptr)->QTextTable::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTextFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextFrame*>(ptr)->QTextFrame::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTextList*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextList*>(ptr)->QTextList::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTextBlockGroup*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QTextObject*>(ptr)->QTextObject::customEvent(static_cast<QEvent*>(event));
	}
}

void QTextObject_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QTextTable*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextTable*>(ptr)->QTextTable::deleteLater();
	} else if (dynamic_cast<QTextFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextFrame*>(ptr)->QTextFrame::deleteLater();
	} else if (dynamic_cast<QTextList*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextList*>(ptr)->QTextList::deleteLater();
	} else if (dynamic_cast<QTextBlockGroup*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::deleteLater();
	} else {
		static_cast<QTextObject*>(ptr)->QTextObject::deleteLater();
	}
}

void QTextObject_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QTextTable*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextTable*>(ptr)->QTextTable::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTextFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextFrame*>(ptr)->QTextFrame::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTextList*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextList*>(ptr)->QTextList::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTextBlockGroup*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QTextObject*>(ptr)->QTextObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QTextObject_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QTextTable*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextTable*>(ptr)->QTextTable::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QTextFrame*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextFrame*>(ptr)->QTextFrame::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QTextList*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextList*>(ptr)->QTextList::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QTextBlockGroup*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QTextObject*>(ptr)->QTextObject::event(static_cast<QEvent*>(e));
	}
}

char QTextObject_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QTextTable*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
			return static_cast<QTextTable*>(ptr)->QTextTable::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QTextTable*>(ptr)->QTextTable::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QTextFrame*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
			return static_cast<QTextFrame*>(ptr)->QTextFrame::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QTextFrame*>(ptr)->QTextFrame::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QTextList*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
			return static_cast<QTextList*>(ptr)->QTextList::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QTextList*>(ptr)->QTextList::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QTextBlockGroup*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
			return static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
			return static_cast<QTextObject*>(ptr)->QTextObject::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QTextObject*>(ptr)->QTextObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void QTextObject_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTextTable*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextTable*>(ptr)->QTextTable::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QTextFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextFrame*>(ptr)->QTextFrame::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QTextList*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextList*>(ptr)->QTextList::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QTextBlockGroup*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QTextObject*>(ptr)->QTextObject::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

Q_DECLARE_METATYPE(QTextOption)
Q_DECLARE_METATYPE(QTextOption*)
void* QTextOption_NewQTextOption()
{
	return new QTextOption();
}

void* QTextOption_NewQTextOption2(long long alignment)
{
	return new QTextOption(static_cast<Qt::AlignmentFlag>(alignment));
}

void* QTextOption_NewQTextOption3(void* other)
{
	return new QTextOption(*static_cast<QTextOption*>(other));
}

long long QTextOption_Alignment(void* ptr)
{
	return static_cast<QTextOption*>(ptr)->alignment();
}

void QTextOption_SetAlignment(void* ptr, long long alignment)
{
	static_cast<QTextOption*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextOption_DestroyQTextOption(void* ptr)
{
	static_cast<QTextOption*>(ptr)->~QTextOption();
}

double QTextOption___setTabArray_tabStops_atList(void* ptr, int i)
{
	return ({qreal tmp = static_cast<QList<qreal>*>(ptr)->at(i); if (i == static_cast<QList<qreal>*>(ptr)->size()-1) { static_cast<QList<qreal>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextOption___setTabArray_tabStops_setList(void* ptr, double i)
{
	static_cast<QList<qreal>*>(ptr)->append(i);
}

void* QTextOption___setTabArray_tabStops_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qreal>();
}

double QTextOption___tabArray_atList(void* ptr, int i)
{
	return ({qreal tmp = static_cast<QList<qreal>*>(ptr)->at(i); if (i == static_cast<QList<qreal>*>(ptr)->size()-1) { static_cast<QList<qreal>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QTextOption___tabArray_setList(void* ptr, double i)
{
	static_cast<QList<qreal>*>(ptr)->append(i);
}

void* QTextOption___tabArray_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qreal>();
}

class MyQTextTable: public QTextTable
{
public:
	void childEvent(QChildEvent * event) { callbackQTextObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQTextObject_CustomEvent(this, event); };
	void deleteLater() { callbackQTextObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQTextObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQTextObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQTextObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQTextObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QTextTable*)
Q_DECLARE_METATYPE(MyQTextTable*)

int QTextTable_QTextTable_QRegisterMetaType(){qRegisterMetaType<QTextTable*>(); return qRegisterMetaType<MyQTextTable*>();}

void QTextTable_AppendColumns(void* ptr, int count)
{
	static_cast<QTextTable*>(ptr)->appendColumns(count);
}

void QTextTable_AppendRows(void* ptr, int count)
{
	static_cast<QTextTable*>(ptr)->appendRows(count);
}

void* QTextTable_CellAt(void* ptr, int row, int column)
{
	return new QTextTableCell(static_cast<QTextTable*>(ptr)->cellAt(row, column));
}

void* QTextTable_CellAt2(void* ptr, int position)
{
	return new QTextTableCell(static_cast<QTextTable*>(ptr)->cellAt(position));
}

void* QTextTable_CellAt3(void* ptr, void* cursor)
{
	return new QTextTableCell(static_cast<QTextTable*>(ptr)->cellAt(*static_cast<QTextCursor*>(cursor)));
}

int QTextTable_Columns(void* ptr)
{
	return static_cast<QTextTable*>(ptr)->columns();
}

void* QTextTable_Format(void* ptr)
{
	return new QTextTableFormat(static_cast<QTextTable*>(ptr)->format());
}

void QTextTable_InsertColumns(void* ptr, int index, int columns)
{
	static_cast<QTextTable*>(ptr)->insertColumns(index, columns);
}

void QTextTable_InsertRows(void* ptr, int index, int rows)
{
	static_cast<QTextTable*>(ptr)->insertRows(index, rows);
}

void QTextTable_RemoveColumns(void* ptr, int index, int columns)
{
	static_cast<QTextTable*>(ptr)->removeColumns(index, columns);
}

void QTextTable_RemoveRows(void* ptr, int index, int rows)
{
	static_cast<QTextTable*>(ptr)->removeRows(index, rows);
}

void QTextTable_Resize(void* ptr, int rows, int columns)
{
	static_cast<QTextTable*>(ptr)->resize(rows, columns);
}

int QTextTable_Rows(void* ptr)
{
	return static_cast<QTextTable*>(ptr)->rows();
}

Q_DECLARE_METATYPE(QTextTableCell)
Q_DECLARE_METATYPE(QTextTableCell*)
void* QTextTableCell_NewQTextTableCell()
{
	return new QTextTableCell();
}

void* QTextTableCell_NewQTextTableCell2(void* other)
{
	return new QTextTableCell(*static_cast<QTextTableCell*>(other));
}

int QTextTableCell_Column(void* ptr)
{
	return static_cast<QTextTableCell*>(ptr)->column();
}

void* QTextTableCell_FirstCursorPosition(void* ptr)
{
	return new QTextCursor(static_cast<QTextTableCell*>(ptr)->firstCursorPosition());
}

void* QTextTableCell_Format(void* ptr)
{
	return new QTextCharFormat(static_cast<QTextTableCell*>(ptr)->format());
}

char QTextTableCell_IsValid(void* ptr)
{
	return static_cast<QTextTableCell*>(ptr)->isValid();
}

int QTextTableCell_Row(void* ptr)
{
	return static_cast<QTextTableCell*>(ptr)->row();
}

void QTextTableCell_DestroyQTextTableCell(void* ptr)
{
	static_cast<QTextTableCell*>(ptr)->~QTextTableCell();
}

Q_DECLARE_METATYPE(QTextTableFormat)
Q_DECLARE_METATYPE(QTextTableFormat*)
void* QTextTableFormat_NewQTextTableFormat()
{
	return new QTextTableFormat();
}

long long QTextTableFormat_Alignment(void* ptr)
{
	return static_cast<QTextTableFormat*>(ptr)->alignment();
}

int QTextTableFormat_Columns(void* ptr)
{
	return static_cast<QTextTableFormat*>(ptr)->columns();
}

void QTextTableFormat_SetAlignment(void* ptr, long long alignment)
{
	static_cast<QTextTableFormat*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void* QTextTableFormat___columnWidthConstraints_atList(void* ptr, int i)
{
	return new QTextLength(({QTextLength tmp = static_cast<QVector<QTextLength>*>(ptr)->at(i); if (i == static_cast<QVector<QTextLength>*>(ptr)->size()-1) { static_cast<QVector<QTextLength>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QTextTableFormat___columnWidthConstraints_setList(void* ptr, void* i)
{
	static_cast<QVector<QTextLength>*>(ptr)->append(*static_cast<QTextLength*>(i));
}

void* QTextTableFormat___columnWidthConstraints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QTextLength>();
}

void* QTextTableFormat___setColumnWidthConstraints_constraints_atList(void* ptr, int i)
{
	return new QTextLength(({QTextLength tmp = static_cast<QVector<QTextLength>*>(ptr)->at(i); if (i == static_cast<QVector<QTextLength>*>(ptr)->size()-1) { static_cast<QVector<QTextLength>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QTextTableFormat___setColumnWidthConstraints_constraints_setList(void* ptr, void* i)
{
	static_cast<QVector<QTextLength>*>(ptr)->append(*static_cast<QTextLength*>(i));
}

void* QTextTableFormat___setColumnWidthConstraints_constraints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QTextLength>();
}

Q_DECLARE_METATYPE(QTransform)
Q_DECLARE_METATYPE(QTransform*)
void* QTransform_NewQTransform2()
{
	return new QTransform();
}

void* QTransform_NewQTransform3(double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33)
{
	return new QTransform(m11, m12, m13, m21, m22, m23, m31, m32, m33);
}

void* QTransform_NewQTransform4(double m11, double m12, double m21, double m22, double dx, double dy)
{
	return new QTransform(m11, m12, m21, m22, dx, dy);
}

void QTransform_Map(void* ptr, double x, double y, double tx, double ty)
{
	static_cast<QTransform*>(ptr)->map(x, y, &tx, &ty);
}

void* QTransform_Map2(void* ptr, void* point)
{
	return ({ QPoint tmpValue = static_cast<QTransform*>(ptr)->map(*static_cast<QPoint*>(point)); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QTransform_Map3(void* ptr, void* p)
{
	return ({ QPointF tmpValue = static_cast<QTransform*>(ptr)->map(*static_cast<QPointF*>(p)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QTransform_Map4(void* ptr, void* l)
{
	return ({ QLine tmpValue = static_cast<QTransform*>(ptr)->map(*static_cast<QLine*>(l)); new QLine(tmpValue.p1(), tmpValue.p2()); });
}

void* QTransform_Map5(void* ptr, void* line)
{
	return ({ QLineF tmpValue = static_cast<QTransform*>(ptr)->map(*static_cast<QLineF*>(line)); new QLineF(tmpValue.p1(), tmpValue.p2()); });
}

void* QTransform_Map6(void* ptr, void* polygon)
{
	return new QPolygonF(static_cast<QTransform*>(ptr)->map(*static_cast<QPolygonF*>(polygon)));
}

void* QTransform_Map7(void* ptr, void* polygon)
{
	return new QPolygon(static_cast<QTransform*>(ptr)->map(*static_cast<QPolygon*>(polygon)));
}

void* QTransform_Map8(void* ptr, void* region)
{
	return new QRegion(static_cast<QTransform*>(ptr)->map(*static_cast<QRegion*>(region)));
}

void* QTransform_Map9(void* ptr, void* path)
{
	return new QPainterPath(static_cast<QTransform*>(ptr)->map(*static_cast<QPainterPath*>(path)));
}

void QTransform_Map10(void* ptr, int x, int y, int tx, int ty)
{
	static_cast<QTransform*>(ptr)->map(x, y, &tx, &ty);
}

void* QTransform_Scale(void* ptr, double sx, double sy)
{
	return new QTransform(static_cast<QTransform*>(ptr)->scale(sx, sy));
}

long long QTransform_Type(void* ptr)
{
	return static_cast<QTransform*>(ptr)->type();
}

class MyQValidator: public QValidator
{
public:
	MyQValidator(QObject *parent = Q_NULLPTR) : QValidator(parent) {QValidator_QValidator_QRegisterMetaType();};
	void Signal_Changed() { callbackQValidator_Changed(this); };
	QValidator::State validate(QString & input, int & pos) const { QByteArray* t140f86 = new QByteArray(input.toUtf8()); QtGui_PackedString inputPacked = { const_cast<char*>(t140f86->prepend("WHITESPACE").constData()+10), t140f86->size()-10, t140f86 };return static_cast<QValidator::State>(callbackQValidator_Validate(const_cast<void*>(static_cast<const void*>(this)), inputPacked, pos)); };
	 ~MyQValidator() { callbackQValidator_DestroyQValidator(this); };
	void childEvent(QChildEvent * event) { callbackQValidator_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQValidator_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQValidator_CustomEvent(this, event); };
	void deleteLater() { callbackQValidator_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQValidator_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQValidator_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQValidator_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQValidator_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQValidator_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQValidator_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QValidator*)
Q_DECLARE_METATYPE(MyQValidator*)

int QValidator_QValidator_QRegisterMetaType(){qRegisterMetaType<QValidator*>(); return qRegisterMetaType<MyQValidator*>();}

void* QValidator_NewQValidator(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQValidator(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQValidator(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQValidator(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQValidator(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQValidator(static_cast<QWindow*>(parent));
	} else {
		return new MyQValidator(static_cast<QObject*>(parent));
	}
}

void QValidator_ConnectChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QValidator*>(ptr), static_cast<void (QValidator::*)()>(&QValidator::changed), static_cast<MyQValidator*>(ptr), static_cast<void (MyQValidator::*)()>(&MyQValidator::Signal_Changed), static_cast<Qt::ConnectionType>(t));
}

void QValidator_DisconnectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValidator*>(ptr), static_cast<void (QValidator::*)()>(&QValidator::changed), static_cast<MyQValidator*>(ptr), static_cast<void (MyQValidator::*)()>(&MyQValidator::Signal_Changed));
}

void QValidator_Changed(void* ptr)
{
	static_cast<QValidator*>(ptr)->changed();
}

long long QValidator_Validate(void* ptr, struct QtGui_PackedString input, int pos)
{
	return static_cast<QValidator*>(ptr)->validate(*(new QString(QString::fromUtf8(input.data, input.len))), pos);
}

void QValidator_DestroyQValidator(void* ptr)
{
	static_cast<QValidator*>(ptr)->~QValidator();
}

void QValidator_DestroyQValidatorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QValidator___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QValidator___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QValidator___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QValidator___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QValidator___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QValidator___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QValidator___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QValidator___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QValidator___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QValidator___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QValidator___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QValidator___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QValidator_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QIntValidator*>(static_cast<QObject*>(ptr))) {
		static_cast<QIntValidator*>(ptr)->QIntValidator::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QValidator*>(ptr)->QValidator::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QValidator_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QIntValidator*>(static_cast<QObject*>(ptr))) {
		static_cast<QIntValidator*>(ptr)->QIntValidator::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QValidator*>(ptr)->QValidator::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QValidator_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QIntValidator*>(static_cast<QObject*>(ptr))) {
		static_cast<QIntValidator*>(ptr)->QIntValidator::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QValidator*>(ptr)->QValidator::customEvent(static_cast<QEvent*>(event));
	}
}

void QValidator_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QIntValidator*>(static_cast<QObject*>(ptr))) {
		static_cast<QIntValidator*>(ptr)->QIntValidator::deleteLater();
	} else {
		static_cast<QValidator*>(ptr)->QValidator::deleteLater();
	}
}

void QValidator_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QIntValidator*>(static_cast<QObject*>(ptr))) {
		static_cast<QIntValidator*>(ptr)->QIntValidator::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QValidator*>(ptr)->QValidator::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QValidator_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QIntValidator*>(static_cast<QObject*>(ptr))) {
		return static_cast<QIntValidator*>(ptr)->QIntValidator::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QValidator*>(ptr)->QValidator::event(static_cast<QEvent*>(e));
	}
}

char QValidator_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QIntValidator*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
			return static_cast<QIntValidator*>(ptr)->QIntValidator::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QIntValidator*>(ptr)->QIntValidator::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
			return static_cast<QValidator*>(ptr)->QValidator::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QValidator*>(ptr)->QValidator::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void QValidator_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QIntValidator*>(static_cast<QObject*>(ptr))) {
		static_cast<QIntValidator*>(ptr)->QIntValidator::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QValidator*>(ptr)->QValidator::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

Q_DECLARE_METATYPE(QVector2D)
Q_DECLARE_METATYPE(QVector2D*)
void* QVector2D_NewQVector2D()
{
	return new QVector2D();
}

void* QVector2D_NewQVector2D3(float xpos, float ypos)
{
	return new QVector2D(xpos, ypos);
}

void* QVector2D_NewQVector2D4(void* point)
{
	return new QVector2D(*static_cast<QPoint*>(point));
}

void* QVector2D_NewQVector2D5(void* point)
{
	return new QVector2D(*static_cast<QPointF*>(point));
}

void* QVector2D_NewQVector2D6(void* vector)
{
	return new QVector2D(*static_cast<QVector3D*>(vector));
}

void* QVector2D_NewQVector2D7(void* vector)
{
	return new QVector2D(*static_cast<QVector4D*>(vector));
}

float QVector2D_Length(void* ptr)
{
	return static_cast<QVector2D*>(ptr)->length();
}

float QVector2D_X(void* ptr)
{
	return static_cast<QVector2D*>(ptr)->x();
}

float QVector2D_Y(void* ptr)
{
	return static_cast<QVector2D*>(ptr)->y();
}

Q_DECLARE_METATYPE(QVector3D)
Q_DECLARE_METATYPE(QVector3D*)
void* QVector3D_NewQVector3D()
{
	return new QVector3D();
}

void* QVector3D_NewQVector3D3(float xpos, float ypos, float zpos)
{
	return new QVector3D(xpos, ypos, zpos);
}

void* QVector3D_NewQVector3D4(void* point)
{
	return new QVector3D(*static_cast<QPoint*>(point));
}

void* QVector3D_NewQVector3D5(void* point)
{
	return new QVector3D(*static_cast<QPointF*>(point));
}

void* QVector3D_NewQVector3D6(void* vector)
{
	return new QVector3D(*static_cast<QVector2D*>(vector));
}

void* QVector3D_NewQVector3D7(void* vector, float zpos)
{
	return new QVector3D(*static_cast<QVector2D*>(vector), zpos);
}

void* QVector3D_NewQVector3D8(void* vector)
{
	return new QVector3D(*static_cast<QVector4D*>(vector));
}

float QVector3D_Length(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->length();
}

void* QVector3D_QVector3D_Normal(void* v1, void* v2)
{
	return new QVector3D(QVector3D::normal(*static_cast<QVector3D*>(v1), *static_cast<QVector3D*>(v2)));
}

void* QVector3D_QVector3D_Normal2(void* v1, void* v2, void* v3)
{
	return new QVector3D(QVector3D::normal(*static_cast<QVector3D*>(v1), *static_cast<QVector3D*>(v2), *static_cast<QVector3D*>(v3)));
}

float QVector3D_X(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->x();
}

float QVector3D_Y(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->y();
}

float QVector3D_Z(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->z();
}

Q_DECLARE_METATYPE(QVector4D)
Q_DECLARE_METATYPE(QVector4D*)
void* QVector4D_NewQVector4D()
{
	return new QVector4D();
}

void* QVector4D_NewQVector4D3(float xpos, float ypos, float zpos, float wpos)
{
	return new QVector4D(xpos, ypos, zpos, wpos);
}

void* QVector4D_NewQVector4D4(void* point)
{
	return new QVector4D(*static_cast<QPoint*>(point));
}

void* QVector4D_NewQVector4D5(void* point)
{
	return new QVector4D(*static_cast<QPointF*>(point));
}

void* QVector4D_NewQVector4D6(void* vector)
{
	return new QVector4D(*static_cast<QVector2D*>(vector));
}

void* QVector4D_NewQVector4D7(void* vector, float zpos, float wpos)
{
	return new QVector4D(*static_cast<QVector2D*>(vector), zpos, wpos);
}

void* QVector4D_NewQVector4D8(void* vector)
{
	return new QVector4D(*static_cast<QVector3D*>(vector));
}

void* QVector4D_NewQVector4D9(void* vector, float wpos)
{
	return new QVector4D(*static_cast<QVector3D*>(vector), wpos);
}

float QVector4D_Length(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->length();
}

void QVector4D_SetW(void* ptr, float w)
{
	static_cast<QVector4D*>(ptr)->setW(w);
}

float QVector4D_W(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->w();
}

float QVector4D_X(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->x();
}

float QVector4D_Y(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->y();
}

float QVector4D_Z(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->z();
}

class MyQWindow: public QWindow
{
public:
	MyQWindow(QScreen *targetScreen = Q_NULLPTR) : QWindow(targetScreen) {QWindow_QWindow_QRegisterMetaType();};
	MyQWindow(QWindow *parent) : QWindow(parent) {QWindow_QWindow_QRegisterMetaType();};
	bool close() { return callbackQWindow_Close(this) != 0; };
	bool event(QEvent * ev) { return callbackQWindow_Event(this, ev) != 0; };
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQWindow_Format(const_cast<void*>(static_cast<const void*>(this)))); };
	void hide() { callbackQWindow_Hide(this); };
	void hideEvent(QHideEvent * ev) { callbackQWindow_HideEvent(this, ev); };
	void keyPressEvent(QKeyEvent * ev) { callbackQWindow_KeyPressEvent(this, ev); };
	void keyReleaseEvent(QKeyEvent * ev) { callbackQWindow_KeyReleaseEvent(this, ev); };
	void lower() { callbackQWindow_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * ev) { callbackQWindow_MouseDoubleClickEvent(this, ev); };
	void mousePressEvent(QMouseEvent * ev) { callbackQWindow_MousePressEvent(this, ev); };
	void mouseReleaseEvent(QMouseEvent * ev) { callbackQWindow_MouseReleaseEvent(this, ev); };
	void resizeEvent(QResizeEvent * ev) { callbackQWindow_ResizeEvent(this, ev); };
	void setGeometry(int posx, int posy, int w, int h) { callbackQWindow_SetGeometry(this, posx, posy, w, h); };
	void setGeometry(const QRect & rect) { callbackQWindow_SetGeometry2(this, const_cast<QRect*>(&rect)); };
	void setMaximumWidth(int w) { callbackQWindow_SetMaximumWidth(this, w); };
	void setMinimumWidth(int w) { callbackQWindow_SetMinimumWidth(this, w); };
	void show() { callbackQWindow_Show(this); };
	void showEvent(QShowEvent * ev) { callbackQWindow_ShowEvent(this, ev); };
	void showMaximized() { callbackQWindow_ShowMaximized(this); };
	QSize size() const { return *static_cast<QSize*>(callbackQWindow_Size(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQWindow() { callbackQWindow_DestroyQWindow(this); };
	void childEvent(QChildEvent * event) { callbackQWindow_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWindow_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWindow_CustomEvent(this, event); };
	void deleteLater() { callbackQWindow_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWindow_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWindow_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWindow_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWindow_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWindow_TimerEvent(this, event); };
	QSurface::SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQWindow_SurfaceType(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QWindow*)
Q_DECLARE_METATYPE(MyQWindow*)

int QWindow_QWindow_QRegisterMetaType(){qRegisterMetaType<QWindow*>(); return qRegisterMetaType<MyQWindow*>();}

void* QWindow_NewQWindow(void* targetScreen)
{
	return new MyQWindow(static_cast<QScreen*>(targetScreen));
}

void* QWindow_NewQWindow2(void* parent)
{
	return new MyQWindow(static_cast<QWindow*>(parent));
}

char QWindow_Close(void* ptr)
{
		bool returnArg;
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QWindow_CloseDefault(void* ptr)
{
		return static_cast<QWindow*>(ptr)->QWindow::close();
}

void QWindow_Create(void* ptr)
{
		static_cast<QWindow*>(ptr)->create();
}

void* QWindow_Cursor(void* ptr)
{
		return new QCursor(static_cast<QWindow*>(ptr)->cursor());
}

void QWindow_Destroy(void* ptr)
{
		static_cast<QWindow*>(ptr)->destroy();
}

char QWindow_Event(void* ptr, void* ev)
{
		return static_cast<QWindow*>(ptr)->event(static_cast<QEvent*>(ev));
}

char QWindow_EventDefault(void* ptr, void* ev)
{
		return static_cast<QWindow*>(ptr)->QWindow::event(static_cast<QEvent*>(ev));
}

void* QWindow_Format(void* ptr)
{
		return new QSurfaceFormat(static_cast<QWindow*>(ptr)->format());
}

void* QWindow_FormatDefault(void* ptr)
{
		return new QSurfaceFormat(static_cast<QWindow*>(ptr)->QWindow::format());
}

void* QWindow_Geometry(void* ptr)
{
		return ({ QRect tmpValue = static_cast<QWindow*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

int QWindow_Height(void* ptr)
{
		return static_cast<QWindow*>(ptr)->height();
}

void QWindow_Hide(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "hide");
}

void QWindow_HideDefault(void* ptr)
{
		static_cast<QWindow*>(ptr)->QWindow::hide();
}

void QWindow_HideEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->hideEvent(static_cast<QHideEvent*>(ev));
}

void QWindow_HideEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::hideEvent(static_cast<QHideEvent*>(ev));
}

void* QWindow_Icon(void* ptr)
{
		return new QIcon(static_cast<QWindow*>(ptr)->icon());
}

void QWindow_KeyPressEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_KeyPressEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_KeyReleaseEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_KeyReleaseEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_Lower(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "lower");
}

void QWindow_LowerDefault(void* ptr)
{
		static_cast<QWindow*>(ptr)->QWindow::lower();
}

void* QWindow_MaximumSize(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWindow*>(ptr)->maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QWindow_MaximumWidth(void* ptr)
{
		return static_cast<QWindow*>(ptr)->maximumWidth();
}

int QWindow_MinimumHeight(void* ptr)
{
		return static_cast<QWindow*>(ptr)->minimumHeight();
}

void* QWindow_MinimumSize(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWindow*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QWindow_MinimumWidth(void* ptr)
{
		return static_cast<QWindow*>(ptr)->minimumWidth();
}

void QWindow_MouseDoubleClickEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseDoubleClickEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MousePressEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MousePressEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseReleaseEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseReleaseEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void* QWindow_Parent(void* ptr, long long mode)
{
		return static_cast<QWindow*>(ptr)->parent(static_cast<QWindow::AncestorMode>(mode));
}

void* QWindow_Parent2(void* ptr)
{
		return static_cast<QWindow*>(ptr)->parent();
}

void* QWindow_Position(void* ptr)
{
		return ({ QPoint tmpValue = static_cast<QWindow*>(ptr)->position(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QWindow_Resize(void* ptr, void* newSize)
{
		static_cast<QWindow*>(ptr)->resize(*static_cast<QSize*>(newSize));
}

void QWindow_Resize2(void* ptr, int w, int h)
{
		static_cast<QWindow*>(ptr)->resize(w, h);
}

void QWindow_ResizeEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QWindow_ResizeEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QWindow_SetGeometry(void* ptr, int posx, int posy, int w, int h)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setGeometry", Q_ARG(int, posx), Q_ARG(int, posy), Q_ARG(int, w), Q_ARG(int, h));
}

void QWindow_SetGeometryDefault(void* ptr, int posx, int posy, int w, int h)
{
		static_cast<QWindow*>(ptr)->QWindow::setGeometry(posx, posy, w, h);
}

void QWindow_SetGeometry2(void* ptr, void* rect)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setGeometry", Q_ARG(const QRect, *static_cast<QRect*>(rect)));
}

void QWindow_SetGeometry2Default(void* ptr, void* rect)
{
		static_cast<QWindow*>(ptr)->QWindow::setGeometry(*static_cast<QRect*>(rect));
}

void QWindow_SetIcon(void* ptr, void* icon)
{
		static_cast<QWindow*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QWindow_SetMaximumWidth(void* ptr, int w)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMaximumWidth", Q_ARG(int, w));
}

void QWindow_SetMaximumWidthDefault(void* ptr, int w)
{
		static_cast<QWindow*>(ptr)->QWindow::setMaximumWidth(w);
}

void QWindow_SetMinimumSize(void* ptr, void* size)
{
		static_cast<QWindow*>(ptr)->setMinimumSize(*static_cast<QSize*>(size));
}

void QWindow_SetMinimumWidth(void* ptr, int w)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMinimumWidth", Q_ARG(int, w));
}

void QWindow_SetMinimumWidthDefault(void* ptr, int w)
{
		static_cast<QWindow*>(ptr)->QWindow::setMinimumWidth(w);
}

void QWindow_SetPosition(void* ptr, void* pt)
{
		static_cast<QWindow*>(ptr)->setPosition(*static_cast<QPoint*>(pt));
}

void QWindow_SetPosition2(void* ptr, int posx, int posy)
{
		static_cast<QWindow*>(ptr)->setPosition(posx, posy);
}

void QWindow_Show(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "show");
}

void QWindow_ShowDefault(void* ptr)
{
		static_cast<QWindow*>(ptr)->QWindow::show();
}

void QWindow_ShowEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->showEvent(static_cast<QShowEvent*>(ev));
}

void QWindow_ShowEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::showEvent(static_cast<QShowEvent*>(ev));
}

void QWindow_ShowMaximized(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showMaximized");
}

void QWindow_ShowMaximizedDefault(void* ptr)
{
		static_cast<QWindow*>(ptr)->QWindow::showMaximized();
}

void* QWindow_Size(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWindow*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QWindow_SizeDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWindow*>(ptr)->QWindow::size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtGui_PackedString QWindow_Title(void* ptr)
{
		return ({ QByteArray* t3f590b = new QByteArray(static_cast<QWindow*>(ptr)->title().toUtf8()); QtGui_PackedString { const_cast<char*>(t3f590b->prepend("WHITESPACE").constData()+10), t3f590b->size()-10, t3f590b }; });
}

long long QWindow_Type(void* ptr)
{
		return static_cast<QWindow*>(ptr)->type();
}

int QWindow_Width(void* ptr)
{
		return static_cast<QWindow*>(ptr)->width();
}

int QWindow_X(void* ptr)
{
		return static_cast<QWindow*>(ptr)->x();
}

int QWindow_Y(void* ptr)
{
		return static_cast<QWindow*>(ptr)->y();
}

void QWindow_DestroyQWindow(void* ptr)
{
	static_cast<QWindow*>(ptr)->~QWindow();
}

void QWindow_DestroyQWindowDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWindow___children_atList(void* ptr, int i)
{
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWindow___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWindow___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>();
}

void* QWindow___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWindow___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWindow___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>();
}

void* QWindow___findChildren_atList(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWindow___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWindow___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QWindow___findChildren_atList3(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWindow___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWindow___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void QWindow_ChildEvent(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWindow_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->QWindow::childEvent(static_cast<QChildEvent*>(event));
}

void QWindow_ConnectNotify(void* ptr, void* sign)
{
		static_cast<QWindow*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWindow_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWindow*>(ptr)->QWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWindow_CustomEvent(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWindow_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->QWindow::customEvent(static_cast<QEvent*>(event));
}

void QWindow_DeleteLater(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "deleteLater");
}

void QWindow_DeleteLaterDefault(void* ptr)
{
		static_cast<QWindow*>(ptr)->QWindow::deleteLater();
}

void QWindow_DisconnectNotify(void* ptr, void* sign)
{
		static_cast<QWindow*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWindow_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWindow*>(ptr)->QWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QWindow_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QWindow*>(ptr)->eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QWindow*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QWindow_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QWindow*>(ptr)->QWindow::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QWindow*>(ptr)->QWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QWindow_TimerEvent(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWindow_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->QWindow::timerEvent(static_cast<QTimerEvent*>(event));
}

long long QWindow_SurfaceType(void* ptr)
{
		return static_cast<QWindow*>(ptr)->surfaceType();
}

long long QWindow_SurfaceTypeDefault(void* ptr)
{
		return static_cast<QWindow*>(ptr)->QWindow::surfaceType();
}

