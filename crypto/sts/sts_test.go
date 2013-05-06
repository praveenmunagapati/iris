// Iris - Distributed Messaging Framework
// Copyright 2013 Peter Szilagyi. All rights reserved.
//
// Iris is dual licensed: you can redistribute it and/or modify it under the
// terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// The framework is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// Alternatively, the Iris framework may be used in accordance with the terms
// and conditions contained in a signed written agreement between you and the
// author(s).
//
// Author: peterke@gmail.com (Peter Szilagyi)
package sts

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	_ "crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"
	"math/big"
	"testing"
)

type stsTest struct {
	cipher func([]byte) (cipher.Block, error)
	bits   int
	hash   crypto.Hash

	group          *big.Int
	generator      *big.Int
	iniExponent    *big.Int
	iniExponential *big.Int
	accExponent    *big.Int
	accExponential *big.Int
}

var stsTests = []stsTest{
	{
		aes.NewCipher,
		128,
		crypto.SHA256,
		new(big.Int).SetBytes([]byte{
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			0xc9, 0x0f, 0xda, 0xa2, 0x21, 0x68, 0xc2, 0x34,
			0xc4, 0xc6, 0x62, 0x8b, 0x80, 0xdc, 0x1c, 0xd1,
			0x29, 0x02, 0x4e, 0x08, 0x8a, 0x67, 0xcc, 0x74,
			0x02, 0x0b, 0xbe, 0xa6, 0x3b, 0x13, 0x9b, 0x22,
			0x51, 0x4a, 0x08, 0x79, 0x8e, 0x34, 0x04, 0xdd,
			0xef, 0x95, 0x19, 0xb3, 0xcd, 0x3a, 0x43, 0x1b,
			0x30, 0x2b, 0x0a, 0x6d, 0xf2, 0x5f, 0x14, 0x37,
			0x4f, 0xe1, 0x35, 0x6d, 0x6d, 0x51, 0xc2, 0x45,
			0xe4, 0x85, 0xb5, 0x76, 0x62, 0x5e, 0x7e, 0xc6,
			0xf4, 0x4c, 0x42, 0xe9, 0xa6, 0x3a, 0x36, 0x20,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		}),
		big.NewInt(2),
		new(big.Int).SetBytes([]byte{
			0x4a, 0xe7, 0xd7, 0xe8, 0xce, 0x57, 0xd8, 0xe0,
			0x37, 0x61, 0xeb, 0x88, 0x38, 0x9f, 0x61, 0xf4,
			0x69, 0x1d, 0xc8, 0xb1, 0x49, 0x39, 0xe9, 0x44,
			0x74, 0x31, 0x49, 0xd4, 0x5b, 0xe1, 0xb8, 0x35,
			0xcf, 0x60, 0xdd, 0x10, 0x64, 0x90, 0x96, 0x5c,
			0xec, 0x46, 0xa7, 0x06, 0x7b, 0x35, 0x8f, 0x07,
			0x1a, 0x7d, 0xee, 0xd0, 0xfd, 0x57, 0x2b, 0xa6,
			0x9d, 0xb9, 0xb2, 0xcf, 0xa6, 0xa3, 0x10, 0x70,
			0x7f, 0x63, 0x82, 0x3d, 0xb4, 0xb9, 0x35, 0x49,
			0x50, 0x89, 0x7e, 0x8b, 0x53, 0xcf, 0x1b, 0x16,
			0x0a, 0x0a, 0x33, 0xb7, 0xb6, 0xc0, 0x1e, 0x01,
			0xcb, 0x1c, 0xa4, 0x99, 0xe6, 0x6c, 0x89, 0xe9,
		}),
		new(big.Int).SetBytes([]byte{
			0xe0, 0x68, 0x84, 0x01, 0xd4, 0xe8, 0x2f, 0x4c,
			0x98, 0x89, 0x0e, 0x62, 0x8f, 0x9a, 0xa6, 0xec,
			0x31, 0xf5, 0xda, 0xf9, 0xde, 0xf9, 0x9a, 0x3c,
			0x91, 0x75, 0x14, 0x83, 0xed, 0x65, 0x34, 0x7b,
			0xb2, 0x73, 0xfd, 0xfb, 0x42, 0x3c, 0x3f, 0xd4,
			0x2f, 0xcf, 0x3a, 0x71, 0x91, 0x12, 0x71, 0x80,
			0x5b, 0x70, 0x18, 0xa3, 0xc6, 0xac, 0x76, 0x82,
			0xb4, 0x36, 0x9a, 0xcf, 0x03, 0x16, 0x0d, 0x3e,
			0xdc, 0x34, 0x51, 0x1a, 0xb0, 0x57, 0x1e, 0xd6,
			0x25, 0x59, 0x7c, 0xaf, 0x13, 0x94, 0x31, 0x7e,
			0xa1, 0xbc, 0x48, 0x36, 0x60, 0x70, 0x65, 0x06,
			0x2a, 0xe7, 0x7c, 0xe5, 0xa9, 0x03, 0x33, 0x85,
		}),
		new(big.Int).SetBytes([]byte{
			0xcc, 0x6d, 0x91, 0x66, 0x68, 0x3e, 0x18, 0xce,
			0xb1, 0xab, 0xaf, 0x02, 0xa6, 0x6e, 0x95, 0x02,
			0xa5, 0x9a, 0x91, 0xb1, 0xc4, 0x28, 0xa7, 0x3f,
			0xa8, 0x06, 0xb3, 0x06, 0xee, 0xaa, 0xee, 0x2b,
			0xed, 0xc9, 0x62, 0xd5, 0xe8, 0x3e, 0x35, 0x49,
			0x0b, 0xc1, 0x9c, 0xa0, 0x32, 0x84, 0x82, 0x14,
			0x7c, 0xc0, 0x1a, 0x15, 0x74, 0x6d, 0x6f, 0xe8,
			0xc0, 0xc8, 0x58, 0x15, 0x04, 0x41, 0xe4, 0x07,
			0x49, 0xad, 0xd2, 0x21, 0xa7, 0x61, 0xa4, 0x0a,
			0x43, 0xd9, 0xbb, 0x47, 0x9a, 0xae, 0x38, 0x7c,
			0x65, 0xd1, 0x62, 0x50, 0xb0, 0x86, 0xbb, 0x90,
			0xde, 0xd4, 0x05, 0xf1, 0xce, 0x52, 0xf4, 0x8d,
		}),
		new(big.Int).SetBytes([]byte{
			0x7e, 0x46, 0x4b, 0x27, 0x66, 0xbf, 0x89, 0x27,
			0x39, 0x85, 0x65, 0x74, 0xeb, 0x28, 0x0c, 0x75,
			0x63, 0x37, 0x72, 0xed, 0xb8, 0xb2, 0x02, 0x59,
			0x96, 0x67, 0xf8, 0xda, 0x91, 0xcf, 0x32, 0x95,
			0x41, 0x63, 0x70, 0x3a, 0x56, 0x4e, 0x76, 0x82,
			0x0a, 0x2e, 0x88, 0xaa, 0xb2, 0xc4, 0x6b, 0x66,
			0x73, 0x95, 0x5b, 0x95, 0xb4, 0x0b, 0xe4, 0x70,
			0xd4, 0x73, 0x1e, 0x0f, 0x12, 0x2f, 0x7c, 0x40,
			0x2e, 0x53, 0xa6, 0x2a, 0x79, 0x8e, 0x08, 0x1d,
			0xa5, 0xd2, 0xc7, 0x90, 0x1a, 0x5a, 0x1a, 0x87,
			0xa3, 0xac, 0xc4, 0x1e, 0x16, 0xa6, 0x94, 0xf9,
			0xd7, 0xbe, 0x56, 0x00, 0xc9, 0x78, 0x65, 0xa2,
		}),
	},
	{
		aes.NewCipher,
		192,
		crypto.MD5,
		new(big.Int).SetBytes([]byte{
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			0xc9, 0x0f, 0xda, 0xa2, 0x21, 0x68, 0xc2, 0x34,
			0xc4, 0xc6, 0x62, 0x8b, 0x80, 0xdc, 0x1c, 0xd1,
			0x29, 0x02, 0x4e, 0x08, 0x8a, 0x67, 0xcc, 0x74,
			0x02, 0x0b, 0xbe, 0xa6, 0x3b, 0x13, 0x9b, 0x22,
			0x51, 0x4a, 0x08, 0x79, 0x8e, 0x34, 0x04, 0xdd,
			0xef, 0x95, 0x19, 0xb3, 0xcd, 0x3a, 0x43, 0x1b,
			0x30, 0x2b, 0x0a, 0x6d, 0xf2, 0x5f, 0x14, 0x37,
			0x4f, 0xe1, 0x35, 0x6d, 0x6d, 0x51, 0xc2, 0x45,
			0xe4, 0x85, 0xb5, 0x76, 0x62, 0x5e, 0x7e, 0xc6,
			0xf4, 0x4c, 0x42, 0xe9, 0xa6, 0x3a, 0x36, 0x20,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		}),
		big.NewInt(2),
		new(big.Int).SetBytes([]byte{
			0x4a, 0xe7, 0xd7, 0xe8, 0xce, 0x57, 0xd8, 0xe0,
			0x37, 0x61, 0xeb, 0x88, 0x38, 0x9f, 0x61, 0xf4,
			0x69, 0x1d, 0xc8, 0xb1, 0x49, 0x39, 0xe9, 0x44,
			0x74, 0x31, 0x49, 0xd4, 0x5b, 0xe1, 0xb8, 0x35,
			0xcf, 0x60, 0xdd, 0x10, 0x64, 0x90, 0x96, 0x5c,
			0xec, 0x46, 0xa7, 0x06, 0x7b, 0x35, 0x8f, 0x07,
			0x1a, 0x7d, 0xee, 0xd0, 0xfd, 0x57, 0x2b, 0xa6,
			0x9d, 0xb9, 0xb2, 0xcf, 0xa6, 0xa3, 0x10, 0x70,
			0x7f, 0x63, 0x82, 0x3d, 0xb4, 0xb9, 0x35, 0x49,
			0x50, 0x89, 0x7e, 0x8b, 0x53, 0xcf, 0x1b, 0x16,
			0x0a, 0x0a, 0x33, 0xb7, 0xb6, 0xc0, 0x1e, 0x01,
			0xcb, 0x1c, 0xa4, 0x99, 0xe6, 0x6c, 0x89, 0xe9,
		}),
		new(big.Int).SetBytes([]byte{
			0xe0, 0x68, 0x84, 0x01, 0xd4, 0xe8, 0x2f, 0x4c,
			0x98, 0x89, 0x0e, 0x62, 0x8f, 0x9a, 0xa6, 0xec,
			0x31, 0xf5, 0xda, 0xf9, 0xde, 0xf9, 0x9a, 0x3c,
			0x91, 0x75, 0x14, 0x83, 0xed, 0x65, 0x34, 0x7b,
			0xb2, 0x73, 0xfd, 0xfb, 0x42, 0x3c, 0x3f, 0xd4,
			0x2f, 0xcf, 0x3a, 0x71, 0x91, 0x12, 0x71, 0x80,
			0x5b, 0x70, 0x18, 0xa3, 0xc6, 0xac, 0x76, 0x82,
			0xb4, 0x36, 0x9a, 0xcf, 0x03, 0x16, 0x0d, 0x3e,
			0xdc, 0x34, 0x51, 0x1a, 0xb0, 0x57, 0x1e, 0xd6,
			0x25, 0x59, 0x7c, 0xaf, 0x13, 0x94, 0x31, 0x7e,
			0xa1, 0xbc, 0x48, 0x36, 0x60, 0x70, 0x65, 0x06,
			0x2a, 0xe7, 0x7c, 0xe5, 0xa9, 0x03, 0x33, 0x85,
		}),
		new(big.Int).SetBytes([]byte{
			0xcc, 0x6d, 0x91, 0x66, 0x68, 0x3e, 0x18, 0xce,
			0xb1, 0xab, 0xaf, 0x02, 0xa6, 0x6e, 0x95, 0x02,
			0xa5, 0x9a, 0x91, 0xb1, 0xc4, 0x28, 0xa7, 0x3f,
			0xa8, 0x06, 0xb3, 0x06, 0xee, 0xaa, 0xee, 0x2b,
			0xed, 0xc9, 0x62, 0xd5, 0xe8, 0x3e, 0x35, 0x49,
			0x0b, 0xc1, 0x9c, 0xa0, 0x32, 0x84, 0x82, 0x14,
			0x7c, 0xc0, 0x1a, 0x15, 0x74, 0x6d, 0x6f, 0xe8,
			0xc0, 0xc8, 0x58, 0x15, 0x04, 0x41, 0xe4, 0x07,
			0x49, 0xad, 0xd2, 0x21, 0xa7, 0x61, 0xa4, 0x0a,
			0x43, 0xd9, 0xbb, 0x47, 0x9a, 0xae, 0x38, 0x7c,
			0x65, 0xd1, 0x62, 0x50, 0xb0, 0x86, 0xbb, 0x90,
			0xde, 0xd4, 0x05, 0xf1, 0xce, 0x52, 0xf4, 0x8d,
		}),
		new(big.Int).SetBytes([]byte{
			0x7e, 0x46, 0x4b, 0x27, 0x66, 0xbf, 0x89, 0x27,
			0x39, 0x85, 0x65, 0x74, 0xeb, 0x28, 0x0c, 0x75,
			0x63, 0x37, 0x72, 0xed, 0xb8, 0xb2, 0x02, 0x59,
			0x96, 0x67, 0xf8, 0xda, 0x91, 0xcf, 0x32, 0x95,
			0x41, 0x63, 0x70, 0x3a, 0x56, 0x4e, 0x76, 0x82,
			0x0a, 0x2e, 0x88, 0xaa, 0xb2, 0xc4, 0x6b, 0x66,
			0x73, 0x95, 0x5b, 0x95, 0xb4, 0x0b, 0xe4, 0x70,
			0xd4, 0x73, 0x1e, 0x0f, 0x12, 0x2f, 0x7c, 0x40,
			0x2e, 0x53, 0xa6, 0x2a, 0x79, 0x8e, 0x08, 0x1d,
			0xa5, 0xd2, 0xc7, 0x90, 0x1a, 0x5a, 0x1a, 0x87,
			0xa3, 0xac, 0xc4, 0x1e, 0x16, 0xa6, 0x94, 0xf9,
			0xd7, 0xbe, 0x56, 0x00, 0xc9, 0x78, 0x65, 0xa2,
		}),
	},
	{
		des.NewTripleDESCipher,
		192,
		crypto.SHA256,
		new(big.Int).SetBytes([]byte{
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			0xc9, 0x0f, 0xda, 0xa2, 0x21, 0x68, 0xc2, 0x34,
			0xc4, 0xc6, 0x62, 0x8b, 0x80, 0xdc, 0x1c, 0xd1,
			0x29, 0x02, 0x4e, 0x08, 0x8a, 0x67, 0xcc, 0x74,
			0x02, 0x0b, 0xbe, 0xa6, 0x3b, 0x13, 0x9b, 0x22,
			0x51, 0x4a, 0x08, 0x79, 0x8e, 0x34, 0x04, 0xdd,
			0xef, 0x95, 0x19, 0xb3, 0xcd, 0x3a, 0x43, 0x1b,
			0x30, 0x2b, 0x0a, 0x6d, 0xf2, 0x5f, 0x14, 0x37,
			0x4f, 0xe1, 0x35, 0x6d, 0x6d, 0x51, 0xc2, 0x45,
			0xe4, 0x85, 0xb5, 0x76, 0x62, 0x5e, 0x7e, 0xc6,
			0xf4, 0x4c, 0x42, 0xe9, 0xa6, 0x3a, 0x36, 0x20,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		}),
		big.NewInt(2),
		new(big.Int).SetBytes([]byte{
			0x4a, 0xe7, 0xd7, 0xe8, 0xce, 0x57, 0xd8, 0xe0,
			0x37, 0x61, 0xeb, 0x88, 0x38, 0x9f, 0x61, 0xf4,
			0x69, 0x1d, 0xc8, 0xb1, 0x49, 0x39, 0xe9, 0x44,
			0x74, 0x31, 0x49, 0xd4, 0x5b, 0xe1, 0xb8, 0x35,
			0xcf, 0x60, 0xdd, 0x10, 0x64, 0x90, 0x96, 0x5c,
			0xec, 0x46, 0xa7, 0x06, 0x7b, 0x35, 0x8f, 0x07,
			0x1a, 0x7d, 0xee, 0xd0, 0xfd, 0x57, 0x2b, 0xa6,
			0x9d, 0xb9, 0xb2, 0xcf, 0xa6, 0xa3, 0x10, 0x70,
			0x7f, 0x63, 0x82, 0x3d, 0xb4, 0xb9, 0x35, 0x49,
			0x50, 0x89, 0x7e, 0x8b, 0x53, 0xcf, 0x1b, 0x16,
			0x0a, 0x0a, 0x33, 0xb7, 0xb6, 0xc0, 0x1e, 0x01,
			0xcb, 0x1c, 0xa4, 0x99, 0xe6, 0x6c, 0x89, 0xe9,
		}),
		new(big.Int).SetBytes([]byte{
			0xe0, 0x68, 0x84, 0x01, 0xd4, 0xe8, 0x2f, 0x4c,
			0x98, 0x89, 0x0e, 0x62, 0x8f, 0x9a, 0xa6, 0xec,
			0x31, 0xf5, 0xda, 0xf9, 0xde, 0xf9, 0x9a, 0x3c,
			0x91, 0x75, 0x14, 0x83, 0xed, 0x65, 0x34, 0x7b,
			0xb2, 0x73, 0xfd, 0xfb, 0x42, 0x3c, 0x3f, 0xd4,
			0x2f, 0xcf, 0x3a, 0x71, 0x91, 0x12, 0x71, 0x80,
			0x5b, 0x70, 0x18, 0xa3, 0xc6, 0xac, 0x76, 0x82,
			0xb4, 0x36, 0x9a, 0xcf, 0x03, 0x16, 0x0d, 0x3e,
			0xdc, 0x34, 0x51, 0x1a, 0xb0, 0x57, 0x1e, 0xd6,
			0x25, 0x59, 0x7c, 0xaf, 0x13, 0x94, 0x31, 0x7e,
			0xa1, 0xbc, 0x48, 0x36, 0x60, 0x70, 0x65, 0x06,
			0x2a, 0xe7, 0x7c, 0xe5, 0xa9, 0x03, 0x33, 0x85,
		}),
		new(big.Int).SetBytes([]byte{
			0xcc, 0x6d, 0x91, 0x66, 0x68, 0x3e, 0x18, 0xce,
			0xb1, 0xab, 0xaf, 0x02, 0xa6, 0x6e, 0x95, 0x02,
			0xa5, 0x9a, 0x91, 0xb1, 0xc4, 0x28, 0xa7, 0x3f,
			0xa8, 0x06, 0xb3, 0x06, 0xee, 0xaa, 0xee, 0x2b,
			0xed, 0xc9, 0x62, 0xd5, 0xe8, 0x3e, 0x35, 0x49,
			0x0b, 0xc1, 0x9c, 0xa0, 0x32, 0x84, 0x82, 0x14,
			0x7c, 0xc0, 0x1a, 0x15, 0x74, 0x6d, 0x6f, 0xe8,
			0xc0, 0xc8, 0x58, 0x15, 0x04, 0x41, 0xe4, 0x07,
			0x49, 0xad, 0xd2, 0x21, 0xa7, 0x61, 0xa4, 0x0a,
			0x43, 0xd9, 0xbb, 0x47, 0x9a, 0xae, 0x38, 0x7c,
			0x65, 0xd1, 0x62, 0x50, 0xb0, 0x86, 0xbb, 0x90,
			0xde, 0xd4, 0x05, 0xf1, 0xce, 0x52, 0xf4, 0x8d,
		}),
		new(big.Int).SetBytes([]byte{
			0x7e, 0x46, 0x4b, 0x27, 0x66, 0xbf, 0x89, 0x27,
			0x39, 0x85, 0x65, 0x74, 0xeb, 0x28, 0x0c, 0x75,
			0x63, 0x37, 0x72, 0xed, 0xb8, 0xb2, 0x02, 0x59,
			0x96, 0x67, 0xf8, 0xda, 0x91, 0xcf, 0x32, 0x95,
			0x41, 0x63, 0x70, 0x3a, 0x56, 0x4e, 0x76, 0x82,
			0x0a, 0x2e, 0x88, 0xaa, 0xb2, 0xc4, 0x6b, 0x66,
			0x73, 0x95, 0x5b, 0x95, 0xb4, 0x0b, 0xe4, 0x70,
			0xd4, 0x73, 0x1e, 0x0f, 0x12, 0x2f, 0x7c, 0x40,
			0x2e, 0x53, 0xa6, 0x2a, 0x79, 0x8e, 0x08, 0x1d,
			0xa5, 0xd2, 0xc7, 0x90, 0x1a, 0x5a, 0x1a, 0x87,
			0xa3, 0xac, 0xc4, 0x1e, 0x16, 0xa6, 0x94, 0xf9,
			0xd7, 0xbe, 0x56, 0x00, 0xc9, 0x78, 0x65, 0xa2,
		}),
	},
	{
		aes.NewCipher,
		192,
		crypto.MD5,
		new(big.Int).SetBytes([]byte{
			0xea, 0x60, 0x76, 0xe5, 0x31, 0x2a, 0xa8, 0xa8,
			0x7d, 0xdd, 0x56, 0x81, 0x3b, 0xd9, 0x78, 0x65,
			0x4e, 0x4b, 0x8a, 0xb8, 0x33, 0x15, 0xd9, 0xa3,
			0x54, 0x94, 0xc8, 0x06, 0x3b, 0x20, 0xa6, 0x31,
			0x47, 0x66, 0x34, 0x0f, 0x7d, 0xa4, 0x09, 0x0f,
			0x87, 0x51, 0x6e, 0xa0, 0xf6, 0x3c, 0x30, 0x43,
			0x18, 0xb9, 0x62, 0x10, 0xe0, 0x90, 0xd1, 0xea,
			0x06, 0x9f, 0xde, 0xd0, 0xee, 0x2d, 0x79, 0x1d,
			0xbd, 0x4b, 0x00, 0x4c, 0xbb, 0x9d, 0xff, 0x9e,
			0xa1, 0x78, 0x58, 0x6c, 0x0d, 0x7f, 0x6f, 0x15,
			0xf4, 0xd9, 0xd8, 0x82, 0xa8, 0xef, 0x3f, 0x50,
			0x8b, 0x3a, 0xaa, 0xc7, 0x23, 0xab, 0x5a, 0x8c,
			0xd6, 0xac, 0xc5, 0xb9, 0xed, 0xe6, 0x93, 0xb1,
			0x25, 0xd4, 0x1a, 0x7b, 0x1a, 0xaf, 0x3a, 0x38,
			0xec, 0xb6, 0xa7, 0x1e, 0x62, 0x83, 0x71, 0x03,
			0x83, 0xf2, 0x34, 0xb0, 0x0d, 0x87, 0x4d, 0x1f,
		}),
		new(big.Int).SetBytes([]byte{
			0xe6, 0x7d, 0xcb, 0xcf, 0xb7, 0xde, 0xfd, 0x40,
			0xcf, 0x0f, 0x85, 0xa3, 0x9b, 0x8d, 0x9a, 0x87,
			0x94, 0x73, 0xe9, 0x3e, 0x52, 0xc0, 0x12, 0xab,
			0xb4, 0xba, 0x1d, 0x2d, 0x5f, 0x2b, 0x11, 0x57,
			0x5d, 0x1b, 0x5b, 0xd6, 0xeb, 0x0a, 0xe0, 0x7f,
			0x03, 0x4f, 0x61, 0x69, 0x1f, 0x6d, 0xfd, 0x29,
			0x8a, 0xe6, 0x64, 0xa7, 0x4d, 0xcc, 0x1f, 0xb3,
			0x83, 0xf7, 0x76, 0x58, 0x37, 0x6e, 0x55, 0x30,
			0xb7, 0xc9, 0x5a, 0x32, 0xcf, 0x8a, 0x7f, 0x23,
			0xee, 0x75, 0x5c, 0xac, 0xb8, 0x87, 0x32, 0xa3,
			0x87, 0x95, 0x24, 0x86, 0x72, 0xfc, 0xfb, 0x1b,
			0xf3, 0x63, 0x44, 0x02, 0x44, 0xb6, 0x5c, 0x01,
			0x1b, 0x76, 0xbc, 0xe8, 0xca, 0xf7, 0xdd, 0xc4,
			0x16, 0xa8, 0x87, 0x80, 0x0c, 0x76, 0xf5, 0x1d,
			0x7e, 0x37, 0xa2, 0xcb, 0x16, 0xb1, 0x4f, 0x61,
			0xde, 0x2c, 0x3b, 0x2a, 0xec, 0x4c, 0xfc, 0xea,
		}),
		new(big.Int).SetBytes([]byte{
			0x30, 0x3e, 0x56, 0x19, 0x06, 0x58, 0x1b, 0x62,
			0xdc, 0x9a, 0xae, 0x7a, 0x91, 0xc4, 0x70, 0xed,
			0xe3, 0xa3, 0x1c, 0xfb, 0xe0, 0x79, 0xe9, 0xca,
			0x23, 0x8e, 0x4b, 0x95, 0xb9, 0xcb, 0xee, 0xa9,
			0x08, 0xc4, 0xea, 0x5c, 0xcb, 0xf0, 0xc3, 0x2f,
			0xd6, 0x7c, 0x34, 0xd6, 0xd2, 0x5e, 0xe9, 0x1e,
			0x82, 0x13, 0x76, 0x37, 0x35, 0x2c, 0xef, 0xb4,
			0x5d, 0x7c, 0xa5, 0x47, 0x00, 0x36, 0xa0, 0x8f,
			0x60, 0x59, 0xe7, 0x93, 0x77, 0x49, 0x6b, 0xb5,
			0x2e, 0xce, 0xc9, 0x83, 0x54, 0x44, 0x7c, 0x5f,
			0x1a, 0x8a, 0xe8, 0x12, 0xa8, 0xe6, 0x1f, 0xed,
			0x48, 0x0e, 0xd5, 0x8a, 0xf4, 0x43, 0x46, 0xa1,
			0xb3, 0x70, 0x19, 0xe1, 0x40, 0x68, 0xe3, 0x05,
			0xe4, 0xb4, 0xd6, 0xd8, 0x13, 0xb6, 0xd6, 0xe1,
			0xaf, 0xf3, 0x86, 0x83, 0x49, 0x96, 0x62, 0xc1,
			0x99, 0x4b, 0x59, 0xd0, 0xa7, 0xb9, 0xb4, 0x4b,
		}),
		new(big.Int).SetBytes([]byte{
			0x96, 0x8a, 0xb5, 0x51, 0x27, 0x73, 0x49, 0x43,
			0xd4, 0x88, 0x5e, 0x03, 0x90, 0x15, 0x6f, 0xee,
			0x29, 0x2b, 0x33, 0xe9, 0x7b, 0xb3, 0x61, 0x12,
			0x17, 0x65, 0x17, 0x61, 0x98, 0xb6, 0xbd, 0xb9,
			0xac, 0xb0, 0x6f, 0x72, 0xb6, 0xbf, 0x93, 0xfe,
			0xb3, 0xb6, 0xd1, 0xef, 0xed, 0x42, 0xa9, 0x4f,
			0x65, 0xbb, 0x08, 0x9b, 0x5b, 0xf3, 0xa6, 0xa1,
			0x96, 0xac, 0x10, 0x80, 0xfa, 0xf1, 0xf9, 0x4a,
			0x5a, 0x19, 0x17, 0xd1, 0x6b, 0x15, 0x1f, 0xb2,
			0x42, 0x7d, 0x37, 0x50, 0x22, 0xc5, 0x1e, 0xf4,
			0x8b, 0xd7, 0x9e, 0x55, 0x4e, 0x90, 0x5b, 0x07,
			0x73, 0x46, 0xf5, 0xeb, 0x9b, 0x3b, 0x11, 0x9f,
			0x57, 0x80, 0x29, 0x69, 0xbe, 0x0b, 0x04, 0x0c,
			0x6c, 0x20, 0xf6, 0xb2, 0x8e, 0xdf, 0x9b, 0x97,
			0x1d, 0xce, 0x19, 0xdc, 0x0a, 0xf6, 0x6b, 0xdb,
			0xb7, 0xd9, 0xa6, 0xcb, 0x0c, 0x3c, 0x93, 0x20,
		}),
		new(big.Int).SetBytes([]byte{
			0x1d, 0x71, 0x4a, 0xa8, 0xac, 0xb3, 0x85, 0xe3,
			0x6f, 0x24, 0x99, 0x68, 0xea, 0xeb, 0xb3, 0x85,
			0xe2, 0x9b, 0x38, 0x5f, 0x34, 0x86, 0x6e, 0x3a,
			0x25, 0xd6, 0xa7, 0x94, 0xb5, 0x87, 0xff, 0xd2,
			0x14, 0x9e, 0x8e, 0xb4, 0x44, 0x81, 0xdd, 0x53,
			0xfc, 0x68, 0xf3, 0x1a, 0x06, 0x15, 0xcc, 0x44,
			0xc2, 0xc5, 0x64, 0x3b, 0x6e, 0xbb, 0xa1, 0xb9,
			0xe9, 0x68, 0x0a, 0x36, 0x3b, 0xb5, 0x95, 0x0b,
			0x6c, 0x54, 0xe0, 0xc1, 0x0f, 0x56, 0xf1, 0x20,
			0xdf, 0x0a, 0xfd, 0x06, 0x83, 0x68, 0xc4, 0xdf,
			0xb6, 0xaa, 0xc6, 0xc6, 0xa9, 0xc8, 0xa9, 0xd2,
			0x5d, 0x06, 0x17, 0x98, 0xa9, 0xce, 0x8a, 0x5e,
			0x2a, 0x21, 0x2c, 0xf2, 0xc7, 0x5c, 0xe8, 0x36,
			0xc8, 0x5a, 0xbd, 0x18, 0x98, 0x77, 0xab, 0x89,
			0x54, 0x7a, 0xc5, 0x13, 0x45, 0xd9, 0xf5, 0x8a,
			0x1c, 0xb3, 0xac, 0xf1, 0xe7, 0x78, 0xe4, 0x4a,
		}),
		new(big.Int).SetBytes([]byte{
			0xb1, 0xb3, 0xf0, 0xc3, 0xe4, 0x02, 0x8d, 0x84,
			0xb5, 0x3e, 0x5f, 0xfb, 0xd4, 0xdd, 0xe5, 0xc5,
			0x5c, 0x72, 0x8c, 0xfe, 0x7e, 0x33, 0x0c, 0x93,
			0x1b, 0x68, 0x76, 0xfb, 0xc4, 0xad, 0x7e, 0x2f,
			0xe6, 0x5b, 0xf4, 0xea, 0x3e, 0xce, 0xd9, 0x47,
			0xba, 0x72, 0x80, 0x30, 0x5e, 0x96, 0x99, 0x9b,
			0xf1, 0x07, 0x35, 0x87, 0x4a, 0x17, 0xf1, 0xee,
			0x1b, 0xa6, 0x08, 0xb4, 0x20, 0xa9, 0xcd, 0x3b,
			0x35, 0x4b, 0x4c, 0xf9, 0xad, 0x9e, 0xd0, 0xd2,
			0x7e, 0x7c, 0x90, 0x3f, 0x2b, 0x16, 0x95, 0xee,
			0xf7, 0xb3, 0xd7, 0x8a, 0xdb, 0x3f, 0x04, 0x75,
			0x96, 0xc8, 0xbf, 0xef, 0xaf, 0x06, 0x48, 0x50,
			0x55, 0x7c, 0xbe, 0xea, 0x12, 0x1d, 0x5b, 0x8c,
			0x60, 0xe2, 0x2c, 0xc7, 0x25, 0xea, 0x5d, 0x51,
			0xbc, 0xc5, 0x2e, 0x2f, 0xc2, 0x79, 0x5d, 0xfe,
			0x8e, 0x81, 0xa9, 0xf1, 0x27, 0xc7, 0x4b, 0x99,
		}),
	},
}

