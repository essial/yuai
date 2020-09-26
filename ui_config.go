package yuai

import "C"
import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type UIConfig struct {
	Colors UIColorConfig `json:"colors"`
	Fonts  FontConfig    `json:"fonts"`
}

type UIColorConfig struct {
	WindowBackground []uint8 `json:"windowBackground"`
	Text             []uint8 `json:"text"`
	Primary          []uint8 `json:"primary"`
	PrimaryHighlight []uint8 `json:"primaryHighlight"`
	Disabled         []uint8 `json:"disabled"`
	DisabledText     []uint8 `json:"disabledText"`
	PanelBackground  []uint8 `json:"panelBackground"`
}

type FontConfig struct {
	Normal     FontItemConfig `json:"normal"`
	Symbols    FontItemConfig `json:"symbols"`
	Monospaced FontItemConfig `json:"monospaced"`
	Info       FontItemConfig `json:"info"`
}

type FontItemConfig struct {
	Face string `json:"face"`
	Size int    `json:"size"`
}

func (c *UIConfig) Load(fileName string) {
	var configBytes []byte
	var err error

	if configBytes, err = ioutil.ReadFile(fileName); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(configBytes, c); err != nil {
		log.Fatal(err)
	}
}
