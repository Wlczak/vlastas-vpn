package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/netip"

	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun/netstack"
)

func main() {
	privateKey := base64tohex("")
	publicKey := base64tohex("")
	endpoint := ""
	presharedKey := base64tohex("")
	allowedIP := ""
	persistentKeepaliveInterval := "25"

	configString := fmt.Sprintf(
		"private_key=%s\npublic_key=%s\npreshared_key=%s\nendpoint=%s\nallowed_ip=%s\npersistent_keepalive_interval=%s\n",
		privateKey, publicKey, presharedKey, endpoint, allowedIP, persistentKeepaliveInterval,
	)
	tun, _, err := netstack.CreateNetTUN(
		[]netip.Addr{netip.MustParseAddr("10.0.0.169")}, // tunnel-internal IP
		[]netip.Addr{netip.MustParseAddr("1.1.1.1")},    // DNS, if needed
		1420, // MTU
	)
	if err != nil {
		panic(err)
	}

	dev := device.NewDevice(tun, conn.NewDefaultBind(), device.NewLogger(device.LogLevelVerbose, ""))
	err = dev.IpcSet(configString) // your peer/key config, same UAPI format as before
	if err != nil {
		panic(err)
	}
	dev.Up()

	for true {

	}
}

func base64tohex(base64String string) string {
	base64Bytes := make([]byte, base64.StdEncoding.DecodedLen(len(base64String)))
	n, err := base64.StdEncoding.Decode(base64Bytes, []byte(base64String))
	if err != nil {
		panic(err)
	}
	base64Bytes = base64Bytes[:n]

	hex := hex.EncodeToString(base64Bytes)

	return hex
}
