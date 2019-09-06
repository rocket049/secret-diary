//生成从1979-01-01 到当前日期的所有日志索引 `~/.sdiary/test/diary.db`，用于测试载入速度

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path"
	"time"
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

var dat []byte = nil
var key []byte

func createDat(home string, user string, id int) error {
	path1 := path.Join(home, ".sdiary", user, fmt.Sprintf("%d.dat", id))
	if dat == nil {
		var err error
		var data qtextFormat
		data.Html = "Test"
		data.Images = nil
		data.Attachments = nil

		buf := bytes.NewBufferString("")
		encoder := gob.NewEncoder(buf)
		err = encoder.Encode(data)
		if err != nil {
			return err
		}
		dat = buf.Bytes()
	}
	return encodeToPathName(dat, path1, key)
}

func createDiarys() {
	home, _ := os.UserHomeDir()
	path1 := path.Join(home, ".sdiary", "test", "diary.db")
	os.Remove(path1)
	err := createUserDb("test", "1234")
	if err != nil {
		panic(err.Error() + "1")
	}
	db, err := getMyDb("test")
	if err != nil {
		panic(err.Error() + "2")
	}
	defer db.Close()
	key, _ = db.GetRealKey("1234")
	tt, err := time.Parse("20060102", "19790101")
	if err != nil {
		panic(err.Error() + "3")
	}
	log.Println("start", tt.Format("2006-01-02"), tt.Unix())
	stop := time.Now().Unix()
	log.Println("end", time.Now().Format("2006-01-02"), stop)
	id := db.NextId()
	log.Println("start id:", id)
	tx, err := db.db.Begin()
	if err != nil {
		panic(err.Error() + "4")
	}
	for tt.Unix() < stop {
		cdate := tt.Format("2006-01-02")
		title := fmt.Sprintf("diary %d", id)
		filename := fmt.Sprintf("%d.dat", id)
		tx.Exec("insert into diaries(id,cdate,title,filename,mtime,category) values(?,?,?,?,?,?);",
			id, cdate, title, filename, time.Now().Format("2006-01-02 15:04:05"), db.category)
		err = createDat(home, "test", id)
		if err != nil {
			log.Println(err, " dat")
		}
		tt = tt.Add(time.Hour * 24)
		id++
	}
	err = tx.Commit()
	if err != nil {
		panic(err.Error() + "5")
	}
	log.Println("success")
}

func main() {
	createDiarys()
}
