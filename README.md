#Dinner Hall

Network Programming first laboratory work. This repository contains **Dinner Hall** web-server implementation.

##Usage

```bash
docker-compose --project-name "[NAME]" up --detach
```

<small>*or*</small>

```bash
docker-compose --project-name "[NAME]" --file "[COMPOSE_FILE_PATH]" \
                run --publish "[HOST_PORT]:[CONTAINER_PORT]" dinner_hall
```

<small>*...if you want to fully control container run...*</small>

 * **[NAME]** - *by default you can use* `restaurant`.
 * **[COMPOSE_FILE_PATH]** - *by default use* `./docker-compose.yml`.
 * **[HOST_PORT]:[CONTAINER_PORT]** - *by default I used* `56565:56565`.