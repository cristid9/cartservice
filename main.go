package main

import (
	lru "github.com/hashicorp/golang-lru"
	"log"
	"lotterProject/cart"
	"net/http"
)

type Service struct {
	cache *lru.Cache
}

type AddProductRequest struct {
	CartId          string `json:"cartId"`
	ProductId       string `json:"productId"`
	ProductQuantity uint   `json:"productQuantity"`
}

type DeleteProductRequest struct {
	CartId               string `json:"cartId"`
	ProductId            string `json:"productId"`
	TargetDeleteQuantity int    `json:"targetDeleteQuantity"`
}

type CreateCartRequest struct {
	UserId string `json:"userId"`
}

type DeleteCartRequest struct {
	CartId string `json:"cartId"`
}

type CreateCartResult struct {
	CartId string `json:"cartId"`
}

type GetCartRequest struct {
	CartId string `json:"cartId"`
}

type GetCartResult struct {
	CartId   string          `json:"cartId"`
	Products []cart.CProduct `json:"products"`
}

const (
	hostport = "localhost:8080"
)

func main() {
	cache, err := lru.New(1024)
	if err != nil {
		log.Panicf("Faild to create cache: %v", err)
	}

	svc := &Service{
		cache: cache,
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/add/", svc.addProductHandler())
	mux.Handle("/v1/delete/", svc.deleteProductHandler())
	mux.Handle("/v1/cart/", svc.getCartHandler())
	mux.Handle("/v1/create/", svc.createCartHandler())
	mux.Handle("/v1/deleteCart/", svc.createCartHandler())

	address := hostport

	log.Print("Listening on ", hostport)
	// When running on docker mac, can't listen only on localhost
	panic(http.ListenAndServe(address, mux))
}
