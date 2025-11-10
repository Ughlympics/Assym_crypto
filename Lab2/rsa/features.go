package rsa

import (
	"fmt"
	"math/big"
	key "rsa/rsa/key_schedule"
)

type User struct {
	ID   int
	Name string

	p *big.Int
	d *big.Int
	q *big.Int

	N *big.Int
	E *big.Int
}

var userCounter = 0
var users []*User

func NewUser(name string) (*User, error) {
	userCounter++
	var p, q, n *big.Int

	for {
		p, q, _ = key.GenKey(32)
		n = new(big.Int).Mul(p, q)
		isBigger := true
		for _, u := range users {
			if n.Cmp(u.N) <= 0 {
				isBigger = false
				break
			}
		}

		if isBigger {
			break
		}
	}

	phi := new(big.Int).Mul(
		new(big.Int).Sub(p, big.NewInt(1)),
		new(big.Int).Sub(q, big.NewInt(1)),
	)

	e := big.NewInt(65537)

	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		return nil, fmt.Errorf("ups, cannot compute modular inverse for e=%s and phi=%s", e.String(), phi.String())
	}

	u := &User{
		ID:   userCounter,
		Name: name,
		p:    p,
		q:    q,
		d:    d,
		N:    n,
		E:    e,
	}

	users = append(users, u)

	return u, nil
}

func (u *User) EncryptUser(message *big.Int) *big.Int {
	if u.N == nil || u.E == nil {
		panic("Encrypt: user has no public key")
	}

	ciphertext := new(big.Int).Exp(message, u.E, u.N)
	return ciphertext
}

func Encrypt(key, message *big.Int) *big.Int {
	ciphertext := new(big.Int).Exp(message, big.NewInt(65537), key)
	return ciphertext
}

func (u *User) DecryptUser(ciphertext *big.Int) *big.Int {
	if u.d == nil || u.N == nil {
		panic("Decrypt: user has no private key")
	}

	plaintext := new(big.Int).Exp(ciphertext, u.d, u.N)
	return plaintext
}

func Decrypt(key, ciphertext, d *big.Int) *big.Int {
	plaintext := new(big.Int).Exp(ciphertext, d, key)
	return plaintext
}

func (u *User) DigitalSign(message *big.Int) *big.Int {
	if u.d == nil || u.N == nil {
		panic("DigitalSign: user has no private key")
	}
	signature := new(big.Int).Exp(message, u.d, u.N)
	return signature
}

func (u *User) VerifySign(signature *big.Int) *big.Int {
	if u.E == nil || u.N == nil {
		panic("VerifySign: user has no public key")
	}
	verifiedMessage := new(big.Int).Exp(signature, u.E, u.N)
	return verifiedMessage
}
