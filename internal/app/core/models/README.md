* structure of directory
  models/ temp/

* cd temp

* excute below commands
gentool -db postgres -dsn "host=localhost user=id password=password dbname=admin port=5432 sslmode=disable" -onlyModel -fieldWithTypeTag -modelPkgName models -tables table_name -withUnitTest

