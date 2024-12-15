# passwordless-auth-server

project-root/
│
├── cmd/
│ └── server/
│ └── main.go # Entry point of the application
│
├── internal/
│ ├── handler/ # HTTP request handlers
│ │ ├── user_handler.go
│ │ └── product_handler.go
│ │
│ ├── service/ # Business logic layer
│ │ ├── user_service.go
│ │ └── product_service.go
│ │
│ ├── repository/ # Data access layer
│ │ ├── user_repository.go
│ │ └── product_repository.go
│ │
│ ├── model/ # Data models/structs
│ │ ├── user.go
│ │ └── product.go
│ │
│ └── middleware/ # Custom middleware
│ ├── auth.go
│ └── logging.go
│
├── config/ # Configuration files
│ └── config.yaml
│
├── pkg/ # Packages that can be imported by external projects
│ └── utils/
│ └── validation.go
│
├── migrations/ # Database migration files
│
├── scripts/ # Utility scripts
│
├── go.mod
├── go.sum
└── README.md
