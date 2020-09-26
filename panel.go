package yuai

import "github.com/hajimehoshi/ebiten"

type Panel struct {
	desktop *Desktop
}

func (d *Desktop) CreatePanel() *Panel {
	result := &Panel{
		desktop: d,
	}

	return result
}

func (p Panel) Render(screen *ebiten.Image, x, y, width, height int) {
	bgColor := p.desktop.Config.Colors.PanelBackground

	DrawColoredRect(screen,
		ScaleToDevice(x), ScaleToDevice(y),
		ScaleToDevice(width), ScaleToDevice(height),
		bgColor[0], bgColor[1], bgColor[2], bgColor[3])
}

func (p Panel) Update() (dirty bool) {
	return false
}

func (p Panel) GetRequestedSize() (int, int) {
	return 1, 1
}

func (p Panel) Invalidate() {
}
