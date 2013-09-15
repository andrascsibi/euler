package inputs

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getPath(dir string, probNum, inputNum int) string {
	filename := fmt.Sprintf("%d_%d.txt", probNum, inputNum)
	return filepath.Join(dir, filename)
}

func String(dir string, probNum, inputNum int) string {
	path := getPath(dir, probNum, inputNum)
	inputBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(inputBytes)
}

func BufReader(dir string, probNum, inputNum int) *bufio.Reader {
	path := getPath(dir, probNum, inputNum)
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	// XXX who's gonna close the file?
	return bufio.NewReader(fi)
}
