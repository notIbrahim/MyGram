package helper

import "golang.org/x/crypto/bcrypt"

func LookupPassword(ReferencePassword string) string {
	Length := 6
	Password := []byte(ReferencePassword)

	Hash, _ := bcrypt.GenerateFromPassword(Password, Length)
	return string(Hash)
}

func PasswordCheck(ReferenceHash, ReferencePassword []byte) bool {
	p_Hash, p_Pass := []byte(ReferenceHash), []byte(ReferencePassword)
	err := bcrypt.CompareHashAndPassword(p_Hash, p_Pass)
	return err == nil
}
