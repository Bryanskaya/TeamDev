package utils

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

const SaltLen = 64
func rndSalt(size int) (salt string){
	genSalt := make([]byte, size)
	for i := 0; i < size; i++ {
		genSalt[i] = byte(rand.Intn(127 - 32) + 32)
	}
	return string(genSalt)
}

/* Hash password
Uses password to create hash. Generates random salt if it wasn't presented
*/
func HashPassword(password string, oldSalt ...string) (salt, hash string) {
	if len(oldSalt) > 0 {
		salt = oldSalt[0]
	} else {
		salt = rndSalt(SaltLen)
	}

	hashByte, _ := bcrypt.GenerateFromPassword([]byte(password + salt), bcrypt.DefaultCost)
	hash = string(hashByte)

	return
}

/* Compare password and hash
Confirming password with stored salt and hash
*/
func CmpPassword(password, salt, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password + salt))
	return err == nil
}