# Go & Vue.js Web app 
___

<img width="1433" alt="Screenshot 2022-08-17 at 22 56 48" src="https://user-images.githubusercontent.com/19240229/185249471-49de37e6-8279-473b-bbfc-60ee7f69cc90.png">


This is a simple web application with a Go server/backend and a Vue.js frontend.

The app is designed to make it easy to use any storage, using Hexagonal architecture.
Organized based on containers, in order to provide a real working application for development/deployment.
Something more than "hello-world" but with the minimum of pre-reqs.

Currently ready for local build using Golang, Vue.js, MongoDB, Docker.

* The Frontend written in [Vue.js 3](https://vuejs.org/guide/quick-start.html)
* The Go component is a Go HTTP server based on [Chi](https://go-chi.io/#/README) routing

**Features**:
* View data
* Add to DB
* Delete object from list 
* Docker compose wrapped


https://user-images.githubusercontent.com/19240229/185249415-e0b50188-f27d-43a7-a53d-f1c49c96f05a.mov


## DONE
___
Front-end:

- [x] Input field
- [x] Add/remove button
- [x] Validation for non-empty input
- [x] Validation for deletion of non-existent item
- [x] Clearing input after adding/deleting the item
- [x] Comma separated output

Back-end:

- [x] Implemented POST end-point
- [x] Validation for non-empty input
- [x] Validation for existing data input
- [x] Implemented core logic for GET and DELETE end-points
- [x] MongoDB storage
- [x] Communication with MongoDB
- [x] Docker compose

## TODO
___
Front-end:
- [ ] Validation clean
- [x] ~~Fetching.~~
- [ ] Separate view list option to another component.
- [ ] CSV rendering.

Back-end:
- [ ] Fix validation
- [x] ~~Fix GET and DELETE end-points.~~
- [x] ~~Fix architecture, folder structures.~~
- [x] Apply clean code.
- [x] ~~Edit status codes view.~~
- [x] ~~Implement unified entrance.~~
- [x] ~~Add Makefile.~~
- [x] ~~Apply hexagonal architecture.~~

General:
- [ ] Add lintering to Makefile.
- [ ] Refactor REST.
- [ ] Add tests.
- [ ] Apply CI/CD.
- [ ] Add Swagger.

## Repo Structure
___
```
├── frontend            Root of the Vue.js project
│ └── src               Vue.js source code
└── server              Go backend server
  ├── cmd               Server main / exec
  ├── internal          Storage packages
  └── pkg               Supporting packages
```
## Building & Running Locally
___
### Pre-reqs
* Be using Linux, WSL or MACOS, with bash, make etc
* [Vue.js](https://vuejs.org/guide/quick-start.html) - run frontend  
* [Vite](https://vitejs.dev/) - tooling for frontend
* [Go 1.16+](https://golang.org/doc/install) - for running locally
* [Docker](https://docs.docker.com/get-docker/) - for running as a container, or image build
* [MongoDB Compass](https://www.mongodb.com/products/compass) - view collection data after applying changes 

Clone the project to any directory where you do development work

```
git clone https://github.com/moontary/email-viewer.git
```

### Makefile

Current version GNU Make file is provided to help with running and building locally.

```
up          Start Build
down        Stop an remove containers
logs        View logs
```

* The server will listen on port 9090 by default, change this by setting the `config/config.go` variable.

## Config

Environmental variables

Frontend:

* `VITE_URL` - Port for serving connection between containers. 

Backend:

* `Port` - Port to listen on (default: `9090`)
* `URI` - Port for DB to listen on (default: `27017`)
