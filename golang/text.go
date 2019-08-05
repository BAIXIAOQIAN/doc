package golang

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

//按行写入文件
func GoWriteMaptoFile(m []string, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, v := range m {
		fmt.Fprintln(w, v)
	}
	return w.Flush()
}

//读取csv文件
type CsvTable struct {
	FileName string
	Records  [][]string
}

//读取csv文件
func GoLoadCsvCfg(filename string) *CsvTable {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if reader == nil {
		return nil
	}
	records, err := reader.ReadAll()
	if err != nil {
		return nil
	}

	var result = &CsvTable{
		filename,
		records,
	}
	return result
}
