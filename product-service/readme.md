curl -i -X POST http://localhost:8001/upstreams \
 --data "name=product-service-upstream"

# Add the first target

curl -i -X POST http://localhost:8001/upstreams/product-service-upstream/targets \
 --data "target=product-service-1:7002"

# Add the second target

curl -i -X POST http://localhost:8001/upstreams/product-service-upstream/targets \
 --data "target=product-service-2:7003"

# Add the third target

curl -i -X POST http://localhost:8001/upstreams/product-service-upstream/targets \
 --data "target=product-service-3:7004"

curl -i -X POST http://localhost:8001/services \
 --data "name=product-service" \
 --data "url=http://product-service-upstream"

curl -i -X POST http://localhost:8001/routes \
 --data "name=product-service-route" \
 --data "paths[]=/product-service" \
 --data "service.name=product-service"

# list the routes

curl -i http://localhost:8001/upstreams/product-service-upstream/targets
curl -i http://localhost:8001/upstreams/product-service-upstream
