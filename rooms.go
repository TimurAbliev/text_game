package main

type Room struct {
	name        string
	objects     map[string]string
	transitions map[string]*Room
	onLock      func() string
	onApply     map[string]func(item string) string
}

func initRooms() {
	rooms = make(map[string]*Room)

	kitchen := &Room{
		name:        "кухня",
		objects:     map[string]string{"чай": "на столе"},
		transitions: make(map[string]*Room), // связывается позже
	}

	hall := &Room{
		name:        "коридор",
		objects:     map[string]string{},
		transitions: make(map[string]*Room),
	}

	room := &Room{
		name: "комната",
		objects: map[string]string{
			"ключи":     "на столе",
			"конспекты": "на столе",
			"рюкзак":    "на стуле",
		},
		transitions: make(map[string]*Room),
	}

	street := &Room{
		name:        "улица",
		objects:     map[string]string{},
		transitions: make(map[string]*Room),
	}

	// Связи
	kitchen.transitions["коридор"] = hall
	hall.transitions["кухня"] = kitchen
	hall.transitions["комната"] = room
	hall.transitions["улица"] = street
	room.transitions["коридор"] = hall
	street.transitions["домой"] = hall

	rooms["кухня"] = kitchen
	rooms["коридор"] = hall
	rooms["комната"] = room
	rooms["улица"] = street
}
