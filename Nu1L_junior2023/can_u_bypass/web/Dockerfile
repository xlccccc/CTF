FROM python:3.7

COPY app.py /app.py

RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple && \
    pip install --upgrade pip &&\
    pip install oracledb flask

CMD ["python", "/app.py"]
