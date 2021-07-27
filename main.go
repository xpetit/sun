package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/nathan-osman/go-sunrise"
	"github.com/xpetit/bar"
	"golang.org/x/term"
)

func main() {
	w, _, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	latitude := flag.Float64("lat", 15.978103, "Latitude")
	longitude := flag.Float64("long", -61.250442, "Longitude")
	flag.Parse()
	now := time.Now()
	y, m, d := now.Date()
	t1, t2 := sunrise.SunriseSunset(*latitude, *longitude, y, m, d)
	if now.Sub(t1) < 0 {
		now = t2
	}
	fmt.Println(bar.Draw(w, float64(now.Sub(t1)), float64(t2.Sub(t1))))
}
