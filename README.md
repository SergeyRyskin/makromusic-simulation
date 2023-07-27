# Makromusic Relationships Structure Simulation

![Makromusic Relationships Structure Simulation](/img/makromusic_simulation.png)


## Introduction of Project 
The aim of this project is to simulate the relationship structure that exists in macromusic. In this simulation: 

- Accepting the Relationship
- Ignoring the Relationship
- Creating the Relationship

There are such cases. Our aim is to develop an application using the technologies and desired ones given in the ```Technology-Tool-Stack``` section below.

# Techstack of the Project

## Project Structure
User information and Users Auth information will be kept in SQL table. **REST Endpoints** will be as follows: 

- *Relationship Acceptance*
- *Relation Ignoring*
- *Relationship Creation*

We should have a *gRPC* module that outputs the relationship information. Generally the *REST API* will be calling **gRPC**. After making the necessary changes in the Database, **gRPC** will send the information ```("new_match","match_accepted","match_ignored")``` to the topics in **Kafka**. Consumers listening to Kafka topics will simply log this data.
**Auth must be done on the REST API side.** For this, as you remember, we created a Users Table. After the data is pulled from *PostgresSQL*, it should be cached in **REDIS**. There should also be a rate limit for each *ENDPOINT*, which must be handled with **Redis**.

In short: It should be a code that creates a relationship, and if certain users don't have a relationship or ignore it, they should call each other gRPC to create a relationship.
## Project File Structure

## Project Architecture

![Project Architecture](/img/simulation_arch.png)

## Technology-Tool-Stack

### Technologys and Tools

- Golang [More information about Golang](https://go.dev/)
- Postgres [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- Redis [Redis Golang documentation](https://redis.io/docs/clients/go/)
- Kafka [Confluent Kafka Golang documentation](https://docs.confluent.io/kafka-clients/go/current/overview.html)
- gRPC [gRPC Golang documentation](https://grpc.io/docs/languages/go/quickstart/)
- Fiber [Fiber Golang documentation](https://docs.gofiber.io/)

# Useful Links and Resources

[Makromusic 😊](https://makromusic.com/)

[How to Create REST API by using Golang and PostgreSQL ?](https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j)

# Not Working on Project

## Kafka
- Unfortunately, Kafka is not working on this project. **kafka** file is unusable... (I don't know why, I tried to fix about 3 day it but it didn't work)

## PostgresSQL run command
docker run --rm -d  --name dev-postgres  --network dev-network  -e POSTGRES_USER=postgres  -e POSTGRES_PASSWORD=password  -e POSTGRES_DB=postgres  -v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data  -p 5432:5432  postgres 