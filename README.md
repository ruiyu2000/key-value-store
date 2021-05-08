## Server:

```
cd server/cmd/
go build -o server
./server
```

OR

```
docker build . -t server
docker run -p 1337:1337 server
```

## Client:

```
cd client/
go build

./client http://localhost:1337 set norway oslo
./client http://localhost:1337 get norway
./client http://localhost:1337 delete norway
```
