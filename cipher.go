package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cipher struct {
	gcm cipher.AEAD
}

func NewCipher() (Cipher, error) {
	key := ""

	aes, err := aes.NewCipher([]byte(mdHashing(key)))
	if err != nil {
		return Cipher{}, err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return Cipher{}, err
	}

	return Cipher{gcm: gcm}, nil
}

func (c *Cipher) encrypt(plain string) {
	nonce := make([]byte, c.gcm.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipherTxt := c.gcm.Seal(nonce, nonce, []byte(plain), nil)

	if _, err := os.Stat("data"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("data", os.ModePerm)
		if err != nil {
			log.Fatalln("Failed to create dir ", err)
		}
	}

	ex, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.OpenFile(filepath.Join(filepath.Dir(ex), "data", "uData"), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Could not open file: ", err)
	}
	if _, err := f.Write(cipherTxt); err != nil {
		log.Fatalln("Failed to write ", err)
	}
	if err := f.Close(); err != nil {
		log.Fatalln("Failed to close file ", err)
	}
}

func (c *Cipher) decrypt() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.OpenFile(filepath.Join(filepath.Dir(ex), "data", "uData"), os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln("Could not open file: ", err)
	}

	cipherTxt, err := io.ReadAll(f)
	if err != nil {
		log.Fatalln("Failed to read file:, ", err)
	}

	nonceSize := c.gcm.NonceSize()
	nonce, cipherTxt := cipherTxt[:nonceSize], cipherTxt[nonceSize:]

	original, err := c.gcm.Open(nil, nonce, cipherTxt, nil)
	if err != nil {
		log.Fatalln("Failed to decrypt: ", err)
	}

	return string(original)
}

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
}
