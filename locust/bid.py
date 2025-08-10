import json
from locust import HttpUser, events, task, between
import random

import websocket

def generate_ws_url(auction_id):
    return f"ws://localhost:8080/ws?auction_id={auction_id}"

class AuctionUser(HttpUser):
    wait_time = between(1, 10)  # 模拟真实用户操作间隔

    def on_start(self):
        """登录后获取 token"""
        response = self.client.post("/login", json={
            "name": "test",
            "password": "test"
        })

        if response.status_code == 200:
            self.token = response.json().get("access_token")
        else:
            self.token = None
        
        self.ws = websocket.WebSocketApp(
            generate_ws_url(2),
            on_message=self.on_message,
            on_error=self.on_error,
            on_close=self.on_close,
        )

        self.bid_amount = 0
    
    def on_stop(self):
        if self.ws:
            self.ws.close()
    
    def on_message(self, ws, message):
        try:
            data = json.loads(message)
            print(f"[收到消息] {message}")
            if "bid_info" in data:
                events.request.fire(
                    request_type="WebSocket",
                    name="receive_bid",
                    response_time=0,  # WebSocket 接收消息没延迟意义，这里用 0
                    response_length=len(message),
                    exception=None
                )
                self.bid_amount = data["bid_info"]["price"]
        except Exception as e:
            events.request.fire(
                request_type="WebSocket",
                name="receive_bid",
                response_time=0,
                response_length=0,
                exception=e
            )

    def on_error(self, ws, error):
        events.request.fire(
            request_type="WebSocket",
            name="ws_error",
            response_time=0,
            response_length=0,
            exception=error
        )

    def on_close(self, ws, close_status_code, close_msg):
        print(f"[连接关闭] code={close_status_code} msg={close_msg}")

    @task
    def bid(self):
        """先获取当前价格，再进行竞价"""
        if not self.token:
            return  # 没登录成功就跳过

        headers = {
            "Authorization": f"Bearer {self.token}"
        }

        # 第一步：获取当前最高出价 建立ws连接
        auction_id = 2
        current_price = self.bid_amount

        # 第二步：生成一个比当前高的价格
        increment = random.randint(-2, 2)
        new_price = current_price + increment

        # 第三步：发起出价请求
        payload = {
            "auction_id": auction_id,
            "price": new_price
        }

        bid_resp = self.client.post("/api/bid", json=payload, headers=headers)
        if bid_resp.status_code != 200:
            print("Bid failed:", bid_resp.text)
