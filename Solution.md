# Solution

For this assignment I've decided to go with a gRPC server that exposes a REST gateway.
I think gRPC is a great choice for a micro-services project (even though this assignment only consists of one service) as it is very performant and has a strict specification.
By choosing gRPC we get a public REST API for clients and also get a gRPC API which could be used for internal inter-service communication.

I wanted to implement pagination as the fee data might not always be shown in its entirety.
Due to the fact that this data is a result of processing data from 2 tables, I could not implement standard SQL offset and limit, hence I had to do it in Go rather than SQL.

If I had more time I would have also implemented caching for the processed fees, so we do not have to read from DB and process the transactions everytime we make a request.
I think metric collections would also be good to have, maybe using something like Prometheus.
Tracing wouldn't really make sense in this scenario as we only have one service, but could definitely be added in the future.

