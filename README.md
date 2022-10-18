# Momo Store aka Пельменная №2

Репозиторий предназначен для хранения кода приложения Momo Store.

Исходный код инфраструктуры и инструкция по развертыванию приложения
лежит [здесь](https://gitlab.praktikum-services.ru/anton-alekseyev/momo-store-infra)

Внешний вид стартого экрана приложения:
<img width="900" alt="image" src="https://user-images.githubusercontent.com/9394918/167876466-2c530828-d658-4efe-9064-825626cc6db5.png">

## Структура репозитория

```
├── backend - код бекенда пельменной
│   ├── .gitlab-ci.yml - CI/CD манифест для бекенда
│   ├── Dockerfile - докерфайл для бекенда
├── frontend - код фронтенда пельменной
│   ├── .gitlab-ci.yml - CI/CD манифест для фронтенда
│   ├── Dockerfile - докерфайл для фронтенда
├── .gitlab-ci.yml - CI/CD манифест для приложения
├── docker-compose.local.yml - docker compose для локального запуска и тестирования
```

## Инструкция пользователя

### Frontend

Запуск приложения локально

```bash
cd frontend
npm install
VUE_APP_API_URL=http://localhost:8081 VUE_APP_VERSION=0.0.1 npm run serve
```
При конвейерной сборки образ контейнера загружается в Gitlab Container Registry репозитория.

Артефакт после сборки в виде архива загружается в [nexus](https://nexus.praktikum-services.ru/repository/momostore-alekseev-anton-frontend/)

Также в код frontend добавлена поддержка вывод версионирования (в интерфейсе отображается Пельменная №$VERSION).
Исправлены замечания от sonarqube.

### Backend

Запуск приложения локально

```bash
cd backend
go run ./cmd/api
go test -v ./...
```
При конвейерной сборки образ контейнера загружается в Gitlab Container Registry репозитория.

Артефакт после сборки в виде архива загружается в [nexus](https://nexus.praktikum-services.ru/repository/momostore-alekseev-anton-backend/)
### Docker compose

Запуск docker compose через

```shell
docker-compose --file docker-compose.local.yml up --build
```