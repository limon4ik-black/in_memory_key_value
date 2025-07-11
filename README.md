# In-Memory Key-Value Store

Простое хранилище ключ-значение в памяти для быстрого доступа к данным.

## Что это

Это библиотека для хранения данных в памяти в формате ключ-значение. Все данные хранятся в оперативной памяти, что обеспечивает очень быстрый доступ к ним.


## Как запустить

```bash
git clone https://github.com/limon4ik-black/in_memory_key_value.git
cd in_memory_key_value
make clean
make build
make run-server
make run-client
```

## Что умеет

- **SET** - сохранить значение по ключу
- **GET** - получить значение по ключу  
- **DEL** - удалить ключ


## Пример использования

```cpp
InMemoryKV store;

SET name John
GET name
DEL name
```
##  Особенность (ну или нет)
WAL
