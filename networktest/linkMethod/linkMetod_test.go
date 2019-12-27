package linkMethod

import (
	"fmt"
	"testing"
)

func TestPostUrl(t *testing.T) {
	fmt.Println(string(PostUrl("http://www.baidu.com")))
}
