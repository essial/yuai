package yuai

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var imgSquare *ebiten.Image

func DrawColoredRect(target *ebiten.Image, x, y, w, h int, r, g, b, alpha uint8) {
	var err error

	if imgSquare == nil {
		imgSquare, err = ebiten.NewImage(1, 1, ebiten.FilterNearest)
		if err != nil {
			log.Panic(err)
		}

		if err = imgSquare.Fill(color.Black); err != nil {
			log.Panic(err)
		}
	}

	drawOptions := &ebiten.DrawImageOptions{}

	drawOptions.GeoM.Translate(float64(x)*(1/float64(w)), float64(y)*(1/float64(h)))
	drawOptions.GeoM.Scale(float64(w), float64(h))
	drawOptions.ColorM.Translate(float64(r)/255, float64(g)/255, float64(b)/255, float64(alpha)/255)

	_ = target.DrawImage(imgSquare, drawOptions)
}
