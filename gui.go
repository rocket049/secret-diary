package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/therecipe/qt/gui"

	"github.com/rocket049/gettext-go/gettext"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func init() {
	exe1, _ := os.Executable()
	dir1 := path.Dir(exe1)
	locale1 := path.Join(dir1, "locale")
	gettext.BindTextdomain("sdiary", locale1, nil)
	gettext.Textdomain("sdiary")
}

var curDiary struct {
	Item      *gui.QStandardItem
	Day       string
	YearMonth string
	Modified  bool
}

func T(v string) string {
	return gettext.T(v)
}

type myWindow struct {
	app                 *widgets.QApplication
	window              *widgets.QMainWindow
	tree                *widgets.QTreeView
	model               *gui.QStandardItemModel
	editor              *widgets.QTextEdit
	actionTextBold      *widgets.QAction
	actionTextUnderline *widgets.QAction
	actionTextItalic    *widgets.QAction
	actionStrikeOut     *widgets.QAction
	user                string
	key                 []byte
	db                  *myDb
	bridge              *QmlBridge
	document            qtextFormat
}

func (s *myWindow) getNewId() (res int) {
	return s.db.NextId()
}

func (s *myWindow) setMenuBar() {
	s.document.Images = make(map[string][]byte)

	menubar := s.window.MenuBar()
	menu := menubar.AddMenu2(T("Manage"))

	ren := menu.AddAction(T("Rename"))
	ren.ConnectTriggered(func(b bool) {
		s.rename()
	})

	pwdNew := menu.AddAction(T("ModifyPassword"))
	pwdNew.ConnectTriggered(func(b bool) {
		s.updatePwd()
	})

	menu = menubar.AddMenu2(T("Help"))

	about := menu.AddAction(T("About"))
	about.ConnectTriggered(func(b bool) {
		s.showMsg(T("About"), T("Copy Right: Fu Huizhong <fuhuizn@163.com>\nSecurity Diary Tool.\nCrypto with AES256"))
	})

	bak := menu.AddAction(T("Backup"))
	bak.ConnectTriggered(func(b bool) {
		s.showMsg(T("Backup"), T("Your DATA storge path is '")+dataDir+"'\nYou can backup the directory yourself.")
	})

	pwd := menu.AddAction(T("Password"))
	pwd.ConnectTriggered(func(b bool) {
		s.showMsg(T("Password"), T("There is no way to recover password. \nPlease write it on paper!"))
	})
}

func (s *myWindow) setToolBar() {
	bar := widgets.NewQToolBar("File", nil)
	act1 := bar.AddAction(T("New"))
	act1.ConnectTriggered(func(b bool) {
		s.setStatusBar(T("New Diary"))
		now := time.Now()
		s.addDiary(now.Format("2006-01"), now.Format("02"), T("New Dialy"))
	})

	act2 := bar.AddAction(T("Save"))
	act2.ConnectTriggered(func(b bool) {

		s.saveCurDiary()
	})

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

	s.addJustifyActions(bar)

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

	s.window.AddToolBar2(bar)
}

func (s *myWindow) setStatusBar(msg string) {
	bar := s.window.StatusBar()
	bar.ShowMessage(msg, 20000)
}

