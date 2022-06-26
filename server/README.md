# Postgres and Go inside WSL2

Networking is different from WSL2 and Windows. The following links help resolve connecting to localhost from WSL2.

https://stackoverflow.com/questions/56824788/how-to-connect-to-windows-postgres-database-from-wsl

https://docs.microsoft.com/en-us/windows/wsl/networking#accessing-windows-networking-apps-from-linux-host-ip

```go
// Example connection string to read/write to db in Windows network (not preferred)
pg := postgres.NewDB("postgres://postgres:secret@172.26.144.1:5432/postgres?sslmode=disable&connect_timeout=10")

// The host `localhost` will point to postgres inside WSL2 network. (preferred)
pg := postgres.NewDB("postgres://postgres:secret@localhost:5432/postgres?sslmode=disable&connect_timeout=10")

```

# Install and Startup Postgres inside WSL2

1. Install PostgreSQL - https://docs.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql

2. Make sure Postgres service is running inside WSL2

```sh
$ sudo service postgresql start
```

3. Pass a connection string to our app. The host is `localhost` when running postgres inside WSL2

```
"postgres://{user}:{password}@{host}:{port}/{database_name}?{optional_params}={param}"
```

# General Postgres Setup after install

https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-20-04

1. Create a new user/role

```sh
$ sudo -u postgres createuser --interactive

Output
Enter name of role to add: sammy
Shall the new role be a superuser? (y/n) y
```

2. Create a database with same name as new user

- If logged in as postgres

```sh
postgres@server:~$ createdb sammy
```

- Without switching accounts

```sh
$ sudo -u postgres createdb sammy
```

3. Open Postgres (psql) with new user/role

```sh
$ sudo adduser sammy

# Switch account and connect to database.
$ sudo -i -u sammy
$ psql

# Can do inline
$ sudo -u sammy psql

# Connect to different database
$ psql -d postgres
```

Get connection info inside postgres shell:

```
sammy=# \conninfo
```
