// +build !minimal

#define protected public
#define private public

#include "xml.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChar>
#include <QDomAttr>
#include <QDomCDATASection>
#include <QDomCharacterData>
#include <QDomComment>
#include <QDomDocument>
#include <QDomDocumentFragment>
#include <QDomDocumentType>
#include <QDomElement>
#include <QDomEntity>
#include <QDomEntityReference>
#include <QDomImplementation>
#include <QDomNamedNodeMap>
#include <QDomNode>
#include <QDomNodeList>
#include <QDomNotation>
#include <QDomProcessingInstruction>
#include <QDomText>
#include <QIODevice>
#include <QLatin1String>
#include <QString>
#include <QTextStream>
#include <QXmlAttributes>
#include <QXmlContentHandler>
#include <QXmlDTDHandler>
#include <QXmlDeclHandler>
#include <QXmlDefaultHandler>
#include <QXmlEntityResolver>
#include <QXmlErrorHandler>
#include <QXmlInputSource>
#include <QXmlLexicalHandler>
#include <QXmlLocator>
#include <QXmlNamespaceSupport>
#include <QXmlParseException>
#include <QXmlReader>
#include <QXmlSimpleReader>
#include <QXmlStreamReader>

Q_DECLARE_METATYPE(QDomAttr)
Q_DECLARE_METATYPE(QDomAttr*)
void* QDomAttr_NewQDomAttr()
{
	return new QDomAttr();
}

void* QDomAttr_NewQDomAttr2(void* x)
{
	return new QDomAttr(*static_cast<QDomAttr*>(x));
}

struct QtXml_PackedString QDomAttr_Name(void* ptr)
{
	return ({ QByteArray* t1fdad4 = new QByteArray(static_cast<QDomAttr*>(ptr)->name().toUtf8()); QtXml_PackedString { const_cast<char*>(t1fdad4->prepend("WHITESPACE").constData()+10), t1fdad4->size()-10, t1fdad4 }; });
}

void* QDomAttr_OwnerElement(void* ptr)
{
	return new QDomElement(static_cast<QDomAttr*>(ptr)->ownerElement());
}

void QDomAttr_SetValue(void* ptr, struct QtXml_PackedString v)
{
	static_cast<QDomAttr*>(ptr)->setValue(QString::fromUtf8(v.data, v.len));
}

char QDomAttr_Specified(void* ptr)
{
	return static_cast<QDomAttr*>(ptr)->specified();
}

struct QtXml_PackedString QDomAttr_Value(void* ptr)
{
	return ({ QByteArray* tcc8e31 = new QByteArray(static_cast<QDomAttr*>(ptr)->value().toUtf8()); QtXml_PackedString { const_cast<char*>(tcc8e31->prepend("WHITESPACE").constData()+10), tcc8e31->size()-10, tcc8e31 }; });
}

Q_DECLARE_METATYPE(QDomCDATASection)
Q_DECLARE_METATYPE(QDomCDATASection*)
void* QDomCDATASection_NewQDomCDATASection()
{
	return new QDomCDATASection();
}

void* QDomCDATASection_NewQDomCDATASection2(void* x)
{
	return new QDomCDATASection(*static_cast<QDomCDATASection*>(x));
}

Q_DECLARE_METATYPE(QDomCharacterData)
Q_DECLARE_METATYPE(QDomCharacterData*)
void* QDomCharacterData_NewQDomCharacterData()
{
	return new QDomCharacterData();
}

void* QDomCharacterData_NewQDomCharacterData2(void* x)
{
	return new QDomCharacterData(*static_cast<QDomCharacterData*>(x));
}

void QDomCharacterData_AppendData(void* ptr, struct QtXml_PackedString arg)
{
	static_cast<QDomCharacterData*>(ptr)->appendData(QString::fromUtf8(arg.data, arg.len));
}

struct QtXml_PackedString QDomCharacterData_Data(void* ptr)
{
	return ({ QByteArray* t6bdb28 = new QByteArray(static_cast<QDomCharacterData*>(ptr)->data().toUtf8()); QtXml_PackedString { const_cast<char*>(t6bdb28->prepend("WHITESPACE").constData()+10), t6bdb28->size()-10, t6bdb28 }; });
}

void QDomCharacterData_DeleteData(void* ptr, unsigned long offset, unsigned long count)
{
	static_cast<QDomCharacterData*>(ptr)->deleteData(offset, count);
}

void QDomCharacterData_InsertData(void* ptr, unsigned long offset, struct QtXml_PackedString arg)
{
	static_cast<QDomCharacterData*>(ptr)->insertData(offset, QString::fromUtf8(arg.data, arg.len));
}

int QDomCharacterData_Length(void* ptr)
{
	return static_cast<QDomCharacterData*>(ptr)->length();
}

void QDomCharacterData_ReplaceData(void* ptr, unsigned long offset, unsigned long count, struct QtXml_PackedString arg)
{
	static_cast<QDomCharacterData*>(ptr)->replaceData(offset, count, QString::fromUtf8(arg.data, arg.len));
}

void QDomCharacterData_SetData(void* ptr, struct QtXml_PackedString v)
{
	static_cast<QDomCharacterData*>(ptr)->setData(QString::fromUtf8(v.data, v.len));
}

struct QtXml_PackedString QDomCharacterData_SubstringData(void* ptr, unsigned long offset, unsigned long count)
{
	return ({ QByteArray* tbe0a67 = new QByteArray(static_cast<QDomCharacterData*>(ptr)->substringData(offset, count).toUtf8()); QtXml_PackedString { const_cast<char*>(tbe0a67->prepend("WHITESPACE").constData()+10), tbe0a67->size()-10, tbe0a67 }; });
}

Q_DECLARE_METATYPE(QDomComment)
Q_DECLARE_METATYPE(QDomComment*)
void* QDomComment_NewQDomComment()
{
	return new QDomComment();
}

void* QDomComment_NewQDomComment2(void* x)
{
	return new QDomComment(*static_cast<QDomComment*>(x));
}

Q_DECLARE_METATYPE(QDomDocument)
Q_DECLARE_METATYPE(QDomDocument*)
void* QDomDocument_NewQDomDocument()
{
	return new QDomDocument();
}

void* QDomDocument_NewQDomDocument2(struct QtXml_PackedString name)
{
	return new QDomDocument(QString::fromUtf8(name.data, name.len));
}

void* QDomDocument_NewQDomDocument3(void* doctype)
{
	return new QDomDocument(*static_cast<QDomDocumentType*>(doctype));
}

void* QDomDocument_NewQDomDocument4(void* x)
{
	return new QDomDocument(*static_cast<QDomDocument*>(x));
}

void* QDomDocument_CreateAttribute(void* ptr, struct QtXml_PackedString name)
{
	return new QDomAttr(static_cast<QDomDocument*>(ptr)->createAttribute(QString::fromUtf8(name.data, name.len)));
}

void* QDomDocument_CreateAttributeNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString qName)
{
	return new QDomAttr(static_cast<QDomDocument*>(ptr)->createAttributeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(qName.data, qName.len)));
}

void* QDomDocument_CreateCDATASection(void* ptr, struct QtXml_PackedString value)
{
	return new QDomCDATASection(static_cast<QDomDocument*>(ptr)->createCDATASection(QString::fromUtf8(value.data, value.len)));
}

void* QDomDocument_CreateComment(void* ptr, struct QtXml_PackedString value)
{
	return new QDomComment(static_cast<QDomDocument*>(ptr)->createComment(QString::fromUtf8(value.data, value.len)));
}

void* QDomDocument_CreateDocumentFragment(void* ptr)
{
	return new QDomDocumentFragment(static_cast<QDomDocument*>(ptr)->createDocumentFragment());
}

void* QDomDocument_CreateElement(void* ptr, struct QtXml_PackedString tagName)
{
	return new QDomElement(static_cast<QDomDocument*>(ptr)->createElement(QString::fromUtf8(tagName.data, tagName.len)));
}

void* QDomDocument_CreateElementNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString qName)
{
	return new QDomElement(static_cast<QDomDocument*>(ptr)->createElementNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(qName.data, qName.len)));
}

void* QDomDocument_CreateEntityReference(void* ptr, struct QtXml_PackedString name)
{
	return new QDomEntityReference(static_cast<QDomDocument*>(ptr)->createEntityReference(QString::fromUtf8(name.data, name.len)));
}

void* QDomDocument_CreateProcessingInstruction(void* ptr, struct QtXml_PackedString target, struct QtXml_PackedString data)
{
	return new QDomProcessingInstruction(static_cast<QDomDocument*>(ptr)->createProcessingInstruction(QString::fromUtf8(target.data, target.len), QString::fromUtf8(data.data, data.len)));
}

void* QDomDocument_CreateTextNode(void* ptr, struct QtXml_PackedString value)
{
	return new QDomText(static_cast<QDomDocument*>(ptr)->createTextNode(QString::fromUtf8(value.data, value.len)));
}

void* QDomDocument_Doctype(void* ptr)
{
	return new QDomDocumentType(static_cast<QDomDocument*>(ptr)->doctype());
}

void* QDomDocument_DocumentElement(void* ptr)
{
	return new QDomElement(static_cast<QDomDocument*>(ptr)->documentElement());
}

void* QDomDocument_ElementById(void* ptr, struct QtXml_PackedString elementId)
{
	return new QDomElement(static_cast<QDomDocument*>(ptr)->elementById(QString::fromUtf8(elementId.data, elementId.len)));
}

void* QDomDocument_ElementsByTagName(void* ptr, struct QtXml_PackedString tagname)
{
	return new QDomNodeList(static_cast<QDomDocument*>(ptr)->elementsByTagName(QString::fromUtf8(tagname.data, tagname.len)));
}

void* QDomDocument_ElementsByTagNameNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString localName)
{
	return new QDomNodeList(static_cast<QDomDocument*>(ptr)->elementsByTagNameNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(localName.data, localName.len)));
}

void* QDomDocument_Implementation(void* ptr)
{
	return new QDomImplementation(static_cast<QDomDocument*>(ptr)->implementation());
}

void* QDomDocument_ImportNode(void* ptr, void* importedNode, char deep)
{
	return new QDomNode(static_cast<QDomDocument*>(ptr)->importNode(*static_cast<QDomNode*>(importedNode), deep != 0));
}

char QDomDocument_SetContent(void* ptr, void* data, char namespaceProcessing, struct QtXml_PackedString errorMsg, int errorLine, int errorColumn)
{
	return static_cast<QDomDocument*>(ptr)->setContent(*static_cast<QByteArray*>(data), namespaceProcessing != 0, new QString(QString::fromUtf8(errorMsg.data, errorMsg.len)), &errorLine, &errorColumn);
}

char QDomDocument_SetContent2(void* ptr, struct QtXml_PackedString text, char namespaceProcessing, struct QtXml_PackedString errorMsg, int errorLine, int errorColumn)
{
	return static_cast<QDomDocument*>(ptr)->setContent(QString::fromUtf8(text.data, text.len), namespaceProcessing != 0, new QString(QString::fromUtf8(errorMsg.data, errorMsg.len)), &errorLine, &errorColumn);
}

char QDomDocument_SetContent3(void* ptr, void* dev, char namespaceProcessing, struct QtXml_PackedString errorMsg, int errorLine, int errorColumn)
{
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QIODevice*>(dev), namespaceProcessing != 0, new QString(QString::fromUtf8(errorMsg.data, errorMsg.len)), &errorLine, &errorColumn);
}

char QDomDocument_SetContent5(void* ptr, void* buffer, struct QtXml_PackedString errorMsg, int errorLine, int errorColumn)
{
	return static_cast<QDomDocument*>(ptr)->setContent(*static_cast<QByteArray*>(buffer), new QString(QString::fromUtf8(errorMsg.data, errorMsg.len)), &errorLine, &errorColumn);
}

char QDomDocument_SetContent6(void* ptr, struct QtXml_PackedString text, struct QtXml_PackedString errorMsg, int errorLine, int errorColumn)
{
	return static_cast<QDomDocument*>(ptr)->setContent(QString::fromUtf8(text.data, text.len), new QString(QString::fromUtf8(errorMsg.data, errorMsg.len)), &errorLine, &errorColumn);
}

char QDomDocument_SetContent9(void* ptr, void* reader, char namespaceProcessing, struct QtXml_PackedString errorMsg, int errorLine, int errorColumn)
{
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QXmlStreamReader*>(reader), namespaceProcessing != 0, new QString(QString::fromUtf8(errorMsg.data, errorMsg.len)), &errorLine, &errorColumn);
}

void* QDomDocument_ToByteArray(void* ptr, int indent)
{
	return new QByteArray(static_cast<QDomDocument*>(ptr)->toByteArray(indent));
}

struct QtXml_PackedString QDomDocument_ToString(void* ptr, int indent)
{
	return ({ QByteArray* t6ed640 = new QByteArray(static_cast<QDomDocument*>(ptr)->toString(indent).toUtf8()); QtXml_PackedString { const_cast<char*>(t6ed640->prepend("WHITESPACE").constData()+10), t6ed640->size()-10, t6ed640 }; });
}

void QDomDocument_DestroyQDomDocument(void* ptr)
{
	static_cast<QDomDocument*>(ptr)->~QDomDocument();
}

Q_DECLARE_METATYPE(QDomDocumentFragment)
Q_DECLARE_METATYPE(QDomDocumentFragment*)
void* QDomDocumentFragment_NewQDomDocumentFragment()
{
	return new QDomDocumentFragment();
}

void* QDomDocumentFragment_NewQDomDocumentFragment2(void* x)
{
	return new QDomDocumentFragment(*static_cast<QDomDocumentFragment*>(x));
}

Q_DECLARE_METATYPE(QDomDocumentType)
Q_DECLARE_METATYPE(QDomDocumentType*)
void* QDomDocumentType_NewQDomDocumentType()
{
	return new QDomDocumentType();
}

void* QDomDocumentType_NewQDomDocumentType2(void* n)
{
	return new QDomDocumentType(*static_cast<QDomDocumentType*>(n));
}

void* QDomDocumentType_Entities(void* ptr)
{
	return new QDomNamedNodeMap(static_cast<QDomDocumentType*>(ptr)->entities());
}

struct QtXml_PackedString QDomDocumentType_InternalSubset(void* ptr)
{
	return ({ QByteArray* t955e99 = new QByteArray(static_cast<QDomDocumentType*>(ptr)->internalSubset().toUtf8()); QtXml_PackedString { const_cast<char*>(t955e99->prepend("WHITESPACE").constData()+10), t955e99->size()-10, t955e99 }; });
}

struct QtXml_PackedString QDomDocumentType_Name(void* ptr)
{
	return ({ QByteArray* tfa1c67 = new QByteArray(static_cast<QDomDocumentType*>(ptr)->name().toUtf8()); QtXml_PackedString { const_cast<char*>(tfa1c67->prepend("WHITESPACE").constData()+10), tfa1c67->size()-10, tfa1c67 }; });
}

void* QDomDocumentType_Notations(void* ptr)
{
	return new QDomNamedNodeMap(static_cast<QDomDocumentType*>(ptr)->notations());
}

struct QtXml_PackedString QDomDocumentType_PublicId(void* ptr)
{
	return ({ QByteArray* t0e356e = new QByteArray(static_cast<QDomDocumentType*>(ptr)->publicId().toUtf8()); QtXml_PackedString { const_cast<char*>(t0e356e->prepend("WHITESPACE").constData()+10), t0e356e->size()-10, t0e356e }; });
}

struct QtXml_PackedString QDomDocumentType_SystemId(void* ptr)
{
	return ({ QByteArray* tf7dc95 = new QByteArray(static_cast<QDomDocumentType*>(ptr)->systemId().toUtf8()); QtXml_PackedString { const_cast<char*>(tf7dc95->prepend("WHITESPACE").constData()+10), tf7dc95->size()-10, tf7dc95 }; });
}

Q_DECLARE_METATYPE(QDomElement)
Q_DECLARE_METATYPE(QDomElement*)
void* QDomElement_NewQDomElement()
{
	return new QDomElement();
}

void* QDomElement_NewQDomElement2(void* x)
{
	return new QDomElement(*static_cast<QDomElement*>(x));
}

