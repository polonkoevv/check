#!/bin/bash
requests=10

echo "Тестирование балансировки нагрузки..."
echo "Отправка $requests запросов..."

for i in $(seq 1 $requests)
do
    curl --insecure -s -I -H "Host: localhost" https://127.0.0.1/api/books | grep X-Upstream
    sleep 0.1
done

echo "Тестирование завершено."