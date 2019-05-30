package main

import (
	"bytes"
	"encoding/gob"
	"reflect"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

type attachmentFile struct {
	Name string
	Data []byte
}

type qtextFormat struct {
	Html        string
	Images      map[string][]byte
	Attachments []attachmentFile
}

func (s *myWindow) getRichText() []byte {
	var data qtextFormat
	data.Html = s.editor.ToHtml()
	data.Images = s.getImagesMap(data.Html)
	data.Attachments = s.document.Attachments

	buf := bytes.NewBufferString("")
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(data)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func (s *myWindow) getQText(data []byte) (*qtextFormat, error) {
	buf := bytes.NewReader(data)
	decoder := gob.NewDecoder(buf)
	s.document.Images = make(map[string][]byte)
	s.document.Attachments = []attachmentFile{}
	err := decoder.Decode(&s.document)
	if err == nil {
		s.loadImgResources(s.document.Images)
	}
	return &s.document, err
}

func (s *myWindow) getImagesMap(html string) map[string][]byte {
	urls := s.getImageList(html)
	imgMap := make(map[string][]byte)
	for _, url := range urls {
		buf, ok := s.document.Images[url]
		if ok {
			imgMap[url] = buf
		}
	}
	return imgMap
}

func (s *myWindow) loadImgResources(imgs map[string][]byte) {
	for url, data := range imgs {
		img := gui.NewQImage()
		ok := img.LoadFromData(data, len(data), "")
		if !ok {
			return
		}
		uri := core.NewQUrl3(url, core.QUrl__TolerantMode)

		s.editor.Document().AddResource(int(gui.QTextDocument__ImageResource), uri, img.ToVariant())
	}
}

func (s *myWindow) removeCurAttachment() {
	idx := s.comboAttachs.CurrentIndex()
	if idx == 0 {
		return
	}
	attachs := sliceValueRemoveIndex(s.document.Attachments, idx-1)
	s.document.Attachments = attachs.([]attachmentFile)
	s.showAttachList()
	s.editor.Document().SetModified(true)
	//curDiary.Modified = true
}

func sliceValueRemoveIndex(v interface{}, idx int) interface{} {
	val := reflect.ValueOf(v)
	if val.Type().Kind() != reflect.Slice {
		return v
	}
	length := val.Len()
	if length == 0 {
		return v
	}
	if idx >= length || idx < 0 {
		return v
	}

	res := reflect.MakeSlice(val.Type(), length-1, length-1)
	for i := 0; i < idx; i++ {
		res.Index(i).Set(val.Index(i))
	}
	for i := idx + 1; i < length; i++ {
		res.Index(i - 1).Set(val.Index(i))
	}
	return res.Interface()
}
