package byten

import (
	"fmt"
	"log"
	"os"
)

func ExampleSize_SizeOfCurrentFile() {
	// get FileInfo
	file, err := os.Stat("size_example.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Size(file.Size()))
	// Output:
	// 237B
}
