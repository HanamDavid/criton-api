ADR 0001: Arquitectura Hexagonal y Almacenamiento en GCS
Estado

Aceptado (Marzo 2026)
Contexto

El sistema Criton requiere gestionar registros de asistencia que incluyen firmas en formato PNG. Necesitamos una solución que:

    Sea altamente testeable sin depender de servicios externos (Base de Datos o Nube).

    Mantenga la base de datos ligera, evitando el almacenamiento de grandes volúmenes de datos binarios (Blobs/Base64).

    Permita el crecimiento del equipo técnico (Andrés, Hinara, María) sin generar conflictos de dependencias.

Decisión

Hemos decidido implementar los siguientes patrones arquitectónicos y de infraestructura:

1. Arquitectura Hexagonal (Ports & Adapters)

Dividiremos el código en tres capas estrictas:

    Domain: Contendrá las entidades (ej. Attendance) y Value Objects, libres de cualquier framework o librería externa.

    Application (Services): Orquestará los casos de uso (ej. RegisterCheckIn) utilizando únicamente interfaces (Puertos).

    Infrastructure (Adapters): Implementará los detalles técnicos como la conexión a PostgreSQL y el cliente de Google Cloud Storage.

2. Almacenamiento Externo de Binarios

Las firmas PNG no se guardarán en la base de datos relacional. En su lugar:

    Se utilizará un Bucket de Google Cloud Storage (GCS).

    La base de datos solo almacenará la SignatureURL (referencia al objeto en GCS).

3. Evitar "Primitive Obsession"

Utilizaremos tipos definidos (ej. type UserID string) en lugar de tipos primitivos para los identificadores, asegurando que el sistema sea auto-documentado y seguro en tiempo de compilación.
Consecuencias
Positivas

    Testabilidad: Podemos hacer mocks de los puertos de salida (Storage y DB) para probar la lógica de negocio en milisegundos.

    Escalabilidad: GCS maneja el almacenamiento de forma infinita y económica, mientras que Postgres se mantiene rápido para consultas de logs.

    Intercambiabilidad: Si en el futuro decidimos migrar de GCS a AWS S3, solo debemos crear un nuevo adaptador en la capa de infraestructura sin tocar el core del sistema.

Negativas

    Boilerplate: Requiere más archivos y carpetas iniciales en comparación con un diseño MVC tradicional.

    Curva de aprendizaje: El equipo debe entender el flujo de inversión de dependencias para no acoplar la lógica al driver de la base de datos.
