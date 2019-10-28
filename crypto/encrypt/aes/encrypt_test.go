package aes

import (
	"bytes"
	"testing"

	"github.com/jinmukeji/go-pkg/crypto/rand"
)

func TestAESGCMEncrypt(t *testing.T) {

	// 生成一个 key
	key, _ := rand.GenerateRecommendedKey()

	type args struct {
		key            []byte
		plainText      []byte
		additionalData []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "0字节加密",
			args: args{
				key:            key[:],
				plainText:      []byte{},
				additionalData: nil,
			},
			wantErr: false,
		},
		{
			name: "短字节加密",
			args: args{
				key:            key[:],
				plainText:      []byte("hello"),
				additionalData: nil,
			},
			wantErr: false,
		},
		{
			name: "一般长度字节加密",
			args: args{
				key: key[:],
				plainText: []byte(
					`If AES is required or chosen, AES-GCM is often the best choice; 
					it pairs the AES block cipher with the GCM block cipher mode. It is an AEAD cipher: authenticated 
					encryption with additional data. It encrypts some data, which will be authenticated along with some 
					optional additional data that is not encrypted. The key length is 16 bytes for AES-128, 24 bytes 
					for AES-192, or 32 bytes for AES-256. It also takes a nonce as input, and the same caveats apply to 
					the nonce selection here. Another caveat is that GCM is difficult to implement properly, so it is 
					important to vet the quality of the packages that may be used in a system using AES-GCM.`),
				additionalData: nil,
			},
			wantErr: false,
		},
		{
			name: "带 additionalData 验证的加密",
			args: args{
				key: key[:],
				plainText: []byte(
					`If AES is required or chosen, AES-GCM is often the best choice; 
					it pairs the AES block cipher with the GCM block cipher mode. It is an AEAD cipher: authenticated 
					encryption with additional data. It encrypts some data, which will be authenticated along with some 
					optional additional data that is not encrypted. The key length is 16 bytes for AES-128, 24 bytes 
					for AES-192, or 32 bytes for AES-256. It also takes a nonce as input, and the same caveats apply to 
					the nonce selection here. Another caveat is that GCM is difficult to implement properly, so it is 
					important to vet the quality of the packages that may be used in a system using AES-GCM.`),
				additionalData: []byte("additional data"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AESGCMEncrypt(tt.args.key, tt.args.plainText, tt.args.additionalData)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESGCMEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			plainText, err := AESGCMDecrypt(tt.args.key, got, tt.args.additionalData)
			if err != nil {
				t.Errorf("AESGCMDecrypt() error = %v", err)
				return
			}

			if !bytes.Equal(plainText, tt.args.plainText) {
				t.Errorf("Original plain text = %v, but decrypted plain text %v", plainText, tt.args.plainText)
			}
		})
	}
}
