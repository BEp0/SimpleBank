begin;
create table if not exists account (
    accont_id bigserial primary key,
    owner varchar(255) not null,
    balance bigint not null,
    currency varchar(3) not null,
    created_at timestamp default NOW(),
    check(currency in ('USD', 'EUR', 'BR'))
);
create index idx_owner on account(owner);

create table if not exists entrie (
    entrie_id bigserial primary key,
    accont_id bigint references ACCONT(accont_id) not null
);
create index idx_account_id on entrie(accont_id);

create table if not exists transfer (
    transfer_id bigserial primary key,
    from_account_id bigint references account(accont_id) not null,
    to_account_id bigint references account(accont_id) not null
);
create index idx_account_ids on transfer(from_accont_id, to_accont_id);
commit;