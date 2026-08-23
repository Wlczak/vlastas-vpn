package routertun

import (
	"fmt"
	"net/netip"
	"os"
	"syscall"

	"golang.zx2c4.com/wireguard/tun"
	"gvisor.dev/gvisor/pkg/buffer"
	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/header"
	"gvisor.dev/gvisor/pkg/tcpip/link/channel"
	"gvisor.dev/gvisor/pkg/tcpip/network/ipv4"
	"gvisor.dev/gvisor/pkg/tcpip/network/ipv6"
	"gvisor.dev/gvisor/pkg/tcpip/stack"
	"gvisor.dev/gvisor/pkg/tcpip/transport/icmp"
	"gvisor.dev/gvisor/pkg/tcpip/transport/tcp"
	"gvisor.dev/gvisor/pkg/tcpip/transport/udp"
)

type Interface struct {
	TUN   tun.Device
	EP    *channel.Endpoint
	NICID tcpip.NICID
}
type Router struct {
	Stack    *stack.Stack
	Up, Down *Interface
}
type netTun struct {
	ep       *channel.Endpoint
	stack    *stack.Stack
	events   chan tun.Event
	notify   *channel.NotificationHandle
	incoming chan *buffer.View
	mtu      int
	nic      tcpip.NICID
}
type endpoint struct {
	*channel.Endpoint
	tun *netTun
}

func (e *endpoint) Attach(d stack.NetworkDispatcher)             { e.tun.ep.Attach(d) }
func (e *endpoint) IsAttached() bool                             { return e.tun.ep.IsAttached() }
func (e *endpoint) MTU() uint32                                  { return e.tun.ep.MTU() }
func (e *endpoint) Capabilities() stack.LinkEndpointCapabilities { return 0 }
func (e *endpoint) MaxHeaderLength() uint16                      { return 0 }
func (e *endpoint) LinkAddress() tcpip.LinkAddress               { return "" }
func (e *endpoint) WritePacket(_ stack.RouteInfo, _ tcpip.NetworkProtocolNumber, p *stack.PacketBuffer) tcpip.Error {
	v := p.ToView()
	p.DecRef()
	e.tun.incoming <- v
	return nil
}
func (e *endpoint) WritePackets(ps stack.PacketBufferList) (int, tcpip.Error) {
	for _, p := range ps.AsSlice() {
		if err := e.WritePacket(stack.RouteInfo{}, 0, p); err != nil {
			return 0, err
		}
	}
	return ps.Len(), nil
}
func (e *endpoint) WriteRawPacket(*stack.PacketBuffer) tcpip.Error { return nil }
func (e *endpoint) ARPHardwareType() header.ARPHardwareType        { return header.ARPHardwareNone }
func (e *endpoint) AddHeader(*stack.PacketBuffer) {
}

func New() *Router {
	return &Router{Stack: stack.New(stack.Options{NetworkProtocols: []stack.NetworkProtocolFactory{ipv4.NewProtocol, ipv6.NewProtocol}, TransportProtocols: []stack.TransportProtocolFactory{tcp.NewProtocol, udp.NewProtocol, icmp.NewProtocol4, icmp.NewProtocol6}, HandleLocal: true})}
}
func (r *Router) CreateInterface(id tcpip.NICID, addr netip.Addr, prefix, mtu int) (*Interface, error) {
	if !addr.IsValid() || prefix < 0 || prefix > addr.BitLen() {
		return nil, fmt.Errorf("invalid address or prefix")
	}
	e := channel.New(1024, uint32(mtu), "")
	n := &netTun{ep: e, stack: r.Stack, events: make(chan tun.Event, 1), incoming: make(chan *buffer.View), mtu: mtu, nic: id}
	n.notify = e.AddNotify(n)
	if err := r.Stack.CreateNIC(id, &endpoint{Endpoint: e, tun: n}); err != nil {
		return nil, fmt.Errorf("CreateNIC: %v", err)
	}
	proto := ipv6.ProtocolNumber
	if addr.Is4() {
		proto = ipv4.ProtocolNumber
	}
	pa := tcpip.ProtocolAddress{Protocol: proto, AddressWithPrefix: tcpip.AddressWithPrefix{Address: tcpip.AddrFromSlice(addr.AsSlice()), PrefixLen: prefix}}
	if err := r.Stack.AddProtocolAddress(id, pa, stack.AddressProperties{}); err != nil {
		r.Stack.RemoveNIC(id)
		return nil, fmt.Errorf("AddProtocolAddress: %v", err)
	}
	if addr.Is4() {
		r.Stack.AddRoute(tcpip.Route{Destination: header.IPv4EmptySubnet, NIC: id})
	} else {
		r.Stack.AddRoute(tcpip.Route{Destination: header.IPv6EmptySubnet, NIC: id})
	}
	n.events <- tun.EventUp
	return &Interface{TUN: n, EP: e, NICID: id}, nil
}
func (r *Router) EnableForwarding() error {
	if err := r.Stack.SetForwardingDefaultAndAllNICs(ipv4.ProtocolNumber, true); err != nil {
		return fmt.Errorf("IPv4 forwarding: %v", err)
	}
	if err := r.Stack.SetForwardingDefaultAndAllNICs(ipv6.ProtocolNumber, true); err != nil {
		return fmt.Errorf("IPv6 forwarding: %v", err)
	}
	return nil
}
func (r *Router) Close()                   { r.Stack.Close() }
func (t *netTun) Name() (string, error)    { return fmt.Sprintf("router-%d", t.nic), nil }
func (t *netTun) File() *os.File           { return nil }
func (t *netTun) Events() <-chan tun.Event { return t.events }
func (t *netTun) Read(b [][]byte, s []int, o int) (int, error) {
	v, ok := <-t.incoming
	if !ok {
		return 0, os.ErrClosed
	}
	n, e := v.Read(b[0][o:])
	if e != nil {
		return 0, e
	}
	s[0] = n
	return 1, nil
}
func (t *netTun) Write(bs [][]byte, o int) (int, error) {
	for _, b := range bs {
		b = b[o:]
		if len(b) == 0 {
			continue
		}
		p := stack.NewPacketBuffer(stack.PacketBufferOptions{Payload: buffer.MakeWithData(b)})
		switch b[0] >> 4 {
		case 4:
			t.ep.InjectInbound(ipv4.ProtocolNumber, p)
		case 6:
			t.ep.InjectInbound(ipv6.ProtocolNumber, p)
		default:
			p.DecRef()
			return 0, syscall.EAFNOSUPPORT
		}
	}
	return len(bs), nil
}
func (t *netTun) WriteNotify() {
	p := t.ep.Read()
	if p == nil {
		return
	}
	v := p.ToView()
	p.DecRef()
	t.incoming <- v
}
func (t *netTun) Close() error {
	t.stack.RemoveNIC(t.nic)
	t.ep.RemoveNotify(t.notify)
	t.ep.Close()
	close(t.events)
	close(t.incoming)
	return nil
}
func (t *netTun) MTU() (int, error) { return t.mtu, nil }
func (t *netTun) BatchSize() int    { return 1 }
