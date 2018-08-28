import jsonRequest from './jsonRequest.js';

const state = {};

const email = document.querySelector('input[name="email"]');
const password = document.querySelector('input[name="password"]');

document.querySelector('#signin-button').addEventListener('click', () => {
  const body = {email: email.value, password: password.value};
  jsonRequest('/api/session', body, 'POST').then((response) => {
    console.log(response);
    state.authenticated = true;
  }).catch((sessionError) => {
    console.error(sessionError);
    state.authenticated = false;
  });
})

document.querySelector('#signout-button').addEventListener('click', () => {
  jsonRequest('/api/session', undefined, 'DELETE').then((response) => {
    console.log(response);
    state.authenticated = false;
  }).catch((sessionError) => {
    console.error(sessionError);
  });
})

setInterval(() => {
  if (state.authenticated) {
    document.querySelector('#signout-button').style.display = 'block';
    document.querySelector('#signin').style.display = 'none';
  } else {
    document.querySelector('#signout-button').style.display = 'none';
    document.querySelector('#signin').style.display = 'block';
  }
}, 50);
