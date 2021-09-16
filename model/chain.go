package model

import (
	"strconv"
	"strings"
)

type Chain struct {
	Type    string
	Length  int
	Value   string
}

type ChainService interface {
	ConfirmChain(*Chain) (bool, error)
}

func NewChain(s string) *Chain{
	len,_ := strconv.Atoi(s[2:4])
	return &Chain{s[0:2],len,s[4:]}
}

type chainService struct{}

func (cs *chainService) ConfirmChain(c *Chain) (bool,error) {
	if c.Type == "NN"{
		if strings.ContainsAny(c.Value,"ABCDEFGHIJKLMNÑOPQRSTUVWXYZ"){
			return false,nil
		}

	}else if c.Type == "TX"{
		if strings.ContainsAny(c.Value,"0123456789"){
			return false,nil
		}
	}
	if c.Length != len(c.Value) {
		return false,nil
	}

	return true,nil
}

func NewChainService() (ChainService,error)  {
	return &chainService{},nil
}