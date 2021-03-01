FROM golang:1.14.6

# Set the Current Working Directory inside the container
RUN mkdir ./app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . ./app

WORKDIR ./app

# Download all the dependencies
RUN go get -d -v

# Install the package
RUN go install -v

ENV CONFIG=config.yaml

# This container exposes port 80 to the outside world
EXPOSE 80

# Run the executable
CMD ["logo-user-management"]