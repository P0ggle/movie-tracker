# frontend/Dockerfile
FROM node:22-alpine

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . .

RUN npm install -g typescript

# Run the development server
CMD ["npm", "start"]

