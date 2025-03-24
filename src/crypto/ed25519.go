package crypto

import (
    "crypto/ed25519"
    "crypto/rand"
    "io"
)

func GenerateKeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
    publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
    if err != nil {
        return nil, nil, err
    }
    return publicKey, privateKey, nil
}

func SignMessage(privateKey ed25519.PrivateKey, message []byte) ([]byte, error) {
    return ed25519.Sign(privateKey, message), nil
}

func VerifySignature(publicKey ed25519.PublicKey, message, signature []byte) bool {
    return ed25519.Verify(publicKey, message, signature)
}

