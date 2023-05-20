SET
search_path TO public;


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

CREATE TABLE public.users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name  VARCHAR(50),
    username   VARCHAR(50),
    email      VARCHAR(50),
    role_id    INTEGER REFERENCES roles (id) NOT NULL,
    is_new     BOOLEAN DEFAULT TRUE          NOT NULL,
    created_at TIMESTAMP,
    created_by TEXT                          NOT NULL,
    updated_at TIMESTAMP,
    updated_by TEXT                          NOT NULL,
    UNIQUE (username)
);


CREATE TABLE public.genres
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255)             NOT NULL,
    type_code  VARCHAR(25)              NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by TEXT                     NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by TEXT                     NOT NULL
);

CREATE TABLE public.movies
(
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(255)             NOT NULL,
    type_code    VARCHAR(25)              NOT NULL,
    release_date TIMESTAMP                NOT NULL,
    runtime      INTEGER                  NOT NULL,
    mpaa_rating  VARCHAR(25)              NOT NULL,
    description  TEXT                     NOT NULL,
    image_path   VARCHAR(255),
    video_path   VARCHAR(255),
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by   TEXT                     NOT NULL,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by   TEXT                     NOT NULL
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
    id         SERIAL PRIMARY KEY,
    code       VARCHAR(25),
    name       VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by TEXT                     NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by TEXT                     NOT NULL
);

CREATE TABLE public.views
(
    id        SERIAL PRIMARY KEY,
    viewed_by TEXT                     NOT NULL,
    viewed_at TIMESTAMP WITH TIME ZONE NOT NULL,
    movie_id  INTEGER REFERENCES movies (id)
);

INSERT INTO mpaa(code, name, created_at, created_by, updated_at, updated_by)
VALUES ('G', 'G', now(), 'admin', now(), 'admin'),
       ('PG', 'PG', now(), 'admin', now(), 'admin'),
       ('PG13', 'PG-13', now(), 'admin', now(), 'admin'),
       ('R', 'R', now(), 'admin', now(), 'admin'),
       ('NC17', 'NC-17', now(), 'admin', now(), 'admin'),
       ('18A', '18A', now(), 'admin', now(), 'admin');

INSERT INTO public.genres (name, type_code, created_at, created_by, updated_at, updated_by)
VALUES ('Comedy', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Science Fiction', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Horror', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Romance', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Action', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Thriller', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Drama', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Mystery', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Crime', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Animation', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Adventure', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Fantasy', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Superhero', 'MOVIE', now(), 'admin', now(), 'admin'),
       ('Comedy', 'TV', now(), 'admin', now(), 'admin'),
       ('Soap', 'TV', now(), 'admin', now(), 'admin'),
       ('Action & Adventure', 'TV', now(), 'admin', now(), 'admin'),
       ('Horror', 'TV', now(), 'admin', now(), 'admin'),
       ('Family', 'TV', now(), 'admin', now(), 'admin'),
       ('Thriller', 'TV', now(), 'admin', now(), 'admin'),
       ('Drama', 'TV', now(), 'admin', now(), 'admin'),
       ('Kids', 'TV', now(), 'admin', now(), 'admin'),
       ('Crime', 'TV', now(), 'admin', now(), 'admin'),
       ('Animation', 'TV', now(), 'admin', now(), 'admin'),
       ('Reality', 'TV', now(), 'admin', now(), 'admin'),
       ('Fantasy', 'TV', now(), 'admin', now(), 'admin'),
       ('News', 'TV', now(), 'admin', now(), 'admin'),
       ('Talk', 'TV', now(), 'admin', now(), 'admin'),
       ('War & Politics', 'TV', now(), 'admin', now(), 'admin');

INSERT INTO public.movies (title, type_code, release_date, runtime, mpaa_rating, description, image_path, created_at,
                           created_by,
                           updated_at, updated_by)
