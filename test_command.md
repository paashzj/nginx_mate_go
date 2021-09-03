## tcp static route
### add amqp route
```bash
curl -X POST localhost:31014/v1/nginx/route/static/tcp -d '{"in_port":5671, "in_ssl":true, "out_host":"localhost", "out_port":5672}'
```
### delete amqp route
```bash
curl -X DELETE localhost:31014/v1/nginx/route/static/tcp/0
```