func (s *myWindow) Create(app *widgets.QApplication) {
	s.app = app
	s.window = widgets.NewQMainWindow(nil, core.Qt__Window)

	s.window.SetWindowTitle(T("UserName"))
	s.window.SetMinimumSize2(800, 600)
	s.window.SetWindowIcon(gui.NewQIcon5(":/qml/icons/Sd.png"))

	grid := widgets.NewQGridLayout2()

	frame := widgets.NewQFrame(nil, core.Qt__Widget)

	s.window.SetCentralWidget(frame)

	s.tree = widgets.NewQTreeView(s.window)
	s.tree.SetFixedWidth(240)
	s.tree.SetSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Expanding)
	s.tree.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAsNeeded)
	s.tree.SetAutoScroll(true)
	s.model = gui.NewQStandardItemModel2(0, 1, s.tree)
	s.model.SetHorizontalHeaderLabels([]string{T("Diary List")})

	s.tree.SetModel(s.model)

	grid.AddWidget(s.tree, 0, 0, 0)

	s.editor = widgets.NewQTextEdit(s.window)
	s.editor.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	s.editor.SetTabChangesFocus(false)

	grid.AddWidget(s.editor, 0, 1, 0)

	grid.SetAlign(core.Qt__AlignTop)

	s.setToolBar()

	s.setMenuBar()

	frame.SetLayout(grid)

	app.SetActiveWindow(s.window)
	app.SetApplicationDisplayName(T("Secret Diary"))
	app.SetApplicationName("sdiary")
	app.SetApplicationVersion("1.0.1")

	s.setEditorFuncs()
	s.setTreeFuncs()

	once := sync.Once{}
	s.window.ConnectShowEvent(func(e *gui.QShowEvent) {
		once.Do(s.login)
	})

	s.window.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		s.db.Close()
	})
	s.window.ShowMaximized()
}

func (s *myWindow) saveCurDiary() {
	if !curDiary.Modified || s.editor.Document().IsModified() == false {
		s.setStatusBar(T("No Diary Saved"))
		return
	}
	//fmt.Println("save", curDiary.Item.AccessibleText(), s.editor.ToHtml())
	var first bool = false
	if len(curDiary.Item.AccessibleText()) == 0 {
		p := curDiary.Item.Parent()
		pos := curDiary.Item.Index()
		item := p.TakeChild(pos.Row(), pos.Column())
		item.SetAccessibleText(fmt.Sprintf("%d", s.db.NextId()))
		p.SetChild(pos.Row(), pos.Column(), item)
		curDiary.Item = item
		first = true
	}

	filename := curDiary.Item.AccessibleText() + ".dat"
	id, err := strconv.Atoi(curDiary.Item.AccessibleText())
	if err != nil {
		s.setStatusBar(T("Error ID"))
		return
	}
	vs := strings.SplitN(curDiary.Item.Text(), "-", 2)
	title := vs[1]
	//fmt.Println(filename, title)
	if first {
		s.db.AddDiary(id, time.Now().Format("2006-01-02"), title, filename)
	} else {
		s.db.UpdateDiaryTitle(id, title)
	}

	encodeToFile(s.getRichText(), filename, s.key)
	s.setStatusBar(T("Save Diary") + fmt.Sprintf(" %s(%s)", title, filename))
	curDiary.Modified = false
}

func (s *myWindow) addDiary(yearMonth, day, title string) {
	if curDiary.Item != nil {
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

	curDiary.Item = diary
	curDiary.Day = day
	curDiary.YearMonth = yearMonth

	s.setTitle(title)

	s.tree.Expand(month.Index())
}

func (s *myWindow) addYearMonthsFromDb() {
	s.setStatusBar(T("Loading Diary List..."))

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
		topidx := core.NewQModelIndex()
		pidx := s.model.Parent(topidx)
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
				s.bridge.AddDiary(strconv.Itoa(items[i].Id), items[i].Day, items[i].Title, items[i].MTime, r, c)
			}

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

	s.model.AppendRow2(item)

	return item
}

func (s *myWindow) selectDiary(idx *core.QModelIndex) {
	diary := s.model.ItemFromIndex(idx)
	if diary.Pointer() == curDiary.Item.Pointer() {
		return
	}
	if len(diary.AccessibleText()) == 0 {
		return
	}
	filename := diary.AccessibleText() + ".dat"
	data, err := decodeFromFile(filename, s.key)

	if err == nil {
		_, err = s.getQText(data)
	}
	if err != nil {
		//log.Println(err)
		if curDiary.Item != nil {
			s.saveCurDiary()
		}
		curDiary.Item = diary
		curDiary.Modified = false
		curDiary.YearMonth = diary.Parent().Text()
		vs := strings.Index(diary.Text(), "-")
		curDiary.Day = diary.Text()[:vs]
		s.setTitle(diary.Text()[vs+1:])

	} else {
		if curDiary.Item != nil {
			s.saveCurDiary()
		}
		curDiary.Item = diary
		curDiary.Modified = false
		curDiary.YearMonth = diary.Parent().Text()
		vs := strings.Index(diary.Text(), "-")
		curDiary.Day = diary.Text()[:vs]

		s.editor.Document().SetHtml(s.document.Html)
		s.editor.Document().SetModified(false)
	}

}

