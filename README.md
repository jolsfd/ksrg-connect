# KSRG-Connect

## Installation

First you need to clone the repository.
```
git clone https://github.com/jolsfd/ksrg-connect.git
```

Before you build your docker images edit the `frontend/.env` and `backend/public/config.yaml` file with your personal settings.
After you have done this you need to build two docker images.
```bash
# backend image
cd backend
docker build . -t ksrg-connect-backend

# frontend image
cd frontend
docker build . -t ksrg-connect-frontend
```

## Deploying

To run the website on your machine or server you need to run the docker conatiner. For the backend you have two options.
One option is to run without a volume. **You need to know that data in this container will be lost if the container is powered off.**


The recommended option is to run the container with a volume for your configuration and database. 
For that you need to create a folder `ksrg-connect` in your `/var` directory.
**Make sure you have copied your `config.yaml` file in this new folder.**


The following commands are used to run the containers:
```bash
# frontend
cd frontend
docker run -d -p 127.0.0.1:8080:8080 ksrg-connect-frontend

# backend
cd backend
docker run -d -p 127.0.0.1:8000:8000 -v /var/ksrg-connect/:/usr/src/app/backend/public/ ksrg-connect-backend

# without volume
docker run -d -p 127.0.0.1:8000:8000 ksrg-connect-backend
```


To serve the website in the internet a reverse proxy is used.
Caddy sets up the reverse proxy and handles ssl certificates automatically.
```bash
sudo caddy run
```

## Deploying on a VPS

* add an user `ksrg-connect` with root privilegies
* add public key to the VPS with `ssh-copy-id`
---
* edit `/etc/ssh/sshd_config`
```
PasswordAuthentication no
PermitRootLogin no
```
---
* setup up a firewall `ufw`
* setup `fail2ban`
* install `caddy`
* install `docker`
* setup `docker`
* change domain and subdomain in `Caddyfile`
* edit settings in the backend `config.yaml` for example password, admins, ...
* add legal information for example DSGVO, Impressum, ...
* refactor FAQ, start page, ...
* point AAAA record to the server's ip address
* point the subdomain to the server
* be happy 🎉