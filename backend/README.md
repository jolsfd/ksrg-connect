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

To run with a persistent database:

```bash
docker run -p 127.0.0.1:8000:8000 -v /var/ksrg-connect/:/usr/src/app/backend/database/ ksrg-connect-backend
```

or without persistent data

```bash
docker run -p 127.0.0.1:8000:8000 ksrg-connect-backend
```