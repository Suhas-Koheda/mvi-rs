package hashpkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(unhashed string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(unhashed), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hashed string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
