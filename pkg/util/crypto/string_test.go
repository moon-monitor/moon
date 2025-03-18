package crypto_test

import (
	"strings"
	"testing"

	"github.com/moon-monitor/moon/pkg/util/crypto"
)

func TestString_Scan_Success(t *testing.T) {
	var s crypto.String = "decrypted"
	val, err := s.Value()
	if err != nil {
		t.Fatalf("Expected no error, got %v\n", err)
	}

	t.Logf("val: %v", string(val.([]byte)))
	var got crypto.String
	if err := got.Scan(val); err != nil {
		t.Fatalf("Expected no error, got %v\n", err)
	}
	if !strings.EqualFold(string(s), string(got)) {
		t.Errorf("Expected '%v', got '%v'", s, got)
	}
}
