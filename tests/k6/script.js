import http from 'k6/http';

const BASE_URL = 'http://localhost:8888';

export const options = {
  stages: [
    { duration: '30s', target: 100 },
    { duration: '30s', target: 0 },
  ]
};

export default function () {
  const requests = {
      'Create create_orbit': {
        method: 'GET',
        url: `${BASE_URL}/create_orbit`,
      },
      'get_orbitst': {
        method: 'GET',
        url: `${BASE_URL}/get_orbits`,
      },
      'create_satellite': {
        method: 'GET',
        url: `${BASE_URL}/create_satellite`,
      },
      'get_satellites': {
        method: 'GET',
        url: `${BASE_URL}/get_satellites`,
      },
    };

    http.batch(requests);
}