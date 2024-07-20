# list available recipes
list:
  @just --list

# start sfs
start:
  docker compose up --detach --build --force-recreate

# stop sfs
stop:
  docker compose down -v
