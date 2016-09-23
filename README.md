#HFAT - Http Forwading AgenT

Would you like to capture http requests and the mirror them to several target servers?
Well this is what HFAT is indended to be for.

Since it is a few hours project it actually supports forwarding of get requests over http.

in HFAT directory go and start it using

    go run main.go
    
use **config.json** to define target Servers

In the config file you may mark a server as **Primary: true**
This effects that the response of the primary server is returned to the caller.