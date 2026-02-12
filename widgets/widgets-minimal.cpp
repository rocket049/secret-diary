// +build minimal

#define protected public
#define private public

#include "widgets-minimal.h"
#include "_cgo_export.h"

#include <QAbstractButton>
#include <QAbstractItemModel>
#include <QAbstractItemView>
#include <QAbstractScrollArea>
#include <QAbstractSlider>
#include <QAction>
#include <QApplication>
#include <QBoxLayout>
#include <QButtonGroup>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QColor>
#include <QColorDialog>
#include <QComboBox>
#include <QContextMenuEvent>
#include <QCursor>
#include <QDial>
#include <QDialog>
#include <QDialogButtonBox>
#include <QDir>
#include <QDirModel>
#include <QDockWidget>
#include <QEvent>
#include <QFileDialog>
#include <QFileSystemModel>
#include <QFont>
#include <QFontDialog>
#include <QFontMetrics>
#include <QFrame>
#include <QGraphicsItem>
#include <QGraphicsItemGroup>
#include <QGraphicsLayout>
#include <QGraphicsLayoutItem>
#include <QGraphicsObject>
#include <QGraphicsSceneContextMenuEvent>
#include <QGraphicsSceneEvent>
#include <QGraphicsSceneMouseEvent>
#include <QGraphicsSceneResizeEvent>
#include <QGraphicsTransform>
#include <QGraphicsWidget>
#include <QGridLayout>
#include <QHBoxLayout>
#include <QHeaderView>
#include <QHideEvent>
#include <QIcon>
#include <QInputDialog>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QKeyEvent>
#include <QKeySequence>
#include <QLabel>
#include <QLayout>
#include <QLayoutItem>
#include <QLineEdit>
#include <QMainWindow>
#include <QMatrix4x4>
#include <QMenu>
#include <QMenuBar>
#include <QMessageBox>
#include <QMetaMethod>
#include <QMetaObject>
#include <QModelIndex>
#include <QMouseEvent>
#include <QObject>
#include <QPagedPaintDevice>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPixmap>
#include <QPoint>
#include <QPointF>
#include <QPushButton>
#include <QRect>
#include <QRectF>
#include <QRegExp>
#include <QRegion>
#include <QRegularExpression>
#include <QResizeEvent>
#include <QScreen>
#include <QScrollArea>
#include <QScrollBar>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QSizePolicy>
#include <QSpacerItem>
#include <QSplitter>
#include <QStatusBar>
#include <QString>
#include <QStyle>
#include <QStyleHintReturn>
#include <QStyleOption>
#include <QStyleOptionComplex>
#include <QStyleOptionGraphicsItem>
#include <QTextCharFormat>
#include <QTextCursor>
#include <QTextDocument>
#include <QTextEdit>
#include <QTimerEvent>
#include <QToolBar>
#include <QTransform>
#include <QTreeView>
#include <QUrl>
#include <QVBoxLayout>
#include <QValidator>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>
#include <QStringList>

class MyQAbstractButton: public QAbstractButton
{
public:
	MyQAbstractButton(QWidget *parent = Q_NULLPTR) : QAbstractButton(parent) {QAbstractButton_QAbstractButton_QRegisterMetaType();};
	void click() { callbackQAbstractButton_Click(this); };
	void Signal_Clicked(bool checked) { callbackQAbstractButton_Clicked(this, checked); };
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void keyReleaseEvent(QKeyEvent * e) { callbackQWidget_KeyReleaseEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	void paintEvent(QPaintEvent * e) { callbackQAbstractButton_PaintEvent(this, e); };
	void setChecked(bool vbo) { callbackQAbstractButton_SetChecked(this, vbo); };
	void timerEvent(QTimerEvent * e) { callbackQWidget_TimerEvent(this, e); };
	 ~MyQAbstractButton() { callbackQAbstractButton_DestroyQAbstractButton(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QAbstractButton*)
Q_DECLARE_METATYPE(MyQAbstractButton*)

int QAbstractButton_QAbstractButton_QRegisterMetaType(){qRegisterMetaType<QAbstractButton*>(); return qRegisterMetaType<MyQAbstractButton*>();}

void* QAbstractButton_NewQAbstractButton(void* parent)
{
		return new MyQAbstractButton(static_cast<QWidget*>(parent));
}

void QAbstractButton_Click(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "click");
}

void QAbstractButton_ClickDefault(void* ptr)
{
	if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::click();
	} else {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::click();
	}
}

void QAbstractButton_ConnectClicked(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Clicked), static_cast<Qt::ConnectionType>(t));
}

void QAbstractButton_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Clicked));
}

void QAbstractButton_Clicked(void* ptr, char checked)
{
	static_cast<QAbstractButton*>(ptr)->clicked(checked != 0);
}

void* QAbstractButton_Group(void* ptr)
{
	return static_cast<QAbstractButton*>(ptr)->group();
}

void* QAbstractButton_Icon(void* ptr)
{
	return new QIcon(static_cast<QAbstractButton*>(ptr)->icon());
}

char QAbstractButton_IsChecked(void* ptr)
{
	return static_cast<QAbstractButton*>(ptr)->isChecked();
}

void QAbstractButton_PaintEvent(void* ptr, void* e)
{
	static_cast<QAbstractButton*>(ptr)->paintEvent(static_cast<QPaintEvent*>(e));
}

void QAbstractButton_SetCheckable(void* ptr, char vbo)
{
	static_cast<QAbstractButton*>(ptr)->setCheckable(vbo != 0);
}

void QAbstractButton_SetChecked(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "setChecked", Q_ARG(bool, vbo != 0));
}

void QAbstractButton_SetCheckedDefault(void* ptr, char vbo)
{
	if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::setChecked(vbo != 0);
	} else {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::setChecked(vbo != 0);
	}
}

void QAbstractButton_SetIcon(void* ptr, void* icon)
{
	static_cast<QAbstractButton*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QAbstractButton_SetShortcut(void* ptr, void* key)
{
	static_cast<QAbstractButton*>(ptr)->setShortcut(*static_cast<QKeySequence*>(key));
}

void QAbstractButton_SetText(void* ptr, struct QtWidgets_PackedString text)
{
	static_cast<QAbstractButton*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

void* QAbstractButton_Shortcut(void* ptr)
{
	return new QKeySequence(static_cast<QAbstractButton*>(ptr)->shortcut());
}

struct QtWidgets_PackedString QAbstractButton_Text(void* ptr)
{
	return ({ QByteArray* t8ae7db = new QByteArray(static_cast<QAbstractButton*>(ptr)->text().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t8ae7db->prepend("WHITESPACE").constData()+10), t8ae7db->size()-10, t8ae7db }; });
}

void QAbstractButton_DestroyQAbstractButton(void* ptr)
{
	static_cast<QAbstractButton*>(ptr)->~QAbstractButton();
}

void QAbstractButton_DestroyQAbstractButtonDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQAbstractItemView: public QAbstractItemView
{
public:
	MyQAbstractItemView(QWidget *parent = Q_NULLPTR) : QAbstractItemView(parent) {QAbstractItemView_QAbstractItemView_QRegisterMetaType();};
	void Signal_Activated(const QModelIndex & index) { callbackQAbstractItemView_Activated(this, const_cast<QModelIndex*>(&index)); };
	void clearSelection() { callbackQAbstractItemView_ClearSelection(this); };
	void Signal_Clicked(const QModelIndex & index) { callbackQAbstractItemView_Clicked(this, const_cast<QModelIndex*>(&index)); };
	void edit(const QModelIndex & index) { callbackQAbstractItemView_Edit(this, const_cast<QModelIndex*>(&index)); };
	bool edit(const QModelIndex & index, QAbstractItemView::EditTrigger trigger, QEvent * event) { return callbackQAbstractItemView_Edit2(this, const_cast<QModelIndex*>(&index), trigger, event) != 0; };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	bool eventFilter(QObject * object, QEvent * event) { return callbackQWidget_EventFilter(this, object, event) != 0; };
	int horizontalOffset() const { return callbackQAbstractItemView_HorizontalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	QModelIndex indexAt(const QPoint & point) const { return *static_cast<QModelIndex*>(callbackQAbstractItemView_IndexAt(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPoint*>(&point))); };
	bool isIndexHidden(const QModelIndex & index) const { return callbackQAbstractItemView_IsIndexHidden(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index)) != 0; };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, Qt::KeyboardModifiers modifiers) { return *static_cast<QModelIndex*>(callbackQAbstractItemView_MoveCursor(this, cursorAction, modifiers)); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void scrollTo(const QModelIndex & index, QAbstractItemView::ScrollHint hint) { callbackQAbstractItemView_ScrollTo(this, const_cast<QModelIndex*>(&index), hint); };
	void selectionChanged(const QItemSelection & selected, const QItemSelection & deselected) { callbackQAbstractItemView_SelectionChanged(this, const_cast<QItemSelection*>(&selected), const_cast<QItemSelection*>(&deselected)); };
	void setCurrentIndex(const QModelIndex & index) { callbackQAbstractItemView_SetCurrentIndex(this, const_cast<QModelIndex*>(&index)); };
	void setModel(QAbstractItemModel * model) { callbackQAbstractItemView_SetModel(this, model); };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags flags) { callbackQAbstractItemView_SetSelection(this, const_cast<QRect*>(&rect), flags); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	void update(const QModelIndex & index) { callbackQAbstractItemView_Update(this, const_cast<QModelIndex*>(&index)); };
	int verticalOffset() const { return callbackQAbstractItemView_VerticalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	QRect visualRect(const QModelIndex & index) const { return *static_cast<QRect*>(callbackQAbstractItemView_VisualRect(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QRegion visualRegionForSelection(const QItemSelection & selection) const { return *static_cast<QRegion*>(callbackQAbstractItemView_VisualRegionForSelection(const_cast<void*>(static_cast<const void*>(this)), const_cast<QItemSelection*>(&selection))); };
	 ~MyQAbstractItemView() { callbackQAbstractItemView_DestroyQAbstractItemView(this); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QAbstractItemView*)
Q_DECLARE_METATYPE(MyQAbstractItemView*)

int QAbstractItemView_QAbstractItemView_QRegisterMetaType(){qRegisterMetaType<QAbstractItemView*>(); return qRegisterMetaType<MyQAbstractItemView*>();}

void* QAbstractItemView_NewQAbstractItemView(void* parent)
{
		return new MyQAbstractItemView(static_cast<QWidget*>(parent));
}

void QAbstractItemView_ConnectActivated(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::activated), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Activated), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemView_DisconnectActivated(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::activated), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Activated));
}

void QAbstractItemView_Activated(void* ptr, void* index)
{
	static_cast<QAbstractItemView*>(ptr)->activated(*static_cast<QModelIndex*>(index));
}

void QAbstractItemView_ClearSelection(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "clearSelection");
}

void QAbstractItemView_ClearSelectionDefault(void* ptr)
{
	if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::clearSelection();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::clearSelection();
	} else {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::clearSelection();
	}
}

void QAbstractItemView_ConnectClicked(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::clicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Clicked), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemView_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::clicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Clicked));
}

void QAbstractItemView_Clicked(void* ptr, void* index)
{
	static_cast<QAbstractItemView*>(ptr)->clicked(*static_cast<QModelIndex*>(index));
}

void* QAbstractItemView_CurrentIndex(void* ptr)
{
	return new QModelIndex(static_cast<QAbstractItemView*>(ptr)->currentIndex());
}

void QAbstractItemView_Edit(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "edit", Q_ARG(const QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_EditDefault(void* ptr, void* index)
{
	if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::edit(*static_cast<QModelIndex*>(index));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::edit(*static_cast<QModelIndex*>(index));
	} else {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::edit(*static_cast<QModelIndex*>(index));
	}
}

char QAbstractItemView_Edit2(void* ptr, void* index, long long trigger, void* event)
{
	return static_cast<QAbstractItemView*>(ptr)->edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
}

char QAbstractItemView_Edit2Default(void* ptr, void* index, long long trigger, void* event)
{
	if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTreeView*>(ptr)->QTreeView::edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHeaderView*>(ptr)->QHeaderView::edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
	} else {
		return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
	}
}

int QAbstractItemView_HorizontalOffset(void* ptr)
{
	return static_cast<QAbstractItemView*>(ptr)->horizontalOffset();
}

void* QAbstractItemView_IndexAt(void* ptr, void* point)
{
	return new QModelIndex(static_cast<QAbstractItemView*>(ptr)->indexAt(*static_cast<QPoint*>(point)));
}

char QAbstractItemView_IsIndexHidden(void* ptr, void* index)
{
	return static_cast<QAbstractItemView*>(ptr)->isIndexHidden(*static_cast<QModelIndex*>(index));
}

void* QAbstractItemView_Model(void* ptr)
{
	return static_cast<QAbstractItemView*>(ptr)->model();
}

void* QAbstractItemView_MoveCursor(void* ptr, long long cursorAction, long long modifiers)
{
	return new QModelIndex(static_cast<QAbstractItemView*>(ptr)->moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

void* QAbstractItemView_RootIndex(void* ptr)
{
	return new QModelIndex(static_cast<QAbstractItemView*>(ptr)->rootIndex());
}

void QAbstractItemView_ScrollTo(void* ptr, void* index, long long hint)
{
	static_cast<QAbstractItemView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QAbstractItemView_SelectionChanged(void* ptr, void* selected, void* deselected)
{
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "selectionChanged", Q_ARG(const QItemSelection, *static_cast<QItemSelection*>(selected)), Q_ARG(const QItemSelection, *static_cast<QItemSelection*>(deselected)));
}

void QAbstractItemView_SelectionChangedDefault(void* ptr, void* selected, void* deselected)
{
	if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
	} else {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
	}
}

long long QAbstractItemView_SelectionMode(void* ptr)
{
	return static_cast<QAbstractItemView*>(ptr)->selectionMode();
}

void QAbstractItemView_SetAutoScroll(void* ptr, char enable)
{
	static_cast<QAbstractItemView*>(ptr)->setAutoScroll(enable != 0);
}

void QAbstractItemView_SetCurrentIndex(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "setCurrentIndex", Q_ARG(const QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_SetCurrentIndexDefault(void* ptr, void* index)
{
	if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::setCurrentIndex(*static_cast<QModelIndex*>(index));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::setCurrentIndex(*static_cast<QModelIndex*>(index));
	} else {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::setCurrentIndex(*static_cast<QModelIndex*>(index));
	}
}

void QAbstractItemView_SetModel(void* ptr, void* model)
{
	static_cast<QAbstractItemView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QAbstractItemView_SetModelDefault(void* ptr, void* model)
{
	if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::setModel(static_cast<QAbstractItemModel*>(model));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::setModel(static_cast<QAbstractItemModel*>(model));
	} else {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::setModel(static_cast<QAbstractItemModel*>(model));
	}
}

void QAbstractItemView_SetSelection(void* ptr, void* rect, long long flags)
{
	static_cast<QAbstractItemView*>(ptr)->setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(flags));
}

void QAbstractItemView_SetSelectionMode(void* ptr, long long mode)
{
	static_cast<QAbstractItemView*>(ptr)->setSelectionMode(static_cast<QAbstractItemView::SelectionMode>(mode));
}

void QAbstractItemView_Update(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "update", Q_ARG(const QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_UpdateDefault(void* ptr, void* index)
{
	if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::update(*static_cast<QModelIndex*>(index));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::update(*static_cast<QModelIndex*>(index));
	} else {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::update(*static_cast<QModelIndex*>(index));
	}
}

int QAbstractItemView_VerticalOffset(void* ptr)
{
	return static_cast<QAbstractItemView*>(ptr)->verticalOffset();
}

void* QAbstractItemView_VisualRect(void* ptr, void* index)
{
	return ({ QRect tmpValue = static_cast<QAbstractItemView*>(ptr)->visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QAbstractItemView_VisualRegionForSelection(void* ptr, void* selection)
{
	return new QRegion(static_cast<QAbstractItemView*>(ptr)->visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

void QAbstractItemView_DestroyQAbstractItemView(void* ptr)
{
	static_cast<QAbstractItemView*>(ptr)->~QAbstractItemView();
}

void QAbstractItemView_DestroyQAbstractItemViewDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

int QAbstractItemView___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QAbstractItemView___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QAbstractItemView___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QAbstractItemView___selectedIndexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemView___selectedIndexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemView___selectedIndexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

class MyQAbstractScrollArea: public QAbstractScrollArea
{
public:
	MyQAbstractScrollArea(QWidget *parent = Q_NULLPTR) : QAbstractScrollArea(parent) {QAbstractScrollArea_QAbstractScrollArea_QRegisterMetaType();};
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void mouseDoubleClickEvent(QMouseEvent * e) { callbackQWidget_MouseDoubleClickEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQAbstractScrollArea() { callbackQAbstractScrollArea_DestroyQAbstractScrollArea(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QAbstractScrollArea*)
Q_DECLARE_METATYPE(MyQAbstractScrollArea*)

int QAbstractScrollArea_QAbstractScrollArea_QRegisterMetaType(){qRegisterMetaType<QAbstractScrollArea*>(); return qRegisterMetaType<MyQAbstractScrollArea*>();}

void* QAbstractScrollArea_NewQAbstractScrollArea(void* parent)
{
		return new MyQAbstractScrollArea(static_cast<QWidget*>(parent));
}

void* QAbstractScrollArea_HorizontalScrollBar(void* ptr)
{
	return static_cast<QAbstractScrollArea*>(ptr)->horizontalScrollBar();
}

long long QAbstractScrollArea_HorizontalScrollBarPolicy(void* ptr)
{
	return static_cast<QAbstractScrollArea*>(ptr)->horizontalScrollBarPolicy();
}

void QAbstractScrollArea_SetHorizontalScrollBar(void* ptr, void* scrollBar)
{
	static_cast<QAbstractScrollArea*>(ptr)->setHorizontalScrollBar(static_cast<QScrollBar*>(scrollBar));
}

void QAbstractScrollArea_SetHorizontalScrollBarPolicy(void* ptr, long long vqt)
{
	static_cast<QAbstractScrollArea*>(ptr)->setHorizontalScrollBarPolicy(static_cast<Qt::ScrollBarPolicy>(vqt));
}

void QAbstractScrollArea_SetVerticalScrollBar(void* ptr, void* scrollBar)
{
	static_cast<QAbstractScrollArea*>(ptr)->setVerticalScrollBar(static_cast<QScrollBar*>(scrollBar));
}

void QAbstractScrollArea_SetVerticalScrollBarPolicy(void* ptr, long long vqt)
{
	static_cast<QAbstractScrollArea*>(ptr)->setVerticalScrollBarPolicy(static_cast<Qt::ScrollBarPolicy>(vqt));
}

void* QAbstractScrollArea_VerticalScrollBar(void* ptr)
{
	return static_cast<QAbstractScrollArea*>(ptr)->verticalScrollBar();
}

long long QAbstractScrollArea_VerticalScrollBarPolicy(void* ptr)
{
	return static_cast<QAbstractScrollArea*>(ptr)->verticalScrollBarPolicy();
}

void QAbstractScrollArea_DestroyQAbstractScrollArea(void* ptr)
{
	static_cast<QAbstractScrollArea*>(ptr)->~QAbstractScrollArea();
}

void QAbstractScrollArea_DestroyQAbstractScrollAreaDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAbstractScrollArea___scrollBarWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractScrollArea___scrollBarWidgets_setList(void* ptr, void* i)
{
		static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QAbstractScrollArea___scrollBarWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

class MyQAbstractSlider: public QAbstractSlider
{
public:
	MyQAbstractSlider(QWidget *parent = Q_NULLPTR) : QAbstractSlider(parent) {QAbstractSlider_QAbstractSlider_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	void keyPressEvent(QKeyEvent * ev) { callbackQWidget_KeyPressEvent(this, ev); };
	void setOrientation(Qt::Orientation vqt) { callbackQAbstractSlider_SetOrientation(this, vqt); };
	void setRange(int min, int max) { callbackQAbstractSlider_SetRange(this, min, max); };
	void timerEvent(QTimerEvent * e) { callbackQWidget_TimerEvent(this, e); };
	 ~MyQAbstractSlider() { callbackQAbstractSlider_DestroyQAbstractSlider(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QAbstractSlider*)
Q_DECLARE_METATYPE(MyQAbstractSlider*)

int QAbstractSlider_QAbstractSlider_QRegisterMetaType(){qRegisterMetaType<QAbstractSlider*>(); return qRegisterMetaType<MyQAbstractSlider*>();}

void* QAbstractSlider_NewQAbstractSlider(void* parent)
{
		return new MyQAbstractSlider(static_cast<QWidget*>(parent));
}

int QAbstractSlider_Maximum(void* ptr)
{
	return static_cast<QAbstractSlider*>(ptr)->maximum();
}

int QAbstractSlider_Minimum(void* ptr)
{
	return static_cast<QAbstractSlider*>(ptr)->minimum();
}

long long QAbstractSlider_Orientation(void* ptr)
{
	return static_cast<QAbstractSlider*>(ptr)->orientation();
}

void QAbstractSlider_SetMaximum(void* ptr, int vin)
{
	static_cast<QAbstractSlider*>(ptr)->setMaximum(vin);
}

void QAbstractSlider_SetMinimum(void* ptr, int vin)
{
	static_cast<QAbstractSlider*>(ptr)->setMinimum(vin);
}

void QAbstractSlider_SetOrientation(void* ptr, long long vqt)
{
	qRegisterMetaType<Qt::Orientation>();
	QMetaObject::invokeMethod(static_cast<QAbstractSlider*>(ptr), "setOrientation", Q_ARG(Qt::Orientation, static_cast<Qt::Orientation>(vqt)));
}

void QAbstractSlider_SetOrientationDefault(void* ptr, long long vqt)
{
	if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::setOrientation(static_cast<Qt::Orientation>(vqt));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::setOrientation(static_cast<Qt::Orientation>(vqt));
	} else {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::setOrientation(static_cast<Qt::Orientation>(vqt));
	}
}

void QAbstractSlider_SetRange(void* ptr, int min, int max)
{
	QMetaObject::invokeMethod(static_cast<QAbstractSlider*>(ptr), "setRange", Q_ARG(int, min), Q_ARG(int, max));
}

void QAbstractSlider_SetRangeDefault(void* ptr, int min, int max)
{
	if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::setRange(min, max);
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::setRange(min, max);
	} else {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::setRange(min, max);
	}
}

int QAbstractSlider_Value(void* ptr)
{
	return static_cast<QAbstractSlider*>(ptr)->value();
}

void QAbstractSlider_DestroyQAbstractSlider(void* ptr)
{
	static_cast<QAbstractSlider*>(ptr)->~QAbstractSlider();
}

void QAbstractSlider_DestroyQAbstractSliderDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQAction: public QAction
{
public:
	MyQAction(QObject *parent = Q_NULLPTR) : QAction(parent) {QAction_QAction_QRegisterMetaType();};
	MyQAction(const QString &text, QObject *parent = Q_NULLPTR) : QAction(text, parent) {QAction_QAction_QRegisterMetaType();};
	MyQAction(const QIcon &icon, const QString &text, QObject *parent = Q_NULLPTR) : QAction(icon, text, parent) {QAction_QAction_QRegisterMetaType();};
	void Signal_Changed() { callbackQAction_Changed(this); };
	bool event(QEvent * e) { return callbackQAction_Event(this, e) != 0; };
	void setChecked(bool vbo) { callbackQAction_SetChecked(this, vbo); };
	void setDisabled(bool b) { callbackQAction_SetDisabled(this, b); };
	void setEnabled(bool vbo) { callbackQAction_SetEnabled(this, vbo); };
	void trigger() { callbackQAction_Trigger(this); };
	void Signal_Triggered(bool checked) { callbackQAction_Triggered(this, checked); };
	 ~MyQAction() { callbackQAction_DestroyQAction(this); };
	void childEvent(QChildEvent * event) { callbackQAction_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAction_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAction_CustomEvent(this, event); };
	void deleteLater() { callbackQAction_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAction_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAction_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAction_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQAction_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAction_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QAction*)
Q_DECLARE_METATYPE(MyQAction*)

int QAction_QAction_QRegisterMetaType(){qRegisterMetaType<QAction*>(); return qRegisterMetaType<MyQAction*>();}

void* QAction_NewQAction(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAction(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAction(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAction(static_cast<QWindow*>(parent));
	} else {
		return new MyQAction(static_cast<QObject*>(parent));
	}
}

void* QAction_NewQAction2(struct QtWidgets_PackedString text, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQAction(QString::fromUtf8(text.data, text.len), static_cast<QObject*>(parent));
	}
}

void* QAction_NewQAction3(void* icon, struct QtWidgets_PackedString text, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QObject*>(parent));
	}
}

void QAction_Activate(void* ptr, long long event)
{
	static_cast<QAction*>(ptr)->activate(static_cast<QAction::ActionEvent>(event));
}

void QAction_ConnectChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::changed), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Changed), static_cast<Qt::ConnectionType>(t));
}

void QAction_DisconnectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::changed), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Changed));
}

void QAction_Changed(void* ptr)
{
	static_cast<QAction*>(ptr)->changed();
}

void* QAction_Data(void* ptr)
{
	return new QVariant(static_cast<QAction*>(ptr)->data());
}

char QAction_EventDefault(void* ptr, void* e)
{
		return static_cast<QAction*>(ptr)->QAction::event(static_cast<QEvent*>(e));
}

void* QAction_Font(void* ptr)
{
	return new QFont(static_cast<QAction*>(ptr)->font());
}

void* QAction_Icon(void* ptr)
{
	return new QIcon(static_cast<QAction*>(ptr)->icon());
}

char QAction_IsChecked(void* ptr)
{
	return static_cast<QAction*>(ptr)->isChecked();
}

void* QAction_Menu(void* ptr)
{
	return static_cast<QAction*>(ptr)->menu();
}

long long QAction_Priority(void* ptr)
{
	return static_cast<QAction*>(ptr)->priority();
}

void QAction_SetCheckable(void* ptr, char vbo)
{
	static_cast<QAction*>(ptr)->setCheckable(vbo != 0);
}

void QAction_SetChecked(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setChecked", Q_ARG(bool, vbo != 0));
}

void QAction_SetCheckedDefault(void* ptr, char vbo)
{
		static_cast<QAction*>(ptr)->QAction::setChecked(vbo != 0);
}

void QAction_SetData(void* ptr, void* userData)
{
	static_cast<QAction*>(ptr)->setData(*static_cast<QVariant*>(userData));
}

void QAction_SetDisabled(void* ptr, char b)
{
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setDisabled", Q_ARG(bool, b != 0));
}

void QAction_SetDisabledDefault(void* ptr, char b)
{
		static_cast<QAction*>(ptr)->QAction::setDisabled(b != 0);
}

void QAction_SetEnabled(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QAction_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QAction*>(ptr)->QAction::setEnabled(vbo != 0);
}

void QAction_SetFont(void* ptr, void* font)
{
	static_cast<QAction*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QAction_SetIcon(void* ptr, void* icon)
{
	static_cast<QAction*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QAction_SetPriority(void* ptr, long long priority)
{
	static_cast<QAction*>(ptr)->setPriority(static_cast<QAction::Priority>(priority));
}

void QAction_SetShortcut(void* ptr, void* shortcut)
{
	static_cast<QAction*>(ptr)->setShortcut(*static_cast<QKeySequence*>(shortcut));
}

void QAction_SetText(void* ptr, struct QtWidgets_PackedString text)
{
	static_cast<QAction*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

void QAction_SetToolTip(void* ptr, struct QtWidgets_PackedString tip)
{
	static_cast<QAction*>(ptr)->setToolTip(QString::fromUtf8(tip.data, tip.len));
}

void* QAction_Shortcut(void* ptr)
{
	return new QKeySequence(static_cast<QAction*>(ptr)->shortcut());
}

struct QtWidgets_PackedString QAction_Text(void* ptr)
{
	return ({ QByteArray* t0f57fa = new QByteArray(static_cast<QAction*>(ptr)->text().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t0f57fa->prepend("WHITESPACE").constData()+10), t0f57fa->size()-10, t0f57fa }; });
}

struct QtWidgets_PackedString QAction_ToolTip(void* ptr)
{
	return ({ QByteArray* td308d4 = new QByteArray(static_cast<QAction*>(ptr)->toolTip().toUtf8()); QtWidgets_PackedString { const_cast<char*>(td308d4->prepend("WHITESPACE").constData()+10), td308d4->size()-10, td308d4 }; });
}

void QAction_Trigger(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "trigger");
}

void QAction_TriggerDefault(void* ptr)
{
		static_cast<QAction*>(ptr)->QAction::trigger();
}

void QAction_ConnectTriggered(void* ptr, long long t)
{
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::triggered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Triggered), static_cast<Qt::ConnectionType>(t));
}

void QAction_DisconnectTriggered(void* ptr)
{
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::triggered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Triggered));
}

void QAction_Triggered(void* ptr, char checked)
{
	static_cast<QAction*>(ptr)->triggered(checked != 0);
}

void QAction_DestroyQAction(void* ptr)
{
	static_cast<QAction*>(ptr)->~QAction();
}

void QAction_DestroyQActionDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAction___associatedGraphicsWidgets_atList(void* ptr, int i)
{
	return ({QGraphicsWidget * tmp = static_cast<QList<QGraphicsWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsWidget *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___associatedGraphicsWidgets_setList(void* ptr, void* i)
{
		static_cast<QList<QGraphicsWidget *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
}

void* QAction___associatedGraphicsWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsWidget *>();
}

void* QAction___associatedWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___associatedWidgets_setList(void* ptr, void* i)
{
		static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QAction___associatedWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* QAction___setShortcuts_shortcuts_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAction___setShortcuts_shortcuts_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QAction___setShortcuts_shortcuts_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QAction___shortcuts_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAction___shortcuts_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QAction___shortcuts_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QAction___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QAction___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QAction___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAction___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAction___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAction___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QAction___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAction___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAction___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QAction___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QAction_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QAction*>(ptr)->QAction::childEvent(static_cast<QChildEvent*>(event));
}

void QAction_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QAction*>(ptr)->QAction::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAction_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QAction*>(ptr)->QAction::customEvent(static_cast<QEvent*>(event));
}

void QAction_DeleteLaterDefault(void* ptr)
{
		static_cast<QAction*>(ptr)->QAction::deleteLater();
}

void QAction_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QAction*>(ptr)->QAction::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAction_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
		return static_cast<QAction*>(ptr)->QAction::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QAction*>(ptr)->QAction::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
		return static_cast<QAction*>(ptr)->QAction::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QAction*>(ptr)->QAction::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QAction*>(ptr)->QAction::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QAction_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QAction*>(ptr)->QAction::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQApplication: public QApplication
{
public:
	MyQApplication(int &argc, char **argv) : QApplication(argc, argv) {QApplication_QApplication_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQApplication_Event(this, e) != 0; };
	void setStyleSheet(const QString & sheet) { QByteArray* t542ebc = new QByteArray(sheet.toUtf8()); QtWidgets_PackedString sheetPacked = { const_cast<char*>(t542ebc->prepend("WHITESPACE").constData()+10), t542ebc->size()-10, t542ebc };callbackQApplication_SetStyleSheet(this, sheetPacked); };
	 ~MyQApplication() { callbackQApplication_DestroyQApplication(this); };
	void quit() { callbackQApplication_Quit(this); };
	void childEvent(QChildEvent * event) { callbackQApplication_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQApplication_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQApplication_CustomEvent(this, event); };
	void deleteLater() { callbackQApplication_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQApplication_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQApplication_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQApplication_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQApplication_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQApplication_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QApplication*)
Q_DECLARE_METATYPE(MyQApplication*)

int QApplication_QApplication_QRegisterMetaType(){qRegisterMetaType<QApplication*>(); return qRegisterMetaType<MyQApplication*>();}

void* QApplication_NewQApplication(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new MyQApplication(argcs, argvs);
}

void* QApplication_QApplication_ActiveWindow()
{
	return QApplication::activeWindow();
}

char QApplication_EventDefault(void* ptr, void* e)
{
		return static_cast<QApplication*>(ptr)->QApplication::event(static_cast<QEvent*>(e));
}

int QApplication_QApplication_Exec()
{
	return QApplication::exec();
}

void* QApplication_QApplication_Font()
{
	return new QFont(QApplication::font());
}

void* QApplication_QApplication_Font2(void* widget)
{
		return new QFont(QApplication::font(static_cast<QWidget*>(widget)));
}

void* QApplication_QApplication_Font3(char* className)
{
	return new QFont(QApplication::font(const_cast<const char*>(className)));
}

void* QApplication_QApplication_FontMetrics()
{
	return new QFontMetrics(QApplication::fontMetrics());
}

void QApplication_QApplication_SetActiveWindow(void* active)
{
		QApplication::setActiveWindow(static_cast<QWidget*>(active));
}

void QApplication_QApplication_SetFont(void* font, char* className)
{
	QApplication::setFont(*static_cast<QFont*>(font), const_cast<const char*>(className));
}

void QApplication_QApplication_SetStyle(void* style)
{
	QApplication::setStyle(static_cast<QStyle*>(style));
}

void* QApplication_QApplication_SetStyle2(struct QtWidgets_PackedString style)
{
	return QApplication::setStyle(QString::fromUtf8(style.data, style.len));
}

void QApplication_SetStyleSheet(void* ptr, struct QtWidgets_PackedString sheet)
{
	QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "setStyleSheet", Q_ARG(const QString, QString::fromUtf8(sheet.data, sheet.len)));
}

void QApplication_SetStyleSheetDefault(void* ptr, struct QtWidgets_PackedString sheet)
{
		static_cast<QApplication*>(ptr)->QApplication::setStyleSheet(QString::fromUtf8(sheet.data, sheet.len));
}

void QApplication_QApplication_SetWindowIcon(void* icon)
{
	QApplication::setWindowIcon(*static_cast<QIcon*>(icon));
}

void* QApplication_QApplication_Style()
{
	return QApplication::style();
}

struct QtWidgets_PackedString QApplication_StyleSheet(void* ptr)
{
	return ({ QByteArray* ta53b66 = new QByteArray(static_cast<QApplication*>(ptr)->styleSheet().toUtf8()); QtWidgets_PackedString { const_cast<char*>(ta53b66->prepend("WHITESPACE").constData()+10), ta53b66->size()-10, ta53b66 }; });
}

void* QApplication_QApplication_WindowIcon()
{
	return new QIcon(QApplication::windowIcon());
}

void QApplication_DestroyQApplication(void* ptr)
{
	static_cast<QApplication*>(ptr)->~QApplication();
}

void QApplication_DestroyQApplicationDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QApplication___allWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___allWidgets_setList(void* ptr, void* i)
{
		static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QApplication___allWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* QApplication___topLevelWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___topLevelWidgets_setList(void* ptr, void* i)
{
		static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QApplication___topLevelWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* QApplication___screens_atList(void* ptr, int i)
{
	return ({QScreen * tmp = static_cast<QList<QScreen *>*>(ptr)->at(i); if (i == static_cast<QList<QScreen *>*>(ptr)->size()-1) { static_cast<QList<QScreen *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___screens_setList(void* ptr, void* i)
{
	static_cast<QList<QScreen *>*>(ptr)->append(static_cast<QScreen*>(i));
}

void* QApplication___screens_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QScreen *>();
}

void* QApplication___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QApplication___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QApplication___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QApplication___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QApplication___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QApplication___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QApplication___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QApplication___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QApplication___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QApplication___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QApplication_QuitDefault(void* ptr)
{
		static_cast<QApplication*>(ptr)->QApplication::quit();
}

void QApplication_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QApplication*>(ptr)->QApplication::childEvent(static_cast<QChildEvent*>(event));
}

void QApplication_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QApplication*>(ptr)->QApplication::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QApplication_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QApplication*>(ptr)->QApplication::customEvent(static_cast<QEvent*>(event));
}

void QApplication_DeleteLaterDefault(void* ptr)
{
		static_cast<QApplication*>(ptr)->QApplication::deleteLater();
}

void QApplication_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QApplication*>(ptr)->QApplication::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QApplication_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
		return static_cast<QApplication*>(ptr)->QApplication::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QApplication*>(ptr)->QApplication::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
		return static_cast<QApplication*>(ptr)->QApplication::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QApplication*>(ptr)->QApplication::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QApplication*>(ptr)->QApplication::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QApplication_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QApplication*>(ptr)->QApplication::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQBoxLayout: public QBoxLayout
{
public:
	MyQBoxLayout(QBoxLayout::Direction dir, QWidget *parent = Q_NULLPTR) : QBoxLayout(dir, parent) {QBoxLayout_QBoxLayout_QRegisterMetaType();};
	void addItem(QLayoutItem * item) { callbackQBoxLayout_AddItem(this, item); };
	int count() const { return callbackQBoxLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayout_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQLayoutItem_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQLayoutItem_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void invalidate() { callbackQLayoutItem_Invalidate(this); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQBoxLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayout_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	int minimumHeightForWidth(int w) const { return callbackQLayoutItem_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayout_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	void setGeometry(const QRect & r) { callbackQLayout_SetGeometry(this, const_cast<QRect*>(&r)); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQBoxLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQBoxLayout_TakeAt(this, index)); };
	 ~MyQBoxLayout() { callbackQBoxLayout_DestroyQBoxLayout(this); };
	void childEvent(QChildEvent * e) { callbackQLayout_ChildEvent(this, e); };
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackQLayoutItem_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayout_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	int indexOf(QWidget * widget) const { return callbackQLayout_IndexOf(const_cast<void*>(static_cast<const void*>(this)), widget); };
	bool isEmpty() const { return callbackQLayout_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	void connectNotify(const QMetaMethod & sign) { callbackQLayout_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLayout_CustomEvent(this, event); };
	void deleteLater() { callbackQLayout_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLayout_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLayout_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQLayout_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLayout_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQLayout_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLayout_TimerEvent(this, event); };
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackQLayoutItem_SpacerItem(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
};

Q_DECLARE_METATYPE(QBoxLayout*)
Q_DECLARE_METATYPE(MyQBoxLayout*)

int QBoxLayout_QBoxLayout_QRegisterMetaType(){qRegisterMetaType<QBoxLayout*>(); return qRegisterMetaType<MyQBoxLayout*>();}

void* QBoxLayout_NewQBoxLayout(long long dir, void* parent)
{
		return new MyQBoxLayout(static_cast<QBoxLayout::Direction>(dir), static_cast<QWidget*>(parent));
}

void QBoxLayout_AddItem(void* ptr, void* item)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
		static_cast<QBoxLayout*>(ptr)->addItem(static_cast<QLayout*>(item));
	} else {
		static_cast<QBoxLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
	}
}

void QBoxLayout_AddItemDefault(void* ptr, void* item)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
			static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::addItem(static_cast<QLayout*>(item));
		} else {
			static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::addItem(static_cast<QLayoutItem*>(item));
		}
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
			static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::addItem(static_cast<QLayout*>(item));
		} else {
			static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::addItem(static_cast<QLayoutItem*>(item));
		}
	} else {
		if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
			static_cast<QBoxLayout*>(ptr)->QBoxLayout::addItem(static_cast<QLayout*>(item));
		} else {
			static_cast<QBoxLayout*>(ptr)->QBoxLayout::addItem(static_cast<QLayoutItem*>(item));
		}
	}
}

void QBoxLayout_AddLayout(void* ptr, void* layout, int stretch)
{
		static_cast<QBoxLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), stretch);
}

