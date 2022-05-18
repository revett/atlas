package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora/v3"
)

// Question asks the user a question and reads the input.
func Question(question string, prefix string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(question) // nolint:forbidigo
	s := fmt.Sprintf("> %s.", prefix)
	fmt.Print(aurora.Faint(s)) // nolint:forbidigo

	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input from use: %w", err)
	}

	answer = strings.ReplaceAll(answer, "\n", "")
	return answer, nil
}
