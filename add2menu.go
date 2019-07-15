package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func copyFile(src, dst string) error {
	fp1, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fp1.Close()
	fp2, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fp2.Close()
	_, err = io.Copy(fp1, fp2)
	log.Println("copy", src, dst)
	return err
}

func makeLauncher(launcherName, fileName, iconName string, force bool) {
	_, err := os.Lstat("/usr/local/share/applications/Secret-Diary.desktop")
	if err == nil {
		//install from deb
		return
	}
	exe1, _ := os.Executable()
	appdir := filepath.Dir(exe1)
	appname := filepath.Join(appdir, fileName)
	home, _ := os.UserHomeDir()
	os.MkdirAll(filepath.Join(home, ".local", "share", "applications"), os.ModePerm)
	dst := filepath.Join(home, ".local", "share", "applications", launcherName+".desktop")
	iconSrc := filepath.Join(appdir, iconName)

	if isNewer(dst, iconSrc) && force == false {
		//already exist
		return
	}

	data := struct {
		LauncherName string
		Name         string
		Icon         string
	}{launcherName, appname, iconSrc}

	tpl := `[Desktop Entry]
Name={{.LauncherName}}
Comment={{.LauncherName}}
Exec="{{.Name}}" %U
Icon={{.Icon}}
Terminal=false
Type=Application
StartupNotify=true
Categories=Office;
	
`
	t := template.New("")
	t.Parse(tpl)
	fp, err := os.Create(dst)
	if err != nil {
		log.Println(err)
		return
	}
	defer fp.Close()
	t.Execute(fp, data)
}

func isNewer(fileNew, fileOld string) bool {
	infoNew, err := os.Lstat(fileNew)
	if err != nil {
		return false
	}
	infoOld, err := os.Lstat(fileOld)
	if err != nil {
		return true
	}
	if infoNew.ModTime().Unix() > infoOld.ModTime().Unix() {
		return true
	}
	return false
}

func appimageLauncher(force bool) error {
	desktop := "Secret-Diary.desktop"
	icon := "Sd.png"

	appimage := os.Getenv("APPIMAGE")
	appdir := os.Getenv("APPDIR")
	if len(appimage) == 0 || len(appdir) == 0 {
		return errors.New("Not Appimage")
	}

	home, _ := os.UserHomeDir()

	src := filepath.Join(appdir, desktop)
	dst := filepath.Join(home, ".local", "share", "applications", desktop)
	if isNewer(dst, appimage) && force == false {
		return nil
	}

	iconSrc := filepath.Join(appdir, icon)

	iconDir := filepath.Join(home, ".local", "share", "icons")
	iconDst := filepath.Join(iconDir, icon)

	os.MkdirAll(iconDir, os.ModePerm)

	err := copyFile(iconSrc, iconDst)
	if err != nil {
		log.Println(err)
	}

	srcFp, err := os.Open(src)
	if err != nil {
		log.Println(err)
		return err
	}
	defer srcFp.Close()
	reader := bufio.NewReader(srcFp)

	fp, err := os.Create(dst)
	if err != nil {
		log.Println(err)
		return err
	}
	defer fp.Close()
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if bytes.HasPrefix(line, []byte("Exec=")) {
			fp.WriteString("Exec=" + appimage + "\n")
		} else {
			fp.Write(line)
			fp.WriteString("\n")
		}
	}
	return nil
}

func addToMenu() {
	err := appimageLauncher(true)
	if err != nil {
		makeLauncher("Secret-Diary", "secret-diary", "Sd.png", true)
	}
}
