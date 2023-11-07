begin;--開始交易
insert into products (
    name,price
)values(
    'Rashguard',1500
), (
    'Trousers',1200
);

commit;--交易結束