package parameters

import (
	"dayu/env"
	"dayu/icmp"
	"dayu/parse"
	"net"
	"os"
)

var myicmp icmp.Icmp

type Par struct {
	Len    int
	Func   string
	IPList []net.IP
}

func (p *Par) Init() {
	p.Len = len(os.Args) - 1
	p.Func = os.Args[1]
	p.IPList, _ = parse.IP(os.Args[2])
}

func (p *Par) Switch(env env.Env) {
	switch p.Func {
	case "icmp":
		myicmp.Check(p.IPList)
		myicmp.Switch(env)
	}
}
