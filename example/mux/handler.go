package main

import (
	"sync"

	"github.com/tidwall/redcon"
)

type Handler struct {
	itemsMux sync.RWMutex
	items    map[string][]byte
}

func NewHandler() *Handler { _ = "STUB: not implemented"; return nil }

func (h *Handler) detach(conn redcon.Conn, cmd redcon.Command) { _ = "STUB: not implemented"; return }

func (h *Handler) ping(conn redcon.Conn, cmd redcon.Command) { _ = "STUB: not implemented"; return }

func (h *Handler) quit(conn redcon.Conn, cmd redcon.Command) { _ = "STUB: not implemented"; return }

func (h *Handler) set(conn redcon.Conn, cmd redcon.Command) { _ = "STUB: not implemented"; return }

func (h *Handler) get(conn redcon.Conn, cmd redcon.Command) { _ = "STUB: not implemented"; return }

func (h *Handler) setnx(conn redcon.Conn, cmd redcon.Command) { _ = "STUB: not implemented"; return }

func (h *Handler) delete(conn redcon.Conn, cmd redcon.Command) { _ = "STUB: not implemented"; return }
