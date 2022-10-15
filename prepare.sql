create database if not exists todo;

create table todo.task (
    id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);