# Golang Framework

## CQRS

### Repository handle
https://outcrawl.com/go-microservices-cqrs-docker
https://shijuvar.medium.com/building-microservices-with-event-sourcing-cqrs-in-go-using-grpc-nats-streaming-and-cockroachdb-983f650452aa

### RSocket
https://github.com/jjeffcaii

### MongoDB
https://github.com/AleksK1NG
https://github.com/avelino/awesome-go#database

### Example CQRS Financial
https://github.com/pperaltaisern	
https://github.com/shijuvar

### Golang init project
go mod init <name>
go mod init melody-io/core-es  
go mod init melody-io/midware-es  

go get github.com/modernice/goes/...@main
go mod tidy
go mod edit -replace melody-io/core-es/=core-es

# Manual Install Kafka
1. Install jdk 13
a. download from sunsdk and run installation, execute following:
b. source ~/.zshrc
c. check java version:
	java -version

Or openjdk

> brew cask
> brew tap adoptopenjdk/openjdk
> brew install adoptopenjdk13
> brew autoremove

2. Kafka
https://www.tutorialkart.com/apache-kafka/install-apache-kafka-on-mac/
https://kafka.apache.org/quickstart
https://david-yappeter.medium.com/golang-pass-by-value-vs-pass-by-reference-e48aac8b2716

3. Kafka UI
https://github.com/provectus/kafka-ui

## Build Kafka UI
./mvnw clean install -Pprod

## Run Kafka
$ sh bin/zookeeper-server-start.sh config/zookeeper.properties
$ sh bin/kafka-server-start.sh config/server.properties

## Run Nats
export GO111MODULE=on
export GOMODCACHE=~/golang/pkg/mod

sudo killall -HUP mDNSResponder

brew services stop mongodb-community
nats-streaming-server
kill $(lsof -ti:4222)

# Run Todo Examples
export MONGO_URL="mongodb://localhost:27017/eventstore"
export NATS_URL="nats://localhost:4222" 

# Installing multiple Go versions
You can install multiple Go versions on the same machine. For example, you might want to test your code on multiple Go versions. For a list of versions you can install this way, see the download page.

Note: To install using the method described here, you'll need to have git installed.

## To install additional Go versions, 
run the go install command, specifying the download location of the version you want to install. The following example illustrates with version 1.10.7:

$ go install golang.org/dl/go1.10.7@latest
$ go1.10.7 download
To run go commands with the newly-downloaded version, append the version number to the go command, as follows:

$ go1.10.7 version
go version go1.10.7 linux/amd64
When you have multiple versions installed, you can discover where each is installed, look at the version's GOROOT value. For example, run a command such as the following:

$ go1.10.7 env GOROOT
To uninstall a downloaded version, just remove the directory specified by its GOROOT environment variable and the goX.Y.Z binary.

## Uninstalling Go linux / macOS / FreeBSD
Delete the go directory.
This is usually /usr/local/go.

Remove the Go bin directory from your PATH environment variable.
Under Linux and FreeBSD, edit /etc/profile or $HOME/.profile. If you installed Go with the macOS package, remove the /etc/paths.d/go file

