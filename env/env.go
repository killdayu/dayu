package env

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"runtime"
	"strings"
)

type Env struct {
	UserName string
	HostName string
	OS       string
	Arch     string
	Pid      int
	Process  string
	OSVer    string
}

func (e *Env) Set() {
	u, _ := user.Current()
	e.UserName = u.Name

	e.HostName, _ = os.Hostname()
	e.OS = runtime.GOOS
	e.Arch = runtime.GOARCH
	e.Pid = os.Getpid()
	e.Process = path.Base(os.Args[0])

	var cmd *exec.Cmd
	switch e.OS {
	case "linux":
		cmd = exec.Command("uname", "-a")
	case "darwin":
		cmd = exec.Command("uname", "-a")
	case "windows":
		cmd = exec.Command("cmd", "/c", "ver")
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	e.OSVer = strings.TrimSpace(string(output))

}

func (e *Env) Show() {
	fmt.Println(e)
}
