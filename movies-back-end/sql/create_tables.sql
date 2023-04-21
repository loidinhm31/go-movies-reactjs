SET
search_path TO public;

CREATE TABLE public.users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name  VARCHAR(50),
    email      VARCHAR(50),
    password   VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE public.roles
(
    id         SERIAL PRIMARY KEY       NOT NULL,
    role_name  TEXT                     NOT NULL,
    role_code  TEXT                     NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by TEXT                     NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by TEXT                     NOT NULL
);

CREATE TABLE user_role
(
    id         SERIAL PRIMARY KEY       NOT NULL,
    user_id    SERIAL                   NOT NULL,
    role_id    SERIAL                   NOT NULL,
    is_active  BOOLEAN                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by TEXT                     NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by TEXT                     NOT NULL
);

ALTER TABLE user_role
    ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_role
    ADD FOREIGN KEY (role_id) REFERENCES roles (id);

CREATE TABLE public.genres
(
    id         SERIAL PRIMARY KEY,
    genre      VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE public.movies
(
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    release_date TIMESTAMP NOT NULL,
    runtime      INTEGER NOT NULL,
    mpaa_rating  VARCHAR(25) NOT NULL,
    description  TEXT NOT NULL,
    image        VARCHAR(255),
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE public.movies_genres
(
    id       SERIAL PRIMARY KEY,
    movie_id INTEGER REFERENCES movies (id),
    genre_id INTEGER REFERENCES genres (id),
    UNIQUE (movie_id, genre_id)
);

CREATE TABLE public.mpaa
(
    id   SERIAL PRIMARY KEY,
    code VARCHAR(25),
    name VARCHAR(50)
);

INSERT INTO mpaa(code, name)
VALUES ('G', 'G'),
       ('PG', 'PG'),
       ('PG13', 'PG-13'),
       ('R', 'R'),
       ('NC17', 'NC-17'),
       ('18A', '18A');

INSERT INTO public.genres (genre, created_at, updated_at)
VALUES ('Comedy', now(), now()),
       ('Science Fiction', now(), now()),
       ('Horror', now(), now()),
       ('Romance', now(), now()),
       ('Action', now(), now()),
       ('Thriller', now(), now()),
       ('Drama', now(), now()),
       ('Mystery', now(), now()),
       ('Crime', now(), now()),
       ('Animation', now(), now()),
       ('Adventure', now(), now()),
       ('Fantasy', now(), now()),
       ('Superhero', now(), now());

INSERT INTO public.movies (title, release_date, runtime, mpaa_rating, description, image, created_at, updated_at)
VALUES ('Highlander', '1986-03-07', 116, 'R',
        'He fought his first battle on the Scottish Highlands in 1536. He will fight his greatest battle on the streets of New York City in 1986. His name is Connor MacLeod. He is immortal.',
        '/8Z8dptJEypuLoOQro1WugD855YE.jpg', now(), now()),
       ('Raiders of the Lost Ark', '1981-06-12', 115, 'PG-13',
        'Archaeology professor Indiana Jones ventures to seize a biblical artefact known as the Ark of the Covenant. While doing so, he puts up a fight against Renee and a troop of Nazis.',
        '/ceG9VzoRAVGwivFU403Wc3AHRys.jpg', now(), now()),
       ('The Godfather', '1972-03-24', 175, '18A',
        'The aging patriarch of an organized crime dynasty in postwar New York City transfers control of his clandestine empire to his reluctant youngest son.',
        '/3bhkrj58Vtu7enYsRolD1fZdja1.jpg', now(), now()),
       ('Thor: Ragnarok', '2017-03-11', 131, 'PG-13',
        'Thor is imprisoned on the other side of the universe and finds himself in a race against time to get back to Asgard to stop Ragnarok, the destruction of his home-world and the end of Asgardian civilization, at the hands of a powerful new threat, the ruthless Hela.',
        '/rzRwTcFvttcN1ZpX2xv4j3tSdJu.jpg', now(), now());

INSERT INTO public.movies_genres (movie_id, genre_id)
VALUES ((SELECT m.id FROM movies m WHERE m.title = 'Highlander'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Action')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Highlander'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Fantasy')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Raiders of the Lost Ark'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Action')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Raiders of the Lost Ark'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Adventure')),
       ((SELECT m.id FROM movies m WHERE m.title = 'The Godfather'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Drama')),
       ((SELECT m.id FROM movies m WHERE m.title = 'The Godfather'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Crime')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Action')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Adventure')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Fantasy')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Science Fiction')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok'),
        (SELECT g.id FROM genres g WHERE g.genre = 'Comedy'));

INSERT INTO public.users (first_name, last_name, email, password, created_at, updated_at)
VALUES ('Admin', 'User', 'admin@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy',
        now(), now());

