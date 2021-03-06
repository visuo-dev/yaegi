package main

type Coord struct {
	x, y int
}

func (c Coord) dist() int { return c.x*c.x + c.y*c.y }

func main() {
	o := &Coord{3, 4}
	println(o.dist())
}

// Output:
// 25
