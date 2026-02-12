// +build minimal

#define protected public
#define private public

#include "core-minimal.h"
#include "_cgo_export.h"

#ifndef QT_CORE_LIB
	#error ------------------------------------------------------------------
	#error please run: '$(go env GOPATH)/bin/qtsetup'
	#error more info here: https://github.com/therecipe/qt/wiki/Installation
	#error ------------------------------------------------------------------
#endif
#include <QAbstractItemModel>
#include <QBitArray>
#include <QBuffer>
#include <QByteArray>
#include <QCalendar>
#include <QCborValue>
#include <QChar>
#include <QChildEvent>
#include <QCoreApplication>
#include <QDataStream>
#include <QDate>
#include <QDateTime>
#include <QDir>
#include <QEasingCurve>
#include <QEvent>
#include <QFile>
#include <QFileDevice>
#include <QFileInfo>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QIODevice>
#include <QImage>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QItemSelectionRange>
#include <QJsonArray>
#include <QJsonDocument>
#include <QJsonObject>
#include <QJsonValue>
#include <QLatin1Char>
#include <QLatin1String>
#include <QLayout>
#include <QLine>
#include <QLineF>
#include <QLocale>
#include <QMap>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMetaProperty>
#include <QModelIndex>
#include <QObject>
#include <QPersistentModelIndex>
#include <QPoint>
#include <QPointF>
#include <QRect>
#include <QRectF>
#include <QRegExp>
#include <QRegularExpression>
#include <QRegularExpressionMatch>
#include <QSize>
#include <QSizeF>
#include <QString>
#include <QStringRef>
#include <QSysInfo>
#include <QTextStream>
#include <QTime>
#include <QTimeZone>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QUuid>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>
#include <QTextDocument>
#include <QObject>
#include <QStringList>

class MyQAbstractItemModel: public QAbstractItemModel
{
public:
	MyQAbstractItemModel(QObject *parent = Q_NULLPTR) : QAbstractItemModel(parent) {QAbstractItemModel_QAbstractItemModel_QRegisterMetaType();};
	int columnCount(const QModelIndex & parent) const { return callbackQAbstractItemModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQAbstractItemModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQAbstractItemModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQAbstractItemModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	int rowCount(const QModelIndex & parent) const { return callbackQAbstractItemModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQAbstractItemModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	 ~MyQAbstractItemModel() { callbackQAbstractItemModel_DestroyQAbstractItemModel(this); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QAbstractItemModel*)
Q_DECLARE_METATYPE(MyQAbstractItemModel*)

int QAbstractItemModel_QAbstractItemModel_QRegisterMetaType(){qRegisterMetaType<QAbstractItemModel*>(); return qRegisterMetaType<MyQAbstractItemModel*>();}

void* QAbstractItemModel_NewQAbstractItemModel(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QWindow*>(parent));
	} else {
		return new MyQAbstractItemModel(static_cast<QObject*>(parent));
	}
}

int QAbstractItemModel_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void* QAbstractItemModel_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QAbstractItemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

char QAbstractItemModel_HasChildren(void* ptr, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_HasChildrenDefault(void* ptr, void* parent)
{
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QAbstractItemModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

char QAbstractItemModel_InsertColumn(void* ptr, int column, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->insertColumn(column, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_InsertColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_InsertRow(void* ptr, int row, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->insertRow(row, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_InsertRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct QtCore_PackedList QAbstractItemModel_Match(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValuef30906 = new QList<QModelIndex>(static_cast<QAbstractItemModel*>(ptr)->match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtCore_PackedList { tmpValuef30906, tmpValuef30906->size() }; });
}

struct QtCore_PackedList QAbstractItemModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
		return ({ QList<QModelIndex>* tmpValue4d7ea4 = new QList<QModelIndex>(static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtCore_PackedList { tmpValue4d7ea4, tmpValue4d7ea4->size() }; });
}

void* QAbstractItemModel_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

char QAbstractItemModel_RemoveColumn(void* ptr, int column, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->removeColumn(column, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_RemoveColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_RemoveRow(void* ptr, int row, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->removeRow(row, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_RemoveRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_SetData(void* ptr, void* index, void* value, int role)
{
	return static_cast<QAbstractItemModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char QAbstractItemModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QAbstractItemModel_DestroyQAbstractItemModel(void* ptr)
{
	static_cast<QAbstractItemModel*>(ptr)->~QAbstractItemModel();
}

void QAbstractItemModel_DestroyQAbstractItemModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAbstractItemModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int QAbstractItemModel___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QAbstractItemModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QAbstractItemModel___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QAbstractItemModel___doSetRoleNames_roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___doSetRoleNames_roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QAbstractItemModel___doSetRoleNames_roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtCore_PackedList QAbstractItemModel___doSetRoleNames_roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtCore_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

void* QAbstractItemModel___encodeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___encodeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___encodeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModel___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QAbstractItemModel___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtCore_PackedList QAbstractItemModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* QAbstractItemModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QAbstractItemModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QAbstractItemModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QAbstractItemModel___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QAbstractItemModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModel___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QAbstractItemModel___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtCore_PackedList QAbstractItemModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtCore_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

void* QAbstractItemModel___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QAbstractItemModel___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtCore_PackedList QAbstractItemModel___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* QAbstractItemModel___setRoleNames_roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___setRoleNames_roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QAbstractItemModel___setRoleNames_roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtCore_PackedList QAbstractItemModel___setRoleNames_roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtCore_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

int QAbstractItemModel_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModel_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModel_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModel_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModel_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModel_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModel_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModel_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModel_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModel_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModel_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModel_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

Q_DECLARE_METATYPE(QBitArray)
Q_DECLARE_METATYPE(QBitArray*)
void* QBitArray_NewQBitArray()
{
	return new QBitArray();
}

void* QBitArray_NewQBitArray2(int size, char value)
{
	return new QBitArray(size, value != 0);
}

void* QBitArray_NewQBitArray3(void* other)
{
	return new QBitArray(*static_cast<QBitArray*>(other));
}

void* QBitArray_NewQBitArray4(void* other)
{
	return new QBitArray(*static_cast<QBitArray*>(other));
}

char QBitArray_At(void* ptr, int i)
{
	return static_cast<QBitArray*>(ptr)->at(i);
}

void QBitArray_Clear(void* ptr)
{
	static_cast<QBitArray*>(ptr)->clear();
}

int QBitArray_Count(void* ptr)
{
	return static_cast<QBitArray*>(ptr)->count();
}

int QBitArray_Count2(void* ptr, char on)
{
	return static_cast<QBitArray*>(ptr)->count(on != 0);
}

char QBitArray_IsEmpty(void* ptr)
{
	return static_cast<QBitArray*>(ptr)->isEmpty();
}

void QBitArray_Resize(void* ptr, int size)
{
	static_cast<QBitArray*>(ptr)->resize(size);
}

int QBitArray_Size(void* ptr)
{
	return static_cast<QBitArray*>(ptr)->size();
}

class MyQBuffer: public QBuffer
{
public:
	MyQBuffer(QObject *parent = Q_NULLPTR) : QBuffer(parent) {QBuffer_QBuffer_QRegisterMetaType();};
	MyQBuffer(QByteArray *byteArray, QObject *parent = Q_NULLPTR) : QBuffer(byteArray, parent) {QBuffer_QBuffer_QRegisterMetaType();};
	void close() { callbackQIODevice_Close(this); };
	bool open(QIODevice::OpenMode flags) { return callbackQIODevice_Open(this, flags) != 0; };
	qint64 pos() const { return callbackQIODevice_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 size() const { return callbackQIODevice_Size(const_cast<void*>(static_cast<const void*>(this))); };
	 ~MyQBuffer() { callbackQBuffer_DestroyQBuffer(this); };
	qint64 readData(char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { data, maxSize, NULL };return callbackQBuffer_ReadData(this, dataPacked, maxSize); };
	qint64 writeData(const char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { const_cast<char*>(data), maxSize, NULL };return callbackQBuffer_WriteData(this, dataPacked, maxSize); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QBuffer*)
Q_DECLARE_METATYPE(MyQBuffer*)

int QBuffer_QBuffer_QRegisterMetaType(){qRegisterMetaType<QBuffer*>(); return qRegisterMetaType<MyQBuffer*>();}

void* QBuffer_NewQBuffer(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBuffer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBuffer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBuffer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBuffer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBuffer(static_cast<QWindow*>(parent));
	} else {
		return new MyQBuffer(static_cast<QObject*>(parent));
	}
}

void* QBuffer_NewQBuffer2(void* byteArray, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBuffer(static_cast<QByteArray*>(byteArray), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBuffer(static_cast<QByteArray*>(byteArray), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBuffer(static_cast<QByteArray*>(byteArray), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBuffer(static_cast<QByteArray*>(byteArray), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBuffer(static_cast<QByteArray*>(byteArray), static_cast<QWindow*>(parent));
	} else {
		return new MyQBuffer(static_cast<QByteArray*>(byteArray), static_cast<QObject*>(parent));
	}
}

void* QBuffer_Buffer(void* ptr)
{
	return new QByteArray(static_cast<QBuffer*>(ptr)->buffer());
}

void* QBuffer_Buffer2(void* ptr)
{
	return const_cast<QByteArray*>(&static_cast<QBuffer*>(ptr)->buffer());
}

void* QBuffer_Data(void* ptr)
{
	return const_cast<QByteArray*>(&static_cast<QBuffer*>(ptr)->data());
}

void QBuffer_SetData(void* ptr, void* data)
{
	static_cast<QBuffer*>(ptr)->setData(*static_cast<QByteArray*>(data));
}

void QBuffer_SetData2(void* ptr, char* data, int size)
{
	static_cast<QBuffer*>(ptr)->setData(const_cast<const char*>(data), size);
}

void QBuffer_DestroyQBuffer(void* ptr)
{
	static_cast<QBuffer*>(ptr)->~QBuffer();
}

void QBuffer_DestroyQBufferDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QBuffer_ReadData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QBuffer*>(ptr)->readData(data, maxSize);
}

long long QBuffer_ReadDataDefault(void* ptr, char* data, long long maxSize)
{
		return static_cast<QBuffer*>(ptr)->QBuffer::readData(data, maxSize);
}

long long QBuffer_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QBuffer*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

long long QBuffer_WriteDataDefault(void* ptr, char* data, long long maxSize)
{
		return static_cast<QBuffer*>(ptr)->QBuffer::writeData(const_cast<const char*>(data), maxSize);
}

Q_DECLARE_METATYPE(QByteArray)
Q_DECLARE_METATYPE(QByteArray*)
void* QByteArray_NewQByteArray()
{
	return new QByteArray();
}

void* QByteArray_NewQByteArray2(char* data, int size)
{
	return new QByteArray(const_cast<const char*>(data), size);
}

void* QByteArray_NewQByteArray3(int size, char* ch)
{
	return new QByteArray(size, *ch);
}

void* QByteArray_NewQByteArray4(void* other)
{
	return new QByteArray(*static_cast<QByteArray*>(other));
}

void* QByteArray_Append(void* ptr, void* ba)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(*static_cast<QByteArray*>(ba)));
}

void* QByteArray_Append2(void* ptr, char* ch)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(*ch));
}

void* QByteArray_Append3(void* ptr, int count, char* ch)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(count, *ch));
}

void* QByteArray_Append4(void* ptr, char* str)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(const_cast<const char*>(str)));
}

void* QByteArray_Append5(void* ptr, char* str, int l)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->append(const_cast<const char*>(str), l));
}

struct QtCore_PackedString QByteArray_At(void* ptr, int i)
{
	return ({ char t8d8fc5 = static_cast<QByteArray*>(ptr)->at(i); QtCore_PackedString { &t8d8fc5, -1, NULL }; });
}

struct QtCore_PackedString QByteArray_Back(void* ptr)
{
	return ({ char te8eea8 = static_cast<QByteArray*>(ptr)->back(); QtCore_PackedString { &te8eea8, -1, NULL }; });
}

void QByteArray_Clear(void* ptr)
{
	static_cast<QByteArray*>(ptr)->clear();
}

int QByteArray_Count(void* ptr, void* ba)
{
	return static_cast<QByteArray*>(ptr)->count(*static_cast<QByteArray*>(ba));
}

int QByteArray_Count2(void* ptr, char* ch)
{
	return static_cast<QByteArray*>(ptr)->count(*ch);
}

int QByteArray_Count3(void* ptr, char* str)
{
	return static_cast<QByteArray*>(ptr)->count(const_cast<const char*>(str));
}

int QByteArray_Count4(void* ptr)
{
	return static_cast<QByteArray*>(ptr)->count();
}

struct QtCore_PackedString QByteArray_Data(void* ptr)
{
	return QtCore_PackedString { static_cast<QByteArray*>(ptr)->data(), static_cast<QByteArray*>(ptr)->size(), NULL };
}

struct QtCore_PackedString QByteArray_Data2(void* ptr)
{
	return QtCore_PackedString { const_cast<char*>(static_cast<QByteArray*>(ptr)->data()), static_cast<QByteArray*>(ptr)->size(), NULL };
}

int QByteArray_IndexOf(void* ptr, void* ba, int from)
{
	return static_cast<QByteArray*>(ptr)->indexOf(*static_cast<QByteArray*>(ba), from);
}

int QByteArray_IndexOf2(void* ptr, char* ch, int from)
{
	return static_cast<QByteArray*>(ptr)->indexOf(*ch, from);
}

int QByteArray_IndexOf3(void* ptr, char* str, int from)
{
	return static_cast<QByteArray*>(ptr)->indexOf(const_cast<const char*>(str), from);
}

void* QByteArray_Insert(void* ptr, int i, void* ba)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, *static_cast<QByteArray*>(ba)));
}

void* QByteArray_Insert2(void* ptr, int i, char* ch)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, *ch));
}

void* QByteArray_Insert3(void* ptr, int i, int count, char* ch)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, count, *ch));
}

void* QByteArray_Insert4(void* ptr, int i, char* str)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, const_cast<const char*>(str)));
}

void* QByteArray_Insert5(void* ptr, int i, char* str, int l)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, const_cast<const char*>(str), l));
}

char QByteArray_IsEmpty(void* ptr)
{
	return static_cast<QByteArray*>(ptr)->isEmpty();
}

void* QByteArray_Left(void* ptr, int l)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->left(l));
}

int QByteArray_Length(void* ptr)
{
	return static_cast<QByteArray*>(ptr)->length();
}

void* QByteArray_Remove(void* ptr, int pos, int l)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->remove(pos, l));
}

void* QByteArray_Replace(void* ptr, int pos, int l, void* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(pos, l, *static_cast<QByteArray*>(after)));
}

void* QByteArray_Replace2(void* ptr, int pos, int l, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(pos, l, const_cast<const char*>(after)));
}

void* QByteArray_Replace3(void* ptr, int pos, int l, char* after, int alen)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(pos, l, const_cast<const char*>(after), alen));
}

void* QByteArray_Replace4(void* ptr, char* before, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*before, const_cast<const char*>(after)));
}

void* QByteArray_Replace5(void* ptr, char* before, void* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*before, *static_cast<QByteArray*>(after)));
}

void* QByteArray_Replace6(void* ptr, char* before, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(const_cast<const char*>(before), const_cast<const char*>(after)));
}

void* QByteArray_Replace7(void* ptr, char* before, int bsize, char* after, int asize)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(const_cast<const char*>(before), bsize, const_cast<const char*>(after), asize));
}

void* QByteArray_Replace8(void* ptr, void* before, void* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*static_cast<QByteArray*>(before), *static_cast<QByteArray*>(after)));
}

void* QByteArray_Replace9(void* ptr, void* before, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*static_cast<QByteArray*>(before), const_cast<const char*>(after)));
}

void* QByteArray_Replace10(void* ptr, char* before, void* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(const_cast<const char*>(before), *static_cast<QByteArray*>(after)));
}

void* QByteArray_Replace11(void* ptr, char* before, char* after)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->replace(*before, *after));
}

void QByteArray_Resize(void* ptr, int size)
{
	static_cast<QByteArray*>(ptr)->resize(size);
}

void* QByteArray_Right(void* ptr, int l)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->right(l));
}

int QByteArray_Size(void* ptr)
{
	return static_cast<QByteArray*>(ptr)->size();
}

struct QtCore_PackedList QByteArray_Split(void* ptr, char* sep)
{
	return ({ QList<QByteArray>* tmpValue17cac8 = new QList<QByteArray>(static_cast<QByteArray*>(ptr)->split(*sep)); QtCore_PackedList { tmpValue17cac8, tmpValue17cac8->size() }; });
}

