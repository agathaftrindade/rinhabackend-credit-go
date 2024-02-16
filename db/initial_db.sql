create table Accounts(
    Account_ID serial primary key,
    ALimit bigint,
    Balance bigint
);

create type Transaction_Type AS ENUM ('c', 'd');
create table Transactions (
    Transaction_ID serial primary key,
    Amount bigint,
    Account_ID bigint references Accounts(Account_ID),
    Type Transaction_Type,
    Description varchar(15),
    CreatedAt timestamp default now()
);


INSERT INTO Accounts
(Account_ID, Alimit, balance)
VALUES
(1, 100000, 0),
(2, 80000, 0),
(3, 1000000, 0),
(4, 10000000, 0),
(5, 500000,	0);

SELECT setval('accounts_account_id_seq', (SELECT MAX(Account_ID) FROM Accounts) + 1);