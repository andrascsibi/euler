package inputs

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ToString(dir string, probNum, inputNum int) string {
	filename := fmt.Sprintf("%d_%d.txt", probNum, inputNum)
	path := filepath.Join(dir, filename)
	inputBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	return string(inputBytes)
}
