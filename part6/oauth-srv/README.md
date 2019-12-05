# Oauth Service

This is the Oauth service

Generated with

```
micro new oauth-srv --namespace=io.github.entere --alias=oauth --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: io.github.entere.srv.oauth
- Type: srv
- Alias: oauth

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./oauth-srv
```

Build a docker image
```
make docker
```