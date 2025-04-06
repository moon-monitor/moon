package crypto

import (
	"database/sql/driver"
	"encoding/base64"
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
	if value == nil {
		*s = ""
		return nil
	}
	val := ""
	switch value.(type) {
	case string:
		val = value.(string)
		if len(val) == 0 {
			*s = ""
			return nil
		}
	case []byte:
		val = string(value.([]byte))
		if len(val) == 0 {
			*s = ""
			return nil
		}
	}
	decodedString, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return err
	}
	decrypt, err := aes.Decrypt(decodedString)
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
	encrypt, err := aes.Encrypt([]byte(s))
	if err != nil {
		return nil, err
	}
	encodeToString := base64.StdEncoding.EncodeToString(encrypt)
	return encodeToString, nil
}
