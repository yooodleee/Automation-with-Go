package main 

import (
	"io"
	"os"
	"net/http"
)


func main() {
	res, err := http.Get("https://www.tkng.io/")
	if err != nil {
		panic(err)
	}

	src := res.Body
	defer src.Close()
	dst := os.Stdout

	io.Copy(dst, src)
}