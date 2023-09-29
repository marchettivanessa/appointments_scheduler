# Appointments Scheduler
This is an exercise project for creating appointments on calendar.
If you are a little curious about this little experiment, I should tell you that first:

1. If you don't have it, install (golang)[https://go.dev/doc/install]
2. (Echo)[https://echo.labstack.com/docs/quick-start#installation] is needed as well.
3. For migrations, go to (golang-migrate)[https://github.com/golang-migrate/migrate].

## Running docker:
This project uses docker default commands:
1. `docker-compose build` for building the docker image;
2. `docker-compose up -d` for initializing the docker containers;

And, if you want to drop all containers, use `docker-compose down`

## Running the api:
For now, a basic `go run main.go` do the trick. You'll see echo logo being shown as the application is started.