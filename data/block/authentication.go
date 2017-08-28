package block

import (
	"bytes"

	"crypto/md5"

	"github.com/pinfake/pes6go/crypt"
)

type Authentication struct {
	Key          string
	PasswordHash []byte
	Password     string
	Unknown      []byte
	RosterHash   []byte
}

type AuthenticationInternal struct {
	data [74]byte
}

func (info Authentication) GetPasswordHash() []byte {
	var keypadded [36]byte
	copy(keypadded[:], []byte(info.Key))
	var buf bytes.Buffer
	buf.Write(keypadded[:])
	buf.Write([]byte(info.Password))
	md5sum := md5.Sum(buf.Bytes())
	return crypt.Encrypt(md5sum[:])
}

func (info Authentication) buildInternal() PieceInternal {
	var internal AuthenticationInternal
	var tmp [74]byte

	copy(tmp[:20], []byte(info.Key))
	copy(tmp[32:48], info.GetPasswordHash())
	copy(tmp[58:], info.RosterHash)
	copy(internal.data[:], crypt.Encrypt(tmp[:]))
	return internal
}

func NewAthentication(b *Block) Authentication {
	data := b.Body.GetBytes()
	deciphered := crypt.Decrypt(data)

	return Authentication{
		Key:          string(deciphered[:20]),
		PasswordHash: data[32:48],
		Unknown:      deciphered[48:58],
		RosterHash:   deciphered[58:74],
	}
}
