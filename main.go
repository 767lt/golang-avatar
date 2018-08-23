package main

import (
	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
	"log"
	"flag"
)

func main() {

	var rectSize int
	var backgroundColor string
	var foregroundColor string
	var text string
	var fontSize int
	var fileName string
	var font string
	var alpha bool

	flag.IntVar(&rectSize, "rect-size", 512, "rectangle size")
	flag.BoolVar(&alpha, "alpha", false, "if set true use alpha background")
	flag.StringVar(&backgroundColor, "background-color", "#FFF", "background color hex value")
	flag.StringVar(&foregroundColor, "foreground-color", "#000", "foreground color hex value")
	flag.StringVar(&text, "text", "", "text to print")
	flag.IntVar(&fontSize, "font-size", 96, "font size in points")
	flag.StringVar(&fileName, "file-name", "out.png", "file name")
	flag.StringVar(&font, "font", "./fonts/Arial.ttf", "font path")

	flag.Parse()

	dc := gg.NewContext(rectSize, rectSize)

	if !alpha {
		fillBackground(dc, backgroundColor)
	}

	if err := dc.LoadFontFace(font, float64(fontSize)); err != nil {
		panic(err)
	}

	drawStringWithColor(dc, text, float64(rectSize), foregroundColor)

	dc.Clip()
	dc.SavePNG(fileName)
}

func setColorFromHex(context *gg.Context, hex string) {
	c, err := colorful.Hex(hex)
	if err != nil {
		log.Fatal(err)
	}
	context.SetColor(c)
}

func fillBackground(context *gg.Context, hex string) {
	setColorFromHex(context, hex)
	context.Clear()
}

func drawStringWithColor(context *gg.Context, text string, rectSize float64, hex string) {
	setColorFromHex(context, hex)
	drawString(context, text, rectSize)
}

func drawString(context *gg.Context, text string, rectSize float64) {
	context.DrawStringAnchored(text, rectSize/2, rectSize/2, 0.5, 0.5)
}
