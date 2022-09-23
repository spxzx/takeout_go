package md5

import (
	"crypto/md5"
	"fmt"
)

func Transfer(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
