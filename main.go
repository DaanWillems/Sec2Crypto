package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/crypt", cryptHandler)

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(htmlTemplate))
}

func cryptHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if len(r.Form["message"][0]) > 0 {
		w.Write([]byte((encrypt(r.Form["pass"][0], r.Form["message"][0]))))
	} else {
		w.Write([]byte(decrypt(r.Form["pass"][0], "0zvJ3LiQCwltILKV3TTo3WgfBDiQokPk")))
	}
}

func encrypt(key string, text string) string {

	key = createHash(key)

	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key string, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
	key = createHash(key)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

var htmlTemplate = `
	<div>
		<form action="/crypt" method="post">
			Naam: <input type="text" name="name"></input></br>
			Geheime tekst: <textarea name="message" style="height: 300px; width: 300px;"></textarea></br>
			Wachtwoord: <input type="text" name="pass"></input></br>
			<button type="submit">Verstuur</button>
		</form>
	</div>
`
