### `evaultz`

THIS PROJECT IS STILL IN THE INITIAL PHASE, STAY TUNED

### `Description`

# Welcome to evaultz app 👋

This is an file storage app, we can store any file like docs, videos, pdf etc.

### `Installation` ✔

Make sure you have latest GoLang installed on your workstation, and you can download the code and use it directly on any IDE

### `Usage` ▶

- Create a config.yaml file, example as below. Make sure you don't store this config file in a public repository or network share for security reasons
  You can store the passwords in an external secret management vault like AKeyless Vault or Hashicorp Vault

```bash
env: "dev"
cm_secret:
  base_url: "https://yourciphertrustip.com/api/"
  version: "v1"
  cm_user: "Your CipherTrust Manager Username"
  cm_password: "Your CipherTrust Manager User Password"
  encryption_key: "Your CM Encryption Key"
akeyless_secret:
  username: "Your Akeyless console email (https://console.akeyless.io/)"
  password: "Your Akeyless console password"
```

To run the program you need to execute the run command in the below format

```bash
go run cmd/evaultz/main.go -configfile config.yaml
```

Build the program into an executable file for your specific platform, Example below for Mac/Linux

```bash
go build -o evaultz cmd/evaultz/main.go -configfile config.yaml
```

### `Contributing` 🤝

Contributions, issues and feature requests are welcome, but it is paused for now.
Feel free to check issues page if you want to contribute in future
