psql -U postgres -d tododb << "EOSQL"
create table todo (ID integer PRIMARY KEY, Title varchar(10) NOT NULL, Content varchar(100) , Status varchar(1));
insert into todo (ID, Title, Content, Status) values (1, 'todo1', 'sample-todo', '0');
EOSQL
