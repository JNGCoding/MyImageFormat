package Encoder

import (
	"encoding/binary"
	"image/color"
	"log"
	"os"

	"golang.org/x/image/bmp"
)

const (
	ROW_START_BYTE  = byte(0x00)
	LINE_START_BYTE = byte(0xFA)
	ROW_END_BYTE    = byte(0xFF)
	IMAGE_END_BYTE  = byte(0xFB)
)

type Line struct {
	ComparatorPixel color.Color
	Width           uint16
}

func Dummy() {}

func Encode(ImageFile *os.File, OutputFile *os.File) {
	/*
		Plan : (We are gonna store image data in the form of lines containing same pixels instead of storing each pixel individually.)
		1) Get the first pixel.
		2) Create a rectangle with pixel data = pixel.data, width = 1
		3) Goto the next pixel if pixel is same as first pixel then width += 1
		4) Write in bytes ---> Row Start Byte + Lines [ Stored as - Line Start Byte + Line Pixel Data (3 bytes) + Width Byte (2 Bytes) ] + Row End Byte
	*/

	MyImage, err := bmp.Decode(ImageFile)
	if err != nil {
		log.Fatalln("Failed to load Image file.")
	}

	// Writing the size of image.
	bounds := MyImage.Bounds()
	width := uint32(bounds.Dx())
	height := uint32(bounds.Dy())
	buf := make([]byte, 8)
	binary.BigEndian.PutUint32(buf[0:4], width)
	binary.BigEndian.PutUint32(buf[4:8], height)

	OutputFile.Write(buf) // Writing the size of the image.

	LineObject := Line{}
	FormingLine := false
	for row := 0; row < MyImage.Bounds().Dy(); row++ {
		OutputFile.Write([]byte{ROW_START_BYTE})
		for column := 0; column < MyImage.Bounds().Dx(); column++ {
			pixel := MyImage.At(column, row)
			PR, PG, PB, PA := pixel.RGBA()
			if !FormingLine {
				LineObject.ComparatorPixel = pixel
				LineObject.Width = 1
				FormingLine = true
				OutputFile.Write([]byte{LINE_START_BYTE})
				OutputFile.Write([]byte{byte(PR >> 8), byte(PG >> 8), byte(PB >> 8)})
			} else {
				CR, CG, CB, CA := LineObject.ComparatorPixel.RGBA()
				if CR == PR && CG == PG && CB == PB && CA == PA {
					LineObject.Width++
				} else {
					OutputFile.Write([]byte{byte(LineObject.Width >> 8 & 0xFF), byte(LineObject.Width & 0xFF)})
					FormingLine = false
					column-- // Reiterating to Create a new line.
				}
			}
		}

		// Trigger this if We reach the end of the row but Forming Line is true
		if FormingLine {
			OutputFile.Write([]byte{byte(LineObject.Width >> 8 & 0xFF), byte(LineObject.Width & 0xFF)})
			FormingLine = false
		}

		OutputFile.Write([]byte{ROW_END_BYTE})
	}

	OutputFile.Write([]byte{IMAGE_END_BYTE})
}
