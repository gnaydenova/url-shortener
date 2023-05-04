package base62

const (
	base    uint = 62
	charSet      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// EncodeToString encodes a uint to a base62 string.
func EncodeToString(num uint) string {
	var encoded []byte
	for num > 0 {
		r := num % base
		num /= base
		encoded = append(encoded, charSet[r])
	}

	return string(encoded)
}
