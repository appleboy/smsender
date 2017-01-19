package store

import "github.com/minchao/smsender/smsender/model"

type StoreResult struct {
	Data interface{}
	Err  error
}

type StoreChannel chan StoreResult

type Store interface {
	Route() RouteStore
	Message() MessageStore
}

type RouteStore interface {
	GetAll() StoreChannel
	SaveAll(routes []*model.Route) StoreChannel
}

type MessageStore interface {
	Get(id string) StoreChannel
	GetByIds(ids []string) StoreChannel
	Save(message *model.MessageRecord) StoreChannel
	Update(message *model.MessageRecord) StoreChannel
	UpdateReceipt(receipt *model.MessageReceipt) StoreChannel
}
