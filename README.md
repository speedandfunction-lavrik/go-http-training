# HTTP and unit testing 

1. Create new repository called `go-http-training` and initialize go modules. Also copy paste `.github` directory from root of this project to add CI to the project.

2. You need to http API server with following endpoints using in memory storage (build those endpoints with unit tests in mind). Please add unit tests to each and every one of those endpoints (create a separate pull request for every API endpoint):

    * `GET` /users/list - returns list of users as an array
    * `POST` /users/create - create user
    * `PUT` /users/update - update user
    * `GET` /users/get?username=test - get single user
    
    Use following description of user for your API (this is only suggestion feel fre to modify the data structure as you wish):

    ```json
    {
      "username": "john_doe",
      "full_name": "John Doe",
      "age": 20,
      "gender": "male"
    }
    ```

    We have already created postman collection `Go_Training.postman_collection.json` so feel free to import that into your [Postman](https://www.postman.com/) environment to help you with testing.

3. Add CORS middleware to your application (don't forget to add unit tests to it).
