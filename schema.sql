begin;
create table if not exists User (
    user_id bigserial primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    document varchar(14) not null,
    balance decimal(12, 2) not null,
    appurtenance varchar(255) not null
    check(document in ('CPF', 'CNPJ')),
    check(balance >= 0),
    check(appurtenance in ('COMMON', 'MERCHANT'))
);
create index idx_user_email on User(email);
create index idx_user_document on User(document);

create table if not exists operation (
    operation_id bigserial primary key,
    sender_user_id bigint references USER(user_id) not null
    receiver_user_id bigint references USER(user_id) not null
);
create index idx_users_operation on User(sender_user_id, receiver_user_id);
commit;