package identicon

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

const (

	// Types
	Normal  int = iota
	Mirrorh     // Mirroring horizontal
	Mirrorv     // Mirroring vertical

	// Themes
	White
	Black
	Gray
	Free
)

type Identicon struct {
	Type   int
	Theme  int
	Q      int // Quantity px
	Div    int // Divisions
	Margin int // px
}

func New() *Identicon {
	rand.Seed(time.Now().UnixNano())
	id := new(Identicon)
	id.Type = Mirrorh
	id.Theme = White
	id.Q = 70
	id.Div = 5
	id.Margin = 35
	return id
}

func (id *Identicon) randRGB(pad int) [3]int {
	var r [3]int
	if id.Theme == Gray {
		ri := rand.Intn(255)
		for i := 0; i < 3; i++ {
			r[i] = ri
		}
	} else {
		for i := 0; i < 3; i++ {
			r[i] = rand.Intn(255)
		}
	}
	sum := r[0] + r[1] + r[2]
	if id.Theme == White && sum > (765-pad) {
		i := 0
		if r[0] < r[1] {
			i = 1
		}
		if r[i] < r[2] {
			i = 2
		}
		r[i] = r[i] - pad
	} else if id.Theme == Black && sum < pad {
		i := 0
		if r[0] > r[1] {
			i = 1
		}
		if r[i] > r[2] {
			i = 2
		}
		r[i] = r[i] + pad
	}
	return r
}

func (id *Identicon) generate() (image.Image, error) {
	if id.Type < Normal || id.Type > Mirrorv || id.Theme < White || id.Theme > Free || id.Q < 1 || id.Div < 1 || id.Margin < 0 {
		return nil, fmt.Errorf("%v", "Wrong fields")
	}
	side := id.Q*id.Div + id.Margin*2
	half := side / 2
	m := image.NewRGBA(image.Rect(0, 0, side, side))
	r := id.randRGB(5)
	fCol := color.RGBA{uint8(r[0]), uint8(r[1]), uint8(r[2]), 255}
	var bCol color.RGBA
	if id.Theme == White {
		bCol = color.RGBA{255, 255, 255, 255}
	} else if id.Theme == Black {
		bCol = color.RGBA{0, 0, 0, 255}
	} else {
		r = id.randRGB(0)
		bCol = color.RGBA{uint8(r[0]), uint8(r[1]), uint8(r[2]), 255}
	}
	if id.Margin > 0 {
		draw.Draw(m, image.Rect(m.Rect.Min.X, m.Rect.Min.Y, m.Rect.Max.X, m.Rect.Max.Y), &image.Uniform{bCol}, image.ZP, draw.Src)
	}
	colPair := []color.Color{fCol, bCol}
	for y := m.Rect.Min.Y + id.Margin; y < m.Rect.Max.Y-id.Margin; y += id.Q {
		for x := m.Rect.Min.X + id.Margin; x < m.Rect.Max.X-id.Margin; x += id.Q {
			dx := x + id.Q
			dy := y + id.Q
			var c color.Color
			if id.Type == Mirrorh && x >= half {
				c = m.At(side-dx, y)
			} else if id.Type == Mirrorv && y >= half {
				c = m.At(x, side-dy)
			} else {
				c = colPair[rand.Intn(2)]
			}
			if c == fCol {
				draw.Draw(m, image.Rect(x, y, dx, dy), &image.Uniform{c}, image.ZP, draw.Src)
			}
		}
	}
	return m, nil
}

func (id *Identicon) GeneratePNG(w io.Writer) error {
	m, err := id.generate()
	if err != nil {
		return err
	}
	return png.Encode(w, m)
}

func (id *Identicon) GeneratePNGToFile(p string) error {
	if _, err := os.Stat(path.Dir(p)); os.IsNotExist(err) {
		if err = os.MkdirAll(path.Dir(p), os.ModePerm); err != nil {
			return err
		}
	}
	file, err := os.Create(fmt.Sprintf("%v%v", p, ".png"))
	if err != nil {
		return err
	}
	defer file.Close()
	m, err := id.generate()
	if err != nil {
		return err
	}
	return png.Encode(file, m)
}

func (id *Identicon) GenerateRandomThemes(p string, n int) error {
	if n < 1 {
		n = 1
	}
	for i := 0; i < n; i++ {
		th := rand.Intn(3) - 1
		if th == 1 {
			th = rand.Intn(3) - 1
			if th == 1 {
				th = 2
			} else if th == 2 {
				th = 3
			}
		} else if th == 2 {
			th = 3
		}
		var t string
		switch th {
		default:
			id.Theme = Free
			t = "free"
		case 0:
			id.Theme = White
			t = "white"
		case 1:
			id.Theme = Black
			t = "black"
		case 2:
			id.Theme = Gray
			t = "gray"
		}
		s := p + strconv.Itoa(i+1) + "_" + t
		if err := id.GeneratePNGToFile(s); err != nil {
			return err
		}
	}
	return nil
}

func (id *Identicon) GenerateSequentialThemes(p string, n int) error {
	if n < 1 {
		n = 1
	}
	for j := White; j <= Free; j++ {
		id.Theme = j
		for i := 0; i < n; i++ {
			var t string
			switch j {
			default:
				t = "free"
			case White:
				t = "white"
			case Black:
				t = "black"
			case Gray:
				t = "gray"
			}
			s := p + strconv.Itoa(i+1) + "_" + t
			if err := id.GeneratePNGToFile(s); err != nil {
				return err
			}
		}
	}
	return nil
}
