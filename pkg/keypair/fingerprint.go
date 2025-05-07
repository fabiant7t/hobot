package keypair

import (
	"crypto/md5"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

func MD5Fingerprint(authorizedKey []byte) (string, error) {
	publicKey, _, _, _, err := ssh.ParseAuthorizedKey(authorizedKey)
	if err != nil {
		return "", fmt.Errorf("Failed to parse SSH public key: %w", err)
	}
	hash := md5.Sum(publicKey.Marshal())
	var sb strings.Builder
	for i, b := range hash {
		if i > 0 {
			sb.WriteString(":")
		}
		sb.WriteString(fmt.Sprintf("%02x", b))
	}
	return sb.String(), nil
}
