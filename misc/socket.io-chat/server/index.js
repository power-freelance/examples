import "dotenv/config";
import {createServer} from "http";
import express from "express";
import socketIO from "socket.io";

const app = express();
const http = createServer(app);
const io = socketIO(http);

io.on('connection', async function (socket) {
  console.log('socket -> connection');

  socket.on('disconnect', async function () {
    console.log('socket -> disconnected');
  });


  socket.on('message', async function(message) {
    io.emit('message', message);
  })
});

http.listen(process.env.LISTEN_PORT, function () {
  console.log(`listening on *:${process.env.LISTEN_PORT}`);
});

