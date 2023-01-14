# todo-backend

## Todo

- [x] tasks on memory
- [ ] .env configuration
- [ ] Build as a Container via buildpacks
- [ ] Build as a container image
- [ ] Logging - zap
- [ ] disable linter structcheck due to warninng
- [ ] lint with golangci-lint
- [ ] go test
- [ ] int test - user
- [ ] int test - task

## Overview

- Onion Architecture

## libs
- echo

## Manual Test

### Create User
```
curl -XPOST localhost:1323/api/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"hogename"}'
```

## Usage

```
make
```
- go test
- go run

