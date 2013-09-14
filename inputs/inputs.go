package inputs

import (
	"fmt"
	"io/ioutil"
)

func ToString(probNum, inputNum int) (input string, err error) {
	filename := fmt.Sprintf("%d_%d.txt", probNum, inputNum)
	var inputBytes []byte
	inputBytes, err = ioutil.ReadFile(filename)
	input = string(inputBytes)
	return
}
