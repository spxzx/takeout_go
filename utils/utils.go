package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func StringsToInts(s string) []int {
	str := strings.Split(s, ",")
	res := make([]int, len(str))
	for i, v := range str {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func GenerateValidateCode() string {
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	return code
}