struct QtXml_PackedString QDomElement_Attribute(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString defValue)
{
	return ({ QByteArray* t268480 = new QByteArray(static_cast<QDomElement*>(ptr)->attribute(QString::fromUtf8(name.data, name.len), QString::fromUtf8(defValue.data, defValue.len)).toUtf8()); QtXml_PackedString { const_cast<char*>(t268480->prepend("WHITESPACE").constData()+10), t268480->size()-10, t268480 }; });
}

struct QtXml_PackedString QDomElement_AttributeNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString localName, struct QtXml_PackedString defValue)
{
	return ({ QByteArray* t495754 = new QByteArray(static_cast<QDomElement*>(ptr)->attributeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(localName.data, localName.len), QString::fromUtf8(defValue.data, defValue.len)).toUtf8()); QtXml_PackedString { const_cast<char*>(t495754->prepend("WHITESPACE").constData()+10), t495754->size()-10, t495754 }; });
}

void* QDomElement_AttributeNode(void* ptr, struct QtXml_PackedString name)
{
	return new QDomAttr(static_cast<QDomElement*>(ptr)->attributeNode(QString::fromUtf8(name.data, name.len)));
}

void* QDomElement_AttributeNodeNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString localName)
{
	return new QDomAttr(static_cast<QDomElement*>(ptr)->attributeNodeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(localName.data, localName.len)));
}

void* QDomElement_ElementsByTagName(void* ptr, struct QtXml_PackedString tagname)
{
	return new QDomNodeList(static_cast<QDomElement*>(ptr)->elementsByTagName(QString::fromUtf8(tagname.data, tagname.len)));
}

void* QDomElement_ElementsByTagNameNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString localName)
{
	return new QDomNodeList(static_cast<QDomElement*>(ptr)->elementsByTagNameNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(localName.data, localName.len)));
}

char QDomElement_HasAttribute(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QDomElement*>(ptr)->hasAttribute(QString::fromUtf8(name.data, name.len));
}

char QDomElement_HasAttributeNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString localName)
{
	return static_cast<QDomElement*>(ptr)->hasAttributeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(localName.data, localName.len));
}

void QDomElement_RemoveAttribute(void* ptr, struct QtXml_PackedString name)
{
	static_cast<QDomElement*>(ptr)->removeAttribute(QString::fromUtf8(name.data, name.len));
}

void QDomElement_RemoveAttributeNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString localName)
{
	static_cast<QDomElement*>(ptr)->removeAttributeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(localName.data, localName.len));
}

void* QDomElement_RemoveAttributeNode(void* ptr, void* oldAttr)
{
	return new QDomAttr(static_cast<QDomElement*>(ptr)->removeAttributeNode(*static_cast<QDomAttr*>(oldAttr)));
}

void QDomElement_SetAttribute(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString value)
{
	static_cast<QDomElement*>(ptr)->setAttribute(QString::fromUtf8(name.data, name.len), QString::fromUtf8(value.data, value.len));
}

void QDomElement_SetAttribute2(void* ptr, struct QtXml_PackedString name, long long value)
{
	static_cast<QDomElement*>(ptr)->setAttribute(QString::fromUtf8(name.data, name.len), value);
}

void QDomElement_SetAttribute3(void* ptr, struct QtXml_PackedString name, unsigned long long value)
{
	static_cast<QDomElement*>(ptr)->setAttribute(QString::fromUtf8(name.data, name.len), value);
}

void QDomElement_SetAttribute4(void* ptr, struct QtXml_PackedString name, int value)
{
	static_cast<QDomElement*>(ptr)->setAttribute(QString::fromUtf8(name.data, name.len), value);
}

void QDomElement_SetAttribute5(void* ptr, struct QtXml_PackedString name, unsigned int value)
{
	static_cast<QDomElement*>(ptr)->setAttribute(QString::fromUtf8(name.data, name.len), value);
}

void QDomElement_SetAttribute6(void* ptr, struct QtXml_PackedString name, float value)
{
	static_cast<QDomElement*>(ptr)->setAttribute(QString::fromUtf8(name.data, name.len), value);
}

void QDomElement_SetAttribute7(void* ptr, struct QtXml_PackedString name, double value)
{
	static_cast<QDomElement*>(ptr)->setAttribute(QString::fromUtf8(name.data, name.len), value);
}

void QDomElement_SetAttributeNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString qName, struct QtXml_PackedString value)
{
	static_cast<QDomElement*>(ptr)->setAttributeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(qName.data, qName.len), QString::fromUtf8(value.data, value.len));
}

void QDomElement_SetAttributeNS2(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString qName, int value)
{
	static_cast<QDomElement*>(ptr)->setAttributeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(qName.data, qName.len), value);
}

void QDomElement_SetAttributeNS3(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString qName, unsigned int value)
{
	static_cast<QDomElement*>(ptr)->setAttributeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(qName.data, qName.len), value);
}

void QDomElement_SetAttributeNS4(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString qName, long long value)
{
	static_cast<QDomElement*>(ptr)->setAttributeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(qName.data, qName.len), value);
}

void QDomElement_SetAttributeNS5(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString qName, unsigned long long value)
{
	static_cast<QDomElement*>(ptr)->setAttributeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(qName.data, qName.len), value);
}

void QDomElement_SetAttributeNS6(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString qName, double value)
{
	static_cast<QDomElement*>(ptr)->setAttributeNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(qName.data, qName.len), value);
}

void* QDomElement_SetAttributeNode(void* ptr, void* newAttr)
{
	return new QDomAttr(static_cast<QDomElement*>(ptr)->setAttributeNode(*static_cast<QDomAttr*>(newAttr)));
}

void* QDomElement_SetAttributeNodeNS(void* ptr, void* newAttr)
{
	return new QDomAttr(static_cast<QDomElement*>(ptr)->setAttributeNodeNS(*static_cast<QDomAttr*>(newAttr)));
}

void QDomElement_SetTagName(void* ptr, struct QtXml_PackedString name)
{
	static_cast<QDomElement*>(ptr)->setTagName(QString::fromUtf8(name.data, name.len));
}

struct QtXml_PackedString QDomElement_TagName(void* ptr)
{
	return ({ QByteArray* tf9a4f6 = new QByteArray(static_cast<QDomElement*>(ptr)->tagName().toUtf8()); QtXml_PackedString { const_cast<char*>(tf9a4f6->prepend("WHITESPACE").constData()+10), tf9a4f6->size()-10, tf9a4f6 }; });
}

struct QtXml_PackedString QDomElement_Text(void* ptr)
{
	return ({ QByteArray* t99336a = new QByteArray(static_cast<QDomElement*>(ptr)->text().toUtf8()); QtXml_PackedString { const_cast<char*>(t99336a->prepend("WHITESPACE").constData()+10), t99336a->size()-10, t99336a }; });
}

Q_DECLARE_METATYPE(QDomEntity)
Q_DECLARE_METATYPE(QDomEntity*)
void* QDomEntity_NewQDomEntity()
{
	return new QDomEntity();
}

void* QDomEntity_NewQDomEntity2(void* x)
{
	return new QDomEntity(*static_cast<QDomEntity*>(x));
}

struct QtXml_PackedString QDomEntity_NotationName(void* ptr)
{
	return ({ QByteArray* t33d72c = new QByteArray(static_cast<QDomEntity*>(ptr)->notationName().toUtf8()); QtXml_PackedString { const_cast<char*>(t33d72c->prepend("WHITESPACE").constData()+10), t33d72c->size()-10, t33d72c }; });
}

struct QtXml_PackedString QDomEntity_PublicId(void* ptr)
{
	return ({ QByteArray* t03d4e3 = new QByteArray(static_cast<QDomEntity*>(ptr)->publicId().toUtf8()); QtXml_PackedString { const_cast<char*>(t03d4e3->prepend("WHITESPACE").constData()+10), t03d4e3->size()-10, t03d4e3 }; });
}

struct QtXml_PackedString QDomEntity_SystemId(void* ptr)
{
	return ({ QByteArray* tb68351 = new QByteArray(static_cast<QDomEntity*>(ptr)->systemId().toUtf8()); QtXml_PackedString { const_cast<char*>(tb68351->prepend("WHITESPACE").constData()+10), tb68351->size()-10, tb68351 }; });
}

Q_DECLARE_METATYPE(QDomEntityReference)
Q_DECLARE_METATYPE(QDomEntityReference*)
void* QDomEntityReference_NewQDomEntityReference()
{
	return new QDomEntityReference();
}

void* QDomEntityReference_NewQDomEntityReference2(void* x)
{
	return new QDomEntityReference(*static_cast<QDomEntityReference*>(x));
}

Q_DECLARE_METATYPE(QDomImplementation)
Q_DECLARE_METATYPE(QDomImplementation*)
void* QDomImplementation_NewQDomImplementation()
{
	return new QDomImplementation();
}

void* QDomImplementation_NewQDomImplementation2(void* x)
{
	return new QDomImplementation(*static_cast<QDomImplementation*>(x));
}

void* QDomImplementation_CreateDocument(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString qName, void* doctype)
{
	return new QDomDocument(static_cast<QDomImplementation*>(ptr)->createDocument(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(qName.data, qName.len), *static_cast<QDomDocumentType*>(doctype)));
}

void* QDomImplementation_CreateDocumentType(void* ptr, struct QtXml_PackedString qName, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId)
{
	return new QDomDocumentType(static_cast<QDomImplementation*>(ptr)->createDocumentType(QString::fromUtf8(qName.data, qName.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len)));
}

char QDomImplementation_HasFeature(void* ptr, struct QtXml_PackedString feature, struct QtXml_PackedString version)
{
	return static_cast<QDomImplementation*>(ptr)->hasFeature(QString::fromUtf8(feature.data, feature.len), QString::fromUtf8(version.data, version.len));
}

long long QDomImplementation_QDomImplementation_InvalidDataPolicy()
{
	return QDomImplementation::invalidDataPolicy();
}

char QDomImplementation_IsNull(void* ptr)
{
	return static_cast<QDomImplementation*>(ptr)->isNull();
}

void QDomImplementation_QDomImplementation_SetInvalidDataPolicy(long long policy)
{
	QDomImplementation::setInvalidDataPolicy(static_cast<QDomImplementation::InvalidDataPolicy>(policy));
}

void QDomImplementation_DestroyQDomImplementation(void* ptr)
{
	static_cast<QDomImplementation*>(ptr)->~QDomImplementation();
}

Q_DECLARE_METATYPE(QDomNamedNodeMap)
Q_DECLARE_METATYPE(QDomNamedNodeMap*)
void* QDomNamedNodeMap_NewQDomNamedNodeMap()
{
	return new QDomNamedNodeMap();
}

void* QDomNamedNodeMap_NewQDomNamedNodeMap2(void* n)
{
	return new QDomNamedNodeMap(*static_cast<QDomNamedNodeMap*>(n));
}

char QDomNamedNodeMap_Contains(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QDomNamedNodeMap*>(ptr)->contains(QString::fromUtf8(name.data, name.len));
}

int QDomNamedNodeMap_Count(void* ptr)
{
	return static_cast<QDomNamedNodeMap*>(ptr)->count();
}

char QDomNamedNodeMap_IsEmpty(void* ptr)
{
	return static_cast<QDomNamedNodeMap*>(ptr)->isEmpty();
}

void* QDomNamedNodeMap_Item(void* ptr, int index)
{
	return new QDomNode(static_cast<QDomNamedNodeMap*>(ptr)->item(index));
}

int QDomNamedNodeMap_Length(void* ptr)
{
	return static_cast<QDomNamedNodeMap*>(ptr)->length();
}

void* QDomNamedNodeMap_NamedItem(void* ptr, struct QtXml_PackedString name)
{
	return new QDomNode(static_cast<QDomNamedNodeMap*>(ptr)->namedItem(QString::fromUtf8(name.data, name.len)));
}

void* QDomNamedNodeMap_NamedItemNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString localName)
{
	return new QDomNode(static_cast<QDomNamedNodeMap*>(ptr)->namedItemNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(localName.data, localName.len)));
}

void* QDomNamedNodeMap_RemoveNamedItem(void* ptr, struct QtXml_PackedString name)
{
	return new QDomNode(static_cast<QDomNamedNodeMap*>(ptr)->removeNamedItem(QString::fromUtf8(name.data, name.len)));
}

void* QDomNamedNodeMap_RemoveNamedItemNS(void* ptr, struct QtXml_PackedString nsURI, struct QtXml_PackedString localName)
{
	return new QDomNode(static_cast<QDomNamedNodeMap*>(ptr)->removeNamedItemNS(QString::fromUtf8(nsURI.data, nsURI.len), QString::fromUtf8(localName.data, localName.len)));
}

void* QDomNamedNodeMap_SetNamedItem(void* ptr, void* newNode)
{
	return new QDomNode(static_cast<QDomNamedNodeMap*>(ptr)->setNamedItem(*static_cast<QDomNode*>(newNode)));
}

void* QDomNamedNodeMap_SetNamedItemNS(void* ptr, void* newNode)
{
	return new QDomNode(static_cast<QDomNamedNodeMap*>(ptr)->setNamedItemNS(*static_cast<QDomNode*>(newNode)));
}

int QDomNamedNodeMap_Size(void* ptr)
{
	return static_cast<QDomNamedNodeMap*>(ptr)->size();
}

void QDomNamedNodeMap_DestroyQDomNamedNodeMap(void* ptr)
{
	static_cast<QDomNamedNodeMap*>(ptr)->~QDomNamedNodeMap();
}

Q_DECLARE_METATYPE(QDomNode)
Q_DECLARE_METATYPE(QDomNode*)
void* QDomNode_NewQDomNode()
{
	return new QDomNode();
}

void* QDomNode_NewQDomNode2(void* n)
{
	return new QDomNode(*static_cast<QDomNode*>(n));
}

void* QDomNode_AppendChild(void* ptr, void* newChild)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->appendChild(*static_cast<QDomNode*>(newChild)));
}

void* QDomNode_ChildNodes(void* ptr)
{
	return new QDomNodeList(static_cast<QDomNode*>(ptr)->childNodes());
}

void QDomNode_Clear(void* ptr)
{
	static_cast<QDomNode*>(ptr)->clear();
}

void* QDomNode_CloneNode(void* ptr, char deep)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->cloneNode(deep != 0));
}

int QDomNode_ColumnNumber(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->columnNumber();
}

void* QDomNode_FirstChild(void* ptr)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->firstChild());
}

void* QDomNode_FirstChildElement(void* ptr, struct QtXml_PackedString tagName)
{
	return new QDomElement(static_cast<QDomNode*>(ptr)->firstChildElement(QString::fromUtf8(tagName.data, tagName.len)));
}

char QDomNode_HasAttributes(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->hasAttributes();
}

char QDomNode_HasChildNodes(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->hasChildNodes();
}

void* QDomNode_InsertAfter(void* ptr, void* newChild, void* refChild)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->insertAfter(*static_cast<QDomNode*>(newChild), *static_cast<QDomNode*>(refChild)));
}

void* QDomNode_InsertBefore(void* ptr, void* newChild, void* refChild)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->insertBefore(*static_cast<QDomNode*>(newChild), *static_cast<QDomNode*>(refChild)));
}

char QDomNode_IsAttr(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isAttr();
}

char QDomNode_IsCDATASection(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isCDATASection();
}

char QDomNode_IsCharacterData(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isCharacterData();
}

char QDomNode_IsComment(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isComment();
}

char QDomNode_IsDocument(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isDocument();
}

char QDomNode_IsDocumentFragment(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isDocumentFragment();
}

char QDomNode_IsDocumentType(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isDocumentType();
}

char QDomNode_IsElement(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isElement();
}

char QDomNode_IsEntity(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isEntity();
}

char QDomNode_IsEntityReference(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isEntityReference();
}

char QDomNode_IsNotation(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isNotation();
}

char QDomNode_IsNull(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isNull();
}

char QDomNode_IsProcessingInstruction(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isProcessingInstruction();
}

char QDomNode_IsSupported(void* ptr, struct QtXml_PackedString feature, struct QtXml_PackedString version)
{
	return static_cast<QDomNode*>(ptr)->isSupported(QString::fromUtf8(feature.data, feature.len), QString::fromUtf8(version.data, version.len));
}

char QDomNode_IsText(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->isText();
}

void* QDomNode_LastChild(void* ptr)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->lastChild());
}

void* QDomNode_LastChildElement(void* ptr, struct QtXml_PackedString tagName)
{
	return new QDomElement(static_cast<QDomNode*>(ptr)->lastChildElement(QString::fromUtf8(tagName.data, tagName.len)));
}

