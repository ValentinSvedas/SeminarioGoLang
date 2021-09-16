package main

import (
	"bufio"
	"fmt"
	"github.com/valentinsvedas/tpGoLang/model"
	"os"
)

func main(){
	aux := false

	for aux == false {
		var c bool
		var i int
		var chainServ model.ChainService
		chainServ, err := model.NewChainService()
		fmt.Print("Ingrese cadena para comprobar: ")
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')
		r := model.NewChain(entrada)
		c, err = chainServ.ConfirmChain(r)
		fmt.Print(c)
		if err != nil {
			panic(err)
		}
		fmt.Print("Para comprobar otra cadena ingrese 0: ")
		fmt.Scanf("%d\n", &i)
		if i != 0 {
			aux=true
		}
	}
}
