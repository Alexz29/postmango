# Free local mock server for Postman

This is mock server written in Golang that works with postman mock server json file

[![Build Status](https://travis-ci.org/agilov/postmango.svg)](https://travis-ci.org/agilov/postmango)

## Installation

```bash
wget -N https://github.com/agilov/postmango/releases/download/v0.0.1/postmango-0.0.1-linux-amd64.zip -P /tmp
unzip /tmp/postmango-0.0.1-linux-amd64.zip -d /tmp
mv /tmp/postmango-0.0.1-linux-amd64/postmango /path/for/postmango
chmod +x /path/for/postmango

rm /tmp/postmango-0.0.1-linux-amd64.zip
rm -rf /tmp/postmango-0.0.1-linux-amd64
```

## Usage example

```bash
./postmango -f ./path/to/mock/server/file.json -p 8888 -h localhost
```

## Params
-f Path to file with postman exported mock server
-p Listen on port
-h Listen on host

By default port is 8080 host - localhost and filename - ./server.json