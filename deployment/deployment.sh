#!/bin/sh

# Generate keys into a temporary directory.
echo "Generating TLS keys ..."
"./webhook-generate-keys.sh" "keys/"

# Create the TLS secret for the generated keys.
echo "Creating secret ..."
kubectl -n preset create secret tls preset-pod \
    --cert "keys/webhook-server-tls.crt" \
    --key "keys/webhook-server-tls.key"

# Read the PEM-encoded CA certificate, base64 encode it, and replace the `${CA_PEM_B64}` placeholder in the YAML
# template with it. Then, create the Kubernetes resources.
echo "Deployment ..."
ca_pem_b64="$(openssl base64 -A < "keys/ca.crt")"
sed -e 's@${CA_PEM_B64}@'"$ca_pem_b64"'@g' <"deployment_all_in_one.yaml" \
    | kubectl create -f -