package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"
)

const (
	addr = ":8000"
)

func main() {

	http.HandleFunc("/", ServeMandelbrot)

	fmt.Printf("Starting server on [%s]\n", addr)

	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	for range time.Tick(time.Second) {
		resp, err := http.Get("http://localhost:8000/")
		if err != nil {
			fmt.Println(err)
			continue
		}
		resp.Body.Close()
	}

}

func ServeMandelbrot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving [%s]\n", r.URL.String())

	// Source from "github.com/campoy/mandelbrot/mandelbrot"
	data := make([][]color.RGBA, 256)
	m := &img{256, 256, data}

	var wg sync.WaitGroup
	wg.Add(m.h * m.w)
	for i, row := range m.m {
		for j := range row {
			go func(i, j int) {
				fillPixel(m, i, j)
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()

	err := png.Encode(w, m)
	if err != nil {
		fmt.Println(err)
	}

}

type img struct {
	h, w int
	m    [][]color.RGBA
}

func fillPixel(m *img, i, j int) {
	// normalized from -2.5 to 1
	xi := 3.5*float64(i)/float64(m.w) - 2.5
	// normalized from -1 to 1
	yi := 2*float64(j)/float64(m.h) - 1

	const maxI = 1000
	x, y := 0., 0.
	for i := 0; (x*x+y*y < 4) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	paint(&m.m[i][j], x, y)
}

func paint(c *color.RGBA, x, y float64) {
	n := byte(x * y)
	c.R, c.G, c.B, c.A = n, n, n, 255
}

func (m *img) At(x, y int) color.Color { return m.m[x][y] }
func (m *img) ColorModel() color.Model { return color.RGBAModel }
func (m *img) Bounds() image.Rectangle { return image.Rect(0, 0, m.h, m.w) }
