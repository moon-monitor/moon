package crypto

import (
	"database/sql/driver"
)

type String struct {
	original string
	hash     []byte
}

func (s *String) Scan(value interface{}) error {
	aes, err := WithAes()
	if err != nil {
		return err
	}
	decrypt, err := aes.Decrypt(value.([]byte))
	if err != nil {
		return err
	}
	s.hash = value.([]byte)
	s.original = string(decrypt)
	return nil
}

func (s String) Value() (driver.Value, error) {
	aes, err := WithAes()
	if err != nil {
		return nil, err
	}
	return aes.Encrypt([]byte(s.original))
}
