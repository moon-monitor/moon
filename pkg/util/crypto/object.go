package crypto

import (
	"database/sql/driver"
	"encoding/json"
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
	decrypt, err := aes.Decrypt(value.([]byte))
	if err != nil {
		return err
	}
	return json.Unmarshal(decrypt, o)
}

func (o Object[T]) Value() (driver.Value, error) {
	aes, err := WithAes()
	if err != nil {
		return nil, err
	}
	bs, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	return aes.Encrypt(bs)
}
