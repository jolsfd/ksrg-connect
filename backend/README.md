# Ksrg-Connect-Backend

## Build image

You need to configure your server in the `config.yaml` file.

```bash
nano config.yaml
```

```bash
docker build . -t ksrg-connect-backend
```

## Run image

If you run the container with a volume make sure that you created a `config.yaml` file in it. To run with a persistent database:

```bash
docker run -p 127.0.0.1:8000:8000 -v /var/ksrg-connect/:/usr/src/app/backend/public/ ksrg-connect-backend
```

or without persistent data

```bash
docker run -p 127.0.0.1:8000:8000 ksrg-connect-backend
```

## Configuration

* **Port**: Port where webserver is running
* **Domain**: Used for json web token creation
* **Origins**: Allowed origins for example your frontend url (CORS)
* **Database**: Path to database file
* **Duration**: Duration in which the json web token is valid in hours
* **Secret**: Secret string for json web token generation
* **Password**: Authentication password for sign up
* **Admins**: Usernames for users who have admin rights. **Warning!**