void QBoxLayout_AddWidget(void* ptr, void* widget, int stretch, long long alignment)
{
		static_cast<QBoxLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), stretch, static_cast<Qt::AlignmentFlag>(alignment));
}

int QBoxLayout_Count(void* ptr)
{
	return static_cast<QBoxLayout*>(ptr)->count();
}

int QBoxLayout_CountDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::count();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::count();
	} else {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::count();
	}
}

long long QBoxLayout_Direction(void* ptr)
{
	return static_cast<QBoxLayout*>(ptr)->direction();
}

void* QBoxLayout_ItemAt(void* ptr, int index)
{
	return static_cast<QBoxLayout*>(ptr)->itemAt(index);
}

void* QBoxLayout_ItemAtDefault(void* ptr, int index)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::itemAt(index);
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::itemAt(index);
	} else {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::itemAt(index);
	}
}

void QBoxLayout_SetStretch(void* ptr, int index, int stretch)
{
	static_cast<QBoxLayout*>(ptr)->setStretch(index, stretch);
}

char QBoxLayout_SetStretchFactor(void* ptr, void* widget, int stretch)
{
		return static_cast<QBoxLayout*>(ptr)->setStretchFactor(static_cast<QWidget*>(widget), stretch);
}

char QBoxLayout_SetStretchFactor2(void* ptr, void* layout, int stretch)
{
		return static_cast<QBoxLayout*>(ptr)->setStretchFactor(static_cast<QLayout*>(layout), stretch);
}

void* QBoxLayout_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QBoxLayout*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QBoxLayout_SizeHintDefault(void* ptr)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QBoxLayout*>(ptr)->QBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

int QBoxLayout_Stretch(void* ptr, int index)
{
	return static_cast<QBoxLayout*>(ptr)->stretch(index);
}

void* QBoxLayout_TakeAt(void* ptr, int index)
{
	return static_cast<QBoxLayout*>(ptr)->takeAt(index);
}

void* QBoxLayout_TakeAtDefault(void* ptr, int index)
{
	if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::takeAt(index);
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::takeAt(index);
	} else {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::takeAt(index);
	}
}

void QBoxLayout_DestroyQBoxLayout(void* ptr)
{
	static_cast<QBoxLayout*>(ptr)->~QBoxLayout();
}

void QBoxLayout_DestroyQBoxLayoutDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQButtonGroup: public QButtonGroup
{
public:
	MyQButtonGroup(QObject *parent = Q_NULLPTR) : QButtonGroup(parent) {QButtonGroup_QButtonGroup_QRegisterMetaType();};
	 ~MyQButtonGroup() { callbackQButtonGroup_DestroyQButtonGroup(this); };
	void childEvent(QChildEvent * event) { callbackQButtonGroup_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQButtonGroup_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQButtonGroup_CustomEvent(this, event); };
	void deleteLater() { callbackQButtonGroup_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQButtonGroup_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQButtonGroup_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQButtonGroup_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQButtonGroup_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQButtonGroup_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQButtonGroup_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QButtonGroup*)
Q_DECLARE_METATYPE(MyQButtonGroup*)

int QButtonGroup_QButtonGroup_QRegisterMetaType(){qRegisterMetaType<QButtonGroup*>(); return qRegisterMetaType<MyQButtonGroup*>();}

void* QButtonGroup_NewQButtonGroup(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQButtonGroup(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQButtonGroup(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQButtonGroup(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQButtonGroup(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQButtonGroup(static_cast<QWindow*>(parent));
	} else {
		return new MyQButtonGroup(static_cast<QObject*>(parent));
	}
}

void* QButtonGroup_Button(void* ptr, int id)
{
	return static_cast<QButtonGroup*>(ptr)->button(id);
}

struct QtWidgets_PackedList QButtonGroup_Buttons(void* ptr)
{
	return ({ QList<QAbstractButton *>* tmpValuea4306e = new QList<QAbstractButton *>(static_cast<QButtonGroup*>(ptr)->buttons()); QtWidgets_PackedList { tmpValuea4306e, tmpValuea4306e->size() }; });
}

int QButtonGroup_Id(void* ptr, void* button)
{
	return static_cast<QButtonGroup*>(ptr)->id(static_cast<QAbstractButton*>(button));
}

void QButtonGroup_DestroyQButtonGroup(void* ptr)
{
	static_cast<QButtonGroup*>(ptr)->~QButtonGroup();
}

void QButtonGroup_DestroyQButtonGroupDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QButtonGroup___buttons_atList(void* ptr, int i)
{
	return ({QAbstractButton * tmp = static_cast<QList<QAbstractButton *>*>(ptr)->at(i); if (i == static_cast<QList<QAbstractButton *>*>(ptr)->size()-1) { static_cast<QList<QAbstractButton *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QButtonGroup___buttons_setList(void* ptr, void* i)
{
	static_cast<QList<QAbstractButton *>*>(ptr)->append(static_cast<QAbstractButton*>(i));
}

void* QButtonGroup___buttons_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAbstractButton *>();
}

void* QButtonGroup___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QButtonGroup___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QButtonGroup___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QButtonGroup___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QButtonGroup___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QButtonGroup___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QButtonGroup___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QButtonGroup___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QButtonGroup___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QButtonGroup___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QButtonGroup___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QButtonGroup___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QButtonGroup_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QButtonGroup*>(ptr)->QButtonGroup::childEvent(static_cast<QChildEvent*>(event));
}

void QButtonGroup_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QButtonGroup*>(ptr)->QButtonGroup::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QButtonGroup_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QButtonGroup*>(ptr)->QButtonGroup::customEvent(static_cast<QEvent*>(event));
}

void QButtonGroup_DeleteLaterDefault(void* ptr)
{
		static_cast<QButtonGroup*>(ptr)->QButtonGroup::deleteLater();
}

void QButtonGroup_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QButtonGroup*>(ptr)->QButtonGroup::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QButtonGroup_EventDefault(void* ptr, void* e)
{
		return static_cast<QButtonGroup*>(ptr)->QButtonGroup::event(static_cast<QEvent*>(e));
}

char QButtonGroup_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
		return static_cast<QButtonGroup*>(ptr)->QButtonGroup::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QButtonGroup*>(ptr)->QButtonGroup::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
		return static_cast<QButtonGroup*>(ptr)->QButtonGroup::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QButtonGroup*>(ptr)->QButtonGroup::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QButtonGroup*>(ptr)->QButtonGroup::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QButtonGroup_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QButtonGroup*>(ptr)->QButtonGroup::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQColorDialog: public QColorDialog
{
public:
	MyQColorDialog(QWidget *parent = Q_NULLPTR) : QColorDialog(parent) {QColorDialog_QColorDialog_QRegisterMetaType();};
	MyQColorDialog(const QColor &initial, QWidget *parent = Q_NULLPTR) : QColorDialog(initial, parent) {QColorDialog_QColorDialog_QRegisterMetaType();};
	void done(int result) { callbackQColorDialog_Done(this, result); };
	 ~MyQColorDialog() { callbackQColorDialog_DestroyQColorDialog(this); };
	void accept() { callbackQDialog_Accept(this); };
	void closeEvent(QCloseEvent * e) { callbackQWidget_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	bool eventFilter(QObject * o, QEvent * e) { return callbackQWidget_EventFilter(this, o, e) != 0; };
	int exec() { return callbackQDialog_Exec(this); };
	void Signal_Finished(int result) { callbackQDialog_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void resizeEvent(QResizeEvent * vqr) { callbackQWidget_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QColorDialog*)
Q_DECLARE_METATYPE(MyQColorDialog*)

int QColorDialog_QColorDialog_QRegisterMetaType(){qRegisterMetaType<QColorDialog*>(); return qRegisterMetaType<MyQColorDialog*>();}

void* QColorDialog_NewQColorDialog(void* parent)
{
		return new MyQColorDialog(static_cast<QWidget*>(parent));
}

void* QColorDialog_NewQColorDialog2(void* initial, void* parent)
{
		return new MyQColorDialog(*static_cast<QColor*>(initial), static_cast<QWidget*>(parent));
}

void QColorDialog_Done(void* ptr, int result)
{
	static_cast<QColorDialog*>(ptr)->done(result);
}

void QColorDialog_DoneDefault(void* ptr, int result)
{
		static_cast<QColorDialog*>(ptr)->QColorDialog::done(result);
}

void* QColorDialog_QColorDialog_GetColor(void* initial, void* parent, struct QtWidgets_PackedString title, long long options)
{
		return new QColor(QColorDialog::getColor(*static_cast<QColor*>(initial), static_cast<QWidget*>(parent), QString::fromUtf8(title.data, title.len), static_cast<QColorDialog::ColorDialogOption>(options)));
}

void QColorDialog_Open(void* ptr, void* receiver, char* member)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(receiver))) {
		static_cast<QColorDialog*>(ptr)->open(static_cast<QGraphicsObject*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(receiver))) {
		static_cast<QColorDialog*>(ptr)->open(static_cast<QGraphicsWidget*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(receiver))) {
		static_cast<QColorDialog*>(ptr)->open(static_cast<QLayout*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(receiver))) {
		static_cast<QColorDialog*>(ptr)->open(static_cast<QWidget*>(receiver), const_cast<const char*>(member));
	} else {
		static_cast<QColorDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
	}
}

void QColorDialog_DestroyQColorDialog(void* ptr)
{
	static_cast<QColorDialog*>(ptr)->~QColorDialog();
}

void QColorDialog_DestroyQColorDialogDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQComboBox: public QComboBox
{
public:
	MyQComboBox(QWidget *parent = Q_NULLPTR) : QComboBox(parent) {QComboBox_QComboBox_QRegisterMetaType();};
	void Signal_Activated(int index) { callbackQComboBox_Activated(this, index); };
	void clear() { callbackQComboBox_Clear(this); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	void hideEvent(QHideEvent * e) { callbackQWidget_HideEvent(this, e); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void keyReleaseEvent(QKeyEvent * e) { callbackQWidget_KeyReleaseEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	void resizeEvent(QResizeEvent * e) { callbackQWidget_ResizeEvent(this, e); };
	void setCurrentIndex(int index) { callbackQComboBox_SetCurrentIndex(this, index); };
	void setCurrentText(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtWidgets_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQComboBox_SetCurrentText(this, textPacked); };
	void showEvent(QShowEvent * e) { callbackQWidget_ShowEvent(this, e); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQComboBox() { callbackQComboBox_DestroyQComboBox(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QComboBox*)
Q_DECLARE_METATYPE(MyQComboBox*)

int QComboBox_QComboBox_QRegisterMetaType(){qRegisterMetaType<QComboBox*>(); return qRegisterMetaType<MyQComboBox*>();}

int QComboBox_AdjustToMinimumContentsLengthWithIcon_Type()
{
	return QComboBox::AdjustToMinimumContentsLengthWithIcon;
}

void* QComboBox_NewQComboBox(void* parent)
{
		return new MyQComboBox(static_cast<QWidget*>(parent));
}

void QComboBox_ConnectActivated(void* ptr, long long t)
{
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::activated), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Activated), static_cast<Qt::ConnectionType>(t));
}

void QComboBox_DisconnectActivated(void* ptr)
{
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::activated), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Activated));
}

void QComboBox_Activated(void* ptr, int index)
{
	static_cast<QComboBox*>(ptr)->activated(index);
}

void QComboBox_AddItem(void* ptr, struct QtWidgets_PackedString text, void* userData)
{
	static_cast<QComboBox*>(ptr)->addItem(QString::fromUtf8(text.data, text.len), *static_cast<QVariant*>(userData));
}

void QComboBox_AddItem2(void* ptr, void* icon, struct QtWidgets_PackedString text, void* userData)
{
	static_cast<QComboBox*>(ptr)->addItem(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), *static_cast<QVariant*>(userData));
}

void QComboBox_AddItems(void* ptr, struct QtWidgets_PackedString texts)
{
	static_cast<QComboBox*>(ptr)->addItems(QString::fromUtf8(texts.data, texts.len).split("¡¦!", QString::SkipEmptyParts));
}

void QComboBox_Clear(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "clear");
}

void QComboBox_ClearDefault(void* ptr)
{
		static_cast<QComboBox*>(ptr)->QComboBox::clear();
}

int QComboBox_Count(void* ptr)
{
	return static_cast<QComboBox*>(ptr)->count();
}

int QComboBox_CurrentIndex(void* ptr)
{
	return static_cast<QComboBox*>(ptr)->currentIndex();
}

struct QtWidgets_PackedString QComboBox_CurrentText(void* ptr)
{
	return ({ QByteArray* t0977ce = new QByteArray(static_cast<QComboBox*>(ptr)->currentText().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t0977ce->prepend("WHITESPACE").constData()+10), t0977ce->size()-10, t0977ce }; });
}

void* QComboBox_LineEdit(void* ptr)
{
	return static_cast<QComboBox*>(ptr)->lineEdit();
}

void* QComboBox_Model(void* ptr)
{
	return static_cast<QComboBox*>(ptr)->model();
}

struct QtWidgets_PackedString QComboBox_PlaceholderText(void* ptr)
{
	return ({ QByteArray* t542619 = new QByteArray(static_cast<QComboBox*>(ptr)->placeholderText().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t542619->prepend("WHITESPACE").constData()+10), t542619->size()-10, t542619 }; });
}

void QComboBox_SetCurrentIndex(void* ptr, int index)
{
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QComboBox_SetCurrentIndexDefault(void* ptr, int index)
{
		static_cast<QComboBox*>(ptr)->QComboBox::setCurrentIndex(index);
}

void QComboBox_SetCurrentText(void* ptr, struct QtWidgets_PackedString text)
{
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setCurrentText", Q_ARG(const QString, QString::fromUtf8(text.data, text.len)));
}

void QComboBox_SetCurrentTextDefault(void* ptr, struct QtWidgets_PackedString text)
{
		static_cast<QComboBox*>(ptr)->QComboBox::setCurrentText(QString::fromUtf8(text.data, text.len));
}

void QComboBox_SetEditable(void* ptr, char editable)
{
	static_cast<QComboBox*>(ptr)->setEditable(editable != 0);
}

void QComboBox_SetModel(void* ptr, void* model)
{
	static_cast<QComboBox*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QComboBox_SetPlaceholderText(void* ptr, struct QtWidgets_PackedString placeholderText)
{
	static_cast<QComboBox*>(ptr)->setPlaceholderText(QString::fromUtf8(placeholderText.data, placeholderText.len));
}

void QComboBox_SetValidator(void* ptr, void* validator)
{
	static_cast<QComboBox*>(ptr)->setValidator(static_cast<QValidator*>(validator));
}

void* QComboBox_Validator(void* ptr)
{
	return const_cast<QValidator*>(static_cast<QComboBox*>(ptr)->validator());
}

void* QComboBox_View(void* ptr)
{
	return static_cast<QComboBox*>(ptr)->view();
}

void QComboBox_DestroyQComboBox(void* ptr)
{
	static_cast<QComboBox*>(ptr)->~QComboBox();
}

void QComboBox_DestroyQComboBoxDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQDial: public QDial
{
public:
	MyQDial(QWidget *parent = Q_NULLPTR) : QDial(parent) {QDial_QDial_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	void resizeEvent(QResizeEvent * e) { callbackQWidget_ResizeEvent(this, e); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQDial() { callbackQDial_DestroyQDial(this); };
	void keyPressEvent(QKeyEvent * ev) { callbackQWidget_KeyPressEvent(this, ev); };
	void setOrientation(Qt::Orientation vqt) { callbackQAbstractSlider_SetOrientation(this, vqt); };
	void setRange(int min, int max) { callbackQAbstractSlider_SetRange(this, min, max); };
	void timerEvent(QTimerEvent * e) { callbackQWidget_TimerEvent(this, e); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QDial*)
Q_DECLARE_METATYPE(MyQDial*)

int QDial_QDial_QRegisterMetaType(){qRegisterMetaType<QDial*>(); return qRegisterMetaType<MyQDial*>();}

void* QDial_NewQDial(void* parent)
{
		return new MyQDial(static_cast<QWidget*>(parent));
}

void QDial_DestroyQDial(void* ptr)
{
	static_cast<QDial*>(ptr)->~QDial();
}

void QDial_DestroyQDialDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQDialog: public QDialog
{
public:
	MyQDialog(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QDialog(parent, ff) {QDialog_QDialog_QRegisterMetaType();};
	void accept() { callbackQDialog_Accept(this); };
	void closeEvent(QCloseEvent * e) { callbackQWidget_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	void done(int r) { callbackQDialog_Done(this, r); };
	bool eventFilter(QObject * o, QEvent * e) { return callbackQWidget_EventFilter(this, o, e) != 0; };
	int exec() { return callbackQDialog_Exec(this); };
	void Signal_Finished(int result) { callbackQDialog_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void open() { callbackQDialog_Open(this); };
	void resizeEvent(QResizeEvent * vqr) { callbackQWidget_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQDialog() { callbackQDialog_DestroyQDialog(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QDialog*)
Q_DECLARE_METATYPE(MyQDialog*)

int QDialog_QDialog_QRegisterMetaType(){qRegisterMetaType<QDialog*>(); return qRegisterMetaType<MyQDialog*>();}

void* QDialog_NewQDialog(void* parent, long long ff)
{
		return new MyQDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void QDialog_Accept(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "accept");
}

void QDialog_AcceptDefault(void* ptr)
{
	if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::accept();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::accept();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::accept();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::accept();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::accept();
	} else {
		static_cast<QDialog*>(ptr)->QDialog::accept();
	}
}

void QDialog_Done(void* ptr, int r)
{
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "done", Q_ARG(int, r));
}

void QDialog_DoneDefault(void* ptr, int r)
{
	if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::done(r);
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::done(r);
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::done(r);
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::done(r);
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::done(r);
	} else {
		static_cast<QDialog*>(ptr)->QDialog::done(r);
	}
}

int QDialog_Exec(void* ptr)
{
	int returnArg;
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "exec", Q_RETURN_ARG(int, returnArg));
	return returnArg;
}

int QDialog_ExecDefault(void* ptr)
{
	if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMessageBox*>(ptr)->QMessageBox::exec();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QInputDialog*>(ptr)->QInputDialog::exec();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFontDialog*>(ptr)->QFontDialog::exec();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDialog*>(ptr)->QFileDialog::exec();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QColorDialog*>(ptr)->QColorDialog::exec();
	} else {
		return static_cast<QDialog*>(ptr)->QDialog::exec();
	}
}

void QDialog_ConnectFinished(void* ptr, long long t)
{
	QObject::connect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)(int)>(&QDialog::finished), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)(int)>(&MyQDialog::Signal_Finished), static_cast<Qt::ConnectionType>(t));
}

void QDialog_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)(int)>(&QDialog::finished), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)(int)>(&MyQDialog::Signal_Finished));
}

void QDialog_Finished(void* ptr, int result)
{
	static_cast<QDialog*>(ptr)->finished(result);
}

void QDialog_Open(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "open");
}

void QDialog_OpenDefault(void* ptr)
{
	if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::open();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::open();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::open();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::open();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::open();
	} else {
		static_cast<QDialog*>(ptr)->QDialog::open();
	}
}

int QDialog_Result(void* ptr)
{
	return static_cast<QDialog*>(ptr)->result();
}

void QDialog_SetModal(void* ptr, char modal)
{
	static_cast<QDialog*>(ptr)->setModal(modal != 0);
}

void QDialog_DestroyQDialog(void* ptr)
{
	static_cast<QDialog*>(ptr)->~QDialog();
}

void QDialog_DestroyQDialogDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

int QDialogButtonBox_AndroidLayout_Type()
{
	return QDialogButtonBox::AndroidLayout;
}

int QDirModel_FilePathRole_Type()
{
	return QDirModel::FilePathRole;
}

class MyQDockWidget: public QDockWidget
{
public:
	MyQDockWidget(const QString &title, QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QDockWidget(title, parent, flags) {QDockWidget_QDockWidget_QRegisterMetaType();};
	MyQDockWidget(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QDockWidget(parent, flags) {QDockWidget_QDockWidget_QRegisterMetaType();};
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	 ~MyQDockWidget() { callbackQDockWidget_DestroyQDockWidget(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QDockWidget*)
Q_DECLARE_METATYPE(MyQDockWidget*)

int QDockWidget_QDockWidget_QRegisterMetaType(){qRegisterMetaType<QDockWidget*>(); return qRegisterMetaType<MyQDockWidget*>();}

void* QDockWidget_NewQDockWidget(struct QtWidgets_PackedString title, void* parent, long long flags)
{
		return new MyQDockWidget(QString::fromUtf8(title.data, title.len), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QDockWidget_NewQDockWidget2(void* parent, long long flags)
{
		return new MyQDockWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QDockWidget_SetWidget(void* ptr, void* widget)
{
		static_cast<QDockWidget*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void* QDockWidget_Widget(void* ptr)
{
	return static_cast<QDockWidget*>(ptr)->widget();
}

void QDockWidget_DestroyQDockWidget(void* ptr)
{
	static_cast<QDockWidget*>(ptr)->~QDockWidget();
}

void QDockWidget_DestroyQDockWidgetDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQFileDialog: public QFileDialog
{
public:
	MyQFileDialog(QWidget *parent, Qt::WindowFlags flags) : QFileDialog(parent, flags) {QFileDialog_QFileDialog_QRegisterMetaType();};
	MyQFileDialog(QWidget *parent = Q_NULLPTR, const QString &caption = QString(), const QString &directory = QString(), const QString &filter = QString()) : QFileDialog(parent, caption, directory, filter) {QFileDialog_QFileDialog_QRegisterMetaType();};
	void accept() { callbackQFileDialog_Accept(this); };
	void done(int result) { callbackQFileDialog_Done(this, result); };
	 ~MyQFileDialog() { callbackQFileDialog_DestroyQFileDialog(this); };
	void closeEvent(QCloseEvent * e) { callbackQWidget_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	bool eventFilter(QObject * o, QEvent * e) { return callbackQWidget_EventFilter(this, o, e) != 0; };
	int exec() { return callbackQDialog_Exec(this); };
	void Signal_Finished(int result) { callbackQDialog_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void resizeEvent(QResizeEvent * vqr) { callbackQWidget_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QFileDialog*)
Q_DECLARE_METATYPE(MyQFileDialog*)

int QFileDialog_QFileDialog_QRegisterMetaType(){qRegisterMetaType<QFileDialog*>(); return qRegisterMetaType<MyQFileDialog*>();}

void* QFileDialog_NewQFileDialog(void* parent, long long flags)
{
		return new MyQFileDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QFileDialog_NewQFileDialog2(void* parent, struct QtWidgets_PackedString caption, struct QtWidgets_PackedString directory, struct QtWidgets_PackedString filter)
{
		return new MyQFileDialog(static_cast<QWidget*>(parent), QString::fromUtf8(caption.data, caption.len), QString::fromUtf8(directory.data, directory.len), QString::fromUtf8(filter.data, filter.len));
}

void QFileDialog_Accept(void* ptr)
{
	static_cast<QFileDialog*>(ptr)->accept();
}

void QFileDialog_AcceptDefault(void* ptr)
{
		static_cast<QFileDialog*>(ptr)->QFileDialog::accept();
}

void* QFileDialog_Directory(void* ptr)
{
	return new QDir(static_cast<QFileDialog*>(ptr)->directory());
}

void QFileDialog_Done(void* ptr, int result)
{
	static_cast<QFileDialog*>(ptr)->done(result);
}

void QFileDialog_DoneDefault(void* ptr, int result)
{
		static_cast<QFileDialog*>(ptr)->QFileDialog::done(result);
}

long long QFileDialog_Filter(void* ptr)
{
	return static_cast<QFileDialog*>(ptr)->filter();
}

struct QtWidgets_PackedString QFileDialog_QFileDialog_GetExistingDirectory(void* parent, struct QtWidgets_PackedString caption, struct QtWidgets_PackedString dir, long long options)
{
		return ({ QByteArray* t875e51 = new QByteArray(QFileDialog::getExistingDirectory(static_cast<QWidget*>(parent), QString::fromUtf8(caption.data, caption.len), QString::fromUtf8(dir.data, dir.len), static_cast<QFileDialog::Option>(options)).toUtf8()); QtWidgets_PackedString { const_cast<char*>(t875e51->prepend("WHITESPACE").constData()+10), t875e51->size()-10, t875e51 }; });
}

struct QtWidgets_PackedString QFileDialog_QFileDialog_GetOpenFileName(void* parent, struct QtWidgets_PackedString caption, struct QtWidgets_PackedString dir, struct QtWidgets_PackedString filter, struct QtWidgets_PackedString selectedFilter, long long options)
{
		return ({ QByteArray* t2c631a = new QByteArray(QFileDialog::getOpenFileName(static_cast<QWidget*>(parent), QString::fromUtf8(caption.data, caption.len), QString::fromUtf8(dir.data, dir.len), QString::fromUtf8(filter.data, filter.len), new QString(QString::fromUtf8(selectedFilter.data, selectedFilter.len)), static_cast<QFileDialog::Option>(options)).toUtf8()); QtWidgets_PackedString { const_cast<char*>(t2c631a->prepend("WHITESPACE").constData()+10), t2c631a->size()-10, t2c631a }; });
}

struct QtWidgets_PackedString QFileDialog_QFileDialog_GetSaveFileName(void* parent, struct QtWidgets_PackedString caption, struct QtWidgets_PackedString dir, struct QtWidgets_PackedString filter, struct QtWidgets_PackedString selectedFilter, long long options)
{
		return ({ QByteArray* t495363 = new QByteArray(QFileDialog::getSaveFileName(static_cast<QWidget*>(parent), QString::fromUtf8(caption.data, caption.len), QString::fromUtf8(dir.data, dir.len), QString::fromUtf8(filter.data, filter.len), new QString(QString::fromUtf8(selectedFilter.data, selectedFilter.len)), static_cast<QFileDialog::Option>(options)).toUtf8()); QtWidgets_PackedString { const_cast<char*>(t495363->prepend("WHITESPACE").constData()+10), t495363->size()-10, t495363 }; });
}

void QFileDialog_Open(void* ptr, void* receiver, char* member)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(receiver))) {
		static_cast<QFileDialog*>(ptr)->open(static_cast<QGraphicsObject*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(receiver))) {
		static_cast<QFileDialog*>(ptr)->open(static_cast<QGraphicsWidget*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(receiver))) {
		static_cast<QFileDialog*>(ptr)->open(static_cast<QLayout*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(receiver))) {
		static_cast<QFileDialog*>(ptr)->open(static_cast<QWidget*>(receiver), const_cast<const char*>(member));
	} else {
		static_cast<QFileDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
	}
}

void QFileDialog_SetReadOnly(void* ptr, char enabled)
{
	static_cast<QFileDialog*>(ptr)->setReadOnly(enabled != 0);
}

void QFileDialog_DestroyQFileDialog(void* ptr)
{
	static_cast<QFileDialog*>(ptr)->~QFileDialog();
}

void QFileDialog_DestroyQFileDialogDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QFileDialog___getOpenFileUrls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QFileDialog___getOpenFileUrls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QFileDialog___getOpenFileUrls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QFileDialog___selectedUrls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QFileDialog___selectedUrls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QFileDialog___selectedUrls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QFileDialog___setSidebarUrls_urls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QFileDialog___setSidebarUrls_urls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QFileDialog___setSidebarUrls_urls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QFileDialog___sidebarUrls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QFileDialog___sidebarUrls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QFileDialog___sidebarUrls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QFileDialog___urlsSelected_urls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QFileDialog___urlsSelected_urls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QFileDialog___urlsSelected_urls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

int QFileSystemModel_FilePathRole_Type()
{
	return QFileSystemModel::FilePathRole;
}

int QFileSystemModel_FileNameRole_Type()
{
	return QFileSystemModel::FileNameRole;
}

int QFileSystemModel_FilePermissions_Type()
{
	return QFileSystemModel::FilePermissions;
}

class MyQFontDialog: public QFontDialog
{
public:
	MyQFontDialog(QWidget *parent = Q_NULLPTR) : QFontDialog(parent) {QFontDialog_QFontDialog_QRegisterMetaType();};
	MyQFontDialog(const QFont &initial, QWidget *parent = Q_NULLPTR) : QFontDialog(initial, parent) {QFontDialog_QFontDialog_QRegisterMetaType();};
	void done(int result) { callbackQFontDialog_Done(this, result); };
	void accept() { callbackQDialog_Accept(this); };
	void closeEvent(QCloseEvent * e) { callbackQWidget_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	bool eventFilter(QObject * o, QEvent * e) { return callbackQWidget_EventFilter(this, o, e) != 0; };
	int exec() { return callbackQDialog_Exec(this); };
	void Signal_Finished(int result) { callbackQDialog_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void resizeEvent(QResizeEvent * vqr) { callbackQWidget_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QFontDialog*)
Q_DECLARE_METATYPE(MyQFontDialog*)

int QFontDialog_QFontDialog_QRegisterMetaType(){qRegisterMetaType<QFontDialog*>(); return qRegisterMetaType<MyQFontDialog*>();}

void* QFontDialog_NewQFontDialog(void* parent)
{
		return new MyQFontDialog(static_cast<QWidget*>(parent));
}

void* QFontDialog_NewQFontDialog2(void* initial, void* parent)
{
		return new MyQFontDialog(*static_cast<QFont*>(initial), static_cast<QWidget*>(parent));
}

void* QFontDialog_CurrentFont(void* ptr)
{
	return new QFont(static_cast<QFontDialog*>(ptr)->currentFont());
}

void QFontDialog_Done(void* ptr, int result)
{
	static_cast<QFontDialog*>(ptr)->done(result);
}

void QFontDialog_DoneDefault(void* ptr, int result)
{
		static_cast<QFontDialog*>(ptr)->QFontDialog::done(result);
}

void* QFontDialog_QFontDialog_GetFont(char* ok, void* initial, void* parent, struct QtWidgets_PackedString title, long long options)
{
		return new QFont(QFontDialog::getFont(reinterpret_cast<bool*>(ok), *static_cast<QFont*>(initial), static_cast<QWidget*>(parent), QString::fromUtf8(title.data, title.len), static_cast<QFontDialog::FontDialogOption>(options)));
}

void* QFontDialog_QFontDialog_GetFont2(char* ok, void* parent)
{
		return new QFont(QFontDialog::getFont(reinterpret_cast<bool*>(ok), static_cast<QWidget*>(parent)));
}

void QFontDialog_Open(void* ptr, void* receiver, char* member)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(receiver))) {
		static_cast<QFontDialog*>(ptr)->open(static_cast<QGraphicsObject*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(receiver))) {
		static_cast<QFontDialog*>(ptr)->open(static_cast<QGraphicsWidget*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(receiver))) {
		static_cast<QFontDialog*>(ptr)->open(static_cast<QLayout*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(receiver))) {
		static_cast<QFontDialog*>(ptr)->open(static_cast<QWidget*>(receiver), const_cast<const char*>(member));
	} else {
		static_cast<QFontDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
	}
}

class MyQFrame: public QFrame
{
public:
	MyQFrame(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QFrame(parent, ff) {QFrame_QFrame_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQFrame() { callbackQFrame_DestroyQFrame(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QFrame*)
Q_DECLARE_METATYPE(MyQFrame*)

int QFrame_QFrame_QRegisterMetaType(){qRegisterMetaType<QFrame*>(); return qRegisterMetaType<MyQFrame*>();}

void* QFrame_NewQFrame(void* parent, long long ff)
{
		return new MyQFrame(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void QFrame_DestroyQFrame(void* ptr)
{
	static_cast<QFrame*>(ptr)->~QFrame();
}

void QFrame_DestroyQFrameDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQGraphicsItem: public QGraphicsItem
{
public:
	MyQGraphicsItem(QGraphicsItem *parent = Q_NULLPTR) : QGraphicsItem(parent) {QGraphicsItem_QGraphicsItem_QRegisterMetaType();};
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsItem_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQGraphicsItem_ContextMenuEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQGraphicsItem_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQGraphicsItem_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MousePressEvent(this, event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MouseReleaseEvent(this, event); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsItem_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	int type() const { return callbackQGraphicsItem_Type(const_cast<void*>(static_cast<const void*>(this))); };
	 ~MyQGraphicsItem() { callbackQGraphicsItem_DestroyQGraphicsItem(this); };
};

Q_DECLARE_METATYPE(MyQGraphicsItem*)

int QGraphicsItem_QGraphicsItem_QRegisterMetaType(){qRegisterMetaType<QGraphicsItem*>(); return qRegisterMetaType<MyQGraphicsItem*>();}

void* QGraphicsItem_NewQGraphicsItem(void* parent)
{
	return new MyQGraphicsItem(static_cast<QGraphicsItem*>(parent));
}

void* QGraphicsItem_BoundingRect(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QGraphicsObject*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QRectF tmpValue = static_cast<QGraphicsItem*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	}
}

void QGraphicsItem_ContextMenuEvent(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
	}
}

void QGraphicsItem_ContextMenuEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
	} else if (dynamic_cast<QGraphicsItemGroup*>(static_cast<QGraphicsItem*>(ptr))) {
		static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->QGraphicsItem::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
	}
}

void* QGraphicsItem_Cursor(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QCursor(static_cast<QGraphicsObject*>(ptr)->cursor());
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QCursor(static_cast<QGraphicsWidget*>(ptr)->cursor());
	} else {
		return new QCursor(static_cast<QGraphicsItem*>(ptr)->cursor());
	}
}

void* QGraphicsItem_Data(void* ptr, int key)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QGraphicsObject*>(ptr)->data(key));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QGraphicsWidget*>(ptr)->data(key));
	} else {
		return new QVariant(static_cast<QGraphicsItem*>(ptr)->data(key));
	}
}

void* QGraphicsItem_Group(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->group();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->group();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->group();
	}
}

void QGraphicsItem_Hide(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->hide();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->hide();
	} else {
		static_cast<QGraphicsItem*>(ptr)->hide();
	}
}

void QGraphicsItem_KeyPressEvent(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
	}
}

void QGraphicsItem_KeyPressEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QGraphicsItemGroup*>(static_cast<QGraphicsItem*>(ptr))) {
		static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->QGraphicsItem::keyPressEvent(static_cast<QKeyEvent*>(event));
	}
}

void QGraphicsItem_KeyReleaseEvent(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
	}
}

void QGraphicsItem_KeyReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QGraphicsItemGroup*>(static_cast<QGraphicsItem*>(ptr))) {
		static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->QGraphicsItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	}
}

