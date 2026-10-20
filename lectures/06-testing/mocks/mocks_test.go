package mocks

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	saved := notifyUser

	t.Cleanup(func() { notifyUser = saved })

	var notifiedUser, notifiedMsg string

	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	// bytesInUse returns 980 MB in this example.
	const user = "joe@example.org"
	CheckQuota(user)

	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}

	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s",
			notifiedUser, user)
	}

	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}
