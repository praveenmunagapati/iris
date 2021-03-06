// Iris - Decentralized cloud messaging
// Copyright (c) 2013 Project Iris. All rights reserved.
//
// Community license: for open source projects and services, Iris is free to use,
// redistribute and/or modify under the terms of the GNU Affero General Public
// License as published by the Free Software Foundation, either version 3, or (at
// your option) any later version.
//
// Evaluation license: you are free to privately evaluate Iris without adhering
// to either of the community or commercial licenses for as long as you like,
// however you are not permitted to publicly release any software or service
// built on top of it without a valid license.
//
// Commercial license: for commercial and/or closed source projects and services,
// the Iris cloud messaging system may be used in accordance with the terms and
// conditions contained in an individually negotiated signed written agreement
// between you and the author(s).

package sortext

import (
	"math/big"
	"testing"
)

// Smoke tests for convenience wrappers - not comprehensive.
var idata = []*big.Int{big.NewInt(-5), big.NewInt(0), big.NewInt(11), big.NewInt(100)}
var rdata = []*big.Rat{big.NewRat(-314, 100), big.NewRat(0, 1), big.NewRat(1, 1), big.NewRat(2, 1), big.NewRat(10007, 10)}

var wrappertests = []struct {
	name   string
	result int
	i      int
}{
	{"SearchBigInts", SearchBigInts(idata, big.NewInt(11)), 2},
	{"SearchBigRats", SearchBigRats(rdata, big.NewRat(21, 10)), 4},
	{"BigIntSlice.Search", BigIntSlice(idata).Search(big.NewInt(0)), 1},
	{"BigRatSlice.Search", BigRatSlice(rdata).Search(big.NewRat(20, 10)), 3},
}

func TestSearchWrappers(t *testing.T) {
	for _, e := range wrappertests {
		if e.result != e.i {
			t.Errorf("%s: expected index %d; got %d", e.name, e.i, e.result)
		}
	}
}