void QGraphicsItem_MouseDoubleClickEvent(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	}
}

void QGraphicsItem_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else if (dynamic_cast<QGraphicsItemGroup*>(static_cast<QGraphicsItem*>(ptr))) {
		static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->QGraphicsItem::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	}
}

void QGraphicsItem_MousePressEvent(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	}
}

void QGraphicsItem_MousePressEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else if (dynamic_cast<QGraphicsItemGroup*>(static_cast<QGraphicsItem*>(ptr))) {
		static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->QGraphicsItem::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	}
}

void QGraphicsItem_MouseReleaseEvent(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	}
}

void QGraphicsItem_MouseReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else if (dynamic_cast<QGraphicsItemGroup*>(static_cast<QGraphicsItem*>(ptr))) {
		static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else {
		static_cast<QGraphicsItem*>(ptr)->QGraphicsItem::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	}
}

void QGraphicsItem_Paint(void* ptr, void* painter, void* option, void* widget)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	} else {
		static_cast<QGraphicsItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	}
}

void* QGraphicsItem_Pos(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({ QPointF tmpValue = static_cast<QGraphicsObject*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QPointF tmpValue = static_cast<QGraphicsWidget*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
	} else {
		return ({ QPointF tmpValue = static_cast<QGraphicsItem*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
	}
}

void QGraphicsItem_Scroll(void* ptr, double dx, double dy, void* rect)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->scroll(dx, dy, *static_cast<QRectF*>(rect));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->scroll(dx, dy, *static_cast<QRectF*>(rect));
	} else {
		static_cast<QGraphicsItem*>(ptr)->scroll(dx, dy, *static_cast<QRectF*>(rect));
	}
}

void QGraphicsItem_SetActive(void* ptr, char active)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->setActive(active != 0);
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setActive(active != 0);
	} else {
		static_cast<QGraphicsItem*>(ptr)->setActive(active != 0);
	}
}

void QGraphicsItem_SetData(void* ptr, int key, void* value)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->setData(key, *static_cast<QVariant*>(value));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setData(key, *static_cast<QVariant*>(value));
	} else {
		static_cast<QGraphicsItem*>(ptr)->setData(key, *static_cast<QVariant*>(value));
	}
}

void QGraphicsItem_SetEnabled(void* ptr, char enabled)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->setEnabled(enabled != 0);
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setEnabled(enabled != 0);
	} else {
		static_cast<QGraphicsItem*>(ptr)->setEnabled(enabled != 0);
	}
}

void QGraphicsItem_SetFocus(void* ptr, long long focusReason)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->setFocus(static_cast<Qt::FocusReason>(focusReason));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setFocus(static_cast<Qt::FocusReason>(focusReason));
	} else {
		static_cast<QGraphicsItem*>(ptr)->setFocus(static_cast<Qt::FocusReason>(focusReason));
	}
}

void QGraphicsItem_SetPos(void* ptr, void* pos)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->setPos(*static_cast<QPointF*>(pos));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setPos(*static_cast<QPointF*>(pos));
	} else {
		static_cast<QGraphicsItem*>(ptr)->setPos(*static_cast<QPointF*>(pos));
	}
}

void QGraphicsItem_SetPos2(void* ptr, double x, double y)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->setPos(x, y);
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setPos(x, y);
	} else {
		static_cast<QGraphicsItem*>(ptr)->setPos(x, y);
	}
}

void QGraphicsItem_SetToolTip(void* ptr, struct QtWidgets_PackedString toolTip)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->setToolTip(QString::fromUtf8(toolTip.data, toolTip.len));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setToolTip(QString::fromUtf8(toolTip.data, toolTip.len));
	} else {
		static_cast<QGraphicsItem*>(ptr)->setToolTip(QString::fromUtf8(toolTip.data, toolTip.len));
	}
}

void QGraphicsItem_Show(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->show();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->show();
	} else {
		static_cast<QGraphicsItem*>(ptr)->show();
	}
}

struct QtWidgets_PackedString QGraphicsItem_ToolTip(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray* t7daf65 = new QByteArray(static_cast<QGraphicsObject*>(ptr)->toolTip().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t7daf65->prepend("WHITESPACE").constData()+10), t7daf65->size()-10, t7daf65 }; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray* t7daf65 = new QByteArray(static_cast<QGraphicsWidget*>(ptr)->toolTip().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t7daf65->prepend("WHITESPACE").constData()+10), t7daf65->size()-10, t7daf65 }; });
	} else {
		return ({ QByteArray* t7daf65 = new QByteArray(static_cast<QGraphicsItem*>(ptr)->toolTip().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t7daf65->prepend("WHITESPACE").constData()+10), t7daf65->size()-10, t7daf65 }; });
	}
}

void* QGraphicsItem_Transform(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QTransform(static_cast<QGraphicsObject*>(ptr)->transform());
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QTransform(static_cast<QGraphicsWidget*>(ptr)->transform());
	} else {
		return new QTransform(static_cast<QGraphicsItem*>(ptr)->transform());
	}
}

int QGraphicsItem_Type(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->type();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->type();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->type();
	}
}

int QGraphicsItem_TypeDefault(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::type();
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::type();
	} else if (dynamic_cast<QGraphicsItemGroup*>(static_cast<QGraphicsItem*>(ptr))) {
		return static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::type();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->QGraphicsItem::type();
	}
}

void QGraphicsItem_Update(void* ptr, void* rect)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->update(*static_cast<QRectF*>(rect));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->update(*static_cast<QRectF*>(rect));
	} else {
		static_cast<QGraphicsItem*>(ptr)->update(*static_cast<QRectF*>(rect));
	}
}

void QGraphicsItem_Update2(void* ptr, double x, double y, double width, double height)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsObject*>(ptr)->update(x, y, width, height);
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->update(x, y, width, height);
	} else {
		static_cast<QGraphicsItem*>(ptr)->update(x, y, width, height);
	}
}

void* QGraphicsItem_Window(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->window();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->window();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->window();
	}
}

double QGraphicsItem_X(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->x();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->x();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->x();
	}
}

double QGraphicsItem_Y(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsObject*>(ptr)->y();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->y();
	} else {
		return static_cast<QGraphicsItem*>(ptr)->y();
	}
}

void QGraphicsItem_DestroyQGraphicsItem(void* ptr)
{
	static_cast<QGraphicsItem*>(ptr)->~QGraphicsItem();
}

void QGraphicsItem_DestroyQGraphicsItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGraphicsItem___childItems_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsItem___childItems_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	}
}

void* QGraphicsItem___childItems_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else {
		return new QList<QGraphicsItem *>();
	}
}

void* QGraphicsItem___children_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsItem___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	}
}

void* QGraphicsItem___children_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else {
		return new QList<QGraphicsItem *>();
	}
}

void* QGraphicsItem___collidingItems_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsItem___collidingItems_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else {
			static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
		}
	}
}

void* QGraphicsItem___collidingItems_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsItem *>();
	} else {
		return new QList<QGraphicsItem *>();
	}
}

void* QGraphicsItem___setTransformations_transformations_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsItem___setTransformations_transformations_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	} else {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	}
}

void* QGraphicsItem___setTransformations_transformations_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsTransform *>();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsTransform *>();
	} else {
		return new QList<QGraphicsTransform *>();
	}
}

void* QGraphicsItem___transformations_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsItem___transformations_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	} else {
		static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
	}
}

void* QGraphicsItem___transformations_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsTransform *>();
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QGraphicsTransform *>();
	} else {
		return new QList<QGraphicsTransform *>();
	}
}

class MyQGraphicsItemGroup: public QGraphicsItemGroup
{
public:
	MyQGraphicsItemGroup(QGraphicsItem *parent = Q_NULLPTR) : QGraphicsItemGroup(parent) {QGraphicsItemGroup_QGraphicsItemGroup_QRegisterMetaType();};
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsItemGroup_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	int type() const { return callbackQGraphicsItem_Type(const_cast<void*>(static_cast<const void*>(this))); };
	 ~MyQGraphicsItemGroup() { callbackQGraphicsItemGroup_DestroyQGraphicsItemGroup(this); };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQGraphicsItem_ContextMenuEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQGraphicsItem_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQGraphicsItem_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MousePressEvent(this, event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MouseReleaseEvent(this, event); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsItemGroup_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
};

Q_DECLARE_METATYPE(QGraphicsItemGroup*)
Q_DECLARE_METATYPE(MyQGraphicsItemGroup*)

int QGraphicsItemGroup_QGraphicsItemGroup_QRegisterMetaType(){qRegisterMetaType<QGraphicsItemGroup*>(); return qRegisterMetaType<MyQGraphicsItemGroup*>();}

void* QGraphicsItemGroup_NewQGraphicsItemGroup(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsItemGroup(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsItemGroup(static_cast<QGraphicsWidget*>(parent));
	} else {
		return new MyQGraphicsItemGroup(static_cast<QGraphicsItem*>(parent));
	}
}

void* QGraphicsItemGroup_BoundingRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QGraphicsItemGroup*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsItemGroup_BoundingRectDefault(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QGraphicsItemGroup_DestroyQGraphicsItemGroup(void* ptr)
{
	static_cast<QGraphicsItemGroup*>(ptr)->~QGraphicsItemGroup();
}

void QGraphicsItemGroup_DestroyQGraphicsItemGroupDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QGraphicsItemGroup_Paint(void* ptr, void* painter, void* option, void* widget)
{
		static_cast<QGraphicsItemGroup*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsItemGroup_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
		static_cast<QGraphicsItemGroup*>(ptr)->QGraphicsItemGroup::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

class MyQGraphicsLayout: public QGraphicsLayout
{
public:
	MyQGraphicsLayout(QGraphicsLayoutItem *parent = Q_NULLPTR) : QGraphicsLayout(parent) {QGraphicsLayout_QGraphicsLayout_QRegisterMetaType();};
	int count() const { return callbackQGraphicsLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	void invalidate() { callbackQGraphicsLayout_Invalidate(this); };
	QGraphicsLayoutItem * itemAt(int i) const { return static_cast<QGraphicsLayoutItem*>(callbackQGraphicsLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), i)); };
	void removeAt(int index) { callbackQGraphicsLayout_RemoveAt(this, index); };
	 ~MyQGraphicsLayout() { callbackQGraphicsLayout_DestroyQGraphicsLayout(this); };
	void setGeometry(const QRectF & rect) { callbackQGraphicsLayoutItem_SetGeometry(this, const_cast<QRectF*>(&rect)); };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQGraphicsLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)), which, const_cast<QSizeF*>(&constraint))); };
};

Q_DECLARE_METATYPE(QGraphicsLayout*)
Q_DECLARE_METATYPE(MyQGraphicsLayout*)

int QGraphicsLayout_QGraphicsLayout_QRegisterMetaType(){qRegisterMetaType<QGraphicsLayout*>(); return qRegisterMetaType<MyQGraphicsLayout*>();}

void* QGraphicsLayout_NewQGraphicsLayout(void* parent)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsLayout(static_cast<QGraphicsWidget*>(parent));
	} else {
		return new MyQGraphicsLayout(static_cast<QGraphicsLayoutItem*>(parent));
	}
}

void QGraphicsLayout_Activate(void* ptr)
{
	static_cast<QGraphicsLayout*>(ptr)->activate();
}

int QGraphicsLayout_Count(void* ptr)
{
	return static_cast<QGraphicsLayout*>(ptr)->count();
}

void QGraphicsLayout_Invalidate(void* ptr)
{
	static_cast<QGraphicsLayout*>(ptr)->invalidate();
}

void QGraphicsLayout_InvalidateDefault(void* ptr)
{
		static_cast<QGraphicsLayout*>(ptr)->QGraphicsLayout::invalidate();
}

void* QGraphicsLayout_ItemAt(void* ptr, int i)
{
	return static_cast<QGraphicsLayout*>(ptr)->itemAt(i);
}

void QGraphicsLayout_RemoveAt(void* ptr, int index)
{
	static_cast<QGraphicsLayout*>(ptr)->removeAt(index);
}

void QGraphicsLayout_DestroyQGraphicsLayout(void* ptr)
{
	static_cast<QGraphicsLayout*>(ptr)->~QGraphicsLayout();
}

void QGraphicsLayout_DestroyQGraphicsLayoutDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGraphicsLayout_SizeHint(void* ptr, long long which, void* constraint)
{
	return ({ QSizeF tmpValue = static_cast<QGraphicsLayout*>(ptr)->sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsLayout_SizeHintDefault(void* ptr, long long which, void* constraint)
{
	Q_UNUSED(ptr);
	Q_UNUSED(which);
	Q_UNUSED(constraint);
	
}

class MyQGraphicsLayoutItem: public QGraphicsLayoutItem
{
public:
	MyQGraphicsLayoutItem(QGraphicsLayoutItem *parent = Q_NULLPTR, bool isLayout = false) : QGraphicsLayoutItem(parent, isLayout) {QGraphicsLayoutItem_QGraphicsLayoutItem_QRegisterMetaType();};
	void setGeometry(const QRectF & rect) { callbackQGraphicsLayoutItem_SetGeometry(this, const_cast<QRectF*>(&rect)); };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQGraphicsLayoutItem_SizeHint(const_cast<void*>(static_cast<const void*>(this)), which, const_cast<QSizeF*>(&constraint))); };
	 ~MyQGraphicsLayoutItem() { callbackQGraphicsLayoutItem_DestroyQGraphicsLayoutItem(this); };
};

Q_DECLARE_METATYPE(QGraphicsLayoutItem*)
Q_DECLARE_METATYPE(MyQGraphicsLayoutItem*)

int QGraphicsLayoutItem_QGraphicsLayoutItem_QRegisterMetaType(){qRegisterMetaType<QGraphicsLayoutItem*>(); return qRegisterMetaType<MyQGraphicsLayoutItem*>();}

void* QGraphicsLayoutItem_NewQGraphicsLayoutItem(void* parent, char isLayout)
{
	return new MyQGraphicsLayoutItem(static_cast<QGraphicsLayoutItem*>(parent), isLayout != 0);
}

void* QGraphicsLayoutItem_Geometry(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->geometry(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QRectF tmpValue = static_cast<QGraphicsLayoutItem*>(ptr)->geometry(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	}
}

void* QGraphicsLayoutItem_MaximumSize(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->maximumSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSizeF tmpValue = static_cast<QGraphicsLayoutItem*>(ptr)->maximumSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
	}
}

double QGraphicsLayoutItem_MaximumWidth(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->maximumWidth();
	} else {
		return static_cast<QGraphicsLayoutItem*>(ptr)->maximumWidth();
	}
}

double QGraphicsLayoutItem_MinimumHeight(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->minimumHeight();
	} else {
		return static_cast<QGraphicsLayoutItem*>(ptr)->minimumHeight();
	}
}

void* QGraphicsLayoutItem_MinimumSize(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->minimumSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSizeF tmpValue = static_cast<QGraphicsLayoutItem*>(ptr)->minimumSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
	}
}

double QGraphicsLayoutItem_MinimumWidth(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->minimumWidth();
	} else {
		return static_cast<QGraphicsLayoutItem*>(ptr)->minimumWidth();
	}
}

void QGraphicsLayoutItem_SetGeometry(void* ptr, void* rect)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
	} else {
		static_cast<QGraphicsLayoutItem*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
	}
}

void QGraphicsLayoutItem_SetGeometryDefault(void* ptr, void* rect)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::setGeometry(*static_cast<QRectF*>(rect));
	} else if (dynamic_cast<QGraphicsLayout*>(static_cast<QGraphicsLayoutItem*>(ptr))) {
		static_cast<QGraphicsLayout*>(ptr)->QGraphicsLayout::setGeometry(*static_cast<QRectF*>(rect));
	} else {
		static_cast<QGraphicsLayoutItem*>(ptr)->QGraphicsLayoutItem::setGeometry(*static_cast<QRectF*>(rect));
	}
}

void QGraphicsLayoutItem_SetMaximumWidth(void* ptr, double width)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setMaximumWidth(width);
	} else {
		static_cast<QGraphicsLayoutItem*>(ptr)->setMaximumWidth(width);
	}
}

void QGraphicsLayoutItem_SetMinimumSize(void* ptr, void* size)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setMinimumSize(*static_cast<QSizeF*>(size));
	} else {
		static_cast<QGraphicsLayoutItem*>(ptr)->setMinimumSize(*static_cast<QSizeF*>(size));
	}
}

void QGraphicsLayoutItem_SetMinimumSize2(void* ptr, double w, double h)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setMinimumSize(w, h);
	} else {
		static_cast<QGraphicsLayoutItem*>(ptr)->setMinimumSize(w, h);
	}
}

void QGraphicsLayoutItem_SetMinimumWidth(void* ptr, double width)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setMinimumWidth(width);
	} else {
		static_cast<QGraphicsLayoutItem*>(ptr)->setMinimumWidth(width);
	}
}

void QGraphicsLayoutItem_SetSizePolicy(void* ptr, void* policy)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setSizePolicy(*static_cast<QSizePolicy*>(policy));
	} else {
		static_cast<QGraphicsLayoutItem*>(ptr)->setSizePolicy(*static_cast<QSizePolicy*>(policy));
	}
}

void QGraphicsLayoutItem_SetSizePolicy2(void* ptr, long long hPolicy, long long vPolicy, long long controlType)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setSizePolicy(static_cast<QSizePolicy::Policy>(hPolicy), static_cast<QSizePolicy::Policy>(vPolicy), static_cast<QSizePolicy::ControlType>(controlType));
	} else {
		static_cast<QGraphicsLayoutItem*>(ptr)->setSizePolicy(static_cast<QSizePolicy::Policy>(hPolicy), static_cast<QSizePolicy::Policy>(vPolicy), static_cast<QSizePolicy::ControlType>(controlType));
	}
}

void* QGraphicsLayoutItem_SizeHint(void* ptr, long long which, void* constraint)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSizeF tmpValue = static_cast<QGraphicsLayoutItem*>(ptr)->sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
	}
}

void* QGraphicsLayoutItem_SizePolicy(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QSizePolicy(static_cast<QGraphicsWidget*>(ptr)->sizePolicy());
	} else {
		return new QSizePolicy(static_cast<QGraphicsLayoutItem*>(ptr)->sizePolicy());
	}
}

void QGraphicsLayoutItem_DestroyQGraphicsLayoutItem(void* ptr)
{
	static_cast<QGraphicsLayoutItem*>(ptr)->~QGraphicsLayoutItem();
}

void QGraphicsLayoutItem_DestroyQGraphicsLayoutItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQGraphicsObject: public QGraphicsObject
{
public:
	MyQGraphicsObject(QGraphicsItem *parent = Q_NULLPTR) : QGraphicsObject(parent) {QGraphicsObject_QGraphicsObject_QRegisterMetaType();};
	bool event(QEvent * ev) { return callbackQGraphicsObject_Event(this, ev) != 0; };
	 ~MyQGraphicsObject() { callbackQGraphicsObject_DestroyQGraphicsObject(this); };
	void childEvent(QChildEvent * event) { callbackQGraphicsObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGraphicsObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGraphicsObject_CustomEvent(this, event); };
	void deleteLater() { callbackQGraphicsObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGraphicsObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGraphicsObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGraphicsObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGraphicsObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsObject_TimerEvent(this, event); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsObject_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQGraphicsItem_ContextMenuEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQGraphicsItem_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQGraphicsItem_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MousePressEvent(this, event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MouseReleaseEvent(this, event); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsObject_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	int type() const { return callbackQGraphicsItem_Type(const_cast<void*>(static_cast<const void*>(this))); };
};

Q_DECLARE_METATYPE(QGraphicsObject*)
Q_DECLARE_METATYPE(MyQGraphicsObject*)

int QGraphicsObject_QGraphicsObject_QRegisterMetaType(){qRegisterMetaType<QGraphicsObject*>(); return qRegisterMetaType<MyQGraphicsObject*>();}

void* QGraphicsObject_NewQGraphicsObject(void* parent)
{
	return new MyQGraphicsObject(static_cast<QGraphicsItem*>(parent));
}

char QGraphicsObject_Event(void* ptr, void* ev)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->event(static_cast<QEvent*>(ev));
	} else {
		return static_cast<QGraphicsObject*>(ptr)->event(static_cast<QEvent*>(ev));
	}
}

char QGraphicsObject_EventDefault(void* ptr, void* ev)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::event(static_cast<QEvent*>(ev));
	} else {
		return static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::event(static_cast<QEvent*>(ev));
	}
}

void QGraphicsObject_DestroyQGraphicsObject(void* ptr)
{
	static_cast<QGraphicsObject*>(ptr)->~QGraphicsObject();
}

void QGraphicsObject_DestroyQGraphicsObjectDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QGraphicsObject_SetEnabled(void* ptr, char enabled)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setEnabled(enabled != 0);
	} else {
		static_cast<QGraphicsObject*>(ptr)->setEnabled(enabled != 0);
	}
}

void* QGraphicsObject_Parent(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->parent();
	} else {
		return static_cast<QGraphicsObject*>(ptr)->parent();
	}
}

void* QGraphicsObject_Pos(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QPointF tmpValue = static_cast<QGraphicsWidget*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
	} else {
		return ({ QPointF tmpValue = static_cast<QGraphicsObject*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
	}
}

void QGraphicsObject_SetPos(void* ptr, void* pos)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->setPos(*static_cast<QPointF*>(pos));
	} else {
		static_cast<QGraphicsObject*>(ptr)->setPos(*static_cast<QPointF*>(pos));
	}
}

double QGraphicsObject_Scale(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->scale();
	} else {
		return static_cast<QGraphicsObject*>(ptr)->scale();
	}
}

double QGraphicsObject_X(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->x();
	} else {
		return static_cast<QGraphicsObject*>(ptr)->x();
	}
}

double QGraphicsObject_Y(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGraphicsWidget*>(ptr)->y();
	} else {
		return static_cast<QGraphicsObject*>(ptr)->y();
	}
}

void* QGraphicsObject___children_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsObject___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
		}
	}
}

void* QGraphicsObject___children_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QObject *>();
	} else {
		return new QList<QObject *>();
	}
}

void* QGraphicsObject___dynamicPropertyNames_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
	} else {
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
	}
}

void QGraphicsObject___dynamicPropertyNames_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
	} else {
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
	}
}

void* QGraphicsObject___dynamicPropertyNames_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QByteArray>();
	} else {
		return new QList<QByteArray>();
	}
}

void* QGraphicsObject___findChildren_atList(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsObject___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
		}
	}
}

void* QGraphicsObject___findChildren_newList(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QObject*>();
	} else {
		return new QList<QObject*>();
	}
}

void* QGraphicsObject___findChildren_atList3(void* ptr, int i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
	} else {
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
	}
}

void QGraphicsObject___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
		} else {
			static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
		}
	}
}

void* QGraphicsObject___findChildren_newList3(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return new QList<QObject*>();
	} else {
		return new QList<QObject*>();
	}
}

void QGraphicsObject_ChildEvent(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QGraphicsObject*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
	}
}

void QGraphicsObject_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QGraphicsObject_ConnectNotify(void* ptr, void* sign)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QGraphicsObject*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QGraphicsObject_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QGraphicsObject_CustomEvent(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QGraphicsObject*>(ptr)->customEvent(static_cast<QEvent*>(event));
	}
}

void QGraphicsObject_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::customEvent(static_cast<QEvent*>(event));
	}
}

void QGraphicsObject_DeleteLater(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		QMetaObject::invokeMethod(static_cast<QGraphicsWidget*>(ptr), "deleteLater");
	} else {
		QMetaObject::invokeMethod(static_cast<QGraphicsObject*>(ptr), "deleteLater");
	}
}

void QGraphicsObject_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::deleteLater();
	} else {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::deleteLater();
	}
}

void QGraphicsObject_DisconnectNotify(void* ptr, void* sign)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QGraphicsObject*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QGraphicsObject_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QGraphicsObject_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsWidget*>(ptr)->eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsWidget*>(ptr)->eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsWidget*>(ptr)->eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsWidget*>(ptr)->eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QGraphicsWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsObject*>(ptr)->eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsObject*>(ptr)->eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsObject*>(ptr)->eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsObject*>(ptr)->eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QGraphicsObject*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

char QGraphicsObject_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void QGraphicsObject_TimerEvent(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QGraphicsObject*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void QGraphicsObject_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QGraphicsObject*>(ptr)->QGraphicsObject::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QGraphicsObject_BoundingRect(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
	
	}
}

void* QGraphicsObject_BoundingRectDefault(void* ptr)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
	
	}
}

void QGraphicsObject_Paint(void* ptr, void* painter, void* option, void* widget)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	} else {
	
	}
}

void QGraphicsObject_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
	if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	} else {
	
	}
}

class MyQGraphicsSceneContextMenuEvent: public QGraphicsSceneContextMenuEvent
{
public:
	 ~MyQGraphicsSceneContextMenuEvent() { callbackQGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEvent(this); };
};

Q_DECLARE_METATYPE(QGraphicsSceneContextMenuEvent*)
Q_DECLARE_METATYPE(MyQGraphicsSceneContextMenuEvent*)

int QGraphicsSceneContextMenuEvent_QGraphicsSceneContextMenuEvent_QRegisterMetaType(){qRegisterMetaType<QGraphicsSceneContextMenuEvent*>(); return qRegisterMetaType<MyQGraphicsSceneContextMenuEvent*>();}

void* QGraphicsSceneContextMenuEvent_Pos(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEvent(void* ptr)
{
	static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->~QGraphicsSceneContextMenuEvent();
}

void QGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEventDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQGraphicsSceneEvent: public QGraphicsSceneEvent
{
public:
	 ~MyQGraphicsSceneEvent() { callbackQGraphicsSceneEvent_DestroyQGraphicsSceneEvent(this); };
};

Q_DECLARE_METATYPE(QGraphicsSceneEvent*)
Q_DECLARE_METATYPE(MyQGraphicsSceneEvent*)

int QGraphicsSceneEvent_QGraphicsSceneEvent_QRegisterMetaType(){qRegisterMetaType<QGraphicsSceneEvent*>(); return qRegisterMetaType<MyQGraphicsSceneEvent*>();}

void* QGraphicsSceneEvent_Widget(void* ptr)
{
	return static_cast<QGraphicsSceneEvent*>(ptr)->widget();
}

void QGraphicsSceneEvent_DestroyQGraphicsSceneEvent(void* ptr)
{
	static_cast<QGraphicsSceneEvent*>(ptr)->~QGraphicsSceneEvent();
}

void QGraphicsSceneEvent_DestroyQGraphicsSceneEventDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQGraphicsSceneMouseEvent: public QGraphicsSceneMouseEvent
{
public:
	 ~MyQGraphicsSceneMouseEvent() { callbackQGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent(this); };
};

Q_DECLARE_METATYPE(QGraphicsSceneMouseEvent*)
Q_DECLARE_METATYPE(MyQGraphicsSceneMouseEvent*)

int QGraphicsSceneMouseEvent_QGraphicsSceneMouseEvent_QRegisterMetaType(){qRegisterMetaType<QGraphicsSceneMouseEvent*>(); return qRegisterMetaType<MyQGraphicsSceneMouseEvent*>();}

long long QGraphicsSceneMouseEvent_Button(void* ptr)
{
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->button();
}

long long QGraphicsSceneMouseEvent_Buttons(void* ptr)
{
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->buttons();
}

void* QGraphicsSceneMouseEvent_Pos(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QGraphicsSceneMouseEvent*>(ptr)->pos(); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent(void* ptr)
{
	static_cast<QGraphicsSceneMouseEvent*>(ptr)->~QGraphicsSceneMouseEvent();
}

void QGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEventDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQGraphicsSceneResizeEvent: public QGraphicsSceneResizeEvent
{
public:
	MyQGraphicsSceneResizeEvent() : QGraphicsSceneResizeEvent() {QGraphicsSceneResizeEvent_QGraphicsSceneResizeEvent_QRegisterMetaType();};
	 ~MyQGraphicsSceneResizeEvent() { callbackQGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEvent(this); };
};

Q_DECLARE_METATYPE(QGraphicsSceneResizeEvent*)
Q_DECLARE_METATYPE(MyQGraphicsSceneResizeEvent*)

int QGraphicsSceneResizeEvent_QGraphicsSceneResizeEvent_QRegisterMetaType(){qRegisterMetaType<QGraphicsSceneResizeEvent*>(); return qRegisterMetaType<MyQGraphicsSceneResizeEvent*>();}

void* QGraphicsSceneResizeEvent_NewQGraphicsSceneResizeEvent2()
{
	return new MyQGraphicsSceneResizeEvent();
}

void QGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEvent(void* ptr)
{
	static_cast<QGraphicsSceneResizeEvent*>(ptr)->~QGraphicsSceneResizeEvent();
}

void QGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEventDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQGraphicsTransform: public QGraphicsTransform
{
public:
	MyQGraphicsTransform(QObject *parent = Q_NULLPTR) : QGraphicsTransform(parent) {QGraphicsTransform_QGraphicsTransform_QRegisterMetaType();};
	void applyTo(QMatrix4x4 * matrix) const { callbackQGraphicsTransform_ApplyTo(const_cast<void*>(static_cast<const void*>(this)), matrix); };
	void update() { callbackQGraphicsTransform_Update(this); };
	 ~MyQGraphicsTransform() { callbackQGraphicsTransform_DestroyQGraphicsTransform(this); };
	void childEvent(QChildEvent * event) { callbackQGraphicsTransform_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGraphicsTransform_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGraphicsTransform_CustomEvent(this, event); };
	void deleteLater() { callbackQGraphicsTransform_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGraphicsTransform_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGraphicsTransform_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGraphicsTransform_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGraphicsTransform_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGraphicsTransform_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsTransform_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QGraphicsTransform*)
Q_DECLARE_METATYPE(MyQGraphicsTransform*)

int QGraphicsTransform_QGraphicsTransform_QRegisterMetaType(){qRegisterMetaType<QGraphicsTransform*>(); return qRegisterMetaType<MyQGraphicsTransform*>();}

void* QGraphicsTransform_NewQGraphicsTransform(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsTransform(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsTransform(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsTransform(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsTransform(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGraphicsTransform(static_cast<QWindow*>(parent));
	} else {
		return new MyQGraphicsTransform(static_cast<QObject*>(parent));
	}
}

void QGraphicsTransform_ApplyTo(void* ptr, void* matrix)
{
	static_cast<QGraphicsTransform*>(ptr)->applyTo(static_cast<QMatrix4x4*>(matrix));
}

void QGraphicsTransform_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsTransform*>(ptr), "update");
}

void QGraphicsTransform_UpdateDefault(void* ptr)
{
		static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::update();
}

void QGraphicsTransform_DestroyQGraphicsTransform(void* ptr)
{
	static_cast<QGraphicsTransform*>(ptr)->~QGraphicsTransform();
}

void QGraphicsTransform_DestroyQGraphicsTransformDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGraphicsTransform___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsTransform___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGraphicsTransform___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QGraphicsTransform___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGraphicsTransform___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGraphicsTransform___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGraphicsTransform___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsTransform___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGraphicsTransform___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGraphicsTransform___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsTransform___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGraphicsTransform___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QGraphicsTransform_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsTransform_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsTransform_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::customEvent(static_cast<QEvent*>(event));
}

void QGraphicsTransform_DeleteLaterDefault(void* ptr)
{
		static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::deleteLater();
}

void QGraphicsTransform_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGraphicsTransform_EventDefault(void* ptr, void* e)
{
		return static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::event(static_cast<QEvent*>(e));
}

char QGraphicsTransform_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
		return static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
		return static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QGraphicsTransform_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsTransform*>(ptr)->QGraphicsTransform::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQGraphicsWidget: public QGraphicsWidget
{
public:
	MyQGraphicsWidget(QGraphicsItem *parent = Q_NULLPTR, Qt::WindowFlags wFlags = Qt::WindowFlags()) : QGraphicsWidget(parent, wFlags) {QGraphicsWidget_QGraphicsWidget_QRegisterMetaType();};
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQGraphicsWidget_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQGraphicsWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQGraphicsWidget_CloseEvent(this, event); };
	bool event(QEvent * event) { return callbackQGraphicsObject_Event(this, event) != 0; };
	void hideEvent(QHideEvent * event) { callbackQGraphicsWidget_HideEvent(this, event); };
	void resizeEvent(QGraphicsSceneResizeEvent * event) { callbackQGraphicsWidget_ResizeEvent(this, event); };
	void setGeometry(const QRectF & rect) { callbackQGraphicsLayoutItem_SetGeometry(this, const_cast<QRectF*>(&rect)); };
	void showEvent(QShowEvent * event) { callbackQGraphicsWidget_ShowEvent(this, event); };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQGraphicsWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)), which, const_cast<QSizeF*>(&constraint))); };
	int type() const { return callbackQGraphicsItem_Type(const_cast<void*>(static_cast<const void*>(this))); };
	 ~MyQGraphicsWidget() { callbackQGraphicsWidget_DestroyQGraphicsWidget(this); };
	void childEvent(QChildEvent * event) { callbackQGraphicsObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGraphicsObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGraphicsObject_CustomEvent(this, event); };
	void deleteLater() { callbackQGraphicsObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGraphicsObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGraphicsObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGraphicsObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQGraphicsObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsObject_TimerEvent(this, event); };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQGraphicsItem_ContextMenuEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQGraphicsItem_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQGraphicsItem_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MousePressEvent(this, event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsItem_MouseReleaseEvent(this, event); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQGraphicsObject_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
};

