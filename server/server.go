package server

import "wkv/storage"

var (
	DB storage.Storage
)

func init() {
	DB = storage.New()
}
