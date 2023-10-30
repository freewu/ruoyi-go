package lib

import (
	"fmt"
	"testing"
)

func TestStrToTime(t *testing.T) {
	time, err := StrToTime("2023-10-31 15:53:11")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("2023-10-31 15:53:11 = %#v\n", time)

	time, err = StrToTime("2023-10-31 05:05:05")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("2023-10-31 5:5:5 = %#v\n", time)
}
