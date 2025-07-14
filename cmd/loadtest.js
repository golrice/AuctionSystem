import http from 'k6/http';
import { sleep, check } from 'k6';

export let options = {
  stages: [
    { duration: '10s', target: 50 },   // 预热
    { duration: '20s', target: 1000 },  // 压力阶段
    { duration: '10s', target: 0 },    // 冷却
  ],
};

export default function () {
  let res = http.get('http://localhost:8080/ping');

  check(res, {
    'status is 200': (r) => r.status === 200,
    'body is pong': (r) => r.body.includes('pong'),
  });

  sleep(1);
}
