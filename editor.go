package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func (s *myWindow) setHeader(level int) {
	var cfmt = gui.NewQTextCharFormat()
	cfmt.SetFont2(s.app.Font())
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

func (s *myWindow) setStandard() {
	var cfmt = gui.NewQTextCharFormat()
	cfmt.SetFont2(s.app.Font())
	cfmt.SetFontPointSize(14)
	cfmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__black), core.Qt__SolidPattern))

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
		cursor.SetBlockFormat(bfmt)
	}
}

// addIndent n=1 or n=-1
func (s *myWindow) addIndent(n int) {
	cursor := s.editor.TextCursor()
	cursor.BeginEditBlock()

	var (
		blockFmt = cursor.BlockFormat()
		listFmt  = gui.NewQTextListFormat()
	)

	if cursor.CurrentList().Pointer() != nil {
		listFmt = gui.NewQTextListFormatFromPointer(cursor.CurrentList().Format().Pointer())
		if listFmt.Indent()+n >= 0 {
			listFmt.SetIndent(listFmt.Indent() + n)
			cursor.CreateList(listFmt)
		}

	} else {
		if blockFmt.Indent()+n >= 0 {
			blockFmt.SetIndent(blockFmt.Indent() + n)
			cursor.SetBlockFormat(blockFmt)
		}

	}

	cursor.EndEditBlock()
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

func (s *myWindow) textBgColor() {
	var col = widgets.QColorDialog_GetColor(s.editor.TextColor(), s.editor, "", 0)
	if !col.IsValid() {
		return
	}
	var cfmt = gui.NewQTextCharFormat()
	cfmt.SetBackground(gui.NewQBrush3(col, core.Qt__SolidPattern))
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

func (s *myWindow) textStrikeOut() {
	var afmt = gui.NewQTextCharFormat()
	afmt.SetFontStrikeOut(s.actionStrikeOut.IsChecked())
	s.mergeFormatOnLineOrSelection(afmt)
}

func (s *myWindow) textItalic() {
	var afmt = gui.NewQTextCharFormat()
	afmt.SetFontItalic(s.actionTextItalic.IsChecked())
	s.mergeFormatOnLineOrSelection(afmt)
}

func (s *myWindow) insertImage() {
	filename := widgets.QFileDialog_GetOpenFileName(s.window, "select a file", ".", "Image (*.png *.jpg)", "Image (*.png *.jpg)", widgets.QFileDialog__ReadOnly)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	img := gui.NewQImage()
	ok := img.LoadFromData(data, len(data), "")
	if !ok {
		return
	}

	uri := core.NewQUrl3("rc://"+filename, core.QUrl__TolerantMode)

	img = s.scaleImage(img)

	s.editor.Document().AddResource(int(gui.QTextDocument__ImageResource), uri, img.ToVariant())
	url := uri.Url(core.QUrl__None)
	cursor := s.editor.TextCursor()
	cursor.InsertImage4(img, url)

	ba := core.NewQByteArray()

	iod := core.NewQBuffer2(ba, nil)

	iod.Open(core.QIODevice__WriteOnly)

	ok = img.Save2(iod, filepath.Ext(filename)[1:], -1)
	//fmt.Println(filepath.Ext(filename))
	if ok {
		s.document.Images[url] = []byte(ba.Data())
	}

	//fmt.Println("save image:", ok)
}

func (s *myWindow) scaleImage(src *gui.QImage) (res *gui.QImage) {

	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetWindowTitle(T("Scale Image Size"))

	grid := widgets.NewQGridLayout(dlg)

	width := widgets.NewQLabel2(fmt.Sprintf("%s : %d =>", T("Width"), src.Width()), dlg, core.Qt__Widget)
	grid.AddWidget(width, 0, 0, 0)
	scaledW := src.Width()
	scaledH := src.Height()
	delta := 30
	if scaledW > s.editor.Width()-delta {
		scaledW = s.editor.Geometry().Width() - delta
		scaledH = int(float64(src.Height()) * float64(scaledW) / float64(src.Width()))
	}
	wValidor := gui.NewQIntValidator(dlg)
	wValidor.SetRange(10, scaledW)
	hValidor := gui.NewQIntValidator(dlg)
	hValidor.SetRange(10, scaledH)

	widthInput := widgets.NewQLineEdit(dlg)
	widthInput.SetText(strconv.Itoa(scaledW))
	widthInput.SetValidator(wValidor)
	grid.AddWidget(widthInput, 0, 1, 0)

	height := widgets.NewQLabel2(fmt.Sprintf("%s : %d =>", T("Height"), src.Height()), dlg, core.Qt__Widget)

	grid.AddWidget(height, 1, 0, 0)

	heightInput := widgets.NewQLineEdit(dlg)
	heightInput.SetText(strconv.Itoa(scaledH))
	heightInput.SetValidator(hValidor)
	grid.AddWidget(heightInput, 1, 1, 0)

	btb := widgets.NewQGridLayout(nil)

	okBtn := widgets.NewQPushButton2(T("OK"), dlg)
	btb.AddWidget(okBtn, 0, 0, 0)

	cancelBtn := widgets.NewQPushButton2(T("Cancel"), dlg)
	btb.AddWidget(cancelBtn, 0, 1, 0)

	grid.AddLayout2(btb, 2, 0, 1, 2, 0)

	dlg.SetLayout(grid)

	widthInput.ConnectKeyReleaseEvent(func(e *gui.QKeyEvent) {
		w, err := strconv.Atoi(widthInput.Text())
		if err != nil {
			return
		}
		w0 := float64(src.Width())
		h0 := float64(src.Height())
		h := float64(w) * h0 / w0
		heightInput.SetText(strconv.Itoa(int(h)))
	})
	heightInput.ConnectKeyReleaseEvent(func(e *gui.QKeyEvent) {
		h, err := strconv.Atoi(heightInput.Text())
		if err != nil {
			return
		}
		w0 := float64(src.Width())
		h0 := float64(src.Height())
		w := float64(h) * w0 / h0
		widthInput.SetText(strconv.Itoa(int(w)))
	})

	okBtn.ConnectClicked(func(b bool) {
		w, err := strconv.Atoi(widthInput.Text())
		if err != nil {
			res = src
		}
		h, err := strconv.Atoi(heightInput.Text())
		if err != nil {
			res = src
		}
		res = src.Scaled2(w, h, core.Qt__KeepAspectRatioByExpanding, core.Qt__SmoothTransformation)
		dlg.Hide()
		dlg.Destroy(true, true)
	})

	cancelBtn.ConnectClicked(func(b bool) {
		res = src
		dlg.Hide()
		dlg.Destroy(true, true)
	})

	dlg.Exec()
	if res == nil {
		res = src
	}
	return
}

func (s *myWindow) getImageList(html string) []string {
	r := strings.NewReader(html)
	bufr := bufio.NewReader(r)
	reg1, err := regexp.Compile(`<img[^><]+/>`)
	if err != nil {
		//fmt.Println(err)
		return nil
	}
	reg2, err := regexp.Compile(`src="([^"]+)"`)
	if err != nil {
		//fmt.Println(err)
		return nil
	}
	imgs := []string{}
	for line, _, err := bufr.ReadLine(); err == nil; line, _, err = bufr.ReadLine() {
		line1 := string(line)
		res1 := reg1.FindAllString(line1, -1)
		imgs = append(imgs, res1...)
	}

	res := []string{}
	for _, img := range imgs {
		res2 := reg2.FindStringSubmatch(img)
		res = append(res, res2[1])
	}
	//fmt.Println(res)
	return res
}

func (s *myWindow) insertTable() {
	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetWindowTitle(T("Table Rows and Columns"))

	grid := widgets.NewQGridLayout(dlg)

	row := widgets.NewQLabel2(T("Rows:"), dlg, core.Qt__Widget)
	grid.AddWidget(row, 0, 0, 0)

	rowInput := widgets.NewQLineEdit(dlg)
	rowInput.SetText("3")
	rowInput.SetValidator(gui.NewQIntValidator(dlg))
	grid.AddWidget(rowInput, 0, 1, 0)

	col := widgets.NewQLabel2(T("Columns:"), dlg, core.Qt__Widget)

	grid.AddWidget(col, 1, 0, 0)

	colInput := widgets.NewQLineEdit(dlg)
	colInput.SetText("3")
	colInput.SetValidator(gui.NewQIntValidator(dlg))
	grid.AddWidget(colInput, 1, 1, 0)

	btb := widgets.NewQGridLayout(nil)

	okBtn := widgets.NewQPushButton2(T("OK"), dlg)
	btb.AddWidget(okBtn, 0, 0, 0)

	cancelBtn := widgets.NewQPushButton2(T("Cancel"), dlg)
	btb.AddWidget(cancelBtn, 0, 1, 0)

	grid.AddLayout2(btb, 2, 0, 1, 2, 0)

	dlg.SetLayout(grid)

	okBtn.ConnectClicked(func(b bool) {
		cursor := s.editor.TextCursor()
		r, err := strconv.Atoi(rowInput.Text())
		if err != nil {
			return
		}
		c, err := strconv.Atoi(colInput.Text())
		if err != nil {
			return
		}
		tbl := cursor.InsertTable2(r, c)
		tbl.Format().SetBorderBrush(gui.NewQBrush2(core.Qt__SolidPattern))
		dlg.Hide()
		dlg.Destroy(true, true)
	})

	cancelBtn.ConnectClicked(func(b bool) {
		dlg.Hide()
		dlg.Destroy(true, true)
	})

	dlg.SetModal(true)
	dlg.Show()
}

func (s *myWindow) addJustifyActions(tb *widgets.QToolBar) {
	rsrcPath := ":/qml/icons"
	var leftIcon = gui.QIcon_FromTheme2("format-justify-left", gui.NewQIcon5(rsrcPath+"/textleft.png"))
	actionAlignLeft := tb.AddAction2(leftIcon, "&Left")
	actionAlignLeft.SetPriority(widgets.QAction__LowPriority)
	actionAlignLeft.ConnectTriggered(func(b bool) {
		s.textAlign(1)
	})

	var centerIcon = gui.QIcon_FromTheme2("format-justify-center", gui.NewQIcon5(rsrcPath+"/textcenter.png"))
	actionAlignCenter := tb.AddAction2(centerIcon, "C&enter")
	actionAlignCenter.SetPriority(widgets.QAction__LowPriority)
	actionAlignCenter.ConnectTriggered(func(b bool) {
		s.textAlign(2)
	})

	var rightIcon = gui.QIcon_FromTheme2("format-justify-right", gui.NewQIcon5(rsrcPath+"/textright.png"))
	actionAlignRight := tb.AddAction2(rightIcon, "&Right")
	actionAlignRight.SetPriority(widgets.QAction__LowPriority)
	actionAlignRight.ConnectTriggered(func(b bool) {
		s.textAlign(3)
	})

	var fillIcon = gui.QIcon_FromTheme2("format-justify-fill", gui.NewQIcon5(rsrcPath+"/textjustify.png"))
	actionAlignJustify := tb.AddAction2(fillIcon, "&Justify")
	actionAlignJustify.SetPriority(widgets.QAction__LowPriority)
	actionAlignJustify.ConnectTriggered(func(b bool) {
		s.textAlign(4)
	})
}

func (s *myWindow) textAlign(n int) {
	switch n {
	case 1:

		s.editor.SetAlignment(core.Qt__AlignLeft | core.Qt__AlignAbsolute)

	case 2:

		s.editor.SetAlignment(core.Qt__AlignHCenter)

	case 3:

		s.editor.SetAlignment(core.Qt__AlignRight | core.Qt__AlignAbsolute)

	case 4:

		s.editor.SetAlignment(core.Qt__AlignJustify)

	}
}

func (s *myWindow) getTable() (t *gui.QTextTable, cell *gui.QTextTableCell) {
	cursor := s.editor.TextCursor()
	blk := s.editor.Document().FindBlock(cursor.Position())
	for _, frame := range blk.Document().RootFrame().ChildFrames() {
		//fmt.Println("table cell:", frame.FrameFormat().IsTableCellFormat(), "table:", frame.FrameFormat().IsTableFormat())
		if frame.FrameFormat().IsTableFormat() {
			table := gui.NewQTextTableFromPointer(frame.Pointer())
			cell := table.CellAt2(cursor.Position())
			//fmt.Println(cell.Row(), cell.Column())
			return table, cell
		}
	}

	return nil, nil
}

func (s *myWindow) clearFormatAtCursor() {
	cursor := s.editor.TextCursor()

	block := cursor.Block()
	cfmt := block.CharFormat()
	cfmt.ClearBackground()
	cfmt.ClearForeground()
	cfmt.SetFontUnderline(false)
	cfmt.SetFontWeight(int(gui.QFont__Normal))
	cfmt.SetFontPointSize(14)
	cfmt.SetFontItalic(false)
	cfmt.SetFontStrikeOut(false)

	var bfmt = gui.NewQTextBlockFormat()
	bfmt.SetObjectIndex(-1)

	if !cursor.HasSelection() {
		cursor.Select(gui.QTextCursor__LineUnderCursor)
	}
	cursor.SetCharFormat(cfmt)
	cursor.SetBlockFormat(bfmt)
	s.editor.SetCurrentCharFormat(cfmt)
}

func OpenDiaryNewWindow(parent *myWindow, id int) *myWindow {
	win := new(myWindow)
	win.OpenNewWindow(parent, id)
	return win
}

func (s *myWindow) charWidth() int {
	// font := gui.NewQFont()
	// font.SetFamily("Serif")
	// if runtime.GOOS == "windows" {
	// 	font.SetPointSize(16)
	// } else {
	// 	//font.SetFamily("Serif Regular")
	// 	font.SetPointSize(12)
	// }

	// fm := gui.NewQFontMetrics(font)
	// return fm.BoundingRect2("M").Height()
	return 24
}

func (s *myWindow) OpenNewWindow(parent *myWindow, id int) {
	s.key = parent.key
	s.curDiary.Id = id
	s.db = parent.db

	s.window = widgets.NewQMainWindow(parent.window, core.Qt__Window)

	s.window.SetWindowTitle(parent.user)
	s.window.SetMinimumSize2(800, 600)
	s.window.SetWindowIcon(gui.NewQIcon5(":/qml/icons/Sd.png"))

	grid := widgets.NewQGridLayout2()

	frame := widgets.NewQFrame(s.window, core.Qt__Widget)

	s.window.SetCentralWidget(frame)

	charW := s.charWidth()
	editor := s.createEditor(charW)
	s.window.SetMinimumWidth(s.editor.Width() + 100)

	grid.AddWidget3(editor, 0, 0, 1, 1, 0)

	comboBox := s.setupComboAttachs()
	grid.AddLayout(comboBox, 1, 0, 0)

	grid.SetAlign(core.Qt__AlignTop)

	s.setToolBar()

	s.setMenuBar()

	s.exportEnc.SetEnabled(true)
	s.exportPdf.SetEnabled(true)
	s.importEnc.SetDisabled(true)
	s.newDiary.SetDisabled(true)
	s.renDiary.SetDisabled(true)
	s.modifyPwd.SetDisabled(true)

	//s.setStandaloneFuncs()
	s.setEditorFuncs()

	frame.SetLayout(grid)

	once := sync.Once{}
	s.window.ConnectShowEvent(func(e *gui.QShowEvent) {
		once.Do(func() {
			s.loadDiary()
		})
	})

	s.window.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		if s.editor.Document().IsModified() {
			ret := widgets.QMessageBox_Question(s.window, T("Close"), T("Do you want to save the document?"), widgets.QMessageBox__Yes|widgets.QMessageBox__No, widgets.QMessageBox__Yes)
			if ret == widgets.QMessageBox__Yes {
				s.saveCurDiary()
			}
		}

	})
	s.window.Show()
}

