package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	timestamp := time.Now().Unix()
	filename := strconv.FormatInt(timestamp, 10) + ".txt"
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	for {
		mleft := robotgo.AddEvent("mleft")
		if mleft == 0 {
			x, y := robotgo.GetMousePos()
			color := robotgo.GetPixelColor(x, y)
			_, err = f.WriteString(fmt.Sprintf("%v %v\n", x, y))
			if err != nil {
				panic(err)
			}

			_, err = f.WriteString(fmt.Sprintf("%v\n", color))
			if err != nil {
				panic(err)
			}
		}
	}
}
