package blueprint_test

import (
	"encoding/json"
	"fmt"
	"room_activity/blueprint"
	"testing"
)

func BlueprintEqualsJSON(blueprint blueprint.Blueprint, json blueprint.BlueprintJSON) error {
	if len(blueprint) != len(json) {
		return fmt.Errorf("len(blueprint) != len(json)")
	}

	for roomId, roomJSON := range json {
		room, exists := blueprint[roomId]
		if !exists {
			return fmt.Errorf("%v exists in json but not in blueprint", roomId)
		}

		if err := RoomEqualsJson(*room, roomJSON); err != nil {
			return err
		}
	}

	return nil
}

func Test_Blueprint_UnmarshalJSON(t *testing.T) {
	blueprintJSON := blueprint.BlueprintJSON{
		"entrance": blueprint.RoomJSON{
			Entrance: true,
			Doorways: []string{
				"living-room",
				"hallway",
				"office",
				"kitchen",
			},
		},
		"living-room": blueprint.RoomJSON{
			Entrance: true,
			Doorways: []string{
				"patio",
				"entrance",
				"kitchen",
			},
		},
		"hallway": blueprint.RoomJSON{
			Doorways: []string{
				"entrance",
				"bathroom",
				"master-bedroom",
				"bedroom-1",
				"bedroom-2",
				"bedroom-3",
			},
		},
		"bathroom": blueprint.RoomJSON{
			Doorways: []string{
				"hallway",
			},
		},
		"master-bedroom": blueprint.RoomJSON{
			Doorways: []string{
				"hallway",
				"master-bathroom",
				"patio",
			},
		},
		"master-bathroom": blueprint.RoomJSON{
			Doorways: []string{
				"master-bedroom",
			},
		},
		"bedroom-1": blueprint.RoomJSON{
			Doorways: []string{
				"hallway",
			},
		},
		"bedroom-2": blueprint.RoomJSON{
			Doorways: []string{
				"hallway",
			},
		},
		"bedroom-3": blueprint.RoomJSON{
			Doorways: []string{
				"hallway",
			},
		},
		"patio": blueprint.RoomJSON{
			Entrance: true,
			Doorways: []string{
				"kitchen",
				"living-room",
				"master-bedroom",
			},
		},
		"kitchen": blueprint.RoomJSON{
			Entrance: true,
			Doorways: []string{
				"patio",
				"entrance",
				"living-room",
				"half-bath",
			},
		},
		"office": blueprint.RoomJSON{
			Doorways: []string{
				"entrance",
			},
		},
		"half-bath": blueprint.RoomJSON{
			Doorways: []string{
				"kitchen",
			},
		},
	}

	data, err := json.Marshal(blueprintJSON)
	if err != nil {
		t.Fatalf("Error marshaling test json: %v", err)
	}

	blueprint := blueprint.MakeBlueprint()
	err = json.Unmarshal(data, &blueprint)
	if err != nil {
		t.Errorf("Error unmarshaling data into blueprint: %v", err)
	}

	if err := BlueprintEqualsJSON(blueprint, blueprintJSON); err != nil {
		t.Errorf("blueprint != blueprintJSON: %v", err)
	}
}
