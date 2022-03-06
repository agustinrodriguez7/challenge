CREATE TABLE CATEGORIES
(
    Id    SERIAL PRIMARY KEY,
    Label VARCHAR NOT NULL
);

CREATE TABLE TYPES
(
    Id    SERIAL PRIMARY KEY,
    Label VARCHAR NOT NULL
);

CREATE TABLE PRODUCTS
(
    Id          SERIAL PRIMARY KEY,
    CategoryId  INTEGER NOT NULL,
    Label       VARCHAR,
    Type        INTEGER NOT NULL,
    DownloadUrl VARCHAR,
    Weight      float,
    CONSTRAINT fk_products_categories
        FOREIGN KEY (CategoryId)
            REFERENCES CATEGORIES (Id),
    CONSTRAINT fk_products_types
        FOREIGN KEY (Type)
            REFERENCES TYPES (Id)

);

CREATE TABLE CLIENTS
(
    Id   SERIAL PRIMARY KEY,
    Name VARCHAR NOT NULL
);

CREATE TABLE CARTS
(
    Id       SERIAL PRIMARY KEY,
    ClientId INTEGER NOT NULL,
    CONSTRAINT fk_carts_clients
        FOREIGN KEY (ClientId)
            REFERENCES CLIENTS (Id)

);

CREATE TABLE PRODUCTS_CARTS
(
    Id        SERIAL PRIMARY KEY,
    ProductId INTEGER NOT NULL,
    CartId    INTEGER NOT NULL,
    CONSTRAINT fk_products_carts_products
        FOREIGN KEY (ProductId)
            REFERENCES PRODUCTS (Id),
    CONSTRAINT fk_products_carts_carts
        FOREIGN KEY (CartId)
            REFERENCES CARTS (Id)

);

CREATE TABLE ORDERS
(
    Id       SERIAL PRIMARY KEY,
    ClientId INTEGER NOT NULL,
    CONSTRAINT fk_orders_clients
        FOREIGN KEY (ClientId)
            REFERENCES CLIENTS (Id)
);

CREATE TABLE PRODUCTS_ORDERS
(
    Id        SERIAL PRIMARY KEY,
    ProductId INTEGER NOT NULL,
    OrderId   INTEGER NOT NULL,
    CONSTRAINT fk_products_orders_products
        FOREIGN KEY (ProductId)
            REFERENCES PRODUCTS (Id),
    CONSTRAINT fk_products_orders_orders
        FOREIGN KEY (OrderId)
            REFERENCES ORDERS (Id)

);

--categories
INSERT INTO CATEGORIES (Label)
VALUES ('Technology');
INSERT INTO CATEGORIES (Label)
VALUES ('Vehicles');
INSERT INTO CATEGORIES (Label)
VALUES ('Construction');

--types
INSERT INTO TYPES (Label)
VALUES ('Physical');
INSERT INTO TYPES (Label)
VALUES ('Digital');

--products
INSERT INTO PRODUCTS (CategoryId, Label, Type, DownloadUrl, Weight)
VALUES (1, 'Asus Laptop', 1, null, 5);
INSERT INTO PRODUCTS (CategoryId, Label, Type, DownloadUrl, Weight)
VALUES (1, 'Mechanical keyboard', 1, null, 2);
INSERT INTO PRODUCTS (CategoryId, Label, Type, DownloadUrl, Weight)
VALUES (1, 'Mouse', 1, null, 1);
--clients
INSERT INTO CLIENTS (NAME) VALUES ('John Doe');
INSERT INTO CLIENTS (NAME) VALUES ('Jane Doe');

--carts
INSERT INTO CARTS (ClientId) VALUES (1);
INSERT INTO CARTS (ClientId) VALUES (2);