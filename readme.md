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
