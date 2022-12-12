package encryption

import (
	"proyectoarqsoft/dtos"

	"crypto/sha256"
	"encoding/hex"

	"github.com/golang-jwt/jwt"
)

func Signedlogintoken(cred dtos.Credentials) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usename":  cred.Username,
		"password": cred.Password,
	})

	// pasamos la llave que vamos a usar a un arreglo de bytes
	var key = []byte("soy una llave")
	// signedstring crea el kwt token y lo firma
	// recibe una llave como parametro
	jwtstring, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return jwtstring, nil
}

func GetSha256Hash(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
