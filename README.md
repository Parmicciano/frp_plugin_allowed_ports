# frp_plugin_allowed_ports

frp server plugin to define allowed ports for a specific user for [frp](https://github.com/fatedier/frp).




### Features

* Support the verification of the port used by the users by ports & subdomain saved in a file. 

### Download

Download frp_plugin_allowed_ports binary file from [Release](https://github.com/Parmicciano/frp_plugin_allowed_ports/releases).

### Requirements

frp version >= v0.42.0

It is possible that the plugin works for older version even though it has not been tested. 

### Usage
It works with custom_domains, tcp and subdomains.
1. Create file `ports` including all support usernames and ports.

    ```
    user1=65536
    user2=80
    user2=525
    user1=6980
    user2=subdomain
    user1=subdomain2
    user1=service.masternetwork.us
    ```

    One user each line. Username and token are split by `=`.

2. Run frp and the plugin:

    `./frps -c ./frps.ini`
    `./frp_plugin_allowed_ports -c ./frps.ini`

3. Register plugin in frps.

    
```

 [common]
bind_port = 7000
vhost_http_port = 80
dashboard_port = 7500

dashboard_user = admin
dashboard_pwd = admin
subdomain_host = masternetwork.us

[plugin.multiuser]
addr = 127.0.0.1:8000
path = /handler
ops = Login

[plugin.frp_plugin_allowed_ports]
addr = 127.0.0.1:9001
path = /handler
ops = NewProxy
```
4. Frpc file :

    User field is required

    ```
    # frpc.ini
    [common]
    server_addr = x.x.x.x
    server_port = 7000
    user = user1

    [ssh]
    type = tcp
    local_port = 22
    remote_port = 6000
    ```

