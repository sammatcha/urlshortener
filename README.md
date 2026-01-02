## Description
This is a project to shorten urls.

The project uses the following technologies: <br>
- Go
- Vite + React + Tailwind
- PostgreSQL

### Setup
1. Download and Install [Go](https://go.dev/) <br>
   a. Create Go module<br>
   ```go mod init [module-path]```
2. Create directory for front-end<br>
   a. Scaffold Vite + React<br>
    ```
    npm create vite@latest my-vue-app -- --template react
    npm install
    npm run dev
    ```
    b. Install Tailwind CSS
    ```
    npm install -D tailwindcss postcss autoprefixer
    npx tailwindcss init
    ```
4. Database
   - Deploy postgres db: `docker compose up -d`
   - Stop postgres db: `docker compose down`

### Commands
`npm run dev` <br>
`go run cmd/main.go`

### Testing
shorten url:
```
curl -X POST -d "url=<insert-url-here>" http://localhost:8080/shorten
```

redirect:
```
curl http://localhost:8080/<short-url> 
```


### Roadmap
- [X] Build out simple UI
- [ ] Create spacing for shortened url output
- [ ] Work on error handling + validating urls for db
- [ ]  Create unit testing


