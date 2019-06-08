package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func (s *myWindow) showCategores(menubar *widgets.QMenuBar) {
	s.categories = menubar.AddMenu2(T("Categories"))
	newcategory := s.categories.AddAction(T("New Category"))
	newcategory.ConnectTriggered(func(b bool) {
		var ok bool
		name := widgets.QInputDialog_GetText(s.window, T("New Category Name"), T("Category"), 0, "", &ok, core.Qt__Dialog, 0)
		if ok {
			id := s.db.AddCategory(name)
			s.addCategoryMenuItem(id, name)
		}
	})

	renameCategory := s.categories.AddAction(T("Rename Category"))
	renameCategory.ConnectTriggered(func(b bool) {
		var ok bool
		name := widgets.QInputDialog_GetText(s.window, T("Rename Category")+" : "+s.categoryName, T("Rename Category")+" : "+s.categoryName, 0, "", &ok, core.Qt__Dialog, 0)
		if ok {
			c := s.categoryItem

			id, ok := c.Data().ToInterface().(int)
			if ok && id > 0 {
				s.db.RenameCategory(id, name)
				c.SetText(name)
				s.categoryName = name
				s.model.SetHorizontalHeaderLabels([]string{T("Diary List") + " - " + s.categoryName})
				s.modelFind.SetHorizontalHeaderLabels([]string{T("Result List") + " - " + s.categoryName})
			}
		}
	})

	removeCategory := s.categories.AddAction(T("Remove Category"))
	removeCategory.ConnectTriggered(func(b bool) {
		s.removeCurCategory()
	})

	s.categories.AddSeparator()
}

func (s *myWindow) loadUserCategorys() {
	c := s.addCategoryMenuItem(0, T("Default"))

	categoryDict := s.db.GetCategories()
	for id, name := range categoryDict {
		s.addCategoryMenuItem(id, name)
	}

	c.Triggered(true)
}

func (s *myWindow) addCategoryMenuItem(id int, name string) *widgets.QAction {
	c := s.categories.AddAction(name)
	c.SetCheckable(true)
	c.SetData(core.NewQVariant1(id))
	c.ConnectTriggered(func(b bool) {
		c.SetChecked(b)
		if b {
			data, ok := c.Data().ToInterface().(int)
			if ok == false {
				return
			}

			s.saveCurDiary()
			s.curDiary.Item = nil

			s.category = data
			s.categoryName = c.Text()
			if s.categoryItem != nil {
				s.categoryItem.SetChecked(false)
			}
			s.categoryItem = c
			s.showCurCategory(data)
		}
	})
	return c
}

func (s *myWindow) showCurCategory(c int) {
	s.db.setCategory(c)
	s.model.Clear()
	s.model.SetHorizontalHeaderLabels([]string{T("Diary List") + " - " + s.categoryName})
	s.modelFind.Clear()
	s.modelFind.SetHorizontalHeaderLabels([]string{T("Result List") + " - " + s.categoryName})

	s.addYearMonthsFromDb()
}

func (s *myWindow) removeCurCategory() {
	if s.category == 0 {
		widgets.QMessageBox_Information(s.window, T("Error"), T("Can not remove default category."), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		return
	}

	ret := widgets.QMessageBox_Question(s.window, T("Remove Category")+" : "+s.categoryName, T("Are you sure? All files in this category will move to default category."), widgets.QMessageBox__Yes|widgets.QMessageBox__No, widgets.QMessageBox__Yes)
	if ret == widgets.QMessageBox__No {
		return
	}
	s.db.RemoveCategory(s.category)
	s.categories.RemoveAction(s.categoryItem)
	s.categoryItem = nil

	for _, v := range s.categories.Children() {
		item := widgets.NewQActionFromPointer(v.Pointer())
		id, ok := item.Data().ToInterface().(int)
		if ok && id == 0 {
			item.Triggered(true)
		}
	}
}
