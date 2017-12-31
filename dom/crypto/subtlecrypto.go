package crypto


import (
	"github.com/matthewmueller/joy/macro"
)

// SubtleCrypto struct
// js:"SubtleCrypto,omit"
type SubtleCrypto struct {
}

// Decrypt fn
// js:"decrypt"
func (*SubtleCrypto) Decrypt(algorithm interface{}, key *CryptoKey, data []byte) (i interface{}) {
	macro.Rewrite("await $_.decrypt($1, $2, $3)", algorithm, key, data)
	return i
}

// DeriveBits fn
// js:"deriveBits"
func (*SubtleCrypto) DeriveBits(algorithm interface{}, baseKey *CryptoKey, length uint) (i interface{}) {
	macro.Rewrite("await $_.deriveBits($1, $2, $3)", algorithm, baseKey, length)
	return i
}

// DeriveKey fn
// js:"deriveKey"
func (*SubtleCrypto) DeriveKey(algorithm interface{}, baseKey *CryptoKey, derivedKeyType interface{}, extractable bool, keyUsages []string) (i interface{}) {
	macro.Rewrite("await $_.deriveKey($1, $2, $3, $4, $5)", algorithm, baseKey, derivedKeyType, extractable, keyUsages)
	return i
}

// Digest fn
// js:"digest"
func (*SubtleCrypto) Digest(algorithm interface{}, data []byte) (i interface{}) {
	macro.Rewrite("await $_.digest($1, $2)", algorithm, data)
	return i
}

// Encrypt fn
// js:"encrypt"
func (*SubtleCrypto) Encrypt(algorithm interface{}, key *CryptoKey, data []byte) (i interface{}) {
	macro.Rewrite("await $_.encrypt($1, $2, $3)", algorithm, key, data)
	return i
}

// ExportKey fn
// js:"exportKey"
func (*SubtleCrypto) ExportKey(format string, key *CryptoKey) (i interface{}) {
	macro.Rewrite("await $_.exportKey($1, $2)", format, key)
	return i
}

// GenerateKey fn
// js:"generateKey"
func (*SubtleCrypto) GenerateKey(algorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	macro.Rewrite("await $_.generateKey($1, $2, $3)", algorithm, extractable, keyUsages)
	return i
}

// ImportKey fn
// js:"importKey"
func (*SubtleCrypto) ImportKey(format string, keyData []byte, algorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	macro.Rewrite("await $_.importKey($1, $2, $3, $4, $5)", format, keyData, algorithm, extractable, keyUsages)
	return i
}

// Sign fn
// js:"sign"
func (*SubtleCrypto) Sign(algorithm interface{}, key *CryptoKey, data []byte) (i interface{}) {
	macro.Rewrite("await $_.sign($1, $2, $3)", algorithm, key, data)
	return i
}

// UnwrapKey fn
// js:"unwrapKey"
func (*SubtleCrypto) UnwrapKey(format string, wrappedKey []byte, unwrappingKey *CryptoKey, unwrapAlgorithm interface{}, unwrappedKeyAlgorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	macro.Rewrite("await $_.unwrapKey($1, $2, $3, $4, $5, $6, $7)", format, wrappedKey, unwrappingKey, unwrapAlgorithm, unwrappedKeyAlgorithm, extractable, keyUsages)
	return i
}

// Verify fn
// js:"verify"
func (*SubtleCrypto) Verify(algorithm interface{}, key *CryptoKey, signature []byte, data []byte) (i interface{}) {
	macro.Rewrite("await $_.verify($1, $2, $3, $4)", algorithm, key, signature, data)
	return i
}

// WrapKey fn
// js:"wrapKey"
func (*SubtleCrypto) WrapKey(format string, key *CryptoKey, wrappingKey *CryptoKey, wrapAlgorithm interface{}) (i interface{}) {
	macro.Rewrite("await $_.wrapKey($1, $2, $3, $4)", format, key, wrappingKey, wrapAlgorithm)
	return i
}
