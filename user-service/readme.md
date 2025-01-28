curl -i -X POST http://localhost:8001/upstreams \
 --data "name=user-service-upstream"

# Add the first target

curl -i -X POST http://localhost:8001/upstreams/user-service-upstream/targets \
 --data "target=user-service-1:7022"

# Add the second target

curl -i -X POST http://localhost:8001/upstreams/user-service-upstream/targets \
 --data "target=user-service-2:7023"

# Add the third target

curl -i -X POST http://localhost:8001/upstreams/user-service-upstream/targets \
 --data "target=user-service-3:7024"

curl -i -X POST http://localhost:8001/services \
 --data "name=user-service" \
 --data "url=http://user-service-upstream"

curl -i -X POST http://localhost:8001/routes \
 --data "name=user-service-route" \
 --data "paths[]=/user-service" \
 --data "service.name=user-service"

# list the routes

curl -i http://localhost:8001/upstreams/user-service-upstream/targets
curl -i http://localhost:8001/upstreams/user-service-upstream
