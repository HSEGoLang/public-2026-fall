package example

//go:generate go run github.com/golang/mock/mockgen -source=example.go -destination=mock.go -package=example

type Notifier interface {
	Send(user, message string) error
}

func Welcome(n Notifier, user string) error {
	return n.Send(user, "Welcome to Go!")
}
