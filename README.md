## Environment Setup:

**Step 1**: You need to create the production env variables for the service.   
[never push the env variables such as secret to github or an repo]

example .env file is added to the zip (with local set up values)

`source .env`

**Step 2**: Bring up the database via docker (in this case we are using mongodb) using   
`docker-compose up -d`

## How to run the code:

### With Docker:
**Step 1**: Build the docker image using `docker build .`
**Step 2**: Create a `.docker.env` file [sample .docker.env file is attached in zip]
**Step 3**: To run the built docker image use,
`docker run --env-file=.docker.env --net=host image-id` 
`

### Without Docker using make commands:
* We can use command `make run` to run the service.
* We can use command `make test` to execute tests.  
* We can also use command `make cover` to check the coverage of tests in the service.


## Additional Information:
We can use command `make build` to build the service.
[refer Makefile]

## Further Improvements:
* We can add lint to improve the quality and maintainability of project.
* We can add logger to improve debugging.
