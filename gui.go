package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/printsupport"

	"github.com/rocket049/gettext-go/gettext"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const version = "1.1.18"

func init() {
	exe1, _ := os.Executable()
	dir1 := path.Dir(exe1)
	locale1 := path.Join(dir1, "locale")
	gettext.BindTextdomain("sdiary", locale1, nil)
	gettext.Textdomain("sdiary")
}

type diaryPointer struct {
	Item      *gui.QStandardItem
	Id        int
	Day       string
	YearMonth string
	//Modified  bool
}

func T(v string) string {
	return gettext.T(v)
}

type formatData struct {
	BF *gui.QTextBlockFormat
	CF *gui.QTextCharFormat
}

type myWindow struct {
	app                 *widgets.QApplication
	window              *widgets.QMainWindow
	tree                *widgets.QTreeView
	model               *gui.QStandardItemModel
	treeFind            *widgets.QTreeView
	modelFind           *gui.QStandardItemModel
	editor              *widgets.QTextEdit
	comboAttachs        *widgets.QComboBox
	actionTextBold      *widgets.QAction
	actionTextUnderline *widgets.QAction
	actionTextItalic    *widgets.QAction
	actionStrikeOut     *widgets.QAction
	exportEnc           *widgets.QAction
	exportPdf           *widgets.QAction
	importEnc           *widgets.QAction
	paste               *widgets.QAction
	search              *widgets.QAction
	replace             *widgets.QAction
	lineMargin          *widgets.QAction
	searchInput         *widgets.QLineEdit
	replaceInput        *widgets.QLineEdit
	newDiary            *widgets.QAction
	saveDiary           *widgets.QAction
	renDiary            *widgets.QAction
	modifyPwd           *widgets.QAction
	fb                  *widgets.QAction
	categories          *widgets.QMenu
	category            int
	categoryName        string
	categoryItem        *widgets.QAction
	curDiary            diaryPointer
	user                string
	key                 []byte
	db                  *myDb
	bridge              *QmlBridge
	document            qtextFormat
	format              formatData
}

func (s *myWindow) getNewId() (res int) {
	return s.db.NextId()
}

func (s *myWindow) setMenuBar() {
	s.document.Images = make(map[string][]byte)

	menubar := s.window.MenuBar()
	menu := menubar.AddMenu2(T("File"))

	s.exportEnc = menu.AddAction(T("Export Encrypted Diary"))
	s.exportEnc.SetDisabled(true)
	s.exportEnc.ConnectTriggered(func(b bool) {
		s.saveCurDiary()
		s.exportEncryptedDiary()
	})

	s.exportPdf = menu.AddAction(T("Export PDF..."))
	s.exportPdf.SetDisabled(true)
	s.exportPdf.ConnectTriggered(func(b bool) {
		s.exportAsPdf()
	})

	s.importEnc = menu.AddAction(T("Import Encrypted Diary"))
	s.importEnc.ConnectTriggered(func(b bool) {
		s.importEncryptedDiary()
	})

	exportAll := menu.AddAction((T("Export All")))
	exportAll.ConnectTriggered(func(b bool) {
		home, _ := os.UserHomeDir()
		dir := widgets.QFileDialog_GetExistingDirectory(s.window, T("Select export directory"), home, 0)
		if len(dir) == 0 {
			return
		}
		fromDir := filepath.Join(home, ".sdiary", s.user)
		toFile := filepath.Join(home, s.user+".bak")
		err := zipData(fromDir, toFile)
		if err != nil {
			widgets.QMessageBox_About(s.window, T("Error"), err.Error())
		}
	})

	importAll := menu.AddAction(T("Import All"))
	importAll.ConnectTriggered(func(b bool) {
		home, _ := os.UserHomeDir()
		filename := widgets.QFileDialog_GetOpenFileName(s.window, T("Select a bakup file"), home, ".bak (*.bak)", ".bak (*.bak)", 0)
		if len(filename) == 0 {
			return
		}
		var ok bool
		pwd := widgets.QInputDialog_GetText(s.window, T("The password of the bakup file"), T("Password:"), widgets.QLineEdit__Password, "",
			&ok, core.Qt__Dialog, core.Qt__ImhNone)
		if len(filename) < 4 || !ok {
			return
		}
		err := s.importFromZip(filename, pwd)
		if err != nil {
			widgets.QMessageBox_About(s.window, T("Error"), err.Error())
		}

		s.clearModel()

		s.addYearMonthsFromDb()
	})

	menu = menubar.AddMenu2(T("Edit"))
	s.paste = menu.AddAction(T("Paste Text"))
	s.paste.SetShortcut(gui.QKeySequence_FromString("Ctrl+d", gui.QKeySequence__NativeText))
	s.paste.ConnectTriggered(func(b bool) {
		s.pasteText()
	})

	s.search = menu.AddAction(T("Search"))
	s.search.SetShortcut(gui.QKeySequence_FromString("Ctrl+f", gui.QKeySequence__NativeText))
	s.search.ConnectTriggered(func(b bool) {
		s.findText()
	})

	s.replace = menu.AddAction(T("Replace"))
	s.replace.SetShortcut(gui.QKeySequence_FromString("Ctrl+r", gui.QKeySequence__NativeText))
	s.replace.ConnectTriggered(func(b bool) {
		s.replaceText()
	})

	s.lineMargin = menu.AddAction(T("Top Margin"))
	s.lineMargin.SetShortcut(gui.QKeySequence_FromString("Ctrl+l", gui.QKeySequence__NativeText))
	s.lineMargin.ConnectTriggered(func(b bool) {
		s.changeLineMargin()
	})

	menu = menubar.AddMenu2(T("Table"))

	newTable := menu.AddAction(T("Insert Table"))
	newTable.ConnectTriggered(func(b bool) {
		s.insertTable()
	})

	addPreRow := menu.AddAction(T("Insert Row"))
	addPreRow.ConnectTriggered(func(b bool) {
		t, c := s.getTable()
		if t == nil {
			return
		}
		t.InsertRows(c.Row(), 1)
	})
	appendRow := menu.AddAction(T("Append Row"))
	appendRow.ConnectTriggered(func(b bool) {
		t, _ := s.getTable()
		if t == nil {
			return
		}
		t.AppendRows(1)
	})

	addPreCol := menu.AddAction(T("Insert Column"))
	addPreCol.ConnectTriggered(func(b bool) {
		t, c := s.getTable()
		if t == nil {
			return
		}
		t.InsertColumns(c.Column(), 1)
	})
	appendCol := menu.AddAction(T("Append Column"))
	appendCol.ConnectTriggered(func(b bool) {
		t, _ := s.getTable()
		if t == nil {
			return
		}
		t.AppendColumns(1)
	})

	removeRow := menu.AddAction(T("Remove Row"))
	removeRow.ConnectTriggered(func(b bool) {
		t, c := s.getTable()
		if t == nil {
			return
		}
		t.RemoveRows(c.Row(), 1)
	})

	removeCol := menu.AddAction(T("Remove Column"))
	removeCol.ConnectTriggered(func(b bool) {
		t, c := s.getTable()
		if t == nil {
			return
		}
		t.RemoveColumns(c.Column(), 1)
	})

	s.showCategores(menubar)

	menu = menubar.AddMenu2(T("Manage"))

	s.renDiary = menu.AddAction(T("Rename"))
	s.renDiary.ConnectTriggered(func(b bool) {
		s.rename()
	})

	s.modifyPwd = menu.AddAction(T("ModifyPassword"))
	s.modifyPwd.ConnectTriggered(func(b bool) {
		s.updatePwd()
	})

	menu = menubar.AddMenu2(T("Help"))

	about := menu.AddAction(T("About"))
	about.ConnectTriggered(func(b bool) {
		s.showMsg(T("About"), T("Copy Right: Fu Huizhong <fuhuizn@163.com>\nSecurity Diary ")+version+"\n"+T("Crypto with AES256")+"\nHomepage: https://github.com/rocket049/secret-diary")
	})

	if runtime.GOOS != "windows" {
		add2menu := menu.AddAction(T("Add To Menu"))
		add2menu.ConnectTriggered(func(b bool) {
			addToMenu()
			s.showMsg(T("Add To Menu"), T("You can start this program from 'Start'->'Office'"))
		})
	}

	bak := menu.AddAction(T("Backup & Delete"))
	bak.ConnectTriggered(func(b bool) {
		s.showMsg(T("Backup & Delete"), T("Your DATA storge path is '")+dataDir+"'\n"+T("You can backup and delete the directory yourself."))
	})

	pwd := menu.AddAction(T("Password"))
	pwd.ConnectTriggered(func(b bool) {
		s.showMsg(T("Password"), T("There is no way to recover password. \nPlease write it on paper!"))
	})

	support := menu.AddAction(T("Support Author"))
	support.ConnectTriggered(func(b bool) {
		s.showPay()
	})

}

