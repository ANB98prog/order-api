-- +goose Up
CREATE TABLE public.users (
    id bigserial NOT NULL,
    "name" text NULL,
    phone text NULL,
    CONSTRAINT uni_users_phone UNIQUE (phone),
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE public.products (
     id bigserial NOT NULL,
     "name" text NULL,
     description text NULL,
     price numeric NULL,
     images _text NULL,
     CONSTRAINT products_pkey PRIMARY KEY (id),
     CONSTRAINT uni_products_name UNIQUE ("name")
);

CREATE TABLE public.orders (
     id bigserial NOT NULL,
     user_id int8 NULL,
     "date" timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
     total numeric NULL,
     CONSTRAINT orders_pkey PRIMARY KEY (id),
     CONSTRAINT fk_users_order FOREIGN KEY (user_id) REFERENCES public.users(id)
);
CREATE INDEX idx_orders_user_id ON public.orders USING btree (user_id);

CREATE TABLE public.order_items (
     id bigserial NOT NULL,
     order_id int8 NULL,
     product_id int8 NULL,
     quantity int8 NULL,
     price numeric NULL,
     CONSTRAINT order_items_pkey PRIMARY KEY (id),
     CONSTRAINT fk_order_items_product FOREIGN KEY (product_id) REFERENCES public.products(id),
     CONSTRAINT fk_orders_items FOREIGN KEY (order_id) REFERENCES public.orders(id)
);
CREATE INDEX idx_order_items_order_id ON public.order_items USING btree (order_id);
CREATE INDEX idx_order_items_product_id ON public.order_items USING btree (product_id);


-- +goose Down
drop table if exists public.users;
drop table if exists public.products;
drop table if exists public.orders;
drop table if exists public.order_items;
