package yuai

import (
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

type Desktop struct {
	x              int
	y              int
	width          int
	height         int
	Config         UIConfig
	normalFont     font.Face
	symbolFont     font.Face
	monospaceFont  font.Face
	ttNormal       *truetype.Font
	ttMono         *truetype.Font
	ttSymbols      *truetype.Font
	screenWidth    int
	screenHeight   int
	mouseX, mouseY int
	widget         Widget
}

func CreateDesktop(x, y, width, height int, configFile string) (*Desktop, error) {
	result := &Desktop{
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}

	SetDeviceScale(ebiten.DeviceScaleFactor())

	result.Config.Load(configFile)
	if err := result.configureFonts(); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Desktop) configureFonts() error {
	var ttNormalBytes, ttMonoBytes, ttSymbolBytes []byte

	var err error

	// Load font data from the files into a byte array
	if ttNormalBytes, err = ioutil.ReadFile(d.Config.Fonts.Normal.Face); err != nil {
		return err
	}

	if ttSymbolBytes, err = ioutil.ReadFile(d.Config.Fonts.Symbols.Face); err != nil {
		return err
	}

	if ttMonoBytes, err = ioutil.ReadFile(d.Config.Fonts.Monospaced.Face); err != nil {
		return err
	}

	// Parse the TTF font files
	if d.ttNormal, err = truetype.Parse(ttNormalBytes); err != nil {
		return err
	}

	if d.ttSymbols, err = truetype.Parse(ttSymbolBytes); err != nil {
		return err
	}

	if d.ttMono, err = truetype.Parse(ttMonoBytes); err != nil {
		return err
	}

	// Generate the fonts
	d.RegenerateFonts()

	return nil
}

// RegenerateFonts regenerates the fonts (at startup and when dragging to a different DPI display).
func (d *Desktop) RegenerateFonts() {
	deviceScale := ebiten.DeviceScaleFactor()

	d.normalFont = truetype.NewFace(d.ttNormal, &truetype.Options{
		Size:    float64(d.Config.Fonts.Normal.Size),
		DPI:     96 * deviceScale,
		Hinting: font.HintingNone,
	})

	d.symbolFont = truetype.NewFace(d.ttSymbols, &truetype.Options{
		Size:    float64(d.Config.Fonts.Symbols.Size),
		DPI:     96 * deviceScale,
		Hinting: font.HintingNone,
	})

	d.monospaceFont = truetype.NewFace(d.ttMono, &truetype.Options{
		Size:    float64(d.Config.Fonts.Monospaced.Size),
		DPI:     96 * deviceScale,
		Hinting: font.HintingNone,
	})
}

func (d *Desktop) Draw(screen *ebiten.Image) {
	frameColor := d.Config.Colors.WindowBackground

	// Fill the window with the frame color
	_ = screen.Fill(color.RGBA{R: frameColor[0], G: frameColor[1], B: frameColor[2], A: frameColor[3]})

	d.widget.Render(screen, d.x, d.y, d.width, d.height)
}

func (d *Desktop) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Store off the screen size for easy access
	d.screenWidth = outsideWidth
	d.screenHeight = outsideHeight

	// Return the actual resolution, determined by the virtual size and screen device scale
	return ScaleToDevice(outsideWidth), ScaleToDevice(outsideHeight)
}

func (d *Desktop) Update() {
	deviceScale := ebiten.DeviceScaleFactor()
	d.mouseX, d.mouseY = ebiten.CursorPosition()

	// If the device scale has changed, we need to regenerate the fonts
	if deviceScale != GetLastDeviceScale() {
		SetDeviceScale(deviceScale)
		d.RegenerateFonts()
		d.widget.Invalidate()
	}

	d.widget.Update()
}

func (d *Desktop) SetPrimaryWidget(widget Widget) {
	d.widget = widget
}
