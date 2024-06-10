# Utilizar la imagen oficial de Golang como base
FROM golang:1.18-alpine

# Setear el directorio de trabajo
WORKDIR /app

# Copiar el archivo go.mod y go.sum al contenedor
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el código fuente al contenedor
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto 2000
EXPOSE 2000

# Setear variables de entorno (asegúrate de tener un archivo .env en el directorio)
COPY .env .env

# Comando para ejecutar la aplicación
CMD ["./main"]