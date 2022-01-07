# Golang - React Todo App

[![go-github docker-image (latest SemVer)](https://img.shields.io/github/v/release/google/go-github?sort=semver)](https://hub.docker.com/repository/docker/barisakdas/react-todo-ui)
[![go-github docker-image (latest SemVer)](https://img.shields.io/github/v/release/google/go-github?sort=semver)](https://hub.docker.com/repository/docker/barisakdas/golang-todo-api)

The project is designed as an api-client application. The api side is built with golang using standard libraries. Postgresql is selected as the database. The client side is created with React.

You can build the whole project using the docker-compose.yaml file.

The `docker-compose up` command is enough to run the project.

## Installation

If you want to download and use the projects separately, you can take the images from `hub.docker.com` and create your own containers.

```bash
docker pull barisakdas/react-todo-ui
docker pull barisakdas/golang-todo-api
```

If you want to run the codes by getting the codes on github, the project can be reached by pulling the repository as follows.

```go
git clone https://github.com/barisakdas/Golang-React-Todo-App
```

## Usage

The application can be used from the front side. For this, you can do the operations from the interface that will appear at `localhost:3000`.

The API is running in the background with `localhost:8080`. You can learn the endpoints on the API and questions such as how it can be used by sending a get request to `localhost:8080/`. This address will give you all the endpoints as json, their purpose and usage examples.

```bash
curl http://localhost:3000/
```

When it is desired to run with the codes on Github, the `npm install` command should be run first to get the node packages. This command will include the required packages in the project.Then you can run the project with the `yarn start` command.

Please contact the repository owner for any necessary questions.
