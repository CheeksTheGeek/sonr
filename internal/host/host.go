package host

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	secio "github.com/libp2p/go-libp2p-secio"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
)

// ^ NewHost: Creates a host with: (MDNS, TCP, QUIC on UDP) ^
func NewHost(ctx context.Context, olc string) (host.Host, error) {
	// @1. Find IPv4 Address
	point := "/sonr/dht/" + olc
	ip, ipErr := localIP()
	if ipErr != nil {
		ip = "0.0.0.0"
	}

	// @2. Create Libp2p Host
	h, err := libp2p.New(ctx,
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ip),
			"/ip6/::/tcp/0",

			fmt.Sprintf("/ip4/%s/udp/0/quic", ip),
			"/ip6/::/udp/0/quic"),

		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),

		// support secio connections
		libp2p.Security(secio.ID, secio.New),

		// support QUIC
		libp2p.Transport(libp2pquic.NewTransport),

		// support any other default transports (TCP)
		libp2p.DefaultTransports,

		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),

		// Let this host use the DHT to find other hosts
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create New IDHT
			idht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}

			// Connect to bootstrap nodes
			var wg sync.WaitGroup
			for _, peerAddr := range dht.DefaultBootstrapPeers {
				peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
				wg.Add(1)
				go func() {
					defer wg.Done()
					// We ignore errors as some bootstrap peers may be down
					h.Connect(ctx, *peerinfo) //nolint
				}()

			}
			wg.Wait()

			// We use a rendezvous point "meet me here" to announce our location.
			// This is like telling your friends to meet you at the Eiffel Tower.
			routingDiscovery := discovery.NewRoutingDiscovery(idht)
			discovery.Advertise(ctx, routingDiscovery, point)
			go handleKademliaDiscovery(ctx, h, routingDiscovery, point)
			return idht, err
		}),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableAutoRelay(),
		libp2p.EnableNATService(),
	)
	if err != nil {
		log.Fatalln("Error starting node: ", err)
	}

	// setup local mDNS discovery
	err = startMDNS(ctx, h, olc)
	fmt.Println("MDNS Started")
	return h, err
}

// ^ Handles Peers that appear on DHT ^
func handleKademliaDiscovery(ctx context.Context, h host.Host, disc *discovery.RoutingDiscovery, point string) {
	// Timer checks to dispose of peers
	peerChan, err := disc.FindPeers(ctx, point, discovery.Limit(15))
	if err != nil {
		log.Println("Failed to get DHT Peer Channel: ", err)
		return
	}

	// Start Routing Discovery
	for {
		select {
		case peer := <-peerChan:
			var wg sync.WaitGroup
			if peer.ID == h.ID() {
				continue
			} else {
				wg.Add(1)
				go func() {
					defer wg.Done()
					// We ignore errors as some bootstrap peers may be down
					h.Connect(ctx, peer) //nolint
				}()
			}
			wg.Wait()
		case <-ctx.Done():
			return
		}
	}
}
