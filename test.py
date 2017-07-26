from threading import Thread
import requests
import timeit


def worker():
    threads = []
    start_time = timeit.default_timer()
    for a in range(0, 1000):
        t = Thread(target=fucc)
        threads.append(t)
        t.start()
    for thread in threads:
        thread.join()
    elapsed = timeit.default_timer() - start_time
    print(elapsed)
    return


def fucc():
    r = requests.get("https://api.mercadolibre.com/sites")
    print(r.status_code)

if __name__ == '__main__':
    worker()
