package password

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const bcryptSchemeName = "BLF-CRYPT"

type Bcrypt struct {
	cost int
}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{cost: 12}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (b Bcrypt) Hash(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	pass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pass, b.cost)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("{%s}%s", bcryptSchemeName, hash), nil
}
