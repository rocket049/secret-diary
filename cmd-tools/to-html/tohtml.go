package main

// Run program in directory $HOME/.sdiary/$USERNAME/
// It will generate index.html

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	_ "embed"
)

//go:embed index.tpl
var tplIndex string

func main() {
	fmt.Print("Password:")
	var pwd string
	n, err := fmt.Scan(&pwd)
	if err != nil || n == 0 {
		panic(err.Error())
	}
	fmt.Println(pwd)
	db, err := openMyDb("diary.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	key, err := db.GetRealKey(pwd)
	if err != nil || n == 0 {
		panic(err.Error())
	}
	fmt.Printf("%x\n", key)

	diaries, err := db.ListAllDiary()
	if err != nil {
		panic(err)
	}
	for _, diary := range diaries {
		fmt.Printf("%v\n%v\n%v %v\n", diary.Id, diary.Title, diary.Day, diary.MTime)
		//转换为html
		convertFile(fmt.Sprintf("%v", diary.Id), key)
	}

	t := template.New("")
	t.Parse(tplIndex)
	fp, err := os.Create("index.html")
	if err != nil {
		panic(err.Error())
	}
	defer fp.Close()
	err = t.Execute(fp, diaries)
	if err != nil {
		panic(err.Error())
	}
}

type attachmentFile struct {
	Name string
	Data []byte
}

type qtextFormat struct {
	Html        string
	Images      map[string][]byte
	Attachments []attachmentFile
}

func convertFile(name string, key []byte) error {
	data, err := decodeFromFile(name+".dat", key)
	if err != nil {
		return err
	}

	buf := bytes.NewReader(data)

	decoder := gob.NewDecoder(buf)

	var document qtextFormat
	document.Images = make(map[string][]byte)
	document.Attachments = []attachmentFile{}

	err = decoder.Decode(&document)
	if err != nil {
		return err
	}
	os.MkdirAll(name, os.ModePerm)
	htmlImages := ""
	for k, img := range document.Images {
		fn := path.Base(k)
		ioutil.WriteFile(filepath.Join(name, fn), img, 0644)
		htmlImages = htmlImages + fmt.Sprintf("\n<p><img src='%v' /></p>", fn)
		//println(htmlImages)
	}

	htmlAttach := ""
	for _, item := range document.Attachments {
		ioutil.WriteFile(filepath.Join(name, item.Name), item.Data, 0644)
		htmlAttach = htmlAttach + fmt.Sprintf("\n<p><a href='%v' />%v</a></p>", item.Name, item.Name)
		//println(htmlAttach)
	}

	pos := strings.Index(document.Html, "</body>")

	html := document.Html[:pos] + htmlImages + htmlAttach + document.Html[pos:]

	ioutil.WriteFile(filepath.Join(name, name+".html"), []byte(html), 0644)
	return err
}
