/******
** @创建时间 : 2020/8/20 14:23
** @作者 : SongZhiBin
******/
package randomCode

import (
	"fmt"
	"testing"
)

func TestGetCode(t *testing.T) {
	fmt.Println("code:", string(GetCode(2)))
}
