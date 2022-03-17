# Ksrg-Connect-Frontend

## Build image

First you need to configure your api url in the `.env` file.

```bash
nano .env
```

```bash
docker build . -t ksrg-connect-frontend
```

## Run image

```bash
docker run -p 127.0.0.1:8080:8080 ksrg-connect-frontend
```