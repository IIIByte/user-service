# user-Service (BACK-END)

project/
├── api/                  # Определения API (gRPC, REST, Swagger)
│   ├── proto/           # .proto файлы
│   │   └── user.proto
│   └── swagger/         # Сгенерированная Swagger-документация
├── cmd/                 # Точка входа для запуска приложения
│   └── server/         # Основной сервер (gRPC + REST)
│       └── main.go
├── internal/            # Внутренняя логика приложения (закрытая от внешнего мира)
│   ├── entity/         # Сущности (Entities) - бизнес-объекты
│   │   └── user.go
│   ├── usecase/        # Случаи использования (Use Cases) - бизнес-логика
│   │   └── user/
│   │       ├── interface.go  # Интерфейсы для репозиториев и сервисов
│   │       └── usecase.go    # Реализация бизнес-логики
│   ├── repository/     # Репозитории (Adapters) - работа с хранилищем данных
│   │   └── user/
│   │       └── repository.go
│   └── delivery/       # Доставка данных (Controllers) - адаптеры для gRPC/REST
│       └── grpc/
│           ├── handler.go    # Обработчики gRPC-запросов
│           └── server.go     # Настройка сервера
├── pkg/                 # Общие утилиты (доступные извне, если нужно)
│   └── logger/         # Пример: логгер
│       └── logger.go
├── config/              # Конфигурация приложения
│   └── config.yaml
└── Makefile             # Автоматизация задач