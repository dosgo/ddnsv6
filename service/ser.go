package service

import (
	"github.com/chai2010/winsvc"
	"log"
	"os"
	"path/filepath"
)


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


