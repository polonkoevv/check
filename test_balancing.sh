#!/bin/bash

# Количество запросов
requests=50

echo "Тестирование балансировки нагрузки..."
echo "Отправка $requests запросов..."

for i in $(seq 1 $requests)
do
    curl -s -I -H "Host: example.com" http://http://83.222.17.152:80/ | grep X-Upstream
    sleep 0.1
done

echo "Тестирование завершено."