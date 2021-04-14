#!/bin/sh
#
# Generate a (self-signed) CA certificate and a certificate and private key to be used by the webhook demo server.
# The certificate will be issued for the Common Name (CN) of `webhook-server.webhook-demo.svc`, which is the
# cluster-internal DNS name for the service.
#
# NOTE: THIS SCRIPT EXISTS FOR DEMO PURPOSES ONLY. DO NOT USE IT FOR YOUR PRODUCTION WORKLOADS.

: "$1?'missing key directory'"

key_dir="$1"

chmod 0700 "$key_dir"
cd "$key_dir"

[ -z "$service" ] && service=preset-pod
[ -z "$namespace" ] && namespace=preset

# 生成CA证书和CA私钥
openssl req -nodes -new -x509 -keyout ca.key -out ca.crt -subj "/CN=Admission Controller Webhook Kingfisher" -days 36500
# 生成Webhook服务的私钥
openssl genrsa -out webhook-server-tls.key 2048
# Generate a Certificate Signing Request (CSR) for the private key, and sign it with the private key of the CA.
# 为Webhook服务的私钥生成证书签名请求(CSR)，并使用CA的私钥对其进行签名
openssl req -new -key webhook-server-tls.key -subj "/CN=${service}.${namespace}.svc"  \
    | openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -out webhook-server-tls.crt -days 36500