package benchmark

import (
	"fmt"
	"strconv"
)

func Int2String_Fmt(n int) string{
	return fmt.Sprintf("%d", n)
}

func Int2String_strconv(n int64) string{
	return strconv.FormatInt(n, 10)
}

func Int2String_strconv2(n int) string{
	return strconv.Itoa(n)
}