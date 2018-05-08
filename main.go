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

// Main holds the command line options for encfile.
type Main struct {
	Filename string `help:"Name of file to encrypt or decrypt."`
	Decrypt  bool   `help:"Set for decryption - otherwise, encrypt."`
	Key      string `help:"Key for encryption/decryption. Must be 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256."`
}

// NewMain gets a new Main with the default options.
func NewMain() *Main {
	return &Main{}
}

// Run runs the file encryption/decryption utiltity.
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
	// Set up flags and run Main.
	err := commandeer.Run(NewMain())
	if err != nil {
		log.Fatal(err)
	}
}