Q_DECLARE_METATYPE(QGraphicsWidget*)
Q_DECLARE_METATYPE(MyQGraphicsWidget*)

int QGraphicsWidget_QGraphicsWidget_QRegisterMetaType(){qRegisterMetaType<QGraphicsWidget*>(); return qRegisterMetaType<MyQGraphicsWidget*>();}

void* QGraphicsWidget_NewQGraphicsWidget(void* parent, long long wFlags)
{
	return new MyQGraphicsWidget(static_cast<QGraphicsItem*>(parent), static_cast<Qt::WindowType>(wFlags));
}

struct QtWidgets_PackedList QGraphicsWidget_Actions(void* ptr)
{
		return ({ QList<QAction *>* tmpValuec4928c = new QList<QAction *>(static_cast<QGraphicsWidget*>(ptr)->actions()); QtWidgets_PackedList { tmpValuec4928c, tmpValuec4928c->size() }; });
}

void QGraphicsWidget_AddAction(void* ptr, void* action)
{
		static_cast<QGraphicsWidget*>(ptr)->addAction(static_cast<QAction*>(action));
}

void* QGraphicsWidget_BoundingRect(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsWidget_BoundingRectDefault(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

char QGraphicsWidget_Close(void* ptr)
{
		bool returnArg;
	QMetaObject::invokeMethod(static_cast<QGraphicsWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QGraphicsWidget_CloseDefault(void* ptr)
{
		return static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::close();
}

void QGraphicsWidget_CloseEvent(void* ptr, void* event)
{
		static_cast<QGraphicsWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QGraphicsWidget_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void* QGraphicsWidget_Font(void* ptr)
{
		return new QFont(static_cast<QGraphicsWidget*>(ptr)->font());
}

void QGraphicsWidget_HideEvent(void* ptr, void* event)
{
		static_cast<QGraphicsWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QGraphicsWidget_HideEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void* QGraphicsWidget_Layout(void* ptr)
{
		return static_cast<QGraphicsWidget*>(ptr)->layout();
}

void* QGraphicsWidget_Rect(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->rect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QGraphicsWidget_RemoveAction(void* ptr, void* action)
{
		static_cast<QGraphicsWidget*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QGraphicsWidget_Resize(void* ptr, void* size)
{
		static_cast<QGraphicsWidget*>(ptr)->resize(*static_cast<QSizeF*>(size));
}

void QGraphicsWidget_Resize2(void* ptr, double w, double h)
{
		static_cast<QGraphicsWidget*>(ptr)->resize(w, h);
}

void QGraphicsWidget_ResizeEvent(void* ptr, void* event)
{
		static_cast<QGraphicsWidget*>(ptr)->resizeEvent(static_cast<QGraphicsSceneResizeEvent*>(event));
}

void QGraphicsWidget_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::resizeEvent(static_cast<QGraphicsSceneResizeEvent*>(event));
}

void QGraphicsWidget_SetFont(void* ptr, void* font)
{
		static_cast<QGraphicsWidget*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsWidget_SetGeometry2(void* ptr, double x, double y, double w, double h)
{
		static_cast<QGraphicsWidget*>(ptr)->setGeometry(x, y, w, h);
}

void QGraphicsWidget_SetLayout(void* ptr, void* layout)
{
		static_cast<QGraphicsWidget*>(ptr)->setLayout(static_cast<QGraphicsLayout*>(layout));
}

void QGraphicsWidget_SetStyle(void* ptr, void* style)
{
		static_cast<QGraphicsWidget*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

void QGraphicsWidget_SetWindowTitle(void* ptr, struct QtWidgets_PackedString title)
{
		static_cast<QGraphicsWidget*>(ptr)->setWindowTitle(QString::fromUtf8(title.data, title.len));
}

void QGraphicsWidget_ShowEvent(void* ptr, void* event)
{
		static_cast<QGraphicsWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QGraphicsWidget_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::showEvent(static_cast<QShowEvent*>(event));
}

void* QGraphicsWidget_Size(void* ptr)
{
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->size(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsWidget_SizeHint(void* ptr, long long which, void* constraint)
{
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsWidget_SizeHintDefault(void* ptr, long long which, void* constraint)
{
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->QGraphicsWidget::sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsWidget_Style(void* ptr)
{
		return static_cast<QGraphicsWidget*>(ptr)->style();
}

struct QtWidgets_PackedString QGraphicsWidget_WindowTitle(void* ptr)
{
		return ({ QByteArray* t8cb41d = new QByteArray(static_cast<QGraphicsWidget*>(ptr)->windowTitle().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t8cb41d->prepend("WHITESPACE").constData()+10), t8cb41d->size()-10, t8cb41d }; });
}

void QGraphicsWidget_DestroyQGraphicsWidget(void* ptr)
{
	static_cast<QGraphicsWidget*>(ptr)->~QGraphicsWidget();
}

void QGraphicsWidget_DestroyQGraphicsWidgetDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGraphicsWidget_Geometry(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QGraphicsWidget*>(ptr)->geometry(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsWidget_MaximumSize(void* ptr)
{
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->maximumSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsWidget_MinimumSize(void* ptr)
{
		return ({ QSizeF tmpValue = static_cast<QGraphicsWidget*>(ptr)->minimumSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void QGraphicsWidget_SetMinimumSize(void* ptr, void* minimumSize)
{
		static_cast<QGraphicsWidget*>(ptr)->setMinimumSize(*static_cast<QSizeF*>(minimumSize));
}

void* QGraphicsWidget_SizePolicy(void* ptr)
{
		return new QSizePolicy(static_cast<QGraphicsWidget*>(ptr)->sizePolicy());
}

void QGraphicsWidget_SetSizePolicy(void* ptr, void* sizePolicy)
{
		static_cast<QGraphicsWidget*>(ptr)->setSizePolicy(*static_cast<QSizePolicy*>(sizePolicy));
}

void* QGraphicsWidget___actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsWidget___actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QGraphicsWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QGraphicsWidget___addActions_actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsWidget___addActions_actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QGraphicsWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QGraphicsWidget___insertActions_actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGraphicsWidget___insertActions_actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QGraphicsWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

class MyQGridLayout: public QGridLayout
{
public:
	MyQGridLayout(QWidget *parent) : QGridLayout(parent) {QGridLayout_QGridLayout_QRegisterMetaType();};
	MyQGridLayout() : QGridLayout() {QGridLayout_QGridLayout_QRegisterMetaType();};
	void addItem(QLayoutItem * item) { callbackQGridLayout_AddItem2(this, item); };
	int count() const { return callbackQGridLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayout_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQLayoutItem_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQLayoutItem_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void invalidate() { callbackQLayoutItem_Invalidate(this); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQGridLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayout_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	int minimumHeightForWidth(int w) const { return callbackQLayoutItem_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayout_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	void setGeometry(const QRect & rect) { callbackQLayout_SetGeometry(this, const_cast<QRect*>(&rect)); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQGridLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQGridLayout_TakeAt(this, index)); };
	 ~MyQGridLayout() { callbackQGridLayout_DestroyQGridLayout(this); };
	void childEvent(QChildEvent * e) { callbackQLayout_ChildEvent(this, e); };
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackQLayoutItem_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayout_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	int indexOf(QWidget * widget) const { return callbackQLayout_IndexOf(const_cast<void*>(static_cast<const void*>(this)), widget); };
	bool isEmpty() const { return callbackQLayout_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	void connectNotify(const QMetaMethod & sign) { callbackQLayout_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLayout_CustomEvent(this, event); };
	void deleteLater() { callbackQLayout_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLayout_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLayout_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQLayout_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLayout_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQLayout_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLayout_TimerEvent(this, event); };
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackQLayoutItem_SpacerItem(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
};

Q_DECLARE_METATYPE(QGridLayout*)
Q_DECLARE_METATYPE(MyQGridLayout*)

int QGridLayout_QGridLayout_QRegisterMetaType(){qRegisterMetaType<QGridLayout*>(); return qRegisterMetaType<MyQGridLayout*>();}

void* QGridLayout_NewQGridLayout(void* parent)
{
		return new MyQGridLayout(static_cast<QWidget*>(parent));
}

void* QGridLayout_NewQGridLayout2()
{
	return new MyQGridLayout();
}

void QGridLayout_AddItem(void* ptr, void* item, int row, int column, int rowSpan, int columnSpan, long long alignment)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
		static_cast<QGridLayout*>(ptr)->addItem(static_cast<QLayout*>(item), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
	} else {
		static_cast<QGridLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
	}
}

void QGridLayout_AddItem2(void* ptr, void* item)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
		static_cast<QGridLayout*>(ptr)->addItem(static_cast<QLayout*>(item));
	} else {
		static_cast<QGridLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
	}
}

void QGridLayout_AddItem2Default(void* ptr, void* item)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
		static_cast<QGridLayout*>(ptr)->QGridLayout::addItem(static_cast<QLayout*>(item));
	} else {
		static_cast<QGridLayout*>(ptr)->QGridLayout::addItem(static_cast<QLayoutItem*>(item));
	}
}

void QGridLayout_AddLayout(void* ptr, void* layout, int row, int column, long long alignment)
{
		static_cast<QGridLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddLayout2(void* ptr, void* layout, int row, int column, int rowSpan, int columnSpan, long long alignment)
{
		static_cast<QGridLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddWidget2(void* ptr, void* widget, int row, int column, long long alignment)
{
		static_cast<QGridLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddWidget3(void* ptr, void* widget, int fromRow, int fromColumn, int rowSpan, int columnSpan, long long alignment)
{
		static_cast<QGridLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), fromRow, fromColumn, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

int QGridLayout_ColumnCount(void* ptr)
{
	return static_cast<QGridLayout*>(ptr)->columnCount();
}

int QGridLayout_Count(void* ptr)
{
	return static_cast<QGridLayout*>(ptr)->count();
}

int QGridLayout_CountDefault(void* ptr)
{
		return static_cast<QGridLayout*>(ptr)->QGridLayout::count();
}

void* QGridLayout_ItemAt(void* ptr, int index)
{
	return static_cast<QGridLayout*>(ptr)->itemAt(index);
}

void* QGridLayout_ItemAtDefault(void* ptr, int index)
{
		return static_cast<QGridLayout*>(ptr)->QGridLayout::itemAt(index);
}

int QGridLayout_RowCount(void* ptr)
{
	return static_cast<QGridLayout*>(ptr)->rowCount();
}

void* QGridLayout_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QGridLayout*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QGridLayout_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QGridLayout*>(ptr)->QGridLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QGridLayout_TakeAt(void* ptr, int index)
{
	return static_cast<QGridLayout*>(ptr)->takeAt(index);
}

void* QGridLayout_TakeAtDefault(void* ptr, int index)
{
		return static_cast<QGridLayout*>(ptr)->QGridLayout::takeAt(index);
}

void QGridLayout_DestroyQGridLayout(void* ptr)
{
	static_cast<QGridLayout*>(ptr)->~QGridLayout();
}

void QGridLayout_DestroyQGridLayoutDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQHBoxLayout: public QHBoxLayout
{
public:
	MyQHBoxLayout() : QHBoxLayout() {QHBoxLayout_QHBoxLayout_QRegisterMetaType();};
	MyQHBoxLayout(QWidget *parent) : QHBoxLayout(parent) {QHBoxLayout_QHBoxLayout_QRegisterMetaType();};
	 ~MyQHBoxLayout() { callbackQHBoxLayout_DestroyQHBoxLayout(this); };
	void addItem(QLayoutItem * item) { callbackQBoxLayout_AddItem(this, item); };
	int count() const { return callbackQBoxLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayout_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQLayoutItem_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQLayoutItem_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void invalidate() { callbackQLayoutItem_Invalidate(this); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQBoxLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayout_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	int minimumHeightForWidth(int w) const { return callbackQLayoutItem_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayout_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	void setGeometry(const QRect & r) { callbackQLayout_SetGeometry(this, const_cast<QRect*>(&r)); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQBoxLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQBoxLayout_TakeAt(this, index)); };
	void childEvent(QChildEvent * e) { callbackQLayout_ChildEvent(this, e); };
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackQLayoutItem_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayout_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	int indexOf(QWidget * widget) const { return callbackQLayout_IndexOf(const_cast<void*>(static_cast<const void*>(this)), widget); };
	bool isEmpty() const { return callbackQLayout_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	void connectNotify(const QMetaMethod & sign) { callbackQLayout_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLayout_CustomEvent(this, event); };
	void deleteLater() { callbackQLayout_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLayout_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLayout_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQLayout_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLayout_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQLayout_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLayout_TimerEvent(this, event); };
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackQLayoutItem_SpacerItem(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
};

Q_DECLARE_METATYPE(QHBoxLayout*)
Q_DECLARE_METATYPE(MyQHBoxLayout*)

int QHBoxLayout_QHBoxLayout_QRegisterMetaType(){qRegisterMetaType<QHBoxLayout*>(); return qRegisterMetaType<MyQHBoxLayout*>();}

void* QHBoxLayout_NewQHBoxLayout()
{
	return new MyQHBoxLayout();
}

void* QHBoxLayout_NewQHBoxLayout2(void* parent)
{
		return new MyQHBoxLayout(static_cast<QWidget*>(parent));
}

void QHBoxLayout_DestroyQHBoxLayout(void* ptr)
{
	static_cast<QHBoxLayout*>(ptr)->~QHBoxLayout();
}

void QHBoxLayout_DestroyQHBoxLayoutDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQHeaderView: public QHeaderView
{
public:
	MyQHeaderView(Qt::Orientation orientation, QWidget *parent = Q_NULLPTR) : QHeaderView(orientation, parent) {QHeaderView_QHeaderView_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	void mouseDoubleClickEvent(QMouseEvent * e) { callbackQWidget_MouseDoubleClickEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	void setModel(QAbstractItemModel * model) { callbackQAbstractItemView_SetModel(this, model); };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags flags) { callbackQHeaderView_SetSelection(this, const_cast<QRect*>(&rect), flags); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQHeaderView() { callbackQHeaderView_DestroyQHeaderView(this); };
	void Signal_Activated(const QModelIndex & index) { callbackQAbstractItemView_Activated(this, const_cast<QModelIndex*>(&index)); };
	void clearSelection() { callbackQAbstractItemView_ClearSelection(this); };
	void Signal_Clicked(const QModelIndex & index) { callbackQAbstractItemView_Clicked(this, const_cast<QModelIndex*>(&index)); };
	void edit(const QModelIndex & index) { callbackQAbstractItemView_Edit(this, const_cast<QModelIndex*>(&index)); };
	bool edit(const QModelIndex & index, QAbstractItemView::EditTrigger trigger, QEvent * event) { return callbackQAbstractItemView_Edit2(this, const_cast<QModelIndex*>(&index), trigger, event) != 0; };
	bool eventFilter(QObject * object, QEvent * event) { return callbackQWidget_EventFilter(this, object, event) != 0; };
	int horizontalOffset() const { return callbackQHeaderView_HorizontalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	QModelIndex indexAt(const QPoint & point) const { return *static_cast<QModelIndex*>(callbackQHeaderView_IndexAt(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPoint*>(&point))); };
	bool isIndexHidden(const QModelIndex & index) const { return callbackQHeaderView_IsIndexHidden(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index)) != 0; };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, Qt::KeyboardModifiers modifiers) { return *static_cast<QModelIndex*>(callbackQHeaderView_MoveCursor(this, cursorAction, modifiers)); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void scrollTo(const QModelIndex & index, QAbstractItemView::ScrollHint hint) { callbackQHeaderView_ScrollTo(this, const_cast<QModelIndex*>(&index), hint); };
	void selectionChanged(const QItemSelection & selected, const QItemSelection & deselected) { callbackQAbstractItemView_SelectionChanged(this, const_cast<QItemSelection*>(&selected), const_cast<QItemSelection*>(&deselected)); };
	void setCurrentIndex(const QModelIndex & index) { callbackQAbstractItemView_SetCurrentIndex(this, const_cast<QModelIndex*>(&index)); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	void update(const QModelIndex & index) { callbackQAbstractItemView_Update(this, const_cast<QModelIndex*>(&index)); };
	int verticalOffset() const { return callbackQHeaderView_VerticalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	QRect visualRect(const QModelIndex & index) const { return *static_cast<QRect*>(callbackQHeaderView_VisualRect(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QRegion visualRegionForSelection(const QItemSelection & selection) const { return *static_cast<QRegion*>(callbackQHeaderView_VisualRegionForSelection(const_cast<void*>(static_cast<const void*>(this)), const_cast<QItemSelection*>(&selection))); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QHeaderView*)
Q_DECLARE_METATYPE(MyQHeaderView*)

int QHeaderView_QHeaderView_QRegisterMetaType(){qRegisterMetaType<QHeaderView*>(); return qRegisterMetaType<MyQHeaderView*>();}

void* QHeaderView_NewQHeaderView(long long orientation, void* parent)
{
		return new MyQHeaderView(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

int QHeaderView_Count(void* ptr)
{
	return static_cast<QHeaderView*>(ptr)->count();
}

int QHeaderView_Length(void* ptr)
{
	return static_cast<QHeaderView*>(ptr)->length();
}

long long QHeaderView_Orientation(void* ptr)
{
	return static_cast<QHeaderView*>(ptr)->orientation();
}

void QHeaderView_SetSelection(void* ptr, void* rect, long long flags)
{
	static_cast<QHeaderView*>(ptr)->setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(flags));
}

void QHeaderView_SetSelectionDefault(void* ptr, void* rect, long long flags)
{
		static_cast<QHeaderView*>(ptr)->QHeaderView::setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(flags));
}

void QHeaderView_DestroyQHeaderView(void* ptr)
{
	static_cast<QHeaderView*>(ptr)->~QHeaderView();
}

void QHeaderView_DestroyQHeaderViewDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

int QHeaderView_HorizontalOffset(void* ptr)
{
	return static_cast<QHeaderView*>(ptr)->horizontalOffset();
}

int QHeaderView_HorizontalOffsetDefault(void* ptr)
{
		return static_cast<QHeaderView*>(ptr)->QHeaderView::horizontalOffset();
}

void* QHeaderView_IndexAt(void* ptr, void* point)
{
	return new QModelIndex(static_cast<QHeaderView*>(ptr)->indexAt(*static_cast<QPoint*>(point)));
}

void* QHeaderView_IndexAtDefault(void* ptr, void* point)
{
		return new QModelIndex(static_cast<QHeaderView*>(ptr)->QHeaderView::indexAt(*static_cast<QPoint*>(point)));
}

char QHeaderView_IsIndexHidden(void* ptr, void* index)
{
	return static_cast<QHeaderView*>(ptr)->isIndexHidden(*static_cast<QModelIndex*>(index));
}

char QHeaderView_IsIndexHiddenDefault(void* ptr, void* index)
{
		return static_cast<QHeaderView*>(ptr)->QHeaderView::isIndexHidden(*static_cast<QModelIndex*>(index));
}

void* QHeaderView_MoveCursor(void* ptr, long long cursorAction, long long modifiers)
{
	return new QModelIndex(static_cast<QHeaderView*>(ptr)->moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

void* QHeaderView_MoveCursorDefault(void* ptr, long long cursorAction, long long modifiers)
{
		return new QModelIndex(static_cast<QHeaderView*>(ptr)->QHeaderView::moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

void QHeaderView_ScrollTo(void* ptr, void* index, long long hint)
{
	static_cast<QHeaderView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QHeaderView_ScrollToDefault(void* ptr, void* index, long long hint)
{
		static_cast<QHeaderView*>(ptr)->QHeaderView::scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

int QHeaderView_VerticalOffset(void* ptr)
{
	return static_cast<QHeaderView*>(ptr)->verticalOffset();
}

int QHeaderView_VerticalOffsetDefault(void* ptr)
{
		return static_cast<QHeaderView*>(ptr)->QHeaderView::verticalOffset();
}

void* QHeaderView_VisualRect(void* ptr, void* index)
{
	return ({ QRect tmpValue = static_cast<QHeaderView*>(ptr)->visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QHeaderView_VisualRectDefault(void* ptr, void* index)
{
		return ({ QRect tmpValue = static_cast<QHeaderView*>(ptr)->QHeaderView::visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QHeaderView_VisualRegionForSelection(void* ptr, void* selection)
{
	return new QRegion(static_cast<QHeaderView*>(ptr)->visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

void* QHeaderView_VisualRegionForSelectionDefault(void* ptr, void* selection)
{
		return new QRegion(static_cast<QHeaderView*>(ptr)->QHeaderView::visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

class MyQInputDialog: public QInputDialog
{
public:
	MyQInputDialog(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QInputDialog(parent, flags) {QInputDialog_QInputDialog_QRegisterMetaType();};
	void done(int result) { callbackQInputDialog_Done(this, result); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQInputDialog() { callbackQInputDialog_DestroyQInputDialog(this); };
	void accept() { callbackQDialog_Accept(this); };
	void closeEvent(QCloseEvent * e) { callbackQWidget_CloseEvent(this, e); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	bool eventFilter(QObject * o, QEvent * e) { return callbackQWidget_EventFilter(this, o, e) != 0; };
	int exec() { return callbackQDialog_Exec(this); };
	void Signal_Finished(int result) { callbackQDialog_Finished(this, result); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void resizeEvent(QResizeEvent * vqr) { callbackQWidget_ResizeEvent(this, vqr); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QInputDialog*)
Q_DECLARE_METATYPE(MyQInputDialog*)

int QInputDialog_QInputDialog_QRegisterMetaType(){qRegisterMetaType<QInputDialog*>(); return qRegisterMetaType<MyQInputDialog*>();}

void* QInputDialog_NewQInputDialog(void* parent, long long flags)
{
		return new MyQInputDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QInputDialog_Done(void* ptr, int result)
{
	static_cast<QInputDialog*>(ptr)->done(result);
}

void QInputDialog_DoneDefault(void* ptr, int result)
{
		static_cast<QInputDialog*>(ptr)->QInputDialog::done(result);
}

struct QtWidgets_PackedString QInputDialog_QInputDialog_GetItem(void* parent, struct QtWidgets_PackedString title, struct QtWidgets_PackedString label, struct QtWidgets_PackedString items, int current, char editable, char* ok, long long flags, long long inputMethodHints)
{
		return ({ QByteArray* t589f5f = new QByteArray(QInputDialog::getItem(static_cast<QWidget*>(parent), QString::fromUtf8(title.data, title.len), QString::fromUtf8(label.data, label.len), QString::fromUtf8(items.data, items.len).split("¡¦!", QString::SkipEmptyParts), current, editable != 0, reinterpret_cast<bool*>(ok), static_cast<Qt::WindowType>(flags), static_cast<Qt::InputMethodHint>(inputMethodHints)).toUtf8()); QtWidgets_PackedString { const_cast<char*>(t589f5f->prepend("WHITESPACE").constData()+10), t589f5f->size()-10, t589f5f }; });
}

struct QtWidgets_PackedString QInputDialog_QInputDialog_GetText(void* parent, struct QtWidgets_PackedString title, struct QtWidgets_PackedString label, long long mode, struct QtWidgets_PackedString text, char* ok, long long flags, long long inputMethodHints)
{
		return ({ QByteArray* t03d5ef = new QByteArray(QInputDialog::getText(static_cast<QWidget*>(parent), QString::fromUtf8(title.data, title.len), QString::fromUtf8(label.data, label.len), static_cast<QLineEdit::EchoMode>(mode), QString::fromUtf8(text.data, text.len), reinterpret_cast<bool*>(ok), static_cast<Qt::WindowType>(flags), static_cast<Qt::InputMethodHint>(inputMethodHints)).toUtf8()); QtWidgets_PackedString { const_cast<char*>(t03d5ef->prepend("WHITESPACE").constData()+10), t03d5ef->size()-10, t03d5ef }; });
}

void QInputDialog_Open(void* ptr, void* receiver, char* member)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(receiver))) {
		static_cast<QInputDialog*>(ptr)->open(static_cast<QGraphicsObject*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(receiver))) {
		static_cast<QInputDialog*>(ptr)->open(static_cast<QGraphicsWidget*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(receiver))) {
		static_cast<QInputDialog*>(ptr)->open(static_cast<QLayout*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(receiver))) {
		static_cast<QInputDialog*>(ptr)->open(static_cast<QWidget*>(receiver), const_cast<const char*>(member));
	} else {
		static_cast<QInputDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
	}
}

void QInputDialog_DestroyQInputDialog(void* ptr)
{
	static_cast<QInputDialog*>(ptr)->~QInputDialog();
}

void QInputDialog_DestroyQInputDialogDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQLabel: public QLabel
{
public:
	MyQLabel(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QLabel(parent, ff) {QLabel_QLabel_QRegisterMetaType();};
	MyQLabel(const QString &text, QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QLabel(text, parent, ff) {QLabel_QLabel_QRegisterMetaType();};
	void clear() { callbackQLabel_Clear(this); };
	void contextMenuEvent(QContextMenuEvent * ev) { callbackQWidget_ContextMenuEvent(this, ev); };
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void keyPressEvent(QKeyEvent * ev) { callbackQWidget_KeyPressEvent(this, ev); };
	void mousePressEvent(QMouseEvent * ev) { callbackQWidget_MousePressEvent(this, ev); };
	void mouseReleaseEvent(QMouseEvent * ev) { callbackQWidget_MouseReleaseEvent(this, ev); };
	void setPixmap(const QPixmap & vqp) { callbackQLabel_SetPixmap(this, const_cast<QPixmap*>(&vqp)); };
	void setText(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQLabel_SetText(this, vqsPacked); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQLabel() { callbackQLabel_DestroyQLabel(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QLabel*)
Q_DECLARE_METATYPE(MyQLabel*)

int QLabel_QLabel_QRegisterMetaType(){qRegisterMetaType<QLabel*>(); return qRegisterMetaType<MyQLabel*>();}

void* QLabel_NewQLabel(void* parent, long long ff)
{
		return new MyQLabel(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void* QLabel_NewQLabel2(struct QtWidgets_PackedString text, void* parent, long long ff)
{
		return new MyQLabel(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

long long QLabel_Alignment(void* ptr)
{
	return static_cast<QLabel*>(ptr)->alignment();
}

void QLabel_Clear(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "clear");
}

void QLabel_ClearDefault(void* ptr)
{
		static_cast<QLabel*>(ptr)->QLabel::clear();
}

int QLabel_Indent(void* ptr)
{
	return static_cast<QLabel*>(ptr)->indent();
}

int QLabel_Margin(void* ptr)
{
	return static_cast<QLabel*>(ptr)->margin();
}

struct QtWidgets_PackedString QLabel_SelectedText(void* ptr)
{
	return ({ QByteArray* t8b2afb = new QByteArray(static_cast<QLabel*>(ptr)->selectedText().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t8b2afb->prepend("WHITESPACE").constData()+10), t8b2afb->size()-10, t8b2afb }; });
}

int QLabel_SelectionStart(void* ptr)
{
	return static_cast<QLabel*>(ptr)->selectionStart();
}

void QLabel_SetAlignment(void* ptr, long long vqt)
{
	static_cast<QLabel*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(vqt));
}

void QLabel_SetIndent(void* ptr, int vin)
{
	static_cast<QLabel*>(ptr)->setIndent(vin);
}

void QLabel_SetPixmap(void* ptr, void* vqp)
{
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setPixmap", Q_ARG(const QPixmap, *static_cast<QPixmap*>(vqp)));
}

void QLabel_SetPixmapDefault(void* ptr, void* vqp)
{
		static_cast<QLabel*>(ptr)->QLabel::setPixmap(*static_cast<QPixmap*>(vqp));
}

void QLabel_SetSelection(void* ptr, int start, int length)
{
	static_cast<QLabel*>(ptr)->setSelection(start, length);
}

void QLabel_SetText(void* ptr, struct QtWidgets_PackedString vqs)
{
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setText", Q_ARG(const QString, QString::fromUtf8(vqs.data, vqs.len)));
}

void QLabel_SetTextDefault(void* ptr, struct QtWidgets_PackedString vqs)
{
		static_cast<QLabel*>(ptr)->QLabel::setText(QString::fromUtf8(vqs.data, vqs.len));
}

struct QtWidgets_PackedString QLabel_Text(void* ptr)
{
	return ({ QByteArray* t0d779d = new QByteArray(static_cast<QLabel*>(ptr)->text().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t0d779d->prepend("WHITESPACE").constData()+10), t0d779d->size()-10, t0d779d }; });
}

void QLabel_DestroyQLabel(void* ptr)
{
	static_cast<QLabel*>(ptr)->~QLabel();
}

void QLabel_DestroyQLabelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQLayout: public QLayout
{
public:
	MyQLayout(QWidget *parent) : QLayout(parent) {QLayout_QLayout_QRegisterMetaType();};
	MyQLayout() : QLayout() {QLayout_QLayout_QRegisterMetaType();};
	void addItem(QLayoutItem * item) { callbackQLayout_AddItem(this, item); };
	void childEvent(QChildEvent * e) { callbackQLayout_ChildEvent(this, e); };
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackQLayoutItem_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	int count() const { return callbackQLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayout_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayout_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	int indexOf(QWidget * widget) const { return callbackQLayout_IndexOf(const_cast<void*>(static_cast<const void*>(this)), widget); };
	void invalidate() { callbackQLayoutItem_Invalidate(this); };
	bool isEmpty() const { return callbackQLayout_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayout_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayout_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	void setGeometry(const QRect & r) { callbackQLayout_SetGeometry(this, const_cast<QRect*>(&r)); };
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQLayout_TakeAt(this, index)); };
	void connectNotify(const QMetaMethod & sign) { callbackQLayout_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLayout_CustomEvent(this, event); };
	void deleteLater() { callbackQLayout_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLayout_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLayout_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQLayout_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLayout_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQLayout_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLayout_TimerEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQLayoutItem_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int vin) const { return callbackQLayoutItem_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), vin); };
	int minimumHeightForWidth(int w) const { return callbackQLayoutItem_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackQLayoutItem_SpacerItem(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
};

Q_DECLARE_METATYPE(QLayout*)
Q_DECLARE_METATYPE(MyQLayout*)

int QLayout_QLayout_QRegisterMetaType(){qRegisterMetaType<QLayout*>(); return qRegisterMetaType<MyQLayout*>();}

void* QLayout_NewQLayout(void* parent)
{
	return new MyQLayout(static_cast<QWidget*>(parent));
}

void* QLayout_NewQLayout2()
{
	return new MyQLayout();
}

char QLayout_Activate(void* ptr)
{
		return static_cast<QLayout*>(ptr)->activate();
}

void QLayout_AddItem(void* ptr, void* item)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(item))) {
		static_cast<QLayout*>(ptr)->addItem(static_cast<QLayout*>(item));
	} else {
		static_cast<QLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
	}
}

void QLayout_AddWidget(void* ptr, void* w)
{
		static_cast<QLayout*>(ptr)->addWidget(static_cast<QWidget*>(w));
}

void QLayout_ChildEvent(void* ptr, void* e)
{
		static_cast<QLayout*>(ptr)->childEvent(static_cast<QChildEvent*>(e));
}

void QLayout_ChildEventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QGridLayout*>(ptr)->QGridLayout::childEvent(static_cast<QChildEvent*>(e));
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::childEvent(static_cast<QChildEvent*>(e));
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::childEvent(static_cast<QChildEvent*>(e));
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxLayout*>(ptr)->QBoxLayout::childEvent(static_cast<QChildEvent*>(e));
	} else {
		static_cast<QLayout*>(ptr)->QLayout::childEvent(static_cast<QChildEvent*>(e));
	}
}

int QLayout_Count(void* ptr)
{
		return static_cast<QLayout*>(ptr)->count();
}

long long QLayout_ExpandingDirections(void* ptr)
{
		return static_cast<QLayout*>(ptr)->expandingDirections();
}

long long QLayout_ExpandingDirectionsDefault(void* ptr)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::expandingDirections();
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::expandingDirections();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::expandingDirections();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::expandingDirections();
	} else {
		return static_cast<QLayout*>(ptr)->QLayout::expandingDirections();
	}
}

void* QLayout_Geometry(void* ptr)
{
		return ({ QRect tmpValue = static_cast<QLayout*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QLayout_GeometryDefault(void* ptr)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QRect tmpValue = static_cast<QGridLayout*>(ptr)->QGridLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QRect tmpValue = static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QRect tmpValue = static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QRect tmpValue = static_cast<QBoxLayout*>(ptr)->QBoxLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QRect tmpValue = static_cast<QLayout*>(ptr)->QLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	}
}

int QLayout_IndexOf(void* ptr, void* widget)
{
		return static_cast<QLayout*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

int QLayout_IndexOfDefault(void* ptr, void* widget)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::indexOf(static_cast<QWidget*>(widget));
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::indexOf(static_cast<QWidget*>(widget));
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::indexOf(static_cast<QWidget*>(widget));
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::indexOf(static_cast<QWidget*>(widget));
	} else {
		return static_cast<QLayout*>(ptr)->QLayout::indexOf(static_cast<QWidget*>(widget));
	}
}

int QLayout_IndexOf2(void* ptr, void* layoutItem)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(layoutItem))) {
		return static_cast<QLayout*>(ptr)->indexOf(static_cast<QLayout*>(layoutItem));
	} else {
		return static_cast<QLayout*>(ptr)->indexOf(static_cast<QLayoutItem*>(layoutItem));
	}
}

char QLayout_IsEmpty(void* ptr)
{
		return static_cast<QLayout*>(ptr)->isEmpty();
}

char QLayout_IsEmptyDefault(void* ptr)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::isEmpty();
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::isEmpty();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::isEmpty();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::isEmpty();
	} else {
		return static_cast<QLayout*>(ptr)->QLayout::isEmpty();
	}
}

void* QLayout_ItemAt(void* ptr, int index)
{
		return static_cast<QLayout*>(ptr)->itemAt(index);
}

void* QLayout_MaximumSize(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QLayout_MaximumSizeDefault(void* ptr)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QGridLayout*>(ptr)->QGridLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QBoxLayout*>(ptr)->QBoxLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->QLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

void* QLayout_MenuBar(void* ptr)
{
		return static_cast<QLayout*>(ptr)->menuBar();
}

void* QLayout_MinimumSize(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QLayout_MinimumSizeDefault(void* ptr)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QGridLayout*>(ptr)->QGridLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QBoxLayout*>(ptr)->QBoxLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->QLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

char QLayout_SetAlignment(void* ptr, void* w, long long alignment)
{
		return static_cast<QLayout*>(ptr)->setAlignment(static_cast<QWidget*>(w), static_cast<Qt::AlignmentFlag>(alignment));
}

char QLayout_SetAlignment2(void* ptr, void* l, long long alignment)
{
		return static_cast<QLayout*>(ptr)->setAlignment(static_cast<QLayout*>(l), static_cast<Qt::AlignmentFlag>(alignment));
}

void QLayout_SetEnabled(void* ptr, char enable)
{
		static_cast<QLayout*>(ptr)->setEnabled(enable != 0);
}

void QLayout_SetGeometry(void* ptr, void* r)
{
		static_cast<QLayout*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void QLayout_SetGeometryDefault(void* ptr, void* r)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QGridLayout*>(ptr)->QGridLayout::setGeometry(*static_cast<QRect*>(r));
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::setGeometry(*static_cast<QRect*>(r));
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::setGeometry(*static_cast<QRect*>(r));
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxLayout*>(ptr)->QBoxLayout::setGeometry(*static_cast<QRect*>(r));
	} else {
		static_cast<QLayout*>(ptr)->QLayout::setGeometry(*static_cast<QRect*>(r));
	}
}

void* QLayout_TakeAt(void* ptr, int index)
{
		return static_cast<QLayout*>(ptr)->takeAt(index);
}

void QLayout_Update(void* ptr)
{
		static_cast<QLayout*>(ptr)->update();
}

void* QLayout___children_atList(void* ptr, int i)
{
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLayout___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QLayout___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>();
}

void* QLayout___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QLayout___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QLayout___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>();
}

void* QLayout___findChildren_atList(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLayout___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QLayout___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QLayout___findChildren_atList3(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLayout___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QLayout___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void QLayout_ConnectNotify(void* ptr, void* sign)
{
		static_cast<QLayout*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLayout_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QGridLayout*>(ptr)->QGridLayout::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxLayout*>(ptr)->QBoxLayout::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QLayout*>(ptr)->QLayout::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QLayout_CustomEvent(void* ptr, void* event)
{
		static_cast<QLayout*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLayout_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QGridLayout*>(ptr)->QGridLayout::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxLayout*>(ptr)->QBoxLayout::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QLayout*>(ptr)->QLayout::customEvent(static_cast<QEvent*>(event));
	}
}

void QLayout_DeleteLater(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QLayout*>(ptr), "deleteLater");
}

void QLayout_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QGridLayout*>(ptr)->QGridLayout::deleteLater();
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::deleteLater();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::deleteLater();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxLayout*>(ptr)->QBoxLayout::deleteLater();
	} else {
		static_cast<QLayout*>(ptr)->QLayout::deleteLater();
	}
}

void QLayout_DisconnectNotify(void* ptr, void* sign)
{
		static_cast<QLayout*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLayout_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QGridLayout*>(ptr)->QGridLayout::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxLayout*>(ptr)->QBoxLayout::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QLayout*>(ptr)->QLayout::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QLayout_Event(void* ptr, void* e)
{
		return static_cast<QLayout*>(ptr)->event(static_cast<QEvent*>(e));
}

char QLayout_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QLayout*>(ptr)->QLayout::event(static_cast<QEvent*>(e));
	}
}

char QLayout_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
		return static_cast<QLayout*>(ptr)->eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QLayout*>(ptr)->eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
		return static_cast<QLayout*>(ptr)->eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QLayout*>(ptr)->eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QLayout*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QLayout_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QGridLayout*>(ptr)->QGridLayout::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QGridLayout*>(ptr)->QGridLayout::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QGridLayout*>(ptr)->QGridLayout::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QGridLayout*>(ptr)->QGridLayout::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QGridLayout*>(ptr)->QGridLayout::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QBoxLayout*>(ptr)->QBoxLayout::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QBoxLayout*>(ptr)->QBoxLayout::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QBoxLayout*>(ptr)->QBoxLayout::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QBoxLayout*>(ptr)->QBoxLayout::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QBoxLayout*>(ptr)->QBoxLayout::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QLayout*>(ptr)->QLayout::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QLayout*>(ptr)->QLayout::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QLayout*>(ptr)->QLayout::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QLayout*>(ptr)->QLayout::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QLayout*>(ptr)->QLayout::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void QLayout_TimerEvent(void* ptr, void* event)
{
		static_cast<QLayout*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLayout_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QGridLayout*>(ptr)->QGridLayout::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxLayout*>(ptr)->QBoxLayout::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QLayout*>(ptr)->QLayout::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QLayout_SizeHint(void* ptr)
{
	Q_UNUSED(ptr);
	
}

void* QLayout_SizeHintDefault(void* ptr)
{
	if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QGridLayout*>(ptr)->QGridLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QBoxLayout*>(ptr)->QBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
	
	}
}

class MyQLayoutItem: public QLayoutItem
{
public:
	MyQLayoutItem(Qt::Alignment alignment = Qt::Alignment()) : QLayoutItem(alignment) {QLayoutItem_QLayoutItem_QRegisterMetaType();};
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackQLayoutItem_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayoutItem_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayoutItem_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQLayoutItem_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int vin) const { return callbackQLayoutItem_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), vin); };
	void invalidate() { callbackQLayoutItem_Invalidate(this); };
	bool isEmpty() const { return callbackQLayoutItem_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayoutItem_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	int minimumHeightForWidth(int w) const { return callbackQLayoutItem_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayoutItem_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	void setGeometry(const QRect & r) { callbackQLayoutItem_SetGeometry(this, const_cast<QRect*>(&r)); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQLayoutItem_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackQLayoutItem_SpacerItem(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
	 ~MyQLayoutItem() { callbackQLayoutItem_DestroyQLayoutItem(this); };
};

Q_DECLARE_METATYPE(QLayoutItem*)
Q_DECLARE_METATYPE(MyQLayoutItem*)

int QLayoutItem_QLayoutItem_QRegisterMetaType(){qRegisterMetaType<QLayoutItem*>(); return qRegisterMetaType<MyQLayoutItem*>();}

void* QLayoutItem_NewQLayoutItem(long long alignment)
{
	return new MyQLayoutItem(static_cast<Qt::AlignmentFlag>(alignment));
}

long long QLayoutItem_Alignment(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->alignment();
	} else {
		return static_cast<QLayoutItem*>(ptr)->alignment();
	}
}

long long QLayoutItem_ControlTypes(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->controlTypes();
	} else {
		return static_cast<QLayoutItem*>(ptr)->controlTypes();
	}
}

long long QLayoutItem_ControlTypesDefault(void* ptr)
{
	if (dynamic_cast<QSpacerItem*>(static_cast<QLayoutItem*>(ptr))) {
		return static_cast<QSpacerItem*>(ptr)->QSpacerItem::controlTypes();
	} else if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::controlTypes();
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::controlTypes();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::controlTypes();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::controlTypes();
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->QLayout::controlTypes();
	} else {
		return static_cast<QLayoutItem*>(ptr)->QLayoutItem::controlTypes();
	}
}

long long QLayoutItem_ExpandingDirections(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->expandingDirections();
	} else {
		return static_cast<QLayoutItem*>(ptr)->expandingDirections();
	}
}

void* QLayoutItem_Geometry(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QRect tmpValue = static_cast<QLayout*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QRect tmpValue = static_cast<QLayoutItem*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	}
}

char QLayoutItem_HasHeightForWidth(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->hasHeightForWidth();
	} else {
		return static_cast<QLayoutItem*>(ptr)->hasHeightForWidth();
	}
}

char QLayoutItem_HasHeightForWidthDefault(void* ptr)
{
	if (dynamic_cast<QSpacerItem*>(static_cast<QLayoutItem*>(ptr))) {
		return static_cast<QSpacerItem*>(ptr)->QSpacerItem::hasHeightForWidth();
	} else if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::hasHeightForWidth();
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::hasHeightForWidth();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::hasHeightForWidth();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::hasHeightForWidth();
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->QLayout::hasHeightForWidth();
	} else {
		return static_cast<QLayoutItem*>(ptr)->QLayoutItem::hasHeightForWidth();
	}
}

int QLayoutItem_HeightForWidth(void* ptr, int vin)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->heightForWidth(vin);
	} else {
		return static_cast<QLayoutItem*>(ptr)->heightForWidth(vin);
	}
}

int QLayoutItem_HeightForWidthDefault(void* ptr, int vin)
{
	if (dynamic_cast<QSpacerItem*>(static_cast<QLayoutItem*>(ptr))) {
		return static_cast<QSpacerItem*>(ptr)->QSpacerItem::heightForWidth(vin);
	} else if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::heightForWidth(vin);
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::heightForWidth(vin);
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::heightForWidth(vin);
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::heightForWidth(vin);
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->QLayout::heightForWidth(vin);
	} else {
		return static_cast<QLayoutItem*>(ptr)->QLayoutItem::heightForWidth(vin);
	}
}

void QLayoutItem_Invalidate(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QLayout*>(ptr)->invalidate();
	} else {
		static_cast<QLayoutItem*>(ptr)->invalidate();
	}
}

void QLayoutItem_InvalidateDefault(void* ptr)
{
	if (dynamic_cast<QSpacerItem*>(static_cast<QLayoutItem*>(ptr))) {
		static_cast<QSpacerItem*>(ptr)->QSpacerItem::invalidate();
	} else if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QGridLayout*>(ptr)->QGridLayout::invalidate();
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::invalidate();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::invalidate();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxLayout*>(ptr)->QBoxLayout::invalidate();
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QLayout*>(ptr)->QLayout::invalidate();
	} else {
		static_cast<QLayoutItem*>(ptr)->QLayoutItem::invalidate();
	}
}

char QLayoutItem_IsEmpty(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->isEmpty();
	} else {
		return static_cast<QLayoutItem*>(ptr)->isEmpty();
	}
}

void* QLayoutItem_Layout(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->layout();
	} else {
		return static_cast<QLayoutItem*>(ptr)->layout();
	}
}

void* QLayoutItem_LayoutDefault(void* ptr)
{
	if (dynamic_cast<QSpacerItem*>(static_cast<QLayoutItem*>(ptr))) {
		return static_cast<QSpacerItem*>(ptr)->QSpacerItem::layout();
	} else if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::layout();
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::layout();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::layout();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::layout();
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->QLayout::layout();
	} else {
		return static_cast<QLayoutItem*>(ptr)->QLayoutItem::layout();
	}
}

void* QLayoutItem_MaximumSize(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QLayoutItem*>(ptr)->maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

int QLayoutItem_MinimumHeightForWidth(void* ptr, int w)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->minimumHeightForWidth(w);
	} else {
		return static_cast<QLayoutItem*>(ptr)->minimumHeightForWidth(w);
	}
}

int QLayoutItem_MinimumHeightForWidthDefault(void* ptr, int w)
{
	if (dynamic_cast<QSpacerItem*>(static_cast<QLayoutItem*>(ptr))) {
		return static_cast<QSpacerItem*>(ptr)->QSpacerItem::minimumHeightForWidth(w);
	} else if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::minimumHeightForWidth(w);
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::minimumHeightForWidth(w);
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::minimumHeightForWidth(w);
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::minimumHeightForWidth(w);
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->QLayout::minimumHeightForWidth(w);
	} else {
		return static_cast<QLayoutItem*>(ptr)->QLayoutItem::minimumHeightForWidth(w);
	}
}

void* QLayoutItem_MinimumSize(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QLayoutItem*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

void QLayoutItem_SetAlignment(void* ptr, long long alignment)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QLayout*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
	} else {
		static_cast<QLayoutItem*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
	}
}

void QLayoutItem_SetGeometry(void* ptr, void* r)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		static_cast<QLayout*>(ptr)->setGeometry(*static_cast<QRect*>(r));
	} else {
		static_cast<QLayoutItem*>(ptr)->setGeometry(*static_cast<QRect*>(r));
	}
}

void* QLayoutItem_SizeHint(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QLayout*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QLayoutItem*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

void* QLayoutItem_SpacerItem(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->spacerItem();
	} else {
		return static_cast<QLayoutItem*>(ptr)->spacerItem();
	}
}

void* QLayoutItem_SpacerItemDefault(void* ptr)
{
	if (dynamic_cast<QSpacerItem*>(static_cast<QLayoutItem*>(ptr))) {
		return static_cast<QSpacerItem*>(ptr)->QSpacerItem::spacerItem();
	} else if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::spacerItem();
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::spacerItem();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::spacerItem();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::spacerItem();
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->QLayout::spacerItem();
	} else {
		return static_cast<QLayoutItem*>(ptr)->QLayoutItem::spacerItem();
	}
}

void* QLayoutItem_Widget(void* ptr)
{
	if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->widget();
	} else {
		return static_cast<QLayoutItem*>(ptr)->widget();
	}
}

void* QLayoutItem_WidgetDefault(void* ptr)
{
	if (dynamic_cast<QSpacerItem*>(static_cast<QLayoutItem*>(ptr))) {
		return static_cast<QSpacerItem*>(ptr)->QSpacerItem::widget();
	} else if (dynamic_cast<QGridLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGridLayout*>(ptr)->QGridLayout::widget();
	} else if (dynamic_cast<QVBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVBoxLayout*>(ptr)->QVBoxLayout::widget();
	} else if (dynamic_cast<QHBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHBoxLayout*>(ptr)->QHBoxLayout::widget();
	} else if (dynamic_cast<QBoxLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxLayout*>(ptr)->QBoxLayout::widget();
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLayout*>(ptr)->QLayout::widget();
	} else {
		return static_cast<QLayoutItem*>(ptr)->QLayoutItem::widget();
	}
}

void QLayoutItem_DestroyQLayoutItem(void* ptr)
{
	static_cast<QLayoutItem*>(ptr)->~QLayoutItem();
}

void QLayoutItem_DestroyQLayoutItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQLineEdit: public QLineEdit
{
public:
	MyQLineEdit(QWidget *parent = Q_NULLPTR) : QLineEdit(parent) {QLineEdit_QLineEdit_QRegisterMetaType();};
	MyQLineEdit(const QString &contents, QWidget *parent = Q_NULLPTR) : QLineEdit(contents, parent) {QLineEdit_QLineEdit_QRegisterMetaType();};
	void clear() { callbackQLineEdit_Clear(this); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	void copy() const { callbackQLineEdit_Copy(const_cast<void*>(static_cast<const void*>(this))); };
	void Signal_EditingFinished() { callbackQLineEdit_EditingFinished(this); };
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	void mouseDoubleClickEvent(QMouseEvent * e) { callbackQWidget_MouseDoubleClickEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	void paste() { callbackQLineEdit_Paste(this); };
	void Signal_SelectionChanged() { callbackQLineEdit_SelectionChanged(this); };
	void setText(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQLineEdit_SetText(this, vqsPacked); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_TextChanged(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtWidgets_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQLineEdit_TextChanged(this, textPacked); };
	 ~MyQLineEdit() { callbackQLineEdit_DestroyQLineEdit(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QLineEdit*)
Q_DECLARE_METATYPE(MyQLineEdit*)

int QLineEdit_QLineEdit_QRegisterMetaType(){qRegisterMetaType<QLineEdit*>(); return qRegisterMetaType<MyQLineEdit*>();}

void* QLineEdit_NewQLineEdit(void* parent)
{
		return new MyQLineEdit(static_cast<QWidget*>(parent));
}

void* QLineEdit_NewQLineEdit2(struct QtWidgets_PackedString contents, void* parent)
{
		return new MyQLineEdit(QString::fromUtf8(contents.data, contents.len), static_cast<QWidget*>(parent));
}

void QLineEdit_AddAction(void* ptr, void* action, long long position)
{
	static_cast<QLineEdit*>(ptr)->addAction(static_cast<QAction*>(action), static_cast<QLineEdit::ActionPosition>(position));
}

void* QLineEdit_AddAction2(void* ptr, void* icon, long long position)
{
	return static_cast<QLineEdit*>(ptr)->addAction(*static_cast<QIcon*>(icon), static_cast<QLineEdit::ActionPosition>(position));
}

long long QLineEdit_Alignment(void* ptr)
{
	return static_cast<QLineEdit*>(ptr)->alignment();
}

void QLineEdit_Clear(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "clear");
}

void QLineEdit_ClearDefault(void* ptr)
{
		static_cast<QLineEdit*>(ptr)->QLineEdit::clear();
}

void QLineEdit_Copy(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "copy");
}

void QLineEdit_CopyDefault(void* ptr)
{
		static_cast<QLineEdit*>(ptr)->QLineEdit::copy();
}

void* QLineEdit_CreateStandardContextMenu(void* ptr)
{
	return static_cast<QLineEdit*>(ptr)->createStandardContextMenu();
}

int QLineEdit_CursorPosition(void* ptr)
{
	return static_cast<QLineEdit*>(ptr)->cursorPosition();
}

void QLineEdit_Del(void* ptr)
{
	static_cast<QLineEdit*>(ptr)->del();
}

long long QLineEdit_EchoMode(void* ptr)
{
	return static_cast<QLineEdit*>(ptr)->echoMode();
}

void QLineEdit_ConnectEditingFinished(void* ptr, long long t)
{
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::editingFinished), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_EditingFinished), static_cast<Qt::ConnectionType>(t));
}

void QLineEdit_DisconnectEditingFinished(void* ptr)
{
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::editingFinished), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_EditingFinished));
}

void QLineEdit_EditingFinished(void* ptr)
{
	static_cast<QLineEdit*>(ptr)->editingFinished();
}

void QLineEdit_End(void* ptr, char mark)
{
	static_cast<QLineEdit*>(ptr)->end(mark != 0);
}

void QLineEdit_Home(void* ptr, char mark)
{
	static_cast<QLineEdit*>(ptr)->home(mark != 0);
}

void QLineEdit_Insert(void* ptr, struct QtWidgets_PackedString newText)
{
	static_cast<QLineEdit*>(ptr)->insert(QString::fromUtf8(newText.data, newText.len));
}

char QLineEdit_IsModified(void* ptr)
{
	return static_cast<QLineEdit*>(ptr)->isModified();
}

void QLineEdit_Paste(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "paste");
}

void QLineEdit_PasteDefault(void* ptr)
{
		static_cast<QLineEdit*>(ptr)->QLineEdit::paste();
}

struct QtWidgets_PackedString QLineEdit_PlaceholderText(void* ptr)
{
	return ({ QByteArray* t749e76 = new QByteArray(static_cast<QLineEdit*>(ptr)->placeholderText().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t749e76->prepend("WHITESPACE").constData()+10), t749e76->size()-10, t749e76 }; });
}

struct QtWidgets_PackedString QLineEdit_SelectedText(void* ptr)
{
	return ({ QByteArray* t1e6e00 = new QByteArray(static_cast<QLineEdit*>(ptr)->selectedText().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t1e6e00->prepend("WHITESPACE").constData()+10), t1e6e00->size()-10, t1e6e00 }; });
}

void QLineEdit_ConnectSelectionChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::selectionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_SelectionChanged), static_cast<Qt::ConnectionType>(t));
}

void QLineEdit_DisconnectSelectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::selectionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_SelectionChanged));
}

void QLineEdit_SelectionChanged(void* ptr)
{
	static_cast<QLineEdit*>(ptr)->selectionChanged();
}

int QLineEdit_SelectionEnd(void* ptr)
{
	return static_cast<QLineEdit*>(ptr)->selectionEnd();
}

int QLineEdit_SelectionStart(void* ptr)
{
	return static_cast<QLineEdit*>(ptr)->selectionStart();
}

void QLineEdit_SetAlignment(void* ptr, long long flag)
{
	static_cast<QLineEdit*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(flag));
}

void QLineEdit_SetEchoMode(void* ptr, long long vql)
{
	static_cast<QLineEdit*>(ptr)->setEchoMode(static_cast<QLineEdit::EchoMode>(vql));
}

void QLineEdit_SetModified(void* ptr, char vbo)
{
	static_cast<QLineEdit*>(ptr)->setModified(vbo != 0);
}

void QLineEdit_SetPlaceholderText(void* ptr, struct QtWidgets_PackedString vqs)
{
	static_cast<QLineEdit*>(ptr)->setPlaceholderText(QString::fromUtf8(vqs.data, vqs.len));
}

void QLineEdit_SetReadOnly(void* ptr, char vbo)
{
	static_cast<QLineEdit*>(ptr)->setReadOnly(vbo != 0);
}

void QLineEdit_SetSelection(void* ptr, int start, int length)
{
	static_cast<QLineEdit*>(ptr)->setSelection(start, length);
}

void QLineEdit_SetText(void* ptr, struct QtWidgets_PackedString vqs)
{
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "setText", Q_ARG(const QString, QString::fromUtf8(vqs.data, vqs.len)));
}

void QLineEdit_SetTextDefault(void* ptr, struct QtWidgets_PackedString vqs)
{
		static_cast<QLineEdit*>(ptr)->QLineEdit::setText(QString::fromUtf8(vqs.data, vqs.len));
}

void QLineEdit_SetValidator(void* ptr, void* v)
{
	static_cast<QLineEdit*>(ptr)->setValidator(static_cast<QValidator*>(v));
}

struct QtWidgets_PackedString QLineEdit_Text(void* ptr)
{
	return ({ QByteArray* te2605d = new QByteArray(static_cast<QLineEdit*>(ptr)->text().toUtf8()); QtWidgets_PackedString { const_cast<char*>(te2605d->prepend("WHITESPACE").constData()+10), te2605d->size()-10, te2605d }; });
}

void QLineEdit_ConnectTextChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextChanged), static_cast<Qt::ConnectionType>(t));
}

void QLineEdit_DisconnectTextChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextChanged));
}

void QLineEdit_TextChanged(void* ptr, struct QtWidgets_PackedString text)
{
	static_cast<QLineEdit*>(ptr)->textChanged(QString::fromUtf8(text.data, text.len));
}

void* QLineEdit_Validator(void* ptr)
{
	return const_cast<QValidator*>(static_cast<QLineEdit*>(ptr)->validator());
}

void QLineEdit_DestroyQLineEdit(void* ptr)
{
	static_cast<QLineEdit*>(ptr)->~QLineEdit();
}

void QLineEdit_DestroyQLineEditDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQMainWindow: public QMainWindow
{
public:
	MyQMainWindow(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QMainWindow(parent, flags) {QMainWindow_QMainWindow_QRegisterMetaType();};
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	 ~MyQMainWindow() { callbackQMainWindow_DestroyQMainWindow(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QMainWindow*)
Q_DECLARE_METATYPE(MyQMainWindow*)

int QMainWindow_QMainWindow_QRegisterMetaType(){qRegisterMetaType<QMainWindow*>(); return qRegisterMetaType<MyQMainWindow*>();}

void* QMainWindow_NewQMainWindow(void* parent, long long flags)
{
		return new MyQMainWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QMainWindow_AddToolBar(void* ptr, long long area, void* toolbar)
{
	static_cast<QMainWindow*>(ptr)->addToolBar(static_cast<Qt::ToolBarArea>(area), static_cast<QToolBar*>(toolbar));
}

void QMainWindow_AddToolBar2(void* ptr, void* toolbar)
{
	static_cast<QMainWindow*>(ptr)->addToolBar(static_cast<QToolBar*>(toolbar));
}

void* QMainWindow_AddToolBar3(void* ptr, struct QtWidgets_PackedString title)
{
	return static_cast<QMainWindow*>(ptr)->addToolBar(QString::fromUtf8(title.data, title.len));
}

void* QMainWindow_CentralWidget(void* ptr)
{
	return static_cast<QMainWindow*>(ptr)->centralWidget();
}

void* QMainWindow_MenuBar(void* ptr)
{
	return static_cast<QMainWindow*>(ptr)->menuBar();
}

void QMainWindow_SetCentralWidget(void* ptr, void* widget)
{
		static_cast<QMainWindow*>(ptr)->setCentralWidget(static_cast<QWidget*>(widget));
}

void* QMainWindow_StatusBar(void* ptr)
{
	return static_cast<QMainWindow*>(ptr)->statusBar();
}

void QMainWindow_DestroyQMainWindow(void* ptr)
{
	static_cast<QMainWindow*>(ptr)->~QMainWindow();
}

void QMainWindow_DestroyQMainWindowDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QMainWindow___resizeDocks_docks_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QMainWindow___resizeDocks_docks_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* QMainWindow___resizeDocks_docks_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

int QMainWindow___resizeDocks_sizes_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QMainWindow___resizeDocks_sizes_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QMainWindow___resizeDocks_sizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QMainWindow___tabifiedDockWidgets_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QMainWindow___tabifiedDockWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* QMainWindow___tabifiedDockWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

class MyQMenu: public QMenu
{
public:
	MyQMenu(QWidget *parent = Q_NULLPTR) : QMenu(parent) {QMenu_QMenu_QRegisterMetaType();};
	MyQMenu(const QString &title, QWidget *parent = Q_NULLPTR) : QMenu(title, parent) {QMenu_QMenu_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	void hideEvent(QHideEvent * vqh) { callbackQWidget_HideEvent(this, vqh); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void timerEvent(QTimerEvent * e) { callbackQWidget_TimerEvent(this, e); };
	void Signal_Triggered(QAction * action) { callbackQMenu_Triggered(this, action); };
	 ~MyQMenu() { callbackQMenu_DestroyQMenu(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QMenu*)
Q_DECLARE_METATYPE(MyQMenu*)

int QMenu_QMenu_QRegisterMetaType(){qRegisterMetaType<QMenu*>(); return qRegisterMetaType<MyQMenu*>();}

void* QMenu_NewQMenu(void* parent)
{
		return new MyQMenu(static_cast<QWidget*>(parent));
}

void* QMenu_NewQMenu2(struct QtWidgets_PackedString title, void* parent)
{
		return new MyQMenu(QString::fromUtf8(title.data, title.len), static_cast<QWidget*>(parent));
}

void* QMenu_AddAction(void* ptr, struct QtWidgets_PackedString text)
{
	return static_cast<QMenu*>(ptr)->addAction(QString::fromUtf8(text.data, text.len));
}

void* QMenu_AddAction2(void* ptr, void* icon, struct QtWidgets_PackedString text)
{
	return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len));
}

void* QMenu_AddAction3(void* ptr, struct QtWidgets_PackedString text, void* receiver, char* member, void* shortcut)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenu*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsObject*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenu*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsWidget*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenu*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QLayout*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenu*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
	} else {
		return static_cast<QMenu*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QObject*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
	}
}

void* QMenu_AddAction4(void* ptr, void* icon, struct QtWidgets_PackedString text, void* receiver, char* member, void* shortcut)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QGraphicsObject*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QGraphicsWidget*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QLayout*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
	} else {
		return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QObject*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
	}
}

void* QMenu_AddMenu(void* ptr, void* menu)
{
	return static_cast<QMenu*>(ptr)->addMenu(static_cast<QMenu*>(menu));
}

void* QMenu_AddMenu2(void* ptr, struct QtWidgets_PackedString title)
{
	return static_cast<QMenu*>(ptr)->addMenu(QString::fromUtf8(title.data, title.len));
}

void* QMenu_AddMenu3(void* ptr, void* icon, struct QtWidgets_PackedString title)
{
	return static_cast<QMenu*>(ptr)->addMenu(*static_cast<QIcon*>(icon), QString::fromUtf8(title.data, title.len));
}

void* QMenu_AddSeparator(void* ptr)
{
	return static_cast<QMenu*>(ptr)->addSeparator();
}

void QMenu_Clear(void* ptr)
{
	static_cast<QMenu*>(ptr)->clear();
}

int QMenu_ColumnCount(void* ptr)
{
	return static_cast<QMenu*>(ptr)->columnCount();
}

void* QMenu_Exec(void* ptr)
{
	return static_cast<QMenu*>(ptr)->exec();
}

void* QMenu_Exec2(void* ptr, void* p, void* action)
{
	return static_cast<QMenu*>(ptr)->exec(*static_cast<QPoint*>(p), static_cast<QAction*>(action));
}

void* QMenu_QMenu_Exec3(void* actions, void* pos, void* at, void* parent)
{
		return QMenu::exec(({ QList<QAction *>* tmpP = static_cast<QList<QAction *>*>(actions); QList<QAction *> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }), *static_cast<QPoint*>(pos), static_cast<QAction*>(at), static_cast<QWidget*>(parent));
}

void* QMenu_Icon(void* ptr)
{
	return new QIcon(static_cast<QMenu*>(ptr)->icon());
}

char QMenu_IsEmpty(void* ptr)
{
	return static_cast<QMenu*>(ptr)->isEmpty();
}

void QMenu_Popup(void* ptr, void* p, void* atAction)
{
	static_cast<QMenu*>(ptr)->popup(*static_cast<QPoint*>(p), static_cast<QAction*>(atAction));
}

void QMenu_SetIcon(void* ptr, void* icon)
{
	static_cast<QMenu*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

struct QtWidgets_PackedString QMenu_Title(void* ptr)
{
	return ({ QByteArray* tbbbdf8 = new QByteArray(static_cast<QMenu*>(ptr)->title().toUtf8()); QtWidgets_PackedString { const_cast<char*>(tbbbdf8->prepend("WHITESPACE").constData()+10), tbbbdf8->size()-10, tbbbdf8 }; });
}

void QMenu_ConnectTriggered(void* ptr, long long t)
{
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::triggered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Triggered), static_cast<Qt::ConnectionType>(t));
}

void QMenu_DisconnectTriggered(void* ptr)
{
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::triggered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Triggered));
}

