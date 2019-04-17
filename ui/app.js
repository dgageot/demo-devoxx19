'use strict';

var proxy = require('express-http-proxy');
const express = require('express');
const app = express();

if (app.get('env') === 'development') {
    app.use(require('easy-livereload')({
        checkFunc: (file) => {
            console.log(file);
            return true;
        }
    }));
}

app.use(express.static('public'));
app.use(proxy('http://localhost:8888'));

app.listen(8080, () => {
  console.log(`App listening on port 8080`);
});