int QDomNode_LineNumber(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->lineNumber();
}

struct QtXml_PackedString QDomNode_LocalName(void* ptr)
{
	return ({ QByteArray* tfb3dda = new QByteArray(static_cast<QDomNode*>(ptr)->localName().toUtf8()); QtXml_PackedString { const_cast<char*>(tfb3dda->prepend("WHITESPACE").constData()+10), tfb3dda->size()-10, tfb3dda }; });
}

void* QDomNode_NamedItem(void* ptr, struct QtXml_PackedString name)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->namedItem(QString::fromUtf8(name.data, name.len)));
}

struct QtXml_PackedString QDomNode_NamespaceURI(void* ptr)
{
	return ({ QByteArray* tdc1910 = new QByteArray(static_cast<QDomNode*>(ptr)->namespaceURI().toUtf8()); QtXml_PackedString { const_cast<char*>(tdc1910->prepend("WHITESPACE").constData()+10), tdc1910->size()-10, tdc1910 }; });
}

void* QDomNode_NextSibling(void* ptr)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->nextSibling());
}

void* QDomNode_NextSiblingElement(void* ptr, struct QtXml_PackedString tagName)
{
	return new QDomElement(static_cast<QDomNode*>(ptr)->nextSiblingElement(QString::fromUtf8(tagName.data, tagName.len)));
}

struct QtXml_PackedString QDomNode_NodeName(void* ptr)
{
	return ({ QByteArray* t13e4b7 = new QByteArray(static_cast<QDomNode*>(ptr)->nodeName().toUtf8()); QtXml_PackedString { const_cast<char*>(t13e4b7->prepend("WHITESPACE").constData()+10), t13e4b7->size()-10, t13e4b7 }; });
}

long long QDomNode_NodeType(void* ptr)
{
	return static_cast<QDomNode*>(ptr)->nodeType();
}

struct QtXml_PackedString QDomNode_NodeValue(void* ptr)
{
	return ({ QByteArray* ta94d2a = new QByteArray(static_cast<QDomNode*>(ptr)->nodeValue().toUtf8()); QtXml_PackedString { const_cast<char*>(ta94d2a->prepend("WHITESPACE").constData()+10), ta94d2a->size()-10, ta94d2a }; });
}

void QDomNode_Normalize(void* ptr)
{
	static_cast<QDomNode*>(ptr)->normalize();
}

void* QDomNode_OwnerDocument(void* ptr)
{
	return new QDomDocument(static_cast<QDomNode*>(ptr)->ownerDocument());
}

void* QDomNode_ParentNode(void* ptr)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->parentNode());
}

struct QtXml_PackedString QDomNode_Prefix(void* ptr)
{
	return ({ QByteArray* t80c566 = new QByteArray(static_cast<QDomNode*>(ptr)->prefix().toUtf8()); QtXml_PackedString { const_cast<char*>(t80c566->prepend("WHITESPACE").constData()+10), t80c566->size()-10, t80c566 }; });
}

void* QDomNode_PreviousSibling(void* ptr)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->previousSibling());
}

void* QDomNode_PreviousSiblingElement(void* ptr, struct QtXml_PackedString tagName)
{
	return new QDomElement(static_cast<QDomNode*>(ptr)->previousSiblingElement(QString::fromUtf8(tagName.data, tagName.len)));
}

void* QDomNode_RemoveChild(void* ptr, void* oldChild)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->removeChild(*static_cast<QDomNode*>(oldChild)));
}

void* QDomNode_ReplaceChild(void* ptr, void* newChild, void* oldChild)
{
	return new QDomNode(static_cast<QDomNode*>(ptr)->replaceChild(*static_cast<QDomNode*>(newChild), *static_cast<QDomNode*>(oldChild)));
}

void QDomNode_Save(void* ptr, void* stream, int indent, long long encodingPolicy)
{
	static_cast<QDomNode*>(ptr)->save(*static_cast<QTextStream*>(stream), indent, static_cast<QDomNode::EncodingPolicy>(encodingPolicy));
}

void QDomNode_SetNodeValue(void* ptr, struct QtXml_PackedString v)
{
	static_cast<QDomNode*>(ptr)->setNodeValue(QString::fromUtf8(v.data, v.len));
}

void QDomNode_SetPrefix(void* ptr, struct QtXml_PackedString pre)
{
	static_cast<QDomNode*>(ptr)->setPrefix(QString::fromUtf8(pre.data, pre.len));
}

void* QDomNode_ToAttr(void* ptr)
{
	return new QDomAttr(static_cast<QDomNode*>(ptr)->toAttr());
}

void* QDomNode_ToCDATASection(void* ptr)
{
	return new QDomCDATASection(static_cast<QDomNode*>(ptr)->toCDATASection());
}

void* QDomNode_ToCharacterData(void* ptr)
{
	return new QDomCharacterData(static_cast<QDomNode*>(ptr)->toCharacterData());
}

void* QDomNode_ToComment(void* ptr)
{
	return new QDomComment(static_cast<QDomNode*>(ptr)->toComment());
}

void* QDomNode_ToDocument(void* ptr)
{
	return new QDomDocument(static_cast<QDomNode*>(ptr)->toDocument());
}

void* QDomNode_ToDocumentFragment(void* ptr)
{
	return new QDomDocumentFragment(static_cast<QDomNode*>(ptr)->toDocumentFragment());
}

void* QDomNode_ToDocumentType(void* ptr)
{
	return new QDomDocumentType(static_cast<QDomNode*>(ptr)->toDocumentType());
}

void* QDomNode_ToElement(void* ptr)
{
	return new QDomElement(static_cast<QDomNode*>(ptr)->toElement());
}

void* QDomNode_ToEntity(void* ptr)
{
	return new QDomEntity(static_cast<QDomNode*>(ptr)->toEntity());
}

void* QDomNode_ToEntityReference(void* ptr)
{
	return new QDomEntityReference(static_cast<QDomNode*>(ptr)->toEntityReference());
}

void* QDomNode_ToNotation(void* ptr)
{
	return new QDomNotation(static_cast<QDomNode*>(ptr)->toNotation());
}

void* QDomNode_ToProcessingInstruction(void* ptr)
{
	return new QDomProcessingInstruction(static_cast<QDomNode*>(ptr)->toProcessingInstruction());
}

void* QDomNode_ToText(void* ptr)
{
	return new QDomText(static_cast<QDomNode*>(ptr)->toText());
}

void QDomNode_DestroyQDomNode(void* ptr)
{
	static_cast<QDomNode*>(ptr)->~QDomNode();
}

Q_DECLARE_METATYPE(QDomNodeList)
Q_DECLARE_METATYPE(QDomNodeList*)
void* QDomNodeList_NewQDomNodeList()
{
	return new QDomNodeList();
}

void* QDomNodeList_NewQDomNodeList2(void* n)
{
	return new QDomNodeList(*static_cast<QDomNodeList*>(n));
}

void* QDomNodeList_At(void* ptr, int index)
{
	return new QDomNode(static_cast<QDomNodeList*>(ptr)->at(index));
}

int QDomNodeList_Count(void* ptr)
{
	return static_cast<QDomNodeList*>(ptr)->count();
}

char QDomNodeList_IsEmpty(void* ptr)
{
	return static_cast<QDomNodeList*>(ptr)->isEmpty();
}

void* QDomNodeList_Item(void* ptr, int index)
{
	return new QDomNode(static_cast<QDomNodeList*>(ptr)->item(index));
}

int QDomNodeList_Length(void* ptr)
{
	return static_cast<QDomNodeList*>(ptr)->length();
}

int QDomNodeList_Size(void* ptr)
{
	return static_cast<QDomNodeList*>(ptr)->size();
}

void QDomNodeList_DestroyQDomNodeList(void* ptr)
{
	static_cast<QDomNodeList*>(ptr)->~QDomNodeList();
}

Q_DECLARE_METATYPE(QDomNotation)
Q_DECLARE_METATYPE(QDomNotation*)
void* QDomNotation_NewQDomNotation()
{
	return new QDomNotation();
}

void* QDomNotation_NewQDomNotation2(void* x)
{
	return new QDomNotation(*static_cast<QDomNotation*>(x));
}

struct QtXml_PackedString QDomNotation_PublicId(void* ptr)
{
	return ({ QByteArray* t66d021 = new QByteArray(static_cast<QDomNotation*>(ptr)->publicId().toUtf8()); QtXml_PackedString { const_cast<char*>(t66d021->prepend("WHITESPACE").constData()+10), t66d021->size()-10, t66d021 }; });
}

struct QtXml_PackedString QDomNotation_SystemId(void* ptr)
{
	return ({ QByteArray* t08bdf6 = new QByteArray(static_cast<QDomNotation*>(ptr)->systemId().toUtf8()); QtXml_PackedString { const_cast<char*>(t08bdf6->prepend("WHITESPACE").constData()+10), t08bdf6->size()-10, t08bdf6 }; });
}

Q_DECLARE_METATYPE(QDomProcessingInstruction)
Q_DECLARE_METATYPE(QDomProcessingInstruction*)
void* QDomProcessingInstruction_NewQDomProcessingInstruction()
{
	return new QDomProcessingInstruction();
}

void* QDomProcessingInstruction_NewQDomProcessingInstruction2(void* x)
{
	return new QDomProcessingInstruction(*static_cast<QDomProcessingInstruction*>(x));
}

struct QtXml_PackedString QDomProcessingInstruction_Data(void* ptr)
{
	return ({ QByteArray* t7d3e8d = new QByteArray(static_cast<QDomProcessingInstruction*>(ptr)->data().toUtf8()); QtXml_PackedString { const_cast<char*>(t7d3e8d->prepend("WHITESPACE").constData()+10), t7d3e8d->size()-10, t7d3e8d }; });
}

void QDomProcessingInstruction_SetData(void* ptr, struct QtXml_PackedString d)
{
	static_cast<QDomProcessingInstruction*>(ptr)->setData(QString::fromUtf8(d.data, d.len));
}

struct QtXml_PackedString QDomProcessingInstruction_Target(void* ptr)
{
	return ({ QByteArray* td6875f = new QByteArray(static_cast<QDomProcessingInstruction*>(ptr)->target().toUtf8()); QtXml_PackedString { const_cast<char*>(td6875f->prepend("WHITESPACE").constData()+10), td6875f->size()-10, td6875f }; });
}

Q_DECLARE_METATYPE(QDomText)
Q_DECLARE_METATYPE(QDomText*)
void* QDomText_NewQDomText()
{
	return new QDomText();
}

void* QDomText_NewQDomText2(void* x)
{
	return new QDomText(*static_cast<QDomText*>(x));
}

void* QDomText_SplitText(void* ptr, int offset)
{
	return new QDomText(static_cast<QDomText*>(ptr)->splitText(offset));
}

class MyQXmlAttributes: public QXmlAttributes
{
public:
	MyQXmlAttributes() : QXmlAttributes() {QXmlAttributes_QXmlAttributes_QRegisterMetaType();};
	 ~MyQXmlAttributes() { callbackQXmlAttributes_DestroyQXmlAttributes(this); };
};

Q_DECLARE_METATYPE(QXmlAttributes*)
Q_DECLARE_METATYPE(MyQXmlAttributes*)

int QXmlAttributes_QXmlAttributes_QRegisterMetaType(){qRegisterMetaType<QXmlAttributes*>(); return qRegisterMetaType<MyQXmlAttributes*>();}

void* QXmlAttributes_NewQXmlAttributes()
{
	return new MyQXmlAttributes();
}

void QXmlAttributes_Append(void* ptr, struct QtXml_PackedString qName, struct QtXml_PackedString uri, struct QtXml_PackedString localPart, struct QtXml_PackedString value)
{
	static_cast<QXmlAttributes*>(ptr)->append(QString::fromUtf8(qName.data, qName.len), QString::fromUtf8(uri.data, uri.len), QString::fromUtf8(localPart.data, localPart.len), QString::fromUtf8(value.data, value.len));
}

void QXmlAttributes_Clear(void* ptr)
{
	static_cast<QXmlAttributes*>(ptr)->clear();
}

int QXmlAttributes_Count(void* ptr)
{
	return static_cast<QXmlAttributes*>(ptr)->count();
}

int QXmlAttributes_Index(void* ptr, struct QtXml_PackedString qName)
{
	return static_cast<QXmlAttributes*>(ptr)->index(QString::fromUtf8(qName.data, qName.len));
}

int QXmlAttributes_Index2(void* ptr, void* qName)
{
	return static_cast<QXmlAttributes*>(ptr)->index(*static_cast<QLatin1String*>(qName));
}

int QXmlAttributes_Index3(void* ptr, struct QtXml_PackedString uri, struct QtXml_PackedString localPart)
{
	return static_cast<QXmlAttributes*>(ptr)->index(QString::fromUtf8(uri.data, uri.len), QString::fromUtf8(localPart.data, localPart.len));
}

int QXmlAttributes_Length(void* ptr)
{
	return static_cast<QXmlAttributes*>(ptr)->length();
}

struct QtXml_PackedString QXmlAttributes_LocalName(void* ptr, int index)
{
	return ({ QByteArray* tb67c45 = new QByteArray(static_cast<QXmlAttributes*>(ptr)->localName(index).toUtf8()); QtXml_PackedString { const_cast<char*>(tb67c45->prepend("WHITESPACE").constData()+10), tb67c45->size()-10, tb67c45 }; });
}

struct QtXml_PackedString QXmlAttributes_QName(void* ptr, int index)
{
	return ({ QByteArray* t2cc196 = new QByteArray(static_cast<QXmlAttributes*>(ptr)->qName(index).toUtf8()); QtXml_PackedString { const_cast<char*>(t2cc196->prepend("WHITESPACE").constData()+10), t2cc196->size()-10, t2cc196 }; });
}

void QXmlAttributes_Swap(void* ptr, void* other)
{
	static_cast<QXmlAttributes*>(ptr)->swap(*static_cast<QXmlAttributes*>(other));
}

struct QtXml_PackedString QXmlAttributes_Type(void* ptr, int index)
{
	return ({ QByteArray* tb3583b = new QByteArray(static_cast<QXmlAttributes*>(ptr)->type(index).toUtf8()); QtXml_PackedString { const_cast<char*>(tb3583b->prepend("WHITESPACE").constData()+10), tb3583b->size()-10, tb3583b }; });
}

struct QtXml_PackedString QXmlAttributes_Type2(void* ptr, struct QtXml_PackedString qName)
{
	return ({ QByteArray* t067cce = new QByteArray(static_cast<QXmlAttributes*>(ptr)->type(QString::fromUtf8(qName.data, qName.len)).toUtf8()); QtXml_PackedString { const_cast<char*>(t067cce->prepend("WHITESPACE").constData()+10), t067cce->size()-10, t067cce }; });
}

struct QtXml_PackedString QXmlAttributes_Type3(void* ptr, struct QtXml_PackedString uri, struct QtXml_PackedString localName)
{
	return ({ QByteArray* t06dc3c = new QByteArray(static_cast<QXmlAttributes*>(ptr)->type(QString::fromUtf8(uri.data, uri.len), QString::fromUtf8(localName.data, localName.len)).toUtf8()); QtXml_PackedString { const_cast<char*>(t06dc3c->prepend("WHITESPACE").constData()+10), t06dc3c->size()-10, t06dc3c }; });
}

struct QtXml_PackedString QXmlAttributes_Uri(void* ptr, int index)
{
	return ({ QByteArray* t2a53d2 = new QByteArray(static_cast<QXmlAttributes*>(ptr)->uri(index).toUtf8()); QtXml_PackedString { const_cast<char*>(t2a53d2->prepend("WHITESPACE").constData()+10), t2a53d2->size()-10, t2a53d2 }; });
}

struct QtXml_PackedString QXmlAttributes_Value(void* ptr, int index)
{
	return ({ QByteArray* ta313f3 = new QByteArray(static_cast<QXmlAttributes*>(ptr)->value(index).toUtf8()); QtXml_PackedString { const_cast<char*>(ta313f3->prepend("WHITESPACE").constData()+10), ta313f3->size()-10, ta313f3 }; });
}

struct QtXml_PackedString QXmlAttributes_Value2(void* ptr, struct QtXml_PackedString qName)
{
	return ({ QByteArray* t36da80 = new QByteArray(static_cast<QXmlAttributes*>(ptr)->value(QString::fromUtf8(qName.data, qName.len)).toUtf8()); QtXml_PackedString { const_cast<char*>(t36da80->prepend("WHITESPACE").constData()+10), t36da80->size()-10, t36da80 }; });
}

