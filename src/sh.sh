#!/bin/bash

# Проверка без роли
echo "Тест 1 - Без роли:"
curl -v -X POST http://localhost:8081/dummyLogin \
     -H "Content-Type: application/json" \
     -d '{}'

echo -e "\n\nТест 2 - Неверная роль:"
curl -v -X POST http://localhost:8081/dummyLogin \
     -H "Content-Type: application/json" \
     -d '{"role":"admin"}'

echo -e "\n\nТест 3 - Корректный запрос:"
curl -v -X POST http://localhost:8081/dummyLogin \
     -H "Content-Type: application/json" \
     -d '{"role":"employee"}'
