package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func (s *myWindow) setHeader(level int) {
	var cfmt = gui.NewQTextCharFormat()
	switch level {
	case 1:
		cfmt.SetFontPointSize(18)
	case 2:
		cfmt.SetFontPointSize(16)
	case 3:
		cfmt.SetFontPointSize(14)
	default:
		cfmt.SetFontPointSize(12)
	}

	cfmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__blue), core.Qt__SolidPattern))
	s.mergeFormatOnLineOrSelection(cfmt)
}

func (s *myWindow) textStyle(styleIndex int) {
	var cursor = s.editor.TextCursor()

	if styleIndex != 0 {

		var style = gui.QTextListFormat__ListDisc

		switch styleIndex {
		case 1:
			{
				style = gui.QTextListFormat__ListDisc
			}

		case 2:
			{
				style = gui.QTextListFormat__ListCircle
			}

		case 3:
			{
				style = gui.QTextListFormat__ListSquare
			}

		case 4:
			{
				style = gui.QTextListFormat__ListDecimal
			}

		case 5:
			{
				style = gui.QTextListFormat__ListLowerAlpha
			}

		case 6:
			{
				style = gui.QTextListFormat__ListUpperAlpha
			}

		case 7:
			{
				style = gui.QTextListFormat__ListLowerRoman
			}

		case 8:
			{
				style = gui.QTextListFormat__ListUpperRoman
			}
		}

		cursor.BeginEditBlock()

		var (
			blockFmt = cursor.BlockFormat()
			listFmt  = gui.NewQTextListFormat()
		)

		if cursor.CurrentList().Pointer() != nil {
			listFmt = gui.NewQTextListFormatFromPointer(cursor.CurrentList().Format().Pointer())
		} else {
			listFmt.SetIndent(blockFmt.Indent() + 1)
			blockFmt.SetIndent(0)
			cursor.SetBlockFormat(blockFmt)
		}

		listFmt.SetStyle(style)
		cursor.CreateList(listFmt)

		cursor.EndEditBlock()

	} else {
		var bfmt = gui.NewQTextBlockFormat()
		bfmt.SetObjectIndex(-1)
		cursor.MergeBlockFormat(bfmt)
	}
}

func (s *myWindow) textColor() {
	var col = widgets.QColorDialog_GetColor(s.editor.TextColor(), s.editor, "", 0)
	if !col.IsValid() {
		return
	}
	var cfmt = gui.NewQTextCharFormat()
	cfmt.SetForeground(gui.NewQBrush3(col, core.Qt__SolidPattern))
	s.mergeFormatOnLineOrSelection(cfmt)
}

func (s *myWindow) textBold() {
	var afmt = gui.NewQTextCharFormat()
	var fw = gui.QFont__Normal
	if s.actionTextBold.IsChecked() {
		fw = gui.QFont__Bold
	}
	afmt.SetFontWeight(int(fw))
	s.mergeFormatOnLineOrSelection(afmt)
}

func (s *myWindow) textUnderline() {
	var afmt = gui.NewQTextCharFormat()
	afmt.SetFontUnderline(s.actionTextUnderline.IsChecked())
	s.mergeFormatOnLineOrSelection(afmt)
}

func (s *myWindow) textItalic() {
	var afmt = gui.NewQTextCharFormat()
	afmt.SetFontItalic(s.actionTextItalic.IsChecked())
	s.mergeFormatOnLineOrSelection(afmt)
}
