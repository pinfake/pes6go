package crypt

var xorMask = []byte{0xa6, 0x77, 0x95, 0x7c}

func ApplyMask(data []byte) []byte {
	decoded := []byte{}
	var i, j int = 0, 0
	for i < len(data) {
		decoded = append(decoded, data[i]^xorMask[j])
		j++
		i++
		if j%4 == 0 {
			j = 0
		}
	}
	return decoded
}
