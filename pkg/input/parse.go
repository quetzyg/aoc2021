package input

import (
	"bufio"
	"io"
)

func Parse(reader io.Reader, split bufio.SplitFunc, callback func(string) error) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(split)

	for scanner.Scan() {
		err := callback(scanner.Text())

		if err != nil {
			return err
		}
	}

	return scanner.Err()
}
