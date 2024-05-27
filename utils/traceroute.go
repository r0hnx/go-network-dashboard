package utils

import (
	"os/exec"
	"strconv"
	"strings"
)

type TracerouteHop struct {
	TTL  int    `json:"ttl"`
	IP   string `json:"ip"`
	RTT  string `json:"rtt"`
	AS   string `json:"as,omitempty"`
	Host string `json:"host,omitempty"`
}

func RunTraceroute(target string) ([]TracerouteHop, error) {
	cmd := exec.Command("traceroute", target)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var hops []TracerouteHop

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}

		ttl, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}

		rtts := fields[2:]
		var ip, rtt, as, host string
		if len(fields) > 3 && fields[1] != "*" {
			ip = fields[1]
			rtt = strings.Join(rtts, " ")
			if len(fields) > 4 {
				as = fields[3]
				host = fields[4]
			}
		} else {
			ip = fields[1]
			rtt = "Timeout"
		}

		hop := TracerouteHop{
			TTL:  ttl,
			IP:   ip,
			RTT:  rtt,
			AS:   as,
			Host: host,
		}
		hops = append(hops, hop)
	}

	return hops, nil
}
