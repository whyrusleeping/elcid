package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	cid "github.com/ipfs/go-cid"
	mcp "github.com/multiformats/go-multicodec-packed"
	mh "github.com/multiformats/go-multihash"
)

func fatal(i interface{}) {
	fmt.Println(i)
	os.Exit(1)
}

func main() {
	typ := flag.String("type", "", "type to convert to/from cid")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fatal("must specify 'e' or 'd' for encode or decode")
	}

	var encode bool
	switch flag.Arg(0) {
	case "e":
		encode = true
		if *typ == "" {
			fatal("please specify type of input for encoding")
		}
	case "d":
		encode = false
	default:
		fatal("must specify 'e' or 'd' for encode or decode")
	}

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		l := scan.Text()
		if encode {
			out, err := encodeToCid(l, *typ)
			if err != nil {
				fatal(err)
			}
			fmt.Println(out)
		} else {
			t, h, v, err := decodeToInfo(l)
			if err != nil {
				fatal(err)
			}

			fmt.Printf("%s\t%s\t%s\n", t, h, v)
		}
	}
}

func encodeToCid(l string, t string) (string, error) {
	switch t {
	case "zcash-block":
		return encodeBtc(l, cid.ZcashBlock), nil
	case "zcash-tx":
		return encodeBtc(l, cid.ZcashTx), nil
	case "bitcoin-block":
		return encodeBtc(l, cid.BitcoinBlock), nil
	case "bitcoin-tx":
		return encodeBtc(l, cid.BitcoinTx), nil
	case "eth-block":
		return encodeEth(l, cid.EthBlock)
	case "eth-tx":
		return encodeEth(l, cid.EthTx)
	default:
		return "", fmt.Errorf("unrecognized input type: %s", t)
	}
}

func encodeEth(l string, mcd uint64) (string, error) {
	out, err := hex.DecodeString(l)
	if err != nil {
		return "", err
	}

	h, err := mh.Encode(out, mh.KECCAK_256)
	if err != nil {
		return "", err
	}

	return cid.NewCidV1(mcd, mh.Multihash(h)).String(), nil
}

func encodeBtc(l string, mcd uint64) string {
	out, err := hex.DecodeString(l)
	if err != nil {
		fatal(err)
	}

	hval := reverse(out)
	h, err := mh.Encode(hval, mh.DBL_SHA2_256)
	if err != nil {
		fatal(err)
	}

	c := cid.NewCidV1(mcd, h)
	return c.String()
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
