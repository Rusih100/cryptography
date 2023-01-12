package ciphers

// Шифр Эль-Гамаля

type ElGamal struct {
	publicKey  *PublicKeyElGamal
	privateKey *PrivateKeyElGamal
}
