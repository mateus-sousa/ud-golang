CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publishes;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id INT auto_increment primary key,
    name VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers(
    user_id int not null,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    follower_id int not null,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    primary key (user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE publishes(
    id int auto_increment primary key,
    title varchar(50) not null,
    content varchar(300) not null,
    author_id int not null,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE,
    likes int default 0,
    created_at timestamp default current_timestamp
)