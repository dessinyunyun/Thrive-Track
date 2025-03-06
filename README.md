# How to running :

- in root project
- cmd: go run .

# How to Add Schema or Create Table Database :

- cmd: cd database
- cmd: go run -mod=mod entgo.io/ent/cmd/ent new TABLE_NAME

# Generate command from repository to database

- cd database
- go generate ./ent

# How to generate update create table

- cd database
- cmd: go generate ./...

# See Documentation :

- After run the app, you can visit to {HOST}/{YOUR_PREFIX}/documentation/index.html
