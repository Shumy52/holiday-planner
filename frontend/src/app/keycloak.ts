import Keycloak from 'keycloak-js';

export const keycloak = new Keycloak({
  url: (window as any)['KEYCLOAK_URL'] || 'http://localhost:8080',
  realm: (window as any)['KEYCLOAK_REALM'] || 'vacations',
  clientId: (window as any)['KEYCLOAK_CLIENT'] || 'spa-frontend'
});

export async function initKeycloak() {
  await keycloak.init({ onLoad: 'login-required', checkLoginIframe: false });
  (window as any).kc = keycloak; // debug
}
