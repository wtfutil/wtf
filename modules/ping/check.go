package ping

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/wtfutil/wtf/logger"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

const (
	ProtocolICMP     = 1
	ProtocolIPv6ICMP = 58
	msgFail          = "fail"
	msgSuccess       = "success"
)

func checkIPV4(target string, pingTimeout int, logging bool) (result string) {
	dst, err := net.ResolveIPAddr("ip4", target)
	if err != nil {
		return msgFail
	}

	var conn *icmp.PacketConn
	conn, err = icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		if logging {
			logger.Log(fmt.Sprintf("%s | failed to listen for ip4:icmp packets", moduleName))
		}
		return msgFail
	}

	defer func() {
		_ = conn.Close()
	}()

	pSeq := rand.Intn(65536)
	pID := rand.Intn(65536)
	pData := []byte("wtf")

	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID:   pID,
			Seq:  pSeq,
			Data: pData,
		},
	}

	var b []byte

	b, err = m.Marshal(nil)
	if err != nil {
		return msgFail
	}

	var n int

	if logging {
		logger.Log(fmt.Sprintf("%s | pinging: %s", moduleName, dst.String()))
	}

	n, err = conn.WriteTo(b, dst)
	if err != nil || n != len(b) {
		if logging {
			logger.Log(fmt.Sprintf("%s | failed to send ping to %s", moduleName, dst.String()))
		}

		return msgFail
	}

	reply := make([]byte, 1500)

	waitStart := time.Now()
	waitDuration := time.Duration(pingTimeout) * time.Second

	for {
		if time.Since(waitStart) > waitDuration {
			if logging {
				logger.Log(fmt.Sprintf("%s | timed out waiting for: %s", moduleName, target))
			}

			return msgFail
		}

		if err = conn.SetReadDeadline(time.Now().Add(waitDuration)); err != nil {
			if logging {
				logger.Log(fmt.Sprintf("%s | failed to set read deadline: %v", moduleName, err))
			}
			return msgFail
		}

		var peer net.Addr

		n, peer, err = conn.ReadFrom(reply)
		if err != nil {
			if logging {
				logger.Log(fmt.Sprintf("%s | failed to read reply for target: %s %v", moduleName, target, err))
			}

			return msgFail
		}

		// if we received a reply from a different request then ignore
		if dst.String() != peer.String() {
			continue
		}

		// if we received a reply from the intended recipient then validate content
		var rm *icmp.Message
		rm, err = icmp.ParseMessage(ProtocolICMP, reply[:n])
		if err != nil {
			return msgFail
		}

		if rm.Type == ipv4.ICMPTypeEchoReply {
			pe := rm.Body.(*icmp.Echo)
			if pID == pe.ID && pSeq == pe.Seq && bytes.Equal(pe.Data, pData) {
				if logging {
					logger.Log(fmt.Sprintf("%s | got reply for %s", moduleName, dst.String()))
				}
				return msgSuccess
			}
			continue
		}

		break
	}

	return msgFail
}

func checkIPV6(target string, pingTimeout int, logging bool) (result string) {
	dst, err := net.ResolveIPAddr("ip6", target)
	if err != nil {
		if logging {
			logger.Log(fmt.Sprintf("%s | failed to resolve %s", moduleName, target))
		}

		return msgFail
	}

	var conn6 *icmp.PacketConn

	conn6, err = icmp.ListenPacket("ip6:ipv6-icmp", "::")
	if err != nil {
		if logging {
			logger.Log(fmt.Sprintf(" %s | failed to listen for ip6:ipv6-icmp packets", moduleName))
		}
		return msgFail
	}

	defer func() {
		_ = conn6.Close()
	}()

	pSeq := rand.Intn(65536)
	pID := rand.Intn(65536)
	pData := []byte("wtf")

	m := icmp.Message{
		Type: ipv6.ICMPTypeEchoRequest, Code: 0,
		Body: &icmp.Echo{
			ID:   pID,
			Seq:  pSeq,
			Data: pData,
		},
	}

	var b []byte

	b, err = m.Marshal(nil)
	if err != nil {
		return msgFail
	}

	var n int

	if logging {
		logger.Log(fmt.Sprintf("%s | pinging: %s", moduleName, dst.String()))
	}

	n, err = conn6.WriteTo(b, dst)
	if err != nil || n != len(b) {
		if logging {
			logger.Log(fmt.Sprintf("%s | failed to send ping to %s", moduleName, dst.String()))
		}

		return msgFail
	}

	reply := make([]byte, 1500)

	waitStart := time.Now()
	waitDuration := time.Duration(pingTimeout) * time.Second

	for {
		if time.Since(waitStart) > waitDuration {
			if logging {
				logger.Log(fmt.Sprintf("%s | timed out waiting for: %s", moduleName, target))
			}

			return msgFail
		}

		if err = conn6.SetReadDeadline(time.Now().Add(waitDuration)); err != nil {
			if logging {
				logger.Log(fmt.Sprintf("%s | failed to set read deadline: %v", moduleName, err))
			}
			return msgFail
		}

		var peer net.Addr

		n, peer, err = conn6.ReadFrom(reply)
		if err != nil {
			return msgFail
		}

		if dst.String() != peer.String() {
			continue
		}

		var rm *icmp.Message

		rm, err = icmp.ParseMessage(ProtocolIPv6ICMP, reply[:n])
		if err != nil {
			return msgFail
		}

		if rm.Type == ipv6.ICMPTypeEchoReply {
			pe := rm.Body.(*icmp.Echo)
			if pID == pe.ID && pSeq == pe.Seq && bytes.Equal(pe.Data, pData) {
				if logging {
					logger.Log(fmt.Sprintf("%s | got reply for %s", moduleName, dst.String()))
				}
				return msgSuccess
			}
			continue
		}

		if logging {
			logger.Log(fmt.Sprintf("%s | got reply for %s", moduleName, dst.String()))
		}

		break
	}

	return msgFail
}

func checkTarget(t *net.IP, pingTimeout int, logging bool) (result string) {
	if t.To4() != nil {
		result = checkIPV4(t.String(), pingTimeout, logging)
		return
	}

	result = checkIPV6(t.String(), pingTimeout, logging)

	return
}

func ipsFromTarget(t string) (ips []*net.IP, isIP bool, err error) {
	// try to parse target as an IP
	if i := net.ParseIP(t); i != nil {
		// return the target parsed as an IP
		return []*net.IP{&i}, true, nil
	}

	// try to parse target as an FQDN
	var pIPs []string

	pIPs, err = net.LookupHost(t)
	if err != nil {
		logger.Log(fmt.Sprintf("%s | lookup failed for: %s", moduleName, t))
		return
	}

	// parse each IP
	for x := 0; x < len(pIPs); x++ {
		pIP := net.ParseIP(pIPs[x])
		ips = append(ips, &pIP)
	}

	return
}

type target struct {
	raw string
	ips []*net.IP
	err error
}

func parseTargets(ts []string) (targets []target) {
	for _, t := range ts {
		res, isIP, err := ipsFromTarget(t)
		// if parsing shortens ipv6 address then use shorter version for checking and display
		if isIP && res[0].String() != t {
			t = res[0].String()
		}

		targets = append(targets, target{
			raw: t,
			ips: res,
			err: err,
		})
	}

	return
}