func (s *myWindow) loadDiary() {
	filename := strconv.Itoa(s.curDiary.Id) + ".dat"
	data, err := decodeFromFile(filename, s.key)
	if err != nil {
		s.setStatusBar(err.Error())
	} else {
		s.getQText(data)
		s.editor.Document().SetHtml(s.document.Html)

		s.showAttachList()
		s.editor.Document().SetModified(false)
		s.editor.SetReadOnly(false)
	}
}

func (s *myWindow) saveDiaryAlone() {
	if s.editor.Document().IsModified() == false {
		s.setStatusBar(T("No Diary Saved"))
		return
	}
	filename := strconv.Itoa(s.curDiary.Id) + ".dat"
	encodeToFile(s.getRichText(), filename, s.key)

	title := strings.TrimSpace(s.editor.Document().FirstBlock().Text())
	s.db.UpdateDiaryTitle(s.curDiary.Id, title)

	s.setStatusBar(T("Save Diary") + fmt.Sprintf(" %s(%s)", title, filename))
}

func (s *myWindow) searchFromDb(kw string) {
	if len(kw) == 0 {
		s.setStatusBar(T("Must Input Search Keyword!"))
		return
	}
	s.setStatusBar(T("Loading Diary List..."))
	diaryList, err := s.db.SearchTitle(kw)
	if err != nil {
		s.setStatusBar(err.Error())
		return
	}

	s.clearModelFind()

	var ym string
	var r, c int
	for _, diary := range diaryList {
		if diary.Day[:7] != ym {
			ym = diary.Day[:7]
			item := s.addFindYM(ym)
			idx := item.Index()
			r, c = idx.Row(), idx.Column()
			s.treeFind.Expand(idx)
		}
		item := s.modelFind.Item(r, c)
		child := gui.NewQStandardItem2(fmt.Sprintf("%s-%s", diary.Day[8:], diary.Title))
		child.SetEditable(false)
		child.SetAccessibleText(strconv.Itoa(diary.Id))
		child.SetAccessibleDescription("0")
		child.SetToolTip(T("Double Click to Open. ") + T("Last Modified:") + diary.MTime)

		item.AppendRow2(child)
	}
	s.treeFind.ResizeColumnToContents(0)
	return
}