void QMenu_Triggered(void* ptr, void* action)
{
	static_cast<QMenu*>(ptr)->triggered(static_cast<QAction*>(action));
}

void QMenu_DestroyQMenu(void* ptr)
{
	static_cast<QMenu*>(ptr)->~QMenu();
}

void QMenu_DestroyQMenuDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QMenu___exec_actions_atList3(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QMenu___exec_actions_setList3(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QMenu___exec_actions_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

class MyQMenuBar: public QMenuBar
{
public:
	MyQMenuBar(QWidget *parent = Q_NULLPTR) : QMenuBar(parent) {QMenuBar_QMenuBar_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	bool eventFilter(QObject * object, QEvent * event) { return callbackQWidget_EventFilter(this, object, event) != 0; };
	int heightForWidth(int vin) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), vin); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	void resizeEvent(QResizeEvent * vqr) { callbackQWidget_ResizeEvent(this, vqr); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void timerEvent(QTimerEvent * e) { callbackQWidget_TimerEvent(this, e); };
	void Signal_Triggered(QAction * action) { callbackQMenuBar_Triggered(this, action); };
	 ~MyQMenuBar() { callbackQMenuBar_DestroyQMenuBar(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QMenuBar*)
Q_DECLARE_METATYPE(MyQMenuBar*)

int QMenuBar_QMenuBar_QRegisterMetaType(){qRegisterMetaType<QMenuBar*>(); return qRegisterMetaType<MyQMenuBar*>();}

void* QMenuBar_NewQMenuBar(void* parent)
{
		return new MyQMenuBar(static_cast<QWidget*>(parent));
}

void* QMenuBar_AddAction(void* ptr, struct QtWidgets_PackedString text)
{
	return static_cast<QMenuBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len));
}

void* QMenuBar_AddAction2(void* ptr, struct QtWidgets_PackedString text, void* receiver, char* member)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenuBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsObject*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenuBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsWidget*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenuBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QLayout*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(receiver))) {
		return static_cast<QMenuBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(receiver), const_cast<const char*>(member));
	} else {
		return static_cast<QMenuBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QObject*>(receiver), const_cast<const char*>(member));
	}
}

void* QMenuBar_AddMenu(void* ptr, void* menu)
{
	return static_cast<QMenuBar*>(ptr)->addMenu(static_cast<QMenu*>(menu));
}

void* QMenuBar_AddMenu2(void* ptr, struct QtWidgets_PackedString title)
{
	return static_cast<QMenuBar*>(ptr)->addMenu(QString::fromUtf8(title.data, title.len));
}

void* QMenuBar_AddMenu3(void* ptr, void* icon, struct QtWidgets_PackedString title)
{
	return static_cast<QMenuBar*>(ptr)->addMenu(*static_cast<QIcon*>(icon), QString::fromUtf8(title.data, title.len));
}

void* QMenuBar_AddSeparator(void* ptr)
{
	return static_cast<QMenuBar*>(ptr)->addSeparator();
}

void QMenuBar_Clear(void* ptr)
{
	static_cast<QMenuBar*>(ptr)->clear();
}

void QMenuBar_ConnectTriggered(void* ptr, long long t)
{
	QObject::connect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::triggered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Triggered), static_cast<Qt::ConnectionType>(t));
}

void QMenuBar_DisconnectTriggered(void* ptr)
{
	QObject::disconnect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::triggered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Triggered));
}

void QMenuBar_Triggered(void* ptr, void* action)
{
	static_cast<QMenuBar*>(ptr)->triggered(static_cast<QAction*>(action));
}

