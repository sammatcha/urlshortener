# URL Shortener
This is a project to shorten urls

## Tech Stack
Go · React · TypeScript · Vite · Tailwind CSS · PostgreSQL · Docker

## Features
- Shorten long URLs
- Store URLs in PostgreSQL
- Redirect via generated short URLs

## Getting Started

Start the Database:
Deploy PostgreSQL:
   ```
   docker compose up -d
   ```
 Stop PostgreSQL:
 ```
   docker compose down
 ```

Run the Backend:
```
go run cmd/main.go
```
Run the Frontend:
```
cd web
npm install
npm run dev
```

### Testing
shorten url:
```
curl -X POST -d "url=<insert-url-here>" http://localhost:8080/shorten
```

redirect:
```
curl http://localhost:8080/<short-url> 
```


