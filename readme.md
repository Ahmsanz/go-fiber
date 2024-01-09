### Summary
This is a simple project that aims to serve as showcase of how to build a small API using GO and Fiber (and helped me learn a bit about it, myself).

It exposes the CRUD endpoints to manipulate a _Book_ entity, comprised of only three fields (the purpose was not to be exhaustive).


#### Architecture
The project has been structured around the aproximate interpretation of the principles suggested by Eric Evans and his DDD book, with three main layers, from inner to outer:

- Domain: with the entities and interfaces that represent the core "bussiness" logic
- Infrastructure: with all the connections with external functionalities, such as the specific implementation of the repositories that will carry out the operations in the database.
- Application: in charge of the implementation of the domain functionalities. It orchestrates the different use cases accessing the other two layers in order to achieve the desired result.


#### Extra features
- Dockerization: the app has been containerized using docker.
- JWT authentication/authorization middleware: a login endpoint is in place to allow the securization of the *book* routes.


#### Missing pieces
- Testing: at least unit and e2e tests should and will be added (when I find the time)
- Auth missing cases: validation and logout would be useful features
- User verification: jwt middleware shoyld be coupled with verifying whether the user actually exists or not, and if the password they're using is correct. This has been skipped, as it wasn't the priority.

#### Execution
Run ``` go run main.go``` from the root of the project and it will start the server locally. You'll need to install the dependecies first using ``` go install```. **You will also need a Postgres database with an accessible port**.

You can also run it all via Docker, for which you'll need docker compose installed. Run ```docker-compose -f ./iac/docker/docker-compose.yaml up``` and it will spin a database and the server itself for you.

**Please note that you will need a .env file at the root of the project, with at least the values for the variables provided in the .env.sample**


#### Like what you see?
've been a NodeJS developer for some years now and this project was only meant for my own learning and practice process, but it has some hints of what I would apply on a real, productive project.

If you like the potential of where this could be going, feel free to reach me via [LinkedIn](https://www.linkedin.com/in/adrianmohmedsanz/) and I'll be more than happy to connect and have a chat!  