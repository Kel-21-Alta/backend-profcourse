## Prof Course


### Running project

```
    git clone https://github.com/Kel-21-Alta/backend-profcourse.git
    
    git checkout development
    
    # install library 
    go get
    
    # copy config.example.json to config.json
    mv config.example.json config.json
    
    # make your envaronment
    
    # Running project
    go run app/main.go
 
    # Development
    reflex -r '\.go' -s -- sh -c "go run app/main.go"
    
    # Running Test coverage
    go test ./business/... -coverprofile=cover.out && go tool cover -html=cover.out
```