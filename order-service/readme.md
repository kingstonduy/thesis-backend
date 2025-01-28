curl -i -X POST http://localhost:8001/upstreams \
 --data "name=order-service-upstream"

# Add the first target

curl -i -X POST http://localhost:8001/upstreams/order-service-upstream/targets \
 --data "target=order-service-1:7012"

# Add the second target

curl -i -X POST http://localhost:8001/upstreams/order-service-upstream/targets \
 --data "target=order-service-2:7013"

# Add the third target

curl -i -X POST http://localhost:8001/upstreams/order-service-upstream/targets \
 --data "target=order-service-3:7014"

curl -i -X POST http://localhost:8001/services \
 --data "name=order-service" \
 --data "url=http://order-service-upstream"

curl -i -X POST http://localhost:8001/routes \
 --data "name=order-service-route" \
 --data "paths[]=/order-service" \
 --data "service.name=order-service"

# list the routes

curl -i http://localhost:8001/upstreams/order-service-upstream/targets
curl -i http://localhost:8001/upstreams/order-service-upstream
