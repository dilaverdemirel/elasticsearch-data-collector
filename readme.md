go mod init github.com/dilaverdemirel/eslasticsearchdatacollector
go get -u github.com/go-sql-driver/mysql
go run elasticsearchdatacollector.go

# Features
Zamanlanmış ya da manuel olarak tetiklenebilecek veri transfer tanımlarına göre verileri RDBMS kaynaklardan
elasticsearch'e aktaracak yapıyı oluşturmak gerekli.

- Datasource tanımları eklenebilir
- Index aktarım tanımları eklenebilir
    - Reload all ya da iterative güncelleme özelliği sağlanabilir
    - Document ID özelliği sağlanmalı
- Index aktarım tanımlarına schedule etme özelliği sağlanabilir

- SQL preview özelliği eklenebilir

### Datasource
Verinin çekileceği database için kullanılacak datasource tanımlarını saklar.

|Column                 |Type            |Description |
|-----------------------|----------------|------------|
|id                     | varchar(36)    |            |
|name                   | varchar(50)    |            |
|connection_string      | varchar(500)   |            |
|max_pool_size          | int(3)         |            |
|min_idle               | int(3)         |            |
|username               | varchar(50)    |            |
|db_password            | varchar(100)   | encrypted  |
|driver_name            | varchar(30)    |            |


### Index
Uygulama tarafından oluşturulan Elasticsearch indexleri ile ilgili bilgileri yöneten tablo.

|Column                 |Type            |Description |
|-----------------------|----------------|------------|
|id                     | varchar(36)    |            |
|name                   | varchar(50)    |            |
|alias                  | varchar(70)    | Reload all yüklemelerde yeni oluşturulan indexe verilen alias |
|description            | varchar(100)   |            |
|valid                  | varchar(1)     | Y/N        |
|sql_query              | clob           |            |
|scheduled              | varchar(1)     | Y/N        |
|cron_expression        | varchar(100)   |            |
|last_execution_time    | datetime       |            |
|sync_type              | varchar(20)    | ITERATIVE/RELOAD_ALL |
|data_source_id         | varchar(36)    |            |
|document_field_id      | varchar(36)    | ES indexlemesi sırasında document-id olarak belirlenecek sql field name|


### SyncLog
Syncronizasyon işlemlerini loglamak için kullanılır.

|Column                 |Type            |Description |
|-----------------------|----------------|------------|
|id                     | varchar(36)    |            |
|index_id               | varchar(36)    |            |
|document_count         | int(15)        |            |
|start_time             | datetime       |            |
|end_time               | datetime       |            |
|exec_duration_in_sec   | int(15)        |            |
|status                 | varchar(20)    | STARTED/COMPLETED/FAILED |
|status_message         | clob           |            |

## Eksikler

[-] Db password encode edilecek



## End Points

### DataSourceController /data-sources

#### create POST /data-sources
**Request:**
```json
{
  "driverClass": "string",
  "jdbcUrl": "string",
  "maximumPoolSize": 0,
  "minimumIdle": 0,
  "name": "string",
  "password": "string",
  "type": "RDBMS",
  "username": "string"
}
```
**Response:**
```json
{
  "creDate": "2024-01-10T12:12:48.011Z",
  "creUser": "string",
  "driverClass": "string",
  "id": "string",
  "jdbcUrl": "string",
  "maximumPoolSize": 0,
  "minimumIdle": 0,
  "name": "string",
  "orgId": "string",
  "password": "string",
  "type": "RDBMS",
  "updDate": "2024-01-10T12:12:48.011Z",
  "updUser": "string",
  "username": "string"
}
```

#### update PUT /data-sources/{id}
**Request:**
```json
{
  "driverClass": "string",
  "jdbcUrl": "string",
  "maximumPoolSize": 0,
  "minimumIdle": 0,
  "name": "string",
  "password": "string",
  "type": "RDBMS",
  "username": "string"
}
```
**Response:**
```json
{
  "creDate": "2024-01-10T12:12:48.011Z",
  "creUser": "string",
  "driverClass": "string",
  "id": "string",
  "jdbcUrl": "string",
  "maximumPoolSize": 0,
  "minimumIdle": 0,
  "name": "string",
  "orgId": "string",
  "password": "string",
  "type": "RDBMS",
  "updDate": "2024-01-10T12:12:48.011Z",
  "updUser": "string",
  "username": "string"
}
```