struct QtXml_PackedString QXmlAttributes_Value3(void* ptr, void* qName)
{
	return ({ QByteArray* t25af61 = new QByteArray(static_cast<QXmlAttributes*>(ptr)->value(*static_cast<QLatin1String*>(qName)).toUtf8()); QtXml_PackedString { const_cast<char*>(t25af61->prepend("WHITESPACE").constData()+10), t25af61->size()-10, t25af61 }; });
}

struct QtXml_PackedString QXmlAttributes_Value4(void* ptr, struct QtXml_PackedString uri, struct QtXml_PackedString localName)
{
	return ({ QByteArray* t04f934 = new QByteArray(static_cast<QXmlAttributes*>(ptr)->value(QString::fromUtf8(uri.data, uri.len), QString::fromUtf8(localName.data, localName.len)).toUtf8()); QtXml_PackedString { const_cast<char*>(t04f934->prepend("WHITESPACE").constData()+10), t04f934->size()-10, t04f934 }; });
}

void QXmlAttributes_DestroyQXmlAttributes(void* ptr)
{
	static_cast<QXmlAttributes*>(ptr)->~QXmlAttributes();
}

void QXmlAttributes_DestroyQXmlAttributesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXmlContentHandler: public QXmlContentHandler
{
public:
	bool characters(const QString & ch) { QByteArray* t482bd6 = new QByteArray(ch.toUtf8()); QtXml_PackedString chPacked = { const_cast<char*>(t482bd6->prepend("WHITESPACE").constData()+10), t482bd6->size()-10, t482bd6 };return callbackQXmlContentHandler_Characters(this, chPacked) != 0; };
	bool endDocument() { return callbackQXmlContentHandler_EndDocument(this) != 0; };
	bool endElement(const QString & namespaceURI, const QString & localName, const QString & qName) { QByteArray* t120278 = new QByteArray(namespaceURI.toUtf8()); QtXml_PackedString namespaceURIPacked = { const_cast<char*>(t120278->prepend("WHITESPACE").constData()+10), t120278->size()-10, t120278 };QByteArray* t9dcab1 = new QByteArray(localName.toUtf8()); QtXml_PackedString localNamePacked = { const_cast<char*>(t9dcab1->prepend("WHITESPACE").constData()+10), t9dcab1->size()-10, t9dcab1 };QByteArray* tbe6bf1 = new QByteArray(qName.toUtf8()); QtXml_PackedString qNamePacked = { const_cast<char*>(tbe6bf1->prepend("WHITESPACE").constData()+10), tbe6bf1->size()-10, tbe6bf1 };return callbackQXmlContentHandler_EndElement(this, namespaceURIPacked, localNamePacked, qNamePacked) != 0; };
	bool endPrefixMapping(const QString & prefix) { QByteArray* tb4ebfe = new QByteArray(prefix.toUtf8()); QtXml_PackedString prefixPacked = { const_cast<char*>(tb4ebfe->prepend("WHITESPACE").constData()+10), tb4ebfe->size()-10, tb4ebfe };return callbackQXmlContentHandler_EndPrefixMapping(this, prefixPacked) != 0; };
	QString errorString() const { return ({ QtXml_PackedString tempVal = callbackQXmlContentHandler_ErrorString(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool ignorableWhitespace(const QString & ch) { QByteArray* t482bd6 = new QByteArray(ch.toUtf8()); QtXml_PackedString chPacked = { const_cast<char*>(t482bd6->prepend("WHITESPACE").constData()+10), t482bd6->size()-10, t482bd6 };return callbackQXmlContentHandler_IgnorableWhitespace(this, chPacked) != 0; };
	bool processingInstruction(const QString & target, const QString & data) { QByteArray* t0e8a3a = new QByteArray(target.toUtf8()); QtXml_PackedString targetPacked = { const_cast<char*>(t0e8a3a->prepend("WHITESPACE").constData()+10), t0e8a3a->size()-10, t0e8a3a };QByteArray* ta17c9a = new QByteArray(data.toUtf8()); QtXml_PackedString dataPacked = { const_cast<char*>(ta17c9a->prepend("WHITESPACE").constData()+10), ta17c9a->size()-10, ta17c9a };return callbackQXmlContentHandler_ProcessingInstruction(this, targetPacked, dataPacked) != 0; };
	void setDocumentLocator(QXmlLocator * locator) { callbackQXmlContentHandler_SetDocumentLocator(this, locator); };
	bool skippedEntity(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlContentHandler_SkippedEntity(this, namePacked) != 0; };
	bool startDocument() { return callbackQXmlContentHandler_StartDocument(this) != 0; };
	bool startElement(const QString & namespaceURI, const QString & localName, const QString & qName, const QXmlAttributes & atts) { QByteArray* t120278 = new QByteArray(namespaceURI.toUtf8()); QtXml_PackedString namespaceURIPacked = { const_cast<char*>(t120278->prepend("WHITESPACE").constData()+10), t120278->size()-10, t120278 };QByteArray* t9dcab1 = new QByteArray(localName.toUtf8()); QtXml_PackedString localNamePacked = { const_cast<char*>(t9dcab1->prepend("WHITESPACE").constData()+10), t9dcab1->size()-10, t9dcab1 };QByteArray* tbe6bf1 = new QByteArray(qName.toUtf8()); QtXml_PackedString qNamePacked = { const_cast<char*>(tbe6bf1->prepend("WHITESPACE").constData()+10), tbe6bf1->size()-10, tbe6bf1 };return callbackQXmlContentHandler_StartElement(this, namespaceURIPacked, localNamePacked, qNamePacked, const_cast<QXmlAttributes*>(&atts)) != 0; };
	bool startPrefixMapping(const QString & prefix, const QString & uri) { QByteArray* tb4ebfe = new QByteArray(prefix.toUtf8()); QtXml_PackedString prefixPacked = { const_cast<char*>(tb4ebfe->prepend("WHITESPACE").constData()+10), tb4ebfe->size()-10, tb4ebfe };QByteArray* t2c6d68 = new QByteArray(uri.toUtf8()); QtXml_PackedString uriPacked = { const_cast<char*>(t2c6d68->prepend("WHITESPACE").constData()+10), t2c6d68->size()-10, t2c6d68 };return callbackQXmlContentHandler_StartPrefixMapping(this, prefixPacked, uriPacked) != 0; };
	 ~MyQXmlContentHandler() { callbackQXmlContentHandler_DestroyQXmlContentHandler(this); };
};

Q_DECLARE_METATYPE(QXmlContentHandler*)
Q_DECLARE_METATYPE(MyQXmlContentHandler*)

int QXmlContentHandler_QXmlContentHandler_QRegisterMetaType(){qRegisterMetaType<QXmlContentHandler*>(); return qRegisterMetaType<MyQXmlContentHandler*>();}

char QXmlContentHandler_Characters(void* ptr, struct QtXml_PackedString ch)
{
	return static_cast<QXmlContentHandler*>(ptr)->characters(QString::fromUtf8(ch.data, ch.len));
}

char QXmlContentHandler_EndDocument(void* ptr)
{
	return static_cast<QXmlContentHandler*>(ptr)->endDocument();
}

char QXmlContentHandler_EndElement(void* ptr, struct QtXml_PackedString namespaceURI, struct QtXml_PackedString localName, struct QtXml_PackedString qName)
{
	return static_cast<QXmlContentHandler*>(ptr)->endElement(QString::fromUtf8(namespaceURI.data, namespaceURI.len), QString::fromUtf8(localName.data, localName.len), QString::fromUtf8(qName.data, qName.len));
}

char QXmlContentHandler_EndPrefixMapping(void* ptr, struct QtXml_PackedString prefix)
{
	return static_cast<QXmlContentHandler*>(ptr)->endPrefixMapping(QString::fromUtf8(prefix.data, prefix.len));
}

struct QtXml_PackedString QXmlContentHandler_ErrorString(void* ptr)
{
	return ({ QByteArray* t3c3aa4 = new QByteArray(static_cast<QXmlContentHandler*>(ptr)->errorString().toUtf8()); QtXml_PackedString { const_cast<char*>(t3c3aa4->prepend("WHITESPACE").constData()+10), t3c3aa4->size()-10, t3c3aa4 }; });
}

char QXmlContentHandler_IgnorableWhitespace(void* ptr, struct QtXml_PackedString ch)
{
	return static_cast<QXmlContentHandler*>(ptr)->ignorableWhitespace(QString::fromUtf8(ch.data, ch.len));
}

char QXmlContentHandler_ProcessingInstruction(void* ptr, struct QtXml_PackedString target, struct QtXml_PackedString data)
{
	return static_cast<QXmlContentHandler*>(ptr)->processingInstruction(QString::fromUtf8(target.data, target.len), QString::fromUtf8(data.data, data.len));
}

void QXmlContentHandler_SetDocumentLocator(void* ptr, void* locator)
{
	static_cast<QXmlContentHandler*>(ptr)->setDocumentLocator(static_cast<QXmlLocator*>(locator));
}

char QXmlContentHandler_SkippedEntity(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QXmlContentHandler*>(ptr)->skippedEntity(QString::fromUtf8(name.data, name.len));
}

char QXmlContentHandler_StartDocument(void* ptr)
{
	return static_cast<QXmlContentHandler*>(ptr)->startDocument();
}

char QXmlContentHandler_StartElement(void* ptr, struct QtXml_PackedString namespaceURI, struct QtXml_PackedString localName, struct QtXml_PackedString qName, void* atts)
{
	return static_cast<QXmlContentHandler*>(ptr)->startElement(QString::fromUtf8(namespaceURI.data, namespaceURI.len), QString::fromUtf8(localName.data, localName.len), QString::fromUtf8(qName.data, qName.len), *static_cast<QXmlAttributes*>(atts));
}

char QXmlContentHandler_StartPrefixMapping(void* ptr, struct QtXml_PackedString prefix, struct QtXml_PackedString uri)
{
	return static_cast<QXmlContentHandler*>(ptr)->startPrefixMapping(QString::fromUtf8(prefix.data, prefix.len), QString::fromUtf8(uri.data, uri.len));
}

void QXmlContentHandler_DestroyQXmlContentHandler(void* ptr)
{
	static_cast<QXmlContentHandler*>(ptr)->~QXmlContentHandler();
}

void QXmlContentHandler_DestroyQXmlContentHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXmlDTDHandler: public QXmlDTDHandler
{
public:
	QString errorString() const { return ({ QtXml_PackedString tempVal = callbackQXmlDTDHandler_ErrorString(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool notationDecl(const QString & name, const QString & publicId, const QString & systemId) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };QByteArray* tcfc7b7 = new QByteArray(publicId.toUtf8()); QtXml_PackedString publicIdPacked = { const_cast<char*>(tcfc7b7->prepend("WHITESPACE").constData()+10), tcfc7b7->size()-10, tcfc7b7 };QByteArray* te11426 = new QByteArray(systemId.toUtf8()); QtXml_PackedString systemIdPacked = { const_cast<char*>(te11426->prepend("WHITESPACE").constData()+10), te11426->size()-10, te11426 };return callbackQXmlDTDHandler_NotationDecl(this, namePacked, publicIdPacked, systemIdPacked) != 0; };
	bool unparsedEntityDecl(const QString & name, const QString & publicId, const QString & systemId, const QString & notationName) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };QByteArray* tcfc7b7 = new QByteArray(publicId.toUtf8()); QtXml_PackedString publicIdPacked = { const_cast<char*>(tcfc7b7->prepend("WHITESPACE").constData()+10), tcfc7b7->size()-10, tcfc7b7 };QByteArray* te11426 = new QByteArray(systemId.toUtf8()); QtXml_PackedString systemIdPacked = { const_cast<char*>(te11426->prepend("WHITESPACE").constData()+10), te11426->size()-10, te11426 };QByteArray* te46e8f = new QByteArray(notationName.toUtf8()); QtXml_PackedString notationNamePacked = { const_cast<char*>(te46e8f->prepend("WHITESPACE").constData()+10), te46e8f->size()-10, te46e8f };return callbackQXmlDTDHandler_UnparsedEntityDecl(this, namePacked, publicIdPacked, systemIdPacked, notationNamePacked) != 0; };
	 ~MyQXmlDTDHandler() { callbackQXmlDTDHandler_DestroyQXmlDTDHandler(this); };
};

Q_DECLARE_METATYPE(QXmlDTDHandler*)
Q_DECLARE_METATYPE(MyQXmlDTDHandler*)

int QXmlDTDHandler_QXmlDTDHandler_QRegisterMetaType(){qRegisterMetaType<QXmlDTDHandler*>(); return qRegisterMetaType<MyQXmlDTDHandler*>();}

struct QtXml_PackedString QXmlDTDHandler_ErrorString(void* ptr)
{
	return ({ QByteArray* t4437d5 = new QByteArray(static_cast<QXmlDTDHandler*>(ptr)->errorString().toUtf8()); QtXml_PackedString { const_cast<char*>(t4437d5->prepend("WHITESPACE").constData()+10), t4437d5->size()-10, t4437d5 }; });
}

char QXmlDTDHandler_NotationDecl(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId)
{
	return static_cast<QXmlDTDHandler*>(ptr)->notationDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len));
}

char QXmlDTDHandler_UnparsedEntityDecl(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId, struct QtXml_PackedString notationName)
{
	return static_cast<QXmlDTDHandler*>(ptr)->unparsedEntityDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len), QString::fromUtf8(notationName.data, notationName.len));
}

void QXmlDTDHandler_DestroyQXmlDTDHandler(void* ptr)
{
	static_cast<QXmlDTDHandler*>(ptr)->~QXmlDTDHandler();
}

void QXmlDTDHandler_DestroyQXmlDTDHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXmlDeclHandler: public QXmlDeclHandler
{
public:
	bool attributeDecl(const QString & eName, const QString & aName, const QString & ty, const QString & valueDefault, const QString & value) { QByteArray* t029528 = new QByteArray(eName.toUtf8()); QtXml_PackedString eNamePacked = { const_cast<char*>(t029528->prepend("WHITESPACE").constData()+10), t029528->size()-10, t029528 };QByteArray* tb01192 = new QByteArray(aName.toUtf8()); QtXml_PackedString aNamePacked = { const_cast<char*>(tb01192->prepend("WHITESPACE").constData()+10), tb01192->size()-10, tb01192 };QByteArray* td0a3e7 = new QByteArray(ty.toUtf8()); QtXml_PackedString tyPacked = { const_cast<char*>(td0a3e7->prepend("WHITESPACE").constData()+10), td0a3e7->size()-10, td0a3e7 };QByteArray* t4d15e2 = new QByteArray(valueDefault.toUtf8()); QtXml_PackedString valueDefaultPacked = { const_cast<char*>(t4d15e2->prepend("WHITESPACE").constData()+10), t4d15e2->size()-10, t4d15e2 };QByteArray* tf32b67 = new QByteArray(value.toUtf8()); QtXml_PackedString valuePacked = { const_cast<char*>(tf32b67->prepend("WHITESPACE").constData()+10), tf32b67->size()-10, tf32b67 };return callbackQXmlDeclHandler_AttributeDecl(this, eNamePacked, aNamePacked, tyPacked, valueDefaultPacked, valuePacked) != 0; };
	QString errorString() const { return ({ QtXml_PackedString tempVal = callbackQXmlDeclHandler_ErrorString(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool externalEntityDecl(const QString & name, const QString & publicId, const QString & systemId) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };QByteArray* tcfc7b7 = new QByteArray(publicId.toUtf8()); QtXml_PackedString publicIdPacked = { const_cast<char*>(tcfc7b7->prepend("WHITESPACE").constData()+10), tcfc7b7->size()-10, tcfc7b7 };QByteArray* te11426 = new QByteArray(systemId.toUtf8()); QtXml_PackedString systemIdPacked = { const_cast<char*>(te11426->prepend("WHITESPACE").constData()+10), te11426->size()-10, te11426 };return callbackQXmlDeclHandler_ExternalEntityDecl(this, namePacked, publicIdPacked, systemIdPacked) != 0; };
	bool internalEntityDecl(const QString & name, const QString & value) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };QByteArray* tf32b67 = new QByteArray(value.toUtf8()); QtXml_PackedString valuePacked = { const_cast<char*>(tf32b67->prepend("WHITESPACE").constData()+10), tf32b67->size()-10, tf32b67 };return callbackQXmlDeclHandler_InternalEntityDecl(this, namePacked, valuePacked) != 0; };
	 ~MyQXmlDeclHandler() { callbackQXmlDeclHandler_DestroyQXmlDeclHandler(this); };
};

