package main

import (
	"ImageFormat/Decoder"
	"ImageFormat/Encoder"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	DecodedImage *image.Image
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	image__ := ebiten.NewImageFromImage(*g.DecodedImage)
	screen.DrawImage(image__, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// W, H := ebiten.WindowSize()
	return 1024, 1024
}

var IN_DECODING bool = true

func main() {
	fmt.Print("")
	Encoder.Dummy()
	Decoder.Dummy()

	ImageFile, err := os.Open("sample-1024x1024.bmp")
	if err != nil {
		log.Fatalln("Error opening image file.")
	}
	defer ImageFile.Close()

	OutputFile, err := os.OpenFile("Output-1024x1024.ifd", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln("Error Creating image file.")
	}
	defer OutputFile.Close()

	if IN_DECODING {
		var MyImage image.Image
		MyImage, err := Decoder.Decode(OutputFile)
		if err != nil {
			log.Fatalf("Failed Decoding Image. Error = %v\n", err)
		}

		Frame := &Game{}
		ebiten.SetWindowSize(1024/2, 1024/2)
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
		Frame.DecodedImage = &MyImage

		ebiten.SetWindowTitle("Ebitengine Sample")
		if err := ebiten.RunGame(Frame); err != nil {
			log.Fatal(err)
		}
	} else {
		Encoder.Encode(ImageFile, OutputFile)
	}
}
