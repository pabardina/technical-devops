Technical test
====

# How to start

Install resources

```shell
# create namespace
kubectl create ns flinks
kubectl create ns ingress-controller

# install ingress-controller
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update

# custom values to use pod anti affinity
helm install ingress-nginx ingress-nginx/ingress-nginx -n ingress-controller -f kubernetes/helm-nginx-values.yaml

# install cert-manager
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --version v1.2.0 \
  --create-namespace \
  --set installCRDs=true

# Install chart
helm upgrade --install awesome-apps -n flinks kubernetes/awesome-apps
```

# Usage

## Simple ap

https://newcomers.bardina.net

## Api

```shell
export API_HOSTNAME="newcomers-api.bardina.net"
```

## Get all todos

```shell
curl --request GET https://$API_HOSTNAME
```

# Get a todo 

```shell
curl --request GET https://$API_HOSTNAME/1
``` 

# Post a todo

```shell
curl --request POST https://$API_HOSTNAME \
--header 'Content-Type: application/json' \
--data-raw '{"Id":"3","Description":"Fitness"}'
```

# Update a todo

```shell
curl --request PUT https://$API_HOSTNAME/3 \
--header 'Content-Type: application/json' \
--data-raw '{"Description":"Plank"}'
```

# Delete a todo

```shell
curl --request DELETE https://$API_HOSTNAME/1
```

