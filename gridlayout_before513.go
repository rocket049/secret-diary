//+build before513

//兼容 Qt5.12 之前的版本
package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type gridLayout struct {
	AddWidget2 func(widgets.QWidget_ITF, int, int, core.Qt__AlignmentFlag)
	widgets.QGridLayout
}

func newGridLayout(parent widgets.QWidget_ITF) *gridLayout {
	res := &gridLayout{AddWidget2: nil, QGridLayout: *widgets.NewQGridLayout(parent)}
	res.AddWidget2 = res.AddWidget
	return res
}

func newGridLayout2() *gridLayout {
	res := &gridLayout{AddWidget2: nil, QGridLayout: *widgets.NewQGridLayout2()}
	res.AddWidget2 = res.AddWidget
	return res
}

func newQPixMap(fn string, format string, flag core.Qt__ImageConversionFlag) *gui.QPixmap {
	return gui.NewQPixmap5(fn, format, flag)
}
