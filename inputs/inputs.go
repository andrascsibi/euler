package inputs

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"strconv"
	"strings"
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

func BufReader(dir string, probNum, inputNum int) (*bufio.Reader, io.Closer) {
	path := getPath(dir, probNum, inputNum)
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(fi), fi
}

func Grid(dir string, probNum, n int) [][]int {
	br, closer := BufReader(dir, probNum, n)
	defer closer.Close()
	scanner := bufio.NewScanner(br)
	grid := make([][]int, n)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		grid[i] = make([]int, n)
		for j, valString := range strings.Split(line, " ") {
			val, _ := strconv.Atoi(valString)
			grid[i][j] = val
		}
	}
	return grid
}
