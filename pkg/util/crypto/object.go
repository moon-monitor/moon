package crypto

import (
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"

	"github.com/moon-monitor/moon/pkg/merr"
)

type Object[T any] struct {
	Data T `json:"data"`
}

func (o *Object[T]) Get() T {
	return o.Data
}

func (o *Object[T]) Scan(value interface{}) error {
	aes, err := WithAes()
	if err != nil {
		return err
	}
	var origin string
	switch val := value.(type) {
	case []byte:
		origin = string(val)
	case string:
		origin = val
	default:
		return merr.ErrorInternalServerError("invalid value type of crypto.Object")
	}
	decodedString, err := base64.StdEncoding.DecodeString(origin)
	if err != nil {
		return err
	}
	decrypt, err := aes.Decrypt(decodedString)
	if err != nil {
		return err
	}
	return json.Unmarshal(decrypt, o)
}

func (o Object[T]) Value() (driver.Value, error) {
	aes, err := WithAes()
	if err != nil {
		return "", err
	}
	bs, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	encrypt, err := aes.Encrypt(bs)
	if err != nil {
		return "", err
	}
	encodeToString := base64.StdEncoding.EncodeToString(encrypt)
	return encodeToString, nil
}
