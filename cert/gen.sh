#!/bin/sh

rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -nodes -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=BD/ST=Dhaka/L=Dhaka/O=Sazid/OU=none/CN=Sazid/emailAddress=sazidozon@gmail.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=BD/ST=Dhaka/L=Dhaka/O=Sazid/OU=none/CN=learngrpc/emailAddress=sazidozon@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed
#    certificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -extfile server-ext.cnf -out server-cert.pem

echo "Server's CA signed certificate"
openssl x509 -in server-cert.pem -noout -text

# 4. Verify server certificate
echo "Verifying generated server certificate (server-cert.pem)"
openssl verify -CAfile ca-cert.pem server-cert.pem
