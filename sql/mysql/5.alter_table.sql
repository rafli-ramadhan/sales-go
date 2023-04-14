alter table transaction rename column transactionnumber to transaction_number

alter table transaction modify column transaction_number int not null
