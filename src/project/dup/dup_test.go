package dup

import (
	"reflect"
	"testing"
)

func TestGetDuplicates(t *testing.T){
	var tests = []struct{
		input []string
		want []string
	}{
		{
			[]string{},
			[]string{},
		},
		{
			[]string{"one"},
			[]string{},
		},
		{
			[]string{"one", "two", "two"},
			[]string{"two"},
		},
	}

	for _,test := range tests {
		r := GetDuplicates(test.input)

		if !reflect.DeepEqual(r, test.want){
			t.Errorf("%s:  %s, want %s",
				test.input, r, test.want)
		}
	}
}