func (s *myWindow) addFindYM(yearMonth string) *gui.QStandardItem {
	idx := core.NewQModelIndex()
	p := s.modelFind.Parent(idx)
	n := s.modelFind.RowCount(p)
	for i := 0; i < n; i++ {
		item := s.modelFind.Item(i, 0)
		if item.Text() == yearMonth {
			return item
		}
	}
	item := gui.NewQStandardItem2(yearMonth)
	item.SetColumnCount(1)
	item.SetEditable(false)

	item.SetAccessibleText("1")
	item.SetAccessibleDescription("1")

	s.modelFind.AppendRow2(item)

	return item
}

func (s *myWindow) setTreeFindFuncs() {
	s.treeFind.ConnectActivated(func(idx *core.QModelIndex) {
		item := s.modelFind.ItemFromIndex(idx)
		if item.AccessibleDescription() == "1" {
			return
		}
		id, err := strconv.Atoi(item.AccessibleText())
		if err != nil {
			s.setStatusBar(err.Error())
		}
		OpenDiaryNewWindow(s, id)
	})
}

func (s *myWindow) findText() {
	if s.searchInput != nil {
		cursor := s.editor.TextCursor()
		s.searchInput.SetText(cursor.Selection().ToPlainText())
		return
	}
	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetMinimumWidth(s.editor.Width() / 2)

	dlg.SetWindowTitle(T("Text Search"))

	grid := widgets.NewQGridLayout(dlg)

	word := widgets.NewQLineEdit(dlg)
	word.SetPlaceholderText(T("Words to search."))
	grid.AddWidget3(word, 0, 0, 1, 2, 0)
	s.searchInput = word
	cursor := s.editor.TextCursor()
	s.searchInput.SetText(cursor.Selection().ToPlainText())

	backBtn := widgets.NewQPushButton2(T("Last"), dlg)
	grid.AddWidget(backBtn, 1, 0, 0)

	nextBtn := widgets.NewQPushButton2(T("Next"), dlg)
	grid.AddWidget(nextBtn, 1, 1, 0)

	dlg.SetLayout(grid)

	backBtn.ConnectClicked(func(b bool) {
		kw := strings.TrimSpace(word.Text())
		if len(kw) == 0 {
			return
		}
		cursor := s.editor.TextCursor()
		if len(cursor.Selection().ToPlainText()) > 0 {
			cursor.SetPosition(cursor.SelectionStart(), gui.QTextCursor__MoveAnchor)
			s.editor.SetTextCursor(cursor)
		}
		s.editor.Find(kw, gui.QTextDocument__FindBackward)
	})

	nextBtn.ConnectClicked(func(b bool) {
		kw := strings.TrimSpace(word.Text())
		if len(kw) == 0 {
			return
		}
		cursor := s.editor.TextCursor()
		if len(cursor.Selection().ToPlainText()) > 0 {
			cursor.SetPosition(cursor.SelectionEnd(), gui.QTextCursor__MoveAnchor)
			s.editor.SetTextCursor(cursor)
		}
		s.editor.Find(kw, 0)
	})

	dlg.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		dlg.Destroy(true, true)
		s.searchInput = nil
	})

	dlg.Show()
}

