package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/wangtuanjie/ip17mon"
	"strings"
)

func init() {
	ip17mon.Init("17monipdb.dat")
}

func Location(ip string) (map[string]string, error) {

	mapStr := map[string]string{}

	loc, err := ip17mon.Find(ip)
	if err != nil {
		return nil, err
	}

	mapStr["country"] = loc.Country
	mapStr["Region"] = loc.Region
	mapStr["City"] = loc.City

	return mapStr, nil
}

func ipip(ip string) {
	if loc, err := ip17mon.Find(ip); err != nil {
		fmt.Fprintf(os.Stderr, "\t%s: %v\n", ip, err)
	} else {
		fmt.Println("\t"+ip+"\t", loc.Country, loc.Region, loc.City)
	}
}

func is_ipv4(ip string) bool {
	trial := net.ParseIP(ip)

	if trial == nil {
		return false
	}
	return trial.To4() != nil
}

func stdin() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ip := scanner.Text()
		ip = strings.TrimSpace(ip)
		if 0 == len(ip) {
			continue
		}

		ip_location(ip)
	}

	stdin()
}

func ip_location(ip string) {
	if false == is_ipv4(ip) {
		//如果不是IP，尝试DNS获取IP地址
		addrs, err := net.LookupHost(ip)

		if err != nil {
			fmt.Println("Not a valid hostname or ipv4 address: " + ip)
			stdin()
		}

		//google包含了ipv6的地址，过滤掉
		for _, v := range addrs {
			if true == is_ipv4(v) {
				ipip(v)
			}
		}
	} else {
		ipip(ip)
	}
}

func main() {
	if len(os.Args) > 1 {
		for k, ip := range os.Args {
			if k > 0 {
				fmt.Println(ip)
				ip_location(ip)
			}
		}
	} else {
		fmt.Println("input a ipV4 ip address or domain name:")
		stdin()
	}
}
