package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"lotterProject/cart"
	"lotterProject/productservice"
	"net/http"
	"golang.org/x/exp/slices"
)

var (
	pService = productservice.NewProductService()
)

func (s *Service) addProductHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Expecting POST request"))
			return
		}

		request := AddProductRequest{}
		err := json.NewDecoder(io.LimitReader(r.Body, 8*1024)).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to parse request"))
			return
		}

		result, _ := s.processAddProductRequest(request)
		data, err := json.Marshal(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to marshal response"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("content-type", "application/json")
		// refactor this
		w.Write(data)
	})
}

func (s *Service) deleteProductHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Expecting DELETE request"))
			return
		}

		request := DeleteProductRequest{}
		err := json.NewDecoder(io.LimitReader(r.Body, 8*1024)).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to parse request"))
			return
		}

		result, _ := s.processDeleteProductRequest(request)
		data, err := json.Marshal(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to marshal response"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		w.Write(data)

	})
}

func (s *Service) getCartHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Expecting GET request"))
			return
		}

		request := GetCartRequest{}
		err := json.NewDecoder(io.LimitReader(r.Body, 8*1024)).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to parse request"))
			return
		}

		result, _ := s.processGetCartRequest(request)
		data, err := json.Marshal(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to marshal response"))
			return
		}

		w.WriteHeader(http.StatusFound)
		w.Header().Add("content-type", "application/json")
		w.Write(data)
	})
}


func (s *Service) deleteCartHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			fmt.Println(r.Method)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Expecting DELETE request"))
			return
		}

		request := DeleteCartRequest{}
		err := json.NewDecoder(io.LimitReader(r.Body, 8*1024)).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to parse request"))
			return
		}

		result, _ := s.processDeleteCartRequest(request)
		data, err := json.Marshal(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to marshal response"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		w.Write(data)
	})
}

func (s *Service) createCartHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println(r.Method)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Expecting POST request"))
			return
		}

		request := CreateCartRequest{}
		err := json.NewDecoder(io.LimitReader(r.Body, 8*1024)).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to parse request"))
			return
		}

		result, _ := s.processCreateCartRequest(request)
		data, err := json.Marshal(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to marshal response"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("content-type", "application/json")
		w.Write(data)
	})
}

func (s *Service) processDeleteCartRequest(request DeleteCartRequest) (bool, error) {
	_, ok := s.cache.Get(request.CartId)

	if !ok {
		return false, errors.New("cart not found")
	}

	s.cache.Remove(request.CartId)

	return true, nil
}

func (s *Service) processCreateCartRequest(request CreateCartRequest) (CreateCartResult, error) {
	cartId := uuid.NewString()

	c := cart.Cart{
		UserId: request.UserId,
		CartId: cartId,
		Products: make([]cart.CProduct, 0),
	}

	s.cache.Add(cartId, c)

	return CreateCartResult{CartId: cartId}, nil
}

func (s *Service) processGetCartRequest(request GetCartRequest) (GetCartResult, error) {
	cartEntry, ok := s.cache.Get(request.CartId)

	if !ok {
		return GetCartResult{}, errors.New("currentCart not found")
	}

	currentCart := cartEntry.(cart.Cart)

	return GetCartResult { CartId: currentCart.CartId, Products: currentCart.Products }, nil
}

func (s *Service) processAddProductRequest(request AddProductRequest) (bool, error) {
	cartEntry, ok := s.cache.Get(request.CartId)

	if !ok {
		return false, errors.New("cart not found")
	}

	currentCart := cartEntry.(cart.Cart)

	product, err := pService.GetProduct(request.ProductId)

	if err != nil { return false, err }

	cartProduct := cart.CProduct{Product: product, Quantity: int(request.ProductQuantity)}
	currentCart.Products = append(currentCart.Products, cartProduct)

	s.cache.Add(currentCart.CartId, currentCart)

	return true, nil
}

func (s *Service) processDeleteProductRequest(request DeleteProductRequest) (bool, error) {
	cartEntry, ok := s.cache.Get(request.CartId)

	if !ok {
		return false, errors.New("cart not found")
	}

	currentCart := cartEntry.(cart.Cart)

	idx := slices.IndexFunc(currentCart.Products, func(p cart.CProduct) bool { return p.Id == request.ProductId })

	if idx == -1 {
		return false, errors.New("product not in cart, invalid  delete request")
	}

	if request.TargetDeleteQuantity == currentCart.Products[idx].Quantity {
		currentCart.Products = append(currentCart.Products[:idx], currentCart.Products[idx+1:]...)
	} else {
		currentCart.Products[idx].Quantity -= 1
	}

	s.cache.Add(currentCart.CartId, currentCart)

	return true, nil
}