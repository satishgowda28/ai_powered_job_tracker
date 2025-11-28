# ai_powered_job_tracker

A platform for job seekers to track applications, get AI insights, and analyze their job search performance — built with Go (backend) and React + TypeScript (frontend).

```tree
backend/
│
├── cmd/
│   └── server/            # main.go entry
│
├── db/
│   └── generated/            # type safe generated sql
│   └── migration/            # goose migrations schema
│   └── queries/              # sqlc sql queries
├── internal/
│   ├── config/            # env, config loader
│   ├── database/          # DB connection, migrations
│   ├── models/            # ORM models
│   ├── repositories/      # DB access layer
│   ├── services/          # business logic
│   ├── handlers/          # http handlers (controllers)
│   ├── middleware/        # JWT, logging, cors
│   └── routes/            # route definitions
│
├── pkg/
│   └── utils/             # helper functions (hash, jwt, etc.)
│
└── go.mod
```
