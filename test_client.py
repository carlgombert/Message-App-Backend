import websocket
import rel

def on_message(ws, message):
    print("Received:", message)

def on_error(ws, error):
    print("Error:", error)

def on_close(ws):
    print("Connection closed")

def on_open(ws):
    print("Connection established")
    # You can send messages after the connection is established
    ws.send("Hello, WebSocket!")


websocket.enableTrace(True)
ws = websocket.WebSocketApp("ws://localhost:8080/ws",
                            on_message=on_message,
                            on_error=on_error,
                            on_close=on_close)
ws.on_open = on_open
ws.run_forever()
rel.signal(2, rel.abort)