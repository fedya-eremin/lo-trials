## Installation
```bash
git clone https://github.com/fedya-eremin/lo-trials
cd lo-trials
```

## Build
```bash
go build -ldflags "-s -w" -o main cmd/main/main.go
```

## Startup
```bash
./main
```

## Info
Server listens for connections at :8001. Avalaible endpoints are
- GET /tasks?status=active (query param is optional, if omitted, returns all entries)
- GET /tasks/{id} (id uint64)
- POST /tasks (deadline attr is iso8601 offset-aware with Z at end)

Example Curl requests
```bash
# GET /tasks?status=...
curl localhost:8001/tasks?status=active

# GET /tasks/{id}
curl localhost:8001/tasks/1

# POST /tasks
curl -X POST localhost:8001/tasks -d '{"name":"test", "description":"test", "status":"active", "assignee_id":1, "deadline":"2025-08-14T12:36:24Z"}'
```

Checklist:
- [x] Only standard library
- [x] Implemented in-mem storage
- [x] Async logger
- [x] Clean architecture (repo-service-transport model)
- [x] Graceful shutdown (handle os.Interrup & shut goroutines)