func (s *myWindow) setToolBar() {
	bar := widgets.NewQToolBar("File", nil)
	s.newDiary = bar.AddAction(T("New"))
	s.newDiary.ConnectTriggered(func(b bool) {
		s.setStatusBar(T("New Diary"))
		now := time.Now()
		s.addDiary(now.Format("2006-01"), now.Format("02"), T("New Diary")+now.Format("2006-01-02"))
	})

	s.saveDiary = bar.AddAction(T("Save"))
	s.saveDiary.SetToolTip(T("Save") + " Ctrl-S")
	s.saveDiary.SetShortcut(gui.QKeySequence_FromString("CTRL+s", gui.QKeySequence__NativeText))

	s.window.AddToolBar2(bar)

	bar = widgets.NewQToolBar("Format", nil)

	head1 := bar.AddAction("H1")
	head1.SetToolTip(T("Header 1"))
	head1.ConnectTriggered(func(b bool) {
		s.setHeader(1)
	})
	head2 := bar.AddAction("H2")
	head2.SetToolTip(T("Header 2"))
	head2.ConnectTriggered(func(b bool) {
		s.setHeader(2)
	})
	head3 := bar.AddAction("H3")
	head3.SetToolTip(T("Header 3"))
	head3.ConnectTriggered(func(b bool) {
		s.setHeader(3)
	})

	standard := bar.AddAction(T("Standard"))
	standard.ConnectTriggered(func(b bool) {
		var cfmt = gui.NewQTextCharFormat()
		cfmt.SetFont2(s.app.Font())
		cfmt.SetFontPointSize(14)
		cfmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__black), core.Qt__SolidPattern))
		s.mergeFormatOnLineOrSelection(cfmt)
	})

	s.actionTextBold = bar.AddAction("B")
	s.actionTextBold.SetCheckable(true)
	s.actionTextBold.SetToolTip(T("Bold"))
	s.actionTextBold.ConnectTriggered(func(checked bool) {
		s.textBold()
	})

	s.actionTextItalic = bar.AddAction("I")
	s.actionTextItalic.SetCheckable(true)
	s.actionTextItalic.SetToolTip(T("Italic"))
	s.actionTextItalic.ConnectTriggered(func(checked bool) {
		s.textItalic()
	})

	s.actionTextUnderline = bar.AddAction("U")
	s.actionTextUnderline.SetCheckable(true)
	s.actionTextUnderline.SetToolTip(T("Underline"))
	s.actionTextUnderline.ConnectTriggered(func(checked bool) {
		s.textUnderline()
	})

	s.actionStrikeOut = bar.AddAction("D")
	s.actionStrikeOut.SetCheckable(true)
	s.actionStrikeOut.SetToolTip(T("StrikeOut"))
	s.actionStrikeOut.ConnectTriggered(func(checked bool) {
		s.textStrikeOut()
	})

	var icon = gui.NewQIcon5(":/qml/icons/fg.png")
	actionTextColor := bar.AddAction2(icon, T("Text Color..."))
	actionTextColor.ConnectTriggered(func(checked bool) {
		s.textColor()
	})

	icon = gui.NewQIcon5(":/qml/icons/bg.png")
	actionBgColor := bar.AddAction2(icon, T("Background Color..."))
	actionBgColor.ConnectTriggered(func(checked bool) {
		s.textBgColor()
	})

	icon = gui.NewQIcon5(":/qml/icons/draw-eraser.png")
	actionClear := bar.AddAction2(icon, T("Clear Line Format..."))
	actionClear.ConnectTriggered(func(checked bool) {
		s.clearFormatAtCursor()
	})

	s.addJustifyActions(bar)

	icon = gui.NewQIcon5(":/qml/icons/right.png")
	increaseIdent := bar.AddAction2(icon, T("Increase Ident"))
	increaseIdent.ConnectTriggered(func(checked bool) {
		s.addIndent(1)
	})

	icon = gui.NewQIcon5(":/qml/icons/left.png")
	degreeIdent := bar.AddAction2(icon, T("Decrease Ident"))
	degreeIdent.ConnectTriggered(func(checked bool) {
		s.addIndent(-1)
	})

	icon = gui.NewQIcon5(":/qml/icons/brush.png")
	s.fb = bar.AddAction2(icon, T("Format Brush"))
	s.fb.SetCheckable(true)
	s.fb.ConnectTriggered(func(checked bool) {
		cursor := s.editor.TextCursor()
		if !cursor.HasSelection() {
			cursor.Select(gui.QTextCursor__LineUnderCursor)
		}
		s.format.BF = cursor.BlockFormat()
		s.format.CF = cursor.CharFormat()
		s.fb.SetChecked(true)
	})

	comboStyle := widgets.NewQComboBox(bar)
	bar.AddWidget(comboStyle)
	comboStyle.AddItems([]string{
		T("Standard"),
		T("List (Disc)"),
		T("List (Circle)"),
		T("List (Square)"),
		T("List (Decimal)"),
		T("List (Alpha lower)"),
		T("List (Alpha upper)"),
		T("List (Roman lower)"),
		T("List (Roman upper)"),
	})
	comboStyle.ConnectActivated(s.textStyle)

	s.window.AddToolBar2(bar)

	bar = widgets.NewQToolBar("Insert", nil)

	addImage := bar.AddAction(T("Image"))
	addImage.SetToolTip(T("Insert Image"))
	addImage.ConnectTriggered(func(b bool) {
		s.insertImage()
	})

	addTable := bar.AddAction(T("Table"))
	addTable.SetToolTip(T("Insert Table"))
	addTable.ConnectTriggered(func(b bool) {
		s.insertTable()
	})

	attach := bar.AddAction(T("Attach..."))
	attach.SetToolTip(T("Append Attachments"))
	attach.ConnectTriggered(func(b bool) {
		filename := widgets.QFileDialog_GetOpenFileName(s.window, T("Seclect File"), ".", "All (*)", "All (*)", widgets.QFileDialog__ReadOnly)
		if s.addAttachment(filename) {
			s.showAttachList()
			s.editor.Document().SetModified(true)
			//curDiary.Modified = true
		}

	})

	font := bar.AddAction(T("Font"))
	font.SetToolTip(T("CurrentFont"))
	font.ConnectTriggered(func(b bool) {
		var ok bool
		f := widgets.QFontDialog_GetFont2(&ok, s.window)
		if ok {
			cfmt := gui.NewQTextCharFormat()
			cfmt.SetFont2(f)
			s.mergeFormatOnLineOrSelection(cfmt)
		}
	})

	s.window.AddToolBar2(bar)
}

