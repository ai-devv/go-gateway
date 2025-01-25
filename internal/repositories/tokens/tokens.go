package tokens

type Token string

type Repository interface {
	Check(token Token) bool
}