func TestNew(t *testing.T) {
	for i, tt := range stsTests {
		ses, err := New(bytes.NewReader(tt.iniExponent.Bytes()), tt.group, tt.generator, tt.cipher, tt.bits, tt.hash)
		if err != nil {
			t.Errorf("test %d: failed to create session: %v", i, err)
		} else if tt.iniExponent.Cmp(ses.exponent) != 0 {
			t.Errorf("test %d: exponent mismatch: have %v, want %v", i, ses.exponent, tt.iniExponent)
		}
	}
}

func TestInitiate(t *testing.T) {
	for i, tt := range stsTests {
		ses, _ := New(bytes.NewReader(tt.iniExponent.Bytes()), tt.group, tt.generator, tt.cipher, tt.bits, tt.hash)
		exp, err := ses.Initiate()
		if err != nil {
			t.Errorf("test %d: failed to initiate session: %v", i, err)
		} else if tt.iniExponential.Cmp(exp) != 0 {
			t.Errorf("test %d: exponential mismatch: have %v, want %v", i, exp, tt.iniExponential)
		}
	}
}

func TestToken(t *testing.T) {
	iniKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	accKey, _ := rsa.GenerateKey(rand.Reader, 1024)

	for i, tt := range stsTests {
		iniSes, _ := New(bytes.NewReader(tt.iniExponent.Bytes()), tt.group, tt.generator, tt.cipher, tt.bits, tt.hash)
		accSes, _ := New(bytes.NewReader(tt.accExponent.Bytes()), tt.group, tt.generator, tt.cipher, tt.bits, tt.hash)
		iniExp, _ := iniSes.Initiate()

		accExp, accToken, err := accSes.Accept(rand.Reader, accKey, iniExp)
		if err != nil {
			t.Errorf("test %d: failed to accept incoming exchange: %v", i, err)
		} else if tt.accExponential.Cmp(accExp) != 0 {
			t.Errorf("test %d: exponential mismatch: have %v, want %v", i, accExp, tt.accExponential)
		} else {
			iniToken, err := iniSes.Verify(rand.Reader, iniKey, &accKey.PublicKey, accExp, accToken)
			if err != nil {
				t.Errorf("test %d: failed to verify auth token: %v", i, err)
			} else {
				err := accSes.Finalize(&iniKey.PublicKey, iniToken)
				if err != nil {
					t.Errorf("test %d: failed to finalize key exchange: %v", i, err)
				}
			}
		}
	}
}

