#!/bin/bash

# Копирование исполняемого файла и конфигурации на сервер
scp dns-over-mesh user@server:/path/to/deploy
scp config/config.yaml user@server:/path/to/deploy

echo "Deployment completed successfully!"

