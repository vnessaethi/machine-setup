package main

import (
	"fmt"
	"runtime"
	"machine-setup/pkg"
)

func main()  {
	if runtime.GOOS != "linux" {
        fmt.Println("You can't execute this configuration on a", runtime.GOOS, "machine!")
    } else {
		fmt.Println("Initialize machine setup on", runtime.GOOS)		
		setup.InstallVSCode()
		setup.VSCodeExtensions()
		setup.InstallConfigureTmux()
		setup.InstallKeybase()
		setup.InstallWormhole()
		setup.InstallSlack()
		setup.InstallAWSCli() 
	}
}
