package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	ResultTab  = "ResultTab"
	AdvTimeTab = "AdvTimeTab"
)

var restabnew = []string{"Amount of edges", "Amount of vertex", "Path to graph", "Path to result", "time", "result", "mark"}
var restabold = []string{"Amount of edges", "Amount of vertex", "Path to graph", "Path to result", "time", "result"}
var advtime = []string{"size", "mark time", "overall"}

func main() {
	walkDir := os.Args[1]

	if err := filepath.Walk(walkDir, checkFile); err != nil {
		log.Panic(err)
	}
}

func checkFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	switch info.Name() {
	case ResultTab:
		translateToCsv(path, false)
	case AdvTimeTab:
		translateToCsv(path, true)
	default:
		return nil
	}

	return nil
}

func translateToCsv(path string, isTime bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	csvtab := make([][]string, 1)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tab := strings.Fields(scanner.Text())
		csvtab = append(csvtab, tab)
	}

	if isTime {
		csvtab[0] = advtime
	} else {
		if len(csvtab[1]) == 7 {
			csvtab[0] = restabnew
		} else {
			csvtab[0] = restabold
		}
	}

	newFile, err := os.Create(path + ".csv")
	if err != nil {
		return err
	}
	defer newFile.Close()

	writer := csv.NewWriter(newFile)
	writer.WriteAll(csvtab)
	return nil
}
