package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string,error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyPassword(hashPassword, stringPassword string) error {
	return  bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(stringPassword))
}