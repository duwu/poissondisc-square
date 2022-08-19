package main

import (
	"fmt"
	"github.com/duwu/poissondisc-square"
	"github.com/fogleman/gg"
	"log"
	"math"
)

const (
	W = 512
	H = 512
	R = 6
	K = 32
)

func main() {
	edge := 6
	points := poissondisc.Sample(float64(edge), float64(edge), float64(W-edge), float64(H-edge), float64(R), K, nil)
	fmt.Println(len(points), "points")

	cells := fenquan(points)
	dc := gg.NewContext(W, H)
	dc.SetRGB255(255, 255, 255)
	for _, cell := range cells {

		for _, point := range cell {
			dc.SetRGB(0, 0, 0)
			dc.DrawRectangle(point.X, point.Y, 1, 1)
			dc.Fill()

		}

	}
	dc.SavePNG("out.png")

}
func fenquan(points []poissondisc.Point) [][]poissondisc.Point {
	//log.Debug("points", points)
	result := make([][]poissondisc.Point, 0, len(points))
	for _, p := range points {
		arr := make([]poissondisc.Point, 144)

		//arr[0] = p
		//cutSlice(p, lands, float64(w))

		totalNum := 12 * 12
		tx := math.Floor(p.X)
		ty := math.Floor(p.Y)
		circle := 1
		var length float64 = 2
		count := -1
		for true {
			// 打印上边
			for j := ty; j < ty+length; j++ {
				count += 1
				p1 := poissondisc.Point{tx, j}
				arr[count] = p1
			}
			// 打印右边
			for i := tx + 1; i < tx+length; i++ {
				count += 1
				p1 := poissondisc.Point{i, ty + length - 1}
				arr[count] = p1
			}
			// 打印下边
			for j := ty; j < ty+length-1; j++ {
				count += 1
				p1 := poissondisc.Point{tx + length - 1, j}
				arr[count] = p1

			}
			for i := tx + 1; i < tx+length-1; i++ {
				count += 1
				p1 := poissondisc.Point{i, ty}
				arr[count] = p1
			}
			log.Printf("p:%v,circle:%d, length:%f,count:%d", p, circle, length, count)

			if count == totalNum-1 {
				break
			}
			circle++
			length += 2
			tx--
			ty--
		}
		result = append(result, arr)
	}
	//log.Debug("temps:%v", temps)

	return result
}
