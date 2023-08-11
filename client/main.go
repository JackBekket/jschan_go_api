package client

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"jschan/app"
	"jschan/app/models"
)

type Gui struct {
	Thread *models.Post
	jschan *jschan.Client
	ctx    context.Context
	labels map[string]*widget.Label
}

func (x *Gui) newLabel(name string) *widget.Label {
	w := widget.NewLabel("")
	x.labels[name] = w
	return w
}

func NewGui(jschanUrl string) *Gui {
	return &Gui{
		ctx:    context.Background(),
		jschan: jschan.NewClient(jschanUrl),
		labels: make(map[string]*widget.Label),
	}
}

func (x *Gui) DataToScreen() {
	myType := reflect.TypeOf(x.Thread).Elem()
	myValue := reflect.ValueOf(x.Thread).Elem()
	for i := 0; i < myType.NumField(); i++ {
		tag := myType.Field(i).Tag.Get("json")
		ft := myType.Field(i).Type.String()
		switch ft {
		case "string":
			v := myValue.Field(i).String()
			x.labels[tag].SetText(v)
		}
	}
}

func (x *Gui) NewForm(w fyne.Window) fyne.Widget {
	form := &widget.Form{}
	tt := reflect.TypeOf(x.Thread).Elem()
	for i := 0; i < tt.NumField(); i++ {
		fld := tt.Field(i)
		tag := fld.Tag.Get("json")
		ft := fld.Type.String()
		switch ft {
		case "string":
			form.Append(fld.Name, x.newLabel(tag))
		}
	}
	return form
}

func (g *Gui) Refresh(w fyne.Window) {

	res, err := g.jschan.GetOverboardCatalog(g.ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	threadNum := rand.Intn(len(res.Threads))
	g.Thread = &res.Threads[threadNum]

	g.DataToScreen()
}

func Show(win fyne.Window) fyne.CanvasObject {

	g := NewGui("https://94chan.org")

	form := g.NewForm(win)

	refresh := widget.NewButton("Refresh", func() {
		go g.Refresh(win)
	})

	buttons := container.NewHBox(layout.NewSpacer(), refresh)

	return container.NewBorder(form, buttons, nil, nil)
}
