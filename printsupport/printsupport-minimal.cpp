// +build minimal

#define protected public
#define private public

#include "printsupport-minimal.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QPageSize>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPrinter>
#include <QPrinterInfo>
#include <QString>
#include <QStringList>

class MyQPrinter: public QPrinter
{
public:
	MyQPrinter(QPrinter::PrinterMode mode = ScreenResolution) : QPrinter(mode) {QPrinter_QPrinter_QRegisterMetaType();};
	MyQPrinter(const QPrinterInfo &printer, QPrinter::PrinterMode mode = ScreenResolution) : QPrinter(printer, mode) {QPrinter_QPrinter_QRegisterMetaType();};
	 ~MyQPrinter() { callbackQPrinter_DestroyQPrinter(this); };
	bool newPage() { return callbackQPrinter_NewPage(this) != 0; };
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPrinter_Metric(const_cast<void*>(static_cast<const void*>(this)), metric); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPrinter_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(QPrinter*)
Q_DECLARE_METATYPE(MyQPrinter*)

int QPrinter_QPrinter_QRegisterMetaType(){qRegisterMetaType<QPrinter*>(); return qRegisterMetaType<MyQPrinter*>();}

void* QPrinter_NewQPrinter(long long mode)
{
	return new MyQPrinter(static_cast<QPrinter::PrinterMode>(mode));
}

void* QPrinter_NewQPrinter2(void* printer, long long mode)
{
	return new MyQPrinter(*static_cast<QPrinterInfo*>(printer), static_cast<QPrinter::PrinterMode>(mode));
}

char QPrinter_IsValid(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->isValid();
}

struct QtPrintSupport_PackedString QPrinter_OutputFileName(void* ptr)
{
	return ({ QByteArray* tde340f = new QByteArray(static_cast<QPrinter*>(ptr)->outputFileName().toUtf8()); QtPrintSupport_PackedString { const_cast<char*>(tde340f->prepend("WHITESPACE").constData()+10), tde340f->size()-10, tde340f }; });
}

long long QPrinter_OutputFormat(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->outputFormat();
}

int QPrinter_Resolution(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->resolution();
}

void QPrinter_SetOutputFileName(void* ptr, struct QtPrintSupport_PackedString fileName)
{
	static_cast<QPrinter*>(ptr)->setOutputFileName(QString::fromUtf8(fileName.data, fileName.len));
}

void QPrinter_SetOutputFormat(void* ptr, long long format)
{
	static_cast<QPrinter*>(ptr)->setOutputFormat(static_cast<QPrinter::OutputFormat>(format));
}

void QPrinter_DestroyQPrinter(void* ptr)
{
	static_cast<QPrinter*>(ptr)->~QPrinter();
}

void QPrinter_DestroyQPrinterDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

int QPrinter___supportedResolutions_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPrinter___supportedResolutions_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QPrinter___supportedResolutions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

char QPrinter_NewPage(void* ptr)
{
	return static_cast<QPrinter*>(ptr)->newPage();
}

char QPrinter_NewPageDefault(void* ptr)
{
		return static_cast<QPrinter*>(ptr)->QPrinter::newPage();
}

int QPrinter_MetricDefault(void* ptr, long long metric)
{
		return static_cast<QPrinter*>(ptr)->QPrinter::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QPrinter_PaintEngineDefault(void* ptr)
{
		return static_cast<QPrinter*>(ptr)->QPrinter::paintEngine();
}

Q_DECLARE_METATYPE(QPrinterInfo)
Q_DECLARE_METATYPE(QPrinterInfo*)
void* QPrinterInfo_NewQPrinterInfo()
{
	return new QPrinterInfo();
}

void* QPrinterInfo_NewQPrinterInfo2(void* other)
{
	return new QPrinterInfo(*static_cast<QPrinterInfo*>(other));
}

void* QPrinterInfo_NewQPrinterInfo3(void* printer)
{
	return new QPrinterInfo(*static_cast<QPrinter*>(printer));
}

struct QtPrintSupport_PackedString QPrinterInfo_Description(void* ptr)
{
	return ({ QByteArray* t0d7900 = new QByteArray(static_cast<QPrinterInfo*>(ptr)->description().toUtf8()); QtPrintSupport_PackedString { const_cast<char*>(t0d7900->prepend("WHITESPACE").constData()+10), t0d7900->size()-10, t0d7900 }; });
}

void QPrinterInfo_DestroyQPrinterInfo(void* ptr)
{
	static_cast<QPrinterInfo*>(ptr)->~QPrinterInfo();
}

void* QPrinterInfo___availablePrinters_atList(void* ptr, int i)
{
	return new QPrinterInfo(({QPrinterInfo tmp = static_cast<QList<QPrinterInfo>*>(ptr)->at(i); if (i == static_cast<QList<QPrinterInfo>*>(ptr)->size()-1) { static_cast<QList<QPrinterInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPrinterInfo___availablePrinters_setList(void* ptr, void* i)
{
	static_cast<QList<QPrinterInfo>*>(ptr)->append(*static_cast<QPrinterInfo*>(i));
}

void* QPrinterInfo___availablePrinters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPrinterInfo>();
}

long long QPrinterInfo___supportedColorModes_atList(void* ptr, int i)
{
	return ({QPrinter::ColorMode tmp = static_cast<QList<QPrinter::ColorMode>*>(ptr)->at(i); if (i == static_cast<QList<QPrinter::ColorMode>*>(ptr)->size()-1) { static_cast<QList<QPrinter::ColorMode>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPrinterInfo___supportedColorModes_setList(void* ptr, long long i)
{
	static_cast<QList<QPrinter::ColorMode>*>(ptr)->append(static_cast<QPrinter::ColorMode>(i));
}

void* QPrinterInfo___supportedColorModes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPrinter::ColorMode>();
}

long long QPrinterInfo___supportedDuplexModes_atList(void* ptr, int i)
{
	return ({QPrinter::DuplexMode tmp = static_cast<QList<QPrinter::DuplexMode>*>(ptr)->at(i); if (i == static_cast<QList<QPrinter::DuplexMode>*>(ptr)->size()-1) { static_cast<QList<QPrinter::DuplexMode>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPrinterInfo___supportedDuplexModes_setList(void* ptr, long long i)
{
	static_cast<QList<QPrinter::DuplexMode>*>(ptr)->append(static_cast<QPrinter::DuplexMode>(i));
}

void* QPrinterInfo___supportedDuplexModes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPrinter::DuplexMode>();
}

void* QPrinterInfo___supportedPageSizes_atList(void* ptr, int i)
{
	return new QPageSize(({QPageSize tmp = static_cast<QList<QPageSize>*>(ptr)->at(i); if (i == static_cast<QList<QPageSize>*>(ptr)->size()-1) { static_cast<QList<QPageSize>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPrinterInfo___supportedPageSizes_setList(void* ptr, void* i)
{
	static_cast<QList<QPageSize>*>(ptr)->append(*static_cast<QPageSize*>(i));
}

void* QPrinterInfo___supportedPageSizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPageSize>();
}

void* QPrinterInfo___supportedPaperSizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPrinter::PaperSize>();
}

int QPrinterInfo___supportedResolutions_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPrinterInfo___supportedResolutions_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QPrinterInfo___supportedResolutions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

