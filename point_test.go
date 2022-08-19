package poissondisc

import "testing"

func TestPoint_IsCrossLine(t *testing.T) {
	isCrossLine := Point{6, 6}.IsCrossLine(Point{0, 18}, 6)
	t.Log(isCrossLine)
}
