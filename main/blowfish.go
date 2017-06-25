package main

import (
	"fmt"

	"crypto/md5"

	"github.com/andreburgaud/crypt2go/ecb"
	"golang.org/x/crypto/blowfish"
)

var blowfishKey = []byte{
	0x27, 0x50, 0x1f, 0xd0, 0x4e, 0x6b, 0x82, 0xc8,
	0x31, 0x02, 0x4d, 0xac, 0x5c, 0x63, 0x05, 0x22,
	0x19, 0x74, 0xde, 0xb9, 0x38, 0x8a, 0x21, 0x90,
	0x1d, 0x57, 0x6c, 0xbb, 0xe2, 0xf3, 0x77, 0xef,
	0x23, 0xd7, 0x54, 0x86, 0x01, 0x0f, 0x37, 0x81,
	0x9a, 0xfe, 0x6c, 0x32, 0x1a, 0x01, 0x46, 0xd2,
	0x15, 0x44, 0xec, 0x36, 0x5b, 0xf7, 0x28, 0x9a,
}

var cdKey = []byte{
	0x7b, 0xc1, 0xde, 0x08, 0xf6, 0x14, 0x7d, 0xbe,
	0x98, 0xfc, 0x68, 0x60, 0x93, 0xdb, 0xe6, 0x8c,
	0xeb, 0x30, 0x6d, 0xfb, 0x3b, 0xfe, 0x8b, 0x28,
	0x51, 0x39, 0xb7, 0x97, 0x66, 0x1b, 0x08, 0xe3,
}

var iv = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

var txt = []byte{
	0x71, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

var txth = []byte{
	0x76, 0x94, 0xf4, 0xa6, 0x63, 0x16, 0xe5, 0x3c,
	0x8c, 0xdd, 0x9d, 0x99, 0x54, 0xbd, 0x61, 0x1d,
}

var algo = []byte{
	'R', 'F', 'L', 'J', 'Y', '3', '4', 'D',
	'R', 'E', '9', '9', '3', 'H', 'X', '3',
	'E', 'R', '9', '4',
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	'q', 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

//var algo = []byte{
//	'R', 'F', 'L', 'J', 'Y', '3', '4', 'D',
//	'R', 'E', '9', '9', '3', 'H', 'X', '3',
//	'E', 'R', '9', '4', 0x00, 'q', 0x00,
//}

//var txt = []byte{
//	0x36, 0x1e, 0x4f, 0xc3, 0xb7, 0xbd, 0x38, 0x46,
//	0xdc, 0x94, 0x7c, 0x16, 0x65, 0xdc, 0x4b, 0x79,
//}

func main() {
	block, _ := blowfish.NewCipher(blowfishKey)
	mode := ecb.NewECBEncrypter(block)
	md5x := md5.Sum(algo)
	bytes := md5x[:]
	fmt.Printf("% x\n", bytes)
	dst := make([]byte, len(bytes))
	mode.CryptBlocks(dst, bytes)

	fmt.Printf("% x\n", dst)
}
