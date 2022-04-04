package blueprint

import (
	"encoding/json"
	"fmt"
)

type Blueprint map[string]*Room

func MakeBlueprint() Blueprint {
	return Blueprint(make(map[string]*Room, 0))
}

type blueprintJSON map[string]roomJSON

func (blueprint *Blueprint) UnmarshalJSON(data []byte) error {
	if blueprint == nil || *blueprint == nil {
		return fmt.Errorf("Can't unmarhsal into a nil pointer")
	}

	var aux blueprintJSON
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	for k, v := range aux {
		(*blueprint)[k] = &Room{Entrance: v.Entrance}
	}

	for k, room := range aux {
		for _, to := range room.Doorways {
			(*blueprint)[k].Doorways[to] = (*blueprint)[to]
		}
	}

	return nil
}
