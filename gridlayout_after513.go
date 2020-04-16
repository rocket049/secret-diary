//+build !before513

//兼容 Qt5.13 以后的版本
package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type gridLayout struct {
	widgets.QGridLayout
}

func newGridLayout(parent widgets.QWidget_ITF) *gridLayout {
	return &gridLayout{*widgets.NewQGridLayout(parent)}
}

func newGridLayout2() *gridLayout {
	return &gridLayout{*widgets.NewQGridLayout2()}
}

func newQPixMap(fn string, format string, flag core.Qt__ImageConversionFlag) *gui.QPixmap {
	return gui.NewQPixmap3(fn, format, flag)
}
