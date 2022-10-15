Overview
=========

The solution is fairly simple, we expose 5 APIs, for 5 different actions:
* creating a cart -> the API user needs to specify the id of the user creating this cart
* retrieving the cart -> only specify the cart id
* adding a product -> specify the product id and the quantity
* deleting a product -> specify which product from the cart to delete and how many times (since,
  for example, we might have 5 bottles of Pepsi in our cart, and we wish to delete only 3)
* deleting a cart -> simply specify the cart ID and if it is valid it will remove that cart from the cache

For simplicity, we store the current cart in a cache (lru cache) and  we have a mock product service with some hardcoded 
products.



APIs
======

Create Cart
------------

```shell
curl -XPOST 'http://localhost:8080/v1/create/' -d @requests/createCart.json
```

Add product
-----------

Make sure to change the `cartId` value in `requests/addProduct.json `. 

```shell
curl -XPOST 'http://localhost:8080/v1/add/' -d @requests/addProduct.json 
```

Get Cart
---------
Make sure to change the `cartId` value in `requests/addProduct.json`.

```shell
curl -XPOST 'http://localhost:8080/v1/add/' -d @requests/addProduct.json 
```

Delete Product
---------------

Make sure to change the `cartId` value in `request/deleteProduct.json`.

```shell
curl -XDELETE 'http://localhoost::8080/v1/delete/' -d @request/deleteProduct.json
```

Delete cart
-----------

Make sure to change the `cartId` value in `requests/deleteCart.json`.

```shell
curl -XDELETE 'http://localhost:8080/v1/deleteCart/' -d @requests/deleteCart.json
```