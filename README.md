# Assignment 1 Submission


Completed Assignment, enjoy :)
- Uses one HTTP server and one RPC server with Redis as the DB like described in the given architecture
- Messages follow the format described in idl_http.proto
- JSON is used for data serialization
- Postman was used to test the API

  
Below I have some code blocks to use with Postman for testing of core functionalities. First run everything with the docker-compose.yml file and then open Postman.


Test 1 - Import and run the below code block in Postman to save the first message
```
curl --location 'localhost:8080/api/send' \
--header 'Content-Type: application/json' \
--data '{
    "chat": "john:doe",
    "text": "hello world",
    "sender": "john"
}'
```

Test 2 - Pulling the message: Change to GET and run localhost:8080/api/pull with the following Body:
```
{
    "chat": "john:doe",
    "cursor": 0,
    "limit": 10,
    "reverse": false
}
```


Test 3 - Saving a second message with same chat: Change to POST and run localhost:8080/api/send with the following Body:
```
{
    "chat": "john:doe",
    "text": "nice weather today",
    "sender": "doe"
}
```


Test 4 - Pulling all messages: Change to GET and run localhost:8080/api/pull with the following Body:
```
{
    "chat": "john:doe",
    "cursor": 0,
    "limit": 10,
    "reverse": false
}
```

Test 5 - Pulling messages in reverse: Run localhost:8080/api/pull with the following Body:
```
{
    "chat": "john:doe",
    "cursor": 0,
    "limit": 10,
    "reverse": true
}
```

Test 6 - Pulling only the first message: Run localhost:8080/api/pull with the following Body:
```
{
    "chat": "john:doe",
    "cursor": 0,
    "limit": 1,
    "reverse": false
}
```

Test 7 - Saving a message with a different chat: Change to POST and run localhost:8080/api/send with the following Body:
```
{
    "chat": "jack:jill",
    "text": "that is a mountain",
    "sender": "jill"
}
```

Test 8 - Pulling only jack:jill messages: Change to GET and run localhost:8080/api/pull with the following Body:
```
{
    "chat": "jack:jill",
    "cursor": 0,
    "limit": 10,
    "reverse": false
}
```
