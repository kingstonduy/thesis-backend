curl -i -X POST http://localhost:8001/upstreams \
 --data "name=cart-service-upstream"

# Add the first target

curl -i -X POST http://localhost:8001/upstreams/cart-service-upstream/targets \
 --data "target=cart-service-1:7032"

# Add the second target

curl -i -X POST http://localhost:8001/upstreams/cart-service-upstream/targets \
 --data "target=cart-service-2:7033"

# Add the third target

curl -i -X POST http://localhost:8001/upstreams/cart-service-upstream/targets \
 --data "target=cart-service-3:7034"

curl -i -X POST http://localhost:8001/services \
 --data "name=cart-service" \
 --data "url=http://cart-service-upstream"

curl -i -X POST http://localhost:8001/routes \
 --data "name=cart-service-route" \
 --data "paths[]=/cart-service" \
 --data "service.name=cart-service"

# list the routes

curl -i http://localhost:8001/upstreams/cart-service-upstream/targets
curl -i http://localhost:8001/upstreams/cart-service-upstream
