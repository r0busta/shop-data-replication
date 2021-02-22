package handler

import (
	"fmt"

	"github.com/r0busta/shop-data-replication/models"
	"github.com/r0busta/shop-data-replication/storage"
)

type Handler interface {
	OnProductCreate(*models.Product) error
	OnProductUpdate(*models.Product) error
	OnProductDelete(*models.Product) error
}

type DefaultHandler struct {
	storage storage.Storage
}

func NewDefaultHandler(storage storage.Storage) Handler {
	return &DefaultHandler{
		storage: storage,
	}
}

func (h *DefaultHandler) OnProductCreate(p *models.Product) error {
	return h.storage.SaveProduct(p)

}

func (h *DefaultHandler) OnProductUpdate(p *models.Product) error {
	return h.storage.SaveProduct(p)
}

func (h *DefaultHandler) OnProductDelete(p *models.Product) error {
	return fmt.Errorf("Not implemented")
}