#### getById GET /data-sources/{id}

**Response:**
```json
{
  "creDate": "2024-01-10T12:12:48.011Z",
  "creUser": "string",
  "driverClass": "string",
  "id": "string",
  "jdbcUrl": "string",
  "maximumPoolSize": 0,
  "minimumIdle": 0,
  "name": "string",
  "orgId": "string",
  "password": "string",
  "type": "RDBMS",
  "updDate": "2024-01-10T12:12:48.011Z",
  "updUser": "string",
  "username": "string"
}
```

#### find GET /data-sources?parameters....

**Response:**
```json
[{
  "creDate": "2024-01-10T12:12:48.011Z",
  "creUser": "string",
  "driverClass": "string",
  "id": "string",
  "jdbcUrl": "string",
  "maximumPoolSize": 0,
  "minimumIdle": 0,
  "name": "string",
  "orgId": "string",
  "password": "string",
  "type": "RDBMS",
  "updDate": "2024-01-10T12:12:48.011Z",
  "updUser": "string",
  "username": "string"
}]
```

-------------

### IndexController /indices

#### create POST /indices
**Request:**
```json
{
  "dataSourceId": "string",
  "description": "string",
  "fieldMetaData": [
    {
      "dataType": "DATE",
      "fieldName": "string"
    }
  ],
  "indexGenerationQuery": "string",
  "indexName": "string",
  "settings": "string"
}
```
**Response:**
```json
{
  "exampleData": [
    {
      "additionalProp1": {},
      "additionalProp2": {},
      "additionalProp3": {}
    }
  ],
  "metaDataList": [
    {
      "dataType": "DATE",
      "fieldName": "string"
    }
  ]
}
```

#### update PUT /indices/{id}
**Request:**
```json
{
  "id": "string",
  "settings": "string",
  "valid": true
}
```
**Response:**
```json
{
  "alias": "string",
  "creDate": "2024-01-10T12:35:45.865Z",
  "creUser": "string",
  "cronExpression": "string",
  "dataSourceId": "string",
  "description": "string",
  "documentIdField": "string",
  "id": "string",
  "indexGenerationQuery": "string",
  "lastSyncExecutionDate": "2024-01-10T12:35:45.865Z",
  "metaData": {
    "data": "string"
  },
  "name": "string",
  "orgId": "string",
  "scheduled": true,
  "settings": {
    "data": "string"
  },
  "syncType": "ITERATIVE",
  "type": "MANUEL",
  "updDate": "2024-01-10T12:35:45.865Z",
  "updUser": "string",
  "valid": true
}
```

#### delete DELETE /indices/{id}


#### scheduleIndexDataSync PUT /indices/{id}/schedule-data-sync
**Request:**
```json
{
  "cronExpression": "string",
  "documentIdField": "string",
  "indexId": "string",
  "syncType": "ITERATIVE"
}
```
**Response:**
```json
{
  "alias": "string",
  "creDate": "2024-01-10T12:35:45.865Z",
  "creUser": "string",
  "cronExpression": "string",
  "dataSourceId": "string",
  "description": "string",
  "documentIdField": "string",
  "id": "string",
  "indexGenerationQuery": "string",
  "lastSyncExecutionDate": "2024-01-10T12:35:45.865Z",
  "metaData": {
    "data": "string"
  },
  "name": "string",
  "orgId": "string",
  "scheduled": true,
  "settings": {
    "data": "string"
  },
  "syncType": "ITERATIVE",
  "type": "MANUEL",
  "updDate": "2024-01-10T12:35:45.865Z",
  "updUser": "string",
  "valid": true
}
```

#### unscheduleIndexDataSync DELETE /indices/{id}/unschedule-data-sync

