// +build !windows

package service

import (
	"fmt"
	"ddnsv6/param"
)

func StartWin(){
	fmt.Printf("cmd:%s\r\n",param.Cmd)
}