void QMenuBar_DestroyQMenuBar(void* ptr)
{
	static_cast<QMenuBar*>(ptr)->~QMenuBar();
}

void QMenuBar_DestroyQMenuBarDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQMessageBox: public QMessageBox
{
public:
	MyQMessageBox(QWidget *parent = Q_NULLPTR) : QMessageBox(parent) {QMessageBox_QMessageBox_QRegisterMetaType();};
	MyQMessageBox(QMessageBox::Icon icon, const QString &title, const QString &text, QMessageBox::StandardButtons buttons = NoButton, QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::Dialog | Qt::MSWindowsFixedSizeDialogHint) : QMessageBox(icon, title, text, buttons, parent, ff) {QMessageBox_QMessageBox_QRegisterMetaType();};
	void closeEvent(QCloseEvent * e) { callbackQWidget_CloseEvent(this, e); };
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	int exec() { return callbackQDialog_Exec(this); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void showEvent(QShowEvent * e) { callbackQWidget_ShowEvent(this, e); };
	 ~MyQMessageBox() { callbackQMessageBox_DestroyQMessageBox(this); };
	void accept() { callbackQDialog_Accept(this); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	void done(int r) { callbackQDialog_Done(this, r); };
	bool eventFilter(QObject * o, QEvent * e) { return callbackQWidget_EventFilter(this, o, e) != 0; };
	void Signal_Finished(int result) { callbackQDialog_Finished(this, result); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void show() { callbackQWidget_Show(this); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QMessageBox*)
Q_DECLARE_METATYPE(MyQMessageBox*)

int QMessageBox_QMessageBox_QRegisterMetaType(){qRegisterMetaType<QMessageBox*>(); return qRegisterMetaType<MyQMessageBox*>();}

int QMessageBox_ButtonMask_Type()
{
	return QMessageBox::ButtonMask;
}

void* QMessageBox_NewQMessageBox(void* parent)
{
		return new MyQMessageBox(static_cast<QWidget*>(parent));
}

void* QMessageBox_NewQMessageBox2(long long icon, struct QtWidgets_PackedString title, struct QtWidgets_PackedString text, long long buttons, void* parent, long long ff)
{
		return new MyQMessageBox(static_cast<QMessageBox::Icon>(icon), QString::fromUtf8(title.data, title.len), QString::fromUtf8(text.data, text.len), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

void QMessageBox_QMessageBox_About(void* parent, struct QtWidgets_PackedString title, struct QtWidgets_PackedString text)
{
		QMessageBox::about(static_cast<QWidget*>(parent), QString::fromUtf8(title.data, title.len), QString::fromUtf8(text.data, text.len));
}

void* QMessageBox_Button(void* ptr, long long which)
{
	return static_cast<QMessageBox*>(ptr)->button(static_cast<QMessageBox::StandardButton>(which));
}

struct QtWidgets_PackedList QMessageBox_Buttons(void* ptr)
{
	return ({ QList<QAbstractButton *>* tmpValuef25e9a = new QList<QAbstractButton *>(static_cast<QMessageBox*>(ptr)->buttons()); QtWidgets_PackedList { tmpValuef25e9a, tmpValuef25e9a->size() }; });
}

long long QMessageBox_Icon(void* ptr)
{
	return static_cast<QMessageBox*>(ptr)->icon();
}

long long QMessageBox_QMessageBox_Information(void* parent, struct QtWidgets_PackedString title, struct QtWidgets_PackedString text, long long buttons, long long defaultButton)
{
		return QMessageBox::information(static_cast<QWidget*>(parent), QString::fromUtf8(title.data, title.len), QString::fromUtf8(text.data, text.len), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_Open(void* ptr, void* receiver, char* member)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(receiver))) {
		static_cast<QMessageBox*>(ptr)->open(static_cast<QGraphicsObject*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(receiver))) {
		static_cast<QMessageBox*>(ptr)->open(static_cast<QGraphicsWidget*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(receiver))) {
		static_cast<QMessageBox*>(ptr)->open(static_cast<QLayout*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(receiver))) {
		static_cast<QMessageBox*>(ptr)->open(static_cast<QWidget*>(receiver), const_cast<const char*>(member));
	} else {
		static_cast<QMessageBox*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
	}
}

long long QMessageBox_QMessageBox_Question(void* parent, struct QtWidgets_PackedString title, struct QtWidgets_PackedString text, long long buttons, long long defaultButton)
{
		return QMessageBox::question(static_cast<QWidget*>(parent), QString::fromUtf8(title.data, title.len), QString::fromUtf8(text.data, text.len), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_SetIcon(void* ptr, long long vqm)
{
	static_cast<QMessageBox*>(ptr)->setIcon(static_cast<QMessageBox::Icon>(vqm));
}

void QMessageBox_SetStandardButtons(void* ptr, long long buttons)
{
	static_cast<QMessageBox*>(ptr)->setStandardButtons(static_cast<QMessageBox::StandardButton>(buttons));
}

void QMessageBox_SetText(void* ptr, struct QtWidgets_PackedString text)
{
	static_cast<QMessageBox*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

void QMessageBox_SetWindowTitle(void* ptr, struct QtWidgets_PackedString title)
{
	static_cast<QMessageBox*>(ptr)->setWindowTitle(QString::fromUtf8(title.data, title.len));
}

long long QMessageBox_StandardButton(void* ptr, void* button)
{
	return static_cast<QMessageBox*>(ptr)->standardButton(static_cast<QAbstractButton*>(button));
}

long long QMessageBox_StandardButtons(void* ptr)
{
	return static_cast<QMessageBox*>(ptr)->standardButtons();
}

struct QtWidgets_PackedString QMessageBox_Text(void* ptr)
{
	return ({ QByteArray* te2bb59 = new QByteArray(static_cast<QMessageBox*>(ptr)->text().toUtf8()); QtWidgets_PackedString { const_cast<char*>(te2bb59->prepend("WHITESPACE").constData()+10), te2bb59->size()-10, te2bb59 }; });
}

long long QMessageBox_QMessageBox_Warning(void* parent, struct QtWidgets_PackedString title, struct QtWidgets_PackedString text, long long buttons, long long defaultButton)
{
		return QMessageBox::warning(static_cast<QWidget*>(parent), QString::fromUtf8(title.data, title.len), QString::fromUtf8(text.data, text.len), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_DestroyQMessageBox(void* ptr)
{
	static_cast<QMessageBox*>(ptr)->~QMessageBox();
}

void QMessageBox_DestroyQMessageBoxDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QMessageBox___buttons_atList(void* ptr, int i)
{
	return ({QAbstractButton * tmp = static_cast<QList<QAbstractButton *>*>(ptr)->at(i); if (i == static_cast<QList<QAbstractButton *>*>(ptr)->size()-1) { static_cast<QList<QAbstractButton *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QMessageBox___buttons_setList(void* ptr, void* i)
{
	static_cast<QList<QAbstractButton *>*>(ptr)->append(static_cast<QAbstractButton*>(i));
}

void* QMessageBox___buttons_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAbstractButton *>();
}

class MyQPushButton: public QPushButton
{
public:
	MyQPushButton(QWidget *parent = Q_NULLPTR) : QPushButton(parent) {QPushButton_QPushButton_QRegisterMetaType();};
	MyQPushButton(const QString &text, QWidget *parent = Q_NULLPTR) : QPushButton(text, parent) {QPushButton_QPushButton_QRegisterMetaType();};
	MyQPushButton(const QIcon &icon, const QString &text, QWidget *parent = Q_NULLPTR) : QPushButton(icon, text, parent) {QPushButton_QPushButton_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQPushButton() { callbackQPushButton_DestroyQPushButton(this); };
	void click() { callbackQAbstractButton_Click(this); };
	void Signal_Clicked(bool checked) { callbackQAbstractButton_Clicked(this, checked); };
	void keyReleaseEvent(QKeyEvent * e) { callbackQWidget_KeyReleaseEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	void paintEvent(QPaintEvent * e) { callbackQPushButton_PaintEvent(this, e); };
	void setChecked(bool vbo) { callbackQAbstractButton_SetChecked(this, vbo); };
	void timerEvent(QTimerEvent * e) { callbackQWidget_TimerEvent(this, e); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QPushButton*)
Q_DECLARE_METATYPE(MyQPushButton*)

int QPushButton_QPushButton_QRegisterMetaType(){qRegisterMetaType<QPushButton*>(); return qRegisterMetaType<MyQPushButton*>();}

void* QPushButton_NewQPushButton(void* parent)
{
		return new MyQPushButton(static_cast<QWidget*>(parent));
}

void* QPushButton_NewQPushButton2(struct QtWidgets_PackedString text, void* parent)
{
		return new MyQPushButton(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
}

void* QPushButton_NewQPushButton3(void* icon, struct QtWidgets_PackedString text, void* parent)
{
		return new MyQPushButton(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
}

void* QPushButton_Menu(void* ptr)
{
	return static_cast<QPushButton*>(ptr)->menu();
}

void QPushButton_DestroyQPushButton(void* ptr)
{
	static_cast<QPushButton*>(ptr)->~QPushButton();
}

void QPushButton_DestroyQPushButtonDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QPushButton_PaintEvent(void* ptr, void* e)
{
	static_cast<QPushButton*>(ptr)->paintEvent(static_cast<QPaintEvent*>(e));
}

void QPushButton_PaintEventDefault(void* ptr, void* e)
{
		static_cast<QPushButton*>(ptr)->QPushButton::paintEvent(static_cast<QPaintEvent*>(e));
}

class MyQScrollArea: public QScrollArea
{
public:
	MyQScrollArea(QWidget *parent = Q_NULLPTR) : QScrollArea(parent) {QScrollArea_QScrollArea_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	bool eventFilter(QObject * o, QEvent * e) { return callbackQWidget_EventFilter(this, o, e) != 0; };
	void resizeEvent(QResizeEvent * vqr) { callbackQWidget_ResizeEvent(this, vqr); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQScrollArea() { callbackQScrollArea_DestroyQScrollArea(this); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void mouseDoubleClickEvent(QMouseEvent * e) { callbackQWidget_MouseDoubleClickEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QScrollArea*)
Q_DECLARE_METATYPE(MyQScrollArea*)

int QScrollArea_QScrollArea_QRegisterMetaType(){qRegisterMetaType<QScrollArea*>(); return qRegisterMetaType<MyQScrollArea*>();}

void* QScrollArea_NewQScrollArea(void* parent)
{
		return new MyQScrollArea(static_cast<QWidget*>(parent));
}

long long QScrollArea_Alignment(void* ptr)
{
	return static_cast<QScrollArea*>(ptr)->alignment();
}

void QScrollArea_SetAlignment(void* ptr, long long vqt)
{
	static_cast<QScrollArea*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(vqt));
}

void QScrollArea_SetWidget(void* ptr, void* widget)
{
		static_cast<QScrollArea*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void* QScrollArea_Widget(void* ptr)
{
	return static_cast<QScrollArea*>(ptr)->widget();
}

void QScrollArea_DestroyQScrollArea(void* ptr)
{
	static_cast<QScrollArea*>(ptr)->~QScrollArea();
}

void QScrollArea_DestroyQScrollAreaDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQScrollBar: public QScrollBar
{
public:
	MyQScrollBar(QWidget *parent = Q_NULLPTR) : QScrollBar(parent) {QScrollBar_QScrollBar_QRegisterMetaType();};
	MyQScrollBar(Qt::Orientation orientation, QWidget *parent = Q_NULLPTR) : QScrollBar(orientation, parent) {QScrollBar_QScrollBar_QRegisterMetaType();};
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	void hideEvent(QHideEvent * vqh) { callbackQWidget_HideEvent(this, vqh); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQScrollBar() { callbackQScrollBar_DestroyQScrollBar(this); };
	void keyPressEvent(QKeyEvent * ev) { callbackQWidget_KeyPressEvent(this, ev); };
	void setOrientation(Qt::Orientation vqt) { callbackQAbstractSlider_SetOrientation(this, vqt); };
	void setRange(int min, int max) { callbackQAbstractSlider_SetRange(this, min, max); };
	void timerEvent(QTimerEvent * e) { callbackQWidget_TimerEvent(this, e); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QScrollBar*)
Q_DECLARE_METATYPE(MyQScrollBar*)

int QScrollBar_QScrollBar_QRegisterMetaType(){qRegisterMetaType<QScrollBar*>(); return qRegisterMetaType<MyQScrollBar*>();}

void* QScrollBar_NewQScrollBar(void* parent)
{
		return new MyQScrollBar(static_cast<QWidget*>(parent));
}

void* QScrollBar_NewQScrollBar2(long long orientation, void* parent)
{
		return new MyQScrollBar(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

void QScrollBar_DestroyQScrollBar(void* ptr)
{
	static_cast<QScrollBar*>(ptr)->~QScrollBar();
}

void QScrollBar_DestroyQScrollBarDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

Q_DECLARE_METATYPE(QSizePolicy)
Q_DECLARE_METATYPE(QSizePolicy*)
void* QSizePolicy_NewQSizePolicy()
{
	return new QSizePolicy();
}

void* QSizePolicy_NewQSizePolicy2(long long horizontal, long long vertical, long long ty)
{
	return new QSizePolicy(static_cast<QSizePolicy::Policy>(horizontal), static_cast<QSizePolicy::Policy>(vertical), static_cast<QSizePolicy::ControlType>(ty));
}

long long QSizePolicy_ControlType(void* ptr)
{
	return static_cast<QSizePolicy*>(ptr)->controlType();
}

long long QSizePolicy_ExpandingDirections(void* ptr)
{
	return static_cast<QSizePolicy*>(ptr)->expandingDirections();
}

char QSizePolicy_HasHeightForWidth(void* ptr)
{
	return static_cast<QSizePolicy*>(ptr)->hasHeightForWidth();
}

class MyQSpacerItem: public QSpacerItem
{
public:
	MyQSpacerItem(int w, int h, QSizePolicy::Policy hPolicy = QSizePolicy::Minimum, QSizePolicy::Policy vPolicy = QSizePolicy::Minimum) : QSpacerItem(w, h, hPolicy, vPolicy) {QSpacerItem_QSpacerItem_QRegisterMetaType();};
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQSpacerItem_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQSpacerItem_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isEmpty() const { return callbackQSpacerItem_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQSpacerItem_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQSpacerItem_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	void setGeometry(const QRect & r) { callbackQSpacerItem_SetGeometry(this, const_cast<QRect*>(&r)); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQSpacerItem_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackQLayoutItem_SpacerItem(this)); };
	 ~MyQSpacerItem() { callbackQSpacerItem_DestroyQSpacerItem(this); };
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackQLayoutItem_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQLayoutItem_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int vin) const { return callbackQLayoutItem_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), vin); };
	void invalidate() { callbackQLayoutItem_Invalidate(this); };
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	int minimumHeightForWidth(int w) const { return callbackQLayoutItem_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
};

Q_DECLARE_METATYPE(QSpacerItem*)
Q_DECLARE_METATYPE(MyQSpacerItem*)

int QSpacerItem_QSpacerItem_QRegisterMetaType(){qRegisterMetaType<QSpacerItem*>(); return qRegisterMetaType<MyQSpacerItem*>();}

void* QSpacerItem_NewQSpacerItem(int w, int h, long long hPolicy, long long vPolicy)
{
	return new MyQSpacerItem(w, h, static_cast<QSizePolicy::Policy>(hPolicy), static_cast<QSizePolicy::Policy>(vPolicy));
}

long long QSpacerItem_ExpandingDirections(void* ptr)
{
	return static_cast<QSpacerItem*>(ptr)->expandingDirections();
}

long long QSpacerItem_ExpandingDirectionsDefault(void* ptr)
{
		return static_cast<QSpacerItem*>(ptr)->QSpacerItem::expandingDirections();
}

void* QSpacerItem_Geometry(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QSpacerItem*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QSpacerItem_GeometryDefault(void* ptr)
{
		return ({ QRect tmpValue = static_cast<QSpacerItem*>(ptr)->QSpacerItem::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

char QSpacerItem_IsEmpty(void* ptr)
{
	return static_cast<QSpacerItem*>(ptr)->isEmpty();
}

char QSpacerItem_IsEmptyDefault(void* ptr)
{
		return static_cast<QSpacerItem*>(ptr)->QSpacerItem::isEmpty();
}

void* QSpacerItem_MaximumSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSpacerItem*>(ptr)->maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSpacerItem_MaximumSizeDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QSpacerItem*>(ptr)->QSpacerItem::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSpacerItem_MinimumSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSpacerItem*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSpacerItem_MinimumSizeDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QSpacerItem*>(ptr)->QSpacerItem::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QSpacerItem_SetGeometry(void* ptr, void* r)
{
	static_cast<QSpacerItem*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void QSpacerItem_SetGeometryDefault(void* ptr, void* r)
{
		static_cast<QSpacerItem*>(ptr)->QSpacerItem::setGeometry(*static_cast<QRect*>(r));
}

void* QSpacerItem_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QSpacerItem*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSpacerItem_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QSpacerItem*>(ptr)->QSpacerItem::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QSpacerItem_SizePolicy(void* ptr)
{
	return new QSizePolicy(static_cast<QSpacerItem*>(ptr)->sizePolicy());
}

void QSpacerItem_DestroyQSpacerItem(void* ptr)
{
	static_cast<QSpacerItem*>(ptr)->~QSpacerItem();
}

void QSpacerItem_DestroyQSpacerItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQSplitter: public QSplitter
{
public:
	MyQSplitter(QWidget *parent = Q_NULLPTR) : QSplitter(parent) {QSplitter_QSplitter_QRegisterMetaType();};
	MyQSplitter(Qt::Orientation orientation, QWidget *parent = Q_NULLPTR) : QSplitter(orientation, parent) {QSplitter_QSplitter_QRegisterMetaType();};
	void childEvent(QChildEvent * c) { callbackQWidget_ChildEvent(this, c); };
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	void resizeEvent(QResizeEvent * vqr) { callbackQWidget_ResizeEvent(this, vqr); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQSplitter() { callbackQSplitter_DestroyQSplitter(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QSplitter*)
Q_DECLARE_METATYPE(MyQSplitter*)

int QSplitter_QSplitter_QRegisterMetaType(){qRegisterMetaType<QSplitter*>(); return qRegisterMetaType<MyQSplitter*>();}

void* QSplitter_NewQSplitter(void* parent)
{
		return new MyQSplitter(static_cast<QWidget*>(parent));
}

void* QSplitter_NewQSplitter2(long long orientation, void* parent)
{
		return new MyQSplitter(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

void QSplitter_AddWidget(void* ptr, void* widget)
{
		static_cast<QSplitter*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

int QSplitter_Count(void* ptr)
{
	return static_cast<QSplitter*>(ptr)->count();
}

int QSplitter_IndexOf(void* ptr, void* widget)
{
		return static_cast<QSplitter*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

long long QSplitter_Orientation(void* ptr)
{
	return static_cast<QSplitter*>(ptr)->orientation();
}

void QSplitter_Refresh(void* ptr)
{
	static_cast<QSplitter*>(ptr)->refresh();
}

void QSplitter_SetOrientation(void* ptr, long long vqt)
{
	static_cast<QSplitter*>(ptr)->setOrientation(static_cast<Qt::Orientation>(vqt));
}

void QSplitter_SetStretchFactor(void* ptr, int index, int stretch)
{
	static_cast<QSplitter*>(ptr)->setStretchFactor(index, stretch);
}

void* QSplitter_Widget(void* ptr, int index)
{
	return static_cast<QSplitter*>(ptr)->widget(index);
}

void QSplitter_DestroyQSplitter(void* ptr)
{
	static_cast<QSplitter*>(ptr)->~QSplitter();
}

void QSplitter_DestroyQSplitterDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

int QSplitter___setSizes_list_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSplitter___setSizes_list_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSplitter___setSizes_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QSplitter___sizes_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QSplitter___sizes_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QSplitter___sizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

class MyQStatusBar: public QStatusBar
{
public:
	MyQStatusBar(QWidget *parent = Q_NULLPTR) : QStatusBar(parent) {QStatusBar_QStatusBar_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQWidget_Event(this, e) != 0; };
	void resizeEvent(QResizeEvent * e) { callbackQWidget_ResizeEvent(this, e); };
	void showEvent(QShowEvent * vqs) { callbackQWidget_ShowEvent(this, vqs); };
	void showMessage(const QString & message, int timeout) { QByteArray* t6f9b9a = new QByteArray(message.toUtf8()); QtWidgets_PackedString messagePacked = { const_cast<char*>(t6f9b9a->prepend("WHITESPACE").constData()+10), t6f9b9a->size()-10, t6f9b9a };callbackQStatusBar_ShowMessage(this, messagePacked, timeout); };
	 ~MyQStatusBar() { callbackQStatusBar_DestroyQStatusBar(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QStatusBar*)
Q_DECLARE_METATYPE(MyQStatusBar*)

int QStatusBar_QStatusBar_QRegisterMetaType(){qRegisterMetaType<QStatusBar*>(); return qRegisterMetaType<MyQStatusBar*>();}

void* QStatusBar_NewQStatusBar(void* parent)
{
		return new MyQStatusBar(static_cast<QWidget*>(parent));
}

void QStatusBar_AddWidget(void* ptr, void* widget, int stretch)
{
		static_cast<QStatusBar*>(ptr)->addWidget(static_cast<QWidget*>(widget), stretch);
}

void QStatusBar_ShowMessage(void* ptr, struct QtWidgets_PackedString message, int timeout)
{
	QMetaObject::invokeMethod(static_cast<QStatusBar*>(ptr), "showMessage", Q_ARG(const QString, QString::fromUtf8(message.data, message.len)), Q_ARG(int, timeout));
}

void QStatusBar_ShowMessageDefault(void* ptr, struct QtWidgets_PackedString message, int timeout)
{
		static_cast<QStatusBar*>(ptr)->QStatusBar::showMessage(QString::fromUtf8(message.data, message.len), timeout);
}

void QStatusBar_DestroyQStatusBar(void* ptr)
{
	static_cast<QStatusBar*>(ptr)->~QStatusBar();
}

void QStatusBar_DestroyQStatusBarDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQStyle: public QStyle
{
public:
	MyQStyle() : QStyle() {QStyle_QStyle_QRegisterMetaType();};
	void drawComplexControl(QStyle::ComplexControl control, const QStyleOptionComplex * option, QPainter * painter, const QWidget * widget) const { callbackQStyle_DrawComplexControl(const_cast<void*>(static_cast<const void*>(this)), control, const_cast<QStyleOptionComplex*>(option), painter, const_cast<QWidget*>(widget)); };
	void drawControl(QStyle::ControlElement element, const QStyleOption * option, QPainter * painter, const QWidget * widget) const { callbackQStyle_DrawControl(const_cast<void*>(static_cast<const void*>(this)), element, const_cast<QStyleOption*>(option), painter, const_cast<QWidget*>(widget)); };
	void drawPrimitive(QStyle::PrimitiveElement element, const QStyleOption * option, QPainter * painter, const QWidget * widget) const { callbackQStyle_DrawPrimitive(const_cast<void*>(static_cast<const void*>(this)), element, const_cast<QStyleOption*>(option), painter, const_cast<QWidget*>(widget)); };
	QPixmap generatedIconPixmap(QIcon::Mode iconMode, const QPixmap & pixmap, const QStyleOption * option) const { return *static_cast<QPixmap*>(callbackQStyle_GeneratedIconPixmap(const_cast<void*>(static_cast<const void*>(this)), iconMode, const_cast<QPixmap*>(&pixmap), const_cast<QStyleOption*>(option))); };
	QStyle::SubControl hitTestComplexControl(QStyle::ComplexControl control, const QStyleOptionComplex * option, const QPoint & position, const QWidget * widget) const { return static_cast<QStyle::SubControl>(callbackQStyle_HitTestComplexControl(const_cast<void*>(static_cast<const void*>(this)), control, const_cast<QStyleOptionComplex*>(option), const_cast<QPoint*>(&position), const_cast<QWidget*>(widget))); };
	int layoutSpacing(QSizePolicy::ControlType control1, QSizePolicy::ControlType control2, Qt::Orientation orientation, const QStyleOption * option, const QWidget * widget) const { return callbackQStyle_LayoutSpacing(const_cast<void*>(static_cast<const void*>(this)), control1, control2, orientation, const_cast<QStyleOption*>(option), const_cast<QWidget*>(widget)); };
	int pixelMetric(QStyle::PixelMetric metric, const QStyleOption * option, const QWidget * widget) const { return callbackQStyle_PixelMetric(const_cast<void*>(static_cast<const void*>(this)), metric, const_cast<QStyleOption*>(option), const_cast<QWidget*>(widget)); };
	QSize sizeFromContents(QStyle::ContentsType ty, const QStyleOption * option, const QSize & contentsSize, const QWidget * widget) const { return *static_cast<QSize*>(callbackQStyle_SizeFromContents(const_cast<void*>(static_cast<const void*>(this)), ty, const_cast<QStyleOption*>(option), const_cast<QSize*>(&contentsSize), const_cast<QWidget*>(widget))); };
	QIcon standardIcon(QStyle::StandardPixmap standardIcon, const QStyleOption * option, const QWidget * widget) const { return *static_cast<QIcon*>(callbackQStyle_StandardIcon(const_cast<void*>(static_cast<const void*>(this)), standardIcon, const_cast<QStyleOption*>(option), const_cast<QWidget*>(widget))); };
	int styleHint(QStyle::StyleHint hint, const QStyleOption * option, const QWidget * widget, QStyleHintReturn * returnData) const { return callbackQStyle_StyleHint(const_cast<void*>(static_cast<const void*>(this)), hint, const_cast<QStyleOption*>(option), const_cast<QWidget*>(widget), returnData); };
	QRect subControlRect(QStyle::ComplexControl control, const QStyleOptionComplex * option, QStyle::SubControl subControl, const QWidget * widget) const { return *static_cast<QRect*>(callbackQStyle_SubControlRect(const_cast<void*>(static_cast<const void*>(this)), control, const_cast<QStyleOptionComplex*>(option), subControl, const_cast<QWidget*>(widget))); };
	QRect subElementRect(QStyle::SubElement element, const QStyleOption * option, const QWidget * widget) const { return *static_cast<QRect*>(callbackQStyle_SubElementRect(const_cast<void*>(static_cast<const void*>(this)), element, const_cast<QStyleOption*>(option), const_cast<QWidget*>(widget))); };
	 ~MyQStyle() { callbackQStyle_DestroyQStyle(this); };
	QPixmap standardPixmap(QStyle::StandardPixmap standardIcon, const QStyleOption * option, const QWidget * widget) const { return *static_cast<QPixmap*>(callbackQStyle_StandardPixmap(const_cast<void*>(static_cast<const void*>(this)), standardIcon, const_cast<QStyleOption*>(option), const_cast<QWidget*>(widget))); };
	void childEvent(QChildEvent * event) { callbackQStyle_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQStyle_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQStyle_CustomEvent(this, event); };
	void deleteLater() { callbackQStyle_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQStyle_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQStyle_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQStyle_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQStyle_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQStyle_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQStyle_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QStyle*)
Q_DECLARE_METATYPE(MyQStyle*)

int QStyle_QStyle_QRegisterMetaType(){qRegisterMetaType<QStyle*>(); return qRegisterMetaType<MyQStyle*>();}

int QStyle_PM_ToolBarIconSize_Type()
{
	#if QT_VERSION >= 0x056000
		return QStyle::PM_ToolBarIconSize;
	#else
		return 0;
	#endif
}

int QStyle_SE_LabelLayoutItem_Type()
{
	#if QT_VERSION >= 0x057000
		return QStyle::SE_LabelLayoutItem;
	#else
		return 0;
	#endif
}

void* QStyle_NewQStyle2()
{
	return new MyQStyle();
}

void QStyle_DrawComplexControl(void* ptr, long long control, void* option, void* painter, void* widget)
{
		static_cast<QStyle*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void QStyle_DrawControl(void* ptr, long long element, void* option, void* painter, void* widget)
{
		static_cast<QStyle*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(element), static_cast<QStyleOption*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void QStyle_DrawPrimitive(void* ptr, long long element, void* option, void* painter, void* widget)
{
		static_cast<QStyle*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(element), static_cast<QStyleOption*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void* QStyle_GeneratedIconPixmap(void* ptr, long long iconMode, void* pixmap, void* option)
{
	return new QPixmap(static_cast<QStyle*>(ptr)->generatedIconPixmap(static_cast<QIcon::Mode>(iconMode), *static_cast<QPixmap*>(pixmap), static_cast<QStyleOption*>(option)));
}

long long QStyle_HitTestComplexControl(void* ptr, long long control, void* option, void* position, void* widget)
{
		return static_cast<QStyle*>(ptr)->hitTestComplexControl(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), *static_cast<QPoint*>(position), static_cast<QWidget*>(widget));
}

int QStyle_LayoutSpacing(void* ptr, long long control1, long long control2, long long orientation, void* option, void* widget)
{
		return static_cast<QStyle*>(ptr)->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

int QStyle_PixelMetric(void* ptr, long long metric, void* option, void* widget)
{
		return static_cast<QStyle*>(ptr)->pixelMetric(static_cast<QStyle::PixelMetric>(metric), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

void* QStyle_SizeFromContents(void* ptr, long long ty, void* option, void* contentsSize, void* widget)
{
		return ({ QSize tmpValue = static_cast<QStyle*>(ptr)->sizeFromContents(static_cast<QStyle::ContentsType>(ty), static_cast<QStyleOption*>(option), *static_cast<QSize*>(contentsSize), static_cast<QWidget*>(widget)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QStyle_StandardIcon(void* ptr, long long standardIcon, void* option, void* widget)
{
		return new QIcon(static_cast<QStyle*>(ptr)->standardIcon(static_cast<QStyle::StandardPixmap>(standardIcon), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget)));
}

int QStyle_StyleHint(void* ptr, long long hint, void* option, void* widget, void* returnData)
{
		return static_cast<QStyle*>(ptr)->styleHint(static_cast<QStyle::StyleHint>(hint), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget), static_cast<QStyleHintReturn*>(returnData));
}

void* QStyle_SubControlRect(void* ptr, long long control, void* option, long long subControl, void* widget)
{
		return ({ QRect tmpValue = static_cast<QStyle*>(ptr)->subControlRect(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), static_cast<QStyle::SubControl>(subControl), static_cast<QWidget*>(widget)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QStyle_SubElementRect(void* ptr, long long element, void* option, void* widget)
{
		return ({ QRect tmpValue = static_cast<QStyle*>(ptr)->subElementRect(static_cast<QStyle::SubElement>(element), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QStyle_DestroyQStyle(void* ptr)
{
	static_cast<QStyle*>(ptr)->~QStyle();
}

void QStyle_DestroyQStyleDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QStyle_StandardPixmap(void* ptr, long long standardIcon, void* option, void* widget)
{
		return new QPixmap(static_cast<QStyle*>(ptr)->standardPixmap(static_cast<QStyle::StandardPixmap>(standardIcon), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget)));
}

void* QStyle___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStyle___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QStyle___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QStyle___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QStyle___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QStyle___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QStyle___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStyle___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QStyle___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QStyle___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QStyle___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QStyle___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QStyle_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QStyle*>(ptr)->QStyle::childEvent(static_cast<QChildEvent*>(event));
}

void QStyle_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QStyle*>(ptr)->QStyle::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QStyle_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QStyle*>(ptr)->QStyle::customEvent(static_cast<QEvent*>(event));
}

void QStyle_DeleteLaterDefault(void* ptr)
{
		static_cast<QStyle*>(ptr)->QStyle::deleteLater();
}

void QStyle_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QStyle*>(ptr)->QStyle::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QStyle_EventDefault(void* ptr, void* e)
{
		return static_cast<QStyle*>(ptr)->QStyle::event(static_cast<QEvent*>(e));
}

char QStyle_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
		return static_cast<QStyle*>(ptr)->QStyle::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QStyle*>(ptr)->QStyle::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
		return static_cast<QStyle*>(ptr)->QStyle::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QStyle*>(ptr)->QStyle::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QStyle*>(ptr)->QStyle::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QStyle_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QStyle*>(ptr)->QStyle::timerEvent(static_cast<QTimerEvent*>(event));
}

Q_DECLARE_METATYPE(QStyleHintReturn*)
void* QStyleHintReturn_NewQStyleHintReturn(int version, int ty)
{
	return new QStyleHintReturn(version, ty);
}

int QStyleHintReturn_Type(void* ptr)
{
	return static_cast<QStyleHintReturn*>(ptr)->type;
}

int QStyleHintReturn_Version(void* ptr)
{
	return static_cast<QStyleHintReturn*>(ptr)->version;
}

Q_DECLARE_METATYPE(QStyleOption*)
void* QStyleOption_NewQStyleOption(int version, int ty)
{
	return new QStyleOption(version, ty);
}

void* QStyleOption_NewQStyleOption2(void* other)
{
	return new QStyleOption(*static_cast<QStyleOption*>(other));
}

void QStyleOption_DestroyQStyleOption(void* ptr)
{
	static_cast<QStyleOption*>(ptr)->~QStyleOption();
}

long long QStyleOption_Direction(void* ptr)
{
	return static_cast<QStyleOption*>(ptr)->direction;
}

void* QStyleOption_FontMetrics(void* ptr)
{
	return new QFontMetrics(static_cast<QStyleOption*>(ptr)->fontMetrics);
}

void* QStyleOption_Rect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QStyleOption*>(ptr)->rect; new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

int QStyleOption_Type(void* ptr)
{
	return static_cast<QStyleOption*>(ptr)->type;
}

int QStyleOption_Version(void* ptr)
{
	return static_cast<QStyleOption*>(ptr)->version;
}

Q_DECLARE_METATYPE(QStyleOptionComplex*)
void* QStyleOptionComplex_NewQStyleOptionComplex(int version, int ty)
{
	return new QStyleOptionComplex(version, ty);
}

void* QStyleOptionComplex_NewQStyleOptionComplex2(void* other)
{
	return new QStyleOptionComplex(*static_cast<QStyleOptionComplex*>(other));
}

Q_DECLARE_METATYPE(QStyleOptionGraphicsItem)
Q_DECLARE_METATYPE(QStyleOptionGraphicsItem*)
void* QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem()
{
	return new QStyleOptionGraphicsItem();
}

void* QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem2(void* other)
{
	return new QStyleOptionGraphicsItem(*static_cast<QStyleOptionGraphicsItem*>(other));
}

class MyQTextEdit: public QTextEdit
{
public:
	MyQTextEdit(QWidget *parent = Q_NULLPTR) : QTextEdit(parent) {QTextEdit_QTextEdit_QRegisterMetaType();};
	MyQTextEdit(const QString &text, QWidget *parent = Q_NULLPTR) : QTextEdit(text, parent) {QTextEdit_QTextEdit_QRegisterMetaType();};
	void append(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtWidgets_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQTextEdit_Append(this, textPacked); };
	void clear() { callbackQTextEdit_Clear(this); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	void copy() { callbackQTextEdit_Copy(this); };
	void insertPlainText(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtWidgets_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQTextEdit_InsertPlainText(this, textPacked); };
	void keyPressEvent(QKeyEvent * e) { callbackQWidget_KeyPressEvent(this, e); };
	void keyReleaseEvent(QKeyEvent * e) { callbackQWidget_KeyReleaseEvent(this, e); };
	void mouseDoubleClickEvent(QMouseEvent * e) { callbackQWidget_MouseDoubleClickEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackQWidget_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQWidget_MouseReleaseEvent(this, e); };
	void paste() { callbackQTextEdit_Paste(this); };
	void resizeEvent(QResizeEvent * e) { callbackQWidget_ResizeEvent(this, e); };
	void Signal_SelectionChanged() { callbackQTextEdit_SelectionChanged(this); };
	void setAlignment(Qt::Alignment a) { callbackQTextEdit_SetAlignment(this, a); };
	void setFontItalic(bool italic) { callbackQTextEdit_SetFontItalic(this, italic); };
	void setFontPointSize(qreal s) { callbackQTextEdit_SetFontPointSize(this, s); };
	void setFontUnderline(bool underline) { callbackQTextEdit_SetFontUnderline(this, underline); };
	void setFontWeight(int weight) { callbackQTextEdit_SetFontWeight(this, weight); };
	void setHtml(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtWidgets_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQTextEdit_SetHtml(this, textPacked); };
	void setText(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); QtWidgets_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackQTextEdit_SetText(this, textPacked); };
	void showEvent(QShowEvent * vqs) { callbackQWidget_ShowEvent(this, vqs); };
	void Signal_TextChanged() { callbackQTextEdit_TextChanged(this); };
	 ~MyQTextEdit() { callbackQTextEdit_DestroyQTextEdit(this); };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QTextEdit*)
Q_DECLARE_METATYPE(MyQTextEdit*)

int QTextEdit_QTextEdit_QRegisterMetaType(){qRegisterMetaType<QTextEdit*>(); return qRegisterMetaType<MyQTextEdit*>();}

void* QTextEdit_NewQTextEdit(void* parent)
{
		return new MyQTextEdit(static_cast<QWidget*>(parent));
}

void* QTextEdit_NewQTextEdit2(struct QtWidgets_PackedString text, void* parent)
{
		return new MyQTextEdit(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
}

long long QTextEdit_Alignment(void* ptr)
{
	return static_cast<QTextEdit*>(ptr)->alignment();
}

void QTextEdit_Append(void* ptr, struct QtWidgets_PackedString text)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "append", Q_ARG(const QString, QString::fromUtf8(text.data, text.len)));
}

void QTextEdit_AppendDefault(void* ptr, struct QtWidgets_PackedString text)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::append(QString::fromUtf8(text.data, text.len));
}

void QTextEdit_Clear(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "clear");
}

void QTextEdit_ClearDefault(void* ptr)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::clear();
}

void QTextEdit_Copy(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "copy");
}

void QTextEdit_CopyDefault(void* ptr)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::copy();
}

void* QTextEdit_CreateStandardContextMenu(void* ptr)
{
	return static_cast<QTextEdit*>(ptr)->createStandardContextMenu();
}

void* QTextEdit_CreateStandardContextMenu2(void* ptr, void* position)
{
	return static_cast<QTextEdit*>(ptr)->createStandardContextMenu(*static_cast<QPoint*>(position));
}

void* QTextEdit_CurrentCharFormat(void* ptr)
{
	return new QTextCharFormat(static_cast<QTextEdit*>(ptr)->currentCharFormat());
}

void* QTextEdit_CurrentFont(void* ptr)
{
	return new QFont(static_cast<QTextEdit*>(ptr)->currentFont());
}

void* QTextEdit_Document(void* ptr)
{
	return static_cast<QTextEdit*>(ptr)->document();
}

char QTextEdit_Find(void* ptr, struct QtWidgets_PackedString exp, long long options)
{
	return static_cast<QTextEdit*>(ptr)->find(QString::fromUtf8(exp.data, exp.len), static_cast<QTextDocument::FindFlag>(options));
}

char QTextEdit_Find2(void* ptr, void* exp, long long options)
{
	return static_cast<QTextEdit*>(ptr)->find(*static_cast<QRegExp*>(exp), static_cast<QTextDocument::FindFlag>(options));
}

char QTextEdit_Find3(void* ptr, void* exp, long long options)
{
	return static_cast<QTextEdit*>(ptr)->find(*static_cast<QRegularExpression*>(exp), static_cast<QTextDocument::FindFlag>(options));
}

char QTextEdit_FontItalic(void* ptr)
{
	return static_cast<QTextEdit*>(ptr)->fontItalic();
}

double QTextEdit_FontPointSize(void* ptr)
{
	return static_cast<QTextEdit*>(ptr)->fontPointSize();
}

char QTextEdit_FontUnderline(void* ptr)
{
	return static_cast<QTextEdit*>(ptr)->fontUnderline();
}

int QTextEdit_FontWeight(void* ptr)
{
	return static_cast<QTextEdit*>(ptr)->fontWeight();
}

void QTextEdit_InsertPlainText(void* ptr, struct QtWidgets_PackedString text)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "insertPlainText", Q_ARG(const QString, QString::fromUtf8(text.data, text.len)));
}

void QTextEdit_InsertPlainTextDefault(void* ptr, struct QtWidgets_PackedString text)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::insertPlainText(QString::fromUtf8(text.data, text.len));
}

void QTextEdit_MergeCurrentCharFormat(void* ptr, void* modifier)
{
	static_cast<QTextEdit*>(ptr)->mergeCurrentCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

void QTextEdit_Paste(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "paste");
}

void QTextEdit_PasteDefault(void* ptr)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::paste();
}

struct QtWidgets_PackedString QTextEdit_PlaceholderText(void* ptr)
{
	return ({ QByteArray* t79a2c0 = new QByteArray(static_cast<QTextEdit*>(ptr)->placeholderText().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t79a2c0->prepend("WHITESPACE").constData()+10), t79a2c0->size()-10, t79a2c0 }; });
}

void QTextEdit_Print(void* ptr, void* printer)
{
#ifndef Q_OS_IOS
	static_cast<QTextEdit*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
#endif
}

void QTextEdit_ConnectSelectionChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::selectionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_SelectionChanged), static_cast<Qt::ConnectionType>(t));
}

void QTextEdit_DisconnectSelectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::selectionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_SelectionChanged));
}

void QTextEdit_SelectionChanged(void* ptr)
{
	static_cast<QTextEdit*>(ptr)->selectionChanged();
}

void QTextEdit_SetAlignment(void* ptr, long long a)
{
	qRegisterMetaType<Qt::Alignment>();
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setAlignment", Q_ARG(Qt::Alignment, static_cast<Qt::Alignment>(static_cast<Qt::AlignmentFlag>(a))));
}

void QTextEdit_SetAlignmentDefault(void* ptr, long long a)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::setAlignment(static_cast<Qt::AlignmentFlag>(a));
}

void QTextEdit_SetCurrentCharFormat(void* ptr, void* format)
{
	static_cast<QTextEdit*>(ptr)->setCurrentCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextEdit_SetFontItalic(void* ptr, char italic)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontItalic", Q_ARG(bool, italic != 0));
}

void QTextEdit_SetFontItalicDefault(void* ptr, char italic)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::setFontItalic(italic != 0);
}

void QTextEdit_SetFontPointSize(void* ptr, double s)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontPointSize", Q_ARG(qreal, s));
}

void QTextEdit_SetFontPointSizeDefault(void* ptr, double s)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::setFontPointSize(s);
}

void QTextEdit_SetFontUnderline(void* ptr, char underline)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontUnderline", Q_ARG(bool, underline != 0));
}

void QTextEdit_SetFontUnderlineDefault(void* ptr, char underline)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::setFontUnderline(underline != 0);
}

void QTextEdit_SetFontWeight(void* ptr, int weight)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontWeight", Q_ARG(int, weight));
}

void QTextEdit_SetFontWeightDefault(void* ptr, int weight)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::setFontWeight(weight);
}

void QTextEdit_SetHtml(void* ptr, struct QtWidgets_PackedString text)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setHtml", Q_ARG(const QString, QString::fromUtf8(text.data, text.len)));
}

void QTextEdit_SetHtmlDefault(void* ptr, struct QtWidgets_PackedString text)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::setHtml(QString::fromUtf8(text.data, text.len));
}

void QTextEdit_SetPlaceholderText(void* ptr, struct QtWidgets_PackedString placeholderText)
{
	static_cast<QTextEdit*>(ptr)->setPlaceholderText(QString::fromUtf8(placeholderText.data, placeholderText.len));
}

void QTextEdit_SetReadOnly(void* ptr, char ro)
{
	static_cast<QTextEdit*>(ptr)->setReadOnly(ro != 0);
}

void QTextEdit_SetTabChangesFocus(void* ptr, char b)
{
	static_cast<QTextEdit*>(ptr)->setTabChangesFocus(b != 0);
}

void QTextEdit_SetTabStopWidth(void* ptr, int width)
{
	static_cast<QTextEdit*>(ptr)->setTabStopWidth(width);
}

void QTextEdit_SetText(void* ptr, struct QtWidgets_PackedString text)
{
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setText", Q_ARG(const QString, QString::fromUtf8(text.data, text.len)));
}

void QTextEdit_SetTextDefault(void* ptr, struct QtWidgets_PackedString text)
{
		static_cast<QTextEdit*>(ptr)->QTextEdit::setText(QString::fromUtf8(text.data, text.len));
}

void QTextEdit_SetTextCursor(void* ptr, void* cursor)
{
	static_cast<QTextEdit*>(ptr)->setTextCursor(*static_cast<QTextCursor*>(cursor));
}

char QTextEdit_TabChangesFocus(void* ptr)
{
	return static_cast<QTextEdit*>(ptr)->tabChangesFocus();
}

int QTextEdit_TabStopWidth(void* ptr)
{
	return static_cast<QTextEdit*>(ptr)->tabStopWidth();
}

void QTextEdit_ConnectTextChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::textChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_TextChanged), static_cast<Qt::ConnectionType>(t));
}

void QTextEdit_DisconnectTextChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::textChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_TextChanged));
}

void QTextEdit_TextChanged(void* ptr)
{
	static_cast<QTextEdit*>(ptr)->textChanged();
}

void* QTextEdit_TextColor(void* ptr)
{
	return new QColor(static_cast<QTextEdit*>(ptr)->textColor());
}

void* QTextEdit_TextCursor(void* ptr)
{
	return new QTextCursor(static_cast<QTextEdit*>(ptr)->textCursor());
}

struct QtWidgets_PackedString QTextEdit_ToHtml(void* ptr)
{
	return ({ QByteArray* t9adf2b = new QByteArray(static_cast<QTextEdit*>(ptr)->toHtml().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t9adf2b->prepend("WHITESPACE").constData()+10), t9adf2b->size()-10, t9adf2b }; });
}

struct QtWidgets_PackedString QTextEdit_ToPlainText(void* ptr)
{
	return ({ QByteArray* te62381 = new QByteArray(static_cast<QTextEdit*>(ptr)->toPlainText().toUtf8()); QtWidgets_PackedString { const_cast<char*>(te62381->prepend("WHITESPACE").constData()+10), te62381->size()-10, te62381 }; });
}

void QTextEdit_DestroyQTextEdit(void* ptr)
{
	static_cast<QTextEdit*>(ptr)->~QTextEdit();
}

void QTextEdit_DestroyQTextEditDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQToolBar: public QToolBar
{
public:
	MyQToolBar(const QString &title, QWidget *parent = Q_NULLPTR) : QToolBar(title, parent) {QToolBar_QToolBar_QRegisterMetaType();};
	MyQToolBar(QWidget *parent = Q_NULLPTR) : QToolBar(parent) {QToolBar_QToolBar_QRegisterMetaType();};
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	 ~MyQToolBar() { callbackQToolBar_DestroyQToolBar(this); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void update() { callbackQWidget_Update(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QToolBar*)
Q_DECLARE_METATYPE(MyQToolBar*)

int QToolBar_QToolBar_QRegisterMetaType(){qRegisterMetaType<QToolBar*>(); return qRegisterMetaType<MyQToolBar*>();}

void* QToolBar_NewQToolBar(struct QtWidgets_PackedString title, void* parent)
{
		return new MyQToolBar(QString::fromUtf8(title.data, title.len), static_cast<QWidget*>(parent));
}

void* QToolBar_NewQToolBar2(void* parent)
{
		return new MyQToolBar(static_cast<QWidget*>(parent));
}

void* QToolBar_AddAction(void* ptr, struct QtWidgets_PackedString text)
{
	return static_cast<QToolBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len));
}

void* QToolBar_AddAction2(void* ptr, void* icon, struct QtWidgets_PackedString text)
{
	return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len));
}

void* QToolBar_AddAction3(void* ptr, struct QtWidgets_PackedString text, void* receiver, char* member)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(receiver))) {
		return static_cast<QToolBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsObject*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(receiver))) {
		return static_cast<QToolBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QGraphicsWidget*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(receiver))) {
		return static_cast<QToolBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QLayout*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(receiver))) {
		return static_cast<QToolBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(receiver), const_cast<const char*>(member));
	} else {
		return static_cast<QToolBar*>(ptr)->addAction(QString::fromUtf8(text.data, text.len), static_cast<QObject*>(receiver), const_cast<const char*>(member));
	}
}

void* QToolBar_AddAction4(void* ptr, void* icon, struct QtWidgets_PackedString text, void* receiver, char* member)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(receiver))) {
		return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QGraphicsObject*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(receiver))) {
		return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QGraphicsWidget*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(receiver))) {
		return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QLayout*>(receiver), const_cast<const char*>(member));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(receiver))) {
		return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(receiver), const_cast<const char*>(member));
	} else {
		return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString::fromUtf8(text.data, text.len), static_cast<QObject*>(receiver), const_cast<const char*>(member));
	}
}

