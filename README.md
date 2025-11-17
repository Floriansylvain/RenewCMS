# RenewCMS

## 🚧 This project is under development !

The features are not complete, not fully tested and subject to many changes.

## Setup

Don't forget to setup the [Environment variables](#environment-variables)!

### Use with Docker

#### Compose

The easiest way to run the project is to use docker-compose.
You still need to set up the environment variables, but you can use the `.env` file for that.
Note that some supplementary variables are required for docker-compose to work, like `DOCKER_DB_FOLDER`.

```bash
docker compose up --build -d
```

### Build and run it yourself

#### Build

There are 2 scripts to build the project:

- `super_build.sh`
- `super_build.bat`

They will build the project for windows, linux and macOS. The binaries will be in the `bin` folder.
You can also really just use `go build` to build the project for your current OS.

#### Run

Once the project is built, you can simply run the binary for your OS. Setting the environment variables is
of course required, but not necessarily via the `.env` file.

### Environment variables

| Name                 | Type   | Description                                   | Comment                                                                                                               |
|----------------------|--------|:----------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| ENVIRONMENT          | string | environment the API is running in             | required, `development` or `production`                                                                               |
| HOST                 | string | the host domain name for emails and callbacks | required, if `localhost`: add the port (e.g. `localhost:8080`)                                                        |
| PORT                 | int    | port the API will use                         | required                                                                                                              |
| JWT_SECRET           | string | secret for the jwt auth                       | required                                                                                                              |
| CORS_ALLOWED_ORIGINS | string | allowed origins for CORS                      | required, semicolon separated list                                                                                    |
| DB_FILE              | string | path to the sqlite db file                    | required, can be at the root but name still required (e.g. `./RenewCMS.db`) ; have to end up with `.db`               |
| DOCKER_DB_FOLDER     | string | path to the sqlite db file's folder           | required with docker, this will basically be the host machine folder (e.g. `./data`) that contains the sqlite db file |
| SMTP_EMAIL           | string | sender email                                  | required                                                                                                              |
| SMTP_PASSWORD        | string | smtp account password                         | required                                                                                                              |
| SMTP_HOST            | string | smtp server address                           | required                                                                                                              | 
| SMTP_PORT            | int    | smtp server port                              | required                                                                                                              |

## API Usage

TODO

## Demo

TODO

## TODOs

- Cancel button when email validation pending or expired
