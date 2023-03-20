# keycloak

`helm repo add codecentric https://codecentric.github.io/helm-charts`

`helm repo up`

`kubectl create ns keycloak`

`helm -n keycloak install keycloak codecentric/keycloak --version=17.0.3 -f values-keycloak.yaml`

Http management interface listening on http://127.0.0.1:9990/management

Admin console listening on http://127.0.0.1:9990

https://keycloak.workshop.metakube.org/

You need local access to create the initial admin user.

`k port-forward  svc/keycloak-http 8080:80`

Open http://localhost:8080/auth
or use the add-user-keycloak script.
