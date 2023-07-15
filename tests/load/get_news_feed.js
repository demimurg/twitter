import grpc from 'k6/net/grpc';
import {check} from 'k6';

// K6_PROMETHEUS_RW_SERVER_URL=http://mimir.demimurg.com/api/v1/push k6 run --vus 10 --rps 50 --duration 1m -o experimental-prometheus-rw tests/load/get_news_feed.js

const client = new grpc.Client();
const packagePrefix = 'github.com.demimurg.twitter.v1.'

export default () => {
    if (__ITER == 0) {
        client.connect('twitter.demimurg.com:30000', {plaintext: true, reflect: true});
    }
    const response = client.invoke(packagePrefix + 'Twitter/GetNewsFeed', {
        user_id: 2, limit: 50,
    });
    check(response, {'status is OK': (r) => r && r.status === grpc.StatusOK});
};
