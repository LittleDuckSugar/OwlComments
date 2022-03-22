# OwlComments
> BERNARD Antoine - 23 mars 2022


## How does it works ?

By default the server will listen on port 3000

### Using Docker
Build docker image
```bash
docker build -t owlcomment .
```

Then run it
```bash
docker run -it --rm -p 8080:3000 owlcomment
```

**By doing this the API will be accessible throw port 8080**


### Without Docker
Simply go inside the API directory and launch main.go file
```bash
cd API
go run main.go
```
