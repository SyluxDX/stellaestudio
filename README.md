# Stella Estudio
⚠️ Under construction 🚧🛠️

## DataBase Schemas

_Add image here_

## Create databases's tables 
```
$ sqlite3 <database_name>.db < init_scripts/01-init_database.sql
```

## Insert dummy data
```
$ cat init_scripts/02-dummy_data.sql | sqlite3 <database_name>.db
```

## Simultaneous create tables and insert data
```
$ cat init_scripts/*.sql | sqlite3 <database_name>.db
```