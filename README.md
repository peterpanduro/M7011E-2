# Dependencies

## Node

[Node website](https://nodejs.org/en/)

## Terraform

~~[Terraform website](https://www.terraform.io/https://)~~ (_Not yet working_)

## Docker

~~[Docker website](https://www.docker.com/)~~ (_Not yet working_)

# Build the project

_Not yet implemented_

## Build individual parts

### Services

(_Guide so far only for Node services_)

1. `cd` to the service you want to build.
2. run `make proto` to generate protobuf files. (Also runs `npm install` since proto generation is a npm dependency).
3. `npm run dev` to run service in developer mode (with hot reload).
