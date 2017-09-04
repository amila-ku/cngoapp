FROM golang:1.8
MAINTAINER Amila Kumaranayaka

#Set the working directry and copy application code
WORKDIR /go/src/app
COPY . .

# Set the PORT environment variable inside the container
ENV PORT 8080

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]