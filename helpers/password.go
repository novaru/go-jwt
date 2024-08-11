package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
  passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  return string(passwordHash), err
}

func VerifyPassword(hash, password string) error {
  return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
