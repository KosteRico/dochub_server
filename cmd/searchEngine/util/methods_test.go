package util

import (
	"fmt"
	"testing"
)

func TestDetectLanguage(t *testing.T) {
	fmt.Println(DetectLanguage("привет"))
}
