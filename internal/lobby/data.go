package lobby

import (
	"errors"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ GetLobbyData returns ALL Lobby Data as Bytes^
func (lob *Lobby) GetLobbyData() []byte {
	// Convert to bytes
	data, err := proto.Marshal(lob.Data)
	if err != nil {
		log.Println("Error Marshaling Lobby Data ", err)
		return nil
	}
	return data
}

// ^ Find returns Pointer to Peer.ID and Peer ^
func (lob *Lobby) Find(q string) (peer.ID, *md.Peer, error) {
	// Retreive Data
	peer := lob.Peer(q)
	id := lob.ID(q)

	if peer == nil || id == "" {
		return "", nil, errors.New("Search Error, peer was not found in map.")
	}

	return id, peer, nil
}

// ^ ID returns ONE Peer.ID in PubSub ^
func (lob *Lobby) ID(q string) peer.ID {
	// Iterate through PubSub in topic
	for _, id := range lob.ps.ListPeers(lob.Data.Code) {
		// If Found Match
		if id.String() == q {
			return id
		}
	}
	return ""
}

// ^ Peer returns ONE Peer in Lobby ^
func (lob *Lobby) Peer(q string) *md.Peer {
	// Iterate Through Peers, Return Matched Peer
	for _, peer := range lob.Data.Available {
		// If Found Match
		if peer.Id == q {
			return peer
		}
	}
	return nil
}

// ^ setBusy changes peer values in Lobby ^
func (lob *Lobby) setUnavailable(event *md.LobbyEvent) {
	// Remove Peer
	delete(lob.Data.Available, event.Id)

	// Add Peer to Unavailable Map
	lob.Data.Unavailable[event.Id] = event.Peer
	lob.Data.Size = int32(len(lob.Data.Available)) + 1 // Account for User

	// Marshal data to bytes
	bytes, err := proto.Marshal(lob.Data)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.callback(bytes)
}

// ^ removePeer deletes peer from all maps ^
func (lob *Lobby) removePeer(event *md.LobbyEvent) {
	// Remove Peer from Available
	delete(lob.Data.Available, event.Id)

	// Remove Peer from Unavailable
	delete(lob.Data.Unavailable, event.Id)
	lob.Data.Size = int32(len(lob.Data.Available)) + 1 // Account for User

	// Marshal data to bytes
	bytes, err := proto.Marshal(lob.Data)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.callback(bytes)
}

// ^ updatePeer changes peer values in Lobby ^
func (lob *Lobby) updatePeer(event *md.LobbyEvent) {
	// Remove Peer
	delete(lob.Data.Unavailable, event.Id)

	// Update Peer with new data
	id := event.Id
	lob.Data.Available[id] = event.Peer
	lob.Data.Size = int32(len(lob.Data.Available)) + 1 // Account for User

	// Marshal data to bytes
	bytes, err := proto.Marshal(lob.Data)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.callback(bytes)
}
