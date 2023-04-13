package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword 加密密码
func HashPassword(password []byte) (string, error) {
	hashPasswd, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashPasswd), nil
}

// ComparePassword 校验密码
func ComparePassword(hashPassword []byte, loginPassword []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashPassword, loginPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}
