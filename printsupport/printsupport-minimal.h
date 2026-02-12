// +build minimal

#pragma once

#ifndef GO_QTPRINTSUPPORT_H
#define GO_QTPRINTSUPPORT_H

#include <stdint.h>

#ifdef __cplusplus
int QPrinter_QPrinter_QRegisterMetaType();
extern "C" {
#endif

struct QtPrintSupport_PackedString { char* data; long long len; void* ptr; };
struct QtPrintSupport_PackedList { void* data; long long len; };
void* QPrinter_NewQPrinter(long long mode);
void* QPrinter_NewQPrinter2(void* printer, long long mode);
char QPrinter_IsValid(void* ptr);
struct QtPrintSupport_PackedString QPrinter_OutputFileName(void* ptr);
long long QPrinter_OutputFormat(void* ptr);
int QPrinter_Resolution(void* ptr);
void QPrinter_SetOutputFileName(void* ptr, struct QtPrintSupport_PackedString fileName);
void QPrinter_SetOutputFormat(void* ptr, long long format);
void QPrinter_DestroyQPrinter(void* ptr);
void QPrinter_DestroyQPrinterDefault(void* ptr);
int QPrinter___supportedResolutions_atList(void* ptr, int i);
void QPrinter___supportedResolutions_setList(void* ptr, int i);
void* QPrinter___supportedResolutions_newList(void* ptr);
char QPrinter_NewPage(void* ptr);
char QPrinter_NewPageDefault(void* ptr);
int QPrinter_MetricDefault(void* ptr, long long metric);
void* QPrinter_PaintEngineDefault(void* ptr);
void* QPrinterInfo_NewQPrinterInfo();
void* QPrinterInfo_NewQPrinterInfo2(void* other);
void* QPrinterInfo_NewQPrinterInfo3(void* printer);
struct QtPrintSupport_PackedString QPrinterInfo_Description(void* ptr);
void QPrinterInfo_DestroyQPrinterInfo(void* ptr);
void* QPrinterInfo___availablePrinters_atList(void* ptr, int i);
void QPrinterInfo___availablePrinters_setList(void* ptr, void* i);
void* QPrinterInfo___availablePrinters_newList(void* ptr);
long long QPrinterInfo___supportedColorModes_atList(void* ptr, int i);
void QPrinterInfo___supportedColorModes_setList(void* ptr, long long i);
void* QPrinterInfo___supportedColorModes_newList(void* ptr);
long long QPrinterInfo___supportedDuplexModes_atList(void* ptr, int i);
void QPrinterInfo___supportedDuplexModes_setList(void* ptr, long long i);
void* QPrinterInfo___supportedDuplexModes_newList(void* ptr);
void* QPrinterInfo___supportedPageSizes_atList(void* ptr, int i);
void QPrinterInfo___supportedPageSizes_setList(void* ptr, void* i);
void* QPrinterInfo___supportedPageSizes_newList(void* ptr);
void* QPrinterInfo___supportedPaperSizes_newList(void* ptr);
int QPrinterInfo___supportedResolutions_atList(void* ptr, int i);
void QPrinterInfo___supportedResolutions_setList(void* ptr, int i);
void* QPrinterInfo___supportedResolutions_newList(void* ptr);

#ifdef __cplusplus
}
#endif

#endif