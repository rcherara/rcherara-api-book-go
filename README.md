# rcherara-api-book

Simple RESTful API to create, read, update and delete books.

The Service  is exposed on port 8900

# Quick Start 

    # Install mux router
    $ go get -u github.com/gorilla/mux

    # Install getenv
    $ go get -u github.com/subosito/gotenv

    # Install db sql
    $ go get -u database/sql

    # Install pq client
    $go get -u github.com/lib/pq

# How it works

# Create new DB rcherara-books on elephantsql.com and run the following request on SQL Browser

    $ create table books (id serial, title varchar, author varchar, year varchar)

    $ insert into books(title, author, year) values ('Spring 5.0 By Example: Grasp the fundamentals of Spring 5.0 to build modern, robust, and scalable Java applications.', 'Mr. Claudio Eduardo de Oliveira', '2010');
    $ insert into books(title, author, year) values ('Network Security with OpenSSL', 'Mr. John Viega.', '2019');
    insert into books(title, author, year) values ('Learn AutoCAD!', 'Mr. David Martin', '2016');
    $ insert into books(title, author, year) values ('Ansible for DevOps', 'Mr. Jeff Geerling', '2015');
 

# Store your ELEPHANTSQL_URL of DB access  in configuration file named  .env  on root diretory of projet 


ELEPHANTSQL_URL="postgres://user:passpassword@baasu.db.elephantsql.com:5432/iwcmsppt"

We read the values of this file .env using this code :

    gotenv.Load()


For more information see  :
https://www.elephantsql.com/docs/go.html

# Endpoints

Get All Books    : GET
Get Single Book  : GET
Delete Book      : DELETE 
Create Book      : POST 
Update Book      : PUT    http://localhost:8900/books 

