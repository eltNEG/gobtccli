# gobtccli

Lightweight package to call bitcoin-cli from your go app. This package has no external dependency. The package however requires that you have bitcoin-cli installed in your env path and also have a valid bitcoin.conf file.


# Sample Usage

```golang
package main

import (
    btccli "github.com/eltNEG/gobtccli"
)

const conf = "/full/path/to/bitcoin.conf"

func main() {
	blockcount, err := btccli.Call[int](conf, "getblockcount")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("blockcount =", blockcount)
}

```

# Contribution

Feel free to raise a PR.

# License

MIT
