
# Nouns
A service for storing people, places and things that interest you.

## NOTE: Just Playing with Go-Kit. For the real source, see: [https://github.com/go-kit](https://github.com/go-kit)

## Getting Started

1. Navigate to the Nouns project

```bash
cd ${GOPATH}/src/github.com/mackdoyle/nouns
```

2. Build the latest source

```bash
go build
```

3. In a second terminal, start three web servers and then a fourth that proxies to the first three via port 9000

```bash
nouns -listen=:8001 & nouns -listen=:8002 & nouns -listen=:8003 & nouns -listen=:9000 -proxy=localhost:8001,localhost:8002,localhost:8003
```

## Posting Data

To post to the noun service, you can use cURL and pass in an example body.

```bash
curl -X POST -d @schemas/examples/noun.json localhost:9000/noun
```

## Quitting
Once you are done testing, you can stop all servers using this nifty command

```bash
kill -9  $(ps aux | grep listen | grep -v grep | awk '{print $2}')
```