func (s *myWindow) setStatusBar(msg string) {
	bar := s.window.StatusBar()
	bar.ShowMessage(msg, 20000)
}

func (s *myWindow) setupComboAttachs() *widgets.QHBoxLayout {
	hbox := widgets.NewQHBoxLayout()

	s.comboAttachs = widgets.NewQComboBox(s.window)
	s.comboAttachs.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Fixed)
	list := []string{fmt.Sprintf("-- %s(%s:0) --", T("Select Attachments"), T("Total"))}
	s.comboAttachs.AddItems(list)
	expBtn := widgets.NewQPushButton2(T("Export As..."), s.window)
	expBtn.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)
	hbox.AddWidget(s.comboAttachs, 1, 0)
	hbox.AddWidget(expBtn, 1, 0)

	delBtn := widgets.NewQPushButton2(T("Remove"), s.window)
	delBtn.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)
	hbox.AddWidget(delBtn, 1, 0)

	expBtn.ConnectClicked(func(b bool) {
		idx := s.comboAttachs.CurrentIndex()
		if idx == 0 {
			return
		}
		idx -= 1
		dlg := widgets.NewQFileDialog(s.window, core.Qt__Dialog)
		filename := dlg.GetSaveFileName(s.window, T("Save File"), s.document.Attachments[idx].Name, "All (*)", "All (*)", 0)
		if len(filename) > 0 {
			ioutil.WriteFile(filename, s.document.Attachments[idx].Data, 0644)
		}

	})

	delBtn.ConnectClicked(func(b bool) {
		if s.comboAttachs.CurrentIndex() == 0 {
			return
		}
		ret := widgets.QMessageBox_Question(s.window, T("Remove"), T("Are you sure?"), widgets.QMessageBox__Yes|widgets.QMessageBox__No, widgets.QMessageBox__Yes)
		if ret == widgets.QMessageBox__Yes {
			s.removeCurAttachment()
		}

	})

	return hbox
}

func (s *myWindow) showAttachList() {
	list := []string{fmt.Sprintf("-- %s(%s:%d) --", T("Select Attachments"), T("Total"), len(s.document.Attachments))}
	for _, v := range s.document.Attachments {
		list = append(list, v.Name)
	}
	s.comboAttachs.Clear()
	s.comboAttachs.AddItems(list)
}

func (s *myWindow) clearAttachs() {
	s.comboAttachs.Clear()
	list := []string{fmt.Sprintf("-- %s(%s:0) --", T("Select Attachments"), T("Total"))}
	s.comboAttachs.AddItems(list)
}

func (s *myWindow) addAttachment(filename string) bool {
	name := filepath.Base(filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return false
	}
	s.document.Attachments = append(s.document.Attachments, attachmentFile{Name: name, Data: data})
	return true
}

func (s *myWindow) Create(app *widgets.QApplication) {

	font := gui.NewQFont()
	font.SetPointSize(12)
	font.SetFamily("Serif")
	font.SetStyleName("Regular")
	app.SetFont(font, "serif-regular")
	s.app = app

	charW := s.charWidth()

	s.window = widgets.NewQMainWindow(nil, core.Qt__Window)

	s.window.SetWindowTitle(T("UserName"))

	s.window.SetWindowIcon(gui.NewQIcon5(":/qml/icons/Sd.png"))

	spliter := widgets.NewQSplitter2(core.Qt__Horizontal, s.window)

	s.window.SetCentralWidget(spliter)

	leftArea := s.createLeftArea(charW)

	spliter.AddWidget(leftArea)

	editor := s.createEditor(charW)

	spliter.AddWidget(editor)

	spliter.SetStretchFactor(0, 1)
	spliter.SetStretchFactor(1, 2)

	s.setToolBar()

	s.setMenuBar()

	app.SetActiveWindow(s.window)
	app.SetApplicationDisplayName(T("Secret Diary"))
	app.SetApplicationName("sdiary")
	app.SetApplicationVersion("1.0.1")

	s.setEditorFuncs()
	s.setTreeFuncs()

	once := sync.Once{}
	s.window.ConnectShowEvent(func(e *gui.QShowEvent) {
		once.Do(func() {
			err := appimageLauncher(false)
			if err != nil {
				log.Println(err)
			}
			s.login()
		})
	})

	s.window.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		if s.editor.Document().IsModified() {
			ret := widgets.QMessageBox_Question(s.window, T("Close"), T("Do you want to save the document?"), widgets.QMessageBox__Yes|widgets.QMessageBox__No, widgets.QMessageBox__Yes)
			if ret == widgets.QMessageBox__Yes {
				s.saveCurDiary()
			}
		}

		s.db.Close()
	})
	s.window.ShowMaximized()
}

func (s *myWindow) createLeftArea(charW int) widgets.QWidget_ITF {
	spliter := widgets.NewQSplitter(nil)
	spliter.SetOrientation(core.Qt__Vertical)
	spliter.SetSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Expanding)

	//mwidth := charW * 18

	s.tree = widgets.NewQTreeView(s.window)
	//s.tree.SetMaximumWidth(mwidth)
	s.tree.SetSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Expanding)
	s.tree.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAsNeeded)
	s.tree.SetAutoScroll(true)
	s.model = gui.NewQStandardItemModel2(0, 1, s.tree)
	s.clearModel()

	s.tree.SetModel(s.model)
	spliter.AddWidget(s.tree)

	search := widgets.NewQWidget(nil, core.Qt__Widget)
	grid := widgets.NewQGridLayout(search)
	keyword := widgets.NewQLineEdit(search)
	grid.AddWidget(keyword, 0, 0, 0)

	btn := widgets.NewQPushButton2(T("Search"), search)
	grid.AddWidget(btn, 0, 1, 0)
	btn.ConnectClicked(func(b bool) {
		s.searchFromDb(strings.TrimSpace(keyword.Text()))
	})
	keyword.ConnectEditingFinished(func() {
		s.searchFromDb(strings.TrimSpace(keyword.Text()))
	})

	s.treeFind = widgets.NewQTreeView(s.window)

	s.treeFind.SetSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Preferred)
	s.treeFind.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAsNeeded)
	s.treeFind.SetAutoScroll(true)
	s.modelFind = gui.NewQStandardItemModel2(0, 1, s.treeFind)
	s.clearModelFind()
	s.treeFind.SetModel(s.modelFind)
	grid.AddWidget3(s.treeFind, 1, 0, 1, 2, 0)
	search.SetLayout(grid)

	spliter.AddWidget(search)

	s.setTreeFindFuncs()
	return spliter
}

