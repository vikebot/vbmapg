package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strconv"

	fatih "github.com/fatih/color"
	"github.com/vikebot/vbmapg/pkg/mapgen"
)

const (
	Water     = 0.05
	Beach     = 0.15
	Grassland = 0.2
	Forest    = 0.6
	Mountain  = 0.7
	Snow      = 0.95
)

var (
	WaterClr     = color.RGBA{62, 96, 193, 0}
	BeachClr     = color.RGBA{93, 128, 253, 0}
	GrasslandClr = color.RGBA{116, 169, 99, 0}
	ForestClr    = color.RGBA{62, 126, 98, 0}
	MountainClr  = color.RGBA{130, 132, 128, 0}
	SnowClr      = color.RGBA{210, 210, 215, 0}
)

func main() {
	if len(os.Args) < 3 {
		println("To few arguments.\nUsage: mapgen <width> <height> [Optional: <filename.jpeg>]")
		os.Exit(-1)
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	m := mapgen.New(width, height, mapgen.Biome{
		Water,
		Beach,
		Grassland,
		Forest,
		Mountain,
		Snow,
	})

	var img *image.RGBA
	if len(os.Args) == 4 {
		img = image.NewRGBA(image.Rect(0, 0, width*10, height*10))
	}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if len(os.Args) == 4 {
				col := resolveClr(m[y][x])
				for yy := 10 * y; yy < 10*y+10; yy++ {
					for xx := 10 * x; xx < 10*x+10; xx++ {
						img.Set(xx, yy, col)
					}
				}
			} else {
				s := resolveFatih(m[y][x])
				fmt.Print(s + s)
				//fmt.Printf("%v ", m[y][x])
			}

		}
		if len(os.Args) != 4 {
			fmt.Printf("\n")
		}
	}

	if len(os.Args) == 4 {
		output, err := os.Create(os.Args[3])
		if err != nil {
			println(err.Error())
			os.Exit(-1)
		}
		defer output.Close()

		err = jpeg.Encode(output, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			println(err.Error())
			os.Exit(-1)
		}
	}
}

func resolveClr(code float64) color.RGBA {
	switch code {
	case Water:
		return WaterClr
	case Beach:
		return BeachClr
	case Grassland:
		return GrasslandClr
	case Forest:
		return ForestClr
	case Mountain:
		return MountainClr
	case Snow:
		return SnowClr
	default:
		return SnowClr
	}
}

func resolveFatih(code float64) string {
	switch code {
	case Water:
		return fatih.BlueString("█")
	case Beach:
		return fatih.CyanString("█")
	case Grassland:
		return fatih.GreenString("█")
	case Forest:
		return fatih.GreenString("█")
	case Mountain:
		return fatih.BlackString("█")
	case Snow:
		return fatih.WhiteString("█")
	default:
		return fatih.WhiteString("█")
	}
}
