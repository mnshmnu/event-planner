package auth

type Auth interface {
	GenerateJWTToken(payload map[string]interface{}) (string, error)

	CompareHash(hPass []byte, pass []byte) (bool, error)
	GenerateHash(pass string) ([]byte, error)
}

type auth struct{}

func New() Auth {
	return &auth{}
}
