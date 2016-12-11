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
			encodeAndPrint(l, *typ)
		} else {
			decodeAndPrint(l)
		}
	}
}

func encodeAndPrint(l string, t string) {
	switch t {
	case "zcash-block":
		fmt.Println(encodeZcash(l, cid.ZcashBlock))
	case "zcash-tx":
		fmt.Println(encodeZcash(l, cid.ZcashTx))
	default:
		fatal("unrecognized input type: " + t)
	}
}

func encodeZcash(l string, mcd uint64) string {
	out, err := hex.DecodeString(l)
	if err != nil {
		fatal(err)
	}

	hval := reverse(out)
	h, err := mh.Encode(hval, mh.DBL_SHA2_256)
	if err != nil {
		fatal(err)
	}

	c := cid.NewCidV1(cid.ZcashBlock, h)
	return c.String()
}

func decodeAndPrint(l string) {
	c, err := cid.Parse(l)
	if err != nil {
		fatal(err)
	}

	dec, _ := mh.Decode(c.Hash())
	cname := mcp.CodeToString(mcp.Code(c.Type()))
	raw := c.Hash()[len(c.Hash())-dec.Length:]
	switch c.Type() {
	case cid.ZcashBlock, cid.ZcashTx:
		raw = reverse(raw)
	}
	fmt.Printf("%s\t%s\t%s\n", cname, dec.Name, hex.EncodeToString(raw))
}

func reverse(b []byte) []byte {
	out := make([]byte, len(b))
	for i, v := range b {
		out[len(b)-(1+i)] = v
	}
	return out
}
