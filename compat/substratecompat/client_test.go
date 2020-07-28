package substratecompat_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/renproject/multichain"
	"github.com/renproject/multichain/compat/substratecompat"
	"github.com/renproject/pack"
)

var _ = Describe("Substrate client", func() {
	Context("when verifying burns", func() {
		It("should verify a valid burn", func() {
			// client, err := substratecompat.NewClient("http://localhost:9933")
			client, err := substratecompat.NewClient("ws://localhost:9944")
			// client, err := substratecompat.NewClient("wss://testnet-node-1.acala.laminar.one/ws")
			if err != nil {
				panic(err)
			}

			nonce := [32]byte{0}
			amount, to, confs, err := client.BurnEvent(context.Background(), multichain.BTC, pack.NewBytes32(nonce), pack.NewU64(3047))
			if err != nil {
				panic(err)
			}

			fmt.Printf("Amount: %v\n", amount)
			fmt.Printf("To: %v\n", to)
			fmt.Printf("Confs: %v\n", confs)

			Expect(amount).Should(Equal(6000))
		})
	})
})
