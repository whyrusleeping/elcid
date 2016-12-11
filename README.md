elcid
==================

> A tool for encoding and decoding content IDs


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

```go
me@home ~> echo zdpuAnXTsoojgUhScFpvU8vyMpdXswNCMUVB28gzy4NJtvUZc | elcid d
dag-cbor        sha2-256        1f426f7fdd365555ab7fea80e2f97f9293d9139834a2ba9efe63730e545da707

me@home ~> echo 000000002f39739a41ee7ceb910d6b2d0084129719e95aa16f7af005380e370d | elcid -type=zcash-block e
z4QJh97xKVfG43tBcdLk31E1An5HD574Xp7fYxi7Mp1vXNGnxvj
```

## Contribute

PRs are welcome!

Small note: If editing the Readme, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Jeromy Johnson
