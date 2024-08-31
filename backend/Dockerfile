 #### Stage 1
 FROM golang:1.21.4-alpine AS build
 RUN mkdir /app
 COPY . /app
 WORKDIR /app
 RUN go build -o main ./cmd/app/
 RUN chmod +x /app/main

 #### Stage 2
 FROM scratch
 COPY --from=build /app/main /main
 COPY --from=build /app/fullchain.pem /fullchain.pem
 COPY --from=build /app/privkey.pem /privkey.pem
 ENTRYPOINT [ "/main" ]
