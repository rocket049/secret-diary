package main

import (
	"fmt"
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
	app    *widgets.QApplication
	window *widgets.QMainWindow
	tree   *widgets.QTreeView
	model  *gui.QStandardItemModel
	editor *widgets.QTextEdit
	user   string
	key    []byte
	db     *myDb
	bridge *QmlBridge
}

func (s *myWindow) getNewId() (res int) {
	return s.db.NextId()
}

func (s *myWindow) setMenuBar() {
	menubar := s.window.MenuBar()
	menu := menubar.AddMenu2("&File")
	open := menu.AddAction("&Open")
	open.ConnectTriggered(func(b bool) {
		log.Println("click open", b)
		s.setStatusBar("click open")
	})
}

func (s *myWindow) setToolBar() {
	bar := widgets.NewQToolBar("toolbar 1", nil)
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

	grid := widgets.NewQGridLayout2()

	frame := widgets.NewQFrame(nil, core.Qt__Widget)

	s.window.SetCentralWidget(frame)

	s.tree = widgets.NewQTreeView(s.window)
	s.tree.SetFixedWidth(200)
	s.tree.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Expanding)
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
	app.SetApplicationVersion("0.0.1")

	s.setEditorFuncs()
	s.setTreeFuncs()

	once := sync.Once{}
	s.window.ConnectShowEvent(func(e *gui.QShowEvent) {
		once.Do(s.login)
	})

	s.window.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		s.db.Close()
	})
	s.window.Show()
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
	fmt.Println(filename, title)
	if first {
		s.db.AddDiary(id, time.Now().Format("2006-01-02"), title, filename)
	} else {
		s.db.UpdateDiaryTitle(id, title)
	}

	encodeToFile(s.editor.ToHtml(), filename, s.key)
	s.setStatusBar(T("Save Diary"))
	curDiary.Modified = false
}

func (s *myWindow) clearEditor() {
	s.editor.SetText("")
}

func (s *myWindow) addDiary(yearMonth, day, title string) {
	if curDiary.Item != nil {
		s.saveCurDiary()
	}

	month := s.addYearMonth(yearMonth)
	diary := gui.NewQStandardItem2(fmt.Sprintf("%s-%s", day, title))
	diary.SetEditable(false)
	diary.SetAccessibleText("")

	//month.AppendRow2(diary)
	month.InsertRow2(0, diary)
	s.tree.SetCurrentIndex(diary.Index())

	curDiary.Item = diary
	curDiary.Day = day
	curDiary.YearMonth = yearMonth

	//s.editor.SetText(title + "\n")
	s.clearEditor()
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

func (s *myWindow) setTreeFuncs() {

	s.tree.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)

	s.tree.ConnectActivated(func(idx *core.QModelIndex) {
		fmt.Println(idx.Row(), idx.Column())
		p := idx.InternalPointer()
		item := gui.NewQStandardItemFromPointer(p)
		diary := item.Child(idx.Row(), idx.Column())

		if diary.AccessibleDescription() == "0" {
			filename := diary.AccessibleText() + ".dat"
			txt, err := decodeFromFile(filename, s.key)
			if err != nil {
				log.Println(err)
				if curDiary.Item != nil {
					s.saveCurDiary()
				}
				curDiary.Item = diary
				curDiary.Modified = false
				curDiary.YearMonth = item.Text()
				vs := strings.SplitN(diary.Text(), "-", 2)
				curDiary.Day = vs[0]
			} else {
				if curDiary.Item != nil {
					s.saveCurDiary()
				}
				curDiary.Item = diary
				curDiary.Modified = false
				curDiary.YearMonth = item.Text()
				vs := strings.SplitN(diary.Text(), "-", 2)
				curDiary.Day = vs[0]
				//s.editor.SetHtml(txt)
				s.editor.Document().SetHtml(txt)
				s.editor.Document().SetModified(false)
			}
		}

	})
	s.tree.ConnectMouseReleaseEvent(func(e *gui.QMouseEvent) {
		fmt.Println(e.Button())
		idx := s.tree.IndexAt(e.Pos())
		fmt.Println(s.model.ItemFromIndex(idx).Text())
	})
}

