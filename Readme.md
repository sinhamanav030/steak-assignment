# Task 
https://docs.google.com/document/d/1soF3puBzTJeArjWBdU2Wx0JXh5zuGYCG/edit?usp=sharing&ouid=101023644063452181190&rtpof=true&sd=true

# How to start server
Run \
``` go mod tidy``` \
```go run main.go```

# Sample api request
```curl -X POST http://localhost:8080/find-pairs -H "Content-Type: application/json" -d '{"numbers": [1, 2, 3, 4, "aa"], "target": 6}'```

