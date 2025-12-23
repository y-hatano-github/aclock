package main

import (
	"log"
	"math"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	c "github.com/y-hatano-github/coordin"
)

const centerX, centerY, axisX, axisY = 20, 10, 20, 10

type model struct {
	Coords
	Args
	SystemBgColor string
}

type tickMsg time.Time

var colors = map[string]string{
	"system":  "x",
	"black":   "0",
	"red":     "1",
	"green":   "2",
	"yellow":  "3",
	"blue":    "4",
	"magenta": "5",
	"cyan":    "6",
	"white":   "7",
	"gray":    "8",
	"purple":  "13",
	"brown":   "130",
	"pink":    "205",
	"orange":  "214",
}

func (m model) tColor(name string) string {
	if c, ok := colors[name]; ok {
		if c == "x" {
			return m.SystemBgColor // system background color
		}
		return c
	}

	return m.SystemBgColor // default system background color
}

type Args struct {
	Background string `help:"Background color of the terminal area surrounding the clock." default:"system"`
	Face       string `help:"Color of the clock face. This is the filled area inside the frame." default:"gray"`
	Frame      string `help:"Color of the outer frame of the clock." default:"white"`
	Hour       string `help:"Color of the hour hand." default:"blue"`
	Min        string `help:"Color of the minute hand." default:"green"`
	Sec        string `help:"Color of the second hand." default:"cyan"`
	Piv        string `help:"Color of the pivot point." default:"white"`
	Tick       string `help:"Color of the tick marks." default:"red"`
}

type Coords struct {
	Fc c.Points // face
	Hm c.Points // HourMark
	Hh c.Points // hours hand
	Mh c.Points // minutes hand
	Sh c.Points // seconds hand
	Fr c.Points // frame
	Pp c.Points // pivot point
}

func main() {
	var args Args

	kong.Parse(&args,
		kong.Name("aclock"),
		kong.Description(`
A colorful analog clock rendered in your terminal.

You can customize the clock's appearance by specifying colors for:
 background
 face
 frame
 hour/minute/second hands
 pivot point
 tick marks

Colors available:
  black, red, green, yellow, blue, magenta, cyan, white
  gray, purple, brown, pink, orange
  system (uses terminal's background color)

Example:
  aclock --face blue --frame white --hour yellow --min green --sec red
   
Controls:
  ESC, Ctrl+C    Exit the application
  `),
	)
	sysbgc := termenv.BackgroundColor()
	p := tea.NewProgram(&model{Args: args, SystemBgColor: sysbgc.Sequence(true)})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m *model) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(200*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.shapeCoords()
		return m, tick()

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *model) View() string {
	width := centerX + axisX + 1
	height := centerY + axisY + 1

	v := make([][]string, width)
	for x := 0; x < width; x++ {
		v[x] = make([]string, height)
		for y := 0; y < height; y++ {
			v[x][y] = " "
		}
	}

	d := func(ps c.Points, c string, v *[][]string) {
		for _, p := range ps {
			if p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height {
				(*v)[p.X][p.Y] = c
			}
		}
	}

	d(m.Fc, m.tColor(m.Args.Face), &v)
	d(m.Hm, m.tColor(m.Args.Tick), &v)
	d(m.Hh, m.tColor(m.Args.Hour), &v)
	d(m.Mh, m.tColor(m.Args.Min), &v)
	d(m.Sh, m.tColor(m.Args.Sec), &v)
	d(m.Fr, m.tColor(m.Args.Frame), &v)
	d(m.Pp, m.tColor(m.Args.Piv), &v)

	p := termenv.ColorProfile()
	var b strings.Builder
	b.Grow(width * height * 10)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if v[x][y] == " " {
				b.WriteString(
					termenv.String(" ").Background(p.Color(m.tColor(m.Args.Background))).String(),
				)

			} else {
				b.WriteString(
					termenv.String(" ").Background(p.Color(v[x][y])).String(),
				)

			}
		}
		b.WriteByte('\n')

	}

	return b.String()
}

// calculate shape coordinates
func (m *model) shapeCoords() {
	t := time.Now()
	hand := func(unit, cx, cy, rx, ry, deg int) c.Points {
		d := float64(deg*unit - 90)
		x := float64(rx) * math.Cos(float64(d)*3.14/180)
		y := float64(ry) * math.Sin(float64(d)*3.14/180)

		return c.Line(c.Point{X: cx, Y: cy}, c.Point{X: cx + int(x), Y: cy + int(y)})
	}

	m.Hm = c.Circled(centerX, centerY, axisX-2, axisY-1, 30)                               // HourMark
	m.Hh = hand(t.Hour()*5+int(t.Minute()/12), centerX, centerY, axisX-6, axisY-3, 360/60) // hours hand
	m.Mh = hand(t.Minute(), centerX, centerY, axisX, axisY, 360/60)                        // minutes hand
	m.Sh = hand(t.Second(), centerX, centerY, axisX, axisY, 360/60)                        // seconds hand
	m.Fr, m.Fc = c.Circle(centerX, centerY, axisX, axisY)                                  // frame
	m.Pp = c.Points{c.Point{X: centerX, Y: centerY}}                                       // pivot point

}
