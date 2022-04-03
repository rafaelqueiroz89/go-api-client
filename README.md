
[![GoDoc](https://godoc.org/github.com/narqo/go-badge?status.svg)](https://godoc.org/github.com/narqo/go-badge)
[![tests](https://github.com/rafaelqueiroz89/form3-client-api/actions/workflows/docker-image-tests.yml/badge.svg?event=push)](https://github.com/rafaelqueiroz89/form3-client-api/actions/workflows/docker-image-tests.yml)

### FORM3 Client API Take Home Exercise. Name: Rafael Queiroz (GO noob)

FORM3 Client API that exposes the methods in he FORM3 API, see the API docs here:

## How to use the Client

There is an example on how to use the Client here.

First you need to import the desired API resource to use, this is done by referecing, for eg.: github.com/rafaelqueiroz89/form3-client-api/src/accounts

Then you need to create a variable to work with the {resource}Operator

`var accountService = accounts.AccountServiceOperator{}`

The resource operations will now be available, the interface exposes the methods you might want to use, currently the available operations are: Fetch, Create and Delete an Account.

### Fetch
`result, resp, err = accountService.Fetch("4c54ff77-8067-43a7-807f-da216d598ad4")`

### Create
`result, resp, err := accountService.Create(&accounts.AccountDataRequest{...})`

### Delete
`resp, err := accountService.Delete("4c54ff77-8067-43a7-807f-da216d598ad4", 0)`