int QByteArray_ToInt(void* ptr, char* ok, int base)
{
	return static_cast<QByteArray*>(ptr)->toInt(reinterpret_cast<bool*>(ok), base);
}

void* QByteArray_ToLower(void* ptr)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->toLower());
}

void QByteArray_DestroyQByteArray(void* ptr)
{
	static_cast<QByteArray*>(ptr)->~QByteArray();
}

void* QByteArray___split_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QByteArray___split_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QByteArray___split_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

Q_DECLARE_METATYPE(QCalendar)
Q_DECLARE_METATYPE(QCalendar*)
void* QCalendar_NewQCalendar()
{
	return new QCalendar();
}

void* QCalendar_NewQCalendar2(long long system)
{
	return new QCalendar(static_cast<QCalendar::System>(system));
}

void* QCalendar_NewQCalendar3(void* name)
{
	return new QCalendar(*static_cast<QLatin1String*>(name));
}

void* QCalendar_NewQCalendar4(void* name)
{
	return new QCalendar(*static_cast<QStringView*>(name));
}

char QCalendar_IsValid(void* ptr)
{
	return static_cast<QCalendar*>(ptr)->isValid();
}

struct QtCore_PackedString QCalendar_Name(void* ptr)
{
	return ({ QByteArray* tda699f = new QByteArray(static_cast<QCalendar*>(ptr)->name().toUtf8()); QtCore_PackedString { const_cast<char*>(tda699f->prepend("WHITESPACE").constData()+10), tda699f->size()-10, tda699f }; });
}

int QCborValue_False_Type()
{
	return QCborValue::False;
}

int QCborValue_True_Type()
{
	return QCborValue::True;
}

int QCborValue_Null_Type()
{
	return QCborValue::Null;
}

int QCborValue_Undefined_Type()
{
	return QCborValue::Undefined;
}

Q_DECLARE_METATYPE(QChar)
Q_DECLARE_METATYPE(QChar*)
void* QChar_NewQChar()
{
	return new QChar();
}

void* QChar_NewQChar2(unsigned short code)
{
	return new QChar(code);
}

void* QChar_NewQChar3(char* cell, char* row)
{
	return new QChar(*static_cast<uchar*>(static_cast<void*>(cell)), *static_cast<uchar*>(static_cast<void*>(row)));
}

void* QChar_NewQChar4(short code)
{
	return new QChar(code);
}

void* QChar_NewQChar5(unsigned int code)
{
	return new QChar(code);
}

void* QChar_NewQChar6(int code)
{
	return new QChar(code);
}

void* QChar_NewQChar7(long long ch)
{
	return new QChar(static_cast<QChar::SpecialCharacter>(ch));
}

void* QChar_NewQChar8(void* ch)
{
	return new QChar(*static_cast<QLatin1Char*>(ch));
}

void* QChar_NewQChar10(char* ch)
{
	return new QChar(*ch);
}

void* QChar_NewQChar11(char* ch)
{
	return new QChar(*static_cast<uchar*>(static_cast<void*>(ch)));
}

long long QChar_Category(void* ptr)
{
	return static_cast<QChar*>(ptr)->category();
}

long long QChar_QChar_Category2(unsigned int ucs4)
{
	return QChar::category(ucs4);
}

struct QtCore_PackedString QChar_Cell(void* ptr)
{
	return ({ uchar pret5084cd = static_cast<QChar*>(ptr)->cell(); char* t5084cd = static_cast<char*>(static_cast<void*>(&pret5084cd)); QtCore_PackedString { t5084cd, -1, NULL }; });
}

long long QChar_Direction(void* ptr)
{
	return static_cast<QChar*>(ptr)->direction();
}

long long QChar_QChar_Direction2(unsigned int ucs4)
{
	return QChar::direction(ucs4);
}

struct QtCore_PackedString QChar_Row(void* ptr)
{
	return ({ uchar pret52b865 = static_cast<QChar*>(ptr)->row(); char* t52b865 = static_cast<char*>(static_cast<void*>(&pret52b865)); QtCore_PackedString { t52b865, -1, NULL }; });
}

void* QChar_ToLower(void* ptr)
{
	return new QChar(static_cast<QChar*>(ptr)->toLower());
}

unsigned int QChar_QChar_ToLower2(unsigned int ucs4)
{
	return QChar::toLower(ucs4);
}

class MyQChildEvent: public QChildEvent
{
public:
	MyQChildEvent(QEvent::Type ty, QObject *child) : QChildEvent(ty, child) {QChildEvent_QChildEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QChildEvent*)
Q_DECLARE_METATYPE(MyQChildEvent*)

int QChildEvent_QChildEvent_QRegisterMetaType(){qRegisterMetaType<QChildEvent*>(); return qRegisterMetaType<MyQChildEvent*>();}

void* QChildEvent_NewQChildEvent(long long ty, void* child)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QGraphicsObject*>(child));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QGraphicsWidget*>(child));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QLayout*>(child));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QWidget*>(child));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QWindow*>(child));
	} else {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QObject*>(child));
	}
}

void* QChildEvent_Child(void* ptr)
{
	return static_cast<QChildEvent*>(ptr)->child();
}

class MyQCoreApplication: public QCoreApplication
{
public:
	MyQCoreApplication(int &argc, char **argv) : QCoreApplication(argc, argv) {QCoreApplication_QCoreApplication_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	void quit() { callbackQCoreApplication_Quit(this); };
	 ~MyQCoreApplication() { callbackQCoreApplication_DestroyQCoreApplication(this); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QCoreApplication*)
Q_DECLARE_METATYPE(MyQCoreApplication*)

int QCoreApplication_QCoreApplication_QRegisterMetaType(){qRegisterMetaType<QCoreApplication*>(); return qRegisterMetaType<MyQCoreApplication*>();}

void* QCoreApplication_NewQCoreApplication(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new MyQCoreApplication(argcs, argvs);
}

struct QtCore_PackedString QCoreApplication_QCoreApplication_ApplicationName()
{
	return ({ QByteArray* t704b73 = new QByteArray(QCoreApplication::applicationName().toUtf8()); QtCore_PackedString { const_cast<char*>(t704b73->prepend("WHITESPACE").constData()+10), t704b73->size()-10, t704b73 }; });
}

struct QtCore_PackedString QCoreApplication_QCoreApplication_ApplicationVersion()
{
	return ({ QByteArray* t9f1c49 = new QByteArray(QCoreApplication::applicationVersion().toUtf8()); QtCore_PackedString { const_cast<char*>(t9f1c49->prepend("WHITESPACE").constData()+10), t9f1c49->size()-10, t9f1c49 }; });
}

int QCoreApplication_QCoreApplication_Exec()
{
	return QCoreApplication::exec();
}

void QCoreApplication_Quit(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCoreApplication*>(ptr), "quit");
}

void QCoreApplication_QuitDefault(void* ptr)
{
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::quit();
}

void QCoreApplication_QCoreApplication_SetApplicationName(struct QtCore_PackedString application)
{
	QCoreApplication::setApplicationName(QString::fromUtf8(application.data, application.len));
}

void QCoreApplication_QCoreApplication_SetApplicationVersion(struct QtCore_PackedString version)
{
	QCoreApplication::setApplicationVersion(QString::fromUtf8(version.data, version.len));
}

void QCoreApplication_DestroyQCoreApplication(void* ptr)
{
	static_cast<QCoreApplication*>(ptr)->~QCoreApplication();
}

void QCoreApplication_DestroyQCoreApplicationDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

Q_DECLARE_METATYPE(QDataStream*)
void* QDataStream_NewQDataStream()
{
	return new QDataStream();
}

void* QDataStream_NewQDataStream2(void* d)
{
	return new QDataStream(static_cast<QIODevice*>(d));
}

void* QDataStream_NewQDataStream3(void* a, long long mode)
{
	return new QDataStream(static_cast<QByteArray*>(a), static_cast<QIODevice::OpenModeFlag>(mode));
}

void* QDataStream_NewQDataStream4(void* a)
{
	return new QDataStream(*static_cast<QByteArray*>(a));
}

void* QDataStream_Device(void* ptr)
{
	return static_cast<QDataStream*>(ptr)->device();
}

long long QDataStream_Status(void* ptr)
{
	return static_cast<QDataStream*>(ptr)->status();
}

int QDataStream_Version(void* ptr)
{
	return static_cast<QDataStream*>(ptr)->version();
}

void QDataStream_DestroyQDataStream(void* ptr)
{
	static_cast<QDataStream*>(ptr)->~QDataStream();
}

Q_DECLARE_METATYPE(QDate)
Q_DECLARE_METATYPE(QDate*)
void* QDate_NewQDate2()
{
	return new QDate();
}

void* QDate_NewQDate3(int y, int m, int d)
{
	return new QDate(y, m, d);
}

int QDate_Day(void* ptr, void* cal)
{
	return static_cast<QDate*>(ptr)->day(*static_cast<QCalendar*>(cal));
}

int QDate_Day2(void* ptr)
{
	return static_cast<QDate*>(ptr)->day();
}

void* QDate_QDate_FromString(struct QtCore_PackedString stri, long long format)
{
	return new QDate(QDate::fromString(QString::fromUtf8(stri.data, stri.len), static_cast<Qt::DateFormat>(format)));
}

void* QDate_QDate_FromString2(struct QtCore_PackedString stri, struct QtCore_PackedString format)
{
	return new QDate(QDate::fromString(QString::fromUtf8(stri.data, stri.len), QString::fromUtf8(format.data, format.len)));
}

void* QDate_QDate_FromString3(struct QtCore_PackedString stri, struct QtCore_PackedString format, void* cal)
{
	return new QDate(QDate::fromString(QString::fromUtf8(stri.data, stri.len), QString::fromUtf8(format.data, format.len), *static_cast<QCalendar*>(cal)));
}

char QDate_IsValid(void* ptr)
{
	return static_cast<QDate*>(ptr)->isValid();
}

char QDate_QDate_IsValid2(int year, int month, int day)
{
	return QDate::isValid(year, month, day);
}

int QDate_Month(void* ptr, void* cal)
{
	return static_cast<QDate*>(ptr)->month(*static_cast<QCalendar*>(cal));
}

int QDate_Month2(void* ptr)
{
	return static_cast<QDate*>(ptr)->month();
}

int QDate_Year(void* ptr, void* cal)
{
	return static_cast<QDate*>(ptr)->year(*static_cast<QCalendar*>(cal));
}

int QDate_Year2(void* ptr)
{
	return static_cast<QDate*>(ptr)->year();
}

Q_DECLARE_METATYPE(QDateTime)
Q_DECLARE_METATYPE(QDateTime*)
void* QDateTime_NewQDateTime()
{
	return new QDateTime();
}

void* QDateTime_NewQDateTime2(void* date)
{
	return new QDateTime(*static_cast<QDate*>(date));
}

void* QDateTime_NewQDateTime3(void* date, void* ti, long long spec)
{
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(ti), static_cast<Qt::TimeSpec>(spec));
}

void* QDateTime_NewQDateTime4(void* date, void* ti, long long spec, int offsetSeconds)
{
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(ti), static_cast<Qt::TimeSpec>(spec), offsetSeconds);
}

void* QDateTime_NewQDateTime5(void* date, void* ti, void* timeZone)
{
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(ti), *static_cast<QTimeZone*>(timeZone));
}

void* QDateTime_NewQDateTime6(void* other)
{
	return new QDateTime(*static_cast<QDateTime*>(other));
}

void* QDateTime_NewQDateTime7(void* other)
{
	return new QDateTime(*static_cast<QDateTime*>(other));
}

void* QDateTime_QDateTime_FromString(struct QtCore_PackedString stri, long long format)
{
	return new QDateTime(QDateTime::fromString(QString::fromUtf8(stri.data, stri.len), static_cast<Qt::DateFormat>(format)));
}

void* QDateTime_QDateTime_FromString2(struct QtCore_PackedString stri, struct QtCore_PackedString format)
{
	return new QDateTime(QDateTime::fromString(QString::fromUtf8(stri.data, stri.len), QString::fromUtf8(format.data, format.len)));
}

void* QDateTime_QDateTime_FromString3(struct QtCore_PackedString stri, struct QtCore_PackedString format, void* cal)
{
	return new QDateTime(QDateTime::fromString(QString::fromUtf8(stri.data, stri.len), QString::fromUtf8(format.data, format.len), *static_cast<QCalendar*>(cal)));
}

char QDateTime_IsValid(void* ptr)
{
	return static_cast<QDateTime*>(ptr)->isValid();
}

void* QDateTime_Time(void* ptr)
{
	return new QTime(static_cast<QDateTime*>(ptr)->time());
}

void QDateTime_DestroyQDateTime(void* ptr)
{
	static_cast<QDateTime*>(ptr)->~QDateTime();
}

Q_DECLARE_METATYPE(QDir*)
void* QDir_NewQDir(void* dir)
{
	return new QDir(*static_cast<QDir*>(dir));
}

void* QDir_NewQDir2(struct QtCore_PackedString path)
{
	return new QDir(QString::fromUtf8(path.data, path.len));
}

void* QDir_NewQDir3(struct QtCore_PackedString path, struct QtCore_PackedString nameFilter, long long sort, long long filters)
{
	return new QDir(QString::fromUtf8(path.data, path.len), QString::fromUtf8(nameFilter.data, nameFilter.len), static_cast<QDir::SortFlag>(sort), static_cast<QDir::Filter>(filters));
}

char QDir_Cd(void* ptr, struct QtCore_PackedString dirName)
{
	return static_cast<QDir*>(ptr)->cd(QString::fromUtf8(dirName.data, dirName.len));
}

unsigned int QDir_Count(void* ptr)
{
	return static_cast<QDir*>(ptr)->count();
}

void* QDir_QDir_Current()
{
	return new QDir(QDir::current());
}

char QDir_Exists(void* ptr, struct QtCore_PackedString name)
{
	return static_cast<QDir*>(ptr)->exists(QString::fromUtf8(name.data, name.len));
}

char QDir_Exists2(void* ptr)
{
	return static_cast<QDir*>(ptr)->exists();
}

long long QDir_Filter(void* ptr)
{
	return static_cast<QDir*>(ptr)->filter();
}

void* QDir_QDir_Home()
{
	return new QDir(QDir::home());
}

char QDir_IsEmpty(void* ptr, long long filters)
{
	return static_cast<QDir*>(ptr)->isEmpty(static_cast<QDir::Filter>(filters));
}

char QDir_QDir_Match(struct QtCore_PackedString filter, struct QtCore_PackedString fileName)
{
	return QDir::match(QString::fromUtf8(filter.data, filter.len), QString::fromUtf8(fileName.data, fileName.len));
}

char QDir_QDir_Match2(struct QtCore_PackedString filters, struct QtCore_PackedString fileName)
{
	return QDir::match(QString::fromUtf8(filters.data, filters.len).split("¡¦!", QString::SkipEmptyParts), QString::fromUtf8(fileName.data, fileName.len));
}

char QDir_Mkdir(void* ptr, struct QtCore_PackedString dirName)
{
	return static_cast<QDir*>(ptr)->mkdir(QString::fromUtf8(dirName.data, dirName.len));
}

struct QtCore_PackedString QDir_Path(void* ptr)
{
	return ({ QByteArray* t1e0939 = new QByteArray(static_cast<QDir*>(ptr)->path().toUtf8()); QtCore_PackedString { const_cast<char*>(t1e0939->prepend("WHITESPACE").constData()+10), t1e0939->size()-10, t1e0939 }; });
}

void QDir_Refresh(void* ptr)
{
	static_cast<QDir*>(ptr)->refresh();
}

char QDir_Remove(void* ptr, struct QtCore_PackedString fileName)
{
	return static_cast<QDir*>(ptr)->remove(QString::fromUtf8(fileName.data, fileName.len));
}

char QDir_Rename(void* ptr, struct QtCore_PackedString oldName, struct QtCore_PackedString newName)
{
	return static_cast<QDir*>(ptr)->rename(QString::fromUtf8(oldName.data, oldName.len), QString::fromUtf8(newName.data, newName.len));
}

void* QDir_QDir_Root()
{
	return new QDir(QDir::root());
}

void* QDir_QDir_Separator()
{
	return new QChar(QDir::separator());
}

char QDir_QDir_SetCurrent(struct QtCore_PackedString path)
{
	return QDir::setCurrent(QString::fromUtf8(path.data, path.len));
}

void* QDir_QDir_Temp()
{
	return new QDir(QDir::temp());
}

