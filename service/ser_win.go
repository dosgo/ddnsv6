// +build windows

package service

import (
	"ddnsv6/param"
	"fmt"
	"github.com/chai2010/winsvc"
	"log"
	"os"
	"path/filepath"
)

func StartWin(){
	if !winsvc.IsAnInteractiveSession() {
		if err := winsvc.RunAsService("ddnsv6", StartServer, StopServer, false); err != nil {
			fmt.Printf("service run err:%s\r\n",err.Error())
		}
		return
	}else{
		if(len(param.Cmd)>0){
			switch (param.Cmd) {
			case "install":
				err:=Install();
				if(err==nil){
					fmt.Printf("install success\r\n")
				}else{
					fmt.Printf("install fail err:"+err.Error()+"\r\n")
				}
				break;
			case "uninstall":
				err:=UnInstall();
				if(err==nil){
					fmt.Printf("uninstall success\r\n")
				}else{
					fmt.Printf("uninstall fail err:"+err.Error()+"\r\n")
				}
				break;
			}
		}
		return ;
	}
}

/*install server*/
func Install() error{
	// change to current dir
	var appPath string
	var err error
	if appPath, err = winsvc.GetAppPath(); err != nil {
		log.Fatal(err)
	}
	if err := os.Chdir(filepath.Dir(appPath)); err != nil {
		log.Fatal(err)
	}
	//参数带上
	err=  winsvc.InstallService(appPath, "ddnsv6", "ddns  server",os.Args[1:]...);
	if(err!=nil){
		return err;
	}
	return winsvc.StartService("ddnsv6")
}

func UnInstall() error{
	// change to current dir
	return  winsvc.RemoveService("ddnsv6");
}


