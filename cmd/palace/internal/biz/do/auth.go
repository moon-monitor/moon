package do

import (
	"database/sql/driver"
	"encoding/json"
)

var _ ORMModel = (*BasicAuth)(nil)

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *BasicAuth) Scan(src any) error {
	switch val := src.(type) {
	case []byte:
		return json.Unmarshal(val, b)
	case string:
		return json.Unmarshal([]byte(val), b)
	default:
		return nil
	}
}

func (b BasicAuth) Value() (driver.Value, error) {
	return json.Marshal(b)
}

var _ ORMModel = (*TLS)(nil)

type TLS struct {
	ServerName string `json:"serverName"`
	ClientCert string `json:"clientCert"`
	ClientKey  string `json:"clientKey"`
}

func (t *TLS) Scan(src any) error {
	switch val := src.(type) {
	case []byte:
		return json.Unmarshal(val, t)
	case string:
		return json.Unmarshal([]byte(val), t)
	default:
		return nil
	}
}

func (t TLS) Value() (driver.Value, error) {
	return json.Marshal(t)
}
