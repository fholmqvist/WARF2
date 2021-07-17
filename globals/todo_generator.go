package globals

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GenerateTodos() {
	todos := []string{}
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if skip(path) {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		content, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		if !bytes.Contains(content, []byte("// TODO")) {
			return nil
		}
		lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
		for idx, line := range lines {
			if !strings.Contains(line, " TODO") {
				continue
			}
			desc := strings.TrimSpace(lines[idx+1])[3:]
			descIdx := 2
			for !strings.Contains(lines[idx+descIdx], "////") {
				if len(lines[idx+descIdx]) < 3 {
					descIdx++
					continue
				}
				desc += "\n\t" + strings.TrimSpace(lines[idx+descIdx])[3:]
				descIdx++
			}
			todo := fmt.Sprintf(
				"%v:%v:\n\t%v",
				path,
				idx,
				desc,
			)
			todos = append(todos, todo+"\n")
		}
		return nil
	})
	file, err := os.Create("todos.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.WriteString(file, strings.Join(todos, "\n"))
	if err != nil {
		panic(err)
	}
}

func skip(path string) bool {
	skips := []string{".git", "todo_generator"}
	for _, skip := range skips {
		if strings.Contains(path, skip) {
			return true
		}
	}
	return false
}
