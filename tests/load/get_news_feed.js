import grpc from 'k6/net/grpc';
import {check} from 'k6';

const client = new grpc.Client();
//client.load(['../../pkg/proto', '.'], 'twitter.proto');

export default () => {
    client.connect('twitter.demimurg.com:30000', {plaintext: true, reflect: true});

    const data = {user_id: 1, limit: 50};
    const response = client.invoke('Twitter/GetNewsFeed', data);
    check(response, {
        'status is OK': (r) => r && r.status === grpc.StatusOK,
    });
    client.close();
};
