package rand

type KeySize int

const (
	KeySize16 KeySize = 16
	KeySize24 KeySize = 24
	KeySize32 KeySize = 32
)

// GenerateKey creates a new random secret key.
func GenerateKey(size KeySize) ([]byte, error) {
	return RandomBytes(int(size))
}

// GenerateKey creates a new random recommended secret key.
func GenerateRecommendedKey() ([]byte, error) {
	return GenerateKey(KeySize32)
}
