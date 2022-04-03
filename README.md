
[![GoDoc](https://godoc.org/github.com/narqo/go-badge?status.svg)](https://godoc.org/github.com/narqo/go-badge)
[![tests](https://github.com/rafaelqueiroz89/form3-client-api/actions/workflows/docker-image-tests.yml/badge.svg?event=push)](https://github.com/rafaelqueiroz89/form3-client-api/actions/workflows/docker-image-tests.yml)

### FORM3 Client API Take Home Exercise. Name: Rafael Queiroz (GO noob)

FORM3 Client API that exposes the methods in he FORM3 API, see the API docs here:

## How to use the Client

There is an example on how to use the Client here: https://github.com/rafaelqueiroz89/form3-client-api/tree/main/src/example.

First you need to import the desired API resource to use, this is done by referecing, for eg.: github.com/rafaelqueiroz89/form3-client-api/src/accounts

Then you need to create a variable to work with the {resource}Operator

`var accountService = accounts.AccountServiceOperator{}`

The resource operations will now be available, the interface exposes the methods you might want to use, currently the available operations are: Fetch, Create and Delete an Account. There is also a function to update the Base URL, the default for local development is http://localhost:8080 while in Docker it is: http://accountapi:8080 

### Fetch
`result, resp, err = accountService.Fetch("4c54ff77-8067-43a7-807f-da216d598ad4")`

### Create
`result, resp, err := accountService.Create(&accounts.AccountDataRequest{...})`

### Delete
`resp, err := accountService.Delete("4c54ff77-8067-43a7-807f-da216d598ad4", 0)`

## Next steps and considerations

Some points to consider:
- I created a workflow in Github Actions to spin up the containers and run the tests, this allows me to test the code without having all the environments setup in my local machine. The target agent is Linux and I code in Windows.
- We should add a Makefile to ease the process of CI and container spin up + integration tests
- I added integration tests which will trigger the local API and a few unit tests of things not fully tested in the integration tests
- I would use the Git flow structure to add more features to the client
- You will see two users commiting, both are me but one of them is my work github profile (apperently I forgot to logout while making the commits), at least this repo is private :) 
- I used Goland to create this project
- I used the builder pattern and fluent pattern to help me create the integration tests which made them easy to create
- I also used the table pattern to create many tests with different outcomes and assert on them
- Github actions are here: https://github.com/rafaelqueiroz89/form3-client-api/actions
