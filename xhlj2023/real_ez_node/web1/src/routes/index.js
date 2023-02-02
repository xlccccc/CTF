var express = require('express');
var http = require('http');
var router = express.Router();
const safeobj = require('safe-obj');
router.get('/',(req,res)=>{
  if (req.query.q) {
    console.log('get q');
  }
  res.render('index');
})
router.post('/copy',(req,res)=>{
  res.setHeader('Content-type','text/html;charset=utf-8')
  var ip = req.connection.remoteAddress;
  console.log(ip);
  var obj = {
      msg: '',
  }
  if (!ip.includes('127.0.0.1')) {
      obj.msg="only for admin"
      res.send(JSON.stringify(obj));
      return 
  }
  let user = {};
  for (let index in req.body) {
      if(!index.includes("__proto__")){
          safeobj.expand(user, index, req.body[index])
      }
    }
  res.render('index');
})

router.get('/curl', function(req, res) {
    var q = req.query.q;
    var resp = "";
    if (q) {
        var url = 'http://localhost:3000/?q=' + q
            try {
                http.get(url,(res1)=>{
                    const { statusCode } = res1;
                    const contentType = res1.headers['content-type'];
                  
                    let error;
                    // 任何 2xx 状态码都表示成功响应，但这里只检查 200。
                    if (statusCode !== 200) {
                      error = new Error('Request Failed.\n' +
                                        `Status Code: ${statusCode}`);
                    }
                    if (error) {
                      console.error(error.message);
                      // 消费响应数据以释放内存
                      res1.resume();
                      return;
                    }
                  
                    res1.setEncoding('utf8');
                    let rawData = '';
                    res1.on('data', (chunk) => { rawData += chunk;
                    res.end('request success') });
                    res1.on('end', () => {
                      try {
                        const parsedData = JSON.parse(rawData);
                        res.end(parsedData+'');
                      } catch (e) {
                        res.end(e.message+'');
                      }
                    });
                  }).on('error', (e) => {
                    res.end(`Got error: ${e.message}`);
                  })
                res.end('ok');
            } catch (error) {
                res.end(error+'');
            }
    } else {
        res.send("search param 'q' missing!");
    }
})
module.exports = router;