func (s *myWindow) diaryPopup(idx *core.QModelIndex, e *gui.QMouseEvent) {
	//fmt.Println("popup")
	diary := s.model.ItemFromIndex(s.tree.CurrentIndex())
	if diary.AccessibleDescription() != "0" {
		return
	}
	menu := widgets.NewQMenu(s.tree)
	item := menu.AddAction(T("Delete"))
	item.ConnectTriggered(func(checked bool) {
		dlg := widgets.NewQMessageBox(s.window)
		dlg.SetWindowTitle(T("Confirm"))
		dlg.SetText(T("Are you sure?"))
		dlg.SetIcon(widgets.QMessageBox__Question)
		dlg.SetStandardButtons(widgets.QMessageBox__Yes | widgets.QMessageBox__Cancel)
		ret := dlg.Exec()
		if ret == int(widgets.QMessageBox__Yes) {
			id := diary.AccessibleText()
			if len(id) > 0 {
				s.db.RemoveDiary(id)
			}
			curDiary.Item = nil
			curDiary.Modified = false
			p := diary.Parent()
			p.RemoveRow(diary.Row())

		}

	})

	menu.Popup(e.GlobalPos(), nil)
}

func (s *myWindow) onSelectItem(idx *core.QModelIndex) {
	item := s.model.ItemFromIndex(idx)
	//fmt.Println("AD:", item.AccessibleDescription())
	switch item.AccessibleDescription() {
	case "0":
		s.selectDiary(idx)
	case "1":
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

		switch e.Button() {
		case core.Qt__LeftButton:
			s.onSelectItem(idx)
		case core.Qt__RightButton:
			s.diaryPopup(idx, e)
		}

	})

}

func (s *myWindow) setTitle(v string) {
	s.editor.Clear()
	var cfmt = gui.NewQTextCharFormat()
	cfmt.SetFontPointSize(18)
	cfmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__blue), core.Qt__SolidPattern))

	cursor := s.editor.TextCursor()
	cursor.InsertText(v)
	s.mergeFormatOnLineOrSelection(cfmt)

	//cursor.MovePosition(gui.QTextCursor__End, gui.QTextCursor__KeepAnchor, 0)
	cursor.InsertText("\n")
	var afmt = gui.NewQTextCharFormat()
	afmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__black), core.Qt__SolidPattern))
	afmt.SetFontPointSize(14)

	s.mergeFormatOnLineOrSelection(afmt)
	cursor.InsertText("\n")

	//cursor.EndEditBlock()
}

func (s *myWindow) setEditorFuncs() {
	var cfmt = gui.NewQTextCharFormat()
	cfmt.SetFontPointSize(18)
	cfmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__blue), core.Qt__SolidPattern))
	s.editor.SetCurrentCharFormat(cfmt)

	s.editor.ConnectTextChanged(func() {
		curDiary.Modified = true
		v := s.editor.ToPlainText()
		secs := strings.Index(v, "\n")
		var disp1 string
		if secs == -1 {
			disp1 = v
		} else {
			disp1 = fmt.Sprintf("%s-%s", curDiary.Day, v[:secs])
		}

		disp0 := curDiary.Item.Text()
		if disp0 != disp1 {
			p := curDiary.Item.Parent()
			pos := curDiary.Item.Index()
			curDiary.Item = p.TakeChild(pos.Row(), pos.Column())
			curDiary.Item.SetText(disp1)
			p.SetChild(pos.Row(), pos.Column(), curDiary.Item)
			s.tree.ResizeColumnToContents(0)
		}

	})
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

