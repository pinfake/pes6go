package network

var key = []byte{0xa6, 0x77, 0x95, 0x7c}

func Mutate(data []byte) ([] byte) {
    decoded := []byte{};
    i := 0;
    j := 0;
    for i < len(data) {
        decoded = append(decoded, data[i]^key[j]);
        j++;
        i++;
        if j%4 == 0 {
            j = 0;
        }
    }
    return decoded;
}
