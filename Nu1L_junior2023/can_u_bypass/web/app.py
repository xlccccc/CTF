import logging

import oracledb
from flask import Flask, request
import socket

app = Flask(__name__)
username = "system"
password = "PaAasSSsSwoRrRDd"
ol_server = socket.gethostbyname("db1")
ol_port = 1521
sid = "orcl"
dsn = "{}:{}/{}".format(ol_server, ol_port, sid)
logging.basicConfig(level=logging.INFO, format='%(asctime)s %(message)s')


def check(sql):
    blacklist = ["select", "insert", "delete", "update", "table", "user", "drop", "alert", "procedure", "exec",
                 "open", ":=", "declare", "runtime", "process", "invoke", "newinstance","parse",
                 ".class", "loader", "script", "url", "xml", "method", "field", "reflect", "defineclass",
                 "getclass", "forname", "constructor", "transform", "sql", "beans", ".net", "http", ".rmi", "naming"
                 ]
    sql = sql.lower()
    for blackword in blacklist:
        if blackword in sql:
            return True


@app.route("/query", methods=["POST"])
def query():
    sql = request.form["sql"]
    if check(sql):
        return "waf"
    else:
        try:
            conn = oracledb.connect(user=username, password=password, dsn=dsn)
            conn.callTimeout = 5000
            cursor = conn.cursor()
            cursor.execute(sql)
            cursor.close()
            conn.close()
        except Exception as e:
            logging.info(str(e))
            return "error"
        return "query success"


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8888)
