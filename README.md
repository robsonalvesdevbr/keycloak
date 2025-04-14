```bash
docker rm $(docker ps --filter 'label=br.dev.robsonalves.service.group=keycloak' -a -q) -f --volumes
```