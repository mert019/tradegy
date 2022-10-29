# docker build  --target tradegy-frontend . -t tradegy-frontend:latest
# docker build  --target tradegy-api . -t tradegy-api:latest


FROM golang:1.19.2-alpine3.15 AS tradegy-api
WORKDIR /app
COPY ./go-backend/go.mod ./
COPY ./go-backend/go.sum ./
RUN go mod download
COPY ./go-backend ./
RUN go build -o ./go-app ./cmd
EXPOSE 8080
CMD [ "./go-app" ]


FROM node:17-alpine AS tradegy-frontend
WORKDIR /app
COPY ./react-frontend/package.json .
RUN npm install
COPY ./react-frontend .
EXPOSE 3000
CMD ["npm", "start"]