func (s *myWindow) replaceText() {
	if s.replaceInput != nil {
		cursor := s.editor.TextCursor()
		s.replaceInput.SetText(cursor.Selection().ToPlainText())
		return
	}

	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetMinimumWidth(s.editor.Width() / 2)

	dlg.SetWindowTitle(T("Text Replace"))

	grid := widgets.NewQGridLayout(dlg)

	wordOld := widgets.NewQLineEdit(dlg)
	wordOld.SetPlaceholderText(T("Old Text."))
	wordOld.SetToolTip(T("Old Text."))
	grid.AddWidget3(wordOld, 0, 0, 1, 3, 0)
	s.replaceInput = wordOld
	cursor := s.editor.TextCursor()
	s.replaceInput.SetText(cursor.Selection().ToPlainText())

	wordNew := widgets.NewQLineEdit(dlg)
	wordNew.SetPlaceholderText(T("New Text."))
	wordNew.SetToolTip(T("New Text."))
	grid.AddWidget3(wordNew, 1, 0, 1, 3, 0)

	backBtn := widgets.NewQPushButton2(T("Last"), dlg)
	grid.AddWidget(backBtn, 2, 0, 0)

	nextBtn := widgets.NewQPushButton2(T("Next"), dlg)
	grid.AddWidget(nextBtn, 2, 1, 0)

	allBtn := widgets.NewQPushButton2(T("All"), dlg)
	grid.AddWidget(allBtn, 2, 2, 0)

	dlg.SetLayout(grid)

	backBtn.ConnectClicked(func(b bool) {
		word0 := wordOld.Text()
		word1 := wordNew.Text()
		if len(word0) == 0 {
			return
		}
		cursor := s.editor.TextCursor()
		if cursor.Selection().ToPlainText() == word0 {
			cursor.RemoveSelectedText()
			cursor.InsertText(word1)

			cursor.SetPosition(cursor.Position()-len(word1), gui.QTextCursor__MoveAnchor)
			s.editor.SetTextCursor(cursor)
		}

		s.editor.Find(word0, gui.QTextDocument__FindBackward)
	})

	nextBtn.ConnectClicked(func(b bool) {
		word0 := wordOld.Text()
		word1 := wordNew.Text()
		if len(word0) == 0 {
			return
		}
		cursor := s.editor.TextCursor()

		if cursor.Selection().ToPlainText() == word0 {
			cursor.RemoveSelectedText()
			cursor.InsertText(word1)
		}
		s.editor.Find(word0, 0)
	})

	allBtn.ConnectClicked(func(b bool) {
		word0 := wordOld.Text()
		if len(word0) == 0 {
			return
		}
		text := s.editor.ToPlainText()
		n := strings.Count(strings.ToLower(text), strings.ToLower(word0))

		cursor := s.editor.TextCursor()
		cursor.SetPosition(0, gui.QTextCursor__MoveAnchor)
		s.editor.SetTextCursor(cursor)

		for i := 0; i < n+1; i++ {
			nextBtn.Clicked(true)
		}
	})

	dlg.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		dlg.Destroy(true, true)
		s.replaceInput = nil
	})

	dlg.Show()
}

func (s *myWindow) pasteText() {
	board := gui.QGuiApplication_Clipboard()
	text := board.Text(gui.QClipboard__Clipboard)
	cursor := s.editor.TextCursor()
	cursor.InsertText(text)
}
