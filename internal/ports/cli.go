package ports

import (
	"fmt"
)

type CLI struct{}

func NewCLI() *CLI {
	return &CLI{}
}

func (cli *CLI) Run(args []string) {
	fmt.Println("CLI started with args:", args)
}
