var createError = require('http-errors');
var express = require('express');
var path = require('path');
var fs = require('fs');
const lodash = require('lodash')
var cookieParser = require('cookie-parser');
var logger = require('morgan');
var session = require('express-session');
var index = require('./routes/index');
var bodyParser = require('body-parser');//解析,用req.body获取post参数
var app = express();
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));
app.use(cookieParser());
app.use(session({
  secret : 'secret', // 对session id 相关的cookie 进行签名
  resave : true,
  saveUninitialized: false, // 是否保存未初始化的会话
  cookie : {
    maxAge : 1000 * 60 * 3, // 设置 session 的有效时间，单位毫秒
  },
}));
// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');
// app.engine('ejs', function (filePath, options, callback) {    // 设置使用 ejs 模板引擎 
//   fs.readFile(filePath, (err, content) => {
//       if (err) return callback(new Error(err))
//       let compiled = lodash.template(content)    // 使用 lodash.template 创建一个预编译模板方法供后面使用
//       let rendered = compiled()

//       return callback(null, rendered)
//   })
// });
app.use(logger('dev'));
app.use(express.static(path.join(__dirname, 'public')));
app.use('/', index);
// app.use('/challenge7', challenge7);
// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

module.exports = app;
