
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
