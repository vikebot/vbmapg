package mapgen

import (
	"sort"
	"time"

	perlin "github.com/aquilax/go-perlin"
)

type Map [][]float64

type Biome []float64

// New generates a new map using the opensimplex algorithm
func New(width, height int, biome Biome) Map {
	pn := perlin.NewPerlin(2, 2, 3, time.Now().UnixNano())

	// allocate map
	m := make([][]float64, height)
	for i := 0; i < height; i++ {
		m[i] = make([]float64, width)
	}

	sort.Float64s(biome)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// receive noise for current block
			f := pn.Noise2D(float64(x)/float64(width), float64(y)/float64(height)) + 0.5

			// find the associated block to the noise
			for i := len(biome) - 1; i >= 0; i-- {
				if f > biome[i] {
					m[y][x] = biome[i]
					break
				}
			}
			if m[y][x] == 0 {
				m[y][x] = biome[0]
			}
		}
	}

	return m
}
