package main

import (
	"testing"
)

type testCase struct {
	Native string
	Type   string
	Hash   string
	Cid    string
}

var cases = []testCase{
	{
		Native: "5a8e677603c409ceb2ee08f53fb530a2a5292455e6d2d52df38de524ac9943ee",
		Type:   "bitcoin-tx",
		Hash:   "dbl-sha2-256",
		Cid:    "z4HhYA9dptHPcigLztHGy4n8Ue2CstZWGVkrvGidtQyS2BGHv33",
	},
}

func TestVerifyCases(t *testing.T) {
	for _, c := range cases {
		cidval, err := encodeToCid(c.Native, c.Type)
		if err != nil {
			t.Fatal(err)
		}

		if cidval != c.Cid {
			t.Fatalf("error on testcase %s, cids did not match", c)
		}

		k, h, v, err := decodeToInfo(cidval)
		if err != nil {
			t.Fatal(err)
		}

		if k != c.Type {
			t.Fatalf("output type was wrong: %s != %s", k, c.Type)
		}

		if h != c.Hash {
			t.Fatalf("output hash was wrong: %s", c)
		}

		if v != c.Native {
			t.Fatalf("output native value was wrong: %s", c)
		}
	}
}
