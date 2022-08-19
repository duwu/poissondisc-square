package poissondisc

import (
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) IsCrossLine(q Point, r float64) bool {
	px1 := p.X - r
	px2 := p.X + r
	py1 := p.Y - r
	py2 := p.Y + r
	qx1 := q.X - r
	qx2 := q.X + r
	qy1 := q.Y - r
	qy2 := q.Y + r

	if math.Abs((px1+px2)/2-(qx1+qx2)/2) < ((px2+qx2-px1-qx1)/2) && math.Abs((py1+py2)/2-(qy1+qy2)/2) < ((py2+qy2-py1-qy1)/2) {
		return true

	}
	return false
}
