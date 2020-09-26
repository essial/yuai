package yuai

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	buttonPaddingH = 32
	buttonPaddingV = 16
)

type Button struct {
	desktop            *Desktop
	caption            string
	textColor          color.Color
	disabledTextColor  color.Color
	fontBoundsX        int
	fontBoundsY        int
	hovered            bool
	canExecuteCallback bool
	dirty              bool
	enabled            bool
	visible            bool
	reqWidth           int
	reqHeight          int
	onClick            func()
}

func (d *Desktop) CreateButton(caption string, onClick func()) *Button {
	tc := d.Config.Colors.Text
	dtc := d.Config.Colors.DisabledText

	result := &Button{
		desktop:            d,
		caption:            caption,
		hovered:            false,
		visible:            true,
		canExecuteCallback: true,
		enabled:            true,
		dirty:              false,
		textColor:          color.RGBA{R: tc[0], G: tc[1], B: tc[2], A: tc[3]},
		disabledTextColor:  color.RGBA{R: dtc[0], G: dtc[1], B: dtc[2], A: dtc[3]},
		onClick:            onClick,
	}

	result.Invalidate()

	return result
}

func (b *Button) Render(screen *ebiten.Image, x, y, width, height int) {
	if width <= 0 || height <= 0 || !b.visible {
		return
	}

	primaryColor := b.desktop.Config.Colors.Primary
	textColor := b.textColor

	if b.enabled {
		mouseX, mouseY := ebiten.CursorPosition()
		b.hovered = false
		if b.canExecuteCallback && mouseX >= ScaleToDevice(x) && mouseX < ScaleToDevice(x+width) &&
			mouseY >= ScaleToDevice(y) && mouseY < ScaleToDevice(y+height) {
			primaryColor = b.desktop.Config.Colors.PrimaryHighlight
			b.hovered = true
		}
	} else {
		primaryColor = b.desktop.Config.Colors.Disabled
		textColor = b.disabledTextColor
	}

	DrawColoredRect(
		screen,
		ScaleToDevice(x), ScaleToDevice(y),
		ScaleToDevice(width), ScaleToDevice(height),
		primaryColor[0], primaryColor[1], primaryColor[2], primaryColor[3])

	font := b.desktop.normalFont
	heightDelta := int(float64(ScaleToDevice(height)-b.fontBoundsY) * 0.30)
	offsetX := ScaleToDevice(x+(width/2)) - (b.fontBoundsX / 2)
	offsetY := ScaleToDevice(y) + b.fontBoundsY + heightDelta

	text.Draw(screen, b.caption, font, offsetX, offsetY, textColor)
}

func (b *Button) Update() (dirty bool) {
	dirty = b.dirty

	if b.dirty {
		b.dirty = false
	}

	if !b.enabled || !b.visible {
		return dirty
	}

	if b.canExecuteCallback {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if b.hovered {
				b.onClick()
			}
			b.canExecuteCallback = false
		}
	} else {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			b.canExecuteCallback = true
		}
	}

	return dirty
}

func (b *Button) GetRequestedSize() (int, int) {
	if !b.visible {
		return 0, 0
	}
	return b.reqWidth, b.reqHeight
}

func (b *Button) Invalidate() {
	font := b.desktop.normalFont
	b.fontBoundsX, b.fontBoundsY = CalculateBounds(b.caption, font)
	b.reqWidth = UnscaleFromDevice(b.fontBoundsX) + (buttonPaddingH * 2)
	b.reqHeight = UnscaleFromDevice(b.fontBoundsY) + (buttonPaddingV * 2)
}

func (b *Button) SetEnabled(enabled bool) {
	b.enabled = enabled
}

func (b *Button) IsEnabled() bool {
	return b.enabled
}

func (b *Button) SetVisible(visible bool) {
	b.visible = visible
	b.dirty = true
}

func (b *Button) GetVisible() bool {
	return b.visible
}
