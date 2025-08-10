import json
import redis
import websocket

def on_message(ws, message):
    print(f"[收到消息] {message}")

def on_error(ws, error):
    print(f"[错误] {error}")

def on_close(ws, close_status_code, close_msg):
    print(f"[连接关闭] code={close_status_code} msg={close_msg}")

def on_open(ws):
    print("[连接打开]")

# ping 的时候，服务器会返回 pong
def on_ping(ws, data):
    print(f"[收到ping] {data}")

app = websocket.WebSocketApp(
    "ws://localhost:8080/ws?auction_id=2",
    on_message=on_message,
    on_error=on_error,
    on_close=on_close,
    on_open=on_open,
    on_ping=on_ping,
)

# 连接到redis并且订阅 auction:2 频道
r = redis.Redis(host='localhost', port=6379, decode_responses=True)

# 要发送的消息
message = {
    "auction_id": 2,
    "bid_info": {
        "price": 10000,
        "timestamp": 1234567890
    }
}

# 发布到频道
# channel = "auction:2"
# r.publish(channel, json.dumps(message))
# 使用一个定时器，定时发送消息
import time
import threading

stop_event = threading.Event()

def send_message():
    while not stop_event.is_set():
        channel = "auction:2"
        r.publish(channel, json.dumps(message))
        time.sleep(30)

threading.Thread(target=send_message, daemon=True).start()

# 启动websocket
try:
    app.run_forever()
except KeyboardInterrupt:
    stop_event.set()
    time.sleep(1)
