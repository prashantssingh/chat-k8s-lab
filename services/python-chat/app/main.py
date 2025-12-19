# services/python-chat/app/main.py
import os
import time
from fastapi import FastAPI, Request
from pydantic import BaseModel
from .messages import pick_reply

app = FastAPI(title="python-chat", version="1.0.0")

class ReplyRequest(BaseModel):
    from_: str | None = None  # we'll map "from" manually
    message: str
    trace_id: str | None = None

@app.get("/healthz")
def healthz():
    return {"status": "ok"}

@app.get("/readyz")
def readyz():
    # Later we can check downstream deps here
    return {"status": "ready"}

@app.post("/reply")
async def reply(request: Request):
    body = await request.json()
    msg = body.get("message", "")
    trace_id = body.get("trace_id")

    r = pick_reply()

    # "Production-ish" structured log (simple)
    print({
        "ts": time.time(),
        "service": "python-chat",
        "path": "/reply",
        "trace_id": trace_id,
        "message_len": len(msg),
        "reply": r,
    })

    return {"from": "python-chat", "reply": r, "trace_id": trace_id}