func (s *myWindow) createEditor(charW int) widgets.QWidget_ITF {
	editorLayout := widgets.NewQVBoxLayout()

	scrollarea := widgets.NewQScrollArea(s.window)
	scrollarea.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAsNeeded)
	scrollarea.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAsNeeded)
	scrollarea.SetAlignment(core.Qt__AlignCenter)
	scrollarea.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	frame := widgets.NewQFrame(s.window, core.Qt__Widget)
	grid := widgets.NewQGridLayout2()
	s.editor = widgets.NewQTextEdit(s.window)
	s.editor.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Expanding)
	s.editor.SetTabChangesFocus(false)
	s.editor.SetReadOnly(true)

	width := charW * 30

	s.editor.SetTabStopWidth(charW * 2)
	s.editor.SetFixedWidth(width)

	scrollarea.SetMinimumWidth(width + 40)

	s.editor.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Expanding)

	scrollarea.ConnectResizeEvent(func(e *gui.QResizeEvent) {
		frame.SetFixedSize(core.NewQSize2(width+40, scrollarea.Geometry().Height()))
		scrollarea.ActivateWindow()
	})

	grid.AddWidget(s.editor, 0, 0, 0)

	comboBox := s.setupComboAttachs()
	//grid.AddLayout(comboBox, 1, 0, 0)

	grid.SetAlign(core.Qt__AlignHCenter)
	frame.SetLayout(grid)
	scrollarea.SetWidget(frame)

	editorLayout.AddWidget(scrollarea, 1, 0)
	editorLayout.AddLayout(comboBox, 1)

	res := widgets.NewQWidget(s.window, core.Qt__Widget)
	res.SetLayout(editorLayout)

	return res
}

func (s *myWindow) saveCurDiary() {
	if s.editor.Document().IsModified() == false {
		s.setStatusBar(T("No Diary Saved"))
		return
	}
	//fmt.Println("save", curDiary.Item.AccessibleText(), s.editor.ToHtml())
	var first bool = false
	if len(s.curDiary.Item.AccessibleText()) == 0 {
		p := s.curDiary.Item.Parent()
		pos := s.curDiary.Item.Index()
		item := p.TakeChild(pos.Row(), pos.Column())
		item.SetAccessibleText(fmt.Sprintf("%d", s.db.NextId()))
		p.SetChild(pos.Row(), pos.Column(), item)
		s.curDiary.Item = item
		first = true
	}

	filename := s.curDiary.Item.AccessibleText() + ".dat"
	id, err := strconv.Atoi(s.curDiary.Item.AccessibleText())
	if err != nil {
		s.setStatusBar(T("Error ID"))
		return
	}
	vs := strings.SplitN(s.curDiary.Item.Text(), "-", 2)
	title := vs[1]
	//fmt.Println(filename, title)
	if first {
		s.db.AddDiary(id, time.Now().Format("2006-01-02"), title, filename)
	} else {
		s.db.UpdateDiaryTitle(id, title)
	}

	encodeToFile(s.getRichText(), filename, s.key)
	s.setStatusBar(T("Save Diary") + fmt.Sprintf(" %s(%s)", title, filename))

	s.editor.Document().SetModified(false)
}

func (s *myWindow) addDiary(yearMonth, day, title string) {
	if s.curDiary.Item != nil {
		s.saveCurDiary()
	}

	month := s.addYearMonth(yearMonth)
	diary := gui.NewQStandardItem2(fmt.Sprintf("%s-%s", day, title))
	diary.SetEditable(false)
	diary.SetAccessibleText("")
	diary.SetAccessibleDescription("0")

	//month.AppendRow2(diary)
	month.InsertRow2(0, diary)

	s.tree.SetCurrentIndex(diary.Index())

	s.curDiary.Item = diary
	s.curDiary.Day = day
	s.curDiary.YearMonth = yearMonth

	s.document.Attachments = []attachmentFile{}
	s.document.Images = make(map[string][]byte)

	s.editor.SetReadOnly(false)
	s.exportEnc.SetEnabled(true)
	s.exportPdf.SetEnabled(true)

	s.setTitle(title)
	s.clearAttachs()

	s.tree.Expand(month.Index())
}

func (s *myWindow) addYearMonthsFromDb() {
	s.setStatusBar(T("Loading Diary List..."))
	s.editor.Document().Clear()
	s.editor.Document().SetModified(false)
	s.editor.SetReadOnly(true)
	s.curDiary.Item = nil
	s.curDiary.Id = 0

	s.bridge = NewQmlBridge(s.window)
	var wg sync.WaitGroup
	s.bridge.ConnectAddYearMonth(func(ym string) {
		s.addYearMonth(ym)
		wg.Done()
	})
	s.bridge.ConnectAddDiary(func(id, day, title, mtime string, r, c int) {
		item := s.model.Item(r, c)
		diary := gui.NewQStandardItem2(fmt.Sprintf("%s-%s", day, title))
		diary.SetEditable(false)
		diary.SetAccessibleText(id)
		diary.SetAccessibleDescription("0")
		diary.SetToolTip(T("Last Modified:") + mtime)

		item.AppendRow2(diary)
		wg.Done()
	})
	s.bridge.ConnectSetMonthFlag(func(r, c int) {
		item := s.model.TakeItem(r, c)
		item.SetAccessibleText("1")
		item.SetAccessibleDescription("1")
		s.model.SetItem(r, c, item)
		s.tree.ResizeColumnToContents(0)
		if r < 2 {
			s.tree.Expand(item.Index())
		}
	})
	go func() {
		yms, err := s.db.GetYearMonths()
		if err != nil {
			return
		}
		for i := 0; i < len(yms); i++ {
			wg.Add(1)
			s.bridge.AddYearMonth(yms[i])
		}
		wg.Wait()
		pidx := s.tree.RootIndex()
		for r := 0; r < s.model.RowCount(pidx); r++ {
			if r > 2 {
				break
			}
			c := 0
			item := s.model.Item(r, c)
			items, err := s.db.GetListFromYearMonth(item.Text())
			if err != nil {
				log.Println(err)
				return
			}
			for i := 0; i < len(items); i++ {
				wg.Add(1)
				s.bridge.AddDiary(strconv.Itoa(items[i].Id), items[i].Day, items[i].Title, items[i].MTime, r, c)
			}
			wg.Wait()
			s.bridge.SetMonthFlag(r, c)

		}
	}()

	return
}

func (s *myWindow) addYearMonth(yearMonth string) *gui.QStandardItem {
	idx := core.NewQModelIndex()
	p := s.model.Parent(idx)
	n := s.model.RowCount(p)
	for i := 0; i < n; i++ {
		item := s.model.Item(i, 0)
		if item.Text() == yearMonth {
			return item
		}
	}
	item := gui.NewQStandardItem2(yearMonth)
	item.SetColumnCount(1)
	item.SetEditable(false)
	item.SetAccessibleText("")

	s.model.InsertRow2(0, item)

	return item
}

func (s *myWindow) selectDiary(idx *core.QModelIndex) {
	diary := s.model.ItemFromIndex(idx)
	if diary.Pointer() == s.curDiary.Item.Pointer() {
		return
	}
	if len(diary.AccessibleText()) == 0 {
		return
	}
	filename := diary.AccessibleText() + ".dat"
	data, err := decodeFromFile(filename, s.key)

	if err != nil {
		//log.Println(err)
		s.getQText(data)
		if s.curDiary.Item != nil {
			s.saveCurDiary()
		}
		s.curDiary.Item = diary
		//curDiary.Modified = false
		s.curDiary.YearMonth = diary.Parent().Text()
		vs := strings.Index(diary.Text(), "-")
		s.curDiary.Day = diary.Text()[:vs]
		s.setTitle(diary.Text()[vs+1:])

	} else {
		if s.curDiary.Item != nil {
			s.saveCurDiary()
		}
		s.getQText(data)
		s.curDiary.Item = diary
		//curDiary.Modified = false
		s.curDiary.YearMonth = diary.Parent().Text()
		vs := strings.Index(diary.Text(), "-")
		s.curDiary.Day = diary.Text()[:vs]
		s.editor.Document().SetHtml(s.document.Html)

		s.showAttachList()
		s.editor.Document().SetModified(false)
	}
	s.tree.SetCurrentIndex(diary.Index())
	s.editor.SetReadOnly(false)
	s.exportEnc.SetEnabled(true)
	s.exportPdf.SetEnabled(true)

}

