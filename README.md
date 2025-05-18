# Snaptrail

Snaptrail is a personal passion project. It's goal is to provide a web application to track and share your journey as a photographer.\
As of now, it is possible for the admin user to create sessions with some basic details describing them, and also uploading one picture for a session - this could be your favorite or most memorable picture. The picture must contain EXIF-Data, as the backend will attempt to paritally extract some information to display - if there is none it will cause the session to not be saveable.
Sessions can be set to public or private - if publically available, they will be visible in a timeline, ordered by date as well as the session list view.
It features a golang backend, an angular frontend and utilizes a PostgreSQL database.\
Currently the project is still a prototype. Large commits and breaking changes are to be expected. Test coverage is also less than ideal - will work on that as things stabilize.

## DISCLAIMER

Be aware the picture you upload will be stored like that in the database and also displayed in the web.
This means you're pictures will end up on the PC of people visiting the page - with no way of preventing downloads. Consider using lower resolution pictures and watermarks!

## Development Setup

You will need:

- Golang (>= 1.24 recommended)
- Angular 19
- relational database (PostgreSQL recommended)
- [dbmate](https://github.com/amacneil/dbmate) for migrations
- docker (recommended)
- SOPS (optional, to en/decrypt secrets)

You can use the `compose.yaml` to locally setup a PostgreSQL database as well as adminer.
The docker image of the application is built for ARM machines, so using that might not be possible.
The application requires some environment variables to be set.
For local development the following are required:

- DATABASE_URL
- ADMIN_USERNAME
- ADMIN_PASSWORD
- JWT_SECRET

And if using docker compose to setup the database (this should match the DATABASE_URL)

- DB_PASSWORD
- DB_USER
- DB

1. Setup your database

```
docker compose up
```

2. Run migrations

```
dbmate up
```

3. Start backend

```
go run internal/main.go
```

4. Start frontend

```
cd web && npm run start
```

And you should be good to go!\
On startup, the backend will query the database for the username foundin the `ADMIN_USERNAME` ENV. If no record can be found, a user will be created taking into account the provided credentials. These can be used to log in.\
Commits on main or dev trigger corresponding GitHub Actions, executing automated tests, building/pushing a Docker image, decrypt secrets and,if possible, deploy the applicaiton with Kustomize

## Deployment

To deploy Snaptrail, some more environment variables are needed:

- DOMAIN_SUFFIX
- UI_DIR

The application is automatically deployed using GitHub Actions and Kustomize to a K3S Cluster in the Hetzner Cloud.\
Using [Kube Hetzner](https://github.com/kube-hetzner/terraform-hcloud-kube-hetzner) the cluster can be created on the fly as needed.
All required environment variables are set directly in the deployment or through secrets.

Currently, there are two environments:\
[INT](https://int.snaptrail.markusharder.com/ui/) and [PROD](https://snaptrail.markusharder.com/ui/)\
These are automatically deployed on successful commits for dev and main respectively.
Since cost is a factor, I don't keep these running at all times. If you are unable to access them, you can reach out to me and I will spin them up as soon as I am able to.

## What's to come?

These are my **current** goals for - they are not final in any way and can (most likely will) change

- [ ] Filters
- [ ] Viewer/Editor roles
- [ ] Simple user management
- [x] Mobile friendly layout
- [ ] Session Gallery - multiple pictures for one session
- [ ] Session Ranking
- [ ] Comments
- [ ] Object storage for pictures
