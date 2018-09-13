package main

import (
	"fmt"
	"strconv"
)

type View interface {
	Show()
	SetContent(string)
}

type TextBox struct {
	content string
}

func (txtBox TextBox) Show() {
	fmt.Println("Text Box content: ", txtBox.content)
}

func (txtBox *TextBox) SetContent(content string) {
	txtBox.content = content
}

type Model interface {
	Set(int)
	Get() int
}

type Database struct {
	data int
}

func (db *Database) Set(value int) {
	db.data = value
}

func (db Database) Get() int {
	return db.data
}

type Controller interface {
	Clicked(View, Model)
}

type TextController struct{}

func (txtController TextController) Clicked(view View, model Model) {
	fmt.Println("A button was clicked")
	model.Set(30)
	value := model.Get()
	content := strconv.Itoa(value)
	view.SetContent(content)
	view.Show()
}

func main() {
	fmt.Println("Minimalistic MVC Demo")
	webView := TextBox{}
	dataModel := Database{}
	defaultController := TextController{}
	defaultController.Clicked(&webView, &dataModel)
}
