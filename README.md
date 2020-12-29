# Supermicro-IPMI-KEYGEN
Supermicro license key generator based on MAC address

<p>This code is going to generate Supermicro IPMI license key by Mac address based on HMAC-SHA1 Algoithm, which takes MAC address as input and claculate hash of MAC address via a key, which is found by P.Kleissner <a href="https://peterkleissner.com/" target="_top">in this document</a>. <br>
  
Usage: 
> $ go build keygen.go <br>
  $ ./keygen "MAC ADDRESS"
