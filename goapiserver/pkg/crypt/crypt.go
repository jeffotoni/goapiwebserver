// Back-End in Go server
// @jeffotoni
// 2019-01-09

package crypt

//
//
//
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	b64 "encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"

	"golang.org/x/crypto/bcrypt"

	// "strings"
	"time"
)

var (
	HASH_SALT      = "$%&*#@..jef3845#@874nw.x48x#@54"
	SHA1_SALT      = "@..39cjeyx#@34848.x48x#@938xw3."
	letters        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	UKK_KEY_CIPHER = []byte("KEUEUXUWYEJDUEUE7383DJDU838CSDXN")
)

// See alternate IV creation from ciphertext below
// var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
func Encrypt(textString string) (string, error) {

	text := []byte(textString)

	block, err := aes.NewCipher(UKK_KEY_CIPHER)

	if err != nil {

		return "", err
	}

	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(crand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return string(ciphertext), nil
}

//
// [Decrypt description]
// @param {[type]} key  [description]
// @param {[type]} text []byte)       ([]byte, error [description]
func Decrypt(textString string) (string, error) {

	text := []byte(textString)

	block, err := aes.NewCipher(UKK_KEY_CIPHER)

	if err != nil {
		return "", err
	}

	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	data, err := base64.StdEncoding.DecodeString(string(text))

	if err != nil {
		return "", err
	}

	return string(data), nil
}

//
//
//
func HashSeed(n int) string {

	b := make([]rune, n)
	for i := range b {

		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

//
//
//
func Random(min, max int) int { rand.Seed(time.Now().Unix()); return rand.Intn(max-min) + min }

// generating single seed, can not repeat, even calling in almost
// the same time interval
func RandUid() string {

	// generate 64 bits timestamp
	unix64bits := uint64(time.Now().UTC().UnixNano())

	buff := make([]byte, 128)

	numRead, err := rand.Read(buff)

	if numRead != len(buff) || err != nil {
		panic(err)
	}

	unixUid := fmt.Sprintf("%x", unix64bits)

	//fmt.Println(unixUid)
	//return unixUid

	unixUid = GSha1(Blowfish(unixUid))
	return unixUid
}

//
//
//
func Blowfish(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {

		panic(err)
	}

	return string(bytes)
}

//
//
//
func CheckBlowfish(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//
//
//
func GHash256(password string) string {

	pass := password + HASH_SALT

	h := sha256.New()
	h.Write([]byte(pass))

	return fmt.Sprintf("%x", h.Sum(nil))
}

//
//
//
func GHash512(password string) string {

	pass := password + HASH_SALT

	h := sha512.Sum512([]byte(pass))

	return fmt.Sprintf("%x", h)
}

//
//
//
func GSha1(key string) string {

	data := []byte(key + SHA1_SALT)
	return (fmt.Sprintf("%x", sha1.Sum(data)))
}

//
func Md5(text string) string {

	h := md5.New()
	io.WriteString(h, text)
	return (fmt.Sprintf("%x", h.Sum(nil)))
}

//
func HashFile(filePath string) (string, error) {

	var returnMD5String string

	file, err := os.Open(filePath)
	if err != nil {

		return returnMD5String, err
	}

	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {

		return returnMD5String, err
	}

	hashInBytes := hash.Sum(nil)[:16]

	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}

//
func UkkBase64Encode(textString string) string {

	text := []byte(textString)

	sEnc := b64.URLEncoding.EncodeToString(text)

	return sEnc
}

//
func UkkBase64Decode(textString string) string {

	//text := []byte(textString)

	sDec, _ := b64.URLEncoding.DecodeString(textString)

	return string(sDec)
}
