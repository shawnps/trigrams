package trigrams

import (
	"errors"
	"fmt"
)

func GenerateTrigrams(input string) (map[string]int, error) {
	f := map[string]int{}
	s := []rune(input)
	if len(s) < 3 {
		return f, errors.New("String is less than 3 characters long!")
	}
	for i := 0; i <= len(s); i++ {
		if i == len(s)-2 {
			break
		}
		t := string(s[i]) + string(s[i+1]) + string(s[i+2])
		if _, ok := f[t]; ok {
			f[t] += 1
		} else {
			f[t] = 1
		}
	}
	return f, nil
}
