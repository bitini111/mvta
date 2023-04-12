package mvta

import "github.com/bitini111/mvta/acceptor"

func main() {
	acceptor.NewWSAcceptor(":3250")
}
