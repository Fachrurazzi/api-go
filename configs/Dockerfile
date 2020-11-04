# start from the latest golang base image
FROM golang:alpine

# Add Maintaner info
LABEL maintainer="Muhammad <*****@gmail.com>"
LABEL build_date="3-June-2020"
RUN apk update && apk add --no-cache gcc && apk add --no-cache libc-dev && apk add --no-cache --update make

# Set the current working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. they will be cached of the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the WORKDIR inisde the container
COPY . .

# Build the Go app
RUN go build .

# Exporse port 3000 or 8000 to the outisde world
EXPOSE 3000
#RUN make -C scripts db_migrate
#RUN make -C scripts db_migrate_test
#RUN make -C scripts db_seed
#RUN make -c scripts db_seed_test

# Command to run the executable
CMD ["make", "-C", "scripts", "prod" ]
