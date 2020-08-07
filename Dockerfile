# The base go-image
FROM golang:1.14-alpine

# Create a directory for the app
RUN mkdir /src

# Copy all files from the current directory to the app directory
COPY . /src

# Set working directory
WORKDIR /src

ENV GOPATH /

# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o /app/server .

# Run the server executable
CMD [ "/app/server" ]
