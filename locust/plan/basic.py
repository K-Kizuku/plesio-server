from locust import HttpUser, task

class HelloWorldUser(HttpUser):
    # 最初の1回呼び出される。
    def on_start(self):
        # dummyAPIの関数URL
        response = self.client.get("https://ih2xtobxskw4lvdhd7sdgscfsi0gtqoi.lambda-url.ap-northeast-1.on.aws/")
        print("Response text:", response.text)

    # 2回目以降、hello_worldが呼び出される。
    @task
    def hello_world(self):
        # dummyAPI2の関数URL (試験開始時に入力するHostの値を使用)
        response = self.client.get("")
        print("Response text:", response.text)