VALUES ('Highlander', 'MOVIE', '1986-03-07', 116, 'R',
        'He fought his first battle on the Scottish Highlands in 1536. He will fight his greatest battle on the streets of New York City in 1986. His name is Connor MacLeod. He is immortal.',
        '/8Z8dptJEypuLoOQro1WugD855YE.jpg', now(), 'admin', now(), 'admin'),
       ('Raiders of the Lost Ark', 'MOVIE', '1981-06-12', 115, 'PG-13',
        'Archaeology professor Indiana Jones ventures to seize a biblical artefact known as the Ark of the Covenant. While doing so, he puts up a fight against Renee and a troop of Nazis.',
        '/ceG9VzoRAVGwivFU403Wc3AHRys.jpg', now(), 'admin', now(), 'admin'),
       ('The Godfather', 'MOVIE', '1972-03-24', 175, '18A',
        'The aging patriarch of an organized crime dynasty in postwar New York City transfers control of his clandestine empire to his reluctant youngest son.',
        '/3bhkrj58Vtu7enYsRolD1fZdja1.jpg', now(), 'admin', now(), 'admin'),
       ('Thor: Ragnarok', 'MOVIE', '2017-03-11', 131, 'PG-13',
        'Thor is imprisoned on the other side of the universe and finds himself in a race against time to get back to Asgard to stop Ragnarok, the destruction of his home-world and the end of Asgardian civilization, at the hands of a powerful new threat, the ruthless Hela.',
        '/rzRwTcFvttcN1ZpX2xv4j3tSdJu.jpg', now(), 'admin', now(), 'admin'),
       ('Harry Maguire', 'TV', '2023-06-06', 11, 'R',
        'I spent 13 hours to watch all defensive skills highlight videos on Harry Maguire Manchester United'' career',
        null, now(), 'admin', now(), 'admin');

INSERT INTO public.movies_genres (movie_id, genre_id)
VALUES ((SELECT m.id FROM movies m WHERE m.title = 'Highlander'),
        (SELECT g.id FROM genres g WHERE g.name = 'Action' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Highlander'),
        (SELECT g.id FROM genres g WHERE g.name = 'Fantasy' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Raiders of the Lost Ark'),
        (SELECT g.id FROM genres g WHERE g.name = 'Action' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Raiders of the Lost Ark'),
        (SELECT g.id FROM genres g WHERE g.name = 'Adventure' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'The Godfather'),
        (SELECT g.id FROM genres g WHERE g.name = 'Drama' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'The Godfather'),
        (SELECT g.id FROM genres g WHERE g.name = 'Crime' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok'),
        (SELECT g.id FROM genres g WHERE g.name = 'Action' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok'),
        (SELECT g.id FROM genres g WHERE g.name = 'Adventure' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok'),
        (SELECT g.id FROM genres g WHERE g.name = 'Fantasy' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok'),
        (SELECT g.id FROM genres g WHERE g.name = 'Science Fiction' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok'),
        (SELECT g.id FROM genres g WHERE g.name = 'Comedy' AND g.type_code = 'MOVIE')),
       ((SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire'),
        (SELECT g.id FROM genres g WHERE g.name = 'Comedy' AND g.type_code = 'TV'));

INSERT INTO public.roles(role_name, role_code, created_at, created_by, updated_at, updated_by)
VALUES ('admin', 'ADMIN', now(), 'admin', now(), 'admin'),
       ('moderator', 'MOD', now(), 'admin', now(), 'admin'),
       ('general', 'GENERAL', now(), 'admin', now(), 'admin'),
       ('banned', 'BANNED', now(), 'admin', now(), 'admin');

INSERT INTO public.users(first_name, last_name, username, email, role_id, is_new, created_at, created_by, updated_at,
                         updated_by)
VALUES ('Admin', 'User', 'root', 'admin@example.com', (SELECT r.id FROM public.roles r WHERE role_code = 'ADMIN'),
        false,
        now(), 'admin', now(), 'admin');

INSERT INTO public.views(viewed_at, viewed_by, movie_id)
VALUES ('2023-03-01', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok')),
       ('2023-03-01', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok')),
       ('2023-04-01', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok')),
       ('2023-04-05', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok')),
       ('2023-04-06', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok')),
       ('2023-04-06', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Thor: Ragnarok')),
       ('2023-04-06', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Highlander')),
       ('2023-04-06', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Highlander')),
       ('2023-04-06', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Highlander')),
       ('2023-04-06', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire'));
