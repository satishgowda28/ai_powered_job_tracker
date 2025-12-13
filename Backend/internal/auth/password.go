package auth

import "github.com/alexedwards/argon2id"

// Hashpassword generates an Argon2id hash from the given plaintext password.
// It returns the encoded hash string and any error encountered while creating the hash.
func Hashpassword(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

// ComparePasswordHash compares the plaintext password with an Argon2id hash.
// It returns true if the password matches the hash, false otherwise.
// If an error occurs during verification it returns false and the error.
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