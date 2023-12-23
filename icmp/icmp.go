package icmp

import (
	"bytes"
	"dayu/env"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"sync"
)

type Icmp struct {
	Test        string
	IPList      []net.IP
	AliveIPList []net.IP
}

func (i *Icmp) Check(ipList []net.IP) {
	i.Test = "ping"
	i.IPList = ipList
}

func (i *Icmp) Switch(env env.Env) {
	switch i.Test {
	case "ping":
		i.pingSwitch(env.OS)
	case "icmp1":
	case "icmp2":
	}
}

func (i *Icmp) pingSwitch(OS string) {
	switch OS {
	case "linux":
		var wg sync.WaitGroup
		for _, ip := range i.IPList {
			wg.Add(1)
			go func(ip net.IP) {
				defer wg.Done()
				i.linuxPing(ip)
			}(ip)
		}
		wg.Wait()
	case "windows":
		var wg sync.WaitGroup
		for _, ip := range i.IPList {
			wg.Add(1)
			go func(ip net.IP) {
				defer wg.Done()
				i.windowsPing(ip)
			}(ip)
		}
		wg.Wait()
	}
}

func (i *Icmp) linuxPing(ip net.IP) {
	var cmd *exec.Cmd
	cmd = exec.Command("/bin/bash", "-c", "ping -c 1 -w 1 "+ip.String()+" && echo true || echo false")
	cmdout := bytes.Buffer{}
	cmd.Stdout = &cmdout
	_ = cmd.Start()
	_ = cmd.Wait()
	if strings.Contains(cmdout.String(), "true") {
		fmt.Println(ip.String())
		i.AliveIPList = append(i.AliveIPList, ip)
	}
}

func (i *Icmp) windowsPing(ip net.IP) {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "ping -n 1 -w 1 "+ip.String()+" && echo true || echo false")
	cmdout := bytes.Buffer{}
	cmd.Stdout = &cmdout
	_ = cmd.Start()
	_ = cmd.Wait()
	if strings.Contains(cmdout.String(), "true") {
		fmt.Println(ip.String())
		i.AliveIPList = append(i.AliveIPList, ip)
	}
}
