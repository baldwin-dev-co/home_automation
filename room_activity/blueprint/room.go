package blueprint

import "fmt"

type Room struct {
	Entrance bool
	Occupants int
	Doorways  map[string]*Room
}

type roomJSON struct {
	Entrance bool `json:"entrance"`
	Doorways []string `json:"doorways"`
}

func (room *Room) Enter(from ...string) (int, error) {
	if room == nil {
		return 0, fmt.Errorf("Can't enter a nil room")
	}

	if len(from) == 0 && !room.Entrance {
		return room.Occupants, fmt.Errorf("Non entrances must include a `from` argument")
	}

	exitedRoom, exists := room.Doorways[from[0]]
	if !exists {
		return room.Occupants, fmt.Errorf("This room doesn't have a doorway to %v", from[0])
	}

	room.Occupants++
	exitedRoom.Occupants--

	return room.Occupants, nil
}

func (room *Room) Exit(to ...string) (int, error) {
	if room == nil {
		return 0, fmt.Errorf("Can't exit a nil room")
	}

	if len(to) == 0 && !room.Entrance {
		return room.Occupants, fmt.Errorf("Non entrances must include a `to` argument")
	}

	enteredRoom, exists := room.Doorways[to[0]]
	if !exists {
		return room.Occupants, fmt.Errorf("This room doesn't have a door to %v", to[0])
	}

	room.Occupants--
	enteredRoom.Occupants++

	return room.Occupants, nil
}
