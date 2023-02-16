package notificator

import "context"

type Notificator interface {
	Send(ctx context.Context, data interface{}) error
}

type Notificatorer struct {
	Notificator Notificator
}

func NewNotificatorer(notificatorer Notificator) *Notificatorer {
	return &Notificatorer{
		Notificator: notificatorer,
	}
}
