# go-chi-openapi-preset
A quickstart for setting up an openapi driven project in go with chi

It is opinionated towards [hexagonal architecture](https://medium.com/ssense-tech/hexagonal-architecture-there-are-always-two-sides-to-every-story-bc0780ed7d9c).

I mainly use it for consistency, as a quick jump off point for small side projects.

It offers a starting point with the things any server should probably have:
 - logging (with zerolog in this case)
 - configuration (using yaml in this case)
 - a sensible dockerfile
 - an example endpoint

Non-Goals are features beyond these basic functions (e.g. database connections, actual business logic)

After forking:
 - change the package name and imports to reflect your repo/project name
 - implement your own service instead of the example service (or rename the example service)
 - start building :)

### Development:

Sync the mod how you normally would.

(Once): Make sure oapi-codegen is installed
````shell
go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
````

#### Generate Stubs from OpenAPI-Specs:
> 🛈 To be executed after each change to the spec (and each checkout)

````shell
go tool oapi-codegen -config ./api-definition/oapi-codegen.yml ./api-definition/openapi.yaml
````

### Contribution
Feel free to fork and build derrived templates and projects.

PRs will be accepted if they fix a problem with the existing setup. 

PRs for additional features will be closed, as they are probably not in the spirit of this template.
