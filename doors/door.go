package door

import "errors"

var (
	ErrLockedDoor = errors.New("Door is locked")
	ErrOpenDoor   = errors.New("Door is already open")
	ErrClosedDoor = errors.New("Door is already closed")
	ErrInvalidKey = errors.New("Invalid Key")
	ErrWrongKey   = errors.New("Wrong Key")
)

type Door struct {
	key      string
	isLocked bool
	isClosed bool
}

func NewDoor(key string) (*Door, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}

	return &Door{key: key}, nil
}

func (door *Door) Open() error {
	if door.isLocked {
		return ErrLockedDoor
	}

	if !door.isClosed {
		return ErrOpenDoor
	}

	door.isClosed = false
	return nil
}

func (door *Door) Close() error {
	if door.isClosed {
		return ErrClosedDoor
	}

	door.isClosed = true
	return nil
}

func (door *Door) Lock() {
	door.isLocked = true
}

func (door *Door) Unlock(key string) error {
	if door.key != key {
		return ErrWrongKey
	}

	door.isLocked = false
	return nil
}

func (door Door) IsOpen() bool {
	return !door.isClosed
}

func (door Door) IsClosed() bool {
	return door.isClosed
}

func (door Door) IsLocked() bool {
	return door.isLocked
}

func (door Door) IsUnlocked() bool {
	return !door.isLocked
}
