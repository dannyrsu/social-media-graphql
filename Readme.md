### Social-Media-GraphQL

- Still need to add graphql

#### Run these to get full sample working
```bash
docker run --rm --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres
```

```bash
export GO111MODULE=on
```