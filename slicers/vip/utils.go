package vip

import (
	"bufio"
	. "github.com/l1va/goosli/primitives"
	. "github.com/l1va/goosli/slicers"
	"os"
)

//PrepareLayers add Walls and fill layers
func PrepareLayers(layers []Layer, settings Settings, planes []Plane) []Layer {
	addWalls := int(settings.WallThickness / settings.LineWidth)
	if addWalls > 0 {
		for i, layer := range layers { //TODO: in parallel
			for _, pt := range layer.Paths {
				if len(pt.Points) < 2 { //TODO: remove this
					continue
				}
				offs := offset(pt, addWalls, settings.LineWidth, layer.Norm)
				layers[i].MiddlePs = append(layers[i].MiddlePs, offs[:len(offs)-1]...)
				layers[i].InnerPs = append(layers[i].InnerPs, offs[len(offs)-1])
			}
		}
	}
	return FillLayers(layers, planes)
}

func offset(pth Path, addWalls int, nozzle float64, norm Vector) []Path {
	var res []Path
	res = append(res, MakeOffset(pth, nozzle, norm))
	for i := 1; i < addWalls; i++ {
		res = append(res, MakeOffset(res[i-1], nozzle, norm))
	}
	return res

}

func readPlanes(f string) []AnalyzedPlane {
	file, err := os.Open(f)
	if err != nil {
		println("File reading error: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := []AnalyzedPlane{}
	for scanner.Scan() {
		ans = append(ans, ParsePlane(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		println("Errors during reading: " + err.Error())
	}
	return ans
}
