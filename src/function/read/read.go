package read

import (
	"bufio"
	"fmt"
	"os"
)

func GetMessage() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write a message:")
	input, err := reader.ReadString('\n')
	if err != nil {
		return _, err
	}
	return input, nil
}
