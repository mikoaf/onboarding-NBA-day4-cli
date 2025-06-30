package utils

import (
	"fmt"
	"strconv"

	"go.bug.st/serial/enumerator"
)

func FindPortByVIDPID(vid, pid string) (portName string, err error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		err = fmt.Errorf("failed to enumerate ports: %v", err)
		return
	}

	for _, port := range ports {
		if port.IsUSB && port.VID == vid && port.PID == pid {
			portName = port.Name
			return
		}
	}
	err = fmt.Errorf("no port found with VID:PID %s:%s", vid, pid)
	return
}

func binaryToHex(binary string) (string, error) {
	// Convert binary string to an integer
	value, err := strconv.ParseUint(binary, 2, 16)
	if err != nil {
		return "", err
	}
	// Convert integer to a hexadecimal string
	return fmt.Sprintf("%04X", value), nil
}

func bytesToInt(b []byte) int {
	result := 0
	for _, v := range b {
		result = (result << 8) | int(v)
	}
	return result
}

func isEnable(value []byte, pos int) int {
	if pos < 8 {
		return int(int(value[0]>>pos) & 1)
	} else {
		return int(int(value[1]>>pos) & 1)
	}
}
