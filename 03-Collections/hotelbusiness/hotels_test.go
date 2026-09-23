package hotelbusiness_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"hse-go/03-Collections/hotelbusiness"
)

func TestComputeLoad(t *testing.T) {
	testCases := []struct {
		name   string
		guests []hotelbusiness.Guest
		want   []hotelbusiness.Load
	}{
		{name: "empty", guests: nil, want: nil},
		{
			name:   "one guest",
			guests: []hotelbusiness.Guest{{CheckInDate: 1, CheckOutDate: 2}},
			want: []hotelbusiness.Load{
				{StartDate: 1, GuestCount: 1},
				{StartDate: 2, GuestCount: 0},
			},
		},
		{
			name: "consecutive guests without load change",
			guests: []hotelbusiness.Guest{
				{CheckInDate: 1, CheckOutDate: 2},
				{CheckInDate: 2, CheckOutDate: 3},
			},
			want: []hotelbusiness.Load{
				{StartDate: 1, GuestCount: 1},
				{StartDate: 3, GuestCount: 0},
			},
		},
		{
			name: "overlapping",
			guests: []hotelbusiness.Guest{
				{CheckInDate: 1, CheckOutDate: 3},
				{CheckInDate: 3, CheckOutDate: 5},
				{CheckInDate: 2, CheckOutDate: 4},
			},
			want: []hotelbusiness.Load{
				{StartDate: 1, GuestCount: 1},
				{StartDate: 2, GuestCount: 2},
				{StartDate: 4, GuestCount: 1},
				{StartDate: 5, GuestCount: 0},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := hotelbusiness.ComputeLoad(testCase.guests)
			if len(testCase.want) == 0 {
				require.Empty(t, got)

				return
			}

			require.Equal(t, testCase.want, got)
		})
	}
}
