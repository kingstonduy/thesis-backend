curl -i -X POST http://localhost:8001/upstreams \
 --data "name=product-service-upstream"

curl -i -X POST http://localhost:8001/upstreams/product-service-upstream/targets \
 --data "target=product-service-1:7002"

curl -i -X POST http://localhost:8001/upstreams/product-service-upstream/targets \
 --data "target=product-service-2:7003"

curl -i -X POST http://localhost:8001/upstreams/product-service-upstream/targets \
 --data "target=product-service-3:7004"

curl -i -X POST http://localhost:8001/services \
 --data "name=product-service" \
 --data "url=http://product-service-upstream"

curl -i -X POST http://localhost:8001/routes \
 --data "name=product-service-route" \
 --data "paths[]=/product-service" \
 --data "service.name=product-service"

curl -i http://localhost:8001/upstreams/product-service-upstream/targets
curl -i http://localhost:8001/upstreams/product-service-upstream

curl -i -X POST http://localhost:8001/upstreams \
 --data "name=product-service-upstream"

curl -i -X POST http://localhost:8001/upstreams/product-service-upstream/targets \
 --data "target=product-service-1:7002"

curl -i -X POST http://localhost:8001/upstreams/product-service-upstream/targets \
 --data "target=product-service-2:7003"

curl -i -X POST http://localhost:8001/upstreams/product-service-upstream/targets \
 --data "target=product-service-3:7004"

curl -i -X POST http://localhost:8001/services \
 --data "name=product-service" \
 --data "url=http://product-service-upstream"

curl -i -X POST http://localhost:8001/routes \
 --data "name=product-service-route" \
 --data "paths[]=/product-service" \
 --data "service.name=product-service"

curl -i http://localhost:8001/upstreams/product-service-upstream/targets
curl -i http://localhost:8001/upstreams/product-service-upstream

curl -i -X POST http://localhost:8001/upstreams \
 --data "name=cart-service-upstream"

curl -i -X POST http://localhost:8001/upstreams/cart-service-upstream/targets \
 --data "target=cart-service-1:7032"

curl -i -X POST http://localhost:8001/upstreams/cart-service-upstream/targets \
 --data "target=cart-service-2:7033"

curl -i -X POST http://localhost:8001/upstreams/cart-service-upstream/targets \
 --data "target=cart-service-3:7034"

curl -i -X POST http://localhost:8001/services \
 --data "name=cart-service" \
 --data "url=http://cart-service-upstream"

curl -i -X POST http://localhost:8001/routes \
 --data "name=cart-service-route" \
 --data "paths[]=/cart-service" \
 --data "service.name=cart-service"

curl -i http://localhost:8001/upstreams/cart-service-upstream/targets
curl -i http://localhost:8001/upstreams/cart-service-upstream

curl -i -X POST http://localhost:8001/upstreams \
 --data "name=order-service-upstream"

curl -i -X POST http://localhost:8001/upstreams/order-service-upstream/targets \
 --data "target=order-service-1:7012"

curl -i -X POST http://localhost:8001/upstreams/order-service-upstream/targets \
 --data "target=order-service-2:7013"

curl -i -X POST http://localhost:8001/upstreams/order-service-upstream/targets \
 --data "target=order-service-3:7014"

curl -i -X POST http://localhost:8001/services \
 --data "name=order-service" \
 --data "url=http://order-service-upstream"

curl -i -X POST http://localhost:8001/routes \
 --data "name=order-service-route" \
 --data "paths[]=/order-service" \
 --data "service.name=order-service"

curl -i http://localhost:8001/upstreams/order-service-upstream/targets
curl -i http://localhost:8001/upstreams/order-service-upstream
