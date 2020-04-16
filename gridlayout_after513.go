//+build !before513
//兼容 Qt5.13 以后的版本
package main

import (
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
