package logicinterface

import "time"

type IAuthInformation interface {
	GetUserName() string

	GetPasswordHash() []byte

	GetCreateOn() time.Time
}
