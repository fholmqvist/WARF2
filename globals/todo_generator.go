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

/*
	Each TODO must be in the following format:
		/////////
		// TODO
		// [non-empty description]
		/////////
	Where the start and end slashes must be
	four or more in length.
*/
func GenerateTodos() {
	todos := []string{}
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if skip(err, info, path) {
			return nil
		}
		content, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		if !isTodo(content) {
			return nil
		}
		lines := split(content)
		for idx, line := range lines {
			if !strings.Contains(line, " TODO") {
				continue
			}
			desc := strings.TrimSpace(lines[idx+1])[3:]
			descIdx := 2
			current := lines[idx+descIdx]
			for !isEndMarker(current) {
				current = lines[idx+descIdx]
				if len(current) < 3 {
					descIdx++
					continue
				}
				desc += "\n\t" + strings.TrimSpace(current)[3:]
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

func skip(err error, info fs.FileInfo, path string) bool {
	if err != nil {
		return true
	}
	if info.IsDir() {
		return true
	}
	skips := []string{".git", "todo_generator"}
	for _, skip := range skips {
		if strings.Contains(path, skip) {
			return true
		}
	}
	return false
}

func isTodo(content []byte) bool {
	return bytes.Contains(content, []byte("// TODO"))
}

func isEndMarker(line string) bool {
	return strings.Contains(line, "////////")
}

func split(content []byte) []string {
	return strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
}
