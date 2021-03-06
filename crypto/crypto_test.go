/*
 * BSD 3-Clause License
 *
 * Copyright (c) 2015-2018, Virgil Security, Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *  Redistributions of source code must retain the above copyright notice, this
 *   list of conditions and the following disclaimer.
 *
 *  Redistributions in binary form must reproduce the above copyright notice,
 *   this list of conditions and the following disclaimer in the documentation
 *   and/or other materials provided with the distribution.
 *
 *  Neither the name of the copyright holder nor the names of its
 *   contributors may be used to endorse or promote products derived from
 *   this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package crypto_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/VirgilSecurity/virgil-commkit-go/crypto"
	"github.com/VirgilSecurity/virgil-commkit-go/crypto/wrapper/foundation"
)

func TestSignVerify(t *testing.T) {
	vcrypto := &crypto.Crypto{}

	// make random data
	data := make([]byte, 257)
	rand.Read(data)

	signerKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	sign, err := vcrypto.Sign(data, signerKey)
	require.NoError(t, err)

	err = vcrypto.VerifySignature(data, sign, signerKey.PublicKey())
	require.NoError(t, err)
}

func TestEncryptDecrypt(t *testing.T) {
	vcrypto := &crypto.Crypto{}

	// make random data
	data := make([]byte, 257)
	rand.Read(data)

	encryptKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	cipherText, err := vcrypto.Encrypt(data, encryptKey.PublicKey())
	require.NoError(t, err)

	actualData, err := vcrypto.Decrypt(cipherText, encryptKey)
	require.NoError(t, err)
	require.Equal(t, data, actualData)
}

func TestEncryptWithPaddingDecrypt(t *testing.T) {
	vcrypto := &crypto.Crypto{}

	// make random data
	data := make([]byte, 257)
	rand.Read(data)

	encryptKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	cipherText, err := vcrypto.EncryptWithPadding(data, true, encryptKey.PublicKey())
	require.NoError(t, err)

	actualData, err := vcrypto.Decrypt(cipherText, encryptKey)
	require.NoError(t, err)
	require.Equal(t, data, actualData)
}

func TestStreamCipher(t *testing.T) {
	vcrypto := &crypto.Crypto{}
	key, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	plainBuf := make([]byte, 102301)
	rand.Read(plainBuf)

	plain := bytes.NewReader(plainBuf)
	cipheredStream := bytes.NewBuffer(nil)
	err = vcrypto.EncryptStream(plain, cipheredStream, key.PublicKey())
	require.NoError(t, err)

	// decrypt with key
	cipheredInputStream := bytes.NewReader(cipheredStream.Bytes())
	plainOutBuffer := bytes.NewBuffer(nil)
	err = vcrypto.DecryptStream(cipheredInputStream, plainOutBuffer, key)
	require.NoError(t, err, "decrypt with correct key")
	require.Equal(t, plainBuf, plainOutBuffer.Bytes(), "decrypt with correct key: plain & decrypted buffers do not match")

	// decrypt with wrong id must fail
	wrongKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	cipheredInputStream = bytes.NewReader(cipheredStream.Bytes())
	plainOutBuffer = bytes.NewBuffer(nil)

	err = vcrypto.DecryptStream(cipheredInputStream, plainOutBuffer, wrongKey)
	require.Error(t, err, "decrypt with incorrect key")
}

func TestStreamCipherWithPadding(t *testing.T) {
	vcrypto := &crypto.Crypto{}
	key, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	plainBuf := make([]byte, 102301)
	rand.Read(plainBuf)

	plain := bytes.NewReader(plainBuf)
	cipheredStream := bytes.NewBuffer(nil)
	err = vcrypto.EncryptStreamWithPadding(plain, cipheredStream, true, key.PublicKey())
	require.NoError(t, err)

	// decrypt with key
	cipheredInputStream := bytes.NewReader(cipheredStream.Bytes())
	plainOutBuffer := bytes.NewBuffer(nil)
	err = vcrypto.DecryptStream(cipheredInputStream, plainOutBuffer, key)
	require.NoError(t, err, "decrypt with correct key")
	require.Equal(t, plainBuf, plainOutBuffer.Bytes())

	// decrypt with wrong id must fail
	wrongKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	cipheredInputStream = bytes.NewReader(cipheredStream.Bytes())
	plainOutBuffer = bytes.NewBuffer(nil)

	err = vcrypto.DecryptStream(cipheredInputStream, plainOutBuffer, wrongKey)
	require.Error(t, err, "decrypt with incorrect key")
}

func TestStreamSigner(t *testing.T) {
	vcrypto := &crypto.Crypto{}
	key, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	plainBuf := make([]byte, 1023)
	rand.Read(plainBuf)
	plain := bytes.NewBuffer(plainBuf)
	sign, err := vcrypto.SignStream(plain, key)
	require.NoError(t, err)

	// verify signature
	plain = bytes.NewBuffer(plainBuf)
	err = vcrypto.VerifyStream(plain, sign, key.PublicKey())
	require.NoError(t, err)

	// verify with wrong key must fail
	wrongKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	err = vcrypto.VerifyStream(plain, sign, wrongKey.PublicKey())
	require.Error(t, crypto.ErrSignVerification, err)

	// verify with wrong signature must fail
	plain = bytes.NewBuffer(plainBuf)
	sign[len(sign)-1] = ^sign[len(sign)-1] // invert last byte

	err = vcrypto.VerifyStream(plain, sign, wrongKey.PublicKey())
	require.Equal(t, crypto.ErrSignVerification, err)
}

func TestExportImportKeys(t *testing.T) {
	vcrypto := &crypto.Crypto{}
	key, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	pubb, err := vcrypto.ExportPublicKey(key.PublicKey())
	require.NoError(t, err)

	privb, err := vcrypto.ExportPrivateKey(key)
	require.NoError(t, err)

	pub, err := vcrypto.ImportPublicKey(pubb)
	require.NoError(t, err)

	priv, err := vcrypto.ImportPrivateKey(privb)
	require.NoError(t, err)

	data := make([]byte, 257)
	rand.Read(data)

	// check that import keys was correct
	{
		cipherText, err := vcrypto.SignThenEncrypt(data, key, key.PublicKey())
		require.NoError(t, err)

		plaintext, err := vcrypto.DecryptThenVerify(cipherText, priv, pub)
		require.NoError(t, err)
		require.Equal(t, plaintext, data)
	}
}

func TestSignAndEncryptAndDecryptAndVerify(t *testing.T) {
	vcrypto := &crypto.Crypto{}

	signKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	encryptKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	data := make([]byte, 257)
	rand.Read(data)

	cipherText, err := vcrypto.SignAndEncrypt(data, signKey, encryptKey.PublicKey())
	require.NoError(t, err)

	plaintext, err := vcrypto.DecryptAndVerify(cipherText, encryptKey, signKey.PublicKey(), encryptKey.PublicKey())
	require.NoError(t, err)
	require.Equal(t, data, plaintext)
}

func TestSignAndEncryptWithPaddingAndDecryptAndVerify(t *testing.T) {
	vcrypto := &crypto.Crypto{}

	signKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	encryptKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	data := make([]byte, 257)
	rand.Read(data)

	cipherText, err := vcrypto.SignAndEncryptWithPadding(data, signKey, true, encryptKey.PublicKey())
	require.NoError(t, err)

	plaintext, err := vcrypto.DecryptAndVerify(cipherText, encryptKey, signKey.PublicKey(), encryptKey.PublicKey())
	require.NoError(t, err)
	require.Equal(t, data, plaintext)
}

func TestSignThenEncryptAndDecryptThenVerify(t *testing.T) {
	vcrypto := &crypto.Crypto{}

	signKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	encryptKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	data := make([]byte, 257)
	rand.Read(data)

	cipherText, err := vcrypto.SignThenEncrypt(data, signKey, encryptKey.PublicKey())
	require.NoError(t, err)

	plaintext, err := vcrypto.DecryptThenVerify(cipherText, encryptKey, signKey.PublicKey(), encryptKey.PublicKey())
	require.NoError(t, err)
	require.Equal(t, data, plaintext)
}

func TestSignThenEncryptWithPaddingAndDecryptThenVerify(t *testing.T) {
	vcrypto := &crypto.Crypto{}

	signKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	encryptKey, err := vcrypto.GenerateKeypair()
	require.NoError(t, err)

	data := make([]byte, 257)
	rand.Read(data)

	cipherText, err := vcrypto.SignThenEncryptWithPadding(data, signKey, true, encryptKey.PublicKey())
	require.NoError(t, err)

	plaintext, err := vcrypto.DecryptThenVerify(cipherText, encryptKey, signKey.PublicKey(), encryptKey.PublicKey())
	require.NoError(t, err)
	require.Equal(t, data, plaintext)
}

func TestGenerateKeypairFromKeyMaterial(t *testing.T) {
	seed := make([]byte, 384)
	for i := range seed {
		seed[i] = byte(i)
	}

	pub1, priv1 := GenKeysFromSeed(t, seed)

	for i := 0; i < 10; i++ {
		pub2, priv2 := GenKeysFromSeed(t, seed)
		require.Equal(t, pub1, pub2)
		require.Equal(t, priv1, priv2)
	}

	// check if we change seed than key pair is different
	{
		seed[383]++
		pub3, priv3 := GenKeysFromSeed(t, seed)
		require.NotEqual(t, pub1, pub3)
		require.NotEqual(t, priv1, priv3)
	}
}

func GenKeysFromSeed(t *testing.T, seed []byte) (publicKey []byte, privateKey []byte) {
	vcrypto := &crypto.Crypto{}
	key, err := vcrypto.GenerateKeypairFromKeyMaterial(seed)
	require.NoError(t, err)

	publicKey, err = vcrypto.ExportPublicKey(key.PublicKey())
	require.NoError(t, err)

	privateKey, err = vcrypto.ExportPrivateKey(key)
	require.NoError(t, err)

	return publicKey, privateKey
}

func TestGenerateKeypairFromKeyMaterialBadCase(t *testing.T) {
	table := []struct {
		name string
		size int
	}{
		{"less 32", 31},
		{"greater 512", 513},
	}
	vcrypto := &crypto.Crypto{}

	for _, test := range table {
		data, err := vcrypto.Random(test.size)
		require.NoError(t, err)

		_, err = vcrypto.GenerateKeypairFromKeyMaterial(data)
		require.Equal(t, crypto.ErrInvalidSeedSize, err, test.name)
	}
}

func TestKeyTypes(t *testing.T) {
	vcrypto := &crypto.Crypto{}
	m, err := vcrypto.Random(128)
	require.NoError(t, err)

	table := []struct {
		kt            crypto.KeyType
		expectedError error
	}{
		{crypto.DefaultKeyType, nil},
		{crypto.Rsa2048, nil},
		{crypto.P256r1, nil},
		{crypto.Curve25519, nil},
		{crypto.Ed25519, nil},
		{crypto.Curve25519Ed25519, nil},
		{crypto.Curve25519Round5Ed25519Falcon, nil},
		{crypto.KeyType(100), crypto.ErrUnsupportedKeyType},
	}

	fs := []func(kt crypto.KeyType) error{
		func(kt crypto.KeyType) error {
			vcrypto.KeyType = kt
			_, err := vcrypto.GenerateKeypair()
			return err
		},
		func(kt crypto.KeyType) error {
			_, err := vcrypto.GenerateKeypairFromKeyMaterialForType(kt, m)
			return err
		},
	}
	for _, test := range table {
		for _, f := range fs {
			err := f(test.kt)
			require.Equal(t, test.expectedError, err, test.kt)
		}
	}
}

func TestImport(t *testing.T) {
	pubKey := []byte{0x30, 0x82, 0x0B, 0x0E, 0x30, 0x51, 0x06, 0x0A, 0x2B, 0x06, 0x01, 0x04, 0x01, 0x83, 0xAC, 0x1B,
		0x01, 0x01, 0x30, 0x43, 0x30, 0x24, 0x06, 0x0A, 0x2B, 0x06, 0x01, 0x04, 0x01, 0x83, 0xAC, 0x1B,
		0x01, 0x02, 0x30, 0x16, 0x30, 0x05, 0x06, 0x03, 0x2B, 0x65, 0x6E, 0x30, 0x0D, 0x06, 0x0B, 0x2B,
		0x06, 0x01, 0x04, 0x01, 0x83, 0xAC, 0x1B, 0x02, 0x02, 0x0B, 0x30, 0x0C, 0x06, 0x0A, 0x2B, 0x06,
		0x01, 0x04, 0x01, 0x83, 0xAC, 0x1B, 0x02, 0x01, 0x30, 0x0D, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01,
		0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x03, 0x82, 0x0A, 0xB7, 0x00, 0x30, 0x82, 0x0A, 0xB2,
		0x04, 0x82, 0x03, 0xFC, 0x30, 0x82, 0x03, 0xF8, 0x04, 0x20, 0x48, 0x03, 0x46, 0x4A, 0x60, 0xB8,
		0x8D, 0x79, 0x7E, 0xED, 0xE6, 0xF3, 0xF3, 0x4A, 0xFB, 0xCC, 0xE9, 0x24, 0x8D, 0xC6, 0x09, 0x0E,
		0xF3, 0x39, 0xEC, 0x09, 0xA8, 0x34, 0xF3, 0xBD, 0x21, 0x5A, 0x04, 0x82, 0x03, 0xD2, 0x0A, 0x42,
		0x45, 0x43, 0x6A, 0xA0, 0xD3, 0x82, 0x28, 0x42, 0xEA, 0xD2, 0xFA, 0x27, 0xFD, 0xDF, 0xD3, 0x47,
		0x1D, 0x32, 0x34, 0xA8, 0x79, 0x3B, 0x4E, 0xC8, 0x75, 0xD1, 0xF6, 0xF7, 0xFF, 0x75, 0xB6, 0x67,
		0x57, 0xC0, 0xEC, 0xAC, 0xF4, 0xBB, 0x9B, 0x1A, 0xA3, 0x93, 0x8C, 0xA5, 0x7E, 0x4D, 0x5F, 0x01,
		0xA8, 0x69, 0x9A, 0x86, 0xD1, 0x86, 0xB8, 0xA9, 0x81, 0x0E, 0x69, 0x75, 0xD3, 0xD8, 0x4E, 0x11,
		0x8E, 0xDE, 0x3F, 0x37, 0x83, 0x66, 0x91, 0xDE, 0xC6, 0x21, 0xF2, 0xEA, 0x50, 0xDE, 0x6B, 0xDF,
		0x4D, 0xCB, 0xB5, 0xC4, 0xD4, 0x65, 0xAC, 0xA3, 0xCD, 0x03, 0xDE, 0xFE, 0xFB, 0x92, 0x44, 0x37,
		0xF3, 0x01, 0xFD, 0x2E, 0xE8, 0x88, 0xAD, 0x92, 0x89, 0xE5, 0xAB, 0x82, 0x69, 0x69, 0x2F, 0xAA,
		0x53, 0xB0, 0x64, 0x3D, 0x89, 0x8B, 0x0A, 0xF5, 0x68, 0x4F, 0x3B, 0xBE, 0xB2, 0x93, 0x81, 0xEA,
		0xFA, 0x02, 0x3E, 0x16, 0xA1, 0xD2, 0x8E, 0x88, 0x97, 0xA2, 0x47, 0x6A, 0xCF, 0xDB, 0x19, 0xE6,
		0xF0, 0xBE, 0xCB, 0x85, 0x74, 0xF8, 0xC0, 0x15, 0x7A, 0x42, 0x56, 0x37, 0xF6, 0x28, 0x9B, 0xCB,
		0x1B, 0xAA, 0x2E, 0xCC, 0xE8, 0x34, 0xBE, 0xA5, 0x3E, 0xC9, 0xC4, 0x27, 0xDA, 0x4D, 0xCF, 0x79,
		0x10, 0xDB, 0xD5, 0x16, 0x96, 0x46, 0xF8, 0x26, 0x1A, 0x69, 0x25, 0x26, 0x12, 0x1E, 0x27, 0xFC,
		0x63, 0xF9, 0x0C, 0x1F, 0x2C, 0xA0, 0x20, 0x54, 0xBB, 0x3B, 0x08, 0x7C, 0x82, 0xD8, 0xE4, 0x96,
		0x83, 0x6B, 0x08, 0xBD, 0x32, 0x4B, 0xF2, 0x45, 0x33, 0xC6, 0xAC, 0x37, 0xF5, 0xEA, 0x5E, 0x17,
		0x13, 0x4F, 0x66, 0x94, 0xD7, 0x7B, 0x46, 0x87, 0x41, 0xB6, 0xB2, 0xAD, 0x31, 0x87, 0xD4, 0x3B,
		0x96, 0x8B, 0xB7, 0x50, 0x28, 0x0C, 0x08, 0x80, 0x00, 0xA9, 0x22, 0x7A, 0xD8, 0xED, 0x4A, 0x23,
		0x79, 0xCA, 0xD3, 0xAD, 0xBC, 0x25, 0x19, 0x30, 0x8E, 0xE6, 0xE9, 0x8A, 0xB7, 0x20, 0x86, 0x95,
		0xDA, 0xDC, 0x66, 0x07, 0xF9, 0x41, 0x98, 0x4B, 0xA9, 0x31, 0x68, 0xC4, 0x63, 0x1A, 0x4B, 0x2C,
		0x1C, 0x05, 0x67, 0x26, 0x63, 0x3C, 0xCA, 0x5D, 0x50, 0x87, 0xE4, 0xC1, 0x61, 0xDB, 0x70, 0x8F,
		0x9A, 0x40, 0xF1, 0xB0, 0x38, 0x0B, 0x51, 0x41, 0xC7, 0x80, 0x78, 0x21, 0x10, 0x47, 0x5D, 0x5D,
		0x11, 0x1F, 0x50, 0x19, 0x95, 0x90, 0x34, 0xF2, 0x29, 0x5C, 0x1E, 0xEE, 0x10, 0xCE, 0x6A, 0xB0,
		0x91, 0xA6, 0x21, 0xED, 0xDC, 0x3D, 0xAC, 0xDE, 0xFF, 0x95, 0x8A, 0xFA, 0xCD, 0x2A, 0xC1, 0x58,
		0x1F, 0x91, 0x7A, 0x69, 0xCC, 0x07, 0x51, 0x0F, 0x73, 0xD4, 0x00, 0xB3, 0xED, 0xAF, 0x32, 0x2C,
		0x6D, 0x61, 0x34, 0x2D, 0x5E, 0x75, 0x31, 0xC5, 0xD4, 0x88, 0x6A, 0x56, 0xBA, 0x97, 0x4D, 0xF6,
		0xC3, 0xEA, 0xC0, 0x32, 0xEF, 0xB8, 0xAD, 0x6B, 0x9A, 0x37, 0x90, 0xBE, 0x24, 0x43, 0xD8, 0x86,
		0xC6, 0x89, 0xAF, 0xD5, 0xC6, 0x1B, 0x3C, 0xEC, 0xC0, 0x6B, 0x0B, 0x3E, 0xBF, 0x6B, 0x22, 0xF5,
		0xE8, 0x9F, 0xBD, 0x73, 0x78, 0xF7, 0xD0, 0x04, 0x5C, 0xDD, 0xA2, 0x59, 0xF7, 0x2B, 0xB6, 0x5E,
		0x26, 0xB2, 0x60, 0x13, 0x96, 0x64, 0x90, 0xE0, 0x8E, 0x5A, 0x2E, 0xC2, 0xD6, 0xB2, 0xC1, 0xF8,
		0x38, 0x8F, 0xFA, 0x3D, 0x56, 0x4A, 0x95, 0xC8, 0xD0, 0xB0, 0x47, 0x7F, 0x7D, 0x9B, 0xD6, 0x29,
		0xF2, 0xDE, 0xB8, 0xFD, 0x29, 0xD2, 0x5B, 0x2F, 0xEF, 0x0D, 0x74, 0x30, 0xC8, 0x11, 0x9B, 0x2C,
		0x75, 0x71, 0xD0, 0x5D, 0x23, 0xB3, 0x33, 0x74, 0x17, 0xC9, 0xDC, 0x97, 0x6F, 0x74, 0xDA, 0x1C,
		0x5D, 0x7D, 0xA8, 0xC7, 0x1C, 0xA5, 0x84, 0x34, 0x79, 0x13, 0x70, 0x6D, 0x53, 0x31, 0x6F, 0x72,
		0x94, 0x40, 0xD0, 0x88, 0x00, 0xB2, 0xDC, 0x24, 0xF5, 0xE0, 0xB8, 0xC7, 0x0C, 0x6D, 0x04, 0xD8,
		0x3A, 0x26, 0x19, 0xDD, 0xAA, 0x3A, 0x61, 0xBB, 0x46, 0x87, 0x0E, 0x8C, 0x2B, 0x17, 0xE0, 0x35,
		0xFE, 0x94, 0xF8, 0x20, 0xD7, 0xE2, 0xA4, 0x54, 0x2F, 0xC6, 0x8A, 0xCD, 0xBB, 0xD4, 0x21, 0x71,
		0x48, 0xC4, 0x39, 0xC8, 0xA5, 0xA2, 0x8A, 0x1A, 0x07, 0x8F, 0xDB, 0xF3, 0x92, 0x4B, 0x7F, 0x60,
		0x4C, 0xBA, 0x6F, 0x19, 0x69, 0x7E, 0x14, 0x41, 0x95, 0x24, 0x13, 0x3C, 0x4D, 0x54, 0x0B, 0x61,
		0x8C, 0x27, 0x34, 0x44, 0x12, 0x29, 0x8B, 0xB4, 0x6C, 0x99, 0xFE, 0x27, 0xEF, 0xF3, 0x4D, 0x2C,
		0x42, 0x10, 0xBF, 0x75, 0xF7, 0x73, 0xF1, 0xF1, 0xAB, 0x51, 0xD3, 0x74, 0xA3, 0x6C, 0x13, 0x1D,
		0x1F, 0xC3, 0xFE, 0xCA, 0x54, 0xE4, 0x58, 0x77, 0x8A, 0xFB, 0x92, 0xA5, 0x20, 0x3E, 0xD0, 0x40,
		0xC7, 0x60, 0x5F, 0xA2, 0x3B, 0xF4, 0x5A, 0x1A, 0xAC, 0x09, 0xD9, 0xE6, 0xD0, 0x43, 0x82, 0x2F,
		0xCE, 0xC5, 0x89, 0x3B, 0x9B, 0xD0, 0xA1, 0x6D, 0xA6, 0x5D, 0xB4, 0x41, 0x6A, 0xB1, 0xD7, 0x80,
		0x31, 0x94, 0x4C, 0x12, 0x8A, 0xC0, 0xB4, 0x33, 0x0E, 0x70, 0x38, 0x25, 0xD3, 0x71, 0x34, 0xBF,
		0xD1, 0x52, 0xFC, 0x07, 0x7E, 0x0D, 0xBC, 0x3B, 0x30, 0xCA, 0x1B, 0x06, 0xAA, 0xB2, 0xAD, 0x75,
		0x23, 0xF6, 0x34, 0x6C, 0x53, 0xB8, 0x9D, 0xD2, 0xA4, 0xDB, 0xBD, 0x77, 0x88, 0x9D, 0x61, 0x8A,
		0xBE, 0xAB, 0x6E, 0x83, 0x89, 0xDD, 0x23, 0x6C, 0x54, 0xE4, 0xDB, 0x95, 0x4F, 0x7E, 0xBB, 0x9D,
		0x61, 0x11, 0xBE, 0xE9, 0xDA, 0x23, 0x58, 0xA4, 0xA6, 0xFD, 0x17, 0x99, 0x19, 0x1E, 0xB2, 0xC4,
		0x5A, 0xBA, 0xF7, 0x77, 0x40, 0x7B, 0x9E, 0x92, 0xE9, 0x72, 0x1C, 0xD4, 0xC3, 0xAE, 0x2E, 0x48,
		0x62, 0x11, 0x96, 0xC1, 0xF9, 0xB9, 0x33, 0x2D, 0x30, 0x1E, 0xD4, 0xED, 0xCB, 0x0D, 0xE7, 0x77,
		0x1C, 0x1F, 0x48, 0x3E, 0xE2, 0x89, 0xFF, 0xC0, 0x6D, 0x8B, 0x41, 0x10, 0xC8, 0xDA, 0x86, 0x29,
		0xFF, 0xE2, 0xB4, 0xF3, 0x7A, 0xB2, 0x73, 0xB6, 0x59, 0x41, 0x45, 0x1A, 0x24, 0xF5, 0x43, 0xED,
		0x90, 0xDC, 0x1F, 0xBC, 0x8B, 0xCD, 0x5D, 0xF8, 0x3A, 0xD5, 0x68, 0x4F, 0x30, 0xA8, 0x5B, 0x03,
		0x69, 0x53, 0x9A, 0x4C, 0xA3, 0x9F, 0xF7, 0x30, 0x57, 0xA0, 0x56, 0xBC, 0x30, 0xDB, 0xA4, 0xF1,
		0x0D, 0x29, 0x5D, 0x53, 0xC4, 0xD7, 0x2B, 0xB5, 0x41, 0x67, 0x3B, 0xC1, 0x88, 0x06, 0x4B, 0x3E,
		0xC7, 0xB7, 0xA3, 0x28, 0x6C, 0x7C, 0x7E, 0x9A, 0x37, 0x37, 0x6A, 0xD6, 0xF7, 0x5B, 0xDB, 0xBB,
		0xF3, 0x4A, 0xEB, 0x1E, 0x4E, 0x00, 0xF4, 0x66, 0xAD, 0xD4, 0x88, 0xDA, 0x82, 0x49, 0x9E, 0x85,
		0x44, 0x1C, 0xFC, 0x40, 0xD9, 0x08, 0x27, 0xF3, 0xA0, 0xD5, 0x40, 0x9D, 0x10, 0xA7, 0xB0, 0xB0,
		0xFF, 0xFB, 0x45, 0xD0, 0x2B, 0xBD, 0x5B, 0xCE, 0x3B, 0xAF, 0xB8, 0x4B, 0xC1, 0x3D, 0xA2, 0xA7,
		0x5B, 0x36, 0x46, 0xAF, 0x24, 0x16, 0x1A, 0xC6, 0x0B, 0x71, 0x06, 0x97, 0xCF, 0x66, 0xF3, 0xFC,
		0xAC, 0x6F, 0xE8, 0xA4, 0xA0, 0x55, 0x29, 0x9F, 0x47, 0xEC, 0x09, 0xB1, 0x5D, 0x0B, 0xE7, 0x91,
		0x55, 0x92, 0xD1, 0xB8, 0xA7, 0x01, 0xB7, 0xA2, 0x81, 0xE2, 0x66, 0x98, 0xB7, 0xF1, 0xA0, 0xFD,
		0x04, 0x82, 0x03, 0x81, 0x09, 0x99, 0x28, 0x89, 0xAA, 0x74, 0xA5, 0x80, 0xAE, 0x60, 0x3C, 0x12,
		0x22, 0x54, 0x66, 0x87, 0x25, 0x9B, 0xE4, 0x0A, 0xA2, 0xFE, 0x73, 0x8D, 0x8F, 0x75, 0xAF, 0x83,
		0x8E, 0x67, 0x76, 0xF6, 0xF1, 0xFA, 0x5D, 0x20, 0x29, 0x5C, 0x38, 0x58, 0xF5, 0x58, 0x5E, 0x51,
		0x4D, 0xCB, 0xCB, 0xD2, 0xDF, 0x9D, 0x18, 0x6E, 0x0E, 0x5A, 0xCD, 0x41, 0x0A, 0x73, 0x96, 0xE6,
		0x56, 0x31, 0x5E, 0x6E, 0x73, 0x9C, 0x64, 0xCB, 0x3D, 0xA4, 0x9F, 0x65, 0x0E, 0x2B, 0x98, 0x6D,
		0x89, 0x00, 0x23, 0x66, 0x78, 0x46, 0xF0, 0x01, 0x25, 0x78, 0xA4, 0x68, 0xA0, 0x62, 0xE5, 0xBD,
		0x21, 0x28, 0x13, 0x6B, 0xC1, 0x5E, 0x3E, 0x53, 0xCA, 0x80, 0x15, 0x12, 0x5F, 0x42, 0x8E, 0xDE,
		0x8F, 0xC2, 0xE8, 0x49, 0xA9, 0x31, 0xEC, 0xD5, 0x80, 0x4E, 0xE4, 0x29, 0x1A, 0x0D, 0x8C, 0x77,
		0x87, 0xA7, 0x46, 0x13, 0x68, 0x6C, 0xB0, 0xD7, 0x1C, 0x67, 0x4C, 0x7C, 0x45, 0x26, 0xD5, 0x55,
		0xA1, 0x86, 0x56, 0xE5, 0x2A, 0x1B, 0xE9, 0x27, 0x44, 0x2A, 0x56, 0x5A, 0xD9, 0x1C, 0x65, 0x25,
		0x01, 0xA7, 0x76, 0x21, 0xC2, 0x50, 0x3B, 0xCC, 0xBD, 0xD4, 0x4D, 0x4A, 0x63, 0xB7, 0x2E, 0xA1,
		0xCB, 0x9A, 0x55, 0x55, 0x01, 0xAA, 0x44, 0x91, 0xDC, 0xC0, 0xBD, 0x1F, 0x78, 0xE8, 0xC9, 0xC6,
		0x52, 0xBE, 0x03, 0x91, 0x50, 0x90, 0xC6, 0xAE, 0xC1, 0x6E, 0xD9, 0x6E, 0x30, 0xAC, 0xDB, 0x96,
		0x14, 0x31, 0x7F, 0x8A, 0xF5, 0x07, 0xB4, 0x6A, 0xE0, 0xA4, 0x86, 0x0D, 0x5B, 0xA9, 0x0A, 0xA5,
		0x65, 0x50, 0x8D, 0x5A, 0xDD, 0x24, 0xD2, 0x33, 0xAA, 0x23, 0x5B, 0xBD, 0x05, 0xC4, 0xD3, 0xF2,
		0x61, 0x15, 0x3A, 0xA3, 0x42, 0xC4, 0x60, 0x98, 0xED, 0x1D, 0x8A, 0x76, 0x65, 0x4A, 0x48, 0xA4,
		0xC2, 0x1D, 0x4A, 0x4A, 0x08, 0xB5, 0xA5, 0xF9, 0x5E, 0x01, 0x79, 0xF3, 0x2B, 0x13, 0x93, 0x49,
		0x05, 0xBB, 0x64, 0x57, 0x8E, 0x39, 0x25, 0x46, 0xC8, 0xE6, 0x3F, 0xDA, 0x14, 0x04, 0xC4, 0x6C,
		0x96, 0x5D, 0xD3, 0x2E, 0x03, 0x30, 0xE9, 0x19, 0xC0, 0xC2, 0x73, 0x53, 0x2E, 0x3E, 0xCB, 0x98,
		0x4A, 0x1C, 0x18, 0x89, 0x8F, 0x30, 0xAD, 0x2A, 0xED, 0x9A, 0x32, 0x9F, 0x90, 0x35, 0xD4, 0x79,
		0xBE, 0xBC, 0x67, 0x0B, 0x19, 0x87, 0xC9, 0x7D, 0x99, 0x51, 0xDB, 0xEA, 0xAC, 0x35, 0x8E, 0x34,
		0x5C, 0xA3, 0xE0, 0x91, 0x81, 0xBA, 0x66, 0x0C, 0xFB, 0x3F, 0xD9, 0xD9, 0x9D, 0x00, 0xB0, 0x26,
		0x98, 0x0F, 0xB3, 0x3B, 0x42, 0x27, 0x84, 0xC5, 0x07, 0x02, 0x6D, 0xBE, 0xFB, 0xC8, 0x5A, 0x84,
		0x4C, 0x9A, 0xAE, 0x3C, 0x9B, 0xA6, 0xDC, 0x32, 0x8D, 0x49, 0x11, 0x30, 0x07, 0xE0, 0xFF, 0x2E,
		0x06, 0xBE, 0x85, 0x01, 0xA3, 0xB2, 0xB6, 0x30, 0x2A, 0x10, 0x06, 0x8A, 0x8C, 0x05, 0x6A, 0x44,
		0x44, 0xA0, 0x94, 0xB2, 0xA3, 0x86, 0xAD, 0x05, 0x8B, 0x57, 0x4E, 0x61, 0x1C, 0xBC, 0xBA, 0x90,
		0x5C, 0x85, 0x05, 0x02, 0xA8, 0x91, 0xE2, 0x2A, 0x8C, 0x2A, 0x61, 0x9B, 0x50, 0xA4, 0x58, 0x56,
		0x3B, 0x26, 0x11, 0x50, 0xB2, 0x5C, 0xE4, 0x95, 0x84, 0xA9, 0x99, 0x93, 0x0B, 0xA5, 0x76, 0x40,
		0x89, 0x67, 0xF4, 0x0D, 0x3F, 0x6F, 0x4E, 0x69, 0xF2, 0xF5, 0xEB, 0xE7, 0x60, 0xAC, 0x33, 0xF0,
		0x89, 0xC3, 0x8C, 0x16, 0x50, 0x38, 0x93, 0xD6, 0x8A, 0x09, 0x39, 0x95, 0x92, 0xDB, 0x8C, 0x6A,
		0x3C, 0x5D, 0x99, 0x0C, 0x74, 0x02, 0xAE, 0xD5, 0x40, 0xCD, 0x59, 0x11, 0xD2, 0x83, 0x2D, 0x6B,
		0x46, 0x2C, 0xB2, 0x5E, 0xEB, 0x3D, 0x29, 0x86, 0xA7, 0xA9, 0xD7, 0xCC, 0xC1, 0xB5, 0x90, 0xEC,
		0x06, 0xE7, 0xA3, 0x7E, 0x08, 0x55, 0x1B, 0x91, 0x2A, 0x5B, 0x02, 0x9D, 0xCE, 0x1B, 0xD1, 0x3A,
		0x5B, 0xC9, 0x8A, 0xC1, 0xF7, 0x24, 0x2F, 0x4F, 0xA8, 0x0D, 0xEA, 0x3C, 0xA6, 0xA0, 0x13, 0xE0,
		0x16, 0x1D, 0x26, 0xD2, 0x55, 0x58, 0xA2, 0x05, 0xFD, 0x23, 0x63, 0xD5, 0x0C, 0xE2, 0xAF, 0x96,
		0x7B, 0xBA, 0xC3, 0x25, 0x58, 0x7A, 0xEA, 0x58, 0x25, 0x42, 0x50, 0x34, 0x66, 0x64, 0xAA, 0xC0,
		0xEB, 0x69, 0x92, 0x15, 0xD6, 0x68, 0x97, 0xAB, 0x09, 0x81, 0x3A, 0xD8, 0x2A, 0xF3, 0xEF, 0x21,
		0x86, 0x32, 0xE9, 0x69, 0x2A, 0x6B, 0x86, 0xDB, 0x40, 0x29, 0xCE, 0x77, 0x4F, 0xCA, 0x3C, 0xB0,
		0xA4, 0x0F, 0x71, 0x51, 0x87, 0x80, 0x98, 0xA1, 0x67, 0x31, 0x6E, 0x29, 0xBE, 0x77, 0x32, 0x6D,
		0x53, 0x9A, 0x88, 0x56, 0x10, 0x24, 0xC1, 0x47, 0x9D, 0x12, 0x76, 0x21, 0xD5, 0x76, 0xC1, 0x1C,
		0xAA, 0x94, 0x83, 0xF6, 0xDB, 0x4B, 0x33, 0x9B, 0x31, 0x88, 0xD5, 0x59, 0xE8, 0xD2, 0x95, 0xE0,
		0x0B, 0x44, 0xC3, 0xC6, 0x99, 0x6E, 0xDD, 0xB9, 0x3A, 0x7D, 0xB5, 0xE1, 0x1A, 0x89, 0x52, 0x32,
		0x09, 0x83, 0xE7, 0x6D, 0xA3, 0xBF, 0xC2, 0xA2, 0xB8, 0x53, 0x25, 0x7E, 0x59, 0xC6, 0x96, 0x99,
		0x4D, 0x91, 0xCA, 0xB0, 0x7A, 0xC6, 0x03, 0xC0, 0x49, 0xC3, 0x96, 0x9C, 0xBC, 0x79, 0xCF, 0x99,
		0x8C, 0x51, 0x5D, 0xAB, 0xD3, 0xA5, 0xAC, 0x44, 0x42, 0xDE, 0xC9, 0xC8, 0x61, 0xA8, 0x3B, 0x3C,
		0xDD, 0x20, 0x19, 0xEE, 0x04, 0x58, 0x1F, 0xEE, 0x62, 0x09, 0xE7, 0x20, 0xD6, 0xB4, 0xB8, 0xB9,
		0xC1, 0x7B, 0x4D, 0xCA, 0x4B, 0x3D, 0x40, 0x69, 0x38, 0x88, 0x45, 0x23, 0x36, 0x68, 0x47, 0xDE,
		0xEB, 0xC9, 0xA6, 0x99, 0xBD, 0x45, 0x29, 0x00, 0x7D, 0x93, 0x59, 0x1F, 0xA8, 0xE5, 0x2E, 0x53,
		0xA4, 0xED, 0x2E, 0xD5, 0x2E, 0xC8, 0x56, 0xB9, 0x6D, 0x20, 0x88, 0x1E, 0x6A, 0x1B, 0x4A, 0xBE,
		0x5E, 0x00, 0x39, 0x05, 0x60, 0x3A, 0x66, 0x93, 0x6B, 0x44, 0xA2, 0xB2, 0x02, 0x66, 0x14, 0xCB,
		0xF8, 0xDF, 0x4A, 0x3B, 0xF5, 0x3B, 0xC1, 0xDF, 0x2A, 0x60, 0x2D, 0x8C, 0x2F, 0xA4, 0x19, 0x19,
		0xEE, 0xAC, 0x71, 0xFC, 0xAA, 0x8B, 0x40, 0xD8, 0x40, 0x39, 0x55, 0xC2, 0xE0, 0x51, 0xB7, 0x85,
		0xA8, 0x3D, 0x98, 0x3A, 0x01, 0xAF, 0x39, 0x44, 0xCC, 0x41, 0x0F, 0xCD, 0x4B, 0x99, 0x12, 0xC6,
		0xD5, 0x86, 0x97, 0x48, 0x7E, 0x1C, 0xF1, 0x77, 0xEB, 0x5A, 0xF3, 0x23, 0x95, 0x87, 0xD2, 0xFB,
		0xE0, 0x49, 0xA6, 0xBC, 0x3C, 0x1B, 0x58, 0x85, 0x7E, 0x57, 0x7A, 0x3C, 0xCB, 0xA1, 0x86, 0xAD,
		0x92, 0x7A, 0x4B, 0x00, 0xC9, 0x63, 0xD8, 0x57, 0xF0, 0x84, 0x00, 0xCA, 0x0A, 0x1B, 0x28, 0xCE,
		0xD5, 0x42, 0x60, 0xEA, 0xA6, 0x04, 0x82, 0x03, 0x29, 0x59, 0xBB, 0x7B, 0x0C, 0x24, 0x3A, 0xD6,
		0x1B, 0x60, 0x05, 0x55, 0x8F, 0x9F, 0x15, 0xCF, 0x1A, 0x5D, 0x58, 0xD5, 0xB8, 0x7B, 0x37, 0xDC,
		0x2A, 0x44, 0x78, 0x42, 0x31, 0xA5, 0xDF, 0xE7, 0xD4, 0xAB, 0x1D, 0x13, 0x38, 0xC6, 0xB7, 0x03,
		0x46, 0x7A, 0x10, 0xCF, 0xD0, 0x03, 0xD0, 0x7D, 0x06, 0xE1, 0x63, 0x06, 0x60, 0x43, 0x08, 0x91,
		0x41, 0x06, 0x00, 0x4E, 0x0C, 0x40, 0x2C, 0xFD, 0x80, 0xA1, 0x0E, 0xE0, 0x03, 0x02, 0x41, 0x40,
		0x01, 0x6F, 0x6A, 0x05, 0x70, 0x30, 0xF4, 0xC0, 0x10, 0x16, 0xB0, 0x80, 0xF6, 0xE0, 0x5F, 0x05,
		0xE0, 0x07, 0x11, 0xCF, 0x58, 0xEB, 0xE0, 0xB1, 0xF1, 0xFF, 0xBA, 0xFF, 0xB0, 0x45, 0x01, 0x0F,
		0x6A, 0xE1, 0x70, 0x5C, 0xF7, 0x4F, 0xF9, 0x0B, 0xD0, 0x81, 0x11, 0xEF, 0x42, 0xFA, 0xDF, 0x84,
		0xF5, 0x40, 0x2C, 0xFE, 0x8F, 0x74, 0xFD, 0x7F, 0x2B, 0xF3, 0x3F, 0x66, 0xF2, 0x50, 0x12, 0xF3,
		0xCF, 0xF7, 0x00, 0x60, 0x91, 0x00, 0xE1, 0x28, 0xF5, 0xCF, 0x23, 0x00, 0x8F, 0xB8, 0x07, 0x80,
		0x60, 0x00, 0x9F, 0xD6, 0xFE, 0xEF, 0x9C, 0x04, 0x4F, 0xA5, 0x0F, 0x8F, 0x9B, 0x0D, 0x80, 0x9F,
		0x0F, 0xBF, 0xD8, 0xFD, 0x91, 0x3F, 0x02, 0x40, 0x34, 0xF7, 0x81, 0x68, 0xFC, 0xE0, 0x7F, 0xF4,
		0x5E, 0xD1, 0xFC, 0x0F, 0xB2, 0xF6, 0x50, 0x32, 0xF9, 0xEE, 0x97, 0xF4, 0xBE, 0xE7, 0x0E, 0x2F,
		0x41, 0xF8, 0xFF, 0x7A, 0x04, 0x1F, 0x72, 0xFF, 0xA0, 0x44, 0x0C, 0xE0, 0x50, 0x0A, 0x50, 0x64,
		0x0B, 0x70, 0x0A, 0xED, 0xEF, 0xDD, 0x0A, 0x5F, 0x95, 0x14, 0x20, 0x22, 0xF6, 0xAF, 0x9A, 0xF0,
		0xBF, 0x32, 0xFB, 0x6F, 0x2F, 0x08, 0x4F, 0xDA, 0x0B, 0x5F, 0x9D, 0xFF, 0x60, 0x60, 0x12, 0x71,
		0x46, 0xF7, 0x1F, 0xB8, 0x07, 0xD0, 0xDA, 0x01, 0xF0, 0x81, 0xFB, 0x50, 0x4D, 0xFC, 0x71, 0x16,
		0xF9, 0xE0, 0x78, 0x05, 0xFF, 0x73, 0x06, 0xDF, 0xA9, 0xF5, 0x60, 0xB4, 0x16, 0x71, 0x06, 0x0B,
		0xFF, 0x84, 0x12, 0x30, 0xEE, 0x08, 0x6E, 0xCE, 0xFA, 0x10, 0xF7, 0xEC, 0xEF, 0xE4, 0x00, 0x00,
		0x46, 0x07, 0xB0, 0xE9, 0x02, 0xBF, 0xFF, 0x0A, 0xCF, 0xF9, 0x10, 0x90, 0x5C, 0x08, 0xC0, 0xE6,
		0xE7, 0xBF, 0x9C, 0xFD, 0xBF, 0x94, 0x01, 0xF0, 0x55, 0xFF, 0xA0, 0x14, 0xFB, 0x10, 0x3F, 0x0C,
		0xC0, 0x88, 0x03, 0xEF, 0xBE, 0xF0, 0xD0, 0x47, 0xE4, 0xAF, 0x3B, 0xFC, 0x5F, 0x03, 0x10, 0xE0,
		0x08, 0x10, 0x0E, 0x89, 0x0A, 0x00, 0x54, 0xEE, 0x20, 0x54, 0xF4, 0x81, 0x0B, 0x01, 0xD0, 0xAA,
		0x04, 0xF1, 0x6A, 0xE7, 0xEF, 0x87, 0xF2, 0x00, 0x29, 0x00, 0x6E, 0x97, 0xF8, 0x60, 0x18, 0xFF,
		0xE0, 0x30, 0xFB, 0x4F, 0x5E, 0xF9, 0x3F, 0xE4, 0xF1, 0xDF, 0x49, 0x01, 0xD0, 0x5F, 0xFC, 0x9F,
		0x68, 0x0C, 0xEE, 0xD3, 0xF8, 0x2F, 0x9B, 0xF8, 0xA0, 0xA7, 0x00, 0x01, 0x7D, 0x06, 0x4F, 0xA5,
		0xFC, 0x6F, 0xD7, 0x0B, 0x30, 0x15, 0xEA, 0x0F, 0x8E, 0xF6, 0x3F, 0x8D, 0x00, 0xF1, 0x6D, 0xFB,
		0xE0, 0x92, 0xF5, 0xF0, 0xA0, 0x0C, 0x50, 0x36, 0x02, 0x5F, 0x90, 0xFB, 0xFF, 0xCD, 0x0E, 0xFF,
		0x63, 0x03, 0xC1, 0x59, 0x0C, 0xDD, 0x8E, 0xF5, 0xEF, 0x7C, 0x06, 0xC0, 0xDB, 0x02, 0x5E, 0x82,
		0xE9, 0x20, 0x64, 0x03, 0xBE, 0xCA, 0x07, 0xEF, 0xA3, 0x0D, 0xA0, 0x76, 0xF3, 0x0F, 0x3E, 0x09,
		0x00, 0x30, 0xFD, 0x0F, 0x89, 0xFC, 0x20, 0xCA, 0xF3, 0x40, 0xC5, 0x0A, 0x0F, 0xC8, 0xFC, 0x40,
		0x62, 0xF5, 0x4F, 0xE9, 0x00, 0x3F, 0x23, 0x05, 0x0E, 0x72, 0x0A, 0xDE, 0x92, 0xF1, 0x40, 0x1C,
		0x04, 0xDE, 0xE1, 0x05, 0x70, 0xA3, 0xFF, 0x9F, 0xDE, 0x0A, 0x50, 0x36, 0x0D, 0xBF, 0xDF, 0x0B,
		0x7F, 0x54, 0xFC, 0x3F, 0xC9, 0x06, 0x41, 0x01, 0x0A, 0xF0, 0x39, 0xF3, 0x80, 0x47, 0x06, 0x0F,
		0x09, 0x00, 0xDF, 0x66, 0x06, 0x10, 0x6C, 0x07, 0xBF, 0x44, 0xF9, 0x3F, 0x9B, 0xF7, 0x80, 0x20,
		0x0F, 0x00, 0x58, 0x09, 0xC0, 0xE6, 0xF9, 0x10, 0x3D, 0xFE, 0x4F, 0xB3, 0xF4, 0xBF, 0xF0, 0xFE,
		0x0F, 0xC2, 0x08, 0x80, 0xA8, 0xFE, 0x10, 0x2F, 0x05, 0x10, 0x94, 0xEF, 0x5F, 0xC1, 0xFD, 0x5E,
		0x6E, 0xF9, 0x8E, 0xE3, 0xF4, 0xF0, 0x65, 0xFA, 0xDF, 0xF3, 0x12, 0x9F, 0x83, 0x06, 0x30, 0x1F,
		0xFB, 0xA0, 0x2E, 0xF6, 0x5F, 0xFF, 0xFC, 0xEF, 0xBA, 0xF5, 0x8F, 0x7B, 0xF7, 0x5F, 0x23, 0xFB,
		0x40, 0xD4, 0xF6, 0x70, 0xF9, 0x08, 0x00, 0xBD, 0xF8, 0x10, 0x12, 0x00, 0x30, 0x75, 0xFA, 0x10,
		0xCC, 0x0C, 0xCF, 0xA8, 0x0A, 0x9F, 0x7F, 0xE8, 0x70, 0x34, 0x19, 0x3F, 0xEF, 0xFB, 0x70, 0xD1,
		0x09, 0x3E, 0x94, 0x00, 0xC0, 0x07, 0xF3, 0x80, 0x23, 0xF0, 0x90, 0x75, 0xFA, 0x20, 0x31, 0xFE,
		0x80, 0xFF, 0xF7, 0x9F, 0x11, 0xFC, 0xD0, 0x6F, 0x00, 0x3F, 0xE6, 0xF2, 0x9E, 0xF8, 0x10, 0xE0,
		0xFF, 0xF3, 0xF0, 0xCD, 0xF8, 0xB0, 0x54, 0x04, 0x40, 0x52, 0x0F, 0x30, 0x2E, 0x0B, 0xAF, 0x36,
		0xF1, 0x71, 0xB1, 0xFB, 0xDE, 0xF5, 0xFB, 0xB0, 0x17, 0x0D, 0x7F, 0x31, 0xF9, 0xEF, 0xFC, 0xFE,
		0x3F, 0x3F, 0xF6, 0x8F, 0x92, 0xF8, 0x31, 0x00, 0xF4, 0x50, 0x7F, 0xF5, 0x5F, 0xA3, 0xFE, 0xAF,
		0x90, 0xFD, 0x6F, 0xF4, 0x09, 0x6F, 0xA6, 0xF1, 0xB0, 0x38, 0x0D, 0xFF, 0x12, 0x15, 0x5F, 0x61,
		0x04, 0x90, 0x2F, 0x03, 0x21, 0x3E, 0xF3, 0x80, 0xF5, 0x08, 0x30, 0x24, 0x0F, 0x61, 0x09, 0x0D,
		0x50, 0x3A, 0xFB, 0x0F, 0xAB, 0xED, 0x40, 0x47, 0x10, 0x20, 0x88, 0xE9, 0xAF, 0x31, 0xFA, 0xD0,
		0x7A, 0x02, 0x60, 0xA0, 0x03, 0xEE, 0x6C, 0x09, 0x50, 0x35, 0x0A, 0xA0, 0x1F, 0x0A, 0xB0, 0x1E,
		0x0D, 0xFE, 0x35, 0x04, 0xEF, 0x6E, 0x0D, 0xBF, 0xB7, 0x04, 0x10, 0x2A, 0x0E, 0x90, 0x53, 0x08,
		0x4E, 0xCB}

	c := &crypto.Crypto{}
	_, err := c.ImportPublicKey(pubKey)
	require.EqualError(t, err, foundation.FoundationErrorHandleStatus(-223).Error())
}
