package inputs

import (
	"fmt"
	"io/ioutil"
)

func ToString(probNum, inputNum int) string {
	filename := fmt.Sprintf("%d_%d.txt", probNum, inputNum)
	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	return string(inputBytes)
}
