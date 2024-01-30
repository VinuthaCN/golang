package main

import "fmt"

func main() {

	event1 := NSPEvent{
		EventName:      "Sankranthi",
		EventDate:      "23rd Jan",
		PrimaryContact: "Giridharan",
		OrganizingTeam: OrganizingTeam{OrganizingTeamMembers: []string{"Shruthi", "Sahana", "Giridharan"}, PrimaryContact: "Giridharan"},
	}

	fmt.Println("Event1 Details ", event1)
}

type NSPEvent struct {
	EventDate      string
	EventName      string
	PrimaryContact string
	OrganizingTeam OrganizingTeam
}

type OrganizingTeam struct {
	OrganizingTeamMembers []string
	PrimaryContact        string
}
