FROM golang as builder
#copy in all source files
WORKDIR /go/src/github.com/duxbuse/LogisticsMaps
COPY . /go/src/github.com/duxbuse/LogisticsMaps/

#run the go unit tests
RUN go test

WORKDIR /go/src/github.com/duxbuse/LogisticsMaps/cmd
#build and install source files into binary including c lib into the binary
RUN CGO_ENABLED=0 GOOS=linux go install -a
#run the go app tests
RUN go test

#------------------------------------
FROM alpine
#Expose host:container
EXPOSE 9000
WORKDIR /root/
#copy binary across
COPY --from=builder /go/bin/ /cmd
#copy html across
COPY views /views
WORKDIR /cmd/
#run server
CMD ["./cmd"]