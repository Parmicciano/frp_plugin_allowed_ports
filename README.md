# frp_plugin_allowed_ports

frp server plugin to define allowed ports for a specific user for [frp](https://github.com/fatedier/frp).




### Features

* Support the verification of the port used by the users by ports & subdomain saved in a file. 

### Download

Download fp-multiuser binary file from [Release](https://github.com/Parmicciano/frp_plugin_allowed_ports/releases).

### Requirements

frp version >= v0.42.0

It is possible that the plugin works for older version even though it has not been tested. 

### Usage

1. Create file `ports` including all support usernames and ports.

    ```
    user1=65536
    user2=80
    user2=525
    user1=6980
    user2=subdomain
    ```

    One user each line. Username and token are split by `=`.

2. Run fp-multiuser:

    `./fp-multiuser -l 127.0.0.1:7200 -f ./tokens`

3. Register plugin in frps.

    ```
    # frps.ini
    [common]
    bind_port = 7000

    [plugin.multiuser]
    addr = 127.0.0.1:7200
    path = /handler
    ops = Login
    ```

4. Specify username and meta_token in frpc configure file.

    For user1:

    ```
    # frpc.ini
    [common]
    server_addr = x.x.x.x
    server_port = 7000
    user = user1
    meta_token = 123

    [ssh]
    type = tcp
    local_port = 22
    remote_port = 6000
    ```

    For user2:

    ```
    # frpc.ini
    [common]
    server_addr = x.x.x.x
    server_port = 7000
    user = user2
    meta_token = abc

    [ssh]
    type = tcp
    local_port = 22
    remote_port = 6000
    ```

