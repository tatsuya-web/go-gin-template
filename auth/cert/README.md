# generate secret this directly!

public.pem

> openssl genrsa 4096 > secret.pem

secret.pem

> openssl rsa -pubout < secret.pem > public.pem
