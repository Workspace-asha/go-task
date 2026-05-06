package auth

import "testing"

func TestGenerateAndValidateToken(t *testing.T) {
	token, err := GenerateToken(1)
	if err != nil {
		t.Fatalf("failed to generate token")
	}

	_, err = ValidateToken(token)
	if err != nil {
		t.Fatalf("token should be valid")
	}
}