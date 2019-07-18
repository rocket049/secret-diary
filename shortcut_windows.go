package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	locale "github.com/rocket049/go-locale"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func makeShortcut(force bool) error {
	exe1, _ := os.Executable()
	appdir := filepath.Dir(exe1)
	appname := filepath.Join(appdir, "secret-diary.exe")
	script := filepath.Join(appdir, "shortcut.vbs")
	bat := filepath.Join(appdir, "shortcut.bat")
	_, err := os.Lstat(bat)
	if err == nil && force == false {
		return err
	}

	writeBat(bat, "wscript", script, fmt.Sprintf(`/target:"%s"`, appname), "\r\nexit\r\n")

	open.Start(bat)

	return nil
}

func writeBat(bat string, msgs ...interface{}) {
	var buf = []byte(fmt.Sprintln(msgs...))
	var err error = nil

	loc, _ := locale.DetectLocale()
	switch loc {
	case "zh_CN":
		enc := simplifiedchinese.GB18030.NewEncoder()
		buf, err = enc.Bytes(buf)
	}

	if err == nil {
		ioutil.WriteFile(bat, buf, 0644)
	}

}
