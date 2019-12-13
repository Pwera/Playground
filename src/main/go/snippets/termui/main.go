package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"image"
	"log"
	"math"
	"math/rand"
	"time"
)

type WidgetCustomInterface interface {
	setText()
}

type MyParagraph struct {
	widgets.Paragraph
	Coords         *image.Point
	UpdateFunction func() string
}
type MyBarChart struct {
	widgets.BarChart
	Coords         *image.Point
	UpdateFunction func() []float64
}
type MyGauge struct {
	widgets.Gauge
	Coords         *image.Point
	UpdateFunction func() int
}
type MyList struct {
	widgets.List
	Coords         *image.Point
	UpdateFunction func() []string
}

func NewMyParagraph(coords *image.Point, x int, y int) *MyParagraph {
	fu := func() string {
		terminalX, terminalY := ui.TerminalDimensions()
		return fmt.Sprintf("[%d - %d]", terminalX, terminalY)
	}
	paragraph := MyParagraph{
		Paragraph:      *widgets.NewParagraph(),
		Coords:         coords,
		UpdateFunction: fu,
	}

	xcoord := coords.X
	ycoord := coords.Y
	paragraph.SetRect(xcoord, ycoord, coords.X+x, y)
	coords.X = coords.X + x
	return &paragraph
}
func (mnp *MyParagraph) setText() {
	mnp.Text = mnp.UpdateFunction() + "\n" + mnp.Text
}

func NewBarChart(coords *image.Point, x int, y int, updatefunction func() []float64) *MyBarChart {
	bc := *widgets.NewBarChart()
	bc.Data = []float64{3, 2, 5, 3, 9, 3}
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.Title = "Bar Chart"
	xcoord := coords.X
	ycoord := coords.Y
	bc.SetRect(xcoord, ycoord, coords.X+x, y)
	coords.X = coords.X + x
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen, ui.ColorYellow, ui.ColorMagenta, ui.ColorBlue, ui.ColorCyan}
	bc.LabelStyles = []ui.Style{ui.NewStyle(bc.BarColors[0]), ui.NewStyle(bc.BarColors[1]), ui.NewStyle(bc.BarColors[2]), ui.NewStyle(bc.BarColors[3]), ui.NewStyle(bc.BarColors[4]), ui.NewStyle(bc.BarColors[5])}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}
	return &MyBarChart{
		Coords:         coords,
		BarChart:       bc,
		UpdateFunction: updatefunction,
	}
}
func (mnp *MyBarChart) setText() {
	mnp.Data = mnp.UpdateFunction()
}

func NewGauge(coords *image.Point, x int, y int, updatefunction func() int) *MyGauge {
	g0 := *widgets.NewGauge()
	g0.Title = "Slim Gauge"
	xcoord := coords.X
	ycoord := coords.Y
	g0.SetRect(xcoord, ycoord, coords.X+x, y)
	coords.X = coords.X + x
	g0.Percent = 0
	g0.BarColor = ui.ColorRed
	g0.BorderStyle.Fg = ui.ColorWhite
	g0.TitleStyle.Fg = ui.ColorWhite
	if updatefunction == nil {
		updatefunction = func() int {
			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(30)
			g0.Percent += i
			return g0.Percent % 100
		}
	}

	return &MyGauge{
		Coords:         coords,
		Gauge:          g0,
		UpdateFunction: updatefunction,
	}
}
func (mnp *MyGauge) setText() {
	if mnp.UpdateFunction != nil {
		mnp.Percent = mnp.UpdateFunction()
	}
}

func NewList(coords *image.Point, x int, y int, updatefunction func() []string) *MyList {
	l := widgets.NewList()
	l.Title = "List"
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	xcoord := coords.X
	ycoord := coords.Y
	l.SetRect(xcoord, ycoord, coords.X+x, y)
	coords.X = coords.X + x
	return &MyList{
		Coords:         coords,
		UpdateFunction: updatefunction,
		List:           *l,
	}
}
func (mnp *MyList) setText() {
	mnp.Rows = append(mnp.Rows, mnp.UpdateFunction()[0])
}

var (
	m = make(map[string]WidgetCustomInterface)
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	coords := image.Point{X: 0, Y: 0}
	bytesupdatefunction := func() []float64 {
		//TODO: find better solution
		return []float64{math.Floor(rand.Float64()*10) / 10, math.Floor(rand.Float64()*10) / 10, math.Floor(rand.Float64()*10) / 10, math.Floor(rand.Float64()*10) / 10, math.Floor(rand.Float64()*10) / 10, math.Floor(rand.Float64()*10) / 10}
	}
	p := NewMyParagraph(&coords, 25, 25)
	bc := NewBarChart(&coords, 50, 25, bytesupdatefunction)
	gauge1 := NewGauge(&coords, 30, 5, nil)
	list := NewList(&coords, 10, 10, func() []string {
		return []string{"s1", "s2"}
	})

	m["P1"] = p
	m["P2"] = bc
	m["G1"] = gauge1
	m["L1"] = list

	refresh := func() {
		//x, y := ui.TerminalDimensions()
		for k := range m {
			m[k].setText()
		}
		//TODO: iterate over map and render items
		ui.Render(p, bc, gauge1, list)
	}
	go func() {
		time.Sleep(4 * time.Second)
		gauge1.setText()
	}()

	for {
		e := <-ui.PollEvents()

		switch e.ID {
		case "q", "<C-c>":
			return
		case "w":

			break
		case "j", "<Down>":
			list.ScrollDown()
			break
		case "k", "<Up>":
			list.ScrollUp()
			break
		}
		switch e.Type {
		case ui.ResizeEvent:
			//refresh()
			break
		}
		refresh()
	}
}
