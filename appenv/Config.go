package appenv

import (
	"encoding/hex"
	"os"

	"github.com/google/uuid"
)

func GetChipherKey() string {
	var key = os.Getenv("ES_DATA_COLLECTOR_CHIPHER_KEY")
	if key == "" {
		key = "es-data-collector-key-0123456789"
	}

	keyLength := 32
	if len(key) < keyLength {
		expectedLength := keyLength - len(key)
		key = key + uuid.NewString()[:expectedLength]
	}

	keyStr := hex.EncodeToString([]byte(key))
	return keyStr //encode key in bytes to string for saving

}
