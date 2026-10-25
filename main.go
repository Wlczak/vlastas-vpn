package main

import (
	"log"
	"os"

	"github.com/Wlczak/vlastas-vpn/vpn"
	"github.com/joho/godotenv"
)

func main() {
	// r := SetupRouter()

	// go func() {
	// 	err := r.Run()
	// 	h.PanicChckErr(err)
	// }()

	// state.SetStateDefaults()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	config := vpn.VpnConfig{
		UpPrivateKey:                  os.Getenv("UP_PRIVATE_KEY"),
		UpPeerPublicKey:               os.Getenv("UP_PEER_PUBLIC_KEY"),
		UpEndpoint:                    os.Getenv("UP_ENDPOINT"),
		UpAllowedIP:                   os.Getenv("UP_ALLOWED_IP"),
		UpPersistentKeepaliveInterval: os.Getenv("UP_PERSISTENT_KEEPALIVE_INTERVAL"),
		UpInterfaceIp:                 os.Getenv("UP_INTERFACE_IP"),
		DownPrivateKey:                os.Getenv("DOWN_PRIVATE_KEY"),
		DownClientPublicKey:           os.Getenv("DOWN_CLIENT_PUBLIC_KEY"),
	}

	if _, err := vpn.StartVpn(config); err != nil {
		log.Fatal(err)
	}
	select {}
}
