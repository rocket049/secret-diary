package go_locale

import (
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func DetectLocale() (string, error) {
	out, err := getWinCommandOutput("wmic", "os", "get", "locale")
	if err != nil {
		return "", err
	}

	out = strings.Replace(out, "Locale", "", -1)
	out = strings.TrimSpace(out)

	id, err := strconv.ParseInt(out, 16, 64)
	if err != nil {
		return "", err
	}

	lcid := LCID()
	locale, err := lcid.ById(int(id))
	if err != nil {
		return "", err
	}

	return locale, nil
}

func getWinCommandOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