func (s *myWindow) popupOnMonth(item *gui.QStandardItem, point *core.QPoint) {
	s.LockEditor()
	menu := widgets.NewQMenu(s.tree)
	refreshItem := menu.AddAction(T("Refresh"))
	refreshItem.ConnectTriggered(func(checked bool) {
		items, err := s.db.GetListFromYearMonth(item.Text())
		idx := item.Index()
		item.RemoveRows(0, item.RowCount())
		r, c := idx.Row(), idx.Column()
		if err != nil {
			log.Println(err)
			return
		}
		for i := 0; i < len(items); i++ {
			id := strconv.Itoa(items[i].Id)
			day := items[i].Day
			title := items[i].Title
			mtime := items[i].MTime

			diary := gui.NewQStandardItem2(fmt.Sprintf("%s-%s", day, title))
			diary.SetEditable(false)
			diary.SetAccessibleText(id)
			diary.SetAccessibleDescription("0")
			diary.SetToolTip(T("Last Modified:") + mtime)

			item.AppendRow2(diary)
		}

		s.bridge.SetMonthFlag(r, c)
		s.tree.Expand(idx)
	})
	menu.Popup(point, nil)
}

func (s *myWindow) diaryPopup(idx *core.QModelIndex, e *gui.QMouseEvent) {
	//fmt.Println("popup")
	diary := s.model.ItemFromIndex(s.tree.CurrentIndex())
	if diary.AccessibleDescription() != "0" {
		s.popupOnMonth(diary, e.GlobalPos())
		return
	}
	s.selectDiary(idx)

	menu := widgets.NewQMenu(s.tree)
	delItem := menu.AddAction(T("Delete"))
	delItem.ConnectTriggered(func(checked bool) {
		dlg := widgets.NewQMessageBox(s.window)
		dlg.SetWindowTitle(T("Confirm Delete"))
		dlg.SetText(T("Are you sure?"))
		dlg.SetIcon(widgets.QMessageBox__Question)
		dlg.SetStandardButtons(widgets.QMessageBox__Yes | widgets.QMessageBox__Cancel)
		ret := dlg.Exec()
		if ret == int(widgets.QMessageBox__Yes) {
			id := diary.AccessibleText()
			if len(id) > 0 {
				s.db.RemoveDiary(id)
			}
			s.curDiary.Item = nil
			//curDiary.Modified = false
			p := diary.Parent()
			p.RemoveRow(diary.Row())

		}

	})

	openItem := menu.AddAction(T("Open in New Window"))
	openItem.ConnectTriggered(func(checked bool) {
		id, err := strconv.Atoi(diary.AccessibleText())
		if err != nil {
			s.setStatusBar(err.Error())
			return
		}
		OpenDiaryNewWindow(s, id)
	})

	category := menu.AddAction(T("Change Category"))
	category.ConnectTriggered(func(checked bool) {
		id, err := strconv.Atoi(diary.AccessibleText())
		if err != nil {
			s.setStatusBar(err.Error())
			return
		}
		dict := s.db.GetCategories()
		dict1 := make(map[string]int)
		dict1[T("Default")] = 0
		items := []string{T("Default")}
		for k, v := range dict {
			dict1[v] = k
			items = append(items, v)
		}
		var ok bool
		ret := widgets.QInputDialog_GetItem(s.window, T("Category"), T("Select a category"), items, 0, false, &ok, core.Qt__Dialog, 0)
		if ok && dict1[ret] != s.category {
			s.db.UpdateDiaryCategory(id, dict1[ret])
			s.model.RemoveRow(diary.Row(), diary.Index().Parent())
		}

	})

	menu.QWidget.AddAction(s.exportEnc)

	menu.QWidget.AddAction(s.exportPdf)

	menu.Popup(e.GlobalPos(), nil)
}

func (s *myWindow) LockEditor() {
	s.saveCurDiary()
	s.curDiary.Item = nil
	s.editor.Clear()
	s.editor.SetReadOnly(true)
	s.editor.Document().SetModified(false)
}

func (s *myWindow) onSelectItem(idx *core.QModelIndex) {
	item := s.model.ItemFromIndex(idx)
	//fmt.Println("AD:", item.AccessibleDescription())
	switch item.AccessibleDescription() {
	case "0":
		s.selectDiary(idx)
	case "1":
		s.LockEditor()
		return

	default:
		items, err := s.db.GetListFromYearMonth(item.Text())
		r, c := idx.Row(), idx.Column()
		if err != nil {
			log.Println(err)
			return
		}
		for i := 0; i < len(items); i++ {
			s.bridge.AddDiary(strconv.Itoa(items[i].Id), items[i].Day, items[i].Title, items[i].MTime, r, c)
		}

		s.bridge.SetMonthFlag(r, c)
		s.tree.Expand(idx)
	}
}

func (s *myWindow) setTreeFuncs() {

	s.tree.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)

	s.tree.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		idx := s.tree.IndexAt(e.Pos())
		if !idx.IsValid() {
			return
		}
		s.tree.SetCurrentIndex(idx)
		switch e.Button() {
		case core.Qt__LeftButton:
			return
		case core.Qt__RightButton:
			s.diaryPopup(idx, e)
		}

	})

	s.tree.ConnectSelectionChanged(func(selected *core.QItemSelection, deselected *core.QItemSelection) {
		idxes := selected.Indexes()
		if len(idxes) == 1 {
			s.onSelectItem(idxes[0])
		}

	})

}

func (s *myWindow) setTitle(v string) {
	s.editor.Clear()
	s.editor.Document().Clear()
	var cfmt = gui.NewQTextCharFormat()
	cfmt.SetFont2(s.app.Font())
	cfmt.SetFontPointSize(18)
	cfmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__blue), core.Qt__SolidPattern))

	s.editor.SetFontPointSize(18)

	s.editor.SetAlignment(core.Qt__AlignHCenter)

	cursor := s.editor.TextCursor()
	cursor.SetCharFormat(cfmt)
	cursor.InsertText(v)

	cursor.InsertText("\n")
	s.editor.SetTextCursor(cursor)

	var afmt = gui.NewQTextCharFormat()
	afmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__black), core.Qt__SolidPattern))
	afmt.SetFont2(s.app.Font())
	afmt.SetFontPointSize(14)

	s.editor.SetAlignment(core.Qt__AlignLeft | core.Qt__AlignAbsolute)

	cursor.InsertText("\n\t")
	s.mergeFormatOnLineOrSelection(afmt)
	s.editor.SetTextCursor(cursor)
}

