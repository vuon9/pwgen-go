package main

import (
	"reflect"
	"testing"
)

func TestNewItems(t *testing.T) {
	type testCase struct {
		name string
		arg  []cmdable
		want []cmdable
	}

	testCases := []*testCase{
		{
			name: "create new cmdable slice properly",
			arg: []cmdable{
				NewBoolCommand(cmdHelp, "h", false, "Get help"),
				NewBoolCommand(cmdNoCapitalize, "A", false, "Don't include capital letters in the password"),
			},
			want: []cmdable{
				NewBoolCommand(cmdHelp, "h", false, "Get help"),
				NewBoolCommand(cmdNoCapitalize, "A", false, "Don't include capital letters in the password"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			actual := NewItems(testCase.arg...)
			if reflect.DeepEqual(testCase.want, actual) {
				tt.Errorf("NewItems(): want %v, actual %v", testCase.want, actual)
			}
		})
	}
}
