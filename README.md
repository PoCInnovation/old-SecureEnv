# SecureEnv

Provide a secure environement for developers working whit Credentials you don't have to care about this tools update environment for you

# How does it work?

![image](./.github/assets/1Vault-1611045210156.webp)

# Getting Started

Make sure you have to install docker and run the Dockerfile like this run the project like this
```
docker build . -t ${NAME}

docker run -d --network="host"  --name secureEnv ${NAME}
```

# Installation

#### Make sure you installed [Goland](https://go.dev/doc/install) <br />

#### Install Goland dependencies
```shell
go get github.com/hashicorp/vault/api
```
#### Create directory for vault
```shell
mkdir -p $GOPATH/src/github.com/hashicorp && cd $_
```

#### Git clone vault-hashicorp repository
```shell
git clone https://github.com/hashicorp/vault.git
```
#### Change directory in vault
```shell
cd vault
```
#### Make dependencies
```shell
make bootstrap
```

#### Build project
```shell
make dev
```
#### Check vault version
```shell
vault -h
```

# Quickstart

#### Launch vault local server
```shell
vault server -dev
```
Get the Hash from the terminal where you launched the vault server you have to realize a .env like this
```
ADRESS=http://127.0.0.1:8200

HASH=hvs.PA3O9TDvERLIp7U9fNqM9uG8

PATHREAD=${PATH}
```


## Get involved

You're invited to join this project ! Check out the [contributing guide](./CONTRIBUTING.md).

If you're interested in how the project is organized at a higher level, please contact the current project manager.

## Our PoC team :heart:

Developers
| [<img src="https://github.com/tonida-rodda.png?size=85" width=85><br><sub>Toni Da-rodda</sub>](https://github.com/tonida-rodda) | [<img src="https://github.com/florianepitech.png?size=85" width=85><br><sub>Florian Damiot</sub>](https://github.com/florianepitech) |
| :---: | :---: |

Manager
| [<img src="https://github.com/adrienfort.png?size=85" width=85><br><sub>Adrien Fort</sub>](https://github.com/adrienfort)
| :---: |

<h2 align=center>
Organization
</h2>

<p align='center'>
    <a href="https://www.linkedin.com/company/pocinnovation/mycompany/">
        <img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white">
    </a>
    <a href="https://www.instagram.com/pocinnovation/">
        <img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white">
    </a>
    <a href="https://twitter.com/PoCInnovation">
        <img src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white">
    </a>
    <a href="https://discord.com/invite/Yqq2ADGDS7">
        <img src="https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white">
    </a>
</p>
<p align=center>
    <a href="https://www.poc-innovation.fr/">
        <img src="https://img.shields.io/badge/WebSite-1a2b6d?style=for-the-badge&logo=GitHub Sponsors&logoColor=white">
    </a>
</p>

> :rocket: Don't hesitate to follow us on our different networks, and put a star 🌟 on `PoC's` repositories

> Made with :heart: by PoC