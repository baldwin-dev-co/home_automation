package blueprint

import "fmt"

type Room struct {
	Occupants int
	Doorways  map[string]*Room
}

func (room *Room) Exit(to string) (int, error) {
	if room == nil {
		return 0, fmt.Errorf("Can't exit a nil room")
	}

	enteredRoom, exists := room.Doorways[to]
	if !exists {
		return room.Occupants, fmt.Errorf("This room doesn't have a door to %v", to)
	}

	room.Occupants--
	enteredRoom.Occupants++

	return room.Occupants, nil
}
