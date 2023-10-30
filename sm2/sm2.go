// Copyright 2023 The mkcert Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sm2

import (
	"crypto"

	gmssl "github.com/GmSSL/GmSSL-Go"
)

type PrivateKey struct {
	key *gmssl.Sm2Key
}

func GenerateKey() (crypto.PrivateKey, error) {
	key, error := gmssl.GenerateSm2Key()

	priv := &PrivateKey{
		key: key,
	}

	return priv, error
}

// Public returns the public key corresponding to priv.
func (priv *PrivateKey) Public() crypto.PublicKey {
	return priv
}
