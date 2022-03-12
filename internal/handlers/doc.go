/*
dbhandlers cервис генерирует короткие URL из оригинальных URL и сохраняет их в базе данных.

Используются Get и Post http запросы.

Для роутера используется библиотека chi: https://github.com/go-chi/chi

В новой версии реализация сохранения в базу данных и реализация сохранения в файл и память разделены на уровне хендлеров.

Используется аутентификация на уровне cookies.

Используются ендпоинты:

В бд:

 POST /
 POST /api/shorten
 GET /{short}
 POST /api/shorten/batch
 GET /ping
 GET /api/user/urls
 DELETE /api/user/urls
*/
package handlers
