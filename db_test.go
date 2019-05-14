package main

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func TestAddDiary(t *testing.T) {
	home, _ := os.UserHomeDir()
	path1 := path.Join(home, ".sdiary", "test", "diary.db")
	os.Remove(path1)
	err := createUserDb("test", "1234")
	if err != nil {
		t.Fatal(err)
	}
	db, err := getMyDb("test")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	tt, err := time.Parse("20060102", "19790101")
	if err != nil {
		t.Fatal(err)
	}
	stop := time.Now().Unix()
	id := db.NextId()
	for tt.Unix() < stop {

		cdate := tt.Format("2006-01-02")
		title := fmt.Sprintf("diary %d", id)
		filename := fmt.Sprintf("%d.dat", id)
		db.AddDiary(id, cdate, title, filename)
		tt = tt.Add(time.Hour * 24)
		id++
	}
}

func TestUpdatePwd(t *testing.T) {
	home, _ := os.UserHomeDir()
	path1 := path.Join(home, ".sdiary", "test2", "diary.db")
	os.Remove(path1)
	err := createUserDb("test2", "1234")
	if err != nil {
		t.Fatal(err)
	}
	db, err := getMyDb("test2")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.UpdateRealKeyAndSha40("1234", "12345")
	_, err = db.GetRealKey("12345")
	if err != nil {
		t.Fatal(err)
	}
}
