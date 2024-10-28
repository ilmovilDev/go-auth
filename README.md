project-root/
├── cmd/                // Punto de entrada de aplicaciones
│   └── app/            // Aplicación principal
│       └── main.go     // main de la aplicación
├── configs/            // Archivos de configuración y carga de Envaronment
├── internal/           // Lógica de negocio privada a la aplicación
│   ├── domain/         // Modelos de dominio o entidades
│   ├── repository/     // Interfaz de acceso a datos (p. ej., SQL, MongoDB)
│   ├── service/        // Lógica de negocio
│   └── handler/        // Manejadores de HTTP o controladores
├── pkg/                // Paquetes reutilizables y utilidades
│   └── logger/         // Logger u otras utilidades
├── migrations/         // Archivos SQL de migración de base de datos
├── docs/               // Documentación (Swagger, README)
├── go.mod              // Declaración de módulos y dependencias
└── go.sum              // Checksum de módulos