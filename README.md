# To-Do List

## 1. Use o docker para rodar
```sh
docker compose up --build
```

## 2. Crie uma tabela em um banco de dados

Primeiro, acesse o container do PostgreSQL:
```sh
docker exec -it to-do-list-postgres-1 psql -U seu_user -d seu_banco
```

Dentro do terminal do PostgreSQL, crie o banco e a tabela:
```sql
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    task TEXT NOT NULL,
    status BOOLEAN NOT NULL DEFAULT false
);
```
