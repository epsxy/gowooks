# gowooks
[![Go report](https://goreportcard.com/badge/github.com/epsxy/gowooks)](https://goreportcard.com/badge/github.com/epsxy/gowooks)
[![Build Status](https://travis-ci.org/epsxy/gowooks.svg?branch=master)](https://travis-ci.org/epsxy/gowooks)
[![Coverage Status](https://coveralls.io/repos/github/epsxy/gowooks/badge.svg?branch=master)](https://coveralls.io/github/epsxy/gowooks?branch=master)

A lightweight secure git webhook server

# Quickstart

You can easily build and run using predefined makefile rules. The `.env` file contains variables and secrets for your server.

```bash
# copy .env file. 
# You have to update it with your own values
cp DEV.env .env

# test
make test

# build
make build

# run
make run
```

# Env file

Below the description of all parameters

- `SECRET`: contains the secret you have set up in your Git server
- `PORT`: Contains the port the server will be listening on
- `ROUTE`: The route the server must listening your hook on

# Licence

The MIT Licence