
# Number Classification API

## Overview

This API classifies a given number by checking if it is perfect, prime, armstrong, odd/ even, and returns a fact about it

## Features
- Check if a number is prime
- Check if a number is perfect
- Check if a number is an armstrong number
- Identify if a number is even or odd
- Calculate the digit sum of a number
- Retrieve a fun fact about the number

## API Endpoint
GET /api/classify-number?number=NUMBER

## Installation

### Prerequsites
- Go installed

### Clone Repository
```git clone https://github.com/girlincyberspace/HngStage1.git```
```cd HngStage1```

### Run the Server
```go run ./cmd/web```

### Test the API
```curl -X GET http://localhost:8080/api/classify-number?number=371```
