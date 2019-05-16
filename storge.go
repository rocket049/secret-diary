package main

import (
	"bytes"
	"encoding/gob"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

type qtextFormat struct {
	Html   string
	Images map[string][]byte
}

func (s *myWindow) getRichText() []byte {
	var data qtextFormat
	data.Html = s.editor.ToHtml()
	data.Images = s.getImagesMap(data.Html)

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
