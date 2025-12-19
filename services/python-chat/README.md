### From services/python-chat

```bash
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
uvicorn app.main:app --host 0.0.0.0 --port 8080
```

### Test:
```bash
curl -s http://localhost:8080/healthz
curl -s http://localhost:8080/readyz

curl -s -X POST http://localhost:8080/reply \
  -H 'Content-Type: application/json' \
  -d '{"from":"go-chat","message":"hello","trace_id":"t-123"}'
```