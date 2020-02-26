package transcoder

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func downloadTestFile(name, ext string) *os.File {
	resp, err := http.Get(name)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := ioutil.TempFile("./", "*."+ext)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
	return file
}
