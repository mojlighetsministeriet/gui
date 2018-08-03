import docCookies from './docCookies.js';

export default function jsonRequest(url, body, method = 'GET') {
  method = method.toUpperCase();

  if (body instanceof Object) {
    try {
      body = JSON.stringify(body);
    } catch (serializeError) {
      console.error(serializeError);
    }
  }

  return new Promise((resolve, reject) => {
    const options = {
      credentials: 'same-origin',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      method,
      body
    };

    const csrf = docCookies.getItem('csrf');
    if (csrf) {
      options.headers['X-CSRF-Token'] = csrf;
    }

    fetch(url, options).then((response) => {
      if (response.ok) {
        return response.json();
      }
      throw response.json();
    }).then(resolve).catch(reject);
  });
}
