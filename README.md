
# Introduction

This API allows managing invoices via a dockerised API. 
It has a set of 17 users for which one may:
- create invoices via `/invoice`
- pay invoices via `/transaction` (the amount must match exactly the invoice's one)
- retrieve the list of all users via `/users`, or a single user (by id) via `/user`
- retrieve the list of all invoices via `/invoices`, or a single user (by id) via `/user-invoices`

It was implemented in Golang (1.19) and uses the official docker image. 
The database uses postgres:14.2-alpine.

# Environment

The environment variables are retrieved from the .env file. 
Before running the API, you should set one with the following command.
By default, the API runs on the port 4000 and the database on the port 5433.

```bash
cp .env.sample .env
```

# How to run it

To run the API, first run ```build-docker-image``` to build the docker image.
Then run ```make run``` to start the server on Docker.
API logs will display the responses to requests.

```make build``` allows you to rebuild the binary directly in the container once it has started. 
```make build-and-restart``` will do the same thing and restart the container, allowing you to test quickly any change brought to the code.

Requests examples:

```bash
curl -X GET -i http://localhost:4000/users

curl -X GET -i http://localhost:4000/invoices

curl -X GET -i http://localhost:4000/user-invoices -d "12"

```

# Next steps

I did not implement an authentication here because it seemed to me that to make sense, an authentification should come with a permission system (or at least with roles).

Regarding the existing endpoints:
	/users: only (super) admins should access all users
	/invoice: each user should only be able to create his own invoices
	/transaction: this operation should only be triggered by a user's client

This should be the next thing to implement.

Here are other things that would be needed or useful:
- a whole CRUD regarding the users would be needed
- it could be useful to store the freelances' customers in DB, because they are the entities that emit the transactions
- it may be useful to allow the users to withdraw money
- transaction references are sent in input, but not used. If deemed useful, they should be added to the data model.

On a more technical side:
- handling invoices statuses as tinyints in database could give us more flexibility
- real health checks would be useful


# Tests

The endpoints are very simple and quite centred around database operations. 
This led me to implement integration tests, and no unitary tests.
These integrations tests can be run with Newman in command line via the command ```make tests``` (only once the server is running).
This requires a local version of Newman, which can be installed with NPM: ``` npm install -g newman ```

They can also be visualised and run with Postman, which offers a clear interface.
Finally: they do not rely on specific data from the database and only require that one user exists in the table `users`.

