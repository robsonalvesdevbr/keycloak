```bash
docker rm $(docker ps --filter 'label=br.dev.robsonalves.service=keycloak' -a -q) -f --volumes
```