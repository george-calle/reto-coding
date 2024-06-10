# Usa una imagen base oficial de Node.js
FROM node:14

# Establece el directorio de trabajo en el contenedor
WORKDIR /usr/src/app

# Copia el archivo package.json y package-lock.json
COPY package*.json ./

# Instala las dependencias del proyecto
RUN npm install

# Copia el resto de los archivos del proyecto
COPY . .

# Expon el puerto que usará la aplicación
EXPOSE 3000

# Define la variable de entorno para el puerto
ENV PORT=3000

# Comando para ejecutar la aplicación
CMD [ "npm", "start" ]