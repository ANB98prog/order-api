-- +goose Up
insert into public.products("name", "description", "price")
values ('MacBook Air', 'A good one laptop', 120000),
       ('Lenovo', 'Another one', 100000),
       ('Surface', 'Another one', 160000);


-- +goose Down
-- +goose StatementBegin
delete from public.products
where "name" in ('MacBook Air', 'Lenovo', 'Surface')
-- +goose StatementEnd
