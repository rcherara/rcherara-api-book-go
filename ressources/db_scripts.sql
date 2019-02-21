CREATE USER postgres_user WITH PASSWORD 'password';
CREATE DATABASE my_postgres_db OWNER postgres_user;


CREATE USER postgres_api_user WITH PASSWORD 'password';
CREATE DATABASE my_postgres_db_api_book OWNER postgres_api_user;

CREATE TABLE books (id serial, title varchar, author varchar, year varchar);

INSERT INTO   books(title, author, year) values ('Spring 5.0 By Example', 'Mr. Claudio Eduardo de Oliveira', '2010');
INSERT INTO   books(title, author, year) values ('Network Security with OpenSSL', 'Mr. John Viega.', '2019');
INSERT INTO   books(title, author, year) values ('Learn AutoCAD!', 'Mr. David Martin', '2016');
INSERT INTO   books(title, author, year) values ('Ansible for DevOps', 'Mr. Jeff Geerling', '2015');
 