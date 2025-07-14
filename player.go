package main

import "strings"

type Player struct {
	location   *Room
	inventory  map[string]bool
	wearingBag bool
	doorIsOpen bool
}

func (p *Player) Look() string {
	room := p.location
	switch room.name {
	case "кухня":
		if p.wearingBag {
			return "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"
		}
		return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
	case "комната":
		list := []string{}
		for item := range room.objects {
			list = append(list, item)
		}
		if len(list) == 0 {
			return "пустая комната. можно пройти - коридор"
		}

		s := "на столе:"
		for item, place := range room.objects {
			if place == "на столе" {
				s += " " + item + ","
			}
		}
		s = strings.TrimSuffix(s, ",")
		if strings.Contains(s, ":") && !strings.Contains(s, " ") {
			s = ""
		}
		stul := ""
		for item, place := range room.objects {
			if place == "на стуле" {
				stul += item
			}
		}
		if s != "" && stul != "" {
			return s + ". на стуле: " + stul + ". можно пройти - коридор"
		} else if s != "" {
			return s + ". можно пройти - коридор"
		} else if stul != "" {
			return "на стуле: " + stul + ". можно пройти - коридор"
		}
	}
	if room.name == "коридор" {
		return "ничего интересного. можно пройти - кухня, комната, улица"
	}
	if room.name == "улица" {
		return "на улице весна. можно пройти - домой"
	}
	return "неизвестная локация"
}

func (p *Player) Move(args []string) string {
	if len(args) < 1 {
		return "куда идти?"
	}
	dest := args[0]
	if p.location.name == "улица" && dest == "улица" {
		return "уже на улице"
	}
	if p.location.name == "коридор" && dest == "улица" && !p.doorIsOpen {
		return "дверь закрыта"
	}
	if next, ok := p.location.transitions[dest]; ok {
		p.location = next
		if next.name == "кухня" {
			return "кухня, ничего интересного. можно пройти - коридор"
		}
		if next.name == "комната" {
			return "ты в своей комнате. можно пройти - коридор"
		}
		if next.name == "улица" {
			return "на улице весна. можно пройти - домой"
		}
		if next.name == "коридор" {
			return "ничего интересного. можно пройти - кухня, комната, улица"
		}
	}
	return "нет пути в " + dest
}

func (p *Player) Take(args []string) string {
	if len(args) < 1 {
		return "что взять?"
	}
	item := args[0]
	room := p.location
	if _, ok := room.objects[item]; !ok {
		return "нет такого"
	}
	if !p.wearingBag {
		return "некуда класть"
	}
	delete(room.objects, item)
	p.inventory[item] = true
	return "предмет добавлен в инвентарь: " + item
}

func (p *Player) Wear(args []string) string {
	if len(args) < 1 {
		return "что надеть?"
	}
	if args[0] == "рюкзак" {
		if place, ok := p.location.objects["рюкзак"]; ok && place == "на стуле" {
			delete(p.location.objects, "рюкзак")
			p.wearingBag = true
			return "вы надели: рюкзак"
		}
	}
	return "нет такого"
}

func (p *Player) Apply(args []string) string {
	if len(args) < 2 {
		return "не хватает аргументов"
	}
	item := args[0]
	target := args[1]

	if !p.inventory[item] {
		return "нет предмета в инвентаре - " + item
	}

	if item == "ключи" && target == "дверь" {
		p.doorIsOpen = true
		return "дверь открыта"
	}

	if target == "шкаф" {
		return "не к чему применить"
	}

	return "неизвестная команда"
}
