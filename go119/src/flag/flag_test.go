package flag_test

import (
	"flag"
	"math/big"
	"net"
	"testing"
	"time"
)

func TestTextVar(t *testing.T) {
	var ipaddr net.IP
	flag.TextVar(&ipaddr, "ipaddr", net.IPv4(192, 168, 0, 100), "ip address")
	t.Log(ipaddr.String()) // 192.168.0.100

	var st time.Time
	flag.TextVar(&st, "time", time.Now(), "time now")
	t.Log(st.Format("2006-01-02 15:04:05")) // 2022-08-xx xx:xx:xx

	var bi big.Int
	flag.TextVar(&bi, "big", big.NewInt(1000), "big int")
	t.Log(bi.String()) // 1000
}
