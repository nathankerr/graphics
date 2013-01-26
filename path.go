package graphics

type Point struct {
	x float32
	y float32
}

func (g *Graphic) NewPath() {
	g.cairo.newPath()
}

func (g *Graphic) ClosePath() {
	g.cairo.closePath()
}

func (g *Graphic) Arc(center Point, radius float32, angle1 float32, angle2 float32) {
	g.cairo.arc(center, radius, angle1, angle2)
}

func (g *Graphic) ArcNegative(center Point, radius float32, angle1 float32, angle2 float32) {
	g.cairo.arcNegative(center, radius, angle1, angle2)
}

func (g *Graphic) CurveTo(p1 Point, p2 Point, p3 Point) {
	g.cairo.curveTo(p1, p2, p3)
}

func (g *Graphic) LineTo(p Point) {
	g.cairo.lineTo(p)
}

func (g *Graphic) MoveTo(p Point) {
	g.cairo.moveTo(p)
}

func (g *Graphic) Rectangle(topLeft Point, width float32, height float32) {
	g.cairo.rectangle(topLeft, width, height)
}

func (g *Graphic) RelCurveTo(dp1 Point, dp2 Point, dp3 Point) {
	g.cairo.relCurveTo(dp1, dp2, dp3)
}

func (g *Graphic) RelLineTo(dp Point) {
	g.cairo.relLineTo(dp)
}

func (g *Graphic) RelMoveTo(dp Point) {
	g.cairo.relMoveTo(dp)
}
