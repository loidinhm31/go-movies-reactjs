# STRUCTURE

1. `API`: port `9090`
2. `Website`: port `3000`
3. `Keycloak Admin Console`: port `8086`

## ShiftFlix API
This API build using
### Major Dependencies
1.  [Gin](https://gin-gonic.com/): A web framework written in Golang.
2.  [GORM](https://gorm.io/): The fantastic ORM library for Golang.
3.  [Gocloak](https://github.com/Nerzal/gocloak/): golang keycloak client.

### Database
1. Use `POSTGRESQL` with port `5432`, create a new database named `mdb`
2. SQL script under `movies-back-end/sql`

### Code Layout
#### Golang Code

1. Main function to run under `cmd/api`
2. All configurations are stored in `config` folder for different environments.
3. All routes are init under `internal/server`

## ShiftFlix NextJS Website
### Major Dependencies
This website build using
1.  [npm](https://www.npmjs.com/): The node package manager for building.
2.  [React](https://reactjs.org/): The core frontend framework.
3.  [Next.js](https://nextjs.org/): A React scaffolding framework to streamline development.
4.  [NextAuth.js](https://next-auth.js.org/): A user authentication framework to ensure we handle accounts with best
    practices.
5.  [MUI](https://mui.com/): A wide collection of pre-built UI components that generally look pretty go

### Purpose
1. User registration using Keycloak.
2. Visit and manage movies
3. Search information about movies

#### Set up your environment
1.  Node 16: if you are on windows, you can [download node from their website](https://nodejs.org/en/download/releases),
    if you are on linux, use [NVM](https://github.com/nvm-sh/nvm) (Once installed, run `nvm use 16`)
2.  [Docker](https://www.docker.com/): This project use docker to simplify running dependent services.

#### Using debug user credentials

You can use the debug credentials provider to log in without fancy emails or OAuth.

1. This feature is automatically on in development mode, i.e. when you run `npm run dev`. In case you want to do the
   same with a production build (for example, the docker image), then run the website with environment variable
   `DEBUG_LOGIN=true`.
2. Use the `SIGN IN` button in the top right to go to the login page.
3. You should see a section for debug credentials. Enter any username you wish, you will be logged in as that user.


### Code Layout

#### React Code

All react code is under `src/` with a few subdirectories:

1.  `pages/`: All pages a user could navigate too and API URLs which are under `pages/api/`.
2.  `components/`: All re-usable React components. If something gets used twice we should create a component and put it
    here.
3.  `lib/`: A generic place to store library files that are used anywhere. This doesn't have much structure yet.

NOTE: `styles/` can be ignored for now.

## Getting everything up and running

If you're doing active development we suggest the following workflow:

1. Make sure you run `npm i` and `npm run build` in `movies-front-end` folder, and `go build -o apiBinary ./cmd/api` in `movies-back-end` folder
    - Currently, docker files is not support to build on docker, this project just create images locally, and transfer them for docker compose
2. Open the terminal, navigate to the `project` folder.
3. Run `docker compose up --build`. You can optionally include `-d` to
    detach and later track the logs if desired.
4. If you don't want to run docker compose on your own, you can use `Makefile` to automatically execute this process to build images and run docker
    - Make sure you have Make on your system
    - At `project` folder, open the terminal and run `make up_build`
5. Now, you can access Keycloak Admin Console at `localhost:8086/admin` with default username `admin` and password `admin`

## Export realm from Keycloak
```
/opt/keycloak/bin/kc.sh export --dir /opt/keycloak/data/import --realm mdb --users realm_file
```
```
docker cp <container_id>:/opt/keycloak/data/import/mdb-realm.json c:/temp/mdb-realm.json
```
## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.