Q_DECLARE_METATYPE(QXmlDeclHandler*)
Q_DECLARE_METATYPE(MyQXmlDeclHandler*)

int QXmlDeclHandler_QXmlDeclHandler_QRegisterMetaType(){qRegisterMetaType<QXmlDeclHandler*>(); return qRegisterMetaType<MyQXmlDeclHandler*>();}

char QXmlDeclHandler_AttributeDecl(void* ptr, struct QtXml_PackedString eName, struct QtXml_PackedString aName, struct QtXml_PackedString ty, struct QtXml_PackedString valueDefault, struct QtXml_PackedString value)
{
	return static_cast<QXmlDeclHandler*>(ptr)->attributeDecl(QString::fromUtf8(eName.data, eName.len), QString::fromUtf8(aName.data, aName.len), QString::fromUtf8(ty.data, ty.len), QString::fromUtf8(valueDefault.data, valueDefault.len), QString::fromUtf8(value.data, value.len));
}

struct QtXml_PackedString QXmlDeclHandler_ErrorString(void* ptr)
{
	return ({ QByteArray* t70b5b5 = new QByteArray(static_cast<QXmlDeclHandler*>(ptr)->errorString().toUtf8()); QtXml_PackedString { const_cast<char*>(t70b5b5->prepend("WHITESPACE").constData()+10), t70b5b5->size()-10, t70b5b5 }; });
}

char QXmlDeclHandler_ExternalEntityDecl(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId)
{
	return static_cast<QXmlDeclHandler*>(ptr)->externalEntityDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len));
}

char QXmlDeclHandler_InternalEntityDecl(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString value)
{
	return static_cast<QXmlDeclHandler*>(ptr)->internalEntityDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(value.data, value.len));
}

void QXmlDeclHandler_DestroyQXmlDeclHandler(void* ptr)
{
	static_cast<QXmlDeclHandler*>(ptr)->~QXmlDeclHandler();
}

void QXmlDeclHandler_DestroyQXmlDeclHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXmlDefaultHandler: public QXmlDefaultHandler
{
public:
	MyQXmlDefaultHandler() : QXmlDefaultHandler() {QXmlDefaultHandler_QXmlDefaultHandler_QRegisterMetaType();};
	bool attributeDecl(const QString & eName, const QString & aName, const QString & ty, const QString & valueDefault, const QString & value) { QByteArray* t029528 = new QByteArray(eName.toUtf8()); QtXml_PackedString eNamePacked = { const_cast<char*>(t029528->prepend("WHITESPACE").constData()+10), t029528->size()-10, t029528 };QByteArray* tb01192 = new QByteArray(aName.toUtf8()); QtXml_PackedString aNamePacked = { const_cast<char*>(tb01192->prepend("WHITESPACE").constData()+10), tb01192->size()-10, tb01192 };QByteArray* td0a3e7 = new QByteArray(ty.toUtf8()); QtXml_PackedString tyPacked = { const_cast<char*>(td0a3e7->prepend("WHITESPACE").constData()+10), td0a3e7->size()-10, td0a3e7 };QByteArray* t4d15e2 = new QByteArray(valueDefault.toUtf8()); QtXml_PackedString valueDefaultPacked = { const_cast<char*>(t4d15e2->prepend("WHITESPACE").constData()+10), t4d15e2->size()-10, t4d15e2 };QByteArray* tf32b67 = new QByteArray(value.toUtf8()); QtXml_PackedString valuePacked = { const_cast<char*>(tf32b67->prepend("WHITESPACE").constData()+10), tf32b67->size()-10, tf32b67 };return callbackQXmlDefaultHandler_AttributeDecl(this, eNamePacked, aNamePacked, tyPacked, valueDefaultPacked, valuePacked) != 0; };
	bool characters(const QString & ch) { QByteArray* t482bd6 = new QByteArray(ch.toUtf8()); QtXml_PackedString chPacked = { const_cast<char*>(t482bd6->prepend("WHITESPACE").constData()+10), t482bd6->size()-10, t482bd6 };return callbackQXmlDefaultHandler_Characters(this, chPacked) != 0; };
	bool comment(const QString & ch) { QByteArray* t482bd6 = new QByteArray(ch.toUtf8()); QtXml_PackedString chPacked = { const_cast<char*>(t482bd6->prepend("WHITESPACE").constData()+10), t482bd6->size()-10, t482bd6 };return callbackQXmlDefaultHandler_Comment(this, chPacked) != 0; };
	bool endCDATA() { return callbackQXmlDefaultHandler_EndCDATA(this) != 0; };
	bool endDTD() { return callbackQXmlDefaultHandler_EndDTD(this) != 0; };
	bool endDocument() { return callbackQXmlDefaultHandler_EndDocument(this) != 0; };
	bool endElement(const QString & namespaceURI, const QString & localName, const QString & qName) { QByteArray* t120278 = new QByteArray(namespaceURI.toUtf8()); QtXml_PackedString namespaceURIPacked = { const_cast<char*>(t120278->prepend("WHITESPACE").constData()+10), t120278->size()-10, t120278 };QByteArray* t9dcab1 = new QByteArray(localName.toUtf8()); QtXml_PackedString localNamePacked = { const_cast<char*>(t9dcab1->prepend("WHITESPACE").constData()+10), t9dcab1->size()-10, t9dcab1 };QByteArray* tbe6bf1 = new QByteArray(qName.toUtf8()); QtXml_PackedString qNamePacked = { const_cast<char*>(tbe6bf1->prepend("WHITESPACE").constData()+10), tbe6bf1->size()-10, tbe6bf1 };return callbackQXmlDefaultHandler_EndElement(this, namespaceURIPacked, localNamePacked, qNamePacked) != 0; };
	bool endEntity(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlDefaultHandler_EndEntity(this, namePacked) != 0; };
	bool endPrefixMapping(const QString & prefix) { QByteArray* tb4ebfe = new QByteArray(prefix.toUtf8()); QtXml_PackedString prefixPacked = { const_cast<char*>(tb4ebfe->prepend("WHITESPACE").constData()+10), tb4ebfe->size()-10, tb4ebfe };return callbackQXmlDefaultHandler_EndPrefixMapping(this, prefixPacked) != 0; };
	bool error(const QXmlParseException & exception) { return callbackQXmlDefaultHandler_Error(this, const_cast<QXmlParseException*>(&exception)) != 0; };
	QString errorString() const { return ({ QtXml_PackedString tempVal = callbackQXmlDefaultHandler_ErrorString(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool externalEntityDecl(const QString & name, const QString & publicId, const QString & systemId) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };QByteArray* tcfc7b7 = new QByteArray(publicId.toUtf8()); QtXml_PackedString publicIdPacked = { const_cast<char*>(tcfc7b7->prepend("WHITESPACE").constData()+10), tcfc7b7->size()-10, tcfc7b7 };QByteArray* te11426 = new QByteArray(systemId.toUtf8()); QtXml_PackedString systemIdPacked = { const_cast<char*>(te11426->prepend("WHITESPACE").constData()+10), te11426->size()-10, te11426 };return callbackQXmlDefaultHandler_ExternalEntityDecl(this, namePacked, publicIdPacked, systemIdPacked) != 0; };
	bool fatalError(const QXmlParseException & exception) { return callbackQXmlDefaultHandler_FatalError(this, const_cast<QXmlParseException*>(&exception)) != 0; };
	bool ignorableWhitespace(const QString & ch) { QByteArray* t482bd6 = new QByteArray(ch.toUtf8()); QtXml_PackedString chPacked = { const_cast<char*>(t482bd6->prepend("WHITESPACE").constData()+10), t482bd6->size()-10, t482bd6 };return callbackQXmlDefaultHandler_IgnorableWhitespace(this, chPacked) != 0; };
	bool internalEntityDecl(const QString & name, const QString & value) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };QByteArray* tf32b67 = new QByteArray(value.toUtf8()); QtXml_PackedString valuePacked = { const_cast<char*>(tf32b67->prepend("WHITESPACE").constData()+10), tf32b67->size()-10, tf32b67 };return callbackQXmlDefaultHandler_InternalEntityDecl(this, namePacked, valuePacked) != 0; };
	bool notationDecl(const QString & name, const QString & publicId, const QString & systemId) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };QByteArray* tcfc7b7 = new QByteArray(publicId.toUtf8()); QtXml_PackedString publicIdPacked = { const_cast<char*>(tcfc7b7->prepend("WHITESPACE").constData()+10), tcfc7b7->size()-10, tcfc7b7 };QByteArray* te11426 = new QByteArray(systemId.toUtf8()); QtXml_PackedString systemIdPacked = { const_cast<char*>(te11426->prepend("WHITESPACE").constData()+10), te11426->size()-10, te11426 };return callbackQXmlDefaultHandler_NotationDecl(this, namePacked, publicIdPacked, systemIdPacked) != 0; };
	bool processingInstruction(const QString & target, const QString & data) { QByteArray* t0e8a3a = new QByteArray(target.toUtf8()); QtXml_PackedString targetPacked = { const_cast<char*>(t0e8a3a->prepend("WHITESPACE").constData()+10), t0e8a3a->size()-10, t0e8a3a };QByteArray* ta17c9a = new QByteArray(data.toUtf8()); QtXml_PackedString dataPacked = { const_cast<char*>(ta17c9a->prepend("WHITESPACE").constData()+10), ta17c9a->size()-10, ta17c9a };return callbackQXmlDefaultHandler_ProcessingInstruction(this, targetPacked, dataPacked) != 0; };
	void setDocumentLocator(QXmlLocator * locator) { callbackQXmlDefaultHandler_SetDocumentLocator(this, locator); };
	bool skippedEntity(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlDefaultHandler_SkippedEntity(this, namePacked) != 0; };
	bool startCDATA() { return callbackQXmlDefaultHandler_StartCDATA(this) != 0; };
	bool startDTD(const QString & name, const QString & publicId, const QString & systemId) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };QByteArray* tcfc7b7 = new QByteArray(publicId.toUtf8()); QtXml_PackedString publicIdPacked = { const_cast<char*>(tcfc7b7->prepend("WHITESPACE").constData()+10), tcfc7b7->size()-10, tcfc7b7 };QByteArray* te11426 = new QByteArray(systemId.toUtf8()); QtXml_PackedString systemIdPacked = { const_cast<char*>(te11426->prepend("WHITESPACE").constData()+10), te11426->size()-10, te11426 };return callbackQXmlDefaultHandler_StartDTD(this, namePacked, publicIdPacked, systemIdPacked) != 0; };
	bool startDocument() { return callbackQXmlDefaultHandler_StartDocument(this) != 0; };
	bool startElement(const QString & namespaceURI, const QString & localName, const QString & qName, const QXmlAttributes & atts) { QByteArray* t120278 = new QByteArray(namespaceURI.toUtf8()); QtXml_PackedString namespaceURIPacked = { const_cast<char*>(t120278->prepend("WHITESPACE").constData()+10), t120278->size()-10, t120278 };QByteArray* t9dcab1 = new QByteArray(localName.toUtf8()); QtXml_PackedString localNamePacked = { const_cast<char*>(t9dcab1->prepend("WHITESPACE").constData()+10), t9dcab1->size()-10, t9dcab1 };QByteArray* tbe6bf1 = new QByteArray(qName.toUtf8()); QtXml_PackedString qNamePacked = { const_cast<char*>(tbe6bf1->prepend("WHITESPACE").constData()+10), tbe6bf1->size()-10, tbe6bf1 };return callbackQXmlDefaultHandler_StartElement(this, namespaceURIPacked, localNamePacked, qNamePacked, const_cast<QXmlAttributes*>(&atts)) != 0; };
	bool startEntity(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlDefaultHandler_StartEntity(this, namePacked) != 0; };
	bool startPrefixMapping(const QString & prefix, const QString & uri) { QByteArray* tb4ebfe = new QByteArray(prefix.toUtf8()); QtXml_PackedString prefixPacked = { const_cast<char*>(tb4ebfe->prepend("WHITESPACE").constData()+10), tb4ebfe->size()-10, tb4ebfe };QByteArray* t2c6d68 = new QByteArray(uri.toUtf8()); QtXml_PackedString uriPacked = { const_cast<char*>(t2c6d68->prepend("WHITESPACE").constData()+10), t2c6d68->size()-10, t2c6d68 };return callbackQXmlDefaultHandler_StartPrefixMapping(this, prefixPacked, uriPacked) != 0; };
	bool unparsedEntityDecl(const QString & name, const QString & publicId, const QString & systemId, const QString & notationName) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };QByteArray* tcfc7b7 = new QByteArray(publicId.toUtf8()); QtXml_PackedString publicIdPacked = { const_cast<char*>(tcfc7b7->prepend("WHITESPACE").constData()+10), tcfc7b7->size()-10, tcfc7b7 };QByteArray* te11426 = new QByteArray(systemId.toUtf8()); QtXml_PackedString systemIdPacked = { const_cast<char*>(te11426->prepend("WHITESPACE").constData()+10), te11426->size()-10, te11426 };QByteArray* te46e8f = new QByteArray(notationName.toUtf8()); QtXml_PackedString notationNamePacked = { const_cast<char*>(te46e8f->prepend("WHITESPACE").constData()+10), te46e8f->size()-10, te46e8f };return callbackQXmlDefaultHandler_UnparsedEntityDecl(this, namePacked, publicIdPacked, systemIdPacked, notationNamePacked) != 0; };
	bool warning(const QXmlParseException & exception) { return callbackQXmlDefaultHandler_Warning(this, const_cast<QXmlParseException*>(&exception)) != 0; };
	 ~MyQXmlDefaultHandler() { callbackQXmlDefaultHandler_DestroyQXmlDefaultHandler(this); };
};

Q_DECLARE_METATYPE(QXmlDefaultHandler*)
Q_DECLARE_METATYPE(MyQXmlDefaultHandler*)

int QXmlDefaultHandler_QXmlDefaultHandler_QRegisterMetaType(){qRegisterMetaType<QXmlDefaultHandler*>(); return qRegisterMetaType<MyQXmlDefaultHandler*>();}

void* QXmlDefaultHandler_NewQXmlDefaultHandler()
{
	return new MyQXmlDefaultHandler();
}

char QXmlDefaultHandler_AttributeDecl(void* ptr, struct QtXml_PackedString eName, struct QtXml_PackedString aName, struct QtXml_PackedString ty, struct QtXml_PackedString valueDefault, struct QtXml_PackedString value)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->attributeDecl(QString::fromUtf8(eName.data, eName.len), QString::fromUtf8(aName.data, aName.len), QString::fromUtf8(ty.data, ty.len), QString::fromUtf8(valueDefault.data, valueDefault.len), QString::fromUtf8(value.data, value.len));
}

char QXmlDefaultHandler_AttributeDeclDefault(void* ptr, struct QtXml_PackedString eName, struct QtXml_PackedString aName, struct QtXml_PackedString ty, struct QtXml_PackedString valueDefault, struct QtXml_PackedString value)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::attributeDecl(QString::fromUtf8(eName.data, eName.len), QString::fromUtf8(aName.data, aName.len), QString::fromUtf8(ty.data, ty.len), QString::fromUtf8(valueDefault.data, valueDefault.len), QString::fromUtf8(value.data, value.len));
}

char QXmlDefaultHandler_Characters(void* ptr, struct QtXml_PackedString ch)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->characters(QString::fromUtf8(ch.data, ch.len));
}

char QXmlDefaultHandler_CharactersDefault(void* ptr, struct QtXml_PackedString ch)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::characters(QString::fromUtf8(ch.data, ch.len));
}

char QXmlDefaultHandler_Comment(void* ptr, struct QtXml_PackedString ch)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->comment(QString::fromUtf8(ch.data, ch.len));
}

char QXmlDefaultHandler_CommentDefault(void* ptr, struct QtXml_PackedString ch)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::comment(QString::fromUtf8(ch.data, ch.len));
}

char QXmlDefaultHandler_EndCDATA(void* ptr)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->endCDATA();
}

char QXmlDefaultHandler_EndCDATADefault(void* ptr)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::endCDATA();
}

