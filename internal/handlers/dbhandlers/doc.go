/*
dbhandlers cервис генерирует короткие URL из оригинальных URL и сохраняет их в базе данных.

Используются Get и Post http запросы.

Для роутера используется библиотека chi: https://github.com/go-chi/chi

В новой версии реализация сохранения в базу данных и реализация сохранения в файл и память разделены на уровне хендлеров.

Используется аутентификация на уровне cookies.

Используются ендпоинты:

В бд:

 / [post]
 /api/shorten [post] -
 /{short} [get]
 /api/shorten/batch [post]
 /ping
 /api/user/urls [get]
 /api/user/urls [delete]
*/
package dbhandlers
