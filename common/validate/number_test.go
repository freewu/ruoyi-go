package validate

import (
	"fmt"
	"regexp"
	"testing"
)

func TestValidateUintFormat(t *testing.T) {
	matched, err := regexp.MatchString(`^\d+$`, "20102")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)
}