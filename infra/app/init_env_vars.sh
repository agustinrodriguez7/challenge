export INSERT_PRODUCT_TO_CART_QUERY='insert into products_carts (productid, cartid) values ( %+v , (select carts.id from carts, clients where clients.id = carts.clientId and clients.id = %+v ))'
export SELECT_PRODUCTS_FROM_PRODUCTS_CARTS_QUERY='select products.id, products.categoryId, products.label, products.type, products.downloadUrl, products.weight from products_carts, carts, clients, products where clients.id = carts.clientid and products_carts.cartid = carts.id and products_carts.productid = products.id and clients.id = %+v'
export SELECT_PRODUCTS_FROM_PRODUCTS_ORDERS_QUERY='select products.id, products.categoryId, products.label, products.type, products.downloadUrl, products.weight from products_orders, orders, clients, products where clients.id = orders.clientid and products_orders.orderid = orders.id and products_orders.productid = products.id and clients.id = %+v and orders.id = %+v'
export INSERT_ORDERS_QUERY='INSERT INTO ORDERS (clientid) values (%+v)'
export SELECT_ORDER_ID_QUERY='SELECT id FROM orders WHERE orders.clientid = %+v ORDER BY id desc'
export SELECT_PRODUCTS_ID_FROM_PRODUCTS_CARTS_QUERY='SELECT products.id FROM products_carts, carts, clients, products WHERE carts.clientid = clients.id AND products_carts.cartid = carts.id and products_carts.productId = products.id and clients.id = %+v'
export INSERT_PRODUCTS_ORDERS_QUERY='INSERT INTO products_orders (productId, orderId) VALUES (%+v, %+v)'
export DELETE_PRODUCTS_CARTS_QUERY='delete from products_carts where products_carts.cartid = (select id from carts where carts.id = %+v)'