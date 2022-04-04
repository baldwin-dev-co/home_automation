package blueprint

import "fmt"

type Room struct {
	Entrance  bool
	Occupants int
	Doorways  map[string]*Room
}

func MakeRoom() Room {
	return Room{
		Doorways: make(map[string]*Room),
	}
}

type RoomJSON struct {
	Entrance bool     `json:"entrance"`
	Doorways []string `json:"doorways"`
}

func (room *Room) Enter(from ...string) error {
	if room == nil {
		return fmt.Errorf("Can't enter a nil room")
	}

	if len(from) == 0 && !room.Entrance {
		return fmt.Errorf("Non entrances must include a `from` argument")
	}

	room.Occupants++

	if len(from) > 0 && len(from[0]) > 0 {
		exitedRoom, exists := room.Doorways[from[0]]
		if !exists {
			return fmt.Errorf("This room doesn't have a doorway to %v", from[0])
		}

		exitedRoom.Occupants--
	}

	return nil
}

func (room *Room) Exit(to ...string) error {
	if room == nil {
		return fmt.Errorf("Can't exit a nil room")
	}

	if len(to) == 0 && !room.Entrance {
		return fmt.Errorf("Non entrances must include a `to` argument")
	}

	room.Occupants--
	
	if len(to) > 0 && len(to[0]) > 0 {
		enteredRoom, exists := room.Doorways[to[0]]
		if !exists {
			return fmt.Errorf("This room doesn't have a door to %v", to[0])
		}
		
		enteredRoom.Occupants++
	}

	return nil
}
