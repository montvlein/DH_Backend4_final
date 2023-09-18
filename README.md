# Desafio final Especialización en Back End 3 Grupo 8


## Tabla de Contenido


- [Configuración del archivo .env](#configuración-del-archivo-env)
- [Montar Contenedor Docker](#montar-contenedor-docker)
- [Documentación Swagger](#documentación-swagger)
- [Autenticación](#autenticación)


## Configuración del archivo .env

Antes de ejecutar el proyecto, asegúrate de configurar el archivo .env en el directorio principal. El archivo .env debe contener las siguientes variables de entorno:


```bash
env=local
SECRET_KEY=testAdmin
HOST=localhost:8080
POSTGRES_PASSWORD=grupo8
```

## Montar Contenedor Docker

Para ejecutar el proyecto, puedes utilizar Docker Compose. Asegúrate de tener Docker Compose instalado en tu sistema. Luego, sigue estos pasos:

Navega al directorio principal de tu proyecto donde se encuentra el archivo docker-compose.yml.

Ejecuta el siguiente comando para montar y ejecutar el contenedor Docker:

```bash
docker-compose up -d
```

Esto iniciará el contenedor en segundo plano (-d) según la configuración definida en docker-compose.yml.

Para detener el contenedor cuando hayas terminado, utiliza el siguiente comando:

```bash
docker-compose down
```

Esto detendrá y eliminará el contenedor Docker.


## Documentación Swagger
La documentación de Swagger para este proyecto está disponible en la siguiente URL:

```bash
http://localhost:8080/docs/index.html
```
Puedes acceder a esta URL en tu navegador para explorar la documentación de la API. La documentación de Swagger contiene información detallada sobre todos los endpoints existentes en el proyecto, incluyendo descripciones, métodos HTTP y parámetros.

![imagen](https://github.com/montvlein/DH_Backend4_final/assets/44759635/6f74036a-83a2-4760-a6d4-b3a87b75098c)


## Autenticación

Para autenticarse en los endpoints de la API, utiliza las siguientes credenciales en el encabezado (header) de las solicitudes:

```bash
PRIVATE-KEY: testAdmin
PUBLIC-KEY: localAdmin
``` 

![image](https://github.com/montvlein/DH_Backend4_final/assets/75510452/024e3c03-170b-4427-b91e-feefb0bdf02a)
    
Asegúrate de incluir estas credenciales correctamente en tus solicitudes para acceder a los recursos protegidos de la API.






