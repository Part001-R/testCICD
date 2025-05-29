
# Создание образов
## Сервер
Создание образа: `docker build -f configs/server/Dockerfile -t srv_img:0.1.0 .`
`-f configs/server/Dockerfile` - путь к файлу
`-t srv_img:0.1.0` - присвоение имени и тега, к создаваемому образу. Имя и версия, произвольные.

## Клиент
Создание образа: `docker build -f configs/client/Dockerfile -t cl_img:0.1.0 .`
`-f configs/client/Dockerfile` - путь к файлу
`-t cl_img:0.1.0` - присвоение имени и тега, к создаваемому образу. Имя и версия, произвольные.

## Проверка
`docker images` - вывод списка образов.

# Compose
Запуск контейнеров -`docker compose up`
Работа сопровождается выводом в терминал генерируемых контейнерами сообщений.

Подключение к выводу - `docker logs -f <контейнер>`

# Версии
`v0.1.0` - Actions

