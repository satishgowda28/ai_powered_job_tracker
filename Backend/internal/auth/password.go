package auth

import "github.com/alexedwards/argon2id"

func Hashpassword(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

func ComparePasswordHash(password, hashPassword string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hashPassword)
	if err != nil {
		return false, err
	}
	if !match {
		return false, nil
	}

	return true, nil
}
