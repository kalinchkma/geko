package temp

import (
	"embed"
	"fmt"
)

//go:embed "sec"
var FS embed.FS

func ReadFile() {
	temp, err := FS.ReadFile("sec/tt.html")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(temp)
}
