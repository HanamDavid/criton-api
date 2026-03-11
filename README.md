# 🏛️ Criton-API

Sistema de backend robusto para la **logística de asistencia** empresarial, utilizando firmas digitales (PNG) y arquitectura limpia.

## 🎯 Propósito del Proyecto

Este sistema reemplaza las planillas físicas de asistencia por un flujo digital auditable. El MVP está diseñado para manejar ~100 usuarios con una infraestructura serverless de bajo costo.

## 🏗️ Arquitectura del Sistema

### 1. Diagrama C4 (Contenedores)

Este diagrama muestra cómo interactúa el binario de Go con los servicios de Google Cloud y el cliente final.

```mermaid
graph TD
    User((Empleado))
    Admin((Administrador))

    subgraph GCP ["Google Cloud Platform / Serverless"]
        API["Go Backend (Cloud Run)"]
        DB[("PostgreSQL (Cloud SQL)")]
        Bucket[("GCS Bucket (Firmas PNG)")]
    end

    User -- "Registra asistencia (JWT + PNG)" --> API
    Admin -- "Consulta reportes y logs" --> API
    
    API -- "CRUD Operaciones" --> DB
    API -- "Upload firmada" --> Bucket
```

```mermaid
sequenceDiagram
    participant U as Usuario
    participant A as API (Gin Handler)
    participant S as AttendanceService
    participant G as GCS Adapter
    participant R as Postgres Repo

    U->>A: POST /check-in (JWT + PNG)
    A->>A: Validar Token JWT
    A->>S: RegisterCheckIn(UserID, Image)
    S->>R: ¿Ya marcó hoy?
    R-->>S: No (OK)
    S->>G: UploadSignature(PNG)
    G-->>S: URL de la imagen
    S->>R: Save(Attendance Record)
    R-->>S: Success
    S-->>A: Domain Entity
    A-->>U: 201 Created (Success)
```

```mermaid
classDiagram
    class Attendance {
        +AttendanceID ID
        +UserID User
        +SignatureURL Signature
        +DateTime CreatedAt
        +String Location
    }
    class UserID {
        +String value
    }
    class AttendanceID {
        +String value
    }
    <<ValueObject>> UserID
    <<ValueObject>> AttendanceID
    Attendance *-- UserID
    Attendance *-- AttendanceID
```

### 2. Entidades

```mermaid
erDiagram
    USERS ||--o{ ATTENDANCE_LOGS : "registra"
    USERS ||--|| USER_STATS : "posee"
    
    USERS {
        uuid id PK
        string name
        string email UK
        string password_hash
        string role "admin | staff | employee"
        boolean is_active
        timestamp created_at
    }

    ATTENDANCE_LOGS {
        uuid id PK
        uuid user_id FK
        timestamp timestamp
        string signature_url "URL de GCS"
        inet ip_address
        string status "pending | verified | rejected"
        timestamp updated_at
    }

    USER_STATS {
        uuid user_id PK, FK
        int total_checkins
        float punctuality_rate
        int current_streak
        timestamp last_activity
    }
```

🛠️ Stack Tecnológico

    Lenguaje: Go 1.21+ (Fuerte énfasis en tipos y concurrencia).

    Framework: Gin Gonic para el ruteo HTTP.

    Arquitectura: Hexagonal (Ports & Adapters).

    Infraestructura: - Google Cloud Run (Compute)

        Google Cloud Storage (Firmas)

        PostgreSQL (Persistencia)

    Calidad: CI/CD con SonarQube "Zero Trust" quality gates.

🚀 Guía de Inicio Rápido
Requisitos

    Go instalado.

    Google Cloud SDK configurado.

    Instancia de Postgres (puedes usar el docker-compose.yml incluido).

Instalación

    Clonar el repo:
    Bash

    git clone [github.com/HanamDavid/criton-api](https://github.com/HanamDavid/criton-api)

    Instalar dependencias:
    Bash

    go mod tidy

    Configurar variables de entorno:
    Bash

    cp .env.example .env

📜 ADRs (Architecture Decision Records)

Las decisiones técnicas importantes están documentadas en docs/adr/:

    ADR-001: Arquitectura Hexagonal

---
