import time
import urllib.request

URLS = [
    "http://localhost:6060/debug/pprof/heap",
    "http://localhost:6060/debug/pprof/goroutine?debug=1",
]

print("Simple test client for pprof endpoints. Press Ctrl+C to stop.")
while True:
    for url in URLS:
        try:
            with urllib.request.urlopen(url, timeout=5) as response:
                print(url, response.status, len(response.read()))
        except Exception as exc:
            print(url, "ERROR", exc)
    time.sleep(5)
