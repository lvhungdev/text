package editor

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

func FromFile(path string) (Editor, error) {
	editor := Editor{
		path: path,
	}

	file, err := os.Open(path)
	if err != nil {
		editor.data = [][]rune{{}}
		return editor, nil
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var content []byte
	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return editor, err
			}
		}

		content = append(content, b)
	}

	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines {
		editor.data = append(editor.data, []rune(string(line)))
	}

	return editor, nil
}

func (e *Editor) save() error {
	if e.path == "" {
		return errors.New("editor.save: path is empty")
	}

	file, err := os.Create(e.path)
	if err != nil {
		return err
	}

	var content string
	for _, line := range e.data {
		content = content + string(line) + "\n"
	}

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}

	return writer.Flush()
}
