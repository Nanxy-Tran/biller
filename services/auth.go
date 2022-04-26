package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(rawPassword []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(rawPassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func IsPasswordMatched(comingPassword []byte, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, comingPassword)
}