void* QToolBar_AddSeparator(void* ptr)
{
	return static_cast<QToolBar*>(ptr)->addSeparator();
}

void* QToolBar_AddWidget(void* ptr, void* widget)
{
		return static_cast<QToolBar*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

void QToolBar_Clear(void* ptr)
{
	static_cast<QToolBar*>(ptr)->clear();
}

long long QToolBar_Orientation(void* ptr)
{
	return static_cast<QToolBar*>(ptr)->orientation();
}

void QToolBar_SetOrientation(void* ptr, long long orientation)
{
	static_cast<QToolBar*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void QToolBar_DestroyQToolBar(void* ptr)
{
	static_cast<QToolBar*>(ptr)->~QToolBar();
}

void QToolBar_DestroyQToolBarDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQTreeView: public QTreeView
{
public:
	MyQTreeView(QWidget *parent = Q_NULLPTR) : QTreeView(parent) {QTreeView_QTreeView_QRegisterMetaType();};
	void collapse(const QModelIndex & index) { callbackQTreeView_Collapse(this, const_cast<QModelIndex*>(&index)); };
	void Signal_Collapsed(const QModelIndex & index) { callbackQTreeView_Collapsed(this, const_cast<QModelIndex*>(&index)); };
	void expand(const QModelIndex & index) { callbackQTreeView_Expand(this, const_cast<QModelIndex*>(&index)); };
	void Signal_Expanded(const QModelIndex & index) { callbackQTreeView_Expanded(this, const_cast<QModelIndex*>(&index)); };
	QModelIndex indexAt(const QPoint & point) const { return *static_cast<QModelIndex*>(callbackQTreeView_IndexAt(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPoint*>(&point))); };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void resizeColumnToContents(int column) { callbackQTreeView_ResizeColumnToContents(this, column); };
	void selectionChanged(const QItemSelection & selected, const QItemSelection & deselected) { callbackQTreeView_SelectionChanged(this, const_cast<QItemSelection*>(&selected), const_cast<QItemSelection*>(&deselected)); };
	void setModel(QAbstractItemModel * model) { callbackQAbstractItemView_SetModel(this, model); };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags command) { callbackQTreeView_SetSelection(this, const_cast<QRect*>(&rect), command); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	 ~MyQTreeView() { callbackQTreeView_DestroyQTreeView(this); };
	void Signal_Activated(const QModelIndex & index) { callbackQAbstractItemView_Activated(this, const_cast<QModelIndex*>(&index)); };
	void clearSelection() { callbackQAbstractItemView_ClearSelection(this); };
	void Signal_Clicked(const QModelIndex & index) { callbackQAbstractItemView_Clicked(this, const_cast<QModelIndex*>(&index)); };
	void edit(const QModelIndex & index) { callbackQAbstractItemView_Edit(this, const_cast<QModelIndex*>(&index)); };
	bool edit(const QModelIndex & index, QAbstractItemView::EditTrigger trigger, QEvent * event) { return callbackQAbstractItemView_Edit2(this, const_cast<QModelIndex*>(&index), trigger, event) != 0; };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	bool eventFilter(QObject * object, QEvent * event) { return callbackQWidget_EventFilter(this, object, event) != 0; };
	int horizontalOffset() const { return callbackQTreeView_HorizontalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	bool isIndexHidden(const QModelIndex & index) const { return callbackQTreeView_IsIndexHidden(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index)) != 0; };
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, Qt::KeyboardModifiers modifiers) { return *static_cast<QModelIndex*>(callbackQTreeView_MoveCursor(this, cursorAction, modifiers)); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void scrollTo(const QModelIndex & index, QAbstractItemView::ScrollHint hint) { callbackQTreeView_ScrollTo(this, const_cast<QModelIndex*>(&index), hint); };
	void setCurrentIndex(const QModelIndex & index) { callbackQAbstractItemView_SetCurrentIndex(this, const_cast<QModelIndex*>(&index)); };
	void update(const QModelIndex & index) { callbackQAbstractItemView_Update(this, const_cast<QModelIndex*>(&index)); };
	int verticalOffset() const { return callbackQTreeView_VerticalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	QRect visualRect(const QModelIndex & index) const { return *static_cast<QRect*>(callbackQTreeView_VisualRect(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QRegion visualRegionForSelection(const QItemSelection & selection) const { return *static_cast<QRegion*>(callbackQTreeView_VisualRegionForSelection(const_cast<void*>(static_cast<const void*>(this)), const_cast<QItemSelection*>(&selection))); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQWidget_ContextMenuEvent(this, e); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QTreeView*)
Q_DECLARE_METATYPE(MyQTreeView*)

int QTreeView_QTreeView_QRegisterMetaType(){qRegisterMetaType<QTreeView*>(); return qRegisterMetaType<MyQTreeView*>();}

void* QTreeView_NewQTreeView(void* parent)
{
		return new MyQTreeView(static_cast<QWidget*>(parent));
}

void QTreeView_Collapse(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "collapse", Q_ARG(const QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QTreeView_CollapseDefault(void* ptr, void* index)
{
		static_cast<QTreeView*>(ptr)->QTreeView::collapse(*static_cast<QModelIndex*>(index));
}

void QTreeView_ConnectCollapsed(void* ptr, long long t)
{
	QObject::connect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::collapsed), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Collapsed), static_cast<Qt::ConnectionType>(t));
}

void QTreeView_DisconnectCollapsed(void* ptr)
{
	QObject::disconnect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::collapsed), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Collapsed));
}

void QTreeView_Collapsed(void* ptr, void* index)
{
	static_cast<QTreeView*>(ptr)->collapsed(*static_cast<QModelIndex*>(index));
}

void QTreeView_Expand(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "expand", Q_ARG(const QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QTreeView_ExpandDefault(void* ptr, void* index)
{
		static_cast<QTreeView*>(ptr)->QTreeView::expand(*static_cast<QModelIndex*>(index));
}

void QTreeView_ConnectExpanded(void* ptr, long long t)
{
	QObject::connect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::expanded), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Expanded), static_cast<Qt::ConnectionType>(t));
}

void QTreeView_DisconnectExpanded(void* ptr)
{
	QObject::disconnect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::expanded), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Expanded));
}

void QTreeView_Expanded(void* ptr, void* index)
{
	static_cast<QTreeView*>(ptr)->expanded(*static_cast<QModelIndex*>(index));
}

void* QTreeView_Header(void* ptr)
{
	return static_cast<QTreeView*>(ptr)->header();
}

void* QTreeView_IndexAt(void* ptr, void* point)
{
	return new QModelIndex(static_cast<QTreeView*>(ptr)->indexAt(*static_cast<QPoint*>(point)));
}

void* QTreeView_IndexAtDefault(void* ptr, void* point)
{
		return new QModelIndex(static_cast<QTreeView*>(ptr)->QTreeView::indexAt(*static_cast<QPoint*>(point)));
}

char QTreeView_IsExpanded(void* ptr, void* index)
{
	return static_cast<QTreeView*>(ptr)->isExpanded(*static_cast<QModelIndex*>(index));
}

void QTreeView_ResizeColumnToContents(void* ptr, int column)
{
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "resizeColumnToContents", Q_ARG(int, column));
}

void QTreeView_ResizeColumnToContentsDefault(void* ptr, int column)
{
		static_cast<QTreeView*>(ptr)->QTreeView::resizeColumnToContents(column);
}

void QTreeView_SelectionChanged(void* ptr, void* selected, void* deselected)
{
	static_cast<QTreeView*>(ptr)->selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
}

void QTreeView_SelectionChangedDefault(void* ptr, void* selected, void* deselected)
{
		static_cast<QTreeView*>(ptr)->QTreeView::selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
}

