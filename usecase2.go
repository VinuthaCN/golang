package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	NspEventList := &NSPEventList{}
	NspEventList.addEvent("Sankranthi", "23rd Jan", "Giridharan", []string{"Shruthi", "Sahana", "Giridharan"}, "Giridharan")
	NspEventList.addEvent("Eid Party", "16th April", "Umesan", []string{"Umesan", "Sunny", "Deepak"}, "Umesan")

	NspEventList.Display()

	NspEventList.removeEvent("Eid Party")

	fmt.Println("After Removing")

	NspEventList.Display()
}

func (list *NSPEventList) Display() {
	current := list.Head
	for current != nil {
		b, err := json.Marshal(current)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))

		//fmt.Println(current)

		current = current.NextEvent
	}
}

func (EventList *NSPEventList) addEvent(EventName string, EventDate string, PrimaryContact string, OrganizingTeamList []string, OrganizingTeamContact string) {
	newNode := &NSPEvent{
		EventName:      EventName,
		EventDate:      EventDate,
		PrimaryContact: PrimaryContact,
		OrganizingTeam: OrganizingTeam{OrganizingTeamMembers: OrganizingTeamList, PrimaryContact: OrganizingTeamContact},
	}

	if EventList.Head == nil {
		EventList.Head = newNode
		return
	}

	curr := EventList.Head
	for curr.NextEvent != nil {
		curr = curr.NextEvent
	}

	curr.NextEvent = newNode
}

func (NspEventList *NSPEventList) removeEvent(EventName string) {
	if NspEventList.Head == nil {
		return
	}

	if NspEventList.Head.EventName == EventName {
		NspEventList.Head = NspEventList.Head.NextEvent
		return
	}

	curr := NspEventList.Head
	for curr.NextEvent != nil && curr.NextEvent.EventName != EventName {
		curr = curr.NextEvent
	}

	if curr.NextEvent != nil {
		curr.NextEvent = curr.NextEvent.NextEvent
	}
}

type NSPEventList struct {
	Head *NSPEvent
}

type NSPEvent struct {
	EventDate      string         `json:"Event Date"`
	EventName      string         `json:"Event Name"`
	PrimaryContact string         `json:"Primary Contact"`
	OrganizingTeam OrganizingTeam `json:"Organizing Team Detail"`
	NextEvent      *NSPEvent      `json:"_"`
}

type OrganizingTeam struct {
	OrganizingTeamMembers []string `json:"Organizing Team Members"`
	PrimaryContact        string   `json:"Organizing Team Contact"`
}
