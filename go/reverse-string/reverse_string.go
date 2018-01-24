package reverse

import "bytes"

func String(s string) string {
	var reversed = bytes.Buffer{}
	// Iterate the string backwards, pushing letter one at a time into the buffer
	for i:=len(s)-1; i >= 0; i-- {
		reversed.WriteByte(s[i])
	}
	return reversed.String()
}
