package main

import (
	"gomodules.xyz/zoom-lib-golang"
	"log"
)

func main() {
	user, err := zoom.GetUser(zoom.GetUserOpts{EmailOrID: "tamal@appscode.com"})
	if err != nil {
		log.Fatalf("failed to get zoom user: %v", err)
	}

	// Set HostID
	createMeetingOpt := zoom.CreateMeetingOptions{
		HostID: user.ID,
	}

	meeting, err := zoom.CreateMeeting(createMeetingOpt)

	if err != nil {
		log.Printf("Error: %+v\n\n", err)
	}
	log.Printf("Meeting: %+v\n\n", meeting)
}
