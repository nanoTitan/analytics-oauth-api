

# Cassandra
# -------------------------------------------------------------------------
# Install
Download and install latest JRE
Add JAVA_HOME environment variable for JRE path
Test - echo %JAVA_HOME% (Should show environment path like C:\Program Files\Java\jre1.8.0_241)

# Start Cassandra DB
open cmd prompt
cd C:\Program Files\apache-cassandra-3.11.6\bin
cassandra.bat -f

# Start Cassandra cqlsh
open cmd prompt
cd C:\Program Files\apache-cassandra-3.11.6\bin
cqlsh

# Run Application
# -------------------------------------------------------------------------
cd ~/go/src/github.com/nanoTitan/analytics-oauth-api/src

export postgres_users_host="localhost"
export postgres_users_port=5432
export postgres_users_user="postgres"
export postgres_users_password="admin"
export postgres_users_dbname="users"

go run main.go

# Run Tests
# -------------------------------------------------------------------------
cd src
# run single test
go test

# run all tests
go test ./...
