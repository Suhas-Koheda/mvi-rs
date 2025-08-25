package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(unhashed string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(unhashed), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
