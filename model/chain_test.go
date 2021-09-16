package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {
		s, err := ParseString([]byte(testData.Input))
		// acá agregar chequeos propios del test por ejemplo:
		r:= NewChain(s) //Creo una nueva cadena
		var chainServ ChainService
		chainServ, err = NewChainService()
		b,err := chainServ.ConfirmChain(r)
		assert.Equal(t, b, testData.Success, "Cadena no aceptada")
		assert.NotEqual(nil, err, "Hay un error")

	}

}

func ParseString(bytes []byte) (string, error) { //devuelve primer resultado
      myString := string(bytes[:])
	  return myString,nil
}



