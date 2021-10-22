package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Print("GoKalkulator")
	a := app.New()
	w := a.NewWindow("GoKalkulator")

	cont := widget.NewLabel("GoKalkulator")
	w.SetContent(container.NewVBox(cont,
		widget.NewButton("Go", func() {
			cont.SetText("GoNow")
		})))

	w.ShowAndRun()

}
