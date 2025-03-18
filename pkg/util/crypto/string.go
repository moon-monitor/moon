package crypto

import (
	"database/sql/driver"
)

type String string

func (s String) EQ(a String) bool {
	return string(s) == string(a)
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
	*s = String(decrypt)
	return nil
}

func (s String) Value() (driver.Value, error) {
	aes, err := WithAes()
	if err != nil {
		return nil, err
	}
	return aes.Encrypt([]byte(s))
}
