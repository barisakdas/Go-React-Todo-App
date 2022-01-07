# React Todo App UI

[![go-github docker-image (latest SemVer)](https://img.shields.io/github/v/release/google/go-github?sort=semver)](https://hub.docker.com/repository/docker/barisakdas/react-todo-ui)

This application is a ui project with react hooks for the todo app application. In it, the necessary libraries are used to design and call APIs in a simple way. It is designed in a way that is separated from each other with its component structures.

## Installation

The program can be run easily by downloading the publicly published image on `hub.docker.com`:

```bash
docker pull barisakdas/react-todo-ui
```

If you want to run the codes by getting the codes on github, the project can be reached by pulling the repository as follows.

```go
git clone https://github.com/barisakdas/React-Todo-UI
```

## Usage

In order for the application to work, the api application written for it must be in working condition. Data is received via inputs through the application and sent to the API side.

```bash
curl http://localhost:3000/
```

When it is desired to run with the codes on Github, the `npm install` command should be run first to get the node packages. This command will include the required packages in the project.Then you can run the project with the `yarn start` command.

Please contact the repository owner for any necessary questions.