func TestSecret(t *testing.T) {
	iniKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	accKey, _ := rsa.GenerateKey(rand.Reader, 1024)

	for i, tt := range stsTests {
		iniSes, _ := New(bytes.NewReader(tt.iniExponent.Bytes()), tt.group, tt.generator, tt.cipher, tt.bits, tt.hash)
		accSes, _ := New(bytes.NewReader(tt.accExponent.Bytes()), tt.group, tt.generator, tt.cipher, tt.bits, tt.hash)
		iniExp, _ := iniSes.Initiate()
		accExp, accToken, _ := accSes.Accept(rand.Reader, accKey, iniExp)
		iniToken, _ := iniSes.Verify(rand.Reader, iniKey, &accKey.PublicKey, accExp, accToken)
		accSes.Finalize(&iniKey.PublicKey, iniToken)

		iniSecret, err := iniSes.Secret()
		if err != nil {
			t.Errorf("test %d: failed to retrieve initiator's secret: %v", i, err)
		} else {
			accSecret, err := accSes.Secret()
			if err != nil {
				t.Errorf("test %d: failed to retrieve acceptor's secret: %v", i, err)
			} else {
				if !bytes.Equal(iniSecret, accSecret) {
					t.Errorf("test %d: secret mismatch: initiator %v, acceptor %v", i, iniSecret, accSecret)
				}
			}
		}
	}
}
