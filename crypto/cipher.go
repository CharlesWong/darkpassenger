package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rc4"
	"golang.org/x/crypto/blowfish"
	"golang.org/x/crypto/cast5"
	"log"
)

type Cipher struct {
	enc cipher.Stream
	dec cipher.Stream
}

type chiperCreator func(key []byte) (*Cipher, error)

var cipherMap = map[string]chiperCreator{
	"rc4":         newRC4Cipher,
	"aes-256-cfb": newAES256CFBCipher,
	"aes-128-cfb": newAES128CFBCipher,
	"aes-192-cfb": newAES192CFBCipher,
	"des-cfb":     newDESCipher,
	"bf-cfb":      newBFCipher,
	"cast5-cfb":   newCast5Cipher,
}

func secretToKey(secret []byte, size int) []byte {
	// size mod 16 must be 0
	h := md5.New()
	buf := make([]byte, size)
	count := size / md5.Size
	// repeatly fill the key with the secret
	for i := 0; i < count; i++ {
		h.Write(secret)
		copy(buf[md5.Size*i:md5.Size*(i+1)-1], h.Sum(nil))
	}
	return buf
}

func newRC4Cipher(secret []byte) (*Cipher, error) {
	ec, err := rc4.NewCipher(secretToKey(secret, 16))
	if err != nil {
		return nil, err
	}
	dc := *ec

	return &Cipher{ec, &dc}, nil
}

func newAES256CFBCipher(secret []byte) (*Cipher, error) {
	return newAESCipher(secret, 32)
}

func newAES192CFBCipher(secret []byte) (*Cipher, error) {
	return newAESCipher(secret, 24)
}

func newAES128CFBCipher(secret []byte) (*Cipher, error) {
	return newAESCipher(secret, 16)
}

func newAESCipher(secret []byte, b int) (*Cipher, error) {
	key := secretToKey(secret, b)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return newBlockCipher(block, key)
}

func newDESCipher(secret []byte) (*Cipher, error) {
	key := secretToKey(secret, 8)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return newBlockCipher(block, key)
}

func newBFCipher(secret []byte) (*Cipher, error) {
	key := secretToKey(secret, 16)
	block, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return newBlockCipher(block, key)
}

func newCast5Cipher(secret []byte) (*Cipher, error) {
	key := secretToKey(secret, 16)
	block, err := cast5.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return newBlockCipher(block, key)
}

func newBlockCipher(block cipher.Block, key []byte) (*Cipher, error) {
	ec := cipher.NewCFBEncrypter(block, key[:block.BlockSize()])
	dc := cipher.NewCFBDecrypter(block, key[:block.BlockSize()])

	return &Cipher{ec, dc}, nil
}

func NewCipher(cryptoMethod string, secret []byte) *Cipher {
	cc := cipherMap[cryptoMethod]
	if cc == nil {
		log.Fatalf("unsupported crypto method %s", cryptoMethod)
	}
	c, err := cc(secret)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func (c *Cipher) Encrypt(dst, src []byte) {
	c.enc.XORKeyStream(dst, src)
}

func (c *Cipher) Decrypt(dst, src []byte) {
	c.dec.XORKeyStream(dst, src)
}
