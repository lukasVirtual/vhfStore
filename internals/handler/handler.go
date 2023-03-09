package handler

import application "www.github.com/vhfStore/internals/usecases/Application"

type Handler struct {
	applicationRepository application.ApplicationInterface
}

func New(as application.ApplicationInterface) *Handler {
	return &Handler{
		applicationRepository: as,
	}
}
