# Usamos la imagen oficial de Go
FROM golang:1.20-alpine

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /workspace

# Instalamos herramientas adicionales (opcional)
RUN apk add --no-cache \
    git \
    make \
    bash

# Configuramos el entorno de Go
ENV GO111MODULE=on
ENV GOPATH=/go
ENV PATH=$PATH:/go/bin

# Copiamos el código fuente al contenedor (opcional, si tienes código)
#COPY . .

# Expone el puerto 8080 (ajusta según tu aplicación)
EXPOSE 8080

# Comando por defecto (abre una terminal interactiva)
CMD ["/bin/sh"]