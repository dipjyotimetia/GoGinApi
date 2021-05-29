import http from "k6/http";
import {check, sleep} from "k6";
import NewRel from './newrelic'

const apiKey = process.env.NR_APIKEY
const nr = new NewRel(apiKey);

export let options = {
    stages: [
        {duration: "30s", target: 5},
        {duration: "1m30s", target: 2},
        {duration: "20s", target: 0},
    ]
}

export function setup() {
    nr.PrintAlertingStatus();
    // nr.Notify(
    //     'Sunday load test - short',
    //     'START',
    //     'Beginning E2E load test script',
    //     'test@gmail.com',
    // );
}


export default function () {
    let url = 'http://172.17.0.1:8082/api/login';
    let payload = JSON.stringify(
        {
            "email": "test1@gmail.com",
            "password": "password1"
        });
    let params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };
    let res = http.post(url, payload, params);

    check(res, {
        "status was 200": (r) => r.status === 200,
        "transaction time OK": (r) => r.timings.duration < 200
    });
    sleep(1);
}

export function teardown() {
//  nr.Notify(
//    'CI load test',
//    'END',
//    'Finishing CI load test script',
//    'test@gmail.com'
//    );
    nr.PrintAlertingStatus();
}