**Response:**
```json
{
  "alias": "string",
  "creDate": "2024-01-10T12:35:45.865Z",
  "creUser": "string",
  "cronExpression": "string",
  "dataSourceId": "string",
  "description": "string",
  "documentIdField": "string",
  "id": "string",
  "indexGenerationQuery": "string",
  "lastSyncExecutionDate": "2024-01-10T12:35:45.865Z",
  "metaData": {
    "data": "string"
  },
  "name": "string",
  "orgId": "string",
  "scheduled": true,
  "settings": {
    "data": "string"
  },
  "syncType": "ITERATIVE",
  "type": "MANUEL",
  "updDate": "2024-01-10T12:35:45.865Z",
  "updUser": "string",
  "valid": true
}
```


#### getById GET /indices/{id}

**Response:**
```json
{
  "alias": "string",
  "creDate": "2024-01-10T12:35:45.865Z",
  "creUser": "string",
  "cronExpression": "string",
  "dataSourceId": "string",
  "description": "string",
  "documentIdField": "string",
  "id": "string",
  "indexGenerationQuery": "string",
  "lastSyncExecutionDate": "2024-01-10T12:35:45.865Z",
  "metaData": {
    "data": "string"
  },
  "name": "string",
  "orgId": "string",
  "scheduled": true,
  "settings": {
    "data": "string"
  },
  "syncType": "ITERATIVE",
  "type": "MANUEL",
  "updDate": "2024-01-10T12:35:45.865Z",
  "updUser": "string",
  "valid"
}
```

#### find GET /indices?parameters....

**Response:**
```json
[{
  "alias": "string",
  "creDate": "2024-01-10T12:35:45.865Z",
  "creUser": "string",
  "cronExpression": "string",
  "dataSourceId": "string",
  "description": "string",
  "documentIdField": "string",
  "id": "string",
  "indexGenerationQuery": "string",
  "lastSyncExecutionDate": "2024-01-10T12:35:45.865Z",
  "metaData": {
    "data": "string"
  },
  "name": "string",
  "orgId": "string",
  "scheduled": true,
  "settings": {
    "data": "string"
  },
  "syncType": "ITERATIVE",
  "type": "MANUEL",
  "updDate": "2024-01-10T12:35:45.865Z",
  "updUser": "string",
  "valid"
}]
```

-------------

### QueryMetaDataController /query-meta-data

#### preview POST /query-meta-data/preview
**Request:**
```json
{
  "dataSourceId": "string",
  "query": "string"
}
```
**Response:**
```json
{
  "exampleData": [
    {
      "additionalProp1": {},
      "additionalProp2": {},
      "additionalProp3": {}
    }
  ],
  "metaDataList": [
    {
      "dataType": "DATE",
      "fieldName": "string"
    }
  ]
}
```

-------------

### SynchronizationController /synchronization

#### start POST /synchronization/{indexId}/start

#### find GET /query-meta-data/logs?parameters

**Response:**
```json
{
  "content": [
    {
      "creDate": "2024-01-10T12:46:42.876Z",
      "creUser": "string",
      "documentCount": 0,
      "endDate": "2024-01-10T12:46:42.876Z",
      "executionDuration": 0,
      "id": "string",
      "indexId": "string",
      "orgId": "string",
      "startDate": "2024-01-10T12:46:42.876Z",
      "status": "COMPLETED",
      "statusMessage": "string",
      "updDate": "2024-01-10T12:46:42.876Z",
      "updUser": "string"
    }
  ],
  "empty": true,
  "first": true,
  "last": true,
  "number": 0,
  "numberOfElements": 0,
  "pageable": {
    "offset": 0,
    "pageNumber": 0,
    "pageSize": 0,
    "paged": true,
    "sort": {
      "empty": true,
      "sorted": true,
      "unsorted": true
    },
    "unpaged": true
  },
  "size": 0,
  "sort": {
    "empty": true,
    "sorted": true,
    "unsorted": true
  },
  "totalElements": 0,
  "totalPages": 0
}
```