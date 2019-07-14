package main

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func copyFile(src, dst string, mode os.FileMode) error {
	fp1, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, mode)
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
	return err
}

func makeLauncher(launcherName, fileName, iconName string) {
	exe1, _ := os.Executable()
	appdir := filepath.Dir(exe1)
	appname := filepath.Join(appdir, fileName)
	home, _ := os.UserHomeDir()
	os.MkdirAll(filepath.Join(home, ".local", "share", "applications"), os.ModePerm)
	dst := filepath.Join(home, ".local", "share", "applications", launcherName+".desktop")
	iconSrc := filepath.Join(appdir, iconName)

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
	appname := "Secret-Diary"
	icon := "Sd.png"
	category := "Office;"

	appimage := os.Getenv("APPIMAGE")
	appdir := os.Getenv("APPDIR")
	if len(appimage) == 0 || len(appdir) == 0 {
		return errors.New("Not Appimage")
	}

	home, _ := os.UserHomeDir()
	dst := filepath.Join(home, ".local", "share", "applications", appname+".desktop")
	if isNewer(dst, appimage) && force == false {
		return nil
	}

	iconSrc := filepath.Join(appdir, icon)

	iconDir := filepath.Join(home, ".local", "share", "icons", appname)
	iconDst := filepath.Join(iconDir, icon)

	os.MkdirAll(iconDir, os.ModePerm)

	copyFile(iconSrc, iconDst, 0644)

	data := struct {
		AppName  string
		Name     string
		Icon     string
		Category string
	}{appname, appimage, iconDst, category}

	tpl := `[Desktop Entry]
Name={{.AppName}}
Comment={{.AppName}}
Exec="{{.Name}}" %U
Icon={{.Icon}}
Terminal=false
Type=Application
StartupNotify=true
Categories={{.Category}};
	
`
	t := template.New("")
	t.Parse(tpl)
	fp, err := os.Create(dst)
	if err != nil {
		log.Println(err)
		return err
	}
	defer fp.Close()
	t.Execute(fp, data)
	return nil
}

func addToMenu() {
	err := appimageLauncher(true)
	if err != nil {
		makeLauncher("Secret-Diary", "secret-diary", "Sd.png")
	}
}
