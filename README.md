# go-task


## Description
building a task managemnt system  using graphql and with a very good test coverage for
repository, usecase and graphql playground test.


### Technical tools used

- Golang
- Docker
- Postgres (included in docker)
- Migrations
- Graphql
- jwt
- testify
- mock

### Running

Setting up all containers

```console
$ make postgress
```

### Create and drop DB

Setting up all containers

```console
$ make createdb

$ make dropdb
```
### Migration

Migrating sql file into db

```console
$ make migrationup
```

Dropping tables 

```console
$ make migrationdown
```
Note: Run migration when docker is running

### Run Test

Run all the test in the project with

```console
$ make test
```


## Sections 

- ### Users
- ### Task
- ### Label


- ## User:

```console


mutation registerUser {
  createUser(
    input: {username: "emeka pas", email: "emeka233@gmail.com", passwordHash: "Password@"}
  ) {
    token
    expired_at
    user {
      id
      username
      email
      passwordHash
      created_at
      updated_at
    }
  }
}

mutation loginUser {
  loginUser(input: {email: "emeka233@gmail.com", password: "Password@"}) {
    token
    expired_at
    user {
      id
      username
      email
      passwordHash
      created_at
      updated_at
    }
  }
}

query getUserByID {
  getUserById(id: "a59ac614-96d9-4439-84dd-0ab2586ccd0a") {
    id
    username
    email
    created_at
    updated_at
  }
}

Header
{"Authorization":"bearer JWT_TOKEN_URL"}
```


- ### Tasks
```console
mutation createTask {
  createTask (
    input: {title:"title1", 
      description:"description", 
      priority:"high", 
      due_date:"2023-08-15 16:37:00.798196"}
  ) {
    Data{
      id
      title
      description
      due_date
      status
      priority
      user_id
      created_at
      updated_at
      
    }
  }
}


mutation updateTask {
  updateTask (
    input: {
      status:"in-progress"
      title:"title1", 
      id:"df51048f-6063-40f2-9ea2-5cf96f8bbfb5"
      description:"description", 
      priority:"high", 
      due_date:"2023-08-15 16:37:00.798196"}
  ) {
    Data{
      id
      title
      description
      due_date
      status
      priority
      user_id
      created_at
      updated_at
      
    }
  }
}

query getTaskById{
  getTaskById(id:"bab811c3-4ea6-4b74-b933-dfa760ab1be4"){
    Data{
       id
      title
      description
      due_date
      status
      priority
      user_id
      created_at
      updated_at
    }
  }
}

query fetchTasks{
  ListTask{
    Data{
       id
      title
      description
      due_date
      status
      priority
      user_id
      created_at
      updated_at
    }
  }
}

mutation deleteTaskById{
  deleteTask(id:"f6f2bc5f-3692-4d92-93a7-16a648180481")
}
```



- ### Labels
```console
mutation createLabel {
  createLabel (
    input: {
      name: "label 1"
    }
  ) {
    Data{
      id
     name
      created_at
      updated_at
      
    }
  }
}
```




### Tasks
- [x] Users api implementaion
- [x] Users test coverage for the implemented graphql
- [x] Tasks implementation
- [x] Tasks test coverage for the implemented graphql
- [x] Label implementaion
- [ ] write test for label graphql playground


