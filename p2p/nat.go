package p2p

import (
	"fmt"

	"github.com/huin/goupnp/dcps/internetgateway1"
	"github.com/huin/goupnp/dcps/internetgateway2"
)

///외부IP 가져오기
func GetExternalIP_ver1() string {
	var ip string
	clients, errors, err := internetgateway1.NewWANIPConnection1Clients()

	if errors != nil {
		return ""
	}
	if err != nil {
		return ""
	}

	for _, client := range clients {
		extip, err := client.GetExternalIPAddress()

		if err != nil {
			return ""
		}

		if extip != "" {
			ip = extip
		}
	}

	return ip
}

func GetExternalIP() string {
	clients, errors, err := internetgateway1.NewWANIPConnection1Clients()

	var ip string
	if errors != nil {
		return ""
	}
	if err != nil {
		return ""
	}

	for _, client := range clients {
		extip, err := client.GetExternalIPAddress()

		if err != nil {
			return ""
		}

		if extip != "" {
			ip = extip
		}
	}

	return ip
}

///외부IP 가져오기
func GetExternalIP_ver2() string {
	var ip string
	clients, errors, err := internetgateway2.NewWANIPConnection2Clients()

	if errors != nil {
		return ""
	}
	if err != nil {
		return ""
	}

	for _, client := range clients {
		extip, err := client.GetExternalIPAddress()

		if err != nil {
			return ""
		}

		if extip != "" {
			ip = extip
		}
	}

	return ip
}

/**
 * NAT 포트 요청
 * arg1 : 현재 사용자IP
 * arg2 : 외부포트
 * arg3 : 사설포트
 * arg4 : 프로토콜 TCP OR UDP
 */
func AddPortMapping_ver1(PrivateIP string, ExternalPort uint16, InternalPort uint16, Protocol string) {
	clients, errors, err := internetgateway1.NewWANIPConnection1Clients()

	if errors != nil {
		return
	}
	if err != nil {
		return
	}

	for _, client := range clients {
		if err := client.AddPortMapping(PrivateIP, ExternalPort, Protocol, InternalPort, PrivateIP, true, "test", 0); err != nil {
			fmt.Println(err.Error())
		}

	}

}

/**
 * NAT 포트 요청
 * arg1 : 현재 사용자IP
 * arg2 : 외부포트
 * arg3 : 사설포트
 * arg4 : 프로토콜 TCP OR UDP
 */
func AddPortMapping_ver2(PrivateIP string, ExternalPort uint16, InternalPort uint16, Protocol string) {
	clients, errors, err := internetgateway2.NewWANIPConnection2Clients()

	if errors != nil {
		return
	}
	if err != nil {
		return
	}

	for _, client := range clients {
		if err := client.AddPortMapping(PrivateIP, ExternalPort, Protocol, InternalPort, PrivateIP, true, "test", 0); err != nil {
			fmt.Println(err.Error())
		}

	}

}

/**
 * NAT 포트 삭제
 * arg1 : 현재 사용자 IP
 * arg2 : 외부포트
 * arg3 : 프로토콜 TCP OR UDP
 */
func DelPortMapping_ver2(PrivateIP string, ExternalPort uint16, Protocol string) {
	clients, errors, err := internetgateway2.NewWANIPConnection2Clients()

	if errors != nil {
		return
	}
	if err != nil {
		return
	}

	for _, client := range clients {
		if err := client.DeletePortMapping(PrivateIP, ExternalPort, Protocol); err != nil {
			fmt.Println(err.Error())
		}

	}
}

/**
 * NAT 포트 삭제
 * arg1 : 현재 사용자 IP
 * arg2 : 외부포트
 * arg3 : 프로토콜 TCP OR UDP
 */
func DelPortMapping_ver1(PrivateIP string, ExternalPort uint16, Protocol string) {
	clients, errors, err := internetgateway2.NewWANIPConnection1Clients()

	if errors != nil {
		return
	}
	if err != nil {
		return
	}

	for _, client := range clients {
		if err := client.DeletePortMapping(PrivateIP, ExternalPort, "tcp"); err != nil {
			fmt.Println(err.Error())
		}

	}
}
