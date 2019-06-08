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
	db       *sql.DB
	category int
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
		"create table diaries (id integer unique,cdate text,title text,filename text,mtime text,category int default 0);",
		"create index idx1 on diaries(cdate);",
		"create table categories(id int,name text);"}
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
	res := &myDb{db: db}
	res.CheckAndUpgrade()
	return res, nil
}

func (s *myDb) setCategory(c int) {
	s.category = c
}

func (s *myDb) CheckAndUpgrade() error {
	diaryStruct, err := s.db.Query("PRAGMA table_info(diaries);")
	if err != nil {
		return err
	}
	var rows = 0
	for diaryStruct.Next() {
		rows++
	}
	diaryStruct.Close()
	if rows == 5 {
		s.db.Exec("alter table diaries add column category int default 0;")
		s.db.Exec("create table categories(id int,name text);")
	}
	return nil
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
	_, err := s.db.Exec("insert into diaries(id,cdate,title,filename,mtime,category) values(?,?,?,?,?,?);",
		id, cdate, title, filename, time.Now().Format("2006-01-02 15:04:05"), s.category)
	if err != nil {
		log.Println("UpdateDiaryTitle")
		return s.UpdateDiaryTitle(id, title)
	}
	return nil
}

func (s *myDb) AddDiary2(id int, cdate, title, filename string, category int) error {
	_, err := s.db.Exec("insert into diaries(id,cdate,title,filename,mtime,category) values(?,?,?,?,?,?);",
		id, cdate, title, filename, time.Now().Format("2006-01-02 15:04:05"), category)
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

func (s *myDb) UpdateDiaryCategory(id, c int) error {
	res, err := s.db.Exec("update diaries set category=?,mtime=? where id=?", c, time.Now().Format("2006-01-02 15:04:05"), id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return errors.New("UpdateDiaryCategory fail.")
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
	rows, err := s.db.Query("select distinct substr(cdate,0,8) from diaries where category=? order by substr(cdate,0,8) desc;", s.category)
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
	rows, err := s.db.Query("select id,substr(cdate,9,11),title,mtime from diaries where substr(cdate,0,8)=? and category=? order by substr(cdate,9,11) desc;", ym, s.category)
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

func (s *myDb) SearchTitle(kw string) ([]diaryItem, error) {
	rows, err := s.db.Query("select id,cdate,title,mtime from diaries where instr(title,?)>0 and category=? order by cdate desc;", kw, s.category)
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

func (s *myDb) AddCategory(name string) int {
	res := s.db.QueryRow("select id from categories order by id desc limit 1;")

	var id int
	err := res.Scan(&id)
	if err != nil {
		id = 1
	} else {
		id++
	}

	s.db.Exec("insert into categories(id,name) values (?,?);", id, name)
	return id
}

func (s *myDb) RenameCategory(id int, name string) error {
	_, err := s.db.Exec("update categories set name=? where id=?;", name, id)
	return err
}

func (s *myDb) RemoveCategory(c int) {
	s.db.Exec("delete from categories where id=?;", c)
	s.db.Exec("update diaries set category=0 where category=?;", c)
}

func (s *myDb) GetCategories() map[int]string {
	res := make(map[int]string)

	rows, err := s.db.Query("select id,name from categories order by id asc;")
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if rows.Scan(&id, &name) == nil {
			res[id] = name
		}
	}

	return res
}