func (s *myWindow) mergeFormatOnLineOrSelection(format *gui.QTextCharFormat) {
	var cursor = s.editor.TextCursor()
	if !cursor.HasSelection() {
		cursor.Select(gui.QTextCursor__LineUnderCursor)
	}
	cursor.MergeCharFormat(format)
	s.editor.MergeCurrentCharFormat(format)
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
		fmt.Println(ret)
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

func (s *myWindow) login() {
	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetWindowTitle(T("Login"))

	grid := widgets.NewQGridLayout(dlg)

	name := widgets.NewQLabel2(T("Name:"), dlg, core.Qt__Widget)
	grid.AddWidget(name, 0, 0, 0)

	nameInput := widgets.NewQLineEdit(dlg)
	nameInput.SetText(s.lastUser())
	grid.AddWidget(nameInput, 0, 1, 0)

	pwd := widgets.NewQLabel2(T("Password:"), dlg, core.Qt__Widget)

	grid.AddWidget(pwd, 1, 0, 0)

	pwdInput := widgets.NewQLineEdit(dlg)
	pwdInput.SetEchoMode(widgets.QLineEdit__Password)
	pwdInput.SetPlaceholderText(T("Length Must >= 4"))
	if len(nameInput.Text()) > 0 {
		pwdInput.SetFocus2()
	}
	grid.AddWidget(pwdInput, 1, 1, 0)

	btb := widgets.NewQGridLayout(nil)

	okBtn := widgets.NewQPushButton2(T("OK"), dlg)
	btb.AddWidget(okBtn, 0, 0, 0)

	regBtn := widgets.NewQPushButton2(T("Register"), dlg)
	btb.AddWidget(regBtn, 0, 1, 0)

	cancelBtn := widgets.NewQPushButton2(T("Cancel"), dlg)
	btb.AddWidget(cancelBtn, 0, 2, 0)

	grid.AddLayout2(btb, 2, 0, 1, 2, 0)

	dlg.SetLayout(grid)

	dlg.SetModal(true)

	minLen := 4

	okBtn.ConnectClicked(func(b bool) {
		var err error
		if len(pwdInput.Text()) < minLen || len(nameInput.Text()) < 1 {
			return
		}
		s.db, err = getMyDb(nameInput.Text())
		if err != nil {
			widgets.QMessageBox_Information(dlg, T("Information"), T("Fail")+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			dlg.Hide()
			s.window.Close()
		}
		s.key, err = s.db.GetRealKey(pwdInput.Text())
		if err != nil {
			widgets.QMessageBox_Information(dlg, T("Information"), T("Fail")+err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			dlg.Hide()
			s.window.Close()
		}
		s.saveLastUser(nameInput.Text())
		s.window.SetWindowTitle(nameInput.Text())
		dlg.Hide()

	})

	regBtn.ConnectClicked(func(b bool) {
		if len(pwdInput.Text()) < minLen || len(nameInput.Text()) < 1 {
			return
		}
		var ok bool
		confirm := widgets.QInputDialog_GetText(dlg, T("Confirm"), T("Name:")+nameInput.Text(), widgets.QLineEdit__Password, "",
			&ok, core.Qt__Dialog, core.Qt__ImhNone)
		if !ok {
			return
		}
		if confirm != pwdInput.Text() {
			s.showMsg(T("Fail"), T("Password Not Match"))
			return
		}
		err := createUserDb(nameInput.Text(), pwdInput.Text())
		if err != nil {
			panic(err)
		}
		s.db, err = getMyDb(nameInput.Text())
		if err != nil {
			panic(err)
		}
		s.key, err = s.db.GetRealKey(pwdInput.Text())
		if err != nil {
			panic(err)
		}
		s.saveLastUser(nameInput.Text())
		s.window.SetWindowTitle(nameInput.Text())
		dlg.Hide()

	})

	cancelBtn.ConnectClicked(func(b bool) {
		dlg.Destroy(true, true)
		s.window.Destroy(true, true)
		s.app.Quit()
	})

	dlg.ConnectHideEvent(func(e *gui.QHideEvent) {
		//log.Println("load list")
		s.addYearMonthsFromDb()
		dlg.Destroy(true, true)
	})

	dlg.Show()
}

func (s *myWindow) showMsg(title, msg string) {
	dlg := widgets.NewQMessageBox(s.window)
	dlg.SetWindowTitle(title)
	dlg.SetText(msg)
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
