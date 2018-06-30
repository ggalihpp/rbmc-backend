package example

import (
	"strconv"
)

// This layer wil handle the business proccess/any logic that happen in your app

func logicExample(input string) (int, error) {
	toNum, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return toNum + 42, nil
}
