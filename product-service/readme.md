# custom image

```
docker build --network=host -t ubuntu-confluent .
docker run -d --name confluent-container -p 9092:9092 -p 9021:9021 -p 8083:8083 ubuntu-confluent
```

# debezium

````shell
# Start the topology as defined in https://debezium.io/documentation/reference/stable/tutorial.html
docker-compose -f docker-compose-postgres.yaml up -d

# set up postgres
```sql
ALTER SYSTEM SET wal_level = logical;
select * from pg_settings where name ='wal_level';
```

# Start Postgres connector
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @register-postgres.json

# Consume messages from a Debezium topic
docker-compose -f docker-compose-postgres.yaml exec kafka /kafka/bin/kafka-console-consumer.sh \
    --bootstrap-server kafka:9092 \
    --from-beginning \
    --property print.key=true \
    --topic dbserver1.inventory.customers

# Modify records in the database via Postgres client
docker-compose -f docker-compose-postgres.yaml exec postgres env PGOPTIONS="--search_path=inventory" bash -c 'psql -U $POSTGRES_USER postgres'

# Shut down the cluster
docker-compose -f docker-compose-postgres.yaml down
````
