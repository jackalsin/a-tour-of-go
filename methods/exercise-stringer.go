package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
//func (ip IPAddr) String() string {
//	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
//}

// this won't work; it seems the indirection only works for struct
func (ip *IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type Type1 struct {
	val string
}

func (t *Type1) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", t.val, t.val, t.val, t.val)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	fmt.Println(&Type1{"val Here"})
}
