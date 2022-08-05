# Email-viewer

This description shortly introduces how to run the front-end part of app. Back-end handy version in progress.

## DONE

Front-end:

- [x] Input field.
- [x] Add/remove button.
- [x] Validation for non-empty input.
- [x] Validation for deletion of non-existent item.
- [x] Clearing input after adding/deleting the item.
- [x] Comma separated output.

Back-end:

- [x] Implemented POST end-point.
- [x] Validation for non-empty input.
- [x] Validation for existing data input.
- [x] Implemented core logic for GET and DELETE end-points.
- [x] [TO BE DELETED, CHANGE TO CSV] MongoDB storage.  
- [x] [TO BE DELETED, CHANGE TO CSV] Communication with MongoDB.
- [x] Docker compose base. Testing required.

## TODO

Front-end:
- [ ] Validation cleaning.
- [ ] ~~Fetching.~~
- [ ] Separate view list option to another component.
- [ ] CSV rendering.

Back-end:
- [ ] Fix validation.
- [x] ~~Fix GET and DELETE end-points.~~
- [x] ~~Fix architecture, folder structures.~~
- [ ] Apply clean code.
- [ ] Edit status codes view.
- [ ] ~~Implement unified entrance.~~
- [ ] Add Makefile.

## RUN Front-end

#### Compile and Hot-Reload for Development

```sh 
npm run dev 
```  

#### Compile and Minify for Production

```sh 
npm run build 
```   

## RUN back-end

```sh 
go run backend/main.go 
```  

## End-points for testing

```postman 
Route: / 
GET: /emails 
POST: /email 
DELETE: /{email} 
```  
