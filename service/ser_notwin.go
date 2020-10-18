// +build !windows

package service

import (
	"fmt"
)

func StartWin(cmd string){
	fmt.Printf("cmd:%s\r\n",cmd)
}