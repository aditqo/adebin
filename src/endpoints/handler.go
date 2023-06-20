package endpoints

import "adebin/store"

type Handler struct {
	Store store.Client
}