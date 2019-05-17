package main

import (
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

func main() {
	makeLauncher("Secret-Diary", "secret-diary", "Sd.png")
}
