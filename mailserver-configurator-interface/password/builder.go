package password

import (
	"fmt"
	"strings"
)

type PasswordHashBuilder interface {
	Hash(password string) (string, error)
}

func GetPasswordHashBuilder(hashType string) PasswordHashBuilder {
	switch strings.ToUpper(hashType) {
	case ssha512SchemeName:
		return NewSsha512()
	case argon2SchemeName:
		return NewArgon2()
	case bcryptSchemeName:
		return NewBcrypt()
	default:
		panic(fmt.Sprintf("%s hash not implemented", hashType))
	}
}
