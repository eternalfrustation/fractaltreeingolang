package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"github.com/fogleman/gg"
	"image/color"
	"math"
	"runtime/debug"
	"strconv"
)

var container *fyne.Container
var angletheonethatstarteditall float64
var lowestlen int
var height int
var width int
var refresh bool
var button *widget.Button
var button2 *widget.Button
var button3 *widget.Button
var button4 *widget.Button
var hbo *widget.Box
var angaccpos bool
var can fyne.Canvas
var slide *widget.Slider
var dc *gg.Context

func main() {
	const W = 1600
	const H = 900
	dc = gg.NewContext(W, H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	debug.SetGCPercent(5)
	red := new(color.RGBA)
	red.R = 255
	width = 500
	height = 500
	app := app.New()
	iterationsbox := widget.NewEntry()
	var err error
	container = fyne.NewContainer()
	w := app.NewWindow("Visualiser")
	can = w.Canvas()
	slide = widget.NewSlider(-2*math.Pi, 2*math.Pi)
	button = widget.NewButton("Render", func() {
		clear()
		lowestlens := iterationsbox.Text
		lowestlen, err = strconv.Atoi(lowestlens)
		angletheonethatstarteditall = slide.Value
		if !refresh {
			height = w.Content().Size().Height
			width = w.Content().Size().Width
		} else {
			height = dc.Height()
			width = dc.Width()
		}
		//fmt.Println(lowestlen)
		lineer(width/2, height, width/2, (height/2)+100)
	})
	drawregion := gg.NewContext(width, height)
	drawregion.SetRGBA(255, 255, 255, 1)
	drawregion.LineTo(0, 500)
	button2 = widget.NewButton("Clear", clear)
	iterationsbox.Show()
	//fmt.Println(err)
	slide.Step = 0.001
	label := widget.NewLabel("Constant Based mode")
	label.Show()
	button3 = widget.NewButton("Toggle Instant or Animated mode Currently Off", func() {
		if refresh {
			refresh = false
			button3.Text = "Toggle Instant or Animated mode Currently Instant"

		} else {
			refresh = true
			button3.Text = "Toggle Instant or Animated mode Currently Animated"
		}
	})
	button4 := widget.NewButton("Toggle Angle Mode", func() {
		if !angaccpos {
			angaccpos = true
			label.Text = "Position Based mode"
			label.Refresh()
		} else {
			angaccpos = false
			label.Text = "Constant Based mode"
			label.Refresh()
		}
	})
	w.Resize(fyne.NewSize(width, height))
	container.Show()
	hbo = widget.NewHBox(button4, button, button2, iterationsbox, label)
	hbo.Resize(fyne.NewSize(width, 100))
	hbo.Show()
	w.SetContent(widget.NewVBox(hbo, slide, button3, container))
	slide.OnChanged = func(val float64) {
		clear()
		refresh = false
		angletheonethatstarteditall = slide.Value
		height = w.Content().Size().Height
		width = w.Content().Size().Width
		lowestlen = 25
		lineer(width/2, height, width/2, (height/2)+100)
	}
	w.ShowAndRun()
	w.SetPadded(false)

}
func lineer(x1 int, y1 int, x2 int, y2 int) {
	hbo.Hide()
	slide.Hide()
	button3.Hide()
	line := canvas.NewLine(color.Black)
	frameno = 0
	red := new(color.RGBA)
	red.R = 100
	red.G = 50
	red.B = 45
	red.A = 125
	initangle := angletheonethatstarteditall
	container.Hidden = false
	line.Position1.X = int(x1)
	line.Position1.Y = int(y1)
	line.Position2.X = int(x2)
	line.Position2.Y = int(y2)
	line.StrokeColor = red
	line.StrokeWidth = 8
	line.Show()
	deltaangle := 0.0
	if angaccpos {
		deltaangle = map1(float64(y2+x2), 0, float64(height+width)*2/3, -0.8, 0.8)
	} else {
		deltaangle = math.Pi / 3
	}
	iter = 1
	defer ended()
	container.AddObject(line)
	if refresh {
		drawline(float64(x1), float64(y1), float64(x2), float64(y2), float64(red.R), float64(red.G), float64(red.B), float64(red.A), 8)
	}

	recurse(float64(line.Position2.X), float64(line.Position2.Y), initangle-deltaangle, float64(line.Position2.Y)*0.2)
	recurse(float64(line.Position2.X), float64(line.Position2.Y), initangle+deltaangle+math.Pi, float64(line.Position2.Y)*0.2)
	//fmt.Println("length")
}
func lineerc(x1 float64, y1 float64, x2 float64, y2 float64, r uint8, g uint8, b uint8, a uint8, width float32) {
	line := canvas.NewLine(color.Black)
	container.Hidden = false
	linecolor := new(color.RGBA)
	linecolor.R = r
	linecolor.G = g
	linecolor.B = b
	linecolor.A = a
	line.Position1.X = int(x1)
	line.Position1.Y = int(y1)
	line.Position2.X = int(x2)
	line.Position2.Y = int(y2)
	line.StrokeWidth = width
	line.StrokeColor = linecolor
	line.Show()
	container.AddObject(line)
}

var iter int

func recurse(x1 float64, y1 float64, angle float64, length float64) {
	x2 := x1 + length*math.Cos(angle)
	y2 := y1 + length*math.Sin(angle)
	iter++
	r := uint8(map1(float64(x1), 0, float64(width), 0, 255))
	g := uint8(map1(float64(x2), 0, float64(width), 0, 255))
	b := uint8(map1(float64(y1), 0, float64(height), 0, 255))
	a := uint8(map1(float64(y2+x2)/2, 0, float64(height+width)/2, 100, 255))
	thicc := float32(map1(length, 0, 125, 2, 12))
	lineerc(x1, y1, x2, y2, r, g, b, a, thicc)
	if refresh {
		container.Refresh()
		//time.Sleep(time.Millisecond * 125)
		drawline(float64(x1), float64(y1), float64(x2), float64(y2), float64(r), float64(g), float64(b), float64(a), float64(thicc))
	}
	deltaangle := 0.0
	if angaccpos {
		deltaangle = map1(float64(x2+y2), 0, float64(height+width)/2, -math.Pi/3, math.Pi/3)
	} else {
		deltaangle = math.Pi / 3
	}
	defer rerecurse(x2, y2, deltaangle, length, angle)

}

func clear() {
	for i := 0; i < len(container.Objects); i++ {
		container.Objects[i].Hide()
	}
}

func map1(value float64, istart float64, istop float64, ostart float64, ostop float64) float64 {
	return ostart + (ostop-ostart)*((value-istart)/(istop-istart))
}

var frameno int

func rerecurse(x2 float64, y2 float64, deltaangle float64, length float64, angle float64) {
	if length > float64(lowestlen) {
		recurse(x2, y2, angle-deltaangle, length*2/3)
		defer recurse(x2, y2, angle+deltaangle, length*2/3)
	}

}
func ended() {
	hbo.Show()
	slide.Show()
	button3.Show()
}
func drawline(x1, y1, x2, y2, r, g, b, a, w float64) {
	dc.SetRGBA(r/255, g/255, b/255, a/255)
	dc.SetLineWidth(w)
	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()
	frameno++
	dc.SavePNG(strconv.Itoa(frameno) + ".png")
}
