JWT - JSON Web Token
JWT  is a means of exchanging information between two parties
Information is embedded inside the payload part of the token, which ois digitally signed

Structure:
{Base64 encoded Header}.{Base64 encoded Payload}.{Signature}

The header contains algorithm and token type

{
    "alg": "HS256",
    "typ": "JWT"
}

The payload can carry claims:
- User and additional data such as the token expiry etc.
- Three types of claims: Registered, Public and Private
  
Signature:
Computed from the header, Payload and a secret
- An algorithm to generate the Signature
- Digitally signed using a secret string only known to the developer





