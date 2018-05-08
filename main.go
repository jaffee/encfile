package main

import (
	"crypto/aes"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jaffee/commandeer"
	"github.com/pkg/errors"
)

type Main struct {
	Filename string `help:"filename to encrypt or decrypt"`
	Decrypt  bool   `help:"set to decrypt - otherwise, encrypt"`
	Key      string `help:"key for encryption/decryption. Must be 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256."`
}

func NewMain() *Main {
	return &Main{}
}

func (m *Main) Run() error {
	fname := m.Filename
	f, err := os.Open(fname)
	if err != nil {
		return errors.Wrap(err, "opening file")
	}
	fbytes, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "reading file")
	}

	ciph, err := aes.NewCipher([]byte(m.Key))
	if err != nil {
		return errors.Wrap(err, "getting cipher")
	}

	if m.Decrypt {
		ciph.Decrypt(fbytes, fbytes)
		fname = strings.TrimSuffix(fname, ".encrypted") + ".decrypted"
	} else {
		ciph.Encrypt(fbytes, fbytes)
		fname += ".encrypted"
	}

	err = ioutil.WriteFile(fname, fbytes, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "writing file")
	}

	return nil
}

func main() {
	err := commandeer.Run(NewMain())
	if err != nil {
		log.Fatal(err)
	}
}
