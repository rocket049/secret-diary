//+build before513
//兼容 Qt5.12 之前的版本
package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type gridLayout struct {
	widgets.QGridLayout
	AddWidget2 func(widgets.QWidget_ITF, int, int, core.Qt__AlignmentFlag)
}

func newGridLayout(parent widgets.QWidget_ITF) *gridLayout {
	res := &gridLayout{*widgets.NewQGridLayout(parent)}
	res.AddWidget2 = res.AddWidget
	return res
}

func newGridLayout2() *gridLayout {
	res := &gridLayout{*widgets.NewQGridLayout2()}
	res.AddWidget2 = res.AddWidget
	return res
}
