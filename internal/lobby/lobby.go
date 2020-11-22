package lobby

import (
	"context"
	"fmt"
	"sync"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// Define Function Types
type OnRefreshed func(data []byte)
type OnError func(err error, method string)

// Struct to Implement Node Callback Methods
type LobbyCallback struct {
	Refreshed OnRefreshed
	Error     OnError
}

// Lobby represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Lobby.Publish, and received
// messages are pushed to the Messages channel.
type Lobby struct {
	// Public Vars
	Messages chan *pb.LobbyEvent
	Data     *pb.Lobby

	// Private Vars
	ctx    context.Context
	call   LobbyCallback
	doneCh chan struct{}
	mutex  sync.Mutex
	ps     *pubsub.PubSub
	topic  *pubsub.Topic
	self   peer.ID
	sub    *pubsub.Subscription
}

// ^ Enter Joins/Subscribes to pubsub topic, Initializes BadgerDB, and returns Lobby ^
func Enter(ctx context.Context, callback LobbyCallback, ps *pubsub.PubSub, id peer.ID, olc string) (*Lobby, error) {
	// Join the pubsub Topic
	topic, err := ps.Join(olc)
	if err != nil {
		return nil, err
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Initialize Lobby for Peers
	lobInfo := &pb.Lobby{
		Code:  olc,
		Size:  1,
		Peers: make(map[string]*pb.Peer),
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:    ctx,
		call:   callback,
		doneCh: make(chan struct{}, 1),
		ps:     ps,
		topic:  topic,
		sub:    sub,
		self:   id,

		Data:     lobInfo,
		Messages: make(chan *pb.LobbyEvent, ChatRoomBufSize),
	}

	// start reading messages
	go lob.handleMessages()
	go lob.handleEvents()
	return lob, nil
}

// ^ Info returns ALL Lobby Data as Bytes^
func (lob *Lobby) Info() []byte {
	// Convert to bytes
	data, err := proto.Marshal(lob.Data)
	if err != nil {
		fmt.Println("Error Marshaling Lobby Data ", err)
		return nil
	}
	return data
}

// ^ Find returns Pointer to Peer.ID and Peer ^
func (lob *Lobby) Find(q string) (peer.ID, *pb.Peer) {
	// Retreive Data
	peer := lob.Peer(q)
	id := lob.ID(q)

	return id, peer
}

// ^ Send publishes a message to the pubsub topic OLC ^
func (lob *Lobby) Update(peer *pb.Peer) error {
	// Create Lobby Event
	event := pb.LobbyEvent{
		Event: pb.LobbyEvent_UPDATE,
		Peer:  peer,
	}

	// Convert Event to Proto Binary
	bytes, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	// Publish to Topic
	err = lob.topic.Publish(lob.ctx, bytes)
	if err != nil {
		return err
	}
	return nil
}

// ^ End terminates lobby loop ^
func (lob *Lobby) Exit() {
	// Create Lobby Event
	event := pb.LobbyEvent{
		Event: pb.LobbyEvent_EXIT,
		Id:    lob.self.String(),
	}

	// Convert Event to Proto Binary, Suppress Error
	bytes, _ := proto.Marshal(&event)

	// Publish to Topic
	lob.topic.Publish(lob.ctx, bytes)
	lob.doneCh <- struct{}{}
}
