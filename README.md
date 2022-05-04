> If you see this section, you've just created a repository using [PoC Innovation's Open-Source project template](https://github.com/PoCInnovation/open-source-project-template). Check the [getting started guide](./.github/getting-started.md).

# SecureEnv

Provide a secure environment for all developers.

## How does it work?

Secret are store in a vault and our go program get these secret for update your environement variables

## Getting Started


### Installation

For install golang-tar [https://go.dev/doc/install]

```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

go version

go get github.com/hashicorp/vault/api

mkdir -p $GOPATH/src/github.com/hashicorp && cd $_

git clone https://github.com/hashicorp/vault.git

cd vault

make bootstrap

make dev

vault -h
```

### Quickstart


```
vault server -dev
```
get the hash from the vault and insert it in .env
```
ADRESS=http://127.0.0.1:8200

HASH=hvs.PA3O9TDvERLIp7U9fNqM9uG8

PATHREAD=${PATH}
```
sure you have to install docker and run the Dockerfile like this

```
sudo docker build . -t ${NAME}

sudo docker run -d --network="host"  --name secureEnv ${NAME}
```


## Get involved

You're invited to join this project ! Check out the [contributing guide](./CONTRIBUTING.md).

If you're interested in how the project is organized at a higher level, please contact the current project manager.

## Our PoC team :heart:

Developers
| [<img src="https://github.com/tonida-rodda.png?size=85" width=85><br><sub>[Toni Da-rodda]</sub>](https://github.com/tonida-rodda) | [<img src="https://github.com/florianepitech.png?size=85" width=85><br><sub>[Florian Damiot]</sub>](https://github.com/florianepitech) | <br></br>
| :---: | :---: | :---: |

Manager
| [<img src="https://github.com/adrienfort.png?size=85" width=85><br><sub>[Adrien Fort]</sub>](https://github.com/adrienfort)
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