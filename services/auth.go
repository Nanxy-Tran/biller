package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(rawPassword string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func IsPasswordMatched(comingPassword string, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(comingPassword))
}
