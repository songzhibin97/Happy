package randomCode

import (
	"fmt"
	"testing"
)

func TestGetCode(t *testing.T) {
	fmt.Println("code:", string(GetCode(2)))
}
