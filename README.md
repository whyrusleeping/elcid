elcid
=============
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

A tool for encoding and decoding content IDs

![](https://jeanarogers.files.wordpress.com/2011/05/el-cid.jpg)

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

```sh
go get github.com/whyrusleeping/elcid
```

## Usage

To use Elcid, you can either encode or decode CIDs.

####Encoding

To encode content into a CID, use the following command:

```shell
echo "<content>" | elcid -type<input_type> e
```
The available input types are:

* zcash-block
* zcash-tx
* bitcoin-block
* bitcoin-tx
* eth-block
* eth-tx

####Decoding

To decode a CID, use the following command:

```shell
echo "<cid>" | elcid d
```
The output will contain the codec name, multihash type, and raw hash of the CID.

####Examples

```shell
me@home ~> echo 000000002f39739a41ee7ceb910d6b2d0084129719e95aa16f7af005380e370d | elcid -type=zcash-block e
z4QJh97xKVfG43tBcdLk31E1An5HD574Xp7fYxi7Mp1vXNGnxvj

me@home ~> echo zdpuAnXTsoojgUhScFpvU8vyMpdXswNCMUVB28gzy4NJtvUZc | elcid d
dag-cbor        sha2-256        1f426f7fdd365555ab7fea80e2f97f9293d9139834a2ba9efe63730e545da707
```
## Maintainers

[@whyrusleeping](https://github.com/whyrusleeping)

## Contribute

PRs are welcome!

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Jeromy Johnson