func (s *myWindow) setEditorFuncs() {

	s.saveDiary.ConnectTriggered(func(b bool) {
		if s.model != nil {
			s.saveCurDiary()
		} else {
			s.saveDiaryAlone()
		}

	})
	if s.model != nil {
		s.editor.ConnectTextChanged(s.OnTextChanged)
	}

	s.editor.ConnectContextMenuEvent(func(e *gui.QContextMenuEvent) {

		menu := s.editor.CreateStandardContextMenu()

		imgUrl := s.getSelectedImage()
		if len(imgUrl) > 0 {
			//addaction export image
			act := menu.AddAction(T("Export Image As..."))
			act.ConnectTriggered(func(b bool) {
				ext := filepath.Ext(imgUrl)
				filter := fmt.Sprintf("%s Image (*%s)", ext, ext)
				filename := widgets.QFileDialog_GetSaveFileName(s.window, T("Export Image As..."), filepath.Base(imgUrl), filter, filter, 0)
				if strings.HasSuffix(filename, ext) == false {
					filename = filename + ext
				}

				if len(filename) > 0 {
					ioutil.WriteFile(filename, s.document.Images[imgUrl], 0644)
				}
			})

			actScale := menu.AddAction(T("Scale Image"))
			actScale.ConnectTriggered(func(b bool) {
				v := s.editor.Document().Resource(int(gui.QTextDocument__ImageResource), core.NewQUrl3(imgUrl, core.QUrl__TolerantMode))
				img := gui.NewQImageFromPointer(v.ToImage())
				res := s.scaleImage(img)

				s.editor.Document().AddResource(int(gui.QTextDocument__ImageResource), core.NewQUrl3(imgUrl, core.QUrl__TolerantMode), res.ToVariant())
				//save data
				ba := core.NewQByteArray()
				iod := core.NewQBuffer2(ba, nil)
				iod.Open(core.QIODevice__WriteOnly)

				ok := res.Save2(iod, filepath.Ext(filepath.Base(imgUrl))[1:], -1)
				//fmt.Println(filepath.Ext(filename))
				if ok {
					s.document.Images[imgUrl] = []byte(ba.Data())
				}
				s.editor.Document().SetModified(true)
			})
		}

		//add search
		menu.QWidget.AddAction(s.paste)
		menu.QWidget.AddAction(s.search)
		menu.QWidget.AddAction(s.replace)

		menu.QWidget.AddAction(s.lineMargin)

		menu.Popup(e.GlobalPos(), nil)
	})

	s.editor.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		defer e.Accept()
		if s.format.BF == nil {
			return
		}
		cursor := s.editor.TextCursor()
		if !cursor.HasSelection() {
			cursor.Select(gui.QTextCursor__LineUnderCursor)
		}
		cursor.SetBlockFormat(s.format.BF)
		cursor.SetCharFormat(s.format.CF)
		s.format.BF = nil
		s.format.CF = nil
		s.fb.SetChecked(false)
	})

}

func (s *myWindow) OnTextChanged() {
	if s.curDiary.Item == nil {
		return
	}

	disp1 := s.curDiary.Day + "-" + strings.TrimSpace(s.editor.Document().FirstBlock().Text())

	disp0 := s.curDiary.Item.Text()
	if disp0 != disp1 {
		p := s.curDiary.Item.Parent()
		pos := s.curDiary.Item.Index()
		s.curDiary.Item = p.TakeChild(pos.Row(), pos.Column())
		s.curDiary.Item.SetText(disp1)
		p.SetChild(pos.Row(), pos.Column(), s.curDiary.Item)
		s.tree.SetCurrentIndex(s.curDiary.Item.Index())
		s.tree.ResizeColumnToContents(0)
		if s.editor.CurrentCharFormat().FontPointSize() != 18 {
			var cfmt = gui.NewQTextCharFormat()
			cfmt.SetFont2(s.app.Font())
			cfmt.SetFontPointSize(18)
			cfmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__blue), core.Qt__SolidPattern))
			s.mergeFormatOnLineOrSelection(cfmt)
		}

		s.editor.Document().SetModified(true)
	}

}

//return url or ""
func (s *myWindow) getSelectedImage() (url string) {
	cursor := s.editor.TextCursor()
	if cursor == nil {
		return ""
	}

	fragment := cursor.Selection()
	if fragment.IsEmpty() {
		return ""
	}

	html := fragment.ToHtml(core.NewQByteArray2("UTF-8", 5))
	reg1, err := regexp.Compile(`<!--StartFragment-->.+<!--EndFragment-->`)
	if err != nil {
		log.Println(err)
		return ""
	}
	html = reg1.FindString(html)

	reg2, err := regexp.Compile(`<img src="([^"]+)" />`)
	if err != nil {
		log.Println(err)
		return ""
	}

	urls := reg2.FindAllStringSubmatch(html, -1)

	if len(urls) != 1 {
		return ""
	}

	return urls[0][1]
}

func (s *myWindow) clearModel() {
	s.model.Clear()
	s.model.SetHorizontalHeaderLabels([]string{T("Diary List") + " - " + s.categoryName})
}

func (s *myWindow) clearModelFind() {
	s.modelFind.Clear()
	s.modelFind.SetHorizontalHeaderLabels([]string{T("Result List") + " - " + s.categoryName})
}

func (s *myWindow) lastUser() string {
	home, _ := os.UserHomeDir()
	v, err := ioutil.ReadFile(path.Join(home, ".sdiary", "user.txt"))
	if err != nil {
		return ""
	}
	return string(v)
}

func (s *myWindow) saveLastUser(name string) {
	home, _ := os.UserHomeDir()
	path1 := path.Join(home, ".sdiary", "user.txt")
	ioutil.WriteFile(path1, []byte(name), 0644)
}

func (s *myWindow) mergeFormatOnLineOrSelection(format *gui.QTextCharFormat) *gui.QTextCursor {
	var cursor = s.editor.TextCursor()
	if !cursor.HasSelection() {
		cursor.Select(gui.QTextCursor__LineUnderCursor)
	}
	cursor.MergeCharFormat(format)
	s.editor.MergeCurrentCharFormat(format)

	return cursor
}

