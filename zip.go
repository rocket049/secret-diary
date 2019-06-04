package main

import (
	"archive/zip"
	"crypto/aes"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func zipData(fromDir, toFile string) error {
	dir, err := os.Open(fromDir)
	if err != nil {
		return err
	}
	defer dir.Close()

	//open zip
	fileToZip, err := os.Create(toFile)
	if err != nil {
		return err
	}
	defer fileToZip.Close()
	zipFile := zip.NewWriter(fileToZip)
	defer zipFile.Close()

	files, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range files {
		filename := filepath.Join(fromDir, name)
		fp, err := os.Open(filename)
		if err != nil {
			return err

		}
		//add to zip
		writer, err := zipFile.Create(name)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, fp)
		if err != nil {
			return err
		}
		fp.Close()
	}
	return nil
}

func (s *myWindow) importFromZip(filename string, pwd string) error {
	zipFile, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	//unzip diary.db
	var diaryDbName string
	for _, zf := range zipFile.File {
		if zf.Name == "diary.db" {
			diaryDbName = filepath.Join(os.TempDir(), "diary.db")
			fp, err := os.Create(diaryDbName)
			if err != nil {
				return err
			}

			reader, err := zf.Open()
			if err != nil {
				fp.Close()
				return err
			}

			_, err = io.Copy(fp, reader)

			reader.Close()
			if err != nil {
				fp.Close()
				reader.Close()
				return err
			}

			fp.Close()
			break
		}
	}
	db, err := sql.Open("sqlite3", diaryDbName)
	if err != nil {
		return err
	}
	defer os.Remove(diaryDbName)
	defer db.Close()

	res, err := db.Query("select cdate,title,filename from diaries")
	if err != nil {
		return err
	}
	//get filename title dict
	type diaryRecod struct {
		Cdate string
		Title string
	}
	var titleDict = make(map[string]diaryRecod)
	var cdate, title, fname string
	for res.Next() {
		err := res.Scan(&cdate, &title, &fname)
		if err != nil {
			res.Close()
			return err
		}
		titleDict[fname] = diaryRecod{cdate, title}
	}
	res.Close()
	//get real key
	key, err := dbGetRealKey(db, pwd)
	if err != nil {
		return err
	}
	//import every file from zip
	for _, zf := range zipFile.File {
		if zf.Name == "diary.db" {
			break
		}
		record, ok := titleDict[zf.Name]
		if !ok {
			break
		}
		reader, err := zf.Open()
		if err != nil {
			return err
		}
		data, err := decodeFromReader(reader, key)
		reader.Close()
		if err != nil {
			return err
		}

		//save diary
		id := s.db.NextId()
		datName := fmt.Sprintf("%d.dat", id)
		encodeToFile(data, datName, s.key)
		s.db.AddDiary(id, record.Cdate, record.Title, datName)
	}
	return nil
}

func dbGetRealKey(db *sql.DB, pwd string) ([]byte, error) {
	sha40 := getSha40String(pwd)
	row := db.QueryRow("select realkey from user where sha40=?;", sha40)
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
