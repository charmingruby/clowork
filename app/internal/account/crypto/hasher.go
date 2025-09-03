package crypto

type HashComparer interface {
	Compare(plain, hash string) bool
}

type HashGenerator interface {
	Hash(plain string) (string, error)
}

type Hasher interface {
	HashComparer
	HashGenerator
}
