package keypair

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

func NewEd25519Keypair(comment, passphrase string) (authorizedKey []byte, privatePEM []byte, err error) {
	commentSanitized := strings.ReplaceAll(comment, "\n", " ")
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return authorizedKey, privatePEM, fmt.Errorf("error generating key: %w", err)
	}
	if passphrase == "" { // no passphrase
		block, err := ssh.MarshalPrivateKey(priv, commentSanitized)
		if err != nil {
			return authorizedKey, privatePEM, fmt.Errorf("error marshalling private key: %w", err)
		}
		privatePEM = pem.EncodeToMemory(block)
	} else { // with passphrase
		block, err := ssh.MarshalPrivateKeyWithPassphrase(priv, commentSanitized, []byte(passphrase))
		if err != nil {
			return authorizedKey, privatePEM, fmt.Errorf("error marshalling private key: %w", err)
		}
		privatePEM = pem.EncodeToMemory(block)
	}
	sshPub, err := ssh.NewPublicKey(ed25519.PublicKey(pub))
	if err != nil {
		return authorizedKey, privatePEM, fmt.Errorf("error creating public key: %w", err)
	}
	authorizedKey = ssh.MarshalAuthorizedKey(sshPub)
	return
}
