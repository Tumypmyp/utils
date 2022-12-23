package utils

import (
	"bufio"
	"log"
	"os"
)

// Head returns the top n lines og a file like the UNIX command.
func Head(path string) ([]string, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}
	log.Print(file)

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0, 2)
	for i := 0; scanner.Scan() && i < 2; i++ {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
