package main

import (
	"log"

	"github.com/essial/yuai"

	"github.com/hajimehoshi/ebiten"
)

type SimpleTest struct {
	desktop      *yuai.Desktop
	testbox      *yuai.VBox
	screenWidth  int
	screenHeight int
}

func (s *SimpleTest) Update(screen *ebiten.Image) error {
	s.desktop.Update()
	return nil
}

func (s *SimpleTest) Draw(screen *ebiten.Image) {
	s.desktop.Draw(screen)
}

func (s *SimpleTest) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.desktop.Layout(outsideWidth, outsideHeight)
}

func main() {
	ebiten.SetWindowTitle("Yuai Simple Test")
	ebiten.SetWindowSize(1280, 720)

	desktop, err := yuai.CreateDesktop(0, 0, 1280, 720, "../shared/config.json")
	if err != nil {
		log.Fatal(err)
	}

	app := &SimpleTest{
		desktop: desktop,
	}

	app.testbox = desktop.CreateVBox()
	app.desktop.SetPrimaryWidget(app.testbox)

	app.testbox.SetPadding(8)
	app.testbox.SetExpandChild(true)

	hbox := desktop.CreateHBox()
	hbox.SetExpandChild(true)
	hbox.AddChild(desktop.CreateButton("Left", func() {}))
	hbox.AddChild(desktop.CreateButton("Center", func() {}))
	hb1 := desktop.CreateButton("Right", func() {})
	hbox.AddChild(hb1)

	b1 := desktop.CreateButton("Align Middle", func() { app.testbox.SetAlignment(yuai.VAlignMiddle) })
	app.testbox.AddChild(desktop.CreateButton("Align Top", func() { app.testbox.SetAlignment(yuai.VAlignTop) }))
	app.testbox.AddChild(b1)
	app.testbox.AddChild(desktop.CreateButton("Align Bottom", func() { app.testbox.SetAlignment(yuai.VAlignBottom) }))
	app.testbox.AddChild(desktop.CreateButton("Vis Toggle", func() { b1.SetVisible(!b1.GetVisible()) }))
	app.testbox.AddChild(desktop.CreateButton("Toggle Expand Child", func() { app.testbox.SetExpandChild(!app.testbox.GetExpandChild()) }))
	app.testbox.AddChild(desktop.CreateButton("Child Spacing +", func() { app.testbox.SetChildSpacing(app.testbox.GetChildSpacing() + 1) }))
	app.testbox.AddChild(desktop.CreateButton("Child Spacing -", func() { app.testbox.SetChildSpacing(app.testbox.GetChildSpacing() - 1) }))
	app.testbox.AddChild(hbox)

	app.testbox.SetStretchComponent(b1)
	hbox.SetStretchComponent(hb1)

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
