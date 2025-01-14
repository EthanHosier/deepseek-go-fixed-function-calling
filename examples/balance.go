package deepseek_examples

import (
	"context"
	"log"
	"os"

	deepseek "github.com/EthanHosier/deepseek-go-fixed-function-calling"
)

func Balance() {
	client := deepseek.NewClient(os.Getenv("DEEPSEEK_KEY"))
	ctx := context.Background()
	balance, err := deepseek.GetBalance(client, ctx)
	if err != nil {
		log.Fatalf("Error getting balance: %v", err)
	}

	if balance == nil {
		log.Fatalf("Balance is nil")
	}

	if len(balance.BalanceInfos) == 0 {
		log.Fatalf("No balance information returned")
	}
	log.Printf("%+v\n", balance)
}
