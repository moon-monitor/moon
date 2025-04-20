package do

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TLS struct {
	ServerName string `json:"serverName"`
	ClientCert string `json:"clientCert"`
	ClientKey  string `json:"clientKey"`
}
