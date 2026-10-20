package example

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestWelcome(t *testing.T) {
	ctrl := gomock.NewController(t)
	n := NewMockNotifier(ctrl)

	n.EXPECT().Send("alice", "Welcome to Go!").Return(nil)

	if err := Welcome(n, "alice"); err != nil {
		t.Fatal(err)
	}
}

func TestWelcomeSendError(t *testing.T) {
	ctrl := gomock.NewController(t)
	n := NewMockNotifier(ctrl)
	want := errors.New("notification service unavailable")

	n.EXPECT().Send("alice", "Welcome to Go!").Return(want)

	if err := Welcome(n, "alice"); !errors.Is(err, want) {
		t.Errorf("Welcome error = %v, want %v", err, want)
	}
}