char QXmlDefaultHandler_EndDTD(void* ptr)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->endDTD();
}

char QXmlDefaultHandler_EndDTDDefault(void* ptr)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::endDTD();
}

char QXmlDefaultHandler_EndDocument(void* ptr)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->endDocument();
}

char QXmlDefaultHandler_EndDocumentDefault(void* ptr)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::endDocument();
}

char QXmlDefaultHandler_EndElement(void* ptr, struct QtXml_PackedString namespaceURI, struct QtXml_PackedString localName, struct QtXml_PackedString qName)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->endElement(QString::fromUtf8(namespaceURI.data, namespaceURI.len), QString::fromUtf8(localName.data, localName.len), QString::fromUtf8(qName.data, qName.len));
}

char QXmlDefaultHandler_EndElementDefault(void* ptr, struct QtXml_PackedString namespaceURI, struct QtXml_PackedString localName, struct QtXml_PackedString qName)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::endElement(QString::fromUtf8(namespaceURI.data, namespaceURI.len), QString::fromUtf8(localName.data, localName.len), QString::fromUtf8(qName.data, qName.len));
}

char QXmlDefaultHandler_EndEntity(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->endEntity(QString::fromUtf8(name.data, name.len));
}

char QXmlDefaultHandler_EndEntityDefault(void* ptr, struct QtXml_PackedString name)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::endEntity(QString::fromUtf8(name.data, name.len));
}

char QXmlDefaultHandler_EndPrefixMapping(void* ptr, struct QtXml_PackedString prefix)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->endPrefixMapping(QString::fromUtf8(prefix.data, prefix.len));
}

char QXmlDefaultHandler_EndPrefixMappingDefault(void* ptr, struct QtXml_PackedString prefix)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::endPrefixMapping(QString::fromUtf8(prefix.data, prefix.len));
}

char QXmlDefaultHandler_Error(void* ptr, void* exception)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->error(*static_cast<QXmlParseException*>(exception));
}

char QXmlDefaultHandler_ErrorDefault(void* ptr, void* exception)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::error(*static_cast<QXmlParseException*>(exception));
}

struct QtXml_PackedString QXmlDefaultHandler_ErrorString(void* ptr)
{
	return ({ QByteArray* tffa037 = new QByteArray(static_cast<QXmlDefaultHandler*>(ptr)->errorString().toUtf8()); QtXml_PackedString { const_cast<char*>(tffa037->prepend("WHITESPACE").constData()+10), tffa037->size()-10, tffa037 }; });
}

struct QtXml_PackedString QXmlDefaultHandler_ErrorStringDefault(void* ptr)
{
		return ({ QByteArray* tdd1e32 = new QByteArray(static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::errorString().toUtf8()); QtXml_PackedString { const_cast<char*>(tdd1e32->prepend("WHITESPACE").constData()+10), tdd1e32->size()-10, tdd1e32 }; });
}

char QXmlDefaultHandler_ExternalEntityDecl(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->externalEntityDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len));
}

char QXmlDefaultHandler_ExternalEntityDeclDefault(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::externalEntityDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len));
}

char QXmlDefaultHandler_FatalError(void* ptr, void* exception)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->fatalError(*static_cast<QXmlParseException*>(exception));
}

char QXmlDefaultHandler_FatalErrorDefault(void* ptr, void* exception)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::fatalError(*static_cast<QXmlParseException*>(exception));
}

char QXmlDefaultHandler_IgnorableWhitespace(void* ptr, struct QtXml_PackedString ch)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->ignorableWhitespace(QString::fromUtf8(ch.data, ch.len));
}

char QXmlDefaultHandler_IgnorableWhitespaceDefault(void* ptr, struct QtXml_PackedString ch)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::ignorableWhitespace(QString::fromUtf8(ch.data, ch.len));
}

char QXmlDefaultHandler_InternalEntityDecl(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString value)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->internalEntityDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(value.data, value.len));
}

char QXmlDefaultHandler_InternalEntityDeclDefault(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString value)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::internalEntityDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(value.data, value.len));
}

char QXmlDefaultHandler_NotationDecl(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->notationDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len));
}

char QXmlDefaultHandler_NotationDeclDefault(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::notationDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len));
}

char QXmlDefaultHandler_ProcessingInstruction(void* ptr, struct QtXml_PackedString target, struct QtXml_PackedString data)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->processingInstruction(QString::fromUtf8(target.data, target.len), QString::fromUtf8(data.data, data.len));
}

char QXmlDefaultHandler_ProcessingInstructionDefault(void* ptr, struct QtXml_PackedString target, struct QtXml_PackedString data)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::processingInstruction(QString::fromUtf8(target.data, target.len), QString::fromUtf8(data.data, data.len));
}

void QXmlDefaultHandler_SetDocumentLocator(void* ptr, void* locator)
{
	static_cast<QXmlDefaultHandler*>(ptr)->setDocumentLocator(static_cast<QXmlLocator*>(locator));
}

void QXmlDefaultHandler_SetDocumentLocatorDefault(void* ptr, void* locator)
{
		static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::setDocumentLocator(static_cast<QXmlLocator*>(locator));
}

char QXmlDefaultHandler_SkippedEntity(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->skippedEntity(QString::fromUtf8(name.data, name.len));
}

char QXmlDefaultHandler_SkippedEntityDefault(void* ptr, struct QtXml_PackedString name)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::skippedEntity(QString::fromUtf8(name.data, name.len));
}

char QXmlDefaultHandler_StartCDATA(void* ptr)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->startCDATA();
}

char QXmlDefaultHandler_StartCDATADefault(void* ptr)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::startCDATA();
}

char QXmlDefaultHandler_StartDTD(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->startDTD(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len));
}

char QXmlDefaultHandler_StartDTDDefault(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::startDTD(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len));
}

char QXmlDefaultHandler_StartDocument(void* ptr)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->startDocument();
}

char QXmlDefaultHandler_StartDocumentDefault(void* ptr)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::startDocument();
}

char QXmlDefaultHandler_StartElement(void* ptr, struct QtXml_PackedString namespaceURI, struct QtXml_PackedString localName, struct QtXml_PackedString qName, void* atts)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->startElement(QString::fromUtf8(namespaceURI.data, namespaceURI.len), QString::fromUtf8(localName.data, localName.len), QString::fromUtf8(qName.data, qName.len), *static_cast<QXmlAttributes*>(atts));
}

char QXmlDefaultHandler_StartElementDefault(void* ptr, struct QtXml_PackedString namespaceURI, struct QtXml_PackedString localName, struct QtXml_PackedString qName, void* atts)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::startElement(QString::fromUtf8(namespaceURI.data, namespaceURI.len), QString::fromUtf8(localName.data, localName.len), QString::fromUtf8(qName.data, qName.len), *static_cast<QXmlAttributes*>(atts));
}

char QXmlDefaultHandler_StartEntity(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->startEntity(QString::fromUtf8(name.data, name.len));
}

char QXmlDefaultHandler_StartEntityDefault(void* ptr, struct QtXml_PackedString name)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::startEntity(QString::fromUtf8(name.data, name.len));
}

char QXmlDefaultHandler_StartPrefixMapping(void* ptr, struct QtXml_PackedString prefix, struct QtXml_PackedString uri)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->startPrefixMapping(QString::fromUtf8(prefix.data, prefix.len), QString::fromUtf8(uri.data, uri.len));
}

char QXmlDefaultHandler_StartPrefixMappingDefault(void* ptr, struct QtXml_PackedString prefix, struct QtXml_PackedString uri)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::startPrefixMapping(QString::fromUtf8(prefix.data, prefix.len), QString::fromUtf8(uri.data, uri.len));
}

char QXmlDefaultHandler_UnparsedEntityDecl(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId, struct QtXml_PackedString notationName)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->unparsedEntityDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len), QString::fromUtf8(notationName.data, notationName.len));
}

char QXmlDefaultHandler_UnparsedEntityDeclDefault(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId, struct QtXml_PackedString notationName)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::unparsedEntityDecl(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len), QString::fromUtf8(notationName.data, notationName.len));
}

char QXmlDefaultHandler_Warning(void* ptr, void* exception)
{
	return static_cast<QXmlDefaultHandler*>(ptr)->warning(*static_cast<QXmlParseException*>(exception));
}

char QXmlDefaultHandler_WarningDefault(void* ptr, void* exception)
{
		return static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::warning(*static_cast<QXmlParseException*>(exception));
}

void QXmlDefaultHandler_DestroyQXmlDefaultHandler(void* ptr)
{
	static_cast<QXmlDefaultHandler*>(ptr)->~QXmlDefaultHandler();
}