struct QtCore_PackedString QDir_QDir_ToNativeSeparators(struct QtCore_PackedString pathName)
{
	return ({ QByteArray* tf0acff = new QByteArray(QDir::toNativeSeparators(QString::fromUtf8(pathName.data, pathName.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(tf0acff->prepend("WHITESPACE").constData()+10), tf0acff->size()-10, tf0acff }; });
}

void QDir_DestroyQDir(void* ptr)
{
	static_cast<QDir*>(ptr)->~QDir();
}

void* QDir___drives_atList(void* ptr, int i)
{
	return new QFileInfo(({QFileInfo tmp = static_cast<QList<QFileInfo>*>(ptr)->at(i); if (i == static_cast<QList<QFileInfo>*>(ptr)->size()-1) { static_cast<QList<QFileInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDir___drives_setList(void* ptr, void* i)
{
	static_cast<QList<QFileInfo>*>(ptr)->append(*static_cast<QFileInfo*>(i));
}

void* QDir___drives_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QFileInfo>();
}

void* QDir___entryInfoList_atList(void* ptr, int i)
{
	return new QFileInfo(({QFileInfo tmp = static_cast<QList<QFileInfo>*>(ptr)->at(i); if (i == static_cast<QList<QFileInfo>*>(ptr)->size()-1) { static_cast<QList<QFileInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDir___entryInfoList_setList(void* ptr, void* i)
{
	static_cast<QList<QFileInfo>*>(ptr)->append(*static_cast<QFileInfo*>(i));
}

void* QDir___entryInfoList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QFileInfo>();
}

void* QDir___entryInfoList_atList2(void* ptr, int i)
{
	return new QFileInfo(({QFileInfo tmp = static_cast<QList<QFileInfo>*>(ptr)->at(i); if (i == static_cast<QList<QFileInfo>*>(ptr)->size()-1) { static_cast<QList<QFileInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDir___entryInfoList_setList2(void* ptr, void* i)
{
	static_cast<QList<QFileInfo>*>(ptr)->append(*static_cast<QFileInfo*>(i));
}

void* QDir___entryInfoList_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QFileInfo>();
}

Q_DECLARE_METATYPE(QEasingCurve*)
void* QEasingCurve_NewQEasingCurve(long long ty)
{
	return new QEasingCurve(static_cast<QEasingCurve::Type>(ty));
}

void* QEasingCurve_NewQEasingCurve2(void* other)
{
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

void* QEasingCurve_NewQEasingCurve3(void* other)
{
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

long long QEasingCurve_Type(void* ptr)
{
	return static_cast<QEasingCurve*>(ptr)->type();
}

void QEasingCurve_DestroyQEasingCurve(void* ptr)
{
	static_cast<QEasingCurve*>(ptr)->~QEasingCurve();
}

void* QEasingCurve___cubicBezierSpline_atList(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QList<QPointF>*>(ptr)->at(i); if (i == static_cast<QList<QPointF>*>(ptr)->size()-1) { static_cast<QList<QPointF>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QEasingCurve___cubicBezierSpline_setList(void* ptr, void* i)
{
	static_cast<QList<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QEasingCurve___cubicBezierSpline_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPointF>();
}

void* QEasingCurve___toCubicSpline_atList(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QEasingCurve___toCubicSpline_setList(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QEasingCurve___toCubicSpline_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

class MyQEvent: public QEvent
{
public:
	MyQEvent(QEvent::Type ty) : QEvent(ty) {QEvent_QEvent_QRegisterMetaType();};
	 ~MyQEvent() { callbackQEvent_DestroyQEvent(this); };
};

Q_DECLARE_METATYPE(QEvent*)
Q_DECLARE_METATYPE(MyQEvent*)

int QEvent_QEvent_QRegisterMetaType(){qRegisterMetaType<QEvent*>(); return qRegisterMetaType<MyQEvent*>();}

void* QEvent_NewQEvent(long long ty)
{
	return new MyQEvent(static_cast<QEvent::Type>(ty));
}

void QEvent_Accept(void* ptr)
{
	static_cast<QEvent*>(ptr)->accept();
}

long long QEvent_Type(void* ptr)
{
	return static_cast<QEvent*>(ptr)->type();
}

void QEvent_DestroyQEvent(void* ptr)
{
	static_cast<QEvent*>(ptr)->~QEvent();
}

void QEvent_DestroyQEventDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQFile: public QFile
{
public:
	MyQFile() : QFile() {QFile_QFile_QRegisterMetaType();};
	MyQFile(const QString &name) : QFile(name) {QFile_QFile_QRegisterMetaType();};
	MyQFile(QObject *parent) : QFile(parent) {QFile_QFile_QRegisterMetaType();};
	MyQFile(const QString &name, QObject *parent) : QFile(name, parent) {QFile_QFile_QRegisterMetaType();};
	QString fileName() const { return ({ QtCore_PackedString tempVal = callbackQFileDevice_FileName(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool open(QIODevice::OpenMode mode) { return callbackQIODevice_Open(this, mode) != 0; };
	bool resize(qint64 sz) { return callbackQFileDevice_Resize(this, sz) != 0; };
	qint64 size() const { return callbackQIODevice_Size(const_cast<void*>(static_cast<const void*>(this))); };
	 ~MyQFile() { callbackQFile_DestroyQFile(this); };
	void close() { callbackQIODevice_Close(this); };
	qint64 pos() const { return callbackQIODevice_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 readData(char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { data, maxSize, NULL };return callbackQFileDevice_ReadData(this, dataPacked, maxSize); };
	qint64 writeData(const char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { const_cast<char*>(data), maxSize, NULL };return callbackQFileDevice_WriteData(this, dataPacked, maxSize); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QFile*)
Q_DECLARE_METATYPE(MyQFile*)

int QFile_QFile_QRegisterMetaType(){qRegisterMetaType<QFile*>(); return qRegisterMetaType<MyQFile*>();}

void* QFile_NewQFile()
{
	return new MyQFile();
}

void* QFile_NewQFile2(struct QtCore_PackedString name)
{
	return new MyQFile(QString::fromUtf8(name.data, name.len));
}

void* QFile_NewQFile3(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QWindow*>(parent));
	} else {
		return new MyQFile(static_cast<QObject*>(parent));
	}
}

void* QFile_NewQFile4(struct QtCore_PackedString name, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QObject*>(parent));
	}
}

char QFile_Copy(void* ptr, struct QtCore_PackedString newName)
{
	return static_cast<QFile*>(ptr)->copy(QString::fromUtf8(newName.data, newName.len));
}

char QFile_QFile_Copy2(struct QtCore_PackedString fileName, struct QtCore_PackedString newName)
{
	return QFile::copy(QString::fromUtf8(fileName.data, fileName.len), QString::fromUtf8(newName.data, newName.len));
}

char QFile_QFile_Exists(struct QtCore_PackedString fileName)
{
	return QFile::exists(QString::fromUtf8(fileName.data, fileName.len));
}

char QFile_Exists2(void* ptr)
{
	return static_cast<QFile*>(ptr)->exists();
}

char QFile_Open3(void* ptr, int fd, long long mode, long long handleFlags)
{
	return static_cast<QFile*>(ptr)->open(fd, static_cast<QIODevice::OpenModeFlag>(mode), static_cast<QFileDevice::FileHandleFlag>(handleFlags));
}

char QFile_Remove(void* ptr)
{
	return static_cast<QFile*>(ptr)->remove();
}

char QFile_QFile_Remove2(struct QtCore_PackedString fileName)
{
	return QFile::remove(QString::fromUtf8(fileName.data, fileName.len));
}

char QFile_Rename(void* ptr, struct QtCore_PackedString newName)
{
	return static_cast<QFile*>(ptr)->rename(QString::fromUtf8(newName.data, newName.len));
}

char QFile_QFile_Rename2(struct QtCore_PackedString oldName, struct QtCore_PackedString newName)
{
	return QFile::rename(QString::fromUtf8(oldName.data, oldName.len), QString::fromUtf8(newName.data, newName.len));
}

char QFile_QFile_Resize2(struct QtCore_PackedString fileName, long long sz)
{
	return QFile::resize(QString::fromUtf8(fileName.data, fileName.len), sz);
}

void QFile_DestroyQFile(void* ptr)
{
	static_cast<QFile*>(ptr)->~QFile();
}

void QFile_DestroyQFileDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQFileDevice: public QFileDevice
{
public:
	void close() { callbackQIODevice_Close(this); };
	QString fileName() const { return ({ QtCore_PackedString tempVal = callbackQFileDevice_FileName(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	qint64 pos() const { return callbackQIODevice_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	bool resize(qint64 sz) { return callbackQFileDevice_Resize(this, sz) != 0; };
	qint64 size() const { return callbackQIODevice_Size(const_cast<void*>(static_cast<const void*>(this))); };
	 ~MyQFileDevice() { callbackQFileDevice_DestroyQFileDevice(this); };
	bool open(QIODevice::OpenMode mode) { return callbackQIODevice_Open(this, mode) != 0; };
	qint64 readData(char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { data, maxSize, NULL };return callbackQFileDevice_ReadData(this, dataPacked, maxSize); };
	qint64 writeData(const char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { const_cast<char*>(data), maxSize, NULL };return callbackQFileDevice_WriteData(this, dataPacked, maxSize); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QFileDevice*)
Q_DECLARE_METATYPE(MyQFileDevice*)

int QFileDevice_QFileDevice_QRegisterMetaType(){qRegisterMetaType<QFileDevice*>(); return qRegisterMetaType<MyQFileDevice*>();}

long long QFileDevice_Error(void* ptr)
{
	return static_cast<QFileDevice*>(ptr)->error();
}

struct QtCore_PackedString QFileDevice_FileName(void* ptr)
{
	return ({ QByteArray* t4e2118 = new QByteArray(static_cast<QFileDevice*>(ptr)->fileName().toUtf8()); QtCore_PackedString { const_cast<char*>(t4e2118->prepend("WHITESPACE").constData()+10), t4e2118->size()-10, t4e2118 }; });
}

struct QtCore_PackedString QFileDevice_FileNameDefault(void* ptr)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray* tb198ab = new QByteArray(static_cast<QFile*>(ptr)->QFile::fileName().toUtf8()); QtCore_PackedString { const_cast<char*>(tb198ab->prepend("WHITESPACE").constData()+10), tb198ab->size()-10, tb198ab }; });
	} else {
		return ({ QByteArray* tb198ab = new QByteArray(static_cast<QFileDevice*>(ptr)->QFileDevice::fileName().toUtf8()); QtCore_PackedString { const_cast<char*>(tb198ab->prepend("WHITESPACE").constData()+10), tb198ab->size()-10, tb198ab }; });
	}
}

struct QtCore_PackedString QFileDevice_Map(void* ptr, long long offset, long long size, long long flags)
{
	return ({ char* t4b412c = static_cast<char*>(static_cast<void*>(static_cast<QFileDevice*>(ptr)->map(offset, size, static_cast<QFileDevice::MemoryMapFlags>(flags)))); QtCore_PackedString { t4b412c, -1, NULL }; });
}

char QFileDevice_Resize(void* ptr, long long sz)
{
	return static_cast<QFileDevice*>(ptr)->resize(sz);
}

char QFileDevice_ResizeDefault(void* ptr, long long sz)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::resize(sz);
	} else {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::resize(sz);
	}
}

void QFileDevice_DestroyQFileDevice(void* ptr)
{
	static_cast<QFileDevice*>(ptr)->~QFileDevice();
}

void QFileDevice_DestroyQFileDeviceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QFileDevice_ReadData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QFileDevice*>(ptr)->readData(data, maxSize);
}

long long QFileDevice_ReadDataDefault(void* ptr, char* data, long long maxSize)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::readData(data, maxSize);
	} else {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::readData(data, maxSize);
	}
}

long long QFileDevice_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QFileDevice*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

long long QFileDevice_WriteDataDefault(void* ptr, char* data, long long maxSize)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::writeData(const_cast<const char*>(data), maxSize);
	} else {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::writeData(const_cast<const char*>(data), maxSize);
	}
}

Q_DECLARE_METATYPE(QFileInfo*)
void* QFileInfo_NewQFileInfo2()
{
	return new QFileInfo();
}

void* QFileInfo_NewQFileInfo3(struct QtCore_PackedString file)
{
	return new QFileInfo(QString::fromUtf8(file.data, file.len));
}

void* QFileInfo_NewQFileInfo4(void* file)
{
	return new QFileInfo(*static_cast<QFile*>(file));
}

void* QFileInfo_NewQFileInfo5(void* dir, struct QtCore_PackedString file)
{
	return new QFileInfo(*static_cast<QDir*>(dir), QString::fromUtf8(file.data, file.len));
}

void* QFileInfo_NewQFileInfo6(void* fileinfo)
{
	return new QFileInfo(*static_cast<QFileInfo*>(fileinfo));
}

void* QFileInfo_Dir(void* ptr)
{
	return new QDir(static_cast<QFileInfo*>(ptr)->dir());
}

char QFileInfo_Exists(void* ptr)
{
	return static_cast<QFileInfo*>(ptr)->exists();
}

char QFileInfo_QFileInfo_Exists2(struct QtCore_PackedString file)
{
	return QFileInfo::exists(QString::fromUtf8(file.data, file.len));
}

struct QtCore_PackedString QFileInfo_FileName(void* ptr)
{
	return ({ QByteArray* t8cf8a1 = new QByteArray(static_cast<QFileInfo*>(ptr)->fileName().toUtf8()); QtCore_PackedString { const_cast<char*>(t8cf8a1->prepend("WHITESPACE").constData()+10), t8cf8a1->size()-10, t8cf8a1 }; });
}

struct QtCore_PackedString QFileInfo_Group(void* ptr)
{
	return ({ QByteArray* ta89964 = new QByteArray(static_cast<QFileInfo*>(ptr)->group().toUtf8()); QtCore_PackedString { const_cast<char*>(ta89964->prepend("WHITESPACE").constData()+10), ta89964->size()-10, ta89964 }; });
}

char QFileInfo_IsDir(void* ptr)
{
	return static_cast<QFileInfo*>(ptr)->isDir();
}

struct QtCore_PackedString QFileInfo_Path(void* ptr)
{
	return ({ QByteArray* tdcd027 = new QByteArray(static_cast<QFileInfo*>(ptr)->path().toUtf8()); QtCore_PackedString { const_cast<char*>(tdcd027->prepend("WHITESPACE").constData()+10), tdcd027->size()-10, tdcd027 }; });
}

void QFileInfo_Refresh(void* ptr)
{
	static_cast<QFileInfo*>(ptr)->refresh();
}

long long QFileInfo_Size(void* ptr)
{
	return static_cast<QFileInfo*>(ptr)->size();
}

struct QtCore_PackedString QFileInfo_Suffix(void* ptr)
{
	return ({ QByteArray* t1b5684 = new QByteArray(static_cast<QFileInfo*>(ptr)->suffix().toUtf8()); QtCore_PackedString { const_cast<char*>(t1b5684->prepend("WHITESPACE").constData()+10), t1b5684->size()-10, t1b5684 }; });
}

void QFileInfo_DestroyQFileInfo(void* ptr)
{
	static_cast<QFileInfo*>(ptr)->~QFileInfo();
}

class MyQIODevice: public QIODevice
{
public:
	MyQIODevice() : QIODevice() {QIODevice_QIODevice_QRegisterMetaType();};
	MyQIODevice(QObject *parent) : QIODevice(parent) {QIODevice_QIODevice_QRegisterMetaType();};
	void close() { callbackQIODevice_Close(this); };
	bool open(QIODevice::OpenMode mode) { return callbackQIODevice_Open(this, mode) != 0; };
	qint64 pos() const { return callbackQIODevice_Pos(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 readData(char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { data, maxSize, NULL };return callbackQIODevice_ReadData(this, dataPacked, maxSize); };
	qint64 size() const { return callbackQIODevice_Size(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 writeData(const char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { const_cast<char*>(data), maxSize, NULL };return callbackQIODevice_WriteData(this, dataPacked, maxSize); };
	 ~MyQIODevice() { callbackQIODevice_DestroyQIODevice(this); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QIODevice*)
Q_DECLARE_METATYPE(MyQIODevice*)

int QIODevice_QIODevice_QRegisterMetaType(){qRegisterMetaType<QIODevice*>(); return qRegisterMetaType<MyQIODevice*>();}

void* QIODevice_NewQIODevice()
{
	return new MyQIODevice();
}

void* QIODevice_NewQIODevice2(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QWindow*>(parent));
	} else {
		return new MyQIODevice(static_cast<QObject*>(parent));
	}
}

void QIODevice_Close(void* ptr)
{
	static_cast<QIODevice*>(ptr)->close();
}

void QIODevice_CloseDefault(void* ptr)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::close();
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::close();
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		static_cast<QBuffer*>(ptr)->QBuffer::close();
	} else {
		static_cast<QIODevice*>(ptr)->QIODevice::close();
	}
}

char QIODevice_Open(void* ptr, long long mode)
{
	return static_cast<QIODevice*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QIODevice_OpenDefault(void* ptr, long long mode)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::open(static_cast<QIODevice::OpenModeFlag>(mode));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::open(static_cast<QIODevice::OpenModeFlag>(mode));
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBuffer*>(ptr)->QBuffer::open(static_cast<QIODevice::OpenModeFlag>(mode));
	} else {
		return static_cast<QIODevice*>(ptr)->QIODevice::open(static_cast<QIODevice::OpenModeFlag>(mode));
	}
}

long long QIODevice_Pos(void* ptr)
{
	return static_cast<QIODevice*>(ptr)->pos();
}

long long QIODevice_PosDefault(void* ptr)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::pos();
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::pos();
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBuffer*>(ptr)->QBuffer::pos();
	} else {
		return static_cast<QIODevice*>(ptr)->QIODevice::pos();
	}
}

long long QIODevice_Read(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->read(data, maxSize);
}

void* QIODevice_Read2(void* ptr, long long maxSize)
{
	return new QByteArray(static_cast<QIODevice*>(ptr)->read(maxSize));
}

long long QIODevice_ReadData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->readData(data, maxSize);
}

long long QIODevice_ReadLine(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->readLine(data, maxSize);
}

void* QIODevice_ReadLine2(void* ptr, long long maxSize)
{
	return new QByteArray(static_cast<QIODevice*>(ptr)->readLine(maxSize));
}

long long QIODevice_Size(void* ptr)
{
	return static_cast<QIODevice*>(ptr)->size();
}

long long QIODevice_SizeDefault(void* ptr)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::size();
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::size();
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBuffer*>(ptr)->QBuffer::size();
	} else {
		return static_cast<QIODevice*>(ptr)->QIODevice::size();
	}
}

long long QIODevice_Write(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->write(const_cast<const char*>(data), maxSize);
}

long long QIODevice_Write2(void* ptr, char* data)
{
	return static_cast<QIODevice*>(ptr)->write(const_cast<const char*>(data));
}

long long QIODevice_Write3(void* ptr, void* byteArray)
{
	return static_cast<QIODevice*>(ptr)->write(*static_cast<QByteArray*>(byteArray));
}

long long QIODevice_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

void QIODevice_DestroyQIODevice(void* ptr)
{
	static_cast<QIODevice*>(ptr)->~QIODevice();
}

void QIODevice_DestroyQIODeviceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

Q_DECLARE_METATYPE(QItemSelection*)
void* QItemSelection_NewQItemSelection()
{
	return new QItemSelection();
}

void* QItemSelection_NewQItemSelection2(void* topLeft, void* bottomRight)
{
	return new QItemSelection(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

struct QtCore_PackedList QItemSelection_Indexes(void* ptr)
{
	return ({ QList<QModelIndex>* tmpValue5090cf = new QList<QModelIndex>(static_cast<QItemSelection*>(ptr)->indexes()); QtCore_PackedList { tmpValue5090cf, tmpValue5090cf->size() }; });
}

void QItemSelection_Merge(void* ptr, void* other, long long command)
{
	static_cast<QItemSelection*>(ptr)->merge(*static_cast<QItemSelection*>(other), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QItemSelection_Select(void* ptr, void* topLeft, void* bottomRight)
{
	static_cast<QItemSelection*>(ptr)->select(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

void QItemSelection_QItemSelection_Split(void* ran, void* other, void* result)
{
	QItemSelection::split(*static_cast<QItemSelectionRange*>(ran), *static_cast<QItemSelectionRange*>(other), static_cast<QItemSelection*>(result));
}

void* QItemSelection___indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QItemSelection___indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QItemSelection___indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

Q_DECLARE_METATYPE(QItemSelectionRange*)
void* QItemSelectionRange_NewQItemSelectionRange()
{
	return new QItemSelectionRange();
}

void* QItemSelectionRange_NewQItemSelectionRange2(void* other)
{
	return new QItemSelectionRange(*static_cast<QItemSelectionRange*>(other));
}

void* QItemSelectionRange_NewQItemSelectionRange4(void* topLeft, void* bottomRight)
{
	return new QItemSelectionRange(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight));
}

void* QItemSelectionRange_NewQItemSelectionRange5(void* index)
{
	return new QItemSelectionRange(*static_cast<QModelIndex*>(index));
}

int QItemSelectionRange_Height(void* ptr)
{
	return static_cast<QItemSelectionRange*>(ptr)->height();
}

struct QtCore_PackedList QItemSelectionRange_Indexes(void* ptr)
{
	return ({ QList<QModelIndex>* tmpValuecb2327 = new QList<QModelIndex>(static_cast<QItemSelectionRange*>(ptr)->indexes()); QtCore_PackedList { tmpValuecb2327, tmpValuecb2327->size() }; });
}

char QItemSelectionRange_IsEmpty(void* ptr)
{
	return static_cast<QItemSelectionRange*>(ptr)->isEmpty();
}

char QItemSelectionRange_IsValid(void* ptr)
{
	return static_cast<QItemSelectionRange*>(ptr)->isValid();
}

int QItemSelectionRange_Left(void* ptr)
{
	return static_cast<QItemSelectionRange*>(ptr)->left();
}

void* QItemSelectionRange_Model(void* ptr)
{
	return const_cast<QAbstractItemModel*>(static_cast<QItemSelectionRange*>(ptr)->model());
}

void* QItemSelectionRange_Parent(void* ptr)
{
	return new QModelIndex(static_cast<QItemSelectionRange*>(ptr)->parent());
}

int QItemSelectionRange_Right(void* ptr)
{
	return static_cast<QItemSelectionRange*>(ptr)->right();
}

int QItemSelectionRange_Top(void* ptr)
{
	return static_cast<QItemSelectionRange*>(ptr)->top();
}

int QItemSelectionRange_Width(void* ptr)
{
	return static_cast<QItemSelectionRange*>(ptr)->width();
}

void* QItemSelectionRange___indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QItemSelectionRange___indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QItemSelectionRange___indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

Q_DECLARE_METATYPE(QJsonArray)
Q_DECLARE_METATYPE(QJsonArray*)
void* QJsonArray_NewQJsonArray()
{
	return new QJsonArray();
}

void* QJsonArray_NewQJsonArray3(void* other)
{
	return new QJsonArray(*static_cast<QJsonArray*>(other));
}

void* QJsonArray_NewQJsonArray4(void* other)
{
	return new QJsonArray(*static_cast<QJsonArray*>(other));
}

void QJsonArray_Append(void* ptr, void* value)
{
	static_cast<QJsonArray*>(ptr)->append(*static_cast<QJsonValue*>(value));
}

void* QJsonArray_At(void* ptr, int i)
{
	return new QJsonValue(static_cast<QJsonArray*>(ptr)->at(i));
}

int QJsonArray_Count(void* ptr)
{
	return static_cast<QJsonArray*>(ptr)->count();
}

char QJsonArray_Empty(void* ptr)
{
	return static_cast<QJsonArray*>(ptr)->empty();
}

void* QJsonArray_First(void* ptr)
{
	return new QJsonValue(static_cast<QJsonArray*>(ptr)->first());
}

void QJsonArray_Insert(void* ptr, int i, void* value)
{
	static_cast<QJsonArray*>(ptr)->insert(i, *static_cast<QJsonValue*>(value));
}

char QJsonArray_IsEmpty(void* ptr)
{
	return static_cast<QJsonArray*>(ptr)->isEmpty();
}

void* QJsonArray_Last(void* ptr)
{
	return new QJsonValue(static_cast<QJsonArray*>(ptr)->last());
}

void QJsonArray_Replace(void* ptr, int i, void* value)
{
	static_cast<QJsonArray*>(ptr)->replace(i, *static_cast<QJsonValue*>(value));
}

int QJsonArray_Size(void* ptr)
{
	return static_cast<QJsonArray*>(ptr)->size();
}

void* QJsonArray_TakeAt(void* ptr, int i)
{
	return new QJsonValue(static_cast<QJsonArray*>(ptr)->takeAt(i));
}

void QJsonArray_DestroyQJsonArray(void* ptr)
{
	static_cast<QJsonArray*>(ptr)->~QJsonArray();
}

void* QJsonArray___fromVariantList_list_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJsonArray___fromVariantList_list_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QJsonArray___fromVariantList_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QJsonArray___toVariantList_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJsonArray___toVariantList_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QJsonArray___toVariantList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

Q_DECLARE_METATYPE(QJsonDocument)
Q_DECLARE_METATYPE(QJsonDocument*)
void* QJsonDocument_NewQJsonDocument()
{
	return new QJsonDocument();
}

void* QJsonDocument_NewQJsonDocument2(void* object)
{
	return new QJsonDocument(*static_cast<QJsonObject*>(object));
}

void* QJsonDocument_NewQJsonDocument3(void* array)
{
	return new QJsonDocument(*static_cast<QJsonArray*>(array));
}

void* QJsonDocument_NewQJsonDocument4(void* other)
{
	return new QJsonDocument(*static_cast<QJsonDocument*>(other));
}

void* QJsonDocument_NewQJsonDocument5(void* other)
{
	return new QJsonDocument(*static_cast<QJsonDocument*>(other));
}

void* QJsonDocument_Array(void* ptr)
{
	return new QJsonArray(static_cast<QJsonDocument*>(ptr)->array());
}

char QJsonDocument_IsEmpty(void* ptr)
{
	return static_cast<QJsonDocument*>(ptr)->isEmpty();
}

void* QJsonDocument_Object(void* ptr)
{
	return new QJsonObject(static_cast<QJsonDocument*>(ptr)->object());
}

void QJsonDocument_SetObject(void* ptr, void* object)
{
	static_cast<QJsonDocument*>(ptr)->setObject(*static_cast<QJsonObject*>(object));
}

void* QJsonDocument_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QJsonDocument*>(ptr)->toVariant());
}

void QJsonDocument_DestroyQJsonDocument(void* ptr)
{
	static_cast<QJsonDocument*>(ptr)->~QJsonDocument();
}

Q_DECLARE_METATYPE(QJsonObject)
Q_DECLARE_METATYPE(QJsonObject*)
void* QJsonObject_NewQJsonObject()
{
	return new QJsonObject();
}

void* QJsonObject_NewQJsonObject3(void* other)
{
	return new QJsonObject(*static_cast<QJsonObject*>(other));
}

void* QJsonObject_NewQJsonObject4(void* other)
{
	return new QJsonObject(*static_cast<QJsonObject*>(other));
}

int QJsonObject_Count(void* ptr)
{
	return static_cast<QJsonObject*>(ptr)->count();
}

char QJsonObject_Empty(void* ptr)
{
	return static_cast<QJsonObject*>(ptr)->empty();
}

char QJsonObject_IsEmpty(void* ptr)
{
	return static_cast<QJsonObject*>(ptr)->isEmpty();
}

int QJsonObject_Length(void* ptr)
{
	return static_cast<QJsonObject*>(ptr)->length();
}

void QJsonObject_Remove(void* ptr, struct QtCore_PackedString key)
{
	static_cast<QJsonObject*>(ptr)->remove(QString::fromUtf8(key.data, key.len));
}

void QJsonObject_Remove2(void* ptr, void* key)
{
	static_cast<QJsonObject*>(ptr)->remove(*static_cast<QStringView*>(key));
}

void QJsonObject_Remove3(void* ptr, void* key)
{
	static_cast<QJsonObject*>(ptr)->remove(*static_cast<QLatin1String*>(key));
}

int QJsonObject_Size(void* ptr)
{
	return static_cast<QJsonObject*>(ptr)->size();
}

void* QJsonObject_Take(void* ptr, struct QtCore_PackedString key)
{
	return new QJsonValue(static_cast<QJsonObject*>(ptr)->take(QString::fromUtf8(key.data, key.len)));
}

void* QJsonObject_Take2(void* ptr, void* key)
{
	return new QJsonValue(static_cast<QJsonObject*>(ptr)->take(*static_cast<QStringView*>(key)));
}

void* QJsonObject_Take3(void* ptr, void* key)
{
	return new QJsonValue(static_cast<QJsonObject*>(ptr)->take(*static_cast<QLatin1String*>(key)));
}

void* QJsonObject_Value(void* ptr, struct QtCore_PackedString key)
{
	return new QJsonValue(static_cast<QJsonObject*>(ptr)->value(QString::fromUtf8(key.data, key.len)));
}

void* QJsonObject_Value2(void* ptr, void* key)
{
	return new QJsonValue(static_cast<QJsonObject*>(ptr)->value(*static_cast<QStringView*>(key)));
}

void* QJsonObject_Value3(void* ptr, void* key)
{
	return new QJsonValue(static_cast<QJsonObject*>(ptr)->value(*static_cast<QLatin1String*>(key)));
}

void QJsonObject_DestroyQJsonObject(void* ptr)
{
	static_cast<QJsonObject*>(ptr)->~QJsonObject();
}

void* QJsonObject___fromVariantHash_hash_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QJsonObject___fromVariantHash_hash_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QJsonObject___fromVariantHash_hash_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QJsonObject___fromVariantHash_hash_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QJsonObject___toVariantHash_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QJsonObject___toVariantHash_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QJsonObject___toVariantHash_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QJsonObject___toVariantHash_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QJsonObject___toVariantMap_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QJsonObject___toVariantMap_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QJsonObject___toVariantMap_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtCore_PackedList QJsonObject___toVariantMap_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtCore_PackedString QJsonObject_____fromVariantHash_hash_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QJsonObject_____fromVariantHash_hash_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____fromVariantHash_hash_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QJsonObject_____fromVariantMap_map_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QJsonObject_____fromVariantMap_map_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____fromVariantMap_map_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QJsonObject_____toVariantHash_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QJsonObject_____toVariantHash_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____toVariantHash_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QJsonObject_____toVariantMap_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QJsonObject_____toVariantMap_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____toVariantMap_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

Q_DECLARE_METATYPE(QJsonValue*)
void* QJsonValue_NewQJsonValue(long long ty)
{
	return new QJsonValue(static_cast<QJsonValue::Type>(ty));
}

void* QJsonValue_NewQJsonValue2(char b)
{
	return new QJsonValue(b != 0);
}

void* QJsonValue_NewQJsonValue3(double v)
{
	return new QJsonValue(v);
}

void* QJsonValue_NewQJsonValue4(int v)
{
	return new QJsonValue(v);
}

void* QJsonValue_NewQJsonValue5(long long v)
{
	return new QJsonValue(v);
}

void* QJsonValue_NewQJsonValue6(struct QtCore_PackedString s)
{
	return new QJsonValue(QString::fromUtf8(s.data, s.len));
}

void* QJsonValue_NewQJsonValue7(void* s)
{
	return new QJsonValue(*static_cast<QLatin1String*>(s));
}

void* QJsonValue_NewQJsonValue8(char* s)
{
	return new QJsonValue(const_cast<const char*>(s));
}

void* QJsonValue_NewQJsonValue9(void* a)
{
	return new QJsonValue(*static_cast<QJsonArray*>(a));
}

void* QJsonValue_NewQJsonValue10(void* o)
{
	return new QJsonValue(*static_cast<QJsonObject*>(o));
}

void* QJsonValue_NewQJsonValue11(void* other)
{
	return new QJsonValue(*static_cast<QJsonValue*>(other));
}

void* QJsonValue_NewQJsonValue12(void* other)
{
	return new QJsonValue(*static_cast<QJsonValue*>(other));
}

int QJsonValue_ToInt(void* ptr, int defaultValue)
{
	return static_cast<QJsonValue*>(ptr)->toInt(defaultValue);
}

void* QJsonValue_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QJsonValue*>(ptr)->toVariant());
}

long long QJsonValue_Type(void* ptr)
{
	return static_cast<QJsonValue*>(ptr)->type();
}

void QJsonValue_DestroyQJsonValue(void* ptr)
{
	static_cast<QJsonValue*>(ptr)->~QJsonValue();
}

Q_DECLARE_METATYPE(QLatin1Char*)
void* QLatin1Char_NewQLatin1Char(char* c)
{
	return new QLatin1Char(*c);
}

Q_DECLARE_METATYPE(QLatin1String)
Q_DECLARE_METATYPE(QLatin1String*)
void* QLatin1String_NewQLatin1String()
{
	return new QLatin1String();
}

void* QLatin1String_NewQLatin1String2(char* str)
{
	return new QLatin1String(const_cast<const char*>(str));
}

void* QLatin1String_NewQLatin1String3(char* first, char* last)
{
	return new QLatin1String(const_cast<const char*>(first), const_cast<const char*>(last));
}

void* QLatin1String_NewQLatin1String4(char* str, int size)
{
	return new QLatin1String(const_cast<const char*>(str), size);
}

void* QLatin1String_NewQLatin1String5(void* str)
{
	return new QLatin1String(*static_cast<QByteArray*>(str));
}

struct QtCore_PackedString QLatin1String_Data(void* ptr)
{
	return QtCore_PackedString { const_cast<char*>(static_cast<QLatin1String*>(ptr)->data()), static_cast<QLatin1String*>(ptr)->size(), NULL };
}

int QLatin1String_IndexOf(void* ptr, void* str, int from, long long cs)
{
	return static_cast<QLatin1String*>(ptr)->indexOf(*static_cast<QStringView*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QLatin1String_IndexOf2(void* ptr, void* l1, int from, long long cs)
{
	return static_cast<QLatin1String*>(ptr)->indexOf(*static_cast<QLatin1String*>(l1), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QLatin1String_IndexOf3(void* ptr, void* c, int from, long long cs)
{
	return static_cast<QLatin1String*>(ptr)->indexOf(*static_cast<QChar*>(c), from, static_cast<Qt::CaseSensitivity>(cs));
}

char QLatin1String_IsEmpty(void* ptr)
{
	return static_cast<QLatin1String*>(ptr)->isEmpty();
}

void* QLatin1String_Left(void* ptr, int length)
{
	return new QLatin1String(static_cast<QLatin1String*>(ptr)->left(length));
}

void* QLatin1String_Right(void* ptr, int length)
{
	return new QLatin1String(static_cast<QLatin1String*>(ptr)->right(length));
}

int QLatin1String_Size(void* ptr)
{
	return static_cast<QLatin1String*>(ptr)->size();
}

Q_DECLARE_METATYPE(QLine)
Q_DECLARE_METATYPE(QLine*)
void* QLine_NewQLine()
{
	return new QLine();
}

void* QLine_NewQLine2(void* p1, void* p2)
{
	return new QLine(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void* QLine_NewQLine3(int x1, int y1, int x2, int y2)
{
	return new QLine(x1, y1, x2, y2);
}

void* QLine_Center(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QLine*>(ptr)->center(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

Q_DECLARE_METATYPE(QLineF)
Q_DECLARE_METATYPE(QLineF*)
void* QLineF_NewQLineF()
{
	return new QLineF();
}

void* QLineF_NewQLineF2(void* p1, void* p2)
{
	return new QLineF(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void* QLineF_NewQLineF3(double x1, double y1, double x2, double y2)
{
	return new QLineF(x1, y1, x2, y2);
}

void* QLineF_NewQLineF4(void* line)
{
	return new QLineF(*static_cast<QLine*>(line));
}

void* QLineF_Center(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QLineF*>(ptr)->center(); new QPointF(tmpValue.x(), tmpValue.y()); });
}

double QLineF_Length(void* ptr)
{
	return static_cast<QLineF*>(ptr)->length();
}

Q_DECLARE_METATYPE(QLocale)
Q_DECLARE_METATYPE(QLocale*)
void* QLocale_NewQLocale()
{
	return new QLocale();
}

void* QLocale_NewQLocale2(struct QtCore_PackedString name)
{
	return new QLocale(QString::fromUtf8(name.data, name.len));
}

void* QLocale_NewQLocale3(long long language, long long country)
{
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Country>(country));
}

void* QLocale_NewQLocale4(long long language, long long scri, long long country)
{
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Script>(scri), static_cast<QLocale::Country>(country));
}

void* QLocale_NewQLocale5(void* other)
{
	return new QLocale(*static_cast<QLocale*>(other));
}

void* QLocale_QLocale_C()
{
	return new QLocale(QLocale::c());
}

struct QtCore_PackedString QLocale_Name(void* ptr)
{
	return ({ QByteArray* tee867e = new QByteArray(static_cast<QLocale*>(ptr)->name().toUtf8()); QtCore_PackedString { const_cast<char*>(tee867e->prepend("WHITESPACE").constData()+10), tee867e->size()-10, tee867e }; });
}

int QLocale_ToInt(void* ptr, struct QtCore_PackedString s, char* ok)
{
	return static_cast<QLocale*>(ptr)->toInt(QString::fromUtf8(s.data, s.len), reinterpret_cast<bool*>(ok));
}

int QLocale_ToInt2(void* ptr, void* s, char* ok)
{
	return static_cast<QLocale*>(ptr)->toInt(*static_cast<QStringRef*>(s), reinterpret_cast<bool*>(ok));
}

int QLocale_ToInt3(void* ptr, void* s, char* ok)
{
	return static_cast<QLocale*>(ptr)->toInt(*static_cast<QStringView*>(s), reinterpret_cast<bool*>(ok));
}

struct QtCore_PackedString QLocale_ToLower(void* ptr, struct QtCore_PackedString str)
{
	return ({ QByteArray* t112fbe = new QByteArray(static_cast<QLocale*>(ptr)->toLower(QString::fromUtf8(str.data, str.len)).toUtf8()); QtCore_PackedString { const_cast<char*>(t112fbe->prepend("WHITESPACE").constData()+10), t112fbe->size()-10, t112fbe }; });
}

void QLocale_DestroyQLocale(void* ptr)
{
	static_cast<QLocale*>(ptr)->~QLocale();
}

void* QLocale___matchingLocales_atList(void* ptr, int i)
{
	return new QLocale(({QLocale tmp = static_cast<QList<QLocale>*>(ptr)->at(i); if (i == static_cast<QList<QLocale>*>(ptr)->size()-1) { static_cast<QList<QLocale>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QLocale___matchingLocales_setList(void* ptr, void* i)
{
	static_cast<QList<QLocale>*>(ptr)->append(*static_cast<QLocale*>(i));
}

void* QLocale___matchingLocales_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QLocale>();
}

long long QLocale___weekdays_atList(void* ptr, int i)
{
	return ({Qt::DayOfWeek tmp = static_cast<QList<Qt::DayOfWeek>*>(ptr)->at(i); if (i == static_cast<QList<Qt::DayOfWeek>*>(ptr)->size()-1) { static_cast<QList<Qt::DayOfWeek>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocale___weekdays_setList(void* ptr, long long i)
{
	static_cast<QList<Qt::DayOfWeek>*>(ptr)->append(static_cast<Qt::DayOfWeek>(i));
}

void* QLocale___weekdays_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Qt::DayOfWeek>();
}

long long QMetaMethod_Access(void* ptr)
{
	return static_cast<QMetaMethod*>(ptr)->access();
}

char QMetaMethod_IsValid(void* ptr)
{
	return static_cast<QMetaMethod*>(ptr)->isValid();
}

void* QMetaMethod_Name(void* ptr)
{
	return new QByteArray(static_cast<QMetaMethod*>(ptr)->name());
}

void* QMetaMethod___parameterNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QMetaMethod___parameterNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMetaMethod___parameterNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QMetaMethod___parameterTypes_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QMetaMethod___parameterTypes_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMetaMethod___parameterTypes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QMetaObject_Constructor(void* ptr, int index)
{
	return new QMetaMethod(static_cast<QMetaObject*>(ptr)->constructor(index));
}

void* QMetaObject_Method(void* ptr, int index)
{
	return new QMetaMethod(static_cast<QMetaObject*>(ptr)->method(index));
}

char QMetaProperty_IsValid(void* ptr)
{
	return static_cast<QMetaProperty*>(ptr)->isValid();
}

struct QtCore_PackedString QMetaProperty_Name(void* ptr)
{
	return QtCore_PackedString { const_cast<char*>(static_cast<QMetaProperty*>(ptr)->name()), -1, NULL };
}

void* QMetaProperty_Read(void* ptr, void* object)
{
	return new QVariant(static_cast<QMetaProperty*>(ptr)->read(static_cast<QObject*>(object)));
}

long long QMetaProperty_Type(void* ptr)
{
	return static_cast<QMetaProperty*>(ptr)->type();
}

char QMetaProperty_Write(void* ptr, void* object, void* value)
{
	return static_cast<QMetaProperty*>(ptr)->write(static_cast<QObject*>(object), *static_cast<QVariant*>(value));
}

Q_DECLARE_METATYPE(QModelIndex)
Q_DECLARE_METATYPE(QModelIndex*)
void* QModelIndex_NewQModelIndex()
{
	return new QModelIndex();
}

int QModelIndex_Column(void* ptr)
{
	return static_cast<QModelIndex*>(ptr)->column();
}

void* QModelIndex_Data(void* ptr, int role)
{
	return new QVariant(static_cast<QModelIndex*>(ptr)->data(role));
}

char QModelIndex_IsValid(void* ptr)
{
	return static_cast<QModelIndex*>(ptr)->isValid();
}

void* QModelIndex_Model(void* ptr)
{
	return const_cast<QAbstractItemModel*>(static_cast<QModelIndex*>(ptr)->model());
}

void* QModelIndex_Parent(void* ptr)
{
	return new QModelIndex(static_cast<QModelIndex*>(ptr)->parent());
}

int QModelIndex_Row(void* ptr)
{
	return static_cast<QModelIndex*>(ptr)->row();
}

class MyQObject: public QObject
{
public:
	MyQObject(QObject *parent = Q_NULLPTR) : QObject(parent) {QObject_QObject_QRegisterMetaType();};
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
	 ~MyQObject() { callbackQObject_DestroyQObject(this); };
};

Q_DECLARE_METATYPE(MyQObject*)

int QObject_QObject_QRegisterMetaType(){qRegisterMetaType<QObject*>(); return qRegisterMetaType<MyQObject*>();}

void* QObject_NewQObject(void* parent)
{
	return new MyQObject(static_cast<QObject*>(parent));
}

void QObject_ChildEvent(void* ptr, void* event)
{
	static_cast<QObject*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QObject_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		static_cast<QBuffer*>(ptr)->QBuffer::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QObject*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
	}
}

struct QtCore_PackedList QObject_Children(void* ptr)
{
	return ({ QList<QObject *>* tmpValue53f390 = new QList<QObject *>(static_cast<QObject*>(ptr)->children()); QtCore_PackedList { tmpValue53f390, tmpValue53f390->size() }; });
}

void QObject_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QObject*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QObject_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		static_cast<QBuffer*>(ptr)->QBuffer::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QObject*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QObject_CustomEvent(void* ptr, void* event)
{
	static_cast<QObject*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QObject_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		static_cast<QBuffer*>(ptr)->QBuffer::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QObject*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
	}
}

void QObject_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QObject*>(ptr), "deleteLater");
}

void QObject_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::deleteLater();
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::deleteLater();
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::deleteLater();
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		static_cast<QBuffer*>(ptr)->QBuffer::deleteLater();
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::deleteLater();
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::deleteLater();
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::deleteLater();
	} else {
		static_cast<QObject*>(ptr)->QObject::deleteLater();
	}
}

void QObject_ConnectDestroyed(void* ptr, long long t)
{
	QObject::connect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed), static_cast<Qt::ConnectionType>(t));
}

void QObject_DisconnectDestroyed(void* ptr)
{
	QObject::disconnect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed));
}

void QObject_Destroyed(void* ptr, void* obj)
{
	static_cast<QObject*>(ptr)->destroyed(static_cast<QObject*>(obj));
}

char QObject_QObject_Disconnect(void* sender, char* sign, void* receiver, char* method)
{
	return QObject::disconnect(static_cast<QObject*>(sender), const_cast<const char*>(sign), static_cast<QObject*>(receiver), const_cast<const char*>(method));
}

char QObject_QObject_Disconnect2(void* sender, void* sign, void* receiver, void* method)
{
	return QObject::disconnect(static_cast<QObject*>(sender), *static_cast<QMetaMethod*>(sign), static_cast<QObject*>(receiver), *static_cast<QMetaMethod*>(method));
}

char QObject_Disconnect3(void* ptr, char* sign, void* receiver, char* method)
{
	return static_cast<QObject*>(ptr)->disconnect(const_cast<const char*>(sign), static_cast<QObject*>(receiver), const_cast<const char*>(method));
}

char QObject_Disconnect4(void* ptr, void* receiver, char* method)
{
	return static_cast<QObject*>(ptr)->disconnect(static_cast<QObject*>(receiver), const_cast<const char*>(method));
}

void QObject_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QObject*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QObject_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		static_cast<QBuffer*>(ptr)->QBuffer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QObject*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QObject_Event(void* ptr, void* e)
{
	return static_cast<QObject*>(ptr)->event(static_cast<QEvent*>(e));
}

char QObject_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTimer*>(ptr)->QTimer::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBuffer*>(ptr)->QBuffer::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QIODevice*>(ptr)->QIODevice::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCoreApplication*>(ptr)->QCoreApplication::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QObject*>(ptr)->QObject::event(static_cast<QEvent*>(e));
	}
}

char QObject_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QObject*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QObject_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTimer*>(ptr)->QTimer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBuffer*>(ptr)->QBuffer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QIODevice*>(ptr)->QIODevice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCoreApplication*>(ptr)->QCoreApplication::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QObject*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

struct QtCore_PackedString QObject_ObjectName(void* ptr)
{
	return ({ QByteArray* te77be1 = new QByteArray(static_cast<QObject*>(ptr)->objectName().toUtf8()); QtCore_PackedString { const_cast<char*>(te77be1->prepend("WHITESPACE").constData()+10), te77be1->size()-10, te77be1 }; });
}

void QObject_ConnectObjectNameChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged), static_cast<Qt::ConnectionType>(t));
}

void QObject_DisconnectObjectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged));
}

void* QObject_Parent(void* ptr)
{
	return static_cast<QObject*>(ptr)->parent();
}

void* QObject_Property(void* ptr, char* name)
{
	return new QVariant(static_cast<QObject*>(ptr)->property(const_cast<const char*>(name)));
}

void QObject_SetObjectName(void* ptr, struct QtCore_PackedString name)
{
	static_cast<QObject*>(ptr)->setObjectName(QString::fromUtf8(name.data, name.len));
}

void QObject_TimerEvent(void* ptr, void* event)
{
	static_cast<QObject*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QObject_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QBuffer*>(static_cast<QObject*>(ptr))) {
		static_cast<QBuffer*>(ptr)->QBuffer::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QObject*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

struct QtCore_PackedString QObject_QObject_Tr(char* sourceText, char* disambiguation, int n)
{
	return ({ QByteArray* te5b32b = new QByteArray(QObject::tr(const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8()); QtCore_PackedString { const_cast<char*>(te5b32b->prepend("WHITESPACE").constData()+10), te5b32b->size()-10, te5b32b }; });
}

void QObject_DestroyQObject(void* ptr)
{
	static_cast<QObject*>(ptr)->~QObject();
}

void QObject_DestroyQObjectDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QObject_ToVariant(void* ptr)
{
	return new QVariant(QVariant::fromValue(static_cast<QObject*>(ptr)));
}

void* QObject___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QObject___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QObject___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QObject___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QObject___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QObject___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QObject___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

Q_DECLARE_METATYPE(QPersistentModelIndex*)
void* QPersistentModelIndex_NewQPersistentModelIndex2(void* index)
{
	return new QPersistentModelIndex(*static_cast<QModelIndex*>(index));
}

void* QPersistentModelIndex_NewQPersistentModelIndex3(void* other)
{
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

void* QPersistentModelIndex_NewQPersistentModelIndex4(void* other)
{
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

int QPersistentModelIndex_Column(void* ptr)
{
	return static_cast<QPersistentModelIndex*>(ptr)->column();
}

void* QPersistentModelIndex_Data(void* ptr, int role)
{
	return new QVariant(static_cast<QPersistentModelIndex*>(ptr)->data(role));
}

char QPersistentModelIndex_IsValid(void* ptr)
{
	return static_cast<QPersistentModelIndex*>(ptr)->isValid();
}

void* QPersistentModelIndex_Model(void* ptr)
{
	return const_cast<QAbstractItemModel*>(static_cast<QPersistentModelIndex*>(ptr)->model());
}

void* QPersistentModelIndex_Parent(void* ptr)
{
	return new QModelIndex(static_cast<QPersistentModelIndex*>(ptr)->parent());
}

int QPersistentModelIndex_Row(void* ptr)
{
	return static_cast<QPersistentModelIndex*>(ptr)->row();
}

Q_DECLARE_METATYPE(QPoint)
Q_DECLARE_METATYPE(QPoint*)
void* QPoint_NewQPoint()
{
	return new QPoint();
}

void* QPoint_NewQPoint2(int xpos, int ypos)
{
	return new QPoint(xpos, ypos);
}

int QPoint_X(void* ptr)
{
	return static_cast<QPoint*>(ptr)->x();
}

int QPoint_Y(void* ptr)
{
	return static_cast<QPoint*>(ptr)->y();
}

Q_DECLARE_METATYPE(QPointF)
Q_DECLARE_METATYPE(QPointF*)
void* QPointF_NewQPointF()
{
	return new QPointF();
}

void* QPointF_NewQPointF2(void* point)
{
	return new QPointF(*static_cast<QPoint*>(point));
}

void* QPointF_NewQPointF3(double xpos, double ypos)
{
	return new QPointF(xpos, ypos);
}

double QPointF_X(void* ptr)
{
	return static_cast<QPointF*>(ptr)->x();
}

double QPointF_Y(void* ptr)
{
	return static_cast<QPointF*>(ptr)->y();
}

Q_DECLARE_METATYPE(QRect)
Q_DECLARE_METATYPE(QRect*)
void* QRect_NewQRect()
{
	return new QRect();
}

void* QRect_NewQRect2(void* topLeft, void* bottomRight)
{
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QPoint*>(bottomRight));
}

void* QRect_NewQRect3(void* topLeft, void* size)
{
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QSize*>(size));
}

void* QRect_NewQRect4(int x, int y, int width, int height)
{
	return new QRect(x, y, width, height);
}

void* QRect_Center(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QRect*>(ptr)->center(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

int QRect_Height(void* ptr)
{
	return static_cast<QRect*>(ptr)->height();
}

char QRect_IsEmpty(void* ptr)
{
	return static_cast<QRect*>(ptr)->isEmpty();
}

char QRect_IsValid(void* ptr)
{
	return static_cast<QRect*>(ptr)->isValid();
}

int QRect_Left(void* ptr)
{
	return static_cast<QRect*>(ptr)->left();
}

int QRect_Right(void* ptr)
{
	return static_cast<QRect*>(ptr)->right();
}

void QRect_SetSize(void* ptr, void* size)
{
	static_cast<QRect*>(ptr)->setSize(*static_cast<QSize*>(size));
}

void QRect_SetTop(void* ptr, int y)
{
	static_cast<QRect*>(ptr)->setTop(y);
}

void* QRect_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QRect*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QRect_Top(void* ptr)
{
	return static_cast<QRect*>(ptr)->top();
}

int QRect_Width(void* ptr)
{
	return static_cast<QRect*>(ptr)->width();
}

int QRect_X(void* ptr)
{
	return static_cast<QRect*>(ptr)->x();
}

int QRect_Y(void* ptr)
{
	return static_cast<QRect*>(ptr)->y();
}

Q_DECLARE_METATYPE(QRectF)
Q_DECLARE_METATYPE(QRectF*)
void* QRectF_NewQRectF()
{
	return new QRectF();
}

void* QRectF_NewQRectF2(void* topLeft, void* size)
{
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QSizeF*>(size));
}

void* QRectF_NewQRectF3(void* topLeft, void* bottomRight)
{
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QPointF*>(bottomRight));
}

void* QRectF_NewQRectF4(double x, double y, double width, double height)
{
	return new QRectF(x, y, width, height);
}

void* QRectF_NewQRectF5(void* rectangle)
{
	return new QRectF(*static_cast<QRect*>(rectangle));
}

void* QRectF_Center(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QRectF*>(ptr)->center(); new QPointF(tmpValue.x(), tmpValue.y()); });
}

double QRectF_Height(void* ptr)
{
	return static_cast<QRectF*>(ptr)->height();
}

char QRectF_IsEmpty(void* ptr)
{
	return static_cast<QRectF*>(ptr)->isEmpty();
}

char QRectF_IsValid(void* ptr)
{
	return static_cast<QRectF*>(ptr)->isValid();
}

double QRectF_Left(void* ptr)
{
	return static_cast<QRectF*>(ptr)->left();
}

double QRectF_Right(void* ptr)
{
	return static_cast<QRectF*>(ptr)->right();
}

void QRectF_SetSize(void* ptr, void* size)
{
	static_cast<QRectF*>(ptr)->setSize(*static_cast<QSizeF*>(size));
}

void QRectF_SetTop(void* ptr, double y)
{
	static_cast<QRectF*>(ptr)->setTop(y);
}

void* QRectF_Size(void* ptr)
{
	return ({ QSizeF tmpValue = static_cast<QRectF*>(ptr)->size(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

double QRectF_Top(void* ptr)
{
	return static_cast<QRectF*>(ptr)->top();
}

double QRectF_Width(void* ptr)
{
	return static_cast<QRectF*>(ptr)->width();
}

double QRectF_X(void* ptr)
{
	return static_cast<QRectF*>(ptr)->x();
}

double QRectF_Y(void* ptr)
{
	return static_cast<QRectF*>(ptr)->y();
}

Q_DECLARE_METATYPE(QRegExp)
Q_DECLARE_METATYPE(QRegExp*)
void* QRegExp_NewQRegExp()
{
	return new QRegExp();
}

void* QRegExp_NewQRegExp2(struct QtCore_PackedString pattern, long long cs, long long syntax)
{
	return new QRegExp(QString::fromUtf8(pattern.data, pattern.len), static_cast<Qt::CaseSensitivity>(cs), static_cast<QRegExp::PatternSyntax>(syntax));
}

void* QRegExp_NewQRegExp3(void* rx)
{
	return new QRegExp(*static_cast<QRegExp*>(rx));
}

char QRegExp_IsEmpty(void* ptr)
{
	return static_cast<QRegExp*>(ptr)->isEmpty();
}

char QRegExp_IsValid(void* ptr)
{
	return static_cast<QRegExp*>(ptr)->isValid();
}

struct QtCore_PackedString QRegExp_Pattern(void* ptr)
{
	return ({ QByteArray* t70b801 = new QByteArray(static_cast<QRegExp*>(ptr)->pattern().toUtf8()); QtCore_PackedString { const_cast<char*>(t70b801->prepend("WHITESPACE").constData()+10), t70b801->size()-10, t70b801 }; });
}

int QRegExp_Pos(void* ptr, int nth)
{
	return static_cast<QRegExp*>(ptr)->pos(nth);
}

void QRegExp_DestroyQRegExp(void* ptr)
{
	static_cast<QRegExp*>(ptr)->~QRegExp();
}

Q_DECLARE_METATYPE(QRegularExpression)
Q_DECLARE_METATYPE(QRegularExpression*)
void* QRegularExpression_NewQRegularExpression()
{
	return new QRegularExpression();
}

void* QRegularExpression_NewQRegularExpression2(struct QtCore_PackedString pattern, long long options)
{
	return new QRegularExpression(QString::fromUtf8(pattern.data, pattern.len), static_cast<QRegularExpression::PatternOption>(options));
}

void* QRegularExpression_NewQRegularExpression3(void* re)
{
	return new QRegularExpression(*static_cast<QRegularExpression*>(re));
}

char QRegularExpression_IsValid(void* ptr)
{
	return static_cast<QRegularExpression*>(ptr)->isValid();
}

void* QRegularExpression_Match(void* ptr, struct QtCore_PackedString subject, int offset, long long matchType, long long matchOptions)
{
	return new QRegularExpressionMatch(static_cast<QRegularExpression*>(ptr)->match(QString::fromUtf8(subject.data, subject.len), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

void* QRegularExpression_Match2(void* ptr, void* subjectRef, int offset, long long matchType, long long matchOptions)
{
	return new QRegularExpressionMatch(static_cast<QRegularExpression*>(ptr)->match(*static_cast<QStringRef*>(subjectRef), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

struct QtCore_PackedString QRegularExpression_Pattern(void* ptr)
{
	return ({ QByteArray* t585ea0 = new QByteArray(static_cast<QRegularExpression*>(ptr)->pattern().toUtf8()); QtCore_PackedString { const_cast<char*>(t585ea0->prepend("WHITESPACE").constData()+10), t585ea0->size()-10, t585ea0 }; });
}

void QRegularExpression_DestroyQRegularExpression(void* ptr)
{
	static_cast<QRegularExpression*>(ptr)->~QRegularExpression();
}

Q_DECLARE_METATYPE(QRegularExpressionMatch)
Q_DECLARE_METATYPE(QRegularExpressionMatch*)
void* QRegularExpressionMatch_NewQRegularExpressionMatch()
{
	return new QRegularExpressionMatch();
}

void* QRegularExpressionMatch_NewQRegularExpressionMatch2(void* match)
{
	return new QRegularExpressionMatch(*static_cast<QRegularExpressionMatch*>(match));
}

char QRegularExpressionMatch_IsValid(void* ptr)
{
	return static_cast<QRegularExpressionMatch*>(ptr)->isValid();
}

void QRegularExpressionMatch_DestroyQRegularExpressionMatch(void* ptr)
{
	static_cast<QRegularExpressionMatch*>(ptr)->~QRegularExpressionMatch();
}

Q_DECLARE_METATYPE(QSize)
Q_DECLARE_METATYPE(QSize*)
void* QSize_NewQSize()
{
	return new QSize();
}

void* QSize_NewQSize2(int width, int height)
{
	return new QSize(width, height);
}

int QSize_Height(void* ptr)
{
	return static_cast<QSize*>(ptr)->height();
}

char QSize_IsEmpty(void* ptr)
{
	return static_cast<QSize*>(ptr)->isEmpty();
}

char QSize_IsValid(void* ptr)
{
	return static_cast<QSize*>(ptr)->isValid();
}

void QSize_Scale(void* ptr, int width, int height, long long mode)
{
	static_cast<QSize*>(ptr)->scale(width, height, static_cast<Qt::AspectRatioMode>(mode));
}

void QSize_Scale2(void* ptr, void* size, long long mode)
{
	static_cast<QSize*>(ptr)->scale(*static_cast<QSize*>(size), static_cast<Qt::AspectRatioMode>(mode));
}

void* QSize_Scaled(void* ptr, int width, int height, long long mode)
{
	return ({ QSize tmpValue = static_cast<QSize*>(ptr)->scaled(width, height, static_cast<Qt::AspectRatioMode>(mode)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSize_Scaled2(void* ptr, void* s, long long mode)
{
	return ({ QSize tmpValue = static_cast<QSize*>(ptr)->scaled(*static_cast<QSize*>(s), static_cast<Qt::AspectRatioMode>(mode)); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QSize_Width(void* ptr)
{
	return static_cast<QSize*>(ptr)->width();
}

Q_DECLARE_METATYPE(QSizeF)
Q_DECLARE_METATYPE(QSizeF*)
void* QSizeF_NewQSizeF()
{
	return new QSizeF();
}

void* QSizeF_NewQSizeF2(void* size)
{
	return new QSizeF(*static_cast<QSize*>(size));
}

void* QSizeF_NewQSizeF3(double width, double height)
{
	return new QSizeF(width, height);
}

double QSizeF_Height(void* ptr)
{
	return static_cast<QSizeF*>(ptr)->height();
}

char QSizeF_IsEmpty(void* ptr)
{
	return static_cast<QSizeF*>(ptr)->isEmpty();
}

char QSizeF_IsValid(void* ptr)
{
	return static_cast<QSizeF*>(ptr)->isValid();
}

void QSizeF_Scale(void* ptr, double width, double height, long long mode)
{
	static_cast<QSizeF*>(ptr)->scale(width, height, static_cast<Qt::AspectRatioMode>(mode));
}

void QSizeF_Scale2(void* ptr, void* size, long long mode)
{
	static_cast<QSizeF*>(ptr)->scale(*static_cast<QSizeF*>(size), static_cast<Qt::AspectRatioMode>(mode));
}

void* QSizeF_Scaled(void* ptr, double width, double height, long long mode)
{
	return ({ QSizeF tmpValue = static_cast<QSizeF*>(ptr)->scaled(width, height, static_cast<Qt::AspectRatioMode>(mode)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* QSizeF_Scaled2(void* ptr, void* s, long long mode)
{
	return ({ QSizeF tmpValue = static_cast<QSizeF*>(ptr)->scaled(*static_cast<QSizeF*>(s), static_cast<Qt::AspectRatioMode>(mode)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

double QSizeF_Width(void* ptr)
{
	return static_cast<QSizeF*>(ptr)->width();
}

Q_DECLARE_METATYPE(QStringRef)
Q_DECLARE_METATYPE(QStringRef*)
void* QStringRef_NewQStringRef()
{
	return new QStringRef();
}

void* QStringRef_NewQStringRef2(struct QtCore_PackedString stri, int position, int length)
{
	return new QStringRef(new QString(QString::fromUtf8(stri.data, stri.len)), position, length);
}

void* QStringRef_NewQStringRef3(struct QtCore_PackedString stri)
{
	return new QStringRef(new QString(QString::fromUtf8(stri.data, stri.len)));
}

void* QStringRef_NewQStringRef4(void* other)
{
	return new QStringRef(*static_cast<QStringRef*>(other));
}

void* QStringRef_At(void* ptr, int position)
{
	return new QChar(static_cast<QStringRef*>(ptr)->at(position));
}

void* QStringRef_Back(void* ptr)
{
	return new QChar(static_cast<QStringRef*>(ptr)->back());
}

void QStringRef_Clear(void* ptr)
{
	static_cast<QStringRef*>(ptr)->clear();
}

int QStringRef_Count(void* ptr)
{
	return static_cast<QStringRef*>(ptr)->count();
}

int QStringRef_Count2(void* ptr, struct QtCore_PackedString str, long long cs)
{
	return static_cast<QStringRef*>(ptr)->count(QString::fromUtf8(str.data, str.len), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count3(void* ptr, void* ch, long long cs)
{
	return static_cast<QStringRef*>(ptr)->count(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count4(void* ptr, void* str, long long cs)
{
	return static_cast<QStringRef*>(ptr)->count(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

void* QStringRef_Data(void* ptr)
{
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->data());
}

int QStringRef_IndexOf(void* ptr, struct QtCore_PackedString str, int from, long long cs)
{
	return static_cast<QStringRef*>(ptr)->indexOf(QString::fromUtf8(str.data, str.len), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf2(void* ptr, void* str, int from, long long cs)
{
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QStringRef*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf3(void* ptr, void* str, int from, long long cs)
{
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QStringView*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf4(void* ptr, void* ch, int from, long long cs)
{
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QChar*>(ch), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf5(void* ptr, void* str, int from, long long cs)
{
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QLatin1String*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

char QStringRef_IsEmpty(void* ptr)
{
	return static_cast<QStringRef*>(ptr)->isEmpty();
}

void* QStringRef_Left(void* ptr, int n)
{
	return new QStringRef(static_cast<QStringRef*>(ptr)->left(n));
}

int QStringRef_Length(void* ptr)
{
	return static_cast<QStringRef*>(ptr)->length();
}

int QStringRef_Position(void* ptr)
{
	return static_cast<QStringRef*>(ptr)->position();
}

void* QStringRef_Right(void* ptr, int n)
{
	return new QStringRef(static_cast<QStringRef*>(ptr)->right(n));
}

int QStringRef_Size(void* ptr)
{
	return static_cast<QStringRef*>(ptr)->size();
}

struct QtCore_PackedList QStringRef_Split(void* ptr, struct QtCore_PackedString sep, long long behavior, long long cs)
{
	return ({ QVector<QStringRef>* tmpValue06f589 = new QVector<QStringRef>(static_cast<QStringRef*>(ptr)->split(QString::fromUtf8(sep.data, sep.len), static_cast<Qt::SplitBehaviorFlags>(behavior), static_cast<Qt::CaseSensitivity>(cs))); QtCore_PackedList { tmpValue06f589, tmpValue06f589->size() }; });
}

struct QtCore_PackedList QStringRef_Split4(void* ptr, void* sep, long long behavior, long long cs)
{
	return ({ QVector<QStringRef>* tmpValue2d06d3 = new QVector<QStringRef>(static_cast<QStringRef*>(ptr)->split(*static_cast<QChar*>(sep), static_cast<Qt::SplitBehaviorFlags>(behavior), static_cast<Qt::CaseSensitivity>(cs))); QtCore_PackedList { tmpValue2d06d3, tmpValue2d06d3->size() }; });
}

struct QtCore_PackedString QStringRef_String(void* ptr)
{
	return ({ QByteArray* t75a779 = new QByteArray(static_cast<QStringRef*>(ptr)->string()->toUtf8()); QtCore_PackedString { const_cast<char*>(t75a779->prepend("WHITESPACE").constData()+10), t75a779->size()-10, t75a779 }; });
}

int QStringRef_ToInt(void* ptr, char* ok, int base)
{
	return static_cast<QStringRef*>(ptr)->toInt(reinterpret_cast<bool*>(ok), base);
}

void QStringRef_DestroyQStringRef(void* ptr)
{
	static_cast<QStringRef*>(ptr)->~QStringRef();
}

void* QStringRef___split_atList(void* ptr, int i)
{
	return new QStringRef(({QStringRef tmp = static_cast<QVector<QStringRef>*>(ptr)->at(i); if (i == static_cast<QVector<QStringRef>*>(ptr)->size()-1) { static_cast<QVector<QStringRef>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QStringRef___split_setList(void* ptr, void* i)
{
	static_cast<QVector<QStringRef>*>(ptr)->append(*static_cast<QStringRef*>(i));
}

void* QStringRef___split_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QStringRef>();
}

void* QStringRef___split_atList2(void* ptr, int i)
{
	return new QStringRef(({QStringRef tmp = static_cast<QVector<QStringRef>*>(ptr)->at(i); if (i == static_cast<QVector<QStringRef>*>(ptr)->size()-1) { static_cast<QVector<QStringRef>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QStringRef___split_setList2(void* ptr, void* i)
{
	static_cast<QVector<QStringRef>*>(ptr)->append(*static_cast<QStringRef*>(i));
}

void* QStringRef___split_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QStringRef>();
}

void* QStringRef___split_atList4(void* ptr, int i)
{
	return new QStringRef(({QStringRef tmp = static_cast<QVector<QStringRef>*>(ptr)->at(i); if (i == static_cast<QVector<QStringRef>*>(ptr)->size()-1) { static_cast<QVector<QStringRef>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QStringRef___split_setList4(void* ptr, void* i)
{
	static_cast<QVector<QStringRef>*>(ptr)->append(*static_cast<QStringRef*>(i));
}

void* QStringRef___split_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QStringRef>();
}

unsigned int QStringRef___toUcs4_atList(void* ptr, int i)
{
	return ({uint tmp = static_cast<QVector<uint>*>(ptr)->at(i); if (i == static_cast<QVector<uint>*>(ptr)->size()-1) { static_cast<QVector<uint>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QStringRef___toUcs4_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<uint>*>(ptr)->append(i);
}

void* QStringRef___toUcs4_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<uint>();
}

Q_DECLARE_METATYPE(QStringView)
Q_DECLARE_METATYPE(QStringView*)
void* QStringView_NewQStringView()
{
	return new QStringView();
}

void* QStringView_NewQStringView7(struct QtCore_PackedString str)
{
	return new QStringView(QString::fromUtf8(str.data, str.len));
}

void* QStringView_NewQStringView8(void* str)
{
	return new QStringView(*static_cast<QStringRef*>(str));
}

void* QStringView_Back(void* ptr)
{
	return new QChar(static_cast<QStringView*>(ptr)->back());
}

char QStringView_Empty(void* ptr)
{
	return static_cast<QStringView*>(ptr)->empty();
}

void* QStringView_First(void* ptr)
{
	return new QChar(static_cast<QStringView*>(ptr)->first());
}

char QStringView_IsEmpty(void* ptr)
{
	return static_cast<QStringView*>(ptr)->isEmpty();
}

void* QStringView_Last(void* ptr)
{
	return new QChar(static_cast<QStringView*>(ptr)->last());
}

int QStringView_Length(void* ptr)
{
	return static_cast<QStringView*>(ptr)->length();
}

struct QtCore_PackedList QStringView_Split(void* ptr, void* sep, long long behavior, long long cs)
{
	return ({ QList<QStringView>* tmpValue10d515 = new QList<QStringView>(static_cast<QStringView*>(ptr)->split(*static_cast<QStringView*>(sep), static_cast<Qt::SplitBehaviorFlags>(behavior), static_cast<Qt::CaseSensitivity>(cs))); QtCore_PackedList { tmpValue10d515, tmpValue10d515->size() }; });
}

struct QtCore_PackedList QStringView_Split2(void* ptr, void* sep, long long behavior, long long cs)
{
	return ({ QList<QStringView>* tmpValueeba59c = new QList<QStringView>(static_cast<QStringView*>(ptr)->split(*static_cast<QChar*>(sep), static_cast<Qt::SplitBehaviorFlags>(behavior), static_cast<Qt::CaseSensitivity>(cs))); QtCore_PackedList { tmpValueeba59c, tmpValueeba59c->size() }; });
}

struct QtCore_PackedList QStringView_Split3(void* ptr, void* sep, long long behavior)
{
	return ({ QList<QStringView>* tmpValuef05c29 = new QList<QStringView>(static_cast<QStringView*>(ptr)->split(*static_cast<QRegularExpression*>(sep), static_cast<Qt::SplitBehaviorFlags>(behavior))); QtCore_PackedList { tmpValuef05c29, tmpValuef05c29->size() }; });
}

unsigned int QStringView___convertToUcs4_atList(void* ptr, int i)
{
	return ({uint tmp = static_cast<QVector<uint>*>(ptr)->at(i); if (i == static_cast<QVector<uint>*>(ptr)->size()-1) { static_cast<QVector<uint>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QStringView___convertToUcs4_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<uint>*>(ptr)->append(i);
}

void* QStringView___convertToUcs4_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<uint>();
}

void* QStringView___split_atList(void* ptr, int i)
{
	return new QStringView(({QStringView tmp = static_cast<QList<QStringView>*>(ptr)->at(i); if (i == static_cast<QList<QStringView>*>(ptr)->size()-1) { static_cast<QList<QStringView>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStringView___split_setList(void* ptr, void* i)
{
	static_cast<QList<QStringView>*>(ptr)->append(*static_cast<QStringView*>(i));
}

void* QStringView___split_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStringView>();
}

void* QStringView___split_atList2(void* ptr, int i)
{
	return new QStringView(({QStringView tmp = static_cast<QList<QStringView>*>(ptr)->at(i); if (i == static_cast<QList<QStringView>*>(ptr)->size()-1) { static_cast<QList<QStringView>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStringView___split_setList2(void* ptr, void* i)
{
	static_cast<QList<QStringView>*>(ptr)->append(*static_cast<QStringView*>(i));
}

void* QStringView___split_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStringView>();
}

void* QStringView___split_atList3(void* ptr, int i)
{
	return new QStringView(({QStringView tmp = static_cast<QList<QStringView>*>(ptr)->at(i); if (i == static_cast<QList<QStringView>*>(ptr)->size()-1) { static_cast<QList<QStringView>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStringView___split_setList3(void* ptr, void* i)
{
	static_cast<QList<QStringView>*>(ptr)->append(*static_cast<QStringView*>(i));
}

void* QStringView___split_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStringView>();
}

unsigned int QStringView___toUcs4_atList(void* ptr, int i)
{
	return ({uint tmp = static_cast<QVector<uint>*>(ptr)->at(i); if (i == static_cast<QVector<uint>*>(ptr)->size()-1) { static_cast<QVector<uint>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QStringView___toUcs4_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<uint>*>(ptr)->append(i);
}

void* QStringView___toUcs4_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<uint>();
}

int QSysInfo_WordSize_Type()
{
	return QSysInfo::WordSize;
}

class MyQTextStream: public QTextStream
{
public:
	MyQTextStream() : QTextStream() {QTextStream_QTextStream_QRegisterMetaType();};
	MyQTextStream(QIODevice *device) : QTextStream(device) {QTextStream_QTextStream_QRegisterMetaType();};
	MyQTextStream(QString *stri, QIODevice::OpenMode openMode = QIODevice::ReadWrite) : QTextStream(stri, openMode) {QTextStream_QTextStream_QRegisterMetaType();};
	MyQTextStream(QByteArray *array, QIODevice::OpenMode openMode = QIODevice::ReadWrite) : QTextStream(array, openMode) {QTextStream_QTextStream_QRegisterMetaType();};
	MyQTextStream(const QByteArray &array, QIODevice::OpenMode openMode = QIODevice::ReadOnly) : QTextStream(array, openMode) {QTextStream_QTextStream_QRegisterMetaType();};
	 ~MyQTextStream() { callbackQTextStream_DestroyQTextStream(this); };
};

Q_DECLARE_METATYPE(QTextStream*)
Q_DECLARE_METATYPE(MyQTextStream*)

int QTextStream_QTextStream_QRegisterMetaType(){qRegisterMetaType<QTextStream*>(); return qRegisterMetaType<MyQTextStream*>();}

void* QTextStream_NewQTextStream()
{
	return new MyQTextStream();
}

void* QTextStream_NewQTextStream2(void* device)
{
	return new MyQTextStream(static_cast<QIODevice*>(device));
}

void* QTextStream_NewQTextStream4(struct QtCore_PackedString stri, long long openMode)
{
	return new MyQTextStream(new QString(QString::fromUtf8(stri.data, stri.len)), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void* QTextStream_NewQTextStream5(void* array, long long openMode)
{
	return new MyQTextStream(static_cast<QByteArray*>(array), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void* QTextStream_NewQTextStream6(void* array, long long openMode)
{
	return new MyQTextStream(*static_cast<QByteArray*>(array), static_cast<QIODevice::OpenModeFlag>(openMode));
}

void* QTextStream_Device(void* ptr)
{
	return static_cast<QTextStream*>(ptr)->device();
}

long long QTextStream_Pos(void* ptr)
{
	return static_cast<QTextStream*>(ptr)->pos();
}

struct QtCore_PackedString QTextStream_Read(void* ptr, long long maxlen)
{
	return ({ QByteArray* t1ba29b = new QByteArray(static_cast<QTextStream*>(ptr)->read(maxlen).toUtf8()); QtCore_PackedString { const_cast<char*>(t1ba29b->prepend("WHITESPACE").constData()+10), t1ba29b->size()-10, t1ba29b }; });
}

struct QtCore_PackedString QTextStream_ReadLine(void* ptr, long long maxlen)
{
	return ({ QByteArray* t7fab2e = new QByteArray(static_cast<QTextStream*>(ptr)->readLine(maxlen).toUtf8()); QtCore_PackedString { const_cast<char*>(t7fab2e->prepend("WHITESPACE").constData()+10), t7fab2e->size()-10, t7fab2e }; });
}

long long QTextStream_Status(void* ptr)
{
	return static_cast<QTextStream*>(ptr)->status();
}

struct QtCore_PackedString QTextStream_String(void* ptr)
{
	return ({ QByteArray* t586a42 = new QByteArray(static_cast<QTextStream*>(ptr)->string()->toUtf8()); QtCore_PackedString { const_cast<char*>(t586a42->prepend("WHITESPACE").constData()+10), t586a42->size()-10, t586a42 }; });
}

void QTextStream_DestroyQTextStream(void* ptr)
{
	static_cast<QTextStream*>(ptr)->~QTextStream();
}

void QTextStream_DestroyQTextStreamDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

Q_DECLARE_METATYPE(QTime)
Q_DECLARE_METATYPE(QTime*)
void* QTime_NewQTime2()
{
	return new QTime();
}

void* QTime_NewQTime3(int h, int m, int s, int ms)
{
	return new QTime(h, m, s, ms);
}

void* QTime_QTime_FromString(struct QtCore_PackedString stri, long long format)
{
	return new QTime(QTime::fromString(QString::fromUtf8(stri.data, stri.len), static_cast<Qt::DateFormat>(format)));
}

void* QTime_QTime_FromString2(struct QtCore_PackedString stri, struct QtCore_PackedString format)
{
	return new QTime(QTime::fromString(QString::fromUtf8(stri.data, stri.len), QString::fromUtf8(format.data, format.len)));
}

char QTime_IsValid(void* ptr)
{
	return static_cast<QTime*>(ptr)->isValid();
}

char QTime_QTime_IsValid2(int h, int m, int s, int ms)
{
	return QTime::isValid(h, m, s, ms);
}

Q_DECLARE_METATYPE(QTimeZone)
Q_DECLARE_METATYPE(QTimeZone*)
void* QTimeZone_NewQTimeZone()
{
	return new QTimeZone();
}

void* QTimeZone_NewQTimeZone2(void* ianaId)
{
	return new QTimeZone(*static_cast<QByteArray*>(ianaId));
}

void* QTimeZone_NewQTimeZone3(int offsetSeconds)
{
	return new QTimeZone(offsetSeconds);
}

void* QTimeZone_NewQTimeZone4(void* ianaId, int offsetSeconds, struct QtCore_PackedString name, struct QtCore_PackedString abbreviation, long long country, struct QtCore_PackedString comment)
{
	return new QTimeZone(*static_cast<QByteArray*>(ianaId), offsetSeconds, QString::fromUtf8(name.data, name.len), QString::fromUtf8(abbreviation.data, abbreviation.len), static_cast<QLocale::Country>(country), QString::fromUtf8(comment.data, comment.len));
}

void* QTimeZone_NewQTimeZone5(void* other)
{
	return new QTimeZone(*static_cast<QTimeZone*>(other));
}

struct QtCore_PackedString QTimeZone_Comment(void* ptr)
{
	return ({ QByteArray* t2b73f6 = new QByteArray(static_cast<QTimeZone*>(ptr)->comment().toUtf8()); QtCore_PackedString { const_cast<char*>(t2b73f6->prepend("WHITESPACE").constData()+10), t2b73f6->size()-10, t2b73f6 }; });
}

struct QtCore_PackedString QTimeZone_DisplayName(void* ptr, void* atDateTime, long long nameType, void* locale)
{
	return ({ QByteArray* t168cd9 = new QByteArray(static_cast<QTimeZone*>(ptr)->displayName(*static_cast<QDateTime*>(atDateTime), static_cast<QTimeZone::NameType>(nameType), *static_cast<QLocale*>(locale)).toUtf8()); QtCore_PackedString { const_cast<char*>(t168cd9->prepend("WHITESPACE").constData()+10), t168cd9->size()-10, t168cd9 }; });
}

struct QtCore_PackedString QTimeZone_DisplayName2(void* ptr, long long timeType, long long nameType, void* locale)
{
	return ({ QByteArray* tc6cf47 = new QByteArray(static_cast<QTimeZone*>(ptr)->displayName(static_cast<QTimeZone::TimeType>(timeType), static_cast<QTimeZone::NameType>(nameType), *static_cast<QLocale*>(locale)).toUtf8()); QtCore_PackedString { const_cast<char*>(tc6cf47->prepend("WHITESPACE").constData()+10), tc6cf47->size()-10, tc6cf47 }; });
}

void* QTimeZone_Id(void* ptr)
{
	return new QByteArray(static_cast<QTimeZone*>(ptr)->id());
}

char QTimeZone_IsValid(void* ptr)
{
	return static_cast<QTimeZone*>(ptr)->isValid();
}

void QTimeZone_DestroyQTimeZone(void* ptr)
{
	static_cast<QTimeZone*>(ptr)->~QTimeZone();
}

void* QTimeZone___availableTimeZoneIds_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___availableTimeZoneIds_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___availableTimeZoneIds_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___availableTimeZoneIds_atList2(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___availableTimeZoneIds_setList2(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___availableTimeZoneIds_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___availableTimeZoneIds_atList3(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___availableTimeZoneIds_setList3(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___availableTimeZoneIds_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___windowsIdToIanaIds_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___windowsIdToIanaIds_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___windowsIdToIanaIds_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___windowsIdToIanaIds_atList2(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___windowsIdToIanaIds_setList2(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___windowsIdToIanaIds_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

class MyQTimer: public QTimer
{
public:
	MyQTimer(QObject *parent = Q_NULLPTR) : QTimer(parent) {QTimer_QTimer_QRegisterMetaType();};
	void start(int msec) { callbackQTimer_Start(this, msec); };
	void start() { callbackQTimer_Start2(this); };
	void stop() { callbackQTimer_Stop(this); };
	void timerEvent(QTimerEvent * e) { callbackQObject_TimerEvent(this, e); };
	 ~MyQTimer() { callbackQTimer_DestroyQTimer(this); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QTimer*)
Q_DECLARE_METATYPE(MyQTimer*)

int QTimer_QTimer_QRegisterMetaType(){qRegisterMetaType<QTimer*>(); return qRegisterMetaType<MyQTimer*>();}

void* QTimer_NewQTimer(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QWindow*>(parent));
	} else {
		return new MyQTimer(static_cast<QObject*>(parent));
	}
}

void QTimer_Start(void* ptr, int msec)
{
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "start", Q_ARG(int, msec));
}

void QTimer_StartDefault(void* ptr, int msec)
{
		static_cast<QTimer*>(ptr)->QTimer::start(msec);
}

void QTimer_Start2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "start");
}

void QTimer_Start2Default(void* ptr)
{
		static_cast<QTimer*>(ptr)->QTimer::start();
}

void QTimer_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "stop");
}

void QTimer_StopDefault(void* ptr)
{
		static_cast<QTimer*>(ptr)->QTimer::stop();
}

void QTimer_DestroyQTimer(void* ptr)
{
	static_cast<QTimer*>(ptr)->~QTimer();
}

void QTimer_DestroyQTimerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQTimerEvent: public QTimerEvent
{
public:
	MyQTimerEvent(int timerId) : QTimerEvent(timerId) {QTimerEvent_QTimerEvent_QRegisterMetaType();};
};

Q_DECLARE_METATYPE(QTimerEvent*)
Q_DECLARE_METATYPE(MyQTimerEvent*)

int QTimerEvent_QTimerEvent_QRegisterMetaType(){qRegisterMetaType<QTimerEvent*>(); return qRegisterMetaType<MyQTimerEvent*>();}

void* QTimerEvent_NewQTimerEvent(int timerId)
{
	return new MyQTimerEvent(timerId);
}

Q_DECLARE_METATYPE(QUrl)
Q_DECLARE_METATYPE(QUrl*)
void* QUrl_NewQUrl()
{
	return new QUrl();
}

void* QUrl_NewQUrl2(void* other)
{
	return new QUrl(*static_cast<QUrl*>(other));
}

void* QUrl_NewQUrl3(struct QtCore_PackedString url, long long parsingMode)
{
	return new QUrl(QString::fromUtf8(url.data, url.len), static_cast<QUrl::ParsingMode>(parsingMode));
}

void* QUrl_NewQUrl4(void* other)
{
	return new QUrl(*static_cast<QUrl*>(other));
}

void QUrl_Clear(void* ptr)
{
	static_cast<QUrl*>(ptr)->clear();
}

struct QtCore_PackedString QUrl_FileName(void* ptr, long long options)
{
	return ({ QByteArray* t4c468f = new QByteArray(static_cast<QUrl*>(ptr)->fileName(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8()); QtCore_PackedString { const_cast<char*>(t4c468f->prepend("WHITESPACE").constData()+10), t4c468f->size()-10, t4c468f }; });
}

struct QtCore_PackedString QUrl_Fragment(void* ptr, long long options)
{
	return ({ QByteArray* t754185 = new QByteArray(static_cast<QUrl*>(ptr)->fragment(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8()); QtCore_PackedString { const_cast<char*>(t754185->prepend("WHITESPACE").constData()+10), t754185->size()-10, t754185 }; });
}

char QUrl_IsEmpty(void* ptr)
{
	return static_cast<QUrl*>(ptr)->isEmpty();
}

char QUrl_IsValid(void* ptr)
{
	return static_cast<QUrl*>(ptr)->isValid();
}

struct QtCore_PackedString QUrl_Password(void* ptr, long long options)
{
	return ({ QByteArray* t55f068 = new QByteArray(static_cast<QUrl*>(ptr)->password(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8()); QtCore_PackedString { const_cast<char*>(t55f068->prepend("WHITESPACE").constData()+10), t55f068->size()-10, t55f068 }; });
}

struct QtCore_PackedString QUrl_Path(void* ptr, long long options)
{
	return ({ QByteArray* t70ef65 = new QByteArray(static_cast<QUrl*>(ptr)->path(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8()); QtCore_PackedString { const_cast<char*>(t70ef65->prepend("WHITESPACE").constData()+10), t70ef65->size()-10, t70ef65 }; });
}

struct QtCore_PackedString QUrl_Query(void* ptr, long long options)
{
	return ({ QByteArray* t911c73 = new QByteArray(static_cast<QUrl*>(ptr)->query(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8()); QtCore_PackedString { const_cast<char*>(t911c73->prepend("WHITESPACE").constData()+10), t911c73->size()-10, t911c73 }; });
}

struct QtCore_PackedString QUrl_Url(void* ptr, long long options)
{
	return ({ QByteArray* t2ea726 = new QByteArray(static_cast<QUrl*>(ptr)->url(static_cast<QUrl::UrlFormattingOption>(options)).toUtf8()); QtCore_PackedString { const_cast<char*>(t2ea726->prepend("WHITESPACE").constData()+10), t2ea726->size()-10, t2ea726 }; });
}

struct QtCore_PackedString QUrl_UserName(void* ptr, long long options)
{
	return ({ QByteArray* tb9c277 = new QByteArray(static_cast<QUrl*>(ptr)->userName(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8()); QtCore_PackedString { const_cast<char*>(tb9c277->prepend("WHITESPACE").constData()+10), tb9c277->size()-10, tb9c277 }; });
}

void QUrl_DestroyQUrl(void* ptr)
{
	static_cast<QUrl*>(ptr)->~QUrl();
}

void* QUrl___allEncodedQueryItemValues_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QUrl___allEncodedQueryItemValues_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QUrl___allEncodedQueryItemValues_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QUrl___fromStringList_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QUrl___fromStringList_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QUrl___fromStringList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QUrl___toStringList_urls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QUrl___toStringList_urls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QUrl___toStringList_urls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

Q_DECLARE_METATYPE(QUuid)
Q_DECLARE_METATYPE(QUuid*)
void* QUuid_NewQUuid2()
{
	return new QUuid();
}

void* QUuid_NewQUuid3(unsigned int l, unsigned short w1, unsigned short w2, char* b1, char* b2, char* b3, char* b4, char* b5, char* b6, char* b7, char* b8)
{
	return new QUuid(l, w1, w2, *static_cast<uchar*>(static_cast<void*>(b1)), *static_cast<uchar*>(static_cast<void*>(b2)), *static_cast<uchar*>(static_cast<void*>(b3)), *static_cast<uchar*>(static_cast<void*>(b4)), *static_cast<uchar*>(static_cast<void*>(b5)), *static_cast<uchar*>(static_cast<void*>(b6)), *static_cast<uchar*>(static_cast<void*>(b7)), *static_cast<uchar*>(static_cast<void*>(b8)));
}

void* QUuid_NewQUuid4(struct QtCore_PackedString text)
{
	return new QUuid(QString::fromUtf8(text.data, text.len));
}

void* QUuid_NewQUuid(void* text)
{
	return new QUuid(*static_cast<QByteArray*>(text));
}

void* QUuid_QUuid_FromString(void* text)
{
	return new QUuid(QUuid::fromString(*static_cast<QStringView*>(text)));
}

void* QUuid_QUuid_FromString2(void* text)
{
	return new QUuid(QUuid::fromString(*static_cast<QLatin1String*>(text)));
}

long long QUuid_Variant(void* ptr)
{
	return static_cast<QUuid*>(ptr)->variant();
}

long long QUuid_Version(void* ptr)
{
	return static_cast<QUuid*>(ptr)->version();
}

Q_DECLARE_METATYPE(QVariant*)
void* QVariant_NewQVariant()
{
	return new QVariant();
}

void* QVariant_NewQVariant2(long long ty)
{
	return new QVariant(static_cast<QVariant::Type>(ty));
}

void* QVariant_NewQVariant3(int typeId, void* copy)
{
	return new QVariant(typeId, copy);
}

void* QVariant_NewQVariant4(void* s)
{
	return new QVariant(*static_cast<QDataStream*>(s));
}

void* QVariant_NewQVariant5(int val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant6(unsigned int val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant7(long long val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant8(unsigned long long val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant9(char val)
{
	return new QVariant(val != 0);
}

void* QVariant_NewQVariant10(double val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant11(float val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant12(char* val)
{
	return new QVariant(const_cast<const char*>(val));
}

void* QVariant_NewQVariant13(void* val)
{
	return new QVariant(*static_cast<QByteArray*>(val));
}

void* QVariant_NewQVariant14(void* val)
{
	return new QVariant(*static_cast<QBitArray*>(val));
}

void* QVariant_NewQVariant15(struct QtCore_PackedString val)
{
	return new QVariant(QString::fromUtf8(val.data, val.len));
}

void* QVariant_NewQVariant16(void* val)
{
	return new QVariant(*static_cast<QLatin1String*>(val));
}

void* QVariant_NewQVariant17(struct QtCore_PackedString val)
{
	return new QVariant(QString::fromUtf8(val.data, val.len).split("¡¦!", QString::SkipEmptyParts));
}

void* QVariant_NewQVariant18(void* c)
{
	return new QVariant(*static_cast<QChar*>(c));
}

void* QVariant_NewQVariant19(void* val)
{
	return new QVariant(*static_cast<QDate*>(val));
}

void* QVariant_NewQVariant20(void* val)
{
	return new QVariant(*static_cast<QTime*>(val));
}

void* QVariant_NewQVariant21(void* val)
{
	return new QVariant(*static_cast<QDateTime*>(val));
}

void* QVariant_NewQVariant22(void* val)
{
	return new QVariant(({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(val); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* QVariant_NewQVariant23(void* val)
{
	return new QVariant(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(val); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void* QVariant_NewQVariant24(void* val)
{
	return new QVariant(({ QHash<QString, QVariant>* tmpP = static_cast<QHash<QString, QVariant>*>(val); QHash<QString, QVariant> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }));
}

void* QVariant_NewQVariant25(void* val)
{
	return new QVariant(*static_cast<QSize*>(val));
}

void* QVariant_NewQVariant26(void* val)
{
	return new QVariant(*static_cast<QSizeF*>(val));
}

void* QVariant_NewQVariant27(void* val)
{
	return new QVariant(*static_cast<QPoint*>(val));
}

void* QVariant_NewQVariant28(void* val)
{
	return new QVariant(*static_cast<QPointF*>(val));
}

void* QVariant_NewQVariant29(void* val)
{
	return new QVariant(*static_cast<QLine*>(val));
}

void* QVariant_NewQVariant30(void* val)
{
	return new QVariant(*static_cast<QLineF*>(val));
}

void* QVariant_NewQVariant31(void* val)
{
	return new QVariant(*static_cast<QRect*>(val));
}

void* QVariant_NewQVariant32(void* val)
{
	return new QVariant(*static_cast<QRectF*>(val));
}

void* QVariant_NewQVariant33(void* l)
{
	return new QVariant(*static_cast<QLocale*>(l));
}

void* QVariant_NewQVariant34(void* regExp)
{
	return new QVariant(*static_cast<QRegExp*>(regExp));
}

void* QVariant_NewQVariant35(void* re)
{
	return new QVariant(*static_cast<QRegularExpression*>(re));
}

void* QVariant_NewQVariant36(void* val)
{
	return new QVariant(*static_cast<QEasingCurve*>(val));
}

void* QVariant_NewQVariant37(void* val)
{
	return new QVariant(*static_cast<QUuid*>(val));
}

void* QVariant_NewQVariant38(void* val)
{
	return new QVariant(*static_cast<QUrl*>(val));
}

void* QVariant_NewQVariant39(void* val)
{
	return new QVariant(*static_cast<QJsonValue*>(val));
}

void* QVariant_NewQVariant40(void* val)
{
	return new QVariant(*static_cast<QJsonObject*>(val));
}

void* QVariant_NewQVariant41(void* val)
{
	return new QVariant(*static_cast<QJsonArray*>(val));
}

void* QVariant_NewQVariant42(void* val)
{
	return new QVariant(*static_cast<QJsonDocument*>(val));
}

void* QVariant_NewQVariant43(void* val)
{
	return new QVariant(*static_cast<QModelIndex*>(val));
}

void* QVariant_NewQVariant44(void* val)
{
	return new QVariant(*static_cast<QPersistentModelIndex*>(val));
}

void* QVariant_NewQVariant45(void* other)
{
	return new QVariant(*static_cast<QVariant*>(other));
}

char QVariant_CanConvert(void* ptr, int targetTypeId)
{
	return static_cast<QVariant*>(ptr)->canConvert(targetTypeId);
}

void QVariant_Clear(void* ptr)
{
	static_cast<QVariant*>(ptr)->clear();
}

char QVariant_IsNull(void* ptr)
{
	return static_cast<QVariant*>(ptr)->isNull();
}

char QVariant_IsValid(void* ptr)
{
	return static_cast<QVariant*>(ptr)->isValid();
}

char QVariant_ToBool(void* ptr)
{
	return static_cast<QVariant*>(ptr)->toBool();
}

double QVariant_ToDouble(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toDouble(reinterpret_cast<bool*>(ok));
}

struct QtCore_PackedList QVariant_ToHash(void* ptr)
{
	return ({ QHash<QString, QVariant>* tmpValue00701e = new QHash<QString, QVariant>(static_cast<QVariant*>(ptr)->toHash()); QtCore_PackedList { tmpValue00701e, tmpValue00701e->size() }; });
}

int QVariant_ToInt(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toInt(reinterpret_cast<bool*>(ok));
}

struct QtCore_PackedList QVariant_ToList(void* ptr)
{
	return ({ QList<QVariant>* tmpValue8f6950 = new QList<QVariant>(static_cast<QVariant*>(ptr)->toList()); QtCore_PackedList { tmpValue8f6950, tmpValue8f6950->size() }; });
}

long long QVariant_ToLongLong(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toLongLong(reinterpret_cast<bool*>(ok));
}

struct QtCore_PackedList QVariant_ToMap(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue1e0d76 = new QMap<QString, QVariant>(static_cast<QVariant*>(ptr)->toMap()); QtCore_PackedList { tmpValue1e0d76, tmpValue1e0d76->size() }; });
}

struct QtCore_PackedString QVariant_ToString(void* ptr)
{
	return ({ QByteArray* tf9e1e4 = new QByteArray(static_cast<QVariant*>(ptr)->toString().toUtf8()); QtCore_PackedString { const_cast<char*>(tf9e1e4->prepend("WHITESPACE").constData()+10), tf9e1e4->size()-10, tf9e1e4 }; });
}

struct QtCore_PackedString QVariant_ToStringList(void* ptr)
{
	return ({ QByteArray* tf99cb6 = new QByteArray(static_cast<QVariant*>(ptr)->toStringList().join("¡¦!").toUtf8()); QtCore_PackedString { const_cast<char*>(tf99cb6->prepend("WHITESPACE").constData()+10), tf99cb6->size()-10, tf99cb6 }; });
}

unsigned int QVariant_ToUInt(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toUInt(reinterpret_cast<bool*>(ok));
}

unsigned long long QVariant_ToULongLong(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toULongLong(reinterpret_cast<bool*>(ok));
}

long long QVariant_Type(void* ptr)
{
	return static_cast<QVariant*>(ptr)->type();
}

void QVariant_DestroyQVariant(void* ptr)
{
	static_cast<QVariant*>(ptr)->~QVariant();
}

void* QVariant_ToImage(void* ptr)
{
	return new QImage(qvariant_cast<QImage>(*static_cast<QVariant*>(ptr)));
}

void* QVariant___QVariant_val_atList22(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVariant___QVariant_val_setList22(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QVariant___QVariant_val_newList22(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QVariant___QVariant_val_atList23(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVariant___QVariant_val_setList23(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___QVariant_val_newList23(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtCore_PackedList QVariant___QVariant_val_keyList23(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QVariant___QVariant_val_atList24(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QVariant___QVariant_val_setList24(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___QVariant_val_newList24(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QVariant___QVariant_val_keyList24(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QVariant___toHash_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QVariant___toHash_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___toHash_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QVariant___toHash_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QVariant___toList_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVariant___toList_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QVariant___toList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QVariant___toMap_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVariant___toMap_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___toMap_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtCore_PackedList QVariant___toMap_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtCore_PackedString QVariant_____QVariant_val_keyList_atList23(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVariant_____QVariant_val_keyList_setList23(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____QVariant_val_keyList_newList23(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QVariant_____QVariant_val_keyList_atList24(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVariant_____QVariant_val_keyList_setList24(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____QVariant_val_keyList_newList24(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QVariant_____toHash_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVariant_____toHash_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____toHash_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QVariant_____toMap_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray* t94aa5e = new QByteArray(({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8()); QtCore_PackedString { const_cast<char*>(t94aa5e->prepend("WHITESPACE").constData()+10), t94aa5e->size()-10, t94aa5e }; });
}

void QVariant_____toMap_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____toMap_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

int Qt_LastGestureType_Type()
{
	return Qt::LastGestureType;
}

