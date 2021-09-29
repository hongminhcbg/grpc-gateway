# grpc-gateway: Run http server and grpc server in same pod

I. How to run this example

    $ go build -o server cmd/server/main.go && ./server
    
    you will see server listening in two port
    http at 9000 and grpc at 10000
    after that you can test http server with postman or curl 
    curl example
    $ curl -d '{"name":"you name", "age": 18}' -X POST 'http://localhost:9000/api/v1/users'


    Test grpc server with example client
    $ go build -o client cmd/client/client.go && ./client
    
    expected client log
    [hongminh] result after create user, resp = {"code":1,"messsage":"success","data":{"user_id":1632568428}}, error = <nil>    
    [minh] result after create user, resp = null, error = rpc error: code = Unknown desc = name existed


II. headers support:
	case
		"Accept",
		"Accept-Charset",
		"Accept-Language",
		"Accept-Ranges",
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Cookie",
		"Date",
		"Expect",
		"From",
		"Host",
		"If-Match",
		"If-Modified-Since",
		"If-None-Match",
		"If-Schedule-Tag-Match",
		"If-Unmodified-Since",
		"Max-Forwards",
		"Origin",
		"Pragma",
		"Referer",
		"User-Agent",
		"Via",
		"Warning":