void QXmlDefaultHandler_DestroyQXmlDefaultHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXmlEntityResolver: public QXmlEntityResolver
{
public:
	QString errorString() const { return ({ QtXml_PackedString tempVal = callbackQXmlEntityResolver_ErrorString(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	 ~MyQXmlEntityResolver() { callbackQXmlEntityResolver_DestroyQXmlEntityResolver(this); };
};

Q_DECLARE_METATYPE(QXmlEntityResolver*)
Q_DECLARE_METATYPE(MyQXmlEntityResolver*)

int QXmlEntityResolver_QXmlEntityResolver_QRegisterMetaType(){qRegisterMetaType<QXmlEntityResolver*>(); return qRegisterMetaType<MyQXmlEntityResolver*>();}

struct QtXml_PackedString QXmlEntityResolver_ErrorString(void* ptr)
{
	return ({ QByteArray* te6feb0 = new QByteArray(static_cast<QXmlEntityResolver*>(ptr)->errorString().toUtf8()); QtXml_PackedString { const_cast<char*>(te6feb0->prepend("WHITESPACE").constData()+10), te6feb0->size()-10, te6feb0 }; });
}

void QXmlEntityResolver_DestroyQXmlEntityResolver(void* ptr)
{
	static_cast<QXmlEntityResolver*>(ptr)->~QXmlEntityResolver();
}

void QXmlEntityResolver_DestroyQXmlEntityResolverDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXmlErrorHandler: public QXmlErrorHandler
{
public:
	bool error(const QXmlParseException & exception) { return callbackQXmlErrorHandler_Error(this, const_cast<QXmlParseException*>(&exception)) != 0; };
	QString errorString() const { return ({ QtXml_PackedString tempVal = callbackQXmlErrorHandler_ErrorString(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool fatalError(const QXmlParseException & exception) { return callbackQXmlErrorHandler_FatalError(this, const_cast<QXmlParseException*>(&exception)) != 0; };
	bool warning(const QXmlParseException & exception) { return callbackQXmlErrorHandler_Warning(this, const_cast<QXmlParseException*>(&exception)) != 0; };
	 ~MyQXmlErrorHandler() { callbackQXmlErrorHandler_DestroyQXmlErrorHandler(this); };
};

Q_DECLARE_METATYPE(QXmlErrorHandler*)
Q_DECLARE_METATYPE(MyQXmlErrorHandler*)

int QXmlErrorHandler_QXmlErrorHandler_QRegisterMetaType(){qRegisterMetaType<QXmlErrorHandler*>(); return qRegisterMetaType<MyQXmlErrorHandler*>();}

char QXmlErrorHandler_Error(void* ptr, void* exception)
{
	return static_cast<QXmlErrorHandler*>(ptr)->error(*static_cast<QXmlParseException*>(exception));
}

struct QtXml_PackedString QXmlErrorHandler_ErrorString(void* ptr)
{
	return ({ QByteArray* t2e99c1 = new QByteArray(static_cast<QXmlErrorHandler*>(ptr)->errorString().toUtf8()); QtXml_PackedString { const_cast<char*>(t2e99c1->prepend("WHITESPACE").constData()+10), t2e99c1->size()-10, t2e99c1 }; });
}

char QXmlErrorHandler_FatalError(void* ptr, void* exception)
{
	return static_cast<QXmlErrorHandler*>(ptr)->fatalError(*static_cast<QXmlParseException*>(exception));
}

char QXmlErrorHandler_Warning(void* ptr, void* exception)
{
	return static_cast<QXmlErrorHandler*>(ptr)->warning(*static_cast<QXmlParseException*>(exception));
}

void QXmlErrorHandler_DestroyQXmlErrorHandler(void* ptr)
{
	static_cast<QXmlErrorHandler*>(ptr)->~QXmlErrorHandler();
}

void QXmlErrorHandler_DestroyQXmlErrorHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXmlInputSource: public QXmlInputSource
{
public:
	MyQXmlInputSource() : QXmlInputSource() {QXmlInputSource_QXmlInputSource_QRegisterMetaType();};
	MyQXmlInputSource(QIODevice *dev) : QXmlInputSource(dev) {QXmlInputSource_QXmlInputSource_QRegisterMetaType();};
	QString data() const { return ({ QtXml_PackedString tempVal = callbackQXmlInputSource_Data(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void fetchData() { callbackQXmlInputSource_FetchData(this); };
	QString fromRawData(const QByteArray & data, bool beginning) { return ({ QtXml_PackedString tempVal = callbackQXmlInputSource_FromRawData(this, const_cast<QByteArray*>(&data), beginning); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QChar next() { return *static_cast<QChar*>(callbackQXmlInputSource_Next(this)); };
	void reset() { callbackQXmlInputSource_Reset(this); };
	void setData(const QString & dat) { QByteArray* tfbd0b9 = new QByteArray(dat.toUtf8()); QtXml_PackedString datPacked = { const_cast<char*>(tfbd0b9->prepend("WHITESPACE").constData()+10), tfbd0b9->size()-10, tfbd0b9 };callbackQXmlInputSource_SetData(this, datPacked); };
	void setData(const QByteArray & dat) { callbackQXmlInputSource_SetData2(this, const_cast<QByteArray*>(&dat)); };
	 ~MyQXmlInputSource() { callbackQXmlInputSource_DestroyQXmlInputSource(this); };
};

Q_DECLARE_METATYPE(QXmlInputSource*)
Q_DECLARE_METATYPE(MyQXmlInputSource*)

int QXmlInputSource_QXmlInputSource_QRegisterMetaType(){qRegisterMetaType<QXmlInputSource*>(); return qRegisterMetaType<MyQXmlInputSource*>();}

void* QXmlInputSource_NewQXmlInputSource()
{
	return new MyQXmlInputSource();
}

void* QXmlInputSource_NewQXmlInputSource2(void* dev)
{
	return new MyQXmlInputSource(static_cast<QIODevice*>(dev));
}

struct QtXml_PackedString QXmlInputSource_Data(void* ptr)
{
	return ({ QByteArray* tba80df = new QByteArray(static_cast<QXmlInputSource*>(ptr)->data().toUtf8()); QtXml_PackedString { const_cast<char*>(tba80df->prepend("WHITESPACE").constData()+10), tba80df->size()-10, tba80df }; });
}

struct QtXml_PackedString QXmlInputSource_DataDefault(void* ptr)
{
		return ({ QByteArray* ta6ef1f = new QByteArray(static_cast<QXmlInputSource*>(ptr)->QXmlInputSource::data().toUtf8()); QtXml_PackedString { const_cast<char*>(ta6ef1f->prepend("WHITESPACE").constData()+10), ta6ef1f->size()-10, ta6ef1f }; });
}

void QXmlInputSource_FetchData(void* ptr)
{
	static_cast<QXmlInputSource*>(ptr)->fetchData();
}

void QXmlInputSource_FetchDataDefault(void* ptr)
{
		static_cast<QXmlInputSource*>(ptr)->QXmlInputSource::fetchData();
}

struct QtXml_PackedString QXmlInputSource_FromRawData(void* ptr, void* data, char beginning)
{
	return ({ QByteArray* t4f313f = new QByteArray(static_cast<QXmlInputSource*>(ptr)->fromRawData(*static_cast<QByteArray*>(data), beginning != 0).toUtf8()); QtXml_PackedString { const_cast<char*>(t4f313f->prepend("WHITESPACE").constData()+10), t4f313f->size()-10, t4f313f }; });
}

struct QtXml_PackedString QXmlInputSource_FromRawDataDefault(void* ptr, void* data, char beginning)
{
		return ({ QByteArray* tded04d = new QByteArray(static_cast<QXmlInputSource*>(ptr)->QXmlInputSource::fromRawData(*static_cast<QByteArray*>(data), beginning != 0).toUtf8()); QtXml_PackedString { const_cast<char*>(tded04d->prepend("WHITESPACE").constData()+10), tded04d->size()-10, tded04d }; });
}

void* QXmlInputSource_Next(void* ptr)
{
	return new QChar(static_cast<QXmlInputSource*>(ptr)->next());
}

void* QXmlInputSource_NextDefault(void* ptr)
{
		return new QChar(static_cast<QXmlInputSource*>(ptr)->QXmlInputSource::next());
}

void QXmlInputSource_Reset(void* ptr)
{
	static_cast<QXmlInputSource*>(ptr)->reset();
}

void QXmlInputSource_ResetDefault(void* ptr)
{
		static_cast<QXmlInputSource*>(ptr)->QXmlInputSource::reset();
}

void QXmlInputSource_SetData(void* ptr, struct QtXml_PackedString dat)
{
	static_cast<QXmlInputSource*>(ptr)->setData(QString::fromUtf8(dat.data, dat.len));
}

void QXmlInputSource_SetDataDefault(void* ptr, struct QtXml_PackedString dat)
{
		static_cast<QXmlInputSource*>(ptr)->QXmlInputSource::setData(QString::fromUtf8(dat.data, dat.len));
}

void QXmlInputSource_SetData2(void* ptr, void* dat)
{
	static_cast<QXmlInputSource*>(ptr)->setData(*static_cast<QByteArray*>(dat));
}

void QXmlInputSource_SetData2Default(void* ptr, void* dat)
{
		static_cast<QXmlInputSource*>(ptr)->QXmlInputSource::setData(*static_cast<QByteArray*>(dat));
}

void QXmlInputSource_DestroyQXmlInputSource(void* ptr)
{
	static_cast<QXmlInputSource*>(ptr)->~QXmlInputSource();
}

void QXmlInputSource_DestroyQXmlInputSourceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXmlLexicalHandler: public QXmlLexicalHandler
{
public:
	bool comment(const QString & ch) { QByteArray* t482bd6 = new QByteArray(ch.toUtf8()); QtXml_PackedString chPacked = { const_cast<char*>(t482bd6->prepend("WHITESPACE").constData()+10), t482bd6->size()-10, t482bd6 };return callbackQXmlLexicalHandler_Comment(this, chPacked) != 0; };
	bool endCDATA() { return callbackQXmlLexicalHandler_EndCDATA(this) != 0; };
	bool endDTD() { return callbackQXmlLexicalHandler_EndDTD(this) != 0; };
	bool endEntity(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlLexicalHandler_EndEntity(this, namePacked) != 0; };
	QString errorString() const { return ({ QtXml_PackedString tempVal = callbackQXmlLexicalHandler_ErrorString(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool startCDATA() { return callbackQXmlLexicalHandler_StartCDATA(this) != 0; };
	bool startDTD(const QString & name, const QString & publicId, const QString & systemId) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };QByteArray* tcfc7b7 = new QByteArray(publicId.toUtf8()); QtXml_PackedString publicIdPacked = { const_cast<char*>(tcfc7b7->prepend("WHITESPACE").constData()+10), tcfc7b7->size()-10, tcfc7b7 };QByteArray* te11426 = new QByteArray(systemId.toUtf8()); QtXml_PackedString systemIdPacked = { const_cast<char*>(te11426->prepend("WHITESPACE").constData()+10), te11426->size()-10, te11426 };return callbackQXmlLexicalHandler_StartDTD(this, namePacked, publicIdPacked, systemIdPacked) != 0; };
	bool startEntity(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlLexicalHandler_StartEntity(this, namePacked) != 0; };
	 ~MyQXmlLexicalHandler() { callbackQXmlLexicalHandler_DestroyQXmlLexicalHandler(this); };
};

Q_DECLARE_METATYPE(QXmlLexicalHandler*)
Q_DECLARE_METATYPE(MyQXmlLexicalHandler*)

int QXmlLexicalHandler_QXmlLexicalHandler_QRegisterMetaType(){qRegisterMetaType<QXmlLexicalHandler*>(); return qRegisterMetaType<MyQXmlLexicalHandler*>();}

char QXmlLexicalHandler_Comment(void* ptr, struct QtXml_PackedString ch)
{
	return static_cast<QXmlLexicalHandler*>(ptr)->comment(QString::fromUtf8(ch.data, ch.len));
}

char QXmlLexicalHandler_EndCDATA(void* ptr)
{
	return static_cast<QXmlLexicalHandler*>(ptr)->endCDATA();
}

char QXmlLexicalHandler_EndDTD(void* ptr)
{
	return static_cast<QXmlLexicalHandler*>(ptr)->endDTD();
}

char QXmlLexicalHandler_EndEntity(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QXmlLexicalHandler*>(ptr)->endEntity(QString::fromUtf8(name.data, name.len));
}

struct QtXml_PackedString QXmlLexicalHandler_ErrorString(void* ptr)
{
	return ({ QByteArray* t0032b9 = new QByteArray(static_cast<QXmlLexicalHandler*>(ptr)->errorString().toUtf8()); QtXml_PackedString { const_cast<char*>(t0032b9->prepend("WHITESPACE").constData()+10), t0032b9->size()-10, t0032b9 }; });
}

char QXmlLexicalHandler_StartCDATA(void* ptr)
{
	return static_cast<QXmlLexicalHandler*>(ptr)->startCDATA();
}

char QXmlLexicalHandler_StartDTD(void* ptr, struct QtXml_PackedString name, struct QtXml_PackedString publicId, struct QtXml_PackedString systemId)
{
	return static_cast<QXmlLexicalHandler*>(ptr)->startDTD(QString::fromUtf8(name.data, name.len), QString::fromUtf8(publicId.data, publicId.len), QString::fromUtf8(systemId.data, systemId.len));
}

char QXmlLexicalHandler_StartEntity(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QXmlLexicalHandler*>(ptr)->startEntity(QString::fromUtf8(name.data, name.len));
}

void QXmlLexicalHandler_DestroyQXmlLexicalHandler(void* ptr)
{
	static_cast<QXmlLexicalHandler*>(ptr)->~QXmlLexicalHandler();
}

void QXmlLexicalHandler_DestroyQXmlLexicalHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXmlLocator: public QXmlLocator
{
public:
	MyQXmlLocator() : QXmlLocator() {QXmlLocator_QXmlLocator_QRegisterMetaType();};
	int columnNumber() const { return callbackQXmlLocator_ColumnNumber(const_cast<void*>(static_cast<const void*>(this))); };
	int lineNumber() const { return callbackQXmlLocator_LineNumber(const_cast<void*>(static_cast<const void*>(this))); };
	 ~MyQXmlLocator() { callbackQXmlLocator_DestroyQXmlLocator(this); };
};

Q_DECLARE_METATYPE(QXmlLocator*)
Q_DECLARE_METATYPE(MyQXmlLocator*)

int QXmlLocator_QXmlLocator_QRegisterMetaType(){qRegisterMetaType<QXmlLocator*>(); return qRegisterMetaType<MyQXmlLocator*>();}

void* QXmlLocator_NewQXmlLocator()
{
	return new MyQXmlLocator();
}

int QXmlLocator_ColumnNumber(void* ptr)
{
	return static_cast<QXmlLocator*>(ptr)->columnNumber();
}

int QXmlLocator_LineNumber(void* ptr)
{
	return static_cast<QXmlLocator*>(ptr)->lineNumber();
}

void QXmlLocator_DestroyQXmlLocator(void* ptr)
{
	static_cast<QXmlLocator*>(ptr)->~QXmlLocator();
}

void QXmlLocator_DestroyQXmlLocatorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

Q_DECLARE_METATYPE(QXmlNamespaceSupport*)
void* QXmlNamespaceSupport_NewQXmlNamespaceSupport()
{
	return new QXmlNamespaceSupport();
}

void QXmlNamespaceSupport_PopContext(void* ptr)
{
	static_cast<QXmlNamespaceSupport*>(ptr)->popContext();
}

struct QtXml_PackedString QXmlNamespaceSupport_Prefix(void* ptr, struct QtXml_PackedString uri)
{
	return ({ QByteArray* t38829d = new QByteArray(static_cast<QXmlNamespaceSupport*>(ptr)->prefix(QString::fromUtf8(uri.data, uri.len)).toUtf8()); QtXml_PackedString { const_cast<char*>(t38829d->prepend("WHITESPACE").constData()+10), t38829d->size()-10, t38829d }; });
}

struct QtXml_PackedString QXmlNamespaceSupport_Prefixes(void* ptr)
{
	return ({ QByteArray* t8faa65 = new QByteArray(static_cast<QXmlNamespaceSupport*>(ptr)->prefixes().join("¡¦!").toUtf8()); QtXml_PackedString { const_cast<char*>(t8faa65->prepend("WHITESPACE").constData()+10), t8faa65->size()-10, t8faa65 }; });
}

struct QtXml_PackedString QXmlNamespaceSupport_Prefixes2(void* ptr, struct QtXml_PackedString uri)
{
	return ({ QByteArray* t763c75 = new QByteArray(static_cast<QXmlNamespaceSupport*>(ptr)->prefixes(QString::fromUtf8(uri.data, uri.len)).join("¡¦!").toUtf8()); QtXml_PackedString { const_cast<char*>(t763c75->prepend("WHITESPACE").constData()+10), t763c75->size()-10, t763c75 }; });
}

void QXmlNamespaceSupport_ProcessName(void* ptr, struct QtXml_PackedString qname, char isAttribute, struct QtXml_PackedString nsuri, struct QtXml_PackedString localname)
{
	static_cast<QXmlNamespaceSupport*>(ptr)->processName(QString::fromUtf8(qname.data, qname.len), isAttribute != 0, *(new QString(QString::fromUtf8(nsuri.data, nsuri.len))), *(new QString(QString::fromUtf8(localname.data, localname.len))));
}

void QXmlNamespaceSupport_PushContext(void* ptr)
{
	static_cast<QXmlNamespaceSupport*>(ptr)->pushContext();
}

void QXmlNamespaceSupport_Reset(void* ptr)
{
	static_cast<QXmlNamespaceSupport*>(ptr)->reset();
}

void QXmlNamespaceSupport_SetPrefix(void* ptr, struct QtXml_PackedString pre, struct QtXml_PackedString uri)
{
	static_cast<QXmlNamespaceSupport*>(ptr)->setPrefix(QString::fromUtf8(pre.data, pre.len), QString::fromUtf8(uri.data, uri.len));
}

void QXmlNamespaceSupport_SplitName(void* ptr, struct QtXml_PackedString qname, struct QtXml_PackedString prefix, struct QtXml_PackedString localname)
{
	static_cast<QXmlNamespaceSupport*>(ptr)->splitName(QString::fromUtf8(qname.data, qname.len), *(new QString(QString::fromUtf8(prefix.data, prefix.len))), *(new QString(QString::fromUtf8(localname.data, localname.len))));
}

struct QtXml_PackedString QXmlNamespaceSupport_Uri(void* ptr, struct QtXml_PackedString prefix)
{
	return ({ QByteArray* tede5ec = new QByteArray(static_cast<QXmlNamespaceSupport*>(ptr)->uri(QString::fromUtf8(prefix.data, prefix.len)).toUtf8()); QtXml_PackedString { const_cast<char*>(tede5ec->prepend("WHITESPACE").constData()+10), tede5ec->size()-10, tede5ec }; });
}

void QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(void* ptr)
{
	static_cast<QXmlNamespaceSupport*>(ptr)->~QXmlNamespaceSupport();
}

Q_DECLARE_METATYPE(QXmlParseException*)
void* QXmlParseException_NewQXmlParseException(struct QtXml_PackedString name, int c, int l, struct QtXml_PackedString p, struct QtXml_PackedString s)
{
	return new QXmlParseException(QString::fromUtf8(name.data, name.len), c, l, QString::fromUtf8(p.data, p.len), QString::fromUtf8(s.data, s.len));
}

void* QXmlParseException_NewQXmlParseException2(void* other)
{
	return new QXmlParseException(*static_cast<QXmlParseException*>(other));
}

int QXmlParseException_ColumnNumber(void* ptr)
{
	return static_cast<QXmlParseException*>(ptr)->columnNumber();
}

int QXmlParseException_LineNumber(void* ptr)
{
	return static_cast<QXmlParseException*>(ptr)->lineNumber();
}

struct QtXml_PackedString QXmlParseException_Message(void* ptr)
{
	return ({ QByteArray* ta2e963 = new QByteArray(static_cast<QXmlParseException*>(ptr)->message().toUtf8()); QtXml_PackedString { const_cast<char*>(ta2e963->prepend("WHITESPACE").constData()+10), ta2e963->size()-10, ta2e963 }; });
}

struct QtXml_PackedString QXmlParseException_PublicId(void* ptr)
{
	return ({ QByteArray* t653e96 = new QByteArray(static_cast<QXmlParseException*>(ptr)->publicId().toUtf8()); QtXml_PackedString { const_cast<char*>(t653e96->prepend("WHITESPACE").constData()+10), t653e96->size()-10, t653e96 }; });
}

struct QtXml_PackedString QXmlParseException_SystemId(void* ptr)
{
	return ({ QByteArray* t3cb8ec = new QByteArray(static_cast<QXmlParseException*>(ptr)->systemId().toUtf8()); QtXml_PackedString { const_cast<char*>(t3cb8ec->prepend("WHITESPACE").constData()+10), t3cb8ec->size()-10, t3cb8ec }; });
}

void QXmlParseException_DestroyQXmlParseException(void* ptr)
{
	static_cast<QXmlParseException*>(ptr)->~QXmlParseException();
}

class MyQXmlReader: public QXmlReader
{
public:
	QXmlDTDHandler * DTDHandler() const { return static_cast<QXmlDTDHandler*>(callbackQXmlReader_DTDHandler(const_cast<void*>(static_cast<const void*>(this)))); };
	QXmlContentHandler * contentHandler() const { return static_cast<QXmlContentHandler*>(callbackQXmlReader_ContentHandler(const_cast<void*>(static_cast<const void*>(this)))); };
	QXmlDeclHandler * declHandler() const { return static_cast<QXmlDeclHandler*>(callbackQXmlReader_DeclHandler(const_cast<void*>(static_cast<const void*>(this)))); };
	QXmlEntityResolver * entityResolver() const { return static_cast<QXmlEntityResolver*>(callbackQXmlReader_EntityResolver(const_cast<void*>(static_cast<const void*>(this)))); };
	QXmlErrorHandler * errorHandler() const { return static_cast<QXmlErrorHandler*>(callbackQXmlReader_ErrorHandler(const_cast<void*>(static_cast<const void*>(this)))); };
	bool feature(const QString & name, bool * ok) const { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlReader_Feature(const_cast<void*>(static_cast<const void*>(this)), namePacked, reinterpret_cast<char*>(ok)) != 0; };
	bool hasFeature(const QString & name) const { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlReader_HasFeature(const_cast<void*>(static_cast<const void*>(this)), namePacked) != 0; };
	bool hasProperty(const QString & name) const { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlReader_HasProperty(const_cast<void*>(static_cast<const void*>(this)), namePacked) != 0; };
	QXmlLexicalHandler * lexicalHandler() const { return static_cast<QXmlLexicalHandler*>(callbackQXmlReader_LexicalHandler(const_cast<void*>(static_cast<const void*>(this)))); };
	void * property(const QString & name, bool * ok) const { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlReader_Property(const_cast<void*>(static_cast<const void*>(this)), namePacked, reinterpret_cast<char*>(ok)); };
	void setContentHandler(QXmlContentHandler * handler) { callbackQXmlReader_SetContentHandler(this, handler); };
	void setDTDHandler(QXmlDTDHandler * handler) { callbackQXmlReader_SetDTDHandler(this, handler); };
	void setDeclHandler(QXmlDeclHandler * handler) { callbackQXmlReader_SetDeclHandler(this, handler); };
	void setEntityResolver(QXmlEntityResolver * handler) { callbackQXmlReader_SetEntityResolver(this, handler); };
	void setErrorHandler(QXmlErrorHandler * handler) { callbackQXmlReader_SetErrorHandler(this, handler); };
	void setFeature(const QString & name, bool value) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQXmlReader_SetFeature(this, namePacked, value); };
	void setLexicalHandler(QXmlLexicalHandler * handler) { callbackQXmlReader_SetLexicalHandler(this, handler); };
	void setProperty(const QString & name, void * value) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQXmlReader_SetProperty(this, namePacked, value); };
	 ~MyQXmlReader() { callbackQXmlReader_DestroyQXmlReader(this); };
};

Q_DECLARE_METATYPE(QXmlReader*)
Q_DECLARE_METATYPE(MyQXmlReader*)

int QXmlReader_QXmlReader_QRegisterMetaType(){qRegisterMetaType<QXmlReader*>(); return qRegisterMetaType<MyQXmlReader*>();}

void* QXmlReader_DTDHandler(void* ptr)
{
	return static_cast<QXmlReader*>(ptr)->DTDHandler();
}

void* QXmlReader_ContentHandler(void* ptr)
{
	return static_cast<QXmlReader*>(ptr)->contentHandler();
}

void* QXmlReader_DeclHandler(void* ptr)
{
	return static_cast<QXmlReader*>(ptr)->declHandler();
}

void* QXmlReader_EntityResolver(void* ptr)
{
	return static_cast<QXmlReader*>(ptr)->entityResolver();
}

void* QXmlReader_ErrorHandler(void* ptr)
{
	return static_cast<QXmlReader*>(ptr)->errorHandler();
}

char QXmlReader_Feature(void* ptr, struct QtXml_PackedString name, char* ok)
{
	return static_cast<QXmlReader*>(ptr)->feature(QString::fromUtf8(name.data, name.len), reinterpret_cast<bool*>(ok));
}

char QXmlReader_HasFeature(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QXmlReader*>(ptr)->hasFeature(QString::fromUtf8(name.data, name.len));
}

char QXmlReader_HasProperty(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QXmlReader*>(ptr)->hasProperty(QString::fromUtf8(name.data, name.len));
}

void* QXmlReader_LexicalHandler(void* ptr)
{
	return static_cast<QXmlReader*>(ptr)->lexicalHandler();
}

void* QXmlReader_Property(void* ptr, struct QtXml_PackedString name, char* ok)
{
	return static_cast<QXmlReader*>(ptr)->property(QString::fromUtf8(name.data, name.len), reinterpret_cast<bool*>(ok));
}

void QXmlReader_SetContentHandler(void* ptr, void* handler)
{
	static_cast<QXmlReader*>(ptr)->setContentHandler(static_cast<QXmlContentHandler*>(handler));
}

void QXmlReader_SetDTDHandler(void* ptr, void* handler)
{
	static_cast<QXmlReader*>(ptr)->setDTDHandler(static_cast<QXmlDTDHandler*>(handler));
}

void QXmlReader_SetDeclHandler(void* ptr, void* handler)
{
	static_cast<QXmlReader*>(ptr)->setDeclHandler(static_cast<QXmlDeclHandler*>(handler));
}

void QXmlReader_SetEntityResolver(void* ptr, void* handler)
{
	static_cast<QXmlReader*>(ptr)->setEntityResolver(static_cast<QXmlEntityResolver*>(handler));
}

void QXmlReader_SetErrorHandler(void* ptr, void* handler)
{
	static_cast<QXmlReader*>(ptr)->setErrorHandler(static_cast<QXmlErrorHandler*>(handler));
}

void QXmlReader_SetFeature(void* ptr, struct QtXml_PackedString name, char value)
{
	static_cast<QXmlReader*>(ptr)->setFeature(QString::fromUtf8(name.data, name.len), value != 0);
}

void QXmlReader_SetLexicalHandler(void* ptr, void* handler)
{
	static_cast<QXmlReader*>(ptr)->setLexicalHandler(static_cast<QXmlLexicalHandler*>(handler));
}

void QXmlReader_SetProperty(void* ptr, struct QtXml_PackedString name, void* value)
{
	static_cast<QXmlReader*>(ptr)->setProperty(QString::fromUtf8(name.data, name.len), value);
}

void QXmlReader_DestroyQXmlReader(void* ptr)
{
	static_cast<QXmlReader*>(ptr)->~QXmlReader();
}

void QXmlReader_DestroyQXmlReaderDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXmlSimpleReader: public QXmlSimpleReader
{
public:
	MyQXmlSimpleReader() : QXmlSimpleReader() {QXmlSimpleReader_QXmlSimpleReader_QRegisterMetaType();};
	QXmlDTDHandler * DTDHandler() const { return static_cast<QXmlDTDHandler*>(callbackQXmlSimpleReader_DTDHandler(const_cast<void*>(static_cast<const void*>(this)))); };
	QXmlContentHandler * contentHandler() const { return static_cast<QXmlContentHandler*>(callbackQXmlSimpleReader_ContentHandler(const_cast<void*>(static_cast<const void*>(this)))); };
	QXmlDeclHandler * declHandler() const { return static_cast<QXmlDeclHandler*>(callbackQXmlSimpleReader_DeclHandler(const_cast<void*>(static_cast<const void*>(this)))); };
	QXmlEntityResolver * entityResolver() const { return static_cast<QXmlEntityResolver*>(callbackQXmlSimpleReader_EntityResolver(const_cast<void*>(static_cast<const void*>(this)))); };
	QXmlErrorHandler * errorHandler() const { return static_cast<QXmlErrorHandler*>(callbackQXmlSimpleReader_ErrorHandler(const_cast<void*>(static_cast<const void*>(this)))); };
	bool feature(const QString & name, bool * ok) const { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlSimpleReader_Feature(const_cast<void*>(static_cast<const void*>(this)), namePacked, reinterpret_cast<char*>(ok)) != 0; };
	bool hasFeature(const QString & name) const { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlSimpleReader_HasFeature(const_cast<void*>(static_cast<const void*>(this)), namePacked) != 0; };
	bool hasProperty(const QString & name) const { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlSimpleReader_HasProperty(const_cast<void*>(static_cast<const void*>(this)), namePacked) != 0; };
	QXmlLexicalHandler * lexicalHandler() const { return static_cast<QXmlLexicalHandler*>(callbackQXmlSimpleReader_LexicalHandler(const_cast<void*>(static_cast<const void*>(this)))); };
	bool parse(const QXmlInputSource & input) { return callbackQXmlSimpleReader_Parse(this, const_cast<QXmlInputSource*>(&input)) != 0; };
	bool parse(const QXmlInputSource * input) { return callbackQXmlSimpleReader_Parse2(this, const_cast<QXmlInputSource*>(input)) != 0; };
	bool parse(const QXmlInputSource * input, bool incremental) { return callbackQXmlSimpleReader_Parse3(this, const_cast<QXmlInputSource*>(input), incremental) != 0; };
	bool parseContinue() { return callbackQXmlSimpleReader_ParseContinue(this) != 0; };
	void * property(const QString & name, bool * ok) const { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };return callbackQXmlSimpleReader_Property(const_cast<void*>(static_cast<const void*>(this)), namePacked, reinterpret_cast<char*>(ok)); };
	void setContentHandler(QXmlContentHandler * handler) { callbackQXmlSimpleReader_SetContentHandler(this, handler); };
	void setDTDHandler(QXmlDTDHandler * handler) { callbackQXmlSimpleReader_SetDTDHandler(this, handler); };
	void setDeclHandler(QXmlDeclHandler * handler) { callbackQXmlSimpleReader_SetDeclHandler(this, handler); };
	void setEntityResolver(QXmlEntityResolver * handler) { callbackQXmlSimpleReader_SetEntityResolver(this, handler); };
	void setErrorHandler(QXmlErrorHandler * handler) { callbackQXmlSimpleReader_SetErrorHandler(this, handler); };
	void setFeature(const QString & name, bool enable) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQXmlSimpleReader_SetFeature(this, namePacked, enable); };
	void setLexicalHandler(QXmlLexicalHandler * handler) { callbackQXmlSimpleReader_SetLexicalHandler(this, handler); };
	void setProperty(const QString & name, void * value) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtXml_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQXmlSimpleReader_SetProperty(this, namePacked, value); };
	 ~MyQXmlSimpleReader() { callbackQXmlSimpleReader_DestroyQXmlSimpleReader(this); };
};

Q_DECLARE_METATYPE(QXmlSimpleReader*)
Q_DECLARE_METATYPE(MyQXmlSimpleReader*)

int QXmlSimpleReader_QXmlSimpleReader_QRegisterMetaType(){qRegisterMetaType<QXmlSimpleReader*>(); return qRegisterMetaType<MyQXmlSimpleReader*>();}

void* QXmlSimpleReader_DTDHandler(void* ptr)
{
	return static_cast<QXmlSimpleReader*>(ptr)->DTDHandler();
}

void* QXmlSimpleReader_DTDHandlerDefault(void* ptr)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::DTDHandler();
}

void* QXmlSimpleReader_NewQXmlSimpleReader()
{
	return new MyQXmlSimpleReader();
}

void* QXmlSimpleReader_ContentHandler(void* ptr)
{
	return static_cast<QXmlSimpleReader*>(ptr)->contentHandler();
}

void* QXmlSimpleReader_ContentHandlerDefault(void* ptr)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::contentHandler();
}

void* QXmlSimpleReader_DeclHandler(void* ptr)
{
	return static_cast<QXmlSimpleReader*>(ptr)->declHandler();
}

void* QXmlSimpleReader_DeclHandlerDefault(void* ptr)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::declHandler();
}

void* QXmlSimpleReader_EntityResolver(void* ptr)
{
	return static_cast<QXmlSimpleReader*>(ptr)->entityResolver();
}

void* QXmlSimpleReader_EntityResolverDefault(void* ptr)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::entityResolver();
}

void* QXmlSimpleReader_ErrorHandler(void* ptr)
{
	return static_cast<QXmlSimpleReader*>(ptr)->errorHandler();
}

void* QXmlSimpleReader_ErrorHandlerDefault(void* ptr)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::errorHandler();
}

char QXmlSimpleReader_Feature(void* ptr, struct QtXml_PackedString name, char* ok)
{
	return static_cast<QXmlSimpleReader*>(ptr)->feature(QString::fromUtf8(name.data, name.len), reinterpret_cast<bool*>(ok));
}

char QXmlSimpleReader_FeatureDefault(void* ptr, struct QtXml_PackedString name, char* ok)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::feature(QString::fromUtf8(name.data, name.len), reinterpret_cast<bool*>(ok));
}

char QXmlSimpleReader_HasFeature(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QXmlSimpleReader*>(ptr)->hasFeature(QString::fromUtf8(name.data, name.len));
}

char QXmlSimpleReader_HasFeatureDefault(void* ptr, struct QtXml_PackedString name)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::hasFeature(QString::fromUtf8(name.data, name.len));
}

char QXmlSimpleReader_HasProperty(void* ptr, struct QtXml_PackedString name)
{
	return static_cast<QXmlSimpleReader*>(ptr)->hasProperty(QString::fromUtf8(name.data, name.len));
}

char QXmlSimpleReader_HasPropertyDefault(void* ptr, struct QtXml_PackedString name)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::hasProperty(QString::fromUtf8(name.data, name.len));
}

void* QXmlSimpleReader_LexicalHandler(void* ptr)
{
	return static_cast<QXmlSimpleReader*>(ptr)->lexicalHandler();
}

void* QXmlSimpleReader_LexicalHandlerDefault(void* ptr)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::lexicalHandler();
}

char QXmlSimpleReader_Parse(void* ptr, void* input)
{
	return static_cast<QXmlSimpleReader*>(ptr)->parse(*static_cast<QXmlInputSource*>(input));
}

char QXmlSimpleReader_ParseDefault(void* ptr, void* input)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::parse(*static_cast<QXmlInputSource*>(input));
}

char QXmlSimpleReader_Parse2(void* ptr, void* input)
{
	return static_cast<QXmlSimpleReader*>(ptr)->parse(static_cast<QXmlInputSource*>(input));
}

char QXmlSimpleReader_Parse2Default(void* ptr, void* input)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::parse(static_cast<QXmlInputSource*>(input));
}

char QXmlSimpleReader_Parse3(void* ptr, void* input, char incremental)
{
	return static_cast<QXmlSimpleReader*>(ptr)->parse(static_cast<QXmlInputSource*>(input), incremental != 0);
}

char QXmlSimpleReader_Parse3Default(void* ptr, void* input, char incremental)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::parse(static_cast<QXmlInputSource*>(input), incremental != 0);
}

