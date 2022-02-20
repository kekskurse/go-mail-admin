package password

import (
	"encoding/base64"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/argon2"
)

const argon2SchemeName = "ARGON2ID"

type Argon2 struct {
	saltSize                uint32
	time, memory, keyLenght uint32
	threads                 uint8
}

func NewArgon2() *Argon2 {
	return &Argon2{
		saltSize:  16,
		keyLenght: 32,
		threads:   2,
		memory:    64 * 1024,
		time:      3,
	}
}

func (s Argon2) Hash(password string) (string, error) {
	salt := make([]byte, s.saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	pass := []byte(password)

	hash := argon2.IDKey(pass, salt, s.time, s.memory, s.threads, s.keyLenght)

	hstr := base64.RawStdEncoding.EncodeToString(hash)
	sstr := base64.RawStdEncoding.EncodeToString(salt)

	return fmt.Sprintf("{%s}$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2SchemeName, argon2.Version, s.memory, s.time, s.threads, sstr, hstr), nil
}
