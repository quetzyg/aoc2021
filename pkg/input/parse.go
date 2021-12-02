package input

import (
	"bufio"
	"fmt"
	"io"
)

func Parse(reader io.Reader, split bufio.SplitFunc, callback func(string) error) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(split)

	line := 1

	for scanner.Scan() {
		err := callback(scanner.Text())

		if err != nil {
			return fmt.Errorf("on input line %d: %w", line, err)
		}

		line++
	}

	return scanner.Err()
}