func (s *myWindow) setTitle(v string) {

	//cursor.BeginEditBlock()
	s.editor.MoveCursor(gui.QTextCursor__Start, gui.QTextCursor__MoveAnchor)
	cursor := s.editor.TextCursor()
	var cfmt = gui.NewQTextCharFormat()
	cfmt.SetFontPointSize(18)

	cfmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__blue), core.Qt__SolidPattern))

	cursor.MergeBlockCharFormat(cfmt)

	cursor.InsertText(v)

	//cursor.EndEditBlock()

	cursor.InsertText("\n\n")

	//cursor.BeginEditBlock()

	var afmt = gui.NewQTextCharFormat()
	afmt.SetForeground(gui.NewQBrush3(gui.NewQColor2(core.Qt__black), core.Qt__SolidPattern))
	afmt.SetFontPointSize(14)
	cursor.MergeBlockCharFormat(afmt)

	//cursor.EndEditBlock()
}

func (s *myWindow) setEditorFuncs() {
	s.editor.ConnectTextChanged(func() {
		curDiary.Modified = true
		v := s.editor.ToPlainText()
		secs := strings.SplitN(v, "\n", 2)
		disp1 := fmt.Sprintf("%s-%s", curDiary.Day, secs[0])
		disp0 := curDiary.Item.Text()
		if disp0 != disp1 {
			p := curDiary.Item.Parent()
			pos := curDiary.Item.Index()
			curDiary.Item = p.TakeChild(pos.Row(), pos.Column())
			curDiary.Item.SetText(disp1)
			p.SetChild(pos.Row(), pos.Column(), curDiary.Item)
		}
	})
}

func (s *myWindow) login() {
	dlg := widgets.NewQDialog(s.window, core.Qt__Dialog)
	dlg.SetWindowTitle(T("Login"))

	grid := widgets.NewQGridLayout(dlg)

	name := widgets.NewQLabel2(T("Name:"), dlg, core.Qt__Widget)
	grid.AddWidget(name, 0, 0, 0)

	nameInput := widgets.NewQLineEdit(dlg)
	grid.AddWidget(nameInput, 0, 1, 0)

	pwd := widgets.NewQLabel2(T("Password:"), dlg, core.Qt__Widget)
	grid.AddWidget(pwd, 1, 0, 0)

	pwdInput := widgets.NewQLineEdit(dlg)
	pwdInput.SetEchoMode(widgets.QLineEdit__Password)
	pwdInput.SetPlaceholderText(T("Length Must >= 4"))
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

	okBtn.ConnectClicked(func(b bool) {
		var err error
		if len(pwdInput.Text()) < 3 || len(nameInput.Text()) < 1 {
			return
		}
		s.db, err = getMyDb(nameInput.Text())
		if err != nil {
			panic(err)
		}
		s.key, err = s.db.GetRealKey(pwdInput.Text())
		if err != nil {
			panic(err)
		}

		s.window.SetWindowTitle(nameInput.Text())
		dlg.Hide()

	})

	regBtn.ConnectClicked(func(b bool) {
		if len(pwdInput.Text()) < 3 || len(nameInput.Text()) < 1 {
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

		s.window.SetWindowTitle(nameInput.Text())
		dlg.Hide()

	})

	cancelBtn.ConnectClicked(func(b bool) {
		dlg.Destroy(true, true)
		s.window.Destroy(true, true)
		s.app.Quit()
	})

	dlg.ConnectHideEvent(func(e *gui.QHideEvent) {
		log.Println("load list")
		s.addYearMonthsFromDb()
		dlg.Destroy(true, true)
	})

	dlg.Show()
}

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := new(myWindow)
	window.Create(app)
	res := app.Exec()
	os.Exit(res)
}