void QTreeView_SetSelection(void* ptr, void* rect, long long command)
{
	static_cast<QTreeView*>(ptr)->setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTreeView_SetSelectionDefault(void* ptr, void* rect, long long command)
{
		static_cast<QTreeView*>(ptr)->QTreeView::setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTreeView_DestroyQTreeView(void* ptr)
{
	static_cast<QTreeView*>(ptr)->~QTreeView();
}

void QTreeView_DestroyQTreeViewDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

int QTreeView_HorizontalOffset(void* ptr)
{
	return static_cast<QTreeView*>(ptr)->horizontalOffset();
}

int QTreeView_HorizontalOffsetDefault(void* ptr)
{
		return static_cast<QTreeView*>(ptr)->QTreeView::horizontalOffset();
}

char QTreeView_IsIndexHidden(void* ptr, void* index)
{
	return static_cast<QTreeView*>(ptr)->isIndexHidden(*static_cast<QModelIndex*>(index));
}

char QTreeView_IsIndexHiddenDefault(void* ptr, void* index)
{
		return static_cast<QTreeView*>(ptr)->QTreeView::isIndexHidden(*static_cast<QModelIndex*>(index));
}

void* QTreeView_MoveCursor(void* ptr, long long cursorAction, long long modifiers)
{
	return new QModelIndex(static_cast<QTreeView*>(ptr)->moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

void* QTreeView_MoveCursorDefault(void* ptr, long long cursorAction, long long modifiers)
{
		return new QModelIndex(static_cast<QTreeView*>(ptr)->QTreeView::moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

void QTreeView_ScrollTo(void* ptr, void* index, long long hint)
{
	static_cast<QTreeView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QTreeView_ScrollToDefault(void* ptr, void* index, long long hint)
{
		static_cast<QTreeView*>(ptr)->QTreeView::scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

int QTreeView_VerticalOffset(void* ptr)
{
	return static_cast<QTreeView*>(ptr)->verticalOffset();
}

int QTreeView_VerticalOffsetDefault(void* ptr)
{
		return static_cast<QTreeView*>(ptr)->QTreeView::verticalOffset();
}

void* QTreeView_VisualRect(void* ptr, void* index)
{
	return ({ QRect tmpValue = static_cast<QTreeView*>(ptr)->visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QTreeView_VisualRectDefault(void* ptr, void* index)
{
		return ({ QRect tmpValue = static_cast<QTreeView*>(ptr)->QTreeView::visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QTreeView_VisualRegionForSelection(void* ptr, void* selection)
{
	return new QRegion(static_cast<QTreeView*>(ptr)->visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

void* QTreeView_VisualRegionForSelectionDefault(void* ptr, void* selection)
{
		return new QRegion(static_cast<QTreeView*>(ptr)->QTreeView::visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

class MyQVBoxLayout: public QVBoxLayout
{
public:
	MyQVBoxLayout() : QVBoxLayout() {QVBoxLayout_QVBoxLayout_QRegisterMetaType();};
	MyQVBoxLayout(QWidget *parent) : QVBoxLayout(parent) {QVBoxLayout_QVBoxLayout_QRegisterMetaType();};
	 ~MyQVBoxLayout() { callbackQVBoxLayout_DestroyQVBoxLayout(this); };
	void addItem(QLayoutItem * item) { callbackQBoxLayout_AddItem(this, item); };
	int count() const { return callbackQBoxLayout_Count(const_cast<void*>(static_cast<const void*>(this))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQLayout_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQLayoutItem_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQLayoutItem_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void invalidate() { callbackQLayoutItem_Invalidate(this); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQBoxLayout_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQLayout_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	int minimumHeightForWidth(int w) const { return callbackQLayoutItem_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQLayout_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	void setGeometry(const QRect & r) { callbackQLayout_SetGeometry(this, const_cast<QRect*>(&r)); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQBoxLayout_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQBoxLayout_TakeAt(this, index)); };
	void childEvent(QChildEvent * e) { callbackQLayout_ChildEvent(this, e); };
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackQLayoutItem_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQLayout_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	int indexOf(QWidget * widget) const { return callbackQLayout_IndexOf(const_cast<void*>(static_cast<const void*>(this)), widget); };
	bool isEmpty() const { return callbackQLayout_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	QLayout * layout() { return static_cast<QLayout*>(callbackQLayoutItem_Layout(this)); };
	void connectNotify(const QMetaMethod & sign) { callbackQLayout_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLayout_CustomEvent(this, event); };
	void deleteLater() { callbackQLayout_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLayout_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLayout_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQLayout_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLayout_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQLayout_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLayout_TimerEvent(this, event); };
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackQLayoutItem_SpacerItem(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQLayoutItem_Widget(this)); };
};

Q_DECLARE_METATYPE(QVBoxLayout*)
Q_DECLARE_METATYPE(MyQVBoxLayout*)

int QVBoxLayout_QVBoxLayout_QRegisterMetaType(){qRegisterMetaType<QVBoxLayout*>(); return qRegisterMetaType<MyQVBoxLayout*>();}

void* QVBoxLayout_NewQVBoxLayout()
{
	return new MyQVBoxLayout();
}

void* QVBoxLayout_NewQVBoxLayout2(void* parent)
{
		return new MyQVBoxLayout(static_cast<QWidget*>(parent));
}

void QVBoxLayout_DestroyQVBoxLayout(void* ptr)
{
	static_cast<QVBoxLayout*>(ptr)->~QVBoxLayout();
}

void QVBoxLayout_DestroyQVBoxLayoutDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQWidget: public QWidget
{
public:
	MyQWidget(QWidget *parent = Q_NULLPTR, Qt::WindowFlags ff = Qt::WindowFlags()) : QWidget(parent, ff) {QWidget_QWidget_QRegisterMetaType();};
	bool close() { return callbackQWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQWidget_ContextMenuEvent(this, event); };
	bool event(QEvent * event) { return callbackQWidget_Event(this, event) != 0; };
	bool hasHeightForWidth() const { return callbackQWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackQWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQWidget_HideEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQWidget_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQWidget_MouseReleaseEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQWidget_SetFocus2(this); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); QtWidgets_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackQWidget_SetStyleSheet(this, styleSheetPacked); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); QtWidgets_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackQWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQWidget_ShowEvent(this, event); };
	void showMaximized() { callbackQWidget_ShowMaximized(this); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void update() { callbackQWidget_Update(this); };
	 ~MyQWidget() { callbackQWidget_DestroyQWidget(this); };
	void childEvent(QChildEvent * event) { callbackQWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWidget_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtWidgets_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWidget_TimerEvent(this, event); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QWidget*)
Q_DECLARE_METATYPE(MyQWidget*)

int QWidget_QWidget_QRegisterMetaType(){qRegisterMetaType<QWidget*>(); return qRegisterMetaType<MyQWidget*>();}

void* QWidget_NewQWidget(void* parent, long long ff)
{
	return new MyQWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(ff));
}

struct QtWidgets_PackedString QWidget_AccessibleDescription(void* ptr)
{
		return ({ QByteArray* tb88bc9 = new QByteArray(static_cast<QWidget*>(ptr)->accessibleDescription().toUtf8()); QtWidgets_PackedString { const_cast<char*>(tb88bc9->prepend("WHITESPACE").constData()+10), tb88bc9->size()-10, tb88bc9 }; });
}

struct QtWidgets_PackedList QWidget_Actions(void* ptr)
{
		return ({ QList<QAction *>* tmpValue268e58 = new QList<QAction *>(static_cast<QWidget*>(ptr)->actions()); QtWidgets_PackedList { tmpValue268e58, tmpValue268e58->size() }; });
}

void QWidget_ActivateWindow(void* ptr)
{
		static_cast<QWidget*>(ptr)->activateWindow();
}

void QWidget_AddAction(void* ptr, void* action)
{
		static_cast<QWidget*>(ptr)->addAction(static_cast<QAction*>(action));
}

char QWidget_Close(void* ptr)
{
		bool returnArg;
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QWidget_CloseDefault(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QToolBar*>(ptr)->QToolBar::close();
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QStatusBar*>(ptr)->QStatusBar::close();
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenuBar*>(ptr)->QMenuBar::close();
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenu*>(ptr)->QMenu::close();
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMainWindow*>(ptr)->QMainWindow::close();
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLineEdit*>(ptr)->QLineEdit::close();
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSplitter*>(ptr)->QSplitter::close();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLabel*>(ptr)->QLabel::close();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextEdit*>(ptr)->QTextEdit::close();
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollArea*>(ptr)->QScrollArea::close();
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTreeView*>(ptr)->QTreeView::close();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHeaderView*>(ptr)->QHeaderView::close();
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::close();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::close();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFrame*>(ptr)->QFrame::close();
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDockWidget*>(ptr)->QDockWidget::close();
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMessageBox*>(ptr)->QMessageBox::close();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QInputDialog*>(ptr)->QInputDialog::close();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFontDialog*>(ptr)->QFontDialog::close();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDialog*>(ptr)->QFileDialog::close();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QColorDialog*>(ptr)->QColorDialog::close();
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDialog*>(ptr)->QDialog::close();
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QComboBox*>(ptr)->QComboBox::close();
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollBar*>(ptr)->QScrollBar::close();
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDial*>(ptr)->QDial::close();
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::close();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPushButton*>(ptr)->QPushButton::close();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractButton*>(ptr)->QAbstractButton::close();
	} else {
		return static_cast<QWidget*>(ptr)->QWidget::close();
	}
}

void QWidget_CloseEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QWidget_CloseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::closeEvent(static_cast<QCloseEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::closeEvent(static_cast<QCloseEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::closeEvent(static_cast<QCloseEvent*>(event));
	}
}

void QWidget_ContextMenuEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QWidget_ContextMenuEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
	}
}

void QWidget_Create(void* ptr, uintptr_t window, char initializeWindow, char destroyOldWindow)
{
		static_cast<QWidget*>(ptr)->create(window, initializeWindow != 0, destroyOldWindow != 0);
}

void* QWidget_Cursor(void* ptr)
{
		return new QCursor(static_cast<QWidget*>(ptr)->cursor());
}

void QWidget_Destroy(void* ptr, char destroyWindow, char destroySubWindows)
{
		static_cast<QWidget*>(ptr)->destroy(destroyWindow != 0, destroySubWindows != 0);
}

char QWidget_Event(void* ptr, void* event)
{
		return static_cast<QWidget*>(ptr)->event(static_cast<QEvent*>(event));
}

char QWidget_EventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QToolBar*>(ptr)->QToolBar::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QStatusBar*>(ptr)->QStatusBar::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenuBar*>(ptr)->QMenuBar::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenu*>(ptr)->QMenu::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMainWindow*>(ptr)->QMainWindow::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLineEdit*>(ptr)->QLineEdit::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSplitter*>(ptr)->QSplitter::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLabel*>(ptr)->QLabel::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextEdit*>(ptr)->QTextEdit::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollArea*>(ptr)->QScrollArea::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTreeView*>(ptr)->QTreeView::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHeaderView*>(ptr)->QHeaderView::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFrame*>(ptr)->QFrame::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDockWidget*>(ptr)->QDockWidget::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMessageBox*>(ptr)->QMessageBox::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QInputDialog*>(ptr)->QInputDialog::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFontDialog*>(ptr)->QFontDialog::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDialog*>(ptr)->QFileDialog::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QColorDialog*>(ptr)->QColorDialog::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDialog*>(ptr)->QDialog::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QComboBox*>(ptr)->QComboBox::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollBar*>(ptr)->QScrollBar::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDial*>(ptr)->QDial::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPushButton*>(ptr)->QPushButton::event(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractButton*>(ptr)->QAbstractButton::event(static_cast<QEvent*>(event));
	} else {
		return static_cast<QWidget*>(ptr)->QWidget::event(static_cast<QEvent*>(event));
	}
}

void* QWidget_QWidget_Find(uintptr_t id)
{
		return QWidget::find(id);
}

void* QWidget_Font(void* ptr)
{
		return const_cast<QFont*>(&static_cast<QWidget*>(ptr)->font());
}

void* QWidget_FontMetrics(void* ptr)
{
		return new QFontMetrics(static_cast<QWidget*>(ptr)->fontMetrics());
}

void* QWidget_Geometry(void* ptr)
{
		return const_cast<QRect*>(&static_cast<QWidget*>(ptr)->geometry());
}

char QWidget_HasHeightForWidth(void* ptr)
{
		return static_cast<QWidget*>(ptr)->hasHeightForWidth();
}

char QWidget_HasHeightForWidthDefault(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QToolBar*>(ptr)->QToolBar::hasHeightForWidth();
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QStatusBar*>(ptr)->QStatusBar::hasHeightForWidth();
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenuBar*>(ptr)->QMenuBar::hasHeightForWidth();
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenu*>(ptr)->QMenu::hasHeightForWidth();
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMainWindow*>(ptr)->QMainWindow::hasHeightForWidth();
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLineEdit*>(ptr)->QLineEdit::hasHeightForWidth();
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSplitter*>(ptr)->QSplitter::hasHeightForWidth();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLabel*>(ptr)->QLabel::hasHeightForWidth();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextEdit*>(ptr)->QTextEdit::hasHeightForWidth();
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollArea*>(ptr)->QScrollArea::hasHeightForWidth();
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTreeView*>(ptr)->QTreeView::hasHeightForWidth();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHeaderView*>(ptr)->QHeaderView::hasHeightForWidth();
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::hasHeightForWidth();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::hasHeightForWidth();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFrame*>(ptr)->QFrame::hasHeightForWidth();
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDockWidget*>(ptr)->QDockWidget::hasHeightForWidth();
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMessageBox*>(ptr)->QMessageBox::hasHeightForWidth();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QInputDialog*>(ptr)->QInputDialog::hasHeightForWidth();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFontDialog*>(ptr)->QFontDialog::hasHeightForWidth();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDialog*>(ptr)->QFileDialog::hasHeightForWidth();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QColorDialog*>(ptr)->QColorDialog::hasHeightForWidth();
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDialog*>(ptr)->QDialog::hasHeightForWidth();
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QComboBox*>(ptr)->QComboBox::hasHeightForWidth();
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollBar*>(ptr)->QScrollBar::hasHeightForWidth();
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDial*>(ptr)->QDial::hasHeightForWidth();
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::hasHeightForWidth();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPushButton*>(ptr)->QPushButton::hasHeightForWidth();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractButton*>(ptr)->QAbstractButton::hasHeightForWidth();
	} else {
		return static_cast<QWidget*>(ptr)->QWidget::hasHeightForWidth();
	}
}

int QWidget_Height(void* ptr)
{
		return static_cast<QWidget*>(ptr)->height();
}

int QWidget_HeightForWidth(void* ptr, int w)
{
		return static_cast<QWidget*>(ptr)->heightForWidth(w);
}

int QWidget_HeightForWidthDefault(void* ptr, int w)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QToolBar*>(ptr)->QToolBar::heightForWidth(w);
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QStatusBar*>(ptr)->QStatusBar::heightForWidth(w);
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenuBar*>(ptr)->QMenuBar::heightForWidth(w);
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenu*>(ptr)->QMenu::heightForWidth(w);
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMainWindow*>(ptr)->QMainWindow::heightForWidth(w);
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLineEdit*>(ptr)->QLineEdit::heightForWidth(w);
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSplitter*>(ptr)->QSplitter::heightForWidth(w);
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLabel*>(ptr)->QLabel::heightForWidth(w);
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextEdit*>(ptr)->QTextEdit::heightForWidth(w);
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollArea*>(ptr)->QScrollArea::heightForWidth(w);
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTreeView*>(ptr)->QTreeView::heightForWidth(w);
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHeaderView*>(ptr)->QHeaderView::heightForWidth(w);
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::heightForWidth(w);
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::heightForWidth(w);
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFrame*>(ptr)->QFrame::heightForWidth(w);
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDockWidget*>(ptr)->QDockWidget::heightForWidth(w);
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMessageBox*>(ptr)->QMessageBox::heightForWidth(w);
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QInputDialog*>(ptr)->QInputDialog::heightForWidth(w);
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFontDialog*>(ptr)->QFontDialog::heightForWidth(w);
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDialog*>(ptr)->QFileDialog::heightForWidth(w);
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QColorDialog*>(ptr)->QColorDialog::heightForWidth(w);
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDialog*>(ptr)->QDialog::heightForWidth(w);
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QComboBox*>(ptr)->QComboBox::heightForWidth(w);
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollBar*>(ptr)->QScrollBar::heightForWidth(w);
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDial*>(ptr)->QDial::heightForWidth(w);
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::heightForWidth(w);
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPushButton*>(ptr)->QPushButton::heightForWidth(w);
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractButton*>(ptr)->QAbstractButton::heightForWidth(w);
	} else {
		return static_cast<QWidget*>(ptr)->QWidget::heightForWidth(w);
	}
}

void QWidget_Hide(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "hide");
}

void QWidget_HideDefault(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::hide();
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::hide();
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::hide();
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::hide();
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::hide();
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::hide();
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::hide();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::hide();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::hide();
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::hide();
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::hide();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::hide();
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::hide();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::hide();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::hide();
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::hide();
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::hide();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::hide();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::hide();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::hide();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::hide();
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::hide();
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::hide();
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::hide();
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::hide();
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::hide();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::hide();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::hide();
	} else {
		static_cast<QWidget*>(ptr)->QWidget::hide();
	}
}

void QWidget_HideEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QWidget_HideEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::hideEvent(static_cast<QHideEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::hideEvent(static_cast<QHideEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::hideEvent(static_cast<QHideEvent*>(event));
	}
}

void QWidget_KeyPressEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QWidget_KeyPressEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
	}
}

void QWidget_KeyReleaseEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	}
}

void* QWidget_Layout(void* ptr)
{
		return static_cast<QWidget*>(ptr)->layout();
}

void QWidget_Lower(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "lower");
}

void QWidget_LowerDefault(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::lower();
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::lower();
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::lower();
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::lower();
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::lower();
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::lower();
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::lower();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::lower();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::lower();
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::lower();
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::lower();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::lower();
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::lower();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::lower();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::lower();
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::lower();
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::lower();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::lower();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::lower();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::lower();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::lower();
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::lower();
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::lower();
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::lower();
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::lower();
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::lower();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::lower();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::lower();
	} else {
		static_cast<QWidget*>(ptr)->QWidget::lower();
	}
}

void* QWidget_MaximumSize(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWidget*>(ptr)->maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QWidget_MaximumWidth(void* ptr)
{
		return static_cast<QWidget*>(ptr)->maximumWidth();
}

int QWidget_Metric(void* ptr, long long m)
{
		return static_cast<QWidget*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

int QWidget_MetricDefault(void* ptr, long long m)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QToolBar*>(ptr)->QToolBar::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QStatusBar*>(ptr)->QStatusBar::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenuBar*>(ptr)->QMenuBar::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenu*>(ptr)->QMenu::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMainWindow*>(ptr)->QMainWindow::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLineEdit*>(ptr)->QLineEdit::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSplitter*>(ptr)->QSplitter::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLabel*>(ptr)->QLabel::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextEdit*>(ptr)->QTextEdit::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollArea*>(ptr)->QScrollArea::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTreeView*>(ptr)->QTreeView::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHeaderView*>(ptr)->QHeaderView::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFrame*>(ptr)->QFrame::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDockWidget*>(ptr)->QDockWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMessageBox*>(ptr)->QMessageBox::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QInputDialog*>(ptr)->QInputDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFontDialog*>(ptr)->QFontDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDialog*>(ptr)->QFileDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QColorDialog*>(ptr)->QColorDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDialog*>(ptr)->QDialog::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QComboBox*>(ptr)->QComboBox::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollBar*>(ptr)->QScrollBar::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDial*>(ptr)->QDial::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPushButton*>(ptr)->QPushButton::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractButton*>(ptr)->QAbstractButton::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	} else {
		return static_cast<QWidget*>(ptr)->QWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
	}
}

int QWidget_MinimumHeight(void* ptr)
{
		return static_cast<QWidget*>(ptr)->minimumHeight();
}

void* QWidget_MinimumSize(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWidget*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QWidget_MinimumWidth(void* ptr)
{
		return static_cast<QWidget*>(ptr)->minimumWidth();
}

void QWidget_MouseDoubleClickEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
	}
}

void QWidget_MousePressEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QWidget_MousePressEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::mousePressEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
	}
}

void QWidget_MouseReleaseEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
	}
}

void QWidget_Move(void* ptr, void* vqp)
{
		static_cast<QWidget*>(ptr)->move(*static_cast<QPoint*>(vqp));
}

void QWidget_Move2(void* ptr, int x, int y)
{
		static_cast<QWidget*>(ptr)->move(x, y);
}

void* QWidget_Pos(void* ptr)
{
		return ({ QPoint tmpValue = static_cast<QWidget*>(ptr)->pos(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QWidget_Rect(void* ptr)
{
		return ({ QRect tmpValue = static_cast<QWidget*>(ptr)->rect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QWidget_RemoveAction(void* ptr, void* action)
{
		static_cast<QWidget*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QWidget_Resize(void* ptr, void* vqs)
{
		static_cast<QWidget*>(ptr)->resize(*static_cast<QSize*>(vqs));
}

void QWidget_Resize2(void* ptr, int w, int h)
{
		static_cast<QWidget*>(ptr)->resize(w, h);
}

void QWidget_ResizeEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QWidget_ResizeEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::resizeEvent(static_cast<QResizeEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::resizeEvent(static_cast<QResizeEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::resizeEvent(static_cast<QResizeEvent*>(event));
	}
}

void QWidget_Scroll(void* ptr, int dx, int dy)
{
		static_cast<QWidget*>(ptr)->scroll(dx, dy);
}

void QWidget_Scroll2(void* ptr, int dx, int dy, void* r)
{
		static_cast<QWidget*>(ptr)->scroll(dx, dy, *static_cast<QRect*>(r));
}

void QWidget_SetAccessibleDescription(void* ptr, struct QtWidgets_PackedString description)
{
		static_cast<QWidget*>(ptr)->setAccessibleDescription(QString::fromUtf8(description.data, description.len));
}

void QWidget_SetDisabled(void* ptr, char disable)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QWidget_SetDisabledDefault(void* ptr, char disable)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::setDisabled(disable != 0);
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::setDisabled(disable != 0);
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::setDisabled(disable != 0);
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::setDisabled(disable != 0);
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::setDisabled(disable != 0);
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::setDisabled(disable != 0);
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::setDisabled(disable != 0);
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::setDisabled(disable != 0);
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::setDisabled(disable != 0);
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::setDisabled(disable != 0);
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::setDisabled(disable != 0);
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::setDisabled(disable != 0);
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::setDisabled(disable != 0);
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::setDisabled(disable != 0);
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::setDisabled(disable != 0);
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::setDisabled(disable != 0);
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::setDisabled(disable != 0);
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::setDisabled(disable != 0);
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::setDisabled(disable != 0);
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::setDisabled(disable != 0);
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::setDisabled(disable != 0);
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::setDisabled(disable != 0);
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::setDisabled(disable != 0);
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::setDisabled(disable != 0);
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::setDisabled(disable != 0);
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::setDisabled(disable != 0);
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::setDisabled(disable != 0);
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::setDisabled(disable != 0);
	} else {
		static_cast<QWidget*>(ptr)->QWidget::setDisabled(disable != 0);
	}
}

void QWidget_SetEnabled(void* ptr, char vbo)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QWidget_SetEnabledDefault(void* ptr, char vbo)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::setEnabled(vbo != 0);
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::setEnabled(vbo != 0);
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::setEnabled(vbo != 0);
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::setEnabled(vbo != 0);
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::setEnabled(vbo != 0);
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::setEnabled(vbo != 0);
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::setEnabled(vbo != 0);
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::setEnabled(vbo != 0);
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::setEnabled(vbo != 0);
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::setEnabled(vbo != 0);
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::setEnabled(vbo != 0);
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::setEnabled(vbo != 0);
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::setEnabled(vbo != 0);
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::setEnabled(vbo != 0);
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::setEnabled(vbo != 0);
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::setEnabled(vbo != 0);
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::setEnabled(vbo != 0);
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::setEnabled(vbo != 0);
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::setEnabled(vbo != 0);
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::setEnabled(vbo != 0);
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::setEnabled(vbo != 0);
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::setEnabled(vbo != 0);
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::setEnabled(vbo != 0);
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::setEnabled(vbo != 0);
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::setEnabled(vbo != 0);
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::setEnabled(vbo != 0);
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::setEnabled(vbo != 0);
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::setEnabled(vbo != 0);
	} else {
		static_cast<QWidget*>(ptr)->QWidget::setEnabled(vbo != 0);
	}
}

void QWidget_SetFixedSize(void* ptr, void* s)
{
		static_cast<QWidget*>(ptr)->setFixedSize(*static_cast<QSize*>(s));
}

void QWidget_SetFixedSize2(void* ptr, int w, int h)
{
		static_cast<QWidget*>(ptr)->setFixedSize(w, h);
}

void QWidget_SetFixedWidth(void* ptr, int w)
{
		static_cast<QWidget*>(ptr)->setFixedWidth(w);
}

void QWidget_SetFocus(void* ptr, long long reason)
{
		static_cast<QWidget*>(ptr)->setFocus(static_cast<Qt::FocusReason>(reason));
}

void QWidget_SetFocus2(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setFocus");
}

void QWidget_SetFocus2Default(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::setFocus();
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::setFocus();
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::setFocus();
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::setFocus();
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::setFocus();
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::setFocus();
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::setFocus();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::setFocus();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::setFocus();
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::setFocus();
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::setFocus();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::setFocus();
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::setFocus();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::setFocus();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::setFocus();
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::setFocus();
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::setFocus();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::setFocus();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::setFocus();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::setFocus();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::setFocus();
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::setFocus();
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::setFocus();
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::setFocus();
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::setFocus();
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::setFocus();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::setFocus();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::setFocus();
	} else {
		static_cast<QWidget*>(ptr)->QWidget::setFocus();
	}
}

void QWidget_SetFont(void* ptr, void* vqf)
{
		static_cast<QWidget*>(ptr)->setFont(*static_cast<QFont*>(vqf));
}

void QWidget_SetGeometry(void* ptr, void* vqr)
{
		static_cast<QWidget*>(ptr)->setGeometry(*static_cast<QRect*>(vqr));
}

void QWidget_SetGeometry2(void* ptr, int x, int y, int w, int h)
{
		static_cast<QWidget*>(ptr)->setGeometry(x, y, w, h);
}

void QWidget_SetLayout(void* ptr, void* layout)
{
		static_cast<QWidget*>(ptr)->setLayout(static_cast<QLayout*>(layout));
}

void QWidget_SetMaximumWidth(void* ptr, int maxw)
{
		static_cast<QWidget*>(ptr)->setMaximumWidth(maxw);
}

void QWidget_SetMinimumSize(void* ptr, void* vqs)
{
		static_cast<QWidget*>(ptr)->setMinimumSize(*static_cast<QSize*>(vqs));
}

void QWidget_SetMinimumSize2(void* ptr, int minw, int minh)
{
		static_cast<QWidget*>(ptr)->setMinimumSize(minw, minh);
}

void QWidget_SetMinimumWidth(void* ptr, int minw)
{
		static_cast<QWidget*>(ptr)->setMinimumWidth(minw);
}

void QWidget_SetSizePolicy(void* ptr, void* vqs)
{
		static_cast<QWidget*>(ptr)->setSizePolicy(*static_cast<QSizePolicy*>(vqs));
}

void QWidget_SetSizePolicy2(void* ptr, long long horizontal, long long vertical)
{
		static_cast<QWidget*>(ptr)->setSizePolicy(static_cast<QSizePolicy::Policy>(horizontal), static_cast<QSizePolicy::Policy>(vertical));
}

void QWidget_SetStyle(void* ptr, void* style)
{
		static_cast<QWidget*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

void QWidget_SetStyleSheet(void* ptr, struct QtWidgets_PackedString styleSheet)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setStyleSheet", Q_ARG(const QString, QString::fromUtf8(styleSheet.data, styleSheet.len)));
}

void QWidget_SetStyleSheetDefault(void* ptr, struct QtWidgets_PackedString styleSheet)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
	}
}

void QWidget_SetToolTip(void* ptr, struct QtWidgets_PackedString vqs)
{
		static_cast<QWidget*>(ptr)->setToolTip(QString::fromUtf8(vqs.data, vqs.len));
}

void QWidget_SetWindowIcon(void* ptr, void* icon)
{
		static_cast<QWidget*>(ptr)->setWindowIcon(*static_cast<QIcon*>(icon));
}

void QWidget_SetWindowTitle(void* ptr, struct QtWidgets_PackedString vqs)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setWindowTitle", Q_ARG(const QString, QString::fromUtf8(vqs.data, vqs.len)));
}

void QWidget_SetWindowTitleDefault(void* ptr, struct QtWidgets_PackedString vqs)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
	}
}

void QWidget_Show(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "show");
}

void QWidget_ShowDefault(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::show();
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::show();
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::show();
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::show();
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::show();
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::show();
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::show();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::show();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::show();
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::show();
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::show();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::show();
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::show();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::show();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::show();
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::show();
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::show();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::show();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::show();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::show();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::show();
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::show();
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::show();
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::show();
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::show();
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::show();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::show();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::show();
	} else {
		static_cast<QWidget*>(ptr)->QWidget::show();
	}
}

void QWidget_ShowEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QWidget_ShowEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::showEvent(static_cast<QShowEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::showEvent(static_cast<QShowEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::showEvent(static_cast<QShowEvent*>(event));
	}
}

void QWidget_ShowMaximized(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showMaximized");
}

void QWidget_ShowMaximizedDefault(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::showMaximized();
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::showMaximized();
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::showMaximized();
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::showMaximized();
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::showMaximized();
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::showMaximized();
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::showMaximized();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::showMaximized();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::showMaximized();
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::showMaximized();
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::showMaximized();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::showMaximized();
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::showMaximized();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::showMaximized();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::showMaximized();
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::showMaximized();
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::showMaximized();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::showMaximized();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::showMaximized();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::showMaximized();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::showMaximized();
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::showMaximized();
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::showMaximized();
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::showMaximized();
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::showMaximized();
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::showMaximized();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::showMaximized();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::showMaximized();
	} else {
		static_cast<QWidget*>(ptr)->QWidget::showMaximized();
	}
}

void* QWidget_Size(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWidget*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QWidget_SizeHint(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWidget*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QWidget_SizeHintDefault(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QToolBar*>(ptr)->QToolBar::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QStatusBar*>(ptr)->QStatusBar::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QMenuBar*>(ptr)->QMenuBar::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QMenu*>(ptr)->QMenu::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QMainWindow*>(ptr)->QMainWindow::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QLineEdit*>(ptr)->QLineEdit::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QSplitter*>(ptr)->QSplitter::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QLabel*>(ptr)->QLabel::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QTextEdit*>(ptr)->QTextEdit::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QScrollArea*>(ptr)->QScrollArea::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QTreeView*>(ptr)->QTreeView::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QHeaderView*>(ptr)->QHeaderView::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QFrame*>(ptr)->QFrame::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QDockWidget*>(ptr)->QDockWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QMessageBox*>(ptr)->QMessageBox::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QInputDialog*>(ptr)->QInputDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QFontDialog*>(ptr)->QFontDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QFileDialog*>(ptr)->QFileDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QColorDialog*>(ptr)->QColorDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QDialog*>(ptr)->QDialog::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QComboBox*>(ptr)->QComboBox::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QScrollBar*>(ptr)->QScrollBar::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QDial*>(ptr)->QDial::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QPushButton*>(ptr)->QPushButton::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QAbstractButton*>(ptr)->QAbstractButton::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QWidget*>(ptr)->QWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

void* QWidget_SizePolicy(void* ptr)
{
		return new QSizePolicy(static_cast<QWidget*>(ptr)->sizePolicy());
}

void* QWidget_Style(void* ptr)
{
		return static_cast<QWidget*>(ptr)->style();
}

struct QtWidgets_PackedString QWidget_StyleSheet(void* ptr)
{
		return ({ QByteArray* tb180d7 = new QByteArray(static_cast<QWidget*>(ptr)->styleSheet().toUtf8()); QtWidgets_PackedString { const_cast<char*>(tb180d7->prepend("WHITESPACE").constData()+10), tb180d7->size()-10, tb180d7 }; });
}

struct QtWidgets_PackedString QWidget_ToolTip(void* ptr)
{
		return ({ QByteArray* t791287 = new QByteArray(static_cast<QWidget*>(ptr)->toolTip().toUtf8()); QtWidgets_PackedString { const_cast<char*>(t791287->prepend("WHITESPACE").constData()+10), t791287->size()-10, t791287 }; });
}

void QWidget_Update(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "update");
}

void QWidget_UpdateDefault(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::update();
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::update();
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::update();
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::update();
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::update();
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::update();
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::update();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::update();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::update();
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::update();
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::update();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::update();
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::update();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::update();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::update();
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::update();
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::update();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::update();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::update();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::update();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::update();
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::update();
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::update();
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::update();
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::update();
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::update();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::update();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::update();
	} else {
		static_cast<QWidget*>(ptr)->QWidget::update();
	}
}

void QWidget_Update2(void* ptr, int x, int y, int w, int h)
{
		static_cast<QWidget*>(ptr)->update(x, y, w, h);
}

void QWidget_Update3(void* ptr, void* rect)
{
		static_cast<QWidget*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QWidget_Update4(void* ptr, void* rgn)
{
		static_cast<QWidget*>(ptr)->update(*static_cast<QRegion*>(rgn));
}

int QWidget_Width(void* ptr)
{
		return static_cast<QWidget*>(ptr)->width();
}

void* QWidget_Window(void* ptr)
{
		return static_cast<QWidget*>(ptr)->window();
}

void* QWidget_WindowIcon(void* ptr)
{
		return new QIcon(static_cast<QWidget*>(ptr)->windowIcon());
}

struct QtWidgets_PackedString QWidget_WindowTitle(void* ptr)
{
		return ({ QByteArray* tf3cd6c = new QByteArray(static_cast<QWidget*>(ptr)->windowTitle().toUtf8()); QtWidgets_PackedString { const_cast<char*>(tf3cd6c->prepend("WHITESPACE").constData()+10), tf3cd6c->size()-10, tf3cd6c }; });
}

int QWidget_X(void* ptr)
{
		return static_cast<QWidget*>(ptr)->x();
}

int QWidget_Y(void* ptr)
{
		return static_cast<QWidget*>(ptr)->y();
}

void QWidget_DestroyQWidget(void* ptr)
{
	static_cast<QWidget*>(ptr)->~QWidget();
}

void QWidget_DestroyQWidgetDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWidget___actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QWidget___addActions_actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___addActions_actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QWidget___insertActions_actions_atList(void* ptr, int i)
{
		return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___insertActions_actions_setList(void* ptr, void* i)
{
		static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QAction *>();
}

void* QWidget___children_atList(void* ptr, int i)
{
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWidget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>();
}

void* QWidget___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWidget___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWidget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>();
}

void* QWidget___findChildren_atList(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWidget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QWidget___findChildren_atList3(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWidget___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsObject*>(i));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QGraphicsWidget*>(i));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QLayout*>(i));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWidget*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWidget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void QWidget_ChildEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWidget_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QWidget_ConnectNotify(void* ptr, void* sign)
{
		static_cast<QWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QWidget_CustomEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWidget_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::customEvent(static_cast<QEvent*>(event));
	}
}

void QWidget_DeleteLater(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "deleteLater");
}

void QWidget_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::deleteLater();
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::deleteLater();
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::deleteLater();
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::deleteLater();
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::deleteLater();
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::deleteLater();
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::deleteLater();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::deleteLater();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::deleteLater();
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::deleteLater();
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::deleteLater();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::deleteLater();
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::deleteLater();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::deleteLater();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::deleteLater();
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::deleteLater();
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::deleteLater();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::deleteLater();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::deleteLater();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::deleteLater();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::deleteLater();
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::deleteLater();
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::deleteLater();
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::deleteLater();
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::deleteLater();
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::deleteLater();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::deleteLater();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::deleteLater();
	} else {
		static_cast<QWidget*>(ptr)->QWidget::deleteLater();
	}
}

void QWidget_DisconnectNotify(void* ptr, void* sign)
{
		static_cast<QWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QWidget_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
		return static_cast<QWidget*>(ptr)->eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QWidget*>(ptr)->eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
		return static_cast<QWidget*>(ptr)->eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
		return static_cast<QWidget*>(ptr)->eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QToolBar*>(ptr)->QToolBar::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QToolBar*>(ptr)->QToolBar::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QToolBar*>(ptr)->QToolBar::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QToolBar*>(ptr)->QToolBar::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QToolBar*>(ptr)->QToolBar::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QStatusBar*>(ptr)->QStatusBar::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QStatusBar*>(ptr)->QStatusBar::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QStatusBar*>(ptr)->QStatusBar::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QStatusBar*>(ptr)->QStatusBar::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QStatusBar*>(ptr)->QStatusBar::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QMenuBar*>(ptr)->QMenuBar::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QMenuBar*>(ptr)->QMenuBar::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QMenuBar*>(ptr)->QMenuBar::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QMenuBar*>(ptr)->QMenuBar::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QMenuBar*>(ptr)->QMenuBar::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QMenu*>(ptr)->QMenu::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QMenu*>(ptr)->QMenu::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QMenu*>(ptr)->QMenu::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QMenu*>(ptr)->QMenu::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QMenu*>(ptr)->QMenu::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QMainWindow*>(ptr)->QMainWindow::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QMainWindow*>(ptr)->QMainWindow::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QMainWindow*>(ptr)->QMainWindow::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QMainWindow*>(ptr)->QMainWindow::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QMainWindow*>(ptr)->QMainWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QLineEdit*>(ptr)->QLineEdit::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QLineEdit*>(ptr)->QLineEdit::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QLineEdit*>(ptr)->QLineEdit::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QLineEdit*>(ptr)->QLineEdit::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QLineEdit*>(ptr)->QLineEdit::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QSplitter*>(ptr)->QSplitter::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QSplitter*>(ptr)->QSplitter::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QSplitter*>(ptr)->QSplitter::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QSplitter*>(ptr)->QSplitter::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QSplitter*>(ptr)->QSplitter::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QLabel*>(ptr)->QLabel::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QLabel*>(ptr)->QLabel::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QLabel*>(ptr)->QLabel::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QLabel*>(ptr)->QLabel::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QLabel*>(ptr)->QLabel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QTextEdit*>(ptr)->QTextEdit::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QTextEdit*>(ptr)->QTextEdit::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QTextEdit*>(ptr)->QTextEdit::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QTextEdit*>(ptr)->QTextEdit::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QTextEdit*>(ptr)->QTextEdit::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QScrollArea*>(ptr)->QScrollArea::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QScrollArea*>(ptr)->QScrollArea::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QScrollArea*>(ptr)->QScrollArea::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QScrollArea*>(ptr)->QScrollArea::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QScrollArea*>(ptr)->QScrollArea::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QTreeView*>(ptr)->QTreeView::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QTreeView*>(ptr)->QTreeView::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QTreeView*>(ptr)->QTreeView::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QTreeView*>(ptr)->QTreeView::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QTreeView*>(ptr)->QTreeView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QHeaderView*>(ptr)->QHeaderView::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QHeaderView*>(ptr)->QHeaderView::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QHeaderView*>(ptr)->QHeaderView::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QHeaderView*>(ptr)->QHeaderView::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QHeaderView*>(ptr)->QHeaderView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QFrame*>(ptr)->QFrame::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QFrame*>(ptr)->QFrame::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QFrame*>(ptr)->QFrame::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QFrame*>(ptr)->QFrame::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QFrame*>(ptr)->QFrame::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QDockWidget*>(ptr)->QDockWidget::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QDockWidget*>(ptr)->QDockWidget::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QDockWidget*>(ptr)->QDockWidget::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QDockWidget*>(ptr)->QDockWidget::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QDockWidget*>(ptr)->QDockWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QMessageBox*>(ptr)->QMessageBox::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QMessageBox*>(ptr)->QMessageBox::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QMessageBox*>(ptr)->QMessageBox::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QMessageBox*>(ptr)->QMessageBox::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QMessageBox*>(ptr)->QMessageBox::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QInputDialog*>(ptr)->QInputDialog::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QInputDialog*>(ptr)->QInputDialog::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QInputDialog*>(ptr)->QInputDialog::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QInputDialog*>(ptr)->QInputDialog::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QInputDialog*>(ptr)->QInputDialog::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QFontDialog*>(ptr)->QFontDialog::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QFontDialog*>(ptr)->QFontDialog::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QFontDialog*>(ptr)->QFontDialog::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QFontDialog*>(ptr)->QFontDialog::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QFontDialog*>(ptr)->QFontDialog::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QFileDialog*>(ptr)->QFileDialog::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QFileDialog*>(ptr)->QFileDialog::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QFileDialog*>(ptr)->QFileDialog::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QFileDialog*>(ptr)->QFileDialog::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QFileDialog*>(ptr)->QFileDialog::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QColorDialog*>(ptr)->QColorDialog::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QColorDialog*>(ptr)->QColorDialog::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QColorDialog*>(ptr)->QColorDialog::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QColorDialog*>(ptr)->QColorDialog::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QColorDialog*>(ptr)->QColorDialog::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QDialog*>(ptr)->QDialog::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QDialog*>(ptr)->QDialog::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QDialog*>(ptr)->QDialog::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QDialog*>(ptr)->QDialog::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QDialog*>(ptr)->QDialog::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QComboBox*>(ptr)->QComboBox::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QComboBox*>(ptr)->QComboBox::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QComboBox*>(ptr)->QComboBox::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QComboBox*>(ptr)->QComboBox::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QComboBox*>(ptr)->QComboBox::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QScrollBar*>(ptr)->QScrollBar::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QScrollBar*>(ptr)->QScrollBar::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QScrollBar*>(ptr)->QScrollBar::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QScrollBar*>(ptr)->QScrollBar::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QScrollBar*>(ptr)->QScrollBar::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QDial*>(ptr)->QDial::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QDial*>(ptr)->QDial::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QDial*>(ptr)->QDial::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QDial*>(ptr)->QDial::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QDial*>(ptr)->QDial::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QPushButton*>(ptr)->QPushButton::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QPushButton*>(ptr)->QPushButton::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QPushButton*>(ptr)->QPushButton::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QPushButton*>(ptr)->QPushButton::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QPushButton*>(ptr)->QPushButton::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractButton*>(ptr)->QAbstractButton::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractButton*>(ptr)->QAbstractButton::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractButton*>(ptr)->QAbstractButton::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QAbstractButton*>(ptr)->QAbstractButton::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QAbstractButton*>(ptr)->QAbstractButton::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(watched))) {
			return static_cast<QWidget*>(ptr)->QWidget::eventFilter(static_cast<QGraphicsObject*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QWidget*>(ptr)->QWidget::eventFilter(static_cast<QGraphicsWidget*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(watched))) {
			return static_cast<QWidget*>(ptr)->QWidget::eventFilter(static_cast<QLayout*>(watched), static_cast<QEvent*>(event));
		} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(watched))) {
			return static_cast<QWidget*>(ptr)->QWidget::eventFilter(static_cast<QWidget*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QWidget*>(ptr)->QWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void QWidget_TimerEvent(void* ptr, void* event)
{
		static_cast<QWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWidget_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QToolBar*>(ptr)->QToolBar::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QStatusBar*>(ptr)->QStatusBar::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenuBar*>(ptr)->QMenuBar::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		static_cast<QMenu*>(ptr)->QMenu::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		static_cast<QMainWindow*>(ptr)->QMainWindow::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineEdit*>(ptr)->QLineEdit::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplitter*>(ptr)->QSplitter::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		static_cast<QLabel*>(ptr)->QLabel::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		static_cast<QTextEdit*>(ptr)->QTextEdit::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollArea*>(ptr)->QScrollArea::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		static_cast<QTreeView*>(ptr)->QTreeView::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		static_cast<QHeaderView*>(ptr)->QHeaderView::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		static_cast<QFrame*>(ptr)->QFrame::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		static_cast<QDockWidget*>(ptr)->QDockWidget::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QMessageBox*>(ptr)->QMessageBox::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QInputDialog*>(ptr)->QInputDialog::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFontDialog*>(ptr)->QFontDialog::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDialog*>(ptr)->QFileDialog::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QColorDialog*>(ptr)->QColorDialog::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		static_cast<QDialog*>(ptr)->QDialog::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		static_cast<QComboBox*>(ptr)->QComboBox::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		static_cast<QScrollBar*>(ptr)->QScrollBar::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		static_cast<QDial*>(ptr)->QDial::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QPushButton*>(ptr)->QPushButton::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractButton*>(ptr)->QAbstractButton::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QWidget*>(ptr)->QWidget::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QWidget_PaintEngine(void* ptr)
{
		return static_cast<QWidget*>(ptr)->paintEngine();
}

void* QWidget_PaintEngineDefault(void* ptr)
{
	if (dynamic_cast<QToolBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QToolBar*>(ptr)->QToolBar::paintEngine();
	} else if (dynamic_cast<QStatusBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QStatusBar*>(ptr)->QStatusBar::paintEngine();
	} else if (dynamic_cast<QMenuBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenuBar*>(ptr)->QMenuBar::paintEngine();
	} else if (dynamic_cast<QMenu*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMenu*>(ptr)->QMenu::paintEngine();
	} else if (dynamic_cast<QMainWindow*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMainWindow*>(ptr)->QMainWindow::paintEngine();
	} else if (dynamic_cast<QLineEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLineEdit*>(ptr)->QLineEdit::paintEngine();
	} else if (dynamic_cast<QSplitter*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSplitter*>(ptr)->QSplitter::paintEngine();
	} else if (dynamic_cast<QLabel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLabel*>(ptr)->QLabel::paintEngine();
	} else if (dynamic_cast<QTextEdit*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTextEdit*>(ptr)->QTextEdit::paintEngine();
	} else if (dynamic_cast<QScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollArea*>(ptr)->QScrollArea::paintEngine();
	} else if (dynamic_cast<QTreeView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTreeView*>(ptr)->QTreeView::paintEngine();
	} else if (dynamic_cast<QHeaderView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHeaderView*>(ptr)->QHeaderView::paintEngine();
	} else if (dynamic_cast<QAbstractItemView*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractItemView*>(ptr)->QAbstractItemView::paintEngine();
	} else if (dynamic_cast<QAbstractScrollArea*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractScrollArea*>(ptr)->QAbstractScrollArea::paintEngine();
	} else if (dynamic_cast<QFrame*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFrame*>(ptr)->QFrame::paintEngine();
	} else if (dynamic_cast<QDockWidget*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDockWidget*>(ptr)->QDockWidget::paintEngine();
	} else if (dynamic_cast<QMessageBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMessageBox*>(ptr)->QMessageBox::paintEngine();
	} else if (dynamic_cast<QInputDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QInputDialog*>(ptr)->QInputDialog::paintEngine();
	} else if (dynamic_cast<QFontDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFontDialog*>(ptr)->QFontDialog::paintEngine();
	} else if (dynamic_cast<QFileDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDialog*>(ptr)->QFileDialog::paintEngine();
	} else if (dynamic_cast<QColorDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QColorDialog*>(ptr)->QColorDialog::paintEngine();
	} else if (dynamic_cast<QDialog*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDialog*>(ptr)->QDialog::paintEngine();
	} else if (dynamic_cast<QComboBox*>(static_cast<QObject*>(ptr))) {
		return static_cast<QComboBox*>(ptr)->QComboBox::paintEngine();
	} else if (dynamic_cast<QScrollBar*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScrollBar*>(ptr)->QScrollBar::paintEngine();
	} else if (dynamic_cast<QDial*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDial*>(ptr)->QDial::paintEngine();
	} else if (dynamic_cast<QAbstractSlider*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractSlider*>(ptr)->QAbstractSlider::paintEngine();
	} else if (dynamic_cast<QPushButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPushButton*>(ptr)->QPushButton::paintEngine();
	} else if (dynamic_cast<QAbstractButton*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractButton*>(ptr)->QAbstractButton::paintEngine();
	} else {
		return static_cast<QWidget*>(ptr)->QWidget::paintEngine();
	}
}

