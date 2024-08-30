insert into users (name, nick, email, password)
values
("User 1", "user_1", "user1@gmail.com", "$2y$10$rx/TunEnBaU4sLhKyaA4j.vSaEYgrogriMWX41Cu.gI.wngL0PIc."), 
("User 2", "user_2", "user2@gmail.com", "$2y$10$rx/TunEnBaU4sLhKyaA4j.vSaEYgrogriMWX41Cu.gI.wngL0PIc."),
("User 3", "user_3", "user3@gmail.com", "$2y$10$rx/TunEnBaU4sLhKyaA4j.vSaEYgrogriMWX41Cu.gI.wngL0PIc.");

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);