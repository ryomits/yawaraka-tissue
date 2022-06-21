package apikey_test

import (
	"testing"

	"yawaraka-tissue/domain/apikey"
)

func TestKeyVerifier_Verify(t *testing.T) {
	type args struct {
		validMinute int32
		issuer      string
		allowed     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				validMinute: 10,
				issuer:      "http://test.example.com",
				allowed: []string{
					"http://test.example.com",
				},
			},
			wantErr: false,
		},
		{
			name: "Expired",
			args: args{
				validMinute: -1,
				issuer:      "http://test.example.com",
				allowed: []string{
					"http://test.example.com",
				},
			},
			wantErr: true,
		},
		{
			name: "UnknownIssuer",
			args: args{
				validMinute: 10,
				issuer:      "http://test.example.com",
				allowed: []string{
					"http://example.com",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kid := "kid1"
			secrets := map[string]string{kid: "test"}
			e := apikey.NewEncoder(secrets, tt.args.issuer)
			st, err := e.Encode(kid, tt.args.validMinute)
			if err != nil {
				t.Fatalf("NewAPIKeyEncoder.Encode() failed to encode token err=%v", err)
			}

			kv := apikey.NewVerifier(secrets, tt.args.allowed)
			if err := kv.Verify(st); (err != nil) != tt.wantErr {
				t.Errorf("APIKeyVerifier.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	t.Run("InvalidSign", func(t *testing.T) {
		kid := "kid1"
		secrets := map[string]string{kid: "test"}
		e := apikey.NewEncoder(secrets, "http://example.com")
		st, err := e.Encode(kid, 10)
		if err != nil {
			t.Fatalf("NewAPIKeyEncoder.Encode() failed to encode token err=%v", err)
		}

		kid2 := "kid2"
		secret2 := map[string]string{kid2: "test"}
		kv := apikey.NewVerifier(secret2, []string{"http://example.com"})
		if err := kv.Verify(st); (err != nil) != true {
			t.Errorf("APIKeyVerifier.Verify() error = %v, wantErr %v", err, true)
		}
	})
}
