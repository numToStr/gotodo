package utils

import "golang.org/x/crypto/bcrypt"

// Password is used for generating and verfifying passwords
type Password struct {
	String string
}

// Generate return a hashed password
func (pwd *Password) Generate() string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd.String), 10)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

// Verify compares a hashed password with plaintext password
func (pwd *Password) Verify(hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd.String))
}
