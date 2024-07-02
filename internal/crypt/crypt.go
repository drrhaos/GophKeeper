// Package crypt предназначен для шифрования данных.
package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"gophkeeper/pkg/proto"
)

// EncField шифрует данные
func EncField(dec *proto.FieldKeep, secretKey string) (enc *proto.FieldKeep) {
	enc = &proto.FieldKeep{
		Name:       Encrypt(dec.Name, secretKey),
		Login:      Encrypt(dec.Login, secretKey),
		Password:   Encrypt(dec.Password, secretKey),
		Data:       Encrypt(dec.Data, secretKey),
		CardNumber: Encrypt(dec.CardNumber, secretKey),
		CardCVC:    Encrypt(dec.CardCVC, secretKey),
		CardDate:   Encrypt(dec.CardDate, secretKey),
		CardOwner:  Encrypt(dec.CardOwner, secretKey),
		FileName:   Encrypt(dec.FileName, secretKey),
	}
	return enc
}

// DecField расшифровывает данные
func DecField(enc *proto.FieldKeep, secretKey string) (dec *proto.FieldKeep) {

	dec = &proto.FieldKeep{
		Name:       Decrypt(enc.Name, secretKey),
		Login:      Decrypt(enc.Login, secretKey),
		Password:   Decrypt(enc.Password, secretKey),
		Data:       Decrypt(enc.Data, secretKey),
		CardNumber: Decrypt(enc.CardNumber, secretKey),
		CardCVC:    Decrypt(enc.CardCVC, secretKey),
		CardDate:   Decrypt(enc.CardDate, secretKey),
		CardOwner:  Decrypt(enc.CardOwner, secretKey),
		FileName:   Decrypt(enc.FileName, secretKey),
	}

	return dec
}

// Encrypt шифрование сообщения открытым ключом
func Encrypt(plaintext string, secretKey string) string {
	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return ""
	}
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return ""
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return ""
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

// Decrypt дешифрование сообщения закрытым ключом
func Decrypt(data string, secretKey string) string {
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if len(ciphertext) == 0 {
		return ""
	}
	if err != nil {
		return ""
	}

	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return ""
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return ""
	}

	nonceSize := gcm.NonceSize()

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), ciphertext, nil)
	if err != nil {
		return ""
	}

	return string(plaintext)
}
