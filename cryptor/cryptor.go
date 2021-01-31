package cryptor

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	. "github.com/alameddinc/passwordBook-golang/record"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Cryptor struct {
	Key *rsa.PrivateKey
}

var FILENAME = "store.dat"

func (c *Cryptor) Generator() error {
	var err error
	c.Key, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	pemPrivateFile, err := os.Create("private_key.pem")
	pemPrivateBlok := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(c.Key),
	}
	err = pem.Encode(pemPrivateFile, &pemPrivateBlok)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pemPrivateFile.Close()
	return nil
}

func (c *Cryptor) Load() error {
	privateKeyFile, err := os.Open("private_key.pem")
	if err != nil {
		fmt.Println(err)
		return c.Generator()
	}
	pemfileinfo, _ := privateKeyFile.Stat()

	pembytes := make([]byte, pemfileinfo.Size())
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode(pembytes)
	privateKeyFile.Close()

	c.Key, err = x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		fmt.Println(err)
		return errors.New("Can not readed")
	}
	return nil
}

func (c *Cryptor) Encoded() error {
	var record Record
	f, err := os.OpenFile(FILENAME,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	record.Init()
	storedText, _ := json.Marshal(record)
	encryptedBytes, _ := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&c.Key.PublicKey,
		storedText,
		nil)
	if _, err := f.WriteString(string(encryptedBytes) + "**_"); err != nil {
		log.Println(err)
	}
	return nil
}

func (c *Cryptor) Unload(s string) {
	records := []Record{}
	dat, _ := ioutil.ReadFile(FILENAME)
	arr := strings.Split(string(dat), "**_")
	for i := 0; i < len(arr)-1; i++ {
		record := Record{}
		decryptedBytes, err := c.Key.Decrypt(nil, []byte(arr[i]), &rsa.OAEPOptions{Hash: crypto.SHA256})
		if err != nil {
			panic(err)
		}
		json.Unmarshal(decryptedBytes, &record)
		if record.URL == s {
			records = append(records, record)
		}
	}

	fmt.Printf("Username\t Password \t Created At\n")
	for _, v := range records {
		fmt.Printf("%s\t%s\t%v\n", v.Username, v.Password, v.CreatedAt)
	}
}
