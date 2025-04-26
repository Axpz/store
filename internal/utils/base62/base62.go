package base62

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Encode encodes an int64 into a base62 string.
func Encode(n int64) string {
	if n == 0 {
		return "0"
	}
	var encoded string
	for n > 0 {
		r := n % 62
		n = n / 62
		encoded = string(charset[r]) + encoded
	}
	return encoded
}

// Decode decodes a base62 string back into an int64.
func Decode(s string) int64 {
	var decoded int64
	for _, c := range s {
		decoded *= 62
		switch {
		case '0' <= c && c <= '9':
			decoded += int64(c - '0')
		case 'A' <= c && c <= 'Z':
			decoded += int64(c - 'A' + 10)
		case 'a' <= c && c <= 'z':
			decoded += int64(c - 'a' + 36)
		}
	}
	return decoded
}
