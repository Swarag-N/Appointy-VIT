package meeting

import (
	"appointy/api/user"
)

// TimeValidator takes Current User and To Respond Meeting
// Validates the Users free time, by
// Make a query to to retrive the list of Meetings in User Future Meeting
// Filter to get "yes" meeting
func TimeValidator(u *user.User, m *Meeting) {
	// Logic when either start_time or end_time (Timestamps) of new meeting
	// lie between the prerecorded meetings return false

}
