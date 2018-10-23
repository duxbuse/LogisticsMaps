FROM golang as builder

WORKDIR /go/src/github.com/duxbuse/LogisticsMaps
COPY . /go/src/github.com/duxbuse/LogisticsMaps/

WORKDIR /go/src/github.com/duxbuse/LogisticsMaps/cmd
RUN CGO_ENABLED=0 GOOS=linux go install -a

#------------------------------------
FROM alpine:latest
#Expose host:container
EXPOSE 8080:8080
WORKDIR /root/
COPY --from=builder /go/bin/ .
CMD ["./cmd"]