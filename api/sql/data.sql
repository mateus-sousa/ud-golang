insert into users (id, name, nick, email, password) values
    (null, "usuario1", "usuario1", "usuario1@email.com", "$2a$10$tTJtk1QwZCbF6fa3JZMoRefKL6Wr2exj8ev6GRGKmE8cHaZwRjsyG"),
    (null, "usuario2", "usuario2", "usuario2@email.com", "$2a$10$Zqk5MSEASfW81/r3qIR3SOfgU9ue0QbGJNnOVKTjrnId7cSdSGQn2"),
    (null, "usuario3", "usuario3", "usuario3@email.com", "$2a$10$cQcWVFakmP7htTozkBG9FOSh75ALSehuAUXxTU2Hn1V4Gy63WIara");

insert into followers (user_id, follower_id) values
(1, 2),
(3,1),
(1,3);