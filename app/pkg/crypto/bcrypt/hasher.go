package bcrypt

import "golang.org/x/crypto/bcrypt"

type Hasher struct{}

func NewHasher() Hasher {
	return Hasher{}
}

func (c Hasher) Compare(plain, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}

func (c Hasher) Hash(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
