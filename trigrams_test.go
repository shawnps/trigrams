package trigrams

import (
	"reflect"
	"testing"
)

func TestEnglish(t *testing.T) {
	m := map[string]int{"Thi": 1, "is ": 2, "h t": 1, " an": 1, " En": 1, "Eng": 1, "sh ": 1, "his": 1, "s a": 1, "n E": 1, "lis": 1, "tes": 1, "s i": 1, " is": 1, "an ": 1, "ngl": 1, "gli": 1, "ish": 1, " te": 1, "est": 1}
	in, out := "This is an English test", m
	if x, _ := GenerateTrigrams(in); reflect.DeepEqual(x, out) != true {
		t.Errorf("GenerateTrigrams(%v) = %v, want %v", in, x, out)
	}
}
