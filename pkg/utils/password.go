package utils

import "golang.org/x/crypto/bcrypt"

const DefaultCost = bcrypt.DefaultCost

// HashPassword hashes a plain text password.
func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with a plain text password.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}