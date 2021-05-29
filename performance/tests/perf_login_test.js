import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  stages: [
    { duration: "30s", target: 5 },
    { duration: "1m30s", target: 2 },
    { duration: "20s", target: 0 },
  ]
}

export default function() {
var url = 'http://localhost:8082/api/login';
  var payload = JSON.stringify(
  {
    "email": "test1@gmail.com",
    "password": "password1"
  });
  var params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };
  let res = http.post(url, payload, params);

  check(res, {
    "status was 200": (r) => r.status == 200,
    "transaction time OK": (r) => r.timings.duration < 200
  });
  sleep(1);
}