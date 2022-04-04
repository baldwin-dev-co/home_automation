package blueprint_test

import (
	"fmt"
	"room_activity/blueprint"
)

func RoomEqualsJson(room blueprint.Room, json blueprint.RoomJSON) error {
	if room.Entrance != json.Entrance {
		return fmt.Errorf("room.Entrance != json.Entrance")
	}

	if len(room.Doorways) != len(json.Doorways) {
		return fmt.Errorf("len(room.Doorways) != len(json.Doorways)")
	}

	for _, roomId := range json.Doorways {
		if _, exists := room.Doorways[roomId]; !exists {
			return fmt.Errorf("%v exists in json.Doorways but not room.Doorways", roomId)
		}
	}

	return nil
}