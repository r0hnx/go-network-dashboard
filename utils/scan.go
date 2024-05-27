package utils

import (
	"encoding/xml"
	"errors"
	"os/exec"
)

type NmapResult struct {
	Host struct {
		Ports struct {
			Port []struct {
				PortID int `xml:"portid,attr"`
				State  struct {
					State string `xml:"state,attr"`
				} `xml:"state"`
				Service struct {
					Name string `xml:"name,attr"`
				} `xml:"service"`
			} `xml:"port"`
		} `xml:"ports"`
	} `xml:"host"`
}

func RunNmap(target string) (NmapResult, error) {
	cmd := exec.Command("nmap", "-oX", "-", target)
	output, err := cmd.Output()
	if err != nil {
		return NmapResult{}, err
	}

	var result NmapResult
	if err := xml.Unmarshal(output, &result); err != nil {
		return NmapResult{}, errors.New("failed to parse nmap output")
	}

	return result, nil
}
