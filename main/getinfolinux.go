package main

import (
	"bytes"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"net"
	"os/exec"
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	infos, err := cpu.Info()
	if err != nil {
		fmt.Printf("error:%v\n")
		return
	}
	for _, info := range infos {
		fmt.Printf("id:%v, length(%v)\n", info.VendorID, len(infos))
	}
	fmt.Printf("cpuinfo:%v,length(%v)\n", infos[0], len(infos))

	n1, err := cpu.Counts(true)
	n2, err := cpu.Counts(false)
	fmt.Printf("logical:%v\n", n1)
	fmt.Printf("physical:%v\n", n2)
	/////////////////////
	var outInfo bytes.Buffer
	cmd := exec.Command("ifconfig")

	// 设置接收
	cmd.Stdout = &outInfo

	// 执行
	cmd.Run()

	//fmt.Println(outInfo.String())
	///////////////////////
	//baseNicPath := "/sys/class/net/"
	//cmd = exec.Command("ls", baseNicPath)
	//buf, err := cmd.Output()
	//if err != nil {
	//	//fmt.Println("Error:", err)
	//	return
	//}
	//output := string(buf)
	//
	//for _, device := range strings.Split(output, "\n") {
	//	if len(device) > 1 {
	//		fmt.Println(device)
	//		ethHandle, err := ethtool.NewEthtool()
	//		if err != nil {
	//			panic(err.Error())
	//		}
	//		defer ethHandle.Close()
	//		stats, err := ethHandle.LinkState(device)
	//		if err != nil {
	//			panic(err.Error())
	//		}
	//		fmt.Printf("LinkName: %s LinkState: %d\n", device, stats)
	//	}
	//
	//}
	/////
	fmt.Printf("mac addrs: %q\n", getMacAddrs())
	fmt.Printf("ips: %q\n", getIPs())
}

func getMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		fmt.Printf("ether name:%v\n", netInterface.Name)
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		if netInterface.Name[0] != 'e' {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

func getIPs() (ips []string) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
