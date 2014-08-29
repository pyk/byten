package byten

import (
	"fmt"
	"math"
)

func index(s int64) float64 {
	x := math.Log(float64(s)) / math.Log(1024)
	return math.Floor(x)
}

// Size return a formated string from file size
//
// Size(6314666666666665984) # => 5.5EB
func Size(size int64) string {

	symbols := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	if size < 10 {
		return fmt.Sprintf("%dB", size)
	}
	i := index(size)
	val := float64(size) / math.Pow(1024, math.Floor(i))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+"%s", val, symbols[int(i)])
}
