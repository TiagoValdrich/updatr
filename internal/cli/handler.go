package cli

import (
	"flag"
)

type Handler struct {
	Path *string
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ReadArguments() {
	h.Path = flag.String("path", "./", "")
	flag.Parse()
}
