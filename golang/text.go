package golang

import (
	"bufio"
	"fmt"
	"os"
)

//按行写入文件
func WriteMaptoFile(m []string, filePath string) error {
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


//