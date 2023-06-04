package DBReader

import (
	"errors"
	"os"
)

func OpenFile(str *string) []byte {
	data, err := os.ReadFile(*str)
	if err != nil {
		errors.New("not open")
	}
	return data
}
