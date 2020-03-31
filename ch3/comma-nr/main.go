package comma

import (
	"bytes"
)

func comma(s string) string {
	var b bytes.Buffer
	n := len(s)
	res := []byte(s)
	for i := 0; i < n-1; i++ {
		if i <= 3 {
			b.Write(byte(","))
		} else {
			b.Write(res[i])
		}
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
