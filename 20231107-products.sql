create table products
(
    -- id serial primary key, --serial自動編號主鍵
    id bigint primary key generated by default as identity,--自動編號主鍵（目前常用）
    name text not null,
    price int not null,
    check (price >= 0)
);