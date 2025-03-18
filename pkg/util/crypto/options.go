package crypto

import (
	"github.com/moon-monitor/moon/pkg/config"
)

// WithIV sets the IV for the AES cipher.
func WithIV(iv []byte) AesOption {
	return func(a *aesImpl) {
		a.iv = iv
	}
}

// WithKey sets the key for the AES cipher.
func WithKey(key []byte) AesOption {
	return func(a *aesImpl) {
		a.key = key
	}
}

// WithMod sets the AES mode.
func WithMod(mod config.Crypto_AesConfig_MOD) AesOption {
	return func(a *aesImpl) {
		a.mod = mod
	}
}
