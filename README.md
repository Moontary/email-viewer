# email-viewer

This description shortly introduces how to run front-end part of app. Back-end handy version in progress.

## DONE 

Front-end:
- Add/remove component.
- Validation for non-empty input.
- Validation for deletion of non-existent item.
- Clearing input after adding/deliting the item.
- Comma seperated output.

Back-end:
- Implemented POST end-point.
- Validation for non-empty input.
- Validation for existed data input.
- Added status codes support.
- Implemented core logic for GET and DELETE end-points.
- [TO BE DELETED] MongoDB storage.
- [TO BE DELETED] Communication with MongoDB.
- Docker compose base. Added yaml config file.

## TODO

Front-end:
- Validation.
- Fetching.
- Separate view list option to another component.
- CSV rendering.

Back-end:
- Fix validation.
- Fix GET and DELETE end-points.
- Fix architecture, foldering.
- Apply clean code.
- Edit status codes view.
- Implement universal entrance.
 
## Config setup

#### Compile and Hot-Reload for Development

```sh
npm run dev
```

#### Compile and Minify for Production

```sh
npm run build
```

## Back-end patterns for testing

```sh
go run backend/main.go
```

```postman
Route: /
[IN PROGRESS]GET: /emails
POST: /email
[IN PROGRESS]DELETE: /{email}
```
