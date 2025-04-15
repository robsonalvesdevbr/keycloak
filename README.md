```bash
docker rm $(docker ps --filter 'label=br.dev.robsonalves.service.group=keycloak' -a -q) -f --volumes
docker compose -f docker-compose.yaml -f docker-compose.override.yaml up -d
```