package main

import (
	"bytes"
	"encoding/gob"
)

type qtextFormat struct {
	Html   string
	Images map[string][]byte
}

func (s *myWindow) getRichText() []byte {
	var data qtextFormat
	data.Html = s.editor.ToHtml()
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
	res := new(qtextFormat)
	err := decoder.Decode(res)
	return res, err
}
