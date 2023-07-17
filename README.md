# MateMind

[![Test](https://github.com/good1hare/MateMind/actions/workflows/ci.yml/badge.svg)](https://github.com/good1hare/MateMind/blob/master/.github/workflows/ci.yml)
![Supported Go Versions](https://img.shields.io/badge/Go-1.20-lightgrey.svg)
[![License](https://img.shields.io/github/license/good1hare/MateMind.svg)](https://github.com/good1hare/MateMind/blob/master/LICENSE)
[![Release](https://img.shields.io/github/v/release/good1hare/MateMind.svg)](https://github.com/good1hare/MateMind/releases/)

MateMind - телеграмм бот, который поможет вам управлять задачами, устанавливать напоминания и создавать быстрые заметки.
Он предоставляет функции по созданию, редактированию и отслеживанию задач, а также возможность создавать заметки и
сохранять их в удобном формате. Бот позволяет быть в курсе всех своих задач и не пропустить ни одного важного события,
благодаря функции напоминаний. Начните использовать MateMind уже сегодня и станьте более продуктивным!

## Installation

Конфиги читаются из config/config.yml(для разработки локально), затем если есть перезаписываются файлом config.env(для
прода или дев на сервере)

Описание взаимодействия базы данных и таблиц [PostgreSQL](README_PostgreSQL.md)

Создание миграций
migrate create -ext sql -dir migrations table_name