func (s *myWindow) rename() {
	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetWindowTitle(T("Rename"))

	grid := widgets.NewQGridLayout(dlg)

	name := widgets.NewQLabel2(T("New Name:"), dlg, core.Qt__Widget)
	grid.AddWidget(name, 0, 0, 0)

	nameInput := widgets.NewQLineEdit(dlg)
	grid.AddWidget(nameInput, 0, 1, 0)

	okAct := widgets.NewQPushButton2(T("OK"), dlg)
	grid.AddWidget(okAct, 1, 0, 0)
	cancelAct := widgets.NewQPushButton2(T("Cancel"), dlg)
	grid.AddWidget(cancelAct, 1, 1, 0)

	okAct.ConnectClicked(func(checked bool) {
		ret := widgets.QMessageBox_Question(dlg, T("Confirm"), T("Are you sure?"), widgets.QMessageBox__Yes|widgets.QMessageBox__Cancel, widgets.QMessageBox__Yes)
		//fmt.Println(ret)
		if ret == widgets.QMessageBox__Yes {
			if len(nameInput.Text()) == 0 {
				return
			}
			s.db.Close()
			dstName := path.Join(path.Dir(dataDir), nameInput.Text())
			err := os.Rename(dataDir, dstName)
			if err == nil {
				widgets.QMessageBox_Information(dlg, T("Information"), T("Success,Please login again."), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

			} else {
				widgets.QMessageBox_Information(dlg, T("Information"), T("Fail,Please login again."), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			}
		}
		dlg.Hide()
		s.window.Close()
	})

	cancelAct.ConnectClicked(func(c bool) {
		dlg.Hide()
		dlg.Destroy(true, true)
	})

	dlg.SetLayout(grid)
	dlg.SetModal(true)
	dlg.Show()
}

func (s *myWindow) updatePwd() {
	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetWindowTitle(T("Modify Password"))

	grid := widgets.NewQGridLayout(dlg)

	pwd1 := widgets.NewQLabel2(T("Old Password:"), dlg, core.Qt__Widget)

	grid.AddWidget(pwd1, 0, 0, 0)

	pwdInput1 := widgets.NewQLineEdit(dlg)
	pwdInput1.SetEchoMode(widgets.QLineEdit__Password)
	pwdInput1.SetPlaceholderText(T("Length Must >= 4"))
	grid.AddWidget(pwdInput1, 0, 1, 0)

	pwd2 := widgets.NewQLabel2(T("New Password:"), dlg, core.Qt__Widget)

	grid.AddWidget(pwd2, 1, 0, 0)

	pwdInput2 := widgets.NewQLineEdit(dlg)
	pwdInput2.SetEchoMode(widgets.QLineEdit__Password)
	pwdInput2.SetPlaceholderText(T("Length Must >= 4"))
	grid.AddWidget(pwdInput2, 1, 1, 0)

	pwd3 := widgets.NewQLabel2(T("Confirm Password:"), dlg, core.Qt__Widget)

	grid.AddWidget(pwd3, 2, 0, 0)

	pwdInput3 := widgets.NewQLineEdit(dlg)
	pwdInput3.SetEchoMode(widgets.QLineEdit__Password)
	pwdInput3.SetPlaceholderText(T("Length Must >= 4"))
	grid.AddWidget(pwdInput3, 2, 1, 0)

	okAct := widgets.NewQPushButton2(T("OK"), dlg)
	grid.AddWidget(okAct, 3, 0, 0)
	cancelAct := widgets.NewQPushButton2(T("Cancel"), dlg)
	grid.AddWidget(cancelAct, 3, 1, 0)

	okAct.ConnectClicked(func(checked bool) {
		ret := widgets.QMessageBox_Question(dlg, T("Confirm"), T("Are you sure?"), widgets.QMessageBox__Yes|widgets.QMessageBox__Cancel, widgets.QMessageBox__Yes)

		if ret == widgets.QMessageBox__Yes {
			if len(pwdInput3.Text()) < 4 {
				return
			}
			if pwdInput2.Text() != pwdInput3.Text() {
				return
			}
			err := s.db.UpdateRealKeyAndSha40(pwdInput1.Text(), pwdInput2.Text())
			if err == nil {
				widgets.QMessageBox_Information(dlg, T("Information"), T("Success,Please login again."), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				dlg.Hide()
				s.window.Close()
			} else {
				widgets.QMessageBox_Information(dlg, T("Information"), T("Fail")+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			}
		}
		dlg.Hide()
		dlg.Destroy(true, true)
	})

	cancelAct.ConnectClicked(func(c bool) {
		dlg.Hide()
		dlg.Destroy(true, true)
	})

	dlg.SetLayout(grid)
	dlg.SetModal(true)
	dlg.Show()
}

func (s *myWindow) getUsers() []string {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".sdiary")
	hdir, err := os.Open(dir)
	if err != nil {
		return []string{}
	}
	dirs, err := hdir.Readdir(-1)
	if err != nil {
		return nil
	}
	res := []string{}
	for _, v := range dirs {
		if v.IsDir() {
			res = append(res, v.Name())
		}
	}
	return res
}

func (s *myWindow) login() {
	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetWindowTitle(T("Login"))

	grid := widgets.NewQGridLayout(dlg)

	name := widgets.NewQLabel2(T("Name:"), dlg, core.Qt__Widget)
	grid.AddWidget(name, 0, 0, 0)

	nameInput := widgets.NewQComboBox(dlg)
	nameInput.SetEditable(true)

	nameInput.AddItems(s.getUsers())
	nameInput.SetCurrentText(s.lastUser())

	grid.AddWidget(nameInput, 0, 1, 0)

	pwd := widgets.NewQLabel2(T("Password:"), dlg, core.Qt__Widget)

	grid.AddWidget(pwd, 1, 0, 0)

	pwdInput := widgets.NewQLineEdit(dlg)
	pwdInput.SetEchoMode(widgets.QLineEdit__Password)
	pwdInput.SetPlaceholderText(T("Length Must >= 4"))
	if len(nameInput.CurrentText()) > 0 {
		pwdInput.SetFocus2()
	}
	grid.AddWidget(pwdInput, 1, 1, 0)

	btb := widgets.NewQGridLayout(nil)

	okBtn := widgets.NewQPushButton2(T("OK"), dlg)
	btb.AddWidget(okBtn, 0, 0, 0)

	regBtn := widgets.NewQPushButton2(T("Register"), dlg)
	regBtn.SetToolTip(T("Please input the name and password to register before you click this button!"))
	btb.AddWidget(regBtn, 0, 1, 0)

	cancelBtn := widgets.NewQPushButton2(T("Cancel"), dlg)
	btb.AddWidget(cancelBtn, 0, 2, 0)

	grid.AddLayout2(btb, 2, 0, 1, 2, 0)

	dlg.SetLayout(grid)

	dlg.SetModal(true)

	minLen := 4

	okBtn.ConnectClicked(func(b bool) {
		var err error
		if len(pwdInput.Text()) < minLen || len(nameInput.CurrentText()) < 1 {
			return
		}
		s.db, err = getMyDb(nameInput.CurrentText())
		if err != nil {
			widgets.QMessageBox_Information(dlg, T("Information"), T("Fail"), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			nameInput.SetFocus2()
			return
		}
		s.key, err = s.db.GetRealKey(pwdInput.Text())
		if err != nil {
			widgets.QMessageBox_Information(dlg, T("Information"), T("Fail"), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			pwdInput.SetFocus2()
			return
		}
		s.saveLastUser(nameInput.CurrentText())
		s.user = nameInput.CurrentText()
		s.window.SetWindowTitle(nameInput.CurrentText())
		dlg.Hide()

	})

	regBtn.ConnectClicked(func(b bool) {
		if len(pwdInput.Text()) < minLen || len(nameInput.CurrentText()) < 1 {
			widgets.QMessageBox_Warning(s.window, T("Register"),
				T("Please input the name and password to register before you click this button!"), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			return
		}
		var ok bool
		confirm := widgets.QInputDialog_GetText(dlg, T("Confirm"), T("Name:")+nameInput.CurrentText(), widgets.QLineEdit__Password, "",
			&ok, core.Qt__Dialog, core.Qt__ImhNone)
		if !ok {
			return
		}
		if confirm != pwdInput.Text() {
			s.showMsg(T("Fail"), T("Password Not Match"))
			return
		}
		err := createUserDb(nameInput.CurrentText(), pwdInput.Text())
		if err != nil {
			panic(err)
		}
		s.db, err = getMyDb(nameInput.CurrentText())
		if err != nil {
			panic(err)
		}
		s.key, err = s.db.GetRealKey(pwdInput.Text())
		if err != nil {
			panic(err)
		}
		s.saveLastUser(nameInput.CurrentText())
		s.user = nameInput.CurrentText()
		s.window.SetWindowTitle(nameInput.CurrentText())
		dlg.Hide()

	})

	cancelBtn.ConnectClicked(func(b bool) {
		dlg.Destroy(true, true)
		s.window.Destroy(true, true)
		s.app.Quit()
	})

	dlg.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		dlg.Destroy(true, true)
		s.window.Destroy(true, true)
		s.app.Quit()
	})

	dlg.ConnectHideEvent(func(e *gui.QHideEvent) {
		//log.Println("load list")
		s.loadUserCategorys()
		dlg.Destroy(true, true)
	})

	dlg.Show()

	dlg.Move2(0, 0)
	dlg.Move2(s.window.X()+(s.window.Width()-dlg.Width())/2, s.window.Y()+(s.window.Height()-dlg.Width())/2)
}

func (s *myWindow) exportEncryptedDiary() {
	//export private format .egf	(encrypted gob file)
	idx := s.tree.CurrentIndex()
	if idx == nil {
		return
	}
	item := s.model.ItemFromIndex(idx)
	if item == nil {
		return
	}
	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetWindowTitle(T("Export Encrypted Diary") + "...")

	grid := widgets.NewQGridLayout2()

	name := widgets.NewQLabel2(T("FileName:"), dlg, core.Qt__Widget)
	grid.AddWidget(name, 0, 0, 0)

	nameInput := widgets.NewQLineEdit(dlg)
	grid.AddWidget(nameInput, 0, 1, 0)
	nameInput.SetPlaceholderText(T("Double click to select..."))
	nameInput.SetMinimumWidth(240)

	passwd := widgets.NewQLabel2(T("Password:"), dlg, core.Qt__Widget)
	grid.AddWidget(passwd, 1, 0, 0)

	passwdInput := widgets.NewQLineEdit(dlg)
	grid.AddWidget(passwdInput, 1, 1, 0)
	passwdInput.SetEchoMode(widgets.QLineEdit__Password)
	passwdInput.SetPlaceholderText(T("Length Must >= 4"))

	hbox := widgets.NewQHBoxLayout()

	okBtn := widgets.NewQPushButton2(T("OK"), dlg)
	hbox.AddWidget(okBtn, 1, 0)

	cancelBtn := widgets.NewQPushButton2(T("Cancel"), dlg)
	hbox.AddWidget(cancelBtn, 1, 0)

	cancelBtn.ConnectClicked(func(b bool) {
		dlg.Hide()
		dlg.Destroy(true, true)
	})

	nameInput.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {
		filename := widgets.QFileDialog_GetSaveFileName(dlg, T("Export To..."), ".", "Encrypted Diary (*.egf)", "Encrypted Diary (*.egf)", 0)
		if strings.HasSuffix(strings.ToLower(filename), ".egf") {
			nameInput.SetText(filename)
		} else {
			nameInput.SetText(filename + ".egf")
		}

	})

	okBtn.ConnectClicked(func(b bool) {
		if len(passwdInput.Text()) < 4 {
			return
		}

		filename := strings.TrimSpace(nameInput.Text())
		if len(filename) == 0 {
			return
		}
		key := getSha4(passwdInput.Text())
		data := s.getRichText()
		err := encodeToPathName(data, filename, key)
		if err != nil {
			s.showMsg(T("Error"), err.Error())
		} else {
			s.showMsg(T("Success"), filename)
		}
		dlg.Hide()
		dlg.Destroy(true, true)
	})

	grid.AddLayout2(hbox, 2, 0, 1, 2, 0)

	dlg.SetLayout(grid)
	dlg.Exec()
}

func (s *myWindow) importEncryptedDiary() {
	//import private format .egf	(encrypted gob file)

	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetWindowTitle(T("Import Encrypted Diary") + "...")

	grid := widgets.NewQGridLayout2()

	name := widgets.NewQLabel2(T("FileName:"), dlg, core.Qt__Widget)
	grid.AddWidget(name, 0, 0, 0)

	nameInput := widgets.NewQLineEdit(dlg)
	grid.AddWidget(nameInput, 0, 1, 0)
	nameInput.SetPlaceholderText(T("Double click to select..."))
	nameInput.SetMinimumWidth(240)

	passwd := widgets.NewQLabel2(T("Password:"), dlg, core.Qt__Widget)
	grid.AddWidget(passwd, 1, 0, 0)

	passwdInput := widgets.NewQLineEdit(dlg)
	grid.AddWidget(passwdInput, 1, 1, 0)
	passwdInput.SetEchoMode(widgets.QLineEdit__Password)

	hbox := widgets.NewQHBoxLayout()

	okBtn := widgets.NewQPushButton2(T("OK"), dlg)
	hbox.AddWidget(okBtn, 1, 0)

	cancelBtn := widgets.NewQPushButton2(T("Cancel"), dlg)
	hbox.AddWidget(cancelBtn, 1, 0)

	cancelBtn.ConnectClicked(func(b bool) {
		dlg.Hide()
		dlg.Destroy(true, true)
	})

	nameInput.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {
		filename := widgets.QFileDialog_GetOpenFileName(dlg, T("Import File..."), ".", "Encrypted Diary .egf (*.egf)", "Encrypted Diary .egf (*.egf)", widgets.QFileDialog__ReadOnly)
		nameInput.SetText(filename)

	})

	okBtn.ConnectClicked(func(b bool) {
		if len(passwdInput.Text()) < 4 {
			return
		}

		filename := strings.TrimSpace(nameInput.Text())
		if len(filename) == 0 {
			return
		}
		key := getSha4(passwdInput.Text())
		data, err := decodeFromPathName(filename, key)
		if err != nil {
			return
		}
		now := time.Now()
		s.addDiary(now.Format("2006-01"), now.Format("02"), "")

		_, err = s.getQText(data)

		if err == nil {

			s.editor.SetHtml(s.document.Html)
			s.showAttachList()
			//curDiary.Modified = true
			s.editor.Document().SetModified(true)

			s.saveCurDiary()
		} else {
			s.showMsg(T("Error"), err.Error())
		}

		dlg.Hide()
		dlg.Destroy(true, true)
	})

	grid.AddLayout2(hbox, 2, 0, 1, 2, 0)

	dlg.SetLayout(grid)
	dlg.Exec()
}

func (s *myWindow) exportAsPdf() {
	filename := widgets.QFileDialog_GetSaveFileName(s.window, T("Export To..."), "out.pdf", "PDF (*.pdf)", "Encrypted Diary (*.pdf)", 0)
	if !strings.HasSuffix(strings.ToLower(filename), ".pdf") {
		filename = filename + ".pdf"
	}
	var (
		fileName = filename
		printer  = printsupport.NewQPrinter(printsupport.QPrinter__HighResolution)
	)
	printer.SetOutputFormat(printsupport.QPrinter__PdfFormat)

	printer.SetOutputFileName(fileName)
	s.editor.Document().Print(printer)
	s.setStatusBar(fmt.Sprintf("Exported %v", core.QDir_ToNativeSeparators(fileName)))
}

func (s *myWindow) showMsg(title, msg string) {
	dlg := widgets.NewQMessageBox(s.window)
	dlg.SetWindowTitle(title)
	dlg.SetText(msg)

	dlg.Exec()
	dlg.Destroy(true, true)
}

func (s *myWindow) showPay() {
	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetWindowTitle(T("Pay Voluntary!"))
	box := widgets.NewQHBoxLayout2(dlg)
	height := 512

	labelAli := widgets.NewQLabel(dlg, core.Qt__Widget)
	img1 := gui.NewQPixmap5(":/qml/pay/alipay.png", "png", core.Qt__NoFormatConversion)
	img1 = img1.ScaledToHeight(height, core.Qt__SmoothTransformation)
	labelAli.SetPixmap(img1)
	box.AddWidget(labelAli, 1, 0)

	labelWx := widgets.NewQLabel(dlg, core.Qt__Widget)
	img2 := gui.NewQPixmap5(":/qml/pay/wxpay.png", "png", core.Qt__NoFormatConversion)
	img2 = img2.ScaledToHeight(height, core.Qt__SmoothTransformation)
	labelWx.SetPixmap(img2)
	box.AddWidget(labelWx, 1, 0)

	dlg.SetLayout(box)

	dlg.Exec()
	dlg.Destroy(true, true)
}

func main() {
	//defer gettext.SaveLog()
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := new(myWindow)
	window.Create(app)
	app.Exec()
}
