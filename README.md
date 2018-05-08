# encfile
Simple, command line file encryption program in ~60 lines of Go.


I wanted to create an encrypted image of a directory, but Disk Utility on my Mac was not behaving. Instead of searching the internet for an encryption utility like a sane person, I wrote this tiny program. I think it's a nice little example of how to build a command line tool in Go.

# install
Should be as easy as `go get github.com/jaffee/encfile`. 

# usage 

```
$ encfile -h
Usage of encfile:
  -decrypt
    	set to decrypt - otherwise, encrypt
  -filename string
    	filename to encrypt or decrypt
  -key string
    	key for encryption/decryption. Must be 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
```

# contributions
Sure, send a PR. Would be cool if you could read the key from a file rather than having to pass it in at the command line. *cough* hint *cough*
