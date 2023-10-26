FROM node:20

WORKDIR /app

COPY . .

RUN npm run build

CMD ["serve", "-s", "build", "-l", "3000"]
