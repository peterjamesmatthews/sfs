# list all available commands
list:
  @just --list

# start sfs
start:
  docker compose up -d --build --force-recreate

# stop sfs
stop:
  docker compose down -v
