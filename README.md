go-chat → python-chat request

POST /reply

```json
{
  "from": "go-chat",
  "message": "hello",
  "trace_id": "uuid-or-rand"
}
```

```json
python-chat response
{
  "from": "python-chat",
  "reply": "hi there",
  "trace_id": "same-trace-id"
}
```