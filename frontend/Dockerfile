# frontend/Dockerfile
FROM node:22-alpine

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

# Copy all source files
COPY . .

RUN npm install -g typescript

EXPOSE 3000

CMD ["npm", "start"]


