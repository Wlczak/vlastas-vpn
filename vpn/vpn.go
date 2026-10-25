package vpn

import (
	"fmt"
	"net/netip"

	"github.com/Wlczak/vlastas-vpn/helpers"
	"github.com/Wlczak/vlastas-vpn/routertun"
	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
)

type VpnConfig struct {
	UpPrivateKey                  string
	UpPeerPublicKey               string
	UpEndpoint                    string
	UpAllowedIP                   string
	UpPersistentKeepaliveInterval string
	UpInterfaceIp                 string
	DownPrivateKey                string
	DownClientPublicKey           string
}

type VpnGateway struct {
	router     *routertun.Router
	upDevice   *device.Device
	downDevice *device.Device
}

func StartVpn(config VpnConfig) (*VpnGateway, error) {
	upPrivateKey := helpers.Base64tohex(config.UpPrivateKey)
	upPeerPublicKey := helpers.Base64tohex(config.UpPeerPublicKey)
	upEndpoint := config.UpEndpoint
	upAllowedIP := config.UpAllowedIP
	upPersistentKeepaliveInterval := config.UpPersistentKeepaliveInterval
	upInterfaceIp := config.UpInterfaceIp

	upConfigString := fmt.Sprintf(
		"private_key=%s\npublic_key=%s\nendpoint=%s\nallowed_ip=%s\npersistent_keepalive_interval=%s\n",
		upPrivateKey, upPeerPublicKey, upEndpoint, upAllowedIP, upPersistentKeepaliveInterval,
	)

	downPrivateKey := helpers.Base64tohex(config.DownPrivateKey)
	downClientPublicKey := helpers.Base64tohex(config.DownClientPublicKey)

	downConfigString := fmt.Sprintf(
		"private_key=%s\nlisten_port=51821\npublic_key=%s\nallowed_ip=10.1.0.2/32\n",
		downPrivateKey,
		downClientPublicKey,
	)
	router := routertun.New()
	up, err := router.CreateInterface(1, netip.MustParseAddr(upInterfaceIp), 32, 1420, true)
	if err != nil {
		return nil, err
	}

	down, err := router.CreateInterface(2, netip.MustParseAddr("10.1.0.1"), 24, 1420, false)
	if err != nil {
		return nil, err
	}

	if err := router.EnableForwarding(); err != nil {
		return nil, err
	}
	if err := router.EnableNAT(netip.MustParsePrefix("10.1.0.0/24"), netip.MustParseAddr(upInterfaceIp)); err != nil {
		return nil, err
	}
	router.Up, router.Down = up, down
	upDev := device.NewDevice(
		up.TUN,
		conn.NewDefaultBind(),
		device.NewLogger(device.LogLevelVerbose, ""),
	)

	downDev := device.NewDevice(
		down.TUN,
		conn.NewDefaultBind(),
		device.NewLogger(device.LogLevelVerbose, ""),
	)

	err = upDev.IpcSet(upConfigString)
	if err != nil {
		return nil, err
	}

	if err := upDev.Up(); err != nil {
		return nil, err
	}

	if err := downDev.IpcSet(downConfigString); err != nil {
		return nil, err
	}

	if err := downDev.Up(); err != nil {
		return nil, err
	}
	return &VpnGateway{router: router, upDevice: upDev, downDevice: downDev}, nil
}
