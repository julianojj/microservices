package adapters

type Hash interface {
	Encrypt(password []byte) ([]byte, error)
	Decrypt(encryptedPassword []byte, password []byte) error
}
