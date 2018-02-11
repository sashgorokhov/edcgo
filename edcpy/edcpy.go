package edcpy

import "edcgo/edjournal"

type Api struct {
	AccessToken string
}

func (a *Api) push(entry *edjournal.Entry) error {
	return nil
}