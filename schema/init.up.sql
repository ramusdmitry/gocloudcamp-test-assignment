CREATE TABLE playlist
(
    id serial not null unique,
    name varchar(255) not null,
)

CREATE TABLE songs
(
    id serial not null unique,
    playlist_id int references playlist (id) on delete no action on update no action not null,
    title varchar(255) not null,
    duration int not null
)
