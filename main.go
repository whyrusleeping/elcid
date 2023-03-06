package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"

	cid "github.com/ipfs/go-cid"
	mcp "github.com/multiformats/go-multicodec-packed"
	mh "github.com/multiformats/go-multihash"
)

func main() {
	typ := flag.String("type", "", "type to convert to/from cid")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatal("must specify 'e' or 'd' for encode or decode")
	}

	encode, err := parseArgs(flag.Arg(0), *typ)
	if err != nil {
		log.Fatal(err)
	}

	err = processInput(encode, *typ)
	if err != nil {
		log.Fatal(err)
	}
}

func parseArgs(arg string, typ string) (bool, error) {
	switch arg {
	case "e":
		if typ == "" {
			return false, fmt.Errorf("please specify type of input for encoding")
		}
		return true, nil
	case "d":
		return false, nil
	default:
		return false, fmt.Errorf("must specify 'e' or 'd' for encode or decode")
	}
}

func processInput(encode bool, typ string) error {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		l := scan.Text()
		if encode {
			out, err := encodeToCid(l, typ)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(out)
		} else {
			t, h, v, err := decodeToInfo(l)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s\t%s\t%s\n", t, h, v)
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func encodeToCid(l string, t string) (string, error) {
	switch t {
	case "zcash-block":
		return encode(l, cid.ZcashBlock, mh.DBL_SHA2_256)
	case "zcash-tx":
		return encode(l, cid.ZcashTx, mh.DBL_SHA2_256)
	case "bitcoin-block":
		return encode(l, cid.BitcoinBlock, mh.DBL_SHA2_256)
	case "bitcoin-tx":
		return encode(l, cid.BitcoinTx, mh.DBL_SHA2_256)
	case "eth-block":
		return encode(l, cid.EthBlock, mh.KECCAK_256)
	case "eth-tx":
		return encode(l, cid.EthTx, mh.KECCAK_256)
	default:
		return "", fmt.Errorf("unrecognized input type: %s", t)
	}
}

func encode(l string, mcd uint64, mhType uint64) (string, error) {
	out, err := hex.DecodeString(l)
	if err != nil {
		return "", err
	}

	h, err := mh.Encode(out, mh.SHA2_256)
	if err != nil {
		return "", err
	}

	return cid.NewCidV1(mcd, mh.Multihash(h)).String(), nil
}

func decodeToInfo(l string) (string, string, string, error) {
	c, err := cid.Parse(l)
	if err != nil {
		return "", "", "", err
	}

	dec, _ := mh.Decode(c.Hash())
	cname := mcp.CodeToString(mcp.Code(c.Type()))
	raw := c.Hash()[len(c.Hash())-dec.Length:]
	switch c.Type() {
	case cid.ZcashBlock, cid.ZcashTx, cid.BitcoinBlock, cid.BitcoinTx:
		raw = reverse(raw)
	}

	return cname, dec.Name, hex.EncodeToString(raw), nil
}

func reverse(b []byte) []byte {
	out := make([]byte, len(b))
	for i, v := range b {
		out[len(b)-(1+i)] = v
	}
	return out
}
