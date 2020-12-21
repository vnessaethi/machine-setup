package setup

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func InstallVSCode() {
	fmt.Println("VSCode download packages...")

	signingKey := "wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > packages.microsoft.gpg ;"
	addGpgKey := "sudo install -o root -g root -m 644 packages.microsoft.gpg /etc/apt/trusted.gpg.d/ ;"
	addSourceList := "echo 'deb [arch=amd64,arm64,armhf signed-by=/etc/apt/trusted.gpg.d/packages.microsoft.gpg] https://packages.microsoft.com/repos/code stable main' > /etc/apt/sources.list.d/vscode.list"

	//To-Do - Validate if the program already installed
	commandDownloadDep := exec.Command("sudo", "/bin/sh", "-c", signingKey, addGpgKey, addSourceList)

	depOut, err := commandDownloadDep.CombinedOutput()
	if err != nil {
		log.Println("Its not possible download VSCode dependencies! Please check the error: ", err)
	}

	log.Println(string(depOut))

	//To-Do - Validate if the program already installed
	fmt.Println("Installing vscode via apt command...")
	commandInstallVscode := exec.Command("sudo", "/bin/sh", "-c", "apt-get install apt-transport-https ; apt-get update ; apt-get install code")

	installOut, err := commandInstallVscode.CombinedOutput()
	if err != nil {
		log.Println("Its not possible install VSCode! Please check the error: ", err)
	}

	log.Println(string(installOut))
}

func VSCodePlugins() {
	pluginList := []string{
		"4ops.terraform", "abusaidm.html-snippets", "aleksandra.go-group-imports", "amazonwebservices.aws-toolkit-vscode", "ansenhuang.vscode-view-readme", "dhdersch.vscode-terraform0-12", "eamodio.gitlens", "ecmel.vscode-html-css", "erd0s.terraform-autocomplete", "geyao.html-snippets", "golang.go", "hashicorp.terraform", "joelalejandro.nrql-language", "jvandyke.vscode-circleci", "kuromoka.circleci-status", "mohd-akram.vscode-html-format", "mohsen1.prettify-json", "ms-azuretools.vscode-docker", "ms-kubernetes-tools.vscode-kubernetes-tools", "ms-python.python", "ms-toolsai.jupyter", "ms-vscode-remote.remote-containers", "msyrus.go-doc", "p1c2u.docker-compose", "pecigonzalo.vscode-terraform-syntax", "redhat.vscode-yaml", "sidthesloth.html5-boilerplate", "skip1.go-swagger", "tomoki1207.pdf",
	}

	for _, plugin := range pluginList {
		fmt.Println("Installing", plugin, "vscode plugin...")
		installPlugins := exec.Command("/bin/sh", "-c", "code -r --install-extension", plugin, "--force")
		
		installOut, err := installPlugins.CombinedOutput()
		if err != nil {
			log.Println("Its not possible to install vscode plugins! Please check the error: ", err)
		}

		log.Println(string(installOut))
	}
}

func InstallConfigureTmux() {
	//To-Do - Validate if the program already installed
	fmt.Println("Installing Tmux...")
	commandInstallTmux := exec.Command("sudo", "/bin/sh", "-c", "apt-get install tmux")

	installOut, err := commandInstallTmux.CombinedOutput()
	if err != nil {
		log.Println("Its not possible to install tmux! Please check the error: ", err)
	}

	log.Println(string(installOut))

	//To-Do - Validate if the program and file already installed
	configureTmux := exec.Command("sudo", "/bin/sh", "-c", "cp .tmux.conf $HOME/.tmux.conf ; apt-get install xclip")

	configureOut, err := configureTmux.CombinedOutput()
	if err != nil {
		log.Println("Its not possible to configure tmux.conf! Please check the error: ", err)
	}

	log.Println(string(configureOut))
}

func InstallKeybase() {
	//To-Do - Validate if the program already installed
	fmt.Println("Installing Keybase...")

	if _, err := os.Stat("./keybase_amd64.deb"); os.IsNotExist(err) {
		installKeybase := exec.Command("sudo", "/bin/sh", "-c", "curl --remote-name https://prerelease.keybase.io/keybase_amd64.deb ; sudo apt install ./keybase_amd64.deb")

		installOut, err := installKeybase.CombinedOutput()
		if err != nil {
			log.Println("Its not possible to install keybase! Please check the error: ", err)
		}
		log.Println(string(installOut))

	} else {
		fmt.Println("Keybase already installed!")
	}
}

func InstallWormhole() {
	//To-Do - Validate if the program already installed
	fmt.Println("Installing wormhole...")

	installWormhole := exec.Command("sudo", "/bin/sh", "-c", "apt-get install magic-wormhole")

	installOut, err := installWormhole.CombinedOutput()
	if err != nil {
		log.Println("Its not possible to install wormhole! Please check the error: ", err)
	}

	log.Println(string(installOut))
}

func InstallSlack()  {
	//To-Do - Validate if the program already installed
	fmt.Println("Installing slack...")

	if _, err := os.Stat("./slack-desktop-4.12.0-amd64.deb"); os.IsNotExist(err) {
		fmt.Println("Downloading and installing slack...")
		InstallSlack := exec.Command("sudo", "/bin/sh", "-c", "wget https://downloads.slack-edge.com/linux_releases/slack-desktop-4.12.0-amd64.deb ; dpkg -i slack-desktop-4.12.0-amd64.deb")

		installOut, err := InstallSlack.CombinedOutput()
		if err != nil {
			log.Println("Its not possible to install Slack! Please check the error: ", err)
		}
		log.Println(string(installOut))
		
	} else {
		fmt.Println("Slack already installed!")
	}
}
