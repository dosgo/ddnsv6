// +build windows

package service

import (
	"ddnsv6/param"
	"fmt"
	"github.com/chai2010/winsvc"
	"os"
	"path/filepath"
)

func StartWin() bool{
	if !winsvc.IsAnInteractiveSession() {
		if err := winsvc.RunAsService("ddnsv6", StartServer, StopServer, false); err != nil {
			fmt.Printf("service run err:%s\r\n",err.Error())
		}
		return true
	}else{
		if len(param.Cmd)>0 {
			switch param.Cmd {
			case "install":
				err:=Install();
				if(err==nil){
					fmt.Printf("install success\r\n")
				}else{
					fmt.Printf("install fail err:"+err.Error()+"\r\n")
				}
				return true
			case "uninstall":
				err:=UnInstall();
				if err==nil {
					fmt.Printf("uninstall success\r\n")
				}else{
					fmt.Printf("uninstall fail err:"+err.Error()+"\r\n")
				}
				return true
			}
		}
	}
	// run as normal
	StartServer()
	return false;
}

/*install server*/
func Install() error{
	// change to current dir
	var appPath string
	var err error
	if appPath, err = winsvc.GetAppPath(); err != nil {
		fmt.Printf("err:%v\r\n",err)
	}
	if err := os.Chdir(filepath.Dir(appPath)); err != nil {
		fmt.Printf("err:%v\r\n",err)
	}
	//参数带上
	err=  winsvc.InstallService(appPath, "ddnsv6", "ddns  server",os.Args[1:]...);
	if err!=nil {
		return err;
	}
	return winsvc.StartService("ddnsv6")
}

func UnInstall() error{
	// change to current dir
	return  winsvc.RemoveService("ddnsv6");
}


