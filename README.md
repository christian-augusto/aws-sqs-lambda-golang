### Executando localmente
Você pode executar o lambda localmente via localstack, como mostramos a seguir:

Execute o comando abaixo a partir do diretório raiz para subir os containeres do ambiente (terminal 1):
```sh
docker-compose up -d --build
```

compilar o projeto (terminal 1):
```sh
bash scripts/build.sh
```

configurar infra (terminal 2):
```sh 
docker container exec -it localstack bash
bash /app/scripts/start.sh
```

executar projeto (terminal 2):
```sh 
bash /app/scripts/send-message.sh
```

verificar registro redis (terminar 3):
```sh 
docker container exec -it redis sh
redis-cli
```

parar projeto:
```sh
bash scripts/stop.sh
```
