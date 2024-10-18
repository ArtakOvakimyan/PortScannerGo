package cmd

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
)

const portsNumber = 65535
const regexpPattern = `^(\d+)(?::(\d+))?$`

func scanPorts(ip, protocol, targetPorts string) error {
	portNums, err := parsePortsArg(targetPorts)
	if err != nil {
		return err
	}

	if len(portNums) == 2 {
		_ = checkPorts(ip, protocol, portNums[0], portNums[1])
	} else if len(portNums) == 1 {
		_ = checkPorts(ip, protocol, portNums[0], portNums[0])
	}
	return nil
}

func parsePortsArg(portsArg string) ([]int, error) {
	re, _ := regexp.Compile(regexpPattern)
	matches := re.FindStringSubmatch(portsArg)
	if len(matches) == 0 {
		return nil, fmt.Errorf("Неправильный формат ввода данных : %v", portsArg)
	}
	groupsCount := countGroups(matches)
	portNums := make([]int, groupsCount)
	for i := 1; i <= groupsCount; i++ {
		number, err := strconv.Atoi(matches[i])
		if err != nil {
			return nil, fmt.Errorf("Введите целое число: %v", matches[i])
		} else if !validatePortNum(number) {
			return nil, fmt.Errorf("Несуществующий номер порта: %d", number)
		}
		portNums[i-1] = number
	}
	return portNums, nil
}

func checkPorts(ip string, protocol string, startPort, endPort int) error {
	for port := startPort; port <= endPort; port++ {
		conn, err := net.Dial(protocol, fmt.Sprintf("%s:%d", ip, port))
		if err == nil {
			defer conn.Close()
			fmt.Printf("Порт %d (%s) - открыт\n", port, protocol)
		} else {
			fmt.Printf("Порт %d (%s) - закрыт\n", port, protocol)
		}
	}
	return nil
}

func validatePortNum(portNum int) bool {
	return 0 <= portNum && portNum <= portsNumber
}
