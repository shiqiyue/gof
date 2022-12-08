package passwords

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type BCryptPasswordEncoder struct {
}

func (B BCryptPasswordEncoder) Encode(raw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
}

func (B BCryptPasswordEncoder) Match(raw, encode string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(encode), []byte(raw)); err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}
