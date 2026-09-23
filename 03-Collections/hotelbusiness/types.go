//go:build !change

package hotelbusiness

// Guest describes one hotel stay. CheckOutDate is not part of the stay.
type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

// Load is the hotel occupancy starting from StartDate.
type Load struct {
	StartDate  int
	GuestCount int
}
