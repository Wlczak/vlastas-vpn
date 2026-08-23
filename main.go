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

	upConfigString := fmt.Sprintf(
		"private_key=%s\npublic_key=%s\npreshared_key=%s\nendpoint=%s\nallowed_ip=%s\npersistent_keepalive_interval=%s\n",
		upPrivateKey, upPublicKey, upPresharedKey, upEndpoint, upAllowedIP, upPersistentKeepaliveInterval,
	)

	downPrivateKey := ""
	downClientPublicKey := ""

	downConfigString := fmt.Sprintf(
		"private_key=%s\nlisten_port=51821\npublic_key=%s\nallowed_ip=10.1.0.2/32\n",
		downPrivateKey,
		downClientPublicKey,
	)
	upTun, upNet, err := netstack.CreateNetTUN(
		[]netip.Addr{
			netip.MustParseAddr("10.0.0.169"),
		},
		[]netip.Addr{
			netip.MustParseAddr("1.1.1.1"),
		},
		1420,
	)
	if err != nil {
		panic(err)
	}

	downTun, downNet, err := netstack.CreateNetTUN(
		[]netip.Addr{
			netip.MustParseAddr("10.1.0.1"),
		},
		[]netip.Addr{
			netip.MustParseAddr("1.1.1.1"),
		},
		1420,
	)
	if err != nil {
		panic(err)
	}

	upDev := device.NewDevice(
		upTun,
		conn.NewDefaultBind(),
		device.NewLogger(device.LogLevelVerbose, ""),
	)

	downDev := device.NewDevice(
		downTun,
		conn.NewDefaultBind(),
		device.NewLogger(device.LogLevelVerbose, ""),
	)

	err = upDev.IpcSet(upConfigString)
	if err != nil {
		panic(err)
	}

	if err := upDev.Up(); err != nil {
		panic(err)
	}

	if err := downDev.IpcSet(downConfigString); err != nil {
		panic(err)
	}

	if err := downDev.Up(); err != nil {
		panic(err)
	}

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
