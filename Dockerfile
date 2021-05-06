FROM golang:1.16-alpine3.13 as build-image

WORKDIR /go/src
COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/twitter-lambda

FROM public.ecr.aws/lambda/go:1

COPY --from=build-image /go/bin/ /var/task/

CMD ["twitter-lambda"]