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
    image_url    VARCHAR(255) DEFAULT NULL,
    video_path   VARCHAR(255) DEFAULT NULL,
    price        FLOAT        DEFAULT NULL,
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

CREATE TABLE public.ratings
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

CREATE TABLE public.seasons
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50)                    NOT NULL,
    air_date    TIMESTAMP                      NOT NULL,
    description TEXT                           NOT NULL,
    movie_id    INTEGER REFERENCES movies (id) NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE       NOT NULL,
    created_by  TEXT                           NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE       NOT NULL,
    updated_by  TEXT                           NOT NULL
);

CREATE TABLE public.episodes
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(50)                     NOT NULL,
    air_date   TIMESTAMP                       NOT NULL,
    runtime    INTEGER                         NOT NULL,
    video_path VARCHAR(255) DEFAULT NULL,
    season_id  INTEGER REFERENCES seasons (id) NOT NULL,
    price      FLOAT        DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE        NOT NULL,
    created_by TEXT                            NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE        NOT NULL,
    updated_by TEXT                            NOT NULL
);

CREATE TABLE public.payments
(
    id                  SERIAL PRIMARY KEY,
    user_id             INTEGER REFERENCES users (id),
    ref_id              INTEGER                  NOT NULL,
    type_code           VARCHAR(10)              NOT NULL,
    provider            VARCHAR(255)             NOT NULL,
    provider_payment_id VARCHAR(255) DEFAULT NULL,
    amount              FLOAT                    NOT NULL,
    received            FLOAT                    NOT NULL,
    currency            VARCHAR(10)              NOT NULL,
    payment_method      VARCHAR(100)             NOT NULL,
    status              VARCHAR(50)              NOT NULL,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by          TEXT                     NOT NULL
);

CREATE TABLE public.collections
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (id),
    movie_id   INTEGER DEFAULT NULL,
    episode_id INTEGER DEFAULT NULL,
    type_code  VARCHAR(10)              NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by TEXT                     NOT NULL
);

INSERT INTO public.ratings(code, name, created_at, created_by, updated_at, updated_by)
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

INSERT INTO public.movies (title, type_code, release_date, runtime, mpaa_rating, description, image_url, price,
                           created_at,
                           created_by,
                           updated_at, updated_by)
VALUES ('Highlander', 'MOVIE', '1986-03-07', 116, 'R',
        'He fought his first battle on the Scottish Highlands in 1536. He will fight his greatest battle on the streets of New York City in 1986. His name is Connor MacLeod. He is immortal.',
        'https://image.tmdb.org/t/p/w200/8Z8dptJEypuLoOQro1WugD855YE.jpg',
        199.0,
        now(), 'admin', now(), 'admin'),
       ('Raiders of the Lost Ark', 'MOVIE', '1981-06-12', 115, 'PG-13',
        'Archaeology professor Indiana Jones ventures to seize a biblical artefact known as the Ark of the Covenant. While doing so, he puts up a fight against Renee and a troop of Nazis.',
        'https://image.tmdb.org/t/p/w200/ceG9VzoRAVGwivFU403Wc3AHRys.jpg',
        299.0,
        now(), 'admin', now(), 'admin'),
       ('The Godfather', 'MOVIE', '1972-03-24', 175, '18A',
        'The aging patriarch of an organized crime dynasty in postwar New York City transfers control of his clandestine empire to his reluctant youngest son.',
        'https://image.tmdb.org/t/p/w200/3bhkrj58Vtu7enYsRolD1fZdja1.jpg',
        null,
        now(), 'admin', now(), 'admin'),
       ('Thor: Ragnarok', 'MOVIE', '2017-03-11', 131, 'PG-13',
        'Thor is imprisoned on the other side of the universe and finds himself in a race against time to get back to Asgard to stop Ragnarok, the destruction of his home-world and the end of Asgardian civilization, at the hands of a powerful new threat, the ruthless Hela.',
        'https://image.tmdb.org/t/p/w200/rzRwTcFvttcN1ZpX2xv4j3tSdJu.jpg',
        null,
        now(), 'admin', now(), 'admin'),
       ('Harry Maguire', 'TV', '2023-06-06', 11, 'R',
        'I spent 13 hours to watch all defensive skills highlight videos on Harry Maguire Manchester United\'' career',
        'https://res.cloudinary.com/dln5uctjy/image/upload/c_scale,w_200,h_300/shiftflix/images/346637967_915526433037296_4038910118938564751_n_n0otyr.jpg',
        null,
        now(), 'admin', now(), 'admin');

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

INSERT INTO public.seasons(name, air_date, description, movie_id, created_at, created_by, updated_at, updated_by)
VALUES ('Season 1', '2019-06-01', 'First Season at OT', (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire'),
        now(), 'admin', now(), 'admin'),
       ('Season 2', '2020-06-01', 'Second Season at OT', (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire'),
        now(), 'admin', now(), 'admin'),
       ('Season 3', '2021-06-01', 'Third Season at OT', (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire'),
        now(), 'admin', now(), 'admin'),
       ('Season 4', '2022-06-01', 'Fourth Season at OT', (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire'),
        now(), 'admin', now(), 'admin');

INSERT INTO public.episodes(name, air_date, runtime, season_id, created_at, created_by, updated_at, updated_by)
VALUES ('Episode 1', '2019-06-01', 60, (SELECT s.id
                                        FROM seasons s
                                        WHERE s.name = 'Season 1'
                                          AND s.movie_id = (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire')),
        now(), 'admin', now(), 'admin'),
       ('Episode 2', '2020-06-01', 60, (SELECT s.id
                                        FROM seasons s
                                        WHERE s.name = 'Season 1'
                                          AND s.movie_id = (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire')),
        now(), 'admin', now(), 'admin'),
       ('Episode 3', '2021-06-01', 60, (SELECT s.id
                                        FROM seasons s
                                        WHERE s.name = 'Season 1'
                                          AND s.movie_id = (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire')),
        now(), 'admin', now(), 'admin'),
       ('Episode 4', '2022-06-01', 60, (SELECT s.id
                                        FROM seasons s
                                        WHERE s.name = 'Season 1'
                                          AND s.movie_id = (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire')),
        now(), 'admin', now(), 'admin'),
       ('Episode 1', '2022-06-01', 60, (SELECT s.id
                                        FROM seasons s
                                        WHERE s.name = 'Season 2'
                                          AND s.movie_id = (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire')),
        now(), 'admin', now(), 'admin'),
       ('Episode 2', '2022-06-01', 60, (SELECT s.id
                                        FROM seasons s
                                        WHERE s.name = 'Season 2'
                                          AND s.movie_id = (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire')),
        now(), 'admin', now(), 'admin'),
       ('Episode 3', '2022-06-01', 60, (SELECT s.id
                                        FROM seasons s
                                        WHERE s.name = 'Season 2'
                                          AND s.movie_id = (SELECT m.id FROM movies m WHERE m.title = 'Harry Maguire')),
        now(), 'admin', now(), 'admin');

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
       ('2023-04-06', 'anonymous', (SELECT m.id FROM movies m WHERE m.title = 'Highlander'));
