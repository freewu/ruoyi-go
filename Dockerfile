FROM alpine
ADD rouyi-go-service /rouyi-go-service
ENTRYPOINT [ "/rouyi-go-service" ]