#!/bin/bash

mysql -h localhost -u root -proot orders < internal/infra/database/migrations/000001_create_orders_table.up.sql 