# golang-proj-REST
![GCI Badge](https://img.shields.io/badge/Google%20Code%20In-JBoss%20Community-red?style=flatr&labelColor=fdb900)

GoLang Project to support REST endpoints GCI task

The Server runs on port 8080 and returns JSON or text


### To view all data use :

localhost:8080/get/all



### To select by id or use:

localhost:8080/get/{id}



### To select by name(exact match) use:

localhost:8080/get/name/{name}



### To add new data by cURL or a Form (HTTP method must be POST) use :

localhost:8080/post/new



### To delete data from the database by id (HTTP Method must be DELETE) use:

### The data should be sent with cURL or a form or a tool like postman

localhost:8080/delete



### To update data: name,price,time use any one of: 

### The data should be sent with cURL or a form or a tool like postman

localhost:8080/update/{id}/name

localhost:8080/update/{id}/price

localhost:8080/update/{id}/time



#### Google Code-in and the Google Code-in logo are trademarks of Google Inc.
