package gen

import (
	"encoding/json"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/vikebot/vbcore"
	"github.com/vikebot/vbmapg/pkg/mapgen"
	"go.uber.org/zap"
)

// Create creates a map and stores it as an image
func Create(height, width int, log *zap.Logger) {
	m := mapgen.New(width, height, mapgen.Biome{
		vbcore.BlockWaterSeed,
		vbcore.BlockLightDirtSeed,
		vbcore.BlockGrassSeed,
		vbcore.BlockTreeSeed,
		vbcore.BlockLightMountainSeed,
		vbcore.BlockMountainSeed,
	})

	var img *image.RGBA
	img = image.NewRGBA(image.Rect(0, 0, width*10, height*10))

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			col := resolveClr(m[y][x])
			for yy := 10 * y; yy < 10*y+10; yy++ {
				for xx := 10 * x; xx < 10*x+10; xx++ {
					img.Set(xx, yy, col)
				}
			}
		}
	}

	filename := "temp_map"
	dir := "vbmapg"

	// create dir if not exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Error("gen: Unable to create dir", zap.Error(err))
		}
	}

	// clear directory
	err := removeContent(dir)
	if err != nil {
		log.Error("gen: Unable to remove content of dir", zap.String("dir", dir), zap.Error(err))
	}

	output, err := os.Create(path.Join(dir, filename+".jpg"))
	if err != nil {
		log.Error("gen: Unable to create file", zap.String("filename", filename), zap.Error(err))
		return
	}
	defer output.Close()

	err = jpeg.Encode(output, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		log.Error("gen: Unable to create image", zap.Error(err))
		return
	}

	matrix := make([][]string, height)
	for i := 0; i < height; i++ {
		matrix[i] = make([]string, width)
	}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			col := resolveClr(m[y][x])
			switch col {
			case vbcore.BlockWaterClr:
				matrix[y][x] = vbcore.BlockWater
			case vbcore.BlockLightDirtClr:
				matrix[y][x] = vbcore.BlockLightDirt
			case vbcore.BlockGrassClr:
				matrix[y][x] = vbcore.BlockGrass
			case vbcore.BlockTreeClr:
				matrix[y][x] = vbcore.BlockTree
			case vbcore.BlockLightMountainClr:
				matrix[y][x] = vbcore.BlockLightMountain
			case vbcore.BlockMountainClr:
				matrix[y][x] = vbcore.BlockMountain
			}
		}
	}

	data, err := json.Marshal(matrix)
	if err != nil {
		log.Warn("unable to marshal matrix", zap.Error(err))
	}

	ioutil.WriteFile(path.Join(dir, filename+".json"), data, 0644)
}

func resolveClr(code float64) color.RGBA {
	switch code {
	case vbcore.BlockWaterSeed:
		return vbcore.BlockWaterClr
	case vbcore.BlockLightDirtSeed:
		return vbcore.BlockLightDirtClr
	case vbcore.BlockGrassSeed:
		return vbcore.BlockGrassClr
	case vbcore.BlockTreeSeed:
		return vbcore.BlockTreeClr
	case vbcore.BlockLightMountainSeed:
		return vbcore.BlockLightMountainClr
	case vbcore.BlockMountainSeed:
		return vbcore.BlockMountainClr
	default:
		return vbcore.BlockGrassClr
	}
}

// removeContent removes everything what is inside of the given directory
func removeContent(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}

	return nil
}
