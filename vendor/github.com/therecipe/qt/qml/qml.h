// +build !minimal

#pragma once

#ifndef GO_QTQML_H
#define GO_QTQML_H

#include <stdint.h>

#ifdef __cplusplus
int QJSEngine_QJSEngine_QRegisterMetaType();
int QQmlAbstractUrlInterceptor_QQmlAbstractUrlInterceptor_QRegisterMetaType();
int QQmlApplicationEngine_QQmlApplicationEngine_QRegisterMetaType();
int QQmlComponent_QQmlComponent_QRegisterMetaType();
int QQmlContext_QQmlContext_QRegisterMetaType();
int QQmlEngine_QQmlEngine_QRegisterMetaType();
int QQmlEngineExtensionPlugin_QQmlEngineExtensionPlugin_QRegisterMetaType();
int QQmlExpression_QQmlExpression_QRegisterMetaType();
int QQmlFileSelector_QQmlFileSelector_QRegisterMetaType();
int QQmlImageProviderBase_QQmlImageProviderBase_QRegisterMetaType();
int QQmlIncubationController_QQmlIncubationController_QRegisterMetaType();
int QQmlIncubator_QQmlIncubator_QRegisterMetaType();
int QQmlNetworkAccessManagerFactory_QQmlNetworkAccessManagerFactory_QRegisterMetaType();
int QQmlParserStatus_QQmlParserStatus_QRegisterMetaType();
int QQmlPropertyMap_QQmlPropertyMap_QRegisterMetaType();
int QQmlPropertyValueSource_QQmlPropertyValueSource_QRegisterMetaType();
extern "C" {
#endif

struct QtQml_PackedString { char* data; long long len; void* ptr; };
struct QtQml_PackedList { void* data; long long len; };
void* QJSEngine_NewQJSEngine();
void* QJSEngine_NewQJSEngine2(void* parent);
void QJSEngine_CollectGarbage(void* ptr);
void* QJSEngine_Evaluate(void* ptr, struct QtQml_PackedString program, struct QtQml_PackedString fileName, int lineNumber);
void* QJSEngine_FromScriptValue(void* ptr, void* value);
void* QJSEngine_GlobalObject(void* ptr);
void* QJSEngine_ImportModule(void* ptr, struct QtQml_PackedString fileName);
void QJSEngine_InstallExtensions(void* ptr, long long extensions, void* object);
char QJSEngine_IsInterrupted(void* ptr);
void* QJSEngine_NewArray(void* ptr, unsigned int length);
void* QJSEngine_NewErrorObject(void* ptr, long long errorType, struct QtQml_PackedString message);
void* QJSEngine_NewObject(void* ptr);
void* QJSEngine_NewQMetaObject(void* ptr, void* metaObject);
void* QJSEngine_NewQObject(void* ptr, void* object);
void QJSEngine_SetInterrupted(void* ptr, char interrupted);
void QJSEngine_SetUiLanguage(void* ptr, struct QtQml_PackedString language);
void QJSEngine_ThrowError(void* ptr, struct QtQml_PackedString message);
void QJSEngine_ThrowError2(void* ptr, long long errorType, struct QtQml_PackedString message);
void* QJSEngine_ToScriptValue(void* ptr, void* value);
struct QtQml_PackedString QJSEngine_UiLanguage(void* ptr);
void QJSEngine_ConnectUiLanguageChanged(void* ptr, long long t);
void QJSEngine_DisconnectUiLanguageChanged(void* ptr);
void QJSEngine_UiLanguageChanged(void* ptr);
void QJSEngine_DestroyQJSEngine(void* ptr);
void QJSEngine_DestroyQJSEngineDefault(void* ptr);
void* QJSEngine_QJSEngine_qjsEngine(void* object);
void* QJSEngine___children_atList(void* ptr, int i);
void QJSEngine___children_setList(void* ptr, void* i);
void* QJSEngine___children_newList(void* ptr);
void* QJSEngine___dynamicPropertyNames_atList(void* ptr, int i);
void QJSEngine___dynamicPropertyNames_setList(void* ptr, void* i);
void* QJSEngine___dynamicPropertyNames_newList(void* ptr);
void* QJSEngine___findChildren_atList(void* ptr, int i);
void QJSEngine___findChildren_setList(void* ptr, void* i);
void* QJSEngine___findChildren_newList(void* ptr);
void* QJSEngine___findChildren_atList3(void* ptr, int i);
void QJSEngine___findChildren_setList3(void* ptr, void* i);
void* QJSEngine___findChildren_newList3(void* ptr);
void QJSEngine_ChildEventDefault(void* ptr, void* event);
void QJSEngine_ConnectNotifyDefault(void* ptr, void* sign);
void QJSEngine_CustomEventDefault(void* ptr, void* event);
void QJSEngine_DeleteLaterDefault(void* ptr);
void QJSEngine_DisconnectNotifyDefault(void* ptr, void* sign);
char QJSEngine_EventDefault(void* ptr, void* e);
char QJSEngine_EventFilterDefault(void* ptr, void* watched, void* event);
void* QJSEngine_MetaObjectDefault(void* ptr);
void QJSEngine_TimerEventDefault(void* ptr, void* event);
void* QJSValue_NewQJSValue(long long value);
void* QJSValue_NewQJSValue2(void* other);
void* QJSValue_NewQJSValue3(void* other);
void* QJSValue_NewQJSValue4(char value);
void* QJSValue_NewQJSValue5(int value);
void* QJSValue_NewQJSValue6(unsigned int value);
void* QJSValue_NewQJSValue7(double value);
void* QJSValue_NewQJSValue8(struct QtQml_PackedString value);
void* QJSValue_NewQJSValue9(void* value);
void* QJSValue_NewQJSValue10(char* value);
void* QJSValue_Call(void* ptr, void* args);
void* QJSValue_CallAsConstructor(void* ptr, void* args);
void* QJSValue_CallWithInstance(void* ptr, void* instance, void* args);
char QJSValue_DeleteProperty(void* ptr, struct QtQml_PackedString name);
char QJSValue_Equals(void* ptr, void* other);
long long QJSValue_ErrorType(void* ptr);
char QJSValue_HasOwnProperty(void* ptr, struct QtQml_PackedString name);
char QJSValue_HasProperty(void* ptr, struct QtQml_PackedString name);
char QJSValue_IsArray(void* ptr);
char QJSValue_IsBool(void* ptr);
char QJSValue_IsCallable(void* ptr);
char QJSValue_IsDate(void* ptr);
char QJSValue_IsError(void* ptr);
char QJSValue_IsNull(void* ptr);
char QJSValue_IsNumber(void* ptr);
char QJSValue_IsObject(void* ptr);
char QJSValue_IsQMetaObject(void* ptr);
char QJSValue_IsQObject(void* ptr);
char QJSValue_IsRegExp(void* ptr);
char QJSValue_IsString(void* ptr);
char QJSValue_IsUndefined(void* ptr);
char QJSValue_IsVariant(void* ptr);
void* QJSValue_Property(void* ptr, struct QtQml_PackedString name);
void* QJSValue_Property2(void* ptr, unsigned int arrayIndex);
void* QJSValue_Prototype(void* ptr);
void QJSValue_SetProperty(void* ptr, struct QtQml_PackedString name, void* value);
void QJSValue_SetProperty2(void* ptr, unsigned int arrayIndex, void* value);
void QJSValue_SetPrototype(void* ptr, void* prototype);
char QJSValue_StrictlyEquals(void* ptr, void* other);
char QJSValue_ToBool(void* ptr);
void* QJSValue_ToDateTime(void* ptr);
int QJSValue_ToInt(void* ptr);
double QJSValue_ToNumber(void* ptr);
void* QJSValue_ToQMetaObject(void* ptr);
void* QJSValue_ToQObject(void* ptr);
struct QtQml_PackedString QJSValue_ToString(void* ptr);
unsigned int QJSValue_ToUInt(void* ptr);
void* QJSValue_ToVariant(void* ptr);
void QJSValue_DestroyQJSValue(void* ptr);
void* QJSValue___call_args_atList(void* ptr, int i);
void QJSValue___call_args_setList(void* ptr, void* i);
void* QJSValue___call_args_newList(void* ptr);
void* QJSValue___callAsConstructor_args_atList(void* ptr, int i);
void QJSValue___callAsConstructor_args_setList(void* ptr, void* i);
void* QJSValue___callAsConstructor_args_newList(void* ptr);
void* QJSValue___callWithInstance_args_atList(void* ptr, int i);
void QJSValue___callWithInstance_args_setList(void* ptr, void* i);
void* QJSValue___callWithInstance_args_newList(void* ptr);
void* QJSValueIterator_NewQJSValueIterator(void* object);
char QJSValueIterator_HasNext(void* ptr);
struct QtQml_PackedString QJSValueIterator_Name(void* ptr);
char QJSValueIterator_Next(void* ptr);
void* QJSValueIterator_Value(void* ptr);
void QJSValueIterator_DestroyQJSValueIterator(void* ptr);
void* QQmlAbstractUrlInterceptor_NewQQmlAbstractUrlInterceptor();
void* QQmlAbstractUrlInterceptor_Intercept(void* ptr, void* url, long long ty);
void QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(void* ptr);
void QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptorDefault(void* ptr);
void* QQmlApplicationEngine_NewQQmlApplicationEngine(void* parent);
void* QQmlApplicationEngine_NewQQmlApplicationEngine2(void* url, void* parent);
void* QQmlApplicationEngine_NewQQmlApplicationEngine3(struct QtQml_PackedString filePath, void* parent);
void QQmlApplicationEngine_Load(void* ptr, void* url);
void QQmlApplicationEngine_LoadDefault(void* ptr, void* url);
void QQmlApplicationEngine_Load2(void* ptr, struct QtQml_PackedString filePath);
void QQmlApplicationEngine_Load2Default(void* ptr, struct QtQml_PackedString filePath);
void QQmlApplicationEngine_LoadData(void* ptr, void* data, void* url);
void QQmlApplicationEngine_LoadDataDefault(void* ptr, void* data, void* url);
void QQmlApplicationEngine_ConnectObjectCreated(void* ptr, long long t);
void QQmlApplicationEngine_DisconnectObjectCreated(void* ptr);
void QQmlApplicationEngine_ObjectCreated(void* ptr, void* object, void* url);
struct QtQml_PackedList QQmlApplicationEngine_RootObjects(void* ptr);
void QQmlApplicationEngine_DestroyQQmlApplicationEngine(void* ptr);
void QQmlApplicationEngine_DestroyQQmlApplicationEngineDefault(void* ptr);
void* QQmlApplicationEngine___rootObjects_atList(void* ptr, int i);
void QQmlApplicationEngine___rootObjects_setList(void* ptr, void* i);
void* QQmlApplicationEngine___rootObjects_newList(void* ptr);
void* QQmlApplicationEngine___setInitialProperties_initialProperties_atList(void* ptr, struct QtQml_PackedString v, int i);
void QQmlApplicationEngine___setInitialProperties_initialProperties_setList(void* ptr, struct QtQml_PackedString key, void* i);
void* QQmlApplicationEngine___setInitialProperties_initialProperties_newList(void* ptr);
struct QtQml_PackedList QQmlApplicationEngine___setInitialProperties_initialProperties_keyList(void* ptr);
struct QtQml_PackedString QQmlApplicationEngine_____setInitialProperties_initialProperties_keyList_atList(void* ptr, int i);
void QQmlApplicationEngine_____setInitialProperties_initialProperties_keyList_setList(void* ptr, struct QtQml_PackedString i);
void* QQmlApplicationEngine_____setInitialProperties_initialProperties_keyList_newList(void* ptr);
void* QQmlComponent_NewQQmlComponent2(void* engine, void* parent);
void* QQmlComponent_NewQQmlComponent3(void* engine, struct QtQml_PackedString fileName, void* parent);
void* QQmlComponent_NewQQmlComponent4(void* engine, struct QtQml_PackedString fileName, long long mode, void* parent);
void* QQmlComponent_NewQQmlComponent5(void* engine, void* url, void* parent);
void* QQmlComponent_NewQQmlComponent6(void* engine, void* url, long long mode, void* parent);
void* QQmlComponent_BeginCreate(void* ptr, void* publicContext);
void* QQmlComponent_BeginCreateDefault(void* ptr, void* publicContext);
void QQmlComponent_CompleteCreate(void* ptr);
void QQmlComponent_CompleteCreateDefault(void* ptr);
void* QQmlComponent_Create(void* ptr, void* context);
void* QQmlComponent_CreateDefault(void* ptr, void* context);
void QQmlComponent_Create2(void* ptr, void* incubator, void* context, void* forContext);
void* QQmlComponent_CreateWithInitialProperties(void* ptr, void* initialProperties, void* context);
void* QQmlComponent_CreationContext(void* ptr);
void* QQmlComponent_Engine(void* ptr);
struct QtQml_PackedList QQmlComponent_Errors(void* ptr);
char QQmlComponent_IsError(void* ptr);
char QQmlComponent_IsLoading(void* ptr);
char QQmlComponent_IsNull(void* ptr);
char QQmlComponent_IsReady(void* ptr);
void QQmlComponent_LoadUrl(void* ptr, void* url);
void QQmlComponent_LoadUrlDefault(void* ptr, void* url);
void QQmlComponent_LoadUrl2(void* ptr, void* url, long long mode);
void QQmlComponent_LoadUrl2Default(void* ptr, void* url, long long mode);
double QQmlComponent_Progress(void* ptr);
void QQmlComponent_ConnectProgressChanged(void* ptr, long long t);
void QQmlComponent_DisconnectProgressChanged(void* ptr);
void QQmlComponent_ProgressChanged(void* ptr, double progress);
void QQmlComponent_SetData(void* ptr, void* data, void* url);
void QQmlComponent_SetDataDefault(void* ptr, void* data, void* url);
long long QQmlComponent_Status(void* ptr);
void QQmlComponent_ConnectStatusChanged(void* ptr, long long t);
void QQmlComponent_DisconnectStatusChanged(void* ptr);
void QQmlComponent_StatusChanged(void* ptr, long long status);
void* QQmlComponent_Url(void* ptr);
void QQmlComponent_DestroyQQmlComponent(void* ptr);
void QQmlComponent_DestroyQQmlComponentDefault(void* ptr);
void* QQmlComponent___createWithInitialProperties_initialProperties_atList(void* ptr, struct QtQml_PackedString v, int i);
void QQmlComponent___createWithInitialProperties_initialProperties_setList(void* ptr, struct QtQml_PackedString key, void* i);
void* QQmlComponent___createWithInitialProperties_initialProperties_newList(void* ptr);
struct QtQml_PackedList QQmlComponent___createWithInitialProperties_initialProperties_keyList(void* ptr);
void* QQmlComponent___errors_atList(void* ptr, int i);
void QQmlComponent___errors_setList(void* ptr, void* i);
void* QQmlComponent___errors_newList(void* ptr);
void* QQmlComponent___setInitialProperties_properties_atList(void* ptr, struct QtQml_PackedString v, int i);
void QQmlComponent___setInitialProperties_properties_setList(void* ptr, struct QtQml_PackedString key, void* i);
void* QQmlComponent___setInitialProperties_properties_newList(void* ptr);
struct QtQml_PackedList QQmlComponent___setInitialProperties_properties_keyList(void* ptr);
struct QtQml_PackedString QQmlComponent_____createWithInitialProperties_initialProperties_keyList_atList(void* ptr, int i);
void QQmlComponent_____createWithInitialProperties_initialProperties_keyList_setList(void* ptr, struct QtQml_PackedString i);
void* QQmlComponent_____createWithInitialProperties_initialProperties_keyList_newList(void* ptr);
struct QtQml_PackedString QQmlComponent_____setInitialProperties_properties_keyList_atList(void* ptr, int i);
void QQmlComponent_____setInitialProperties_properties_keyList_setList(void* ptr, struct QtQml_PackedString i);
void* QQmlComponent_____setInitialProperties_properties_keyList_newList(void* ptr);
void* QQmlComponent___children_atList(void* ptr, int i);
void QQmlComponent___children_setList(void* ptr, void* i);
void* QQmlComponent___children_newList(void* ptr);
void* QQmlComponent___dynamicPropertyNames_atList(void* ptr, int i);
void QQmlComponent___dynamicPropertyNames_setList(void* ptr, void* i);
void* QQmlComponent___dynamicPropertyNames_newList(void* ptr);
void* QQmlComponent___findChildren_atList(void* ptr, int i);
void QQmlComponent___findChildren_setList(void* ptr, void* i);
void* QQmlComponent___findChildren_newList(void* ptr);
void* QQmlComponent___findChildren_atList3(void* ptr, int i);
void QQmlComponent___findChildren_setList3(void* ptr, void* i);
void* QQmlComponent___findChildren_newList3(void* ptr);
void QQmlComponent_ChildEventDefault(void* ptr, void* event);
void QQmlComponent_ConnectNotifyDefault(void* ptr, void* sign);
void QQmlComponent_CustomEventDefault(void* ptr, void* event);
void QQmlComponent_DeleteLaterDefault(void* ptr);
void QQmlComponent_DisconnectNotifyDefault(void* ptr, void* sign);
char QQmlComponent_EventDefault(void* ptr, void* e);
char QQmlComponent_EventFilterDefault(void* ptr, void* watched, void* event);
void* QQmlComponent_MetaObjectDefault(void* ptr);
void QQmlComponent_TimerEventDefault(void* ptr, void* event);
void* QQmlContext_NewQQmlContext(void* engine, void* parent);
void* QQmlContext_NewQQmlContext2(void* parentContext, void* parent);
void* QQmlContext_BaseUrl(void* ptr);
void* QQmlContext_ContextObject(void* ptr);
void* QQmlContext_ContextProperty(void* ptr, struct QtQml_PackedString name);
void* QQmlContext_Engine(void* ptr);
char QQmlContext_IsValid(void* ptr);
struct QtQml_PackedString QQmlContext_NameForObject(void* ptr, void* object);
void* QQmlContext_ParentContext(void* ptr);
void* QQmlContext_ResolvedUrl(void* ptr, void* src);
void QQmlContext_SetBaseUrl(void* ptr, void* baseUrl);
void QQmlContext_SetContextObject(void* ptr, void* object);
void QQmlContext_SetContextProperty(void* ptr, struct QtQml_PackedString name, void* value);
void QQmlContext_SetContextProperty2(void* ptr, struct QtQml_PackedString name, void* value);
void QQmlContext_DestroyQQmlContext(void* ptr);
void QQmlContext_DestroyQQmlContextDefault(void* ptr);
void* QQmlContext___setContextProperties_properties_newList(void* ptr);
void* QQmlContext___children_atList(void* ptr, int i);
void QQmlContext___children_setList(void* ptr, void* i);
void* QQmlContext___children_newList(void* ptr);
void* QQmlContext___dynamicPropertyNames_atList(void* ptr, int i);
void QQmlContext___dynamicPropertyNames_setList(void* ptr, void* i);
void* QQmlContext___dynamicPropertyNames_newList(void* ptr);
void* QQmlContext___findChildren_atList(void* ptr, int i);
void QQmlContext___findChildren_setList(void* ptr, void* i);
void* QQmlContext___findChildren_newList(void* ptr);
void* QQmlContext___findChildren_atList3(void* ptr, int i);
void QQmlContext___findChildren_setList3(void* ptr, void* i);
void* QQmlContext___findChildren_newList3(void* ptr);
void QQmlContext_ChildEventDefault(void* ptr, void* event);
void QQmlContext_ConnectNotifyDefault(void* ptr, void* sign);
void QQmlContext_CustomEventDefault(void* ptr, void* event);
void QQmlContext_DeleteLaterDefault(void* ptr);
void QQmlContext_DisconnectNotifyDefault(void* ptr, void* sign);
char QQmlContext_EventDefault(void* ptr, void* e);
char QQmlContext_EventFilterDefault(void* ptr, void* watched, void* event);
void* QQmlContext_MetaObjectDefault(void* ptr);
void QQmlContext_TimerEventDefault(void* ptr, void* event);
void* QQmlEngine_NewQQmlEngine(void* parent);
void QQmlEngine_AddImageProvider(void* ptr, struct QtQml_PackedString providerId, void* provider);
void QQmlEngine_AddImportPath(void* ptr, struct QtQml_PackedString path);
void QQmlEngine_AddPluginPath(void* ptr, struct QtQml_PackedString path);
void* QQmlEngine_BaseUrl(void* ptr);
void QQmlEngine_ClearComponentCache(void* ptr);
void* QQmlEngine_QQmlEngine_ContextForObject(void* object);
void QQmlEngine_ConnectExit(void* ptr, long long t);
void QQmlEngine_DisconnectExit(void* ptr);
void QQmlEngine_Exit(void* ptr, int retCode);
void* QQmlEngine_ImageProvider(void* ptr, struct QtQml_PackedString providerId);
struct QtQml_PackedString QQmlEngine_ImportPathList(void* ptr);
char QQmlEngine_ImportPlugin(void* ptr, struct QtQml_PackedString filePath, struct QtQml_PackedString uri, void* errors);
void* QQmlEngine_IncubationController(void* ptr);
void* QQmlEngine_NetworkAccessManager(void* ptr);
void* QQmlEngine_NetworkAccessManagerFactory(void* ptr);
long long QQmlEngine_QQmlEngine_ObjectOwnership(void* object);
struct QtQml_PackedString QQmlEngine_OfflineStorageDatabaseFilePath(void* ptr, struct QtQml_PackedString databaseName);
struct QtQml_PackedString QQmlEngine_OfflineStoragePath(void* ptr);
char QQmlEngine_OutputWarningsToStandardError(void* ptr);
struct QtQml_PackedString QQmlEngine_PluginPathList(void* ptr);
void QQmlEngine_ConnectQuit(void* ptr, long long t);
void QQmlEngine_DisconnectQuit(void* ptr);
void QQmlEngine_Quit(void* ptr);
void QQmlEngine_RemoveImageProvider(void* ptr, struct QtQml_PackedString providerId);
void QQmlEngine_Retranslate(void* ptr);
void QQmlEngine_RetranslateDefault(void* ptr);
void* QQmlEngine_RootContext(void* ptr);
void QQmlEngine_SetBaseUrl(void* ptr, void* url);
void QQmlEngine_QQmlEngine_SetContextForObject(void* object, void* context);
void QQmlEngine_SetImportPathList(void* ptr, struct QtQml_PackedString paths);
void QQmlEngine_SetIncubationController(void* ptr, void* controller);
void QQmlEngine_SetNetworkAccessManagerFactory(void* ptr, void* factory);
void QQmlEngine_QQmlEngine_SetObjectOwnership(void* object, long long ownership);
void QQmlEngine_SetOfflineStoragePath(void* ptr, struct QtQml_PackedString dir);
void QQmlEngine_SetOutputWarningsToStandardError(void* ptr, char enabled);
void QQmlEngine_SetPluginPathList(void* ptr, struct QtQml_PackedString paths);
void QQmlEngine_TrimComponentCache(void* ptr);
void QQmlEngine_ConnectWarnings(void* ptr, long long t);
void QQmlEngine_DisconnectWarnings(void* ptr);
void QQmlEngine_Warnings(void* ptr, void* warnings);
void QQmlEngine_DestroyQQmlEngine(void* ptr);
void QQmlEngine_DestroyQQmlEngineDefault(void* ptr);
int QQmlEngine_QQmlEngine_QmlRegisterSingletonType(void* url, char* uri, int versionMajor, int versionMinor, char* qmlName);
int QQmlEngine_QQmlEngine_QmlRegisterType(void* url, char* uri, int versionMajor, int versionMinor, char* qmlName);
int QQmlEngine_QQmlEngine_QmlRegisterSingletonInstance(char* uri, int versionMajor, int versionMinor, char* typeName, void* cppObject);
void* QQmlEngine___importPlugin_errors_atList(void* ptr, int i);
void QQmlEngine___importPlugin_errors_setList(void* ptr, void* i);
void* QQmlEngine___importPlugin_errors_newList(void* ptr);
void* QQmlEngine___warnings_warnings_atList(void* ptr, int i);
void QQmlEngine___warnings_warnings_setList(void* ptr, void* i);
void* QQmlEngine___warnings_warnings_newList(void* ptr);
void* QQmlEngineExtensionPlugin_NewQQmlEngineExtensionPlugin3(void* parent);
void QQmlEngineExtensionPlugin_InitializeEngine(void* ptr, void* engine, char* uri);
void QQmlEngineExtensionPlugin_InitializeEngineDefault(void* ptr, void* engine, char* uri);
void* QQmlEngineExtensionPlugin___children_atList(void* ptr, int i);
void QQmlEngineExtensionPlugin___children_setList(void* ptr, void* i);
void* QQmlEngineExtensionPlugin___children_newList(void* ptr);
void* QQmlEngineExtensionPlugin___dynamicPropertyNames_atList(void* ptr, int i);
void QQmlEngineExtensionPlugin___dynamicPropertyNames_setList(void* ptr, void* i);
void* QQmlEngineExtensionPlugin___dynamicPropertyNames_newList(void* ptr);
void* QQmlEngineExtensionPlugin___findChildren_atList(void* ptr, int i);
void QQmlEngineExtensionPlugin___findChildren_setList(void* ptr, void* i);
void* QQmlEngineExtensionPlugin___findChildren_newList(void* ptr);
void* QQmlEngineExtensionPlugin___findChildren_atList3(void* ptr, int i);
void QQmlEngineExtensionPlugin___findChildren_setList3(void* ptr, void* i);
void* QQmlEngineExtensionPlugin___findChildren_newList3(void* ptr);
void QQmlEngineExtensionPlugin_ChildEventDefault(void* ptr, void* event);
void QQmlEngineExtensionPlugin_ConnectNotifyDefault(void* ptr, void* sign);
void QQmlEngineExtensionPlugin_CustomEventDefault(void* ptr, void* event);
void QQmlEngineExtensionPlugin_DeleteLaterDefault(void* ptr);
void QQmlEngineExtensionPlugin_DisconnectNotifyDefault(void* ptr, void* sign);
char QQmlEngineExtensionPlugin_EventDefault(void* ptr, void* e);
char QQmlEngineExtensionPlugin_EventFilterDefault(void* ptr, void* watched, void* event);
void* QQmlEngineExtensionPlugin_MetaObjectDefault(void* ptr);
void QQmlEngineExtensionPlugin_TimerEventDefault(void* ptr, void* event);
void* QQmlError_NewQQmlError();
void* QQmlError_NewQQmlError2(void* other);
int QQmlError_Column(void* ptr);
struct QtQml_PackedString QQmlError_Description(void* ptr);
char QQmlError_IsValid(void* ptr);
int QQmlError_Line(void* ptr);
void* QQmlError_Object(void* ptr);
void QQmlError_SetColumn(void* ptr, int column);
void QQmlError_SetDescription(void* ptr, struct QtQml_PackedString description);
void QQmlError_SetLine(void* ptr, int line);
void QQmlError_SetObject(void* ptr, void* object);
void QQmlError_SetUrl(void* ptr, void* url);
struct QtQml_PackedString QQmlError_ToString(void* ptr);
void* QQmlError_Url(void* ptr);
void* QQmlExpression_NewQQmlExpression();
void* QQmlExpression_NewQQmlExpression2(void* ctxt, void* scope, struct QtQml_PackedString expression, void* parent);
void* QQmlExpression_NewQQmlExpression3(void* scri, void* ctxt, void* scope, void* parent);
void QQmlExpression_ClearError(void* ptr);
int QQmlExpression_ColumnNumber(void* ptr);
void* QQmlExpression_Context(void* ptr);
void* QQmlExpression_Engine(void* ptr);
void* QQmlExpression_Error(void* ptr);
void* QQmlExpression_Evaluate(void* ptr, char* valueIsUndefined);
struct QtQml_PackedString QQmlExpression_Expression(void* ptr);
char QQmlExpression_HasError(void* ptr);
int QQmlExpression_LineNumber(void* ptr);
char QQmlExpression_NotifyOnValueChanged(void* ptr);
void* QQmlExpression_ScopeObject(void* ptr);
void QQmlExpression_SetExpression(void* ptr, struct QtQml_PackedString expression);
void QQmlExpression_SetNotifyOnValueChanged(void* ptr, char notifyOnChange);
void QQmlExpression_SetSourceLocation(void* ptr, struct QtQml_PackedString url, int line, int column);
struct QtQml_PackedString QQmlExpression_SourceFile(void* ptr);
void QQmlExpression_ConnectValueChanged(void* ptr, long long t);
void QQmlExpression_DisconnectValueChanged(void* ptr);
void QQmlExpression_ValueChanged(void* ptr);
void QQmlExpression_DestroyQQmlExpression(void* ptr);
void QQmlExpression_DestroyQQmlExpressionDefault(void* ptr);
void* QQmlExpression___children_atList(void* ptr, int i);
void QQmlExpression___children_setList(void* ptr, void* i);
void* QQmlExpression___children_newList(void* ptr);
void* QQmlExpression___dynamicPropertyNames_atList(void* ptr, int i);
void QQmlExpression___dynamicPropertyNames_setList(void* ptr, void* i);
void* QQmlExpression___dynamicPropertyNames_newList(void* ptr);
void* QQmlExpression___findChildren_atList(void* ptr, int i);
void QQmlExpression___findChildren_setList(void* ptr, void* i);
void* QQmlExpression___findChildren_newList(void* ptr);
void* QQmlExpression___findChildren_atList3(void* ptr, int i);
void QQmlExpression___findChildren_setList3(void* ptr, void* i);
void* QQmlExpression___findChildren_newList3(void* ptr);
void QQmlExpression_ChildEventDefault(void* ptr, void* event);
void QQmlExpression_ConnectNotifyDefault(void* ptr, void* sign);
void QQmlExpression_CustomEventDefault(void* ptr, void* event);
void QQmlExpression_DeleteLaterDefault(void* ptr);
void QQmlExpression_DisconnectNotifyDefault(void* ptr, void* sign);
char QQmlExpression_EventDefault(void* ptr, void* e);
char QQmlExpression_EventFilterDefault(void* ptr, void* watched, void* event);
void* QQmlExpression_MetaObjectDefault(void* ptr);
void QQmlExpression_TimerEventDefault(void* ptr, void* event);
void* QQmlFileSelector_NewQQmlFileSelector(void* engine, void* parent);
void* QQmlFileSelector_QQmlFileSelector_Get(void* engine);
void* QQmlFileSelector_Selector(void* ptr);
void QQmlFileSelector_SetExtraSelectors(void* ptr, struct QtQml_PackedString strin);
void QQmlFileSelector_SetExtraSelectors2(void* ptr, struct QtQml_PackedString strin);
void QQmlFileSelector_SetSelector(void* ptr, void* selector);
void QQmlFileSelector_DestroyQQmlFileSelector(void* ptr);
void QQmlFileSelector_DestroyQQmlFileSelectorDefault(void* ptr);
void* QQmlFileSelector___children_atList(void* ptr, int i);
void QQmlFileSelector___children_setList(void* ptr, void* i);
void* QQmlFileSelector___children_newList(void* ptr);
void* QQmlFileSelector___dynamicPropertyNames_atList(void* ptr, int i);
void QQmlFileSelector___dynamicPropertyNames_setList(void* ptr, void* i);
void* QQmlFileSelector___dynamicPropertyNames_newList(void* ptr);
void* QQmlFileSelector___findChildren_atList(void* ptr, int i);
void QQmlFileSelector___findChildren_setList(void* ptr, void* i);
void* QQmlFileSelector___findChildren_newList(void* ptr);
void* QQmlFileSelector___findChildren_atList3(void* ptr, int i);
void QQmlFileSelector___findChildren_setList3(void* ptr, void* i);
void* QQmlFileSelector___findChildren_newList3(void* ptr);
void QQmlFileSelector_ChildEventDefault(void* ptr, void* event);
void QQmlFileSelector_ConnectNotifyDefault(void* ptr, void* sign);
void QQmlFileSelector_CustomEventDefault(void* ptr, void* event);
void QQmlFileSelector_DeleteLaterDefault(void* ptr);
void QQmlFileSelector_DisconnectNotifyDefault(void* ptr, void* sign);
char QQmlFileSelector_EventDefault(void* ptr, void* e);
char QQmlFileSelector_EventFilterDefault(void* ptr, void* watched, void* event);
void* QQmlFileSelector_MetaObjectDefault(void* ptr);
void QQmlFileSelector_TimerEventDefault(void* ptr, void* event);
long long QQmlImageProviderBase_Flags(void* ptr);
long long QQmlImageProviderBase_ImageType(void* ptr);
void* QQmlIncubationController_NewQQmlIncubationController2();
void* QQmlIncubationController_Engine(void* ptr);
void QQmlIncubationController_IncubateFor(void* ptr, int msecs);
int QQmlIncubationController_IncubatingObjectCount(void* ptr);
void QQmlIncubationController_IncubatingObjectCountChanged(void* ptr, int incubatingObjectCount);
void QQmlIncubationController_IncubatingObjectCountChangedDefault(void* ptr, int incubatingObjectCount);
void* QQmlIncubator_NewQQmlIncubator2(long long mode);
void QQmlIncubator_Clear(void* ptr);
struct QtQml_PackedList QQmlIncubator_Errors(void* ptr);
void QQmlIncubator_ForceCompletion(void* ptr);
long long QQmlIncubator_IncubationMode(void* ptr);
char QQmlIncubator_IsError(void* ptr);
char QQmlIncubator_IsLoading(void* ptr);
char QQmlIncubator_IsNull(void* ptr);
char QQmlIncubator_IsReady(void* ptr);
void* QQmlIncubator_Object(void* ptr);
void QQmlIncubator_SetInitialState(void* ptr, void* object);
void QQmlIncubator_SetInitialStateDefault(void* ptr, void* object);
long long QQmlIncubator_Status(void* ptr);
void QQmlIncubator_StatusChanged(void* ptr, long long status);
void QQmlIncubator_StatusChangedDefault(void* ptr, long long status);
void* QQmlIncubator___errors_atList(void* ptr, int i);
void QQmlIncubator___errors_setList(void* ptr, void* i);
void* QQmlIncubator___errors_newList(void* ptr);
void* QQmlIncubator___setInitialProperties_initialProperties_atList(void* ptr, struct QtQml_PackedString v, int i);
void QQmlIncubator___setInitialProperties_initialProperties_setList(void* ptr, struct QtQml_PackedString key, void* i);
void* QQmlIncubator___setInitialProperties_initialProperties_newList(void* ptr);
struct QtQml_PackedList QQmlIncubator___setInitialProperties_initialProperties_keyList(void* ptr);
struct QtQml_PackedString QQmlIncubator_____setInitialProperties_initialProperties_keyList_atList(void* ptr, int i);
void QQmlIncubator_____setInitialProperties_initialProperties_keyList_setList(void* ptr, struct QtQml_PackedString i);
void* QQmlIncubator_____setInitialProperties_initialProperties_keyList_newList(void* ptr);
void* QQmlListReference_NewQQmlListReference();
void* QQmlListReference_NewQQmlListReference2(void* object, char* property, void* engine);
char QQmlListReference_Append(void* ptr, void* object);
void* QQmlListReference_At(void* ptr, int index);
char QQmlListReference_CanAppend(void* ptr);
char QQmlListReference_CanAt(void* ptr);
char QQmlListReference_CanClear(void* ptr);
char QQmlListReference_CanCount(void* ptr);
char QQmlListReference_CanRemoveLast(void* ptr);
char QQmlListReference_CanReplace(void* ptr);
char QQmlListReference_Clear(void* ptr);
int QQmlListReference_Count(void* ptr);
char QQmlListReference_IsManipulable(void* ptr);
char QQmlListReference_IsReadable(void* ptr);
char QQmlListReference_IsValid(void* ptr);
void* QQmlListReference_ListElementType(void* ptr);
void* QQmlListReference_Object(void* ptr);
char QQmlListReference_RemoveLast(void* ptr);
char QQmlListReference_Replace(void* ptr, int index, void* object);
void* QQmlNetworkAccessManagerFactory_Create(void* ptr, void* parent);
void QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(void* ptr);
void QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactoryDefault(void* ptr);
void* QQmlNetworkAccessManagerFactory_NewQQmlNetworkAccessManagerFactory();
void QQmlParserStatus_ClassBegin(void* ptr);
void QQmlParserStatus_ComponentComplete(void* ptr);
void* QQmlProperty_NewQQmlProperty();
void* QQmlProperty_NewQQmlProperty2(void* obj);
void* QQmlProperty_NewQQmlProperty3(void* obj, void* ctxt);
void* QQmlProperty_NewQQmlProperty4(void* obj, void* engine);
void* QQmlProperty_NewQQmlProperty5(void* obj, struct QtQml_PackedString name);
void* QQmlProperty_NewQQmlProperty6(void* obj, struct QtQml_PackedString name, void* ctxt);
void* QQmlProperty_NewQQmlProperty7(void* obj, struct QtQml_PackedString name, void* engine);
void* QQmlProperty_NewQQmlProperty8(void* other);
char QQmlProperty_ConnectNotifySignal(void* ptr, void* dest, char* slot);
char QQmlProperty_ConnectNotifySignal2(void* ptr, void* dest, int method);
char QQmlProperty_HasNotifySignal(void* ptr);
int QQmlProperty_Index(void* ptr);
char QQmlProperty_IsDesignable(void* ptr);
char QQmlProperty_IsProperty(void* ptr);
char QQmlProperty_IsResettable(void* ptr);
char QQmlProperty_IsSignalProperty(void* ptr);
char QQmlProperty_IsValid(void* ptr);
char QQmlProperty_IsWritable(void* ptr);
void* QQmlProperty_Method(void* ptr);
struct QtQml_PackedString QQmlProperty_Name(void* ptr);
char QQmlProperty_NeedsNotifySignal(void* ptr);
void* QQmlProperty_Object(void* ptr);
int QQmlProperty_PropertyType(void* ptr);
long long QQmlProperty_PropertyTypeCategory(void* ptr);
struct QtQml_PackedString QQmlProperty_PropertyTypeName(void* ptr);
void* QQmlProperty_Read(void* ptr);
void* QQmlProperty_QQmlProperty_Read2(void* object, struct QtQml_PackedString name);
void* QQmlProperty_QQmlProperty_Read3(void* object, struct QtQml_PackedString name, void* ctxt);
void* QQmlProperty_QQmlProperty_Read4(void* object, struct QtQml_PackedString name, void* engine);
char QQmlProperty_Reset(void* ptr);
long long QQmlProperty_Type(void* ptr);
char QQmlProperty_Write(void* ptr, void* value);
char QQmlProperty_QQmlProperty_Write2(void* object, struct QtQml_PackedString name, void* value);
char QQmlProperty_QQmlProperty_Write3(void* object, struct QtQml_PackedString name, void* value, void* ctxt);
char QQmlProperty_QQmlProperty_Write4(void* object, struct QtQml_PackedString name, void* value, void* engine);
void* QQmlPropertyMap_NewQQmlPropertyMap(void* parent);
void QQmlPropertyMap_Clear(void* ptr, struct QtQml_PackedString key);
char QQmlPropertyMap_Contains(void* ptr, struct QtQml_PackedString key);
int QQmlPropertyMap_Count(void* ptr);
void QQmlPropertyMap_Insert(void* ptr, struct QtQml_PackedString key, void* value);
char QQmlPropertyMap_IsEmpty(void* ptr);
struct QtQml_PackedString QQmlPropertyMap_Keys(void* ptr);
int QQmlPropertyMap_Size(void* ptr);
void* QQmlPropertyMap_UpdateValue(void* ptr, struct QtQml_PackedString key, void* input);
void* QQmlPropertyMap_UpdateValueDefault(void* ptr, struct QtQml_PackedString key, void* input);
void* QQmlPropertyMap_Value(void* ptr, struct QtQml_PackedString key);
void QQmlPropertyMap_ConnectValueChanged(void* ptr, long long t);
void QQmlPropertyMap_DisconnectValueChanged(void* ptr);
void QQmlPropertyMap_ValueChanged(void* ptr, struct QtQml_PackedString key, void* value);
void QQmlPropertyMap_DestroyQQmlPropertyMap(void* ptr);
void QQmlPropertyMap_DestroyQQmlPropertyMapDefault(void* ptr);
void* QQmlPropertyMap___children_atList(void* ptr, int i);
void QQmlPropertyMap___children_setList(void* ptr, void* i);
void* QQmlPropertyMap___children_newList(void* ptr);
void* QQmlPropertyMap___dynamicPropertyNames_atList(void* ptr, int i);
void QQmlPropertyMap___dynamicPropertyNames_setList(void* ptr, void* i);
void* QQmlPropertyMap___dynamicPropertyNames_newList(void* ptr);
void* QQmlPropertyMap___findChildren_atList(void* ptr, int i);
void QQmlPropertyMap___findChildren_setList(void* ptr, void* i);
void* QQmlPropertyMap___findChildren_newList(void* ptr);
void* QQmlPropertyMap___findChildren_atList3(void* ptr, int i);
void QQmlPropertyMap___findChildren_setList3(void* ptr, void* i);
void* QQmlPropertyMap___findChildren_newList3(void* ptr);
void QQmlPropertyMap_ChildEventDefault(void* ptr, void* event);
void QQmlPropertyMap_ConnectNotifyDefault(void* ptr, void* sign);
void QQmlPropertyMap_CustomEventDefault(void* ptr, void* event);
void QQmlPropertyMap_DeleteLaterDefault(void* ptr);
void QQmlPropertyMap_DisconnectNotifyDefault(void* ptr, void* sign);
char QQmlPropertyMap_EventDefault(void* ptr, void* e);
char QQmlPropertyMap_EventFilterDefault(void* ptr, void* watched, void* event);
void* QQmlPropertyMap_MetaObjectDefault(void* ptr);
void QQmlPropertyMap_TimerEventDefault(void* ptr, void* event);
void* QQmlPropertyValueSource_NewQQmlPropertyValueSource();
void QQmlPropertyValueSource_SetTarget(void* ptr, void* property);
void QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(void* ptr);
void QQmlPropertyValueSource_DestroyQQmlPropertyValueSourceDefault(void* ptr);
void* QQmlScriptString_NewQQmlScriptString();
void* QQmlScriptString_NewQQmlScriptString2(void* other);
char QQmlScriptString_BooleanLiteral(void* ptr, char* ok);
char QQmlScriptString_IsEmpty(void* ptr);
char QQmlScriptString_IsNullLiteral(void* ptr);
char QQmlScriptString_IsUndefinedLiteral(void* ptr);
double QQmlScriptString_NumberLiteral(void* ptr, char* ok);
struct QtQml_PackedString QQmlScriptString_StringLiteral(void* ptr);

#ifdef __cplusplus
}
#endif

#endif