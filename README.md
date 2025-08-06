mod-ctf/
├── cmd/
│   └── server/
│       └── main.go  # Основной сервер на Go
├── internal/
│   ├── handlers/
│   │   └── handlers.go  # Обработчики HTTP для этапов CTF
│   ├── models/
│   │   └── models.go  # Модели данных (флаги, этапы)
│   └── utils/
│       └── utils.go  # Утилиты (например, генерация флагов)
├── pkg/
│   └── auth/
│       └── auth.go  # Простая аутентификация для участников
├── Dockerfile  # Для контейнеризации
├── kubernetes/
│   ├── deployment.yaml  # Kubernetes Deployment
│   ├── service.yaml     # Kubernetes Service
│   └── configmap.yaml   # ConfigMap для конфигурации (флаги, правила)
├── go.mod  # Go модуль
├── go.sum  # Go зависимости
├── README.md  # Инструкции
└── LICENSE  # MIT License (пример)
