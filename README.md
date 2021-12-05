# API Ilúvatar - Proyecto Kümelen

## Descripción del Proyecto

"Kümelen" nace desde la necesidad de implementar, con ayuda de elementos tecnológicos, nuevos canales de comunicación
que le permitan a los estudiantes de la UTEM conocer e informarse sobre los servicios estudiantiles que pueden apoyar
en las dimensiones extraacadémicas de su vida universitaria.

## Descripción del artefacto

El artefacto "Ilúvatar" es una API que realiza funciones de autenticación, consumiendo los recursos que provee la API
"Mi UTEM".

## Variables de entorno

Las variables de entorno necesarias para la ejecución del artefacto son:

- PORT
- APP_VERSION
- ENV
- MI_UTEM_API_HOST
- DB_SERVER
- DB_PORT
- DB_NAME
- DB_USER
- DB_PASSWORD
- JWT_TOKEN_SEED
- ADMIN_ACCOUNT_EMAIL
- ADMIN_ACCOUNT_PASS
- CLOUD_MESSAGE_API_HOST
- CLOUD_MESSAGE_API_KEY

## Entorno de compilación/ejecución

Para compilar el código es necesario configurar un ambiente de Go en su versión 1.15

La ejecución del mismo es totalmente agnóstica, dado que de la compilación es generado un archivo binario ejecutable

## Generación/Visualización de documentación

Para la generación/actualización de la documentación del proyecto se requiere ejecutar el siguiente comando:

```console
foo@bar:~$ swag init -g src/main.go
```

Luego, para su visualización, es necesario entrar a la siguiente ruta:

http://servidor.com/swagger/index.html

### Nombre en clave

El nombre en clave utilizado para este artefacto, "Ilúvatar", viene de la mitología creada por el autor J.R.R. Tolkien,
donde Ilúvatar representa al ente creador de todo el universo donde se llevan a cabo las aventuras de libros como "El
Hobbit" y "El Señor de los Anillos" [Ver Más](https://es.wikipedia.org/wiki/Il%C3%BAvatar)