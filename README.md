# ozon-test
# Сервис сокращения ссылок на Go

## Запуск проекта
Выбор хранения данных:
Перейдите в файл config.yaml.
```bash
storage:
  type: "in-memory"  # или "postgres"
postgres:
  dsn: "postgres://ROOT:password@localhost:5432/urls?sslmode=disable"
server:
  port: 8080
```
в строчке type: в скобках укажите нужный тип хранения
```bash
docker-compose up --build
```
## Кодирование ссылки
```bash
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "ссылка_на_сайт"}'
```

## Декодирование ссылки из базы данных
```bash
curl http://localhost:8080/expand/закодированная_ссылка
```

## Просмотр базы данных PostgreSQL
Для этого, во время запущенной программы, в терминале введите 
```bash 
psql -U ROOT -d urls
```
Далее введите:
```bash
SELECT * FROM urls
```