char QXmlSimpleReader_ParseContinue(void* ptr)
{
	return static_cast<QXmlSimpleReader*>(ptr)->parseContinue();
}

char QXmlSimpleReader_ParseContinueDefault(void* ptr)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::parseContinue();
}

void* QXmlSimpleReader_Property(void* ptr, struct QtXml_PackedString name, char* ok)
{
	return static_cast<QXmlSimpleReader*>(ptr)->property(QString::fromUtf8(name.data, name.len), reinterpret_cast<bool*>(ok));
}

void* QXmlSimpleReader_PropertyDefault(void* ptr, struct QtXml_PackedString name, char* ok)
{
		return static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::property(QString::fromUtf8(name.data, name.len), reinterpret_cast<bool*>(ok));
}

void QXmlSimpleReader_SetContentHandler(void* ptr, void* handler)
{
	static_cast<QXmlSimpleReader*>(ptr)->setContentHandler(static_cast<QXmlContentHandler*>(handler));
}

void QXmlSimpleReader_SetContentHandlerDefault(void* ptr, void* handler)
{
		static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setContentHandler(static_cast<QXmlContentHandler*>(handler));
}

void QXmlSimpleReader_SetDTDHandler(void* ptr, void* handler)
{
	static_cast<QXmlSimpleReader*>(ptr)->setDTDHandler(static_cast<QXmlDTDHandler*>(handler));
}

void QXmlSimpleReader_SetDTDHandlerDefault(void* ptr, void* handler)
{
		static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setDTDHandler(static_cast<QXmlDTDHandler*>(handler));
}

void QXmlSimpleReader_SetDeclHandler(void* ptr, void* handler)
{
	static_cast<QXmlSimpleReader*>(ptr)->setDeclHandler(static_cast<QXmlDeclHandler*>(handler));
}

void QXmlSimpleReader_SetDeclHandlerDefault(void* ptr, void* handler)
{
		static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setDeclHandler(static_cast<QXmlDeclHandler*>(handler));
}

void QXmlSimpleReader_SetEntityResolver(void* ptr, void* handler)
{
	static_cast<QXmlSimpleReader*>(ptr)->setEntityResolver(static_cast<QXmlEntityResolver*>(handler));
}

void QXmlSimpleReader_SetEntityResolverDefault(void* ptr, void* handler)
{
		static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setEntityResolver(static_cast<QXmlEntityResolver*>(handler));
}

void QXmlSimpleReader_SetErrorHandler(void* ptr, void* handler)
{
	static_cast<QXmlSimpleReader*>(ptr)->setErrorHandler(static_cast<QXmlErrorHandler*>(handler));
}

void QXmlSimpleReader_SetErrorHandlerDefault(void* ptr, void* handler)
{
		static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setErrorHandler(static_cast<QXmlErrorHandler*>(handler));
}

void QXmlSimpleReader_SetFeature(void* ptr, struct QtXml_PackedString name, char enable)
{
	static_cast<QXmlSimpleReader*>(ptr)->setFeature(QString::fromUtf8(name.data, name.len), enable != 0);
}

void QXmlSimpleReader_SetFeatureDefault(void* ptr, struct QtXml_PackedString name, char enable)
{
		static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setFeature(QString::fromUtf8(name.data, name.len), enable != 0);
}

void QXmlSimpleReader_SetLexicalHandler(void* ptr, void* handler)
{
	static_cast<QXmlSimpleReader*>(ptr)->setLexicalHandler(static_cast<QXmlLexicalHandler*>(handler));
}

void QXmlSimpleReader_SetLexicalHandlerDefault(void* ptr, void* handler)
{
		static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setLexicalHandler(static_cast<QXmlLexicalHandler*>(handler));
}

void QXmlSimpleReader_SetProperty(void* ptr, struct QtXml_PackedString name, void* value)
{
	static_cast<QXmlSimpleReader*>(ptr)->setProperty(QString::fromUtf8(name.data, name.len), value);
}

void QXmlSimpleReader_SetPropertyDefault(void* ptr, struct QtXml_PackedString name, void* value)
{
		static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setProperty(QString::fromUtf8(name.data, name.len), value);
}

void QXmlSimpleReader_DestroyQXmlSimpleReader(void* ptr)
{
	static_cast<QXmlSimpleReader*>(ptr)->~QXmlSimpleReader();
}

void QXmlSimpleReader_DestroyQXmlSimpleReaderDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

