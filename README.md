## How to run
### Run all services and graphql playground
- System: linux
- Command:
```
make up
```
- Playground will be host at: *localhost:8081*
- Example queries' written in playground-examples

### Run single service
- Supported profile: flight, user, customer, graphql, booking
- Example for user service:
```
make upProfile Profile="user"
```

## Planned feature
- Implement transaction.

## Database schema
[Schema Link](https://dbdiagram.io/embed/64547112dca9fb07c48a6dfc)