package main

import (
	"github.com/therecipe/qt/core"
)

type QmlBridge struct {
	core.QObject

	_ func(id, day, title, mtime string, r, c int) `signal:"addDiary"`
	_ func(r, c int)                               `signal:"setMonthFlag"`
	_ func(ym string)                              `signal:"addYearMonth"`
}
