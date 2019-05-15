package main

import (
	"crypto/aes"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"path"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type myDb struct {
	db *sql.DB
}

func createUserDb(name string, pwd string) error {
	home1, _ := os.UserHomeDir()
	dataDir = path.Join(home1, ".sdiary", name)
	os.MkdirAll(dataDir, os.ModePerm)
	userDb := path.Join(dataDir, "diary.db")
	if _, err := os.Stat(userDb); err == nil {
		//exists
		return errors.New("Already exists.")
	}
	db, err := sql.Open("sqlite3", userDb)
	if err != nil {
		return err
	}
	defer db.Close()
	sqls := []string{"create table user (id integer unique,cdata text,sha40 TEXT not null,realkey text not null,mtime text);",
		"create table diaries (id integer unique,cdate text,title text,filename text,mtime text);",
		"create index idx1 on diaries(cdate);"}
	//"create table months (ym text);"}
	for i := 0; i < len(sqls); i++ {
		_, err = db.Exec(sqls[i])
		if err != nil {
			return err
		}
	}

	id := 1
	now := time.Now()
	cdate := now.Format("2006-01-02")
	sha40 := getSha40String(pwd)
	realKey, err := getRealKeyString(pwd)
	if err != nil {
		return err
	}
	mtime := now.Format("2006-01-02 15:04:05")

	_, err = db.Exec("insert into user(id,cdata,sha40,realkey,mtime) values(?,?,?,?,?);", id, cdate, sha40, realKey, mtime)

	return err
}

func getMyDb(name string) (*myDb, error) {
	home1, _ := os.UserHomeDir()
	dataDir = path.Join(home1, ".sdiary", name)
	os.MkdirAll(dataDir, os.ModePerm)

	userDb := path.Join(dataDir, "diary.db")
	if _, err := os.Stat(userDb); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", userDb)
	if err != nil {
		return nil, err
	}
	return &myDb{db: db}, nil
}

func (s *myDb) Close() {
	s.db.Close()
}

func (s *myDb) GetRealKey(pwd string) ([]byte, error) {
	sha40 := getSha40String(pwd)
	row := s.db.QueryRow("select realkey from user where sha40=?;", sha40)
	var realKey string
	err := row.Scan(&realKey)
	if err != nil {
		return nil, err
	}
	key, err := hex.DecodeString(realKey)
	if err != nil {
		return nil, err
	}
	if len(key) != 32+aes.BlockSize {
		return nil, errors.New("Data invalid.")
	}
	iv := key[:aes.BlockSize]
	data := key[aes.BlockSize:]
	ukey := getSha4(pwd)
	return decodeRealKey(ukey, data, iv)
}

func (s *myDb) UpdateRealKeyAndSha40(pwdOld, pwdNew string) error {
	sha40 := getSha40String(pwdOld)
	row := s.db.QueryRow("select realkey from user where sha40=?;", sha40)
	var realKey string
	err := row.Scan(&realKey)
	if err != nil {
		return err
	}
	key, err := hex.DecodeString(realKey)
	if err != nil {
		return err
	}
	if len(key) != 32+aes.BlockSize {
		return errors.New("Data invalid.")
	}
	iv := key[:aes.BlockSize]
	data := key[aes.BlockSize:]
	ukey := getSha4(pwdOld)
	rkey, err := decodeRealKey(ukey, data, iv)
	if err != nil {
		return err
	}
	rkeyStr, err := newRealKeyString(rkey, pwdNew)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("update user set realkey=?,sha40=? where id=1;", rkeyStr, getSha40String(pwdNew))
	return err
}

func (s *myDb) AddDiary(id int, cdate, title, filename string) error {
	_, err := s.db.Exec("insert into diaries(id,cdate,title,filename,mtime) values(?,?,?,?,?);",
		id, cdate, title, filename, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println("UpdateDiaryTitle")
		return s.UpdateDiaryTitle(id, title)
	}
	return nil
}

func (s *myDb) UpdateDiaryTitle(id int, title string) error {
	res, err := s.db.Exec("update diaries set title=?,mtime=? where id=?", title, time.Now().Format("2006-01-02 15:04:05"), id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return errors.New("UpdateDiaryTitle fail.")
	}
	return nil
}

func (s *myDb) TouchDiary(id int) error {
	res, err := s.db.Exec("update diaries set mtime=? where id=?;", time.Now().Format("2006-01-02 15:04:05"), id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return errors.New("AddDiary fail.")
	}
	return nil
}

func (s *myDb) NextId() int {
	res := s.db.QueryRow("select id from diaries order by id desc limit 1;")

	var id int
	err := res.Scan(&id)
	if err != nil {
		log.Println(err)
		return 1
	}
	return id + 1
}

func (s *myDb) GetYearMonths() ([]string, error) {
	rows, err := s.db.Query("select distinct substr(cdate,0,8) from diaries order by substr(cdate,0,8) desc;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := []string{}
	for rows.Next() {
		var v string
		rows.Scan(&v)
		res = append(res, v)
	}
	return res, nil
}

type diaryItem struct {
	Id    int
	Day   string
	Title string
	MTime string
}

func (s *myDb) GetListFromYearMonth(ym string) ([]diaryItem, error) {
	rows, err := s.db.Query("select id,substr(cdate,9,11),title,mtime from diaries where substr(cdate,0,8)=? order by substr(cdate,9,11) desc;", ym)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := []diaryItem{}
	for rows.Next() {
		var v diaryItem
		rows.Scan(&v.Id, &v.Day, &v.Title, &v.MTime)
		res = append(res, v)
	}
	return res, nil
}

func (s *myDb) RemoveDiary(id string) error {
	row := s.db.QueryRow("select filename from diaries where id=?;", id)
	var filename string
	err := row.Scan(&filename)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("delete from diaries where id=?;", id)
	if err != nil {
		return err
	}
	return os.Remove(path.Join(dataDir, filename))
}
