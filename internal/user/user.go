package user

import (
	"google.golang.org/protobuf/proto"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
)

// @ Constant Variables
const K_SONR_USER_PATH = "user.snr"
const K_SONR_PRIV_KEY = "snr-peer.privkey"

// @ Sonr User Struct
type User struct {
	// Properties
	Call md.NodeCallback
	FS   *sf.FileSystem

	// User Data
	contact  *md.Contact
	device   *md.Device
	devices  []*md.Device
	protoRef *md.User
}

// ^ Method Initializes User Info Struct ^ //
func NewUser(cr *md.ConnectionRequest, callback md.NodeCallback) (*User, error) {
	// @ Init FileSystem
	fs, err := sf.NewFs(cr, callback)
	if err != nil {
		return nil, err
	}

	// @ Create Devices
	devices := make([]*md.Device, 32)
	devices = append(devices, cr.Device)

	// @ Return
	return &User{
		Call:    callback,
		contact: cr.GetContact(),
		device:  cr.Device,
		devices: devices,
		FS:      fs,
	}, nil
}

// ^ Method Loads User Data from Disk ^ //
func (u *User) LoadUser() (*md.User, error) {
	// Read File
	dat, err := u.FS.ReadFile(K_SONR_USER_PATH)
	if err != nil {
		return nil, err
	}

	// Get User Data
	user := &md.User{}
	err = proto.Unmarshal(dat, user)
	if err != nil {
		return nil, err
	}

	// Set and Return
	u.protoRef = user
	return user, nil
}
