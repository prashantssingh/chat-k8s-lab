import random

REPLIES = [
    "Hey! 👋",
    "Got it.",
    "Tell me more.",
    "Interesting—why do you think that?",
    "👍",
    "Haha 😄",
]

def pick_reply() -> str:
    return random.choice(REPLIES)