package Decoder

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"os"
)

const (
	ROW_START_BYTE  = byte(0x00)
	LINE_START_BYTE = byte(0xFA)
	ROW_END_BYTE    = byte(0xFF)
	IMAGE_END_BYTE  = byte(0xFB)
)

func read_width_height(file *os.File) (uint32, uint32) {
	WidthBytes := make([]byte, 4)
	HeightBytes := make([]byte, 4)

	file.Read(WidthBytes)
	file.Read(HeightBytes)

	Width := binary.BigEndian.Uint32(WidthBytes)
	Height := binary.BigEndian.Uint32(HeightBytes)

	return Width, Height
}

func get_line_object_data(file *os.File) []byte {
	result := make([]byte, 5)
	file.Read(result)

	return result
}

func read_byte(file *os.File) byte {
	result := []byte{0}
	file.Read(result)
	return result[0]
}

func Dummy() {}

func Decode(InputFile *os.File) (image.Image, error) {
	/*
		Plan : (to decode)
		1) Get the Width, Height of image.
		2) Check the ROW_START_BYTE.
		3) Get the line object data (5 bytes forward).
		4) Check if the next byte is ROW_END_BYTE or LINE_START_BYTE.
		5) If LINE_START_BYTE --> goto Step(3)
		6) If ROW_END_BYTE --> Check if the byte ahead is IMAGE_END_BYTE or ROW_START_BYTE
		7) If ROW_START_BYTE --> goto Step(2)
		8) If IMAGE_END_BYTE --> return formed image
	*/

	W, H := read_width_height(InputFile)
	storingImage := image.NewRGBA(image.Rect(0, 0, int(W), int(H)))

	cur_x := 0
	cur_y := -1

	for {
		data := read_byte(InputFile)
		if data == IMAGE_END_BYTE {
			break
		} else if data == ROW_START_BYTE {
			// Read the row till ROW END BYTE
			for {
				sub_data := read_byte(InputFile)
				if sub_data == LINE_START_BYTE {
					line_data := get_line_object_data(InputFile)
					width := uint16(line_data[3])<<8 | uint16(line_data[4])&0xFF
					Linecolor := color.RGBA{uint8(line_data[0]), uint8(line_data[1]), uint8(line_data[2]), 255}
					// fmt.Printf("Width = %d\n", width)
					// fmt.Printf("Color - %d, %d, %d, %d\n", Linecolor.R, Linecolor.G, Linecolor.B, Linecolor.A)

					for range width {
						storingImage.SetRGBA(cur_x, cur_y, Linecolor)
						cur_x++
					}
				} else if sub_data == ROW_END_BYTE {
					cur_y++
					cur_x = 0
					// fmt.Printf("CurX, CurY - %d, %d\n", cur_x, cur_y)
					break
				} else {
					return nil, fmt.Errorf("error: unexpected byte %b", sub_data)
				}
			}
		} else {
			return nil, fmt.Errorf("error: unexpected byte %b", data)
		}
	}

	// fmt.Println("Decoding finished.")
	return storingImage, nil
}
