# guardianlink
There is two nodes nodea and nodeb sync together if any one create or update


# Distributed System with NodeA and NodeB Service


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>



### Built Withss

These framework/Libraries and tools are require to build Nodaa and Nodeb servicess

* ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
* ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
* ![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)



<!-- GETTING STARTED -->
## Getting Started

Given below prerequisite require to install yours system.

### Application 

There are atleast two nodes commincate with each other if any one node date create other data will create and update if any one node fails rollback from both nodes

### Architecture Diagram :
[diagram](https://excalidraw.com/#json=iPBq2dBHbbDu3LoOtBenT,i5MiNiIGzQ84ZP-aKnh-Aw)

![alt text](image-1.png)

#### Auth Secuirty

##### SINGUP  NODE A

```
 curl --location 'http://localhost:8080/api/signup' \
--header 'Content-Type: application/json' \
--data-raw '{"username": "Vivekanand1235", "password": "Password@12"}'  
    
```

##### SIGNIN NODE A

```
curl --location 'http://localhost:8080/api/signin' \
--header 'Content-Type: application/json' \
--data-raw '{"username": "Vivekanand1235", "password": "Password@12"}'

```

#### Node A 

##### CREATE BLOCK
```
curl --location 'http://localhost:8080/api/blocks' \
--header 'Content-Type: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlZpdmVrYW5hbmQxMjM1IiwiZXhwIjoxNzEyNzUwNzQ2fQ.NrV0ihvKa2RD7Lhc09eMfn93GUIPJGHZDGbz0nlkdXI' \
--data '{
  "height":99,
  "hash": "",
  "parent": "gaurdianlink",
  "status":"inprogress",
  "encoded": "YWJjZA=="  
}
'

```

##### GET BLOCK RECORDS

```
curl --location 'http://localhost:8080/api/blocks' \
--header 'Content-Type: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlZpdmVrYW5hbmQxMjM1IiwiZXhwIjoxNzEyNzUwNzQ2fQ.NrV0ihvKa2RD7Lhc09eMfn93GUIPJGHZDGbz0nlkdXI' \
--data '{
  "height":99,
  "hash": "",
  "parent": "gaurdianlink",
  "status":"inprogress",
  "encoded": "YWJjZA=="  
}
'
```
##### UPDATE BLOCK

```
curl --location --request PUT 'http://localhost:8080/api/blocks/2' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzEyNTc2MDcxfQ.r9QnU7Mux22Bt6ia4liAcyiFSIoZErJXGcieFkcHLN4' \
--header 'Content-Type: application/json' \
--data '{
           "height": 2,
           "hash": "Hash something ",
           "parent": "guardian link",
           "status":"inprogress"
           "encoded": "YWJjZA=="
         }'

```


##### SINGUP  NODE B

```
 curl --location 'http://localhost:8081/api/signup' \
--header 'Content-Type: application/json' \
--data-raw '{"username": "Vivekanand1235", "password": "Password@12"}'  
    
```

##### SIGNIN NODE B

```
curl --location 'http://localhost:8081/api/signin' \
--header 'Content-Type: application/json' \
--data-raw '{"username": "Vivekanand1235", "password": "Password@12"}'

```

#### Node B 

##### CREATE BLOCK
```
curl --location 'http://localhost:8081/api/blocks' \
--header 'Content-Type: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlZpdmVrYW5hbmQxMjM1IiwiZXhwIjoxNzEyNzUwNzQ2fQ.NrV0ihvKa2RD7Lhc09eMfn93GUIPJGHZDGbz0nlkdXI' \
--data '{
  "height":99,
  "hash": "",
  "parent": "gaurdianlink",
  "status":"inprogress",
  "encoded": "YWJjZA=="  
}
'

```

##### GET BLOCK RECORDS

```
curl --location 'http://localhost:8081/api/blocks' \
--header 'Content-Type: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlZpdmVrYW5hbmQxMjM1IiwiZXhwIjoxNzEyNzUwNzQ2fQ.NrV0ihvKa2RD7Lhc09eMfn93GUIPJGHZDGbz0nlkdXI' \
--data '{
  "height":99,
  "hash": "",
  "parent": "gaurdianlink",
  "status":"inprogress",
  "encoded": "YWJjZA=="  
}
'
```
##### UPDATE BLOCK

```
curl --location --request PUT 'http://localhost:8081/api/blocks/2' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzEyNTc2MDcxfQ.r9QnU7Mux22Bt6ia4liAcyiFSIoZErJXGcieFkcHLN4' \
--header 'Content-Type: application/json' \
--data '{
           "height": 2,
           "hash": "Hash something ",
           "parent": "guardian link",
           "status":"inprogress"
           "encoded": "YWJjZA=="
         }'

```




### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Golang
  ```sh
  Linux:
  
  1. $ sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz
  2. export PATH=$PATH:/usr/local/go/bin
  
  MacOs:
  brew install go@1.21
   
  ```
* Windows
  [Golang Installation](https://go.dev/doc/install)

* Docker
 ```
  brew cask install docker
  ```

### Installation


### Run Application Node A
 There is three environment 
 ### Manual

##### Local Environment

 ```
go run nodea/main.go -config config/local
 ```
 ##### Non prod Environment

 ```
    go run nodea/main.go -config config/non-prod
 ```
 ##### Prod Environment

```
    go run nodea/main.go -config config/prod
 ```

 ### Run Application Node B
 There is three environment 
 ### Manual

##### Local Environment

 ```
go run nodea/main.go -config config/local
 ```
 ##### Non prod Environment

 ```
    go run nodea/main.go -config config/non-prod
 ```
 ##### Prod Environment

```
    go run nodea/main.go -config config/prod
 ```

### Docker

   Which environment you want to run application use environment according yours requirement

   In docker-compose file change the 
   ```
    environment:
      - CONFIG_PATH=config/local or config/non-prod or config/prod
   ```

  ```
  docker-compose up 
```
## Swagger

  ```sh
  http://localhost:8080/swagger/index.html
   ```

<!-- USAGE EXAMPLES -->
## Usage

Integration with third parties tools

