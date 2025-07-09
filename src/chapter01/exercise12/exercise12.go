// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//
// Modified by Andy Dalton to implement exercise solution.
//

// Lissajous generates GIF animations of random Lissajous figures.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type options struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func makeOptions(r *http.Request) *options {
	opt := options{
		cycles:  5,
		res:     0.001, // angular resolution
		size:    100,   // image canvas covers [-size..+size]
		nframes: 64,    // number of animation frames
		delay:   8,     // delay between frames in 10ms units
	}

	if r == nil {
		return &opt
	}

	if cycles := r.URL.Query().Get("cycles"); cycles != "" {
		opt.cycles, _ = strconv.Atoi(cycles)
	}

	if res := r.URL.Query().Get("res"); res != "" {
		opt.res, _ = strconv.ParseFloat(res, 64)
	}

	if size := r.URL.Query().Get("size"); size != "" {
		opt.size, _ = strconv.Atoi(size)
	}

	if nframes := r.URL.Query().Get("nframes"); nframes != "" {
		opt.nframes, _ = strconv.Atoi(nframes)
	}

	if delay := r.URL.Query().Get("delay"); delay != "" {
		opt.delay, _ = strconv.Atoi(delay)
	}

	return &opt
}

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w, makeOptions(r))
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	lissajous(os.Stdout, makeOptions(nil))
}

func lissajous(out io.Writer, opt *options) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: opt.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < opt.nframes; i++ {
		rect := image.Rect(0, 0, 2*opt.size+1, 2*opt.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(opt.cycles)*2.0*math.Pi; t += opt.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(opt.size+int(x*float64(opt.size)+0.5), opt.size+int(y*float64(opt.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, opt.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
