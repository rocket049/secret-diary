package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
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

var id int

func getNewId() (res int) {
	res = id
	id++
	return
}

func T(v string) string {
	return gettext.T(v)
}

type myWindow struct {
	window *widgets.QMainWindow
	tree   *widgets.QTreeView
	model  *gui.QStandardItemModel
	editor *widgets.QTextEdit
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

	s.window.Show()
}

func (s *myWindow) saveCurDiary() {
	if !curDiary.Modified {
		s.setStatusBar(T("No Diary Saved"))
		return
	}
	fmt.Println("save", curDiary.Item.AccessibleText(), s.editor.ToHtml())
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

	month := s.addYearMonty(yearMonth)
	diary := gui.NewQStandardItem2(fmt.Sprintf("%s-%s", day, title))
	diary.SetEditable(false)
	diary.SetAccessibleText(fmt.Sprintf("%d", getNewId()))

	month.AppendRow2(diary)

	curDiary.Item = diary
	curDiary.Day = day
	curDiary.YearMonth = yearMonth

	//s.editor.SetText(title + "\n")
	s.clearEditor()
	s.setTitle(title)

	s.tree.Expand(month.Index())
}

func (s *myWindow) addYearMonty(yearMonth string) *gui.QStandardItem {
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
	s.model.AppendRow2(item)
	return item
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

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := new(myWindow)
	window.Create(app)
	res := app.Exec()
	os.Exit(res)
}
